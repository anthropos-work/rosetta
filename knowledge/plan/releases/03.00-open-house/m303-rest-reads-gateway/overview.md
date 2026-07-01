---
milestone: M303
slug: rest-reads-gateway
version: v3.0 "open house" (PROPOSAL — awaiting review)
milestone_shape: iterative
exit_gate: "Every declared R1 read use case (UC1 people.member.list, UC2 people.member.get, UC3 learning.skill-path.list, UC4 learning.skill-path.get, UC5 learning.session.list, UC6 verification.verified-skill.list, UC7 audit.event.list) has a PASSING contract test on an integration stack: OpenAPI conformance + Principal→org isolation (a Principal from org A gets 0 rows/404 from org B's data) + audit-row written + rate-limit fires on burst — 0 cross-tenant leakage over 5 consecutive integration runs."
iteration_protocol_ref: knowledge/plan/spec-drafts/customer-api-mcp/spec.md (§4 use-case catalog; §6.3 R1 MVP scope; the closed-per-resource loop)
status: planned
created: 2026-07-01
last_updated: 2026-07-01
complexity: large
delivers: the R1 READ catalog — one `catalog.yaml` entry + one Connect-RPC-backed handler + one OpenAPI-generated endpoint per resource, closed-per-resource on: contract test + cross-tenant isolation test + audit row + rate-limit fire. R1 ships **7 resources across 4 products** — People (member.list + member.get), Learning (skill-path.list + skill-path.get + session.list), Verification (verified-skill.list), Audit (event.list — the customer's own audit trail, feature-flagged).
depends_on: M302
spec_ref: knowledge/plan/spec-drafts/customer-api-mcp/spec.md (the consolidated pillar spec, v0.1)
---

# M303 — REST reads gateway

## Goal
Land the **R1 READ catalog** — every FIRST-USABLE (UC1–UC4) + FIRST-USEFUL (UC5–UC7) read resource — as a
production endpoint under the new `customer-api-mcp` façade, each closed on the same per-resource gate:
OpenAPI entry + contract test + cross-tenant isolation test + audit-row lands + rate-limit fires. When the
last resource closes the gate, the pillar has a **working read half** — a customer with a minted API key
(M302) can list/get their org's people, skill paths, sessions, verified skills, and their own audit trail.

## Exit gate (objective, machine-verifiable)
**Every declared R1 read use case has a PASSING contract test on an integration stack**, with:
1. **OpenAPI conformance** — response shape + status codes match the generated `openapi.yaml` (schemathesis or
   equivalent).
2. **Principal → org isolation** — a `Principal` from org A calling any UC returns 0 rows or 404 for org B's
   data (never a leak).
3. **Audit row written** — every request (success + failure) lands a row in `customer_api.audit_events` with the
   correct `principal_id` + `organization_id` + `resource_id` + `action`.
4. **Rate-limit fires** — a burst above the default budget returns 429 with `Retry-After` + `X-RateLimit-*`
   headers.

**Sustained bar:** 0 cross-tenant leakage over 5 consecutive integration runs (borrowing the M203 zero-false-fail
posture — determinism under real data).

## Why iterative (not section)
The 7 resources are **declarable** (UC1–UC7), but each one's Connect-RPC-to-REST projection surfaces its own
failure modes — a missing filter, a shape mismatch, a Sentinel policy that doesn't scope to `organization_id`,
a Redis-Streams write path we thought was read-side but isn't, an audit-row schema that needs a new field for a
list-response. Each iter picks the next resource → writes its `catalog.yaml` entry → wires the handler → the
contract test fails → diagnose → fix (adapter code / catalog / audit / policy) → close-per-resource, until all
7 gates are GREEN. Build with `/developer-kit:build-mstone-iters`.

## Iteration protocol
The loop is the pillar spec: [`knowledge/plan/spec-drafts/customer-api-mcp/spec.md`](../../../spec-drafts/customer-api-mcp/spec.md)
§4 (use-case catalog, the FIRST-USABLE flags on UC1–UC4) + §6.3 (R1 MVP scope) + Appendix A (the real-mutation
gap analysis — for reads, "gap" means "no existing Connect-RPC that carries the shape we need"; a missing read
surfaces as a **catalog-entry escalation**, never a shim).

**Per-iter close bar** (the closed-per-resource loop):
1. `catalog.yaml` entry lands + validator green.
2. Handler compiles + unit-tests green (adapter to existing Connect-RPC).
3. OpenAPI regenerated + committed.
4. Contract test written + green.
5. Cross-tenant isolation test written + green (the load-bearing safety test).
6. Integration run confirms audit row + rate-limit headers.

## Re-scope trigger
A resource that **can't** be adapted without touching a shared Connect-RPC contract (e.g., an existing endpoint
returns cross-org data by default and would need a filter added upstream) → **escalate, don't wedge a filter
inside the façade** (would poison the isolation guarantee). Log to `decisions.md`, close the iter with a scoped
gap, and reopen when the upstream fix lands. Mirrors M203's `unimplementable-without-platform-edit` escalation
posture — but here platform edits *are* allowed (this is the platform), just not silent hacks in the façade.

If Redis rate-limit fights the shared Redis budget (per M302's re-scope trigger) — inherit that fallback.

## Depends on / Parallel with
- **Depends on:** **M302** — needs the API-key `Principal`, the `ApiKeyIdentityProvider`, the audit-write
  middleware, and the rate-limit middleware. Without them the per-resource close bar is not verifiable.
- **Parallel with:** **M304 (docs-lite)** — the customer UI + docs site can develop against the M302
  `/v1/access/whoami` echo + the M303-in-progress catalog snapshots (a customer-visible endpoint list). Docs
  publish in lockstep as resources close their gates.

## KB dependencies
Read as contract:
- [`knowledge/plan/spec-drafts/customer-api-mcp/spec.md`](../../../spec-drafts/customer-api-mcp/spec.md) — the
  pillar spec (§3 principles P1 real-mutations-only + P4 catalog-first + P5 principal-scoped, §4 use-case
  catalog, §5 tech approach, §6.3 R1 scope, Appendix A gap analysis).
- [`corpus/services/backend.md`](../../../../../corpus/services/backend.md) — the `app` service the façade lands
  inside.
- [`corpus/services/skiller.md`](../../../../../corpus/services/skiller.md), [`corpus/services/skillpath.md`](../../../../../corpus/services/skillpath.md),
  [`corpus/services/cms.md`](../../../../../corpus/services/cms.md) — the read sources the R1 catalog projects
  from.
- [`corpus/architecture/security_compliance.md`](../../../../../corpus/architecture/security_compliance.md) —
  the tenant-isolation posture the isolation test asserts against.
- [`corpus/architecture/dependency_map.md`](../../../../../corpus/architecture/dependency_map.md) — the
  Connect-RPC surfaces the handlers wrap.

## Delivers →
- The **R1 read catalog** — 7 endpoints across 4 products (People, Learning, Verification, Audit), each with
  its `catalog.yaml` entry + handler + generated OpenAPI + contract test + cross-tenant isolation test.
- The **first customer-visible read** — a minted API key gets a real answer from `GET /v1/people/members`.
- The **isolation-guarantee proof** — the load-bearing "org A can't see org B's data" test at every resource,
  run per-integration-run, 0 false-fails over 5 consecutive runs at close.
- The **first audit-populated stream** — real request rows land in `customer_api.audit_events`, ready for
  UC7's `/v1/audit/events` read.
