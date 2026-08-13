# iter33 — Clause-5 KB-fidelity audit, group G5 (18 files, 1897 lines)

Auditor: full top-to-bottom read of every assigned file. Platform ground truth read from the clone at
`/Users/marco/workspace/anthropos/rosetta/stack-demo/platform` @ `2adcf71` and the peer repo clones
(`app` @ `5ba17044` = v1.363.2, `next-web-app` @ `bb3313bc0`, `sentinel`, `storage`, `messenger`,
`roadrunner`, `cms`, `jobsimulation`). **Read-only — no file was edited.**

---

## 1. Positive control — per-file read confirmation

| File (under `corpus/services/`) | `wc -l` | Read status |
|---|---:|---|
| `storage.md` | 166 | read to line 166 |
| `roadrunner.md` | 164 | read to line 164 |
| `sentinel.md` | 159 | read to line 159 |
| `ai-labs.md` | 156 | read to line 156 |
| `academy-backend.md` | 139 | read to line 139 |
| `coursebuilder.md` | 132 | read to line 132 |
| `messenger.md` | 128 | read to line 128 |
| `clerk-integration.md` | 128 | read to line 128 |
| `next-web-app.md` | 126 | read to line 126 |
| `askengine.md` | 121 | read to line 121 |
| `skillpath.md` | 92 | read to line 92 |
| `gotenberg.md` | 82 | read to line 82 |
| `README.md` | 79 | read to line 79 |
| `customerio-sync.md` | 75 | read to line 75 |
| `skiller.md` | 55 | read to line 55 |
| `TEMPLATE.md` | 46 | read to line 46 |
| `db-backup.md` | 31 | read to line 31 |
| `intelligence.md` | 18 | read to line 18 |
| **total** | **1897** | **18 / 18 read in full** |

---

## 2. Findings

### BLOCKER-1 — `storage.md:5` and `storage.md:104`

> `:5` — "Other services (`jobsimulation`, `cms`, `app`) push and pull binary objects through it instead of dealing with S3 themselves."
> `:104` — "**Upstream consumers**: jobsimulation (recordings, simulation documents), cms (content assets, media), app (user files, profile images)"

**Why false at HEAD.** `jobsimulation` and `cms` are folded into `app` (`repos.yml:14-19` — `migrations: false # legacy — folded into app`); their compose services still start but are unfederated husks off every request path. The only process that calls storage at HEAD is `app`:
- recordings: `app/internal/jobsimulation/recording/recording.go:12` + `anticheat/anticheat.go:34` + `simulator/…` import `storage/sdk/storage/v1`
- cms assets: `app/main.go:983` — `cmsStorage := storage.NewClient(os.Getenv("STORAGE_RPC_ADDR"), storagens.CMS)`
- user files: `app/internal/app/users/services/service.go:9-10`, `internal/publicstorage/publicstorage.go:11`

`storage.md` is the **only** file in this group that never acknowledges the monolith merge anywhere — no banner, no fence, none of the words a term-scoped sweep would grep (`merged`, `husk`, `monolith`, `in-process`, `M810`). Its Dependencies block is the operative "who calls storage" map, and it reads as a three-service world. A reader debugging a missing recording upload on a local stack is sent to the `jobsimulation` container/repo, which is frozen legacy. This is the studio-room archetype shape.

**Grade: BLOCKER** (weakest of the blockers — the names `jobsimulation`/`cms` are also literal `app/internal/` domain package names, so a charitable reading is domain-level; I grade it BLOCKER because the doc frames them as *services* at `:5` and never fences it).

**Suggested correction (one line):** add the standing consumer fence — "**Sole live caller: `app`.** The `jobsimulation` and `cms` domains are in-process inside `app` since the folds; their husk containers still start (teardown M810) but are off every storage path."

---

### minor-2 — `sentinel.md:12`

> "* **Language**: Go 1.25"

**Why false at HEAD.** `sentinel/go.mod:3` reads `go 1.26.0`. (Peer check: `storage`, `roadrunner`, `messenger` really are `1.25.0`, so this is a genuine per-repo drift, not a corpus-wide convention.)

**Grade: minor.** **Correction:** `Go 1.26`.

---

### minor-3 — `sentinel.md:5`

> "…still receive `AUTHORIZATION_ADDRESS=http://sentinel:8087` at `docker-compose.yml:97,158`"

**Why false at HEAD.** Those lines are the two `environment:` keys. The actual `AUTHORIZATION_ADDRESS` rows are `docker-compose.yml:99` (jobsimulation) and `:160` (cms) — both off by exactly +2.

**Grade: minor (wrong line number).** **Correction:** `docker-compose.yml:99,160`.

---

### minor-4 — `sentinel.md:82`

> "* **Upstream consumers**: every other Anthropos service that gates requests (`app`, `cms`, `jobsimulation`, `messenger`)"

**Why misleading at HEAD.** Unfenced restatement of the four-caller world that the doc's own `:5` correctly fences ("Its live callers are **`app`** … and **`messenger`**"). `cms` and `jobsimulation` are husks off every request path. A reader who skips to Dependencies gets the pre-merge picture.

**Grade: minor** (the doc self-corrects 77 lines earlier). **Correction:** "`app` (incl. the folded cms + jobsimulation authz call sites) and `messenger`; the cms/jobsimulation husks still receive the address but call nothing."

---

### minor-5 — `messenger.md:110`

> "skill-path data is read via the CMS client (`internal/flow/assignments.go:815`)."

**Why false at HEAD.** `messenger/internal/flow/assignments.go:815` is inside `getEmailNotificationForSimulation` (`Type: notificationType,`). The CMS skill-path read is `getSkillPath` at `:827-828` (`h.cms.GetSkillPath(ctx, connect.NewRequest(&cmsv1.GetSkillPathRequest{`).

**Grade: minor (wrong line number).** **Correction:** `internal/flow/assignments.go:827-828`.

---

### minor-6 — `messenger.md:23-40` (the "Key directories" block) + `:58-60` ("What triggers Messenger?")

> The block enumerates `internal/flow/` as: `flow.go`, `assignments.go`, `cms.go`, `jobsimulations.go`, `organizations.go`, `organizations_db.go`, `whitelabel.go`.

**Why stale at HEAD.** `internal/flow/` holds 12 non-test files; the doc omits five, all of them net-new flows: `ai_readiness.go` (the v0.41 M408 per-org copy override — six files incl. tests), `content_assigned.go` + `content_completed.go` (the v0.42 unified content-assignment / assigner-completion handlers, visible at `internal/flow/flow.go:85-89`), `coursebuilder.go`, `invitation_reminders.go`. Consequently "What triggers Messenger?" names only jobsimulation/cms/backend streams and never mentions AI-readiness, content-assignment digests, or the Course Builder author-lifecycle emails — even though `coursebuilder.md:81` asserts those emails exist ("colony pub/sub (author-lifecycle emails → messenger → Brevo)"), so the two docs disagree.

**Grade: minor** (incompleteness, not a false statement — but it is a real blind area, and it is cross-doc-inconsistent). **Correction:** add the five files + an "AI-readiness / content-assignment / coursebuilder" row to the trigger list.

---

### minor-7 — `messenger.md:42-44`

> "### Whitelabel rendering (2026-Q2) — Recent work in v0.34.0 added **whitelabel support**"

**Why stale at HEAD.** `messenger/CHANGELOG.md` HEAD is **v0.42.0 (2026-07-24)**; v0.34.x dates to 2026-05/06. The whitelabel feature itself is still present and the `READONLY_DB_CONNECTION` claim at `:44` verifies exactly (`cmd/root.go:147`), so the content is true — only the "recent" framing is eight minors out of date.

**Grade: minor (true-but-confusing).** **Correction:** drop "Recent work in" → "v0.34.0 added whitelabel support (messenger is at v0.42.0 today)".

---

### minor-8 — `README.md:20`

> "And **three of the four** (cms, jobsimulation, roadrunner) still start CONTAINERS locally in the default `graphql` profile as unfederated husks"

**Why false at HEAD.** The doc's own `:11` defines "the four" as *skiller, skillpath, jobsimulation, cms*, and `:15` explicitly names roadrunner "**the fifth**". Roadrunner is therefore not one of "the four". The underlying compose fact is right (`docker-compose.yml:144,83,281`, all `profiles: [graphql, …]`) — only the arithmetic is wrong.

**Grade: minor.** **Correction:** "three of the five".

---

### minor-9 — `README.md:28`

> "…**plus** the folded skiller (taxonomy, matching, embeddings), skillpath, jobsimulation, cms and **roadrunner domains**"

**Why misleading at HEAD.** There is **no `app/internal/roadrunner/`** (verified: absent from `ls app/internal/`). The Judge0 runner was absorbed as `app/internal/jobsimulation/runner/`, constructed at `app/internal/jobsimwiring/wiring.go:118` (`jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))`). The index's own roadrunner row at `:33` states this correctly, so `:28` sends a reader looking for a package that does not exist.

**Grade: minor.** **Correction:** "…jobsimulation (which absorbed the roadrunner Judge0 runner as `internal/jobsimulation/runner/`) and cms domains".

---

### minor-10 — `next-web-app.md:124`

> "* [GraphQL Gateway](./graphql-wundergraph.md) — the federated endpoint this app consumes"

**Why false at HEAD.** Locally the app consumes `backend` directly at `:8082/graphql/query` (`docker-compose.yml:352,361`); `graphql-wundergraph` was deleted from `repos.yml` + `docker-compose.yml` at `2adcf71` and the repo is archived. The doc body handles this correctly three times (`:14`, `:47`, `:53`, `:96`); only the Related-Documentation caption still asserts the router is the endpoint.

**Grade: minor (stale caption / effectively-dead link semantics).** **Correction:** "— the Cosmo Router, **prod-only** since `2adcf71`; locally this app hits `backend` at `:8082/graphql/query`".

---

### minor-11 — `next-web-app.md:11`

> "**Primary Goal**: The main user-facing frontend — a pnpm + Turborepo monorepo of Next.js apps that **consume the federated GraphQL gateway** and authenticate with Clerk."

**Why misleading at HEAD.** The federated *gateway* no longer exists locally. Self-corrected three lines later at `:14`, so the exposure is one summary sentence.

**Grade: minor.** **Correction:** "…that consume the platform GraphQL endpoint (locally `backend:8082/graphql/query`; the Cosmo Router remains only in prod)".

---

## 3. Per-file clean verdicts

Explicit clean verdict — no finding of any grade — for the following 12 files. Each was read in full and its
load-bearing claims spot-verified against source (verification notes below):

* **`roadrunner.md` — CLEAN.** The nuanced prod-vs-local case is handled exactly right. Verified: `roadrunner/terraform/main.tf:19` = `service_desired_count = 1` (vs `cms/terraform/main.tf:39` = 0, `jobsimulation/terraform/main.tf:40` = 0); `repos.yml` really does hold **9** repos with roadrunner among them; `docker-compose.yml:281` starts it with `profiles: [graphql, roadrunner, all]`; `ROADRUNNER_RPC_ADDR=http://roadrunner:10401` is set at `:118` but read by no Go code; `internal/lsp/lsp.go` exists and is unwired; `queues.DefaultQueue = "roadrunner:default"`, `asynq.MaxRetry(3)` (`internal/runner/runner.go:126`), `Concurrency: 10` (`internal/worker/worker.go:25`); asynq `v0.25.1`, `gorilla/websocket v1.5.3`; **zero** `*_test.go` files; Go 1.25.0. Both `app/internal/jobsimulation/runner/runner.go` and `jobsimulation/internal/runner/runner.go` carry the quoted *"formerly the standalone 'roadrunner' service"* header. Per the assignment's special note, the MERGED/ORPHANED opening is not flagged.
* **`ai-labs.md` — CLEAN.** Verified: `app` @ `5ba17044` = `v1.363.2`; `internal/{labs,credits,payments,subscriptions}` + `stripe/` all present; `credits/cost.go:231` `DefaultSeedBalance int64 = 500`, `:133` `MarginMarkup = 1.40`, `:142` `PricePerCreditUSD = 0.45`; `internal/web/backend/credits/handler.go` header records "POST /credits/purchase was removed in Wave 13" and `Register` mounts only `GET /balance` + `GET /transactions`; the Stripe webhook switch (`internal/web/backend/api/api.go:315-323`) handles `customer.created` + `customer.subscription.{created,updated,deleted}` and **no** `checkout.session.completed`. The "there IS a separate repo" self-correction and the two-meanings-of-v6.0 fence are both explicit.
* **`academy-backend.md` — CLEAN.** Verified: `academy_embedding_refresh` nightly @ 03:00 (`internal/worker/tasks/academy_embedding_refresh.go:10`, `internal/worker/worker.go:156`); `GET /content/catalog.json` (`internal/web/backend/content.go:23`) + the `ACADEMY_CONTENT_API_TOKEN` shared-token `/content/admin` group (`content_admin.go:35`). The router-deletion fact (#4) is handled correctly and explicitly at `:13-15`.
* **`coursebuilder.md` — CLEAN.** Verified: `go.mod:3` = `go 1.26.4`; `DefaultAuthorModelID = "eu.anthropic.claude-opus-4-8"` / `DefaultGraderModelID = "eu.anthropic.claude-sonnet-4-6"` (`internal/coursebuilder/bedrock.go:23,29`); `imagegen/openai.go:22` `defaultModel = "gpt-image-2"`; `DefaultSessionsPerOrgPerDay = 50` (`web/backend/coursebuilder/rate.go:28`); `COURSEBUILDER_MAX_MONTHLY_COGS_USD` → 500 default, `DefaultMaxDailyCOGSUSD = 0.0` (`budget.go:33,42`); app version + date match.
* **`clerk-integration.md` — CLEAN.** Verified: `app/go.mod:31` = `github.com/clerk/clerk-sdk-go/v2 v2.7.0` — the doc's own drift call-out is exactly right, line number included; **12** webhook event types in `internal/clerk/events/events.go:121-190` (`user.{created,updated,deleted}`, `organization.{created,deleted,updated}`, `organizationInvitation.{accepted,created,revoked}`, `organizationMembership.{created,deleted,updated}`); svix at `events.go:27,69`; `apps/web/src/proxy.ts` exists, `middleware.ts` does not.
* **`askengine.md` — CLEAN.** Verified: `DefaultModelID = "eu.anthropic.claude-sonnet-4-6"` (`internal/askengine/bedrock.go:25`), `ASK_MODEL_ID` override at `:163`; `askEngineMaxConns = 6` (`app/main.go:149`) fed from `COPILOT_DB_CONN` (`main.go:312,331,335`); `maxAgenticIterations = 15` + `loopTimeout = 10 * time.Minute` + `context.WithoutCancel` (`internal/web/backend/ask/handler.go:34,41,248,503`); `MaxInlineRows = 200` / `MaxCellLength = 400` (`executor.go:18,21`).
* **`skillpath.md` — CLEAN.** Correct merge-banner redirect. Verified: `repos.yml` = **9** repos with 0 skillpath and no `skillpath` compose service; `SKILLPATH_STREAM=skillpath` survives at `docker-compose.yml:64`; `app/internal/skillpath/` + `app/internal/skillpaths/` exist. The historical "3 subgraphs at the time" is correctly fenced with "it is **1** now".
* **`gotenberg.md` — CLEAN.** Verified line-for-line: image `gotenberg/gotenberg:8`, the exact four-flag command and `3200:3200` (`docker-compose.yml:371-381`), `profiles: [graphql, backend, all]` (`:384`); `GOTENBERG_URL=http://gotenberg:3200` (`:51`); `app/internal/converter/gotenberg.go:13` `Timeout: 90 * time.Second`, `:16` signature, `:31` `POST …/forms/libreoffice/convert`.
* **`customerio-sync.md` — CLEAN.** The quoted compose block matches `docker-compose.yml:220-238` byte-for-byte in substance (GitHub-URL build context, `ssh: ["default"]`, `8080:8080`, `search_path=public` DSN, `profiles: [customerio-sync, all]`); absent from `repos.yml` as claimed.
* **`skiller.md` — CLEAN.** Correct merge-banner redirect. Verified: `app/internal/rpc/skillerrpc/skiller.go` exists; `SKILLER_RPC_ADDR=http://backend:8083` on backend/jobsimulation/cms/messenger (`docker-compose.yml:62,121,174,265`); `internal/web/backend/graphql/graph/schemas/skiller_taxonomy.graphqls:7` states `categoryTree`/`fullCategoryTree` "stay unported — no consumers". The taxonomy-library-is-NodeID-only claim matches ground-truth #10.
* **`TEMPLATE.md` — CLEAN.** A pattern skeleton with no factual claims about the platform.
* **`db-backup.md` — CLEAN.** Production-only and explicitly fenced as such (`:27`, and README `:37`). No local-stack claim to falsify.
* **`intelligence.md` — CLEAN.** Correct archived-stub. "The Intelligence GitHub repository still exists" is compatible with ground-truth #11 (archived repos still exist), and the banner + "Historical Role (pre-2026-Q2)" heading fence the whole body.

Files with findings, for completeness: `storage.md` (1 blocker), `sentinel.md` (3 minors),
`messenger.md` (3 minors), `README.md` (2 minors), `next-web-app.md` (2 minors).

---

## 4. Totals

| | Count |
|---|---:|
| **BLOCKER** | **1** |
| **minor** | **10** |
| Files fully read | 18 / 18 |
| Files not fully read | **0** |

**Router-drift check (ground-truth fact #4), the flag risk for this group:** every file in G5 that touches the
GraphQL path was checked. `next-web-app.md`, `academy-backend.md`, `README.md` and `skillpath.md` have all
already been swept for it and state it correctly in the body; the only residue is the two `next-web-app.md`
captions above (minor-10, minor-11). No file in this group lists `graphql` as a local container, and no file
claims 2+ subgraphs as a present-tense fact.
