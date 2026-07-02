# M204 Retro — Manager-vantage coverage

## Summary
The FINAL v2.0 "opening night" milestone. Took the **manager vantage** to green as Playthroughs on the M202
foundation: all 3 declared manager journeys (Workforce funnel + member roster, per-member activity-dashboard
drill-down, succession/at-risk) — 4 use cases — pass on a cold reset-to-seed demo, with 0 false-fails over 5
consecutive reset runs. Iterative, **closed-on-gate** at iter-05 (1 bootstrap tok + 4 tiks, all closed-fixed).
Tooling + docs only — zero platform-repo edits, zero new deps. The corpus now stands at **10 live Playthroughs
(6 employee + 4 manager), 1 declared TODO** (the assign-WRITE half). rext code-of-record: tag `opening-night-m204`
@ `c81c6dd`.

## Incidents This Cycle
None. No P2 flakes, no regressions. The one iter-loop side-fix (iter-02 runner reporter-override, a stale-JSON
hazard) was caught + fixed in-iter and pinned by a drift guard at harden Pass-1; the iter-03 SPA-URL race +
out-of-`<main>` table scope were fixed in-iter (D1/D2). Flake gate at close: 5/5 clean (Go + TS unit).

## What Went Well
- **The M203→M204 loop transferred cleanly.** The measure→declare→page-object→play→diagnose→re-measure loop that
  took the employee vantage to green worked verbatim for the manager vantage on the same M202 foundation. The
  shared page-object/locator layer was a genuinely additive merge (each vantage adds its own surfaces; no collision).
- **All 4 manager UCs are READ/monitoring flows** — the pre-flight prediction held: the risk was seed-scale
  render + antd grid ambiguity, not mutation-determinism. Org A (Meridian Labs, size 40) rendered the M36
  org-dashboard aggregates without demo-patches; the 5-run gate was a determinism formality that passed 5/5.
- **The deferral was honest from the start.** The assign-WRITE UC1 was declared out-of-scope at M201 and kept as
  a tracked `playthrough: TODO` build-reference gap — surfaced as `unimplemented` (a first-class state), never a
  silent drop, and presence-pinned by a harden test so it can't vanish unnoticed.

## What Didn't
- **A close code-review finding read a kept-symmetric-API as dead code.** The 5 manager `isOn*` predicate
  functions have no page-object consumer — but the review missed that the M203 `isOnSkillsTab`/`looksLikeTimelineEntry`
  are equally unconsumed and were deliberately kept as part of the single-source-pinned predicate API. Resolved by
  documenting the intent (not pruning), but it cost a round-trip to confirm the M203 precedent. Lesson: when a
  symmetric API has intentionally-unconsumed members, say so at authoring time.

## Carried Forward
- **`assignment-monitoring.assign-and-track.UC1` (the assign-WRITE half)** — a two-backend org-admin WRITE flow →
  **Fate-2**, tracked in-manifest as `playthrough: TODO` (D-CLOSE-1). No current v2.0 milestone owns the
  manager-WRITE class; a future manager-write tier is the natural home. The manifest TODO is the ready-made
  declaration to build against. Not a repeat-defer.

## Metrics Delta
(from `metrics.json`)
- **Gate:** MET — 4/4 manager UCs, 5/5 cold reset-to-seed runs, 0 false-fails (distance 0).
- **Live Playthroughs:** +4 manager (`pt-workforce-funnel`, `-roster`, `-succession`, `pt-activity-drilldown`) → corpus total 10, 1 TODO.
- **Go tests:** 105 test+fuzz funcs (+2 vs M203's 103 — the 2 harden pins). vet + gofmt clean, shuffle 5/5.
- **TS unit specs:** 58 (+20 vs M203's 38). tsc clean, unit flake 5/5.
- **Coverage:** manifest 100.0% stmts (0.0% delta across all 3 harden passes — invariant-pinning, not %-gains).
- **flake_count:** 0.
- **Close review:** 5 findings, all Fate-1 (1 code-quality should-fix + 3 docs + 1 decision-triage; 0 must-fix, 0 escape-hatch).
- **Deferral audit:** GREEN.
