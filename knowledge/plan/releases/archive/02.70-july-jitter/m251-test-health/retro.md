# M251 — Retro (test-health)

## Summary
The test-health fan-out lane of v2.7 "july jitter" — realizes the reserved v2.6→v2.7 test-health carry.
Discharged the standing demo-stack test debt without a live box: rostered the 2 orphan unit specs (clearing
the RED `UnitSpecsAreExecuted` guard / runner exit 2) and re-pointed the 7 mechanical assertions (6
`test_cockpit` overlay/academy + `test_public_host` port-13001) at the **deliberately-changed** behaviour
(overlay 30s-window removal @ M218; per-hero academy link removal 2026-07-15; hiring 3001→13001 UI-tier
fronting @ M226) as **removal-guards**. Section close, all sections Fate 1, 0 platform-repo edits. Code-of-record:
rext tag `july-jitter-m251-test-health` @ `e9e29a1` (on origin).

## Incidents This Cycle
None. 0 P2 flakes (flake gate 5/5 clean), 0 regressions (M251's touched files are 207/207 green; the 8
full-suite failures live in untouched `test_purge`/`test_ant_academy*` files → M254). No bug surfaced during
build, harden, or close.

## What Went Well
- **Mutation-verify over line-coverage.** For re-pointed removal-guards the right robustness signal is "does
  the guard go RED when the removed behaviour returns" — 4/4 fire. A guard proven live is not vacuous; this
  pre-empted the exact false-confidence failure mode a removal-guard invites (Phase 2c adversarial scenario).
- **The composition was confirmed against disk before build** (Phase 0b GREEN, spec-notes topic→doc→code
  triples), so there was no exploratory drift — a clean section shape with no surprises.
- **Scope discipline at close.** The full-suite run surfaced 8 live-gated failures whose enumeration differs
  slightly run-to-run (host-sensitive: academy reap ×6 + no host-isolation this run, vs the ×5 + host-isolation
  in decisions.md). Rather than chase them into M251, they were correctly confirmed Fate-2 → M254 (they need a
  live box) — the count holds at 8, which is the load-bearing fact.

## What Didn't (go as smoothly)
- **A pre-existing, self-documented handbook drift** surfaced during Phase 4 reconciliation: the demo-stack
  README quotes "730 tests" vs the actual 869 collected (the README itself flags "this line drifted by 306";
  the number is machine-UNGUARDED by design). NOT M251-caused (its re-points are count-neutral renames) and
  NOT in a file M251 touched → routed to rext-hygiene, not fixed here (fixing it would re-tag the frozen
  code-of-record for an unrelated drift). The guarded GUIDE count ("47", `TestGuideDocTruth`) is accurate.
- **`pytest-randomly` is not installed**, so the flake gate ran plain repeated runs rather than randomized
  order. Acceptable for these order-independent unittest-style render/logic tests (no shared mutable
  temp-dirs/ports/caches), but a randomized-order flake gate would be a stronger signal for the demo-stack
  suite generally — a rext-tooling nit, not an M251 concern.

## Carried Forward
- **8 live/env/docker-gated demo-stack failures** (`test_purge` + `test_ant_academy*` launcher/reap +
  `test_ant_academy_clerk_wiring`) → **M254** gate parts (g)+(h) — need a live billion box. ⚠️ **M254's
  overview names "~2"; the real set is 8** — flagged to the orchestrator to correct M254's overview when M254
  runs (not edited from this close window; M254 not active).
- **Optional `corpus/ops/verification.md` demo-stack-suite index anchor** → **M247** (owns the ops-doc
  reground; Fate 2).
- **Inherited M246 rows** naming M251 as a possible dest: D-03/D-04 (`test_injection.py` /
  `exposure_claim_guard.py` skillpath fixtures) → M247 drift-ledger triage; D-08 (fake-FAPI probe) + D-09
  (academy peripheral) → M254. All Fate-2, all confirmed, none aged out.

## Metrics Delta
(from `metrics.json`)
- run-unit: 7 → **9** executed specs (172 tests, exit 0); `UnitSpecsAreExecuted` RED → **GREEN**.
- demo-stack mechanical failures: **7 → 0** (the 6 `test_cockpit` + `test_public_host` re-points now pass).
- demo-stack full suite: **861 pass / 8 fail** (was 839/8 baseline; the 8 fail are the live-gated carry → M254).
- Mutation-verify: **4/4** removal-guards fire. Flake: **0** (5/5 clean). Lint: clean. Platform-repo edits: **0**.
