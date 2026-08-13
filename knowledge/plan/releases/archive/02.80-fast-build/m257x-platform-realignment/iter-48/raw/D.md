# Seat D — iter-48

Repo: `/Users/marco/workspace/anthropos/rosetta` @ `m257x/platform-realignment` (`cabc3b1`).
Ground truth: `stack-demo/app` @ `5ba17044` (`chore(version): v1.363.2`), incl. the Studio-Room Python
pipeline at `stack-demo/app/studio/`; plus `stack-demo/{platform@2adcf71, next-web-app@bb3313bc,
roadrunner@87d8d44, jobsimulation, sentinel, storage, studio-desk}` and the `rosetta-extensions`
authoring copy at `.agentspace/rosetta-extensions` @ `932554e`.

## Coverage (file, wc -l, lines read)

| # | File | `wc -l` | Lines read |
|---|---|---|---|
| 1 | `corpus/services/studio-room.md` | 473 | **all 473** |
| 2 | `corpus/services/clerkenstein.md` | 366 | **all 366** |
| 3 | `corpus/services/chronos.md` | 245 | **all 245** |
| 4 | `corpus/services/roadrunner.md` | 171 | **all 171** |
| 5 | `corpus/services/next-web-app.md` | 126 | **all 126** |
| 6 | `corpus/services/gotenberg.md` | 82 | **all 82** |
| 7 | `corpus/services/intelligence.md` | 18 | **all 18** |
| | **total** | **1481** | **all 1481** |

Every file was read top-to-bottom in full via `Read` with no `offset`/`limit`, so the read covers
exactly the line count above. No file was sampled.

### Search hygiene (platform-alignment.md §5)

- Every `grep` in this pass was run with stderr visible. One rejection was caught and re-run: an
  unquoted `--include=*.py` was eaten by zsh globbing (`(eval):1: no matches found: --include=*.py`)
  and would otherwise have read as "zero `template` consumers" — the exact rule-1 failure. Re-run
  quoted, it returned 40 hits.
- Positive controls run in the same pass as each negative: `grep "def parse_argument"` (2 hits)
  alongside the `template`-consumer sweep; `grep JUDGE0_BASE_URL` in `app` (1 hit) alongside the
  `ROADRUNNER_RPC_ADDR` sweep (0); `grep "package runner"` in `jobsimulation` (3) alongside its
  roadrunner sweep (0); `grep -ril bedrock stack-demo/app/internal/askengine/` (2 files) alongside
  `grep -rin 'bedrock|boto3' stack-demo/app/studio/` (**0**).
- Every quoted `file:line` below was opened and read with its surrounding block, not grepped to.

### THE FLAGGED ITEM — Studio-Room's AI providers (RESOLVED: the doc is correct)

Read from source, not inferred:

- `stack-demo/app/studio/services/ai.py:704-724` — `get_client(engine, target_override=None)` builds
  `providers = {'openai': OpenAIProvider, 'azure': AzureProvider, 'anthropic': AnthropicProvider}`
  and does `providers.get(target_engine['service'])`, raising
  `ValueError(f"Unknown AI service: {target_engine['service']}")` on a miss. **Three providers, no
  fourth.** The classes are defined at `:334` (`OpenAIProvider`), `:490` (`AzureProvider`),
  `:627` (`AnthropicProvider`); imports at `:1-2` are `from openai import OpenAI, AzureOpenAI` /
  `from anthropic import Anthropic` — the **first-party** Anthropic SDK.
- **Selection mechanism:** per-`GenMode`, per-branch, from the INI. `configs/{env}_config.ini`
  `[SERVICES]` rows are `{MODE}_AI_{BRANCH}_MODEL = service, model, thinking` (the comment in all
  three tracked configs reads literally `# TARGET SERVICE: openai, azure, anthropic`). `get_client`
  indexes `engine[<mode>.value]` and dispatches on that row's `service` token. `GenMode`
  (`services/ai.py:35-42`) is `FAST|STRICT|EXECUTION|CREATIVE|REASONING` with `DEFAULT = EXECUTION`.
- **AWS Bedrock is NOT among them.** `grep -rin 'bedrock|boto3' stack-demo/app/studio/` → **0**
  (positive control: the same pattern hits 2 files under `app/internal/askengine/`). There is no
  `boto3`/`botocore` in `studio/requirements.txt`.

`studio-room.md:36` ("**AI Providers** | OpenAI, Azure OpenAI, Anthropic"), `:49` (mermaid
`AI Service<br/>OpenAI/Azure/Anthropic`), `:113` and `:253-268` are all **correct**, and consistent
with the corpus's own retraction at `corpus/architecture/ai_architecture.md:45-46` and `:84`
("Studio-Room was never on Bedrock"). **No finding.** The conflation the seat was asked to hunt is
already fixed in these files.

## Blockers

**None.** Zero BLOCKERs found across all seven files.

Every mechanism, number, guard polarity and cross-corpus twin I could reach was checked against
platform source or the cited corpus site, and each held. The verification actually performed
(non-exhaustive):

| Claim | Verified against |
|---|---|
| `gen.py` registers **exactly nine** args at `gen.py:484-492` | exact — 9 `add_argument` calls, lines 484–492 |
| `parse_argument` folds unknown `--k v` pairs via `parse_known_args` | `gen.py:18-28` verbatim |
| `--template` is silently swallowed; **zero** consumers | quoted repo-wide grep: only the legacy-strip + an argparse `description` string |
| `--blueprint` + `--template` fails loud with that exact message | `gen.py:241-271`, whitelist excludes `template`, message matches verbatim |
| legacy names = exactly `micro/scenario/collaborative challenge`; anything else warns "*is ignored; asset type is now inferred from task interactions*" | `gen.py:205-238`, string matches verbatim |
| blueprint mode merges over a whitelist of execution controls | `gen.py:433-441` — the same 9 keys |
| `worklog_path` in-code fallback is a literal `worklog/` | `gen.py:450-455` (`path_key.replace('_path','')`) |
| `max_tokens = 4000` in **all three** tracked configs | template / development / production `.ini`, all `4000` |
| env vars override INI at load; `configparser` does not interpolate | `gen.py:41-53`, the 6-key `secrets_keys` list |
| studio queue weight **3** vs ai_video **7**, worker `Concurrency: 5` | `app/internal/cms/worker/worker.go:29,32-33` |
| Go side invokes `studio/gen.py`; venv at `studio/studio-venv` | `internal/cms/studio/studioManager.go:119`, `:94` |
| zip in **each** of `postgen/` + `published/` | `exporter.py:518-519` exact |
| unpack → write `simulation.json` → re-archive → `rmtree` | `exporter.py:514-550` |
| taxonomy memo is the only cache | `agents/simulation/model.py:59`, `:467-469` exact |
| 4 selectable post-gen targets; `testing` is a module, not a target; export is not a target | `agents/simulation/postgen/__init__.py:47-75` |
| `--media/--simid/--target` all `required=True` | `postgen.py:396-398` |
| runtime image `python:3.11-slim`; pulled via `additional_repo` at app **v1.360.1** | `app/Dockerfile:28,42-46`; `.github/workflows/build-production.yml:29`; `CHANGELOG.md:80-82` |
| `configs/local_*` + `configs/test_*` gitignored | `studio/.gitignore` |
| `ExitRegressed=2` / `ExitUnmeasurable=3` at `run.go:134-135` | exact, with the surrounding rationale block read |
| `store.go:138` `SeedOrgIdentity`, `:151` `LookupOrgEid` | exact |
| `clerk-2.6.0.json:131` carries the "97.2% -> 100%" sentence | exact |
| gene counts 27/14 · 9 · 9 · 13/5 · 7 | parsed all five DNA JSONs |
| deploy pinned `colony v0.34.3`; `app` on `v0.35.2`; sentinel+storage on `v0.34.3`; `app` on `clerk-sdk-go/v2 v2.7.0` | `clerk-deploy-1.json:6`; `app/go.mod:16,31`; `sentinel/go.mod:8`; `storage/go.mod:7` — **the iter-23 drift ⚠️ is exactly right** |
| FAPI twin at `clerk-frontend/server.go:186`, BAPI at `clerk-backend/server.go:47` | exact |
| `clerk-js-5.json` has a `Me` capability but **no** gene for `/v1/me/organization_memberships` | enumerated all 6 capabilities / 9 variants |
| single-seat: one `activeKey`, one `clientID`/`sessID` const, `r.Cookie` called **nowhere** | `registry.go:24,67,75`; `server.go:125,665`; 0 hits |
| all five `server_test.go` test names | `:256,286,390,427,461` |
| `clerkJSFetchTimeout = 15s`, "Explicitly NOT http.DefaultClient", `FAKE_FAPI_CLERKJS_CACHE` | `clerk-frontend/server.go:35-68` |
| conditional-emit `if u.OrgIsHiring { pm["isHiring"]=true }` | `resources.go:271-276`; producer `stack-seeding/seeders/roster.go:77,155,246` |
| studio-desk accepts both role forms at `src/index.ts:96` + `app/services/userService.ts:16` | exact, incl. the quoted comment |
| `ALIGN_DIR` default `../../alignment`; rext `alignment/` has no `scripts/`; workflow inert (no repo-root `.github/workflows`) | `gate.sh:30`, `drift-check.sh:16`, `ls` |
| roadrunner: 9 repos in `repos.yml` (10 → 9 at `2adcf71`), `:29-31`; compose `:281`, jobsim `:83`/`repos.yml:17`; `ROADRUNNER_RPC_ADDR` at compose `:118`; **0** Go readers in `app` **and** `jobsimulation` | `git show 2adcf71^:repos.yml` = 10 entries incl. `graphql-wundergraph` |
| `terraform/main.tf:19` = `service_desired_count = 1` | exact |
| roadrunner: 0 `*_test.go`; `go test` at `Dockerfile` line **18**; queue `roadrunner:default`; `MaxRetry(3)`; `Concurrency: 10`; 15 polls @ 1s; LSP unwired | `internal/worker/queues/queues.go:4`, `internal/runner/runner.go:126,244,273`, `internal/worker/worker.go:25` |
| next-web: Next **16.2.7** across exactly 4 apps, React 19.2.7, Node ≥24, pnpm 10.30.3, Turborepo ^2.9.6 | `package.json` / `apps/*/package.json` |
| `proxy.ts` not `middleware.ts`; repo `CLAUDE.md:55` says so; `CLAUDE.md:15` says "Next.js 16 App Router" | both anchors exact |
| public allowlist contains all 7 named routes; `/print` HMAC-gated by `PRINT_ROUTE_SECRET` | `proxy.ts:7-57,59,66-72`; `packages/core-js/src/security/printToken.ts:16,49` |
| **only one Dockerfile** (`Dockerfile.dev`) while repo `CLAUDE.md` still says "Two Dockerfiles" | `ls` + `CLAUDE.md` — the doc's gotcha is correct |
| no `storybook` script, no `.storybook/`, only `configs/tailwind/storybooks.css` | exact |
| 8 locales on disk (de,en,es,fr,it,ja,nl,pt) while repo `CLAUDE.md` says 7 | `ls configs/i18n/messages/` |
| `docker-compose.yml:352` bakes `:8082/graphql/query`; was `:5050/graphql` pre-`2adcf71` | exact, incl. `git show 2adcf71^` |
| `graphql-request` + TanStack React Query, **no** Apollo; AntD 6; `please-use-pnpm` | `apps/web/package.json:27,32,39`; root `:15,17` |
| gotenberg image/ports/command/`--libreoffice-restart-after=50`, `GOTENBERG_URL`, 90 s client timeout, `ConvertToPDF` signature, `POST /forms/libreoffice/convert` | `platform/docker-compose.yml:371-384,51`; `app/internal/converter/gotenberg.go:13,16,31` |
| chronos `045857c` / intelligence `fdfa189` / jobsim `09631fb2` all exist with the quoted subjects; both services absent from compose+repos.yml+Makefile | `git log` in `platform` and `jobsimulation` |

**Not verifiable in this environment (stated, not counted as findings):** the GitHub *archived* flags
and push dates asserted in `roadrunner.md:25-37` (skiller 2026-07-01, skillpath 2026-07-31,
jobsimulation 2026-07-31, chronos NOT archived / last push 2026-04-23) — `gh` is not installed on
this box and there is no network path. The internal code claims of `chronos.md` (schema.sql,
`LIMIT 10` + `FOR UPDATE SKIP LOCKED`, 3 s run timeout, Sentry heartbeat) and `intelligence.md`
(port 8080, `/_meta` only, 5-minute ticker) are likewise unreachable — neither repo is cloned into
any `stack-*/` workspace, which is itself consistent with both docs' "no longer cloned by
`make init`".

## Minors

1. **`corpus/services/studio-room.md:388`** — "studio-room makes no GraphQL or Directus calls; **its
   only outbound API call is to the skills taxonomy service** (`api.anthropos.work`) via
   `services/taxonomy.py`." Over-broad. It also calls the AI provider APIs
   (`services/ai.py:383,530,664` construct `OpenAI`/`AzureOpenAI`/`Anthropic` clients) and fetches
   generated images over HTTP — `agents/simulation/export.py:51`:
   `response = requests.get(image_url, stream=True)`. The **predicate** the sentence exists to
   support (studio-room is not a GraphQL/Directus client — CMS orchestrates) is TRUE, and the AI
   egress is documented at length three sections above, so nothing load-bearing is wrong. Suggest
   "its only *platform*-API call". The image fetch is undocumented anywhere in the file.

2. **`corpus/services/studio-room.md:350`** — the requirements block annotates `mistralai` as
   `# AI provider`. There is **no Mistral provider**: `services/ai.py:704-709` maps only
   `openai`/`azure`/`anthropic`, and `:713-715` raises `ValueError(f"Unknown AI service: ...")` for
   anything else. `mistralai` is present-but-unused in `requirements.txt`. Worth fixing precisely
   because this file is the provider-conflation surface — the authoritative row at `:36` is right,
   and this annotation is the one place a reader could re-derive a fourth provider.

3. **`corpus/services/studio-room.md:210,281`** — `translate_legacy_blueprint` is cited as
   `gen.py:205-238`; the `def` is at **212** (205-209 is the `_LEGACY_TEMPLATE_DEFAULTS` dict it
   reads, so the range is defensible as a block, just not as the function). The three other gen.py
   anchors (`18-28`, `241-271`, `273-282`, `484-492`, `450-456`) are exact.

4. **`corpus/services/studio-room.md:426`** — `exporter.py:513-550` for `_create_export_package`;
   the function spans **514-550** (`513` is a blank line). Off-by-one at the head only.

5. **`corpus/services/studio-room.md:61`** — the Project Structure tree is rooted at `studio-room/`,
   which the same file later corrects at `:337` ("studio-room's root IS `app/studio/`"). It also
   omits `tests/`, `tools/`, `pytest.ini`, `changelog.md`, `cog.toml`. (`workspace/` and `tools/`
   are in `studio/.gitignore`, which is why they are absent from a fresh clone — worth a parenthetical
   so a reader isn't surprised the tree's `workspace/` doesn't exist yet.)

6. **`corpus/services/roadrunner.md:33`** — cites `../architecture/architecture_overview.md:188` as
   a twin for "jobsimulation … remains in `docker-compose.yml:83` … starts on a bare `make up`".
   Line **188** is the **Skiller** row (`| **Skiller** | Merged into Backend/App (July 2026) — repo
   legacy/decommissioned |`). The row that actually carries the claim — "**Container still starts
   locally** (`docker-compose.yml:83`, default profile) as an unfederated husk" — is line **189**.
   Off-by-one anchor; the claim is true. (The companion citation `README.md:20-21` resolves to the
   sibling `corpus/services/README.md:20-21` and **is correct** — it reads "*three of the four (cms,
   jobsimulation, roadrunner) still start CONTAINERS locally in the default `graphql` profile*".)

7. **`corpus/services/roadrunner.md:21`** — "(M247 re-grepped `app` + `jobsimulation` on the
   consolidated clones — **zero hits outside CHANGELOG**)". There is one non-CHANGELOG hit for
   `RoadRunnerService` in `app`:
   `knowledge/plan/releases/07.00-jobsim-in-app/RE-PORT-CHECKLIST.md:10`. The load-bearing half —
   "no `ROADRUNNER_RPC_ADDR` / `RoadRunnerService` / `roadrunner:10401` read in any service's **Go
   code**" — is TRUE and re-verified this pass: 0 Go hits in `app`, 0 in `jobsimulation`.

8. **`corpus/services/roadrunner.md:15-16`** — "`roadrunner/terraform/main.tf:19` still reads
   `service_desired_count = 1` and **has not been touched since `87d8d44`** (2026-06-19, before the
   fold)". `87d8d44` is the repo's **HEAD** and is a CI commit that does not touch terraform at all;
   the last commit to touch `terraform/main.tf` is `e45eb61` (2026-05-27). Every asserted fact is
   true (`= 1` at line 19, exact; untouched; pre-fold) — the commit is just attributed to the wrong
   change. Reads more cleanly as "the repo has not been touched since `87d8d44`".

9. **`corpus/services/clerkenstein.md:3`** — "**Last updated:** 2026-07-14", and the status line
   stops at "M218 the roster-aware fake BAPI". The body documents M219 (`:51-58`), M220 (`:301-308`),
   M224 (`:228-251`), v2.8 M256 (`:156-205`) and M257x iter-23 (`:270-275`). Header is stale against
   its own contents.

10. **`corpus/services/clerkenstein.md:18`** — enumerates the monorepo as having sections
    "(`clerkenstein`, `demo-stack`, `stack-injection`, `stack-core`, `stack-seeding`, `alignment`)".
    The authoring copy has **11**: those six plus `dev-stack`, `playthroughs`, `stack-secrets`,
    `stack-snapshot`, `stack-verify`. Reads as a complete enumeration. (Rosetta's own `CLAUDE.md`
    lists 9 — also incomplete, omitting `playthroughs` and `stack-secrets`; same drift, wider.)

11. **`corpus/services/clerkenstein.md:101`** — the `cmd/` table row lists "`mintpk` … `fake-fapi` /
    `fake-bapi`". `clerkenstein/cmd/` also contains **`jwtkey`**.

12. **`corpus/services/clerkenstein.md:42-44`** — a markdown table row is wrapped across three
    blockquote lines (the cell text continues on `> now ExitRegressed …` / `> as a missing Node
    module. |`). A table row cannot wrap in markdown, so this row will not render as a table row.
    Formatting only; the content is correct (`run.go:134-135` verified exact).

13. **`corpus/services/gotenberg.md:14`** — "**Profile**: `graphql` (default) and `backend`".
    `platform/docker-compose.yml:384` reads `profiles: [graphql, backend, all]` — the `all` profile
    is omitted. (The same file's compose excerpt at `:20-27` also renders `command:` as a block
    sequence where the real file uses a flow sequence — cosmetic, semantically identical YAML.)

14. **`corpus/services/chronos.md:5`** and **`corpus/services/intelligence.md:5`** — both say the
    service was removed "in **mid-2026**". Both commits are dated **2026-04-17**
    (`platform` `045857c` "chore: remove chronos service from orchestration and update related
    documentation"; `fdfa189` "chore: remove intelligence service from local dev orchestration") —
    early Q2, not mid-year. The jobsimulation companion `09631fb2` is the same day.

15. **`corpus/services/chronos.md:22,171,191`** — "**Codebase**: `stack-dev/chronos`" plus
    `cd stack-dev/chronos` in both the Local Development and Testing blocks, while the banner at
    `:9` correctly states the repo "is no longer cloned by `make init`". No `chronos` directory
    exists under `stack-dev/` or `stack-demo/`. The instructions are unrunnable as written (they
    need a manual `git clone` first) — the section is explicitly historical, so harmless, but a
    one-line "clone it first" would close the gap.

16. **`corpus/services/chronos.md:9`** — "have moved to **in-process Asynq** running inside
    jobsimulation". Correct as of the mid-2026 removal, but since **jobsim-in-app** that engine runs
    inside `app` (`app/internal/jobsimulation/`, wired by `internal/jobsimwiring/`) — there is no
    jobsimulation process to be "inside" on a current stack. Same class of stale-but-historical
    framing as `intelligence.md:8`, which *does* carry the corresponding "(Note: skiller has since
    been merged into app …)" parenthetical; chronos has no equivalent.

---

### Note for the next iteration

These seven files are in unusually good shape — eight passes in, I could not find a single false
mechanism, inverted guard, wrong number, or contradicted twin sentence in any of them, and several
of the riskiest claims (the `colony v0.34.3` vs `v0.35.2` deployment-DNA drift; the "one seat per
stack" single-tenancy disclosure; the "only one Dockerfile"/"8 locales"/"`proxy.ts` not
`middleware.ts`" corrections against `next-web-app`'s own `CLAUDE.md`; the roadrunner
orphaned-not-absent framing incl. the 10→9 `repos.yml` count) are exactly right down to the line
number. The 16 minors above are anchors, enumerations, one stale header and one broken table row.
The highest-value two to fix are **#1** and **#2** (both in `studio-room.md`), because they sit on
the provider surface this audit was pointed at: the authoritative provider row is correct, but a
reader skimming `:350` can still walk away thinking Mistral is wired, and `:388` understates the
service's egress.
