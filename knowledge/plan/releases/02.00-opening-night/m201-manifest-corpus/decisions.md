# M201 Decisions

Decisions with rationale (recorded during the user-guided curation). Design-time decisions live in
[`overview.md`](overview.md) + the consolidated capability spec
[`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../spec-drafts/playthroughs/spec.md) (v0.3).

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| D1 | **Product spine = 9 active products** (onboarding, skill-paths, ai-simulations, profile-skills, workforce-intelligence, org-admin-settings, assignment-monitoring, studio, talk-to-data). | The goal-aligned must-work surface, walked + locked with the user one product at a time. Hiring-recruiting (beyond onboarding), billing, academy-internal, mobile = OUT. | 2026-06-28 |
| D2 | **Onboarding = the only login-adjacent flow in scope** — 4 flavors (individual, workforce-standard {self-import + manager-prepared}, workforce-ai-readiness, hiring). Login/signup + checkout themselves are OUT. | The user: "we don't care about covering login/signup: only the onboarding of a user matters." Hiring onboarding = "the user lands AND has an assigned hiring simulation." | 2026-06-28 |
| D3 | **AI Simulations are non-voice this release** (chat / code / interview); voice (LiveKit) deferred to next release / M206. | User scope; voice needs a mirror tier. Interview is a sim_type on the chat engine, not a distinct surface (verify-confirmed). | 2026-06-28 |
| D4 | **AI Labs DEFERRED** (not an active product). | AI-Labs-class trap: backend `lab.v1.LabSessionService` persists a row but the labs-api client is nil (no VM boots), and there is no frontend launch surface. Nothing to be-the-human against. User: "defer to next release." | 2026-06-29 |
| D5 | **All 4 open flags RESOLVED by the adversarial verify** (code evidence): import → KEEP (`/reimport-profile` real); growth → VIEW-ONLY (target-role setter unwired); WI → ADD organization-feedback story; feature-config → KEEP (`/enterprise/settings` is a real multi-toggle). | Each flag was a runnability/route question the verify settled against the live clones. | 2026-06-29 |
| D6 | **`skill-paths.academy.UC1` LEFT AS-AUTHORED** despite a not-runnable verdict in the demo. | The verify found the academy not runnable in the demo (anonymous launch + no assessment module + unseeded catalog). User: **"leave here"** — re-decide in the v1.10 backfill rather than defer now. | 2026-06-29 |
| D7 | **`onboarding.enterprise-workforce-ai-readiness.UC1` KEPT — the verify verdict is a STALE-CLONE FALSE NEGATIVE.** | User: "the code exists … we have customers already using them right now." The verify grounded against next-web @ v2.33.2 (115+ commits behind prod), where the member surface isn't yet present. Re-ground post-sync. | 2026-06-29 |
| D8 | **Adversarial fidelity verification adopted as the close-quality gate** (11 agents, workflow `wvpnpvozh`) before accepting the sign-off. | Confirmation bias hides non-runnable flows (the AI-Labs precedent). The pass caught a 2nd trap + the staleness — paying for itself. | 2026-06-29 |
| D9 | **STALE-CLONE DISCOVERY → v2.0 PAUSED for a v1.10 backfill.** stack-demo clones are 5 weeks / 115+ commits behind prod; corpus + clones never caught up to what shipped. | All downstream (demo, seeders, corpus, the verify itself) is graded against stale code → systemic drift. The user opens a dedicated backfill (re-sync + re-ground + re-validate) before resuming v2.0. M201 corpus preserved as the v2.0 spec. | 2026-06-29 |
| D10 | **Close M201 on `release/02.00-opening-night` + merge that branch to `main`** (no tag, release NOT closed). | No `m201/*` branch was cut (worked on the release branch). User: "close properly … make sure to merge this mston/branch on main." Merging consolidates the M201 work + v2.0 scaffolding onto main so the backfill starts from it; v2.0 stays paused-not-shipped (no `v2.0` tag). | 2026-06-29 |

## Adversarial review (close)

The corpus was itself the artifact under adversarial review (workflow `wvpnpvozh`): every use-case re-grounded
against the live clones by an independent verifier + a completeness critic + a coherence critic. Scenario it was
built to catch — *a use-case that looks real but isn't runnable* (the AI-Labs class). It found two more
(academy; and a third trap-shaped finding, the unwired `createUserTargetRole`), and — critically — surfaced that
the **grounding baseline itself was stale**, which is the more important finding: it means the verify's own
negative verdicts are not yet trustworthy and must be re-run against fresh clones (handed to the backfill).
