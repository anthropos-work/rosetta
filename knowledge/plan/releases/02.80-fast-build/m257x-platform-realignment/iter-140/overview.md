---
iter: 140
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-08
---

# iter-140 — the receipts, and whether they reproduce

**Type:** tik
**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Step 0 — Re-survey

`adj-B`'s **P-1** (`sentinel.md:5`) is live and unrepaired at HEAD (`aaf28ab`): the page publishes a
`git grep` **with a claimed result** — *"returns **one unrelated hit**"* — and the command returns **zero**.

## Cluster / target identified — and why this class is different from iter-138's

iter-139 retracted the bare-`:NN` census because those pins have **no resolvable head**. A **published
receipt is the opposite**: it names its own command, its own pathspec and its own ref. **It is
self-contained, therefore re-runnable, therefore genuinely censusable** — the thing `TOK-08` asks for and
the thing iter-138's subject was not.

Population enumerated **before** any verdict: **22 receipts carrying a claimed count, across 15 files**
(`corpus/**`; a backticked `git grep`/`grep`/`git log -S`/`git ls-tree`/`git show` followed within 160
chars by a result verb *and* a number). Concentration: `sentinel.md` 4 · `external_services.md` 3 ·
`security_compliance.md` 2 · `studio-desk.md` 2.

## Hypothesis

**A published receipt that does not reproduce is worse than an uncited claim**, because it invites the
reader to run it and then contradicts itself in front of them. `adj-B` found one. The question is whether
it is alone.

## Expected lift

No `N` movement claimed (no reading). Deliverable: `sentinel.md:5` repaired, and the receipt class
**checked against a rule fixed in advance**, with reproduced / not-reproduced / not-checkable published as
three separate counts and a stated denominator.

## Phase plan (declared 2-step shape)

1. **Priority 1** — verify `adj-B`'s P-1 independently and repair it.
2. **Priority 2** — **run the checkable receipts.**

**Check rule, sealed before any receipt is run:**
- **Checkable** = the command's target is a repo/path this box holds (`stack-demo/*`, the rosetta tree,
  `.agentspace/rosetta-extensions`) at the ref the receipt names.
- Every checkable receipt is run **verbatim as published**, in listed order — no re-wording to make it
  pass. A receipt that needs re-wording to reproduce is **not-reproduced**.
- **Not-checkable** (repo absent, ref absent) is reported as its own count, never folded into either
  other bucket — iter-139's lesson: a bucket you cannot decide is published, not absorbed.

## Escalation conditions

3rd unplanned line → tripwire. If a receipt's failure changes a **conclusion** (not just a count), that is
a finding to repair, not a footnote.

## Acceptable close-no-lift outcomes

**All checkable receipts reproducing is a first-class result** — it would make `adj-B`'s P-1 a singleton
and close the class.
