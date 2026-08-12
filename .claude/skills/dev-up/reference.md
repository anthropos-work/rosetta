# Dev Up — Technical Reference

Quick reference for verification, health checks, and error recovery for the consolidated dev bring-up.
Full instructions: `corpus/ops/setup_guide.md` (first-time build) + `corpus/ops/run_guide.md` (start/health).

## File references

| Type | Path |
|------|------|
| Build guide | `corpus/ops/setup_guide.md` |
| Start guide | `corpus/ops/run_guide.md` |
| Registry + offset ports | `corpus/ops/rosetta_demo.md` |
| Set-dress (snapshot + seed) | `corpus/ops/snapshot-spec.md` + `corpus/ops/seeding-spec.md` |
| Ops reports | `stack-dev/ops-reports/` |
| Working dir | `stack-dev/` |

## Prerequisites (verify before the first build)

```bash
git --version
docker --version && docker compose version
go version
node --version && pnpm --version   # Node must be v24+ (next-web-app engines.node ">=24.0.0")
python3 --version
atlas version
ssh -T git@github.com               # GitHub SSH (run /setup-github if this fails)
```

## Mode A — main dev stack (N=0): build + start

```bash
# First-time build (in stack-dev/platform):
make init                 # clone the 3 repos in repos.yml: app, next-web-app, studio-desk
                          # (ant-academy is NOT in repos.yml by design — clone it by hand if you need it;
                          #  the old `cd cms && make init-studio` step is dead: cms is not cloned.)
git clone git@github.com:anthropos-work/anthropos-studio-room.git app/studio
                          # REQUIRED for N=0 (M257x iter-262/270). "studio-room ships inside the app
                          # image" is CI's story and it inverts the causality locally: app/Dockerfile
                          # hard-COPYs studio/ + pip-installs it, so a local `make up` DIES with
                          # `"/build/studio": not found` if the tree is absent. `dev-stack up N`
                          # (N >= 1) does this for you; the make-driven N=0 path does not.
# PostgreSQL schemas (before migrations):
docker exec anthropos-postgresql-1 psql -U postgres \
  -c "CREATE SCHEMA IF NOT EXISTS extensions; CREATE EXTENSION IF NOT EXISTS vector SCHEMA extensions; CREATE EXTENSION IF NOT EXISTS pg_trgm SCHEMA extensions; CREATE SCHEMA IF NOT EXISTS sentinel;"
make up                   # build from local code + start (core profile — the default) — 4 containers, not 11
make migrate              # apply migrations — `app` is the ONLY repo that has any (repos.yml)

# Start / restart an already-built stack:
make up                   # rebuild + start
make ps                   # expect 5 — the old "11" is three merge waves stale (see the service-set table below)
```

### Expected service set (default `core` profile, main dev stack)

| Container | Port(s) | Notes |
|-----------|---------|-------|
| anthropos-postgresql-1 | 5432 | Health gate for others |
| anthropos-redis-1 | 6379 | Health gate for others |
| ~~anthropos-sentinel-1~~ | ~~8087~~ | **GONE since platform `766df6c`** — the Casbin PDP is `app/internal/sentinel/`, in-process |
| anthropos-backend-1 | 8081-8083 | **The monolith.** Serves the **seven** merged skiller / skillpath / cms / jobsimulation / storage / messenger / customerio-sync domains in-process, **and GraphQL itself at `:8082/graphql/query`**. (`roadrunner` was listed here as an eighth domain until M257x iter-137 — it was *deleted*, not merged; Judge0 is reached from inside the jobsimulation domain) |
| anthropos-gotenberg-1 | 3200 | Third-party PDF conversion |

> **⚠️ FOUR containers since platform `766df6c` (was five; `sentinel` folded into `app` at v11.0) — re-derived from the platform clone at origin
> `0c91421`** (`docker-compose.yml`, 186 lines; M257x iter-87). Compose declares only **5** services
> (`sentinel`, `backend`, `studio-desk`, `next-web-app`, `gotenberg`) plus `postgresql` + `redis` from
> `common.yml`, and `core` selects five of those seven. Everything this table used to list is gone from
> compose, in three waves: **(1)** `anthropos-skillpath-1` at platform M507 and `anthropos-graphql-1`
> (`:5050`) at `2adcf71` — nothing listens on `:5050`, GraphQL is `backend`'s own `:8082/graphql/query`
> (`/graphql` serves the Apollo Sandbox UI); **(2)** the cms / jobsimulation / roadrunner husks, which
> still existed as no-traffic rollback containers at `2adcf71` (where this note's previous revision
> measured them) and no longer exist at `0c91421`; **(3)** the `storage`, `messenger` and
> `customerio-sync` services at `838d907` (2026-08-05, PR #26 *"drop the support-service containers"*),
> which also removed `storage` + `messenger` from `repos.yml`. **Do not trust a container count from
> memory** — the old "11 healthy containers" figure is three waves stale. See [`corpus/architecture/platform-migration-status.md`](../../../corpus/architecture/platform-migration-status.md).

Not in this profile (don't expect running): `next-web-app` (`frontend`/`all`) and `studio-desk`
(`studio-desk`/`all`) — the only two profile-gated services left; `ant-academy` (native-only on port
3077, never in docker-compose). **No local container at all** (merged into `app`): `cms`,
`jobsimulation`, `roadrunner`, `storage`, `messenger`, `customerio-sync`, `skiller`, `skillpath` —
asking for one of their retired profile tokens **exits 0 and starts only the floor**, so the stack
looks alive with the application absent. Archived: `chronos`, `intelligence`.

## Mode B — additional dev-N (N ≥ 1): bring up + set-dress

```bash
DEV=stack-dev/rosetta-extensions/dev-stack
"$DEV/dev-stack" up N                 # allocate N (unified registry) + offset-port bring-up + set-dress
"$DEV/dev-stack" up N --no-snapshot      # seed only
"$DEV/dev-stack" up N --no-setdress      # bare bring-up
"$DEV/dev-stack" up N --local-content    # EXECUTE a per-stack Directus (dev opt-in; content self-contained)
"$DEV/dev-stack" status                  # list live dev-N (or /stack-list for dev + demo)
```

`dev-N` maps host port `P → P + N*OFFSET` (default offset 10000, shared with demo-stack). The set-dress
pass (`dev-setdress.sh`) is default-on + non-fatal: cache-first snapshot replay (`taxonomy` lands) →
`dev-min` seed (~1 org, ~10 users, fixed admin `dev@anthropos.test`), plus the per-stack-Directus firewall
check. The per-stack Directus is **opt-in for dev** via `--local-content` (v1.5 M22/M23): **with** it the
recipe is EXECUTED (bootstrap → apply-structure → replay → boot the offset-port Directus, `directus` replay
**exits 0**) and `backend`'s cms domain is cut over so content is self-contained; **without** it the recipe is print-only and
the `directus` replay skips with exit 4 — the stack reads content live from prod (the documented fallback).

## Quick health checks

```bash
docker info > /dev/null 2>&1 && echo "Docker OK" || echo "Start Docker"
docker exec anthropos-postgresql-1 pg_isready -U postgres
docker exec anthropos-redis-1 redis-cli ping
curl -s http://localhost:8082/health && echo "GraphQL OK"     # main stack; offset for dev-N
# NB: NOT :5050 — nothing listens on :5050. The Cosmo router was deleted from compose at
# platform 2adcf71, and GraphQL is now backend's own :8082/graphql/query. Curling :5050
# fails against a port with no listener, and so does every other retired port: cms 8090/8091,
# jobsimulation 8400/8401 and roadrunner (d11a403); storage 8300/8301, messenger 8200/8201
# and customerio-sync 8080 (838d907).
docker ps --filter "name=anthropos-" --format "table {{.Names}}\t{{.Status}}"
```

## Error recovery

### Port already in use
```bash
lsof -i :8082        # find the holder (was :5050 before the router was deleted at platform 2adcf71)
# kill -9 <PID> (ask the user first) — or bring the stack up as dev-N on offset ports.
```

### Missing pgvector extension (migrations fail: `schema 'extensions' does not exist`)
```bash
docker exec anthropos-postgresql-1 psql -U postgres \
  -c "CREATE SCHEMA IF NOT EXISTS extensions; CREATE EXTENSION IF NOT EXISTS vector SCHEMA extensions;"
make migrate
```

### Sentinel crash-loops (`pq: no schema has been selected`)
```bash
docker exec anthropos-postgresql-1 psql -U postgres -c "CREATE SCHEMA IF NOT EXISTS sentinel;"
# then restart sentinel
```

### CMS Docker build fails (`"/studio": not found`)
Obsolete since cms-in-app: there is no `cms` service to build and `make init` does not clone the repo,
so `cd cms` fails outright. `studio-room` is pulled into the `app` image by CI, not by a local submodule.

### Docker build fails (SSH / private Go modules)
```bash
ssh-add -l || { eval "$(ssh-agent -s)"; ssh-add ~/.ssh/id_ed25519; }
ssh -T git@github.com
# Builds pull private Go modules via GH_PAT/GOPRIVATE — confirm GH_PAT is set in platform/.env.
```

### Set-dress pass reported a stale/missing snapshot cache
Non-fatal — the seed still ran. To set-dress fully, capture/refresh the snapshot then re-run:
`/stack-snapshot dev-N replay` (see `corpus/ops/snapshot-spec.md`).

## Ops report template

`stack-dev/ops-reports/op_YYYYMMDD_HHMMSS_devup_<topic>.md`:

```markdown
# Ops Report: [Brief Title]
**Date**: YYYY-MM-DD HH:MM
**Skill**: /dev-up
**OS**: [macOS 14.x / Ubuntu 22.04 / etc.]
**Phase**: [Prerequisites / Repos / Docker / Migrations / Start / Set-dress]

## Issue Encountered
[Exact error message]

## Context
[What was being done, what commands ran]

## Resolution
[How fixed, or "Unresolved"]

## Suggested Documentation Update
[What to add/change in setup_guide.md / run_guide.md]
```

## Related skills

| Skill | Use when |
|-------|----------|
| `/dev-down` | Stop / reclaim a dev stack |
| `/stack-update` | Sync code/deps/schemas before running |
| `/stack-list` | List live dev + demo stacks |
| `/update-knowledge` | Process ops-reports into the corpus |
