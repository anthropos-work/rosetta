---
iter: 06
milestone: M244
iteration_type: tik
status: closed-fixed
active_strategy: TOK-01
created: 2026-07-22
---

# iter-06 — gate (g) full green: the plan-section-id alignment assertion (S-8/S-9)

**Type:** tik under **TOK-01** (staged billion bring-up → gate parts a–h one-cluster-per-tik).

## Phase 0 type selection
TIK. Not bootstrap (iter-06). Not a triggered tok — the no-prog streak is < 3: iter-03 moved gate (b)
37→46 (measurable), iter-05 root-caused + capture-fixed gate (g) with a live render 0→520 (measurable);
iter-04 was closed-no-lift-falsification (excluded from the streak). Default: tik under TOK-01.

## Active strategy reference
TOK-01 — discharge exit-gate parts in dependency order, one cluster per tik, asserting from a tailnet peer.
This iter targets gate **(g)** "interview plan-section-id alignment assertion added + green (S-8/S-9)."

## Cluster / target identified
iter-05 landed the ROOT-CAUSE fix (capture `directus.simulations_extraction`, rext e74e563) and proved the
plan load-bearing (render 0→520), but gate (g) is NOT green: the alignment assertion was never authored.
Fresh evidence this iter (read-only prod, non-PII section-id metadata):
- The interview sim `6d6cdf39` (`ai-readiness-interview-d62`) plan is **v1.4 +v1.4-mastery-axes** with 14
  sections: 12 manager-scope (success_score, usage_profile, adoption, transformation, originality, depth,
  ownership_governance, multiplier_blocker, top_concerns, top_unexpected, recurring_themes, sentiment),
  1 summary-scope (summary), 1 player-scope (candidate_wrapup).
- **`intv-voice-pass`** (session `cba53b09`) is PERFECTLY aligned: all 12 manager_report keys ∈ plan.
- **`intv-voice-fail`** (session `43a92fc0`) is **MISALIGNED** — its report data carries 3 orphan keys
  (`breadth`/`context_fit`/`frequency`) from an **older plan version (v1.3)**; 5 v1.4 manager sections have
  no data. This is exactly the S-8/S-9 defect: a source-pinned session whose report was produced under a
  plan version that has since evolved → orphaned data (lost on render) + empty plan sections.

## Hypothesis
Building a strict plan-section-id alignment assertion (report-data keys ⊆ captured-plan section-ids, per
scope) will (a) CATCH the intv-voice-fail version drift and (b) turn green once intv-voice-fail is re-pinned
to a v1.4-aligned FAILED interview session. Candidate: **`05dae0f7`** — same sim, completion_status=failed,
score 0 (matches the current pin), 11 in-plan manager keys, 0 orphans, italian, candidate_wrapup player key.

## Planned scope (deliverables)
1. **Re-pin `intv-voice-fail`** `content-sessions.yaml` `43a92fc0 → 05dae0f7` (+ metadata: interaction_count
   0→1); regenerate `intv-voice-fail.json` via `content-capture --only intv-voice-fail --dsn <read-only prod>`
   (scrubbed, read-only, the designed re-pin mechanism; `~/.pgpass` + Tailscale confirmed reachable).
2. **The alignment assertion** in rext `contentsession/`: a checked-in golden of the captured interview plan
   section-ids per scope + `TestInterviewPlanSectionAlignment` asserting, per interview fixture:
   (i) plan non-null ≥1 section [catches iter-05 null-plan], (ii) manager_report.results keys ⊆ plan manager
   ids [strict alignment / version-drift catch], (iii) user_report.results keys ⊆ plan player ids.
3. `go test` green (contentsession + the touched packages) + cleanliness (surviving-token) green on the new
   fixture.
4. Tag + push rext (move the `sound-check-m244-content-sweep-robustness` consumption tag; push main + tag).
5. **Live-render confirmation** on billion: the interview extraction sections render behind "Explore Key
   Moments" for the aligned `intv-voice-pass` session (extends iter-05's 0→520 shell to the deep sections).

## Expected lift
Gate (g) discharged → metric **2/8 → 3/8**. Assertion added + green; the real version-drift defect fixed
(not allowlisted).

## Phase plan
Phase A: re-pin + re-capture (read-only prod). Phase B: author assertion + golden. Phase C: go test gate.
Phase D: tag+push. Phase E: live-render confirm on billion (peer Playwright, settle+retry template).

## Escalation conditions
- content-capture cannot connect to prod / the new fixture fails cleanliness → FALL BACK (orchestrator-blessed):
  keep the assertion, prove it green for the aligned pass fixture, DOCUMENT the intv-voice-fail drift + route
  the re-pin forward (Fate-3), record the capture-path limitation honestly. Do not fabricate a live capture.
- Any platform-repo edit demanded → ESCALATE (blocker); route to a sha-pinned demopatch instead.

## Acceptable close-no-lift outcomes
If prod capture is infeasible AND the assertion cannot be made honestly green without the re-pin, close with
the assertion authored + the drift characterized + the re-pin routed forward (a complete falsification of the
"just add the assertion" premise). But the primary path (re-pin feasible) is expected to land full green.

## Routes that may carry forward
- The LIVE directus capture-path exercise (re-pin billion rext → re-capture → re-replay) — gate (g)'s
  capture-path proof; may route to the gate-(h) cold reset-to-seed.
- gate (a) ORG-CLEAN LIVE re-verify for the re-pinned fixture at the next cold reset-to-seed.
