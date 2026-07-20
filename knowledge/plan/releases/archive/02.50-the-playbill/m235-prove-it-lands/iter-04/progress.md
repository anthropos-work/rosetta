# iter-04 — progress

**Type:** tik (under TOK-01, Track A step 1) — fixture-matrix closure.

## Execution
- Sourced 4 cells from prod read-only (postgres MCP, non-PII metadata; confirmed prod via a known pin's slug):
  richest-fan-out first, public predicate + completed + pass/fail, excluding the 9 existing pins.
- Pinned the 4 in `content-sessions.yaml`; strengthened `TestEmbedded` (13 + PASSED-set modality contract +
  per-type passed-and-failed).
- Captured all 13 through the fixed path (counts-only). The **fail-closed gate FIRED** on the first pass
  (asmt-voice-pass-2, a 4-char token in interaction.payload) — diagnosed as a `SurvivingToken` vs `Scrub`
  word-boundary MISMATCH (the leak check used plain substring; Scrub uses word boundaries). Fixed
  `SurvivingToken` to mirror Scrub + 2 tests; re-captured — all 13 clean.
- Regenerated both honesty-gated presets (`content-manifest.json` 13 sessions; `seed-generation-manifest.yaml`
  13 pins + fixed its stale synthesize-first header line). Updated the 2 projection-count tests (9→13).
- rext commit `590082a`, tag `playbill-m235-fixture-matrix`.

## Re-measure
- Fixture = 13 sessions. Assessment PASSED set = 2 voice + 1 code + 1 document (the gate's modality contract).
- All 4 sim types present in BOTH passed and not-passed. The 4 new render a fan-out (skills 2–3 / crit 4–7;
  interview report on intv-voice-fail). All leak-clean (fail-closed gate + offline gate).
- The 9 existing re-captured BYTE-IDENTICAL (deterministic scrub — no diff). Full module go vet/test green.

## Close — 2026-07-19

**Outcome:** Fixture-matrix closed — 13 sessions; assessment PASSED = 2 voice / 1 code / 1 document; every sim
type present passed AND not-passed; both honesty-gated manifests regenerated. Milestone live gate not moved (Track-A).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (Track-A readiness — the live (session×action) landing gate is Track-B → coverage/Playthrough tiks + M236)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** iter-04/decisions.md (cell sourcing).
**Side-deliverables:** `SurvivingToken` word-boundary fix (a correctness fix to the iter-03 leak check, surfaced by this iter's capture; the fail-closed gate did its job — halted loudly rather than writing a doubtful fixture).
**Routes carried forward:** Playthrough use-case + coverage descriptors per (session × action); non-sim product sections; M230 clone re-anchor + ANT_ACADEMY coverage descriptor. The LIVE (session×action) landing proof (cold reset-to-seed + Playwright) → Track-B, routed to M236 per TOK-01 (no local stack).
**Lessons:** a fail-closed gate must share EXACT matching semantics with the mechanism it guards, or it false-positives on the very cases the mechanism correctly handles. Better it halted loudly than shipped a doubtful fixture.
