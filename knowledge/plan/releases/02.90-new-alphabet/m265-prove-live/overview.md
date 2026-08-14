---
milestone: M265
title: "Prove it live"
milestone_shape: iterative
status: planned
release: "02.90-new-alphabet"
exit_gate: "On a cold bring-up, ALL of: (1) `/demo-up` green end-to-end on the new canon; (2) `/dev-up` green on the new canon; (3) the full Playthrough suite passing, INCLUDING the net-new taxonomy Playthrough; (4) seed closure green WITH the per-hero richness floor satisfied; (5) `/taxonomy` navigable live."
iteration_protocol_ref: "corpus/ops/verification.md"
depends_on: "M262, M263, M264"
parallel_with: "none"
complexity: large
last_updated: "2026-08-14"
---

# M265: Prove it live

**Goal:** A demo AND a dev stack come up cold on the new canon and prove themselves.

## Exit gate

On a cold bring-up, ALL of: (1) `/demo-up` green end-to-end on the new canon; (2) `/dev-up` green on the new canon; (3) the full Playthrough suite passing, INCLUDING the net-new taxonomy Playthrough; (4) seed closure green WITH the per-hero richness floor satisfied; (5) `/taxonomy` navigable live.

Every clause is measured on a **cold** bring-up. A clause that cannot be measured is **not met** — the v2.8
M258 lesson, where a projection was mistaken for a measurement and the milestone closed off-gate.

## Iteration protocol

[`verification.md`](../../../../corpus/ops/verification.md) — the M258 batch-gate lineage.

## Why iterative (not section)

The failure set of a canon swap is not enumerable up front — it is whatever the first cold run surfaces. v2.1, v2.6 and v2.8 each showed that set is discovered, not predicted.

## Depends on

M262, M263, M264

## Re-scope trigger

If 5 consecutive toks fail to produce a viable strategy, escalate to user-strategic-replan.

## KB dependencies

- [`verification.md`](../../../../corpus/ops/verification.md)
- [`playthroughs.md`](../../../../corpus/ops/demo/playthroughs.md)
- [`rosetta_demo.md`](../../../../corpus/ops/rosetta_demo.md)

## Constraints (release-wide, non-negotiable)

- **Zero platform-repo edits.**
- All stack tooling in `rosetta-extensions` at a pinned tag — **and the tag pushed to origin**, or a remote
  stack cannot obtain it (the M236 pre-flight rung zero).
