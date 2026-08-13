# iter-254 — decisions

## `D-M257x-254-1` — the precondition predicate lives in the module that CENSUSES the class

`clone_set_present()`, `node_modules_present()` and their two reason strings ship in `suite_census.py`,
beside `--fresh-checkout` and `ENV_GATED`, not in a new helper module and not copied per test file.

The rule: **the predicate that flags a test as fresh-checkout-hostile and the predicate that excuses it
must be the same one.** Two copies drift, and the drifted state is the worst one available — a test
simultaneously excused by its own decorator and reported by the census, or reported by neither.

It also settles a smaller mess measured on the way in: the 12 holding files spell their root **six
different ways** (`parents[4]`, `parents[2]`, `parents[1]`, `parent.parent`, an env override, a
fixtures-relative path). `rosetta_root()` **walks** to the checkout — an ancestor carrying both `corpus/`
and `.git` — instead of counting levels, so the answer does not depend on where the caller happens to sit.

## `D-M257x-254-2` — `PR-1` refuted: the class needs TWO preconditions, and the second was invisible

The convenient story was one precondition (a clone set). Reading the real failure of all 22 on the frozen
tree refuted it:

| precondition | why the failure misreads |
|---|---|
| **the clone set** (`stack-demo/`/`stack-dev/`) | the corpus cites files that live in those git-ignored clones — `UPGRADE-IMPACT-next16.md`, cited by the corpus, lives at `stack-demo/next-web-app/` — so their absence prints *"resolves in neither pool and is undeclared"* |
| **an installed dependency tree** (`node_modules`) | three arms measure the TypeScript population or the prune-exclusion rate; a fresh clone has never run `npm install`, so they read *"the TypeScript population fell to 0 tests"* and *"prune_census removed nothing"* |

The second is invisible to inspection: `test_prune_census_can_return_NON_ZERO`'s own assertion message
says *"the tooling repo, **which has a node_modules**"* — the precondition was written down, in the
failure text, as an assumption. Reading the failures is what surfaced it; predicting them would not have.

## `D-M257x-254-3` — `PR-5` refuted, by a bug in this iter's own predicate, caught by the live half

The first `node_modules_present` globbed two levels. That is enough from the **rext** root
(`stack-verify/e2e/node_modules`) and **not** enough from the **rosetta** root, where the shallowest is at
four and this repo's own at five. One call site passes the rosetta root — so the predicate answered
`False` on a machine that plainly has one, and `test_prune_census_can_return_NON_ZERO`, which passes here,
**began to SKIP**.

That is the other way to break a suite, and it is quieter than a false RED: **a test that stops running
looks exactly like a test that passed.** It was caught by the *live* direction of the verification, not by
reading the diff — which is the whole argument for checking both directions rather than only the one the
iter is aiming at. Pinned by a regression test asserting the predicate answers `True` from **both** roots,
and bounded (`max_depth`) because an unbounded `rglob` under a root carrying ~180,000 files is not a
predicate anyone keeps calling.

## `D-M257x-254-4` — do not decorate a defect: the three battery failures were a CASCADE, and proving it cost nothing

`test_m257x_mechanical_fences_mutation_battery`'s 3 failures were never independent — the battery runs the
suite and asserts its baseline is green, so `test_18` and `test_22` failing made it fail three times over.
After those two declared their precondition, **2 of the 3 passed with no edit of their own**. The
remaining one (`test_01_every_mutant_matches_its_DECLARED_verdict`) did not, so it is *not* the same thing
and is routed rather than assumed.

The general form: **before writing a declaration, check whether the failure is downstream of another one.**
A decorator on a cascade hides the cause and permanently retires a control that was working.

## `D-M257x-254-5` — the ratchet re-pin is now a standing closing step of any iter that writes prose

Second consecutive iter to breach all-or-most of the three literal ceilings with its own edits, and this
time the COMMENT arrow **breached the ceiling it was written to raise** — the population includes its own
explanation, which that block's own docstring predicts. Convergence took two passes and is recorded as the
expected shape, not an anomaly. Reinforces `ROUTE-M257x-253-the-iter-loop-runs-no-ratchet`: the reading
costs ~1 s (`derivation_registry.py --ceilings`) and nothing in the loop calls it.
