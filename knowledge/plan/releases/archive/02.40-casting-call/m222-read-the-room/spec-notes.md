# M222 — Spec notes

_Technical detail accumulated during the milestone: file:line maps, command transcripts, measured numbers._

## Hiring model (for `corpus/services/hiring.md`)
- **Org-type gate — a DUAL-WRITE.** Backend: `public.organizations.is_hiring boolean NOT NULL default false`. Client:
  `isHiringOrg = Boolean(organization.publicMetadata?.isHiring)` + `organizationId = organization.publicMetadata?.eid`
  (`apps/web/src/hooks/useGetClerkOrganization.tsx:20-21`). Both must be set (D3).
- **The `candidate` membership role** rides the existing `role_mix.candidate` in the blueprint (RoleMix already
  carries `Candidate`). The hiring org is `RoleMix ≈ 0.1 admin / 0.9 candidate`, no `member` (M223).
- **Hiring sims** = `SIMULATION_TYPE_HIRING`. The (optional) `job_position` relation is `JobSimulation.jobPosition`
  — **absent in the captured snapshot** (0 rows) and **unused by the comparison scoreboard** (D4).

## Render-probe results (BA-1 / BA-2 / BA-3 / BA-6) — S3, done on the live `billion` substrate
- **BA-3 → GO.** The comparison surface `/enterprise/activity-dashboard → AI-Simulations → [simId]` is in the
  **dockerized `apps/web`** (`@anthropos/web-app`), renders from DATA ALONE, and **survives the `is_hiring` flip**
  (relabel-only: `useNavbarSections.tsx:300-307`). No escalation. R2 retired.
- **BA-1 → the mirror-table contract.** The score is `app.public.local_jobsimulation_sessions.score` (Float32), NOT
  `jobsimulation.sessions.score`. Read-path: `simulationScoreColumn.tsx:54,95-97` → `insights.ts:31-82` →
  `resolver_queries.go:1088,1134` → `IntelligenceManager.InsightsJobSimulationByMemberships` reads ONLY
  `LocalJobsimulationSession` (`intelligence.go:1692`), `row_number() ORDER BY score DESC` (`:1728-1735`), `ls.Score`
  (`:1801`); Ent `local_jobsimulation_session.go:52` `field.Float32("score")`. **The list score needs NO per-session
  `validation_*`/eval row** — those (`validation_attempt_results`/`_skill_results`/`_criterion_results`) are the
  DRILL-DOWN (BA-4), `persona_write.go:69-71,143-167`.
- **BA-2 → enumerated.** `isEnterprise` divergence: nav `Boolean(organization)` (`template.tsx:90`, stays TRUE) vs
  billing `!isHiringOrg && organizationId` (`FreeTrialContainer.tsx:29`, flips FALSE). Under `is_hiring` the nav
  relabels activity-dashboard "Results", trims library to AI-Simulations, hides some non-admin member surfaces,
  gates Workforce Intelligence off. The comparison surface itself is unaffected (relabel only).
- **BA-6 → GO.** Captured snapshot has **87 real `SIMULATION_TYPE_HIRING` sims** (published + public) — pick 5 as
  the positions. `directus.job_position` = **0 rows captured** (443 was a PROD count) → **no job_position replay**
  (D4, an M223 Scope.In refinement).

## The seeder-output contract (the GO/NO-GO record)
- **Minimal write-set per (candidate × sim):** (1) `public.local_jobsimulation_sessions` (score source;
  `completition_status` is the misspelled column); (2) `jobsimulation.sessions` (the federated `Session!` pair — else
  NULL-bubble; **393/393 local rows on `billion` have the matching pair**); (3) `public.memberships` active.
- **Org prereqs:** `organizations.is_hiring=true` + the **`OrgFeatureInsights` Casbin permission** (else silent 403
  at `resolver_queries.go:1089`).
- **Sort/cohort:** one best-attempt row per `user_id` (highest score), `score DESC, completition_status ASC,
  session_started_at DESC` (`intelligence.go:1738-1751`); same `jobsimulation_id`+`organization_id` = one cohort.
- **PersonaSeeder already writes exactly this pair** (`persona_write.go:68-73`, `sessionCols()`/`localSessionCols()`
  `:125-141`). M223 = a 45×5 generalization, NOT net-new.

## `is_hiring` gate thread (S2 — landed in M222)
- `blueprint.StoryOrg.IsHiring bool` (yaml `is_hiring`) — the explicit gate field (`blueprint/blueprint.go`).
- `blueprint.HiringNarrative = "hiring"` + `ResolvedStory.IsHiringOrg()` = `IsHiring || Narrative == "hiring"` — the
  discriminator recognition, sibling to seeders' `aiReadinessNarrative` (`blueprint/stories.go`).
- `ResolvedStory.IsHiring` threaded from `StoryOrg.IsHiring` in `EffectiveStories`.
- `seeders/org.go` now writes `st.IsHiringOrg()` into `public.organizations.is_hiring` (was hardcoded `false`).
- Reserved: `blueprint.HiringStoryID = "hiring"` → `blueprint.HiringOrgID()` = `StoryOrgID("hiring")` (the
  deterministic hiring-org UUID M223/M224 reference).
- Also projected into the auditable manifest: `manifest.Org.IsHiring` (`yaml:is_hiring,omitempty`).
