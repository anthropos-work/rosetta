# Stack Update — Technical Reference

Quick reference for update commands and error recovery. Full instructions: `corpus/ops/update_guide.md`.
Targets the **dev** side (main dev stack or a `dev-N`); demo stacks are reproduced by teardown + bring-up at
a pinned tag, not updated in place.

## File references

| Type | Path |
|------|------|
| Master Guide | `corpus/ops/update_guide.md` |
| Ops Reports | `stack-dev/ops-reports/` |
| Working Dir | `stack-dev/` (or the target `dev-N` stack dir) |

## Update scenarios

| Scenario | Use When | Scope |
|----------|----------|-------|
| Daily | Starting daily work | Main repos, frontend deps |
| Weekly | Weekly maintenance, after being away | All repos, all deps, migrations |
| Full | Breaking changes, major version bumps | Fresh pull, clean install, rebuild |

## Quick state check

```bash
# Running services (main dev stack):
docker ps --filter "name=anthropos-" --format "table {{.Names}}\t{{.Status}}"

# Git status across repos — the platform clone + the 4 repos.yml repos @ 0c91421.
# (This loop used to name `cms` and `jobsimulation` and omit `sentinel` + `studio-desk`. Both of those
#  are frozen legacy since the merges; `make init` no longer clones them, so they printed "Not found"
#  while two repos that DO need updating were never checked.)
cd stack-dev
for repo in platform app sentinel next-web-app studio-desk; do
  echo "=== $repo ==="
  (cd "$repo" && git status -s) 2>/dev/null || echo "  Not found"
done

# Pending migrations:
(cd app && atlas migrate status --env local)
```

## Common update commands

```bash
# Stop services before pulling (in stack-dev/platform):
make down

# Pull all repos (auto-stashes dirty changes):
make pull

# Update dependencies:
(cd next-web-app && pnpm install)
(cd studio-desk && npm install)
(cd ant-academy/code && npm install)   # internal learning portal, native only

# Apply migrations — `app` is the ONLY repo with any (`repos.yml` @ 0c91421 marks exactly one
# `migrations: true`, schema `public`). The old "4 services: app, cms, jobsimulation, skillpath"
# comment was wrong on 3 of its 4: cms/jobsimulation are `migrations: false` with their `schema:`
# keys deleted, and skillpath is not in repos.yml at all.
make migrate

# Rebuild and start:
make up
```

## Error recovery

### Git conflict during pull
`make pull` auto-stashes dirty changes. If a conflict remains, resolve it in the affected repo, then re-run.
(Do NOT discard the user's local changes without confirmation.)

### Package not found after pull
```bash
(cd next-web-app && pnpm install)     # frontend
(cd studio-desk && npm install)       # studio-desk
# If still failing: rm -rf node_modules && pnpm install
```

### Migration failed
```bash
(cd app && atlas migrate status --env local)
# If a dirty state needs a fresh DB (DATA LOSS — confirm first):
make reset-db
```

### Docker build fails
```bash
docker system prune -f                 # clear cache
# then rebuild from local code:
make up
```

### Service won't start after update
```bash
make logs S=[service]
# Common causes: missing env vars (platform/.env) · schema mismatch (apply migrations) · stale image (make up).
```

## Repository update order (the make-driven `pull` order)

`make pull` walks `repos.yml` in file order, so the whole list is **four** repos plus the `platform`
clone you run it from (`Makefile:31-45`, `repos.yml` @ platform `0c91421`):

```
1. platform       (Docker configs, shared .env — the repo you are standing in)
2. app            (the monolith: skiller, skillpath, roadrunner, jobsimulation, cms,
                   storage, messenger and customerio-sync all served in-process)
3. sentinel       (authorization — the one other Go service still deployed)
4. next-web-app   (frontend)
5. studio-desk
```

`ant-academy` is **not** in `repos.yml` by design, so `make pull` never touches it — update it by hand in
whatever clone you made (`git -C ant-academy pull`); it is independent of `backend`, so it is safe to do
last, or not at all.

No longer cloned or orchestrated, and therefore not in this order: `chronos`, `intelligence`, `skiller`
(merged into `app` in July 2026), `skillpath` (M502→M507), `cms` + `jobsimulation` + `roadrunner`
(entries deleted at `d11a403`) and `storage` + `messenger` (deleted at `838d907`). **None of those repos
were deleted** — clone one by hand if you need to read the pre-merge source.

## Updating a stack's tooling (tag bump, not in-place edit)

A `stack-*/` dir consumes `rosetta-extensions` at a pinned tag (`stack-<role>/rosetta-extensions @ <tag>`).
Update tooling by bumping the pinned tag (re-clone / checkout at the new tag) — never hand-edit scripts in
the stack dir. New tooling is authored + TESTED + TAGGED first in `.agentspace/rosetta-extensions/`.

## Ops report template

`stack-dev/ops-reports/op_YYYYMMDD_HHMMSS_update_<topic>.md`:

```markdown
# Ops Report: [Brief Title]
**Date**: YYYY-MM-DD HH:MM
**Skill**: /stack-update
**OS**: [macOS 14.x / Ubuntu 22.04 / etc.]
**Phase**: [Stop / Pull / Dependencies / Migrations / Rebuild / Verify]

## Issue Encountered
[Exact error message]

## Context
[What was being done, what commands ran]

## Resolution
[How fixed, or "Unresolved"]

## Suggested Documentation Update
[What to add/change in update_guide.md]
```

## Related skills

| Skill | Use when |
|-------|----------|
| `/dev-up` | First-time environment build / start a stack after updating |
| `/stack-list` | List live stacks |
| `/update-knowledge` | Process ops-reports into the corpus |
