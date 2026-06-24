---
milestone: M42m
slug: manager-coverage
version: v1.10 "method acting"
milestone_shape: iterative
exit_gate: "logged in as a MANAGER-vantage hero (e.g. Dan Rossi / Leah Donovan) on demo-3, the Playwright coverage sweep (manager vantage) reports 0 failing pages (no empty/error pages) AND 0 out-of-demo escapes across every manager-reachable surface — including the M36 Workforce-Intelligence dashboard (mapped→verified funnel + teams + role gap/mobility + succession + feedback + the org-scale claimed-vs-verified gap) + any admin pages"
iteration_protocol_ref: corpus/ops/demo/coverage-protocol.md
status: planned
created: 2026-06-24
last_updated: 2026-06-24
complexity: medium-large
delivers: updates across the touched specs as manager-vantage gaps close — stories-spec.md (the M36 Workforce-dashboard surfaces) + frontend-tier.md + any manager-only-route fixes landed into rosetta-extensions
depends_on: M42e (reuses the harness + protocol) + M39 + M40 + M41
spec_ref: .agentspace/profile_gaps.md (live-demo review, 2026-06-24; root-cause workflow w7t4wq2z4)
---

# M42m — Manager 100% demo coverage

## Goal
A hero of the **MANAGER** vantage (e.g. Dan Rossi / Leah Donovan) sees **100% of the demo platform's
manager-reachable pages populated** — no empty pages, no error pages, and no out-of-demo escapes. Logging in
as that manager on demo-3, every surface the manager can reach fills with believable seeded story data and
stays inside the demo world.

## Exit gate (observable, machine-verifiable)
Identical to M42e's gate, run as a **MANAGER** hero. The Playwright coverage sweep, pointed at a manager-roster
hero, must report **0 failing pages** (no empty/error pages) and **0 out-of-demo escapes** across every
manager-reachable surface. This covers the **manager-only** surfaces:
- the **M36 Workforce-Intelligence dashboard** — the mapped→verified funnel + teams + role gap/mobility +
  succession + feedback + the org-scale claimed-vs-verified gap;
- plus **any admin pages** reachable from the manager vantage.

Gate = the Playwright sweep (manager vantage) reports **0 failing pages + 0 escapes**.

## Why iterative (not section)
Same loop as M42e: the manager-vantage **page set and its failure modes are discovered by the sweep, not
enumerable up front**. Which manager-reachable routes exist, which render empty, and which escape the demo
world only surface when the harness crawls the live surface as a manager hero — so the commitment is the gate,
and the fix list emerges per-tik from the sweep's evidence.

## Iteration protocol
`corpus/ops/demo/coverage-protocol.md` — **REUSED** (M42e authors this protocol; M42m reuses the harness +
protocol for the manager vantage). Each iter: run the sweep against a manager hero → triage the failing
pages / escapes → land fixes into `rosetta-extensions` → re-run the sweep → close the iter.

**Re-scope trigger:** same as M42e — escalate a **platform-only blocker** (a failing page whose only fix would
require a platform-repo edit, which this release forbids) to a user-strategic-replan rather than absorbing it.

## Approach sketch
Point the **M42e harness** at a **manager-roster hero** and iterate fixes — likely concentrated in the **M36
Workforce-dashboard surfaces** and manager-only routes — into `rosetta-extensions` until the gate is **GREEN**
for a manager. Fixes are tooling-side only (seeding / set-dress / route handling in rext); the platform repos
(next-web / app / cms / jobsimulation) stay **read-only** (ZERO platform-repo edits — the v1.10 hard line).

## Depends on / Parallel with
- **Depends on:** M42e (reuses the Playwright harness + the coverage protocol it authors) + M39 + M40 + M41
  (the profile-identity / directus-serve / profile-depth groundwork the manager surfaces read from).
- **Parallel with:** none.

## Open questions
- Which manager-roster hero is canonical for the sweep — Dan Rossi vs Leah Donovan (lean: pick one in iter-01
  from the demo-3 roster; the gate is per-hero).
- Whether any "admin page" reachable from the manager vantage is in-scope for *populate* or merely *in-demo*
  (lean: resolve per-iter against what the sweep flags — empty-but-in-demo vs error vs escape).

## KB dependencies
This milestone reads, as contract:
- `corpus/ops/demo/coverage-protocol.md` — the iteration protocol + harness (authored by M42e).
- `corpus/ops/demo/stories-spec.md` — the **M36 Workforce-Intelligence dashboard surfaces** (the
  mapped→verified funnel + teams + role gap/mobility + succession + feedback + the org-scale claimed-vs-verified
  gap) that the manager vantage renders.
- `corpus/ops/demo/frontend-tier.md` — the demo UI tier the sweep crawls.

## Delivers →
Updates across the touched specs **as manager-vantage gaps close**:
- `corpus/ops/demo/stories-spec.md` — the M36 Workforce-dashboard surfaces, revised as fixes land.
- `corpus/ops/demo/frontend-tier.md` — revised for any manager-only-route coverage fixes.
- the manager-vantage fixes themselves land in `rosetta-extensions` (tooling-only; no platform-repo edits).
