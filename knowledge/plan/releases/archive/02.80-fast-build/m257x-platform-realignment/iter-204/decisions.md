# iter-204 — decisions

## `D-M257x-204-1` — a vocabulary that cannot be derived gets its REACH measured instead

`_MEASURED_NOUNS` is a hand-maintained tuple inside the module written to end hand-maintained tuples,
and it is the lens for **two** censuses since iter-203. It cannot be derived: the superset is *every
number followed by a word*, and only a human can say that `4 says` is not a count.

So it gets what `§5` prescribes for a declared population — **reach, measured against the superset, with
the misses named**:

| | before | after |
|---|---|---|
| number+word occurrences scanned | 478 | 483 |
| matched by the vocabulary | **106 (22.2 %)** | **183 (37.9 %)** |
| distinct uncovered **plural-shaped** words | **57** | — |
| addressable residual (plural-shaped, not a named verb) | 57 | **0** |

`noun_vocabulary_reach()` returns the audit; a fence asserts the residual is empty, that the scan is
non-empty (a zero residual over an empty scan is not a measurement), and — the mutation control — that
gutting the vocabulary to one noun **reopens** the residual, so the audit is provably keyed on the thing
it claims to audit.

## `D-M257x-204-2` — 37 nouns taken, and the verbs are WRITTEN DOWN rather than filtered

Plural *shape* is not nouniness: `says`, `closes`, `misses`, `enumerates`, `ships` all follow a number in
this repo and none is a count. The split is judgement, so it is a named set — `_NOT_NOUNS` — and the
reach audit reports it as excluded-by-name. **A correct exclusion is still a defect while it is silent**;
a filter nobody can count is exactly that.

37 measurement nouns were taken from the corpus's own uncovered list — `blockers` (13 occurrences),
`columns` (6), `names` (5), `citations`, `records`, `predicates`, `anchors`, `repairs`, `orgs`,
`basenames`, `literals`, `routes`, `dispositions`, `playthroughs`, `subtests`, `schemas`, `bytes` and the
rest.

## `D-M257x-204-3` — iter-199's zero SURVIVED the widening, and that is the result worth keeping

The printed-measurement-literal census read **zero** `literal` findings under the 29-noun vocabulary. At
37 nouns more it still reads **zero** — `{guarded-zero: 8, ordinal: 2}`, no findings.

That distinction was previously unavailable and it is the one a declared vocabulary always leaves open:
*a zero under a narrow lens* and *a zero under a wide one* are different claims. The printed sites are
genuinely clean. Pinned by `test_the_printed_census_stayed_at_ZERO_through_the_widening` so it cannot
quietly stop being true.

**The docstring class was the opposite.** Same widening, same tree, and the population moves **94 → 162**
— a 72 % undercount one iter old. Two censuses over one vocabulary, and only one of them was
vocabulary-limited; nothing before this iter could have told you which.

## `D-M257x-204-4` — the ceiling came from the census, not from the dry run — 162, not 160

The simulation predicted **160**. The census returns **162**, and the two extra members are
`noun_vocabulary_reach`'s **own docstring**, which states the reach measurement and thereby joins the
population it measures.

Taking the ceiling from the simulation would have repeated `D-M257x-203-2` — *a census's denominator must
come from the census* — one iter after recording it. The ratchet is re-baselined 95 → **162** with the
reason stated at the constant: a ceiling counts what the instrument can see, so a wider lens moves it
without a single new literal being written. **A ratchet over a declared vocabulary bounds only what the
vocabulary admits** — which is why the vocabulary now has a fence of its own.
