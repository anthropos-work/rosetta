# iter-248 — decisions

## `D-M257x-248-1` — the fix for a two-for-two process defect is a DISCLOSURE, not a runner change

The tempting repair is to make `guard_family` run the test suite. It is the wrong one, for the reason
`D-M255-1` already records about a different pair of consumers: **one measurement, two contracts.**
`guard_family` is a fast tree-state check an iter runs many times; the suite is a ~26-minute run. Folding
them makes the fast one slow and gives an operator no way to ask the cheap question.

What was actually broken was **legibility**: `29 GREEN · 0 RED` answers a narrower question than the one
a reader brings to it, and nothing said so. The repair is the SCOPE line — stated on every run, green or
red, with the file count **derived from disk** (M220 D1) and **the two iters that paid for it named**, so a
later reader cannot mistake it for boilerplate and delete it.

## `D-M257x-248-2` — the disclosure names iters 239 and 240 explicitly

Evidence travels with the rule or the rule reads as noise. This milestone has deleted its own correct
exclusions before for exactly that reason. A test asserts the two numbers are present, so a tidy-up that
strips the evidence fails.
