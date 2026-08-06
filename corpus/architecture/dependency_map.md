# Anthropos Services Dependency Map

This document outlines the inter-service dependencies inferred from configuration files (`docker-compose.yaml`) and code inspections.

## Dependency Matrix

Sourced from `platform/docker-compose.yml` `depends_on:` declarations and environment variables (`*_RPC_ADDR`).

> Since the monolith merge most of this matrix collapsed: `skiller`, `skillpath`, `roadrunner`, `jobsimulation`, `cms` — and, since **v9.0 "support-in-app"** (2026-08-04), `messenger`, `storage` and `customerio-sync` — are domains inside **Backend (`app`)**, so their edges are in-process calls, not dependencies.
>
> **What is left is one edge: `backend` → `sentinel`.** That is the whole inter-process matrix now (Gotenberg is a third-party container, not an Anthropos service).

| Service | Depends On (Direct) | Infrastructure |
| :--- | :--- | :--- |
| **Backend** (`app`) — the monolith | **Sentinel** (compose `depends_on`; the only Anthropos service it calls); Gotenberg (runtime HTTP, no startup-order dep) | Postgres (`public` schema; `pgvector` in `extensions` — skiller embeddings, skill-path sessions, the 23 jobsim run-state tables, the cms similarity/Studio tables), Redis (cache + streams + asynq), **S3** (both buckets, written directly since v9.0), **Clerk**, **Directus**, **Judge0**, **LiveKit**, **AWS Chime**, **Brevo**, **AI Providers** |
| **Sentinel** | - | Postgres |
| ~~**CMS**~~ | **Merged into `app`** ("cms-in-app v8.0", app v1.360.0) — the content layer + Studio run in-process; Directus stays external | *(no standalone service)* |
| ~~**Jobsimulation**~~ | **Merged into `app`** ("jobsim-in-app") — the session engine runs in-process; simulation definitions come from the cms domain by ID without an RPC hop | *(no standalone service)* |
| ~~**Skillpath**~~ | **Merged into `app`** ("skillpath-in-app", M502→M507) — the skill-path engine's dependencies (cms content by ID, the jobsimulation Redis Stream, Sentinel) are now `app`'s, in-process | *(no standalone service)* |
| ~~**Roadrunner**~~ | **Merged into `app`** with jobsim-in-app — `backend` calls Judge0 directly via `JUDGE0_BASE_URL` | *(no standalone service)* |
| ~~**Storage**~~ | **Merged into `app`** (v9.0) — `backend` reads/writes both S3 buckets in-process; `STORAGE_RPC_ADDR` is gone. Still startable from the `storage-legacy` profile as the rollback path | *(no standalone service; the terraform module survives and owns the buckets + CloudFront)* |
| ~~**Messenger**~~ | **Merged into `app`** (v9.0) — the mailer runs in-process behind `MESSENGER_ENABLED`, on messenger's own Redis consumer group. Still startable from the `messenger` profile as the rollback path | *(no standalone service)* |
| ~~**CustomerIO Sync**~~ | **Merged into `app`** (v9.0) — the 10-minute **Brevo** marketing-contact push runs on app's asynq scheduler behind `CUSTOMERIO_SYNC_ENABLED`. Terraform module deleted: **no rollback path**. Still declared in compose (profiles `customerio-sync`, `all`) | *(no standalone service)* |
| ~~**Graphql (Cosmo Router)**~~ | **RETIRED 2026-07-31 — no service, no compose entry, no row.** With `backend` the only subgraph, the router was a pure extra hop. Every GraphQL client now depends on **Backend** directly | - |
| **Gotenberg** | - | - (stateless conversion service, third-party image) |
| **Studio-Desk** (opt-in profile) | **Backend** (GraphQL, `VITE_GRAPHQL_ENDPOINT` → `:8082/graphql/query`) | **Clerk**, **OpenAI / Azure OpenAI / Anthropic** (Copilot, via `AI_PROVIDER_CHAIN`) |
| **Studio-Room** | (runs inside the `app` container; depends on the backend process) | **OpenAI**, **Anthropic**, **Mistral** |

> **Skiller merged into app (July 2026):** the standalone skiller service is gone from the compose file. Its RPC surface is now served by **backend** — consumers keep the `SKILLER_RPC_ADDR` env var, re-pointed at `http://backend:8083` (production terraform: `skiller_rpc_addr = http://backend:8081`). See [Backend](../services/backend.md) and the [skiller stub](../services/skiller.md).
>
> **Skillpath merged into app (skillpath-in-app, M502→M507):** the standalone skillpath service is gone from the compose file / repos.yml / and (while it still existed) the supergraph. Its skill-path progression engine now runs **in-process inside `app`**, with session state in `public.skill_path_sessions` (the legacy `skillpath` schema is an empty husk). See [Backend](../services/backend.md) and the [skillpath stub](../services/skillpath.md).
>
> **Jobsimulation + cms merged into app (jobsim-in-app, cms-in-app v8.0):** the last two subgraph services are gone from the compose file / repos.yml / supergraph — which took the federation to **one** subgraph and made the router redundant. **The router was then retired 2026-07-31**; there is no supergraph at all now, and clients call `backend`'s own endpoint. Their tables were re-created in `public` (the legacy `jobsimulation` and `cms` schemas are non-authoritative). Their ECS modules are **still declared** in production terraform as the rollback path and take no traffic; teardown is **M810**. See [Jobsimulation](../services/jobsimulation.md) and [CMS](../services/cms.md).
>
> **Messenger + storage + customerio-sync merged into app (v9.0 "support-in-app", 2026-08-04):** the last three support services folded in, leaving **`sentinel` as the only out-of-process Anthropos service**. Three consequences for this matrix: (1) `backend`'s Connect-RPC mux has **no external callers left** — `messenger` was the last, so the mux's remaining service definitions exist for the frozen repos' pinned builds, not for live traffic; (2) `backend` → S3 is now a **direct** edge, not an RPC hop, and its bucket names come from `STORAGE_S3_BUCKET` / `STORAGE_S3_PUBLIC_BUCKET` (`STORAGE_RPC_ADDR` is gone); (3) **Brevo is now a direct `backend` dependency** — one `BREVO_KEY` covers transactional mail, product tracking and the marketing-contact sync. Both outbound subsystems are gated (`MESSENGER_ENABLED`, `CUSTOMERIO_SYNC_ENABLED`) and off unless switched on by name. See [Messenger](../services/messenger.md), [Storage](../services/storage.md), [CustomerIO Sync](../services/customerio-sync.md).
>
> **Content-vs-runtime dependency (unchanged, now in-process):** both the skill-path engine and the jobsimulation engine depend on the **cms domain for content/definitions** — cms is the content layer; they are runtime/session engines that hold no content and reference cms artifacts **by ID**. The skill-path engine fetches a path's chapter/step structure when (re)building a session; the jobsimulation engine loads a simulation's definition before running it. Both calls used to be Connect-RPC (`CMS_RPC_ADDR`, `cms.GetSimulation`); they are **plain function calls** now. The jobsim domain still holds no `DIRECTUS_BASE_ADDR` of its own — its Directus reads flow *through* the cms domain. (See [CMS](../services/cms.md), [Skillpath](../services/skillpath.md), [Jobsimulation](../services/jobsimulation.md).)

Production-only:
| Service | Depends On (Direct) | Infrastructure |
| :--- | :--- | :--- |
| **db-backup** | - | Postgres, **S3**, **Azure**, **Hetzner** |

### Shared Libraries

Imported as private Go modules (not deployed, **not** cloned by `make init`). Full reference: [Shared Libraries](./shared_libraries.md).

| Library | Used By |
| :--- | :--- |
| **colony** | `app` + `sentinel` (deployed) and the frozen `storage` / `messenger` repos, which still import it at pinned tags (logging, DB, Redis, middleware, pub/sub); also bundles `authn` |
| **proto** | The same four repos. It still carries the `MessengerService` / `StorageService` contracts — **don't delete them**, they are what the frozen repos build against |
| **ai** | app — i.e. every folded domain (AI provider wrapper — Go services only, not Studio-Desk). Cost & routing live in the consumers, not the lib |
| **authn** | Imported via `colony/authn` by app (standalone `authn` repo is legacy; the former cms/jobsimulation/skillpath usage is folded into app) |
| **taxonomy** | **node-id library** (not data): direct — app, messenger; indirect — storage, sentinel |

> Only **two** of those four repos still deploy: `app` and `sentinel`. `storage` and `messenger` are frozen rollback targets — they build, they import the libraries at pinned tags, and they take no traffic.

## Event Streams (Redis Streams via Watermill)

Services communicate asynchronously through named Redis Streams. Stream names come from `*_STREAM` env vars in `platform/docker-compose.yml`.

| Stream Name | Producer | Consumer(s) | Events |
| :--- | :--- | :--- | :--- |
| `backend` | App | CMS | User/org updates |
| `skiller` | App | App | Skill score changes — both producer and consumer live inside app since the skiller→app merge (stream name retained) |
| `jobsimulation` | App | App — the jobsim engine + the skill-path engine on ONE subscriber, **plus the folded messenger on its OWN subscriber** (see below) | Session completed, insights generated |
| `cms` | App (+ **Directus webhooks** → `POST /api/webhook/directus`, now authenticated via `DIRECTUS_WEBHOOK_SECRET`) | App (the cms similarity/Studio handlers + the jobsim handlers, merged onto ONE subscriber) | Content published/updated, translation & clone requests |
| `skillpath` | App | App | Session updated, chapters completed — both producer and consumer live inside app since the skillpath→app merge (stream name retained) |
| ~~`roadrunner`~~ | — | **no producer or consumer** — roadrunner is merged into app, which calls Judge0 synchronously | ~~`RoadrunnerSubmissionCompleted`~~ |
| `AI` | (multiple) | (multiple) | AI usage / cost telemetry — see `AI_USAGE_STREAM=AI` env var |

> **Note**: The `chronos` stream was previously used by Chronos for timer events but is gone with the chronos service removal. Jobsimulation no longer has chronos as a dependency.

> **The one place two subscribers share a stream — deliberately.** Everywhere else, `app` merges new
> handlers onto the **existing** subscriber with `.AddHandler(...)`, because colony keys subscribers by
> stream name and a second `AddSubscriber` for the same stream silently overwrites the first. The
> folded **messenger** is the exception: it runs on a **second, dedicated `SubscriberServer`** attached
> to messenger's **own** Redis consumer group (the literal `messenger`). Two reasons — messenger
> subscribes to streams `app` already subscribes to, so a shared server would have silently replaced
> app's handlers; and re-using the pre-existing group means Redis keeps the cursor, so the cutover from
> the standalone had **no gap**. Different consumer groups on the same stream each see every entry, so
> the two subscribers do not steal work from one another. `backend` verifies the group exists at boot
> rather than creating a fresh one, and the whole block only runs when `MESSENGER_ENABLED` is on.
>
> The corollary is the operational rule: **never run the standalone `messenger` container alongside a
> `MESSENGER_ENABLED=true` backend.** Those two DO share a group, and entries claimed by one are
> invisible to the other — you get a coin flip over which process sends each email. Platform `0dab54d`
> dropped `messenger` from the `all` compose profile for exactly this reason.

## Key Flows

### 1. User Authentication
`Frontend` -> `Backend` -> `Sentinel`
*   The Backend validates requests using Sentinel.
*   **Studio Desk** authenticates directly via **Clerk**.

### 2. Job Simulation
`Frontend` -> `Backend` / `Jobsimulation`
*   The jobsimulation engine fetches the simulation **definition** (the `simulations` content/blueprint) from the cms domain by ID — in-process since cms-in-app. It owns no content, only the run/session state.
*   The jobsimulation engine stores its session/run **state** (interactions, recordings, validation results, anti-cheat) in the **`public`** schema (the legacy `jobsimulation` schema is non-authoritative), and its binary objects **straight to S3** — the storage manager is in-process since v9.0, so there is no `Storage` RPC hop.
*   Voice flows go through LiveKit; video recordings via AWS Chime SDK.

### 3. Content Delivery
`Frontend` -> `CMS` -> `Directus`
*   CMS acts as the gateway to Directus content.

### 4. Studio Content Creation
`Studio Desk` → `CMS` → (in-process) `Studio Room`
*   **Studio-Desk** (TypeScript) creates blueprints, sent to CMS as `StudioDocument` rows.
*   **CMS** (Go) creates `StudioTask` records and dispatches generation work.
*   **Studio-Room** (Python, embedded inside the **`app`** container since cms-in-app) executes the generation pipeline against AI providers (OpenAI, Anthropic, Mistral).
*   Final content is persisted via the CMS service; **Directus** is the underlying storage backend.

### 5. Skill Path Progress (Event-Driven)
`Jobsimulation` -> `Redis Stream` -> `App (skill-path engine, in-process)`
*   When a user completes a simulation, **Jobsimulation** publishes an event.
*   The **skill-path engine — now inside `app`** (merged from the standalone skillpath service, M502→M507) — subscribes to the Jobsimulation stream and updates step/chapter/path progress **state** (`SkillPathSession → ChapterSession → StepSession`, in `public.skill_path_sessions`) — it owns no content, only the per-user progression state.
*   The engine queries **CMS** (RPC) for the skill-path **content** structure (chapters → steps it tracks against) and **Sentinel** for authorization. All of this now runs in-process inside `app`.

### 6. Document → PDF Conversion
`Backend (app)` → `Gotenberg`
*   The backend service uses Gotenberg's `/forms/libreoffice/convert` endpoint to render Office documents to PDF. See `app/internal/converter/gotenberg.go`.
*   `GOTENBERG_URL=http://gotenberg:3200` is injected via the backend's compose env.

### 7. Transactional Email (Event-Driven)
`App (any domain)` → `Redis Stream` → `App (messenger flow handlers, in-process)` → `Brevo`
*   A domain event (session completed, assignment created, org invitation, …) is published to one of the application streams.
*   The **folded messenger** — `app/internal/messenger/flow/`, running on its **own** subscriber and its **own** consumer group — decides whether the event should produce an email, picks the template, applies staleness guards and per-org whitelabel branding, renders through Liquid, and sends via **Brevo**.
*   Gated by **`MESSENGER_ENABLED`**; `BREVO_KEY` is required when it is on. With the switch off (the default on a developer machine) no handler is registered at all — the events still flow, nothing mails.

### 8. Marketing-Contact Sync (Scheduled)
`App (asynq scheduler, every 10 min)` → `public` tables → `Brevo` contacts
*   `app/internal/customeriosync/` reads a fixed **overlap window** wider than the schedule — the standalone's in-memory `lastSyncTime` did not survive a multi-replica worker that restarts on every deploy — and pushes each user as a Brevo marketing contact.
*   The push is idempotent (`CreateContact` with `UpdateEnabled`), so overlapping windows cost nothing and a missed run is covered by the next one.
*   Gated by **`CUSTOMERIO_SYNC_ENABLED`**. The name is a fossil: the destination is Brevo, not Customer.io.
*   It no longer reads the `public.customer_io_sync_table` view, which was **`public`'s only cross-schema dependency** — it would have been a silent casualty of any legacy-schema drop. The query now lives in `internal/customeriosync/sync_query.sql` against final `public` tables.
