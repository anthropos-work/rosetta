---
iter: 05
milestone: M42e
iteration_type: tik
iter_shape: tooling
status: closed-fixed
created: 2026-06-25
---

# iter-05 — tooling-iter: crawl-scope the per-session sim-result deep-links

**Type:** tik / **iter_shape: tooling** (coverage-protocol.md "Iter type selection -> Tooling-iter"). iter-04
falsified seeding the 5 sim-result empties + routed the correct fix here: a harness `skipPaths` rule. Ships the
capability AND uses it (re-sweep) in the same iter.

## Active strategy reference
**TOK-01: sweep-then-route-by-leverage.** Highest-leverage residual move: one harness skip-rule clears the 5
sim-result empties (5 of 8 failures) -> expected `failing` 8 -> 3 in a single change.

## Re-survey (Phase 1 Step 0)
iter-03 re-sweep is the current measured state `(failing=8, escapes=1)`; iter-04 was no-lift (no metric move).
The 5 sim-result empties are still present + confirmed runtime-computed-not-seedable (iter-04 D1). The
crawl-scope fix (the `skipPaths` rule) is still untouched in committed form.

## Cluster / target identified
The crawl's `skipPaths` set (`stack-verify/e2e/lib/crawl.ts` DEFAULT_SKIP + the spec's per-vantage `skipPaths`).
Add a pattern that drops per-session sim/skill-path RESULT deep-links (`/sim/.../result/` and any
`/skill-path/.../result/`) from the BFS frontier so they are not enqueued/scored -- they are runtime-computed
surfaces outside a seeded demo's meaningfully-reachable set (iter-04 D2).

## Hypothesis
Adding the `skipPaths` rule removes the 5 `/sim/.../result/<uuid>` pages from the reachable set; the re-sweep
drops `failing` 8 -> 3 (the residual: 2 empty skill-paths + 1 `/` sentinel false-positive). escapes still 1.

## Expected lift
`failing` 8 -> 3. The 3 remaining = the 2 empty skill-paths (real seedable content gap) + the `/` false-positive
(assertion-tune) -- both routed forward to the next tik(s).

## Phase plan (multi-step tooling-iter -- planned 2 lines)
- **Line 1 (ship the capability):** add the result-deep-link `skipPaths` pattern to the employee sweep
  (`coverage.spec.ts` skipPaths, or DEFAULT_SKIP in crawl.ts). Compile-validate.
- **Line 2 (use it):** Phase D re-sweep vs live demo-3; record the new `(failing, escapes)` + delta.
- **Phase E -- close:** grade on whether the 5 sim-result pages left the reachable set + failing dropped to ~3.
- Fold the seedable-row-vs-runtime-computed + the result-scope rule into coverage-protocol.md (protocol-evolution).

## Escalation conditions
- The skip-rule over-excludes (drops a legitimate page) -> tighten the pattern; re-sweep. Not a blocker.
- The re-sweep shows the result pages were ALSO reached by a non-result path -> investigate; route forward.

## Acceptable close-no-lift outcomes
- If the skip-rule lands but the re-sweep surfaces the result pages are still enqueued via a different link
  shape, close on the landed rule + route the residual; the scoping capability still advances the harness.
