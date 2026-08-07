# TIER-1 ADJUDICATION BATCH 05 — 44 pairs

Each item: the CLAIMING UNIT (corpus prose) and the CITED CONTENT it points at (source lines,
line-numbered, +-3 lines of context). Answer ONE question per item.

## 05-001
- **id**: `B05-001`
- **corpus site**: `corpus/services/jobsimulation.md:232-245` (paragraph)
- **citation**: `app/main.go:670`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
**The bind survived the fold and moved onto `backend`.** `docker-compose.yml:91` binds
`$HOME/.aws/credentials:/root/.aws/credentials:ro` — the **only** AWS bind in the file — under `backend`'s
`volumes:` (`:90`), and compose's own comment says why (`:88-89`: *"jobsim-in-app's Chime/LiveKit recording
managers use the AWS SDK default credential chain — the mount the standalone jobsimulation container had."*).
Measured at platform `0dab54d`. **When the host path does not exist, Docker auto-creates it as an empty
DIRECTORY.** The container then sees a *directory* where a file belongs, and `aws-sdk-go-v2`'s
`config.LoadDefaultConfig()` **opens it successfully** (opening a directory succeeds!) before failing `EISDIR`
on the read — so it is *not* skipped as an unreadable file. In the standalone binary that error propagated out
of `ai.NewAIManager` → the root `RunE` → cobra's usage block → `exit 1`. **The CAUSE is inherited; the
SIGNATURE is not, and the container name is not the only thing that changed.** In `backend` the identical
`config.LoadDefaultConfig` failure comes out of `jsai.NewAIManager` (`app/internal/jobsimulation/ai/ai.go:90`,
`can't load AWS config: %w`), is returned unwrapped by `jobsimwiring.Wire`
(`app/internal/jobsimwiring/wiring.go:147-148`) and dies at `log.Fatalf` in `app/main.go:670` — one timestamped
line, no `Error:` prefix, no usage block (`app` `9d00a313` v1.367.0; the fatal is `:614` @ `b948604`).
```

**CITED CONTENT**

```
   667  		skillerAzureKeyEu = &v
   668  	}
   669  	if v := os.Getenv("SKILLER_AZURE_OPENAI_ENDPOINT_URL"); v != "" {
   670  		skillerAzureEndpointEu = &v
   671  	}
   672  	if v := os.Getenv("SKILLER_OPENAI_KEY"); v != "" {
   673  		skillerOpenAIKey = &v
```

## 05-002
- **id**: `B05-002`
- **corpus site**: `corpus/services/messenger.md:17-17` (paragraph)
- **citation**: `messenger/internal/flow/flow.go:72-104`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/messenger/internal/flow/flow.go`  (135 lines)

**CLAIMING UNIT**

```md
Callers don't talk to Brevo directly — they **publish Redis Stream events that the messenger flow consumes** (`messenger/internal/flow/flow.go:72-104` @ `fa47850` adds a subscriber on the `backend` stream with **21** handlers — 22 `pubsub.EventHandler(…)` lines of which one, `OrgJobSimulationAssignmentPastDueHandler`, is commented out *"not implemented"*; `app` runs that same subscriber now, on messenger's own consumer group). Messenger then decides whether to send immediately, apply org-level whitelabel branding, or skip the message entirely based on per-domain notification rules (e.g., it skips job-sim emails for stale/re-triggered sessions). (Scheduling RPCs exist in the proto but are not yet implemented — they return Unimplemented.)
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

## 05-003
- **id**: `B05-003`
- **corpus site**: `corpus/services/messenger.md:19-35` (paragraph)
- **citation**: `messenger/cmd/root.go:118-142`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/messenger/cmd/root.go`  (202 lines)

**CLAIMING UNIT**

```md
> **⚠️ Nobody "fires a Messenger RPC", and nobody ever did.** The standalone *exposed* a `MessengerService`
> Connect-RPC surface, but **no service in the platform ever constructed a client for it**: `MESSENGER_RPC_ADDR`
> occurs in **no** repo — measured at each clone's own named ref, and including the two NESTED repos a
> host-ref grep cannot see (`app/studio` and `cms/studio`, both `anthropos-studio-room` @ `aeec036`).
> And `git -C stack-demo/platform log -S 'MESSENGER_RPC' --oneline 0c91421d | wc -l` returns **0**
> commits over the platform's whole 121-commit history (**positive control**, same repo and ref:
> `-S 'SKILLER_RPC'` returns **7**; add `--all` and both become 8-and-0, the 8th being `464dfe3` on a
> non-main branch — so state the scope). The earlier control figure here was **3**, which is reproducible
> at no repo, ref, spelling or scope; corrected M257x iter-96. The RPC traffic ran the other way — messenger
> called **out** to `backend` on four addresses, all `http://backend:8083` (`messenger/cmd/root.go:118-142`)
> — an **end-state** claim, and the true one. Not to be restated as *"`d11a403` re-pointed all four"*: that
> commit moved **two** (`CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR`); the other two already held that value at
> `d11a403^` (M257x iter-115).
> **Compose set those four on the `messenger` service block and nowhere else**, so deleting the block
> deleted them: since `838d907` **no compose file sets any `*_RPC_ADDR` at all**, and there is no
> messenger process left to hold a client. Corrected M257x iter-85, re-derived at iter-87; the same
> sentence stood in [`README.md`](README.md) and was repaired in the same pass.
```

**CITED CONTENT**

```
   115  			pubsub.WithServerLogger(logger),
   116  		)
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
   144  		cache := cache.New(5*time.Minute, 10*time.Minute)
   145  
```

## 05-004
- **id**: `B05-004`
- **corpus site**: `corpus/services/messenger.md:37-59` (paragraph)
- **citation**: `terraform/main.tf:19`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/main.tf`  (787 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — the v9.0 fold landed 2026-08-04**, in the same program that folded
> `storage`, and on the same morning; the container went the next day, with storage's and
> customerio-sync's, at platform `838d907` (merged `0c91421`, 2026-08-05, *"drop the storage,
> messenger and customerio-sync containers"*). Re-derived at `app` `9d00a313` v1.367.0 /
> `messenger` `e9421c6`.
>
> **Which tree settles which row, stated because getting it backwards inverts the answer.** `e9421c6` is
> `messenger`'s **`origin/main`**, and the **prod** row is a claim about *production infrastructure*, so
> origin/main is the tree that settles it. The demo's `messenger` clone is a frozen legacy checkout **7
> commits behind** at `fa47850d`, where `terraform/main.tf:19` reads `service_desired_count = 1` — **one,
> not zero** — and there is no `:29` declaration at all; grading the prod row against that clone returns
> the opposite verdict. The **local** row, by contrast, is a claim about a stack and is settled by platform `0c91421`
> + the clone set. (Noted M257x iter-102.)
>
> | side | measured |
> |---|---|
> | **prod** | `messenger/terraform/main.tf:29` `service_desired_count = 0` **@ `messenger` `e9421c6` (= `origin/main`)** — the compute is stopped, the cms precedent again. Image and task definition stay declared: this is the rollback path, a one-line revert plus an apply (`:27-28`, the in-file comment saying exactly that) |
> | **consumer** | `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63` @ `app` **`ad9f3c49`**) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`). It does **not** merge messenger's handlers onto app's own subscribers — it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's, and it is a literal on purpose: the standalone read `cmp.Or(os.Getenv("SERVICE_NAME"), "messenger")` and nothing in terraform ever set `SERVICE_NAME` for it (`:1416-1421`). **Every anchor in this row moved between `9d00a313` and `2035f9a` without the code moving** — re-derived M257x iter-87, ref re-stated M257x iter-100, and re-stated **again** at iter-102. **This cell is the reference specimen of the stale-currency-pin class:** iter-100 correctly replaced the bare `9d00a313` — but rep
```

**CITED CONTENT**

```
    16  // This is used to determine which migration to run.
    17  data "atlas_migration" "backend_migration" {
    18    dir = "${path.module}/migrations?format=atlas"
    19    url = "${aws_ssm_parameter.supabase_db_conn.value}?search_path=public"
    20  }
    21  
    22  // Sync the state of the target database with the migrations directory.
```

## 05-005
- **id**: `B05-005`
- **corpus site**: `corpus/services/messenger.md:37-59` (paragraph)
- **citation**: `messenger/terraform/main.tf:29`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/messenger/terraform/main.tf`  (112 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — the v9.0 fold landed 2026-08-04**, in the same program that folded
> `storage`, and on the same morning; the container went the next day, with storage's and
> customerio-sync's, at platform `838d907` (merged `0c91421`, 2026-08-05, *"drop the storage,
> messenger and customerio-sync containers"*). Re-derived at `app` `9d00a313` v1.367.0 /
> `messenger` `e9421c6`.
>
> **Which tree settles which row, stated because getting it backwards inverts the answer.** `e9421c6` is
> `messenger`'s **`origin/main`**, and the **prod** row is a claim about *production infrastructure*, so
> origin/main is the tree that settles it. The demo's `messenger` clone is a frozen legacy checkout **7
> commits behind** at `fa47850d`, where `terraform/main.tf:19` reads `service_desired_count = 1` — **one,
> not zero** — and there is no `:29` declaration at all; grading the prod row against that clone returns
> the opposite verdict. The **local** row, by contrast, is a claim about a stack and is settled by platform `0c91421`
> + the clone set. (Noted M257x iter-102.)
>
> | side | measured |
> |---|---|
> | **prod** | `messenger/terraform/main.tf:29` `service_desired_count = 0` **@ `messenger` `e9421c6` (= `origin/main`)** — the compute is stopped, the cms precedent again. Image and task definition stay declared: this is the rollback path, a one-line revert plus an apply (`:27-28`, the in-file comment saying exactly that) |
> | **consumer** | `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63` @ `app` **`ad9f3c49`**) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`). It does **not** merge messenger's handlers onto app's own subscribers — it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's, and it is a literal on purpose: the standalone read `cmp.Or(os.Getenv("SERVICE_NAME"), "messenger")` and nothing in terraform ever set `SERVICE_NAME` for it (`:1416-1421`). **Every anchor in this row moved between `9d00a313` and `2035f9a` without the code moving** — re-derived M257x iter-87, ref re-stated M257x iter-100, and re-stated **again** at iter-102. **This cell is the reference specimen of the stale-currency-pin class:** iter-100 correctly replaced the bare `9d00a313` — but rep
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

## 05-006
- **id**: `B05-006`
- **corpus site**: `corpus/services/messenger.md:37-59` (paragraph)
- **citation**: `app/main.go:15`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — the v9.0 fold landed 2026-08-04**, in the same program that folded
> `storage`, and on the same morning; the container went the next day, with storage's and
> customerio-sync's, at platform `838d907` (merged `0c91421`, 2026-08-05, *"drop the storage,
> messenger and customerio-sync containers"*). Re-derived at `app` `9d00a313` v1.367.0 /
> `messenger` `e9421c6`.
>
> **Which tree settles which row, stated because getting it backwards inverts the answer.** `e9421c6` is
> `messenger`'s **`origin/main`**, and the **prod** row is a claim about *production infrastructure*, so
> origin/main is the tree that settles it. The demo's `messenger` clone is a frozen legacy checkout **7
> commits behind** at `fa47850d`, where `terraform/main.tf:19` reads `service_desired_count = 1` — **one,
> not zero** — and there is no `:29` declaration at all; grading the prod row against that clone returns
> the opposite verdict. The **local** row, by contrast, is a claim about a stack and is settled by platform `0c91421`
> + the clone set. (Noted M257x iter-102.)
>
> | side | measured |
> |---|---|
> | **prod** | `messenger/terraform/main.tf:29` `service_desired_count = 0` **@ `messenger` `e9421c6` (= `origin/main`)** — the compute is stopped, the cms precedent again. Image and task definition stay declared: this is the rollback path, a one-line revert plus an apply (`:27-28`, the in-file comment saying exactly that) |
> | **consumer** | `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63` @ `app` **`ad9f3c49`**) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`). It does **not** merge messenger's handlers onto app's own subscribers — it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's, and it is a literal on purpose: the standalone read `cmp.Or(os.Getenv("SERVICE_NAME"), "messenger")` and nothing in terraform ever set `SERVICE_NAME` for it (`:1416-1421`). **Every anchor in this row moved between `9d00a313` and `2035f9a` without the code moving** — re-derived M257x iter-87, ref re-stated M257x iter-100, and re-stated **again** at iter-102. **This cell is the reference specimen of the stale-currency-pin class:** iter-100 correctly replaced the bare `9d00a313` — but rep
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

## 05-007
- **id**: `B05-007`
- **corpus site**: `corpus/services/messenger.md:37-59` (paragraph)
- **citation**: `repos.yml:21-23`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` — the v9.0 fold landed 2026-08-04**, in the same program that folded
> `storage`, and on the same morning; the container went the next day, with storage's and
> customerio-sync's, at platform `838d907` (merged `0c91421`, 2026-08-05, *"drop the storage,
> messenger and customerio-sync containers"*). Re-derived at `app` `9d00a313` v1.367.0 /
> `messenger` `e9421c6`.
>
> **Which tree settles which row, stated because getting it backwards inverts the answer.** `e9421c6` is
> `messenger`'s **`origin/main`**, and the **prod** row is a claim about *production infrastructure*, so
> origin/main is the tree that settles it. The demo's `messenger` clone is a frozen legacy checkout **7
> commits behind** at `fa47850d`, where `terraform/main.tf:19` reads `service_desired_count = 1` — **one,
> not zero** — and there is no `:29` declaration at all; grading the prod row against that clone returns
> the opposite verdict. The **local** row, by contrast, is a claim about a stack and is settled by platform `0c91421`
> + the clone set. (Noted M257x iter-102.)
>
> | side | measured |
> |---|---|
> | **prod** | `messenger/terraform/main.tf:29` `service_desired_count = 0` **@ `messenger` `e9421c6` (= `origin/main`)** — the compute is stopped, the cms precedent again. Image and task definition stay declared: this is the rollback path, a one-line revert plus an apply (`:27-28`, the in-file comment saying exactly that) |
> | **consumer** | `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63` @ `app` **`ad9f3c49`**) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`). It does **not** merge messenger's handlers onto app's own subscribers — it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's, and it is a literal on purpose: the standalone read `cmp.Or(os.Getenv("SERVICE_NAME"), "messenger")` and nothing in terraform ever set `SERVICE_NAME` for it (`:1416-1421`). **Every anchor in this row moved between `9d00a313` and `2035f9a` without the code moving** — re-derived M257x iter-87, ref re-stated M257x iter-100, and re-stated **again** at iter-102. **This cell is the reference specimen of the stale-currency-pin class:** iter-100 correctly replaced the bare `9d00a313` — but rep
```

**CITED CONTENT**

```
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

## 05-008
- **id**: `B05-008`
- **corpus site**: `corpus/services/messenger.md:61-68` (paragraph)
- **citation**: `app/env_guards.go:61`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/env_guards.go`  (202 lines)

**CLAIMING UNIT**

```md
> **You cannot run it locally at all.** There is no messenger container and no `messenger` profile —
> both were deleted at `838d907`. `make up` never started it even before that (it was opt-in), and
> since the v9.0 fold `backend` does its work in-process, gated by `MESSENGER_ENABLED`
> (`app/env_guards.go:61` @ `app` **`ad9f3c49`** — `origin/main` and the demo's build pin on 2026-08-06;
> identical at `2035f9a4`), which defaults to **off** on a developer machine. **The ref is not optional
> on this anchor** (M257x iter-102): `env_guards.go` **did not exist** at the demo's former pin
> `b948604f` (`git -C stack-demo/app ls-tree b948604f -- env_guards.go` → empty), so this citation used to
> resolve at no ref the document named.
```

**CITED CONTENT**

```
    58  // the whole failure mode here is somebody's environment accidentally satisfying a
    59  // predicate they never read.
    60  const (
    61  	envMessengerEnabled      = "MESSENGER_ENABLED"
    62  	envCustomerIOSyncEnabled = "CUSTOMERIO_SYNC_ENABLED"
    63  )
    64  
```

## 05-009
- **id**: `B05-009`
- **corpus site**: `corpus/services/messenger.md:77-77` (bullet)
- **citation**: `cmd/root.go:63`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/cmd/root.go`  (287 lines)

**CLAIMING UNIT**

```md
* **Ports**: `8200` (HTTP) and `8201` (Connect-RPC) were published 1:1 by the compose block until `838d907` deleted it; **nothing publishes them on a stack now**. The binary's own defaults are 8080/8081 (`cmd/root.go:63`, `:64`)
```

**CITED CONTENT**

```
    60  
    61  // rootCmd represents the base command when called without any subcommands
    62  var rootCmd = &cobra.Command{
    63  	Use:   "cms",
    64  	Short: "CMS Service",
    65  	RunE: func(cmd *cobra.Command, args []string) error {
    66  		environment := colony.ReadFromEnvVar()
```

## 05-010
- **id**: `B05-010`
- **corpus site**: `corpus/services/messenger.md:103-103` (paragraph)
- **citation**: `cmd/root.go:147`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/cmd/root.go`  (287 lines)

**CLAIMING UNIT**

```md
Recent work in v0.34.0 added **whitelabel support**: when an org has custom branding (logo URL, custom invitation templates), Messenger renders subject and body separately so the Brevo send can include the org's logo and styling. The org lookup uses a **read-only Postgres connection** (`READONLY_DB_CONNECTION`, formerly `COPILOT_DB_CONNECTION` — see `cmd/root.go:147`) so the rendering path doesn't contend with the write-heavy backend load.
```

**CITED CONTENT**

```
   144  		}
   145  
   146  		// Asynq client
   147  		redisWorkerIndex, err := strconv.Atoi(os.Getenv("REDIS_WORKER_INDEX"))
   148  		if err != nil {
   149  			logger.Error("can't convert REDIS_WORKER_INDEX to int", "error", err)
   150  			return fmt.Errorf("can't convert REDIS_WORKER_INDEX to int %w", err)
```

## 05-011
- **id**: `B05-011`
- **corpus site**: `corpus/services/messenger.md:112-112` (table-row)
- **citation**: `internal/rpcsrv/rpcsrv.go:25-30`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/jobsimulation/internal/rpcsrv/rpcsrv.go`  (243 lines)

**CLAIMING UNIT**

```md
| `Schedule(message, schedule_for)` | Schedule a future email | Stub — returns `Unimplemented` (`internal/rpcsrv/rpcsrv.go:25-30`) |
```

**CITED CONTENT**

```
    22  }
    23  
    24  type RPCServer struct {
    25  	simManager        manager.Manager
    26  	inbound           inbound.Inbound
    27  	internalConverter internalConverter.Converter
    28  	anticheatManager  *anticheat.AnticheatManager
    29  }
    30  
    31  func New(conf Config) *RPCServer {
    32  	return &RPCServer{
    33  		simManager:        conf.SimManager,
```

## 05-012
- **id**: `B05-012`
- **corpus site**: `corpus/services/messenger.md:113-113` (table-row)
- **citation**: `internal/rpcsrv/rpcsrv.go:25-30`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/jobsimulation/internal/rpcsrv/rpcsrv.go`  (243 lines)

**CLAIMING UNIT**

```md
| `CancelScheduledMessage(id)` | Cancel a previously scheduled message | Stub — returns `Unimplemented` (`internal/rpcsrv/rpcsrv.go:25-30`) |
```

**CITED CONTENT**

```
    22  }
    23  
    24  type RPCServer struct {
    25  	simManager        manager.Manager
    26  	inbound           inbound.Inbound
    27  	internalConverter internalConverter.Converter
    28  	anticheatManager  *anticheat.AnticheatManager
    29  }
    30  
    31  func New(conf Config) *RPCServer {
    32  	return &RPCServer{
    33  		simManager:        conf.SimManager,
```

## 05-013
- **id**: `B05-013`
- **corpus site**: `corpus/services/messenger.md:119-119` (paragraph)
- **citation**: `app/main.go:1095`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
Most messenger sends are reactive — driven by **Redis Streams** events on the `jobsimulation`, `cms` and `backend` streams. The stream *names* outlived the services: since the merges they are published from inside `app` (e.g. the `CMS_STREAM` publisher at `app/main.go:1095`, and the whole subscriber stream binding at `:1478-1484` @ `app` `9d00a313` v1.367.0), so there is no separate producer service in compose behind any of them. The corresponding flow handlers in `internal/flow/` decide whether a stream event should produce an email, what template to use, and whether to apply staleness guards (e.g., for job-sim completions it drops the email if the session ended >2h ago, or has no end time and started >12h ago — `internal/flow/jobsimulations.go:140-151`). See `internal/flow/jobsimulations.go` for examples.
```

**CITED CONTENT**

```
  1092  		log.Fatalf("DIRECTUS_BASE_ADDR is required to wire the cms-in-app managers; the external Directus edge is unconfigured")
  1093  	}
  1094  
  1095  	cmsWorkerIndex, _ := strconv.Atoi(os.Getenv("REDIS_WORKER_INDEX"))
  1096  	cmsAsynq := cmsasync.NewClient(redisAddr, cmsWorkerIndex)
  1097  	// Storage namespace is the S3 key PREFIX, bound at client construction — it is part of the
  1098  	// object's physical address. cms's objects were written by the standalone cms service under
```

## 05-014
- **id**: `B05-014`
- **corpus site**: `corpus/services/messenger.md:119-119` (paragraph)
- **citation**: `internal/flow/jobsimulations.go:140-151`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/messenger/internal/flow/jobsimulations.go`  (515 lines)

**CLAIMING UNIT**

```md
Most messenger sends are reactive — driven by **Redis Streams** events on the `jobsimulation`, `cms` and `backend` streams. The stream *names* outlived the services: since the merges they are published from inside `app` (e.g. the `CMS_STREAM` publisher at `app/main.go:1095`, and the whole subscriber stream binding at `:1478-1484` @ `app` `9d00a313` v1.367.0), so there is no separate producer service in compose behind any of them. The corresponding flow handlers in `internal/flow/` decide whether a stream event should produce an email, what template to use, and whether to apply staleness guards (e.g., for job-sim completions it drops the email if the session ended >2h ago, or has no end time and started >12h ago — `internal/flow/jobsimulations.go:140-151`). See `internal/flow/jobsimulations.go` for examples.
```

**CITED CONTENT**

```
   137  		return nil
   138  	}
   139  
   140  	// skip email if the session ended more than 2 hours ago (catch-all for retriggered old sessions
   141  	// where the event ended_at gets set to now but started_at is never overwritten)
   142  	if session.GetEndedAt() != nil && time.Since(session.GetEndedAt().AsTime()) > 2*time.Hour {
   143  		h.logger.With("sim", e.SimulationId, "session", e.SessionId).Info("session ended more than 2 hours ago, skipping email")
   144  		return nil
   145  	}
   146  
   147  	// skip email if the session has no end time and started more than 12 hours ago
   148  	if session.GetEndedAt() == nil && session.GetStartedAt() != nil && time.Since(session.GetStartedAt().AsTime()) > 12*time.Hour {
   149  		h.logger.With("sim", e.SimulationId, "session", e.SessionId).Info("session started more than 12 hours ago, skipping email")
   150  		return nil
   151  	}
   152  
   153  	totalTimeSpent := humanize.RelTime(session.GetEndedAt().AsTime(), session.GetStartedAt().AsTime(), "", "")
   154  	totalTimeSpent = strings.Trim(totalTimeSpent, " ")
```

## 05-015
- **id**: `B05-015`
- **corpus site**: `corpus/services/messenger.md:123-123` (bullet)
- **citation**: `cmd/root.go:118-142`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/cmd/root.go`  (287 lines)

**CLAIMING UNIT**

```md
* **RPC clients**: the binary still constructs four Connect-RPC clients — CMS, backend users + organizations, skiller, and jobsimulation (`cmd/root.go:118-142`) — each reading its address from the environment. At `0dab54d` compose supplied all four, and all four resolved to the one `backend` mux (`http://backend:8083`) — **two of them because `d11a403` re-pointed them, two because they had always read `backend`**; `838d907` deleted the service block, and with it **every `*_RPC_ADDR` compose ever set** — there are now zero across `docker-compose.yml`, `common.yml` and `.env_example`. So on a stack today those clients are neither constructed nor addressed: the process does not run. The `cms` and `jobsimulation` services they were named for went earlier, at `d11a403`; their surfaces are registered on `app`'s RPC server. Skill-path notifications arrive as Redis Streams events on the `backend` subscriber (`OrgSkillPath*` handlers in `internal/flow/flow.go:74-78`), not via a direct Skillpath RPC.
```

**CITED CONTENT**

```
   115  		skillerClient := skillerv1connect.NewSkillerServiceClient(
   116  			rpc.NewHttpClient(),
   117  			os.Getenv("SKILLER_RPC_ADDR"),
   118  			rpc.DefaultInterceptors,
   119  		)
   120  
   121  		jobSimulationClient := jobsimulationv1connect.NewJobSimulationServiceClient(
   122  			rpc.NewHttpClient(),
   123  			os.Getenv("JOBSIMULATION_RPC_ADDR"),
   124  			rpc.DefaultInterceptors,
   125  		)
   126  
   127  		dirClient := directus.New(
   128  			os.Getenv("DIRECTUS_BASE_ADDR"),
   129  			os.Getenv("DIRECTUS_PUBLIC_BASE_ADDR"),
   130  			os.Getenv("DIRECTUS_TOKEN"),
   131  			redisClient,
   132  			skillerClient,
   133  		)
   134  
   135  		orgClient := organizationsv1connect.NewOrganizationsServiceClient(
   136  			rpc.NewHttpClient(),
   137  			os.Getenv("BACKEND_USERS_RPC_ADDR"),
   138  			rpc.DefaultInterceptors,
   139  		)
   140  
   141  		storageClient := storage.NewClient(os.Getenv("STORAGE_RPC_ADDR"), serviceName)
   142  		if err != nil {
   143  			return fmt.Errorf("can't init storage client %w", err)
   144  		}
   145  
```

## 05-016
- **id**: `B05-016`
- **corpus site**: `corpus/services/messenger.md:123-123` (bullet)
- **citation**: `internal/flow/flow.go:74-78`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/messenger/internal/flow/flow.go`  (135 lines)

**CLAIMING UNIT**

```md
* **RPC clients**: the binary still constructs four Connect-RPC clients — CMS, backend users + organizations, skiller, and jobsimulation (`cmd/root.go:118-142`) — each reading its address from the environment. At `0dab54d` compose supplied all four, and all four resolved to the one `backend` mux (`http://backend:8083`) — **two of them because `d11a403` re-pointed them, two because they had always read `backend`**; `838d907` deleted the service block, and with it **every `*_RPC_ADDR` compose ever set** — there are now zero across `docker-compose.yml`, `common.yml` and `.env_example`. So on a stack today those clients are neither constructed nor addressed: the process does not run. The `cms` and `jobsimulation` services they were named for went earlier, at `d11a403`; their surfaces are registered on `app`'s RPC server. Skill-path notifications arrive as Redis Streams events on the `backend` subscriber (`OrgSkillPath*` handlers in `internal/flow/flow.go:74-78`), not via a direct Skillpath RPC.
```

**CITED CONTENT**

```
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
```

## 05-017
- **id**: `B05-017`
- **corpus site**: `corpus/services/messenger.md:146-149` (paragraph)
- **citation**: `docker-compose.yml:84-92`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
**To exercise the mail path today you enable it inside `backend`**: set `MESSENGER_ENABLED=true` in
`platform/.env` (compose deliberately sets no value for it — pinning one there would override `.env`;
see the comment on the `backend` block, `docker-compose.yml:84-92`). Know what that does: `app`
attaches to messenger's **live** Redis consumer group, and a non-empty `BREVO_KEY` sends real mail.
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

## 05-018
- **id**: `B05-018`
- **corpus site**: `corpus/services/messenger.md:162-167` (paragraph)
- **citation**: `app/main.go:295-300`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
Set `BREVO_KEY=""` to route through the **console sender** (`internal/messenger/console/`) instead of
hitting Brevo — emails print to stdout. (That fallback is standalone-only: `app` did **not** port it —
`app/main.go:295-300` @ `app` **`ad9f3c49`** (identical at `2035f9a4`): the condition is at `:295`
(`MESSENGER_ENABLED` **or** `CUSTOMERIO_SYNC_ENABLED` on with an empty `BREVO_KEY`) and the `log.Fatalf`
at `:296`. Ref re-stated M257x iter-102 — the citation was previously unpinned and present-tense, and
`:295` at the demo's former pin `b948604f` is a different construct.)
```

**CITED CONTENT**

```
   292  	// one: sender.NewFromEnv always returns the Brevo sender (the console fallback is
   293  	// deliberately not ported), so an empty key means every send and every contact
   294  	// write 401s while app boots healthy and the deploy goes green.
   295  	if (messengerEnabled || customerIOSyncEnabled) && os.Getenv("BREVO_KEY") == "" {
   296  		log.Fatalf("messenger-in-app: BREVO_KEY is required when %s or %s is on; "+
   297  			"sender.NewFromEnv always returns the Brevo sender, so an empty key means every send "+
   298  			"fails 401 at Brevo — app would boot healthy and discard all outbound email",
   299  			envMessengerEnabled, envCustomerIOSyncEnabled)
   300  	}
   301  
   302  	dbConn, err := colony.NewDBStdConn(serverContext, cancelServerContext, logger, os.Getenv("SUPABASE_DB_CONN"),
   303  		colony.WithDBCustomSearchPath("public", "extensions"),
```

## 05-019
- **id**: `B05-019`
- **corpus site**: `corpus/services/messenger.md:184-184` (table-row)
- **citation**: `cmd/root.go:107`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/cmd/root.go`  (287 lines)

**CLAIMING UNIT**

```md
| `REDIS_WORKER_INDEX` | `0` | Was set in docker-compose (=0) but NOT read by the code — there is no worker pool / separate worker Redis index; only `REDIS_STREAMS_INDEX` is consumed (`cmd/root.go:107`). |
```

**CITED CONTENT**

```
   104  		}
   105  		// Ent db connection
   106  		entClient := ent.NewClient(ent.Driver(
   107  			entsql.OpenDB(dialect.Postgres, db),
   108  		))
   109  
   110  		azureAIClient, err := openai.NewAzure(os.Getenv("CMS_AZURE_OPENAI_KEY"), os.Getenv("CMS_AZURE_OPENAI_ENDPOINT_URL"), nil)
```

## 05-020
- **id**: `B05-020`
- **corpus site**: `corpus/services/messenger.md:186-186` (table-row)
- **citation**: `app/main.go:1205-1211`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| `CMS_RPC_ADDR` | *(unset — was `http://backend:8083`)* | CMS RPC. M809 re-pointed it off the standalone `cms` onto the `backend` mux at `d11a403`; `838d907` then removed the variable altogether. The earlier `http://cms:8091` was true at `2adcf71` only. `app`'s own comment at `app/main.go:1205-1211` (@ `b948604` v1.366.0) still says *"additive + DORMANT … until the M809 re-point"* and is **stale in `app`** |
```

**CITED CONTENT**

```
  1202  	// standalone cms. Active whenever the Directus edge is configured (the release sets it);
  1203  	// the external client the switch was seeded with is only the construction-time placeholder.
  1204  	cmsReaderSw.set(cmsRPCServer)
  1205  	// M805: consume the cms studio + ai_video Asynq queue in-process (the app is the sole
  1206  	// consumer post-release — the standalone cms takes no traffic). The consumer polls the SAME
  1207  	// DB index the enqueue client writes to (audit R2). The studio gen.py/postgen.py pipeline
  1208  	// is argv-safe (M809b H-1 fixed).
  1209  	cmsWorker := cmsworker.NewServer(redisAddr, cmsWorkerIndex, logger)
  1210  	wg.Go(func() {
  1211  		defer cancelServerContext()
  1212  		if err := cmsWorker.Start(serverContext, cmsManagers.Studio, cmsManagers.AiVideo); err != nil {
  1213  			logger.Info("shutting down the cms worker", "error", err)
  1214  		}
```

## 05-021
- **id**: `B05-021`
- **corpus site**: `corpus/services/messenger.md:189-189` (table-row)
- **citation**: `internal/flow/assignments.go:828`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/messenger/internal/flow/assignments.go`  (847 lines)

**CLAIMING UNIT**

```md
| ~~`SKILLPATH_RPC_ADDR`~~ | *(removed earlier)* | **Gone from docker-compose** since skillpath was decommissioned into `app` ("skillpath-in-app", M502→M507) — only the residual `SKILLPATH_STREAM=skillpath` remains, on `backend`. Messenger never had a Skillpath RPC client anyway; skill-path data is read via the CMS client (`internal/flow/assignments.go:828`, in `getSkillPath`). |
```

**CITED CONTENT**

```
   825  }
   826  
   827  func (h *Manager) getSkillPath(ctx context.Context, skillPathId string) (*skillpath.SkillPath, error) {
   828  	res, err := h.cms.GetSkillPath(ctx, connect.NewRequest(&cmsv1.GetSkillPathRequest{
   829  		Id: skillPathId,
   830  	}))
   831  	if err != nil {
```

## 05-022
- **id**: `B05-022`
- **corpus site**: `corpus/services/messenger.md:191-191` (paragraph)
- **citation**: `cmd/root.go:63`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/cmd/root.go`  (287 lines)

**CLAIMING UNIT**

```md
> The binary's built-in fallbacks when the env var is unset are `PORT=8080` (`cmd/root.go:63`), `RPC_PORT=8081` (`cmd/root.go:64`), `REDIS_STREAMS_INDEX=2` (`cmd/root.go:107`).
```

**CITED CONTENT**

```
    60  
    61  // rootCmd represents the base command when called without any subcommands
    62  var rootCmd = &cobra.Command{
    63  	Use:   "cms",
    64  	Short: "CMS Service",
    65  	RunE: func(cmd *cobra.Command, args []string) error {
    66  		environment := colony.ReadFromEnvVar()
```

## 05-023
- **id**: `B05-023`
- **corpus site**: `corpus/services/messenger.md:191-191` (paragraph)
- **citation**: `cmd/root.go:64`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/cmd/root.go`  (287 lines)

**CLAIMING UNIT**

```md
> The binary's built-in fallbacks when the env var is unset are `PORT=8080` (`cmd/root.go:63`), `RPC_PORT=8081` (`cmd/root.go:64`), `REDIS_STREAMS_INDEX=2` (`cmd/root.go:107`).
```

**CITED CONTENT**

```
    61  // rootCmd represents the base command when called without any subcommands
    62  var rootCmd = &cobra.Command{
    63  	Use:   "cms",
    64  	Short: "CMS Service",
    65  	RunE: func(cmd *cobra.Command, args []string) error {
    66  		environment := colony.ReadFromEnvVar()
    67  		logger := colony.InitLogger(versionConfig.Name,
```

## 05-024
- **id**: `B05-024`
- **corpus site**: `corpus/services/messenger.md:191-191` (paragraph)
- **citation**: `cmd/root.go:107`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/cmd/root.go`  (287 lines)

**CLAIMING UNIT**

```md
> The binary's built-in fallbacks when the env var is unset are `PORT=8080` (`cmd/root.go:63`), `RPC_PORT=8081` (`cmd/root.go:64`), `REDIS_STREAMS_INDEX=2` (`cmd/root.go:107`).
```

**CITED CONTENT**

```
   104  		}
   105  		// Ent db connection
   106  		entClient := ent.NewClient(ent.Driver(
   107  			entsql.OpenDB(dialect.Postgres, db),
   108  		))
   109  
   110  		azureAIClient, err := openai.NewAzure(os.Getenv("CMS_AZURE_OPENAI_KEY"), os.Getenv("CMS_AZURE_OPENAI_ENDPOINT_URL"), nil)
```

## 05-025
- **id**: `B05-025`
- **corpus site**: `corpus/services/next-web-app.md:37-61` (paragraph)
- **citation**: `demo-stack/up-injected.sh:1076-1085`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/demo-stack/up-injected.sh`  (2695 lines)

**CLAIMING UNIT**

```md
> **⚠️ RETRACTED — "the recruiter scoreboard is an `is_hiring` org-type surface in the dockerized `apps/web`, not
> the Hiring app."** That sentence stood in the Hiring row (and was implied by the *Key Functions* bullet above)
> and is **false in both directions**. Two independent adjudicator readings booked the same anchor; corrected
> M257x iter-102.
>
> * **Why it is false.** `/enterprise/activity-dashboard` exists in *both* apps, so the route's presence in
>   `apps/web` proves nothing. What decides it is a global product-boundary guard: measured @ `next-web-app`
>   **`8297c684`**, `apps/web/src/context/UserStatusContext.tsx:144-148` computes `userHasAllHiringOrgs` from
>   `membership.organization.publicMetadata.isHiring`, and when it holds, `:168-172` sets
>   `window.location.href = buildSwitchHandoffUrl({ targetProduct: 'hiring', … next: '/home' })` — the recruiter
>   is **ejected out of `apps/web`**, on a direct navigation too. So *"the org genuinely reads as hiring"* and
>   *"the scoreboard is reachable in `apps/web`"* are **mutually exclusive**. The screen that actually renders the
>   comparison is `apps/hiring/src/components/containers/InsightsByMembersContainer.tsx:108`, mounted at
>   `apps/hiring/…/enterprise/activity-dashboard/@tabs/ai-simulations/[simId]/page.tsx:14`.
> * **Which half was true.** The scoreboard *is* driven by the `is_hiring` **org-type** and *does* render from
>   seedable data with no platform edit — that half stands. Only the **app** was wrong.
> * **Consequence for the "Dockerized?" column.** `apps/hiring` is still absent from **platform** compose
>   (`platform` `0c91421` `docker-compose.yml` declares `sentinel`, `backend`, `studio-desk`, `next-web-app`,
>   `gotenberg` — the frontend service is `apps/web` only, at `:143`), and the repo ships one
>   `Dockerfile.dev`. But a **demo** builds `apps/hiring` as a **second UI container** from the same unmodified
>   clone using rext's own `demo-stack/frontend/hiring.Dockerfile` (`demo-stack/up-injected.sh:1076-1085`, image
>   `demo-<N>-hiring`, port `3001`+offset) — still zero platform-repo edits. `❌ Vercel-only` was therefore too
>   strong as well.
>
> Authoritative statement, with the render proof: [`hiring.md`](hiring.md) § *The render path* (M224).
```

**CITED CONTENT**

```
  1073  }
  1074  
  1075  # ── M224 TOK-02: per-demo HIRING app image build (the SECOND UI container — a two-app demo) ───────────────
  1076  # Build the REAL apps/hiring from the SAME UNMODIFIED next-web-app clone the web app builds from (build
  1077  # CONTEXT only — zero platform-repo edits), baked with this stack's OFFSET URLs + the minted Clerk pk. Uses a
  1078  # rext-owned frontend/hiring.Dockerfile (the platform Dockerfile.dev hardcodes the WEB filter/port, so it
  1079  # cannot be reused verbatim — this one filters @anthropos/hiring-app and serves its port 3001). Tag-guarded +
  1080  # minted-pk-in-bundle validated exactly like next-web (fail-safe: anything unverifiable rebuilds). NO
  1081  # demopatches in tik A (the urls.ts HIRING_APP_URL chain-patch is an iter-09 concern). PK_DEMO must be set.
  1082  build_frontend_hiring() {
  1083    local ctx="$DEMO/next-web-app" img="demo-$N-hiring"   # M26: the demo's OWN monorepo clone; hiring is apps/hiring within it
  1084    local dockerfile="$HERE/frontend/hiring.Dockerfile"   # rext-owned (the platform Dockerfile.dev hardcodes the web filter/port)
  1085    [ -f "$ctx/apps/hiring/package.json" ] || { log "⚠ hiring: $ctx/apps/hiring not found — skipping UI build (non-fatal)"; return 0; }
  1086    [ -f "$dockerfile" ] || { log "⚠ hiring: $dockerfile not found — skipping UI build (non-fatal)"; return 0; }
  1087    # [M224 tik C/D] the demo-patch set THIS hiring image bakes — declared ABOVE the cache check because the reuse
  1088    # check now keys on their fingerprint. The §5-bis hazard ("applying a patch is not shipping it") applies to
```

## 05-026
- **id**: `B05-026`
- **corpus site**: `corpus/services/next-web-app.md:74-74` (bullet)
- **citation**: `docker-compose.yml:151`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
* **GraphQL**: single endpoint `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` — compose bakes `http://${PUBLIC_HOST:-localhost}:8082/graphql/query`, as a build arg (`docker-compose.yml:151`) and again in the runtime environment (`:160`), re-anchored at platform `0c91421` (it was `:236` at `0dab54d`); the env-var NAME still says wundergraph, the router behind it is gone locally; Clerk bearer token injected via React Query `defaultOptions.queries.meta.getToken`.
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

## 05-027
- **id**: `B05-027`
- **corpus site**: `corpus/services/next-web-app.md:75-75` (bullet)
- **citation**: `CLAUDE.md:55`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/CLAUDE.md`  (581 lines)

**CLAIMING UNIT**

```md
* **Auth edge**: **`apps/web/src/proxy.ts`** (and `apps/hiring/src/proxy.ts`, and `apps/integration/src/proxy.ts`) — **not `middleware.ts`**, which exists nowhere in the repo at `next-web-app` **`8297c684`** (re-derived 2026-08-06; the label here read the moving *"origin HEAD"* until M257x iter-102 — a pin is checkable, a moving label rots): **Next 16 renamed the `middleware.ts` convention to `proxy.ts`** (the repo's own `CLAUDE.md:55` says so, verbatim at that ref). `clerkMiddleware` protects every non-public route; public allowlist includes `/login`, `/sign-up`, `/checkout`, `/free-trial`, `/monitoring`, `/print`, `/api/bunny/thumbnail`. `/print` routes are HMAC-gated (`PRINT_ROUTE_SECRET`) for Puppeteer PDF generation.
```

**CITED CONTENT**

```
    52  `corpus/ops/setup_guide.md` (first-time build) + `corpus/ops/run_guide.md` (start + health) with:
    53  - Verification before/after each step + user confirmation before destructive operations
    54  - Progress tracking via TodoWrite
    55  - For an additional `dev-N`: the M13 set-dress pass (cache-first snapshot replay + a light `dev-min` seed
    56    + the per-stack-Directus firewall check), default-on + non-fatal. The per-stack Directus itself is
    57    **opt-in for dev** via `--local-content` (v1.5 M22/M23): with it the recipe is EXECUTED (a per-stack
    58    Directus boots on an offset port + `backend`'s cms domain is cut over → content self-contained); without it the stack
```

## 05-028
- **id**: `B05-028`
- **corpus site**: `corpus/services/next-web-app.md:113-117` (paragraph)
- **citation**: `docker-compose.yml:165-167`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> ⚠️ **`make up PROFILE=frontend` on its own EXITS 1 — it builds nothing.** `next-web-app` declares
> `depends_on: backend` (`docker-compose.yml:165-167`) and `backend` is `profiles: [core, backend, all]`
> (`:110`), which the `frontend` profile does not select — so compose rejects the whole project with
> *"service `next-web-app` depends on undefined service `backend`: invalid compose project."* Use
> `make up-frontend` (which adds `core`), or `make up PROFILE=all`.
```

**CITED CONTENT**

```
   162        - NEXT_PUBLIC_HOSTING_URL=http://${PUBLIC_HOST:-localhost}:3000
   163      networks:
   164        - app-network
   165      depends_on:
   166        backend:
   167          condition: service_started
   168      profiles: [frontend, all]
   169  
   170    gotenberg:
```

## 05-029
- **id**: `B05-029`
- **corpus site**: `corpus/services/next-web-app.md:160-160` (table-row)
- **citation**: `apps/web/src/app/api/dev/login-as/route.ts:79`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/apps/web/src/app/api/dev/login-as/route.ts`  (100 lines)

**CLAIMING UNIT**

```md
| `apps/web/src/app/api/dev/login-as/route.ts:79` | dev *"log in as a real Clerk user"* route → `/dev/accept` | `DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production'` (`apps/web/src/lib/devLogin.ts:28`); hard-404 otherwise. **The same boolean adds `/api/dev/login-as` + `/dev/accept` to the PUBLIC route list** (`apps/web/src/proxy.ts:56`) — it must be reachable before a session exists |
```

**CITED CONTENT**

```
    76      }
    77  
    78      // 5. Mint a one-time sign-in token (10-minute validity is plenty).
    79      const { token } = await client.signInTokens.createSignInToken({
    80        userId: user.id,
    81        expiresInSeconds: 600,
    82      });
```

## 05-030
- **id**: `B05-030`
- **corpus site**: `corpus/services/next-web-app.md:160-160` (table-row)
- **citation**: `apps/web/src/lib/devLogin.ts:28`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/apps/web/src/lib/devLogin.ts`  (35 lines)

**CLAIMING UNIT**

```md
| `apps/web/src/app/api/dev/login-as/route.ts:79` | dev *"log in as a real Clerk user"* route → `/dev/accept` | `DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production'` (`apps/web/src/lib/devLogin.ts:28`); hard-404 otherwise. **The same boolean adds `/api/dev/login-as` + `/dev/accept` to the PUBLIC route list** (`apps/web/src/proxy.ts:56`) — it must be reachable before a session exists |
```

**CITED CONTENT**

```
    25  // genuinely local-dev-only — the endpoint hard-404s anywhere it is deployed.
    26  // It also does nothing without CLERK_SECRET_KEY configured locally.
    27  
    28  export const DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production';
    29  
    30  // Optional convenience: when set, `GET /api/dev/login-as` with NO `email=`
    31  // query param signs you in as this address. Lets the agentic workflow use a
```

## 05-031
- **id**: `B05-031`
- **corpus site**: `corpus/services/next-web-app.md:160-160` (table-row)
- **citation**: `apps/web/src/proxy.ts:56`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/apps/web/src/proxy.ts`  (83 lines)

**CLAIMING UNIT**

```md
| `apps/web/src/app/api/dev/login-as/route.ts:79` | dev *"log in as a real Clerk user"* route → `/dev/accept` | `DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production'` (`apps/web/src/lib/devLogin.ts:28`); hard-404 otherwise. **The same boolean adds `/api/dev/login-as` + `/dev/accept` to the PUBLIC route list** (`apps/web/src/proxy.ts:56`) — it must be reachable before a session exists |
```

**CITED CONTENT**

```
    53    // both must be reachable BEFORE a session exists, so they're public. Gated to
    54    // local dev; production (incl. Vercel Preview) drops them and the route handler
    55    // itself hard-404s. See docs/testing-with-clerk.md.
    56    ...(DEV_LOGIN_ENABLED ? ['/api/dev/login-as(.*)', '/dev/accept(.*)'] : []),
    57  ]);
    58  
    59  const isPrintRoute = createRouteMatcher(['/print(.*)']);
```

## 05-032
- **id**: `B05-032`
- **corpus site**: `corpus/services/next-web-app.md:161-161` (table-row)
- **citation**: `e2e/auth.setup.ts:72`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/e2e/auth.setup.ts`  (87 lines)

**CLAIMING UNIT**

```md
| `e2e/auth.setup.ts:72` | the Playwright auth fixture — it mints a ticket instead of driving the password form | **no `NODE_ENV` gate.** It is a test-runner file, never in an app build, but it runs against a **real Clerk instance** |
```

**CITED CONTENT**

```
    69    if (!user) {
    70      throw new Error(`No Clerk user found with email ${email}`);
    71    }
    72    const { token: ticket } = await clerkClient.signInTokens.createSignInToken({
    73      userId: user.id,
    74      expiresInSeconds: 300,
    75    });
```

## 05-033
- **id**: `B05-033`
- **corpus site**: `corpus/services/next-web-app.md:163-168` (paragraph)
- **citation**: `e2e/auth.setup.ts:57-62`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/e2e/auth.setup.ts`  (87 lines)

**CLAIMING UNIT**

```md
**Why `auth.setup.ts` needs no `E2E_TEST_PASSWORD`, stated because the old testing note implied one:** the
e2e account *"enforces 2FA (email_code as second factor); password signin returns `needs_second_factor` and
never produces a session"*, and the ticket path means *"Clerk treats it as fully authenticated and **skips
both factors**"* (`e2e/auth.setup.ts:57-62`, the file's own words). So the suite needs `E2E_TEST_EMAIL` +
`CLERK_SECRET_KEY`, and the token is a **deliberate second-factor bypass** — the platform-side reading of
that is filed in `knowledge/plan/platform-defect-register.md`, not asserted here.
```

**CITED CONTENT**

```
    54  
    55    // Mint a one-time sign-in ticket on the backend instead of going through the
    56    // password form. Two reasons:
    57    //   1. Stefano's account enforces 2FA (email_code as second factor); password
    58    //      signin returns `needs_second_factor` and never produces a session.
    59    //   2. Tickets don't burn dev-tier signin rate limits the way repeated
    60    //      password attempts do.
    61    // The ticket is consumed in-page via `clerk.signIn({ strategy: 'ticket' })`
    62    // — Clerk treats it as fully authenticated and skips both factors.
    63    const clerkClient = createClerkClient({ secretKey });
    64    const { data: users } = await clerkClient.users.getUserList({
    65      emailAddress: [email],
```

## 05-034
- **id**: `B05-034`
- **corpus site**: `corpus/services/next-web-app.md:172-172` (bullet)
- **citation**: `CLAUDE.md:15`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/CLAUDE.md`  (581 lines)

**CLAIMING UNIT**

```md
* **Next.js 16 / React 19** — the repo went 15 → 16 and the corpus missed it for four releases. Its own `CLAUDE.md:15` says *"Next.js 16 App Router"* and is **current** (an older note here claimed it still said 14; it does not). `knowledge/next15-adoption-plan.md` survives as a superseded plan beside `UPGRADE-IMPACT-next16.md`.
```

**CITED CONTENT**

```
    12  This is NOT the Anthropos platform source code - it's the documentation about it. The actual platform code lives in separate repositories under the `anthropos-work` GitHub organization.
    13  
    14  ## Development Commands
    15  
    16  ### Available Skills
    17  
    18  | Skill | Purpose | Guide |
```

## 05-035
- **id**: `B05-035`
- **corpus site**: `corpus/services/roadrunner.md:5-74` (paragraph)
- **citation**: `roadrunner/terraform/main.tf:19`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/roadrunner/terraform/main.tf`  (96 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` / ORPHANED — nothing calls this service any more (verified v2.5 M231 KB-6; re-verified v2.7 "july jitter"
> M247 against the CONSOLIDATED platform — the ~386-commit `app` bump).** Code execution moved **in-process into
> jobsimulation** (`jobsimulation/internal/runner/runner.go`, an in-process Judge0 client whose own header comment
> reads *"formerly the standalone 'roadrunner' service"*) — and with the **jobsim-in-app** merge that runner now
> lives inside **`app`**. `backend` reads `JUDGE0_BASE_URL` and calls Judge0 directly, and **nothing calls the
> roadrunner service any more.**
>
> **⚠️ Precision, because the declarations disagree (v2.8 M257x).** *"There is no roadrunner service in
> production"* overstates it: `roadrunner/terraform/main.tf:19` still reads `service_desired_count = 1` and has
> has **not been touched since `84a4b4f` (2025-12-15)** — the commit that first added `terraform/main.tf`,
> seven months before the fold. ⚠️ **This said "`87d8d44` (2026-06-19)" until M257x iter-115, and that is not
> the line's provenance**: `87d8d44` is the repo's HEAD and touches exactly one file,
> `.github/workflows/bump-version.yml` (3 insertions) — it never goes near terraform, so *"not touched since
> it"* is vacuous by construction while the parenthetical presented it as the date of the last touch. The
> subject of *"has not been touched"* is the **line**, not the repo. `git blame -L 19,19 87d8d443 --
> terraform/main.tf` names `84a4b4f`; a file-level `git log` is not line provenance (the file's own most recent
> touch is `e45eb61`, 2026-05-27, a line-11 module-source swap). The corpus's own fenced authority,
> [`platform-migration-status.md`](../architecture/platform-migration-status.md), had this right all along and
> this document never named `84a4b4f` anywhere. The conclusion — *before the fold* — survives; the sha did not.
> Meanwhile the platform has removed it from its own
> clone set: `roadrunner` had a `repos.yml` entry reading *"legacy — folded into app"* as late as `2adcf71`
> (`repos.yml:29-31` **at that ref**), and platform `d11a403` (2026-08-03) **deleted the entry outright** —
> **grade that commit by its diff, not its message.** Its message asserts roadrunner's *"repos.yml clone entry
> was already gone"*; `git show d11a403 -- repos.yml` shows **that very commit** removing `- name: cms`,
> `- name: jobsimulation` *
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

## 05-036
- **id**: `B05-036`
- **corpus site**: `corpus/services/roadrunner.md:5-74` (paragraph)
- **citation**: `repos.yml:29-31`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` / ORPHANED — nothing calls this service any more (verified v2.5 M231 KB-6; re-verified v2.7 "july jitter"
> M247 against the CONSOLIDATED platform — the ~386-commit `app` bump).** Code execution moved **in-process into
> jobsimulation** (`jobsimulation/internal/runner/runner.go`, an in-process Judge0 client whose own header comment
> reads *"formerly the standalone 'roadrunner' service"*) — and with the **jobsim-in-app** merge that runner now
> lives inside **`app`**. `backend` reads `JUDGE0_BASE_URL` and calls Judge0 directly, and **nothing calls the
> roadrunner service any more.**
>
> **⚠️ Precision, because the declarations disagree (v2.8 M257x).** *"There is no roadrunner service in
> production"* overstates it: `roadrunner/terraform/main.tf:19` still reads `service_desired_count = 1` and has
> has **not been touched since `84a4b4f` (2025-12-15)** — the commit that first added `terraform/main.tf`,
> seven months before the fold. ⚠️ **This said "`87d8d44` (2026-06-19)" until M257x iter-115, and that is not
> the line's provenance**: `87d8d44` is the repo's HEAD and touches exactly one file,
> `.github/workflows/bump-version.yml` (3 insertions) — it never goes near terraform, so *"not touched since
> it"* is vacuous by construction while the parenthetical presented it as the date of the last touch. The
> subject of *"has not been touched"* is the **line**, not the repo. `git blame -L 19,19 87d8d443 --
> terraform/main.tf` names `84a4b4f`; a file-level `git log` is not line provenance (the file's own most recent
> touch is `e45eb61`, 2026-05-27, a line-11 module-source swap). The corpus's own fenced authority,
> [`platform-migration-status.md`](../architecture/platform-migration-status.md), had this right all along and
> this document never named `84a4b4f` anywhere. The conclusion — *before the fold* — survives; the sha did not.
> Meanwhile the platform has removed it from its own
> clone set: `roadrunner` had a `repos.yml` entry reading *"legacy — folded into app"* as late as `2adcf71`
> (`repos.yml:29-31` **at that ref**), and platform `d11a403` (2026-08-03) **deleted the entry outright** —
> **grade that commit by its diff, not its message.** Its message asserts roadrunner's *"repos.yml clone entry
> was already gone"*; `git show d11a403 -- repos.yml` shows **that very commit** removing `- name: cms`,
> `- name: jobsimulation` *
```

**CITED CONTENT**

```
    26    - name: studio-desk
    27      type: node-npm
    28      migrations: false
    29  
```

## 05-037
- **id**: `B05-037`
- **corpus site**: `corpus/services/roadrunner.md:5-74` (paragraph)
- **citation**: `jobsimulation/knowledge/operational.md:68`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/jobsimulation/knowledge/operational.md`  (155 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` / ORPHANED — nothing calls this service any more (verified v2.5 M231 KB-6; re-verified v2.7 "july jitter"
> M247 against the CONSOLIDATED platform — the ~386-commit `app` bump).** Code execution moved **in-process into
> jobsimulation** (`jobsimulation/internal/runner/runner.go`, an in-process Judge0 client whose own header comment
> reads *"formerly the standalone 'roadrunner' service"*) — and with the **jobsim-in-app** merge that runner now
> lives inside **`app`**. `backend` reads `JUDGE0_BASE_URL` and calls Judge0 directly, and **nothing calls the
> roadrunner service any more.**
>
> **⚠️ Precision, because the declarations disagree (v2.8 M257x).** *"There is no roadrunner service in
> production"* overstates it: `roadrunner/terraform/main.tf:19` still reads `service_desired_count = 1` and has
> has **not been touched since `84a4b4f` (2025-12-15)** — the commit that first added `terraform/main.tf`,
> seven months before the fold. ⚠️ **This said "`87d8d44` (2026-06-19)" until M257x iter-115, and that is not
> the line's provenance**: `87d8d44` is the repo's HEAD and touches exactly one file,
> `.github/workflows/bump-version.yml` (3 insertions) — it never goes near terraform, so *"not touched since
> it"* is vacuous by construction while the parenthetical presented it as the date of the last touch. The
> subject of *"has not been touched"* is the **line**, not the repo. `git blame -L 19,19 87d8d443 --
> terraform/main.tf` names `84a4b4f`; a file-level `git log` is not line provenance (the file's own most recent
> touch is `e45eb61`, 2026-05-27, a line-11 module-source swap). The corpus's own fenced authority,
> [`platform-migration-status.md`](../architecture/platform-migration-status.md), had this right all along and
> this document never named `84a4b4f` anywhere. The conclusion — *before the fold* — survives; the sha did not.
> Meanwhile the platform has removed it from its own
> clone set: `roadrunner` had a `repos.yml` entry reading *"legacy — folded into app"* as late as `2adcf71`
> (`repos.yml:29-31` **at that ref**), and platform `d11a403` (2026-08-03) **deleted the entry outright** —
> **grade that commit by its diff, not its message.** Its message asserts roadrunner's *"repos.yml clone entry
> was already gone"*; `git show d11a403 -- repos.yml` shows **that very commit** removing `- name: cms`,
> `- name: jobsimulation` *
```

**CITED CONTENT**

```
    65  | `BACKEND_USERS_RPC_ADDR` | App service address |
    66  | `STORAGE_RPC_ADDR` | Storage service address |
    67  | `REALTIME_RPC_ADDR` | Realtime service address |
    68  | `ROADRUNNER_RPC_ADDR` | Roadrunner service address |
    69  | `AUTHORIZATION_ADDRESS` | Sentinel service address |
    70  | `AZURE_OPENAI_KEY_EU` + `AZURE_OPENAI_ENDPOINT_URL_EU` | AI gameplay (EU) |
    71  | `AZURE_OPENAI_KEY_RESULTS` + `AZURE_OPENAI_ENDPOINT_URL_RESULTS` | AI validation |
```

## 05-038
- **id**: `B05-038`
- **corpus site**: `corpus/services/roadrunner.md:5-74` (paragraph)
- **citation**: `app/knowledge/plan/releases/07.00-jobsim-in-app/RE-PORT-CHECKLIST.md:10`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/knowledge/plan/releases/07.00-jobsim-in-app/RE-PORT-CHECKLIST.md`  (150 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` / ORPHANED — nothing calls this service any more (verified v2.5 M231 KB-6; re-verified v2.7 "july jitter"
> M247 against the CONSOLIDATED platform — the ~386-commit `app` bump).** Code execution moved **in-process into
> jobsimulation** (`jobsimulation/internal/runner/runner.go`, an in-process Judge0 client whose own header comment
> reads *"formerly the standalone 'roadrunner' service"*) — and with the **jobsim-in-app** merge that runner now
> lives inside **`app`**. `backend` reads `JUDGE0_BASE_URL` and calls Judge0 directly, and **nothing calls the
> roadrunner service any more.**
>
> **⚠️ Precision, because the declarations disagree (v2.8 M257x).** *"There is no roadrunner service in
> production"* overstates it: `roadrunner/terraform/main.tf:19` still reads `service_desired_count = 1` and has
> has **not been touched since `84a4b4f` (2025-12-15)** — the commit that first added `terraform/main.tf`,
> seven months before the fold. ⚠️ **This said "`87d8d44` (2026-06-19)" until M257x iter-115, and that is not
> the line's provenance**: `87d8d44` is the repo's HEAD and touches exactly one file,
> `.github/workflows/bump-version.yml` (3 insertions) — it never goes near terraform, so *"not touched since
> it"* is vacuous by construction while the parenthetical presented it as the date of the last touch. The
> subject of *"has not been touched"* is the **line**, not the repo. `git blame -L 19,19 87d8d443 --
> terraform/main.tf` names `84a4b4f`; a file-level `git log` is not line provenance (the file's own most recent
> touch is `e45eb61`, 2026-05-27, a line-11 module-source swap). The corpus's own fenced authority,
> [`platform-migration-status.md`](../architecture/platform-migration-status.md), had this right all along and
> this document never named `84a4b4f` anywhere. The conclusion — *before the fold* — survives; the sha did not.
> Meanwhile the platform has removed it from its own
> clone set: `roadrunner` had a `repos.yml` entry reading *"legacy — folded into app"* as late as `2adcf71`
> (`repos.yml:29-31` **at that ref**), and platform `d11a403` (2026-08-03) **deleted the entry outright** —
> **grade that commit by its diff, not its message.** Its message asserts roadrunner's *"repos.yml clone entry
> was already gone"*; `git show d11a403 -- repos.yml` shows **that very commit** removing `- name: cms`,
> `- name: jobsimulation` *
```

**CITED CONTENT**

```
     7    - port: `/home/devops/ant/app/internal/jobsimulation/ (no agent/ subdir)` | std: `/home/devops/ant/jobsimulation/internal/agent/{report_agent.go,narrative.go,data_store.go,report_agent_test.go}`
     8  - [ ] **(missing) agent-and-new-packages** — The entire `internal/runner/` package is absent from the port. It is the in-process Judge0 sandboxed-code-execution client (RunnerManager: NewRunnerManager, CreateSubmission, GetSubmission, WaitForSubmission with poll budget, plus languages.go mapping runtime->Judge0 language id). Its own doc comment states it replaces the formerly-standalone 'roadrunner' service. The port has no equivalent and still depends on the roadrunner RPC service instead. 3 files (runner.go, languages.go, runner_test.go).
     9    - port: `/home/devops/ant/app/internal/jobsimulation/ (no runner/ subdir)` | std: `/home/devops/ant/jobsimulation/internal/runner/{runner.go,languages.go,runner_test.go}`
    10  - [ ] **(changed) agent-and-new-packages** — Code-execution path diverges materially. In the port, simulator/manager/manager.go injects a roadrunnerv1connect.RoadRunnerServiceClient and calls m.roadrunner.SubmissionPackage / m.roadrunner.SubmissionResult over Connect-RPC (the removed roadrunner microservice). In current main the same manager injects *runner.RunnerManager and calls m.runner.CreateSubmission / m.runner.WaitForSubmission directly against Judge0 in-process. Same code-submission feature, different transport and dependency — re-porting must swap the RPC client for the runner package and drop the roadrunner proto dependency.
    11    - port: `/home/devops/ant/app/internal/jobsimulation/simulator/manager/manager.go (roadrunnerv1connect client; lines ~57-58,214,234,1725,1779,1965)` | std: `/home/devops/ant/jobsimulation/internal/simulator/manager/manager.go (runner.RunnerManager; lines ~38,212,231,1704,1758,1965)`
    12  - [ ] **(missing) agent-and-new-packages** — analytics/aggregated_report.go is absent from the port — the driver that consumes the agent package. It defines the persisted envelope contract (AggregatedReportSchemaVersion, status values ok/insufficient_data/partial_narrative, MinSessionsForNarrative=3), constructs the Bedrock-backed ReportAgent, schedules/runs GenerateAggregatedReport, and persists interview_aggregated_report. Also missing: aggregator.go, aggregator_v2.go, plan_loader.go and their tests
```

## 05-039
- **id**: `B05-039`
- **corpus site**: `corpus/services/roadrunner.md:5-74` (paragraph)
- **citation**: `repos.yml:17-19`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/repos.yml`  (29 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` / ORPHANED — nothing calls this service any more (verified v2.5 M231 KB-6; re-verified v2.7 "july jitter"
> M247 against the CONSOLIDATED platform — the ~386-commit `app` bump).** Code execution moved **in-process into
> jobsimulation** (`jobsimulation/internal/runner/runner.go`, an in-process Judge0 client whose own header comment
> reads *"formerly the standalone 'roadrunner' service"*) — and with the **jobsim-in-app** merge that runner now
> lives inside **`app`**. `backend` reads `JUDGE0_BASE_URL` and calls Judge0 directly, and **nothing calls the
> roadrunner service any more.**
>
> **⚠️ Precision, because the declarations disagree (v2.8 M257x).** *"There is no roadrunner service in
> production"* overstates it: `roadrunner/terraform/main.tf:19` still reads `service_desired_count = 1` and has
> has **not been touched since `84a4b4f` (2025-12-15)** — the commit that first added `terraform/main.tf`,
> seven months before the fold. ⚠️ **This said "`87d8d44` (2026-06-19)" until M257x iter-115, and that is not
> the line's provenance**: `87d8d44` is the repo's HEAD and touches exactly one file,
> `.github/workflows/bump-version.yml` (3 insertions) — it never goes near terraform, so *"not touched since
> it"* is vacuous by construction while the parenthetical presented it as the date of the last touch. The
> subject of *"has not been touched"* is the **line**, not the repo. `git blame -L 19,19 87d8d443 --
> terraform/main.tf` names `84a4b4f`; a file-level `git log` is not line provenance (the file's own most recent
> touch is `e45eb61`, 2026-05-27, a line-11 module-source swap). The corpus's own fenced authority,
> [`platform-migration-status.md`](../architecture/platform-migration-status.md), had this right all along and
> this document never named `84a4b4f` anywhere. The conclusion — *before the fold* — survives; the sha did not.
> Meanwhile the platform has removed it from its own
> clone set: `roadrunner` had a `repos.yml` entry reading *"legacy — folded into app"* as late as `2adcf71`
> (`repos.yml:29-31` **at that ref**), and platform `d11a403` (2026-08-03) **deleted the entry outright** —
> **grade that commit by its diff, not its message.** Its message asserts roadrunner's *"repos.yml clone entry
> was already gone"*; `git show d11a403 -- repos.yml` shows **that very commit** removing `- name: cms`,
> `- name: jobsimulation` *
```

**CITED CONTENT**

```
    14    - name: app
    15      type: go
    16      migrations: true
    17      schema: public
    18    - name: sentinel
    19      type: go
    20      migrations: false
    21  
    22    # Frontend
```

## 05-040
- **id**: `B05-040`
- **corpus site**: `corpus/services/roadrunner.md:5-74` (paragraph)
- **citation**: `docker-compose.yml:83`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> **⚠️ MERGED INTO `app` / ORPHANED — nothing calls this service any more (verified v2.5 M231 KB-6; re-verified v2.7 "july jitter"
> M247 against the CONSOLIDATED platform — the ~386-commit `app` bump).** Code execution moved **in-process into
> jobsimulation** (`jobsimulation/internal/runner/runner.go`, an in-process Judge0 client whose own header comment
> reads *"formerly the standalone 'roadrunner' service"*) — and with the **jobsim-in-app** merge that runner now
> lives inside **`app`**. `backend` reads `JUDGE0_BASE_URL` and calls Judge0 directly, and **nothing calls the
> roadrunner service any more.**
>
> **⚠️ Precision, because the declarations disagree (v2.8 M257x).** *"There is no roadrunner service in
> production"* overstates it: `roadrunner/terraform/main.tf:19` still reads `service_desired_count = 1` and has
> has **not been touched since `84a4b4f` (2025-12-15)** — the commit that first added `terraform/main.tf`,
> seven months before the fold. ⚠️ **This said "`87d8d44` (2026-06-19)" until M257x iter-115, and that is not
> the line's provenance**: `87d8d44` is the repo's HEAD and touches exactly one file,
> `.github/workflows/bump-version.yml` (3 insertions) — it never goes near terraform, so *"not touched since
> it"* is vacuous by construction while the parenthetical presented it as the date of the last touch. The
> subject of *"has not been touched"* is the **line**, not the repo. `git blame -L 19,19 87d8d443 --
> terraform/main.tf` names `84a4b4f`; a file-level `git log` is not line provenance (the file's own most recent
> touch is `e45eb61`, 2026-05-27, a line-11 module-source swap). The corpus's own fenced authority,
> [`platform-migration-status.md`](../architecture/platform-migration-status.md), had this right all along and
> this document never named `84a4b4f` anywhere. The conclusion — *before the fold* — survives; the sha did not.
> Meanwhile the platform has removed it from its own
> clone set: `roadrunner` had a `repos.yml` entry reading *"legacy — folded into app"* as late as `2adcf71`
> (`repos.yml:29-31` **at that ref**), and platform `d11a403` (2026-08-03) **deleted the entry outright** —
> **grade that commit by its diff, not its message.** Its message asserts roadrunner's *"repos.yml clone entry
> was already gone"*; `git show d11a403 -- repos.yml` shows **that very commit** removing `- name: cms`,
> `- name: jobsimulation` *
```

**CITED CONTENT**

```
    80        - AWS_REGION=eu-west-1
    81        - AWS_DEFAULT_REGION=eu-west-1
    82        - STORAGE_S3_BUCKET=production-storage20240826131618541000000005
    83        - STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001
    84        # messenger-in-app (v9.0) and customerio-sync-in-app are folded into this container
    85        # too, but deliberately have NO variables here. Both reach outside the process on a
    86        # stream or a timer — they send mail and rewrite Brevo contacts — so app gates them
```

## 05-041
- **id**: `B05-041`
- **corpus site**: `corpus/services/roadrunner.md:87-87` (bullet)
- **citation**: `cmd/root.go:110`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/cmd/root.go`  (287 lines)

**CLAIMING UNIT**

```md
* **Ports**: **8080 (HTTP — `/_meta` health only), 8081 (Connect-RPC) — the binary's own defaults, and now the only ones there are**: `cmd/root.go:110` `cmp.Or(os.Getenv("PORT"), "8080")`, `:84` `cmp.Or(os.Getenv("RPC_PORT"), "8081")`. The **10400 / 10401** pair quoted throughout this corpus was **compose-supplied by a service that no longer exists**: `docker-compose.yml` set `PORT=10400` (`:298`) / `RPC_PORT=10401` (`:302`) and published `10400:10400` / `10401:10401` (`:291-292`) — **at `2adcf71`**. At `0dab54d` there is no `roadrunner` service, so nothing sets them and nothing is published. **Treat 10400/10401 as historical, not as an address**
```

**CITED CONTENT**

```
   107  			entsql.OpenDB(dialect.Postgres, db),
   108  		))
   109  
   110  		azureAIClient, err := openai.NewAzure(os.Getenv("CMS_AZURE_OPENAI_KEY"), os.Getenv("CMS_AZURE_OPENAI_ENDPOINT_URL"), nil)
   111  		if err != nil {
   112  			return fmt.Errorf("can't init Azure AI client %w", err)
   113  		}
```

## 05-042
- **id**: `B05-042`
- **corpus site**: `corpus/services/roadrunner.md:130-130` (paragraph)
- **citation**: `internal/runner/runner.go:3`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/jobsimulation/internal/runner/runner.go`  (215 lines)

**CLAIMING UNIT**

```md
On completion the worker publishes a `RoadrunnerSubmissionCompleted` event (carrying the Judge0 token) to Redis Streams (`REDIS_STREAMS_INDEX`) via colony pubsub — **and nothing consumes it.** The jobsimulation consumer was **deleted, not moved**: at `jobsimulation 462343b0` the repo's **Go source** contains exactly one `roadrunner` mention (`internal/runner/runner.go:3`, a comment), with no handler and no event reference — **scope that to Go, not to the repo**: `git -C stack-demo/jobsimulation grep -in roadrunner 462343b0` returns **14 lines across 8 files** (5 exact-case across 3), the rest being CHANGELOG and `knowledge/*.md`; in `app`, `internal/jobsimulation/simulator/stream_handlers.go:30-34` states that the roadrunner-submission pubsub event was removed upstream and the code-submission result now arrives as `HandleCodeSubmissionResultTask` on `CodeRunQueue` — *"NOT stream handlers."* Consistent with the **Upstream consumers** bullet under § *Dependencies* below (*"none (orphaned — see the banner at the top)"*) — **named, not pinned:** this said `:124` **below**, and at M257x iter-120 `:124` was **above** this very line and carried no such text.
```

**CITED CONTENT**

```
     1  // Package runner is an in-process client for the Judge0 sandboxed code
     2  // execution API. It powers the code-execution tasks of coding/GenAI
     3  // simulations (formerly the standalone "roadrunner" service).
     4  package runner
     5  
     6  import (
```

## 05-043
- **id**: `B05-043`
- **corpus site**: `corpus/services/roadrunner.md:130-130` (paragraph)
- **citation**: `internal/jobsimulation/simulator/stream_handlers.go:30-34`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimulation/simulator/stream_handlers.go`  (35 lines)

**CLAIMING UNIT**

```md
On completion the worker publishes a `RoadrunnerSubmissionCompleted` event (carrying the Judge0 token) to Redis Streams (`REDIS_STREAMS_INDEX`) via colony pubsub — **and nothing consumes it.** The jobsimulation consumer was **deleted, not moved**: at `jobsimulation 462343b0` the repo's **Go source** contains exactly one `roadrunner` mention (`internal/runner/runner.go:3`, a comment), with no handler and no event reference — **scope that to Go, not to the repo**: `git -C stack-demo/jobsimulation grep -in roadrunner 462343b0` returns **14 lines across 8 files** (5 exact-case across 3), the rest being CHANGELOG and `knowledge/*.md`; in `app`, `internal/jobsimulation/simulator/stream_handlers.go:30-34` states that the roadrunner-submission pubsub event was removed upstream and the code-submission result now arrives as `HandleCodeSubmissionResultTask` on `CodeRunQueue` — *"NOT stream handlers."* Consistent with the **Upstream consumers** bullet under § *Dependencies* below (*"none (orphaned — see the banner at the top)"*) — **named, not pinned:** this said `:124` **below**, and at M257x iter-120 `:124` was **above** this very line and carried no such text.
```

**CITED CONTENT**

```
    27  	}
    28  }
    29  
    30  // NOTE (jobsim-in-app re-sync to v0.253.0): chronos, realtime, and the roadrunner-submission
    31  // pubsub event were all removed upstream ("remove chronos & realtime" PR + the Judge0 asynq migration).
    32  // The session timeout and code-submission result are now delayed/queued Asynq TASKS
    33  // (HandleSessionTimeoutTask on TimersQueue, HandleCodeSubmissionResultTask on CodeRunQueue), registered
    34  // on the worker — NOT stream handlers. So there is no Chronos/Realtime/Roadrunner stream to expose here.
    35  
```

## 05-044
- **id**: `B05-044`
- **corpus site**: `corpus/services/sentinel.md:5-5` (paragraph)
- **citation**: `docker-compose.yml:48`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
Sentinel is the **centralized authorization service** of the platform. Its **only** live caller is **`app`** — including the jobsimulation and cms authz call sites it absorbed in-process — which reaches it over Connect-RPC to check permissions before executing operations. (There are no `cms` or `jobsimulation` containers left to receive the address: platform `d11a403` deleted both compose services along with `roadrunner`, so at `0dab54d` `AUTHORIZATION_ADDRESS` is set in exactly **one** block — backend's, `docker-compose.yml:48`.) **`messenger` is not a caller** — and ⚠️ **the evidence clause has to be past tense, because there is no messenger compose block to read (corrected M257x iter-115).** At platform `0c91421d`, `docker-compose.yml` declares **five** services — `sentinel` (`:5`), `backend` (`:28`), `studio-desk` (`:112`), `next-web-app` (`:143`), `gotenberg` (`:170`) — and `git grep -n messenger 0c91421d -- docker-compose.yml common.yml repos.yml` returns **only comments**. `838d907` (*"drop the storage, messenger and customerio-sync containers"*, 2026-08-05) deleted it. The sentence read *"its compose block sets no `AUTHORIZATION_ADDRESS` and declares no `depends_on: sentinel`"* in the **present** tense, presupposing a block that does not exist — true at `0dab54d` (where the block began at `docker-compose.yml:156`) and silently expired. **Ten other corpus sites already recorded the deletion, two of them in this same file**, neither framed as a retraction of this sentence — one survivor against ten witnesses. What survives, and was re-derived: messenger's Go source imports no authorization client (`git grep "authorization\|AUTHORIZATION_ADDRESS" fa47850d -- '*.go' go.mod` returns one unrelated hit, against `colony` present as a positive control); [`clerk-integration.md`](./clerk-integration.md) says the same ("storage, messenger — no auth"). It wraps **Casbin v3** with a PostgreSQL-backed policy store and a single in-memory enforcer that handles all of Anthropos's authorization patterns.
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
