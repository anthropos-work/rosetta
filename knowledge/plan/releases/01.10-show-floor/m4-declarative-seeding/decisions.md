# M4 — Decisions

## M4-D1 — fold validation/dry-run into M4 (no separate M4b yet) (design, 2026-06-03)
The Phase 2b tooling pass flagged a candidate **M4b (seed-config validator + dry-run)**. Decision: **first fold
`--validate` + `--dry-run` into the M4 seeder itself** (same author-context, cheap). Split out an M4b only if that
surface grows beyond a flag (e.g. a standalone DAG resolver + linter + `/demo-seed --check` skill worth its own
milestone). Recorded so the build knows the hedge.

## Open (resolve during build)
- single config vs per-store seeders (recommend single, fanned out).
- pre-embedded skiller snapshot provenance (no prod data).
- Directus content tenancy (shared vs per-demo).
- CLI-as-binary vs linked `app/internal/bootstrap` for 1k-scale speed.
- backdating fidelity (ent-Immutable / DB-default timestamps).
