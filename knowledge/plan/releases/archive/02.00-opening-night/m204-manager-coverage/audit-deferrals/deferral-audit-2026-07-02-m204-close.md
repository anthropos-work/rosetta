---
title: "Deferral Audit — M204 Manager-vantage coverage (close)"
date: 2026-07-02
scope: milestone
invoked-by: close-milestone
---

## Verdict
**GREEN**

- No repeat deferrals. The one milestone-relevant deferral (the assign-WRITE UC) is a **single, in-manifest
  Fate-2 tracked gap** — declared once (M201), never pushed forward milestone-after-milestone.
- No aging trigger fires (declared this release, not routed-to-a-closed-milestone, area not re-scoped).
- Prior v2.0 close audits (M202, M203) were both GREEN; standing backlog (DEF-M10-01 / DEF-M21-01 / M25-D9)
  is pre-v2.0 and unrelated to M204.

## Summary
- Total deferrals in scope (M204-relevant): **1**
- Single deferrals: **1**
- Repeat deferrals: **0**
- Chronic patterns flagged: **0**
- Aged-out: **0**

## Deferral Inventory

```yaml
- id: DEF-M204-01
  item: "assignment-monitoring.assign-and-track.UC1 — the assign-WRITE half (a manager assigns a legacy
         skill-path/sim/interview to a person/team with a deadline; a two-backend org-admin WRITE flow)."
  origin_milestone: M201 (declared in the manifest corpus alongside UC2)
  first_deferred_on: 2026-06-29   # M201 manifest-draft authoring — declared as one of two flows
  last_seen_in: >
    .agentspace/rosetta-extensions/playthroughs/manifest/assignment-monitoring.yaml:50
    (playthrough: TODO — "build-reference gap: the assign WRITE half is out of M204's declared manager journeys")
  destination: "tracked in-manifest as a declared TODO (build-reference gap); reports `unimplemented`, NOT `failing`"
  reason_recorded: >
    M204's declared 3 manager journeys cover the MONITORING (drill-down) half — UC2, the per-member activity
    dashboard, which LANDED green (pt-activity-drilldown). UC1 is the WRITE half (a two-backend org-admin write)
    and was OUT of M204's declared manager journeys from the start (overview.md Goal + decisions TOK-01).
  partial_attempted: no   # the two halves are distinct flows; UC2 landed in full, UC1 was never in M204 scope
```

## Repeat-Deferral Patterns
None. `assign-and-track.UC1` was declared ONCE (M201) as one of two distinct flows in the `assign-and-track`
story. M204 covered the sibling flow (UC2, monitoring). This is not the same item deferred across ≥2 milestones —
it is a single declared build-reference gap tracked in the corpus manifest.

## Fate-1 Investigation

### DEF-M204-01 — "assign-WRITE UC1"
- **Fate-1 (land now, complete) feasible:** no.
- **Why a complete landing now is genuinely infeasible in M204:** UC1 is a **WRITE** flow (create an assignment
  → target a person/team → set a deadline → confirm), spanning **two backends** (the assignment write + the
  target-member fan-out), and is **out of M204's declared 3 manager journeys** (which are all READ/monitoring
  flows). M204's exit gate is defined over the 3 declared manager journeys (funnel+roster, activity-dashboard
  drill-down, succession/at-risk) — all four UCs of that gate landed green with 0 false-fails / 5. Pulling a new
  two-backend WRITE journey into M204 at close is scope creep, not a Fate-1 completion of the milestone's
  declared work. It is a *new manager-WRITE journey*, not the tail of an in-flight one.
- **Which fate applies:** **Fate-2** — the item is already **tracked** (in the corpus, durably): the
  `assignment-monitoring.yaml` manifest declares `UC1 … playthrough: TODO` with header prose recording it as a
  build-reference gap, the `ptreport` 4-state map surfaces it as `unimplemented` (a first-class tracked state, not
  a silent drop), and the M204 harden Pass-2 deliverable-presence pin (`TestRealCorpus_ManagerCoverageIsPresent`)
  **enforces** that the UC1 gap stays declared-as-TODO (fail-red if silently removed). No current v2.0 milestone
  owns the manager-WRITE class — M205 is Hiring/tier-gates, M206 AI-sim-mirror-tier, M207 Academy. A future
  manager-write tier is the natural home; the manifest TODO is the durable pointer that survives this close.
- **Enumerate what failed last time:** nothing "failed" — UC1 was never attempted in M204; it was declared
  out-of-scope at M201 (two distinct flows) and confirmed out-of-scope in the M204 bootstrap tok (TOK-01: the 3
  declared journeys are the funnel/roster, drill-down, succession).

## Recommendations
- **DEF-M204-01 → LAND-NEXT (Fate-2, no plan edit).** Confirmed tracked in-manifest as a declared build-reference
  gap (`playthrough: TODO`, `unimplemented` state, presence-pinned by the harden Pass-2 test). No milestone-plan
  edit is required: the manifest IS the durable tracking surface, and it survives the close as a corpus artifact.
  When a future major schedules a manager-WRITE tier, this manifest TODO is the ready-made declaration to build
  against. Recorded as the close decision D-CLOSE-1 (see `decisions.md`).

## Applied Changes
- Recorded the Fate-2 verdict for DEF-M204-01 as D-CLOSE-1 in `m204-manager-coverage/decisions.md`.
- No plan edit (Fate-2 = confirm-covered, not annotate). No `overview.md` / `roadmap-vision.md` mutation.

## Blocking Items (require user decision)
None. Zero repeat-deferrals, zero aged-out items, zero escape-hatch entries. GREEN.
