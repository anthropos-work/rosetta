**Type:** tik — under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

# iter-208 — the milestone's own NOT-COVERED clause names ten Python sections; the repo derives four

## The clause, and the two independent ways it is wrong

Six harden passes closed with the `§5` rule-60 scope disclosure. Counted mechanically over
`hardening-ledger.md`:

| form | instances | correct? |
|---|---|---|
| *"the ten non-`stack-core` **sections**"* | **1** | ✅ 11 sections on disk, minus `stack-core` |
| *"the ten non-`stack-core` **Python** sections"* | **5** | ❌ there are **four** |

**One adjective was inserted into a correct sentence and the number was not re-derived.** Six of the
"ten" carry **zero** `test_*.py` files — `alignment`, `clerkenstein`, `playthroughs`, `stack-secrets`,
`stack-seeding`, `stack-snapshot` — which is precisely what `LANGUAGE_EXCLUDED_SECTIONS` says about each
of them, in the module the ledger quotes for its *"5 of 11 sections"* line. The wrong form then
propagated verbatim through five consecutive passes and into the orchestration prompt of the run that
caught it.

**The second error is independent of the first and worse.** *"NOT COVERED"* was true of the pass that
wrote it and was read as *"never read"*. The same ledger holds the tables that refute it:

- **iter-145** — *"the four never-run sections were RUN"*, 21 failures graded individually.
- **harden pass 35** — all five sections in one table: `demo-stack` 9 failed · 1,038 passed · 11
  skipped; `dev-stack` 151 passed; `stack-injection` 335 passed; `stack-verify` 12 failed · 225 passed.
- a later pass — `demo-stack` 9 failed · 1,055 passed · 2 skipped; `stack-verify` 0 failed · 275 passed;
  `dev-stack` 151 passed.

`§5` iter-178 says a NOT-REACHED clause is a measurement or it is a mood. This is the failure one step
later: **a clause that WAS a measurement, copied forward until it read as a property of the milestone.**

> **Caught by Step 0, not by the close.** This iter's first plan said the four sections would get
> *"their first verdict."* Re-surveying the milestone's own ledger before running falsified that half
> and it never reached a commit. The `overview.md` records both the claim and its falsification.

## The reading — the four sections, now

`/usr/bin/python3 -m pytest` (**8.4.2** under CPython **3.9.6**), one section at a time, **rext tree
frozen from the commit before the run to the commit after** (nine runs on this milestone have been
discarded as confounded by a mid-run edit):

| section | result | wall |
|---|---|---|
| `demo-stack` | **9 failed · 1,063 passed · 2 skipped** | 3:42 |
| `dev-stack` | **151 passed** | 1:40 |
| `stack-injection` | **335 passed** | 0:07 |
| `stack-verify` | **275 passed** | 6:15 |
| **total** | **1,824 passed · 9 failed · 2 skipped** | 11:44 |

**All nine failures are exactly the nine entries of `suite_census.ENV_GATED`** — the repo's own
declaration of which tests need a live clone or a live stack — so the census's own grammar grades the
section **ENV-GATED, not RED**: `undeclared = [f for f in failures if f not in ENV_GATED]` is empty.
**Zero undeclared failures across the four sections.**

Against the last recorded reading of the same four: `demo-stack` 1,055 → **1,063**, `dev-stack` 151 →
**151**, `stack-injection` 335 → **335**, `stack-verify` 275 → **275**. Stable, +8.

## What shipped — the missing third of the language triple

`suite_census` derives `go_sections()` from `go.mod` and `ts_sections()` from
`e2e/playwright.config.ts`, and each has an arm asserting the derivation agrees with the hand-written
registry. **Python — the language this repo is mostly written in — had no derivation at all.**
`derive_sections()` looks like one and is not: it returns *everything not in
`LANGUAGE_EXCLUDED_SECTIONS`*, a subtraction over a declaration. `PYTHON_TEST_GLOB`'s own note already
states the rule that violates:

> *any predicate that claims to answer "does this section carry Python tests this census will run?" must
> derive from it rather than restate it*

**That is why the clause could rot unnoticed: the set it named had nothing to be graded against.**

- **`python_sections(root)`** — sections carrying ≥1 `test_*.py`, from the collector's own glob.
- Three arms in `tests/test_suite_census_population.py`
  (`TheLanguageTripleIsDerivedForPythonToo`), none carrying a figure:
  1. the glob derivation and the registry subtraction **agree today** — asserted so a *disagreement*
     becomes visible, not because agreement is the property;
  2. **independence** — a staged tree in which a language-excluded section carries a Python test file.
     The registry still excludes it; the glob must see it. If they agree *there*, arm 1 is measuring one
     derivation twice — the agreeing-reconstruction shape `TheDenominatorIsReadNotReconstructed` exists
     for;
  3. the two counts a disclosure can confuse — non-`stack-core` **sections** vs non-`stack-core`
     **Python** sections — must differ on this tree, because while they were equal no one could have
     caught the substitution by counting.

`corpus/ops/platform-alignment.md` `§5` gains the rule, in this iter's commit per the protocol-evolution
rule: **a NOT-COVERED clause DECAYS — it names a RUN and gets read as a PROPERTY.**

And `derivation_registry`'s completeness fence went RED on `python_sections` within a minute of the
function existing — the **seventh** consecutive time this table has caught its own author, and the
seventh piece of evidence that `unclassified()` has not fallen behind the tree. Registered
(`REGISTERED`, same class as its two siblings), with both entries kept: `derive_sections` and
`python_sections` answer different questions and arm 1 asserts they agree, which is the whole value of
having both.

## Close — 2026-08-09

**Outcome:** the milestone's standing scope disclaimer overstated the unread Python surface by **six
sections** (ten claimed, four derived) and understated its own coverage (the same document holds three
tables of those sections being run). The four sections were re-read: **1,824 passed, 9 failed, 2
skipped — and all 9 failures are the 9 declared `ENV_GATED` entries, so 0 undeclared.** The set the
clause names now has a derivation (`python_sections`), which it never had, plus an independence
mutation control proving the derivation is not the registry in disguise.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (fortieth consecutive `closed-fixed`; **no
`P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: **n, and
it is graded rather than assumed** — the 9 `demo-stack` failures are a *measurement of a suite this iter
did not touch*, pre-declared in `overview.md` as such, and they are the repo's own `ENV_GATED` set, so
the census grades the section ENV-GATED. This is not the Phase 5 § 4 test-gate RED, which is about the
iter's own changes; the iter's own changes are green — (5) cap-reached: n — **counted: iters 207, 208 =
two tiks this run against a cap of five** — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**
**Decisions:** `D-M257x-208-1` … `D-M257x-208-3` (see [`decisions.md`](decisions.md))

**Audit:** `/usr/bin/python3 -m pytest` (**8.4.2** / CPython **3.9.6**), **Python**:
**42 passed · 1 skipped** in `test_suite_census_population.py` (the changed fence, +3 arms), and the
whole-section readings above — **`demo-stack` 1,063 passed / `dev-stack` 151 / `stack-injection` 335 /
`stack-verify` 275**, tree frozen throughout.
**RED-proof battery, mtime-mitigated (`§5` r77):** `python_sections`'s glob predicate was replaced by the
registry subtraction; **the independence arm went RED and the agreement arm stayed green** — which is
the finding restated as a control, since agreement is exactly what cannot detect the substitution.
Restore sha-verified against `b652ad17…`.
*Scope, stated rather than implied (`§5` r60) — and stated with a **derived** denominator this time:
the **four** non-`stack-core` **Python** sections were run; `stack-core` was **not** re-run this iter
(iter-207's nine-module reading stands); the **six** sections carrying no Python test file are Go or
mixed-toolchain and were not run in any language; **no Go and no TypeScript verdict is claimed.***

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-iter208-env-gated-keys-are-not-nodeids` — **NEW.** All nine `ENV_GATED` keys are
  `file.py::test_name`, but every one of those tests is a method on a `unittest.TestCase`, so the real
  nodeid carries a class segment. The keys are **internally consistent** — `RE_FAIL_PYTEST`'s
  `(?:\S+::)?` deliberately drops the class, and `stale_declarations` greps `def {test}(` — so this is
  **not a live defect in the census**. It is a private grammar spelled like a public one: `pytest
  --deselect` with these keys matches nothing **and says nothing**, which this iter hit directly. `§5`,
  *an abbreviated id is not an id*. Routed rather than landed under the scope-creep tripwire — it is a
  third line and the census is not wrong today.
- `SURVEY-M257x-iter208-the-wrong-clause-is-still-in-the-hardening-ledger` — **NEW.** The five wrong
  instances live in `hardening-ledger.md`, which **`/developer-kit:build-mstone-iters` is forbidden to
  write** (*"that file is owned exclusively by `harden-mstone-iters`"*). The correction is recorded here
  and in the milestone `progress.md`; **the next harden pass must fix its own five clauses and stop
  copying the sentence forward.**
- `SURVEY-M257x-iter208-a-language-triples-third-leg-was-missing-for-twelve-iters` — **NEW.**
  `go_sections` landed at iter-195, `ts_sections` at iter-196; the Python leg arrived at iter-208. When
  a family of derivations is built one language at a time, the language everybody assumes is covered is
  the one that gets skipped.
- Unchanged and still open: all of iter-207's routes, plus the standing queue.

**Lessons:**
- **A scope disclosure is a claim about a population.** It gets a derived denominator or it decays. Ours
  decayed by one adjective and five copy-forwards.
- **"Not covered" and "never read" are different claims, and only one of them was true.** A clause that
  does not say which it means will be read as the stronger one.
- **Build a language triple all at once, or the missing leg is the language you are writing in.**
