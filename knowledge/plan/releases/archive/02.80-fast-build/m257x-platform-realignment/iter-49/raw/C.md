# Seat C — M257x iter-49 KB-fidelity audit (ninth clause-5 reading)

## Header — shas consulted

| Repo | sha | Notes |
|---|---|---|
| `rosetta` (this corpus) | `2fc633a2c5c09a6034e5ab4e29d509dfcadcbd8a` | branch `m257x/platform-realignment` |
| `stack-demo/app` | `5ba17044` | v1.363.2 — the Go monolith |
| `stack-demo/platform` | `2adcf71` | compose, `repos.yml`, Makefile, `common.yml` |
| `stack-demo/sentinel` | `88bc559` | **not `60c229f`** — the brief's sha was wrong for this repo |
| `stack-demo/storage` | `4ce8ece` | **not `60c229f`** |
| `stack-demo/messenger` | `fa47850` | **not `60c229f`** |
| `stack-demo/graphql-wundergraph` | `60c229f` | this is the repo `60c229f` actually belongs to |
| `stack-demo/cms` | `ca50c81` | husk repo (go.mod pins) |
| `stack-demo/jobsimulation` | `462343b0` | husk repo (go.mod pins) |
| `stack-demo/roadrunner` | `87d8d44` | husk repo (go.mod pins) |
| `stack-demo/next-web-app` | `bb3313bc0` | v2.133.0 |
| `stack-demo/studio-desk` | `14a5442` | v0.152.4 |
| `stack-demo/ant-academy` | `9c3843cd` | v2.34.2 |
| `.agentspace/rosetta-extensions` | `4d03b53` | |
| Go module cache | `~/go/pkg/mod/github.com/anthropos-work/ai@v1.40.1` | the ONLY shared lib present locally |

**Tooling gaps declared up front (not "0 hits"):**
- `gh` is **not installed** (`command not found`). Every GitHub-archive-status claim
  (skiller ARCHIVED 2026-07-01; jobsimulation/skillpath ARCHIVED 2026-07-31;
  graphql-wundergraph ARCHIVED 2026-07-30; cms "frozen, not archived"; chronos "NOT
  archived, last push 2026-04-23") is **UNVERIFIED**, not refuted.
- `colony`, `proto`, `taxonomy` and `chronos` are **not cloned and not in the module cache**.
  `shared_libraries.md`'s colony package inventory, the "12 Connect-RPC services" proto
  count, the "131-line `node.go`" taxonomy claim and the chronos `colony v0.30.1` pin are
  **UNVERIFIED**. The `ai` library WAS checkable (cache @ v1.40.1, one patch behind the
  documented v1.40.2) and every claim in that section verified.

---

## Coverage

| # | file | `wc -l` | lines read |
|---|---|---|---|
| 1 | `corpus/architecture/architecture_overview.md` | 349 | 349 (all) |
| 2 | `corpus/architecture/security_compliance.md` | 265 | 265 (all) |
| 3 | `corpus/architecture/service_taxonomy.md` | 440 | 440 (all) |
| 4 | `corpus/architecture/shared_libraries.md` | 242 | 242 (all) |
| 5 | `corpus/architecture/README.md` | 38 | 38 (all) |
| 6 | `corpus/services/README.md` | 79 | 79 (all) |
| | **total** | **1413** | **1413** |

Each file was read in a single `Read` call returning the complete file (all well under the
2000-line cap), then re-walked passage-by-passage against source.

---

## BLOCKERS

| # | site (file:line) | the false claim | what is true (with platform file:line) |
|---|---|---|---|
| C-1 | `security_compliance.md:76`, `security_compliance.md:120`, `architecture_overview.md:299-300` | "**31 schemas auto-filter by ORGANIZATION** (the 30 mixin users + `Membership`)" / "Ent privacy policies auto-filter by organization on **only 31 of 135 schemas**" | The count is **32**. The `Organization` schema declares its OWN `Policy()` (`app/internal/data/ent/schema/organization.go:56`) whose **Query** policy is `rule.DenyIfNoOrganizationInContext()` + **`rule.FilterSameOrganizations()`** + `AlwaysAllowRule()` (`organization.go:94-97`), and `FilterSameOrganizations()` pins the query to `organization.ID(org.ID())` (`app/internal/data/ent/rule/organization.go:41-49`) — an unambiguous auto-filter by organization. `Organization` uses **neither** mixin (`organization.go:17-23` = PrimaryKey/CreatedAt/UpdatedAt only), so it is neither one of the 30 nor `Membership`. The corpus's own re-measurement paragraph (`security_compliance.md:92-97`) names `organization.go` as one of the only four `Policy()`-declaring files and then never accounts for it. This is the fence the text itself says "has now been wrong FOUR times." |
| C-2 | `shared_libraries.md:145` | authn — "**Imported by** \| via colony: app (the former cms / jobsimulation / skillpath usage is all folded in)" | The still-running `cms` and `jobsimulation` **husk repos still import `colony/authn` directly in their own Go source** — 7 import sites in `stack-demo/cms`, 9 in `stack-demo/jobsimulation` (`grep -rn "colony/authn" --include="*.go"`) — and the default `graphql` profile still **starts both containers** (`platform/docker-compose.yml:144` cms, `:83` jobsimulation, `profiles: [graphql, …]` at `:187`/`:140`). Every *other* library row in this same file gets this right and says so explicitly: colony `:42` ("plus the `cms`, `jobsimulation` and `roadrunner` containers… **husks**"), ai `:105` ("**and, as their own direct `go.mod` requires, the still-running `cms` and `jobsimulation` husk containers**"), taxonomy `:181` (same). The authn row alone drops them, so it reads as "no running container outside `app` does Clerk JWT auth" — false while M810 is open. |

---

## MINORS

1. **`security_compliance.md:73`** — cites `mixin.go:98` for `UserMixin`'s `Policy()`. Actual:
   `func (UserMixin) Policy() ent.Policy` is at **`mixin.go:99`**; `:98` is its doc comment.
   (`OrganizationMixin`'s `Policy()` at `mixin.go:126` — cited at `:69` — is **exact**.)
2. **`security_compliance.md:216-217`** — cites `skills.go:53-64` for `calculateSkillScore`.
   The function spans **`:53-62`**; `:64` is the header of the *next* function (`run`). The
   companion anchor `:75` (`score = utils.RoundFloat(float32(passed)/float32(total)*100, 2)`)
   is **exact**.
3. **`services/README.md:20`** — "And **three of the four** (cms, jobsimulation, roadrunner)
   still start CONTAINERS locally". Line **15** of the *same blockquote* says roadrunner "is
   the **fifth**", explicitly excluding it from "the four" (skiller, skillpath, jobsimulation,
   cms). Only **two** of the four still start containers; it is three of the **five**. The
   parenthetical names the right three services, so a reader still acts correctly — hence
   MINOR, not BLOCKER — but the arithmetic contradicts its own line 15.
4. **`service_taxonomy.md:37` (mermaid) and `:224`** — "Academy -->|GraphQL - **academy
   subgraph**| Backend" / "a GraphQL client of the platform `app` **academy subgraph**".
   There is **no academy subgraph**: `graphql-wundergraph/supergraph-config-prod.yaml` @ `60c229f`
   declares exactly one (`- name: backend`), and `academy.graphqls` is a schema *file* inside
   the backend subgraph (`app/internal/web/backend/graphql/graph/schemas/academy.graphqls`).
   Contradicts this same file's `:339` ("**`backend` alone (1)**"). `services/README.md:57`
   uses the correct phrasing ("over the `app` subgraph").
5. **Cross-file, outside my set (FYI for the seat that owns it):**
   `external_services.md:572` — item 4 of the residency list still reads *"the easiest of the
   **four** to miss"* after iter-49 grew that list from four to **five** (`:569`). Both
   `architecture_overview.md:247` and `security_compliance.md:186` correctly say "five".

---

## What I checked hardest, and what passed

The two files this iteration's repair touched were re-derived from scratch, not read.

### Multi-tenancy / Layer 1 (the four-times-wrong fence) — re-derived end to end
Run in `app/internal/data/ent/schema` @ `5ba17044`:

| Claim | Measured | Verdict |
|---|---|---|
| 139 `.go` files, 4 declaring no schema | 139; non-schema = `database_types.go`, `mixin.go`, `skiller_mixins.go`, `skillpath_mixins.go` | ✅ |
| 135 `ent.Schema` files | 135 | ✅ |
| 30 `OrganizationMixin{}` | 30 (and all 30 are real schemas; `OrganizationIDMixin{}` does not false-match) | ✅ |
| 7 `OrganizationIDMixin{}` = category, jobrole, similarity, skill, specialization, studio_document, studio_task | exactly those 7 | ✅ |
| plain `organization_id`, neither mixin → 18 | 18 | ✅ |
| minus `org_membership.go` (own policy) + `academy_feedback.go` (`UserMixin{}` at `:64`) → **16**, and the named 16 | exact set match, all 16 names | ✅ |
| 23 = 16 + 7 unpoliced-with-`organization_id` | ✅ | ✅ |
| only FOUR files declare any `Policy()`: `organization.go`, `mixin.go`, `user.go`, `org_membership.go` | exactly those 4 | ✅ |
| `Membership.Policy()` at `org_membership.go:172-188`, ending `privacy.AlwaysDenyRule()` | `:172` func, `:186` AlwaysDeny (Query), `:189` close | ✅ (range exact) |
| `job_simulation_session.go:5` self-statement | `:5` = *"L2: NO Ent privacy Policy; owner/org/tenant are plain fields."* | ✅ verbatim |
| `jobrole.go:18` / `category.go:15` "globally readable" | both inside the cited comment blocks | ✅ |
| **31 auto-filter by organization** | **32** — see BLOCKER C-1 | ❌ |

Only one schema dir exists in `app` (`find -type d -name schema -path "*ent*"` → one hit), so
139/135 is the complete universe.

### EU-egress enumeration (the other repaired passage) — all five paths traced
- `external_services.md:541`, `:545`, `:569` — all three cross-anchors land **exactly** on the
  Anthropic-Direct provider row, the "no ordered EU-first fallback chain" retraction, and the
  "**Five** things can send a request outside the EU" list. ✅
- `getClient` defaults `azureClientEu`, swaps on `flag_use_azure_us` → `ai.go:259/263/266/274` ✅
  (cited `:262-276`, `:263-277`, `:262-266` all contain it).
- 429 retry → direct OpenAI: `isThrottlingError` at `:129` ✅, applied `:166` ✅ and `:325` ✅,
  and `ChatCompletion`'s retry body **overrides `vendor = Openai`** (`:151-154`) — the claim is
  not just an anchor, the mechanism is real. ✅
- Bedrock pinned `eu-west-1`: `config.WithRegion("eu-west-1")` at `:87`, `NewAnthropic(&cfg,nil)`
  at `:92` (cited `:85-88` / `:85-95`). ✅
- Course Builder → first-party Anthropic: `coursebuilder/bedrock.go:109-112`
  (`newUnderlyingClient` → `NewAnthropicClientWithModel`) ✅, `ModelBackendName()` returns
  `"anthropic-api"` at `:100` ✅.
- Mistral OCR only: `internal/cms/studio/markdownManager.go:11` (import) + `:19`
  (`mistral.NewMistral(nil, os.Getenv("MISTRAL_API_KEY"))`) ✅, called from
  `studioManager.go:583` ✅.

### EU AI Act blockquote — the whole chain re-walked
`skills.go:75` divides ✅; `criterion.go:127` is the hardcoded `case check.EngineLlm:` ✅;
`:428` `NewLLMBulkChecker(c.logger)` ✅; `:168` dispatches `validateCodeDiff` ✅; `:450-475`
sets `success` from `cdiff.Diff == ""`, no model ✅; `checkValidationBulk.tmpl:27` asks the
model to *"assess whether the `<asset>` … meets or does not meet"* ✅ and returns
`{check_id, feedback, success}` ✅; **`checkerEngines` is assigned (`criterion.go:90`,
`validator.go:60`) and never read** ✅ — the "cite the DISPATCH, not that map" instruction is
correct. The "*most*, not all" hedge is accurate: `EngineTextDiff` results are appended
alongside the LLM ones. `ai_architecture.md:261-282` does carry the twin correction, as
`security_compliance.md:234` claims ✅.

### Counts re-derived rather than trusted
- **Supergraph ladder 5 → 4 → 3 → 1**: `supergraph-config-prod.yaml` at `749dc86^` = 5
  (backend, skiller, jobsimulation, cms, skillpath), `749dc86` = 4, `7c17e63` = 3,
  `915da06` = 1 ✅. Dates 2026-06-24 / 2026-07-21 / 2026-07-29 ✅. `915da06` **does** delete
  `cms.graphqls` **and** `jobsimulation.graphqls` in one commit — so the corpus's "**3→1**,
  the commit subject's 2→1 is wrong" is correct, and the commit subject really does say
  "supergraph 2→1" ✅.
- **"six Go services"**: `graphql` profile = backend + jobsimulation + cms + storage +
  roadrunner + always-on sentinel = **6**, plus gotenberg ✅.
- **"23 run-state tables"**: `app/terraform/migrations/20260722081626_jobsim_data_model.sql`
  creates exactly **23** tables ✅.
- **Taxonomy figures**: `.agentspace/snapshots/taxonomy/5afc0bcc…/manifest.json` —
  skills **42,790**, job_roles **22,470**, job_role_embeddings **18,919**, `public_only: true`,
  `predicate: org-null`, `source: primary-read`, `captured_at 2026-06-29` — every number and
  every provenance field ✅.
- **27 service docs**: 29 `.md` minus `README.md`/`TEMPLATE.md` = 27, and the index enumerates
  **all 27** (10 + 5 + 8 + 4) with **zero** broken links ✅. `architecture/README.md`'s 13 links
  all resolve ✅.
- **`authn` not a dependency**: 0 hits for `anthropos-work/authn` across all 7 `go.mod` **and**
  all 7 `go.sum` ✅ — `architecture/README.md:21` verified.
- **Shared-lib version pins** (every one exact): colony app/messenger `v0.35.2`,
  cms/jobsimulation `v0.35.1`, sentinel/storage/roadrunner `v0.34.3`; proto `v1.210.0` /
  `v1.207.0` / `v1.205.0` / `v1.200.0` / `v1.196.0`; ai `v1.40.2` in `cms/go.mod:9` +
  `jobsimulation/go.mod:11`; taxonomy `v1.2.0` in `cms/go.mod:13` + `jobsimulation/go.mod:15`
  (neither `// indirect`), indirect in storage/sentinel, absent from roadrunner ✅.
- **"6 of the 7 live Go service repos" import taxonomy, roadrunner the exception** ✅.
- **`app/main.go` six Connect handlers** at `:1178, :1179, :1187, :1195, :1204, :1218` — every
  anchor exact, and no `SkillPathSessionService` / RoadRunner handler exists ✅.
  `ROADRUNNER_RPC_ADDR` → **0 hits** in `app`'s Go source (grep ran clean, exit 1 = genuine
  absence) ✅, and the var itself is at `docker-compose.yml:118` ✅.
- **`roadrunner/main.go:7`** imports colony; `NewVersionConfig` used at `main.go:17` ✅.

### Compose / repos.yml / terraform anchors — every one hit
`repos.yml:10-13` (app, migrations true, schema public) ✅ · `docker-compose.yml:18`
`search_path=sentinel` ✅ · `:83` jobsimulation ✅ · `:118` ROADRUNNER_RPC_ADDR ✅ · `:144` cms ✅ ·
`:256` messenger's `CMS_RPC_ADDR=http://cms:8091` ✅ · `:281` roadrunner ✅ · `:311`/`:342`
studio-desk + `profiles: [studio-desk, all]` ✅ · `:318`/`:334` VITE_GRAPHQL_ENDPOINT →
`:8082/graphql/query` ✅ · `:352`/`:361` NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT → `:8082` ✅ ·
**no `graphql` service and no `graphql-wundergraph` repos.yml entry** ✅ · every port in the
Tier-1 table (8087, 8081-8083 + `META_PORT=8084`, 8400-8401, 8090-8091, 8300-8301, 10400-10401,
3200, 8200-8201, 8080, 9000/9100, 3000, 5432, 6379 `bitnamilegacy/redis:latest`) ✅ ·
platform `b56d731`+`360efd4` merged as `2adcf71` on **2026-07-31** ✅ ·
`a2a3ee6` (2026-02-27) and its parent's `docker-compose.yml:384` `directus/directus:10.10.1`,
`:386` `8055:8055`, `:409` `ADMIN_PASSWORD=password` — **all three exact**, so the
"correction of the correction" at `service_taxonomy.md:299-305` is itself correct ✅ ·
`graphql-wundergraph/terraform/main.tf:20` `= 1` ✅, `locals.tf:8` `port = 8080` ✅,
`main.tf:48-49` container→host 8080 ✅, `config.prod.yaml:5` `listen_addr: 0.0.0.0:8080` ✅ ·
`roadrunner/terraform/main.tf:19` `= 1` ✅, `cms/terraform/main.tf:39` `= 0` ✅,
`jobsimulation/terraform/main.tf:40` `= 0` ✅.

### Tier-2 / frontend claims
`next: ^16.2.7` in **exactly four** next-web-app apps (web, hiring, integration, maintenance) ✅ ·
Clerk org-switch anchors `useOrgSelection.tsx:94` (web **and** hiring), `useResolveActiveOrg.tsx:107`,
`useActivateMembershipOrg.tsx:81` — all exact, all `clerk.setActive({ organization })`, no
re-auth ✅ · studio-desk: **0** react/vue/angular deps across both `package.json`s, **0**
`.tsx`/`.jsx` in the repo, GPT-5.x present, and `app/public/l12n/` holds **exactly 7** locale
files (de, en, es, fr, it, ja, nl) — the "7 languages" claim ✅ · `gen.py` registers **exactly
nine** arguments, exactly the nine named, `parse_known_args` at `gen.py:19`, **no `--template`** ✅,
and `studio/CLAUDE.md:12-14` is verbatim the quoted entry point ✅ · Studio-Room pulled in by CI
via `additional_repo: "anthropos-studio-room:studio"` (`app/.github/workflows/build-production.yml:29`)
and the image is `python:3.11-slim` (`app/Dockerfile.dev:26`) ✅ · ant-academy: `server.js:14,18`
(throws when unset) ✅, `serverTenant.js:145` is the DB-authoritative read and `:115-145` carries
the verbatim *"NO FS-as-published fallback … not reversible-on-error"* ✅, `backendContent.js:36`
+ `:102-103` ✅, `beacon/route.js:36` + `:41-55` (`UPSERT_CHAPTER_PROGRESS`, `SET_LAST_ACTIVITY`) ✅,
**0** serwist/workbox in `package.json` ✅, `RegisterServiceWorker.jsx` is a kill-switch that
`unregister()`s ✅, regression fences at `next-scaffold.test.js:106,111` and
`react-compiler-config.test.js:41` ✅, manifest survives with `"display": "standalone"` declared at
`layout.jsx:132` ✅, ports 3077/8555 ✅, Expo `~54.0.33` ✅, React `^19.2.5` ✅,
`NEXT_PUBLIC_FEATURE_TRAINING_COACH` default-off ✅, `gpt-5.2` at
`code/ucourses/ucourse-engine/assistant/agent.js:13` ✅ · rext `demo-stack/patches/academy-fs-published-fallback`
exists ✅ · `directus/directus:11.6.1` pinned in rext ✅.

### `ai` library (from the module cache @ v1.40.1)
`ai.AI` has **exactly** the nine methods listed ✅ · `openai.New`/`NewOpenAI`/`NewAzure`,
`anthropic.NewAnthropic(cfg,key)`, `mistral.NewMistral` ✅ · Mistral panics ×4, Anthropic
panics ×2 ✅ · Azure default API version `2025-04-01-preview` ✅ · `retry.Attempts(10)` ✅ ·
`MetaData{Usage any; Model Model}` — **no cost field**, so "the `ai` library does NOT track
cost" ✅ · the Anthropic `{"` prefill at `anthropic/completion.go:106,143` ✅ ·
consumer-side cost in `app/internal/aiusage/ai_usage.go:36` (hardcoded model switch) +
`:98` `Event_AiUsage` handler ✅ · `app/internal/askengine/bedrock.go:15-18` raw
`anthropic-sdk-go` ✅.

**Result: 2 blockers, 5 minors (one of which is outside my file set).** The file that came
closest to a clean sheet is `architecture/README.md` (38 lines, every claim and every link
verified true). The densest file, `service_taxonomy.md`, produced only a terminology minor
despite ~60 checkable anchors.
