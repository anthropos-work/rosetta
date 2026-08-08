**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*).

# iter-170 — census the RED-at-HEAD population

## The result, first

**110 Python test modules, five sections, both interpreters.** Denominator stated; Go-suite sections
(264 `_test.go`) are out by declaration, not omission.

| runner | GREEN | ENV-GATED | RED | TIMEOUT | tests |
|---|---|---|---|---|---|
| `/usr/bin/python3` 3.9.6 + pytest (**the fleet runner**) | 106 | 4 | **0 actionable** | 0 | 3,332 collected |
| `python3` 3.14.6 + unittest | 101 | — | 8 | 1 | 3,279 executed |

**Zero unknown REDs at HEAD.** Every failure resolves to an already-routed class:

* **6 sha-drift** against the demo's *present* live clones — `FIX-M257x-iter145-sha-baseline-drift`. This
  is the **freshness signal**, and re-pinning it away would convert a signal into a chore. Not repaired,
  deliberately.
* **3 live-stack** — `FIX-M257x-iter145-migrate-race-needs-a-host-postgres`. No Postgres here.
* **4 runner artifacts** (see below), of which **1 is repaired in this iter**.

## The finding: "the suite" is not one command

`/usr/bin/python3` is **3.9.6** and is the **only interpreter on this box with pytest** — the fleet runner.
The working interpreter is **3.14.6** and has **no pytest at all**. So every "the suite is green" in this
milestone has been scoped by an axis nobody had named: **the interpreter**.

Measured, the two runners **disagree about four modules**:

| module | 3.14 + unittest | 3.9.6 + pytest | why |
|---|---|---|---|
| `test_claim_census_guard.py` | **cannot load** | GREEN | `import pytest` |
| `test_gen_override_home_binds.py` | **cannot load** | GREEN | `import pytest` |
| `test_progress_beacon.py` | **cannot load** | GREEN | relied on pytest putting a test's own dir on `sys.path` |
| `test_cockpit.py` | 2 RED | GREEN | server-binding tests, runner-dependent |

This is `§5` rule 60 with the scope being the interpreter. It also explains the shape of iter-169: the
rotted `Thread._stop` assertion was in the remainder of iter-167's *17 GREEN · 0 RED · **7 not-run***, and
`test_claim_census_guard` — one of the four waiver-carrying guards iter-166 shipped — **has never executed
on the modern interpreter at all.** Booked as `§5` **rule 75**.

## The instrument committed the defect it was built to prevent

The census's third bucket exists because a module needing a live stack is neither green nor a defect
(`§5` rule 73; `SURVEY-M257x-iter152`). The first draft **sniffed** that bucket from error-message
substrings and returned **ZERO against nine genuinely environment-gated failures** — the real signal was
never in the output at all. **A third bucket that never fires is a two-bucket partition wearing three
labels**, which is the exact defect rule 73 names, committed by the instrument written to honour it.

The repair is the split rule 73 already prescribes: **keep the partition DECLARED, derive its
COMPLETENESS.** The nine are named in `ENV_GATED`; a declaration whose test no longer exists is reported
**STALE**; an undeclared RED is reported **ACTIONABLE**. `test_suite_census.py` proves **each of the four
buckets fires**, that the ENV bucket does *not* swallow an undeclared failure, and that the staleness check
is not vacuous.

## What landed

* **`stack-core/suite_census.py`** — the census as a repeatable tool. `--runner both` is the only reading
  that can be called a census; it refuses the pytest half loudly if `/usr/bin/python3` is absent rather
  than reporting half a population as a whole one, and it prints the runner disagreement explicitly.
* **`stack-core/tests/test_suite_census.py`** — 7 tests: every bucket fires, the declared bucket is
  complete and non-vacuous, the enumeration covers all five sections.
* **`test_progress_beacon.py`** — one line, and it moves a module from *runs under one runner* to *runs
  under both*. Every sibling test file in that directory already carried it.
* **`§5` rule 75** in the protocol doc.

## Verification, and what it did NOT cover (`§5` rule 60)

| run | result |
|---|---|
| the census itself, 110 modules × 3.14/unittest | 101 GREEN · 8 RED · 1 TIMEOUT, 3,279 tests |
| the census's 9 non-green, re-run under the fleet runner | 131 passed / 9 failed, all 9 declared |
| `pytest --collect-only` over all five sections | **3,332 collected, 0 collection errors** |
| `test_suite_census` (new) | **7/7 OK** |
| `test_progress_beacon` under both runners | **9/9 OK · 9 passed** |
| shipped tool, smoke-run on the 34-module demo-stack slice | 30 GREEN · 4 ENV-GATED · **0 ACTIONABLE** |

**Not covered:** the TIMEOUT (`test_m257x_mechanical_fences_mutation_battery`, >600 s) is a
**self-contention** artifact of 5-way parallelism, not a verdict — the known effect `§5` rule 51 records;
it is GREEN when run alone (iter-168). The six Go-suite sections were not measured. And the census is a
reading of **this box**; the fleet's interpreter matrix is not necessarily this one.

## Close — 2026-08-08

**Outcome:** the RED-at-HEAD population is **enumerated for the first time**, with a stated denominator of
110 modules and, decisively, **a stated RUNNER** — the axis that had never been named. **Zero unknown REDs**;
all 13 non-green results resolve to routed classes or runner artifacts. The census also found that
`test_claim_census_guard` — a guard test iter-166 shipped — **has never run on the modern interpreter**, and
that the instrument's own third bucket was vacuous until repaired.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (two consecutive `closed-fixed` iters; no
no-prog streak, and no `N`/`P` reading was taken so the metric is UNMEASURED, not unmoved — `§9`) —
(3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks this run) — (6) protocol-stop: n —
(7) budget-exhausted: **y** — Outcome: **exit-7** (BETWEEN ITERS, tree clean: both repos committed, rext
pushed. The 110-module census consumed ~12 minutes of wall-clock inside a ~55-minute run and a third iter
would risk a mid-iter stop.)
**Decisions:** `D-M257x-170-1` … `D-M257x-170-3` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M257x-iter170-two-modules-cannot-run-on-the-modern-interpreter` — **NEW.**
  `test_claim_census_guard` (a waiver-carrying guard's own suite) and `test_gen_override_home_binds` use
  `@pytest.fixture` + `@pytest.mark.parametrize` and therefore cannot execute under 3.14. Converting them
  is a real change with its own failure modes, not a one-line fix — and per iter-158 the conversion must
  be shown not to weaken what they assert.
- `SURVEY-M257x-iter170-cockpit-runner-dependence` — **NEW.** Two `test_cockpit` server-binding tests pass
  under the fleet runner and fail under 3.14/unittest. Either a runner-dependent harness assumption or a
  real 3.14 behaviour difference; unresolved, and it is the only disagreement not explained by imports.
- `FIX-M257x-iter145-sha-baseline-drift` — **quantified, not repaired.** Exactly **6** tests in 3 modules.
  Still the freshness signal; still must not be re-pinned away.
- `SURVEY-M257x-iter169-rotted-assertions-beyond-Thread-_stop` — **CLOSED by measurement.** The hazard was
  censused rather than asserted: no second rotted assertion exists in the population under either runner.
- The standing queue, unchanged.

**Lessons:** **name the runner, or the suite verdict has an unstated scope.** `§8` has required *state the
environment with every number* since M255 — for timings. This iter shows the same rule owns **pass/fail**:
the same 110 modules yield 101 green under one interpreter and 106 under another, and the difference is
neither noise nor defects.

And the one that stings: **the instrument built to honour rule 73 broke rule 73.** Its third bucket was
sniffed from error text, returned zero against nine live instances, and would have published *"9 REDs"*
where the truth is *"9 declared environment-gated tests, 0 actionable."* **A bucket you did not watch fire
is a bucket you do not have** — the same sentence as iter-166's ACCEPT-side finding, one layer down.
