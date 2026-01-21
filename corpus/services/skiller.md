# Skiller Service

## Role & Responsibility
*   **Primary Goal**: Manages user skills, assessments, and competency tracking.
*   **Key Functions**:
    *   **Skill Tracking**: Records user progress in various skills.
    *   **Assessments**: Likely manages logic for testing skills.

## Architecture & Code Map
*   **Codebase**: **Remote Only** (`git@github.com:anthropos-work/skiller.git`).
*   **Language**: Go (Inferred).
*   **Database**: PostgreSQL.
*   **Port**: `8085` (HTTP), `8086` (RPC).

## Interface Discovery
*   **API**: Exposed via RPC/HTTP on ports 8085/8086.
*   **Search**: Look for `SKILLER_RPC_ADDR` references in `backend` or `cms` to see how it is called.

## Local Development
*   **Docker Only**:
    ```bash
    cd platform
    docker compose up -d skiller
    ```
