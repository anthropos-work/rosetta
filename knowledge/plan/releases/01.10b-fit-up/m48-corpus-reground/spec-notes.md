# M48 — spec notes

_Technical notes accumulate here during build (file:line surfaces, schema findings, the drift gap-list)._

## Pre-flight audits — S1/S2 (corpus re-ground)

**Phase 0b KB-fidelity verdict: GREEN-by-design** (recorded 2026-06-29).

M48 is the inverse of the usual gate: its *deliverable* is re-grounded docs, and the contract it reads is the
**current platform code** (the M47-confirmed-current clones), not knowledge docs. "Stale corpus docs" is not a
blind area that blocks M48 — it is M48's job. The one genuinely-undocumented surface (member-AI-readiness) is the
milestone's load-bearing **new** deliverable (S2), authored from the code. No blocker; proceed.

## Investigation (3 parallel read-only agents, launched 2026-06-29)
- **AI-readiness backend data model** (app/jobsimulation/skiller/skillpath/cms) — the seeding contract for M51.
- **AI-readiness frontend + manager dashboard** (next-web `ai-readiness/`) — the surface map.
- **Corpus drift survey** (corpus/architecture + corpus/services vs current clones) — the S1 material-lag gap list
  + the AI-readiness doc placement + the ant-academy stale-claim confirmation.

### S1 — drift gap-list (from the drift-survey agent, 2026-06-29)

**Corpus inventory:** 10 `architecture/*.md` + 22 `services/*.md` + CLAUDE.md.

**CRITICAL:**
- **member-AI-readiness — ZERO corpus coverage** (the headline gap → S2). Backend: `app/internal/workforce/`
  (`ai_readiness.go`, `ai_readiness_v2.go`, `readiness_narrative.go`, `readiness_steps.go`, …) + GraphQL
  `app/internal/web/backend/graphql/graph/schemas/ai_readiness.graphqls` + ~10 REST/RPC handlers in
  `internal/web/backend/api/api.go` (`workforce-ai-readiness`, `ai-readiness-narrative`, `ai-readiness-cycles-*`,
  `ai-readiness-csv`, `ai-readiness-compare`, `workforce-sv-role-readiness`) + a background
  `internal/worker/tasks/ai_readiness_refresh.go`. Frontend: `next-web-app/packages/ui/src/Workforce/AIReadinessView/`
  (+ `StepsCompletionDrawer`), org setting `AIReadinessSetting.tsx`, navbar gated on `showAIReadiness`.
  **3-step scoring: skill-mapping 30 → simulation 40 → interview-extraction 30 (= 100).** Org-level `ai_readiness`
  setting gates the feature. → **NEW doc `corpus/services/ai-readiness.md`** (S2).
- **ant-academy `repos.yml` claim is FALSE** (→ S4 doc-fix; M49 #5 owns the repos.yml fix). `CLAUDE.md:199` +
  `corpus/services/ant-academy.md:26` claim "cloned via repos.yml / In repos.yml: Yes", but
  `stack-demo/platform/repos.yml` lists 13 repos with **no ant-academy**.

**MINOR (material-lag, S3):**
- `corpus/services/backend.md` (~L84-92 "Recent Feature Additions") — omits AI-readiness.
- `corpus/services/next-web-app.md` — omits AIReadinessView / AIReadinessSetting / navbar item.
- `corpus/architecture/architecture_overview.md` (service inventory + mermaid) — Backend row omits AI-readiness.
- `corpus/architecture/service_taxonomy.md` — recent-additions note omits AI-readiness.

### S2 — AI-readiness FRONTEND surface map (from the frontend agent, 2026-06-29)

- **Route:** `/app/(authenticated)/(verified)/ai-readiness/` (`AIReadinessClient.tsx`). **Access gate chain:** PostHog
  `flag_ai_readiness` **AND** GraphQL `aiReadinessEnabled` (org setting) → else bounce to `/home`. Admin alt-entry
  under `enterprise/workforce/.../ai-readiness/`.
- **Manager dashboard:** HeroCard (org **score** 0-100 + dominant **archetype** + **Steps-Completion %** = stage3/members)
  + 4 tabs: **Snapshot** (archetype 2×2 Knowledge×Usage matrix, donuts, by-tag table), **How-we-measure** (the 3-step
  ribbon + skill strengths/gaps + sims + interview findings), **What-to-do-next** (archetype action groups +
  per-person Diagnose narrative), **Compare** (cycle deltas, gated off by default). Cycle selector (`CyclePill`).
- **3-step funnel** (`components/ai-readiness/aiReadiness.constants.ts` STEPS): **1 Skill-mapping 30pts** (framework
  modal) · **2 Hands-on 40pts** (a simulation, `simSlug`) · **3 Usage-discovery 30pts** (an interview, `interviewSlug`)
  = 100. Member-facing `AIReadinessHero.tsx` (modes new/progress/done/archived).
- **STARTED vs COMPLETED:** person `stage` ∈ {1,2,3} (0 = done/none), `score` (null = not-completed, number = completed);
  org aggregate `stage1/stage2/stage3` = counts reaching each stage → the funnel + the "X% completed all 3 steps".
- **Data contract:** REST `/api/workforce/ai-readiness` (+ `/cycles`, `/steps-completion`, `/narrative`, `/compare`,
  `/export.csv`) → `AIReadinessResponse {org, byTeam[], people[], howWeMeasure, cycle}`. GraphQL (member program):
  `aiReadinessEnabled`, `aiReadinessUserPlanProgress`, `aiReadinessSkills`, `completeAiReadinessSkillMapping`.
- **Empty when:** flag off / org setting off / 0 people / step data null. **Populated when:** `org.members>0` + people
  records carry stage+score + howWeMeasure skill/sim/interview non-empty.
- **Seed-from-the-frontend view (M51):** enable flag + org `ai_readiness` setting; create a cycle; populate people
  with stage+score (≈80% at stage3/completed) + the howWeMeasure skill/sim/interview aggregates. (Exact tables ⇐ the
  BACKEND agent — pending.)

### S2 — AI-readiness BACKEND data model (from the backend agent, 2026-06-29) — the M51 seeding contract

All tables in **`app`** (`public` schema), ent under `app/internal/data/ent/schema/ai_readiness_*.go` (9 schemas
verified present). Enablement: `organization_settings` row `setting='ai_readiness', is_enabled=true`
(enum `OrganizationSettingAIReadiness` @ `enum/organization_settings.go:47`; checked `readiness_steps.go:452`).

**9 tables:** `ai_readiness_cycles` (one active/org), `ai_readiness_steps` (org 3-step plan), `ai_readiness_skills`
(node_id+weight 1.0/0.5), `ai_readiness_sims` (sim_ref, step_type sim/interview), `ai_readiness_user_step_progress`
(status not_started/in_progress/completed + completed_at), `ai_readiness_live_snapshots` (mutable per-member
score/stage/archetype), `ai_readiness_snapshots` (frozen per-cycle), `ai_readiness_diagnose_narratives` (persisted
LLM, sha256-signals-keyed), `ai_readiness_text_translations`.

**Scoring** (`ai_readiness.go:173-229`): step1 skill-mapping ≤30 (matched-weight/total ×30) · step2 sim ≤40
(raw/100×40) · step3 interview ≤30. score=Σ (≤100); knowledge=(s1+s2)/70×100; usage=s3-scaled. Archetypes
(threshold 50): champion/hidden_talent/explorer/standby. stage∈{1,2,3} (0=done/none), score null=not-done.

**M51 seed contract** (→ now documented in `corpus/services/ai-readiness.md` §"Seeding contract"): org config
(~10 rows: setting + steps + ~5 skills [real node_ids] + 2 sims + 1 active cycle) + per-member (≈160 "completed":
signals OR direct progress+live-snapshot; 1 "started" stage=1/score≈30; 1 "completed" stage=3/score≈100). **No AI
keys needed** (narratives fall back to static). Two strategies (signals-true vs snapshot-direct) — **an M51 decision**.

### S2 done — `corpus/services/ai-readiness.md` authored (the load-bearing deliverable)
New doc covers: role, the org-enablement gate (setting + PostHog flag), the 3-step framework + scoring, the 9-table
data model code-map, the GraphQL+REST interface, the manager-dashboard + member surfaces, narrative generation, and
the **M51 seeding contract**. Cross-linked from CLAUDE.md, backend.md, architecture_overview.md, next-web-app.md,
service_taxonomy.md. All cross-refs verified resolving; fences balanced.
