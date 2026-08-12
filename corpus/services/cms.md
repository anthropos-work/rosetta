# CMS Service

> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of **cms-in-app v8.0** (`app` **v1.360.0**, July 2026), the standalone `cms` Go microservice has been
> **merged into the `app` monolith** (the service the platform calls "backend"). CMS no longer runs as a
> separate service **in production**. Its subgraph is gone from the supergraph, and its ECS service is
> **DESTROYED — corrected M257x iter-127.** **RESOLVED at M257x iter-123/127 — the cms ECS service is DESTROYED.** `infrastructure` @ `13c248e6` declares **no `module "cms"` at all**, and `infrastructure/terraform/production/services.tf:64-70` records what the apply destroyed (ECS service, task definition, ECR repository, IAM roles, security group, Cloud Map entry, log group, alarms, the ten `/production/cms/*` SSM parameters). **`cms/terraform/main.tf:39` is ORPHANED DEAD CODE** — a `service_desired_count` in a module no root module instantiates describes nothing ([`org-repos.md` § 3](../architecture/org-repos.md)). The legacy **schema** drop is a separate, still-pending M810 step. **This banner read *"scaled to zero, not deleted … the one M810 row whose terraform module block has not moved"* until iter-127**, four days after the measurement that settled it.
> **⚠️ But cms HAS taken an M810 step since, and the corpus's "it has not moved" was becoming stale:**
> `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** `.github/workflows/build-production.yml` under the
> subject *"the cms ECR repository is decommissioned (M810)"*, its body stating that M810 *"deletes
> `module "cms_euwest1"` from the platform's `services.tf`, which destroys the ECS service and the
> production-cms ECR repository"* — the workflow went because it *"would try to push an image into a
> registry that no longer exists."* **The two facts in this repo LOOKED like they pointed opposite ways**
> (a module block that still declares the service; a CI commit asserting the registry is already gone),
> and the deciding declaration lands in `infrastructure`. **That was a CLONE-SET limit, not a measurement
> limit, and it is SETTLED: the ECS service is DESTROYED.** `infrastructure` was read at `13c248e6` (M257x
> iter-123, re-verified at iter-132 — `git ls-remote` puts that sha at origin `HEAD`): there is **no
> `module "cms"` declaration anywhere in it**, and `infrastructure/terraform/production/services.tf:64-70`
> records what the apply destroyed. So the CI commit was the correct signal and `cms/terraform/main.tf:39`
> `service_desired_count = 0` is **orphaned dead code** — no root module instantiates that file
> ([`org-repos.md` § 3](../architecture/org-repos.md)). **The pending M810 step is the legacy *schema*
> drop, not the service.** This paragraph read *"which has never been in any clone set we have — do not
> assert either way"* for four iterations after the read that settled it; **`infrastructure` is indeed not
> in the standing clone set, and that never entailed unmeasurable.** It was the **fourth** engine consolidated into `app`, after
> [skiller](./skiller.md), [skillpath](./skillpath.md) and [jobsimulation](./jobsimulation.md) — **not the
> last.** The v9.0 program (2026-08-04) then folded [`storage`](./storage.md),
> [`messenger`](./messenger.md) and [`customerio-sync`](./customerio-sync.md), and platform `838d907`
> (merged `0c91421`, 2026-08-05) deleted all three containers the next day. See the fenced map,
> [`platform-migration-status.md`](../architecture/platform-migration-status.md).
>
> **✅ The husk is GONE locally, and M809 has landed (re-measured at platform `0c91421`).**
> There is no `cms` compose service, no `cms` entry in `repos.yml` (4 entries: app, sentinel,
> next-web-app, studio-desk) and no `cms` profile. Nor is there a `CMS_RPC_ADDR` any more: M809
> re-pointed it at `http://backend:8083` on the `messenger` block — **one of the MIDDLE TWO `d11a403`
> moved** (with `JOBSIMULATION_RPC_ADDR`); `BACKEND_USERS_RPC_ADDR` and `SKILLER_RPC_ADDR` already held
> that value at `d11a403^`, so "all four" (which
> [`service_taxonomy.md`](../architecture/service_taxonomy.md) asserted until M257x iter-115) was never
> true — and `838d907` deleted that block —
> **compose now sets zero `*_RPC_ADDR` values**, and the cms domain is reached in-process.
> *(Until `2adcf71` all of the above was false, and this banner said so; the M809 re-point is what
> changed it.)* **Scope note: this is the LOCAL compose topology only.** Whether production's
> `module.cms_euwest1` rollback path has also been torn down (**M810**) is now MEASURED and the answer is
> **yes** — `infrastructure` `13c248e6` (2026-08-07) declares no `module "cms_euwest1"` at all. Still do not
> read the local removal *as* the production one; read the production one from
> `infrastructure/terraform/production/services.tf:64-70`, which is where it is recorded. See
> [`platform-migration-status.md`](../architecture/platform-migration-status.md) and
> [`org-repos.md` § 3](../architecture/org-repos.md).
>
> **The Directus content edge stays external.** The merge moved the *Go service*; the authored content still
> lives in Directus at `content.anthropos.work`, which `app` reads over HTTP.
>
> Where everything went:
>
> * **Domain** — `app/internal/cms/` (directus, similarity, studio, library, importer/exporter, aivideo,
>   contentread, jobsimimport, rpcsrv, worker, …), wired from `app/internal/cms/wiring.go`.
> * **Data** — `similarities`, `similarity_categories`, `similarity_features`, `similarity_skills`,
>   `studio_documents`, `studio_tasks` were re-created in the **`public` schema** by
>   `app/terraform/migrations/20260724132049_cms_data_model.sql`, with the **same table names**. The old `cms`
>   DB schema is **legacy — no longer authoritative**.
> * **RPC** — `CMSService` is served on `app`'s single RPC mux, and **nothing outside the process
>   reaches it**. `messenger` was the last caller; `CMS_RPC_ADDR` was `http://backend:8083` at
>   `0dab54d` — `d11a403` put it there, **one of the two that commit moved, not one of four** — set on
>   messenger's block alone, and `838d907` deleted that block — so the variable is
>   set by no compose file today.
>   **The M809 re-point had already landed** and there was no husk container left to reach either — `cms` is
>   not among the **five** services compose declares at platform `0c91421` (**seven** effective, once
>   `include: common.yml` adds the `postgresql`/`redis` floor). (`http://cms:8091`, still quoted around this
>   corpus, was true at `2adcf71`.) **The production address is not stated here in either direction:** no `.tf` file in any clone names `http://backend.internal.anthropos:8081` (0 hits over all 44 tracked `.tf` files in the 13 `stack-demo` repos at each clone's own HEAD, 2026-08-06; the literal's 6 non-terraform occurrences are counted once in [`backend.md`](./backend.md), the load-bearing one being the past-tense `app/knowledge/service-dependencies.md:52` @ `app` `ad9f3c49`), **and the deciding declaration lives in `infrastructure` — not in the standing clone set, and read anyway.** Production DOES name it, exactly once: `infrastructure/terraform/production/services.tf:346` @ `13c248e6` sets `cms_rpc_address` to `http://backend.internal.anthropos:8081` **as an input to `module "backend_euwest1"`** — so **M809 landed in production too, in the same shape as locally** (M257x iter-132). Derived once in [`backend.md`](./backend.md)'s *RPC re-pointed, then un-set* bullet and not restated here. See also the *RPC* line under *Interface Discovery* below.
>   `app`'s own source comment at `app/main.go:1205-1211` (@ `b948604` v1.366.0) still calls the in-app
>   edge *"additive + DORMANT … until the **M809** re-point"* — **that comment is stale in `app`**, and
>   there is no compose value left to grade it against. `app` itself makes **no** outbound cms RPC.
> * **GraphQL** — the cms subgraph was folded into `app`'s `backend` subgraph. That single commit,
>   `graphql-wundergraph@915da06` (2026-07-29), deleted **both** `schemas/cms.graphqls` **and**
>   `schemas/jobsimulation.graphqls`, taking the supergraph from **3 subgraphs to 1** — not 2 to 1. (The
>   jobsimulation subgraph outlived jobsim-in-app and was removed here; the ladder before it was 5 → 4 at
>   `749dc86` → 3 at `7c17e63`.) Public (unauthenticated) library content queries are preserved — see app
>   v1.360.2/v1.360.3.
> * **Events** — `app` owns the `CMS_STREAM` subscriber. The folded similarity re-index + Studio handlers are
>   merged onto app's **existing** CMS subscriber via `.AddHandler(...)`; they act on disjoint rows, so they
>   compose. Directus webhooks land on `POST /api/webhook/directus`, which now **fails closed** without
>   `DIRECTUS_WEBHOOK_SECRET` (the standalone webhook was unauthenticated).
> * **Caching** — the Directus item cache still lives in its own Redis DB, `REDIS_CMS_CACHE_INDEX` (5).
> * **Studio** — the Python `anthropos-studio-room` project is now pulled into the **`app`** image via the CI
>   `additional_repo` mechanism (app v1.360.1), the same way `cms` used to do it.
> * **Infrastructure** — **the ECS service is DESTROYED (iter-123, propagated iter-127).** This repo's own
>   `cms/terraform/main.tf:39` still reads `service_desired_count = 0` in an otherwise-whole 191-line
>   module, **and that describes nothing**: no root module instantiates it. *"That is all that can be said
>   from here"* was true only while `infrastructure` was in no clone set — a **clone-set limit, not a
>   measurement limit**, and the fix was to clone the repo. Jobsimulation's ECS service was destroyed outright at
>   `6092c6d2`, and generalising that to `cms` is exactly the mistake
>   [`platform-migration-status.md`](../architecture/platform-migration-status.md) fences.
>   **⚠️ Do NOT extend "the module block has not moved" to "the rollback path is intact".** `6efa1d5`
>   (merged `f38c0c4`, 2026-08-04) deleted this repo's build-production workflow under the subject *"the cms
>   ECR repository is decommissioned (M810)"*, because it *"would try to push an image into a registry that
>   no longer exists"* — so the two measured facts in this repo appeared to point opposite ways.
>   **They never did, and iter-123 settled it by cloning the repo this bullet said could not be read:**
>   `infrastructure` `13c248e6` declares **no `module "cms_euwest1"`**, and
>   `terraform/production/services.tf:64-70` records what its deletion destroyed (ECS service, task
>   definition, ECR repository, IAM roles, security group, Cloud Map entry, log group, alarms, ten
>   `/production/cms/*` SSM parameters), with a `removed { destroy = false }` for the Atlas tracker at
>   `:88-94` and the legacy **schema** deliberately untouched (`:85-86` — that drop is a separate,
>   still-pending M810 step). **The CI commit was the correct signal; `cms/terraform/main.tf:39` is
>   ORPHANED DEAD CODE**, an input to a module no root module instantiates. The general rule and the
>   three sibling repos it also settles: [`org-repos.md` § 3](../architecture/org-repos.md).
> * **Repo** — the `cms` git repo still exists but is **frozen/legacy**; make changes in `app`.
>
> For current documentation of this domain, see [Backend (`app`)](./backend.md).

## Role & Responsibility

The CMS service is the **content layer of the platform** — it owns the authored, versioned, published **CONTENT / DEFINITIONS** and serves them to everyone else. It does three things:

1. **Serves content** to the rest of the platform via GraphQL Federation and internal RPC — **skill paths** (title, description, cover/video, curators a.k.a. "Meet the Experts", library categories, **chapters → steps**, the job-simulation steps inside a chapter, skills-to-verify, settings, versioning — the `skill_paths` Directus collection), **job-simulation blueprints** (the `simulations` collection + `sequences`, roles, tasks, validation criteria), and the **content library** (`library_categories`, `library_macro_categories`, `resource`) — all proxied through Directus with Anthropos-specific business logic on top.
2. **Owns the Studio data model** — `StudioDocument` (**a customer-UPLOADED attachment converted to
   Markdown for AI context** — `storage_document_id`, `name`, `content_type`, `markdown`, `tokens`),
   `StudioTask` (generation jobs), and related entities for the content-authoring workflow.
   ⚠️ *`StudioDocument` was described as "(simulation blueprints)" until run 81 and that is FALSE* —
   blueprints live in Directus (`simulations` + `sequences`), exactly as the bullet above states.
   `app/internal/data/ent/schema/studio_document.go:9-11` @ `ad9f3c498`: *"caches a user-uploaded
   attachment converted to Markdown for AI context in the Studio generation pipeline."* **This is a
   data-classification error in the doc that also declares `cms.studio_*` 100 % customer data** — the
   table holds customer uploads and their extracted full text, which is a *higher* sensitivity than
   platform-authored blueprints, and the schema notes it carries no Ent privacy policy.
3. **Runs the AI generation pipeline** in-process. The Python project `anthropos-studio-room` is pulled into the image and dispatched as a subprocess **in argv (exec) form — never through a shell** (`app/internal/cms/studio/studioManager.go:1099-1101` @ `app` `ad9f3c49`; `git grep -n '"bash"' ad9f3c49 -- '*.go'` over the whole tree returns **0**); the Go side dispatches generation work, the Python code executes it against **OpenAI, Azure OpenAI or Anthropic** — those three and no others. The provider registry is a three-entry dict: `{'openai': OpenAIProvider, 'azure': AzureProvider, 'anthropic': AnthropicProvider}` (`services/ai.py:705-708` @ `anthropos-studio-room` `aeec036` v0.51.1), and `services/ai.py:1-2` imports only `openai`/`anthropic`. **There is no Mistral path in the Python *generation* engine** — but `mistralai` is **not** unimported. `tools/pdf2md.py:24` does `from mistralai import Mistral` (client at `:96`, `model="mistral-ocr-latest"` at `:127`): a **standalone CLI OCR utility**, one leg of the `tools/r3.py` offline PDF→markdown chain, that nothing on the generation path calls — `gen.py` never imports `tools`, nothing outside `tools/` references it, and no Go caller exists (Go execs **two** studio scripts and neither is `pdf2md.py`: `studio/gen.py` at `studioManager.go:119` and `studio/postgen.py` at `:1045`, both @ `app b948604f`). `git -C app/studio grep -i mistral aeec036a` returns **22 hits in 3 files** (`requirements.txt:8`, `tools/pdf2md.py`, `tools/r3.py`), not one. So Mistral is **OCR-only on both sides** — Go-side for studio attachments, Python-side for that offline tool — and on the generation path on neither (see the Downstream-dependencies bullet below, and [`studio-room.md`](./studio-room.md) for the grep caveat that hid `tools/`).

This last point was the first structural shift: **studio-room is not a standalone deployable**. Since cms-in-app it rides in the **`app`** image rather than the cms one.

> [!IMPORTANT]
> **CMS owns content; the runtime engines own state.** Do not conflate the **skill-path engine** with skill-path content, or the **`jobsimulation`** service with simulation content. Those are **runtime/session engines** that hold *no* content and reference CMS artifacts **by ID**:
> - **The [skill-path engine](./skillpath.md)** (merged into `app` — "skillpath-in-app", M502→M507; formerly the standalone `skillpath` service) tracks per-user progression *state* (`SkillPathSession → ChapterSession → StepSession`, progress %); it reads the skill-path *structure* it tracks against from the **cms domain in-process** — `app/internal/skillpath/session.go:205-207` (`// cms-in-app deseam: cms is in-process`) calls `contentread.CmsContentReader.GetSkillPathDomain`. It was a `CMS_RPC_ADDR` Connect-RPC hop until both merged into `app`.
> - **[`jobsimulation`](./jobsimulation.md)** runs the interactive simulation *session*; it reads the simulation *definition* it runs from the cms domain **in-process** (it was a `cms.GetSimulation` Connect-RPC hop until both merged into `app`) — it has no `DIRECTUS_BASE_ADDR` of its own, so all its content reads go *through* CMS.
>
> So **content = CMS/Directus; the like-named service = the state machine over that content.** This split is the source of a recurring naming confusion — see the [Service Taxonomy](../architecture/service_taxonomy.md) and [Architecture Overview](../architecture/architecture_overview.md) content-vs-runtime callouts.

> **Demo/dev set-dressing (v1.2 → v1.5 "prop room"):** the **public** content templates (the `directus` schema of the prod app DB — `private = false AND tenant_id IS NULL AND status = 'published'`) are captured read-only by the snapshot mechanism, then served from a **per-stack Directus**. The collection-schema gap that once forced live-prod reads is **closed**: M21 captures + auto-provisions the content-model structure (DDL + serve rows), M22 boots a per-stack Directus as a compose service (offset port, torn down with the stack), and **M23 re-points `DIRECTUS_BASE_ADDR` at that local instance — for `cms` AND for `backend`** (⚠️ the `cms` re-point alone is **not sufficient**: since cms-in-app, `backend` is the in-process Directus reader via `app/cms_reader_switch.go`, so a stack that re-points only `cms` still reads prod. M257x iter-24 measured that as 96 all-403 Directus lines in `backend`'s log; `DIRECTUS_DATA_CONSUMERS` now names both) (`http://directus:8055`, the in-network service, #M23-D1) so a `--local-content` stack (demo default; dev opt-in) serves its **own** captured catalog — no live-prod read. The **asset plane stays on prod**: `DIRECTUS_PUBLIC_BASE_ADDR` keeps pointing at `content.anthropos.work`, so browser images load real `<...>/assets/<uuid>` URLs (the data-plane-local / asset-plane-prod split; the captured `directus_files` refs resolve those uuids). A **non-`--local-content`** stack still reads the public content **live from prod** (a demo does so **anonymously**, the prod token stripped — the documented prod-read fallback) — see [`corpus/ops/snapshot-spec.md`](../ops/snapshot-spec.md) (the M10 content surface + the M23 cutover). The app-Postgres `cms.studio_*` tables (`StudioDocument` / `StudioTask`) are **100% customer data** and are never captured (the tenant firewall).

## Architecture & Code Map

* **Codebase**: `cms` — repo `git@github.com:anthropos-work/cms.git`. **Not cloned by `make init`**: no `repos.yml` entry since `d11a403`. Clone it by hand to read the pre-merge source; the live code is `app/internal/cms/`
* **Language**: Go 1.26 (primary — `cms/go.mod:3` `go 1.26.4`) + Python 3.11 (studio-room)
* **Database**: ~~PostgreSQL `cms` schema~~ — **`public`, via `app`'s Ent**. The `cms` schema is a legacy husk since cms-in-app v8.0; the similarity + Studio tables moved to `public`
* **Ports**: **8080 (GraphQL/HTTP), 8081 (Connect-RPC) — the binary's own defaults**, and now the only ones there are: `cms/cmd/root.go:77` `cmp.Or(os.Getenv("PORT"), "8080")` / `:78` `cmp.Or(os.Getenv("RPC_PORT"), "8081")`. The **8090 / 8091** pair quoted throughout this corpus was **compose-supplied by a service that no longer exists**: `docker-compose.yml` set `PORT=8090` (`:169`) / `RPC_PORT=8091` (`:173`) and published `8090:8090` / `8091:8091` (`:154-155`) — **at `2adcf71`**. At `0dab54d` there is no `cms` service, so nothing sets them and nothing is published; **8090/8091 are historical, not an address you can reach.** The domain's live surface is `backend`'s (`:8082/graphql/query`, RPC on `:8083`)
* **Docker image**: Two-stage build — Go binary built in `golang:1.26-bookworm` (`cms/Dockerfile:2`), copied into a `python:3.11-slim` final stage (`:23`) along with `cms/studio/` and its `pip install -r studio/requirements.txt`. The Go binary is the entrypoint; it shells out to Python when a generation task fires.

### Key directories

```
cmd/                       Service entrypoints
internal/
  graph/                   GraphQL layer (gqlgen)
    schemas/*.graphqls     API contract — simulation.graphqls, skills.graphqls, studio.graphqls
    *.resolvers.go         Hand-written resolvers
    model/models_gen.go    Auto-generated (DO NOT EDIT)
  directus/                Directus client + collection queries
  rpcsrv/                  Connect-RPC server (binds RPC_PORT — see § Ports; 8091 was compose-supplied)
  auth/                    Authn middleware
  event/                   Watermill event handling
  worker/                  Background workers (Redis Streams consumers)
  studio/                  Studio data-model business logic (StudioManager)
  library/                 Content library
  exporter/                Content export
  importer/                Content import
  aivideo/                 AI video processing (HeyGen integration)
  similarity/              Similarity/matching algorithms
ent/                       Ent schema + generated code
studio/                    Python AI generation pipeline (cloned via `make init-studio`)
  gen.py                   Pipeline entrypoint
  postgen.py               Post-generation steps
  agents/                  Agent definitions
  configs/                 Per-environment AI model slots (`{env}_config.ini`)
  services/                Provider wrappers (ai.py, …)
  knowledge/, tools/       Pipeline knowledge + helper tooling
  requirements.txt         openai, anthropic, rich, pyyaml, requests, jinja2, mistralai, pytest,
                           pytest-asyncio — the file verbatim, 9 packages. **`python-docx` is NOT among
                           them** and never was: it was listed here until v2.8 M257x, and neither
                           `cms/studio/requirements.txt` nor the in-image `app/studio/requirements.txt`
                           contains it. (The only `.docx` in the tree is a filename in an authoring
                           guideline and an asset-example README — no dependency.)
terraform/                 IaC
```

> Note: local proto development requires the developer to create their own (uncommitted) `go.work` linking `../proto`; it is not committed to the repo.

## Studio Generation Pipeline

The Studio entities and the Python pipeline are tightly coupled. The flow:

```mermaid
sequenceDiagram
    participant Desk as Studio-Desk
    participant CMS as CMS (Go)
    participant DB as PostgreSQL
    participant Studio as studio/gen.py (Python — argv exec, no shell)
    participant AI as AI Providers

    Desk->>CMS: createStudioDocument(blueprint)
    CMS->>DB: INSERT studio_documents
    Desk->>CMS: generateContent(documentId)
    CMS->>DB: INSERT studio_tasks (pending)
    CMS->>Studio: runCommand(python3, ["studio/gen.py", "--media", ...]) — argv, never a shell
    Studio->>AI: prompts (FAST → STRICT → EXECUTION → CREATIVE slots)
    AI-->>Studio: generated content
    Studio->>CMS: results (stdout / files)
    CMS->>DB: UPDATE studio_tasks (completed) + persist content
```

### Studio entities

* **StudioDocument** (`ent/schema/studioDocument.go`): the blueprint a Studio-Desk user authored
* **StudioTask** (`ent/schema/studioTask.go`): a generation job — status, progress, params

## Directus integration

CMS acts as a proxy + business-logic layer over Directus:

```
Frontend / Studio-Desk → CMS GraphQL → Business Logic → Redis Cache → Directus API → PostgreSQL
```

Why this pattern: business rules and validation live in CMS, caching reduces Directus load, and the abstraction makes it easier to swap the storage backend later.

## Interface Discovery

* **GraphQL**: since cms-in-app the schemas live with the rest of app's at `app/internal/web/backend/graphql/graph/schemas/*.graphqls`, served on the `backend` subgraph. The Directus webhook receiver moved to `POST /api/webhook/directus` on app's web server and **fails closed** without `DIRECTUS_WEBHOOK_SECRET` (the standalone receiver at `:8090/webhooks/` was unauthenticated).
* **RPC**: `app/internal/cms/rpcsrv` — served on app's single RPC mux, and **every caller is in-process**. `messenger` was the last external one: M809 pointed its `CMS_RPC_ADDR` at `http://backend:8083` — **`d11a403` moved exactly two variables on that block, `CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`; the other two already read `http://backend:8083` at `d11a403^` and were untouched (measured M257x iter-115)** — and `838d907` deleted the messenger service and that variable with it, so compose sets it nowhere. **No `.tf` file in any clone names `http://backend.internal.anthropos:8081`** — 0 hits measured 2026-08-06 over all 44 tracked `.tf` files in the 13 `stack-demo` repos at each clone's own HEAD, and 0 again over the 59 `.tf` files a raw filesystem sweep of that workspace finds. The literal does occur in the clone set — **6 times, none of them terraform**; the count and its per-repo derivation are stated once, in [`backend.md`](./backend.md)'s *RPC re-pointed, then un-set* bullet, and are not restated here. The one that matters is a **markdown KB page** — `app/knowledge/service-dependencies.md:52` @ `app` `ad9f3c49` — which is not terraform, and which puts it in the **past** tense: *"it used to reach the users, cms, jobsimulation and skiller surfaces at `http://backend.internal.anthropos:8081`, and folding it in at v9.0 closed that edge"*, under the heading *"**There are no external callers of app's RPC mux left.**"* **The production declaration is not in this repo — it is in `infrastructure`, which is not in the standing clone set and HAS been read: production names the literal exactly once, as `module "backend_euwest1"`'s `cms_rpc_address` input** (`infrastructure/terraform/production/services.tf:346` @ `13c248e6`, M257x iter-132). Stated once in [`backend.md`](./backend.md)'s *RPC re-pointed, then un-set* bullet; not restated here.
* **Federation**: **there is none left to speak of.** The cms subgraph was folded into `backend` at cms-in-app v8.0 — the **3 → 1** step, because `graphql-wundergraph@915da06` deleted `cms.graphqls` and `jobsimulation.graphqls` in the same commit. Then platform `2adcf71` (2026-07-31, PR #23 *"drop the WunderGraph router"*) **deleted the Cosmo/WunderGraph router itself** — service, `repos.yml` entry and clone. So this line's old ending, *"Cosmo Router now composes `backend` alone"*, names a component that no longer exists: **nothing composes anything.** GraphQL is served **directly by `backend`** at `:8082/graphql/query` — note the path moved with it (`/graphql` → `/graphql/query`), so a host-only re-point 404s rather than errors.

### Upstream consumers
* Next Web App (GraphQL)
* Studio-Desk (GraphQL for studio entities)
* Backend (`app`) — the skill-path engine, the jobsimulation engine and the cms domain all run **in the same
  process**, so those hops are plain function calls, **not RPC**; the Redis Streams edge has `app` on both
  ends. (**messenger** was the one remaining out-of-process consumer; M809 re-pointed its `CMS_RPC_ADDR`
  at `backend`, and `838d907` removed the messenger container and the variable. There is no husk `cms`
  container left to receive RPC either.)

### Downstream dependencies
* Directus (content storage)
* PostgreSQL (Ent ORM, **`public` schema** — the cms tables were re-created there at cms-in-app v8.0; the
  legacy `cms` schema is non-authoritative. Consistent with the **Data** bullet of the *Where everything
  went* list in the banner at the top of this doc — **named, not pinned:** it carried a line range until
  M257x iter-120, by then the **Domain** bullet, one bullet off)
* Redis (cache, Watermill streams)
* AI providers — **OpenAI / Azure OpenAI / Anthropic** for the `studio/` Python generation pipeline
  (`services/ai.py:705-708`). **Mistral is NOT one of them**: every use of it is **OCR**, never generation. The Go one —
  `app/internal/cms/studio/markdownManager.go:10` imports `internal/cms/studio/mistralocr` and `:30`
  builds the client (`mistralocr.New(aiKey)` inside `NewMarkdownManager`, re-derived at `app`
  `2035f9a` — a **pin**, not a moving label: that was origin/main on 2026-08-05, and both offsets still name the same two constructs at today's origin/main, `ad9f3c49` (re-checked 2026-08-06); it was `:11`/`:19` and a `mistral.NewMistral(nil, MISTRAL_API_KEY)` call before the key-plumbing
  fix that stopped it reading `os.Getenv` behind the caller's back).
  Its single use is `OCRProcess` (document → markdown) on the studio attachment path
  (`studioManager.go:531` *"supported ocr content types for mistral ocr"*, `:583`, and `xlsx.go:13` — xlsx is
  rendered locally precisely because Mistral OCR rejects it). Nothing generates through it. There is also a
  **Python-side** Mistral OCR user in the same image — `app/studio/tools/pdf2md.py:24`
  `from mistralai import Mistral` (`mistral-ocr-latest`), a standalone CLI **no Go caller and no `gen.py`
  path dispatches** — `tools/r3.py:139`/`:190`/`:199-206` DOES exec it as step 2 of the offline chain, so
  the flat *"nothing dispatches it"* this line carried is withdrawn
  (`git -C app/studio grep -i mistral aeec036a` → 22 hits / 3 files; `git -C app grep -- studio/`
  returns 0 because `studio/` is untracked in `app`, `app/.gitignore:79`)

## Local Development

### First-time setup

> **⚠️ HISTORICAL — `cd cms; make init-studio` is NOT the onboarding path any more.** Since cms-in-app v8.0
> the studio-room pipeline is pulled into the **`app`** image by CI via the `additional_repo` mechanism (app
> v1.360.1) — see the **Studio** bullet of the *Where everything went* list in the banner at the top of this
> doc (**named, not pinned:** it carried a line range until M257x iter-120, by then the **Events** bullet). Work on this domain in **`app`**, not in the
> frozen `cms` repo. The block below is kept only because the legacy repo still carries these targets.

> **⚠️ The demo tooling used to ENTER this repo to fetch Studio. It does not any more — FIXED at M257x
> iter-270, and this block is the retraction.** Until then `demo-stack/ensure-clones.sh` opened its
> studio-consumer list with a **hardcoded** `_studio_repos="cms"`, deriving only the rest from `repos.yml`
> (which has not listed `cms` since `d11a403`) — so the decommissioned repo was the *preferred* fetcher, and
> on any box still carrying a populated `stack-demo/cms/studio` it was the branch that actually ran.
> At rext **`e64a3cd3b`** the hardcode, the `make init-studio` special case, and the preference are all
> **gone**: `demo-stack/ensure-clones.sh:314` derives the set by calling `studio_consumer_names`
> (`stack-core/lib/studio.sh:121`) against the platform clone's `repos.yml`, and **refuses the
> bring-up** if it cannot be derived, and
> acquisition is a plain `git clone` for every consumer — `init-studio` was literally that same clone, so
> the special case bought nothing and cost a hardcoded corpse.
>
> **Two things iter-268 got wrong, and the second is the one worth carrying.** *"Nothing is broken by
> it"* was true only of the branch iter-268 looked at. iter-270 graded all **8** platform-topology
> derivations in the demo bring-up path and found this one **failed OPEN**: an unreadable `repos.yml`
> collapsed the consumer set to **`cms` alone**, dropping every live consumer and re-arming the
> `/build/studio: not found` failure the phase exists to pre-empt. **Neither arm errored, which is why it
> survived four releases** — a preference does not fail; on a fresh box the guard skips it, on a stale box
> it silently wins, and both look like success.
>
> **The structural lesson stands and is now paid**: the clone *set* was correctly fenced —
> `clone_pin_guard.py` removed five phantom pin keys at iter-222, `cms` among them — but **the
> studio-consumer list was a SECOND registry one file over, and that sweep did not reach it**
> (`platform-alignment.md` §5's *"a named-consumer list survives the merge that moved the consumer"*,
> occurring inside the repo that wrote the rule down). `FIX-M257x-268-ensure-clones-hardcodes-cms-as-studio-fetcher`
> is **CLOSED**; this block asserted it open until M257x iter-278.
>
> **What is still true:** `stack-demo/` on a long-lived box carries **6** clones `repos.yml` does not name
> (`cms`, `graphql-wundergraph`, `jobsimulation`, `messenger`, `roadrunner`, `storage`), and
> `stack-demo/cms/studio` stays populated where it already was. They carry **no compose service and no
> build context**, and nothing fetches through them any more. That they remain on disk is a **measured
> and accepted** state, not backlog: `ROUTE-M257x-265-stack-demo-carries-six-dead-clones` was **closed
> at M257x iter-268**, whose deliverable was the census itself — *nothing is deleted in this iter* was
> one of its sealed pre-registrations, because a stale clone is the evidence.

**⚠️ Read the tense.** The Python studio tree had to be cloned **before** any docker build, or `make up`
failed with `"/studio": not found` — and **that is still true today, of `app`, not of `cms`.** The
requirement did not die with the service; `fdb8034a` moved it, and `app/Dockerfile:45-46` hard-COPYs
`/build/studio` on every `make up`. The past tense below is scoped to the **`cms` targets**, never to the
dependency. The live procedure is [`setup_guide.md` § Acquire the Studio
runtime](../ops/setup_guide.md#acquire-the-studio-runtime--required-before-make-up-or-the-backend-build-fails);
three troubleshooting entries told operators to *delete* the COPY lines until M257x iter-265 corrected them
(`D-M257x-265-1`). The `cms`-side history:

```bash
cd cms
make init-studio   # HISTORICAL — clones anthropos-studio-room into cms/studio/
make setup         # installs ent, atlas, gqlgen
make gen           # regenerates GraphQL resolvers + Ent code
```

### Run in Docker (with the rest of the platform)

```bash
cd platform
make up                  # the `core` profile — `backend` (app) serves the cms domain in-process
# There is NO cms profile and no cms container. Asking for one does NOT fail: it exits 0
# and starts postgresql and redis — a stack with no application in it.
```

### Run natively (single service)

```bash
cd platform
make dev S=backend       # stops the backend container
cd ../app
go run .                 # the cms domain runs inside this process
```

For Python pipeline development — **in `app/studio/`, not `cms/studio/`.** `cms` has no `repos.yml` entry, so
`make init` does not clone it; the pipeline that actually ships rides in the `app` image (`additional_repo`,
app v1.360.1) and the two `requirements.txt` are byte-identical. The `cms/studio/` path below is the
historical one.

```bash
cd app/studio          # was: cms/studio
pip install -r requirements.txt
# the repo's own entry point (studio/CLAUDE.md:12-14)
python gen.py --media simulation --prompt "..." --evaluation_skills "skill1, skill2" --branch stable
# or, from a reusable blueprint JSON in the attachments directory
python gen.py --media simulation --blueprint <file>.json
```

> **⚠️ There is no `--template` flag** — and a stray one is **silently swallowed**, not rejected.
> `gen.py:484-492` registers exactly nine arguments (`-i/--interactive`, `-m/--media`, `-f/--force`,
> `--simid`, `--branch`, `--prompt`, `--annotations`, `--pipeline`, `--blueprint`), and
> `parse_argument` (`gen.py:18-28`) calls `parse_known_args` and merges the leftovers into the args
> dict. So `--template foo` parses cleanly, sets a key **nothing in the codebase reads**, and the
> command *succeeds* while generating something unrelated. The reusable unit is a **blueprint**, not
> a template — see [studio-room.md](./studio-room.md#blueprints-not-templates).

> ⚠️ **CORRECTED M257x iter-115 — this note asserted the exact inversion of a shipped security property.**
> It said the Go service invokes `python3 studio/gen.py ...` / `studio/postgen.py` **"via `bash -c`"**, in the
> present tense and with no HISTORICAL marker (unlike both of its neighbours in this section). The live code is
> the opposite, deliberately: at `app` `ad9f3c49`,
> `app/internal/cms/studio/studioManager.go:1096-1098` reads *"runCommand executes name+args in **argv (exec)
> form — NEVER through a shell**… nothing is string-interpolated into a command line (M809b H-1/M-1)"*, and
> `:1101` is `pycmd := exec.CommandContext(ctx, name, args...)`. `:100-103` says it in the caller's own words —
> *"It MUST NOT be interpolated into a shell … **No `bash -c`**"* — and `:119` is
> `s.runCommand(ctx, pyBin, append([]string{"studio/gen.py"}, tokens...))`. Measured: `git grep -n '"bash"'
> ad9f3c49 -- '*.go'` over the whole `app` tree returns **0**.
>
> **What survives.** Dev mode does auto-provision a venv at `studio/studio-venv` and run
> `pip3 install -r studio/requirements.txt` — as **fixed argv** (`:126`, `:129`), *"previously chained into the
> same `bash -c` string that carried the tainted args"* (`:122-124`). Paths are still `studio/...`, not from
> inside `studio/`. For standalone Python work, use a venv to match the service's behavior.
>
> **Why it read as true:** the claim is still correct about the **frozen** `cms` repo
> (`ca50c817:internal/studio/studioManager.go:967` = `exec.Command("bash", "-c", command)`) — right about the
> dead code, wrong about the shipped code, and wrong about the direction of a deliberate hardening.

### Sync the studio submodule

**HISTORICAL — this is no longer how the shipped pipeline is refreshed.** Since cms-in-app v8.0 CI pulls
`anthropos-studio-room` into the **`app`** image via `additional_repo`; there is no manual sync step, and the
`cms` repo is not cloned by `make init`. Kept because the frozen repo still carries the target:

```bash
cd cms
make update-studio       # cd studio && git pull
```

## Testing

```bash
go test ./...            # Go tests
cd studio && pytest      # Python tests (requires `pip install -r requirements.txt`)
```

## Related Documentation

* [Skillpath](./skillpath.md) — the skill-path runtime engine (merged into `app`, M502→M507) that tracks progress against CMS-owned skill-path content (the content-vs-runtime split)
* [Jobsimulation](./jobsimulation.md) — the runtime service that *runs* simulations defined as CMS content
* [AI Architecture](../architecture/ai_architecture.md) — model routing, generation slots
* [Service Taxonomy](../architecture/service_taxonomy.md) — orchestration profile + the content-vs-runtime callout
* [Dependency Map](../architecture/dependency_map.md) — RPC and event-stream relationships
