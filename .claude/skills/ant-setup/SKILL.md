---
name: ant-setup
description: Build Anthropos dev environment following setup guide with auto-improvement
argument-hint: [step-name or 'full']
---

# Anthropos Development Environment Setup

Build the Anthropos development environment by following `corpus/ops/setup_guide.md` with systematic verification.

## Your Mission

1. **Read the guide**: `corpus/ops/setup_guide.md` is your source of truth
2. **Apply STEP RUN principles**: Verify before/after, request confirmation
3. **Track progress**: Use TodoWrite for each phase (no separate checklist file)
4. **Report issues**: Create ops reports for errors and fixes discovered

## STEP RUN Principles

Apply to EVERY step in the guide:

| Principle | Action |
|-----------|--------|
| Verify Before Install | Check if tool exists before attempting installation |
| Request Confirmation | Ask user before installing tools or modifying system |
| Verify After Install | Confirm installation succeeded |
| Report Issues | Create ops report when errors or improvements found |

## Confirmation Policy

**ALWAYS ask for confirmation before**:
- Installing system tools (brew, apt, etc.)
- Cloning repositories
- Starting services
- Creating/modifying .env files

## Error Handling

1. Do NOT skip errors - resolve first
2. Document error message verbatim
3. Research and test solution
4. Create ops report (see below)
5. Continue

## Ops Reports (Instead of Auto-Improving Docs)

When you discover errors, missing steps, or better approaches:

1. Create a report file: `anthropos-dev/ops-reports/op_YYYYMMDD_HHMMSS_setup_<topic>.md`
2. Use this template:

```markdown
# Ops Report: [Brief Title]

**Date**: YYYY-MM-DD HH:MM
**Skill**: /ant-setup
**OS**: [macOS/Linux/version]
**Phase**: [Which setup phase]

## Issue Encountered
[Exact error message or problem description]

## Context
[What step was being executed, what commands were run]

## Resolution
[How it was fixed, or "Unresolved" if still broken]

## Suggested Documentation Update
[What should be added/changed in setup_guide.md]
```

3. These reports feed into `/ant-document` for corpus improvements

## Progress Tracking

Use TodoWrite with phases from the guide:
- Prerequisites verified (Git, Docker, Go, Node, pnpm)
- GitHub SSH access configured
- Workspace created (anthropos-dev/)
- Platform repo cloned
- All repos cloned via `make init`
- Environment file configured (platform/.env)
- Services started via `make up`
- Database migrations applied via `make migrate`
- Frontend running (native or containerized)
- Studio-Desk running (native or containerized)

## Critical Rules

- Work in `anthropos-dev/` scratchpad only
- Never commit .env files
- Use `make` commands for all platform operations
- Ask before every install/modification
- Follow the guide - don't improvise

## Success Criteria

Setup complete when:
1. All tools installed and verified
2. All repositories cloned (via `make init`)
3. Environment file configured (platform/.env with GH_PAT, Clerk, AI keys)
4. Docker services running and healthy (`make ps`)
5. Frontend accessible at localhost:3000
6. Studio-Desk accessible at localhost:9100

## Additional Resources

- For complete setup instructions, read `corpus/ops/setup_guide.md`
- For technical reference, see [reference.md](reference.md)
