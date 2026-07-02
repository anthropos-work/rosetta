---
milestone: M302
slug: access-primitive
version: v3.0 "open house" (PROPOSAL — awaiting review)
milestone_shape: section
status: planned
created: 2026-07-01
last_updated: 2026-07-01
complexity: medium
delivers: the customer-facing API-key primitive (mint / rotate / revoke) + the `ApiKeyIdentityProvider` adapter + the append-only audit ledger (`customer_api.audit_events` Postgres table) + the shared audit-write middleware + the Redis token-bucket rate-limit middleware + the `admin`-tier catalog entries (`access.api-key.*`). No customer-data endpoint yet — the primitive is proven with a `/v1/access/whoami` + `/v1/access/api-keys` echo.
depends_on: M301
spec_ref: knowledge/plan/spec-drafts/customer-api-mcp/spec.md (the consolidated pillar spec, v0.1)
---

# M302 — Access primitive

## Goal
Land the **three cross-cutting middleware surfaces** every subsequent read + write endpoint depends on:
1. The **customer-facing API-key credential** (mint / rotate / revoke) with its `ApiKeyIdentityProvider`
   adapter — a `Principal` source that sits alongside the Clerk-JWT one from M301.
2. The **append-only audit ledger** + shared middleware that writes an audit row for every request.
3. The **Redis token-bucket rate-limit middleware** that budgets every request against a
   `Principal.id` + `rate_limit_bucket` key.

Still no customer-data endpoint. The primitive is proven with a `/v1/access/whoami` + `/v1/access/api-keys` echo
under the new admin-tier catalog entries.

## Scope
**In:**
- **(1) API-key store** — Postgres `customer_api.api_keys` (columns: `id`, `organization_id`, `hashed_key`,
  `key_prefix`, `label`, `scopes[]`, `created_at`, `last_used_at`, `revoked_at`, `rotated_from_id`). Keys are
  **stored hashed** (argon2id — decided at this milestone if not before), never plaintext.
- **(2) `ApiKeyIdentityProvider`** — reads `Authorization: Bearer ak_live_...`, hashes + lookups, produces a
  `Principal` with `identity_source = "api-key"`. Rejects with a plain 401 on unknown / revoked.
- **(3) The mint/rotate/revoke path** — three admin-tier catalog entries (`access.api-key.create`,
  `access.api-key.rotate`, `access.api-key.revoke`). Only a **Clerk-authenticated** Principal with the admin
  scope can call them (the reverse-dependency: customer UI → Clerk → mint an API-key Principal for scripts).
- **(4) Audit ledger** — the append-only `customer_api.audit_events` table (§5.6) + the shared middleware that
  writes a row for every request (success + failure). W2 writes will only persist `input_hash`, never raw.
- **(5) Rate-limit middleware** — Redis token-bucket keyed by `Principal.id` + `rate_limit_bucket`; standard
  `X-RateLimit-*` headers + HTTP 429 + `Retry-After` on exhaustion. Default budgets (60 req/min, 10k req/day).
- **(6) A proof `/v1/access/whoami` + `/v1/access/api-keys` (list only)** — feature-flagged, internal-only in
  M302 (customer UI opens in M304). Proves the three middlewares wire together.

**Out:** any customer-data endpoint (M303); the customer self-serve UI in Workforce (M304); rate-limit
per-tenant overrides via a customer surface (the override lives in the platform-internal admin surface — out of
scope for the customer API entirely).

## Why section (not iterative)
The API-key primitive + audit ledger + rate-limit middleware are three well-understood, decomposable pieces with
established patterns (argon2id, Postgres append-only + S3 archive, Redis token-bucket). No exploration; a
`section` build. `/developer-kit:build-milestone`.

## Depends on / Parallel with
- **Depends on:** **M301** — reuses the `Principal` DTO + `IdentityProvider` adapter port + the catalog registry
  (the three access-tier entries live in the catalog).
- **Parallel with:** **none** — M303 and M304 both depend on this middleware existing.

## KB dependencies
Read as contract:
- [`knowledge/plan/spec-drafts/customer-api-mcp/spec.md`](../../../spec-drafts/customer-api-mcp/spec.md) — the
  pillar spec (§5.5 API keys, §5.6 audit + rate limits).
- [`corpus/services/backend.md`](../../../../../corpus/services/backend.md) — the `app` service the primitive
  lands inside.
- [`corpus/architecture/security_compliance.md`](../../../../../corpus/architecture/security_compliance.md) —
  the tenant isolation + audit posture the primitive lands on.
- OWASP guidance on password / secret hashing — argon2id parameters (decision recorded in `decisions.md`).

## Re-scope trigger
If the Redis token-bucket approach fights the shared Redis's memory budget under expected R1 traffic (multiple
buckets per key), fall back to Postgres-backed leaky-bucket for R1 (slower but bounded), documented in
`decisions.md`. Never ship without a working rate-limit floor.

## Delivers →
- The API-key primitive + `ApiKeyIdentityProvider` — the customer's script credential.
- The append-only audit ledger + shared middleware — the compliance floor (G6).
- The rate-limit middleware — the abuse floor (G6).
- Three admin-tier catalog entries — the first `admin` action_type in the catalog; the customer's
  key-management surface (surfaced in Workforce at M304).

