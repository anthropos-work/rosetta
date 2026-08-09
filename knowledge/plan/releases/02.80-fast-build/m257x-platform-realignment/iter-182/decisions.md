# iter-182 — decisions

## `D-M257x-182-1` — a two-member route is sized BEFORE it is planned, and these two were not one job

`FIX-M257x-iter170-two-modules-cannot-run-on-the-modern-interpreter` names two modules as if they were a
pair. Measured at HEAD:

| module | lines | tests | pytest surface |
|---|---|---|---|
| `test_gen_override_home_binds.py` | 165 | **16** | one `@fixture`, one `@parametrize`, one `monkeypatch` |
| `test_claim_census_guard.py` | 481 | **25** | `@fixture` **plus** `tmp_path`, `monkeypatch`, `capsys` |

The first is a mechanical conversion. The second requires re-implementing three pytest builtins, each
with its own failure modes.

**Decision: convert the first completely; register the second inside a fence.** Not a half-landing —
Fate 1 for one member and Fate 3 for the other, with the reason a measurement rather than an estimate.
A route that names N things has not thereby established that they are one piece of work.

## `D-M257x-182-2` — the conversion's real risk was SILENT NARROWING, and the count is part of the contract

The obvious translation of `@pytest.mark.parametrize` with 5 cases is one `subTest` loop. Every
assertion still executes and every failure is still reported — **and the module collects 12 tests instead
of 16.** A count that moves during a change whose entire contract is *nothing moves* is the signature
this milestone spends its iters on.

**Decision: five named methods, and prove the count on both interpreters.** Measured:

| runner | before | after |
|---|---|---|
| pytest 8.4.2 · `/usr/bin/python3` 3.9.6 | 16 collected | **16 passed** |
| unittest · `/opt/homebrew/bin/python3` 3.14.6 | **could not run at all** | **Ran 16 … OK** |

Every `assert` statement, message and docstring is unchanged; the fixture became a plain builder called
on the first line of each test that took it; `monkeypatch.setattr` became an explicit save/restore in
`try/finally`. iter-158's rule — *a narrowing that grades a broken check green is a defect, not a fix* —
applied to a conversion rather than to a fence.

## `D-M257x-182-3` — the predicate is the DECORATOR, not the import, and the difference is the whole population

`pytest.skip()` is a call inside a test body: a module using only that still collects and runs under
plain `unittest`. `@pytest.fixture` and `@pytest.mark.*` change how the module is **collected**, which
`unittest` cannot do at all.

**Decision: key the census on the decorators.** Keyed on `import pytest` instead, the population would
include every module that merely calls `pytest.skip` and would read as a fleet-wide crisis rather than as
**one** open member. A control asserts that more modules import pytest than are blocked by it, so the
distinction stays load-bearing rather than becoming a comment.

## `D-M257x-182-4` — the class is closed by an enumeration, and the open member lives INSIDE it

Population after the conversion, derived repo-wide: **1** — `test_claim_census_guard.py`.

**Decision: fence it as a classified population with a ratchet, in the module whose subject it already
is.** `test_test_collection_fence.py` exists for *"a test file may not hide tests from the runner that
reads it"* and already carries two arms — statement ORDER and failure to IMPORT. **Needing a runner is
the third way**, and it belongs beside them: no new module, no new registry tax (iters 178–181, fourth
consecutive iter). The remaining member is an `UNITTEST_BLOCKED` entry naming its cost and its owner, so
the obligation is enumerated by something that keeps running rather than by a markdown bullet.

All three new arms were **RED-proven live** against the real population — an unclassified intruder, the
ratchet at 2, and a `@fixture` creeping back into the converted module.
