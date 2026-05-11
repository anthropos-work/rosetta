---
name: ant-run
description: Start and manage the Anthropos platform locally with health verification. Use when you are asked to start or restart the anthropos platform locally.
argument-hint: [scenario or 'full']
---

# Anthropos Platform Run

Start the Anthropos platform by following `corpus/ops/run_guide.md` with systematic health verification.

## Your Mission

1. **Read the guide**: `corpus/ops/run_guide.md` is your source of truth
2. **Apply RUN STEP principles**: Check before, verify after, handle conflicts
3. **Track progress**: Use TodoWrite for each phase
4. **Report issues**: Create ops reports for errors and fixes discovered

## RUN STEP Principles

Apply to EVERY step in the guide:

| Principle | Action |
|-----------|--------|
| Check Before Start | Verify if service already running before starting |
| Verify After Start | Confirm health with explicit checks |
| Handle Conflicts | Identify and resolve port conflicts |
| Report Issues | Create ops report when errors or improvements found |

## Confirmation Policy

**Proceed WITHOUT confirmation**:
- Starting Docker services (`make up`)
- Starting frontend/studio servers
- Health checks and verifications

**ASK for confirmation before**:
- Stopping or restarting services
- Killing processes (port conflicts)
- Data-destructive operations (`make reset-db`)
- Changing ports or environment configuration

## Error Handling

1. Do NOT skip errors - resolve first
2. Check logs: `cd anthropos-dev/platform && make logs S=[service]`
3. Create ops report (see below)
4. Continue

## Ops Reports (Instead of Auto-Improving Docs)

When you discover errors, missing steps, or better approaches:

1. Create a report file: `anthropos-dev/ops-reports/op_YYYYMMDD_HHMMSS_run_<topic>.md`
2. Use this template:

```markdown
# Ops Report: [Brief Title]

**Date**: YYYY-MM-DD HH:MM
**Skill**: /ant-run
**OS**: [macOS/Linux/version]
**Phase**: [Docker / Backend / Frontend / Studio]

## Issue Encountered
[Exact error message or problem description]

## Context
[What step was being executed, what commands were run]

## Resolution
[How it was fixed, or "Unresolved" if still broken]

## Suggested Documentation Update
[What should be added/changed in run_guide.md]
```

3. These reports feed into `/ant-document` for corpus improvements

## Progress Tracking

Use TodoWrite with these phases:
- Docker environment verified
- Backend services started (`make up`)
- All containers healthy (`make ps`) — expect 12 containers in default `graphql` profile: postgresql, redis, sentinel, backend, cms, skiller, skillpath, jobsimulation, storage, roadrunner, graphql, gotenberg
- GraphQL gateway healthy (localhost:5050)
- Frontend running (native or containerized; requires Node v24+)
- Studio-Desk running (native or containerized)

## Expected Service Set (default `graphql` profile)

| Container | Port(s) | Notes |
|-----------|---------|-------|
| anthropos-postgresql-1 | 5432 | Healthy gate for others |
| anthropos-redis-1 | 6379 | Healthy gate for others |
| anthropos-sentinel-1 | 8087 | Always on (no profile) |
| anthropos-backend-1 | 8081-8083 | |
| anthropos-skiller-1 | 8085-8086 | |
| anthropos-skillpath-1 | 8100-8101 | |
| anthropos-cms-1 | 8090-8091 | Includes embedded Python studio-room |
| anthropos-jobsimulation-1 | 8400-8401 | |
| anthropos-storage-1 | 8300-8301 | |
| anthropos-roadrunner-1 | 10400-10401 | |
| anthropos-graphql-1 | 5050 | Cosmo Router |
| anthropos-gotenberg-1 | 3200 | Third-party PDF conversion |

Services in `repos.yml` but NOT in this profile (don't expect them running): `messenger`, `customerio-sync` (need explicit profile). Services no longer orchestrated locally: `chronos`, `intelligence` (archived 2026-Q2).

## Critical Rules

- Work in `anthropos-dev/` only
- Use `make` commands for all platform operations
- Verify health after every start
- Follow the guide - don't improvise

## Additional Resources

- For complete startup instructions, read `corpus/ops/run_guide.md`
- For technical reference (health check patterns, error recovery), see [reference.md](reference.md)
