# The `platform` Repo (Dev-Environment Control Plane)

> Reference for the orchestrator repo itself. The step-by-step *guides* live in
> [setup_guide.md](./setup_guide.md), [run_guide.md](./run_guide.md), and
> [update_guide.md](./update_guide.md); this page documents **what the repo contains**
> and **what each Make target / profile / file does**.

## Role & Responsibility

`platform` is **not a deployed service**. It is the dev-environment control plane: a
**Makefile + Docker Compose** orchestrator that clones the 6 sibling repos and
builds/runs the microservices locally **from source**. It is the one repo you `cd` into to
operate everything else.

* **Repo**: `git@github.com:anthropos-work/platform` → cloned to `stack-dev/platform`
* **Drivers**: GNU Make (`SHELL=/bin/bash`), Docker Compose v2, YAML
* **No application code, no tests, no CI** (orchestration only; per-service tests run inside each cloned repo)

## Repo Layout

```
Makefile            Single entry point for all dev ops (parses repos.yml with awk — no yq/python)
docker-compose.yml  11 app service definitions; `include: [common.yml]`
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
| `make migrate [S=svc]` | `atlas migrate apply --env local`. **`app` is the only migration repo now** — the cms and jobsim tables were re-created in `public` under `app/terraform/migrations/`, so `repos.yml` drops both to `migrations: false` |
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

`docker-compose.yml` defines **11 app services**: `graphql`, `sentinel`, `backend`,
`storage`, `customerio-sync`, `messenger`, `studio-desk`, `next-web-app` — plus the
third-party `gotenberg` image and the two base services from `common.yml`.

> **Five services were folded into `backend`.** `skiller`, `skillpath`, `roadrunner`,
> `jobsimulation` (jobsim-in-app) and `cms` (cms-in-app v8.0) all run in-process inside
> `app`; their compose services and profiles are gone, and the federation composes a single
> `backend` subgraph. `backend` also has no `*_RPC_ADDR` loopbacks any more — only
> `messenger` still reaches those surfaces, at `http://backend:8083`.

(The former `skiller`
service was merged into `app`/`backend` in July 2026 — its RPC surface is now served
by `backend`, `SKILLER_RPC_ADDR=http://backend:8083` in compose. The former `skillpath`
service was likewise merged into `app`/`backend` — "skillpath-in-app", M502→M507 — and is
**gone from compose**; only the residual `SKILLPATH_STREAM=skillpath` env plumbing remains.)

| Profile | Services started (besides always-on `postgresql`, `redis`, `sentinel`) |
|---------|------------------------------------------------------------------------|
| `core` *(default — `PROFILE ?= core`)* | backend, gotenberg |
| `backend` | backend, gotenberg |
| `all` | backend, gotenberg, customerio-sync, next-web-app, studio-desk |
| `storage-legacy` | storage — the rollback path only. `docker-compose.yml:130-133`: *"v9.0: NOT in the default profiles any more — app serves storage in-process, and running both means two writers on one bucket. Kept startable for rollback comparison."* |
| `customerio-sync` | customerio-sync |
| `messenger` / `frontend` / `studio-desk` | **selecting one alone exits 1** — each service declares `depends_on: backend`, which its own profile does not select, so compose rejects the project as invalid. Combine with `core` |

**Retired tokens.** `graphql` (*renamed* to `core` at `0dab54d`), `storage`, `cms`, `jobsimulation`
and `roadrunner` are not profiles any more — and asking for one **exits 0**, starting only the
always-on floor. Deliberately not spelled above in runnable form.

> **Gotchas:**
> * `sentinel`, `postgresql`, `redis` have **no `profiles:` line** → they start with *every* profile.
> * **There is no `graphql` gateway service any more.** Platform `2adcf71` (2026-07-31) deleted the Cosmo/WunderGraph router from `docker-compose.yml` outright — service, `repos.yml` entry and clone. GraphQL is served by **`backend` itself at `:8082/graphql/query`** (the `/graphql` path serves the Apollo Sandbox UI). The `graphql` **profile name survives** and still selects the seven-service set, so nothing about the profile wiring warns you. `make up PROFILE=backend` therefore *does* give you a usable GraphQL endpoint — on `:8082`, not the retired `:5050`.
> * `customerio-sync` is **built from a GitHub URL** (`context: git@github.com:anthropos-work/customerio-sync.git#main`) and is **not** in `repos.yml`, so `make init` never clones it.
> * Every Go service hardcodes build arg `ARCH: arm64` (Apple-Silicon-first) — x86 hosts must override it.
> * All app builds use BuildKit SSH forwarding (`ssh: ["default"]`) + `GH_ACCESS_TOKEN=$GH_PAT` to pull private Go modules — needs a loaded SSH agent **and** `GH_PAT` in `.env`.

Use `docker compose --profile <name> config --services` to confirm a profile's exact members.

## `repos.yml` (what `make init` clones)

Entries with `name` / `type` / `migrations` (+ `schema` for Go services with migrations):

* **Go**: `app` (public) is the only `migrations: true` entry. `cms`, `jobsimulation`, `roadrunner`, `sentinel`, `storage`, `messenger` are `migrations: false` — the first three are frozen legacy repos whose tables all live in `app`'s `public` schema. (`skillpath` and `skiller` are decommissioned and no longer in `repos.yml` at all.)
* **Node**: `next-web-app` (node-pnpm), `studio-desk` (node-npm), `ant-academy` (node-npm), `graphql-wundergraph` (node-npm).

> `ant-academy` is cloned but has **no compose service** (runs natively / Vercel). The
> shared libraries (colony, authn, proto, ai, taxonomy) are **not** here — they are pulled
> as Go modules, see [Shared Libraries](../architecture/shared_libraries.md).

## Ports

| Service | Host port(s) |
|---------|--------------|
| postgresql / redis | 5432 / 6379 |
| backend (`app`) | 8081, 8082 (`PORT`), 8083 (RPC — one mux serving **six** Connect handlers: `BackendUsers`, `BackendOrganizations`, `SkillerService`, `JobSimulationService`, `CMSService` and `lab.v1.LabSessionService`. **There is no `SkillPathSessionService`** — skillpath-in-app M506 *removed* the RPC rather than re-hosting it; likewise no RoadRunner service, `backend` calling Judge0 over plain HTTP), 8084 (`META_PORT`) |
| sentinel | 8087 |

| messenger | 8200, 8201 (RPC) |
| storage | 8300, 8301 (RPC) |

| studio-desk | 9000 (backend), 9100 (frontend) |

| ~~graphql (WunderGraph/Cosmo)~~ | ~~5050 → container 8080~~ — **deleted from compose at platform `2adcf71`**; GraphQL is `backend`'s `8082/graphql/query` |
| next-web-app | 3000 |
| customerio-sync | 8080 |
| gotenberg | 3200 |

## Infrastructure (`common.yml`)

* **PostgreSQL 15** — a **built** image (`postgresql/Dockerfile` compiles **pgvector v0.4.4** onto `bitnamilegacy/postgresql:15`), `ALLOW_EMPTY_PASSWORD=yes`, `pg_isready` healthcheck, data persisted via `./data/postgresql`. Schema isolation by `search_path` per service (sentinel uses `sentinel`; everything else defaults to `public`). Since the merges, **all** application data — skills, skill-path sessions, jobsim run state, cms similarity/Studio tables — lives in `public`; the old `skiller`, `skillpath`, `jobsimulation` and `cms` schemas are legacy.
* **Redis** — `bitnamilegacy/redis:latest`, no password; Watermill streams at `REDIS_STREAMS_INDEX=4` plus per-service worker/recording indexes.

## Environment

Every app service uses `env_file: .env` — a **single centralized secrets file** in this
repo. Copy the template and fill it in (never commit `.env`):

```bash
cd platform
cp .env_example .env
brew install ariga/tap/atlas   # required on the host for migrations
make init && make up && make migrate
open http://localhost:8082/graphql   # GraphQL playground (Apollo Sandbox; the endpoint is /graphql/query)
```

Key variables include `GH_PAT` (private Go modules), `CLERK_SECRET_KEY`, the
`AZURE_OPENAI_*` / `OPENAI_KEY` / `ANTHROPIC_API_KEY` / `MISTRAL_API_KEY` AI set,
`LIVEKIT_*`, `BUNNY_*`, `DIRECTUS_*`, and `PUBLIC_HOST` (compose-only; bakes
`NEXT_PUBLIC_*` URLs for remote VMs). Non-secret config baked into `docker-compose.yml`
includes the Judge0 sandbox URL, the LiveKit cloud URL, and the Directus address.

> Two OpenAI keys coexist and are easy to confuse: **`OPENAI_KEY`** (the app/jobsim domains)
> vs **`OPENAI_API_KEY`** (the cms domain). The cms domain also has its own
> `CMS_AZURE_OPENAI_*` and `AZURE_API_KEY`. Since cms-in-app all of these are read by the
> single `backend` process.

## Related Documentation

* [Setup Guide](./setup_guide.md) · [Run Guide](./run_guide.md) · [Update Guide](./update_guide.md)
* [Service Taxonomy](../architecture/service_taxonomy.md) · [Shared Libraries](../architecture/shared_libraries.md)
* [GraphQL Gateway](../services/graphql-wundergraph.md) · [Next Web App](../services/next-web-app.md)
