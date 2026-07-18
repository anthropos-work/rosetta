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
- Net-new CODE (roadrunner) + DOCUMENT (upload/PDF) assessment modalities
- Enforce owner-is-player-vantage, never a manager seat; pin the prod source session-id + anonymization transform in seed-generation-manifest.yaml (deterministic reseed)
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
