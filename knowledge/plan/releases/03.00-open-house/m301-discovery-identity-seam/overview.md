---
milestone: M301
slug: discovery-identity-seam
version: v3.0 "open house" (PROPOSAL — awaiting review)
milestone_shape: section
status: planned
created: 2026-07-01
last_updated: 2026-07-01
complexity: medium
delivers: the API resource-catalog registry (Products → Resources → Actions → Tools, §4.1); the internal `Principal` DTO + `IdentityProvider` adapter port (§5.4); the `ClerkIdentityProvider` adapter (Clerk today, swappable tomorrow); the "no `clerk.*` above the adapter" lint rule + review guardrail; the OpenAPI-3.1 machine source (with `x-anthropos-*` extension for MCP fields). NO public endpoint ships in M301 — this is the seam.
depends_on: none (v3.0's foundation; the R1 spine)
spec_ref: knowledge/plan/spec-drafts/customer-api-mcp/spec.md (the consolidated pillar spec, v0.1)
---

# M301 — Discovery + Identity seam

## Goal
Land the **two load-bearing seams** of the Customer API + MCP pillar in a single milestone: the
**resource-catalog registry** (the machine source that both shells will project from) and the
**auth-vendor-independent identity seam** (the internal `Principal` DTO + `IdentityProvider` adapter port).

No public endpoint ships in M301. This is the **seam** — the shape that everything else in R1 (and R2..R6)
consumes.

## Scope
**In:**
- **(1) The resource-catalog registry** — the machine source (§4.1–§4.3): `product` / `resource` / `action` /
  `action_type` / `principal_scope` / `entitlement_tier` / `input_schema` / `output_schema` /
  `rate_limit_bucket` / `audit_shape` / `docs`. **OpenAPI 3.1 + a small `x-anthropos-*` extension** for the MCP
  fields (spec-progress #2, decided at this milestone).
- **(2) The `Principal` DTO** (§5.4): `id`, `organization_id`, `user_id?`, `scopes[]`, `entitlement_tier`,
  `identity_source`. Not a Clerk claim, not a JWT — an internal type.
- **(3) The `IdentityProvider` adapter port** — one implementation lands here (`ClerkIdentityProvider`); the
  second (`ApiKeyIdentityProvider`) lands in M302 alongside the key primitive.
- **(4) The lint / review guardrail** — a static check (`ClerkGuardrails`) forbidding `clerk.*` imports in
  `internal/customerapi/` above the adapter package. **Machine-enforced** (P3), not review-hopes.
- **(5) A round-trip `/v1/access/whoami`-shaped smoke handler** — proves the seam works end-to-end (Principal
  built from a Clerk JWT, delivered to a handler, echoed as a Principal). Behind an internal-only feature flag;
  not customer-visible.

**Out:** the API-key primitive + audit ledger + rate-limit middleware (M302); any customer-data read (M303);
customer UI + docs site (M304).

## Why section (not iterative)
The catalog registry, the DTO, the adapter port, the Clerk adapter, the lint, and the smoke handler are a
**known, enumerable checklist** — each item has a clear "done" state, and the interactions among them are the
kind of layered-scaffolding a `section` milestone is for. Build with `/developer-kit:build-milestone`.

## Depends on / Parallel with
- **Depends on:** **none** — this is R1's spine. It reuses only what already exists in `app` (the Clerk JWT
  middleware, the Sentinel client) — behind the adapter.
- **Parallel with:** **none** — M302, M303, M304 all depend on this seam and are strictly sequential after it.

## KB dependencies
Read as contract:
- [`knowledge/plan/spec-drafts/customer-api-mcp/spec.md`](../../../spec-drafts/customer-api-mcp/spec.md) — the
  consolidated pillar spec (P3, P5, P7, §4.1–§4.3, §5.4).
- [`corpus/services/backend.md`](../../../../../corpus/services/backend.md) — the `app` service the façade lands
  inside.
- [`corpus/services/sentinel.md`](../../../../../corpus/services/sentinel.md) — the authz call the handler asks
  per-resource.
- [`corpus/architecture/external_services.md`](../../../../../corpus/architecture/external_services.md) — Clerk
  as the *current* identity vendor (the swap horizon this milestone contains).

## Re-scope trigger
If the catalog registry's expressive power turns out to fight OpenAPI 3.1's constraints (e.g. we can't cleanly
express the `x-anthropos-*` MCP fields inside a valid OpenAPI document), fall back to a **homegrown YAML
catalog + an OpenAPI generator** — but only after documenting the trigger in `decisions.md`. Never bend the
catalog to fit a doc format at the cost of the P5 "one contract, two shells" principle.

## Delivers →
- The catalog registry (machine source) — read by M302's key-scope UI, by M303's per-resource handler
  scaffolding, and by M304's docs-generator + Workforce UI.
- The `Principal` DTO + `IdentityProvider` adapter port + `ClerkIdentityProvider` — the auth-vendor-swap
  containment surface.
- The `ClerkGuardrails` lint — machine-enforced P3.
- The `/v1/access/whoami` smoke handler (internal-only, feature-flagged) — the seam's end-to-end proof.

