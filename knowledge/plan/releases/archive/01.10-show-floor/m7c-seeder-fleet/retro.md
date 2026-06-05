# M7c — Retro (iterative, closed gate-met-over-reachable + waiver)

**Summary:** Built the **seeder fleet** — the believability core of a demo world — across 5 iters, driving
data-DNA coverage **40% → 80%** and meeting **3 of the 4 gate conditions** outright; the 4th (≥90% coverage)
resolved by **waiving the two snapshot-blocked surfaces** (taxonomy, content) per the Re-scope trigger
(user-confirmed). The reachable subset conforms **100% / critical 100%**, the demo identity logs in **200**, the
full ~8,500-row seed runs in **0.69s** (gate: <2 min), and the isolation audit is **clean**.

## Gate Outcome Ledger
| Gate condition | Outcome |
|---|---|
| (a) demo identity logs in → **200** | ✅ **met** — HTTP 200 (`membershipsCount: 1001`) live on demo-1 |
| (b) data-DNA coverage **≥ 90% / critical 100%** | ✅ **met-over-reachable** — 100% over the 8 non-waived surfaces (critical 100%); taxonomy + content **waived** (the hard line) |
| (c) seeding **< 2 min** | ✅ **met** — 0.69s for ~8,500 rows (8 seeders) |
| (d) isolation audit **zero shared/prod writes** | ✅ **met** — clean on every run |

**Waiver (Re-scope trigger, user-confirmed):** `taxonomy` (needs the pre-embedded skiller node-hierarchy
snapshot — empty in demo-1) and `content` (the shared Directus instance — the isolation guard blocks live
writes; snapshot-replay only) are M7c's declared hard line — structural-snapshot / shared-store work, not
seeding. Recorded as `waived-m7c` in the data-DNA (excluded from the coverage denominator); a natural **v1.2**
seed (richer demo worlds).

## What went well
- **The reachability survey (TOK-01) was the whole game.** The key finding — the believability core (sessions/
  assignments/activity) references content via *free values*, not DB-enforced FKs — meant the entire user-
  activity layer was seedable **without** touching the shared Directus. That's what made M7c tractable.
- **One pattern, four surfaces.** jobsim-sessions established the deterministic backdated-activity generator
  (time-distributed over the span, pass/fail per `pass_rate`, no random source); skillpath-sessions, assignments,
  activity reused it. The fleet is ~4 small files.
- **Live seeding kept catching real constraints** unit tests can't: the skillpath UNIQUE `(user_id,
  skill_path_id, version)` (fixed by indexing the path by session number), and the M7b harness's
  Validate-before-introspect bug (a freshly-promoted surface couldn't load). Both fixed + pinned.
- **The perf path delivered:** ~8,500 rows across 5 schemas in **0.69s** — the direct-`COPY` architecture (M7a-D3)
  paid off; the <2-min gate was never in doubt.
- **The data-DNA closed the loop:** every seeder promoted planned→seeded, introspected, and conformance-gated;
  the `waived` status keeps the manifest honest about the hard line.

## What didn't / constraints
- **Two surfaces genuinely out of reach** without a snapshot mechanism — correctly waived rather than faked.
- **Run across 2 sessions** (iter-01/02, then iter-03/04/05 + close) — the right cadence for an iterative
  milestone; each session left a clean boundary.

## Carried forward → M8 / v1.2
- **M8 (recipes):** 2–3 seed presets (200/500/1k) over the now-complete fleet; the express-gate CI carry-forward;
  the discoverability story.
- **v1.2 seed:** the snapshot mechanism — taxonomy (skiller node-hierarchy import) + Directus content
  snapshot-replay — to lift the waived surfaces and reach richer demo worlds.

## Metrics
See [metrics.json](metrics.json). 5 iters, coverage 40%→80% (100% over reachable), **20 seeder tests / 145
module tests**, all gates green; full 8-seeder seed 0.69s; measure/critical 100%; login 200; isolation clean.
2 bugs caught live (skillpath UNIQUE, introspect-load) + the waived-status discipline added.
