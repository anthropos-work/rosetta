---
iter: 79
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-05
---

# iter-79 — size the cross-repo-pin class, and apply the rule to the whole family

**Active strategy:** `TOK-05` — *stop repairing claims; fence the predicates under them.*
`D-M257x-77-1` (fence what is decidable without reading a sentence), `D-M257x-78-3` (the class is
live).

## Step 0 — re-survey (mandatory)

Re-run at open, from the corpus root against `--platform stack-demo/platform` (`0dab54d`):

```
6 repo(s), 10 compose service(s), 8 legal profile(s), default `core` selects 5, migrating ['app']
G10 4 compose-service-count claim(s); G9 4/19 repos.yml citation(s) graded
platform_predicate_guard: OK
```

iter-78's target is absorbed and `CHECK-M257x-iter76-compose-service-count` is closed.
`CHECK-M257x-iter77-cross-repo-pin` is **live, and was upgraded by iter-78** from *"unmeasured
whether any of the 145 dates a platform claim"* to *"at least two do, and both were false."*

## Cluster / target identified

`CHECK-M257x-iter77-cross-repo-pin`. Two iterations found the same mechanism independently —
`roadrunner.md:14` (G9) and `external_services.md:296` (G10) — and each fixed it **locally, inside
the assertion that happened to trip over it**. G2, G4 and G5 assert claims about platform files and
still take any sha in the block as a date. **A rule owned twice and implemented twice is a rule not
implemented.**

## Hypothesis

That the routed **145** is an upper bound of the same kind every routed count in this milestone has
been, and that the true class — pin-exempted blocks where a **foreign sha dates a platform-file
claim** — is **small**, because most of the 145 are legitimate `app`-repo citations about `app`
files. Prediction, registered before building: **fewer than 10**, and **0 new findings** once the
rule is generalised, since iter-77 and iter-78 already repaired the two that were false.

## Expected lift

- The class **sized by measurement**, closing or bounding the route rather than carrying it.
- The resolve-in-repo rule applied **once, in the shared helper**, instead of a third local copy.
- Watched: reach up, findings unchanged — and if findings move, each one adjudicated by reading.

## Phase plan

A size the class · B implement the rule in the shared helper, derived not listed · C measure reach
and findings, adjudicate anything new by reading it · D positive controls + regression tests ·
E re-measure and route the residual honestly.

## Escalation conditions

- Any new finding that reads FALSE on inspection → the rule is reported as a measured negative and
  NOT shipped, exactly as iter-77 did with the all-pins widening.
- More than ~10 newly-graded claims turning RED → measure and route; do not repair inside this iter.

## Acceptable close-no-lift outcomes

Sizing the class and reporting that no generalisation is safe would be a complete iter. So would
shipping the rule with 0 findings: **reach is the deliverable here, not a repair count.**
