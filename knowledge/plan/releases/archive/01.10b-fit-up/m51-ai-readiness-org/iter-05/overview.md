---
iter: iter-05
milestone: M51
iteration_type: tik
iter_shape: tooling
status: closed-fixed-partial
created: 2026-06-30
---

# iter-05 — tooling-iter (deepen the manager-grid WARM so the cold ~11.6s members query resolves pre-sweep)

## Type
tooling-iter (coverage-protocol.md "Iter type selection" refinement) — under TOK-01 (active-cycle
signals-true). iter-04 closed-no-lift BLOCKED on a harness measurement gap: the manager warm step bails at the
4s networkidle ceiling, so the org-scale members grid (cold ~11.6s query) is still loading at the authoritative
visit → a skeleton-frame false-fail. This iter SHIPS the harness capability (a content-presence warm for the
heavy manager grids) AND USES it within the same iter for the coverage drive. coverage-protocol Phase A–E.

## Step 0 — Re-survey before targeting
Re-confirmed against the iter-04 GATED sweep (`(failing=6, escapes=0)` frontier-exhausted) + the live state:
- demo-1 is UP at fit-up-m50 with all 3 perf demo-patches baked + the AI-readiness showcase org seeded (verified
  in the DB: Northwind 200, ENABLED, 78.4% all-3, heroes pinned). No re-up / re-seed needed — only the harness
  changes + a re-sweep.
- The 6 fails are all `kind:empty` "re-asserted 6× — genuinely below bar" on the base-Workforce grids; the
  screenshot is real-chrome-over-skeleton; the GraphQL latency is ~11.6s (slow-not-erroring). The DB has the data.
Target (the harness warm/poll) is untouched + meaningful. No substitution.

## Active strategy reference
**TOK-01** (`../decisions.md`) — active-cycle signals-true. This iter is the coverage-drive strand (TOK-01
step 4): the fix surface is the HARNESS (the demo-UP perf-patches already landed in iter-04 and reduced the
wall; the residual is a measurement-budget gap), per the coverage-protocol "Org-scale grid perf wall" +
"bound the settle to the heaviest DATA GRID" lessons.

## Cluster / target identified
The single highest-leverage cluster: ALL 6 failing sections share one root cause — the org-scale members/grid
query (~11.6s cold) outruns the harness warm (4s ceiling) so the authoritative visit pays the cold cost and
captures a skeleton. ONE harness fix — a content-presence WARM for the manager's heavy `/enterprise/*` grids
(navigate, then wait for REAL ROWS up to a generous ceiling, not just networkidle) — primes the GraphQL/React
caches so every subsequent authoritative visit reads a hydrated grid. (The members-roster + members-location +
assign-roster + activity-table + the two workforce sections are all the same backing-query family.)

## Hypothesis
A deepened warm that waits for real rows (content-presence) on the manager's heavy grids
(`/enterprise/members`, `/enterprise/activity-dashboard`, `/enterprise/workforce`, `/enterprise/assignments`)
with a generous ceiling (~25s, > the observed 11.6s) — instead of bailing at the 4s networkidle ceiling — lets
the cold query resolve + the result cache during warm, so the authoritative visit + the existing bounded
re-assert poll read hydrated grids → the 6 skeleton false-fails flip to real rows → `(failing→0, escapes=0)`.
(A secondary lever if warm alone is insufficient: widen the per-section heavy-grid poll for the org-scale grids.)

## Expected lift
A net reduction of the 6 perf-wall failing sections toward 0. The data is confirmed present + the wall is
already 76.4s→11.6s; the residual is purely the warm not reaching the cold query, so a content-presence warm
should clear all 6 (they share the backing query family). Success signal: the re-sweep `(failing)` drops.

## Phase plan
- Phase A (sweep): inherit the iter-04 GATED `(6,0)` as the pre-iter metric.
- Phase B (triage): done in Step-0 — all 6 = the harness-warm gap, routed to the harness fix surface.
- Phase C (fix): deepen `warmStack` (or add a manager-grid `warmHeavyGrids`) in
  `stack-verify/e2e/lib/section-assert.ts` to wait for real rows (content-presence) with a generous ceiling on
  the manager's heavy grids; wire it into `tests/coverage.spec.ts`'s warm step. Run from the AUTHORING copy
  against the LIVE demo-1 (harness-only change — no rebuild/re-seed needed; the harness runs against the live
  offset ports).
- Phase D (re-sweep): GATED manager-vantage sweep on demo-1 → record `(failing, escapes)` delta.
- Phase E (close): grade on whether the perf-wall sections cleared; route the manifest AI-readiness assertion +
  cockpit jump_to (TOK-01 strand-4) forward if not yet landed.

## Escalation conditions
- If a content-presence warm + a widened poll STILL can't get the cold query under any reasonable budget (the
  query is genuinely pathological at 200 members, not just cold) → re-triage: is a further demo-local query
  optimization available, or is this a documented residual to escalate as a re-scope-trigger? (NOT expected —
  the M46 close cleared the analogous grids demo-local.)
- If the harness change introduces a regression in the employee-vantage sweep (the warm is shared) → user-blocker.

## Acceptable close-no-lift outcomes
If the deepened warm lands clean but a residual perf gap keeps `(failing)` above 0 → close-fixed-partial (the
sections that cleared) or close-no-lift with the falsification "a content-presence warm cleared N of 6; the
residual M need a further lever". A documented characterization of exactly which grids resist warming is a valid
outcome that sets up the next iter.
