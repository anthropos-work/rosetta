# Seat A — M257x iter-49, ninth clause-5 reading

**Corpus repo**: `/Users/marco/workspace/anthropos/rosetta` @ `2fc633a2c5c09a6034e5ab4e29d509dfcadcbd8a`, branch `m257x/platform-realignment`

**Ground-truth clones consulted (all read directly, never assumed):**

| clone | sha |
|---|---|
| `stack-demo/app` | `5ba17044` |
| `stack-demo/app/studio` (anthropos-studio-room) | `aeec036` |
| `stack-demo/platform` | `2adcf71` |
| `stack-demo/graphql-wundergraph` | `60c229f` |
| `stack-demo/next-web-app` | `bb3313bc0` |
| `stack-demo/sentinel` | `88bc559` |
| `stack-demo/storage` | `4ce8ece` |
| `stack-demo/messenger` | `fa47850` |
| `stack-demo/studio-desk` | `14a5442` |
| `.agentspace/rosetta-extensions` | `4d03b53` |
| `.agentspace/snapshots/taxonomy/5afc0bcc…/manifest.json` | (read, parsed) |

---

## Coverage

| # | file | `wc -l` | lines read |
|---|---|---|---|
| 1 | `corpus/architecture/external_services.md` | 814 | 814 (all) |
| 2 | `corpus/services/backend.md` | 271 | 271 (all) |
| 3 | `corpus/services/graphql-wundergraph.md` | 265 | 265 (all) |
| 4 | `corpus/services/academy-backend.md` | 141 | 141 (all) |
| 5 | `corpus/services/coursebuilder.md` | 139 | 139 (all) |
| 6 | `corpus/services/skiller.md` | 66 | 66 (all) |
| 7 | `corpus/services/TEMPLATE.md` | 46 | 46 (all) |
| | **total** | **1742** | **1742** |

Every file was read end to end via `Read` (line-numbered 1..N), not sampled.

---

## BLOCKERS

| # | site (file:line) | the false claim | what is true (platform file:line) |
|---|---|---|---|
| A1 | `corpus/architecture/external_services.md:668` | "**The EU agent is the bare `anthropos-agent`** (`calls/livekit.go:110,120`); **only the US one is suffixed**, `anthropos-agent-us` (`:126`)." — an exclusivity claim that the US variant is the *only* suffixed LiveKit agent name. | There is a **third, suffixed** agent name. `app/internal/jobsimulation/calls/livekit.go:115` — `agentName = "anthropos-agent-chain"` — is taken whenever `CreateAgentDispatch` is called with `chain == true` (the livekitchain voice engine), and it is selected **before** the eu/us location branch. Full census of agent names in the platform (`grep -rn anthropos-agent stack-demo/app/internal/`): `:110` `anthropos-agent`, `:115` `anthropos-agent-chain`, `:120` `anthropos-agent`, `:126` `fmt.Sprintf("anthropos-agent-%s", *location)`, `:142` `anthropos-agent`. A reader provisioning LiveKit agent dispatch from this sentence registers two names and silently loses every chain-engine call. **The rest of the correction stands and is verified**: `anthropos-agent-eu` really does appear nowhere (`grep -rn "anthropos-agent-eu" stack-demo/app` → exit 1, 0 hits; positive control: the unsuffixed grep above returns 5 lines), and the eu/us split really does live on the endpoint (`euAgentEndpoints = {azure-eu, azure-eu-fr}` at `:101-104`, `agentEndpoint = fmt.Sprintf("azure-%s", *location)` at `:127`). Only the word **"only"** is wrong. |

**Blocker count: 1.**

---

## MINORS

1. **`backend.md:253`** — "…not in the top-level `migrations/` dir (**which holds only `atlas.sum`**)". There is **no** top-level `migrations/` directory in `app` @ `5ba17044`: `ls migrations` → *No such file or directory*; `find . -maxdepth 2 -name migrations -type d` returns `./terraform/migrations` and nothing else. Not booked as a blocker because the operative instruction — versioned Atlas migrations live in `terraform/migrations/`, per `atlas.hcl` `dir = "file://terraform/migrations"` — is **correct and verified**, so no reader action is misdirected. The parenthetical just describes a directory that no longer exists.

2. **`backend.md:39`** — anchor drift. "the mux registers five handlers unconditionally (`main.go:1178-1218`)". The five unconditional registrations are Users (`:1178`), Organizations (`:1179`), Skiller (`:1187-1189`), JobSimulation (`:1195`) and LabSession — but LabSession's `mux.Handle` is at **`:1219`**, one line past the cited range (`:1218` is the `NewLabSessionServiceHandler(...)` call that produces the path/handler pair). The *count* (five) and the conditional-`CMSService` anchor (`:1203-1205`, exact) are both correct.

3. **`external_services.md:541` and `:573`** — both cite `app/internal/coursebuilder/bedrock.go:106-113` for the key-set→first-party / key-unset→Bedrock switch. The function `newUnderlyingClient` is actually at **`:109-114`** (`:105-108` are its doc comment). ~3 lines of drift. Note `coursebuilder.md:48` cites the same behavior as `:105-114`, which **is** exact — the two docs disagree with each other on the anchor.

4. **`external_services.md:590`** — cites `config_template.ini:30-31` for the Studio-Room OpenAI endpoint. The file is at **`app/studio/configs/config_template.ini`** (the path in the doc omits `configs/`; there is no `config_template.ini` at the studio root). Lines 30-31 of the real file are `OPENAI_API_KEY = sk-aaa` / `OPENAI_ENDPOINT = https://api.openai.com/v1/chat/completions`, so the content claim is exact.

5. **`graphql-wundergraph.md:82`** — self-reference drift. "…while `:174-176` of this same doc already said `localhost:5050` refuses the connection". That sentence is at **`:178`**; `:174-176` is the section heading and the `make up` sentence.

6. **`coursebuilder.md:114`** — "`go test ./internal/web/backend/coursebuilder/...` # **~25** boundary test files". Actual: **32** (`ls internal/web/backend/coursebuilder/*_test.go | wc -l`). The sibling figures check out: `~100 Go files` in `internal/coursebuilder/` → 98; `~55 test files` → 59.

7. **`coursebuilder.md:48`** — "`ssm.tf:328-334`". The `aws_ssm_parameter "anthropic_api_key"` resource block is `:328-333`. Trivial. Its two companions are exact: `terraform/variables.tf:635-638` (the `sensitive`, no-default declaration) and `terraform/main.tf:555` (`"name": "ANTHROPIC_API_KEY"`).

8. **`external_services.md:376`** — "the six `Subscription` hits in that SDL are all Stripe/plan field names (`activeSubscription`, `stripeSubscriptionId`, `type PlanSubscription`)". The **count is exactly right** (`grep -n Subscription schemas/backend.graphqls` → 6 lines: 1528, 1555, 4385, 4389, 4391, 4418), but the characterization is loose: `:4391` is a **doc-comment** line, not a field name, and `:4389` is `stripeSubscriptionItemId`, which is not one of the three names listed. The load-bearing claim — **no `type Subscription`** in `backend.graphqls`, only `type Mutation` (`:4053`) and `type Query` (`:4912`) — is verified exact.

9. **Observation, not a defect.** `next-web-app/Dockerfile.dev:23` still carries `ARG NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=http://localhost:5050/graphql` as its *default*, and the gateway repo's own `Makefile` `run` target still does `docker run -p 5050:8080`. The corpus's "there is no `:5050` on a local stack" is nonetheless **true as scoped**: `platform/docker-compose.yml` passes the arg explicitly at `:352` (build) and `:361` (runtime), both `http://${PUBLIC_HOST:-localhost}:8082/graphql/query`, so the Dockerfile default never applies on the compose path — and `grep -n 5050 platform/docker-compose.yml` returns nothing, exactly as `graphql-wundergraph.md:80` claims.

---

## What was checked hardest, and came back clean

A blocker count of 1 across 1742 lines is a low yield, so here is the audited-zero evidence for the claims most likely to be wrong. **Every one of these was executed against a clone, not inferred.**

### The router-deletion / supergraph-count spine (the highest-risk surface, ~30 distinct claims)

- **No `graphql` compose service, no `repos.yml` entry** — `grep -n "^  [a-z0-9_-]*:" platform/docker-compose.yml` lists `sentinel:5, backend:28, jobsimulation:83, cms:144, storage:189, customerio-sync:220, messenger:240, roadrunner:281, studio-desk:311, next-web-app:344, gotenberg:371`. No `graphql:`. `repos.yml` (read in full) has 9 entries, none of them `graphql-wundergraph`. ✅
- **The "nine containers on the `graphql` profile" claim (`external_services.md:168-174`)** — `grep -n "profiles:"` shows six services carrying `graphql` (backend `:81`, jobsimulation `:140`, cms `:187`, storage `:218`, roadrunner `:309`, gotenberg `:384`); `sentinel` carries **no** `profiles:` key so it always starts; `common.yml` adds `postgresql` + `redis`, also profile-less. 6 + 1 + 2 = **9**, and the seven named application services are exactly right. ✅
- **The 5→4→3→1 subgraph ladder** — read `supergraph-config-prod.yaml` at each of the five cited trees (`bash -c 'git show "$s:…"'`; note the naive zsh form fails with *bad substitution*, which is why I re-ran under bash rather than reading an error as an empty result): `749dc86~1` = backend/skiller/jobsimulation/cms/skillpath (**5**), `749dc86` (**4**), `7c17e63` (**3**), `915da06~1` (**3**), `915da06` (**1**). `git show --name-status 915da06` marks **both** `schemas/cms.graphqls` and `schemas/jobsimulation.graphqls` as `D`. ✅ The corpus's insistence that **cms-in-app was 3→1, not 2→1** is correct, and `915da06`'s own subject line does read "supergraph 2→1" — the corpus is right to distrust it.
- **The tempting refutation that isn't** — `508ea37` ("remove the jobsimulation subgraph (supergraph 3->2)") *does* delete `schemas/jobsimulation.graphqls`, which would break the corpus's story. It is **not on mainline**: `git merge-base --is-ancestor 508ea37 HEAD` → rc **1**, `git branch -a --contains` names only `remotes/origin/feat/cms-in-app`. The corpus's account survives. ✅
- **No subscriptions** — `grep -rn "sse\|subscription" *.yaml` → exit 1 (0 hits) with the stated positive control (`grep -rln backend *.yaml` → all three files). `915da06~1:supergraph-config-prod.yaml` still reads `protocol: "sse_post"`. `git log -S 'protocol: "ws"' --all` returns **only** `bba862f`, which is not an ancestor of HEAD (rc 1) and lives only on `remotes/origin/feat/use-web-socket`. ✅ Every clause verified.
- **The Dockerfile-era table** — `2c85211^:docker-compose.yml` has `context: ../graphql-wundergraph` with **no `dockerfile:` key** (so Docker defaults to the production `Dockerfile`); `2c85211` adds `dockerfile: Dockerfile.dev`; `1e8e754` line 8 reads `dockerfile: graphql-wundergraph/Dockerfile.dev`; `b56d731:docker-compose.yml` still has `  graphql:` at **`:22`** under `profiles: [wundergraph-deprecated]`; `360efd4` has no such key, `  backend:` at **`:28`**. Every cited line number is exact. ✅
- **Prod-side still declared** — `terraform/main.tf:20` `service_desired_count = 1`, `locals.tf:8` `port = 8080`, `main.tf:48-49` containerPort/hostPort, `config.prod.yaml:5` `listen_addr: 0.0.0.0:8080`, `playground_enabled: false` + `introspection_enabled: false` at `:10`/`:12` (vs `true`/`true` in compose/dev). All exact. ✅
- **`ci/update-subgraph.sh:9`** — exactly one `gh release download`, `-R anthropos-work/app`. The iter-49 repair to `graphql-wundergraph.md:171` is correct. ✅
- **`Dockerfile.dev`** — one schema `COPY` at `:18`, the orphaned comment at `:19-20`, one `awk` at `:23`, `wgc@0.104.0` at `:7`, `node:22.11-alpine` at `:2`, `router:0.275.0` at `:29`. Every anchor in `external_services.md:433-447` exact. ✅
- **`package.json`** is literally `{"name":"graphql-wundegraph"}` and `CLAUDE.md`'s heading carries the same misspelling. ✅
- **The `graphql-wundergraph/CLAUDE.md:39` counter-claim** the corpus warns about is real and still there ("Since cms-in-app the platform compose `graphql` service builds from the **production** Dockerfile"), written at `60c229f` (2026-07-30, *"correct the compose build path"*). The corpus's fence — *the compose file wins, and it lives in `platform`* — is the right call. ✅

### The Directus / cms-in-app posture

- No `directus` service in the platform compose ✅. `backend`'s `environment:` block is `:43-67` and carries **no** `DIRECTUS_*` ✅ (verified line-by-line); `cms`'s `DIRECTUS_BASE_ADDR`/`DIRECTUS_PUBLIC_BASE_ADDR` are at **`:164-165`** exactly ✅. `app/cms_reader_switch.go` exists ✅. `app/main.go:971-973` is exactly the `DIRECTUS_BASE_ADDR` `log.Fatalf` ✅. `app/internal/cms/directus/` exists ✅.
- The rext half: `stack-injection/gen_injected_override.py:53` is `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")` ✅, the re-point is at `:636-637` ✅, `test_injection.py:1051` is `test_backend_the_actual_reader_is_repointed` ✅, and `test_only_cms_is_repointed_not_other_services` is **gone** ✅. The whole iter-24 narrative checks out.
- Frontend endpoints: studio-desk `:318` (build arg) / `:334` (env), next-web-app `:352` / `:361` — **all four** `…:8082/graphql/query` ✅.

### `app` internals

- `SkillPathSessionService`: **0** occurrences in Go, no `skillpath…v1connect` import ✅ — and `app/CLAUDE.md:72` + `app/knowledge/architecture.md:28` *do* both still list it, exactly as `backend.md:43` warns. The Trap-C framing is correct.
- `main.go:735-738` `LABS_API_URL` conditional ✅ exact. `main.go:762` `coursebuilder.ModelBackendName()` boot log ✅ exact. `main.go:816-819` `OPENAI_KEY` cover generator ✅ exact. `main.go:1234` `rpc.WithWriteTimeout(60*time.Second)` ✅. `docker-compose.yml:255-265` messenger's four RPC addrs (two at `backend:8083`, `CMS_RPC_ADDR=http://cms:8091`, `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401`) ✅ exact. `main.go:1196-1202` really does say *"Additive + DORMANT … until the M809 re-point"* ✅.
- The five Redis streams: `AddSubscriber` for `SKILLPATH_STREAM`, `SKILLER_STREAM`, `JOBSIMULATION_STREAM`, `CMS_STREAM`, `AI_USAGE_STREAM`, plus `subServer.AddSubscriber(serviceName, backendSelfSub)` at `:1311` (serviceName = `backend`) ✅ — six subscribers, five application streams + the AI usage stream, exactly as claimed.
- **13** `ai_readiness_*` Ent schema files ✅, and the four the corpus names as the ones a "9" omits (`recommendations`, `email_overrides`, `notification_logs`, `notification_optouts`) are all present, with `live_snapshots` + `text_translations` correctly identified as part of the original nine.
- `POST /api/webhook/directus` skips Clerk authn and authenticates on `DIRECTUS_WEBHOOK_SECRET` (`backend.go:137-138`, `main.go:1076-1080`) ✅ fail-closed. `backend.go:130` is `"/api/webhook/clerk"` ✅ exact.
- Ports: compose publishes 8081/8082/8083; env sets `PORT=8082`, `RPC_PORT=8083`, `META_PORT=8084` ✅ — `backend.md:115` exact, including "8081 is reserved/unused".

### The taxonomy figures (the claim most often mis-transcribed)

Parsed the manifest directly: `public.skills` **42,790**, `public.job_roles` **22,470**, `public.skill_embeddings` **42,790**, `public.job_role_embeddings` **18,919**, `categories` 23, `specializations` 1,447, `public_only: true`, `predicate: org-null`, `captured_at 2026-06-29`. Every number in `external_services.md:632-645`, `backend.md:9/36/71-74` and `skiller.md:40-49` matches, and the "18,919 was transcribed onto the role count" diagnosis is arithmetically sound. `backend.md`'s 2026-07-08 prod figures are internally consistent too (42,790 public + 794 org-private = 43,584 total = the `skill_embeddings` total). ✅

### The residency argument (`external_services.md:543-600`) — the newest text, checked clause by clause

Every anchor in `internal/jobsimulation/ai/ai.go` is **exact**: vendor consts `:30-33`, `getClient` `:259-289`, the Azure/PostHog flag block `:264-276`, `AnthropicAws`+`Anthropic` → `anthropicClient` `:280-283`, `isThrottlingError` `:130-141`, the ChatCompletion throttle override `:150-155`, the `Response` mirror `:296-302` and `:326`, `openai.NewOpenAI(openaiKey)` at `:80`. The cms-layer default is exact: `AIVendor *AIVendor` at `collections/jobsimulation.go:905`, `aiVendor := simulation.Openai` at `:1302-1305`. The mapping is exact: `simulator/ai/ai.go:58-59` and the `default:` arm at `:114-115`. Mistral's *only* uses in `app` are `cms/studio/markdownManager.go:19` and `studioManager.go:583` ✅. Item 5 (the iter-49 addition) is exact: `studio/services/ai.py:704-724` providers dict, `:383` bare `OpenAI(api_key=self.api_key)`, `:627-664` `AnthropicProvider`, and **0 hits for `bedrock|boto3` anywhere under `app/studio/`** (grep exit 1) — so "Studio-Room was never on Bedrock" is right. The coursebuilder prod-key argument is right too: `variables.tf:635-638` sensitive/no-default, `ssm.tf:328` parameter, `main.tf:555` injection.

### Course Builder / Academy

- `creditCost` map at `internal/credits/cost.go:86-90` is `build:5, refine:1, translate:1` — **`coursebuilder.md:102`'s "`course.refine`=1" is correct**, and the *platform's own* package doc comment at `cost.go:29` ("5 credits per refine turn") is the stale one. This was my best blocker candidate and it resolved in the corpus's favour. ✅
- `DefaultSessionsPerOrgPerDay = 50` (`rate.go:28`), `DefaultMaxMonthlyCOGSUSD = 500.0` (`budget.go:27`), `DefaultMaxDailyCOGSUSD = 0.0` (`:41`) ✅. Migration `20260717151144.sql` creates exactly `course_builder_sessions`, `credit_transactions`, `organization_credits` ✅. Route table at `handler.go:2859-2926` matches the doc's list ✅. `/coursebuilder` Echo group + org-admin gate at `backend.go:229` ✅.
- `COURSEBUILDER_OPENAI_IMAGE_KEY` — **0** Go references; the only survivals are `internal/coursebuilder/GO-LIVE-RUNBOOK.md:35` and `README.md:429`, i.e. exactly the "stale in-repo markdown" the doc names; and `git show 68c24512` does remove `os.Getenv("COURSEBUILDER_OPENAI_IMAGE_KEY")`. The whole warning is precisely right. ✅
- Academy: all 11 `academy_*` tables exist with the **plural** names the doc insists on (`academy_chapter_progresses`, `academy_certificates`, `academy_feedbacks`, `academy_last_activities`, …) ✅ — that "Ent LABELS vs plural TABLE names" warning is real and useful. `academy_embedding_refresh` Asynq task exists (`worker/tasks/academy_embedding_refresh.go:10`, "03:00" per `embeddings.go:108`) ✅. `/content/catalog.json` (`content.go:23`) + `/content/admin` shared-token group ✅. **No Connect-RPC for academy** ✅. `internal/aiacademy` is **gone** and `internal/academy/academy.go:6-9` says exactly what the doc quotes ✅. PR `0e37771f` is real, dated 2026-06-05, titled *"Academy backend v1.0 'ground truth'"* ✅. App version `v1.363.2` @ `5ba17044` ✅ (`CHANGELOG.md:5`).

### Cross-cutting spot checks

Sentinel really has **zero** clerk/authn references and **zero** webhook routes, and depends on `casbin/v3` — so both the "authorization-only, no Clerk validation" claim (`:73-76`) and the "webhooks go to backend, not Sentinel" claim (`:736-737`) hold ✅. Clerk SDKs per app all match (`@clerk/nextjs` web/hiring/integration, `@clerk/clerk-expo` mobile, `@clerk/clerk-js`+`@clerk/express` studio-desk, `clerk-sdk-go/v2 v2.7.0`) ✅. `NEXT_PUBLIC_GRAPHQL_ENDPOINT` has **0** occurrences in `next-web-app` ✅. `enum ContentLanguage` has exactly **8** values ✅ (skiller.md:55). Studio-desk's `AI_PROVIDER_CHAIN` default is `azure-openai,openai` per its own `.env.example:57` ✅, and the circuit-breaker rotation is real (`aiService.ts:10`) ✅. `gqlauthz.go:186` is exactly the `unknown viewer: Forbidden` return, and both named regression tests exist (`gqlauthz_test.go:132`, `:176`) ✅. `make up` default profile is `graphql` (`platform/Makefile:10`) ✅.

`TEMPLATE.md` (46 lines) makes **no** claims about the platform — it is a pure skeleton with bracketed placeholders. Nothing to falsify; audited as a genuine zero.

### Failed / empty commands, disclosed

- `git show "$s:supergraph-config-prod.yaml"` under zsh returned *"(eval):1: bad substitution"* for all five refs — a **failed command**, not evidence. Re-run under `bash -c` and reported only the bash results.
- `grep -rn "anthropos-agent" --include=*.go .` failed under zsh (*"no matches found"* — glob expansion, not a grep result). Re-run without the glob; the 5-line result above is from the successful run.
- `grep -rniE "bedrock|boto3" app/studio/` → **exit 1 with no error output** = a genuine 0 hits, distinguishable from the failures above.
