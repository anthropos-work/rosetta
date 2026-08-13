# iter-173 — decisions

## `D-M257x-173-1` — split the routed class by DERIVABILITY, and say which half the green covers

iter-172 routed *"every pytest count published before this iter is a `passed` count, therefore an
undercount of the executed population … nobody has enumerated where they are."* Taken whole, that class
is unaffordable: most of its members are **observed** counts read off a runner at a ref that no longer
exists, and re-checking one means re-running a suite — the cost `§5` rule 51's unusable timing leg and
iter-172's own 50-minute two-runner census both bill for.

**Decision: split the class by whether a member carries its own operands.**

- **DERIVED** — a total over a table, a delta over its endpoints, a percentage over its own numerator and
  denominator. Re-derivable *on the page*, with no runner, no host, no clone. **Censusable, and the
  census runs in under a second.**
- **OBSERVED** — `1 failed · 1229 passed`, read off a runner. Not re-checkable here at any price this
  iter can pay. **Stays routed, and this iter's green says nothing about it** (`§5` rule 60).

This also **qualifies the routed item's own premise**, per the standing rule that a route's proposed
repair is a hypothesis (iter-158). *"A `passed` count is therefore an undercount"* does not hold
universally: it is exact whenever the run had zero failures and zero skips. This iter's own controls are
the demonstration — **12 tests, `Ran 12` under unittest/3.14.6 and `12 passed` under pytest/3.9.6**, the
two units coinciding because there is nothing to drop. The defect is not that `passed` was published; it
is that `passed` was published **under the word `tests`** and then summed.

## `D-M257x-173-2` — repair the claim's own sites; ROUTE the ones another skill owns

Five sites publish the false denominator. Two of them are in `hardening-ledger.md`, which
`/developer-kit:harden-mstone-iters` owns exclusively — `build-mstone-iters` reads that file and does not
write it.

**Decision: repair the three sites this iter owns, and route the two it does not with the derivation
pre-computed**, so the next harden pass applies an answer rather than re-deriving one. Handler:
`FIX-M257x-iter173-ledger-denominator`. A correction that respects file ownership arrives later and
intact; one that does not arrives as a merge conflict.

## `D-M257x-173-3` — a verbatim quote of a false claim is left VERBATIM and annotated beneath

`iter-145/overview.md:20` block-quotes the ledger sentence carrying the wrong figure. Editing the quote
would make the record cite a sentence that was never written.

**Decision: the quotation stands; the correction goes underneath it, naming what the quote asserts and
what is true.** The claim's *own* sites are where the number changes. This is the same
correction-vs-retraction seam `FIX-M257x-iter144-correction-vs-retraction-unfenced` still has open, and
this iter is one more instance for it, not a closure of it.

## `D-M257x-173-4` — fence the OPERANDS when the claim itself is not machine-reachable

`N of M` prose is not derivable: `M` names no source, and deciding which nearby table a sentence
summarises is exactly the inference this milestone has spent nine iters proving instruments cannot afford
(`D-M257x-117-2`, iter-119's refutation).

**Decision: do not build an inference; build the guard one level down.**
`stack-core/derived_count_guard.py` fences the **table totals, explicit deltas and percent-triples** the
`N of M` claim must be derived FROM. The prose repair stays a hand repair — but it now rests on a fenced
ground truth instead of a second reading. **The tool prints its NOT-REACHED clause on every run, green
included**, so the reach can never be quoted without its limit (`§5` rule 68).

**Measured reach, with its denominator:** 25 derivable sites evaluated across **698 markdown files**
(240 excluded — `raw/` and `evidence/` hold pre-adjudication seat artifacts and captured logs, which are
inputs committed verbatim by protocol, not claims; editing one to satisfy a fence would corrupt an
evidence artifact). Arms: 16 table-totals · 5 explicit-deltas · 4 percent-triples.

## `D-M257x-173-5` — the new fence was graded by the fence family before it was trusted, and FAILED three times

The guard's own controls were green (12/12, both runners) while the guard was still **non-conformant with
the family it was joining**. Three separate family fences caught it:

| fence | what it caught |
|---|---|
| `test_fence_provenance::test_every_member_stamps_on_direct_execution` | no `fence_provenance.stamp_main()` — the verdict did not state the tree it was taken with |
| `test_fence_provenance::test_machine_mode_is_parseable_and_self_describing` | `--json` did not ride the tree inside the document |
| `repair_postcondition` | declared `FENCE_KIND = "postcondition"` while exposing no `postcondition_sites()` — a **copied declaration**, the registry-rot shape |

And the family view surfaced a fourth by inspection: the guard printed **nothing** when green, so
`guard_family.run_one`'s `lines[-1]` rendered it as a blank summary. All four fixed; the one-line verdict
now prints last and carries its reach clause with it.

**The lesson, and it is `§8`'s own rule turned on its author:** *a new fence must pass the fence family's
contracts before its own green means anything.* A guard that is right about its subject and wrong about
its interface is invisible in the one view anybody reads.

## `D-M257x-173-6` — a census whose SUBJECT includes its own REPORT has a moving population

The guard scans the published record. **This iter's report is part of the published record.** So the
moment the census table was written into `iter-173/progress.md`, the population it reports grew — by
exactly two, the two numeric columns of that table's own `total` row.

Measured: **25 evaluated** at the tree the census ran on, **27** once this report joined it. Both runs
green.

**Decision: publish the number WITH the tree it was taken at, and state the self-reference, rather than
silently restating it as 27.** Restating would have been arithmetically fine and epistemically wrong — it
would present a figure as a standing fact when it is a reading, which is the precise failure this whole
iter is about. It is `§5` rule 51(b)'s *state-the-ref* discipline reached from an unusual direction: not
"the tree may change under you" but **"your own act of reporting changes it."**

The practical corollary for any future fence over the milestone record: **its own iter's write-up is
inside its subject.** That is not a bug to exclude — the report's arithmetic should be checked like
anything else, and here it is — but a green quoted without its ref will drift by the size of whatever the
iter wrote.

## `D-M257x-173-7` — the whole-suite run paid for itself again, and the registry enumeration was incomplete

The scoped runs were green. The **whole `stack-core` suite** — `3 failed · 1,525 passed in 1,368 s
(22:48)`, `/usr/bin/python3` 3.9.6 (state the runner: iter-170) — returned three, and the grading matters:

| failure | whose |
|---|---|
| `test_iter45_mechanical_fences::test_21_the_shipped_baseline_records_EVERY_participating_fence` | **THIS ITER'S.** A net-new regression: the new fence was absent from the repair-postcondition ratchet baseline |
| `test_frozen_expectation_census::test_every_executable_derivation_is_classified` | **pre-existing, and this iter ADDED a third member to it** (`derived_count_guard.py::postcondition_sites`) |
| `test_battery_stage::test_a_stdlib_shadow_is_refused_not_staged` | pre-existing, unrelated (`RuntimeError not raised`, no mention of this iter's module) — **stays routed** |

**The finding underneath the finding.** Before the run, this iter had enumerated the registries a new
fence must join by grepping the tooling for a sibling's name: `guard_family.py`, `derivation_registry.py`,
`stack-core/README.md`. **That enumeration was incomplete and its incompleteness was invisible** — the
fourth registry is a **JSON ratchet baseline** (`repair_postcondition_baseline.json`), which names no
sibling module in any `.py`, so a name-grep over source could not see it. Only running the suite did.

**This is `FIX-M257x-iter142-whole-suite-owed` firing for the third consecutive iter**, and the sharpest
instance yet: *a change-derived scoped suite cannot see the fence that grades it* — and here the author
had explicitly tried to derive the scope by enumeration first, and the enumeration was itself a list, not
a derivation (`§5` rule 73: **a glob is not a derivation** — nor is a grep).

**Decision: fix both of this iter's own contributions; leave the unrelated one routed.** The
derivation-registry entry could not be verified in isolation — a single test fails if *any* site is
unclassified — so the two pre-existing entries were classified as well, closing that half of
`SURVEY-M257x-iter172-two-preexisting-actionable-reds`. **Entanglement forced the closure: a RED this
iter contributes to cannot be shown fixed while other contributors remain.** All three were read before
classification and given three *different* reasons, because a blanket verdict would have hidden that
`suite_census.py::modules` is a tree-scan while `::stale_declarations` is a verdict.
