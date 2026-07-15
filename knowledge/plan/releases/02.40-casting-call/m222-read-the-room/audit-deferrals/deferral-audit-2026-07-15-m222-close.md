---
title: "Deferral Audit — milestone (M222 read-the-room, close)"
date: 2026-07-15
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals; no chronic patterns; no aged-out items.
- The single scope-refinement (D4, job_position replay dropped) is an **applied Fate-3**, not a pending deferral.

## Summary
- Total deferrals in scope: **1** (a resolved Fate-3 scope-reduction — not a punt of desired work forward)
- Single deferrals: 0 pending
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0

M222 is the **first milestone of v2.4** — there are no inherited deferrals from prior milestones in this release.
All of M222's `overview.md` `Out:` items (the full 50-person seed, the assessment funnel, the cockpit heroes, any
latency work) are **Fate-2** items explicitly owned by named downstream milestones in the sequential release graph
(M223/M224/M225/M226) — the planned execution graph, not deferrals.

## Deferral Inventory

```yaml
- id: DEF-M222-01
  item: "Snapshot extension to REPLAY directus.job_position rows (originally M223 Scope.In #4)"
  origin_milestone: M222
  first_deferred_on: 2026-07-15
  last_seen_in: m222/decisions.md:D4 ; m223/overview.md:Scope.In #4 (struck through)
  destination: "DROPPED — refined into M223 Scope.In (the 5 positions ARE 5 real captured HIRING sims)"
  reason_recorded: "M222 BA-6 measured the captured snapshot: 0 job_position rows captured (the '443' was a PROD
    count, never captured); the comparison scoreboard does NOT read job_position (JobSimulation.jobPosition is
    optional/unused). Nothing to replay."
  partial_attempted: no
```

## Repeat-Deferral Patterns
None. DEF-M222-01 is a first-and-only occurrence, created today.

## Fate-1 Investigation

### DEF-M222-01 — "replay directus.job_position rows"
- **Fate-1 (land now, complete) feasible:** N/A — this is a scope **reduction**, not a punt of desired work. The
  render-probe (S3) empirically proved the work is **unnecessary** (0 rows captured; the scoreboard doesn't read the
  entity). There is nothing to land.
- **Fate applied:** **Fate-3** — annotate/attach to the receiving milestone. M223's `overview.md` Scope.In #4 was
  edited (struck through with the full rationale) and its frontmatter `delivers:` line updated to state "NO
  directus.job_position replay, per M222 BA-6/Fate-3." Recorded in M222 `decisions.md` D4.
- **Aging:** not aged out (created 2026-07-15; destination milestone M223 has not yet closed).

## Recommendations
- **DEF-M222-01 → LAND-NEXT (Fate-3, applied):** confirmed. The scope-reduction is already reflected in M223's
  `overview.md`; no further action. This is the single routed item the close report notes.

## Applied Changes
No new edits required by this audit — the Fate-3 was already applied at build time (M223 `overview.md` +
M222 `decisions.md` D4). This report records the audit trail.

## Blocking Items (require user decision)
None.
