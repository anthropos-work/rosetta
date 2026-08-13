---
iter: 94
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-05
closed: 2026-08-05
---

# iter-94 — the family's own green: one member was reporting a pass over zero docs

**Type:** tik, under `TOK-05`.

## Active strategy reference

`TOK-05`. Answers `CHECK-M257x-iter91-guard-repo-root-scoping`, surfaced by iter-91's own mutation run.

## Step 0 — re-survey

Still live, and it bears directly on clause 5: **the guard family's green is the evidence this milestone
quotes**, so a member that can report a pass it did not earn contaminates every reading taken with it.

## Cluster / target identified

With iter-91's freshness refusal mutated out, the family ran against an empty temp dir and reported
`2 GREEN · 2 RED · 9 could-not-check`. Nine members correctly said *COULD NOT RUN — no corpus/*.
**`story_org_count_guard` and `union_apply_guard` returned GREEN.**

## Hypothesis

At least one of the two is answering about a tree the caller did not name.

## Expected lift

No reading. One fewer way for the family's green to be unearned.

## Phase plan

1. Adjudicate each of the two separately — rext-scoped-by-design is not the same defect as vacuous-pass.
2. Fix what is a defect; record what is correct.

## Escalation conditions

- If the fix would make a legitimately rext-scoped guard fail in a rext-only checkout, stop: the guard
  family is consumed per-stack, not only from the rosetta worktree.
