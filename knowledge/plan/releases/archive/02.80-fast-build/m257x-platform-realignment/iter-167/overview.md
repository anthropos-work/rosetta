---
iter: 167
milestone: M257x
iteration_type: tik
status: closed-fixed
date: 2026-08-08
---

# iter-167 — the answer key that has been RED at HEAD, and why it was right to be

**Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).
A frozen fixture's agreement with a live derivation is mechanically decidable — the fence either
fires on the fixture or it does not — so this is a census target, not a reading.

**Cluster / target identified.** `FIX-M257x-iter166-iter48-answer-key-red-at-head`, routed by the
iter that verified it: `test_claim_twin_guard_iter48_answer_key::test_02` fails at HEAD, and staging
`git archive HEAD stack-core` into a scratch tree reproduced it with none of iter-166's code present.
It is the concrete instance of the class the last harden pass named — **two fences RED at HEAD since
iters 162/163, with three iters shipped over them.**

**Hypothesis.** The failure is not brittleness in the fence and not a defect in the corpus. It is a
**frozen expectation graded against a moving denominator**: the fixture is pinned at rosetta
`cabc3b1`, while `claim_twin_guard` re-derives its claim ledger from the milestone's blocker-ledgers
on every run, and that ledger has grown by ~117 iters of adjudications since the capture.

**Expected lift.** No `P`/`N` reading (`TOK-08`'s sealed rule; `§9` reads UNMEASURED as UNMEASURED).
The deliverable is a **RED at HEAD, correctly diagnosed and closed** — plus the generalized rule, if
the diagnosis holds.

**Phase plan.** A — reproduce and identify the firing claim exactly. B — decide which of the three
possible subjects is at fault (fence / corpus / fixture-vs-ledger coupling). C — repair without
spending the perishable answer key. D — prove the repair is not a vacuous narrowing. E — record the
rule.

**Escalation conditions.** **A narrowing is the obvious repair here and therefore the dangerous one**
(iter-158: a proposed narrowing would have graded 14 of 14 broken checks green). If the scoped
assertion cannot be shown to still fire, STOP and route instead. Editing the fixture to make the
test quiet is forbidden outright — `§5` rule 21, it would spend an answer key that cannot be
re-captured.

**Acceptable close-no-lift outcomes.** A finding that the fence really is brittle would be a complete
result, and a more important one than the repair.
