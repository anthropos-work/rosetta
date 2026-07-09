# Release Review: v2.1 "quick change"

**Date:** 2026-07-09
**Milestones:** M208 (re-sync & ground-truth) · M209 (rext re-ground) · M210 (corpus re-ground) · M211 (bring-up acceptance, iterative closed-on-gate)
**Method:** 8-agent parallel review workflow (`wf_01ccca96`), two-repo (rosetta docs/plan @ `release/02.10-quick-change` + rext code @ `quick-change-m211`=`2039103`).

**Headline:** clean release. Both blocking gates GREEN + adversarially verified (1b deferral, 4b metrics). No must-fix. All findings are should-fix / nice-to-have.

## Scope (Phase 1/1b) — GREEN
- [x] All 4 milestones shipped every promised capability; **0 Fate-3-undelivered** (migrate-dev.sh M208→M211, cache-migration M209→M211, KB-1/2/3 flips M208/M209→M210 all landed — traced end-to-end).
- [ ] [should-fix] **M314b** claimed "tracked in roadmap-vision.md" (state.md + context.md) but has no literal entry there — add a one-line unscheduled-backlog entry OR repoint the phrasing to where it lives (coverage-protocol.md / ai-readiness.md).

## Supply chain (Phase 0) — YELLOW (no blocker)
- [ ] [should-fix] Bump `toolchain go1.25.11 → go1.25.12` across the 6 rext `go.mod` — clears two inherited Go-stdlib advisories (GO-2026-5856 crypto/tls ECH MEDIUM *called*; GO-2026-4970 os.Root HIGH *not reachable*), both DB-published since the v2.0 close, both drop-in fixed at go1.25.12. Re-scan to green.
- [x] 0 net-new deps; npm 0 vulns; 0 GPL/AGPL; v2.0 x/net HIGH stayed resolved. `dependencies.lock` written.

## Code Quality (Phase 2) — GREEN
- [ ] [should-fix] `stack-verify/e2e` declares a `typecheck: tsc --noEmit` script but has no `typescript`/`@types/node` devDependency → `npm run typecheck` fails; v2.1 lands new TS (academy/cross-port) into exactly this project.
- [ ] [should-fix] The M209 "drop skiller from verify + demo bring-up" sweep missed two orchestration files still referencing skiller as a **live service**: `stack-injection/gen_injected_override.py` INJECTED map + `stack-verify/repos/run.sh` per-repo test cases.
- [ ] [nice-to-have] Cross-port app classifier `%10000===3077` duplicated (live inline in `coverage.spec.ts` + a re-implemented `isAcademyPort` the unit test asserts against → test validates a copy).
- [ ] [nice-to-have] Cosmetic double-word from the mechanical replace: "the replayed PUBLIC public taxonomy" (doc-comments).
- [x] 0 residual `skiller.<table>` in production paths; `go vet` clean; shellcheck clean; digest-narrowing + migrate-dev.sh + MinRows floor all clean.

## Documentation (Phase 3/3b) — GREEN
- [ ] [should-fix] `service_taxonomy.md:378` "9 Go services" contradicts `architecture_overview.md:54` "8 Go microservices".
- [ ] [should-fix] `context.md` still narrates v2.1 "IN DEVELOPMENT (active) / Next: build M208" while state.md says COMPLETE (partly Phase-10 demote work).
- [ ] [nice-to-have] `idempotency.md` migrate row indexes only `migrate-demo.sh`, not the new `migrate-dev.sh`.
- [x] 0 stale `skiller.<table>` tooling refs corpus-wide; 4 subgraphs uniform; backend.md↔skiller.md cross-link resolves; no doc >1500 lines (bar the frozen v1.x archive).

## Tests & Benchmarks (Phase 4/4b) — GREEN
- [x] Go 6/6 modules green (**1764** test funcs, +19 vs v2.0 1745, every module flat-or-up); TS **103** (+3); shellcheck 25/25; **flake 0**; alignment 100%/100% carry; **all 4 regression rules PASS**.
- [ ] [nice-to-have] `stack-verify/e2e` `tsc --noEmit` can't run (no typescript devDep) — same as the Phase-2 item.

## Decision Consolidation (Phase 5) — YELLOW
- [ ] [should-fix] **The one real cross-milestone contradiction:** M209 routed a schema+**count** fix to M210, but M210 flipped the schema token everywhere while leaving **42,763** as the current count in 3 tooling docs — contradicting M208's authoritative fact-sheet (**42,790**). Fix the 3 docs to 42,790.
- [ ] [should-fix] `stack-snapshot/SKILL.md` body prose still calls the replay target a `skiller` schema — contradicts the merge banner the same milestone (M210) added ~45 lines above.
- [ ] [should-fix] Capture the **cache-migration-as-recapture** pattern (the substantive M211 mechanism) as a durable corpus note (snapshot-spec.md).
- [ ] [should-fix] Capture the recurring **stale-clone / build-scratch** class (M208 re-sync + M211 build-scratch) as a durable note.
- [x] Re-ground substantively complete + correct; M210 correctly refused the phantom 43/44 fabrication.

## Resolution (Phase 7 — all findings landed)
**Corpus/planning (rosetta `47d9ff5`):** count 42,763→42,790 (3 tooling docs) · stack-snapshot/SKILL.md replay-target skiller→public · service_taxonomy 9→8 Go services · M314b roadmap-vision entry · idempotency migrate-dev.sh row · 2 KB notes (cache-migration-recapture + build-scratch freshness).
**Rext (`quick-change-m211` → `7906f3f`):** FIX A toolchain go1.25.11→go1.25.12 (govulncheck clean all 6) · FIX B stack-verify/e2e typescript + tsconfig (typecheck exit 0) · FIX C dropped skiller from `gen_injected_override.py` INJECTED + `run.sh` (suites kept green: test_injection 96, test_verify 104) · FIX D 5 doc-comment double-words · FIX E exported `isAcademyPort` (test validates shipped code).
**Verification:** Phase 8 both trees clean; Phase 8b triple-clean 3/3 (-shuffle, go1.25.12); Go 6/6 green, TS typecheck+units green, govulncheck clean.
**Deferred to the v2.1 rext roll (Phase 10):** TEST-1 + DOC-1 (rext README reconciliations).
