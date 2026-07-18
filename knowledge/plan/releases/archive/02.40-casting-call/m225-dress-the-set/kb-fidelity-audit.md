---
title: "KB Fidelity Audit — M225 dress-the-set"
date: 2026-07-16
scope: milestone:M225
invoked-by: build-milestone
---

## Verdict
**YELLOW** — proceed with tracking. No blind areas; no stale load-bearing **corpus** claims (the corpus is fully
aligned, including the `job_position` drop). One stale premise in the M225 **plan docs** (S1 `job_position` replay),
reconciled inline (KB-1). The docs M225 DELIVERS (hiring sections of `coverage-protocol.md` + `playthroughs.md`) do
not exist yet — that is expected (they are the milestone's deliverable), not a blind area.

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Auto-set-dress HIRING-sim replay (S1) | `corpus/ops/snapshot-spec.md` · `corpus/ops/demo/recipe-snapshot-world.md` | rext `demo-stack/up-injected.sh` (auto-set-dress pass, `build_frontend_hiring`) · `stack-snapshot/` · `stack-seeding` `readHiringSimPool`/`HiringConfigSeeder`/`HiringFunnelSeeder` · `presets/stories.seed.yaml` (hiring story) | PAIRED |
| Hiring coverage manifest (S2) | `corpus/ops/demo/coverage-protocol.md` | rext `stack-verify/e2e/lib/coverage-manifest.ts` (`manifestFor(vantage, expectedOrg?, identityKey?)`) · `coverage.spec.ts` · `persona-check.spec.ts` | PAIRED (hiring section = DELIVERS/DOC-ONLY) |
| Hiring playthrough + pt-world (S3) | `corpus/ops/demo/playthroughs.md` | rext `playthroughs/manifest/*.yaml` + `manifest.go`/`seed_worlds.go`/`validator.go` · `playthroughs/seed/pt-world.seed.yaml` (3 orgs A/B/C) · `playthroughs/e2e/` | PAIRED (hiring.yaml + org = DELIVERS/DOC-ONLY) |
| Hiring domain / read-path / `is_hiring` | `corpus/services/hiring.md` (M222) | app `intelligence.go` `InsightsJobSimulationByMemberships` · `public.local_jobsimulation_sessions` mirror | PAIRED |
| Demo UI tier (two-app hiring container) | `corpus/ops/demo/frontend-tier.md` · `corpus/ops/demo/demopatch-spec.md` | rext `demo-stack/up-injected.sh` `build_frontend_hiring` + 4 hiring patches | PAIRED |

## Fidelity Findings

1. **`manifestFor` signature — ALIGNED.** `coverage-protocol.md` documents `manifestFor(vantage, expectedOrg,
   identityKey)` (3-arg, org/identity-conditional). Code: `coverage-manifest.ts:851`
   `manifestFor(vantage, expectedOrg?, identityKey?)` — matches, with the AI-readiness `Northwind Aviation`
   showcase-org gate as the exact precedent M225 S2 reuses (`AI_READINESS_SHOWCASE_ORG` + `isShowcaseOrg` substring
   match). **Fix owner:** none.

2. **`job_position` drop — corpus ALIGNED, plan-doc STALE (KB-1).** Corpus (`snapshot-spec.md:386`,
   `seeding-spec.md:352`, `stories-spec.md:633`, `hiring.md:75`) correctly states there is NO `directus.job_position`
   replay (0 rows captured; scoreboard doesn't read it; M222 D4 / M223 D4). The M225 **plan docs** (`overview.md`
   Scope.In #1, `progress.md` S1, `spec-notes.md`) still said "fold the `job_position` replay". **STALE**, and
   load-bearing for S1. **Fix owner:** update the plan docs (the code/corpus win). **Applied inline** — see Applied
   Fixes.

3. **pt-world model — ALIGNED.** `playthroughs.md:285-323` documents pt-world's 3 orgs (A Meridian Labs / B Halcyon
   Retail / C Vertex Logistics `ai-readiness`), the M202-D4 anchor-story workaround, and the reset-to-seed
   lifecycle. Code `pt-world.seed.yaml` matches exactly. M225 S3 adds a 4th (hiring) org here. **Fix owner:** none.

4. **Playthrough count — ALIGNED.** `playthroughs.md:102` states "14 live Playthroughs, 1 TODO". S3 adds a hiring
   Playthrough (one GREEN). **Fix owner:** none.

## Completeness Gaps
None critical. The hiring vantage sections of `coverage-protocol.md` and `playthroughs.md` are the milestone's
`Delivers →` targets (authored in S4/Phase 5), not pre-existing blind areas.

## Applied Fixes
- **KB-1:** reconciled the stale S1 `job_position`-replay premise in the M225 plan docs:
  - `overview.md` §Why-section (dropped "replay `job_position`" → "replay the HIRING-sim content") + §Scope.In #1
    (reframed to "fold the HIRING-sim (`SIMULATION_TYPE_HIRING`) capture + replay"; added the NB citing M222 BA-6 /
    M223 D4 + `readHiringSimPool`).
  - `progress.md` S1 line — same reframe + the no-`job_position` NB.
  - `spec-notes.md` — reframed the §Auto-set-dress section, added the topic→doc→code triples + the KB-1 note.
  - `decisions.md` — recorded KB-1.

## Open Items (require user decision)
None.

## Gate Result
**YELLOW: proceed with tracking.** KB-1 tracked in `decisions.md` + applied inline. build-milestone may enter
Phase 1.
