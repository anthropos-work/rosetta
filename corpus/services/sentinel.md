# Sentinel Service

## Role & Responsibility
*   **Primary Goal**: The Authentication and Authorization authority for the platform.
*   **Key Functions**:
    *   **Authentication**: Verifies user identities (likely handling Tokens/Sessions).
    *   **Authorization**: Manages access control (Service-to-Service and User privileges).

## Architecture & Code Map
*   **Codebase**: **Remote Only** (Not checked out locally).
    *   Repo: `git@github.com:anthropos-work/sentinel.git`
*   **Language**: Go (Inferred from generic platform stack).
*   **Database**: PostgreSQL (`search_path=sentinel`).

## Interface Discovery
*   **How to find the API**:
    *   Since the code is not local, check `platform/docker-compose.yml` for exposed ports (`8087`).
    *   In `backend` service, look for usage of `AUTHORIZATION_ADDRESS` to see how it is called (likely HTTP/RPC).

## Local Development (The "How-To")

### 1. Running in Docker (Recommended)
Since we don't have the source code locally, we must run the Docker image built from git.

*   **Service Name**: `sentinel`
*   **Command**:
    ```bash
    cd platform
    docker compose up -d sentinel
    ```

### 2. Running Standalone
*   *Not possible without cloning the repository.*
