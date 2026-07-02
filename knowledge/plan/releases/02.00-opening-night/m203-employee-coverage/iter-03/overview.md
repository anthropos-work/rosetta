---
iteration: 03
iteration_type: tik
milestone: M203
status: closed-fixed
created: 2026-07-02
---

# iter-03 — TIK (Profile work/education timeline)

## Active strategy reference
**TOK-01** (Deterministic-read-first). This tik completes TOK-01's Profile step by landing the third
gate-relevant Profile piece.

## Cluster / target identified
iter-02 routed forward `profile.self-evaluation.UC1` (the WRITE flow). **Re-survey (Phase 1 Step 0) substituted
the target:** probing the self-rate flow found a drivability quirk — the "Re-rate skills" / per-skill "Edit"
controls have their click intercepted (pointer-events / overlay; the rate MODAL "Rate your skills as a DevOps
Engineer" exists in the DOM but stays hidden on click). More importantly, re-reading the M203 **exit gate** —
Profile = "verified-skill chart + claimed-vs-verified gap + **work/education timeline**" — self-eval is a
non-gate M201-corpus EXTRA, whereas the **work/education timeline** is the third GATE-relevant Profile piece
and still uncovered. A probe of the Career Profile tab confirmed it renders a real 3-entry work history
(Meridian Labs / Drift Analytics) + University of Edinburgh education (the M41 ProfileSeeder). Substituted
`profile.timeline.UC1` as this iter's target (still under TOK-01: a deterministic read); re-routed self-eval
to a later iter as a non-gate extra.

## Hypothesis
The work/education timeline renders real seeded entries on the Career Profile tab, so a Playthrough that opens
/profile and asserts the Work + Education sections + a real dated entry present will PASS deterministically.

## Expected lift
+1 employee use case (profile.timeline.UC1), completing the gate's Profile journey; no regression.

## Phase plan
Declare profile.timeline.UC1 (ptvalidate green) → extend ProfilePage with the Career-tab timeline accessors →
build the spec → run against demo-1 → reconcile (no-regressions).

## Escalation conditions / acceptable close-no-lift
Empty timeline despite the seed → diagnose seed-vs-platform drift (P6). A timeline un-assertable without
pixel/copy coupling → close-no-lift with falsification (but the probe found sanctioned semantic anchors).

## Close → see progress.md
