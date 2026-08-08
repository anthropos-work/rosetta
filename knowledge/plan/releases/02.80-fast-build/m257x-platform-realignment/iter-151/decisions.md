# iter-151 — decisions

## `D-M257x-151-1` — grade an absent-value default by the SIDE that reads it

iter-147 booked `STACK_PROJECT` / `STACK_OFFSET` as an open hazard on the **shape** of the default alone:
unset, `STACK_PROJECT` resolves to `anthropos`, the developer's own main dev stack, and `STACK_OFFSET`
derives from it.

The census that settles it is not *"where is this read"* but **"is any reader a writer"**. Every read in
the monorepo is inside `stack-verify/`, a read-only probe section; no seeding, snapshot, injection,
secrets, demo, dev, playthrough, clerkenstein or alignment code reads either variable. So the live
failure mode is a probe pointed at the wrong stack — loud, recoverable, and the class iter-148 already
repaired at the caller — and not a write into someone's dev stack because a variable was missing from an
environment.

**Decision: close the route as a falsification and fence the PARTITION rather than the default.** The
default is fine while only readers read it. The fence asserts no write-side section reads either
variable, carries anti-vacuity on the subject (a rename would empty the census and make every assertion
trivially true) and a RED-proof (the census returns zero, so the write-side predicate is shown matching
synthetic content — `D-M257x-149-3`).

**And the failure message names the right repair, not the easy one.** Someone will eventually have a good
reason to give a write path a target variable. The message says: do **not** add the section to the
read-side list — make its resolution REFUSE an absent value, which is iter-147's own rule. A fence that
does not say what to do when it fires becomes an obstacle, and obstacles get allowlisted.
