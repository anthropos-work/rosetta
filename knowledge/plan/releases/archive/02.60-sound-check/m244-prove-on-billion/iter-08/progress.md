**Type:** tik — gate (b) #3 (intv-ack), under TOK-01. Run 4, tik 3.

# iter-08 — progress

## Root cause (proven, from the saved sweep artifacts + billion DB + platform source)
The iter-07 sweep saved: `intv-voice-pass · player` → **mainLen 0 (EMPTY)**; `intv-voice-fail · player` →
**mainLen 205** ("Congratulations! Interview completed / Thank you for your time. Your responses have been
recorded…" — the ACK). So it is NOT a language/regex issue: the PASSED interview player page renders EMPTY,
the FAILED one renders the ack.

The two seeded clones differ ONLY in `evaluation_status` (validation_attempt_results): pass=`passed` (score 81)
vs fail=`failed` (score 0); identical otherwise (val_attempts=1, extraction=1). In
`AISimulationResult.tsx`, the interview render gate (line ~549) is:
`isInterview && plan && data && (isManagerReportEnabled || isPlayerReportEnabled || <M232 demo-widen>)` →
`<InterviewReport>` : `<AISimulationUserHiddenResult>` (the ACK). The M232 demopatch
`next-web-interview-flag-result` widened that gate for a demo (no PostHog) for **BOTH scopes**. So a PASSED
interview (extraction report available) took the player-scope `InterviewReport` branch and rendered EMPTY,
while a FAILED interview fell through to the ACK. **But an INTERVIEW has no player-facing scored report** — the
candidate sees only an acknowledgement, which is the product's real behaviour AND exactly what the M236 sweep's
`player-interview` shape asserts ("the candidate is NOT shown a scored report for an interview").

## Fix (rext demopatch, committed + tested + pushed — f46e082, tag → 8756ec0)
Scope the demo-widen in `next-web-interview-flag-result` to **isManagerScope**:
`(… || (isManagerScope && !(POSTHOG absent)) …)`. The interview PLAYER now always falls to the ACK (pass AND
fail); the MANAGER report (gate g, iter-06) is untouched (manager scope still opens the gate). The container
FETCH gate intentionally stays open (harmless — data fetched, render gate closed for the player).
- `post_sha256` recomputed = `99903aec…`, **sha-method-verified** (my apply-replay reproduces the CURRENT
  post_sha `f4eaea3a…` exactly, so the new value is trustworthy).
- `isManagerScope` (= isManagerView) confirmed in scope at the anchor; pristine file matches `pre_sha256`.
- + a regression test (`test_result_demo_widen_is_manager_scoped`) pinning the manager-scoping.
- demopatch tests (12) + tooling + frontend-build (254) green.

## Effect (logical + coupled to gate h for the live proof)
With the render gate manager-scoped, the interview PLAYER renders `UserHiddenResult` (the ack) for BOTH pass
and fail → the sweep's `player-interview` ack passes. Combined with iter-07 (2 voice cells → presence-only,
denom 47), **all 3 gate-(b) residuals are resolved**. The live 47/47 lands at the gate-(h) cold reset-to-seed
(re-pin billion to `8756ec0` → next-web re-bake bakes this demopatch + re-seed). 0 platform edits.

## Close — 2026-07-22

**Outcome:** gate (b) #3 resolved — root-caused the passed-interview player empty render to the M232 demopatch
opening the render gate for the player scope; fixed by scoping the demo-widen to isManagerScope (player →
ack, product-real + sweep-expected; manager report untouched). sha-verified + regression-tested + pushed. All
3 gate-(b) residuals now resolved (2 voice presence-only + this). Metric stays **3/8** (gate (b) counted live
at gate h). closed-fixed.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (milestone gate a–h not fully met; 3/8; gate (b) all 3 residuals now resolved, awaiting live 47/47 at gate h)
**Phase 5 grading:** (1) gate-met: n (3/8) — (2) triggered-tok: n (iter-06/07/08 all made measurable progress) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik 3 of run 4) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (root cause = demopatch scope, not seed/language), D2 (fix at the demopatch render gate scoped to manager, container fetch left open), D3 (sha-method verified against the current post_sha before trusting the new one) — iter-08/decisions.md
**Side-deliverables:** none (all in the #3 planned scope; 0 platform edits).
**Routes carried forward:**
- **live 47/47** (gate b) + this demopatch's live effect (interview player renders the ack) → gate-(h) cold
  reset-to-seed (next-web re-bake at tag `8756ec0` + re-seed). Named handler: gate (h).
**Lessons:** (1) a "renders empty" PLAYER surface for one pass/fail variant but not the other points at a
render-gate branch keyed on `evaluation_status`, not the seed data (both clones had identical data). (2) a
demo-widen demopatch that ignores SCOPE can over-open a gate — scope the widen to the vantage that actually
has a surface (here: manager); the sweep's own shape comment is the spec for what each vantage should render.
(3) always verify an sha-computation METHOD against a known-good value (the current post_sha) before trusting
a recomputed one.
