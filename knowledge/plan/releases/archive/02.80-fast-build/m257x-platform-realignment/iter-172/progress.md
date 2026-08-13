**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*), turned on the
milestone's **own instrument**.

# iter-172 — the census counted two different things and printed them in one column

## Phase A — decompose the 11 before interpreting them

iter-171 closed with a route that carried a number *and* a reading of it:

> the same 35 modules yield **1073** tests under unittest and **1062** under pytest — *"11 tests execute
> under one runner and not the other, and nobody has named which."*

The number is real. The reading assumes both columns denote the same set. Measured per module, per runner:

| module | unittest | pytest | diff | cause |
|---|---|---|---|---|
| `test_migrate_race_live.py` | 3 | **0** | 3 | 3 ENV-GATED failures |
| `test_ssr_origin_chain.py` | 12 | 9 | 3 | 3 ENV-GATED failures |
| `test_demopatch.py` | 62 | 60 | 2 | 2 ENV-GATED failures |
| `test_ant_academy.py` | 63 | 62 | 1 | 1 ENV-GATED failure |
| `test_interview_flag_patch_m232.py` | 12 | 11 | 1 | **1 skip** |
| `test_purge.py` | 5 | 4 | 1 | **1 skip** |
| | | | **11** | **9 failures + 2 skips — zero residue** |

**Not one test executes under one runner and not the other.** Both ran all 1073. `test_migrate_race_live`
is the tell that should have been read first: **3 → 0**, a module reported as having *no tests* by a census
whose stated job is to enumerate tests.

## Phase B — the unit, named

```
unittest:  RE_RAN = re.compile(r"^Ran (\d+) test")     # passes + failures + errors + SKIPS
pytest:    re.search(r"(\d+) passed", out)             # passes only
```

Both are correct parses of their own runner. Neither parses the same *thing*. And the column they fed is
printed under the banner **"runners DISAGREE about N module(s)"** — a report whose whole purpose is
cross-runner comparison. **A column with two units is worse than a missing column**, because it invites
exactly the population reading iter-171 published.

**Why it survived 35 modules:** the two units **coincide on any all-passing, skip-free module** — 29 of 35
here. A defect that is invisible on the healthy majority is found by the sick minority or not at all.

## Phase C — the repair

`pytest_executed()` sums `passed + failed + error(s) + skipped + xfailed + xpassed` from the **last**
summary-bearing line. Two sub-decisions, both `§5`-shaped:

- **`deselected` is not summed.** It never ran, and `Ran N` would not have counted it. Summing it would make
  the columns agree by *inflating* one — the flattering fix.
- **Last line, not whole output.** pytest prints a per-test `FAILED …` block above its summary; a
  whole-output scan double-counts every failure. Fenced against a literal fixture carrying both.

Repaired **toward the larger set**: a census of a *population* must count what exists, not what passed —
otherwise a module gets smaller as it breaks.

## Phase D — the fence, and why its input is what it is

`stack-core/tests/test_suite_census.py`, +5 tests (11 total in the module), plain `unittest`, **green on
both interpreters**.

The fixture is **one pass, one failure, one skip** — the minimal input on which a passed-only counter and an
executed counter *must* disagree. A fixture of three passing tests would make the new assert green and prove
nothing (the two units coincide there). So the mutation control pins the **separation** itself: the old
`(\d+) passed` parse must still read **1** on this input. If it ever stops undercounting, the fence fails
rather than grading a degenerate input green — iter-166's ACCEPT-side lesson, applied to a unit rather than
a waiver.

| assertion | fires when |
|---|---|
| both runners agree on the fixture (3 = 3) | the units diverge again |
| the old parse reads **1** on the same input | the fixture stops separating the units — control gone vacuous |
| `deselected` excluded; mixed summaries sum | a unit boundary moves |
| last-line rule survives a `FAILED` prefix | the parser regresses to a whole-output scan |

## Phase E — re-measure, runner named

`suite_census.py --only demo-stack --runner both`, 35 modules:

| runner | GREEN | ENV-GATED | RED | tests |
|---|---|---|---|---|
| unittest (3.14.6) | 31 | 4 | 0 | **1073** |
| pytest (3.9.6) | 31 | 4 | 0 | **1073** |

**Modules with a count difference: 6 → 0.** The gap is closed at its cause, not papered over.

### The whole-population re-measure — and what it caught

iter-170's headline — *"the pytest interpreter collects **3,332** tests, 0 collection errors"* — was
produced by the passed-only parse. Re-measured with the corrected unit across **112 modules / 5 sections**
(110 at iter-170 + the 2 this iter and iter-171 added), both runners:

| runner | GREEN | ENV-GATED | RED | TIMEOUT | modules | tests (executed) |
|---|---|---|---|---|---|---|
| unittest (3.14.6) | 104 | 4 | 4 | 0 | 112 | **3311** |
| pytest (3.9.6) | 105 | 4 | 3 | 0 | 112 | **3350** |

**Do not read `3,332 → 3,350` as the size of the unit error** — the population grew by 2 modules in the same
window, and this iter refuses that arithmetic for the same reason it refused iter-171's. What is measured
cleanly is the `demo-stack` slice at a fixed population: **+11 of 1073, 1.02 %**.

**The whole-population columns still differ, by 39 — and now the difference IS a population fact.** The
unit fix removed the artifact; what remains is the three modules the two runners genuinely disagree about
(`test_claim_census_guard` and `test_gen_override_home_binds` RED under 3.14 — the routed
`FIX-M257x-iter170-two-modules-cannot-run-on-the-modern-interpreter`; `test_battery_stage` RED under
pytest). That is the shape iter-171's route *guessed*; it is now isolated rather than assumed.

**And the run caught a regression this iter had just introduced.** `test_test_collection_fence`
(*"the main guard is the last thing in every rext test module"*) went RED under **both** runners: the new
census tests were appended with `cat >>`, landing **after** `test_suite_census.py`'s
`if __name__ == "__main__"` block. Repaired (guard moved to EOF), re-run **27 tests OK on both**. This is
`FIX-M257x-iter142-whole-suite-owed` firing in its useful direction for once — the change-derived scoped run
(`tests.test_suite_census`, green) could not see the fence that grades it; only the whole-population run
could, and it ran *before* the commit rather than after.

Two ACTIONABLE REDs are **pre-existing and unrelated**, verified by reading their failures rather than
assuming: `test_frozen_expectation_census_m257x` names `suite_census.py::modules` and `::stale_declarations`
— both shipped at iter-170, neither touched here — and `test_battery_stage`'s stdlib-shadow refusal fails
only under 3.9.6, which is itself rule-76 shaped. Both routed, neither repaired.

## Close — 2026-08-08

**Outcome:** iter-171's *"11 tests execute under one runner and not the other"* is **REFUTED and corrected
in place**. The census's `tests` column carried **two units** — `Ran N` (executed) under unittest, `N passed`
under pytest — so it under-reported pytest by every failure and every skip; `test_migrate_race_live` was
printed as having **0 tests**. All 11 decompose with **zero residue** (9 ENV-GATED failures + 2 skips), and
not one test executes under only one runner. Repaired toward the honest set, fenced with a fixture chosen so
the control cannot be vacuous, and **demo-stack now reads 1073 = 1073, count differences 6 → 0**. The
whole-population figure iter-170 published (**3,332**) is re-measured at the corrected unit (**3,350**
pytest / **3,311** unittest over 112 modules) — and the run **caught a regression this iter had just
introduced**, a main-guard placement RED that the change-derived scoped run could not see.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (fourth consecutive `closed-fixed`; no no-prog
streak, and **no `P`/`N` reading was taken, so the metric is UNMEASURED, not unmoved** — `§9`) —
(3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted:
**y** — Outcome: **exit-7** (BETWEEN ITERS, tree clean; see the close note below)
**Decisions:** `D-M257x-172-1` … `D-M257x-172-4` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none.

**Routes carried forward:**
- `SURVEY-M257x-iter171-runner-test-count-gap` — **CLOSED by refutation.** Not a population difference; an
  instrument with two units. Corrected at all three publishing sites.
- `SURVEY-M257x-iter172-two-preexisting-actionable-reds` — **NEW.**
  `test_frozen_expectation_census_m257x::test_every_executable_derivation_is_classified` (both runners;
  names `suite_census.py::modules` + `::stale_declarations`, unclassified in `derivation_registry`) and
  `test_battery_stage::test_a_stdlib_shadow_is_refused_not_staged` (**pytest/3.9.6 only** — rule-76 shaped).
  Verified pre-existing by reading the failures, not assumed. Neither repaired here.
- `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` — **NEW.** Every test-count this milestone
  published under the pytest runner before this iter is a *passed* count, therefore an undercount of the
  executed population. iter-170's `3,332` is corrected here; **any other quoted pytest count is not**, and
  nobody has enumerated where they are.
- `SURVEY-M257x-iter171-anchor-guard-detects-structure-not-staleness` — unchanged; open.
- `FIX-M257x-iter170-two-modules-cannot-run-on-the-modern-interpreter` — unchanged; open.
- The standing queue, unchanged.

**Lessons:** **name the unit, or the column has an unstated scope.** iter-170 earned *name the runner*;
iter-171 earned *an unexplained disagreement is a defect until proven otherwise*; this iter is the third
turn of the same screw and the sharpest, because the offender was **the instrument built to enforce the
first two**. A report that puts two numbers side by side under the word *"DISAGREE"* has already asserted
they are commensurable — and nothing checked that.

The general form, worth carrying: **when two measurements of the same thing differ, decompose the
difference to zero residue before interpreting any of it.** iter-171 interpreted first and published a
population claim; the decomposition took under ten minutes and refuted it entirely. And the corollary that
made it findable: **a defect invisible on the healthy majority is found by the sick minority** — 29 of 35
modules agreed, and the six that did not were the whole signal.
