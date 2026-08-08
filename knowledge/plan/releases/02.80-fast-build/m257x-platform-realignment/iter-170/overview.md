---
iter: 170
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-08
---

# iter-170 — census the RED-at-HEAD population

## Step 0 — re-survey before targeting

iter-169 routed `SURVEY-M257x-iter169-rotted-assertions-beyond-Thread-_stop` an hour ago and it is the
freshest evidence in the queue. It is also the direct test of a **standing inference** two iters old:

> iter-167: *"A full suite run should now come back genuinely clean; if you see a failure, it is new, not
> the old standing one."*

iter-169 found a failure that was **not new** — `test_buildbench`'s sampler assertion has been RED since
this box moved to CPython 3.14, and it was invisible because the only thing that surfaced it was a *different*
battery's baseline, whose red read as a *different* class. iter-167's family re-measurement was **17 GREEN ·
0 RED · 7 not-run**, and the module iter-169 found red was in the not-run remainder.

So the re-survey confirms the target and sharpens it: the question is not *"are there more `Thread._stop`s"*
(a class defined by one instance) but **"what is the RED-at-HEAD population of this repo, and what is its
denominator?"**

## Active strategy reference

`TOK-08` — *census the mechanical classes; stop sampling them.* "Does this test module pass on this
interpreter, at this HEAD" is as mechanical as a question gets: it is decided by an exit code. Every reading
so far has been a **sample** (a family, a battery, a scoped run), and iter-169 is the second time in four
iters that the un-sampled remainder held a real RED.

## Cluster / target identified

The Python test population of `rosetta-extensions`, enumerated per section:

| section | `test_*.py` |
|---|---|
| stack-core | 59 |
| demo-stack | 34 |
| stack-injection | 7 |
| stack-verify | 5 |
| dev-stack | 5 |
| **total** | **110** |

Go-suite sections (`stack-seeding`, `stack-snapshot`, `clerkenstein`, `playthroughs`, `alignment`,
`stack-secrets` — 264 `_test.go` files) are a different toolchain and are **out of this iter's denominator by
declaration, not by omission**.

## Hypothesis

1. **The RED-at-HEAD population is not zero**, and it has never been enumerated — only inferred from scoped
   greens. `§5` rule 60 says a scoped green is evidence about its scope alone; the corollary nobody has
   spent is that **the un-scoped remainder is therefore UNMEASURED, not green.**
2. **A per-module census is affordable where a whole-suite run is not.** Running modules individually gives
   an *attributable* RED list instead of one aggregate verdict, and the expensive members (the seven
   mutation batteries, ~100 s each) are a known, bounded minority.
3. **The partition needs a third bucket** (`§5` rule 73, and `SURVEY-M257x-iter152-half-up-services-are-ungradeable`):
   a module that requires a live stack is neither green nor a defect. A two-bucket census would report the
   platform's absence as this repo's failure.

## Expected lift

No `N`/`P` reading (`§9`: the metric stays UNMEASURED, not unmoved). The deliverable is the **enumerated
population with a stated denominator** — how many modules run, how many are RED at HEAD, how many are
un-gradeable here and why — plus triage of every RED found.

## Phase plan

- **A — build the census.** Per-module subprocess, bounded parallelism, per-module timeout, three-way (plus
  timeout) classification. Prove the instrument before believing a zero (`§9`).
- **B — run it and report the population**, per section, with the denominator stated.
- **C — triage every RED.** For each: rotted assertion / real defect / environment. Grade by consequence,
  not by class (`TOK-08`'s carried finding).
- **D — land what is landable**, route the rest with named handlers.

## Escalation conditions

- If the census returns **zero REDs**, that is a result only if the instrument is proven to be able to
  report one (`§9`) — a deliberately-broken module must be shown to land in the RED bucket.
- If a RED is a real platform-behaviour defect rather than a rotted assertion, it may exceed this iter;
  route it with a named handler rather than opening a third line of investigation.

## Acceptable close-no-lift outcomes

- The census runs, the population is enumerated, and every RED turns out to be environment-gated
  (no live stack, no Docker) — the denominator and the third bucket are still the deliverable.
