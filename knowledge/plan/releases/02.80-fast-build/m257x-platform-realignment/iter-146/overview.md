---
iter: 146
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-08
---

# iter-146 — audit iter-13's own re-point for completeness

## Step 0 — re-survey before targeting

iter-145 proved that **M257x iter-13's re-point off the deleted GraphQL router was incomplete** — it
updated `stack-verify/lib/services.sh` and left the test side's copy, RED for 132 iters. The re-survey
asks the obvious next question, which iter-145 deliberately routed rather than answered:

> **iter-13 missed one place. Did it miss others?**

That is a mechanical, censusable question over a finite population — `TOK-08`'s shape exactly — and it
audits a repair *this milestone made*, which is where iter-145's evidence says the risk is.

## Active strategy reference

`TOK-08` — census the mechanical classes; stop sampling them. The class here is *"tooling that still
names the deleted router as a reachable endpoint."* A port either has a listener or it does not.

## Cluster / target identified

Every reference to the deleted Cosmo/WunderGraph router in `rosetta-extensions`: the port `5050`,
the product names, and the `graphql` compose-service name. Denominator to be stated.

## Hypothesis

The misses are **not randomly distributed**. iter-145's miss landed in the one test section nothing
runs. If the pattern holds, the remaining misses sit on paths that are **not exercised** — rarely-taken
branches and build-arg defaults — and the executed paths will be clean.

## Expected lift

No `N` reading. The deliverable is the census with its denominator, each reference classified
(live-defect · latent · correct-repoint · fence-asserting-absence · inert prose), and every live
defect repaired with a fence.

## Phase plan

- **A** — census every `5050` / router-name reference in rext; classify each by file role.
- **B** — verify each candidate live defect against the mechanism (does a listener exist? is the front
  applied? is the default overridden?), not against its wording.
- **C** — repair what is live; fence it.
- **D** — re-run every section the repair touches (rule 68's own instruction), close.

## Escalation conditions

- A live defect on a **write path** the gate's clause 4 covers → surface rather than silently repair.
- A repair that would need a platform-repo edit → stop (the v2.8 zero-platform-edit constraint).

## Acceptable close-no-lift outcomes

If the census finds the executed paths clean and no live defect at all, that is a complete iter: it
would establish that iter-145's miss was isolated to the never-run section, which is a *stronger*
version of rule 68 than the one currently written.
