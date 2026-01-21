# Service Name (e.g., Backend)

## Role & Responsibility
*   **Primary Goal**: [One sentence summary] (e.g., "The main API gateway and user management service.")
*   **Key Functions**:
    *   [Function 1]
    *   [Function 2]
    *   [Function 3]

## Architecture & Code Map
*   **Codebase**: `[path/to/repo]`
*   **Language**: [Language] (Version)
*   **Database**: [Database] (e.g., Postgres, Redis)
*   **Key Directories**:
    *   `cmd/`: Entry points.
    *   `internal/`: Core logic.
    *   `rpc/` or `rpc.go`: API Definitions.

## Interface Discovery
*   **How to find the API**:
    *   [Instructions] (e.g., "Check `rpc.go` for the struct `RPC`", or "Look at `schema.graphql`")
*   **Dependencies**:
    *   Downstream: [Services this calls]
    *   Upstream: [Services that call this]

## Local Development (The "How-To")

### 1. Running Standalone
*   **Prerequisites**: [Env vars, other services running]
*   **Command**:
    ```bash
    [command]
    ```

### 2. Running in Docker
*   **Service Name**: `[service_name]` (in platform `docker-compose.yml`)
*   **Command**:
    ```bash
    docker compose up -d [service_name]
    ```

### 3. Testing
*   **Command**:
    ```bash
    [test command]
    ```
