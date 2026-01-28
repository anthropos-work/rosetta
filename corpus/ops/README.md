# Operations Guides

This directory contains guides for operating the Anthropos platform locally.

## Available Operations

| Guide | Purpose | When to Use |
|-------|---------|-------------|
| [Platform Setup](./setup_guide.md) | Build the development environment | First time on a new machine |
| [Platform Run](./run_guide.md) | Start the platform locally | Daily development work |
| [Webhook Setup](./webhook_setup.md) | Configure Clerk webhooks for user sync | When you need user/org data locally |
| [Platform Update](./update_guide.md) | Sync code, deps, and schemas | After being away or before new features |
| [Quick Ops](./quick_ops.md) | Common commands reference | When you need a quick command |

## Workflow

```
┌─────────────────────────────────────────────────────────────────┐
│                     First Time Setup                             │
│  /ant-setup  or  corpus/ops/setup_guide.md                       │
│  Install tools, clone repos, configure environment               │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Daily Development                           │
│  /ant-run  or  corpus/ops/run_guide.md                           │
│  Start Docker, backend services, frontend                        │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Periodic Updates                              │
│  /ant-update  or  corpus/ops/update_guide.md                     │
│  Pull latest code, install deps, run migrations                  │
└─────────────────────────────────────────────────────────────────┘
```

## Progress Tracking

All operations use Claude's **TodoWrite** tool for real-time progress tracking.

## Ops Reports

When `/ant-setup`, `/ant-run`, or `/ant-update` encounter errors or discover improvements, they create **ops reports** in `anthropos-dev/ops-reports/`:

```
anthropos-dev/ops-reports/
├── op_20250127_143022_setup_pgvector.md
├── op_20250127_151045_run_port_conflict.md
└── op_20250128_092311_update_migration_fail.md
```

### Report Format

```markdown
# Ops Report: [Brief Title]

**Date**: YYYY-MM-DD HH:MM
**Skill**: /ant-setup | /ant-run | /ant-update
**OS**: [macOS 14.x / Ubuntu 22.04 / etc.]
**Phase**: [Which operation phase]

## Issue Encountered
[Exact error message]

## Context
[What was being done]

## Resolution
[How fixed, or "Unresolved"]

## Suggested Documentation Update
[What to add/change in the guides]
```

### Integration Workflow

Ops reports are **not** automatically applied to documentation. Instead:

1. Skills create reports during execution
2. Run `/ant-integrate` to review and apply improvements
3. Human reviews and commits changes

This separates "live execution" from "corpus maintenance".

## Future Operations

This directory may grow to include:
- `deploy_guide.md` - Deployment procedures
- `debug_guide.md` - Debugging and diagnostics
