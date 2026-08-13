# TIER-1 ADJUDICATION BATCH 01 — 44 pairs

Each item: the CLAIMING UNIT (corpus prose) and the CITED CONTENT it points at (source lines,
line-numbered, +-3 lines of context). Answer ONE question per item.

## 01-001
- **id**: `B01-001`
- **corpus site**: `corpus/services/README.md:11-33` (paragraph)
- **citation**: `roadrunner/terraform/main.tf:19`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/roadrunner/terraform/main.tf`  (96 lines)

**CLAIMING UNIT**

```md
> **⚠️ `app` is the backend monolith.** **Seven** services in this index — skiller, skillpath,
> jobsimulation, cms, storage, messenger and customerio-sync — are **folded into `app`**: each domain runs
> in-process and none has a compose service or a local container. Their docs are kept for domain knowledge
> and carry a merge banner at the top. **"No longer deploy separately" is a claim about PROD, and it is not
> uniformly measurable** — for `customerio-sync` in particular the standalone's terraform lives in a repo
> that has never been in any clone set, so its prod half is asserted from `app`'s side only
> ([`platform-migration-status.md:101`](../architecture/platform-migration-status.md)). The fenced map is
> authoritative per service; this banner is about the LOCAL stack.
>
> **`roadrunner` is the eighth, and it is different: orphaned, not merged-and-undeployed.** Nothing calls it,
> but `roadrunner/terraform/main.tf:19` still reads `service_desired_count = 1` — **last changed at `84a4b4f`
> (2025-12-15), seven months before the fold, and not a decision about it** (`git blame -L 19,19`; M257x
> iter-115) — so it **does** still deploy,
> unlike cms (`cms/terraform/main.tf:39` = 0) and jobsimulation, whose ECS service **M810 has already destroyed** (`6092c6d2`; `service_desired_count` no longer appears in `jobsimulation/terraform/main.tf` at all — `:15-22`). It is the one row where prod and the platform's own
> `repos.yml` contradict each other. See [`platform-migration-status.md`](../architecture/platform-migration-status.md).
>
> And **none of them starts a container any more.** cms, jobsimulation and roadrunner did run locally as
> unfederated husks, but platform **`d11a403`** (2026-08-03) deleted all three from `docker-compose.yml`
> **and** from `repos.yml`; **`838d907`** (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and
> customerio-sync containers"*) then did the same to the last three.
> `docker-compose.yml` declares **5** services (7 effective, with `common.yml`'s `postgresql` +
> `redis`), and `repos.yml` carries **4** entries — `app`, `sentinel`, `next-web-app`, `studio-desk`.
> Read [`backend.md`](backend.md) for the current shape.
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

## 01-002
- **id**: `B01-002`
- **corpus site**: `corpus/services/README.md:11-33` (paragraph)
- **citation**: `cms/terraform/main.tf:39`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/terraform/main.tf`  (192 lines)

**CLAIMING UNIT**

```md
> **⚠️ `app` is the backend monolith.** **Seven** services in this index — skiller, skillpath,
> jobsimulation, cms, storage, messenger and customerio-sync — are **folded into `app`**: each domain runs
> in-process and none has a compose service or a local container. Their docs are kept for domain knowledge
> and carry a merge banner at the top. **"No longer deploy separately" is a claim about PROD, and it is not
> uniformly measurable** — for `customerio-sync` in particular the standalone's terraform lives in a repo
> that has never been in any clone set, so its prod half is asserted from `app`'s side only
> ([`platform-migration-status.md:101`](../architecture/platform-migration-status.md)). The fenced map is
> authoritative per service; this banner is about the LOCAL stack.
>
> **`roadrunner` is the eighth, and it is different: orphaned, not merged-and-undeployed.** Nothing calls it,
> but `roadrunner/terraform/main.tf:19` still reads `service_desired_count = 1` — **last changed at `84a4b4f`
> (2025-12-15), seven months before the fold, and not a decision about it** (`git blame -L 19,19`; M257x
> iter-115) — so it **does** still deploy,
> unlike cms (`cms/terraform/main.tf:39` = 0) and jobsimulation, whose ECS service **M810 has already destroyed** (`6092c6d2`; `service_desired_count` no longer appears in `jobsimulation/terraform/main.tf` at all — `:15-22`). It is the one row where prod and the platform's own
> `repos.yml` contradict each other. See [`platform-migration-status.md`](../architecture/platform-migration-status.md).
>
> And **none of them starts a container any more.** cms, jobsimulation and roadrunner did run locally as
> unfederated husks, but platform **`d11a403`** (2026-08-03) deleted all three from `docker-compose.yml`
> **and** from `repos.yml`; **`838d907`** (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and
> customerio-sync containers"*) then did the same to the last three.
> `docker-compose.yml` declares **5** services (7 effective, with `common.yml`'s `postgresql` +
> `redis`), and `repos.yml` carries **4** entries — `app`, `sentinel`, `next-web-app`, `studio-desk`.
> Read [`backend.md`](backend.md) for the current shape.
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

## 01-003
- **id**: `B01-003`
- **corpus site**: `corpus/services/README.md:39-39` (table-row)
- **citation**: `app/internal/jobsimwiring/wiring.go:123`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimwiring/wiring.go`  (283 lines)

**CLAIMING UNIT**

```md
| [`backend.md`](backend.md) | Backend (`app`) | **The monolith.** Main API gateway + user/org management, **plus** the **seven** folded domains this index's banner names — skiller (taxonomy, matching, embeddings), skillpath, jobsimulation, cms, storage, messenger, customerio-sync (`app/internal/{skiller,skillpath,jobsimulation,cms,storage,messenger,customeriosync}/`, all present @ `app` `ad9f3c49`) — and the AI-readiness subsystem, academy store, AI Labs LabSession. **`roadrunner` is NOT one of them**: `app/internal/roadrunner/` exists at no ref — at `ad9f3c49` **no path in the whole tree matches `roadrunner`** — and Judge0 execution was absorbed into the *jobsim* domain as `app/internal/jobsimulation/runner/`, wired at `app/internal/jobsimwiring/wiring.go:123`. This row listed a "roadrunner domain" until M257x iter-102, contradicting the ⚠️ at `:20-23` of this same file and the `app` row of [`platform-migration-status.md`](../architecture/platform-migration-status.md) |
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

## 01-004
- **id**: `B01-004`
- **corpus site**: `corpus/services/README.md:44-44` (table-row)
- **citation**: `roadrunner/terraform/main.tf:19`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/roadrunner/terraform/main.tf`  (96 lines)

**CLAIMING UNIT**

```md
| [`roadrunner.md`](roadrunner.md) | Roadrunner — **orphaned** (not "merged and undeployed") | Code-execution proxy to the Judge0 sandbox. Execution moved in-process with the jobsim engine and `backend` calls Judge0 directly via `JUDGE0_BASE_URL` — but prod terraform still reads `= 1` (`roadrunner/terraform/main.tf:19`, unchanged since `84a4b4f` / 2025-12-15 — it predates the fold and nobody has been back), even though `d11a403` removed its local container **and** its `repos.yml` entry |
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

## 01-005
- **id**: `B01-005`
- **corpus site**: `corpus/services/README.md:46-46` (table-row)
- **citation**: `messenger/internal/flow/flow.go:72-104`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/messenger/internal/flow/flow.go`  (135 lines)

**CLAIMING UNIT**

```md
| [`messenger.md`](messenger.md) | Messenger — **merged into `app`** | Centralized transactional email via Brevo + Liquid templates. Folded in at v9.0 "support-in-app"; container, `repos.yml` entry and `messenger` profile all deleted at `838d907`, and `app` gates the domain behind `MESSENGER_ENABLED` (unset = off on a laptop). Other services never called Brevo directly — they **publish Redis Stream events** the domain consumes (`messenger/internal/flow/flow.go:72-104` @ `fa47850`, `AddSubscriber("backend", …)`, 21 live handlers); `app` took over messenger's own consumer group. It *exposes* a `MessengerService` Connect-RPC surface, but **no service ever constructed a client for it**: `MESSENGER_RPC_ADDR` appears in no repo — every clone at its own named ref, nested repos included — and `git -C stack-demo/platform log -S 'MESSENGER_RPC' --oneline 0c91421d` returns **0** commits that ever set it (positive control at the same repo+ref: `-S 'SKILLER_RPC'` returns **7**) |
```

**CITED CONTENT**

```
    69  
    70  func (h *Manager) Subscribe() {
    71  	sub := pubsub.NewSubscriber()
    72  	h.subServer.AddSubscriber("backend", sub.AddHandler(
    73  		// skill paths
    74  		pubsub.EventHandler(h.OrgSkillPathAssignedHandler),
    75  		pubsub.EventHandler(h.OrgSkillPathUnassignedHandler),
    76  		pubsub.EventHandler(h.OrgSkillPathAssignmentCompletedHandler),
    77  		pubsub.EventHandler(h.OrgSkillPathAssignmentPastDueHandler),
    78  		pubsub.EventHandler(h.OrgSkillPathAssignmentDueDateUpdatedHandler),
    79  		// Job Simulation
    80  		pubsub.EventHandler(h.OrgJobSimulationAssignedHandler),
    81  		pubsub.EventHandler(h.OrgJobSimulationUnassignedHandler),
    82  		pubsub.EventHandler(h.OrgJobSimulationAssignmentCompletedHandler),
    83  		// pubsub.EventHandler(h.OrgJobSimulationAssignmentPastDueHandler), // not implemented
    84  		pubsub.EventHandler(h.OrgJobSimulationAssignmentDueDateUpdatedHandler),
    85  		// Content-agnostic assignment digest (skill path / sim / academy / lab), flag-gated
    86  		pubsub.EventHandler(h.OrgContentAssignedHandler),
    87  		// Content-agnostic COMPLETED notification to the assigner (academy / lab — the
    88  		// types with no dedicated completed event; sp/sim keep their flows above)
    89  		pubsub.EventHandler(h.OrgContentCompletedHandler),
    90  		// Organization
    91  		pubsub.EventHandler(h.OrganizationMemberInvitedHanler),
    92  		pubsub.EventHandler(h.MemberInvitationReminderHandler),
    93  		// AI Readiness
    94  		pubsub.EventHandler(h.AIReadinessCompletedHandler),
    95  		// AI Readiness Notifications (M400 stubs — bodied in M402/M404)
    96  		pubsub.EventHandler(h.AIReadinessMemberInvitedHandler),
    97  		pubsub.EventHandler(h.AIReadinessMemberReminderHandler),
    98  		pubsub.EventHandler(h.AIReadinessCycleLaunchedHandler),
    99  		pubsub.EventHandler(h.AIReadinessManagerDigestHandler),
   100  		// Course Builder
   101  		pubsub.EventHandler(h.CourseBuildCompletedHandler),
   102  		pubsub.EventHandler(h.CourseBuildFailedHandler),
   103  		pubsub.EventHandler(h.CoursePublishedHandler),
   104  	))
   105  	h.subServer.AddSubscriber("jobsimulation", sub.AddHandler(
   106  		pubsub.EventHandler(h.JobsimulationSessionStartedHandler),
   107  		pubsub.EventHandler(h.JobsimulationSessionEndedHandler),
```

## 01-006
- **id**: `B01-006`
- **corpus site**: `corpus/services/academy-backend.md:52-64` (bullet)
- **citation**: `app/main.go:471-472`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
*   **Managers** (constructed in `app/main.go`):
    *   `academy.Manager` (`academy.go`) — per-user study-state service (self-only identity via `authn`; upsert-retry
        on the first-write race; `SELECT … FOR UPDATE` in production).
    *   `academy.ContentManager` (`content.go`, `body.go`, `content_import.go`) — the global catalog + bodies with
        **fail-closed tenancy** (`TenancyResolver`; `ErrNotFound` = a deliberate hard-404 so "no access" is
        indistinguishable from "no such row") + a **tier gate** (`PremiumChecker`/`SubscriptionPremiumChecker`).
    *   `academy.EmbeddingsManager` (`embeddings.go`) — materializes + serves path embeddings; feeds
        `internal/aireadiness/suggested_path.go`.
    *   `academy.AssetUploader` (`asset.go`) — uploads path cover images + intro audio to the **public S3
        bucket**. Since the **storage-in-app v9.0 cutover** (`app` `9d00a313` = v1.367.0) that is an
        **in-process** public storage manager, not an RPC hop: `app/main.go:471-472` constructs the two
        managers from `STORAGE_S3_BUCKET` / `STORAGE_S3_PUBLIC_BUCKET`, and `STORAGE_RPC_ADDR` is gone from
        the Go source (3 remaining occurrences at that ref, **all comments**, zero reads).
```

**CITED CONTENT**

```
   468  	if err != nil {
   469  		logger.Error("bedrock client unavailable; talk-to-data disabled", "error", err)
   470  		return
   471  	}
   472  
   473  	// Managers are constructed here in the root and assembled into the App data
   474  	// struct below — there is no functional-options builder anymore (the app.With*
   475  	// options were removed). Construction order follows data dependencies; global
```

## 01-007
- **id**: `B01-007`
- **corpus site**: `corpus/services/academy-backend.md:116-120` (bullet)
- **citation**: `cmd/academyImport/main.go:239-243`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/cmd/academyImport/main.go`  (265 lines)

**CLAIMING UNIT**

```md
*   **Prerequisites**: runs **as part of `app`** — no standalone academy process. Uses the shared platform Postgres
    (`public` schema). Env: `SUPABASE_DB_CONN` (the DSN the cmd tools use), **`STORAGE_S3_PUBLIC_BUCKET`** +
    AWS credentials for asset upload (**not** `STORAGE_RPC_ADDR` — `cmd/academyImport/main.go:239-243`
    @ `app` `9d00a313`: *"the standalone storage service is gone … it therefore needs the BUCKET, not an RPC
    address"*, and it hard-errors on an empty bucket at `:244-246`), `ACADEMY_CONTENT_API_TOKEN` (the write API).
```

**CITED CONTENT**

```
   236  func buildUploader() (*academy.AssetUploader, error) {
   237  	coverNS := os.Getenv("ACADEMY_COVER_STORAGE_PATH")
   238  	audioNS := os.Getenv("ACADEMY_AUDIO_STORAGE_PATH")
   239  	// v9.0 cutover: the standalone storage service is gone, so this CLI talks to S3
   240  	// through the same in-process public manager the service uses. It therefore needs
   241  	// the BUCKET, not an RPC address — and AWS credentials in its own environment.
   242  	// An empty bucket would silently write to local disk, so require it explicitly.
   243  	publicBucket := os.Getenv(appstorage.EnvPublicBucket)
   244  	if publicBucket == "" {
   245  		return nil, fmt.Errorf("%s is required (or pass --metadata-only)", appstorage.EnvPublicBucket)
   246  	}
```

## 01-008
- **id**: `B01-008`
- **corpus site**: `corpus/services/academy-backend.md:121-138` (bullet)
- **citation**: `serverTenant.js:115-145`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/src/lib/serverTenant.js`  (337 lines)

**CLAIMING UNIT**

```md
*   **cmd binaries**:
    *   `cmd/academy-seed` — local dev/test seeder. Seeds chapter-progress + last-activity **through the Manager**
        (so writes go through the same monotonic-merge + self-only privacy paths, idempotent by construction). Flags:
        `--user-email`/`--user-id`, `--fixture`, `--reset`, `--dry-run`, `--list`. Local dev only.
        **Demo caveat**: on a demo the frontend has **no `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`**, so any rows
        `academy-seed` writes have **no reader** — i.e. `academy-seed` is **moot on a demo**. **It does NOT
        "serve its committed FS catalog": there is no FS-as-published fallback in the app** — this passage
        said there was until M257x iter-108, contradicting
        [`ant-academy.md`](./ant-academy.md)`:82-88`, which states the opposite in bold.
        `getServerCatalogView()` resolves a null backend result to the **empty view**
        (`serverTenant.js:115-145` — *"the cutover is intentional, not reversible-on-error"*). A demo grid
        renders cards **only** because the rext demo-patch `demo-stack/patches/academy-fs-published-fallback`
        *restores* that removed fallback on the demo's ephemeral clone; if the patch is refused, the grid is
        empty. The demo CTA is a real `/courses/<slug>` link (see
        [`../ops/demo/content-stories-routes.md`](../ops/demo/content-stories-routes.md)).
    *   `cmd/academyImport` — idempotent, resumable catalog/metadata/**bodies** importer from a manifest JSON
        (`--manifest`, `--content-root`, `--checkpoint`), uploading covers/audio via the asset uploader.
    *   `cmd/coursebuilder-e2e`, `cmd/coursebuilder-liverun` — exercise the Course Builder → academy publish path.
```

**CITED CONTENT**

```
   112   * tenant metadata and is the one piece not yet modeled in the backend catalog,
   113   * so it passes through verbatim (same as the backend adapter does).
   114   */
   115  function emptyCatalogView() {
   116    return { chapters: [], skillPaths: {}, series: [], bundles: PUBLIC_BUNDLES, catalogVersion: CATALOG_VERSION }
   117  }
   118  
   119  /**
   120   * The catalog the current (authenticated) user is allowed to receive: public +
   121   * the user's org tenants. This is what authed RSCs pass to the client instead
   122   * of importing the raw CHAPTERS array — so tenant content never ships to a
   123   * browser that can't see it.
   124   *
   125   * DB-authoritative (M7): the catalog is read UNCONDITIONALLY from the academy
   126   * backend (`getBackendCatalogView(eids)` → already tenant-filtered server-side,
   127   * with a chapter-level eid refinement in the adapter). DB presence gates
   128   * visibility — a path/chapter renders only if the backend returns it; there is
   129   * NO FS-as-published fallback. The eids are threaded in for the adapter's
   130   * chapter-level tenant refinement. A null backend result (not-composed /
   131   * outage) resolves to the EMPTY view rather than the committed FS catalog: the
   132   * cutover is intentional, not reversible-on-error.
   133   *
   134   * DEV DRAFT LAYER (M8): when `draftsEnabled()` (dev + the ACADEMY_SHOW_DRAFTS
   135   * opt-in; NEVER in production — hard-blocked), FS-only content (present-on-FS ∧
   136   * absent-from-DB) is merged on top of the DB view tagged `_draft: true`, so an
   137   * author previews locally-committed-but-not-yet-exported content. Off (incl. all
   138   * production) → the DB view passes through verbatim, zero behavior change. The
   139   * merge runs server-side; the client still receives only the threaded view
   140   * (catalog-client-boundary intact). The eids thread through so a tenant-scoped
   141   * draft stays tenant-gated.
   142   */
   143  export async function getServerCatalogView() {
   144    const eids = await getUserEids()
   145    const view = (await getBackendCatalogView(eids)) ?? (process.env.ACADEMY_DEMO_FS_PUBLISHED === '1' ? (v => ({ ...v, chapters: (v.chapters ?? []).map(({ _draft, _origin, ...c }) => c), series: (v.series ?? []).map(({ _draft, _origin, ...s }) => s), skillPaths: Object.fromEntries(Object
```

## 01-009
- **id**: `B01-009`
- **corpus site**: `corpus/services/ai-labs.md:11-21` (paragraph)
- **citation**: `internal/web/backend/credits/handler.go:12`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/web/backend/credits/handler.go`  (140 lines)

**CLAIMING UNIT**

```md
> ### ⚠️ Two different meanings of "v6.0 shared purse" — read this first
> - The credit **ledger primitive** (`internal/credits`, "Wave 6") **is built and live** — but wired **only to
>   [Course Builder](./coursebuilder.md)**, not to AI Labs.
> - The **AI Labs self-serve credits initiative** branded **v6.0 "shared purse"** (org self-serve *buy* AI-Labs
>   credits + an *enforcing* shared-pool wallet) is **DESIGNED / QUEUED, NOT BUILT** — a knowledge-plan release
>   (`app` `knowledge/plan/releases/06.00-shared-purse/`, milestones M600–M607, all planned). Measured at `app`
>   **`ad9f3c49`** (which was `origin/main` on 2026-08-06) there is **no `checkout.session.completed` webhook**
>   (0 occurrences in Go source), **no labs↔credits linkage** (`internal/labs/` imports `internal/credits`
>   nowhere), **and `/credits/purchase` was removed (Wave 13)** (`internal/web/backend/credits/handler.go:12`).
>   `v6.0` is a **knowledge-plan release number, NOT the `app` SemVer** (those run in the `v1.3xx` range — newest tag **`v1.369.0`**, re-read 2026-08-06 at that same `ad9f3c49`, seven commits past the tag: `git describe --tags ad9f3c49` → `v1.369.0-7-gad9f3c498`. **A version — and a branch label — is a reading at a ref, never a standing "current"**: this line read `v1.363.2` @ `5ba17044` for six `app` releases, then `v1.369.0` @ *origin/main* `2035f9a4` until that label expired five commits later. `2035f9a4` is still a valid **pin** (`v1.369.0-2-g2035f9a40`); only the moving label rotted).
> This doc documents the **shipped reality**; the shared-purse unification is flagged as planned where relevant.
```

**CITED CONTENT**

```
     9  //
    10  // There is deliberately no self-serve purchase route: top-ups are
    11  // handled out-of-band via support@anthropos.work (Wave 12). The mock
    12  // POST /credits/purchase was removed in Wave 13 — the informational
    13  // packages still ride on GET /balance so the UI can render tiers.
    14  package credits
    15  
```

## 01-010
- **id**: `B01-010`
- **corpus site**: `corpus/services/ai-readiness.md:13-68` (paragraph)
- **citation**: `app-aireadiness-snapshot-loadmembers.yaml:42`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/demo-stack/patches/app-aireadiness-snapshot-loadmembers/app-aireadiness-snapshot-loadmembers.yaml`  (86 lines)

**CLAIMING UNIT**

```md
> ### ⚠️ Package refactor — `internal/workforce/` → `internal/aireadiness/` (app v1.351.1, M247 refresh)
>
> The whole AI-Readiness domain moved out of `app/internal/workforce/` into a **new package
> `app/internal/aireadiness/`** (package name `aireadiness`) — commit `4c28365f` ("Refactor AI Readiness domain:
> migrate workforce dependencies to aireadiness package", 2026-07-22). `workforce` keeps the org-analytics KPIs;
> `aireadiness` owns everything readiness-scoped. **The only remaining dependency on `workforce` is the member
> directory** (the `WorkforceDirectory` interface — `LoadMembers`/`LoadMembersByUserIDs`, whose implementations
> **stayed** in `app/internal/workforce/members.go`).
>
> **File renames** (older `app/internal/workforce/…` anchors elsewhere in this doc refer to the pre-refactor
> location — resolve them under `internal/aireadiness/`):
>
> | Old (`internal/workforce/`) | New (`internal/aireadiness/`) |
> |---|---|
> | `ai_readiness.go` | **`readiness.go`** (the scoring engine + read entrypoints) |
> | `ai_readiness_v2.go` | `scoring.go` (archetype/axis math + bands) |
> | `ai_readiness_csv.go` | `csv.go` |
> | `readiness_steps.go` | `steps.go` |
> | `readiness_narrative.go` | `narrative.go` |
> | `how_we_measure_v2.go` | folded into `how_we_measure.go` (`computeInterviewInsightsV2`) |
> | `cycles.go`/`compare.go`/`diagnosis.go`/`provision.go`/`defaults.go`/… | same names under `internal/aireadiness/` |
> | `emailoverride/`, `emailpreview/`, `notifications/` | same, under `internal/aireadiness/` |
> | `*_test.go` (scoring/steps/cycle suites) | moved with the package; harness `testdb_test.go` (pgtest) |
>
> **D-07 demopatch re-anchor (load-bearing).** The `app-aireadiness-snapshot-loadmembers` demo-patch anchored on
> `app/internal/workforce/ai_readiness.go` at the `buildResponseFromSnapshots → loadMembers(orgID,"")` call. That
> file no longer exists — the call is now at **`app/internal/aireadiness/readiness.go`**, `buildResponseFromSnapshots`,
> as **`m.workforce.LoadMembers(ctx, orgID, "")`** *through the `WorkforceDirectory` interface* (the bounded swap
> `LoadMembers → LoadMembersByUserIDs` is now expressible at that interface call site, since `WorkforceDirectory`
> already exposes `LoadMembersByUserIDs`). **The re-anchor has LANDED** — v2.7 **M254**, not the M250 the M246
> drift-ledger's **D-07** item originally assigned it to: th
```

**CITED CONTENT**

```
    39  
    40  id: app-aireadiness-snapshot-loadmembers
    41  repo: app
    42  path: internal/aireadiness/readiness.go
    43  scope: demo
    44  
    45  # sha256 of the PRISTINE on-disk file (demo clone @ app 3df8536 / v2.7 pin) — the baseline; apply_patch.py
```

## 01-011
- **id**: `B01-011`
- **corpus site**: `corpus/services/ai-readiness.md:86-91` (bullet)
- **citation**: `app/internal/data/ent/enum/organization_settings.go:47`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/data/ent/enum/organization_settings.go`  (90 lines)

**CLAIMING UNIT**

```md
1. **Org setting** — a row in `organization_settings` with `setting = 'ai_readiness'`, `is_enabled = true`
   (`app/internal/data/ent/enum/organization_settings.go:47` → `OrganizationSettingAIReadiness = "ai_readiness"`;
   checked by `isAIReadinessEnabled` in `app/internal/aireadiness/steps.go` — formerly
   `workforce/readiness_steps.go::isAIReadinessEnabled`). No row =
   off. Exposed to the FE as the GraphQL query `aiReadinessEnabled: Boolean!`
   (`resolver_ai_readiness.go` — returns `false`, not an error, for non-enabled orgs).
```

**CITED CONTENT**

```
    44  	// "session ended" events are consumed to advance the plan. Enterprise toggle
    45  	// set directly in the DB by ops; intentionally NOT exposed through the
    46  	// GraphQL/RPC OrganizationSettings APIs.
    47  	OrganizationSettingAIReadiness OrganizationSetting = "ai_readiness"
    48  	// OrganizationSettingAIReadinessAutoAssign turns ON automatic assignment of a
    49  	// person's recommended Academy path the moment they finish all three AI
    50  	// Readiness steps. DEFAULT OFF: a missing row (the state every org starts in)
```

## 01-012
- **id**: `B01-012`
- **corpus site**: `corpus/services/ai-readiness.md:92-95` (bullet)
- **citation**: `apps/web/src/components/ai-readiness/data/useAiReadinessActive.ts:22`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/apps/web/src/components/ai-readiness/data/useAiReadinessActive.ts`  (38 lines)

**CLAIMING UNIT**

```md
2. **PostHog flag** `flag_ai_readiness` — gates the **member-facing** surface only. Its **sole consumer
   repo-wide** is `apps/web/src/components/ai-readiness/data/useAiReadinessActive.ts:22`
   (`useFeatureFlagEnabled(AI_READINESS_FLAG)`; the constant is declared at
   `components/ai-readiness/aiReadiness.constants.ts:26`).
```

**CITED CONTENT**

```
    19  export function useAiReadinessActive(): { active: boolean; loading: boolean } {
    20    // Sticky-resolve the flag so a transient `undefined` (PostHog re-bootstrapping
    21    // its cache, e.g. on tab refocus) never momentarily reads as "off".
    22    const rawFlag = useFeatureFlagEnabled(AI_READINESS_FLAG);
    23    const stickyFlag = useRef<boolean | undefined>(undefined);
    24    if (rawFlag !== undefined) stickyFlag.current = rawFlag;
    25    const flagEnabled = stickyFlag.current === true;
```

## 01-013
- **id**: `B01-013`
- **corpus site**: `corpus/services/ai-readiness.md:92-95` (bullet)
- **citation**: `components/ai-readiness/aiReadiness.constants.ts:26`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/apps/web/src/components/ai-readiness/aiReadiness.constants.ts`  (226 lines)

**CLAIMING UNIT**

```md
2. **PostHog flag** `flag_ai_readiness` — gates the **member-facing** surface only. Its **sole consumer
   repo-wide** is `apps/web/src/components/ai-readiness/data/useAiReadinessActive.ts:22`
   (`useFeatureFlagEnabled(AI_READINESS_FLAG)`; the constant is declared at
   `components/ai-readiness/aiReadiness.constants.ts:26`).
```

**CITED CONTENT**

```
    23  // Consumed client-side via useFeatureFlagEnabled() in useAIReadiness.ts.
    24  // TODO(ai-readiness): when GA-ing, widen the cohort / rollout in PostHog (no
    25  // code change needed), and optionally AND-in a real per-org setting here.
    26  export const AI_READINESS_FLAG = 'flag_ai_readiness';
    27  
    28  // Due-date text turns red under this many days.
    29  export const DUE_SOON_DAYS = 5;
```

## 01-014
- **id**: `B01-014`
- **corpus site**: `corpus/services/ai-readiness.md:110-164` (paragraph)
- **citation**: `apps/web/src/components/ai-readiness/data/useAiReadinessActive.ts:19-32`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/apps/web/src/components/ai-readiness/data/useAiReadinessActive.ts`  (38 lines)

**CLAIMING UNIT**

```md
> **These two gates are different layers — not a contradiction.** `stories-spec.md` (the `OrgSettingsSeeder` row)
> calls enablement "an **org setting**, not a PostHog flag": that is precise about the **enablement/data layer**
> (gate 1) the seeder writes — a `organization_settings` DB row, resolved from the M48 contract, which is *not*
> stored in PostHog. It does **not** deny gate 2: the next-web **member** surface *additionally* checks the
> PostHog `flag_ai_readiness` before rendering. Seeder-writes-the-setting (gate 1) and
> member-UI-also-checks-the-flag (gate 2) are complementary — but per the callout above, **only gate 1 is
> required for the manager dashboard**. The demo must satisfy gate 2 for the member journey, which is what
> the next section is about.
>
> ### ⚠️ How the demo satisfies gate 2 (the FE flag) — CORRECTED, M219 (v2.3 "cue to cue")
>
> **This section previously asserted the exact opposite of the truth, and the error is instructive.** It said:
> *"the demo next-web bakes no `NEXT_PUBLIC_POSTHOG_KEY`, so the client-side flag check has no PostHog backend
> to consult and does not block the route"* — i.e. that absence of PostHog **defaults the flag through**.
>
> **It does not. Absence of PostHog makes the flag `undefined`, and the code demands `=== true`:**
>
> ```ts
> const rawFlag     = useFeatureFlagEnabled(AI_READINESS_FLAG);  // no PostHog → undefined, FOREVER
> const flagEnabled = stickyFlag.current === true;               // undefined === true → FALSE
> const { orgEnabled } = useAiReadinessEnabled(flagEnabled);     // queried ONLY when the flag is on
> active = flagEnabled && orgEnabled === true;                   // → never active
> ```
>
> **The construct this paraphrases, named (M257x iter-115).** It is `useAiReadinessActive()` in
> `apps/web/src/components/ai-readiness/data/useAiReadinessActive.ts:19-32`, **@ `next-web-app`
> `8297c684`** — `:22` `rawFlag`, `:25` `flagEnabled`, `:29` `useAiReadinessEnabled(flagEnabled)`,
> `:32` `active: flagEnabled && orgEnabled === true`. **This block carried no file and no ref until
> iter-115**, which is the same defect the four `AIReadinessClient.tsx` anchors in this document carried
> — a citation-free paraphrase cannot be re-derived, and cannot be caught when it drifts. **Distinct
> surface from the manager dashboard**: this is the *member* gate (gate 2); the manager dashboard
> calls 
```

**CITED CONTENT**

```
    16   * user to `/home` only makes sense when the AI Readiness component will actually
    17   * be there.
    18   */
    19  export function useAiReadinessActive(): { active: boolean; loading: boolean } {
    20    // Sticky-resolve the flag so a transient `undefined` (PostHog re-bootstrapping
    21    // its cache, e.g. on tab refocus) never momentarily reads as "off".
    22    const rawFlag = useFeatureFlagEnabled(AI_READINESS_FLAG);
    23    const stickyFlag = useRef<boolean | undefined>(undefined);
    24    if (rawFlag !== undefined) stickyFlag.current = rawFlag;
    25    const flagEnabled = stickyFlag.current === true;
    26  
    27    // The org boolean (data, never an error) is the real enablement; query it only
    28    // when the flag is on, so non-flagged users fire no request.
    29    const { orgEnabled, loading } = useAiReadinessEnabled(flagEnabled);
    30  
    31    return {
    32      active: flagEnabled && orgEnabled === true,
    33      // Only meaningfully "loading" while the flag is on and the org boolean is
    34      // still unknown; a flag-off user resolves immediately to inactive.
    35      loading: flagEnabled && orgEnabled === undefined && loading,
```

## 01-015
- **id**: `B01-015`
- **corpus site**: `corpus/services/ai-readiness.md:249-255` (bullet)
- **citation**: `stack-seeding/seeders/ai_readiness_interview_report.go:185-189`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/stack-seeding/seeders/ai_readiness_interview_report.go`  (431 lines)

**CLAIMING UNIT**

```md
- the **five** usage KPIs — `avg_adoption` · `avg_depth` · `avg_ownership` · `avg_transformation` ·
  `avg_originality`, the set `aiReadinessUsageKPIs` returns at
  `stack-seeding/seeders/ai_readiness_interview_report.go:185-189`, whose own comment at `:149` reads *"All
  five are always emitted here"* — are **DERIVED from the org's own seeded Step-3 session scores** (the same
  raw numbers the frozen snapshot rolls up), so the tiles agree with the funnel rather than being invented.
  (This bullet said *"four"* until M257x; it is the same five ids the `catalog_kpis[]` row above enumerates
  and the same five *part 5* of § *The FILLED-ness contract* re-derives.);
```

**CITED CONTENT**

```
   182  	// (every usageDimSpec bands high ≥61 / mid ≥41 / low <41). Adoption leads (the beachhead),
   183  	// depth close behind; ownership sits mid; transformation and originality trail into the softer bands.
   184  	return []map[string]any{
   185  		{"id": "avg_adoption", "value": clamp(mean + 15)},
   186  		{"id": "avg_depth", "value": clamp(mean + 3)},
   187  		{"id": "avg_ownership", "value": clamp(mean - 4)},
   188  		{"id": "avg_transformation", "value": clamp(mean - 12)},
   189  		{"id": "avg_originality", "value": clamp(mean - 20)},
   190  	}
   191  }
   192  
```

## 01-016
- **id**: `B01-016`
- **corpus site**: `corpus/services/ai-readiness.md:352-363` (bullet)
- **citation**: `packages/core-js/src/constants/urls.ts:51`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/packages/core-js/src/constants/urls.ts`  (70 lines)

**CLAIMING UNIT**

```md
- **`/ai-readiness` is the only readiness route the navbar links** — `AI_READINESS_URL`
  (`packages/core-js/src/constants/urls.ts:51`), consumed by `packages/ui/src/NavBar/useNavbarSections.tsx`
  — imported at `:4`, built into `aiReadinessMenuItem` at `:401-408` (`key: AI_READINESS_URL` at `:403`),
  gated at `:569` (`showAIReadiness ? aiReadinessMenuItem : null`). ⚠️ **This doc said `urls.ts:52` for four
  readings, and `:52` has never been `AI_READINESS_URL` at any ref reachable from this clone** — over every
  commit touching that file the constant sits at line **41, 50 or 51 only**; `:52` is `TALK_TO_DATA_URL` at
  `8297c684` and was `ORGANIZATION_FEEDBACK_URL` at `bb3313bc`. The citation *resolved*,
  which is why an anchor-existence check passed it while a reader following it landed on a different route;
  corrected M257x iter-102. A repo-wide grep finds the constant in exactly those
  two **source** files — the only other non-`node_modules` hit is the platform's own KB page,
  `knowledge/ai-readiness/frontend-architecture.md:15`. `/ai-readiness` is also the only readiness route
  next-web's own e2e covers (`e2e/specs/web.ai-readiness.spec.ts`).
```

**CITED CONTENT**

```
    48  // ENTERPRISE
    49  export const WORKFORCE_URL = '/enterprise/workforce';
    50  export const WORKFORCE_ANALYTICS_URL = '/enterprise/workforce/analytics';
    51  export const AI_READINESS_URL = '/ai-readiness';
    52  export const TALK_TO_DATA_URL = '/enterprise/talk-to-data';
    53  export const ORGANIZATION_FEEDBACK_URL = '/enterprise/organization-feedback';
    54  export const INSIGHTS_URL = '/enterprise/activity-dashboard';
```

## 01-017
- **id**: `B01-017`
- **corpus site**: `corpus/services/ai-readiness.md:352-363` (bullet)
- **citation**: `urls.ts:52`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/packages/core-js/src/constants/urls.ts`  (70 lines)

**CLAIMING UNIT**

```md
- **`/ai-readiness` is the only readiness route the navbar links** — `AI_READINESS_URL`
  (`packages/core-js/src/constants/urls.ts:51`), consumed by `packages/ui/src/NavBar/useNavbarSections.tsx`
  — imported at `:4`, built into `aiReadinessMenuItem` at `:401-408` (`key: AI_READINESS_URL` at `:403`),
  gated at `:569` (`showAIReadiness ? aiReadinessMenuItem : null`). ⚠️ **This doc said `urls.ts:52` for four
  readings, and `:52` has never been `AI_READINESS_URL` at any ref reachable from this clone** — over every
  commit touching that file the constant sits at line **41, 50 or 51 only**; `:52` is `TALK_TO_DATA_URL` at
  `8297c684` and was `ORGANIZATION_FEEDBACK_URL` at `bb3313bc`. The citation *resolved*,
  which is why an anchor-existence check passed it while a reader following it landed on a different route;
  corrected M257x iter-102. A repo-wide grep finds the constant in exactly those
  two **source** files — the only other non-`node_modules` hit is the platform's own KB page,
  `knowledge/ai-readiness/frontend-architecture.md:15`. `/ai-readiness` is also the only readiness route
  next-web's own e2e covers (`e2e/specs/web.ai-readiness.spec.ts`).
```

**CITED CONTENT**

```
    49  export const WORKFORCE_URL = '/enterprise/workforce';
    50  export const WORKFORCE_ANALYTICS_URL = '/enterprise/workforce/analytics';
    51  export const AI_READINESS_URL = '/ai-readiness';
    52  export const TALK_TO_DATA_URL = '/enterprise/talk-to-data';
    53  export const ORGANIZATION_FEEDBACK_URL = '/enterprise/organization-feedback';
    54  export const INSIGHTS_URL = '/enterprise/activity-dashboard';
    55  export const ENTERPRISE_LABS_URL = '/enterprise/labs';
```

## 01-018
- **id**: `B01-018`
- **corpus site**: `corpus/services/ai-readiness.md:352-363` (bullet)
- **citation**: `knowledge/ai-readiness/frontend-architecture.md:15`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/knowledge/ai-readiness/frontend-architecture.md`  (64 lines)

**CLAIMING UNIT**

```md
- **`/ai-readiness` is the only readiness route the navbar links** — `AI_READINESS_URL`
  (`packages/core-js/src/constants/urls.ts:51`), consumed by `packages/ui/src/NavBar/useNavbarSections.tsx`
  — imported at `:4`, built into `aiReadinessMenuItem` at `:401-408` (`key: AI_READINESS_URL` at `:403`),
  gated at `:569` (`showAIReadiness ? aiReadinessMenuItem : null`). ⚠️ **This doc said `urls.ts:52` for four
  readings, and `:52` has never been `AI_READINESS_URL` at any ref reachable from this clone** — over every
  commit touching that file the constant sits at line **41, 50 or 51 only**; `:52` is `TALK_TO_DATA_URL` at
  `8297c684` and was `ORGANIZATION_FEEDBACK_URL` at `bb3313bc`. The citation *resolved*,
  which is why an anchor-existence check passed it while a reader following it landed on a different route;
  corrected M257x iter-102. A repo-wide grep finds the constant in exactly those
  two **source** files — the only other non-`node_modules` hit is the platform's own KB page,
  `knowledge/ai-readiness/frontend-architecture.md:15`. `/ai-readiness` is also the only readiness route
  next-web's own e2e covers (`e2e/specs/web.ai-readiness.spec.ts`).
```

**CITED CONTENT**

```
    12  
    13  ## What it will cover
    14  
    15  1. **Route registration** — `apps/web/src/app/(authenticated)/(verified)/ai-readiness/page.tsx`, the manager auth gate via `EnterpriseWrapper`, the `AI_READINESS_URL` constant.
    16  2. **Component tree** — `AIReadinessClient` → `PageShell` → tabs → matrix / drawers / action plan.
    17  3. **Container → View convention** — containers in `apps/web/src/components/organisms/AIReadiness/`; pure views in `packages/ui/src/AIReadiness/`. Same split as Workforce Intelligence.
    18  4. **Hook pattern** — `useAIReadiness({ cycleId? })` based on `useWorkforceAIReadiness` template; React Query with `STALE_SLOW` cache tier.
```

## 01-019
- **id**: `B01-019`
- **corpus site**: `corpus/services/ai-readiness.md:364-366` (bullet)
- **citation**: `hooks/useWorkforceAIReadiness.ts:23-27`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/apps/web/src/hooks/useWorkforceAIReadiness.ts`  (37 lines)

**CLAIMING UNIT**

```md
- **The legacy route was an orphan**: no nav entry, no workforce tab (`WorkforceNewClient.tsx:125-151` omitted it),
  no redirect points at it. Its hook (`hooks/useWorkforceAIReadiness.ts:23-27`) calls
  `GET /api/workforce/ai-readiness?tag=` — **there is no `cycle` param in it at all**, and it never calls `/cycles`.
```

**CITED CONTENT**

```
    20      queryFn: async ({ signal }) => {
    21        const token = await getTokenRef.current();
    22        if (!token) throw new Error('Unauthorized');
    23        return callWorkforceAPI<WorkforceAIReadinessResponse>(
    24          '/ai-readiness',
    25          token,
    26          { searchParams: { tag }, signal }
    27        );
    28      },
    29      enabled: !!orgId && (options?.enabled ?? true),
    30      meta: { error: 'Failed to fetch workforce AI readiness' },
```

## 01-020
- **id**: `B01-020`
- **corpus site**: `corpus/services/ai-readiness.md:370-372` (paragraph)
- **citation**: `useAiReadinessActive.ts:22`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/apps/web/src/components/ai-readiness/data/useAiReadinessActive.ts`  (38 lines)

**CLAIMING UNIT**

```md
**The `flag_ai_readiness` PostHog flag gates the EMPLOYEE side only** (`useAiReadinessActive.ts:22`). It does
**not** select between manager trees — it never did, back when there were two. The (one) manager
dashboard gates purely on the GraphQL `aiReadinessEnabled` boolean plus `isEnterprise` nav visibility.
```

**CITED CONTENT**

```
    19  export function useAiReadinessActive(): { active: boolean; loading: boolean } {
    20    // Sticky-resolve the flag so a transient `undefined` (PostHog re-bootstrapping
    21    // its cache, e.g. on tab refocus) never momentarily reads as "off".
    22    const rawFlag = useFeatureFlagEnabled(AI_READINESS_FLAG);
    23    const stickyFlag = useRef<boolean | undefined>(undefined);
    24    if (rawFlag !== undefined) stickyFlag.current = rawFlag;
    25    const flagEnabled = stickyFlag.current === true;
```

## 01-021
- **id**: `B01-021`
- **corpus site**: `corpus/services/ai-readiness.md:499-507` (bullet)
- **citation**: `ops/demo/stories-spec.md:599`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/corpus/ops/demo/stories-spec.md`  (841 lines)

**CLAIMING UNIT**

```md
- **Closed cycle → the dashboard reads frozen snapshots.** `buildResponseFromSnapshots` reads `ai_readiness_snapshots`
  directly, so a **closed**-cycle showcase can be seeded **snapshot-direct** (write the `frozen_*` rows + flip the
  cycle to `closed`) with **no underlying signals** — the world reads as a *finished* assessment. **This is the
  strategy M51 shipped** (`AIReadinessConfigSeeder` writes the cycle `closed` + `AIReadinessFunnelSeeder` writes 199
  frozen `ai_readiness_snapshots`), after iters 03→06 falsified the active-signals path — on the premise that the
  live-recompute never completed in the coverage harness budget (a per-skill federated translation N+1, the M46
  per-object-RPC class). ⚠️ **That premise was refuted at M219: the recompute takes 2.09 s.** The strategy M51
  shipped is unchanged; the reason given for it is not. Retracted here and at
  [`ops/demo/stories-spec.md:599`](../ops/demo/stories-spec.md) at M257x iter-49.
```

**CITED CONTENT**

```
   596  The M48 contract offers two seed strategies (see [`../../services/ai-readiness.md`](../../services/ai-readiness.md)):
   597  an **active** cycle (the dashboard live-recomputes from signals) or a **closed** cycle (the dashboard reads
   598  pre-computed `ai_readiness_snapshots` directly). M51's iters 03→06 built and then **falsified** the active-signals
   599  path — believing the live-recompute never completed in the coverage harness's budget (a per-skill federated
   600  translation N+1, the M46 per-object-RPC class).
   601  > **⚠️ That premise was REFUTED at v2.3 M219: the live recompute completes in 2.09 s.** The
   602  > never-completes claim is retracted, and is recorded as retracted at
```

## 01-022
- **id**: `B01-022`
- **corpus site**: `corpus/services/ai-readiness.md:509-510` (paragraph)
- **citation**: `app/internal/aireadiness/readiness.go:289`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/aireadiness/readiness.go`  (1238 lines)

**CLAIMING UNIT**

```md
  **The frozen path is reachable BOTH cycle-scoped and by default.** `GetAIReadinessWithOptions`
  (`app/internal/aireadiness/readiness.go:289`) has two routes into `buildResponseFromSnapshots`:
```

**CITED CONTENT**

```
   286  
   287  // GetAIReadinessWithOptions is the v3.0 entrypoint that supports cycle scoping +
   288  // per-person enumeration.
   289  func (m *Manager) GetAIReadinessWithOptions(ctx context.Context, orgID uuid.UUID, opts GetAIReadinessOptions) (*AIReadinessResponse, error) {
   290  	// Cycle path: if a closed cycle is requested, return snapshot data.
   291  	if opts.CycleID != nil {
   292  		cycle, err := m.GetAIReadinessCycle(ctx, orgID, *opts.CycleID)
```

## 01-023
- **id**: `B01-023`
- **corpus site**: `corpus/services/ai-readiness.md:519-543` (paragraph)
- **citation**: `useWorkforceAIReadiness.ts:23-27`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/apps/web/src/hooks/useWorkforceAIReadiness.ts`  (37 lines)

**CLAIMING UNIT**

```md
  > **✅ CORRECTED M219 (v2.3 "cue to cue") — the old M51 iter-07 caveat here was MISATTRIBUTED, and it sent a
  > later milestone hunting for a demo-patch that was never needed.**
  >
  > The retracted claim: *"the demo FE fires the data GET WITHOUT `?cycle=` … and never fires the `/cycles` list
  > that supplies `latestClosedCycle.id`"*, concluded to be **platform-bound**.
  >
  > **What is actually true.** The **CURRENT** dashboard (`AIReadinessClient.tsx:155-156`, **@ `next-web-app`
  > `8297c684`**) computes
  > `effectiveCycleId = selectedCycle ?? activeCycle?.id ?? latestClosedCycle?.id` and gates the data GET on
  > `enabled: featureOn && cyclesQ.isFetched` at **`:171`** — i.e. it **waits for `/cycles`, then passes
  > `?cycle=`**. (These read `:153-154` and `:166-170` until M257x iter-115: the first is the comment above
  > the const, and the second is a range that stops **one line short** of the gate it is offered as evidence
  > for — `:168-170` is the `useAIReadiness({` call and its first two options.) Verified live
  > against a running demo (authenticated as the manager hero): `/cycles` returns the seeded cycle, and the
  > frozen read answers **HTTP 200 in 24 ms**.
  >
  > The iter-07 probe was watching the **LEGACY** page (`/enterprise/workforce/ai-readiness`), whose hook
  > (`useWorkforceAIReadiness.ts:23-27`) has **no `cycle` param at all** and **never calls `/cycles`** — which is
  > exactly the behavior that was observed and then attributed to the platform. **It was a pointer bug, not a
  > platform gap.** See § Surfaces (UI) above.
  >
  > **And the live path does not "never complete".** Measured on the same 199-member org:
  > **LIVE `GET /api/workforce/ai-readiness` → HTTP 200 · 2.09 s · 304 KB.** The M51-era "translation-N+1 that
  > never completes in-budget" is **not reproducible** on the app tag the demo builds today. Re-measure before
  > relying on either number; do not re-derive them from prose.
```

**CITED CONTENT**

```
    20      queryFn: async ({ signal }) => {
    21        const token = await getTokenRef.current();
    22        if (!token) throw new Error('Unauthorized');
    23        return callWorkforceAPI<WorkforceAIReadinessResponse>(
    24          '/ai-readiness',
    25          token,
    26          { searchParams: { tag }, signal }
    27        );
    28      },
    29      enabled: !!orgId && (options?.enabled ?? true),
    30      meta: { error: 'Failed to fetch workforce AI readiness' },
```

## 01-024
- **id**: `B01-024`
- **corpus site**: `corpus/services/ai-readiness.md:653-667` (paragraph)
- **citation**: `apps/web/src/hooks/useAIReadiness.ts:326`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/apps/web/src/hooks/useAIReadiness.ts`  (753 lines)

**CLAIMING UNIT**

```md
> **Measured, not assumed — and it corrects the finding that opened this thread.** The **current** dashboard's
> *"✨ Handled for you this cycle"* tile renders **`skillsMapped` / `handsOnMinutes` / `interviewMinutes`** —
> and **does not render `interviewQuestions` at all** (**all anchors in this parenthetical re-derived at
> `next-web-app` `8297c684`, 2026-08-06**: `HowWeMeasureTab.tsx:1879` opens the
> `{/* ===== C · Handled for you this cycle ===== */}` block, label `:1903`, then exactly three cells —
> `skillsMapped` `:1915`, `handsOnMinutes` `:1921`, `interviewMinutes` `:1927`; `grep -c interviewQuestions`
> over that 1,989-line file returns **0**. The field exists in the API and in the FE's TypeScript type,
> `apps/web/src/hooks/useAIReadiness.ts:326` — inside `export interface AIReadinessCycleTotals` (`:323-330`)
> — and is drawn by nothing. ⚠️ **Read that `:326` as a pin, not a standing line**: it was `:274` at
> `bb3313bc`, this doc carried `:274` and `:326` in successive iters, and at each ref the *other* number named
> an unrelated construct. The interface name is the contract; the line number is the convenience). So its zero was a
> **payload** zero, not a visible empty cell. Filled regardless — an interview with no questions is not real
> data — but the honest claim is that this tile's *visible* zero-risk lives in the three cells that do render,
> which the coverage sweep now fences with a **non-zero-value** assert rather than a label assert (a section
> that renders with all-zero numbers is an empty section wearing a hat).
```

**CITED CONTENT**

```
   323  export interface AIReadinessCycleTotals {
   324    skillsMapped: number;
   325    handsOnMinutes: number;
   326    interviewQuestions: number;
   327    // Total minutes across AI-interview sessions (→ "hours saved" tile). Optional:
   328    // present only once the backend computes it.
   329    interviewMinutes?: number;
```

## 01-025
- **id**: `B01-025`
- **corpus site**: `corpus/services/ant-academy.md:65-77` (paragraph)
- **citation**: `src/writeThrough/index.js:247`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/src/writeThrough/index.js`  (319 lines)

**CLAIMING UNIT**

```md
> **⚠️ The write path is the CLIENT harness, not the beacon route** (corrected M257x iter-102 — this sentence
> sourced the mutations to `code/app/api/academy/beacon/route.js`, which states the mechanism in the opposite
> order). Measured @ `ant-academy` `22df69dd`: **every in-session write is fired from `code/src/progress/store.js`**
> — `saveChapterProgress` (`:150`) calls the injected authed requester with `UPSERT_CHAPTER_PROGRESS` at `:162`,
> `saveLastActivity` (`:202`) with `SET_LAST_ACTIVITY` at `:210` — i.e. **straight to the supergraph** over the
> cross-origin GraphQL endpoint with a Clerk Bearer token. The beacon route is the **exception the old sentence
> presented as the rule**: an *on-unload last-ditch flush*, passed as the `beacon:` option at `store.js:169` /
> `:215` and reached only by `navigator.sendBeacon` / `fetch({keepalive:true})` on pagehide
> (`src/writeThrough/index.js:247`, `:259`). It exists precisely **because** `sendBeacon` cannot set an
> `Authorization` header, so this **same-origin** route re-issues the mutation server-side from the Clerk session
> cookie — its own header comment says so: *"a best-effort last-ditch flush for a write that would otherwise be
> lost if the tab closes mid-retry"* (`route.js:1-18`). The mutation NAMES are correct either way; only the
> attribution of where they are posted from was wrong.
```

**CITED CONTENT**

```
   244      const body = typeof desc.body === 'string' ? desc.body : JSON.stringify(desc.body ?? {})
   245      let sent = false
   246      try {
   247        if (typeof navigator !== 'undefined' && typeof navigator.sendBeacon === 'function') {
   248          sent = navigator.sendBeacon(desc.url, new Blob([body], { type: desc.contentType || 'application/json' }))
   249        }
   250      } catch {
```

## 01-026
- **id**: `B01-026`
- **corpus site**: `corpus/services/ant-academy.md:79-91` (paragraph)
- **citation**: `serverTenant.js:115-145`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/src/lib/serverTenant.js`  (337 lines)

**CLAIMING UNIT**

```md
This makes a "played academy session" a **seedable server row** (via `app/cmd/academy-seed`) — **on a
backend-wired deployment. That binary is MOOT on a demo stack** (M236 iter-08): a demo academy has no
`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`, so the backend read yields nothing and **nothing ever reads the seeded
`academy_chapter_progresses` rows**. Seeding them on a demo changes no pixel. **⚠️ It does *not* "fall back to
the committed FS catalog"** — there is **no FS-as-published fallback** in the app; `getServerCatalogView()`
resolves a null backend result to the **empty view**, and `serverTenant.js:115-145` says so in-code: *"the
cutover is intentional, not reversible-on-error."* A demo grid renders cards only because the rext demo-patch
**`demo-stack/patches/academy-fs-published-fallback`** *restores* that removed fallback on the demo's ephemeral
clone (applied by `demo-stack/ant-academy.sh` before `next dev`, env-gated on `ACADEMY_DEMO_FS_PUBLISHED`,
reverted on `--stop`); if that patch is refused, the grid is empty. See
[the empty-grid analysis below](#the-content-model--db-authoritative-catalog-v051-m7). The demo's academy
story is therefore **presence-only** — a real `/courses/<slug>` link into a grid of 65 real cards — not a
progress/result surface. See [`../ops/demo/content-stories-routes.md`](../ops/demo/content-stories-routes.md).
```

**CITED CONTENT**

```
   112   * tenant metadata and is the one piece not yet modeled in the backend catalog,
   113   * so it passes through verbatim (same as the backend adapter does).
   114   */
   115  function emptyCatalogView() {
   116    return { chapters: [], skillPaths: {}, series: [], bundles: PUBLIC_BUNDLES, catalogVersion: CATALOG_VERSION }
   117  }
   118  
   119  /**
   120   * The catalog the current (authenticated) user is allowed to receive: public +
   121   * the user's org tenants. This is what authed RSCs pass to the client instead
   122   * of importing the raw CHAPTERS array — so tenant content never ships to a
   123   * browser that can't see it.
   124   *
   125   * DB-authoritative (M7): the catalog is read UNCONDITIONALLY from the academy
   126   * backend (`getBackendCatalogView(eids)` → already tenant-filtered server-side,
   127   * with a chapter-level eid refinement in the adapter). DB presence gates
   128   * visibility — a path/chapter renders only if the backend returns it; there is
   129   * NO FS-as-published fallback. The eids are threaded in for the adapter's
   130   * chapter-level tenant refinement. A null backend result (not-composed /
   131   * outage) resolves to the EMPTY view rather than the committed FS catalog: the
   132   * cutover is intentional, not reversible-on-error.
   133   *
   134   * DEV DRAFT LAYER (M8): when `draftsEnabled()` (dev + the ACADEMY_SHOW_DRAFTS
   135   * opt-in; NEVER in production — hard-blocked), FS-only content (present-on-FS ∧
   136   * absent-from-DB) is merged on top of the DB view tagged `_draft: true`, so an
   137   * author previews locally-committed-but-not-yet-exported content. Off (incl. all
   138   * production) → the DB view passes through verbatim, zero behavior change. The
   139   * merge runs server-side; the client still receives only the threaded view
   140   * (catalog-client-boundary intact). The eids thread through so a tenant-scoped
   141   * draft stays tenant-gated.
   142   */
   143  export async function getServerCatalogView() {
   144    const eids = await getUserEids()
   145    const view = (await getBackendCatalogView(eids)) ?? (process.env.ACADEMY_DEMO_FS_PUBLISHED === '1' ? (v => ({ ...v, chapters: (v.chapters ?? []).map(({ _draft, _origin, ...c }) => c), series: (v.series ?? []).map(({ _draft, _origin, ...s }) => s), skillPaths: Object.fromEntries(Object
```

## 01-027
- **id**: `B01-027`
- **corpus site**: `corpus/services/ant-academy.md:152-154` (paragraph)
- **citation**: `code/src/lib/serverTenant.js:145`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/src/lib/serverTenant.js`  (337 lines)

**CLAIMING UNIT**

```md
`getServerCatalogView()` is literally `const view = (await getBackendCatalogView(eids)) ?? emptyCatalogView()`
(`code/src/lib/serverTenant.js:145` @ `ant-academy` `22df69dd` — byte-exact), so on **any** failure the catalog
becomes `emptyCatalogView()` → **0 cards**.
```

**CITED CONTENT**

```
   142   */
   143  export async function getServerCatalogView() {
   144    const eids = await getUserEids()
   145    const view = (await getBackendCatalogView(eids)) ?? (process.env.ACADEMY_DEMO_FS_PUBLISHED === '1' ? (v => ({ ...v, chapters: (v.chapters ?? []).map(({ _draft, _origin, ...c }) => c), series: (v.series ?? []).map(({ _draft, _origin, ...s }) => s), skillPaths: Object.fromEntries(Object.entries(v.skillPaths ?? {}).map(([k, { _draft, _origin, ...p }]) => [k, p])) }))(mergeDrafts(emptyCatalogView(), eids)) : emptyCatalogView())
   146    return draftsEnabled() ? mergeDrafts(view, eids) : view
   147  }
   148  
```

## 01-028
- **id**: `B01-028`
- **corpus site**: `corpus/services/ant-academy.md:156-166` (paragraph)
- **citation**: `serverTenant.js:115-117`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/src/lib/serverTenant.js`  (337 lines)

**CLAIMING UNIT**

```md
> **The shape of `emptyCatalogView()` is FIVE keys, not three** (corrected M257x iter-102 — this passage asserted
> it by `=` as the 3-key literal `{ chapters: [], skillPaths: {}, series: [] }`). Measured at
> `serverTenant.js:115-117`: `return { chapters: [], skillPaths: {}, series: [], bundles: PUBLIC_BUNDLES,
> catalogVersion: CATALOG_VERSION }`. The two extra keys are **not** empty — `PUBLIC_BUNDLES`
> (`code/ucourses/catalog.js:961`) is a populated exported array of curated bundle objects and `CATALOG_VERSION`
> (`:31`) is `'1.0'` — they pass through verbatim from the committed FS tree because, in the function's own
> words (`:111-113`), `bundles` *"carries no tenant metadata and is the one piece not yet modeled in the backend
> catalog."* **The `→ 0 cards` conclusion is unaffected and stands:** a bundle stripe's path cards are derived
> from the (empty) `chapters` — `AcademyClient.jsx:1363-1365` drops every path whose `scopedChapters` filter is
> empty — so the audience views render bundle chrome with **zero** path cards, and the grid itself renders none.
> Do not "fix" this by deleting the bundles key from the shape; the shape is what the code returns.
```

**CITED CONTENT**

```
   112   * tenant metadata and is the one piece not yet modeled in the backend catalog,
   113   * so it passes through verbatim (same as the backend adapter does).
   114   */
   115  function emptyCatalogView() {
   116    return { chapters: [], skillPaths: {}, series: [], bundles: PUBLIC_BUNDLES, catalogVersion: CATALOG_VERSION }
   117  }
   118  
   119  /**
   120   * The catalog the current (authenticated) user is allowed to receive: public +
```

## 01-029
- **id**: `B01-029`
- **corpus site**: `corpus/services/ant-academy.md:156-166` (paragraph)
- **citation**: `code/ucourses/catalog.js:961`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/ucourses/catalog.js`  (1361 lines)

**CLAIMING UNIT**

```md
> **The shape of `emptyCatalogView()` is FIVE keys, not three** (corrected M257x iter-102 — this passage asserted
> it by `=` as the 3-key literal `{ chapters: [], skillPaths: {}, series: [] }`). Measured at
> `serverTenant.js:115-117`: `return { chapters: [], skillPaths: {}, series: [], bundles: PUBLIC_BUNDLES,
> catalogVersion: CATALOG_VERSION }`. The two extra keys are **not** empty — `PUBLIC_BUNDLES`
> (`code/ucourses/catalog.js:961`) is a populated exported array of curated bundle objects and `CATALOG_VERSION`
> (`:31`) is `'1.0'` — they pass through verbatim from the committed FS tree because, in the function's own
> words (`:111-113`), `bundles` *"carries no tenant metadata and is the one piece not yet modeled in the backend
> catalog."* **The `→ 0 cards` conclusion is unaffected and stands:** a bundle stripe's path cards are derived
> from the (empty) `chapters` — `AcademyClient.jsx:1363-1365` drops every path whose `scopedChapters` filter is
> empty — so the audience views render bundle chrome with **zero** path cards, and the grid itself renders none.
> Do not "fix" this by deleting the bundles key from the shape; the shape is what the code returns.
```

**CITED CONTENT**

```
   958  // The START HERE pill is rendered by the AUDIENCE view renderer for the
   959  // FIRST stripe in each audience's AUDIENCE_BUNDLE_ORDER — not a per-bundle
   960  // flag — so each audience gets exactly one canonical entry point.
   961  export const PUBLIC_BUNDLES = [
   962      // ─── For Beginners ──────────────────────────────────────────
   963      {
   964          key: 'beginner-track',
```

## 01-030
- **id**: `B01-030`
- **corpus site**: `corpus/services/ant-academy.md:260-286` (bullet)
- **citation**: `src/i18n/locale.js:10`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/src/i18n/locale.js`  (84 lines)

**CLAIMING UNIT**

```md
- **There are TWO switchers, and only one of them is the 2-way toggle.** `src/i18n/LocaleSwitch.jsx` **is** a 2-way
  EN↔IT toggle `<Link>` that sets `?lang=it` on the current path — and it is mounted **only in the public-storefront
  header** (`src/views/public/PublicHeader.jsx:20`). **That header renders on `/library` ALONE** — corrected M257x
  iter-102, which struck the trailing two-route gloss this line used to carry: measured @ `22df69dd`,
  `PublicHeader` has exactly one mount site, `code/app/(public)/library/page.jsx:28`, and `/free` renders no
  header of its own — its whole body is `redirect('/?tier=free')` (`code/app/(public)/free/page.jsx:18`), which
  lands on the **app-shell** home and therefore serves the *other* switcher. (The *"only in the public-storefront
  header"* half was true and is kept.)
  The app-shell header mounts a
  different component, and that one **IS a dropdown menu**: `src/components/LanguageSelector.jsx`
  (`src/components/TopBar.jsx:76`) renders a flag-only trigger opening a `role="menu"` panel
  (`LanguageSelector.jsx:88`) of **7** `role="menuitemradio"` options (`:97`) over
  `SUPPORTED_LOCALES = ['en', 'it', 'es', 'fr', 'de', 'nl', 'pt']`
  (`src/i18n/locale.js:10`) — and **`TopBar`'s surface set is SEVEN routes, not five** (also corrected at iter-102;
  this line named `/`, `/chapters/*`, `/latest`, `/bookmarks`, `/my-activity` and closed the list). Measured
  @ `22df69dd` by mount site: **`/`, `/courses`, `/courses/[slug]`** — all three render `AcademyClient`
  (`app/(authed)/page.jsx:151`, `courses/page.jsx:92`, `courses/[slug]/page.jsx:219`), whose `TopBar` is at
  `AcademyClient.jsx:1906` — plus **`/chapters/[slug]`** (`CourseClient.jsx:2091`, `:2141`), **`/latest`**
  (`LatestClient.jsx:128`), **`/bookmarks`** (`BookmarksClient.jsx:508`) and **`/my-activity`**
  (`MyActivityClient.jsx:161`). `/my-certificates` renders **no** `TopBar` (`MyCertificatesClient.jsx` imports
  none). The two routes the old list omitted — `/courses` and `/courses/[slug]` — are precisely the demo's landing
  routes. **So "the switcher shows no menu" cannot be dismissed as "there is no menu."** This bullet read *"not a dropdown
  menu"* until M257x; that half is **retracted**, and it was wrong **when written**, not merely stale — the dropdown
  has been mounted in `TopBar` since `5b05b7d9` (2026-05-05) and 7-locale since `e22f3230` (2026-
```

**CITED CONTENT**

```
     7  // Hard-refreshing a bare URL therefore lands the user in DEFAULT_LOCALE
     8  // (en) — that is by design.
     9  
    10  export const SUPPORTED_LOCALES = ['en', 'it', 'es', 'fr', 'de', 'nl', 'pt']
    11  export const DEFAULT_LOCALE = 'en'
    12  
    13  export const LOCALE_LABELS = {
```

## 01-031
- **id**: `B01-031`
- **corpus site**: `corpus/services/ant-academy.md:364-387` (paragraph)
- **citation**: `demo-stack/ant-academy.sh:576-583`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/demo-stack/ant-academy.sh`  (718 lines)

**CLAIMING UNIT**

```md
> **In a demo the academy is AUTHENTICATED-as-a-member via real Clerkenstein keys (v2.3 M220 S5/i).**
> ⚠ **This block described the OPPOSITE model until M257x iter-98, and was self-contradictory besides** — it
> opened by saying the cockpit sets the `e2e_persona` cookie, denied it mid-paragraph, then asserted it again
> in the closing line. Measured at rext `main`, both halves are now the other way round:
>
> * **The `e2e_persona` BYPASS is gone from the academy's launch env.** `/demo-up` sets **neither**
>   `BENCHMARK_VISUAL_BYPASS=1` nor `NEXT_PUBLIC_E2E_AUTH=1` (`demo-stack/ant-academy.sh:576-583`), because the
>   demo academy now gets real Clerkenstein keys. Keeping both would be *worse* than either: `proxy.js`
>   short-circuits on the persona cookie **before** it resolves the real session, so a presenter logged in as
>   Maya would be rendered a generic **"E2E Member"**. Fenced by two tests
>   (`test_the_e2e_persona_bypass_is_gone_from_the_launch_env`, `test_e2e_persona_bypass_is_not_in_the_launch_env`).
> * **The cockpit DOES still set the cookie — at two live paths**, contrary to the retracted sentence:
>   client-side on click (`demo-stack/cockpit.py:812`, `_ACADEMY_JS`) and as a `Set-Cookie` on the `/go` 302
>   (`:1496`); `:327` names both. It is the content-stories academy deep-link that uses them. With the launch-env
>   bypass gone the cookie is **inert** on a stock demo — set, and not honoured.
>
> The historical keyless model (server RSC `anonymous=false` + a synthetic **`E2E Member`**, no real Clerk keys)
> is what v1.10b M53 F6 shipped and is **no longer how a demo runs**. (Separately, and still true:
> the academy grid renders **empty** in a demo — the v2.4 **F4** carry, **NOT** a client-side render defect: the
> catalog is [DB-authoritative](#the-content-model--db-authoritative-catalog-v051-m7) and the demo neither sets
> `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` nor holds academy rows → `emptyCatalogView()`. v2.5 **M230** fills it
> production-faithfully, zero academy-repo edits). The **Cosmo AI assistant stays absent** in a demo (its flag +
> OpenAI key are deliberately not provisioned — the AI-keys policy). Zero academy-repo edits: the keys live in
> the gitignored `code/.env.local`. Full mechanics: [`../ops/demo/frontend-tier.md` § ant-academy](../ops/demo/frontend-tier.md).
```

**CITED CONTENT**

```
   573  else
   574    log "⚠ ant-academy-back-to-cockpit patch skipped/refused (non-fatal) — the academy user menu will have no 'Back to Cockpit' item"
   575  fi
   576  # M220 (S5/i): the e2e_persona bypass (BENCHMARK_VISUAL_BYPASS + NEXT_PUBLIC_E2E_AUTH) is GONE from the
   577  # launch env. It existed to fake an authenticated academy session on a KEYLESS academy — a workaround for
   578  # the very wiring gap this milestone closes. Keeping it alongside real Clerkenstein keys would be actively
   579  # worse than either alone: `proxy.js` short-circuits on the persona cookie BEFORE it resolves the real
   580  # session, so the academy would render a generic **"E2E Member"** to a presenter logged in as Maya — the
   581  # persona self-consistency defect the coverage protocol exists to catch, shipped by our own launcher.
   582  # With a real (fake-)Clerk instance the hero is genuinely signed in, so the mock persona has nothing left
   583  # to do. NODE_ENV stays `development` because this is `next dev`.
   584  ( cd "$ACADEMY" && launch_detached "$PIDFILE" -- \
   585      env "${devorigin_env[@]+"${devorigin_env[@]}"}" "${fspub_env[@]+"${fspub_env[@]}"}" NODE_ENV=development npm run dev -- --port "$PORT" "${bind_args[@]+"${bind_args[@]}"}" \
   586    ) </dev/null >>"$LOG" 2>&1
```

## 01-032
- **id**: `B01-032`
- **corpus site**: `corpus/services/ant-academy.md:364-387` (paragraph)
- **citation**: `demo-stack/cockpit.py:812`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/demo-stack/cockpit.py`  (1756 lines)

**CLAIMING UNIT**

```md
> **In a demo the academy is AUTHENTICATED-as-a-member via real Clerkenstein keys (v2.3 M220 S5/i).**
> ⚠ **This block described the OPPOSITE model until M257x iter-98, and was self-contradictory besides** — it
> opened by saying the cockpit sets the `e2e_persona` cookie, denied it mid-paragraph, then asserted it again
> in the closing line. Measured at rext `main`, both halves are now the other way round:
>
> * **The `e2e_persona` BYPASS is gone from the academy's launch env.** `/demo-up` sets **neither**
>   `BENCHMARK_VISUAL_BYPASS=1` nor `NEXT_PUBLIC_E2E_AUTH=1` (`demo-stack/ant-academy.sh:576-583`), because the
>   demo academy now gets real Clerkenstein keys. Keeping both would be *worse* than either: `proxy.js`
>   short-circuits on the persona cookie **before** it resolves the real session, so a presenter logged in as
>   Maya would be rendered a generic **"E2E Member"**. Fenced by two tests
>   (`test_the_e2e_persona_bypass_is_gone_from_the_launch_env`, `test_e2e_persona_bypass_is_not_in_the_launch_env`).
> * **The cockpit DOES still set the cookie — at two live paths**, contrary to the retracted sentence:
>   client-side on click (`demo-stack/cockpit.py:812`, `_ACADEMY_JS`) and as a `Set-Cookie` on the `/go` 302
>   (`:1496`); `:327` names both. It is the content-stories academy deep-link that uses them. With the launch-env
>   bypass gone the cookie is **inert** on a stock demo — set, and not honoured.
>
> The historical keyless model (server RSC `anonymous=false` + a synthetic **`E2E Member`**, no real Clerk keys)
> is what v1.10b M53 F6 shipped and is **no longer how a demo runs**. (Separately, and still true:
> the academy grid renders **empty** in a demo — the v2.4 **F4** carry, **NOT** a client-side render defect: the
> catalog is [DB-authoritative](#the-content-model--db-authoritative-catalog-v051-m7) and the demo neither sets
> `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` nor holds academy rows → `emptyCatalogView()`. v2.5 **M230** fills it
> production-faithfully, zero academy-repo edits). The **Cosmo AI assistant stays absent** in a demo (its flag +
> OpenAI key are deliberately not provisioned — the AI-keys policy). Zero academy-repo edits: the keys live in
> the gitignored `code/.env.local`. Full mechanics: [`../ops/demo/frontend-tier.md` § ant-academy](../ops/demo/frontend-tier.md).
```

**CITED CONTENT**

```
   809        var persona = this.getAttribute('data-academy-persona') || 'member';
   810        // Host-only cookie for localhost (no Domain attr → the current host; shared across ports).
   811        // path=/ so every academy route sees it; SameSite=Lax so the top-level navigation carries it.
   812        try { document.cookie = 'e2e_persona=' + encodeURIComponent(persona) + '; path=/; SameSite=Lax'; } catch (e) {}
   813      });
   814    }
   815  })();
```

## 01-033
- **id**: `B01-033`
- **corpus site**: `corpus/services/ant-academy.md:474-474` (bullet)
- **citation**: `code/src/lib/platformUrls.js:1-32`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/src/lib/platformUrls.js`  (87 lines)

**CLAIMING UNIT**

```md
- **Clerk (shared)**: Uses the same Clerk app as the rest of the platform. ⚠️ **"Domain-gated to `@anthropos.work` so external users cannot enter" is FALSE and was removed at M257x iter-115.** Measured at `ant-academy` `22df69dd`: `code/src/lib/platformUrls.js:1-32` calls the Academy *"a **storefront** in front of the Anthropos platform"* and defines **FLOW A — Account gate** (*"the PRIMARY CTA for this flow — **a new visitor registers**"*) and **FLOW B — Checkout gate** (*"contextually registers → signs in → subscribes **for an anonymous visitor**"*); `code/src/lib/pricing.js` sets `STANDARD_YEARLY = { usd: 399, eur: 349 }` with a live launch promo and coupon codes; `code/src/components/TopBar.jsx:77-88` renders a *"Buy AI Academy"* CTA **to anonymous visitors** and navigates them to platform checkout; `code/src/lib/schema.js:3` publishes `SITE_URL = 'https://aiacademy.anthropos.work'` with an `EducationalOrganization` schema.org block and public SEO copy decks. The repo's own `knowledge/user-types.md` enumerates four user types — Anonymous · Signed-in (free) · Subscriber · Enterprise/Org member — and **no `@anthropos.work` predicate appears in the detection list at all**. A product that sells a $399/yr subscription to an anonymous visitor through a public checkout funnel is not one external users cannot enter. **This document contradicted itself 213 lines earlier**, where it documents anonymous browsing, a *"Phase-1 public launch"*, *"the public surface is much wider than 'a few auth pages'"*, and the literal phrase **"public-storefront"**. (The error is inherited rather than invented — the ant-academy repo's own `CLAUDE.md:11` still says *"internal learning portal for Anthropic employees"* — but the corpus stated it in its own voice, present tense, with no attribution.) **What is true:** Clerk is shared with the platform, and the org-membership gate still applies to the *enterprise* surfaces.
```

**CITED CONTENT**

```
     1  // Canonical platform conversion URLs — the single source of truth for the two
     2  // Academy → platform conversion flows. See knowledge/onboarding-and-conversion.md.
     3  //
     4  // The Academy is a storefront in front of the Anthropos platform
     5  // (`app.anthropos.work`, overridable via NEXT_PUBLIC_STUDIO_URL). Account +
     6  // billing live on the platform, so every conversion CTA leaves the Academy and
     7  // lands on a platform page:
     8  //
     9  //   FLOW A — Account gate (free content, login required)
    10  //     platformSignUpUrl(returnTo) → `${STUDIO}/sign-up?redirect_url=<returnTo>`
    11  //     The PRIMARY CTA for this flow — a new visitor registers, then is
    12  //     redirected back to `returnTo`. Maps to lockReason 'signup'.
    13  //     platformLoginUrl(returnTo)  → `${STUDIO}/login?redirect_url=<returnTo>`
    14  //     The existing-account branch ("Log in" / "Already have an account? Sign
    15  //     in"). Both pages accept `redirect_url`; pick by what the CTA says.
    16  //
    17  //   FLOW B — Checkout gate (paid content: subscription or enterprise membership)
    18  //     platformCheckoutUrl(returnTo) → `${STUDIO}/checkout?post_checkout_url=<returnTo>`
    19  //     Checkout contextually registers → signs in → subscribes for an anonymous
    20  //     visitor (and signs in an existing account that just needs to pay). Maps to
    21  //     lockReason 'subscribe'. The param is `post_checkout_url`, NOT `redirect_url`:
    22  //     the checkout page embeds Clerk's <SignIn>, and Clerk gives a `redirect_url`
    23  //     query param precedence over the page's own redirect props — so users were
    24  //     bounced back here right after signup, BEFORE the payment step. The platform
    25  //     captures `post_checkout_url` instead (allowlisted to *.anthropos.work) and
    26  //     redirects to it only once checkout is actually complete — payment done, or
    27  //     the visitor turns out to be already entitled (premium / org member).
    28  //
    29  // Both destinations let an existing-account holder sign IN instead of
    30  // registering — there is no separate "sign in" URL; /login and /checkout each
    31  // handle it. `process.env.NEXT_PUBLIC_STUDIO_URL` is inlined by Next at build
    32  // time, so this module is safe to import from client components.
    33  
```

## 01-034
- **id**: `B01-034`
- **corpus site**: `corpus/services/ant-academy.md:474-474` (bullet)
- **citation**: `code/src/lib/schema.js:3`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/src/lib/schema.js`  (61 lines)

**CLAIMING UNIT**

```md
- **Clerk (shared)**: Uses the same Clerk app as the rest of the platform. ⚠️ **"Domain-gated to `@anthropos.work` so external users cannot enter" is FALSE and was removed at M257x iter-115.** Measured at `ant-academy` `22df69dd`: `code/src/lib/platformUrls.js:1-32` calls the Academy *"a **storefront** in front of the Anthropos platform"* and defines **FLOW A — Account gate** (*"the PRIMARY CTA for this flow — **a new visitor registers**"*) and **FLOW B — Checkout gate** (*"contextually registers → signs in → subscribes **for an anonymous visitor**"*); `code/src/lib/pricing.js` sets `STANDARD_YEARLY = { usd: 399, eur: 349 }` with a live launch promo and coupon codes; `code/src/components/TopBar.jsx:77-88` renders a *"Buy AI Academy"* CTA **to anonymous visitors** and navigates them to platform checkout; `code/src/lib/schema.js:3` publishes `SITE_URL = 'https://aiacademy.anthropos.work'` with an `EducationalOrganization` schema.org block and public SEO copy decks. The repo's own `knowledge/user-types.md` enumerates four user types — Anonymous · Signed-in (free) · Subscriber · Enterprise/Org member — and **no `@anthropos.work` predicate appears in the detection list at all**. A product that sells a $399/yr subscription to an anonymous visitor through a public checkout funnel is not one external users cannot enter. **This document contradicted itself 213 lines earlier**, where it documents anonymous browsing, a *"Phase-1 public launch"*, *"the public surface is much wider than 'a few auth pages'"*, and the literal phrase **"public-storefront"**. (The error is inherited rather than invented — the ant-academy repo's own `CLAUDE.md:11` still says *"internal learning portal for Anthropic employees"* — but the corpus stated it in its own voice, present tense, with no attribution.) **What is true:** Clerk is shared with the platform, and the org-membership gate still applies to the *enterprise* surfaces.
```

**CITED CONTENT**

```
     1  import { BRAND } from './brand.js';
     2  
     3  const SITE_URL = 'https://aiacademy.anthropos.work';
     4  
     5  export function organizationSchema() {
     6    return {
```

## 01-035
- **id**: `B01-035`
- **corpus site**: `corpus/services/ant-academy.md:474-474` (bullet)
- **citation**: `CLAUDE.md:11`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/CLAUDE.md`  (581 lines)

**CLAIMING UNIT**

```md
- **Clerk (shared)**: Uses the same Clerk app as the rest of the platform. ⚠️ **"Domain-gated to `@anthropos.work` so external users cannot enter" is FALSE and was removed at M257x iter-115.** Measured at `ant-academy` `22df69dd`: `code/src/lib/platformUrls.js:1-32` calls the Academy *"a **storefront** in front of the Anthropos platform"* and defines **FLOW A — Account gate** (*"the PRIMARY CTA for this flow — **a new visitor registers**"*) and **FLOW B — Checkout gate** (*"contextually registers → signs in → subscribes **for an anonymous visitor**"*); `code/src/lib/pricing.js` sets `STANDARD_YEARLY = { usd: 399, eur: 349 }` with a live launch promo and coupon codes; `code/src/components/TopBar.jsx:77-88` renders a *"Buy AI Academy"* CTA **to anonymous visitors** and navigates them to platform checkout; `code/src/lib/schema.js:3` publishes `SITE_URL = 'https://aiacademy.anthropos.work'` with an `EducationalOrganization` schema.org block and public SEO copy decks. The repo's own `knowledge/user-types.md` enumerates four user types — Anonymous · Signed-in (free) · Subscriber · Enterprise/Org member — and **no `@anthropos.work` predicate appears in the detection list at all**. A product that sells a $399/yr subscription to an anonymous visitor through a public checkout funnel is not one external users cannot enter. **This document contradicted itself 213 lines earlier**, where it documents anonymous browsing, a *"Phase-1 public launch"*, *"the public surface is much wider than 'a few auth pages'"*, and the literal phrase **"public-storefront"**. (The error is inherited rather than invented — the ant-academy repo's own `CLAUDE.md:11` still says *"internal learning portal for Anthropic employees"* — but the corpus stated it in its own voice, present tense, with no attribution.) **What is true:** Clerk is shared with the platform, and the org-membership gate still applies to the *enterprise* surfaces.
```

**CITED CONTENT**

```
     8  1. **Documentation Repository**: Comprehensive architecture guides for developers
     9  2. **Environment Setup**: Manual for humans and AI agents to build local development environments
    10  3. **Recursive Inspection**: Tool for reverse-engineering and documenting the platform itself
    11  
    12  This is NOT the Anthropos platform source code - it's the documentation about it. The actual platform code lives in separate repositories under the `anthropos-work` GitHub organization.
    13  
    14  ## Development Commands
```

## 01-036
- **id**: `B01-036`
- **corpus site**: `corpus/services/askengine.md:88-90` (bullet)
- **citation**: `main.go:467-471`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
*   **Prerequisites**: runs **as part of `app`** (no dedicated `cmd/`). ⚠️ **A failed Bedrock init takes the WHOLE
    `backend` process down — nothing is "disabled".** Measured at `app` `ad9f3c49` (`== origin/main`, 2026-08-06),
    `main.go:467-471`:
```

**CITED CONTENT**

```
   464  	// manager through the adapter so preview == what actually sends.
   465  	aiReadinessPreviewRenderer := emailpreview.NewRenderer(aiReadinessManager, orgManager, aireadiness.PreviewOverrideLookup{Mgr: aiReadinessOverrides})
   466  
   467  	bedrockClient, err := askengine.NewBedrockClient(serverContext)
   468  	if err != nil {
   469  		logger.Error("bedrock client unavailable; talk-to-data disabled", "error", err)
   470  		return
   471  	}
   472  
   473  	// Managers are constructed here in the root and assembled into the App data
   474  	// struct below — there is no functional-options builder anymore (the app.With*
```

## 01-037
- **id**: `B01-037`
- **corpus site**: `corpus/services/askengine.md:100-114` (paragraph)
- **citation**: `main.go:229`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
    That `return` sits at **one level inside `func main()`** (`main.go:229`; no other `func` declaration falls
    between `:229` and `:467`), so it returns from `main` itself — a *normal* return, hence exit status 0, which
    is why it reads as a clean shutdown rather than a crash. Everything constructed **after** it never runs: the
    Connect-RPC mux (`:1295`), the meta HTTP server (`:1361`), the Echo router, the Asynq pools and the Redis
    subscribers (`:1438` onward). **The trap is the platform's own log string** — `"bedrock client unavailable;
    talk-to-data disabled"` occurs **exactly once** in `app`'s Go source, and this doc had read it as a
    description of behaviour rather than of intent (retracted M257x iter-102).
    **How narrow the path is, stated rather than smoothed:** `askengine.NewBedrockClient`
    (`internal/askengine/bedrock.go:161`) has exactly **one** error return — `config.LoadDefaultConfig` failing
    (a malformed AWS config/profile), not merely absent credentials, which the SDK resolves lazily at call time.
    So a bare local shell with no AWS creds normally boots fine and fails *per request*; but **when** the
    constructor does error, the consequence is process exit, not a disabled subsystem.
    **Beyond that**, it needs AWS Bedrock access (IAM via the default credential chain) with Claude Sonnet 4.6
    enabled in **eu-west-1**, a Postgres with the platform schema, and an embeddings provider for the shared `ai`
    client.
```

**CITED CONTENT**

```
   226  	return dsn + " statement_timeout=" + ms
   227  }
   228  
   229  func main() {
   230  	serviceName := os.Getenv("SERVICE_NAME")
   231  	if serviceName == "" {
   232  		serviceName = "backend"
```

## 01-038
- **id**: `B01-038`
- **corpus site**: `corpus/services/askengine.md:100-114` (paragraph)
- **citation**: `internal/askengine/bedrock.go:161`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/askengine/bedrock.go`  (787 lines)

**CLAIMING UNIT**

```md
    That `return` sits at **one level inside `func main()`** (`main.go:229`; no other `func` declaration falls
    between `:229` and `:467`), so it returns from `main` itself — a *normal* return, hence exit status 0, which
    is why it reads as a clean shutdown rather than a crash. Everything constructed **after** it never runs: the
    Connect-RPC mux (`:1295`), the meta HTTP server (`:1361`), the Echo router, the Asynq pools and the Redis
    subscribers (`:1438` onward). **The trap is the platform's own log string** — `"bedrock client unavailable;
    talk-to-data disabled"` occurs **exactly once** in `app`'s Go source, and this doc had read it as a
    description of behaviour rather than of intent (retracted M257x iter-102).
    **How narrow the path is, stated rather than smoothed:** `askengine.NewBedrockClient`
    (`internal/askengine/bedrock.go:161`) has exactly **one** error return — `config.LoadDefaultConfig` failing
    (a malformed AWS config/profile), not merely absent credentials, which the SDK resolves lazily at call time.
    So a bare local shell with no AWS creds normally boots fine and fails *per request*; but **when** the
    constructor does error, the consequence is process exit, not a disabled subsystem.
    **Beyond that**, it needs AWS Bedrock access (IAM via the default credential chain) with Claude Sonnet 4.6
    enabled in **eu-west-1**, a Postgres with the platform schema, and an embeddings provider for the shared `ai`
    client.
```

**CITED CONTENT**

```
   158  
   159  // NewBedrockClient builds a client using the default AWS credential chain
   160  // (env, shared config, IAM role) and routes requests through Bedrock.
   161  func NewBedrockClient(ctx context.Context) (*BedrockClient, error) {
   162  	region := envOr("AWS_REGION", DefaultRegion)
   163  	modelID := envOr("ASK_MODEL_ID", DefaultModelID)
   164  
```

## 01-039
- **id**: `B01-039`
- **corpus site**: `corpus/services/backend.md:3-99` (paragraph)
- **citation**: `main.go:1297-1338`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> ## `app` is the backend monolith
>
> **Eight former microservices now run inside `app`**, in merge order:
>
> | Merged service | Program | What moved in |
> |---|---|---|
> | [skiller](./skiller.md) | skiller-in-app (v2.1 "quick change", July 2026) | the skills-taxonomy graph (**≥42,790 skills / ≥22,470 job roles** — public subset; [not "60K/18K"](../architecture/shared_libraries.md#taxonomy-figures)), embeddings, AI matching |
> | [skillpath](./skillpath.md) | skillpath-in-app (M502→M507) | skill-path progression engine, session state |
> | [roadrunner](./roadrunner.md) | with jobsim-in-app | Judge0 code execution (called directly via `JUDGE0_BASE_URL`) |
> | [jobsimulation](./jobsimulation.md) | jobsim-in-app (prod ECS teardown **M810 — LANDED**, `6092c6d2`) | the simulation session engine — `internal/jobsimulation/`, wired by `internal/jobsimwiring/wiring.go` |
> | [cms](./cms.md) | cms-in-app v8.0, app **v1.360.0** (prod teardown **M810 — NOT MEASURABLE here**; report both, assert neither — see the *M810 prod teardown is UNEVEN* bullet below) | content layer + Directus edge + Studio — `internal/cms/` |
> | [storage](./storage.md) | v9.0 "support-in-app", 2026-08-04 | the private + public object-storage managers — `internal/storage/`, `internal/storagens/`, `internal/publicstorage/` |
> | [messenger](./messenger.md) | v9.0 "support-in-app", 2026-08-04 | transactional email (Brevo + Liquid) and messenger's **own** Redis consumer group — `internal/messenger/`; switch-gated by `MESSENGER_ENABLED` |
> | [customerio-sync](./customerio-sync.md) | v9.0 "support-in-app" | the one-way Brevo marketing-contact push — `internal/customeriosync/`; switch-gated by `CUSTOMERIO_SYNC_ENABLED` |
>
> The last three lost their **containers** a day later, at platform `838d907` (merged `0c91421`,
> 2026-08-05): compose now declares **five** services and `repos.yml` **four** entries.
>
> Consequences that hold platform-wide:
> * **The federation composes ONE subgraph** (`backend`). cms-in-app was the **3 → 1** step: the single
>   commit `graphql-wundergraph@915da06` (2026-07-29) deleted **both** `schemas/cms.graphqls` **and**
>   `schemas/jobsimulation.graphqls`, taking the supergraph from (backend, jobsimulation, cms) to
>   (backend) alone. The jobsimulation subgraph therefore **survived jobsim-in-app** and was removed
>   here, not at its own merge.
> * **All of their tabl
```

**CITED CONTENT**

```
  1294  	// RPC
  1295  	mux := http.NewServeMux()
  1296  	// Reuse the same handlers the jobsim engine calls in-process (built above).
  1297  	mux.Handle(usersv1connect.NewUsersServiceHandler(usersRPCHandler))
  1298  	mux.Handle(organizationsv1connect.NewOrganizationsServiceHandler(orgsRPCHandler))
  1299  	// skiller-in-app M206: the in-app edge for the former standalone skiller
  1300  	// service. Serves the 5 externally-reached methods (GetSkill/GetSkills/
  1301  	// SearchSkill/MatchSkill/GetJobRole) from the ported managers; the 8
  1302  	// backend-only methods return CodeUnimplemented (internalized per D5,
  1303  	// callers cut over to managers in M208). Same service path skiller served,
  1304  	// so external consumers (jobsimulation/cms/messenger) only need to re-point
  1305  	// their base URL at cutover.
  1306  	mux.Handle(skillerv1connect.NewSkillerServiceHandler(
  1307  		skillerInProcess, // jobsim-in-app: the same in-process skiller handler built above (also the jobsim validator's collaborator)
  1308  	))
  1309  	// jobsim-in-app: the in-app JobSimulationService edge — serves the externally-reached subset
  1310  	// (messenger's GetSession/GetSessionInsights + cms's GetUserSessions) from the ported managers, with
  1311  	// jobsim's source error contract (CodeNotFound etc.). This IS the jobsim service now — the standalone
  1312  	// is decommissioned and the internalized methods return CodeUnimplemented. SaveInboundEmail is a no-op
  1313  	// stub (internal/inbound not ported, OD-11).
  1314  	mux.Handle(jobsimulationv1connect.NewJobSimulationServiceHandler(jobsimDj.RPCServer))
  1315  	// cms-in-app M807: the in-app CMSService edge — serves the externally-reached subset
  1316  	// (messenger's GetSimulation/GetSkillPath/GetEmailNotifications) from the ported cms
  1317  	// managers, with cms's source error contract. Additive + DORMANT: external callers
  1318  	// (messenger) keep hitting the standalone cms via CMS_RPC_ADDR until the M809 re-point;
  1319  	// registered only when the Directus edge is configured (cmsManagers built). LibraryExport
  1320  	// is a dead RPC (dropped, OD-10 — returns CodeUnimplemented from the ported server);
  1321  	// SaveHeyGenGeneration collapses to the in-process webhook path at M809 (OD-9).
  1322  	if cmsRPCServer != nil {
  1323  		mux.Handle(cmsv1connect.NewCMSServiceHandler(cmsRPC
```

## 01-040
- **id**: `B01-040`
- **corpus site**: `corpus/services/backend.md:3-99` (paragraph)
- **citation**: `cmd/root.go:62-66`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/cmd/root.go`  (287 lines)

**CLAIMING UNIT**

```md
> ## `app` is the backend monolith
>
> **Eight former microservices now run inside `app`**, in merge order:
>
> | Merged service | Program | What moved in |
> |---|---|---|
> | [skiller](./skiller.md) | skiller-in-app (v2.1 "quick change", July 2026) | the skills-taxonomy graph (**≥42,790 skills / ≥22,470 job roles** — public subset; [not "60K/18K"](../architecture/shared_libraries.md#taxonomy-figures)), embeddings, AI matching |
> | [skillpath](./skillpath.md) | skillpath-in-app (M502→M507) | skill-path progression engine, session state |
> | [roadrunner](./roadrunner.md) | with jobsim-in-app | Judge0 code execution (called directly via `JUDGE0_BASE_URL`) |
> | [jobsimulation](./jobsimulation.md) | jobsim-in-app (prod ECS teardown **M810 — LANDED**, `6092c6d2`) | the simulation session engine — `internal/jobsimulation/`, wired by `internal/jobsimwiring/wiring.go` |
> | [cms](./cms.md) | cms-in-app v8.0, app **v1.360.0** (prod teardown **M810 — NOT MEASURABLE here**; report both, assert neither — see the *M810 prod teardown is UNEVEN* bullet below) | content layer + Directus edge + Studio — `internal/cms/` |
> | [storage](./storage.md) | v9.0 "support-in-app", 2026-08-04 | the private + public object-storage managers — `internal/storage/`, `internal/storagens/`, `internal/publicstorage/` |
> | [messenger](./messenger.md) | v9.0 "support-in-app", 2026-08-04 | transactional email (Brevo + Liquid) and messenger's **own** Redis consumer group — `internal/messenger/`; switch-gated by `MESSENGER_ENABLED` |
> | [customerio-sync](./customerio-sync.md) | v9.0 "support-in-app" | the one-way Brevo marketing-contact push — `internal/customeriosync/`; switch-gated by `CUSTOMERIO_SYNC_ENABLED` |
>
> The last three lost their **containers** a day later, at platform `838d907` (merged `0c91421`,
> 2026-08-05): compose now declares **five** services and `repos.yml` **four** entries.
>
> Consequences that hold platform-wide:
> * **The federation composes ONE subgraph** (`backend`). cms-in-app was the **3 → 1** step: the single
>   commit `graphql-wundergraph@915da06` (2026-07-29) deleted **both** `schemas/cms.graphqls` **and**
>   `schemas/jobsimulation.graphqls`, taking the supergraph from (backend, jobsimulation, cms) to
>   (backend) alone. The jobsimulation subgraph therefore **survived jobsim-in-app** and was removed
>   here, not at its own merge.
> * **All of their tabl
```

**CITED CONTENT**

```
    59  )
    60  
    61  // rootCmd represents the base command when called without any subcommands
    62  var rootCmd = &cobra.Command{
    63  	Use:   "cms",
    64  	Short: "CMS Service",
    65  	RunE: func(cmd *cobra.Command, args []string) error {
    66  		environment := colony.ReadFromEnvVar()
    67  		logger := colony.InitLogger(versionConfig.Name,
    68  			colony.LoggingConfig{
    69  				SentryDSN:   os.Getenv("SENTRY_DSN"),
```

## 01-041
- **id**: `B01-041`
- **corpus site**: `corpus/services/backend.md:3-99` (paragraph)
- **citation**: `cmd/root.go:87`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/cmd/root.go`  (287 lines)

**CLAIMING UNIT**

```md
> ## `app` is the backend monolith
>
> **Eight former microservices now run inside `app`**, in merge order:
>
> | Merged service | Program | What moved in |
> |---|---|---|
> | [skiller](./skiller.md) | skiller-in-app (v2.1 "quick change", July 2026) | the skills-taxonomy graph (**≥42,790 skills / ≥22,470 job roles** — public subset; [not "60K/18K"](../architecture/shared_libraries.md#taxonomy-figures)), embeddings, AI matching |
> | [skillpath](./skillpath.md) | skillpath-in-app (M502→M507) | skill-path progression engine, session state |
> | [roadrunner](./roadrunner.md) | with jobsim-in-app | Judge0 code execution (called directly via `JUDGE0_BASE_URL`) |
> | [jobsimulation](./jobsimulation.md) | jobsim-in-app (prod ECS teardown **M810 — LANDED**, `6092c6d2`) | the simulation session engine — `internal/jobsimulation/`, wired by `internal/jobsimwiring/wiring.go` |
> | [cms](./cms.md) | cms-in-app v8.0, app **v1.360.0** (prod teardown **M810 — NOT MEASURABLE here**; report both, assert neither — see the *M810 prod teardown is UNEVEN* bullet below) | content layer + Directus edge + Studio — `internal/cms/` |
> | [storage](./storage.md) | v9.0 "support-in-app", 2026-08-04 | the private + public object-storage managers — `internal/storage/`, `internal/storagens/`, `internal/publicstorage/` |
> | [messenger](./messenger.md) | v9.0 "support-in-app", 2026-08-04 | transactional email (Brevo + Liquid) and messenger's **own** Redis consumer group — `internal/messenger/`; switch-gated by `MESSENGER_ENABLED` |
> | [customerio-sync](./customerio-sync.md) | v9.0 "support-in-app" | the one-way Brevo marketing-contact push — `internal/customeriosync/`; switch-gated by `CUSTOMERIO_SYNC_ENABLED` |
>
> The last three lost their **containers** a day later, at platform `838d907` (merged `0c91421`,
> 2026-08-05): compose now declares **five** services and `repos.yml` **four** entries.
>
> Consequences that hold platform-wide:
> * **The federation composes ONE subgraph** (`backend`). cms-in-app was the **3 → 1** step: the single
>   commit `graphql-wundergraph@915da06` (2026-07-29) deleted **both** `schemas/cms.graphqls` **and**
>   `schemas/jobsimulation.graphqls`, taking the supergraph from (backend, jobsimulation, cms) to
>   (backend) alone. The jobsimulation subgraph therefore **survived jobsim-in-app** and was removed
>   here, not at its own merge.
> * **All of their tabl
```

**CITED CONTENT**

```
    84  		}
    85  		authnMan, err := authn.NewManager(authnClerk)
    86  		if err != nil {
    87  			logger.Error("Authn with Clerk crashed", "error", err)
    88  			return err
    89  		}
    90  
```

## 01-042
- **id**: `B01-042`
- **corpus site**: `corpus/services/backend.md:3-99` (paragraph)
- **citation**: `docker-compose.yml:48`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> ## `app` is the backend monolith
>
> **Eight former microservices now run inside `app`**, in merge order:
>
> | Merged service | Program | What moved in |
> |---|---|---|
> | [skiller](./skiller.md) | skiller-in-app (v2.1 "quick change", July 2026) | the skills-taxonomy graph (**≥42,790 skills / ≥22,470 job roles** — public subset; [not "60K/18K"](../architecture/shared_libraries.md#taxonomy-figures)), embeddings, AI matching |
> | [skillpath](./skillpath.md) | skillpath-in-app (M502→M507) | skill-path progression engine, session state |
> | [roadrunner](./roadrunner.md) | with jobsim-in-app | Judge0 code execution (called directly via `JUDGE0_BASE_URL`) |
> | [jobsimulation](./jobsimulation.md) | jobsim-in-app (prod ECS teardown **M810 — LANDED**, `6092c6d2`) | the simulation session engine — `internal/jobsimulation/`, wired by `internal/jobsimwiring/wiring.go` |
> | [cms](./cms.md) | cms-in-app v8.0, app **v1.360.0** (prod teardown **M810 — NOT MEASURABLE here**; report both, assert neither — see the *M810 prod teardown is UNEVEN* bullet below) | content layer + Directus edge + Studio — `internal/cms/` |
> | [storage](./storage.md) | v9.0 "support-in-app", 2026-08-04 | the private + public object-storage managers — `internal/storage/`, `internal/storagens/`, `internal/publicstorage/` |
> | [messenger](./messenger.md) | v9.0 "support-in-app", 2026-08-04 | transactional email (Brevo + Liquid) and messenger's **own** Redis consumer group — `internal/messenger/`; switch-gated by `MESSENGER_ENABLED` |
> | [customerio-sync](./customerio-sync.md) | v9.0 "support-in-app" | the one-way Brevo marketing-contact push — `internal/customeriosync/`; switch-gated by `CUSTOMERIO_SYNC_ENABLED` |
>
> The last three lost their **containers** a day later, at platform `838d907` (merged `0c91421`,
> 2026-08-05): compose now declares **five** services and `repos.yml` **four** entries.
>
> Consequences that hold platform-wide:
> * **The federation composes ONE subgraph** (`backend`). cms-in-app was the **3 → 1** step: the single
>   commit `graphql-wundergraph@915da06` (2026-07-29) deleted **both** `schemas/cms.graphqls` **and**
>   `schemas/jobsimulation.graphqls`, taking the supergraph from (backend, jobsimulation, cms) to
>   (backend) alone. The jobsimulation subgraph therefore **survived jobsim-in-app** and was removed
>   here, not at its own merge.
> * **All of their tabl
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

## 01-043
- **id**: `B01-043`
- **corpus site**: `corpus/services/backend.md:3-99` (paragraph)
- **citation**: `docker-compose.yml:57`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> ## `app` is the backend monolith
>
> **Eight former microservices now run inside `app`**, in merge order:
>
> | Merged service | Program | What moved in |
> |---|---|---|
> | [skiller](./skiller.md) | skiller-in-app (v2.1 "quick change", July 2026) | the skills-taxonomy graph (**≥42,790 skills / ≥22,470 job roles** — public subset; [not "60K/18K"](../architecture/shared_libraries.md#taxonomy-figures)), embeddings, AI matching |
> | [skillpath](./skillpath.md) | skillpath-in-app (M502→M507) | skill-path progression engine, session state |
> | [roadrunner](./roadrunner.md) | with jobsim-in-app | Judge0 code execution (called directly via `JUDGE0_BASE_URL`) |
> | [jobsimulation](./jobsimulation.md) | jobsim-in-app (prod ECS teardown **M810 — LANDED**, `6092c6d2`) | the simulation session engine — `internal/jobsimulation/`, wired by `internal/jobsimwiring/wiring.go` |
> | [cms](./cms.md) | cms-in-app v8.0, app **v1.360.0** (prod teardown **M810 — NOT MEASURABLE here**; report both, assert neither — see the *M810 prod teardown is UNEVEN* bullet below) | content layer + Directus edge + Studio — `internal/cms/` |
> | [storage](./storage.md) | v9.0 "support-in-app", 2026-08-04 | the private + public object-storage managers — `internal/storage/`, `internal/storagens/`, `internal/publicstorage/` |
> | [messenger](./messenger.md) | v9.0 "support-in-app", 2026-08-04 | transactional email (Brevo + Liquid) and messenger's **own** Redis consumer group — `internal/messenger/`; switch-gated by `MESSENGER_ENABLED` |
> | [customerio-sync](./customerio-sync.md) | v9.0 "support-in-app" | the one-way Brevo marketing-contact push — `internal/customeriosync/`; switch-gated by `CUSTOMERIO_SYNC_ENABLED` |
>
> The last three lost their **containers** a day later, at platform `838d907` (merged `0c91421`,
> 2026-08-05): compose now declares **five** services and `repos.yml` **four** entries.
>
> Consequences that hold platform-wide:
> * **The federation composes ONE subgraph** (`backend`). cms-in-app was the **3 → 1** step: the single
>   commit `graphql-wundergraph@915da06` (2026-07-29) deleted **both** `schemas/cms.graphqls` **and**
>   `schemas/jobsimulation.graphqls`, taking the supergraph from (backend, jobsimulation, cms) to
>   (backend) alone. The jobsimulation subgraph therefore **survived jobsim-in-app** and was removed
>   here, not at its own merge.
> * **All of their tabl
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

## 01-044
- **id**: `B01-044`
- **corpus site**: `corpus/services/backend.md:3-99` (paragraph)
- **citation**: `docker-compose.yml:183`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> ## `app` is the backend monolith
>
> **Eight former microservices now run inside `app`**, in merge order:
>
> | Merged service | Program | What moved in |
> |---|---|---|
> | [skiller](./skiller.md) | skiller-in-app (v2.1 "quick change", July 2026) | the skills-taxonomy graph (**≥42,790 skills / ≥22,470 job roles** — public subset; [not "60K/18K"](../architecture/shared_libraries.md#taxonomy-figures)), embeddings, AI matching |
> | [skillpath](./skillpath.md) | skillpath-in-app (M502→M507) | skill-path progression engine, session state |
> | [roadrunner](./roadrunner.md) | with jobsim-in-app | Judge0 code execution (called directly via `JUDGE0_BASE_URL`) |
> | [jobsimulation](./jobsimulation.md) | jobsim-in-app (prod ECS teardown **M810 — LANDED**, `6092c6d2`) | the simulation session engine — `internal/jobsimulation/`, wired by `internal/jobsimwiring/wiring.go` |
> | [cms](./cms.md) | cms-in-app v8.0, app **v1.360.0** (prod teardown **M810 — NOT MEASURABLE here**; report both, assert neither — see the *M810 prod teardown is UNEVEN* bullet below) | content layer + Directus edge + Studio — `internal/cms/` |
> | [storage](./storage.md) | v9.0 "support-in-app", 2026-08-04 | the private + public object-storage managers — `internal/storage/`, `internal/storagens/`, `internal/publicstorage/` |
> | [messenger](./messenger.md) | v9.0 "support-in-app", 2026-08-04 | transactional email (Brevo + Liquid) and messenger's **own** Redis consumer group — `internal/messenger/`; switch-gated by `MESSENGER_ENABLED` |
> | [customerio-sync](./customerio-sync.md) | v9.0 "support-in-app" | the one-way Brevo marketing-contact push — `internal/customeriosync/`; switch-gated by `CUSTOMERIO_SYNC_ENABLED` |
>
> The last three lost their **containers** a day later, at platform `838d907` (merged `0c91421`,
> 2026-08-05): compose now declares **five** services and `repos.yml` **four** entries.
>
> Consequences that hold platform-wide:
> * **The federation composes ONE subgraph** (`backend`). cms-in-app was the **3 → 1** step: the single
>   commit `graphql-wundergraph@915da06` (2026-07-29) deleted **both** `schemas/cms.graphqls` **and**
>   `schemas/jobsimulation.graphqls`, taking the supergraph from (backend, jobsimulation, cms) to
>   (backend) alone. The jobsimulation subgraph therefore **survived jobsim-in-app** and was removed
>   here, not at its own merge.
> * **All of their tabl
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
