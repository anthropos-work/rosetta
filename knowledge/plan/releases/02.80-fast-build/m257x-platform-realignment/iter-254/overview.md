---
iter: 254
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
route: ROUTE-M257x-249-fresh-checkout-hostile-tests
---

# iter-254 — declare the 22 preconditions, so a fresh clone stops accusing the corpus

**Active strategy reference:** `TOK-08` — census the mechanical classes; stop sampling them. iter-253 built
the census and ran it to a **named** population. `TOK-08`'s sentence is *"run it to zero, and keep it
green"*; this iter is the run-to-zero.

## Step 0 — re-survey (mandatory)

The target was measured **this session**, at rosetta `1f1e0be` / rext `d739952`, and the names are
checked in at `iter-253/evidence/iter253-box-fresh-checkout-hostile.txt`. No substitution; nothing has
moved under it.

## Cluster / target identified

The **22 BOX node-ids across 12 files**: tests that pass on this box and fail on a clean clone of both
repos, where they do not report a missing precondition but instead print sentences like *"the live corpus
must resolve every rext path"* — false about the corpus, true about the machine.

Handler: `FIX-M257x-249-declare-the-clone-precondition`. The acceptance test is the instrument iter-253
shipped: `suite_census.py --fresh-checkout` goes green when they SKIP.

## Hypothesis

Each of the 22 needs an `unittest.skipUnless` naming the untracked state it reads (`D-M257x-249-2`: skip
beats fail, because a skip states what is missing while a failure states that the corpus is wrong), and
the declaration must be attached at the **grain that actually fails** — per `D-M257x-253-5`, a class-level
decorator protects the class and not the file.

**The preconditions are READ from the failures, never guessed.** Phase A extracts the real reason for each
of the 22 on the frozen tree before a single decorator is written.

## Expected lift

The BOX population **22 → 0**, verified in both directions: frozen (each now SKIPS, with a reason naming
its precondition) and live (each still RUNS and passes — a repair that silences a test on the box where it
works is not a repair).

## Phase plan

- **A** — extract the per-test failure reason on the frozen tree; group into distinct preconditions.
- **B** — write the declarations at the failing grain.
- **C** — verify BOTH directions: the 12 files frozen (0 failed, ≥22 skipped) and live (0 skipped among
  the 22).
- **D** — close.

## Pre-registrations (sealed in this iter's first commit, before Phase A's result is read)

| # | claim | prediction |
|---|---|---|
| **PR-1** | all 22 share ONE precondition (a clone set under `stack-demo/`/`stack-dev/`) | **false** — ≥ 2 distinct preconditions |
| **PR-2** | the repair is purely additive — decorators only, no assertion logic altered | **true** |
| **PR-3** | after repair the 12 files run **0 failed** on the frozen pair | **true** |
| **PR-4** | ≥ 1 of the 22 is on inspection NOT environmental (the control mis-classified it) | **false** — the control is sound |
| **PR-5** | after repair all 22 still RUN and pass on the live tree (0 become skips here) | **true** |

`PR-4` is the one that matters: it predicts against finding a flattering extra defect, and `PR-1`
predicts against the convenient single-fix story.

## Escalation conditions

- A test whose failure has **no** declarable precondition (it is simply wrong on a clean tree) → it leaves
  this class, is named, and routes out. Do not decorate a defect.
- If the frozen verification cannot run, the iter closes with the refusal recorded rather than an
  asserted repair.

## Acceptable close-no-lift outcomes

- The reasons turn out to be heterogeneous enough that a correct declaration cannot be written for some
  subset without changing what the test asserts — that subset is characterised and routed, and the
  falsification is the deliverable.
