**Type:** tik (standard shape; §9 iter-type refinements consulted, none selected).

# iter-182 — the third way a test file hides its tests: it needs a runner

**Controlling strategy:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Phase A — size the route before planning it

`FIX-M257x-iter170-two-modules-cannot-run-on-the-modern-interpreter` names two modules as a pair. They
are not one job (`D-M257x-182-1`):

| module | lines | tests | pytest surface |
|---|---|---|---|
| `test_gen_override_home_binds.py` | 165 | **16** | one `@fixture`, one `@parametrize`, one `monkeypatch` |
| `test_claim_census_guard.py` | 481 | **25** | `@fixture` **plus** `tmp_path`, `monkeypatch`, `capsys` |

## Phase B — convert the tractable one, assertion for assertion

The risk was never correctness; it was **silent narrowing**. `@parametrize` with 5 cases collects as
**5 tests**; the obvious `subTest` translation collects as **1**, taking the module 16 → 12 while every
assertion still runs. So the five cases became **five named methods** (`D-M257x-182-2`):

| runner | before | after |
|---|---|---|
| pytest 8.4.2 · `/usr/bin/python3` 3.9.6 | 16 collected | **16 passed** |
| unittest · `/opt/homebrew/bin/python3` 3.14.6 | **could not run at all** | **Ran 16 … OK** |

Every `assert` statement, message and docstring is unchanged; the fixture became a plain builder called
on the first line of each test that took it; `monkeypatch.setattr` became an explicit save/restore in
`try/finally`. iter-158's rule applied to a conversion instead of a fence.

## Phase C — census the class, and key it on the right thing

Derived repo-wide after the conversion: **1** module still needs pytest to be collected.

The predicate is the **decorator**, not the import (`D-M257x-182-3`). `pytest.skip()` is a call inside a
test body and runs fine under `unittest`; `@pytest.fixture` / `@pytest.mark.*` change how the module is
*collected*, which `unittest` cannot do at all. Keyed on `import pytest`, the population would have
included every module that merely calls `skip` and read as a fleet-wide crisis instead of **one** open
member. A control keeps the distinction load-bearing.

## Phase D — where the fence went, and why not a new module

`test_test_collection_fence.py` exists for *"a test file may not hide tests from the runner that reads
it"* and already carries two arms — statement **ORDER** (a class below the `__main__` guard is never
registered) and failure to **IMPORT** (a PEP 604 annotation on a 3.9 runner aborts the whole run).
**Needing a runner is the third way**, and it belongs beside them. Fourth consecutive iter with no new
module and no new registry tax.

Shipped there: the classified population (both directions — unclassified is RED, a stale blocker is RED),
a **ratchet** floor stating 2 at iter-170 → **1** at iter-182, a regression arm on the converted module,
and the decorator-vs-import control. All three assertion arms **RED-proven live** against the real
population — an unclassified intruder, the ratchet at 2, and a `@fixture` creeping back into the
converted module.

## Runs — scope stated, and what it did NOT cover

| run | result | wall |
|---|---|---|
| the converted module, **pytest 8.4.2 / 3.9.6** | **16 passed** | 0.02 s |
| the converted module, **unittest / 3.14.6** | **Ran 16 … OK** | 0.004 s |
| `test_test_collection_fence.py` | **20 passed** (was 16) | 1.85 s |
| `stack-core` minus all 7 mutation batteries | **1,535 passed · 2 skipped · 0 failed** | 455.54 s (7:35) |

**The arithmetic is checkable:** iter-179's sweep was **1,521**; iter-180 added 5, iter-181 added 5,
iter-182 added 4, and the conversion added **0** — `1,521 + 14 = 1,535`. With the batteries' **41**,
`stack-core` is **1,576** collected.

**Not covered:** the 7 mutation batteries (none stages either module touched here) and the four other
rext sections. A scoped green is evidence about its scope alone (rule 60).

## Close — 2026-08-09

**Outcome:** one of the two modules that could not run on the modern interpreter now does — **16
collected before, 16 after, on both runners**, with every assertion, message and docstring unchanged —
and the class is closed by a **derived enumeration that keeps running** rather than by the conversion.
The remaining member is an `UNITTEST_BLOCKED` entry naming its cost and owner, under a ratchet that
states its floor, in the module whose subject *"a test file may not hide tests from the runner that reads
it"* already was.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (fourteenth consecutive `closed-fixed`; **no
`P`/`N` reading taken, so the metric is UNMEASURED, not unmoved** — `§9`, and `TOK-08` declares the
class-by-class sweep order in advance) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n —
(6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7** (BETWEEN ITERS, tree clean)
**Decisions:** `D-M257x-182-1` … `D-M257x-182-4` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none.

**Routes carried forward:**
- `FIX-M257x-iter170-two-modules-cannot-run-on-the-modern-interpreter` — **half CLOSED and the other
  half RE-HOMED.** `test_gen_override_home_binds.py` is converted; `test_claim_census_guard.py` is now
  an `UNITTEST_BLOCKED` entry inside a running fence, with its cost measured (25 tests, three pytest
  builtins) rather than estimated.
- `SURVEY-M257x-iter170-cockpit-runner-dependence` — unchanged; open. Same axis, different mechanism
  (two `test_cockpit` tests pass under one runner and fail under the other).
- `SURVEY-M257x-iter181-*`, `SURVEY-M257x-iter180-relation-grammar-supports-only-equality`,
  `SURVEY-M257x-iter179-thirty-battery-tests-unrun`, `FIX-M257x-iter173-ledger-denominator`,
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` — unchanged; open.
- The standing queue, unchanged.

**Lessons:** **a conversion that moves a COUNT is not a conversion** — the obvious `subTest` translation
would have taken the module 16 → 12 with every assertion still running, and nothing would have flagged
it. And the framing that made the iter cheap: **a route that names N things has not established that
they are one piece of work.** Sizing the two members first turned an intimidating open item into one
mechanical conversion plus one measured, registered obligation. Written into `platform-alignment.md` §8
in this iter's commit.
