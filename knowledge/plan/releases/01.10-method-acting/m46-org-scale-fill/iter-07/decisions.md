# M46 · iter-07 — decisions

## D1 — The 998→~500 population-math fix is CONTAINED in the descriptor, not the seeder
The 998-member double-size bug (iter-07's first re-seed) is: the curated `UsersSeeder` seeds a full `size`
synthetic body AND the `fill:true` batch adds `size − heroes` generated members → the org lands at ~2×`size`.
The contained fix keeps the descriptor's `size` at 250 (Cervato) / 120 (Solvantis) so the org lands at a
believable ~500 (Cervato 498) / ~237 (Solvantis) — half curated, half generated. The `StoryHasFillBatch`
helper was added to `blueprint/batch.go` documenting the heroes-only intent, but the realized fix is the
descriptor size, not a seeder rewrite (the heroes-only-UsersSeeder refactor is the named **re-scope** target,
deferred — see D3, and it would not have helped the grid wall anyway).

## D2 — The warm-grid harness fix (bounded re-assert POLL + vantage-aware warm) LANDED but did NOT close the gate
`section-assert.ts` got a bounded re-assert poll (up to 6×, return on first pass, poll only a paint-timing
`skeleton`/`empty` kind, no false-pass) and `coverage.spec.ts` got a manager-vantage warm of
`/enterprise/{members,activity-dashboard,workforce}`. Both are correct, defensible org-scale extensions of the
iter-09 settle lesson and are kept. But they did not close the gate — the failure is not slow PAINT, it is a
never-resolving QUERY (below).

## D3 — RE-SCOPE-TRIGGER: the 5th gate face is platform-bound, not seedable/harness/resize-fixable
**Decision: surface a re-scope-trigger; do NOT fake the gate.**

The manager M42 sweep on the corrected ~500-member org returned `failingSections=3` (the SAME 3:
`/enterprise/members`, `/enterprise/activity-dashboard`, `/enterprise/settings`), `personaFailures=0`,
`escapes=0`, `crossPortFollowFails=0`, persona/cross-port all `ok`. The decisive evidence:

1. **Not empty, never-resolving.** The three screenshots show full chrome + Cervato org + table headers but a
   `…` loading spinner / skeleton rows + a `0 / ∞` count — a GraphQL query that never resolves in the window,
   NOT a content gap. The data IS present: 498 Cervato memberships + 1000 activity_events seeded; raw SQL for
   the same shape is **31–121 ms** (fast).
2. **The slowness is in the FEDERATED GraphQL layer.** Cosmo-router latency in the sweep window: **11 requests
   at 10 s+**, peaking **83.9 s / 80.4 s / 60.0 s / 16.5 s**. `organizationMembers` + `membershipsCount` + the
   activity-dashboard aggregation fan out per-row resolvers (`jobRole`, `targetRole`, `tags`,
   `lastActivityDate`, `organizationFeatures`) — each a subgraph round-trip + a Sentinel authz check — an N+1
   across the federation. The router `/health` is 66 ms (the router is fine; the heavy queries are not).
3. **Resize did not help.** 998→500 barely moved the query time (10.88 s @ 998 → 10.5 s @ 500) ⇒ member-COUNT
   is not the dominant factor; per-row fan-out is. The contained levers (resize + a ~12 s warm-grid poll) are
   exhausted and cannot catch an 84 s query.
4. **The manager gate last PASSED at 221 members** (commit `63d3fad`, M42m/04). M46 "org-scale fill" inherently
   pushes past 221, colliding with the platform grid wall.
5. **No legitimate contained fix remains.** A platform resolver fix is forbidden (zero-canonical-edit line). A
   `demopatch` of the federation N+1 is out of scope (the demo-patch mechanism is for content-anchored
   believability hunks, not a multi-file resolver perf refactor) and would be faking the gate. Resizing back to
   ~221 would pass but defeats the milestone (221 is the pre-M46 baseline, not org-scale).

**Conclusion:** the exit-gate criterion *"the M42 semantic gate PASSES on a ~500-member generated org"* is
**unsatisfiable against the unmodifiable platform** — the enterprise org grids do not hydrate org-scale data in
any reasonable window. Even the named re-scope refactor (heroes-only-`UsersSeeder` + fill-aware downstream
seeders) only fixes the population MATH (org = `size`, not 2×`size`); it does not touch the grid wall. The
re-scope is therefore deeper than population math: it is a **gate-criterion re-scope** — the M46 gate must
either (a) measure org-scale believability on the surfaces that DO render at scale (employee profile, the
seeded population via DB/closure, the workforce-intelligence aggregates if they paginate), and treat the
enterprise members/activity grids as a documented platform-perf exception at org scale; or (b) cap the
"headline" org at the platform's render threshold and prove org-scale fill on the generated population's
correctness rather than the enterprise-grid render. This is a roadmap-owner decision, surfaced — not faked.

## Escalation
Per the iter-07 overview's escalation conditions: *"An unstabilizable believability gap that is NOT seedable /
fixable in-rext, and is platform-bound → re-scope-trigger (the coverage-protocol zero-edit line)."* This is
that condition, proven empirically.
