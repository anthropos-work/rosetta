# Anthropos Services Dependency Map

This document outlines the inter-service dependencies inferred from configuration files (`docker-compose.yaml`) and code inspections.

## Dependency Matrix

Sourced from `platform/docker-compose.yml` `depends_on:` declarations and environment variables (`*_RPC_ADDR`).

> Since the monolith merge most of this matrix collapsed: `skiller`, `skillpath`, `roadrunner`, `jobsimulation` and `cms` are domains inside **Backend (`app`)**, so their edges are in-process calls, not dependencies.

| Service | Depends On (Direct) | Infrastructure |
| :--- | :--- | :--- |
| **Backend** (`app`) — the monolith | Sentinel, Storage (compose `depends_on`); Gotenberg (runtime HTTP, no startup-order dep) | Postgres (`public` schema; `pgvector` in `extensions` — skiller embeddings, skill-path sessions, the 23 jobsim run-state tables, the cms similarity/Studio tables), Redis, **Clerk**, **Directus**, **Judge0**, **LiveKit**, **AWS Chime**, **AI Providers** |
| **Sentinel** | - | Postgres |
| ~~**CMS**~~ | **Merged into `app`** ("cms-in-app v8.0", app v1.360.0) — the content layer + Studio run in-process; Directus stays external | *(no standalone service)* |
| ~~**Jobsimulation**~~ | **Merged into `app`** ("jobsim-in-app") — the session engine runs in-process; simulation definitions come from the cms domain by ID without an RPC hop | *(no standalone service)* |
| ~~**Skillpath**~~ | **Merged into `app`** ("skillpath-in-app", M502→M507) — the skill-path engine's dependencies (cms content by ID, the jobsimulation Redis Stream, Sentinel) are now `app`'s, in-process | *(no standalone service)* |
| ~~**Roadrunner**~~ | **Merged into `app`** with jobsim-in-app — `backend` calls Judge0 directly via `JUDGE0_BASE_URL` | *(no standalone service)* |
| **Storage** | - | Postgres, Redis, **S3** |
| **Gotenberg** | - | - (stateless conversion service) |
| **Messenger** (opt-in profile) | Backend (users, cms, jobsimulation and skiller RPC all at `http://backend:8083`) | Postgres, Redis, **Brevo** (email delivery) |
| **CustomerIO Sync** (opt-in profile) | Postgres | **Customer.io** |
| **Graphql (Cosmo Router)** | Backend (the **only** subgraph), Storage | - |
| **Studio-Desk** (opt-in profile) | Graphql | **Clerk**, **OpenAI / Azure OpenAI / Anthropic** (Copilot, via `AI_PROVIDER_CHAIN`) |
| **Studio-Room** | (runs inside the `app` container; depends on the backend process) | **OpenAI**, **Anthropic**, **Mistral** |

> **Skiller merged into app (July 2026):** the standalone skiller service is gone from the compose file. Its RPC surface is now served by **backend** — consumers keep the `SKILLER_RPC_ADDR` env var, re-pointed at `http://backend:8083` (production terraform: `skiller_rpc_addr = http://backend:8081`). See [Backend](../services/backend.md) and the [skiller stub](../services/skiller.md).
>
> **Skillpath merged into app (skillpath-in-app, M502→M507):** the standalone skillpath service is gone from the compose file / repos.yml / supergraph. Its skill-path progression engine now runs **in-process inside `app`**, with session state in `public.skill_path_sessions` (the legacy `skillpath` schema is an empty husk). See [Backend](../services/backend.md) and the [skillpath stub](../services/skillpath.md).
>
> **Jobsimulation + cms merged into app (jobsim-in-app, cms-in-app v8.0):** the last two subgraph services are gone from the compose file / repos.yml / supergraph — the federation now composes **one** subgraph. Their tables were re-created in `public` (the legacy `jobsimulation` and `cms` schemas are non-authoritative). Their ECS modules are **still declared** in production terraform as the rollback path and take no traffic; teardown is **M810**. See [Jobsimulation](../services/jobsimulation.md) and [CMS](../services/cms.md).
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
| `backend` | App | CMS | User/org updates |
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
`Frontend` -> `CMS` -> `Directus`
*   CMS acts as the gateway to Directus content.

### 4. Studio Content Creation
`Studio Desk` → `CMS` → (in-process) `Studio Room`
*   **Studio-Desk** (TypeScript) creates blueprints, sent to CMS as `StudioDocument` rows.
*   **CMS** (Go) creates `StudioTask` records and dispatches generation work.
*   **Studio-Room** (Python, embedded inside the **`app`** container since cms-in-app) executes the generation pipeline against AI providers (OpenAI, Anthropic, Mistral).
*   Final content is persisted via the CMS service; **Directus** is the underlying storage backend.

### 5. Skill Path Progress (Event-Driven)
`Jobsimulation` -> `Redis Stream` -> `App (skill-path engine, in-process)`
*   When a user completes a simulation, **Jobsimulation** publishes an event.
*   The **skill-path engine — now inside `app`** (merged from the standalone skillpath service, M502→M507) — subscribes to the Jobsimulation stream and updates step/chapter/path progress **state** (`SkillPathSession → ChapterSession → StepSession`, in `public.skill_path_sessions`) — it owns no content, only the per-user progression state.
*   The engine queries **CMS** (RPC) for the skill-path **content** structure (chapters → steps it tracks against) and **Sentinel** for authorization. All of this now runs in-process inside `app`.

### 6. Document → PDF Conversion
`Backend (app)` → `Gotenberg`
*   The backend service uses Gotenberg's `/forms/libreoffice/convert` endpoint to render Office documents to PDF. See `app/internal/converter/gotenberg.go`.
*   `GOTENBERG_URL=http://gotenberg:3200` is injected via the backend's compose env.
