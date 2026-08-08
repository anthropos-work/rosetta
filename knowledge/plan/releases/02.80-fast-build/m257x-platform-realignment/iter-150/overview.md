---
iter: 150
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-08
---

# iter-150 — which declared lists claim a derivation they never perform?

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Cluster / target identified

`SURVEY-M257x-iter149-declared-lists-unfenced-against-layout`, opened by iter-149 hours ago and the
newest evidence in the queue. iter-149 found `claim_census_guard.REXT_SECTION_NAMES` describing itself as
*"derived from the monorepo's own layout"* while being a hand-typed tuple that had **drifted** — 10 names
for 11 sections, so every claim naming the section behind `/stack-secrets` resolved to no known artifact
and left the census silently. It closed with the obvious question unasked: **how many other declared
tuples make that claim?**

This is a mechanical class in `TOK-08`'s exact sense — the subject (a module-level literal) carries its
own comment, so the population is enumerable by parse rather than by reading.

**Step-0 re-survey:** the target is one iter old and untouched; nothing has absorbed it.

## Hypothesis

A constant whose comment says *derived* and whose value is a literal is a **claim with no mechanism**. The
population is enumerable: parse every `.py` in the monorepo, take module-level `UPPER = <literal>`
assignments, and read the contiguous comment block immediately above each for a self-directed derivation
claim. Then grade each by hand — because the word appears for at least three other reasons (something
*else* derives FROM this constant; the constant is a fixture for a derivation under test; the comment
discusses derivation as a design choice) and a token-level count would report all of them.

## Expected lift

No `N` reading is planned, so **no `N` movement will be claimed** (`§9` guard-rail 1). The deliverable is
`TOK-08`'s per-class triple: population enumerated · claims graded · defects fenced.

## Phase plan

1. Parse-level census over every Python file in the monorepo; report constants scanned, comment blocks
   claiming a derivation, and — after hand-grading — the self-claiming subset.
2. Grade each self-claim: **fenced** (a named test compares it to its source), **correct-but-loose**
   (the comment is about a neighbouring value), or **unfenced claim** (the defect).
3. Fence the unfenced ones at the point of use, per `§2`'s derive-at-the-point-of-use doctrine, or —
   where the module genuinely cannot derive — fence the property the comment asserts, which is the
   `D-M257x-149-5` shape.
4. RED-proof each fence against content that should trip it.

## Escalation conditions

- If a defect's repair would change a guard's verdict on the live corpus → route forward rather than
  land a fence and a verdict shift in the same iter.

## Acceptable close-no-lift outcomes

A census showing every self-claim is already fenced is a complete iter — iter-149's route would close
with a falsification, which is the outcome its own lesson 2 predicts against.
