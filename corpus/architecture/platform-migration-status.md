# Platform migration status — every service, its state, cited

**What this is.** The single place that records *where the microservice-into-`app` consolidation actually is*.
Every service the platform has ever had gets one row, and **every row carries two states — one for
production, one for a fresh local stack** — because those genuinely differ and collapsing them is why the
drift class in [`corpus/ops/platform-alignment.md`](../ops/platform-alignment.md) recurred three times.

**Every claim is cited to a sha or a `file:line`.** A row with no evidence is not a claim, and the fence
below rejects it.

**It is machine-fenced against the platform's own `repos.yml`, both ways** — see [§4](#4-the-fence). A
service that enters or leaves the clone set without a row here turns the guard RED. That is the whole point:
this file is *allowed* to be out of date only for as long as it takes CI to say so.

> **Re-measured 2026-08-04 against platform origin HEAD [`0dab54d`](https://github.com/anthropos-work/platform)**
> (local clone and origin level; `app` @ `b948604` v1.366.0). The prior reading was 2026-08-03 at `ef32d4c`
> ("Merge pull request #24 … chore/prune-merged-services"); first measured 2026-08-01 at `2adcf71`. Re-run [§4 of the protocol](../ops/platform-alignment.md#4--detection--six-signals-cheapest-first)
> before trusting any row older than a release.
>
> **`2adcf71 → ef32d4c` moved three rows**, and the fence in [§4](#4-the-fence) found it — direction B, three
> departures, unprompted, on a tree nobody had touched. `d11a403` deletes the **cms**, **jobsimulation** and
> **roadrunner** compose services *and* their `repos.yml` entries: those three are now decommissioned locally
> and are no longer cloned by `make init`. **This is the removal [§5](#5-what-this-map-says-about-the-program)
> named as the one that would have armed the failure this milestone exists to fence** — it landed under
> `chore/prune-merged-services`, ahead of the M810 it was expected to wait for, and it landed *after* the
> fix. See §5 row 3: the tooling had already been re-derived, so the removal passed through it harmlessly.

---

## 1. How to read a row

| state | means |
|---|---|
| `live-standalone` | its own process, still on the traffic path |
| `merged-into-app` | `app` owns the code and calls it unconditionally, the tables live in `public`, and the standalone is scaled to zero — **all three**, per [§6 of the protocol](../ops/platform-alignment.md#6--classification--the-map) |
| `running_but_unfederated` | the container still starts, but it owns no schema and is not a subgraph — a **husk**, not a service |
| `mid-fold` | **a half-landed fold: the config side says removed and the consumer side says live.** Neither `live-standalone` nor `merged-into-app` — and it is recorded on **both** sides, cited, or not at all. Added M257x iter-64; the gap it closes was stated in [§6 of the protocol](../ops/platform-alignment.md#6--classification--the-map) at iter-59. **No row carries it today** (M257x iter-68): its only holder, `storage`, completed its fold four iterations later. The token stays in the vocabulary — the fold program is not finished, and a state you can only name *after* you need it is the state you will get wrong |
| `decommissioned` | gone from the orchestration; the repo may still exist as a rollback reference |
| `net-new` | exists in the org, is not in `repos.yml`, and the corpus has never described it |
| `external` | third-party or separately-deployed; never in the local Go clone set |
| `library` | imported as a private Go module, never a process |

**Two traps this table exists to keep straight:**

- **`migrations: false` entails nothing on its own.** `sentinel` is `migrations: false` *and* alive *with its
  own `sentinel` schema* (`docker-compose.yml:18`, `search_path=sentinel`). Read the `prod` and
  `fresh local stack` columns, never the flag alone. Live at `ef32d4c`.
- **Absent from `repos.yml` no longer means "never was a service".** Until `d11a403` the declared topology
  and the actual one disagreed — `repos.yml` called cms / jobsimulation / roadrunner *"legacy"* while compose
  still started all three. `d11a403` closed that gap by deleting both sides at once, so the three now look
  exactly like `skiller` and `skillpath`: no row in `repos.yml`, no compose service, **repo still on GitHub as
  the pre-merge reference**. The clone set is therefore no longer a census of what the platform has ever run —
  which is what this file is for.

---

## 2. The services

**Completeness is measured, not asserted.** The row set is the union of *every name that has ever appeared in
`repos.yml`* (`git log -p --follow -- repos.yml` → **14** names: app · chronos · cms · graphql-wundergraph ·
intelligence · jobsimulation · messenger · next-web-app · roadrunner · sentinel · skiller · skillpath ·
storage · studio-desk — all 14 have rows) and *every service that has ever appeared in `docker-compose.yml`*
(same command on that file → 26 names, including the pre-history the clone set never knew: `nats`,
`web-app`, `chromedp`, `simulator`, `realtime`). Re-run those two commands to audit this table; a name they
return that has no row is a gap.

<!-- fence:services:begin -->

| repo | prod | fresh local stack | in `repos.yml` | evidence |
|---|---|---|---|---|
| `app` | live-standalone | live-standalone | yes | the monolith. `app/terraform/main.tf:44` `service_desired_count = 1`; **the only migrating repo** — `repos.yml:11-14` (`migrations: true`, `schema: public`); compose service is named `backend` (`docker-compose.yml:28`). Owns **six** domains in-process — the four folded before v9.0 plus storage and messenger — each with **its own** wiring call site: skiller `app/main.go:637` (`skiller.NewSkillerManager`), jobsimulation `:668` (`jobsimwiring.Wire`), skillpath `:698` (`skillpath.NewSessionManager`), cms `:1099` (`appcms.Wire`), storage `:471` (`internalstorage.NewManager`), messenger `:1423` (`msgadapters.Wire`) (`app/internal/{cms,jobsimulation,skiller,skillpath,storage,messenger}/`, `app` @ `9d00a313` v1.367.0 — **re-resolved M257x iter-68**; the first four stood at `:581`/`:612`/`:643`/`:1043` at `b948604` v1.366.0 and at `:573`/`:604`/`:634`/`:1034` at `5ba17044` v1.363.2, which is what one working day costs a line-number citation). An earlier revision attached the jobsimulation site alone to all four, where it wires jobsimulation only; corrected M257x iter-46. **`app/internal/roadrunner/` does not exist** — the Judge0 runner was absorbed as `app/internal/jobsimulation/runner/`, constructed at `app/internal/jobsimwiring/wiring.go:123` (`jsrunner.NewRunnerManager(JUDGE0_API_KEY, JUDGE0_BASE_URL)`) |
| `cms` | merged-into-app | decommissioned | no | `cms/terraform/main.tf:39` `service_desired_count = 0`; code in `app/internal/cms/`; folded by platform `236771f` (2026-07-29, cms-in-app v8.0). **Compose service and `repos.yml` entry both deleted by `d11a403`** (merged `ef32d4c`, 2026-08-03) — `make init` no longer clones it. Repo **not** archived; `repos.yml:9-10` names infrastructure's `services.tf` as the prod rollback path until M810 |
| `jobsimulation` | merged-into-app | decommissioned | no | `jobsimulation/terraform/main.tf:40` `service_desired_count = 0`; code in `app/internal/jobsimulation/`, wired unconditionally at `app/main.go:612` (`jobsimwiring.Wire`, @ `app` `b948604` v1.366.0); tables re-created in `public`; folded by platform `236771f`. **Compose service and `repos.yml` entry both deleted by `d11a403`.** **Repo ARCHIVED on GitHub 2026-07-31** |
| `roadrunner` | live-standalone | decommissioned | no | **Compose service and `repos.yml` entry both deleted by `d11a403`, in that one commit.** Its message says the clone entry *"was already gone, so the `../roadrunner` build context could no longer resolve"* — **the message is wrong; the diff is the fact.** `git show d11a403 -- repos.yml` shows that very commit deleting `- name: roadrunner` alongside `- name: cms` and `- name: jobsimulation`, and the compose file at `d11a403^` still declares a `roadrunner:` service block (it was one of eleven there; eight remain). The service was legacy, not unbuildable. (An earlier revision of this row promoted that message into a conclusion — *"the service had been unbuildable, not merely legacy"*; corrected against the diff in M257x. **A commit message is testimony, not evidence** — grade a change by its diff.) Judge0 is reached directly: `JUDGE0_BASE_URL` moved onto `backend` (`docker-compose.yml:59`) for `app/internal/jobsimulation/runner/` (`app/internal/jobsimwiring/wiring.go:123` @ `app` `9d00a313` v1.367.0 — it was `:118` at `b948604`, which is the line `d11a403`'s own message quotes). **The prod contradiction is now explained but still not verified:** `roadrunner/terraform/main.tf:19` remains `service_desired_count = 1` — last changed at **`84a4b4f` (2025-12-15)**, the commit that first added `terraform/main.tf`, and untouched by everything up to the repo's HEAD `87d8d44` (2026-06-19). **That count is not a decision about the fold; it predates it by seven months and nobody has been back.** (An earlier revision of this row dated it to `e45eb61` (2026-05-27) — that commit is the file's most recent touch but it changed **line 11 only**, a one-line module-source URL swap whose own message says *"Module contents are identical; this is a pure source-URL swap"*. `git blame -L 19,19` names `84a4b4f`; a file-level `git log` is not line provenance.) Meanwhile `repos.yml:9-10` says the authoritative rollback declaration lives in **infrastructure's `services.tf`** — a repo this map has never read. Repo **not** archived |
| `sentinel` | live-standalone | live-standalone | yes | `sentinel/terraform/main.tf:19` `= 1`; `docker-compose.yml:5`, own `sentinel` schema via `search_path=sentinel` (`:18`) **despite `migrations: false`** (`repos.yml:15-17`) — the Trap-A row |
| `storage` | merged-into-app | merged-into-app (startable via `storage-legacy`) | yes | **The v9.0 fold COMPLETED between M257x iter-64 and iter-68** — this row read `mid-fold` for four iterations and the half it was waiting on landed in one working morning (re-derived M257x iter-68 at platform `0dab54d` / `app` `9d00a313` v1.367.0 / `storage` `63bffc8`). **Prod — stopped:** `storage/terraform/main.tf:38` `service_desired_count = 0`, following the cms precedent, with the ordering constraint in-comment (`:29-33` — it must land *after* app is confirmed reading/writing S3, because scaling to 0 empties the Cloud Map record). The buckets, CloudFront distribution and media DNS record are **not** touched — declared in the same module under `prevent_destroy`, custody transfer is M903 (`:35-37`). **Config side — removed:** `STORAGE_RPC_ADDR` occurs **0** times across `docker-compose.yml`, `common.yml` and `.env_example`; `STORAGE_S3_BUCKET`/`STORAGE_S3_PUBLIC_BUCKET` were added to `backend` (`docker-compose.yml:82`); the standalone moved to `profiles: [storage-legacy]` (`:134`) with the rationale in-comment (`:131-133` — two writers on one bucket), so `make up` no longer starts it. **Consumer side — also removed:** `app` serves object storage in-process — `internalstorage.NewManager` / `NewPublicManager` at `app/main.go:471`, `:472`, consumed at `:494` and `:1048`; the constants at `app/internal/storage/service.go:22`, `:24` are the bucket **env-var NAMES** (`EnvBucket = "STORAGE_S3_BUCKET"`, `EnvPublicBucket = "STORAGE_S3_PUBLIC_BUCKET"`) — the bucket names themselves are still read from the environment, at `app/main.go:463`, `:464`, and guarded there against empty (`:465-469`). **`STORAGE_RPC_ADDR` is read by nothing at this ref** — a Go grep across `9d00a313` returns **3 hits, every one of them a comment** (`app/main.go:451`, `app/internal/jobsimwiring/wiring.go:101`, `app/internal/storagens/callsites_test.go:189`), and the first says it in words: *"the standalone service takes no traffic and STORAGE_RPC_ADDR is gone."* **State the ref or the sentence flips:** at the OLDER `b948604` (v1.366.0 — the checkout this document's banner names) it is genuinely read, in `main.go` **and in all three `cmd/` tools** — `git -C stack-demo/app grep -n STORAGE_RPC_ADDR b948604 -- '*.go'` returns **15** hits, of which **7 are env lookups** — `main.go` ×3, `internal/jobsimwiring/wiring.go` ×1 and one in each of the three `cmd/` tools (the other 8 are two doc comments, two error strings, **three** `t.Setenv` calls and **one `t.Fatal` message** — the fourth test-file hit *names* the variable inside a failure string rather than setting it, so counting all four as `t.Setenv` overstates the arrange side by one). Six of the seven are spelled `os.Getenv`; the seventh uses a lowercase `getenv` helper, which is why a regex fitted to `os.Getenv` reports **six** and is not wrong to. **Line anchors for that side are deliberately omitted:** they resolve at `b948604` and nowhere else, and naming a second ref inside this cell makes every anchor in it ungradeable (the citation resolver reads one ref per block and reports `ambiguous` when a block names two — M257x run-53). Run the grep; it is the ref, not the line number, that carries the claim. That interval — `b948604` → `9d00a313` — *is* the consumer half of the fold, and it is the evidence that moved this row off `mid-fold`. `repos.yml:18-20` still clones the repo and the standalone stays startable — the rollback path, exactly as `cms` and `jobsimulation` are kept |
| `messenger` | merged-into-app | merged-into-app (startable via the `messenger` profile) | yes | **Folded in the same v9.0 program as `storage`, and it landed the same morning** (re-derived M257x iter-68 at platform `0dab54d` / `app` `9d00a313` v1.367.0 / `messenger` `a0ec933`). **Prod — stopped:** `messenger/terraform/main.tf:29` `service_desired_count = 0`, the cms precedent again; the image and task definition stay declared as the rollback path (`:27-28`). **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:61`, `:62`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1387`, wired at `:1423` with `msgsender.NewFromEnv`) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap (`:1330-1340`). The group name is messenger's and is a deliberate literal (`:1362-1365`). Local: `docker-compose.yml:156`, `messenger` profile (`:195`) — not started by the default `core` profile, and `0dab54d` also dropped it from `all`, because running both puts two consumers on one group; `repos.yml:21-23` still clones it. It was **the last service still talking to cms/jobsimulation as RPC peers** — `d11a403` had repointed both edges at the monolith (`CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` → `http://backend:8083`, `docker-compose.yml:174`, `:176`) because `messenger/cmd/root.go:120-140` genuinely reads all four |
| `next-web-app` | external (Vercel) | live-standalone | yes | `repos.yml:26-28`; `docker-compose.yml:228` (`frontend` profile, `:253`). Points at `backend` directly since the router drop — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=…:8082/graphql/query` (`docker-compose.yml:245`) |
| `studio-desk` | live-standalone | live-standalone | yes | `repos.yml:29-31`; `docker-compose.yml:197` (`studio-desk` profile, `:226`) |
| `graphql-wundergraph` | live-standalone | decommissioned | no | **the router, dropped from local dev mid-milestone.** Deleted from `repos.yml` **and** compose by `b56d731` + `360efd4`, merged `2adcf71` (2026-07-31); local dev now points at `backend`. In prod it is still declared — `graphql-wundergraph/terraform/main.tf:20` `= 1` — while the **repo is ARCHIVED on GitHub 2026-07-30**. Supergraph is **one** subgraph: `supergraph-config-prod.yaml` lists `backend` alone, `schemas/` holds `backend.graphqls` alone, `subgraphs.conf` = `BACKEND=v1.360.0` (folded by `915da06`, 2026-07-29) |
| `skiller` | merged-into-app | decommissioned | no | removed from compose + `repos.yml` by platform `21429b7` (2026-07-07); code in `app/internal/skiller/`; taxonomy data in `public`. **Repo ARCHIVED 2026-07-01** |
| `skillpath` | merged-into-app | decommissioned | no | decommissioned by platform `a4db680` (2026-07-21, M507); code in `app/internal/skillpath/`; session state in `public.skill_path_sessions`. **Repo ARCHIVED 2026-07-31** |
| `chronos` | decommissioned | decommissioned | no | removed from orchestration by platform `045857c` (2026-04-17). **Repo is NOT archived on GitHub** (last push 2026-04-23) — the corpus called it archived; the org disagrees |
| `intelligence` | decommissioned | decommissioned | no | removed from orchestration by platform `fdfa189` (2026-04-17). **Repo ARCHIVED 2026-04-02** |
| `customerio-sync` | live-standalone | live-standalone (opt-in profile) | no | never cloned — compose builds it straight from the git URL (`docker-compose.yml:136-138`, `context: git@github.com:anthropos-work/customerio-sync.git#main`), `customerio-sync` profile (`:154`) |
| `db-backup` | live-standalone | — | no | production-only; no compose service and no `repos.yml` entry at `ef32d4c` |
| `anthropos-studio-room` | merged-into-app | merged-into-app | no | the Python generation pipeline is pulled into the `app` image by CI and orchestrated from `app/internal/cms/studio/`, which spawns it as a subprocess. Not a deployment, not in `repos.yml`. **The repo name is `anthropos-studio-room`, not `studio-room`** |
| `ant-academy` | external (Vercel) | external | no | deliberately absent from `repos.yml` — run natively, never containerised |
| `gotenberg` | external | live-standalone | no | third-party image, `docker-compose.yml:255-256` (`gotenberg/gotenberg:8`), default `core` profile (`:268` — renamed from `graphql` by `0dab54d`, since the WunderGraph router the name described is gone) |
| `colony` | library | library | no | private Go module (framework + `colony/authn`); pulled at Docker build via `GH_PAT`/`GOPRIVATE`, never cloned by `make init` |
| `proto` | library | library | no | private Go module — RPC contracts + domain types |
| `ai` | library | library | no | private Go module — the multi-provider `ai.AI` wrapper |
| `authn` | library | library | no | legacy standalone module; the live copy ships **inside** colony as `colony/authn` |
| `taxonomy` | library | library | no | private Go module — the `NodeID` type only. **Not** the 60K-skill dataset, which lives in `app`'s `public` schema |
| `postgresql` | external | external | no | the shared database. Not in `docker-compose.yml` at all — it lives in the **included** `common.yml:2` (`docker-compose.yml:1-2`, `include: - common.yml`), which is why a top-level grep of the compose file finds no database. Its healthcheck gained a `start_period: 120s` at `6060315` (`common.yml:22`) because permission re-application on a grown data dir outlasted the 25 s the retries allowed — a **bring-up-timing** change, so any cold-cycle timing baseline taken before `ef32d4c` is measuring a different startup contract |
| `redis` | external | external | no | `common.yml:24`. Streams transport for the Watermill pub/sub |
| `directus` | external | external | no | the headless CMS at `content.anthropos.work`. **Removed from compose** at `a2a3ee6` (2026-02-27); a local stack gets its own only via rext's `--local-content` cutover, never from the platform repo |
| `chromedp` | decommissioned | decommissioned | no | headless-Chrome renderer, present from the first commit `cb6ebf5` (2023-04-30), last touched `ef4b449` (2024-09-02) |
| `simulator` | decommissioned | decommissioned | no | added `1474b1f` (2023-12-01), **replaced by `jobsimulations`** at `84862d1` (2024-05-29) — the first ancestor of what is now `app/internal/jobsimulation/` |
| `realtime` | decommissioned | decommissioned | no | added `b43b99a` (2026-02-26), gone by `c17cc9a` (2026-04-15). The org still holds an undocumented `realtime-python` repo (§3) |
| `web-app` | decommissioned | decommissioned | no | the pre-Next frontend service; first commit `cb6ebf5` (2023-04-30), reorganised out at `467965a` (2023-09-11) |
| `nats` | decommissioned | decommissioned | no | the original message bus, removed at `8770fe6` (2023-05-04) — four days after the first commit. Redis Streams replaced it |

<!-- fence:services:end -->

---

## 3. Net-new — in the org, in neither `repos.yml` nor the corpus

**93 repositories** in `anthropos-work` (GitHub org API, 2026-08-01 — independently reproducing iter-01's
count), of which **6** are in `repos.yml` and **46** are named by no corpus document at all. The 6 are
**enumerated from the file, not carried forward as a sum**: `app` · `sentinel` · `storage` · `messenger` ·
`next-web-app` · `studio-desk` (a 31-line `repos.yml` @ `0dab54d`). **It read 9 until `d11a403`** removed the
`cms` / `jobsimulation` / `roadrunner` entries; the 46 does not move with it, because all three *are* named
by corpus documents. The table below
is the subset that is **unarchived and pushed since 2025** — the ones that could plausibly matter and that
nothing in the corpus has ever looked for.

<!-- fence:census:begin -->

| repo | last push | why it matters |
|---|---|---|
| `kb-ant-product` | 2026-08-01 | knowledge-base repo; the org's most recently touched repo |
| `ant-observability` | 2026-07-31 | observability stack — the corpus documents no observability tier |
| `sim-qa` | 2026-07-31 | simulation QA tooling. iter-01 recorded that it **writes to prod** and is unmarked as such |
| `kb-certifications-iso27001` | 2026-07-07 | compliance KB; `security_compliance.md` cites no ISO-27001 programme |
| `livekit-agent-chain` | 2026-07-03 | one of **five** LiveKit agent repos. `ai_architecture.md` documents the LiveKit *engine* and **none** of the agents |
| `livekit-agent` | 2026-05-20 | (same family) |
| `livekit-agent-azure-us` | 2026-05-20 | (same family) |
| `livekit-agent-azure-eu` | 2026-04-22 | (same family) |
| `livekit-agent-azure-eu-fr` | 2026-06-03 | (same family) — iter-01 measured `azure-eu` + `azure-eu-fr` as dispatching nothing |
| `github-runner-config` | 2026-06-26 | CI runner configuration |
| `kb-migration-plan` | 2026-06-18 | a migration-planning KB — directly adjacent to this document's subject |
| `simulation-form` | 2026-06-14 | unknown surface |
| `customer-orbyta` | 2026-05-19 | a per-customer repo; the corpus describes no per-customer repo pattern |
| `kb-domain-singularity` | 2026-05-12 | knowledge-base repo |
| `bench-analysis-transcripts` | 2026-03-26 | benchmark/transcript analysis |
| `transcoder` | 2025-11-11 | media transcoding — adjacent to `media-substrate-spec.md`'s Bunny.net path |
| `realtime-python` | 2025-09-12 | realtime Python service |
| `studio-tools` | 2025-02-18 | studio tooling |
| `analytics-go` | 2025-02-12 | analytics service |

<!-- fence:census:end -->

**`auth` is not on this list, deliberately.** iter-01 measured it as a 3-commit, 8-hour spike marked *"Not yet
deployed"*, with all human activity stopped 2026-06-18 and BetterAuth a dependency of nothing in the org.
**Clerk is not being replaced.** Recording that here is the point: the scary-looking repo was measured, and
the measurement is the deliverable.

**`AI-Labs` is not on this list either** — it *is* named by the corpus (`corpus/services/ai-labs.md`), but as a
subsystem rather than as the live Go control plane repo it is. That is a fidelity gap, not a discovery gap.

---

## 4. The fence

`rosetta-extensions/stack-core/platform_alignment_guard.py` — layer 1 of the three in
[§8 of the protocol](../ops/platform-alignment.md#8--fence--so-it-cannot-silently-recur). It reads this file
and the platform's own `repos.yml` and asserts **both directions**:

| # | assertion | the miss it catches |
|---|---|---|
| A | every repo in `repos.yml` has a row here with `in repos.yml = yes` | a service **enters** the clone set and nobody writes it down |
| B | every row here claiming `in repos.yml = yes` is really in `repos.yml` | a service **leaves** (skiller, skillpath, the router) and the map keeps asserting it |
| C | every state is one of the **eight** in §1 (markdown emphasis stripped before the check) | a row invents a state, so the vocabulary stops meaning anything |
| D | every row cites evidence | a claim with no sha and no `file:line` |
| E | no §3 census row is in `repos.yml` | the census silently overlaps the clone set it is defined as the complement of |
| F | every citation resolves, and lands in **its own subject's compose block** | the platform edits `docker-compose.yml`, every line number below it shifts, and a citation slides into the neighbouring service — still resolving, still non-blank, silently about something else |

**Assertion F (M257x iter-57) exists because A–E were GREEN over a map with five false claims.** D checks
only that a row cites *something*; a cell full of citations to lines that no longer exist is, to D,
indistinguishable from a correct one. F parses `docker-compose.yml`'s own service blocks and its
`context:` lines — so `backend` is known to be the `app` repo because compose says so, not because this
guard was told — and asks whether each cited line sits in the block it claims. On first run it named
**8 dead citations**; a hand reading had found 2. A row may cite into *another* service's block when the
cell **names** that block (`roadrunner`'s Judge0 line genuinely lives in `backend`); declaring the wrong
block still fires.

### What is fenced, and what is prose-under-review

The protocol's rule is **derive, else fence, else declare it prose-under-review** — and the third category
only works if it is *visible*, which is why it is stated here rather than left implied:

| the claim | standing |
|---|---|
| `in repos.yml` membership, both directions | **derived + fenced** (A/B/E) — tracked the storage/messenger move with no human action |
| the state vocabulary | **fenced** (C) |
| `docker-compose.yml` citations — 20 of them | **fenced** (F) |
| every other `file:line` — terraform, `repos.yml`, `common.yml`, `app/…` (28 citations) | **resolution + range only.** No derivable notion of *whose* line it is, so F does not pretend |
| the prod column, PR/rollback narrative, §5's ordering | **prose-under-review.** Neither derived nor fenced. Re-check at each platform ref bump |

F's reach is printed on **every** run, GREEN or RED — a fence whose coverage shrinks in silence is the
failure this milestone has now found four times. A run that subject-checks nothing **refuses (exit 2)**
rather than reporting the map clean.

```bash
# from a rosetta checkout, against any stack's platform clone
PLATFORM_REPOS_YML=stack-demo/platform/repos.yml \
  python3 .agentspace/rosetta-extensions/stack-core/platform_alignment_guard.py
```

Exit 0 = aligned. Exit 1 = the drift is named, by repo, in the direction it drifted.

**Direction B is the one that has actually fired in anger.** All three occurrences of this class — skiller,
skillpath, jobsimulation — were a *departure* the corpus never noticed, not an arrival.

---

## 5. What this map says about the program

The consolidation is a **program with a published order**, not three accidents:
v2.0 skiller → v5.0 skillpath → v7.0 jobsimulation → v8.0 cms → **v9.0 `storage` + `messenger`, whose
compose half LANDED at `0dab54d` (2026-08-03)** — both now served in-process by `backend`, both moved out
of the default profile, both still startable for rollback.
Two of the rows above are therefore already known to be wrong on a schedule.

The rows to watch, in order:

1. **`storage` and `messenger`** — the named next fold. **The signal this row used to give was already
   dead when it was written:** it said *"when `repos.yml` flips either to `migrations: false`, the fold has
   landed"* — but both have read `migrations: false` since long before the fold was announced (`repos.yml:18-23`
   @ `ef32d4c`), exactly the Trap-A error §1 warns about. The signal that actually fires is **departure**:
   the row leaves `repos.yml` and the compose service is deleted, which is what direction B in [§4](#4-the-fence)
   watches and what caught `d11a403`. `messenger` is the more exposed of the two: it is the last process that
   still calls cms and jobsimulation over RPC, and `d11a403` had to repoint both edges by hand.
2. **`roadrunner`** — still the only row where a repo's own terraform and the platform's declaration
   disagree, and now the disagreement is one this map cannot settle without reading **infrastructure**.
3. ~~**`cms` / `jobsimulation` local husks**~~ — **this happened, on 2026-08-03, ahead of M810 — and the
   predicted failure did not occur, because the fix had already landed.** `d11a403` removed the containers
   *and* the clone entries under `chore/prune-merged-services`. What it met on our side:
   - **The time bomb was disarmed six weeks early, by this milestone.** M257x iter-01 named it:
     `migrate-demo.sh` created the legacy schemas itself and atlas-applied a hand-maintained
     `app:public cms:cms jobsimulation:jobsimulation skillpath:skillpath` tuple behind a silent
     `[ -d ] || continue`. iter-02 (rext `54bccf7`) replaced the tuple with a set **derived from `repos.yml`**
     and made an absent clone LOUD; iters 06 and 07 re-pointed the last cms/jobsimulation writes, emptying
     `REXT_TRANSITIONAL_SCHEMAS`. Measured against `repos.yml` @ `ef32d4c` on 2026-08-03: migration set =
     `app:public`, schema-create set = `extensions sentinel public`, transitional debt = **empty**. The
     derived set followed the platform's removal with **zero human action** — which is the whole thesis.
   - **What is still unproven is the cold path, not the logic.** The three directories still exist on any
     machine that cloned before 2026-08-03, so no local run exercises a genuinely fresh `make init` against
     this HEAD. A local "it still works" is not evidence about a cold box.

   > **This row asserted the opposite for one commit.** M257x iter-54 first wrote it up as *"the armed
   > failure is now armed"*, citing `migrate-demo.sh:81-85` / `:106` — line anchors and a code shape that
   > iter-02 had already deleted. The claim was quoted forward from iter-01 without re-measuring against
   > this milestone's own repair. It is the milestone's founding class, committed into the map built to stop
   > it, and the membership fence in §4 cannot see it: the fence checks who is in `repos.yml`, not whether
   > the prose about our own tooling is still true. Corrected the same day; recorded rather than erased.

4. **The merge itself keeps dropping configuration the merged code still reads.** `d11a403`'s own message
   records that deleting the three containers *"silently dropped env that `app` still reads in-process"* —
   `JUDGE0_BASE_URL`, `DIRECTUS_PUBLIC_BASE_ADDR`, `REDIS_WORKER_INDEX`, the LiveKit and Chime blocks and the
   `~/.aws/credentials` mount, all restored onto `backend`. Merging a service moves its **code**; its
   **environment** has to be moved separately, and nothing checks that it was.

---

## See also

- [`corpus/ops/platform-alignment.md`](../ops/platform-alignment.md) — the procedure that produces and
  maintains this map: detection signals, the re-point steps, the three fence layers.
- [`corpus/architecture/service_taxonomy.md`](service_taxonomy.md) — the three-tier categorisation.
- [`corpus/services/README.md`](../services/README.md) — the per-service docs this map is the index of truth for.
