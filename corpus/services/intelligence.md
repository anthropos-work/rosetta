# Intelligence Service

## Role & Responsibility
*   **Primary Goal**: AI/ML Integration service.
*   **Key Functions**:
    *   Likely interfaces with LLMs (Mistral, OpenAI) or internal models.

## Architecture & Code Map
*   **Codebase**: **Remote Only** (`git@github.com:anthropos-work/intelligence.git`).
*   **Language**: Go.
*   **Connection**: Interfaces with Backend and Skiller via DB Connections (strange pattern? `DB_CONNECTION_BACKEND`, `DB_CONNECTION_SKILLER` env vars in compose).
*   **Port**: Not exposed to host in `docker-compose.yml`, suggesting it might process tasks asynchronously (Redis Streams?) or be internal only.

## Interface Discovery
*   **Async**: Likely listens on Redis Streams or connects directly to other service DBs (based on env vars).

## Local Development
*   **Docker Only**:
    ```bash
    cd platform
    docker compose up -d intelligence
    ```
