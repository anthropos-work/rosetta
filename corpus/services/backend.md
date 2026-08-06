# Backend Service (`app`)

> ## `app` is the backend monolith
>
> **Eight former microservices now run inside `app`**, in merge order:
>
> | Merged service | Program | What moved in |
> |---|---|---|
> | [skiller](./skiller.md) | skiller-in-app (v2.1 "quick change", July 2026) | the skills-taxonomy graph (**≥42,790 skills / ≥22,470 job roles** — public subset; [not "60K/18K"](../architecture/shared_libraries.md#taxonomy-figures)), embeddings, AI matching |
> | [skillpath](./skillpath.md) | skillpath-in-app (M502→M507) | skill-path progression engine, session state |
> | [roadrunner](./roadrunner.md) | with jobsim-in-app | Judge0 code execution (called directly via `JUDGE0_BASE_URL`) |
> | [jobsimulation](./jobsimulation.md) | jobsim-in-app (prod ECS teardown **M810 — LANDED**, `6092c6d2`) | the simulation session engine — `internal/jobsimulation/`, wired by `internal/jobsimwiring/wiring.go` |
> | [cms](./cms.md) | cms-in-app v8.0, app **v1.360.0** (prod teardown **M810 — NOT MEASURABLE here**; report both, assert neither — see the *M810 prod teardown is UNEVEN* bullet below) | content layer + Directus edge + Studio — `internal/cms/` |
> | [storage](./storage.md) | v9.0 "support-in-app", 2026-08-04 | the private + public object-storage managers — `internal/storage/`, `internal/storagens/`, `internal/publicstorage/` |
> | [messenger](./messenger.md) | v9.0 "support-in-app", 2026-08-04 | transactional email (Brevo + Liquid) and messenger's **own** Redis consumer group — `internal/messenger/`; switch-gated by `MESSENGER_ENABLED` |
> | [customerio-sync](./customerio-sync.md) | v9.0 "support-in-app" | the one-way Brevo marketing-contact push — `internal/customeriosync/`; switch-gated by `CUSTOMERIO_SYNC_ENABLED` |
>
> The last three lost their **containers** a day later, at platform `838d907` (merged `0c91421`,
> 2026-08-05): compose now declares **five** services and `repos.yml` **four** entries.
>
> Consequences that hold platform-wide:
> * **The federation composes ONE subgraph** (`backend`). cms-in-app was the **3 → 1** step: the single
>   commit `graphql-wundergraph@915da06` (2026-07-29) deleted **both** `schemas/cms.graphqls` **and**
>   `schemas/jobsimulation.graphqls`, taking the supergraph from (backend, jobsimulation, cms) to
>   (backend) alone. The jobsimulation subgraph therefore **survived jobsim-in-app** and was removed
>   here, not at its own merge.
> * **All of their tables live in `public`**, with the same table names. The `skiller`, `skillpath`,
>   `jobsimulation` and `cms` DB schemas are legacy and non-authoritative.
> * **Only THREE of the eight have a Connect-RPC handler on `app`'s mux — and nothing outside the process
>   calls any of them.** The mux carries **six** handlers in total (`main.go:1297-1338` @ `app` `ad9f3c49`;
>   `:1185-1228` @ the demo pin `b948604f` — the per-handler enumeration is under *Role & Responsibility*
>   below): `SkillerService`, `JobSimulationService`, and `CMSService` **only when the Directus edge is
>   configured** (`if cmsRPCServer != nil`) — plus `UsersService`, `OrganizationsService` and
>   `lab.v1.LabSessionService`, which are `app`'s **own** surfaces and were never folded in.
>   **`skillpath`, `roadrunner`, `storage`, `messenger` and `customerio-sync` have no handler on it at
>   all.** M506 *removed* `SkillPathSessionService` rather than re-hosting it (**0** occurrences in `app`
>   Go source at `ad9f3c49`). Three of the others still *declare* a Connect service — but only in their
>   own now-frozen repos, on their own muxes, never on `app`'s: storage's `StorageService`
>   (`storage@4ce8ece5:sdk/storage/v1/service.go`), messenger's `MessengerService`
>   (`messenger@fa47850d:internal/rpcsrv/rpcsrv.go`) and roadrunner's `RoadRunnerService`
>   (`roadrunner@87d8d443:cmd/root.go:87` — `app` reaches Judge0 over plain HTTP via `JUDGE0_BASE_URL`
>   instead). For `customerio-sync` nothing can be said about its own surface from here: its repo has
>   never been in a clone set. What *is* measured is that `app`'s mux carries no handler for it.
>   **The second half of this bullet is the true half:**
>   `messenger` was the last external caller, and `838d907` deleted its container and the four
>   `*_RPC_ADDR` values that addressed the mux, so compose sets **none**.
> * **The only cross-process *Connect-RPC* edge out of `backend` on a `core` stack is `backend → sentinel`**
>   (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48` @ platform `0c91421`), and there
>   are **zero `*_RPC_ADDR` variables anywhere in compose**. **It is not the only cross-process edge:**
>   `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`,
>   `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at
>   `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL`
>   (`docker-compose.yml:59`). The correctly-scoped model wording is
>   [`architecture_overview.md:321`](../architecture/architecture_overview.md).
> * **`app` is both producer and consumer of FIVE Redis Streams** — `backend`, `skillpath`,
>   `jobsimulation`, `cms` and `ai_usage`. **Do not drop `backend` from that set** — this bullet used to,
>   while also closing the set, which made the omission a false exhaustiveness claim rather than a merely
>   short list. `app` publishes to its own self-stream at `main.go:325`
>   (`pubsub.NewPublisher(serviceName, …)`, `SERVICE_NAME` defaulting to `backend` at `:230-232`) and
>   subscribes to it in the same registration map (`:1534-1541`, `:1579-1581`; `subs[d.Streams.Backend]`
>   at `subscriber_wiring.go:248`, whose own comment reads *"Backend is app's OWN self-stream … events
>   app publishes and also consumes in-process"* — `:112-113`, all @ `app` `ad9f3c49`).
>   **`skiller` is the one `app` only CONSUMES**: nothing in `app` publishes to it, so it is a **sixth**
>   stream — not a fifth member of the both-ways set. Six subscribers, five publishers. Enumerated in
>   *Redis Streams* under *Interface Discovery* below, and in the skiller fact-sheet's
>   *No skiller container* bullet.
>   Merge handlers with `.AddHandler(...)`; a second `AddSubscriber` on one stream overwrites the first.
> * **The M810 prod teardown is UNEVEN — do not state the two together.** `cms`'s **terraform module block
>   has not moved**: it is still declared and takes no traffic (`cms/terraform/main.tf:39`
>   `service_desired_count = 0`, re-measured at `f38c0c4`). **But `cms` HAS taken an M810 step since, and
>   it cuts against the older reading:** `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted**
>   `.github/workflows/build-production.yml` with the subject *"the cms ECR repository is decommissioned
>   (M810)"*, and its body states that M810 *"deletes `module "cms_euwest1"` from the platform's
>   `services.tf`, which destroys the ECS service and the production-cms ECR repository"* — the workflow
>   went because it *"would try to push an image into a registry that no longer exists."*
>   **Whether that infrastructure-side deletion has actually been applied is NOT MEASURABLE from any clone
>   set we have** — `infrastructure` has never been in one. So do not assert either way: what is measured is
>   a `cms`-repo commit asserting the destruction, and a `cms`-repo terraform block that still declares the
>   module. The fenced map states this limit explicitly and is authoritative
>   ([`platform-migration-status.md`](../architecture/platform-migration-status.md), the `cms` row). This
>   bullet previously said *"`module.cms_euwest1` is still declared as the rollback path"* as a flat fact,
>   which the map already said it could not see. **`jobsimulation`'s ECS service is already
>   destroyed** — `6092c6d2` deleted the `module "jobsimulation"` block with its task definition and ECR
>   repository (`jobsimulation/terraform/main.tf:15-22`); the module file survives only to own the
>   LiveKit/Chime recording buckets `backend` reads by literal name, the `/production/jobsimulation/*` SSM
>   parameters and the atlas tracker (`:24-40`), and dropping the legacy `jobsimulation` schema is a separate,
>   still-pending M810 step. Fenced statement:
>   [`platform-migration-status.md`](../architecture/platform-migration-status.md).
>
> The skiller-specific detail below is the authoritative
> [**§ Skiller-in-app merge — fact-sheet**](#skiller-in-app-merge--fact-sheet-v21-quick-change).

## Role & Responsibility

`app` is the **main API gateway** of the platform — the service that frontends, hiring apps, and other backend services talk to first. It owns the `public` schema (users, organizations, memberships, assignments, subscriptions, payments) and, since the **skiller-in-app merge (July 2026)**, the **skills taxonomy domain** — the skills graph (**≥42,790 skills** across **≥22,470 job roles**; that is the measured *public* subset, `organization_id IS NULL`, 2026-06-29 — the long-quoted "60K skills / 18K roles" is not a measurement, and [18K is outright refuted](../architecture/shared_libraries.md#taxonomy-figures)), skill/job-role embeddings, and AI skill matching formerly owned by the standalone [skiller](./skiller.md) service. It exposes:

* **GraphQL Federation v2 subgraph** for high-level user / organization / assignment queries — plus the taxonomy types/queries absorbed from the former skiller subgraph (`graph/schemas/skiller_taxonomy.graphqls`)
* **Connect-RPC** for inter-service calls (**no external caller is left** — `messenger` was the last, and `838d907` removed its container) — the mux registers five handlers unconditionally (`main.go:1185-1228` @ `app` `b948604` v1.366.0): `UsersService` (`:1187`), `OrganizationsService` (`:1188`), `SkillerService` (`:1196`), `JobSimulationService` (`:1204`) and `lab.v1.LabSessionService` (`:1228`), plus **`CMSService` only when the Directus edge is configured** (`if cmsRPCServer != nil`, `:1212-1214`).

  **There is no `SkillPathSessionService`** — measured: **0** occurrences in Go source, and no `skillpath…v1connect` package is imported. Skill-path session state lives in `public.skill_path_sessions` and is reached through the GraphQL subgraph and in-process calls, not over RPC.

  > **⚠️ `app`'s OWN docs still list it** (`app/CLAUDE.md:109`, `app/knowledge/architecture.md:28` — re-derived at `app` **`ad9f3c49`**, which was `origin/main` on 2026-08-06; both anchors are unchanged from `2035f9a4`, the ref this corpus used to call "origin/main" and which is now 5 commits behind. The CLAUDE.md line was `:80` when this was first measured, so re-find the sentence rather than trusting the offset), which is where this corpus previously got the claim. That is Trap C in [`../ops/platform-alignment.md`](../ops/platform-alignment.md) — *the platform's planning docs lag its own code*. **Grade against `main.go`, not against `app/CLAUDE.md`.**
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
- **RPC re-pointed, then un-set** — the `SkillerService` Connect-RPC surface is served **by app itself**
  (`internal/rpc/skillerrpc/`). Consumers kept the env var `SKILLER_RPC_ADDR`, re-pointed at
  `http://backend:8083`. **That count was always ref-relative, and it has now reached zero:** four
  occurrences in `docker-compose.yml` @ platform `0808b92` (the ref this fact-sheet was first ground
  against — `backend`, `jobsimulation`, `cms` and `messenger` each carried one); **one** @ `0dab54d`,
  messenger's, after `d11a403` deleted the `jobsimulation` and `cms` blocks and dropped it from
  `backend`, which no longer addresses a surface it serves itself; and **none** @ `0c91421`, because
  `838d907` deleted the `messenger` block that held the last one. **No compose file sets any
  `*_RPC_ADDR` today.** **No terraform in the clone set names `http://backend.internal.anthropos:8081`,
  and this doc no longer asserts that any does.** Measured 2026-08-06 by two independent mechanisms:
  `git grep` at each clone's own HEAD over the **44** tracked `.tf` files in the 13-repo `stack-demo`
  clone set → **0 files**; a raw filesystem `find … -name '*.tf' | grep` over the same working trees,
  **59** files → **0** (positive control on `service_discovery_namespace_id`: 25 files). The literal's only
  occurrence anywhere in the clone set is a markdown KB page — `app/knowledge/service-dependencies.md:52`
  @ `ad9f3c49` — which is **not** terraform and is itself in the **past** tense: *"it used to reach the
  users, cms, jobsimulation and skiller surfaces at `http://backend.internal.anthropos:8081`, and folding
  it in at v9.0 closed that edge"*, under the heading *"**There are no external callers of app's RPC mux
  left.**"* **And the one tree that could settle the production value is not measurable from here:** the
  deciding declaration lives in the `infrastructure` repo, which is in no clone set — see
  [`platform-migration-status.md`](../architecture/platform-migration-status.md) for the fenced
  unmeasurable-claims convention.
  **What IS measurable is only the shape the address would be built from, and it is a derivation, not a
  reading:** the `app` service deploys as `local.project = "backend"` (`app/terraform/locals.tf:6` @
  `b948604f`, unchanged at `ad9f3c49`); it takes a Cloud Map service-discovery namespace **id** as an
  input variable (`app/terraform/main.tf:58` @ `b948604f`) — the namespace *name* `internal.anthropos`
  appears in **no `.tf` file in any clone**; and it exposes `local.rpc_port = 8081` (`locals.tf:8`, mapped
  at `main.tf:185-186`, both @ `b948604f`). Every cloned service module only *declares* its `*_rpc_addr`
  variable, with **no default** (`messenger/terraform/variables.tf:77,82,87,92` @ `fa47850d`;
  `app/terraform/variables.tf:197,230` @ `b948604f` — that is a pin, and it has since moved: at
  `ad9f3c49` `cms_rpc_address` is at `:309` and `storage_rpc_addr` is **deleted**, replaced by
  `storage_s3_bucket` / `storage_s3_public_bucket`). **The unqualified `backend:8081` form, which this
  corpus quoted in several places, is that same derived endpoint written short — not a second variable.**
  **Do not repoint any of these citations at `service-dependencies.md` to make the claim resolve:** a
  markdown page is not the production terraform, and a correctly-cited false statement is worse than a
  stale one.
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
  `skiller` Redis stream — but only as a **consumer**: nothing in `app` publishes to it. `main.go:1276`
  @ `b948604f` (the demo pin) is a direct `AddSubscriber`; at `app` **`ad9f3c49`** — `origin/main` on
  2026-08-06, 5 commits past the `2035f9a4` this corpus used to *label* "origin/main" — that
  registration has moved into the map-built subscriber set (`subs[d.Streams.Skiller]`,
  `subscriber_wiring.go:209`), applied by the single loop at `main.go:1579-1581`. **No `NewPublisher`
  names `SKILLER_STREAM` at any of those refs** — re-derived at `b948604f`, `2035f9a4` and `ad9f3c49`.
  The producer was the standalone skiller service and went with it. See the
  Redis Streams section below.
- **Clean-bring-up prerequisite:** the merged migrations create the taxonomy vector columns as
  `extensions.vector(1536)` and a GIN-trigram index via `extensions.gin_trgm_ops`, so the **`extensions`
  schema (pgvector + `pg_trgm`) must be bootstrapped before `make migrate`** on a clean DB — else app
  `20260518125439` and cms `20250116133510` fail with `schema "extensions" does not exist`. (Bring-up
  ordering, tracked for M211; not a merge defect.)

**Live de-risk (measured 2026-07-08):** a cold containerized `make up` on stack-dev built the 86-commit
merged image and brought up the federation with **no skiller container** — `SKILLER_RPC_ADDR` pointed at
`http://backend:8083`, which is what compose set at that ref and sets nowhere now
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
* **Profile**: `core` (the default), `backend`, `all` — `profiles: [core, backend, all]` (`docker-compose.yml:110`, derived from `docker-compose.yml` @ platform `0c91421`; it was `:100` at `0dab54d`, and compose clean-ups move it). The default profile is `core`, not `graphql`: `0dab54d` renamed it. Corrected M257x iter-68
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
* **Connect-RPC**: `rpc.go` is the top-level wire-up. Look there for the implemented services. **There is no external caller left, and no address to be one with.** `messenger` was the last, and it reached four surfaces — `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR`, `SKILLER_RPC_ADDR` — all pointed at `http://backend:8083` by `d11a403` (M809), all set on messenger's own compose block and nowhere else, under compose's own comment *"cms + jobsimulation are folded into app: all four RPC edges are the one backend mux"*. `838d907` deleted that block, so **compose now sets zero `*_RPC_ADDR` variables** and the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is `backend → sentinel` (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48` @ platform `0c91421`). **It is not the only cross-process edge:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The earlier two-of-four split (`http://cms:8091` / `http://jobsimulation:8401`, true at platform `2adcf71`) is history twice over. `app`'s own source comment still says *"additive + DORMANT: external callers (messenger) keep hitting the standalone cms via `CMS_RPC_ADDR` **until the M809 re-point**"* (`app/main.go:1205-1211` @ `b948604` v1.366.0) — **that comment is now stale in `app`**; grade the address against compose, not against the comment. **This bullet used to close by naming a production terraform address. That assertion is DROPPED, not softened.** Measured 2026-08-06 by two mechanisms — `git grep` at each clone's own HEAD over the 44 tracked `.tf` files in the 13-repo `stack-demo` clone set, and a raw filesystem grep over the 59 `.tf` files in the same working trees — **no terraform anywhere in the clone set names `http://backend.internal.anthropos:8081`**; the literal's sole occurrence is a past-tense markdown KB page, which is not terraform and must not be cited as if it were. **And the one tree that could settle the production value, `infrastructure`, is in no clone set** — see the *RPC re-pointed, then un-set* bullet in the fact-sheet above for the full derivation, and [`platform-migration-status.md`](../architecture/platform-migration-status.md) for the fenced unmeasurable-claims convention. Services include `lab.v1.LabSessionService`, `SkillerService` (`internal/rpc/skillerrpc/`), `JobSimulationService` and `CMSService`. Note the RPC server runs with a **60s write timeout** — the ported skiller RAG/LLM methods can exceed the old 10s default.
* **HTTP** (port 8082): Clerk webhooks, payment webhooks, document upload/convert endpoints, "Talk to Data" SSE.

### Upstream consumers

* Next Web App (GraphQL via **`backend`'s own endpoint, `:8082/graphql/query`** on a local stack since platform `2adcf71` deleted the router — the Cosmo Router is prod-only now — plus direct HTTP for SSE and webhooks)
* Hiring App
* Mobile App
* Studio-Desk (for org-level metadata)

### Downstream dependencies

* **Sentinel** — authz on every request
* **Object storage** — **in-process since the v9.0 fold** (2026-08-04), not a service hop: `app` constructs the private and public managers itself (`internalstorage.NewManager` / `NewPublicManager` at `app/main.go:524`, `:525`, re-derived at `app` **`ad9f3c49`** — `origin/main` on 2026-08-06; `main.go` is **byte-identical** to `2035f9a4`, the ref this corpus used to label "origin/main", so every `main.go` line number pinned to `2035f9a4` still resolves) and threads them to each consumer. `STORAGE_RPC_ADDR` has **0 read sites** — its 3 remaining occurrences are comments, one of which (`app/main.go:504`) says *"STORAGE_RPC_ADDR is gone"*. There is no `storage` compose service to address either, since `838d907`. See [Storage](./storage.md)
* **Directus** (`content.anthropos.work`) — the external content edge read by the in-process cms domain
* **Judge0** — sandboxed code execution, called directly (`JUDGE0_BASE_URL`) since roadrunner merged in
* **LiveKit / AWS Chime** — simulation voice + recording, for the in-process jobsim engine
* **Gotenberg** — Office → PDF conversion
* **PostgreSQL** (`public` schema), **Redis** (cache + streams)
* **External**: Clerk (auth), Stripe (payments), PostHog, Bedrock (AI), AI providers via the shared `ai` library (embeddings + skill matching — merged skiller domain), **Brevo** — reached in-process by both folded subsystems (transactional mail via messenger-in-app, marketing contacts via customerio-sync-in-app), each behind its own switch — and Sentry

### Redis Streams

* app is **both producer and consumer** of **four** of the five application streams — `backend`, `skillpath`, `jobsimulation`, `cms` — plus the `AI`/`ai_usage` usage stream: **five both-ways streams in all**, and with consumer-only `skiller` on top, **six subscribers against five publishers**. (Watch the partition when quoting a number from here: *four* is the application-stream subtotal, *five* the both-ways total, *six* the subscriber count.) **`skiller` is the exception: `app` only SUBSCRIBES to it and nothing publishes to it.** Enumerated over every publisher constructor in `app` @ `b948604f`, the topics are `backend` (`main.go:287`), `SKILLPATH_STREAM` (`:637`), `CMS_STREAM` (`:1039`), `AI_USAGE_STREAM` and `JOBSIMULATION_STREAM` (`internal/jobsimwiring/wiring.go:127`, `:180`); `SKILLER_STREAM` occurs once in Go **at that ref**, at `main.go:1276`, and it is an `AddSubscriber` call. At `app` `ad9f3c49` the same five publishers sit at `main.go:325`, `:746`, `:1149` and `wiring.go:132`, `:185`, while the subscriber side has been rebuilt as a map (`buildStreamSubscribers`, `subscriber_wiring.go:203-248`, applied by one loop at `main.go:1579-1581`) — so there is no standalone skiller `AddSubscriber` line to cite there. **Grade the shape at the ref you name.** The producer was the standalone skiller service, which is decommissioned — the fact was **deleted, not moved**
* The one external producer left is **Directus**, whose webhooks feed the `cms` stream
* Each stream has exactly **one** subscriber with multiple handlers merged via `.AddHandler(...)` — colony keys by stream name, so a second `AddSubscriber` for the same stream silently overwrites the first

## Local Development

### Run in Docker

```bash
cd platform
make up                # the default core profile (Makefile:10 PROFILE ?= core) — recommended
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

Versioned Atlas migrations live in `terraform/migrations/` (per `atlas.hcl`: `dir = "file://terraform/migrations"`, source `ent://internal/data/ent/schema`). **There is no top-level `migrations/` dir** — it held only `atlas.sum` and `6a46e8445` (2026-06-18, *"chore(migrations): remove obsolete atlas.sum file"*) deleted it; `git ls-tree b948604f migrations/` is empty, as it is at `2035f9a4` and at **`ad9f3c49`** (`origin/main` on 2026-08-06 — `2035f9a4` is a pin, not the tip, and is 5 commits behind). Generate a new migration after an Ent schema change with `make migrations` (`atlas migrate diff --env local`); apply with `atlas migrate apply --env local` (or `make migrate S=app`).

The `public` schema is the largest in the platform. **The migration set has NOT been static since May 2026.** This line used to date the head of the queue to May 2026 and describe it as simulation-type definitions plus content JSON defaults; that was **refuted** — it was two to three months stale, and neither the date nor the subject matter survives. Measured 2026-08-06 over `terraform/migrations/*.sql`: **170** migrations at `app` `ad9f3c49` (`origin/main` that day), **169** at the demo pin `b948604f`. **46** of them landed after 2026-05-31, so the last *May* migration (`20260529072659_add_lab_session.sql`) sits 46 back from the head, not at it. The head at `ad9f3c49` is `20260803143844_ai_readiness_recommendation_path.sql` (2026-08-03 — adds `ordinal` + `path_size` to `ai_readiness_recommendations` and re-keys its unique index); the head at `b948604f` is `20260731154527_academy_chapter_progress_completed_at.sql`. The three immediately preceding that one are `20260731131307.sql` (`course_builder_sessions.brief` + `credits_spent`), `20260729133514.sql` (the M709c `local_*` session-mirror collapse — the mirrors are re-pointed at `job_simulation_sessions` / `skill_path_sessions`, then **dropped**) and `20260728103254_ai_readiness_snapshot_frozen_matched_sources.sql`. **None of these five touches simulation-type definitions or content JSON defaults.** Re-derive the head of the queue rather than quoting a date from here: `git -C app ls-tree --name-only <ref> terraform/migrations/ | sort | tail -3`.

## Testing

```bash
go test ./...
# Heavy components have isolated test suites:
go test ./internal/askengine/...
```

## Related Documentation

* [AI Architecture](../architecture/ai_architecture.md) — Bedrock routing, cost tracking
* [CMS](./cms.md), [Jobsimulation](./jobsimulation.md) — the merged domains, now inside `app` (no compose service since `d11a403`)
* [Skiller](./skiller.md) — the former standalone skills-taxonomy service, merged into app (July 2026)
* [Skillpath](./skillpath.md) — the former standalone skill-path runtime engine, merged into app ("skillpath-in-app", M502→M507)
* [Storage](./storage.md), [Messenger](./messenger.md), [CustomerIO Sync](./customerio-sync.md) — the v9.0 "support-in-app" trio; no compose service since `838d907`
* [Gotenberg](./gotenberg.md) — PDF conversion sidecar
