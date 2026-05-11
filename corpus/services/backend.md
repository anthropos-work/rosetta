# Backend Service (`app`)

## Role & Responsibility

`app` is the **main API gateway** of the platform — the service that frontends, hiring apps, and other backend services talk to first. It owns the `public` schema (users, organizations, memberships, assignments, subscriptions, payments) and exposes:

* **GraphQL Federation v2 subgraph** for high-level user / organization / assignment queries
* **Connect-RPC** for inter-service calls (consumed by skiller, jobsimulation, skillpath, cms)
* **HTTP** endpoints on port 8082 for webhooks and miscellaneous integrations

It also hosts a growing number of cross-cutting features that don't fit neatly into any other service:

* **Talk to Data** (`internal/ask`, `internal/askengine`) — SSE-streaming natural-language Q&A over the platform's data, powered by Bedrock (Anthropic) with a SQL-validation sandbox. Added 2026-Q2 (v1.266+).
* **Workforce analytics** (`internal/workforce`) — aggregations of skills, simulations, and growth across org members
* **Job-simulation feedback** (`internal/jobsimfeedback`) — post-session signals routed back to skiller
* **AI usage / cost tracking** (`internal/aiusage`) — central ledger driven by the `AI` Redis Stream
* **Bootstrap & admin** (`internal/admin`, `internal/bootstrap`, `cmd/bootstrap-org`) — provisioning utilities
* **Copilot** (`internal/copilot`) — internal assistant flows
* **Document → PDF conversion** (`internal/converter/gotenberg.go`) — via the Gotenberg service

## Architecture & Code Map

* **Codebase**: `app` (local) — repo `git@github.com:anthropos-work/app`
* **Language**: Go 1.25
* **Database**: PostgreSQL `public` schema (Ent ORM + Atlas migrations)
* **Ports**: 8081 (HTTP), 8082 (HTTP/GraphQL), 8083 (Connect-RPC)
* **Profile**: `graphql` (default) and `backend`
* **Versioning**: Semantic; CHANGELOG.md is generated from conventional commits. Tags trigger production deploys.

### Key directories

```
main.go, rpc.go             Entry points
cmd/                        CLIs (bootstrap-org, migrations utilities)
internal/
  admin/                    Admin operations
  ai/, aiusage/             AI provider wrapper consumers + cost tracking
  analytics/                PostHog / internal analytics
  app/                      Component wire-up
  ask/, askengine/          "Talk to Data" — SSE streaming SQL Q&A
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
  data/ent/                 Ent schema + generated code (public schema)
  deadletterqueue/          DLQ handling for Redis Streams
  experiencepoint/          User XP tracking
  jobsimfeedback/           Post-session signal routing
  jobsimulations/           Backend's view of jobsim data
  linkedin/                 LinkedIn import / profile sync
  meta/                     Metadata utilities
  organization/             Org domain logic
  payments/                 Stripe integration
  resource/                 Resource entities
  roles/                    User roles
  rpc/                      Connect-RPC server
  set/                      Set / collection utilities
  skill/                    Skill domain
  skiller/, skillpaths/     Backend's view of skiller/skillpath data
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
* **Hiring talk-to-data** (`feat/hiring-talk-to-data` branch): Variant scoped to hiring workflows.
* **Bedrock task role policy statements** (v1.267.1): IAM additions for Bedrock model access from the prod ECS task role.
* **Company context (M1/M2)** (`feat/company-context-m1m2` branch): Org-level context propagation through AI calls.
* **Taxonomy translations** (`feat/taxonomy-translations` branch): Localized skill/role labels.

## Interface Discovery

* **GraphQL Federation**: schemas at `internal/web/backend/graphql/graph/schemas/*.graphqls`. Federated into the Cosmo Router supergraph as the `backend` subgraph.
* **Connect-RPC**: `rpc.go` is the top-level wire-up. Look there for the implemented services. Used by skiller, jobsim, skillpath, cms via `BACKEND_USERS_RPC_ADDR=http://backend:8083`.
* **HTTP** (port 8082): Clerk webhooks, payment webhooks, document upload/convert endpoints, "Talk to Data" SSE.

### Upstream consumers

* Next Web App (GraphQL via Cosmo, plus direct HTTP for SSE and webhooks)
* Hiring App
* Mobile App
* Studio-Desk (for org-level metadata)

### Downstream dependencies

* **Sentinel** — authz on every request
* **CMS** — content RPC for assignments, simulation metadata
* **Skiller** — taxonomy and matching RPC
* **Skillpath** — skill-progression queries
* **Storage** — file uploads
* **Gotenberg** — Office → PDF conversion
* **PostgreSQL** (`public` schema), **Redis** (cache + streams)
* **External**: Clerk (auth), Stripe (payments), Customer.io, PostHog, Bedrock (AI), Brevo (via Messenger), Sentry

### Redis Streams

* Producer: `backend` stream (user/org updates)
* Consumer: `cms`, `skiller` events; `AI` usage stream (also produces)

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

The `public` schema is the largest in the platform; the most recent set of migrations (May 2026) touched simulation-type definitions and content JSON defaults.

## Testing

```bash
go test ./...
# Heavy components have isolated test suites:
go test ./internal/askengine/...
```

## Related Documentation

* [AI Architecture](../architecture/ai_architecture.md) — Bedrock routing, cost tracking
* [CMS](./cms.md), [Skiller](./skiller.md), [Skillpath](./skillpath.md), [Jobsimulation](./jobsimulation.md) — downstream services
* [Gotenberg](./gotenberg.md) — PDF conversion sidecar
