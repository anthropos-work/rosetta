---
iter: 156
milestone: M257x
iteration_type: tik
status: closed-fixed
date: 2026-08-08
---

# iter-156 — the family runner reported a line the guard never wrote

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)
— census the mechanical classes; stop sampling them.

**Step 0 — re-survey.** `TOK-08` names no single next target; the open-route queue does. Re-surveyed it and
took `SURVEY-M257x-iter152-other-guards-may-read-prose-as-data`, which is the only queued item whose
subject is *mechanically decidable* (a marker either occurs outside its structural position or it does
not) and which iter-152 explicitly left with instructions: **derive the guard list from disk, and grade the
PROPERTY — not the anchoring.** Its stated property:

> *the marker string does not occur in the scanned tree outside its structural position.*

`FIX-M257x-iter155-add-injected-rows-vs-the-pinned-liveness-only-decision` was **deliberately not taken**:
iter-155 reverted it because landing it reverses a pinned decision on partly-unobservable grounds, and this
session can observe no more of a live stack than that one could. Re-landing it silently is the one thing
its own close forbids. Left queued, unchanged.

**Cluster / target identified.** 32 guard modules on disk in `stack-core`; 23 of them read a prose tree
(`corpus/**` or `knowledge/plan/**`); 180 module-level compiled patterns among them. That is the
population the route asked to be enumerated rather than sampled.

**Hypothesis.** iter-152 found one instance (`blocking_state_guard`'s unanchored `search()`). If the class
is real rather than a one-off, the census finds more — and if the census over the corpus-reading guards
comes back clean, then either the class was a one-off, or the census is pointed at the wrong layer.

**Expected lift.** No `N` reading planned. The deliverable is the enumerated population with its
denominator stated, plus a fence over whatever it finds.

**Escalation conditions.** A finding that requires a live stack → route forward (iter-152's half-up-services
precedent). A finding that would reverse a pinned decision → do not land it in-iter (iter-155's precedent).

**Acceptable close-no-lift outcomes.** A census that returns zero and PROVES ITS INSTRUMENT (§9) is a
complete iter — the class is then measured, not assumed.
