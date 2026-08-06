# Anthropos Services Dependency Map

This document outlines the inter-service dependencies inferred from configuration files (`docker-compose.yml`) and code inspections.

## Dependency Matrix

Sourced from `platform/docker-compose.yml` `depends_on:` declarations and service-address environment variables — of which, at platform `0c91421`, there is exactly **one** left, `AUTHORIZATION_ADDRESS` (`docker-compose.yml:48`), and **zero `*_RPC_ADDR` variables** in any compose file: the `messenger` block that set the last four went with the service at `838d907`.

> Since the monolith merge most of this matrix collapsed: `skiller`, `skillpath`, `roadrunner`, `jobsimulation`, `cms`, `storage`, `messenger` and `customerio-sync` are all domains inside **Backend (`app`)**, so their edges are in-process calls, not dependencies. The rows below that are struck through have **no compose service at all**.

| Service | Depends On (Direct) | Infrastructure |
| :--- | :--- | :--- |
| **Backend** (`app`) — the monolith | Sentinel, Redis, Postgres — the whole of its compose `depends_on` at platform `0c91421` (`docker-compose.yml:101-109`). **The cms, storage, messenger and customerio-sync edges are gone as *facts*, not merely moved:** `d11a403` deleted the cms container and `838d907` deleted the other three outright — compose says so in-line where those edges used to be (`:102-103`, *"storage, messenger and customerio-sync are not services any more — this one container serves all three in-process."*). Gotenberg is a runtime HTTP call with no startup-order dep | Postgres (`public` schema; `pgvector` in `extensions` — skiller embeddings, skill-path sessions, the 23 jobsim run-state tables, the cms similarity/Studio tables), Redis, **Clerk**, **Directus**, **Judge0**, **LiveKit**, **AWS Chime**, **AI Providers** |
| **Sentinel** | - | Postgres |
| ~~**CMS**~~ | **Merged into `app`** ("cms-in-app v8.0", app v1.360.0) — the content layer + Studio run in-process; Directus stays external | *(**no container** at platform `0c91421` — the husk is gone from compose and from `repos.yml`. **M809 landed** at `d11a403`, which re-pointed `messenger`'s `CMS_RPC_ADDR` at `backend:8083`; `838d907` then deleted `messenger` too, so **no compose file sets `CMS_RPC_ADDR` at all**)* |
| ~~**Jobsimulation**~~ | **Merged into `app`** ("jobsim-in-app") — the session engine runs in-process; simulation definitions come from the cms domain by ID without an RPC hop | *(**no container** at platform `0dab54d` — gone from compose and from `repos.yml`)* |
| ~~**Skillpath**~~ | **Merged into `app`** ("skillpath-in-app", M502→M507) — the skill-path engine's dependencies (cms content by ID, the jobsimulation Redis Stream, Sentinel) are now `app`'s, in-process | *(no standalone service)* |
| ~~**Roadrunner**~~ | **Merged into `app`** with jobsim-in-app — `backend` calls Judge0 directly via `JUDGE0_BASE_URL` | *(**no container** at platform `0dab54d` — gone from compose; `ROADRUNNER_RPC_ADDR` is set nowhere. Prod terraform still reads `= 1`)* |
| ~~**Storage**~~ | **Merged into `app`** (v9.0 "support-in-app") — `backend` serves object storage in-process; platform `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service **and** the `repos.yml` entry | *(**no container** at platform `0c91421`.)* Its data path was **S3** only — plus a local-filesystem fallback per bucket. While the service existed it declared `depends_on: postgres, redis` (both `service_healthy`) and **read neither**: no `DB_CONNECTION`/`REDIS_ADDR` in its compose env, no redis in its `go.mod`, which is what its own doc says at [`storage.md:40,47`](../services/storage.md). The compose ordering edge and the runtime data path disagreed **by design**, which is why [`service_taxonomy.md`](service_taxonomy.md) used to list postgresql + redis under storage. Corrected M257x iter-49, over-corrected to `-`, re-corrected iter-52; the row went historical when the service was deleted |
| **Gotenberg** | - | - (stateless conversion service) |
| ~~**Messenger**~~ | **Merged into `app`** (v9.0 "support-in-app") — `838d907` deleted the compose service and the `repos.yml` entry. While it ran, **all four of its addresses reached `backend`** — `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` and `SKILLER_RPC_ADDR`, all `http://backend:8083` in its compose block at `0dab54d`, under compose's own comment *"cms + jobsimulation are folded into app: all four RPC edges are the one backend mux"*. (No `file:line` here on purpose — the block is deleted, so every line number in it now points at somebody else's service.) **The M809 re-point landed** at `d11a403`; the two-of-four split was true at `2adcf71` and is history. Deleting the service deleted all four — **compose now sets zero `*_RPC_ADDR` variables** | Postgres, Redis, **Brevo** (email delivery) — now `app`'s, and gated **OFF** on a developer machine behind `MESSENGER_ENABLED` |
| ~~**CustomerIO Sync**~~ | **Merged into `app`** — it runs on `backend`'s asynq scheduler; `838d907` deleted the compose service. It was still in the **`all`** profile until then, so `make up-all` started a second Brevo contact pusher alongside `backend`'s own | **Customer.io** — now `app`'s, gated **OFF** on a developer machine behind `CUSTOMERIO_SYNC_ENABLED` |
| ~~**Graphql (Cosmo Router)**~~ | **Not in a local stack** — platform `2adcf71` deleted the compose service and the `repos.yml` entry; the frontends call `backend` directly at `:8082/graphql/query`. Still declared in production terraform; repo archived 2026-07-30. Composed `backend` alone (1 subgraph) | *(no local service)* |
| **Studio-Desk** (opt-in profile) | `backend`'s GraphQL endpoint directly (`:8082/graphql/query`) — the router it used to depend on is gone locally. Compose `depends_on` is **`backend`, and only `backend`** (`docker-compose.yml:138-140`) — the cms edge went with the cms container at `d11a403`, so this is now a one-edge block, not a two-edge one | **Clerk**, **OpenAI / Azure OpenAI / Anthropic** (Copilot, via `AI_PROVIDER_CHAIN`) |
| **Studio-Room** | (runs inside the `app` container; depends on the backend process) | **OpenAI**, **Azure OpenAI**, **Anthropic** — **not Mistral** on the generation path (`services/ai.py:705-708` is the whole provider registry). The `mistralai` requirement is **not** dead code, though: `tools/pdf2md.py:24` imports it for a standalone CLI OCR utility (`mistral-ocr-latest`) that nothing in the pipeline dispatches (`git -C app/studio grep -i mistral aeec036a` → 22 hits / 3 files) |

> **Skiller merged into app (July 2026):** the standalone skiller service is gone from the compose file. Its RPC surface is now served by **backend** — consumers keep the `SKILLER_RPC_ADDR` env var, re-pointed at `http://backend:8083`; but `messenger` was the last thing that set it locally, and since `838d907` deleted that service **no compose file sets it at all** (production terraform: `skiller_rpc_addr = http://backend.internal.anthropos:8081` — the app mux's Cloud Map name; see [Backend](../services/backend.md) for the derivation). See [Backend](../services/backend.md) and the [skiller stub](../services/skiller.md).
>
> **Skillpath merged into app (skillpath-in-app, M502→M507):** the standalone skillpath service is gone from the compose file / repos.yml / supergraph. Its skill-path progression engine now runs **in-process inside `app`**, with session state in `public.skill_path_sessions` (the legacy `skillpath` schema is an empty husk). See [Backend](../services/backend.md) and the [skillpath stub](../services/skillpath.md).
>
> **Jobsimulation + cms merged into app (jobsim-in-app, cms-in-app v8.0):** the last two subgraph services are gone from the **supergraph** — the federation now composes **one** subgraph. **And at platform `0c91421` they are gone from compose and from `repos.yml` too** — 7 compose services in the effective topology (5 declared in `docker-compose.yml`), 4 repo entries, neither list containing cms or jobsimulation. (It was 10 and 6 at `0dab54d`, before `838d907` dropped the last three support containers.) (Both still started as unfederated husks right up to `d11a403` — at `2adcf71` compose still declared all three of cms, jobsimulation and roadrunner — and `d11a403` is what changed it. `2adcf71` is the *router* deletion; do not conflate the two commits.) Their tables were re-created in `public` (the legacy `jobsimulation` and `cms` schemas are non-authoritative). **Their production dispositions have since diverged — do not state them as one:** `cms`'s ECS module is still declared **in its own repo** and takes no traffic (`cms/terraform/main.tf:39` `service_desired_count = 0`) — but **do not read that as "the prod rollback path stands"**: `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** cms's build-production workflow saying *"the cms ECR repository is decommissioned (M810)"*, so the repo holds two measured facts pointing opposite ways, and the deletion itself lands in `infrastructure`, **which has never been in any clone set** — UNMEASURABLE, report both and assert neither. While **jobsimulation's ECS service is destroyed — M810 has landed for that row** (`6092c6d2` deleted the `module "jobsimulation"` block along with its task definition and ECR repository; the file survives owning only the LiveKit/Chime buckets, the `/production/jobsimulation/*` SSM parameters and the atlas tracker — `jobsimulation/terraform/main.tf:15-40`, the legacy-schema drop being a separate, still-pending M810 step). See [Jobsimulation](../services/jobsimulation.md) and [CMS](../services/cms.md).
>
> **Content-vs-runtime dependency (unchanged, now in-process):** both the skill-path engine and the jobsimulation engine depend on the **cms domain for content/definitions** — cms is the content layer; they are runtime/session engines that hold no content and reference cms artifacts **by ID**. The skill-path engine fetches a path's chapter/step structure when (re)building a session; the jobsimulation engine loads a simulation's definition before running it. Both calls used to be Connect-RPC (`CMS_RPC_ADDR`, `cms.GetSimulation`); they are **plain function calls** now. The jobsim domain still holds no `DIRECTUS_BASE_ADDR` of its own — its Directus reads flow *through* the cms domain. (See [CMS](../services/cms.md), [Skillpath](../services/skillpath.md), [Jobsimulation](../services/jobsimulation.md).)

Production-only:
| Service | Depends On (Direct) | Infrastructure |
| :--- | :--- | :--- |
| **db-backup** | - | Postgres, **S3**, **Azure**, **Hetzner** |

### Shared Libraries

Imported as private Go modules (not deployed, **not** cloned by `make init`). Full reference: [Shared Libraries](./shared_libraries.md).

| Library | Used By |
| :--- | :--- |
| **colony** | All Go services (logging, DB, Redis, middleware, pub/sub); also bundles `authn` |
| **proto** | All Go services using RPC (contract definitions) + domain types |
| **ai** | app — i.e. every folded domain (AI provider wrapper — Go services only, not Studio-Desk). Cost & routing live in the consumers, not the lib |
| **authn** | Imported via `colony/authn` by app (standalone `authn` repo is legacy; the former cms/jobsimulation/skillpath usage is folded into app) |
| **taxonomy** | **node-id library** (not data): of the Go repos a stack clones at platform `0c91421` — direct: app (`app/go.mod:20` @ `b948604f`); indirect: sentinel (`sentinel/go.mod:21` @ `88bc5592`). The messenger (direct) / storage (indirect) requirements are frozen — `838d907` dropped both from `repos.yml` |

## Event Streams (Redis Streams via Watermill)

Services communicate asynchronously through named Redis Streams. Stream names come from `*_STREAM` env vars in `platform/docker-compose.yml`.

| Stream Name | Producer | Consumer(s) | Events |
| :--- | :--- | :--- | :--- |
| `backend` | App | App (cms **domain** in `app`) — **and two more; the old "and nothing else / no second subscriber locally" here was refuted twice.** (a) the standalone `messenger` service subscribed to this same stream — `messenger/internal/flow/flow.go:72` is a literal `AddSubscriber("backend", …)` over 21 live handlers — though since `838d907` deleted that compose service and its profile there is no longer any way to start it, so this subscriber now exists only in the frozen repo. (b) Since the v9.0 messenger-in-app fold (`app` `9d00a313`), **`app` itself runs a second subscriber server** on messenger's *own* Redis consumer group and attaches it to this stream by name — `app/main.go:1442` guards the takeover against `StreamBackend`, which is the literal `"backend"` (`app/internal/messenger/flow/streams.go:65`). That one is not opt-in: it is the stock `core` selection. What `d11a403` removed was only the `cms` husk container, and that was never the whole answer | User/org updates |
| `skiller` | **none** — the producer was the standalone skiller service, decommissioned at the merge; **the fact was deleted, not moved** | App (consumer only) | `SkillerCustomJobRoleCreated` → migrate an org's members to a new custom job role. **`app` runs a live subscriber on a stream nothing publishes to.** Enumerated over every publisher constructor in `app` @ **`b948604f` only** — `backend` (`main.go:287`), `SKILLPATH_STREAM` (`:637`), `CMS_STREAM` (`:1039`), `AI_USAGE_STREAM` + `JOBSIMULATION_STREAM` (`internal/jobsimwiring/wiring.go:127`, `:180`) — `SKILLER_STREAM` is not among them; its one Go occurrence, `main.go:1276`, is an `AddSubscriber` (handler: `internal/roles/roles.go:791`). **One ref, deliberately.** This cell also named origin/main `2035f9a4` until M257x iter-98, where **not one of those line numbers resolves** (`main.go:1276` is `apiKeyManager,`) and `SKILLER_STREAM` has **6** Go occurrences across 4 files, not one — so the sentence was false at the second of the two refs it cited, and a block naming two refs is `ambiguous` to the citation resolver besides. **The consumer-only finding itself holds at both**: `2035f9a4` still has no `NewPublisher` naming `SKILLER_STREAM`, and `internal/roles/roles.go:791` is `SkillerSubscriber()` at each. Compose still sets the name (`docker-compose.yml:71` @ `0c91421`). This row's Events cell previously named skill-score changes — that was never this stream's payload |
| `jobsimulation` | App | App (the jobsim engine + the skill-path engine, on ONE subscriber), the standalone Messenger (`messenger/internal/flow/flow.go:105` — no longer startable since `838d907`), **and app's messenger-in-app subscriber**: the same takeover as the `backend` row attaches to `StreamJobSimulation` too | Session completed, insights generated |
| `cms` | App (+ **Directus webhooks** → `POST /api/webhook/directus`, now authenticated via `DIRECTUS_WEBHOOK_SECRET`) | App (the cms similarity/Studio handlers + the jobsim handlers, merged onto ONE subscriber), the standalone Messenger (`messenger/internal/flow/flow.go:109` — no longer startable since `838d907`), **and app's messenger-in-app subscriber**: the takeover names all three of messenger's streams, this one included | Content published/updated, translation & clone requests |
| `skillpath` | App | App | Session updated, chapters completed — both producer and consumer live inside app since the skillpath→app merge (stream name retained) |
| ~~`roadrunner`~~ | — | **no producer or consumer** — roadrunner is merged into app, which calls Judge0 synchronously | ~~`RoadrunnerSubmissionCompleted`~~ |
| `AI` | (multiple) | (multiple) | AI usage / cost telemetry — see `AI_USAGE_STREAM=AI` env var |

> **Note**: The `chronos` stream was previously used by Chronos for timer events but is gone with the chronos service removal. Jobsimulation no longer has chronos as a dependency.

## Key Flows

### 1. User Authentication
`Frontend` -> `Backend` -> `Sentinel`
*   The Backend validates requests using Sentinel.
*   **Studio Desk** authenticates directly via **Clerk**.

### 2. Job Simulation
`Frontend` -> `Backend` (`app` — the **jobsimulation domain**, in-process; there is no jobsimulation service to reach)
*   The jobsimulation engine fetches the simulation **definition** (the `simulations` content/blueprint) from the cms domain by ID — in-process since cms-in-app. It owns no content, only the run/session state.
*   The jobsimulation engine stores its session/run **state** (interactions, recordings, validation results, anti-cheat) via the **storage domain inside `app`** (in-process since v9.0 — not the standalone `storage` service, which `838d907` deleted from compose altogether) or directly to the **`public`** schema (the legacy `jobsimulation` schema is non-authoritative).
*   Voice flows go through LiveKit; video recordings via AWS Chime SDK.

### 3. Content Delivery
`Frontend` -> `app` (cms domain) -> `Directus`   *(was `Frontend -> CMS -> Directus` before cms-in-app)*
*   The cms **domain inside `app`** acts as the gateway to Directus content — `Frontend -> app (cms domain) -> Directus`. It was a separate `CMS` service until cms-in-app v8.0; the hop is in-process now.

### 4. Studio Content Creation
`Studio Desk` → `app` (cms domain) → (in-process) `Studio Room`   *(was `Studio Desk → CMS → Studio Room` before cms-in-app)*
*   **Studio-Desk** (TypeScript) creates blueprints, sent to the **cms domain inside `app`** as `StudioDocument` rows
    over `backend`'s GraphQL endpoint (`:8082/graphql/query`).
*   The **cms domain** (Go, `app/internal/cms/`) creates `StudioTask` records and dispatches generation work —
    an in-process hop since cms-in-app v8.0, not a service call. Consistent with :9/:15/:31 above.
*   **Studio-Room** (Python, embedded inside the **`app`** container since cms-in-app) executes the generation pipeline against AI providers (OpenAI, Azure OpenAI, Anthropic — **not** Mistral; see the table above).
*   Final content is persisted via the **cms domain** (in `app`); **Directus** is the underlying storage backend.

### 5. Skill Path Progress (Event-Driven)
`App (jobsimulation domain)` -> `Redis Stream` -> `App (skill-path engine, in-process)`
*   When a user completes a simulation, the **jobsimulation domain inside `app`** publishes an event — both ends of this stream have been the same process since jobsim-in-app; the stream survives because the *name* was retained, not because a second service does.
*   The **skill-path engine — now inside `app`** (merged from the standalone skillpath service, M502→M507) — subscribes to the Jobsimulation stream and updates step/chapter/path progress **state** (`SkillPathSession → ChapterSession → StepSession`, in `public.skill_path_sessions`) — it owns no content, only the per-user progression state.
*   The engine reads the skill-path **content** structure (chapters → steps it tracks against) from the **cms domain in-process** — it was a Connect-RPC call to the CMS service before the merge — and calls **Sentinel** for authorization, which is still a real network hop.

### 6. Document → PDF Conversion
`Backend (app)` → `Gotenberg`
*   The backend service uses Gotenberg's `/forms/libreoffice/convert` endpoint to render Office documents to PDF. See `app/internal/converter/gotenberg.go`.
*   `GOTENBERG_URL=http://gotenberg:3200` is injected via the backend's compose env.
