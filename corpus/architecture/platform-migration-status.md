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

> **Re-measured 2026-08-03 against platform origin HEAD [`ef32d4c`](https://github.com/anthropos-work/platform)**
> ("Merge pull request #24 … chore/prune-merged-services", 2026-08-03). First measured 2026-08-01 at
> `2adcf71`. Re-run [§4 of the protocol](../ops/platform-alignment.md#4--detection--six-signals-cheapest-first)
> before trusting any row older than a release.
>
> **`2adcf71 → ef32d4c` moved three rows**, and the fence in [§4](#4-the-fence) found it — direction B, three
> departures, unprompted, on a tree nobody had touched. `d11a403` deletes the **cms**, **jobsimulation** and
> **roadrunner** compose services *and* their `repos.yml` entries: those three are now decommissioned locally
> and are no longer cloned by `make init`. **This is the removal [§5](#5-what-this-map-says-about-the-program)
> named as the one that arms the failure this milestone exists to fence** — it landed under
> `chore/prune-merged-services`, ahead of the M810 it was expected to wait for.

---

## 1. How to read a row

| state | means |
|---|---|
| `live-standalone` | its own process, still on the traffic path |
| `merged-into-app` | `app` owns the code and calls it unconditionally, the tables live in `public`, and the standalone is scaled to zero — **all three**, per [§6 of the protocol](../ops/platform-alignment.md#6--classification--the-map) |
| `running_but_unfederated` | the container still starts, but it owns no schema and is not a subgraph — a **husk**, not a service |
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
| `app` | live-standalone | live-standalone | yes | the monolith. `app/terraform/main.tf:44` `service_desired_count = 1`; **the only migrating repo** — `repos.yml:10-13` (`migrations: true`, `schema: public`); compose service is named `backend` (`docker-compose.yml:28`). Owns four domains in-process, each with **its own** wiring call site — skiller `app/main.go:573`, jobsimulation `:604`, skillpath `:634`, cms `:1034` (`app/internal/{cms,jobsimulation,skiller,skillpath}/`, `app` @ `5ba17044` v1.363.2). An earlier revision attached `:604` alone to all four, where it wires jobsimulation only; corrected M257x iter-46. **`app/internal/roadrunner/` does not exist** — the Judge0 runner was absorbed as `app/internal/jobsimulation/runner/`, constructed at `app/internal/jobsimwiring/wiring.go:118` (`jsrunner.NewRunnerManager(JUDGE0_API_KEY, JUDGE0_BASE_URL)`) |
| `cms` | merged-into-app | decommissioned | no | `cms/terraform/main.tf:39` `service_desired_count = 0`; code in `app/internal/cms/`; folded by platform `236771f` (2026-07-29, cms-in-app v8.0). **Compose service and `repos.yml` entry both deleted by `d11a403`** (merged `ef32d4c`, 2026-08-03) — `make init` no longer clones it. Repo **not** archived; `repos.yml:9-10` names infrastructure's `services.tf` as the prod rollback path until M810 |
| `jobsimulation` | merged-into-app | decommissioned | no | `jobsimulation/terraform/main.tf:40` `service_desired_count = 0`; code in `app/internal/jobsimulation/`, wired unconditionally at `app/main.go:604` (`jobsimwiring.Wire`); tables re-created in `public`; folded by platform `236771f`. **Compose service and `repos.yml` entry both deleted by `d11a403`.** **Repo ARCHIVED on GitHub 2026-07-31** |
| `roadrunner` | live-standalone | decommissioned | no | **Compose service and `repos.yml` entry both deleted by `d11a403`** — whose message states the clone entry *"was already gone, so the `../roadrunner` build context could no longer resolve"*, i.e. the service had been unbuildable, not merely legacy. Judge0 is reached directly: `JUDGE0_BASE_URL` moved onto `backend` (`docker-compose.yml:56`) for `app/internal/jobsimulation/runner/` (`app/internal/jobsimwiring/wiring.go:118`). **The prod contradiction is now explained but still not verified:** `roadrunner/terraform/main.tf:19` remains `service_desired_count = 1`, untouched since `87d8d44` (2026-06-19), while `repos.yml:9-10` says the authoritative rollback declaration lives in **infrastructure's `services.tf`** — a repo this map has never read. Repo **not** archived |
| `sentinel` | live-standalone | live-standalone | yes | `sentinel/terraform/main.tf:19` `= 1`; `docker-compose.yml:5`, own `sentinel` schema via `search_path=sentinel` (`:18`) **despite `migrations: false`** (`repos.yml:15-17`) — the Trap-A row |
| `storage` | live-standalone | live-standalone | yes | `storage/terraform/main.tf:19` `= 1`; `docker-compose.yml:90`; `repos.yml:18-20`. **Named as the next fold** — `app` PR #1103 (v9.0 "support-in-app") folds storage + messenger |
| `messenger` | live-standalone | live-standalone (opt-in profile) | yes | `messenger/terraform/main.tf:19` `= 1`; `docker-compose.yml:141`, `messenger` profile (`:178`) — not started by the default `graphql` profile; `repos.yml:21-23`. **The only surviving service that still talks to cms/jobsimulation as RPC peers**, and `d11a403` repointed both edges at the monolith: `CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR` now read `http://backend:8083` (`docker-compose.yml:159`, `:161`) instead of the dead `cms:8091` / `jobsimulation:8401`, because `messenger/cmd/root.go:120-140` genuinely reads all four addrs and `backend` registers `CMSService` + `JobSimulationService` on its own mux. Also in the v9.0 fold |
| `next-web-app` | external (Vercel) | live-standalone | yes | `repos.yml:26-28`; `docker-compose.yml:211` (`frontend` profile, `:236`). Points at `backend` directly since the router drop — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=…:8082/graphql/query` (`docker-compose.yml:228`) |
| `studio-desk` | live-standalone | live-standalone | yes | `repos.yml:29-31`; `docker-compose.yml:180` (`studio-desk` profile, `:209`) |
| `graphql-wundergraph` | live-standalone | decommissioned | no | **the router, dropped from local dev mid-milestone.** Deleted from `repos.yml` **and** compose by `b56d731` + `360efd4`, merged `2adcf71` (2026-07-31); local dev now points at `backend`. In prod it is still declared — `graphql-wundergraph/terraform/main.tf:20` `= 1` — while the **repo is ARCHIVED on GitHub 2026-07-30**. Supergraph is **one** subgraph: `supergraph-config-prod.yaml` lists `backend` alone, `schemas/` holds `backend.graphqls` alone, `subgraphs.conf` = `BACKEND=v1.360.0` (folded by `915da06`, 2026-07-29) |
| `skiller` | merged-into-app | decommissioned | no | removed from compose + `repos.yml` by platform `21429b7` (2026-07-07); code in `app/internal/skiller/`; taxonomy data in `public`. **Repo ARCHIVED 2026-07-01** |
| `skillpath` | merged-into-app | decommissioned | no | decommissioned by platform `a4db680` (2026-07-21, M507); code in `app/internal/skillpath/`; session state in `public.skill_path_sessions`. **Repo ARCHIVED 2026-07-31** |
| `chronos` | decommissioned | decommissioned | no | removed from orchestration by platform `045857c` (2026-04-17). **Repo is NOT archived on GitHub** (last push 2026-04-23) — the corpus called it archived; the org disagrees |
| `intelligence` | decommissioned | decommissioned | no | removed from orchestration by platform `fdfa189` (2026-04-17). **Repo ARCHIVED 2026-04-02** |
| `customerio-sync` | live-standalone | live-standalone (opt-in profile) | no | never cloned — compose builds it straight from the git URL (`docker-compose.yml:121-123`, `context: git@github.com:anthropos-work/customerio-sync.git#main`), `customerio-sync` profile (`:139`) |
| `db-backup` | live-standalone | — | no | production-only; no compose service and no `repos.yml` entry at `ef32d4c` |
| `anthropos-studio-room` | merged-into-app | merged-into-app | no | the Python generation pipeline is pulled into the `app` image by CI and orchestrated from `app/internal/cms/studio/`, which spawns it as a subprocess. Not a deployment, not in `repos.yml`. **The repo name is `anthropos-studio-room`, not `studio-room`** |
| `ant-academy` | external (Vercel) | external | no | deliberately absent from `repos.yml` — run natively, never containerised |
| `gotenberg` | external | live-standalone | no | third-party image, `docker-compose.yml:238-239` (`gotenberg/gotenberg:8`), default `graphql` profile (`:251`) |
| `colony` | library | library | no | private Go module (framework + `colony/authn`); pulled at Docker build via `GH_PAT`/`GOPRIVATE`, never cloned by `make init` |
| `proto` | library | library | no | private Go module — RPC contracts + domain types |
| `ai` | library | library | no | private Go module — the multi-provider `ai.AI` wrapper |
| `authn` | library | library | no | legacy standalone module; the live copy ships **inside** colony as `colony/authn` |
| `taxonomy` | library | library | no | private Go module — the `NodeID` type only. **Not** the 60K-skill dataset, which lives in `app`'s `public` schema |
| `postgresql` | external | external | no | the shared database. Not in `docker-compose.yml` at all — it lives in the **included** `common.yml:2` (`docker-compose.yml:1-2`, `include: - common.yml`), which is why a top-level grep of the compose file finds no database. Its healthcheck gained a `start_period: 120s` at `6060315` (`common.yml:22`) because permission re-application on a grown data dir outlasted the 25 s the retries allowed — a **bring-up-timing** change, so any cold-cycle timing baseline taken before `ef32d4c` is measuring a different startup contract |
| `redis` | external | external | no | `common.yml:20`. Streams transport for the Watermill pub/sub |
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
count), of which **9** are in `repos.yml` and **46** are named by no corpus document at all. The table below
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
| C | every state is one of the seven in §1 | a row invents a state, so the vocabulary stops meaning anything |
| D | every row cites evidence | a claim with no sha and no `file:line` |
| E | no §3 census row is in `repos.yml` | the census silently overlaps the clone set it is defined as the complement of |

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
v2.0 skiller → v5.0 skillpath → v7.0 jobsimulation → v8.0 cms → **v9.0 `storage` + `messenger`, PR #1103 open**.
Two of the rows above are therefore already known to be wrong on a schedule.

The rows to watch, in order:

1. **`storage` and `messenger`** — the named next fold. When `repos.yml` flips either to `migrations: false`
   with a `legacy` comment, the fold has landed. `messenger` is the more exposed of the two: it is the last
   process that still calls cms and jobsimulation over RPC, and `d11a403` had to repoint both edges by hand.
2. **`roadrunner`** — still the only row where a repo's own terraform and the platform's declaration
   disagree, and now the disagreement is one this map cannot settle without reading **infrastructure**.
3. ~~**`cms` / `jobsimulation` local husks**~~ — **this happened, on 2026-08-03, ahead of M810.** `d11a403`
   removed the containers *and* the clone entries under `chore/prune-merged-services`. Two consequences,
   both live:
   - **The armed failure is now armed.** Tooling that iterates the clone set — `demo-stack/migrate-demo.sh:81-85`
     creates the legacy schemas itself and `:106` atlas-applies a hand-maintained 4-tuple, guarded by
     `[ -d ] || continue` — will now **silently skip** three repos rather than fail. This is the exact shape
     M257x iter-01 predicted and named a time bomb.
   - **A box with stale clones cannot observe it.** The three directories still exist on any machine that
     cloned before 2026-08-03, so the skip is invisible there and fires only on a genuinely fresh `make init`.
     A local "it still works" is not evidence about a cold box.

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
