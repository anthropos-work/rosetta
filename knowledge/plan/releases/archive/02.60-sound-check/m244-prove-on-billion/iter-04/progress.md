**Type:** tik — gate (g), under TOK-01. Run 2, tik 1. (Ended early on a user-requested pause; iter cleanly closed.)

# iter-04 — progress

## Investigation (gate g — the interview plan-section-id alignment)
Established what gate (g) actually is and PROVED the gap is a real defect (not just "unproven"):

1. **What the gate is.** `session-clone-spec.md:208-212`: the interview surface's "exact plan-section match" was routed to M235's prove-it-lands gate, which **closed incomplete and never ran**; M236 only proved the report renders **non-empty**, not that each plan section matches the copied report. Gate (g) is to add + green that exact-match assertion.
2. **The render mechanism.** The interview report (`packages/ui/.../AISimulationResultContainer.tsx`) renders via `InterviewReport` from an `interviewExtractionPlan` (`.sections[].sectionId`, parsed from a `raw` JSON — `:519-523`) looked up against the seeded `interviewExtractionData`. It is flag-gated: `isExtractionEnabled = (isManager && flag_interview_manager_report) || (player && flag_interview_player_report)` (`:499-505`) — patched by the `next-web-interview-flag-{container,result}` demopatches (confirmed baked in the built bundle: the `flag_interview_manager_report`/`_player_report` strings are present in `/app/apps/web/.next/server/...`).
3. **The seeded data** (`interview_extraction_results.manager_report.results`) has section keys **`depth` / `adoption` / `sentiment` / …** (each with `key_quote`, a score, an Italian summary). Full content present in the DB.
4. **LIVE PROOF on billion (the decisive finding).** Drove the manager interview view as dan-manager via the cockpit: the attempt-list VIEW renders (breadcrumb "AI Interviews / AI Readiness Interview / Nadia Ferrari" + a "Completed … View Report" row). Clicking **"View Report"** navigates to `/sim/ai-readiness-interview-d62/<reportId>/result/97f3f681…` — and that report page renders **EMPTY**: `depth`/`adoption`/`sentiment`/`Cloud Code` (a seeded key_quote) are **ALL ABSENT** from the DOM. The intv-voice-pass **player** result is the SAME empty page. So the seeded interview extraction does **NOT render** on the report surface.

⇒ **The plan-section-id alignment assertion, once added, is RED live** — the S-8/S-9 gap is a real defect, not merely unproven. The intv-voice-pass gate-(b) residual is this same failure.

## Root-cause surface (for the fix, routed forward)
The report renders empty despite full seeded data + the flag demopatch baked in. Two candidates to decide between next: (a) the `interviewExtractionPlan` `raw` is empty/null for the seeded session (the seed provides the extraction DATA but no PLAN, or the plan source the query reads is absent) → no sections to render; (b) the flag demopatch, though present in the bundle, isn't effectively forcing `isFeatureEnabled` true at render (so `isExtractionEnabled` is false → report hidden). Distinguishing these (and the fix — a seed plan or a demopatch correction, 0 platform edits) is the next step. NB: intv-voice-**fail**/player renders the terse ack (205 chars) while intv-voice-**pass**/player renders empty — a pass-vs-fail routing difference worth checking.

## Close — 2026-07-22

**Outcome:** gate (g) characterized as a real live defect — the interview report renders EMPTY on billion (seeded depth/adoption/sentiment absent); the alignment assertion would be RED. No fix landed (the fix — seed-plan vs flag-demopatch — needs one more determination, routed forward). Ended early on a user-requested pause.
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET (gate g not discharged; 2/8 overall; gate b stays 46/49)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (prior tiks iter-02/iter-03 both made progress) — (3) re-scope: n — (4) user-blocker: n (route-forward-able) — (5) cap-reached: **y** (used as the user-pause marker per the orchestrator) — (6) protocol-stop: n — Outcome: exit-5 (pause)
**Decisions:** D1 (gate-g is a live RED, not just unproven) — iter-04/decisions.md
**Side-deliverables:** none (0 platform edits; investigation only; temp capture spec deleted, trees clean).
**Routes carried forward:** the interview-report empty-render fix (determine seed-plan vs flag-demopatch → fix under 0 platform edits → add the alignment assertion → green) — lands intv-voice-pass too. Then the 2 voice cells (Bunny-key check → video or presence-only 47/47) + re-prove gate (b); then gates (c)/(d)/(f)/(h).
**Lessons:** gate (g) is a real defect surfaced live, exactly the class M235 deferred as "unproven." A "renders non-empty" check (M236) is NOT a plan-section-match proof — the report can pass the weak check (attempt row) while the actual report page renders empty. Always drive the surface the gate names (behind "View Report"), not the list that links to it.
