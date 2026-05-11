# Skiller Service

## Role & Responsibility

Skiller is the **skills graph** of the platform. It owns the 60K-skill / 18K-job-role Anthropos taxonomy, computes vector embeddings for skills and roles (for RAG-based matching), and exposes job-role matching and skill-management APIs to the rest of the platform.

This is the AI-heaviest pure-Go service: every other AI surface (jobsimulation, cms) uses skiller for skill-aware behavior.

## Architecture & Code Map

* **Codebase**: `skiller` (Local directory; repo `git@github.com:anthropos-work/skiller`)
* **Language**: Go 1.25
* **Frameworks**: gqlgen (GraphQL Federation v2), Connect-RPC, Ent ORM, **Asynq** (background task queue)
* **Database**: PostgreSQL `skiller` schema, with `pgvector` from the `extensions` schema
* **Ports**: 8085 (GraphQL), 8086 (Connect-RPC)
* **Profile**: `graphql` (default) and `skiller`
* **Deploy**: Docker → ECR → ECS

### Key directories

```
main.go                       Entry point
cmd/
  importer/                   Bulk taxonomy import tool
  jrembeddings/               Job-role embeddings backfill CLI
  skillembeddings/            Skill embeddings backfill CLI
  jobroleMeta/                Job-role metadata utility
  jobroleSkills/              Job-role ↔ skill linking utility
  skillmatchbenchmark/        Benchmark/eval harness
internal/
  ai/                         AI integration (matching prompts)
  authorization/              Sentinel client
  cache/                      Redis caching
  embeddings/                 Vector embedding generation + storage
  rag/                        Retrieval-augmented generation for role matching
  jobrole/                    Job-role business logic
  organization/               Org-scoping for taxonomy
  rpcsrv/                     Connect-RPC server
  search/                     Skill / role search
  taxonomy/                   Anthropos taxonomy ops (60K skills, 18K roles)
  templates/                  Prompt templates
  worker/                     Async background workers
graph/
  schemas/schema.graphqls     GraphQL contract (federated by Cosmo Router)
ent/schema/                   Ent entity definitions:
                              skill.go, jobrole.go, category.go, specialization.go,
                              jobroleEmbeddings.go, skillEmbeddings.go,
                              jobroleskill.go, jobroleCategory.go, mixin.go
```

## Vector Storage (2026-Q2)

Embeddings live in dedicated tables, not on the entity tables themselves. Migrations `20260417103036` and `20260417120309` created the new layout and dropped the old denormalized columns.

```
job_role_embeddings(
  id BIGSERIAL PK,
  job_role_id UUID FK → job_roles.id,
  small_embedding3 extensions.vector(1536),  -- OpenAI text-embedding-3-small
  -- IVFFLAT index on small_embedding3
)

skill_embeddings(
  id BIGSERIAL PK,
  skill_id UUID FK → skills.id,
  small_embedding3 extensions.vector(1536),
  -- IVFFLAT index on small_embedding3
)
```

The `extensions` schema (which houses `pgvector`) must exist before applying these migrations. The setup guide creates it as part of the first-run flow.

See [AI Architecture → Embeddings & RAG](../architecture/ai_architecture.md#embeddings--rag-skiller) for the full picture.

## Interface Discovery

* **GraphQL**: `graph/schemas/schema.graphqls`. Skiller is one of 5 subgraphs in the Cosmo Router federation.
* **Connect-RPC**: `internal/rpcsrv/`. Consumed via `SKILLER_RPC_ADDR=http://skiller:8086`.

### Notable RPC operations

* `MatchJobRole(name, context, organization_id)` — AI-powered job role matching using embeddings + RAG
* Job-role and skill CRUD types

### Upstream consumers

* Backend (`app`) — user-skill data, taxonomy queries
* CMS — skill metadata for content
* Jobsimulation — skill metadata during sessions
* Skillpath — skill progression contexts

### Downstream dependencies

* **Sentinel** — authz
* **Backend (app)** — user data via RPC (`BACKEND_USERS_RPC_ADDR=http://backend:8083`)
* **AI providers** — OpenAI + Anthropic via the shared `ai` library (embeddings + matching)
* PostgreSQL (with pgvector), Redis

## Local Development

### Run in Docker

```bash
cd platform
make up                       # default graphql profile
# or just skiller:
make up PROFILE=skiller
```

### Run natively

```bash
cd platform
make dev S=skiller
cd ../skiller
go generate ./...             # gqlgen + Ent codegen
go run .
```

### Migrations

```bash
cd platform
make migrate S=skiller
```

The `extensions` schema must exist (run setup_guide §6 once) before vector-table migrations will apply.

### Backfill embeddings

After importing or updating the taxonomy:

```bash
cd skiller
go run ./cmd/jrembeddings    # Job-role embeddings
go run ./cmd/skillembeddings # Skill embeddings
```

### Benchmarks

```bash
go run ./cmd/skillmatchbenchmark
```

## Testing

```bash
go test ./...
```

## Related Documentation

* [AI Architecture](../architecture/ai_architecture.md) — embeddings, RAG, provider routing
* [Backend](./backend.md) — user identity / org scoping
* [Dependency Map](../architecture/dependency_map.md)
