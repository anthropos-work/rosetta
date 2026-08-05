---
iter: 91
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-05
closed: 2026-08-05
---

# iter-91 — grade the cannot-tell: the stale-clone class, and the guard-pair enumeration

**Type:** tik, under `TOK-05`.

## Active strategy reference

`TOK-05: stop repairing claims; fence the predicates under them`. Both halves of this iter are predicate
work: the predicate *"this family reading is a measurement"* is fenced by making an unmeasurable run say so,
and the predicate *"these guards were tested"* is fenced by enumerating the PAIRS rather than the members.

## Step 0 — re-survey

iter-90 routed three items here and they are all still live: `FENCE-M257x-iter91-clone-freshness`,
`CHECK-M257x-iter90-revert-idempotency`, `CHECK-M257x-iter90-realmanifest-baseline`, plus Decision 1 item 3
(the 7-guard conjunction-pair sweep). Re-measured at open: the guard family reads **15 GREEN · 0 RED** at
iter-90's commit with every `stack-demo` clone fetched, so the corpus precondition for a reading is intact
and the blocker is instrument-side, exactly as routed.

## Cluster / target identified

The user's standing correction — *treat any guard result as INVALID unless the clone was fetched first* —
and their question: should `guard_family.py` **refuse to run** against a stale clone rather than answering
from what it can see?

iter-90 measured the mechanism underneath: `platform_alignment_guard` resolves citations at `origin/main`,
then `HEAD`, then **silently at the worktree**, and grades `unresolvable` as nothing at all. The two
references disagree — `auto` reads GREEN/0-unresolvable, the worktree fallback reads RED/8 findings — so
*which reference answered* is not a detail, it is the verdict.

## Hypothesis

The cannot-tell must be graded at **two** levels, and the split is the deliverable: at the point of use
(only the guard knows which refs it needs) and at the runner (only the runner can record the reference the
whole family was taken against).

## Expected lift

No movement on the clause-5 reading. The deliverable is that an unmeasurable guard run can no longer report
a verdict, and that every family transcript from now on names its commits.

## Phase plan (declared multi-step shape)

1. `platform_alignment_guard`: count the silent worktree fallback, grade it and `unresolvable` as
   **UNMEASURED (exit 2)**.
2. `guard_family.py`: print the reference (both shas) on every run; refuse `exit 2` when a platform-facing
   run has no `origin/main`; add opt-in `--verify-remote` for the only honest staleness check.
3. Enumerate the **7-guard interacting pairs** (Decision 1 item 3), land what is cheap, route the rest with
   named CHECK ids.

## Escalation conditions

- If refusing on staleness cannot be made both locally-decidable and non-vacuous, **say so with the
  measurement** rather than shipping a heuristic that fires on the wrong thing.

## Acceptable close-no-lift outcomes

A measured demonstration that the family-level refusal is the wrong layer — with the reasoning — is a
complete iter, provided the point-of-use fix lands.
