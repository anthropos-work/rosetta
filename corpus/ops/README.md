# Operations Guides

This directory contains guides for operating the Anthropos platform locally.

## Available Operations

| Guide | Purpose | When to Use |
|-------|---------|-------------|
| **[Staging Bringup](./staging-bringup.md)** | **Full personal-staging onboarding (fresh VM → working Tailscale staging + live prod data + dev Clerk + daily sync). Includes the [colony v2-JWT vendoring recipe](./staging-bringup.md#bringup-quirks-consolidated-as-a-procedural-narrative) (Quirk #11), the [Atlas migrations gap fix](./staging-bringup.md#45-apply-pending-atlas-migrations) (§4.5), and the [known-schema-drifts table](./staging-bringup.md#105-known-schema-drifts-expected-on-staging) (§10.5).** | **New engineer (or AI agent) joining the team — start here** |
| [Staging Sync](./staging-sync.md) | Daily force-reset to `origin/main` + skip-worktree mechanics + recovery. **Note:** Atlas migrations are NOT in the daily run — see [Atlas migrations are NOT run by sync](./staging-sync.md#atlas-migrations-are-not-run-by-sync). | Understanding what the daily 06:00 UTC routine does, recovering clobbered WIP, or remembering to run Atlas periodically |
| [Staging Clerk](./staging-clerk.md) | Dev Clerk app `national-elk-17`, shared cross-engineer test login, the load-bearing `clerk-fetch-fix.js` monkey-patch, and the [v2 session-token anatomy](./staging-clerk.md#anatomy-of-a-v2-session-token) | Setting up auth on a new staging or debugging Clerk symptoms |
| [Platform Setup](./setup_guide.md) | Build the development environment | First time on a new machine (no prod-dump path) |
| [Personal Staging from a Prod Dump](./staging_from_dump.md) | Restore a prod DB dump, rebind to a dev Clerk app, kill outbound email, apply colony/Clerk patches | Engineer-rebind reference (called from `staging-bringup.md`) |
| [Platform Run](./run_guide.md) | Start the platform locally | Daily development work |
| [Webhook Setup](./webhook_setup.md) | Configure Clerk webhooks for user sync | When you need user/org data locally |
| [Platform Update](./update_guide.md) | Sync code, deps, and schemas | After being away or before new features (superseded by `staging-sync.md` on staging hosts) |
| [Quick Ops](./quick_ops.md) | Common commands reference | When you need a quick command |
| [Demo Stacks](./demo_stacks.md) | **Disposable, isolated demo stacks (`demo-N`) alongside the dev stack — Clerkenstein-wired, offset ports, killable cleanly, zero platform-repo change. Skills `/demo-up`, `/demo-down`, `/demo-status`.** (v1.1/M3) | Spinning up a throwaway demo environment to seed (M4) + show |

## Workflow

```
┌─────────────────────────────────────────────────────────────────┐
│                     First Time Setup                             │
│  /setup-platform  or  corpus/ops/setup_guide.md                  │
│  Install tools, clone repos, configure environment               │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Daily Development                           │
│  /start-platform  or  corpus/ops/run_guide.md                    │
│  Start Docker, backend services, frontend                        │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Periodic Updates                              │
│  /update-platform  or  corpus/ops/update_guide.md                │
│  Pull latest code, install deps, run migrations                  │
└─────────────────────────────────────────────────────────────────┘
```

## Progress Tracking

All operations use Claude's **TodoWrite** tool for real-time progress tracking.

## Ops Reports

When `/setup-platform`, `/start-platform`, or `/update-platform` encounter errors or discover improvements, they create **ops reports** in `anthropos-dev/ops-reports/`:

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
**Skill**: /setup-platform | /start-platform | /update-platform
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
2. Run `/update-knowledge` to review and apply improvements
3. Human reviews and commits changes

This separates "live execution" from "corpus maintenance".

## Future Operations

This directory may grow to include:
- `deploy_guide.md` - Deployment procedures
- `debug_guide.md` - Debugging and diagnostics
