# Operations Guides

This directory contains guides for operating the Anthropos platform locally.

## Available Operations

| Guide | Purpose | When to Use |
|-------|---------|-------------|
| [Setup Guide](./setup/setup_guide.md) | Build the development environment | First time on a new machine |
| [Run Guide](./run/run_guide.md) | Start the platform locally | Daily development work |

## Workflow

```
┌─────────────────────────────────────────────────────────────────┐
│                        First Time Setup                          │
│  /anthropos-setup  or  corpus/ops/setup/setup_guide.md          │
│  Install tools, clone repos, configure environment               │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Daily Development                           │
│  corpus/ops/run/run_guide.md                                     │
│  Start Docker, backend services, frontend                        │
└─────────────────────────────────────────────────────────────────┘
```

## Progress Tracking

Each guide has a companion checklist:
- `setup/setup_checklist_macos.md` / `setup_checklist_linux.md`
- `run/run_checklist.md`

**Usage**: Copy the checklist to `anthropos-dev/` and track your progress there. The originals are only updated when the guide structure changes.

## Future Operations

This directory is designed to accommodate additional operation types:
- `deploy/` - Deployment procedures
- `migrate/` - Database migration guides
- `debug/` - Debugging and diagnostics
