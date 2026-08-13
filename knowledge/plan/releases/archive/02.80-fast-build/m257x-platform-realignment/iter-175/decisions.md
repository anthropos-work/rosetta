# iter-175 — decisions

## `D-M257x-175-1` — the route's item was third by size; the target is substituted, the strategy is not

`FIX-M257x-iter174-accept-registers-one-registry-of-two` names the mechanical-fences battery's fence-seed
list. Phase 1 Step 0's mandatory re-survey ran the census that route asks for **before** targeting, and
the seed list is the **third** row by size — and it is already reported, late but correctly, by the
battery's own `test_000`. The top row is reported by nothing.

**Decision: substitute the target under `TOK-08`, keep the route open.** `TOK-08` names a *method* —
*work the classes in descending measured size* — not a file. Substituting a bigger member of the same
class is the strategy executing, not a re-scope. The seed-list route stays open and unmodified.

## `D-M257x-175-2` — instrument B over instrument A, and the rejected one is recorded

| instrument | predicate | population |
|---|---|---|
| A | a file naming ≥2 fence modules as quoted literals, anywhere | **39** |
| B | a **collection literal** (py `List`/`Tuple`/`Set`/`Dict`-keys/call-args, or a JSON array/object) holding ≥2 fence-module names | **5** |

**Decision: B.** A measures *mentions*; the claim is about **a set that must track the tree**, so A is
graded at the wrong grain (§9, iter-159). A's 39 would have produced 39 classification lines, ~34 of them
declining a test that happens to name two guards — volume that reads like rigour and answers a different
question. Both are recorded because *the rejected instrument is part of the measurement*: had B come back
empty, A's 39 is where the next reading would start.

Neither script is checked in (read-only, `.agentspace/scratch/work-m257x/`). Both are reproducible from
the predicate as stated; the checked-in artifact is the **fence**, per §8.

## `D-M257x-175-3` — the repair is the UNION, and NOT the symmetry with iter-157

The obvious move is to swap `glob("*_guard.py")` for the `FENCE_KIND` declaration, exactly as iter-157 did
in `repair_postcondition`. **Rejected: it is a weakening dressed as a tidy-up.** A `*_guard.py` that
declares nothing would silently leave the family — and that file is precisely the one worth catching.

**Decision: `census := spelled ∪ declared ∪ extra`.** A member needs only ONE property, so both gaps close
at once and neither can lose a member. This is the only direction iter-158's rule permits (*a narrowing
that grades a broken check green is a defect, not a fix*), and it is why a proposed repair is a hypothesis
until it is graded against what it would let through.

`EXTRA_CENSUS_MEMBERS` survives, demoted in the comment to **additive only**: it may add a member, never
substitute for a property. `repair_postcondition` declares no `FENCE_KIND` (it is the module that *reads*
them) and is not spelled `*_guard`, so neither derived property reaches it and naming it is honest.

**One reader, two consumers.** `declaring_modules()` calls `repair_postcondition.declared_kind` rather
than re-implementing the AST read. A private copy of the rule here would be a *third* derivation of the
same population — this iter's own defect, committed while repairing it.

## `D-M257x-175-4` — `predicate_enumerator` is invoked, not excluded; and its NOT-RUN reason had to be true

It requires `--ledger`, exactly like `repair_reach_guard`, so it is an `input`-class member. But `run_one`
derived the *range* requirement from the class — `cls in ("commit", "input")` — and `predicate_enumerator`
takes no `--range`. Left alone it would report **"needs --range, not supplied"**, a NOT-RUN reason naming
a flag the guard does not accept, sending the reader to fix the wrong thing.

**Decision: `needs_range` is DECLARED per member, defaulting to the class rule** (`spec.get("needs_range",
spec["cls"] in ("commit", "input"))`). The default keeps every existing member's behaviour byte-identical;
the one member that differs says so. A reason that is not true is not a reason.

## `D-M257x-175-5` — an excluded member is PRINTED, and the exclusion table is held to both directions

`guard_family` cannot run itself. But this module's founding sentence is §5 rule 8 — *a guard that was not
run reads exactly like a guard that passed* — and a member omitted from the census is that sentence with
the evidence removed.

**Decision: `CENSUS_EXCLUSIONS` is a table of reasons, printed on every run**, and `reconcile()` gains its
two directions: an exclusion the census does not reach is **stale** (it subtracts nothing, so it excuses
nothing, and it is a place a future member can be parked), and an exclusion also named in `INVOCATIONS` is
**ambiguous**. Without those, the table is a place to put a member to make `reconcile()` quiet — the
failure mode, not the feature.

## `D-M257x-175-6` — four synthetic fixtures staged a guard dir with no runner in it, and one asserted the right code for the wrong reason

`test_an_unreconciled_census_exits_2_not_0` writes one `orphan_guard.py` and asserts `main()` returns 2.
Once `CENSUS_EXCLUSIONS` landed it still returned 2 — from a **stale-exclusion** complaint, with the
orphan never reached. The test passed throughout and proved nothing.

**Decision: make the fixtures faithful AND assert the sentence.** `stage_runner()` writes the runner into
every synthetic guard dir (a guard dir without `guard_family.py` is not a guard dir, and since this iter
the difference decides the run), and the exit-2 test now reads stdout and asserts it names `orphan_guard`.
**An exit code is not a diagnosis.** This is iter-91's *grade the cannot-tell* and §8's *write the control
against the guard's SUBJECT* arriving together, in a test that had been green for 89 iters.
