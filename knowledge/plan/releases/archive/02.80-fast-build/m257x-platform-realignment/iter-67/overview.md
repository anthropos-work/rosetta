---
iter: 67
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-04
---

# iter-67 — G7: the service list beside a profile

**Active strategy reference:** `TOK-05`. Consumes `FENCE-M257x-iter66-tier-membership`, opened by
iter-66 one iteration earlier — the routed item and its build in adjacent iters, so the rule is
proven by use rather than by assertion (the §7 rule-4 discipline).

## Cluster / target identified

iter-66 corrected a prose sentence placing `storage` in the default selection and observed that
**no fence could have caught it**: G1 checks that a token is legal and selects *something*, G3 checks
the default's *count*. Nothing checks the **list**. A profile-reference table
(`| Profile | Services started … |`) states that list in a construct that IS mechanically decidable.

## Hypothesis

`compose.beyond_floor(tok)` is the legal set. One assertion over the parsed compose the guard already
holds closes the class in every table the corpus has.

## Expected lift

G7 landed, watched RED on fixtures and mutants, GREEN live; the reach reported (checked rows vs
prose-cell rows UNREACHED).

## Phase plan

A. Read the services column by header, like the profile column; read its cells by shape.
B. Assert `stated == beyond_floor(tok)`, both directions (MISSING / NOT STARTED).
C. Mutants; suite; §5 rule 34 re-point.

## Escalation conditions

- If the live corpus goes RED at scale, route the residual WHOLE rather than weakening the rule.

## Acceptable close-no-lift outcomes

- Every live row already correct → recorded, with the fixtures as the evidence the rule can fire.
