# iter-235 — decisions

## `D-M257x-235-1` — a repo-relative command is graded against the CLONE SET, never one Makefile

The first reading compared all 394 documented `make <target>` invocations to `platform/Makefile` and
reported 21 distinct / 61 sites failing. That number was **denominator, not defect**: the corpus's `make`
invocations are repo-relative by design (`cd cms && make init-studio`, "per service"), so the population
of legitimate targets is every Makefile in the clone set. Widened, 43 of the 61 are correct documentation.

Recorded because the wrong reading was *plausible* and would have published a 15 %-broken-command headline
— the same shape as iter-234's three instrument corrections and as `§8`'s iter-229 rule (*before comparing
a DECLARED number to an OBSERVED one, ask what the declared number is a number OF*).

## `D-M257x-235-2` — `make force-gen` is UNMEASURABLE, not wrong; not repaired

The single documented target absent from every cloned Makefile is `make force-gen`
(`shared_libraries.md:152`), a **`proto`** target. `proto` is pulled at Docker build via
`GH_PAT`/`GOPRIVATE` and is **never cloned** — so the clone set cannot decide it either way.

Marking it as a defect would repeat exactly the error iter-123 corrected for `infrastructure`: treating a
**clone-set limit** as a measurement limit. It stays as written, and the census reports it in its own
class rather than folding it into either verdict.

## `D-M257x-235-3` — the two repairs are DIRECTORY claims, and were graded as runnable, not as prose

Neither repaired site was wrong about the platform in the abstract. Both were wrong about **what happens
when you paste them**:

- `quick_ops.md` named `backend` (no such directory — the repo is `app`) and `cms` / `jobsimulation`
  (directories `make init` does not create, schemas the platform does not create).
- `graphql-wundergraph.md:254` told you to `cd` into a repo deleted from the platform at `2adcf71`,
  without the caveat its five sibling archived-service docs all carry.

The grading standard applied is the corpus's own, from the compose-profile lesson: **does the documented
command still do the thing**, never does it still parse.

## `D-M257x-235-4` — the clone set was read, never written

No `git fetch`, `reset`, `checkout` or `clean` ran. Every Makefile was read with `git show HEAD:Makefile`
rather than from the working tree, so a dirty clone could not have influenced the result.
