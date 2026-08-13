---
milestone: M258
iter: 1
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
created: 2026-08-12
---

# M258 iter-01 — bootstrap tok

**Type:** tok (bootstrap) · **Authors:** `TOK-01` · **Does NOT terminate the call** — the loop continues
into iter-02 as a tik under the strategy authored here.

## Inputs

`overview.md` (scope, `exit_gate`, the world contract, the inherited lists) · the milestone's
`iteration_protocol_ref` (`corpus/ops/verification.md`, the prove-on-billion lineage) ·
`roadmap.md` § v2.8 · M257's `progress.md` (the composition's measured first half) · M256's
`progress.md` (the second half's only timing evidence) · M257x's `carry-forward.md` (5 clusters) ·
the iter-01 Phase-0b audit (`../kb-fidelity-audit.md`, **YELLOW**).

## What the bootstrap tok must settle

1. **The world contract** — `overview.md` says explicitly it *"MUST be decided at iter-01"* and offers
   two pre-authorised resolutions. Picking one is not a re-plan; declining to pick is the failure.
2. **The opening strategy** — what the first batch of tiks does, in what order, and why that order.
3. **The known-context** the audit surfaced, carried forward rather than rediscovered.

## Non-goals for this iter

No metric movement (a tok does not move the gate). No bring-up. No code change beyond what the
Phase-0b audit already applied as stale-reference repair.

## Phase plan

- Phase 0b (done, before this plan): `/developer-kit:audit-kb-fidelity` → **YELLOW**, fixes applied.
- Phase 0d: **SKIPPED** — a strategy-authoring iter wires no artifacts through a gate pipeline.
- Phase 1: author `TOK-01` (this iter's deliverable).
- Phase 3: no metric delta; record the signal + the next-tik direction.
- Phase 4: close, append `TOK-01` to the **milestone-root** `decisions.md`, commit.

## Escalation conditions

Escalate only if the world contract turns out to have **no** admissible resolution (both (a) and (b)
refuted by evidence) — that would be a genuine re-cut. A choice *between* two pre-authorised options is
this iter's job, not the user's.
