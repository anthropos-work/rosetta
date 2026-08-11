# iter-193 — decisions

## `D-M257x-193-1` — the census's FIRST reach reading was wrong, and it was wrong the way this milestone's subjects are

The selector ran over `*.py` + `lib/*.py` per section. That reported **`dev-stack` 0 modules,
`stack-verify` 0 modules** — and both were false: their Python does not sit at those two depths. A
census whose own walk cannot reach two of the sections it claims to cover is the defect class
(`SURVEY-M257x-iter188`, the walk that prunes and does not say so), committed by the tool built to find
it. Fixed to a recursive walk with `node_modules`/`.venv` excluded **component-exact**, and the scope
is now printed with every run rather than assumed.

**Corrected reach, and it answers a question the milestone has been deferring:**

| | |
|---|---|
| non-test Python modules, repo-wide | **62** |
| sections carrying any non-test Python | **4** of 11 |
| printed add/subtract totals (the population) | **4** |
| …with a module-level registry operand (the class) | **1** |

The six Go sections carry **zero** Python modules, so for *this* class the Python-only scope is not a
gap in reach — it is the whole reachable population. That is a materially better answer than "5 of 11
sections" and it is stated as class-specific: the same defect could exist in Go and would need a
different instrument.

## `D-M257x-193-2` — the single hit is a route that has been open and invisible since harden pass 36

`labeled_spelling_pins.prove()`:

```python
print(f"  RECALL   {recall_hits}/{recall_eligible} of the text-shaped instances "
      f"(denominator excludes {len(LABELED_SET) - recall_eligible} declared-blind)")
```

The subtraction is the declared-blind count **only while every instance is readable at both refs.** An
instance whose file is absent at `repair_commit^` hits `prove`'s `continue` and increments **neither**
counter, so it silently joined the figure labelled *"declared-blind"* — and, far worse, **left the
RECALL denominator.**

**So the published recall RISES when the instrument loses the ability to read a case.** A recall that
improves on losing a case is not a recall. Measured today: 7 labeled, 1 declared blind, 6 eligible —
`7 − 6 = 1`, which **agrees**. The same agreeing-reconstruction signature as iter-192, one iter later,
in a different module.

**The census re-found `FIX-M257x-h36-labeled-prover-denominator` — open and invisible since harden pass
36** because the backlog fence could not read harden-routed items. It was found here by enumeration
rather than by memory, which is the strongest evidence available that a census beats a reading: the
route was in the brief the whole time and four iters of judgement had walked past it.

## `D-M257x-193-3` — repair: count each bucket, and say the flattering direction out loud

- `unreadable` is now **counted** in the `continue` branch instead of being inferred by subtraction.
- `declared_blind` is **derived** (`sum(1 for i in LABELED_SET if i.expect_blind)`).
- The exclusion line names **both** causes: `excludes N declared-blind and M UNREADABLE`.
- A loud follow-up fires when `unreadable > 0`, stating that recall is quoted over the shrunken
  denominator — *because that is the direction nobody audits.*
- A **partition assertion** — `declared_blind + unreadable + recall_eligible == len(LABELED_SET)` — that
  refuses to let a ratio be read when the accounting is ambiguous.

Verdict on the healthy tree is **unchanged**: `RECALL 4/6 … excludes 1 declared-blind`.

## `D-M257x-193-4` — the census caught its own author, one minute after the repair landed

The repair's own follow-up line was written as
`{len(LABELED_SET) - declared_blind} text-shaped instances` — **a fresh instance of the very class the
iter is closing**, in the function being repaired. `printed_arithmetic_totals` returned it on the next
run and it was replaced with a counted `text_shaped`.

This is the iter's best single piece of evidence, and it is worth more than the repair: **judgement had
just spent twenty minutes on this exact defect in this exact function and reproduced it anyway.** An
enumeration that keeps running is not a nicer way to find what a careful reader would find — it catches
what a careful reader, freshly primed, actively re-introduces.

The registry's completeness fence then went RED for `printed_arithmetic_totals` itself within a minute
of the function existing (adjudicated `DECLINE:verdict` — it returns findings). Two fences firing on
their own author inside one iter.

## `D-M257x-193-5` — the fence, and the two arms that stop the zero being decoration

Shipped in `derivation_registry` (which already owns *derivations vs declarations*) rather than as a new
guard module, and fenced from `test_frozen_expectation_census_m257x.py`:

- `test_no_printed_total_is_assembled_from_a_module_level_registry` — the zero, with the scope in the
  failure message so a future reader gets the denominator without re-deriving it.
- `test_the_census_STATES_its_scope_and_the_scope_is_not_empty` — a floor on modules and sections, so a
  **shrunken** walk cannot present as a clean tree. `§9`, and the exact failure `D-M257x-193-1` made.
- `test_the_census_FIRES_on_the_shape_it_exists_to_find` — a synthetic module carrying the defect.
- `test_the_census_does_NOT_fire_on_arithmetic_over_LOCALS` — the other direction; adding two local
  counts is ordinary reporting.
- `test_TEST_modules_are_out_of_scope_and_that_is_deliberate` — a test asserting `len(A) + len(B)`
  states an identity on purpose, which is the opposite defect; the exclusion is named, not silent
  (iter-186's rule).

Plus three arms on `prove()` itself, including a **partition control that fires on a real configuration**
— an instance both declared-blind *and* unreadable is counted twice, which is what happens the day a
declared-blind instance's file moves. The first draft of that control monkeypatched `sum`; it was
replaced, because a control that needs a builtin patched is testing the patch.
