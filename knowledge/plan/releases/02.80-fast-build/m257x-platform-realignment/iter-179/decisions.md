# iter-179 — decisions

## `D-M257x-179-1` — the routed repair is REFUTED, and the refutation is the deliverable

`FIX-M257x-iter174-accept-registers-one-registry-of-two` names its own repair: *"`--accept` registers one
registry of two"* — read forward, *make it write the second.* **It must not, and the reason is already
written in the artifact the route is about.** The second registry is the mechanical-fences battery's
`_COPY_FILES`, and `test_000` asserts *staged ⊇ the baseline's fence names*. A `--accept` that wrote both
from one source would make that assertion **compare a set with itself** — iter-158's shape, and the
battery's own seed-list comment says so in its own words.

**Decision: perform the refutation, not the repair, and record it where the next reader of the route
will hit it** — in the population fence's verdict text for that site, in the battery's own `test_000`
docstring, and here. *A routed item's proposed repair is a HYPOTHESIS, not a plan* (iter-158) — applied
for the fourth time on this milestone, and the first time to a route whose proposal was **wrong rather
than merely over-broad**.

## `D-M257x-179-2` — the direction the route feared is ENTAILED, and the entailment's PREMISE is now asserted next to its conclusion

The worry behind the route is that only one direction of a two-way contract is fenced. Measured, it is
one direction of a **sufficient pair**:

| conjunct | asserted by | cost |
|---|---|---|
| **P1** real tree: `baseline names == participating` | `test_iter45_mechanical_fences.py::test_21` | sub-second |
| **P2** `_COPY_FILES ⊇ baseline names` | the battery's `test_000` | sub-second |
| **⟹** staged tree: `staged-participating == staged baseline` | the staged run of `test_21` | minutes |

`participating ∩ staged = baseline ∩ staged = baseline` whenever `baseline ⊆ staged`, so P1 ∧ P2 give the
conclusion outright. **Nothing recorded that** — and P1 is not a stable premise: it was a **hard-coded set
of four** until iter-118, so the thing P2's sufficiency rests on has already been the weaker of the two
once. Weakening it back would silently un-sufficient P2 with no fence anywhere going red.

**Decision: assert P1 where P2 is asserted, and say why.** Not a second derivation — the same
`discover_fences()` + the same checked-in baseline, called from a second place, with the other call site
named in the docstring (§8 iter-175: *two derivations of ONE population must be COMPARED*; the cheaper
compliance is to have one derivation and two call sites).

## `D-M257x-179-3` — the defect was REACHABILITY, not cost; measured before it was believed

The route reads *"reported, but only by a ~14-minute battery, one iter after the fact."* Measured at
iter-179, units and runners named:

| reading | value |
|---|---|
| the whole contract computed standalone | **0.10 s** — `/usr/bin/python3` 3.9.6 **and** `/opt/homebrew/bin/python3` 3.14.6, agreeing |
| `test_000` alone | **0.04 s** — pytest 8.4.2 on 3.9.6, `-k test_000` |
| the battery containing it | minutes — it runs a nested unittest suite per mutant |

So the check was never expensive. It was **only reachable by naming a file whose name says
`mutation_battery`**, and the standing practice on this milestone is scoped runs that exclude exactly
those files. iter-173's own post-fix scoped re-run — **167 passed, green** — structurally could not see it.

**Decision: fix the reachability, keep the ordering.** The contract moves to
`tests/test_battery_baseline_stage_m257x.py`, which a plain scoped run over `tests/` reaches in **1.4 s**;
`test_000` stays and delegates to that module's `unstaged_fences()`, because **its ordering is its
deliverable** — it must fire before `test_00_` *inside* the battery so the failure arrives as *"you forgot
a file"* rather than *"the fence is broken."* → §8 rule.

## `D-M257x-179-4` — the obligation is NOT universal, and it is DECLARED per site rather than inferred

Derived, not remembered: **two** modules stage the postcondition baseline into a temp tree. Only one owes
anything.

| stager | staged participating fences | baseline names | unstaged | verdict |
|---|---|---|---|---|
| `test_m257x_mechanical_fences_mutation_battery.py` | **6** | 6 | 0 | **REQUIRED** — its staged suite carries `test_21` |
| `test_m257x_repair_postcondition_mutation_battery.py` | **1** (`claim_twin_guard`) | 6 | **5** | **DECLINE** — and it is green |

The decline is **proved rather than argued**: that battery is green today while five of the six
baseline-named fences are absent from its staged tree, because its staged suite is
`tests/test_repair_postcondition.py`, which calls `discover_fences()` three times and never compares the
result to the baseline (it asserts `on_disk - registry == set()`, a different claim). `grade()` books a
baseline-named fence absent from disk as `removed` and a discovered fence absent from the baseline as
`registered`, and **neither is a non-zero exit**.

**Decision: derive the POPULATION, declare the VERDICT.** Membership is computed on every run from
`_COPY_FILES` + `repair_postcondition.BASELINE_REL` (read from its owner, never spelled — §8 iters 70/71),
so a battery that starts staging the baseline tomorrow enrols itself and is RED until classified.
Inferring the verdict would mean deciding *"does this staged suite assert the equality?"* from source —
the wrong-construct guess this milestone spends its iters finding.

## `D-M257x-179-5` — a new TEST module owes no registry; a new FENCE module owes four. The distinction is measured

iter-178 put its new arm **inside an existing guard** precisely to avoid the registry tax — *"a new module
drags four registries behind it and three have been caught rotting here."* One iter later this iter adds a
new module, so the exemption is stated with its evidence rather than assumed: the four registries
(`guard_family.INVOCATIONS`, `derivation_registry.DECISIONS`, the postcondition baseline, the battery seed
lists) are all keyed on **fences** — `*_guard.py` files and `FENCE_KIND` declarations — and a `tests/`
module is in none of those populations. Checked, not asserted: with the new file present,
`test_fence_registry_population_m257x.py` + `test_fence_registry_completeness_m257x.py` are **23 passed**
and `test_frozen_expectation_census_m257x.py` + `test_guard_family.py` are **77 passed**.

**One registry it DOES owe: the prose one.** `stack-core/README.md` indexes batteries and test modules
alongside guards, and it carried **no row** for this module — nor, as it turns out, for **iter-176's
`test_fence_registry_population_m257x.py`**. Both rows are added here. That the index had been missing a
three-iter-old fence is a fresh datum for
`SURVEY-M257x-iter175-readme-fence-index-is-16-of-27`: the disclosed limit is measured over *fences*, and
the index's coverage of **test modules** was never measured at all.
