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

> **Re-measured 2026-08-05 against platform origin HEAD [`0c91421`](https://github.com/anthropos-work/platform)**
> (local clone and origin level; `app` **pinned** at `2035f9a`, post-v1.369.0 — a sha, not a moving label; see the `app` row). The prior readings were 2026-08-04 at
> `0dab54d`, and 2026-08-03 at `ef32d4c` ("Merge pull request #24 … chore/prune-merged-services"); first
> measured 2026-08-01 at `2adcf71`. Re-run [§4 of the protocol](../ops/platform-alignment.md#4-detection--six-signals-cheapest-first)
> before trusting any row older than a release.
>
> **`0dab54d → 0c91421` moved three more rows, and the fence named two of them the same day.** `838d907`
> (merged `0c91421`, 2026-08-05, PR #26 *"drop the storage, messenger and customerio-sync containers"*)
> **deletes those three compose services outright** — not parks them behind a legacy profile — and removes
> `storage` + `messenger` from `repos.yml`, which is direction B again: two departures, on a tree the map had
> read one day earlier. `repos.yml` is now **4** entries (`app` · `sentinel` · `next-web-app` · `studio-desk`)
> and `docker-compose.yml` is **186** lines, down from **271** at `0dab54d` (`838d907` is +11/−96 on that
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

**Two traps this table exists to keep straight:**

- **`migrations: false` entails nothing on its own.** `sentinel` is `migrations: false` *and* alive *with its
  own `sentinel` schema* (`docker-compose.yml:18`, `search_path=sentinel`). Read the `prod` and
  `fresh local stack` columns, never the flag alone. Live at `0c91421` (`repos.yml:18-20`).
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
| `app` | live-standalone | live-standalone | yes | the monolith. `app/terraform/main.tf:181` `service_desired_count = 1`; **the only migrating repo** — `repos.yml:14-17` (`migrations: true`, `schema: public`); compose service is named `backend` (`docker-compose.yml:28`). Owns **seven** domains in-process — the four folded before v9.0, plus storage, messenger and customerio-sync — each with **its own** wiring call site: skiller `app/main.go:690` (`skiller.NewSkillerManager`), jobsimulation `:721` (`jobsimwiring.Wire`), skillpath `:751` (`skillpath.NewSessionManager`), cms `:1153` (`appcms.Wire`), storage `:524` (`internalstorage.NewManager`), messenger `:1471` (`msgadapters.Wire`), customerio-sync `:395` (`customeriosync.New`) — `app/internal/{cms,jobsimulation,skiller,skillpath,storage,messenger,customeriosync}/`. **Anchors re-resolved M257x iter-87 at `app` `2035f9a` (post-v1.369.0) — a PIN, not a moving label.** `2035f9a` *was* `origin/main` on 2026-08-05; re-checked 2026-08-06 it is **five commits behind**, and `origin/main` is now **`ad9f3c49`**. Those five touch `.claude/skills/publish/SKILL.md`, `CLAUDE.md`, `knowledge/deployment.md`, `terraform/main.tf` and `terraform/variables.tf` — **no Go source at all**, and the single `terraform/main.tf` hunk rewrites one precondition error message in place, so every anchor in this cell resolves to the same construct at either ref — `ad9f3c49` is a currency note, not a fifth re-derivation, which is why the count below still reads four. **This cell — and three others in this table — wrote the sha as `origin/main` until M257x iter-102:** the sha is a pin and still means what it meant; it is the *label* that expired, and a label that moves under a citation is how a correct anchor becomes a wrong one without anybody editing it. The six that were already cited have moved at **every one of the four refs this map has read in a week** — `5ba17044` v1.363.2, `b948604` v1.366.0, `9d00a313` v1.367.0 (iter-68), and now `2035f9a`; three re-derivations for six anchors, none of them caused by a change to the code they point at. **The older refs are named without their line numbers on purpose** — a block naming two refs is `ambiguous` to the citation resolver, which then falls back to origin/main and grades every anchor in the cell against a file the cell did not mean (M257x run-53; `storage`'s row omits them for the same reason). An earlier revision attached the jobsimulation site alone to all four, where it wires jobsimulation only; corrected M257x iter-46. **`app/internal/roadrunner/` does not exist** — the Judge0 runner was absorbed as `app/internal/jobsimulation/runner/`, constructed at `app/internal/jobsimwiring/wiring.go:123` (`jsrunner.NewRunnerManager(JUDGE0_API_KEY, JUDGE0_BASE_URL)`) |
| `cms` | merged-into-app | decommissioned | no | `cms/terraform/main.tf:39` `service_desired_count = 0`; code in `app/internal/cms/`; folded by platform `236771f` (2026-07-29, cms-in-app v8.0). **Compose service and `repos.yml` entry both deleted by `d11a403`** (merged `ef32d4c`, 2026-08-03) — `make init` no longer clones it. Repo **not** archived. **The pointer to the prod rollback path is gone from `repos.yml`:** its header named infrastructure's `services.tf` as that path until M810 (`repos.yml:9-10` @ `0dab54d`), and `838d907` rewrote the header without it — what stands now says only that the frozen repos own no schema, no compose service and no clone entry, and that *"None of them are deleted"* (`repos.yml:2-10`). **Whether that rollback declaration still stands is not something this map can see** — it never could, since infrastructure has never been in the clone set; it had a pointer, and `838d907` removed the pointer. Absence of the sentence is not evidence the declaration went with it. **M257x iter-92 — cms has since taken an M810 step, and it points the OTHER way:** `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** `.github/workflows/build-production.yml`, subject *"the cms ECR repository is decommissioned (M810)"*, body: M810 *"deletes `module \"cms_euwest1\"` … which destroys the ECS service and the production-cms ECR repository"*, the workflow being dropped because it *"would try to push an image into a registry that no longer exists."* So this repo now holds **two measured facts pointing opposite ways** — `cms/terraform/main.tf:39` still declaring the module, and a CI commit asserting the registry is already gone. **Still UNMEASURABLE here**, and now unmeasurable with contrary evidence on both sides rather than one: report both, assert neither |
| `jobsimulation` | merged-into-app | decommissioned | no | **M810 has LANDED for this service, and this map had not noticed** — `6092c6d2` (*"remove the jobsimulation ECS service and ECR repository (M810)"*) deleted the `module "jobsimulation"` block outright, destroying the ECS service, task definition, ECR repository, task/execution IAM roles, security group, Cloud Map entry, log group and alarms (`jobsimulation/terraform/main.tf:15-22` @ `82cb66e`). It had run at `service_desired_count = 0` since app v1.360.0; **that line does not exist any more**, which is why this row cites the decommission comment and not a count — the old citation to it resolved to an unrelated comment about the atlas tracker, the exact silent-slide failure [§4](#4-the-fence)'s assertion F was built for, in the one file class F does not reach. The module deliberately survives because it still owns the **LiveKit and Chime recording buckets `backend` reads by literal name**, the `/production/jobsimulation/*` SSM parameters, and the atlas tracker for the legacy `jobsimulation` schema — dropping that schema is a separate M810 step (`:24-40`). **Do not generalise M810 from this row — and do not read `cms` as standing still either.** `cms` holds **two measured facts pointing opposite ways**: `cms/terraform/main.tf:39` still declares `service_desired_count = 0` in a module that still exists, **and** `6efa1d5` (merged `f38c0c4`, 2026-08-04) deleted that repo's build-production workflow because *"the cms ECR repository is decommissioned (M810)"*. The `cms` row directly above **reports both and asserts neither**, and so must anything reading across from this one; the destruction itself lands in **infrastructure**, which is in no clone set. **This row said *"`cms` has not moved"* until M257x iter-102** — a flat assertion the row above it had already retracted, which is how one table came to hold both readings. Code in `app/internal/jobsimulation/`, wired unconditionally at `app/main.go:721` (`jobsimwiring.Wire`, @ `app` `2035f9a` — a pin; see the re-measured banner at the top of this file for where `origin/main` has gone since); tables re-created in `public`; folded by platform `236771f`. **Compose service and `repos.yml` entry both deleted by `d11a403`.** **Repo archive state — REPORT BOTH, ASSERT NEITHER.** This row asserted a GitHub archive on 2026-07-31 until M257x, and the clone refutes the flat form: `origin/main` carries **four commits dated 2026-08-04** — `6092c6d2` (the M810 terraform delete), `caf36c96` (**Merge pull request #439**, committer `GitHub`, i.e. the merge button), `1e40d184`, and `82cb66ec` (v0.254.0) — while an archived GitHub repo is **read-only**: no push, no PR merge. So either it was not archived on that date, or it was un-archived to land M810. **Which one is not something this map can see** — archive state lives in the GitHub org API, never in the git objects (`gh api repos/anthropos-work/jobsimulation --jq .archived`; census signal 6, [`platform-alignment.md`](../ops/platform-alignment.md)). A dated archive assertion is a **snapshot with an expiry**, not a derived fact |
| `roadrunner` | live-standalone | decommissioned | no | **Compose service and `repos.yml` entry both deleted by `d11a403`, in that one commit.** Its message says the clone entry *"was already gone, so the `../roadrunner` build context could no longer resolve"* — **the message is wrong; the diff is the fact.** `git show d11a403 -- repos.yml` shows that very commit deleting `- name: roadrunner` alongside `- name: cms` and `- name: jobsimulation`, and the compose file at `d11a403^` still declares a `roadrunner:` service block (it was one of eleven there; `d11a403` left eight, and `838d907` has since taken that count to **five**). The service was legacy, not unbuildable. (An earlier revision of this row promoted that message into a conclusion — *"the service had been unbuildable, not merely legacy"*; corrected against the diff in M257x. **A commit message is testimony, not evidence** — grade a change by its diff.) Judge0 is reached directly: `JUDGE0_BASE_URL` moved onto `backend` (`docker-compose.yml:59`) for `app/internal/jobsimulation/runner/` (`app/internal/jobsimwiring/wiring.go:123` @ `app` `2035f9a` — a pin, per the `app` row; the same line at `9d00a313`; it was five lines earlier at `b948604`, which is the number `d11a403`'s own message quotes, and it is not repeated here for the reason the `app` row states). **The prod contradiction is now explained but still not verified:** `roadrunner/terraform/main.tf:19` remains `service_desired_count = 1` — last changed at **`84a4b4f` (2025-12-15)**, the commit that first added `terraform/main.tf`, and untouched by everything up to the repo's HEAD `87d8d44` (2026-06-19). **That count is not a decision about the fold; it predates it by seven months and nobody has been back.** (An earlier revision of this row dated it to `e45eb61` (2026-05-27) — that commit is the file's most recent touch but it changed **line 11 only**, a one-line module-source URL swap whose own message says *"Module contents are identical; this is a pure source-URL swap"*. `git blame -L 19,19` names `84a4b4f`; a file-level `git log` is not line provenance.) **This row was RIGHT and the service doc was WRONG for four readings** — [`roadrunner.md`](../services/roadrunner.md) dated the line to `87d8d44` (the repo HEAD, which touches only a workflow file) and **never named `84a4b4f` anywhere**, while pointing the reader here two lines below its own error. Corrected at M257x iter-115; re-verified at that iter, unchanged. Meanwhile the authoritative rollback declaration lives in **infrastructure's `services.tf`** — a repo this map has never read — which `repos.yml:9-10` @ `0dab54d` said in words until `838d907` rewrote that header (see the `cms` row). Losing the citation did not settle the disagreement; it removed the one pointer this map had to the thing that could. Repo **not** archived |
| `sentinel` | live-standalone | live-standalone | yes | `sentinel/terraform/main.tf:19` `= 1`; `docker-compose.yml:5`, own `sentinel` schema via `search_path=sentinel` (`:18`) **despite `migrations: false`** (`repos.yml:18-20`) — the Trap-A row. Since `838d907` it is the **only** Go service besides `app` that a local stack clones or runs, which `repos.yml` now states in its own header (`repos.yml:12-13`) |
| `storage` | merged-into-app | decommissioned | no | **The v9.0 fold is finished on both sides, and the container is now gone.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service **and** the `repos.yml` clone entry in one commit, so `make init` no longer clones it and **the `storage-legacy` profile is gone** — removed with the service it selected. Asking compose for a retired token still exits 0 and starts only the always-on floor (postgresql, redis, sentinel), so the stack looks alive with the application absent: grade a documented compose command on *does it still select something*, never on *does it still parse*. This row read `mid-fold` for four M257x iterations (iter-64 → iter-68), then `merged-into-app` with a legacy rollback container for exactly one platform day. **Prod — the ECS service is DELETED, not scaled to zero.** The module is **18 lines** and declares no service at all; `storage/terraform/main.tf:9-11` says it in the platform's own words — *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS."* The module deliberately survives, because deleting the block would destroy the two buckets, the CloudFront distribution and the media DNS record along with the `prevent_destroy` guards that are read from configuration (`:13-16`). **There is no custody-transfer clause in that file and `:18` is not one** — line 18 is the module's closing comment, *"See outputs.tf — consumers should reference these by output, never by literal name."* The word *custody* occurs **0** times in the entire storage repo at `9f8cb53` (`git -C stack-demo/storage grep -n -i custody 9f8cb532` → no match). **M903 was never executed and is now explicitly superseded.** Its single occurrence in the repo is in a different file, `storage/terraform/storage.tf:22-25`: *"An earlier plan, M903, instead proposed `moved`-ing these resources into a local module in the infrastructure repo. That was never executed, and the shipped design supersedes it: the assets stay here and the module stays declared."* There are **no `moved` blocks** anywhere in the repo. The plan text itself lives in a third repo — `app/knowledge/plan/releases/09.00-support-in-app/m903-s3-custody/` @ app `2035f9a` — and self-declares `status: planned` in its `overview.md` front-matter, and *"Planned, not started — no terraform applied, no branch"* in its `progress.md` (only the read-only P0.2 state capture is ticked; Steps 1–4 are all open). `d3e6d32` (2026-08-05, *"correct the guard header, which contradicted the shipped design"*) is the commit that retired it, in its own words: *"M903 never ran."* **This corrects a paired claim.** This map and its neighbours cited line 38 of that file for `service_desired_count = 0` and paired it with messenger's. It was true up to `9f8cb53`^ and is now past EOF in an 18-line file — the line number is not re-pointable because the construct it named no longer exists. **Scaled-to-zero is `messenger`'s state, not storage's**, so the paired sentence has to be split, not re-anchored. **Config side:** `STORAGE_RPC_ADDR` occurs **0** times across `docker-compose.yml`, `common.yml` and `.env_example`; `STORAGE_S3_BUCKET` / `STORAGE_S3_PUBLIC_BUCKET` sit on `backend` (`docker-compose.yml:82`, `:83`) with the reason in-comment (`:73-75`), and `backend`'s `depends_on` block now states the disappearance directly (`:102-103` — *"storage, messenger and customerio-sync are not services any more — this one container serves all three in-process"*). **Consumer side:** `app` serves object storage in-process and unconditionally — `internalstorage.NewManager` / `NewPublicManager` at `app/main.go:524`, `:525`, consumed at `:547` (the resource manager) and `:1102` (`cmsStorage`); the bucket names are read from the environment at `:516`, `:517` and guarded against empty outside a developer machine at `:518-523`; the constants at `app/internal/storage/service.go:22`, `:24` are the bucket **env-var NAMES** (`EnvBucket = "STORAGE_S3_BUCKET"`, `EnvPublicBucket = "STORAGE_S3_PUBLIC_BUCKET"`), not the bucket names themselves. **`STORAGE_RPC_ADDR` is read by nothing** — a Go grep at `app` `2035f9a` (a pin, per the `app` row) returns **3 hits, every one of them a comment** (`app/main.go:504`, `app/internal/jobsimwiring/wiring.go:101`, `app/internal/storagens/callsites_test.go:189`), and the first says it in words: *"the standalone service takes no traffic and STORAGE_RPC_ADDR is gone."* **State the ref or the sentence flips:** at the OLDER `b948604` (v1.366.0) it is genuinely read, in `main.go` **and in all three `cmd/` tools** — `git -C stack-demo/app grep -n STORAGE_RPC_ADDR b948604 -- '*.go'` returns **15** hits, of which **7 are env lookups** — `main.go` ×3, `internal/jobsimwiring/wiring.go` ×1 and one in each of the three `cmd/` tools (the other 8 are two doc comments, two error strings, **three** `t.Setenv` calls and **one `t.Fatal` message** — the fourth test-file hit *names* the variable inside a failure string rather than setting it, so counting all four as `t.Setenv` overstates the arrange side by one). Six of the seven are spelled `os.Getenv`; the seventh uses a lowercase `getenv` helper, which is why a regex fitted to `os.Getenv` reports **six** and is not wrong to. **Line anchors for that side are deliberately omitted:** they resolve at `b948604` and nowhere else, and naming a second ref inside this cell makes every anchor in it ungradeable (the citation resolver reads one ref per block and reports `ambiguous` when a block names two — M257x run-53). Run the grep; it is the ref, not the line number, that carries the claim. That interval — `b948604` → `9d00a313` — *is* the consumer half of the fold, and it is what moved this row off `mid-fold`. The repo is **not** archived and **not** deleted: `repos.yml`'s header now says so in as many words — *"None of them are deleted"* — and tells you to clone it by hand to read the pre-merge source (`repos.yml:2-10`) |
| `messenger` | merged-into-app | decommissioned | no | **Folded in the same v9.0 program as `storage`, and its container went the same way, one day later.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service and the `repos.yml` clone entry together, and **the `messenger` profile is gone** with them — the same retired-token reading as `storage`'s row. **Prod — stopped, and here "scaled to zero" is still the right words:** `messenger/terraform/main.tf:29` `service_desired_count = 0` in a module that is otherwise intact (121 lines), the cms precedent again, with the reason in-comment (`:19-25` — `app` has taken the Redis consumer group over, and an untargeted `terraform apply` would silently undo a hand-made `aws ecs update-service`); the image and task definition stay declared as the rollback path (`:27-28`). **Do not pair this with `storage`'s prod state:** that module's service block is deleted outright, so a sentence reading *"each `service_desired_count = 0`"* is half false. **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's and is a deliberate literal (`:1416-1421`). **All of it sits behind a switch, which is new since iter-68:** `MESSENGER_ENABLED` (resolved at `app/main.go:285`, read at `:1445` and `:1552`) is **OFF when unset on a developer machine and a BOOT FAILURE when unset in a deployed one** (`app/env_guards.go:92-111`); prod sets it to `"true"` in the task definition (`app/terraform/main.tf:415-416`), and compose sets it **nowhere on purpose** — pinning it to `false` there would override `.env` and make opting in impossible (`docker-compose.yml:84-92`). **The RPC edge is gone, not re-pointed.** messenger's compose block was the only place in the platform that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` or `SKILLER_RPC_ADDR`; `d11a403` had re-pointed two of them at the monolith by hand, and `838d907` deleted the block, so **all four are now set by no compose file at all** — there are **zero `*_RPC_ADDR` variables anywhere in compose**, and the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is **`backend → sentinel`** (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48`). **It is not the only cross-process edge, and compose does not set exactly one service address:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is a live container in the default `core` profile, `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). **This row generalised that to *"the only cross-process service address left in a local stack"* until M257x iter-102** — the `*_RPC_ADDR`-is-zero half was and is true; the generalisation from *no RPC edge* to *no cross-process edge* was not, and the `gotenberg` row of this same table already graded that container `live-standalone` on a fresh local stack. `messenger/cmd/root.go:120-140` still reads all four, which is what made that re-point necessary while messenger was a process; it is not one any more. Repo **not** archived and **not** deleted (`repos.yml:2-10`) |
| `next-web-app` | external (Vercel) | live-standalone | yes | `repos.yml:23-25`; `docker-compose.yml:143` (`frontend` profile, `:168`). Points at `backend` directly since the router drop — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` resolves to `…:8082/graphql/query`, baked as a build arg (`docker-compose.yml:151`) and set again in the environment (`:160`) |
| `studio-desk` | live-standalone | live-standalone | yes | `repos.yml:26-28`; `docker-compose.yml:112` (`studio-desk` profile, `:141`). Same re-point — `VITE_GRAPHQL_ENDPOINT` build arg at `:119`, environment at `:135` |
| `graphql-wundergraph` | live-standalone | decommissioned | no | **the router, dropped from local dev mid-milestone.** Deleted from `repos.yml` **and** compose by `b56d731` + `360efd4`, merged `2adcf71` (2026-07-31); local dev now points at `backend`. In prod it is still declared — `graphql-wundergraph/terraform/main.tf:20` `= 1` — while the **repo is ARCHIVED on GitHub 2026-07-30**. Supergraph is **one** subgraph: `supergraph-config-prod.yaml` lists `backend` alone, `schemas/` holds `backend.graphqls` alone, `subgraphs.conf` = `BACKEND=v1.360.0` (folded by `915da06`, 2026-07-29) |
| `skiller` | merged-into-app | decommissioned | no | removed from compose + `repos.yml` by platform `21429b7` (2026-07-07); code in `app/internal/skiller/`; taxonomy data in `public`. **Repo ARCHIVED 2026-07-01** |
| `skillpath` | merged-into-app | decommissioned | no | decommissioned by platform `a4db680` (2026-07-21, M507); code in `app/internal/skillpath/`; session state in `public.skill_path_sessions`. **Repo ARCHIVED 2026-07-31** |
| `chronos` | decommissioned | decommissioned | no | removed from orchestration by platform `045857c` (2026-04-17). **Repo is NOT archived on GitHub** (last push 2026-04-23) — the corpus called it archived; the org disagrees |
| `intelligence` | decommissioned | decommissioned | no | removed from orchestration by platform `fdfa189` (2026-04-17). **Repo ARCHIVED 2026-04-02** |
| `customerio-sync` | merged-into-app | decommissioned | no | **A state transition this map had never recorded, and one no membership assertion could have caught** — it was never in `repos.yml`, so directions A and B are both blind to it. It was `live-standalone` on both sides until `838d907` (merged `0c91421`, 2026-08-05), which deleted the compose service — it had been built straight from a git URL rather than cloned — and with it **the `customerio-sync` profile is gone**. The commit states a hazard — *"was still in the `all` profile, so `make up-all` started a second Brevo contact pusher alongside backend's own."* — and **the second half of that sentence is false; it is quoted here as the platform's wording, not endorsed** (corrected M257x iter-102). The `all`-profile half is true (`profiles: [customerio-sync, all]`, `0dab54d:docker-compose.yml:154`); the *"second pusher"* half is not, because `backend`'s own in-process pusher is gated behind `CUSTOMERIO_SYNC_ENABLED`, unset and therefore **off** on a developer machine — so `make up-all` started exactly **one**. **This is the corpus inheriting a false claim by quoting a commit message as authoritative**, which is worth naming as a class: a platform commit message is evidence of *intent*, never a measurement. **Consumer side:** the code is `app/internal/customeriosync/`, constructed at `app/main.go:395` (`customeriosync.New`) behind `CUSTOMERIO_SYNC_ENABLED` (resolved at `:286`, read at `:394`) — same switch semantics as messenger's: off when unset on a developer machine, a boot failure when unset in a deployed one (`app/env_guards.go:92-111`), and set to `"true"` in prod's task definition (`app/terraform/main.tf:419-420`). compose sets it nowhere, deliberately, and says why in-comment (`docker-compose.yml:84-92`); `backend`'s `depends_on` block states the disappearance (`:102-103`). **The prod half is asserted from `app`'s side only:** the standalone's own terraform lives in a repo that has never been in the clone set and that this map has therefore never read — the same gap the `roadrunner` row carries, recorded rather than papered over |
| `db-backup` | live-standalone | — | no | production-only; no compose service and no `repos.yml` entry at `0c91421` |
| `anthropos-studio-room` | merged-into-app | merged-into-app | no | the Python generation pipeline is pulled into the `app` image by CI and orchestrated from `app/internal/cms/studio/`, which spawns it as a subprocess. Not a deployment, not in `repos.yml`. **The repo name is `anthropos-studio-room`, not `studio-room`** |
| `ant-academy` | external (Vercel) | external | no | deliberately absent from `repos.yml` — run natively, never containerised |
| `gotenberg` | external | live-standalone | no | third-party image, `docker-compose.yml:170-171` (`gotenberg/gotenberg:8`), default `core` profile (`:183` — renamed from `graphql` by `0dab54d`, since the WunderGraph router the name described is gone) |
| `colony` | library | library | no | private Go module (framework + `colony/authn`); pulled at Docker build via `GH_PAT`/`GOPRIVATE`, never cloned by `make init` |
| `proto` | library | library | no | private Go module — RPC contracts + domain types |
| `ai` | library | library | no | private Go module — the multi-provider `ai.AI` wrapper |
| `authn` | library | library | no | legacy standalone module; the live copy ships **inside** colony as `colony/authn` |
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
**enumerated from the file, not carried forward as a sum**: `app` · `sentinel` · `next-web-app` ·
`studio-desk` (a 28-line `repos.yml` @ `0c91421`). **It read 9 until `d11a403`** removed the
`cms` / `jobsimulation` / `roadrunner` entries, **and 6 until `838d907`** removed `storage` and `messenger`;
the 46 does not move with either, because all five *are* named by corpus documents. The table below
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
[§8 of the protocol](../ops/platform-alignment.md#8-fence--so-it-cannot-silently-recur). It reads this file
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
| `docker-compose.yml` citations — **25** of them | **fenced** (F) |
| every other `file:line` — terraform, `repos.yml`, `common.yml`, `app/…` (**70** citations, + **1** outside any service block) | **resolution + range only.** No derivable notion of *whose* line it is, so F does not pretend |
| the prod column, PR/rollback narrative, §5's ordering | **prose-under-review.** Neither derived nor fenced. Re-check at each platform ref bump |

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
