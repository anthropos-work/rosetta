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

> **Measured 2026-08-01 against platform origin HEAD [`2adcf71`](https://github.com/anthropos-work/platform)**
> ("Merge pull request #23 … chore/drop-wundergraph", 2026-07-31), with every peer repo's clone verified
> `behind=0` versus its own origin. Re-run [§4 of the protocol](../ops/platform-alignment.md#4--detection--six-signals-cheapest-first)
> before trusting any row older than a release.

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

**Two traps this table exists to keep straight** (both live at `2adcf71`):

- **`migrations: false` entails nothing on its own.** `sentinel` is `migrations: false` *and* alive *with its
  own `sentinel` schema* (`docker-compose.yml:18`, `search_path=sentinel`). Read the `prod` and
  `fresh local stack` columns, never the flag alone.
- **The declared topology and the actual topology disagree by design.** `repos.yml:14-31` calls cms /
  jobsimulation / roadrunner *"legacy — folded into app"*, while `docker-compose.yml` still defines all three
  **in the default `graphql` profile** — so a fresh local stack still starts them. Docs merged, compose
  deferred.

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
| `app` | live-standalone | live-standalone | yes | the monolith. `app/terraform/main.tf:44` `service_desired_count = 1`; **the only migrating repo** — `repos.yml:10-13` (`migrations: true`, `schema: public`); compose service is named `backend` (`docker-compose.yml:28`). Owns cms · jobsimulation · roadrunner · skiller · skillpath in-process (`app/internal/{cms,jobsimulation,roadrunner,skiller,skillpath}/`, wired at `app/main.go:604`, `app` @ `5ba17044` v1.363.2) |
| `cms` | merged-into-app | running_but_unfederated | yes | `cms/terraform/main.tf:39` `service_desired_count = 0`; code in `app/internal/cms/`; folded by platform `236771f` (2026-07-29, cms-in-app v8.0) — **but** `docker-compose.yml:144` still defines the service in the default `graphql` profile, and `repos.yml:14-16` marks it `migrations: false # legacy`. Repo **not** archived |
| `jobsimulation` | merged-into-app | running_but_unfederated | yes | `jobsimulation/terraform/main.tf:40` `service_desired_count = 0`; code in `app/internal/jobsimulation/`, wired unconditionally at `app/main.go:604` (`jobsimwiring.Wire`); tables re-created in `public`; folded by platform `236771f`. Container still starts (`docker-compose.yml:83`). **Repo ARCHIVED on GitHub 2026-07-31** |
| `roadrunner` | live-standalone | running_but_unfederated | yes | **contradiction, recorded not resolved:** `repos.yml:29-31` says *"legacy — folded into app; backend calls Judge0 directly"*, while `roadrunner/terraform/main.tf:19` still reads `service_desired_count = 1` — that file has not been touched since `87d8d44` (2026-06-19), before the fold. Container still starts (`docker-compose.yml:281`). Repo **not** archived |
| `sentinel` | live-standalone | live-standalone | yes | `sentinel/terraform/main.tf:19` `= 1`; `docker-compose.yml:5`, own `sentinel` schema via `search_path=sentinel` (`:18`) **despite `migrations: false`** (`repos.yml:20-22`) — the Trap-A row |
| `storage` | live-standalone | live-standalone | yes | `storage/terraform/main.tf:19` `= 1`; `docker-compose.yml:189`; `repos.yml:23-25`. **Named as the next fold** — `app` PR #1103 (v9.0 "support-in-app") folds storage + messenger |
| `messenger` | live-standalone | live-standalone (opt-in profile) | yes | `messenger/terraform/main.tf:19` `= 1`; `docker-compose.yml:240`, `messenger` profile — not started by the default `graphql` profile; `repos.yml:26-28`. Also in the v9.0 fold |
| `next-web-app` | external (Vercel) | live-standalone | yes | `repos.yml:34-36`; `docker-compose.yml:344` (`frontend` profile). Points at `backend` directly since the router drop — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=…:8082/graphql/query` (`docker-compose.yml:352`) |
| `studio-desk` | live-standalone | live-standalone | yes | `repos.yml:37-39`; `docker-compose.yml:311` (`studio-desk` profile) |
| `graphql-wundergraph` | live-standalone | decommissioned | no | **the router, dropped from local dev mid-milestone.** Deleted from `repos.yml` **and** compose by `b56d731` + `360efd4`, merged `2adcf71` (2026-07-31); local dev now points at `backend`. In prod it is still declared — `graphql-wundergraph/terraform/main.tf:20` `= 1` — while the **repo is ARCHIVED on GitHub 2026-07-30**. Supergraph is **one** subgraph: `supergraph-config-prod.yaml` lists `backend` alone, `schemas/` holds `backend.graphqls` alone, `subgraphs.conf` = `BACKEND=v1.360.0` (folded by `915da06`, 2026-07-29) |
| `skiller` | merged-into-app | decommissioned | no | removed from compose + `repos.yml` by platform `21429b7` (2026-07-07); code in `app/internal/skiller/`; taxonomy data in `public`. **Repo ARCHIVED 2026-07-01** |
| `skillpath` | merged-into-app | decommissioned | no | decommissioned by platform `a4db680` (2026-07-21, M507); code in `app/internal/skillpath/`; session state in `public.skill_path_sessions`. **Repo ARCHIVED 2026-07-31** |
| `chronos` | decommissioned | decommissioned | no | removed from orchestration by platform `045857c` (2026-04-17). **Repo is NOT archived on GitHub** (last push 2026-04-23) — the corpus called it archived; the org disagrees |
| `intelligence` | decommissioned | decommissioned | no | removed from orchestration by platform `fdfa189` (2026-04-17). **Repo ARCHIVED 2026-04-02** |
| `customerio-sync` | live-standalone | live-standalone (opt-in profile) | no | never cloned — compose builds it straight from the git URL (`docker-compose.yml:220-222`, `context: git@github.com:anthropos-work/customerio-sync.git#main`), `customerio-sync` profile (`:238`) |
| `db-backup` | live-standalone | — | no | production-only; no compose service and no `repos.yml` entry at `2adcf71` |
| `anthropos-studio-room` | merged-into-app | merged-into-app | no | the Python generation pipeline is pulled into the `app` image by CI and orchestrated from `app/internal/cms/studio/`, which spawns it as a subprocess. Not a deployment, not in `repos.yml`. **The repo name is `anthropos-studio-room`, not `studio-room`** |
| `ant-academy` | external (Vercel) | external | no | deliberately absent from `repos.yml` — run natively, never containerised |
| `gotenberg` | external | live-standalone | no | third-party image, `docker-compose.yml:371-372` (`gotenberg/gotenberg:8`) |
| `colony` | library | library | no | private Go module (framework + `colony/authn`); pulled at Docker build via `GH_PAT`/`GOPRIVATE`, never cloned by `make init` |
| `proto` | library | library | no | private Go module — RPC contracts + domain types |
| `ai` | library | library | no | private Go module — the multi-provider `ai.AI` wrapper |
| `authn` | library | library | no | legacy standalone module; the live copy ships **inside** colony as `colony/authn` |
| `taxonomy` | library | library | no | private Go module — the `NodeID` type only. **Not** the 60K-skill dataset, which lives in `app`'s `public` schema |
| `postgresql` | external | external | no | the shared database. Not in `docker-compose.yml` at all — it lives in the **included** `common.yml:2` (`docker-compose.yml:1-2`, `include: - common.yml`), which is why a top-level grep of the compose file finds no database |
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
   with a `legacy` comment, the fold has landed.
2. **`roadrunner`** — the only row where prod and the platform's own declaration contradict each other.
3. **`cms` / `jobsimulation` local husks** — `running_but_unfederated` until platform M810 removes the
   containers and the repos from the clone set. That removal is exactly what arms the failure this milestone
   exists to fence: tooling that iterates the clone set silently skips what is no longer cloned.

---

## See also

- [`corpus/ops/platform-alignment.md`](../ops/platform-alignment.md) — the procedure that produces and
  maintains this map: detection signals, the re-point steps, the three fence layers.
- [`corpus/architecture/service_taxonomy.md`](service_taxonomy.md) — the three-tier categorisation.
- [`corpus/services/README.md`](../services/README.md) — the per-service docs this map is the index of truth for.
