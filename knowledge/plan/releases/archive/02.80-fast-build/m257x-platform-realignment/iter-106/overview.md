---
iter: 106
milestone: M257x
iteration_type: tik
iter_shape: fence
status: closed-fixed
opened: 2026-08-06
---

# iter-106 — the drift fence: watch the inflow, not the puddle

**Type:** tik · **Active strategy: `TOK-06` step 1** — `FIX-M257x-iter103-drift-fence-gap`, which iter-103
ranked **above** the read-union repair.

## Step 0 — re-survey before targeting

Re-derived at this open, not inherited:

| measured now | value |
|---|---|
| `app/go.mod` requires `github.com/anthropos-work/ai` | **no** — the fold is real and the corpus calls it imported in 6 places |
| `sentinel/go.mod` proto pin | **`v1.210.0`** — `shared_libraries.md:85` says `v1.200.0` and calls the live skew *"two"* |
| `sentinel/go.mod` colony pin | **`v0.35.2`** — `clerkenstein.md:275` says sentinel is *"still on `v0.34.3`"* |
| clone HEADs vs the shas the corpus cites | `sentinel` is at `f2c46190`; the corpus cites `88bc5592` |

**The 33 anchors of iter-103 are still live** — iter-103 was a measuring pass with no repair inside it — so
the drift this fence targets is present in the tree **right now** and can be used as its own answer key.

Target confirmed, not substituted.

## Cluster / target identified

**61 % of `N` is clone-advance drift and no guard fences it.** `platform_alignment_guard` fences `repos.yml`
membership; `platform_predicate_guard` fences compose profile tokens; `derived_value_guard` fences two
scalars in a service doc against that service's own repo. **None of them notices that a clone moved.**

## Hypothesis — and it reframes what the fence should assert

The naive fence — *"a version the corpus states must equal the clone's"* — **cries wolf**, and the
measurement above shows exactly where. `shared_libraries.md:85` states `sentinel v1.200.0` and cites
`sentinel/go.mod:9 @ 88bc5592`; at **that** ref the claim was true. §5 rules 41/44 make this a
ref-scoped claim, so a fence that calls it *false* is asserting something it did not measure.

**So the fence must not adjudicate truth-at-a-ref. It must watch the ADVANCE**, which is the actual inflow:

> A clone the corpus cites has moved past the commit the corpus was last reconciled against, and here are
> the sites at risk.

That is a fence rather than a linter because it is a **ratchet against a checked-in baseline** — the same
shape `repair_postcondition` and `value_change_guard` already use. It cannot be green while a cited clone
has advanced un-reconciled, and it fires **at the moment the drift is born**, not one full reading cycle
later.

## Phase plan

`stack-core/clone_drift_guard.py`, two assertions:

- **D1 — cited-clone advance (the inflow watcher, the real deliverable).** A checked-in
  `clone_reconciliation_baseline.json` maps repo → the sha the corpus was last reconciled against. RED when
  a repo **the corpus actually cites** has moved past it, naming the citing sites. `--accept` re-baselines,
  and it is a **ratchet**: accepting is an explicit act that lands in a diff.
  > ⚠ **PLAN SUPERSEDED IN FLIGHT — `D-M257x-106-2`. There is no baseline file, and there should not be.**
  > A checked-in `repo → sha` map is **§2's hand-maintained tuple in a new costume**, and its *first* value
  > would have to be asserted rather than measured — there is no honest sha to seed it with, because
  > iter-103 proved the corpus is not reconciled to the current clones. The shipped fence derives the
  > baseline from **the corpus's own sha citations**, resolved per-repo with `git cat-file -t`. Recorded
  > here rather than quietly rewritten: a plan that changed and left no trace is the shape this milestone
  > exists to catch.
- **D2 — pin agreement, conservative form.** Fires only on the unambiguous parsed construct: a site naming
  `` `<repo>/go.mod` `` **and** a `<module> v<semver>` token. A site carrying a sha for that repo that the
  clone is not at is **UNMEASURED and named**, never silently passed (§8's three-verdicts rule).

Plus, per TOK-06's binding clause: a mutation battery and an anti-vacuity control **that can fire**.

## Expected lift

**The fence is expected to go RED on the live tree**, and that is the deliverable, not a failure — clause 4's
own wording is *"asserted by a FENCE that is watched going RED, not by inspection."* TOK-06 sequences the
fences **before** the repair precisely so the repair (step 3) has something watching it. A green here would
mean the fence cannot see the drift a double reading found.

**No movement on `N` is claimed.** Clause 3's instrument, not clause 5's.

## Escalation conditions

- If D2's recognizer produces any false positive on the live tree, D2 ships **narrower** or not at all — a
  fence that cries wolf gets suppressed, and a suppressed fence is worse than no fence.
- If D1 cannot name citing sites (i.e. the corpus cites no repo shas mechanically), the ratchet has no
  subject and the iter closes `closed-no-lift` with that falsification.

## Acceptable close-no-lift outcomes

D2 proving unbuildable at acceptable precision is a first-class outcome **if** it is recorded with the
false-positive measurement that killed it. D1 landing alone is `closed-fixed-partial`.
