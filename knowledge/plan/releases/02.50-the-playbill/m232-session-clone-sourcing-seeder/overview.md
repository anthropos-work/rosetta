---
milestone: M232
slug: session-clone-sourcing-seeder
version: v2.5 "the playbill"
milestone_shape: section
status: planned
created: 2026-07-19
last_updated: 2026-07-19
depends_on: M231
delivers: corpus/ops/demo/session-clone-spec.md
---

# M232 — session clone sourcing seeder

**Status:** `planned`  ·  **Shape:** `section`  ·  **Complexity:** large  ·  **Depends on:** M231

## Goal
Build the seeder that COPIES real production sessions, anonymized where possible, re-tenanted into a manifest org, non-manager-played, and source-pinned by prod session-id — the deterministic realization of 'clone real sessions'.

## Scope
### In
- Read the selected real prod sessions (via the db-access read path, at authoring time) and reconstruct the full seedable result substrate per session in the target org (jobsimulation.sessions + public.local_jobsimulation_sessions mirror + validation_attempt_results/_skill_results/_criterion_results + actors/interactions transcript + interview_extraction_results.user_report/manager_report), passed + not-passed, all G14-valid enums
- Anonymize where possible (structured fields scrubbed; free-text handled per M231's contract)
- Net-new CODE (in-process Judge0 via `jobsimulation/internal/runner/` — the standalone roadrunner service is retired; corrected per M231 KB-5) + DOCUMENT (upload/PDF / `CollaborativeAsset`) assessment modalities
- Enforce owner-is-player-vantage, never a manager seat; pin the prod source session-id + anonymization transform in seed-generation-manifest.yaml (deterministic reseed)
- **Enable the interview render flags in the demo (added M231 D3, Fate-3):** the interview player+manager result surfaces gate on `posthog.isFeatureEnabled('flag_interview_{player,manager}_report')` (`AISimulationResultContainer.tsx:499-506`) — a seeded `interview_extraction_results` row is necessary but NOT sufficient. This seeder's demo config must turn these two PostHog flags ON (demo PostHog bootstrap, or a sha-pinned `demopatch` forcing `isFeatureEnabled` true for the two interview flags). See `content-stories-routes.md` § Interview.
- **Source only public-anchored sessions (added M231 D6):** the sourcing query MUST inner-join `directus.simulations` on the public predicate (`private=false AND tenant_id IS NULL AND status='published'`) so each cloned session's `sim_id` resolves in the demo (already snapshot-replayed). Pin by `jobsimulation.sessions.id`. Note: the platform's own `jobsimulation clone-session --session-id --user-id` (`cmd/clone_session.go`, running the built binary = within the zero-edit wall) re-players to a new userId but does NOT anonymize free-text or re-tenant `organization_id` — a rext layer still owns anonymize + re-tenant + the mirror co-write.
- AMEND corpus/ops/safety.md Part 3 to the honest posture (content-story demos carry anonymized-real session data, VPN/tailnet-scoped, source-pinned — the 'nothing behind the door' guarantee gains a documented, bounded exception)

### Out
- Reading/copying a customer session into a WIDER-than-VPN exposure
- The manifest projection + cockpit tab (M233/M234)
- Playable voice/Chime/LiveKit recording artifacts (transcript-only)
- AI-labs sessions unless M231 ruled them feasible

## Delivers
`corpus/ops/demo/session-clone-spec.md`

## Open questions
- Reuse existing hero seats as players or mint per-session anonymized player seats (brief leans mint; each must map to a real seeded public.users row)?
- Are a synthesized/scrubbed transcript + code-submission + document sufficient, or must a real recording be playable (DEF-M10-01)?

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.5 "the playbill" for the authoritative milestone design + the release-level decisions/risks (research provenance: `.agentspace/scratch/roadmap-research-2026-07-19` via the design-content-stories-research workflow).
