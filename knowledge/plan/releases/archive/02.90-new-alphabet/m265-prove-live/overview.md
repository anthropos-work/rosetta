---
milestone: M265
title: "Prove it live"
milestone_shape: iterative
status: archived
release: "02.90-new-alphabet"
exit_gate: "On a cold bring-up, ALL of: (1) `/demo-up` green end-to-end on the new canon; (2) `/dev-up` green on the new canon; (3) the full Playthrough suite passing, INCLUDING the net-new taxonomy Playthrough; (4) seed closure green WITH the per-hero richness floor satisfied; (5) `/taxonomy` navigable live."
iteration_protocol_ref: "corpus/ops/verification.md"
depends_on: "M262, M263, M264"
parallel_with: "none"
complexity: large
last_updated: "2026-08-16"
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

## ~~Routed in from M262~~ — CLOSED 2026-08-15, before this milestone started

**An AI key. RESOLVED** — found in `.agentspace/secrets/studio-desk/.env` (`AI_OPENAI_API_KEY`), not `app/.env`. Both items below are DONE; kept for the record because the search, not the key, is the reusable lesson.

1. **The canon has no embeddings.** `taxonomy-load` ended with *"vectors not computed: the canon is
   loaded but does not take part in matching until this is re-run"* — no embedding manager was
   configured. Browsing, listing and joining all work; **AI skill-matching does not.**
2. **The generated member profiles were not regenerated**, so they still name the old taxonomy's
   skills. `D-v29-3` set a $200 ceiling and price-before-spend; there is nothing to price until a key
   exists.

**Clause 4 of this gate — seed closure green WITH the per-hero richness floor — is what makes this
blocking rather than cosmetic.** A stack whose heroes hold no verified-skill chain is not a proven
stack, and M262's new floor will now say so out loud instead of passing.

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
