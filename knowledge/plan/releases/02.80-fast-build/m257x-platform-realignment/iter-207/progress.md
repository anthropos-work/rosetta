**Type:** tik — under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them),
*census the mechanical classes; stop sampling them.*

# iter-207 — the three literal censuses drop every test module, and the dropped population is the bigger one

## What was measured

All three literal censuses in `stack-core/derivation_registry.py` share one population line:

```python
if any(part in _CENSUS_SKIP for part in path.parts) or path.name.startswith("test_"):
```

The exclusion is stated **twice as a rationale** (*"a test stating `4 of 7` is stating an expectation on
purpose, which is the opposite defect"*) and **never once as a size**. Sized this iter, on the working
tree at close, with the shipped instrument:

| population | rows | `standing` |
|---|---|---|
| **censused** — non-test, docstrings ∪ comments ∪ other literals | **322** | **164** |
| **silently excluded** — `test_*.py`, same grammar, same classifier | **462** | **314** |

**The excluded population is the larger one on both units** — 1.43× the rows, 1.91× the `standing`
figures — across **79 modules in five sections** (`demo-stack`, `dev-stack`, `stack-core`,
`stack-injection`, `stack-verify`), so it is not a `stack-core` fact either.

The rationale is right about **assertions** and wrong about **prose**, and a test module carries both.
`self.assertEqual(len(rows), 7)` is an expectation on purpose and goes RED when wrong. A test module's
docstrings, comments and message strings are asserted by nothing, are read as evidence, and rot exactly
like the production prose these censuses exist to catch. `§5`: **a membership check cannot see a hole
inside a member** — the exclusion is at FILE grain, the distinction it wants is at CONSTRUCT grain, so
it drops both halves to be rid of one. And `§5`: **a CORRECT exclusion is still a defect while it is
silent.**

## The pre-registered subset — 3 of 6, and the split is structural rather than lucky

Six derivable rows were named in [`overview.md`](overview.md) and sealed in this iter's **first** commit
(`91ee86f`) before any of them was derived.

| # | site | figure | verdict |
|---|---|---|---|
| 1 | `test_suite_census_population.py:90` | `all 11 sections` | ✅ live — `all_sections()` = 11 |
| 2 | `test_suite_census_population.py:432` | `424 tests` / 75 files | ✅ live — `TS_FIRST_ENUMERATION`; derived ts files 45 + 30 = 75 |
| 3 | `test_suite_census_population.py:586` | `2,204 + 510 subtests` | ✅ live — `GO_FIRST_READING` |
| 4 | `test_fence_registry_completeness_m257x.py:2` | `25 modules declared; 23 enumerated` | ❌ **stale on BOTH operands — live 27 declared / 27 enumerated** |
| 5 | `test_test_collection_fence.py:286` | `2837 tests collected, 1 error` | ❌ historical (harden pass 26), graded as a claim about today |
| 6 | `test_story_org_count_guard.py:252` | `119 of the 164 files` | ❌ historical (iter-191), graded as a claim about today |

The 3/3 split falls exactly along one line, which is worth more than the ratio: **rows 1–3 name a value
some constant or function in the repo still computes; rows 4–6 name a measurement taken at a past iter
and repaired since.** The subset was selection-biased toward derivability (stated in `overview.md`
before the reading), so **3 of 6 is not an estimate for the other 308** — the same disclosure iter-206
made for its own 4 of 6.

Row 4 is the one that was wrong in both senses at once: it is historical *and* its figures no longer
describe the registry — `RP.discover_fences()` returns 6 participating + 21 standalone = **27**, so a
reader taking that sentence as the registry's size would be wrong by 2 on the declaration count and by
4 on the enumeration count, in the module whose entire subject is that registry's completeness.

## Why rows 4–6 read `standing`, and what that does to the 314

`_classify_measurement` grades `dated` on a `_DATED_MARKERS` hit within **±120 characters** of the
match. A test module names its provenance in a first-line module or class header — *"M257x iter-157 —
the fence registry's membership must follow the DECLARATION"* — and the figure it dates sits further
than 120 characters below it. Every such row is graded as a claim about **today**.

`dating_window_sensitivity()` derives the size of that exposure on every run, in **both** populations:

| population | `standing` rows | of which the enclosing unit carries a dating marker **outside** the window |
|---|---|---|
| test modules | **314** | **93** (29.6 %) |
| non-test modules | **164** | **119** (72.6 %) |

**This is a sensitivity reading, not a reclassification.** A unit carrying a marker somewhere does not
make every figure in it historical — rows 1–3 above are exactly that case, three live figures inside
docstrings that name an iter. What the reading establishes is the **shape** (`§5`: *print the SIZE,
assert the SHAPE*): the `standing` bucket is **not robust to the window width in either population**, so
its size is an **upper bound** on historical-free standing figures rather than a count of them. The
non-test share being the higher of the two is the part that matters — the class this milestone has been
working for six iters was never a count of live claims.

## CORRECTION — the standing class is 157 at `5f4b779`, not 147

Harden pass 50 corrected the milestone's standing-class figure **168 → 147** and recorded it as *"the
standing class is 147, not 168."* Re-derived this iter, per the standing rule that a figure from iters
197–206 is re-derived and never carried:

> **`git archive HEAD` of `rosetta-extensions` at `5f4b779`, unpacked to a scratch tree, censused by
> that tree's OWN `derivation_registry.py`** — i.e. the instrument that ships in the same commit:
>
> docstrings **178 rows / 80 standing** · comments **136 / 77** · **union 314 distinct rows / 157
> standing.** The ceilings in that tree read 179 and 137, consistent with the row counts and not with
> pass 50's 164 and 117.

Pass 50's own method statement explains the gap and is not itself wrong: it re-derived *"with iter-206's
own vocabulary, and the window fixed — so the only variable is the defect."* Holding the vocabulary
fixed was the right way to isolate the window defect. What travelled up to the milestone was the
**narrow-lens** number without that qualifier, published as the class's size. The shipped instrument
reads **157**.

**That is the third milestone-level value for one class: 168 → 147 → 157.** Recorded as a marked
correction appended to the milestone `progress.md`, never substituted — the same treatment pass 50 gave
iters 205/206.

## What shipped

`stack-core/derivation_registry.py`:

- **`_measurement_units(src, tree)`** — one extraction of the two window shapes (whole string literal
  outside a `print(...)`; whole contiguous `#` comment block), shared by everything below so the fourth
  census and the sensitivity reading cannot become the pair of agreeing reconstructions this module has
  already paid for twice.
- **`_census_rows(root, *, tests)`** — the census, parameterised on the one predicate that separated the
  two populations. `_unit_line` resolves a match offset to a line for either window shape.
- **`excluded_test_module_literals(root)`** — the fourth site-kind, sibling row grammar, sibling
  classifier, plus `TEST_MODULE_LITERAL_CEILING` on the same ratchet contract.
- **`dating_window_sensitivity(root)`** — derived on every call; **no docstring carries the numbers**
  (harden pass 50's rule). The unit is a **distinct row**, matching the censuses' `sorted(set(...))`;
  counting occurrences instead read **162 against 157** for one population and would have looked like a
  disagreement about the corpus rather than about the unit (`§5` r75).
- `_MEASURED_NOUNS` **+1**: `class(?:es)?`, forced by the residual arm exactly as `matches` was at
  harden pass 48 — **and the reason it had never appeared is the subject of this iter.** The word is
  written in test modules, which no census could read. **A vocabulary derived from a population with a
  hole in it inherits the hole**, so closing the fourth site-kind widened the other three by
  construction: `DOCSTRING_LITERAL_CEILING` **179 → 181**, `COMMENT_LITERAL_CEILING` **137 → 141**.
- `DECISIONS` **+1** — the sixth consecutive RED this table has raised on its own author within a minute
  of the function existing.

Six arms in `stack-core/tests/test_frozen_expectation_census_m257x.py`
(`TheFourthSiteIsTheOneTheOtherThreeDROP`), none of which carries a figure:

1. the shared extraction reproduces **both** siblings with an **empty symmetric difference** on the
   non-test population — measured **319 = 319, 0 either way** at the time it was written. Not an
   agreeing total, which is what an agreeing reconstruction always gives.
2. the two populations **partition** the tree — no row on both sides, every row's filename on the side
   its predicate says.
3. the **direction** is the assertion and the sizes are printed: excluded > censused on rows and on
   `standing`, and the population spans more than one section.
4. **anti-vacuity against the subject** (`§5` iter-94): a staged tree carrying the identical sentence in
   a production module and a test module; the siblings must see exactly one and the fourth census the
   other. An arm that only counted rows stays green if the predicate is dropped on both sides.
5. window sensitivity, asserted as a shape in both populations, plus a cross-check that the sensitivity
   reading and the two siblings agree on the non-test size — i.e. that they count the same unit.
6. the fourth ratchet, with the siblings' anti-prediction slack.

And `TheCeilingProseDoesNotContradictTheCeiling`'s **own name list was a registry with two entries while
the module had three ceilings** (`§5` iter-184). It now derives the list from the module, so a fourth
ratchet enrols itself.

Three historical figures were repaired by **making the sentence self-dating** rather than by re-pinning
the number — `§5` iter-111, *a verdict with a GRAMMAR states its provenance INSIDE the payload*, and
iter-206's *de-literalise beats re-pin*. All three now grade `dated`, mechanically verified.

## Close — 2026-08-09

**Outcome:** the population three censuses silently drop is **462 rows / 314 standing** against the
**322 / 164** they keep — larger on both units, across five sections — and it is now enumerated by a
fence that keeps running. The pre-registered subset came back **3 of 6**, split structurally rather than
by luck. Two corrections landed: the fence-registry sentence was stale on both operands (25/23 against a
live 27/27), and the milestone's own standing-class figure is **157 at `5f4b779`**, not pass 50's 147 —
a third value for one class.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (thirty-ninth consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — **counted: iter-207 is the first tik of run 22, 1 of 5** — (6) protocol-stop: n —
(7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-207-1` … `D-M257x-207-4` (see [`decisions.md`](decisions.md))

**Audit:** `stack-core`, `/usr/bin/python3 -m pytest` (**3.9.6**), **Python** — **91 passed** in
`test_frozen_expectation_census_m257x.py` (the changed fence, +6 arms) and **314 passed · 1 skipped**
across the nine touched-and-dependent modules
(`test_frozen_expectation_census_m257x` + `test_fence_registry_completeness_m257x` +
`test_test_collection_fence` + `test_story_org_count_guard` + `test_claim_census_guard` +
`test_guard_family` + `test_suite_census_collection` + `test_suite_census_population` +
`test_repair_postcondition`).
**RED-proof battery, mtime-mitigated (`§5` r77):** two mutations, both fired and both restores
sha-verified against `303a3fc1…` — inverting `_census_rows`'s population predicate took **5 of the 6
new arms RED and left exactly the one that does not read it green**; blinding `dating_window_sensitivity`
took **the 1 arm that does** RED. Four further REDs were observed unstaged, on this iter's own author:
the registry-completeness arm, the noun-vocabulary residual arm, and both sibling ratchets.
*Scope, stated rather than implied (`§5` r60): `stack-core` only, Python only, changed-code reach. No
whole-section run this iter — the tree was edited throughout, and nine runs on this milestone have been
discarded as confounded for exactly that. No Go, no TypeScript, and the other ten rext sections were
not run.*

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-iter207-the-standing-bucket-is-an-upper-bound-not-a-count` — **NEW.** 119 of the 164
  non-test `standing` rows (72.6 %) sit in a unit whose dating marker falls outside the ±120-char
  window. Whether the window should widen, or the classifier should read the enclosing *construct*, is
  a design question the sensitivity reading sizes but does not answer. **Do not widen it blind** — that
  moves rows between two reported buckets and would look like the class shrinking.
- `SURVEY-M257x-iter207-the-excluded-test-population-has-314-unverified-standing-figures` — **NEW.**
  The fourth census enumerates them; none is verified. Same per-module shape as its non-test sibling
  route, at roughly twice the size.
- `SURVEY-M257x-iter206-the-other-161-standing-figures-are-in-modules-with-no-fence` — **RE-STATED at
  the corrected denominator: 164 non-test standing rows live, 7 derived.** The route's numerator was
  itself one of the milestone's stale figures.
- `SURVEY-M257x-iter206-a-figure-can-be-stale-before-its-own-iter-closes` — **still open, and this iter
  is a fourth instance**: the fourth ratchet's first draft (301) came from this iter's own pre-flight
  probe, which read docstrings and comments only and so reproduced, in the sizing, the exact narrowing
  the census exists to fix. The census read 462. The wrong first draft is kept as the arrow's left
  operand on purpose.
- Unchanged and still open: `SURVEY-M257x-iter203-the-standing-class-is-not-mechanically-decidable` ·
  `SURVEY-M257x-iter202-published-citation-figures-predate-the-truncation-fix` ·
  `SURVEY-M257x-iter202-anchor-subject-census-extension-vocabulary-is-narrower-than-the-census` ·
  `SURVEY-M257x-iter202-the-eighteen-false-RED-pairs-remain-substrate-dependent` ·
  `SURVEY-M257x-iter201-published-suite-totals-predate-the-runner-gap-closing` ·
  `SURVEY-M257x-h45-printed-measurement-literals-uncensused` ·
  `SURVEY-M257x-iter200-battery-stagers-are-safe-by-isolation-not-by-discipline` ·
  `SURVEY-M257x-iter200-only-one-test-module-ever-clears-a-memo` ·
  `SURVEY-M257x-iter197-the-derivation-registry-sees-only-NAME-shaped-derivations` ·
  `SURVEY-M257x-iter196-no-typescript-test-has-ever-been-EXECUTED` · and the standing queue.

**Lessons:**
- **A census's exclusion is part of its claim, and a file-grain exclusion cannot make a
  construct-grain distinction.** The rationale here was correct about assertions and simply had no
  vehicle for saying so, and the price was dropping a population 1.4× the size of the one kept.
- **A vocabulary derived from a population with a hole in it inherits the hole.** `classes` is an
  obvious measurement noun; it was missing for six iters because the sentences that use it live where
  no census could look. Closing a site-kind therefore widens the others by construction, and the
  ratchets must be re-taken in the same iter.
- **Isolating a variable and publishing the result are two different acts.** Pass 50 was right to hold
  the vocabulary fixed to isolate the window defect. The number that then travelled up to the milestone
  carried none of that qualification and became the class's size.
