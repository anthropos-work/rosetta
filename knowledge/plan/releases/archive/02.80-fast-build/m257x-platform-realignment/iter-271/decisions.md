# iter-271 — decisions

## D-M257x-271-1 — clause 1 is RE-EARNED, not widened; and "cold" excludes the clone set

**Context.** iter-260 established gate clause 1 with three consecutive green cycles at rext
`fast-build-m257x-iter-101`. iter-270 spent that frozen-pin control deliberately and bumped the pin **206
commits**, changing the service-injection and studio-acquisition paths — the code a cold cycle exercises.

**Decision.** Clause 1 is re-proven at the shipping pin before anything is built on it: three consecutive
cycles, `DOWN_RC=0` / `UP_RC=0`, `autoverify green:true / warnings:0` (cycles 2/3/4, plus cycle 1 as a
fourth). **The clause is not re-cut, relaxed or re-interpreted** — it is the same clause, re-earned against
the tooling that will ship.

**And the scope is stated with the green, not left to be assumed.** `down --purge` removes containers,
network, volumes, data and images; **it does not remove the clone set**. So the acquisition path prints
`app: studio/ already populated — reusing (idempotent)` and its fetch arm does not run — on any number of
consecutive cycles. Three greens are therefore **not** evidence about the studio repair iter-270 shipped.
That repair was graded directly instead, with its precondition removed (rc 1 on a missing `repos.yml`,
rc 1 on one declaring no repo, the four live repos when present).

**Rejected:** purging the clone set too. It would turn every cycle into a full platform re-clone and make
the gate measure GitHub's availability alongside ours. The purge boundary is right; the omission was
saying where it falls.

## D-M257x-271-2 — a detached launch costs an `rc`; instrument the runner, not the operator

**Context.** Cycle 1's bring-up was launched `nohup … &` so the session stayed responsive. `READY` is
defined as *`up-injected.sh` exits 0 **and** `autoverify.json` is green* (`build-budget.md`). The detached
launch captured the second and **discarded the first** — cycle 1 has a green verdict and a success marker
in its log, but no exit code, and no way to recover one after the fact.

**Decision.** Cycles 2–4 run through a fixed three-line runner that records `DOWN_RC`, `UP_RC` and real UTC
timestamps around both halves. Cycle 1 is reported as **corroborating**, explicitly held out of the three
that carry the clause, and the gap is disclosed rather than smoothed by asserting the log line is
equivalent to an `rc`.

**The generalization** is Lesson 2 of this iter: convenience at the call site is where evidence goes
missing quietly. A launcher that cannot report its own exit status is not an instrument, however green the
thing it launched.
