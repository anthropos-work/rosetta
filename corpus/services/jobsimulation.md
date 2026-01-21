# Jobsimulation Service

## Role & Responsibility
*   **Primary Goal**: Manages and executes job simulations, which are core interactive experiences for the user.
*   **Key Functions**:
    *   **Simulation Engine**: Runs the logic for simulations.
    *   **GraphQL API**: Exposes simulation state and activities (`internal/graph/schemas/simulation.graphqls`).
    *   **RPC Server**: Accepts internal commands (`internal/rpcsrv`).

## Architecture & Code Map
*   **Codebase**: `jobsimulation` (Local directory).
*   **Language**: Go.
*   **Database**: PostgreSQL (via `ent`).
*   **Key Directories**:
    *   `internal/graph`: **GraphQL Implementation**.
    *   `internal/rpcsrv`: **RPC Server**.
    *   `cmd/`: Entry point.
    *   `internal/colony`: Likely shared infrastructure code.

## Interface Discovery
*   **How to find the API**:
    *   High-level interactive API: **GraphQL** (`internal/graph/schemas`).
    *   Internal Service API: **RPC** (`internal/rpcsrv`).
*   **Dependencies**:
    *   **Downstream**:
        *   **Backend** (Users).
        *   **CMS** (Content).
        *   **Storage** (Files).
        *   **Chronos** (Scheduling).

## Local Development (The "How-To")

### 1. Running Standalone
*   **Setup**:
    ```bash
    make setup  # Installs tools
    make gen    # Generates code
    ```
*   **Run**:
    ```bash
    go run main.go
    ```
    *Note: Ensure `.env` is set and dependent services (Redis, Postgres) are reachable.*

### 2. Running in Docker
*   **Service Name**: `jobsimulation`
*   **Command**:
    ```bash
    cd platform
    docker compose up -d jobsimulation
    ```
