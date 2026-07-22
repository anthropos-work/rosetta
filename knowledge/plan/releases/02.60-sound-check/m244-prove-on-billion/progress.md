# M244 — Progress

Iterative milestone (the closer). Primary metric: the multi-part exit gate (a–h) met cold on `billion`.

## Running ledger

- iter-01 (tok/bootstrap): TOK-01 authored (staged cold billion bring-up → gate-parts a–h one-per-tik); Phase 0b KB-fidelity YELLOW (denominator=49, spec=40); pre-flight rung zero mapped (billion bare, pin=m243) — see iter-01/progress.md
- iter-02 (tik): Foundation GREEN — re-pinned billion to m243 (rung-zero), gate (a) ORG-CLEAN PASS (23 fixtures/0 tokens live), gate (e) DEF-M226-01 serve-reap CONFIRMED (7→0 ports), cold reset-to-seed BRINGUP_EXIT=0 + fresh green autoverify (12:51Z) + all peer origins serving + 42,790 skills/1,644 sessions/49-pair manifest. Metric 0/8→2/8. — see iter-02/progress.md

- iter-03 (tik/tooling-iter): gate (b) — fixed the content-stories runner's bash-3.2 parse wall (unrunnable on macOS peers) + tailnet cold-render flakiness (settle-budget + per-pair warm retry); 3 rext fixes shipped+tagged+pushed (`sound-check-m244-content-sweep-robustness`→01206e7). Gate (b) 37 cold → **46/49 reliable**; 3 DETERMINISTIC residuals (voice/interview render, full DB content) route to gate (g) + a voice-presence-only decision. closed-fixed-partial. Metric 2/8 (b not green). — see iter-03/progress.md
- iter-04 (tik): gate (g) — LIVE-PROVED the interview report renders **EMPTY** on billion (seeded depth/adoption/sentiment absent behind "View Report"); the plan-section-id alignment assertion would be **RED**. The S-8/S-9 gap is a real defect (not just unproven); intv-voice-pass residual is the same failure. Root-cause surface identified (seed-plan vs flag-demopatch); fix routed forward (0 platform edits). closed-no-lift. Ended on a user-requested pause. — see iter-04/progress.md
- iter-05 (tik): gate (g) FIX — ROOT CAUSE = `directus.simulations_extraction` (the interview PLAN) **never captured by the snapshot** (0 rows on demo) → `interviewExtractionPlan` null → report empty. FIX: added the table to the directus capture surface (`stack-snapshot`, parent-scoped public-only), tested + pushed (rext `e74e563`, tag moved). Plan loaded to the demo → report render 0→520 (plan is load-bearing). NOT green: full section render (behind "Explore Key Moments") + the alignment assertion + the live capture-path exercise routed. **intv-voice-pass/player is a SEPARATE ack-page issue** (gate b stayed 46/49). closed-fixed-partial. — see iter-05/progress.md

## Next-iter queue (RESUME HERE — run 3)
- **iter-05 (tik): fix the interview report empty-render** (gate g) — determine seed-plan-missing vs flag-demopatch-ineffective (`AISimulationResultContainer.tsx:519-523` plan `raw`; `:499-505` flag gate), fix under 0 platform edits, add the alignment assertion + prove green. Lands intv-voice-pass → gate (b) 47/49.
- then: the **2 voice cells** (hire-voice-fail, asmt-voice-pass-en) — **Bunny-key check NOT DONE yet**: check `BUNNY_RECORDING_CDN_TOKEN_KEY` + `BUNNY_RECORDING_PULL_ZONE_HOST` reachability on billion → if reachable wire the video exhibit (gate b → 49/49, discharges DEF-M240-01); else presence-only (drop the 2 cells from the denominator → gate b 47/47, update `content-denominator.json`). Then re-prove gate (b).
- then gates (c) 40 specs · (d) academy twin · (f) 3 drift-carries · (h) v2.6 fixes + p95<5s. NB reuse the content-stories settle+retry template for gate (c)'s tailnet flakiness.

## Carry / disposition tracking (gate parts a–h + inherited)
- (a) ORG-CLEAN ✅ (23 fixtures, 0 tokens, live billion) · (e) DEF-M226-01 ✅ (serve-reap 7→0, tested+passing) — **2/8 discharged**.
- (b) content-stories: **46/49 reliable** (flakiness FIXED; 3 deterministic voice/interview-render residuals → gate g + voice-presence-only decision) — near-green, NOT counted.
- (c) 40 specs · (d) academy twin · (f) 3 drift-carries · (g) interview alignment assertion · (h) v2.6 fixes + p95<5s — OPEN (billion green + serving; unblocked).
- Inherited: DEF-M239-01 (ENOSPC loud-build-fail) · reap-17700 standing-9 (test isolation) · DEF-M240-01 (video exhibit, conditional on Bunny keys) — OPEN.
- Enabling foundation: billion demo-1 GREEN at m243, cold reset-to-seed, fresh green autoverify 2026-07-22T12:51:20Z.
