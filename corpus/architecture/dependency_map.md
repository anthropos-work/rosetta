# Anthropos Services Dependency Map

This document outlines the inter-service dependencies inferred from configuration files (`docker-compose.yaml`) and code inspections.

## Dependency Matrix

| Service | Depends On (Direct) | Infrastructure |
| :--- | :--- | :--- |
| **Backend** (`app`) | Sentinel, CMS, Skillpath (Inferred) | Postgres, Redis, **Clerk** |
| **CMS** | - | Postgres, Redis, **Directus** |
| **Sentinel** | - | Postgres |
| **Jobsimulation** | Sentinel, Backend, Storage, Chronos, CMS | Postgres, Redis |
| **Skiller** | - | Postgres, Redis |
| **Skillpath** | CMS, Jobsimulation (RPC + Redis Stream), Sentinel | Postgres, Redis |
| **Storage** | CMS | Postgres, Redis |
| **Chronos** | - | Postgres, Redis |
| **Intelligence** | - | Postgres (Connection to Backend/Skiller DB) |
| **Graphql** | Backend, Skiller, Jobsimulation, Intelligence, CMS | - |
| **Studio Desk** | - | **Clerk**, **OpenAI** (Copilot) |
| **Studio Room** | CMS | **OpenAI**, **Anthropic**, **Azure OpenAI** |

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
