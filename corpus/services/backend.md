# Backend Service (`app`)

> ## `app` is the backend monolith
>
> **Five former microservices now run inside `app`**, in merge order:
>
> | Merged service | Program | What moved in |
> |---|---|---|
> | [skiller](./skiller.md) | skiller-in-app (v2.1 "quick change", July 2026) | the skills-taxonomy graph (**≥42,790 skills / ≥22,470 job roles** — public subset; [not "60K/18K"](../architecture/shared_libraries.md#taxonomy-figures)), embeddings, AI matching |
> | [skillpath](./skillpath.md) | skillpath-in-app (M502→M507) | skill-path progression engine, session state |
> | [roadrunner](./roadrunner.md) | with jobsim-in-app | Judge0 code execution (called directly via `JUDGE0_BASE_URL`) |
> | [jobsimulation](./jobsimulation.md) | jobsim-in-app (teardown **M810**) | the simulation session engine — `internal/jobsimulation/`, wired by `internal/jobsimwiring/wiring.go` |
> | [cms](./cms.md) | cms-in-app v8.0, app **v1.360.0** (teardown **M810**) | content layer + Directus edge + Studio — `internal/cms/` |
>
> Consequences that hold platform-wide:
> * **The federation composes ONE subgraph** (`backend`). cms-in-app was the **3 → 1** step: the single
>   commit `graphql-wundergraph@915da06` (2026-07-29) deleted **both** `schemas/cms.graphqls` **and**
>   `schemas/jobsimulation.graphqls`, taking the supergraph from (backend, jobsimulation, cms) to
>   (backend) alone. The jobsimulation subgraph therefore **survived jobsim-in-app** and was removed
>   here, not at its own merge.
> * **All of their tables live in `public`**, with the same table names. The `skiller`, `skillpath`,
>   `jobsimulation` and `cms` DB schemas are legacy and non-authoritative.
> * **All of their Connect-RPC surfaces are served on `app`'s single RPC mux.** `messenger` is the only
>   remaining external caller.
> * **`app` owns the `skiller`, `skillpath`, `jobsimulation`, `cms` and `ai_usage` Redis Streams** — both
>   producer and consumer are in-process. Merge new handlers onto the existing subscriber with
>   `.AddHandler(...)`; a second `AddSubscriber` for the same stream silently overwrites the first.
> * **`module.jobsimulation_euwest1` and `module.cms_euwest1` are still declared in production terraform**
>   as the rollback path and take no traffic. Teardown is **M810**.
>
> The skiller-specific detail below is the authoritative
> [**§ Skiller-in-app merge — fact-sheet**](#skiller-in-app-merge--fact-sheet-v21-quick-change).

## Role & Responsibility

`app` is the **main API gateway** of the platform — the service that frontends, hiring apps, and other backend services talk to first. It owns the `public` schema (users, organizations, memberships, assignments, subscriptions, payments) and, since the **skiller-in-app merge (July 2026)**, the **skills taxonomy domain** — the skills graph (**≥42,790 skills** across **≥22,470 job roles**; that is the measured *public* subset, `organization_id IS NULL`, 2026-06-29 — the long-quoted "60K skills / 18K roles" is not a measurement, and [18K is outright refuted](../architecture/shared_libraries.md#taxonomy-figures)), skill/job-role embeddings, and AI skill matching formerly owned by the standalone [skiller](./skiller.md) service. It exposes:

* **GraphQL Federation v2 subgraph** for high-level user / organization / assignment queries — plus the taxonomy types/queries absorbed from the former skiller subgraph (`graph/schemas/skiller_taxonomy.graphqls`)
* **Connect-RPC** for inter-service calls (the only remaining external caller is **messenger**) — the mux registers five handlers unconditionally (`main.go:1185-1228` @ `app` `b948604` v1.366.0): `UsersService` (`:1187`), `OrganizationsService` (`:1188`), `SkillerService` (`:1196`), `JobSimulationService` (`:1204`) and `lab.v1.LabSessionService` (`:1228`), plus **`CMSService` only when the Directus edge is configured** (`if cmsRPCServer != nil`, `:1212-1214`).

  **There is no `SkillPathSessionService`** — measured: **0** occurrences in Go source, and no `skillpath…v1connect` package is imported. Skill-path session state lives in `public.skill_path_sessions` and is reached through the GraphQL subgraph and in-process calls, not over RPC.

  > **⚠️ `app`'s OWN docs still list it** (`app/CLAUDE.md:72`, `app/knowledge/architecture.md:28`), which is where this corpus previously got the claim. That is Trap C in [`../ops/platform-alignment.md`](../ops/platform-alignment.md) — *the platform's planning docs lag its own code*. **Grade against `main.go`, not against `app/CLAUDE.md`.**
* **HTTP** endpoints on port 8082 (local; 8080 in production) for webhooks and miscellaneous integrations — including `POST /api/webhook/directus`, which **fails closed** without `DIRECTUS_WEBHOOK_SECRET`

It also hosts a growing number of cross-cutting features that don't fit neatly into any other service:

* **Talk to Data** (`internal/askengine`) — SSE-streaming natural-language Q&A over the platform's data, powered by Bedrock (Anthropic) with a SQL-validation sandbox. Added 2026-Q2 (v1.266+).
* **Workforce analytics** (`internal/workforce`) — aggregations of skills, simulations, and growth across org members
* **Job-simulation feedback** (`internal/jobsimfeedback`) — post-session signals routed back to the skills domain (in-process since the skiller merge)
* **AI usage / cost tracking** (`internal/aiusage`) — central ledger driven by the `AI` Redis Stream
* **Bootstrap & admin** (`internal/admin`, `internal/bootstrap`, `cmd/bootstrap-org`) — provisioning utilities
* **AI Labs LabSession** (`internal/labs/session`; siblings `internal/labs/labsapi`, `internal/labs/adapter`, `internal/labs/catalog`) — Connect-RPC `lab.v1.LabSessionService` (Create/Get/List/Cancel/ReportEvent) plus a `lab_sessions` Ent table. The labs-api client is wired **only when `LABS_API_URL` is set** (`main.go:743-746` @ `app` `b948604` v1.366.0); with it unset — the usual local/demo case — Create persists a session row without booting a VM and Cancel marks the row cancelled without calling labs-api (see Recent Feature Additions). It is NOT unconditionally nil.
* **Document → PDF conversion** (`internal/converter/gotenberg.go`) — via the Gotenberg service

## Skiller-in-app merge — fact-sheet (v2.1 "quick change")

The standalone `skiller` microservice was **merged into `app`** (July 2026). This is the authoritative,
verified statement of the merged shape — the contract the v2.1 re-ground grades against. Verified
2026-07-08 against the re-synced stack-dev clone (`app@c3c45e01` v1.334.1, `platform@0808b92`), a live
containerized bring-up + migrate, and read-only prod.

- **Domain → the `public` schema, table names unchanged (`skiller.X → public.X`).** The moved tables:
  `skills`, `job_roles`, `categories`, `specializations`, `skill_embeddings`, `job_role_embeddings`,
  `skill_translations`, `job_role_translations`, `job_role_skills`, `job_role_categories` (Ent models now
  in `app/internal/data/ent/schema/`; port migrations in `terraform/migrations/`, merge commit
  `1fc00c78 Deprecate skiller schema`). The legacy `skiller` DB schema still exists on prod as a
  **deprecated mirror** — `public.*` is authoritative.
- **Public predicate `organization_id IS NULL`** (the public taxonomy; customer-private rows carry a real
  `organization_id`). Measured on prod 2026-07-08: **`public.skills WHERE organization_id IS NULL` =
  42,790** (43,584 total incl. 794 org-private), `public.job_roles` (org NULL) = 22,490, `categories` = 23,
  `specializations` = 1,447, `public.skill_embeddings` = 43,584. (The ~42,763 figure quoted in the roadmap
  is this count; taxonomy grows over time.) The independent 2026-06-29 public-only snapshot capture agrees
  within that drift — 42,790 skills / 22,470 job roles / 18,919 job-role embeddings; see
  [the canonical "60K / 18K" statement](../architecture/shared_libraries.md#taxonomy-figures). **These are
  floors, not totals** — a public-only measurement cannot see org-private rows.
- **RPC re-pointed** — the `SkillerService` Connect-RPC surface is served **by app itself**
  (`internal/rpc/skillerrpc/`). Consumers keep the env var, re-pointed: `SKILLER_RPC_ADDR=http://backend:8083`
  locally (all four occurrences in the merged `docker-compose.yml`), `http://backend:8081` in prod terraform.
- **Federation is now 1 subgraph**: **backend**. The skiller subgraph was removed at the skiller merge
  (`schemas/skiller.graphqls` deleted at `graphql-wundergraph@749dc86`, "remove skiller subgraph and update
  related configurations", 2026-06-24), which left **4** — backend, jobsimulation, cms, skillpath. The
  **skillpath** subgraph went next when the skillpath service merged into `app` ("skillpath-in-app", platform
  M502→M507; `schemas/skillpath.graphqls` deleted at `graphql-wundergraph@7c17e63`, 2026-07-21) → **3**. The
  last step is `graphql-wundergraph@915da06` (2026-07-29), which deleted **both** `schemas/cms.graphqls` and
  `schemas/jobsimulation.graphqls` in one commit — **3 → 1**. (So the jobsimulation subgraph outlived
  jobsim-in-app; its removal was staged and landed with cms-in-app.) The former skiller taxonomy
  types/queries (`Skill`, `jobRoleMatch`, `similarJobRoles`,
  `mostPopularSkills`, `jobRoleCount`, …) **and** the skill-path session types/queries
  (`getOrCreateSkillPathSession`, `completeSkillPathStep`, …) are all served by the **backend** subgraph;
  `categoryTree`/`fullCategoryTree` were dropped, not ported.
- **No skiller container / repo / schema search-path.** Not in `repos.yml` or `docker-compose.yml`; the app
  DB connection uses the default `public` search_path (no `search_path=skiller`); `app` subscribes to the
  `skiller` Redis stream **in-process** (both ends now inside app).
- **Clean-bring-up prerequisite:** the merged migrations create the taxonomy vector columns as
  `extensions.vector(1536)` and a GIN-trigram index via `extensions.gin_trgm_ops`, so the **`extensions`
  schema (pgvector + `pg_trgm`) must be bootstrapped before `make migrate`** on a clean DB — else app
  `20260518125439` and cms `20250116133510` fail with `schema "extensions" does not exist`. (Bring-up
  ordering, tracked for M211; not a merge defect.)

**Live de-risk (2026-07-08):** a cold containerized `make up` on stack-dev built the 86-commit merged
image and brought up the federation with **no skiller container** (`SKILLER_RPC_ADDR=http://backend:8083`)
— 4 subgraphs as it stood then; skillpath, jobsimulation and cms have since also merged into `app`, so the
current supergraph is **1 subgraph** (backend).
A clean-slate `make reset-db` + `make migrate` created the full `public` taxonomy from scratch —
`public.skills` (with an `organization_id` column), `job_roles`, `job_role_skills`, `skill_embeddings`,
`categories`, `specializations` — with **no `skiller` schema on a clean DB**, once the `extensions` schema
was bootstrapped (see prerequisite above).

## Architecture & Code Map

* **Codebase**: `app` (local) — repo `git@github.com:anthropos-work/app`
* **Language**: Go 1.26
* **Database**: PostgreSQL `public` schema (Ent ORM + Atlas migrations)
* **Ports**: 8082 (HTTP/GraphQL — `PORT`), 8083 (Connect-RPC — `RPC_PORT`), 8084 (meta/health — `META_PORT`). Container publishes 8081/8082/8083; 8081 is reserved/unused.
* **Profile**: `core` (the default), `backend`, `all` — `profiles: [core, backend, all]` (`docker-compose.yml:100`, derived from `docker-compose.yml` @ platform `0dab54d`). The default profile is `core`, not `graphql`: `0dab54d` renamed it. Corrected M257x iter-68
* **Versioning**: Semantic; CHANGELOG.md is generated from conventional commits. Tags trigger production deploys.

### Key directories

```
main.go, rpc.go             Entry points
cmd/                        CLIs (bootstrap-org, migrations utilities)
internal/
  academy/                  Server-owned academy domain (chapter progress + catalog: academy_series /
                            academy_skill_paths / academy_chapters / academy_chapter_bodies). These ARE the
                            tables "Talk to Data" reads. NB: the legacy `internal/aiacademy` sync package and
                            its `aiacademy_courses` read-model were REMOVED when this domain took ownership
                            (`internal/academy/academy.go:6-9` states so).
  admin/                    Admin operations
  aiusage/                  AI usage / cost tracking ledger (AI Redis Stream)
  analytics/                PostHog / internal analytics
  app/                      Component wire-up
  askengine/                "Talk to Data" — SSE streaming SQL Q&A
    rules.md                Source of truth for SQL guardrails + business rules
    bedrock.go              AWS Bedrock client middleware
    sandbox.go              SQL validator (whitelist + read-only enforcement)
    executor.go             SQL execution & streaming
    followups.go            Follow-up suggestion extraction
  assignments/              Assignment lifecycle
  authorization/            Sentinel client
  bootstrap/                First-run / new-org provisioning
  cache/                    Redis caching layer
  clerk/                    Clerk webhook handlers
  companysearch/            Company search (LinkedIn / external sources)
  converter/                gotenberg.go for Office → PDF
  cors/                     CORS configuration
  data/ent/                 Ent schema + generated code (public schema)
  deadletterqueue/          DLQ handling for Redis Streams
  experiencepoint/          User XP tracking
  cms/                      Merged cms domain: directus edge, similarity, studio, library, importer/exporter (cms-in-app)
  jobsimfeedback/           Post-session signal routing
  jobsimulation/            Merged jobsim domain: session engine, actors, calls, recording, anticheat, analytics (jobsim-in-app)
  jobsimwiring/             Single construction entry point for the merged jobsim engine
  jobsimulations/           Backend's view of jobsim data
  labs/session/             AI Labs LabSession RPC handlers (+ labs/labsapi, labs/adapter, labs/catalog)
  linkedin/                 LinkedIn import / profile sync
  meta/                     Metadata utilities
  organization/             Org domain logic
  payments/                 Stripe integration
  resource/                 Resource entities
  roles/                    User roles
  rpc/                      Connect-RPC server
  set/                      Set / collection utilities
  skill/                    Skill domain
  skiller/, skillerai/      Merged skiller domain: taxonomy, embeddings, AI matching (skiller-in-app)
  skillpaths/               Backend's view of skillpath data
  subscriptions/            Subscription lifecycle
  taxonomy/                 Taxonomy access
  templates/                Email / message templates
  user/                     User domain
  utils/                    Shared helpers
  web/                      HTTP + GraphQL handlers
    backend/graphql/graph/schemas/   Federation v2 GQL schemas
  worker/                   Redis Streams consumers (Watermill)
  workforce/                Workforce analytics aggregations
```

## Recent Feature Additions (Q1-Q2 2026)

* **Talk to Data** (v1.266.0+, May 2026): SSE-streaming Q&A on the platform's data. Bedrock-backed Anthropic streaming, SQL validation sandbox in `internal/askengine/sandbox.go`, business rules in `internal/askengine/rules.md`. Has its own conversation table and rate-limiting.
* **Workforce analytics** (v1.266.2): Skill + sim aggregations across org members with date filtering.
* **AI Readiness** (v1.266+, the `internal/aireadiness` package): org-level AI-capability diagnostics — a 3-step onboarding/evaluation (skill-mapping 30 → simulation 40 → interview 30) yielding a per-member score + archetype, an org **manager dashboard** (funnel + Knowledge×Usage matrix + per-team/person drill-down), **org-gated** via `organization_settings.ai_readiness`, with persisted LLM diagnosis narratives. Engine: its own top-level package **`app/internal/aireadiness/`** (`manager.go`, `cycles.go`,
  `diagnosis.go`, `compare.go`, `csv.go`, …) — **not** `internal/workforce/`, which contains no `readi*`
  file at HEAD; GraphQL `graph/schemas/ai_readiness.graphqls`; ~10 `/api/workforce/ai-readiness*` REST handlers + an `ai_readiness_refresh` worker task; **13** `ai_readiness_*` ent tables (`select table_name … where table_name like 'ai_readiness%'` on a migrated stack). The four a "9" omits — the ones a seeder or schema audit then misses — are `ai_readiness_recommendations` (M219), `ai_readiness_email_overrides` (M408) and the notification **pair** `ai_readiness_notification_logs` + `ai_readiness_notification_optouts` (M400/M403). (`ai_readiness_live_snapshots` and `ai_readiness_text_translations` were already among the original 9 — they are *not* omissions.) **Full doc, with the authoritative per-table breakdown: [`ai-readiness.md`](ai-readiness.md).**
* **Hiring talk-to-data** (`feat/hiring-talk-to-data` branch): Variant scoped to hiring workflows.
* **Bedrock task role policy statements** (v1.267.1): IAM additions for Bedrock model access from the prod ECS task role.
* **Company context (M1/M2)** (`feat/company-context-m1m2` branch): Org-level context propagation through AI calls.
* **Taxonomy translations** (`feat/taxonomy-translations` branch): Localized skill/role labels.
* **AI Labs LabSession** (Phase B PR 2, #896): Connect-RPC `lab.v1.LabSessionService` (Create/Get/List/Cancel/ReportEvent) plus a new `lab_sessions` Ent table — `id` supplied by labs-api as a 12-char hex (not a UUID); `user_id`, `organization_id` (optional — empty for individual payers), `template`, `mode` (test/build/teach), `status` (booting/ready/grading/stopped/failed/cancelled), `budget_usd`/`spend_usd`/`total_tokens`, `started_at`/`stopped_at`, `grade_result` JSON. Registered as a third RPC handler in `main.go` after Users and Organizations. **The real HTTP client has since LANDED** and is wired conditionally — `main.go:743-746` @ `app` `b948604` v1.366.0: `if labsAPIURL := os.Getenv("LABS_API_URL"); labsAPIURL != "" { labsAPI = adapter.New(labsapi.New(...)) }`, backed by real `internal/labs/labsapi/` + `internal/labs/adapter/` packages. `LabsAPIClient` is nil **only** on the unset-`LABS_API_URL` local-dev path, where Create persists the LabSession row without booting a VM (no `ide_url`/`preview_url`) and Cancel marks the row cancelled without calling labs-api.

## Interface Discovery

* **GraphQL Federation**: schemas at `internal/web/backend/graphql/graph/schemas/*.graphqls`. Federated into the supergraph as the `backend` subgraph — **the only one left**, and on a local stack the frontends now reach it directly at `:8082/graphql/query` rather than through a router.
* **Connect-RPC**: `rpc.go` is the top-level wire-up. Look there for the implemented services. The only remaining external caller is **messenger**, and **all four of its addresses point here.** At platform `0dab54d`, `docker-compose.yml:173-183` sets `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` and `SKILLER_RPC_ADDR` all to `http://backend:8083`, under its own comment *"cms + jobsimulation are folded into app: all four RPC edges are the one backend mux"* (`:171-172`). **The M809 re-point has landed** — the earlier two-of-four split (`http://cms:8091` / `http://jobsimulation:8401`, true at platform `2adcf71`) is history, not current configuration. `app`'s own source comment still says *"additive + DORMANT: external callers (messenger) keep hitting the standalone cms via `CMS_RPC_ADDR` **until the M809 re-point**"* (`app/main.go:1205-1211` @ `b948604` v1.366.0) — **that comment is now stale in `app`**; grade the address against compose, not against the comment. In production terraform the re-pointed pair is at `http://backend.internal.anthropos:8081`. Services include `lab.v1.LabSessionService`, `SkillerService` (`internal/rpc/skillerrpc/`), `JobSimulationService` and `CMSService`. Note the RPC server runs with a **60s write timeout** — the ported skiller RAG/LLM methods can exceed the old 10s default.
* **HTTP** (port 8082): Clerk webhooks, payment webhooks, document upload/convert endpoints, "Talk to Data" SSE.

### Upstream consumers

* Next Web App (GraphQL via **`backend`'s own endpoint, `:8082/graphql/query`** on a local stack since platform `2adcf71` deleted the router — the Cosmo Router is prod-only now — plus direct HTTP for SSE and webhooks)
* Hiring App
* Mobile App
* Studio-Desk (for org-level metadata)

### Downstream dependencies

* **Sentinel** — authz on every request
* **Storage** — file uploads
* **Directus** (`content.anthropos.work`) — the external content edge read by the in-process cms domain
* **Judge0** — sandboxed code execution, called directly (`JUDGE0_BASE_URL`) since roadrunner merged in
* **LiveKit / AWS Chime** — simulation voice + recording, for the in-process jobsim engine
* **Gotenberg** — Office → PDF conversion
* **PostgreSQL** (`public` schema), **Redis** (cache + streams)
* **External**: Clerk (auth), Stripe (payments), Customer.io, PostHog, Bedrock (AI), AI providers via the shared `ai` library (embeddings + skill matching — merged skiller domain), Brevo (via Messenger), Sentry

### Redis Streams

* app is **both producer and consumer** of all five application streams: `backend`, `skiller`, `skillpath`, `jobsimulation`, `cms`, plus the `AI`/`ai_usage` usage stream
* The one external producer left is **Directus**, whose webhooks feed the `cms` stream
* Each stream has exactly **one** subscriber with multiple handlers merged via `.AddHandler(...)` — colony keys by stream name, so a second `AddSubscriber` for the same stream silently overwrites the first

## Local Development

### Run in Docker

```bash
cd platform
make up                # default graphql profile — recommended
# or just backend:
make up PROFILE=backend # also starts postgresql, redis, sentinel, gotenberg
```

### Run natively

```bash
cd platform
make dev S=backend
cd ../app
make setup             # mockgen, ent, atlas
make gen               # protobuf, ent, gqlgen codegen
go run .
```

You'll need `platform/.env` reachable (or copy relevant vars). The infra services should still run via Docker.

### Migrations

```bash
cd platform
make migrate S=app
```

Versioned Atlas migrations live in `terraform/migrations/` (per `atlas.hcl`: `dir = "file://terraform/migrations"`, source `ent://internal/data/ent/schema`), not in the top-level `migrations/` dir (which holds only `atlas.sum`). Generate a new migration after an Ent schema change with `make migrations` (`atlas migrate diff --env local`); apply with `atlas migrate apply --env local` (or `make migrate S=app`).

The `public` schema is the largest in the platform; the most recent set of migrations (May 2026) touched simulation-type definitions and content JSON defaults.

## Testing

```bash
go test ./...
# Heavy components have isolated test suites:
go test ./internal/askengine/...
```

## Related Documentation

* [AI Architecture](../architecture/ai_architecture.md) — Bedrock routing, cost tracking
* [CMS](./cms.md), [Jobsimulation](./jobsimulation.md) — downstream services
* [Skiller](./skiller.md) — the former standalone skills-taxonomy service, merged into app (July 2026)
* [Skillpath](./skillpath.md) — the former standalone skill-path runtime engine, merged into app ("skillpath-in-app", M502→M507)
* [Gotenberg](./gotenberg.md) — PDF conversion sidecar
