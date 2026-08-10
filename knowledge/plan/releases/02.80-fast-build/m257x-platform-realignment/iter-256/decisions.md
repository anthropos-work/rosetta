# iter-256 — decisions

## `D-M257x-256-1` — the user's closing condition is recorded as binding, and it re-orders the milestone

**Supplied by the user on 2026-08-10, and it supersedes "gate 4 of 5" as the definition of done:**

> the milestone closes ONLY when the CURRENT `main`/tagged branches of the still-relevant platform
> repos assemble into a WORKING STACK — **BOTH demo AND dev** — and the corpus reflects that, with the
> **deprecated/removed repos no longer treated as part of the project**.

**Three consequences, each of which changes what an iter may target:**

1. **Advancing the active clones onto current platform code is IN SCOPE AND REQUIRED.** It had been
   routed and deliberately not taken for many iters, on the ground that it changes what a demo builds
   and `demo-1` was up and green with clauses 1–2 proven on it. Under the user's condition that
   reasoning inverts: **a stack built from a stale pin does not count.**
2. **The working-stack proof must cover DEV as well as DEMO.** Every working-stack proof this milestone
   has ever attempted was demo-only. Treat "dev assembles too" as genuinely unproven — iter-243 already
   found the seeder restarting a *demo* container on dev stacks.
3. **"Deprecated repos no longer part of the project" is a census obligation**, not a one-off. The
   phantom removal from the clone pin (5 repos: cms, jobsimulation, storage, messenger, roadrunner) was
   one instance; where the corpus still treats removed repos as live is squarely in scope.

**Not decided here:** whether the physical checkout advance lands in this iter. That is sized against
the measurement, per §7 rule 4c — *you do not have to TAKE an advance to measure it, and you should
measure it first.*

## `D-M257x-256-2` — the fetch is part of the measurement, and it moved a number in the first minute

The brief carried `ant-academy` at **+9** commits. A real `git fetch origin` against the six clones
made it **+10** (`c885dab2` → `249430c3`), while the other five were unchanged. **A remote-tracking ref
is a cache, not a remote** — stated in the brief as a rule and demonstrated on the very repo set the
advance is about, before any other work.

**Standing consequence for this milestone:** any advance arithmetic (`+28 / +12 / +10`) is a reading
with a timestamp, never a property. Every figure in this iter names the fetch that produced it.
