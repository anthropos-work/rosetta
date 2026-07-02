**Type:** tik — under TOK-01 (manager-surface-per-iter). Protocol: `corpus/ops/demo/playthroughs.md`.

# iter-03 — Member drill-down (activity-dashboard)

Probed the activity dashboard as Morgan → real per-content activity (20 rows, Passed/Failed status mix) + a
drill-down to a per-member "Simulation Results" table. Authored `assignment-monitoring.yaml` (UC2 live; UC1
assign-WRITE `TODO`, a build-reference gap out of M204's declared journeys), `ActivityDashboardPage`, and the
activity route shapes (`ACTIVITY_DASHBOARD_URL` / `ACTIVITY_DRILLDOWN_URL` + pure-logic unit coverage). Two
in-scope drivability fixes got UC2 green (D1 URL race, D2 out-of-main table).

## Close — 2026-07-02

**Outcome:** +1 passing manager UC (`assignment-monitoring.assign-and-track.UC2`); metric 2 → 3. Green + `[PASS]`
reconciled; all 3 manager Playthroughs pass together.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (3 of the 3 declared manager journeys covered — funnel+roster, drill-down, and succession
STILL TO DO; the 5-run zero-false-fail determinism gate not yet run)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (drill-down URL race → waitForURL), D2 (drill-down results table is OUTSIDE `<main>` → page-level scope). Both in iter-03/decisions.md; both in-scope (UC2's own drivability).
**Side-deliverables:** none (D1/D2 are UC2's planned-scope drivability, not unrelated side-fixes).
**Routes carried forward:** iter-04 (next tik) → **Succession / at-risk** — `workforce-intelligence.talent-pool.UC1`
(`/enterprise/workforce/succession`; probe confirmed "Succession Planning" heading + 10 role→candidate rows +
Succession/At-Risk labels render). This is the LAST declared manager journey → then the 5-run reset-to-seed
determinism gate.
**Lessons:** (1) SPA client-side navigations need `page.waitForURL(...)` (auto-waits), never a synchronous
`page.url()` right after a click — the URL lags the click (a green-but-wrong-adjacent race). (2) NOT every
surface renders its table inside `<main>` — the activity drill-down detail renders its per-member table in a
plain-div layout (2 `<main>` present, table `.closest('main')` is false); scope to the page-level single results
table, disambiguated by the surface + heading, not `main()`. Both generalize; noted for future manager surfaces.
