---
iteration_type: tik
status: in-progress
opened: 2026-08-11
active_strategy: TOK-02
---

# iter-08 — `BASELINE-M257-macmini-n3`, take two: from the PINNED clone

**Type:** tik · **Active strategy:** `TOK-02` step 3 — *take the baseline on the contended box, and label it*

## Step 0 — Re-survey before targeting (mandatory)

`TOK-02`'s **Next-tik direction** named iter-06 (the units fix, DONE) and then the campaign. iter-07 ran the
campaign and produced a number that is **not** a baseline; the milestone `progress.md` next-iter row names
**iter-08 = the campaign take two, from a pinned clone**, blocked on a publish. Re-surveyed at open rather
than trusted:

| re-survey check | 2026-08-11 21:2xZ | verdict |
|---|---|---|
| `gated_baseline` in `macmini.json` | still **absent** | target still owed ✅ |
| `demo-1` slot | **no containers** (`docker ps`) | free; the user's `demo-2` (11) + dev (5) untouched ✅ |
| rext authoring copy | clean, **14 commits ahead** of `origin/main`, real writable GitHub remote | publish is available ✅ |
| the arriving dirty file | `corpus/ops/demo/build-budget.md`, iter-07's correct edit | **must be committed (Priority 1)** |
| the two fence REDs | reproduced at open | see below — **guard defect, and a NEW mechanism** |

**The two REDs are NOT what iter-07 characterised them as, and the difference matters.** iter-07 verified the
cited content at **worktree and `HEAD`** in both declared clone roots and concluded the guard resolves to
content *"neither of its own declared clone roots contains."* Re-measured here: the guard names its ref in
its own output — **`read at origin/main@0a9370c`** and **`read at origin/main@766df6c`** — and those are
**remote-tracking refs that are AHEAD of the checkouts**. `stack-demo/app` HEAD `3eaadae68` vs `origin/main`
`0a9370c24`; `stack-demo/platform` HEAD `0c91421` vs `origin/main` `766df6c`. The content IS in a declared
clone root; it is at the root's **checkout**, and the guard graded the **fetched upstream**.

## Cluster / target identified

**Primary target: `BASELINE-M257-macmini-n3`** — unchanged since `TOK-02` step 3, still the only thing
between this milestone and levers. It is reached through three planned, declared steps (below), the first two
of which are the preconditions iter-07 verified and could not take.

## Hypothesis

1. The two fence REDs are a **guard-resolution defect of the corpus's own governing class** — *"cite the sha,
   never the moving label"* — and the fix is to extend an acquittal rule the guard **already implements** for
   block-named refs (iter-100) to the **default ladder's own committed rungs**. An UNPINNED citation names no
   ref; a guard has no warrant to pick one and grade against it.
2. With `rosetta-extensions` tagged and pushed and `stack-demo/rosetta-extensions` re-pinned, the campaign run
   **from that consumption clone** satisfies G6 arm 1 and the `postgres-schemas` probe — the single root cause
   of all three of iter-07's disqualifications.

## Expected lift

- Step 1: the pre-commit fence goes GREEN **without editing a word of the prose it flagged**; the arriving
  correct corpus edit commits. A negative control proves the guard still books a genuinely-rotted anchor.
- Step 3: a campaign whose `autoverify` verdict is about the **stack** rather than about where the harness was
  invoked from. `gated_baseline` filled **only if the run qualifies** — `green:true / 0 warnings`, patches
  applied, HEADROOM read and reported. A refusal is a RESULT.

## Phase plan — a DECLARED multi-step shape (scope-creep tripwire counts against THIS list)

1. **`FIX-M257-anchor-guard-resolution`** + commit `corpus/ops/demo/build-budget.md` (Priority 1).
2. **Publish `rosetta-extensions`**: full guard sweep to **completion** (the sweep iter-07 killed mid-run —
   D2), then tag + `git push --tags` (Priority 2, durably authorised; `CLAUDE.md:154`).
3. **`BASELINE-M257-macmini-n3` take two**: re-pin `stack-demo/rosetta-extensions`, verify the four
   preconditions live, run `n ≥ 3` from that clone, every rep labelled with its `load1`.

Anything else that surfaces is **unplanned** and routes forward.

## Escalation conditions

- The guard REDs turn out to be a **corpus** defect after all → fix the corpus, not the guard.
- The sweep finds a RED that the tag would publish → **do not push**; fix or escalate.
- The campaign's blocker turns out to be something the pinned clone does not fix → close on the finding, route
  the baseline forward again, do **not** invent a number.

## Acceptable close-no-lift outcomes

- A campaign that runs and is **refused** by HEADROOM on a contended box is a **result**, recorded with its
  `load1` — not a failure to measure (`TOK-02` step 3, explicit).
- `gated_baseline` left empty **again**, if and only if the run does not qualify and the reason is measured.
