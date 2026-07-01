# Customer-Facing API + MCP Server — Execution Roadmap

Author: developer-kit (exploration + drafting)
Status: DRAFT — sequential milestones, execute one at a time.
Scope: first customer-facing API layer for the Anthropos platform (REST + GraphQL), an MCP server exposing the same query tools, a Talk-to-Data–style natural-language endpoint, an Integrations page in Enterprise Settings, and developer docs to connect Claude / ChatGPT in minutes.

---

## 0. Vision & framing

**Goal.** Any Anthropos customer admin can, in under 5 minutes:
1. Open `Enterprise Settings → Integrations`, mint an API key (or a personal access token), and
2. Point Claude Desktop / Claude API / ChatGPT / a script at either
   - the **Public API** (REST or GraphQL) — programmatic reads of the org's data (skills, people, roles, simulations, assessments, reports…), or
   - the **Anthropos MCP Server** — a hosted MCP endpoint exposing the same query surface as tools, plus a natural-language `ask` tool that mirrors Talk-to-Data.

Everything the internal `Talk to Data` feature can answer, a customer can now answer programmatically — scoped hard to their own org, over authenticated, audited, rate-limited channels.

**Design tenets.**
- **Read-first**. v1 is read-only across the board. Writes come only when we have a signed contract for them (out of scope for this roadmap).
- **Federation-native**. Do not re-implement domain logic. The customer API is a thin, hardened facade over the existing Cosmo Router supergraph and its 5 subgraphs (`backend`, `cms`, `skiller`, `jobsimulation`, `skillpath`).
- **Tenant isolation by construction**. Every request path resolves an `organization_id` at the edge, injects it as a synthetic Clerk-equivalent context, and every downstream query is Ent-policy-filtered as it is today. There is no path where a caller can omit the org scope.
- **One brain, two mouths (three, counting the UI)**. REST, GraphQL, and MCP all delegate to a single internal `query-service` façade so behavior is identical and evolves together.
- **Bring-your-own-agent from day one**. The MCP endpoint is remote (HTTP+SSE / Streamable HTTP), works with Claude Desktop, `claude` CLI, ChatGPT (once ChatGPT ships MCP client support broadly), and the Anthropic Agent SDK.
- **Do not conflate service names with content.** The customer's `skills`, `roles`, `skill-paths`, `simulations`, `assessments`, `sessions` come from the right owner service (see the content-vs-runtime split in `corpus/architecture/architecture_overview.md`). The API surface hides the split behind one clean product noun set.

**Non-goals for v1.**
- No write endpoints.
- No BYO-model / BYO-Bedrock for the `ask` endpoint.
- No public "webhooks out" surface (events for customer subscribers) — separate track.
- No fine-grained ABAC exposed to customers (they get a single "read all my org data" scope in v1; scopes granularize in v2).

**Grounding — what exists today (from the corpus).**
| Concern | Today | Reference |
|---|---|---|
| Public API | none (only internal) | `corpus/services/backend.md:96-99` |
| GraphQL gateway | Cosmo Router `:5050/graphql`, 5 subgraphs, static composition | `corpus/services/graphql-wundergraph.md` |
| Talk to Data | `app/internal/askengine/{bedrock,sandbox,executor,followups}.go` + `rules.md`; SSE at `/ask/stream`; own conversation table + rate-limit | `corpus/services/backend.md:42-47,86,98` |
| Auth | Clerk JWT → `colony/authn` → Sentinel Casbin | `corpus/services/clerk-integration.md`, `corpus/services/sentinel.md` |
| Tenancy | `organization_id` on every table + Ent auto-filter | `corpus/architecture/security_compliance.md:60-78` |
| Enterprise UI | `apps/web/enterprise/*`, Ant Design 6, GraphQL only | `corpus/services/next-web-app.md` |
| API keys / MCP / integrations | **do not exist** | — |

---

## Milestone list (sequential)

- **M1** — Discovery, contract & security review (docs-only gate)
- **M2** — API-key primitive: issue / rotate / revoke, hashed at rest, org-scoped
- **M3** — Public API gateway (`api.anthropos.work`) + REST read layer v1 (people, skills, roles, paths, simulations, sessions, assessments, reports)
- **M4** — Customer-facing GraphQL surface (persisted queries + the same read schema)
- **M5** — Rate limiting, quotas, audit log, error taxonomy, request tracing
- **M6** — MCP Server (remote, Streamable HTTP), exposing the read tools + resource templates
- **M7** — Natural-language `ask` endpoint (customer-safe generalization of `askengine`) — REST, GraphQL, and MCP tool
- **M8** — `Enterprise Settings → Integrations` UI: key CRUD, MCP enable, quickstart, live docs, usage panel
- **M9** — Developer docs site + Claude / ChatGPT / SDK quickstarts + reference SDKs (TS, Python)
- **M10** — Private beta → GA hardening: versioning policy, deprecation contract, SLOs, on-call runbook

Each milestone below has: **Goal · Scope in / out · Deliverables · Technical approach · Dependencies · Risks · Acceptance**.

---

## M1 — Discovery, contract & security review

**Goal.** Freeze the v1 product surface and its safety contract before any code lands. Turn "customers can query anything" into an enumerated schema + scope model that a security review can sign off on.

**Scope in.**
- Enumerate the v1 read surface (entities + fields) across all 5 subgraphs + `askengine`'s `rules.md` allowlist.
- Draft OpenAPI 3.1 + GraphQL SDL for the customer-facing schema (not the federated internal one).
- Decide the auth model: **API key** (org-scoped, service-to-service) and **Personal Access Token** (user-scoped, inherits the user's Clerk role); v1 ships **API key only**, PAT deferred to M10+.
- Decide the tenancy contract: every key belongs to exactly one Clerk `organization` (Anthropos internal `eid`); no cross-org keys.
- Threat model: key exfiltration, SSRF via `ask`, prompt injection through customer data, over-broad reads that hit the GB tables listed in `corpus/ops/db-access.md:32-38`.

**Scope out.** Any code changes, DB migrations, or UI work.

**Deliverables (docs).**
- `docs/api/surface-v1.md` — the v1 entity catalog with per-entity field list, filter set, and originating subgraph.
- `docs/api/security-model.md` — auth, tenancy, scopes, rotation, revocation, incident response.
- `docs/api/threat-model.md` — STRIDE table + mitigations, sign-off log.
- `docs/api/openapi.yaml` (draft) and `docs/api/schema.graphql` (draft).

**Technical approach.**
- Cross-walk the federated supergraph (`schemas/backend|cms|skiller|jobsimulation|skillpath.graphqls`) against a customer-friendly noun set:
  - `people` ← `backend.User`, `Membership`
  - `skills` / `roles` / `specializations` ← `skiller.Skill`, `skiller.JobRole` (public taxonomy + org overrides)
  - `skill-paths` ← `cms.SkillPath` (content) joined with `skillpath.SkillPathSession` (progress)
  - `simulations` ← `cms.Simulation` (blueprint) joined with `jobsimulation.SimulationSession` (runs)
  - `assessments` / `verified-skills` ← the skill-verification chain (`corpus/ops/demo/stories-spec.md`)
  - `reports` / `workforce-analytics` ← `backend/internal/workforce` (rolled up)
- Codify the **public/customer split** from `corpus/ops/db-access.md:44-55`: public taxonomy is exposed as global reference; org data is always org-filtered.
- For `askengine`, treat `internal/askengine/rules.md` as the customer-safe allowlist floor — anything not in it cannot be reached even by the natural-language endpoint.

**Dependencies.** Security review sign-off (must attend the threat-model meeting).

**Risks.**
- Federated schema drift — v2.6.6+ Cosmo Router bakes composition at build (`corpus/services/graphql-wundergraph.md:49-58`); need a policy that our customer schema pins subgraph versions independently.
- Cross-service joins (`skill-paths` = CMS content + Skillpath progress) become 2× RPC hops at the edge — capture latency budget now.

**Open questions (flag inline).**
- Do we want a single `read:all` scope in v1, or already carve `read:people`, `read:skills`, `read:activity`? Recommendation: single scope, note as v2 seam.
- PAT model: mint via Clerk custom claims, or a separate `personal_tokens` table? Recommendation: separate table with a Clerk `user_id` FK; do not overload Clerk.

**Acceptance.**
- Surface catalog reviewed by product + at least one design partner.
- Threat model signed off by security.
- OpenAPI + GraphQL drafts lint clean (`spectral`, `graphql-schema-linter`).
- No code committed.

---

## M2 — API-key primitive: issue / rotate / revoke

**Goal.** Introduce a first-class, org-scoped `api_key` in the `backend` service (`public` schema) that can be issued, listed, rotated, and revoked, and is the sole thing the future gateway trusts for customer traffic.

**Scope in.**
- Ent schema: `api_keys` on `backend` (`public` schema — same place as `users`, `organizations`, `memberships`).
- CLI + Connect-RPC methods (internal only, admin-gated) to issue/list/rotate/revoke.
- **Hashed at rest** (Argon2id or bcrypt) — plaintext returned once at creation only.
- Prefix + last-4 stored plaintext for display (e.g. `ant_live_ABCD…WXYZ`).
- `organization_id` FK is **required**; `created_by_user_id` FK; `scopes` string array (v1 always `["read:all"]`); `expires_at` optional; `revoked_at`.
- Sentinel policy: only Clerk `org:admin` in the owning org may CRUD their org's keys.

**Scope out.** Gateway edge validation (M3), rate-limit accounting (M5), UI (M8), PATs.

**Deliverables (code).**
- `app/internal/data/ent/schema/apikey.go` (Ent schema) + `atlas migrate diff`ed migration.
- `app/internal/apikey/` package: `manager.go`, `hash.go`, `format.go` (key prefix + secret encoding), `manager_test.go`.
- Connect-RPC: `backend.v1.ApiKeyService/{Create,List,Rotate,Revoke}` in `app/rpc.go`.
- Cobra subcommand `app apikey {create,list,rotate,revoke}` under `cmd/apikey`.
- Sentinel: add `p2` policy row `org:feature:integrations:apikey:manage` gated to `admin`; wire into `init_policy.sql` + a `local_superadmin_grants.sql` note.
- Redis Streams event: `backend` stream carries `ApiKeyCreated|Rotated|Revoked` for downstream audit consumers.

**Technical approach.**
- Key format: `ant_{env}_{22-char-b32}.{40-char-b32}` — `{22-char}` is the record id (short, indexable), `{40-char}` is the secret. Constant-time compare on lookup; index the record id.
- Hash: Argon2id (`golang.org/x/crypto/argon2`), memory=64MB, iters=3, parallelism=2. Cost-tune before ship.
- Rotation: creating a rotation returns a new secret and marks the old key `revoked_at = now() + grace` (default 24h grace, configurable per key).
- **Never log the plaintext.** Reuse the values-blind principle from `corpus/ops/secrets-spec.md`.

**Dependencies.** None. Isolated to `backend`.

**Risks.**
- If we ever leak a plaintext into `ai_usages` or a Sentry breadcrumb, it's game over. Add a `regexp` breadcrumb scrubber for `ant_live_.*\.` and unit-test it.
- Argon2 CPU cost on `create` — fine (rare op), unit-test that `verify` under load stays <5ms P99.

**Acceptance.**
- Create → list → verify (RPC) → rotate → revoke cycle covered by integration tests.
- Sentinel denies non-admin `Create`; permits admin.
- Argon2 hash time within the tuned envelope on the reference CI runner.
- Audit event published to `backend` stream and consumed cleanly in a fixture consumer.

---

## M3 — Public API gateway + REST read layer v1

**Goal.** Stand up `api.anthropos.work` — a new edge service that terminates customer API-key auth, resolves tenant scope, and translates every request into a Cosmo Router GraphQL call. Ship a REST v1 read surface over it.

**Scope in.**
- New repo (or new dir in `platform`) `apigw` — a Go service using colony, sitting in the compose file at `:8500` (dev) / behind ALB (prod), profile `graphql` peers.
- REST v1 endpoints for the M1 catalog (paginated `GET /v1/{people,skills,roles,skill-paths,simulations,sessions,assessments,reports}` + `GET /v1/{entity}/{id}`).
- Key resolution middleware: `Authorization: Bearer ant_live_…` → hash-verify against `api_keys` → hydrate `organization_id` → mint a **short-lived internal Clerk-shaped JWT** signed by an internal HS256 key that `colony/authn`'s `Dummy` provider (or a new `Internal` provider) accepts.
- Query translation: every REST resource maps to a **persisted GraphQL operation** compiled at build time from `docs/api/schema.graphql`; no ad-hoc GraphQL execution from customer input.
- Response shaping: JSON:API-lite (`data`, `meta.page`, `links.next`) — pick a shape, document it, stick to it.

**Scope out.** Rate limiting (M5), audit (M5), MCP (M6), NL `ask` (M7), UI (M8).

**Deliverables.**
- `apigw/` service — `main.go`, `internal/{auth,router,graphql,resp}/`, OpenAPI-generated handler stubs (using `oapi-codegen`).
- `apigw/openapi.yaml` — the runtime spec (frozen from M1 draft).
- Persisted GraphQL operations under `apigw/persisted/*.graphql` + a compile-time SHA registry.
- `platform/docker-compose.yml` entry, `platform/repos.yml` entry, Terraform ECS service block, ALB routing rule for `api.anthropos.work`.
- `authn` extension: an **`Internal` provider** that trusts JWTs signed with a shared secret (`APIGW_INTERNAL_JWT_KEY`), rotated separately; carries the same `eid`/`org`/`org_role` claims Clerk does. Documented in `corpus/architecture/shared_libraries.md`.

**Technical approach.**
- `apigw` **never** talks to Postgres directly. It only calls `http://graphql:8080/graphql` with the internal JWT — that path is already Ent-policy-filtered by `organization_id` in every subgraph. Tenancy is preserved by construction, not by hope.
- Persisted queries are **the only way** to reach the graph. This kills GraphQL introspection abuse and lets us reason about the exact fields customers can pull.
- Pagination: cursor-based (`opaque base64(json)` → `{after, limit}`), max page 200, default 50. Never offset (kills the DB on large orgs).
- Content type: `application/json` only for v1. No SSE except on the `ask` endpoint (M7).

**Dependencies.** M2 (`api_keys` exists).

**Risks.**
- **Internal-JWT provider is the weakest link.** If leaked, an attacker can impersonate any org. Mitigations: rotate every 24h automatically, sign in the gateway process only, never write the key to disk in prod (SecretsManager → memory), narrow the acceptance window to 5 min.
- Persisted-query compilation vs Cosmo static composition — the customer schema must not drift ahead of what the compose-time supergraph exposes. Add a CI check that every persisted op validates against the frozen supergraph.

**Open questions.**
- Do we host `apigw` in-cluster (ECS) or edge-first (Cloudflare Workers proxying to ECS)? Recommendation: ECS for v1, revisit edge for M10 if latency demands it.
- Do we expose `api.anthropos.work` at ALB or via Cosmo Router directly? Recommendation: ALB, distinct target group — Cosmo does not want customer traffic on its box.

**Acceptance.**
- `curl -H "Authorization: Bearer ant_live_..." https://api.anthropos.work/v1/people?limit=10` returns the caller's org's members only.
- Rotating a key immediately (within grace 0) blocks the old key on the next request.
- All persisted operations validate against the CI supergraph.
- No path in `apigw` executes a customer-supplied GraphQL string.

---

## M4 — Customer-facing GraphQL surface (persisted + ad-hoc)

**Goal.** Expose the same catalog as GraphQL at `https://api.anthropos.work/graphql`, so agent frameworks that speak GraphQL can drive it natively. Same auth path as M3.

**Scope in.**
- Add a `/graphql` handler to `apigw` that:
  - Accepts persisted-operation IDs (SHA-256 of the op) with variables — the fast path.
  - Accepts ad-hoc queries only if the request's key has an opt-in flag `allow_adhoc_graphql = true` (default off) — safety valve for power users.
- Query cost analyzer (`graphql-cost-analysis`-equivalent) — reject ops over a per-key budget.
- Depth limit (default 8), breadth limit, argument-value length limit.
- Introspection: **off** for keys without `allow_introspection` flag. Off by default.
- Same underlying persisted-query path as M3 — REST endpoints and named GraphQL ops are the same thing under the hood.

**Scope out.** Subscriptions (streaming), mutations (all write ops), file uploads.

**Deliverables.**
- `apigw/internal/graphql/handler.go`, cost-analyzer wiring, persisted-op registry.
- `docs/api/schema.graphql` — the frozen customer schema (a strict subset of the supergraph).
- CI: fuzzer that throws random ops at `apigw` and asserts we always 400/403 on out-of-catalog fields.

**Technical approach.** Reuse M3's internal JWT bridge; the only new thing is a second HTTP handler that forwards to Cosmo with the same JWT.

**Dependencies.** M3.

**Risks.**
- Ad-hoc GraphQL, even opted in, is a footgun. Keep the opt-in gated behind an org-level admin toggle + audit log entry per session.

**Acceptance.**
- Persisted `PeopleList` op works via `POST /graphql` with `{ opId, variables }`.
- Ad-hoc query rejected without the flag; accepted (and cost-limited) with it.
- Fuzzer produces 0 unintended-data-returned events over 10 000 attempts.

---

## M5 — Rate limiting, quotas, audit log, error taxonomy, tracing

**Goal.** Make the API production-safe. Every request is accounted, rate-limited, audited, and traceable end-to-end.

**Scope in.**
- **Rate limits.** Per-key sliding window (default 60 rpm / 1000 rph / 20k rpd — tunable per key). Enforced in `apigw` via Redis (INCR + TTL) — the same Redis cluster the platform already uses. Return `429` with `Retry-After` + `X-RateLimit-*` headers (limit, remaining, reset).
- **Concurrency limit.** Per-key inflight cap (default 8) — bulkhead against runaway agent loops.
- **Quotas.** Monthly request quotas (soft) + monthly `ask` token budget (hard) — the second is the pain point since `ask` is Bedrock-priced. Reuse the `app/internal/aiusage` ledger (`corpus/architecture/shared_libraries.md:119-125`) — the `Event_AiUsage` stream already lands cost events; add a `key_id` label.
- **Audit log.** Append-only `api_audit_events` table on `backend`: `{id, org_id, key_id, actor_user_id, ts, method, path, op_id, status, latency_ms, req_bytes, resp_bytes, sample_of_error, ip, ua}`. Writer is async via a Watermill publisher on the `backend` stream to avoid slowing the request path. Retention 90 days (matches `corpus/architecture/security_compliance.md:110-119`).
- **Error taxonomy.** Standard JSON error body `{error: {code, message, docs_url, request_id}}`; codes documented in `docs/api/errors.md`.
- **Tracing.** OTEL span per request; correlate through `apigw → cosmo → subgraph → sentinel`. `X-Request-ID` on every response.

**Scope out.** Billing surfaces (out of this roadmap).

**Deliverables.**
- `apigw/internal/ratelimit/`, `internal/audit/`, `internal/errors/`, OTEL wiring.
- `app/internal/data/ent/schema/api_audit_event.go` + migration.
- `docs/api/rate-limits.md`, `docs/api/errors.md`.
- Grafana / CloudWatch dashboard for `apigw`: rpm / p99 / error-rate / 429-rate per key.
- Sentry: gateway-project DSN separate from `app` to keep traffic clean.

**Technical approach.**
- Redis keys `rl:{key_id}:{window}` with TTL == window; single INCR per request. Fixed-window sliding is fine at v1 (revisit if bursty traffic hurts).
- Audit writer publishes to a `api_audit` Watermill topic on the existing Redis Streams infra; a consumer in `backend/internal/worker/` persists to `api_audit_events`. If Redis is unhealthy, `apigw` **degrades open on rate limit, closed on audit** (drops the request rather than losing an audit line).

**Dependencies.** M3.

**Risks.**
- Audit backpressure — if the consumer lags, `api_audit_events` will queue in Redis. Alarm on consumer lag > 30s.
- Cost of Argon2 verify + audit publish on every request. Cache the key hash → key_id resolution in-process with a 60s TTL; measure and hold P99 < 40ms for the auth+audit overhead.

**Acceptance.**
- Load test: 500 rpm per key sustained, 429 clean at limit + 1.
- Every request traceable end-to-end in one Sentry / OTEL view.
- 90-day audit retention verified with a Postgres partition drop job.

---

## M6 — MCP Server: remote Streamable-HTTP endpoint

**Goal.** Ship `mcp.anthropos.work` — a hosted MCP server that exposes the M3–M4 read catalog as MCP tools + resource templates, so Claude Desktop / Claude Code / any MCP-capable client can connect in minutes.

**Scope in.**
- Transport: **Streamable HTTP** (the current spec-recommended remote transport) + SSE fallback for older clients.
- Auth: same `Authorization: Bearer ant_live_…` API key as M3 (MCP allows arbitrary HTTP headers) plus an optional `?key=…` for clients that don't let you set headers (documented as inferior).
- Tools (v1):
  - `list_people`, `get_person`
  - `list_skills`, `get_skill`, `search_skills`
  - `list_roles`, `get_role`
  - `list_skill_paths`, `get_skill_path`, `list_skill_path_sessions`
  - `list_simulations`, `get_simulation`, `list_simulation_sessions`
  - `list_assessments`, `get_assessment`, `list_verified_skills`
  - `run_report` (canonical rollups: workforce_activity, coverage_by_role, growth_by_period)
  - `ask` — deferred to M7, but the tool slot is reserved
- Resources (v1): `person://{id}`, `skill://{id}`, `role://{id}`, `skill-path://{id}`, `simulation://{id}`, `session://{id}` — each resolves to a JSON blob (or a small markdown summary + JSON payload).
- Prompts (v1): a small set of curated prompts — e.g. `weekly-workforce-summary`, `role-skill-gap`, `hiring-funnel-status` — each is a saved prompt template that uses the tools above.
- Discovery: `GET /` on the MCP endpoint returns human-readable "how to add this to Claude Desktop / ChatGPT" instructions and a link to the docs.

**Scope out.** Client-side installers, non-remote (stdio) MCP variant, tool-generated writes.

**Deliverables.**
- New service `mcp-server/` — Go, uses the official `mcp-go` SDK (or the reference Anthropic MCP Go implementation if we adopt it), boots on `:8600` in compose, deployed as its own ECS service.
- Each tool is a thin wrapper that calls `apigw` internally (over an in-cluster HTTP client with a service-account token) — the MCP server is **not** a second query engine, it's a tool-shaped face on the same one.
- `docs/api/mcp.md` — tool catalog, resource URIs, prompt catalog, connect guide.
- Compose + ALB route for `mcp.anthropos.work`.

**Technical approach.**
- Because MCP is invoked by an LLM, tool schemas must be **narrow, typed, and unambiguous** — one filter per parameter, explicit enums, no free-form JSON. Write a per-tool `zod`/JSON-Schema definition and test it against a Claude Sonnet call from a fixture.
- The MCP server holds no data. It resolves the API key on connect, maps it to an `org_id`, and threads that into every tool call as a header to `apigw`.
- Streaming: `ask` returns a stream of MCP progress notifications; other tools return synchronously.

**Dependencies.** M3, M4, M5.

**Risks.**
- MCP spec churn — pin the version we implement in `docs/api/mcp.md`; add a CI job that pins the SDK version and re-tests the connect flow.
- Prompt injection through customer data pulled by a tool. Mitigations documented in `docs/api/prompt-injection.md`; do not concatenate untrusted content into agent instructions on the server side.

**Acceptance.**
- Adding the MCP URL + key to Claude Desktop's config surfaces every tool.
- A Claude Sonnet chat "how many verified Python skills across my org?" resolves via `list_verified_skills` filtered by skill = python.
- MCP-Inspector CLI shows all tools + resources + prompts pass validation.

---

## M7 — Natural-language `ask` endpoint (customer-safe generalization of askengine)

**Goal.** Ship the crown-jewel: a customer-facing NL query API + MCP tool that answers questions against the caller's org data, using the same technique as internal Talk to Data — but sandboxed, org-scoped, budget-capped, and safe for arbitrary customer prompts.

**Scope in.**
- New endpoint `POST /v1/ask` on `apigw` — accepts `{question, conversation_id?}`, streams NDJSON events (`token`, `tool_call`, `sql`, `followup`, `answer`, `done`, `error`).
- MCP tool `ask` — same contract, delivered as MCP progress notifications.
- Reuse of the internal `askengine` code:
  - `app/internal/askengine/bedrock.go` — the Bedrock (Anthropic) agentic loop with prompt caching. **Keep as-is.**
  - `app/internal/askengine/sandbox.go` — the SQL whitelist + read-only enforcement. **Extend** with a mandatory `WHERE organization_id = $1` predicate injection for every generated query. The injection must be part of the sandbox itself, not a linter — a rejected query must be un-rescuable by prompt-tuning.
  - `app/internal/askengine/rules.md` — extend with a `customer_safe = true` section listing the tables allowed in customer mode.
  - `app/internal/askengine/executor.go` — pass `organization_id` into query execution as a bind param.
- **Conversation table**: extend the existing `askengine` conversation table with `api_key_id` (nullable, non-null for API traffic) + `source` (`internal|api|mcp`).
- **Budget cap**: enforce per-request `max_tokens_in + max_tokens_out` and per-key monthly ceiling using `app/internal/aiusage`. Hard-reject on breach.

**Scope out.** Fine-tuning per-org models; retrieval over customer content beyond the M3 catalog.

**Deliverables.**
- `app/internal/askengine` extensions + migration adding `api_key_id`, `source` columns.
- `apigw/internal/ask/` handler; MCP tool wiring in `mcp-server/`.
- `docs/api/ask.md` — the endpoint contract, event schema, examples.
- Playground page under `apps/web/enterprise/settings/integrations/ask-playground` (v0 — send a prompt, see the stream).

**Technical approach.**
- The **only** way `ask` can reach the DB is through the sandbox with the org predicate injected. Add a unit test that fuzzes 1000 attempted SQL bypasses (`;`-terminated second statement, comment-out, `UNION` against pg_catalog, etc.) and asserts every one is rejected. The test set becomes a regression fixture.
- Prompt-caching: reuse the `bedrock.go` cache. Cache keys must include the `organization_id` so cross-tenant cache poisoning is impossible by construction.
- Follow-ups (`followups.go`) return as `followup` events in the NDJSON stream — the UI + agents can offer them as one-click continuations.

**Dependencies.** M3, M5, M6, security review sign-off on the sandbox extension.

**Risks.**
- **Highest-risk milestone.** SQL generation over customer data is a bounty-magnet. Mandatory: threat model refresh, external pentest before beta, hard budget on per-question cost.
- Bedrock cost — an unbounded agentic loop can blow through budget fast. Cap the loop at 8 tool turns (already the internal default per `corpus/architecture/shared_libraries.md:131-132`) and re-verify.

**Open questions.**
- Do we support conversation history across requests in v1, or is every `ask` stateless? Recommendation: stateless in v1 (customer passes their own `history`) — simpler and lets the customer keep memory outside our system if they want.
- Should `ask` be able to call the M3 catalog as tools (self-hosted MCP-style)? Recommendation: yes, but only via the vetted tool list — the model can call `list_people`, `get_report`, etc., as intermediate steps. This is a productive convergence of M6 + M7.

**Acceptance.**
- `curl -N` streams tokens end-to-end.
- SQL-bypass fuzzer: 1000/1000 rejections.
- Org-scope test: two keys from two orgs asking the same question return data from their own org only.
- Bedrock per-question cost < $0.10 P50, < $0.50 P99 on the fixture set; monthly per-key hard cap enforced.

---

## M8 — `Enterprise Settings → Integrations` UI

**Goal.** In `apps/web`, ship the customer-facing surface where admins discover, activate, and manage the API + MCP.

**Scope in.**
- New route `apps/web/src/app/enterprise/settings/integrations/` with sub-pages:
  - **Overview** — what the API + MCP are, in 60 seconds; two big CTAs (Create API key / Enable MCP).
  - **API Keys** — list, create, rotate, revoke; show `prefix + last4`, created-by, last-used-at (from audit), scope, expiry; **plaintext shown once at creation**, with a "download env" affordance.
  - **MCP Server** — one toggle to enable/disable at org level (so an admin can shut off MCP globally without touching keys), the endpoint URL, and per-client connect snippets (Claude Desktop config JSON, `claude` CLI, ChatGPT).
  - **Usage** — chart of requests / errors / cost this month, per key, from `api_audit_events` + `ai_usages`. Read-only rollups.
  - **Docs** — inline links to the developer docs (M9) with the org's example key embedded.
  - **Quickstart** — a 3-step banner: (1) create key, (2) copy this snippet, (3) run it, see it work; shown until the org has made ≥ 1 successful API call.
- Sentinel policy: only `org:admin` can open the Integrations page (mirrors existing `/enterprise/settings` gating).
- All data via the federated GraphQL — the customer-facing UI itself must not use the new customer API (avoid the loop where our own UI depends on our public throttles).

**Scope out.** PAT UI (deferred), team/user-level scopes (deferred), webhooks/events (out of roadmap).

**Deliverables.**
- Next.js routes + Ant Design 6 pages in `apps/web/src/app/enterprise/settings/integrations/`.
- New GraphQL ops in `next-web-app/packages/graphql/src/query/integrations/*.graphql` + generated hooks.
- Backend GraphQL: extend the `backend` subgraph with `orgApiKeys`, `createApiKey`, `rotateApiKey`, `revokeApiKey`, `orgMcpSettings`, `setOrgMcpEnabled`, `orgApiUsage`.
- Storybook stories for the new components; Playwright E2E: create key → hit `/v1/people` with it → see it in Usage.

**Technical approach.**
- Mirror the existing `/enterprise/*` UX patterns already in `apps/web` — TanStack React Query + `useGraphql`, `<Protect>`-equivalent org-admin gate.
- **Plaintext key display** — one-time modal with a copy button, big "you will not see this again" warning, hash-verify test button.
- Endpoint URLs shown per environment (dev / staging / prod) — the UI reads these from `NEXT_PUBLIC_API_BASE_URL` and `NEXT_PUBLIC_MCP_BASE_URL`.

**Dependencies.** M2, M3, M5, M6.

**Risks.**
- Copy-to-clipboard on plaintext key on non-HTTPS localhost silently no-ops — add a fallback and warn.
- Admins mis-reading the "usage" chart as billing — label carefully; link out to Billing.

**Acceptance.**
- Playwright: fresh admin can create a key, run a `curl`, see a green check in Usage within 30s.
- The Integrations page renders empty-state correctly for a brand-new org.
- Non-admin gets 403 on all four `orgApiKey*` mutations.

---

## M9 — Developer docs + Claude / ChatGPT / SDK quickstarts

**Goal.** A public documentation site sufficient for a stranger to go from "signed-up admin" to "Claude Desktop answering questions about my company data" in under 5 minutes.

**Scope in.**
- `docs.anthropos.work/api` (or `/developers`) site. Static generator TBD (Docusaurus, Mintlify, or Nextra — pick one, prefer Mintlify for API docs polish; **open question**).
- Sections:
  - **Get started** — 5-minute quickstart with tabs for `curl` / TypeScript / Python / Claude Desktop.
  - **Authentication** — API keys, headers, rotation, revocation.
  - **REST reference** — auto-generated from `docs/api/openapi.yaml`.
  - **GraphQL reference** — auto-generated from `docs/api/schema.graphql`.
  - **MCP reference** — tool + resource + prompt catalog from M6.
  - **Ask API** — the NL endpoint, event schema, examples.
  - **Rate limits & errors** — from M5.
  - **Recipes** — "connect Claude Desktop", "connect ChatGPT (once MCP client GA)", "use with Claude API + Agent SDK", "use with the Anthropic Python SDK", "build a weekly digest agent".
  - **Changelog** — every schema change surfaces here.
  - **Status** — link to statuspage.
- Reference SDKs:
  - **TypeScript** — `@anthropos/api` — thin OpenAPI-generated client, works in Node ≥ 20 and browsers; MCP helper.
  - **Python** — `anthropos-api` — same shape, ships to PyPI.
  - Both include an `AnthroposClient` with `.people`, `.skills`, `.ask()` async stream.

**Scope out.** Localized docs (v2).

**Deliverables.**
- `docs-site/` in this repo (or split repo, TBD).
- `sdk/ts/`, `sdk/py/` starter packages.
- CI: docs preview per PR; SDK codegen from `openapi.yaml` on tag.

**Technical approach.**
- Generate REST reference straight from OpenAPI; do not hand-write it.
- Each recipe has a working, tested code snippet — CI runs the snippets against a demo org (`stack-demo/`) nightly and posts failures to Slack.

**Dependencies.** M3, M4, M6, M7 (schemas stable).

**Risks.**
- Docs drift is the #1 way this program dies. Auto-generate everything derivable; hand-write only intent, not shape.
- Claude Desktop / ChatGPT config UI can and will change — screenshots go stale. Prefer copy-pasteable JSON snippets over screenshots.

**Acceptance.**
- New Anthropic engineer, no prior context, can complete the 5-minute quickstart end-to-end.
- SDKs pass smoke tests against `stack-dev`.
- All recipes green in nightly CI for 7 consecutive days before GA.

---

## M10 — Private beta → GA hardening

**Goal.** Ship it — to 3 design partners first, then to every enterprise customer — with the runbook a real product needs.

**Scope in.**
- **Versioning policy** — `/v1/`, `/v2/` at the URL. Non-breaking changes only inside `v1`. Breaking-changes land in `v2` with 12-month `v1` support. Documented in `docs/api/versioning.md`.
- **Deprecation contract** — `Sunset` header, deprecation notice in changelog + email to key owners; ≥ 90 days before removal.
- **SLOs** — 99.9% availability (30d), P50 < 150ms, P99 < 1200ms for non-`ask` reads; `ask` P95 < 8s time-to-first-token, hard 60s wall.
- **On-call runbook** — `docs/api/runbook.md` under `corpus/ops/` conventions: key-leak response, gateway down, sandbox bypass discovery, rate-limit tuning, incident comms template.
- **Pentest** — third-party pentest scoped to `apigw` + `mcp-server` + `/v1/ask`. Findings triaged before GA.
- **Compliance** — DPA amendment for API traffic; sub-processor list update if we adopt any new SaaS (statuspage, docs host); audit-log retention aligned to DPA v1.4 in `corpus/architecture/security_compliance.md:135-149`.
- **Design-partner beta** — 3 customers × 4 weeks; weekly office hours; issue-tracker access.
- **GA gate** — signed pentest report + zero P0 in 2 weeks + all runbook steps rehearsed once in a game-day.

**Scope out.** v2 scope planning (separate roadmap).

**Deliverables.**
- `docs/api/versioning.md`, `docs/api/runbook.md`, statuspage set up, pentest report archived.
- Design-partner postmortem doc.
- Public launch blog post + Loom.

**Dependencies.** M1–M9 complete.

**Risks.**
- Pentest finds a sandbox bypass at week -2. Contingency: hold GA; do not ship.
- Bedrock quota — as `ask` scales, region quota may be the bottleneck. Pre-file AWS increases; instrument `throttling` alarms.

**Acceptance.**
- All SLOs hit for the 4-week beta.
- Zero P0 open at GA gate.
- ≥ 1 design partner has a real automation (e.g. weekly workforce digest via Claude + MCP) in production.

---

## Cross-cutting notes

**Rosetta corpus updates** (part of every milestone, not a separate one).
Each milestone that adds a service or a durable operating concern MUST update the corpus alongside code:
- `corpus/services/apigw.md`, `corpus/services/mcp-server.md` — new service docs (M3 / M6).
- `corpus/architecture/architecture_overview.md` — mermaid diagram add `apigw` + `mcp-server` at the edge.
- `corpus/architecture/dependency_map.md` — apigw depends on Cosmo Router; mcp-server depends on apigw.
- `corpus/architecture/security_compliance.md` — add "Customer API" section (auth + tenancy + audit).
- `corpus/ops/api-runbook.md` (M10) — key rotation, rate-limit tuning, sandbox-bypass response.
- Skill: `/update-knowledge` after each milestone close.

**Test topology.**
- Unit tests per package (Go standard).
- Integration tests: extend the demo stack (`/demo-up`) with an `apigw` + `mcp-server` service; add Playwright + a `mcp-inspector`-based test.
- Alignment DNA (`corpus/architecture/alignment_testing.md`) — new gene set for "Customer API contract" so future refactors don't silently regress the surface.

**Naming & branding.**
- Product name: "Anthropos API" (REST + GraphQL) + "Anthropos MCP".
- Env prefixes on keys: `ant_live_…`, `ant_test_…`.
- Base URLs: `api.anthropos.work`, `mcp.anthropos.work`, `docs.anthropos.work`.

**Open questions to resolve before M1 closes.**
1. Docs generator: Mintlify vs Docusaurus vs Nextra.
2. Do PATs ship in v1 (M8) or later (M10+)? Recommendation: later.
3. Should the customer API be behind Cloudflare, or ALB-direct? Recommendation: ALB-direct for v1, Cloudflare in front once traffic warrants.
4. `ask` conversation memory: stateless (customer manages) or stateful (we manage)? Recommendation: stateless in v1.
5. Do we allow ad-hoc GraphQL by default for design partners, or persisted-only? Recommendation: persisted-only, ad-hoc gated per-key.
6. MCP transport: Streamable HTTP only, or add SSE fallback? Recommendation: both; deprecate SSE once clients migrate.
7. Repo layout: new `apigw/` and `mcp-server/` sibling repos (matching `platform/repos.yml` pattern), or a monorepo under a new `platform-api/`? Recommendation: sibling repos, consistent with the current service topology.

---

## Milestone dependency graph

```
M1 (design)  ──►  M2 (api_key)  ──►  M3 (REST gateway) ──►  M4 (GraphQL surface)
                                        │                          │
                                        └──►  M5 (limits+audit) ◄──┘
                                                    │
                                        ┌───────────┼───────────┐
                                        ▼           ▼           ▼
                                       M6 (MCP)    M7 (ask)    M8 (UI)
                                        │           │           │
                                        └────►  M9 (docs+SDKs)  ◄
                                                    │
                                                    ▼
                                                M10 (beta→GA)
```

M6, M7, M8 can proceed in parallel after M5. M9 depends on the schemas being stable across M3/M4/M6/M7. M10 depends on everything.
