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
make init                 # clone all repos from repos.yml (incl. ant-academy)
cd cms && make init-studio && cd ..   # CMS studio submodule
# PostgreSQL schemas (before migrations):
docker exec anthropos-postgresql-1 psql -U postgres \
  -c "CREATE SCHEMA IF NOT EXISTS extensions; CREATE EXTENSION IF NOT EXISTS vector SCHEMA extensions; CREATE EXTENSION IF NOT EXISTS pg_trgm SCHEMA extensions; CREATE SCHEMA IF NOT EXISTS sentinel;"
make up                   # build from local code + start (graphql profile) — expect 11 containers
make migrate              # apply migrations (app, cms, jobsimulation, skillpath)

# Start / restart an already-built stack:
make up                   # rebuild + start
make ps                   # 11 healthy containers in graphql
```

### Expected service set (default `graphql` profile, main dev stack)

| Container | Port(s) | Notes |
|-----------|---------|-------|
| anthropos-postgresql-1 | 5432 | Health gate for others |
| anthropos-redis-1 | 6379 | Health gate for others |
| anthropos-sentinel-1 | 8087 | Always on (no profile) |
| anthropos-backend-1 | 8081-8083 | Also serves the merged skiller RPC surface |
| anthropos-skillpath-1 | 8100-8101 | |
| anthropos-cms-1 | 8090-8091 | Embedded Python studio-room |
| anthropos-jobsimulation-1 | 8400-8401 | |
| anthropos-storage-1 | 8300-8301 | |
| anthropos-roadrunner-1 | 10400-10401 | |
| anthropos-graphql-1 | 5050 | Cosmo Router |
| anthropos-gotenberg-1 | 3200 | Third-party PDF conversion |

Not in this profile (don't expect running): `messenger`, `customerio-sync` (explicit profile),
`ant-academy` (native-only on port 3077, never in docker-compose). Archived (not orchestrated locally):
`chronos`, `intelligence`, `skiller` (merged into `app`, July 2026).

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
**exits 0**) and `cms` is cut over so content is self-contained; **without** it the recipe is print-only and
the `directus` replay skips with exit 4 — the stack reads content live from prod (the documented fallback).

## Quick health checks

```bash
docker info > /dev/null 2>&1 && echo "Docker OK" || echo "Start Docker"
docker exec anthropos-postgresql-1 pg_isready -U postgres
docker exec anthropos-redis-1 redis-cli ping
curl -s http://localhost:5050/health && echo "GraphQL OK"     # main stack; offset for dev-N
docker ps --filter "name=anthropos-" --format "table {{.Names}}\t{{.Status}}"
```

## Error recovery

### Port already in use
```bash
lsof -i :5050        # find the holder
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
Run `cd cms && make init-studio` before `make up`.

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
