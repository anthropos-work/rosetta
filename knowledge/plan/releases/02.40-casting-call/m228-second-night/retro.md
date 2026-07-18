# M228 "Second night" — Retro

## Summary
The corrected (M227) hiring demo re-proven LIVE on `billion` over the tailnet — the 7-condition gate (retuned for
M227) MET 7/7 on a cold-seed instance, all four believability corrections confirmed rendering. Closed-on-gate.

## Incidents this cycle
- **The M227 fix-#1 guard was incomplete (F1/F2/F3).** The live cold bring-up (iter-02) surfaced what the
  deterministic M227 unit test omitted: FeedbackSeeder (mirror rows since M42m → training-sim leak + 2nd session
  per candidate) and SuccessionSeeder (FK to a now-skipped session → seed "failed") also needed the guard. Fixed
  iter-03, regression enumerates all 8 generic seeders. **This is the milestone's headline value: live-prove
  caught a gap deterministic tests structurally could not.**
- **C2 render "failed" 4 times before the real cause was found.** The recruiter drawer is a Next.js intercepting
  route — only the first sim per page-load is cleanly detectable. Cost several long cold-tailnet render cycles of
  blind hypothesis-testing (visible-drawer scoping, re-land-per-sim) before reading the trace/screenshots proved
  the drawers render fine and the fix was "prove each sim as the first" (RENDER_ONLY_SIM). No production impact —
  a test-harness limitation, not a demo defect.

## What went well
- Multi-source verification (app list API + per-sim screenshots + DB) established data correctness independently
  of the flaky harness, so the harness issue never blocked the truth.
- The enumerated `TestGenericActivitySeeders_SkipHiringOrg` table made the F1/F2/F3 fix a one-line-per-seeder add.
- p95 click→ACCESS 1.27 s — the recruiter vantage is fast; ACCESS ≠ the slow insights data render.

## What didn't
- Blind render re-runs over a cold tailnet burned wall-clock. Lesson: read the trace/screenshots BEFORE the third
  identical failure, not after — the DOM snapshot named the cause immediately.
- The demo re-cooled between runs (idle → ~15 min cold SSR each), amplifying the cost of each iteration.

## Carried forward
- None milestone-specific (closed-on-gate). Release-level carries (8 pre-existing demo-stack test failures, M204
  assign-WRITE TODO, DEF-M226-01 pre-bind serve reap) remain routed to `/developer-kit:close-release`.

## Metrics delta
- rext `stack-seeding/seeders` coverage 96.8%; flake 0; platform-repo edits 0. Live: render 5/5, heroes 3/3,
  p95 1.27 s, 42 hiring sessions / 0 leak. See `metrics.json`.
