# M51 — spec notes

_Technical notes accumulate here during build (file:line surfaces, rext tag, schema findings)._

## Pre-flight audits — iter-01

**KB-fidelity gate (2026-06-30): GREEN.** Report: `kb-fidelity-audit.md`. All 8 AI-readiness contract claims
ALIGNED vs `stack-demo/app` (incl. the load-bearing cycle-state read-path); 3 doc-hygiene fixes applied inline. No
blind areas, no stale load-bearing claims.

## Topic → doc → code triples (fast-start for future audits)

| Topic | Doc | Code (live surface) |
|---|---|---|
| AI-readiness model/scoring/cycle-state | `corpus/services/ai-readiness.md` | `stack-demo/app/internal/workforce/{ai_readiness,ai_readiness_v2,readiness_steps,readiness_narrative,cycles,live_snapshots}.go` + `.../data/ent/schema/ai_readiness_*.go` + `.../ent/enum/{ai_readiness,organization_settings}.go` |
| Stories & Heroes seeder | `corpus/ops/demo/stories-spec.md` | rext `stack-seeding/blueprint/{blueprint,stories}.go`, `seeders/{persona,persona_write,users,jobsim_sessions,membership_skills,population_evidence,cockpit}.go`, `presets/stories.seed.yaml` |
| Coverage protocol (manager vantage) | `corpus/ops/demo/coverage-protocol.md` | rext `stack-verify/e2e/` (Playwright; manager manifest in `lib/coverage-manifest.ts`) |
| Seed closure gate | `corpus/ops/seeding-spec.md` | rext `stack-seeding/dna/seed_closure.go`, `cmd/datadna/` (`measure-closure`) |

## Reverse-engineered seeder surface (iter-01 survey)

- **3rd-org entrypoint:** append a `stories[]` entry to `stack-seeding/presets/stories.seed.yaml` (the preset
  `/demo-up` seeds + exports roster + projects cockpit from). The first story keeps the Clerkenstein default org-id;
  every story after the first gets a deterministic distinct org-id from `StoryOrgID(story.ID)` (no code change for
  org identity). Schema: `blueprint.Story{ID,Name,Annotation,Org,Heroes,Batches}`, `Persona{...,Vantage,Trajectory,
  JumpTo,Demonstrates,Skills}`.
- **Reused as-is:** PersonaSeeder 7-table verified chain (`persona.go`/`persona_write.go`), JobsimSessionsSeeder
  (`status='ended'`, `result_status='completed'`, `growthArcScore()`; **no FK on sim_id** so a chosen sim_ref
  inserts cleanly), the `user_skill_evidences` UPSERT (Step-1 signal model), the org-agnostic closure gate.
- **Net-new for M51:** (a) an `organization_settings` `ai_readiness` gate-row writer (zero seeders write that table
  today); (b) a sim-id pin mechanism for Step-2/3 funnel sessions (sim_id is hash-selected from ~5 replayed
  templates today — no YAML hook); (c) a 3-step onboarding/evaluation funnel seeder writing the `ai_readiness_*`
  config (cycle/skills/sims/steps) + `ai_readiness_user_step_progress` + the per-member signals at ~80%-all-3;
  (d) cockpit `DeepLinkCatalog` entries for the manager AI-readiness route + employee onboarding-element routes
  (raw jump_to already flows through, but a catalog entry adds a proper label).
- **Closure-gate constraint:** AI-readiness skills must resolve via `resolveTaxonomyRefs` (real replayed taxonomy);
  AI-named-skill biasing already exists in `membership_skills.go` (`isAISkillName`/`filterAISkills`, mirrors the
  dashboard's `matchAISkill` substring semantics) — reuse for the AI-narrative org bias.

## Live target
demo-1 UP (17 containers, offset 10000: backend :18082=200, next-web :13000, postgres :15432, cockpit :17700).
Consumes rext `fit-up-m49` (the `fit-up-m50` consumption-re-pin is the KEEP deferred to M53). Iterate in place
(re-seed / re-sweep against demo-1; single-machine-one-demo serialization).
