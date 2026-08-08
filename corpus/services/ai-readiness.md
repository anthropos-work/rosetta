# AI Readiness (Workforce) — service documentation

> **Status:** documented 2026-06-29 (v1.10b "fit-up" M48); **re-verified GREEN 2026-06-30** (M51). **Package-refactor
> refresh 2026-07-23 (v2.7 "july jitter" M247)** — see the ⚠️ callout below: the domain moved out of
> `app/internal/workforce/` into its own **`app/internal/aireadiness/`** package (app `v1.351.1`), the archetype
> **scoring bands changed**, and several net-new subsystems landed (notifications, email overrides, one-click
> provisioning, auto-close, a recommendation engine). The feature ships in **`app` v1.315+** (backend) +
> **`next-web-app` v2.89.0+** (UI). This doc is the contract the v1.10b **M51** AI-readiness showcase-org seeder
> builds against; the **demo-seeder fidelity deltas** (the 31-skill default set, track-keyed named sims,
> evaluated-skills set-dress) are owned by **v2.7 M250 (AI-readiness fidelity)** — this refresh covers the
> platform-side package refactor only.

> ### ⚠️ Package refactor — `internal/workforce/` → `internal/aireadiness/` (app v1.351.1, M247 refresh)
>
> The whole AI-Readiness domain moved out of `app/internal/workforce/` into a **new package
> `app/internal/aireadiness/`** (package name `aireadiness`) — commit `4c28365f` ("Refactor AI Readiness domain:
> migrate workforce dependencies to aireadiness package", 2026-07-22). `workforce` keeps the org-analytics KPIs;
> `aireadiness` owns everything readiness-scoped. **The remaining dependency on `workforce` is the member
> directory PLUS the org's skill-scale setting** — the `WorkforceDirectory` interface, which declares
> **FOUR** methods at `app/internal/aireadiness/manager.go:40-51` @ `app` `ad9f3c498`:
> `LoadMembers` (`:43`), `LoadMembersByUserIDs` (`:45`), `BaseMembers` (`:48`) and **`LevelsCount`**
> (`:50`). ⚠️ **Corrected M257x iter-141 — this said *"the ONLY remaining dependency … is the member
> directory"* and named two methods.** The source's own type doc says otherwise in as many words
> (`manager.go:36-39`): *"the slice of the workforce manager this domain needs: the active-member
> directory … **and the org's skill-scale setting**."* `LevelsCount` is an **org setting**, not a member
> call — `readiness.go:770` reads `maxLevel := int(m.workforce.LevelsCount(ctx, orgID))`.
> **And *"whose implementations stayed in `members.go`"* is wrong for the fourth**: `LoadMembers` /
> `LoadMembers` / `LoadMembersByUserIDs` / `BaseMembers` are at `app/internal/workforce/members.go:349`/`:353`/`:357`,
> but **`LevelsCount` is at `app/internal/workforce/manager.go:90`** (`git grep "func .*LevelsCount"`
> returns three sites tree-wide: the unexported `getLevelsCount` at `manager.go:61`, the exported one at
> `:90`, and a test fake). Independently re-derived at source, iter-141.
>
> **File renames** (older `app/internal/workforce/…` anchors elsewhere in this doc refer to the pre-refactor
> location — resolve them under `internal/aireadiness/`):
>
> | Old (`internal/workforce/`) | New (`internal/aireadiness/`) |
> |---|---|
> | `ai_readiness.go` | **`readiness.go`** (the scoring engine + read entrypoints) |
> | `ai_readiness_v2.go` | `scoring.go` (archetype/axis math + bands) |
> | `ai_readiness_csv.go` | `csv.go` |
> | `readiness_steps.go` | `steps.go` |
> | `readiness_narrative.go` | `narrative.go` |
> | `how_we_measure_v2.go` | folded into `how_we_measure.go` (`computeInterviewInsightsV2`) |
> | `cycles.go`/`compare.go`/`diagnosis.go`/`provision.go`/`defaults.go`/… | same names under `internal/aireadiness/` |
> | `emailoverride/`, `emailpreview/`, `notifications/` | same, under `internal/aireadiness/` |
> | `*_test.go` (scoring/steps/cycle suites) | moved with the package; harness `testdb_test.go` (pgtest) |
>
> **D-07 demopatch re-anchor (load-bearing).** The `app-aireadiness-snapshot-loadmembers` demo-patch anchored on
> `app/internal/workforce/ai_readiness.go` at the `buildResponseFromSnapshots → loadMembers(orgID,"")` call. That
> file no longer exists — the call is now at **`app/internal/aireadiness/readiness.go`**, `buildResponseFromSnapshots`,
> as **`m.workforce.LoadMembers(ctx, orgID, "")`** *through the `WorkforceDirectory` interface* (the bounded swap
> `LoadMembers → LoadMembersByUserIDs` is now expressible at that interface call site, since `WorkforceDirectory`
> already exposes `LoadMembersByUserIDs`). **The re-anchor has LANDED** — v2.7 **M254**, not the M250 the M246
> drift-ledger's **D-07** item originally assigned it to: the manifest reads `path:
> internal/aireadiness/readiness.go` (`app-aireadiness-snapshot-loadmembers.yaml:42`) and its own header says
> *"v2.7 M254 RE-POINT"* (`:33`), with `demo-stack/tests/test_aireadiness_snapshot_loadmembers_m254.py`
> fencing it. This paragraph stated completed work as outstanding while the **⚠⚠ M51 iter-08/09** block below
> already described the same call site in its post-refactor location — *"(now `aireadiness/readiness.go`,
> formerly `workforce/ai_readiness.go:512`)"* — i.e. the doc contradicted itself; corrected M257x iter-46.
> ⚠️ **The bare line number was itself the defect.** That cross-reference carried one line number when iter-46 wrote it and
> the next one down after iter-100 mechanically re-pointed it, and **neither line has ever carried the statement** — the later of the two
> merely opens the `> **✅ CORRECTED M219 …**` blockquote, which is about an unrelated claim. The target is
> therefore named by construct above — the **`⚠⚠ M51 iter-08/09`** block, identified by the quoted
> parenthetical *"(now `aireadiness/readiness.go`, formerly `workforce/ai_readiness.go:512`)"*.
> ⚠️ **And the third generation rotted too, which is why this passage no longer carries a line pin at all
> (M257x iter-115).** iter-102 replaced `:459` with **`:496`**; re-derived at iter-115, `:496` was the
> **closing line of the `> **✅ CORRECTED M219 …**` blockquote** — the very blockquote the two sentences above
> name as the *wrong* target for the previous generation. **The pin was deleted rather than re-derived**:
> three generations of one anchor in one file, each off by a handful of lines, is not three accidents — a
> same-file line pin in a document that grows will rot again, and the self-heal clause (*"when the two
> disagree, the name wins"*) lowers a reader's cost without making the number true. **Search the construct.**
>
> ⚠️ **A FOURTH generation, and it was inside this very paragraph (M257x iter-120).** iter-115 deleted the
> *cross-reference* pin and then spent three sentences describing where things were, in bare line numbers —
> *"(`:476-496`)"*, *"the block opens at `:498`"*, *"the parenthetical is at `:500`"*. Re-derived at iter-120
> those are **`:512`**, **`:538`** and **`:540`**: the passage arguing that same-file pins rot had three of
> its own rot inside five iterations. They are now gone. **The surviving numbers — `:458`, `:459`, `:496` —
> are QUOTED HISTORY, not citations**: they are what the doc once *said*, which is a fixed fact, and that is
> the only kind of line number this paragraph is allowed to carry.

> **The demo-patch mechanism is specified in [`../ops/demo/demopatch-spec.md`](../ops/demo/demopatch-spec.md).** It is the sanctioned **zero-platform-edit escape hatch**: patch the demo's own ephemeral clone before the image build, revert after — the canonical repos are never touched. Read it before adding or re-pinning a patch. Since M217 the gate is **self-healing**: the *anchor* is the contract, the whole-file sha is only a baseline.

## Role & Responsibility

**AI Readiness** is an org-level **AI-capability diagnostic**: each member runs a **3-step onboarding/evaluation**,
which produces a per-member readiness **score (0–100)** + an **archetype**, rolled up into a **manager dashboard**
(a funnel + a Knowledge×Usage archetype matrix + per-team/per-person drill-down). It is a subsystem of the **`app`
(backend) service** — its own package **`app/internal/aireadiness/`** (split out of `internal/workforce/`; see the
⚠️ package-refactor callout above) — not a standalone microservice. It is **org-gated** (off by default)
and **member-facing** for the onboarding flow + **manager/analytics-facing** for the dashboard.

## Org enablement (the gate)

The feature is off until an org turns it on. Two gates exist — but **they do not both apply to every
surface**, and assuming they do sends you hunting for a PostHog problem the manager dashboard does not have:

1. **Org setting** — a row in `organization_settings` with `setting = 'ai_readiness'`, `is_enabled = true`
   (`app/internal/data/ent/enum/organization_settings.go:47` → `OrganizationSettingAIReadiness = "ai_readiness"`;
   checked by `isAIReadinessEnabled` in `app/internal/aireadiness/steps.go` — formerly
   `workforce/readiness_steps.go::isAIReadinessEnabled`). No row =
   off. Exposed to the FE as the GraphQL query `aiReadinessEnabled: Boolean!`
   (`resolver_ai_readiness.go` — returns `false`, not an error, for non-enabled orgs).
2. **PostHog flag** `flag_ai_readiness` — gates the **member-facing** surface only. Its **sole consumer
   repo-wide** is `apps/web/src/components/ai-readiness/data/useAiReadinessActive.ts:22`
   (`useFeatureFlagEnabled(AI_READINESS_FLAG)`; the constant is declared at
   `components/ai-readiness/aiReadiness.constants.ts:26`).

> **⚠️ The MANAGER dashboard does NOT check the flag.**
> `apps/web/src/app/(authenticated)/(verified)/ai-readiness/AIReadinessClient.tsx` contains **no PostHog
> reference at all** — measured: 0 occurrences of `posthog` and 0 of `flag_ai_readiness`. It gates on
> `orgEnabled` alone — `const { orgEnabled } = useAiReadinessEnabled(true);` at **`:135`**, **@ `next-web-app`
> `8297c684`** (byte-identical at `origin/main` `f97ba659`, so no ref resolves it away). ⚠️ **This cited
> the two lines above it until M257x iter-115** — the tail of the `// Feature gate:` comment, not the construct.
> **All four of this document's `AIReadinessClient.tsx` anchors were off by the same one-construct drift**
> and are corrected together (`:135` · `:155-156` · `:171` · `:601`); the one anchor *above* the drift point,
> `:78` `const SHOW_SECONDARY_TABS: boolean = false;`, is exact — which is what let the block read as
> verified. So **gate 1 alone**
> renders the admin dashboard; gate 2 applies to the member/onboarding surface. A blank manager dashboard is
> a gate-1 (org-setting) problem, never a PostHog one.

> **These two gates are different layers — not a contradiction.** `stories-spec.md` (the `OrgSettingsSeeder` row)
> calls enablement "an **org setting**, not a PostHog flag": that is precise about the **enablement/data layer**
> (gate 1) the seeder writes — a `organization_settings` DB row, resolved from the M48 contract, which is *not*
> stored in PostHog. It does **not** deny gate 2: the next-web **member** surface *additionally* checks the
> PostHog `flag_ai_readiness` before rendering. Seeder-writes-the-setting (gate 1) and
> member-UI-also-checks-the-flag (gate 2) are complementary — but per the callout above, **only gate 1 is
> required for the manager dashboard**. The demo must satisfy gate 2 for the member journey, which is what
> the next section is about.
>
> ### ⚠️ How the demo satisfies gate 2 (the FE flag) — CORRECTED, M219 (v2.3 "cue to cue")
>
> **This section previously asserted the exact opposite of the truth, and the error is instructive.** It said:
> *"the demo next-web bakes no `NEXT_PUBLIC_POSTHOG_KEY`, so the client-side flag check has no PostHog backend
> to consult and does not block the route"* — i.e. that absence of PostHog **defaults the flag through**.
>
> **It does not. Absence of PostHog makes the flag `undefined`, and the code demands `=== true`:**
>
> ```ts
> const rawFlag     = useFeatureFlagEnabled(AI_READINESS_FLAG);  // no PostHog → undefined, FOREVER
> const flagEnabled = stickyFlag.current === true;               // undefined === true → FALSE
> const { orgEnabled } = useAiReadinessEnabled(flagEnabled);     // queried ONLY when the flag is on
> active = flagEnabled && orgEnabled === true;                   // → never active
> ```
>
> **The construct this paraphrases, named (M257x iter-115).** It is `useAiReadinessActive()` in
> `apps/web/src/components/ai-readiness/data/useAiReadinessActive.ts:19-32`, **@ `next-web-app`
> `8297c684`** — `:22` `rawFlag`, `:25` `flagEnabled`, `:29` `useAiReadinessEnabled(flagEnabled)`,
> `:32` `active: flagEnabled && orgEnabled === true`. **This block carried no file and no ref until
> iter-115**, which is the same defect the four `AIReadinessClient.tsx` anchors in this document carried
> — a citation-free paraphrase cannot be re-derived, and cannot be caught when it drifts. **Distinct
> surface from the manager dashboard**: this is the *member* gate (gate 2); the manager dashboard
> calls `useAiReadinessEnabled(true)` directly and never reaches this hook.
>
> `Analytics.provider.tsx` initializes PostHog only when **both** `NEXT_PUBLIC_POSTHOG_KEY` and
> `NEXT_PUBLIC_POSTHOG_HOST` are present; a demo supplies neither. So on a demo the **member** AI-readiness
> surface **never mounts, for any member, in any cycle state** — and the org-enablement query is never even
> fired, because the hook short-circuits on the flag. Measured on `billion`/demo-1 (cold reset-to-seed, both
> seeded heroes, authenticated): `/home` body contains "AI Readiness" → **NO**; readiness network calls →
> **NONE**.
>
> **Why the old claim survived: it was proven against the wrong vantage.** The "empirical proof" cited was the
> M53 acceptance **AB5**, which renders the **manager** dashboard. But `flag_ai_readiness` gates the
> **EMPLOYEE side only** (`useAiReadinessActive.ts` — see § below); the manager page does not route through
> that hook at all. A manager-side render was therefore never evidence about gate 2. The doc's own parenthetical
> conceded the mechanism was *"inferred … not separately traced in the FE code"* — and the inference was wrong.
> This is the same wrong-vantage trap that made two of M219's own opening premises false.
>
> **How a demo ACTUALLY satisfies gate 2:** the sha-pinned demo-patch **`next-web-aireadiness-flag-gate`**
> (M219) widens the gate to treat *"PostHog is not configured"* as *"no rollout gate"* — behaviour-identical
> wherever PostHog **is** configured, and the ORG boolean still has the final say in every case, so a
> non-readiness org stays dark on a demo too. See [`demo/demopatch-spec.md`](../ops/demo/demopatch-spec.md).
>
> **The genuine platform limitation this records** (the patch does not erase it): a deployment without PostHog
> cannot turn AI Readiness on for members **at all**, whatever its org settings say. The real platform fix
> would be to fall back to the org boolean when the analytics provider was never initialized.

## The 3-step framework + scoring

The evaluation is a fixed **3-step** framework (per-org orderable, canonical default below), each step a scoring
axis (`enum.AIReadinessStepType` defined in `app/internal/data/ent/enum/ai_readiness.go` — `StepSkillMapping` /
`StepSimulation` / `StepInterview`; consumed + scored in `app/internal/aireadiness/readiness.go`, formerly
`workforce/ai_readiness.go`):

| # | Step (`step_type`) | Method | Max pts | Signal that completes it |
|---|--------------------|--------|---------|--------------------------|
| 1 | `skill_mapping`    | self-map AI skills (framework modal) | **30** | ≥1 `user_skill_evidences` row for a skill node_id in the org's `ai_readiness_skills` set (presence-based; level ignored) |
| 2 | `simulation`       | a job-simulation session | **40** | best ended/scored jobsim session whose sim_id ∈ `ai_readiness_sims` (`step_type='simulation'`); `(raw_score/100)×40` |
| 3 | `interview`        | an interview (a sim, different config) | **30** | best ended/scored session whose sim_id ∈ `ai_readiness_sims` (`step_type='interview'`); `(raw_score/100)×30` |

**Composite scores (0–100):** `score = step1+step2+step3` (max 100) · `knowledge = (step1+step2)/70×100` (axis X) ·
`usage = step3 scaled` (axis Y). **Archetype** (the math lives in `aireadiness/scoring.go::classifyArchetype`) — as
of the M247-refresh refactor this is **band-based, not a flat threshold-50 split** (`archetypeHighBand = 75`,
`archetypeLowCeil = 50`): **Champion** = knowledge ≥ 75 AND usage ≥ 75 · **Standby** = both ≤ 50 · else usage ≥
knowledge → **Explorer**, knowledge > usage → **Hidden Talent** (exact tie → Explorer). **Buckets/bands** likewise
changed: **None 0–24 / Low 25–50 / Medium 51–74 / High 75–100** (was none 0-10 / low 11-40 / medium 41-70 /
high 71-100 — the ≥75 Champion requirement deflated the Champion population materially; re-derive any "Champion
30/30" demo beat against these bands).

> **Track-awareness (net-new at the M247-refresh refactor).** Cycles and sims now carry a **`track`**
> (`tech` | `business` | `both`). The Step-2 sim resolves per-user by track (`resolveUserTrack`, business-wins);
> the interview (Step-3) is shared across tracks (`track = "both"`). `defaultReadinessSims` is 3 entries:
> `{simulation, "tech", …}`, `{simulation, "business", …}`, `{interview, "both", …}`. Cycles also gained a
> `launched_by`. (The demo's track↔audience mapping + the named track-keyed sims are v2.7 **M250** fidelity work.)

**Started vs completed (the funnel):** a member carries a `stage` ∈ {1,2,3} (0 = none/done) and a `score` (null =
not-completed). Per-step status lives in `ai_readiness_user_step_progresses` (`not_started`/`in_progress`/`completed`,
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
| `ai_readiness_user_step_progresses` | per-(org,user,step) progress (**plural** — the ent-generated table name; M219) | `status`, `completed_at` |
| `ai_readiness_live_snapshots` | **live** per-member score (mutable, upserted) | `score/knowledge/usage/archetype/stage/...`; exposed to Talk-to-Data SQL |
| `ai_readiness_snapshots` | **frozen** per-(cycle,user) snapshot at close (immutable) | `frozen_*` mirror of live |
| `ai_readiness_diagnose_narratives` | persisted per-member LLM narrative | keyed `(org,user,cycle_ref,lang)` + `signals_hash` |
| `ai_readiness_text_translations` | content-addressed translation cache | `source_hash`+`lang` |
| `ai_readiness_recommendations` | per-member recommended actions (the What-to-do-next drawer's "Recommended actions") | **was missing from this doc until M219**; the demo seeds **0 rows** — the live read derives `people[].diagnosis.recommendations` instead |
| `ai_readiness_email_overrides` (**net-new, M408**) | per-org email-copy override | one row per `(organization, email_type)`; `OrganizationMixin`-scoped; backs the workforce-admin email PUT/GET/DELETE + the preview renderer |
| `ai_readiness_notification_logs` (**net-new, M400**; plural — the ent-generated table name) | notification send log | the invitation/reminder/launch/digest lifecycle audit |
| `ai_readiness_notification_optouts` (**net-new, M400/M403**; plural — the ent-generated table name) | per-member unsubscribe | the reminder-cadence opt-out |
| `organization_settings` (existing) | the enablement gate | `setting='ai_readiness'`, `is_enabled` |
| **`public.interview_aggregated_reports`** | **the org's Step-3 interview AGGREGATE — the SOLE source of all four "AI Interview — breakdown" blocks** | `(organization_id, sim_id, report JSONB, session_count)`. **Added to this doc in M219 R-8; nothing had ever seeded it.** See below. |

### The Step-3 interview findings — `public.interview_aggregated_reports` (M219 R-8)

The manager's **How-we-measure → Step-3 breakdown** panel has four sub-sections. On the shipped demo **three
rendered their HEADINGS WITH NO CONTENT** and **a fourth did not render at all**, and the coverage gate **passed
it under a disclosed exception**. An empty sub-section is a **FINDING, not a pass** — the exception is gone and
the seeder fills the blocks.

**It was blamed on the wrong table.** The milestone's own DB corroboration pointed at
`public.conversation_extractions` (0 rows) — a **red herring**: that table holds transcript interaction
counts and *nothing on this surface reads it*. `interview_extraction_results` (165 rows, written by the
`SuccessionSeeder`) feeds a **different** surface. `app/internal/aireadiness/how_we_measure.go`
(`computeInterviewInsightsV2` — formerly `workforce/how_we_measure_v2.go`, folded in at the M247-refresh refactor)
reads **exactly one table**, and decodes its `report` JSONB:

| `report` key | → renders | Notes |
|---|---|---|
| `catalog_kpis[]` `{id, value}` | **"How they use AI — at a glance"** (5 tiles) | ids `avg_adoption` / `avg_transformation` / `avg_originality` / `avg_depth` / `avg_ownership` (Adoption / Transformation / Originality / Depth·Creation / Critical ownership), each a **0-100 cohort average** (the M219 ids `avg_frequency` / `avg_breadth` / `avg_context_fit` were retired in the post-M246 platform rename — see *part 5* of § *The FILLED-ness contract* below). `usageDimensionsFromReports` **omits** any KPI that is absent or non-numeric — **omitting all of them is why the tile row did not render at all**, rather than rendering empty. |
| `narrative.patterns[]` | **Strengths** | `evidence[0]` **IS** the rendered verbatim quote; `source_session_id` is what `resolveSessionAuthors` joins (`sessions → memberships`) to hydrate the quote's **author name + job role**. |
| `narrative.unexpected[]` | **Unexpected angles** | **NO chart fallback exists** — the narrative is the only way this column can *ever* render. |
| `narrative.insights[]` where `category` contains **`"risk"`** | **What holds them back** | The category string is **load-bearing**: `holdsBackFromInsights` filters on it. Get it wrong and the column silently empties again. |
| `catalog_charts[]` `top_concerns` / `top_unexpected` | the **no-narrative fallbacks** | Back What-holds-them-back / Strengths when the LLM narrative is absent. The seeded row carries a narrative, so these are belt-and-braces. |

**Seeder:** `stack-seeding/seeders/ai_readiness_interview_report.go`, flushed by the `AIReadinessFunnelSeeder`
(one row per AI-readiness org, deterministic id → `ON CONFLICT (id)` makes a re-seed a no-op). The table is in the
`--reset` list. **The honesty rules it holds to:**
- every `source_session_id` / `session_ids` entry is a **REAL seeded Step-3 session id**, so quote attribution
  resolves to a **real seeded member** through the platform's own join — never a fabricated id, never a quote
  from nobody;
- the **five** usage KPIs — `avg_adoption` · `avg_depth` · `avg_ownership` · `avg_transformation` ·
  `avg_originality`, the set `aiReadinessUsageKPIs` returns at
  `stack-seeding/seeders/ai_readiness_interview_report.go:185-189`, whose own comment at `:149` reads *"All
  five are always emitted here"* — are **DERIVED from the org's own seeded Step-3 session scores** (the same
  raw numbers the frozen snapshot rolls up), so the tiles agree with the funnel rather than being invented.
  (This bullet said *"four"* until M257x; it is the same five ids the `catalog_kpis[]` row above enumerates
  and the same five *part 5* of § *The FILLED-ness contract* re-derives.);
- `session_count` is the true number of seeded interviews;
- an org with **zero** seeded Step-3 interviews writes **no row** (nothing to aggregate — honest degradation).

The narrative prose itself is **code-owned demo copy** (like `aiReadinessInterviewPrompts`) — what a real
aggregation LLM would have synthesised. Fenced by `ai_readiness_interview_report_test.go`, which decodes the
seeded row **through the platform's own contract** (transcribed structs), because *"the seeder wrote a row"* is
not the proposition that matters — *"the row makes the four blocks render"* is.

Scoring engine: `app/internal/aireadiness/readiness.go` (`computeAIReadiness`, `GetAIReadinessWithOptions`,
`computeOrgBreakdowns`; formerly `workforce/ai_readiness.go`). Archetype/axis math + bands: `aireadiness/scoring.go`.
Steps/progress: `aireadiness/steps.go`. Cycles + auto-close: `aireadiness/cycles.go`. Narrative:
`aireadiness/narrative.go`. Compare: `aireadiness/compare.go`. CSV: `aireadiness/csv.go`. One-click provisioning:
`aireadiness/provision.go`. (All formerly `internal/workforce/…` — see the ⚠️ package-refactor callout.)

## Interface

- **GraphQL** (`app/internal/web/backend/graphql/graph/schemas/ai_readiness.graphqls`; resolver
  `resolver_ai_readiness.go`): `aiReadinessEnabled`, `aiReadinessUserPlanProgress` (member step status + deadline),
  `aiReadinessSkills` (skills to map), mutation `completeAiReadinessSkillMapping`.
- **REST/workforce API** (`app/internal/web/backend/api/`): `GET /api/workforce/ai-readiness` (→ the
  `AIReadinessResponse` aggregate the manager dashboard consumes), `/cycles` (GET/POST) **+ `/cycles/{cycleID}`
  (GET) + `/cycles/{cycleID}/close` (POST)**, `/steps-completion`, `/narrative` (POST, LLM diagnosis), `/compare`,
  `/export.csv`, and **net-new at the M247-refresh refactor: `/setup` (GET status + POST one-click provision)** plus
  the **M407/M408 email-preview + email-override admin** endpoints.
- **Background:** live-snapshot refresh (`RefreshLiveSnapshots`, worker task) **+ the net-new auto-close scheduler
  `CloseDueAIReadinessCycles`** (sweeps active cycles past `end_date`) + the notification workers (reminder/digest).

## Platform refresh — net-new subsystems (aireadiness-package refactor, app v1.351.1, M247)

The refactor did more than move files — it consolidated several net-new platform subsystems that had **no prior
corpus coverage**. Documented here as platform facts; the demo-seeder consequences are v2.7 **M250** work.

- **The platform DEFAULT provisioning is 31 readiness skills** (`aireadiness/defaults.go` `defaultReadinessSkills` =
  **19 core @ weight 1.0 + 12 enabling @ 0.5**) + **3 default sims** (2 track-keyed simulations + 1 shared
  interview). **v2.7 M250 brought the demo seeder up to this platform default** — it now seeds the full
  **31-skill repertoire** (19 core + 12 enabling, denominator **25.0**) + the **3 real track-keyed default sims**,
  replacing the M51 demo seeder's much smaller *invented* set (~5 core + a few enabling, denominator 6.5). The
  "Seeding contract" section below describes the *demo seeder as it stands after M250*; the platform default it
  now mirrors is described here.
- **One-click self-service provisioning** — `aireadiness/provision.go` `ProvisionAIReadiness` (+ `GetAIReadinessSetupStatus`)
  seeds the 31-skill default + 3 sims + the 3-step plan idempotently, behind the new `/setup` GET/POST endpoints.
- **The notifications lifecycle** — `aireadiness/notifications/` (invitation fan-out, a 5-slot reminder cadence +
  unsubscribe, launch confirmation, the weekly manager digest), backed by the `ai_readiness_notification_logs` +
  `ai_readiness_notification_optouts` tables, emitting proto events consumed by the messenger **domain** — in-process inside `backend` since the v9.0 fold and switch-gated by `MESSENGER_ENABLED`, not a separate service (its container went at `838d907`).
- **Email overrides + preview** (M407/M408) — `aireadiness/emailoverride/` + `emailpreview/`: per-org email-copy
  overrides (`ai_readiness_email_overrides`) validated against **`app/internal/messenger/aireadinessemail`**
  placeholders — imported under exactly that path at `aireadiness/emailoverride/emailoverride.go:29`, **@ `app`
  `ad9f3c49`**, where `git ls-tree ad9f3c49 internal/messenger/aireadinessemail/` lists 9 files. ⚠️ **This read
  `messenger/pkg/aireadinessemail` until M257x iter-115**, with **no `internal/` and no `app/`** — a path that
  resolves only in the standalone `messenger` repo @ `fa47850d`, which `838d907` removed from `repos.yml`, so
  `make init` does not clone it and the citation resolves to nothing on a stack. It is also the odd one out in
  its own bullet list, every sibling of which is `app/internal/`-relative under a heading that scopes the list
  to the live app packages — and two bullets above, this document has just said messenger is *"in-process
  inside `backend` since the v9.0 fold … not a separate service."* (The platform's own comment at
  `emailoverride.go:33` uses the same `pkg/` shorthand, which is where the error came from; it does not make
  it resolve.) With
  an admin preview renderer.
- **Cross-cycle Compare is a fully-built backend** — `aireadiness/compare.go` `CompareCycles` → a 6-section
  `AIReadinessCompareResponse` (Topline / Archetypes / a 4×4 Transitions matrix / TeamDelta / SkillCoverage /
  ThemesShift); both cycles must be `closed`. (The FE Compare tab is still hard-gated off — see § Surfaces.)
- **The auto-close scheduler** — `CloseDueAIReadinessCycles` sweeps active cycles past `end_date` (cross-org under
  `privacy.Allow`).
- **The personalized Academy recommendation engine** — `aireadiness/recommendation_engine.go` + `recommendation_signals.go`,
  embeddings-based via `academy.EmbeddingsManager` (see [Academy backend](./academy-backend.md)), frozen at cycle
  completion.
- **CSV export is now 15 columns** (`aireadiness/csv.go` `ExportAIReadinessCSV`; the recommendation columns were
  dropped from the earlier 19), UTF-8 BOM, formula-injection-neutralized.
- **Manager wiring** — `aireadiness.Manager` (`aireadiness/manager.go`) is constructed with `db`, `ent`, `orgRepo`,
  `ai.AI`, a pub `EventPublisher`, the `academy.EmbeddingsManager`, and the `WorkforceDirectory` (member directory).

## Surfaces (UI) — **current vs legacy** (M219, v2.3 "cue to cue")

> ⚠️ **HISTORICAL (M219) — there is now only ONE manager dashboard.** For four releases every AI-readiness
> demo pointer (the cockpit deep-link catalog, the manager hero's `jump_to`, the coverage sweep's page
> descriptor) targeted a **legacy** dashboard that still rendered but was not the shipped product. **That
> surface no longer exists:** next-web-app `dae0fb2f7` (*"drop orphaned container"*, 2026-07-13) deleted
> `AIReadinessContainer.tsx`, `AIReadinessIntro.tsx` and `AIReadinessView.tsx` (−653 lines), and
> `/enterprise/workforce/ai-readiness` now **404s**. The section is kept because the *class* of defect —
> an unlinked orphan surface that renders — is the one worth recognising, not because there is still a
> choice to make.

| Vantage | Surface | Route | Status |
|---------|---------|-------|--------|
| **Manager** | **`AIReadinessClient`** — HeroCard (org score + dominant archetype + **Steps-Completion %**) + tabs **Snapshot** (archetype matrix + donuts + by-tag), **How-we-measure** (3-step ribbon + skill strengths/gaps + sims + **interview findings**), **What-to-do-next** (archetype action groups + per-person **Diagnose** drawer). **Cycle-aware.** | **`/ai-readiness`** | ✅ **CURRENT** |
| **Manager** | `AIReadinessContainer` → `AIReadinessView` — pre-v3.0 org-summary card + team table. | `/enterprise/workforce/ai-readiness` | 🗑️ **DELETED** at next-web `dae0fb2f7` (2026-07-13); the route 404s |
| **Employee** | `AIReadinessHero` (the 3-step funnel; modes new/progress/done/archived) + `AIReadinessRailCard`. **NO ROUTE OF ITS OWN — both are EMBEDDED in `/home`.** Step 1 = skill-mapping modal, Step 2 → a sim, Step 3 → an interview. | **`/home`** | ✅ **CURRENT** |

**How the orphan was identified, while it still existed** (there was no `@deprecated` marker, no `-v2`
naming and no feature flag switching between them — the legacy one was simply *unlinked*). Retained as the
recognition pattern. **The three orphaned COMPONENTS are deleted; the surrounding anchors below still
resolve** — `urls.ts`, `useNavbarSections.tsx`, the e2e spec, `WorkforceNewClient.tsx` (still
omitting readiness) and `useWorkforceAIReadiness.ts` (still cycle-less) are all live, so they remain
checkable evidence rather than history. **Every line number in the three bullets below is a reading at
`next-web-app` `8297c684`** (re-derived 2026-08-06) — this block used to say *"at HEAD"*, and a bare moving
label is what let the first anchor below rot through four readings:

- **`/ai-readiness` is the only readiness route the navbar links** — `AI_READINESS_URL`
  (`packages/core-js/src/constants/urls.ts:51`), consumed by `packages/ui/src/NavBar/useNavbarSections.tsx`
  — imported at `:4`, built into `aiReadinessMenuItem` at `:401-408` (`key: AI_READINESS_URL` at `:403`),
  gated at `:569` (`showAIReadiness ? aiReadinessMenuItem : null`). ⚠️ **This doc said `urls.ts:52` for four
  readings, and `:52` has never been `AI_READINESS_URL` at any ref reachable from this clone** — over every
  commit touching that file the constant sits at line **41, 50 or 51 only**; `:52` is `TALK_TO_DATA_URL` at
  `8297c684` and was `ORGANIZATION_FEEDBACK_URL` at `bb3313bc`. The citation *resolved*,
  which is why an anchor-existence check passed it while a reader following it landed on a different route;
  corrected M257x iter-102. A repo-wide grep finds the constant in exactly those
  two **source** files — the only other non-`node_modules` hit is the platform's own KB page,
  `knowledge/ai-readiness/frontend-architecture.md:15`. `/ai-readiness` is also the only readiness route
  next-web's own e2e covers (`e2e/specs/web.ai-readiness.spec.ts`).
- **The legacy route was an orphan**: no nav entry, no workforce tab (`WorkforceNewClient.tsx:125-151` omitted it),
  no redirect points at it. Its hook (`hooks/useWorkforceAIReadiness.ts:23-27`) calls
  `GET /api/workforce/ai-readiness?tag=` — **there is no `cycle` param in it at all**, and it never calls `/cycles`.
- The `(new)` in the legacy path is a Next.js **route group** for the workforce refactor — *not* a version marker.
  Don't read it as "the new one".

**The `flag_ai_readiness` PostHog flag gates the EMPLOYEE side only** (`useAiReadinessActive.ts:22`). It does
**not** select between manager trees — it never did, back when there were two. The (one) manager
dashboard gates purely on the GraphQL `aiReadinessEnabled` boolean plus `isEnterprise` nav visibility.

**Also present but not user-reachable:** a 4th manager tab, **Compare** (cycle deltas), is fully built but
**hard-gated off** — `AIReadinessClient.tsx:78` `const SHOW_SECONDARY_TABS: boolean = false;` (exact), read by
the tab filter at **`:601`** — `(SHOW_SECONDARY_TABS && tab.key === 'compare')` — which strips it from the tab
list. **@ `next-web-app` `8297c684`.** (This carried a line pin until M257x iter-115; that line is `tab.key === 'how' ||`,
two lines above the read.)
`/ai-readiness?tab=compare` renders no panel. It is neither current nor legacy: complete-but-disabled.

**The demo's pointers** (all repointed at the current surfaces in M219, and a legacy target is now a **hard
failure** — `stack-seeding/seeders/cockpit.go` `LegacyReadinessPaths` / `ValidateCockpitManifest`):
`stories.seed.yaml` (Dana → `/ai-readiness`; Aria + Ben → `/home`) · the cockpit deep-link catalog (which
gained the **missing** end-user readiness entry) · `stack-verify/e2e/lib/coverage-manifest.ts`.

## Narrative generation

Per-member manager-facing narratives are **persisted, not regenerated per read** (`aireadiness/narrative.go`,
formerly `workforce/readiness_narrative.go`):
a sha256 of the member's signals keys a read-through cache in `ai_readiness_diagnose_narratives`; on a miss/stale
hash it calls the LLM (GPT-5-Mini) and upserts. On AI error it returns empty + the FE falls back to static
per-archetype guidance — **so a demo with no AI key still renders** (narratives just show the static fallback).

## Local development

Enable for an org by inserting the `organization_settings` row (`setting='ai_readiness', is_enabled=true`) +
`flag_ai_readiness` on in PostHog (or the local flag shim). The member flow then needs `ai_readiness_skills` +
`ai_readiness_sims` config + an active cycle; the dashboard reads `GET /api/workforce/ai-readiness`. Tests:
`app/internal/aireadiness/*_test.go` (the scoring/steps/cycle suites moved with the package; Postgres harness
`testdb_test.go` via `internal/testsupport/pgtest`).

## Seeding contract (demo / M51 → 31-skill fidelity, v2.7 M250)

To make a **200-person demo org** show the AI-readiness manager dashboard **enabled**, with **78.4% (≈156 of the
199 frozen snapshots) having completed all 3 steps** (the shipped figure — see `seeding-spec.md`; this supersedes
the earlier round "~80%/≈160" contract prose), plus **one hero "started"** and **one hero "completed"** — the
seeder writes:

**Org config (≈10 rows):**
1. `organization_settings` (`setting='ai_readiness'`, `is_enabled=true`) — the gate.
2. `ai_readiness_steps` × 3 (skill_mapping/simulation/interview, positions 0/1/2) — optional (canonical default if absent).
3. `ai_readiness_skills` × **31 — 19 core (weight 1.0) + 12 enabling (0.5)** (M250; was ~5 core + a few enabling),
   the platform's own `aireadiness/defaults.go` `defaultReadinessSkills` mirrored **verbatim**. `node_id` = **real
   taxonomy node-ids** (route through the existing seeding resolvers — never fabricate, per the closure gate; here
   they are the platform's real defaults, so it is no-fabrication-*by-construction*, and the config seeder writes
   the 31 unconditionally as `provision.go` does).
4. `ai_readiness_sims` × **3 track-keyed** (M250; was 2) — **two distinct-track simulations**
   (tech=`who-can-see-this-document-fc0`, business=`use-ai-to-turn-survey-data-into-a-leadership-email`) **+ one
   shared interview** (`ai-readiness-interview-d62`), each pinned to the platform's real default sim uuid, with the
   `track` column WRITTEN (the schema's `UNIQUE(org, step_type, track)` requires the tech+business pair — both
   `step_type=simulation` — to carry DISTINCT tracks). **Plus a net-new Directus set-dress**
   (`AIReadinessSimSkillsSeeder`) that resolves each wired sim's `directus.sequences.evaluation_skills` node-ids →
   `public.skills.name` and UPDATEs `directus.simulations.skills` for the 3 uuids — this is what fills step-2's
   **evaluated-skills list** and drives the tech/business label **name-heuristic** (`directus.simulations.skills`
   is genuinely NULL in capture, so replay alone leaves it empty — snapshot replay is replay-only, hence a net-new
   write step).
5. `ai_readiness_cycles` **× 2 — one `closed` AND one `active`** (v2.3 M219; see § *The CYCLE-STATE contract —
   seed BOTH cycles* below, which this line used to contradict). The **active** cycle (future `end_date`,
   `participants_filter {"all":true}`) is the one the manager dashboard resolves and LIVE-recomputes — it is what
   fills Ben's funnel, Aria's full hero card, and the `interview` / `diagnosis` / `sources` sub-sections that were
   NULL or absent under closed-only. The **closed** cycle is retained as cycle *history* (its 199 frozen snapshots
   stay meaningful in the CyclePill). M51 shipped `closed` alone on the belief that the active-signals path was
   falsified — **M219 measured the live recompute at 2.09 s and refuted that.**

**Per-member (≈156 of 199 "completed"):** the underlying signals (≥1 `user_skill_evidences` for a configured skill;
jobsim sessions for steps 2/3) **+** `ai_readiness_user_step_progresses` (3× `completed`) **+** an
`ai_readiness_live_snapshots` upsert (`score≈100, stage=3, archetype` per the score). The **"started" hero**: only
the skill_mapping signal + a `stage=1`/`score≈30` live snapshot. The **"completed" hero**: all 3 + `stage=3`.

**⚠️ A hero's stage is DERIVED FROM HER PERSONA KIND — and stage 0 was UNREACHABLE for an end-user until v2.8
M256 iter-26.** `aiReadinessStageFor` maps a hero: **manager → 0** (she reads the dashboard, she is not a funnel
row), **struggling → 1**, and **everything else → 3**. There is no "unspecified" middle: an end-user hero appended
to a readiness org arrives **already COMPLETED**. That is a trap for test authorship rather than for the demo —
a Playthrough built on such a seat drives its hero *past* the flow it means to watch her enter and can **PASS on
the completed surface**. It is now declarable:

- **`Persona.AIReadiness`** — `""` (derive from trajectory, unchanged) | **`"not_started"`** (funnel **stage 0**:
  no evidence, no sessions, no progress row, no snapshot — `seedAIReadinessOrgFunnel` skips her entirely). Checked
  **before** the derived default, read in **exactly one place**, and a **closed enum** (a typo would fall back to
  stage 3, i.e. produce the very already-completed seat, so it fails the seed loudly).
- It is deliberately **not** a third `trajectory` value: readiness stage is **orthogonal** to the life-arc — the
  diagnostic is a time-boxed cycle, and a member who joined mid-cycle has verified skills, a career and an
  activity history while having done none of its three steps. A new trajectory would additionally have claimed her
  skills, level band, growth arc and succession signal were day-0 too, and would have rippled through the five
  seeders that switch on `Trajectory` through a silent `default:` branch in each.
- An end-user hero must still declare `verified > 0` (`blueprint.validate()`), which is the same point from the
  other side: the field does not make a hero **blank**, it makes her **not-yet-diagnosed**.

Fenced by `stack-seeding/seeders/seat_append_test.go` (the mapping, both directions) +
`ai_readiness_funnel_test.go:TestAIReadinessFunnelSeeder_DayZeroHeroGetsNoSignals` (the consequence — she has
zero rows *while the derived heroes still land 3 and 1*, so a seeder that wrote nothing cannot pass it). The
consumer is the Playthrough seat `pt-ai-onboard`; see
[`demo/playthroughs.md`](../ops/demo/playthroughs.md) § *The day-0 readiness seat*.

**⚠️ Which table the dashboard reads depends on the cycle state — this dictates the seed strategy (an M51
decision):**

- **Active cycle → the dashboard RECOMPUTES from signals.** `GetAIReadinessWithOptions` → `buildLiveResponse` →
  `computeOrgBreakdowns` (`aireadiness/readiness.go:330`) re-derives each member's score **from the underlying signals**:
  `user_skill_evidences` (step 1) + the readiness jobsim sessions (steps 2/3) + the `ai_readiness_skills`/
  `ai_readiness_sims` config — and, **for an active cycle, the membership filter is `keepInCycleStep1`**
  (call site `readiness.go:388`; **`func` at `:710`**): it keeps only members who **FINISHED step 1 inside
  the cycle window** (`completed_at >= cycle.start`), scores carried forward all-time. `keepStartedMembers`
  is the **other** branch — the lifetime filter used when there is **no** active cycle (call site `:390`;
  **`func` at `:730`**). This passage named the wrong one of the two until M257x iter-52, and carried the
  wrong `func` anchors until iter-108 — they landed in the body of `audienceScope` (`:655`) and on
  `queryInCycleStep1Completers`, the very function the sentence contrasts this branch against.

  ⚠️ The no-active-cycle filter `keepStartedMembers` **excludes members with no PROGRESS ROW** from the
  aggregate, and **reads no step-1 evidence signal at all**, which this sentence claimed until M257x iter-49:
  `queryReadinessStarters` (`aireadiness/steps.go:964`, its SQL at `:969-974`) is
  `SELECT DISTINCT user_id FROM public.ai_readiness_user_step_progresses … AND status <> 'not_started'` — a
  row in the **progress** table, never a `user_skill_evidences` row. The old wording was wrong in **both**
  directions: a member with step-1 evidence but no progress row is **dropped**, and one with an
  `in_progress` row and zero evidences is **kept**. (This paragraph **previously said** the active-cycle
  path used this filter; that attribution is **refuted** — see the `keepInCycleStep1` correction above.) The platform states this itself at `steps.go:962`
  (*"This DB signal is the only real 'has started' check"*) — this doc cited a different function's doc comment until M257x
  iter-108 — that of `queryReadinessSimCompleters` (`:922`), a
  `job_simulation_sessions` query about a different proposition. So an **active**-cycle dashboard requires the
  **signals-true** seed — and the `ai_readiness_user_step_progresses` rows are what actually decide who is
  counted, not the evidences. **Both branches read that one table, at different strictness**, and the
  active-cycle one is the strict one: `queryInCycleStep1Completers` (`readiness.go:684`) demands
  `step_type = skill_mapping` **AND** `status = completed` **AND** `completed_at >= cycle.start`, so a
  progress row seeded `in_progress` — enough for `keepStartedMembers` — leaves the member **invisible**
  on an active cycle (write the real skill evidences +
  sim sessions + `ai_readiness_user_step_progresses`; reuse the existing verified-skill chain). `ai_readiness_live_snapshots`
  is a **materialized cache** (rewritten by `RefreshLiveSnapshots`, consumed by Talk-to-Data SQL) — **NOT** the
  dashboard's source: seeding it directly does **not** make the live dashboard render and is overwritten on refresh.
- **Closed cycle → the dashboard reads frozen snapshots.** `buildResponseFromSnapshots` reads `ai_readiness_snapshots`
  directly, so a **closed**-cycle showcase can be seeded **snapshot-direct** (write the `frozen_*` rows + flip the
  cycle to `closed`) with **no underlying signals** — the world reads as a *finished* assessment. **This is the
  strategy M51 shipped** (`AIReadinessConfigSeeder` writes the cycle `closed` + `AIReadinessFunnelSeeder` writes 199
  frozen `ai_readiness_snapshots`), after iters 03→06 falsified the active-signals path — on the premise that the
  live-recompute never completed in the coverage harness budget (a per-skill federated translation N+1, the M46
  per-object-RPC class). ⚠️ **That premise was refuted at M219: the recompute takes 2.09 s.** The strategy M51
  shipped is unchanged; the reason given for it is not. Retracted here and at
  [`ops/demo/stories-spec.md:599`](../ops/demo/stories-spec.md) at M257x iter-49.

  **The frozen path is reachable BOTH cycle-scoped and by default.** `GetAIReadinessWithOptions`
  (`app/internal/aireadiness/readiness.go:289`) has two routes into `buildResponseFromSnapshots`:

  1. **cycle-scoped** — `opts.CycleID != nil` AND that cycle's `status == "closed"` (`:290-297`);
  2. **default (`CycleID == nil`)** — when the org has **no active cycle** and a **latest closed cycle**
     exists (`:307-312`). Only an org with neither falls through to `buildLiveResponse` (`:314`).

  Route 2 is exactly the shape M51 seeds, so a snapshot-direct closed-cycle showcase renders frozen data
  **even when the caller passes no cycle id**. The current manager dashboard passes the cycle id anyway.

  > **✅ CORRECTED M219 (v2.3 "cue to cue") — the old M51 iter-07 caveat here was MISATTRIBUTED, and it sent a
  > later milestone hunting for a demo-patch that was never needed.**
  >
  > The retracted claim: *"the demo FE fires the data GET WITHOUT `?cycle=` … and never fires the `/cycles` list
  > that supplies `latestClosedCycle.id`"*, concluded to be **platform-bound**.
  >
  > **What is actually true.** The **CURRENT** dashboard (`AIReadinessClient.tsx:155-156`, **@ `next-web-app`
  > `8297c684`**) computes
  > `effectiveCycleId = selectedCycle ?? activeCycle?.id ?? latestClosedCycle?.id` and gates the data GET on
  > `enabled: featureOn && cyclesQ.isFetched` at **`:171`** — i.e. it **waits for `/cycles`, then passes
  > `?cycle=`**. (These carried two line ranges until M257x iter-115: the first was the comment above
  > the const, and the second a range that stopped **one line short** of the gate it is offered as evidence
  > for — `:168-170` is the `useAIReadiness({` call and its first two options.) Verified live
  > against a running demo (authenticated as the manager hero): `/cycles` returns the seeded cycle, and the
  > frozen read answers **HTTP 200 in 24 ms**.
  >
  > The iter-07 probe was watching the **LEGACY** page (`/enterprise/workforce/ai-readiness`), whose hook
  > (`useWorkforceAIReadiness.ts:23-27`) has **no `cycle` param at all** and **never calls `/cycles`** — which is
  > exactly the behavior that was observed and then attributed to the platform. **It was a pointer bug, not a
  > platform gap.** See § Surfaces (UI) above.
  >
  > **And the live path does not "never complete".** Measured on the same 199-member org:
  > **LIVE `GET /api/workforce/ai-readiness` → HTTP 200 · 2.09 s · 304 KB.** The M51-era "translation-N+1 that
  > never completes in-budget" is **not reproducible** on the app tag the demo builds today. Re-measure before
  > relying on either number; do not re-derive them from prose.

  **⚠⚠ M51 iter-08/09 — the frozen READ is ITSELF org-scale-slow ("frozen" froze the SCORES, not the RESPONSE).**
  Even when the frozen branch IS selected (a direct `?cycle=<closed>` GET), `buildResponseFromSnapshots`
  (now `aireadiness/readiness.go`, formerly `workforce/ai_readiness.go:512`) reads the frozen scores fast but then
  calls **`m.workforce.LoadMembers(ctx, orgID, "")`** *through the `WorkforceDirectory` interface* (the `LoadMembers`
  implementation stayed in `workforce/members.go`) — an
  **unbounded whole-org member hydration** (`hydrateMembers` over ~200 members) to re-join current tags/name/role
  onto each snapshot. At 200 members that member-load is the **same org-scale wall** as the live path: the
  `?cycle=<closed>` GET timed out at 180 s (iter-08's authenticated dual-endpoint probe). It is NOT the
  demo-patchable per-object targetRole Sentinel RPC (`queryBaseMembers` reads `jobRole` from a SQL column). **In the
  demo**, M51 iter-09 bounds it with the `app-aireadiness-snapshot-loadmembers` app read-path demo-patch
  (`loadMembers(orgID,"")` → the bounded sibling `loadMembersByUserIDs` over the ~199 snapshot user-ids — a pure,
  data-identical perf optimization; 180 s → 19 ms). **In PROD** the frozen read still hydrates the whole org and
  would need `loadMembers` bounded in the snapshot path, or a **`frozen_tags jsonb` column** so the snapshot read
  needn't re-join live members (**M314b** — a disclosed demo-perf relaxation, NOT a prod fix). See
  [`../ops/demo/coverage-protocol.md`](../ops/demo/coverage-protocol.md) (the iter-08/09 loadMembers lesson) +
  [`../ops/demo/stories-spec.md`](../ops/demo/stories-spec.md#the-ai-readiness-showcase-org--the-3rd-story-v110b-fit-up-m51)
  (the seeder + demo-patch).

### The CYCLE-STATE contract — seed BOTH cycles (M219, v2.3 "cue to cue")

**The two vantages need opposite cycle states, and one cycle cannot serve both.** The demo therefore seeds
**one CLOSED cycle + one ACTIVE cycle** per readiness org (legal: the *one active cycle per org* partial unique
index permits it).

**Why an ACTIVE cycle is mandatory — the member surface does not exist without one.** `AIReadinessHero` is gated
on `deadline`, and the backend derives `deadline` **only** from an active cycle
(`aireadiness/steps.go` `queryActiveCycleEndDate` → `StatusEQ(active)` → `IsNotFound` → `nil`; formerly
`workforce/readiness_steps.go`).
`deriveMode` (`useAIReadiness.ts:48-62`) then treats a **null deadline as "deadline passed"**. So against a
**closed-only** org:

| Hero | Steps done | Mode | What renders |
|------|-----------|------|--------------|
| the **COMPLETED** hero | 3 / 3 | `archived` | only the compact right-rail mini-card — **not** the full done-hero |
| the **STARTED** hero | 1 / 3 | `progress` | **NOTHING.** `AIReadinessHero.tsx:88` `if (!air.deadline) return null;` |

The started hero — the entire point of the persona — was **invisible**, and no gate caught it, because an absent
section is not an error. The active cycle's `end_date` **must be in the future** (it *is* the member deadline)
and its `participants_filter` must stay `{"all":true}` (a tag-scoped cycle returns a nil deadline for anyone
outside the tags, silently re-hiding the surface for most of the org).

**Why the CLOSED cycle is retained:** it owns the frozen 199-snapshot showcase and gives the dashboard a real
cycle *history* in the picker; a `?cycle=<closed>` read still answers off the frozen rows in ~24 ms.

**What the manager then reads — and why that is the point.** With an active cycle present, `AIReadinessClient`'s
`activeCycle?.id ?? latestClosedCycle?.id` resolves the **active** id, so the dashboard takes `buildLiveResponse`.
That is **deliberate**: the frozen read returns **six sub-sections as null**, and the dashboard renders them as
*absent*:

| API field (FROZEN → LIVE) | Sections it feeds |
|---|---|
| `howWeMeasure.interview` — **null** → present | the whole **Step-3** block + **"How they use AI"** + **"What holds them back"** + **"Strengths"** + **"Unexpected angles"** |
| `people[].diagnosis` — **missing** → present | the Diagnose drawer's **"Recommended actions"** |
| `people[].sources` — **missing** → present | the Diagnose drawer's **"Assessment sources"** (else grey "not started" cards) |

Cost: the manager data-load goes **24 ms → ~2.09 s** (measured, 199 members). **Reported, not gated** — the
milestone that owns login speed is a different one. Both paths fill `org.*`, `byTeam` (13), `people` (199),
`howWeMeasure.{steps,skillInsights,simulations,cycleTotals}`.

**No AI keys needed either way** (diagnosis narratives fall back to static per-archetype text on AI error).

### The FILLED-ness contract — three ways a readiness seed reads as real but is not (M219)

The M219 bar was *"every element and sub-section filled with spot data"*. Raising it turned three quiet
mis-seeds into defects. Each is a **seeder** contract, and each is now fenced by a regression test that was
**proven RED against the pre-fix code**.

**1. A member maps SEVERAL AI skills — not one.** `computeTier1` (`aireadiness/readiness.go:139`) divides the
member's **held** skill-weight by the org's **entire configured repertoire** — **since M250 the platform's real 31
defaults: 19 core @ 1.0 + 12 enabling @ 0.5 = `25.0`** (was the invented 8-skill / `6.5` set) — normalized to 30.
So *one* core skill alone is `round(1.0/25.0*30)` = **1/30**, and one *enabling* skill is also **1/30**: the
single-skill floor. A seeder that wrote exactly one evidence row per member would score the COMPLETED hero, the
org's showcase **"Champion"**, at that **1/30** floor. Non-empty, and not believable.

> **Full marks require holding EVERY configured skill.** The denominator is the whole repertoire, by design —
> "a larger configured set makes a full score harder to reach" (the platform's own comment). A seeder that
> ignores this produces a technically-populated, semantically-broken funnel.

The held-count is now **stage- and hero-aware** (re-derived at the 31-skill / 25.0-weight repertoire — M250): the
COMPLETED hero maps the full repertoire (**30/30**), the STARTED hero **9** of the 19 core skills
(`round(9.0/25.0*30)` = **11/30** — visibly behind the Champion, well clear of the 1/30 floor; was 3 core → 14/30
at the old 8-skill set), and the population spreads by funnel stage (stage 3: 19…31; stage 2: 3…19; stage 1:
1…3). Heroes start core-first (deterministic — a hero's score is a story beat, not a sample); the population
rotates its window so the org's per-skill strengths spread across the repertoire. The **frozen snapshot's
`frozen_step1` is now COMPUTED from the same held weight** (it was a flat constant 5) — so a frozen row and a live
recompute of the same member finally agree. (The double-round divergence the old 6.5 denominator could trip is
**arithmetically unreachable at 25.0** — held/25×100 = held×4 is always integral for half-step weights — so
M250 converted that stale hardcoded fence-triple into a live invariant.)

**2. The readiness sims must be RESERVED — or an unrelated session silently scores a member.** The platform
scores Steps 2/3 from **any** ended session whose `sim_id` is in the org's `ai_readiness_sims` set. It does
**not** consult the step-progress row. The generic session seeders hash their `sim_id` out of the *same* ~50-id
replayed content pool the readiness config draws from — so a member's **unrelated activity session** could land
on the readiness sim by coincidence and score them against a step they never took. That is exactly what
happened to the STARTED hero: his funnel row said `interview: not_started` while the backend read an interview
signal (score 21) off a stray activity session. The two readiness refs now come from a **reserved tail** of the
sims pool that no general picker can draw (`contentref.go`), making the sets **provably disjoint**. The fence is
**structural, not statistical**: asserting "no seeded session happened to collide" clears by luck about one run
in ten.

**3. An interview session with no turns is incoherent data.** `computeCycleTotals`
(`how_we_measure.go:253-261`) counts `interviewQuestions` as `COUNT(public.interactions)` joined through
sessions to the org's interview sim — the live query at `how_we_measure.go:285-287` reads
`FROM public.interactions i JOIN public.job_simulation_sessions s`. The funnel seeded the **session** and not one
interaction, so the field was a hard **0**. The funnel now writes each stage-3 interview's two
`public.actors` (the AI interviewer + the member — the interaction FKs *require* them, and the DB enforces
`source_id <> target_id`) and **6–11** `public.interactions` turns (`action_type='call'` — the platform's enum is exactly `{email, call}`).

> **Measured, not assumed — and it corrects the finding that opened this thread.** The **current** dashboard's
> *"✨ Handled for you this cycle"* tile renders **`skillsMapped` / `handsOnMinutes` / `interviewMinutes`** —
> and **does not render `interviewQuestions` at all** (**all anchors in this parenthetical re-derived at
> `next-web-app` `8297c684`, 2026-08-06**: `HowWeMeasureTab.tsx:1879` opens the
> `{/* ===== C · Handled for you this cycle ===== */}` block, label `:1903`, then exactly three cells —
> `skillsMapped` `:1915`, `handsOnMinutes` `:1921`, `interviewMinutes` `:1927`; `grep -c interviewQuestions`
> over that 1,989-line file returns **0**. The field exists in the API and in the FE's TypeScript type,
> `apps/web/src/hooks/useAIReadiness.ts:326` — inside `export interface AIReadinessCycleTotals` (`:323-330`)
> — and is drawn by nothing. ⚠️ **Read that `:326` as a pin, not a standing line**: this doc carried two
> different line numbers for it in successive iters, and at each ref the *other* number named an unrelated
> construct. **The interface name is the contract; the line number is the convenience** — and the earlier
> number is deliberately no longer reproduced here (M257x iter-141): quoting it kept it live enough to rot,
> and it rotted onto a blank line, turning two fences RED. `§5` rule 63(c)). So its zero was a
> **payload** zero, not a visible empty cell. Filled regardless — an interview with no questions is not real
> data — but the honest claim is that this tile's *visible* zero-risk lives in the three cells that do render,
> which the coverage sweep now fences with a **non-zero-value** assert rather than a label assert (a section
> that renders with all-zero numbers is an empty section wearing a hat).

**Also (a latent hazard closed while scaling #1):** the funnel's Step-1 evidence UPSERT is now
**presence-preserving** — on a conflict with a row the verified-skill chain already wrote it asserts only that
the row exists and is verified, and leaves `level` / `anthropos_level` / `user_level` alone. Step 1 is
presence-based (`queryUserAISkills` selects only `user_id, skill_id, is_verified`), so preserving is both
correct and safer: with a member now mapping up to **31** skills (M250), the old clobbering upsert would have let
the readiness seeder quietly restate a hero's claimed-vs-verified gap.

**4. Evidence distribution — the completed sim's verified skills must reach the profile (M250).** Step-1
presence is not enough; the member's profile and the manager's per-skill dots read a **validation fan-out** off
the completed Step-2/3 sessions. A net-new `ai_readiness_evidence.go` seeder distributes each completed member's
sim **evaluated** node-ids as: one `validation_attempt_results` + one `validation_attempt_skill_results` per
evaluated node-id (the score `computePerformanceBySkill` reads for the manager's dot ratings) + one
**session-backed verified** `user_skill_evidences` per node-id (the profile's verified skill). The evaluated
node-ids are read **closure-safe** from `directus.sequences.evaluation_skills` JOINed to `public.skills`
(unresolved ids drop, never fabricated — the same truthful source as the set-dress). On a cold reset-to-seed this
moved AI-sim `validation_attempt_results` **5 → 345** (+ 897 `validation_attempt_skill_results`, 787
session-backed verified evidences), flipping both the member-profile verified-skill distribution AND the manager
per-skill dots to render real.

**5. Interview-findings must key on the CURRENT KPI ids (M250 iter-07).** The manager's "interview findings"
tiles key on the platform's live `usageDimSpecs` — **avg_adoption / avg_transformation / avg_originality /
avg_depth / avg_ownership** (Adoption / Transformation / Originality / Depth·Creation / Critical ownership). The
M219 seeder wrote the now-retired ids `avg_frequency / avg_breadth / avg_context_fit`; only `avg_depth` survived
the post-M246 platform rename, so 4 of 5 tiles rendered empty. The seeder now emits the 5 current ids (spread
across bands). One of three **post-M246 drifts** the M250 fidelity sweep caught (the other two — the
`by-tag`/`handled-for-you` coverage-manifest phrasing — were manifest-copy reconciles); their **live**
manager-sweep confirmation is a slow ~150-page crawl routed to **M254** (its exit gate re-runs the same sweep on
billion by design).

**End-to-end proof:** the AI-readiness journeys are covered by **5 Playthroughs** (both member vantages + the
manager — see [`../ops/demo/playthroughs.md`](../ops/demo/playthroughs.md#the-ai-readiness-product-m219--and-why-a-blind-area-is-the-worst-kind-of-gap)
— plus the **day-0 guided onboarding** journey `pt-onboarding-aireadiness-guided` on the `not_started` seat
documented above, v2.8 M256 iter-26).
The M250 fidelity gate was proven LIVE-GREEN both vantages (employee `aria-completed` + manager `dana-manager`,
Northwind, cold reset-to-seed, escapes=0) for parts 1/2/3/5 + the core part-4 sections. Code-of-record:
`rosetta-extensions` @ `july-jitter-m250-iter07`. **Re-proven LIVE on billion at M254 (gate (d)):** both
vantages, 8/8 AI-readiness sections + the 3 M250 drift-fixes confirmed cold on the consolidated platform.

## Cross-references

- **Authoritative in-repo deep-dive** (the platform's own KB): `app/knowledge/ai-readiness/overview.md` (start
  there; per-topic docs under `app/knowledge/ai-readiness/`) — the 2-axis/4-archetype model, the 3-step plan, the
  scoring engine, live-vs-frozen cycles, compare, CSV, Talk-to-Data. This corpus doc *summarizes* for the rosetta
  reader + the M51 seeder; that KB is the source of truth for deep work.
- Backend service: [`backend.md`](backend.md) (AI Readiness is the `app/internal/aireadiness/` package, split out of
  `internal/workforce/` at app v1.351.1 — see the ⚠️ package-refactor callout).
- The seeded demo world it plugs into: [`../ops/demo/stories-spec.md`](../ops/demo/stories-spec.md) (the Stories &
  Heroes model — M51 adds the AI-readiness showcase org as a 3rd story).
- Verified-skill chain the Step-1 signal reuses: [`../ops/demo/stories-spec.md`](../ops/demo/stories-spec.md)
  (`user_skill_evidences`).
- AI provider routing for the narrative LLM: [`../architecture/ai_architecture.md`](../architecture/ai_architecture.md).
