# Backend Service (`app`)

> **Since the skiller-in-app merge (v2.1 "quick change", July 2026), `app` also owns the skills-taxonomy
> domain** — the 60K+ skills graph, embeddings, and AI matching formerly owned by the standalone
> [skiller](./skiller.md) service. See the authoritative [**§ Skiller-in-app merge — fact-sheet**](#skiller-in-app-merge--fact-sheet-v21-quick-change)
> below (the merged-shape contract this release grades against). The body of this doc was re-pointed to the
> merged shape in M210 of the v2.1 release.

## Role & Responsibility

`app` is the **main API gateway** of the platform — the service that frontends, hiring apps, and other backend services talk to first. It owns the `public` schema (users, organizations, memberships, assignments, subscriptions, payments) and, since the **skiller-in-app merge (July 2026)**, the **skills taxonomy domain** — the 60K+ skills graph, skill/job-role embeddings, and AI skill matching formerly owned by the standalone [skiller](./skiller.md) service. It exposes:

* **GraphQL Federation v2 subgraph** for high-level user / organization / assignment queries — plus the taxonomy types/queries absorbed from the former skiller subgraph (`graph/schemas/skiller_taxonomy.graphqls`)
* **Connect-RPC** for inter-service calls (consumed by jobsimulation, skillpath, cms, messenger) — including the **skiller RPC surface** (`SkillerService`), now served by app
* **HTTP** endpoints on port 8082 for webhooks and miscellaneous integrations

It also hosts a growing number of cross-cutting features that don't fit neatly into any other service:

* **Talk to Data** (`internal/askengine`) — SSE-streaming natural-language Q&A over the platform's data, powered by Bedrock (Anthropic) with a SQL-validation sandbox. Added 2026-Q2 (v1.266+).
* **Workforce analytics** (`internal/workforce`) — aggregations of skills, simulations, and growth across org members
* **Job-simulation feedback** (`internal/jobsimfeedback`) — post-session signals routed back to the skills domain (in-process since the skiller merge)
* **AI usage / cost tracking** (`internal/aiusage`) — central ledger driven by the `AI` Redis Stream
* **Bootstrap & admin** (`internal/admin`, `internal/bootstrap`, `cmd/bootstrap-org`) — provisioning utilities
* **Copilot** (`internal/copilot`) — internal assistant flows
* **AI Labs LabSession** (`internal/labs/session`; siblings `internal/labs/labsapi`, `internal/labs/adapter`, `internal/labs/catalog`) — Connect-RPC `lab.v1.LabSessionService` (Create/Get/List/Cancel/ReportEvent) plus a `lab_sessions` Ent table. The labs-api client is currently wired as nil, so Create persists a session row without booting a VM and Cancel marks the row cancelled without calling labs-api (see Recent Feature Additions).
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
  is this count; taxonomy grows over time.)
- **RPC re-pointed** — the `SkillerService` Connect-RPC surface is served **by app itself**
  (`internal/rpc/skillerrpc/`). Consumers keep the env var, re-pointed: `SKILLER_RPC_ADDR=http://backend:8083`
  locally (all four occurrences in the merged `docker-compose.yml`), `http://backend:8081` in prod terraform.
- **Federation is now 4 subgraphs** (the skiller subgraph was removed; `schemas/skiller.graphqls` deleted at
  `graphql-wundergraph@c284453`): **backend**, **jobsimulation**, **cms**, **skillpath**. The former skiller
  taxonomy types/queries (`Skill`, `jobRoleMatch`, `similarJobRoles`, `mostPopularSkills`, `jobRoleCount`, …)
  are served by the **backend** subgraph; `categoryTree`/`fullCategoryTree` were dropped, not ported.
- **No skiller container / repo / schema search-path.** Not in `repos.yml` or `docker-compose.yml`; the app
  DB connection uses the default `public` search_path (no `search_path=skiller`); `app` subscribes to the
  `skiller` Redis stream **in-process** (both ends now inside app).
- **Clean-bring-up prerequisite:** the merged migrations create the taxonomy vector columns as
  `extensions.vector(1536)` and a GIN-trigram index via `extensions.gin_trgm_ops`, so the **`extensions`
  schema (pgvector + `pg_trgm`) must be bootstrapped before `make migrate`** on a clean DB — else app
  `20260518125439` and cms `20250116133510` fail with `schema "extensions" does not exist`. (Bring-up
  ordering, tracked for M211; not a merge defect.)

**Live de-risk (2026-07-08):** a cold containerized `make up` on stack-dev built the 86-commit merged
image and brought up the 4-subgraph federation with **no skiller container** (`SKILLER_RPC_ADDR=http://backend:8083`).
A clean-slate `make reset-db` + `make migrate` created the full `public` taxonomy from scratch —
`public.skills` (with an `organization_id` column), `job_roles`, `job_role_skills`, `skill_embeddings`,
`categories`, `specializations` — with **no `skiller` schema on a clean DB**, once the `extensions` schema
was bootstrapped (see prerequisite above).

## Architecture & Code Map

* **Codebase**: `app` (local) — repo `git@github.com:anthropos-work/app`
* **Language**: Go 1.26
* **Database**: PostgreSQL `public` schema (Ent ORM + Atlas migrations)
* **Ports**: 8082 (HTTP/GraphQL — `PORT`), 8083 (Connect-RPC — `RPC_PORT`), 8084 (meta/health — `META_PORT`). Container publishes 8081/8082/8083; 8081 is reserved/unused.
* **Profile**: `graphql` (default) and `backend`
* **Versioning**: Semantic; CHANGELOG.md is generated from conventional commits. Tags trigger production deploys.

### Key directories

```
main.go, rpc.go             Entry points
cmd/                        CLIs (bootstrap-org, migrations utilities)
internal/
  admin/                    Admin operations
  aiacademy/                Periodic AI Academy catalog sync (fetches catalog.json, populates aiacademy_courses for Talk to Data)
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
  copilot/                  Internal copilot flows
  cors/                     CORS configuration
  data/ent/                 Ent schema + generated code (public schema)
  deadletterqueue/          DLQ handling for Redis Streams
  experiencepoint/          User XP tracking
  jobsimfeedback/           Post-session signal routing
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
* **AI Readiness** (v1.266+, the `internal/workforce` subsystem): org-level AI-capability diagnostics — a 3-step onboarding/evaluation (skill-mapping 30 → simulation 40 → interview 30) yielding a per-member score + archetype, an org **manager dashboard** (funnel + Knowledge×Usage matrix + per-team/person drill-down), **org-gated** via `organization_settings.ai_readiness`, with persisted LLM diagnosis narratives. Engine: `internal/workforce/ai_readiness.go` + `readiness_steps.go` + `readiness_narrative.go`; GraphQL `graph/schemas/ai_readiness.graphqls`; ~10 `/api/workforce/ai-readiness*` REST handlers + an `ai_readiness_refresh` worker task; 9 `ai_readiness_*` ent tables. **Full doc: [`ai-readiness.md`](ai-readiness.md).**
* **Hiring talk-to-data** (`feat/hiring-talk-to-data` branch): Variant scoped to hiring workflows.
* **Bedrock task role policy statements** (v1.267.1): IAM additions for Bedrock model access from the prod ECS task role.
* **Company context (M1/M2)** (`feat/company-context-m1m2` branch): Org-level context propagation through AI calls.
* **Taxonomy translations** (`feat/taxonomy-translations` branch): Localized skill/role labels.
* **AI Labs LabSession** (Phase B PR 2, #896): Connect-RPC `lab.v1.LabSessionService` (Create/Get/List/Cancel/ReportEvent) plus a new `lab_sessions` Ent table — `id` supplied by labs-api as a 12-char hex (not a UUID); `user_id`, `organization_id` (optional — empty for individual payers), `template`, `mode` (test/build/teach), `status` (booting/ready/grading/stopped/failed/cancelled), `budget_usd`/`spend_usd`/`total_tokens`, `started_at`/`stopped_at`, `grade_result` JSON. Registered as a third RPC handler in `main.go` after Users and Organizations. The labs-api client (`LabsAPIClient`) is wired as nil for now, so Create persists the LabSession row but does not boot a VM (no `ide_url`/`preview_url` returned) and Cancel marks the row cancelled without calling labs-api; the real HTTP client that drives VM lifecycle lands in PR 6.

## Interface Discovery

* **GraphQL Federation**: schemas at `internal/web/backend/graphql/graph/schemas/*.graphqls`. Federated into the Cosmo Router supergraph as the `backend` subgraph.
* **Connect-RPC**: `rpc.go` is the top-level wire-up. Look there for the implemented services. Used by jobsim, skillpath, cms, messenger via `BACKEND_USERS_RPC_ADDR=http://backend:8083`. Services include `lab.v1.LabSessionService` (Create/Get/List/Cancel/ReportEvent) registered in `main.go` as a third RPC handler after Users and Organizations, and `SkillerService` (`internal/rpc/skillerrpc/`) — consumers reach it via `SKILLER_RPC_ADDR=http://backend:8083` locally (`http://backend:8081` in production terraform).
* **HTTP** (port 8082): Clerk webhooks, payment webhooks, document upload/convert endpoints, "Talk to Data" SSE.

### Upstream consumers

* Next Web App (GraphQL via Cosmo, plus direct HTTP for SSE and webhooks)
* Hiring App
* Mobile App
* Studio-Desk (for org-level metadata)

### Downstream dependencies

* **Sentinel** — authz on every request
* **CMS** — content RPC for assignments, simulation metadata
* **Skillpath** — skill-progression queries
* **Storage** — file uploads
* **Gotenberg** — Office → PDF conversion
* **PostgreSQL** (`public` schema), **Redis** (cache + streams)
* **External**: Clerk (auth), Stripe (payments), Customer.io, PostHog, Bedrock (AI), AI providers via the shared `ai` library (embeddings + skill matching — merged skiller domain), Brevo (via Messenger), Sentry

### Redis Streams

* Producer: `backend` stream (user/org updates); `skiller` stream (skill score changes — both ends of this stream live inside app since the skiller merge)
* Consumer: `cms`, `jobsimulation`, `skillpath`, `skiller` events; `AI` usage stream (also produces)

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
* [CMS](./cms.md), [Skillpath](./skillpath.md), [Jobsimulation](./jobsimulation.md) — downstream services
* [Skiller](./skiller.md) — the former standalone skills-taxonomy service, merged into app (July 2026)
* [Gotenberg](./gotenberg.md) — PDF conversion sidecar
