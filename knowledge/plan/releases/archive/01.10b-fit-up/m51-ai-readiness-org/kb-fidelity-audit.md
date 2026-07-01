---
title: "KB Fidelity Audit — M51 AI-readiness showcase org"
date: 2026-06-30
scope: milestone:M51
invoked-by: build-mstone-iters (Phase 0b pre-flight, iter-01)
---

## Verdict
**GREEN**

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| AI-readiness data model + scoring + cycle-state read-path | `corpus/services/ai-readiness.md` | `stack-demo/app/internal/workforce/`, `.../data/ent/schema/ai_readiness_*.go`, `.../ent/enum/ai_readiness.go` | PAIRED |
| Coverage protocol (manager vantage iteration loop) | `corpus/ops/demo/coverage-protocol.md` | rext `stack-verify/e2e/` | PAIRED |
| Stories & Heroes 7-table chain (the seeder the 3rd org extends) | `corpus/ops/demo/stories-spec.md` | rext `stack-seeding/seeders/persona*.go`, `blueprint/`, `presets/stories.seed.yaml` | PAIRED |
| Seeding spec + closure gate | `corpus/ops/seeding-spec.md` | rext `stack-seeding/dna/seed_closure.go`, `cmd/datadna/` | PAIRED |

No BLIND-AREA topics. The milestone's load-bearing input (the AI-readiness contract) was authored M48 and is PAIRED with verifiable `app` code in the present platform clone (`stack-demo/app`).

## Fidelity Findings (AI-readiness contract — the load-bearing topic)

All 8 audited claims verified against `stack-demo/app`. Compact:

1. **Org-enablement gate** — ALIGNED. `OrganizationSettingAIReadiness = "ai_readiness"` at `ent/enum/organization_settings.go:47`; `isAIReadinessEnabled` at `internal/workforce/readiness_steps.go::isAIReadinessEnabled`.
2. **3-step enum** — ALIGNED (location nit, fixed). Values `skill_mapping`/`simulation`/`interview` exact; enum defined in `ent/enum/ai_readiness.go` (not the workforce file the doc implied) — doc corrected inline.
3. **Scoring 30/40/30 + composite/knowledge/usage formulas** — ALIGNED. `ai_readiness.go:23-25` weights; `ai_readiness_v2.go` `score=t1+t2+t3`, `knowledge=(s1+s2)/70*100`, `usage=t3*100/30`.
4. **9 ent schemas + key cols** — ALIGNED. All 9 schema files present; cycle `status:active|closed` + partial-unique `IndexWhere("status='active'")` (one active cycle/org); skill `node_id`+`weight`; sim `step_type`+`sim_ref`. (One STALE comment in `ai_readiness_skill.go` referencing legacy `membership_skills.skill_id` — but the *doc* matches the live code, so the schema comment is the artifact, not the doc.)
5. **Cycle-state read-path (CRITICAL — the seed-strategy decision)** — ALIGNED. Active/no-cycle → `buildLiveResponse` → `computeOrgBreakdowns` RECOMPUTES from `user_skill_evidences` (`queryUserAISkills`) + `jobsimulation.sessions` (`queryReadinessSimScores`); `keepStartedMembers` excludes no-step-1-signal members. Closed → `buildResponseFromSnapshots` reads frozen `ai_readiness_snapshots`. `ai_readiness_live_snapshots` confirmed a materialized cache (rewritten by `RefreshLiveSnapshots`), NOT the dashboard source. **The signals-true active-cycle seed strategy is sound.**
6. **Step-completion signals** — ALIGNED. Step1 presence-based on `user_skill_evidences` (level ignored, by design); Steps2/3 best `status='ended' AND score IS NOT NULL` jobsim session joined to `ai_readiness_sims` on `sim_ref=sim_id`.
7. **GraphQL/REST interface** — ALIGNED. `aiReadinessEnabled: Boolean!` (false, not error, for non-enabled); REST `GET /api/workforce/ai-readiness`.
8. **Narrative fallback** — ALIGNED. Persisted in `ai_readiness_diagnose_narratives` keyed by signals sha256; AI-error → empty + FE static per-archetype fallback (demo with no AI key renders).

## Completeness Gaps
None critical. The contract's "Seeding contract (demo / M51)" section already enumerates the exact rows the seeder writes per cycle-state.

## Applied Fixes (inline)
1. `corpus/services/ai-readiness.md` — clarified the `AIReadinessStepType` enum location (`ent/enum/ai_readiness.go`, consumed in the workforce file).
2. `corpus/services/ai-readiness.md` — switched the `isAIReadinessEnabled` numeric line-anchor (`:452`, drifted to `:454`) to a symbol-name anchor (drift-proof).
3. `corpus/services/ai-readiness.md` — bumped the Status freshness marker (re-verified GREEN 2026-06-30).

## Open Items (require user decision)
None.

## Gate Result
**GREEN — proceed.** The bootstrap tok may author its strategy against the verified contract. Net-new seeder machinery for M51 (none of which is a knowledge blind area): (a) an `organization_settings` `ai_readiness` gate writer (nothing writes that table today), (b) a way to pin specific `sim_ref` sim-ids for Step-2/3 funnel sessions (sim_id is hash-selected today, not YAML-configurable), (c) a 3-step onboarding/evaluation funnel seeder (`ai_readiness_*` rows + the step-progress + the signals), and (d) optional cockpit `DeepLinkCatalog` entries for the AI-readiness manager route + employee onboarding-element routes.
