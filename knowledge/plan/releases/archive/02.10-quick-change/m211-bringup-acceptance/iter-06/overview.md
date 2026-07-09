---
iter: 06
milestone: M211
iteration_type: tik
iter_shape: cleanup
status: closed-fixed
created: 2026-07-08
---

# iter-06 — tik (cleanup/measurement): establish sub-condition (f) — 0 residual skiller-schema refs

**Type:** tik (cleanup/measurement) — under **TOK-01**. 0 production runtime code modified (the attempted
tidy was investigated + reverted; see below).

## Step 0 — Re-survey
iter-05 noted 2 skiller repo-test-cmd branches in `stack-verify/repos/run.sh`. The open question: is
sub-condition (f) ("0 residual skiller-SCHEMA references in any path the tooling queries") MET, and are
those branches a gate violation?

## Active strategy reference
**TOK-01** — the residual-skiller closure of the warm inner loop.

## Cluster / target
Sub-condition **(f)**: formally determine whether 0 residual skiller-SCHEMA refs remain in the paths the
tooling queries, and resolve the `repos/run.sh` residue.

## Hypothesis
The M209/M210 re-ground left 0 `skiller.<table>` schema refs in queried SQL; the remaining "skiller"
strings are concept-name comments + test fixtures + dead config branches (not schema refs).

## Phase plan (executed)
1. Comprehensive sweep of all rext tooling for `skiller.<table>` schema refs → confirm 0.
2. Classify the remaining "skiller" strings.
3. Assess/attempt the `repos/run.sh` dead-branch tidy; test-gate it.

## Outcome
**Target MET — (f) confirmed.** 0 `skiller.<table>` schema refs in rext production SQL
(stack-snapshot/stack-seeding/stack-verify). The remaining "skiller" strings are: (a) concept-name comments
in seeders (naming the taxonomy source — not schema refs), (b) intentional test fixtures/tokens in
`test_verify.py` (scope-filter exact-match + $DEVDIR path-resolution regression), (c) 2 dead `case skiller)`
branches in `repos/run.sh` — skiller is ABSENT from repos.yml so the branches are never iterated (dead
code), and they are NOT schema refs. **Attempted** removing the dead branches; it broke
`test_absent_repo_resolves_under_stack_root` (which uses "skiller" as a generic absent-repo fixture and
depends on the branch to reach path-resolution). Since (f) is about SCHEMA refs (0, met) and the branches
are harmless dead-but-fixture-referenced code, the removal was **reverted** (don't churn multiple test
fixtures for zero gate value). rext tree left clean; verify suite 104/104 green after revert.
