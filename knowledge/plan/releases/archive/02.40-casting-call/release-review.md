# Release Review: v2.4 "casting call"

**Date:** 2026-07-18
**Milestones:** M222, M223, M224, M225, M226, M227, M228 (7 total)

## Supply chain (Phase 0)
- ✅ npm audit (e2e): **0 vulnerabilities** (prod + all). Dev-only toolchain (@playwright/test, @types/node, typescript).
- ✅ Go deps: `github.com/anthropos-work/ai v1.40.1` — the only third-party dep, **unchanged in v2.4** (added M45).
- ✅ No new deps introduced by v2.4. No license concerns (dev-only TS + one internal Go dep).
- Lockfile: `dependencies.lock`.

## Scope (Phase 1)
- ✅ All 7 milestones delivered: M222 (spike/GO), M223 (seeder), M224 (render, closed-on-gate), M225 (demo-integration),
  M226 (billion-proof, closed-on-gate), M227 (believability fixes), M228 (live re-prove, closed-on-gate).
- ✅ No `closed-incomplete` milestones → no carry-forward.md, no undelivered Fate-3 items.
- ✅ The release thesis — the recruiter-vantage hiring org, 45 candidates on 5 shared positions compared side-by-side —
  is delivered AND proven live on billion.

## Code quality (Phase 2)
- ✅ rosetta release diff = **111 files, 100% docs** (97 knowledge + 14 corpus). 0 code, **0 platform-repo edits**.
- ✅ rext tooling reviewed per-milestone; consistent guard/seeder patterns; the M228 render-probe + seeder-guard
  reviewed at close-milestone (adversarial scenario recorded).

## Documentation (Phase 3)
- ✅ `corpus/services/hiring.md` carries the full render+seed model incl. the M228 live findings (intercepting-route
  drawer + the incomplete-guard correction).
- ✅ `state.md` + `roadmap.md` reflect v2.4 code-complete; all 7 milestone overviews archived.
- [ ] (nice-to-have, non-blocking) The intercepting-route drawer finding is in `hiring.md`; the render-probe
  `RENDER_ONLY_SIM` knob is a rext-internal test detail (documented in-code) — no corpus gap.

## Tests & benchmarks (Phase 4)
- ✅ rext stack-seeding: green, **96.8%** stmt coverage. playthroughs: green. e2e: tsc clean.
- ⚠️ rext demo-stack: **644 passed / 14 failed** — ALL pre-existing or environment-gated carries, NONE v2.4 regressions
  (see Completeness Ledger below). The rext repo is at a fixed commit; the rosetta docs merge+tag does not touch it.

## Decision consolidation (Phase 5)
- ✅ Per-milestone decisions blended into corpus at each close (M222 hiring model, M227 fixes #1/#2/#4, M228 render).
- ✅ No cross-milestone decision conflicts. The mirror-table read-model (M222) held consistently through M228.

## Metrics (Phase 4b)
- ✅ 0 platform edits · supply-chain GREEN · flake 0 · rext seeders 96.8%. Aggregate: `metrics.json`.
- No regression vs v2.3 (v2.4 adds a hiring vantage; no capability removed).
