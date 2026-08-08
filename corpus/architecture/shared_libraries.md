# Shared Libraries

> ## ⚠️ This document's subject set is NOT `app`'s require set — corrected M257x iter-123
>
> The "five shared libraries" is a **historical grouping**, and every count below is taken over it.
> `app`'s actual org-private module requirements, measured at `app` **`3eaadae6`** (v1.371.1),
> `app/go.mod:14-18` — **all five DIRECT, zero `// indirect`, and the file's one `replace` (`:295`) is
> `getsentry/sentry-go/echo`, not an org module**:
>
> | Module | `go.mod` line | Version | In the "five"? |
> |---|---|---|---|
> | `analytics-go` | `:14` | `v0.3.1` | **no — this doc had no section for it** |
> | `colony` | `:15` | `v0.35.2` | yes |
> | `proto` | `:16` | `v1.210.0` | yes |
> | **`storage`** | `:17` | `v0.15.2` | **no — this doc had no section for it** |
> | `taxonomy` | `:18` | `v1.2.0` | yes |
>
> `go.sum:64-73` carries exactly those five, two lines each — **no `ai`, no `authn`, no `messenger`.**
>
> **The one sentence to carry away: "merged into `app`" describes the RUNTIME, not the module graph.**
> `ai` left `go.mod`; **`storage` did not.** Read the two as one rule and you conclude `storage` is
> deletable, which is false twice over — it is a compile-time requirement of the backend monolith *and*
> the repo is still maintained (HEAD **2026-08-05**, tags out to **`v0.15.8`**, six past `app`'s pin).
>
> **`analytics-go` was the sharpest gap**: a live private Go dependency carrying Stripe
> subscription-lifecycle events, with **no** section here and **no** guard anywhere. It is not, however,
> undocumented corpus-wide — [`external_services.md:554`](./external_services.md) already carried the
> correct `go.mod:14-18` enumeration. **The defect was that the library model and the external-services
> doc disagreed and nothing reconciled them**, which is the same failure mode as two corpora
> disagreeing, one file apart. Both new modules are described in `CLAUDE.md`'s Shared Libraries block;
> full sections here are owed and **not yet written** — that is a named gap, not a silent one.
>
> **Do not delete `ai` either.** `rosetta-extensions/stack-seeding` pins `ai v1.40.1`, so the deletion
> this doc's old text invites would break Rosetta's own tooling
> (`app/internal/ai/module_import_guard_test.go:15-17`; `app/CLAUDE.md:289-294`).
>
> **The block is actively shrinking**, so date every reading of it: at `b948604f` it was **seven**
> modules — `ai` at `:14` and `messenger` at `:17` — and at `ad9f3c49` and `3eaadae6` it is the five above.

This document covers **five** internal library repos. **The Go services do not share five**
(corrected M257x iter-102; this line previously said *"The Anthropos Go services share five
internal libraries"*): measured at platform `0c91421`, the two Go repos a stack clones and
builds — `app` and `sentinel` — require **three OF THE FIVE**, colony + proto + taxonomy (`app`
requires **five modules in total**; the other two are outside this doc's subject set — see the
banner). The other two of the five
arrived by **absorption, not dependency**: `authn` ships *inside* colony as `colony/authn`
and is a `require` in **no** repo's `go.mod`, and `ai` was folded into `app` as
`app/internal/ai` at `1e457fa70` (2026-08-04), which dropped its module requirement — it
survives as a requirement only in the frozen `cms` / `jobsimulation` repos. None of the five
is a **deployed service** — there is no container, port, or `docker-compose` entry for any of
them. They are **Go modules** compiled *into* each service's binary.

## High-Level Summary (For PMs & Non-Engineers)

Think of these as the platform's "standard library." Rather than every microservice
re-implementing logging, database wiring, authentication, RPC contracts, or AI calls,
that shared plumbing lives in a handful of small repos that the services pull in like any
third-party dependency. This keeps the services consistent and small. **Five such repos
exist, but no service pulls five** (corrected M257x iter-102; this passage previously said
*"that shared plumbing lives in five small repos that the services pull in like any
third-party dependency"*). Counting `go.mod` requires over the seven Go repos on disk at
their pinned refs: colony **7/7**, proto **7/7**, taxonomy **6/7** (all but roadrunner),
ai **2/7** (only the frozen cms + jobsimulation), authn **0/7**. So **four** of the five are
pulled by at least one repo, and only **three** by the two repos a stack actually builds.
**Those ratios are over THE FIVE, and the denominator hides two live modules** — `app` also requires
`storage` and `analytics-go` (banner above), so *"three"* is the answer to *"how many of the five"*,
never to *"what does `app` depend on"*.

| Library | One-liner |
|---------|-----------|
| **colony** | The framework: DB/Redis, logging+Sentry, GraphQL/RPC servers, CORS, pub/sub, feature flags — and it now contains **authn** |
| **proto** | The single source of truth for service-to-service RPC contracts (Protobuf → Connect-RPC) + hand-written domain types |
| **ai** | A thin wrapper over OpenAI/Azure/Anthropic/Bedrock/Mistral behind one `ai.AI` interface. **Folded into `app` at `1e457fa70`; in no live `go.mod`** |
| **authn** | Clerk JWT authentication (now shipped **inside colony** as `colony/authn`; the standalone repo is legacy) |
| **taxonomy** | The **node-id library** (`NodeID` type + ID generation/validation) — **not** a dataset |
| **storage** *(not one of "the five" — added M257x iter-123)* | `app/go.mod:17` `v0.15.2`. A **type shim**: `app` imports `sdk/storage` + `sdk/storage/v1` (32 files / 36 import lines) and implements the three-method `Service` interface **itself** (`app/internal/storage/service.go:48-56` fills `sdkstorage.Client{V1: NewService(...)}`). **No SDK RPC client is constructed; `STORAGE_RPC_ADDR` occurs in zero Go source.** The service is gone; the module is not |
| **analytics-go** *(not one of "the five" — added M257x iter-123)* | `app/go.mod:14` `v0.3.1`. Two files (`analytics.go` + `brevo/brevo.go`): a `sync.Once` `Init`/`Track` fan-out to Brevo's `EventsApi.CreateEvent`. Wrapped by `app/internal/tracking`; the load-bearing consumer is **Stripe subscription lifecycle → Brevo**, `app/internal/payments/handler.go:302-316` (seven event names switched off `entSub.Status`), wired at **`main.go:494-495`** @ `app` `ad9f3c498` — `trackingManager := tracking.New(os.Getenv("BREVO_KEY"))` and the `payments.New(…, trackingManager)` that consumes it, which are the file's **only two** `trackingManager` lines. ⚠️ **This carried a `main.go` line range until M257x iter-138, by then the storage-in-app comment block**, and it named a *different* construct at all four refs the corpus reads (`b948604f`: a jobsim `BACKEND_USERS_RPC_ADDR` comment; `9d00a313`: an AI-Readiness auto-assign comment; `2035f9a4` + `ad9f3c498`: the storage comment). The `handler.go:302-316` half was **verified and is exact**. Repo untouched since **2025-02-12**; `v0.3.1` is its newest tag — **dormant, not dead** |

> ### How they are consumed (this matters)
> **None of these are cloned by `make init`** — they are **absent from
> `platform/repos.yml`**, so there is no `stack-dev/colony` (etc.) directory.
> **Each of the three a stack actually builds against** — colony, proto, taxonomy — is pulled
> as a **private Go module** during a service's Docker build (`authn` rides *inside* colony
> and is fetched by no `go.mod` line of its own; `ai` has no live puller left since `app`
> folded it in — see its section below):
> `platform/docker-compose.yml` passes `GH_ACCESS_TOKEN=$GH_PAT` as a build arg, and the
> service Dockerfiles set `GOPRIVATE=github.com/anthropos-work/*` plus a
> `git config … url."https://x-access-token:${GH_ACCESS_TOKEN}@github.com/".insteadOf`
> rewrite so `go mod download` can fetch them. **Without a valid `GH_PAT`, the build fails.**
> To work on one locally you clone it and add an (uncommitted) `go.work` with
> `use ( . ../<lib> )` — no `go.work` is committed in any service.

---

## colony

| Property | Value |
|:---------|:------|
| **Module** | `github.com/anthropos-work/colony` |
| **Language** | Go (`go.mod` declares `go 1.25.0`; built with `golang:1.26-bookworm`) |
| **Version pin** | **ONE pin across the live services: `v0.35.2`.** `app` (`app/go.mod:15` @ `ad9f3c49`) and `sentinel` (`sentinel/go.mod:8` @ `f2c46190`) — the only two `type: go` entries in `repos.yml` — now agree. **The long-standing "split" is CLOSED**, and the closing event is dated: `sentinel`'s `88036d7` *"chore(deps): update dependencies to latest versions"* took it `v0.34.3 → v0.35.2`, two commits past the `88bc5592` this row used to cite. The frozen repos keep their own pins and are **not** part of this reading — `storage` `v0.34.3` (`4ce8ece5`), `messenger` `v0.35.2` (`fa47850d`), archived `chronos` `v0.30.1`; the `v0.35.1` third pin went with the `cms` + `jobsimulation` husk containers at `d11a403`. Measured from each repo's `go.mod` at the ref stated beside it. |
| **Imported by** | **Every live Go service — `app` and `sentinel`, and only those two, at platform `0c91421`.** `repos.yml` there lists four entries (app, sentinel, next-web-app, studio-desk) of which two are `type: go`. The four-service reading (app, sentinel, **storage**, **messenger**) was true at `0dab54d` and **`838d907` ended it**: that commit deleted the `storage` and `messenger` clone entries *and* their compose services, so `make init` now clones app, sentinel, next-web-app, studio-desk. Both repos still import colony in their own `go.mod` (storage `v0.34.3` @ `4ce8ece5`, messenger `v0.35.2` @ `fa47850d`) but nothing clones or builds them — they join the `cms`, `jobsimulation` and `roadrunner` repos, **gone from compose** at `0dab54d` (`d11a403`); there is no profile that starts them and no `graphql` profile at all. Their domains run inside `app`, and the three repos are **frozen legacy** — still on GitHub as the pre-merge reference, but with no compose service, no `repos.yml` entry and nothing that starts them, so they are not importers of anything a stack runs. The `roadrunner` repo's own import was minimal but real while it lasted: `roadrunner/main.go:7` imports `colony` for `NewVersionConfig` (`roadrunner/go.mod:7` pins `v0.34.3`) |

The platform framework. Each service composes its server out of colony packages:

| Package | What it provides |
|---------|------------------|
| `colony` (root) | `NewDBStdConn`/`NewDBPool` (pgx v5 Postgres for Ent + raw pools, with a 30s **DB health monitor** that cancels the service context after 3 failed pings), `InitLogger` (slog + Sentry fan-out), `NewGQLHandlerServer` + `Apply*GraphqlMiddlewares`, `NewHTTPServer` (serves `GET /_meta` version/health JSON), `NewCORSHandler`, `Environment`, `NewVersionConfig` (resolves ECS task id) |
| `colony/authn` | Clerk JWT auth (see **authn** below) — the live copy of authn |
| `colony/authorization` | ⚠️ Go **package name is `authorizer`** — `NewSentinelAuthorizer` (Connect-RPC client to Sentinel's `AuthorizationService`), `Authorizer` interface, `Decision` type, ctx helpers |
| `colony/pubsub` | Watermill over **Redis Streams**: `NewPublisher`/`NewSubscriberServer`, generic `EventHandler[T]`, poison-queue DLQ + 3× exponential-backoff retry, proto `eventsv1.Event` envelope |
| `colony/redis` | go-redis/v9 `UniversalClient` factory (pings on startup) |
| `colony/rpc` | h2c (cleartext HTTP/2) Connect-RPC server/client + `DefaultInterceptors` (request logging + proto `Validate()`) |
| `colony/flags` | Feature flags: `PosthogChecker` (PostHog EU, 5-min polling) and `EnvChecker` |

**Notable**: Sentry only initializes when `ENVIRONMENT=production` *or* `FORCE_SENTRY`
is set (dev = plain slog). The GraphQL public-middleware per-IP rate limiter
(`rate.NewLimiter(1,5)`) currently **only logs** — its reject path is commented out, so
colony does **not** actually enforce GraphQL rate limiting today.

---

## proto

| Property | Value |
|:---------|:------|
| **Module** | `github.com/anthropos-work/proto` |
| **Language** | Go (`go 1.25.0`); tooling: `buf` (CI pins `v1.57.0`), protoc-gen-go, protoc-gen-connect-go, goverter |
| **Version pin** | **ONE value across the live services: `v1.210.0`.** `app` (`app/go.mod:16` @ `ad9f3c49`) and `sentinel` (`sentinel/go.mod:9` @ `f2c46190`) agree, so **the live skew is ZERO, not two** — closed by the same `sentinel` `88036d7` dependency bump that closed the colony split (`v1.200.0 → v1.210.0`). At platform `0dab54d` the reading was four repos (app/messenger `v1.210.0`, sentinel `v1.200.0`, storage `v1.196.0`); `838d907` deleted the `storage` and `messenger` clone entries, so `storage v1.196.0` (`4ce8ece5`) and `messenger v1.210.0` (`fa47850d`) are frozen alongside the husks below and are not part of the live reading. The frozen repos keep their own `go.mod` and therefore their own pins (cms `v1.207.0` @ `ca50c817`, jobsimulation `v1.205.0` @ `462343b0`, roadrunner `v1.196.0` @ `87d8d44`), but **nothing clones or builds them any more**: `d11a403` deleted all three `repos.yml` entries *and* their compose services in one commit, so those pins compile nowhere. Reading them as part of the platform's skew was the error this row used to make |
| **Imported by** | every live Go service that does RPC — **`app` and `sentinel` at platform `0c91421`** (it was app, sentinel, storage, messenger at `0dab54d`; `838d907` dropped the last two from `repos.yml`). The cms / jobsimulation / **skiller** RPC surfaces are served in-process by `app`; **skillpath and roadrunner were REMOVED, not re-hosted** — `app/main.go` registers six Connect handlers @ `app` `b948604` v1.366.0 — five unconditionally (Users `app/main.go:1187`, Organizations `app/main.go:1188`, Skiller `app/main.go:1196`, JobSimulation `app/main.go:1204`, LabSession `app/main.go:1228`) plus `CMSService` **only when a cms RPC server was built** (`app/main.go:1212-1214`, `if cmsRPCServer != nil`) — and neither `SkillPathSessionService` nor a RoadRunner service is among them |

The **single source of truth for RPC contracts**. Two layers:

* **Generated** — `proto/<svc>/v1/*.proto` → `go/<svc>/v1/` (message structs) + `go/<svc>/v1/<svc>v1connect/` (Connect-RPC stubs, *do not edit*).
* **Hand-written** — `go/domain/<svc>/` idiomatic Go types (string enums, `time.Time`) plus **goverter**-generated converters. goverter fails codegen if a proto enum value has no matching domain const — the "three-file rule" (proto + domain const + `make gen`).

**At least 13** Connect-RPC services are defined — **this is a floor, not a count**: `proto` is a private Go
module and **is in no clone set**, so the list below is hand-enumerated from consumers and cannot be
verified against the source of truth. It omitted `StorageService` until M257x iter-98, which
[`storage.md:129`](../services/storage.md) documents in full. The named ones:
`UsersService`, `OrganizationsService`,
`CMSService`, `JobSimulationService`, `SkillerService` (all served by app since the merges — one RPC mux),
`SkillPathSessionService` (**contract still in `proto`, but NO LONGER SERVED** — like `ChronosService`. skillpath-in-app M506 *removed* the RPC rather than re-hosting it; `app/internal/skillpaths/skillpaths.go:27-31` calls its replacement "the drop-in for the **removed** skillpath RPC client". Likewise roadrunner: `backend` calls Judge0 over plain HTTP — `jsrunner.NewRunnerManager` at `app/internal/jobsimwiring/wiring.go:123` @ `app` `9d00a313` v1.367.0 — and `ROADRUNNER_RPC_ADDR` is read by no Go code in `app` **and is not in the platform compose at all** — 0 occurrences in either, at `app` `9d00a313` / platform `0dab54d`. (This line long cited `docker-compose.yml:118` for it; that line sets `AWS_REGION`.)),
`LabSessionService` (served by app — the AI Labs domain, see `../services/ai-labs.md`),
`AuthorizationService` (Sentinel), `MessengerService`, `RoadRunnerService`,
`RealtimeService`, `ChronosService` (archived service, contract still present), and
`StorageService` (the storage surface — served in-process by `app` since the v9.0 fold). Plus
`events`/`flags`/`ai` message-only protos used over Redis Streams pub/sub.

```bash
make gen          # buf format → build → breaking → generate → go generate (goverter)
make force-gen    # same, skipping the breaking-change check (dev)
# consumers bump with: GOPRIVATE=github.com/anthropos-work/* go get -u github.com/anthropos-work/proto@latest
```

**Notable**: legacy **buf v1** single-module layout (`proto/buf.yaml`, name
`buf.build/anthropos/platform`) — **no `buf.work.yaml`**. `go/simulator/*` holds ~10
generated packages with **no source `.proto`** in the repo (legacy/vendored stubs,
e.g. `storage/internal/migration` imports `go/simulator/storage/v1` as `legacyStorage`).

---

## ai

| Property | Value |
|:---------|:------|
| **Module** | `github.com/anthropos-work/ai` |
| **Language** | Go (`go 1.25.0`) |
| **Version pin** | **`v1.40.2`** across every repo that *still* requires it — the frozen `cms` and `jobsimulation` — both agree. `app` required `v1.40.2` too, up to `b948604f`; it requires nothing now (next row) |
| **Imported by** | **No repo a stack builds** (corrected M257x iter-102; this row previously said *"`app` alone among the services a stack runs"*). `app` **dropped** the module at `1e457fa70` (2026-08-04, *"refactor(ai): fold the ai library into app as internal/ai"*): `git show ad9f3c49:go.mod` has no `anthropos-work/ai` line and `go.sum` has zero, while `app/internal/ai/` carries the library in-tree — with a one-way door, `internal/ai/module_import_guard_test.go`, whose own comment records that the repo *"was deliberately left in place because at least one consumer outside this codebase (anthropos-work/rosetta-extensions/stack-seeding) pins it."* `sentinel` never required it. The frozen `cms` and `jobsimulation` repos still require it directly (`cms/go.mod:9` @ `ca50c817`, `jobsimulation/go.mod:11` @ `462343b0`, both `v1.40.2`), but neither has a compose service or a `repos.yml` entry, so nothing builds or starts them — a `go.mod` require in a repo nothing compiles is not a live import. Go services only — **not** Studio-Desk, which is TypeScript, and **not** roadrunner, whose only shared-lib requires are colony + proto |

A thin wrapper exposing **one interface, `ai.AI`** (`ChatCompletion`,
`ChatCompletionStream`, `Response`, `CreateEmbeddings`, `CreateSpeech`, `OCRProcess`,
`AudioTranscriptions`, `Tokenize`, `GetEndpoint`) over per-provider constructors:

| Constructor | Provider |
|-------------|----------|
| `openai.New` / `NewOpenAI` | OpenAI direct |
| `openai.NewAzure` | Azure OpenAI (default API version `2025-04-01-preview`) |
| `anthropic.NewAnthropic(cfg, key)` | **AWS Bedrock** (`cfg!=nil`, EU `eu.anthropic.*` model IDs) **or** Anthropic-direct (`key!=nil`) — one constructor for both |
| `mistral.NewMistral` | Mistral — **OCR only** (chat/embeddings/speech `panic`) |

> ### ⚠️ Two corrections to long-standing corpus wording
> 1. **The `ai` library does NOT track cost.** `MetaData.Usage` only carries provider
>    token counts. Dollar cost is computed by the **consumer** in
>    `app/internal/aiusage/ai_usage.go` (a hardcoded model→price switch) and written to
>    the `ai_usage` Postgres table, fed by an `Event_AiUsage` published over Redis Streams.
> 2. **The `ai` library does NOT select a provider.** It only exposes per-provider
>    constructors. (And what the consumers do is not an EU-first fallback *ladder* either —
>    `external_services.md:579` retracts that chain.) **Vendor selection lives in each
>    consuming DOMAIN's own `ai` wrapper — NOT in a file called `internal/ai/ai.go`**
>    (corrected M257x iter-102; this passage previously said it in *"each consumer's own
>    `internal/ai/ai.go` wrapper"*, which names no such code at any ref: `app/internal/ai/`
>    did not exist at all at `b948604f`, and since the fold `app/internal/ai/ai.go` is
>    **21 lines** declaring `type AI interface` + `type TokenEncoder interface` — no Azure
>    client, no PostHog flag, no 429 handling). Measured @ `app` `ad9f3c49` — the two real
>    sites are `internal/jobsimulation/ai/ai.go` and `internal/skillerai/ai.go`, each with
>    its own `AIManager.getClient` (`:259` and `:332`): an EU Azure client by default, a US
>    Azure client swapped in when the PostHog flag `flag_use_azure_us` is enabled
>    (`jobsimulation/ai/ai.go:267` and `:344`, `skillerai/ai.go:347`), and — a **retry
>    target, not a rung** — the vendor overridden to `Openai` on the next attempt once
>    `isThrottlingError` sees an HTTP 429 (`jobsimulation/ai/ai.go:129`, used at `:166`
>    and `:325`; `skillerai/ai.go:128`, used at `:176`). Anthropic is always Bedrock in
>    `eu-west-1`.

**Other gotchas**: capability is asymmetric — only OpenAI/Azure implement embeddings,
speech, OCR, transcription, streaming; **Anthropic `ChatCompletionStream`/`CreateSpeech`
`panic`**. Anthropic has no native JSON mode, so the lib prefills `{"` and prepends it to
the response (parse accordingly). Retry policy: 10 attempts, exponential backoff, never on
401/403/404 or context cancellation. A separate, non-`ai`-library Bedrock path exists in
`app/internal/askengine/bedrock.go` (raw `anthropic-sdk-go`, prompt caching, agentic tool loop).

---

## authn

| Property | Value |
|:---------|:------|
| **Module (standalone)** | `github.com/anthropos-work/authn` — **legacy** (tag `v1.7.0`) |
| **Live form** | `github.com/anthropos-work/colony/authn` (absorbed into colony) |
| **Imported by** | via colony: app — the only service a stack runs that reaches it. **The frozen `cms` and `jobsimulation` repos still import `colony/authn` directly** (6 and 8 `.go` files at `cms ca50c817` / `jobsimulation 462343b0`), but neither has had a compose service or a `repos.yml` entry since `d11a403`, so those imports compile in no image a stack builds. Only the skillpath usage is fully folded in |

Provider-agnostic authentication: verifies bearer tokens (Clerk JWTs in practice) and
injects a typed `User`/`Organization` into request context for `net/http`, Echo, and
GraphQL servers.

> This is the Clerk JWT library; for the platform-wide picture (dependent repos, SDKs, the auth-vs-authz split) see [Clerk Integration](../services/clerk-integration.md).

* `authn.NewManager(providers…)` tries each provider in order; only **Clerk** and a
  **Dummy** (test) provider exist.
* Clerk flow (`provider/clerk`): `jwt.Verify` against Clerk JWKS, then `jwt.Decode` to
  read custom session claims (`eid`, `email`, `firstname`, `lastname`, `org`, `org_id`,
  `org_role`) — a performance optimization to avoid Clerk API round-trips. `User.ID()`
  returns the **internal Anthropos UUID** (`eid`); `AuthID()` returns the Clerk subject —
  the two-identity bridge.

> ### ⚠️ Correction: authn is effectively a colony sub-package now
> **No checked-out service imports the standalone `github.com/anthropos-work/authn`** —
> they all import `github.com/anthropos-work/colony/authn`. The standalone repo is
> legacy/orphaned (and its `HTTPAuthnMiddleware` has a missing `return` after the
> websocket-skip that colony's copy fixed). Document authn as **part of colony**, not as
> an independent dependency.
>
> **Relationship to Sentinel is loose**: Sentinel does **not** import authn. authn only
> *authenticates* (who you are); `app` then maps the resulting User/Org IDs into Sentinel
> Connect-RPC *authorization* calls (`gqlauthz.go`).

---

## taxonomy

| Property | Value |
|:---------|:------|
| **Module** | `github.com/anthropos-work/taxonomy` (README title: **"nodeid"**) |
| **Language** | Go (`go 1.21.0`), **zero external dependencies** (stdlib only) |
| **Version pin** | `v1.2.0` |
| **Imported by** | **Both Go repos a stack still clones and builds at platform `0c91421`** — `app` directly (`app/go.mod:20` @ `b948604f`, `v1.2.0`) and `sentinel` indirectly (`sentinel/go.mod:21` @ `88bc5592`, `v1.2.0 // indirect`). It was **four** at `0dab54d` — directly: app, messenger; indirectly (`// indirect`): storage, sentinel — until `838d907` deleted the `storage` and `messenger` clone entries (their `v1.2.0` requirements, `storage/go.mod:25 // indirect` @ `4ce8ece5` and `messenger/go.mod:9` direct @ `fa47850d`, are now frozen). The frozen `cms` and `jobsimulation` repos also require it directly in their own `go.mod` (`cms/go.mod:13`, `jobsimulation/go.mod:15`, both `v1.2.0`, neither marked `// indirect`) — but `d11a403` deleted their compose services and their `repos.yml` entries, so nothing builds them; they are frozen legacy, not running containers. Counted over the **seven** Go repos a stack still has on disk (app, sentinel, storage, messenger + the frozen cms, jobsimulation, roadrunner), that is **6 of 7** — the sole exception is `roadrunner`, which requires only colony + proto. **Do not read that 7 as "every Go repo the platform has ever cloned": that set is 11.** The union of `type: go` entries across `repos.yml`'s whole history also names `skiller`, `skillpath`, `chronos` and `intelligence` (all four present in the first revision, `a2a3ee6`; the last of them, skillpath, dropped by `a4db680`). None of the four is cloned by any stack today, so none of their `go.mod`s is measured here — 6 of 7 is a count over what is on disk, not over what has ever existed. (The skillpath usage is folded into app.) |

> ### ⚠️ Major correction: taxonomy is a LIBRARY, not data
> Multiple corpus docs called this "Skills taxonomy data (60K skills, 18K roles)". That
> is **wrong twice over** — wrong about *where the data lives*, and wrong about *the
> numbers themselves*.
>
> **Where it lives.** The repo is a **131-line** node-id library (`node.go`) and ships **no
> dataset**. The taxonomy *data* is owned by and stored in **app**'s
> `public` Postgres schema (the merged skiller domain — formerly the standalone skiller
> service's schema), originally loaded from **external** CSV/JSON by the former skiller
> importers (`importSkills` / `importJobRole`); taxonomy CLIs now live under `app/cmd/`
> (e.g. `createTaxonomy`). The taxonomy module
> only supplies the **ID type/format** used as keys.

<a id="taxonomy-figures"></a>
> ### ⚠️ Second correction: the "60K skills / 18K roles" figures
> They are not measurements, and they fail in two *different* ways. Keep them apart.
>
> | Long-quoted figure | Verdict | What was actually measured |
> |:-------------------|:--------|:---------------------------|
> | **"18K roles"** | **REFUTED** | **22,470** public job roles. Public ⊆ total, so prod holds **≥ 22,470** — 18K is below the floor. The 18K almost certainly came from `job_role_embeddings` (**18,919** rows), a different table, mis-transcribed onto the role count |
> | **"60K skills"** | **UNVERIFIED** (not refuted) | **42,790** public skills. A public-only capture cannot see org-scoped *private* skills, so the total could be higher — possibly much higher. Nothing measured supports 60K, and nothing measured rules it out |
>
> **Provenance.** Read-only production capture of the **public subset only**
> (`organization_id IS NULL`): `.agentspace/snapshots/taxonomy/<digest>/manifest.json`,
> `source: primary-read`, `public_only: true`, `predicate: org-null`, captured
> **2026-06-29**. Both counts reproduce exactly against a live stack database
> (`select count(*) … where organization_id is null`).
>
> **So how should this be written?** Say *"≥22,470 job roles and ≥42,790 skills (the public
> subset, measured 2026-06-29; totals including org-private content are unmeasured)"*. Do
> **not** write "42,790 skills" as though it were the total — it is a floor, not a count.
>
> ### ⚠️ A SECOND CORPUS ASSERTS THE REFUTED FIGURE, AND IT IS CUSTOMER-FACING (M257x iter-125)
>
> This is stated here, rather than only in a repo census, because **this is where a reader meets the
> figures** — and because the contradiction is live, not historical.
>
> `anthropos-work/anthropos-knowledge-base` (AKB) carries a **second, parallel platform-architecture
> corpus** — six files under `knowledge/`, ≈1,773 lines — and asserts *"60,000 skills … mapped to
> 18,000 roles"* in **14 places, citing no source in any of them**. The figure is **load-bearing in
> four customer-facing competitor-comparison tables**. AKB ships as a **Claude Code plugin** that
> injects *"full Anthropos context (product details, architecture, …)"* into every engineer's editor,
> and **this corpus's own [`toolchain_overview.md`](../tools/toolchain_overview.md) recommends
> installing it** — so a reader following our advice is served the figure we refute. That
> recommendation now carries this warning at the install line.
>
> | | this corpus | AKB |
> |---|---|---|
> | job roles | **≥ 22,470** — public subset, read-only prod capture (`organization_id IS NULL`), 2026-06-29, reproducible against a live stack DB | **18,000** — no source given, in any of 14 occurrences |
> | skills | **≥ 42,790** — same capture, same date, same predicate | **60,000** — likewise unsourced |
>
> **Which is measured is not a matter of opinion: ours is, and its provenance is above.** "18K roles"
> is **REFUTED** (public ⊆ total, so the true count is at or above 22,470); "60K skills" is
> **UNVERIFIED, not refuted** — a public-only capture cannot see org-private skills, so nothing
> measured supports 60K and nothing rules it out. **Keep those two verdicts apart**; collapsing them
> into "AKB is wrong" over-claims on the skills row.
>
> **And do NOT read this as one corpus correcting another.** On the WunderGraph router's production
> residue **AKB was right and this corpus was wrong, in a fenced table** — because AKB reads the
> `infrastructure` repo this corpus had never cloned ([`org-repos.md` § 11](org-repos.md)). The two
> have different blind spots, not a ranking.
>
> **Reconciliation is not this milestone's to do** — AKB is a different repo, outside the two-repo
> scope, and no edit here can reach it. It is filed with an owner as
> `PLATFORM-M257x-akb-taxonomy-figures-contradict-measurement` in
> [`platform-defect-register.md`](../../knowledge/plan/platform-defect-register.md).

The whole product is the `NodeID` type and its generators/validators:

* `NodeID` (a `string`) with `MarshalGQL`/`UnmarshalGQL` — satisfies gqlgen's
  marshaler **structurally** (no gqlgen import), so malformed IDs are rejected at the
  GraphQL boundary.
* Constructors: `NewSkillID` (K), `NewSpecializationID` (S), `NewCategoryID` (C),
  `NewJobRoleID` (J), `NewIndustryID` (I). As of `v1.2.0` all except Industry take an
  `organization *string` so IDs can be **org-scoped** (org folded into the hash).

**Canonical ID format**: `<PREFIX>-<WORDPART>-<HASH>` matching
`^[CSKJI]-[0-9A-Z_]{6}-[0-9A-F]{4}$` — PREFIX = C/S/K/J/I; WORDPART = 6 uppercased
alphanumerics (X-padded; multi-word = 3 chars/word; `.`/`+`/`#` → `DOT`/`P`/`SHARP`);
HASH = first 4 hex of SHA-1 of the sanitized name(+org). Deterministic and
cross-language-consistent (a matching Python implementation exists for data pipelines).
Example: `NewSkillID("go") = K-GOXXXX-F63F`, `".net" = K-DOTNET-DDE9`, `"c#" = K-CSHAXX-5F5B`.

> Note: service-local packages named `taxonomy` (e.g.
> `app/internal/taxonomy`) are **distinct** from this module — don't confuse them.

---

## Related Documentation

* [Service Taxonomy](./service_taxonomy.md) — where these sit in the three-tier model
* [Dependency Map](./dependency_map.md) — who imports what + the Redis Streams event map
* [AI Architecture](./ai_architecture.md) — model inventory, routing, cost telemetry (the consumer side of `ai`)
* [Sentinel](../services/sentinel.md) — the authorization service `colony/authorization` calls
* [Architecture Overview](./architecture_overview.md)
