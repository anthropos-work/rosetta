# Ask Engine — "Talk to Data" (`app` domain)

> **Not a standalone service — a domain of the `app` monolith** (the service the platform calls "backend").
> No separate container, repo, or subgraph. Code lives in `app/internal/askengine/` (the engine) +
> `app/internal/web/backend/ask/` (the HTTP/SSE handler + agentic loop).

## Role & Responsibility

*   **Primary Goal**: The core of **"Talk to Data"** — a natural-language business-analytics **copilot** embedded in
    the Anthropos Workforce/Hiring dashboard. A non-technical stakeholder asks a question in any language; an
    **agentic LLM** (Claude Sonnet 4.6 on AWS Bedrock) writes SQL, runs it inside an **org-scoped read-only SQL
    sandbox**, and explains the results — never exposing raw schema or another tenant's data.
*   **Key Functions**:
    *   **Agentic tool-calling loop** — streams a Bedrock/Claude turn, executes any tool the model emits, feeds
        results back, repeats until the model stops (max 15 turns).
    *   **Org-scoped SQL sandbox** — validates LLM-authored SQL (SELECT-only) and rewraps it in a CTE prelude that
        shadows every allowed table with an organization-filtered view, so a forgotten `WHERE` clause can never leak
        cross-tenant data. Enforced by a **table allowlist registry** (~60 tables).
    *   **Tool registry** — exposes `query_postgres` (SQL fetch), `render_chart` (inline chart spec), and — in author
        mode — `course_builder` (turn an insight into a training chapter, via [Course Builder](./coursebuilder.md)).
    *   **System-prompt construction** — a large cacheable prompt (an embedded `rules.md`, ~146 KB of schema + golden
        queries) plus per-org / per-user dynamic blocks (custom rules, cross-org auto-distilled rules, RAG golden
        examples, current-user identity).
    *   **Follow-up suggestions** — a streaming filter that strips a `<followups>` block from the visible text and
        surfaces 2–4 suggested next questions.
    *   **Result cross-validation** — warns the LLM about impossible values (percent out of range, negative counts,
        member-count > headcount, duplicate `user_id` = missing GROUP BY).

## Architecture & Code Map

*   **Codebase**: `app/internal/askengine/` (pure SQL/prompt/LLM layers) + `app/internal/web/backend/ask/` (the Echo
    handler, agentic loop, conversation store). Go, part of the `app` monolith.
*   **Database**: PostgreSQL. Reads run against a **dedicated, small read pool** (`COPILOT_DB_CONN` DSN,
    `askEngineMaxConns=6`) so interactive queries can't be starved by heavy dashboard analytics; each query runs in a
    **read-only tx with `SET LOCAL statement_timeout=10s`**. Own persistence tables in the **`public` schema**:
    `ask_conversations`, `ask_messages` (JSONB content blocks), `ask_query_examples` (RAG golden question→SQL pairs
    with a **pgvector** embedding, cosine retrieval), `ask_query_lessons` (captured failed/corrected queries),
    `ask_auto_rules` (cross-org auto-distilled ruleset).
*   **Key files**:
    *   `bedrock.go` — the LLM adapter. `BedrockClient` wraps `anthropic-sdk-go` configured for **AWS Bedrock**;
        model **`eu.anthropic.claude-sonnet-4-6`** (`DefaultModelID`), region **`eu-west-1`**, `Temperature:0`, an
        ephemeral prompt-cache breakpoint on the system prompt. `stripAnthropicAuthMiddleware` forces the SigV4 path.
    *   `sandbox.go` — **the security boundary**: `ValidateSQL` (must start SELECT/WITH; rejects DML, system
        catalogs, `public.` prefixes, `SET`, multi-statements), `WrapQuery` (defends four documented bypass vectors),
        `BuildCTEs` (shadows each referenced table with an org-scoped CTE via `OrgStrategy` Direct/Indirect/Global;
        `$1`=orgID, `$2`=callerUserID for self-scoped tables like `profile_histories`).
    *   `registry.go` — the `TableRegistry` allowlist (~60 `TableDef`s across `public`, `jobsimulation.*`,
        `directus.*`, and deprecated `skiller.*`/`skillpath.*` transition aliases that resolve to `public`).
    *   `executor.go` — `Engine.Execute` (the read-only per-query executor; `MaxInlineRows=200`, `MaxCellLength=400`)
        + `crossValidate`.
    *   `prompt.go` — `BuildSystemPrompt` (embeds `rules.md`; flips product framing Workforce↔Hiring; orders blocks
        to maximize the cacheable prefix).
    *   `followups.go` — the streaming `<followups>` extractor. `coursebuilder_tool.go` + `coursebuilder_rules.md` —
        the author-mode tool (its `brief` must be an **AGGREGATE** with no per-row PII).

## Interface Discovery

*   **How to find the API**: **HTTP (Echo) + SSE only — NOT GraphQL and NOT Connect-RPC.** The engine is driven by
    the sibling package `internal/web/backend/ask/handler.go`; routes are mounted under the **`/ask`** Echo group
    (CORS + Clerk `authn`; the group **skips OpenAPI validation** because SSE isn't swagger-expressible).
*   **Key routes**: `POST /ask/stream` (the **SSE** agentic generation), `GET`/`POST /ask/conversations`,
    `GET /ask/conversations/:id`, `GET /ask/conversations/:id/stream` (SSE reattach to an in-flight run),
    `DELETE /ask/conversations/:id`, `POST /ask/messages/:messageId/feedback`, and the admin auto-rules routes
    (`POST /ask/admin/auto-rules/distill`, `GET`/`activate`/`DELETE`).
*   **The agentic loop** (`handler.go:runAgenticLoop`, `maxAgenticIterations=15`, `loopTimeout=10min`) runs in a
    **detached goroutine** (`context.WithoutCancel` + timeout) so a client disconnect doesn't abort generation; a
    per-conversation `streamRegistry` lets the client reattach. Per iteration: `bedrock.Stream` → accumulate blocks →
    if `stop_reason != tool_use` finish, else `executeTool` each tool (`query_postgres`→`Engine.Execute`,
    `render_chart`, `course_builder`) and feed back the result. Wire results are sanitized (`publicToolResult` strips
    raw error text).
*   **Authorization**: the `/ask` group mounts Clerk auth; **org-admin is resolved per-turn** via
    `OrgCheckFeaturePermission`, gating self-scoped table access + author mode. Tenant scope comes from the `authn`
    user → org.
*   **Dependencies**:
    *   **Upstream (callers)**: the Next.js Workforce/Hiring "Talk to Data" UI (via the `/ask/*` SSE + REST
        endpoints).
    *   **Downstream**: **AWS Bedrock** (Claude Sonnet 4.6, eu-west-1); **PostgreSQL** (read-only, org-scoped); the
        shared **`ai`** library (`CreateEmbeddings`, text-embedding-3-small, for RAG golden-example retrieval);
        [`internal/coursebuilder`](./coursebuilder.md) (author mode); `internal/organization`
        (`GetTalkToDataCustomRules`); `internal/authorization`.

## Local Development

### 1. Running Standalone
*   **Prerequisites**: runs **as part of `app`** (no dedicated `cmd/`); a failed Bedrock init **disables** Talk-to-Data
    but doesn't crash `app`. Needs AWS Bedrock access (IAM via the default credential chain) with Claude Sonnet 4.6
    enabled in **eu-west-1**, a Postgres with the platform schema, and an embeddings provider for the shared `ai`
    client.
*   **Command**: `go run .` in the app repo, or `make up` (ships inside the `backend` container).
*   **Key env vars**: `ASK_MODEL_ID` (default `eu.anthropic.claude-sonnet-4-6`), `AWS_REGION` (default `eu-west-1`),
    `COPILOT_DB_CONN` (the read-pool DSN), the standard AWS creds, and `AZURE_OPENAI_ENDPOINT_URL`/`AZURE_OPENAI_KEY`
    (embeddings). Per-org config is DB-resident, not env: the `talk_to_data_custom_rules` org setting.

### 2. Running in Docker
*   **Service Name**: none of its own — part of the `backend` service (`docker compose up -d backend`).

### 3. Testing
```bash
go test ./internal/askengine/...
```
*   `sandbox_test.go` (49 tests, the bulk) — adversarial/security: each test names the specific SQL-injection or
    CTE-bypass vector it closes.
*   `executor_test.go` (13) — `crossValidate` sanity checks + cell normalization.
*   `prompt_test.go` (10) — workforce-vs-hiring framing, dynamic blocks, and **`RulesBlockIsIdentical`** (prompt-cache
    prefix stability).
*   `followups_test.go` (14) — the streaming `<followups>` extractor edge cases.
No real Bedrock is hit in unit tests — the SQL/prompt/followups layers are pure and table-driven. (The handler
package `internal/web/backend/ask/` carries its own tests: sanitize, stream_registry, examples, lessons, distill,
feedback, store, embed.)

## Status & Notes

*   **Shipped v1.267.0 (2026-05-07)** — the initial "Talk to Data" with SSE streaming + conversation management + the
    Bedrock client + the SQL sandbox. **Author mode** (`course_builder` tool) landed **v1.340.0 (2026-07-17)**.

## Related Documentation

* [Backend (`app`)](./backend.md) — the monolith Ask Engine is a domain of
* [Course Builder (`app` domain)](./coursebuilder.md) — the author-mode target; shares the Bedrock transport
* [AI Architecture](../architecture/ai_architecture.md) — Bedrock routing, embeddings, model inventory
* [Security & Compliance](../architecture/security_compliance.md) — the multi-tenant isolation the SQL sandbox enforces
