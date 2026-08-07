# TIER-1 ADJUDICATION BATCH 08 — 44 pairs

Each item: the CLAIMING UNIT (corpus prose) and the CITED CONTENT it points at (source lines,
line-numbered, +-3 lines of context). Answer ONE question per item.

## 08-001
- **id**: `B08-001`
- **corpus site**: `corpus/architecture/architecture_overview.md:265-284` (bullet)
- **citation**: `docker-compose.yml:57`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
*   **Synchronous**: Connect-RPC/HTTP endpoints — down to **one Connect-RPC edge on a local stack,
    `backend → sentinel`**. At platform `0c91421` that is the only cross-process **Connect-RPC** address,
    `AUTHORIZATION_ADDRESS=http://sentinel:8087`
    (`docker-compose.yml:48`), and there are **zero `*_RPC_ADDR` variables**. **It is NOT the only service
    address compose sets, and not the only cross-process edge** — this passage previously said *"compose
    sets exactly one service address"*, which **is false** and is retracted (corrected M257x iter-102).
    The same `backend` block also sets `GOTENBERG_URL=http://gotenberg:3200` (`docker-compose.yml:57` — a
    second container on the **default** `core` profile at `:183`, reached over **plain HTTP**, not
    Connect-RPC, at `app/internal/converter/gotenberg.go:31` @ `app` `ad9f3c49`), `JUDGE0_BASE_URL`
    (`docker-compose.yml:59`) and `REDIS_ADDR` (`docker-compose.yml:66`). **The correctly-scoped form is
    this document's own local-stack diagram below** — *"the only cross-process **RPC** edge out of backend
    on a core stack"* — which was right while this line was wrong, 55 lines apart in one file.
    On the `*_RPC_ADDR` half: the `messenger` block was the last thing
    that set any (`BACKEND_USERS_`, `CMS_`, `JOBSIMULATION_`, `SKILLER_` — **all four read
    `http://backend:8083`, but `d11a403` moved only the MIDDLE TWO**: `CMS_RPC_ADDR` and
    `JOBSIMULATION_RPC_ADDR`. `BACKEND_USERS_RPC_ADDR` and `SKILLER_RPC_ADDR` already held that value at
    `d11a403^`, and `BACKEND_USERS_RPC_ADDR` never addressed anything but `backend` from its introduction
    at `3e85fce` — it only ever moved ports, so there was nothing to re-point. Corrected M257x iter-115),
    and `838d907` deleted that service. The env-var *names* still exist
    in consumer code; no local compose file configures them
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

## 08-002
- **id**: `B08-002`
- **corpus site**: `corpus/architecture/architecture_overview.md:265-284` (bullet)
- **citation**: `app/internal/converter/gotenberg.go:31`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/converter/gotenberg.go`  (54 lines)

**CLAIMING UNIT**

```md
*   **Synchronous**: Connect-RPC/HTTP endpoints — down to **one Connect-RPC edge on a local stack,
    `backend → sentinel`**. At platform `0c91421` that is the only cross-process **Connect-RPC** address,
    `AUTHORIZATION_ADDRESS=http://sentinel:8087`
    (`docker-compose.yml:48`), and there are **zero `*_RPC_ADDR` variables**. **It is NOT the only service
    address compose sets, and not the only cross-process edge** — this passage previously said *"compose
    sets exactly one service address"*, which **is false** and is retracted (corrected M257x iter-102).
    The same `backend` block also sets `GOTENBERG_URL=http://gotenberg:3200` (`docker-compose.yml:57` — a
    second container on the **default** `core` profile at `:183`, reached over **plain HTTP**, not
    Connect-RPC, at `app/internal/converter/gotenberg.go:31` @ `app` `ad9f3c49`), `JUDGE0_BASE_URL`
    (`docker-compose.yml:59`) and `REDIS_ADDR` (`docker-compose.yml:66`). **The correctly-scoped form is
    this document's own local-stack diagram below** — *"the only cross-process **RPC** edge out of backend
    on a core stack"* — which was right while this line was wrong, 55 lines apart in one file.
    On the `*_RPC_ADDR` half: the `messenger` block was the last thing
    that set any (`BACKEND_USERS_`, `CMS_`, `JOBSIMULATION_`, `SKILLER_` — **all four read
    `http://backend:8083`, but `d11a403` moved only the MIDDLE TWO**: `CMS_RPC_ADDR` and
    `JOBSIMULATION_RPC_ADDR`. `BACKEND_USERS_RPC_ADDR` and `SKILLER_RPC_ADDR` already held that value at
    `d11a403^`, and `BACKEND_USERS_RPC_ADDR` never addressed anything but `backend` from its introduction
    at `3e85fce` — it only ever moved ports, so there was nothing to re-point. Corrected M257x iter-115),
    and `838d907` deleted that service. The env-var *names* still exist
    in consumer code; no local compose file configures them
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

## 08-003
- **id**: `B08-003`
- **corpus site**: `corpus/architecture/architecture_overview.md:265-284` (bullet)
- **citation**: `docker-compose.yml:59`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
*   **Synchronous**: Connect-RPC/HTTP endpoints — down to **one Connect-RPC edge on a local stack,
    `backend → sentinel`**. At platform `0c91421` that is the only cross-process **Connect-RPC** address,
    `AUTHORIZATION_ADDRESS=http://sentinel:8087`
    (`docker-compose.yml:48`), and there are **zero `*_RPC_ADDR` variables**. **It is NOT the only service
    address compose sets, and not the only cross-process edge** — this passage previously said *"compose
    sets exactly one service address"*, which **is false** and is retracted (corrected M257x iter-102).
    The same `backend` block also sets `GOTENBERG_URL=http://gotenberg:3200` (`docker-compose.yml:57` — a
    second container on the **default** `core` profile at `:183`, reached over **plain HTTP**, not
    Connect-RPC, at `app/internal/converter/gotenberg.go:31` @ `app` `ad9f3c49`), `JUDGE0_BASE_URL`
    (`docker-compose.yml:59`) and `REDIS_ADDR` (`docker-compose.yml:66`). **The correctly-scoped form is
    this document's own local-stack diagram below** — *"the only cross-process **RPC** edge out of backend
    on a core stack"* — which was right while this line was wrong, 55 lines apart in one file.
    On the `*_RPC_ADDR` half: the `messenger` block was the last thing
    that set any (`BACKEND_USERS_`, `CMS_`, `JOBSIMULATION_`, `SKILLER_` — **all four read
    `http://backend:8083`, but `d11a403` moved only the MIDDLE TWO**: `CMS_RPC_ADDR` and
    `JOBSIMULATION_RPC_ADDR`. `BACKEND_USERS_RPC_ADDR` and `SKILLER_RPC_ADDR` already held that value at
    `d11a403^`, and `BACKEND_USERS_RPC_ADDR` never addressed anything but `backend` from its introduction
    at `3e85fce` — it only ever moved ports, so there was nothing to re-point. Corrected M257x iter-115),
    and `838d907` deleted that service. The env-var *names* still exist
    in consumer code; no local compose file configures them
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

## 08-004
- **id**: `B08-004`
- **corpus site**: `corpus/architecture/architecture_overview.md:265-284` (bullet)
- **citation**: `docker-compose.yml:66`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
*   **Synchronous**: Connect-RPC/HTTP endpoints — down to **one Connect-RPC edge on a local stack,
    `backend → sentinel`**. At platform `0c91421` that is the only cross-process **Connect-RPC** address,
    `AUTHORIZATION_ADDRESS=http://sentinel:8087`
    (`docker-compose.yml:48`), and there are **zero `*_RPC_ADDR` variables**. **It is NOT the only service
    address compose sets, and not the only cross-process edge** — this passage previously said *"compose
    sets exactly one service address"*, which **is false** and is retracted (corrected M257x iter-102).
    The same `backend` block also sets `GOTENBERG_URL=http://gotenberg:3200` (`docker-compose.yml:57` — a
    second container on the **default** `core` profile at `:183`, reached over **plain HTTP**, not
    Connect-RPC, at `app/internal/converter/gotenberg.go:31` @ `app` `ad9f3c49`), `JUDGE0_BASE_URL`
    (`docker-compose.yml:59`) and `REDIS_ADDR` (`docker-compose.yml:66`). **The correctly-scoped form is
    this document's own local-stack diagram below** — *"the only cross-process **RPC** edge out of backend
    on a core stack"* — which was right while this line was wrong, 55 lines apart in one file.
    On the `*_RPC_ADDR` half: the `messenger` block was the last thing
    that set any (`BACKEND_USERS_`, `CMS_`, `JOBSIMULATION_`, `SKILLER_` — **all four read
    `http://backend:8083`, but `d11a403` moved only the MIDDLE TWO**: `CMS_RPC_ADDR` and
    `JOBSIMULATION_RPC_ADDR`. `BACKEND_USERS_RPC_ADDR` and `SKILLER_RPC_ADDR` already held that value at
    `d11a403^`, and `BACKEND_USERS_RPC_ADDR` never addressed anything but `backend` from its introduction
    at `3e85fce` — it only ever moved ports, so there was nothing to re-point. Corrected M257x iter-115),
    and `838d907` deleted that service. The env-var *names* still exist
    in consumer code; no local compose file configures them
```

**CITED CONTENT**

```
    63        - MEDIA_URL=https://media.anthropos.work
    64        - META_PORT=8084
    65        - PORT=8082
    66        - REDIS_ADDR=redis:6379
    67        - REDIS_STREAMS_INDEX=4
    68        - REDIS_WORKER_INDEX=0
    69        - RPC_PORT=8083
```

## 08-005
- **id**: `B08-005`
- **corpus site**: `corpus/architecture/architecture_overview.md:295-318` (bullet)
- **citation**: `internal/cms/studio/markdownManager.go:29-30`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/studio/markdownManager.go`  (144 lines)

**CLAIMING UNIT**

```md
*   **AI Providers**: the default clients are EU-resident, and **there is no ordered EU-first fallback
    ladder** — the chain *"Azure OpenAI EU → Azure OpenAI US → direct OpenAI"* was retracted at
    [`external_services.md:579`](./external_services.md) and is corrected here (M257x iter-46). Inside the
    AI manager there are two US paths — a **feature flag** and a **429 retry target**, not fallback rungs —
    but **that is not the whole set**: [`external_services.md:602-607`](./external_services.md) enumerates
    **four live** ways a request leaves the EU, of which the two outside the manager are `ANTHROPIC_API_KEY`
    and **an authored sequence with `ai_vendor` unset** — the latter reaching direct US OpenAI
    *unconditionally, on the first attempt, with no flag and no 429*. A fifth arm, **Studio-Room's own
    `openai` `TARGET SERVICE`**, exists in code but is **selected by no shipped config** (all three
    `app/studio/configs/*.ini` pin `azure`). Scope corrected M257x iter-48, count corrected to five at
    iter-49 and to four-live-plus-one-latent at iter-52. Measured at
    `app/internal/jobsimulation/ai/ai.go`: `getClient` defaults to `azureClientEu` and swaps to
    `azureClientUs` when the PostHog flag **`flag_use_azure_us`** is on (`:262-276`); direct OpenAI is the
    retry target on HTTP 429 (`isThrottlingError` at `:129`, applied at `:166` and `:325`). **⚠️ "EU-first"
    is not "EU-only" — a feature flag routes traffic to the US, and an unset `ai_vendor` does so with no flag
    at all.** AWS Bedrock is a *per-call vendor*
    (`AnthropicAws`, pinned to `eu-west-1` at `:85-88`), never a fallback tier. **Mistral is not part of this
    routing chain** — but it *is* live in `app`: `internal/cms/studio/markdownManager.go:29-30` builds a
    Mistral OCR client (`mistralocr.New(aiKey)`) for **Studio document OCR**. **The key reaches it from the
    CALLER, not from the environment**: `studioManager.go:583` reads `os.Getenv("MISTRAL_API_KEY")` and
    passes it in. This sentence used to cite `:11,19` and say the constructor read the variable itself —
    `:11` is the closing paren of the import block and `:19` is a doc comment recording that that very
    `os.Getenv` read was **removed**, so the anchor named the fix as if it were the defect.
    It is a separate, single-purpose provider, not a tier in the simulation cascade
```

**CITED CONTENT**

```
    26  // The error return is kept — it is part of the call site's shape and the OCR client
    27  // may grow a failing constructor again — but there is nothing left in here that can
    28  // fail.
    29  func NewMarkdownManager(aiKey string) (*MarkdownManager, error) {
    30  	return &MarkdownManager{ocr: mistralocr.New(aiKey)}, nil
    31  }
    32  
    33  func (m *MarkdownManager) OCRProcess(ctx context.Context, documentData []byte) (*string, int, error) {
```

## 08-006
- **id**: `B08-006`
- **corpus site**: `corpus/architecture/architecture_overview.md:326-326` (paragraph)
- **citation**: `graphql-wundergraph/terraform/main.tf:20`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/graphql-wundergraph/terraform/main.tf`  (63 lines)

**CLAIMING UNIT**

```md
**In production** (the router still exists there — `graphql-wundergraph/terraform/main.tf:20` `= 1`):
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

## 08-007
- **id**: `B08-007`
- **corpus site**: `corpus/architecture/architecture_overview.md:346-373` (paragraph)
- **citation**: `main.go:451`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> `roadrunner` is **not** a gRPC hop from `backend` in either column — it was folded in with jobsim-in-app and
> `backend` calls Judge0 directly. **Nor is `storage` one in EITHER column** — the production diagram above no
longer lists it, because the edge is dead there too: `storage`'s ECS service block is **gone from terraform
entirely**, and says so in-comment — at `9f8cb53` `storage/terraform/main.tf` is 18 lines (`:9-11`
*"The ECS service that used to live here is GONE (v9.0 'support-in-app')"*), the module kept only so the
buckets, CloudFront distribution and media DNS record keep their `prevent_destroy` guards (`:13-16`). It
read `service_desired_count = 0` at the intermediate `63bffc8`. And `STORAGE_RPC_ADDR` has **zero** reads
anywhere in `app`
(**3 hits in Go source** at `9d00a313`, every one a comment — `main.go:451`,
`internal/jobsimwiring/wiring.go:101`, `internal/storagens/callsites_test.go:189`. **Not 3 repo-wide**:
`git -C stack-demo/app grep -n STORAGE_RPC_ADDR 9d00a313` returns **29 lines across 18 files**, the rest
being CHANGELOG and `knowledge/*.md`. The Go scope is what carries the claim; the repo-wide form was a
mis-transcription of it). The earlier wording scoped this retraction to
*"locally"*, which left the prod edge affirmatively standing (corrected M257x iter-85). Platform
**`0dab54d`** ("storage-in-app,
> v9.0") deleted `STORAGE_RPC_ADDR` from `backend`'s env, dropped `storage` from `backend`'s `depends_on`
> — the replacement comment read *"storage removed at v9.0: served in-process by this container now"*
> — and moved the service to `profiles: [storage-legacy]`; **`838d907`
> (merged `0c91421`, 2026-08-05) then deleted the `storage` service and that profile outright**, so there
> is nothing left to opt into. The app side
> closed at **`app` `9d00a313`** (v1.367.0) and is still closed at `app` **origin/main**, where
> `STORAGE_RPC_ADDR` has **zero reads** and `main.go:504`
> says *"the standalone service takes no traffic and STORAGE_RPC_ADDR is gone."* That comment stood at
> `main.go:451` at `9d00a313`; **the line number moved without the code moving**, which is what a week of
> `app` commits costs an anchor (re-derived M257x iter-87). At the older `b948604`
> v1.366.0 the variable is still read — `internal/jobsimwiring/wiring.go:115` — so **state the ref you mean**. See
> [`roadrunner.md`](../services/roadrunner.md), [`storage
```

**CITED CONTENT**

```
   448  	// event via the notifier — plan A) and the M407/M408 in-product preview +
   449  	// review-and-edit surface, so a saved override is reflected byte-for-byte.
   450  	aiReadinessOverrides := emailoverride.NewManager(ent)
   451  	aiReadinessNotifier := notifications.NewManager(notifications.Options{
   452  		Logger:       logger,
   453  		Ent:          ent,
   454  		Publisher:    pub,
```

## 08-008
- **id**: `B08-008`
- **corpus site**: `corpus/architecture/architecture_overview.md:346-373` (paragraph)
- **citation**: `internal/jobsimwiring/wiring.go:101`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimwiring/wiring.go`  (283 lines)

**CLAIMING UNIT**

```md
> `roadrunner` is **not** a gRPC hop from `backend` in either column — it was folded in with jobsim-in-app and
> `backend` calls Judge0 directly. **Nor is `storage` one in EITHER column** — the production diagram above no
longer lists it, because the edge is dead there too: `storage`'s ECS service block is **gone from terraform
entirely**, and says so in-comment — at `9f8cb53` `storage/terraform/main.tf` is 18 lines (`:9-11`
*"The ECS service that used to live here is GONE (v9.0 'support-in-app')"*), the module kept only so the
buckets, CloudFront distribution and media DNS record keep their `prevent_destroy` guards (`:13-16`). It
read `service_desired_count = 0` at the intermediate `63bffc8`. And `STORAGE_RPC_ADDR` has **zero** reads
anywhere in `app`
(**3 hits in Go source** at `9d00a313`, every one a comment — `main.go:451`,
`internal/jobsimwiring/wiring.go:101`, `internal/storagens/callsites_test.go:189`. **Not 3 repo-wide**:
`git -C stack-demo/app grep -n STORAGE_RPC_ADDR 9d00a313` returns **29 lines across 18 files**, the rest
being CHANGELOG and `knowledge/*.md`. The Go scope is what carries the claim; the repo-wide form was a
mis-transcription of it). The earlier wording scoped this retraction to
*"locally"*, which left the prod edge affirmatively standing (corrected M257x iter-85). Platform
**`0dab54d`** ("storage-in-app,
> v9.0") deleted `STORAGE_RPC_ADDR` from `backend`'s env, dropped `storage` from `backend`'s `depends_on`
> — the replacement comment read *"storage removed at v9.0: served in-process by this container now"*
> — and moved the service to `profiles: [storage-legacy]`; **`838d907`
> (merged `0c91421`, 2026-08-05) then deleted the `storage` service and that profile outright**, so there
> is nothing left to opt into. The app side
> closed at **`app` `9d00a313`** (v1.367.0) and is still closed at `app` **origin/main**, where
> `STORAGE_RPC_ADDR` has **zero reads** and `main.go:504`
> says *"the standalone service takes no traffic and STORAGE_RPC_ADDR is gone."* That comment stood at
> `main.go:451` at `9d00a313`; **the line number moved without the code moving**, which is what a week of
> `app` commits costs an anchor (re-derived M257x iter-87). At the older `b948604`
> v1.366.0 the variable is still read — `internal/jobsimwiring/wiring.go:115` — so **state the ref you mean**. See
> [`roadrunner.md`](../services/roadrunner.md), [`storage
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

## 08-009
- **id**: `B08-009`
- **corpus site**: `corpus/architecture/architecture_overview.md:346-373` (paragraph)
- **citation**: `internal/storagens/callsites_test.go:189`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/storagens/callsites_test.go`  (298 lines)

**CLAIMING UNIT**

```md
> `roadrunner` is **not** a gRPC hop from `backend` in either column — it was folded in with jobsim-in-app and
> `backend` calls Judge0 directly. **Nor is `storage` one in EITHER column** — the production diagram above no
longer lists it, because the edge is dead there too: `storage`'s ECS service block is **gone from terraform
entirely**, and says so in-comment — at `9f8cb53` `storage/terraform/main.tf` is 18 lines (`:9-11`
*"The ECS service that used to live here is GONE (v9.0 'support-in-app')"*), the module kept only so the
buckets, CloudFront distribution and media DNS record keep their `prevent_destroy` guards (`:13-16`). It
read `service_desired_count = 0` at the intermediate `63bffc8`. And `STORAGE_RPC_ADDR` has **zero** reads
anywhere in `app`
(**3 hits in Go source** at `9d00a313`, every one a comment — `main.go:451`,
`internal/jobsimwiring/wiring.go:101`, `internal/storagens/callsites_test.go:189`. **Not 3 repo-wide**:
`git -C stack-demo/app grep -n STORAGE_RPC_ADDR 9d00a313` returns **29 lines across 18 files**, the rest
being CHANGELOG and `knowledge/*.md`. The Go scope is what carries the claim; the repo-wide form was a
mis-transcription of it). The earlier wording scoped this retraction to
*"locally"*, which left the prod edge affirmatively standing (corrected M257x iter-85). Platform
**`0dab54d`** ("storage-in-app,
> v9.0") deleted `STORAGE_RPC_ADDR` from `backend`'s env, dropped `storage` from `backend`'s `depends_on`
> — the replacement comment read *"storage removed at v9.0: served in-process by this container now"*
> — and moved the service to `profiles: [storage-legacy]`; **`838d907`
> (merged `0c91421`, 2026-08-05) then deleted the `storage` service and that profile outright**, so there
> is nothing left to opt into. The app side
> closed at **`app` `9d00a313`** (v1.367.0) and is still closed at `app` **origin/main**, where
> `STORAGE_RPC_ADDR` has **zero reads** and `main.go:504`
> says *"the standalone service takes no traffic and STORAGE_RPC_ADDR is gone."* That comment stood at
> `main.go:451` at `9d00a313`; **the line number moved without the code moving**, which is what a week of
> `app` commits costs an anchor (re-derived M257x iter-87). At the older `b948604`
> v1.366.0 the variable is still read — `internal/jobsimwiring/wiring.go:115` — so **state the ref you mean**. See
> [`roadrunner.md`](../services/roadrunner.md), [`storage
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

## 08-010
- **id**: `B08-010`
- **corpus site**: `corpus/architecture/architecture_overview.md:346-373` (paragraph)
- **citation**: `main.go:504`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> `roadrunner` is **not** a gRPC hop from `backend` in either column — it was folded in with jobsim-in-app and
> `backend` calls Judge0 directly. **Nor is `storage` one in EITHER column** — the production diagram above no
longer lists it, because the edge is dead there too: `storage`'s ECS service block is **gone from terraform
entirely**, and says so in-comment — at `9f8cb53` `storage/terraform/main.tf` is 18 lines (`:9-11`
*"The ECS service that used to live here is GONE (v9.0 'support-in-app')"*), the module kept only so the
buckets, CloudFront distribution and media DNS record keep their `prevent_destroy` guards (`:13-16`). It
read `service_desired_count = 0` at the intermediate `63bffc8`. And `STORAGE_RPC_ADDR` has **zero** reads
anywhere in `app`
(**3 hits in Go source** at `9d00a313`, every one a comment — `main.go:451`,
`internal/jobsimwiring/wiring.go:101`, `internal/storagens/callsites_test.go:189`. **Not 3 repo-wide**:
`git -C stack-demo/app grep -n STORAGE_RPC_ADDR 9d00a313` returns **29 lines across 18 files**, the rest
being CHANGELOG and `knowledge/*.md`. The Go scope is what carries the claim; the repo-wide form was a
mis-transcription of it). The earlier wording scoped this retraction to
*"locally"*, which left the prod edge affirmatively standing (corrected M257x iter-85). Platform
**`0dab54d`** ("storage-in-app,
> v9.0") deleted `STORAGE_RPC_ADDR` from `backend`'s env, dropped `storage` from `backend`'s `depends_on`
> — the replacement comment read *"storage removed at v9.0: served in-process by this container now"*
> — and moved the service to `profiles: [storage-legacy]`; **`838d907`
> (merged `0c91421`, 2026-08-05) then deleted the `storage` service and that profile outright**, so there
> is nothing left to opt into. The app side
> closed at **`app` `9d00a313`** (v1.367.0) and is still closed at `app` **origin/main**, where
> `STORAGE_RPC_ADDR` has **zero reads** and `main.go:504`
> says *"the standalone service takes no traffic and STORAGE_RPC_ADDR is gone."* That comment stood at
> `main.go:451` at `9d00a313`; **the line number moved without the code moving**, which is what a week of
> `app` commits costs an anchor (re-derived M257x iter-87). At the older `b948604`
> v1.366.0 the variable is still read — `internal/jobsimwiring/wiring.go:115` — so **state the ref you mean**. See
> [`roadrunner.md`](../services/roadrunner.md), [`storage
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

## 08-011
- **id**: `B08-011`
- **corpus site**: `corpus/architecture/architecture_overview.md:346-373` (paragraph)
- **citation**: `internal/jobsimwiring/wiring.go:115`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimwiring/wiring.go`  (283 lines)

**CLAIMING UNIT**

```md
> `roadrunner` is **not** a gRPC hop from `backend` in either column — it was folded in with jobsim-in-app and
> `backend` calls Judge0 directly. **Nor is `storage` one in EITHER column** — the production diagram above no
longer lists it, because the edge is dead there too: `storage`'s ECS service block is **gone from terraform
entirely**, and says so in-comment — at `9f8cb53` `storage/terraform/main.tf` is 18 lines (`:9-11`
*"The ECS service that used to live here is GONE (v9.0 'support-in-app')"*), the module kept only so the
buckets, CloudFront distribution and media DNS record keep their `prevent_destroy` guards (`:13-16`). It
read `service_desired_count = 0` at the intermediate `63bffc8`. And `STORAGE_RPC_ADDR` has **zero** reads
anywhere in `app`
(**3 hits in Go source** at `9d00a313`, every one a comment — `main.go:451`,
`internal/jobsimwiring/wiring.go:101`, `internal/storagens/callsites_test.go:189`. **Not 3 repo-wide**:
`git -C stack-demo/app grep -n STORAGE_RPC_ADDR 9d00a313` returns **29 lines across 18 files**, the rest
being CHANGELOG and `knowledge/*.md`. The Go scope is what carries the claim; the repo-wide form was a
mis-transcription of it). The earlier wording scoped this retraction to
*"locally"*, which left the prod edge affirmatively standing (corrected M257x iter-85). Platform
**`0dab54d`** ("storage-in-app,
> v9.0") deleted `STORAGE_RPC_ADDR` from `backend`'s env, dropped `storage` from `backend`'s `depends_on`
> — the replacement comment read *"storage removed at v9.0: served in-process by this container now"*
> — and moved the service to `profiles: [storage-legacy]`; **`838d907`
> (merged `0c91421`, 2026-08-05) then deleted the `storage` service and that profile outright**, so there
> is nothing left to opt into. The app side
> closed at **`app` `9d00a313`** (v1.367.0) and is still closed at `app` **origin/main**, where
> `STORAGE_RPC_ADDR` has **zero reads** and `main.go:504`
> says *"the standalone service takes no traffic and STORAGE_RPC_ADDR is gone."* That comment stood at
> `main.go:451` at `9d00a313`; **the line number moved without the code moving**, which is what a week of
> `app` commits costs an anchor (re-derived M257x iter-87). At the older `b948604`
> v1.366.0 the variable is still read — `internal/jobsimwiring/wiring.go:115` — so **state the ref you mean**. See
> [`roadrunner.md`](../services/roadrunner.md), [`storage
```

**CITED CONTENT**

```
   112  	cmsManager := jscms.New(cmsClient)
   113  	userClient := userReader // in-process (app's users domain reader); replaces the BACKEND_USERS_RPC_ADDR loopback
   114  	// Storage namespace is the S3 key PREFIX, bound at client construction — it is part of the
   115  	// object's physical address. The jobsim corpus (recordings, conversation clips, interaction
   116  	// audio + attachments, interview report CSVs) was written by the standalone jobsimulation
   117  	// service under "jobsimulation", so it must be read under "jobsimulation". Passing serviceName
   118  	// here ("backend") re-points every read of that historical corpus at a prefix that does not
```

## 08-012
- **id**: `B08-012`
- **corpus site**: `corpus/architecture/architecture_overview.md:390-401` (bullet)
- **citation**: `internal/authorization/gqlauthz/gqlauthz.go:222`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/authorization/gqlauthz/gqlauthz.go`  (263 lines)

**CLAIMING UNIT**

```md
2. **Authorization**: Sentinel is the centralized Casbin (RBAC/ABAC) authorization **engine** — **not a
   blanket applied to every API request.** ⚠️ This line said *"validates every API request"* until M257x
   iter-120. Measured at `app` `ad9f3c49`: the GraphQL `AuthorizationMiddleware` is a **viewer** gate
   with **six** paths that reach the resolver before the single Sentinel call
   (`internal/authorization/gqlauthz/gqlauthz.go:222`) — including *"viewer has no active org"* (`:190-191`)
   and *"the operation carries no `userId` variable"* (`:196-197`) — and the REST surface has **no
   BLANKET authz middleware**: authorization there is opt-in per group or per handler, and only 2 of its
   6 Echo groups carry a group-level one (`cbGate`, `internal/web/backend/gate.go:27-49`). The platform's
   own source calls the blanket gate **fail-open**
   (`graph/resolver_skiller_taxonomy_authz.go:53-66`). See
   [Security & Compliance → Layer 2](./security_compliance.md#layer-2-authorization) for the enumerated
   paths and the honest statement
```

**CITED CONTENT**

```
   219  			return next(ctx)
   220  		}
   221  
   222  		isAuthorized, err := authorizationManager.OrgCheckUserPermission(ctx, org.ID(), viewer.ID(), target, action)
   223  		if err != nil || !isAuthorized {
   224  			l.With("error", err).With("authorized", isAuthorized).Error("unauthorized from sentinel")
   225  			return permissionErrorResponse(ctx, action)
```

## 08-013
- **id**: `B08-013`
- **corpus site**: `corpus/architecture/architecture_overview.md:390-401` (bullet)
- **citation**: `internal/web/backend/gate.go:27-49`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/web/backend/gate.go`  (50 lines)

**CLAIMING UNIT**

```md
2. **Authorization**: Sentinel is the centralized Casbin (RBAC/ABAC) authorization **engine** — **not a
   blanket applied to every API request.** ⚠️ This line said *"validates every API request"* until M257x
   iter-120. Measured at `app` `ad9f3c49`: the GraphQL `AuthorizationMiddleware` is a **viewer** gate
   with **six** paths that reach the resolver before the single Sentinel call
   (`internal/authorization/gqlauthz/gqlauthz.go:222`) — including *"viewer has no active org"* (`:190-191`)
   and *"the operation carries no `userId` variable"* (`:196-197`) — and the REST surface has **no
   BLANKET authz middleware**: authorization there is opt-in per group or per handler, and only 2 of its
   6 Echo groups carry a group-level one (`cbGate`, `internal/web/backend/gate.go:27-49`). The platform's
   own source calls the blanket gate **fail-open**
   (`graph/resolver_skiller_taxonomy_authz.go:53-66`). See
   [Security & Compliance → Layer 2](./security_compliance.md#layer-2-authorization) for the enumerated
   paths and the honest statement
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

## 08-014
- **id**: `B08-014`
- **corpus site**: `corpus/architecture/dependency_map.md:7-7` (paragraph)
- **citation**: `docker-compose.yml:48`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
Sourced from `platform/docker-compose.yml` `depends_on:` declarations and service-address environment variables. At platform `0c91421` there are **zero `*_RPC_ADDR` variables** in any compose file — the `messenger` block that set the last four went with the service at `838d907` — so `AUTHORIZATION_ADDRESS=http://sentinel:8087` (`docker-compose.yml:48`) is the only cross-process **Connect-RPC** address left. **It is not the only service address, and `backend → sentinel` is not the only cross-process edge** (corrected M257x iter-102; this line previously said *"there is exactly **one** left"*): the same `backend` `environment:` block also carries `GOTENBERG_URL=http://gotenberg:3200` (`:57` — a second container, on the **default** `core` profile at `:183`, reached over **plain HTTP**, not Connect-RPC, at `app/internal/converter/gotenberg.go:31` @ `app` `ad9f3c49`), `JUDGE0_BASE_URL` (`docker-compose.yml:59`), `REDIS_ADDR` (`docker-compose.yml:66`) and the two Postgres DSNs `SUPABASE_DB_CONN` / `COPILOT_DB_CONN` (`docker-compose.yml:93-94`). The correctly-scoped form is [`architecture_overview.md`](architecture_overview.md)'s local-stack diagram — *"the only cross-process **RPC** edge out of backend on a core stack"* — and **§6. Document → PDF Conversion** below states the gotenberg half of it.
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

## 08-015
- **id**: `B08-015`
- **corpus site**: `corpus/architecture/dependency_map.md:7-7` (paragraph)
- **citation**: `app/internal/converter/gotenberg.go:31`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/converter/gotenberg.go`  (54 lines)

**CLAIMING UNIT**

```md
Sourced from `platform/docker-compose.yml` `depends_on:` declarations and service-address environment variables. At platform `0c91421` there are **zero `*_RPC_ADDR` variables** in any compose file — the `messenger` block that set the last four went with the service at `838d907` — so `AUTHORIZATION_ADDRESS=http://sentinel:8087` (`docker-compose.yml:48`) is the only cross-process **Connect-RPC** address left. **It is not the only service address, and `backend → sentinel` is not the only cross-process edge** (corrected M257x iter-102; this line previously said *"there is exactly **one** left"*): the same `backend` `environment:` block also carries `GOTENBERG_URL=http://gotenberg:3200` (`:57` — a second container, on the **default** `core` profile at `:183`, reached over **plain HTTP**, not Connect-RPC, at `app/internal/converter/gotenberg.go:31` @ `app` `ad9f3c49`), `JUDGE0_BASE_URL` (`docker-compose.yml:59`), `REDIS_ADDR` (`docker-compose.yml:66`) and the two Postgres DSNs `SUPABASE_DB_CONN` / `COPILOT_DB_CONN` (`docker-compose.yml:93-94`). The correctly-scoped form is [`architecture_overview.md`](architecture_overview.md)'s local-stack diagram — *"the only cross-process **RPC** edge out of backend on a core stack"* — and **§6. Document → PDF Conversion** below states the gotenberg half of it.
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

## 08-016
- **id**: `B08-016`
- **corpus site**: `corpus/architecture/dependency_map.md:7-7` (paragraph)
- **citation**: `docker-compose.yml:59`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
Sourced from `platform/docker-compose.yml` `depends_on:` declarations and service-address environment variables. At platform `0c91421` there are **zero `*_RPC_ADDR` variables** in any compose file — the `messenger` block that set the last four went with the service at `838d907` — so `AUTHORIZATION_ADDRESS=http://sentinel:8087` (`docker-compose.yml:48`) is the only cross-process **Connect-RPC** address left. **It is not the only service address, and `backend → sentinel` is not the only cross-process edge** (corrected M257x iter-102; this line previously said *"there is exactly **one** left"*): the same `backend` `environment:` block also carries `GOTENBERG_URL=http://gotenberg:3200` (`:57` — a second container, on the **default** `core` profile at `:183`, reached over **plain HTTP**, not Connect-RPC, at `app/internal/converter/gotenberg.go:31` @ `app` `ad9f3c49`), `JUDGE0_BASE_URL` (`docker-compose.yml:59`), `REDIS_ADDR` (`docker-compose.yml:66`) and the two Postgres DSNs `SUPABASE_DB_CONN` / `COPILOT_DB_CONN` (`docker-compose.yml:93-94`). The correctly-scoped form is [`architecture_overview.md`](architecture_overview.md)'s local-stack diagram — *"the only cross-process **RPC** edge out of backend on a core stack"* — and **§6. Document → PDF Conversion** below states the gotenberg half of it.
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

## 08-017
- **id**: `B08-017`
- **corpus site**: `corpus/architecture/dependency_map.md:7-7` (paragraph)
- **citation**: `docker-compose.yml:66`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
Sourced from `platform/docker-compose.yml` `depends_on:` declarations and service-address environment variables. At platform `0c91421` there are **zero `*_RPC_ADDR` variables** in any compose file — the `messenger` block that set the last four went with the service at `838d907` — so `AUTHORIZATION_ADDRESS=http://sentinel:8087` (`docker-compose.yml:48`) is the only cross-process **Connect-RPC** address left. **It is not the only service address, and `backend → sentinel` is not the only cross-process edge** (corrected M257x iter-102; this line previously said *"there is exactly **one** left"*): the same `backend` `environment:` block also carries `GOTENBERG_URL=http://gotenberg:3200` (`:57` — a second container, on the **default** `core` profile at `:183`, reached over **plain HTTP**, not Connect-RPC, at `app/internal/converter/gotenberg.go:31` @ `app` `ad9f3c49`), `JUDGE0_BASE_URL` (`docker-compose.yml:59`), `REDIS_ADDR` (`docker-compose.yml:66`) and the two Postgres DSNs `SUPABASE_DB_CONN` / `COPILOT_DB_CONN` (`docker-compose.yml:93-94`). The correctly-scoped form is [`architecture_overview.md`](architecture_overview.md)'s local-stack diagram — *"the only cross-process **RPC** edge out of backend on a core stack"* — and **§6. Document → PDF Conversion** below states the gotenberg half of it.
```

**CITED CONTENT**

```
    63        - MEDIA_URL=https://media.anthropos.work
    64        - META_PORT=8084
    65        - PORT=8082
    66        - REDIS_ADDR=redis:6379
    67        - REDIS_STREAMS_INDEX=4
    68        - REDIS_WORKER_INDEX=0
    69        - RPC_PORT=8083
```

## 08-018
- **id**: `B08-018`
- **corpus site**: `corpus/architecture/dependency_map.md:7-7` (paragraph)
- **citation**: `docker-compose.yml:93-94`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
Sourced from `platform/docker-compose.yml` `depends_on:` declarations and service-address environment variables. At platform `0c91421` there are **zero `*_RPC_ADDR` variables** in any compose file — the `messenger` block that set the last four went with the service at `838d907` — so `AUTHORIZATION_ADDRESS=http://sentinel:8087` (`docker-compose.yml:48`) is the only cross-process **Connect-RPC** address left. **It is not the only service address, and `backend → sentinel` is not the only cross-process edge** (corrected M257x iter-102; this line previously said *"there is exactly **one** left"*): the same `backend` `environment:` block also carries `GOTENBERG_URL=http://gotenberg:3200` (`:57` — a second container, on the **default** `core` profile at `:183`, reached over **plain HTTP**, not Connect-RPC, at `app/internal/converter/gotenberg.go:31` @ `app` `ad9f3c49`), `JUDGE0_BASE_URL` (`docker-compose.yml:59`), `REDIS_ADDR` (`docker-compose.yml:66`) and the two Postgres DSNs `SUPABASE_DB_CONN` / `COPILOT_DB_CONN` (`docker-compose.yml:93-94`). The correctly-scoped form is [`architecture_overview.md`](architecture_overview.md)'s local-stack diagram — *"the only cross-process **RPC** edge out of backend on a core stack"* — and **§6. Document → PDF Conversion** below states the gotenberg half of it.
```

**CITED CONTENT**

```
    90        # without editing this file. To exercise either one locally, set it in .env — and
    91        # know that messenger then attaches to the LIVE Redis consumer group and
    92        # customerio-sync writes real Brevo contacts.
    93        - SUPABASE_DB_CONN=postgresql://postgres@postgresql:5432/postgres?sslmode=disable
    94        - COPILOT_DB_CONN=postgresql://postgres@postgresql:5432/postgres?sslmode=disable
    95      networks:
    96        - app-network
    97      # jobsim-in-app's Chime/LiveKit recording managers use the AWS SDK default
```

## 08-019
- **id**: `B08-019`
- **corpus site**: `corpus/architecture/dependency_map.md:13-13` (table-row)
- **citation**: `docker-compose.yml:101-109`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| **Backend** (`app`) — the monolith | Sentinel, Redis, Postgres — the whole of its compose `depends_on` at platform `0c91421` (`docker-compose.yml:101-109`). **The cms, storage, messenger and customerio-sync edges are gone as *facts*, not merely moved:** `d11a403` deleted the cms container and `838d907` deleted the other three outright — compose says so in-line where those edges used to be (`:102-103`, *"storage, messenger and customerio-sync are not services any more — this one container serves all three in-process."*). Gotenberg is a runtime HTTP call with no startup-order dep | Postgres (`public` schema; `pgvector` in `extensions` — skiller embeddings, skill-path sessions, the 23 jobsim run-state tables, the cms similarity/Studio tables), Redis, **Clerk**, **Directus**, **Judge0**, **LiveKit**, **AWS Chime**, **AI Providers** |
```

**CITED CONTENT**

```
    98      # credential chain — the mount the standalone jobsimulation container had.
    99      volumes:
   100        - $HOME/.aws/credentials:/root/.aws/credentials:ro
   101      depends_on:
   102        # storage, messenger and customerio-sync are not services any more — this one
   103        # container serves all three in-process.
   104        redis:
   105          condition: service_healthy
   106        postgresql:
   107          condition: service_healthy
   108        sentinel:
   109          condition: service_started
   110      profiles: [core, backend, all]
   111  
   112    studio-desk:
```

## 08-020
- **id**: `B08-020`
- **corpus site**: `corpus/architecture/dependency_map.md:19-19` (table-row)
- **citation**: `docker-compose.yml:126-131`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| ~~**Storage**~~ | **Merged into `app`** (v9.0 "support-in-app") — `backend` serves object storage in-process; platform `838d907` (merged `0c91421`, 2026-08-05) deleted the compose service **and** the `repos.yml` entry | *(**no container** at platform `0c91421`.)* Its data path was **S3** only — plus a local-filesystem fallback per bucket. While the service existed it declared `depends_on: redis, postgresql` (both `service_healthy`) and **read neither**: no `DB_CONNECTION`/`REDIS_ADDR` among the **seven** variables in its compose env. Measured in the platform's own compose at `0dab54d` — `docker-compose.yml:126-131` for the `depends_on`, `:117-123` for the env block: `AWS_DEFAULT_REGION`, `AWS_REGION`, `ENVIRONMENT`, `PORT`, `RPC_PORT`, `SERVICE_NAME`, `STORAGE_S3_PUBLIC_BUCKET`. ⚠️ **This said "eight … `:116-123`" until M257x iter-115**, and the eight is the cardinality of the cited LINE RANGE, not of the set the sentence ranges over — `:116` is the `environment:` **key**. Enumerate the set before you count it; the predicate the count serves (neither variable is present) survives, and only the number was false. **This clause used to credit [`storage.md`](../services/storage.md)`:40,47` with saying it; that attribution was false and is withdrawn** — `storage.md` contains none of `depends_on`, `DB_CONNECTION`, `REDIS_ADDR`, `go.mod`, `redis` or `postgres` at any line, and never has (`git log -S` finds no commit adding one). The compose file is the source; the service doc was never a second witness. The compose ordering edge and the runtime data path disagreed **by design**, which is why [`service_taxonomy.md`](service_taxonomy.md) used to list postgresql + redis under storage. Corrected M257x iter-49, over-corrected to `-`, re-corrected iter-52; the row went historical when the service was deleted |
```

**CITED CONTENT**

```
   123        - "9000:9000"
   124        - "9100:9100"
   125      env_file:
   126        - .env
   127      environment:
   128        - CLERK_SIGN_IN_URL=http://localhost:3000/login
   129        - DEFAULT_MODEL=gpt-4o
   130        - ENVIRONMENT=development
   131        - FRONTEND_PORT=9100
   132        - NODE_ENV=development
   133        - PORT=9000
   134        - VITE_ENVIRONMENT=production
```

## 08-021
- **id**: `B08-021`
- **corpus site**: `corpus/architecture/dependency_map.md:22-22` (table-row)
- **citation**: `docker-compose.yml:154`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| ~~**CustomerIO Sync**~~ | **Merged into `app`** — it runs on `backend`'s asynq scheduler; `838d907` deleted the compose service. It was still in the **`all`** profile until then (`profiles: [customerio-sync, all]`, `0dab54d:docker-compose.yml:154`) — **that half is true; the "second Brevo pusher" half is false and is retracted** (corrected M257x iter-102; this row previously said `make up-all` started *a second Brevo contact pusher alongside `backend`'s own*). `make up-all` started exactly **one** Brevo contact pusher: `backend`'s own in-process one is gated behind `CUSTOMERIO_SYNC_ENABLED`, which is unset — and therefore **off** — on a developer machine (`docker-compose.yml:84-92` @ `0c91421`). See `service_taxonomy.md:101-103` | **Customer.io** — now `app`'s, gated **OFF** on a developer machine behind `CUSTOMERIO_SYNC_ENABLED` |
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

## 08-022
- **id**: `B08-022`
- **corpus site**: `corpus/architecture/dependency_map.md:22-22` (table-row)
- **citation**: `docker-compose.yml:84-92`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| ~~**CustomerIO Sync**~~ | **Merged into `app`** — it runs on `backend`'s asynq scheduler; `838d907` deleted the compose service. It was still in the **`all`** profile until then (`profiles: [customerio-sync, all]`, `0dab54d:docker-compose.yml:154`) — **that half is true; the "second Brevo pusher" half is false and is retracted** (corrected M257x iter-102; this row previously said `make up-all` started *a second Brevo contact pusher alongside `backend`'s own*). `make up-all` started exactly **one** Brevo contact pusher: `backend`'s own in-process one is gated behind `CUSTOMERIO_SYNC_ENABLED`, which is unset — and therefore **off** — on a developer machine (`docker-compose.yml:84-92` @ `0c91421`). See `service_taxonomy.md:101-103` | **Customer.io** — now `app`'s, gated **OFF** on a developer machine behind `CUSTOMERIO_SYNC_ENABLED` |
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

## 08-023
- **id**: `B08-023`
- **corpus site**: `corpus/architecture/dependency_map.md:24-24` (table-row)
- **citation**: `docker-compose.yml:138-140`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| **Studio-Desk** (opt-in profile) | `backend`'s GraphQL endpoint directly (`:8082/graphql/query`) — the router it used to depend on is gone locally. Compose `depends_on` is **`backend`, and only `backend`** (`docker-compose.yml:138-140`) — the cms edge went with the cms container at `d11a403`, so this is now a one-edge block, not a two-edge one | **Clerk**, **OpenAI / Azure OpenAI / Anthropic** (Copilot, via `AI_PROVIDER_CHAIN`) |
```

**CITED CONTENT**

```
   135        - VITE_GRAPHQL_ENDPOINT=http://localhost:8082/graphql/query
   136      networks:
   137        - app-network
   138      depends_on:
   139        backend:
   140          condition: service_started
   141      profiles: [studio-desk, all]
   142  
   143    next-web-app:
```

## 08-024
- **id**: `B08-024`
- **corpus site**: `corpus/architecture/dependency_map.md:27-33` (paragraph)
- **citation**: `cms/terraform/main.tf:39`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/terraform/main.tf`  (192 lines)

**CLAIMING UNIT**

```md
> **Skiller merged into app (July 2026):** the standalone skiller service is gone from the compose file. Its RPC surface is now served by **backend** — consumers keep the `SKILLER_RPC_ADDR` env var, which read `http://backend:8083` — **and was already reading it before `d11a403`, so that commit did not re-point this one** (it moved `CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR` only; measured at `d11a403^`, M257x iter-115); but `messenger` was the last thing that set it locally, and since `838d907` deleted that service **no compose file sets it at all**. **The production value is not asserted here in either direction** — no `.tf` file in any clone names `http://backend.internal.anthropos:8081`, and the deciding declaration lives in `infrastructure`, which is in no clone set; see [Backend](../services/backend.md) for the derivation and the fenced unmeasurable-claims convention. See [Backend](../services/backend.md) and the [skiller stub](../services/skiller.md).
>
> **Skillpath merged into app (skillpath-in-app, M502→M507):** the standalone skillpath service is gone from the compose file / repos.yml / supergraph. Its skill-path progression engine now runs **in-process inside `app`**, with session state in `public.skill_path_sessions` (the legacy `skillpath` schema is an empty husk). See [Backend](../services/backend.md) and the [skillpath stub](../services/skillpath.md).
>
> **Jobsimulation + cms merged into app (jobsim-in-app, cms-in-app v8.0):** the last two subgraph services are gone from the **supergraph** — the federation now composes **one** subgraph. **And at platform `0c91421` they are gone from compose and from `repos.yml` too** — 7 compose services in the effective topology (5 declared in `docker-compose.yml`), 4 repo entries, neither list containing cms or jobsimulation. (It was 10 and 6 at `0dab54d`, before `838d907` dropped the last three support containers.) (Both still started as unfederated husks right up to `d11a403` — at `2adcf71` compose still declared all three of cms, jobsimulation and roadrunner — and `d11a403` is what changed it. `2adcf71` is the *router* deletion; do not conflate the two commits.) Their tables were re-created in `public` (the legacy `jobsimulation` and `cms` schemas are non-authoritative). **Their production dispositions have since diverged — do not state them as one:** `cms`'s ECS module is still declared **in its own re
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

## 08-025
- **id**: `B08-025`
- **corpus site**: `corpus/architecture/dependency_map.md:27-33` (paragraph)
- **citation**: `jobsimulation/terraform/main.tf:15-40`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/jobsimulation/terraform/main.tf`  (344 lines)

**CLAIMING UNIT**

```md
> **Skiller merged into app (July 2026):** the standalone skiller service is gone from the compose file. Its RPC surface is now served by **backend** — consumers keep the `SKILLER_RPC_ADDR` env var, which read `http://backend:8083` — **and was already reading it before `d11a403`, so that commit did not re-point this one** (it moved `CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR` only; measured at `d11a403^`, M257x iter-115); but `messenger` was the last thing that set it locally, and since `838d907` deleted that service **no compose file sets it at all**. **The production value is not asserted here in either direction** — no `.tf` file in any clone names `http://backend.internal.anthropos:8081`, and the deciding declaration lives in `infrastructure`, which is in no clone set; see [Backend](../services/backend.md) for the derivation and the fenced unmeasurable-claims convention. See [Backend](../services/backend.md) and the [skiller stub](../services/skiller.md).
>
> **Skillpath merged into app (skillpath-in-app, M502→M507):** the standalone skillpath service is gone from the compose file / repos.yml / supergraph. Its skill-path progression engine now runs **in-process inside `app`**, with session state in `public.skill_path_sessions` (the legacy `skillpath` schema is an empty husk). See [Backend](../services/backend.md) and the [skillpath stub](../services/skillpath.md).
>
> **Jobsimulation + cms merged into app (jobsim-in-app, cms-in-app v8.0):** the last two subgraph services are gone from the **supergraph** — the federation now composes **one** subgraph. **And at platform `0c91421` they are gone from compose and from `repos.yml` too** — 7 compose services in the effective topology (5 declared in `docker-compose.yml`), 4 repo entries, neither list containing cms or jobsimulation. (It was 10 and 6 at `0dab54d`, before `838d907` dropped the last three support containers.) (Both still started as unfederated husks right up to `d11a403` — at `2adcf71` compose still declared all three of cms, jobsimulation and roadrunner — and `d11a403` is what changed it. `2adcf71` is the *router* deletion; do not conflate the two commits.) Their tables were re-created in `public` (the legacy `jobsimulation` and `cms` schemas are non-authoritative). **Their production dispositions have since diverged — do not state them as one:** `cms`'s ECS module is still declared **in its own re
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

## 08-026
- **id**: `B08-026`
- **corpus site**: `corpus/architecture/dependency_map.md:48-48` (table-row)
- **citation**: `go.mod:9`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/go.mod`  (296 lines)

**CLAIMING UNIT**

```md
| **ai** | **No repo a stack builds** — `app/go.mod` required it up to `b948604f` and `1e457fa70` removed the requirement when the library was folded in as `app/internal/ai` (guarded one-way by `internal/ai/module_import_guard_test.go`). Still required by the frozen `cms` (`go.mod:9`) and `jobsimulation` (`go.mod:11`) repos, both `v1.40.2`, which nothing clones or compiles. Go services only, never Studio-Desk. Cost & routing live in the consumers, not the lib |
```

**CITED CONTENT**

```
     6  	code.sajari.com/docconv v1.3.8
     7  	connectrpc.com/connect v1.20.0
     8  	entgo.io/ent v0.14.6
     9  	github.com/99designs/gqlgen v0.17.94
    10  	github.com/DATA-DOG/go-sqlmock v1.5.2
    11  	github.com/ThreeDotsLabs/watermill v1.5.2
    12  	github.com/ThreeDotsLabs/watermill-redisstream v1.4.5
```

## 08-027
- **id**: `B08-027`
- **corpus site**: `corpus/architecture/dependency_map.md:48-48` (table-row)
- **citation**: `go.mod:11`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/go.mod`  (296 lines)

**CLAIMING UNIT**

```md
| **ai** | **No repo a stack builds** — `app/go.mod` required it up to `b948604f` and `1e457fa70` removed the requirement when the library was folded in as `app/internal/ai` (guarded one-way by `internal/ai/module_import_guard_test.go`). Still required by the frozen `cms` (`go.mod:9`) and `jobsimulation` (`go.mod:11`) repos, both `v1.40.2`, which nothing clones or compiles. Go services only, never Studio-Desk. Cost & routing live in the consumers, not the lib |
```

**CITED CONTENT**

```
     8  	entgo.io/ent v0.14.6
     9  	github.com/99designs/gqlgen v0.17.94
    10  	github.com/DATA-DOG/go-sqlmock v1.5.2
    11  	github.com/ThreeDotsLabs/watermill v1.5.2
    12  	github.com/ThreeDotsLabs/watermill-redisstream v1.4.5
    13  	github.com/anthropics/anthropic-sdk-go v1.61.0
    14  	github.com/anthropos-work/analytics-go v0.3.1
```

## 08-028
- **id**: `B08-028`
- **corpus site**: `corpus/architecture/dependency_map.md:50-50` (table-row)
- **citation**: `app/go.mod:20`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/go.mod`  (296 lines)

**CLAIMING UNIT**

```md
| **taxonomy** | **node-id library** (not data): of the Go repos a stack clones at platform `0c91421` — direct: app (`app/go.mod:20` @ `b948604f`); indirect: sentinel (`sentinel/go.mod:21` @ `88bc5592`). The messenger (direct) / storage (indirect) requirements are frozen — `838d907` dropped both from `repos.yml` |
```

**CITED CONTENT**

```
    17  	github.com/anthropos-work/storage v0.15.2
    18  	github.com/anthropos-work/taxonomy v1.2.0
    19  	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
    20  	github.com/asticode/go-astisub v0.42.0
    21  	github.com/avast/retry-go/v4 v4.7.0
    22  	github.com/aws/aws-sdk-go-v2 v1.43.0
    23  	github.com/aws/aws-sdk-go-v2/config v1.32.31
```

## 08-029
- **id**: `B08-029`
- **corpus site**: `corpus/architecture/dependency_map.md:50-50` (table-row)
- **citation**: `sentinel/go.mod:21`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/sentinel/go.mod`  (54 lines)

**CLAIMING UNIT**

```md
| **taxonomy** | **node-id library** (not data): of the Go repos a stack clones at platform `0c91421` — direct: app (`app/go.mod:20` @ `b948604f`); indirect: sentinel (`sentinel/go.mod:21` @ `88bc5592`). The messenger (direct) / storage (indirect) requirements are frozen — `838d907` dropped both from `repos.yml` |
```

**CITED CONTENT**

```
    18  require (
    19  	github.com/99designs/gqlgen v0.17.94 // indirect
    20  	github.com/agnivade/levenshtein v1.2.1 // indirect
    21  	github.com/anthropos-work/taxonomy v1.2.0 // indirect
    22  	github.com/bmatcuk/doublestar/v4 v4.10.0 // indirect
    23  	github.com/brunoscheufler/aws-ecs-metadata-go v0.0.0-20221221133751-67e37ae746cd // indirect
    24  	github.com/casbin/govaluate v1.10.0 // indirect
```

## 08-030
- **id**: `B08-030`
- **corpus site**: `corpus/architecture/dependency_map.md:58-58` (table-row)
- **citation**: `messenger/internal/flow/flow.go:72`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/messenger/internal/flow/flow.go`  (135 lines)

**CLAIMING UNIT**

```md
| `backend` | App | App (cms **domain** in `app`) — **and two more in the code, neither of which runs on a stock `core` stack.** (a) the standalone `messenger` service subscribed to this same stream — `messenger/internal/flow/flow.go:72` is a literal `AddSubscriber("backend", …)` over 21 live handlers — though since `838d907` deleted that compose service and its profile there is no longer any way to start it, so this subscriber now exists only in the frozen repo. (b) Since the v9.0 messenger-in-app fold (`app` `9d00a313`), **`app` itself can run a second subscriber server** on messenger's *own* Redis consumer group, attaching by messenger's own stream literals — `StreamBackend` is the literal `"backend"` (`app/internal/messenger/flow/streams.go:65`), named from `main.go:1493`. **That one is opt-in, and OFF by default** — corrected M257x iter-102; this cell previously said *"That one is not opt-in: it is the stock `core` selection"* and cited `app/main.go:1442`, which is a comment line at the newer refs and does not exist at all at `b948604f` (`main.go` is 1361 lines there). Measured: the whole block sits behind `if messengerEnabled {` at **`app/main.go:1445`**, fed by `mustSubsystemSwitch(envMessengerEnabled)` at `:285` (`env_guards.go:61` — `envMessengerEnabled = "MESSENGER_ENABLED"`), and the code states it in-line at `main.go:1437` @ `ad9f3c49`: *"ALL OF IT is behind MESSENGER_ENABLED."* Compose deliberately sets nothing there and says why where the variables would have gone (`docker-compose.yml:84-92` @ `0c91421`: *"which default to OFF on a developer machine"*), so **unset = OFF** — which is what the Messenger row above has said all along. `app` anchors here @ `app` `ad9f3c49` (`origin/main` on 2026-08-06; they hold identically at `2035f9a4`). What `d11a403` removed was only the `cms` husk container, and that was never the whole answer | User/org updates |
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
```

## 08-031
- **id**: `B08-031`
- **corpus site**: `corpus/architecture/dependency_map.md:58-58` (table-row)
- **citation**: `app/internal/messenger/flow/streams.go:65`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/messenger/flow/streams.go`  (143 lines)

**CLAIMING UNIT**

```md
| `backend` | App | App (cms **domain** in `app`) — **and two more in the code, neither of which runs on a stock `core` stack.** (a) the standalone `messenger` service subscribed to this same stream — `messenger/internal/flow/flow.go:72` is a literal `AddSubscriber("backend", …)` over 21 live handlers — though since `838d907` deleted that compose service and its profile there is no longer any way to start it, so this subscriber now exists only in the frozen repo. (b) Since the v9.0 messenger-in-app fold (`app` `9d00a313`), **`app` itself can run a second subscriber server** on messenger's *own* Redis consumer group, attaching by messenger's own stream literals — `StreamBackend` is the literal `"backend"` (`app/internal/messenger/flow/streams.go:65`), named from `main.go:1493`. **That one is opt-in, and OFF by default** — corrected M257x iter-102; this cell previously said *"That one is not opt-in: it is the stock `core` selection"* and cited `app/main.go:1442`, which is a comment line at the newer refs and does not exist at all at `b948604f` (`main.go` is 1361 lines there). Measured: the whole block sits behind `if messengerEnabled {` at **`app/main.go:1445`**, fed by `mustSubsystemSwitch(envMessengerEnabled)` at `:285` (`env_guards.go:61` — `envMessengerEnabled = "MESSENGER_ENABLED"`), and the code states it in-line at `main.go:1437` @ `ad9f3c49`: *"ALL OF IT is behind MESSENGER_ENABLED."* Compose deliberately sets nothing there and says why where the variables would have gone (`docker-compose.yml:84-92` @ `0c91421`: *"which default to OFF on a developer machine"*), so **unset = OFF** — which is what the Messenger row above has said all along. `app` anchors here @ `app` `ad9f3c49` (`origin/main` on 2026-08-06; they hold identically at `2035f9a4`). What `d11a403` removed was only the `cms` husk container, and that was never the whole answer | User/org updates |
```

**CITED CONTENT**

```
    62  const ConsumerGroup = "messenger"
    63  
    64  const (
    65  	StreamBackend       = "backend"
    66  	StreamJobSimulation = "jobsimulation"
    67  	StreamCMS           = "cms"
    68  )
```

## 08-032
- **id**: `B08-032`
- **corpus site**: `corpus/architecture/dependency_map.md:58-58` (table-row)
- **citation**: `main.go:1493`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| `backend` | App | App (cms **domain** in `app`) — **and two more in the code, neither of which runs on a stock `core` stack.** (a) the standalone `messenger` service subscribed to this same stream — `messenger/internal/flow/flow.go:72` is a literal `AddSubscriber("backend", …)` over 21 live handlers — though since `838d907` deleted that compose service and its profile there is no longer any way to start it, so this subscriber now exists only in the frozen repo. (b) Since the v9.0 messenger-in-app fold (`app` `9d00a313`), **`app` itself can run a second subscriber server** on messenger's *own* Redis consumer group, attaching by messenger's own stream literals — `StreamBackend` is the literal `"backend"` (`app/internal/messenger/flow/streams.go:65`), named from `main.go:1493`. **That one is opt-in, and OFF by default** — corrected M257x iter-102; this cell previously said *"That one is not opt-in: it is the stock `core` selection"* and cited `app/main.go:1442`, which is a comment line at the newer refs and does not exist at all at `b948604f` (`main.go` is 1361 lines there). Measured: the whole block sits behind `if messengerEnabled {` at **`app/main.go:1445`**, fed by `mustSubsystemSwitch(envMessengerEnabled)` at `:285` (`env_guards.go:61` — `envMessengerEnabled = "MESSENGER_ENABLED"`), and the code states it in-line at `main.go:1437` @ `ad9f3c49`: *"ALL OF IT is behind MESSENGER_ENABLED."* Compose deliberately sets nothing there and says why where the variables would have gone (`docker-compose.yml:84-92` @ `0c91421`: *"which default to OFF on a developer machine"*), so **unset = OFF** — which is what the Messenger row above has said all along. `app` anchors here @ `app` `ad9f3c49` (`origin/main` on 2026-08-06; they hold identically at `2035f9a4`). What `d11a403` removed was only the `cms` husk container, and that was never the whole answer | User/org updates |
```

**CITED CONTENT**

```
  1490  			if err := verifyConsumerGroupExists(serverContext, messengerRedisClient, msgflow.ConsumerGroup,
  1491  				// messenger's OWN literals: Subscribe() hardcodes these three, so these are the
  1492  				// streams app is about to attach to, regardless of app's own env-derived names.
  1493  				msgflow.StreamBackend, msgflow.StreamJobSimulation, msgflow.StreamCMS); err != nil {
  1494  				log.Fatalf("messenger-in-app: %v; app TAKES OVER the standalone's consumer group rather "+
  1495  					"than creating one, so a missing group means the entire retained stream would be "+
  1496  					"replayed through 24 email handlers", err)
```

## 08-033
- **id**: `B08-033`
- **corpus site**: `corpus/architecture/dependency_map.md:58-58` (table-row)
- **citation**: `app/main.go:1442`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| `backend` | App | App (cms **domain** in `app`) — **and two more in the code, neither of which runs on a stock `core` stack.** (a) the standalone `messenger` service subscribed to this same stream — `messenger/internal/flow/flow.go:72` is a literal `AddSubscriber("backend", …)` over 21 live handlers — though since `838d907` deleted that compose service and its profile there is no longer any way to start it, so this subscriber now exists only in the frozen repo. (b) Since the v9.0 messenger-in-app fold (`app` `9d00a313`), **`app` itself can run a second subscriber server** on messenger's *own* Redis consumer group, attaching by messenger's own stream literals — `StreamBackend` is the literal `"backend"` (`app/internal/messenger/flow/streams.go:65`), named from `main.go:1493`. **That one is opt-in, and OFF by default** — corrected M257x iter-102; this cell previously said *"That one is not opt-in: it is the stock `core` selection"* and cited `app/main.go:1442`, which is a comment line at the newer refs and does not exist at all at `b948604f` (`main.go` is 1361 lines there). Measured: the whole block sits behind `if messengerEnabled {` at **`app/main.go:1445`**, fed by `mustSubsystemSwitch(envMessengerEnabled)` at `:285` (`env_guards.go:61` — `envMessengerEnabled = "MESSENGER_ENABLED"`), and the code states it in-line at `main.go:1437` @ `ad9f3c49`: *"ALL OF IT is behind MESSENGER_ENABLED."* Compose deliberately sets nothing there and says why where the variables would have gone (`docker-compose.yml:84-92` @ `0c91421`: *"which default to OFF on a developer machine"*), so **unset = OFF** — which is what the Messenger row above has said all along. `app` anchors here @ `app` `ad9f3c49` (`origin/main` on 2026-08-06; they hold identically at `2035f9a4`). What `d11a403` removed was only the `cms` husk container, and that was never the whole answer | User/org updates |
```

**CITED CONTENT**

```
  1439  	// messenger's LIVE consumer group and starts claiming entries off the real
  1440  	// streams. A developer running `go run .` against a shared Redis would take
  1441  	// delivery of production notifications and — because the group cursor advances —
  1442  	// they would not be redelivered to the real backend. The switch has to sit
  1443  	// outside the whole block, not around the last step of it.
  1444  	var messengerSubServer *pubsub.SubscriberServer
  1445  	if messengerEnabled {
```

## 08-034
- **id**: `B08-034`
- **corpus site**: `corpus/architecture/dependency_map.md:58-58` (table-row)
- **citation**: `app/main.go:1445`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| `backend` | App | App (cms **domain** in `app`) — **and two more in the code, neither of which runs on a stock `core` stack.** (a) the standalone `messenger` service subscribed to this same stream — `messenger/internal/flow/flow.go:72` is a literal `AddSubscriber("backend", …)` over 21 live handlers — though since `838d907` deleted that compose service and its profile there is no longer any way to start it, so this subscriber now exists only in the frozen repo. (b) Since the v9.0 messenger-in-app fold (`app` `9d00a313`), **`app` itself can run a second subscriber server** on messenger's *own* Redis consumer group, attaching by messenger's own stream literals — `StreamBackend` is the literal `"backend"` (`app/internal/messenger/flow/streams.go:65`), named from `main.go:1493`. **That one is opt-in, and OFF by default** — corrected M257x iter-102; this cell previously said *"That one is not opt-in: it is the stock `core` selection"* and cited `app/main.go:1442`, which is a comment line at the newer refs and does not exist at all at `b948604f` (`main.go` is 1361 lines there). Measured: the whole block sits behind `if messengerEnabled {` at **`app/main.go:1445`**, fed by `mustSubsystemSwitch(envMessengerEnabled)` at `:285` (`env_guards.go:61` — `envMessengerEnabled = "MESSENGER_ENABLED"`), and the code states it in-line at `main.go:1437` @ `ad9f3c49`: *"ALL OF IT is behind MESSENGER_ENABLED."* Compose deliberately sets nothing there and says why where the variables would have gone (`docker-compose.yml:84-92` @ `0c91421`: *"which default to OFF on a developer machine"*), so **unset = OFF** — which is what the Messenger row above has said all along. `app` anchors here @ `app` `ad9f3c49` (`origin/main` on 2026-08-06; they hold identically at `2035f9a4`). What `d11a403` removed was only the `cms` husk container, and that was never the whole answer | User/org updates |
```

**CITED CONTENT**

```
  1442  	// they would not be redelivered to the real backend. The switch has to sit
  1443  	// outside the whole block, not around the last step of it.
  1444  	var messengerSubServer *pubsub.SubscriberServer
  1445  	if messengerEnabled {
  1446  		messengerRedisClient, err := redis.NewClient(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_STREAMS_INDEX"))
  1447  		if err != nil {
  1448  			log.Fatalf("messenger-in-app: can't open its Redis client: %v", err)
```

## 08-035
- **id**: `B08-035`
- **corpus site**: `corpus/architecture/dependency_map.md:58-58` (table-row)
- **citation**: `env_guards.go:61`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/env_guards.go`  (202 lines)

**CLAIMING UNIT**

```md
| `backend` | App | App (cms **domain** in `app`) — **and two more in the code, neither of which runs on a stock `core` stack.** (a) the standalone `messenger` service subscribed to this same stream — `messenger/internal/flow/flow.go:72` is a literal `AddSubscriber("backend", …)` over 21 live handlers — though since `838d907` deleted that compose service and its profile there is no longer any way to start it, so this subscriber now exists only in the frozen repo. (b) Since the v9.0 messenger-in-app fold (`app` `9d00a313`), **`app` itself can run a second subscriber server** on messenger's *own* Redis consumer group, attaching by messenger's own stream literals — `StreamBackend` is the literal `"backend"` (`app/internal/messenger/flow/streams.go:65`), named from `main.go:1493`. **That one is opt-in, and OFF by default** — corrected M257x iter-102; this cell previously said *"That one is not opt-in: it is the stock `core` selection"* and cited `app/main.go:1442`, which is a comment line at the newer refs and does not exist at all at `b948604f` (`main.go` is 1361 lines there). Measured: the whole block sits behind `if messengerEnabled {` at **`app/main.go:1445`**, fed by `mustSubsystemSwitch(envMessengerEnabled)` at `:285` (`env_guards.go:61` — `envMessengerEnabled = "MESSENGER_ENABLED"`), and the code states it in-line at `main.go:1437` @ `ad9f3c49`: *"ALL OF IT is behind MESSENGER_ENABLED."* Compose deliberately sets nothing there and says why where the variables would have gone (`docker-compose.yml:84-92` @ `0c91421`: *"which default to OFF on a developer machine"*), so **unset = OFF** — which is what the Messenger row above has said all along. `app` anchors here @ `app` `ad9f3c49` (`origin/main` on 2026-08-06; they hold identically at `2035f9a4`). What `d11a403` removed was only the `cms` husk container, and that was never the whole answer | User/org updates |
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

## 08-036
- **id**: `B08-036`
- **corpus site**: `corpus/architecture/dependency_map.md:58-58` (table-row)
- **citation**: `main.go:1437`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| `backend` | App | App (cms **domain** in `app`) — **and two more in the code, neither of which runs on a stock `core` stack.** (a) the standalone `messenger` service subscribed to this same stream — `messenger/internal/flow/flow.go:72` is a literal `AddSubscriber("backend", …)` over 21 live handlers — though since `838d907` deleted that compose service and its profile there is no longer any way to start it, so this subscriber now exists only in the frozen repo. (b) Since the v9.0 messenger-in-app fold (`app` `9d00a313`), **`app` itself can run a second subscriber server** on messenger's *own* Redis consumer group, attaching by messenger's own stream literals — `StreamBackend` is the literal `"backend"` (`app/internal/messenger/flow/streams.go:65`), named from `main.go:1493`. **That one is opt-in, and OFF by default** — corrected M257x iter-102; this cell previously said *"That one is not opt-in: it is the stock `core` selection"* and cited `app/main.go:1442`, which is a comment line at the newer refs and does not exist at all at `b948604f` (`main.go` is 1361 lines there). Measured: the whole block sits behind `if messengerEnabled {` at **`app/main.go:1445`**, fed by `mustSubsystemSwitch(envMessengerEnabled)` at `:285` (`env_guards.go:61` — `envMessengerEnabled = "MESSENGER_ENABLED"`), and the code states it in-line at `main.go:1437` @ `ad9f3c49`: *"ALL OF IT is behind MESSENGER_ENABLED."* Compose deliberately sets nothing there and says why where the variables would have gone (`docker-compose.yml:84-92` @ `0c91421`: *"which default to OFF on a developer machine"*), so **unset = OFF** — which is what the Messenger row above has said all along. `app` anchors here @ `app` `ad9f3c49` (`origin/main` on 2026-08-06; they hold identically at `2035f9a4`). What `d11a403` removed was only the `cms` husk container, and that was never the whole answer | User/org updates |
```

**CITED CONTENT**

```
  1434  	// messenger's do unbounded Brevo HTTP with no client timeout, so app's router reaches
  1435  	// close first essentially by construction.
  1436  	//
  1437  	// ALL OF IT is behind MESSENGER_ENABLED. Gating only the *sender* would not be
  1438  	// enough: merely constructing this server and calling Subscribe() attaches app to
  1439  	// messenger's LIVE consumer group and starts claiming entries off the real
  1440  	// streams. A developer running `go run .` against a shared Redis would take
```

## 08-037
- **id**: `B08-037`
- **corpus site**: `corpus/architecture/dependency_map.md:58-58` (table-row)
- **citation**: `docker-compose.yml:84-92`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `backend` | App | App (cms **domain** in `app`) — **and two more in the code, neither of which runs on a stock `core` stack.** (a) the standalone `messenger` service subscribed to this same stream — `messenger/internal/flow/flow.go:72` is a literal `AddSubscriber("backend", …)` over 21 live handlers — though since `838d907` deleted that compose service and its profile there is no longer any way to start it, so this subscriber now exists only in the frozen repo. (b) Since the v9.0 messenger-in-app fold (`app` `9d00a313`), **`app` itself can run a second subscriber server** on messenger's *own* Redis consumer group, attaching by messenger's own stream literals — `StreamBackend` is the literal `"backend"` (`app/internal/messenger/flow/streams.go:65`), named from `main.go:1493`. **That one is opt-in, and OFF by default** — corrected M257x iter-102; this cell previously said *"That one is not opt-in: it is the stock `core` selection"* and cited `app/main.go:1442`, which is a comment line at the newer refs and does not exist at all at `b948604f` (`main.go` is 1361 lines there). Measured: the whole block sits behind `if messengerEnabled {` at **`app/main.go:1445`**, fed by `mustSubsystemSwitch(envMessengerEnabled)` at `:285` (`env_guards.go:61` — `envMessengerEnabled = "MESSENGER_ENABLED"`), and the code states it in-line at `main.go:1437` @ `ad9f3c49`: *"ALL OF IT is behind MESSENGER_ENABLED."* Compose deliberately sets nothing there and says why where the variables would have gone (`docker-compose.yml:84-92` @ `0c91421`: *"which default to OFF on a developer machine"*), so **unset = OFF** — which is what the Messenger row above has said all along. `app` anchors here @ `app` `ad9f3c49` (`origin/main` on 2026-08-06; they hold identically at `2035f9a4`). What `d11a403` removed was only the `cms` husk container, and that was never the whole answer | User/org updates |
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

## 08-038
- **id**: `B08-038`
- **corpus site**: `corpus/architecture/dependency_map.md:59-59` (table-row)
- **citation**: `main.go:287`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| `skiller` | **none** — the producer was the standalone skiller service, decommissioned at the merge; **the fact was deleted, not moved** | App (consumer only) | `SkillerCustomJobRoleCreated` → migrate an org's members to a new custom job role. **`app` runs a live subscriber on a stream nothing publishes to.** Enumerated over every publisher constructor in `app` @ **`b948604f` only** — `backend` (`main.go:287`), `SKILLPATH_STREAM` (`:637`), `CMS_STREAM` (`:1039`), `AI_USAGE_STREAM` + `JOBSIMULATION_STREAM` (`internal/jobsimwiring/wiring.go:127`, `:180`) — `SKILLER_STREAM` is not among them; its one Go occurrence **at that ref**, `main.go:1276`, is an `AddSubscriber` (handler: `internal/roles/roles.go:791`). **One ref for the line numbers, deliberately** — `app` is 98 commits past `b948604f`, and **6 of those 7 anchors drift**: at `app` `2035f9a4`, and identically at `ad9f3c49` (`origin/main` on 2026-08-06), `main.go:287` is `logger.Info("subsystem switches",`, `:637` is `)`, `:1039` is `serverContext,`, `:1276` is `apiKeyManager,`, and `wiring.go:127` / `:180` are an asynq-client construction and a blank line. **The seventh HOLDS** — `internal/roles/roles.go:791` is `func (r *RoleManager) SkillerSubscriber() *pubsub.Subscriber {`, byte-identical at all three refs — so *"not one of those line numbers resolves"*, which this cell asserted from M257x iter-98 through iter-101, is **RETRACTED**: the split is **6 drift / 1 holds**, and a bolded universal falsified by a member of its own enumerated set is the defect, not the drift. At the newer refs `SKILLER_STREAM` has **6 Go occurrences across 3 files** — `git grep -n SKILLER_STREAM ad9f3c49 -- '*.go'` → 6 lines; `git grep -l` → `main.go`, `subscriber_merge_test.go`, `subscriber_wiring.go` — and **6 files** off the `*.go` pathspec. **No scope yields the "4 files" this cell used to claim**; that figure was induced by iter-98's own repair and is retracted with it. **The consumer-only finding itself holds at every ref**: neither `2035f9a4` nor `ad9f3c49` has a `NewPublisher` naming `SKILLER_STREAM` (there it only fills the `Skiller` field of `streamNamesForSubscribers`, `main.go:1537`), and `internal/roles/roles.go:791` is `SkillerSubscriber()` at each. Compose still sets the name (`docker-compose.yml:71` @ `0c91421`). This row's Events cell previously named skill-score changes — that was never this stream's payload
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

## 08-039
- **id**: `B08-039`
- **corpus site**: `corpus/architecture/dependency_map.md:59-59` (table-row)
- **citation**: `internal/jobsimwiring/wiring.go:127`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/jobsimwiring/wiring.go`  (283 lines)

**CLAIMING UNIT**

```md
| `skiller` | **none** — the producer was the standalone skiller service, decommissioned at the merge; **the fact was deleted, not moved** | App (consumer only) | `SkillerCustomJobRoleCreated` → migrate an org's members to a new custom job role. **`app` runs a live subscriber on a stream nothing publishes to.** Enumerated over every publisher constructor in `app` @ **`b948604f` only** — `backend` (`main.go:287`), `SKILLPATH_STREAM` (`:637`), `CMS_STREAM` (`:1039`), `AI_USAGE_STREAM` + `JOBSIMULATION_STREAM` (`internal/jobsimwiring/wiring.go:127`, `:180`) — `SKILLER_STREAM` is not among them; its one Go occurrence **at that ref**, `main.go:1276`, is an `AddSubscriber` (handler: `internal/roles/roles.go:791`). **One ref for the line numbers, deliberately** — `app` is 98 commits past `b948604f`, and **6 of those 7 anchors drift**: at `app` `2035f9a4`, and identically at `ad9f3c49` (`origin/main` on 2026-08-06), `main.go:287` is `logger.Info("subsystem switches",`, `:637` is `)`, `:1039` is `serverContext,`, `:1276` is `apiKeyManager,`, and `wiring.go:127` / `:180` are an asynq-client construction and a blank line. **The seventh HOLDS** — `internal/roles/roles.go:791` is `func (r *RoleManager) SkillerSubscriber() *pubsub.Subscriber {`, byte-identical at all three refs — so *"not one of those line numbers resolves"*, which this cell asserted from M257x iter-98 through iter-101, is **RETRACTED**: the split is **6 drift / 1 holds**, and a bolded universal falsified by a member of its own enumerated set is the defect, not the drift. At the newer refs `SKILLER_STREAM` has **6 Go occurrences across 3 files** — `git grep -n SKILLER_STREAM ad9f3c49 -- '*.go'` → 6 lines; `git grep -l` → `main.go`, `subscriber_merge_test.go`, `subscriber_wiring.go` — and **6 files** off the `*.go` pathspec. **No scope yields the "4 files" this cell used to claim**; that figure was induced by iter-98's own repair and is retracted with it. **The consumer-only finding itself holds at every ref**: neither `2035f9a4` nor `ad9f3c49` has a `NewPublisher` naming `SKILLER_STREAM` (there it only fills the `Skiller` field of `streamNamesForSubscribers`, `main.go:1537`), and `internal/roles/roles.go:791` is `SkillerSubscriber()` at each. Compose still sets the name (`docker-compose.yml:71` @ `0c91421`). This row's Events cell previously named skill-score changes — that was never this stream's payload
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

## 08-040
- **id**: `B08-040`
- **corpus site**: `corpus/architecture/dependency_map.md:59-59` (table-row)
- **citation**: `main.go:1276`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| `skiller` | **none** — the producer was the standalone skiller service, decommissioned at the merge; **the fact was deleted, not moved** | App (consumer only) | `SkillerCustomJobRoleCreated` → migrate an org's members to a new custom job role. **`app` runs a live subscriber on a stream nothing publishes to.** Enumerated over every publisher constructor in `app` @ **`b948604f` only** — `backend` (`main.go:287`), `SKILLPATH_STREAM` (`:637`), `CMS_STREAM` (`:1039`), `AI_USAGE_STREAM` + `JOBSIMULATION_STREAM` (`internal/jobsimwiring/wiring.go:127`, `:180`) — `SKILLER_STREAM` is not among them; its one Go occurrence **at that ref**, `main.go:1276`, is an `AddSubscriber` (handler: `internal/roles/roles.go:791`). **One ref for the line numbers, deliberately** — `app` is 98 commits past `b948604f`, and **6 of those 7 anchors drift**: at `app` `2035f9a4`, and identically at `ad9f3c49` (`origin/main` on 2026-08-06), `main.go:287` is `logger.Info("subsystem switches",`, `:637` is `)`, `:1039` is `serverContext,`, `:1276` is `apiKeyManager,`, and `wiring.go:127` / `:180` are an asynq-client construction and a blank line. **The seventh HOLDS** — `internal/roles/roles.go:791` is `func (r *RoleManager) SkillerSubscriber() *pubsub.Subscriber {`, byte-identical at all three refs — so *"not one of those line numbers resolves"*, which this cell asserted from M257x iter-98 through iter-101, is **RETRACTED**: the split is **6 drift / 1 holds**, and a bolded universal falsified by a member of its own enumerated set is the defect, not the drift. At the newer refs `SKILLER_STREAM` has **6 Go occurrences across 3 files** — `git grep -n SKILLER_STREAM ad9f3c49 -- '*.go'` → 6 lines; `git grep -l` → `main.go`, `subscriber_merge_test.go`, `subscriber_wiring.go` — and **6 files** off the `*.go` pathspec. **No scope yields the "4 files" this cell used to claim**; that figure was induced by iter-98's own repair and is retracted with it. **The consumer-only finding itself holds at every ref**: neither `2035f9a4` nor `ad9f3c49` has a `NewPublisher` naming `SKILLER_STREAM` (there it only fills the `Skiller` field of `streamNamesForSubscribers`, `main.go:1537`), and `internal/roles/roles.go:791` is `SkillerSubscriber()` at each. Compose still sets the name (`docker-compose.yml:71` @ `0c91421`). This row's Events cell previously named skill-score changes — that was never this stream's payload
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

## 08-041
- **id**: `B08-041`
- **corpus site**: `corpus/architecture/dependency_map.md:59-59` (table-row)
- **citation**: `internal/roles/roles.go:791`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/roles/roles.go`  (796 lines)

**CLAIMING UNIT**

```md
| `skiller` | **none** — the producer was the standalone skiller service, decommissioned at the merge; **the fact was deleted, not moved** | App (consumer only) | `SkillerCustomJobRoleCreated` → migrate an org's members to a new custom job role. **`app` runs a live subscriber on a stream nothing publishes to.** Enumerated over every publisher constructor in `app` @ **`b948604f` only** — `backend` (`main.go:287`), `SKILLPATH_STREAM` (`:637`), `CMS_STREAM` (`:1039`), `AI_USAGE_STREAM` + `JOBSIMULATION_STREAM` (`internal/jobsimwiring/wiring.go:127`, `:180`) — `SKILLER_STREAM` is not among them; its one Go occurrence **at that ref**, `main.go:1276`, is an `AddSubscriber` (handler: `internal/roles/roles.go:791`). **One ref for the line numbers, deliberately** — `app` is 98 commits past `b948604f`, and **6 of those 7 anchors drift**: at `app` `2035f9a4`, and identically at `ad9f3c49` (`origin/main` on 2026-08-06), `main.go:287` is `logger.Info("subsystem switches",`, `:637` is `)`, `:1039` is `serverContext,`, `:1276` is `apiKeyManager,`, and `wiring.go:127` / `:180` are an asynq-client construction and a blank line. **The seventh HOLDS** — `internal/roles/roles.go:791` is `func (r *RoleManager) SkillerSubscriber() *pubsub.Subscriber {`, byte-identical at all three refs — so *"not one of those line numbers resolves"*, which this cell asserted from M257x iter-98 through iter-101, is **RETRACTED**: the split is **6 drift / 1 holds**, and a bolded universal falsified by a member of its own enumerated set is the defect, not the drift. At the newer refs `SKILLER_STREAM` has **6 Go occurrences across 3 files** — `git grep -n SKILLER_STREAM ad9f3c49 -- '*.go'` → 6 lines; `git grep -l` → `main.go`, `subscriber_merge_test.go`, `subscriber_wiring.go` — and **6 files** off the `*.go` pathspec. **No scope yields the "4 files" this cell used to claim**; that figure was induced by iter-98's own repair and is retracted with it. **The consumer-only finding itself holds at every ref**: neither `2035f9a4` nor `ad9f3c49` has a `NewPublisher` naming `SKILLER_STREAM` (there it only fills the `Skiller` field of `streamNamesForSubscribers`, `main.go:1537`), and `internal/roles/roles.go:791` is `SkillerSubscriber()` at each. Compose still sets the name (`docker-compose.yml:71` @ `0c91421`). This row's Events cell previously named skill-score changes — that was never this stream's payload
```

**CITED CONTENT**

```
   788  	)
   789  }
   790  
   791  func (r *RoleManager) SkillerSubscriber() *pubsub.Subscriber {
   792  	return pubsub.NewSubscriber().AddHandler(
   793  		pubsub.EventHandler(r.SkillerCustomJobRoleCreatedHandler),
   794  	)
```

## 08-042
- **id**: `B08-042`
- **corpus site**: `corpus/architecture/dependency_map.md:59-59` (table-row)
- **citation**: `main.go:1537`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
| `skiller` | **none** — the producer was the standalone skiller service, decommissioned at the merge; **the fact was deleted, not moved** | App (consumer only) | `SkillerCustomJobRoleCreated` → migrate an org's members to a new custom job role. **`app` runs a live subscriber on a stream nothing publishes to.** Enumerated over every publisher constructor in `app` @ **`b948604f` only** — `backend` (`main.go:287`), `SKILLPATH_STREAM` (`:637`), `CMS_STREAM` (`:1039`), `AI_USAGE_STREAM` + `JOBSIMULATION_STREAM` (`internal/jobsimwiring/wiring.go:127`, `:180`) — `SKILLER_STREAM` is not among them; its one Go occurrence **at that ref**, `main.go:1276`, is an `AddSubscriber` (handler: `internal/roles/roles.go:791`). **One ref for the line numbers, deliberately** — `app` is 98 commits past `b948604f`, and **6 of those 7 anchors drift**: at `app` `2035f9a4`, and identically at `ad9f3c49` (`origin/main` on 2026-08-06), `main.go:287` is `logger.Info("subsystem switches",`, `:637` is `)`, `:1039` is `serverContext,`, `:1276` is `apiKeyManager,`, and `wiring.go:127` / `:180` are an asynq-client construction and a blank line. **The seventh HOLDS** — `internal/roles/roles.go:791` is `func (r *RoleManager) SkillerSubscriber() *pubsub.Subscriber {`, byte-identical at all three refs — so *"not one of those line numbers resolves"*, which this cell asserted from M257x iter-98 through iter-101, is **RETRACTED**: the split is **6 drift / 1 holds**, and a bolded universal falsified by a member of its own enumerated set is the defect, not the drift. At the newer refs `SKILLER_STREAM` has **6 Go occurrences across 3 files** — `git grep -n SKILLER_STREAM ad9f3c49 -- '*.go'` → 6 lines; `git grep -l` → `main.go`, `subscriber_merge_test.go`, `subscriber_wiring.go` — and **6 files** off the `*.go` pathspec. **No scope yields the "4 files" this cell used to claim**; that figure was induced by iter-98's own repair and is retracted with it. **The consumer-only finding itself holds at every ref**: neither `2035f9a4` nor `ad9f3c49` has a `NewPublisher` naming `SKILLER_STREAM` (there it only fills the `Skiller` field of `streamNamesForSubscribers`, `main.go:1537`), and `internal/roles/roles.go:791` is `SkillerSubscriber()` at each. Compose still sets the name (`docker-compose.yml:71` @ `0c91421`). This row's Events cell previously named skill-score changes — that was never this stream's payload
```

**CITED CONTENT**

```
  1534  	streamNamesForSubscribers := streamNames{
  1535  		Backend:       serviceName,
  1536  		SkillPath:     os.Getenv("SKILLPATH_STREAM"),
  1537  		Skiller:       os.Getenv("SKILLER_STREAM"),
  1538  		JobSimulation: os.Getenv("JOBSIMULATION_STREAM"),
  1539  		CMS:           os.Getenv("CMS_STREAM"),
  1540  		AIUsage:       os.Getenv("AI_USAGE_STREAM"),
```

## 08-043
- **id**: `B08-043`
- **corpus site**: `corpus/architecture/dependency_map.md:59-59` (table-row)
- **citation**: `docker-compose.yml:71`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
| `skiller` | **none** — the producer was the standalone skiller service, decommissioned at the merge; **the fact was deleted, not moved** | App (consumer only) | `SkillerCustomJobRoleCreated` → migrate an org's members to a new custom job role. **`app` runs a live subscriber on a stream nothing publishes to.** Enumerated over every publisher constructor in `app` @ **`b948604f` only** — `backend` (`main.go:287`), `SKILLPATH_STREAM` (`:637`), `CMS_STREAM` (`:1039`), `AI_USAGE_STREAM` + `JOBSIMULATION_STREAM` (`internal/jobsimwiring/wiring.go:127`, `:180`) — `SKILLER_STREAM` is not among them; its one Go occurrence **at that ref**, `main.go:1276`, is an `AddSubscriber` (handler: `internal/roles/roles.go:791`). **One ref for the line numbers, deliberately** — `app` is 98 commits past `b948604f`, and **6 of those 7 anchors drift**: at `app` `2035f9a4`, and identically at `ad9f3c49` (`origin/main` on 2026-08-06), `main.go:287` is `logger.Info("subsystem switches",`, `:637` is `)`, `:1039` is `serverContext,`, `:1276` is `apiKeyManager,`, and `wiring.go:127` / `:180` are an asynq-client construction and a blank line. **The seventh HOLDS** — `internal/roles/roles.go:791` is `func (r *RoleManager) SkillerSubscriber() *pubsub.Subscriber {`, byte-identical at all three refs — so *"not one of those line numbers resolves"*, which this cell asserted from M257x iter-98 through iter-101, is **RETRACTED**: the split is **6 drift / 1 holds**, and a bolded universal falsified by a member of its own enumerated set is the defect, not the drift. At the newer refs `SKILLER_STREAM` has **6 Go occurrences across 3 files** — `git grep -n SKILLER_STREAM ad9f3c49 -- '*.go'` → 6 lines; `git grep -l` → `main.go`, `subscriber_merge_test.go`, `subscriber_wiring.go` — and **6 files** off the `*.go` pathspec. **No scope yields the "4 files" this cell used to claim**; that figure was induced by iter-98's own repair and is retracted with it. **The consumer-only finding itself holds at every ref**: neither `2035f9a4` nor `ad9f3c49` has a `NewPublisher` naming `SKILLER_STREAM` (there it only fills the `Skiller` field of `streamNamesForSubscribers`, `main.go:1537`), and `internal/roles/roles.go:791` is `SkillerSubscriber()` at each. Compose still sets the name (`docker-compose.yml:71` @ `0c91421`). This row's Events cell previously named skill-score changes — that was never this stream's payload
```

**CITED CONTENT**

```
    68        - REDIS_WORKER_INDEX=0
    69        - RPC_PORT=8083
    70        - SERVICE_NAME=backend
    71        - SKILLER_STREAM=skiller
    72        - SKILLPATH_STREAM=skillpath
    73        # storage-in-app (v9.0): app owns the buckets in-process; there is no storage service to
    74        # address. AWS credentials come from .env — the same ones the standalone storage
```

## 08-044
- **id**: `B08-044`
- **corpus site**: `corpus/architecture/dependency_map.md:60-60` (table-row)
- **citation**: `messenger/internal/flow/flow.go:105`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/messenger/internal/flow/flow.go`  (135 lines)

**CLAIMING UNIT**

```md
| `jobsimulation` | App | App (the jobsim engine + the skill-path engine, on ONE subscriber), the standalone Messenger (`messenger/internal/flow/flow.go:105` — no longer startable since `838d907`), **and app's messenger-in-app subscriber**: the same takeover as the `backend` row attaches to `StreamJobSimulation` too — and, like it, is **opt-in behind `MESSENGER_ENABLED` and OFF by default** | Session completed, insights generated |
```

**CITED CONTENT**

```
   102  		pubsub.EventHandler(h.CourseBuildFailedHandler),
   103  		pubsub.EventHandler(h.CoursePublishedHandler),
   104  	))
   105  	h.subServer.AddSubscriber("jobsimulation", sub.AddHandler(
   106  		pubsub.EventHandler(h.JobsimulationSessionStartedHandler),
   107  		pubsub.EventHandler(h.JobsimulationSessionEndedHandler),
   108  	))
```
