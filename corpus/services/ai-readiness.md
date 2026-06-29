# AI Readiness (Workforce) — service documentation

> **Status:** documented 2026-06-29 (v1.10b "fit-up" M48 — corpus re-ground). The feature ships in **`app` v1.315+**
> (backend) + **`next-web-app` v2.89.0+** (UI) and had **no prior corpus coverage** (it was invisible to the
> ~1-month-stale clones, which is why M201's verify reported it as a false-negative). This doc is the contract the
> v1.10b **M51** AI-readiness showcase-org seeder builds against.

## Role & Responsibility

**AI Readiness** is an org-level **AI-capability diagnostic**: each member runs a **3-step onboarding/evaluation**,
which produces a per-member readiness **score (0–100)** + an **archetype**, rolled up into a **manager dashboard**
(a funnel + a Knowledge×Usage archetype matrix + per-team/per-person drill-down). It is a subsystem of the **`app`
(backend) service** — `app/internal/workforce/` — not a standalone microservice. It is **org-gated** (off by default)
and **member-facing** for the onboarding flow + **manager/analytics-facing** for the dashboard.

## Org enablement (the gate)

The feature is off until an org turns it on. Two gates compose (both must be true for the UI to render):

1. **Org setting** — a row in `organization_settings` with `setting = 'ai_readiness'`, `is_enabled = true`
   (`app/internal/data/ent/enum/organization_settings.go:47` → `OrganizationSettingAIReadiness = "ai_readiness"`;
   checked by `WorkforceManager.isAIReadinessEnabled`, `app/internal/workforce/readiness_steps.go:452`). No row =
   off. Exposed to the FE as the GraphQL query `aiReadinessEnabled: Boolean!`
   (`resolver_ai_readiness.go` — returns `false`, not an error, for non-enabled orgs).
2. **PostHog flag** `flag_ai_readiness` — the next-web client also gates the route on this flag before it even
   queries `aiReadinessEnabled` (`apps/web/.../ai-readiness/AIReadinessClient.tsx`).

## The 3-step framework + scoring

The evaluation is a fixed **3-step** framework (per-org orderable, canonical default below), each step a scoring
axis (`enum.AIReadinessStepType`; `app/internal/workforce/ai_readiness.go:173-229`):

| # | Step (`step_type`) | Method | Max pts | Signal that completes it |
|---|--------------------|--------|---------|--------------------------|
| 1 | `skill_mapping`    | self-map AI skills (framework modal) | **30** | ≥1 `user_skill_evidences` row for a skill node_id in the org's `ai_readiness_skills` set (presence-based; level ignored) |
| 2 | `simulation`       | a job-simulation session | **40** | best ended/scored jobsim session whose sim_id ∈ `ai_readiness_sims` (`step_type='simulation'`); `(raw_score/100)×40` |
| 3 | `interview`        | an interview (a sim, different config) | **30** | best ended/scored session whose sim_id ∈ `ai_readiness_sims` (`step_type='interview'`); `(raw_score/100)×30` |

**Composite scores (0–100):** `score = step1+step2+step3` (max 100) · `knowledge = (step1+step2)/70×100` (axis X) ·
`usage = step3 scaled` (axis Y). **Archetype** (2×2, threshold 50 per axis): **Champion** (hi knowledge, hi usage) ·
**Hidden Talent** (hi/lo) · **Explorer** (lo/hi) · **Standby** (lo/lo). Buckets/bands: none(0-10)/low(11-40)/
medium(41-70)/high(71-100).

**Started vs completed (the funnel):** a member carries a `stage` ∈ {1,2,3} (0 = none/done) and a `score` (null =
not-completed). Per-step status lives in `ai_readiness_user_step_progress` (`not_started`/`in_progress`/`completed`,
`completed_at` set once, never re-updated). The org funnel = counts of members reaching `stage1`/`stage2`/`stage3`;
the dashboard's "X% completed all 3 steps" = `stage3 / members`.

## Data model (code map)

All tables live in `app` (`public` schema); ent schemas under `app/internal/data/ent/schema/`:

| Table | Purpose | Key cols / states |
|-------|---------|-------------------|
| `ai_readiness_cycles` | a diagnostic cycle per org | `status: active\|closed`, `start/end_date`, `final_score`, `closed_at`; **one active cycle per org** (partial unique index) |
| `ai_readiness_steps` | org's ordered 3-step plan | `step_type`, `position`; default = all 3 canonical if no rows |
| `ai_readiness_skills` | org's AI-skill set | `node_id` (taxonomy), `weight` (1.0 core / 0.5 enabling) — Step 1 scoring |
| `ai_readiness_sims` | org's sim registry | `step_type` (simulation/interview), `sim_ref` (Directus sim id or `PLACEHOLDER-{slug}`) — Steps 2/3 |
| `ai_readiness_user_step_progress` | per-(org,user,step) progress | `status`, `completed_at` |
| `ai_readiness_live_snapshots` | **live** per-member score (mutable, upserted) | `score/knowledge/usage/archetype/stage/...`; exposed to Talk-to-Data SQL |
| `ai_readiness_snapshots` | **frozen** per-(cycle,user) snapshot at close (immutable) | `frozen_*` mirror of live |
| `ai_readiness_diagnose_narratives` | persisted per-member LLM narrative | keyed `(org,user,cycle_ref,lang)` + `signals_hash` |
| `ai_readiness_text_translations` | content-addressed translation cache | `source_hash`+`lang` |
| `organization_settings` (existing) | the enablement gate | `setting='ai_readiness'`, `is_enabled` |

Scoring engine: `app/internal/workforce/ai_readiness.go` (`computeAIReadiness`, `GetAIReadinessWithOptions`,
`computeOrgBreakdowns`). Steps/progress: `readiness_steps.go`. Cycles: `cycles.go`. Narrative: `readiness_narrative.go`.

## Interface

- **GraphQL** (`app/internal/web/backend/graphql/graph/schemas/ai_readiness.graphqls`; resolver
  `resolver_ai_readiness.go`): `aiReadinessEnabled`, `aiReadinessUserPlanProgress` (member step status + deadline),
  `aiReadinessSkills` (skills to map), mutation `completeAiReadinessSkillMapping`.
- **REST/workforce API** (`app/internal/web/backend/api/api.go`): `GET /api/workforce/ai-readiness` (→ the
  `AIReadinessResponse` aggregate the manager dashboard consumes), `/cycles` (GET/POST), `/steps-completion`,
  `/narrative` (POST, LLM diagnosis), `/compare`, `/export.csv`.
- **Background:** `app/internal/worker/tasks/ai_readiness_refresh.go` re-materializes live snapshots.

## Surfaces (UI)

- **Manager dashboard** (`next-web-app` `apps/web/.../ai-readiness/`): HeroCard (org score + dominant archetype +
  **Steps-Completion %**) + tabs **Snapshot** (archetype matrix + donuts + by-tag), **How-we-measure** (the 3-step
  ribbon + skill strengths/gaps + sims + interview findings), **What-to-do-next** (archetype action groups + per-person
  **Diagnose** narrative drawer), **Compare** (cycle deltas — gated off by default).
- **Member onboarding** (`apps/web/src/components/ai-readiness/AIReadinessHero.tsx`): the 3-step funnel
  (modes new/progress/done/archived); Step 1 = skill-mapping modal, Step 2 → a sim, Step 3 → an interview.

## Narrative generation

Per-member manager-facing narratives are **persisted, not regenerated per read** (`readiness_narrative.go:60-98`):
a sha256 of the member's signals keys a read-through cache in `ai_readiness_diagnose_narratives`; on a miss/stale
hash it calls the LLM (GPT-5-Mini) and upserts. On AI error it returns empty + the FE falls back to static
per-archetype guidance — **so a demo with no AI key still renders** (narratives just show the static fallback).

## Local development

Enable for an org by inserting the `organization_settings` row (`setting='ai_readiness', is_enabled=true`) +
`flag_ai_readiness` on in PostHog (or the local flag shim). The member flow then needs `ai_readiness_skills` +
`ai_readiness_sims` config + an active cycle; the dashboard reads `GET /api/workforce/ai-readiness`. Tests:
`app/internal/workforce/*_test.go` (the scoring/steps/cycle suites).

## Seeding contract (demo / M51)

To make a **200-person demo org** show the AI-readiness manager dashboard **enabled**, with **~80% (≈160 members)
having completed all 3 steps**, plus **one hero "started"** and **one hero "completed"** — the seeder writes:

**Org config (≈10 rows):**
1. `organization_settings` (`setting='ai_readiness'`, `is_enabled=true`) — the gate.
2. `ai_readiness_steps` × 3 (skill_mapping/simulation/interview, positions 0/1/2) — optional (canonical default if absent).
3. `ai_readiness_skills` × ~5 core (weight 1.0) + a few enabling (0.5), `node_id` = **real taxonomy node-ids** (route
   through the existing seeding resolvers — never fabricate, per the closure gate).
4. `ai_readiness_sims` × 2 (`step_type` simulation + interview, `sim_ref` = a real Directus sim id or a `PLACEHOLDER-` ref).
5. `ai_readiness_cycles` × 1 (`status='active'`).

**Per-member (≈160 "completed"):** the underlying signals (≥1 `user_skill_evidences` for a configured skill;
jobsim sessions for steps 2/3) **+** `ai_readiness_user_step_progress` (3× `completed`) **+** an
`ai_readiness_live_snapshots` upsert (`score≈100, stage=3, archetype` per the score). The **"started" hero**: only
the skill_mapping signal + a `stage=1`/`score≈30` live snapshot. The **"completed" hero**: all 3 + `stage=3`.

**⚠️ Which table the dashboard reads depends on the cycle state — this dictates the seed strategy (an M51
decision):**

- **Active cycle → the dashboard RECOMPUTES from signals.** `GetAIReadinessWithOptions` → `buildLiveResponse` →
  `computeOrgBreakdowns` (`ai_readiness.go:283-343`) re-derives each member's score **from the underlying signals**:
  `user_skill_evidences` (step 1) + the readiness jobsim sessions (steps 2/3) + the `ai_readiness_skills`/
  `ai_readiness_sims` config — and `keepStartedMembers` **excludes members with no step-1 signal** from the
  aggregate. So an **active**-cycle dashboard requires the **signals-true** seed (write the real skill evidences +
  sim sessions + `ai_readiness_user_step_progress`; reuse the existing verified-skill chain). `ai_readiness_live_snapshots`
  is a **materialized cache** (rewritten by `RefreshLiveSnapshots`, consumed by Talk-to-Data SQL) — **NOT** the
  dashboard's source: seeding it directly does **not** make the live dashboard render and is overwritten on refresh.
- **Closed cycle → the dashboard reads frozen snapshots.** `buildResponseFromSnapshots` reads `ai_readiness_snapshots`
  directly, so a **closed**-cycle showcase can be seeded **snapshot-direct** (write the `frozen_*` rows + flip the
  cycle to `closed`) with **no underlying signals** — lighter, but the world reads as a *finished* assessment.

**No AI keys needed either way** (diagnosis narratives fall back to static per-archetype text on AI error).

## Cross-references

- **Authoritative in-repo deep-dive** (the platform's own KB): `app/knowledge/ai-readiness/overview.md` (start
  there; per-topic docs under `app/knowledge/ai-readiness/`) — the 2-axis/4-archetype model, the 3-step plan, the
  scoring engine, live-vs-frozen cycles, compare, CSV, Talk-to-Data. This corpus doc *summarizes* for the rosetta
  reader + the M51 seeder; that KB is the source of truth for deep work.
- Backend service: [`backend.md`](backend.md) (AI Readiness is an `app/internal/workforce/` subsystem).
- The seeded demo world it plugs into: [`../ops/demo/stories-spec.md`](../ops/demo/stories-spec.md) (the Stories &
  Heroes model — M51 adds the AI-readiness showcase org as a 3rd story).
- Verified-skill chain the Step-1 signal reuses: [`../ops/demo/stories-spec.md`](../ops/demo/stories-spec.md)
  (`user_skill_evidences`).
- AI provider routing for the narrative LLM: [`../architecture/ai_architecture.md`](../architecture/ai_architecture.md).
