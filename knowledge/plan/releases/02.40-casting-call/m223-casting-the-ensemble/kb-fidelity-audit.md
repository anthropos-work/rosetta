---
title: "KB Fidelity Audit — M223 casting-the-ensemble"
date: 2026-07-15
scope: milestone:M223
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| The recruiter comparison read-model (the mirror-table score + the read-path) | `corpus/services/hiring.md` | `app/internal/organization/intelligence.go` (`InsightsJobSimulationByMemberships` :1681, `InsightsByJobSimulations` :1472), `app/.../resolver_queries.go` :1088-1089, `app/internal/data/ent/schema/local_jobsimulation_session.go` :52 | PAIRED |
| The `is_hiring` gate + `narrative: hiring` discriminator | `corpus/services/hiring.md`, `blueprint/stories.go` | `blueprint/stories.go` (`IsHiringOrg`, `HiringOrgID`), `seeders/org.go` :68 | PAIRED |
| The `OrgFeatureInsights` Casbin substrate | `corpus/services/hiring.md` § silent-403 | `sentinel/init_policy.sql` :54 (p3), `sentinel/internal/authorization/casbin.go` (m3), `seeders/identity.go`/`users.go` (g2 admin grant) | PAIRED |
| `SIMULATION_TYPE_HIRING` sims + (absent) `job_position` | `corpus/services/hiring.md`, `corpus/ops/snapshot-spec.md` | `cms/internal/directus/collections/jobsimulation.go` :741, `stack-snapshot/directus/directus.go` :28 | PAIRED |
| The verified-skill / session-pair seeder machinery (the 7-table fan-out) | `corpus/ops/demo/stories-spec.md`, `corpus/ops/seeding-spec.md` | `seeders/persona.go` :267-288, `seeders/persona_write.go` :68-141 | PAIRED |
| The funnel-seeder precedent (config + funnel split) | `corpus/services/ai-readiness.md` | `seeders/ai_readiness_config.go`, `seeders/ai_readiness_funnel.go`, `seeders/contentref.go` (reserved tail) | PAIRED |

## Fidelity Findings

1. **hiring.md — the score is `public.local_jobsimulation_sessions.score` (a Float32 mirror), read directly by the resolver.**
   - Source: `corpus/services/hiring.md` § "The comparison read-model (THE HEADLINE)".
   - Expected: the resolver reads ONLY `LocalJobsimulationSession`, `field.Float32("score")` at `local_jobsimulation_session.go:52`; best-attempt via `row_number() ORDER BY score DESC`.
   - Actual: `intelligence.go:1692` `m.ent.LocalJobsimulationSession.Query()`, `:1740` sort `FieldScore DESC`, `:1801` `score := RoundFloat(float64(ls.Score), 0)`; `local_jobsimulation_session.go:52` `field.Float32("score")`, `:39` `field.String("completition_status")` (misspelled column, as documented).
   - Verdict: **ALIGNED.**

2. **hiring.md — the AI-Simulations LIST derives from the mirror (no config table).**
   - Expected: sims appear because the org has candidate mirror-sessions on them.
   - Actual: `InsightsByJobSimulations` (`intelligence.go:1472,1527`) reads `LocalJobsimulationSession` org-scoped; no config/registry table involved. Verdict: **ALIGNED** (informs D2 — no hiring-config-table write is needed for the surface).

3. **hiring.md — the scoreboard sits behind `OrgFeatureInsights`; the admin role carries it globally.**
   - Expected: without the Casbin insights permission the query silent-403s (`resolver_queries.go:1089`); the seeder must replicate whatever grants the demo orgs the permission.
   - Actual: `resolver_queries.go:1089` `OrgCheckFeaturePermission(ctx, permission.OrgFeatureInsights, organizationID)` → enforce-context "3" → `m3 = g2(org,sub,'admin') && feat` matched against the GLOBAL `p3('default','admin','org:feature:insights')` (`init_policy.sql:54`). Verdict: **ALIGNED** — the hiring org's `admin` members inherit insights from their standard g2 grant (D1); no net-new grant.

4. **hiring.md — the scoreboard lives in the dockerized `apps/web` (not `apps/hiring`-only).**
   - Actual: `simulationScoreColumn.tsx` exists in BOTH `next-web-app/apps/web/...` and `next-web-app/apps/hiring/...` — the demo-built `apps/web` copy renders the score. Verdict: **ALIGNED** (M222 D1 render-proof holds).

5. **hiring.md — `directus.simulations.type` carries `SIMULATION_TYPE_HIRING`; `job_position` is absent from the capture.**
   - Actual: `cms/.../jobsimulation.go:741` `type ... oneof=... SIMULATION_TYPE_HIRING ...`; the capture's public filter is `private=false AND tenant_id IS NULL AND status='published'` (`stack-snapshot/directus/directus.go:28`); M222 measured 87 captured hiring sims / 0 captured `job_position`. Verdict: **ALIGNED** (confirms the type-aware reader is feasible against the snapshot; S3 stays dropped).

## Completeness Gaps

1. **The 4th (hiring) story, the `HiringConfigSeeder`, the `HiringFunnelSeeder`, and the type-aware hiring-sim reader do not exist yet** — this is the M223 net-new work, NOT a blind area. The framework docs (stories-spec, seeding-spec) and the read-model contract (hiring.md) exist; M223 delivers the new seeders + the doc extensions (S7: stories-spec.md + seeding-spec.md + the snapshot-spec job_position-drop note). Tracked as the milestone's own deliverables (overview § Delivers), not a fidelity gap.

## Applied Fixes
None required — all audited claims are ALIGNED, all cited line-anchors resolve within the current files (`intelligence.go:1692`, `resolver_queries.go:1089`, `local_jobsimulation_session.go:52`, `init_policy.sql:54`, `jobsimulation.go:741` all confirmed).

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to Phase 1. Every M223-dependency topic is PAIRED and every load-bearing claim in `corpus/services/hiring.md` (the M222 contract the milestone builds against) is ALIGNED with the READ-ONLY platform clones. The net-new seeders + doc extensions are the milestone's deliverables, not blind areas.
