# Service Documentation — Index

Every service doc in this directory, enumerated. Each follows the
[`TEMPLATE.md`](TEMPLATE.md) pattern (Role & Responsibility · Architecture & Code Map ·
Interface Discovery · Local Development · Testing).

For the *categorised* view (tiers, ports, profiles, which repos are cloned where) see
[`../architecture/service_taxonomy.md`](../architecture/service_taxonomy.md); for how the
services talk to each other see [`../architecture/dependency_map.md`](../architecture/dependency_map.md).

> **⚠️ `app` is the backend monolith.** **Seven** services in this index — skiller, skillpath,
> jobsimulation, cms, storage, messenger and customerio-sync — are **folded into `app`**: each domain runs
> in-process and none has a compose service or a local container. Their docs are kept for domain knowledge
> and carry a merge banner at the top. **"No longer deploy separately" is a claim about PROD, and it is not
> uniformly measurable** — for `customerio-sync` in particular the standalone's terraform lives in a repo
> that has never been in any clone set, so its prod half is asserted from `app`'s side only
> ([`platform-migration-status.md:101`](../architecture/platform-migration-status.md)). The fenced map is
> authoritative per service; this banner is about the LOCAL stack.
>
> **`roadrunner` is the eighth, and it is different: orphaned, not merged-and-undeployed.** Nothing calls it,
> but `roadrunner/terraform/main.tf:19` still reads `service_desired_count = 1` — **last changed at `84a4b4f`
> (2025-12-15), seven months before the fold, and not a decision about it** (`git blame -L 19,19`; M257x
> iter-115) — so it **does** still deploy,
> unlike cms (`cms/terraform/main.tf:39` = 0) and jobsimulation, whose ECS service **M810 has already destroyed** (`6092c6d2`; `service_desired_count` no longer appears in `jobsimulation/terraform/main.tf` at all — `:15-22`). It is the one row where prod and the platform's own
> `repos.yml` contradict each other. See [`platform-migration-status.md`](../architecture/platform-migration-status.md).
>
> And **none of them starts a container any more.** cms, jobsimulation and roadrunner did run locally as
> unfederated husks, but platform **`d11a403`** (2026-08-03) deleted all three from `docker-compose.yml`
> **and** from `repos.yml`; **`838d907`** (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and
> customerio-sync containers"*) then did the same to the last three.
> `docker-compose.yml` declares **5** services (7 effective, with `common.yml`'s `postgresql` +
> `redis`), and `repos.yml` carries **4** entries — `app`, `sentinel`, `next-web-app`, `studio-desk`.
> Read [`backend.md`](backend.md) for the current shape.

## Core backend services (Tier 1 — Go)

| Doc | Service | One-liner |
|---|---|---|
| [`backend.md`](backend.md) | Backend (`app`) | **The monolith.** Main API gateway + user/org management, **plus** the **seven** folded domains this index's banner names — skiller (taxonomy, matching, embeddings), skillpath, jobsimulation, cms, storage, messenger, customerio-sync (`app/internal/{skiller,skillpath,jobsimulation,cms,storage,messenger,customeriosync}/`, all present @ `app` `ad9f3c49`) — and the AI-readiness subsystem, academy store, AI Labs LabSession. **`roadrunner` is NOT one of them**: `app/internal/roadrunner/` exists at no ref — at `ad9f3c49` **no path in the whole tree matches `roadrunner`** — and Judge0 execution was absorbed into the *jobsim* domain as `app/internal/jobsimulation/runner/`, wired at `app/internal/jobsimwiring/wiring.go:123`. This row listed a "roadrunner domain" until M257x iter-102, contradicting the ⚠️ at `:20-23` of this same file and the `app` row of [`platform-migration-status.md`](../architecture/platform-migration-status.md) |
| [`cms.md`](cms.md) | CMS — **merged into `app`** | **The content layer** — owns authored CONTENT/DEFINITIONS (skill paths, simulation blueprints, the library), wrapping Directus as proxy + business logic + cache. Embeds the studio-room generation pipeline. Folded in at cms-in-app v8.0 (app v1.360.0); teardown **M810** |
| [`sentinel.md`](sentinel.md) | Sentinel | **Authorization only** (Casbin RBAC/ABAC). Authentication is Clerk + the `authn` middleware, *not* Sentinel |
| [`jobsimulation.md`](jobsimulation.md) | Jobsimulation — **merged into `app`** | The **runtime/session engine** that *runs* AI simulations (voice, chat, code, documents) and emits completion events. Holds run/session state, never content. Folded in at jobsim-in-app; the prod **ECS service is deleted — M810 landed** (`6092c6d2`), the terraform module surviving only to own the LiveKit/Chime buckets, the SSM parameters and the atlas tracker |
| [`storage.md`](storage.md) | Storage — **merged into `app`** | Centralized file/blob service — private + public S3-backed managers by namespace + UUID. Stateless, owns no DB. Folded in at v9.0 "support-in-app" (2026-08-04); container and `repos.yml` entry deleted at `838d907`. In prod the ECS service is **deleted**, not scaled to zero — the module survives only to keep the buckets/CDN under `prevent_destroy` |
| [`roadrunner.md`](roadrunner.md) | Roadrunner — **orphaned** (not "merged and undeployed") | Code-execution proxy to the Judge0 sandbox. Execution moved in-process with the jobsim engine and `backend` calls Judge0 directly via `JUDGE0_BASE_URL` — but prod terraform still reads `= 1` (`roadrunner/terraform/main.tf:19`, unchanged since `84a4b4f` / 2025-12-15 — it predates the fold and nobody has been back), even though `d11a403` removed its local container **and** its `repos.yml` entry |
| [`gotenberg.md`](gotenberg.md) | Gotenberg | Third-party stateless Office-doc → PDF conversion (LibreOffice headless). One consumer: `app` |
| [`messenger.md`](messenger.md) | Messenger — **merged into `app`** | Centralized transactional email via Brevo + Liquid templates. Folded in at v9.0 "support-in-app"; container, `repos.yml` entry and `messenger` profile all deleted at `838d907`, and `app` gates the domain behind `MESSENGER_ENABLED` (unset = off on a laptop). Other services never called Brevo directly — they **publish Redis Stream events** the domain consumes (`messenger/internal/flow/flow.go:72-104` @ `fa47850`, `AddSubscriber("backend", …)`, 21 live handlers); `app` took over messenger's own consumer group. It *exposes* a `MessengerService` Connect-RPC surface, but **no service ever constructed a client for it**: `MESSENGER_RPC_ADDR` appears in no repo — every clone at its own named ref, nested repos included — and `git -C stack-demo/platform log -S 'MESSENGER_RPC' --oneline 0c91421d` returns **0** commits that ever set it (positive control at the same repo+ref: `-S 'SKILLER_RPC'` returns **7**) |
| [`customerio-sync.md`](customerio-sync.md) | CustomerIO Sync — **merged into `app`** | One-directional background pipeline, Postgres `public` → **Brevo** (the Customer.io name is a fossil), for marketing automation. The last Go service folded into `app`; container deleted at `838d907`, gated by `CUSTOMERIO_SYNC_ENABLED`. Its unique "build straight from a GitHub URL" compose pattern died with it |
| [`db-backup.md`](db-backup.md) | db-backup | A **43-line Bash** script (not Go) dumping Postgres to **two** targets, S3 + Hetzner — **never Azure**. **Its schedule has been commented out since 2025-05-29**, at the commit prod pins. **Production-only** — not in local compose |

## Frontends & gateway

| Doc | Service | One-liner |
|---|---|---|
| [`graphql-wundergraph.md`](graphql-wundergraph.md) | GraphQL Gateway | *was* Apollo Federation v2 via Cosmo Router — **ONE** subgraph (`backend`) since `915da06`. **GONE IN BOTH STATES** (corrected iter-124): deleted from local dev at platform `2adcf71` (2026-07-31) and **destroyed in production** (`infrastructure` `services.tf:509-517`); repo ARCHIVED; the frontends hit `backend` at `:8082/graphql/query` |
| [`next-web-app.md`](next-web-app.md) | Next Web App | The Next.js **16** monorepo on Vercel — Workforce (`apps/web`), Hiring (`apps/hiring`), mobile |
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
| [`skillpath.md`](skillpath.md) | **Merged into `app`** then decommissioned ("skillpath-in-app", platform M502→M507). The runtime session engine now lives in `app`; session state moved to `public.skill_path_sessions`; no skillpath container or subgraph. Skill-path *content* still lives in the cms domain **inside `app`**. Heavily inbound-linked — treat as a redirect |
| [`chronos.md`](chronos.md) | **Decommissioned** — removed from compose + `repos.yml` (platform `045857c`). **The GitHub repo is NOT archived** (last push 2026-04-23) — the corpus called it archived; the org disagrees. Session timeouts are now in-process Asynq |
| [`intelligence.md`](intelligence.md) | **Archived** — removed from compose + `repos.yml` (platform `fdfa189`). Was background sync between the backend and skiller schemas |

## Related

- [`TEMPLATE.md`](TEMPLATE.md) — the pattern every doc here follows
- [`../architecture/service_taxonomy.md`](../architecture/service_taxonomy.md) — tiers, ports, profiles
- [`../architecture/dependency_map.md`](../architecture/dependency_map.md) — who calls whom, and the Redis Streams events
- [`../ops/platform_repo.md`](../ops/platform_repo.md) — the `platform` orchestrator (Make targets, profiles, compose, `repos.yml`)
- [`../tools/README.md`](../tools/README.md) — the tools tier
