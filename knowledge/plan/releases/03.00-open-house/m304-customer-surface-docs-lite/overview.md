---
milestone: M304
slug: customer-surface-docs-lite
version: v3.0 "open house" (PROPOSAL — awaiting review)
milestone_shape: section
status: planned
created: 2026-07-01
last_updated: 2026-07-01
complexity: medium
delivers: the customer-facing Workforce UI for API-key management (list / mint / rotate / revoke via the M302 admin catalog entries) + the docs.anthropos.work/api/v1/ static site (OpenAPI-rendered reference + a Quickstart per FIRST-USABLE UC + the Principles page + the entitlement-tier page). No new backend surface — this milestone consumes M302 + M303.
depends_on: M303
spec_ref: knowledge/plan/spec-drafts/customer-api-mcp/spec.md (the consolidated pillar spec, v0.1)
---

# M304 — Customer surface + docs-lite

## Goal
Turn the R1 read catalog into a **product a developer can actually pick up**: a self-serve UI in Workforce
(list keys, mint, rotate, revoke — driven by the M302 admin catalog entries; Clerk-authenticated in Workforce)
+ a public reference site at `docs.anthropos.work/api/v1/` (OpenAPI-rendered reference + Quickstarts +
Principles + entitlement tiers). "Docs-lite" — reference + Quickstarts + entitlement page, not the full
developer portal (that's a later minor).

R1 becomes visible + usable by a real customer at close of this milestone.

## Scope
**In:**
- **(1) Workforce API-key page** — a new page in `next-web-app` under Enterprise settings: list this org's API
  keys (labels, prefix, scopes, tier, created_at, last_used_at, revoked_at) + [Mint new key] + [Rotate] +
  [Revoke]. The plaintext key is shown **once**, on mint, with a copy-to-clipboard button + a "you can't see
  this again" banner. All UI actions call the M302 admin catalog entries, authenticated as the current Clerk
  user with the admin scope.
- **(2) Docs site skeleton** — a new static site at `docs.anthropos.work/api/v1/`
  (Docusaurus / MkDocs / stoplight — deferred, decided in `decisions.md`). Reads: the generated `openapi.yaml`
  (from M303) + a fixed set of hand-authored MDX pages.
- **(3) OpenAPI reference** — the auto-rendered reference for every R1 resource (from the M303-generated
  `openapi.yaml`). Includes: request/response shapes, auth-header format, `X-RateLimit-*` header contract, the
  4 error shapes.
- **(4) Quickstart per FIRST-USABLE UC** — 4 short recipes (UC1 list members, UC2 get member, UC3 list skill
  paths, UC4 get skill path), each with a `curl` example, a Python snippet, a Node snippet. ~1 page each.
- **(5) Principles page** — the 9 principles (P1–P9) from spec §3, in customer-friendly prose.
- **(6) Entitlement-tier page** — the 4 tiers (free / paying / enterprise / partner) + which resources each
  can call + default rate-limits. Sourced from `catalog.yaml`'s `audience:` field (not hand-maintained).

**Out:** the MCP server (fast-follow after R1, spec §5.2); AI-agent-specific docs (later minor); a change-log
publishing pipeline (R2); interactive OpenAPI playground (R2); SDK downloads (later minor); write-catalog docs
(R2 W1 / R3 W2).

## Why section (not iterative)
The 6 pieces are well-scoped consumers of M302 + M303 outputs — no discovery, no exploration. Standard
frontend page + a static site generator + auto-render from `openapi.yaml`. A `section` build.
`/developer-kit:build-milestone`.

## Depends on / Parallel with
- **Depends on:** **M303** — the docs site auto-renders the R1 `openapi.yaml`; the Quickstarts assert against
  the actual R1 endpoints in a smoke-test job.
- **Depends on (via M303 chain):** **M302** — the API-key UI is a thin frontend over the M302 admin catalog
  entries (`access.api-key.create` / `.rotate` / `.revoke`) + the M302 list endpoint (`/v1/access/api-keys`).
- **Parallel with:** **none** in R1.

## KB dependencies
Read as contract:
- [`knowledge/plan/spec-drafts/customer-api-mcp/spec.md`](../../../spec-drafts/customer-api-mcp/spec.md) — the
  pillar spec (§3 principles, §4 use-case catalog, §6.3 R1 scope). The Principles page + entitlement page + the
  UC Quickstart set trace back to this doc.
- [`corpus/services/next-web-app.md`](../../../../../corpus/services/next-web-app.md) — the Workforce shell
  the API-key page lands inside.
- [`corpus/architecture/frontend_architecture.md`](../../../../../corpus/architecture/frontend_architecture.md)
  — the Next.js monorepo conventions the new page follows.

## Re-scope trigger
If the docs site tooling choice (Docusaurus vs stoplight vs MkDocs) drags — **pick the simplest one that
auto-renders OpenAPI cleanly + supports MDX**, ship a static build, and defer polish to R2. Never ship without
a public reference URL.

## Delivers →
- The **Workforce API-key page** — the customer's key-management surface (list / mint / rotate / revoke).
- The **`docs.anthropos.work/api/v1/` site** — public reference + Quickstarts + Principles + entitlement tiers.
- The **first developer-facing artifact** — a URL a customer can send to their team.
- The **R1 close** — this is the last milestone of v3.0 R1; when it closes, `/developer-kit:close-release`
  runs on the codenamed release.
