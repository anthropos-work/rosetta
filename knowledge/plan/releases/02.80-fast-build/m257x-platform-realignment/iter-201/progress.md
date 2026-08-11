**Type:** tik — under `TOK-08`, closing the class iter-197 censused.

# iter-201 — 25 tests one runner could not see, now seen by both

## The reading, after

`suite_census.py --collect --runner none`, this tree, both runners:

```
124 module(s) · pytest collects 3580 · unittest collects 3580 · gap 0 · 0 module(s) DISAGREE
authoring style: testcase=124   (unittest collects `TestCase` and nothing else)
⚠ STALE declaration: stack-core/tests/test_claim_census_guard.py — no longer disagrees
```

**gap 0. Zero disagreeing modules. 124 of 124 `testcase`.** The class iter-197 enumerated to a single
member is closed, and the aggregate gap `suite_census` had been reporting since iter-170 — 25 tests, in
every whole-population baseline this milestone has quoted — is gone rather than explained.

The `⚠ STALE declaration` line is iter-197's registry grading itself in the other direction: a declared
split that no longer splits is rot, and it said so without being asked.

## The conversion, and why it is a binding change

**No case body was edited.** The 25 cases are renamed `_case_*` — so neither runner double-collects them
— and bound onto a `TestCase` **by generation from the module namespace**:

```python
CASES = sorted(n for n in dict(globals()) if n.startswith("_case_"))
for _name in CASES:
    setattr(ClaimCensusGuard, "test" + _name[len("_case"):], _bind(_name))
```

**That shape is the answer to iter-182's finding, not a stylistic choice.** iter-182 measured that the
obvious hand-translation *loses tests silently* — a hand-written class can drop one and stay green.
Generation from the namespace makes the loss unrepresentable, and it is asserted anyway, twice and
independently: the bound method names must equal `CASES` exactly (both directions, not a count), and a
second arm re-derives the number by counting `^def _case_` in the **file source**. Two derivations, one
number; if they ever disagree the binder is reading something other than this file.

Three fixture APIs were in use — `monkeypatch.setattr` ×6, `capsys.readouterr` ×6, `pytest.skip` ×2 —
measured before planning, because that surface is what decides whether this is an hour or a day. All
three are shimmed in ~40 lines (`_Monkeypatch` with undo, `_Capsys` over redirected streams,
`unittest.SkipTest`, which pytest honours too). **The module now imports pytest nowhere at all** — the
root of the asymmetry rather than a symptom of it.

**29 tests, identical under both runners** (25 cases + 4 controls): `pytest 3.9.6 — 29 passed`;
`unittest 3.9.6 — Ran 29 … OK`.

## The self-referential fence

`test_this_module_no_longer_imports_pytest` was written as `assertNotIn("import pytest", src)` and
**failed on itself** — the method's own body contains the words it searches for. Repaired to match at
line start. *A fence that cannot survive being written down is checking the file, not the import.*

## Keeping the proof after the repair

Emptying `RUNNER_COLLECTION_SPLIT` makes every live declared-vs-derived arm in iter-197's fence
**trivially true**: two empty sets are equal. That is `§9`'s standing trap — *a good repair can destroy
the proof the instrument fires* — so the proof was **moved, not lost**:

- The synthetic battery (`TheCensusCanActuallyFire`) is now the whole of the firing proof: it builds a
  bare-def module and asserts the census finds it, builds a TestCase-only tree and asserts it does not.
- A new arm asserts the empty registry **declares itself a closure** and names where the proof went — an
  empty declaration and a forgotten declaration are otherwise indistinguishable.
- The style arm gained its positive half: every module is `testcase`, which is the property that *makes*
  the zero true, rather than an equality between two empty sets.
- The offsetting-members arm is **kept**, not deleted: two modules at +7 and −7 still sum to zero, and
  the equality against the per-module set is what catches them. A zero total is only evidence when its
  parts are also zero.

## Close — 2026-08-09

**Outcome:** the milestone's last open `FIX-` route is closed, and with it the 25-test runner gap that
has been inside every whole-population figure quoted since iter-170. `test_claim_census_guard.py` — the
repo's only module written as bare fixture-taking functions, and iter-197's single censused member — now
collects **29 tests identically under both runners**, with **no case body edited**: the cases are bound
onto a `TestCase` **by generation from the namespace**, which makes iter-182's silent-loss failure
unrepresentable, and asserted anyway from two independent derivations. The census reads **124 modules ·
3,580 · 3,580 · gap 0 · 0 DISAGREE · testcase 124**, and iter-197's registry reported its own entry
STALE without being asked. The now-vacuous declared-vs-derived arms were repaired rather than left green
— the firing proof moved to the synthetic battery, the empty registry declares itself a closure, and the
offsetting-members arm was kept because a zero total is only evidence when its parts are zero.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (thirty-third consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: **y** — **counted, not felt**: iters 197, 198, 199, 200, 201 = **five** tiks, and the
`fix(M257x/199)` commit is a side-deliverable of iter-200's re-survey, not a sixth iter —
(6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **exit-5**
**Decisions:** `D-M257x-201-1` … `D-M257x-201-3` (see [`decisions.md`](decisions.md))

**Audit:** `stack-core`, `/usr/bin/python3 -m pytest` (3.9.6) — **91 passed · 1 skipped** across
`test_suite_census_collection.py` + `test_claim_census_guard.py` +
`test_claim_census_substrate_m257x.py` + `test_suite_census_population.py`; **45 passed** in
`test_guard_family.py`. Both changed modules green under **both** runners (unittest 3.9.6: `Ran 29 … OK`
and `Ran 14 … OK`). *Scope: `stack-core` only, Python only, changed-code reach (`§5` r60) — no Go, no
TypeScript; the whole-section figure is not re-taken this iter and harden pass 47's **1,699** remains the
last one, now understood to have been **25 tests short of the population** on the unittest side.*

**Side-deliverables:** none this iter. (`fix(M257x/199)`, the route-id repair, belongs to iter-200's
re-survey and is recorded there.)

**Routes carried forward:**
- `FIX-M257x-h44-claim-census-guard-is-single-runner` — **CLOSED.** Converted, both runners collect 29,
  no case body changed, lost-test controls in two independent derivations, and the class census reads
  zero.
- `SURVEY-M257x-iter201-published-suite-totals-predate-the-runner-gap-closing` — **NEW.** Every
  unittest-side total this milestone has published — including harden pass 47's **1,699** whole-section
  figure — was taken while 25 tests were invisible to that runner. The figures are not wrong for their
  runner, but no cross-runner total quoted before this iter describes the same population as one taken
  after it.
- Unchanged and still open: `SURVEY-M257x-iter200-battery-stagers-are-safe-by-isolation-not-by-discipline` ·
  `SURVEY-M257x-iter200-only-one-test-module-ever-clears-a-memo` ·
  `SURVEY-M257x-iter199-the-literal-census-reads-PRINTS-only` ·
  `SURVEY-M257x-iter199-the-noun-list-is-a-declared-vocabulary` ·
  `SURVEY-M257x-iter197-the-derivation-registry-sees-only-NAME-shaped-derivations` ·
  `SURVEY-M257x-iter198-the-nineteen-exposed-pairs-are-unadjudicated` ·
  `SURVEY-M257x-iter198-materialization-reads-the-working-tree-by-construction` ·
  `SURVEY-M257x-iter196-no-typescript-test-has-ever-been-EXECUTED` · and the standing queue.

**Lessons:**
- **Census first, then close.** iter-197 spent an iter proving the class had one member; iter-201 closed
  it in one. Neither would have been safe alone — a conversion without the census would have been a
  repair against a reconstructed size, and the census without the conversion leaves the gap in every
  total.
- **Generate the binding when the failure mode is silent loss.** A hand-written class satisfies the same
  green and cannot be audited; a generated one carries its own completeness.
- **When a repair empties a registry, say so where the registry is.** An empty declaration and a
  forgotten one look identical, and every arm that reads it goes quietly vacuous at the same moment.
