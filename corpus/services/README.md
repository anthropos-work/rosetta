# Service Documentation — Index

Every service doc in this directory, enumerated. Each follows the
[`TEMPLATE.md`](TEMPLATE.md) pattern (Role & Responsibility · Architecture & Code Map ·
Interface Discovery · Local Development · Testing).

For the *categorised* view (tiers, ports, profiles, which repos are cloned where) see
[`../architecture/service_taxonomy.md`](../architecture/service_taxonomy.md); for how the
services talk to each other see [`../architecture/dependency_map.md`](../architecture/dependency_map.md).

> **⚠️ `app` is the backend monolith.** **Eight** services in this index — skiller, skillpath,
> roadrunner, jobsimulation, cms, and (at v9.0 "support-in-app", 2026-08-04) messenger, storage and
> customerio-sync — are **folded into `app`** and no longer deploy separately. Their docs are kept
> for domain knowledge and carry a merge banner at the top. Read [`backend.md`](backend.md) for the
> current shape.
>
> **[`sentinel.md`](sentinel.md) is the only Go support service still running out-of-process.**

## Core backend services (Tier 1 — Go)

| Doc | Service | One-liner |
|---|---|---|
| [`backend.md`](backend.md) | Backend (`app`) | **The monolith.** Main API gateway + user/org management, **plus** the folded skiller (taxonomy, matching, embeddings), skillpath, jobsimulation, cms, roadrunner, messenger, storage and customerio-sync domains — and the AI-readiness subsystem, academy store, AI Labs LabSession |
| [`cms.md`](cms.md) | CMS — **merged into `app`** | **The content layer** — owns authored CONTENT/DEFINITIONS (skill paths, simulation blueprints, the library), wrapping Directus as proxy + business logic + cache. Embeds the studio-room generation pipeline. Folded in at cms-in-app v8.0 (app v1.360.0); teardown **M810** |
| [`sentinel.md`](sentinel.md) | Sentinel | **Authorization only** (Casbin RBAC/ABAC). Authentication is Clerk + the `authn` middleware, *not* Sentinel. **The only support service still deployed alongside `backend`** |
| [`jobsimulation.md`](jobsimulation.md) | Jobsimulation — **merged into `app`** | The **runtime/session engine** that *runs* AI simulations (voice, chat, code, documents) and emits completion events. Holds run/session state, never content. Folded in at jobsim-in-app; teardown **M810** |
| [`storage.md`](storage.md) | Storage — **merged into `app`** | The private + public S3-backed object managers. `backend` reads/writes both buckets directly since v9.0; `STORAGE_RPC_ADDR` is gone. The ECS service is gone but `module.storage-service_euwest1` **must stay** — it owns the buckets, CloudFront and `media.anthropos.work` |
| [`roadrunner.md`](roadrunner.md) | Roadrunner — **merged into `app`** | Code-execution proxy to the Judge0 sandbox. Execution moved in-process with the jobsim engine; `backend` calls Judge0 directly via `JUDGE0_BASE_URL` |
| [`gotenberg.md`](gotenberg.md) | Gotenberg | Third-party stateless Office-doc → PDF conversion (LibreOffice headless). One consumer: `app` |
| [`messenger.md`](messenger.md) | Messenger — **merged into `app`** | Transactional email via Brevo + Liquid templates and the 24 event handlers. In `backend` since v9.0, gated by `MESSENGER_ENABLED`; `app` **takes over messenger's own Redis consumer group** rather than merging handlers. ECS module deleted; still startable from the `messenger` profile as the rollback path |
| [`customerio-sync.md`](customerio-sync.md) | CustomerIO Sync — **merged into `app`** | One-directional background pipeline, Postgres `public` → **Brevo** marketing contacts (the "Customer.io" name is a fossil). In `backend` on the asynq scheduler since v9.0, gated by `CUSTOMERIO_SYNC_ENABLED`. Terraform module deleted — **no rollback path** |
| [`db-backup.md`](db-backup.md) | db-backup | Scheduled Postgres backups every 6 h to three geographies (S3, Azure, Hetzner). **Production-only** — not in local compose |

## Frontends & gateway

| Doc | Service | One-liner |
|---|---|---|
| [`graphql-wundergraph.md`](graphql-wundergraph.md) | GraphQL Gateway | Apollo Federation v2 via Cosmo Router — **one** subgraph (`backend`) since cms-in-app |
| [`next-web-app.md`](next-web-app.md) | Next Web App | The Next.js 15 monorepo on Vercel — Workforce (`apps/web`), Hiring (`apps/hiring`), mobile |
| [`studio-desk.md`](studio-desk.md) | Studio-Desk | TypeScript/Vite/Express design tool for authoring simulation blueprints |
| [`studio-room.md`](studio-room.md) | Studio-Room | Python/asyncio AI content-generation pipeline. **Embedded inside the `app` (backend) container** since cms-in-app |
| [`ant-academy.md`](ant-academy.md) | Ant Academy | Internal Next.js 16 + Expo learning portal for `@anthropos.work` staff. Vercel-deployed, native-only, DB-authoritative catalog |

## Cross-cutting subsystems & domains

| Doc | Subject | One-liner |
|---|---|---|
| [`ai-readiness.md`](ai-readiness.md) | AI Readiness | Org-level AI-capability diagnostics inside `app` (the `internal/aireadiness/` package) — the cycle/funnel model, the gate-by-surface rules, and the demo seeder contract |
| [`coursebuilder.md`](coursebuilder.md) | Course Builder | `app` domain — the in-process author→benchmark→refine AI pipeline that generates Academy chapters/skill-paths (Bedrock Opus author + Sonnet grader, ≥90 gate). HTTP+SSE, no subgraph |
| [`ai-labs.md`](ai-labs.md) | AI Labs + Credits | `app` domains — the hosted AI coding-lab product (catalog + sandbox sessions via the `labs-api` control plane) **and** the credit ledger ("shared purse", live for Course Builder) + Stripe payments/subscriptions |
| [`askengine.md`](askengine.md) | Ask Engine / Talk-to-Data | `app` domain — the NL analytics copilot: an agentic Bedrock LLM writes SQL, runs it in an org-scoped read-only sandbox, explains results. HTTP+SSE, no subgraph |
| [`academy-backend.md`](academy-backend.md) | Academy Backend | `app` domain — the server-authoritative owner of the Academy catalog + per-user study state, served to the [ant-academy frontend](ant-academy.md) over the `app` subgraph (distinct from the frontend doc) |
| [`hiring.md`](hiring.md) | Hiring | The recruiting **org-type** (`is_hiring`) + the candidate-comparison read-model. Authored from a live render-probe, not inferred |
| [`clerk-integration.md`](clerk-integration.md) | Clerk | The cross-cutting single source of truth for how the platform uses Clerk (vs. per-service mentions elsewhere) |
| [`clerkenstein.md`](clerkenstein.md) | Clerkenstein | The **Clerk mock** that makes demo stacks Clerk-free — a `rosetta-extensions` section, consumed per-stack at a pinned tag |

## Archived / merged — kept as redirects

These describe services that no longer run. They stay because many docs still link to them.

| Doc | Fate |
|---|---|
| [`skiller.md`](skiller.md) | **Merged into `app`** (July 2026). The skills domain now lives in `app`'s `public` schema; no skiller container or subgraph. Heavily inbound-linked — treat as a redirect, do not delete |
| [`skillpath.md`](skillpath.md) | **Merged into `app`** then decommissioned ("skillpath-in-app", platform M502→M507). The runtime session engine now lives in `app`; session state moved to `public.skill_path_sessions`; no skillpath container or subgraph. Skill-path *content* still lives in CMS. Heavily inbound-linked — treat as a redirect |
| [`messenger.md`](messenger.md), [`storage.md`](storage.md), [`customerio-sync.md`](customerio-sync.md) | **Merged into `app`** at v9.0 "support-in-app" (2026-08-04). Listed above under Tier 1 because their docs still carry the ported domain detail, but none of the three deploys any more. See [`backend.md`](backend.md) § *The v9.0 "support-in-app" fold* |
| [`chronos.md`](chronos.md) | **Archived** — removed from compose + `repos.yml` (platform `045857c`). Session timeouts are now in-process Asynq |
| [`intelligence.md`](intelligence.md) | **Archived** — removed from compose + `repos.yml` (platform `fdfa189`). Was background sync between the backend and skiller schemas |

## Related

- [`TEMPLATE.md`](TEMPLATE.md) — the pattern every doc here follows
- [`../architecture/service_taxonomy.md`](../architecture/service_taxonomy.md) — tiers, ports, profiles
- [`../architecture/dependency_map.md`](../architecture/dependency_map.md) — who calls whom, and the Redis Streams events
- [`../ops/platform_repo.md`](../ops/platform_repo.md) — the `platform` orchestrator (Make targets, profiles, compose, `repos.yml`)
- [`../tools/README.md`](../tools/README.md) — the tools tier
