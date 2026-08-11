# Academy Backend (`app` domain)

> **Not a standalone service — a domain of the `app` monolith** (the service the platform calls "backend").
> Code lives in `app/internal/academy/`. **This is NOT the [Ant Academy frontend](./ant-academy.md)** (the
> Next.js 16 + Expo Vercel-deployed learning portal). This doc covers the **server-side academy domain** that
> **owns** the course catalog + every learner's study state and **serves** it to the frontend over GraphQL.

> ### The frontend/backend split (read this first)
> - **Frontend — [`ant-academy`](./ant-academy.md)**: the Next.js/Expo *reader/UI* (Vercel-deployed).
> - **Backend — `app/internal/academy/` (this doc)**: the *source of truth*. It owns the `series → skill_path →
>   chapter → body` catalog **and** all per-user study state (progress, time, bookmarks, feedback, certificates),
>   and serves them to the frontend through the **`app` GraphQL subgraph**, reached at
>   `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` — which on a local stack is **`backend`'s own `:8082/graphql/query`**
>   since platform `2adcf71` (2026-07-31) deleted the Cosmo/WunderGraph router from compose. The env-var
>   *name* outlived the router, which is now destroyed in production as well (iter-124). (Consistent with the **Primary — GraphQL, on the `app` (backend) subgraph** bullet under § *Interface Discovery* below — *"there is no separate 'academy subgraph'"*. ⚠️ **This carried a line range until M257x iter-138, by then the Certificate-minting bullet.** The pin was correct when written at `cd16967` (iter-102) and rotted **+8** from unrelated insertions above it, so it is replaced by the construct name rather than re-pinned.)
> - The backend became authoritative in app-side release **"v1.0 ground truth"** (PR #903, `0e37771f`, 2026-06-05):
>   the net-new server-owned academy domain replaced the legacy `internal/aiacademy` sync + `aiacademy_courses`
>   read-model. (The corpus's "since v0.5 M2 backend-authoritative" refers to the **frontend's** own version line —
>   when the FE started reading the DB catalog; the backend domain's labels are its own "v1.0/v1.05" line, distinct
>   from `app` SemVer, which runs in the `v1.3xx` range — **`v1.369.0`** @ `app` **`ad9f3c49`**, re-derived
>   **2026-08-06** (`git describe --tags ad9f3c49` → `v1.369.0-7-gad9f3c498`; the `v1.369.0` tag is dated
>   2026-08-05 in `CHANGELOG.md`). **A version is a reading at a ref, never a standing "current"**: this line
>   said `v1.363.2` while `:61` below cited `v1.367.0`, one document disagreeing with itself — and it then
>   attached the moving label *"origin/main"* to the sha `2035f9a4`, which **expired** when `app`'s branch head
>   moved 5 commits on to `ad9f3c49` (corrected M257x iter-102). `2035f9a4` still resolves and still means
>   exactly what it meant: it is the **label** that rotted, never the pin. A pinned sha here is left alone.)

## Role & Responsibility

*   **Primary Goal**: Be the server-authoritative owner of the AI Academy's content catalog and every learner's study
    state, served to the `ant-academy` frontend over the `app` GraphQL subgraph.
*   **Key Functions**:
    *   **Content catalog** — owns the `series → skill_path → chapter` metadata tree + the chapter **bodies**
        (modules/sections/references + per-locale overlays). Global content, tenant + tier gated.
    *   **Per-user study progress** — per-chapter completed-modules/sections progress (monotonic-merge), plus a
        "resume where you left off" singleton pointer.
    *   **Active-study-time ledger** — per-(user, chapter, device) cumulative time with a completion-vs-reinforcement
        phase split; the server-authoritative backup to PostHog.
    *   **Certificates of Mastery** — server-issued, publicly verifiable credentials (id `AAS-YYYY-XXXXXX`) with a
        PII-minimized public verify projection.
    *   **Bookmarks & annotations**, **feedback**, and **path embeddings** (pgvector) that feed the
        [AI-Readiness](./ai-readiness.md) recommendation engine.
*   **Design invariant**: **additive-only** — every table is NEW with a child-side forward edge to `User`; the legacy
    `skillpath`/`assignment`/`user_*` schemas are untouched. Per-user tables carry `UserMixin` (owner-only privacy;
    every read/write is self-only and fails closed for unauthenticated callers). Catalog tables have NO `UserMixin`
    (global content); their gating is a manager/resolver concern.

## Architecture & Code Map

*   **Codebase**: `app/internal/academy/`. Go, inside the `app` (backend) service. Persistence via Ent into the
    **`public`** Postgres schema (`academy_*` tables).
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
        **in-process** public storage manager, not an RPC hop: `app/main.go:524-525` **@ `ad9f3c49`**
        constructs the two managers (`internalstorage.NewManager` / `NewPublicManager`) from the bucket
        names read at `app/main.go:516-517` (`internalstorage.EnvBucket` / `EnvPublicBucket`, i.e.
        `STORAGE_S3_BUCKET` / `STORAGE_S3_PUBLIC_BUCKET`), and `STORAGE_RPC_ADDR` is gone from the Go
        source (3 remaining occurrences at that ref — `main.go:504`, `env/callsites_test.go:189`,
        `internal/jobsimwiring/wiring.go:101` — **all comments**, zero reads).
        **`:471-472` was this line's citation for four releases and it names neither construct** — at
        `ad9f3c49` it is the closing brace of the Bedrock-error branch plus a blank line. Found by the
        iter-122 claim census, which is why the ref is now stated with the pin rather than left to the
        cutover sha named earlier in the sentence.
*   **Ent tables** (`internal/data/ent/schema/`, `public` schema):
    *   **Per-user (UserMixin, owner-only)** — ⚠️ **these are Ent LABELS; the TABLE names are plural** (Ent
        pluralizes, so `SELECT … FROM academy_certificate` errors with *relation does not exist*):
        `academy_chapter_progresses` (jsonb completed-modules/sections;
        `UNIQUE(user_id, chapter_slug)` — the hottest dataset), `academy_last_activities` (per-user singleton resume
        pointer), `academy_chapter_times` (completion/reinforcement active-seconds split), `academy_certificates`
        (immutable; `cert_id` unique `AAS-YYYY-XXXXXX`; nullable `learner_email` never public), `academy_bookmarks`
        (kind-discriminated; global `external_id` idempotency key), `academy_feedbacks` (one editable row per
        (user, scope, target)).
    *   **Global catalog (no UserMixin):** `academy_series`, `academy_skill_paths` (`paid` default true — the tier
        lives here; nullable `tenant_eid`; `lifecycle`), `academy_chapters` (metadata; own `tenant_eid`),
        `academy_chapter_bodies` (one row per `(chapter_slug, locale)`; `body` jsonb = {meta, modules[], references[]};
        EN base + de/es/fr/it/nl/pt overlays), `academy_path_embeddings` (pgvector `extensions.vector(1536)`,
        `source_hash` skip-key).
*   **Certificate minting** (`certificate.go`): id `AAS-YYYY-XXXXXX` (`AAS` = "Anthropos Academy Skill"), a 6-char
    suffix from a 30-symbol ambiguity-free alphabet via `crypto/rand`. Public verify validates the id **before** any
    DB lookup and maps to a minimized PII-free struct (never `learner_email`; empty name → "AI Academy Learner").

## Interface Discovery

*   **How to find the API**:
    *   **Primary — GraphQL, on the `app` (backend) subgraph.** SDL: `internal/web/backend/graphql/graph/schemas/academy.graphqls`;
        resolvers: `resolver_academy.go`. **There is no separate "academy subgraph"** — these types live in the
        `app`/`backend` federation subgraph — **the only subgraph left**. The frontend reaches them at
        (`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`).
        *   **Queries** (self-only unless `@public`): `academyProgress`, `academyLastActivity`, `academyChapterTime(s)`,
            `academyPathProgress`, `academyCertificates` + `academyCertificate @public` (unauthenticated verify),
            `academyBookmarks`, `academyFeedback(s)`; **catalog** (`@public`, tenant-filtered): `academyCatalogSeries`,
            `academyCatalogSkillPaths`, `academySkillPath`, `academyChapter`; gated body `academyChapterBody` (returns
            a body-stripped teaser + `locked`/`lockReason` for non-entitled viewers, never full `modules`).
        *   **Mutations**: `upsertChapterProgress`/`…Batch`/`…BatchTolerant`, `setLastActivity`,
            `upsertChapterTimeBatch`, `claimPathCertificate`/`…FromLocal`, `add/update/removeAcademyBookmark`,
            `upsertAcademyFeedback`.
    *   **REST** (Echo): `GET /content/catalog.json` (unauthenticated public machine-index — only published, public
        paths; omits `paid`) and the **Academy Content Write API** `POST /content/admin/{series,skill-paths,chapters,
        bodies}` (guarded by a constant-time shared-token middleware reading `ACADEMY_CONTENT_API_TOKEN`, idempotent
        slug-keyed upserts) + an admin embeddings-recompute endpoint.
    *   **No Connect-RPC** for academy (GraphQL + REST only).
*   **Background task**: the nightly Asynq **`academy_embedding_refresh` @ 03:00** materializes
    `academy_path_embeddings` (skip via `source_hash = sha256(model | embed_text)`).
*   **Dependencies**:
    *   **Upstream (consumers)**: the [`ant-academy` frontend](./ant-academy.md) (over GraphQL/WunderGraph); the app's
        own [Ask/"Talk to Data"](./askengine.md) feed (reads `academy_*` directly); [Course Builder](./coursebuilder.md)
        (publishes generated courses into the catalog via `ContentManager` upserts).
    *   **Downstream**: **S3 → CloudFront** for asset upload — **in-process** since storage-in-app v9.0, *not*
        the standalone storage service (that RPC edge no longer exists); the shared **`ai`**
        library (embeddings); `internal/organization` (`OrgManager` for tenancy).

## Local Development

### 1. Running Standalone
*   **Prerequisites**: runs **as part of `app`** — no standalone academy process. Uses the shared platform Postgres
    (`public` schema). Env: `SUPABASE_DB_CONN` (the DSN the cmd tools use), **`STORAGE_S3_PUBLIC_BUCKET`** +
    AWS credentials for asset upload (**not** `STORAGE_RPC_ADDR` — `cmd/academyImport/main.go:239-243`
    @ `app` `9d00a313`: *"the standalone storage service is gone … it therefore needs the BUCKET, not an RPC
    address"*, and it hard-errors on an empty bucket at `:244-246`), `ACADEMY_CONTENT_API_TOKEN` (the write API).
*   **cmd binaries**:
    *   `cmd/academy-seed` — local dev/test seeder. Seeds chapter-progress + last-activity **through the Manager**
        (so writes go through the same monotonic-merge + self-only privacy paths, idempotent by construction). Flags:
        `--user-email`/`--user-id`, `--fixture`, `--reset`, `--dry-run`, `--list`. Local dev only.
        **Demo caveat**: on a demo the frontend has **no `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`**, so any rows
        `academy-seed` writes have **no reader** — i.e. `academy-seed` is **moot on a demo**. **It does NOT
        "serve its committed FS catalog": there is no FS-as-published fallback in the app** — this passage
        said there was until M257x iter-108, contradicting
        [`ant-academy.md`](./ant-academy.md)'s **“no FS-as-published fallback”** paragraph — *"`getServerCatalogView()` resolves a null backend result to the **empty view**"* — which states the opposite in bold. ⚠️ **This carried a line range until M257x iter-138, by then the `store.js`/beacon write-path blockquote.** Correct when written at `f8be5a1` (iter-108), rotted **+14** across files; named, not re-pinned.
        `getServerCatalogView()` resolves a null backend result to the **empty view**
        (`serverTenant.js:115-145` — *"the cutover is intentional, not reversible-on-error"*). A demo grid
        renders cards **only** because the rext demo-patch `demo-stack/patches/academy-fs-published-fallback`
        *restores* that removed fallback on the demo's ephemeral clone; if the patch is refused, the grid is
        empty. The demo CTA is a real `/courses/<slug>` link (see
        [`../ops/demo/content-stories-routes.md`](../ops/demo/content-stories-routes.md)).
    *   `cmd/academyImport` — idempotent, resumable catalog/metadata/**bodies** importer from a manifest JSON
        (`--manifest`, `--content-root`, `--checkpoint`), uploading covers/audio via the asset uploader.
    *   `cmd/coursebuilder-e2e`, `cmd/coursebuilder-liverun` — exercise the Course Builder → academy publish path.

### 2. Running in Docker
*   **Service Name**: none of its own — part of the `backend` service (`docker compose up -d backend`).

### 3. Testing
```bash
go test ./internal/academy/...                          # domain
go test ./internal/web/backend/graphql/graph/...        # resolvers
```
Unit tests use an **in-memory SQLite (`enttest`) harness** (`sqlite_test.go`) — the academy Manager is pure Ent, so
no real Postgres is needed (the one Postgres-only feature, `SELECT … FOR UPDATE`, is `t.Skip`'d under SQLite).
Notable: `certificate_mint_test.go` (cert-id contract over 500 mints + PII rejection); `entitlement_parity_test.go`
(cross-repo parity with the FE `getEntitlement`/`serverTenant.js` gates); `merge_test.go`/`time_test.go`/
`progress_aggregate_test.go` (monotonic-merge + keep-max-by-event-time semantics).

## Related Documentation

* [Ant Academy (frontend)](./ant-academy.md) — the Next.js/Expo reader/UI this domain serves
* [Backend (`app`)](./backend.md) — the monolith this domain lives in
* [Course Builder (`app` domain)](./coursebuilder.md) — publishes generated courses into the `academy_*` catalog
* [AI Readiness](./ai-readiness.md) — consumes the path embeddings for its recommendation engine
* [`../ops/demo/content-stories-routes.md`](../ops/demo/content-stories-routes.md) — the demo FS-catalog caveat
