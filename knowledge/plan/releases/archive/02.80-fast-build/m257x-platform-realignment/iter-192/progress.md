**Type:** tik · **Protocol:** `corpus/ops/platform-alignment.md` · **Strategy:** `TOK-08`

# iter-192 — the denominator was arithmetic over a registry, and the fallback never spoke

## What was taken, and why

iter-191 routed `SURVEY-M257x-iter191-published-denominators-are-unenumerated` with a selector already
written in words: *a printed count whose derivation does not appear in the function that produced the
verdict*. Its two known members — iter-186's `suite_census` scope line and iter-191's
`story_org_count_guard` — were both found **by judgement**, one iter apart, and both already repaired.
TOK-08's claim is that a fence censuses where a reading samples, so the iter's first move was to build
the enumeration rather than look for a third instance by eye.

## The enumeration, and the two selectors that were discarded

Full table in [`decisions.md`](decisions.md) `D-M257x-192-1`. In short:

| selector | population | flagged | verdict |
|---|---|---|---|
| ≥2 distinct file-selection predicates per module | 42 modules | 16 | too coarse |
| reporting-path vs verdict-path predicate closure | 42 modules | 4 | **cannot decide** (`<dyn>` args) |
| green-path printed `len(NAME)` vs verdict-keyed names | 55 sites | **46** | **noise — 84 %** |

The intermediate census that *is* sound and worth keeping as a number: **315 printed cardinality
interpolations across 34 modules** (127 inline `len()`/`sum()`, 188 bare-name), which is the first time
this class has had a denominator at all.

**The conclusion is a refutation, and it is stated as one:** the class is **not mechanically decidable at
the grain iter-191 routed it**. Selector 2 produced the finding by *pointing*, not by deciding — two of
its four flags (`unreadable_repo_claim_guard`, `dev_flag_guard`) are sound on inspection, one of them
carrying a comment describing this very bug and its own earlier repair. A selector that flags 84 % of its
population is a sample wearing a census's clothes, and shipping it as a fence would be the defect this
milestone exists to catch, committed by the instrument built to catch it.

## The defect it led to — in the module behind this milestone's most-quoted number

`suite_census.py`. Two defects, **different strengths**, both sized before any edit (`D-M257x-192-2`).

**D1 — the printed denominator was arithmetic over a declared registry.** `SECTIONS` is derived from
disk; the total was `|SECTIONS| + |LANGUAGE_EXCLUDED_SECTIONS|`. A section retired from the tree but left
in the registry keeps counting, so the tool would print *"5 of 11"* over a tree holding ten.
**Measured: printed 11, disk 11 — they agree**, and the agreement is the finding. `§5` r70/71 in
arithmetic form.

**Its strength is lower than it first looks, and that is recorded rather than glossed:**
`test_the_PYTHON_population_is_5_of_11_and_that_stays_measurable` already compares that same sum against
disk, so a divergence **could not have passed a test run**. What it could do is print unchecked from the
**tool** — which is where the figure was read from and quoted. *Fenced by a test, unfenced in the
instrument.*

**D2 — the fallback was silent, and its own docstring asserted it was not.** `derive_sections` reads
*"falls back … **and says so** — a silent fallback is how a derivation becomes a literal again."* Both
arms returned `_SECTIONS_FALLBACK` emitting **nothing**; measured directly, `stdout=''` and `stderr=''`.
So a census against a tree it could not locate printed `scope: 5 of 11 sections` **in the words of a
measured one**. Nothing tested it — the adjacent arm asserts only that the fallback IS returned.
**D2 is the unfenced one, and it is the one that earned the iter.**

## The repair (`D-M257x-192-3`)

- `all_sections(root)` — the denominator, **read off the tree**.
- `stale_excluded_sections(root)` — declared exclusions naming no directory: reported inline and
  **excluded from the total**, so registry rot cannot pad a denominator silently.
- an **unaccounted** reporter — disk sections neither collected nor declared-excluded — so the scope line
  asserts in both directions instead of subtracting.
- `_fall_back(why)` — writes `WARNING … NOT derived … LITERAL, not a reading` to **stderr** and clears
  the new `SECTIONS_ARE_DERIVED`; `main` prints `scope: UNMEASURED` when it is false.

**Verdict unchanged — `5 of 11`, and the 11 is now read.** A repair that moves a derivation without
moving a number is the kind this class produces, and saying so is part of the claim.

## Proofs (`D-M257x-192-4`)

5 new arms, and two of them exist to stop the other three going vacuous — the mutation control asserts
the old reconstruction and the new reading **disagree** on a synthetic tree with a retired section, and
its twin pins that on the real tree they **agree** (the trap). The D2 pair covers **both** fallback arms
and the both-directions control that the warning does *not* fire on a healthy tree (`§9`).

`§5` rule 77 does not bite: every mutation builds a synthetic tree at runtime, so no `.pyc` invalidation
is involved.

## Close — 2026-08-09

**Outcome:** the routed class was **enumerated and found not to be mechanically decidable at its stated
grain** (three selectors; flag rates 16/42, 4/42-undecidable, **46/55**) — a refutation reported with its
numbers rather than papered over with a noisy fence. The enumeration nonetheless produced the lead, and
the lead landed two defects in the module behind this milestone's most-quoted figures: a **denominator
reconstructed from a declared registry** (fenced by a test, unfenced in the instrument; printed 11 = disk
11, an agreeing reconstruction) and a **fallback that fell back in silence while its docstring asserted
it spoke** (unfenced by anything). Denominator now read from disk, stale and unaccounted sections
reported inline, the fallback made to say so and the scope line made to print `UNMEASURED` when it fires.
Verdict unchanged; 5 new arms, mutation + anti-vacuity controls in both directions.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (twenty-fourth consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n (first tik of this invocation) — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**
**Decisions:** `D-M257x-192-1` … `D-M257x-192-5` (see [`decisions.md`](decisions.md))

**Audit:** `stack-core`, `/usr/bin/python3 -m pytest tests/` (3.9.6), **whole section** —
**1,662 passed · 2 skipped · 2 FAILED** in 24m18s, both failures **RED at HEAD before this iter touched
anything**. After repair, the four affected modules re-run **108 passed**, and every module that reads
the derivation registry **53 passed**. *Scope: `stack-core` only — 1 of 11 sections, Python only
(`§5` r60); the 264 Go and 45 TS tests remain unread.*

**Side-deliverables** (`D-M257x-192-5`; separate from planned scope, and they do not change the status):
- **iter-191's five proof arms were uncollectable by direct execution** —
  `test_story_org_count_guard.py` had its `__main__` guard at line 265 with
  `TheDENOMINATORIsDerivedNotRestated` appended below it, so `python3 test_story_org_count_guard.py`
  never reached them **and still printed OK**. Guard moved to end of file.
- **`claim_census_guard.dot_subsumed` unclassified since iter-188** — the derivation registry's
  completeness fence was RED across iters 189/190/191. Adjudicated `REGISTERED`.
- Shared cause: each of iters 188–191 ran a change-derived **scoped** suite that excluded the module
  *grading* it. The registry's own comments already record this about iter-163 and iter-186; this is its
  **third measured occurrence**, and a whole-section run is the only thing that has ever closed it.

**Routes carried forward:**
- `SURVEY-M257x-iter191-published-denominators-are-unenumerated` — **CLOSED BY REFUTATION.** The class
  is not mechanically decidable at the routed grain; three selectors measured, flag rates recorded. Its
  *decidable residual* is re-routed below rather than left implied.
- `SURVEY-M257x-iter192-printed-cardinality-census-is-one-section-of-eleven` — **NEW.** The
  **315 interpolations across 34 modules** figure is `stack-core` only. Same shape as
  `SURVEY-M257x-iter190-the-dual-reader-census-covers-one-section-of-eleven`; the two should be taken
  together, and the Go sections need a different instrument entirely.
- `SURVEY-M257x-iter192-agreeing-reconstructions-are-unenumerated` — **NEW.** The decidable residual of
  the closed class, and a sharper selector than the one that failed: *a printed total assembled by
  ARITHMETIC over a module-level registry, where a derivation of the same quantity exists in the module.*
  `suite_census` was one; nothing has looked for others.
- `SURVEY-M257x-iter190-one-construct-two-regexes-is-unenumerated` ·
  `SURVEY-M257x-iter190-the-dual-reader-census-covers-one-section-of-eleven` ·
  `SURVEY-M257x-iter187-the-grain-question-is-unasked-elsewhere` ·
  `SURVEY-M257x-iter186-264-go-tests-have-never-been-read` ·
  `SURVEY-M257x-iter185-other-declared-populations-unaudited` · `D-M257x-145-3` (the user's to rule) ·
  `SURVEY-M257x-iter185-reach-percentages-predate-the-citation-widening` ·
  `SURVEY-M257x-iter184-the-standing-queue-wildcard-is-unbounded` ·
  `SURVEY-M257x-iter183-segment-grammar-refuses-28-multi-id-segments` ·
  `SURVEY-M257x-iter179-thirty-battery-tests-unrun` ·
  `SURVEY-M257x-iter180-relation-grammar-supports-only-equality` · `FIX-M257x-iter173-ledger-denominator` ·
  `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` ·
  `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` (observed half) ·
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` · `FIX-M257x-h36-labeled-prover-denominator`
  — unchanged; open. Standing queue unchanged.

**Lessons:** **an agreeing reconstruction is indistinguishable from a reading** — the total that matched
disk was the one nobody checked, for sixteen iters, precisely because it matched. And the companion that
kept this iter honest: **a selector's flag rate decides whether it is a fence or a sample** — 46 of 55 is
noise, and shipping it because the class was routed with a selector already written would have been this
milestone's own defect. Both written into `platform-alignment.md` in this iter's commit.
