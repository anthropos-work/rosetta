# Anthropos Services Dependency Map

This document outlines the inter-service dependencies inferred from configuration files (`docker-compose.yaml`) and code inspections.

## Dependency Matrix

| Service | Depends On (Direct) | Infrastructure |
| :--- | :--- | :--- |
| **Backend** (`app`) | Sentinel, CMS, Skillpath, Messenger | Postgres, Redis, **Clerk** |
| **CMS** | - | Postgres, Redis, **Directus** |
| **Sentinel** | - | Postgres |
| **Jobsimulation** | Sentinel, Backend, Storage, Chronos, CMS, Roadrunner | Postgres, Redis, **LiveKit**, **AWS Chime**, **AI Providers** |
| **Skiller** | - | Postgres, Redis, **AI Providers** (embeddings) |
| **Skillpath** | CMS, Jobsimulation (RPC + Redis Stream), Sentinel | Postgres, Redis |
| **Storage** | CMS | Postgres, Redis, **S3** |
| **Chronos** | - | Postgres, Redis |
| **Intelligence** | - | Postgres (Connection to Backend/Skiller DB) |
| **Messenger** | - | **Brevo** (email delivery) |
| **Roadrunner** | - | **Judge0** (code execution) |
| **db-backup** | - | Postgres, **S3**, **Azure**, **Hetzner** |
| **Graphql (Cosmo Router)** | Backend, Skiller, Jobsimulation, CMS, Skillpath | - |
| **Studio Desk** | - | **Clerk**, **OpenAI** (Copilot) |
| **Studio Room** | CMS | **OpenAI**, **Anthropic**, **Azure OpenAI** |

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

Services communicate asynchronously through named Redis Streams:

| Stream Name | Producer | Consumer(s) | Events |
| :--- | :--- | :--- | :--- |
| `backend` | App | Skiller, Intelligence | User/org updates |
| `skiller` | Skiller | App, Intelligence | Skill score changes |
| `jobsimulation` | Jobsimulation | Skillpath, App | Session completed, insights generated |
| `cms` | CMS | Jobsimulation, Skillpath | Content published |
| `skillpath` | Skillpath | App | Session updated, chapters completed |
| `chronos` | Chronos | Jobsimulation | Timer events (e.g., session timeout) |

## Key Flows

### 1. User Authentication
`Frontend` -> `Backend` -> `Sentinel`
*   The Backend validates requests using Sentinel.
*   **Studio Desk** authenticates directly via **Clerk**.

### 2. Job Simulation
`Frontend` -> `Backend` / `Jobsimulation`
*   Jobsimulation likely requests content from `CMS`.
*   Jobsimulation stores state changes via `Storage` or directly to DB.
*   Jobsimulation schedules events via `Chronos`.

### 3. Content Delivery
`Frontend` -> `CMS` -> `Directus`
*   CMS acts as the gateway to Directus content.

### 4. Studio Content Creation
`Studio Desk` -> `CMS` -> `Studio Room`
*   **Studio Desk** creates blueprints.
*   **Studio Room** reads blueprints from **CMS** and uses AI Providers to generate content.
*   **Directus** serves as the storage backend for CMS content.

### 5. Skill Path Progress (Event-Driven)
`Jobsimulation` -> `Redis Stream` -> `Skillpath`
*   When a user completes a simulation, **Jobsimulation** publishes an event.
*   **Skillpath** subscribes to the Jobsimulation stream and updates step/chapter/path progress.
*   Skillpath queries **CMS** (RPC) for skill path structure and **Sentinel** for authorization.
