# Service Documentation — Index

Every service doc in this directory, enumerated. Each follows the
[`TEMPLATE.md`](TEMPLATE.md) pattern (Role & Responsibility · Architecture & Code Map ·
Interface Discovery · Local Development · Testing).

For the *categorised* view (tiers, ports, profiles, which repos are cloned where) see
[`../architecture/service_taxonomy.md`](../architecture/service_taxonomy.md); for how the
services talk to each other see [`../architecture/dependency_map.md`](../architecture/dependency_map.md).

## Core backend services (Tier 1 — Go)

| Doc | Service | One-liner |
|---|---|---|
| [`backend.md`](backend.md) | Backend (`app`) | The main API gateway + user/org management. **Also owns the skills domain** (taxonomy, assessment, matching, embeddings) since the July 2026 skiller merge, the **AI-readiness** workforce subsystem, the academy store, and AI Labs LabSession |
| [`cms.md`](cms.md) | CMS | **The content layer** — owns authored CONTENT/DEFINITIONS (skill paths, simulation blueprints, the library), wrapping Directus as proxy + business logic + cache. Embeds the studio-room generation pipeline |
| [`sentinel.md`](sentinel.md) | Sentinel | **Authorization only** (Casbin RBAC/ABAC). Authentication is Clerk + the `authn` middleware, *not* Sentinel |
| [`jobsimulation.md`](jobsimulation.md) | Jobsimulation | The **runtime/session engine** that *runs* AI simulations (voice, chat, code, documents) and emits completion events. Holds run/session state, never content |
| [`skillpath.md`](skillpath.md) | Skillpath | The **runtime/session engine** tracking per-user progression *state*. The skill-path content lives in CMS |
| [`storage.md`](storage.md) | Storage | Centralized file/blob service — private + public S3-backed managers by namespace + UUID. Stateless, owns no DB |
| [`roadrunner.md`](roadrunner.md) | Roadrunner | Code-execution proxy to the Judge0 sandbox. **⚠️ ORPHANED** — execution moved in-process to `jobsimulation/internal/runner/`; nothing calls it |
| [`gotenberg.md`](gotenberg.md) | Gotenberg | Third-party stateless Office-doc → PDF conversion (LibreOffice headless). One consumer: `app` |
| [`messenger.md`](messenger.md) | Messenger | Centralized transactional email via Brevo + Liquid templates. Opt-in `messenger` profile; other services fire an RPC rather than calling Brevo |
| [`customerio-sync.md`](customerio-sync.md) | CustomerIO Sync | One-directional background pipeline, Postgres `public` → Customer.io, for marketing automation. Opt-in profile; built from a GitHub URL, not cloned |
| [`db-backup.md`](db-backup.md) | db-backup | Scheduled Postgres backups every 6 h to three geographies (S3, Azure, Hetzner). **Production-only** — not in local compose |

## Frontends & gateway

| Doc | Service | One-liner |
|---|---|---|
| [`graphql-wundergraph.md`](graphql-wundergraph.md) | GraphQL Gateway | Apollo Federation v2 via Cosmo Router — 4 subgraphs (app, jobsimulation, cms, skillpath) |
| [`next-web-app.md`](next-web-app.md) | Next Web App | The Next.js 15 monorepo on Vercel — Workforce (`apps/web`), Hiring (`apps/hiring`), mobile |
| [`studio-desk.md`](studio-desk.md) | Studio-Desk | TypeScript/Vite/Express design tool for authoring simulation blueprints |
| [`studio-room.md`](studio-room.md) | Studio-Room | Python/asyncio AI content-generation pipeline. **Embedded inside the cms container** as `cms/studio/` |
| [`ant-academy.md`](ant-academy.md) | Ant Academy | Internal Next.js 16 + Expo learning portal for `@anthropos.work` staff. Vercel-deployed, native-only, DB-authoritative catalog |

## Cross-cutting subsystems & domains

| Doc | Subject | One-liner |
|---|---|---|
| [`ai-readiness.md`](ai-readiness.md) | AI Readiness | Org-level AI-capability diagnostics inside `app` — the cycle/funnel model, the gate-by-surface rules, and the demo seeder contract |
| [`hiring.md`](hiring.md) | Hiring | The recruiting **org-type** (`is_hiring`) + the candidate-comparison read-model. Authored from a live render-probe, not inferred |
| [`clerk-integration.md`](clerk-integration.md) | Clerk | The cross-cutting single source of truth for how the platform uses Clerk (vs. per-service mentions elsewhere) |
| [`clerkenstein.md`](clerkenstein.md) | Clerkenstein | The **Clerk mock** that makes demo stacks Clerk-free — a `rosetta-extensions` section, consumed per-stack at a pinned tag |

## Archived / merged — kept as redirects

These describe services that no longer run. They stay because many docs still link to them.

| Doc | Fate |
|---|---|
| [`skiller.md`](skiller.md) | **Merged into `app`** (July 2026). The skills domain now lives in `app`'s `public` schema; no skiller container or subgraph. Heavily inbound-linked — treat as a redirect, do not delete |
| [`chronos.md`](chronos.md) | **Archived** — removed from compose + `repos.yml` (platform `045857c`). Session timeouts are now in-process Asynq |
| [`intelligence.md`](intelligence.md) | **Archived** — removed from compose + `repos.yml` (platform `fdfa189`). Was background sync between the backend and skiller schemas |

## Related

- [`TEMPLATE.md`](TEMPLATE.md) — the pattern every doc here follows
- [`../architecture/service_taxonomy.md`](../architecture/service_taxonomy.md) — tiers, ports, profiles
- [`../architecture/dependency_map.md`](../architecture/dependency_map.md) — who calls whom, and the Redis Streams events
- [`../ops/platform_repo.md`](../ops/platform_repo.md) — the `platform` orchestrator (Make targets, profiles, compose, `repos.yml`)
- [`../tools/README.md`](../tools/README.md) — the tools tier
