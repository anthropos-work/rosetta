# Seat A — M257x clause-5 KB-fidelity reading

## 1. Header

Corpus: `/Users/marco/workspace/anthropos/rosetta`, branch `m257x/platform-realignment`,
HEAD `57dfbfded8791fcb12a4651d747247ce9d04d7f0` (verified via `git rev-parse HEAD` / `git branch --show-current`).

Ground-truth clones consulted (all under `stack-demo/`, sha confirmed with `git rev-parse HEAD` in each):

| clone | sha consulted |
|---|---|
| `stack-demo/app` | `5ba1704482cf812b130c2d3673afd09f4f7f22e5` |
| `stack-demo/app/studio` | (in-tree, read directly) |
| `stack-demo/platform` | `2adcf714bd877a205e8948f59a23db49b884c054` |
| `stack-demo/graphql-wundergraph` | `60c229f39adcbbe75c84cd58f0f45052b5423372` |
| `stack-demo/next-web-app`, `studio-desk`, `ant-academy`, `cms`, `jobsimulation`, `messenger` | read as cloned |
| `.agentspace/rosetta-extensions` (authoring copy) | `a91f8f7` |

**Positive control — `wc -l` on every assigned file** (one invocation:
`cd /Users/marco/workspace/anthropos/rosetta && wc -l <7 paths>`):

| file | lines | assigned | read in full |
|---|---:|---:|---|
| `corpus/architecture/external_services.md` | 814 | 814 | yes |
| `corpus/services/backend.md` | 271 | 271 | yes |
| `corpus/services/graphql-wundergraph.md` | 265 | 265 | yes |
| `corpus/services/academy-backend.md` | 141 | 141 | yes |
| `corpus/services/coursebuilder.md` | 139 | 139 | yes |
| `corpus/services/skiller.md` | 66 | 66 | yes |
| `corpus/services/TEMPLATE.md` | 46 | 46 | yes |
| **total** | **1742** | 1742 | — |

No file read short or empty. Every search below was run with a positive control in the same
invocation; where a control returned 0 the pipeline was repaired before concluding.

---

## 2. BLOCKERS

### B1 — `corpus/architecture/external_services.md:604-618` — the `platform/.env` AI-config block names three env vars the platform never reads, misspells a fourth, and omits the one that actually supplies the OpenAI key

**The claim (quoted, `:604` + the fenced block `:606-618`):**

> AI services are configured via environment variables in `platform/.env`:
> ```bash
> # OpenAI
> OPENAI_API_KEY=sk-proj-xxxxx
> OPENAI_ORG_ID=org-xxxxx
> …
> # Azure OpenAI
> AZURE_OPENAI_KEY=xxxxx
> AZURE_OPENAI_ENDPOINT=https://resource.openai.azure.com/
> AZURE_OPENAI_DEPLOYMENT=deployment-name
> ```

**What is true.** The Go platform reads `OPENAI_KEY` and `AZURE_OPENAI_ENDPOINT_URL`.
`OPENAI_ORG_ID`, `AZURE_OPENAI_DEPLOYMENT` and bare `AZURE_OPENAI_ENDPOINT` are read **nowhere**,
and `OPENAI_KEY` — the var that actually feeds the direct-OpenAI client — is **missing from the block**.

Ground truth (invocation: `grep -rn '"<NAME>"' --include="*.go" stack-demo/app/`, counts stated per name):

| name | occurrences in `app/**/*.go` | evidence |
|---|---:|---|
| `"OPENAI_KEY"` | 4 | `app/internal/jobsimwiring/wiring.go:140` `getenv("OPENAI_KEY")`; `app/main.go:817` `os.Getenv("OPENAI_KEY")` (coursebuilder cover generator); `app/cmd/createTaxonomy/main.go:90` |
| `"AZURE_OPENAI_ENDPOINT_URL"` | 12 | e.g. `app/cmd/userSkills/main.go:55` `openai.NewAzure(os.Getenv("AZURE_OPENAI_KEY"), os.Getenv("AZURE_OPENAI_ENDPOINT_URL"), nil)` |
| `"AZURE_OPENAI_ENDPOINT"` (bare) | **0** | — |
| `"OPENAI_ORG_ID"` | **0** | — |
| `"AZURE_OPENAI_DEPLOYMENT"` | **0** | — |
| `"OPENAI_API_KEY"` | 1 | **and it is not a platform read** — `app/internal/cms/studio/studioManager.go:1067` is a *remap target* in `studioSubprocessEnv()`: `"OPENAI_API_KEY": "CMS_STUDIO_OPENAI_API_KEY"`. The surrounding block (`:1055-1076`, read in full) **strips every bare AI name from the inherited environment first**, precisely so a bare `OPENAI_API_KEY` in the container env is *not* handed to the child |

Positive controls for the zero counts: the same grep form returned 12 for `AZURE_OPENAI_ENDPOINT_URL`
and 4 for `OPENAI_KEY` in the same pass, so the pattern engine was working.

Independent corroboration (rext `stack-secrets/secretdna/secret-dna.json`, the 6-repo secret DNA):
its `repo: platform, file: .env` genes are `OPENAI_KEY` ("DISTINCT-SIMILAR from OPENAI_API_KEY — may
hold a DIFFERENT token; do NOT auto-alias"), `AZURE_OPENAI_KEY`, `AZURE_OPENAI_ENDPOINT_URL`,
`SKILLER_*` twins. The DNA has **0** entries for `OPENAI_ORG_ID`, **0** for `AZURE_OPENAI_DEPLOYMENT`,
and **0** for bare `"AZURE_OPENAI_ENDPOINT"`.
(Guard applied: the 4 hits for those names elsewhere under `.agentspace/rosetta-extensions` are
`stack-core/tests/fixtures/**/corpus/architecture/external_services.md:593,601` — copies of *this very
corpus file* used as test fixtures. That is a probe satisfying itself; it was excluded.)

**How a reader is harmed.** Anyone provisioning `platform/.env` from this checklist sets
`AZURE_OPENAI_ENDPOINT` (never read → the Azure client is built with an empty endpoint) and never sets
`OPENAI_KEY`, so `openai.NewOpenAI(openaiKey)` at `app/internal/jobsimulation/ai/ai.go:80` — the client
this very document's residency section (`:584`) says every unset-`ai_vendor` simulation lands on — and
the Course Builder cover generator (`app/main.go:816-819`) both come up keyless. This is a repair
landed at one site and left standing at another: the repo-root `CLAUDE.md` and
`corpus/services/coursebuilder.md:98` both name **`OPENAI_KEY`** and explicitly warn that naming the
wrong var "fixes nothing".

---

No other blocker survived verification. Two candidates were **measured and dropped** (recording them
because both are the shape this milestone keeps producing):

* **`coursebuilder.md:102` `course.refine`=1.** `app/internal/credits/cost.go:29` — the *package
  doc-comment* — says "course.refine → 5 credits per refine turn". Reading around it (rule 5),
  the authoritative map at `cost.go:72-76` is `ActionCourseRefine: 1, // D1 ruling: flat 1 credit per
  refine turn`, and the const's own doc at `:45-53` states the D1 ruling. **The corpus is right and the
  in-repo comment is stale.** Not a finding.
* **`backend.md:91` / `skiller.md:24` "`categoryTree`/`fullCategoryTree` were dropped, not ported".**
  `grep -c` on `graphql-wundergraph/schemas/backend.graphqls` returns 1 for each — but both hits are the
  *same* line, `:5261`, a comment that reads "… matchJobRole, categoryTree, fullCategoryTree) stay
  unported — no consumers." **The corpus is right.** Not a finding.

---

## 3. MINORS (15)

Line drift / mis-anchored cross-references (6):

1. `external_services.md:804` — "Consistent with :512 above, where the same correction is already
   recorded." `:512` is the *subscriptions* bullet; the `docker compose restart graphql` → `restart backend`
   correction is at `:520`.
2. `graphql-wundergraph.md:82` — cites "`:174-176` of this same doc already said `localhost:5050`
   refuses"; that text is at `:178`.
3. `graphql-wundergraph.md:128` — "…which is why a subgraph SDL change rebuilt the router, as `:84`
   describes." `:84` is the "no application source here" blockquote; the rebuild statement is at `:110-111`.
4. `academy-backend.md:15` — "(Consistent with :74-76 below.)" `:74-76` is the section heading +
   "How to find the API"; the matching `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` statement is at `:79-80`.
5. `backend.md:39` — "the mux registers five handlers unconditionally (`main.go:1178-1218`)". The count of
   five is right (Users `:1178`, Organizations `:1179`, Skiller `:1187`, JobSimulation `:1195`, LabSession
   `:1219`), but the **fifth `mux.Handle` is at `:1219`**, one line past the stated range. (`:1203-1205`
   for the conditional CMS edge is exact.)
6. `external_services.md:429` — the quote attributed to `graphql-wundergraph/CLAUDE.md:39`
   ("Since cms-in-app the platform compose `graphql` service builds from the **production** Dockerfile")
   is **verbatim from `:33`**. Line 39 asserts the same thing in different words ("Since the compose stack
   builds from the production Dockerfile…"), and `git blame -L 39,39` does confirm `60c229f3`,
   Luca Casartelli, 2026-07-30 — so the substance and the provenance hold; only the quote's anchor drifts.

Stale / overstated facts (5):

7. `backend.md:253` — "not in the top-level `migrations/` dir (which holds only `atlas.sum`)". There is
   **no top-level `migrations/` dir** at `app@5ba17044` (`ls migrations/` → *No such file or directory*);
   it was removed at `app@6a46e844` "chore(migrations): remove obsolete atlas.sum file", and `atlas.sum`
   now lives in `terraform/migrations/atlas.sum`. The actionable half of the sentence (migrations live in
   `terraform/migrations/`, per `atlas.hcl` `dir = "file://terraform/migrations"`) is correct.
8. `external_services.md:541`, `:575`, `:588` — Studio-Room's `anthropic` and `openai` `TARGET SERVICE`
   are presented as live selections ("selected by `TARGET SERVICE = anthropic` in `configs/*.ini`").
   At HEAD **all 30 model slots across all three `app/studio/configs/*.ini` read `azure`**
   (`production_config.ini:26-36`, `development_config.ini:26-36`, `config_template.ini:37-47`); `anthropic`
   and `openai` appear only in the allowed-values comment (`production_config.ini:25`).
   `grep -rnE '^[A-Z_]+_MODEL = (anthropic|openai)' configs/*.ini` → 0 (positive control:
   `grep -cE '^[A-Z_]+_MODEL = azure'` → 10 per file). The *mechanism* claims are correct — the provider
   set is `{openai, azure, anthropic}` (`services/ai.py:704-709`), the openai arm is a bare
   `OpenAI(api_key=…)` (`:383`), and `gen.py:48-53` does let the env `ANTHROPIC_API_KEY`/`OPENAI_API_KEY`
   override the ini value — but no shipped config selects either arm.
9. `external_services.md:648` — "`AI_PROVIDER_CHAIN`, default `azure-openai,openai`". The env var has
   **no default**; when unset, `loadAIServiceConfig` falls back to Azure → OpenAI → **Anthropic**
   (`studio-desk/src/services/ai/config.ts:165-171`).
10. `graphql-wundergraph.md:243` — "`ARCHITECTURE` … (`linux/amd64`)". The prod CI build arg value is
    `ARCHITECTURE=x86_64` (`graphql-wundergraph/.github/workflows/release.yml:100,103`).
11. `backend.md:186,189` — `feat/hiring-talk-to-data` and `feat/taxonomy-translations` no longer exist in
    the `app` clone (92 branches, 35 `remotes/origin/feat/*`; positive control: `feat/company-context-m1m2`
    at `:188` **does** exist).

Undercounts / omitted list members (4):

12. `coursebuilder.md:113-114` — "~55 test files" / "~25 boundary test files": actual
    `internal/coursebuilder/*_test.go` = **59**, `internal/web/backend/coursebuilder/*_test.go` = **32**
    (`internal/coursebuilder/*.go` = 98 vs "~100" ✓). `:129` "142 changelog lines": `grep -c coursebuilder
    CHANGELOG.md` = **144**.
13. `graphql-wundergraph.md:99` — the `.github/workflows/` list omits `bump-version.yml` (3 workflows, not 2).
14. `academy-backend.md:109-119` — the `cmd` binary list omits `cmd/academy-asset-upload`.
15. `backend.md:124-177` — the "Key directories" tree omits ~20 real `internal/` packages, including
    `credits/` and `embeddings/`, which sibling docs (`coursebuilder.md:83`, `skiller.md`) cite by path.
    Selective by design, but the two named ones are load-bearing elsewhere.

---

## 4. Audited zeros — read in full, found clean

**`corpus/services/TEMPLATE.md` (46/46).** Pure scaffolding, no factual claim to grade. Clean.

**`corpus/services/skiller.md` (66/66).** Every gradeable claim checked and passed:
the merged-shape banner (`public` schema, Ent models in `app/internal/data/ent/schema/`,
`SkillerService` served by `app/internal/rpc/skillerrpc/` — dir exists); the taxonomy floors
(**≥42,790 skills / ≥22,470 job roles**) reproduce **exactly** against
`.agentspace/snapshots/taxonomy/5afc0bccf1df7ef538b643321fc6362f/manifest.json` — `public_only: true`,
`predicate: "org-null"`, `captured_at 2026-06-29T14:43:00Z`, `skills 42790`, `job_roles 22470`,
`skill_embeddings 42790`, `job_role_embeddings 18919`, `categories 23`, `specializations 1447`
(invocation: `python3` over the manifest's `tables[]`), which also confirms the "18,919 is the
*embedding* row count transcribed onto the role count" diagnosis; IVFFLAT indexes on both embedding
tables (`terraform/migrations/20260615130000_skiller_taxonomy.sql:61,251`); `extensions.vector(1536)`;
**8 `ContentLanguage`s** — verified exactly (`backend.graphqls:1469-1478`: english, italian, spanish,
french, german, dutch, japanese, portuguese); `categoryTree`/`fullCategoryTree` dropped
(`backend.graphqls:5261`).

**`corpus/services/coursebuilder.md` (139/139).** All 15 named engine files present; `imagegen/`,
`assets/` present; all 5 in-repo docs present; `TestProductionConfig_WiresPlanner` at
`production_config_test.go:11`; `DefaultTargetScore = 90` (`refine.go:224`);
`DefaultAuthorModelID = "eu.anthropic.claude-opus-4-8"` (`bedrock.go:23`),
`DefaultGraderModelID = "eu.anthropic.claude-sonnet-4-6"` (`:29`); the backend switch
`bedrock.go:105-114` (`ANTHROPIC_API_KEY` set → `askengine.NewAnthropicClientWithModel` → a first-party
`anthropic.NewClient(option.WithAPIKey(...))` with `directModelID`-stripped ids, `anthropic.go:35-51`)
and `ModelBackendName()` at `:98-103` logged at `main.go:762`; the prod-requires-the-key chain
`terraform/variables.tf:635-638` (sensitive, **no default**) → `ssm.tf:328-333` → `main.tf:555` — all
three exact; `main.go:816-819` reads **`OPENAI_KEY`** and `COURSEBUILDER_OPENAI_IMAGE_KEY` survives only
in `internal/coursebuilder/{README,GO-LIVE-RUNBOOK}.md` (0 Go hits), removed at `app@68c24512`;
`DefaultMaxMonthlyCOGSUSD = 500.0` / `DefaultMaxDailyCOGSUSD = 0.0` (`budget.go:27,41`);
`DefaultSessionsPerOrgPerDay = 50` (`rate.go:28`); credits 5/1/1 (`credits/cost.go:72-76`);
`coursebuilderWorkerConcurrency = 20` (`main.go:172`), `TypeCoursebuilderRun`, `MaxRetry(1)`,
`Retention(6*time.Hour)` (`run_task.go:49,197-204`); migration `terraform/migrations/20260717151144.sql`
creates `course_builder_sessions` + `credit_transactions` + `organization_credits` (`:2,40,56`);
every route in the `:72-75` table present at `handler.go:2859-2926`; Go 1.26.4 (`go.mod:3`);
`v1.363.2` @ `5ba17044` (`CHANGELOG.md:5`).

**`corpus/services/academy-backend.md` (141/141).** Provenance exact — `0e37771f` = 2026-06-05
"Academy backend v1.0 'ground truth' … (#903)"; `internal/academy/academy.go:6-9` is verbatim the
`aiacademy`-removal statement and `internal/aiacademy` is gone. The **Ent-label-vs-plural-table
warning is correct**: all 11 named tables exist in the plural form
(`academy_chapter_progresses`, `academy_last_activities`, `academy_chapter_times`,
`academy_certificates`, `academy_bookmarks`, `academy_feedbacks`, `academy_series`,
`academy_skill_paths`, `academy_chapters`, `academy_chapter_bodies`, `academy_path_embeddings`).
Cert contract exact — `AAS-\d{4}-[A-HJKMNP-TV-Z2-9]{6}` (`certificate.go:42`, a 30-symbol alphabet)
via `crypto/rand`; `paid` default true on the skill path (`academy_skill_path.go:69-71`); one body per
`(chapter_slug, locale)` with EN base + de/es/fr/it/nl/pt (`academy_chapter_body.go:18,58-59`);
`SELECT … FOR UPDATE` with the SQLite skip (`academy.go:45-57`, `sqlite_test.go`);
`ACADEMY_CONTENT_API_TOKEN` shared-token middleware (`content_admin.go:35`);
`GET /content/catalog.json` registered at `content.go:23`; nightly
`academy_embedding_refresh @ 03:00` (`internal/worker/tasks/academy_embedding_refresh.go:10`);
all named test files present; all 4 managers present (`academy.go`/`content.go`+`body.go`+
`content_import.go`/`embeddings.go`/`asset.go`); every listed GraphQL query exists in
`academy.graphqls:640-689` with `academyCertificate` and the catalog set carrying `@public`.

**`corpus/services/graphql-wundergraph.md` (265/265)** — the whole "two states" banner and the entire
historical fence verified commit-by-commit, see §5 for the two things I could not check:
`terraform/main.tf:20 service_desired_count = 1` (exact line); `locals.tf:8 port = 8080`;
`main.tf:48-49` containerPort/hostPort 8080; `config.prod.yaml:5 listen_addr: 0.0.0.0:8080`;
`health_check_path = "/health"` (`main.tf:26`); `package.json` is exactly
`{"name":"graphql-wundegraph"}` and the CLAUDE.md heading carries the same misspelling;
`grep -c 5050 platform/docker-compose.yml` → **0**; router `0.275.0`, `wgc@0.104.0`,
`federation_version: =2.3.2`, `node:22.11-alpine` all pinned as stated;
`config.compose.yaml` has `playground_enabled: true`, `introspection_enabled: true`,
`graphql_path: /graphql`, `max_request_body_size: 35MB`, the three apollo-compatibility flags;
`Makefile` has `run` (`-p 5050:8080`, README's aarch note) and `updatesubg`;
`terraform/tests/` exists and there is no other test suite;
`ci/update-subgraph.sh:9` carries **exactly one** `gh release download … -R anthropos-work/app`;
`schemas/` holds `backend.graphqls` alone; `subgraphs.conf` = `BACKEND=v1.360.0`.
The **subgraph ladder is exact** — counting `^  - name:` in `supergraph-config-prod.yaml` per commit:
`749dc86~1` → **5**, `749dc86` (2026-06-24) → **4**, `7c17e63` (2026-07-21) → **3**,
`915da06~1` → **3**, `915da06` (2026-07-29) → **1**; `git show --name-status 915da06` marks **both**
`schemas/cms.graphqls` and `schemas/jobsimulation.graphqls` `D`; and `915da06`'s subject really does
read "supergraph 2→1" against a 3-entry tree — the doc's "the config file, not the commit message, is
the source of truth" is correct.
The **subscription retraction is exact**: `grep -rn "sse\|subscription" graphql-wundergraph/*.yaml` → 0
(positive control `grep -rln backend …*.yaml` → all three supergraph configs); `915da06~1`'s prod config
still carries `subscription: protocol: "sse_post"` on the `jobsimulation` entry alone;
`git log -S 'protocol: "ws"'` over HEAD's history → **0** lines (control: `-S sse_post` → **5** commits);
`git merge-base --is-ancestor bba862f HEAD` → **rc 1**; `git branch -a --contains bba862f` names
`remotes/origin/feat/use-web-socket` and nothing else.
The **compose-era table is exact**: `63d285c` (service `wundergraph`, `context: git@…#main`, **no
`dockerfile:` key**) → `d92e84e` (renamed `graphql`, still no key) → `a2a3ee6` (`context:
../graphql-wundergraph`, still no key) → `719befb` (still no key) → `2c85211` (adds `dockerfile:
Dockerfile.dev`) → `67ba772` (`context: ..`, path `graphql-wundergraph/Dockerfile.dev`) →
`b56d731` (`  graphql:` still present at **`:22`**, parked behind `profiles: [wundergraph-deprecated]`)
→ `360efd4` (no `graphql:` key; control `  backend:` at **`:28`**). Historic routing URLs
`http://jobsimulation:8400/query` and `http://cms:8090/query` confirmed at `915da06~1`.
The smoke-test section is exact: `gqlauthz.go:186` returns `fmt.Errorf("unknown viewer: %w",
errForbidden)` with `errForbidden = errors.New("Forbidden")` (`:251`); the three exemptions at
`:175-178` are `isPublicOnlyOperation` / `isFederationQuery` / `(environment == colony.Development &&
isIntrospectionOnlyQuery)`; both named regression tests exist (`gqlauthz_test.go:132,176`);
`/graphql` = Apollo Sandbox **gated on `colony.Development`** and `/graphql/query` = the handler
(`backend.go:312-317`); `repos.yml` historically tagged it `type: node-npm` (`b56d731:repos.yml:45-47`);
the archived CLAUDE.md's "Version Tracking" (`:70`) and `supergraph-config-local.yaml` (`:85`) are
indeed stale (only `compose`/`dev`/`prod` variants exist).

**`corpus/services/backend.md` (271/271)** — everything except the 3 minors above:
the merge banner's 3→1 ladder (as verified above); `main.go` mux registrations
(`:1178`, `:1179`, `:1187`, `:1195`, conditional CMS `:1203-1205`, LabSession `:1219`);
**`SkillPathSessionService` = 0 occurrences** in `app`'s Go source and no `skillpath…v1connect` import
(positive control: `JobSimulationService` matched 5 files) — and the Trap-C citation is exact:
`app/CLAUDE.md:72` and `app/knowledge/architecture.md:28` both still list it;
`main.go:735-738` is verbatim the `LABS_API_URL` conditional wiring with real
`internal/labs/labsapi/` + `internal/labs/adapter/`; `lab_sessions` Ent table exists;
`POST /api/webhook/directus` registered (`backend.go:324`) and fail-closed on
`DIRECTUS_WEBHOOK_SECRET` (`main.go:1079-1080`); ports 8082/8083/8084 with the container publishing
8081/8082/8083 (`docker-compose.yml:37-40`, `PORT=8082`/`RPC_PORT=8083`/`META_PORT=8084` at `:55-59`);
**13 `ai_readiness_*` tables** — enumerated exactly (`grep -rhoE '"ai_readiness[a-z_]*"'
internal/data/ent/ | sort -u` → the 13 plural names), and the "four a 9 omits" list is right
(`recommendations`, `email_overrides`, `notification_logs`, `notification_optouts`), with
`live_snapshots` + `text_translations` correctly *not* counted as omissions; `internal/aireadiness/`
is a top-level package with `manager.go`/`cycles.go`/`diagnosis.go`/`compare.go`/`csv.go` and
`internal/workforce/` contains **no** `readi*` file (verified by listing the whole dir);
`internal/academy/academy.go:6-9` verbatim; RPC 60s write timeout (`main.go:1234`);
the messenger-address split at `docker-compose.yml:255-256` (`BACKEND_USERS_RPC_ADDR`/`SKILLER_RPC_ADDR`
→ `http://backend:8083`) vs `:256`/`:258` (`CMS_RPC_ADDR=http://cms:8091`,
`JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401`) and `main.go:1196-1202`'s
"Additive + DORMANT … until the M809 re-point" comment — **quoted verbatim, exact line range**;
Atlas config (`atlas.hcl` `dir = "file://terraform/migrations"`, `src = "ent://internal/data/ent/schema"`,
169 migrations, `make migrations` = `atlas migrate diff --env local`); `make up PROFILE=backend` really
does also bring `postgresql`, `redis`, `sentinel` (profile-less), `gotenberg`
(`profiles: [graphql, backend, all]` at `:384`).

**`corpus/architecture/external_services.md` (814/814)** — everything except B1 and its minors:
the Directus posture (no `directus` service in the platform compose; `backend`'s `environment:` block is
**exactly `:43-67`** and carries **no `DIRECTUS_*`**; the only explicit setter is `cms` at
**`:164-165`** — both line anchors exact); the **correction-of-the-correction at `:143-148` is
verified true** — `git show a2a3ee6^:docker-compose.yml` has `:384 image: directus/directus:10.10.1`,
`:386 - 8055:8055`, `:409 - ADMIN_PASSWORD=password`, all three at the exact stated lines, while
`git log --all -S "admin@example.com"` returns **0** commits (control: `-S "ADMIN_PASSWORD=password"`
returns 2) — so "only the email is unfound" is precisely right; the **nine-container** default profile
(profile-less `postgresql`+`redis` from `common.yml` — confirmed no `profiles:` key anywhere in
`common.yml` — plus `sentinel` (also profile-less) · `backend` · `jobsimulation` · `cms` · `storage` ·
`roadrunner` · `gotenberg`) = 9; the four frontend endpoints at `:318`/`:334`/`:352`/`:361`, all
`:8082/graphql/query`; `app/cms_reader_switch.go`'s "a DIRECT domain call — no proto round-trip …
and no internal traffic to a standalone cms" quoted **verbatim**; `main.go:971-973` is exactly the
`DIRECTUS_BASE_ADDR` `log.Fatalf`; the `--local-content` re-point —
`stack-injection/gen_injected_override.py:53` `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")` and
`:636-637` the re-point, with `stack-injection/tests/test_injection.py:1051`
`test_backend_the_actual_reader_is_repointed` present and the old pinning test gone;
`type Mutation` at `backend.graphqls:4053` and `type Query` at `:4912` with **no** `type Subscription`
and exactly **6** `Subscription` substring hits, all Stripe/plan; `Dockerfile.dev` has exactly one
schema `COPY` (`:18`), one `awk` (`:23`), the replacement comment at `:19-20`;
**the AI-routing section is anchor-perfect** — vendor consts `:30-33`, `getClient` `:259-289`,
Azure/PostHog branch `:264-276`, `AnthropicAws`+`Anthropic` → the same client `:280-283`,
`isThrottlingError` `:130-141`, the 429 retry `:150-155` / `:296-302` / `:326`, the OpenAI client at
`:80`; the caller-default chain — `AIVendor *AIVendor` at
`internal/cms/directus/collections/jobsimulation.go:905`, `aiVendor := simulation.Openai` at
**`:1302-1305`**, mapped at `internal/jobsimulation/simulator/ai/ai.go:58-59`, and the `default:` arm
at **`:114-115`** — all four exact; Mistral confined to OCR (`markdownManager.go:19`,
`studioManager.go:583`, and those are the only non-test `mistral` hits in `app/internal/`);
Bedrock model ids at `report_agent.go:31` and `askengine/bedrock.go:25`;
**`0` hits for `bedrock|boto3` under `app/studio/`** (control: `anthropic` matched 5 files), so the
iter-48 "Studio-Room was never on Bedrock" correction holds; Studio-Room anchors `ai.py:627-664`,
`:383`, `:704-724`, `:706-708`, `config_template.ini:30-31` all exact;
**the LiveKit correction is verified in context** (`calls/livekit.go:110` default bare name, `:120` the
explicit `*location == LocationEu` branch also bare, `:126` `fmt.Sprintf("anthropos-agent-%s",
*location)` only in the non-EU `else`; `grep -rn "anthropos-agent-eu"` across the whole `stack-demo`
tree → **0** with control `anthropos-agent-us` → 4 hits);
the taxonomy figures (see skiller.md above); `NEXT_PUBLIC_GRAPHQL_ENDPOINT` genuinely does **not**
exist in `next-web-app` (control: `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` matched 5 files including
`packages/graphql/src/server/server.graphql.ts`); the per-app Clerk SDKs
(`apps/{web,hiring,integration}/package.json` `@clerk/nextjs ^6.39.2`, `apps/mobile` `@clerk/clerk-expo`,
`studio-desk` `@clerk/clerk-js` + `@clerk/express`, `ant-academy/code` `@clerk/nextjs`);
`app/internal/web/backend/backend.go:130` is the `/api/webhook/clerk` entry in the authn skip-list.

---

## 5. Unverified — could not be checked, and why

1. **Production terraform RPC addresses.** `backend.md:195` ("`http://backend.internal.anthropos:8081`"),
   `backend.md:79` and `skiller.md:19` ("`http://backend:8081` in production terraform"). The variables
   are declared **without defaults** (`messenger/terraform/variables.tf:87`, `cms/terraform/variables.tf:76`,
   `app/terraform/variables.tf:197`); the values live in `anthropos-work/infrastructure`, which is not
   cloned. `grep -rn "backend.internal.anthropos"` across every cloned repo matched **only**
   `graphql-wundergraph/supergraph-config-prod.yaml:6`. Note the two corpus statements differ from each
   other (`backend:8081` vs `backend.internal.anthropos:8081`); I cannot adjudicate.
2. **"The repo is ARCHIVED on GitHub, 2026-07-30"** (`graphql-wundergraph.md:8`,
   `external_services.md:3`, `:357`). `gh` is unavailable. Consistent with the clone's last commit
   (`60c229f`, 2026-07-30) but that is not proof of archival.
3. **colony behaviours.** "a second `AddSubscriber` for the same stream silently overwrites the first"
   (`backend.md:27,220`) and the `colony/authn` claims (`external_services.md:75`, `:69`). `colony` is
   not cloned.
4. **`proto` / `taxonomy` library claims.** `ContentLanguage` resolves to
   `proto/go/domain/cms/v1/content.Language` — but I closed this one anyway off the composed SDL
   (`backend.graphqls:1469-1478`, 8 values). The `taxonomy` library being "NodeID helpers only"
   (`skiller.md:37-39`) remains unverified — repo not cloned.
5. **The 2026-07-08 live-prod counts** in `backend.md:70-74` (43,584 total skills incl. 794 org-private;
   `job_roles` org-NULL = 22,490; `skill_embeddings` = 43,584). No prod DB access in this pass. They are
   internally consistent (42,790 + 794 = 43,584) and agree with the 2026-06-29 snapshot within the drift
   the doc itself states; `categories = 23` and `specializations = 1,447` match the snapshot exactly.
6. **Infrastructure-side skiller decommission** (`skiller.md:25-27`): ECS service + terraform module
   removal, "app PR #989", the orphaned ECR repo. `infrastructure` repo not cloned; `gh` unavailable.
7. **Production Directus** (`external_services.md:284-287`, `:739-743`): S3 file storage, CDN, the
   `directus` schema on prod. No prod access; the `DB_SEARCH_PATH=directus` in the deleted historical
   compose block is consistent but is not the prod deployment.
8. **Runtime media claims** (`external_services.md:670`, `:684`): LiveKit MP3 audio, Chime composited
   MP4 grid. Not statically checkable from these clones.
9. **`external_services.md:672`** — "ElevenLabs remains the *active default*". ElevenLabs code is
   present and live (`internal/jobsimulation/calls/elevenlabs.go`,
   `simulator/manager/elevenLabsManager.go`) and `flag_use_realtime_openai` gates the LiveKit path
   (`calls/livekit.go:132-142`), which is consistent — but "active default" is a PostHog rollout state
   I cannot read.
