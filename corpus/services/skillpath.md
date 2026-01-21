# Skillpath Service

## Role & Responsibility
*   **Primary Goal**: Defines and manages the progression paths for skills (Curriculum/Learning Paths).
*   **Key Functions**:
    *   **Path Management**: Organizing skills into sequences or trees.

## Architecture & Code Map
*   **Codebase**: **Remote Only** (`git@github.com:anthropos-work/skillpath.git`).
*   **Language**: Go (Inferred).
*   **Database**: PostgreSQL.
*   **Port**: `8100` (HTTP), `8101` (RPC).

## Interface Discovery
*   **API**: Exposed via RPC/HTTP on ports 8100/8101.

## Local Development
*   **Docker Only**:
    ```bash
    cd platform
    docker compose up -d skillpath
    ```
