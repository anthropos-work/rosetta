# Release Review: v1.5 "prop room"

**Date:** 2026-06-14
**Milestones:** M21 (structure-capture) · M22 (provision-lifecycle) · M23 (content-cutover) · M24 (docs-hygiene) · M25 (field-bake)
**Method:** 5-dimension parallel review workflow (supply-chain · scope/fate/decisions · code-quality · docs/KB · tests/metrics). Code lives in `rosetta-extensions` @ tags `prop-room-m21..m25`; rosetta holds docs/plan.

## Gate summary
- **Supply-chain (P0):** GREEN — 0 called CVEs (govulncheck, all 4 Go modules, against the pinned go1.25.11 stdlib); licenses all permissive (1 transitive-only MPL-2.0, not compiled in). Lockfile written.
- **Scope/Fate/Decisions (P1/1b/5):** YELLOW — all deliverables Fate-1; both M21 Fate-3 → M23 cross-validated landed; DEF-M21-02 resolved in M25; no repeat-defer; decisions all blended. Gap: 2 standing follow-ups lack a roadmap-vision landing spot.
- **Code quality (P2):** YELLOW — vet/build clean, firewall coherent + never weakened. 1 should-fix (test/prod drift), 3 cosmetic gofmt.
- **Docs/KB (P3/3b):** RED — 1 stale false-capability section the M24 sweep missed.
- **Tests/Metrics (P4/4b):** GREEN — all suites pass, Go +131 / Python +99, flake 0, no coverage drop.

## Scope
- [x] Full scope accounted — every roadmap deliverable across M21–M25 landed Fate-1; no Fate-3-undelivered, no unaccounted items. (info)
- [ ] [should-fix] DEF-M21-01 (replayCmd conn-seam) — tracked follow-up with no durable landing spot → add to `roadmap-vision.md` Unscheduled backlog.
- [ ] [should-fix] M25-D9 (dev-2 taxonomy `rc=4` migrate-ordering) — same; add to `roadmap-vision.md` Unscheduled backlog.

## Code Quality
- [ ] [should-fix] `stack-injection/tests/test_injection.py:47` — add `GOTOOLCHAIN=local` to the mintpk test env to mirror M25's `up-injected.sh` fix (else the drift guard fails on any host with go < 1.25.11). (ext)
- [ ] [nice-to-have] `gofmt -w` 3 cosmetic files in stack-snapshot (structure.go, adapters_test.go, directus_test.go). (ext)
- [x] Firewall safety-critical code coherent + never weakened; integration seams clean; 0 new third-party deps. (info)

## Documentation
- [ ] [must-fix] `corpus/ops/snapshot-cold-start.md:144-155` — rewrite the stale "M10 collection-schema gap / nothing executed / exit 4" framing (missed by the M24 sweep; file absent from the release diff) to the v1.5 two-path posture (executed on `--local-content` per M21/M22/M23 → exit 0; prod-read fallback exit 4). It contradicts the release's headline capability.
- [ ] [nice-to-have] `corpus/ops/snapshot-spec.md:362` heading still reads "recipe corrected in fix16" (provenance only; body correct) — normalize to v1.5 vocabulary.

## Tests & Benchmarks
- [x] All suites pass; Go 736→867 (+131), Python 360→459 (+99); flake 0; no coverage >2pp drop. (info)
- [ ] [nice-to-have] Release metrics rollup says Go=850 (M24-frozen); true final HEAD = **867** (stack-snapshot 316→333 from M25's fix+harden tests). Roll the final number into the release metrics.json.

## Decision Consolidation
- [x] All "blend into knowledge" decisions reached their target docs (M23-D1..D5, M25-D2, M24-D2 backref tags grep-confirmed); no cross-milestone conflicts. (info)
