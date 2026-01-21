# Backend Service (`app`)

## Role & Responsibility
*   **Primary Goal**: The central API Gateway and User Management service for the proper platform.
*   **Key Functions**:
    *   **User Management**: Handles user profiles, authentication context (via Sentinel), and retrieval (`GetUser`).
    *   **Organization Management**: Manages organization membership, roles, and details (`GetOrganizationRoles`, `GetUserOrganizations`).
    *   **Feature Gating**: Checks permissions for features (`CanPerformFeatureAction`).
    *   **Experience System**: Tracks and retrieves user experience points (`GetExperiencePoints`).

## Architecture & Code Map
*   **Codebase**: `app` (Local directory), `git@github.com:anthropos-work/app.git` (Remote).
*   **Language**: Go.
*   **Database**: PostgreSQL (managed via `ent` ORM).
*   **Key Directories**:
    *   `rpc.go`: **Main Entry Point for API**. Contains the `backendRPCServer` and `orgRPCServer` implementations.
    *   `internal/app`: likely wire-up of components.
    *   `internal/data/ent`: Database schema and ORM code.
    *   `internal/organization`, `internal/experiencepoint`: Domain logic.

## Interface Discovery
*   **How to find the API**:
    *   The service uses **Connect RPC**.
    *   Look at **`rpc.go`** to see the implemented handlers (e.g., `GetUser`, `GetOrganizationDetails`).
    *   The Protocol Buffers definitions are imported from `github.com/anthropos-work/proto`.
*   **Dependencies**:
    *   **Upstream**: Next Web App (calls these RPCs).
    *   **Downstream**:
        *   **Sentinel** (Authentication/Authorization).
        *   **PostgreSQL** (Data persistence).
        *   **Redis** (Cache & Streams).

## Local Development (The "How-To")

### 1. Running Standalone
*   **Prerequisites**:
    *   `.env` file (ask EM).
    *   Running Postgres and Redis (can be done via `docker compose up -d redis postgresql` in `platform`).
*   **Setup**:
    ```bash
    make setup  # Installs tools like mockgen, ent, atlas
    make gen    # Runs code generation
    ```
*   **Database Migrations**:
    ```bash
    atlas migrate apply --env local
    ```
*   **Run**:
    ```bash
    go run .
    ```

### 2. Running in Docker
*   **Service Name**: `backend`
*   **Command**:
    ```bash
    cd platform
    docker compose up -d backend
    ```

### 3. Testing
*   **Command**:
    ```bash
    go test ./...
    ```
