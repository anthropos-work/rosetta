# iter-197 — decisions

## `D-M257x-197-1` — the unit of this census is COLLECTED tests, and it is declared, not implied

`suite_census`'s existing `tests=` column counts tests that **EXECUTED** — iter-172 settled that unit
deliberately, after `Ran N` and `N passed` had been printed in one column. This iter adds a second count
in the same module with a **different** unit, which is exactly the condition that produced harden pass
47's `2,714` defect.

So the unit is stated in the section header (`COLLECTED tests, never executed ones`), in every docstring,
and in the function names (`collected_by_*`, `collection_census`). The two figures are also *shown* to
differ rather than asserted to: 3,551/3,526 collected against 3,527/3,502 executed, **+24 on each side**,
of which 13 is this iter's own new fence. A unit difference that shows up identically on both runners is
a unit difference; one that shows up on one runner is a bug.

## `D-M257x-197-2` — the census is opt-in on the CLI and unconditional in the gate

`--collect` is a flag; `tests/test_suite_census_collection.py` runs the same census on **every** run of
the Python suite with no env gate. Two consumers, two contracts — `D-M255-1`'s rule.

Deliberate, and against the obvious alternative of gating the live arm behind an env var the way pass 45
gated the Go one. Pass 45's own finding was that *cost cannot explain* leaving the stronger claim
unfenced: the Go census re-runs in 23 s, the TypeScript enumeration in 0.59 s, and **only the weaker claim
had a ratchet.** This census costs **5.9 s**. Gating it would have re-created the finding in the iter that
cites it.

`--runner none` exists for the same reason from the other side: without it, `--collect` would have cost a
whole-population execution census (~24 min) to read a 6-second fact, and a check priced that far above its
value does not get run.

## `D-M257x-197-3` — an unreadable collection is `UNREADABLE`, never `0` and never `1`

Both runners can report a plausible-looking number for a module they cannot read:

- **pytest** prints `no tests collected, 1 error` and exits **2**. Scraping the line yields **0**.
- **unittest**'s `loadTestsFromName` converts an import failure into a synthetic `_FailedTest`, so
  `countTestCases()` returns **1** — a test that does not exist.

Both are the `go_census` silent-zero: a number that sums into a clean total and reads exactly like a
healthy module. The sentinel is `UNREADABLE = -1`, pytest is read by **exit code**
(`PYTEST_RC_READABLE`), unittest by the **loader's own error list**, and
`test_an_UNIMPORTABLE_module_is_UNREADABLE_not_zero_and_not_one` pins both directions.

The pytest half is recorded as a defect *of this iter*, not as a design: the first cut shipped the scrape,
and the control written to prove the census fires is what caught it. That ordering is the point — the
control earned its place before the code did.

## `D-M257x-197-4` — a module NEITHER runner can read is a disagreement

`collection_disagreements` was one expression, `c["pytest"] != c["unittest"]`, and for a doubly-unreadable
module that is `-1 != -1` → **False**. The module left the disagreement set and would have been published
inside *"0 modules disagree"*.

The sentinel is now compared explicitly. Recorded as a decision rather than a fix because the rule
generalises past this function: **any comparison of two derived values over a shared sentinel converts
"both blind" into "both agree"** unless the sentinel is tested for. Two constructed arms pin it
(`test_an_unreadable_side_alone_…`, `test_TWO_BLIND_runners_are_not_an_AGREEMENT`) so the rule holds even
though no module on this tree currently produces the shape — a live population of one cannot prove a rule
about a shape it does not contain.

## `D-M257x-197-5` — the conversion is NOT in this iter, and `FIX-M257x-h44-…` is not re-routed

The obvious "complete" version of this iter converts `test_claim_census_guard.py`'s 25 pytest-fixture
functions to `TestCase` and takes the class to zero. Not taken, for a reason the milestone has measured:
iter-182 found that *the obvious translation loses tests silently*, and the module carries 28 `tmp_path`,
16 `monkeypatch` and 12 `capsys` usages — a careful conversion with its own lost-test control, not a
mechanical one.

Splitting it keeps both halves honest. This iter's deliverable is the **census**, which is complete: it
enumerates the class, grades declared against derived in both directions, and can fire. h44's deliverable
is the **repair**, which now has a correctly-grained size and a fence watching it. This is not a deferral
under the three-fate rule — h44 is a pre-existing open route with a named handler, unchanged by this iter
(Fate 2), and the iter's own planned scope landed in full.
