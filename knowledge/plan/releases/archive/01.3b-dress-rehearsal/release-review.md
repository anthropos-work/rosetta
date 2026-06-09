# Release Review: v1.3b "dress rehearsal"

**Date:** 2026-06-09
**Milestones:** M16 (land-fixes) · M17 (rerun-safety) · M18 (verify-net) · M19 (frontend-tier) · M20 (lifecycle-convergence)
**Branch:** `release/01.3b-dress-rehearsal` → (pending) `main`, tag `v1.3.1`
**Method:** 5 parallel review sweeps (supply-chain · scope+deferrals · code-quality · docs+KB · tests), real `go test -race` / `govulncheck` / `pytest` / `shellcheck` runs. Provenance: `wf_1cf06222-5b0`.

## Verdict summary
| Phase | Verdict |
|---|---|
| 0 Supply-chain | **GREEN** (0 called third-party CVEs; all-permissive; stdlib-only Python) |
| 1 Scope | **GREEN** (all 14 field issues delivered Fate-1) |
| 1b Deferrals | **GREEN** (0 new; DEF-M10-01 → v1.4 signed/untouched/not-aged) |
| 2 Code quality | **GREEN** (gofmt/vet/shellcheck/py_compile clean) |
| 3/3b Docs | **YELLOW → fixed** (1 must-fix + 1 should-fix; 1 finding stale/already-resolved) |
| 4 Tests | **GREEN** (1,096 tests pass; 0 flakes) |
| 4b Metrics | **GREEN** (Go +23, Python grew; no regression) |

## Scope
- [x] All 14 `/demo-up` field issues (ISSUE-1..14) delivered Fate-1 across M16-M20. No scope gaps, no Fate-3-undelivered, no unaccounted items. (info — verified)

## Code Quality
- [x] [info] `dev-stack/dev-setdress.sh` `mktemp -d` build-staging dir is intentional + env-overridable (`DEV_SETDRESS_BIN`) — no action. (verified, no fix)
- [x] Static checks clean: gofmt + `go vet` (4 Go modules), shellcheck (touched bash; info-only SC1091 source directives), py_compile. No must-fix.

## Documentation
- [x] [must-fix] `corpus/ops/README.md` index missing rows for `snapshot-cold-start.md` (M20) + `demo/frontend-tier.md` (M19) — **FIXED** (rows added after the verification row).
- [x] [should-fix] `corpus/ops/demo/README.md` mechanism-guides index missing `verification.md` — **FIXED** (entry added after idempotency).
- [x] [should-fix → stale] `CLAUDE.md` `/demo-up` row auto-set-dress mention — **ALREADY RESOLVED** by M20 close (`d29f513`); the docs agent reviewed a stale view. No action (verified current row reads "full UI tier + auto-set-dressed — the M20 demo set-dress flow, mirroring `/dev-up`").
- [x] [info] All 38 internal links across the 4 new docs + safety.md resolve; no KB debt (largest new doc 175 lines); CLAUDE.md/roadmap.md/state.md mutually coherent.

## Tests & Benchmarks
- [x] Full suites pass: Go 736 test funcs (4 modules, `-race -count=1`), Python 360 collected (5 sections). 1,096 tests, 0 failures, 0 flakes, tree clean. No coverage gaps surfaced.

## Decision Consolidation
- [x] All decision-triage items across M16-M20 processed (blended into the relevant new doc or archived per the v1.3 corpus precedent). No unblended, no cross-milestone contradictions.

## Supply chain
- [x] [should-fix → accept] Go toolchain `go1.25.3` has 12 *called stdlib* advisories per module (net/textproto, crypto/x509, crypto/tls, net/url, os) — cleared by go1.25.11+. **0 called third-party CVEs.** Identical to the v1.3.0 baseline; accepted as an environment/toolchain item (not code), same disposition as v1.3. Lockfile: `dependencies.lock`.

## Recurring-pattern note (for the retro)
The docs YELLOW (new docs not added to the parent index) is the **"new deliverable shipped without its index row"** class that recurred in v1.3 (M14, M15). It recurred again here (M19/M20 docs vs the main ops README). → retro carry: a release-close index-row check (now done) should become a per-milestone close habit on corpus directory READMEs, not just per-unit READMEs.
