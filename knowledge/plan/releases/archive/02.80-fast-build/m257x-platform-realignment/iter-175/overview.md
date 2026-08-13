---
iter: 175
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
---

# iter-175 — two derivations of ONE population, disagreeing, with nothing comparing them

**Type:** tik · **Active strategy:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Step 0 — re-survey (mandatory)

`TOK-08`'s standing direction is *work the mechanical classes in descending measured size, build a fence
that enumerates every instance, run it to zero.* iter-174 routed
`FIX-M257x-iter174-accept-registers-one-registry-of-two` with a sentence that is itself a class, not an
item:

> *"Five registries are now known; **nothing enumerates them**."*

Re-surveyed at open. The "five" is a **remembered list** — iter-173 said four, iter-174 said five, and
each number was reached by grepping for a sibling's name. That is §2's hand-maintained tuple one level
up: the registry **of registries**.

So the re-survey ran the census the route asks for, mechanically, before picking a target. Two
instruments, deliberately, because the first one's answer was not gradeable:

| instrument | predicate | population |
|---|---|---|
| A (rejected) | a file naming ≥2 fence modules as quoted literals, anywhere | **39** sites |
| B (used) | a **collection literal** (py `List`/`Tuple`/`Set`/`Dict`-keys/call-args, or a JSON array/object) holding ≥2 fence-module names | **5** sites |

Instrument A measures *mentions*; the claim is about a **set that must track the tree**, so A is graded
at the wrong grain (§9, iter-159). B is the claim's own grain, and it lands on exactly the sites the
route was groping for. Script: `.agentspace/scratch/work-m257x/` (read-only, not checked in).

**Instrument B's five, with the fence population derived from `FENCE_KIND` declarations (n = 26):**

| site | names | how its membership is derived | completeness fenced? |
|---|---|---|---|
| `guard_family.py:78` `INVOCATIONS` | 24 | reconciled against `census()` **both ways** | yes — but see the target below |
| `repair_postcondition_baseline.json` | 6 | written by `--accept` | ratchet |
| `tests/test_m257x_mechanical_fences_mutation_battery.py:70` `_COPY_FILES` | 6 | **hand-maintained** | only by a 14-min battery (`FIX-M257x-iter174-…`) |
| `tests/test_iter45_mechanical_fences.py:384` | 3 | test subject, not a registry | n/a |
| `tests/test_fence_registry_completeness_m257x.py:83` | 2 | the iter-157 regression pin | n/a |

The route's own named item (the battery seed list) is **third** by size and is already reported, late but
correctly, by `test_000`. The re-survey found something **larger and silent one row above it**, so the
target is substituted under the same strategy — `TOK-08` names the *method*, not this file.

## Cluster / target identified

**Substitution (Phase 1 Step 0):** the route named the battery seed list; the census shows the top row is
the defect. Same strategy, bigger member, and the smaller one stays routed.

**`guard_family.census()` selects the fence family by FILENAME SPELLING. `repair_postcondition.discover_fences()`
selects the same family by DECLARATION. They disagree by three members, in both directions, and no fence
compares them.**

```
stack-core/guard_family.py:222-224
    def census(guard_dir: Path) -> list[str]:
        """Every guard on disk, DERIVED. This is the half that must not be hand-maintained."""
        found = sorted(p.stem for p in guard_dir.glob("*_guard.py"))
        return found + [m for m in EXTRA_CENSUS_MEMBERS if (guard_dir / f"{m}.py").exists()]
```

`EXTRA_CENSUS_MEMBERS = ("repair_postcondition",)` — a hand-maintained escape hatch **substituting for**
the property, which is the shape r70/71 names: *a fence pinned to a SPELLING is not pinned to a
PROPERTY.*

**This is iter-157's defect, in the sibling module, still live.** iter-157 measured *"25 modules declared a
`FENCE_KIND`; 23 were enumerated"* in `repair_postcondition`, repaired it to walk `*.py` and select by
declaration, and shipped `test_fence_registry_completeness_m257x.py` — **which fences that one module's
registry only.** iter-169's rule, one turn on: *closing a class means fencing its POPULATION, not its
last member.*

**The measured disagreement, at open:**

| set | n | members |
|---|---|---|
| declares `FENCE_KIND` | 26 | 24 × `*_guard.py` + `guard_family` + `predicate_enumerator` |
| `guard_family.census()` | 25 | 24 × `*_guard.py` + `repair_postcondition` |
| declared ∖ censused | **2** | `guard_family` (self), **`predicate_enumerator`** |
| censused ∖ declared | **1** | `repair_postcondition` (declares no `FENCE_KIND`) |

**The consequence is the one guard_family exists to prevent.** `predicate_enumerator.py:142` declares
`FENCE_KIND = "standalone"` — it says of itself that it is a fence — and the runner whose docstring
promises to *"run the WHOLE guard family, and name every member"* **has never run it and has never named
it.** Its own §5 rule 8 is the indictment: *a guard that was not run reads exactly like a guard that
passed.* Today it does not even read as NOT-RUN; it is absent from the census, so the family's verdict
line is silent about it.

## Hypothesis

Deriving the census as the **UNION** of the two properties — `*_guard.py` on disk **∪** modules declaring
`FENCE_KIND` — plus the existing additive `EXTRA_CENSUS_MEMBERS`, is **strictly stronger in both
directions** and narrows nothing: a `*_guard.py` that declares nothing stays in (today's behaviour), and a
declaring module that is not spelled `*_guard.py` enters (today's gap). `reconcile()` then forces every new
member to be disposed of — an invocation, or a declared exclusion with a stated reason.

**Explicitly NOT the repair (iter-158's rule — a proposed repair is a hypothesis, and a narrowing that
grades a broken check green is a defect):** switching `census()` *from* spelling *to* declaration. That
reads as the obvious symmetry with iter-157 and it would **weaken** this runner — a `*_guard.py` with no
`FENCE_KIND` would silently leave the family. The union is the only direction that cannot lose a member.

## Expected lift

No `P`/`N` reading. The deliverable is a class closed at its population: two derivations of one set
reconciled, the third member disposed of by invocation rather than by omission, and a fence that goes RED
when any future registry of this family disagrees with the declaration set.

## Phase plan

1. **Census** (done in Step 0) — instrument B, denominator stated, both directions.
2. **Repair** — `census()` := union; `predicate_enumerator` into `INVOCATIONS` as `cls: "input"` (it
   requires `--ledger`, exactly like `repair_reach_guard`); a declared exclusion table with per-member
   reasons for anything deliberately not run; `reconcile()` gains the third direction.
3. **Fence** — a test asserting the two derivations agree modulo the declared exclusions, with a
   mutation control that kills the shipped form and an anti-vacuity control that fires when the
   instrument stops seeing the tree (`§9`).
4. **Whole-population run** — `FIX-M257x-iter142-whole-suite-owed`: a change-derived scoped suite cannot
   see the fence that grades it. The scoped run does not close this iter.

## Escalation conditions

- If the union makes `guard_family` exit 2 for a member with no honest invocation **and** no honest
  exclusion reason → that is a real design question; record it, do not invent a reason to reach green.
- If `predicate_enumerator` cannot be invoked from the family's context at all → route forward rather
  than weaken the census to hide it.

## Acceptable close-no-lift outcomes

If the union turns out to be **already** equivalent to the declaration set on this tree (i.e. the
disagreement is an artifact of the instrument, not the tree), the iter closes no-lift with the
falsification recorded and instrument B retracted in place.
