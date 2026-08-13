# iter-199 — decisions

## `D-M257x-199-1` — the census CLASSIFIES; it does not filter

Three kinds are returned and only `literal` is asserted on. The alternative — dropping `ordinal` and
`guarded-zero` inside the function — was rejected because a census that silently discards two thirds of
its matches publishes a zero whose denominator nobody can name. That is iter-158's shape (14 of 14
broken checks graded green by a narrowing) and iter-114's rule (*a reach metric names its denominator*).

So the classes are reported, and a dedicated arm asserts **both excluded classes are non-empty on this
tree**. An exclusion with no live instances is not conservative — it is unfalsifiable, and it would sit
there accumulating meaning it had never earned.

## `D-M257x-199-2` — `guarded-zero` is a CONDITION on the message, not a rule about the value

The tempting exclusion is *"a printed `0` is never a stale literal"*. True in this repo, for a reason:
the `0` appears in the CANNOT-RUN refusal idiom, where the branch that reaches the print establishes the
value. But *"in this repo"* is the load-bearing clause, and it would be invisible in the code.

Written instead as a requirement that the printed message carry a refusal marker (`CANNOT RUN`,
`COULD NOT RUN`, `UNREADABLE`, `NOT LISTED`, `UNMEASURED`, `DEAD DETECTOR`). Measured: **7 of 7** zero-hits
on this tree satisfy it. A `print("OK — 0 files needed repair")` classifies as **`literal`**, and a paired
control asserts the two sort differently — so the exclusion can be shown to have an edge rather than
being taken on trust.

## `D-M257x-199-3` — the repair is STRUCTURAL, and re-typing a fresher number was the wrong fix

`4 of 7`, `3 of 7`, `4 of 4` were all **correct**. The cheap repair is to leave them and check them
occasionally; the next-cheapest is to add a test asserting they equal the derived values. Both were
rejected — the first is what has failed three times, and the second creates a *second* literal to
maintain (the frozen-expectation class this very module censuses).

`§5` r71's prescribed repair is to derive the expectation from the same source the code derives from, so
`Instance` gained a **`surface`** field and the block counts from `LABELED_SET`. A fourth instance now
enrols itself into the right bucket, the bullet list, and the totals, with no edit.

**The check that the repair is real:** the derived output reproduces the three figures exactly. A
structural repair that changed a number would have meant the literals were also *wrong*, which is a
different (and easier) finding; this one changed only who computes them.

**And `surface` is asserted to be a genuine second axis**, not a rename of `expect_blind`: they agree on
5 of 7. Without that arm, a reviewer could not tell whether the field carried information or duplicated
an existing one.

## `D-M257x-199-4` — the iter published two unmeasured figures in its own progress.md, and they are recorded, not erased

The scope line (*"51 non-test modules across 8 sections"*; really **62 across 4**) and the arm delta
(*"41 → 52, +11"*; really **42 → 52, +10**) were typed from impression in the first draft of this iter's
own record, then caught by running `census_scope` and `git show HEAD:` before committing.

Kept in the record with the correction beside it. Two reasons:

1. **It is the class, committed by the person fencing the class, inside the hour.** Nothing argues more
   directly that the shape is a systematic pull rather than an author's carelessness.
2. **It sizes the residual the fence does NOT reach.** `printed_measurement_literals` reads `print(...)`
   in Python. Markdown prose — where this milestone publishes nearly all of its numbers — is outside its
   scope entirely, and these two are its first measured instances.
   `SURVEY-M257x-iter199-the-literal-census-reads-PRINTS-only` carries it forward; widening the census to
   prose was deliberately not attempted here (a much larger, much less mechanical population, and
   conflating the two would make the `literal = 0` unreadable).
