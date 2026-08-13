---
iteration_type: tik
status: closed-fixed
controlling_strategy: TOK-08
date: 2026-08-09
---

# iter-192 — the denominator was arithmetic over a registry, and the fallback never spoke

**Type:** tik · **Active strategy:** `TOK-08` (census the mechanical classes; stop sampling them)

## Step 0 — Re-survey before targeting

iter-191 closed minutes before this iter opened and routed
`SURVEY-M257x-iter191-published-denominators-are-unenumerated` with a **mechanical selector** already
attached: *a printed count whose derivation does not appear in the function that produced the verdict.*
That is the target. Re-survey confirmed it is un-attempted (routed at the close of the immediately prior
iter, no fence exists for it) and that its two founding cases — iter-186's `suite_census.SECTIONS` and
iter-191's `story_org_count_guard` — are both already repaired, so **the class must be proven to fire on
something other than its founding cases or it is not a class.**

No substitution. The TOK-directed target stands.

## Cluster / target identified

TOK-08 says a reading SAMPLES and a fence CENSUSES, and it says to work the classes in descending
measured size **and state the denominator**. So the iter's first move is an enumeration, not a hunt.

## Hypothesis

The printed-denominator class is enumerable over `stack-core` by AST, and running the enumeration will
surface instances that neither of the two by-judgement discoveries reached.

## Expected lift

An enumerated population for the class with its denominator stated, plus ≥1 live defect the two prior
by-judgement findings did not reach. `P`/`N` is **not** re-cut this iter (clause 5 is deliberately not
re-read; see the milestone's standing position).

## Phase plan

1. Enumerate candidate selectors mechanically; **report the flag rate of each**, because a selector that
   flags most of its population is noise, not a fence.
2. Take the decidable sub-class and measure the live instances.
3. Size every defect **before** editing.
4. Repair; prove each repair fires with a mutation control that would have failed before it.

## Escalation conditions

If the class turns out not to be mechanically decidable, say so with the flag rate and route the
residual — do **not** ship a noisy selector as a fence.

## Acceptable close-no-lift outcomes

A measured refutation that the class is decidable, with the numbers, is a complete iter.
