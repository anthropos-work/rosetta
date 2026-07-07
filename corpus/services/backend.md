# Backend Service (`app`)

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
* **AI Labs LabSession** (`internal/labsession`) — Connect-RPC `lab.v1.LabSessionService` (Create/Get/List/Cancel/ReportEvent) plus a `lab_sessions` Ent table. The labs-api client is currently wired as nil, so Create persists a session row without booting a VM and Cancel marks the row cancelled without calling labs-api (see Recent Feature Additions).
* **Document → PDF conversion** (`internal/converter/gotenberg.go`) — via the Gotenberg service

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
  labsession/               AI Labs LabSession RPC handlers
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
* **Skiller-in-app merge** (July 2026): the standalone `skiller` microservice was merged into `app`. Taxonomy/embeddings Ent models moved into `internal/data/ent/schema/` (data ported from the old `skiller` DB schema into `public`; the old schema is legacy), the `SkillerService` Connect-RPC surface is served from `internal/rpc/skillerrpc/` (`GetSkills`, `GetSkill`, `SearchSkill`, `MatchSkill`, `GetJobRole` are the externally-reached methods), and the taxonomy GraphQL queries were folded into app's subgraph (the skiller subgraph was removed from federation; `categoryTree` / `fullCategoryTree` were dropped, not ported). The internal app→skiller RPC path is retired (PR #989). See [skiller.md](./skiller.md).
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
