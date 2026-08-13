# iter-158 — decisions

## `D-M257x-158-1` — the routed repair was REFUTED by one command

`FIX-M257x-iter156-cannot-run-sniff-reads-merged-stream` proposed narrowing the rc-0 sniff to stdout.
Forced against an empty repo-root, **14 of 14 members report off the `merged-last` rung** — stdout empty,
the could-not-run message on stderr. The narrowing would have graded all fourteen GREEN. iter-156's
routing note (*"narrowing it needs its own evidence"*) is what kept the defect out of the tree; the
evidence, when taken, said do the opposite.

## `D-M257x-158-2` — iter-156's noise disclosure was firing falsely on the same run

`⚠ NOISE demo_knob_guard: … MISSING: the defaults contract … An absent contract is a FINDING, not a skip.`
That is the guard's own sentence. Classifying a stderr line as foreign whenever it fails `speaks_for`
mislabels every guard's continuation prose. **A disclosure that mislabels its subject's own words is the
defect it was built to stop, pointed the other way.**

## `D-M257x-158-3` — two authorship questions on one stream, with OPPOSITE safe defaults

For the **verdict line**, an unrecognised line must not be accepted (accepting an interpreter echo as a
verdict is the iter-156 defect). For **authorship**, an unrecognised line must not be taken away from the
guard (taking its words is the iter-158 defect). So `speaks_for` recognises the guard and defaults to
*not the guard*, while `interpreter_noise` recognises the interpreter and defaults to *the guard*. The
cost is stated rather than hidden: noise from a subprocess or a C-level library is attributed to the
guard, and that blind spot is routed as `SURVEY-M257x-iter158-noise-classifier-is-narrow-by-choice`.

## `D-M257x-158-4` — an inexact fixture under a loose classifier is a test that proves nothing and says it did

iter-156's traceback arm used `Traceback (most recent call)` — a string CPython never emits. Its own loose
classifier accepted it, so the arm passed on a fixture that was not its subject. Tightening the classifier
is what audited the fixture, one iter later. **A test's fixture is a claim too**, and the only thing that
grades it is a mechanism strict enough to reject a wrong one.
