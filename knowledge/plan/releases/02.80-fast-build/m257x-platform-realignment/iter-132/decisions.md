# iter-132 — decisions

## `D-M257x-132-1` — the hedge is settled by a `git clone`, not by better wording

The routed repair was framed as a **prose** job: eleven sites publish *"cms's prod state is
UNMEASURABLE because `infrastructure` has never been in any clone set"*, so replace the sentences with
the settled verdict iter-123 measured.

Seven **more** sites hedged a *different* proposition on the *same* premise — the production RPC
address `http://backend.internal.anthropos:8081`, which the corpus says *"no `.tf` file in any clone
names"* and then declines to assert in either direction, because the deciding declaration lives in
`infrastructure`.

**Decision: measure it instead of re-wording it.** A `--depth 1` clone completed in under a minute and
its `HEAD` is `13c248e6` — the exact sha the corpus already cites 28 times. **Production DOES name the
literal, exactly once**, and the corpus could have said so at any point in the last four iterations.

**Why this is a decision and not just work:** the alternative was available and tempting — write the
honest hedge (*"not in the standing clone set, read transiently for a different question, so not
re-derivable here"*), which is TRUE, defensible, and would have closed the route. It would also have
**re-published a limit that does not exist** in seven more places. `D-M257x-129-4`'s rule — *a
mirror-write instruction is retracted, not re-qualified* — generalises: **when the premise of a hedge
is a habit rather than a barrier, retract the hedge; do not soften it.** → `platform-alignment.md` §5
**rule 61**.

## `D-M257x-132-2` — the fence's NOTE was counting my repair as the disease, and the fix is a third bucket, not better prose

`unreadable_repo_claim_guard`'s NOTE read **"11 site(s) still hedge about `infrastructure` while 14
report having read it"** — up from 9/13 at the iter's open, i.e. **the repair appeared to make the
corpus hedge more.**

Measured rather than assumed: of those 11 paragraphs, **8 also carry a ref-pinned reading.** They quote
the retired wording *in order to retract it* — `architecture_overview.md:227` and
`platform-migration-status.md:93` are the corpus's own model retractions, and both were counted as live
hedges. The guard classifies by substring and tests `marker` **before** `measured`, so a retraction can
never reach the `measured` bucket.

**Two fixes were available and only one of them is honest.**

- **Rejected:** re-word the retractions until the marker substring is gone. That is tuning the corpus to
  the instrument — the number improves and nothing about the corpus does. This milestone has refused
  that class since `D-M257x-122-3`.
- **Taken:** a **third bucket**, `mixed`, reported separately, with a `KNOWN_WEAKNESS` line that says in
  the instrument's own voice that it **cannot** distinguish a quoted retraction from a live hedge. The
  NOTE now fires on live hedges only (**1**, and it is the protocol doc's own historical worked
  example). Exactly `D-M257x-121-4`'s ruling — *do not widen; DISCLOSE the floor in the instrument* —
  reached independently by a second fence.

**Controls, and the honest accounting:** three tests were added; a meta-mutation that deletes the
`mixed` bucket **kills 2 of the 3**. The third survives its own mutant, so it is **renamed** from
`test_MUTATION_…` to `test_FIXTURE_INTEGRITY_…` with the survival stated in its docstring. A control
that cannot fire is the defect class this milestone has caught eight times; naming one MUTATION while it
survives is the same defect wearing a better label.

## `D-M257x-132-3` — the fence caught two unref-pinned assertions in the repair itself, and they were real

Two of the three live hedges at first re-run were **paragraphs iter-132 had just written**: the new
riders in `backend.md` assert that `module.messenger_euwest1` is deleted and that `infrastructure`'s own
narrative docs describe it anyway — **with no sha in that paragraph.**

**Decision: fix the prose, not the fence.** The claim is about a `module.*_euwest1` construct, the fence
asks such a claim to name its ref, and the fence is right. `13c248e6` now appears in the riders.

**Recorded because of what it says about sweeps:** a repair that retires a hedge carries the ref into
*the paragraph where the reading is narrated* and then, characteristically, forgets to carry it into the
paragraphs it spawns. This is the sixth time on this milestone that a fence has fired on the repair
rather than the defect, and the sixth time it was correct to.

## `D-M257x-132-4` — the routed item named a file that was already correct, and that is reported rather than dropped

`FIX-M257x-iter131-infrastructure-hedge-stale` was routed as *"11 sites **+ `CLAUDE.md`**"*, with the
aggravating note that *"`CLAUDE.md` publishes it too, so every agent that loads this repo starts from
the retracted claim."*

**`CLAUDE.md` does not publish it.** It was corrected at iter-124 and carries the settled reading twice
(`:194-203`, `:259`). Its three remaining `clone set` mentions are `stack-demo`'s clone set (`:123`,
`:132`) and `customerio-sync`'s `go.mod` (`:293`) — a different repo, and a claim that still holds.

**Decision: report the over-statement in the iter that inherited it.** The route was written by the
same coordinator whose adjudication `adj-1` had just corrected, in a reading that disclosed the
independence deviation; an inherited route is evidence, not instruction, and `D-M257x-121-1`'s rule
(*re-derive a report at source before filing it*) applies to routes we wrote ourselves.
