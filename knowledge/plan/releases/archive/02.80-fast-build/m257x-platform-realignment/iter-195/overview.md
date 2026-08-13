---
iteration_type: tik
status: closed-fixed
controlling_strategy: TOK-08
date: 2026-08-09
---

# iter-195 — the six sections nobody had ever run

**Type:** tik · **Active strategy:** `TOK-08` · **Protocol:** `corpus/ops/platform-alignment.md`

## Step 0 — Re-survey before targeting

`SURVEY-M257x-iter186-264-go-tests-have-never-been-read` is the milestone's largest standing
measurement gap and has been open for nine iters. The brief states its consequence plainly: *every
"whole-population" figure published before iter-186 describes 5 of 11 sections and one language.*

Re-survey confirmed it is untouched and that nothing in iters 187–194 read a Go test.

## Cluster / target identified

The six Go sections: `alignment`, `clerkenstein`, `playthroughs`, `stack-secrets`, `stack-seeding`,
`stack-snapshot`.

## Hypothesis

They are **unread by this runner**, not unreadable — and the difference has never been tested.

## Expected lift

Either a first reading of the Go population, or a **measured** statement of why it cannot be taken
here. Both close the route; only one of them is good news.

## Phase plan

1. Time-boxed feasibility probe on the smallest section.
2. If runnable: run all six, tally, and **name the unit**.
3. Make it repeatable — derive the Go section set, wire a runner, fence it.

## Escalation conditions

If the Go toolchain needs network or org credentials, record that as the measured answer and stop.

## Acceptable close-no-lift outcomes

A measured *cannot-run-here*, with the reason, is a complete iter — the route is closed either way.
