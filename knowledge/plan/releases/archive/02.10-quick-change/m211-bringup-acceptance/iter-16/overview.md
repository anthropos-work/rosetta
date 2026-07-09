---
iter: 16
milestone: M211
iteration_type: tik
status: planned
created: 2026-07-08
---

# iter-16 — tik: v2.0 Playthroughs suite GREEN on the merged demo-1

## Active strategy reference
**TOK-01** move (4) — close the **Playthroughs** half of sub-condition (e) on demo-1 (function-proof, after
iter-14/15 closed the coverage/presence half).

## Step 0 — Re-survey
Both M42 coverage vantages are GREEN (iter-14 employee, iter-15 manager). The remaining (e) piece is the v2.0
Playthroughs suite (10 live: 6 employee + 4 manager, + 1 declared in-manifest TODO — the assign-WRITE half →
Fate-2). Shipped GREEN in v2.0 (M202→M204) on cold reset-to-seed; M211 re-proves them on the MERGED platform.

## Cluster / target identified
Run `playthroughs/e2e/run-playthroughs.sh 1 --reset` against demo-1: reset-to-seed the dedicated `pt-world`
(the real `stackseed --reset` FK-ordered TRUNCATE + fresh `pt-world.seed.yaml` seed + sentinel reload), run
the 10 Playthroughs serially (workers=1), reconcile inline via `ptreport --gate no-regressions`. Gate-run
prereq (M204 iter-05 D1): prepend demo-1's pinned `bin/` to PATH so the runner's bare `stackseed` resolves.
The harness runs from the authoring copy (HEAD 84e15e9 — all M209/M211 fixes).

## Hypothesis
The merged platform + the re-grounded pinned tooling (post-M209 stackseed resolves public.* node-ids; closure
green for the stories seed already) → the `pt-world` reset-to-seed succeeds + the 10 Playthroughs pass green
(NoRegressions: nothing `failing`; the 1 declared TODO stays TODO). Risk: the next-web v2.106.1 bump may have
drifted a page-object locator; if so, re-pin the affected locator (semantic-by-default + find-only landmark
registry → re-pin is O(surfaces)), a Fate-1 harness fix.

## Expected lift
v2.0 Playthroughs: `ptreport --gate no-regressions` **PASS** (10 passing + 1 TODO, 0 failing) on cold
reset-to-seed → the Playthroughs half of sub-condition (e) MET. With coverage already GREEN, sub-condition
(e) COMPLETE.

## Phase plan
1. Pre-flight: `ptvalidate` static checks (both-way id integrity + precondition-coverage) — fast.
2. Run `run-playthroughs.sh 1 --reset` (reap-safe; the reset-to-seed + serial run + inline ptreport).
3. Triage any `failing` Playthrough: locator drift (re-pin, Fate-1) vs seed/platform drift (diagnose) vs a
   genuine platform-bound wall (`unimplementable-without-platform-edit` escalation).
4. Re-measure: ptreport four-state map → gate no-regressions PASS.

## Escalation conditions
- A Playthrough surface needs a platform-repo edit to drive → `unimplementable-without-platform-edit`
  escalation (declared in report/unimplementable.yaml with rationale; never edit the platform).
- The pt-world reset-to-seed fails closure → diagnose (a seeder re-ground gap, Fate-1) or user-blocker.

## Acceptable close-no-lift outcomes
If a locator drift needs a re-pin that lands + the suite goes green → closed-fixed. If a genuine
platform-bound wall surfaces on 1 Playthrough (escalated), the other 9 + coverage still advance the gate →
closed-fixed-partial with the escalation recorded.
