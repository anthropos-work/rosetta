**Type:** tik

# iter-204 — the vocabulary the literal censuses see through

Two censuses now report through one hand-written list of nouns, and everything either of them says is
conditioned on it. `_MEASURED_NOUNS` is a hand-maintained tuple **inside the module written to end
hand-maintained tuples** — `SURVEY-M257x-iter199-the-noun-list-is-a-declared-vocabulary`, open since
iter-199 and load-bearing for a second consumer since iter-203.

It cannot be derived, so it gets reach instead (`D-M257x-204-1`): scan the superset — *any* number
followed by a word — and report what the vocabulary matches.

| | before | after |
|---|---|---|
| number+word occurrences | 478 | 483 |
| matched by the vocabulary | **106 (22.2 %)** | **183 (37.9 %)** |
| distinct uncovered plural-shaped words | **57** | — |
| addressable residual | 57 | **0** |

37 measurement nouns taken from the corpus's own uncovered list; the rest are verbs — `says`, `closes`,
`misses`, `ships` — and they are **written down** in `_NOT_NOUNS` rather than filtered, because a
correct exclusion is still a defect while it is silent (`D-M257x-204-2`).

## The two censuses answered the widening in opposite directions

- **Printed census: still ZERO.** iter-199's zero was taken under a 29-noun vocabulary; 37 nouns later
  it is still `{guarded-zero: 8, ordinal: 2}` and no findings. *A zero under a narrow lens and a zero
  under a wide one are different claims*, and nothing before this iter could tell them apart. Pinned
  (`D-M257x-204-3`).
- **Docstring census: 94 → 162.** A **72 % undercount**, one iter old. Same vocabulary, same tree,
  opposite verdict — which is exactly why a shared lens has to be measured rather than trusted.

## Close — 2026-08-09

**Outcome:** the declared-vocabulary route is closed with numbers on both sides. Reach goes
**22.2 % → 37.9 %** and the addressable residual to **zero**, fenced with a scan-size arm and a mutation
control that guts the vocabulary and requires the residual to reopen. The widening then separated two
claims that had been indistinguishable: the **printed** census's zero is real and survives it, while the
**docstring** class was undersized by 72 %. The re-baselined ceiling was taken from the census rather
than the dry run — **162, not the 160 predicted**, the two extra being the reach audit's own docstring
joining the population it measures.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (thirty-sixth consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
the one RED was again this iter's own new derivation demanding classification, resolved at source —
(5) cap-reached: n — **counted:** iters 202, 203, 204 = **three** tiks this run — (6) protocol-stop: n —
(7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-204-1` … `D-M257x-204-4` (see [`decisions.md`](decisions.md))

**Audit:** `stack-core`, `/usr/bin/python3 -m pytest` (3.9.6) — **100 passed** across
`test_frozen_expectation_census_m257x.py` + `test_claim_census_substrate_m257x.py`, and **126 passed**
across `test_claim_census_guard.py` + `test_test_collection_fence.py` + `test_suite_census_collection.py`
+ `test_guard_family.py` + `test_claim_census_skip_registry_m257x.py` + `test_retired_service_endpoints.py`.
The changed fence module is green under **both** runners (unittest 3.9.6: `Ran 70 … OK`).
*Scope: `stack-core` only, Python only, changed-code reach (`§5` r60) — no Go, no TypeScript, and the
other ten rext sections were not run.*

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-iter199-the-noun-list-is-a-declared-vocabulary` — **CLOSED.** Reach measured against the
  superset with its denominator stated, residual at zero, exclusions named and sized, mutation-controlled.
- `SURVEY-M257x-iter203-thirty-five-standing-figures-are-sized-but-unverified` — **SUPERSEDED and
  RE-SIZED to 71.** The number in its own title is wrong as of this iter: the `standing` bucket is
  **71**, not 35, because the vocabulary it was counted through was 37 nouns narrower. Still exactly
  **one** of them derived. *A route's own figure is a measurement and inherits every weakness of the
  instrument that took it.*
- `SURVEY-M257x-iter204-the-superset-scan-is-string-literals-only` — **NEW.** `noun_vocabulary_reach`
  walks string constants; a measurement written in a `#` comment is in **no** census here, because
  comments are not in the AST and reading them needs `tokenize`. Both literal censuses inherit this
  hole, and neither has ever named it.
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
- **A zero has a lens, and the lens is part of the claim.** Two censuses shared one vocabulary; widening
  it confirmed one zero and destroyed the other's size. Neither outcome was predictable from the zero.
- **A ratchet over a declared vocabulary bounds what the vocabulary admits, not the class.** The ceiling
  moved 95 → 162 with no new literal written.
- **Write the exclusions down.** `_NOT_NOUNS` is a judgement, and the only way it stays auditable is as
  a named set the reach audit can count.
- **A route's own number ages with its instrument.** iter-203's route says *thirty-five*; one iter later
  the figure it names is 71, and the route title is the last place anyone looks for staleness.
