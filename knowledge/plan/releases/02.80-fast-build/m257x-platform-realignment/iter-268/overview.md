---
iter: 268
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
route: ROUTE-M257x-265-stack-demo-carries-six-dead-clones
---

# iter-268 — "no longer treated as part of the project", measured on disk

**Type:** tik, under `TOK-08`.

## Step 0 — re-survey (mandatory, before targeting)

The user's binding closing condition has a second half this milestone has answered only in prose: *"…with
the **deprecated/removed repos no longer treated as part of the project**."* Every iter so far has read
that as a **corpus** statement. iter-265 noticed it is also a **disk** statement — `stack-demo/` carries
clones of **`cms`, `graphql-wundergraph`, `jobsimulation`, `messenger`, `roadrunner`, `storage`**: six
repos that left `repos.yml`.

Re-surveyed at open, corpus `c33d2c8`: the six directories are present. Nothing else is known — whether a
bring-up still fetches them, whether anything still builds from them, and whether they are fossils or
fresh has not been measured.

## Cluster / target identified

`ROUTE-M257x-265-stack-demo-carries-six-dead-clones`, promoted from a side note because it sits directly
on the user's own words. It also decides a **reach** question iter-265 left open: a `cd cms` fence resolves
on this box *because of these clones* and would not on a fresh one (§8, iter-241).

## Hypothesis

They are fossils — left by an older `repos.yml`, fetched by nothing today, built by nothing. If so the
finding is *"the workspace keeps a corpse the tooling has stopped feeding"*, which is untidy but honest.
**If instead a current bring-up still clones or builds any of them, the user's condition is not met on
disk**, and that is a materially larger finding than a stale directory.

## Expected lift

The second half of limb 3, measured rather than asserted.

## Phase plan

1. Seal pre-registrations (first commit).
2. Enumerate: which repos in each `stack-*/` are absent from `repos.yml`.
3. Determine whether the demo bootstrap (`ensure-clones.sh`) or any compose/build path still names them.
4. Date them against the `repos.yml` commit that removed them — fossil or fresh?
5. Repair the corpus side; route what needs a pin bump.

## Escalation conditions

- **Delete nothing.** A stale clone is evidence; removing it destroys the measurement and is a
  working-tree mutation nobody asked for. This iter reports, it does not tidy.
- If a current bring-up still fetches them, that is a tooling defect needing a tag + pin bump — route it,
  do not spend the frozen-pin control.

## Acceptable close-no-lift outcomes

A documented falsification of PR-2 (a bring-up *does* still clone them) inverts the hypothesis and is the
more valuable outcome; it closes as a finding, not as a fix.
