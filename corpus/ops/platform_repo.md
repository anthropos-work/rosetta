# The `platform` Repo (Dev-Environment Control Plane)

> Reference for the orchestrator repo itself. The step-by-step *guides* live in
> [setup_guide.md](./setup_guide.md), [run_guide.md](./run_guide.md), and
> [update_guide.md](./update_guide.md); this page documents **what the repo contains**
> and **what each Make target / profile / file does**.

## Role & Responsibility

`platform` is **not a deployed service**. It is the dev-environment control plane: a
**Makefile + Docker Compose** orchestrator that clones the six sibling repos still in
`repos.yml` and builds/runs the services locally **from source**. It is the one repo you
`cd` into to operate everything else.

* **Repo**: `git@github.com:anthropos-work/platform` → cloned to `stack-dev/platform`
* **Drivers**: GNU Make (`SHELL=/bin/bash`), Docker Compose v2, YAML
* **No application code, no tests, no CI** (orchestration only; per-service tests run inside each cloned repo)

## Repo Layout

```
Makefile            Single entry point for all dev ops (parses repos.yml with awk — no yq/python)
docker-compose.yml  8 app service definitions; `include: [common.yml]`
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
| `make up [PROFILE=…]` | `docker compose --profile $(PROFILE) up --build -d` — **`PROFILE` defaults to `core`** (renamed from `graphql` at v9.0; there is no `graphql` profile any more) |
| `make up-all` | Start every service (profile `all`) — note `all` no longer includes `messenger` or `storage`, see the profile table |
| `make up-frontend` | Start `next-web-app` together with the `core` backend stack (`--profile core --profile frontend`) |
| `make down` / `make ps` | Stop all services / list containers |
| `make logs [S=svc]` | Tail compose logs, optionally one service |
| `make migrate [S=svc]` | `atlas migrate apply --env local`. **`app` is the only `migrations: true` repo** — every folded domain's tables were re-created in `public` under `app/terraform/migrations/`. The three repos still in `repos.yml` beside it (`sentinel`, `storage`, `messenger`) are all `migrations: false` |
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

`docker-compose.yml` defines **8 app services**: `sentinel`, `backend`, `storage`,
`customerio-sync`, `messenger`, `studio-desk`, `next-web-app` — plus the third-party
`gotenberg` image and the two base services from `common.yml`. (The `graphql` service is
gone: the WunderGraph/Cosmo router was **retired 2026-07-31**, so `:5050` is free and
clients hit `backend`'s own gqlgen endpoint.)

> **Eight services were folded into `backend`.** `skiller`, `skillpath`, `roadrunner`,
> `jobsimulation` (jobsim-in-app) and `cms` (cms-in-app v8.0) all run in-process inside
> `app`, and **v9.0 "support-in-app" (2026-08-04) added `messenger`, `storage` and
> `customerio-sync`**. `sentinel` is now the **only out-of-process Anthropos service** on a
> default stack.
>
> The first five have no compose service left at all. The v9.0 three still have one, but
> **not in any default profile** — they are kept startable purely as the rollback path
> (`storage-legacy` / `messenger` / `customerio-sync`).
>
> `backend` has **no `*_RPC_ADDR` loopbacks**, and since v9.0 its RPC mux has **no external
> callers left**. The `BACKEND_USERS_RPC_ADDR` / `CMS_RPC_ADDR` / `JOBSIMULATION_RPC_ADDR` /
> `SKILLER_RPC_ADDR=http://backend:8083` env block lives in the **`messenger` service
> definition only** — i.e. it is the standalone rollback container's wiring, not something a
> default stack exercises.

(The former `skiller` service was merged into `app`/`backend` in July 2026 — its RPC
surface is now served by `backend`. The former `skillpath` service was likewise merged
into `app`/`backend` — "skillpath-in-app", M502→M507 — and is **gone from compose**; only
the residual `SKILLPATH_STREAM=skillpath` env plumbing remains.)

| Profile | Services started (besides always-on `postgresql`, `redis`, `sentinel`) |
|---------|------------------------------------------------------------------------|
| `core` *(default)* | backend, gotenberg — **renamed from `graphql` at v9.0**; with the router retired and storage folded in, the two members are all that is left |
| `backend` | backend, gotenberg — **identical membership to `core`** since v9.0 |
| `storage-legacy` | storage — **rollback only**, never started by default (`app` serves storage in-process; running both means two writers on one bucket) |
| `messenger` | messenger — **rollback only**. **Never run it alongside a `MESSENGER_ENABLED=true` backend**: `app` takes over messenger's *own* Redis consumer group (the literal `messenger`), so both running = two consumers on one group |
| `customerio-sync` | customerio-sync — **still in `all`** (see below) |
| `frontend` | next-web-app (containerized Workforce) |
| `studio-desk` | studio-desk (containerized) |
| `all` | backend, gotenberg, customerio-sync, studio-desk, next-web-app. **`messenger` was dropped from `all` at v9.0** (the consumer-group clash above) and `storage` is `storage-legacy`-only — so `make up-all` no longer means "everything declared" |

> **Gotchas:**
> * `sentinel`, `postgresql`, `redis` have **no `profiles:` line** → they start with *every* profile.
> * There is **no gateway container to miss** any more — `backend` serves GraphQL itself at `http://localhost:8082/graphql/query`. `:5050` is free.
> * `customerio-sync` is **built from a GitHub URL** (`context: git@github.com:anthropos-work/customerio-sync.git#main`) and is **not** in `repos.yml`, so `make init` never clones it. That is unchanged by v9.0 — but the job it runs is now `backend`'s (`internal/customeriosync`, on app's asynq scheduler, gated by `CUSTOMERIO_SYNC_ENABLED`), so the container only matters if you deliberately select its profile. Its destination is **Brevo**, not Customer.io — the name is a fossil.
> * `backend`'s compose `environment:` hardcodes `STORAGE_S3_BUCKET=production-storage20240826131618541000000005` and `STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001` — **a default local stack points at the REAL PRODUCTION buckets**. `STORAGE_RPC_ADDR` is gone (read by no code).
> * `MESSENGER_ENABLED` and `CUSTOMERIO_SYNC_ENABLED` are **not** in `.env_example`. Unset ⇒ **off** on a developer machine, which is the intended local default; unset in a *deployed* environment ⇒ `backend` **refuses to boot**, and an unparseable value is an error everywhere (`app/env_guards.go`). `BREVO_KEY` is **required** whenever either switch is on — `backend` fails fast on an empty key.
> * Every Go service hardcodes build arg `ARCH: arm64` (Apple-Silicon-first) — x86 hosts must override it.
> * All app builds use BuildKit SSH forwarding (`ssh: ["default"]`) + `GH_ACCESS_TOKEN=$GH_PAT` to pull private Go modules — needs a loaded SSH agent **and** `GH_PAT` in `.env`.

Use `docker compose --profile <name> config --services` to confirm a profile's exact members.

## `repos.yml` (what `make init` clones)

Entries with `name` / `type` / `migrations` (+ `schema` for Go services with migrations).
**Six entries remain:**

* **Go**: `app` (public) is the only `migrations: true` entry. `sentinel`, `storage` and `messenger` are `migrations: false`. Since v9.0 only `sentinel` is a live out-of-process service — **`storage` and `messenger` are frozen legacy too**, folded into `app`, and they are still listed here **on purpose**: `make init` keeps a clone on disk so the rollback containers can be built. (`cms`, `jobsimulation`, `roadrunner`, `skillpath` and `skiller` are decommissioned and no longer in `repos.yml` at all — clone them by hand to read the pre-merge source.)
* **Node**: `next-web-app` (node-pnpm), `studio-desk` (node-npm).

> `ant-academy` is **not** in `repos.yml` (by design) and has no compose service — it runs
> natively / on Vercel, so clone it yourself. `graphql-wundergraph` is gone from `repos.yml`
> with the router's retirement. The shared libraries (colony, authn, proto, ai, taxonomy)
> are **not** here either — they are pulled as Go modules, see
> [Shared Libraries](../architecture/shared_libraries.md).

## Ports

| Service | Host port(s) |
|---------|--------------|
| postgresql / redis | 5432 / 6379 |
| backend (`app`) | 8081, 8082 (`PORT` — HTTP/GraphQL/`/api/health`), 8083 (RPC — one mux serving `BackendUsers`, `BackendOrganizations`, `SkillerService`, `SkillPathSessionService`, `JobSimulationService`, `CMSService` and `lab.v1.LabSessionService`; **no external caller left since v9.0**). `META_PORT=8084` is set but **not published** to the host |
| sentinel | 8087 — the only out-of-process Anthropos service on a default stack |
| studio-desk | 9000 (backend), 9100 (frontend) |
| next-web-app | 3000 |
| gotenberg | 3200 |

**Rollback-only ports.** These are the *rollback target's* ports, not a default stack's — nothing
listens on them unless you deliberately select the profile:

| Service (profile) | Host port(s) |
|---------|--------------|
| messenger (`messenger`) | 8200, 8201 (RPC) |
| storage (`storage-legacy`) | 8300, 8301 (RPC) |
| customerio-sync (`customerio-sync`) | 8080 |

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
open http://localhost:8082/graphql/query   # backend's own GraphQL endpoint (:5050 is retired)
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
