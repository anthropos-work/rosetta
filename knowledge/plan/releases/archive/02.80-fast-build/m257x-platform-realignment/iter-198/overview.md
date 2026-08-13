---
iter: 198
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
---

# iter-198 — "six clones are behind" is a fact about the substrate, not about the evidence

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Step 0 — re-survey (mandatory)

Re-ran the route survey at HEAD `592d583` (iter-197's commit). Of the four routes harden passes 45–47
filed, `FIX-M257x-h44-…` was re-sized (not closed) by iter-197 and the conversion is deliberately its own
work. The next one is stated as an explicit either/or, which makes it unusually well-formed:

> `SURVEY-M257x-h46-stale-substrate-direction-undeclared` — *"`claim_census_guard`'s STALE SUBSTRATE
> warning is conservative (false-RED only) and says neither that nor its verdict consequence. **Declare
> the direction, or grade it.**"*

Confirmed live before planning: the guard prints a six-clone staleness table — `cms` 2, `jobsimulation` 4,
`messenger` 7, `next-web-app` 4, `rosetta-extensions` 1, `storage` 20 commits behind — and then prints
`claim-census: OK — ratchet holds` and returns **0**.

## Cluster / target identified

Two mechanical questions the module poses and does not answer:

1. **How much evidence is actually affected?** `stale_clones` answers *which clones are behind*. A reader
   holding a tier-1 excerpt needs *is this excerpt one of the affected ones* — a different grain, two
   steps further on.
2. **What does the staleness do to the exit code?** Undeclared. This is the batch's own class (*an
   instrument that states its own invalidity must not exit 0*) — but whether this instrument is IN that
   class is itself a question, not an assumption.

## Hypothesis

Both are derivable. Exposure = tier-1 pairs whose **cited line range** differs between the clone's HEAD
and its own fetched `origin/main`. Verdict-dependence = whether the inputs the exit code actually reads
are among the drifted files. Neither needs adjudication, so both are inside `F4`.

## Expected lift

No `P`/`N` reading. Deliverable: the exposure number with its three buckets, a derived verdict-scope
statement, and the direction claim graded — including a retraction if it is one-sided.

## Phase plan

1. `drifted_files` + `substrate_exposure` — the file-grain and pair-grain refinements of `stale_clones`.
2. Print them in the STALE SUBSTRATE block, per-pair.
3. Grade the **direction**: is false-RED really the only reachable error?
4. Grade the **verdict consequence** by deriving it from what the exit code reads, not by asserting it.
5. Fence all of it offline (synthetic `git init` clones), so no arm depends on the demo clone set.

## Escalation conditions

- If the verdict IS substrate-dependent today → the guard is stating its own invalidity and exiting 0;
  that is a finding to report, not to fix by flipping an exit code mid-iter.
- Do not attempt to update the clones. They belong to a live demo stack this milestone may not touch —
  `substrate_of`'s own docstring says so.

## Acceptable close-no-lift outcomes

An exposure of **0** would be a first-class result — it would mean the warning is currently inert and the
table over-warns — provided the instrument is shown able to return non-zero.

## Explicitly NOT in scope

Re-cloning or updating the stale clones; adjudicating any of the affected pairs.
