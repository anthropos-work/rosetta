# iter-160 — decisions

## `D-M257x-160-1` — the value-side predicate must EXECUTE the derivation, not read for it

iter-159's haystack clause is decided by reading. This one cannot be: whether
`"postgresql redis sentinel backend gotenberg"` is a frozen copy depends on what
`platform_topology.default_services()` **returns today**. So the instrument runs the derivation and
compares token sets for **equality**.

Three consequences, each deliberate:

- **Equality, not containment.** A subset is a different claim and admitting it would flood the census.
- **A hit is actionable by construction.** Its repair is never a rewording — it is `import the
  derivation`, which is `§5` rule 71's prescribed structural repair verbatim.
- **The instrument is exposed to `§9`'s failure mode more than any other in the tree.** Its predicate
  depends on an external checkout; on a box without a platform clone every derivation yields nothing and
  the census reports a serene zero. **An unavailable derivation is `CANNOT RUN` (exit 2), never a
  finding of none** — asserted in the fence, in both directions.

## `D-M257x-160-2` — a repair fixes the side that went RED; the other side stays broken

**The instrument's first run found the class alive at HEAD, in the file M257x repaired for this exact
reason.** `test_bringup_verify_scope_m257x.py` feeds the literal
`"postgresql redis sentinel backend gotenberg"` at **8 sites**, while the tree derives that same list
from the platform's own compose.

iter-155 re-pointed that test's **expectation** to a real derivation — correctly, and it is recorded as
a rule-71 repair. It did not touch the **fixture** the test feeds its subject. The expectation had gone
RED and announced itself; the fixture had not, and did not.

**A test has at least two values — what it supplies and what it expects — and only one of them
announces itself.** Five subsequent iters read this file and none saw it, because every instrument was
looking at haystacks or at assertions, and a fixture is neither.

## `D-M257x-160-3` — declare blind spots as PREDICTIONS, before the run

iter-159 discovered its blind spot after measuring and had to be careful not to retro-fit the
denominator. This iter inverted that: the `overview.md` declared, before any code, that the instrument
would be **blind to (b2)** (iter-157's over-strict direction — no literal involved) and **blind to
iter-158's traceback fixture** (an *inexact* copy cannot equal a derived value).

Both held, along with the fire prediction and the anti-vacuity exit. **A declaration made in advance can
fail; a description written afterwards cannot** — so the three blinds are evidence about the taxonomy,
where iter-159's were evidence only about the instrument.

The second blind is the one to carry: **exactness is what this predicate keys on, so near-misses are
invisible to it** — and a near-miss is the more dangerous defect, because it looks right. Routed as
`SURVEY-M257x-iter160-inexact-copies-are-invisible-to-an-equality`.

### And a rule-71 defect in this iter's own fence — the third inside the thread that owns the rule

The mutation control asserted that lowering `MIN_TOKENS` to 1 **changes the live candidate count**. It
does not — every registered derivation has ≥ 2 members and set equality already excludes a one-token
literal — so a **correct guard read RED against a count that cannot move**. Re-pointed to the floor's
real property (*a one-token derivation must not let a bare word match*), asserted in both directions.

Rule 71 says the durable defence is structural rather than vigilance. This is the third time the rule has
caught someone working inside its own thread, which is the evidence for that claim, not against it.
