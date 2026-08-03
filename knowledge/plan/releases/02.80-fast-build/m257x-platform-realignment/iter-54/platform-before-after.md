# The platform, before → now — the user's requested inventory

**Taken 2026-08-03 against platform origin HEAD `ef32d4cd8e0ceecf528a74c37d5e2ae5804ce021`.**
Our `stack-demo/platform` clone was at `2adcf714` (3 behind) at the start of this iteration and was
fast-forwarded to `ef32d4c`. Nothing was committed into it.

The canonical, fenced, per-row version of this is
[`corpus/architecture/platform-migration-status.md`](../../../../../corpus/architecture/platform-migration-status.md).
This file is the **narrative** the user asked for, and the record of **what changed in the last three
commits**.

---

## 1. The one-paragraph answer

**The consolidation is essentially done for the five named services, and as of 2026-08-03 the local stack
finally reflects it.** `app` (the compose service `backend`) is the monolith: skiller, skillpath, roadrunner,
jobsimulation and cms all run in-process. A default `make up` now starts **four Go services** — `backend`,
`sentinel`, `storage`, `gotenberg` — where six weeks ago it started nine plus a federation router. The
federation router itself is gone from local dev. The next fold, `storage` + `messenger`, is in an open PR.

---

## 2. Before → now, by service

| service | was | is now | evidence |
|---|---|---|---|
| **app** | one of several Go services | **the monolith**; the only repo with migrations; compose service `backend` | `repos.yml:11-14` (`migrations: true`, `schema: public`); `docker-compose.yml:28` |
| **skiller** | standalone (taxonomy, assessment, embeddings) | **merged into `app`**, repo ARCHIVED 2026-07-01 | removed from compose + `repos.yml` by `21429b7` (2026-07-07) |
| **skillpath** | standalone (progression runtime) | **merged into `app`**, repo ARCHIVED 2026-07-31 | `a4db680` (2026-07-21, M507) |
| **jobsimulation** | standalone simulation runtime | **merged into `app`**, repo ARCHIVED 2026-07-31; **as of `d11a403` no compose service and no clone entry** | folded `236771f` (2026-07-29); dropped `d11a403` (2026-08-03) |
| **cms** | standalone content layer + Studio | **merged into `app`**; **as of `d11a403` no compose service and no clone entry**; repo not archived | folded `236771f` (cms-in-app v8.0); dropped `d11a403` |
| **roadrunner** | standalone Judge0 proxy | **merged into `app`**; `backend` calls Judge0 directly; **dropped from compose + `repos.yml`** | `d11a403`; `JUDGE0_BASE_URL` now on `backend` (`docker-compose.yml:56`) → `app/internal/jobsimwiring/wiring.go:118` |
| **graphql-wundergraph** (Cosmo Router) | the federation gateway, local + prod | **gone from local dev entirely**; repo ARCHIVED 2026-07-30; still declared in prod | `b56d731` + `360efd4`, merged `2adcf71` (2026-07-31). Frontends point straight at `backend:8082/graphql/query` (`docker-compose.yml:228`) |
| **sentinel** | standalone authorization | **unchanged — still standalone**, still owns the `sentinel` schema | `docker-compose.yml:5`, `:18`; `repos.yml:15-17` |
| **storage** | standalone | **unchanged — but named as the next fold** (`app` PR #1103, v9.0 "support-in-app") | `docker-compose.yml:90`; `repos.yml:18-20` |
| **messenger** | standalone, opt-in profile | **unchanged in shape, repointed in wiring** — its cms + jobsimulation RPC edges now target the monolith | `CMS_RPC_ADDR` / `JOBSIMULATION_RPC_ADDR` = `http://backend:8083` (`docker-compose.yml:159`, `:161`) |
| **chronos**, **intelligence** | standalone | **decommissioned outright** (not merged) — removed 2026-04-17 | `045857c`, `fdfa189` |
| **next-web-app**, **studio-desk** | frontends | **unchanged** | `repos.yml:26-31` |

### Net-new

Nothing net-new entered `repos.yml` in this window. The net-new surface is **outside** the clone set and
unchanged from the iter-20 census: 93 org repos, 9 in `repos.yml` (now **6**), 46 named by no corpus doc —
the observability tier, five LiveKit agent repos, `sim-qa` (which writes to prod), `transcoder`,
`analytics-go`, `realtime-python` and the KB repos. See §3 of the migration-status map.

---

## 3. What changed in the last three commits — and why it matters more than its size

```
ef32d4c  Merge pull request #24 from anthropos-work/chore/prune-merged-services
6060315  fix(compose): give postgres a 120s healthcheck start_period
d11a403  chore(compose): drop roadrunner, prune dead env, repoint messenger
```

`d11a403` is 29 insertions / 170 deletions across three files. It does four things:

1. **Deletes the `cms`, `jobsimulation` and `roadrunner` compose services** and their `repos.yml` entries.
   `make init` no longer clones them. `repos.yml:7-10` now reads: *"Those five repos are frozen legacy: they
   own no local schema, no compose service and no clone entry here."*

2. **Restores environment the merged code still reads.** Its own message: deleting the containers *"silently
   dropped env that `app` still reads in-process"* — `JUDGE0_BASE_URL`, `DIRECTUS_PUBLIC_BASE_ADDR`,
   `REDIS_WORKER_INDEX`, the LiveKit and Chime blocks, and the `~/.aws/credentials` mount — all moved onto
   `backend`. **Merging a service moves its code; its environment has to be moved separately, and nothing
   checks that it was.**

3. **Repoints `messenger`**, the last process that still calls cms and jobsimulation over RPC, at
   `backend:8083`. It also drops `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR`, `SKILLER_RPC_ADDR`,
   `SIMULATOR_SESSIONS_ADDR` and `CHROME_DEBUGGER_URL` from `backend` as verified-dead.

4. **Adds `--remove-orphans`** to the `up` targets so the deleted services' containers get cleaned up.

`6060315` gives postgres a `start_period: 120s` (`common.yml:22`) because permission re-application on a
grown data dir outlasts the 25 s the retries allowed. **This changes the bring-up timing contract**, which
matters for any cold-cycle baseline taken before `ef32d4c`.

---

## 4. Two findings this survey produced

### The fence caught it, unaided, on a tree nobody had touched

Before any corpus edit, the clause-3 fence was run against the new HEAD:

```
platform_alignment_guard: 3 finding(s) …
  [B departure] the map claims cms is in repos.yml, and it is not …
  [B departure] the map claims jobsimulation is in repos.yml, and it is not …
  [B departure] the map claims roadrunner is in repos.yml, and it is not …
EXIT=1
```

After the map was updated: `EXIT=0`. **This is the first time a fence this milestone built has caught a real
departure it was not shown**, rather than a defect deliberately staged to watch it go RED. Direction B is the
direction that fired in anger all three prior occurrences (skiller, skillpath, jobsimulation) and it fired
again, correctly, within hours.

### The armed failure is now armed — and a stale box cannot see it

M257x iter-01 named the time bomb: `demo-stack/migrate-demo.sh:81-85` creates the legacy schemas itself and
`:106` atlas-applies a hand-maintained 4-tuple, guarded by `[ -d ] || continue`. **The condition it was
waiting for has now occurred**, ahead of the M810 everyone expected to wait for. Three repos are no longer
cloned, so that guard now *silently skips* rather than fails.

**And our box cannot observe it.** `stack-demo/{cms,jobsimulation,roadrunner}` still exist on disk — cloned
before 2026-08-03, last commits `ca50c81` / `462343b0` / `87d8d44`. The skip fires only on a genuinely fresh
`make init`. **A local "it still works" is not evidence about a cold box** — which is precisely how B1 and B2
went undetected for four days in M257.

---

## 5. The drift this creates in the corpus — measured, not repaired

Sites in `corpus/**` + root `CLAUDE.md` asserting the now-false shape (*"the container still starts"*,
*"`running_but_unfederated`"*, *"`repos.yml` still lists"*, and the three stale
`docker-compose.yml:{83,144,281}` anchors):

**81 sites across 21 files.**

```
CLAUDE.md
corpus/architecture/{ai_architecture,architecture_overview,dependency_map,external_services,
                     frontend_architecture,service_taxonomy,shared_libraries}.md
corpus/ops/{platform-alignment,update_guide}.md
corpus/ops/demo/content-stories-routes.md
corpus/services/{README,backend,cms,jobsimulation,messenger,roadrunner,sentinel,skillpath,
                 storage,studio-desk}.md
```

**This is the datum that decides TOK-04.** Three platform commits, landed inside one working day, created a
larger drift surface than the 46-item union this milestone has spent ten readings trying to close. The
vocabulary itself died: `running_but_unfederated` was coined at iter-20 for exactly these three services and
now describes none of them.

**Nothing in §5 was repaired in this iter.** It is measurement, deliberately left for the next tik under the
policy TOK-04 establishes.
