# M225 — Spec notes

_Technical detail accumulated during the milestone: file:line maps, command transcripts, measured numbers._

## Pre-flight audits — S1/S2/S3 (2026-07-16)
KB-fidelity audit: `knowledge/plan/releases/02.40-casting-call/m225-dress-the-set/kb-fidelity-audit.md` — **YELLOW**
(corpus fully aligned; one stale PLAN-doc premise reconciled inline, KB-1). Topic → doc → code triples (for reuse):

| Topic | Knowledge doc | Code path | Status |
|---|---|---|---|
| Auto-set-dress HIRING-sim replay | `corpus/ops/snapshot-spec.md` · `corpus/ops/demo/recipe-snapshot-world.md` | rext `demo-stack/up-injected.sh` (auto-set-dress pass) · `stack-snapshot/` · `stack-seeding` `readHiringSimPool`/`HiringConfigSeeder` | PAIRED |
| Hiring coverage manifest | `corpus/ops/demo/coverage-protocol.md` | rext `stack-verify/e2e/lib/coverage-manifest.ts` (`manifestFor`) · `stack-verify/e2e/tests/coverage.spec.ts` · `persona-check.spec.ts` | PAIRED (hiring section = DELIVERS) |
| Hiring playthrough + pt-world | `corpus/ops/demo/playthroughs.md` | rext `playthroughs/manifest/*.yaml` + `manifest.go`/`seed_worlds.go` · `playthroughs/seed/pt-world.seed.yaml` · `playthroughs/e2e/` | PAIRED (hiring.yaml + org = DELIVERS) |
| Hiring domain (read-path / is_hiring) | `corpus/services/hiring.md` (M222) | app `intelligence.go` `InsightsJobSimulationByMemberships` · `local_jobsimulation_sessions` mirror | PAIRED |

## Auto-set-dress: HIRING-sim capture + replay (S1) — NO job_position replay
_Folding the HIRING-sim (`SIMULATION_TYPE_HIRING`) capture + replay into the default `/demo-up` bring-up pass so the
hiring org's 5 positions + candidate funnel come up real with no manual steps._

**KB-1 reconciliation (M222 BA-6 / M223 D4):** the scaffolded S1 line said "`directus.job_position` replay". That
premise was **refuted before M225** — `directus.job_position` captured **0 rows** (the prod "443" was never
captured) and the recruiter scoreboard does **not** read `job_position` (`JobSimulation.jobPosition` is
optional/unused). The 5 "positions" ARE 5 real captured `SIMULATION_TYPE_HIRING` sims, resolved by
`readHiringSimPool` (`stack-seeding`, `directus.simulations WHERE type='SIMULATION_TYPE_HIRING'`, reserved-disjoint)
and materialized as `organization_sim_invitation_links` by the `HiringConfigSeeder`. So S1 = **ensure the HIRING-sim
pool is present in the default auto-set-dress directus replay** (so `readHiringSimPool` resolves ≥5 with no manual
capture step) + the hiring org/funnel seeds by default (already in `presets/stories.seed.yaml`, M223) + the two-app
hiring container comes up by default (M224). No `snapshot-spec.md` `job_position` surface is built.

## Hiring coverage manifest (S2)
_Wired into `manifestFor(vantage, expectedOrg, identityKey)` — the 3-arg org/identity-conditional dispatch already
in `coverage-manifest.ts` (the AI-readiness `Northwind Aviation` showcase-org precedent, M53 AB4). Add a HIRING
manager/candidate manifest set + a `MERIDIAN_TALENT` org gate; assert candidate persona self-consistency
(role↔skills↔score), the compare-surface sections, 0 prod-eject._

## Hiring playthrough + pt-world (S3)
_`playthroughs/manifest/hiring.yaml` (recruiter compares candidates on a shared sim); a DISTINCT hiring org into the
decoupled `pt-world` seed (test data ≠ demo data — NOT "Meridian Talent"/"Meridian Labs"); the GREEN run record._
