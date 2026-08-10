---
iter: 251
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
---

# iter-251 — the same defect, one guard over: the corpus's own file citations graded by existence

**Type:** tik · **Active strategy:** [`TOK-08`](../decisions.md) — census the mechanical classes; stop
sampling them.

## Step 0 — Re-survey before targeting

iter-250 closed minutes ago on `rext_path_guard`, and routed its own generalisation:
`ROUTE-M257x-250-the-runtime-bucket-is-one-guard-wide` — *"every other fence that grades a corpus path by
existence has the same defect"*, naming `corpus_citation_guard` first. Confirmed by reading, before
targeting: `corpus_citation_guard.py:232`, `:255`, `:262` are three `if not tp.exists()` branches, one per
clause (**C1** link target, **C1** backticked path, **C3** line-pinned path). Its subject is the `rosetta`
tree, where `stack-*/` and `.agentspace/` are git-ignored — so the corpus's citations into a stack
workspace resolve here and cannot resolve for a reader.

This is the largest citation fence in the family (iter-221: *"the corpus cites 2,117 files and nothing
checked that any of them exist"*), so if the class is there, it is there at scale.

## Cluster / target identified

`TOK-08` applied to the class iter-250 named and could only close for one guard.
`D-M257x-250-1` is the decided repair shape and it transfers verbatim: **partition by
`git check-ignore`, which decides pathnames whether or not they exist, so the answer comes from the tree.**

## Hypothesis

`corpus_citation_guard`'s GREEN is operator-dependent in the same way: some graded citations resolve only
because this box carries `stack-*/` and `.agentspace/`, and on a fresh checkout the guard reports them as
corpus defects.

## Pre-registered numeric claims — sealed in this iter's FIRST commit

| # | claim | prediction |
|---|---|---|
| **PR-1** | ≥ 1 graded citation resolves only because of state git ignores in `rosetta` | **true** |
| **PR-2** | the count of such citations is **≥ 10** | **true** |
| **PR-3** | run on iter-249's frozen clone, the guard emits ≥ 1 C1/C3 finding that is **not** a real defect | **true** |
| **PR-4** | after the repair, live and frozen give the **identical** verdict | **true** |
| **PR-5** | ≥ 1 of these citations is genuinely **WRONG** — a real defect the operator's tree was hiding | **false** |

**Direction check.** PR-1…PR-4 are the confident structural claims. PR-5 is the one I would rather be wrong
about: if it holds, the operator's tree has been pardoning real corpus defects, which is a worse finding
and a better one to have.

## Phase plan

1. **Seal** as `probe(M257x/251)`.
2. **Census** the graded population by tree-decidability, using the guard's own resolver.
3. **Repair** per `D-M257x-250-1`: an ignored target is RUNTIME — named, never silent, never a finding —
   with a fail-closed `UNDECIDABLE` disclosure.
4. **Verify** live == frozen.
5. **Test**, including a control that keeps a genuinely-missing tracked path RED.
6. **Grade** PR-1…PR-5 and close.

## Escalation conditions

- If PR-5 holds — a real broken citation surfaces — repair it in this iter if it is a one-line re-point,
  otherwise route it with a named handler and say which.
- If the population is 0 (PR-1 refuted), stop: the class does not reach this guard, and that is the
  deliverable. Do not go looking for a third guard to make the iter feel bigger.

## Acceptable close-no-lift outcomes

- PR-1 refuted with a derived zero and a proven instrument: `corpus_citation_guard` never grades an ignored
  path, the class is one-guard-wide after all, and the route closes with a number.
