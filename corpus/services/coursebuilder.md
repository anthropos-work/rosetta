# Course Builder (`app` domain)

> **Not a standalone service — a domain of the `app` monolith** (the service the platform calls "backend").
> Course Builder ships inside `app`; there is no separate container, repo, or subgraph. Code lives in
> `app/internal/coursebuilder/` (the engine) + `app/internal/web/backend/coursebuilder/` (the HTTP/SSE boundary).

## Role & Responsibility

*   **Primary Goal**: The in-process **author → benchmark → refine** AI pipeline that turns a customer-supplied
    source (pasted text, uploaded files, or a workforce-profile brief) into a shippable **Anthropos AI Academy**
    chapter or multi-chapter skill-path — in the target language, gated at a cross-model benchmark score **≥ 90**
    before it can be published. It is the "author your own custom training" surface for enterprise content teams;
    the content it produces is schema-identical to team-authored Academy content and lands in the same
    `academy_*` tables the [academy backend](./academy-backend.md) serves.
*   **Key Functions**:
    *   **Author** a chapter as Academy-schema-identical JSON (`{meta, modules, references}`) from a source
        (single-shot Opus 4.8 call; the earlier parallel-author pipeline was removed 2026-07-21).
    *   **Benchmark + refine** with a cross-model **Sonnet 4.6 grader** (0–100 score + violation list); a
        keep-best / rollback-on-regression loop of patches until the score reaches **≥ 90** or the budget exhausts.
    *   **Multi-chapter path planning** (Wave 24): a planner lays out an N-chapter path, each chapter authored by
        the same iterate-to-gate engine, then assembled into a `Path`.
    *   **Refine / live-steer / rebuild-consent**: surgical chapter-scoped edits, mid-build live steering, and a
        consent gate for rebuild-class instructions.
    *   **Publish** into the `academy_*` tables; **translate** (7 locales); AI **cover-image** generation.
    *   **Cost governance**: per-org **credit** debits (via the [credits ledger](./ai-labs.md)), daily/monthly COGS
        ceilings, a session-count rate cap, and prompt-cache accounting.

## Architecture & Code Map

*   **Codebase**: `app/internal/coursebuilder/` (the engine, ~100 Go files) + `app/internal/web/backend/coursebuilder/`
    (the HTTP/SSE boundary, credits, rate/COGS caps, publisher). In-repo docs are authoritative and worth reading
    first: `internal/coursebuilder/README.md`, `SPEC.md`, `GO-LIVE-RUNBOOK.md`, `STEERING_PLAN.md`,
    `MULTI-CORRECT-FEASIBILITY.md`.
*   **Language**: Go (1.26.4) — part of the `app` monolith.
*   **Database**: PostgreSQL, **`public` schema** (standard `app` Ent). Sole owned table: **`course_builder_sessions`**
    (Ent schema `internal/data/ent/schema/coursebuildersession.go`; migration `terraform/migrations/20260717151144.sql`,
    which also creates the `credit_transactions` + `organization_credits` tables). Key columns: `organization_id`,
    `user_id`, immutable `source_text`/`locale`/`depth`, `draft_json`, `status` (draft/preview/published/cancelled),
    `iterations`, `benchmark_score`, `cost_usd`, `usage_json`, `translations`, `checkpoint_json`, `published_hash`,
    `draft_version` (optimistic lock), `build_ref` (crash-refund key).
*   **Key engine files**: `service.go` (the state machine + `ProductionConfig(author, grader, translator)` wiring —
    `TestProductionConfig_WiresPlanner` locks the nil-planner regression), `refine.go` (the load-bearing ≥ 90 loop:
    keep-best + rollback), `author.go` (single-shot author), `benchmark.go`/`patch.go`/`patchgen.go` (grader +
    patch machinery), `planner.go`/`path.go` (multi-chapter), `checkpoint.go` (resume), `schema.go` (the Academy
    chapter shape + validators), `normalize.go` (widget normalization + XSS sanitization), `bedrock.go`/`model.go`/`usage.go`
    (LLM adapter + `MockClient` + cost formula), `embed.go` (`//go:embed assets/*.md` rubric), `imagegen/` (cover
    images).
*   **LLM usage — AWS Bedrock (eu-west-1)** via the shared `internal/askengine/bedrock.go` transport:
    *   **Author/patch model**: Opus 4.8 (`eu.anthropic.claude-opus-4-8`, env `CB_AUTHOR_MODEL`; streaming, no
        sampling params — Opus 4.8 rejects them — at 32 K max_tokens).
    *   **Grader model**: Sonnet 4.6 (`eu.anthropic.claude-sonnet-4-6`, env `CB_GRADER_MODEL`; deliberately a
        different model to avoid self-grading bias; `temperature=0`).
    *   **Cover images**: OpenAI `gpt-image-2` (`imagegen/openai.go`, separate key).
    *   **Prompt caching** (Wave 2b): static system prompts marked ephemeral cache-control (~85 % input-side saving
        on the static prefix), tracked via `CacheUsageTracker`.
*   **Job structure**: builds run on a **dedicated Asynq worker** (`internal/worker/worker.go` → `CoursebuilderSrv`,
    concurrency 20), task type `TypeCoursebuilderRun` on its own Redis queue (`MaxRetry(1)`, per-depth timeout,
    `Retention(6h)`). If enqueue fails the handler drives the build **in-process** as a fallback. `checkpoint_json`
    lets a deploy-interrupted multi-chapter build resume only the missing chapters.

## Interface Discovery

*   **How to find the API**: **HTTP + SSE only — there is NO GraphQL subgraph and NO Connect-RPC for Course
    Builder.** The routes are an Echo group mounted in `internal/web/backend/backend.go` under **`/coursebuilder`**,
    behind `cors + authn (Clerk JWT via colony/authn) + courseBuilderAccessGate` (**org-admin-gated**). Route table:
    `internal/web/backend/coursebuilder/handler.go:Register`. The whole group is **unmounted** when the Bedrock-backed
    `Service` is nil (missing AWS creds) — no half-working surface.
*   **Key routes**: `POST /coursebuilder/sessions` (+ `/sessions/mixed`, `/sessions/upload`), `GET /sessions` +
    `/sessions/:id`, **`POST /sessions/:id/messages`** (the **SSE** build/refine stream), `/sessions/:id/{queue,steer,
    cancel,publish,unpublish,duplicate,translate,cover}`, `PATCH /sessions/:id/draft`, `GET /sessions/:id/published-diff`,
    `DELETE /sessions/:id`, `GET /people`, `GET /tags`.
*   **SSE wire contract** (`event:` name / `data:` JSON): `session`, `stage`, `outline`, `progress`, `patch_applied`,
    `patch_skipped`, `preview_ready`, `draft_kept`, `translation_ready`, `rebuild_required`, `error`, `cost`, and an
    always-last `done`.
*   **Dependencies**:
    *   **Upstream (callers)**: the Next.js `/enterprise/coursebuilder` UI (via SSE); **Ask "author mode"**
        (`app/internal/askengine/coursebuilder_tool.go` → `ask.SetCourseBuildStarter`); the Asynq worker.
    *   **Downstream**: `internal/askengine` (Bedrock transport, shared with [Ask/Talk-to-Data](./askengine.md));
        AWS Bedrock (Opus 4.8 + Sonnet 4.6); OpenAI `gpt-image-2` (covers); `internal/credits`
        (debits/refunds — see [AI Labs + credits](./ai-labs.md)); Ent/Postgres (`course_builder_sessions` +
        `academy_*` on publish); colony pub/sub (author-lifecycle emails → messenger → Brevo); Redis (Asynq). It does
        **not** call the shared `ai` library, embeddings, taxonomy, or CMS.

## Local Development

### 1. Running Standalone
*   **Prerequisites**: it runs **as part of `app`** (no separate binary/container). Needs AWS creds resolvable by the
    default SDK chain (Bedrock, eu-west-1), Postgres (migrations applied), Redis (the Asynq worker), and a Clerk
    secret. The routes stay unmounted if Bedrock creds are absent.
*   **Command**: `go run .` in the app repo, or the platform `make up` (Course Builder ships inside the `backend`
    container).
*   **Key env vars**: `CB_AUTHOR_MODEL` (default `eu.anthropic.claude-opus-4-8`), `CB_GRADER_MODEL` (default
    `eu.anthropic.claude-sonnet-4-6`), `CB_IMAGE_MODEL` (default `gpt-image-2`), `COURSEBUILDER_OPENAI_IMAGE_KEY`,
    `COURSEBUILDER_PLANNER_ENABLED` (multi-chapter kill-switch), `CB_SOURCE_DISTILL`, `COURSEBUILDER_EMAILS_ENABLED`,
    `COURSEBUILDER_MAX_MONTHLY_COGS_USD` (default **500**, the primary per-org ceiling), `COURSEBUILDER_MAX_DAILY_COGS_USD`
    (default 0 = off), `AWS_REGION`, `CLERK_SECRET_KEY`. Cost/rate: session cap `DefaultSessionsPerOrgPerDay=50`
    (code constant); credits `course.build`=**5/chapter**, `course.refine`=**1**, `course.translate`=**1/locale**.
*   **Operator-only cmd tools** (not in CI, cost money): `cmd/coursebuilder-liverun` (author one course against real
    Bedrock from a Markdown source; writes `chapter.json` + `run-report.json`) and `cmd/coursebuilder-e2e` (full
    seed→author→benchmark→publish proof against a **throwaway** Docker Postgres; `assertNotSharedDSN` refuses any
    shared/prod DSN).

### 2. Running in Docker
*   **Service Name**: none of its own — it is part of the `backend` service (`docker compose up -d backend`).

### 3. Testing
```bash
go test ./internal/coursebuilder/...           # ~55 test files, ~2 s, mock-backed
go test ./internal/web/backend/coursebuilder/... # ~25 boundary test files
```
The deterministic `MockClient` (FIFO scripted responses, usage/cache accounting) stands in for the Bedrock
`ModelClient` — no real LLM in unit tests. Coverage spans author parse/validate/retry, grader scoring, patch
apply round-trips, refine convergence (keep-best + rollback), locale threading, multi-chapter build/publish, XSS
sanitization, the credit/COGS/rate gates, and `-race`. Live-Bedrock proofs live in the two operator-only `cmd/`
tools.

## Status & Notes

*   **Not behind a global feature flag** — GA to every organization, **org-admin-gated** (`courseBuilderAccessGate`).
    The only gate is graceful degradation: routes stay unmounted if Bedrock creds are absent. Per-sub-feature
    kill-switches: `COURSEBUILDER_PLANNER_ENABLED`, `CB_SOURCE_DISTILL`, `COURSEBUILDER_EMAILS_ENABLED`.
*   **Recently added + heavily iterated** — developed through numbered Waves 1→24 (142 changelog lines); current app
    `v1.351.1` (2026-07). The parallel-author pipeline was built then deleted (2026-07-21) — single-shot `Author` is
    the only path now.

## Related Documentation

* [Backend (`app`)](./backend.md) — the monolith Course Builder is a domain of
* [Academy backend (`app` domain)](./academy-backend.md) — the `academy_*` tables Course Builder publishes into
* [AI Labs + credits](./ai-labs.md) — the credit ledger Course Builder debits (its only credits consumer today)
* [Ask engine / Talk-to-Data](./askengine.md) — shares the Bedrock transport; Ask "author mode" invokes Course Builder
* [AI Architecture](../architecture/ai_architecture.md) — Bedrock routing, model inventory
