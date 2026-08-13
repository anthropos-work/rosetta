---
iter: 226
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-09
---

# iter-226 — the supersession that never reached the corpus

**Type:** tik, under `TOK-08` (*a reading SAMPLES; a fence CENSUSES*), on the **corpus half** of the user
redirect of 2026-08-09.

## Step 0 — re-survey before targeting (mandatory)

iter-225 surfaced this and did **not** close it. It repaired **one** banner in `build-budget.md` and
routed the rest, recording the general shape as its Lesson 3:

> *"A supersession that reaches the ledger and not the prose has not landed either."*

Re-surveyed now, the target is still untouched and is **small, bounded and mechanical** — which is
exactly what `TOK-08` says to census rather than sample.

## Cluster / target identified — population measured before any repair

| | count |
|---|---|
| `odysseus` mentions in `corpus/` + `CLAUDE.md` + `.claude/` | **27** |
| — in `corpus/ops/demo/build-budget.md` | 26 |
| — in `CLAUDE.md` | **1** |
| `D-v28-14` mentions | **11** |
| `D-v28-15` mentions | **2** |

The whole class lives in **two files**. `D-v28-14` — the **superseded** decision — outnumbers its
**superseding** successor 11 to 2.

## Hypothesis

`D-v28-15` never propagated into the corpus at all: both of its two mentions are iter-225's own banner,
written 40 minutes ago. Every live "the gate host is `odysseus`" claim is therefore un-retracted, and
`CLAUDE.md` — the file every session loads — carries one of them.

## Expected lift

A **census, not a sample**: all 27 mentions classified as **live claim** vs **correctly-historical
record**, every live one retracted in place, and the correctly-historical ones deliberately left alone
with the distinction stated. `D-v28-15` reaches `CLAUDE.md`.

## Phase plan

1. **Seal predictions** (this commit — `probe(M257x/226)`).
2. Establish whether `D-v28-15` existed in the corpus before iter-225.
3. Classify all 27; repair the live ones; leave the historical ones and say so.
4. Re-run the fences; confirm no citation was re-pointed.

## Escalation conditions

- If a "live" claim turns out to be load-bearing for a *measured number* (a baseline attributed to a
  host), it is not a wording fix — the number and its host travel together, and the fence
  `test_baseline_mirror_fence.py` grades exactly that pairing. Such a site is repaired by naming the host
  correctly, never by moving the number.

## Acceptable close-no-lift outcomes

**Finding that most of the 27 are correctly-historical is a first-class result** — it would mean the
corpus already handles supersession well and only the headline needs the retraction.

## Pre-registered predictions — SEALED IN THIS COMMIT

| id | prediction | rationale |
|---|---|---|
| **P-226-1** | before iter-225, `D-v28-15` appeared **0 times** in `corpus/` + `CLAUDE.md` — the superseding decision **never reached the corpus** | it has 2 mentions now and iter-225 wrote a banner citing it |
| **P-226-2** | **≥ 8 of the 27** `odysseus` mentions are **LIVE** claims (present-tense gate-host assertions), not historical records | `D-v28-14` outnumbers `D-v28-15` 11:2 |
| **P-226-3** | `CLAUDE.md`'s single mention is a **LIVE** claim asserting `odysseus` is the current gate host | it is the release-summary line for `build-budget.md` |
| **P-226-4** | `D-v28-15`'s own text **does** exist in `knowledge/plan/` — the decision was recorded, and only its **propagation** failed | iter-225 quoted it from `state.md` |

**If P-226-2 is refuted (< 8 live), the corpus handles supersession better than iter-225's lesson implies,
and the lesson gets narrowed rather than generalised.**
