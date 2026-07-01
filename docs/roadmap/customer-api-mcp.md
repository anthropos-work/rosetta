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
- **Auth-layer independence (HARD CONSTRAINT — Stefano).** The public API + MCP + Integrations UI depend **only** on a stable internal **identity/principal contract** — `{organization_id, actor (user_id | service_principal_id), scopes, roles, session_id, issued_at, expires_at}` — populated by a thin **`IdentityProvider` adapter**. Clerk is *one* adapter today; it must be swappable (Auth0, WorkOS, self-hosted OIDC, or none) without touching a single line of API/MCP/UI logic. **No milestone below may reference Clerk claims, Clerk sessions, or Clerk SDKs directly** — the only place the word "Clerk" is allowed is inside the adapter package. When we drop Clerk, we replace the adapter and everything above it keeps working.
- **Tenant isolation by construction**. Every request path resolves an `organization_id` at the edge from the internal principal (never from a provider-specific claim), and every downstream query is Ent-policy-filtered as it is today. There is no path where a caller can omit the org scope.
- **One brain, two mouths (three, counting the UI)**. REST, GraphQL, and MCP all delegate to a single internal `query-service` façade so behavior is identical and evolves together.
- **Bring-your-own-agent from day one**. The MCP endpoint is remote (HTTP+SSE / Streamable HTTP), works with Claude Desktop, `claude` CLI, ChatGPT (once ChatGPT ships MCP client support broadly), and the Anthropic Agent SDK.
- **Do not conflate service names with content.** The customer's `skills`, `roles`, `skill-paths`, `simulations`, `assessments`, `sessions` come from the right owner service (see the content-vs-runtime split in `corpus/architecture/architecture_overview.md`). The API surface hides the split behind one clean product noun set.

**Write layer — cross-cutting contract (applies to W1 + W2 + every future write).**
Writes are additive to the read stack, gated behind M5 (audit + limits + tracing must exist first), and follow one uniform contract so REST / GraphQL / MCP surface them identically:
- **Per-resource READ vs WRITE scopes on API keys.** Scopes take the shape `people:read`, `people:write`, `people:delete`, `assignments:write`, `webhooks:manage`, etc. Writes are **DEFAULT-OFF** on every newly minted key (the key CRUD UI + API must require an explicit opt-in per resource, per action class).
- **Destructive-action gating.** Delete + deactivate + role-remove are a **separate elevated scope** (`*:delete`), never granted by `*:write`. **Soft-delete is the default** for every deletable resource (sets `deleted_at`, hidden from all default reads, reversible for 30d); **hard-delete requires the `*:delete:hard` scope AND a step-up confirmation** (`X-Confirm: irreversible` header + `?confirm=<resource_id>` param that must match the target — no cross-request replay).
- **Idempotency keys on all writes.** `Idempotency-Key: <caller-supplied ULID>` is REQUIRED on every mutating request (POST/PATCH/PUT/DELETE); server stores `(key_id, idempotency_key) → response_hash` with 24h TTL; replay returns the cached response unchanged; a mismatched body against the same key returns `409 idempotency_conflict`. GraphQL writes carry the same key in an `x-idempotency-key` header + persisted-op restriction.
- **Async job model for bulk (1 → 20 000).** Every bulk action returns `202 Accepted` with `{job_id, status_url, estimated_duration_s}` — never blocks. Jobs are first-class: `GET /v1/jobs/{id}` (poll), `job.progress` / `job.completed` / `job.failed` outbound webhooks, and per-row result stream so a poison row never rolls back its siblings (mirroring the backend's existing `bulkImportV2 → Job!` pattern). Job records inherit the caller's `Principal`, org, key id, and audit chain.
- **Dry-run / validate mode.** Every write endpoint (single and bulk) accepts `?dry_run=true` (REST) / `dryRun: true` (GraphQL) / `dry_run` arg (MCP): validates + returns the diff the write would produce, **never touches state**, no idempotency slot consumed, no audit row written (a `write_dry_run` row is instead).
- **MCP write tools — explicitly separated + annotated.** MCP tools split into `list_*` / `get_*` (read) vs `create_*` / `update_*` / `assign_*` / `delete_*` (write). Every write tool carries the MCP `destructive` + `non_idempotent` annotations (MCP spec fields) so LLM clients render a confirmation dialog. `delete_*` tools additionally require the client to echo an `acknowledge_destructive: true` argument (dropped silently by non-destructive tools). The `ask` endpoint (M7) remains **read-only** — it cannot invoke a write tool in v1 (the MCP write catalog is opt-in per key).
- **Audit before + after per write.** Every mutating request writes `api_audit_events` with `{before_snapshot, after_snapshot, diff, actor_principal, key_id, idempotency_key, dry_run, request_id}` — the two snapshots are the exact stored JSON blob of the entity pre/post write, `diff` is the RFC 6902 patch. Retention aligned to the M5 90-day window; longer for `*:delete:hard` (7 years, DPA-aligned).
- **Stricter write rate limits.** Per-key write budget defaults are an order of magnitude below reads: 6 rpm / 100 rph / 2 000 rpd for single writes; 10 async jobs/hour with per-job concurrency = 1 per resource kind. Bulk-job size caps: 20 000 rows/job, 5 concurrent jobs/org. Independent 429 accounting from the read budget — a write burn does not throttle reads.
- **Write endpoints go through the SAME persisted-operation registry as reads.** No customer-supplied SQL, no ad-hoc mutations by default (opt-in per key via the same `allow_adhoc_graphql` flag from M4).
- **Auth-layer independence carries over.** Every write handler receives the internal `Principal` from the middleware — never a Clerk claim, never a raw session. Writes carry `Principal.actor_id` into the audit row; when the actor is a service-principal (API key), that's what's recorded.

**Non-goals for v1.**
- No BYO-model / BYO-Bedrock for the `ask` endpoint.
- No fine-grained ABAC exposed to customers (they get resource-level scopes in v1 per the write-scope table above; row-level policies granularize in v2).
- The `ask` endpoint cannot invoke write tools in v1 (agentic writes deferred to a later track with mandatory human-in-the-loop confirmation).

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

**Read track (M1–M10) — original scope, ships first.**
- **M1** — Discovery, contract & security review (docs-only gate) — includes the Clerk-coupling audit + the `IdentityProvider` / `Principal` auth-abstraction seam + the **write-surface audit** (real-mutation inventory + gap table, see Appendix A)
- **M2** — API-key primitive: issue / rotate / revoke, hashed at rest, org-scoped — **our own primitive, IdP-independent**; carries **per-resource READ scopes AND WRITE scopes** (writes DEFAULT-OFF)
- **M3** — Public API gateway (`api.anthropos.work`) + REST read layer v1 (people, skills, roles, paths, simulations, sessions, assessments, reports)
- **M4** — Customer-facing GraphQL surface (persisted queries + the same read schema)
- **M5** — Rate limiting, quotas, audit log, error taxonomy, request tracing — **prerequisite for any write**
- **M6** — MCP Server (remote, Streamable HTTP), exposing the read tools + resource templates
- **M7** — Natural-language `ask` endpoint (customer-safe generalization of `askengine`) — REST, GraphQL, and MCP tool
- **M8** — `Enterprise Settings → Integrations` UI: key CRUD, MCP enable, quickstart, live docs, usage panel

**Write track (W1–W2) — inserted AFTER M5, gated by it; ships in parallel with M6–M8.**
- **W1** — Write foundations + user/org lifecycle + assignments: async job model (job id + poll/webhook), idempotency-key contract, dry-run/validate mode, stricter write rate limits, tighter audit-before/after. Ships PEOPLE (create/invite/update/deactivate/reactivate/soft-delete + elevated hard-delete, HRIS-style bulk upsert), ORG STRUCTURE (teams/departments — new; job-role CRUD — new; assign person↔role), ASSIGNMENTS (assign Sim / Interview / Skill Path / Academy / AI Lab to person or group, reschedule, cancel/revoke, BULK 1→20 000 async).
- **W2** — Skills + learning + sessions + reports + webhooks: SKILLS (assign/remove on person or role, set/override level, trigger reassessment, role-skill requirements+levels), LEARNING (enroll/unenroll skill path or development program, mark complete/reset), ASSESSMENT SESSIONS (launch/schedule sim or interview, resend invite, revoke), REPORTS+EVENTS (trigger report/export async; outbound webhooks — subscribe/manage/rotate-secret for `assessment.completed`, `skill.verified`, `assignment.completed`, `user.created`). Extends M6 with MCP write tools (destructive-annotated, confirmation semantics) + M8 with write-scope toggles + destructive-action confirmation UX.

**Release track.**
- **M9** — Developer docs site + Claude / ChatGPT / SDK quickstarts + reference SDKs (TS, Python) — includes write recipes + webhook receiver templates
- **M10** — Private beta → GA hardening: versioning policy, deprecation contract, SLOs, on-call runbook — separate SLOs + pentest scope for writes

Each milestone below has: **Goal · Scope in / out · Deliverables · Technical approach · Dependencies · Risks · Acceptance**.

---

## M1 — Discovery, contract & security review

**Goal.** Freeze the v1 product surface and its safety contract before any code lands. Turn "customers can query anything" into an enumerated schema + scope model that a security review can sign off on.

**Scope in.**
- Enumerate the v1 read surface (entities + fields) across all 5 subgraphs + `askengine`'s `rules.md` allowlist.
- Draft OpenAPI 3.1 + GraphQL SDL for the customer-facing schema (not the federated internal one).
- **Audit every Clerk coupling point touched by this initiative** — grep the corpus + `backend`, `apps/web`, and any consumer of `colony/authn`; produce `docs/api/clerk-coupling-audit.md` listing each site (file:line), what claim / SDK call is used, and the equivalent internal-principal call it must migrate to. This is the pre-condition for defining the seam.
- **Define the auth-abstraction seam.** Specify the internal `IdentityProvider` interface + the `Principal` DTO (`organization_id, actor_id, actor_kind ∈ {user, service_principal}, scopes[], roles[], session_id, issued_at, expires_at, provider_hint`) as the ONLY thing the public API / MCP / UI ever see. Clerk becomes one implementation (`identity/clerk`); a `Static` and a `Test` implementation ship alongside for local dev and tests. No branching on `provider_hint` above the adapter.
- Decide the auth model at the principal layer: **API key** (org-scoped, service-principal — a first-class actor_kind, no Clerk user needed) and **Personal Access Token** (user-scoped, inherits the actor's roles as resolved by the current IdentityProvider); v1 ships **API key only**, PAT deferred to M10+.
- Decide the tenancy contract: every key belongs to exactly one Anthropos `organization_id` (our own primary key — not Clerk's `eid`; the adapter is responsible for the `eid → organization_id` mapping); no cross-org keys.
- Threat model: key exfiltration, SSRF via `ask`, prompt injection through customer data, over-broad reads that hit the GB tables listed in `corpus/ops/db-access.md:32-38`.

**Scope out.** Any code changes, DB migrations, or UI work.

**Deliverables (docs).**
- `docs/api/surface-v1.md` — the v1 entity catalog with per-entity field list, filter set, and originating subgraph.
- `docs/api/security-model.md` — auth, tenancy, scopes, rotation, revocation, incident response.
- `docs/api/threat-model.md` — STRIDE table + mitigations, sign-off log.
- `docs/api/identity-seam.md` — the `IdentityProvider` interface, the `Principal` DTO, adapter contract, the Clerk-adapter mapping table, and the migration playbook for swapping providers.
- `docs/api/clerk-coupling-audit.md` — the exhaustive audit produced under Scope-in above.
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
- PAT model: separate `personal_tokens` table keyed to the internal `actor_id` (never to a provider-specific user id); the adapter resolves the actor. Do not overload the identity provider.

**Acceptance.**
- Surface catalog reviewed by product + at least one design partner.
- Threat model signed off by security.
- **`docs/api/identity-seam.md` signed off** — reviewers confirm the `Principal` DTO carries every field the API/MCP/UI will ever need, and that no downstream milestone requires a provider-specific claim.
- **`docs/api/clerk-coupling-audit.md` complete** — every coupling point has a named migration target on the seam.
- OpenAPI + GraphQL drafts lint clean (`spectral`, `graphql-schema-linter`).
- No code committed.

---

## M2 — API-key primitive: issue / rotate / revoke

**Goal.** Introduce a first-class, org-scoped `api_key` in the `backend` service (`public` schema) — **entirely our own primitive**, with no runtime dependency on Clerk (or any external IdP) for issuance, storage, or validation. This is the sole thing the future gateway trusts for customer traffic.

**Auth-independence contract (non-negotiable).**
- `apikey` package **must not import** any Clerk SDK, must not read any Clerk-shaped claim, and must not call the IdP at validation time. The only inputs to `apikey.Validate(secret)` are the DB row and the presented secret.
- The one admin-facing touchpoint — "an org admin requests a new key in-app" — goes through the **abstract `IdentityProvider`** defined in M1. The RPC/handler receives a `Principal` (already resolved by the adapter) and authorizes against `Principal.roles`, never against a Clerk claim. Swapping the IdP later touches the adapter only; the `apikey` service is untouched.
- Every `api_key` row references an actor as **`created_by_actor_id` (our own internal id)**, not a Clerk `user_id`. The adapter is responsible for mapping the external user to our internal actor at admin-login time.

**Scope in.**
- Ent schema: `api_keys` on `backend` (`public` schema — same place as `users`, `organizations`, `memberships`).
- CLI + Connect-RPC methods (internal only, admin-gated via the `Principal`) to issue/list/rotate/revoke.
- **Hashed at rest** (Argon2id or bcrypt) — plaintext returned once at creation only.
- Prefix + last-4 stored plaintext for display (e.g. `ant_live_ABCD…WXYZ`).
- `organization_id` FK is **required**; `created_by_actor_id` FK (internal actor, not an external user id); `scopes` string array with the structured resource:action grammar defined in `docs/api/scopes.md`; `expires_at` optional; `revoked_at`.
- **Scope grammar (v1).** Reads: `people:read`, `skills:read`, `roles:read`, `paths:read`, `simulations:read`, `sessions:read`, `assessments:read`, `reports:read` (a legacy `read:all` alias expands to all `*:read` scopes at issuance time; kept for the M3–M4 window). Writes (opt-in, DEFAULT-OFF): `people:write`, `org:write`, `roles:write`, `skills:write`, `assignments:write`, `learning:write`, `sessions:write`, `reports:write`, `webhooks:manage`. Elevated destructive: `people:delete`, `org:delete`, `roles:delete`, `assignments:delete`, `webhooks:delete`, plus the hard-delete tier `*:delete:hard` (each is a distinct scope, never implied by any other). The `ask` endpoint requires `ask:invoke` (separate — expensive).
- **Default-off enforcement at issuance.** `apikey.Manager.Create` rejects a key request that includes any `*:write`, `*:delete`, or `ask:invoke` scope unless the caller passes `explicit_write_optin: true` AND the `Principal` carries `roles ∋ "org:admin"`. The UI (M8) surfaces per-scope toggles + a destructive-cluster confirmation.
- Authorization policy (in Sentinel or its successor — Sentinel is called with `Principal.roles`, not with a Clerk claim): only an actor holding the internal `org:admin` role in the owning org may CRUD their org's keys.

**Scope out.** Gateway edge validation (M3), rate-limit accounting (M5), UI (M8), PATs.

**Deliverables (code).**
- `app/internal/data/ent/schema/apikey.go` (Ent schema) + `atlas migrate diff`ed migration.
- `app/internal/apikey/` package: `manager.go`, `hash.go`, `format.go` (key prefix + secret encoding), `manager_test.go`. No Clerk imports; enforced by an `import`-lint rule in CI.
- Connect-RPC: `backend.v1.ApiKeyService/{Create,List,Rotate,Revoke}` in `app/rpc.go`. Handlers receive a `Principal` from the identity middleware, never a raw Clerk session.
- Cobra subcommand `app apikey {create,list,rotate,revoke}` under `cmd/apikey`.
- Authorization: add the `org:feature:integrations:apikey:manage` policy gated to the internal `admin` role; wire into `init_policy.sql` + a `local_superadmin_grants.sql` note.
- Redis Streams event: `backend` stream carries `ApiKeyCreated|Rotated|Revoked` for downstream audit consumers (payload carries `actor_id`, not any external user id).

**Technical approach.**
- Key format: `ant_{env}_{22-char-b32}.{40-char-b32}` — `{22-char}` is the record id (short, indexable), `{40-char}` is the secret. Constant-time compare on lookup; index the record id.
- Hash: Argon2id (`golang.org/x/crypto/argon2`), memory=64MB, iters=3, parallelism=2. Cost-tune before ship.
- Rotation: creating a rotation returns a new secret and marks the old key `revoked_at = now() + grace` (default 24h grace, configurable per key).
- **Never log the plaintext.** Reuse the values-blind principle from `corpus/ops/secrets-spec.md`.

**Dependencies.** M1 (the `IdentityProvider` seam + `Principal` DTO must exist as spec).

**Risks.**
- If we ever leak a plaintext into `ai_usages` or a Sentry breadcrumb, it's game over. Add a `regexp` breadcrumb scrubber for `ant_live_.*\.` and unit-test it.
- Argon2 CPU cost on `create` — fine (rare op), unit-test that `verify` under load stays <5ms P99.

**Acceptance.**
- Create → list → verify (RPC) → rotate → revoke cycle covered by integration tests, driven by both the Clerk adapter and the Static adapter — behavior identical across both, proving auth-independence.
- Authorization denies a non-admin `Principal` on `Create`; permits admin.
- Import-lint rule fails CI if `apikey/` gains any Clerk-SDK import.
- Argon2 hash time within the tuned envelope on the reference CI runner.
- Audit event published to `backend` stream and consumed cleanly in a fixture consumer.

---

## M3 — Public API gateway + REST read layer v1

**Goal.** Stand up `api.anthropos.work` — a new edge service that terminates customer API-key auth, resolves tenant scope, and translates every request into a Cosmo Router GraphQL call. Ship a REST v1 read surface over it.

**Scope in.**
- New repo (or new dir in `platform`) `apigw` — a Go service using colony, sitting in the compose file at `:8500` (dev) / behind ALB (prod), profile `graphql` peers.
- REST v1 endpoints for the M1 catalog (paginated `GET /v1/{people,skills,roles,skill-paths,simulations,sessions,assessments,reports}` + `GET /v1/{entity}/{id}`).
- Key resolution middleware: `Authorization: Bearer ant_live_…` → hash-verify against `api_keys` → hydrate a `Principal` (`organization_id`, `actor_kind = service_principal`, `actor_id = api_key.id`, `scopes`, `roles = ["api:read"]`) → mint a **short-lived internal JWT signed BY OUR SERVICE against the `Principal`** (HS256, 5-min TTL). The claim set is our own (`org_id`, `actor_id`, `actor_kind`, `scopes`, `roles`) — **NOT** Clerk-shaped, **NOT** derived from any Clerk-issued token. Every subgraph accepts this internal JWT via a new `authn` provider (`identity/internal`), sitting alongside — never downstream of — the Clerk adapter.
- Query translation: every REST resource maps to a **persisted GraphQL operation** compiled at build time from `docs/api/schema.graphql`; no ad-hoc GraphQL execution from customer input.
- Response shaping: JSON:API-lite (`data`, `meta.page`, `links.next`) — pick a shape, document it, stick to it.

**Scope out.** Rate limiting (M5), audit (M5), MCP (M6), NL `ask` (M7), UI (M8).

**Deliverables.**
- `apigw/` service — `main.go`, `internal/{auth,router,graphql,resp}/`, OpenAPI-generated handler stubs (using `oapi-codegen`).
- `apigw/openapi.yaml` — the runtime spec (frozen from M1 draft).
- Persisted GraphQL operations under `apigw/persisted/*.graphql` + a compile-time SHA registry.
- `platform/docker-compose.yml` entry, `platform/repos.yml` entry, Terraform ECS service block, ALB routing rule for `api.anthropos.work`.
- `authn` extension: an **`Internal` identity provider** that trusts JWTs signed by the gateway with a shared secret (`APIGW_INTERNAL_JWT_KEY`, rotated separately). Claims are the internal `Principal` shape (`org_id`, `actor_id`, `actor_kind`, `scopes`, `roles`, `session_id`) — **explicitly not the Clerk claim shape**. Documented in `corpus/architecture/shared_libraries.md` alongside the Clerk adapter as peer implementations of the same `IdentityProvider` interface.

**Technical approach.**
- `apigw` **never** talks to Postgres directly. It only calls `http://graphql:8080/graphql` with the internal JWT — that path is already Ent-policy-filtered by `organization_id` in every subgraph. Tenancy is preserved by construction, not by hope.
- Persisted queries are **the only way** to reach the graph. This kills GraphQL introspection abuse and lets us reason about the exact fields customers can pull.
- Pagination: cursor-based (`opaque base64(json)` → `{after, limit}`), max page 200, default 50. Never offset (kills the DB on large orgs).
- Content type: `application/json` only for v1. No SSE except on the `ask` endpoint (M7).

**Dependencies.** M2 (`api_keys` exists).

**Risks.**
- **Internal-JWT provider is the weakest link.** If leaked, an attacker can impersonate any org. Mitigations: rotate every 24h automatically, sign in the gateway process only, never write the key to disk in prod (SecretsManager → memory), narrow the acceptance window to 5 min.
- **Provider-shape leak.** If any subgraph starts treating an internal-JWT claim as if it were a Clerk claim, we have re-coupled by accident. Guard: the `authn` middleware exposes only the `Principal` DTO to handler code; a CI import-lint rule forbids `context`-typed Clerk claim helpers in any code that also handles `Principal`.
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

## W1 — Write foundations + user/org lifecycle + assignments

**Goal.** Ship the first customer-facing write surface — people lifecycle, org structure (teams/departments + job roles, both new), and assignment lifecycle — on top of the async-job / idempotency / audit-before-after contract defined in the cross-cutting section.

**Ordering.** W1 lands **only after M5**. The cross-cutting write contract (idempotency, audit before/after, stricter limits, async jobs, dry-run) is a hard prerequisite — writes without it are not shipped.

**Scope in — foundations.**
- Async-job service in `apigw`: `POST /v1/jobs`, `GET /v1/jobs/{id}`, `GET /v1/jobs?filter=` (list of the caller's jobs); backing `api_jobs` table on `backend` with `{id, org_id, key_id, actor_id, kind, params, status, progress_pct, results_url, created_at, updated_at, expires_at}`. Reuses the existing `Job!` pattern already returned by `bulkImportV2` (documented at `app/internal/web/backend/graphql/graph/schemas/mutations.graphqls:164` per Appendix A). Watermill-driven worker in `backend`; the worker owns the write, `apigw` never writes directly.
- Idempotency middleware in `apigw` (`internal/idempotency/`): stores response hashes in Redis with 24h TTL; returns `409 idempotency_conflict` on body mismatch.
- Dry-run middleware: injects a `principal.dry_run = true` flag through to the persisted-op executor; every mutation resolver honors it.
- Write-scope enforcement middleware: rejects writes whose key's scopes do not cover the resource:action; destructive-scope + `X-Confirm` header enforcement.
- Audit-before/after writer: pre-write snapshot pulled by the persisted-op (SELECT before UPDATE/DELETE); post-write snapshot from the mutation response; `write_dry_run` audit-kind for dry-runs.

**Scope in — PEOPLE lifecycle.**
- `POST /v1/people` — create + invite (wraps the existing `inviteMember` mutation). Idempotent on `email + organization_id`.
- `PATCH /v1/people/{id}` — update attributes, manager, department, status (wraps existing profile mutations + new `manager_id` + `department_id` fields on `Membership`; the manager/department FKs are new and land as part of W1).
- `POST /v1/people/{id}/deactivate` + `POST /v1/people/{id}/reactivate` — new; adds `membership.status ∈ {active, deactivated}` (new column) so we can pause a seat without hard-removing it. Reversible. Gap-filler per Appendix A.
- `DELETE /v1/people/{id}` — SOFT-delete by default (sets `deleted_at`, hidden from all reads, reversible for 30d).
- `DELETE /v1/people/{id}?hard=true` — HARD-delete, requires `people:delete:hard` + `X-Confirm: irreversible` + `?confirm={id}`. Wraps the existing `removeMember` mutation which is already hard.
- `POST /v1/people:bulk-upsert` — HRIS-style bulk upsert on `{external_id | email}` as merge key: creates, updates, deactivates by-omission-toggle (opt-in via `deactivate_omitted: true`). Async job, up to 20 000 rows. Extends the existing `bulkImportV2` job pattern to full upsert semantics (which today is invite-only per Appendix A).

**Scope in — ORG STRUCTURE.**
- `POST|GET|PATCH|DELETE /v1/teams`, `/v1/departments` — **all new**. No corresponding backend mutations exist today (only `Tag` — Appendix A gap). Ships alongside new `teams`, `departments` tables on `backend` (`public` schema) + membership FKs (`membership.team_id`, `membership.department_id`) + Ent schemas. Soft-delete default; hard-delete tier as above.
- `POST|GET|PATCH|DELETE /v1/roles` — org-scoped job-role CRUD. Currently orgs can only ASSIGN a public-taxonomy `NodeID` role to a member (`createOrganizationRole` — Appendix A); v1 adds org-private role authoring backed by a new `organization_roles` table. Public-taxonomy roles remain read-only + referenced by `taxonomy_node_id`.
- `POST /v1/people/{id}/roles` + `DELETE /v1/people/{id}/roles/{role_id}` — assign / unassign person↔role. Wraps existing `createOrganizationRole` / `deleteOrganizationRole`.

**Scope in — ASSIGNMENTS.**
- `POST /v1/assignments` — assign an AI Simulation, AI Interview, Skill Path, Academy path, or AI Lab to a person or group. **Gap:** backend `AssignmentType` enum today covers only `skillPath` + `jobSimulation` (Appendix A); W1 extends the enum to `aiInterview | academyPath | aiLab` alongside a new `assignment_kind_registry` so future kinds don't require an enum bump. New assignment kinds land in the `assignments.graphqls` federated schema as an additive extension.
- `PATCH /v1/assignments/{id}` — reschedule (`due_date`) or reassign (`assignee_id`). Wraps `bulkUpdateOrganizationAssignments` on the single-row case; extends it with a new `updateOrganizationAssignment` mutation for the individual write (Appendix A gap — only bulk exists today).
- `POST /v1/assignments/{id}:revoke` — cancel / revoke a specific assignment (individual, non-bulk — new; today only `bulkDeleteOrganizationAssignments` exists). Wraps a new `revokeOrganizationAssignment` mutation.
- `POST /v1/assignments:bulk` — assign to 1 → 20 000 people in one async job. Wraps `createOrganizationAssignments` which already accepts a membership list; the async wrapper adds job + progress semantics.

**Scope out.** Skills mutations, learning enroll/unenroll, session launch, reports, webhooks — all W2.

**Deliverables.**
- `apigw/internal/{idempotency,jobs,scopes,audit_writer,dryrun}/`.
- `backend`: new Ent schemas + migrations for `api_jobs`, `teams`, `departments`, `organization_roles`, `membership.status`, `membership.manager_id`, `membership.team_id`, `membership.department_id`.
- New/extended backend GraphQL mutations: `updateOrganizationAssignment`, `revokeOrganizationAssignment`, team/department/role CRUD, `deactivateMember`, `reactivateMember`, `bulkUpsertMembers` (returns `Job!`).
- Persisted write-ops registry under `apigw/persisted/writes/*.graphql` — SHA-registered same as reads.
- `docs/api/writes-v1.md`, `docs/api/scopes.md`, `docs/api/idempotency.md`, `docs/api/jobs.md`, `docs/api/destructive-actions.md`.

**Dependencies.** M2 (scopes), M3 (persisted-op path), M4 (GraphQL path), M5 (audit + limits + tracing).

**Risks.**
- **Assignment enum extension** touches the federated `assignments.graphqls` in the backend subgraph — coordinate with the Cosmo static-composition build (a schema change is a supergraph rebuild + redeploy). Land the extension in an M5 → W1 window with the FE team synced.
- **HRIS bulk upsert** is a classic data-loss vector (a botched delimiter + `deactivate_omitted: true` = mass deactivation). Mandatory: dry-run required on first call per key, 24h "cooling" flag that requires dry-run within the prior week for every subsequent full-upsert.

**Acceptance.**
- Create + rotate a key with `people:write + assignments:write` (destructive OFF); assign a Skill Path to 10 members via `POST /v1/assignments:bulk`; job completes with per-row audit; a repeat call with the same `Idempotency-Key` returns the same job id.
- Attempt `DELETE /v1/people/{id}?hard=true` without `people:delete:hard` → 403; with the scope but without `X-Confirm` → 400; with both → hard-deletes + writes an audit row with retention 7y.
- Dry-run `POST /v1/people:bulk-upsert` with 500 rows returns a diff, touches no state, appears as `write_dry_run` audit-kind.
- The federated supergraph rebuilds cleanly with the extended `AssignmentType`.

---

## W2 — Skills, learning, sessions, reports, webhooks

**Goal.** Complete the write catalog: skill assignment + reassessment triggers, learning enroll/unenroll + mark-complete/reset, session launch/schedule/resend/revoke, report/export triggering, and **outbound webhook management** (subscribe / list / rotate secret / delete for the four event kinds).

**Ordering.** W2 lands after W1 (shares W1's job + idempotency + audit foundations).

**Scope in — SKILLS.**
- `POST /v1/people/{id}/skills` — assign a skill to a person; wraps existing `addUserSkill`.
- `DELETE /v1/people/{id}/skills/{skill_id}` — remove skill(s); wraps `removeUserSkills`.
- `PATCH /v1/people/{id}/skills/{skill_id}` — set/override level; wraps existing `overrideSkillLevel` / `rateUserSkillLevel`.
- `POST /v1/roles/{id}/skills` + `DELETE /v1/roles/{id}/skills/{skill_id}` — **role skill requirements + levels** (org-role side). **Gap:** no such backend mutation today (Appendix A); a new `role_skill_requirements` table + `setRoleSkillRequirement` / `removeRoleSkillRequirement` mutations land on the backend subgraph. Public-taxonomy role requirements stay read-only from customers (they belong to the skiller taxonomy).
- `POST /v1/people/{id}/skills/{skill_id}:reassess` — **trigger reassessment**. **Gap:** no direct mutation today (skill verification happens via simulation completion events per `corpus/services/skillpath.md` + the verified-skill chain). W2 exposes a new `triggerSkillReassessment` mutation that enqueues an assessment assignment for the person — a bridge, not a re-implementation. Async job.

**Scope in — LEARNING.**
- `POST /v1/people/{id}/enrollments` — enroll in a skill path or development program. **Gap:** skillpath auto-creates a session on view (`getOrCreateSkillPathSession`); W2 adds an explicit `enrollSkillPathSession` mutation (wrapping the auto-create with a source=api tag). Development programs don't exist as a first-class type yet — flagged as a spec gap in Appendix A; ships when the platform defines the type.
- `DELETE /v1/people/{id}/enrollments/{enrollment_id}` — unenroll (archives the `SkillPathSession`, soft-delete default). New `archiveSkillPathSession` mutation on the skillpath subgraph.
- `POST /v1/enrollments/{id}:complete` — mark complete (wraps repeated `completeSkillPathStep` calls into a single atomic bulk completion; new mutation).
- `POST /v1/enrollments/{id}:reset` — new `resetSkillPathSession` mutation (clears step completions, keeps history).

**Scope in — ASSESSMENT SESSIONS.**
- `POST /v1/sessions:launch` — launch or schedule an AI Sim / AI Interview session for a person. Wraps existing `createOrganizationSimInvitationLink` (invitation-link path) + adds a new `launchSimSession` mutation (Appendix A gap — today only invitation LINKS exist, not per-candidate launches).
- `POST /v1/sessions/{id}:resend-invite` — **new** (Appendix A gap). Wraps a new `resendSimInvitation` mutation that re-triggers the messenger send.
- `POST /v1/sessions/{id}:revoke` — wraps `revokeOrganizationSimInvitationLink`.

**Scope in — REPORTS + EVENTS.**
- `POST /v1/reports` — trigger a report / export (workforce activity, coverage-by-role, growth-by-period, roster export). Async job returning `results_url` on completion. **Gap:** no report/export trigger mutation today (Appendix A) — a new `queueReportExport` mutation lands on the backend subgraph, worker in `backend/internal/worker/` renders the export to `storage` + signs a URL.
- **Outbound webhooks** — first-class management API:
  - `POST /v1/webhooks` — subscribe: `{url, event_kinds[], secret_lifecycle: "generated" | "provided"}`; server returns a `signing_secret` on create (shown once).
  - `GET /v1/webhooks` — list.
  - `PATCH /v1/webhooks/{id}` — update url / event list.
  - `DELETE /v1/webhooks/{id}` — unsubscribe.
  - `POST /v1/webhooks/{id}:rotate-secret` — rotate signing secret; old secret honored for 24h grace.
  - `POST /v1/webhooks/{id}:test` — send a test event.
  - `GET /v1/webhooks/{id}/deliveries` — recent delivery attempts + retry status.
  - **Event kinds (v1):** `assessment.completed`, `skill.verified`, `assignment.completed`, `user.created` (Stefano's target set). Plus the internal `job.progress` / `job.completed` / `job.failed` from W1's async model.
  - Delivery: HMAC-SHA256 signature in `X-Anthropos-Signature`, timestamp in `X-Anthropos-Timestamp`, retry with exponential backoff (0/15s/1m/10m/1h/6h then dead-letter). Delivery attempts audited.

**Scope in — MCP WRITE TOOLS (extends M6).**
- Every write endpoint above gets an MCP tool variant: `create_person`, `update_person`, `deactivate_person`, `delete_person`, `assign_role`, `create_team`, `create_assignment`, `bulk_upsert_people`, `enroll_skill_path`, `launch_session`, `queue_report`, `subscribe_webhook`, etc.
- Each tool carries the MCP `destructive` + `non_idempotent` annotations per the cross-cutting contract; `delete_*` and `deactivate_*` require an `acknowledge_destructive: true` argument.
- The MCP server exposes write tools only if the connecting key holds the matching write scope — the tool list a client sees is scope-filtered at connect time (so an LLM never even sees a tool it cannot invoke).
- The `ask` endpoint remains read-only in v1.

**Scope out.** Row-level ABAC scopes, delegated tokens on behalf of a user, agent-initiated writes without operator confirmation.

**Deliverables.**
- Backend: new mutations (`setRoleSkillRequirement`, `removeRoleSkillRequirement`, `triggerSkillReassessment`, `archiveSkillPathSession`, `resetSkillPathSession`, `launchSimSession`, `resendSimInvitation`, `queueReportExport`) + `webhook_subscriptions`, `webhook_deliveries` tables + worker.
- New service `webhook-dispatcher/` (or a `backend/internal/webhook_dispatcher` package if we keep it in-process) — consumes Redis Streams events (`backend`, `skillpath`, `jobsimulation`) and dispatches to customer endpoints.
- `apigw`: write-tool wiring for all of the above; MCP write-tool wiring on `mcp-server/`.
- UI (extends M8): write-scope toggle UI + destructive-action confirmation UX (see M8 diff below).
- `docs/api/webhooks.md`, `docs/api/events.md`, `docs/api/writes-v2.md`.

**Dependencies.** W1, M6 (MCP), M8 (UI toggles).

**Risks.**
- Webhook delivery infrastructure is a service tier of its own — undelivered → data-loss customer complaints. SLO: 99.9% of events delivered within 5 min P95; automatic dead-letter after 6h; UI shows failing subscriptions loudly.
- `triggerSkillReassessment` is expensive (spawns a sim/interview) — bill against the org's `ai_usages` budget and rate-limit per-person (max 1 reassessment/skill/24h).
- Enrolling a person in a development program that doesn't exist as a first-class platform type is a spec-gap — ship the skill-path variant in W2, defer development-programs to W2b or the next release.

**Acceptance.**
- Subscribe to `assessment.completed`; complete a sim end-to-end; receive a signed webhook within 30s; rotate the secret; the next delivery uses the new secret; the old secret honored for 24h then rejected.
- `POST /v1/reports` returns a job id; `GET /v1/jobs/{id}` polls to completion; `results_url` downloads the export.
- Claude Desktop (connected via MCP with `assignments:write`) can `create_assignment` for a person after user confirmation; without `assignments:write`, the tool is not visible.
- All writes appear in `api_audit_events` with `before` + `after` snapshots.

---

## M6 — MCP Server: remote Streamable-HTTP endpoint

**Goal.** Ship `mcp.anthropos.work` — a hosted MCP server that exposes the M3–M4 read catalog as MCP tools + resource templates, so Claude Desktop / Claude Code / any MCP-capable client can connect in minutes.

**Scope in.**
- Transport: **Streamable HTTP** (the current spec-recommended remote transport) + SSE fallback for older clients.
- Auth: same `Authorization: Bearer ant_live_…` API key as M3 (MCP allows arbitrary HTTP headers) plus an optional `?key=…` for clients that don't let you set headers (documented as inferior).
- **Read tools (v1)** — always visible:
  - `list_people`, `get_person`
  - `list_skills`, `get_skill`, `search_skills`
  - `list_roles`, `get_role`
  - `list_skill_paths`, `get_skill_path`, `list_skill_path_sessions`
  - `list_simulations`, `get_simulation`, `list_simulation_sessions`
  - `list_assessments`, `get_assessment`, `list_verified_skills`
  - `run_report` (canonical rollups: workforce_activity, coverage_by_role, growth_by_period)
  - `ask` — deferred to M7, but the tool slot is reserved
- **Write tools (v1)** — surfaced only when the connecting key carries the matching write scope; hidden entirely from `tools/list` otherwise (per-scope visibility filter in the MCP server). Every write tool declares the MCP `destructive` / `non_idempotent` annotations honestly and requires the same idempotency key + async-job contract as REST (see W1/W2):
  - `create_person`, `update_person`, `deactivate_person`, `reactivate_person`, `soft_delete_person` — behind `people:write`; `hard_delete_person` behind `people:delete:hard` **plus** an explicit `acknowledge_destructive: true` argument (any write tool tagged `destructive` refuses to run without it — the LLM must be told to pass it).
  - `create_team`, `update_team`, `delete_team`, `create_department`, `update_department`, `delete_department`, `create_role`, `update_role`, `delete_role` — behind `org:write` / `roles:write`.
  - `assign_content`, `bulk_assign_content`, `revoke_assignment` — behind `assignments:write`; bulk returns `{job_id, status_url}` and the MCP server emits a progress notification when the job finishes (or the client can call `get_job`).
  - `add_person_skill`, `remove_person_skill`, `set_role_skill_requirement`, `trigger_skill_reassessment` — behind `skills:write` / `roles:write`.
  - `enroll_person_in_path`, `unenroll_person_from_path`, `launch_session`, `resend_session_invite`, `revoke_session_invite` — behind `learning:write` / `sessions:write`.
  - `queue_report_export`, `create_webhook`, `rotate_webhook_secret`, `delete_webhook` — behind `reports:write` / `webhooks:manage`.
  - `get_job` — job-status polling tool (surfaces to any key that holds any write scope, so async writes are inspectable).
- Resources (v1): `person://{id}`, `skill://{id}`, `role://{id}`, `skill-path://{id}`, `simulation://{id}`, `session://{id}`, `job://{id}` — each resolves to a JSON blob (or a small markdown summary + JSON payload).
- Prompts (v1): a small set of curated prompts — e.g. `weekly-workforce-summary`, `role-skill-gap`, `hiring-funnel-status` — each is a saved prompt template that uses the tools above.
- Discovery: `GET /` on the MCP endpoint returns human-readable "how to add this to Claude Desktop / ChatGPT" instructions and a link to the docs.

**Scope out.** Client-side installers, non-remote (stdio) MCP variant, unbounded free-form write tools (every write tool is one-purpose-one-tool — no `execute_graphql` escape hatch).

**Deliverables.**
- New service `mcp-server/` — Go, uses the official `mcp-go` SDK (or the reference Anthropic MCP Go implementation if we adopt it), boots on `:8600` in compose, deployed as its own ECS service.
- Each tool is a thin wrapper that calls `apigw` internally (over an in-cluster HTTP client with a service-account token) — the MCP server is **not** a second query engine, it's a tool-shaped face on the same one.
- `docs/api/mcp.md` — tool catalog, resource URIs, prompt catalog, connect guide.
- Compose + ALB route for `mcp.anthropos.work`.

**Technical approach.**
- Because MCP is invoked by an LLM, tool schemas must be **narrow, typed, and unambiguous** — one filter per parameter, explicit enums, no free-form JSON. Write a per-tool `zod`/JSON-Schema definition and test it against a Claude Sonnet call from a fixture.
- The MCP server holds no data. It resolves the API key on connect, maps it to an `org_id`, and threads that into every tool call as a header to `apigw`.
- Streaming: `ask` returns a stream of MCP progress notifications; other tools return synchronously.

**Dependencies.** M3, M4, M5, W1, W2 (writes are optional per-scope; the read cut can ship first).

**Risks.**
- MCP spec churn — pin the version we implement in `docs/api/mcp.md`; add a CI job that pins the SDK version and re-tests the connect flow.
- Prompt injection through customer data pulled by a tool. Mitigations documented in `docs/api/prompt-injection.md`; do not concatenate untrusted content into agent instructions on the server side.
- **LLM misuse of destructive tools** — an agent can be tricked into calling `hard_delete_person` if we're careless. Mitigations: `destructive`/`non_idempotent` annotations honestly declared so hosts show confirmation UI; `acknowledge_destructive: true` required argument on any destructive tool; write-scope default-off (M2) means most keys never see these tools; every destructive tool call written to the audit log with the full prompt + tool_call trace.

**Acceptance.**
- Adding the MCP URL + key to Claude Desktop's config surfaces every tool the key is scoped for — and only those.
- A Claude Sonnet chat "how many verified Python skills across my org?" resolves via `list_verified_skills` filtered by skill = python.
- A read-only key sees 0 write tools; a key with `people:write` sees `create_person`/`update_person`/`deactivate_person`/`reactivate_person`/`soft_delete_person` but NOT `hard_delete_person`; a key with `people:delete:hard` sees the hard-delete tool but calling it without `acknowledge_destructive: true` returns a structured error.
- MCP-Inspector CLI shows all tools + resources + prompts pass validation, including annotation correctness on every write tool.

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
  - **API Keys** — list, create, rotate, revoke; show `prefix + last4`, created-by, last-used-at (from audit), scope, expiry; **plaintext shown once at creation**, with a "download env" affordance. The **create-key wizard** is a two-step form:
    1. **Reads** — pick per-resource read scopes from a checklist (`people:read`, `skills:read`, `roles:read`, `paths:read`, `simulations:read`, `sessions:read`, `assessments:read`, `reports:read`) with a "Select all reads" convenience; default = all reads on.
    2. **Writes (advanced, collapsed by default)** — a separately-collapsed panel with a top-of-panel warning banner ("Write scopes let this key modify data. Enable only what you need."). Each write scope is a toggle: `people:write`, `org:write`, `roles:write`, `skills:write`, `assignments:write`, `learning:write`, `sessions:write`, `reports:write`, `webhooks:manage`. Enabling ANY write toggle reveals a mandatory checkbox: "I understand this key can modify {resource} in my organization."
    3. **Destructive scopes (elevated)** — `*:delete` (soft-delete) and `*:delete:hard` (hard-delete) sit under a second, further-collapsed panel gated by a "Show destructive scopes" reveal. Turning on any hard-delete scope forces a typed confirmation: the admin types their org name to unlock the Create button. The `X-Confirm: irreversible` header is documented next to the toggle so the customer's client code makes the requirement legible.
  - **Key detail view** — the scope list is rendered as pill groups (Reads / Writes / Destructive) with per-scope in-place edit and a "revoke scope" affordance; scope revocation is single-click and takes effect immediately (Sentinel policy update within one Reload cycle).
  - **MCP Server** — one toggle to enable/disable at org level (so an admin can shut off MCP globally without touching keys), the endpoint URL, per-client connect snippets (Claude Desktop config JSON, `claude` CLI, ChatGPT), and a "MCP-visible tools for this key" preview that mirrors the M6 scope-visibility filter (admins can see exactly which tools a key will expose to an LLM before they hand it out).
  - **Usage** — chart of requests / errors / cost this month, per key, from `api_audit_events` + `ai_usages`. Read-only rollups; write requests split out visually (a second series) so admins immediately see if a write scope is being exercised.
  - **Jobs** — list of async write jobs (bulk imports, bulk assigns, report exports) with status, key, submitter, counts, and a "view diff" link that opens the audit RFC-6902 diff for the affected rows.
  - **Webhooks** — CRUD on the org's outbound webhook endpoints (URL, event filter, signing-secret rotate, test-fire button, last 10 deliveries with attempt counts); driven by W2.
  - **Docs** — inline links to the developer docs (M9) with the org's example key embedded.
  - **Quickstart** — a 3-step banner: (1) create key, (2) copy this snippet, (3) run it, see it work; shown until the org has made ≥ 1 successful API call.
- **Admin authorization via the abstract identity, not Clerk.** The page gate, the mutation gate, and every conditional render key off the internal `Principal.roles` (e.g. `hasRole('org:admin')`) exposed by a Next.js server-side `getPrincipal()` helper that hides which IdP is behind it. **No component in `apps/web/src/app/enterprise/settings/integrations/` may import `@clerk/*` directly.** Swapping Clerk later replaces the `getPrincipal()` implementation only; the Integrations UI is untouched.
- All data via the federated GraphQL — the customer-facing UI itself must not use the new customer API (avoid the loop where our own UI depends on our public throttles).

**Scope out.** PAT UI (deferred), team/user-level scopes (deferred).

**Deliverables.**
- Next.js routes + Ant Design 6 pages in `apps/web/src/app/enterprise/settings/integrations/`.
- New GraphQL ops in `next-web-app/packages/graphql/src/query/integrations/*.graphql` + generated hooks.
- Backend GraphQL: extend the `backend` subgraph with `orgApiKeys`, `createApiKey`, `rotateApiKey`, `revokeApiKey`, `updateApiKeyScopes`, `orgMcpSettings`, `setOrgMcpEnabled`, `orgApiUsage`, `orgApiJobs`, `orgWebhooks`, `createOrgWebhook`, `updateOrgWebhook`, `rotateOrgWebhookSecret`, `deleteOrgWebhook`, `testOrgWebhook`, `listOrgWebhookDeliveries`.
- Storybook stories for the new components; Playwright E2E: (a) create read-only key → hit `/v1/people` with it → see it in Usage; (b) create write-scoped key → `POST /v1/people` → see the row + audit diff in Jobs; (c) try to enable hard-delete scope without typed org-name confirmation → Create button stays disabled.

**Technical approach.**
- Mirror the existing `/enterprise/*` UX patterns already in `apps/web` — TanStack React Query + `useGraphql`, plus a new `<PrincipalProtect requires="org:admin">` component that wraps the existing IdP-specific gate. Under the hood it reads from `getPrincipal()`; today it delegates to Clerk's `<Protect>`, tomorrow it delegates to whatever replaces it. Consumers never see the swap.
- **Plaintext key display** — one-time modal with a copy button, big "you will not see this again" warning, hash-verify test button.
- **Destructive-confirmation UX** — reused as a component `<DestructiveConfirmModal resource="…" orgName="…">` so the same typed-org-name pattern is reachable from any future admin surface that mints destructive scopes.
- **Write-scope preview** — the wizard's final step shows a **plain-English audit sentence** ("This key can: read people, skills, roles; write people; hard-delete people.") derived from the toggled scopes, so admins double-check semantics before submit.
- Endpoint URLs shown per environment (dev / staging / prod) — the UI reads these from `NEXT_PUBLIC_API_BASE_URL` and `NEXT_PUBLIC_MCP_BASE_URL`.

**Dependencies.** M1 (identity seam), M2, M3, M5, M6, W1, W2 (webhooks + jobs pages need W2's endpoints).

**Risks.**
- Copy-to-clipboard on plaintext key on non-HTTPS localhost silently no-ops — add a fallback and warn.
- Admins mis-reading the "usage" chart as billing — label carefully; link out to Billing.
- **Admin over-scoping** — the natural pull is to check every write box "just in case". Mitigations: reads default-on but writes default-off; a persistent inline hint under the write panel ("Anthropos recommends scoping keys to the narrowest surface a single integration needs.").

**Acceptance.**
- Playwright: fresh admin can create a key, run a `curl`, see a green check in Usage within 30s.
- The Integrations page renders empty-state correctly for a brand-new org.
- Non-admin `Principal` gets 403 on all `orgApiKey*` + `orgWebhook*` mutations.
- Enabling a `*:delete:hard` scope requires the typed-org-name confirm; the Create button remains disabled otherwise (asserted in E2E).
- Grep gate: `apps/web/src/app/enterprise/settings/integrations/**` has zero `@clerk/*` imports.

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
M1 (design)  ──►  M2 (api_key + scope grammar)  ──►  M3 (REST reads) ──►  M4 (GraphQL reads)
                                        │                          │
                                        └──►  M5 (limits+audit) ◄──┘
                                                    │
                                          ┌─────────┴─────────┐
                                          ▼                   ▼
                                       W1 (writes:            M7 (ask)
                                       people/org/                │
                                       assignments)               │
                                          │                       │
                                          ▼                       │
                                       W2 (skills/                │
                                       learning/                  │
                                       sessions/                  │
                                       webhooks)                  │
                                          │                       │
                                          ├──────────►  M6 (MCP: reads + scoped writes)
                                          │                       │
                                          └──────────►  M8 (UI: keys + write toggles + webhooks + jobs)
                                                                  │
                                                                  ▼
                                                              M9 (docs + SDKs)
                                                                  │
                                                                  ▼
                                                              M10 (beta → GA)
```

**Parallelizable slices.** Read track (M1→M5) is a hard prefix. Once M5 lands, W1 and M7 can proceed in parallel; W2 depends on W1's async-job + idempotency middleware but not on its endpoint surface. M6 can ship a **read-only** cut against M5 immediately, with write tools layering in as W1/W2 land. M8 can ship a **read-only Integrations page** off M5 and then extend to write-scope toggles / webhooks / jobs as W1 + W2 land. M9 depends on the schemas being stable across M3/M4/W1/W2/M6/M7. M10 depends on everything.

---

## Appendix A — Write-surface audit: real mutations discovered + gap analysis

This appendix is the evidence backing M1's audit + W1/W2's scope-in. It enumerates the mutations that **already exist** in the platform's GraphQL subgraphs (as of 2026-06 read) and the gaps we must close to deliver the target write catalog. Sources: `internal/web/backend/graphql/graph/schemas/*.graphqls` in the `ant-platform-backend` repo (subgraphs: `backend`, `academy`, `labs`, `assignments`, `ai_readiness`, etc.).

### A.1 What already exists (usable behind M2/W1/W2)

| Domain | Existing mutation | Sub-graph | Notes |
|---|---|---|---|
| People / membership | `inviteMember`, `bulkImportV2(input): Job!`, `removeMember`, `bulkRemoveMembers`, `changeMemberRole` | backend | `bulkImportV2` **already returns `Job!`** — validates the async-job pattern the whole write layer standardises on. Reused by W1's `POST /v1/people:bulk-upsert`. |
| Self-profile | `updateMeInBriefInfo`, `updateProfileInfo`, `addWorkExperience`/`updateWorkExperience`/`deleteWorkExperience`, `addEducation`/…, `addCertification`/…, `addProject`/…, `addVolunteering`/…, `addContent`/… | backend | These are **self-writes** (subject == principal). NOT exposed customer-side in W1 (customer API is admin-side). Kept internal. |
| Org | `createOrganization` | backend | Org creation stays platform-scoped for now — customers don't create orgs via the API. |
| Skills | `addUserSkill`, `removeUserSkills`, `rateUserSkillLevel`, `overrideSkillLevel`, `updateUserCoreSkills` | backend | Backs W2's `POST/DELETE /v1/people/{id}/skills` + `PATCH` (level override). |
| Assignments | `createPersonalAssignment`, `createOrganizationAssignments(input: OrganizationAssignmentInput!)`, `bulkUpdateOrganizationAssignments`, `bulkDeleteOrganizationAssignments` | backend | `createOrganizationAssignments` already handles a *list* of memberships in one call — the bulk path. Backs W1's `POST /v1/assignments` + `:bulk`. |
| Sim invitations | `createOrganizationSimInvitationLink`, `revokeOrganizationSimInvitationLink` | backend | Link-based flow only. Backs W2's `POST /v1/sessions:invite-link` variant. |
| Tags | `addTag`, `editTag`, `deleteTag`, `tagMembers`, `untagMembers` | backend | Not in the customer catalog v1 — parked. |
| Target roles | `createUserTargetRole`, `createOrganizationTargetRole`, `createOrganizationRole` | backend | `createOrganizationRole` is the closest existing hook for W1 `POST /v1/roles`, but only *creates* — no update/delete. |
| Labs | `createLabSession(template, model, agent, budgetUsd, mode)`, `cancelLabSession(id)` | labs | Self-invoked only — no "assign lab to a candidate" mutation. |
| Academy | `upsertChapterProgress`, `claimPathCertificate`, `addAcademyBookmark`, `upsertAcademyFeedback` | academy | Learner-side only — no admin-side "assign academy path" mutation. |
| AI Readiness | `completeAiReadinessSkillMapping: Boolean!` | ai_readiness | Self-completion. Not customer-catalog. |

### A.2 Gaps (net-new mutations W1/W2 add)

| Gap | Milestone | New mutation(s) or table(s) |
|---|---|---|
| Teams as first-class org units (only `Tags` exist today) | W1 | New `teams` table + `createTeam` / `updateTeam` / `deleteTeam` mutations + `membership.team_id` column. |
| Departments as first-class org units | W1 | New `departments` table + full CRUD + `membership.department_id` column. |
| Org-private job-role CRUD (today: create-only; taxonomy roles are read-only) | W1 | New `organization_roles` table + `createOrganizationRole` (extended), `updateOrganizationRole`, `deleteOrganizationRole`. |
| Role → required skills with levels | W2 | New `role_skill_requirements` table + `setRoleSkillRequirement` / `removeRoleSkillRequirement`. |
| Soft-delete vs hard-delete on person removal (today: `removeMember` is one-shot) | W1 | New `membership.status` column (`active|deactivated|deleted`) + `deactivateMember` / `reactivateMember` (soft) alongside `hardDeleteMember` (destructive, elevated). |
| HRIS-style bulk upsert (today: `bulkImportV2` invites only) | W1 | Extend `bulkImportV2` (or new `bulkUpsertMembers`) to accept `{external_id, ...}` and idempotently insert-or-update by external_id. |
| Trigger skill reassessment on demand | W2 | New `triggerSkillReassessment(userId, skillId): Job!` returning the async job. |
| `AssignmentType` enum missing modalities | W1 | Extend enum: `skillPath, jobSimulation` → add `aiInterview | academyPath | aiLab`. Corresponds to new `AssignmentResource` subtypes. |
| Assign academy path (no mutation today) | W1 | New `createAcademyPathAssignment` (via extended `createOrganizationAssignments` with `academyPath` variant). |
| Assign AI Lab (no mutation today) | W1 | New `createLabAssignment` variant on `createOrganizationAssignments`. |
| Per-assignment reschedule / update (today only bulk) | W1 | New singular `updateOrganizationAssignment(id, input)`. |
| Individual assignment revoke (today only `bulkDelete`) | W1 | New `revokeOrganizationAssignment(id)`. |
| Launch a per-candidate session directly (today only invitation-link) | W2 | New `launchSimSession(userId, simulationId, scheduledAt?): Session!`. |
| Resend invite (today none) | W2 | New `resendSimInvitation(sessionId)`. |
| Enroll / unenroll into a skill path atomically (today auto-created by skillpath runtime) | W2 | New `enrollSkillPathSession(userId, skillPathId): SkillPathSession!` + `archiveSkillPathSession(sessionId)`. |
| Mark path complete / reset (admin override) | W2 | New `markSkillPathSessionComplete(sessionId)`, `resetSkillPathSession(sessionId)`. |
| Trigger report / export (today none) | W2 | New `queueReportExport(kind, params): Job!`. |
| Outbound webhooks (today: Clerk-inbound only) | W2 | New `webhook_endpoints` table + full CRUD + signing-secret rotate + `webhook_deliveries` table + retry worker. Event kinds: `assessment.completed`, `skill.verified`, `assignment.completed`, `user.created`, `job.*`. |
| Async-job registry (today: only `bulkImportV2` returns `Job!`, no generic surface) | W1 | Formalise `jobs` table + `jobService.Enqueue/Get/Update` + `GET /v1/jobs/{id}` + a `job.completed` webhook so every write ≥ N rows follows the same shape. |
| Idempotency-key store (today: none) | W1 | New `idempotency_keys` table keyed by `(api_key_id, key)` + middleware; 24h TTL + 409 on payload conflict. |

### A.3 Non-goals (out of the customer write catalogue for v1)

- Direct-to-DB writes bypassing the subgraph (never — every write flows through a subgraph mutation via `apigw`).
- Free-form GraphQL mutations for customers (persisted mutations only; ad-hoc gated per-key like the read side).
- Self-profile edits on behalf of users (customer API is admin-side; user-owned edits stay in `next-web-app`).
- Org creation via the customer API (platform-scoped).
- Tag CRUD (parked — will re-enter roadmap only if a design partner asks).
