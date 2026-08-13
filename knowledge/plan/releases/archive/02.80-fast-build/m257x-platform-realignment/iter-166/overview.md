---
iter: 166
milestone: M257x
iteration_type: tik
status: closed-fixed
date: 2026-08-08
---

# iter-166 — ask the guard: the accept side, reported for the first time

**Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)
— *census the mechanical classes; stop sampling them.* A waiver is the most mechanical clause in the
family: it either suppressed a finding on this run or it did not. That is decidable — but only from
**inside** the guard, which is why no census of it existed.

**Cluster / target identified.** iter-165's close routed
`FIX-M257x-iter165-fences-do-not-report-which-waivers-they-honoured`, and it is the freshest and
sharpest item in the queue. iter-165 tried to audit the accept side with a **second instrument** and
withdrew all 11 of its findings as artifacts of its own normalisation; its falsification named the
only correct instrument — **the guard itself** — and observed that the guards do not report it.

Re-survey (Phase 1 Step 0): the target is untouched. Four waiver files, four consumers, and a
mechanical census over `stack-core/*.py` for the `WAIVERS_REL` / `WAIVER_FILE` construct returns
exactly those four — so the population is enumerated, not hand-listed.

**Hypothesis.** Each waiver-carrying guard can record which waivers it honoured **with the same
matcher it already uses to suppress**, and print it. That makes the accept side measurable from
outside a guard for the first time, without re-implementing anything.

**Expected lift.** Not a `P`/`N` movement — no reading is being taken (`TOK-08`'s sealed rule
forbids a successor strategy, and `§9` reads UNMEASURED as UNMEASURED). The deliverable is a
**standing enumeration**: for each of the four guards, `honoured / dormant` with its denominator.

**Phase plan.**
- **A** — census the accept-side mechanisms mechanically; state the denominator.
- **B** — establish the defect: does any guard name a honoured waiver?
- **C** — build the shared reporting layer; wire all four through their OWN matcher.
- **D** — take the first accept-side reading on each.
- **E** — fence it: the reporting paths get tests, or they are docstrings (`§8` rule 5).

**Escalation conditions.** If a "dead waiver" count appears, **check the instrument before the
corpus** (iter-165's lesson, and this iter's own near-miss proves it was not yet learned). If a
guard's suppression predicate would have to be duplicated to report it, STOP — that is the shape
iter-165 withdrew.

**Acceptable close-no-lift outcomes.** A finding of *0 dormant waivers everywhere* is a complete
result: it converts iter-165's withdrawn-and-uncertain "0 provably dead" into a measured 0, taken by
the only instrument entitled to take it.
