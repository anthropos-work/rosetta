---
iteration: 04
iteration_type: tik
milestone: M203
status: closed-fixed
created: 2026-07-02
---

# iter-04 — TIK (Skill Paths — legacy learn & progress)

## Active strategy reference
**TOK-01** step 2 (Skill Paths — the mutating flows tier, after the deterministic Profile reads).

## Cluster / target identified
iter-03 routed forward Skill Paths (skill-paths.legacy.UC1) as the next GATE journey. Re-survey (Phase 1
Step 0): baseline 4/4 profile UCs passing; probed the real UI — the Skill Paths library is at
`/library/skill-paths` ("Skill Paths Library") with a rich replayed catalog (real path slugs), a path detail
(`/skill-path/<slug>`) with a "Start" CTA + 5 chapters + "0% complete", and Start → the chapter player
(`/skill-path/<slug>/chapter?...`) with a real step + "Mark complete & continue". The path's ASSESSMENT step
(Chapter 05 "Test Yourself") is a VOICE sim ("voice call") → M206 tier, OUT this release.

## Hypothesis
The browse→open→Start→progress leg is fully drivable + deterministic (NON-voice); a Playthrough that opens the
library, opens a real path, Starts it, lands in the player on a real step, and asserts the advance control is
present will PASS. The verify-skill terminal composes with an assessment (P7) — proven on the profile side.

## Expected lift
+1 employee use case (skill-paths.legacy.UC1) — the 2nd of the 3 gate journeys; no regression.

## Phase plan
Declare skill-paths.legacy.UC1 (new skill-paths.yaml product; add `public-catalog` seed-world capability) →
add SkillPathPage page-object → build the spec → run → reconcile.

## Escalation conditions / acceptable close-no-lift
A path undrivable without a platform edit → unimplementable (escalate). The voice-assessment terminal is the
KNOWN boundary (spec §5.8) — asserting at the progress boundary is the sanctioned NON-voice posture, not a
close-no-lift.

## Close → see progress.md
