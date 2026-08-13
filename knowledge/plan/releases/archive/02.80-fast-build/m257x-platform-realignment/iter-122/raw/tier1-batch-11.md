# TIER-1 ADJUDICATION BATCH 11 — 44 pairs

Each item: the CLAIMING UNIT (corpus prose) and the CITED CONTENT it points at (source lines,
line-numbered, +-3 lines of context). Answer ONE question per item.

## 11-001
- **id**: `B11-001`
- **corpus site**: `corpus/architecture/platform-migration-status.md:96-96` (table-row)
- **citation**: `graphql-wundergraph/terraform/main.tf:20`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/graphql-wundergraph/terraform/main.tf`  (63 lines)

**CLAIMING UNIT**

```md
| `graphql-wundergraph` | live-standalone | decommissioned | no | **the router, dropped from local dev mid-milestone.** Deleted from `repos.yml` **and** compose by `b56d731` + `360efd4`, merged `2adcf71` (2026-07-31); local dev now points at `backend`. In prod it is still declared — `graphql-wundergraph/terraform/main.tf:20` `= 1` — while the **repo is ARCHIVED on GitHub 2026-07-30**. Supergraph is **one** subgraph: `supergraph-config-prod.yaml` lists `backend` alone, `schemas/` holds `backend.graphqls` alone, `subgraphs.conf` = `BACKEND=v1.360.0` (folded by `915da06`, 2026-07-29) |
```

**CITED CONTENT**

```
    17    tags                           = var.tags
    18    aws_region                     = var.aws_region
    19    project                        = local.project
    20    service_desired_count          = 1
    21    service_cpu                    = local.service_cpu
    22    service_memory                 = local.service_memory
    23    service_port                   = local.port
```

## 11-002
- **id**: `B11-002`
- **corpus site**: `corpus/architecture/platform-migration-status.md:101-101` (table-row)
- **citation**: `docker-compose.yml:154`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `customerio-sync` | merged-into-app | decommissioned | no | **A state transition this map had never recorded, and one no membership assertion could have caught** — it was never in `repos.yml`, so directions A and B are both blind to it. It was `live-standalone` on both sides until `838d907` (merged `0c91421`, 2026-08-05), which deleted the compose service — it had been built straight from a git URL rather than cloned — and with it **the `customerio-sync` profile is gone**. The commit states a hazard — *"was still in the `all` profile, so `make up-all` started a second Brevo contact pusher alongside backend's own."* — and **the second half of that sentence is false; it is quoted here as the platform's wording, not endorsed** (corrected M257x iter-102). The `all`-profile half is true (`profiles: [customerio-sync, all]`, `0dab54d:docker-compose.yml:154`); the *"second pusher"* half is not, because `backend`'s own in-process pusher is gated behind `CUSTOMERIO_SYNC_ENABLED`, unset and therefore **off** on a developer machine — so `make up-all` started exactly **one**. **This is the corpus inheriting a false claim by quoting a commit message as authoritative**, which is worth naming as a class: a platform commit message is evidence of *intent*, never a measurement. **Consumer side:** the code is `app/internal/customeriosync/`, constructed at `app/main.go:395` (`customeriosync.New`) behind `CUSTOMERIO_SYNC_ENABLED` (resolved at `:286`, read at `:394`) — same switch semantics as messenger's: off when unset on a developer machine, a boot failure when unset in a deployed one (`app/env_guards.go:92-111`), and set to `"true"` in prod's task definition (`app/terraform/main.tf:419-420`). compose sets it nowhere, deliberately, and says why in-comment (`docker-compose.yml:84-92`); `backend`'s `depends_on` block states the disappearance (`:102-103`). **The prod half is asserted from `app`'s side only:** the standalone's own terraform lives in a repo that has never been in the clone set and that this map has therefore never read — the same gap the `roadrunner` row carries, recorded rather than papered over |
```

**CITED CONTENT**

```
   151          NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT: http://${PUBLIC_HOST:-localhost}:8082/graphql/query
   152          NEXT_PUBLIC_BACKEND_API_URL: http://${PUBLIC_HOST:-localhost}:8082
   153          NEXT_PUBLIC_HOSTING_URL: http://${PUBLIC_HOST:-localhost}:3000
   154      ports:
   155        - "3000:3000"
   156      env_file:
   157        - .env
```

## 11-003
- **id**: `B11-003`
- **corpus site**: `corpus/architecture/platform-migration-status.md:101-101` (table-row)
- **citation**: `app/main.go:395`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| `customerio-sync` | merged-into-app | decommissioned | no | **A state transition this map had never recorded, and one no membership assertion could have caught** — it was never in `repos.yml`, so directions A and B are both blind to it. It was `live-standalone` on both sides until `838d907` (merged `0c91421`, 2026-08-05), which deleted the compose service — it had been built straight from a git URL rather than cloned — and with it **the `customerio-sync` profile is gone**. The commit states a hazard — *"was still in the `all` profile, so `make up-all` started a second Brevo contact pusher alongside backend's own."* — and **the second half of that sentence is false; it is quoted here as the platform's wording, not endorsed** (corrected M257x iter-102). The `all`-profile half is true (`profiles: [customerio-sync, all]`, `0dab54d:docker-compose.yml:154`); the *"second pusher"* half is not, because `backend`'s own in-process pusher is gated behind `CUSTOMERIO_SYNC_ENABLED`, unset and therefore **off** on a developer machine — so `make up-all` started exactly **one**. **This is the corpus inheriting a false claim by quoting a commit message as authoritative**, which is worth naming as a class: a platform commit message is evidence of *intent*, never a measurement. **Consumer side:** the code is `app/internal/customeriosync/`, constructed at `app/main.go:395` (`customeriosync.New`) behind `CUSTOMERIO_SYNC_ENABLED` (resolved at `:286`, read at `:394`) — same switch semantics as messenger's: off when unset on a developer machine, a boot failure when unset in a deployed one (`app/env_guards.go:92-111`), and set to `"true"` in prod's task definition (`app/terraform/main.tf:419-420`). compose sets it nowhere, deliberately, and says why in-comment (`docker-compose.yml:84-92`); `backend`'s `depends_on` block states the disappearance (`:102-103`). **The prod half is asserted from `app`'s side only:** the standalone's own terraform lives in a repo that has never been in the clone set and that this map has therefore never read — the same gap the `roadrunner` row carries, recorded rather than papered over |
```

**CITED CONTENT**

```
   392  	// count(*)) are recorded in customeriosync/store.go.
   393  	var customerIOSyncManager *customeriosync.Manager
   394  	if customerIOSyncEnabled {
   395  		customerIOSyncManager = customeriosync.New(logger, copilotDB, os.Getenv("BREVO_KEY"))
   396  	}
   397  	// ent here is the primary-DB ORM client (public schema). The AI Readiness
   398  	// cycles/snapshots/narratives tables it owns live there; the analytics reads
```

## 11-004
- **id**: `B11-004`
- **corpus site**: `corpus/architecture/platform-migration-status.md:101-101` (table-row)
- **citation**: `app/env_guards.go:92-111`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/env_guards.go`  (202 lines)

**CLAIMING UNIT**

```md
| `customerio-sync` | merged-into-app | decommissioned | no | **A state transition this map had never recorded, and one no membership assertion could have caught** — it was never in `repos.yml`, so directions A and B are both blind to it. It was `live-standalone` on both sides until `838d907` (merged `0c91421`, 2026-08-05), which deleted the compose service — it had been built straight from a git URL rather than cloned — and with it **the `customerio-sync` profile is gone**. The commit states a hazard — *"was still in the `all` profile, so `make up-all` started a second Brevo contact pusher alongside backend's own."* — and **the second half of that sentence is false; it is quoted here as the platform's wording, not endorsed** (corrected M257x iter-102). The `all`-profile half is true (`profiles: [customerio-sync, all]`, `0dab54d:docker-compose.yml:154`); the *"second pusher"* half is not, because `backend`'s own in-process pusher is gated behind `CUSTOMERIO_SYNC_ENABLED`, unset and therefore **off** on a developer machine — so `make up-all` started exactly **one**. **This is the corpus inheriting a false claim by quoting a commit message as authoritative**, which is worth naming as a class: a platform commit message is evidence of *intent*, never a measurement. **Consumer side:** the code is `app/internal/customeriosync/`, constructed at `app/main.go:395` (`customeriosync.New`) behind `CUSTOMERIO_SYNC_ENABLED` (resolved at `:286`, read at `:394`) — same switch semantics as messenger's: off when unset on a developer machine, a boot failure when unset in a deployed one (`app/env_guards.go:92-111`), and set to `"true"` in prod's task definition (`app/terraform/main.tf:419-420`). compose sets it nowhere, deliberately, and says why in-comment (`docker-compose.yml:84-92`); `backend`'s `depends_on` block states the disappearance (`:102-103`). **The prod half is asserted from `app`'s side only:** the standalone's own terraform lives in a repo that has never been in the clone set and that this map has therefore never read — the same gap the `roadrunner` row carries, recorded rather than papered over |
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

## 11-005
- **id**: `B11-005`
- **corpus site**: `corpus/architecture/platform-migration-status.md:101-101` (table-row)
- **citation**: `app/terraform/main.tf:419-420`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/main.tf`  (787 lines)

**CLAIMING UNIT**

```md
| `customerio-sync` | merged-into-app | decommissioned | no | **A state transition this map had never recorded, and one no membership assertion could have caught** — it was never in `repos.yml`, so directions A and B are both blind to it. It was `live-standalone` on both sides until `838d907` (merged `0c91421`, 2026-08-05), which deleted the compose service — it had been built straight from a git URL rather than cloned — and with it **the `customerio-sync` profile is gone**. The commit states a hazard — *"was still in the `all` profile, so `make up-all` started a second Brevo contact pusher alongside backend's own."* — and **the second half of that sentence is false; it is quoted here as the platform's wording, not endorsed** (corrected M257x iter-102). The `all`-profile half is true (`profiles: [customerio-sync, all]`, `0dab54d:docker-compose.yml:154`); the *"second pusher"* half is not, because `backend`'s own in-process pusher is gated behind `CUSTOMERIO_SYNC_ENABLED`, unset and therefore **off** on a developer machine — so `make up-all` started exactly **one**. **This is the corpus inheriting a false claim by quoting a commit message as authoritative**, which is worth naming as a class: a platform commit message is evidence of *intent*, never a measurement. **Consumer side:** the code is `app/internal/customeriosync/`, constructed at `app/main.go:395` (`customeriosync.New`) behind `CUSTOMERIO_SYNC_ENABLED` (resolved at `:286`, read at `:394`) — same switch semantics as messenger's: off when unset on a developer machine, a boot failure when unset in a deployed one (`app/env_guards.go:92-111`), and set to `"true"` in prod's task definition (`app/terraform/main.tf:419-420`). compose sets it nowhere, deliberately, and says why in-comment (`docker-compose.yml:84-92`); `backend`'s `depends_on` block states the disappearance (`:102-103`). **The prod half is asserted from `app`'s side only:** the standalone's own terraform lives in a repo that has never been in the clone set and that this map has therefore never read — the same gap the `roadrunner` row carries, recorded rather than papered over |
```

**CITED CONTENT**

```
   416          "value": "true"
   417        },
   418        {
   419          "name": "CUSTOMERIO_SYNC_ENABLED",
   420          "value": "true"
   421        },
   422        {
   423          "name": "SKILLPATH_STREAM",
```

## 11-006
- **id**: `B11-006`
- **corpus site**: `corpus/architecture/platform-migration-status.md:101-101` (table-row)
- **citation**: `docker-compose.yml:84-92`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `customerio-sync` | merged-into-app | decommissioned | no | **A state transition this map had never recorded, and one no membership assertion could have caught** — it was never in `repos.yml`, so directions A and B are both blind to it. It was `live-standalone` on both sides until `838d907` (merged `0c91421`, 2026-08-05), which deleted the compose service — it had been built straight from a git URL rather than cloned — and with it **the `customerio-sync` profile is gone**. The commit states a hazard — *"was still in the `all` profile, so `make up-all` started a second Brevo contact pusher alongside backend's own."* — and **the second half of that sentence is false; it is quoted here as the platform's wording, not endorsed** (corrected M257x iter-102). The `all`-profile half is true (`profiles: [customerio-sync, all]`, `0dab54d:docker-compose.yml:154`); the *"second pusher"* half is not, because `backend`'s own in-process pusher is gated behind `CUSTOMERIO_SYNC_ENABLED`, unset and therefore **off** on a developer machine — so `make up-all` started exactly **one**. **This is the corpus inheriting a false claim by quoting a commit message as authoritative**, which is worth naming as a class: a platform commit message is evidence of *intent*, never a measurement. **Consumer side:** the code is `app/internal/customeriosync/`, constructed at `app/main.go:395` (`customeriosync.New`) behind `CUSTOMERIO_SYNC_ENABLED` (resolved at `:286`, read at `:394`) — same switch semantics as messenger's: off when unset on a developer machine, a boot failure when unset in a deployed one (`app/env_guards.go:92-111`), and set to `"true"` in prod's task definition (`app/terraform/main.tf:419-420`). compose sets it nowhere, deliberately, and says why in-comment (`docker-compose.yml:84-92`); `backend`'s `depends_on` block states the disappearance (`:102-103`). **The prod half is asserted from `app`'s side only:** the standalone's own terraform lives in a repo that has never been in the clone set and that this map has therefore never read — the same gap the `roadrunner` row carries, recorded rather than papered over |
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

## 11-007
- **id**: `B11-007`
- **corpus site**: `corpus/architecture/platform-migration-status.md:105-105` (table-row)
- **citation**: `docker-compose.yml:170-171`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `gotenberg` | external | live-standalone | no | third-party image, `docker-compose.yml:170-171` (`gotenberg/gotenberg:8`), default `core` profile (`:183` — renamed from `graphql` by `0dab54d`, since the WunderGraph router the name described is gone) |
```

**CITED CONTENT**

```
   167          condition: service_started
   168      profiles: [frontend, all]
   169  
   170    gotenberg:
   171      image: gotenberg/gotenberg:8
   172      command:
   173        [
   174          "gotenberg",
```

## 11-008
- **id**: `B11-008`
- **corpus site**: `corpus/architecture/platform-migration-status.md:111-111` (table-row)
- **citation**: `common.yml:2`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/common.yml`  (37 lines)

**CLAIMING UNIT**

```md
| `postgresql` | external | external | no | the shared database. Not in `docker-compose.yml` at all — it lives in the **included** `common.yml:2` (`docker-compose.yml:1-2`, `include: - common.yml`), which is why a top-level grep of the compose file finds no database. Its healthcheck gained a `start_period: 120s` at `6060315` (`common.yml:22`) because permission re-application on a grown data dir outlasted the 25 s the retries allowed — a **bring-up-timing** change, so any cold-cycle timing baseline taken before `ef32d4c` is measuring a different startup contract |
```

**CITED CONTENT**

```
     1  services:
     2    postgresql:
     3      build:
     4        context: postgresql
     5      ports:
```

## 11-009
- **id**: `B11-009`
- **corpus site**: `corpus/architecture/platform-migration-status.md:111-111` (table-row)
- **citation**: `docker-compose.yml:1-2`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `postgresql` | external | external | no | the shared database. Not in `docker-compose.yml` at all — it lives in the **included** `common.yml:2` (`docker-compose.yml:1-2`, `include: - common.yml`), which is why a top-level grep of the compose file finds no database. Its healthcheck gained a `start_period: 120s` at `6060315` (`common.yml:22`) because permission re-application on a grown data dir outlasted the 25 s the retries allowed — a **bring-up-timing** change, so any cold-cycle timing baseline taken before `ef32d4c` is measuring a different startup contract |
```

**CITED CONTENT**

```
     1  include:
     2    - common.yml
     3  
     4  services:
     5    sentinel:
```

## 11-010
- **id**: `B11-010`
- **corpus site**: `corpus/architecture/platform-migration-status.md:111-111` (table-row)
- **citation**: `common.yml:22`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/common.yml`  (37 lines)

**CLAIMING UNIT**

```md
| `postgresql` | external | external | no | the shared database. Not in `docker-compose.yml` at all — it lives in the **included** `common.yml:2` (`docker-compose.yml:1-2`, `include: - common.yml`), which is why a top-level grep of the compose file finds no database. Its healthcheck gained a `start_period: 120s` at `6060315` (`common.yml:22`) because permission re-application on a grown data dir outlasted the 25 s the retries allowed — a **bring-up-timing** change, so any cold-cycle timing baseline taken before `ef32d4c` is measuring a different startup contract |
```

**CITED CONTENT**

```
    19        # Bitnami postgres re-applies permissions on ./data/postgresql at every boot,
    20        # which takes well over the 25s (5s x 5) the retries alone allow once the data
    21        # dir has grown. Failures inside start_period don't count against retries.
    22        start_period: 120s
    23  
    24    redis:
    25      image: "bitnamilegacy/redis:latest"
```

## 11-011
- **id**: `B11-011`
- **corpus site**: `corpus/architecture/platform-migration-status.md:112-112` (table-row)
- **citation**: `common.yml:24`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/common.yml`  (37 lines)

**CLAIMING UNIT**

```md
| `redis` | external | external | no | `common.yml:24`. Streams transport for the Watermill pub/sub |
```

**CITED CONTENT**

```
    21        # dir has grown. Failures inside start_period don't count against retries.
    22        start_period: 120s
    23  
    24    redis:
    25      image: "bitnamilegacy/redis:latest"
    26      environment:
    27        - ALLOW_EMPTY_PASSWORD=yes
```

## 11-012
- **id**: `B11-012`
- **corpus site**: `corpus/architecture/platform-migration-status.md:257-268` (bullet)
- **citation**: `repos.yml:18-23`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
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
    24      type: node-pnpm
    25      migrations: false
    26    - name: studio-desk
```

## 11-013
- **id**: `B11-013`
- **corpus site**: `corpus/architecture/platform-migration-status.md:269-282` (bullet)
- **citation**: `cms/terraform/main.tf:39`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/terraform/main.tf`  (192 lines)

**CLAIMING UNIT**

```md
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

## 11-014
- **id**: `B11-014`
- **corpus site**: `corpus/architecture/platform-migration-status.md:304-309` (paragraph)
- **citation**: `migrate-demo.sh:81-85`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/demo-stack/migrate-demo.sh`  (212 lines)

**CLAIMING UNIT**

```md
   > **This row asserted the opposite for one commit.** M257x iter-54 first wrote it up as *"the armed
   > failure is now armed"*, citing `migrate-demo.sh:81-85` / `:106` — line anchors and a code shape that
   > iter-02 had already deleted. The claim was quoted forward from iter-01 without re-measuring against
   > this milestone's own repair. It is the milestone's founding class, committed into the map built to stop
   > it, and the membership fence in §4 cannot see it: the fence checks who is in `repos.yml`, not whether
   > the prose about our own tooling is still true. Corrected the same day; recorded rather than erased.
```

**CITED CONTENT**

```
    78  # proceeds (the downstream ON_ERROR_STOP=0 + idempotent CREATE+INSERT still recover), so this only ever
    79  # REMOVES flakiness, never adds a hard failure. `wait_pg` runs under a subshell that disarms `set -e`
    80  # for the polled command so a not-yet-ready exec can never abort the script.
    81  wait_pg() { # bounded poll: postgres accepts connections (pg_isready, falling back to SELECT 1)
    82    local tries="${MIGRATE_PG_TRIES:-30}" i=0
    83    while [ "$i" -lt "$tries" ]; do
    84      if docker exec "$PGC" pg_isready -U postgres -d postgres >/dev/null 2>&1; then return 0; fi
    85      # pg_isready may be absent in a minimal image — fall back to a trivial query.
    86      if docker exec "$PGC" psql -U postgres -d postgres -tAc 'SELECT 1' >/dev/null 2>&1; then return 0; fi
    87      i=$((i+1)); sleep 1
    88    done
```

## 11-015
- **id**: `B11-015`
- **corpus site**: `corpus/architecture/security_compliance.md:22-34` (bullet)
- **citation**: `terraform/main.tf:31`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/main.tf`  (787 lines)

**CLAIMING UNIT**

```md
- **Public subnets**: Application Load Balancer (ALB). ⚠️ **The Cosmo Router was listed here until M257x
  iter-115 and the only readable evidence contradicts it.** Re-derived across **all eight** service terraform
  trees in the clone set (`app`, `sentinel`, `graphql-wundergraph`, `messenger`, `cms`, `roadrunner`,
  `storage`, `jobsimulation`): the token `public_subnet` occurs **0 times**, and **every one of the eight**
  passes `private_subnets_ids = var.platform_private_subnets_ids` — the router at
  `graphql-wundergraph@60c229f3:terraform/main.tf:31`, with **no public-subnet argument of any kind**. The
  router uses the same `base_service` module as `app` (`:11`) and `app` passes the same private ids, so these
  two bullets singled the router out for a placement it shares with `backend`, which the next bullet files as
  private. **Residual, stated rather than hidden:** `infrastructure` (which *defines* `base_service`) has never
  been in a clone set and `use_fargate = false` (`:13`) puts tasks on cluster instances this corpus cannot see —
  so the module could in principle place them elsewhere. What is measurable says private; what was published
  said public, with no ref. **Also note the router is gone from local dev entirely** (platform `2adcf71`) and
  the repo is archived on GitHub
```

**CITED CONTENT**

```
    28    dev_url          = var.atlas_dev_url
    29  }
    30  
    31  //TODO: add outputs from atlas_migration
    32  
    33  // ---------------------------------------------------------------------------
    34  // sentinel-in-app v10.0 / M1001 — the SECOND Atlas pipeline: the `sentinel`
```

## 11-016
- **id**: `B11-016`
- **corpus site**: `corpus/architecture/security_compliance.md:153-168` (paragraph)
- **citation**: `clerk-integration.md:40`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/corpus/services/clerk-integration.md`  (185 lines)

**CLAIMING UNIT**

```md
> ⚠️ **CORRECTED M257x iter-120 — this layer said Sentinel *"validates **every** API request"* and that
> *"authorization checks happen before **any** data access."* Both are false, and they are false in the
> direction that makes the platform sound MORE protected than it is.** The same class as the
> `clerk-integration.md:40` *"only"* and the `cms.md` `bash -c` inversion: an absolute quantifier over a
> security surface, published unhedged. Layer 1 directly above has been re-measured four times down to
> an exact schema count and carries its caveat; **Layer 2 sat underneath it, unhedged and unfenced,
> asserting a blanket that does not exist.**
>
> **The platform's own source says so, in a comment written as a post-mortem of this exact misreading**
> (`app/internal/web/backend/graphql/graph/resolver_skiller_taxonomy_authz.go:53-66` @ `app` `ad9f3c49`):
> the M207 skiller-in-app port dropped skiller's per-resolver guards and *"leaned on app's blanket
> `AuthorizationMiddleware` — but that gate is keyed on a `userId` operation variable and **FAILS OPEN**
> for taxonomy operations … That left every taxonomy read/write reachable by any authenticated caller
> (**cross-tenant IDOR + privilege escalation**)."* It ends: ***"Do NOT rely on the blanket gate for this
> surface — it fails open here."*** (That specific hole was closed by restoring per-resolver checks; the
> **general** statement about the gate is what this doc got wrong.)
```

**CITED CONTENT**

```
    37  - **Webhooks** (svix, 12 event types) — Clerk → Postgres + Sentinel sync.
    38  - **Backend API** — org/membership/invitation CRUD, user-create CLI, `external_id` + metadata write-back, lookups.
    39  - **User/org metadata** — `unsafeMetadata` (trial/stripe flags), `publicMetadata` (eid, isHiring, role).
    40  - **Sign-in tokens** — **five** live minting sites, one product and four harnesses. The product one is app-native admin impersonation (chosen over Enterprise-tier Actor Tokens). **See the enumeration below — this bullet used to say "only", and it was false.**
    41  - **Localization** — 8 locales (`@clerk/localizations`).
    42  
    43  **Not used** (available but untouched)
```

## 11-017
- **id**: `B11-017`
- **corpus site**: `corpus/architecture/security_compliance.md:153-168` (paragraph)
- **citation**: `app/internal/web/backend/graphql/graph/resolver_skiller_taxonomy_authz.go:53-66`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/web/backend/graphql/graph/resolver_skiller_taxonomy_authz.go`  (144 lines)

**CLAIMING UNIT**

```md
> ⚠️ **CORRECTED M257x iter-120 — this layer said Sentinel *"validates **every** API request"* and that
> *"authorization checks happen before **any** data access."* Both are false, and they are false in the
> direction that makes the platform sound MORE protected than it is.** The same class as the
> `clerk-integration.md:40` *"only"* and the `cms.md` `bash -c` inversion: an absolute quantifier over a
> security surface, published unhedged. Layer 1 directly above has been re-measured four times down to
> an exact schema count and carries its caveat; **Layer 2 sat underneath it, unhedged and unfenced,
> asserting a blanket that does not exist.**
>
> **The platform's own source says so, in a comment written as a post-mortem of this exact misreading**
> (`app/internal/web/backend/graphql/graph/resolver_skiller_taxonomy_authz.go:53-66` @ `app` `ad9f3c49`):
> the M207 skiller-in-app port dropped skiller's per-resolver guards and *"leaned on app's blanket
> `AuthorizationMiddleware` — but that gate is keyed on a `userId` operation variable and **FAILS OPEN**
> for taxonomy operations … That left every taxonomy read/write reachable by any authenticated caller
> (**cross-tenant IDOR + privilege escalation**)."* It ends: ***"Do NOT rely on the blanket gate for this
> surface — it fails open here."*** (That specific hole was closed by restoring per-resolver checks; the
> **general** statement about the gate is what this doc got wrong.)
```

**CITED CONTENT**

```
    50  
    51  // Per-resolver authorization for the ported skiller taxonomy surface.
    52  //
    53  // skiller enforced authorization PER RESOLVER (skiller graph/permission.go:
    54  // checkIsAuthenticated / checkWritePermission / checkReadPermission /
    55  // checkOrganizationMatch). The skiller-in-app M207 port dropped every one of
    56  // those guards and leaned on app's blanket AuthorizationMiddleware — but that
    57  // gate is keyed on a `userId` operation variable and FAILS OPEN for taxonomy
    58  // operations (which carry {jobRoleId, organization} and no userId): an
    59  // authenticated caller with no org short-circuits to allow, and one with an org
    60  // hits errUnknownTarget → allow. That left every taxonomy read/write reachable
    61  // by any authenticated caller (cross-tenant IDOR + privilege escalation).
    62  //
    63  // These helpers restore skiller's per-resolver model using app's
    64  // AuthorizationManager (same OrgCheckFeaturePermission / UserCheckActionPermission
    65  // primitives skiller's SentinelManager exposed). Do NOT rely on the blanket gate
    66  // for this surface — it fails open here.
    67  
    68  // checkTaxonomyWritePermission mirrors skiller checkWritePermission: require the
    69  // org taxonomy-write feature on the named organization, else fall back to the
```

## 11-018
- **id**: `B11-018`
- **corpus site**: `corpus/architecture/security_compliance.md:170-173` (paragraph)
- **citation**: `app/internal/authorization/gqlauthz/gqlauthz.go:149`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/authorization/gqlauthz/gqlauthz.go`  (263 lines)

**CLAIMING UNIT**

```md
**What is actually enforced.** `AuthorizationMiddleware`'s own doc comment says it *"gates every
operation **on a viewer**"* (`app/internal/authorization/gqlauthz/gqlauthz.go:149` @ `app` `ad9f3c49`) —
an **authentication** gate. The single Sentinel call is `OrgCheckUserPermission` at `:222`, and **six
paths reach the resolver before it**:
```

**CITED CONTENT**

```
   146  	return true
   147  }
   148  
   149  // AuthorizationMiddleware gates every operation on a viewer. environment is used
   150  // only to scope the introspection exemption to local development — see
   151  // isIntrospectionOnlyQuery.
   152  func AuthorizationMiddleware(authorizationManager authorization.Manager, logger *slog.Logger, schema *ast.Schema, environment colony.Environment) graphql.ResponseMiddleware {
```

## 11-019
- **id**: `B11-019`
- **corpus site**: `corpus/architecture/security_compliance.md:202-210` (paragraph)
- **citation**: `internal/web/backend/gate.go:27-49`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/web/backend/gate.go`  (50 lines)

**CLAIMING UNIT**

```md
> ⚠️ **CORRECTED at iter-121, in the OTHER direction.** iter-120's own repair of this paragraph said
> *"every Echo group … and nothing else"*, and cited `:230-231` / `:274-275` — **each one line short of
> the third middleware**. `cbGate := courseBuilderAccessGate(authorizationManager)` (`backend.go:227`,
> defined `internal/web/backend/gate.go:27-49`) **is** a Sentinel-backed group middleware: it requires a
> user, a non-nil active org, and `OrgCheckFeaturePermission(OrgFeatureMembersEdit, orgID)`, returning
> 401/403 before the handler. Two of the six groups carry it. The conclusion — no blanket, authorization
> is opt-in — survives; the absolute quantifier did not. **Same defect class as the sentence it replaced,
> pointing the other way**, and a citation that stops one line short of its own subject is exactly the
> wrong-construct class `anchor_construct_guard` does not detect.
```

**CITED CONTENT**

```
    24  // Org/permission failures return 403 (never 401 — the caller IS
    25  // authenticated, just not authorized). A missing authn user is the only 401
    26  // path.
    27  func courseBuilderAccessGate(authz authorization.Manager) echo.MiddlewareFunc {
    28  	return func(next echo.HandlerFunc) echo.HandlerFunc {
    29  		return func(c echo.Context) error {
    30  			user := authnEcho.UserFromEchoContext(c)
    31  			if user == nil {
    32  				return echo.NewHTTPError(http.StatusUnauthorized, "missing user")
    33  			}
    34  			org := user.GetOrganization()
    35  			if org == nil || org.ID() == uuid.Nil {
    36  				return echo.NewHTTPError(http.StatusForbidden, "missing organization context")
    37  			}
    38  			orgID := org.ID()
    39  			if authz == nil {
    40  				return echo.NewHTTPError(http.StatusForbidden, "authorization unavailable")
    41  			}
    42  			ctx := authn.NewContextWithUser(c.Request().Context(), user)
    43  			if err := authz.OrgCheckFeaturePermission(ctx, permission.OrgFeatureMembersEdit, orgID); err != nil {
    44  				return echo.NewHTTPError(http.StatusForbidden, "organization admin permission required")
    45  			}
    46  			return next(c)
    47  		}
    48  	}
    49  }
    50  
```

## 11-020
- **id**: `B11-020`
- **corpus site**: `corpus/architecture/security_compliance.md:266-266` (table-row)
- **citation**: `README.md:21`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/README.md`  (175 lines)

**CLAIMING UNIT**

```md
| **AI Token Tracking** | Centralized usage, latency, and cost tracking in **`app/internal/aiusage`** — **not** the shared `ai` library, which only wraps providers (consistent with `README.md:21` + `ai_architecture.md`) |
```

**CITED CONTENT**

```
    18  
    19  ## What Is This?
    20  
    21  Project Rosetta is a **documentation repository** - not the Anthropos platform itself. It contains:
    22  
    23  - Architecture guides explaining how the platform works
    24  - Setup instructions for building a local development environment
```

## 11-021
- **id**: `B11-021`
- **corpus site**: `corpus/architecture/security_compliance.md:288-292` (bullet)
- **citation**: `app/internal/jobsimulation/ai/ai.go:263-277`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimulation/ai/ai.go`  (355 lines)

**CLAIMING UNIT**

```md
- **⚠️ "EU-first" is not "EU-only", and the US path is a FLAG, not a fallback.** `getClient` swaps
  `azureClientEu` → **`azureClientUs`** whenever the PostHog flag **`flag_use_azure_us`** is enabled
  (`app/internal/jobsimulation/ai/ai.go:263-277`). That is a deliberate switch that can route live
  simulation traffic to a US region with no error condition involved. Direct OpenAI is additionally used as
  the **retry target on HTTP 429** (`isThrottlingError`, `:129` / `:166` / `:325`)
```

**CITED CONTENT**

```
   260  	var client ai.AI
   261  
   262  	switch vendor {
   263  	case Azure:
   264  		client := a.azureClientEu
   265  		isAzureUsFlagEnabled, err := a.posthogClient.IsFeatureEnabled(
   266  			nil,
   267  			"flag_use_azure_us",
   268  			fflags.WithOnlyEvaluateLocally(true),
   269  		)
   270  		if err != nil {
   271  			a.logger.Error("can't check feature flag, using default client", "error", err)
   272  			return client, nil
   273  		}
   274  		if isAzureUsFlagEnabled {
   275  			client = a.azureClientUs
   276  		}
   277  		return client, nil
   278  	case Openai:
   279  		client = a.openaiClient
   280  	case AnthropicAws:
```

## 11-022
- **id**: `B11-022`
- **corpus site**: `corpus/architecture/security_compliance.md:293-302` (bullet)
- **citation**: `app/internal/coursebuilder/bedrock.go:109-112`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/coursebuilder/bedrock.go`  (276 lines)

**CLAIMING UNIT**

```md
- **Anthropic is reached through AWS Bedrock `eu-west-1` from the AI manager (`:85-95`) — but "Anthropic
  Direct is not used at all" is FALSE at platform HEAD.** Course Builder routes **every** model call to
  first-party `api.anthropic.com` whenever `ANTHROPIC_API_KEY` is set:
  `app/internal/coursebuilder/bedrock.go:109-112` (`newUnderlyingClient` → `NewAnthropicClientWithModel`),
  with `ModelBackendName()` (`:100`) returning `"anthropic-api"` to say so. That is a **US-terminating**
  path outside the Bedrock EU region, selected by an env var rather than a flag — so it is not covered by
  the `flag_use_azure_us` caveat below. [`external_services.md:567`](./external_services.md) carries the
  provider row and `coursebuilder.md:48` calls it *"the shipped path"*; this section said the opposite.
  Corrected M257x iter-46 — *the anchor said `:489`, which is a TypeScript codegen comment, because it was
  transcribed from a blocker ledger instead of re-derived; corrected iter-48*
```

**CITED CONTENT**

```
   106  // ANTHROPIC_API_KEY present → the first-party Anthropic API (with the
   107  // model id normalized to its bare form); absent → AWS Bedrock, the
   108  // legacy path, byte-for-byte what shipped before the switch existed.
   109  func newUnderlyingClient(ctx context.Context, modelID string) (*askengine.BedrockClient, error) {
   110  	if key := strings.TrimSpace(os.Getenv(AnthropicAPIKeyEnv)); key != "" {
   111  		return askengine.NewAnthropicClientWithModel(key, directModelID(modelID))
   112  	}
   113  	return askengine.NewBedrockClientWithModel(ctx, modelID)
   114  }
   115  
```

## 11-023
- **id**: `B11-023`
- **corpus site**: `corpus/architecture/security_compliance.md:293-302` (bullet)
- **citation**: `coursebuilder.md:48`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/corpus/services/coursebuilder.md`  (162 lines)

**CLAIMING UNIT**

```md
- **Anthropic is reached through AWS Bedrock `eu-west-1` from the AI manager (`:85-95`) — but "Anthropic
  Direct is not used at all" is FALSE at platform HEAD.** Course Builder routes **every** model call to
  first-party `api.anthropic.com` whenever `ANTHROPIC_API_KEY` is set:
  `app/internal/coursebuilder/bedrock.go:109-112` (`newUnderlyingClient` → `NewAnthropicClientWithModel`),
  with `ModelBackendName()` (`:100`) returning `"anthropic-api"` to say so. That is a **US-terminating**
  path outside the Bedrock EU region, selected by an env var rather than a flag — so it is not covered by
  the `flag_use_azure_us` caveat below. [`external_services.md:567`](./external_services.md) carries the
  provider row and `coursebuilder.md:48` calls it *"the shipped path"*; this section said the opposite.
  Corrected M257x iter-46 — *the anchor said `:489`, which is a TypeScript codegen comment, because it was
  transcribed from a blocker ledger instead of re-derived; corrected iter-48*
```

**CITED CONTENT**

```
    45      chapter shape + validators), `normalize.go` (widget normalization + XSS sanitization), `bedrock.go`/`model.go`/`usage.go`
    46      (LLM adapter + `MockClient` + cost formula), `embed.go` (`//go:embed assets/*.md` rubric), `imagegen/` (cover
    47      images).
    48  *   **LLM usage — the backend is SELECTED AT START-UP, and production is the first-party Anthropic API, not Bedrock.** `internal/coursebuilder/bedrock.go:105-114` returns an `api.anthropic.com` client with bare model ids whenever `ANTHROPIC_API_KEY` is set, reporting `ModelBackendName() == "anthropic-api"` (`:98-104`, logged at `main.go:770` @ `app` `b948604` v1.366.0); the Bedrock `eu-west-1` path via `internal/askengine/bedrock.go` is the fallback when it is not. **In production the key is required** — at `app` `ad9f3c49`, `terraform/variables.tf:759-763` declares it `sensitive` with no default, `ssm.tf:328-333` creates the SecureString parameter and `main.tf:757-758` injects it from the SSM ARN — so the shipped path is the direct API. ⚠️ **These read `variables.tf:635-638` / `main.tf:555` until M257x iter-115, and both had drifted onto DIFFERENT SUBJECTS**, which is the maximally misleading failure: at `ad9f3c49`, `variables.tf:631-645` is a cms-in-app secrets comment block and `main.tf:555` is `"name": "DIRECTUS_BASE_ADDR"`, so a reader opening the cited line saw a Directus variable and read the whole production-key claim as wrong. **The substantive claim is true and was re-derived, not assumed** — only the citation pair was false. (`ssm.tf:328` verified at both `b948604f` and `ad9f3c49`; the sentence carries no ref of its own — the only pin in the bullet is the parenthetical `@ app b948604` attached to `main.go:770`, which is a different file — so it grades at the checkout, and now names it.) Models:
    49      *   **Author/patch model**: Opus 4.8 (`eu.anthropic.claude-opus-4-8`, env `CB_AUTHOR_MODEL`; streaming, no
    50          sampling params — Opus 4.8 rejects them — at 32 K max_tokens).
    51      *   **Grader model**: Sonnet 4.6 (`eu.anthropic.claude-sonnet-4-6`, env `CB_GRADER_MODEL`; deliberately a
```

## 11-024
- **id**: `B11-024`
- **corpus site**: `corpus/architecture/security_compliance.md:312-337` (paragraph)
- **citation**: `app/internal/jobsimulation/simulator/validation/v3/validator/skills.go:53-64`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimulation/simulator/validation/v3/validator/skills.go`  (165 lines)

**CLAIMING UNIT**

```md
> **⚠️ THE STATED REASON IS FALSE AT PLATFORM HEAD, AND THIS IS A COMPLIANCE CLAIM.** It is a conjunction
> and **both conjuncts fail**. The *aggregation* is deterministic arithmetic — `calculateSkillScore`
> (`app/internal/jobsimulation/simulator/validation/v3/validator/skills.go:53-64`) counts booleans and
> `:75` divides. **The booleans it counts are LLM output.** The validator registers exactly ONE check
> engine — but **cite the DISPATCH, not that map**: `checkerEngines` is stored and never read, so it is
> not the mechanism. The real path is the hardcoded switch at
> `internal/jobsimulation/simulator/validation/basevalidator/criterion.go:127` → `validateLLM` →
> `NewLLMBulkChecker(c.logger)` (`:428`), which sends
> `basevalidator/templates/checkValidationBulk.tmpl` — a prompt asking a model to *"assess whether the
> `<asset>` … meets or does not meet"* each check and to return `{"check_id", "feedback", "success"}`. So
> "AI is used for conversation/generation only" is also false.
>
> **Not ALL verdicts are LLM-produced, and the honest claim is "most":** `EngineTextDiff` checks run
> deterministically alongside them (`criterion.go:168` dispatches `validateCodeDiff`; `:450-475` sets
> `success` from a pure string comparison, no model), and both result sets are appended together.
>
> **What follows is a question for counsel, not for this corpus**: a system that judges workers and
> candidates sits near Annex III. **Do not cite this section as evidence of a Limited-Risk
> classification** — re-derive it. Measured M257x iter-38; the same false premise was stated
> independently in `ai_architecture.md` and is corrected there too.
>
> **Both bullets above are what is STATED, not what this corpus asserts** — including the consequence
> bullet. It previously sat *after* this blockquote, at column 0, drawing the operative legal consequence
> from the classification the blockquote had just retracted three lines earlier; the retraction had been
> spliced into the middle of the list and the list resumed on the far side of it. Moved back inside the
> stated-rationale list so the retraction governs it. Repaired M257x iter-46.
```

**CITED CONTENT**

```
    50  	return float32(math.Max(0.0, float64(percentageScore*2-100)))
    51  }
    52  
    53  func (c skillsValidator) calculateSkillScore(checkResults []validation.CheckResult) (passed int, failed int) {
    54  	for _, check := range checkResults {
    55  		if check.Success {
    56  			passed += 1
    57  		} else {
    58  			failed += 1
    59  		}
    60  	}
    61  	return passed, failed
    62  }
    63  
    64  func (s skillsValidator) run(ctx context.Context, isSessionPassed bool) (validation.SkillResult, error) {
    65  	var passed, failed int
    66  	for _, r := range s.criteriaResults {
    67  		passedChecks, failedChecks := s.calculateSkillScore(r.CheckResults)
```

## 11-025
- **id**: `B11-025`
- **corpus site**: `corpus/architecture/security_compliance.md:312-337` (paragraph)
- **citation**: `internal/jobsimulation/simulator/validation/basevalidator/criterion.go:127`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimulation/simulator/validation/basevalidator/criterion.go`  (575 lines)

**CLAIMING UNIT**

```md
> **⚠️ THE STATED REASON IS FALSE AT PLATFORM HEAD, AND THIS IS A COMPLIANCE CLAIM.** It is a conjunction
> and **both conjuncts fail**. The *aggregation* is deterministic arithmetic — `calculateSkillScore`
> (`app/internal/jobsimulation/simulator/validation/v3/validator/skills.go:53-64`) counts booleans and
> `:75` divides. **The booleans it counts are LLM output.** The validator registers exactly ONE check
> engine — but **cite the DISPATCH, not that map**: `checkerEngines` is stored and never read, so it is
> not the mechanism. The real path is the hardcoded switch at
> `internal/jobsimulation/simulator/validation/basevalidator/criterion.go:127` → `validateLLM` →
> `NewLLMBulkChecker(c.logger)` (`:428`), which sends
> `basevalidator/templates/checkValidationBulk.tmpl` — a prompt asking a model to *"assess whether the
> `<asset>` … meets or does not meet"* each check and to return `{"check_id", "feedback", "success"}`. So
> "AI is used for conversation/generation only" is also false.
>
> **Not ALL verdicts are LLM-produced, and the honest claim is "most":** `EngineTextDiff` checks run
> deterministically alongside them (`criterion.go:168` dispatches `validateCodeDiff`; `:450-475` sets
> `success` from a pure string comparison, no model), and both result sets are appended together.
>
> **What follows is a question for counsel, not for this corpus**: a system that judges workers and
> candidates sits near Annex III. **Do not cite this section as evidence of a Limited-Risk
> classification** — re-derive it. Measured M257x iter-38; the same false premise was stated
> independently in `ai_architecture.md` and is corrected there too.
>
> **Both bullets above are what is STATED, not what this corpus asserts** — including the consequence
> bullet. It previously sat *after* this blockquote, at column 0, drawing the operative legal consequence
> from the classification the blockquote had just retracted three lines earlier; the retraction had been
> spliced into the middle of the list and the list resumed on the far side of it. Moved back inside the
> stated-rationale list so the retraction governs it. Repaired M257x iter-46.
```

**CITED CONTENT**

```
   124  		checksMap[critCheck.CheckID] = critCheck
   125  
   126  		switch critCheck.Engine {
   127  		case check.EngineLlm:
   128  			var p check.ParamsLLM
   129  			if err := jsoniter.Unmarshal(critCheck.Parameters, &p); err != nil {
   130  				return validation.CriterionResult{}, fmt.Errorf("can't unmarshal checkId %s llm parameters: %w", critCheck.CheckID, err)
```

## 11-026
- **id**: `B11-026`
- **corpus site**: `corpus/architecture/service_taxonomy.md:3-3` (paragraph)
- **citation**: `graphql-wundergraph/terraform/main.tf:20`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/graphql-wundergraph/terraform/main.tf`  (63 lines)

**CLAIMING UNIT**

```md
> **⚠️ Router status, two states (v2.8 M257x).** Platform `b56d731`+`360efd4` (merged **`2adcf71`**, 2026-07-31) **deleted the Cosmo Router from local dev** — no `graphql` compose service, no `repos.yml` entry — and re-pointed the frontends at **`backend` directly, `http://localhost:8082/graphql/query`**. **There is no `:5050` on a local stack.** In *production* the router is still declared (`graphql-wundergraph/terraform/main.tf:20` `= 1`), though **the repo is ARCHIVED on GitHub (2026-07-30)**. And the supergraph is **ONE** subgraph — `backend` — since `915da06` (2026-07-29). The fenced source of truth is [`platform-migration-status.md`](./platform-migration-status.md).
```

**CITED CONTENT**

```
    17    tags                           = var.tags
    18    aws_region                     = var.aws_region
    19    project                        = local.project
    20    service_desired_count          = 1
    21    service_cpu                    = local.service_cpu
    22    service_memory                 = local.service_memory
    23    service_port                   = local.port
```

## 11-027
- **id**: `B11-027`
- **corpus site**: `corpus/architecture/service_taxonomy.md:44-58` (paragraph)
- **citation**: `app/internal/cms/studio/studioManager.go:119`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/studio/studioManager.go`  (1224 lines)

**CLAIMING UNIT**

```md
> **Read the generation edge in that direction.** Until this pass the diagram drew `Room --> Desk`, which
> is backwards: Studio-Desk never receives anything from Studio-Room, and Studio-Room never calls
> Studio-Desk. Generation flows **Desk → Backend → Room** — Desk submits/polls `StudioTask` over GraphQL
> (`studio-desk/.env.example:44` bakes `VITE_GRAPHQL_ENDPOINT=http://localhost:8082/graphql/query` @ `41ee3575` — `:45` is `VITE_ENVIRONMENT=production`; the
> `studioTask` / `studioTasks` / `archiveStudioTask` operations are `app`'s — but ⚠️ **they are not all in
> one file, and this bullet supplied a single locator for three constructs until M257x iter-115.** At `app`
> `ad9f3c49`: `studioTask` is `…/graphql/graph/schemas/cms_queries.graphqls:106` and `studioTasks` is `:107`,
> while **`archiveStudioTask` is a MUTATION and lives in `cms_mutations.graphqls:22`** — it occurs nowhere in
> `cms_queries.graphqls`, so this was a wrong *file*, not line drift, and a reader chasing the archive
> operation found nothing. The load-bearing proposition — *these operations are `app`'s, not studio-desk's and
> not a standalone cms's* — is true and re-derived), and the cms domain in `app`
> then runs the pipeline as a **subprocess of its own container**, in **argv (exec) form, never through a
> shell** — `app/internal/cms/studio/studioManager.go:119` runs `studio/gen.py` via `runCommand`, whose
> contract at `:1096-1098` is *"NEVER through a shell"*. Same correction as
> [`dependency_map.md`](./dependency_map.md)'s content-generation flow, which had it right all along.
```

**CITED CONTENT**

```
   116  		}
   117  		pyBin = studioVenvPython
   118  	}
   119  	return s.runCommand(ctx, pyBin, append([]string{"studio/gen.py"}, tokens...))
   120  }
   121  
   122  // ensureStudioVenv creates the local dev virtualenv and installs the Studio Python
```

## 11-028
- **id**: `B11-028`
- **corpus site**: `corpus/architecture/service_taxonomy.md:68-68` (bullet)
- **citation**: `repos.yml:14-17`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
- **Database**: PostgreSQL — **one schema, `public`, owned by `app`**, which is the only repo with migrations (`repos.yml:14-17`). `sentinel` keeps its own `sentinel` schema (`docker-compose.yml:18`, `search_path=sentinel`) **despite `migrations: false`** (`repos.yml:18-20`) — the Trap-A case; the `cms`, `jobsimulation` and `skillpath` schemas are legacy husks
```

**CITED CONTENT**

```
    11    #
    12    # `sentinel` is the one Go service still deployed alongside `backend`, so it is
    13    # the only other backend clone local dev needs.
    14    - name: app
    15      type: go
    16      migrations: true
    17      schema: public
    18    - name: sentinel
    19      type: go
    20      migrations: false
```

## 11-029
- **id**: `B11-029`
- **corpus site**: `corpus/architecture/service_taxonomy.md:68-68` (bullet)
- **citation**: `docker-compose.yml:18`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
- **Database**: PostgreSQL — **one schema, `public`, owned by `app`**, which is the only repo with migrations (`repos.yml:14-17`). `sentinel` keeps its own `sentinel` schema (`docker-compose.yml:18`, `search_path=sentinel`) **despite `migrations: false`** (`repos.yml:18-20`) — the Trap-A case; the `cms`, `jobsimulation` and `skillpath` schemas are legacy husks
```

**CITED CONTENT**

```
    15      env_file:
    16        - .env
    17      environment:
    18        - DB_CONNECTION=postgresql://postgres@postgresql:5432/postgres?search_path=sentinel&sslmode=disable
    19        - ENVIRONMENT=development
    20        - PORT=8087
    21      networks:
```

## 11-030
- **id**: `B11-030`
- **corpus site**: `corpus/architecture/service_taxonomy.md:68-68` (bullet)
- **citation**: `repos.yml:18-20`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
- **Database**: PostgreSQL — **one schema, `public`, owned by `app`**, which is the only repo with migrations (`repos.yml:14-17`). `sentinel` keeps its own `sentinel` schema (`docker-compose.yml:18`, `search_path=sentinel`) **despite `migrations: false`** (`repos.yml:18-20`) — the Trap-A case; the `cms`, `jobsimulation` and `skillpath` schemas are legacy husks
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

## 11-031
- **id**: `B11-031`
- **corpus site**: `corpus/architecture/service_taxonomy.md:99-118` (paragraph)
- **citation**: `docker-compose.yml:84-92`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> **Storage (8300-8301), Messenger (8200-8201) and CustomerIO Sync (8080) were the other three rows.**
> Platform `838d907` (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync
> containers"*) deleted all three service definitions — build contexts, env blocks, ports, `depends_on`
> edges — and dropped `storage` + `messenger` from `repos.yml`. The `storage-legacy` / `messenger` /
> `customerio-sync` profiles are gone with them. All three are served in-process by `backend` (storage + messenger at
> v9.0 "support-in-app", customerio-sync on the asynq scheduler). The last two stay **OFF** on a
> developer machine behind `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED`, which compose deliberately
> does not set — pinning them to `false` there would override `.env` and make opting in impossible
> (`docker-compose.yml:84-92`). `customerio-sync` was still in the **`all`** profile until the deletion
> (`profiles: [customerio-sync, all]`, `0dab54d:docker-compose.yml:154`) — **that half is true; the
> "second Brevo pusher" half is not.** `make up-all` started exactly **one** Brevo contact pusher, the
> container: `backend`'s own was never on locally. Compose sets `ENVIRONMENT=development` on `backend`
> (`0dab54d:docker-compose.yml:56`, still `:56` at `0c91421`), so `deployedEnvironment()` returns
> **false** (`app/env_guards.go:37-44` @ `ad9f3c49`) and an unset `CUSTOMERIO_SYNC_ENABLED` resolves to
> `(false, nil)` rather than an error (`resolveSubsystemSwitch`, `:92-111`) — `main.go:394`'s
> `if customerIOSyncEnabled` never fires. Nor did it before that switch existed: at the fold commit
> itself, `app` `3e5bc33ef:main.go:387` gated the manager on `deployedEnvironment() &&
> os.Getenv("BREVO_KEY") != ""`. True at **every** ref between the fold (2026-08-04) and the container's
> deletion (`838d907`, 2026-08-05) — which is what the `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED`
> sentence just above already said.
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

## 11-032
- **id**: `B11-032`
- **corpus site**: `corpus/architecture/service_taxonomy.md:99-118` (paragraph)
- **citation**: `docker-compose.yml:154`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> **Storage (8300-8301), Messenger (8200-8201) and CustomerIO Sync (8080) were the other three rows.**
> Platform `838d907` (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync
> containers"*) deleted all three service definitions — build contexts, env blocks, ports, `depends_on`
> edges — and dropped `storage` + `messenger` from `repos.yml`. The `storage-legacy` / `messenger` /
> `customerio-sync` profiles are gone with them. All three are served in-process by `backend` (storage + messenger at
> v9.0 "support-in-app", customerio-sync on the asynq scheduler). The last two stay **OFF** on a
> developer machine behind `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED`, which compose deliberately
> does not set — pinning them to `false` there would override `.env` and make opting in impossible
> (`docker-compose.yml:84-92`). `customerio-sync` was still in the **`all`** profile until the deletion
> (`profiles: [customerio-sync, all]`, `0dab54d:docker-compose.yml:154`) — **that half is true; the
> "second Brevo pusher" half is not.** `make up-all` started exactly **one** Brevo contact pusher, the
> container: `backend`'s own was never on locally. Compose sets `ENVIRONMENT=development` on `backend`
> (`0dab54d:docker-compose.yml:56`, still `:56` at `0c91421`), so `deployedEnvironment()` returns
> **false** (`app/env_guards.go:37-44` @ `ad9f3c49`) and an unset `CUSTOMERIO_SYNC_ENABLED` resolves to
> `(false, nil)` rather than an error (`resolveSubsystemSwitch`, `:92-111`) — `main.go:394`'s
> `if customerIOSyncEnabled` never fires. Nor did it before that switch existed: at the fold commit
> itself, `app` `3e5bc33ef:main.go:387` gated the manager on `deployedEnvironment() &&
> os.Getenv("BREVO_KEY") != ""`. True at **every** ref between the fold (2026-08-04) and the container's
> deletion (`838d907`, 2026-08-05) — which is what the `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED`
> sentence just above already said.
```

**CITED CONTENT**

```
   151          NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT: http://${PUBLIC_HOST:-localhost}:8082/graphql/query
   152          NEXT_PUBLIC_BACKEND_API_URL: http://${PUBLIC_HOST:-localhost}:8082
   153          NEXT_PUBLIC_HOSTING_URL: http://${PUBLIC_HOST:-localhost}:3000
   154      ports:
   155        - "3000:3000"
   156      env_file:
   157        - .env
```

## 11-033
- **id**: `B11-033`
- **corpus site**: `corpus/architecture/service_taxonomy.md:99-118` (paragraph)
- **citation**: `docker-compose.yml:56`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> **Storage (8300-8301), Messenger (8200-8201) and CustomerIO Sync (8080) were the other three rows.**
> Platform `838d907` (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync
> containers"*) deleted all three service definitions — build contexts, env blocks, ports, `depends_on`
> edges — and dropped `storage` + `messenger` from `repos.yml`. The `storage-legacy` / `messenger` /
> `customerio-sync` profiles are gone with them. All three are served in-process by `backend` (storage + messenger at
> v9.0 "support-in-app", customerio-sync on the asynq scheduler). The last two stay **OFF** on a
> developer machine behind `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED`, which compose deliberately
> does not set — pinning them to `false` there would override `.env` and make opting in impossible
> (`docker-compose.yml:84-92`). `customerio-sync` was still in the **`all`** profile until the deletion
> (`profiles: [customerio-sync, all]`, `0dab54d:docker-compose.yml:154`) — **that half is true; the
> "second Brevo pusher" half is not.** `make up-all` started exactly **one** Brevo contact pusher, the
> container: `backend`'s own was never on locally. Compose sets `ENVIRONMENT=development` on `backend`
> (`0dab54d:docker-compose.yml:56`, still `:56` at `0c91421`), so `deployedEnvironment()` returns
> **false** (`app/env_guards.go:37-44` @ `ad9f3c49`) and an unset `CUSTOMERIO_SYNC_ENABLED` resolves to
> `(false, nil)` rather than an error (`resolveSubsystemSwitch`, `:92-111`) — `main.go:394`'s
> `if customerIOSyncEnabled` never fires. Nor did it before that switch existed: at the fold commit
> itself, `app` `3e5bc33ef:main.go:387` gated the manager on `deployedEnvironment() &&
> os.Getenv("BREVO_KEY") != ""`. True at **every** ref between the fold (2026-08-04) and the container's
> deletion (`838d907`, 2026-08-05) — which is what the `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED`
> sentence just above already said.
```

**CITED CONTENT**

```
    53        - DIRECTUS_PUBLIC_BASE_ADDR=https://content.anthropos.work
    54        - ELEVENLABS_EU_TEMPLATE_AGENT_ID=agent_4301k834j6pxfefbgf6bg48g8kpq
    55        - ELEVENLABS_TEMPLATE_AGENT_ID=agent_01k07b5k4ge3f9cvv30rv1d49n
    56        - ENVIRONMENT=development
    57        - GOTENBERG_URL=http://gotenberg:3200
    58        - JOBSIMULATION_STREAM=jobsimulation
    59        - JUDGE0_BASE_URL=http://52.48.139.23:2358
```

## 11-034
- **id**: `B11-034`
- **corpus site**: `corpus/architecture/service_taxonomy.md:99-118` (paragraph)
- **citation**: `app/env_guards.go:37-44`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/env_guards.go`  (202 lines)

**CLAIMING UNIT**

```md
> **Storage (8300-8301), Messenger (8200-8201) and CustomerIO Sync (8080) were the other three rows.**
> Platform `838d907` (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync
> containers"*) deleted all three service definitions — build contexts, env blocks, ports, `depends_on`
> edges — and dropped `storage` + `messenger` from `repos.yml`. The `storage-legacy` / `messenger` /
> `customerio-sync` profiles are gone with them. All three are served in-process by `backend` (storage + messenger at
> v9.0 "support-in-app", customerio-sync on the asynq scheduler). The last two stay **OFF** on a
> developer machine behind `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED`, which compose deliberately
> does not set — pinning them to `false` there would override `.env` and make opting in impossible
> (`docker-compose.yml:84-92`). `customerio-sync` was still in the **`all`** profile until the deletion
> (`profiles: [customerio-sync, all]`, `0dab54d:docker-compose.yml:154`) — **that half is true; the
> "second Brevo pusher" half is not.** `make up-all` started exactly **one** Brevo contact pusher, the
> container: `backend`'s own was never on locally. Compose sets `ENVIRONMENT=development` on `backend`
> (`0dab54d:docker-compose.yml:56`, still `:56` at `0c91421`), so `deployedEnvironment()` returns
> **false** (`app/env_guards.go:37-44` @ `ad9f3c49`) and an unset `CUSTOMERIO_SYNC_ENABLED` resolves to
> `(false, nil)` rather than an error (`resolveSubsystemSwitch`, `:92-111`) — `main.go:394`'s
> `if customerIOSyncEnabled` never fires. Nor did it before that switch existed: at the fold commit
> itself, `app` `3e5bc33ef:main.go:387` gated the manager on `deployedEnvironment() &&
> os.Getenv("BREVO_KEY") != ""`. True at **every** ref between the fold (2026-08-04) and the container's
> deletion (`838d907`, 2026-08-05) — which is what the `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED`
> sentence just above already said.
```

**CITED CONTENT**

```
    34  // "production" to Development. A future ENVIRONMENT=staging would therefore silently
    35  // disarm every guard below. This form fails safe instead — an unrecognised value is
    36  // treated as deployed.
    37  func deployedEnvironment() bool {
    38  	switch strings.ToLower(strings.TrimSpace(os.Getenv("ENVIRONMENT"))) {
    39  	case "", "development", "dev", "local", "test":
    40  		return false
    41  	default:
    42  		return true
    43  	}
    44  }
    45  
    46  // Switches for the folded subsystems that reach OUTSIDE the process on a timer or a
    47  // stream — messenger (sends mail) and customerio-sync (rewrites Brevo contacts).
```

## 11-035
- **id**: `B11-035`
- **corpus site**: `corpus/architecture/service_taxonomy.md:99-118` (paragraph)
- **citation**: `main.go:394`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> **Storage (8300-8301), Messenger (8200-8201) and CustomerIO Sync (8080) were the other three rows.**
> Platform `838d907` (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync
> containers"*) deleted all three service definitions — build contexts, env blocks, ports, `depends_on`
> edges — and dropped `storage` + `messenger` from `repos.yml`. The `storage-legacy` / `messenger` /
> `customerio-sync` profiles are gone with them. All three are served in-process by `backend` (storage + messenger at
> v9.0 "support-in-app", customerio-sync on the asynq scheduler). The last two stay **OFF** on a
> developer machine behind `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED`, which compose deliberately
> does not set — pinning them to `false` there would override `.env` and make opting in impossible
> (`docker-compose.yml:84-92`). `customerio-sync` was still in the **`all`** profile until the deletion
> (`profiles: [customerio-sync, all]`, `0dab54d:docker-compose.yml:154`) — **that half is true; the
> "second Brevo pusher" half is not.** `make up-all` started exactly **one** Brevo contact pusher, the
> container: `backend`'s own was never on locally. Compose sets `ENVIRONMENT=development` on `backend`
> (`0dab54d:docker-compose.yml:56`, still `:56` at `0c91421`), so `deployedEnvironment()` returns
> **false** (`app/env_guards.go:37-44` @ `ad9f3c49`) and an unset `CUSTOMERIO_SYNC_ENABLED` resolves to
> `(false, nil)` rather than an error (`resolveSubsystemSwitch`, `:92-111`) — `main.go:394`'s
> `if customerIOSyncEnabled` never fires. Nor did it before that switch existed: at the fold commit
> itself, `app` `3e5bc33ef:main.go:387` gated the manager on `deployedEnvironment() &&
> os.Getenv("BREVO_KEY") != ""`. True at **every** ref between the fold (2026-08-04) and the container's
> deletion (`838d907`, 2026-08-05) — which is what the `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED`
> sentence just above already said.
```

**CITED CONTENT**

```
   391  	// predicate can't be pushed down. Numbers and their trap (don't benchmark this with
   392  	// count(*)) are recorded in customeriosync/store.go.
   393  	var customerIOSyncManager *customeriosync.Manager
   394  	if customerIOSyncEnabled {
   395  		customerIOSyncManager = customeriosync.New(logger, copilotDB, os.Getenv("BREVO_KEY"))
   396  	}
   397  	// ent here is the primary-DB ORM client (public schema). The AI Readiness
```

## 11-036
- **id**: `B11-036`
- **corpus site**: `corpus/architecture/service_taxonomy.md:99-118` (paragraph)
- **citation**: `main.go:387`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> **Storage (8300-8301), Messenger (8200-8201) and CustomerIO Sync (8080) were the other three rows.**
> Platform `838d907` (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync
> containers"*) deleted all three service definitions — build contexts, env blocks, ports, `depends_on`
> edges — and dropped `storage` + `messenger` from `repos.yml`. The `storage-legacy` / `messenger` /
> `customerio-sync` profiles are gone with them. All three are served in-process by `backend` (storage + messenger at
> v9.0 "support-in-app", customerio-sync on the asynq scheduler). The last two stay **OFF** on a
> developer machine behind `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED`, which compose deliberately
> does not set — pinning them to `false` there would override `.env` and make opting in impossible
> (`docker-compose.yml:84-92`). `customerio-sync` was still in the **`all`** profile until the deletion
> (`profiles: [customerio-sync, all]`, `0dab54d:docker-compose.yml:154`) — **that half is true; the
> "second Brevo pusher" half is not.** `make up-all` started exactly **one** Brevo contact pusher, the
> container: `backend`'s own was never on locally. Compose sets `ENVIRONMENT=development` on `backend`
> (`0dab54d:docker-compose.yml:56`, still `:56` at `0c91421`), so `deployedEnvironment()` returns
> **false** (`app/env_guards.go:37-44` @ `ad9f3c49`) and an unset `CUSTOMERIO_SYNC_ENABLED` resolves to
> `(false, nil)` rather than an error (`resolveSubsystemSwitch`, `:92-111`) — `main.go:394`'s
> `if customerIOSyncEnabled` never fires. Nor did it before that switch existed: at the fold commit
> itself, `app` `3e5bc33ef:main.go:387` gated the manager on `deployedEnvironment() &&
> os.Getenv("BREVO_KEY") != ""`. True at **every** ref between the fold (2026-08-04) and the container's
> deletion (`838d907`, 2026-08-05) — which is what the `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED`
> sentence just above already said.
```

**CITED CONTENT**

```
   384  	// folded service brings its own pool" is the shape that stops fitting. This is one
   385  	// query every ten minutes; it does not deserve a standing allocation.
   386  	//
   387  	// Measured against production: ~0.9s for the every-10-minute run, ~6-8s for a full
   388  	// resync, against copilotDB's 30s statement_timeout. That is ~4x headroom, not the
   389  	// comfortable margin the query's size might suggest — and it does NOT shrink with
   390  	// the sync window, because the read model's refresh_date is computed and the
```

## 11-037
- **id**: `B11-037`
- **corpus site**: `corpus/architecture/service_taxonomy.md:130-130` (bullet)
- **citation**: `common.yml:2`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/common.yml`  (37 lines)

**CLAIMING UNIT**

```md
- **PostgreSQL** :5432 (custom image with pgvector extension) — `common.yml:2`, via `include:`
```

**CITED CONTENT**

```
     1  services:
     2    postgresql:
     3      build:
     4        context: postgresql
     5      ports:
```

## 11-038
- **id**: `B11-038`
- **corpus site**: `corpus/architecture/service_taxonomy.md:131-131` (bullet)
- **citation**: `common.yml:24`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/common.yml`  (37 lines)

**CLAIMING UNIT**

```md
- **Redis** :6379 (`bitnamilegacy/redis:latest`) — `common.yml:24`
```

**CITED CONTENT**

```
    21        # dir has grown. Failures inside start_period don't count against retries.
    22        start_period: 120s
    23  
    24    redis:
    25      image: "bitnamilegacy/redis:latest"
    26      environment:
    27        - ALLOW_EMPTY_PASSWORD=yes
```

## 11-039
- **id**: `B11-039`
- **corpus site**: `corpus/architecture/service_taxonomy.md:132-135` (bullet)
- **citation**: `docker-compose.yml:5`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
- **Sentinel** :8087 — `docker-compose.yml:5`. A Tier-1 Go service **and** a floor member. It is the
  third member the *Services* paragraph, the *Profiles* table and the Summary Table below all count
  when they say `core` starts **five** containers; this bullet list said **two** for four releases
  while every other statement of the floor in this file said three.
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

## 11-040
- **id**: `B11-040`
- **corpus site**: `corpus/architecture/service_taxonomy.md:168-168` (table-row)
- **citation**: `jobsimulation/terraform/main.tf:15-40`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/jobsimulation/terraform/main.tf`  (344 lines)

**CLAIMING UNIT**

```md
| **Jobsimulation** | Merged into Backend/App ("jobsim-in-app"); 23 run-state tables → `public`; **no subgraph**; **the prod ECS service is DESTROYED — M810 landed for this row** (`6092c6d2` deleted the `module "jobsimulation"` block; the file survives owning only the LiveKit/Chime buckets, the SSM parameters and the atlas tracker — `jobsimulation/terraform/main.tf:15-40`), unlike CMS below; **repo archive state: report both, assert neither** — this cell asserted a GitHub archive on 2026-07-31, but `origin/main` carries four commits dated **2026-08-04**, including merged PR #439, and an archived repo is read-only; archive state is not visible to this corpus (it lives in the GitHub org API, not a clone). See the fenced map | **NO — gone from compose at platform `0dab54d`** (and from `repos.yml`). Merged into `app`, no subgraph, no container | [jobsimulation.md](../services/jobsimulation.md) |
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
    26    url              = data.atlas_migration.jobsimulation_migrations.url
    27    revisions_schema = "jobsimulation"
    28    dev_url          = var.atlas_dev_url
    29  }
    30  
    31  module "jobsimulation" {
    32    source = "github.com/anthropos-work/infrastructure.git//modules/services/base_internal_service?ref=main"
    33  
    34    use_fargate = false
    35  
    36    environment                    = var.environment
    37    tags                           = var.tags
    38    aws_region                     = var.aws_region
    39    project                        = local.project
    40    service_desired_count          = 0
    41    service_cpu                    = local.service_cpu
    42    service_memory                 = local.service_memory
    43    health_check_path              = "/_meta"
```

## 11-041
- **id**: `B11-041`
- **corpus site**: `corpus/architecture/service_taxonomy.md:169-169` (table-row)
- **citation**: `cms/terraform/main.tf:39`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/terraform/main.tf`  (192 lines)

**CLAIMING UNIT**

```md
| **CMS** | Merged into Backend/App ("cms-in-app v8.0", app v1.360.0); similarity + Studio tables → `public`; supergraph **3→1** (the one commit `graphql-wundergraph@915da06` deleted `cms.graphqls` **and** `jobsimulation.graphqls`); the prod ECS module is **NOT** a settled rollback path — **report both, assert neither**: `cms/terraform/main.tf:39` still declares `module "cms"` at `service_desired_count = 0` in an intact 191-line module, *and* `6efa1d5` (merged `f38c0c4`, 2026-08-04) deleted `.github/workflows/build-production.yml` under the subject *"the cms ECR repository is decommissioned (M810)"*, its body stating that M810 deletes `module.cms_euwest1` and destroys the ECS service and the production-cms ECR repository. Whether that infrastructure-side deletion has been applied is **not visible to this corpus** — `infrastructure` has never been in any clone set; repo frozen, **not** archived (`origin/main` `f38c0c4`, 2026-08-04) | **NO — gone from compose at platform `0dab54d`** (and from `repos.yml`), deleted by `d11a403`. Merged into `app`, no subgraph, no container, no port. `d11a403` re-pointed `messenger`'s `CMS_RPC_ADDR` at `http://backend:8083` — **one of the two variables that commit moved** (with `JOBSIMULATION_RPC_ADDR`), not one of four; `838d907` then deleted `messenger` too, so **no compose file sets that variable at all** | [cms.md](../services/cms.md) |
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

## 11-042
- **id**: `B11-042`
- **corpus site**: `corpus/architecture/service_taxonomy.md:207-212` (bullet)
- **citation**: `docker-compose.yml:112`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
- **Deployment**: standalone processes — typically Vercel or local-only — **but not uniformly outside
  docker-compose**: Studio-Desk, the first Tier-2 member listed below, IS in the platform compose at
  `docker-compose.yml:112` behind `profiles: [studio-desk, all]` (`:141`), so it starts when that profile
  is selected — stacked on `core`, since `studio-desk` **alone** exits 1 (see *Profiles* below) — and not on
  a bare `make up`. The unqualified *"not in main docker-compose"* contradicted the Studio-Desk row above,
  `frontend_architecture.md:11` and `studio-desk.md:21`; corrected M257x iter-46
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

## 11-043
- **id**: `B11-043`
- **corpus site**: `corpus/architecture/service_taxonomy.md:207-212` (bullet)
- **citation**: `frontend_architecture.md:11`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/corpus/architecture/frontend_architecture.md`  (117 lines)

**CLAIMING UNIT**

```md
- **Deployment**: standalone processes — typically Vercel or local-only — **but not uniformly outside
  docker-compose**: Studio-Desk, the first Tier-2 member listed below, IS in the platform compose at
  `docker-compose.yml:112` behind `profiles: [studio-desk, all]` (`:141`), so it starts when that profile
  is selected — stacked on `core`, since `studio-desk` **alone** exits 1 (see *Profiles* below) — and not on
  a bare `make up`. The unqualified *"not in main docker-compose"* contradicted the Studio-Desk row above,
  `frontend_architecture.md:11` and `studio-desk.md:21`; corrected M257x iter-46
```

**CITED CONTENT**

```
     8  > - **[Studio-Desk](../services/studio-desk.md)** — Vite + Express, simulation design tool
     9  > - **[Ant Academy](../services/ant-academy.md)** — Next.js 16 + Expo, internal learning portal for `@anthropos.work` employees
    10  >
    11  > **Studio-Desk** is in `repos.yml` (so `make init` clones it) and *does* have a `studio-desk` compose profile (`docker-compose.yml:141`, `profiles: [studio-desk, all]` — the fact has survived platform `d11a403` **and** `838d907`; only the line number keeps moving — it was 226 before the support containers were deleted). **Ant Academy is deliberately NOT in `repos.yml`** — `make init` never clones it; a demo gets it from `ensure-clones.sh` phase d2, a dev box by hand. At `0c91421` `repos.yml` holds exactly **4** entries — `app`, `sentinel`, `next-web-app`, `studio-desk` (it was 9 at `2adcf71`, before the cms/jobsimulation/roadrunner and storage/messenger drops) — and ant-academy is not one of them. The rest of this document is about `next-web-app` specifically.
    12  
    13  ## Monorepo Structure
    14  
```

## 11-044
- **id**: `B11-044`
- **corpus site**: `corpus/architecture/service_taxonomy.md:207-212` (bullet)
- **citation**: `studio-desk.md:21`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/corpus/services/studio-desk.md`  (450 lines)

**CLAIMING UNIT**

```md
- **Deployment**: standalone processes — typically Vercel or local-only — **but not uniformly outside
  docker-compose**: Studio-Desk, the first Tier-2 member listed below, IS in the platform compose at
  `docker-compose.yml:112` behind `profiles: [studio-desk, all]` (`:141`), so it starts when that profile
  is selected — stacked on `core`, since `studio-desk` **alone** exits 1 (see *Profiles* below) — and not on
  a bare `make up`. The unqualified *"not in main docker-compose"* contradicted the Studio-Desk row above,
  `frontend_architecture.md:11` and `studio-desk.md:21`; corrected M257x iter-46
```

**CITED CONTENT**

```
    18  |:---------|:------|
    19  | **Service Type** | Custom Application (Tier 2 - Studio Services) |
    20  | **Technology Stack** | TypeScript, Vite, Express.js (vanilla TS frontend, no framework) |
    21  | **Deployment** | Runs natively for dev (`npm run dev`), or containerized via the `studio-desk` docker-compose profile (ports 9000/9100). It `depends_on` **`backend` alone** — `docker-compose.yml:138-140` @ platform `0c91421`, with `profiles: [studio-desk, all]` at `:141` (both re-anchored M257x iter-87; they were `:223-225`/`:226` at `0dab54d`, before `838d907` deleted three service blocks above them). It *also* listed **`cms`** (`:337-341` @ `2adcf71`) until that container was deleted from compose at `d11a403`; there is no `cms` service to depend on now, and it never depended on `graphql`, which is likewise no longer a compose service. Built with `VITE_GRAPHQL_ENDPOINT=http://localhost:8082/graphql/query`. **⚠️ Asking for `studio-desk` as the only profile exits 1** — the profile selects `studio-desk` but *not* the `backend` it depends on, so compose rejects the whole project (`service "studio-desk" depends on undefined service "backend": invalid compose project`). Use `PROFILE=all`, which selects both. |
    22  | **Port(s)** | 9100 (frontend), 9000 (backend) - configurable via `.env` |
    23  | **Authentication** | Clerk |
    24  | **Repository** | Local `studio-desk/` (sibling repo cloned by `make init`) |
```
