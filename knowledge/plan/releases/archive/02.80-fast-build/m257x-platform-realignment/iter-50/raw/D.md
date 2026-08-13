# Seat D — M257x iter-50 KB-fidelity reading

## 1. Header

**Corpus under audit:** `/Users/marco/workspace/anthropos/rosetta`, branch `m257x/platform-realignment`,
HEAD `57dfbfded8791fcb12a4651d747247ce9d04d7f0` (`git rev-parse HEAD`).

**Ground truth consulted** (verified in-situ, not assumed):

| clone | sha I measured (`git rev-parse --short HEAD`) |
|---|---|
| `stack-demo/platform` | `2adcf71` |
| `stack-demo/app` | (CHANGELOG head `v1.363.2 - 2026-07-31`) |
| `stack-demo/app/studio` | `aeec036` |
| `stack-demo/next-web-app` | `bb3313bc0` |
| `stack-demo/roadrunner` | `87d8d44` (last commit 2026-06-19) |
| `stack-demo/jobsimulation` | (commit `09631fb2` + PR #395 merge `500b9761` present) |
| `stack-demo/studio-desk` | `14a5442` |
| `stack-demo/graphql-wundergraph` | (commit `915da06` present) |
| `.agentspace/rosetta-extensions` | `a91f8f7` (2026-08-03) |

**Positive control — `wc -l` on every assigned file** (one invocation:
`wc -l corpus/services/{studio-room,clerkenstein,chronos,roadrunner,next-web-app,gotenberg,intelligence}.md`):

| file | lines | briefing said | read in full? |
|---|---|---|---|
| `corpus/services/studio-room.md` | 473 | 473 | yes, 1→473 |
| `corpus/services/clerkenstein.md` | 366 | 366 | yes, 1→366 |
| `corpus/services/chronos.md` | 245 | 245 | yes, 1→245 |
| `corpus/services/roadrunner.md` | 171 | 171 | yes, 1→171 |
| `corpus/services/next-web-app.md` | 126 | 126 | yes, 1→126 |
| `corpus/services/gotenberg.md` | 82 | 82 | yes, 1→82 |
| `corpus/services/intelligence.md` | 18 | 18 | yes, 1→18 |
| **total** | **1481** | 1481 | |

All seven counts match the briefing exactly; no file read short or empty.

**Search-pipeline hygiene.** One grep in this session failed with a zsh glob error
(`(eval):1: no matches found: --include=*.go`) — caught, re-run quoted, and the result re-taken. Every
absence claim below was paired with a control pattern known to match in the same invocation (e.g. the
`GOTENBERG` sweep was paired with `package converter`; the clerkenstein `r.Cookie(` sweep ran in the same
invocation as constant greps that returned hits).

---

## 2. BLOCKERS

| # | corpus `file:line` | the false claim (quoted) | what is true + ground-truth citation | reader harm |
|---|---|---|---|---|
| 1 | `corpus/services/studio-room.md:388` | "studio-room makes no GraphQL or Directus calls; **its only outbound API call is to the skills taxonomy service (`api.anthropos.work`) via `services/taxonomy.py`.**" | False as an absolute. The pipeline's *primary* outbound calls are to the AI providers. `stack-demo/app/studio/services/ai.py:1` `from openai import OpenAI, AzureOpenAI`, `:2` `from anthropic import Anthropic`; clients are instantiated at `services/ai.py:383` (`return OpenAI(api_key=self.api_key)`), `:530` (`AzureOpenAI(`), `:664` (`Anthropic(api_key=self.api_key)`). The same doc contradicts itself at `:36` ("AI Providers \| OpenAI, Azure OpenAI, Anthropic") and `:261-266` (`AZURE_ENDPOINT` / `OPENAI_API_KEY` / `ANTHROPIC_API_KEY` config). | Anyone building an egress allow-list, an air-gapped/offline demo, or a network policy for the `app` container's studio pipeline would allow only `api.anthropos.work` and the pipeline would fail at its first generation step. This is the one *absolute* in the doc that its own §"AI Service Configuration" refutes. |

**Honest severity note on #1.** The sentence sits under *"Integration Points"*, whose subject is
platform-to-platform integration ("no GraphQL or Directus calls"), so the *intent* was plainly "no other
**platform-service** call". I graded it a blocker because as written it is unambiguous, absolute, false,
and actionable; a reader deciding egress would act on it. If the gate wants to downgrade it, the fix is a
two-word scoping edit ("its only outbound *platform* API call") — not a retraction.

**No other blocker found in 1481 lines.** Everything else I could reach resolved.

---

## 3. MINORS (11)

1. **`studio-room.md:60-93`** — the "Project Structure" tree omits real top-level members `tests/`,
   `tools/`, `pytest.ini`, `cog.toml`, `CLAUDE.md`, `changelog.md`, and lists `workspace/` which does not
   exist in the repo (it is runtime-created by `gen.py:450-456`). `ls stack-demo/app/studio`.
2. **`studio-room.md:5` and `:25`** — "Since **cms-in-app v8.0** (`app` **v1.360.1**)". cms-in-app v8.0
   is `v1.360.0` (`stack-demo/app/CHANGELOG.md`, `## v1.360.0 - 2026-07-29`, the `(**cms-in-app**)` block);
   `v1.360.1` is the *follow-up* CI fix "pull studio via additional_repo like cms". The `additional_repo`
   attribution at `:25` is exactly right; the "since v8.0 = v1.360.1" equivalence at `:5` conflates two
   releases.
3. **`studio-room.md:390-400`** — "#### With CMS Service / **The CMS service** drives the full lifecycle"
   is present-tense standalone-service framing that the file's own banner (`:3-11`) retracts. It is the
   cms *domain* inside `app` (`app/internal/cms/worker/worker.go`).
4. **`clerkenstein.md:3`** — "**Last updated:** 2026-07-14" while the body carries v2.8 M256 (`:180-205`)
   and M257x iter-23 (`:270-275`) material dated into August 2026. Stale metadata line.
5. **`clerkenstein.md:18`** — the monorepo sections list ("`clerkenstein`, `demo-stack`,
   `stack-injection`, `stack-core`, `stack-seeding`, `alignment`") omits 5 of the 11 real sections:
   `dev-stack`, `playthroughs`, `stack-secrets`, `stack-snapshot`, `stack-verify`
   (`ls .agentspace/rosetta-extensions`).
6. **`clerkenstein.md:101`** — the `cmd/` row lists `mintpk` / `fake-fapi` / `fake-bapi` but omits
   `jwtkey` (`ls clerkenstein/cmd` → `fake-bapi fake-fapi jwtkey mintpk`).
7. **`clerkenstein.md:168-171`** — "**The read path takes NO request input.** `handleMe`, `handleToken`,
   `handleClient` and `handleMeOrganizationMemberships` all discard (or ignore) the `*http.Request`".
   Three of the four do take `_ *http.Request` (`clerk-frontend/server.go:241, 467, 488`), but
   `handleMeOrganizationMemberships` (`:512`) takes a live `r` and reads it —
   `r.URL.Query().Get("offset")` (`:527`) and `…Get("limit")` (`:531`). The load-bearing half of the claim
   (no *identity* input, no cookie read) is correct and I verified it: `r.Cookie(` appears nowhere in
   non-test `clerkenstein/` Go. The absolute overshoots.
8. **`roadrunner.md:33`** — cites `../architecture/architecture_overview.md:188` as corroborating that
   jobsimulation still starts locally. Line 188 is the **Skiller** row; the Jobsimulation row (which does
   say "**Container still starts locally** (`docker-compose.yml:83`, default profile)") is line **189**.
   Off-by-one. The companion citation `README.md:20-21` resolves correctly to
   `corpus/services/README.md:20-21` ("three of the four (cms, jobsimulation, roadrunner) still start
   CONTAINERS locally in the default `graphql` profile").
9. **`roadrunner.md:21`** — "M247 re-grepped `app` + `jobsimulation` … **zero hits outside CHANGELOG**".
   There is a second non-Go hit: `stack-demo/jobsimulation/knowledge/operational.md:68`
   (`| ROADRUNNER_RPC_ADDR | Roadrunner service address |`). The substantive claim — *no Go code reads
   it* — is fully correct (my grep across app/jobsimulation/cms/sentinel/storage/messenger/next-web-app
   returned only CHANGELOG, that knowledge doc, and `platform/docker-compose.yml:118`).
10. **`next-web-app.md:38-43`** — the shared-packages table omits `@anthropos/design`
    (`packages/design/package.json:2`). Four packages exist: `core-js`, `design`, `graphql`, `ui`.
11. **`chronos.md:9`** and **`:5`/`intelligence.md:5`** — (a) "moved to **in-process Asynq** running
    inside **jobsimulation**" carries no note that jobsimulation has since been folded into `app`
    (`app/internal/jobsimulation/`), the only doc in my set with a decommission banner that lacks the
    merge follow-up; (b) both banners say the removals happened "in mid-2026" — both platform commits are
    **2026-04-17** (`git log -1 045857c` / `fdfa189`).

---

## 4. Audited zeros — read in full, found clean

### `corpus/services/gotenberg.md` (82 lines) — **fully verified, zero findings**
Every claim in this file resolves:
- image `gotenberg/gotenberg:8`, the exact 4-token `command` (`--api-port=3200`, `--api-timeout=60s`,
  `--libreoffice-restart-after=50`), `ports: ["3200:3200"]`, `profiles: [graphql, backend, all]` —
  `platform/docker-compose.yml:371-384`, matching the doc's YAML block character-for-character.
- "the backend service (`app`) is the only consumer" — a repo-wide `GOTENBERG` sweep across
  app/platform/sentinel/storage/messenger/jobsimulation/cms returns exactly 4 hits, all app-side
  (`app/main.go:243`, `app/terraform/main.tf:281`,
  `app/internal/web/backend/coursebuilder/handler.go:242`, `platform/docker-compose.yml:51`).
- `ConvertToPDF(ctx, gotenbergURL, document, filename) ([]byte, error)` — `app/internal/converter/gotenberg.go:16`;
  90 s client timeout — `:13` (`&http.Client{Timeout: 90 * time.Second}`);
  `POST {URL}/forms/libreoffice/convert` — `:31`.
- `GOTENBERG_URL=http://gotenberg:3200` injected via the backend compose `environment:` —
  `docker-compose.yml:51` (inside the `backend:` block that opens at `:28`).

### `corpus/services/roadrunner.md` (171 lines) — every checkable claim verified
The heavily-repaired top banner (`:5-43`) survives scrutiny intact:
- `roadrunner/terraform/main.tf:19` → `service_desired_count          = 1` ✔ (exact line), untouched since
  `87d8d44` dated **2026-06-19** ✔.
- `repos.yml:29-31` → the roadrunner entry with `# legacy — folded into app; backend calls Judge0 directly` ✔.
- `docker-compose.yml:281` → `roadrunner:` service key ✔; profiles `[graphql, roadrunner, all]` (`:309`) ✔.
- jobsimulation "remains in `repos.yml:17` **and** in `docker-compose.yml:83` with
  `profiles: [graphql, jobsimulation, all]`" ✔ (`repos.yml:17` is the jobsimulation entry;
  `docker-compose.yml:83` is the service key; the profiles line is `:140`).
- "**1 of the 9 repos** — the count dropped when platform `2adcf71` deleted the router entry" ✔ —
  repos.yml has exactly 9 entries, and `git diff 2adcf71~1 2adcf71 -- repos.yml` shows the
  `graphql-wundergraph` block deleted.
- "no `ROADRUNNER_RPC_ADDR` / `RoadRunnerService` / `roadrunner:10401` read in any service's Go code" ✔
  (0 Go hits; `RoadRunnerService` returns 0 across app+jobsimulation).
- "code execution moved in-process" ✔ — `jobsimulation/internal/runner/runner.go:1-3` package comment
  reads *"(formerly the standalone \"roadrunner\" service)"*, and `app` runs it as
  `jsrunner "github.com/anthropos-work/app/internal/jobsimulation/runner"`
  (`app/internal/jobsimwiring/wiring.go:34`, `:118` `NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))`).
- Body: Go 1.25 (`go.mod:3`), asynq `v0.25.1` + `gorilla/websocket v1.5.3` (`go.mod:9-10`), ports
  10400/10401 with `PORT`/`RPC_PORT` (`docker-compose.yml:290-302`), `JUDGE0_BASE_URL=http://52.48.139.23:2358`
  (`:297`), `JUDGE0_API_KEY` the one Judge0 var **not** in the compose env block ✔, worker
  `Concurrency: 10` (`internal/worker/worker.go:25`), queue `roadrunner:default`
  (`internal/worker/queues/queues.go:4`), task const `roadrunner:submissionresult`
  (`internal/worker/tasks/tasks.go:4`) with the handler in `internal/runner` (`worker.go:48`),
  `asynq.MaxRetry(3)` (`internal/runner/runner.go:126`), poll `maxRetries = 15` at
  `time.Sleep(1 * time.Second)` (`:244`, `:273`), `RoadrunnerSubmissionCompleted` publish (`:277-278`),
  language map accepts `"py"` (`internal/runner/languages.go:26`), the exact file tree, the unwired
  `internal/lsp/lsp.go`, **zero `*_test.go` files** and `RUN go test -v ./...` at `Dockerfile:18` ✔.

### `corpus/services/next-web-app.md` (126 lines) — verified except minor #10
- Versions: `next ^16.2.7` in all four apps (`apps/{web,hiring,integration,maintenance}/package.json`),
  `react ^19.2.7`, `"node": ">=24.0.0"`, `"packageManager": "pnpm@10.30.3"`, `turbo ^2.9.6`,
  `UPGRADE-IMPACT-next16.md` + `knowledge/next15-adoption-plan.md` both present ✔.
- The "corrected" gotcha at `:115` — the repo's `CLAUDE.md:15` does read *"Next.js 16 App Router"* ✔
  (the retraction of the older "it says 14" note is correct).
- `:48` — `apps/web/src/proxy.ts` and `apps/hiring/src/proxy.ts` exist, **no `middleware.ts` anywhere**;
  the repo's `CLAUDE.md:55` says verbatim *"Clerk middleware lives in `src/proxy.ts` (Next 16 renamed the
  `middleware.ts` convention → `proxy.ts`)"* ✔. Every listed public route is in the real allowlist
  (`proxy.ts:7-57`); `/print` is HMAC-gated via `PRINT_ROUTE_SECRET`
  (`packages/core-js/src/security/printToken.ts:16,49`, imported at `proxy.ts:1`) ✔.
- `:47`/`:96` — `docker-compose.yml:352` bakes
  `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT: http://${PUBLIC_HOST:-localhost}:8082/graphql/query` as a **build
  arg** ✔; `git show 2adcf71~1:docker-compose.yml` shows the prior value was
  `:5050/graphql` ✔; there is no `graphql` router service left in compose ✔; `app` serves
  `/graphql/query` on `PORT=8082` (`app/internal/web/backend/backend.go:317`, compose `:56`) ✔; the
  supergraph 2→1 commit `915da06` is real and lives in **graphql-wundergraph**
  ("fold cms subgraph into backend (supergraph 2→1) — cms-in-app v8.0 (#293)", 2026-07-29) ✔.
- **The `:30` AI-readiness gate correction is right, and I checked both branches** (this is exactly the
  "right citation on the wrong branch" shape the briefing warns about — it isn't one here): the **member**
  funnel reads *both* — `apps/web/src/components/ai-readiness/data/useAiReadinessActive.ts:22`
  (`useFeatureFlagEnabled(AI_READINESS_FLAG)`, `AI_READINESS_FLAG = 'flag_ai_readiness'`,
  `aiReadiness.constants.ts:26`) **AND** `:29` (`useAiReadinessEnabled(flagEnabled)`), returning
  `flagEnabled && orgEnabled === true` (`:32`); the **manager dashboard**
  (`app/(authenticated)/(verified)/ai-readiness/AIReadinessClient.tsx:133`) reads *only*
  `useAiReadinessEnabled(true)` and contains **zero** PostHog references (grep for
  `useFeatureFlag|posthog|AI_READINESS_FLAG` in that file returns nothing).
- Structure: 8 locales on disk (`configs/i18n/messages`: de,en,es,fr,it,ja,nl,pt) ✔; **exactly one**
  Dockerfile at the repo root (`Dockerfile.dev`, `FROM node:24-alpine`) building only
  `--filter=@anthropos/web-app` (`:32`) ✔; **no `.storybook/` and no `storybook` script**, with
  `configs/tailwind/storybooks.css` surviving exactly as claimed ✔; mobile **excluded** from the
  workspace (`pnpm-workspace.yaml` `- '!apps/mobile'`) on Expo port 3031
  (`apps/mobile/package.json:70`) ✔; dev ports 3000/3001/3002 ✔; `test` script present in web + hiring,
  absent in integration + maintenance ✔; `antd ^6.3.7` (Ant Design 6) ✔; `repos.yml` type `node-pnpm` ✔.
- Cross-repo claim at `:13` verified from the other side: `useGetClerkOrganization.ts:20` derives
  `const isHiringOrg = Boolean(organization?.publicMetadata?.isHiring)`.

### `corpus/services/studio-room.md` (473 lines) — every anchor resolves; only the §Integration Points absolute fails
This file is the densest in `file:line` anchors of the seven, and **all of them land**:
- `gen.py:18-28` `parse_argument` (`parse_known_args` + the `zip(unknown[::2], unknown[1::2])` fold) ✔ —
  so the `--template foo` swallow is real, and the "grep: zero consumers" claim holds (the only `template`
  reads in the package are the *legacy blueprint* stripper and the argparse description string).
- `gen.py:205-238` `translate_legacy_blueprint` + `_LEGACY_TEMPLATE_DEFAULTS` with **exactly** the three
  legacy names ✔, and the else-branch warning text *"is ignored; asset type is now inferred from task
  interactions"* is verbatim (`:227-228`) ✔.
- `gen.py:241-271` `validate_blueprint_exclusivity`, the exact whitelist the doc lists at `:282`
  (`media, forced, interactive, branch, pipeline` + the four `*_path`), and the verbatim error
  *"Cannot combine --blueprint with content parameters: …"* ✔; it is actually invoked in blueprint mode
  (`gen.py:423`).
- `gen.py:273-282` `setup_generation_request` (settings first, CLI args over) ✔.
- `gen.py:450-456` `work_paths` with the literal fallback `path_key.replace('_path','')` ✔.
- `gen.py:484-492` — **exactly nine** `add_argument` calls, each flag/dest/default matching the doc's
  table (incl. `-f/--force` → `dest: forced`, `--annotations` default `"{}"`, `--pipeline` default
  `linear`) ✔.
- `agents/simulation/postgen/exporter.py:518-519` — the two zips (`postgen/` + `published/`), and
  `:514-550` the unpack → write `simulation.json` → `_export_assets` → `make_archive` → `rmtree` sequence,
  including the `collaborative_*` / `asset_*` / `internal_*` preservation predicate at `:537` and the
  `source_id`-else-simid naming at `:517` ✔.
- `agents/simulation/model.py:59` + `:467-469` — the sole taxonomy memo ✔, and the doc's absolute
  *"the only cache in the pipeline"* is **correct**: a package-wide `cache|Cache|lru_cache` sweep
  (excluding tests/benchmark) returns those four lines and nothing else.
- `app/internal/cms/studio/studioManager.go:119` invokes `studio/gen.py` ✔; `:94-96` are the
  `studio/studio-venv` constants ✔.
- The repo's own `CLAUDE.md:12-14` is the quoted entry-point command ✔.
- Config: `max_tokens = 4000` in **all three** tracked configs ✔ (`config_template.ini:22`,
  `production_config.ini:21`, `development_config.ini:21`); the five `*_AI_STABLE_MODEL` rows are
  byte-identical to the doc's ini block (`production_config.ini:26-30`) with the matching
  `*_AI_EXPERIMENTAL_MODEL` set (`:32-36`) ✔; `configs/local_*` + `configs/test_*` gitignored
  (`.gitignore:5-6`) ✔; the four `*_path` values are `workspace/{attachments,trace,published,postgen}` ✔.
- `requirements.txt` unpinned, exactly the listed 9 packages, **no `aiohttp`** ✔.
- postgen: `--media`/`--simid`/`--target` all `required=True` (`postgen.py:396-398`) ✔; the four targets
  exist as modules plus `testing.py`, and `exporter.py` is not a selectable target ✔.
- Blueprints: `tests/e2e/blueprints/` holds exactly b2b/business/none_chat/none_voice/technical ✔, and
  the doc's `technical.json` excerpt is a faithful subset of the real file (every quoted key/value
  matches, incl. the `category` uuid, `salt: "0329"` and the `simid`) ✔;
  `knowledge/development/asset-examples/blueprints/` holds the code/micro/scenario trio ✔.
- State files `{simid}_pre_generation.json` / `{simid}_task_state.json` (`gen.py:116-117`) + the per-run
  `{simid}_usage.json` (`gen.py:458`) ✔.
- Deployment: `additional_repo: "anthropos-studio-room:studio"` at
  `app/.github/workflows/build-production.yml:29` ✔; runtime `FROM python:3.11-slim`
  (`app/Dockerfile:28`, `Dockerfile.dev:26`) ✔; the cms Asynq worker is `Concurrency: 5` with
  `AiVideoQueue: 7` / `StudioQueue: 3` (`app/internal/cms/worker/worker.go:29-33`) — the doc's
  "scheduling priorities, not concurrency limits" gloss is correct ✔; studio-room is **not** in
  `repos.yml` ✔.

### `corpus/services/clerkenstein.md` (366 lines) — all cited anchors resolve
The file is almost entirely repaired/retraction text; I checked each anchor rather than the prose:
- `alignment/cmd/alignctl/run.go:134-135` → `ExitRegressed = 2` / `ExitUnmeasurable = 3` ✔, exactly the
  two lines cited, with the surrounding rationale block (`:121-132`) saying what the doc says it says.
- `clerk-backend/store.go:138` `func (s *Store) SeedOrgIdentity(org, eid string)` and `:151`
  `func (s *Store) LookupOrgEid(org string)` ✔ — exact lines.
- `alignment/dna/clerk-2.6.0.json:131` contains verbatim *"M219 landed the fix … taking the Go surface
  97.2% -> 100%."* ✔.
- **Gene counts all check out** (`python3 -c` over each DNA file): `clerk-2.6.0` 14 capabilities / **27**
  genes; `clerk-js-5` **9**; `clerk-multi-1` **9**; `clerk-deploy-1` **7**; `clerk-express-1` **13 genes
  across 5 capabilities** — every number in the ⚠ score table and the M224 `/align-run` record matches.
- The "not yet a measured gene" disclosure at `:141-143` is **true and load-bearing**: `clerk-js-5.json`
  has a `Me` capability (variants `universal-user`, `unauthenticated`) and **zero** occurrences of
  `organization_memberships`. Likewise "no DNA gene covers `GET /npm/`" — `/npm/` count is 0 in all five DNAs.
- `SessionToken/decoded-identity` really is `"criticality": "critical"` with `"operator": "exact"` ✔.
- The BAPI≠FAPI section: `clerk-frontend/meorgmemberships_test.go` exists ✔; the route is registered at
  `clerk-frontend/server.go:186`; all three pinned properties are in the handler
  (`:512-543`) — paginated envelope `{"data":…,"total_count":…}` in `Response` (`:539-541`), `limit`/`offset`
  honoured (`:527`,`:531`) with `total_count` as the **true** total (`:522`) and a non-nil empty slice
  (`:534-537`), unauthenticated ⇒ **401** (`:515-519`) ✔. The `~4.05 s` figure cross-referenced to
  `latency-budget.md` §"Time-to-usable" matches that doc's measured **4049 ms** (`latency-budget.md:445`) ✔.
- studio-desk's `STUDIO_ACCESS_ROLES` is at **`src/index.ts:96`** and **`app/services/userService.ts:16`**
  with the quoted comment *"Both the prefixed (`org:*`) and bare role keys are accepted"* — both exact ✔.
- Single-tenant seat disclosure: `clerk-frontend/registry.go:24,46,67,70,79` (`activeKey`, `active()`,
  `ActiveKey()`, `Select()`) ✔; `clientID: "client_clerkenstein"` (`server.go:125`) and
  `sessID = "sess_clerkenstein"` minted in `establishLocked` (`:663-665`) ✔; `handleSelectIdentity`
  (`:627`) sets `signedIn = false; sessID = ""` (`:646-647`) and deliberately clears `signedOut`
  (`:654`) ✔; `handleSignOut` (`:554`) takes `_ *http.Request` — ignores its `{id}` ✔; **`r.Cookie(`
  appears nowhere** in non-test clerkenstein Go ✔.
- D81: all **five** named tests exist in `clerk-frontend/server_test.go` at `:256, :286, :390, :427, :461` ✔;
  `POST /v1/client/sessions` with `_method=DELETE` dispatch at `server.go:583` ✔; whitelist semantics ✔.
- M220 clerk-js fix: `clerk-frontend/server.go:35-67` — `clerkJSFetchTimeout = 15 * time.Second` (`:59`),
  `clerkJSClient` with the literal comment *"Explicitly NOT http.DefaultClient"* (`:65-68`),
  `FAKE_FAPI_CLERKJS_CACHE` disk cache, and the stated no-`http.Get(` test ✔.
  `FAKE_FAPI_CLERKJS_CDN` override at `:29` over `defaultClerkJSCDN = "https://cdn.jsdelivr.net"` (`:22`) ✔.
- M39 / M224 roster threading, **producer and consumer both**:
  `stack-seeding/seeders/roster.go:53-54` (`org_name`/`org_slug`), `:77` (`org_is_hiring,omitempty`),
  `:140`/`:237` (`orgSlugFor`), `:155`/`:246` (`IsHiringOrg()`); `clerk-frontend/registry.go:125-126,144`
  + `:190` `DisallowUnknownFields()`; `clerk-frontend/resources.go:255` `orgMemberships()` with the
  `orgNameDefault`/`orgSlugDefault` fallbacks (`:50-51`, applied `:263-270`) and the **conditional-emit**
  `if u.OrgIsHiring { orgPublicMeta["isHiring"] = true }` over a base `{"eid": …}` (`:274-276`) ✔ —
  i.e. the "byte-identically `{eid}`" align-safety claim is literally true in the code.
- The M257x iter-23 ⚠ box (`:270-275`) is **correct in both directions**: `app/go.mod:16`
  `colony v0.35.2` and `:31` `clerk-sdk-go/v2 v2.7.0`; `sentinel/go.mod:8` and `storage/go.mod:7` both
  `colony v0.34.3`; clerkenstein's own `go.mod:8` pins `colony v0.34.3`. (Uncontradicted extra I found:
  `messenger/go.mod:7` is also `v0.35.2` — the doc names only sentinel+storage but does not claim
  exhaustivity.)
- Repo-structure table: `authn/ clerk-backend/ clerk-frontend/ clerk-webhook/ shared/ deploy/ cmd/
  alignment/` all exist ✔; `deploy/colony-authn/` holds the drop-in ✔;
  `alignment/{cmd/{clerkrun,jsfapirun,multirun,expressrun,deployrun}, dna(×5), golden{,-js,-multi,-express,-deploy}, scripts}` ✔;
  the sibling `rosetta-extensions/alignment/` section has **no** `scripts/` dir ✔ and
  `ALIGN_DIR` defaults to `$base/../../alignment` (`clerkenstein/alignment/scripts/gate.sh:30`,
  `drift-check.sh:16`) ✔; `.github/workflows/alignment.yml` is git-tracked (`git ls-files .github`) and,
  being below a monorepo *section* rather than a repo root, is inert as described ✔;
  `FAKE_FAPI_ROSTER` read at `cmd/fake-fapi/main.go:25` ✔; `dotless-pk-rejected` is a real gene
  (`clerk-express-1` → `ExpressRequest`) ✔.
- Every KB file named in "Read next" exists in **both** clone roles
  (`kb-index.md`, `scope.md`, `architecture.md`, `injection.md`, `alignment.md`, `coverage-index.md`).

### Link/anchor integrity across all seven files — clean
Every relative link resolves (23 targets checked by existence) and every named section anchor exists:
`alignment_testing.md#how-m1-m1b-m2-and-m2c-consume-this` (`:311`),
`#what-alignment-proves--and-what-it-doesnt-the-m3-lesson` (`:378`),
`external_services.md#clerk-authentication-service` (`:21`),
`cockpit-spec.md` § *Limitation — one seat per stack* (`:414`),
`setup_guide.md` §"Linux host prerequisites (for a remote/VM demo over Tailscale)" (`:110`),
`latency-budget.md` §"Time-to-usable" (`:416`), `recipe-browser-login.md §B` (`:42`),
`safety.md` §3.8 (present). **Zero broken links.**

### `corpus/services/intelligence.md` (18 lines) — the one checkable claim is verified
`platform` commit `fdfa189` exists and its subject is *"chore: remove intelligence service from local dev
orchestration"* (2026-04-17) — the doc quotes it as "remove intelligence service from local dev
orchestration" ✔. `intelligence` appears **nowhere** in `docker-compose.yml`, `repos.yml`, `common.yml`
or the `Makefile` ✔. The skiller→app parenthetical is consistent with `repos.yml:4-9`.

### `corpus/services/chronos.md` (245 lines) — banner verified; body unverifiable (see §5)
Every *checkable* claim in the banner is right, and precisely right:
- platform `045857c` = "chore: remove chronos service from orchestration and update related
  documentation" (2026-04-17) ✔; `chronos` appears nowhere in compose/repos.yml/Makefile/common.yml ✔.
- jobsimulation commit `09631fb2` = "refactor: remove Chronos references and update documentation to
  reflect Asynq integration for session timeout management" ✔, and PR **#395** merged branch
  `feat/remove-chronos-and-realtime` (merge `500b9761`) ✔ — both quoted strings verbatim.
- "moved to in-process Asynq" ✔ — `jobsimulation/go.mod:24` `hibiken/asynq v0.26.0`, and **zero**
  `chronos` references remain in jobsimulation Go code.
- `:186` "Chronos is no longer present in `platform/docker-compose.yml` … `docker compose up -d chronos`
  will fail" ✔.
- The header's "**but the GitHub repo is NOT archived**" is *consistent* with the M257x correction carried
  in `roadrunner.md:35-37` — the two repaired sites agree rather than contradict, which is the failure
  shape the briefing flags. (The underlying archive fact itself is unverifiable here — §5.)

---

## 5. Unverified — and why

1. **`chronos.md:13-245` — the entire technical body.** The `chronos` repo is not cloned (no
   `stack-demo/chronos`, no `stack-dev/` in this checkout). Unverified: the `timers` DDL and its four
   indexes, the `ChronosService` RPC shape, the Triggerer's `LIMIT 10` / `FOR UPDATE SKIP LOCKED` /
   `scheduled_at < now() + interval '1 second'` / 3-second run timeout / per-minute Sentry heartbeat, the
   `EventTimerWentOff` proto, the env-var and CLI-flag tables, `search_path=chronos`, Go 1.25, sqlc+Cobra.
   Not reportable as passed **or** as a blocker.
2. **`intelligence.md:12-18` — the historical body** (`DB_CONNECTION_BACKEND` / `DB_CONNECTION_SKILLER`,
   `:8080` `/_meta`, the 5-minute ticker). Repo not cloned.
3. **All GitHub archive/push states.** `gh` is unavailable, so I cannot confirm: skiller archived
   2026-07-01, skillpath 2026-07-31, jobsimulation 2026-07-31 (`roadrunner.md:25,31`), chronos "NOT
   archived, last push 2026-04-23" (`roadrunner.md:35-37`, `chronos.md:3`), intelligence "still exists"
   (`intelligence.md:8`), roadrunner "still in repos.yml, not archived".
4. **Live alignment SCORES.** I verified the DNA gene *counts*, criticalities and operators from
   `alignment/dna/*.json`, and the exit-code contract in `alignctl`. I did **not** run `alignctl`/
   `deployrun`/`expressrun` — that needs `colony` (not cloned) and `@clerk/express` `node_modules`. So
   the "100.0%/100.0%" figures at `clerkenstein.md:38-46, 68-76, 245-247` are unverified, as is the
   claim that the express runner exits rc=3 on a box without `node_modules` (the *mechanism* at
   `run.go:121-153` is verified; the *observed outcome* is not). The doc's own framing — that the express
   surface is "frequently *unmeasured*" — is exactly the honest posture, so I flag no finding.
5. **`colony` / `proto` / `taxonomy`-dependent claims.** Not cloned. Affects: `roadrunner.md:92,99`
   (`/_meta` served by colony's HTTP server — only *indirectly* corroborated by
   `roadrunner/terraform/main.tf:22` `health_check_path = "/_meta"`), `roadrunner.md:131`
   (`proto/roadrunner/v1/roadrunner.proto`), `chronos.md:101-116` (the events proto),
   `clerkenstein.md:76,264-268` (the `deploy/` drop-in compiling against real colony `v0.34.3`).
6. **Internal ledger identifiers** — `FIX-M219-bapi-org-eid` "is CLOSED",
   `CHECK-M257x-iter22-clerk-sdk-drift`, `D81`, `#M224-D-align`, `#M213-D-*`, `#M39-D2/D3`,
   `DEF-M240-01`. These live under `knowledge/plan/**`, which this seat is forbidden to read. The *code*
   each one points at was verified; only the ledger state was not.
7. **`studio-room.md:331` "Python 3.9+"** — no `python_requires` / `setup.py` / `pyproject.toml` in the
   studio repo to check it against; the *runtime* half of the same line (`python:3.11-slim`) is verified.
8. **Prod-side assertions** — `next-web-app.md:53` "in prod, the router" is *consistent with* the
   surviving `graphql-wundergraph/{terraform,config.prod.yaml,supergraph-config-prod.yaml}`, but I did not
   inspect deployed infrastructure.
