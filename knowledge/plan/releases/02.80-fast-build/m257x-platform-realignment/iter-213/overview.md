---
iter: 213
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-09
---

# iter-213 — the backlog registry reads a route's NAME as its VERDICT

**Type:** tik
**Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop sampling them.*
Class under work: **the milestone's own route registry** — the one iter-212 routed forward.

## Cluster / target identified

iter-212 routed `SURVEY-M257x-iter212-a-retraction-does-not-reach-the-code-that-acts-on-it`. Surveying
the instrument that owns route dispositions — `route_disposition_guard`, `FENCE-M257x-iter183`, *"the
milestone's own BACKLOG is a registry"* — found something one grain below that route:

`classify()` docstrings itself *"Which disposition does this segment record? One segment, one verdict."*
It is handed the segment **with the route id still in it** (`collect()` extracts ids from the same
string it classifies, and never removes them), and `REOPEN_RE` matches `retract|refut|supersede|
RE-?OPEN|CORRECTED|new evidence` case-insensitively. **A route whose SLUG names one of those words is
classified by its own name.**

`FIX-M257x-iter144-correction-vs-retraction-unfenced` is carried forward as a bare id in twelve
consecutive iters. **Eleven of those twelve record nothing at all** — no verdict text, just the id — and
all eleven are booked as **`reopen`**, which is one of the four excuses that suppresses this guard's
contradiction rule.

## Hypothesis

Grading the segment **with the route ids removed** makes `classify()` grade what the segment RECORDS
rather than what the route is CALLED, and the change is provable on a staged milestone rather than
argued.

## Pre-registered, sealed in this iter's FIRST commit — before any repair

Measured at corpus `c3d52b9` / rext `3e012b5`, `/usr/bin/python3` 3.9.6, `stack-core`, Python.

- **S1** — **8 of 367** route ids carrying a disposition match a disposition regex **in their own name**:
  five match `REOPEN_RE` (`…cannot-see-retractions`, `…retraction-idiom-sweep`,
  `…correction-vs-retraction-unfenced`, `…ledger-carries-a-retracted-retraction`,
  `…a-retraction-does-not-reach-the-code-that-acts-on-it`) and three match `PARTIAL_RE` via `\barm\b` /
  `\bhalf\b` (`…path-arm-window`, `…orphan-arm-is-the-residual`, `…half-up-services-are-ungradeable`).
- **S2** — of **1,447** segments carrying at least one route id, **14** change verdict once the id text
  is removed, across **3** distinct ids. Transitions: `reopen → other` ×13, `reopen → open` ×1.
- **S3** — **`violations()` is 0 before and 0 after.** The defect is **LATENT**, not a live false-GREEN,
  and this iter will not claim otherwise. `§5` r60 — *a scoped green is evidence about its scope alone.*
- **S4** — it is nonetheless load-bearing in the direction that matters: `reopen` is one of the four
  excuses that suppress the contradiction rule, so for these routes **every bare re-listing supplies its
  own excuse.** To be proven on a **staged** milestone, since S3 says the live tree cannot show it.
- **S5 — a methodology correction against this iter's OWN first probe.** The first measurement of S2
  returned **40 events across 26 ids** and was **wrong by ~2.9×**: it compared a verdict computed on the
  FULL segment against one recomputed on the **240-character truncation** stored in the event tuple. The
  same class as iter-209's hand-written slugger — *when the question is about the input, hold the
  machinery fixed.* The corrected probe re-walks the sources and classifies both ways on one string.

**STOP CONDITION, sealed before the repair:** if the id-stripping repair changes `violations()` or
`malformed()` on the live milestone — i.e. it manufactures a finding rather than removing a blind spot —
**do not land it.** A false RED in a registry every session brief quotes is worse than the silence it
replaces (`_pass_positions`' own recorded lesson, iter-194).

## Expected lift

The registry grades what a segment records. The blind spot is closed **and proven closable** by a staged
control, which S3 says is the only way it can be proven at all.

## Phase plan

A: seal this record. B: strip ids before classifying. C: staged control proving a suppressed
contradiction now fires. D: anti-vacuity + both-way reconciliation. E: re-run, check the stop condition.

## Escalation conditions

Stop condition fires → land the measurement + the survey only, route the repair.

## Acceptable close-no-lift outcomes

If the staged control cannot be built to fire, S4 is falsified and the finding downgrades to a naming
convention — which would itself be the deliverable.
