# Anthropos Services Dependency Map

This document outlines the inter-service dependencies inferred from configuration files (`docker-compose.yml`) and code inspections.

## Dependency Matrix

Sourced from `platform/docker-compose.yml` `depends_on:` declarations and environment variables (`*_RPC_ADDR`).

> Since the monolith merge most of this matrix collapsed: `skiller`, `skillpath`, `roadrunner`, `jobsimulation` and `cms` are domains inside **Backend (`app`)**, so their edges are in-process calls, not dependencies.

| Service | Depends On (Direct) | Infrastructure |
| :--- | :--- | :--- |
| **Backend** (`app`) — the monolith | Sentinel, **cms**, Storage (compose `depends_on`, `docker-compose.yml:70-80` — yes, the monolith still has a startup edge onto the cms **husk**); Gotenberg (runtime HTTP, no startup-order dep) | Postgres (`public` schema; `pgvector` in `extensions` — skiller embeddings, skill-path sessions, the 23 jobsim run-state tables, the cms similarity/Studio tables), Redis, **Clerk**, **Directus**, **Judge0**, **LiveKit**, **AWS Chime**, **AI Providers** |
| **Sentinel** | - | Postgres |
| ~~**CMS**~~ | **Merged into `app`** ("cms-in-app v8.0", app v1.360.0) — the content layer + Studio run in-process; Directus stays external | *(**no container** at platform `0dab54d` — the husk is gone from compose and from `repos.yml`; **M809 has landed**, so `CMS_RPC_ADDR` resolves to `backend:8083`)* |
| ~~**Jobsimulation**~~ | **Merged into `app`** ("jobsim-in-app") — the session engine runs in-process; simulation definitions come from the cms domain by ID without an RPC hop | *(**no container** at platform `0dab54d` — gone from compose and from `repos.yml`)* |
| ~~**Skillpath**~~ | **Merged into `app`** ("skillpath-in-app", M502→M507) — the skill-path engine's dependencies (cms content by ID, the jobsimulation Redis Stream, Sentinel) are now `app`'s, in-process | *(no standalone service)* |
| ~~**Roadrunner**~~ | **Merged into `app`** with jobsim-in-app — `backend` calls Judge0 directly via `JUDGE0_BASE_URL` | *(**no container** at platform `0dab54d` — gone from compose; `ROADRUNNER_RPC_ADDR` is set nowhere. Prod terraform still reads `= 1`)* |
| **Storage** | Postgres, Redis — **`depends_on` only** (`docker-compose.yml:213-217`, both `service_healthy`) | **S3** only — plus a local-filesystem fallback per bucket. Storage **reads** neither: no `DB_CONNECTION`/`REDIS_ADDR` in its compose env, no redis in its `go.mod`, which is what its own doc says at [`storage.md:14,21`](../services/storage.md). The compose ordering edge and the runtime data path disagree here **by design**, which is why [`service_taxonomy.md`](service_taxonomy.md) lists postgresql + redis under storage. (At platform `0dab54d` the service moved to `profiles: [storage-legacy]`, so a default bring-up no longer starts it at all.) Corrected M257x iter-49, over-corrected to `-`, re-corrected iter-52 |
| **Gotenberg** | - | - (stateless conversion service) |
| **Messenger** (opt-in profile) | **Two of four addresses reach `backend`**, two still reach the husks — `BACKEND_USERS_RPC_ADDR` + `SKILLER_RPC_ADDR` = `http://backend:8083` (`docker-compose.yml:255,265`); `CMS_RPC_ADDR` = `http://cms:8091` (`:256`) and `JOBSIMULATION_RPC_ADDR` = `http://jobsimulation:8401` (`:258`) **until the M809 re-point** (`app/main.go:1196-1202`) | Postgres, Redis, **Brevo** (email delivery) |
| **CustomerIO Sync** (opt-in profile) | Postgres | **Customer.io** |
| ~~**Graphql (Cosmo Router)**~~ | **Not in a local stack** — platform `2adcf71` deleted the compose service and the `repos.yml` entry; the frontends call `backend` directly at `:8082/graphql/query`. Still declared in production terraform; repo archived 2026-07-30. Composed `backend` alone (1 subgraph) | *(no local service)* |
| **Studio-Desk** (opt-in profile) | `backend`'s GraphQL endpoint directly (`:8082/graphql/query`) — the router it used to depend on is gone locally. Compose `depends_on` is **`backend` + `cms`** (`docker-compose.yml:337-341`), not `graphql` | **Clerk**, **OpenAI / Azure OpenAI / Anthropic** (Copilot, via `AI_PROVIDER_CHAIN`) |
| **Studio-Room** | (runs inside the `app` container; depends on the backend process) | **OpenAI**, **Anthropic**, **Mistral** |

> **Skiller merged into app (July 2026):** the standalone skiller service is gone from the compose file. Its RPC surface is now served by **backend** — consumers keep the `SKILLER_RPC_ADDR` env var, re-pointed at `http://backend:8083` (production terraform: `skiller_rpc_addr = http://backend:8081`). See [Backend](../services/backend.md) and the [skiller stub](../services/skiller.md).
>
> **Skillpath merged into app (skillpath-in-app, M502→M507):** the standalone skillpath service is gone from the compose file / repos.yml / supergraph. Its skill-path progression engine now runs **in-process inside `app`**, with session state in `public.skill_path_sessions` (the legacy `skillpath` schema is an empty husk). See [Backend](../services/backend.md) and the [skillpath stub](../services/skillpath.md).
>
> **Jobsimulation + cms merged into app (jobsim-in-app, cms-in-app v8.0):** the last two subgraph services are gone from the **supergraph** — the federation now composes **one** subgraph. **And at platform `0dab54d` they are gone from compose and from `repos.yml` too** — 10 compose services, 6 repo entries, neither list containing cms or jobsimulation. (Until `2adcf71` both still started as unfederated husks; that is what changed.) Their tables were re-created in `public` (the legacy `jobsimulation` and `cms` schemas are non-authoritative). Their ECS modules are **still declared** in production terraform as the rollback path and take no traffic; teardown is **M810**. See [Jobsimulation](../services/jobsimulation.md) and [CMS](../services/cms.md).
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
| **taxonomy** | **node-id library** (not data): direct — app, messenger; indirect — storage, sentinel |

## Event Streams (Redis Streams via Watermill)

Services communicate asynchronously through named Redis Streams. Stream names come from `*_STREAM` env vars in `platform/docker-compose.yml`.

| Stream Name | Producer | Consumer(s) | Events |
| :--- | :--- | :--- | :--- |
| `backend` | App | App (cms **domain** in `app`; the `cms` husk also still subscribes until platform M810) | User/org updates |
| `skiller` | App | App | Skill score changes — both producer and consumer live inside app since the skiller→app merge (stream name retained) |
| `jobsimulation` | App | App (the jobsim engine + the skill-path engine, on ONE subscriber), Messenger (if running) | Session completed, insights generated |
| `cms` | App (+ **Directus webhooks** → `POST /api/webhook/directus`, now authenticated via `DIRECTUS_WEBHOOK_SECRET`) | App (the cms similarity/Studio handlers + the jobsim handlers, merged onto ONE subscriber) | Content published/updated, translation & clone requests |
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
`Frontend` -> `Backend` / `Jobsimulation`
*   The jobsimulation engine fetches the simulation **definition** (the `simulations` content/blueprint) from the cms domain by ID — in-process since cms-in-app. It owns no content, only the run/session state.
*   The jobsimulation engine stores its session/run **state** (interactions, recordings, validation results, anti-cheat) via `Storage` or directly to the **`public`** schema (the legacy `jobsimulation` schema is non-authoritative).
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
*   **Studio-Room** (Python, embedded inside the **`app`** container since cms-in-app) executes the generation pipeline against AI providers (OpenAI, Anthropic, Mistral).
*   Final content is persisted via the **cms domain** (in `app`); **Directus** is the underlying storage backend.

### 5. Skill Path Progress (Event-Driven)
`Jobsimulation` -> `Redis Stream` -> `App (skill-path engine, in-process)`
*   When a user completes a simulation, **Jobsimulation** publishes an event.
*   The **skill-path engine — now inside `app`** (merged from the standalone skillpath service, M502→M507) — subscribes to the Jobsimulation stream and updates step/chapter/path progress **state** (`SkillPathSession → ChapterSession → StepSession`, in `public.skill_path_sessions`) — it owns no content, only the per-user progression state.
*   The engine reads the skill-path **content** structure (chapters → steps it tracks against) from the **cms domain in-process** — it was a Connect-RPC call to the CMS service before the merge — and calls **Sentinel** for authorization, which is still a real network hop.

### 6. Document → PDF Conversion
`Backend (app)` → `Gotenberg`
*   The backend service uses Gotenberg's `/forms/libreoffice/convert` endpoint to render Office documents to PDF. See `app/internal/converter/gotenberg.go`.
*   `GOTENBERG_URL=http://gotenberg:3200` is injected via the backend's compose env.
