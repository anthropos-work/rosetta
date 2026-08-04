# Shared Libraries

The Anthropos Go services share five internal libraries. They are **not deployed
services** — there is no container, port, or `docker-compose` entry for any of them.
They are **Go modules** compiled *into* each service's binary.

## High-Level Summary (For PMs & Non-Engineers)

Think of these as the platform's "standard library." Rather than every microservice
re-implementing logging, database wiring, authentication, RPC contracts, or AI calls,
that shared plumbing lives in five small repos that the services pull in like any
third-party dependency. This keeps the services consistent and small.

| Library | One-liner |
|---------|-----------|
| **colony** | The framework: DB/Redis, logging+Sentry, GraphQL/RPC servers, CORS, pub/sub, feature flags — and it now contains **authn** |
| **proto** | The single source of truth for service-to-service RPC contracts (Protobuf → Connect-RPC) + hand-written domain types |
| **ai** | A thin wrapper over OpenAI/Azure/Anthropic/Bedrock/Mistral behind one `ai.AI` interface |
| **authn** | Clerk JWT authentication (now shipped **inside colony** as `colony/authn`; the standalone repo is legacy) |
| **taxonomy** | The **node-id library** (`NodeID` type + ID generation/validation) — **not** a dataset |

> ### How they are consumed (this matters)
> **None of these are cloned by `make init`** — they are **absent from
> `platform/repos.yml`**, so there is no `stack-dev/colony` (etc.) directory.
> Each is pulled as a **private Go module** during a service's Docker build:
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
| **Version pin** | **Split — there is no single pin.** `app` + `messenger` → **`v0.35.2`**; the still-running `cms` + `jobsimulation` husk containers → **`v0.35.1`**; `sentinel` + `storage` + **`roadrunner`** → `v0.34.3` (archived `chronos` pins `v0.30.1`). Measured from each repo's `go.mod` at platform `2adcf71`. |
| **Imported by** | **Every** live Go service: app, sentinel, storage, messenger — plus the `cms`, `jobsimulation` and **`roadrunner`** containers, which the default `graphql` profile still starts as merged-into-`app` **husks** (rollback path; teardown is M810) even though their domains now run inside `app`. Roadrunner's import is minimal but real: `roadrunner/main.go:7` imports `colony` for `NewVersionConfig` (`go.mod:7` pins `v0.34.3`) |

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
| **Version pin** | one pin per repo, and the skew is **live and wider than the merge suggests**: app/messenger `v1.210.0`, cms `v1.207.0`, jobsimulation `v1.205.0`, sentinel `v1.200.0`, storage/roadrunner `v1.196.0`. The husk repos **still carry their own `go.mod`** — `repos.yml:14-19` still clones them and `docker-compose.yml:83,144` still builds them |
| **Imported by** | every live Go service that does RPC (app, sentinel, storage, messenger). The cms / jobsimulation / **skiller** RPC surfaces are served in-process by `app`; **skillpath and roadrunner were REMOVED, not re-hosted** — `app/main.go` registers six Connect handlers (Users `:1178`, Organizations `:1179`, Skiller `:1187`, JobSimulation `:1195`, CMS `:1204`, LabSession `:1218`) and neither `SkillPathSessionService` nor a RoadRunner service is among them |

The **single source of truth for RPC contracts**. Two layers:

* **Generated** — `proto/<svc>/v1/*.proto` → `go/<svc>/v1/` (message structs) + `go/<svc>/v1/<svc>v1connect/` (Connect-RPC stubs, *do not edit*).
* **Hand-written** — `go/domain/<svc>/` idiomatic Go types (string enums, `time.Time`) plus **goverter**-generated converters. goverter fails codegen if a proto enum value has no matching domain const — the "three-file rule" (proto + domain const + `make gen`).

12 Connect-RPC services are defined: `UsersService`, `OrganizationsService`,
`CMSService`, `JobSimulationService`, `SkillerService` (all served by app since the merges — one RPC mux),
`SkillPathSessionService` (**contract still in `proto`, but NO LONGER SERVED** — like `ChronosService`. skillpath-in-app M506 *removed* the RPC rather than re-hosting it; `app/internal/skillpaths/skillpaths.go:27-31` calls its replacement "the drop-in for the **removed** skillpath RPC client". Likewise roadrunner: `backend` calls Judge0 over plain HTTP at `internal/jobsimwiring/wiring.go:118`, and `ROADRUNNER_RPC_ADDR` (`docker-compose.yml:118`) is read by no Go code in `app`),
`LabSessionService` (served by app — the AI Labs domain, see `../services/ai-labs.md`),
`AuthorizationService` (Sentinel), `MessengerService`, `RoadRunnerService`,
`RealtimeService`, `ChronosService` (archived service, contract still present). Plus
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
| **Version pin** | **`v1.40.2`** across consumers (`app`, and the `cms` / `jobsimulation` husks, all agree) |
| **Imported by** | app — i.e. every folded domain — **and, as their own direct `go.mod` requires, the still-running `cms` and `jobsimulation` husk containers** (`cms/go.mod:9`, `jobsimulation/go.mod:11`, both `v1.40.2`). Go services only — **not** Studio-Desk, which is TypeScript, and **not** roadrunner, whose only shared-lib requires are colony + proto |

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
>    `external_services.md:546` retracts that chain.) Vendor selection lives in each consumer's own
>    `internal/ai/ai.go` wrapper: an EU Azure client by default, a US Azure client gated
>    by the PostHog flag `flag_use_azure_us`, and an Azure→direct-OpenAI fallback on
>    HTTP 429. Anthropic is always Bedrock in `eu-west-1`.

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
| **Imported by** | via colony: app — **and the cms and jobsimulation husk repos still import `colony/authn` directly** (6 and 8 files at `cms ca50c817` / `jobsimulation 462343b0`). Only the skillpath usage is fully folded in |

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
| **Imported by** | directly: app, messenger — **and the still-running `cms` + `jobsimulation` husk containers**, which require it directly in their own `go.mod` (`cms/go.mod:13`, `jobsimulation/go.mod:15`, both `v1.2.0`, neither marked `// indirect`); indirectly (`// indirect`): storage, sentinel. **6 of the 7 live Go service repos** (app, sentinel, storage, messenger, cms, jobsimulation, roadrunner) — the sole exception is `roadrunner`, which requires only colony + proto. (The skillpath usage is folded into app.) |

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
