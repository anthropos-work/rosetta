**Type:** tik — under TOK-01 (manager-surface-per-iter). Protocol: `corpus/ops/demo/playthroughs.md`.

# iter-04 — Succession / at-risk (the last manager surface)

Probed `/enterprise/workforce/succession` as Morgan → "Succession Planning" + "Top talents (ready)" succession
candidates + "People at risk" / "At-risk people" + 10 role→candidate rows render real projections. Authored the
`talent-pool` story + UC1 in `workforce.yaml`, `SuccessionPage`, and `SUCCESSION_URL` (+ pure-logic unit
coverage). The Playthrough passes green; all 4 manager Playthroughs pass together.

## Close — 2026-07-02

**Outcome:** +1 passing manager UC (`workforce-intelligence.talent-pool.UC1`); metric 3 → 4. **All 3 declared
manager journeys / 4 manager UCs pass green.** No seed expansion needed (the audit's at-risk concern did not
materialize — the surface computes at-risk across the org's trajectory data).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (the FUNCTIONAL half is met — every declared manager UC passing; the DETERMINISM half [0 false-fails over 5 consecutive reset-to-seed runs] is NOT yet run — that is iter-05's work)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (3 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** none new (a clean surface add under the established pattern; the succession page object scopes to `main()`, which DOES contain this surface's content, unlike iter-03's drill-down detail).
**Side-deliverables:** none.
**Routes carried forward:** iter-05 (next tik) → the **5-run reset-to-seed determinism gate** — run the full
Playthrough suite (all 4 manager + all M203 employee UCs) under `--reset` for 5 consecutive runs, prove 0
false-fails (the DETERMINISM half of the exit gate). This is the last gate requirement.
**Lessons:** the base pt-world stories model at Org A (size 40) renders ALL the M36 manager org-dashboard
surfaces (funnel, roster, activity, succession/at-risk) with real data — the whole manager vantage was drivable
with NO seed expansion (the Phase 0b known-context resolved fully positive). Not every surface's content is in
`<main>` (iter-03's drill-down detail was the exception); the succession surface IS in `<main>`.
