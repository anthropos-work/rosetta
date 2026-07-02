---
milestone: M303
slug: rest-reads-gateway
version: v3.0 "open house" (PROPOSAL — awaiting review)
milestone_shape: iterative
exit_gate: "Every resource-family in the pillar spec §4.2 (9 products, 35 resources, ~44 endpoints, ~55 backing tables — Talk-to-Data data parity) has a PASSING contract test on an integration stack: OpenAPI conformance + Principal→org isolation (a Principal from org A gets 0 rows/404 from org B's data on EVERY endpoint) + the applicable CR1–CR15 read-contract rules from spec §4.5 all green + audit-row written + rate-limit fires on burst. The 7 FIRST-USABLE UCs (organization, roster, member get, member skills, sim sessions L+G, learning progress, audit trail) are end-to-end scripted. Static-lint gates for CR7/CR8 (no local_jobsimulation_sessions, no local_skill_path_sessions, no membership_skills.skill_level) are CI-enforced on internal/customerapi/. Sustained bar: 0 cross-tenant leakage over 5 consecutive integration runs across the full ~44-endpoint surface."
iteration_protocol_ref: knowledge/plan/spec-drafts/customer-api-mcp/spec.md (§4.2 catalog; §4.4 UCs; §4.5 read-contract rules CR1–CR15; §5.7 testing posture; §6.3 R1 milestone shape)
status: planned
created: 2026-07-01
last_updated: 2026-07-01
complexity: large
delivers: the R1 READ catalog at **Talk-to-Data data parity** — one `catalog.yaml` entry + one Connect-RPC-or-repository-backed handler + one OpenAPI-generated endpoint per resource-family, closed-per-family on: contract test + cross-tenant isolation test + CR-rule tests + audit row + rate-limit fire. R1 ships **9 products / 35 resources / ~44 endpoints / ~55 backing tables** — People, Assignments, Simulations, Learning, Catalog, Taxonomy, Academy, AI Readiness, Audit, Access.
depends_on: M302
spec_ref: knowledge/plan/spec-drafts/customer-api-mcp/spec.md (the consolidated pillar spec, v0.2 — scope corrected to Talk-to-Data parity)
---

# M303 — REST reads gateway

## Goal
Land the **R1 READ catalog at Talk-to-Data data parity** — every domain a customer can query through the AI chat
surface today (`askengine.TableRegistry` in the platform backend, ~55 tables) becomes a stable, versioned,
principal-scoped REST endpoint under the new `customer-api-mcp` façade. Iterated one **resource-family** at a
time (a family = a parent resource + its nested collection endpoints), each closed on the same per-family gate:
OpenAPI entry + contract test + cross-tenant isolation test + the applicable **CR1–CR15 read-contract rules**
(spec §4.5) + audit-row lands + rate-limit fires. When the last family closes the gate, the pillar has a
**working read half** at parity — a customer with a minted API key (M302) can programmatically answer any
question their AI agent could answer via Talk to Data today, under the same principal + audit + rate-limit
floor.

## Exit gate (objective, machine-verifiable)
**Every resource-family in spec §4.2 has PASSING contract tests on an integration stack**, with:
1. **OpenAPI conformance** — response shape + status codes match the generated `openapi.yaml` (schemathesis or
   equivalent).
2. **Principal → org isolation** — a `Principal` from org A calling ANY of the ~44 endpoints returns 0 rows or
   404 for org B's data (never a leak). Full-surface sweep, not a spot check.
3. **CR1–CR15 rule matrix** (spec §4.5) — every applicable rule test green on the relevant endpoint. Fixtures
   include: soft-deleted member (CR2), `levels_count = 7` org (CR6), `timedout` sim session (CR4),
   partial-translation skill (CR11), `draft` academy chapter (CR14), non-admin profile-history request (CR13),
   live vs frozen ai-readiness (CR12).
4. **Static-lint gates** (CR7, CR8) — CI grep-check on `internal/customerapi/` fails the build on any reference
   to `local_jobsimulation_sessions`, `local_skill_path_sessions`, or `membership_skills.skill_level`.
5. **Audit row written** — every request (success + failure) lands a row in `customer_api.audit_events` with the
   correct `principal_id` + `organization_id` + `resource_id` + `action`.
6. **Rate-limit fires** — a burst above the default budget returns 429 with `Retry-After` + `X-RateLimit-*`
   headers.
7. **FIRST-USABLE seven scripted** — an integration script exercises the seven ∗-marked endpoints end-to-end
   (organization → roster → member → member skills → sim session list → sim session get → learning progress →
   audit trail) under one minted API key.

**Sustained bar:** 0 cross-tenant leakage over **5 consecutive integration runs across the full ~44-endpoint
surface** (borrowing the M203 zero-false-fail posture — determinism under real data, at parity scale).

## Why iterative (not section)
The ~44 endpoints across 9 products are **declarable** (spec §4.2), but each resource-family's projection from
its Talk-to-Data backing table(s) surfaces its own failure modes — a missing filter, a shape mismatch, a
Sentinel policy that doesn't scope to `organization_id`, an `askengine`-specific business rule (a CR) that
customers must inherit but the naive Connect-RPC doesn't enforce, a translation join that only fires on
`?language=`, a nested collection whose parent-id path parameter needs to be validated for tenant scope. Each
iter picks the next resource-family → writes its `catalog.yaml` entry → wires the handler (with the applicable
CR-rule enforcement) → the contract test + CR-rule tests + isolation test are exercised → diagnose → fix
(handler / catalog / audit / policy / repository predicate) → close-per-family, until all families are GREEN.
Build with `/developer-kit:build-mstone-iters`.

## Iteration protocol
The loop is the pillar spec: [`knowledge/plan/spec-drafts/customer-api-mcp/spec.md`](../../../spec-drafts/customer-api-mcp/spec.md)
§4.2 (catalog), §4.4 (35 UCs), §4.5 (CR1–CR15 read-contract rules), §5.7 (testing posture), §6.3 (R1 MVP scope).
Appendix A gap analysis applies to writes (R4/R5); for reads, "gap" means "no existing internal query carries
the shape we need at parity" — that surfaces as a **catalog-entry escalation** (with a decision.md record),
never a shim.

**Per-family close bar** (the closed-per-resource loop):
1. `catalog.yaml` entry lands + validator green.
2. Handler compiles + unit-tests green (adapter to existing Connect-RPC or repository query).
3. OpenAPI regenerated + committed.
4. Contract test written + green.
5. **CR-rule tests** written + green for every CR that applies to the family's shape (e.g. sim-session family
   → CR4 + CR6; member-skill family → CR5 + CR6 + CR7; academy family → CR14; profile-history → CR13).
6. Cross-tenant isolation test written + green (the load-bearing safety test).
7. Integration run confirms audit row + rate-limit headers.

**Suggested iter ordering** (deliverable-first, then depth-first per family):
1. **Iter 1** — `people.organization` (the base — needed by CR6 `max_level`).
2. **Iter 2** — `people.member` list + get (FIRST-USABLE; sets the roster + isolation baseline).
3. **Iter 3** — `people.member.skill` (CR5 mapped-vs-verified, CR6 org-scale, CR7 source-column — the highest-
   risk read).
4. **Iter 4–7** — remaining `people.member.*` families (cert/edu/exp/language/target-role/tag/profile-history).
5. **Iter 8** — `people.team`, `people.invitation`, `people.company`.
6. **Iter 9** — `assignments.*` family.
7. **Iter 10** — `simulations.simulation-session` list + get (FIRST-USABLE; CR4 completed-definition, CR6 score-
   on-org-scale).
8. **Iter 11–15** — nested sim-session collections (recording, interaction, realtime-call, code-submission,
   anticheat, activity-event, task-check, conversation-extraction, interview-extraction, validation-result,
   validation-attempt).
9. **Iter 16** — `simulations.simulation-feedback`.
10. **Iter 17** — `learning.skill-path-session` (FIRST-USABLE).
11. **Iter 18** — `catalog.*` (simulation-template + skill-path-template; CR11 localization).
12. **Iter 19** — `taxonomy.*` (skill + job-role + world-language; CR10 catalog resolution, CR11 localization).
13. **Iter 20** — `academy.*` (series + skill-path + chapter + progress; CR14 visibility).
14. **Iter 21** — `ai-readiness.*` (live + cycle + cycle.snapshot; CR12 live-vs-frozen — split as three distinct
    resources).
15. **Iter 22** — `audit.audit-event` (FIRST-USABLE; closes the audit self-serve loop from M302).

The count of iters is not the exit criterion; the exit-gate items are. If a family absorbs into a prior one
(e.g. `academy.chapter-body` folds under `chapter`), count it and move on.

## Re-scope trigger
A resource-family that **can't** be adapted without touching a shared Connect-RPC contract (e.g., an existing
endpoint returns cross-org data by default and would need a filter added upstream) → **escalate, don't wedge a
filter inside the façade** (would poison the isolation guarantee). Log to `decisions.md`, close the iter with a
scoped gap, and reopen when the upstream fix lands. Mirrors M203's `unimplementable-without-platform-edit`
escalation posture — but here platform edits *are* allowed (this is the platform), just not silent hacks in the
façade.

**Rule-fidelity re-scope:** if a CR1–CR15 rule cannot be honored without a schema change (unlikely — the rules
are all read-side predicates), escalate before proceeding. A rule that "mostly holds" is a rule that has failed.

If Redis rate-limit fights the shared Redis budget (per M302's re-scope trigger) — inherit that fallback.

## Depends on / Parallel with
- **Depends on:** **M302** — needs the API-key `Principal`, the `ApiKeyIdentityProvider`, the audit-write
  middleware, and the rate-limit middleware. Without them the per-family close bar is not verifiable.
- **Parallel with:** **M304 (docs-lite)** — the customer UI + docs site can develop against the M302
  `/v1/access/whoami` echo + the M303-in-progress catalog snapshots (a customer-visible endpoint list). Docs
  publish in lockstep as families close their gates. The docs site's entitlement-tier page is auto-generated
  from `catalog.yaml`.

## KB dependencies
Read as contract:
- [`knowledge/plan/spec-drafts/customer-api-mcp/spec.md`](../../../spec-drafts/customer-api-mcp/spec.md) — the
  pillar spec (§3 principles, §4.1 resource hierarchy, §4.2 catalog, §4.4 UCs, §4.5 CR1–CR15, §5 tech approach,
  §5.7 testing posture, §6.3 R1 scope, Appendix A gap analysis).
- **`ant-platform-backend/internal/askengine/registry.go`** (~55 tables) + **`askengine/rules.md`** (Universal
  Rules 1–24 + Cross-Validation + Common Mistakes) — the **authoritative parity source** every iter references.
  A read that disagrees with `askengine` on a filter, a join, or a rule is wrong.
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
- The **R1 read catalog at Talk-to-Data data parity** — 9 products / 35 resources / ~44 endpoints / ~55 backing
  tables, each with its `catalog.yaml` entry + handler + generated OpenAPI + contract test + CR-rule tests +
  cross-tenant isolation test.
- The **first customer-visible read at parity** — a minted API key gets a real answer from any of the ~44
  endpoints, respecting every read-contract rule the internal AI chat surface honors.
- The **isolation-guarantee proof** — the load-bearing "org A can't see org B's data" test at every endpoint,
  run per-integration-run, 0 false-fails over 5 consecutive runs at close.
- The **CR1–CR15 rule matrix green** — the ~15 read-contract invariants encoded as contract tests + static-lint
  gates, so the same bug class doesn't re-emerge under a customer principal (higher blast radius).
- The **first audit-populated stream at scale** — real request rows land in `customer_api.audit_events` across
  ~44 endpoints, ready for the audit-event self-serve read.
