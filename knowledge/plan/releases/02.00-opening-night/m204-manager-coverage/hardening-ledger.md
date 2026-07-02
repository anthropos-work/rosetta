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
