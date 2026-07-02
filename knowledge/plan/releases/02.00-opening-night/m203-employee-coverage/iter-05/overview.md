---
iteration: 05
iteration_type: tik
milestone: M203
status: closed-fixed
created: 2026-07-02
---

# iter-05 — TIK (AI Simulations NON-voice — chat launch) + the sim-feature-enable diagnosis+fix

## Active strategy reference
**TOK-01** step 3 (AI Simulations — the integration-boundary tier). The 3rd + final gate journey.

## Cluster / target identified
iter-04 routed forward AI Simulations (NON-voice) as the last gate journey. Re-survey: the demo sim catalog
is heavily VOICE-typed ("Voice Call"/"Voice Mode" = M206 tier, OUT); found chat-based sims (interview-typed
sims reuse the chat container — M201 finding), e.g. `ai-tools-adoption-developer-interview` (described as a
"conversation"). Target: `ai-simulations.chat.UC1` at the §5.8 launch boundary.

## Hypothesis
A chat-based sim opens + launches to its interactive launch state; a Playthrough that opens the library,
opens the sim, clicks Start Simulation, and asserts the launch confirmation (§5.8 boundary) + the deny modal
absent will PASS.

## What actually happened (a diagnostic+fix iter)
The launch was BLOCKED for the seeded pt-employee: "You cannot start AI Simulations in this organization" — a
deep diagnosis (documented in decisions.md D1) traced it through the FE gate (`canStartAsOrganizationMember`
reads `userMembership.organizationFeatures`) → the GraphQL resolver (`IsMemberAllowedToUseFeature`) → the
Sentinel matcher (`OrgMembershipsAllowedToUseFeature` = the g3 grouping policy). The DB data was all CORRECT
(Pat's membership had the g3 FEATURE_JOB_SIMULATIONS grant + the g2 org grant). ROOT CAUSE: the running
Sentinel casbin enforcer CACHES its policy in-memory; the seeded g3 grants weren't reflected until an explicit
Reload. **FIX (confirmed): a post-seed Sentinel Reload** — added to `run-playthroughs.sh` (idempotent,
non-fatal, zero platform edits). After Reload, pt-employee launches the sim ("Welcome to your AI Simulation").

## Expected lift → realized
+1 employee use case (ai-simulations.chat.UC1) — the 3rd gate journey. Realized: 6/6 passing, no regression.

## Phase plan (as executed)
Declare ai-simulations.chat.UC1 (new ai-simulations.yaml; add `sim-feature-enabled` seed-world capability) →
add SimulationPage page-object → diagnose the launch deny → fix (runner Sentinel Reload) → build spec → run →
reconcile.

## Close → see progress.md
