---
name: ant-update
description: Sync Anthropos code, dependencies, and database schemas with latest changes
argument-hint: [scenario: 'daily' | 'weekly' | 'full']
---

# Anthropos Platform Update

Update the Anthropos platform by following `corpus/ops/update_guide.md` with systematic verification.

## Your Mission

1. **Read the guide**: `corpus/ops/update_guide.md` is your source of truth
2. **Apply UPDATE STEP principles**: Check state, pull before build, verify after
3. **Track progress**: Use TodoWrite for each phase
4. **Report issues**: Create ops reports for errors and fixes discovered

## UPDATE STEP Principles

Apply to EVERY step in the guide:

| Principle | Action |
|-----------|--------|
| Check Current State | Verify what needs updating before making changes |
| Pull Before Build | Always fetch latest code before rebuilding |
| Handle Conflicts | If git conflicts occur, resolve before proceeding |
| Verify After Update | Confirm services still work after updates |

## Confirmation Policy

**Proceed WITHOUT confirmation**:
- Checking git status or service state (`make status`, `make ps`)
- Health checks and verifications

**ASK for confirmation before**:
- Stopping running services (`make down`)
- Pulling repository changes (`make pull`)
- Running database migrations (`make migrate`)
- Rebuilding Docker images (`make up`)
- Data-destructive operations (`make reset-db`)

## Error Handling

1. Do NOT skip errors - resolve first
2. Check logs: `cd anthropos-dev/platform && make logs S=[service]`
3. Create ops report (see below)
4. Continue

## Ops Reports (Instead of Auto-Improving Docs)

When you discover errors, missing steps, or better approaches:

1. Create a report file: `anthropos-dev/ops-reports/op_YYYYMMDD_HHMMSS_update_<topic>.md`
2. Use this template:

```markdown
# Ops Report: [Brief Title]

**Date**: YYYY-MM-DD HH:MM
**Skill**: /ant-update
**OS**: [macOS/Linux/version]
**Phase**: [Stop / Pull / Dependencies / Migrations / Rebuild / Verify]

## Issue Encountered
[Exact error message or problem description]

## Context
[What step was being executed, what commands were run]

## Resolution
[How it was fixed, or "Unresolved" if still broken]

## Suggested Documentation Update
[What should be added/changed in update_guide.md]
```

3. These reports feed into `/ant-document` for corpus improvements

## Progress Tracking

Use TodoWrite with these phases:
- Services stopped (`make down`)
- Repositories updated (`make pull`)
- Dependencies installed (frontend: pnpm install)
- Migrations applied (`make migrate`)
- Services rebuilt and started (`make up`)
- Services verified healthy (`make ps` + health checks)

## Critical Rules

- Work in `anthropos-dev/` only
- Use `make` commands for all platform operations
- Stop services before pulling code
- Handle git conflicts before continuing
- Verify health after updates
- Follow the guide - don't improvise

## Additional Resources

- For complete update instructions, read `corpus/ops/update_guide.md`
- For technical reference (update scenarios, troubleshooting), see [reference.md](reference.md)
