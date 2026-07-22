# M244 — Progress

Iterative milestone (the closer). Primary metric: the multi-part exit gate (a–h) met cold on `billion`.

## Running ledger

- iter-01 (tok/bootstrap): TOK-01 authored (staged cold billion bring-up → gate-parts a–h one-per-tik); Phase 0b KB-fidelity YELLOW (denominator=49, spec=40); pre-flight rung zero mapped (billion bare, pin=m243) — see iter-01/progress.md
- iter-02 (tik): Foundation GREEN — re-pinned billion to m243 (rung-zero), gate (a) ORG-CLEAN PASS (23 fixtures/0 tokens live), gate (e) DEF-M226-01 serve-reap CONFIRMED (7→0 ports), cold reset-to-seed BRINGUP_EXIT=0 + fresh green autoverify (12:51Z) + all peer origins serving + 42,790 skills/1,644 sessions/49-pair manifest. Metric 0/8→2/8. — see iter-02/progress.md

- iter-03 (tik/tooling-iter): gate (b) — fixed the content-stories runner's bash-3.2 parse wall (unrunnable on macOS peers) + tailnet cold-render flakiness (settle-budget + per-pair warm retry); 3 rext fixes shipped+tagged+pushed (`sound-check-m244-content-sweep-robustness`→01206e7). Gate (b) 37 cold → **46/49 reliable**; 3 DETERMINISTIC residuals (voice/interview render, full DB content) route to gate (g) + a voice-presence-only decision. closed-fixed-partial. Metric 2/8 (b not green). — see iter-03/progress.md

## Next-iter queue
- iter-04 (tik): **gate (g)** — the interview plan-section-id alignment assertion; likely lands the intv-voice-pass gate-(b) residual too. Then investigate the 2 voice residuals (hire-voice-fail, asmt-voice-pass-en) → re-prove gate (b).
- gate (h) latency p95<5s + gate (c) the 40 specs + gate (d) academy twin + gate (f) 3 drift-carries in following tiks. NB the content-stories tailnet cold-render flakiness (fixed via settle+retry) likely affects gate (c)'s browser sweeps too — reuse the pattern.

## Carry / disposition tracking (gate parts a–h + inherited)
- (a) ORG-CLEAN ✅ (23 fixtures, 0 tokens, live billion) · (e) DEF-M226-01 ✅ (serve-reap 7→0, tested+passing) — **2/8 discharged**.
- (b) content-stories: **46/49 reliable** (flakiness FIXED; 3 deterministic voice/interview-render residuals → gate g + voice-presence-only decision) — near-green, NOT counted.
- (c) 40 specs · (d) academy twin · (f) 3 drift-carries · (g) interview alignment assertion · (h) v2.6 fixes + p95<5s — OPEN (billion green + serving; unblocked).
- Inherited: DEF-M239-01 (ENOSPC loud-build-fail) · reap-17700 standing-9 (test isolation) · DEF-M240-01 (video exhibit, conditional on Bunny keys) — OPEN.
- Enabling foundation: billion demo-1 GREEN at m243, cold reset-to-seed, fresh green autoverify 2026-07-22T12:51:20Z.
