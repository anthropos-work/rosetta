# iter-187 — decisions

## `D-M257x-187-1` — the substitution is deliberate: enumerate, don't classify

Run 17's brief named `SURVEY-M257x-iter185-other-declared-populations-unaudited` (70 collections needing
population-vs-predicate classification) as the natural next target. **Substituted** under the same
`TOK-08`, per Phase-1 Step 0. iter-186's own close records why: *"there is no syntactic marker, so the
split is judgement."* TOK-08 exists to stop paying for judgement where an enumeration exists, and one
was available — the **(section × language)** test-file matrix, one command, 11 rows.

The route is **not closed and not narrowed by this iter**; it stays open exactly as iter-186 left it.

## `D-M257x-187-2` — the exclusion is declared at (section × language), and only its completeness derived

Two shapes were available for the repair:

1. **Widen the census** to run Playwright specs. Rejected — it changes what the tool *does* under cover
   of fixing what it *says*, and the milestone has no reading of those 30 specs to compare against. It
   also silently makes the total a different total.
2. **Declare the within-section remainder and derive its completeness.** Taken. It is iter-150's split
   (*keep the partition declared, derive its completeness*), the same shape `ENV_GATED` and
   `LANGUAGE_EXCLUDED_SECTIONS` already use in this file, so there is one idiom rather than three.

`UNREAD_IN_COLLECTED` is the declared side; `unread_non_python()` is the derived side; the fence asserts
both directions (undeclared → RED, stale → RED). **Nothing about what the census runs changed.**

## `D-M257x-187-3` — iter-186's `45` is NOT retracted; a repo-wide figure is added beside it

`264 Go + 45 TS` is **true as published** — it is scoped to the six excluded sections and says so at
every site. This iter does not correct it, and the sites are left standing. What was missing is that
**no site published a repo-wide figure at all**, so the six-section one was the only number available to
quote, and a reader assuming it covered the repo would be 30 specs short. Repo-wide: **264 Go + 75 TS**.

A pointer was added at the iter-186 block in `suite_census.py` (*that figure is scoped to the six
excluded sections; it is not the repo-wide one*) rather than an edit to the number itself — per `§5`
rule 8, and because rewriting a correct measurement to mean something else is how provenance is lost.

## `D-M257x-187-4` — the presence arm is re-based on the collector, with the hazard SIZED first

Harden pass 42's `test_every_COLLECTED_section_actually_carries_PYTHON_tests` globbed
`test_*.py` + `*_test.py`; the collector `modules()` globs `test_*.py` alone. The arm's predicate was a
**superset** of the collector's, so a section whose Python tests were all spelled `*_test.py` would pass
the arm and contribute a silent zero — the exact defect the arm was written to prevent (`§5` r70/71).

**Sized before repairing** (`§5`, *measure a hazard's size or it is only a mood*): **0 `*_test.py` files
across all 11 sections.** Latent, not live. Repaired anyway because it is one line: the arm now reads
`S.modules()` itself, and a control asserts the collector still reads the named `PYTHON_TEST_GLOB`
constant — without that control the two are free to drift apart again, which is how the superset arrived.

## `D-M257x-187-5` — the escalation condition was checked and did not fire

The iter's `overview.md` pre-registered an escalation: *if the 30 specs turn out to be collected by some
other instrument this milestone quotes, this is a duplicate-population problem, not a silent-exclusion
one.* Checked: the only Python sources in `stack-core` matching `spec.ts|playwright` are
`claim_census_guard.py` (a keyword in a vocabulary list) and `test_anchor_construct_denominator.py` (a
citation fixture, `playwright.config.js:13`). Neither reads or counts the specs. **No instrument in this
milestone has ever read them.** Condition did not fire; the silent-exclusion reading stands.

## `D-M257x-187-6` — the mutation controls run against the imported module, not the tree

Six mutants were needed and the tree must not be edited mid-run (nine runs on this milestone have been
discarded as confounded for exactly that, and the restore would want a forbidden op). They are applied
to the **imported module's attributes** in a scratch driver
(`.agentspace/scratch/work-m257x/iter187_mutants.py`), which exercises the arms' real predicates — every
new arm reads `S.UNREAD_IN_COLLECTED` / `S.NON_PYTHON_TEST_GLOBS` / `S.modules` at call time.

**Limit, stated:** M5's source-inspection arm is proven by substituting a `modules` built with `exec`,
so it proves the arm reads `inspect.getsource` of whatever `S.modules` is — not that an edit to the file
on disk would be caught. That is the same thing at runtime, and it is worth saying which one was tested.

## `D-M257x-187-7` — the harden-origin route registry fired on this iter's own close block, and was dispositioned

Writing the close block turned `tests/test_harden_origin_route_visibility_m257x.py` **RED**. Its
`LEDGER_ONLY_DISPOSITIONS` registry carries harden-origin routes *no iter ever cited*; this iter's routes
list cites `FIX-M257x-h36-labeled-prover-denominator`, so the backlog fence's own population now reaches
it, and the stale-entry arm fired — correctly, and by design (harden pass 44 dispositioned the same class
of firing).

**Dispositioned by removing the entry, not by dropping the citation.** Suppressing the mention to keep a
fence green would make an open route invisible again, which is the exact defect the module exists to
prevent. The route stays **OPEN and Fate 3** — choosing between `RECALL 4/4` and `4/7` moves a figure the
milestone quotes, a design decision rather than a corollary of a test — it is simply no longer *this*
registry's business.

The module's docstring table is updated truthfully rather than quietly: the rescue was **still** an iter
happening to mention it, which is the same luck the table already calls out. What is new is that the luck
is now **observable**.
