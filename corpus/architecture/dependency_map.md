# Anthropos Services Dependency Map

This document outlines the inter-service dependencies inferred from configuration files (`docker-compose.yaml`) and code inspections.

## Dependency Matrix

Sourced from `platform/docker-compose.yml` `depends_on:` declarations and environment variables (`*_RPC_ADDR`).

| Service | Depends On (Direct) | Infrastructure |
| :--- | :--- | :--- |
| **Backend** (`app`) | Sentinel, CMS, Skiller, Skillpath, Storage, Gotenberg | Postgres, Redis, **Clerk** |
| **CMS** | Sentinel, Skiller, Storage | Postgres, Redis, **Directus**, **AI Providers** (Anthropic, OpenAI, Mistral — via embedded studio-room) |
| **Sentinel** | - | Postgres |
| **Jobsimulation** | Sentinel, Backend, CMS, Roadrunner, Skiller, Storage | Postgres, Redis, **LiveKit**, **AWS Chime**, **AI Providers** |
| **Skiller** | Sentinel | Postgres (with `pgvector` in `extensions` schema), Redis, **AI Providers** (embeddings) |
| **Skillpath** | Sentinel, CMS, Jobsimulation (RPC + Redis Stream) | Postgres, Redis |
| **Storage** | - | Postgres, Redis, **S3** |
| **Roadrunner** | - | Redis, **Judge0** (code execution) |
| **Gotenberg** | - | - (stateless conversion service) |
| **Messenger** (opt-in profile) | Backend, CMS, Jobsimulation, Skiller, Skillpath | Postgres, Redis, **Brevo** (email delivery) |
| **CustomerIO Sync** (opt-in profile) | Postgres | **Customer.io** |
| **Graphql (Cosmo Router)** | Backend, Skiller, Jobsimulation, CMS, Skillpath, Storage | - |
| **Studio-Desk** (opt-in profile) | Graphql, CMS | **Clerk**, **OpenAI** (Copilot) |
| **Studio-Room** | (runs inside CMS container; depends on CMS process) | **OpenAI**, **Anthropic**, **Mistral** |

Production-only:
| Service | Depends On (Direct) | Infrastructure |
| :--- | :--- | :--- |
| **db-backup** | - | Postgres, **S3**, **Azure**, **Hetzner** |

### Shared Libraries

These are imported as Go modules by services, not deployed independently:

| Library | Used By |
| :--- | :--- |
| **colony** | All Go services (logging, DB, Redis, middleware, pub/sub) |
| **authn** | All services needing Clerk JWT auth |
| **proto** | All services using RPC (contract definitions) |
| **ai** | Jobsimulation, Skiller, CMS, Studio-Desk (AI provider wrapper) |
| **taxonomy** | Skiller, CMS (60K skills, 18K roles data) |

## Event Streams (Redis Streams via Watermill)

Services communicate asynchronously through named Redis Streams. Stream names come from `*_STREAM` env vars in `platform/docker-compose.yml`.

| Stream Name | Producer | Consumer(s) | Events |
| :--- | :--- | :--- | :--- |
| `backend` | App | CMS, Skiller | User/org updates |
| `skiller` | Skiller | App | Skill score changes |
| `jobsimulation` | Jobsimulation | Skillpath, App, Messenger (if running) | Session completed, insights generated |
| `cms` | CMS | Jobsimulation, Skillpath, Backend | Content published |
| `skillpath` | Skillpath | App | Session updated, chapters completed |
| `AI` | (multiple) | (multiple) | AI usage / cost telemetry — see `AI_USAGE_STREAM=AI` env var |

> **Note**: The `chronos` stream was previously used by Chronos for timer events but is gone with the chronos service removal. Jobsimulation no longer has chronos as a dependency.

## Key Flows

### 1. User Authentication
`Frontend` -> `Backend` -> `Sentinel`
*   The Backend validates requests using Sentinel.
*   **Studio Desk** authenticates directly via **Clerk**.

### 2. Job Simulation
`Frontend` -> `Backend` / `Jobsimulation`
*   Jobsimulation requests content from `CMS` (RPC + Redis Stream).
*   Jobsimulation stores state changes via `Storage` or directly to DB.
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
*   **Skillpath** subscribes to the Jobsimulation stream and updates step/chapter/path progress.
*   Skillpath queries **CMS** (RPC) for skill path structure and **Sentinel** for authorization.

### 6. Document → PDF Conversion
`Backend (app)` → `Gotenberg`
*   The backend service uses Gotenberg's `/forms/libreoffice/convert` endpoint to render Office documents to PDF. See `app/internal/converter/gotenberg.go`.
*   `GOTENBERG_URL=http://gotenberg:3200` is injected via the backend's compose env.
