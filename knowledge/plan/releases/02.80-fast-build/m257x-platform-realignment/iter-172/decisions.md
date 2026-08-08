# iter-172 — decisions

## `D-M257x-172-1` — iter-171's reading of the 11 is REFUTED, and corrected in place

**Decision:** the 11-test gap is an **instrument artifact**, not a population difference. iter-171's route
(`SURVEY-M257x-iter171-runner-test-count-gap`) said *"11 tests execute under one runner and not the other,
and nobody has named which."* **That reading is withdrawn**, corrected where it was published (this iter's
close, iter-171's close, and the milestone ledger entry) rather than deleted — the milestone's standing
correction-vs-retraction discipline.

**The decomposition, to zero residue.** Measured per module, per runner, before any repair:

| module | unittest | pytest | diff | cause |
|---|---|---|---|---|
| `test_migrate_race_live.py` | 3 | 0 | 3 | 3 ENV-GATED failures |
| `test_ssr_origin_chain.py` | 12 | 9 | 3 | 3 ENV-GATED failures |
| `test_demopatch.py` | 62 | 60 | 2 | 2 ENV-GATED failures |
| `test_ant_academy.py` | 63 | 62 | 1 | 1 ENV-GATED failure |
| `test_interview_flag_patch_m232.py` | 12 | 11 | 1 | **1 skip** |
| `test_purge.py` | 5 | 4 | 1 | **1 skip** |
| | | | **11** | **9 failures + 2 skips, nothing unexplained** |

**Not one test executes under one runner and not the other.** Both runners ran all 1073.

## `D-M257x-172-2` — a column with two units is worse than a missing column

**Decision:** derive the `tests` figure from the same denoted set under both runners — **everything that
executed** — and take it from pytest's whole summary line rather than one term of it.

**The root cause is four characters wide and it published a wrong number twice.**

```
unittest:  RE_RAN = re.compile(r"^Ran (\d+) test")     # passes + failures + errors + SKIPS
pytest:    re.search(r"(\d+) passed", out)             # passes only
```

Both were correct parses of their own runner's output. Neither was a parse of the same *thing*. And the
column they fed was printed side by side under the banner *"runners DISAGREE about N module(s)"* — a
report whose entire purpose is cross-runner comparison.

**Which unit is the honest one is not a coin-flip.** The larger set is the population; the smaller one is a
verdict. A census of *"the RED-at-HEAD population"* must count what exists, not what passed — counting
passes would make a module get *smaller* as it broke. Repaired toward `Ran N`: `pytest_executed()` sums
`passed + failed + error(s) + skipped + xfailed + xpassed` from the **last** summary-bearing line.

**Two sub-decisions worth naming.**

- **`deselected` is deliberately not summed.** A deselected test never ran, and unittest's `Ran N` would not
  have counted it. Including it would make the columns agree by inflating one — the flattering fix `§5`
  refuses.
- **The last line, not the whole output.** pytest prints a per-test `FAILED …` block above its summary; a
  whole-output scan would double-count every failure. Fenced with a literal fixture carrying both.

## `D-M257x-172-3` — the fence's input is chosen so the control cannot be vacuous

**Decision:** the fixture is **one pass, one failure, one skip** — the minimal input on which a passed-only
counter and an executed counter *must* disagree — and the mutation control asserts that the **old** parse
still reads `1` on it.

**Why the control is written that way.** iter-170's own instrument shipped a bucket that returned zero
against nine live instances, and iter-166's finding was that a fence publishing only its FIRE side hides
whether its ACCEPT side works. A fixture of three passing tests would make the new assert green and prove
nothing — the two units coincide on any all-passing module, which is precisely why this defect survived
across 35 modules where **29 of 35 agreed**. So the control pins the *separation*: if the fixture ever
stops making the old parse undercount, the fence says so and fails, rather than quietly grading a
degenerate input.

Four assertions ship: both-runners-agree on the live fixture (3 = 3); the old parse reads 1 on the same
input (separation is real); `deselected` is excluded and mixed summaries sum correctly (unit boundaries);
and the last-line rule survives a `FAILED`-block prefix (no double-count).

## `D-M257x-172-4` — the corrected figure is re-published at the whole-population scope

**Decision:** re-run the census across **all 110 modules / 5 sections**, both runners, and publish the
corrected executed-test count — do not close on the `demo-stack` slice and route the rest.

**Because the wrong unit is already load-bearing in this milestone's record.** iter-170's headline —
*"the pytest interpreter collects 3,332 tests, 0 collection errors"* — was produced by the passed-only
parse. On the `demo-stack` slice the same parse undercounted by **11 of 1073 (1.02 %)**, so the published
whole-population figure is an undercount of unknown size, and every later sentence that quotes it inherits
that. `§5`'s rule against leaving a known-wrong published number standing applies to this milestone's own
instrument first.
