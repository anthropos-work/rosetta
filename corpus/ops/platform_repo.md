# The `platform` Repo (Dev-Environment Control Plane)

> Reference for the orchestrator repo itself. The step-by-step *guides* live in
> [setup_guide.md](./setup_guide.md), [run_guide.md](./run_guide.md), and
> [update_guide.md](./update_guide.md); this page documents **what the repo contains**
> and **what each Make target / profile / file does**.

## Role & Responsibility

`platform` is **not a deployed service**. It is the dev-environment control plane: a
**Makefile + Docker Compose** orchestrator that clones the ~13 sibling repos and
builds/runs the microservices locally **from source**. It is the one repo you `cd` into to
operate everything else.

* **Repo**: `git@github.com:anthropos-work/platform` → cloned to `stack-dev/platform`
* **Drivers**: GNU Make (`SHELL=/bin/bash`), Docker Compose v2, YAML
* **No application code, no tests, no CI** (orchestration only; per-service tests run inside each cloned repo)

## Repo Layout

```
Makefile            Single entry point for all dev ops (parses repos.yml with awk — no yq/python)
docker-compose.yml  12 app service definitions; `include: [common.yml]`
common.yml          Base infra: postgresql + redis (always-on, no profile); declares app-network
repos.yml           Manifest of repos `make init` clones (name / type / migrations / schema)
postgresql/         Custom Postgres image (Dockerfile: compiles pgvector v0.4.4 onto bitnamilegacy/postgresql:15)
data/               Git-ignored Postgres bind-mount (./data/postgresql → /bitnami/postgresql); wiped by reset-db
.env / .env_example .env = real shared secrets (git-ignored); .env_example = tracked template (names only)
README.md / CLAUDE.md   In-repo docs (Make-target table, profile table, port map)
```

## Make Targets

| Target | What it does |
|--------|--------------|
| `make init` | Clone every repo in `repos.yml` not yet present in `../` from `git@github.com:anthropos-work/<name>.git` |
| `make pull` | Checkout + rebase `main` on all repos, auto-stashing dirty trees |
| `make status` | Per-repo branch / dirty / behind table |
| `make up [PROFILE=…]` | `docker compose --profile $(PROFILE) up --build -d` — **`PROFILE` defaults to `graphql`** |
| `make up-all` | Start every service (profile `all`) |
| `make up-frontend` | Start `next-web-app` together with the graphql backend stack |
| `make down` / `make ps` | Stop all services / list containers |
| `make logs [S=svc]` | Tail compose logs, optionally one service |
| `make migrate [S=svc]` | `atlas migrate apply --env local` across the 5 migration repos, or a single repo via `S=` |
| `make dev S=svc` | Stop a service container and print native-run instructions (`cd ../svc && go run .`) |
| `make build-frontend` | `pnpm install && pnpm build` in `../next-web-app` |
| `make reset-db` | **Confirm-gated** wipe of `data/postgresql/`, restart Postgres, re-migrate (waits on `pg_isready`) |
| `make bootstrap-dev` | End-to-end: up + migrate + seed Sentinel policy (`../sentinel/init_policy.sql`) + create a Clerk/DB admin user & org via `../app/cmd` CLIs (needs Go toolchain + `CLERK_SECRET_KEY`) |
| `make help` | Auto-generated target listing from `## ` doc comments |

> **`make migrate` (bulk)** runs each repo with `|| true` — a single repo's migration failure is logged but does **not** abort the run or fail the target, so scan the output for errors. Use `make migrate S=<repo>` to get a hard (non-zero) failure for one repo.
>
> There is **no** `setup`, `gen`, or `init-studio` target in `platform`. Those live in
> the individual service repos (`make gen`/`make setup` per service; `cd cms && make
> init-studio` embeds `anthropos-studio-room`).

## Compose Profiles

`docker-compose.yml` defines **12 app services**: `graphql`, `sentinel`, `backend`,
`jobsimulation`, `cms`, `skillpath`, `storage`, `customerio-sync`,
`messenger`, `roadrunner`, `studio-desk`, `next-web-app` — plus the third-party
`gotenberg` image and the two base services from `common.yml`. (The former `skiller`
service was merged into `app`/`backend` in July 2026 — its RPC surface is now served
by `backend`, `SKILLER_RPC_ADDR=http://backend:8083` in compose.)

| Profile | Services started (besides always-on `postgresql`, `redis`, `sentinel`) |
|---------|------------------------------------------------------------------------|
| `graphql` *(default)* | backend, jobsimulation, cms, skillpath, storage, roadrunner, gotenberg, **graphql** |
| `backend` | backend, gotenberg |
| `jobsimulation` / `cms` / `skillpath` / `storage` / `roadrunner` | **only that one service** |
| `messenger` | messenger (bring up its deps too: backend/cms/jobsimulation/skillpath) |
| `customerio-sync` | customerio-sync |
| `frontend` | next-web-app (containerized Workforce) |
| `studio-desk` | studio-desk (containerized) |
| `all` | everything |

> **Gotchas:**
> * `sentinel`, `postgresql`, `redis` have **no `profiles:` line** → they start with *every* profile.
> * A **single-service profile does NOT start the `graphql` gateway** (it's only in `graphql`/`all`). `make up PROFILE=cms` gives you cms but no usable `:5050` endpoint.
> * `customerio-sync` is **built from a GitHub URL** (`context: git@github.com:anthropos-work/customerio-sync.git#main`) and is **not** in `repos.yml`, so `make init` never clones it.
> * Every Go service hardcodes build arg `ARCH: arm64` (Apple-Silicon-first) — x86 hosts must override it.
> * All app builds use BuildKit SSH forwarding (`ssh: ["default"]`) + `GH_ACCESS_TOKEN=$GH_PAT` to pull private Go modules — needs a loaded SSH agent **and** `GH_PAT` in `.env`.

Use `docker compose --profile <name> config --services` to confirm a profile's exact members.

## `repos.yml` (what `make init` clones)

Entries with `name` / `type` / `migrations` (+ `schema` for Go services with migrations):

* **Go**: `app` (public), `cms` (cms), `jobsimulation` (jobsimulation), `skillpath` (skillpath) — all `migrations: true`; `sentinel`, `storage`, `messenger`, `roadrunner` — `migrations: false`.
* **Node**: `next-web-app` (node-pnpm), `studio-desk` (node-npm), `ant-academy` (node-npm), `graphql-wundergraph` (node-npm).

> `ant-academy` is cloned but has **no compose service** (runs natively / Vercel). The
> shared libraries (colony, authn, proto, ai, taxonomy) are **not** here — they are pulled
> as Go modules, see [Shared Libraries](../architecture/shared_libraries.md).

## Ports

| Service | Host port(s) |
|---------|--------------|
| postgresql / redis | 5432 / 6379 |
| backend (`app`) | 8081, 8082 (`PORT`), 8083 (RPC — also serves the merged skiller RPC surface) |
| sentinel | 8087 |
| cms | 8090, 8091 (RPC) |
| skillpath | 8100, 8101 (RPC) |
| messenger | 8200, 8201 (RPC) |
| storage | 8300, 8301 (RPC) |
| jobsimulation | 8400 (`PORT`), 8401 (RPC) |
| studio-desk | 9000 (backend), 9100 (frontend) |
| roadrunner | 10400, 10401 (RPC) |
| graphql (WunderGraph/Cosmo) | **5050 → container 8080** |
| next-web-app | 3000 |
| customerio-sync | 8080 |
| gotenberg | 3200 |

## Infrastructure (`common.yml`)

* **PostgreSQL 15** — a **built** image (`postgresql/Dockerfile` compiles **pgvector v0.4.4** onto `bitnamilegacy/postgresql:15`), `ALLOW_EMPTY_PASSWORD=yes`, `pg_isready` healthcheck, data persisted via `./data/postgresql`. Schema isolation by `search_path` per service (sentinel uses `sentinel`; the rest default to `public` — skills data lives in `public` since the skiller→app merge; the old `skiller` schema is legacy).
* **Redis** — `bitnamilegacy/redis:latest`, no password; Watermill streams at `REDIS_STREAMS_INDEX=4` plus per-service worker/recording indexes.

## Environment

Every app service uses `env_file: .env` — a **single centralized secrets file** in this
repo. Copy the template and fill it in (never commit `.env`):

```bash
cd platform
cp .env_example .env
brew install ariga/tap/atlas   # required on the host for migrations
make init && make up && make migrate
open http://localhost:5050     # GraphQL playground
```

Key variables include `GH_PAT` (private Go modules), `CLERK_SECRET_KEY`, the
`AZURE_OPENAI_*` / `OPENAI_KEY` / `ANTHROPIC_API_KEY` / `MISTRAL_API_KEY` AI set,
`LIVEKIT_*`, `BUNNY_*`, `DIRECTUS_*`, and `PUBLIC_HOST` (compose-only; bakes
`NEXT_PUBLIC_*` URLs for remote VMs). Non-secret config baked into `docker-compose.yml`
includes the Judge0 sandbox URL, the LiveKit cloud URL, and the Directus address.

> Two OpenAI keys coexist and are easy to confuse: **`OPENAI_KEY`** (app/jobsim)
> vs **`OPENAI_API_KEY`** (cms). CMS also has its own `CMS_AZURE_OPENAI_*` and `AZURE_API_KEY`.

## Related Documentation

* [Setup Guide](./setup_guide.md) · [Run Guide](./run_guide.md) · [Update Guide](./update_guide.md)
* [Service Taxonomy](../architecture/service_taxonomy.md) · [Shared Libraries](../architecture/shared_libraries.md)
* [GraphQL Gateway](../services/graphql-wundergraph.md) · [Next Web App](../services/next-web-app.md)
