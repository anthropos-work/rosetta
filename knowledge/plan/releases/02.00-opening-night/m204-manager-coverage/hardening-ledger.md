# Hardening Ledger — M204 Manager-vantage coverage

## Pass 1 — 2026-07-02 — final

**Iters hardened this pass:** all milestone-touched code (--final; cumulative-scope sweep across iter-01 … iter-05)
**Tiks covered since prior pass:** all iters in milestone (first harden pass; no prior ledger entry)
**Coverage delta on touched files:**
- `playthroughs/manifest/` (Go): 100.0% -> 100.0% stmts (already saturated by the iter loop's per-iter tests; harden value here is untested-INVARIANT pinning, not statement %)
- `playthroughs/e2e/tests/url-shapes.unit.spec.ts` (TS pure-logic): 40 -> 46 tests (+6 manager edge/interaction/fuzz); the manager predicates in `lib/url-shapes.ts` are pure regex `.test()` calls — the delegating-predicate agreement pins + the new edge/fuzz grid exercise every manager pattern's match AND reject arms
**Tests added:**
- all iters -> `playthroughs/manifest/runner_safety_test.go`: 1 drift guard (`TestRunnerSafety_ReporterOverrideGate`) + 1 test helper (`shellCodeLines`)
- all iters -> `playthroughs/e2e/tests/url-shapes.unit.spec.ts`: 6 edge/interaction/fuzz (1 parent-leaf overlap, 1 drill-down segment-count boundary, 1 empty/degenerate-input, 1 mutual-exclusion matrix, 1 path-relative-vs-host-anchored, 1 adversarial-input fuzz)
**Bugs surfaced + fixed inline:** none — the M204 manager surfaces were correct; the gap was an UNPINNED invariant (the iter-02 reporter-override fix had no drift guard), now closed
**Flakes stabilized:** none
**Cross-iter integration findings:** the reporter-override fix (iter-02) is a FOUNDATION-level invariant spanning all iters (every Playthrough's four-state gate depends on the config's json reporter writing a fresh `last-run.json`); it had no drift guard. Pinned via `TestRunnerSafety_ReporterOverrideGate`, verified fail-red when `--reporter=list` is reintroduced on an executable line. Mirrors the M203 `TestRunnerSafety_SentinelReloadGate` cross-iter pin.
**Knowledge backfill:** none needed — the reporter-override rationale is already documented inline in `run-playthroughs.sh` (the M204 iter-02 NB comment) and is now ENFORCED by the drift guard (assertion 3 pins the rationale comment's presence).
**Stop condition:** continue-to-next-pass — one pass cannot compute the "delta < 2% across MULTIPLE passes" stabilization condition; a second pass is required to confirm the dimension scan finds nothing new.

## Pass 2 — 2026-07-02 — final

**Iters hardened this pass:** all milestone-touched code (--final; cumulative-scope re-sweep, dimension scan across all 6 dimensions)
**Tiks covered since prior pass:** all iters in milestone (same cumulative --final scope)
**Coverage delta on touched files:**
- `playthroughs/manifest/` (Go): 100.0% -> 100.0% stmts (delta 0.0%, < 2%) — the new pin is test-side over already-covered production code
**Tests added:**
- all iters -> `playthroughs/manifest/corpus_test.go`: 1 deliverable-presence pin (`TestRealCorpus_ManagerCoverageIsPresent`) + the `strings` import
**Bugs surfaced + fixed inline:** none
**Flakes stabilized:** none
**Cross-iter integration findings:** Dimension-1 orthogonal gap — the real-corpus test proves INTERNAL consistency but not deliverable EXISTENCE, so a silently-dropped manager manifest / renamed playthrough id would pass with the manager coverage gone. Closed with an external presence pin (the four M204 manager PT ids declared non-TODO, every manager UC played by pt-manager, the assign-write UC1 gap kept as a declared TODO). Verified fail-red on a renamed manager PT id. Dimensions 2/3 (edge/error on manifest TODO/EffectiveOutcome) already covered by manifest_test + manifest_edge_test; Dimension 5 (fuzz) already covers the YAML loader incl. a TODO seed (the manager manifests parse through the same path); Dimension 6 (perf) N/A — page objects are locators.
**Knowledge backfill:** none needed — the deliverable scope is documented in the manager manifests' header prose (assignment-monitoring.yaml records UC1 as a build-reference gap); the pin now enforces it.
**Stop condition:** continue-to-next-pass — this pass's dimension scan SURFACED a new gap (the deliverable-presence pin) and closed it, so "the scan found nothing new" is not yet true for this pass; a confirming Pass 3 is required to grade stabilized.

## Pass 3 — 2026-07-02 — final

**Iters hardened this pass:** all milestone-touched code (--final; confirming clean-scan re-sweep across all 6 dimensions)
**Tiks covered since prior pass:** all iters in milestone (same cumulative --final scope)
**Coverage delta on touched files:**
- `playthroughs/manifest/` (Go): 100.0% -> 100.0% stmts (delta 0.0%, < 2%, stable across all 3 passes)
**Tests added:** none — confirming pass; the dimension scan surfaced no new gap (legitimately a zero-test pass, per the harden anti-pattern guard: run the confirming pass even when the prior looked clean, to grade stabilized against a real re-measurement)
**Bugs surfaced + fixed inline:** none
**Flakes stabilized:** none — flake gate ran 3 consecutive clean runs of the newly-added tests (Go drift guards + TS pure-logic unit spec): 3/3 PASS both stacks
**Cross-iter integration findings:** clean re-scan. Dim1 — no PURE (non-live-Page) logic remains in the manager page objects (every method is a Locator return delegating to base helpers; the one async-nav method drillIntoActiveContent delegates to the already-unit-pinned ACTIVITY_DRILLDOWN_URL). The manager Playthrough specs are live e2e journeys (gate-verified green + 5/5 deterministic at iter-05), correctly out of harden's unit-deepening scope. Dims 2/3/5 pre-covered (manifest error/TODO/fuzz); Dim 6 N/A (locators, no perf-sensitive path).
**Knowledge backfill:** none — all M204 invariants are now either documented inline AND enforced by a drift guard (reporter-override, Pass 1) or pinned as a deliverable-existence test (manager coverage, Pass 2).
**Stop condition:** stabilized — coverage delta 0.0% (< 2%) across all 3 passes AND the Pass 3 dimension scan found nothing new. Final harden complete; the milestone is ready for /developer-kit:close-milestone.
