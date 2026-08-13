# TIER-1 ADJUDICATION BATCH 02 — 44 pairs

Each item: the CLAIMING UNIT (corpus prose) and the CITED CONTENT it points at (source lines,
line-numbered, +-3 lines of context). Answer ONE question per item.

## 02-001
- **id**: `B02-001`
- **corpus site**: `corpus/services/backend.md:3-99` (paragraph)
- **citation**: `app/internal/converter/gotenberg.go:31`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/converter/gotenberg.go`  (54 lines)

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
    28  		return nil, fmt.Errorf("gotenberg: can't finalize multipart body: %w", err)
    29  	}
    30  
    31  	req, err := http.NewRequestWithContext(ctx, "POST", gotenbergURL+"/forms/libreoffice/convert", &body)
    32  	if err != nil {
    33  		return nil, fmt.Errorf("gotenberg: can't create request: %w", err)
    34  	}
```

## 02-002
- **id**: `B02-002`
- **corpus site**: `corpus/services/backend.md:3-99` (paragraph)
- **citation**: `docker-compose.yml:59`
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
    56        - ENVIRONMENT=development
    57        - GOTENBERG_URL=http://gotenberg:3200
    58        - JOBSIMULATION_STREAM=jobsimulation
    59        - JUDGE0_BASE_URL=http://52.48.139.23:2358
    60        - LIVEKIT_AWS_SDK_REGION=eu-central-1
    61        - LIVEKIT_HOST_URL=wss://anthropos-pbvktu3v.livekit.cloud
    62        - LIVEKIT_RECORDINGS_BUCKET_NAME=anthropos-livekit-test
```

## 02-003
- **id**: `B02-003`
- **corpus site**: `corpus/services/backend.md:3-99` (paragraph)
- **citation**: `main.go:325`
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
   322  		logger.Error("can't connect to redis", "error", err)
   323  		return
   324  	}
   325  	pub, err := pubsub.NewPublisher(serviceName, redisClientStream)
   326  	if err != nil {
   327  		logger.Error("can't init event publisher", "error", err)
   328  		return
```

## 02-004
- **id**: `B02-004`
- **corpus site**: `corpus/services/backend.md:3-99` (paragraph)
- **citation**: `subscriber_wiring.go:248`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/subscriber_wiring.go`  (282 lines)

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
   245  	if d.CMSStudio != nil {
   246  		backendSelfSub.AddHandler(pubsub.EventHandler(d.CMSStudio.OrganizationMemberDeletedHandler))
   247  	}
   248  	subs[d.Streams.Backend] = backendSelfSub
   249  
   250  	return subs
   251  }
```

## 02-005
- **id**: `B02-005`
- **corpus site**: `corpus/services/backend.md:3-99` (paragraph)
- **citation**: `cms/terraform/main.tf:39`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/terraform/main.tf`  (192 lines)

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
    36    tags                           = var.tags
    37    aws_region                     = var.aws_region
    38    project                        = local.project
    39    service_desired_count          = 0
    40    service_cpu                    = local.service_cpu
    41    service_memory                 = local.service_memory
    42    health_check_path              = "/_meta"
```

## 02-006
- **id**: `B02-006`
- **corpus site**: `corpus/services/backend.md:3-99` (paragraph)
- **citation**: `jobsimulation/terraform/main.tf:15-22`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/jobsimulation/terraform/main.tf`  (344 lines)

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

## 02-007
- **id**: `B02-007`
- **corpus site**: `corpus/services/backend.md:106-106` (bullet)
- **citation**: `main.go:1185-1228`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
* **Connect-RPC** for inter-service calls (**no external caller is left** — `messenger` was the last, and `838d907` removed its container) — the mux registers five handlers unconditionally (`main.go:1185-1228` @ `app` `b948604` v1.366.0): `UsersService` (`:1187`), `OrganizationsService` (`:1188`), `SkillerService` (`:1196`), `JobSimulationService` (`:1204`) and `lab.v1.LabSessionService` (`:1228`), plus **`CMSService` only when the Directus edge is configured** (`if cmsRPCServer != nil`, `:1212-1214`).
```

**CITED CONTENT**

```
  1182  		Ent:                 ent,
  1183  		Skiller:             cmsSkiller,
  1184  		JobSimulationClient: jobsimDj.RPCServer,
  1185  		Studio:              cmsManagers.Studio,
  1186  		Asynq:               cmsAsynq,
  1187  		Pub:                 cmsPub,
  1188  		Storage:             cmsStorage,
  1189  		AiVideo:             cmsManagers.AiVideo,
  1190  	}
  1191  	// cms-in-app M807/M809: one cms RPC server backs BOTH the served CMSService handler
  1192  	// (M807 mux, for external messenger) and the in-process client switch (M809 internal
  1193  	// caller cutover). It satisfies both the connect Handler and Client interfaces.
  1194  	cmsRPCServer = cmsrpcsrv.NewRPCServer(cmsManagers.Directus, cmsManagers.AiVideo)
  1195  	// cms-in-app: the inbound Directus webhook now lands on APP (POST /api/webhook/directus) —
  1196  	// at release the Directus Flows re-point here so cache-clear / re-index / clone / ai-video
  1197  	// creation fire in-process (no traffic to the standalone cms). Authenticated by
  1198  	// DIRECTUS_WEBHOOK_SECRET (M809b M-2, fail-closed).
  1199  	cmsWebhookHandler = cmswebhooks.Handler(os.Getenv("DIRECTUS_WEBHOOK_SECRET"), cmsManagers.Directus, cmsPub, cmsManagers.AiVideo)
  1200  	// M809 internal caller cutover: app owns cms content in-process. jobsim/skillpath/JSManager
  1201  	// read cms via the in-process RPC server instead of over the wire — no traffic to the
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
  1215  	})
  1216  
  1217  	graphHandler := graph.NewHandler(
  
```

## 02-008
- **id**: `B02-008`
- **corpus site**: `corpus/services/backend.md:110-110` (paragraph)
- **citation**: `app/CLAUDE.md:109`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/CLAUDE.md`  (357 lines)

**CLAIMING UNIT**

```md
  > **⚠️ `app`'s OWN docs still list it** (`app/CLAUDE.md:109`, `app/knowledge/architecture.md:28` — re-derived at `app` **`ad9f3c49`**, which was `origin/main` on 2026-08-06; both anchors are unchanged from `2035f9a4`, the ref this corpus used to call "origin/main" and which is now 5 commits behind. The CLAUDE.md line was `:80` when this was first measured, so re-find the sentence rather than trusting the offset), which is where this corpus previously got the claim. That is Trap C in [`../ops/platform-alignment.md`](../ops/platform-alignment.md) — *the platform's planning docs lag its own code*. **Grade against `main.go`, not against `app/CLAUDE.md`.**
```

**CITED CONTENT**

```
   106  
   107  The service runs 4 concurrent servers from `main.go`:
   108  1. **Web server** (port 8080) — Echo HTTP with GraphQL endpoint and REST routes (incl. `POST /api/webhook/directus`, which fails closed without `DIRECTUS_WEBHOOK_SECRET`)
   109  2. **RPC server** (port 8081) — one Connect-RPC mux carrying `BackendUsersService`, `BackendOrganizationsService`, `SkillerService`, `SkillPathSessionService`, `JobSimulationService`, `CMSService` and `lab.v1.LabSessionService`. Runs with a **60s write timeout** — the ported skiller RAG/LLM calls can exceed the old 10s default
   110  3. **Meta server** (port 8083) — Health checks, version info, Asynq inspector
   111  4. **Workers** — Asynq background job processors: the app worker plus the ported skiller, coursebuilder and cms workers
   112  
```

## 02-009
- **id**: `B02-009`
- **corpus site**: `corpus/services/backend.md:110-110` (paragraph)
- **citation**: `app/knowledge/architecture.md:28`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/knowledge/architecture.md`  (131 lines)

**CLAIMING UNIT**

```md
  > **⚠️ `app`'s OWN docs still list it** (`app/CLAUDE.md:109`, `app/knowledge/architecture.md:28` — re-derived at `app` **`ad9f3c49`**, which was `origin/main` on 2026-08-06; both anchors are unchanged from `2035f9a4`, the ref this corpus used to call "origin/main" and which is now 5 commits behind. The CLAUDE.md line was `:80` when this was first measured, so re-find the sentence rather than trusting the offset), which is where this corpus previously got the claim. That is Trap C in [`../ops/platform-alignment.md`](../ops/platform-alignment.md) — *the platform's planning docs lag its own code*. **Grade against `main.go`, not against `app/CLAUDE.md`.**
```

**CITED CONTENT**

```
    25  | Server | Port | Protocol | Purpose |
    26  |--------|------|----------|---------|
    27  | Web | 8080 | HTTP (Echo) | GraphQL endpoint + REST webhook routes |
    28  | RPC | 8081 | Connect-RPC | One mux: `BackendUsersService`, `BackendOrganizationsService`, `SkillerService`, `SkillPathSessionService`, `JobSimulationService`, `CMSService`, `lab.v1.LabSessionService`. Runs with a **60s write timeout** — the ported skiller RAG/LLM calls can exceed the old 10s default |
    29  | Meta | 8083 | HTTP | Health check, version, Asynq inspector |
    30  | Worker | — | Asynq (Redis) | Background job processor (10 concurrent workers) |
    31  
```

## 02-010
- **id**: `B02-010`
- **corpus site**: `corpus/services/backend.md:120-120` (bullet)
- **citation**: `main.go:743-746`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
* **AI Labs LabSession** (`internal/labs/session`; siblings `internal/labs/labsapi`, `internal/labs/adapter`, `internal/labs/catalog`) — Connect-RPC `lab.v1.LabSessionService` (Create/Get/List/Cancel/ReportEvent) plus a `lab_sessions` Ent table. The labs-api client is wired **only when `LABS_API_URL` is set** (`main.go:743-746` @ `app` `b948604` v1.366.0); with it unset — the usual local/demo case — Create persists a session row without booting a VM and Cancel marks the row cancelled without calling labs-api (see Recent Feature Additions). It is NOT unconditionally nil.
```

**CITED CONTENT**

```
   740  	// existing cms/jobsim/authn/authz — no new outbound client (E.2).
   741  	//
   742  	// Its own publisher targets SKILLPATH_STREAM (NOT app's default `pub`, which publishes to
   743  	// the "backend" stream): skillpath emits SkillPathSessionUpdated / ChapterStepSessionCompleted,
   744  	// and app's existing internal/skillpaths consumer subscribes SKILLPATH_STREAM — so keeping
   745  	// the producer on that stream preserves the in-process Redis loopback (E.3/OD-4).
   746  	skillPathPub, err := pubsub.NewPublisher(os.Getenv("SKILLPATH_STREAM"), redisClientStream)
   747  	if err != nil {
   748  		logger.Error("can't init skillpath event publisher", "error", err)
   749  		return
```

## 02-011
- **id**: `B02-011`
- **corpus site**: `corpus/services/backend.md:144-194` (bullet)
- **citation**: `app/knowledge/service-dependencies.md:52`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/knowledge/service-dependencies.md`  (122 lines)

**CLAIMING UNIT**

```md
- **RPC re-pointed, then un-set** — the `SkillerService` Connect-RPC surface is served **by app itself**
  (`internal/rpc/skillerrpc/`). Consumers kept the env var `SKILLER_RPC_ADDR`, re-pointed at
  `http://backend:8083` — a value it held **before** `d11a403` as well as after, which is why that commit
  did not re-point it (M257x iter-115). **That count was always ref-relative, and it has now reached zero:** four
  occurrences in `docker-compose.yml` @ platform `0808b92` (the ref this fact-sheet was first ground
  against — `backend`, `jobsimulation`, `cms` and `messenger` each carried one); **one** @ `0dab54d`,
  messenger's, after `d11a403` deleted the `jobsimulation` and `cms` blocks and dropped it from
  `backend`, which no longer addresses a surface it serves itself (**note `d11a403` did not *re-point*
  this variable — `SKILLER_RPC_ADDR` already read `http://backend:8083` at `d11a403^`; the two it
  re-pointed were `CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`**); and **none** @ `0c91421`, because
  `838d907` deleted the `messenger` block that held the last one. **No compose file sets any
  `*_RPC_ADDR` today.** **No terraform in the clone set names `http://backend.internal.anthropos:8081`,
  and this doc no longer asserts that any does.** Measured 2026-08-06 by two independent mechanisms:
  `git grep` at each clone's own HEAD over the **44** tracked `.tf` files in the 13-repo `stack-demo`
  clone set → **0 files**; a raw filesystem `find … -name '*.tf' | grep` over the same working trees,
  **59** files → **0** (positive control on `service_discovery_namespace_id`: 25 files). **The literal
  occurs SIX times in the clone set, none of them in terraform** — and the *"only occurrence"* wording this
  passage used to carry was **self-refuting at its own stated scope**, which is why the count is now
  spelled out: **1** in `app` (`app/knowledge/service-dependencies.md:52` @ `ad9f3c49`) and **5** in
  `rosetta-extensions` @ the pinned `09d06070` — four inside this corpus's own frozen test fixtures
  (the merge-banner **RPC** bullet and the *Interface Discovery* **RPC** bullet of
  `stack-core/tests/fixtures/repair_leak/{pre,post}/corpus/services/cms.md` — **named, not numbered**:
  those are FROZEN copies of a corpus file, and a citation of the form `…/corpus/services/cms.md:NNN` is
  resolved against the **live** corpus by `anchor_construct_guard`, so a cor
```

**CITED CONTENT**

```
    49  >
    50  > **There are no external callers of app's RPC mux left.** `messenger` was the last one — it used to
    51  > reach the users, cms, jobsimulation and skiller surfaces at
    52  > `http://backend.internal.anthropos:8081`, and folding it in at v9.0 closed that edge. The mux is
    53  > kept because it is how the in-process domains are wired, not because something outside dials it.
    54  > App also keeps emitting `JOBSIMULATION_STREAM` as an in-process loopback — it feeds the
    55  > real consumers (XP/skills/quota/assignment link/AI Readiness); the `LocalJobsimulationSession`
```

## 02-012
- **id**: `B02-012`
- **corpus site**: `corpus/services/backend.md:144-194` (bullet)
- **citation**: `stack-core/tests/test_platform_predicate_guard.py:435`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/stack-core/tests/test_platform_predicate_guard.py`  (2438 lines)

**CLAIMING UNIT**

```md
- **RPC re-pointed, then un-set** — the `SkillerService` Connect-RPC surface is served **by app itself**
  (`internal/rpc/skillerrpc/`). Consumers kept the env var `SKILLER_RPC_ADDR`, re-pointed at
  `http://backend:8083` — a value it held **before** `d11a403` as well as after, which is why that commit
  did not re-point it (M257x iter-115). **That count was always ref-relative, and it has now reached zero:** four
  occurrences in `docker-compose.yml` @ platform `0808b92` (the ref this fact-sheet was first ground
  against — `backend`, `jobsimulation`, `cms` and `messenger` each carried one); **one** @ `0dab54d`,
  messenger's, after `d11a403` deleted the `jobsimulation` and `cms` blocks and dropped it from
  `backend`, which no longer addresses a surface it serves itself (**note `d11a403` did not *re-point*
  this variable — `SKILLER_RPC_ADDR` already read `http://backend:8083` at `d11a403^`; the two it
  re-pointed were `CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`**); and **none** @ `0c91421`, because
  `838d907` deleted the `messenger` block that held the last one. **No compose file sets any
  `*_RPC_ADDR` today.** **No terraform in the clone set names `http://backend.internal.anthropos:8081`,
  and this doc no longer asserts that any does.** Measured 2026-08-06 by two independent mechanisms:
  `git grep` at each clone's own HEAD over the **44** tracked `.tf` files in the 13-repo `stack-demo`
  clone set → **0 files**; a raw filesystem `find … -name '*.tf' | grep` over the same working trees,
  **59** files → **0** (positive control on `service_discovery_namespace_id`: 25 files). **The literal
  occurs SIX times in the clone set, none of them in terraform** — and the *"only occurrence"* wording this
  passage used to carry was **self-refuting at its own stated scope**, which is why the count is now
  spelled out: **1** in `app` (`app/knowledge/service-dependencies.md:52` @ `ad9f3c49`) and **5** in
  `rosetta-extensions` @ the pinned `09d06070` — four inside this corpus's own frozen test fixtures
  (the merge-banner **RPC** bullet and the *Interface Discovery* **RPC** bullet of
  `stack-core/tests/fixtures/repair_leak/{pre,post}/corpus/services/cms.md` — **named, not numbered**:
  those are FROZEN copies of a corpus file, and a citation of the form `…/corpus/services/cms.md:NNN` is
  resolved against the **live** corpus by `anchor_construct_guard`, so a cor
```

**CITED CONTENT**

```
   432          self.assertNotIn("G5 wrong-target-set", kinds(res.findings))
   433  
   434      def test_a_production_fqdn_is_not_a_compose_claim(self) -> None:
   435          body = "In production it is `CMS_RPC_ADDR=http://backend.internal.anthropos:8081`.\n"
   436          corpus = write_corpus(self.root, body)
   437          res = G.check(corpus, self.platform, app_root=None)
   438          self.assertNotIn("G4 stale-address", kinds(res.findings))
```

## 02-013
- **id**: `B02-013`
- **corpus site**: `corpus/services/backend.md:144-194` (bullet)
- **citation**: `app/terraform/locals.tf:6`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/locals.tf`  (22 lines)

**CLAIMING UNIT**

```md
- **RPC re-pointed, then un-set** — the `SkillerService` Connect-RPC surface is served **by app itself**
  (`internal/rpc/skillerrpc/`). Consumers kept the env var `SKILLER_RPC_ADDR`, re-pointed at
  `http://backend:8083` — a value it held **before** `d11a403` as well as after, which is why that commit
  did not re-point it (M257x iter-115). **That count was always ref-relative, and it has now reached zero:** four
  occurrences in `docker-compose.yml` @ platform `0808b92` (the ref this fact-sheet was first ground
  against — `backend`, `jobsimulation`, `cms` and `messenger` each carried one); **one** @ `0dab54d`,
  messenger's, after `d11a403` deleted the `jobsimulation` and `cms` blocks and dropped it from
  `backend`, which no longer addresses a surface it serves itself (**note `d11a403` did not *re-point*
  this variable — `SKILLER_RPC_ADDR` already read `http://backend:8083` at `d11a403^`; the two it
  re-pointed were `CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`**); and **none** @ `0c91421`, because
  `838d907` deleted the `messenger` block that held the last one. **No compose file sets any
  `*_RPC_ADDR` today.** **No terraform in the clone set names `http://backend.internal.anthropos:8081`,
  and this doc no longer asserts that any does.** Measured 2026-08-06 by two independent mechanisms:
  `git grep` at each clone's own HEAD over the **44** tracked `.tf` files in the 13-repo `stack-demo`
  clone set → **0 files**; a raw filesystem `find … -name '*.tf' | grep` over the same working trees,
  **59** files → **0** (positive control on `service_discovery_namespace_id`: 25 files). **The literal
  occurs SIX times in the clone set, none of them in terraform** — and the *"only occurrence"* wording this
  passage used to carry was **self-refuting at its own stated scope**, which is why the count is now
  spelled out: **1** in `app` (`app/knowledge/service-dependencies.md:52` @ `ad9f3c49`) and **5** in
  `rosetta-extensions` @ the pinned `09d06070` — four inside this corpus's own frozen test fixtures
  (the merge-banner **RPC** bullet and the *Interface Discovery* **RPC** bullet of
  `stack-core/tests/fixtures/repair_leak/{pre,post}/corpus/services/cms.md` — **named, not numbered**:
  those are FROZEN copies of a corpus file, and a citation of the form `…/corpus/services/cms.md:NNN` is
  resolved against the **live** corpus by `anchor_construct_guard`, so a cor
```

**CITED CONTENT**

```
     3      Environment = "${var.environment}"
     4      Engine      = "terraform"
     5    }
     6    project   = "backend"
     7    port      = 8080
     8    rpc_port  = 8081
     9    meta_port = 8083
```

## 02-014
- **id**: `B02-014`
- **corpus site**: `corpus/services/backend.md:144-194` (bullet)
- **citation**: `app/terraform/main.tf:58`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/main.tf`  (787 lines)

**CLAIMING UNIT**

```md
- **RPC re-pointed, then un-set** — the `SkillerService` Connect-RPC surface is served **by app itself**
  (`internal/rpc/skillerrpc/`). Consumers kept the env var `SKILLER_RPC_ADDR`, re-pointed at
  `http://backend:8083` — a value it held **before** `d11a403` as well as after, which is why that commit
  did not re-point it (M257x iter-115). **That count was always ref-relative, and it has now reached zero:** four
  occurrences in `docker-compose.yml` @ platform `0808b92` (the ref this fact-sheet was first ground
  against — `backend`, `jobsimulation`, `cms` and `messenger` each carried one); **one** @ `0dab54d`,
  messenger's, after `d11a403` deleted the `jobsimulation` and `cms` blocks and dropped it from
  `backend`, which no longer addresses a surface it serves itself (**note `d11a403` did not *re-point*
  this variable — `SKILLER_RPC_ADDR` already read `http://backend:8083` at `d11a403^`; the two it
  re-pointed were `CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`**); and **none** @ `0c91421`, because
  `838d907` deleted the `messenger` block that held the last one. **No compose file sets any
  `*_RPC_ADDR` today.** **No terraform in the clone set names `http://backend.internal.anthropos:8081`,
  and this doc no longer asserts that any does.** Measured 2026-08-06 by two independent mechanisms:
  `git grep` at each clone's own HEAD over the **44** tracked `.tf` files in the 13-repo `stack-demo`
  clone set → **0 files**; a raw filesystem `find … -name '*.tf' | grep` over the same working trees,
  **59** files → **0** (positive control on `service_discovery_namespace_id`: 25 files). **The literal
  occurs SIX times in the clone set, none of them in terraform** — and the *"only occurrence"* wording this
  passage used to carry was **self-refuting at its own stated scope**, which is why the count is now
  spelled out: **1** in `app` (`app/knowledge/service-dependencies.md:52` @ `ad9f3c49`) and **5** in
  `rosetta-extensions` @ the pinned `09d06070` — four inside this corpus's own frozen test fixtures
  (the merge-banner **RPC** bullet and the *Interface Discovery* **RPC** bullet of
  `stack-core/tests/fixtures/repair_leak/{pre,post}/corpus/services/cms.md` — **named, not numbered**:
  those are FROZEN copies of a corpus file, and a citation of the form `…/corpus/services/cms.md:NNN` is
  resolved against the **live** corpus by `anchor_construct_guard`, so a cor
```

**CITED CONTENT**

```
    55  // Nothing has to be provisioned before an apply. The `sentinel` role that DSN
    56  // carries owns the schema and the table, so CREATE on
    57  // sentinel.atlas_schema_revisions needs no GRANT from anyone. No CREATE ROLE, no
    58  // GRANT, no ALTER ROLE, no ordering constraint between this and anything else.
    59  //
    60  // ACCEPTED TRADE-OFF, stated rather than buried: app therefore runs as the
    61  // SCHEMA OWNER, with DDL rights, at runtime — it can DROP TABLE casbin_rules.
```

## 02-015
- **id**: `B02-015`
- **corpus site**: `corpus/services/backend.md:144-194` (bullet)
- **citation**: `messenger/terraform/variables.tf:77`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/messenger/terraform/variables.tf`  (108 lines)

**CLAIMING UNIT**

```md
- **RPC re-pointed, then un-set** — the `SkillerService` Connect-RPC surface is served **by app itself**
  (`internal/rpc/skillerrpc/`). Consumers kept the env var `SKILLER_RPC_ADDR`, re-pointed at
  `http://backend:8083` — a value it held **before** `d11a403` as well as after, which is why that commit
  did not re-point it (M257x iter-115). **That count was always ref-relative, and it has now reached zero:** four
  occurrences in `docker-compose.yml` @ platform `0808b92` (the ref this fact-sheet was first ground
  against — `backend`, `jobsimulation`, `cms` and `messenger` each carried one); **one** @ `0dab54d`,
  messenger's, after `d11a403` deleted the `jobsimulation` and `cms` blocks and dropped it from
  `backend`, which no longer addresses a surface it serves itself (**note `d11a403` did not *re-point*
  this variable — `SKILLER_RPC_ADDR` already read `http://backend:8083` at `d11a403^`; the two it
  re-pointed were `CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`**); and **none** @ `0c91421`, because
  `838d907` deleted the `messenger` block that held the last one. **No compose file sets any
  `*_RPC_ADDR` today.** **No terraform in the clone set names `http://backend.internal.anthropos:8081`,
  and this doc no longer asserts that any does.** Measured 2026-08-06 by two independent mechanisms:
  `git grep` at each clone's own HEAD over the **44** tracked `.tf` files in the 13-repo `stack-demo`
  clone set → **0 files**; a raw filesystem `find … -name '*.tf' | grep` over the same working trees,
  **59** files → **0** (positive control on `service_discovery_namespace_id`: 25 files). **The literal
  occurs SIX times in the clone set, none of them in terraform** — and the *"only occurrence"* wording this
  passage used to carry was **self-refuting at its own stated scope**, which is why the count is now
  spelled out: **1** in `app` (`app/knowledge/service-dependencies.md:52` @ `ad9f3c49`) and **5** in
  `rosetta-extensions` @ the pinned `09d06070` — four inside this corpus's own frozen test fixtures
  (the merge-banner **RPC** bullet and the *Interface Discovery* **RPC** bullet of
  `stack-core/tests/fixtures/repair_leak/{pre,post}/corpus/services/cms.md` — **named, not numbered**:
  those are FROZEN copies of a corpus file, and a citation of the form `…/corpus/services/cms.md:NNN` is
  resolved against the **live** corpus by `anchor_construct_guard`, so a cor
```

**CITED CONTENT**

```
    74    description = "Redis Streams Index"
    75  }
    76  
    77  variable "cms_rpc_address" {
    78    type        = string
    79    description = "CMS RPC Address"
    80  }
```

## 02-016
- **id**: `B02-016`
- **corpus site**: `corpus/services/backend.md:144-194` (bullet)
- **citation**: `app/terraform/variables.tf:197`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/variables.tf`  (764 lines)

**CLAIMING UNIT**

```md
- **RPC re-pointed, then un-set** — the `SkillerService` Connect-RPC surface is served **by app itself**
  (`internal/rpc/skillerrpc/`). Consumers kept the env var `SKILLER_RPC_ADDR`, re-pointed at
  `http://backend:8083` — a value it held **before** `d11a403` as well as after, which is why that commit
  did not re-point it (M257x iter-115). **That count was always ref-relative, and it has now reached zero:** four
  occurrences in `docker-compose.yml` @ platform `0808b92` (the ref this fact-sheet was first ground
  against — `backend`, `jobsimulation`, `cms` and `messenger` each carried one); **one** @ `0dab54d`,
  messenger's, after `d11a403` deleted the `jobsimulation` and `cms` blocks and dropped it from
  `backend`, which no longer addresses a surface it serves itself (**note `d11a403` did not *re-point*
  this variable — `SKILLER_RPC_ADDR` already read `http://backend:8083` at `d11a403^`; the two it
  re-pointed were `CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`**); and **none** @ `0c91421`, because
  `838d907` deleted the `messenger` block that held the last one. **No compose file sets any
  `*_RPC_ADDR` today.** **No terraform in the clone set names `http://backend.internal.anthropos:8081`,
  and this doc no longer asserts that any does.** Measured 2026-08-06 by two independent mechanisms:
  `git grep` at each clone's own HEAD over the **44** tracked `.tf` files in the 13-repo `stack-demo`
  clone set → **0 files**; a raw filesystem `find … -name '*.tf' | grep` over the same working trees,
  **59** files → **0** (positive control on `service_discovery_namespace_id`: 25 files). **The literal
  occurs SIX times in the clone set, none of them in terraform** — and the *"only occurrence"* wording this
  passage used to carry was **self-refuting at its own stated scope**, which is why the count is now
  spelled out: **1** in `app` (`app/knowledge/service-dependencies.md:52` @ `ad9f3c49`) and **5** in
  `rosetta-extensions` @ the pinned `09d06070` — four inside this corpus's own frozen test fixtures
  (the merge-banner **RPC** bullet and the *Interface Discovery* **RPC** bullet of
  `stack-core/tests/fixtures/repair_leak/{pre,post}/corpus/services/cms.md` — **named, not numbered**:
  those are FROZEN copies of a corpus file, and a citation of the form `…/corpus/services/cms.md:NNN` is
  resolved against the **live** corpus by `anchor_construct_guard`, so a cor
```

**CITED CONTENT**

```
   194      check for years. Local dev shows the same shape outright:
   195      `postgresql://…/postgres?search_path=sentinel&sslmode=disable`.
   196  
   197      The validation below therefore fails the plan rather than letting a DSN
   198      without it through. A missing `search_path=sentinel` does not error at
   199      runtime — it silently resolves `casbin_rules` against `public`, where the
   200      table does not exist or, worse, might one day. A loud plan failure is the
```

## 02-017
- **id**: `B02-017`
- **corpus site**: `corpus/services/backend.md:207-216` (bullet)
- **citation**: `main.go:1276`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
- **No skiller container / repo / schema search-path.** Not in `repos.yml` or `docker-compose.yml`; the app
  DB connection uses the default `public` search_path (no `search_path=skiller`); `app` subscribes to the
  `skiller` Redis stream — but only as a **consumer**: nothing in `app` publishes to it. `main.go:1276`
  @ `b948604f` (the demo pin) is a direct `AddSubscriber`; at `app` **`ad9f3c49`** — `origin/main` on
  2026-08-06, 5 commits past the `2035f9a4` this corpus used to *label* "origin/main" — that
  registration has moved into the map-built subscriber set (`subs[d.Streams.Skiller]`,
  `subscriber_wiring.go:209`), applied by the single loop at `main.go:1579-1581`. **No `NewPublisher`
  names `SKILLER_STREAM` at any of those refs** — re-derived at `b948604f`, `2035f9a4` and `ad9f3c49`.
  The producer was the standalone skiller service and went with it. See the
  Redis Streams section below.
```

**CITED CONTENT**

```
  1273  		cmsWebhookHandler,
  1274  		academyContentManager,
  1275  		academyEmbeddingsManager,
  1276  		apiKeyManager,
  1277  		invitationManager,
  1278  		labsDeps,
  1279  		courseBuilderDeps,
```

## 02-018
- **id**: `B02-018`
- **corpus site**: `corpus/services/backend.md:207-216` (bullet)
- **citation**: `subscriber_wiring.go:209`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/subscriber_wiring.go`  (282 lines)

**CLAIMING UNIT**

```md
- **No skiller container / repo / schema search-path.** Not in `repos.yml` or `docker-compose.yml`; the app
  DB connection uses the default `public` search_path (no `search_path=skiller`); `app` subscribes to the
  `skiller` Redis stream — but only as a **consumer**: nothing in `app` publishes to it. `main.go:1276`
  @ `b948604f` (the demo pin) is a direct `AddSubscriber`; at `app` **`ad9f3c49`** — `origin/main` on
  2026-08-06, 5 commits past the `2035f9a4` this corpus used to *label* "origin/main" — that
  registration has moved into the map-built subscriber set (`subs[d.Streams.Skiller]`,
  `subscriber_wiring.go:209`), applied by the single loop at `main.go:1579-1581`. **No `NewPublisher`
  names `SKILLER_STREAM` at any of those refs** — re-derived at `b948604f`, `2035f9a4` and `ad9f3c49`.
  The producer was the standalone skiller service and went with it. See the
  Redis Streams section below.
```

**CITED CONTENT**

```
   206  	// SkillPath
   207  	subs[d.Streams.SkillPath] = d.SkillPath.SkillPathSubscriber()
   208  	// Skiller
   209  	subs[d.Streams.Skiller] = d.Skiller.SkillerSubscriber()
   210  
   211  	// JobSimulation stream — consumed by the jobsimulation domain + skillpath + the ported
   212  	// jobsim engine (streamrouter composes those three onto one subscriber),.
```

## 02-019
- **id**: `B02-019`
- **corpus site**: `corpus/services/backend.md:207-216` (bullet)
- **citation**: `main.go:1579-1581`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
- **No skiller container / repo / schema search-path.** Not in `repos.yml` or `docker-compose.yml`; the app
  DB connection uses the default `public` search_path (no `search_path=skiller`); `app` subscribes to the
  `skiller` Redis stream — but only as a **consumer**: nothing in `app` publishes to it. `main.go:1276`
  @ `b948604f` (the demo pin) is a direct `AddSubscriber`; at `app` **`ad9f3c49`** — `origin/main` on
  2026-08-06, 5 commits past the `2035f9a4` this corpus used to *label* "origin/main" — that
  registration has moved into the map-built subscriber set (`subs[d.Streams.Skiller]`,
  `subscriber_wiring.go:209`), applied by the single loop at `main.go:1579-1581`. **No `NewPublisher`
  names `SKILLER_STREAM` at any of those refs** — re-derived at `b948604f`, `2035f9a4` and `ad9f3c49`.
  The producer was the standalone skiller service and went with it. See the
  Redis Streams section below.
```

**CITED CONTENT**

```
  1576  	// map makes that unrepresentable: one subscriber per stream, by construction. Order is
  1577  	// irrelevant (the destination is a map, and colony's own init() already iterates it in
  1578  	// randomized order) — proved by TestRegistrationOrderDoesNotMatter.
  1579  	for stream, sub := range streamSubscribers {
  1580  		subServer.AddSubscriber(stream, sub)
  1581  	}
  1582  
  1583  	wg.Go(func() {
  1584  		defer cancelServerContext()
```

## 02-020
- **id**: `B02-020`
- **corpus site**: `corpus/services/backend.md:240-240` (bullet)
- **citation**: `docker-compose.yml:110`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
* **Profile**: `core` (the default), `backend`, `all` — `profiles: [core, backend, all]` (`docker-compose.yml:110`, derived from `docker-compose.yml` @ platform `0c91421`; it was `:100` at `0dab54d`, and compose clean-ups move it). The default profile is `core`, not `graphql`: `0dab54d` renamed it. Corrected M257x iter-68
```

**CITED CONTENT**

```
   107          condition: service_healthy
   108        sentinel:
   109          condition: service_started
   110      profiles: [core, backend, all]
   111  
   112    studio-desk:
   113      build:
```

## 02-021
- **id**: `B02-021`
- **corpus site**: `corpus/services/backend.md:314-314` (bullet)
- **citation**: `main.go:743-746`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
* **AI Labs LabSession** (Phase B PR 2, #896): Connect-RPC `lab.v1.LabSessionService` (Create/Get/List/Cancel/ReportEvent) plus a new `lab_sessions` Ent table — `id` supplied by labs-api as a 12-char hex (not a UUID); `user_id`, `organization_id` (optional — empty for individual payers), `template`, `mode` (test/build/teach), `status` (booting/ready/grading/stopped/failed/cancelled), `budget_usd`/`spend_usd`/`total_tokens`, `started_at`/`stopped_at`, `grade_result` JSON. Registered as a third RPC handler in `main.go` after Users and Organizations. **The real HTTP client has since LANDED** and is wired conditionally — `main.go:743-746` @ `app` `b948604` v1.366.0: `if labsAPIURL := os.Getenv("LABS_API_URL"); labsAPIURL != "" { labsAPI = adapter.New(labsapi.New(...)) }`, backed by real `internal/labs/labsapi/` + `internal/labs/adapter/` packages. `LabsAPIClient` is nil **only** on the unset-`LABS_API_URL` local-dev path, where Create persists the LabSession row without booting a VM (no `ide_url`/`preview_url`) and Cancel marks the row cancelled without calling labs-api.
```

**CITED CONTENT**

```
   740  	// existing cms/jobsim/authn/authz — no new outbound client (E.2).
   741  	//
   742  	// Its own publisher targets SKILLPATH_STREAM (NOT app's default `pub`, which publishes to
   743  	// the "backend" stream): skillpath emits SkillPathSessionUpdated / ChapterStepSessionCompleted,
   744  	// and app's existing internal/skillpaths consumer subscribes SKILLPATH_STREAM — so keeping
   745  	// the producer on that stream preserves the in-process Redis loopback (E.3/OD-4).
   746  	skillPathPub, err := pubsub.NewPublisher(os.Getenv("SKILLPATH_STREAM"), redisClientStream)
   747  	if err != nil {
   748  		logger.Error("can't init skillpath event publisher", "error", err)
   749  		return
```

## 02-022
- **id**: `B02-022`
- **corpus site**: `corpus/services/backend.md:319-319` (bullet)
- **citation**: `docker-compose.yml:48`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
* **Connect-RPC**: `rpc.go` is the top-level wire-up. Look there for the implemented services. **There is no external caller left, and no address to be one with.** `messenger` was the last, and it reached four surfaces — `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR`, `SKILLER_RPC_ADDR` — all reading `http://backend:8083` — **but only two of them because `d11a403` moved them** (`CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`; the other two already held that value at `d11a403^`, and `BACKEND_USERS_RPC_ADDR` never addressed anything but `backend` from `3e85fce` on). *"…pointed at `http://backend:8083` by `d11a403`"* over all four was the false form and is corrected at M257x iter-115; M809 did land, on two variables. All four were set on messenger's own compose block and nowhere else, under compose's own comment *"cms + jobsimulation are folded into app: all four RPC edges are the one backend mux"*. `838d907` deleted that block, so **compose now sets zero `*_RPC_ADDR` variables** and the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is `backend → sentinel` (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48` @ platform `0c91421`). **It is not the only cross-process edge:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The earlier two-of-four split (`http://cms:8091` / `http://jobsimulation:8401`, true at platform `2adcf71`) is history twice over. `app`'s own source comment still says *"additive + DORMANT: external callers (messenger) keep hitting the standalone cms via `CMS_RPC_ADDR` **until the M809 re-point**"* (`app/main.go:1205-1211` @ `b948604` v1.366.0) — **that comment is now stale in `app`**; grade the address against compose, not against the comment. **This bullet used to close by naming a production terraform address. That assertion is DROPPED, not softened.** Measured 2026-08-06 by two mechanisms — `git grep` at each clone's own HEAD over the 44 tracked `.tf` files in the 13-repo `stack-demo` clone set, and a raw filesystem grep over the 59 `.tf` files in the same working trees — **no terraform anywhere in the clone set nam
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

## 02-023
- **id**: `B02-023`
- **corpus site**: `corpus/services/backend.md:319-319` (bullet)
- **citation**: `docker-compose.yml:57`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
* **Connect-RPC**: `rpc.go` is the top-level wire-up. Look there for the implemented services. **There is no external caller left, and no address to be one with.** `messenger` was the last, and it reached four surfaces — `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR`, `SKILLER_RPC_ADDR` — all reading `http://backend:8083` — **but only two of them because `d11a403` moved them** (`CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`; the other two already held that value at `d11a403^`, and `BACKEND_USERS_RPC_ADDR` never addressed anything but `backend` from `3e85fce` on). *"…pointed at `http://backend:8083` by `d11a403`"* over all four was the false form and is corrected at M257x iter-115; M809 did land, on two variables. All four were set on messenger's own compose block and nowhere else, under compose's own comment *"cms + jobsimulation are folded into app: all four RPC edges are the one backend mux"*. `838d907` deleted that block, so **compose now sets zero `*_RPC_ADDR` variables** and the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is `backend → sentinel` (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48` @ platform `0c91421`). **It is not the only cross-process edge:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The earlier two-of-four split (`http://cms:8091` / `http://jobsimulation:8401`, true at platform `2adcf71`) is history twice over. `app`'s own source comment still says *"additive + DORMANT: external callers (messenger) keep hitting the standalone cms via `CMS_RPC_ADDR` **until the M809 re-point**"* (`app/main.go:1205-1211` @ `b948604` v1.366.0) — **that comment is now stale in `app`**; grade the address against compose, not against the comment. **This bullet used to close by naming a production terraform address. That assertion is DROPPED, not softened.** Measured 2026-08-06 by two mechanisms — `git grep` at each clone's own HEAD over the 44 tracked `.tf` files in the 13-repo `stack-demo` clone set, and a raw filesystem grep over the 59 `.tf` files in the same working trees — **no terraform anywhere in the clone set nam
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

## 02-024
- **id**: `B02-024`
- **corpus site**: `corpus/services/backend.md:319-319` (bullet)
- **citation**: `docker-compose.yml:183`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
* **Connect-RPC**: `rpc.go` is the top-level wire-up. Look there for the implemented services. **There is no external caller left, and no address to be one with.** `messenger` was the last, and it reached four surfaces — `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR`, `SKILLER_RPC_ADDR` — all reading `http://backend:8083` — **but only two of them because `d11a403` moved them** (`CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`; the other two already held that value at `d11a403^`, and `BACKEND_USERS_RPC_ADDR` never addressed anything but `backend` from `3e85fce` on). *"…pointed at `http://backend:8083` by `d11a403`"* over all four was the false form and is corrected at M257x iter-115; M809 did land, on two variables. All four were set on messenger's own compose block and nowhere else, under compose's own comment *"cms + jobsimulation are folded into app: all four RPC edges are the one backend mux"*. `838d907` deleted that block, so **compose now sets zero `*_RPC_ADDR` variables** and the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is `backend → sentinel` (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48` @ platform `0c91421`). **It is not the only cross-process edge:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The earlier two-of-four split (`http://cms:8091` / `http://jobsimulation:8401`, true at platform `2adcf71`) is history twice over. `app`'s own source comment still says *"additive + DORMANT: external callers (messenger) keep hitting the standalone cms via `CMS_RPC_ADDR` **until the M809 re-point**"* (`app/main.go:1205-1211` @ `b948604` v1.366.0) — **that comment is now stale in `app`**; grade the address against compose, not against the comment. **This bullet used to close by naming a production terraform address. That assertion is DROPPED, not softened.** Measured 2026-08-06 by two mechanisms — `git grep` at each clone's own HEAD over the 44 tracked `.tf` files in the 13-repo `stack-demo` clone set, and a raw filesystem grep over the 59 `.tf` files in the same working trees — **no terraform anywhere in the clone set nam
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

## 02-025
- **id**: `B02-025`
- **corpus site**: `corpus/services/backend.md:319-319` (bullet)
- **citation**: `app/internal/converter/gotenberg.go:31`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/converter/gotenberg.go`  (54 lines)

**CLAIMING UNIT**

```md
* **Connect-RPC**: `rpc.go` is the top-level wire-up. Look there for the implemented services. **There is no external caller left, and no address to be one with.** `messenger` was the last, and it reached four surfaces — `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR`, `SKILLER_RPC_ADDR` — all reading `http://backend:8083` — **but only two of them because `d11a403` moved them** (`CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`; the other two already held that value at `d11a403^`, and `BACKEND_USERS_RPC_ADDR` never addressed anything but `backend` from `3e85fce` on). *"…pointed at `http://backend:8083` by `d11a403`"* over all four was the false form and is corrected at M257x iter-115; M809 did land, on two variables. All four were set on messenger's own compose block and nowhere else, under compose's own comment *"cms + jobsimulation are folded into app: all four RPC edges are the one backend mux"*. `838d907` deleted that block, so **compose now sets zero `*_RPC_ADDR` variables** and the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is `backend → sentinel` (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48` @ platform `0c91421`). **It is not the only cross-process edge:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The earlier two-of-four split (`http://cms:8091` / `http://jobsimulation:8401`, true at platform `2adcf71`) is history twice over. `app`'s own source comment still says *"additive + DORMANT: external callers (messenger) keep hitting the standalone cms via `CMS_RPC_ADDR` **until the M809 re-point**"* (`app/main.go:1205-1211` @ `b948604` v1.366.0) — **that comment is now stale in `app`**; grade the address against compose, not against the comment. **This bullet used to close by naming a production terraform address. That assertion is DROPPED, not softened.** Measured 2026-08-06 by two mechanisms — `git grep` at each clone's own HEAD over the 44 tracked `.tf` files in the 13-repo `stack-demo` clone set, and a raw filesystem grep over the 59 `.tf` files in the same working trees — **no terraform anywhere in the clone set nam
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

## 02-026
- **id**: `B02-026`
- **corpus site**: `corpus/services/backend.md:319-319` (bullet)
- **citation**: `docker-compose.yml:59`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
* **Connect-RPC**: `rpc.go` is the top-level wire-up. Look there for the implemented services. **There is no external caller left, and no address to be one with.** `messenger` was the last, and it reached four surfaces — `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR`, `SKILLER_RPC_ADDR` — all reading `http://backend:8083` — **but only two of them because `d11a403` moved them** (`CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`; the other two already held that value at `d11a403^`, and `BACKEND_USERS_RPC_ADDR` never addressed anything but `backend` from `3e85fce` on). *"…pointed at `http://backend:8083` by `d11a403`"* over all four was the false form and is corrected at M257x iter-115; M809 did land, on two variables. All four were set on messenger's own compose block and nowhere else, under compose's own comment *"cms + jobsimulation are folded into app: all four RPC edges are the one backend mux"*. `838d907` deleted that block, so **compose now sets zero `*_RPC_ADDR` variables** and the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is `backend → sentinel` (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48` @ platform `0c91421`). **It is not the only cross-process edge:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The earlier two-of-four split (`http://cms:8091` / `http://jobsimulation:8401`, true at platform `2adcf71`) is history twice over. `app`'s own source comment still says *"additive + DORMANT: external callers (messenger) keep hitting the standalone cms via `CMS_RPC_ADDR` **until the M809 re-point**"* (`app/main.go:1205-1211` @ `b948604` v1.366.0) — **that comment is now stale in `app`**; grade the address against compose, not against the comment. **This bullet used to close by naming a production terraform address. That assertion is DROPPED, not softened.** Measured 2026-08-06 by two mechanisms — `git grep` at each clone's own HEAD over the 44 tracked `.tf` files in the 13-repo `stack-demo` clone set, and a raw filesystem grep over the 59 `.tf` files in the same working trees — **no terraform anywhere in the clone set nam
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

## 02-027
- **id**: `B02-027`
- **corpus site**: `corpus/services/backend.md:319-319` (bullet)
- **citation**: `app/main.go:1205-1211`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
* **Connect-RPC**: `rpc.go` is the top-level wire-up. Look there for the implemented services. **There is no external caller left, and no address to be one with.** `messenger` was the last, and it reached four surfaces — `BACKEND_USERS_RPC_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR`, `SKILLER_RPC_ADDR` — all reading `http://backend:8083` — **but only two of them because `d11a403` moved them** (`CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`; the other two already held that value at `d11a403^`, and `BACKEND_USERS_RPC_ADDR` never addressed anything but `backend` from `3e85fce` on). *"…pointed at `http://backend:8083` by `d11a403`"* over all four was the false form and is corrected at M257x iter-115; M809 did land, on two variables. All four were set on messenger's own compose block and nowhere else, under compose's own comment *"cms + jobsimulation are folded into app: all four RPC edges are the one backend mux"*. `838d907` deleted that block, so **compose now sets zero `*_RPC_ADDR` variables** and the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is `backend → sentinel` (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48` @ platform `0c91421`). **It is not the only cross-process edge:** `backend` also calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`; `gotenberg` is in the default `core` profile at `docker-compose.yml:183`, consumed at `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL` (`docker-compose.yml:59`). The earlier two-of-four split (`http://cms:8091` / `http://jobsimulation:8401`, true at platform `2adcf71`) is history twice over. `app`'s own source comment still says *"additive + DORMANT: external callers (messenger) keep hitting the standalone cms via `CMS_RPC_ADDR` **until the M809 re-point**"* (`app/main.go:1205-1211` @ `b948604` v1.366.0) — **that comment is now stale in `app`**; grade the address against compose, not against the comment. **This bullet used to close by naming a production terraform address. That assertion is DROPPED, not softened.** Measured 2026-08-06 by two mechanisms — `git grep` at each clone's own HEAD over the 44 tracked `.tf` files in the 13-repo `stack-demo` clone set, and a raw filesystem grep over the 59 `.tf` files in the same working trees — **no terraform anywhere in the clone set nam
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

## 02-028
- **id**: `B02-028`
- **corpus site**: `corpus/services/backend.md:331-331` (bullet)
- **citation**: `internal/authorization/gqlauthz/gqlauthz.go:190-191`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/authorization/gqlauthz/gqlauthz.go`  (263 lines)

**CLAIMING UNIT**

```md
* **Sentinel** — the authorization **engine**, ⚠️ **not "authz on every request"** (which is what this line said until M257x iter-120). At `app` `ad9f3c49` the GraphQL `AuthorizationMiddleware` reaches the resolver **without** calling Sentinel on six paths — no active org (`internal/authorization/gqlauthz/gqlauthz.go:190-191`), no `userId` **variable** in the operation (`:196-197`), self-targeted (`:202-203`), `@resolverAuthorized` (`:209-219`), unparsed op (`:160-161`), public/federation/dev-introspection (`:174-178`) — and the REST groups in `internal/web/backend/backend.go` carry **no BLANKET authz middleware**: 4 of the 6 are `cors` + (`swagger`) + `authn` only, and 2 (`/coursebuilder` `:229-232`, `/credits` `:273-276`) opt into `cbGate`, a Sentinel-backed group middleware (`internal/web/backend/gate.go:27-49`). ⚠️ iter-120's own repair here said *"authn only"* and iter-121 corrected it — the same absolute-quantifier defect, pointing the other way. See [Security & Compliance → Layer 2](../architecture/security_compliance.md#layer-2-authorization)
```

**CITED CONTENT**

```
   187  		}
   188  		l = l.With("viewer", viewer.ID())
   189  		org := viewer.GetOrganization()
   190  		if org == nil || org.ID() == uuid.Nil {
   191  			return next(ctx)
   192  		}
   193  		target, err := targetFromOpCtx(opCtx)
   194  		if err != nil {
```

## 02-029
- **id**: `B02-029`
- **corpus site**: `corpus/services/backend.md:331-331` (bullet)
- **citation**: `internal/web/backend/gate.go:27-49`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/web/backend/gate.go`  (50 lines)

**CLAIMING UNIT**

```md
* **Sentinel** — the authorization **engine**, ⚠️ **not "authz on every request"** (which is what this line said until M257x iter-120). At `app` `ad9f3c49` the GraphQL `AuthorizationMiddleware` reaches the resolver **without** calling Sentinel on six paths — no active org (`internal/authorization/gqlauthz/gqlauthz.go:190-191`), no `userId` **variable** in the operation (`:196-197`), self-targeted (`:202-203`), `@resolverAuthorized` (`:209-219`), unparsed op (`:160-161`), public/federation/dev-introspection (`:174-178`) — and the REST groups in `internal/web/backend/backend.go` carry **no BLANKET authz middleware**: 4 of the 6 are `cors` + (`swagger`) + `authn` only, and 2 (`/coursebuilder` `:229-232`, `/credits` `:273-276`) opt into `cbGate`, a Sentinel-backed group middleware (`internal/web/backend/gate.go:27-49`). ⚠️ iter-120's own repair here said *"authn only"* and iter-121 corrected it — the same absolute-quantifier defect, pointing the other way. See [Security & Compliance → Layer 2](../architecture/security_compliance.md#layer-2-authorization)
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

## 02-030
- **id**: `B02-030`
- **corpus site**: `corpus/services/backend.md:332-332` (bullet)
- **citation**: `app/main.go:524`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
* **Object storage** — **in-process since the v9.0 fold** (2026-08-04), not a service hop: `app` constructs the private and public managers itself (`internalstorage.NewManager` / `NewPublicManager` at `app/main.go:524`, `:525`, re-derived at `app` **`ad9f3c49`** — `origin/main` on 2026-08-06; `main.go` is **byte-identical** to `2035f9a4`, the ref this corpus used to label "origin/main", so every `main.go` line number pinned to `2035f9a4` still resolves) and threads them to each consumer. `STORAGE_RPC_ADDR` has **0 read sites** — its 3 remaining occurrences are comments, one of which (`app/main.go:504`) says *"STORAGE_RPC_ADDR is gone"*. There is no `storage` compose service to address either, since `838d907`. See [Storage](./storage.md)
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

## 02-031
- **id**: `B02-031`
- **corpus site**: `corpus/services/backend.md:332-332` (bullet)
- **citation**: `app/main.go:504`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
* **Object storage** — **in-process since the v9.0 fold** (2026-08-04), not a service hop: `app` constructs the private and public managers itself (`internalstorage.NewManager` / `NewPublicManager` at `app/main.go:524`, `:525`, re-derived at `app` **`ad9f3c49`** — `origin/main` on 2026-08-06; `main.go` is **byte-identical** to `2035f9a4`, the ref this corpus used to label "origin/main", so every `main.go` line number pinned to `2035f9a4` still resolves) and threads them to each consumer. `STORAGE_RPC_ADDR` has **0 read sites** — its 3 remaining occurrences are comments, one of which (`app/main.go:504`) says *"STORAGE_RPC_ADDR is gone"*. There is no `storage` compose service to address either, since `838d907`. See [Storage](./storage.md)
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

## 02-032
- **id**: `B02-032`
- **corpus site**: `corpus/services/backend.md:342-342` (bullet)
- **citation**: `main.go:287`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
* app is **both producer and consumer** of **four** of the five application streams — `backend`, `skillpath`, `jobsimulation`, `cms` — plus the `AI`/`ai_usage` usage stream: **five both-ways streams in all**, and with consumer-only `skiller` on top, **six subscribers against five publishers**. (Watch the partition when quoting a number from here: *four* is the application-stream subtotal, *five* the both-ways total, *six* the subscriber count.) **`skiller` is the exception: `app` only SUBSCRIBES to it and nothing publishes to it.** Enumerated over every publisher constructor in `app` @ `b948604f`, the topics are `backend` (`main.go:287`), `SKILLPATH_STREAM` (`:637`), `CMS_STREAM` (`:1039`), `AI_USAGE_STREAM` and `JOBSIMULATION_STREAM` (`internal/jobsimwiring/wiring.go:127`, `:180`); `SKILLER_STREAM` occurs once in Go **at that ref**, at `main.go:1276`, and it is an `AddSubscriber` call. At `app` `ad9f3c49` the same five publishers sit at `main.go:325`, `:746`, `:1149` and `wiring.go:132`, `:185`, while the subscriber side has been rebuilt as a map (`buildStreamSubscribers`, `subscriber_wiring.go:203-248`, applied by one loop at `main.go:1579-1581`) — so there is no standalone skiller `AddSubscriber` line to cite there. **Grade the shape at the ref you name.** The producer was the standalone skiller service, which is decommissioned — the fact was **deleted, not moved**
```

**CITED CONTENT**

```
   284  	// resolveSubsystemSwitch for why unset is off on a laptop and fatal in production.
   285  	messengerEnabled := mustSubsystemSwitch(envMessengerEnabled)
   286  	customerIOSyncEnabled := mustSubsystemSwitch(envCustomerIOSyncEnabled)
   287  	logger.Info("subsystem switches",
   288  		envMessengerEnabled, messengerEnabled,
   289  		envCustomerIOSyncEnabled, customerIOSyncEnabled)
   290  
```

## 02-033
- **id**: `B02-033`
- **corpus site**: `corpus/services/backend.md:342-342` (bullet)
- **citation**: `internal/jobsimwiring/wiring.go:127`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimwiring/wiring.go`  (283 lines)

**CLAIMING UNIT**

```md
* app is **both producer and consumer** of **four** of the five application streams — `backend`, `skillpath`, `jobsimulation`, `cms` — plus the `AI`/`ai_usage` usage stream: **five both-ways streams in all**, and with consumer-only `skiller` on top, **six subscribers against five publishers**. (Watch the partition when quoting a number from here: *four* is the application-stream subtotal, *five* the both-ways total, *six* the subscriber count.) **`skiller` is the exception: `app` only SUBSCRIBES to it and nothing publishes to it.** Enumerated over every publisher constructor in `app` @ `b948604f`, the topics are `backend` (`main.go:287`), `SKILLPATH_STREAM` (`:637`), `CMS_STREAM` (`:1039`), `AI_USAGE_STREAM` and `JOBSIMULATION_STREAM` (`internal/jobsimwiring/wiring.go:127`, `:180`); `SKILLER_STREAM` occurs once in Go **at that ref**, at `main.go:1276`, and it is an `AddSubscriber` call. At `app` `ad9f3c49` the same five publishers sit at `main.go:325`, `:746`, `:1149` and `wiring.go:132`, `:185`, while the subscriber side has been rebuilt as a map (`buildStreamSubscribers`, `subscriber_wiring.go:203-248`, applied by one loop at `main.go:1579-1581`) — so there is no standalone skiller `AddSubscriber` line to cite there. **Grade the shape at the ref you name.** The producer was the standalone skiller service, which is decommissioned — the fact was **deleted, not moved**
```

**CITED CONTENT**

```
   124  
   125  	// --- Asynq producer client (task-type / queue name strings are frozen — M705 contract).
   126  	workerIndex, _ := strconv.Atoi(getenv("REDIS_WORKER_INDEX"))
   127  	asynqClient := jsworkerclient.NewClient(getenv("REDIS_ADDR"), workerIndex)
   128  
   129  	// --- The AI usage/billing publisher (OD-9): AI usage events go to AI_USAGE_STREAM so app's existing
   130  	// aiUsageManager.AIUsageSubscriber() consumes them — NOT app's own service stream. (M704: preserve
```

## 02-034
- **id**: `B02-034`
- **corpus site**: `corpus/services/backend.md:342-342` (bullet)
- **citation**: `main.go:1276`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
* app is **both producer and consumer** of **four** of the five application streams — `backend`, `skillpath`, `jobsimulation`, `cms` — plus the `AI`/`ai_usage` usage stream: **five both-ways streams in all**, and with consumer-only `skiller` on top, **six subscribers against five publishers**. (Watch the partition when quoting a number from here: *four* is the application-stream subtotal, *five* the both-ways total, *six* the subscriber count.) **`skiller` is the exception: `app` only SUBSCRIBES to it and nothing publishes to it.** Enumerated over every publisher constructor in `app` @ `b948604f`, the topics are `backend` (`main.go:287`), `SKILLPATH_STREAM` (`:637`), `CMS_STREAM` (`:1039`), `AI_USAGE_STREAM` and `JOBSIMULATION_STREAM` (`internal/jobsimwiring/wiring.go:127`, `:180`); `SKILLER_STREAM` occurs once in Go **at that ref**, at `main.go:1276`, and it is an `AddSubscriber` call. At `app` `ad9f3c49` the same five publishers sit at `main.go:325`, `:746`, `:1149` and `wiring.go:132`, `:185`, while the subscriber side has been rebuilt as a map (`buildStreamSubscribers`, `subscriber_wiring.go:203-248`, applied by one loop at `main.go:1579-1581`) — so there is no standalone skiller `AddSubscriber` line to cite there. **Grade the shape at the ref you name.** The producer was the standalone skiller service, which is decommissioned — the fact was **deleted, not moved**
```

**CITED CONTENT**

```
  1273  		cmsWebhookHandler,
  1274  		academyContentManager,
  1275  		academyEmbeddingsManager,
  1276  		apiKeyManager,
  1277  		invitationManager,
  1278  		labsDeps,
  1279  		courseBuilderDeps,
```

## 02-035
- **id**: `B02-035`
- **corpus site**: `corpus/services/backend.md:342-342` (bullet)
- **citation**: `main.go:325`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
* app is **both producer and consumer** of **four** of the five application streams — `backend`, `skillpath`, `jobsimulation`, `cms` — plus the `AI`/`ai_usage` usage stream: **five both-ways streams in all**, and with consumer-only `skiller` on top, **six subscribers against five publishers**. (Watch the partition when quoting a number from here: *four* is the application-stream subtotal, *five* the both-ways total, *six* the subscriber count.) **`skiller` is the exception: `app` only SUBSCRIBES to it and nothing publishes to it.** Enumerated over every publisher constructor in `app` @ `b948604f`, the topics are `backend` (`main.go:287`), `SKILLPATH_STREAM` (`:637`), `CMS_STREAM` (`:1039`), `AI_USAGE_STREAM` and `JOBSIMULATION_STREAM` (`internal/jobsimwiring/wiring.go:127`, `:180`); `SKILLER_STREAM` occurs once in Go **at that ref**, at `main.go:1276`, and it is an `AddSubscriber` call. At `app` `ad9f3c49` the same five publishers sit at `main.go:325`, `:746`, `:1149` and `wiring.go:132`, `:185`, while the subscriber side has been rebuilt as a map (`buildStreamSubscribers`, `subscriber_wiring.go:203-248`, applied by one loop at `main.go:1579-1581`) — so there is no standalone skiller `AddSubscriber` line to cite there. **Grade the shape at the ref you name.** The producer was the standalone skiller service, which is decommissioned — the fact was **deleted, not moved**
```

**CITED CONTENT**

```
   322  		logger.Error("can't connect to redis", "error", err)
   323  		return
   324  	}
   325  	pub, err := pubsub.NewPublisher(serviceName, redisClientStream)
   326  	if err != nil {
   327  		logger.Error("can't init event publisher", "error", err)
   328  		return
```

## 02-036
- **id**: `B02-036`
- **corpus site**: `corpus/services/backend.md:342-342` (bullet)
- **citation**: `subscriber_wiring.go:203-248`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/subscriber_wiring.go`  (282 lines)

**CLAIMING UNIT**

```md
* app is **both producer and consumer** of **four** of the five application streams — `backend`, `skillpath`, `jobsimulation`, `cms` — plus the `AI`/`ai_usage` usage stream: **five both-ways streams in all**, and with consumer-only `skiller` on top, **six subscribers against five publishers**. (Watch the partition when quoting a number from here: *four* is the application-stream subtotal, *five* the both-ways total, *six* the subscriber count.) **`skiller` is the exception: `app` only SUBSCRIBES to it and nothing publishes to it.** Enumerated over every publisher constructor in `app` @ `b948604f`, the topics are `backend` (`main.go:287`), `SKILLPATH_STREAM` (`:637`), `CMS_STREAM` (`:1039`), `AI_USAGE_STREAM` and `JOBSIMULATION_STREAM` (`internal/jobsimwiring/wiring.go:127`, `:180`); `SKILLER_STREAM` occurs once in Go **at that ref**, at `main.go:1276`, and it is an `AddSubscriber` call. At `app` `ad9f3c49` the same five publishers sit at `main.go:325`, `:746`, `:1149` and `wiring.go:132`, `:185`, while the subscriber side has been rebuilt as a map (`buildStreamSubscribers`, `subscriber_wiring.go:203-248`, applied by one loop at `main.go:1579-1581`) — so there is no standalone skiller `AddSubscriber` line to cite there. **Grade the shape at the ref you name.** The producer was the standalone skiller service, which is decommissioned — the fact was **deleted, not moved**
```

**CITED CONTENT**

```
   200  	return nil
   201  }
   202  
   203  func buildStreamSubscribers(d streamSubscriberDeps) map[string]*pubsub.Subscriber {
   204  	subs := make(map[string]*pubsub.Subscriber, 6)
   205  
   206  	// SkillPath
   207  	subs[d.Streams.SkillPath] = d.SkillPath.SkillPathSubscriber()
   208  	// Skiller
   209  	subs[d.Streams.Skiller] = d.Skiller.SkillerSubscriber()
   210  
   211  	// JobSimulation stream — consumed by the jobsimulation domain + skillpath + the ported
   212  	// jobsim engine (streamrouter composes those three onto one subscriber),.
   213  	jobsimSub := d.JobSimulation.JobSimulationSubscriber()
   214  	subs[d.Streams.JobSimulation] = jobsimSub
   215  
   216  	// CMS — app's JSManager handlers + the ported jobsim engine's cms handlers + (when the
   217  	// Directus edge is configured) the folded cms similarity + studio handlers, on ONE
   218  	// subscriber.
   219  	//
   220  	// cms-in-app M806: the cms similarity re-index and app's jobsim session cleanup fire on the
   221  	// same CmsJobSimulationDeleted/Updated events but act on disjoint rows, so they compose
   222  	// (order-independent, each self-selects by proto payload type).
   223  	cmsSub := d.JobSimsCMS.CMSSubscriber()
   224  	cmsSub.AddHandler(d.JobsimEngine.CmsStreamHandlers()...)
   225  	if d.CMSSimilarity != nil && d.CMSStudio != nil {
   226  		cmsSub.AddHandler(
   227  			pubsub.EventHandler(d.CMSSimilarity.CmsJobSimulationDeletedHandler),
   228  			pubsub.EventHandler(d.CMSSimilarity.CmsJobSimulationUpdatedHandler),
   229  			pubsub.EventHandler(d.CMSStudio.CmsJobSimulationTranslationRequestedHandler),
   230  			pubsub.EventHandler(d.CMSStudio.CmsJobSimulationCloneRequestedHandler),
   231  		)
   232  	}
   233  	subs[d.Streams.CMS] = cmsSub
   234  
   235  	// AI Usage
   236  	subs[d.Streams.AIUsage] = d.AIUsage.AIUsageSubscriber()
   237  
   238  	// AI Readiness audience-match — subscribes to backend's OWN event stream (SERVICE_NAME) so
   239  	// the tag-assigned events published from within the org manager fan out to per-user
   240  	// invitations in the same process. Same-stream self-subscribe is intentional: it keeps the
   241  	// retry / dead-letter semantics consistent with cross-service events.
   242  	backendSelfSub := d.AIReadiness.Subscriber()
   243  	// cms-in-app M806: cms's studio OrganizationMemberDeleted handler, on the
```

## 02-037
- **id**: `B02-037`
- **corpus site**: `corpus/services/backend.md:342-342` (bullet)
- **citation**: `main.go:1579-1581`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
* app is **both producer and consumer** of **four** of the five application streams — `backend`, `skillpath`, `jobsimulation`, `cms` — plus the `AI`/`ai_usage` usage stream: **five both-ways streams in all**, and with consumer-only `skiller` on top, **six subscribers against five publishers**. (Watch the partition when quoting a number from here: *four* is the application-stream subtotal, *five* the both-ways total, *six* the subscriber count.) **`skiller` is the exception: `app` only SUBSCRIBES to it and nothing publishes to it.** Enumerated over every publisher constructor in `app` @ `b948604f`, the topics are `backend` (`main.go:287`), `SKILLPATH_STREAM` (`:637`), `CMS_STREAM` (`:1039`), `AI_USAGE_STREAM` and `JOBSIMULATION_STREAM` (`internal/jobsimwiring/wiring.go:127`, `:180`); `SKILLER_STREAM` occurs once in Go **at that ref**, at `main.go:1276`, and it is an `AddSubscriber` call. At `app` `ad9f3c49` the same five publishers sit at `main.go:325`, `:746`, `:1149` and `wiring.go:132`, `:185`, while the subscriber side has been rebuilt as a map (`buildStreamSubscribers`, `subscriber_wiring.go:203-248`, applied by one loop at `main.go:1579-1581`) — so there is no standalone skiller `AddSubscriber` line to cite there. **Grade the shape at the ref you name.** The producer was the standalone skiller service, which is decommissioned — the fact was **deleted, not moved**
```

**CITED CONTENT**

```
  1576  	// map makes that unrepresentable: one subscriber per stream, by construction. Order is
  1577  	// irrelevant (the destination is a map, and colony's own init() already iterates it in
  1578  	// randomized order) — proved by TestRegistrationOrderDoesNotMatter.
  1579  	for stream, sub := range streamSubscribers {
  1580  		subServer.AddSubscriber(stream, sub)
  1581  	}
  1582  
  1583  	wg.Go(func() {
  1584  		defer cancelServerContext()
```

## 02-038
- **id**: `B02-038`
- **corpus site**: `corpus/services/chronos.md:3-11` (paragraph)
- **citation**: `app/internal/jobsimwiring/worker.go:24-35`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimwiring/worker.go`  (88 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Decommissioned — no longer in local orchestration (but the GitHub repo is NOT archived)
>
> Chronos was removed from `platform/docker-compose.yml` and `platform/repos.yml` in mid-2026:
> - Platform: commit `045857c` — "remove chronos service from orchestration"
> - Jobsimulation: PR `#395` (`feat/remove-chronos-and-realtime`), commit `09631fb2` — "remove Chronos references and update documentation to reflect Asynq integration for session timeout management"
>
> The use cases Chronos covered (session timeouts, delayed events from Jobsimulation) have moved to **in-process [Asynq](https://github.com/hibiken/asynq)** running inside the jobsimulation engine — which has itself since been folded into `app`, so those Asynq workers now run in the **`backend`** process (`app/internal/jobsimwiring/worker.go:24-35`). The Chronos GitHub repository still exists but is no longer cloned by `make init` and no service in the current compose file depends on it.
>
> The detail below is preserved for historical context and in case a future need for a generic timer service resurfaces.
```

**CITED CONTENT**

```
    21  
    22  	jsqueues "github.com/anthropos-work/app/internal/jobsimulation/worker/queues"
    23  	jstasks "github.com/anthropos-work/app/internal/jobsimulation/worker/tasks"
    24  	"github.com/hibiken/asynq"
    25  )
    26  
    27  func StartWorkers(ctx context.Context, logger *slog.Logger, dj *Runtime) {
    28  	workerIndex, _ := strconv.Atoi(os.Getenv("REDIS_WORKER_INDEX"))
    29  	redisOpt := asynq.RedisClientOpt{Addr: os.Getenv("REDIS_ADDR"), DB: workerIndex}
    30  	errH := asynq.ErrorHandlerFunc(func(_ context.Context, task *asynq.Task, err error) {
    31  		logger.Error("jobsim asynq task error", "task", task.Type(), "error", err)
    32  	})
    33  
    34  	// --- Standard pool (heavy/validation/media/anticheat/analytics).
    35  	std := asynq.NewServer(redisOpt, asynq.Config{
    36  		Concurrency:  10,
    37  		Queues:       map[string]int{jsqueues.HighPriorityQueue: 7, jsqueues.TimersQueue: 5, jsqueues.CodeRunQueue: 3, jsqueues.DefaultQueue: 2, jsqueues.LowPriorityQueue: 1},
    38  		LogLevel:     asynq.InfoLevel,
```

## 02-039
- **id**: `B02-039`
- **corpus site**: `corpus/services/clerk-integration.md:65-65` (table-row)
- **citation**: `app/internal/admin/impersonation/manager.go:101`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/admin/impersonation/manager.go`  (218 lines)

**CLAIMING UNIT**

```md
| 1 | `app/internal/admin/impersonation/manager.go:101` (`m.signInTokenCl.Create`) | `app` `ad9f3c49` | **The product feature** — app-native admin impersonation, the one this page always documented | A Sentinel permission check + a non-nil actor, then an audit row on **every** attempt, success or failure (`manager.go:1-4`). **The permission it checks is `ActionObjectTaxonomy` / `UserActionWrite`** (`internal/web/backend/graphql/graph/resolver_admin_audit.go:19-23`) — a **taxonomy-write** permission, not a dedicated impersonation one |
```

**CITED CONTENT**

```
    98  	targetUserID, targetEmail := m.resolveInternalUser(ctx, clerkU.ID, email)
    99  
   100  	// 3) Create the sign-in token.
   101  	tok, err := m.signInTokenCl.Create(ctx, &clerkSignInToken.CreateParams{
   102  		UserID:           &clerkU.ID,
   103  		ExpiresInSeconds: &expiresIn,
   104  	})
```

## 02-040
- **id**: `B02-040`
- **corpus site**: `corpus/services/clerk-integration.md:65-65` (table-row)
- **citation**: `internal/web/backend/graphql/graph/resolver_admin_audit.go:19-23`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/web/backend/graphql/graph/resolver_admin_audit.go`  (72 lines)

**CLAIMING UNIT**

```md
| 1 | `app/internal/admin/impersonation/manager.go:101` (`m.signInTokenCl.Create`) | `app` `ad9f3c49` | **The product feature** — app-native admin impersonation, the one this page always documented | A Sentinel permission check + a non-nil actor, then an audit row on **every** attempt, success or failure (`manager.go:1-4`). **The permission it checks is `ActionObjectTaxonomy` / `UserActionWrite`** (`internal/web/backend/graphql/graph/resolver_admin_audit.go:19-23`) — a **taxonomy-write** permission, not a dedicated impersonation one |
```

**CITED CONTENT**

```
    16  )
    17  
    18  // ImpersonateUser is the resolver for the impersonateUser field.
    19  func (r *mutationResolver) ImpersonateUser(ctx context.Context, email string, expiresInSeconds *int) (*impersonation.ImpersonateResult, error) {
    20  	if err := r.authorizationManager.UserCheckActionPermission(
    21  		ctx, permission.ActionObjectTaxonomy, permission.UserActionWrite,
    22  	); err != nil {
    23  		return nil, fmt.Errorf("forbidden")
    24  	}
    25  	actor := authn.UserFromContext(ctx)
    26  	if actor == nil {
```

## 02-041
- **id**: `B02-041`
- **corpus site**: `corpus/services/clerk-integration.md:66-66` (table-row)
- **citation**: `next-web-app/apps/web/src/app/api/dev/login-as/route.ts:79`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/apps/web/src/app/api/dev/login-as/route.ts`  (100 lines)

**CLAIMING UNIT**

```md
| 2 | `next-web-app/apps/web/src/app/api/dev/login-as/route.ts:79` | `next-web-app` `8297c684` | Dev "log in as a real Clerk user" harness → `/dev/accept` | `DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production'` (`apps/web/src/lib/devLogin.ts:28`); hard-404 otherwise |
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

## 02-042
- **id**: `B02-042`
- **corpus site**: `corpus/services/clerk-integration.md:66-66` (table-row)
- **citation**: `apps/web/src/lib/devLogin.ts:28`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/apps/web/src/lib/devLogin.ts`  (35 lines)

**CLAIMING UNIT**

```md
| 2 | `next-web-app/apps/web/src/app/api/dev/login-as/route.ts:79` | `next-web-app` `8297c684` | Dev "log in as a real Clerk user" harness → `/dev/accept` | `DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production'` (`apps/web/src/lib/devLogin.ts:28`); hard-404 otherwise |
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

## 02-043
- **id**: `B02-043`
- **corpus site**: `corpus/services/clerk-integration.md:67-67` (table-row)
- **citation**: `next-web-app/e2e/auth.setup.ts:72`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/e2e/auth.setup.ts`  (87 lines)

**CLAIMING UNIT**

```md
| 3 | `next-web-app/e2e/auth.setup.ts:72` | `next-web-app` `8297c684` | **Playwright e2e auth setup** — mints a ticket rather than driving the password form | **No `NODE_ENV` gate** — it is a test-runner file, never in an app build, but it runs against a **real Clerk instance** with a real `CLERK_SECRET_KEY` |
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

## 02-044
- **id**: `B02-044`
- **corpus site**: `corpus/services/clerk-integration.md:68-68` (table-row)
- **citation**: `studio-desk/src/routes/dev.ts:83`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/studio-desk/src/routes/dev.ts`  (107 lines)

**CLAIMING UNIT**

```md
| 4 | `studio-desk/src/routes/dev.ts:83` | `studio-desk` `41ee357` | Dev login harness → `dev-accept.html` | `DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production'` (`src/lib/devLogin.ts:33`) |
```

**CITED CONTENT**

```
    80        }
    81  
    82        // Mint a one-time sign-in token (10-minute validity is plenty for a redirect).
    83        const { token } = await clerkClient.signInTokens.createSignInToken({
    84          userId: user.id,
    85          expiresInSeconds: 600,
    86        });
```
