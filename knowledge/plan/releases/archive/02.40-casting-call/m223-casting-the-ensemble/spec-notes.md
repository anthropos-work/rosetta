# M223 — Spec notes

_Technical detail accumulated during the milestone: file:line maps, command transcripts, measured numbers._

## Pre-flight audits — S0 (KB-fidelity)
`/developer-kit:audit-kb-fidelity --milestone=M223` → **GREEN** (2026-07-15, sha at audit `319510d`).
Report: `kb-fidelity-audit.md`. Every M223-dependency topic PAIRED; every load-bearing claim in
`corpus/services/hiring.md` ALIGNED with the READ-ONLY platform clones (mirror-table score, insights
Casbin substrate, `SIMULATION_TYPE_HIRING`, `apps/web` render). Net-new seeders + doc extensions are the
milestone's deliverables, not blind areas.

### Topic → doc → code triples (fast-start for later audits)
- read-model → `corpus/services/hiring.md` → `app/internal/organization/intelligence.go:1681,1472` + `resolver_queries.go:1088`
- score source → `hiring.md` → `app/internal/data/ent/schema/local_jobsimulation_session.go:52` (`field.Float32("score")`), `:39` (`completition_status`)
- insights gate → `hiring.md` → `sentinel/init_policy.sql:54` (`p3 admin org:feature:insights`) + `casbin.go` (m3)
- hiring-sim type → `hiring.md`/`snapshot-spec.md` → `cms/internal/directus/collections/jobsimulation.go:741`
- session pair → `stories-spec.md`/`seeding-spec.md` → `seeders/persona.go:267-288`
- funnel precedent → `ai-readiness.md` → `seeders/ai_readiness_config.go` + `ai_readiness_funnel.go` + `contentref.go`

## Implementation landmarks (rext stack-seeding)
- Session-pair builder to REUSE: `seeders/persona.go:267-288` (`jobsimulation.sessions` + `public.local_jobsimulation_sessions` mirror), cols `persona_write.go:125-141`.
- G14-valid session enums: `seeders/jobsim_sessions.go:60-69` (`sessStatusEnded`/`sessCompletionPassed`/`sessCompletionFailed`/`sessResultCompleted`/…).
- Reserved-tail disjoint mechanism: `seeders/contentref.go` (`reservedSimRefs`/`reservedAt`/`general()`); AI-readiness type-aware reader precedent: `ai_readiness_config.go:290` (`readAIReadinessSkillPool`).
- Role split (candidate = remainder): `seeders/users.go:506` (`roleForIndex`); g2 admin grant → `users.go:234`.
- Share helper: `seeders/target_roles.go:155` (`memberInShare`); story prefix: `helpers.go:22` (`storyKeyPrefix`).
- DAG registration + `resetTables`: `cmd/stackseed/main.go:43` (resetTables), `:418` (`buildRegistry`).
- Honesty gate: `cmd/stackseed/main_test.go:838` (`TestManifest_CanonicalFileMatchesProjection`) — regenerate `presets/seed-generation-manifest.yaml` after editing `presets/stories.seed.yaml`.

## The 4th story entry
_`narrative: hiring`, `RoleMix` ≈ 0.1 admin / 0.9 candidate, hero-trio placeholder._

## HiringConfigSeeder + the type-aware hiring-sim reader
_The 5 shared HIRING sims; the `type=HIRING AND job_position NOT NULL` pattern query; the disjoint reserved tail._

## Snapshot extension — directus.job_position replay
_Replaying all 443 public rows + pinning the 5 chosen HIRING sims; the digest/capture-column changes._

## The candidate-assessment funnel seeder
_Resolve 5 shared sim refs once; MOST candidates on all 5, SOME assigned-not-taken; the differentiated score
distribution; the M219 skill-ladder + closure wiring._

## reset / closure / isolation wiring
_New hiring rows into `resetTables`; `datadna measure-closure`; the `isolation.Guard` audit._
