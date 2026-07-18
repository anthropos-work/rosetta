---
milestone: M235
slug: prove-it-lands
version: v2.5 "the playbill"
milestone_shape: iterative
status: planned
created: 2026-07-19
last_updated: 2026-07-19
depends_on: M234 (+ M230 for the academy section)
exit_gate: "On a cold reset-to-seed, every in-scope (session x action) logs in on the correct org and lands on a NON-EMPTY result page for BOTH player and manager vantages, 0 ejects, with the assessment 2-voice / 1-code / 1-document PASSED set present and each type present in passed AND not-passed states; each product either passes or is declared with a documented fate (AI-labs feasibility answered explicitly)."
iteration_protocol_ref: corpus/ops/demo/playthroughs.md + corpus/ops/demo/coverage-protocol.md
delivers: none
---

# M235 — prove it lands

**Status:** `planned`  ·  **Shape:** `iterative`  ·  **Complexity:** large  ·  **Depends on:** M234 (+ M230 for the academy section)

## Goal
Populate the tab with INTERESTING (not boring) real-shaped sessions per the brief and prove every cockpit action lands on a non-empty, believable result page.

## Exit gate
On a cold reset-to-seed, every in-scope (session x action) logs in on the correct org and lands on a NON-EMPTY result page for BOTH player and manager vantages, 0 ejects, with the assessment 2-voice / 1-code / 1-document PASSED set present and each type present in passed AND not-passed states; each product either passes or is declared with a documented fate (AI-labs feasibility answered explicitly).

**Iteration protocol:** `corpus/ops/demo/playthroughs.md + corpus/ops/demo/coverage-protocol.md`

## Scope
### In
- Pick interesting source sims per type; seed the full required set: training/assessment/hiring/interview each full+complete with PASSED and NOT-PASSED, assessment passed = 2 voice + 1 code + 1 document
- Drive every (session x action) end-to-end via a Playthrough (reuse hero-login.ts) + a coverage descriptor asserting the result page renders real content with non-zero values
- Triage each blank/wrong landing to its true read-model (the M219/M222 mirror-table-vs-base-session trap) and fix in seeder/manifest, or route to a demo-patch / escalate; verify no manager-played, deterministic reseed, 0 prod-ejects

### Out
- Live-on-billion proof (M236)
- Products M231 ruled out

## Open questions
- If /sim/.../result/<sessionId> is runtime-blank, is landing 'as player' on the seedable /profile/activities|/profile/skills composed outcome acceptable, or is a demo-patch authorized?
- Does not-passed render a meaningful result page or blocked/empty?

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.5 "the playbill" for the authoritative milestone design + the release-level decisions/risks (research provenance: `.agentspace/scratch/roadmap-research-2026-07-19` via the design-content-stories-research workflow).
