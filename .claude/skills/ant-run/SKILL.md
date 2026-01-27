---
name: ant-run
description: Start and manage the Anthropos platform locally with health verification
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
- Starting Docker services
- Starting frontend/studio servers
- Health checks and verifications

**ASK for confirmation before**:
- Stopping or restarting services
- Killing processes (port conflicts)
- Data-destructive operations (`docker compose down -v`)
- Changing ports or environment configuration

## Error Handling

1. Do NOT skip errors - resolve first
2. Check logs: `docker compose -p ant-rosetta logs [service]`
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
**Phase**: [Docker / Infrastructure / Backend / Frontend]

## Issue Encountered
[Exact error message or problem description]

## Context
[What step was being executed, what commands were run]

## Resolution
[How it was fixed, or "Unresolved" if still broken]

## Suggested Documentation Update
[What should be added/changed in run_guide.md]
```

3. These reports feed into `/ant-integrate` for corpus improvements

## Progress Tracking

Use TodoWrite with these phases:
- Docker environment verified
- Infrastructure running (PostgreSQL, Redis)
- Backend services running
- GraphQL gateway healthy
- Frontend accessible
- Studio-Desk running

## Critical Rules

- Work in `anthropos-dev/` only
- Use `-p ant-rosetta` for Docker isolation
- Verify health after every start
- Follow the guide - don't improvise

## Additional Resources

- For complete startup instructions, read `corpus/ops/run_guide.md`
- For technical reference (health check patterns, error recovery), see [reference.md](reference.md)
