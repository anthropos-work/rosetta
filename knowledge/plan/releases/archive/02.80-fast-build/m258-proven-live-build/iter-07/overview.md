---
iteration_type: tik
status: closed-fixed-partial
milestone: M258
iter: 07
opened: 2026-08-12
---

# M258 iter-07 — the composed 3× cold campaign (`TOK-01` step 4)

**Active strategy reference:** `TOK-01` — *measure the composition before engineering it* — **step 4**,
the last of its four strictly-ordered steps. Steps 1 (measure both halves), 2 (wire the batch gate) and
3 (the world-contract restore leg) are all discharged; the strategy's own text names step 4 as
*"the composed 3× cold campaign against the gate, with the spread published beside the p50."*

## Step 0 — re-survey before targeting (mandatory)

The TOK-directed target is **still current and still the right one**. Re-surveyed against the live tree,
not assumed:

| check | reading | verdict |
|---|---|---|
| gate measurable? | `up-injected.sh:2839` carries the batch hook (iter-06, proven live) | ✅ unblocked |
| composed p50 taken? | **no** — iter-06 took n=1 at `green:false`, correctly not gate-usable | ⬜ the target |
| batch half p50? | **no** — two samples (129 s iter-04, 160 s iter-06), no p50 | ⬜ `C2` owed |
| bring-up half | gateable single-box **247.79 s** (iter-05) | ✅ reference |

**Nothing has absorbed this target.** No substitution needed.

## The precondition iter-06 discovered — verified, not assumed

iter-06 closed with: *run the campaign from the **consumption clone at a pushed tag**, not the authoring
copy* — an authoring-copy bring-up has no sibling `platform/`, so `autoverify`'s `postgres-schemas`
probe **refuses** to assert and every rep grades `green:false` whatever its timings say. That is the
guard working (`verification.md` § *A guard that cannot find its subject must not exit 0*), not a bug to
route around.

Re-verified this iter against evidence, per the standing *"a routed fix is a hypothesis"* rule — and it
holds in **both** directions:

- **The tag is on origin.** `git ls-remote --tags origin` → `fast-build-m258-iter-06` = `5f3a3815`,
  and the authoring copy's HEAD is the **same sha** with a clean tree. Rung zero satisfied.
- **The consumption clone is genuinely one tag behind**, and the gap is the feature under test:
  at `fast-build-m258-iter-03` **all four** batch-gate files are absent from
  `stack-demo/rosetta-extensions/playthroughs/e2e/` (`batch-gate.sh`, `restore-presenter-world.sh`,
  `check-cockpit-roster.py`, `stack-paths.sh`) and `up-injected.sh` carries **no** `batch-gate` hook.
  This is the M236 shape exactly: *the feature under test could not be obtained at all.*
- **The topology the probe needs is present there.** `stack-demo/platform/repos.yml` exists;
  `.agentspace/platform` does not. So the consumption clone satisfies the sibling-`platform/`
  requirement the authoring copy failed.
- **`buildbench` drives the clone it lives in** — `rext_root()` is `Path(__file__).parents[1]`
  (`demo_knob_guard.py:76`, imported at `buildbench.py:85`), and `argv_up` is
  `ext/"demo-stack"/"up-injected.sh"` (`:1477`). So running the **consumption clone's** `buildbench`
  drives the **consumption clone's** bring-up. The campaign must be launched from that path.

## Hypothesis

With the batch gate wired and the campaign run from a clone whose `autoverify` can actually assert, a
3-rep cold campaign yields **3 green, gate-usable reps** whose composed **p50 ≤ 480 s**.

## Expected lift

The projection standing at iter-06's close is **414.15 s** (gateable bring-up 247.79 + batch/restore
166.36), ~66 s inside the ceiling. **That is arithmetic across two runs, not a measurement** — this iter
replaces it with a p50 over three composed cycles, **with the spread published beside it** (`C2`).

## Phase plan

- **Phase A — re-pin + verify.** `.agentspace/rext.tag` → `fast-build-m258-iter-06`; fetch + checkout in
  `stack-demo/rosetta-extensions`; verify the four files and the hook are now present. Verify the
  user's stacks resident.
- **Phase B — the campaign.** `buildbench run 1 --reps 3 --profile macmini --no-public-host` from the
  consumption clone. Foreground-poll with a heartbeat; never background-and-yield.
- **Phase C — read the result against the gate's FIVE clauses**, and check **relationships between
  outputs**, not just each output's own success (`D20`). Specifically: does `batch-gate.json`'s
  `batch+restore` reconcile with the `batch_gate` phase anchor? Do the phase sub-totals sum to
  `total_s`? Does the cockpit roster post-condition hold on the stack the campaign leaves behind?
- **Phase D — close**: publish p50 **and spread** on the composed number and both halves.

## The gate's five clauses (graded explicitly in Phase C)

1. one cold command brings the stack up **AND** drives the full batch to completion
2. **zero standing red** (consolidated red set empty)
3. composed **p50 ≤ 480 s** over **3 consecutive cold reset-to-seed cycles**, reproducible
4. **0 platform-repo edits**
5. the stack left in a **presenter-usable world**

## Escalation conditions

- **Composed p50 > 600 s** → the milestone's declared `re_scope_trigger`. Note it reads *"after 3
  tiks"*; this is the first composed p50 ever taken, so a single over-ceiling p50 is surfaced **with
  measurements** but is not by itself the trigger.
- **HEADROOM refusal** → a **result**, not a failure. The host is permanently contended and cannot be
  freed. If the gate instrument refuses, fall back to driving the cycle as an *operator* (iter-04's
  precedent) for the halves contention cannot corrupt, and report the refusal as the reading.
- **Non-empty red set** → `D-v28-3` escalation to the user at batch end, once, with the full set.

## Acceptable close-no-lift outcomes

A campaign that completes but grades a rep unusable for a **named, evidenced instrument reason** (as
iter-06's `postgres-schemas` refusal did) is a complete iter: the finding is the deliverable. What is
**not** acceptable is reporting a p50 over reps that are not gate-usable.
