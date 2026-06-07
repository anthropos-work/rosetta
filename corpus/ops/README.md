# Operations Guides

This directory contains guides for operating the Anthropos platform locally.

> **Corpus vs. extensions boundary:** rosetta is a read-only doc corpus + dev-env skills; ALL executable tooling that operates a spawned stack lives in rosetta-extensions — authored in `.agentspace/rosetta-extensions/`, tagged, and consumed per-stack via a pinned-tag clone.

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
| [Demo Stacks](./rosetta_demo.md) | **Disposable, isolated demo stacks (`demo-N`) alongside the dev stack — Clerkenstein-wired, offset ports, killable cleanly, zero platform-repo change. Skills `/demo-up`, `/demo-down`; list via `/stack-list`.** (v1.1/M3) | Spinning up a throwaway demo environment to seed (M4) + show |
| [Seeding Spec](./seeding-spec.md) | Declaratively backfill a stack with structural data (blueprint + DAG + the 3-layer write isolation guard). Skill `/stack-seed` (`dev-N` or `demo-N`). (v1.1/M7) | Populating a demo/dev stack with an org + users + activity |
| [DB Access](./db-access.md) | Read-only prod DB access (the wired `postgres` MCP tool **or** Tailscale + `~/.pgpass`) + the public-vs-customer boundary. Skill `/db-query`. (v1.2/M9a) | Investigating data, sizing a surface, or telling public reference data from customer data |
| [Snapshot Spec](./snapshot-spec.md) | Capture a **public** reference surface once from a safe prod source, manifest-cache it in `.agentspace`, replay per-stack — tenant-data firewall + snapshot-fidelity. `stacksnap` CLI. (v1.2/M9a) | Filling a stack with the real public taxonomy/content library |
| [Demo Recipes](./demo/README.md) | The end-to-end demo-env recipe family (up → **snapshot** → seed → use → down) + presets; the `/stack-snapshot` skill *set-dresses* a stack (`dev-N` or `demo-N`) with the real public taxonomy + Directus content (100% catalog). (v1.2/M11) | Running a believable, full-fidelity demo world |
| [Safety & Security](./safety.md) | **The authoritative, code-cited safety contract of the stack tooling: the read-side (never reads private/customer data — the `AssertPlan`/`AssertCaptured` firewall + public predicates + bounded read-only capture) and the write-side (never touches prod — the 3-layer `CheckWrite`/`PreflightEnv`/`AssertClean` isolation guard + never-write shared Directus/prod-S3 + doubled n=0 guards + audit-proven zero pollution).** (v1.3/M15) | Understanding *why* snapshot/seed/db-query can't read customer data or pollute production — the safety landing page |

## Workflow

```
┌─────────────────────────────────────────────────────────────────┐
│                  First Time Setup + Daily Run                    │
│  /dev-up  or  corpus/ops/setup_guide.md + run_guide.md           │
│  Install tools, clone repos, configure env, start the stack      │
│  (one skill — consolidates the former setup + start)             │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Periodic Updates                              │
│  /stack-update  or  corpus/ops/update_guide.md                   │
│  Pull latest code, install deps, run migrations                  │
└─────────────────────────────────────────────────────────────────┘
```

## Progress Tracking

All operations use Claude's **TodoWrite** tool for real-time progress tracking.

## Ops Reports

When `/dev-up` or `/stack-update` encounter errors or discover improvements, they create **ops reports** in `stack-dev/ops-reports/`:

```
stack-dev/ops-reports/
├── op_20250127_143022_setup_pgvector.md
├── op_20250127_151045_run_port_conflict.md
└── op_20250128_092311_update_migration_fail.md
```

### Report Format

```markdown
# Ops Report: [Brief Title]

**Date**: YYYY-MM-DD HH:MM
**Skill**: /dev-up | /stack-update
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
