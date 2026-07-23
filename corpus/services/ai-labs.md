# AI Labs + Credits (`app` domain)

> **Not a standalone service — a set of `app`-monolith domains** (the service the platform calls "backend").
> There is no separate container, repo, or subgraph. Code lives under `app/internal/labs/`,
> `app/internal/credits/`, `app/internal/payments/`, `app/internal/subscriptions/`, plus the top-level
> `app/stripe/` fixtures.

> ### ⚠️ Two different meanings of "v6.0 shared purse" — read this first
> - The credit **ledger primitive** (`internal/credits`, "Wave 6") **is built and live** — but wired **only to
>   [Course Builder](./coursebuilder.md)**, not to AI Labs.
> - The **AI Labs self-serve credits initiative** branded **v6.0 "shared purse"** (org self-serve *buy* AI-Labs
>   credits + an *enforcing* shared-pool wallet) is **DESIGNED / QUEUED, NOT BUILT** — a knowledge-plan release
>   (`app` `knowledge/plan/releases/06.00-shared-purse/`, milestones M600–M607, all planned). In current `origin/main`
>   there is **no `checkout.session.completed` webhook, no labs↔credits linkage, and `/credits/purchase` was removed
>   (Wave 13)**. `v6.0` is a **knowledge-plan release number, NOT the `app` SemVer** (`app` is at `v1.351.1`).
> This doc documents the **shipped reality**; the shared-purse unification is flagged as planned where relevant.

## Role & Responsibility

Two independent concerns absorbed into `app`, joined today only by a shared *reuse target* (the credit ledger), not
yet by code:

*   **(a) AI Labs** — a hosted, scenario-based **AI coding-lab** product. A user launches an ephemeral cloud sandbox
    (a microVM running an in-VM coding agent — Claude Code or OpenAI "codex") to work a lab scenario, metered in **USD
    per session**. `app` owns the **tenant-facing view** (catalog content + session identity/budget/status/billing
    record + org spend caps); a **separate control plane, `labs-api`**, owns the microVM lifecycle, the per-session
    LiteLLM virtual key, and the live spend numbers. Two sub-domains: a **catalog** (the authored "Labs as a product"
    library) and a **session/runtime** (boot/track/cancel a sandbox + enforce spend caps).
*   **(b) The credit ledger ("shared purse")** — a usage-metering + billing ledger: an **org-level shared credit
    pool** ("purse", one balance per organization that every member draws from) that meters premium generative
    actions. Every billable action writes an **append-only ledger row** *and* updates the org's running balance in
    one atomic transaction. **Today it meters [Course Builder](./coursebuilder.md) only** (`course.build`/`refine`/
    `translate`); AI Labs metering rides on a *separate* USD wallet (`labs_budget`).
*   **(c) Payments / subscriptions** — the Stripe surround: user premium **subscriptions** + org **seat licensing**,
    both backed by Stripe. Distinct from the credit ledger.

## Architecture & Code Map

*   **Codebase**: `app/internal/{labs,credits,payments,subscriptions}/` + `app/stripe/` (fixtures). Go, part of the
    `app` monolith. DB: PostgreSQL, **`public` schema**, Ent + Atlas.

### AI Labs — catalog (`internal/labs/catalog/`)
*   `catalog.Manager` (slug catalog domain) + `ContentManager` (viewer-facing, **fail-closed** tenant-filtered
    reads; `ErrNotFound` is a hard-404 no-existence-oracle), `content_import.go` (idempotent manifest upsert),
    `s3_workspace_store.go` (S3-backed workspace/grader/solution assets — the solution is private, never served to
    learners).
*   **Tables**: `labs` (Ent `lab.go` — metadata: `slug`, `tenant_eid` nullable = org-scoped else public, `lifecycle`,
    `base_image`, `default_model`, `budget_usd`, S3 asset refs) and `lab_bodies` (Ent `lab_body.go` — one row per
    `(lab_slug, locale)`; EN base + 6 overlays; `{meta,brief,tutor,hints}` as one jsonb). Migration
    `20260617203555_add_labs_catalog.sql`.

### AI Labs — session/runtime (`internal/labs/session/`, `labsapi/`, `adapter/`)
*   `labsession.Server` implements the `lab.v1.LabSessionService` Connect-RPC (Create/Get/List/Cancel/ReportEvent);
    `Manager` holds the business logic the GraphQL resolvers call directly. `labs_budget.go` is the **cap/tier
    engine**: `LabsTier` (`essential`/`professional`/`frontier`) → `modelsForTier`, `LabsBudgetOptions` stored as JSON
    on `OrganizationSettings.Options[labs_budget]`, `CheckRequest` → typed `CapError{Reason}` (`model_not_in_tier`,
    `per_session_cap`, `org_total`, `user_monthly`, `user_lifetime`) → Connect `ResourceExhausted`.
    `spend_reconciler.go` is a **30 s pull loop** that denormalizes `spend_usd`/`total_tokens`/`status` from labs-api
    into `lab_sessions` rows (+ an orphan-fail-out pass).
*   `labsapi/client.go` — HTTP client to **labs-api** (default `:7070`, bearer `LABS_API_PLATFORM_TOKEN`); labs-api
    **owns session-ID generation**. `adapter/adapter.go` bridges the client to the `LabsAPIClient` seam.
*   **Table**: `lab_sessions` (Ent `lab_session.go`) — deliberately **no `PrimaryKeyMixin`**: `id` is a **String**
    supplied by labs-api. `user_id`, nullable `organization_id`, frozen `template` (= `Lab.slug`), `model`, `mode`,
    `status` (booting/ready/grading/stopped/failed/cancelled), immutable `budget_usd`, denormalized `spend_usd`/
    `total_tokens`, `grade_result` jsonb. Migrations `20260529072659`, `20260617120000`.

### The credit ledger / "shared purse" (`internal/credits/`)
*   `credits.Manager`: `Balance` (lazy-seeds `organization_credits` at `DefaultSeedBalance=500` on first read),
    `Debit` (atomic conditional decrement + negative ledger row; `InsufficientCreditsError` → **HTTP 402**), `Refund`
    (positive row, **idempotent on `ref_id`**), `Purchase`, `Transactions` — every mutation in one `ent.WithTx` under
    `privacy.Allow`.
*   `cost.go` — the cost model: `creditCost` map (`course.build`=**5/chapter**, `course.refine`=**1**,
    `course.translate`=**1/locale**), `purchasePackages` (starter 50 / team 200 / scale 500), `PricePerCreditUSD` =
    **$0.45** (derived from Bedrock COGS × `MarginMarkup` **1.40**). `reconcile.go` — `ReconcileOrphanedBuilds`
    (refunds crashed pre-terminal Course Builder builds, idempotent) + `ReconcileOrg`/`MarginAlert` (internal
    per-org realized-margin; never served on customer routes).
*   **Tables**: `organization_credits` (Ent `organization_credit.go` — per-org `balance`, UNIQUE on `organization_id`
    = the one-purse-per-org invariant) and `credit_transactions` (Ent `credit_transaction.go` — append-only:
    `action_type`, signed `credits_delta`, `ref_id`, `meta`; a **partial-unique** `(organization_id, action_type,
    ref_id) WHERE ref_id IS NOT NULL` = the idempotency guard). Migration `20260717151144.sql`.

### Payments / subscriptions (Stripe — `internal/payments/`, `internal/subscriptions/`)
*   `payments.Manager` (`stripe-go/v74`): `CreateCustomer`, `CreateCustomerCheckoutSession` (**subscription mode**,
    price by `lookup_key`, 7-day trial), portal session, `ParseEvent` (webhook signature via
    `STRIPE_ENDPOINT_SECRET`), `HandleStripeSubscriptionEvent` (upserts user `subscriptions`), `emitMetaPurchase`
    (Meta Conversions API). `subscription_verifier.go` implements `organization.StripeSubscriptionVerifier` so **org
    seat licensing derives seats from the real Stripe quantity**, not client input. `subscriptions.SubscriptionManager`
    handles user premium subs + **feature-usage quotas** (cover-letters / job-applications / job-simulations) via
    Sentinel.
*   **Tables**: `subscriptions` (user-scoped), `stripe_customers`, `org_subscriptions` (org seat licensing:
    `licenses`/seats, `source` = stripe|manual, Stripe ids, period). Migration `20260626120000_add_org_subscriptions.sql`.

## Interface Discovery

*   **How to find the API** — three surfaces:
    *   **GraphQL** (on the `app`/`backend` subgraph): `internal/web/backend/graphql/graph/schemas/labs.graphqls`
        (`Lab`, `LabSession`, queries `labs`/`labBySlug`/`labSessions`/`mySessions`/`labSpendByDay`/…, mutations
        `createLabSession`/`cancelLabSession`); `billing.graphqls`; org licensing fields in `organizations.graphqls`.
        **Credits has NO GraphQL** — REST only.
    *   **Connect-RPC**: `lab.v1.LabSessionService` (Create/Get/List/Cancel/**ReportEvent**) — uniquely wrapped in
        Clerk `authn.HTTPAuthnMiddleware` (it reads caller identity from context); `ReportEvent` is the **labs-api →
        app callback** (allowlist: `status`, `grade`). Credits/payments/subscriptions expose no RPC of their own.
    *   **HTTP/REST** (Echo): credits `GET /credits/balance` + `/credits/transactions` (org-admin gated; **`POST
        /credits/purchase` was removed, Wave 13**); labs catalog write `POST /v1/labs` + `/v1/labs/bodies` (org
        API-key, scope `labs:write`) + `GET /v1/labs/:slug/workspace.tar.gz`; the **Stripe webhook** `POST
        /api/webhook/stripe` (auth-skipped, signature-verified; handles **`customer.created` +
        `customer.subscription.{created,updated,deleted}` only** — **no `checkout.session.completed`**, the un-built
        shared-purse grant).
*   **Dependencies**:
    *   **Upstream (callers)**: the Next.js `next-web-app` (Labs UI at `/labs`, enterprise dashboards, checkout, the
        credits panel); [Course Builder](./coursebuilder.md) (the only credits `Debit`/`Refund` caller); the org
        licensing manager (calls `SubscriptionVerifier`).
    *   **Downstream**: **Stripe** (`stripe-go/v74`), the **labs-api** control plane (microVMs + a LiteLLM gateway →
        Bedrock/Anthropic + Azure/OpenAI), **AWS S3** (labs workspace assets), the **Meta Conversions API**, and
        Sentinel (authz quotas/seats).

## Local Development

### 1. Running Standalone
*   **Prerequisites**: runs **as part of `app`** (`go run .` / Docker). Labs wire up only when `LABS_API_URL` is set
    (else the labs-api client is nil → in-memory idgen, reconciler no-op); credits seed lazily; payments need Stripe
    keys.
*   **Key env vars**: labs — `LABS_API_URL`, `LABS_API_PLATFORM_TOKEN`, `LABS_WORKSPACE_BUCKET`,
    `LABS_WORKSPACE_PREFIX`, `AWS_REGION`; Stripe — `STRIPE_SECRET_KEY`, `STRIPE_ENDPOINT_SECRET`; Meta —
    `META_PIXEL_ID`, `META_CONVERSION_API_TOKEN` (self-disables when unset).
*   **Stripe fixture flow** (`app/stripe/`): `setup.sh` → `stripe fixtures setup-fixtures.json` creates the premium
    product + monthly/yearly prices; `teardown.sh` archives them; `tests/stripe-test.sh` runs `stripe listen
    --forward-to localhost:8080/api/webhook/stripe`.
*   **Dedicated cmds**: `cmd/labsimport` (idempotent catalog manifest loader), `cmd/labskey` (mint an org API key
    for `labs:write`), `cmd/labsdemo` (standalone local Labs-catalog server on `:8099`, in-memory store, dev only).

### 2. Running in Docker
*   **Service Name**: none of its own — part of the `backend` service (`docker compose up -d backend`).

### 3. Testing
```bash
go test ./internal/labs/... ./internal/credits/... ./internal/subscriptions/...
```
*   **AI Labs**: `manager_test.go`, `labs_budget_test.go` (tier/cap matrix), `spend_reconciler_test.go`,
    `create_e2e_test.go`, `labsapi/client_test.go`, catalog `content_test.go`/`workspace_test.go`.
*   **Credits**: `cost_test.go` (`TestPricingDerivation` guards the $0.45/40 % derivation), `manager_test.go`
    (Balance/Debit/Refund/Purchase + concurrency + lazy-seed), `privacy_backstop_test.go`
    (`TestPrivacyBackstop_AutoScopesToCallerOrg` — proves the `OrganizationMixin` fences cross-org reads),
    `reconcile_orphan_test.go`.
*   **Subscriptions**: `subscriptions_test.go`. **Payments** has no Go unit test — it is exercised via the `stripe/`
    CLI harness (`stripe listen` against `/api/webhook/stripe`).

## Related Documentation

* [Backend (`app`)](./backend.md) — the monolith these domains live in
* [Course Builder (`app` domain)](./coursebuilder.md) — the sole current consumer of the credit ledger
* [Ask engine / Talk-to-Data](./askengine.md) — the shared Bedrock transport
* [AI Architecture](../architecture/ai_architecture.md) — provider routing; [External Services](../architecture/external_services.md) — Stripe/S3
