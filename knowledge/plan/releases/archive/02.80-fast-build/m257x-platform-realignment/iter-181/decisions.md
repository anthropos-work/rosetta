# iter-181 — decisions

## `D-M257x-181-1` — the survey's implied denominator is REFUTED, and publishing it would have been the defect

`SURVEY-M257x-iter179-readme-indexes-test-modules-unmeasured` asks how completely
`stack-core/README.md` indexes **test modules**. Measured at HEAD `3a8f5b4`:

| denominator | reading |
|---|---|
| all `tests/test_*.py` on disk | **10 of 63** |
| **mutation batteries** (`*mutation_battery*.py`) | **6 of 7** |

**`10 of 63` is not a coverage gap.** The index's subject is the fence family and its batteries; the other
53 are per-guard *behaviour* suites, which the index deliberately does not list — it lists the guard. Had
the survey been answered on its own terms, the milestone would have published an 84 % gap that does not
exist, in a corpus whose whole quarrel is with numbers whose denominators nobody stated.

**Decision: answer the question by fixing its denominator, and record the refuted one.** §9 iter-159 —
*grade the instrument at the grain of its claim* — applied to a **question** rather than to a fence. The
survey closes as *ill-posed until scoped*, not as *measured and fine*.

## `D-M257x-181-2` — at the answerable denominator there was a real gap, unseen for 133 iters

`test_repair_leak_guard_mutation_battery.py` shipped at **iter-48** (`932554e`, 2026-08-02) and had **no
README row** at iter-180 — the one battery of seven that nobody could find from the index. 20 mutants, 19
kills, and it is the only proof that the *"did this commit FINISH?"* fence can fail.

**Nothing was wrong with anyone's reading; there was no denominator, so there was nothing to be short
of.** The row is added, and the arm that would have said so ships with it — **RED-proven against the
unrepaired tree first**, naming the file.

## `D-M257x-181-3` — the naive instrument was RUN, DISCARDED, and KEPT as a control

The first form of the resolve arm checked README-named `.py` files against `stack-core/` +
`stack-core/tests/`. It reported `exposure_claim_guard.py` **missing**. It is not missing — it lives in
`stack-injection/`, and the README says so in a note under its own table.

**Decision: scope the arm to the REPO, and keep the discarded version as a mutation control.** A fence
scoped more narrowly than its subject manufactures findings — the instrument-side form of
`D-M257x-122-4` (*a stale substrate FABRICATES defects*). The control asserts that the cross-section
reference is still what makes the narrow scope wrong, so if the README stops naming anything outside
`stack-core/` the control goes RED **on purpose** and the narrowing has to be re-derived rather than
silently adopted.

## `D-M257x-181-4` — one member of a DIFFERENT open survey is measured, not repaired

`repair_leak_guard.py` — the guard, not its battery — also has no README row. That belongs to
`SURVEY-M257x-iter175-readme-fence-index-is-16-of-27`, whose open question is *which derivation the index
is meant to be complete against*.

**Decision: record it as evidence for that survey; do not repair it here.** Adding 1 of its 11 missing
members would move a **published, fenced triple** (`16 of 27` / `16 of 26` / `15 of 26`) for no gain, and
would answer none of the question that survey actually asks. The battery arm is a complete class; the
fence arm is not, and mixing them would leave both half-done.

## `D-M257x-181-5` — the OTHER route probed this iter was set aside WITH a number

`SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` (the observed half) was the first candidate.
Probed before targeting: **412 count-mentions across 115 files** in the milestone dir. The defect is only
where a `passed` count is used *as the executed population*, which no regex decides — and iter-173 already
priced the observed half as out of reach of any instrument that does not re-run a suite at a past ref.

**Decision: set it aside, and say how big it is.** iter-178's rule turned on this milestone's own backlog:
*a NOT-REACHED clause is a measurement or it is a mood.* A route declined without a number cannot be
ranked against the route that was chosen instead.
