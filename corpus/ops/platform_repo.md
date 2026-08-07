# The `platform` Repo (Dev-Environment Control Plane)

> Reference for the orchestrator repo itself. The step-by-step *guides* live in
> [setup_guide.md](./setup_guide.md), [run_guide.md](./run_guide.md), and
> [update_guide.md](./update_guide.md); this page documents **what the repo contains**
> and **what each Make target / profile / file does**.

## Role & Responsibility

`platform` is **not a deployed service**. It is the dev-environment control plane: a
**Makefile + Docker Compose** orchestrator that clones the 4 sibling repos and
builds/runs the services locally **from source**. It is the one repo you `cd` into to
operate everything else. (The clone set shrank to 4 at `838d907` — see the `repos.yml`
section below.)

* **Repo**: `git@github.com:anthropos-work/platform` → cloned to `stack-dev/platform`
* **Drivers**: GNU Make (`SHELL=/bin/bash`), Docker Compose v2, YAML
* **No application code, no tests, no CI** (orchestration only; per-service tests run inside each cloned repo)

## Repo Layout

```
Makefile            Single entry point for all dev ops (parses repos.yml with awk — no yq/python)
docker-compose.yml  5 app service definitions (186 lines); `include: [common.yml]`
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
| `make up [PROFILE=…]` | `docker compose --profile $(PROFILE) up --build -d` — **`PROFILE` defaults to `core`** (`Makefile:10`, `PROFILE ?= core`) |
| `make up-all` | Start every service (profile `all`) |
| `make up-frontend` | Start `next-web-app` together with whatever the default profile selects — `Makefile:120` passes `--profile core --profile frontend`, i.e. `backend` + `gotenberg` + the always-on floor, plus the frontend container |
| `make down` / `make ps` | Stop all services / list containers |
| `make logs [S=svc]` | Tail compose logs, optionally one service |
| `make migrate [S=svc]` | `atlas migrate apply --env local`. **`app` is the only migration repo now** — the cms and jobsim tables were re-created in `public` under `app/terraform/migrations/`, and `d11a403` then removed both repos from `repos.yml` outright |
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

`docker-compose.yml` defines **5 services** at platform `0c91421`: `sentinel`, `backend`,
`studio-desk`, `next-web-app` and the third-party `gotenberg` image — **7 in the effective
topology**, once `include: common.yml` adds the two always-on base services (`postgresql`,
`redis`). The `graphql` service was **deleted** at platform `2adcf71`, and the `storage`,
`messenger` and `customerio-sync` service definitions were **deleted** at `838d907`
(PR #26, 2026-08-05) — all three are now served in-process by `backend`. Corrected M257x
iter-78 + iter-87, and fenced by `platform_predicate_guard` G10.

> **SEVEN services were folded into `backend`** (this said **eight** and included `roadrunner` until
> M257x iter-137 — roadrunner was **deleted, not folded**; no `app/internal/roadrunner/` at any ref).
> `skiller`, `skillpath`,
> `jobsimulation` (jobsim-in-app), `cms` (cms-in-app v8.0), `storage` + `messenger`
> (v9.0 support-in-app) and `customerio-sync` all run in-process inside `app`; their
> compose services and profiles are gone, and the federation composes a single `backend`
> subgraph. **Compose sets no `*_RPC_ADDR` variable at all** — the deleted `messenger`
> block was the last thing that set one. The only service address `docker-compose.yml`
> still sets is `AUTHORIZATION_ADDRESS=http://sentinel:8087` (`docker-compose.yml:48`),
> so `backend → sentinel` is the single cross-process RPC edge of a local stack.

(The former `skiller` service was merged into `app`/`backend` in July 2026 — its RPC surface
is now served by `backend`. `SKILLER_RPC_ADDR` is set **nowhere** in compose: the `messenger`
service block was the only thing that set it, and `838d907` deleted that block together with
`BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`. The former `skillpath`
service was likewise merged into `app`/`backend` — "skillpath-in-app", M502→M507 — and is
**gone from compose**; only the residual `SKILLPATH_STREAM=skillpath` env plumbing remains
(`docker-compose.yml:72`).)

| Profile | Services started (besides always-on `postgresql`, `redis`, `sentinel`) |
|---------|------------------------------------------------------------------------|
| `core` *(default — `PROFILE ?= core`)* | backend, gotenberg |
| `backend` | backend, gotenberg |
| `all` | backend, gotenberg, next-web-app, studio-desk |
| `frontend` / `studio-desk` | **selecting one alone exits 1** — each service declares `depends_on: backend`, which its own profile does not select, so compose rejects the project as invalid. Combine with `core` |

**Retired tokens.** `graphql` (*renamed* to `core` at `0dab54d`), `storage`, `cms`, `jobsimulation`
and `roadrunner` are not profiles any more; `storage-legacy`, `customerio-sync` and `messenger`
were removed at `838d907`, when the three containers themselves were deleted. Asking for any of
them **exits 0**, starting only the always-on floor. Deliberately not spelled above in runnable
form — a copy-pasteable command for a silent no-op is the defect.

> **Gotchas:**
> * `sentinel`, `postgresql`, `redis` have **no `profiles:` line** → they start with *every* profile.
> * **There is no `graphql` gateway service any more.** Platform `2adcf71` (2026-07-31) deleted the Cosmo/WunderGraph router from `docker-compose.yml` outright — service, `repos.yml` entry and clone. GraphQL is served by **`backend` itself at `:8082/graphql/query`** (the `/graphql` path serves the Apollo Sandbox UI). The `graphql` **profile name did not survive either** — `0dab54d` renamed it to `core` — and asking for the retired token **exits 0**, starting only the always-on floor (`postgresql`, `redis`, `sentinel`), so nothing about the profile wiring warns you. `make up PROFILE=backend` therefore *does* give you a usable GraphQL endpoint — on `:8082`, not the retired `:5050`.
> * `customerio-sync` **is no longer a compose service** — `838d907` deleted it, along with `storage` and `messenger`; `backend` serves all three in-process. While it existed it was the one service built straight from a GitHub URL (`context: git@github.com:anthropos-work/customerio-sync.git#main`) rather than from a local clone, which is why it was never in `repos.yml`.
> * `backend` hardcodes build arg `ARCH: arm64` (Apple-Silicon-first) — x86 hosts must override it. It is now the **only** `ARCH` in the file (`docker-compose.yml:37`): `838d907` deleted the `storage` and `messenger` blocks, which carried the other two, and `sentinel` never had one.
> * All app builds use BuildKit SSH forwarding (`ssh: ["default"]`) + `GH_ACCESS_TOKEN=$GH_PAT` to pull private Go modules — needs a loaded SSH agent **and** `GH_PAT` in `.env`.

Use `docker compose --profile <name> config --services` to confirm a profile's exact members.

## `repos.yml` (what `make init` clones)

**Four entries** at `0c91421`, each with `name` / `type` / `migrations` (+ `schema` for Go services with migrations):

* **Go**: `app` — `migrations: true`, `schema: public`, and the only migrating repo. `sentinel` — `migrations: false`, yet very much alive. **The second clause of what this line used to say — *"rather than going through `atlas`"* — is FALSE since 2026-08-04** (corrected M257x iter-130): `68272003` added a **second Atlas pipeline** and it lives in **`app`**, not in `sentinel`. `app/atlas.hcl:50-64` declares `env "sentinel"` (`revisions_schema = "sentinel"`, `dir = file://terraform/migrations-sentinel`, `src = file://terraform/sentinel/schema.sql`), and `app/Makefile:59-60` states that *"`atlas migrate apply --env sentinel` creates the schema itself, and that is what local/CI actually run"* (@ `ad9f3c498`). `migrations: false` is still right **about the `sentinel` repo** — it carries no `atlas.hcl` and no `terraform/migrations/` at `f2c461903` — which is exactly why the flag and the pipeline can both be true at once.
* **Node**: `next-web-app` (node-pnpm), `studio-desk` (node-npm).

> **What left the clone set, and when.** `intelligence` (`fdfa189`) · `chronos` (`045857c`) ·
> `skiller` (`21429b7`) · `skillpath` (`a4db680`) · `graphql-wundergraph` (`360efd4`) ·
> `cms` + `jobsimulation` + `roadrunner` (`d11a403`) · `storage` + `messenger` (`838d907`).
> **None of those repos were deleted** — `make init` simply no longer clones them; clone one
> by hand if you need to read the pre-merge source. `customerio-sync` was never an entry at
> all (it built from a git URL), and neither was `ant-academy`, which is cloned by hand or by
> the demo bring-up and has **no compose service** (runs natively / Vercel).
>
> The shared libraries are **not** here either — they are pulled as Go modules, see
> [Shared Libraries](../architecture/shared_libraries.md). **Two lists, and they are not the same
> list:** the doc's five *subjects* are `colony`, `authn`, `proto`, `ai`, `taxonomy` (historical),
> while the five modules a stack actually **imports** are `analytics-go`, `colony`, `proto`,
> `storage`, `taxonomy` (`app/go.mod:14-18` @ `app` `ad9f3c498`). `ai` was folded into `app`
> in-tree and `authn` ships inside colony, so neither is pulled; `analytics-go` and `storage` are
> pulled and were in neither list here until M257x iter-133.

## Ports

| Service | Host port(s) |
|---------|--------------|
| postgresql / redis | 5432 / 6379 |
| backend (`app`) | 8081, 8082 (`PORT`), 8083 (RPC — one mux serving **six** Connect handlers: `BackendUsers`, `BackendOrganizations`, `SkillerService`, `JobSimulationService`, `CMSService` and `lab.v1.LabSessionService`. **There is no `SkillPathSessionService`** — skillpath-in-app M506 *removed* the RPC rather than re-hosting it; likewise no RoadRunner service, `backend` calling Judge0 over plain HTTP), 8084 (`META_PORT`) |
| sentinel | 8087 |
| studio-desk | 9000 (backend), 9100 (frontend) |
| next-web-app | 3000 |
| gotenberg | 3200 |

> **Ports that no longer exist on a local stack.** `graphql` (WunderGraph/Cosmo) `5050 → container
> 8080`, deleted at platform `2adcf71` — GraphQL is now `backend`'s `8082/graphql/query`. And
> `customerio-sync` `8080`, `messenger` `8200/8201` and `storage` `8300/8301`, all deleted at
> `838d907` — those three surfaces run in-process inside `backend`, on no port of their own.

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
