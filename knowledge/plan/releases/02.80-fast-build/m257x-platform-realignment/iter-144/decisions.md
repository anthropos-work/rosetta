# iter-144 — decisions

## `D-M257x-144-1` — a retraction clause holds the CORRECTED pin as often as the retracted one

Rule 63(c′) is right that the corpus's correction idiom keeps a retracted number live in the text. It
is **incomplete**: a correction has two halves, and the second is a **live pin the corpus is right to
publish**. Both sit in the same sentence, inside the same markers, in the same token shape.

Measured on `retracted_pin_guard`'s wrapped arm over its whole population of 10: **7 true, 3 false —
70 % precision**, and all three falses are corrections rather than retractions (`ai_architecture.md:303`
retracts the *absence of a filename*; `hiring.md:304`'s retracted value is *"an earlier range"*, left
deliberately unnamed, while the pin present is the **fix**).

**The remedy is not to tighten the arm.** A form-matching fence cannot read which half of a correction
it is holding — that is a property of the construct, not a defect in the regex. Harden pass 30 was
right to make this arm SURVEY rather than gating. The remedy is to **grade before repairing**, which
is what this iter did. Landed as `§5` rule 67.

## `D-M257x-144-2` — grade a survey arm's findings before treating its count as a backlog

Pass 30 routed *"the **8** true sites across 6 files"*. Measured at this iter's open: **10 findings**,
of which **7** are true. Neither number was wrong to write — the pass had just decided the arm must be
non-gating and correctly judged that grading six documents of prose was an iter's work, not a harden
inline fix.

The hazard is what happens next: **a routed count is an estimate of WORK, and the moment it is quoted
it becomes a measurement of DEFECTS.** This milestone has now watched that conversion happen twice
(iter-138's census, and the *"class stands at 0"* of pass 32). A routed count should carry the word
*estimate* until something grades it.

## `D-M257x-144-3` — a control earns its place by being RUN, not by firing

iter-143 landed `D-M257x-143-1` — *an audit is a predicate too, and it needs a control that is not
another reading*. Applied here one iter later, the control **confirmed the reading 3 for 3**, where at
iter-143 it had overturned **9 of 92**.

That is the outcome worth recording, because it is the one that tempts a future iter to skip the step.
The control cost one command — resolving three pins against the sources they name. **A control that is
only run when it is expected to fire is not a control**; it is a second opinion solicited on doubt,
which is exactly the reader-grading-the-reader loop `D-M257x-143-1` exists to break.
