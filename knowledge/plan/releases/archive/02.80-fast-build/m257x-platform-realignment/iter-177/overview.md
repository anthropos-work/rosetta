---
iter: 177
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
controlling_strategy: TOK-08
---

# iter-177 — the retraction was wrong: `16 of 27` is a correct reading of a population nobody named

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)
— *census the mechanical classes; stop sampling them.* This iter works a class that is decidable by
arithmetic over sets: **how many distinct derivations of "the fence family" are live in the tree, and do
the numbers published about it name which one they came from.**

## Step 0 — re-survey before targeting (mandatory)

`TOK-08`'s standing direction is "next mechanical class, descending size". The routed queue's top items
are the two `SURVEY-M257x-iter175-*` entries plus harden pass 39's `FIX-M257x-h39-survey-id-embeds-
retracted-figure`. Re-surveyed at rext `c7f4c3d` / corpus `84b127a2` before targeting:

* `stack-core/README.md` still names a strict subset of the fence family — the gap is real and unchanged.
* The pre-flight (`§0d`) ran the live instrument first: `tests/test_fence_registry_population_m257x.py`
  → **8 passed in 6.07 s** under `/usr/bin/python3` (3.9.6, the only interpreter here with pytest —
  `§9` measurement preconditions, r75/76 *name the runner*). Green before any edit.

**The re-survey changed the target.** The queued item said *"the README index is 16 of 27, and harden
pass 39 retracted that to 15 of 26 because both operands were wrong at publication."* Measuring it
first — the cheapest possible check, three lines of Python — falsifies the retraction:

| derivation | what it is | live count | named in `README.md` |
|---|---|---|---|
| `guard_family.union(HERE)` | spelled `*_guard.py` ∪ declared `FENCE_KIND` ∪ `EXTRA_CENSUS_MEMBERS` | **27** | **16** |
| `guard_family.census(HERE)` | `union` − `CENSUS_EXCLUSIONS` | **26** | **16** |
| `fence_population()` / `declaring_modules` / `discover_fences` | declares `FENCE_KIND` | **26** | **15** |

Re-derived at **`5b108d0`** (iter-175's own commit, reconstructed with `git archive` — the same method
harden pass 39 used): **identical**, 15 of 26 · 16 of 26 · 16 of 27.

So `16 of 27` is not an arithmetic error. It is an **exact** reading of `union`. iter-175's own routed
text proves which set it enumerated: it lists the **11** missing members *by name*, and that list
**contains `guard_family`** — which `census` excludes and `declaring` includes. Only `union` has 27
members with those 11 absent.

## Cluster / target identified

`TOK-08` named "the next mechanical class". The re-survey substitutes a target inside the same class and
one layer up: **not the README's incompleteness, but the fact that three derivations of one population
are live at once and every number published about it omits which one it came from.**

The substitution is recorded per Phase 1 Step 0: *the queued survey named the README gap; the re-survey
shows the figure it disputes was correct, and that the actual defect is the missing population label.*

**And the cardinality is a coincidence that hides it.** `census` and `declaring` are **both 26** and
differ by **one member in each direction** — `census` has `repair_postcondition` and not `guard_family`;
`declaring` has `guard_family` and not `repair_postcondition`. **Any comparison of these two sets by
COUNT reads green.** That is this milestone's running thread — *the instrument measured something other
than what it printed* — in its purest arithmetic form, and it is the reason a retraction written by a
careful pass landed on a third population without noticing there were three.

## Hypothesis

If the disclosure is made to publish **all three derivations, each labelled with the function that
produced it**, and if the two 26-member sets are compared **by membership rather than by cardinality**
with every difference carrying a written disposition, then:

1. the `16 of 27` retraction is itself retracted, at the site that carries the live claim;
2. a fourth derivation cannot be silently introduced (it has no row, so the test goes RED);
3. the coincidence that made a count-comparison green is fenced by construction.

## Expected lift

**No `P`/`N` reading is taken this iter, so no clause-5 movement is claimed** (`§9` iter-type refinement
— an iter that took no reading has an UNMEASURED metric, not an unmoved one; `TOK-08` declares the
class-by-class sweep order in advance). The lift is instrument-side and is stated in counts:
derivations published with a name **0 → 3**; set-comparisons of the family asserted by membership
**0 → 1**; retracted-in-error figures restored **1**.

## Phase plan

* **A — measure** (done in Step 0): three derivations at HEAD and at `5b108d0`.
* **B — repair the live instrument**: correct the retraction in
  `tests/test_fence_registry_population_m257x.py`, publish the labelled triple, assert each against its
  own derivation.
* **C — fence the coincidence**: a membership reconcile of `census` vs `declaring` whose differences must
  each be dispositioned, with a mutation control that a count-only comparison would pass.
* **D — run**: the module, then the `stack-core` suite scoped to the fence family, under the named runner.
* **E — route**: the harden ledger's own entry (this skill may not write `hardening-ledger.md`), and the
  README gap.

## Escalation conditions

* If the three derivations turn out to be **the same set** at some ref, the whole premise collapses →
  close `closed-no-lift` with the falsification recorded. (Checked at two refs already; it does not.)
* If correcting the disclosure requires editing `hardening-ledger.md`, **stop and route it** — that file
  is owned by `/developer-kit:harden-mstone-iters` and this skill is forbidden from writing it.
* If a suite run is needed beyond the fence family's own modules, prefer a **scoped** run plus an honest
  statement of what it did not cover (r60/66).

## Acceptable close-no-lift outcomes

* The measurement shows harden pass 39 was right and iter-175's figure was wrong → record it, restore
  nothing, close with the falsification. (Falsified in Step 0.)
* The membership reconcile turns out to already exist somewhere → the fence is redundant; say so and
  close with the pointer rather than shipping a duplicate derivation, which is the very defect.

## Out of scope, routed not taken

* **Filling `stack-core/README.md`'s missing rows.** A second line of investigation, and a different
  class (index completeness, not derivation labelling). It stays a disclosed blind spot pinned by a test
  — which is where an open obligation belongs (iter-176's own rule). The route now has a precise target:
  it must say *which* derivation it intends to be complete against.
