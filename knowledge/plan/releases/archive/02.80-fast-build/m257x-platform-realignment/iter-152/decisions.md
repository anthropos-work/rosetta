# iter-152 — decisions

## `D-M257x-152-1` — fence two hand-maintained copies against each other; never derive one from the other

`services.sh` and the test-side `REGISTRY_BASES` are two copies of one fact, and iter-145 kept them
**both hand-written on purpose** — reading the port expectations out of the table under test would make
the offset sweep assert nothing (`§8`'s anti-vacuity rule).

The obvious-looking next move — *derive `services.sh` from the platform's compose* — has the same defect
one level up: it would delete the independent copy and turn every downstream assertion into a tautology.
**The fence asserts AGREEMENT between independently-maintained artifacts.** That is what makes a
disagreement information.

Generalised: *a fence between two artifacts is only worth what their independence is worth. Deriving one
from the other converts a fence into a tautology, and the conversion is invisible — the tests still pass.*

## `D-M257x-152-2` — a probe registry's denominator is the platform's PUBLISHED PORTS, not its services

Graded per **service**, `services.sh` is complete: 7 rows for 7 compose services. Graded per **published
host port**, it covers **7 of 10**. Neither number is wrong; the first is not an answer to the question
*"can this stack be half-up and still grade green?"* — and it can.

The guard therefore carries **arm C** as a distinct arm from arm B, and a control asserts that the
4th-port mutation produces **no** A/B/D finding — so arm C is proven load-bearing rather than assumed. The
3 currently-unprobed ports are **declared with reasons**, not silently tolerated.

Generalised, and it sharpens `§5` rule 69: *a census cannot find a value absent from its DENOMINATOR, and
the denominator is a choice made before the census runs. State it, and state what it excludes.*

## `D-M257x-152-3` — a corrected falsification: iter-132's dirty-tree attribution does not reproduce

iter-132 recorded `test_fence_provenance::TestFamilyRefusesAnUnstateableTree::test_the_escape_accepts_and_records`
as *"an artifact of the confound — the fence tree was DIRTY while it ran. **Re-run alone on the committed
tree: `1 passed in 83.11 s`.** Not a defect; proven rather than argued."*

**It reproduced here on a committed, clean fence tree.** The cause was a real defect
(`blocking_state_guard`'s unanchored grading search — see `progress.md` D2), fixed in this iter. The
iter-132 row is **corrected, not carried**: what was booked as a proven non-defect was a defect the
re-run happened not to hit.

This is the milestone's own `correction-vs-retraction` distinction applied to itself — the *finding*
(the test failed) stands; the *attribution* (dirtiness caused it) is withdrawn.

Generalised: *"re-ran it and it passed" falsifies "it always fails". It does not establish "it is not a
defect." A flake and a state-dependent true positive are indistinguishable from one green re-run.*
