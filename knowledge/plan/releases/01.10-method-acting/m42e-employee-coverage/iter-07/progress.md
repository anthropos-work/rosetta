**Type:** tik / iter_shape: tooling -- under TOK-01, coverage-protocol.md "Iter type selection -> Tooling-iter".

# iter-07 progress -- tooling-iter: tighten the harness, exhaust the frontier, quote the TRUE residual

Run-2 opener. The run-1 verification (`.agentspace/scratch/work-m42e/verify-run1.md`, YELLOW) proved the
committed `(failing=2, escapes=3)` is HONEST but is a FLOOR over a CAP-SATURATED slice (`reachable===80===cap`).
This tooling-iter ships the 4 mandated harness corrections, then raises the cap until the BFS frontier
EXHAUSTS and re-sweeps to quote the REAL residual (the precondition for triaging it), and DIAGNOSES the empty
pages to exact root cause before any fix (must-fix 2).

## Line 1 -- ship the harness corrections (must-fix 3 + 4)
1. **`coverage.spec.ts`: tighten the `/result` skip regex** (D1) -- `/\/result\b/i` -> segment-anchored
   `[/\/result\/[0-9a-f-]{6,}/i, /\/result\/[^/]+$/i]`. Removes the latent over-skip of any future path merely
   containing "result"; still skips the runtime result deep-links. Zero real pages changed.
2. **`crawl.ts` + `coverage.spec.ts`: cap-hit flag + annotation** (D2) -- added `cappedAtFrontier`,
   `frontierRemaining`, `maxPages` to `CoverageReport`; a runtime CAP-HIT warning; annotated `reachable` so a
   cap-saturated `reachable===maxPages` can never again be misread as a true count. A `(0,0)` over a truncated
   frontier is now structurally flagged as NOT gate-met.
Compile-clean via `playwright test --list`.

## Line 2 -- use it: the raised-cap re-sweep (must-fix 1)
- **Re-sweep at `COVERAGE_MAX_PAGES=150`:** `reachable=87, cappedAtFrontier=FALSE, frontierRemaining=0` -- the
  BFS frontier EXHAUSTED at 87 pages (42 skill-paths + 39 sims + 3 profile + 2 library + 1 home), well under
  the 150 cap. **The TRUE, frontier-exhausted residual = `(failing=3, escapes=3)`** (D3). The verdict's
  ~269-unreached-sims estimate was wrong: the employee/`maya-thriving` vantage's in-app nav LINKS only 39 sim
  pages (the 307 sims exist in the library but aren't all crawl-reachable nav links from this vantage). The
  floor `(2,3)` rises to `(3,3)` -- one NEW failure surfaced by the deeper crawl (a sim-start page).

## Line 3 -- DIAGNOSE the residual (must-fix 2 + 5 prep) -- routed forward
- **2 empty skill-paths (D4 -- ROOT CAUSE FOUND; `directus_versions` hypothesis FALSIFIED):**
  `prompt-engineering-fundamentals` + `foundation-of-artificial-intelligence` render a perpetual loading
  spinner. The `getSkillPath` GraphQL query returns a federation error -- skiller `skill not found
  K-AIFUNX-E658` on the non-nullable `chapters.jobSimulations.simulation.skills.name` -> the whole query nulls
  client-side -> empty page. NOT a directus_versions row gap. The cause is a **STALE TAXONOMY CACHE** (D5): prod
  has 42790 public skills incl `K-AIFUNX-E658`; the cache (captured 2026-06-09) has 42768 and is missing 22.
  -> iter-08 taxonomy RE-CAPTURE (in-rext, zero-platform; handler `SNAP-M42e-stale-taxonomy-recapture`).
- **1 new sim-start empty (D6):** `/sim/.../start` is a RUNTIME simulation-launch surface (empty `<main>`, only
  a successful `jobSimulationBySlug`, client-populated only on a live session) -- crawl-scope class like the
  `/result/` pages (iter-04/05). -> iter-08 `skipPaths` rule (handler `SCOPE-M42e-sim-start-runtime`).
- **3 escapes (D7 -- FULL set confirmed, frontier-exhausted):** strategy-business.com + en.wikipedia.org +
  dremio.com, all editorial citations in replayed `/skill-path/.../chapter` body copy (pages render full +
  pass). Set complete at 3 (no new ones beyond the old cap). -> iter-08 ALLOW-RULE + presenter-notes (per user
  pre-authorization; handler `ESCAPE-M42e-skillpath-external-articles`). Do NOT strip the citations.

## Phase E -- Close

**Outcome:** Tooling-iter corrected the MEASUREMENT (regex tighten + cap-hit flag) and RE-BASELINED the metric
to the frontier-EXHAUSTED true residual `(failing=3, escapes=3)` (replacing the cap-saturated floor `(2,3)`),
then DIAGNOSED the residual to exact root causes BEFORE any fix -- the `directus_versions` hypothesis is
FALSIFIED by direct evidence; the 2 empty skill-paths are a stale-taxonomy-cache gap (missing public skill
`K-AIFUNX-E658`), the new sim-start is a runtime surface (crawl-scope), and the full escape set is exactly 3
editorial citations. The cap-exhaustion + cap-hit lessons folded into coverage-protocol.md.
**Type:** tik (iter_shape: tooling)
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n -- (2) triggered-tok: n -- (3) re-scope: n -- (4) user-blocker: n -- (5) cap-reached: n (tik #1 of session) -- (6) protocol-stop: n -- Outcome: continue
**Decisions:** D1 (regex tighten), D2 (cap-hit flag), D3 (frontier exhausted @87, true residual 3/3), D4 (root cause: skiller federation error, directus_versions falsified), D5 (stale taxonomy cache), D6 (sim-start runtime surface), D7 (full escape set = 3 citations) -- see ./decisions.md
**Side-deliverables (if any):** a reusable `tests/probe-empty.spec.ts` diagnostic probe (network/DOM/GraphQL capture) -- the tooling-iter's diagnostic capability; kept for future empty-page triage.
**Routes carried forward (iter-08):**
- `SNAP-M42e-stale-taxonomy-recapture` (Fate-3, iter-08): re-capture the taxonomy surface (picks up the 22
  missing public skills incl `K-AIFUNX-E658`) + cache swap + re-replay into demo-3 -> the 2 empty skill-paths
  render. The highest-leverage production fix (expected failing 3 -> 1).
- `SCOPE-M42e-sim-start-runtime` (Fate-3, iter-08): a `skipPaths` rule for `/sim/.../start` (runtime launch
  surface) -> failing 1 -> 0 after the recapture.
- `ESCAPE-M42e-skillpath-external-articles` (Fate-3, iter-08): the citation allow-rule + presenter-notes list
  (full set now known = 3) -> escapes 3 -> 0.
**Lessons:** raise the cap until `cappedAtFrontier===false` BEFORE reading the residual -- a cap-saturated
`reachable===maxPages` is a FLOOR, not a count (folded into coverage-protocol.md as the new cap-exhaustion
convention). DIAGNOSE an empty page via a DOM + network + downstream-service-log probe before assuming a fix
surface -- a "content not replayed" guess was wrong; the real cause was a non-nullable GraphQL field nulling
the whole query on ONE missing skiller skill (a stale public-taxonomy cache). The federation-error-nulls-the-
page pattern + the probe technique fold into coverage-protocol.md.
