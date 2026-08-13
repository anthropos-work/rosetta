# TIER-1 ADJUDICATION BATCH 10 — 44 pairs

Each item: the CLAIMING UNIT (corpus prose) and the CITED CONTENT it points at (source lines,
line-numbered, +-3 lines of context). Answer ONE question per item.

## 10-001
- **id**: `B10-001`
- **corpus site**: `corpus/architecture/platform-migration-status.md:87-87` (table-row)
- **citation**: `docker-compose.yml:28`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `app` | live-standalone | live-standalone | yes | the monolith. `app/terraform/main.tf:181` `service_desired_count = 1`; **the only migrating repo** — `repos.yml:14-17` (`migrations: true`, `schema: public`); compose service is named `backend` (`docker-compose.yml:28`). Owns **seven** domains in-process — the four folded before v9.0, plus storage, messenger and customerio-sync — each with **its own** wiring call site: skiller `app/main.go:690` (`skiller.NewSkillerManager`), jobsimulation `:721` (`jobsimwiring.Wire`), skillpath `:751` (`skillpath.NewSessionManager`), cms `:1153` (`appcms.Wire`), storage `:524` (`internalstorage.NewManager`), messenger `:1471` (`msgadapters.Wire`), customerio-sync `:395` (`customeriosync.New`) — `app/internal/{cms,jobsimulation,skiller,skillpath,storage,messenger,customeriosync}/`. **Anchors re-resolved M257x iter-87 at `app` `2035f9a` (post-v1.369.0) — a PIN, not a moving label.** `2035f9a` *was* `origin/main` on 2026-08-05; re-checked 2026-08-06 it is **five commits behind**, and `origin/main` is now **`ad9f3c49`**. Those five touch `.claude/skills/publish/SKILL.md`, `CLAUDE.md`, `knowledge/deployment.md`, `terraform/main.tf` and `terraform/variables.tf` — **no Go source at all**, and the single `terraform/main.tf` hunk rewrites one precondition error message in place, so every anchor in this cell resolves to the same construct at either ref — `ad9f3c49` is a currency note, not a fifth re-derivation, which is why the count below still reads four. **This cell — and three others in this table — wrote the sha as `origin/main` until M257x iter-102:** the sha is a pin and still means what it meant; it is the *label* that expired, and a label that moves under a citation is how a correct anchor becomes a wrong one without anybody editing it. The six that were already cited have moved at **every one of the four refs this map has read in a week** — `5ba17044` v1.363.2, `b948604` v1.366.0, `9d00a313` v1.367.0 (iter-68), and now `2035f9a`; three re-derivations for six anchors, none of them caused by a change to the code they point at. **The older refs are named without their line numbers on purpose** — a block naming two refs is `ambiguous` to the citation resolver, which then falls back to origin/main and grades every anchor in the cell against a file the cell did not mean (M257x run-53; `storage`'s row omits them for the same reaso
```

**CITED CONTENT**

```
    25        postgresql:
    26          condition: service_healthy
    27  
    28    backend:
    29      build:
    30        # Overridable so a branch checked out elsewhere (e.g. a git worktree) can be built
    31        # without touching the default clone:
```

## 10-002
- **id**: `B10-002`
- **corpus site**: `corpus/architecture/platform-migration-status.md:87-87` (table-row)
- **citation**: `app/main.go:690`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| `app` | live-standalone | live-standalone | yes | the monolith. `app/terraform/main.tf:181` `service_desired_count = 1`; **the only migrating repo** — `repos.yml:14-17` (`migrations: true`, `schema: public`); compose service is named `backend` (`docker-compose.yml:28`). Owns **seven** domains in-process — the four folded before v9.0, plus storage, messenger and customerio-sync — each with **its own** wiring call site: skiller `app/main.go:690` (`skiller.NewSkillerManager`), jobsimulation `:721` (`jobsimwiring.Wire`), skillpath `:751` (`skillpath.NewSessionManager`), cms `:1153` (`appcms.Wire`), storage `:524` (`internalstorage.NewManager`), messenger `:1471` (`msgadapters.Wire`), customerio-sync `:395` (`customeriosync.New`) — `app/internal/{cms,jobsimulation,skiller,skillpath,storage,messenger,customeriosync}/`. **Anchors re-resolved M257x iter-87 at `app` `2035f9a` (post-v1.369.0) — a PIN, not a moving label.** `2035f9a` *was* `origin/main` on 2026-08-05; re-checked 2026-08-06 it is **five commits behind**, and `origin/main` is now **`ad9f3c49`**. Those five touch `.claude/skills/publish/SKILL.md`, `CLAUDE.md`, `knowledge/deployment.md`, `terraform/main.tf` and `terraform/variables.tf` — **no Go source at all**, and the single `terraform/main.tf` hunk rewrites one precondition error message in place, so every anchor in this cell resolves to the same construct at either ref — `ad9f3c49` is a currency note, not a fifth re-derivation, which is why the count below still reads four. **This cell — and three others in this table — wrote the sha as `origin/main` until M257x iter-102:** the sha is a pin and still means what it meant; it is the *label* that expired, and a label that moves under a citation is how a correct anchor becomes a wrong one without anybody editing it. The six that were already cited have moved at **every one of the four refs this map has read in a week** — `5ba17044` v1.363.2, `b948604` v1.366.0, `9d00a313` v1.367.0 (iter-68), and now `2035f9a`; three re-derivations for six anchors, none of them caused by a change to the code they point at. **The older refs are named without their line numbers on purpose** — a block naming two refs is `ambiguous` to the citation resolver, which then falls back to origin/main and grades every anchor in the cell against a file the cell did not mean (M257x run-53; `storage`'s row omits them for the same reaso
```

**CITED CONTENT**

```
   687  	// translationManager).
   688  	skillTaxonomyManager := skilltaxonomy.NewTaxonomyManager(logger, ent, orgManager, skillerAIManager, embeddingManager, translationManager)
   689  	jobRoleManager := jobrole.NewJobRoleManager(logger, ent, orgManager, skillTaxonomyManager, skillerAIManager, embeddingManager, pub, workerClient.Client, redisClientStream)
   690  	skillerManager := skiller.NewSkillerManager(logger, jobRoleManager, skillTaxonomyManager, localizationManager)
   691  
   692  	// jobsim-in-app: app's IN-PROCESS skiller handler, serving the external RPC mux registration below
   693  	// (messenger + any remaining external consumers). The ported jobsim validator no longer calls it:
```

## 10-003
- **id**: `B10-003`
- **corpus site**: `corpus/architecture/platform-migration-status.md:87-87` (table-row)
- **citation**: `app/internal/jobsimwiring/wiring.go:123`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimwiring/wiring.go`  (283 lines)

**CLAIMING UNIT**

```md
| `app` | live-standalone | live-standalone | yes | the monolith. `app/terraform/main.tf:181` `service_desired_count = 1`; **the only migrating repo** — `repos.yml:14-17` (`migrations: true`, `schema: public`); compose service is named `backend` (`docker-compose.yml:28`). Owns **seven** domains in-process — the four folded before v9.0, plus storage, messenger and customerio-sync — each with **its own** wiring call site: skiller `app/main.go:690` (`skiller.NewSkillerManager`), jobsimulation `:721` (`jobsimwiring.Wire`), skillpath `:751` (`skillpath.NewSessionManager`), cms `:1153` (`appcms.Wire`), storage `:524` (`internalstorage.NewManager`), messenger `:1471` (`msgadapters.Wire`), customerio-sync `:395` (`customeriosync.New`) — `app/internal/{cms,jobsimulation,skiller,skillpath,storage,messenger,customeriosync}/`. **Anchors re-resolved M257x iter-87 at `app` `2035f9a` (post-v1.369.0) — a PIN, not a moving label.** `2035f9a` *was* `origin/main` on 2026-08-05; re-checked 2026-08-06 it is **five commits behind**, and `origin/main` is now **`ad9f3c49`**. Those five touch `.claude/skills/publish/SKILL.md`, `CLAUDE.md`, `knowledge/deployment.md`, `terraform/main.tf` and `terraform/variables.tf` — **no Go source at all**, and the single `terraform/main.tf` hunk rewrites one precondition error message in place, so every anchor in this cell resolves to the same construct at either ref — `ad9f3c49` is a currency note, not a fifth re-derivation, which is why the count below still reads four. **This cell — and three others in this table — wrote the sha as `origin/main` until M257x iter-102:** the sha is a pin and still means what it meant; it is the *label* that expired, and a label that moves under a citation is how a correct anchor becomes a wrong one without anybody editing it. The six that were already cited have moved at **every one of the four refs this map has read in a week** — `5ba17044` v1.363.2, `b948604` v1.366.0, `9d00a313` v1.367.0 (iter-68), and now `2035f9a`; three re-derivations for six anchors, none of them caused by a change to the code they point at. **The older refs are named without their line numbers on purpose** — a block naming two refs is `ambiguous` to the citation resolver, which then falls back to origin/main and grades every anchor in the cell against a file the cell did not mean (M257x run-53; `storage`'s row omits them for the same reaso
```

**CITED CONTENT**

```
   120  	storageV1Client := appstorage.NewClient(inAppStorage, storagens.JobSimulation).V1
   121  	// Judge0 sandbox runner (IN-PROCESS; replaces the removed roadrunner RPC edge — resync to jobsim main
   122  	// v0.253.0, which deleted chronos + realtime + the roadrunner-submission event).
   123  	runnerManager := jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))
   124  
   125  	// --- Asynq producer client (task-type / queue name strings are frozen — M705 contract).
   126  	workerIndex, _ := strconv.Atoi(getenv("REDIS_WORKER_INDEX"))
```

## 10-004
- **id**: `B10-004`
- **corpus site**: `corpus/architecture/platform-migration-status.md:88-88` (table-row)
- **citation**: `cms/terraform/main.tf:39`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/terraform/main.tf`  (192 lines)

**CLAIMING UNIT**

```md
| `cms` | merged-into-app | decommissioned | no | `cms/terraform/main.tf:39` `service_desired_count = 0`; code in `app/internal/cms/`; folded by platform `236771f` (2026-07-29, cms-in-app v8.0). **Compose service and `repos.yml` entry both deleted by `d11a403`** (merged `ef32d4c`, 2026-08-03) — `make init` no longer clones it. Repo **not** archived. **The pointer to the prod rollback path is gone from `repos.yml`:** its header named infrastructure's `services.tf` as that path until M810 (`repos.yml:9-10` @ `0dab54d`), and `838d907` rewrote the header without it — what stands now says only that the frozen repos own no schema, no compose service and no clone entry, and that *"None of them are deleted"* (`repos.yml:2-10`). **Whether that rollback declaration still stands is not something this map can see** — it never could, since infrastructure has never been in the clone set; it had a pointer, and `838d907` removed the pointer. Absence of the sentence is not evidence the declaration went with it. **M257x iter-92 — cms has since taken an M810 step, and it points the OTHER way:** `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** `.github/workflows/build-production.yml`, subject *"the cms ECR repository is decommissioned (M810)"*, body: M810 *"deletes `module \"cms_euwest1\"` … which destroys the ECS service and the production-cms ECR repository"*, the workflow being dropped because it *"would try to push an image into a registry that no longer exists."* So this repo now holds **two measured facts pointing opposite ways** — `cms/terraform/main.tf:39` still declaring the module, and a CI commit asserting the registry is already gone. **Still UNMEASURABLE here**, and now unmeasurable with contrary evidence on both sides rather than one: report both, assert neither |
```

**CITED CONTENT**

```
    36    tags                           = var.tags
    37    aws_region                     = var.aws_region
    38    project                        = local.project
    39    service_desired_count          = 0
    40    service_cpu                    = local.service_cpu
    41    service_memory                 = local.service_memory
    42    health_check_path              = "/_meta"
```

## 10-005
- **id**: `B10-005`
- **corpus site**: `corpus/architecture/platform-migration-status.md:88-88` (table-row)
- **citation**: `repos.yml:9-10`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
| `cms` | merged-into-app | decommissioned | no | `cms/terraform/main.tf:39` `service_desired_count = 0`; code in `app/internal/cms/`; folded by platform `236771f` (2026-07-29, cms-in-app v8.0). **Compose service and `repos.yml` entry both deleted by `d11a403`** (merged `ef32d4c`, 2026-08-03) — `make init` no longer clones it. Repo **not** archived. **The pointer to the prod rollback path is gone from `repos.yml`:** its header named infrastructure's `services.tf` as that path until M810 (`repos.yml:9-10` @ `0dab54d`), and `838d907` rewrote the header without it — what stands now says only that the frozen repos own no schema, no compose service and no clone entry, and that *"None of them are deleted"* (`repos.yml:2-10`). **Whether that rollback declaration still stands is not something this map can see** — it never could, since infrastructure has never been in the clone set; it had a pointer, and `838d907` removed the pointer. Absence of the sentence is not evidence the declaration went with it. **M257x iter-92 — cms has since taken an M810 step, and it points the OTHER way:** `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** `.github/workflows/build-production.yml`, subject *"the cms ECR repository is decommissioned (M810)"*, body: M810 *"deletes `module \"cms_euwest1\"` … which destroys the ECS service and the production-cms ECR repository"*, the workflow being dropped because it *"would try to push an image into a registry that no longer exists."* So this repo now holds **two measured facts pointing opposite ways** — `cms/terraform/main.tf:39` still declaring the module, and a CI commit asserting the registry is already gone. **Still UNMEASURABLE here**, and now unmeasurable with contrary evidence on both sides rather than one: report both, assert neither |
```

**CITED CONTENT**

```
     6    # served in-process. Their Ent entities were re-created in the `public` schema
     7    # under app/terraform/migrations/, so `app` is the ONLY repo with migrations to
     8    # run. Those repos are frozen legacy: they own no local schema, no compose service
     9    # and no clone entry here. `make init` therefore does not clone them — clone them
    10    # by hand if you need to read the pre-merge source. None of them are deleted.
    11    #
    12    # `sentinel` is the one Go service still deployed alongside `backend`, so it is
    13    # the only other backend clone local dev needs.
```

## 10-006
- **id**: `B10-006`
- **corpus site**: `corpus/architecture/platform-migration-status.md:88-88` (table-row)
- **citation**: `repos.yml:2-10`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
| `cms` | merged-into-app | decommissioned | no | `cms/terraform/main.tf:39` `service_desired_count = 0`; code in `app/internal/cms/`; folded by platform `236771f` (2026-07-29, cms-in-app v8.0). **Compose service and `repos.yml` entry both deleted by `d11a403`** (merged `ef32d4c`, 2026-08-03) — `make init` no longer clones it. Repo **not** archived. **The pointer to the prod rollback path is gone from `repos.yml`:** its header named infrastructure's `services.tf` as that path until M810 (`repos.yml:9-10` @ `0dab54d`), and `838d907` rewrote the header without it — what stands now says only that the frozen repos own no schema, no compose service and no clone entry, and that *"None of them are deleted"* (`repos.yml:2-10`). **Whether that rollback declaration still stands is not something this map can see** — it never could, since infrastructure has never been in the clone set; it had a pointer, and `838d907` removed the pointer. Absence of the sentence is not evidence the declaration went with it. **M257x iter-92 — cms has since taken an M810 step, and it points the OTHER way:** `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** `.github/workflows/build-production.yml`, subject *"the cms ECR repository is decommissioned (M810)"*, body: M810 *"deletes `module \"cms_euwest1\"` … which destroys the ECS service and the production-cms ECR repository"*, the workflow being dropped because it *"would try to push an image into a registry that no longer exists."* So this repo now holds **two measured facts pointing opposite ways** — `cms/terraform/main.tf:39` still declaring the module, and a CI commit asserting the registry is already gone. **Still UNMEASURABLE here**, and now unmeasurable with contrary evidence on both sides rather than one: report both, assert neither |
```

**CITED CONTENT**

```
     1  repos:
     2    # Go backend services
     3    #
     4    # `app` (deployed as `backend`) is the merged monolith: skiller, skillpath,
     5    # roadrunner, jobsimulation, cms, messenger, storage and customerio-sync are all
     6    # served in-process. Their Ent entities were re-created in the `public` schema
     7    # under app/terraform/migrations/, so `app` is the ONLY repo with migrations to
     8    # run. Those repos are frozen legacy: they own no local schema, no compose service
     9    # and no clone entry here. `make init` therefore does not clone them — clone them
    10    # by hand if you need to read the pre-merge source. None of them are deleted.
    11    #
    12    # `sentinel` is the one Go service still deployed alongside `backend`, so it is
    13    # the only other backend clone local dev needs.
```

## 10-007
- **id**: `B10-007`
- **corpus site**: `corpus/architecture/platform-migration-status.md:89-89` (table-row)
- **citation**: `jobsimulation/terraform/main.tf:15-22`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/jobsimulation/terraform/main.tf`  (344 lines)

**CLAIMING UNIT**

```md
| `jobsimulation` | merged-into-app | decommissioned | no | **M810 has LANDED for this service, and this map had not noticed** — `6092c6d2` (*"remove the jobsimulation ECS service and ECR repository (M810)"*) deleted the `module "jobsimulation"` block outright, destroying the ECS service, task definition, ECR repository, task/execution IAM roles, security group, Cloud Map entry, log group and alarms (`jobsimulation/terraform/main.tf:15-22` @ `82cb66e`). It had run at `service_desired_count = 0` since app v1.360.0; **that line does not exist any more**, which is why this row cites the decommission comment and not a count — the old citation to it resolved to an unrelated comment about the atlas tracker, the exact silent-slide failure [§4](#4-the-fence)'s assertion F was built for, in the one file class F does not reach. The module deliberately survives because it still owns the **LiveKit and Chime recording buckets `backend` reads by literal name**, the `/production/jobsimulation/*` SSM parameters, and the atlas tracker for the legacy `jobsimulation` schema — dropping that schema is a separate M810 step (`:24-40`). **Do not generalise M810 from this row — and do not read `cms` as standing still either.** `cms` holds **two measured facts pointing opposite ways**: `cms/terraform/main.tf:39` still declares `service_desired_count = 0` in a module that still exists, **and** `6efa1d5` (merged `f38c0c4`, 2026-08-04) deleted that repo's build-production workflow because *"the cms ECR repository is decommissioned (M810)"*. The `cms` row directly above **reports both and asserts neither**, and so must anything reading across from this one; the destruction itself lands in **infrastructure**, which is in no clone set. **This row said *"`cms` has not moved"* until M257x iter-102** — a flat assertion the row above it had already retracted, which is how one table came to hold both readings. Code in `app/internal/jobsimulation/`, wired unconditionally at `app/main.go:721` (`jobsimwiring.Wire`, @ `app` `2035f9a` — a pin; see the re-measured banner at the top of this file for where `origin/main` has gone since); tables re-created in `public`; folded by platform `236771f`. **Compose service and `repos.yml` entry both deleted by `d11a403`.** **Repo archive state — REPORT BOTH, ASSERT NEITHER.** This row asserted a GitHub archive on 2026-07-31 until M257x, and the clone refutes th
```

**CITED CONTENT**

```
    12    }
    13  }
    14  
    15  // Inspect the target database and load its state.
    16  // This is used to determine which migration to run.
    17  data "atlas_migration" "jobsimulation_migrations" {
    18    dir = "${path.module}/migrations?format=atlas"
    19    url = "${aws_ssm_parameter.db_connection.value}?search_path=jobsimulation"
    20  }
    21  
    22  // Sync the state of the target database with the migrations directory.
    23  resource "atlas_migration" "jobsimulation_migrations" {
    24    dir              = "${path.module}/migrations?format=atlas"
    25    version          = data.atlas_migration.jobsimulation_migrations.latest # Use latest to run all migrations
```

## 10-008
- **id**: `B10-008`
- **corpus site**: `corpus/architecture/platform-migration-status.md:89-89` (table-row)
- **citation**: `cms/terraform/main.tf:39`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/terraform/main.tf`  (192 lines)

**CLAIMING UNIT**

```md
| `jobsimulation` | merged-into-app | decommissioned | no | **M810 has LANDED for this service, and this map had not noticed** — `6092c6d2` (*"remove the jobsimulation ECS service and ECR repository (M810)"*) deleted the `module "jobsimulation"` block outright, destroying the ECS service, task definition, ECR repository, task/execution IAM roles, security group, Cloud Map entry, log group and alarms (`jobsimulation/terraform/main.tf:15-22` @ `82cb66e`). It had run at `service_desired_count = 0` since app v1.360.0; **that line does not exist any more**, which is why this row cites the decommission comment and not a count — the old citation to it resolved to an unrelated comment about the atlas tracker, the exact silent-slide failure [§4](#4-the-fence)'s assertion F was built for, in the one file class F does not reach. The module deliberately survives because it still owns the **LiveKit and Chime recording buckets `backend` reads by literal name**, the `/production/jobsimulation/*` SSM parameters, and the atlas tracker for the legacy `jobsimulation` schema — dropping that schema is a separate M810 step (`:24-40`). **Do not generalise M810 from this row — and do not read `cms` as standing still either.** `cms` holds **two measured facts pointing opposite ways**: `cms/terraform/main.tf:39` still declares `service_desired_count = 0` in a module that still exists, **and** `6efa1d5` (merged `f38c0c4`, 2026-08-04) deleted that repo's build-production workflow because *"the cms ECR repository is decommissioned (M810)"*. The `cms` row directly above **reports both and asserts neither**, and so must anything reading across from this one; the destruction itself lands in **infrastructure**, which is in no clone set. **This row said *"`cms` has not moved"* until M257x iter-102** — a flat assertion the row above it had already retracted, which is how one table came to hold both readings. Code in `app/internal/jobsimulation/`, wired unconditionally at `app/main.go:721` (`jobsimwiring.Wire`, @ `app` `2035f9a` — a pin; see the re-measured banner at the top of this file for where `origin/main` has gone since); tables re-created in `public`; folded by platform `236771f`. **Compose service and `repos.yml` entry both deleted by `d11a403`.** **Repo archive state — REPORT BOTH, ASSERT NEITHER.** This row asserted a GitHub archive on 2026-07-31 until M257x, and the clone refutes th
```

**CITED CONTENT**

```
    36    tags                           = var.tags
    37    aws_region                     = var.aws_region
    38    project                        = local.project
    39    service_desired_count          = 0
    40    service_cpu                    = local.service_cpu
    41    service_memory                 = local.service_memory
    42    health_check_path              = "/_meta"
```

## 10-009
- **id**: `B10-009`
- **corpus site**: `corpus/architecture/platform-migration-status.md:89-89` (table-row)
- **citation**: `app/main.go:721`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| `jobsimulation` | merged-into-app | decommissioned | no | **M810 has LANDED for this service, and this map had not noticed** — `6092c6d2` (*"remove the jobsimulation ECS service and ECR repository (M810)"*) deleted the `module "jobsimulation"` block outright, destroying the ECS service, task definition, ECR repository, task/execution IAM roles, security group, Cloud Map entry, log group and alarms (`jobsimulation/terraform/main.tf:15-22` @ `82cb66e`). It had run at `service_desired_count = 0` since app v1.360.0; **that line does not exist any more**, which is why this row cites the decommission comment and not a count — the old citation to it resolved to an unrelated comment about the atlas tracker, the exact silent-slide failure [§4](#4-the-fence)'s assertion F was built for, in the one file class F does not reach. The module deliberately survives because it still owns the **LiveKit and Chime recording buckets `backend` reads by literal name**, the `/production/jobsimulation/*` SSM parameters, and the atlas tracker for the legacy `jobsimulation` schema — dropping that schema is a separate M810 step (`:24-40`). **Do not generalise M810 from this row — and do not read `cms` as standing still either.** `cms` holds **two measured facts pointing opposite ways**: `cms/terraform/main.tf:39` still declares `service_desired_count = 0` in a module that still exists, **and** `6efa1d5` (merged `f38c0c4`, 2026-08-04) deleted that repo's build-production workflow because *"the cms ECR repository is decommissioned (M810)"*. The `cms` row directly above **reports both and asserts neither**, and so must anything reading across from this one; the destruction itself lands in **infrastructure**, which is in no clone set. **This row said *"`cms` has not moved"* until M257x iter-102** — a flat assertion the row above it had already retracted, which is how one table came to hold both readings. Code in `app/internal/jobsimulation/`, wired unconditionally at `app/main.go:721` (`jobsimwiring.Wire`, @ `app` `2035f9a` — a pin; see the re-measured banner at the top of this file for where `origin/main` has gone since); tables re-created in `public`; folded by platform `236771f`. **Compose service and `repos.yml` entry both deleted by `d11a403`.** **Repo archive state — REPORT BOTH, ASSERT NEITHER.** This row asserted a GitHub archive on 2026-07-31 until M257x, and the clone refutes th
```

**CITED CONTENT**

```
   718  	// jobsim now, so wiring is FATAL, not best-effort: a jobsim-less boot must fail loud, never silently
   719  	// serve a half-wired domain. The GraphQL Session type, the Redis-stream subscribers, and the jobsim
   720  	// Asynq pools are all served by app unconditionally (no dormant gate).
   721  	jobsimDj, err := jobsimwiring.Wire(serverContext, logger, serviceName, ent, pub, redisClientStream, cmsReaderSw, posthogClient, jobsimUsers, jobsimSkiller, copilotDB, authz, storageManager)
   722  	if err != nil {
   723  		log.Fatalf("jobsim-in-app: engine wiring failed (is jobsim env provisioned?): %v", err)
   724  	}
```

## 10-010
- **id**: `B10-010`
- **corpus site**: `corpus/architecture/platform-migration-status.md:90-90` (table-row)
- **citation**: `docker-compose.yml:59`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `roadrunner` | live-standalone | decommissioned | no | **Compose service and `repos.yml` entry both deleted by `d11a403`, in that one commit.** Its message says the clone entry *"was already gone, so the `../roadrunner` build context could no longer resolve"* — **the message is wrong; the diff is the fact.** `git show d11a403 -- repos.yml` shows that very commit deleting `- name: roadrunner` alongside `- name: cms` and `- name: jobsimulation`, and the compose file at `d11a403^` still declares a `roadrunner:` service block (it was one of eleven there; `d11a403` left eight, and `838d907` has since taken that count to **five**). The service was legacy, not unbuildable. (An earlier revision of this row promoted that message into a conclusion — *"the service had been unbuildable, not merely legacy"*; corrected against the diff in M257x. **A commit message is testimony, not evidence** — grade a change by its diff.) Judge0 is reached directly: `JUDGE0_BASE_URL` moved onto `backend` (`docker-compose.yml:59`) for `app/internal/jobsimulation/runner/` (`app/internal/jobsimwiring/wiring.go:123` @ `app` `2035f9a` — a pin, per the `app` row; the same line at `9d00a313`; it was five lines earlier at `b948604`, which is the number `d11a403`'s own message quotes, and it is not repeated here for the reason the `app` row states). **The prod contradiction is now explained but still not verified:** `roadrunner/terraform/main.tf:19` remains `service_desired_count = 1` — last changed at **`84a4b4f` (2025-12-15)**, the commit that first added `terraform/main.tf`, and untouched by everything up to the repo's HEAD `87d8d44` (2026-06-19). **That count is not a decision about the fold; it predates it by seven months and nobody has been back.** (An earlier revision of this row dated it to `e45eb61` (2026-05-27) — that commit is the file's most recent touch but it changed **line 11 only**, a one-line module-source URL swap whose own message says *"Module contents are identical; this is a pure source-URL swap"*. `git blame -L 19,19` names `84a4b4f`; a file-level `git log` is not line provenance.) **This row was RIGHT and the service doc was WRONG for four readings** — [`roadrunner.md`](../services/roadrunner.md) dated the line to `87d8d44` (the repo HEAD, which touches only a workflow file) and **never named `84a4b4f` anywhere**, while pointing the reader here two lines below its own er
```

**CITED CONTENT**

```
    56        - ENVIRONMENT=development
    57        - GOTENBERG_URL=http://gotenberg:3200
    58        - JOBSIMULATION_STREAM=jobsimulation
    59        - JUDGE0_BASE_URL=http://52.48.139.23:2358
    60        - LIVEKIT_AWS_SDK_REGION=eu-central-1
    61        - LIVEKIT_HOST_URL=wss://anthropos-pbvktu3v.livekit.cloud
    62        - LIVEKIT_RECORDINGS_BUCKET_NAME=anthropos-livekit-test
```

## 10-011
- **id**: `B10-011`
- **corpus site**: `corpus/architecture/platform-migration-status.md:90-90` (table-row)
- **citation**: `app/internal/jobsimwiring/wiring.go:123`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimwiring/wiring.go`  (283 lines)

**CLAIMING UNIT**

```md
| `roadrunner` | live-standalone | decommissioned | no | **Compose service and `repos.yml` entry both deleted by `d11a403`, in that one commit.** Its message says the clone entry *"was already gone, so the `../roadrunner` build context could no longer resolve"* — **the message is wrong; the diff is the fact.** `git show d11a403 -- repos.yml` shows that very commit deleting `- name: roadrunner` alongside `- name: cms` and `- name: jobsimulation`, and the compose file at `d11a403^` still declares a `roadrunner:` service block (it was one of eleven there; `d11a403` left eight, and `838d907` has since taken that count to **five**). The service was legacy, not unbuildable. (An earlier revision of this row promoted that message into a conclusion — *"the service had been unbuildable, not merely legacy"*; corrected against the diff in M257x. **A commit message is testimony, not evidence** — grade a change by its diff.) Judge0 is reached directly: `JUDGE0_BASE_URL` moved onto `backend` (`docker-compose.yml:59`) for `app/internal/jobsimulation/runner/` (`app/internal/jobsimwiring/wiring.go:123` @ `app` `2035f9a` — a pin, per the `app` row; the same line at `9d00a313`; it was five lines earlier at `b948604`, which is the number `d11a403`'s own message quotes, and it is not repeated here for the reason the `app` row states). **The prod contradiction is now explained but still not verified:** `roadrunner/terraform/main.tf:19` remains `service_desired_count = 1` — last changed at **`84a4b4f` (2025-12-15)**, the commit that first added `terraform/main.tf`, and untouched by everything up to the repo's HEAD `87d8d44` (2026-06-19). **That count is not a decision about the fold; it predates it by seven months and nobody has been back.** (An earlier revision of this row dated it to `e45eb61` (2026-05-27) — that commit is the file's most recent touch but it changed **line 11 only**, a one-line module-source URL swap whose own message says *"Module contents are identical; this is a pure source-URL swap"*. `git blame -L 19,19` names `84a4b4f`; a file-level `git log` is not line provenance.) **This row was RIGHT and the service doc was WRONG for four readings** — [`roadrunner.md`](../services/roadrunner.md) dated the line to `87d8d44` (the repo HEAD, which touches only a workflow file) and **never named `84a4b4f` anywhere**, while pointing the reader here two lines below its own er
```

**CITED CONTENT**

```
   120  	storageV1Client := appstorage.NewClient(inAppStorage, storagens.JobSimulation).V1
   121  	// Judge0 sandbox runner (IN-PROCESS; replaces the removed roadrunner RPC edge — resync to jobsim main
   122  	// v0.253.0, which deleted chronos + realtime + the roadrunner-submission event).
   123  	runnerManager := jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))
   124  
   125  	// --- Asynq producer client (task-type / queue name strings are frozen — M705 contract).
   126  	workerIndex, _ := strconv.Atoi(getenv("REDIS_WORKER_INDEX"))
```

## 10-012
- **id**: `B10-012`
- **corpus site**: `corpus/architecture/platform-migration-status.md:90-90` (table-row)
- **citation**: `roadrunner/terraform/main.tf:19`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/roadrunner/terraform/main.tf`  (96 lines)

**CLAIMING UNIT**

```md
| `roadrunner` | live-standalone | decommissioned | no | **Compose service and `repos.yml` entry both deleted by `d11a403`, in that one commit.** Its message says the clone entry *"was already gone, so the `../roadrunner` build context could no longer resolve"* — **the message is wrong; the diff is the fact.** `git show d11a403 -- repos.yml` shows that very commit deleting `- name: roadrunner` alongside `- name: cms` and `- name: jobsimulation`, and the compose file at `d11a403^` still declares a `roadrunner:` service block (it was one of eleven there; `d11a403` left eight, and `838d907` has since taken that count to **five**). The service was legacy, not unbuildable. (An earlier revision of this row promoted that message into a conclusion — *"the service had been unbuildable, not merely legacy"*; corrected against the diff in M257x. **A commit message is testimony, not evidence** — grade a change by its diff.) Judge0 is reached directly: `JUDGE0_BASE_URL` moved onto `backend` (`docker-compose.yml:59`) for `app/internal/jobsimulation/runner/` (`app/internal/jobsimwiring/wiring.go:123` @ `app` `2035f9a` — a pin, per the `app` row; the same line at `9d00a313`; it was five lines earlier at `b948604`, which is the number `d11a403`'s own message quotes, and it is not repeated here for the reason the `app` row states). **The prod contradiction is now explained but still not verified:** `roadrunner/terraform/main.tf:19` remains `service_desired_count = 1` — last changed at **`84a4b4f` (2025-12-15)**, the commit that first added `terraform/main.tf`, and untouched by everything up to the repo's HEAD `87d8d44` (2026-06-19). **That count is not a decision about the fold; it predates it by seven months and nobody has been back.** (An earlier revision of this row dated it to `e45eb61` (2026-05-27) — that commit is the file's most recent touch but it changed **line 11 only**, a one-line module-source URL swap whose own message says *"Module contents are identical; this is a pure source-URL swap"*. `git blame -L 19,19` names `84a4b4f`; a file-level `git log` is not line provenance.) **This row was RIGHT and the service doc was WRONG for four readings** — [`roadrunner.md`](../services/roadrunner.md) dated the line to `87d8d44` (the repo HEAD, which touches only a workflow file) and **never named `84a4b4f` anywhere**, while pointing the reader here two lines below its own er
```

**CITED CONTENT**

```
    16    tags                           = var.tags
    17    aws_region                     = var.aws_region
    18    project                        = local.project
    19    service_desired_count          = 1
    20    service_cpu                    = local.service_cpu
    21    service_memory                 = local.service_memory
    22    health_check_path              = "/_meta"
```

## 10-013
- **id**: `B10-013`
- **corpus site**: `corpus/architecture/platform-migration-status.md:90-90` (table-row)
- **citation**: `repos.yml:9-10`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
| `roadrunner` | live-standalone | decommissioned | no | **Compose service and `repos.yml` entry both deleted by `d11a403`, in that one commit.** Its message says the clone entry *"was already gone, so the `../roadrunner` build context could no longer resolve"* — **the message is wrong; the diff is the fact.** `git show d11a403 -- repos.yml` shows that very commit deleting `- name: roadrunner` alongside `- name: cms` and `- name: jobsimulation`, and the compose file at `d11a403^` still declares a `roadrunner:` service block (it was one of eleven there; `d11a403` left eight, and `838d907` has since taken that count to **five**). The service was legacy, not unbuildable. (An earlier revision of this row promoted that message into a conclusion — *"the service had been unbuildable, not merely legacy"*; corrected against the diff in M257x. **A commit message is testimony, not evidence** — grade a change by its diff.) Judge0 is reached directly: `JUDGE0_BASE_URL` moved onto `backend` (`docker-compose.yml:59`) for `app/internal/jobsimulation/runner/` (`app/internal/jobsimwiring/wiring.go:123` @ `app` `2035f9a` — a pin, per the `app` row; the same line at `9d00a313`; it was five lines earlier at `b948604`, which is the number `d11a403`'s own message quotes, and it is not repeated here for the reason the `app` row states). **The prod contradiction is now explained but still not verified:** `roadrunner/terraform/main.tf:19` remains `service_desired_count = 1` — last changed at **`84a4b4f` (2025-12-15)**, the commit that first added `terraform/main.tf`, and untouched by everything up to the repo's HEAD `87d8d44` (2026-06-19). **That count is not a decision about the fold; it predates it by seven months and nobody has been back.** (An earlier revision of this row dated it to `e45eb61` (2026-05-27) — that commit is the file's most recent touch but it changed **line 11 only**, a one-line module-source URL swap whose own message says *"Module contents are identical; this is a pure source-URL swap"*. `git blame -L 19,19` names `84a4b4f`; a file-level `git log` is not line provenance.) **This row was RIGHT and the service doc was WRONG for four readings** — [`roadrunner.md`](../services/roadrunner.md) dated the line to `87d8d44` (the repo HEAD, which touches only a workflow file) and **never named `84a4b4f` anywhere**, while pointing the reader here two lines below its own er
```

**CITED CONTENT**

```
     6    # served in-process. Their Ent entities were re-created in the `public` schema
     7    # under app/terraform/migrations/, so `app` is the ONLY repo with migrations to
     8    # run. Those repos are frozen legacy: they own no local schema, no compose service
     9    # and no clone entry here. `make init` therefore does not clone them — clone them
    10    # by hand if you need to read the pre-merge source. None of them are deleted.
    11    #
    12    # `sentinel` is the one Go service still deployed alongside `backend`, so it is
    13    # the only other backend clone local dev needs.
```

## 10-014
- **id**: `B10-014`
- **corpus site**: `corpus/architecture/platform-migration-status.md:91-91` (table-row)
- **citation**: `sentinel/terraform/main.tf:19`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/sentinel/terraform/main.tf`  (67 lines)

**CLAIMING UNIT**

```md
| `sentinel` | live-standalone | live-standalone | yes | `sentinel/terraform/main.tf:19` `= 1`; `docker-compose.yml:5`, own `sentinel` schema via `search_path=sentinel` (`:18`) **despite `migrations: false`** (`repos.yml:18-20`) — the Trap-A row. Since `838d907` it is the **only** Go service besides `app` that a local stack clones or runs, which `repos.yml` now states in its own header (`repos.yml:12-13`) |
```

**CITED CONTENT**

```
    16    tags                           = var.tags
    17    aws_region                     = var.aws_region
    18    project                        = local.project
    19    service_desired_count          = 1
    20    service_cpu                    = local.service_cpu
    21    service_memory                 = local.service_memory
    22    health_check_path              = "/_meta"
```

## 10-015
- **id**: `B10-015`
- **corpus site**: `corpus/architecture/platform-migration-status.md:91-91` (table-row)
- **citation**: `docker-compose.yml:5`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `sentinel` | live-standalone | live-standalone | yes | `sentinel/terraform/main.tf:19` `= 1`; `docker-compose.yml:5`, own `sentinel` schema via `search_path=sentinel` (`:18`) **despite `migrations: false`** (`repos.yml:18-20`) — the Trap-A row. Since `838d907` it is the **only** Go service besides `app` that a local stack clones or runs, which `repos.yml` now states in its own header (`repos.yml:12-13`) |
```

**CITED CONTENT**

```
     2    - common.yml
     3  
     4  services:
     5    sentinel:
     6      build:
     7        context: ../sentinel
     8        dockerfile: Dockerfile.dev
```

## 10-016
- **id**: `B10-016`
- **corpus site**: `corpus/architecture/platform-migration-status.md:91-91` (table-row)
- **citation**: `repos.yml:18-20`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
| `sentinel` | live-standalone | live-standalone | yes | `sentinel/terraform/main.tf:19` `= 1`; `docker-compose.yml:5`, own `sentinel` schema via `search_path=sentinel` (`:18`) **despite `migrations: false`** (`repos.yml:18-20`) — the Trap-A row. Since `838d907` it is the **only** Go service besides `app` that a local stack clones or runs, which `repos.yml` now states in its own header (`repos.yml:12-13`) |
```

**CITED CONTENT**

```
    15      type: go
    16      migrations: true
    17      schema: public
    18    - name: sentinel
    19      type: go
    20      migrations: false
    21  
    22    # Frontend
    23    - name: next-web-app
```

## 10-017
- **id**: `B10-017`
- **corpus site**: `corpus/architecture/platform-migration-status.md:91-91` (table-row)
- **citation**: `repos.yml:12-13`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
| `sentinel` | live-standalone | live-standalone | yes | `sentinel/terraform/main.tf:19` `= 1`; `docker-compose.yml:5`, own `sentinel` schema via `search_path=sentinel` (`:18`) **despite `migrations: false`** (`repos.yml:18-20`) — the Trap-A row. Since `838d907` it is the **only** Go service besides `app` that a local stack clones or runs, which `repos.yml` now states in its own header (`repos.yml:12-13`) |
```

**CITED CONTENT**

```
     9    # and no clone entry here. `make init` therefore does not clone them — clone them
    10    # by hand if you need to read the pre-merge source. None of them are deleted.
    11    #
    12    # `sentinel` is the one Go service still deployed alongside `backend`, so it is
    13    # the only other backend clone local dev needs.
    14    - name: app
    15      type: go
    16      migrations: true
```

## 10-018
- **id**: `B10-018`
- **corpus site**: `corpus/architecture/platform-migration-status.md:92-92` (table-row)
- **citation**: `storage/terraform/main.tf:9-11`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/storage/terraform/main.tf`  (101 lines)

**CLAIMING UNIT**

```md
| `storage` | merged-into-app | decommissioned | no | **The v9.0 fold is finished on both sides, and the container is now gone.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service **and** the `repos.yml` clone entry in one commit, so `make init` no longer clones it and **the `storage-legacy` profile is gone** — removed with the service it selected. Asking compose for a retired token still exits 0 and starts only the always-on floor (postgresql, redis, sentinel), so the stack looks alive with the application absent: grade a documented compose command on *does it still select something*, never on *does it still parse*. This row read `mid-fold` for four M257x iterations (iter-64 → iter-68), then `merged-into-app` with a legacy rollback container for exactly one platform day. **Prod — the ECS service is DELETED, not scaled to zero.** The module is **18 lines** and declares no service at all; `storage/terraform/main.tf:9-11` says it in the platform's own words — *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS."* The module deliberately survives, because deleting the block would destroy the two buckets, the CloudFront distribution and the media DNS record along with the `prevent_destroy` guards that are read from configuration (`:13-16`). **There is no custody-transfer clause in that file and `:18` is not one** — line 18 is the module's closing comment, *"See outputs.tf — consumers should reference these by output, never by literal name."* The word *custody* occurs **0** times in the entire storage repo at `9f8cb53` (`git -C stack-demo/storage grep -n -i custody 9f8cb532` → no match). **M903 was never executed and is now explicitly superseded.** Its single occurrence in the repo is in a different file, `storage/terraform/storage.tf:22-25`: *"An earlier plan, M903, instead proposed `moved`-ing these resources into a local module in the infrastructure repo. That was never executed, and the shipped design supersedes it: the assets stay here and the module stays declared."* There are **no `moved` blocks** anywhere in the repo. The plan text itself lives in a third repo — `app/knowledge/plan/releases/09.00-support-in-app/m903-s3-custody/` @ app `2035f9a` — and self-declares `status: planned` in its `overview.md` front-matter, and *"Planned, not started — no ter
```

**CITED CONTENT**

```
     6      }
     7    }
     8  }
     9  
    10  module "storage" {
    11    source = "github.com/anthropos-work/infrastructure.git//modules/services/base_internal_service?ref=main"
    12  
    13    use_fargate = false
    14  
```

## 10-019
- **id**: `B10-019`
- **corpus site**: `corpus/architecture/platform-migration-status.md:92-92` (table-row)
- **citation**: `storage/terraform/storage.tf:22-25`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/storage/terraform/storage.tf`  (206 lines)

**CLAIMING UNIT**

```md
| `storage` | merged-into-app | decommissioned | no | **The v9.0 fold is finished on both sides, and the container is now gone.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service **and** the `repos.yml` clone entry in one commit, so `make init` no longer clones it and **the `storage-legacy` profile is gone** — removed with the service it selected. Asking compose for a retired token still exits 0 and starts only the always-on floor (postgresql, redis, sentinel), so the stack looks alive with the application absent: grade a documented compose command on *does it still select something*, never on *does it still parse*. This row read `mid-fold` for four M257x iterations (iter-64 → iter-68), then `merged-into-app` with a legacy rollback container for exactly one platform day. **Prod — the ECS service is DELETED, not scaled to zero.** The module is **18 lines** and declares no service at all; `storage/terraform/main.tf:9-11` says it in the platform's own words — *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS."* The module deliberately survives, because deleting the block would destroy the two buckets, the CloudFront distribution and the media DNS record along with the `prevent_destroy` guards that are read from configuration (`:13-16`). **There is no custody-transfer clause in that file and `:18` is not one** — line 18 is the module's closing comment, *"See outputs.tf — consumers should reference these by output, never by literal name."* The word *custody* occurs **0** times in the entire storage repo at `9f8cb53` (`git -C stack-demo/storage grep -n -i custody 9f8cb532` → no match). **M903 was never executed and is now explicitly superseded.** Its single occurrence in the repo is in a different file, `storage/terraform/storage.tf:22-25`: *"An earlier plan, M903, instead proposed `moved`-ing these resources into a local module in the infrastructure repo. That was never executed, and the shipped design supersedes it: the assets stay here and the module stays declared."* There are **no `moved` blocks** anywhere in the repo. The plan text itself lives in a third repo — `app/knowledge/plan/releases/09.00-support-in-app/m903-s3-custody/` @ app `2035f9a` — and self-declares `status: planned` in its `overview.md` front-matter, and *"Planned, not started — no ter
```

**CITED CONTENT**

```
    19  #
    20  # KNOWN LIMIT — `prevent_destroy` is read from CONFIGURATION, not state. It does not fire when the
    21  # `module "storage-service_euwest1"` block is deleted, which is exactly what the M907
    22  # decommission does. These guards stop accidental drift and targeted destroys; they do NOT stop
    23  # module removal. That is what M903's `moved` blocks are for — relocate the assets out of this
    24  # module BEFORE M907, and the guards become irrelevant to it rather than load-bearing.
    25  #
    26  # These guards are intentionally load-bearing: a plan that would destroy any of them fails at
    27  # plan time. If a future change genuinely needs to remove one, delete its lifecycle block in a
    28  # separate, deliberate commit that says why.
```

## 10-020
- **id**: `B10-020`
- **corpus site**: `corpus/architecture/platform-migration-status.md:92-92` (table-row)
- **citation**: `docker-compose.yml:82`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `storage` | merged-into-app | decommissioned | no | **The v9.0 fold is finished on both sides, and the container is now gone.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service **and** the `repos.yml` clone entry in one commit, so `make init` no longer clones it and **the `storage-legacy` profile is gone** — removed with the service it selected. Asking compose for a retired token still exits 0 and starts only the always-on floor (postgresql, redis, sentinel), so the stack looks alive with the application absent: grade a documented compose command on *does it still select something*, never on *does it still parse*. This row read `mid-fold` for four M257x iterations (iter-64 → iter-68), then `merged-into-app` with a legacy rollback container for exactly one platform day. **Prod — the ECS service is DELETED, not scaled to zero.** The module is **18 lines** and declares no service at all; `storage/terraform/main.tf:9-11` says it in the platform's own words — *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS."* The module deliberately survives, because deleting the block would destroy the two buckets, the CloudFront distribution and the media DNS record along with the `prevent_destroy` guards that are read from configuration (`:13-16`). **There is no custody-transfer clause in that file and `:18` is not one** — line 18 is the module's closing comment, *"See outputs.tf — consumers should reference these by output, never by literal name."* The word *custody* occurs **0** times in the entire storage repo at `9f8cb53` (`git -C stack-demo/storage grep -n -i custody 9f8cb532` → no match). **M903 was never executed and is now explicitly superseded.** Its single occurrence in the repo is in a different file, `storage/terraform/storage.tf:22-25`: *"An earlier plan, M903, instead proposed `moved`-ing these resources into a local module in the infrastructure repo. That was never executed, and the shipped design supersedes it: the assets stay here and the module stays declared."* There are **no `moved` blocks** anywhere in the repo. The plan text itself lives in a third repo — `app/knowledge/plan/releases/09.00-support-in-app/m903-s3-custody/` @ app `2035f9a` — and self-declares `status: planned` in its `overview.md` front-matter, and *"Planned, not started — no ter
```

**CITED CONTENT**

```
    79        # silently writes every upload to the container's ephemeral disk.
    80        - AWS_REGION=eu-west-1
    81        - AWS_DEFAULT_REGION=eu-west-1
    82        - STORAGE_S3_BUCKET=production-storage20240826131618541000000005
    83        - STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001
    84        # messenger-in-app (v9.0) and customerio-sync-in-app are folded into this container
    85        # too, but deliberately have NO variables here. Both reach outside the process on a
```

## 10-021
- **id**: `B10-021`
- **corpus site**: `corpus/architecture/platform-migration-status.md:92-92` (table-row)
- **citation**: `app/main.go:524`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| `storage` | merged-into-app | decommissioned | no | **The v9.0 fold is finished on both sides, and the container is now gone.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service **and** the `repos.yml` clone entry in one commit, so `make init` no longer clones it and **the `storage-legacy` profile is gone** — removed with the service it selected. Asking compose for a retired token still exits 0 and starts only the always-on floor (postgresql, redis, sentinel), so the stack looks alive with the application absent: grade a documented compose command on *does it still select something*, never on *does it still parse*. This row read `mid-fold` for four M257x iterations (iter-64 → iter-68), then `merged-into-app` with a legacy rollback container for exactly one platform day. **Prod — the ECS service is DELETED, not scaled to zero.** The module is **18 lines** and declares no service at all; `storage/terraform/main.tf:9-11` says it in the platform's own words — *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS."* The module deliberately survives, because deleting the block would destroy the two buckets, the CloudFront distribution and the media DNS record along with the `prevent_destroy` guards that are read from configuration (`:13-16`). **There is no custody-transfer clause in that file and `:18` is not one** — line 18 is the module's closing comment, *"See outputs.tf — consumers should reference these by output, never by literal name."* The word *custody* occurs **0** times in the entire storage repo at `9f8cb53` (`git -C stack-demo/storage grep -n -i custody 9f8cb532` → no match). **M903 was never executed and is now explicitly superseded.** Its single occurrence in the repo is in a different file, `storage/terraform/storage.tf:22-25`: *"An earlier plan, M903, instead proposed `moved`-ing these resources into a local module in the infrastructure repo. That was never executed, and the shipped design supersedes it: the assets stay here and the module stays declared."* There are **no `moved` blocks** anywhere in the repo. The plan text itself lives in a third repo — `app/knowledge/plan/releases/09.00-support-in-app/m903-s3-custody/` @ app `2035f9a` — and self-declares `status: planned` in its `overview.md` front-matter, and *"Planned, not started — no ter
```

**CITED CONTENT**

```
   521  			"ephemeral disk", internalstorage.EnvBucket, internalstorage.EnvPublicBucket,
   522  			storageBucket, storagePublicBucket)
   523  	}
   524  	storageManager := internalstorage.NewManager(storageBucket)
   525  	publicStorageManager := internalstorage.NewPublicManager(storagePublicBucket)
   526  	// Prove the task role can actually reach both buckets under these names, now,
   527  	// rather than at the first user upload. See verifyBucketAccess in env_guards.go —
```

## 10-022
- **id**: `B10-022`
- **corpus site**: `corpus/architecture/platform-migration-status.md:92-92` (table-row)
- **citation**: `app/internal/storage/service.go:22`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/storage/service.go`  (252 lines)

**CLAIMING UNIT**

```md
| `storage` | merged-into-app | decommissioned | no | **The v9.0 fold is finished on both sides, and the container is now gone.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service **and** the `repos.yml` clone entry in one commit, so `make init` no longer clones it and **the `storage-legacy` profile is gone** — removed with the service it selected. Asking compose for a retired token still exits 0 and starts only the always-on floor (postgresql, redis, sentinel), so the stack looks alive with the application absent: grade a documented compose command on *does it still select something*, never on *does it still parse*. This row read `mid-fold` for four M257x iterations (iter-64 → iter-68), then `merged-into-app` with a legacy rollback container for exactly one platform day. **Prod — the ECS service is DELETED, not scaled to zero.** The module is **18 lines** and declares no service at all; `storage/terraform/main.tf:9-11` says it in the platform's own words — *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS."* The module deliberately survives, because deleting the block would destroy the two buckets, the CloudFront distribution and the media DNS record along with the `prevent_destroy` guards that are read from configuration (`:13-16`). **There is no custody-transfer clause in that file and `:18` is not one** — line 18 is the module's closing comment, *"See outputs.tf — consumers should reference these by output, never by literal name."* The word *custody* occurs **0** times in the entire storage repo at `9f8cb53` (`git -C stack-demo/storage grep -n -i custody 9f8cb532` → no match). **M903 was never executed and is now explicitly superseded.** Its single occurrence in the repo is in a different file, `storage/terraform/storage.tf:22-25`: *"An earlier plan, M903, instead proposed `moved`-ing these resources into a local module in the infrastructure repo. That was never executed, and the shipped design supersedes it: the assets stay here and the module stays declared."* There are **no `moved` blocks** anywhere in the repo. The plan text itself lives in a third repo — `app/knowledge/plan/releases/09.00-support-in-app/m903-s3-custody/` @ app `2035f9a` — and self-declares `status: planned` in its `overview.md` front-matter, and *"Planned, not started — no ter
```

**CITED CONTENT**

```
    19  const (
    20  	// EnvBucket is the private bucket. Empty means local-dev filesystem mode —
    21  	// main.go refuses to boot on an empty value outside a developer machine.
    22  	EnvBucket = "STORAGE_S3_BUCKET"
    23  	// EnvPublicBucket is the public bucket. Same rule.
    24  	EnvPublicBucket = "STORAGE_S3_PUBLIC_BUCKET"
    25  )
```

## 10-023
- **id**: `B10-023`
- **corpus site**: `corpus/architecture/platform-migration-status.md:92-92` (table-row)
- **citation**: `app/main.go:504`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| `storage` | merged-into-app | decommissioned | no | **The v9.0 fold is finished on both sides, and the container is now gone.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service **and** the `repos.yml` clone entry in one commit, so `make init` no longer clones it and **the `storage-legacy` profile is gone** — removed with the service it selected. Asking compose for a retired token still exits 0 and starts only the always-on floor (postgresql, redis, sentinel), so the stack looks alive with the application absent: grade a documented compose command on *does it still select something*, never on *does it still parse*. This row read `mid-fold` for four M257x iterations (iter-64 → iter-68), then `merged-into-app` with a legacy rollback container for exactly one platform day. **Prod — the ECS service is DELETED, not scaled to zero.** The module is **18 lines** and declares no service at all; `storage/terraform/main.tf:9-11` says it in the platform's own words — *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS."* The module deliberately survives, because deleting the block would destroy the two buckets, the CloudFront distribution and the media DNS record along with the `prevent_destroy` guards that are read from configuration (`:13-16`). **There is no custody-transfer clause in that file and `:18` is not one** — line 18 is the module's closing comment, *"See outputs.tf — consumers should reference these by output, never by literal name."* The word *custody* occurs **0** times in the entire storage repo at `9f8cb53` (`git -C stack-demo/storage grep -n -i custody 9f8cb532` → no match). **M903 was never executed and is now explicitly superseded.** Its single occurrence in the repo is in a different file, `storage/terraform/storage.tf:22-25`: *"An earlier plan, M903, instead proposed `moved`-ing these resources into a local module in the infrastructure repo. That was never executed, and the shipped design supersedes it: the assets stay here and the module stays declared."* There are **no `moved` blocks** anywhere in the repo. The plan text itself lives in a third repo — `app/knowledge/plan/releases/09.00-support-in-app/m903-s3-custody/` @ app `2035f9a` — and self-declares `status: planned` in its `overview.md` front-matter, and *"Planned, not started — no ter
```

**CITED CONTENT**

```
   501  		log.Fatalf("can't init authentication manager: %v", err)
   502  	}
   503  	// storage-in-app (v9.0 cutover): the in-process object managers. app IS storage
   504  	// now — the standalone service takes no traffic and STORAGE_RPC_ADDR is gone.
   505  	// One private manager + one public manager, constructed HERE like every other
   506  	// manager and passed explicitly to every consumer (the resource manager, the
   507  	// public clients, cmsStorage and jobsimwiring.Wire). Per-namespace clients are
```

## 10-024
- **id**: `B10-024`
- **corpus site**: `corpus/architecture/platform-migration-status.md:92-92` (table-row)
- **citation**: `app/internal/jobsimwiring/wiring.go:101`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimwiring/wiring.go`  (283 lines)

**CLAIMING UNIT**

```md
| `storage` | merged-into-app | decommissioned | no | **The v9.0 fold is finished on both sides, and the container is now gone.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service **and** the `repos.yml` clone entry in one commit, so `make init` no longer clones it and **the `storage-legacy` profile is gone** — removed with the service it selected. Asking compose for a retired token still exits 0 and starts only the always-on floor (postgresql, redis, sentinel), so the stack looks alive with the application absent: grade a documented compose command on *does it still select something*, never on *does it still parse*. This row read `mid-fold` for four M257x iterations (iter-64 → iter-68), then `merged-into-app` with a legacy rollback container for exactly one platform day. **Prod — the ECS service is DELETED, not scaled to zero.** The module is **18 lines** and declares no service at all; `storage/terraform/main.tf:9-11` says it in the platform's own words — *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS."* The module deliberately survives, because deleting the block would destroy the two buckets, the CloudFront distribution and the media DNS record along with the `prevent_destroy` guards that are read from configuration (`:13-16`). **There is no custody-transfer clause in that file and `:18` is not one** — line 18 is the module's closing comment, *"See outputs.tf — consumers should reference these by output, never by literal name."* The word *custody* occurs **0** times in the entire storage repo at `9f8cb53` (`git -C stack-demo/storage grep -n -i custody 9f8cb532` → no match). **M903 was never executed and is now explicitly superseded.** Its single occurrence in the repo is in a different file, `storage/terraform/storage.tf:22-25`: *"An earlier plan, M903, instead proposed `moved`-ing these resources into a local module in the infrastructure repo. That was never executed, and the shipped design supersedes it: the assets stay here and the module stays declared."* There are **no `moved` blocks** anywhere in the repo. The plan text itself lives in a third repo — `app/knowledge/plan/releases/09.00-support-in-app/m903-s3-custody/` @ app `2035f9a` — and self-declares `status: planned` in its `overview.md` front-matter, and *"Planned, not started — no ter
```

**CITED CONTENT**

```
    98  	// serves both app and the ported engine.
    99  	authz *authorization.SentinelManager,
   100  	// storage-in-app (v9.0 cutover): app's own in-process private object manager,
   101  	// replacing the STORAGE_RPC_ADDR edge. Threaded in rather than constructed here
   102  	// so the namespace literal below stays next to the incident comment that
   103  	// explains it.
   104  	inAppStorage *appstorage.Manager,
```

## 10-025
- **id**: `B10-025`
- **corpus site**: `corpus/architecture/platform-migration-status.md:92-92` (table-row)
- **citation**: `app/internal/storagens/callsites_test.go:189`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/storagens/callsites_test.go`  (298 lines)

**CLAIMING UNIT**

```md
| `storage` | merged-into-app | decommissioned | no | **The v9.0 fold is finished on both sides, and the container is now gone.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service **and** the `repos.yml` clone entry in one commit, so `make init` no longer clones it and **the `storage-legacy` profile is gone** — removed with the service it selected. Asking compose for a retired token still exits 0 and starts only the always-on floor (postgresql, redis, sentinel), so the stack looks alive with the application absent: grade a documented compose command on *does it still select something*, never on *does it still parse*. This row read `mid-fold` for four M257x iterations (iter-64 → iter-68), then `merged-into-app` with a legacy rollback container for exactly one platform day. **Prod — the ECS service is DELETED, not scaled to zero.** The module is **18 lines** and declares no service at all; `storage/terraform/main.tf:9-11` says it in the platform's own words — *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS."* The module deliberately survives, because deleting the block would destroy the two buckets, the CloudFront distribution and the media DNS record along with the `prevent_destroy` guards that are read from configuration (`:13-16`). **There is no custody-transfer clause in that file and `:18` is not one** — line 18 is the module's closing comment, *"See outputs.tf — consumers should reference these by output, never by literal name."* The word *custody* occurs **0** times in the entire storage repo at `9f8cb53` (`git -C stack-demo/storage grep -n -i custody 9f8cb532` → no match). **M903 was never executed and is now explicitly superseded.** Its single occurrence in the repo is in a different file, `storage/terraform/storage.tf:22-25`: *"An earlier plan, M903, instead proposed `moved`-ing these resources into a local module in the infrastructure repo. That was never executed, and the shipped design supersedes it: the assets stay here and the module stays declared."* There are **no `moved` blocks** anywhere in the repo. The plan text itself lives in a third repo — `app/knowledge/plan/releases/09.00-support-in-app/m903-s3-custody/` @ app `2035f9a` — and self-declares `status: planned` in its `overview.md` front-matter, and *"Planned, not started — no ter
```

**CITED CONTENT**

```
   186  }
   187  
   188  // TestNoRPCStorageClientsRemain is the cutover's one-way door. app owns the buckets
   189  // in process; the standalone storage service is scaled to 0 and STORAGE_RPC_ADDR is
   190  // gone from the task definition. An sdk/storage.NewClient(addr, ns) reintroduced
   191  // anywhere would construct a client pointed at nothing and fail at the first call,
   192  // in whatever code path happens to reach it first.
```

## 10-026
- **id**: `B10-026`
- **corpus site**: `corpus/architecture/platform-migration-status.md:92-92` (table-row)
- **citation**: `repos.yml:2-10`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
| `storage` | merged-into-app | decommissioned | no | **The v9.0 fold is finished on both sides, and the container is now gone.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service **and** the `repos.yml` clone entry in one commit, so `make init` no longer clones it and **the `storage-legacy` profile is gone** — removed with the service it selected. Asking compose for a retired token still exits 0 and starts only the always-on floor (postgresql, redis, sentinel), so the stack looks alive with the application absent: grade a documented compose command on *does it still select something*, never on *does it still parse*. This row read `mid-fold` for four M257x iterations (iter-64 → iter-68), then `merged-into-app` with a legacy rollback container for exactly one platform day. **Prod — the ECS service is DELETED, not scaled to zero.** The module is **18 lines** and declares no service at all; `storage/terraform/main.tf:9-11` says it in the platform's own words — *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS."* The module deliberately survives, because deleting the block would destroy the two buckets, the CloudFront distribution and the media DNS record along with the `prevent_destroy` guards that are read from configuration (`:13-16`). **There is no custody-transfer clause in that file and `:18` is not one** — line 18 is the module's closing comment, *"See outputs.tf — consumers should reference these by output, never by literal name."* The word *custody* occurs **0** times in the entire storage repo at `9f8cb53` (`git -C stack-demo/storage grep -n -i custody 9f8cb532` → no match). **M903 was never executed and is now explicitly superseded.** Its single occurrence in the repo is in a different file, `storage/terraform/storage.tf:22-25`: *"An earlier plan, M903, instead proposed `moved`-ing these resources into a local module in the infrastructure repo. That was never executed, and the shipped design supersedes it: the assets stay here and the module stays declared."* There are **no `moved` blocks** anywhere in the repo. The plan text itself lives in a third repo — `app/knowledge/plan/releases/09.00-support-in-app/m903-s3-custody/` @ app `2035f9a` — and self-declares `status: planned` in its `overview.md` front-matter, and *"Planned, not started — no ter
```

**CITED CONTENT**

```
     1  repos:
     2    # Go backend services
     3    #
     4    # `app` (deployed as `backend`) is the merged monolith: skiller, skillpath,
     5    # roadrunner, jobsimulation, cms, messenger, storage and customerio-sync are all
     6    # served in-process. Their Ent entities were re-created in the `public` schema
     7    # under app/terraform/migrations/, so `app` is the ONLY repo with migrations to
     8    # run. Those repos are frozen legacy: they own no local schema, no compose service
     9    # and no clone entry here. `make init` therefore does not clone them — clone them
    10    # by hand if you need to read the pre-merge source. None of them are deleted.
    11    #
    12    # `sentinel` is the one Go service still deployed alongside `backend`, so it is
    13    # the only other backend clone local dev needs.
```

## 10-027
- **id**: `B10-027`
- **corpus site**: `corpus/architecture/platform-migration-status.md:93-93` (table-row)
- **citation**: `messenger/terraform/main.tf:29`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/messenger/terraform/main.tf`  (112 lines)

**CLAIMING UNIT**

```md
| `messenger` | merged-into-app | decommissioned | no | **Folded in the same v9.0 program as `storage`, and its container went the same way, one day later.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service and the `repos.yml` clone entry together, and **the `messenger` profile is gone** with them — the same retired-token reading as `storage`'s row. **Prod — stopped, and here "scaled to zero" is still the right words:** `messenger/terraform/main.tf:29` `service_desired_count = 0` in a module that is otherwise intact (121 lines), the cms precedent again, with the reason in-comment (`:19-25` — `app` has taken the Redis consumer group over, and an untargeted `terraform apply` would silently undo a hand-made `aws ecs update-service`); the image and task definition stay declared as the rollback path (`:27-28`). **Do not pair this with `storage`'s prod state:** that module's service block is deleted outright, so a sentence reading *"each `service_desired_count = 0`"* is half false. **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's and is a deliberate literal (`:1416-1421`). **All of it sits behind a switch, which is new since iter-68:** `MESSENGER_ENABLED` (resolved at `app/main.go:285`, read at `:1445` and `:1552`) is **OFF when unset on a developer machine and a BOOT FAILURE when unset in a deployed one** (`app/env_guards.go:92-111`); prod sets it to `"true"` in the task definition (`app/terraform/main.tf:415-416`), and compose sets it **nowhere on purpose** — pinning it to `false` there would override `.env` and make opting in impossible (`docker-compose.yml:84-92`). **The RPC edge is gone, not re-pointed.** messenger's compose block was the only place in the platform that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` or `SKILLER_RPC_ADDR`; `d11a403` had re-pointed two of them at the monolith by hand, and `838d907` deleted the block, so **all four are now set by no compose file at all** — there are **zero `*_RPC_ADDR` variables anywhere 
```

**CITED CONTENT**

```
    26    private_subnets_ids            = var.platform_private_subnets_ids
    27    service_discovery_namespace_id = var.service_discovery_namespace_id
    28    monitoring_sns_topic_arn       = var.monitoring_sns_topic_arn
    29    container_definitions          = <<EOF
    30  [
    31    {
    32      "name": "${local.project}",
```

## 10-028
- **id**: `B10-028`
- **corpus site**: `corpus/architecture/platform-migration-status.md:93-93` (table-row)
- **citation**: `app/main.go:15`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| `messenger` | merged-into-app | decommissioned | no | **Folded in the same v9.0 program as `storage`, and its container went the same way, one day later.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service and the `repos.yml` clone entry together, and **the `messenger` profile is gone** with them — the same retired-token reading as `storage`'s row. **Prod — stopped, and here "scaled to zero" is still the right words:** `messenger/terraform/main.tf:29` `service_desired_count = 0` in a module that is otherwise intact (121 lines), the cms precedent again, with the reason in-comment (`:19-25` — `app` has taken the Redis consumer group over, and an untargeted `terraform apply` would silently undo a hand-made `aws ecs update-service`); the image and task definition stay declared as the rollback path (`:27-28`). **Do not pair this with `storage`'s prod state:** that module's service block is deleted outright, so a sentence reading *"each `service_desired_count = 0`"* is half false. **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's and is a deliberate literal (`:1416-1421`). **All of it sits behind a switch, which is new since iter-68:** `MESSENGER_ENABLED` (resolved at `app/main.go:285`, read at `:1445` and `:1552`) is **OFF when unset on a developer machine and a BOOT FAILURE when unset in a deployed one** (`app/env_guards.go:92-111`); prod sets it to `"true"` in the task definition (`app/terraform/main.tf:415-416`), and compose sets it **nowhere on purpose** — pinning it to `false` there would override `.env` and make opting in impossible (`docker-compose.yml:84-92`). **The RPC edge is gone, not re-pointed.** messenger's compose block was the only place in the platform that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` or `SKILLER_RPC_ADDR`; `d11a403` had re-pointed two of them at the monolith by hand, and `838d907` deleted the block, so **all four are now set by no compose file at all** — there are **zero `*_RPC_ADDR` variables anywhere 
```

**CITED CONTENT**

```
    12  	"syscall"
    13  	"time"
    14  
    15  	msgflow "github.com/anthropos-work/app/internal/messenger/flow"
    16  
    17  	"entgo.io/ent/dialect"
    18  	entsql "entgo.io/ent/dialect/sql"
```

## 10-029
- **id**: `B10-029`
- **corpus site**: `corpus/architecture/platform-migration-status.md:93-93` (table-row)
- **citation**: `app/main.go:285`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| `messenger` | merged-into-app | decommissioned | no | **Folded in the same v9.0 program as `storage`, and its container went the same way, one day later.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service and the `repos.yml` clone entry together, and **the `messenger` profile is gone** with them — the same retired-token reading as `storage`'s row. **Prod — stopped, and here "scaled to zero" is still the right words:** `messenger/terraform/main.tf:29` `service_desired_count = 0` in a module that is otherwise intact (121 lines), the cms precedent again, with the reason in-comment (`:19-25` — `app` has taken the Redis consumer group over, and an untargeted `terraform apply` would silently undo a hand-made `aws ecs update-service`); the image and task definition stay declared as the rollback path (`:27-28`). **Do not pair this with `storage`'s prod state:** that module's service block is deleted outright, so a sentence reading *"each `service_desired_count = 0`"* is half false. **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's and is a deliberate literal (`:1416-1421`). **All of it sits behind a switch, which is new since iter-68:** `MESSENGER_ENABLED` (resolved at `app/main.go:285`, read at `:1445` and `:1552`) is **OFF when unset on a developer machine and a BOOT FAILURE when unset in a deployed one** (`app/env_guards.go:92-111`); prod sets it to `"true"` in the task definition (`app/terraform/main.tf:415-416`), and compose sets it **nowhere on purpose** — pinning it to `false` there would override `.env` and make opting in impossible (`docker-compose.yml:84-92`). **The RPC edge is gone, not re-pointed.** messenger's compose block was the only place in the platform that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` or `SKILLER_RPC_ADDR`; `d11a403` had re-pointed two of them at the monolith by hand, and `838d907` deleted the block, so **all four are now set by no compose file at all** — there are **zero `*_RPC_ADDR` variables anywhere 
```

**CITED CONTENT**

```
   282  	// anything. Both subsystems act on the world — one sends mail, one rewrites
   283  	// marketing contacts — and neither is inferred from the environment any more: see
   284  	// resolveSubsystemSwitch for why unset is off on a laptop and fatal in production.
   285  	messengerEnabled := mustSubsystemSwitch(envMessengerEnabled)
   286  	customerIOSyncEnabled := mustSubsystemSwitch(envCustomerIOSyncEnabled)
   287  	logger.Info("subsystem switches",
   288  		envMessengerEnabled, messengerEnabled,
```

## 10-030
- **id**: `B10-030`
- **corpus site**: `corpus/architecture/platform-migration-status.md:93-93` (table-row)
- **citation**: `app/env_guards.go:92-111`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/env_guards.go`  (202 lines)

**CLAIMING UNIT**

```md
| `messenger` | merged-into-app | decommissioned | no | **Folded in the same v9.0 program as `storage`, and its container went the same way, one day later.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service and the `repos.yml` clone entry together, and **the `messenger` profile is gone** with them — the same retired-token reading as `storage`'s row. **Prod — stopped, and here "scaled to zero" is still the right words:** `messenger/terraform/main.tf:29` `service_desired_count = 0` in a module that is otherwise intact (121 lines), the cms precedent again, with the reason in-comment (`:19-25` — `app` has taken the Redis consumer group over, and an untargeted `terraform apply` would silently undo a hand-made `aws ecs update-service`); the image and task definition stay declared as the rollback path (`:27-28`). **Do not pair this with `storage`'s prod state:** that module's service block is deleted outright, so a sentence reading *"each `service_desired_count = 0`"* is half false. **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's and is a deliberate literal (`:1416-1421`). **All of it sits behind a switch, which is new since iter-68:** `MESSENGER_ENABLED` (resolved at `app/main.go:285`, read at `:1445` and `:1552`) is **OFF when unset on a developer machine and a BOOT FAILURE when unset in a deployed one** (`app/env_guards.go:92-111`); prod sets it to `"true"` in the task definition (`app/terraform/main.tf:415-416`), and compose sets it **nowhere on purpose** — pinning it to `false` there would override `.env` and make opting in impossible (`docker-compose.yml:84-92`). **The RPC edge is gone, not re-pointed.** messenger's compose block was the only place in the platform that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` or `SKILLER_RPC_ADDR`; `d11a403` had re-pointed two of them at the monolith by hand, and `838d907` deleted the block, so **all four are now set by no compose file at all** — there are **zero `*_RPC_ADDR` variables anywhere 
```

**CITED CONTENT**

```
    89  	return on
    90  }
    91  
    92  func resolveSubsystemSwitch(key, raw string, deployed bool) (bool, error) {
    93  	switch strings.ToLower(strings.TrimSpace(raw)) {
    94  	case "true", "1", "yes", "on":
    95  		return true, nil
    96  	case "false", "0", "no", "off":
    97  		return false, nil
    98  	case "":
    99  		if deployed {
   100  			return false, fmt.Errorf("%s is not set. Deployed environments must state this "+
   101  				"explicitly (\"true\" or \"false\") — an unset switch would silently disable the "+
   102  				"subsystem, and for messenger that means every email is dropped while the service "+
   103  				"reports healthy. Set it in app/terraform/main.tf's container environment", key)
   104  		}
   105  		return false, nil
   106  	default:
   107  		return false, fmt.Errorf("%s=%q is not a boolean. Use true/1/yes/on or false/0/no/off; "+
   108  			"an unrecognised value is rejected rather than read as \"false\" so a typo can't "+
   109  			"silently disable the subsystem", key, raw)
   110  	}
   111  }
   112  
   113  // verifyBucketAccess proves, at boot, that the task role can actually reach both
   114  // buckets under the names it was given.
```

## 10-031
- **id**: `B10-031`
- **corpus site**: `corpus/architecture/platform-migration-status.md:93-93` (table-row)
- **citation**: `app/terraform/main.tf:415-416`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/main.tf`  (787 lines)

**CLAIMING UNIT**

```md
| `messenger` | merged-into-app | decommissioned | no | **Folded in the same v9.0 program as `storage`, and its container went the same way, one day later.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service and the `repos.yml` clone entry together, and **the `messenger` profile is gone** with them — the same retired-token reading as `storage`'s row. **Prod — stopped, and here "scaled to zero" is still the right words:** `messenger/terraform/main.tf:29` `service_desired_count = 0` in a module that is otherwise intact (121 lines), the cms precedent again, with the reason in-comment (`:19-25` — `app` has taken the Redis consumer group over, and an untargeted `terraform apply` would silently undo a hand-made `aws ecs update-service`); the image and task definition stay declared as the rollback path (`:27-28`). **Do not pair this with `storage`'s prod state:** that module's service block is deleted outright, so a sentence reading *"each `service_desired_count = 0`"* is half false. **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's and is a deliberate literal (`:1416-1421`). **All of it sits behind a switch, which is new since iter-68:** `MESSENGER_ENABLED` (resolved at `app/main.go:285`, read at `:1445` and `:1552`) is **OFF when unset on a developer machine and a BOOT FAILURE when unset in a deployed one** (`app/env_guards.go:92-111`); prod sets it to `"true"` in the task definition (`app/terraform/main.tf:415-416`), and compose sets it **nowhere on purpose** — pinning it to `false` there would override `.env` and make opting in impossible (`docker-compose.yml:84-92`). **The RPC edge is gone, not re-pointed.** messenger's compose block was the only place in the platform that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` or `SKILLER_RPC_ADDR`; `d11a403` had re-pointed two of them at the monolith by hand, and `838d907` deleted the block, so **all four are now set by no compose file at all** — there are **zero `*_RPC_ADDR` variables anywhere 
```

**CITED CONTENT**

```
   412          "value": "false"
   413        },
   414        {
   415          "name": "MESSENGER_ENABLED",
   416          "value": "true"
   417        },
   418        {
   419          "name": "CUSTOMERIO_SYNC_ENABLED",
```

## 10-032
- **id**: `B10-032`
- **corpus site**: `corpus/architecture/platform-migration-status.md:93-93` (table-row)
- **citation**: `docker-compose.yml:84-92`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `messenger` | merged-into-app | decommissioned | no | **Folded in the same v9.0 program as `storage`, and its container went the same way, one day later.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service and the `repos.yml` clone entry together, and **the `messenger` profile is gone** with them — the same retired-token reading as `storage`'s row. **Prod — stopped, and here "scaled to zero" is still the right words:** `messenger/terraform/main.tf:29` `service_desired_count = 0` in a module that is otherwise intact (121 lines), the cms precedent again, with the reason in-comment (`:19-25` — `app` has taken the Redis consumer group over, and an untargeted `terraform apply` would silently undo a hand-made `aws ecs update-service`); the image and task definition stay declared as the rollback path (`:27-28`). **Do not pair this with `storage`'s prod state:** that module's service block is deleted outright, so a sentence reading *"each `service_desired_count = 0`"* is half false. **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's and is a deliberate literal (`:1416-1421`). **All of it sits behind a switch, which is new since iter-68:** `MESSENGER_ENABLED` (resolved at `app/main.go:285`, read at `:1445` and `:1552`) is **OFF when unset on a developer machine and a BOOT FAILURE when unset in a deployed one** (`app/env_guards.go:92-111`); prod sets it to `"true"` in the task definition (`app/terraform/main.tf:415-416`), and compose sets it **nowhere on purpose** — pinning it to `false` there would override `.env` and make opting in impossible (`docker-compose.yml:84-92`). **The RPC edge is gone, not re-pointed.** messenger's compose block was the only place in the platform that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` or `SKILLER_RPC_ADDR`; `d11a403` had re-pointed two of them at the monolith by hand, and `838d907` deleted the block, so **all four are now set by no compose file at all** — there are **zero `*_RPC_ADDR` variables anywhere 
```

**CITED CONTENT**

```
    81        - AWS_DEFAULT_REGION=eu-west-1
    82        - STORAGE_S3_BUCKET=production-storage20240826131618541000000005
    83        - STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001
    84        # messenger-in-app (v9.0) and customerio-sync-in-app are folded into this container
    85        # too, but deliberately have NO variables here. Both reach outside the process on a
    86        # stream or a timer — they send mail and rewrite Brevo contacts — so app gates them
    87        # behind MESSENGER_ENABLED / CUSTOMERIO_SYNC_ENABLED, which default to OFF on a
    88        # developer machine (ENVIRONMENT=development is what makes unset mean off).
    89        # Pinning them to `false` here would override .env and make opting in impossible
    90        # without editing this file. To exercise either one locally, set it in .env — and
    91        # know that messenger then attaches to the LIVE Redis consumer group and
    92        # customerio-sync writes real Brevo contacts.
    93        - SUPABASE_DB_CONN=postgresql://postgres@postgresql:5432/postgres?sslmode=disable
    94        - COPILOT_DB_CONN=postgresql://postgres@postgresql:5432/postgres?sslmode=disable
    95      networks:
```

## 10-033
- **id**: `B10-033`
- **corpus site**: `corpus/architecture/platform-migration-status.md:93-93` (table-row)
- **citation**: `docker-compose.yml:48`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `messenger` | merged-into-app | decommissioned | no | **Folded in the same v9.0 program as `storage`, and its container went the same way, one day later.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service and the `repos.yml` clone entry together, and **the `messenger` profile is gone** with them — the same retired-token reading as `storage`'s row. **Prod — stopped, and here "scaled to zero" is still the right words:** `messenger/terraform/main.tf:29` `service_desired_count = 0` in a module that is otherwise intact (121 lines), the cms precedent again, with the reason in-comment (`:19-25` — `app` has taken the Redis consumer group over, and an untargeted `terraform apply` would silently undo a hand-made `aws ecs update-service`); the image and task definition stay declared as the rollback path (`:27-28`). **Do not pair this with `storage`'s prod state:** that module's service block is deleted outright, so a sentence reading *"each `service_desired_count = 0`"* is half false. **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's and is a deliberate literal (`:1416-1421`). **All of it sits behind a switch, which is new since iter-68:** `MESSENGER_ENABLED` (resolved at `app/main.go:285`, read at `:1445` and `:1552`) is **OFF when unset on a developer machine and a BOOT FAILURE when unset in a deployed one** (`app/env_guards.go:92-111`); prod sets it to `"true"` in the task definition (`app/terraform/main.tf:415-416`), and compose sets it **nowhere on purpose** — pinning it to `false` there would override `.env` and make opting in impossible (`docker-compose.yml:84-92`). **The RPC edge is gone, not re-pointed.** messenger's compose block was the only place in the platform that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` or `SKILLER_RPC_ADDR`; `d11a403` had re-pointed two of them at the monolith by hand, and `838d907` deleted the block, so **all four are now set by no compose file at all** — there are **zero `*_RPC_ADDR` variables anywhere 
```

**CITED CONTENT**

```
    45        - .env
    46      environment:
    47        - AI_USAGE_STREAM=AI
    48        - AUTHORIZATION_ADDRESS=http://sentinel:8087
    49        - AWS_CHIME_SDK_REGION=eu-central-1
    50        - CHIME_RECORDINGS_BUCKET_NAME=ant-prod-chime-demo
    51        - CMS_STREAM=cms
```

## 10-034
- **id**: `B10-034`
- **corpus site**: `corpus/architecture/platform-migration-status.md:93-93` (table-row)
- **citation**: `docker-compose.yml:57`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `messenger` | merged-into-app | decommissioned | no | **Folded in the same v9.0 program as `storage`, and its container went the same way, one day later.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service and the `repos.yml` clone entry together, and **the `messenger` profile is gone** with them — the same retired-token reading as `storage`'s row. **Prod — stopped, and here "scaled to zero" is still the right words:** `messenger/terraform/main.tf:29` `service_desired_count = 0` in a module that is otherwise intact (121 lines), the cms precedent again, with the reason in-comment (`:19-25` — `app` has taken the Redis consumer group over, and an untargeted `terraform apply` would silently undo a hand-made `aws ecs update-service`); the image and task definition stay declared as the rollback path (`:27-28`). **Do not pair this with `storage`'s prod state:** that module's service block is deleted outright, so a sentence reading *"each `service_desired_count = 0`"* is half false. **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's and is a deliberate literal (`:1416-1421`). **All of it sits behind a switch, which is new since iter-68:** `MESSENGER_ENABLED` (resolved at `app/main.go:285`, read at `:1445` and `:1552`) is **OFF when unset on a developer machine and a BOOT FAILURE when unset in a deployed one** (`app/env_guards.go:92-111`); prod sets it to `"true"` in the task definition (`app/terraform/main.tf:415-416`), and compose sets it **nowhere on purpose** — pinning it to `false` there would override `.env` and make opting in impossible (`docker-compose.yml:84-92`). **The RPC edge is gone, not re-pointed.** messenger's compose block was the only place in the platform that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` or `SKILLER_RPC_ADDR`; `d11a403` had re-pointed two of them at the monolith by hand, and `838d907` deleted the block, so **all four are now set by no compose file at all** — there are **zero `*_RPC_ADDR` variables anywhere 
```

**CITED CONTENT**

```
    54        - ELEVENLABS_EU_TEMPLATE_AGENT_ID=agent_4301k834j6pxfefbgf6bg48g8kpq
    55        - ELEVENLABS_TEMPLATE_AGENT_ID=agent_01k07b5k4ge3f9cvv30rv1d49n
    56        - ENVIRONMENT=development
    57        - GOTENBERG_URL=http://gotenberg:3200
    58        - JOBSIMULATION_STREAM=jobsimulation
    59        - JUDGE0_BASE_URL=http://52.48.139.23:2358
    60        - LIVEKIT_AWS_SDK_REGION=eu-central-1
```

## 10-035
- **id**: `B10-035`
- **corpus site**: `corpus/architecture/platform-migration-status.md:93-93` (table-row)
- **citation**: `docker-compose.yml:183`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `messenger` | merged-into-app | decommissioned | no | **Folded in the same v9.0 program as `storage`, and its container went the same way, one day later.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service and the `repos.yml` clone entry together, and **the `messenger` profile is gone** with them — the same retired-token reading as `storage`'s row. **Prod — stopped, and here "scaled to zero" is still the right words:** `messenger/terraform/main.tf:29` `service_desired_count = 0` in a module that is otherwise intact (121 lines), the cms precedent again, with the reason in-comment (`:19-25` — `app` has taken the Redis consumer group over, and an untargeted `terraform apply` would silently undo a hand-made `aws ecs update-service`); the image and task definition stay declared as the rollback path (`:27-28`). **Do not pair this with `storage`'s prod state:** that module's service block is deleted outright, so a sentence reading *"each `service_desired_count = 0`"* is half false. **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's and is a deliberate literal (`:1416-1421`). **All of it sits behind a switch, which is new since iter-68:** `MESSENGER_ENABLED` (resolved at `app/main.go:285`, read at `:1445` and `:1552`) is **OFF when unset on a developer machine and a BOOT FAILURE when unset in a deployed one** (`app/env_guards.go:92-111`); prod sets it to `"true"` in the task definition (`app/terraform/main.tf:415-416`), and compose sets it **nowhere on purpose** — pinning it to `false` there would override `.env` and make opting in impossible (`docker-compose.yml:84-92`). **The RPC edge is gone, not re-pointed.** messenger's compose block was the only place in the platform that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` or `SKILLER_RPC_ADDR`; `d11a403` had re-pointed two of them at the monolith by hand, and `838d907` deleted the block, so **all four are now set by no compose file at all** — there are **zero `*_RPC_ADDR` variables anywhere 
```

**CITED CONTENT**

```
   180        - "3200:3200"
   181      networks:
   182        - app-network
   183      profiles: [core, backend, all]
   184  
   185  networks:
   186    app-network:
```

## 10-036
- **id**: `B10-036`
- **corpus site**: `corpus/architecture/platform-migration-status.md:93-93` (table-row)
- **citation**: `app/internal/converter/gotenberg.go:31`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/converter/gotenberg.go`  (54 lines)

**CLAIMING UNIT**

```md
| `messenger` | merged-into-app | decommissioned | no | **Folded in the same v9.0 program as `storage`, and its container went the same way, one day later.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service and the `repos.yml` clone entry together, and **the `messenger` profile is gone** with them — the same retired-token reading as `storage`'s row. **Prod — stopped, and here "scaled to zero" is still the right words:** `messenger/terraform/main.tf:29` `service_desired_count = 0` in a module that is otherwise intact (121 lines), the cms precedent again, with the reason in-comment (`:19-25` — `app` has taken the Redis consumer group over, and an untargeted `terraform apply` would silently undo a hand-made `aws ecs update-service`); the image and task definition stay declared as the rollback path (`:27-28`). **Do not pair this with `storage`'s prod state:** that module's service block is deleted outright, so a sentence reading *"each `service_desired_count = 0`"* is half false. **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's and is a deliberate literal (`:1416-1421`). **All of it sits behind a switch, which is new since iter-68:** `MESSENGER_ENABLED` (resolved at `app/main.go:285`, read at `:1445` and `:1552`) is **OFF when unset on a developer machine and a BOOT FAILURE when unset in a deployed one** (`app/env_guards.go:92-111`); prod sets it to `"true"` in the task definition (`app/terraform/main.tf:415-416`), and compose sets it **nowhere on purpose** — pinning it to `false` there would override `.env` and make opting in impossible (`docker-compose.yml:84-92`). **The RPC edge is gone, not re-pointed.** messenger's compose block was the only place in the platform that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` or `SKILLER_RPC_ADDR`; `d11a403` had re-pointed two of them at the monolith by hand, and `838d907` deleted the block, so **all four are now set by no compose file at all** — there are **zero `*_RPC_ADDR` variables anywhere 
```

**CITED CONTENT**

```
    28  		return nil, fmt.Errorf("gotenberg: can't finalize multipart body: %w", err)
    29  	}
    30  
    31  	req, err := http.NewRequestWithContext(ctx, "POST", gotenbergURL+"/forms/libreoffice/convert", &body)
    32  	if err != nil {
    33  		return nil, fmt.Errorf("gotenberg: can't create request: %w", err)
    34  	}
```

## 10-037
- **id**: `B10-037`
- **corpus site**: `corpus/architecture/platform-migration-status.md:93-93` (table-row)
- **citation**: `docker-compose.yml:59`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `messenger` | merged-into-app | decommissioned | no | **Folded in the same v9.0 program as `storage`, and its container went the same way, one day later.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service and the `repos.yml` clone entry together, and **the `messenger` profile is gone** with them — the same retired-token reading as `storage`'s row. **Prod — stopped, and here "scaled to zero" is still the right words:** `messenger/terraform/main.tf:29` `service_desired_count = 0` in a module that is otherwise intact (121 lines), the cms precedent again, with the reason in-comment (`:19-25` — `app` has taken the Redis consumer group over, and an untargeted `terraform apply` would silently undo a hand-made `aws ecs update-service`); the image and task definition stay declared as the rollback path (`:27-28`). **Do not pair this with `storage`'s prod state:** that module's service block is deleted outright, so a sentence reading *"each `service_desired_count = 0`"* is half false. **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's and is a deliberate literal (`:1416-1421`). **All of it sits behind a switch, which is new since iter-68:** `MESSENGER_ENABLED` (resolved at `app/main.go:285`, read at `:1445` and `:1552`) is **OFF when unset on a developer machine and a BOOT FAILURE when unset in a deployed one** (`app/env_guards.go:92-111`); prod sets it to `"true"` in the task definition (`app/terraform/main.tf:415-416`), and compose sets it **nowhere on purpose** — pinning it to `false` there would override `.env` and make opting in impossible (`docker-compose.yml:84-92`). **The RPC edge is gone, not re-pointed.** messenger's compose block was the only place in the platform that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` or `SKILLER_RPC_ADDR`; `d11a403` had re-pointed two of them at the monolith by hand, and `838d907` deleted the block, so **all four are now set by no compose file at all** — there are **zero `*_RPC_ADDR` variables anywhere 
```

**CITED CONTENT**

```
    56        - ENVIRONMENT=development
    57        - GOTENBERG_URL=http://gotenberg:3200
    58        - JOBSIMULATION_STREAM=jobsimulation
    59        - JUDGE0_BASE_URL=http://52.48.139.23:2358
    60        - LIVEKIT_AWS_SDK_REGION=eu-central-1
    61        - LIVEKIT_HOST_URL=wss://anthropos-pbvktu3v.livekit.cloud
    62        - LIVEKIT_RECORDINGS_BUCKET_NAME=anthropos-livekit-test
```

## 10-038
- **id**: `B10-038`
- **corpus site**: `corpus/architecture/platform-migration-status.md:93-93` (table-row)
- **citation**: `messenger/cmd/root.go:120-140`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/messenger/cmd/root.go`  (202 lines)

**CLAIMING UNIT**

```md
| `messenger` | merged-into-app | decommissioned | no | **Folded in the same v9.0 program as `storage`, and its container went the same way, one day later.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service and the `repos.yml` clone entry together, and **the `messenger` profile is gone** with them — the same retired-token reading as `storage`'s row. **Prod — stopped, and here "scaled to zero" is still the right words:** `messenger/terraform/main.tf:29` `service_desired_count = 0` in a module that is otherwise intact (121 lines), the cms precedent again, with the reason in-comment (`:19-25` — `app` has taken the Redis consumer group over, and an untargeted `terraform apply` would silently undo a hand-made `aws ecs update-service`); the image and task definition stay declared as the rollback path (`:27-28`). **Do not pair this with `storage`'s prod state:** that module's service block is deleted outright, so a sentence reading *"each `service_desired_count = 0`"* is half false. **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's and is a deliberate literal (`:1416-1421`). **All of it sits behind a switch, which is new since iter-68:** `MESSENGER_ENABLED` (resolved at `app/main.go:285`, read at `:1445` and `:1552`) is **OFF when unset on a developer machine and a BOOT FAILURE when unset in a deployed one** (`app/env_guards.go:92-111`); prod sets it to `"true"` in the task definition (`app/terraform/main.tf:415-416`), and compose sets it **nowhere on purpose** — pinning it to `false` there would override `.env` and make opting in impossible (`docker-compose.yml:84-92`). **The RPC edge is gone, not re-pointed.** messenger's compose block was the only place in the platform that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` or `SKILLER_RPC_ADDR`; `d11a403` had re-pointed two of them at the monolith by hand, and `838d907` deleted the block, so **all four are now set by no compose file at all** — there are **zero `*_RPC_ADDR` variables anywhere 
```

**CITED CONTENT**

```
   117  
   118  		cmsV1Client := cmsv1connect.NewCMSServiceClient(
   119  			rpc.NewHttpClient(),
   120  			os.Getenv("CMS_RPC_ADDR"),
   121  			rpc.DefaultInterceptors,
   122  		)
   123  		usersClient := usersv1connect.NewUsersServiceClient(
   124  			rpc.NewHttpClient(),
   125  			os.Getenv("BACKEND_USERS_RPC_ADDR"),
   126  			rpc.DefaultInterceptors,
   127  		)
   128  		organizationsClient := organizationsv1connect.NewOrganizationsServiceClient(
   129  			rpc.NewHttpClient(),
   130  			os.Getenv("BACKEND_USERS_RPC_ADDR"),
   131  			rpc.DefaultInterceptors,
   132  		)
   133  		skillerClient := skillerv1connect.NewSkillerServiceClient(
   134  			rpc.NewHttpClient(),
   135  			os.Getenv("SKILLER_RPC_ADDR"),
   136  			rpc.DefaultInterceptors,
   137  		)
   138  		jobsimulationClient := jobsimulationv1connect.NewJobSimulationServiceClient(
   139  			rpc.NewHttpClient(),
   140  			os.Getenv("JOBSIMULATION_RPC_ADDR"),
   141  			rpc.DefaultInterceptors,
   142  		)
   143  
```

## 10-039
- **id**: `B10-039`
- **corpus site**: `corpus/architecture/platform-migration-status.md:93-93` (table-row)
- **citation**: `repos.yml:2-10`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
| `messenger` | merged-into-app | decommissioned | no | **Folded in the same v9.0 program as `storage`, and its container went the same way, one day later.** `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service and the `repos.yml` clone entry together, and **the `messenger` profile is gone** with them — the same retired-token reading as `storage`'s row. **Prod — stopped, and here "scaled to zero" is still the right words:** `messenger/terraform/main.tf:29` `service_desired_count = 0` in a module that is otherwise intact (121 lines), the cms precedent again, with the reason in-comment (`:19-25` — `app` has taken the Redis consumer group over, and an untargeted `terraform apply` would silently undo a hand-made `aws ecs update-service`); the image and task definition stay declared as the rollback path (`:27-28`). **Do not pair this with `storage`'s prod state:** that module's service block is deleted outright, so a sentence reading *"each `service_desired_count = 0`"* is half false. **Consumer side:** `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`) — it does not merge messenger's handlers onto its own subscribers, it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's and is a deliberate literal (`:1416-1421`). **All of it sits behind a switch, which is new since iter-68:** `MESSENGER_ENABLED` (resolved at `app/main.go:285`, read at `:1445` and `:1552`) is **OFF when unset on a developer machine and a BOOT FAILURE when unset in a deployed one** (`app/env_guards.go:92-111`); prod sets it to `"true"` in the task definition (`app/terraform/main.tf:415-416`), and compose sets it **nowhere on purpose** — pinning it to `false` there would override `.env` and make opting in impossible (`docker-compose.yml:84-92`). **The RPC edge is gone, not re-pointed.** messenger's compose block was the only place in the platform that set `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR` or `SKILLER_RPC_ADDR`; `d11a403` had re-pointed two of them at the monolith by hand, and `838d907` deleted the block, so **all four are now set by no compose file at all** — there are **zero `*_RPC_ADDR` variables anywhere 
```

**CITED CONTENT**

```
     1  repos:
     2    # Go backend services
     3    #
     4    # `app` (deployed as `backend`) is the merged monolith: skiller, skillpath,
     5    # roadrunner, jobsimulation, cms, messenger, storage and customerio-sync are all
     6    # served in-process. Their Ent entities were re-created in the `public` schema
     7    # under app/terraform/migrations/, so `app` is the ONLY repo with migrations to
     8    # run. Those repos are frozen legacy: they own no local schema, no compose service
     9    # and no clone entry here. `make init` therefore does not clone them — clone them
    10    # by hand if you need to read the pre-merge source. None of them are deleted.
    11    #
    12    # `sentinel` is the one Go service still deployed alongside `backend`, so it is
    13    # the only other backend clone local dev needs.
```

## 10-040
- **id**: `B10-040`
- **corpus site**: `corpus/architecture/platform-migration-status.md:94-94` (table-row)
- **citation**: `repos.yml:23-25`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
| `next-web-app` | external (Vercel) | live-standalone | yes | `repos.yml:23-25`; `docker-compose.yml:143` (`frontend` profile, `:168`). Points at `backend` directly since the router drop — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` resolves to `…:8082/graphql/query`, baked as a build arg (`docker-compose.yml:151`) and set again in the environment (`:160`) |
```

**CITED CONTENT**

```
    20      migrations: false
    21  
    22    # Frontend
    23    - name: next-web-app
    24      type: node-pnpm
    25      migrations: false
    26    - name: studio-desk
    27      type: node-npm
    28      migrations: false
```

## 10-041
- **id**: `B10-041`
- **corpus site**: `corpus/architecture/platform-migration-status.md:94-94` (table-row)
- **citation**: `docker-compose.yml:143`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `next-web-app` | external (Vercel) | live-standalone | yes | `repos.yml:23-25`; `docker-compose.yml:143` (`frontend` profile, `:168`). Points at `backend` directly since the router drop — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` resolves to `…:8082/graphql/query`, baked as a build arg (`docker-compose.yml:151`) and set again in the environment (`:160`) |
```

**CITED CONTENT**

```
   140          condition: service_started
   141      profiles: [studio-desk, all]
   142  
   143    next-web-app:
   144      build:
   145        context: ../next-web-app
   146        dockerfile: Dockerfile.dev
```

## 10-042
- **id**: `B10-042`
- **corpus site**: `corpus/architecture/platform-migration-status.md:94-94` (table-row)
- **citation**: `docker-compose.yml:151`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `next-web-app` | external (Vercel) | live-standalone | yes | `repos.yml:23-25`; `docker-compose.yml:143` (`frontend` profile, `:168`). Points at `backend` directly since the router drop — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` resolves to `…:8082/graphql/query`, baked as a build arg (`docker-compose.yml:151`) and set again in the environment (`:160`) |
```

**CITED CONTENT**

```
   148        # localhost; set it in platform/.env on a remote VM (e.g. PUBLIC_HOST=trillion)
   149        # so the client bundle resolves the user-visible hostname.
   150        args:
   151          NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT: http://${PUBLIC_HOST:-localhost}:8082/graphql/query
   152          NEXT_PUBLIC_BACKEND_API_URL: http://${PUBLIC_HOST:-localhost}:8082
   153          NEXT_PUBLIC_HOSTING_URL: http://${PUBLIC_HOST:-localhost}:3000
   154      ports:
```

## 10-043
- **id**: `B10-043`
- **corpus site**: `corpus/architecture/platform-migration-status.md:95-95` (table-row)
- **citation**: `repos.yml:26-28`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
| `studio-desk` | live-standalone | live-standalone | yes | `repos.yml:26-28`; `docker-compose.yml:112` (`studio-desk` profile, `:141`). Same re-point — `VITE_GRAPHQL_ENDPOINT` build arg at `:119`, environment at `:135` |
```

**CITED CONTENT**

```
    23    - name: next-web-app
    24      type: node-pnpm
    25      migrations: false
    26    - name: studio-desk
    27      type: node-npm
    28      migrations: false
    29  
```

## 10-044
- **id**: `B10-044`
- **corpus site**: `corpus/architecture/platform-migration-status.md:95-95` (table-row)
- **citation**: `docker-compose.yml:112`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `studio-desk` | live-standalone | live-standalone | yes | `repos.yml:26-28`; `docker-compose.yml:112` (`studio-desk` profile, `:141`). Same re-point — `VITE_GRAPHQL_ENDPOINT` build arg at `:119`, environment at `:135` |
```

**CITED CONTENT**

```
   109          condition: service_started
   110      profiles: [core, backend, all]
   111  
   112    studio-desk:
   113      build:
   114        context: ../studio-desk
   115        dockerfile: Dockerfile.dev
```
