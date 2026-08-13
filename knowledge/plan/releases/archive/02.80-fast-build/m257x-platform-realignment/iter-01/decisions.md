---
milestone: M257x
iter: 01
---

# iter-01 — decisions

Milestone-level doctrine from this iter lives in the milestone-root `decisions.md` as `D-M257x-1/2/3` and
`TOK-01`. Intra-iter notes only below.

- **d1 — measure before authoring strategy.** The five open questions were answered by parallel probes against
  origin HEAD *before* `TOK-01` was written, because a strategy authored first would necessarily encode the
  stale picture the overview warns about.
- **d2 — probe reports are evidence, not conclusions.** Three probe claims were refuted on re-verification
  (see `D-M257x-1`). Every probe finding carried into a deliverable was independently re-checked by this tok.
- **d3 — the protocol doc's procedure was executed, not just written.** All six §4 detection signals were run
  end-to-end; an unexecuted procedure is a hypothesis.
- **d4 — the `corpus_index_guard` was watched going RED before being trusted**, per the release's own
  "prove a check can go RED" rule (M256 found 43 that could not).
- **d5 — `re-scope-trigger` deliberately NOT fired.** The target is provably moving (platform PR #20 and `app`
  PR #1103 both open), which is tempting. But the trigger's condition is *two consecutive alignment attempts
  invalidated by mid-milestone commits*, and **zero** attempts have been made. Firing it now would skip the
  work to reach its own precondition. Recorded instead as the standing risk `TOK-01` is designed around, with
  the instrument-first ordering as the mitigation.
