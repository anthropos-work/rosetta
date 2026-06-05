# Release Review: v1.1 "show floor"

**Date:** 2026-06-05
**Milestones:** M3 · M4 · M5 · M6 · M7a · M7b · M7c · M8 (8 milestones)
**Review:** 4 parallel sweeps (scope+deferrals · supply-chain+code · docs+decisions · tests+metrics) + live
boundary gates, all driven from `release/01.10-show-floor`.

## Verdicts
- **Scope:** GREEN — every milestone delivered its `In:` scope; the 2 M3-deferred injection recipes landed in
  M8; M7c's waived surfaces (taxonomy + content) homed to v1.2; no Fate-3-undelivered / unaccounted / repeat-defer.
- **Supply chain + code:** GREEN — all deps standard + permissively licensed (BSD/MIT/Apache); `go vet` clean.
- **Docs:** YELLOW → fixed (see below).
- **Tests + metrics:** GREEN — test count 175 (v1.0) → **409** (v1.1) (+134%, no regression); all gates green.

## Findings + fixes (all resolved in Phase 7)

### Scope
- [x] [info] roadmap-vision.md didn't explicitly seed the v1.2 "snapshot mechanism" → **added** an explicit
  bullet (lift M7c's `waived` taxonomy + content via skiller-snapshot + Directus snapshot-replay).

### Code Quality
- [x] [should-fix] gofmt misalignment in `stack-seeding/cmd/stackseed/main.go:93` (the M7c seeder-registration
  comments) → **gofmt -w**, clean.

### Documentation
- [x] [must-fix] M8 milestone records carried the stale **"M5"** label (renumber leftover) in
  `overview.md` / `decisions.md` / `spec-notes.md` → **relabeled to M8**.
- [x] [should-fix] M8 `decisions.md` was a stale placeholder (old M5 pre-build open questions) → **rewritten**
  with the actual M8 decisions (M8-D1 open-questions-resolved, M8-D2 express-gate-CI, M8-D3 cert-redirect-recipe).
- [x] [should-fix] the demo recipes reference the `datadna` CLI but `seeding-spec.md` didn't document it →
  **added** a "Verifying a seed — `datadna`" subsection (catalog/introspect/measure/diff) to `seeding-spec.md`.

### Tests & Benchmarks
- [x] No gaps. The full surface is green at the boundary (see below).

### Decision Consolidation
- [x] No unblended/conflicting decisions across the 8 milestones. The strategy decisions (M7a-D3 direct-Postgres,
  M7b-D1 extend-M0-to-data, M7c re-scope waiver) are coherent + documented in `alignment_testing.md` + the records.

## Live boundary verification (re-confirmed green)
- **3 Go suites:** rosetta `test/alignment` (7 pkgs) · stack-seeding (8 pkgs) · clerkenstein (13 pkgs) — all `ok`.
- **4 Clerkenstein alignment gates:** Go **22/22** · JS/FAPI **9/9** · `@clerk/express` **9/9** · deployment/injection
  **7/7** — all **100%/100%, no divergences**.
- **Seeding gates:** isolation guard (3-layer, audit clean), data-DNA `measure` **100%/critical 100%** over the 8
  reachable surfaces, `diff` drift-detection proven, login→**200** on demo-1.
- **Triple-clean (Phase 8b):** stack-seeding suite 3/3 clean (offline-deterministic; project flake history 0).

**Conclusion:** clean release. No blockers, no escape-hatch deferrals, no Fate-3-undelivered. Ready to merge +
tag `v1.1`.
