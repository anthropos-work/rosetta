# CMS Service

## Role & Responsibility
*   **Primary Goal**: Manages content delivery specific to the platform, acting as a smart layer on top of Directus (Headless CMS).
*   **Key Functions**:
    *   **Content API**: Exposes a **GraphQL API** for the frontend to query content (Simulations, Skills, Resources).
    *   **Directus Integration**: Proxies or enhances data stored in Directus.
    *   **RPC Server**: Likely serves internal requests from other services (e.g., Backend, JobSim).

## Architecture & Code Map
*   **Codebase**: `cms` (Local directory).
*   **Language**: Go.
*   **Database**: PostgreSQL (via `ent`). Connects to Directus.
*   **Key Directories**:
    *   `internal/graph`: **GraphQL Implementation** (using `gqlgen`).
    *   `internal/graph/schemas/`: **GraphQL Definitions** (`*.graphqls`). Look here for the API contract!
    *   `internal/rpcsrv`: RPC Server implementation.
    *   `internal/directus`: Directus integration logic.
    *   `cmd/`: Entry point.

## Interface Discovery
*   **How to find the API**:
    *   **GraphQL**: The primary interface. Check `internal/graph/schemas/` for `.graphqls` files (e.g., `queries.graphqls`, `simulations.graphqls`).
    *   **RPC**: Check `internal/rpcsrv`.
*   **Dependencies**:
    *   **Upstream**: Next Web App (GraphQL).
    *   **Downstream**:
        *   **Directus** (Content Source).
        *   **PostgreSQL**.

## Local Development (The "How-To")

### 1. Running Standalone
*   **Prerequisites**:
    *   Access to Directus (URL/Token in `.env`).
    *   Postgres/Redis.
*   **Setup**:
    ```bash
    make setup  # Installs ent, atlas, gqlgen
    make gen    # Regenerates GraphQL resolvers & Ent code
    ```
*   **Run**:
    ```bash
    go run main.go
    ```

### 2. Running in Docker
*   **Service Name**: `cms`
*   **Command**:
    ```bash
    cd platform
    docker compose up -d cms
    ```
