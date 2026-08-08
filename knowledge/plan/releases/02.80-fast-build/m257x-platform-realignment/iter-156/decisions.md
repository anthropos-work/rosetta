# iter-156 — decisions

## `D-M257x-156-1` — the defect is in the RUNNER; fixing the warning alone would have hidden it

`claim_census_guard.py:441` passed `maxsplit` positionally, so py3.13+ emitted a `DeprecationWarning` and
CPython echoed the source line to stderr. The one-line fix removes today's instance and **leaves the
mechanism**: `guard_family` merged stdout+stderr and reported the last merged line as a GREEN member's
verdict, so any library that writes to stderr re-creates the defect on any member, silently. Both were
fixed, and the ordering matters — the fence was RED-proofed against the *unfixed* source, because after
the source fix the live arm has nothing to fire on.

## `D-M257x-156-2` — the RED path was worse than the GREEN one

A warning echo is **indented**, which is exactly the shape `headline()` selects findings on. A member
going RED while anything wrote to stderr could have had a line of Python reported as its *first finding* —
the failure iter-87 wrote `headline()` to fix, re-entering through a door iter-87 could not see. Findings
are now taken from the guard's own stdout when it has any; the merged stream stays the fallback so §5
rule 8 and harden pass-20's echo case are untouched.

## `D-M257x-156-3` — the fence is interpreter-dependent, and it says so

The motivating warning exists on py3.13+ and not on py3.9, and **the only interpreter on this host carrying
pytest is 3.9** — the fence's first full run was **18/18 green while the defect was live**. Rather than
pin an interpreter (which would make the fence unrunnable where the suite runs), the mechanism is proven
**interpreter-independently** by driving `run_one` against a synthetic member that writes to stderr on
purpose, and the live census arm **states the interpreter it measured with**. A green there is scoped
evidence about that interpreter (§5 rule 60), never a claim about all of them.

## `D-M257x-156-4` — the speaker test is DERIVED, and a prefix test, on evidence

The three live spellings are `anchor-construct-guard:` (hyphens), `union_apply_guard:` (underscores) and
`claim-census:` — **a prefix of its module name**. Equality would have rejected the summary of the very
guard the repair was written for. A "does it look like a summary" heuristic was rejected outright: it
re-creates the defect one remove up, because a warning's echo looks like whatever the source line looks
like. The live arm asserts **all 17 runnable members report on rung `own-summary`**, so the convention is
measured rather than assumed — and a member that stops honouring it surfaces as a disclosed fallback rung,
which is a finding about that member, not a failure of this fence's premise.

## `D-M257x-156-5` — noise is DECLARED but does not turn a member RED

Dropping non-subject output is the same swallow in the other direction, and is how this warning survived
four releases *while being printed as a verdict*. But a warning is not a finding, and grading it as one
would be the runner inventing a verdict — the thing it exists to stop. **The run discloses (`⚠ NOISE`, plus
a count in the summary line); the fence gates.** `D-M255-1` is the precedent for one assert carrying two
contracts.

## `D-M257x-156-6` — the `CANNOT RUN` sniff still reads the merged stream, and was NOT changed

`guard_family:277` sniffs `"CANNOT RUN"` / `"Nothing was checked"` over merged stdout+stderr. That is the
same class, and narrowing it needs its own evidence — harden pass-20 wrote that line deliberately after
measuring a RED guard echoing a corpus line, and its rc-0 scoping already carries most of the weight.
Routed as `FIX-M257x-iter156-cannot-run-sniff-reads-merged-stream`, **not** silently altered inside an
iter whose planned scope was the verdict line.

## `D-M257x-156-7` — `FIX-M257x-iter155-add-injected-rows…` was deliberately NOT taken

iter-155 built that repair, measured it working, and reverted it because landing it reverses a pinned
decision on partly-unobservable grounds. This session can observe no more of a live stack than that one
could, so re-landing it would be exactly the half-argued reversal iter-155 refused. Left queued and
unchanged; **re-landing it is a decision for the user to see, not a repair for an iter to make.**
