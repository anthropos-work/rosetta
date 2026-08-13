# iter-167 — decisions

## `D-M257x-167-1` — a fixture is frozen; a derived ledger is not; the assertion between them must say which

`claim_twin_guard`'s claim ledger is **re-derived from the milestone's blocker-ledgers on every run**.
The iter-48 answer-key fixture is **pinned at rosetta `cabc3b1`**. The green-twin assertion was written
as *"no refuted claim fires here"* when the capture could only ever support *"no refuted claim KNOWN AT
CAPTURE fires here"*. iter-49 re-adjudicated the same corpus region on a different form and the two
propositions came apart — after which the test was RED at HEAD and three iters shipped over it.

**Rule:** an assertion joining a frozen artifact to a live derivation must be scoped to the frozen
artifact's own denominator, and must say so at the site. This generalizes past this fixture: the two
other answer keys (iter-41's 18, iter-47's 7) have the same coupling and the same exposure.

## `D-M257x-167-2` — the fixture was NOT edited, and that was the harder choice

Making the green twin quiet by deleting iter-49's sentence from it would have taken about a minute and
been a real defect: the capture is perishable (`§5` rule 21), it exists to support a claim about the
INSTRUMENT — *a seven-auditor read missed these while they sat in its own assigned file sets* — and it
cannot be re-taken now that the corpus is repaired. **A fixture edited to make a test pass is a fixture
that has stopped being evidence.**

## `D-M257x-167-3` — the narrowing was required to prove itself before it landed

The repair is a scope narrowing, which is the shape iter-158 caught grading 14 of 14 broken checks
green. So it carries two controls, both green:

- **`test_02b`** (permanent): the same scoped predicate must still find every captured claim on the
  RED fixture. A fence that stopped detecting them fails here instead of passing there.
- **A mutation run at authoring time**: `CAPTURE_ITER = 50` makes iter-49's hits inadmissible and the
  residual clause fires — *"49 not greater than 50 … This is an in-capture miss wearing a different
  path."* The clause is live.

**A narrowing that has not been shown to still fire is not a repair, it is a silencer.**

## `D-M257x-167-4` — the class was closed by MEASUREMENT, not by inference

Two repairs (iter-166's value-change mutant, this iter's answer key) could have been reported as
closing the harden pass's *"RED at HEAD since iters 162/163"* finding. Instead `guard_family.py` was
run at HEAD: **17 GREEN · 0 RED · 0 could-not-check · 7 not-run**, each not-run named with the input it
lacks. Two repairs plus an argument is not the same object as a family-wide run, and this milestone has
spent enough iters on the difference.
