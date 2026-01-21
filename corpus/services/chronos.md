# Chronos Service

## Role & Responsibility
*   **Primary Goal**: Scheduler and Time-based Event Manager.
*   **Key Functions**:
    *   Handling delayed jobs, crons, or time-sensitive triggers for simulation/workflows.

## Architecture & Code Map
*   **Codebase**: **Remote Only** (`git@github.com:anthropos-work/chronos.git`).
*   **Language**: Go.
*   **Database**: PostgreSQL (`search_path=chronos`).
*   **Port**: `8500` (HTTP), `8501` (RPC).

## Interface Discovery
*   **API**: RPC/HTTP on 8500/8501.

## Local Development
*   **Docker Only**:
    ```bash
    cd platform
    docker compose up -d chronos
    ```
