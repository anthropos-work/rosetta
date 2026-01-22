# Anthropos Services Dependency Map

This document outlines the inter-service dependencies inferred from configuration files (`docker-compose.yaml`) and code inspections.

## Dependency Matrix

| Service | Depends On (Direct) | Infrastructure |
| :--- | :--- | :--- |
| **Backend** (`app`) | Sentinel, CMS, Skillpath (Inferred) | Postgres, Redis |
| **CMS** | - | Postgres, Redis, **Directus** |
| **Sentinel** | - | Postgres |
| **Jobsimulation** | Sentinel, Backend, Storage, Chronos, CMS | Postgres, Redis |
| **Skiller** | - | Postgres, Redis |
| **Skillpath** | CMS | Postgres, Redis |
| **Storage** | CMS | Postgres, Redis |
| **Chronos** | - | Postgres, Redis |
| **Intelligence** | - | Postgres (Connection to Backend/Skiller DB) |
| **Graphql** | Backend, Skiller, Jobsimulation, Intelligence, CMS | - |

## Key Flows

### 1. User Authentication
`Frontend` -> `Backend` -> `Sentinel`
*   The Backend validates requests using Sentinel.

### 2. Job Simulation
`Frontend` -> `Backend` / `Jobsimulation`
*   Jobsimulation likely requests content from `CMS`.
*   Jobsimulation stores state changes via `Storage` or directly to DB.
*   Jobsimulation schedules events via `Chronos`.

### 3. Content Delivery
`Frontend` -> `CMS` -> `Directus`
*   CMS acts as the gateway to Directus content.
