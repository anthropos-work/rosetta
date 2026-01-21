# Anthropos Architecture Overview

This document provides a high-level overview of the Anthropos platform architecture, reversed engineered from the codebase and configuration files.

## High-Level Summary (For PMs & Non-Engineers)

Anthropos is a web platform composed of a **modern Frontend** (what users see) and a powerful **Backend** (the logic and data).

*   **The Frontend**: Built with **Next.js**, providing a fast and interactive user interface.
*   **The Backend**: A collection of specialized "microservices" (like individual workers) built with **Go** (a high-performance language). These services handle specific tasks like:
    *   **Sentinel**: Security and access control (the bouncer).
    *   **Skiller/Skillpath**: Managing user skills and learning paths.
    *   **Jobsimulation**: Running realistic job scenarios.
    *   **Intelligence**: Providing AI capabilities.
*   **The Data**: We use **PostgreSQL** to store structured data and **Directus** for managing content easily.

## Technical Deep Dive (For Engineers)

The Anthropos platform follows a **Microservices architecture**, primarily built with **Go** for the backend services and **Next.js** for the frontend. The services communicate via **RPC/HTTP** for synchronous operations and **Redis Streams** for asynchronous event-driven messaging.

```mermaid
graph TD
    subgraph Frontend
        Web[<a href='./frontend_architecture.md'>Next Web App</a>]
    end

    subgraph "Core Backend Services (Go)"
        Gateway[Backend / App Gateway]
        Sentinel[Sentinel (Auth Server)]
        CMS_Service[CMS Service]
        Skiller[Skiller]
        JobSim[Job Simulation]
        Skillpath[Skillpath]
        Storage[Storage]
        Chronos[Chronos]
        Intelligence[Intelligence]
    end

    subgraph "Data & Infra"
        Postgres[(PostgreSQL)]
        Redis[(Redis)]
        Directus[Directus CMS]
    end

    Web --> Gateway
    Gateway --> Sentinel
    Gateway --> Skiller
    Gateway --> CMS_Service
    Gateway --> JobSim
    
    CMS_Service --> Directus
    
    %% Service Inter-dependencies (Inferred)
    JobSim --> CMS_Service
    JobSim --> Storage
    JobSim --> Chronos
    
    %% Data connections
    Gateway --> Postgres
    Gateway --> Redis
    Skiller --> Postgres
    Skiller --> Redis
    JobSim --> Postgres
    JobSim --> Redis
```

### Service Inventory

| Service Name | Technology | Responsibility (Inferred) | Source |
| :--- | :--- | :--- | :--- |
| **Backend** (`app`) | Go | [Main API Gateway / User Backend](./services/backend.md) | `git@github.com:anthropos-work/app.git` |
| **CMS** | Go | [Content Management / Directus Proxy](./services/cms.md) | Local `../cms` |
| **Sentinel** | Go | [Authorization & Authentication](./services/sentinel.md) | `git@github.com:anthropos-work/sentinel.git` |
| **Jobsimulation** | Go | [Simulates job environments/tasks](./services/jobsimulation.md) | `git@github.com:anthropos-work/jobsimulation.git` |
| **Skiller** | Go | [Skill management & assessment](./services/skiller.md) | `git@github.com:anthropos-work/skiller.git` |
| **Skillpath** | Go | [Skill progression paths](./services/skillpath.md) | `git@github.com:anthropos-work/skillpath.git` |
| **Storage** | Go | [File/Blob storage management](./services/storage.md) | `git@github.com:anthropos-work/storage.git` |
| **Chronos** | Go | [Scheduling / Time-based events](./services/chronos.md) | `git@github.com:anthropos-work/chronos.git` |
| **Intelligence** | Go | [AI/ML Integration](./services/intelligence.md) | `git@github.com:anthropos-work/intelligence.git` |
| **Directus** | Node/Docker | Headless CMS (Content Store) | Docker Image |
| **Next Web App** | Next.js | Main User Interface | Local `next-web-app` |

### Communication Patterns
*   **Synchronous**: Services expose HTTP endpoints (often configured via `*_RPC_ADDR` env vars).
*   **Asynchronous**: Extensive use of **Redis Streams** for decoupling services (inferred from `REDIS_STREAM` and `REDIS_INDEX` configurations). Consumers are likely implemented in Go to process these streams.

