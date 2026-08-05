---
iter: 78
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-05
---

# iter-78 — settle the denominator iter-76 left unsettled, and fence it

**Active strategy:** `TOK-05` — *stop repairing claims; fence the predicates under them.*
`D-M257x-59-1` (repair by predicate), `D-M257x-77-1` (fence only what is decidable without reading a
sentence).

## Step 0 — re-survey (mandatory)

Re-run at open, from the corpus root against `--platform stack-demo/platform` (`0dab54d`):

```
6 repo(s), 10 compose service(s) across docker-compose.yml+common.yml, floor 3,
8 legal profile(s), default `core` selects 5, migrating ['app']
platform_predicate_guard: OK
```

iter-77's targets are absorbed and its fence is green. `CHECK-M257x-iter76-compose-service-count` is
**live and unabsorbed** — iter-76 recorded it as *"8 vs 9 vs 10; my one-line grep disagreed with the
tested parser and the disagreement is recorded rather than resolved by assertion."* It is the one
routed item in this milestone explicitly left **unsettled**, and it is a denominator, which under
TOK-05 makes it a predicate rather than a claim.

## Cluster / target identified

`CHECK-M257x-iter76-compose-service-count`. Chosen over the other iter-77 routes because it is a
**denominator the guard already derives and prints on every run** while the corpus states a third
number in three places — the exact shape `D-M257x-59-1` calls a predicate, and the cheapest
remaining one.

## Hypothesis

That all three numbers are **right about different things**, and that the disagreement is a missing
qualifier rather than a wrong measurement: **8** is `docker-compose.yml` alone, **10** is the
effective set once `include: common.yml` is resolved — and **9** is a count of nothing at any ref.

If so the repair is not "pick the right number"; it is to make each site say *which set it counts*,
and the fence must assert the pair, never a single value.

## Expected lift

- The unsettled denominator settled **by derivation across refs**, not by assertion.
- Each of the three `nine services` sites adjudicated and repaired with its qualifier.
- A fence for the predicate **only if it clears precision on the live corpus** — measured first.

## Phase plan

A derive both counts + the ref history · B adjudicate every "N services" site · C repair by
predicate · D fence, or report the measured negative · E re-measure + regression tests.

## Escalation conditions

- A fence that cannot reach FP-free precision on the live corpus is **reported as a measured
  negative and NOT shipped** (§4 Trap A) — the iter-77 precedent, applied without hesitation.
- More than ~10 adjudicated repair sites → repair the predicate's live set and route the remainder.

## Acceptable close-no-lift outcomes

Settling the denominator with a derivation and repairing the sites is the deliverable; a fence is a
bonus, and reporting that no sound fence exists for this construct is a complete iter.
