**Type:** tik — gate (g) fix, under TOK-01. Run 3, tik 1.

# iter-05 — progress

## ROOT CAUSE (proven)
The interview report renders empty because **`directus.simulations_extraction` has 0 rows on the demo** — the snapshot **never captured that table**. The render chain (`AISimulationResultContainer.tsx`):
- `isExtractionEnabled = isInterview && (flag_interview_manager_report || flag_interview_player_report)` (`:499-505`) — **both flag demopatches confirmed baked** in the built bundle, so the gate is open.
- `useGetSimulationExtractionSchema(simulation.id)` → `simulationExtractionSchema` (a JSON string) → `JSON.parse` → **`interviewExtractionPlan`** (`:519-523`). The schema is stored in **`directus.simulations_extraction.schema`**.
- `useGetInterviewExtractionReport(sessionId)` → **`interviewExtractionData`** (the seeded `interview_extraction_results`, present).
With the plan table empty, `interviewExtractionPlan` is null → the report has no sections to render → empty. Prod HAS the plan (sim `6d6cdf39` → a 65 KB extraction plan v1.4 with `sections` depth/adoption/sentiment); the demo lacked it purely because the capture surface omitted the table.

## FIX (committed + tested + pushed)
Added `simulations_extraction` to the snapshot **directus capture surface** (`stack-snapshot/directus/directus.go` `Surface()`), parent-scoped via `jobsimulation → simulations` (public-only; a tenant sim's plan never leaks). Columns id/date_created/date_updated/jobsimulation/schema. Golden surface tests updated + green (14→15 tables, 8→9 parent-scoped, + a direct-scope assertion). **rext `e74e563` on main; tag `sound-check-m244-content-sweep-robustness` moved + pushed to origin.** 0 platform edits.

## Live effect DEMONSTRATED (partial)
Pulled the real plan from prod (via the postgres MCP → a saved file → base64 → an inline-SQL `INSERT`, avoiding a 65 KB context pull) and loaded it into the demo's `directus.simulations_extraction`. The interview report render then changed from **empty (mainLen 0) → a rendered shell (mainLen 520)**: "Simulation Results / Nadia Ferrari / INTERVIEW AI Readiness Interview / Coming Soon [skills eval] / Evaluation Materials / Explore Key Moments / Explore Full Conversation." **So the plan was the load-bearing missing piece.** The full extraction sections (depth/adoption/sentiment) are behind "Explore Key Moments" — a deeper surface whose render I did not confirm (the drive timed out).

## Two things this iter SEPARATED (important)
1. **The gate-(g) extraction REPORT** (plan sections) — fixed at the root by the capture change; report shell now renders. Full-section render + the alignment assertion routed.
2. **The intv-voice-pass/PLAYER gate-(b) residual is a DIFFERENT surface** — the terse-ack page `/sim/<slug>/result/<sessionId>` (no reportId), which still fails "no interview acknowledgement text." Re-running gate (b) after the plan load stayed **46/49** — so this residual is UNRELATED to the extraction plan (likely a pass-vs-fail player-interview render / a sweep-assertion mismatch: a PASSED interview shows the result shell, not the ack the sweep's regex expects). Routed forward as its own item.

## Close — 2026-07-22

**Outcome:** gate (g) root cause PROVEN + the durable capture fix committed/tested/pushed + the plan shown load-bearing live (report render 0→520). NOT green: the full section render (behind "Explore Key Moments") unconfirmed, the alignment assertion not added, the live capture-path not exercised (data loaded manually), and the intv-voice-pass/player residual is a SEPARATE ack-page issue. Gate (b) stays 46/49.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET (gate g not green; 2/8 overall; gate b 46/49)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (iter-03 made progress; iter-04 closed-no-lift is excluded from the streak; only iter-05 is a no-prog tik) — (3) re-scope: n — (4) user-blocker: n (route-forward-able) — (5) cap-reached: n (tik 1 of run-3) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (capture-surface fix, not a seed), D2 (manual data-load proof vs live-capture-path) — iter-05/decisions.md
**Side-deliverables:** none beyond the capture fix (0 platform edits).
**Routes carried forward:** (g-full-green) exercise the live capture path (re-pin billion to the fixed rext → re-capture directus → re-replay → the extraction sections render from the REAL capture) + verify the sections behind "Explore Key Moments" + add the plan-section-id alignment assertion. (b-residual) the intv-voice-pass/player ack-page issue (separate; likely a sweep-assertion fix for passed interviews). (voice cells) the Bunny-key check still pending.
**Lessons:** (1) a "renders empty" interview surface can be a MISSING CONTENT TABLE in the capture surface, not a flag/seed bug — check what the render's data queries actually read (here: directus.simulations_extraction) and whether the snapshot captures it. (2) the interview has THREE distinct surfaces (player ack page / manager attempt-list / the extraction report behind "Explore Key Moments") — do not conflate them; the gate-(b) sweep and gate (g) touch different ones.