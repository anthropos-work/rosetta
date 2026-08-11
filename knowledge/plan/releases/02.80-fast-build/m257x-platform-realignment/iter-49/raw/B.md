# Seat B — M257x iter-49 KB-fidelity audit (9th clause-5 reading)

## Shas consulted

| Repo | sha |
|---|---|
| `rosetta` (branch `m257x/platform-realignment`) | `2fc633a2c5c09a6034e5ab4e29d509dfcadcbd8a` |
| `stack-demo/app` | `5ba1704482cf812b130c2d3673afd09f4f7f22e5` (`v1.363.2`) |
| `stack-demo/platform` | `2adcf714bd877a205e8948f59a23db49b884c054` |
| `stack-demo/next-web-app` | `bb3313bc0133ee5728ce83fda485e95bfea1a6c6` |
| `stack-demo/sentinel` | `88bc55929dde7ba43913966ec3fc36372e4ff32a` |
| `stack-demo/storage` | `4ce8ece52adb7c095e792e235da4a8913214d190` |
| `stack-demo/messenger` | `fa47850d9c507d1928da7a38f7b37bac1bb8fabc` |
| `.agentspace/rosetta-extensions` | `4d03b53a5e524e9abb020c1a4534ec968c25072b` |

## Coverage

| # | file | `wc -l` | lines read |
|---|---|---|---|
| 1 | `corpus/services/ai-readiness.md` | 634 | all 634 |
| 2 | `corpus/architecture/ai_architecture.md` | 302 | all 302 |
| 3 | `corpus/architecture/alignment_testing.md` | 521 | all 521 |
| 4 | `corpus/services/askengine.md` | 121 | all 121 |
| 5 | `corpus/services/ai-labs.md` | 156 | all 156 |
| | **total** | **1734** | **1734** |

Every file was read top-to-bottom in a single `Read` call with no `offset`/`limit`, so the whole file was in
context. Verification greps were then run against the ground-truth clones above.

---

## BLOCKERS

| # | site (`file:line`) | the false claim | what is true (platform `file:line`) |
|---|---|---|---|
| **B-1** | `corpus/services/ai-readiness.md:406-419` (the bullet **"Active cycle → the dashboard RECOMPUTES from signals"**, the passage this iteration's repair rewrote) | The active-cycle population filter is `keepStartedMembers` / `queryReadinessStarters`, i.e. `SELECT DISTINCT user_id FROM public.ai_readiness_user_step_progresses … AND status <> 'not_started'` — and therefore *"one with an `in_progress` row and zero evidences is **kept**"*. | On the **active-cycle** path `buildLiveResponse` does **not** call `keepStartedMembers` at all. `app/internal/aireadiness/readiness.go:387-392`:<br>`if cyc != nil { members, breakdowns, err = m.keepInCycleStep1(ctx, orgID, cyc.StartDate, …) } else { … m.keepStartedMembers(…) }`.<br>`keepStartedMembers`/`queryReadinessStarters` (`steps.go:915-939`) is the **no-active-cycle** branch only.<br>The active-cycle filter is `keepInCycleStep1` → `queryInCycleStep1Completers` (`readiness.go:638-660`), which requires **all three** of: `StepTypeEQ(enum.StepSkillMapping)` **AND** `StatusEQ(enum.StepCompleted)` **AND** `CompletedAtGTE(cycle.StartDate)`.<br>So on an active cycle: an `in_progress` row is **dropped** (not "kept"); a `completed` row on step 2 or 3 alone is **dropped**; and a `skill_mapping` row whose `completed_at` predates the active cycle's `start_date` is **dropped**. |

### Why B-1 is actionable, not cosmetic

This document is explicitly the contract the AI-readiness demo seeder builds against, and § *The CYCLE-STATE
contract* (`:508-509`) establishes that with an active cycle present the dashboard **does** take
`buildLiveResponse` — so the wrong branch is precisely the branch a seeder author will be on.

A seeder written to this paragraph writes `ai_readiness_user_step_progresses` rows with any status past
`not_started` and any `completed_at`, then expects the 199-member active-cycle dashboard to populate. Under
`queryInCycleStep1Completers` the entire population is filtered out unless each member has a **`skill_mapping`**
row at **`completed`** with **`completed_at >= cycle.start_date`**. That is the exact failure shape — an empty
but error-free manager dashboard — that this file's own § *The FILLED-ness contract* exists to prevent.

The correction *did* fix the original error it targeted (the old text credited `user_skill_evidences` as the
"has started" signal, which is refuted by `steps.go:907-914`'s own comment). It replaced it with the right
mechanism for the **wrong branch**. This is the "a correction introduced a NEW inconsistency in the same file"
case.

**Suggested repair shape:** split the bullet in two — *active cycle* → `keepInCycleStep1` (skill_mapping ·
completed · in-window), *no active cycle* → `keepStartedMembers` / `queryReadinessStarters` (any step ·
`status <> 'not_started'` · unbounded). Both facts are already true; they are attached to the wrong headings.

---

## MINORS

Anchor drift, path ambiguity, and over-narrow phrasing. None would mislead a reader about behaviour.

1. **`ai-readiness.md:316`** — cites `AIReadinessClient.tsx:69` for `const SHOW_SECONDARY_TABS = false;`.
   Actual: **`:78`**, and the declaration is `const SHOW_SECONDARY_TABS: boolean = false;`. The behavioural
   claim is correct — the Compare tab is stripped by the `.filter()` at `:595-600` (`… || (SHOW_SECONDARY_TABS
   && tab.key === 'compare')`).
2. **`ai-readiness.md:448-451`** — cites `AIReadinessClient.tsx:137-138` for `effectiveCycleId = selectedCycle
   ?? activeCycle?.id ?? latestClosedCycle?.id` (actual **`:153-154`**) and `:150-154` for the
   `cyclesQ.isFetched` gate (actual **`:169`** and **`:176`**, `enabled: featureOn && cyclesQ.isFetched`).
   Both claims are substantively correct.
3. **`ai-readiness.md:411`** — `queryReadinessStarters` given as `steps.go:915-938`; the function runs
   **`:915-939`**. The quoted SQL is verbatim-correct.
4. **`ai-readiness.md:565`** — `computeCycleTotals` given as `how_we_measure.go:253-261`; the func signature is
   at **`:260`** (doc comment `:256-259`). The companion anchor `:285-287` for `FROM public.interactions i JOIN
   public.job_simulation_sessions s` is **exact**.
5. **`ai_architecture.md:232`** — the CMS `ai_model` enum given as
   `internal/cms/directus/collections/jobsimulation.go:983-990`; the const block is **`:979-991`**. The cited
   range omits `anthropic-35-sonnet-aws` (`:980`), `anthropic-37-sonnet-aws` (`:981`),
   `anthropic-4-sonnet-aws` (`:982`) and `gpt-5.1` (`:990` is in range but `)` is `:991`).
6. **`ai_architecture.md:263`** — `calculateSkillScore` given as `v3/validator/skills.go:53-64`; the function is
   **`:53-62`** (`:64` is `func (s skillsValidator) run`). The `:75` anchor for `passed / total * 100` and the
   `:40-51` anchor for `calculateCompetencyLevelScore` are **exact**.
7. **`ai_architecture.md:264-272`** — `basevalidator/criterion.go:127` etc. is written unprefixed immediately
   after a fully-qualified `internal/jobsimulation/simulator/validation/v3/validator/validator.go` path. The
   real path is `internal/jobsimulation/simulator/validation/basevalidator/criterion.go` — **not** under
   `v3/validator/`. A reader who concatenates gets a non-existent file. (All four line anchors — `:127`,
   `:168`, `:428`, `:450-475` — resolve exactly at the correct path.)
8. **`ai_architecture.md:208-212`** — "all it does is swap the dispatched **endpoint** to `openai-hosted`". The
   flag branch (`calls/livekit.go:140-144`) also resets `agentName = "anthropos-agent"` (`:142`), which
   overrides a location-derived `anthropos-agent-us`. Incomplete, not false; the residency point stands. All
   other anchors in this ⚠ block verify exactly — the read is `:131-135`, the effect `:140-144`, and it is
   indeed inside `CreateAgentDispatch` (`:106`).
9. **`ai_architecture.md:128-129`** — "**`gpt-5.2` appears in no studio config at all** (only as a pricing entry
   in `studio/services/ai.py:356,508`)". The *config* half is verified (no `.ini` under `studio/configs/`
   contains it). But "only as a pricing entry" over-narrows: it also appears in `studio/changelog.md:106`,
   `studio/knowledge/operations/running-tests.md:32` and `studio/knowledge/concepts/tests-system.md:84` — the
   latter two asserting `configs/local_config.ini` uses it, which the configs refute. Worth a parenthetical,
   since it is a live platform-KB contradiction a reader will hit.
10. **`alignment_testing.md:193`** — "`gate.sh:61` calls `alignctl dna coverage --dna … --if-declared`". The
    call is at **`clerkenstein/alignment/scripts/gate.sh:69`**; `:61` sits inside the explanatory comment
    block above it. The claim itself is exactly right.
11. **`alignment_testing.md:511-519`** (§ Layout) — omits **`internal/canon`** (which exists alongside
    `dna`/`outcome`/`compare`/`report`), and lists `cmd/alignctl  run | capture | dna list|diff|validate`
    **without `coverage`** — contradicting this same document's `alignctl` reference at `:245` and the whole
    M218 section, where `dna coverage` is a real, gate-binding subcommand. Self-inconsistency inside one file.
12. **Informational, not a corpus defect** — `ai-readiness.md:270` ("CSV export is now 15 columns") is
    **correct** (`AIReadinessCSVColumns`, `internal/aireadiness/csv.go:17-32`, 15 entries), but the platform's
    own doc comment at `csv.go:34` still says *"writes the 19-column CSV"*. A reader grepping the source will
    find an apparent contradiction; a one-clause note would pre-empt it.

---

## What was checked hardest, and came back CLEAN

Recorded so the audited-zero areas are visible rather than merely unreported.

**`ai-readiness.md` — the M247 refactor + scoring surfaces (the other edited area).** All verified exact:
`internal/aireadiness/` package exists with every renamed file in the ⚠ table; `archetypeHighBand = 75` /
`archetypeLowCeil = 50` and the Champion/Standby/Explorer/Hidden-Talent classification incl. the exact-tie →
Explorer rule (`scoring.go:22-55`); the band table **None 0-24 / Low 25-50 / Medium 51-74 / High 75-100**
(`scoring.go:59-77`); `knowledge = ((step1+step2)/70)×100` (`scoring.go:92-94`) and usage = step3 scaled
(`:99-105`); step maxima 30/40/30 (`readiness.go:29-31`); **31 default skills = 19 core @1.0 + 12 enabling @0.5
→ denominator 25.0** (`defaults.go`, counted entry-by-entry); **3 default sims**, two track-keyed + one shared
interview (`defaults.go`); `resolveUserTrack` business-wins (`cycles.go:150-171`); the **partial unique index**
`organization_id WHERE status='active'` (`schema/ai_readiness_cycle.go:125-129`); `computeTier1`'s
whole-repertoire denominator + the platform's own *"a larger configured set makes a full score harder to
reach"* comment (`readiness.go:139-176`); `queryUserAISkills` selecting exactly `user_id, skill_id,
is_verified` (`readiness.go:82-90`); `GetAIReadinessWithOptions`'s **two** routes into
`buildResponseFromSnapshots` at `:289`/`:291-297`/`:309-312`/`:314`; `computeOrgBreakdowns` at `:330`;
`CompareCycles` requiring **both** cycles `closed` (`compare.go:163,170`); `CloseDueAIReadinessCycles`
(`cycles.go:554`); `queryActiveCycleEndDate` (`steps.go:316`); the `aireadiness.Manager` ctor arg list
(`manager.go:95-106`); the `WorkforceDirectory` interface with implementations left in
`workforce/members.go:349,353`; `OrganizationSettingAIReadiness` at
`enum/organization_settings.go:47` **exactly**; the `interactions.action_type` enum being **exactly
`{email, call}`** (`ent/interaction/interaction.go:92-96`); `usageDimSpecs`'s **five** current KPI ids
(`how_we_measure.go:1138-1158`) and `usageDimensionsFromReports` omitting absent/non-numeric KPIs
(`:1164-1197`); `interview_aggregated_reports` being the **only** table `computeInterviewInsightsV2` reads
(`:813-830`, `:1055`); the narrative LLM being **GPT-5-Mini** (`narrative.go:39,85`).

**next-web-app side.** `useAiReadinessActive.ts:22` **exact**; `aiReadiness.constants.ts:26` **exact**; the
flag's sole consumer confirmed by repo-wide grep (4 hits, all in those two files); `AIReadinessClient.tsx`
`grep -c posthog` = **0**; `:133-134` `const { orgEnabled } = useAiReadinessEnabled(true)` **exact**;
`urls.ts:52` **exact**; `useNavbarSections.tsx:4` and `:400` **exact**;
`useWorkforceAIReadiness.ts:23-27` has no `cycle` param and never calls `/cycles` — **confirmed**;
`AIReadinessHero.tsx:88` `if (!air.deadline) return null;` **exact**; `deriveMode` null-deadline →
`deadlinePassed = true` (`components/ai-readiness/useAIReadiness.ts:48-61`); the legacy route directory
`enterprise/workforce/ai-readiness/` **does not exist** and all three orphan components are **gone** (`find`
returns nothing) — the 404 claim holds; `HowWeMeasureTab.tsx` is **1,989 lines**, `grep -c
interviewQuestions` = **0**, and `:1879` / `:1915` / `:1921` / `:1927` are all **exact**;
`hooks/useAIReadiness.ts:274 interviewQuestions: number;` **exact**.

**`ai_architecture.md` routing (the other edited area) — clean.** Vendor consts `:30-33` **exact**; `getClient`
`:259-288` **exact**; both `anthropic-aws` and `anthropic` → the same `anthropicClient`, built with
`config.WithRegion("eu-west-1")` (`:87`); `isThrottlingError` → `vendor = Openai` the only automatic fallback
(`:154`, `:166`, `:300`, `:325`); the cited wrapper anchors `jobsimulation/ai/ai.go:267,344` and
`skillerai/ai.go:347` **all exact**; `internal/ai/ai.go` confirmed **absent**; Mistral confined to
`cms/studio/markdownManager.go:19` + `studioManager.go:583`; `coursebuilder/bedrock.go` key-set → first-party
Anthropic (`:109-114`), and it never touches `AIManager`; `grep -rin 'bedrock\|boto3' studio/` → **0 hits**
(command ran, empty output), so the Studio-Room retraction holds. The **five** Directus `AIVendor` members vs
the **four** switch cases, with `azureglobal` (`:971`) as the case-less one; `AIVendor *AIVendor` at `:905`;
nil→`simulation.Openai` at `:1302-1305` **before** the sequence is built at `:1307`; `simulator/ai/ai.go`
`:58-59`, `:65-66`, `:82-83`, `:98-99`, `:110-111`, `:114-115`, `:126-127`, `:20-26` — **every one exact**.
Voice: 4-member `SimulationVoiceEngine` enum at `:1079-1086`, nil→`gptrealtime` at `:1594-1596`, ElevenLabs
EU/US clients live (`calls/elevenlabs.go:28-38`), both call resolvers present. Evaluation: `checkerEngines`
stored at `validator.go:43,60-61` and passed at `:595` but **never consulted for dispatch** — `criterion.go:428`
hardcodes `NewLLMBulkChecker`; `EngineTextDiff` at `criterion.go:142`; temperature `0.0` +
`checkValidationBulk.tmpl` + `{check_id, success, feedback}` (`bulkChecks.go:56-58,85,118`);
`convertLevelTo100` at `skill/skill.go:617`. Studio slots: `production_config.ini:26-36` is **exactly** the
published table, stable == experimental, and `development_config.ini:26-36` **diffs clean** against it;
`config_template.ini:39,40` carry `gpt-4o` as claimed.

**`askengine.md` — clean, every checkable number exact.** `DefaultModelID = "eu.anthropic.claude-sonnet-4-6"`,
`DefaultRegion = "eu-west-1"`, `ASK_MODEL_ID`/`AWS_REGION` overrides, `Temperature: 0`,
`stripAnthropicAuthMiddleware` (`bedrock.go:25,26,162,163,185,252,315`); `askEngineMaxConns = 6` and
`COPILOT_DB_CONN` (`main.go:149,312,335`); `MaxInlineRows = 200`, `MaxCellLength = 400`
(`executor.go:18,21`); `maxAgenticIterations = 15`, `loopTimeout = 10 * time.Minute`,
`context.WithoutCancel` + `streamRegistry` (`handler.go:34,41,57,248`); the `TableRegistry` at **60** entries
(57 `public` + 3 `directus`) with the dotted `jobsimulation.*` names documented in-code as transition aliases
resolving to public (`registry.go:188-189`). **Test counts verified by `grep -c '^func Test'`: sandbox 49,
executor 13, prompt 10, followups 14 — all four exact.**

**`ai-labs.md` — clean, including the two most falsifiable numbers.** `app` really is **`v1.363.2`** at
`5ba17044` (`git tag | sort -V | tail`). The `creditCost` map is `build:5, refine:1, translate:1`
(`credits/cost.go:86-90`) — matching the doc, **against** the platform's own stale header comment at `cost.go:29`
which says *"5 credits per refine turn"*; the doc read the map, not the comment. `purchasePackages` 50/200/500
(`:200-204`); `PricePerCreditUSD = 0.45` (`:142`) derived via `MarginMarkup = 1.40` (`:133`).
`POST /credits/purchase` removal confirmed (`web/backend/credits/handler_test.go:128`). The Stripe webhook
handles `customer.created` + `customer.subscription.{created,updated,deleted}` and **no
`checkout.session.completed`** (`web/backend/api/api.go:315-323`, repo-wide grep). `stripe-go/v74 v74.30.0`
(`go.mod:62`); `LabsTier` essential/professional/frontier (`labs_budget.go:26-28`); the reconciler's
`Interval: 30 * time.Second` (`spend_reconciler.go:60`); labs-api default `:7070` (`labsapi/client.go:6`);
`labs.graphqls` carrying every named query/mutation.

**`alignment_testing.md` — clean; every enumerable claim measured.** `alignment/` holds exactly
`README.md cmd examples go.mod internal` — **no `scripts/`**, as claimed; module `anthropos.dev/alignment`;
`clerkenstein/alignment/scripts/{gate,drift-check}.sh` exist mirror-side; five DNAs and five runners
(`clerkenstein/alignment/cmd/{clerkrun,jsfapirun,expressrun,deployrun,multirun}`). **Every gene/capability
count in the scores table verified by parsing the DNA JSON**: `clerk-2.6.0` 14 caps / 27 genes;
`clerk-js-5` 6/9; `clerk-multi-1` 5/9; `clerk-deploy-1` 3/7; `clerk-express-1` 5/13 — and **only
`clerk-2.6.0` declares a `consumed_surface`** (15 entries; the other four declare none), exactly as the doc's
carefully-hedged §"what it does NOT guarantee" states. `MembershipOrgIdentity/real-org-eid` is present and is
a **`standard`** gene. The exit-code split is real: `ExitRegressed = 2` / `ExitUnmeasurable = 3` with the
banner (`cmd/alignctl/run.go:134-135,142`). Both zero-critical guards exist: `Validate` rejects a DNA with no
critical capability (`internal/dna/dna.go:276-280` + `validate_test.go:50-73`) and `GateMet` refuses a
non-zero critical threshold at `CriticalGenes == 0` (`internal/compare/compare.go:56-60`). `idPattern`
`^[A-Za-z0-9][A-Za-z0-9_-]*$` (`dna.go:21`) and `maxWeight = 1_000_000` (`:24,269-270`). The workflow file is
**git-tracked** at `clerkenstein/.github/workflows/alignment.yml` and its lines **`:10-11`** carry the
self-describing "currently inert" note **verbatim** as quoted. The toy arithmetic reconciles: `Add` critical
(3 genes @ w3) + `Greet` standard (3 @ w2) ⇒ 13/15 = **86.7% overall**, **100% critical**, **5/6 genes**.

**`ai-readiness.md` demo/seeder cross-claims — clean.** The M254 demopatch re-anchor is real:
`demo-stack/patches/app-aireadiness-snapshot-loadmembers/app-aireadiness-snapshot-loadmembers.yaml` carries
`# v2.7 M254 RE-POINT` at **`:33`** and `path: internal/aireadiness/readiness.go` at **`:42`**, with
`demo-stack/tests/test_aireadiness_snapshot_loadmembers_m254.py` present — so the ⚠ callout at `:37-47`
(itself a fenced self-correction from iter-46) is accurate and correctly *not* bookable. `cockpit.go`
`LegacyReadinessPaths` / `ValidateCockpitManifest` exist (`:122-135`); the seeders
`ai_readiness_interview_report.go`, `ai_readiness_evidence.go`, `ai_readiness_sim_skills.go`,
`seat_append_test.go` and
`ai_readiness_funnel_test.go:TestAIReadinessFunnelSeeder_DayZeroHeroGetsNoSignals` (`:230`) all exist, and
`Persona.AIReadiness = "not_started"` is read in `aiReadinessStageFor` (`ai_readiness_funnel.go:205`).

### Commands that failed (recorded so no empty output reads as absence)

- `grep -rn "ActionType" --include=*.go …` — zsh glob-expanded `--include=*.go` and errored. Re-run quoted as
  `--include='*.go'`; the enum was then located via the generated validator.
- `grep -rn "action_type" internal/data/migrations/*.sql` — no such directory (`app` uses Atlas elsewhere).
  The enum was confirmed from `ent/interaction/interaction.go:92-96` instead.
- `find / -type d -path "*anthropos-work/proto*/…/interactions"` — 0 hits (the private module is not in the
  local mod cache). Superseded by the generated-validator evidence above; the `{email, call}` claim is
  confirmed, not assumed.
