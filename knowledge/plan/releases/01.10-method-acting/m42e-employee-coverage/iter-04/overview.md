---
iter: 04
milestone: M42e
iteration_type: tik
status: closed-no-lift
created: 2026-06-25
---

# iter-04 — the 5 empty sim-result pages (highest-leverage residual cluster)

**Type:** tik -- production-fix under TOK-01, coverage-protocol.md Phase A-E. Targets the largest residual
cluster from iter-03's re-sweep.

## Active strategy reference
**TOK-01: sweep-then-route-by-leverage.** Leverage-first: the 5 `/sim/.../result/<uuid>` empties are the
largest cluster (5 of 8 residual failures), all one fix surface (`stack-seeding`), so clearing them is the
biggest single-fix reduction toward `(0,0)`.

## Re-survey (Phase 1 Step 0)
iter-03 re-sweep is the current measured state: `(failing=8, escapes=1)`. The 5 sim-result empties are
untouched + confirmed: nav-discovered from `/profile/activities` (5 `a[href*="/result/"]` links, each with
`?sessionId=<uuid>&organizationId=22222222-2222-2222-2222-222222222222`). Probe-confirmed the result page is
`http=200` with `<main>` present but **`mainTextLen=0`** (an empty result-view shell). Still the dominant cluster.

## Cluster / target identified
The 5 sim **result** pages (`/sim/<slug>/result/<sessionId>`) reached from `/profile/activities`. The sessions
ARE seeded (`JobsimSessionsSeeder` writes `jobsimulation.sessions` -> they list in activities) but the result
page's data (validation results / score breakdown / transcript it reads to fill `<main>`) is NOT seeded -> the
result `<main>` renders empty.

## Hypothesis
The result page reads a result/validation surface keyed by `sessionId` that the session seeder doesn't write.
Seeding that surface (in `stack-seeding`, the routed fix surface) for the activities-listed sessions fills the
result `<main>` -> those 5 pages stop being empty. Phase B maps the exact data contract (which table/query the
result page reads) before writing the seed.

## Expected lift
`failing` 8 -> 3 (clear all 5 sim-result empties) if the result-data seed lands cleanly. Partial (e.g. 8 -> 5)
if only some sessions get result data. The re-sweep delta is the truth.

## Phase plan
- **Phase B -- map the data contract:** determine what the result page queries (GraphQL/RPC -> which
  jobsimulation/validation table keyed by sessionId). Read the next-web result route + the jobsim/validation
  schema (READ-ONLY -- zero platform edits).
- **Phase C -- fix:** extend the seeder (`stack-seeding`, likely `jobsim_sessions.go` / a result/validation
  seeder) to write the result data for the activities sessions. Re-seed the live demo.
- **Phase D -- re-sweep:** re-run the employee sweep; record the new `(failing, escapes)` + the per-cluster delta.
- **Phase E -- close:** grade on whether the 5 sim-result empties cleared; route the rest forward.

## Escalation conditions
- The result page reads a surface that ONLY a platform change can populate (e.g. a computed validation the app
  derives, not a seedable row) -> re-scope-trigger (zero-edit line); record + escalate, do NOT edit the platform.
- The data contract is large/unclear -> land what's seedable (partial), route the remainder forward.

## Acceptable close-no-lift outcomes
- If Phase B reveals these per-session result deep-links are NOT a legitimate employee coverage commitment
  (e.g. they're an internal/manager surface leaking into the employee crawl), the correct fix is crawl-scope
  (exclude them) not seeding -- a falsification that still advances the gate definition. Record + route to a
  crawl-scope tooling-iter.
