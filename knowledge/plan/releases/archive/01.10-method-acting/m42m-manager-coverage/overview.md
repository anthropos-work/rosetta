---
milestone: M42m
slug: manager-coverage
version: v1.10 "method acting"
milestone_shape: iterative
exit_gate: "Identical bar to M42e's re-scoped gate, run as a MANAGER hero (Dan Rossi / Leah Donovan) on a stack from a FRESH demo-up: per page AND per section/element of every manager-reachable surface — (a) REAL semantic content (not placeholder/empty-state copy or bare chrome), (b) SUBSTANTIAL cardinality (not just 1, documented 0/1 exceptions), (c) PERSONA self-consistency (role↔skills↔bio↔real-person avatar↔team↔activity cohere; org has a name+logo), (d) NO prod-eject escape (legitimate external reference links allowed but disclosed). This covers the manager-only surfaces — the M36 Workforce-Intelligence dashboard (mapped→verified funnel + teams + role gap/mobility + succession + feedback + the org-scale claimed-vs-verified gap) + any admin pages, which must show real seeded org/team data, not empty widgets. Gate = 0 sections below the bar + 0 prod-eject escapes, reproduced on a FRESH demo-up."
iteration_protocol_ref: corpus/ops/demo/coverage-protocol.md
status: archived
created: 2026-06-24
last_updated: 2026-06-26
complexity: medium-large
delivers: updates across the touched specs as manager-vantage gaps close — stories-spec.md (the M36 Workforce-dashboard surfaces) + frontend-tier.md + any manager-only-route fixes landed into rosetta-extensions
depends_on: M42e (reuses the harness + protocol) + M39 + M40 + M41
spec_ref: .agentspace/profile_gaps.md (live-demo review, 2026-06-24; root-cause workflow w7t4wq2z4)
---

# M42m — Manager 100% demo coverage

## ⚠ Re-scope (2026-06-25 — inherits M42e's re-scope)
Same correction as M42e: the gate is the **believability bar** (real semantic content + substantial per-section
cardinality + persona/org self-consistency, reproducible on a fresh demo-up), measured by the **new semantic
coverage harness** — NOT the original DOM text-density check. See M42e's overview + [[demo-coverage-semantic-content-gate]].

## Goal
A hero of the **MANAGER** vantage (Dan Rossi / Leah Donovan) sees a **believable, fully-populated manager
experience** — every reachable page and **every section/element** (especially the **M36 Workforce-Intelligence
dashboard**) shows real seeded org/team story data (no empty widgets, no placeholders, no single-row "filler"),
the manager persona + their org/team are **internally coherent**, and nothing ejects the presenter to
production. It must look this way on a **fresh `demo-up`**.

## Exit gate (observable, machine-verifiable)
The **same bar as M42e's re-scoped gate (a)–(d)**, run as a **MANAGER** hero on a stack from a **fresh
`demo-up`** — asserted per page AND **per section/element** of every manager-reachable surface, including the
manager-only ones:
- the **M36 Workforce-Intelligence dashboard** — the mapped→verified funnel + teams + role gap/mobility +
  succession + feedback + the org-scale claimed-vs-verified gap — must render **real seeded org/team data with
  substantial cardinality**, not empty/placeholder widgets;
- plus **any admin pages** reachable from the manager vantage.

**Gate = 0 sections below the content/cardinality/consistency bar + 0 prod-eject escapes (legitimate external
reference links allowed but disclosed), reproduced on a FRESH `demo-up`.**

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
