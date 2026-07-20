---
iter: 5
milestone: M236
iteration_type: tik
status: closed-fixed
date: 2026-07-20
metric_pre: "16/31"
metric_post: "27/31"
---

# iter-05 — tik: the MANAGER vantage

## Step 0 — re-survey

iter-04's reading stands: 16/31, with **13 of the 15 non-landing pairs on the manager vantage**. The
TOK-01 Phase-L direction ("one arm per iter, descending evidence density") names this as the next target
and it is untouched. No substitution.

## Active strategy reference

**TOK-01 "publish-then-prove", Phase L** — the manager arm.

## Cluster / target identified

One symptom cluster, 11 pairs: the manager scoreboard renders its header but shows **"No data"** for the
attempts table and **"undefined undefined"** for the player. Two further interview pairs fail *differently*
(no header at all), suggesting a second cause.

## Hypothesis

The manager page's failure is a **read-path** defect, not a seeding gap — iter-04 established that the
mirror row exists (13/13) and the user row is correct (`Clara Romano`), and the *player's own* page renders
her name from the same data. If one query is erroring, both symptoms fall out of it at once.

## Expected lift

**+11 to +13 pairs** if the single-cause hypothesis holds.

## Phase plan

1. Determine whether the defect is content-story-specific or demo-wide (probe a hero's drill-down).
2. Capture the network/GraphQL layer to name the failing query.
3. Read the resolver; confirm the contract it actually expects.
4. Prove the fix-shape live BEFORE writing code.
5. Fix in tooling (zero platform edits), re-export, re-measure.

## Escalation conditions

- If the fix requires a **platform-repo edit** → hard stop; escalate as a user-blocker.
- If the cause is in the platform's own code with no seam → demopatch is the sanctioned route; assess cost
  before committing.

## Acceptable close-no-lift outcomes

A precisely-named root cause with a proven-but-unlandable fix would still be a complete iter.
