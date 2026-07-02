---
iter: 04
milestone: M204
iteration_type: tik
status: closed-fixed
created: 2026-07-02
---

# M204 iter-04 — tik: Succession / at-risk (the last manager surface)

**Type:** tik

**Active strategy reference:** TOK-01 (manager-surface-per-iter).

**Step 0 re-survey:** metric = 3/4 declared manager UCs passing (funnel, roster, drill-down). TOK-01's last
named target (succession/at-risk) is untouched + meaningful. No substitution.

**Cluster / target identified:** journey 3 of 3 (the LAST) — succession / at-risk / mobility
(`workforce-intelligence.talent-pool.UC1`, `/enterprise/workforce/succession`). Probe confirmed the surface
renders "Succession Planning" + "Top talents (ready)" succession candidates + "People at risk" / "At-risk people"
+ 10 role→candidate rows.

**Hypothesis:** the base pt-world model's activity + trajectory data drives the computed succession/at-risk
projections at Org A → the surface goes green with a page object + spec, no seed expansion.

**Expected lift:** +1 passing manager UC (3 → 4) — completes all 3 declared manager journeys.

**Phase plan:** probe → declare (extend `workforce.yaml` with the `talent-pool` story + UC1) → page-object
(`succession-page.ts` + `SUCCESSION_URL` route shape) → play → re-measure.

**Escalation conditions:** empty projections at Org A → seed-expansion route-forward; un-drivable → escalate.

**Acceptable close-no-lift outcomes:** if succession/at-risk rendered empty/"too sparse" at size 40 (the audit
flagged Org A has only a thriving end-user hero, no struggling one), the falsification "the at-risk projection
needs a struggling-trajectory member in Org A" would route a seed-expansion iter.

## Outcome
+1 passing manager UC (3 → 4) — **all 3 declared manager journeys / 4 manager UCs now pass green.** The base
pt-world model's data was sufficient at Org A (size 40) — succession/at-risk render real projections without seed
expansion (the audit's at-risk concern did not materialize — the surface computes at-risk across the org's
trajectory data, which the base activity provides). `succession-page.ts` + the `talent-pool` UC + the spec;
green + `[PASS]`; all 4 manager Playthroughs pass together.

The FUNCTIONAL half of the gate is met (every declared manager UC passing). The DETERMINISM half (0 false-fails
over 5 consecutive reset runs) is the next iter's work.

## Close
See `progress.md`.
