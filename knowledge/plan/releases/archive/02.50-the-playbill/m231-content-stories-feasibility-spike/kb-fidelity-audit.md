---
title: "KB Fidelity Audit — M231 content-stories-feasibility-spike"
date: 2026-07-19
scope: milestone:M231
invoked-by: build-milestone
---

## Verdict
**YELLOW** — proceed with tracking. No blind area blocks the spike; the milestone's own deliverable
(`corpus/ops/demo/content-stories-routes.md`) is the consolidation home for the result-route/read-model knowledge
the topic docs are silent or stale on. All stale claims were verified against code by the spike's own discovery, so
the deliverable is built on verified facts, not on the stale docs. Findings tracked as `KB-1..KB-8` in `decisions.md`.

## Topic Inventory
| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Sim session/result read-model | `corpus/services/jobsimulation.md` | `stack-dev/{jobsimulation,next-web-app,app}` | PAIRED (doc SILENT on read-model — see KB-1) |
| Hiring render path (player+manager) | `corpus/services/hiring.md` | `stack-dev/next-web-app/apps/{web,hiring}`, `stack-dev/app` | PAIRED |
| Skill-path session/result | `corpus/services/skillpath.md` | `stack-dev/skillpath`, `stack-dev/next-web-app`, `stack-dev/app` | PAIRED (incomplete — KB-4) |
| Assessment modalities | `corpus/architecture/ai_architecture.md` | `stack-dev/{jobsimulation,roadrunner,app}` | PAIRED |
| Code-exec engine | `corpus/services/roadrunner.md` | `stack-dev/jobsimulation/internal/runner` | PAIRED (stale — KB-6) |
| Doc conversion | `corpus/services/gotenberg.md` | `stack-dev/app/internal/converter` | PAIRED (aligned) |
| AI-labs LabSession | `corpus/services/backend.md` | `stack-dev/app/internal/labs/session` | PAIRED (path stale — KB-8) |
| Academy session/progress | `corpus/services/ant-academy.md` | `stack-dev/ant-academy`, `stack-dev/app/internal/data/ent/schema/academy_*` | PAIRED (stale — KB-7) |
| Prod read path | `corpus/ops/db-access.md` | wired `postgres` MCP (`marco_read`/`10.2.22.13`) | PAIRED (aligned) |
| Demo exposure/anon posture | `corpus/ops/safety.md` Part 3 | tooling | PAIRED (current) |
| Content-stories result-route map | `corpus/ops/demo/content-stories-routes.md` | — | DOC-ONLY (M231 DELIVERS it — expected) |
| Session-clone sourcing | `corpus/ops/demo/session-clone-spec.md` | — | DOC-ONLY (M232 delivers — expected) |

## Fidelity Findings
1. **KB-1 — `jobsimulation.md` is SILENT on the session/result read-model + the `local_jobsimulation_sessions` mirror.** Verdict UNVERIFIABLE-by-absence → the content the spike leans on lives in `seeding-spec.md` + `hiring.md`. Fix owner: M231's `content-stories-routes.md` becomes the consolidation home; add a pointer from `jobsimulation.md`. Not blocking (knowledge exists elsewhere).
2. **KB-2 — `jobsimulation.md` ports "8400/8401"** conflict with the repo's own CLAUDE.md ("8080/8081"). Verdict STALE, tangential (M231 does not depend on ports). Track; do not fix under M231 (offset-port nuance needs its own verification).
3. **KB-3 — `hiring.md` § "The render probe is intercepting-route-aware (M228)"** claims the recruiter drawer is a Next.js intercepting route at `…/@tabs/(.)ai-simulations/[simId]`. Verdict STALE — verified: ZERO intercepting-route dirs exist in `apps/`; the drawer is a plain Ant `<Drawer>` (`InsightsByMembersContainer.tsx:359`) on the ordinary `[simId]/page.tsx` leaf. **FIXED INLINE** (surgical correction; kept the `.ant-drawer` detectability + first-sim-per-page-load probe narrative).
4. **KB-4 — `skillpath.md` omits the manager-side mirror.** The manager insights surface (`insightsSkillPathByMemberships`) reads the `app`-side `public.local_skill_path_session` MIRROR (`app/internal/organization/intelligence.go:997/1142`), NOT the skillpath runtime — same render-gate trap as hiring's `local_jobsimulation_sessions`. Verdict STALE/incomplete. Fix owner: add a mirror pointer to `skillpath.md` (done inline) + full treatment in `content-stories-routes.md`.
5. **KB-5 — `ai_architecture.md` code modality = "Judge0 via the Roadrunner service".** Verdict STALE — code execution is now IN-PROCESS in jobsimulation (`jobsimulation/internal/runner/runner.go` — comment: "formerly the standalone 'roadrunner' service"). Fix owner: architecture-doc pass; `content-stories-routes.md` states the correct fact. Track (not blocking; deliverable carries the correct fact).
6. **KB-6 — `roadrunner.md` describes a live jobsim consumer** (`ROADRUNNER_RPC_ADDR`). Verdict STALE — no `ROADRUNNER_RPC_ADDR` remains in jobsimulation; the standalone repo is orphaned. Track for an architecture-doc pass.
7. **KB-7 — `ant-academy.md` "no backend writes / read-only GraphQL client at runtime".** Verdict STALE — since v0.5 M2 the academy is a full read/**write** GraphQL client persisting progress/last-activity/bookmarks/certs to `app internal/academy`. This is LOAD-BEARING for the academy content-story verdict (academy IS seedable BECAUSE it's backend-authoritative). Fix owner: `ant-academy.md` correction (pointer added inline) + `content-stories-routes.md` carries the verified academy-session model.
8. **KB-8 — `backend.md` labs package path "`internal/labsession`".** Verdict STALE — actual path is `internal/labs/session` (+ `internal/labs/labsapi`, `internal/labs/adapter`); also `grade_result` exists on the Ent table but is NOT GraphQL-exposed. Track.

## Completeness Gaps
- The **result read-model as a persisted-vs-recomputed contract** was undocumented as a first-class topic anywhere. M231's deliverable fills this (the central spike output). Classified critical → it IS the milestone deliverable.
- The **manager-view mirror pattern** (`local_jobsimulation_sessions` / `local_skill_path_session`) is documented for hiring but not generalized. `content-stories-routes.md` generalizes it.

## Applied Fixes
- `corpus/services/hiring.md` — corrected the M228 "intercepting-route" mechanism claim to "plain Ant `<Drawer>` on the leaf route" (KB-3).
- `corpus/services/skillpath.md` — added a manager-side mirror pointer (KB-4).
- `corpus/services/ant-academy.md` — corrected the stale "no backend writes at runtime" claim to backend-authoritative-since-v0.5-M2 (KB-7).

## Open Items (require user decision)
None. AI-labs OUT is a spike verdict (no user decision needed). The INTERVIEW flag-gate + skill-path-player-blank demo-patch decisions are recorded in `decisions.md` and routed Fate-3 to M232/M235 (normal spike output).

## Gate Result
**YELLOW: proceed with tracking.** No blind area; no stale claim that the spike's deliverable would inherit as truth
(all corrected by verified discovery). KB-1..KB-8 tracked in `decisions.md`; KB-3/KB-4/KB-7 fixed inline; the rest
routed to their fix-owners (M231 deliverable + architecture-doc pass). `SEVERITY: warning`.
