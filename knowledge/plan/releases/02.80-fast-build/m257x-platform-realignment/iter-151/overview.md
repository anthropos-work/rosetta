---
iter: 151
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-08
---

# iter-151 — the last open arm of the absent-value class

**Active strategy reference:** `TOK-08` — census the mechanical classes; stop sampling them.

## Cluster / target identified

`SURVEY-M257x-iter147-absent-value-class`, the one route iter-148 left **partially** closed. iter-147
censused the seven compose-profile choice points and found two that brought up empty stacks because a
token census cannot see an ABSENT value; iter-148 closed the `STACK_SERVICES` arm and graded
`rosetta-demo`'s `--ref`/`--only`/`--services` sound. It left this: **`STACK_PROJECT` and `STACK_OFFSET`,
whose unset defaults silently target the main dev stack.**

**Step-0 re-survey:** untouched; iters 149–150 worked other classes.

## Hypothesis

The severity of an absent-value default depends entirely on **which side reads it**. Unset,
`STACK_PROJECT` resolves to `anthropos` — the developer's own main dev stack — and `STACK_OFFSET` derives
from it. On a read-only probe that is a wrong answer, loud and recoverable. On any write path it is a
seeder, a snapshot replay or an injection writing into the wrong stack because a variable was missing
from an environment. **So the census to run is not "where is it read" but "is any READER a WRITER".**

## Expected lift

No `N` reading; no `N` movement claimed (`§9` guard-rail 1). The deliverable is the graded read set and a
fence that holds whatever it says.

## Phase plan

1. Census every non-comment, non-prose read of both variables, repo-wide.
2. Partition the read sites into probe-side and write-side sections.
3. Fence the partition — with anti-vacuity on the subject and a RED-proof, because a census that returns
   zero cannot demonstrate its own instrument (`D-M257x-149-3`).

## Acceptable close-no-lift outcomes

If every read is probe-side, the route closes with a falsification and a fence that keeps it true. That
is the expected outcome and it is a complete iter.
