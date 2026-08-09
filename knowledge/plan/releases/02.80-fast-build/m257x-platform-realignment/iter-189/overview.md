---
iter: 189
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-09
controlling_strategy: TOK-08
---

# iter-189 — one population, two readers, two different exclusion rules — and the repo already wrote down the one they violate

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

## Step 0 — re-survey before targeting

iter-188 routed `SURVEY-M257x-iter188-the-other-walks-are-unmeasured` **with its number**: the same grep
that found `_SKIP_DIRS` found **two more** name-based prune rules in `stack-core`, and the larger of them
is `platform_predicate_guard`'s pair at `:752` and `:806`. Re-surveyed and confirmed still live; targeted
directly rather than re-derived.

## Cluster / target identified

`app_rpc_reads()` — the reader behind **G6**, the guard whose output is the evidence for the corpus's
*"`STORAGE_RPC_ADDR` occurs in zero Go source"* and *"zero `*_RPC_ADDR` variables anywhere"* — has two
implementations, and `_reads_at_ref`'s own docstring calls itself *"**Same derivation**, read from the
object store instead of the checkout."* They exclude differently:

| reader | rule | shape |
|---|---|---|
| `_reads_worktree` `:752` | `parts = set(path.parts)`; `"vendor" in parts` | **component-exact** |
| `_reads_at_ref` `:806` | `"vendor/" in rel` | **substring** |

`"cloud-vendor/x.go"` contains `"vendor/"`. So `_reads_at_ref` **over-excludes** every path with a
directory whose name merely *ends* in `vendor` or `node_modules`, and `_reads_worktree` does not. On the
consumer side of this guard, over-exclusion means a read that exists is reported as absent — the
direction `app_rpc_reads`'s own docstring is most careful about (*"None is not zero"*).

**And the rule they violate is already written down, in this repo, by a sibling guard.**
`story_org_count_guard.py:125` — *"an exclusion that can swallow the whole repo. **Match components;
never substrings.**"* One guard states the rule; another breaks it 60 lines from a docstring claiming the
two are the same derivation.

## Hypothesis

Extracting **one** component-exact predicate that both readers call, and adding a **differential** arm
that runs both against a synthetic git tree, converts a latent silent divergence into a fenced identity —
and gives the milestone the comparison iter-175's rule demands but nothing has ever run for this pair.

## Expected lift

No `P`/`N` reading (`§9`). One shared predicate replacing two; the substring form deleted; ≥6 arms
including a differential with its own instrument control; both hazards sized rather than asserted.

## Phase plan

- **A** — size both hazards and run the comparison that has never been run.
- **B** — extract the shared component-exact predicate; delete the substring form.
- **C** — fence: both readers call it · it is component-exact · the two readers **agree** on a synthetic
  tree · the differential can detect a divergence (`§9`).
- **D** — mutation-prove; both runners.
- **E** — publish the `§8` rule; route residuals.

## Escalation conditions

- If the two readers disagree **today** on the real `app` clone, this is a live defect in a published
  corpus claim, not a latent one — stop and grade it as such before repairing.
- If the substring form turns out to be deliberate (some path shape the component match would miss),
  keep it, name the shape, and fence the difference instead of removing it.

## Acceptable close-no-lift outcomes

- The two readers are already reconciled somewhere upstream (a caller normalises paths) → the finding
  shrinks to a duplicated-rule cleanup; record the falsification and close.
