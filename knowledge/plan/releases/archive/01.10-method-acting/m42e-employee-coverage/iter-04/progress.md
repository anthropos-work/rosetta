**Type:** tik (production-fix attempt -> falsified) -- under TOK-01, coverage-protocol.md Phase A-E.

# iter-04 progress -- the 5 empty sim-result pages (investigation + falsification)

Targeted the largest residual cluster (5 of 8 failures): the empty `/sim/<slug>/result/<sessionId>` pages.
Phase B falsified the planned `stack-seeding` fix; the correct fix is crawl-scope, routed to a tooling-iter.

## Phase A -- (carried) measured state
iter-03 re-sweep: `(failing=8, escapes=1)` over reachable=80. Target cluster: the 5 sim-result empties.

## Phase B -- map the data contract (the falsification)
Read (READ-ONLY, zero platform edits): the result route `sim/[slug]/result/[sessionId]/page.tsx` ->
`AISimulationResultContainer` -> `useGetResult` (`GET_SIMULATION_RESULT -> jobSimulationResult.evaluationStatus`,
keyed by sessionId). The container guards `!resultData || evaluationStatus !== Pending` and renders an **empty
`<main>`** when there is no result.

**Finding (D1): the result is a RUNTIME-COMPUTED artifact, not a seedable row.** A backdated SEEDED session
(written by `JobsimSessionsSeeder` into `jobsimulation.sessions`, which is why it lists in `/profile/activities`)
has **no `jobSimulationResult`** -- the evaluation is computed server-side by the jobsimulation AI pipeline from
the transcript (cf. the `useRecalculateEvaluationResult` retry). Seeding a believable one means synthesizing the
AI pipeline's whole output (result + interactions + interaction-calls + tasks) across coupled tables AND
satisfying the computed `evaluationStatus` -- brittle, and the only robust way to populate it is to RUN the
evaluation (a platform action, forbidden this release).

## Phase B -> routing (D2/D3/D4)
- **Correct fix surface = crawl-scope, NOT seeding (D2).** Per coverage-protocol "Scope = the vantage's
  reachable set": a presenter lands on `/profile/activities` (renders fine, text=998, lists the sims); a
  specific historical session's computed AI evaluation result is a runtime-only surface a seeded demo can't fill.
  These per-session result deep-links are not a legitimate demo coverage commitment for the employee vantage.
- **Routed to iter-05 (tooling-iter, handler `TOOL-M42e-iter05-result-scope`, D3):** add a `skipPaths` pattern
  (`/sim/.../result/`) so the sweep doesn't enqueue/score per-session result deep-links; re-sweep. Not landed
  this iter (a harness change is a 2nd line -> scope-creep tripwire; the protocol routes harness changes to a
  tooling-iter).
- **Zero-edit line held (D4):** the only in-scope-links path to fill these is a platform change -- NOT taken.
  This is NOT a re-scope-trigger (a viable in-scope fix exists: crawl-scope), just a routed-forward tooling fix.

## Phase C/D -- fix / re-sweep
None this iter (the planned seed was falsified; no fix landed -> no re-sweep). Metric unchanged at `(8,1)`.

## Close -- 2026-06-25

**Outcome:** Investigated the 5 sim-result empties; FALSIFIED the planned `stack-seeding` fix -- the result is a
runtime-computed AI-evaluation artifact, not a seedable row, and the zero-edit line blocks the platform path.
Correct fix = crawl-scope (exclude per-session result deep-links from the reachable set), routed to iter-05
(tooling-iter). Metric unchanged `(8,1)` -- a complete investigation ending in characterization.
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n -- (2) triggered-tok: n -- (3) re-scope: n -- (4) user-blocker: n -- (5) cap-reached: n -- (6) protocol-stop: n -- Outcome: continue
**Decisions:** D1 (runtime-computed, seed falsified), D2 (crawl-scope is the fix), D3 (route to iter-05 tooling-iter), D4 (zero-edit line held) -- see ./decisions.md
**Side-deliverables (if any):** none
**Routes carried forward:**
- iter-05 (tooling-iter, Fate-3, handler `TOOL-M42e-iter05-result-scope`): add the `skipPaths` `/sim/.../result/`
  exclusion to the employee sweep + re-sweep. (Expected to drop failing 8 -> 3.)
- iter >= 05 (Fate-3, handler `SEED-M42e-empty-skillpaths`): the 2 empty skill-path indexes
  (`prompt-engineering-fundamentals`, `foundation-of-artificial-intelligence`) -- a real seedable/serve-grant
  content gap (still production-fix surface).
- iter >= 05 (Fate-3, handler `ASSERT-M42e-root-sentinel`): the `/` sentinel false-positive (visible-main-region
  assertion / root normalization).
- iter >= 05 (Fate-3, handler `ESCAPE-M42e-skillpath-external-article`): the `strategy-business.com` escape.
**Lessons:** Not every empty page is a seed gap -- a RUNTIME-COMPUTED surface (an AI evaluation result keyed by
sessionId) cannot be seeded structurally; under the zero-edit line the correct resolution is crawl-scope (the
gate is over the vantage's MEANINGFULLY-reachable set, and a seeded demo's per-session result deep-link is not
that). This distinction (seedable structural row vs runtime-computed artifact) generalizes to the manager
vantage (M42m) -- a candidate measurement-conventions note for coverage-protocol.md when iter-05 lands the scope rule.
