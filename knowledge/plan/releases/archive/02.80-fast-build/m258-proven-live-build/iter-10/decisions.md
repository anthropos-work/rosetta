# M258 iter-10 — decisions

## D48 — Fix the class the way the repo already fixed it, not a better way.

`census_pruned` exists because M257x iter-277 lost the better part of an hour to a `.venv-check/`
adding **+43 / +42** to two ratchets parked at their ceilings, and its docstring records the rule:
*"the runner must not live inside its own subject."* A per-stack workspace is the same class — a tree
**this repo's own tooling creates, inside itself, that is not its subject** — so the fix is one name
in the existing checked-in registry, not a new mechanism. Component-exact is safe: `dev-stack/stacks/`
and `demo-stack/stacks/` are the only directories named `stacks` in the tree, and both are gitignored
by their own sections.

## D49 — The prediction was stated before the measurement, and it held exactly.

248 / 236 / 657 predicted from the pristine extract; 248 / 236 / 657 measured after the prune. The
value of predicting first is that the alternative — prune, then read whatever comes out — cannot
distinguish "removed the pollution" from "removed too much." **Over-pruning a census fails GREEN**,
which is why the negative control (Phase B) matters more than the positive one.

## D50 — The fourth consumer, and why `population()` was worth fixing rather than routing.

It bypassed the shared helper with a hand-rolled substring filter, and it was the **only** consumer
that surfaced the pollution as a failing test rather than as a quietly-inflated number — it demanded
`DECISIONS` entries for `demo-stack/stacks/demo-1/clones/app/.claude/skills/…`. Fixed rather than
routed because its own docstring already restricts it to *"a non-test **rext** module"*: routing
through `census_pruned` **enforces the stated contract**, it is one line, it uses machinery that
already exists, and it flipped a pre-existing RED to green with no new failures across all four
modules that import the registry (`unclassified` `[]`, `stale_decisions` `[]`, `reach` (20, 109, 192)).

**Not claimed:** that this reaches `decommissioned_instruction_guard` or
`test_fence_provenance::test_the_escape_accepts_and_records`. They carry their own filters and were
**not run**. Routed as `ROUTE-M258-iter10-hand-rolled-path-filters`.

## D51 — A pristine-checkout baseline cannot see an untracked cause.

The first A/B used `git archive HEAD`, which **omits gitignored paths** — so the extract had no
`stacks/` tree and could not express the defect at all. That made a pre-existing failure look like a
regression I had introduced. The correct control for this class is the **working tree with the change
reverted**, which is how it was re-tested and cleared.

Worth stating as a rule because the wrong control was *more* rigorous-looking, not less: it is the
same instrument the ratchet baseline legitimately needs (there, tracked content is exactly the
subject) applied to a question about untracked content.

## D52 — The milestone is ACHIEVED BY USER RULING. Clause 3 is NOT met and must never be recorded as met.

The user ruled mid-iter that M258 has achieved its goal, accepting it on clauses **1, 2, 4, 5** plus
the **~402 s clean projection**, having concluded the CPU contention on this box is not something they
can remove.

**Recorded exactly as: *achieved by user ruling, timing clause unmeasured under load*** — the shape of
M257x's `TOK-09`, and **not** `gate-met`. Specifically, and not to be softened:

- the **840.01 s** contended figure remains **instrument-rejected** (3/3 `headroom=FAIL`,
  `peak_load1` 40.09 / 74.77 / 51.80 against a limit of 10) and stays published as such;
- **401.60 s is a PROJECTION**, composed of iter-05/06's gateable bring-up half and iter-08's
  best-rep contended `batch_gate`. It has never been measured as one cycle;
- no clean p50 over 3 consecutive cold cycles has ever been taken.

The auto-arm **stays armed** (`campaign-iter09/`, `fast-build-m258-iter-09` pin): clause 3 is now an
opportunistic bonus, not the objective.

**New scope, 5 iters:** low-hanging build-time fruit, and a **net-new SPACE axis** — pre-build,
post-build and post-teardown. There has never been a space budget the way there is a time budget.
