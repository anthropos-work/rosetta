# Storage Service

## Role & Responsibility
*   **Primary Goal**: Centralized file and blob storage management.
*   **Key Functions**:
    *   Likely handles S3/Local file uploads and retrieval URLs.

## Architecture & Code Map
*   **Codebase**: **Remote Only** (`git@github.com:anthropos-work/storage.git`).
*   **Language**: Go.
*   **Port**: `8300` (HTTP), `8301` (RPC).

## Interface Discovery
*   **API**: RPC/HTTP on 8300/8301.

## Local Development
*   **Docker Only**:
    ```bash
    cd platform
    docker compose up -d storage
    ```
