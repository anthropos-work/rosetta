# iter-192 — decisions

## `D-M257x-192-1` — three selectors were tried, and the flag rate is the reason two were discarded

iter-191 routed the class with a selector in words: *a printed count whose derivation does not appear in
the function that produced the verdict.* Turning that into an instrument took three attempts, and
**recording the two that failed is the point** — TOK-08's whole claim is that a census beats a sample,
and a census whose selector flags most of its population is a sample wearing a census's clothes.

| # | selector | population | flagged | verdict |
|---|---|---|---|---|
| 1 | modules with ≥2 distinct file-selection predicates (glob / walk / iterdir / suffix-in) | 42 modules | **16** | too coarse — a module may read two populations for two honest reasons |
| 2 | reporting-path vs verdict-path predicate closure, per module, over the intra-module call graph | 42 modules | **4** | **cannot decide** — a predicate passed as an ARGUMENT unparses to `<dyn>`, so `scan(corpus, ("**/*.md",))` is invisible to a per-function closure and reads as divergence |
| 3 | green-path printed `len(NAME)` vs the names the verdict is keyed on | 55 sites | **46** | **noise** — 84 % flag rate; legitimate sub-counts of a checked population are indistinguishable from the defect |

Selector 2 produced the finding by pointing, not by deciding: `unreadable_repo_claim_guard` flagged and
is **sound** — it carries a comment recording this exact bug and its own repair (*"ONE file set, used for
both the findings and the denominator. They diverged before"*). `dev_flag_guard` flagged and is also
sound: on a green run `parser - hinted` and `hinted - parser` are both empty, so `parser == hinted` and
the printed `len(parser)` is correct **by construction**, not by luck.

**The honest reading: this class is NOT mechanically decidable at the grain iter-191 stated it.** The
enumeration is still what found the defect — but it found it as a *lead*, and the deciding was judgement.
Routed forward rather than papered over; a noisy selector shipped as a fence would be the exact defect
this milestone keeps finding.

## `D-M257x-192-2` — the defect is in the module behind this milestone's most-quoted number, and it is TWO defects of different strength

`suite_census.py` — the instrument behind **`3,369 passed · 9 failed · 4 skipped`** and behind the
`scope: 5 of 11 sections` line iter-186 shipped. Both defects sized before any edit.

### D1 — the printed denominator was arithmetic over a declared registry

```python
_n_excl = len(LANGUAGE_EXCLUDED_SECTIONS)
print(f"  scope: {len(SECTIONS)} of {len(SECTIONS) + _n_excl} sections — Python only. …")
```

`SECTIONS` is derived from disk (iter-186 did that correctly). The **total** is not: it is
`|SECTIONS| + |LANGUAGE_EXCLUDED_SECTIONS|`, the size of the **declared** world. A section retired from
the tree but left in the registry keeps counting, so the tool would publish *"5 of 11"* over a tree
holding ten.

**Measured: printed 11, disk 11 — they agree.** That agreement is the finding, not a refutation of it:
`§5` r70/71, *a fence pinned to a SPELLING is not pinned to a PROPERTY*. An agreeing reconstruction is
indistinguishable from a reading until the day it stops agreeing, which is why sixteen iters of quoting
the figure never caught it.

**And the strength must be stated honestly, because it is lower than it first looks.**
`test_the_PYTHON_population_is_5_of_11_and_that_stays_measurable` already asserts that same sum against
disk, so a divergence **could not have passed a test run**. What it could do is print unchecked from the
**tool** — which is where the figure was actually read from and quoted. So D1 is *fenced by a test and
not by the instrument*, and the repair's job is to move the check into the instrument. Claiming more
than that would be the over-reading this milestone has retracted twice.

### D2 — the fallback was silent, and the docstring asserted it was not

```python
"""Falls back to the last measured tuple only when the repo root cannot be located, and says so — a
silent fallback is how a derivation becomes a literal again."""
```

Both arms — `not root.is_dir()` and an empty `found` — returned `_SECTIONS_FALLBACK` and emitted
**nothing** on stdout or stderr. Measured directly: `derive_sections(Path("/nonexistent"))` →
`('demo-stack', …)`, `stdout=''`, `stderr=''`.

The consequence is concrete: a census run against a tree it could not locate prints
`scope: 5 of 11 sections` **in exactly the words of a measured one** — a fully synthetic figure presented
as a reading, by the module whose own docstring names that failure mode.

**Nothing tested it.** The adjacent arm `test_the_derivation_is_not_a_literal_wearing_a_function`
asserts only that the fallback IS returned; it never asserted anything was said. This is iter-189's
lesson (*a stated-but-unfenced rule is a comment*) with an aggravating detail: here the prose asserted
the mitigation already existed, so a reader auditing the module would have ticked it off.

**D2 is the unfenced one, and it is the one worth the iter.**

## `D-M257x-192-3` — repair shape: read the denominator, and make the fallback speak

- `all_sections(root)` — the denominator, read off the tree (dirs minus `NON_SECTION_DIRS`). The
  language-excluded sections **belong** in it; being unread is what the scope line reports.
- `stale_excluded_sections(root)` — declared exclusions naming no directory. **Reported inline and
  excluded from the total**, so registry rot can no longer pad a denominator silently.
- An **unaccounted** reporter — sections on disk that are neither collected nor declared-excluded — so
  the scope line is an assertion in both directions rather than a subtraction.
- `_fall_back(why)` — takes the frozen tuple, writes a `WARNING … NOT derived … LITERAL, not a reading`
  line to **stderr**, and clears a new module flag `SECTIONS_ARE_DERIVED`.
- `main` reads that flag and, when it is false, prints `scope: UNMEASURED` instead of a figure. The
  fallback can no longer wear the words of a measurement.

**Verdict unchanged: `5 of 11`, and 11 is now read.** A repair that moves a number is a different kind of
claim from one that moves a derivation, and this is the second kind.

## `D-M257x-192-4` — five arms, and two of them exist to stop the other three going vacuous

- `test_the_denominator_is_the_DISK_not_the_registry` — `all_sections` is a reading.
- `test_a_RETIRED_section_left_in_the_registry_no_longer_pads_the_total` — **the mutation control**: a
  synthetic tree with one declared exclusion absent. It asserts the old reconstruction and the new
  reading **DISAGREE** there — without that `assertNotEqual` the arm would pass against the unrepaired
  code and prove nothing.
- `test_a_registry_that_matches_disk_is_NOT_evidence_the_formula_was_right` — the anti-vacuity twin,
  pinning that on the real tree the two *do* agree, which is the trap the mutation arm exists to escape.
- `test_the_fallback_SAYS_SO_on_both_arms` — D2, **both** arms (`cannot-locate` and `empty-tree`),
  asserting the emission, the word `LITERAL`, and the cleared flag.
- `test_a_DERIVED_run_leaves_the_flag_alone` — the other direction. A warning that fires on a healthy
  tree means nothing; `§9` says a census returning zero must prove its instrument, and this is the arm
  that keeps the D2 pair honest.

`§5` rule 77 (size-preserving mutations are invisible to the pytest bytecode cache) **does not bite
here**: every mutation builds a synthetic tree at runtime rather than editing a module, so no `.pyc`
invalidation is involved.

## `D-M257x-192-5` — the whole-section run found TWO REDs at HEAD that four scoped runs had walked past

Side-discoveries, not planned scope. Both were RED at HEAD **before this iter touched anything**, and
both were invisible to the change-derived scoped suites iters 188–191 ran. They are recorded here rather
than folded into the close status (`§` Phase 4 Step 0).

**Side-1 — iter-191's five proof arms were hidden from the runner its own module advertises.**
`test_story_org_count_guard.py` carried `if __name__ == "__main__": unittest.main(verbosity=2)` at
**line 265**, and iter-191 appended `TheDENOMINATORIsDerivedNotRestated` (5 tests) **below** it.
`unittest.main()` exits inside the guard, so `python3 test_story_org_count_guard.py` never reached them
**and still printed OK** — five arms proving iter-191's repair, silently uncollected, reporting success.
pytest collects them, which is exactly why nothing noticed: *the suite was green about a file direct
execution was lying about.* `§5` r75/76 (**name the runner**) in its sharpest form yet — the two runners
disagreed about the *existence* of five tests, not about their verdict. Guard moved to the end of file.

**Side-2 — `claim_census_guard.dot_subsumed` has been an unclassified derivation since iter-188.**
`derivation_registry`'s completeness fence went RED the day iter-188 added it and **stayed** RED through
iters 189, 190 and 191. Adjudicated REGISTERED (derived from `SKIP_DIRS`, a checked-in module registry,
no caller input).

**The shared cause is already written in the registry's own comments, twice, about iter-163 and about
iter-186 — and it recurred anyway.** Each of iters 188–191 ran a change-derived scoped suite that did not
include the module *grading* it. That is `§5` rule 60 stating its own price: **a scoped green is evidence
about its scope alone**, and the module that grades your change is frequently outside the scope your
change derives. Recorded not as a new rule but as its **third measured occurrence** — the rule is not
under-written, it is under-applied, and the only thing that has ever closed it is a whole-section run.

**This iter's own suite was therefore the whole section**, 24m18s, and it is what produced both findings.
The cost is real and so is the yield: the four preceding iters bought their pace with these two REDs.
