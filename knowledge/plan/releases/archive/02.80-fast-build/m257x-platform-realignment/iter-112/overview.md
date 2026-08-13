---
iter: 112
milestone: M257x
iteration_type: tik
iter_shape: tooling
status: archived
opened: 2026-08-06
---

# iter-112 — TOK-07 step 1: the enumerator

**Active strategy reference:** `TOK-07: enumerate the predicate, not the anchor`, **step 1** —
`FIX-M257x-iter109-repair-scope-is-detection-bounded`'s named mechanism.

## Step 0 re-survey

`TOK-07` step 0 closed at iter-111 (both items, one decided and one refuted). Its own reason for
ordering held and was worth the wait: the enumerator's entire output is machine-read, and it ships
`--json` — which would have been unparseable, and would have needed the hidden env var, had step 0 not
gone first. Target confirmed.

## Cluster / target identified

A **per-predicate, corpus-wide site sweep that runs BEFORE any repair and produces the repair's
denominator.** iter-108's reach was graded 46/46 = 100 % against a denominator derived from
`iter-103/raw/` — a prior reading's **detections** — with per-pass recall at 33–83 %, so twins survived
and one became a self-contradiction.

## Hypothesis

The judgement/mechanism boundary can be drawn cleanly: **choosing a predicate's search FORM is
judgement; enumerating that form over the corpus is mechanical and complete.** If the judgement half is
derived by default and **fenced by seed recall** — a form must find the site it came from — then the
enumeration is trustworthy without being unsupervised.

## Expected lift

**No `N` movement, and none is claimed** — the read is step 3. The deliverable is the first
**per-predicate multiplier** this milestone has ever reported, and an instrument that refuses to report
a number it cannot justify.

## Phase plan

Two planned lines (tooling-iter shape): ship the enumerator with its controls; run it on iter-109's 24
predicates and report.

## Escalation conditions

If the enumeration cannot be made trustworthy inside the iter, **say so and route the untrustworthy
part** — a fabricated denominator is strictly worse than none, and it is the exact defect `TOK-07`
exists to end.

## Acceptable close-no-lift outcomes

A measured statement of the instrument's reach — *which predicates it can enumerate and which it
cannot, and why* — is a first-class outcome. The fence refusing its own first run would be evidence the
controls work, not evidence the iter failed.
