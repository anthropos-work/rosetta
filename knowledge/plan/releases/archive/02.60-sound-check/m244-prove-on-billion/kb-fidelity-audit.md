---
title: "KB Fidelity Audit — M244 prove-on-billion"
date: 2026-07-22
scope: milestone:M244
invoked-by: build-mstone-iters (Phase 0b, iter-01 bootstrap tok)
---

## Verdict
**YELLOW** — no blind areas, no stale claim the live proof reads as truth and is misled by; two stale-narrative
counts (cheap), recorded + partially applied.

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| ORG-CLEAN / content-story scrub-cleanliness (gate a) | `session-clone-spec.md`, `safety.md §3.8`, `content-stories-routes.md` | `stack-seeding/contentsession/cleanliness_test.go`, `stack-seeding/scrub/` | PAIRED |
| content-stories sweep `run-content-stories.sh` (gate b) | `coverage-protocol.md` §content-stories sweep, `content-stories-spec.md` | `stack-verify/e2e/run-content-stories.sh`, `content-denominator.json`, `aggregate-content.py` | PAIRED |
| 40 live-browser specs (gate c) | `coverage-protocol.md`, `playthroughs.md` | `stack-verify/e2e/tests/*.spec.ts` (19 live + unit), `playthroughs/e2e/tests/*.spec.ts` (16 live + 2 unit) | PAIRED |
| anonymous academy /library+/free (gate d) | `content-stories-routes.md` (Academy IN), `services/ant-academy.md` | ant-academy (native), demo set-dress | PAIRED |
| serve-reap DEF-M226-01 (gate e) | `tailscale-serve.md` Step 7 (F12) | `demo-stack/rosetta-demo` down / serve-reset | PAIRED |
| 3 v2.3 drift-carries (gate f) | `tailscale-serve.md`, `latency-budget.md` | rext demo-stack + patches | PAIRED |
| interview plan-section-id alignment assertion (gate g) | `content-stories-routes.md` (interview), `session-clone-spec.md` | to be ADDED by M244 (`added + green`) | DOC-ONLY (build item, not blind) |
| v2.6 fixes + p95 latency (gate h) | `latency-budget.md`, M238–M243 docs | `stack-verify/e2e/tests/latency.spec.ts`, demopatches | PAIRED |
| billion remote bring-up (all gates) | `tailscale-serve.md`, `verification.md` (pre-flight rung zero) | `demo-stack/up-injected.sh`, `ensure-clones.sh` | PAIRED |

## Fidelity Findings

### KB-1 — content-stories denominator: doc-narrative says 29, tooling says 49
- **Source:** `roadmap.md` M244 goal ("v2.5's headline 29/29"), `state.md`, `corpus/ops/demo/coverage-protocol.md:916-918` ("`EXPECTED_PAIRS=29`", "29/29").
- **Expected (per prose):** 29/29 content-story pairs.
- **Actual (per tooling):** `stack-verify/e2e/content-denominator.json` → `expected_pairs: 49`. `run-content-stories.sh:80-134` reads this file as `EXPECTED_PAIRS` and cross-checks the served manifest against it before the sweep. M241 (v2.6) grew simulation sessions 13→23 by pinning EN/IT language counterparts (23×2=46 + 2 skill-path-legacy-player + 1 skill-path-new = 49; ai-labs presence-only excluded by design).
- **Verdict:** STALE (narrative only).
- **Fix owner:** doc. The live proof reads the denominator FILE (49), not the prose, so the implementation is NOT misled — this is YELLOW, not RED. Recorded prominently in `spec-notes.md` so every tik targets 49/49. `coverage-protocol.md:916-918` + roadmap/state "29/29" flagged for close/harden reconciliation once the live sweep lands 49/49 (deferred to have the live result in hand — a pre-iter audit does not rewrite corpus historical-example prose speculatively).

### KB-2 — live-browser spec count: 40, not 39
- **Source:** `spec-notes.md` gate (c), `roadmap.md:462`.
- **Expected (per those two):** 39 live-browser specs.
- **Actual:** 40 (24 stack-verify + 16 Playthroughs). Verified: 16 live-browser playthrough spec files (`playthroughs/e2e/tests/*.spec.ts` minus `url-shapes.unit`/`stack-env.unit`), incl. M243's `assignment-assign.spec.ts`. `overview.md` exit_gate already says 40.
- **Verdict:** STALE.
- **Fix owner:** doc. Applied to `spec-notes.md` (39→40). `roadmap.md:462` flagged for close (align its M244 detail copy with `overview.md`).

## Completeness Gaps
None critical. Gate (g)'s "interview plan-section-id alignment assertion" is net-new tooling M244 builds (correctly DOC-ONLY / a build item, not an undocumented behavior or blind area).

## Applied Fixes
- `spec-notes.md`: gate (c) count 39→40 (+ rationale); added the KB-1/KB-2 pre-flight audit record + the pre-flight rung-zero state (billion bare; pin must be `sound-check-m243-assign-write-playthrough`).

## Open Items (require user decision)
None blocking. Two doc-reconciliation items flagged for close-milestone/close-release (not this pre-flight):
1. `coverage-protocol.md:916-918` + `roadmap.md` M244 goal + `state.md`: reconcile "29/29" → "49/49" once the live sweep confirms 49/49 on billion.
2. `roadmap.md:462`: gate (c) count 39→40 (align with `overview.md`).

## Gate Result
**YELLOW: proceed with tracking.** No blind areas; all gate parts have tooling + doc anchors. The two stale counts are narrative-only — the implementation reads files (49 denominator, 16 playthrough spec files), not the stale prose. Findings recorded in `spec-notes.md`; corpus reconciliation deferred to close with the live result in hand.
