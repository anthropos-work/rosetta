---
iter: 03
milestone: M204
iteration_type: tik
status: closed-fixed
created: 2026-07-02
---

# M204 iter-03 — tik: Member drill-down (activity-dashboard)

**Type:** tik

**Active strategy reference:** TOK-01 (manager-surface-per-iter).

**Step 0 re-survey:** metric = 2/4 declared manager UCs passing (funnel + roster, iter-02). TOK-01's next target
(the activity-dashboard drill-down) is untouched + meaningful. No substitution.

**Cluster / target identified:** journey 2 of 3 — the per-member activity dashboard
(`assignment-monitoring.assign-and-track.UC2`, `/enterprise/activity-dashboard` + the per-content drill-down to a
per-member breakdown). The probe confirmed the surface renders real per-content activity (20 rows) and drilling a
content row opens a "Simulation Results" per-member table (Passed/Failed status mix across real members).

**Hypothesis:** the base pt-world model seeds real sim activity at Org A → the dashboard + drill-down go green
with a page object + spec, no seed expansion.

**Expected lift:** +1 passing manager UC (2 → 3).

**Phase plan:** probe → declare (`assignment-monitoring.yaml`; UC1 assign-WRITE stays `TODO`, a build-reference
gap out of M204's declared journeys) → page-object (`activity-dashboard-page.ts` + `ACTIVITY_DASHBOARD_URL` /
`ACTIVITY_DRILLDOWN_URL` route shapes) → play → diagnose → re-measure.

**Escalation conditions:** un-drivable drill-down → `unimplementable-without-platform-edit`; empty per-member
breakdown at Org A → seed-expansion route-forward.

**Acceptable close-no-lift outcomes:** if the activity/breakdown rendered empty at size 40, the falsification
"the base model does not seed drillable per-member sim activity" would route a seed-expansion iter.

## Outcome
+1 passing manager UC (2 → 3). Authored `assignment-monitoring.yaml` (UC2 live, UC1 `TODO`) + the
`ActivityDashboardPage` page object + the activity route shapes; the drill-down Playthrough passes green;
`ptreport` reconciles it `[PASS]`. All 3 manager Playthroughs pass together.

**Two in-scope diagnostic fixes (D1, D2) landed within the iter's planned scope** (getting the one UC2 green):
the drill-down URL race (waitForURL vs synchronous currentUrl) and the drill-down results table being OUTSIDE
`<main>` (page-level scope, not main-scope). Both are UC2's own drivability, not side-discoveries.

## Close
See `progress.md`.
