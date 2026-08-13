---
iter: 62
milestone: M257x
iteration_type: tik
status: closed
opened: 2026-08-04
active_strategy: TOK-05
refs:
  platform: 0dab54d
  app: v1.366.0
  rext_pin: fast-build-m257x-iter-61
---

# iter-62 — repair `FIX-M257x-iter61-profile-prose-class`, whole

**Active strategy:** [`TOK-05`](../decisions.md). iter-61 landed the widened fence and left it **RED
at 2 findings / 22 sites / 12 files**, enumerated in
[`../iter-61/evidence/residual.md`](../iter-61/evidence/residual.md) and routed **whole** rather than
in subsets (§5 rule 19's scope-edge corollary).

## Step 0 — re-survey

Re-ran the guard at open: the residual reproduced exactly — 22 sites, 12 files, at platform `0dab54d`
(re-verified level with origin). No substitution.

## Hypothesis

The class is one predicate with one legal set, so it closes in one pass. What it will *not* close in
one pass is the **second** predicate riding along in the same rows — several sites assert not just a
`graphql` profile but *"the husk container still starts"*, which is separately false at `0dab54d`.

## Expected lift

Guard GREEN; `markdown_structure_guard` still clean.
