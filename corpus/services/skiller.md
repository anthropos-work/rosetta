# Skiller Service

## Role & Responsibility

Skiller is the **skills graph** of the platform. It owns the 60K-skill / 18K-job-role Anthropos taxonomy, computes vector embeddings for skills and roles (for RAG-based matching), and exposes job-role matching and skill-management APIs to the rest of the platform.

This is the AI-heaviest pure-Go service: every other AI surface (jobsimulation, cms) uses skiller for skill-aware behavior.

The 60K-skill / 18K-role taxonomy **data** is loaded into the skiller DB by the `importskills` and `importjobroles` cobra subcommands (`cmd/importSkills`, `cmd/importJobRole`), which read the CSVs and call `TaxonomyManager.CreateSkill` / `JobRoleManager.CreateJobRole`. The `anthropos-work/taxonomy` library only supplies NodeID generation helpers, not the data.

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
  importSkills/               Bulk skill/specialization/category import from CSV (cobra subcommand: importskills)
  importJobRole/              Bulk job-role import + skill linking from CSV (cobra subcommand: importjobroles)
  importer/                   Standalone taxonomy-directory importer (reads a skill-taxonomy/ path arg)
  jrembeddings/               Job-role embeddings backfill CLI
  skillembeddings/            Skill embeddings backfill CLI
  jrtranslations/             Translate job roles into ContentLanguages (cobra subcommand: jrtranslations)
  skilltranslations/          Translate skills into ContentLanguages (cobra subcommand: skilltranslations)
  backfilltranslations/       Seed english *_translations rows from legacy fields (cobra subcommand: backfilltranslations)
  jobroleMeta/                Job-role metadata utility
  jobroleSkills/              Job-role ↔ skill linking utility
  skillmatchbenchmark/        Benchmark/eval harness
internal/
  ai/                         AI integration (matching prompts)
  authorization/              Sentinel client
  cache/                      Redis caching
  content/                    Localized content access
  embeddings/                 Vector embedding generation + storage
  localization/               Per-skill / per-role translation management
  rag/                        Retrieval-augmented generation for role matching
  jobrole/                    Job-role business logic
  organization/               Org-scoping for taxonomy
  rpcsrv/                     Connect-RPC server
  search/                     Skill / role search
  taxonomy/                   Anthropos taxonomy ops (60K skills, 18K roles)
  templates/                  Prompt templates
  translation/                Translation generation pipeline
  worker/                     Async background workers
graph/
  schemas/schema.graphqls     GraphQL contract (federated by Cosmo Router)
ent/schema/                   Ent entity definitions:
                              skill.go, jobrole.go, category.go, specialization.go,
                              jobroleEmbeddings.go, skillEmbeddings.go,
                              skillTranslation.go, jobroleTranslation.go,
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

## Localization / Multilingual content

Skiller stores per-skill / per-job-role translations (`skill_translations`, `job_role_translations` tables; `ent/schema/skillTranslation.go`, `jobroleTranslation.go`) across 8 `ContentLanguage`s (english, italian, spanish, french, german, dutch, japanese, portuguese).

Most GraphQL queries/mutations accept an optional `language: ContentLanguage` arg (`skillDetails`, `skillsByName`, `matchSkill`, `jobRoleDetails`, `matchJobRole`, `jobRoleSkills`, etc.); `Skill` / `JobRole` expose `language` and `availableLanguages`.

Translations are generated/seeded via cobra subcommands of the skiller binary:

```bash
go run . skilltranslations <node_id>...
go run . jrtranslations <node_id>...
go run . backfilltranslations          # seeds english rows from legacy fields; idempotent
```

A `localizationManager` (`cmd/root.go:172`) is wired into the Connect-RPC server.

## Interface Discovery

* **GraphQL**: `graph/schemas/schema.graphqls`. Skiller is one of 5 subgraphs in the Cosmo Router federation.
* **Connect-RPC**: `internal/rpcsrv/`. Consumed via `SKILLER_RPC_ADDR=http://skiller:8086`.

### Notable RPC operations

(per `internal/rpcsrv/rpc.go`)

* `MatchJobRole(name, context, organization_id)` — AI-powered job role matching using embeddings + RAG
* `MatchMultipleSkills` — batch skill matching
* `MatchMultipleJobRoles` — batch job-role matching
* `GetSimilarJobRoles` — nearest-neighbour job roles via embeddings
* `GetJobRoleMatch` — fetch a previously computed job-role match
* `FilterSkills` — filtered skill lookup
* `SearchSkill` — skill search
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
go generate ./...             # runs all 3 //go:generate directives: gqlgen (generate.go), Ent/entc (ent/generate.go), mockgen (internal/cache/memoize.go)
go run .
```

`go generate ./...` regenerates Ent (`ent/generate.go` → `go run entc.go`), gqlgen, and mockgen mocks — all committed. The `ent`/atlas tooling (installed by `make setup`) is needed to produce **migrations** via `make migrations` (`atlas migrate diff --env local`).

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
