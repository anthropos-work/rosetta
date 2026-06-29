# M203 Spec Notes

Technical notes accumulate here during the iter loop. The authoritative design lives in [`overview.md`](overview.md)
+ the consolidated capability spec
[`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../spec-drafts/playthroughs/spec.md) (v0.3). M203 is
**tooling + docs only — zero platform-repo edits**; an un-drivable surface escalates via
`unimplementable-without-platform-edit`, it never edits the platform.

## Scope (see overview.md for the authoritative gate)
The declared **employee-vantage** use cases — Maya's core journeys (declared in the **M201 manifest corpus**):
- **Skill Paths** — browse → enroll → complete → verify-skill.
- **AI Simulations** — chat/code launch → complete → score-in-range (**NON-voice**; assert at the
  launch/completion boundary per spec §5.8 — voice/recording is the M206 mirror tier, future).
- **Profile** — verified-skill chart + claimed-vs-verified gap + work/education timeline.

Each becomes one Playthrough (one use case ↔ one Playthrough), played as Maya via the M202 foundation.

## Reuse paths (the M202 foundation + the shared e2e foundation)
- The **M202 page-object layer** (the per-surface locator/landmark registry every Playthrough imports) — extend
  it with the employee surfaces; re-pin is O(surfaces).
- The **M202 dedicated decoupled seed** + the reset-to-seed serial-default runner.
- `stack-demo/rosetta-extensions/stack-verify/e2e/lib/cockpit-login.ts` — Maya's login (the M37 handshake).
- `stack-demo/rosetta-extensions/stack-seeding/` — the seeding machinery the reset rides.

## Per-vantage assertion notes (the AI-sim boundary)
- **Scoring is deterministic rubric-based** (0–100, EU-AI-Act — NOT an AI scorer; spec §5.8). Assert the score's
  **structure / range** ("score-in-range"), never AI prose.
- A **live-AI** sim flow is asserted at the **launch/completion boundary** (the flow launched + reached an
  interactive state + the outcome artifact materialized), NOT turn-by-turn (the only thing provable under P6 with
  a live LLM in the loop).
- AI/voice creds **are** provisioned on demo-N; the barrier for voice is the missing **mirror** + determinism,
  not missing wiring (spec §5.8). NON-voice chat/code is in scope here.

## Tag / two-repo state
TODO (iter loop): per-iter rext authoring commits + tags; consumption-clone checkouts; the corpus m203 branch.

## Open questions (resolve in the iter loop; record in decisions.md)
- Which surfaces need a landmark anchor vs play purely on semantic locators (discovered per-iter against the
  false-fail rate).
- The exact "verify-skill" outcome assertion for the async verified-skill projection (P1b: waiting for an async
  platform projection to settle is ordinary P2 + Playwright web-first waiting, NOT a backdoor).
- The sim launch/completion boundary's concrete outcome marker (a session row / a result / a completion marker).
