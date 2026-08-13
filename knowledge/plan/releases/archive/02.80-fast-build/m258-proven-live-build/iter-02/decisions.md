# M258 iter-02 — decisions

## D5 — the batch half was NOT measured, deliberately, rather than measured meaninglessly

The tik's primary deliverable was the batch half. The stack came up `rc=0` and `autoverify green:true`,
so running the suite was *possible*. It was not run, because the ISOLATION evidence shows all three UI
images are wired to a real Clerk app rather than Clerkenstein — and under that wiring the cockpit's
password-free hero logins cannot establish a session (`verification.md` § M218 iter-03 D8/F-6). The
suite would have produced a wall-clock **30 login failures long**.

A number obtained that way is the failure mode `TOK-01` was written against: it would enter the record
as "the batch half", survive into the composed arithmetic, and be wrong by an unknown amount. **A
refusal to measure is a result; a meaningless measurement is a liability.** Recorded as an explicit
non-measurement with its reason, not as an omission.

## D6 — this is `budget-exhausted`, not `user-blocker` — re-graded before reporting

Phase 5 §4 lists *"the protocol's test gates return RED"* as a user-blocker, and ISOLATION is a gate
clause that went RED. Graded against the section's own NOT-list, it is not one:

- **Nothing here needs a user *decision*.** There is no fork. The mechanism is named to two `file:line`
  writers (`inject.py:89`, `up-injected.sh:2036`); the repair is to stop appending and stop swallowing.
  A defect with one obvious remedy is a fix, not a choice.
- **I did not introduce it.** Checked before routing, not asserted: the re-pin's diff over
  `up-injected.sh` is 21 insertions / 9 deletions and **every hunk is a comment or a `log` string**.
- **The blast radius is bounded and the user's stack is clean.** `demo-2`'s last publishable key is the
  minted one; `demo-1` is `--no-public-host`, localhost-bound.
- Phase 5 §4's NOT-list names this shape directly: *"new findings discovered mid-iter → route forward,
  continue."*

So the exit is the honest one: the session's working budget is spent with the iter **closed and
committed** and the tree clean — `budget-exhausted (between iters)`. The orchestrator has twice had to
overturn an over-escalation this release; this is deliberately not a third.

## D7 — `demo-1` is left UP, and that is the contract, not an oversight

`overview.md` § *Batch-gate behaviour* says the stack is left UP regardless, and the `autoverify`
precedent is that a test bug must never cost a good demo. Two further reasons here:

1. **It is the reproduction iter-03 needs.** The 24-block `.env.demo-1` and the three foreign-keyed
   images are the evidence; a teardown destroys them.
2. **It cannot be reached.** Brought up `--no-public-host` per `D2`, so it is localhost-bound.

⚠️ **It must not be browsed.** Its UI tier would talk to a real Clerk app, which `safety.md` forbids
for a demo. Stated here so a later reader treats it as a quarantined artifact rather than a usable
stack. The user's `demo-2` and dev stack were untouched throughout — verified resident before and after.
