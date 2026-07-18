# iter-01 — decisions (iter-local)

_Strategy-level decisions live in the milestone-root `decisions.md` (TOK-01). These are intra-iter._

| # | Decision | Rationale | Date |
|---|----------|-----------|------|
| D1 | Fix the `hiring.md` isHiring-wiring pointer inline (Phase 0b applied fix): name the FAPI `clerk-frontend/resources.go::orgMemberships` as the browser-visible org-publicMetadata emission point (fed by the `RosterEntry` roster thread), demote the BAPI `clerk-backend` citation to "server-side, optional". | The old pointer sent an M224 author to the wrong file — the client re-skin reads the FAPI, not the BAPI. Small + unambiguous (Explore traced it definitively) → apply inline per audit Phase 6. | 2026-07-16 |
| D2 | **Recruiter-render-first** ordering (→ TOK-01): drive the recruiter scoreboard to green BEFORE the 2 candidate `/profile` heroes. | The exit gate IS the recruiter's comparison surface (≥40 rows/sim × 5); the candidate `/profile` heroes are In-scope polish (swept at M225), not the gate metric. Front-load the gate. | 2026-07-16 |
| D3 | The primary metric is `min(rows-painted across the 5 sims)` on a cold-seeded scoreboard, not a sum/mean. | The gate is per-EACH-of-5 → the binding number is the minimum; a sum could hide a 0-row sim. | 2026-07-16 |
