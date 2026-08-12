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

> ## ⚠️ RE-MEASURED 2026-08-12 at platform `766df6c` — **the 8th merge: `sentinel` is folded into `app`**
>
> **This is a new row state, not a re-check**, and the fence caught it exactly as designed: run against
> `766df6c` it went **`rc=1`, 17 findings — one `[B departure]` (*"the map claims sentinel is in
> repos.yml, and it is not"*) plus 16 citation failures.** Repaired at M258 iter-18; **`rc=0`** now.
> Every clone read is at `origin/main` **and verified equal to the remote the same minute** (`git
> ls-remote`, 2026-08-12T12:07Z): platform `766df6c` · `app` `c52dbc51e` · `sentinel` `f2c46190`.
>
> - **`repos.yml` is now THREE entries** — `app` · `next-web-app` · `studio-desk` — in a **13**-line file
>   (was 4 in 28 lines at `0c91421`). `make init` no longer clones `sentinel`.
> - **`docker-compose.yml` is 164 lines** (was 190 at `0c91421`, 271 at `0dab54d`) and declares **four**
>   services: `backend`, `studio-desk`, `next-web-app`, `gotenberg`. **`postgresql` and `redis` are not in
>   it at all** — they live in the **included** `common.yml`, and they carry **no `profiles:` key**, so the
>   always-on floor is those **two**. A `core` stack is therefore **four** containers, not five.
> - **`app` deleted its Connect-RPC listener** with the fold (`app/main.go:1310`), so a local stack has
>   **no cross-process Connect-RPC edge at all** — the `backend → sentinel` edge this corpus called *"the
>   only one left"* is gone, not re-pointed.
> - **Prod is `mid-fold`, and deliberately not `merged-into-app`** — the policy tables stay in the
>   `sentinel` schema and the ECS service is not at zero. See the `sentinel` row for the three-part test.
>
> **Two method findings from the repair, both worth more than the repair:** (1) a citation that drifts
> **within** a file's own line count and **into another service's block** is invisible to a reader and
> caught only by assertion F — 8 of the 16 were that shape; (2) **quoting a retracted citation verbatim
> re-arms it.** Writing *"this used to say `docker-compose.yml:102-103`"* in backticks makes the guard
> parse it as a live citation, and because the retraction usually also names the block the line drifted
> into, it **silently satisfies the declared-cross-block escape**. Measured on this edit: the fence went
> green while three retracted citations were still being graded. Retracted citations are now written
> without the `path:line` form.

> **Prior reading — 2026-08-05 against platform origin HEAD [`0c91421`](https://github.com/anthropos-work/platform)**
> (local clone and origin level; `app` **pinned** at `2035f9a`, post-v1.369.0 — a sha, not a moving label; see the `app` row). The prior readings were 2026-08-04 at
> `0dab54d`, and 2026-08-03 at `ef32d4c` ("Merge pull request #24 … chore/prune-merged-services"); first
> measured 2026-08-01 at `2adcf71`. Re-run [§4 of the protocol](../ops/platform-alignment.md#4-detection--six-signals-cheapest-first)
> before trusting any row older than a release.
>
> **`0dab54d → 0c91421` moved three more rows, and the fence named two of them the same day.** `838d907`
> (merged `0c91421`, 2026-08-05, PR #26 *"drop the storage, messenger and customerio-sync containers"*)
> **deletes those three compose services outright** — not parks them behind a legacy profile — and removes
> `storage` + `messenger` from `repos.yml`, which is direction B again: two departures, on a tree the map had
> read one day earlier. `repos.yml` was **4** entries (`app` · `sentinel` · `next-web-app` · `studio-desk`)
> and `docker-compose.yml` **186** lines, down from **271** at `0dab54d` (**both superseded at `766df6c`** —
> 3 entries / 164 lines; see the 2026-08-12 banner above) (`838d907` is +11/−96 on that
> file alone) — which is why nearly every compose citation in this file had to be re-derived rather than
> merely re-checked. The third row, `customerio-sync`, was never in
> `repos.yml` at all, so no membership assertion could see it; it moves `live-standalone → merged-into-app` on
> compose + `app` evidence alone, which is exactly what [§4](#4-the-fence)'s *prose-under-review* line is for.
> The commit's own framing names the hazard it removed: both were *"one `--profile` flag from a second writer
> on the prod bucket and a second consumer on messenger's Redis group"*, and `customerio-sync` *"was still in
> the `all` profile"*.
>
> **`2adcf71 → ef32d4c` moved three rows**, and the fence in [§4](#4-the-fence) found it — direction B, three
> departures, unprompted, on a tree nobody had touched. `d11a403` deletes the **cms**, **jobsimulation** and
> **roadrunner** compose services *and* their `repos.yml` entries: those three are now decommissioned locally
> and are no longer cloned by `make init`. **This is the removal [§5](#5-what-this-map-says-about-the-program)
> named as the one that would have armed the failure this milestone exists to fence** — it landed under
> `chore/prune-merged-services`, ahead of the M810 it was expected to wait for, and it landed *after* the
> fix. See §5 row 4: the tooling had already been re-derived, so the removal passed through it harmlessly.

---

> ### ⚠️ An archived repo's OWN `CLAUDE.md` is not a source of truth for its status — four of the six say they are live
>
> **Censused M257x iter-227**, all six archived repos read at their `origin/main` tips (fetched and
> advanced at iter-224). The question was iter-224's lesson turned into a check — *"a resolved anchor
> quoting a verbatim line is not the source's position; read the repo's **retraction surface**"* — and the
> census says **that method only works for two of them.**
>
> | repo | what the repo's own `CLAUDE.md` says about itself | verdict |
> |---|---|---|
> | `messenger` | **⚠️ FROZEN — this service no longer runs** | **agrees** |
> | `storage` | **⚠️ FROZEN — the service no longer runs, but the terraform module is LIVE** | **agrees** |
> | `cms` | *"Content layer for the platform: serves job simulations, skill path chapters, and the content library **via GraphQL Federation**"* — **no status statement anywhere** | **corpus-ahead** |
> | `jobsimulation` | a live-service doc; **no status statement**, though `6092c6d2` destroyed its ECS service and ECR repository | **corpus-ahead** |
> | `graphql-wundergraph` | *"**it is currently the sole subgraph**"* — a live router, with **no mention** that platform `2adcf71` deleted it. Also carries the **`2 → 1`** subgraph figure this corpus corrected to **`3 → 1`** | **corpus-ahead ×2** |
> | `roadrunner` | *"Sandboxed code execution service… **Deploy: Docker -> ECR -> ECS**… used exclusively by `jobsimulation`"* — a live service, consumed by another that is itself gone | **corpus-ahead** |
>
> **Two of six carry a freeze banner; four describe themselves as running services.** Not one of the six
> states a status this map does not already hold, and **not one contradicts this map on direction** — every
> disagreement is the repo being behind, never the corpus being wrong. That is the reassuring half.
>
> **The trap is the other half, and it is aimed at exactly the reader this corpus is written for:** an
> agent that clones `cms` or `roadrunner` to check a claim will open its `CLAUDE.md` and be told the
> service is live and deployed to ECS. The freeze wave of 2026-08-05 reached `storage` and `messenger`
> and stopped. **So "check the repo's own docs" is a method with a 2-in-6 hit rate here, and the
> platform's `repos.yml` + `infrastructure`'s `services.tf` remain the authorities** — as [§4 The
> fence](#4-the-fence) already asserts mechanically.

## 1. How to read a row

| state | means |
|---|---|
| `live-standalone` | its own process, still on the traffic path |
| `merged-into-app` | `app` owns the code and calls it unconditionally, the tables live in `public`, and the standalone is scaled to zero — **all three**, per [§6 of the protocol](../ops/platform-alignment.md#6-classification--the-map) |
| `running_but_unfederated` | the container still starts, but it owns no schema and is not a subgraph — a **husk**, not a service |
| `mid-fold` | **a half-landed fold: the config side says removed and the consumer side says live.** Neither `live-standalone` nor `merged-into-app` — and it is recorded on **both** sides, cited, or not at all. Added M257x iter-64; the gap it closes was stated in [§6 of the protocol](../ops/platform-alignment.md#6-classification--the-map) at iter-59. **No row carries it today** (re-checked M257x iter-87 at `0c91421`): its only holder, `storage`, completed its fold four iterations later and had its container deleted the day after that. The token stays in the vocabulary — the fold program is not finished, and a state you can only name *after* you need it is the state you will get wrong |
| `decommissioned` | gone from the orchestration; the repo may still exist as a rollback reference |
| `net-new` | exists in the org, is not in `repos.yml`, and the corpus has never described it |
| `external` | third-party or separately-deployed; never in the local Go clone set |
| `library` | imported as a private Go module, never a process |
| `library-unimported` | the repo exists and the module is published, but **no `go.mod` a stack builds requires it** — the live copy lives somewhere else (vendored in-tree, or shipped inside another module). Added **M257x iter-130** with assertion G, which fired on `ai` and `authn` the first time it ran. ⚠️ **This row was missing from this table for three iterations while assertion C's description already said "nine"** — the vocabulary change reached the checker (`ALLOWED_STATES`) and not the definition, so the guard was green over a document that defined eight. Caught by two independent seats at M257x iter-131 (P7) and repaired at iter-133. **`library` is not its superset and `decommissioned` is not its synonym**: `library` asserts an import that does not exist, and `decommissioned` describes an orchestration lifecycle a library was never in |

**Two traps this table exists to keep straight:**

- **`migrations: false` entails nothing on its own.** The exemplar used to be `sentinel` — `migrations: false`
  *and* alive *with its own `sentinel` schema*. ⚠️ **That instance died at `766df6c`** (no `repos.yml` entry
  at all now), and the trap survived it in a sharper form: the `sentinel` **schema** still exists and is still
  not `public`, but it is now owned by `app`, whose row reads `migrations: true`, `schema: public`
  (`repos.yml:3-6`) — one flag, one schema named, and **two** schemas actually written
  (`docker-compose.yml:25`, `SENTINEL_DB_CONNECTION` with `search_path=sentinel`). Read the `prod` and
  `fresh local stack` columns, never the flag alone.
- **Absent from `repos.yml` no longer means "never was a service".** Until `d11a403` the declared topology
  and the actual one disagreed — `repos.yml` called cms / jobsimulation / roadrunner *"legacy"* while compose
  still started all three. `d11a403` closed that gap by deleting both sides at once, so the three now look
  exactly like `skiller` and `skillpath`: no row in `repos.yml`, no compose service, **repo still on GitHub as
  the pre-merge reference**. `838d907` did the same for `storage` and `messenger` on 2026-08-05, which took the
  clone set to **four**. The clone set is therefore no longer a census of what the platform has ever run —
  which is what this file is for.

---

## 2. The services

**Completeness is measured, not asserted.** The row set is the union of *every name that has ever appeared in `repos.yml`* (`git log -p --follow -- repos.yml` → **14** names: app ·
chronos · cms · graphql-wundergraph · intelligence · jobsimulation · messenger · next-web-app · roadrunner · sentinel · skiller · skillpath · storage · studio-desk — all 14 have
rows) and *every service that has ever appeared in `docker-compose.yml`* (same command on that file → **25** names, including the pre-history the clone set never knew: `nats`,
`web-app`, `chromedp`, `simulator`, `realtime`). **Collect only the keys under `services:`, not every two-space key in the file:** a section-blind pass returns **26**, and the one
extra token is `app-network` — the **network** declared under `networks:` (`docker-compose.yml:185-186` @ `0c91421`), never a service and correctly without a row. **This passage
said 26 until M257x iter-102**, which turned its own audit instruction into a false alarm on every re-run. Re-run those two commands to audit this table; a name they return that
has no row is a gap — three of the 25 have their row filed under the repo name rather than the compose key (`backend` → `app`; `graphql` and `wundergraph` → `graphql-wundergraph`).

<!-- fence:services:begin -->

| repo | prod | fresh local stack | in `repos.yml` | evidence |
|---|---|---|---|---|
| `app` | live-standalone | live-standalone | yes | the monolith. `app/terraform/main.tf:181` `service_desired_count = 1`; **the only migrating repo** — `repos.yml:3-6` (`migrations: true`, `schema: public`); compose service is named `backend` (`docker-compose.yml:5`). ⚠️ **Both of those citations were `repos.yml` lines 14-17 and `docker-compose.yml` line 28 until M258 iter-18** — correct at `0c91421`, and stale the moment `766df6c` took `repos.yml` 28 → **13** lines and `docker-compose.yml` 190 → **164**. *(Retracted citations are written **without** the `path:line` form on purpose: a retraction quoted verbatim re-arms the dead citation as a live one, and — where the row also names the block the stale line drifted into — silently satisfies assertion F's declared-cross-block escape. Measured on this very edit, M258 iter-18.)* Owns **eight** domains in-process — the four folded before v9.0, plus storage, messenger and customerio-sync, **plus `sentinel` at v11.0** (`app/internal/sentinel/`, wired at `app/main.go:305`; see that row) — each with **its own** wiring call site: skiller `app/main.go:706` (`skiller.NewSkillerManager`), jobsimulation `:734` (`jobsimwiring.Wire`), skillpath `:764` (`skillpath.NewSessionManager`), cms `:1167` (`appcms.Wire`), storage `:537` (`internalstorage.NewManager`), messenger `:1458` (`msgadapters.Wire`), customerio-sync `:396` (`customeriosync.New`), **and sentinel `:305` (`sentinel.Open`)** — `app/internal/{cms,jobsimulation,skiller,skillpath,storage,messenger,customeriosync,sentinel}/`. ⚠️ **All seven of the original anchors were re-resolved at M258 iter-20 against `app` `c52dbc51e`** (= `origin/main`, verified at the remote): they were pinned at `2035f9a` and every one had drifted by 12–20 lines. **No fence saw it** — they are `app/main.go:NNN` citations, which assertion F grades range-only, and a 1,635-line file swallows a 20-line slip. The sibling anchor in this same table that WAS graded (`messenger`'s `:1485`) had already slid onto a closing brace. *Anchors into a file with no block structure are the class the fences cannot see; re-derive them at every ref bump, or do not pin them.* **Anchors re-resolved M257x iter-87 at `app` `2035f9a` (post-v1.369.0) — a PIN, not a moving label.** `2035f9a` *was* `origin/main` on 2026-08-05; re-checked 2026-08-06 it is **five commits behind**, and `origin/main` is now **`ad9f3c49`**. Those five touch `.claude/skills/publish/SKILL.md`, `CLAUDE.md`, `knowledge/deployment.md`, `terraform/main.tf` and `terraform/variables.tf` — **no Go source at all**, and the single `terraform/main.tf` hunk rewrites one precondition error message in place, so every anchor in this cell resolves to the same construct at either ref — `ad9f3c49` is a currency note, not a fifth re-derivation, which is why the count below still reads four. **This cell — and three others in this table — wrote the sha as `origin/main` until M257x iter-102:** the sha is a pin and still means what it meant; it is the *label* that expired, and a label that moves under a citation is how a correct anchor becomes a wrong one without anybody editing it. The six that were already cited have moved at **every one of the four refs this map has read in a week** — `5ba17044` v1.363.2, `b948604` v1.366.0, `9d00a313` v1.367.0 (iter-68), and now `2035f9a`; three re-derivations for six anchors, none of them caused by a change to the code they point at. **The older refs are named without their line numbers on purpose** — a block naming two refs is `ambiguous` to the citation resolver, which then falls back to origin/main and grades every anchor in the cell against a file the cell did not mean (M257x run-53; `storage`'s row omits them for the same reason). An earlier revision attached the jobsimulation site alone to all four, where it wires jobsimulation only; corrected M257x iter-46. **`app/internal/roadrunner/` does not exist** — the Judge0 runner was absorbed as `app/internal/jobsimulation/runner/`, constructed at `app/internal/jobsimwiring/wiring.go:123` (`jsrunner.NewRunnerManager(JUDGE0_API_KEY, JUDGE0_BASE_URL)`) |
| `cms` | merged-into-app | decommissioned | no | `cms/terraform/main.tf:39` `service_desired_count = 0`; code in `app/internal/cms/`; folded by platform `236771f` (2026-07-29, cms-in-app v8.0). **Compose service and `repos.yml` entry both deleted by `d11a403`** (merged `ef32d4c`, 2026-08-03) — `make init` no longer clones it. Repo **not** archived. **The pointer to the prod rollback path is gone from `repos.yml`:** its header named infrastructure's `services.tf` as that path until M810 (`repos.yml:9-10` @ `0dab54d`), and `838d907` rewrote the header without it — what stands now says only that the frozen repos own no schema, no compose service and no clone entry, and that *"None of them are deleted"* (`repos.yml:2-10`). **Whether that rollback declaration still stands is not something this map can see** — it never could, since infrastructure has never been in the clone set; it had a pointer, and `838d907` removed the pointer. Absence of the sentence is not evidence the declaration went with it. **M257x iter-92 — cms has since taken an M810 step, and it points the OTHER way:** `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** `.github/workflows/build-production.yml`, subject *"the cms ECR repository is decommissioned (M810)"*, body: M810 *"deletes `module \"cms_euwest1\"` … which destroys the ECS service and the production-cms ECR repository"*, the workflow being dropped because it *"would try to push an image into a registry that no longer exists."* So this repo now holds **two measured facts pointing opposite ways** — `cms/terraform/main.tf:39` still declaring the module, and a CI commit asserting the registry is already gone. **RESOLVED at M257x iter-123 — the ECS service is DESTROYED.** `infrastructure` was cloned and read (`13c248e6`, 2026-08-07): there is **no `module "cms"` declaration anywhere in it** (`grep -rn 'module "cms' terraform/ modules/` → zero). In its place `infrastructure/terraform/production/services.tf:64-70` states it in the platform's own words — *"M810: cms was removed (module block deleted here)… Deleting the module destroys its ECS service, task definition, ECR repository, IAM roles, security group, Cloud Map entry, log group, alarms and the ten `/production/cms/*` SSM parameters"* — with a `removed { … lifecycle { destroy = false } }` for the Atlas tracker at `infrastructure/terraform/production/services.tf:88-94` and `infrastructure/terraform/production/services.tf:85-86` noting the legacy **schema** is untouched (**that drop is a separate, still-pending M810 step**). So the two facts were never contradictory: **the CI commit was the correct signal, and `cms/terraform/main.tf:39` is ORPHANED DEAD CODE** — no root module instantiates that file, so its `service_desired_count = 0` describes nothing. **The blocker this cell named for four iterations — *"the destruction happens in infrastructure's `services.tf`, which we cannot read"* — was a clone-set limit, not a measurement limit, and the fix was to clone the repo.** Full derivation + the sibling rows it also settles: [`org-repos.md` § 3](org-repos.md) |
| `jobsimulation` | merged-into-app | decommissioned | no | **M810 has LANDED for this service, and this map had not noticed** — `6092c6d2` (*"remove the jobsimulation ECS service and ECR repository (M810)"*) deleted the `module "jobsimulation"` block outright, destroying the ECS service, task definition, ECR repository, task/execution IAM roles, security group, Cloud Map entry, log group and alarms (`jobsimulation/terraform/main.tf:15-22` @ `82cb66e`). It had run at `service_desired_count = 0` since app v1.360.0; **that line does not exist any more**, which is why this row cites the decommission comment and not a count — the old citation to it resolved to an unrelated comment about the atlas tracker, the exact silent-slide failure [§4](#4-the-fence)'s assertion F was built for, in the one file class F does not reach. The module deliberately survives because it still owns the **LiveKit and Chime recording buckets `backend` reads by literal name**, the `/production/jobsimulation/*` SSM parameters, and the atlas tracker for the legacy `jobsimulation` schema — dropping that schema is a separate M810 step (`:24-40`). **Do not generalise M810 from this row — and do not read `cms` as standing still either.** `cms` holds **two measured facts pointing opposite ways**: `cms/terraform/main.tf:39` still declares `service_desired_count = 0` in a module that still exists, **and** `6efa1d5` (merged `f38c0c4`, 2026-08-04) deleted that repo's build-production workflow because *"the cms ECR repository is decommissioned (M810)"*. The `cms` row directly above **has since RESOLVED it, and anything reading across from this row must take the resolution, not the hedge**: `infrastructure` was read at `13c248e6` and declares no `module "cms"` — `cms`'s ECS service is destroyed too, and the repo-side `service_desired_count = 0` is orphaned dead code. **This cell said *"report both and assert neither … the destruction lands in infrastructure, which is in no clone set"* for four iterations after the read.** Not being in the standing clone set is a fact about our habits, never about what is measurable. **This row said *"`cms` has not moved"* until M257x iter-102** — a flat assertion the row above it had already retracted, which is how one table came to hold both readings. Code in `app/internal/jobsimulation/`, wired unconditionally at `app/main.go:721` (`jobsimwiring.Wire`, @ `app` `2035f9a` — a pin; see the re-measured banner at the top of this file for where `origin/main` has gone since); tables re-created in `public`; folded by platform `236771f`. **Compose service and `repos.yml` entry both deleted by `d11a403`.** **Repo archive state — REPORT BOTH, ASSERT NEITHER.** This row asserted a GitHub archive on 2026-07-31 until M257x, and the clone refutes the flat form: `origin/main` carries **four commits dated 2026-08-04** — `6092c6d2` (the M810 terraform delete), `caf36c96` (**Merge pull request #439**, committer `GitHub`, i.e. the merge button), `1e40d184`, and `82cb66ec` (v0.254.0) — while an archived GitHub repo is **read-only**: no push, no PR merge. So either it was not archived on that date, or it was un-archived to land M810. **Which one is not something this map can see** — archive state lives in the GitHub org API, never in the git objects (`gh api repos/anthropos-work/jobsimulation --jq .archived`; census signal 6, [`platform-alignment.md`](../ops/platform-alignment.md)). A dated archive assertion is a **snapshot with an expiry**, not a derived fact |
| `roadrunner` | decommissioned | decommissioned | no | **⚠️ The prod cell read `live-standalone` until M257x iter-280, four iters after this row's own prose retracted it.** §1 defines `live-standalone` as *"its own process, still on the traffic path"* — and the evidence below, measured at iter-123, says roadrunner appears in NO terraform in `infrastructure` at all and that there is no roadrunner ECS service. **The row argued against its own state token, and the token is what a reader scans.** Corrected to `decommissioned` (*"gone from the orchestration; the repo may still exist as a rollback reference"* — the repo is not archived). **Scope of this correction, stated so it cannot be over-read:** it makes the CELL agree with the measured prose already in this row; `infrastructure` is not in the clone set and was NOT re-read at iter-280. **Compose service and `repos.yml` entry both deleted by `d11a403`, in that one commit.** Its message says the clone entry *"was already gone, so the `../roadrunner` build context could no longer resolve"* — **the message is wrong; the diff is the fact.** `git show d11a403 -- repos.yml` shows that very commit deleting `- name: roadrunner` alongside `- name: cms` and `- name: jobsimulation`, and the compose file at `d11a403^` still declares a `roadrunner:` service block (it was one of eleven there; `d11a403` left eight, and `838d907` has since taken that count to **five**). The service was legacy, not unbuildable. (An earlier revision of this row promoted that message into a conclusion — *"the service had been unbuildable, not merely legacy"*; corrected against the diff in M257x. **A commit message is testimony, not evidence** — grade a change by its diff.) Judge0 is reached directly: `JUDGE0_BASE_URL` moved onto `backend` (`docker-compose.yml:59`) for `app/internal/jobsimulation/runner/` (`app/internal/jobsimwiring/wiring.go:123` @ `app` `2035f9a` — a pin, per the `app` row; the same line at `9d00a313`; it was five lines earlier at `b948604`, which is the number `d11a403`'s own message quotes, and it is not repeated here for the reason the `app` row states). **MEASURED AT LAST (M257x iter-123, `infrastructure` `13c248e6`):** a service repo's own `service_desired_count` is **not evidence of production state** — it is an input to a module that must be *instantiated* by `infrastructure/terraform/production/services.tf` to mean anything, and **four repos declare a count that instantiates nothing.** **`roadrunner` appears in NO terraform in `infrastructure` at all** — 7 org-wide hits, all of them `judge0_*` secret names still labelled *"Roadrunner"* in two CI workflows plus one KB line. **There is no roadrunner ECS service, and `roadrunner/terraform/main.tf:19` describes nothing.** The whole line of enquiry below is therefore SETTLED, and is kept because its method — *a commit message is testimony, grade a change by its diff* — outlived its subject. [`org-repos.md` § 3](org-repos.md). **The prod contradiction WAS explained but not verified, and here is that record:** `roadrunner/terraform/main.tf:19` remains `service_desired_count = 1` — last changed at **`84a4b4f` (2025-12-15)**, the commit that first added `terraform/main.tf`, and untouched by everything up to the repo's HEAD `87d8d44` (2026-06-19). **That count is not a decision about the fold; it predates it by seven months and nobody has been back.** (An earlier revision of this row dated it to `e45eb61` (2026-05-27) — that commit is the file's most recent touch but it changed **line 11 only**, a one-line module-source URL swap whose own message says *"Module contents are identical; this is a pure source-URL swap"*. `git blame -L 19,19` names `84a4b4f`; a file-level `git log` is not line provenance.) **This row was RIGHT and the service doc was WRONG for four readings** — [`roadrunner.md`](../services/roadrunner.md) dated the line to `87d8d44` (the repo HEAD, which touches only a workflow file) and **never named `84a4b4f` anywhere**, while pointing the reader here two lines below its own error. Corrected at M257x iter-115; re-verified at that iter, unchanged. The authoritative declaration lives in **infrastructure's `services.tf`**, which `repos.yml:9-10` @ `0dab54d` said in words until `838d907` rewrote that header (see the `cms` row). **⚠️ This clause called that *"a repo this map has never read"* until M257x iter-137 — stale since iter-123, which read it, and re-read independently at iter-137.** Losing the citation did not settle the disagreement; reading the repo did. **And the read supplied the POSITIVE half this row lacked:** the `production_roadrunner_judge0_*` secrets feed `TF_VAR_judge0_{api_key,base_url}` (`infrastructure/.github/workflows/wf-terraform-deploy.yml:209-211`) into **`module "backend_euwest1"`** (`infrastructure/terraform/production/services.tf:384-385`) — production wiring Judge0 straight into `backend`. Repo **not** archived. **⚠️ THE REFERENCE THIS MAP IS FENCED AGAINST STILL SAYS THE OPPOSITE, IN PROSE** (measured M257x harden pass 72): `repos.yml`'s own header comment at `0c91421` reads *"skiller, skillpath, **roadrunner**, jobsimulation, cms, messenger, storage and customerio-sync are all served in-process"* — **eight**, roadrunner among them. That sentence is where the refuted claim came from, and `platform_alignment_guard` cannot catch it: the fence grades `repos.yml`'s **membership** (which `- name:` entries exist), not its **comments**, and on membership the two agree exactly. So an agent who opens the fenced reference to check this row is told roadrunner was folded, and the correction at M257x iter-137 — `git log --all --diff-filter=A -- internal/roadrunner` → **0 commits, ever**, in a 6,728-ref clone — reads as the corpus contradicting the platform. It is not: it is the corpus having *measured* what the platform's comment *asserts*. **Do not "fix" this row from that comment.** |
| `sentinel` | **mid-fold** | decommissioned | no | **THE 8th MERGE — and this row read `live-standalone` / `live-standalone` / `yes` until M258 iter-18, which is verbatim the departure class assertion B exists for.** It fired: *"the map claims sentinel is in repos.yml, and it is not."* Platform **`766df6c`** (*"remove sentinel service and related configurations"*), on top of `befca6d` (*"SENTINEL_MODE + SENTINEL_DB_CONNECTION for backend (app v11.0)"*) and `48de408` (*"backend is the PDP — drop SENTINEL_MODE and the sentinel dependency"*), deleted the compose service **and** the `repos.yml` entry together, so `make init` no longer clones it and `repos.yml` declares **three** repos — `app` (`repos.yml:3-6`), `next-web-app` (`:8-10`), `studio-desk` (`:11-13`) — not four. **Consumer side — folded, unconditionally:** `app/internal/sentinel/` is the Casbin PDP, ported from the standalone repo at tag **v0.24.2** (`app/internal/sentinel/doc.go:10`; `f2c46190` is `sentinel` origin/main today, so the port source and the repo head are the same commit), wired exactly once at `app/main.go:305` under the source's own *"There is no switch and no RPC path: app IS the PDP"*, `log.Fatalf` on failure, `app/main.go:314` logging `authorization: in-process PDP active`. **Config side — still a standalone in production:** `sentinel/terraform/main.tf:19` `service_desired_count = 1`, and unlike the `cms`/`messenger`/`roadrunner` class that count **is instantiated** — `module "sentinel_euwest1"` is the first of the ten declarations in `infrastructure/terraform/production/services.tf` as measured at `13c248e6` (2026-08-07; [`org-repos.md` §3](org-repos.md)). ⚠️ **That half is NOT re-measured here** — `infrastructure` is in no clone set and assertion F reports it unclonable; what IS measured in-tree is `app`'s own pin of the same transition, `app/sentinel_wiring_test.go:57` (*"TestNoRPCPathSurvives … the variable is gone too — infrastructure drops its services.tf reference in the same release"*). **Neither cell is `merged-into-app`, and the reason is §1's own three-part test:** the code is owned ✅, but the policy tables stay in their **own `sentinel` schema** (`docker-compose.yml:25`, `SENTINEL_DB_CONNECTION` with `search_path=sentinel`) ❌ and the standalone is not scaled to zero ❌. `mid-fold` is the token the vocabulary reserved for exactly this and has had **no holder since `storage`**; the platform names the remaining step itself at `docker-compose.yml:85` — *"sentinel removed at v11.0: backend answers authorization in-process and does not talk to that container at all… until M1103 decommissions it."* (⚠️ that comment's *"It stays defined above as the rollback target"* is **already false**: `766df6c` deleted the block it points at.) **The RPC edge is gone, not re-pointed — and it took the whole listener with it.** `AUTHORIZATION_ADDRESS` occurs **zero** times across `docker-compose.yml`, `common.yml` and `repos.yml` (measured at `766df6c`), and `app` deleted its Connect-RPC server outright (`app/main.go:1310`, *"NO RPC SERVER"* — the port-8081 mux that carried Users/Organizations/Skiller/JobSimulation/CMS/lab), so a `core` stack now has **no cross-process Connect-RPC edge at all**. Net-new and previously documented nowhere: cross-replica policy invalidation is **Redis Pub/Sub fan-out** on channel `sentinel:policy:invalidate` (`app/internal/sentinel/watcher.go:55`), deliberately **not** the Watermill consumer-group plumbing, which delivers to one consumer. Repo **not** archived and **not** deleted |
| `storage` | merged-into-app | decommissioned | no | **The v9.0 fold is finished on both sides, and the container is now gone.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service **and** the `repos.yml` clone entry in one commit, so `make init` no longer clones it and **the `storage-legacy` profile is gone** — removed with the service it selected. Asking compose for a retired token still exits 0 and starts only the always-on floor — **postgresql and redis**; ⚠️ **this cell read *"(postgresql, redis, sentinel)"* until M258 iter-18**, true at `0c91421` and retracted at `766df6c`, which deleted the `sentinel` service and took the floor from three to two (see the `sentinel` row above) — so the stack looks alive with the application absent: grade a documented compose command on *does it still select something*, never on *does it still parse*. This row read `mid-fold` for four M257x iterations (iter-64 → iter-68), then `merged-into-app` with a legacy rollback container for exactly one platform day. **Prod — the ECS service is DELETED, not scaled to zero.** The module is **18 lines** and declares no service at all; `storage/terraform/main.tf:9-11` says it in the platform's own words — *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS."* The module deliberately survives, because deleting the block would destroy the two buckets, the CloudFront distribution and the media DNS record along with the `prevent_destroy` guards that are read from configuration (`:13-16`). **There is no custody-transfer clause in that file and `:18` is not one** — line 18 is the module's closing comment, *"See outputs.tf — consumers should reference these by output, never by literal name."* The word *custody* occurs **0** times in the entire storage repo at `9f8cb53` (`git -C stack-demo/storage grep -n -i custody 9f8cb532` → no match). **M903 was never executed and is now explicitly superseded.** Its single occurrence in the repo is in a different file, `storage/terraform/storage.tf:22-25`: *"An earlier plan, M903, instead proposed `moved`-ing these resources into a local module in the infrastructure repo. That was never executed, and the shipped design supersedes it: the assets stay here and the module stays declared."* There are **no `moved` blocks** anywhere in the repo. The plan text itself lives in a third repo — `app/knowledge/plan/releases/09.00-support-in-app/m903-s3-custody/` @ app `2035f9a` — and self-declares `status: planned` in its `overview.md` front-matter, and *"Planned, not started — no terraform applied, no branch"* in its `progress.md` (only the read-only P0.2 state capture is ticked; Steps 1–4 are all open). `d3e6d32` (2026-08-05, *"correct the guard header, which contradicted the shipped design"*) is the commit that retired it, in its own words: *"M903 never ran."* **This corrects a paired claim.** This map and its neighbours cited line 38 of that file for `service_desired_count = 0` and paired it with messenger's. It was true up to `9f8cb53`^ and is now past EOF in an 18-line file — the line number is not re-pointable because the construct it named no longer exists. **Scaled-to-zero is `messenger`'s state, not storage's**, so the paired sentence has to be split, not re-anchored. **Config side:** `STORAGE_RPC_ADDR` occurs **0** times across `docker-compose.yml`, `common.yml` and `.env_example`; `STORAGE_S3_BUCKET` / `STORAGE_S3_PUBLIC_BUCKET` sit on `backend` (`docker-compose.yml:59`, `:60`) with the reason in-comment (`:61-69`), and `backend`'s `depends_on` block now states the disappearance directly (`:79-80` — *"storage, messenger and customerio-sync are not services any more — this one container serves all three in-process"*). ⚠️ **Those four were `docker-compose.yml` lines 82 / 83 / 73-75 / 102-103 until M258 iter-18** — correct at `0c91421`, and after `766df6c` the last pair pointed into the studio-desk block, at its published ports. **Consumer side:** `app` serves object storage in-process and unconditionally — `internalstorage.NewManager` / `NewPublicManager` at `app/main.go:524`, `:525`, consumed at `:547` (the resource manager) and `:1102` (`cmsStorage`); the bucket names are read from the environment at `:516`, `:517` and guarded against empty outside a developer machine at `:518-523`; the constants at `app/internal/storage/service.go:22`, `:24` are the bucket **env-var NAMES** (`EnvBucket = "STORAGE_S3_BUCKET"`, `EnvPublicBucket = "STORAGE_S3_PUBLIC_BUCKET"`), not the bucket names themselves. **`STORAGE_RPC_ADDR` is read by nothing** — a Go grep at `app` `2035f9a` (a pin, per the `app` row) returns **3 hits, every one of them a comment** (`app/main.go:504`, `app/internal/jobsimwiring/wiring.go:101`, `app/internal/storagens/callsites_test.go:189`), and the first says it in words: *"the standalone service takes no traffic and STORAGE_RPC_ADDR is gone."* **State the ref or the sentence flips:** at the OLDER `b948604` (v1.366.0) it is genuinely read, in `main.go` **and in all three `cmd/` tools** — `git -C stack-demo/app grep -n STORAGE_RPC_ADDR b948604 -- '*.go'` returns **15** hits, of which **7 are env lookups** — `main.go` ×3, `internal/jobsimwiring/wiring.go` ×1 and one in each of the three `cmd/` tools (the other 8 are two doc comments, two error strings, **three** `t.Setenv` calls and **one `t.Fatal` message** — the fourth test-file hit *names* the variable inside a failure string rather than setting it, so counting all four as `t.Setenv` overstates the arrange side by one). Six of the seven are spelled `os.Getenv`; the seventh uses a lowercase `getenv` helper, which is why a regex fitted to `os.Getenv` reports **six** and is not wrong to. **Line anchors for that side are deliberately omitted:** they resolve at `b948604` and nowhere else, and naming a second ref inside this cell makes every anchor in it ungradeable (the citation resolver reads one ref per block and reports `ambiguous` when a block names two — M257x run-53). Run the grep; it is the ref, not the line number, that carries the claim. That interval — `b948604` → `9d00a313` — *is* the consumer half of the fold, and it is what moved this row off `mid-fold`. The repo is **not** archived and **not** deleted: `repos.yml`'s header now says so in as many words — *"None of them are deleted"* — and tells you to clone it by hand to read the pre-merge source (`repos.yml:2-10`) |
| `messenger` | merged-into-app | decommissioned | no | **Folded in the same v9.0 program as `storage`, and its container went the same way, one day later.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service and the `repos.yml` clone entry together, and **the `messenger` profile is gone** with them — the same retired-token reading as `storage`'s row. **Prod — stopped, and here "scaled to zero" is still the right words:** `messenger/terraform/main.tf:29` `service_desired_count = 0` in a module that is otherwise intact (121 lines), the cms precedent again, with the reason in-comment (`:19-25` — `app` has taken the Redis consumer group over, and an untargeted `terraform apply` would silently undo a hand-made `aws ecs update-service`);  **⚠️ AND THAT DESCRIBES NOTHING IN PRODUCTION — corrected M257x iter-123.** `module.messenger_euwest1` is **deleted** from `infrastructure/terraform/production/services.tf` (`13c248e6`; `infrastructure/terraform/production/services.tf:622` — *"module.messenger_euwest1 is deleted above"* — with only its ECR `removed{}` surviving at `infrastructure/terraform/production/services.tf:664`). So `service_desired_count = 0` is an input to a module no root module instantiates, and **the in-comment claim that it is "the rollback path" is the repo describing an intent that production no longer holds.** The line is real, the count is real, and neither is evidence: [`org-repos.md` § 3](org-repos.md) has the rule and the three sibling repos it also settles. **⚠️ and the "rollback path" is gone too — corrected M257x iter-224.** `messenger/terraform/main.tf:27-28` does still declare the image and task definition *"the rollback path… a one-line revert plus an apply, not a re-provision"*, and this row quoted it. **The same commit that is `origin/main` retracts it elsewhere in the repo:** `459b184` (merged `e9421c6`, 2026-08-05, *"correct the ECR claim — production-messenger was destroyed, not preserved"*) rewrote `messenger/CLAUDE.md` to record that the `production-messenger` **ECR repository was DESTROYED on 2026-08-05**, hand-deleted with `production-storage`, `production-customerio-sync`, `production-skiller` and `production-wundergraph` (infrastructure #3253), leaving *"exactly six repositories, one per live service"* in eu-west-1 and the `removed { destroy = false }` block **inert**. There is no image to revert to; restoring is *"a re-provision, not a revert."* **The terraform comment was never updated — so the anchor resolves, the quote is verbatim, and the claim is still retracted.** The same two-step teardown applies to `storage` (`bef79bf`, same day, same wording). **Do not pair this with `storage`'s prod state:** that module's service block is deleted outright, so a sentence reading *"each `service_desired_count = 0`"* is half false. **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` and runs a **second subscriber server on messenger's OWN Redis consumer group** (`app/main.go:1437`, wired at `:1458`, sender at `:1460` via `msgsender.NewFromEnv`) — **re-resolved at M258 iter-18 against `app` `c52dbc51e`**; the previous anchors were taken at `2035f9a` and had drifted (`:1485` now lands on a closing brace, which is exactly the `anchor_construct_guard` finding class) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's and is a deliberate literal (`:1430-1435`). **All of it sits behind a switch, which is new since iter-68:** `MESSENGER_ENABLED` (resolved at `app/main.go:286`, read at `:1459` and `:1576`) is **OFF when unset on a developer machine and a BOOT FAILURE when unset in a deployed one** (`app/env_guards.go:92-111`); prod sets it to `"true"` in the task definition (`app/terraform/main.tf:415-416`), and compose sets it **nowhere on purpose** — pinning it to `false` there would override `.env` and make opting in impossible (`docker-compose.yml:84-92`). **The RPC edge is gone, not re-pointed.** messenger's compose block was the only place in the platform that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` or `SKILLER_RPC_ADDR`; `d11a403` had re-pointed two of them at the monolith by hand, and `838d907` deleted the block, so **all four are now set by no compose file at all** — there are **zero `*_RPC_ADDR` variables anywhere in compose**. ⚠️ **This clause used to continue *"…and the only cross-process Connect-RPC edge out of `backend` on a `core` stack is `backend → sentinel` (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, docker-compose.yml line 48)"* — RETRACTED at M258 iter-18.** That edge is **gone**: `766df6c` deleted the `sentinel` service, `AUTHORIZATION_ADDRESS` occurs **0** times across `docker-compose.yml`, `common.yml` and `repos.yml`, and `app` deleted its Connect-RPC listener with it (`app/main.go:1310`). **A `core` stack now has NO cross-process Connect-RPC edge at all** — see the `sentinel` row. What survives of the qualification, and is still the point: `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:34`; `gotenberg` is a live container in the default `core` profile, `docker-compose.yml:161`, consumed at `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:36`) — so *"no RPC edge"* still does **not** mean *"no cross-process edge."* **This row generalised that to *"the only cross-process service address left in a local stack"* until M257x iter-102** — the `*_RPC_ADDR`-is-zero half was and is true; the generalisation from *no RPC edge* to *no cross-process edge* was not, and the `gotenberg` row of this same table already graded that container `live-standalone` on a fresh local stack. `messenger/cmd/root.go:120-140` still reads all four, which is what made that re-point necessary while messenger was a process; it is not one any more. Repo **not** archived and **not** deleted (`repos.yml:2-10`) |
| `next-web-app` | external (Vercel) | live-standalone | yes | `repos.yml:8-10`; `docker-compose.yml:121` (`frontend` profile, `:146`). Points at `backend` directly since the router drop — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` resolves to `…:8082/graphql/query`, baked as a build arg (`docker-compose.yml:129`) and set again in the environment (`:138`). ⚠️ **All five citations re-resolved at M258 iter-18** — they read `repos.yml` lines 23-25 and `docker-compose.yml` lines 143 / 168 / 151 / 160, correct at `0c91421`; at `766df6c` three were out of range and two had slid into the gotenberg block |
| `studio-desk` | live-standalone | live-standalone | yes | `repos.yml:11-13`; `docker-compose.yml:90` (`studio-desk` profile, `:119`). Same re-point — `VITE_GRAPHQL_ENDPOINT` build arg at `:97`, environment at `:113`. ⚠️ **All five re-resolved at M258 iter-18** (they read `repos.yml` lines 26-28 and `docker-compose.yml` lines 112 / 141 / 119 / 135 — correct at `0c91421`; at `766df6c` two had slid into the next-web-app block, and line 119 had become studio-desk's *profiles* line, which is how a stale citation keeps landing plausibly) |
| `graphql-wundergraph` | decommissioned | decommissioned | no | **⚠️ The prod cell read `live-standalone` until M257x iter-280 — the SECOND row with this defect, which is why iter-280 enumerated the class instead of repairing the `roadrunner` site it was pointed at.** The prose below says *"In prod it is DESTROYED — corrected M257x iter-124"* and records that `module.wundergraph_euwest1` is deleted, the ECS service / task definition / target group / ALB rule destroyed, and the repo ARCHIVED on GitHub — while the state token next to it still said the service was on the traffic path. iter-124 repaired this row's PROSE and left its TOKEN; the same iter's lesson was that repairing a site leaves the class. **Same scope caveat as the `roadrunner` row:** the cell is made to agree with the measured prose already here; `infrastructure` was not re-read at iter-280. **the router, dropped from local dev mid-milestone.** Deleted from `repos.yml` **and** compose by `b56d731` + `360efd4`, merged `2adcf71` (2026-07-31); local dev now points at `backend`. **In prod it is DESTROYED — corrected M257x iter-124.** `module.wundergraph_euwest1` is deleted from `infrastructure/terraform/production/services.tf` @ `13c248e6`; `infrastructure/terraform/production/services.tf:509-517` records that the apply destroyed *"its ECS service, task definition, target group, ALB rule (priority 810), Cloud Map entry, log group, ACM cert and the `wundergraph.anthropos.work` alias"*, leaving only a `removed{}` for the ECR (`infrastructure/terraform/production/services.tf:521`), hand-deleted **2026-08-05** — *"so production-wundergraph is gone and this block is now inert."* **This cell read *"in prod it is still declared"* until iter-124, citing that archived repo's own `service_desired_count = 1`**; that declaration is **orphaned dead code**, and iter-123 measured the rule it violates — *a service repo's own `service_desired_count` is not evidence of production state* ([`org-repos.md` § 3](org-repos.md)). **iter-123 corrected the sibling rows (`cms`, `roadrunner`, `messenger`) and this one was missed**, which is why iter-124 enumerated the class rather than repairing a site. The **repo is ARCHIVED on GitHub 2026-07-30**. Supergraph is **one** subgraph: `supergraph-config-prod.yaml` lists `backend` alone, `schemas/` holds `backend.graphqls` alone, `subgraphs.conf` = `BACKEND=v1.360.0` (folded by `915da06`, 2026-07-29) |
| `skiller` | merged-into-app | decommissioned | no | removed from compose + `repos.yml` by platform `21429b7` (2026-07-07); code in `app/internal/skiller/`; taxonomy data in `public`. **Repo ARCHIVED 2026-07-01** |
| `skillpath` | merged-into-app | decommissioned | no | decommissioned by platform `a4db680` (2026-07-21, M507); code in `app/internal/skillpath/`; session state in `public.skill_path_sessions`. **Repo ARCHIVED 2026-07-31** |
| `chronos` | decommissioned | decommissioned | no | removed from orchestration by platform `045857c` (2026-04-17). **Repo is NOT archived on GitHub** (last push 2026-04-23) — the corpus called it archived; the org disagrees |
| `intelligence` | decommissioned | decommissioned | no | removed from orchestration by platform `fdfa189` (2026-04-17). **Repo ARCHIVED 2026-04-02** |
| `customerio-sync` | merged-into-app | decommissioned | no | **A state transition this map had never recorded, and one no membership assertion could have caught** — it was never in `repos.yml`, so directions A and B are both blind to it. It was `live-standalone` on both sides until `838d907` (merged `0c91421`, 2026-08-05), which deleted the compose service — it had been built straight from a git URL rather than cloned — and with it **the `customerio-sync` profile is gone**. The commit states a hazard — *"was still in the `all` profile, so `make up-all` started a second Brevo contact pusher alongside backend's own."* — and **the second half of that sentence is false; it is quoted here as the platform's wording, not endorsed** (corrected M257x iter-102). The `all`-profile half is true (`profiles: [customerio-sync, all]`, `0dab54d:docker-compose.yml:154`); the *"second pusher"* half is not, because `backend`'s own in-process pusher is gated behind `CUSTOMERIO_SYNC_ENABLED`, unset and therefore **off** on a developer machine — so `make up-all` started exactly **one**. **This is the corpus inheriting a false claim by quoting a commit message as authoritative**, which is worth naming as a class: a platform commit message is evidence of *intent*, never a measurement. **Consumer side:** the code is `app/internal/customeriosync/`, constructed at `app/main.go:395` (`customeriosync.New`) behind `CUSTOMERIO_SYNC_ENABLED` (resolved at `:286`, read at `:394`) — same switch semantics as messenger's: off when unset on a developer machine, a boot failure when unset in a deployed one (`app/env_guards.go:92-111`), and set to `"true"` in prod's task definition (`app/terraform/main.tf:419-420`). compose sets it nowhere, deliberately, and says why in-comment (`docker-compose.yml:61-69`); `backend`'s `depends_on` block states the disappearance (`:79-80`). ⚠️ **Both re-resolved at M258 iter-18** (they read `docker-compose.yml` lines 84-92 / 102-103, correct at `0c91421`; `766df6c` put the second inside the studio-desk block). **The prod half is asserted from `app`'s side only:** the standalone's own terraform lives in a repo that has never been in the clone set and that this map has therefore never read — the same gap the `roadrunner` row carries, recorded rather than papered over |
| `db-backup` | live-standalone | — | no | production-only; no compose service and no `repos.yml` entry at `0c91421`. **Deployed but not triggered** — `infrastructure/terraform/production/services.tf:571` pins `ref=v0.3.3` = `6e1fb15b` = HEAD, and at that commit the EventBridge rule + target are commented out (`db-backup/terraform/main.tf:10-27`, since `7dd1b80` 2025-05-29). The task definition, ECR repo, IAM roles and S3 bucket all still exist. **Bash, not Go; S3 + Hetzner, never Azure** — [`db-backup.md`](../services/db-backup.md) |
| `anthropos-studio-room` | merged-into-app | merged-into-app | no | the Python generation pipeline is pulled into the `app` image by CI and orchestrated from `app/internal/cms/studio/`, which spawns it as a subprocess. Not a deployment, not in `repos.yml`. **The repo name is `anthropos-studio-room`, not `studio-room`** |
| `ant-academy` | external (Vercel) | external | no | deliberately absent from `repos.yml` — run natively, never containerised |
| `gotenberg` | external | live-standalone | no | third-party image, `docker-compose.yml:148-149` (`gotenberg/gotenberg:8`), default `core` profile (`:161` — renamed from `graphql` by `0dab54d`, since the WunderGraph router the name described is gone). ⚠️ **Re-resolved at M258 iter-18 from `docker-compose.yml` lines 170-171 / 183** — both past the end of a 164-line file at `766df6c`. **Since that commit this is the ONLY non-`app` service block in `docker-compose.yml` that a `core` stack starts**, so `backend → gotenberg` (plain HTTP, `:34`) is the last cross-process application edge in the default profile |
| `colony` | library | library | no | private Go module (framework + `colony/authn`); pulled at Docker build via `GH_PAT`/`GOPRIVATE`, never cloned by `make init` |
| `proto` | library | library | no | private Go module — RPC contracts + domain types |
| `ai` | **merged-into-app** | **merged-into-app** | no | ⚠️ **CORRECTED TWICE. iter-129 fixed the `fresh local stack` cell and left `prod` reading `library` — a repair that reached one of the row's two cells, inside the row it cited as the rule-54 exemplar. The prod cell was corrected at iter-130 by the new assertion G, which fired on it.** Prod and local are the *same* answer here and always were: production runs the `app` image built from the **same `app/go.mod`**, so if no local build requires the module, no prod build does either. Original finding: this row said `library` / `library`, and it is the ONLY library row the iter-102 fold never reached. `app` **folded the library in-tree as `internal/ai` at `1e457fa70`** (2026-08-04, *"refactor(ai): fold the ai library into app as internal/ai"*) and **dropped the module requirement**: at `ad9f3c498`, `app/go.mod`'s `anthropos-work` requires are `analytics-go`, `colony`, `proto`, `storage`, `taxonomy` — **no `ai`** — and `go.sum` carries **0** hits for it. By this map's own § 1 vocabulary (*a `library` is imported as a private Go module*), **no repo a stack builds imports it**: `sentinel` `f2c461903`, `storage` `9f8cb5322`, `messenger` `e9421c68f` and `roadrunner` `87d8d4438` none require it; only the frozen `cms` `f38c0c4a4` and `jobsimulation` `82cb66ecc` still do, and nothing builds either. `app/internal/ai/module_import_guard_test.go:18-38` is a **one-way door** against re-acquiring it. The repo survives because `rosetta-extensions/stack-seeding` pins `v1.40.1` for its own generator. **The corpus already knew** — [`shared_libraries.md`](shared_libraries.md) recorded the fold at iter-102 and the neighbouring `authn` row here carries the right caveat; this row was the miss, found by the iter-129 complement read |
| `authn` | **library-unimported** | **library-unimported** | no | ⚠️ **CORRECTED M257x iter-130 — this row said `library` / `library`, and assertion G fired on it the first time it ran.** The repo exists and the module is published, but **nothing imports it**: at `ad9f3c498` / `f2c461903`, neither `app/go.mod` nor `sentinel/go.mod` requires `github.com/anthropos-work/authn`, and both `go.sum`s carry **0** hits (control: `colony` = 2 hits in each, so the search is not vacuous). `rosetta-extensions` requires it in **no** `go.mod` either. The live copy ships **inside** colony as `colony/authn`. This is the *same class* as the `ai` row — a library row whose state outlived its import — and it is why iter-130 added the `library-unimported` token: §1 defines `library` as *"imported as a private Go module"*, which this is not, and `decommissioned` describes an orchestration lifecycle a library was never in. **The corpus had this right in prose and wrong in the fenced cell** — [`shared_libraries.md`](shared_libraries.md) and [`service_taxonomy.md`](service_taxonomy.md) both say in prose that no service's `go.mod` requires it — which is the whole reason the cell needed a fence rather than another sweep |
| `taxonomy` | library | library | no | private Go module — the `NodeID` type only. **Not** the taxonomy dataset, which lives in `app`'s `public` schema: **≥42,790 skills / ≥22,470 job roles** (the *public* subset, `organization_id IS NULL`, measured 2026-06-29; totals including org-private content are unmeasured). **This row quoted a 60K figure until M257x** — ["60K / 18K" is not a measurement](shared_libraries.md#taxonomy-figures): the roles figure is **REFUTED** (below the 22,470 public floor; 18,919 is the `job_role_embeddings` row count, a different table) and the skills figure **UNVERIFIED** (42,790 is a floor, not a total). The index of truth must not restate a figure the corpus elsewhere fences |
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
count), of which **4** are in `repos.yml` and **46** are named by no corpus document at all. The 4 are
**enumerated from the file, not carried forward as a sum**: `app` · `next-web-app` · `studio-desk` — **3**,
in a 13-line `repos.yml` @ `766df6c` (it was 4 in 28 lines @ `0c91421`, before `sentinel` left). **It read 9 until `d11a403`** removed the
`cms` / `jobsimulation` / `roadrunner` entries, **and 6 until `838d907`** removed `storage` and `messenger`;
the 46 does not move with either, because all five *are* named by corpus documents. The table below
is the subset that is **unarchived and pushed since 2025** — the ones that could plausibly matter and that
nothing in the corpus has ever looked for.

<!-- fence:census:begin -->

| repo | last push | why it matters |
|---|---|---|
| `kb-ant-product` | 2026-08-01 | knowledge-base repo; the org's most recently touched repo |
| `ant-observability` | 2026-08-05 | observability stack. **The corpus documented no observability tier until iter-123** — now [`observability.md`](../ops/observability.md). Holds the platform's live outside-in `product-monitoring/` **and a production read path** (asynq → prod ElastiCache) no safety doc enumerates |
| `sim-qa` | 2026-07-31 | simulation QA tooling. **iter-01's *"writes to prod and is unmarked as such"* is CORRECTED at iter-123: write-capable yes (7 mutations, a real user's JWT from an `sk_live_` key), but NOT standing (`ls .github` → absent: no workflow, no cron) and NOT unmarked — sessions are `is_test=true` by DEFAULT (`src/flow/scenario.ts:156`, `:245`), which `app` honours (`jobsimulations.graphqls:750`, `manager.go:445`, column `is_test`). The stale source of the wrong belief was sim-qa's own README (`:36-42`).** [`org-repos.md` § 9](org-repos.md) |
| `kb-certifications-iso27001` | 2026-07-07 | compliance KB; `security_compliance.md` cites no ISO-27001 programme |
| `livekit-agent-chain` | 2026-08-03 | one of **five** LiveKit agent repos, and **not a regional variant** — a different architecture (STT→LLM→TTS: Gemini 2.5 Flash + ElevenLabs), selected **by CMS content** via `voiceEngine == livekitchain`, not by a flag. **The prior cell said `ai_architecture.md` documents *none* of the agents — false; it names all five. Corrected iter-123.** [`org-repos.md` § 8](org-repos.md) |
| `livekit-agent` | 2026-05-20 | (same family) |
| `livekit-agent-azure-us` | 2026-05-20 | (same family) |
| `livekit-agent-azure-eu` | 2026-04-22 | (same family) |
| `livekit-agent-azure-eu-fr` | 2026-06-03 | (same family) — iter-01 measured `azure-eu` + `azure-eu-fr` as dispatching nothing, and **iter-123 found why: they were ABSORBED, not abandoned.** `livekit-agent` now serves four endpoints chosen by dispatch metadata; these three hard-wire one and never read the key |
| `github-runner-config` | 2026-06-26 | CI runner configuration — a 51-line non-idempotent bootstrap that registers no runner. Its value is the fact it records: **CI runs on self-hosted EU runners reaching AWS over Tailscale**. **Not a runbook.** [`org-repos.md`](org-repos.md) |
| `kb-migration-plan` | 2026-06-18 | a migration-planning KB — directly adjacent to this document's subject |
| `simulation-form` | 2026-06-14 | unknown surface |
| `customer-orbyta` | 2026-05-19 | a per-customer repo; the corpus describes no per-customer repo pattern |
| `kb-domain-singularity` | 2026-05-12 | knowledge-base repo |
| `bench-analysis-transcripts` | 2026-03-26 | benchmark/transcript analysis |
| `transcoder` | 2025-11-11 | media transcoding — adjacent to `media-substrate-spec.md`'s Bunny.net path |
| `realtime-python` | 2025-09-12 | realtime Python service |
| `studio-tools` | 2025-02-18 | studio tooling |
| `analytics-go` | 2025-02-12 | **NOT a service — a Go library, and a DIRECT compile-time dependency of the backend monolith** (`app/go.mod:14`, `v0.3.1`). Two files; carries **Stripe subscription-lifecycle events → Brevo** (`app/internal/payments/handler.go:302-316`). Dormant ~18 months and load-bearing: **do not delete.** [`shared_libraries.md`](shared_libraries.md) |
| `infrastructure` | 2026-08-07 | **the Terraform monorepo — the platform's authoritative deployment ledger, and it was never in a clone set.** Reading it at iter-123 settled the `cms` row above. Mutually pinned with the service repos it sources at tags (`sentinel`, `directus`, `storage`, `next-web-app`, `app`, `jobsimulation`, `studio-desk`, `db-backup`, `metabase`) — **a set that is not `repos.yml`'s and must not be read as it**. [`org-repos.md` § 3](org-repos.md) |
| `anthropos-knowledge-base` | 2026-08-06 | **a SECOND org corpus covering this document's subject** — and it contradicts this one: 60K skills / 18K roles asserted in **14 places with no source**, against this corpus's measured 42,790 / 22,470. Nothing reconciled the two until iter-123. [`org-repos.md` § 11](org-repos.md) |
| `hyper-studio` | 2026-08-06 | the org's most actively developed repo (827 commits since 2026-06-17) — the successor to **studio-room**, **PRE-INTEGRATION** (zero platform coupling, no deployment). `secrets-spec.md:370` already borrows its `.env.example`. [`org-repos.md` § 10](org-repos.md) |
| `AI-Labs` | 2026-08-04 | the `labs-api` **Go** control plane (stdlib-only, Firecracker microVMs, `anthroposlabs.com`). **The corpus DOES document it** — [`ai-labs.md`](../services/ai-labs.md); what was missing is *where it runs* (Ansible + systemd on a single tailnet VM, no Terraform in-repo). **96 scenario templates, not the 15 its own README claims** |
| `directus` | 2026-06-05 | **not a fork** — the extension-builder sidecar for the stock CMS image, holding **four in-house extensions no dev or demo stack installs**. An authoring-time gap, not a rendering one. [`org-repos.md` § 4](org-repos.md) |
| `metabase` | 2026-05-27 | Terraform-only; a **live** BI console (`desired_count = 1`) on its own hostname, reading the **production database outside the Sentinel authorization layer**. [`org-repos.md` § 6](org-repos.md) |
| `judge0` | 2026-02-19 | a **vendored copy** of Judge0 CE v1.13.1 + the org's patches, and **the only record of how that box is built** — the one production box that is not infrastructure-as-code. [`org-repos.md` § 5](org-repos.md) |
| `watermill-redisstream` | 2024-03-21 | the org's **only fork**, and **inert** — `app/go.mod:12` and `colony/go.mod:9` both require **upstream** `ThreeDotsLabs/watermill-redisstream v1.4.5` with **no `replace`**. Recorded so the next reader does not re-spend the hour proving it |

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
[§8 of the protocol](../ops/platform-alignment.md#8-fence--so-it-cannot-silently-recur). It reads this file
and the platform's own `repos.yml` and asserts **both directions**:

| # | assertion | the miss it catches |
|---|---|---|
| A | every repo in `repos.yml` has a row here with `in repos.yml = yes` | a service **enters** the clone set and nobody writes it down |
| B | every row here claiming `in repos.yml = yes` is really in `repos.yml` | a service **leaves** (skiller, skillpath, the router) and the map keeps asserting it |
| C | every state is one of the **nine** in §1 (markdown emphasis stripped before the check) | a row invents a state, so the vocabulary stops meaning anything |
| D | every row cites evidence | a claim with no sha and no `file:line` |
| E | no §3 census row is in `repos.yml` | the census silently overlaps the clone set it is defined as the complement of |
| G | every **library row**'s state agrees with the **module graph** — the `go.mod` of every repo `repos.yml` declares — in both directions | a library is folded in-tree and its row keeps saying `library`; or a module enters the build with no row at all |
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
| `docker-compose.yml` citations — **25** of them | **fenced** (F) |
| every other `file:line` — terraform, `repos.yml`, `common.yml`, `app/…` (**70** citations, + **1** outside any service block) | **resolution + range only.** No derivable notion of *whose* line it is, so F does not pretend |
| **library-row states**, both cells | **derived + fenced** (G, M257x iter-130) — against `app/go.mod` + `sentinel/go.mod` read at their refs |
| the prod column, PR/rollback narrative, §5's ordering | **prose-under-review.** Neither derived nor fenced. Re-check at each platform ref bump |

#### Which ROWS each assertion reaches — and the classes still unfenced

**Added M257x iter-130, and the reason is this table's own history.** The `ai` row sat wrong for four
days *inside this fenced file*, and the fence was GREEN the whole time — correctly, because `ai` is a
**module, not a clone**, so it has no `repos.yml` row for A/B to disagree with. A fenced artifact with a
silently unfenced class in it is worse than an unfenced one: the fence's green is read as a warrant over
the whole file. So the reach is now declared per row class, not per assertion.

| row class | membership (`in repos.yml`) | **state cells** |
|---|---|---|
| repos `repos.yml` declares (`app`, `next-web-app`, `studio-desk` — `sentinel` left at `766df6c`) | **fenced** A/B | **prose-under-review** — `live-standalone` is asserted, not derived |
| library rows (`colony`, `proto`, `taxonomy`, `authn`; `ai` until it left the class) | **fenced** B | **fenced** G — against the module graph |
| merged / decommissioned services (`cms`, `jobsimulation`, `skiller`, `chronos`, …) | **fenced** B | **prose-under-review** — the sha in the evidence cell is the only check, and assertion D (`stack-core/platform_alignment_guard.py`) only asks that *some* evidence exists, never that it supports the state |
| external + third-party (`postgresql`, `redis`, `directus`, `gotenberg`, `ant-academy`) | **fenced** B | **prose-under-review** — no assertion in `stack-core/platform_alignment_guard.py` reads a third-party state |
| §3 census rows | **fenced** E (must not overlap the clone set) | not a state column |

**So: every row's MEMBERSHIP is fenced; only the library rows' STATE is** — derived from the assertion
set in `stack-core/platform_alignment_guard.py` (A/B membership, G the module graph), not asserted here.
The rest of the state column is prose-under-review and must be re-derived at each platform ref bump —
that is not a promise this file keeps by itself, and pretending otherwise is how the `ai` row happened.
When the next class becomes mechanically decidable, fence it and move its line up.

F's reach is printed on **every** run, GREEN or RED — a fence whose coverage shrinks in silence is the
failure this milestone has now found four times. A run that subject-checks nothing **refuses (exit 2)**
rather than reporting the map clean.

**Take those three numbers from the run, never from this table.** They are a reading, and the reading
moves whenever the map gains a citation. M257x iter-98 found the row above saying 23/66 while the guard
said 22/69 — the coverage figure had drifted from the coverage. It moved again at M257x iter-102 (22/69 →
**25/70**), because repairing the messenger row's cross-process claim added four citations. The run prints
it in one line:

```bash
# from the rosetta repo root. Swap .agentspace/ for stack-<role>/ to use a stack's pinned clone.
python3 .agentspace/rosetta-extensions/stack-core/platform_alignment_guard.py \
    corpus/architecture/platform-migration-status.md stack-demo/platform/repos.yml
# assertion F resolved 96 citation(s) — 25 subject-checked, 70 range-only,
# 1 outside any service block; 0 unresolvable; 0 read from the WORKTREE (no ref resolved)
```

**Both paths are relative to the same root on purpose** — the recipe this replaced named the guard
relative to a stack workspace and its two arguments relative to the rosetta root, so pasting it
anywhere resolved at most one of the three.

```bash
# from a rosetta checkout, against any stack's platform clone
PLATFORM_REPOS_YML=stack-demo/platform/repos.yml \
  python3 .agentspace/rosetta-extensions/stack-core/platform_alignment_guard.py
```

Exit 0 = aligned. Exit 1 = the drift is named, by repo, in the direction it drifted.

**Direction B is the one that has actually fired in anger, and every occurrence of this class has been a
*departure*, never an arrival** — skiller, skillpath and the router left while the corpus went on describing
them; `d11a403`'s three and `838d907`'s two (`storage`, `messenger`) were named by this fence within a day of
landing, which is the whole difference the file is here to make.

---

## 5. What this map says about the program

The consolidation is a **program with a published order**, not three accidents:
v2.0 skiller → v5.0 skillpath → v7.0 jobsimulation → v8.0 cms → **v9.0 `storage` + `messenger`, whose compose
half landed at `0dab54d` (2026-08-03) and whose containers were then deleted outright at `838d907` (merged
`0c91421`, 2026-08-05) — together with `customerio-sync`, which the published order never named.** All three
are now served in-process by `backend`, and none of them is startable locally at all.
**Every row this section flagged as "already known to be wrong on a schedule" has since moved, on that
schedule** — which is the argument for keeping a section like this one rather than only a table.

The rows to watch, in order:

1. ~~**`storage` and `messenger`** — the named next fold.~~ — **this happened, on 2026-08-05, and the fence
   caught it the same day.** `838d907` removed both clone entries and all three containers. **The signal this
   row used to give was already dead when it was written:** it said *"when `repos.yml` flips either to
   `migrations: false`, the fold has landed"* — but both had read `migrations: false` since long before the
   fold was announced (`repos.yml:18-23` @ `ef32d4c`), exactly the Trap-A error §1 warns about. The signal
   that actually fired was **departure** — the row leaves `repos.yml` and the compose service is deleted —
   which is what direction B in [§4](#4-the-fence) watches, and it has now fired twice in three days
   (`d11a403`, then `838d907`). This row also predicted the right exposure for the wrong reason: `messenger`
   *was* the more exposed of the two because it was the last process calling cms and jobsimulation over RPC,
   and `d11a403` re-pointed both edges by hand — but the resolution was not a third re-point. **The process
   was deleted, and with it the only compose block that set any `*_RPC_ADDR` variable.** A dependency you keep
   re-pointing is one measurement away from being a dependency you can delete.
2. **The teardown phase has started, and it is uneven.** The program's second half — destroying what the
   folds left scaled to zero — is booked as **M810**, and this map has been carrying it as future work.
   **It has already landed for `jobsimulation`** (`6092c6d2` deleted the ECS service *and* the ECR repository)
   while `cms` holds **two measured facts pointing opposite ways** — its module still declares
   `service_desired_count = 0` (`cms/terraform/main.tf:39`) *and* `6efa1d5` deleted its build-production
   workflow because *"the cms ECR repository is decommissioned (M810)"* — so the `cms` row **reports both and
   asserts neither**, and `storage`'s service block is gone by a different route entirely (`838d907`'s v9.0
   sibling, not M810). **This clause read *"`cms` sits untouched"* until M257x iter-102**, which is the same
   flat assertion the `cms` row had already retracted; the destruction lands in **infrastructure**, a repo no
   clone set has, so *unmeasurable* is the state, not *unmoved*. Three folded services, three different prod
   dispositions, none of them derivable from the fold's own version number. **The map found this only because
   a citation had gone stale, not because anything watched for it:** prod terraform lives in repos that are not
   in `repos.yml`, so directions A and B cannot see any of it, and assertion F does not reach `.tf` files. The
   prod column is [§4](#4-the-fence)'s *prose-under-review* row, and this is what that costs.
3. **`roadrunner`** — still the row where a repo's own terraform and the platform's declaration disagree, and
   the disagreement is still one this map cannot settle without reading **infrastructure**. It got *harder* at
   `838d907`, which rewrote `repos.yml`'s header and removed the sentence naming infrastructure's `services.tf`
   as the rollback path — the map's only pointer at the thing that could settle it. **`customerio-sync` now
   sits in the same blind spot** from the other end: its prod state is asserted from `app`'s side only,
   because its repo has never been in the clone set. Two rows, one missing input.
4. ~~**`cms` / `jobsimulation` local husks**~~ — **this happened, on 2026-08-03, ahead of M810 — and the
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

5. **The merge itself keeps dropping configuration the merged code still reads.** `d11a403`'s own message
   records that deleting the three containers *"silently dropped env that `app` still reads in-process"* —
   `JUDGE0_BASE_URL`, `DIRECTUS_PUBLIC_BASE_ADDR`, `REDIS_WORKER_INDEX`, the LiveKit and Chime blocks and the
   `~/.aws/credentials` mount, all restored onto `backend`. Merging a service moves its **code**; its
   **environment** has to be moved separately, and nothing checks that it was.
   **`838d907` is the first commit in this program to answer that in its own message** — it carries a *Kept*
   paragraph enumerating what stays on `backend` and why (*"Nothing in that block addressed a deleted
   container"*), and it states the one case where the right move was to set **nothing**: `MESSENGER_ENABLED`
   and `CUSTOMERIO_SYNC_ENABLED` are deliberately absent from compose, because pinning either to `false`
   would override `.env` and make opting in impossible. That is still a human writing a paragraph, not a
   check — but it is the first time the class has been named on the way past rather than after.

---

## See also

- [`corpus/ops/platform-alignment.md`](../ops/platform-alignment.md) — the procedure that produces and
  maintains this map: detection signals, the re-point steps, the three fence layers.
- [`corpus/architecture/service_taxonomy.md`](service_taxonomy.md) — the three-tier categorisation.
- [`corpus/services/README.md`](../services/README.md) — the per-service docs this map is the index of truth for.
