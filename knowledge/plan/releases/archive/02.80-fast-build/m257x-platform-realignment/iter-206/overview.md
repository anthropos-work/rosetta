---
iter: 206
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
---

# iter-206 — deriving the comment figures, starting where this run has been standing

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.* Fifth
consecutive iter on the measurement-literal class, and the first to **verify** members rather than
enumerate them.

**Step 0 — re-survey (mandatory).** iter-205 sized the class and routed
`SURVEY-M257x-iter205-comment-provenance-notes-are-the-highest-risk-standing-figures`: **96 standing
comment figures**, of which **six** sit in `claim_census_guard.py` and describe populations that module
computes. That is the smallest subset where *derivation* is available rather than judgement, and it is
in the file this whole run has been editing.

**Cluster / target identified.** iter-205 already had one confirmed stale instance in hand (`949 pairs`
against a printed 1,015) and closed at the cap without repairing it. Deriving all six is bounded, and
the module is the one whose populations moved most this run.

**Hypothesis.** Figures inside the module that computes them are the ones a fence can actually hold, so
this subset is where the `standing` bucket stops being merely *sized*. The expected finding is that the
run's own edits made some of them stale — the strongest possible demonstration of the class, and the
least comfortable.

**Expected lift.** All six derived; the stale ones repaired; the whole subset fenced so it recomputes.

**Phase plan.** A: derive each of the six. B: repair. C: fence with an anti-vacuity arm. D: close.

**Escalation conditions.** If a figure cannot be derived without running something expensive, say so and
leave it reported — a fence that needs a demo stack is not a fence.

**Acceptable close-no-lift outcomes.** If all six turn out current, that is a complete iter: it would be
the first *verified* subset of the standing class, and the route's *"sized but unverified"* wording could
finally be narrowed with evidence.
