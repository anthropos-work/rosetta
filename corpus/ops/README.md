# Operations Guides

This directory contains guides for operating the Anthropos platform locally.

## Available Operations

| Guide | Purpose | When to Use |
|-------|---------|-------------|
| [Platform Setup](./platform-setup/setup_guide.md) | Build the development environment | First time on a new machine |
| [Platform Run](./platform-run/run_guide.md) | Start the platform locally | Daily development work |
| [Platform Update](./platform-update/update_guide.md) | Sync code, deps, and schemas | After being away or before new features |
| [Quick Ops](./quick_ops.md) | Common commands reference | When you need a quick command |

## Workflow

```
┌─────────────────────────────────────────────────────────────────┐
│                     First Time Setup                             │
│  /anthropos-setup  or  corpus/ops/platform-setup/setup_guide.md │
│  Install tools, clone repos, configure environment               │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Daily Development                           │
│  corpus/ops/platform-run/run_guide.md                            │
│  Start Docker, backend services, frontend                        │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Periodic Updates                              │
│  corpus/ops/platform-update/update_guide.md                      │
│  Pull latest code, install deps, run migrations                  │
└─────────────────────────────────────────────────────────────────┘
```

## Progress Tracking

Each guide has a companion checklist:
- `platform-setup/setup_checklist_macos.md` / `setup_checklist_linux.md`
- `platform-run/run_checklist.md`
- `platform-update/update_checklist.md`

**Usage**: Copy the checklist to `anthropos-dev/` and track your progress there. The originals are only updated when the guide structure changes.

## Future Operations

This directory is designed to accommodate additional operation types:
- `platform-deploy/` - Deployment procedures
- `platform-debug/` - Debugging and diagnostics
