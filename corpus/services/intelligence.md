# Intelligence Service

> ## ⚠️ Archived — no longer in local orchestration
>
> Intelligence was removed from `platform/docker-compose.yml` and `platform/repos.yml` in mid-2026:
> - Platform: commit `fdfa189` — "remove intelligence service from local dev orchestration"
>
> The Intelligence GitHub repository still exists but is no longer cloned by `make init` and no service in the current compose file depends on it. The original role was background data sync between the backend (`public`) and skiller schemas via direct DB connections (`DB_CONNECTION_BACKEND`, `DB_CONNECTION_SKILLER`). The cross-schema reads it performed are now handled inside the consuming services themselves. *(Note: skiller has since been merged into app — the taxonomy domain and its data now live in the `app` service's `public` schema; see [backend.md](./backend.md).)*
>
> The skeleton below is preserved for historical context. If intelligence is ever brought back, the doc should be rewritten against the current code in the repo.

## Historical Role (pre-2026-Q2)

* **Primary Goal**: Background data sync between Backend and Skiller schemas.
* **Codebase**: `git@github.com:anthropos-work/intelligence.git` (no longer cloned locally)
* **Language**: Go
* **Connection**: Interfaced with Backend (`public`) and Skiller schemas via direct DB connections (`DB_CONNECTION_BACKEND`, `DB_CONNECTION_SKILLER` env vars).
* **Port**: Exposed an HTTP server on :8080 (PORT env override) serving only a /_meta health endpoint for ECS checks; all real work ran on a 5-minute async ticker, not via the API.
