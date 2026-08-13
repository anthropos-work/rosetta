---
iter: 145
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-08
---

# iter-145 — grade the 21, and answer the scope call with evidence

## Step 0 — re-survey before targeting

`TOK-08`'s `Next-tik direction` is *"work the mechanical classes in descending measured size."* The
re-survey does not read the corpus-citation classes again — iters 138–144 worked them and iter-143's
census left the residual (`orphan-no-referent`, 333 of 380) explicitly **not mechanically decidable at
fence quality**. What the re-survey DOES find is a mechanical class of size **21 that has never been
graded at all**, sitting in this milestone's own tooling:

> `hardening-ledger.md:2827` — *"'the whole suite' in this ledger has always meant `stack-core` alone —
> one section of five, 1,280 of 3,040 tests. 21 failures sit in sections no harden pass or iter close
> has ever executed, and they were invisible for the whole milestone."*

⚠️ **The quotation is accurate; the claim inside it is not** (M257x iter-173). The quote is left verbatim
because altering a block-quote to fix its source is how a record stops being a record. **`1,280 of 3,040`
is `1,281 of 3,062`** — the ledger's denominator was `2,978 passed + 11 skipped`, dropping that same
table's **22 failures**, so it had switched unit from *executed* to *passed-and-skipped*. The correction
is made at the claim's own sites (`corpus/ops/platform-alignment.md`, the milestone `progress.md`), and
the two `hardening-ledger.md` sites are routed to the next harden pass, which owns that file.

That count has been **routed three times** (passes 30, 31, 32) with a one-line characterisation —
*"provably not ours … live-clone / live-container assertions … pre-existing, environment-coupled"* —
that was derived from `git diff --name-only` **scope**, never from reading a single failure. It is
therefore exactly the shape iter-144 closed on, one iter ago:

> `D-M257x-144-2` — **grade a survey arm's findings before treating its count as a backlog; a routed
> count is an estimate of work, not a measurement of defects.** iter-144's own route said *"8 true
> sites"* and graded out at **10 findings, 7 true**.

## Active strategy reference

`TOK-08` — *census the mechanical classes; stop sampling them.* A test suite is the most mechanical
census the milestone has: every instance is enumerated by construction, and each one is decidable by
running it. This iter censuses the one class `TOK-08` never reached because nobody ran it.

## Cluster / target identified

`FIX-M257x-h30-nonstackcore-suite` — the four rext sections (`demo-stack`, `dev-stack`,
`stack-injection`, `stack-verify`) that **no iter close and no harden pass in 144 iters has ever
run**, and the 21 failures in them.

## Hypothesis

The 21 are **not one phenomenon**. The routed characterisation assigns all 21 to a single cause
(environment coupling) on the strength of a scope argument that cannot distinguish causes. Grading
them individually will split the population — and any member that is NOT environment-coupled is a
live defect in the tooling this milestone exists to re-align, invisible for the milestone's whole life.

## Expected lift

No `N` reading. The deliverable is the **graded census**: 21 (or whatever the count is at this ref)
enumerated, each with a cause, denominator stated. Any member that grades as a real defect and is
repairable in scope lands here (Fate 1); the rest route with named handlers.

## Phase plan

- **A** — re-run all five sections at the current ref; record the failure set + its denominator. **No
  tree edits while a suite runs** (three prior runs were discarded as confounded for exactly that).
- **B** — grade EVERY failure individually: real-defect · environment-coupled · stale-expectation.
- **C** — repair what is correctly repairable in scope; route the rest.
- **D** — state the scope-call position as a decision (the orchestrator's standing instruction: an iter
  that needs a position states the assumption and records it), close.

## Escalation conditions

- A graded **real defect** in a write path the gate's clause 4 covers → surface, do not silently repair.
- The suites cannot be run at all on this host → route forward with the host finding, close-no-lift.

## Acceptable close-no-lift outcomes

If all 21 grade as environment-coupled with per-failure evidence, that is a **complete** iter: the
routed characterisation would then be confirmed by measurement instead of asserted from scope, and
the scope call gets its answer (widening the suite would have caught nothing). Falsification is the
deliverable either way.
