---
iter: 176
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
---

# iter-176 — ship the instrument as a fence: the registry POPULATION, not its last member

**Type:** tik · **Active strategy:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Step 0 — re-survey (mandatory)

iter-175 built the instrument the standing route asks for and **used it once**. The route it was built
for is still open, and it is open at its population:

> `FIX-M257x-iter174-accept-registers-one-registry-of-two` — *"Five registries are now known; **nothing
> enumerates them**."*

Re-run at HEAD (`3b5a82d` / rext `5b108d0`), unchanged in shape, one row moved by iter-175's own repair:

```
fence population = 26
SITES with a collection literal holding >=2 fence names: 5
   25  stack-core/guard_family.py:85                                    (was 24 — predicate_enumerator joined)
    6  stack-core/repair_postcondition_baseline.json
    6  stack-core/tests/test_m257x_mechanical_fences_mutation_battery.py:70
    3  stack-core/tests/test_iter45_mechanical_fences.py:384
    2  stack-core/tests/test_fence_registry_completeness_m257x.py:83
```

**`TOK-08`'s sentence is *"build or extend a fence that enumerates every instance in the corpus, run it to
zero, and keep it green."*** iter-175 did the first clause with a scratch script that was deliberately not
checked in — correct for that iter, whose deliverable was the union repair, and **exactly the gap this
iter closes.** A reading SAMPLES; a fence CENSUSES; an instrument that ran once and was deleted is a
reading wearing a fence's clothes.

The target is live and untouched: nothing in the tree enumerates the five, and iter-175's own routes
demonstrate the cost twice over — a **sixth** registry (the section README, 16 of 27) was found by hand,
and the **fifth** (`derivation_registry`) announced itself only by going RED 34 minutes into a
whole-population run.

## Cluster / target identified

**The population of REGISTRIES is itself unfenced.** iter-169's rule — *closing a class means fencing its
POPULATION, not its last member* — applied to the class iter-174 named and iter-175 measured. Every repair
in this thread has added **one** member to **one** registry and discovered the next registry by running a
34-minute suite:

| iter | registry found | how it was found |
|---|---|---|
| 173 | `guard_family` · `derivation_registry` · `fence_provenance` | grep for a sibling's name |
| 173 | `repair_postcondition_baseline.json` (4th) | whole-suite run, after the fact |
| 174 | the mechanical-fences battery seed list (5th) | whole-suite run, one iter later |
| 175 | `derivation_registry` again + a 6th (README, prose) | whole-suite run + a hand check |

**Four consecutive discoveries by the most expensive instrument available.** That is the signature of a
population nobody enumerates.

## Hypothesis

Shipping iter-175's instrument B as a **checked-in fence with a declared classification table** — every
collection literal holding ≥2 fence-module names must be classified `REGISTRY:<who keeps it in sync>` or
`DECLINE:<class>: <reason>`, and an **unclassified site is RED** — moves the discovery of registry #7 from
a 34-minute suite run to a sub-second static check, and closes `FIX-M257x-iter174-…` at its population
rather than at its last member.

## Expected lift

No `P`/`N` reading. Deliverable: the enumeration is checked in, at zero unclassified, with a mutation
control and an anti-vacuity control that can actually fire; and the standing route closes.

## Phase plan

1. **Author** `tests/test_fence_registry_population_m257x.py` — instrument B, hardened, plus the
   classification table and its two directions.
2. **Classify** all five sites, per site, with a reason (`D-M257x-159-4`'s rule: explicit, never inferred).
3. **Disclose the instrument's grain** — the README row is *prose*, not a collection literal, so this
   fence cannot reach it. State that in the fence, on every run, the way `derived_count_guard` prints its
   NOT-REACHED clause.
4. **Whole-population run** — `FIX-M257x-iter142-whole-suite-owed`, which has now paid for itself four
   consecutive iters, three of them on the iter's own commit.

## Escalation conditions

- If a site cannot be honestly classified either way, record it and route — do not invent a decline class
  to reach zero.
- If the fence's own file becomes a sixth site of its own subject (it will name fence modules in its
  controls), that is **not** a reason to weaken the predicate: it is the `D-M257x-173-6` case — a census
  whose subject includes its own report — and it is handled by classifying it, not by excluding it.

## Acceptable close-no-lift outcomes

If the five sites turn out not to share a checkable obligation — i.e. "registry" is not one class but
several with nothing in common — the iter closes no-lift with that falsification recorded, and the route
is re-shaped rather than re-attempted.
