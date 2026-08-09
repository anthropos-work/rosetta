**Type:** tik — under `TOK-08`.

# iter-199 — the printed-measurement-literal class, enumerated and taken to zero

## The census

`derivation_registry.printed_measurement_literals(root)` — the sibling `printed_arithmetic_totals`
(iter-193) never got. It walks every **non-test** Python module in `rosetta-extensions`, finds each
`print(...)`, and matches a digit group followed by a **measurement noun** inside its *literal* string
segments. Interpolated values are deliberately not collected: `f"{n:,} passed"` is the **repaired** form,
and a census that flagged it would be telling authors to undo the fix.

**It classifies rather than filters**, because the naive predicate is dominated by identifiers, and an
exclusion nobody can size is the narrowing iter-158 caught grading 14 of 14 broken checks green:

| kind | this tree | why it is not a finding |
|---|---|---|
| **`literal`** | **2** | **the defect** |
| `ordinal` | 2 | `TIER 1  pairs` — the digit names a tier, not a quantity |
| `guarded-zero` | 7 | `CANNOT RUN — 0 file(s) in scope` — the branch that reaches the print proves the value |
| **total** | **11** | scope: **62 non-test modules across 4 sections** (`census_scope`) |

**`guarded-zero` is written as a CONDITION, not a blanket.** The `0` must appear in a message carrying a
refusal marker (`CANNOT RUN` / `COULD NOT RUN` / `UNREADABLE` / `UNMEASURED` / …). A bare `OK — 0 files
needed repair` stays a **finding**, and a paired control proves the two sort differently. Measured: 7 of
7 zero-hits on this tree satisfy the condition — which is a reading, not the assumption it replaced.

I classified all 11 by hand before trusting the classifier. It agreed on 11 of 11.

## The two findings, and the repair

Both in `labeled_spelling_pins.prove()`, printed on **every** run of the labeled-set prover:

```
    (a) IN THE HAYSTACK …
        4 of 7 instances. … flagged the exact repaired assertion in 4 of 4.
    (b) IN THE VALUE …  3 of 7 instances:
          · iter-155 scope fence            …
```

**Two lines below `print(f"  LABELED SET     {len(LABELED_SET)} confirmed instances")`** — the set size
derived, the taxonomy split hand-written. And **all three figures were correct.** That is the class's
whole signature and the reason four prior readings walked past them: a literal that is right is
indistinguishable from a derivation until the population moves.

Repaired **structurally** (`§5` r71 — *derive the expectation from the same source the code derives
from*), not by re-typing a fresher number: `Instance` gained a **`surface`** field (`haystack` | `value`)
recording *where the pinned spelling lives*, the three `value` instances were marked, and the printed
block now counts from `LABELED_SET`. The three bullets naming the value-surface instances are derived
too — `iter_found`, `label` and `note` off the same objects, so a fourth instance enrols itself.

**Verification that matters more than the green:** the derived output reproduces the hand-written figures
**exactly** — `4 of 7`, `4 of 4`, `3 of 7`. The repair changed no fact. It changed who computes it.

`surface` is a **separate axis** from the existing `expect_blind`, and that is asserted rather than
assumed: they agree on **5 of 7**, so `surface` is a measurement and not a rename. (`expect_blind` is
about *this instrument's* visibility; `surface` is about where the pinned literal lives.)

## Post-repair reading

**`literal` = 0**, `ordinal` = 2, `guarded-zero` = 7. The class is closed *by an enumeration that keeps
running*, which is the only sense in which this milestone counts a class closed.

## The registry's completeness fence fired on its own author, again

`derivation_registry.DECISIONS` went RED for `printed_measurement_literals` within a minute of the
function existing — the second time in the same file (iter-193's sibling was the first). Classified at
source as `DECLINE:verdict`. Worth recording because it is the fence working on the person adding to it,
and because it independently corroborates
`SURVEY-M257x-iter197-the-derivation-registry-sees-only-NAME-shaped-derivations`: this function returns
`list[str]` and **was** seen; iter-197's five returned `dict`/`int` and were **not**.

## Close — 2026-08-09

**Outcome:** the class that had been hand-repaired in three consecutive work units — harden pass 45
(three sites), harden pass 47 (`2,714` under the wrong unit), iter-197 (*"the other 121 modules"*, stale
inside its own iter) — now has a census. **11 hits over 62 non-test modules in 4 sections: 2 `literal`,
2 `ordinal`, 7 `guarded-zero`**, with the two excluded classes reported rather than dropped and the
zero-exclusion written as a checkable **condition** (a bare `OK — 0 files` still counts). Both `literal`
findings sat in `labeled_spelling_pins`, two lines below a derived `len(LABELED_SET)`, and **all three
figures were correct** — which is why four readings missed them. Repaired by deriving them from a new
`Instance.surface` axis; the derived output reproduces `4 of 7` / `4 of 4` / `3 of 7` exactly.
Post-repair: **`literal` = 0**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (thirty-first consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — **counted:** iters 197, 198, 199 = **three** tiks this run — (6) protocol-stop: n —
(7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-199-1` … `D-M257x-199-4` (see [`decisions.md`](decisions.md))

**Audit:** `stack-core`, `/usr/bin/python3 -m pytest` (3.9.6) — **52 passed** in
`test_frozen_expectation_census_m257x.py`; the module gained **10 arms** this iter, **42 → 52**, counted
from `git show HEAD:` against the working file. **75 passed** across it plus
`test_spelling_pin_census_m257x.py`, and **33 passed** in
`test_derived_count_guard.py` + `test_fence_registry_population_m257x.py`. The changed module is green
under **both** runners (unittest 3.9.6: `Ran 52 tests … OK`). *Scope: `stack-core` only, Python only,
changed-code reach (`§5` r60) — no Go, no TypeScript, and the whole-section figure is not re-taken.*

**Side-deliverables:** none.

**Routes carried forward:**
- `SURVEY-M257x-h45-printed-measurement-literals-uncensused` — **CLOSED.** Censused corpus-wide over
  rext's non-test Python, both findings repaired structurally, class at zero, controls fire in both
  directions, and the two exclusions are sized rather than silent.
- `SURVEY-M257x-iter199-the-literal-census-reads-PRINTS-only` — **NEW.** Scope is `print(...)` calls in
  non-test modules. **Comments, docstrings and module constants are not read**, and pass 45's own
  findings included a docstring table. The census's zero is a statement about printed output; the same
  literal one line up in a comment is invisible to it. Stated rather than implied (`§5` r60).
- `SURVEY-M257x-iter199-the-noun-list-is-a-declared-vocabulary` — **NEW.** `_MEASURED_NOUNS` is a
  hand-listed set of ~30 words. A count printed beside a noun outside it (`3 heroes`, `5 orgs`) is not
  seen. This is `§5` r70/71's *a fence pinned to a SPELLING is not pinned to a PROPERTY*, applying to
  this iter's own instrument, and it bounds the zero above.
- Unchanged and still open: `SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` · `FIX-M257x-h44-claim-census-guard-is-single-runner` ·
  `SURVEY-M257x-iter197-the-derivation-registry-sees-only-NAME-shaped-derivations` (independently
  corroborated this iter) · `SURVEY-M257x-iter198-the-nineteen-exposed-pairs-are-unadjudicated` ·
  `SURVEY-M257x-iter198-materialization-reads-the-working-tree-by-construction` ·
  `SURVEY-M257x-iter196-no-typescript-test-has-ever-been-EXECUTED` · and the standing queue.

## The iter wrote two of these itself, in this file

The first draft of the scope line above read *"51 non-test modules across 8 sections"* and the audit line
read *"41 → 52, +11 arms"*. **Both were typed from impression, neither was measured, and both were
wrong** — `census_scope` says **62 modules across 4 sections**, and `git show HEAD:` against the working
file says **42 → 52, +10**. Caught by running the two commands before committing, in the iter whose
entire subject is published numbers nobody derives.

Recorded rather than quietly fixed, because it is the strongest available evidence for the class's
reach: **prose is not covered by this iter's census.** `printed_measurement_literals` reads `print(...)`
calls in Python modules. A knowledge-base markdown file full of counts is outside it, and two defects of
exactly the fenced shape landed there while the fence was being written.
`SURVEY-M257x-iter199-the-literal-census-reads-PRINTS-only` is where that residual lives, and this is its
first measured instance.

**Lessons:**
- **A number that is CORRECT is the hardest kind of rot to see.** All three repaired figures matched
  reality on the day they were read; nothing in a reading distinguishes a right literal from a
  derivation. Only the enumeration does.
- **The repeat-repair count is the signal.** Three hand-fixes of one shape in three consecutive work
  units said more about the class than any of the three individual defects did.
- **An exclusion should be a condition, not a category.** "A printed `0` is fine" would have been a
  blanket; "a printed `0` inside a message that says the tool refused" is checkable, and it leaves the
  bare case a finding.
