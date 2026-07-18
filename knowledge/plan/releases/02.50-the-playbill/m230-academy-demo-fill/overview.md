---
milestone: M230
slug: academy-demo-fill
version: v2.5 "the playbill"
milestone_shape: iterative
status: planned
created: 2026-07-19
last_updated: 2026-07-19
depends_on: M229
exit_gate: "On a cold /demo-up, the academy home grid renders real cards (count >= floor) for the employee vantage, NO 'Draft' chip, via the real DB-authoritative GraphQL path (or a faithful equivalent), 0 prod-ejects, verified by the coverage sweep on a RENDERED-CARD COUNT (not the M53 port-serves + SSR-title check that let F4 slip)."
iteration_protocol_ref: corpus/ops/verification.md + corpus/ops/demo/coverage-protocol.md
delivers: corpus/ops/demo/frontend-tier.md
---

# M230 — academy demo fill

**Status:** `planned`  ·  **Shape:** `iterative`  ·  **Complexity:** medium  ·  **Depends on:** M229

## Goal
Make the demo (and dev) ant-academy home grid render REAL academy content the way taxonomy/skill-path do — PRODUCTION-FAITHFUL, no 'Draft' chip (user decision) — closing the year-old F4 carry inside the zero-platform-edit wall.

## Exit gate
On a cold /demo-up, the academy home grid renders real cards (count >= floor) for the employee vantage, NO 'Draft' chip, via the real DB-authoritative GraphQL path (or a faithful equivalent), 0 prod-ejects, verified by the coverage sweep on a RENDERED-CARD COUNT (not the M53 port-serves + SSR-title check that let F4 slip).

**Iteration protocol:** `corpus/ops/verification.md + corpus/ops/demo/coverage-protocol.md`

## Scope
### In
- First tik decides the faithful path: Option C (sha-pinned rext demo-patch restoring the M7 FS-as-published fallback on the ephemeral clone) vs Option B (a net-new firewalled academy-content snapshot surface, capture->replay + wire endpoint + compose subgraph). Draft-layer Option A REJECTED (visible chip).
- Verify the full committed catalog merges/renders + chapter bodies serve unlocked
- Correct frontend-tier.md's F4 attribution + document the shipped fill mechanism

### Out
- Any ant-academy platform-repo edit (routes to a demo-patch or escalates)
- An academy SESSION/progress model (Thread-B concern)

## Delivers
`corpus/ops/demo/frontend-tier.md`

## Open questions
- Does prod academy content live in app internal/academy as firewallable public rows (Option B), and what is its public predicate?
- Is a demo-patch (Option C) sufficient + revert-clean?

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.5 "the playbill" for the authoritative milestone design + the release-level decisions/risks (research provenance: `.agentspace/scratch/roadmap-research-2026-07-19` via the design-content-stories-research workflow).
