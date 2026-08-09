**Type:** tik

# iter-203 — the measurement literals a `print()`-only census cannot see

iter-199 shipped a census of hand-written measurement counts and walked **`print(...)` calls only** —
recorded against itself as `SURVEY-M257x-iter199-the-literal-census-reads-PRINTS-only`. A number in a
**docstring** rots identically and is read more often: the docstring is where a guard states what it
measured.

`docstring_measurement_literals` enumerates that site-kind across rext non-test modules — **95
measurement-shaped numbers across 62 sites**, split **53 dated · 7 document-relative · 35 standing**
(`D-M257x-203-1`). It **asserts on none of them**, and the reason is the one this milestone has now paid
for twice in two iters: `dated` and `doc-relative` are mechanically decidable, *"this present-tense
sentence asserts a current property"* is not, and failing on it would grade prose. What is fenced is the
**size** — a ratchet the population may fall below and may not rise above without a reason — plus
non-emptiness of both mechanical classes and a **mutation control** that strips the dated marker from one
sentence and requires the classification to move.

## The census's own denominator was wrong, and that is the class

This iter's `overview.md` and the first draft of the census's comment both said **85 across 55**. The
census returns **95 across 62**. The 85 came from a scratch probe written minutes earlier that did not
apply the ordinal rule. **Quoting a throwaway measurement beside a fence is precisely what the fence
enumerates** — committed while describing it (`D-M257x-203-2`). Corrected at both sites with the
provenance stated.

## The first member: a figure iter-202 invalidated one iter ago

`basename_index`'s docstring carried *"205 of 695 line-pinned citations"* as a standing property of the
corpus. Derived now:

| grain | bare | line-pinned |
|---|---|---|
| pair (claiming unit × citation) | **410** | **979** |
| **distinct citation text** — what the sentence means | **292** | **704** |

Both operands had moved, and the numerator moved mostly **because of iter-202**: un-truncating
`.json`/`.jsx`/`.tsx`/`.graphqls` recovered 35 invisible citations and re-spelled 63, and those are
overwhelmingly bare filenames. **The missing unit is half the defect** — *"205 of 695"* is checkable
against neither grain because it names neither.

Repaired to state both grains, and fenced by `TheBasenameShareIsDERIVED`, which recomputes both operands
from the live census and asserts the docstring carries them — with an anti-vacuity arm requiring the two
grains to actually **differ** (coincide and the unit label is untested) and one requiring the words
*distinct* and *pair* to be present at all (`D-M257x-203-3`).

## Close — 2026-08-09

**Outcome:** the `print()`-only route is closed at the site-kind it named. The class outside `print()` is
**95 across 62 sites**, classified three ways, ratcheted, and reported rather than graded — because only
one of the three classes is mechanically decidable and grading the other would be the over-reach this
milestone priced twice in iter-202. Its first member was a figure **iter-202 itself invalidated**:
`205 of 695` → **292 of 704** distinct, **410 of 979** at pair grain, now derived by a fence instead of
believed. And the census's own denominator was quoted from a throwaway probe — the class caught in the
act of being described.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (thirty-fifth consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
the one RED (`test_every_executable_derivation_is_classified`) was **this iter's own new derivation**
demanding classification, resolved at source in the same minute (`D-M257x-203-4`) — (5) cap-reached: n —
**counted:** iters 202, 203 = **two** tiks this run — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**
**Decisions:** `D-M257x-203-1` … `D-M257x-203-4` (see [`decisions.md`](decisions.md))

**Audit:** `stack-core`, `/usr/bin/python3 -m pytest` (3.9.6) — **95 passed** across
`test_frozen_expectation_census_m257x.py` + `test_claim_census_substrate_m257x.py` (the two changed
fences), and **121 passed** across `test_claim_census_guard.py` + `test_test_collection_fence.py` +
`test_suite_census_collection.py` + `test_guard_family.py` + `test_claim_census_skip_registry_m257x.py`.
Both changed modules green under **both** runners (unittest 3.9.6: `Ran 95 … OK`).
`claim_census_guard --check` green (**1,130** unevidenced, baseline 1,164); `guard_family` drives every
guard end-to-end and reports its own anti-vacuity verdict.
*Scope: `stack-core` only, Python only, changed-code reach (`§5` r60) — no Go, no TypeScript, and the
other ten rext sections were not run. Harden pass 47's **1,699** remains the last whole-section
`stack-core` figure and predates iter-201, iter-202 and this iter.*

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-iter199-the-literal-census-reads-PRINTS-only` — **CLOSED.** The site-kind it named is
  enumerated, classified, ratcheted and mutation-controlled.
- `SURVEY-M257x-iter203-thirty-five-standing-figures-are-sized-but-unverified` — **NEW.** The `standing`
  bucket is **35** and exactly **one** of them has been derived. The other 34 are *reported*, not
  *checked* — the census sizes the class and says nothing about how many are stale. Deriving them is
  per-site work (each needs the module's own population recomputed), and a sampled estimate would be the
  sampling `TOK-08` replaced.
- `SURVEY-M257x-iter203-the-standing-class-is-not-mechanically-decidable` — **NEW.** `standing` is the
  residual after two decidable classes, so its miss rate is unmeasured in **both** directions: a dated
  sentence whose marker sits >120 chars from the number reads as standing, and a standing sentence that
  happens to mention a version reads as dated. Declared as a limit rather than fixed.
- Unchanged and still open: `SURVEY-M257x-iter202-published-citation-figures-predate-the-truncation-fix`
  (this iter closed **one** of its instances, inside the instrument) ·
  `SURVEY-M257x-iter202-anchor-subject-census-extension-vocabulary-is-narrower-than-the-census` ·
  `SURVEY-M257x-iter202-the-eighteen-false-RED-pairs-remain-substrate-dependent` ·
  `SURVEY-M257x-iter201-published-suite-totals-predate-the-runner-gap-closing` ·
  `SURVEY-M257x-h45-printed-measurement-literals-uncensused` ·
  `SURVEY-M257x-iter200-battery-stagers-are-safe-by-isolation-not-by-discipline` ·
  `SURVEY-M257x-iter200-only-one-test-module-ever-clears-a-memo` ·
  `SURVEY-M257x-iter199-the-noun-list-is-a-declared-vocabulary` ·
  `SURVEY-M257x-iter197-the-derivation-registry-sees-only-NAME-shaped-derivations` ·
  `SURVEY-M257x-iter196-no-typescript-test-has-ever-been-EXECUTED` · and the standing queue.

**Lessons:**
- **A census's denominator must come from the census.** The probe that sized the population and the
  function that enumerates it disagreed by 10, and the probe's number was already written into two
  documents before the function ran once.
- **Report the class you cannot decide; ratchet its size.** Three classes, one undecidable, and the
  honest instrument is the one that says so and still bounds the total — rather than either grading prose
  or leaving the class unmeasured.
- **A repair in one iter is a staleness event in the next.** iter-202 changed what the citation
  population *is*; every figure describing that population became a candidate the same moment, including
  one inside the instrument's own docstring.
