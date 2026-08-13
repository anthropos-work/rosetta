# iter-208 — decisions

## `D-M257x-208-1` — the derivation is a READING of the collector's glob, not a shorter route to the registry

`python_sections()` could have been written three ways:

1. `set(all_sections()) - set(LANGUAGE_EXCLUDED_SECTIONS)` — what `derive_sections()` already does. A
   subtraction over a declaration; it can never disagree with the declaration, so it cannot grade it.
2. `SECTIONS` — the same thing with an extra hop.
3. **`any(d.rglob(PYTHON_TEST_GLOB))`** — the collector's own glob. Chosen.

Only (3) can go RED. The three agree on this tree, which is exactly the condition under which the wrong
one gets shipped, and this module has been burned by it three times (`all_sections` at iter-192, the two
compose readers at iter-190, the shared classifier at harden pass 48). **The independence arm exists so
the choice is provable rather than stated**: it stages a language-excluded section carrying a Python
test file, where (1) and (3) must disagree.

## `D-M257x-208-2` — the ENV_GATED nodeid finding is ROUTED, not landed

Running the four sections surfaced that all nine `ENV_GATED` keys are `file.py::test_name` while every
one of those tests is a method on a `unittest.TestCase`. `pytest --deselect` with such a key matches
nothing **and reports nothing** — this iter hit it directly, deselecting three tests that then ran and
failed anyway.

**It is not a live defect in the census**, and saying so is part of the finding: the keys are consumed
only by `run_one` (whose `failures` list is built by `RE_FAIL_PYTEST`, whose `(?:\S+::)?` deliberately
drops the class) and by `stale_declarations` (which splits once and greps `def {test}(`). Internally
consistent. What is wrong is that a **private grammar is spelled exactly like a public one**, so any
consumer who reads a key as a pytest nodeid gets a silent no-op — `§5`, *an abbreviated id is not an id*.

Routed under the scope-creep tripwire: it is a **third distinct line** in an iter with two planned ones,
and the correct repair (declare the grammar, or store real nodeids and adapt both consumers) is a
design choice, not a one-liner. `SURVEY-M257x-iter208-env-gated-keys-are-not-nodeids`.

## `D-M257x-208-3` — the wrong clause is NOT edited in place, because this skill may not write that file

Five instances of the wrong disclosure live in `hardening-ledger.md`. `/developer-kit:build-mstone-iters`
states the ownership rule without an exception:

> This skill does NOT write to `hardening-ledger.md` — that file is owned exclusively by
> `/developer-kit:harden-mstone-iters`. Even if a tik surfaces a documentation insight worth recording
> in the ledger, leave it for the next harden invocation to capture.

So the correction is recorded where this skill *does* own the surface — the iter's own record, the
milestone `progress.md`, and `§5` of the protocol doc — and the ledger repair is routed to the next
harden pass by name. **Editing another skill's file to fix a documentation defect would have been the
same class of act as re-pinning a stale sha to make a signal go quiet.**
