# iter-01 — progress

**Type:** tok (bootstrap) — per `coverage-protocol.md` §"Iter type selection" (the bootstrap tok authors the
strategy + scaffolds; does not terminate the call).

## Work
- Loaded the M42 coverage protocol + the annotation field-review + the four KB-contract docs + the rext
  seeder fleet.
- Ran the Phase 0b KB-fidelity gate → YELLOW (recorded in `spec-notes.md`; F1 + F2 are the must-knows).
- **Re-diagnosed the FRESH demo-1** (the orchestration's critical constraint) against its live Postgres
  (`:15432`) — demo-1 is seeded with the full stories world, so several annotation gaps from the old stale
  demo are already closed; the genuine remaining seed gaps are catalogued in `spec-notes.md`.
- Installed the coverage harness deps + chromium in the rext authoring copy (`stack-verify/e2e/`).
- Authored TOK-01 (milestone-root `decisions.md`).

## Re-diagnosis result (the strategy's evidence base)
- GENUINE seed gaps: member location/last_activity/joined_at (0/221); languages (0 rows, +world_languages
  reference empty); certs roster-coverage (2/221, hero-only); Maya XP (0) + skill-path app-mirror (0).
- LIKELY-NOT seed gaps (data present, confirm by sweep): library skill-paths (22 published), target-roles
  (76), assignments (114) — empty render is probably federation/frontend, not seed.

## Close — 2026-06-30

**Outcome:** TOK-01 authored (initial strategy: sweep-driven seed-fill of the genuine empties, re-seed to
iterate, COLD reset reserved for the exit-gate proof). FRESH-demo-1 re-diagnosis done; 4 genuine seed-gap
clusters identified + the has-data surfaces flagged for sweep confirmation.
**Type:** tok (bootstrap)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap, does not exit) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** TOK-01 (milestone-root decisions.md)
**Side-deliverables (if any):** none
**Routes carried forward:** AI-provider-keys policy (F7) → a future tik (decision deliverable, record in
secrets-spec.md; academy AI chat documented-as-absent, not a gate blocker). Academy menu-link + non-anonymous
session (F6) → a future tik. The four genuine seed-gap clusters → iter-02 baseline-then-target.
**Lessons:** A FRESH-demo re-diagnosis before authoring fixes is load-bearing — it split the annotation's gap
list into genuine seed gaps vs has-data-but-empty-render surfaces, preventing speculative seed work on the
latter. The live schema differs from the specs' `_id` column guesses (`memberships.organization`/`.user`,
`user_*` tables' `"user"` FK) — verify columns against demo-1, never trust the doc guess.
