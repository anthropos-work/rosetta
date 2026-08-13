# Seat A — iter-48

Repo: `/Users/marco/workspace/anthropos/rosetta` @ `m257x/platform-realignment` `cabc3b15`.
Ground truth read directly from:
- `stack-demo/app` @ `5ba17044` (v1.363.2, 2026-07-31)
- `stack-demo/platform` @ `2adcf714` (merge #23 "chore/drop-wundergraph", 2026-07-31)
- `stack-demo/graphql-wundergraph` @ `60c229f3` (2026-07-30)
- `stack-demo/messenger` @ HEAD, `stack-demo/studio-desk`, `stack-demo/next-web-app` @ `bb3313bc`
- `.agentspace/rosetta-extensions` (stack-injection)

## Coverage (file, wc -l, lines read)

| # | file | `wc -l` | lines read |
|---|---|---|---|
| 1 | `corpus/architecture/external_services.md` | 808 | **all 808** |
| 2 | `corpus/services/backend.md` | 271 | **all 271** |
| 3 | `corpus/services/graphql-wundergraph.md` | 265 | **all 265** |
| 4 | `corpus/services/academy-backend.md` | 141 | **all 141** |
| 5 | `corpus/services/coursebuilder.md` | 139 | **all 139** |
| 6 | `corpus/services/skiller.md` | 66 | **all 66** |
| 7 | `corpus/services/TEMPLATE.md` | 46 | **all 46** |

Total 1736 lines, read top-to-bottom in full (no sampling).

### Search-hygiene notes (rule 1 + 2)
Two searches in this pass returned **rc=0 with empty output** for reasons that had nothing to do
with the corpus, and both were caught only because stderr was read:
- `grep -rn X --include=*.go .` → zsh expanded `*.go` before grep saw it → `(eval):1: no matches found`.
  Re-run quoted (`--include="*.go"`) with a positive control (`SkillerService` → 4 files).
- `git show 749dc86^:file` / `git show "$c:file"` → zsh read `:s…` as a parameter **modifier** →
  `bad substitution`, which looks exactly like "commit has no such file". Re-run as `"${c}":file`.
Every negative result below is paired with a positive control that matched.

### What was verified GREEN (the bulk of both fenced sections)
Recorded so the next iteration doesn't re-spend budget on it.

*Router / supergraph.* `2adcf71` is an ancestor of platform HEAD; `grep -c 5050 docker-compose.yml` = **0**;
no `graphql` service key and no `graphql`/`wundergraph` entry in `repos.yml` (positive control: `app` at
`repos.yml:10`); `graphql` survives only as a profile label. `terraform/main.tf:20 service_desired_count = 1`,
`terraform/locals.tf:8 port = 8080`, `terraform/main.tf:48-49` containerPort/hostPort `${local.port}`,
`config.prod.yaml:5 listen_addr: 0.0.0.0:8080` — all exact.
Subgraph ladder, read off `supergraph-config-prod.yaml` per commit: `749dc86~1` = **5**, `749dc86` = **4**,
`7c17e63` = **3**, `915da06~1` = **3**, `915da06` = **1** — so **3 → 1**, and `915da06`'s own subject
("supergraph 2→1") is wrong exactly as both docs say. `git show --name-status 915da06` marks
`schemas/cms.graphqls` **and** `schemas/jobsimulation.graphqls` `D`. `subgraphs.conf` = `BACKEND=v1.360.0` alone.
Subscriptions: `grep -rn "sse\|subscription" ./*.yaml` → no match (positive control: `backend` matches all
three); `git log -S 'protocol: "ws"'` on mainline → **nothing**, `-S sse_post` → exactly **5** commits;
`bba862f` is **not** an ancestor of HEAD and lives only on `remotes/origin/feat/use-web-socket`.
(Also checked and *not* a refutation: `508ea37` "remove the jobsimulation subgraph (3->2)" exists but is
only on `remotes/origin/feat/cms-in-app` — it never landed separately, so the "the jobsim entry outlived
jobsim-in-app" claim survives.) `schemas/backend.graphqls`: `type Mutation` at **:4053**, `type Query` at
**:4912**, **no** `type Subscription`, and exactly **6** `Subscription` substring hits, all Stripe/plan.
`Dockerfile.dev`: one schema `COPY` at **:18**, the orphaned comment at **:19-20**, one `awk` at **:23**,
`wgc@0.104.0` at `:7`, `router:0.275.0` at `:29`. `package.json` = `{"name":"graphql-wundegraph"}` (misspelt).
Compose build-path history: `63d285c` (2024-06-20) introduces the service as `wundergraph:` with a `git@…`
context and **no `dockerfile:` key**; `d92e84e` renames it to `graphql:`; `a2a3ee6` moves the context to
`../graphql-wundergraph` still with no key; `719befb` still has no key; `2c85211` adds
`dockerfile: Dockerfile.dev`; `67ba772` raises the context to `..`; `b56d731` still has `  graphql:` at `:22`
behind `profiles: [wundergraph-deprecated]`; `360efd4` has no such key (positive control: `  backend:` at `:28`).
`1e8e754` lines 6-8 read `context: ..` / `dockerfile: graphql-wundergraph/Dockerfile.dev`. All as documented.

*Directus / compose posture.* `backend`'s compose `environment:` block is `:43-67` and carries **no**
`DIRECTUS_*`; the only explicit setter is `cms` at `:164-165`. Frontend endpoints at `:318`/`:334`
(studio-desk) and `:352`/`:361` (next-web-app), all four `:8082/graphql/query`.
`app/cms_reader_switch.go:28-29` reads verbatim *"a DIRECT domain call — no proto round-trip … and no
internal traffic to a standalone cms"*; `app/main.go:971-973` `log.Fatalf`s without `DIRECTUS_BASE_ADDR`.
Nine-container `graphql` profile confirmed: `postgresql` + `redis` (profile-less in `common.yml`) + the
profile-less `sentinel` + `backend`/`jobsimulation`/`cms`/`storage`/`roadrunner`/`gotenberg`.
The historical Directus service is real: `git show a2a3ee6~1:docker-compose.yml` → `:384 image:
directus/directus:10.10.1`, `:386 - 8055:8055`, `:409 - ADMIN_PASSWORD=password`; `a2a3ee6` itself has zero
`directus` hits. rext: `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")` at
`stack-injection/gen_injected_override.py:53`, consumed at `:636-637`; the replacement test
`test_backend_the_actual_reader_is_repointed` is at `stack-injection/tests/test_injection.py:1051` and the
old inverted test is gone (only its name survives, in that test's own comment).

*AI routing.* Every anchor in `external_services.md:534-594` checked line-by-line and correct:
vendor consts `internal/jobsimulation/ai/ai.go:30-33`; `getClient` `:259-289`; Azure/`flag_use_azure_us`
`:264-276` with error → keep-EU; `AnthropicAws`+`Anthropic` → `a.anthropicClient` `:280-283`;
`isThrottlingError` 429-only `:130-141`; `throttled → vendor = Openai` `:150-155` and `:296-302`/`:326`;
`openai.NewOpenAI(openaiKey)` at `:80`. Caller default: `internal/cms/directus/collections/jobsimulation.go:905`
`AIVendor *AIVendor` and `:1302` `aiVendor := simulation.Openai`; mapped at
`internal/jobsimulation/simulator/ai/ai.go:58-59`, `default:` arm `:114-115`. Mistral is OCR-only, cms-only
(`internal/cms/studio/markdownManager.go:19`, `studioManager.go:583`; 5 total `mistral` hits in `internal/`).
Model ids `coursebuilder/bedrock.go:23,29`, `askengine/bedrock.go:25`, `jobsimulation/agent/report_agent.go:31`
all exact. `app/studio/services/ai.py:627` `class AnthropicProvider` exists; `grep -riE 'bedrock|boto3' studio/`
→ **no match** (positive control: `anthropic` matches `studio/gen.py`, `requirements.txt`, …), so the
iter-48 correction *"Studio-Room was never on Bedrock"* holds.

*`app` internals.* RPC mux: `Users`(:1178) `Organizations`(:1179) `Skiller`(:1187) `JobSimulation`(:1195)
unconditional, `CMSService` behind `if cmsRPCServer != nil` at `:1203-1205` with the quoted
"Additive + DORMANT … until the M809 re-point" comment verbatim at `:1196-1202`; 60 s write timeout at
`main.go:1234`. `SkillPathSessionService`: **0** hits in Go source (positive control: `SkillerService` → 4
files) — it survives only in `app/CLAUDE.md:72` and `app/knowledge/architecture.md:28`, exactly as
`backend.md:43` says. `labsAPI` conditional at `main.go:735-738`; `internal/labs/{session,labsapi,adapter,catalog}`
all present. Ports: compose publishes 8081/8082/8083, env `PORT=8082`/`RPC_PORT=8083`/`META_PORT=8084`.
Messenger's four addresses at `docker-compose.yml:255-265`: `BACKEND_USERS_RPC_ADDR` + `SKILLER_RPC_ADDR`
→ `backend:8083`, `CMS_RPC_ADDR` → `cms:8091`, `JOBSIMULATION_RPC_ADDR` → `jobsimulation:8401` — exactly the
split `backend.md:195` describes. Redis streams: `AddSubscriber` for SKILLPATH(:1265) SKILLER(:1267)
JOBSIMULATION(:1276) CMS(:1294) AI_USAGE(:1296) + self/`serviceName`(:1311), with `.AddHandler` merges at
:1279/:1287/:1309. **13** `ai_readiness_*` tables in `terraform/migrations` (counted excluding index/FK
names) — `backend.md:185`'s "13" and its four-named-omissions list are both right; `internal/workforce/`
contains no `readi*` file (positive control: `ls internal/workforce/` returns activity.go, growth.go, …).
`internal/rpc/skillerrpc/skiller.go` implements the 5 externally-reached methods with real bodies
(:73/:93/:116/:147/:187); the other 8 are stubs.

*Academy.* `internal/academy/academy.go:6-9` states the `aiacademy` removal verbatim. All 11 Ent schema files
present; migration table names confirm the **plural** forms (`academy_chapter_progresses`,
`academy_last_activities`, `academy_chapter_times`, `academy_certificates`, `academy_bookmarks`,
`academy_feedbacks`, `academy_series`, `academy_skill_paths`, `academy_chapters`, `academy_chapter_bodies`,
`academy_path_embeddings`). Nightly task `academy_embedding_refresh` registered `"0 3 * * *"` at
`internal/worker/worker.go:158` → 03:00 ✓. `cmd/academy-seed` flags `--user-email/--user-id/--fixture/
--reset/--dry-run/--list` exact. `ACADEMY_CONTENT_API_TOKEN` shared-token middleware at
`internal/web/backend/content_admin.go:35`; public `/content/catalog.json` at `internal/web/backend/content.go`.
Certificate: `AAS-YYYY-XXXXXX`, `crypto/rand`, regex `^AAS-\d{4}-[A-HJKMNP-TV-Z2-9]{6}$` = a 30-symbol
ambiguity-free alphabet. `0e37771f` = PR #903, 2026-06-05, "Academy backend v1.0 'ground truth'". App
version `v1.363.2` @ `5ba17044` ✓.

*Course Builder.* `bedrock.go:98-104 ModelBackendName()`, `:105-114 newUnderlyingClient`, logged at
`main.go:762`; `terraform/variables.tf:635-638` (sensitive, no default), `ssm.tf:328-334`, `main.tf:555`
inject `ANTHROPIC_API_KEY` — the "production path is the first-party API" claim is fully supported.
`terraform/migrations/20260717151144.sql` creates `course_builder_sessions` (:2) + `credit_transactions` (:40)
+ `organization_credits`. `DefaultMaxMonthlyCOGSUSD = 500.0`, `DefaultMaxDailyCOGSUSD = 0.0`,
`DefaultSessionsPerOrgPerDay = 50`, `coursebuilderWorkerConcurrency = 20` (`main.go:172`),
`Retention(6h)`. Credits map `internal/credits/cost.go:86-90`: build **5**, refine **1**, translate **1**.
`OPENAI_KEY` at `main.go:816-819`; `COURSEBUILDER_OPENAI_IMAGE_KEY` exists only in two stale in-repo
markdown files — exactly as the doc says. Waves top out at 24. `handler.go:2842 func (h *Handler) Register`.
`gqlauthz.go:186` is the `unknown viewer: Forbidden` return, and `gqlauthz_test.go:132` is
`"bare __typename is not exempt"` with `TestAnonymousRejectionLogsAtWarn` at `:176`.
`NEXT_PUBLIC_GRAPHQL_ENDPOINT` does **not** exist in `next-web-app` (positive control:
`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` matches 5+ files).

*TEMPLATE.md* makes no falsifiable claim about the platform — it is a skeleton. Nothing found.

## Blockers

| # | site | the false claim (verbatim) | what is TRUE | citation |
|---|---|---|---|---|
| 1 | `corpus/services/graphql-wundergraph.md:171` | "**CI/prod**: GitHub Releases on `anthropos-work/{app,jobsimulation,cms}` (schema artifacts) + `anthropos-work/infrastructure` Terraform + `release-service.yml`." | CI pulls a schema artifact from **`anthropos-work/app` only**. `ci/update-subgraph.sh` is 9 lines and has exactly **one** `gh release download`, targeting `-R anthropos-work/app`; the `jobsimulation` and `cms` downloads were deleted at `915da06`. The bullet is also contradicted by this same doc's `:50` ("`subgraphs.conf` carries a single `BACKEND=` pin") and `:96` ("`schemas/` … now just backend.graphqls"), and it sits under a plain `## Dependencies` heading with **no** historical fence — unlike the two bullets directly above it (`:169`, `:170`), which do say "historically". A reader wiring or debugging the supergraph-update path is sent to two repos that no longer publish into it. | `stack-demo/graphql-wundergraph` @ `60c229f3`: `ci/update-subgraph.sh:9` (sole download line, `-R anthropos-work/app`); `subgraphs.conf` = `BACKEND=v1.360.0`; `git show 915da06` commit body: *"…dropped … the gh-release-download in ci/update-subgraph.sh…"*, and `git show --name-status 915da06` marks `ci/update-subgraph.sh` `M` |
| 2 | `corpus/architecture/external_services.md:662` | "The platform runs **GPT Realtime agents** (`anthropos-agent-eu` / `anthropos-agent-us`) inside LiveKit rooms" | **`anthropos-agent-eu` does not exist anywhere in the platform.** The EU agent name is the bare **`anthropos-agent`** (`livekit.go:110` default, re-asserted at `:120` inside the `LocationEu` branch); only the non-EU branch builds a suffixed name, `fmt.Sprintf("anthropos-agent-%s", *location)` → `anthropos-agent-us` (`:126`). The eu/us split the doc has attached to the *agent name* actually lives on the **endpoint** — `azure-eu` (`:111`) vs `azure-us` (`:127`), plus `openai-hosted` when `flag_use_realtime_openai` is on (`:143`). There is also a third agent the doc omits, `anthropos-agent-chain` (`:115`). `grep -rn "anthropos-agent" --include="*.go" internal/jobsimulation/` returns 7 lines, none containing `-eu`. | `stack-demo/app` @ `5ba17044`: `internal/jobsimulation/calls/livekit.go:110,111,115,120,126,127,143` |

## Minors

1. **`corpus/services/backend.md:39`** — anchor short by one line. "the mux registers five handlers unconditionally (`main.go:1178-1218`)". The claim is true (Users/Organizations/Skiller/JobSimulation/LabSession are all unconditional, CMSService is not), but the fifth registration is `mux.Handle(labSessionPath, authn.HTTPAuthnMiddleware(labAuthn)(labSessionHandler))` at **`main.go:1219`**; `:1218` is only the `NewLabSessionServiceHandler` call that produces the pair. Range should read `:1178-1219`.

2. **`corpus/services/backend.md:253`** — describes a directory that no longer exists. "…not in the top-level `migrations/` dir (**which holds only `atlas.sum`**)". At `app@5ba17044` there is no top-level `migrations/` at all — `ls migrations/` → *No such file or directory*, `git ls-files migrations/` → empty (positive control: `git ls-files terraform/migrations/` → 169 files), and `atlas.sum` lives at `terraform/migrations/atlas.sum`. The operative claim (migrations live in `terraform/migrations`, per `atlas.hcl` `dir = "file://terraform/migrations"`, `src = "ent://internal/data/ent/schema"`) is exactly right — only the parenthetical is stale, and it steers *away* from the wrong place, so it misleads nobody.

3. **`corpus/architecture/external_services.md:429`** — quotation cited to the wrong line. The quoted sentence *"Since cms-in-app the platform compose `graphql` service builds from the **production** Dockerfile"* is at `graphql-wundergraph/CLAUDE.md:**33**` (the "Platform compose" row of the Schema-Management table), not `:39`. Line 39 asserts the same wrong thing in different words ("Since the compose stack builds from the production Dockerfile, a schema change in `app` does not appear at `:5050`…"), so the *substance* of the corpus's correction stands at both lines; only the anchor + the quote-marks pairing drifted.

4. **`corpus/architecture/external_services.md:144`** — anchor off by three. "repaired at [`service_taxonomy.md:296-303`]". The repair paragraph ("*That retraction over-corrected…*") actually runs `service_taxonomy.md:**299-305**`; `:296-303` starts inside the *retraction* it is correcting. Note also that the twin there is stamped **M257x iter-46** while this one is stamped iter-48 — consistent, not a conflict, but worth knowing when re-reading the pair.

5. **`corpus/services/coursebuilder.md:114`** — count drift. "`go test ./internal/web/backend/coursebuilder/...` # **~25** boundary test files"; actual `ls internal/web/backend/coursebuilder/*_test.go | wc -l` = **32**. (The two sibling approximations are fine: "~100 Go files" → 98, "~55 test files" → 59.)

6. **`corpus/services/graphql-wundergraph.md:82`** — self-reference off by three. "while `:174-176` of this same doc already said `localhost:5050` refuses the connection"; that sentence is at `:177-179` (`:175` is the heading).

7. **`corpus/services/graphql-wundergraph.md:171`, tail** — `release-service.yml` is named as a CI artifact but is not in this repo's `.github/workflows/` (which holds `bump-version.yml`, `release.yml`, `supergraph-update.yml`). If it means the `anthropos-work/infrastructure` workflow, that repo is not cloned here — **unverified, not refuted**; recorded so it isn't re-chased. (The rest of that bullet is Blocker #1.)

8. **Cross-corpus tension, owner is another seat.** `corpus/architecture/service_taxonomy.md:290-292` still frames the reader as `cms`: *"`cms` reaches Directus over the network via `DIRECTUS_BASE_ADDR` / `DIRECTUS_PUBLIC_BASE_ADDR` env vars (the only service the compose gives these)"*. `external_services.md:198-204` (my file) explicitly corrects this — `backend` is the actual reader (`app/cms_reader_switch.go:28-29`, `app/main.go:971-973`) and the compose-env observation is a *separate* fact about which service gets the var **explicitly**. **My file is the correct one**; flagging so the seat that owns `service_taxonomy.md` reconciles it rather than "re-correcting" `external_services.md` from it.

9. **Platform-side trap, recorded so a future auditor does not "fix" the corpus from it.** `app/internal/credits/cost.go:29` (package doc comment) says *"course.refine → **5** credits per refine turn"*, but the authoritative `creditCost` map at `:86-90` is `ActionCourseRefine: 1`, and the const doc at `:59-67` spells out the D1 ruling ("a refine turn is a FLAT 1 CREDIT"). `coursebuilder.md:102`'s `course.refine`=**1** is **correct** — the platform's own comment is the stale one. This is a live instance of Trap C (`platform-alignment.md`): grepping to `cost.go` and reading the header comment would produce a confidently wrong "blocker".
