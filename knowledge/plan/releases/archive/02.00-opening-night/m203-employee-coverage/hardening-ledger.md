# Hardening Ledger — M203 Employee-vantage coverage

Final-mode cumulative-scope hardening of the M203 footprint (the employee-coverage additions to the
`playthroughs/` section, built on the M202 foundation). Ran after the exit gate fired at iter-06 (6/6 employee
use cases pass on cold reset-to-seed, 0 false-fails over 5 runs), before `/developer-kit:close-milestone`.

Scope (cumulative milestone-touched code, `release/02.00-opening-night...HEAD`): the rext authoring copy's
`playthroughs/` section — the 3 manifest YAMLs (`profile.yaml`, `skill-paths.yaml`, `ai-simulations.yaml`), the
`seed/seed-worlds.yaml` employee-data additions, the per-surface page objects
(`e2e/lib/{profile,skill-path,simulation}-page.ts`), and the `e2e/run-playthroughs.sh` iter-05 Sentinel-reload
fix. The Playthroughs themselves (`e2e/tests/*.spec.ts`) already pass green + 5/5 deterministic; harden deepened
the SUPPORTING code, not the browser drive.

## Pass 1 — 2026-07-02 — final

**Iters hardened this pass:** all milestone-touched code (final cumulative scope) — TS page-object subsystem.
**Tiks covered since prior pass:** all iters in milestone (first harden pass).
**Coverage delta on touched files:**
- e2e/lib/url-shapes.ts: NEW (extracted) -> ~100% stmts/branches (22 pure-logic cases)
- e2e/lib/{skill-path,simulation,profile}-page.ts: route/landmark decision logic now unit-covered via delegation (was browser-only)
**Tests added:**
- final -> e2e/tests/url-shapes.unit.spec.ts: 22 unit/edge cases (chapter-player-vs-detail, sim-launch-vs-detail, timeline dated-range, sample-slug catalog pins, `/g`-flag lastIndex hazard, single-source pattern pins)
**Bugs surfaced + fixed inline:**
- The chapter-player + sim-launch route predicates used a bare `\b` terminal, which false-matched look-alike sibling segments (`/skill-path/x/chapter-list`, `/sim/x/start-now`) — a green-but-wrong hazard. Anchored the terminal segment with `(?:[/?#]|$)` (commit ff38105)
**Flakes stabilized:** none (no pre-existing flakes; new tests deterministic).
**Knowledge backfill:** none this pass (deferred to Pass 3 with the discipline note).
**Stop condition:** continue-to-next-pass — Go seed/manifest/validator + the bash Sentinel-reload gate not yet swept.

## Pass 2 — 2026-07-02 — final

**Iters hardened this pass:** all milestone-touched code — Go manifest/seed subsystem + the bash runner gate (cross-iter integration).
**Tiks covered since prior pass:** all iters in milestone (final cumulative).
**Coverage delta on touched files:**
- manifest/seed_worlds.go: 100% -> 100% (already saturated)
- manifest/validator.go: precondition-coverage positive+negative arms all covered (unknown world/hero/tier/capability, empty-world, blank-precondition-skip, free-form-actor) — confirmed, no gap
- the 3 M203 manifest YAMLs + seed/seed-worlds.yaml: regression-guarded in lockstep by manifest/corpus_test.go (pre-existing)
**Tests added:**
- final -> manifest/runner_safety_test.go: 1 cross-iter drift guard (`TestRunnerSafety_SentinelReloadGate`) pinning the iter-05 post-seed Sentinel Reload (runs the reload, drives Sentinel's own AuthorizationService/Reload RPC, per-stack offset base 8087+OFFSET, non-fatal)
**Bugs surfaced + fixed inline:** none.
**Flakes stabilized:** none.
**Knowledge backfill:** none this pass.
**Stop condition:** continue-to-next-pass — the Sentinel-reload drift guard is a new dimension found this pass; re-measure needed to confirm nothing else remains.

## Pass 3 — 2026-07-02 — final

**Iters hardened this pass:** all milestone-touched code (final cumulative re-sweep + stabilization measurement).
**Tiks covered since prior pass:** all iters in milestone (final cumulative).
**Coverage delta on touched files:**
- Go footprint: 94.8% / 97.6% / 99.4% / 100% — UNCHANGED vs Pass 2 (delta 0%)
- TS footprint: url-shapes.ts ~100%, no new pure-logic surface found (every remaining page-object method is a branch-free live-Page locator construction — no unit seam)
**Tests added:** none (dimension scan found nothing new).
**Bugs surfaced + fixed inline:** none.
**Flakes stabilized:** none — flake gate PASSED (3/3 consecutive clean runs of the new TS + Go tests).
**Knowledge backfill:** `corpus/ops/demo/playthroughs.md` §"The per-surface page-object / locator layer" — documented the M203 employee-journey surfaces + the pure `url-shapes.ts` predicate module + the route-shape discipline truth (anchor the terminal segment, never a bare `\b` — the look-alike-segment hazard surfaced in Pass 1).
**Stop condition:** stabilized — coverage delta < 2% across Pass 2 -> Pass 3 (0%) AND the Phase 2 dimension scan found nothing new.
