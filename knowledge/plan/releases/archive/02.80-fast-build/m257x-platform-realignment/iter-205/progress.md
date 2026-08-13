**Type:** tik

# iter-205 — the third site-kind: `#` comments

A `#` comment is a **token**, discarded before an AST exists. `printed_measurement_literals` walks
`print(...)` calls; `docstring_measurement_literals` walks string constants; **neither could ever see a
comment** — and the comment is where this repo writes its *"Measured: N of M"* provenance, the sentences
most likely to be read as evidence.

`comment_measurement_literals` reads them with `tokenize`: **118 measurement-shaped numbers**, of which
**95 `standing`** — a far higher standing share than either sibling (docstrings 73 of 164; prints 0 of
10), which is exactly what comments are for (`D-M257x-205-1`). The classification rule is now
**one function with two callers**, fenced as such: writing it twice was the available shortcut and it is
the shape iter-202 paid 16-against-19 for. The three site-kinds **partition**, proven by staging one
sentence three ways and requiring each census to see only its own.

## The vocabulary was case-sensitive — and that blinded all three at once

iter-204's residual arm returned exactly one entry: `playthroughs::1`. But `playthroughs?` **is** in the
vocabulary. The occurrence is *"23 Playthroughs stayed green"* — and `_MEASURED_RE` was
case-**sensitive** while the noun list is written lower-case, so **every capitalised measurement noun was
invisible to all three censuses simultaneously**, in a repo that capitalises Playthroughs, Stories and
Heroes (`D-M257x-205-2`).

**Nobody would have found this by re-reading the word list — the word was in it.** It took an arm that
compares the vocabulary against what the tree actually writes, which is the difference between a reach
audit and a longer list.

## The reach audit had the blind spot it exists to find

`noun_vocabulary_reach` scanned string literals only, for exactly one iter — the identical hole it was
built to expose in its consumers (`D-M257x-205-3`). Extended to comments: superset **483 → 799**, reach
**38.4 %**, residual **zero**.

## Close — 2026-08-09

**Outcome:** the last unreachable site-kind is censused, ratcheted and partitioned from its two siblings
by a shared classifier rather than a second copy of one. Along the way the vocabulary turned out to be
**case-sensitive**, which had been hiding capitalised measurement nouns from all three censuses at once —
found by a fence, not a reading — and the reach audit turned out to be carrying the exact blind spot it
exists to find. Both ceilings were taken from the census after the last change, which is the third
consecutive iter in which taking the number from anywhere else would have set it through a narrower lens.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (thirty-seventh consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
the one RED was again this iter's own new derivation demanding classification, resolved at source
(`D-M257x-205-5`) — (5) cap-reached: **n** — **counted, not felt, and CORRECTED before commit**: iters
202, 203, 204, 205 = **four** tiks this run against a cap of **five**. The first draft of this line
graded it `y`/`exit-5`, which is the mis-grade run 19 made (a fourth tik read as the fifth) and the one
this run was warned about by name. Four is not five — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**
**Decisions:** `D-M257x-205-1` … `D-M257x-205-5` (see [`decisions.md`](decisions.md))

**Audit:** `stack-core`, `/usr/bin/python3 -m pytest` (3.9.6) — **76 passed** in
`test_frozen_expectation_census_m257x.py` (the changed fence) and **151 passed** across
`test_claim_census_substrate_m257x.py` + `test_claim_census_guard.py` + `test_test_collection_fence.py`
+ `test_suite_census_collection.py` + `test_guard_family.py` + `test_claim_census_skip_registry_m257x.py`.
Green under **both** runners (unittest 3.9.6: `Ran 76 … OK`).
*Scope: `stack-core` only, Python only, changed-code reach (`§5` r60) — no Go, no TypeScript, and the
other ten rext sections were not run.*

**Side-deliverables:** none this iter.

**Disclosed defect in this iter's rext commit message** (`f9a0db1`): a backtick-quoted token was
consumed by the shell before `git commit` saw it, so the sentence about the residual arm reads
*"returned  for a word that IS in the vocabulary"* with the value missing. The commit is already on
origin and amending it would need a force-push, which is forbidden here — so it is disclosed rather
than rewritten. The value is `playthroughs::1`, and the substance is in `D-M257x-205-2`.

**Routes carried forward:**
- `SURVEY-M257x-iter204-the-superset-scan-is-string-literals-only` — **CLOSED.** Comments are censused,
  and the reach audit reads them too.
- `SURVEY-M257x-iter205-the-standing-buckets-total-168-and-one-is-derived` — **NEW, and it SUPERSEDES
  the iter-203 and iter-204 forms of the same route.** Across the three site-kinds the `standing` class
  is **73 (docstrings) + 95 (comments) = 168**, and exactly **one** has been derived (`basename_index`,
  iter-203). The class is now fully *sized* and almost entirely *unverified*; the route's figure has been
  restated three times in three iters by instrument changes alone, which is itself the finding.
- `SURVEY-M257x-iter205-comment-provenance-notes-are-the-highest-risk-standing-figures` — **NEW.** 95 of
  118 comment figures are `standing`, against 73 of 164 in docstrings. Comments hold the provenance
  notes — *"Measured: N of M"* — so the site-kind with the least derivation carries the sentences most
  likely to be quoted as evidence. A known-stale instance is in hand and unrepaired:
  `claim_census_guard.py:1275` says *"949 pairs"* where the census now prints **1,015** (moved by
  iter-202). Repairing it means deriving from `substrate_exposure`'s own `under_clones`, which is a code
  change to a comment's meaning and was not attempted at the cap.
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
- **An instrument inherits the blind spot of whatever it reads.** The reach audit was string-literals-only
  while its whole job was finding that hole in its consumers.
- **A residual arm finds what re-reading a list cannot.** The case-sensitivity defect was invisible to
  inspection because the missing word was *present* in the vocabulary; only comparing the vocabulary to
  the tree exposed it.
- **Take the ceiling after the last change, from the instrument.** Three iters running, quoting an
  earlier scan would have fixed a ratchet against a tree that no longer existed.
- **Repeated REDs on a registry's own author are a result.** Five in four iters is the strongest evidence
  available that the completeness check has not fallen behind.
