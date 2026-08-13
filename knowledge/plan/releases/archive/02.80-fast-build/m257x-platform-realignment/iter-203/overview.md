---
iter: 203
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
---

# iter-203 — the measurement literals a `print()`-only census cannot see

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.* Same class
iter-199 closed for one site-kind, at the site-kind it left out.

**Step 0 — re-survey (mandatory).** iter-199 shipped `printed_measurement_literals`, which walks
`print(...)` calls only — recorded as `SURVEY-M257x-iter199-the-literal-census-reads-PRINTS-only` and
still open. Measured now over rext **non-test** modules: **95 measurement-shaped numbers across 62 sites** (the census itself; an early scratch probe said 85 across 55 — see `D-M257x-203-2`)
live in string literals that are *not* inside a `print()` — module and function docstrings, mostly. The
route is live, the population is bounded, and it is the same defect shape: a number nobody derives, that
nothing reads, and that stays green through every change to what it describes.

**Cluster / target identified — and the sharpest instance is one iter-202 created.**
`claim_census_guard.py:884` states *"205 of 695 line-pinned citations"* are bare basenames, as a standing
property of the corpus. Derived on this tree it is **292 of 704**. Both operands moved, and the numerator
moved mostly *because of iter-202's own parser fix* — the recovered `.tsx`/`.json` targets are
overwhelmingly bare filenames. So the first member of this class is a figure this milestone invalidated
one iter ago and did not notice, which is precisely
`SURVEY-M257x-iter202-published-citation-figures-predate-the-truncation-fix` landing inside the instrument
itself.

**And the grain is part of the defect.** *"205 of 695"* names no unit. This tree derives **979 / 410** at
pair grain and **704 / 292** at distinct-citation-text grain; only the second is what the sentence means.
A figure whose unit is unstated cannot be checked, which is `§9`'s *name the unit* rule applied to a
docstring.

**Hypothesis.** The class is bounded and mostly *narrative* (dated design history, document-relative
distances), so a blanket fence would be wrong. What is mechanical is: (1) **enumerate and ratchet** the
population so it cannot grow unseen, and (2) for the figures that state a **current, derivable** property,
make them derive themselves — the iter-199 repair, whose docstring analogue is a fence that recomputes
and asserts rather than a self-formatting `print`.

**Expected lift.** The `print`-only route closed; the class sized with its denominator stated; the one
proven-stale figure repaired **and** fenced so it cannot rot again.

**Phase plan.** A: census the non-print population + classify. B: derive and repair the stale figure with
its grain named. C: fence it. D: ratchet the population. E: close.

**Escalation conditions.** If the census's undated/derivable split turns out not to be mechanically
separable, report the population and the limit rather than shipping a fence that grades narrative prose
as a defect — that would be the over-reach `D-M257x-202-4` priced twice already.

**Acceptable close-no-lift outcomes.** If every undated figure turns out to be narrative and the only
derivable one is the one already known, that is a complete iter: the class is **sized, bounded and
declared small**, which is what a census is for even when it finds one member.
