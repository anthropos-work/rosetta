---
iter: 60
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-04
active_strategy: TOK-05
refs:
  platform: 0dab54d          # local clone == origin, re-verified at open
  app: v1.366.0
  rext_pin_at_open: fast-build-m257x-iter-58
  rext_pin_at_close: fast-build-m257x-iter-60
---

# iter-60 — build the sibling predicate guard, and let G1 close the profile class

**Active strategy:** [`TOK-05: stop repairing claims; fence the predicates under them`](../decisions.md).
First step of `D-M257x-59-5`'s ordering — **fence → citations → map state → read**. This is *fence*.

## Step 0 — re-survey before targeting (mandatory)

TOK-05 named iter-60's target one day ago and pre-registered it as refutable. Re-measured at open,
against the artifacts rather than against the tok:

| pre-registered | measured at open | verdict |
|---|---|---|
| platform at `0dab54d`, level with origin | `git rev-parse HEAD` == `git rev-parse origin/main` == `0dab54d` | **holds** |
| 6 denominators (10 svc / 6 repos / core→5 / 4 RPC / 1 migrating / mid-fold) | all six reproduced, and cross-checked against `docker compose --profile X config --services` | **holds** |
| G1 goes RED naming `graphql`, and `cms` + `storage` too | RED on **five** tokens — `graphql`, `cms`, `storage`, plus `jobsimulation` and `roadrunner` | **holds, wider** |
| "17 files / 30 occurrences" | **26 live docs / 56 lines** by whole-tree grep; the fence's own parsed-construct reach is **53 sites over 11 tokens** | **corrected** |
| `cmd/academyImport/main.go:235` hard-requires `STORAGE_RPC_ADDR` | the `Getenv` is `:231`; `:235` is the `is required` return. Both true, different lines | **corrected** |

Target stands, unsubstituted.

## Hypothesis

Three residual classes rest on predicates whose legal sets are derivable from `repos.yml` +
`docker-compose.yml` (include-resolved) + `Makefile`. A guard that derives those sets closes the class
in one build rather than in 56 edits — and, unlike a repair, keeps it closed.

## Expected lift

- A new sibling guard with 6 assertions, each watched **RED then GREEN**, with a surviving no-op
  positive control and **inversion** mutants (§8 rule 5 — removal mutants do not catch inversion).
- The `graphql`-profile predicate class repaired to **zero fence findings**.
- The three protocol-doc lessons TOK-05 deliberately withheld, written into §5 and §6.

## Phase plan

A. Re-derive all six denominators against the platform clone (done above).
B. Build `stack-core/platform_predicate_guard.py`; watch RED; fix RULES not thresholds; watch GREEN.
C. Repair the predicate class the fence names.
D. Write §5 rules and the §6 lesson.
E. Re-measure: fence GREEN, suite at baseline, both trees clean.

## Escalation conditions

- A fence draft whose false-positive rate cannot be removed by **replacing the rule** → route the
  assertion forward rather than ship a fence that gets disabled on first contact (§4 Trap A).
- A repair that would require a platform-repo edit → stop; the release forbids it.

## Acceptable close-no-lift outcomes

If the six denominators had failed to reproduce at `0dab54d`, the iter closes on the falsification —
TOK-05's whole premise is that the legal sets are derivable *today*, and a denominator that will not
re-derive refutes it.
