---
iter: 250
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
---

# iter-250 — the corpus's runnable inputs, split by what a fresh checkout can actually reach

**Type:** tik · **Active strategy:** [`TOK-08`](../decisions.md) — census the mechanical classes; stop
sampling them.

## Step 0 — Re-survey before targeting

iter-249 closed minutes ago with `ROUTE-M257x-249-fresh-checkout-hostile-tests` as its largest finding: 29
tests read untracked local state without declaring it. The obvious next move is to fence that class — and
it is routed. But the same measurement exposed something **on the user's redirect** that is cheaper and
more directly about the corpus, and iter-249 recorded it without following it:

> `fence_command_guard`, on a fresh clone, printed **`a fenced command names a target that does not
> exist`**. `rext_path_guard` printed **`the live corpus must resolve every rext path`**. Both are GREEN on
> this box. Both statements are *about the corpus* and both were caused by an absent `stack-*/` workspace.

So the guards grade a corpus command as resolving when it resolves **against a tree the reader does not
have yet**. That is not a test-hygiene problem like iter-249's; it is a claim the corpus is making about
its own runnability, and nothing measures it. The redirect names exactly this: *be able to build a working
stack*.

Substrate confirmed present at open (from the guard's own docstring, to be re-derived not carried): the
live corpus has 621 fenced blocks / 2,872 runnable lines, of which 149 `cd` + 108 `make` + 69 `npm|pnpm`
are of a checkable shape.

## Cluster / target identified

`TOK-08`, applied to the runnable-input class the milestone has already censused **four times without ever
splitting it**: every graded target is scored `exists / does not exist`, against a filesystem that includes
the operator's `stack-dev/` and `stack-demo/`. The missing axis is **reachability**:

| tier | meaning | who can run it |
|---|---|---|
| **R** | the target is checked into `rosetta` (or `rosetta-extensions`) | anyone, immediately after clone |
| **W** | the target lives inside a `stack-*/` workspace | only after `make init` / `/dev-up` / `/demo-up` |
| **X** | resolves nowhere | nobody — a real defect |

The guards conflate **R** and **W** into one GREEN. A doc could migrate wholly into **W** and no fence
would move.

## Hypothesis

A large share of the corpus's *graded* runnable targets are tier **W**, and the fences' GREEN therefore
carries a precondition none of them states. Making the split explicit is a reach disclosure of the kind
`§5` already requires — *a verdict without its reach is not a verdict* — and it is derivable from data the
guards already compute.

## Pre-registered numeric claims — sealed in this iter's FIRST commit

Stated before any measurement of the split. Graded verbatim at close.

| # | claim | prediction |
|---|---|---|
| **PR-1** | of `fence_command_guard`'s **graded** targets, the share in tier **W** | **≥ 50 %** |
| **PR-2** | ≥ 1 graded target is tier **X** on this box (a genuinely dead path the live-green hides) | **false** — the guard is green here, so X should be empty |
| **PR-3** | `rext_path_guard` likewise grades ≥ 1 tier-**W** path | **true** |
| **PR-4** | either guard's printed verdict already distinguishes **R** from **W** | **false** — that is the gap this iter exists to close |
| **PR-5** | the count of tier-**W** targets is **stable** between the live tree and a fresh clone (i.e. the split is a property of the corpus text, not of the box) | **true** |

**Direction check.** iter-249 refuted four of five predictions, all because I expected the REDs to be about
the corpus. PR-2 and PR-4 are deliberately set to the boring answer for the same reason; PR-1 is the one I
would bet on and is stated as a floor, not a point estimate.

## Phase plan

1. **Seal** this overview as `probe(M257x/250)`.
2. **Derive** the R/W/X split from `fence_command_guard`'s own machinery — never a hand-written re-implementation
   (iter-209's lesson: a hand-rolled second opinion was 16× wrong in one direction).
3. **Repeat** against iter-249's frozen clone pair to grade PR-5.
4. **Disclose**: make the verdict line state the split, so a GREEN says which tier it is green over.
5. **Test** the disclosure, including a control that can actually fire.
6. **Grade** PR-1…PR-5 verbatim and close.

## Escalation conditions

- If the split cannot be derived from the guard's existing structures without re-implementing its parser,
  stop and route — a second parser is the defect this milestone keeps finding.
- If tier **X** is non-empty, that is a real corpus defect: repair it in this iter or route it with a named
  handler, and say which.

## Acceptable close-no-lift outcomes

- The split comes back overwhelmingly **R** (PR-1 refuted): the fences' GREEN is nearly reader-reachable,
  the disclosure is still correct to ship, and the finding is that this worry was unfounded — recorded with
  the number.
