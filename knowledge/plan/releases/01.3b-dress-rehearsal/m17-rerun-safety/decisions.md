# M17 — Decisions

_Implementation decisions with rationale. ID scheme: M17-D1, M17-D2, … Open questions from `overview.md` get resolved here during build._

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M17-D1 | The `set -e` race audit found **three** kin sites of the ISSUE-7 class (not just the one fixed). Fix: `up-injected.sh` GH_PAT → **fail loud** (`\|\| true` + empty-check); `rosetta-demo` + `dev-stack` DEV_PROJECT → `\|\| true` so the documented `${DEV_PROJECT:-anthropos}` fallback actually runs (it was unreachable — the grep-miss aborted first). | All three abort the script under `set -euo pipefail` on a grep-miss/missing-.env, before their guard/fallback can run. GH_PAT-miss = fail-fast (build needs it) → loud; DEV_PROJECT-miss = has a documented fallback → keep it live. Verified by simulation. | 2026-06-08 |
| M17-D2 | wait-for-sentinel-ready defense-in-depth lives in **migrate-demo.sh** (bounded `wait_pg` via `pg_isready`/`SELECT 1` + `wait_sentinel_running` via `docker inspect`), both **non-fatal** (timeout → WARN → proceed). | migrate-demo.sh owns the sentinel touch + the first `docker exec psql`; waiting there closes the ISSUE-7 race at its SOURCE (proactive) on top of the existing reactive `\|\| echo 0`. Non-fatal so it only removes flakiness, never adds a hard failure (downstream is `ON_ERROR_STOP=0` + idempotent CREATE+INSERT). Functions used as `if`-conditions → `set -e` suspended in their body (verified). | 2026-06-08 |
