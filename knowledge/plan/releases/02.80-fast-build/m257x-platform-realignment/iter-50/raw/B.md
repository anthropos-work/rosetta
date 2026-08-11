# Seat B — M257x clause-5 KB-fidelity reading

## 1. Header

**Corpus under audit:** `/Users/marco/workspace/anthropos/rosetta`, branch `m257x/platform-realignment`,
HEAD `57dfbfded8791fcb12a4651d747247ce9d04d7f0` (confirmed via `git rev-parse HEAD`).

**Ground-truth clones consulted** (all under `stack-demo/`):

| clone | used for |
|---|---|
| `app` (`5ba17044`, `v1.363.2` per `CHANGELOG.md:5`) | aireadiness pkg, jobsim AI manager, validation engine, Ent schemas, credits/labs/payments, coursebuilder, cms/directus DTOs, clerk events |
| `app/studio` (`aeec036a`) | studio-room config slots, pricing table, bedrock/boto3 absence |
| `platform` (`2adcf714`) | `repos.yml`, `docker-compose.yml` (customerio-sync) |
| `next-web-app` (`bb3313bc`) | all AI-readiness FE anchors, Clerk pkgs, `setActive`, competency mapping |
| `sentinel` (`88bc5592`), `storage` (`4ce8ece5`), `messenger` (`fa47850d`) | authn/Clerk import absence |
| `studio-desk` (`14a5442a`) | `STUDIO_ACCESS_ROLES`, `AI_PROVIDER_CHAIN`, Clerk pkgs |
| `ant-academy` (`9c3843cd`) | Clerk pkgs, `proxy.js` org gate |
| `.agentspace/rosetta-extensions` | seeders, demopatch manifest, playthrough manifests, coverage-manifest |

**Positive control — `wc -l` on every assigned file** (single invocation:
`wc -l corpus/services/ai-readiness.md corpus/architecture/ai_architecture.md … corpus/architecture/README.md`):

| file | lines | briefed | match |
|---|---:|---:|:--:|
| `corpus/services/ai-readiness.md` | 634 | 634 | ✅ |
| `corpus/architecture/ai_architecture.md` | 302 | 302 | ✅ |
| `corpus/architecture/security_compliance.md` | 265 | 265 | ✅ |
| `corpus/services/ai-labs.md` | 156 | 156 | ✅ |
| `corpus/services/clerk-integration.md` | 128 | 128 | ✅ |
| `corpus/services/customerio-sync.md` | 75 | 75 | ✅ |
| `corpus/architecture/README.md` | 38 | 38 | ✅ |
| **total** | **1598** | **1598** | ✅ |

All seven read IN FULL, top-to-bottom, via `Read` with no offset/limit.

**Pipeline-integrity notes (briefing rule 3).** Two searches returned "no matches found: `--include=*.go`" /
`--include=*.ts` — that was **zsh globbing the flag**, not an empty result. Both were re-run with quoted globs
and produced hits. Every negative finding below was paired with a known-matching control in the same
invocation (recorded per-item). One grep hit a binary (`app/schema.svg`) and was re-scoped with `-I` +
`--include`.

---

## 2. BLOCKERS

**None. 0 blockers.**

I could not construct a single finding meeting the bar — "a claim a reader would ACT ON that is FALSE, or a
load-bearing `file:line` anchor that does not resolve to what the text says is there" — that also survived
verification. Every substantive claim I tested in these seven files verified TRUE against the clones, and in
two cases the corpus is **more accurate than a comment in the platform source** (see Audited zeros §4, items
CSV-15 and credits-refine).

The closest candidates were three line-number drifts inside one repaired blockquote in `ai-readiness.md`;
all three underlying claims verify TRUE, the symbols are unique and one `grep` away, so I graded them MINOR
per the briefing's own naming of "line drift" as the MINOR archetype. They are listed first under MINORS and
flagged as a cluster so the caller can re-grade if drift is scored harder this pass.

---

## 3. MINORS

**7 minors.** (Counts: 6 line-drift/range, 1 omitted effect.)

### Cluster A — three anchor drifts in one file, `apps/web/.../ai-readiness/AIReadinessClient.tsx` (655 lines)

| # | corpus `file:line` | cited anchor | actual location | claim itself |
|---|---|---|---|---|
| M1 | `corpus/services/ai-readiness.md:316` | `AIReadinessClient.tsx:69` `const SHOW_SECONDARY_TABS = false;` | **:78** `const SHOW_SECONDARY_TABS: boolean = false;` (:69 is an `import`) | **TRUE** — also used at :599 to strip the `compare` tab |
| M2 | `corpus/services/ai-readiness.md:449-450` | `:137-138` computes `effectiveCycleId = selectedCycle ?? activeCycle?.id ?? latestClosedCycle?.id` | **:153-154** (:137 is blank, :138 is `const cyclesQ = …`) | **TRUE**, verbatim |
| M3 | `corpus/services/ai-readiness.md:450-451` | gates the data GET on `cyclesQ.isFetched` (`:150-154`) | **:169** and **:176** (`enabled: featureOn && cyclesQ.isFetched`); :150-154 is the `effectiveCycleId` block | **TRUE** |

Note the drift is **non-uniform** (+9, +16, +19) and the neighbouring anchor `:133-134`
(`const { orgEnabled } = useAiReadinessEnabled(true);` / `const featureOn = …`) resolves **exactly**, so this
is not a single insertion offset. Worth a re-pin sweep of this one file rather than a blanket shift.

### Cluster B — over- or under-wide ranges (anchor lands on/near the right construct)

| # | corpus `file:line` | cited range | actual | note |
|---|---|---|---|---|
| M4 | `corpus/architecture/ai_architecture.md:232` | "the enum is `…/jobsimulation.go:983-990`" | AIModel enum is `const (` **:979** … **:991**; members :980-990 | cited range omits the first 3 members (`anthropic-35/37/4-sonnet-aws`). All three examples the sentence gives (`gpt-5` :987, `gpt-4.1` :985, `anthropic-45-sonnet-aws` :983) **are** inside the cited range |
| M5 | `ai_architecture.md:274` + `security_compliance.md:220` | `calculateSkillScore` (`v3/validator/skills.go:53-64`) | function is **:53-62**; :63 blank, :64 is `run`'s signature | trivial over-range; `:75` (the `passed/total*100` line) is **exact** |
| M6 | `corpus/services/ai-readiness.md:564-565` | `computeCycleTotals` (`how_we_measure.go:253-261`) | function signature at **:260**; :253 is the previous function's `return` | range starts inside the prior function. The companion anchor `:285-287` (`FROM public.interactions i JOIN public.job_simulation_sessions s`) is **exact** |

*(A 7th of this shape — `ai_architecture.md:23` citing `getClient` as `:259-289` when the function is
`:259-288` — is a one-line tail and is not counted separately.)*

### Cluster C — omitted effect

| # | corpus `file:line` | issue |
|---|---|---|
| M7 | `corpus/architecture/ai_architecture.md:211-212` | "all it does is swap the dispatched **endpoint** to `openai-hosted`". Measured at `internal/jobsimulation/calls/livekit.go:140-144`: the flag body sets **both** `agentName = "anthropos-agent"` (:142) **and** `agentEndpoint = "openai-hosted"` (:143). On the US branch (`:126-127` sets `anthropos-agent-us` / `azure-us`) the flag therefore also renames the agent. The residency conclusion the paragraph draws is unaffected and correct; the "all it does" is an undercount. Both cited ranges (`:131-135` read, `:140-144` effect) resolve **exactly**, and the enclosing function is `CreateAgentDispatch` (`:106`) as stated. |

---

## 4. Audited zeros — read in full, measured, found clean

### `corpus/architecture/README.md` (38 lines) — CLEAN
- Lists all **10** non-README files in `corpus/architecture/` (`ls` shows 11 entries incl. README). No omission.
- All 14 relative links resolve (link checker validated with a deliberate broken-link probe — it flagged).
- `../ops/platform-alignment.md` exists; `rosetta-extensions/stack-core/platform_alignment_guard.py` exists.
- **`:21` "`authn` is not a dependency of any service"** — VERIFIED: `grep -rn "anthropos-work/authn" --include="go.mod" --include="go.sum"` over `stack-demo/` = **0** (control: `grep "anthropos-work" app/go.mod` returns `ai v1.40.2`, `colony v0.35.2`, `proto v1.210.0`, `taxonomy v1.2.0` + others). The "four imported as private modules (ai/colony/proto/taxonomy)" set is exactly right for the five-library scope.

### `corpus/services/customerio-sync.md` (75 lines) — CLEAN
- Compose block matches `platform/docker-compose.yml:220-238` field-for-field (`context: git@…#main`, `ssh: ["default"]`, `VERSION: dev`, `GH_ACCESS_TOKEN: $GH_PAT`, `8080:8080`, the `DB_CONNECTION_BACKEND` DSN incl. `search_path=public`, `profiles: [customerio-sync, all]`, `depends_on postgresql: service_healthy`). The doc's snippet omits `env_file` and `networks` but is visibly abridged — not reported.
- "not cloned locally by `make init`" — VERIFIED: absent from `platform/repos.yml` (which lists app, cms, jobsimulation, sentinel, storage, messenger, roadrunner, next-web-app, studio-desk).

### `corpus/services/clerk-integration.md` (128 lines) — CLEAN
Every checkable claim verified:
- **`app/go.mod:31` = `github.com/clerk/clerk-sdk-go/v2 v2.7.0`** ✅ exact line, exact version.
- **`@clerk/nextjs ^6.39.2` aligned across four**: `apps/{web,hiring,integration}/package.json:{10,10,9}` + `ant-academy/code/package.json:52` ✅.
- **`@clerk/clerk-expo` misaligned**: `next-web-app/apps/mobile/package.json:6` = `~2.6.18`; `ant-academy/mobile/package.json:18` = `~2.19.36` ✅ both exact lines; "thirteen minor versions apart" ✅ (2.6→2.19).
- **"Webhooks (svix, 12 event types)"** ✅ — `internal/clerk/events/events.go` switch has exactly 12 cases (:121-190): `user.{created,updated,deleted}`, `organization.{created,deleted,updated}`, `organizationInvitation.{accepted,created,revoked}`, `organizationMembership.{created,deleted,updated}`. svix import at :27.
- **`metabase/route.ts` checks `orgRole === 'admin'`** ✅ — `apps/web/src/app/api/metabase/route.ts:35` `authData.orgRole !== 'admin'` (bare form, as claimed).
- **`STUDIO_ACCESS_ROLES` = `org:admin` **and** `content_creator`** ✅ — `studio-desk/src/index.ts:96` + `app/services/userService.ts:16`: `['admin','org:admin','content_creator','org:content_creator']`.
- **ant-academy `REQUIRE_ORGANIZATION_MEMBERSHIP` → `/no-organization`** ✅ — `ant-academy/code/proxy.js:91,131,318`.
- **"Sentinel … no Clerk/authn import"** ✅ — `grep -rn "colony/authn" sentinel --include="*.go"` = 0; control: `grep "anthropos-work/colony" sentinel` hits `main.go:7`, `cmd/root.go:12-13`. So Sentinel imports colony-the-framework but not `colony/authn` — the doc's row is precise.
- **"storage, messenger — no auth"** ✅ — 0 `colony/authn` imports in each.
- **"`AuthRole()` has zero call sites on the allow/deny path"** ✅ — `grep -rn "AuthRole" app --include="*.go"` = **0** repo-wide (control in same invocation: `OrgCheckUserPermission` hits `internal/testsupport/authz.go:15,24,25`).

### `corpus/services/ai-labs.md` (156 lines) — CLEAN
- `app` at **`v1.363.2`** ✅ (`CHANGELOG.md:5`), matching the doc's `:18`.
- Code map: `internal/{labs,credits,payments,subscriptions}/` + `stripe/` all exist; `internal/labs/{adapter,catalog,labsapi,session}` ✅.
- **Cost model** — `internal/credits/cost.go`: `creditCost` map (:86-91) = `course.build:5`, **`course.refine:1`**, `course.translate:1` ✅. *Note the doc is RIGHT where the source's own header comment (:29, "5 credits per refine turn") is STALE — the map carries `// D1 ruling: flat 1 credit per refine turn`.* `MarginMarkup = 1.40` (:133), `PricePerCreditUSD = 0.45` (:142), `DefaultSeedBalance int64 = 500` (:231), `purchasePackages` starter 50 / team 200 / scale 500 (:201-203) ✅ all exact.
- **`POST /credits/purchase` removed (Wave 13)** ✅ — stated at `internal/web/backend/credits/handler.go:12`; only `GET /credits/balance` + `/credits/transactions` remain (:3-4).
- **Stripe webhook handles `customer.created` + `customer.subscription.{created,updated,deleted}` only, NO `checkout.session.completed`** ✅ — `internal/web/backend/api/api.go:315` + `:322-324`; `grep -rn "checkout.session.completed" --include="*.go"` = **0** repo-wide.
- **`lab_sessions` deliberately no `PrimaryKeyMixin`, `id` a String** ✅ — `internal/data/ent/schema/lab_session.go:21-28` says so verbatim; `field.String("id")` at :36.
- **LabsTier `essential`/`professional`/`frontier`** ✅ (`labs_budget.go:26-28`); **5 `CapReason`s** `model_not_in_tier`/`per_session_cap`/`org_total`/`user_monthly`/`user_lifetime` ✅ (:247-251); **30 s reconciler** ✅ (`spend_reconciler.go:60` `Interval: 30 * time.Second`).
- **labs-api default `:7070`** ✅ (`labsapi/client.go:6`).
- **All 5 migrations exist** ✅ — `20260617203555_add_labs_catalog.sql`, `20260529072659_add_lab_session.sql`, `20260617120000_add_lab_session_model.sql`, `20260717151144.sql`, `20260626120000_add_org_subscriptions.sql`.
- **GraphQL** `schemas/{labs,billing,organizations}.graphqls` ✅; **cmds** `labsimport`/`labskey`/`labsdemo` ✅; **`/v1/labs` group behind `labs:write`** ✅ (`labs_admin.go:31`) + `GET /v1/labs/:slug/workspace.tar.gz` (:40).
- **"v6.0 shared purse is a knowledge-plan release, milestones M600–M607, all planned"** ✅ — `app/knowledge/plan/releases/06.00-shared-purse/` contains `m600…m607` (+`m601b`). *(Directory listing only; no file under any `knowledge/plan/` was read, and this is the `app` repo, not rosetta's.)*
- Test-file inventory (`cost_test.go`, `manager_test.go`, `privacy_backstop_test.go`, `reconcile_orphan_test.go`, `labs_budget_test.go`, `spend_reconciler_test.go`, `create_e2e_test.go`, `labsapi/client_test.go`, catalog `content_test.go`/`workspace_test.go`, `subscriptions_test.go`) — all present; "Payments has no Go unit test" consistent with the listing.

### `corpus/architecture/ai_architecture.md` (302 lines) — CLEAN apart from M4/M7 above
The whole **Provider Routing Strategy** retraction block verifies line-exact:
- vendor consts at **`internal/jobsimulation/ai/ai.go:30-33`** ✅ (`azure`/`openai`/`anthropic-aws`/`anthropic` — exactly four).
- `getClient` a plain `switch`, no ordering/probing ✅ (:259-288); **a failed flag lookup keeps the EU client** ✅ (`return client, nil` with `client := a.azureClientEu`, :264+:272).
- `anthropic-aws` and `anthropic` both → the same Bedrock client ✅ (:278-281), constructed `config.WithRegion("eu-west-1")` ✅ (**:85-88**, exactly as cited by `security_compliance.md`).
- **429-only fallback → `vendor = Openai`** ✅ — `throttled` set at :167/:326 inside `isThrottlingError` guards (**:129** exact), consumed at :152-154 (`if throttled { vendor = Openai }`) and :299-301.
- **Mistral nowhere in the AI path; OCR only** ✅ — `internal/cms/studio/markdownManager.go:19` and `studioManager.go:583` both exact; no other non-comment Mistral use.
- **`app/internal/ai/ai.go` "no such file"** ✅ — `ls` errors; wrapper anchors `jobsimulation/ai/ai.go:267,344` and `skillerai/ai.go:347` all land on `"flag_use_azure_us"` selection lines ✅.
- **The Directus-DTO correction (iter-48) is exactly right**: `AIVendor *AIVendor` at `cms/directus/collections/jobsimulation.go:905` ✅; nil→`simulation.Openai` at **:1302** with the `if` at :1303-1305 ✅; sequence built at **:1307** ✅; model default `simulation.Gpt5Point1` at **:1297** ✅. The AIVendor enum has **five** members with `Azureglobal` at **:971** ✅ and the runtime switch **four** cases at `simulator/ai/ai.go` **:58 / :69 / :86 / :102** ✅ — every one exact.
- **The three simulation model defaults** table: content `gpt-5.1` (:1297/:1302) ✅; runtime unmatched-model `gpt-4.1` at **:65-66 / :82-83 / :126-127** ✅ all exact; scoring hardcoded Azure `gpt-4.1` + summarize `gpt-4.1-mini` at **`simulator/ai/ai.go:20-26`** ✅ exact. Anthropic arm fallbacks **:98-99** (Claude 3.7 Bedrock) and **:110-111** (Claude 3.5) ✅ exact. `default:` → `internalAi.Openai` at **:114-115** ✅ exact.
- **Studio-Room slots**: `studio/configs/production_config.ini:26-36` ✅ — the cited range spans exactly the 5 stable + 5 experimental keys, and every cell of the table (FAST/STRICT `gpt-5-mini` none; EXECUTION `gpt-5.4` none; CREATIVE `gpt-5.4` low; REASONING `gpt-5.4` medium) matches. `development_config.ini:26-36` byte-identical ✅. Template `:39`/`:40` = `azure, gpt-4o, none` ✅ exact. **`gpt-5.2` in no studio config** ✅ (`grep` over `studio/configs/` exits 1; only `services/ai.py:356,508` pricing entries — both exact).
- **`grep -rin 'bedrock\|boto3' app/studio/` = 0** ✅ (control in same invocation: `anthropic` = 34 hits — pipeline sound). The "Studio-Room was never on Bedrock" correction holds.
- **Voice**: `voice_engine` 4-member enum at `jobsimulation.go:1079-1085` ✅ exact; nil→`gptrealtime` in `voiceEngineFromDirectus` at **:1594-1600** ✅ exact. ElevenLabs call/reply pipeline resolvers `getJobSimulationCallSignedUrl` / `getJobSimulationCallConversationToken` both exist (`resolver_jobsimulations.go:572,656`) ✅. LiveKit MP3 ✅ (`livekit.go:176` `EncodedFileType_MP3`).
- **Evaluation System** (shared with `security_compliance.md`): `checkerEngines` **stored and never read** ✅ — all references are declaration/assignment (`validator.go:43,60,595`; `criterion.go:34,72,90`), never an index or range. Dispatch is the hardcoded switch: `criterion.go:126` `switch critCheck.Engine`, `:127 case check.EngineLlm`, `:142 case check.EngineTextDiff`, `default: → error` ✅. `NewLLMBulkChecker(c.logger)` at **:428** ✅ exact; template `templates/checkValidationBulk.tmpl` exists and `:27` reads *"assess whether the `<asset>` shared by the user (the player) meets or does not meet"* ✅ **verbatim**, returning `{check_id, feedback, success}` ✅; **temperature 0.0** at `bulkChecks.go:118` ✅. `validateCodeDiff` at **:168**, run **concurrently** via `pool.New().WithMaxGoroutines(2)` (:162-173) ✅ — the "concurrently" is literally true, not loose. `:450-475` = the deterministic `cdiff.Diff == ""` comparison, no model ✅.
- **"No 60/65/75/85/95 threshold ladder"** ✅ — `grep -rInE '\b65\b.*\b75\b.*\b85\b.*\b95\b'` over `app`/`next-web-app`/`cms`/`jobsimulation` (Go/TS only) = 0; control in same invocation: `calculateCompetencyLevelScore` returns 3 hits. The real formula at `skills.go:40-51` matches **word for word**, including the `// TODO fix this formula` comment at :49.
- `convertLevelTo100` at **`app/internal/skill/skill.go:617-623`** ✅ exact; `CompetencyReadLevel.tsx:18` is the `getSkillScoreForSimulation(...)` call, whose body (`skillScore.ts`) is a plain `Math.round((levelsCount*score)/100)` ✅.
- Cost tracking in `internal/aiusage/ai_usage.go` fed by `Event_AiUsage` ✅ (directory present, consistent with `security_compliance.md:170`).
- Cross-file anchors all resolve: `external_services.md:541` = the Anthropic-Direct provider row ✅; `:545` = "There is **no ordered EU-first fallback chain.**" ✅; `:569` = "**Five** things can send a request outside the EU" ✅; `:577-587` = the `ai_vendor`-unset derivation ✅; heading `#routing-what-is-actually-implemented` at `:543` ✅; `shared_libraries.md#taxonomy-figures` anchor at `:196` ✅; `#ai` heading at `:98` ✅.
- **`app/internal/cms/studio` runs `studio/gen.py` as a subprocess** ✅ (`studioManager.go:119`), corroborating "runs as a subprocess inside the `app` container".
- `AI_PROVIDER_CHAIN` for studio-desk ✅ (`studio-desk/.env.example:57`).

### `corpus/architecture/security_compliance.md` (265 lines) — CLEAN
The Layer-1 Ent-schema census — the most heavily-repaired block in the file, and the one the doc itself warns
"has now been wrong FOUR times" — **reproduces exactly**, running the doc's own derivation verbatim:

| doc claim | measured | ✓ |
|---|---|:--:|
| 139 `.go` files in `internal/data/ent/schema/` | `ls *.go \| wc -l` = **139** | ✅ |
| 135 Ent schemas (4 declare none) | `grep -l 'ent.Schema' *.go \| wc -l` = **135** | ✅ |
| `OrganizationMixin{}` in 30 | **30** | ✅ |
| `OrganizationIDMixin{}` in 7 | **7** | ✅ |
| plain `organization_id`, neither mixin = 18 | doc's `comm -23 … \| xargs grep -l` = **18** | ✅ |
| subtract self-policing + owner-filtered → 16 | `org_membership.go` (own `Policy()`) + `academy_feedback.go` (`UserMixin{}`) are the only two → **16** | ✅ |
| the 16 named files | set-identical to my measured 18 minus those 2 | ✅ |
| the 7 `OrganizationIDMixin` users named (`category`, `jobrole`, `similarity`, `skill`, `specialization`, `studio_document`, `studio_task`) | **exactly those 7** | ✅ |
| "only FOUR files declare any `Policy()`" | `organization.go`, `mixin.go`, `user.go`, `org_membership.go` — **4** | ✅ |
| 31 auto-filter by org (30 + `Membership`) | consistent | ✅ |

Supporting anchors, all exact: `mixin.go:126` = `func (OrganizationMixin) Policy()` ✅; `mixin.go:98` =
`func (UserMixin) Policy()` whose body carries `rule.FilterOwnerRule()` ✅; `org_membership.go:172-188` =
`func (Membership) Policy()` ending in `privacy.AlwaysDenyRule()` ✅; `job_simulation_session.go:5` =
*"L2: NO Ent privacy Policy; owner/org/tenant are plain fields."* ✅ **verbatim**; `jobrole.go:18` and
`category.go:15` both = *"…no UserMixin/OrganizationMixin and no Policy(), so the taxonomy stays globally
readable"* ✅.

Also verified in this file:
- **Layer 3 org-switch = `clerk.setActive`, not re-auth** ✅ — `useOrgSelection.tsx:94`, `useResolveActiveOrg.tsx:107`, `useActivateMembershipOrg.tsx:81` — **all three exact**.
- **EU-residency bullets** ✅ — `ai.go:262-266` (Azure EU arm), `:85-88` (Bedrock `eu-west-1`), `:263-277` (the `flag_use_azure_us` swap — exactly the Azure case), `isThrottlingError` `:129`/`:166`/`:325` — **every one exact**.
- **"Anthropic Direct is not used at all is FALSE"** ✅ — `coursebuilder/bedrock.go:109-112` = `newUnderlyingClient` → `NewAnthropicClientWithModel` ✅ exact; `:100` = `return "anthropic-api"` ✅ exact; and **"every model call"** holds — all three model roles (`:229`, `:244`, `:270`) route through `newUnderlyingClient`.
- **EU AI Act blockquote** — same validation-engine evidence as above; all anchors exact. The corpus's careful "*most*, not all" hedging (with `EngineTextDiff` as the stated exception) is precisely what the code shows.
- The `:489` self-note ("a TypeScript codegen comment") is a *historical* note about a since-corrected anchor; `external_services.md:489` is indeed a TS snippet (`variables: { id: '123' }`). Not actionable.

### `corpus/services/ai-readiness.md` (634 lines) — CLEAN apart from M1/M2/M3/M6 above
This is the largest and most heavily repaired file; I verified it densely.

**Package refactor (the ⚠️ callout)** ✅ — `internal/aireadiness/` exists with `readiness.go`, `scoring.go`,
`csv.go`, `steps.go`, `narrative.go`, `how_we_measure.go` (with `computeInterviewInsightsV2` at :819),
`cycles.go`, `compare.go`, `diagnosis.go`, `provision.go`, `defaults.go`, `emailoverride/`, `emailpreview/`,
`notifications/`, `testdb_test.go`. `internal/workforce/ai_readiness.go`, `ai_readiness_v2.go`,
`readiness_steps.go` **all gone** ✅. `workforce/members.go` retains `LoadMembers` (:349) +
`LoadMembersByUserIDs` (:353) ✅. `WorkforceDirectory` interface at `aireadiness/manager.go:38` exposing both
✅ — so the doc's "the bounded swap is now expressible at that interface call site" is exactly right.

**D-07 demopatch re-anchor** ✅ — `app-aireadiness-snapshot-loadmembers.yaml` **`:42` reads
`path: internal/aireadiness/readiness.go`** ✅ exact, **`:33` reads "v2.7 M254 RE-POINT"** ✅ exact, and
`demo-stack/tests/test_aireadiness_snapshot_loadmembers_m254.py` exists ✅. The iter-46 correction (that this
paragraph had stated completed work as outstanding) is itself correct.

**Gates** ✅ — `enum/organization_settings.go:47` = `OrganizationSettingAIReadiness OrganizationSetting = "ai_readiness"` ✅ exact; `isAIReadinessEnabled` in `aireadiness/steps.go:541` reading `IsOrganizationSettingEnabled` ✅; `useAiReadinessActive.ts:22` = `const rawFlag = useFeatureFlagEnabled(AI_READINESS_FLAG);` ✅ exact; `aiReadiness.constants.ts:26` = `export const AI_READINESS_FLAG = 'flag_ai_readiness';` ✅ exact; `AIReadinessClient.tsx:133-134` (the manager gates on `orgEnabled` alone) ✅ exact. The `=== true` / sticky-flag code block quoted in the M219 correction matches `useAiReadinessActive.ts:22-25` in substance and order.

**Scoring** ✅ — `enum/ai_readiness.go:9-11` step types ✅; `scoring.go` `archetypeHighBand = 75` (:26) and `archetypeLowCeil = 50` (:30) ✅; `classifyArchetype` (:46-57) = Champion both ≥75 / Standby both ≤50 / `usage >= knowledge` → Explorer / else Hidden Talent ✅ **including the exact-tie→Explorer**; buckets `None 0-24 / Low 25-50 / Medium 51-74 / High 75-100` ✅ **verbatim from the source's own comment block (:63-66)**.

**Defaults / 31 skills** ✅ — `defaults.go` `defaultReadinessSkills` = **19 @ 1.0 + 12 @ 0.5 = 31**, counted entry by entry; denominator **19 + 6 = 25.0** ✅. `defaultReadinessSims` = **3** entries `{simulation,"tech"}`, `{simulation,"business"}`, `{interview,"both"}` (:76-78) ✅ exact. The doc's arithmetic checks out independently: `round(1.0/25*30)=1`, `round(0.5/25*30)=1` (the 1/30 floor), `round(9/25*30)=11`, `25/25*30=30` — all four as stated.

**The iter-49 `queryReadinessStarters` repair — the highest-risk passage in the file — is CORRECT** ✅.
`aireadiness/steps.go:915` is `func (m *Manager) queryReadinessStarters(...)` and its SQL (:921-925) is
`SELECT DISTINCT user_id FROM public.ai_readiness_user_step_progresses WHERE organization_id = $1 AND
user_id = ANY($2::uuid[]) AND status <> 'not_started'` — **verbatim as quoted**, reading the *progress* table
and **no** `user_skill_evidences`. The platform's self-statement *"This DB signal is the only real 'has
started' check"* sits at **:913**, inside the cited `:907-914` ✅. `keepStartedMembers` at `readiness.go:684`,
called at `:390` ✅. Both directions of the old error are as the doc describes.

**Live-vs-frozen routing** ✅ — `GetAIReadinessWithOptions` at **`readiness.go:289`** ✅ exact; route 1
(`opts.CycleID != nil` + `status == "closed"`) at **:290-297** ✅ exact; route 2 (no active cycle + latest
closed) at **:307-312** ✅ exact; fall-through `buildLiveResponse` at **:314** ✅ exact. `computeOrgBreakdowns`
at **:330** ✅ exact. `computeTier1` at **:139** ✅ exact. `buildResponseFromSnapshots` (:771) calls
**`m.workforce.LoadMembers(ctx, orgID, "")`** at :779 ✅ — the whole-org hydration the doc describes.

**Interview findings / `interview_aggregated_reports`** ✅ — `how_we_measure.go:1055` is the **only** read of
that table in the package, and `conversation_extractions` has **0** occurrences in `internal/aireadiness/`
✅ (the "blamed on the wrong table" correction is right). `holdsBackFromInsights` (:986) filters
`strings.Contains(strings.ToLower(in.Category), "risk")` at :994 ✅ — the "category string is load-bearing"
claim is literal. `usageDimensionsFromReports` (:1164) **omits** absent/non-numeric KPIs (:1170, :1181) ✅.
`resolveSessionAuthors` (:1093) joins `public.job_simulation_sessions s JOIN public.memberships m`
(:1103-1104) ✅. Ent schema columns `(organization_id, sim_id, report JSONB, session_count)` ✅
(`interview_aggregated_report.go:23-26`).
**KPI ids** ✅ — `usageDimSpecs` (:1139-1155) is exactly `avg_adoption / avg_transformation / avg_originality
/ avg_depth / avg_ownership`; the retired `avg_frequency` / `avg_breadth` / `avg_context_fit` return **0**
hits repo-wide. Part 5's "only `avg_depth` survived the rename" is consistent.

**Surfaces** ✅ — `AIReadinessContainer.tsx`, `AIReadinessIntro.tsx`, `AIReadinessView.tsx` **all absent**;
the legacy route dir `…/enterprise/workforce/ai-readiness` **does not exist** (404 ✅). Commit **`dae0fb2f7`**
exists, is titled *"…drop orphaned container"*, dated **2026-07-13** ✅, and `--numstat` gives
**103 + 220 + 330 = 653** deleted lines ✅ **exact**. `AI_READINESS_URL` at `urls.ts:52` ✅ exact; imported at
`useNavbarSections.tsx:4` ✅, built at `:398-400` ✅, gated at `:547` ✅ — and a repo-wide grep finds the
constant in **exactly those two** non-`node_modules` files ✅. `e2e/specs/web.ai-readiness.spec.ts` exists ✅.
`WorkforceNewClient.tsx:125-151` is the 5-tab list (growth/verification/talent/assignments/activity),
omitting readiness ✅ exact. `useWorkforceAIReadiness.ts:23-27` passes only `{ searchParams: { tag } }` —
**no `cycle` param at all** ✅ exact.

**The "Handled for you" tile correction** ✅ — `HowWeMeasureTab.tsx` is **1,989 lines** ✅ exact;
`grep -c interviewQuestions` = **0** ✅ exact; `:1879` = `{/* ===== C · Handled for you this cycle ===== */}`
✅ **verbatim**; the three rendered cells at **:1915 / :1921 / :1927** ✅ all exact; and the orphan type field
`interviewQuestions: number;` at **`apps/web/src/hooks/useAIReadiness.ts:274`** ✅ exact.

**Cycle-state contract** ✅ — `deriveMode` at `components/ai-readiness/useAIReadiness.ts:48-62` ✅ exact,
treating a null deadline as passed (:55-57); `AIReadinessHero.tsx:88` = `if (!air.deadline) return null;` ✅
**verbatim**. The "one active cycle per org (partial unique index)" is real:
`terraform/migrations/20260618160827.sql:18` = `CREATE UNIQUE INDEX … ON "ai_readiness_cycles"
("organization_id") WHERE ((status)::text = 'active'::text)` — so one closed + one active is indeed legal ✅.
`participants_filter` jsonb `{all, tags, roles}` ✅.

**CSV — the doc beats the source comment.** `csv.go:35` says *"writes the 19-column CSV"*, but
`AIReadinessCSVColumns` (:17-33) has **15** entries and the row literal (:93-109) has **15** values. The
corpus's "**now 15 columns** … the recommendation columns were dropped from the earlier 19" is **correct**
and matches the source's own header note at :14-16. UTF-8 BOM (:44) and formula-injection neutralization
(:112-113) ✅.

**Other platform subsystems** ✅ — `CloseDueAIReadinessCycles` (`cycles.go:554` + worker
`ai_readiness_autoclose.go:37`) ✅; `RefreshLiveSnapshots` (`live_snapshots.go:71`) ✅;
`recommendation_engine.go` + `recommendation_signals.go` ✅ with `academy.EmbeddingsManager` on the Manager
(:88) ✅; `emailoverride/` validating against `messenger/pkg/aireadinessemail` ✅; `CompareCycles`
(`compare.go:154`) → a **6-section** response (Topline/Archetypes/Transitions/TeamDelta/SkillCoverage/
ThemesShift, :35-40) with **both cycles required `closed`** (:163, :170) ✅. Narrative uses `openai.GPT5Mini`
(`narrative.go:39,85`) ✅ = the doc's "GPT-5-Mini". REST surface confirmed in `api/server.gen.go:84-114` +
router `:1768-1776`: `/workforce/ai-readiness`, `/cycles` GET+POST, `/cycles/{cycleID}` GET,
`/cycles/{cycleID}/close` POST, `/steps-completion`, `/narrative`, `/compare`, `/export.csv`, **`/setup`
GET+POST** ✅ — every endpoint the doc lists.

**Seeder / tooling side** ✅ — all named artifacts exist in `rosetta-extensions`:
`stack-seeding/seeders/{ai_readiness_interview_report,ai_readiness_evidence,cockpit,contentref}.go`,
`seat_append_test.go`, `ai_readiness_funnel_test.go` (containing
`TestAIReadinessFunnelSeeder_DayZeroHeroGetsNoSignals` at :230 ✅),
`stack-verify/e2e/lib/coverage-manifest.ts`, `LegacyReadinessPaths` (:135) + `ValidateCockpitManifest` (:161)
✅. **`aiReadinessStageFor`** (`ai_readiness_funnel.go:205-226`) maps exactly as documented: explicit
`not_started` → **0** checked **FIRST** (:208), manager → **0** (:210), struggling → **1** (:212),
`default:` → **3** (:214) ✅ — including the doc's point that the explicit check precedes the derived
default. `blueprint.AIReadinessNotStarted = "not_started"` (:316) ✅.
**"5 Playthroughs"** ✅ — `playthroughs/manifest/ai-readiness.yaml` declares 4
(`pt-aireadiness-member-{done,progress}`, `pt-aireadiness-manager-{dashboard,howwemeasure}`) plus
`pt-onboarding-aireadiness-guided` in `onboarding.yaml` = **5**; the seat `pt-ai-onboard` exists in
`seed/pt-world.seed.yaml:302` and `seed-worlds.yaml:35` ✅.

**Cross-reference** `app/knowledge/ai-readiness/overview.md` exists ✅.

### Link integrity across all seven files
64 relative links extracted; **0 broken**. The checker was validated with a deliberate `../nope/missing.md`
probe, which it correctly flagged (a check that silently skips would otherwise read as a pass).

---

## 5. Unverified — and why

These are claims I could **not** check with the material available. Per the briefing they are reported as
neither passed nor blocking.

1. **`ai_architecture.md:105,138,146` — DB row counts.** `skill_embeddings = 42,790`,
   `job_role_embeddings = 18,919`, `≥22,470` public job roles. Require a live/prod Postgres; no running stack
   and no DB access in this session. The doc dates them ("measured 2026-06-29") and explicitly frames 42,790
   as a floor, which is the honest shape.
2. **`ai_architecture.md:183-189` — the five LiveKit agent repos** (`livekit-agent`, `-chain`, `-azure-us`,
   `-azure-eu`, `-azure-eu-fr`) and the M257x iter-01 "azure-eu/-eu-fr dispatch nothing" measurement. `gh` is
   unavailable and none of these repos is cloned, so org membership cannot be enumerated. *Partial
   corroboration only:* `calls/livekit.go:120-127` does construct `anthropos-agent` / `anthropos-agent-us`
   and `azure-<location>` endpoints, consistent with the naming.
3. **`ai_architecture.md:151` — skiller migrations `20260417103036` / `20260417120309`.** The `skiller` repo
   is decommissioned and not cloned (absent from `repos.yml`). The *result* is corroborated: the dedicated
   embedding tables exist in `app/terraform/migrations/20260615130000_skiller_taxonomy.sql`, and
   `extensions.vector` appears in three app migrations.
4. **`security_compliance.md` §§ Network Security, Transport, Access Management, Backup & DR, Server &
   Runtime, Monitoring, GDPR, Sub-Processors** — VPC CIDR `10.0.0.0/16`, subnet tiering, TLS 1.3 cipher set,
   AES-256/KMS, CloudTrail, "every 6 hours → S3/Azure/Hetzner", "DR site: US AWS region", ECS 30-second health
   checks, 90-day log retention, **DPA v1.4**, **18 approved sub-processors**, 90-day auto-deletion, "CV data
   never used for AI training". These are production-infrastructure and legal/contractual facts. **The
   `stack-demo/platform` clone contains no Terraform at all** — `find . -name "*.tf"` = **0** (the clone holds
   only `Makefile`, `docker-compose.yml`, `common.yml`, `repos.yml`, `postgresql/`, docs), so the briefing's
   description of that clone as including "terraform" does not hold for this checkout. `app/terraform/`
   exists but contains only `migrations/`. Nothing here is contradicted; none of it is checkable.
5. **`clerk-integration.md` §1/§2 — `colony/authn` internals.** JWKS verification via `clerk-sdk-go/v2`
   `jwt.Verify` + `jwks.Client`, the **1-minute leeway**, `jwt.Decode`, the 401-on-invalid middleware, the
   claim→getter mapping table, and "`AuthRole()` is never enforced inside `authn`". `colony` is a private Go
   module and is not cloned. *Consumer-side corroboration:* `AuthRole()` has **0** call sites in `app`, which
   is consistent with (but not proof of) the non-enforcement claim.
6. **`clerk-integration.md` — Clerk dashboard state.** "custom claims baked into the default token via
   dashboard config", the whole **Not used** feature list (MFA/TOTP/passkeys, SAML/SSO, Clerk Billing, Actor
   Tokens, Waitlist, `@clerk/themes`, …), and "8 locales via `@clerk/localizations`". These describe SaaS
   tenant configuration, not repo content. The *negative* half is partly corroborable — e.g. `@clerk/themes`
   appears in no `package.json` I enumerated — but dashboard toggles are not observable from here.
7. **`ai-labs.md:5-7` — the existence of `anthropos-work/AI-Labs` as a repo.** `gh` unavailable and the repo
   is not cloned. **Strongly corroborated in-code**, however: `internal/labs/labsapi/client.go:14` cites
   `AI-Labs/docs/PHASE_B_LABSESSION.md`, `catalog/content_import.go:14` refers to "the SAME shape the AI-Labs
   runtime enforces", and `anthroposlabs.com` URLs appear in `labsapi/client_test.go:43-44`. I did not
   escalate this as verified because the repo itself was never observed.
8. **`ai-readiness.md` — every live runtime measurement.** The 2.09 s live recompute, the 24 ms frozen read,
   the 180 s iter-08 timeout, 180 s → 19 ms from the demo-patch, `78.4%` / ≈156 of 199 frozen snapshots, the
   `/home` "AI Readiness → NO" probe on `billion`, and the M250/M254 LIVE-GREEN gate results. No stack is
   running; these are reproducible only against a live demo. The doc itself says *"Re-measure before relying
   on either number; do not re-derive them from prose"* — which is the right instruction.
9. **`customerio-sync.md`** — the service's internal protocol/field mapping and "Port 8080 *likely* serves a
   health/metrics endpoint". The repo is intentionally never cloned (built from the GitHub URL), so this is
   unverifiable by construction; the doc already hedges with "likely" and points the reader at `gh repo
   clone`.
10. **`ai_architecture.md:225` — "Both recordings are stored in S3".** Partly checkable and **not
    contradicted**: LiveKit writes MP3 (`livekit.go:176`) and Chime has an S3 `SinkArn`
    (`internal/jobsimulation/recording/chime.go:189`). The delivery-side Bunny.net CDN layer that
    `media-substrate-spec.md` describes (`BunnyVideoResource` resolvers,
    `resolver_jobsimulations.go:1519`) sits *above* capture and does not falsify the sentence, but I could
    not confirm bucket-level retention.

**One incidental observation, offered as context rather than a finding:** `how_we_measure.go:1106` still
issues `LEFT JOIN skiller.job_roles jr` — a live query against the legacy `skiller` schema that the corpus
elsewhere calls a decommissioned husk. No claim in *my* seven files asserts otherwise (they scope the
"legacy skiller schema" statement to the *embedding* tables, which did move to `public`), so this is not a
defect against them — but it may be worth a look from whichever seat owns `backend.md` / `skiller.md`.
