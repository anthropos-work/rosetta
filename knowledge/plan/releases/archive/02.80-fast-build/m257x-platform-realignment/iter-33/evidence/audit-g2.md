# iter33 KB-fidelity audit — group G2

Read-only audit. Ground truth: `iter33-groundtruth.md`. Platform clone
`/Users/marco/workspace/anthropos/rosetta/stack-demo/platform` @ `2adcf71`, peer repos
`stack-demo/{app,cms,jobsimulation,sentinel,storage,messenger,roadrunner,studio-desk,next-web-app}`
(next-web-app @ `bb3313bc0` v2.133.0, studio-desk @ `14a5442` v0.152.4).

## 1. Positive control

| File | `wc -l` | Read |
|---|---|---|
| `corpus/services/ai-readiness.md` | 590 | read to line 590 (full) |
| `corpus/services/studio-desk.md` | 419 | read to line 419 (full) |
| `corpus/services/studio-room.md` | 414 | read to line 414 (full) |
| `corpus/architecture/shared_libraries.md` | 219 | read to line 219 (full) |

No file was sampled, grepped-only, or truncated.

## 2. Findings

### BLOCKER-1 — `corpus/architecture/shared_libraries.md:69`

> `| **Version pin** | one pin per live repo — the cms/jobsimulation version skew disappeared with the merge (they no longer have their own `go.mod`) |`

**False at HEAD.** Both husk repos still carry their own Go modules, and the skew is *larger* than
before, not gone:

| repo | `github.com/anthropos-work/proto` |
|---|---|
| `app` | `v1.210.0` |
| `messenger` | `v1.210.0` |
| `cms` | `v1.207.0` |
| `jobsimulation` | `v1.205.0` |
| `sentinel` | `v1.200.0` |
| `storage` / `roadrunner` | `v1.196.0` |

`stack-demo/cms/go.mod` and `stack-demo/jobsimulation/go.mod` both exist; `platform/repos.yml:14-19`
still clones both; `docker-compose.yml:144` / `:83` still build them. The same doc contradicts itself
five lines earlier — `:41` says "the still-running `cms` + `jobsimulation` husk containers → **v0.35.1**.
Measured from each repo's `go.mod`". Acting on `:69` (e.g. bumping proto, or reasoning about who must
rebuild on a contract change) misdirects.

**Grade: BLOCKER.**
**Fix (one line):** replace with "one pin per repo, and the skew is live — app/messenger `v1.210.0`,
cms `v1.207.0`, jobsimulation `v1.205.0`, sentinel `v1.200.0`, storage/roadrunner `v1.196.0`; the husk
repos still have their own `go.mod`."

---

### BLOCKER-2 — `corpus/architecture/shared_libraries.md:70` and `:79`

> `:70` — "…(app, sentinel, storage, messenger; the cms / jobsimulation / skiller / skillpath / **roadrunner** RPC surfaces are all served in-process by app)"
> `:79` — "`SkillPathSessionService` (served by app since the skillpath merge, M502→M507),"

**False at HEAD for `skillpath` and `roadrunner`.** `app/main.go` registers exactly six Connect-RPC
handlers:

```
main.go:1178 usersv1connect.NewUsersServiceHandler
main.go:1179 organizationsv1connect.NewOrganizationsServiceHandler
main.go:1187 skillerv1connect.NewSkillerServiceHandler
main.go:1195 jobsimulationv1connect.NewJobSimulationServiceHandler
main.go:1204 cmsv1connect.NewCMSServiceHandler
main.go:1219 labv1connect.NewLabSessionServiceHandler
```

There is **no** `SkillPathSessionService` and **no** `RoadRunnerService` handler anywhere in `app`
(`grep -r "skillpathv1connect\|roadrunnerv1connect" app` → 0 hits). Those two surfaces were **removed**,
not re-hosted:

- `app/internal/skillpaths/skillpaths.go:27-31` — "`sessionReader` is the narrow read the loopback
  consumer needs from the in-process skillpath SessionManager — **the drop-in for the removed skillpath
  RPC client** (skillpath-in-app M506)". It is a Go interface call, not an RPC surface.
- roadrunner: `app/internal/jobsimwiring/wiring.go:118` constructs `jsrunner.NewRunnerManager(JUDGE0_API_KEY,
  JUDGE0_BASE_URL)` — direct Judge0 HTTP. `ROADRUNNER_RPC_ADDR=http://roadrunner:10401` is still set on
  `backend` (`docker-compose.yml:118`) but is **read by no Go code in `app`**.

Contrast `skiller`, where the doc's claim *is* true (`main.go:1187`) — the asymmetry is the drift.
A consumer wired to call `SkillPathSessionService` or `RoadRunnerService` on `backend` finds nothing.

**Grade: BLOCKER.**
**Fix (one line):** drop `skillpath`/`roadrunner` from the "served in-process by app" list and mark
`SkillPathSessionService` + `RoadRunnerService` as contract-still-in-proto-but-no-longer-served (like
`ChronosService`), noting skillpath is now an in-process interface and roadrunner is direct Judge0 HTTP.

---

### BLOCKER-3 — `corpus/services/ai-readiness.md:259-291` (§ *Surfaces (UI) — current vs legacy*)

> `:261-265` — "⚠️ **There are TWO manager dashboards. Only one of them is the product.** … Nothing ever failed, because the legacy page *does* render. … **Establish which surface you are on before you conclude anything about AI readiness.**"
> `:270` — "| **Manager** | `AIReadinessContainer` → `AIReadinessView` — pre-v3.0 org-summary card + team table. … | `/enterprise/workforce/ai-readiness` | ❌ **LEGACY** |"
> `:279-281` — "**The legacy route is an orphan**: no nav entry, no workforce tab (`WorkforceNewClient.tsx:125-151` omits it), no redirect points at it. Its hook (`useWorkforceAIReadiness.ts:23-27`) calls `GET /api/workforce/ai-readiness?tag=`"
> `:282-283` — "The `(new)` in the legacy path is a Next.js **route group** … Don't read it as 'the new one'."

**False at HEAD — the legacy surface no longer exists.** It was deleted from `next-web-app` at commit
`dae0fb2f7` ("fix(ai-readiness): … **drop orphaned container**", 2026-07-13), which removed all three files:

```
packages/ui/src/organisms/Workforce/AIReadinessContainer.tsx        -103
packages/ui/src/.../Workforce/AIReadinessView/AIReadinessIntro.tsx  -220
packages/ui/src/.../Workforce/AIReadinessView/AIReadinessView.tsx   -330
```

At `bb3313bc0` (v2.133.0): `grep -r AIReadinessContainer` → **0 hits**; the only `ai-readiness` route dir
is `apps/web/src/app/(authenticated)/(verified)/ai-readiness`; `enterprise/workforce/(new)/` contains only
`succession/` and `trends/`; `packages/ui/src/Workforce/AIReadinessView/` retains only
`StepsCompletionDrawer.tsx` (consumed by the *current* client at `AIReadinessClient.tsx:6`). The hook
`apps/web/src/hooks/useWorkforceAIReadiness.ts` survives with **zero consumers** (dead code).

So there is exactly **one** manager dashboard, `/ai-readiness`. The section's central instruction —
"establish which surface you are on" — now sends a reader hunting for a route that 404s, and the two
code anchors offered to "tell them apart" (`WorkforceNewClient.tsx:125-151`, the legacy route) no longer
describe anything. This is unfenced present tense, not narrated history (the *historical* references at
`:401-416`, which explain that the M219 iter-07 probe watched the legacy page, are correctly fenced as
past and are **not** part of this finding).

**Grade: BLOCKER.**
**Fix (one line):** collapse the table to the two current surfaces and rewrite the ⚠️ callout as history —
"the legacy `AIReadinessContainer`/`AIReadinessView` was deleted at next-web-app `dae0fb2f7` (2026-07-13);
`/ai-readiness` is now the only manager surface, and `useWorkforceAIReadiness.ts` is orphaned dead code."

---

### minor-1 — `corpus/services/ai-readiness.md:164`, `:165` (repeated at `:241-242`)

> "`ai_readiness_notification_log` (**net-new, M400**) | notification send log"
> "`ai_readiness_notification_optout` (**net-new, M400/M403**) | per-member unsubscribe"

Both table names are singular; the shipped tables are **plural**:
`app/terraform/migrations/20260722160000_m400_ai_readiness_notifications.sql:4` →
`CREATE TABLE "ai_readiness_notification_logs"`, `:25` → `CREATE TABLE "ai_readiness_notification_optouts"`.
Fails loudly rather than silently (a query errors), and the same doc elsewhere makes a point of the
ent-generated plural (`:157` `ai_readiness_user_step_progresses`), so this is inconsistency rather than
misdirection.

**Grade: minor.**
**Fix (one line):** pluralize to `ai_readiness_notification_logs` / `ai_readiness_notification_optouts`
in all four places.

---

### minor-2 — `corpus/services/ai-readiness.md:277`, `:290`, `:409`, `:534` (line-anchor drift)

> `:277` "`AI_READINESS_URL` (`packages/core-js/src/constants/urls.ts:50`)" → actual `:52`
> `:290` "`AIReadinessClient.tsx:69` `const SHOW_SECONDARY_TABS = false;`" → actual `:78`
> `:409` "(`AIReadinessClient.tsx:137-138`) computes `effectiveCycleId`… gates the data GET on `cyclesQ.isFetched` (`:150-154`)" → actual `:153` and `:169`
> `:534` "`HowWeMeasureTab.tsx:2773-2797`" → **out of range**; the file is 1989 lines, the tile block is `:1915-1927`
> `:534` "`useAIReadiness.ts:250`" → actual `:274`

Every underlying *claim* verified TRUE (SHOW_SECONDARY_TABS is `false`; the current client does compute
`effectiveCycleId` and gate on `cyclesQ.isFetched`; the tile renders `skillsMapped` / `handsOnMinutes` /
`interviewMinutes` and `interviewQuestions` is typed-but-undrawn). Only the anchors drifted; one is past EOF.

**Grade: minor.**
**Fix (one line):** re-pin the five anchors against next-web-app `bb3313bc0`.

---

### minor-3 — `corpus/architecture/shared_libraries.md:42`

> "**Imported by** | **Every** live Go service: app, sentinel, storage, messenger — plus the `cms` and `jobsimulation` containers, which the default `graphql` profile still starts as merged-into-`app` **husks**"

The enumeration omits the **third** husk. `roadrunner` also imports colony (`roadrunner/go.mod` →
`colony v0.34.3`), is also started by the default profile (`docker-compose.yml:309`
`profiles: [graphql, roadrunner, all]`), and is also a merged-into-`app` husk with the same teardown
milestone. (The claim that the `graphql` **profile** still exists and still starts the husks is TRUE —
the profile survives on `backend`/`jobsimulation`/`cms`/`storage`/`roadrunner`/`gotenberg`; only the
`graphql-wundergraph` **service** was deleted — so that half is not a finding.)

**Grade: minor.**
**Fix (one line):** add `roadrunner` (`colony v0.34.3`) to the husk list in both the pin row and the
imported-by row.

---

### minor-4 — `corpus/services/studio-room.md:336`, `:338-344` (incomplete harden-6 correction)

> `:336` "Orchestration is performed by the **CMS Go code**, not by studio-room itself."
> `:338` "#### With CMS Service"
> `:339` "The CMS service drives the full lifecycle:"

Unfenced present tense naming a **service** that no longer exists as a process. The orchestrator is the
cms **domain inside `backend`** — `app/internal/cms/studio/studioManager.go:119`
(`s.runCommand(ctx, pyBin, append([]string{"studio/gen.py"}, tokens...))`), driven by the asynq worker in
`app/internal/cms/worker/worker.go`. The `cms` husk container still runs (`docker-compose.yml:144`), so a
reader who takes "the CMS service" literally will attach to the **wrong container** when debugging a
generation job. The standing ⚠ banner at `:3-11` fences the doc overall — which is why this is not graded
BLOCKER — but this section is the residue the harden-6 sweep did not reach.

**Grade: minor.**
**Fix (one line):** rename the heading to "With the cms domain (inside `backend`)" and change "The CMS
service drives" → "The cms domain in `backend` drives".

---

### minor-5 — `corpus/services/studio-room.md:61` vs `:287-291` (inconsistency left by the correction)

> `:61` — the Project Structure tree is rooted at "`studio-room/`"
> `:287-291` — "studio-room's root **IS** `app/studio/`. There is no `studio/studio-room` path."

The tree root contradicts the Installation section's own (correct) statement, and the repo is named
`anthropos-studio-room` while its in-image path is `app/studio/` (CI: `app/.github/workflows/build-production.yml:29`
`additional_repo: "anthropos-studio-room:studio"`). Verified on disk: `stack-demo/app/studio/` holds
`gen.py`, `postgen.py`, `agents/`, `services/`, `configs/`, `requirements.txt` exactly as the tree lists.
Answer to the tasking question: **the correction is substantively complete but cosmetically inconsistent** —
this line and minor-4 are the two loose ends, and neither reintroduces a live-pipeline reading.

**Grade: minor.**
**Fix (one line):** re-root the tree at `app/studio/` (repo `anthropos-studio-room`).

---

### minor-6 — `corpus/services/studio-room.md:82-84`

> "`├── configs/            # Environment configs` / `│   ├── local_config.ini` / `│   └── production_config.ini`"

The shipped `configs/` at HEAD is `config_template.ini`, `development_config.ini`, `production_config.ini`.
`local_config.ini` is **gitignored** (`app/studio/.gitignore` → `configs/local_*`) so it is absent from a
fresh checkout, and the tree omits the two files that *are* shipped. The doc is self-consistent
elsewhere (`:234` correctly says `configs/local_*` and `configs/test_*` are gitignored; `:316` tells you to
create it), and `gen.py:31-33` does default `ENVIRONMENT` to `local` — so this is a tree-listing gap, not a
false mechanism.

**Grade: minor.**
**Fix (one line):** list `config_template.ini` + `development_config.ini` + `production_config.ini` and
annotate `local_config.ini` as "gitignored, create locally".

---

### minor-7 — `corpus/services/studio-room.md:231`

> "`max_tokens = 4096`"

Both shipped configs set `max_tokens = 4000` (`app/studio/configs/production_config.ini:21`,
`development_config.ini:21`). Presented inside a block labelled as the config's contents.

**Grade: minor.**
**Fix (one line):** change `4096` → `4000`.

## 3. Per-file verdicts

**`corpus/services/studio-desk.md` — CLEAN. No findings.** Every checkable claim verified TRUE at HEAD,
including the ones most at risk from the `graphql-wundergraph` drop, which this doc has already swept
correctly:
- `:21` `depends_on` **backend + cms** at `docker-compose.yml:337-341` — exact match; `profiles: [studio-desk, all]` at `:342`; ports 9000/9100 at `:321-323`.
- `:21`/`:132` `VITE_GRAPHQL_ENDPOINT=http://localhost:8082/graphql/query` baked at `docker-compose.yml:334` (and build arg `:318`) — exact.
- `:46`/`:124`/`:156`/`:280-284` "backend directly; there is no `graphql` service locally" — correct; `graphql-wundergraph` is absent from both `repos.yml` and `docker-compose.yml`.
- `:24` cloned by `make init` — `repos.yml` lists `studio-desk` (`type: node-npm`).
- `:34` the 9 prod entry points + dev-only `dev-accept.html` — exact match to `vite.config.ts` `rollupOptions.input` (incl. the `isProduction ? {} : {'dev-accept'}` guard).
- `:99` `/api/skillpath` "~61KB" — `src/routes/skillpath.ts` is 61,080 bytes; `youtube.ts` present.
- `:104` tiers/models — `src/services/ai/config.ts:20-40`: OpenAI+Azure `gpt-5.2` / `gpt-5-mini` / `gpt-5-nano`; Anthropic `claude-opus-4-5-20251101` / `claude-sonnet-4-5-20250929` / `claude-haiku-4-5-20251001`. `.env.example:57,61` `AI_PROVIDER_CHAIN=azure-openai,openai`, `AI_DEFAULT_TIER=fast`.
- `:153` Node — `package.json` `engines.node >=24`, `Dockerfile`/`Dockerfile.dev` `FROM node:24-alpine`.
- `:164` in-code PORT fallback 9100 — `src/index.ts:60` `const backendPort = process.env.PORT || 9100;`.
- `:289` `STUDIO_ACCESS_ROLES = ['admin','org:admin','content_creator','org:content_creator']` — `src/index.ts:96`, consumed `:115`, redirect `:118-119`.
- `:371-378` boot sequence — `app/core/main.ts:97` `preloadCriticalCSS()`, `:104` `Sentry.init`, `:117` `posthog.init`, `:182` `clerk.load()`, `:191` `l12nService.init()`, `:199` `userService.canAccess()`, `:206` `new PageWrapper()`. Both quoted anchors (L97, L206) exact.
- (Cosmetic only, not raised as a finding: `:156` "Access to CMS service" in Prerequisites carries the same merged-service phrasing as minor-4, but the same line already pins the real dependency — "`backend` on `:8082`".)

**`corpus/architecture/shared_libraries.md`** — 2 BLOCKERS (1, 2) + 1 minor (minor-3). Everything else
verified TRUE: colony pins app/messenger `v0.35.2`, cms/jobsimulation `v0.35.1`, sentinel/storage `v0.34.3`;
`ai v1.40.2` across app+husks; `taxonomy v1.2.0`, direct in app+messenger, indirect in sentinel+storage;
**no repo imports the standalone `anthropos-work/authn`** (0 hits across all 7 Go repos) while `colony/authn`
has 129 importing files in `app` — the `:160-165` correction stands; the `:182-190` "taxonomy is a library,
not data" correction stands.

**`corpus/services/ai-readiness.md`** — 1 BLOCKER (3) + 2 minors (1, 2). The platform-side body is
otherwise strong and re-verified true at HEAD: `internal/aireadiness/` exists with exactly the claimed file
set and `internal/workforce/members.go` retains `LoadMembers`; `scoring.go:26,30` `archetypeHighBand = 75` /
`archetypeLowCeil = 50` with the `:61` "Medium starts at 51" comment; `defaults.go` = **31** skills
(19 @1.0 + 12 @0.5) + **3** track-keyed sims (`tech`/`business`/`both`); `csv.go:13-33` = **15** columns with
the BOM at `:44`; `enum/organization_settings.go` `OrganizationSettingAIReadiness = "ai_readiness"`;
`enum/ai_readiness.go:9-11` the three step types; `how_we_measure.go:1139-1155` the five **current** KPI ids
(`avg_adoption`/`avg_transformation`/`avg_originality`/`avg_depth`/`avg_ownership`) — the retired
`avg_frequency`/`avg_breadth`/`avg_context_fit` appear nowhere; `computeCycleTotals` at `:260` with the
`FROM public.interactions i JOIN public.job_simulation_sessions s` query at `:285-286`;
`CloseDueAIReadinessCycles` at `cycles.go:554`; `RefreshLiveSnapshots` at `live_snapshots.go:71`;
`interview_aggregated_reports` created at `terraform/migrations/20260722081626_jobsim_data_model.sql:47`;
the GraphQL ops (`aiReadinessEnabled` `:73`, `aiReadinessUserPlanProgress` `:78`, `aiReadinessSkills` `:85`,
`completeAiReadinessSkillMapping` `:8`) and the REST surface incl. `/cycles/{cycleID}`, `/cycles/{cycleID}/close`
and the net-new `/setup` (`api/server.gen.go:84-99`).

**`corpus/services/studio-room.md`** — 0 BLOCKERS, 4 minors (4, 5, 6, 7). **The harden-6 correction is
substantively complete**: the ⚠ banner (`:3-11`), the High-Level Summary close (`:25`), the Deployment row
(`:35`) and the Repository row (`:37`) all now read as merged-and-embedded, and every merge-side fact checks
out — `app/.github/workflows/build-production.yml:29` `additional_repo: "anthropos-studio-room:studio"`,
`app/Dockerfile:28` + `Dockerfile.dev:26` `FROM python:3.11-slim`, `internal/cms/worker/worker.go:29-34`
`Concurrency: 5` with `ai_video: 7` / `studio: 3`, `studioManager.go:94` `studio/studio-venv` and `:119`
`studio/gen.py` (postgen at `:1045`). Also verified: `services/taxonomy.py:11` `BASE_URL =
"https://api.anthropos.work/api"` (the sole outbound call, `:336`); `postgen.py:396-399` `--media`/`--simid`/
`--target` all `required=True`; the four post-gen targets plus `testing.py` and a separate `exporter.py`
(`agents/simulation/postgen/`) — matching `:158-167`; `worklog_path = workspace/trace` in both shipped configs
(`:262`); no top-level `templates/` dir (`:244`); `requirements.txt` unpinned and exactly the 8 listed packages
with no `aiohttp` (`:295-307`). It did **not** reintroduce a live-pipeline reading; the two loose ends are the
"CMS service" heading (minor-4) and the tree root (minor-5).

## 4. Totals

- **BLOCKERS: 3** — `shared_libraries.md:69`; `shared_libraries.md:70,79`; `ai-readiness.md:259-291`.
- **minors: 7** — `ai-readiness.md` ×2, `shared_libraries.md` ×1, `studio-room.md` ×4.
- **Files not fully read: none.** All 4 read top-to-bottom (1642 lines total).
- **Verification gap to disclose:** `colony`, `proto`, `taxonomy` and standalone `authn` are private Go
  modules not present on disk (`GOMODCACHE` holds only `ai@v1.40.1`), so `shared_libraries.md`'s claims about
  those modules' *internals* — the colony package table (`:48-54`), the GraphQL rate-limiter note (`:56-59`),
  the "12 Connect-RPC services are defined" count (`:78-83`), the buf-v1 layout note (`:91-94`), the `ai.AI`
  method/constructor tables (`:107-116`), and the NodeID format/examples (`:201-206`) — could be checked only
  from the consumer side. Nothing there contradicted the consumers; BLOCKER-2 is a claim about **who serves
  what in `app`**, which *is* fully verifiable from the app clone, and is the one I found false.
