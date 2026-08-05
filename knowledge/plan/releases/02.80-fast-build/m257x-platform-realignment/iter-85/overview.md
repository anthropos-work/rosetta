---
iter: 85
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-05
---

# iter-85 — repair Q2: the claims about facts that were DELETED

**Type:** tik, under `TOK-05`. The first repair since iter-81, and the first ever graded by a
post-condition.

## Step 0 — re-survey

Gate **4 of 5**, unchanged. Both repos clean and pushed (`rosetta c1d4a27`, `rext 24819f0`). 6 corpus
guards exit 0. Ground truth unchanged since iter-84's open.

## Cluster / target identified — and why NOT all 40

iter-84 adjudicated 40 upheld across five predicate classes. **This iter repairs Q2 only**, plus the two
confirmed leak sites and the live rext defect. The scope is declared narrow **on purpose**:

- **Q2 is the class where re-anchoring is the WRONG repair.** Seven present-tense claims about facts the
  platform *deleted*; §4 Trap A says re-pointing them produces a correctly-cited false statement. They
  need restating or dropping, which is judgement work that does not parallelise.
- **Q1 (13) is one re-derivation per cluster, not 13** — the `ai-readiness.md` seven are a single rext
  commit. It is cheap, but it is cheap *later*, and mixing it in here would blur the grading.
- **A repair I cannot finish reproduces iter-81 exactly.** iter-83 measured what happens when a repair
  reports a completeness it did not measure. Declaring a scope I can complete, and grading against it,
  is the whole lesson.

**Q1, Q3, Q4, Q5 and the P4 membership sweep route to iter-86** with the adjudication ledger as their
work list.

## Hypothesis

Repairing by predicate with a **per-anchor post-condition** yields 100 % reach over the declared scope —
against iter-81's 74.1 %. **Pre-registered:** `repair_reach_guard` reports **0 unreached** over this
iter's own ledger, or the iter does not close as `closed-fixed`.

## Phase plan — three planned lines (declared)

1. **Q2 — the 7 deleted-fact claims.** Restate or drop; never re-anchor.
2. **The 2 confirmed leak sites** (`CLAUDE.md:285`, `platform-alignment.md:1305`).
3. **The live rext defect** — `dev-stack:186`/`:414` `profile="graphql"`, plus the stale
   `gen_injected_override.py` comments that describe a live `graphql` profile.

## Escalation conditions

- Any Q2 site where the correct restatement is not derivable from source at the ref → leave unchanged
  and record, rather than guess.
- `DEF-M257x-iter80-storage-prod-bucket` stays held; `storage.md:55,:154,:181` unchanged.
- **`ai_architecture.md:225` must NOT be edited** — adjudicated CORRECT (`D-M257x-84-5`).

## Acceptable close-no-lift outcomes

If a Q2 restatement turns out to need a platform-repo fact this milestone cannot establish, that site
closes unrepaired with the falsification recorded.
