---
iter: 269
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-10
active_strategy: TOK-08
route: FIX-M257x-262-demo-env-append-is-not-idempotent
---

# iter-269 — the demo `.env` is one block appended 31 times, and last-wins is a hazard

**Type:** tik, under `TOK-08`.

## Step 0 — re-survey (mandatory, before targeting)

iter-262 found `stack-demo/platform/.env` is *"the same 18-key block appended 31 times with
`DIRECTUS_TOKEN` BLANK in all 31"*, called it harmless to a demo, and routed
`FIX-M257x-262-demo-env-append-is-not-idempotent` with the instruction *"find the writer and make it
replace-or-skip."*

Re-surveyed at open, corpus `5c2de87`: **470 lines**, `GH_PAT` × **31**, and five more keys at exactly 31.
The condition is live and has not been touched since.

## Cluster / target identified

`FIX-M257x-262-demo-env-append-is-not-idempotent`. Two things were left unmeasured and both matter more
than the tidiness: **who writes it**, and **whether last-wins makes duplication a correctness hazard
rather than a cosmetic one**.

## Hypothesis

Compose reads `.env` **last-wins**, so N identical blocks are harmless *only while the blocks are
identical*. The moment one writer appends a **blank or stale** value after a good one — which is exactly
the shape `DIRECTUS_TOKEN`-blank-in-all-31 already has — the append order silently decides the value, and
the symptom is the classic *stack boots, catalog empty*.

## Expected lift

Names the writer and converts a "harmless duplication" note into a stated hazard with its mechanism —
or falsifies the hazard, which is equally publishable.

## Phase plan

1. Seal pre-registrations (first commit).
2. Measure: are the 31 blocks byte-identical? Which keys vary?
3. Find the writer(s) that append.
4. Establish the last-wins semantics and whether any key's value differs across blocks.
5. Repair the corpus side; route the tooling fix.

## Escalation conditions

- **Do not rewrite `stack-demo/platform/.env`.** It is a live stack's environment and the evidence; a
  "helpful" dedupe is a mutation nobody asked for and would destroy the measurement.
- The writer fix needs a tag + pin bump → route it, do not spend the frozen-pin control.

## Acceptable close-no-lift outcomes

A documented falsification of the hazard — every block byte-identical **and** no writer capable of
appending a differing value — closes the route as cosmetic, with evidence, which is a real result.
