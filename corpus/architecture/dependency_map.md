# Anthropos Services Dependency Map

This document outlines the inter-service dependencies inferred from configuration files (`docker-compose.yaml`) and code inspections.

## Dependency Matrix

Sourced from `platform/docker-compose.yml` `depends_on:` declarations and environment variables (`*_RPC_ADDR`).

| Service | Depends On (Direct) | Infrastructure |
| :--- | :--- | :--- |
| **Backend** (`app`) | Sentinel, CMS, Skillpath, Storage (compose `depends_on`); Gotenberg (runtime HTTP, no startup-order dep) | Postgres (with `pgvector` in `extensions` schema — embeddings of the merged skiller domain), Redis, **Clerk**, **AI Providers** (embeddings + skill matching) |
| **CMS** | Sentinel, Storage; Backend (skiller RPC surface via `SKILLER_RPC_ADDR=http://backend:8083`) | Postgres, Redis, **Directus**, **AI Providers** (Anthropic, OpenAI, Mistral — via embedded studio-room) |
| **Sentinel** | - | Postgres |
| **Jobsimulation** | Sentinel, Backend (user context + the skiller RPC surface since the merge), CMS (simulation *definitions* by ID via `cms.GetSimulation` RPC), Roadrunner, Storage | Postgres, Redis, **LiveKit**, **AWS Chime**, **AI Providers** |
| **Skillpath** | Sentinel, CMS (skill-path *content* structure by ID via `CMS_RPC_ADDR`), Jobsimulation (RPC + Redis Stream) | Postgres, Redis |
| **Storage** | - | Postgres, Redis, **S3** |
| **Roadrunner** | - | Redis, **Judge0** (code execution) |
| **Gotenberg** | - | - (stateless conversion service) |
| **Messenger** (opt-in profile) | Backend, CMS, Jobsimulation, Skillpath | Postgres, Redis, **Brevo** (email delivery) |
| **CustomerIO Sync** (opt-in profile) | Postgres | **Customer.io** |
| **Graphql (Cosmo Router)** | Backend, Jobsimulation, CMS, Skillpath, Storage | - |
| **Studio-Desk** (opt-in profile) | Graphql, CMS | **Clerk**, **OpenAI / Azure OpenAI / Anthropic** (Copilot, via `AI_PROVIDER_CHAIN`) |
| **Studio-Room** | (runs inside CMS container; depends on CMS process) | **OpenAI**, **Anthropic**, **Mistral** |

> **Skiller merged into app (July 2026):** the standalone skiller service is gone from the compose file. Its RPC surface is now served by **backend** — consumers keep the `SKILLER_RPC_ADDR` env var, re-pointed at `http://backend:8083` (production terraform: `skiller_rpc_addr = http://backend:8081`). See [Backend](../services/backend.md) and the [skiller stub](../services/skiller.md).
>
> **Content-vs-runtime dependency:** both `skillpath` and `jobsimulation` depend on **CMS for content/definitions** — CMS is the content layer; they are runtime/session engines that hold no content and reference CMS artifacts **by ID**. `skillpath` calls CMS over Connect-RPC (`CMS_RPC_ADDR=http://cms:8091`) to fetch a skill path's chapter/step structure when (re)building a session; `jobsimulation` calls CMS over Connect-RPC (`cms.GetSimulation`) to load a simulation's definition before running it. Note `jobsimulation` does **not** hold its own `DIRECTUS_BASE_ADDR` — all its Directus reads flow *through* CMS. (See [CMS](../services/cms.md), [Skillpath](../services/skillpath.md), [Jobsimulation](../services/jobsimulation.md).)

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
| **ai** | app, cms, jobsimulation (AI provider wrapper — Go services only, not Studio-Desk). Cost & routing live in the consumers, not the lib |
| **authn** | Imported via `colony/authn` by app, cms, jobsimulation, skillpath (standalone `authn` repo is legacy) |
| **taxonomy** | **node-id library** (not data): direct — app, cms, jobsimulation, messenger; indirect — skillpath, storage, sentinel |

## Event Streams (Redis Streams via Watermill)

Services communicate asynchronously through named Redis Streams. Stream names come from `*_STREAM` env vars in `platform/docker-compose.yml`.

| Stream Name | Producer | Consumer(s) | Events |
| :--- | :--- | :--- | :--- |
| `backend` | App | CMS | User/org updates |
| `skiller` | App | App | Skill score changes — both producer and consumer live inside app since the skiller→app merge (stream name retained) |
| `jobsimulation` | Jobsimulation | Skillpath, App, Messenger (if running) | Session completed, insights generated |
| `cms` | CMS | Jobsimulation, Skillpath, Backend | Content published |
| `skillpath` | Skillpath | App | Session updated, chapters completed |
| `roadrunner` | Roadrunner | Jobsimulation | `RoadrunnerSubmissionCompleted` (code execution finished; carries the Judge0 token) |
| `AI` | (multiple) | (multiple) | AI usage / cost telemetry — see `AI_USAGE_STREAM=AI` env var |

> **Note**: The `chronos` stream was previously used by Chronos for timer events but is gone with the chronos service removal. Jobsimulation no longer has chronos as a dependency.

## Key Flows

### 1. User Authentication
`Frontend` -> `Backend` -> `Sentinel`
*   The Backend validates requests using Sentinel.
*   **Studio Desk** authenticates directly via **Clerk**.

### 2. Job Simulation
`Frontend` -> `Backend` / `Jobsimulation`
*   Jobsimulation fetches the simulation **definition** (the `simulations` content/blueprint) from `CMS` by ID (`cms.GetSimulation` RPC) — it owns no content, only the run/session state.
*   Jobsimulation stores its session/run **state** (interactions, recordings, validation results, anti-cheat) via `Storage` or directly to its own `jobsimulation` DB schema.
*   Voice flows go through LiveKit; video recordings via AWS Chime SDK.

### 3. Content Delivery
`Frontend` -> `CMS` -> `Directus`
*   CMS acts as the gateway to Directus content.

### 4. Studio Content Creation
`Studio Desk` → `CMS` → (in-process) `Studio Room`
*   **Studio-Desk** (TypeScript) creates blueprints, sent to CMS as `StudioDocument` rows.
*   **CMS** (Go) creates `StudioTask` records and dispatches generation work.
*   **Studio-Room** (Python, embedded inside the CMS container at `cms/studio/`) executes the generation pipeline against AI providers (OpenAI, Anthropic, Mistral).
*   Final content is persisted via the CMS service; **Directus** is the underlying storage backend.

### 5. Skill Path Progress (Event-Driven)
`Jobsimulation` -> `Redis Stream` -> `Skillpath`
*   When a user completes a simulation, **Jobsimulation** publishes an event.
*   **Skillpath** subscribes to the Jobsimulation stream and updates step/chapter/path progress **state** (`SkillPathSession → ChapterSession → StepSession`) — it owns no content, only the per-user progression state.
*   Skillpath queries **CMS** (RPC) for the skill-path **content** structure (chapters → steps it tracks against) and **Sentinel** for authorization.

### 6. Document → PDF Conversion
`Backend (app)` → `Gotenberg`
*   The backend service uses Gotenberg's `/forms/libreoffice/convert` endpoint to render Office documents to PDF. See `app/internal/converter/gotenberg.go`.
*   `GOTENBERG_URL=http://gotenberg:3200` is injected via the backend's compose env.
