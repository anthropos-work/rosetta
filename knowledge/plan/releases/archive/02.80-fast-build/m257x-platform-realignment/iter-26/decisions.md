# iter-26 decisions

## D-M257x-26-1 — the completed run is USED, not discarded and repeated

iter-25 closed with the re-measure routed forward on the grounds that (a) it was partial and (b) it ran under
a clone predating the second-pass runner fix, so its roster and cockpit manifest described the pre-reset
world. Both grounds were correct when written. Both were checked at iter-26's open, and neither survived:

- **(a) is simply false now** — the run completed, 31 ptreport rows, gate verdict emitted, no scoping flag.
- **(b) is real but immaterial**, and that was established by measurement: the roster's `pt-employee` `eid`
  resolves to the expected user in the DB *after* the reset+reseed. The seed is deterministic; the roster's
  load-bearing field is unchanged.

Discarding a valid full run to produce an identical one would have cost roughly an hour of serial suite time
and bought nothing but the appearance of rigour. **The cheap check that discriminates beats the expensive
check that reassures.**

What the run cannot support, and is not asked to: the cockpit-manifest half was also stale. No Playthrough
clicks a cockpit button (the runner's own comment says so), so nothing in this measurement depends on it —
but a future run that exercises cockpit surfaces must use `fast-build-m257x-iter-25b` or later.

## D-M257x-26-2 — `hiring.recruiter-comparison.UC1` is an OBSERVATION, not an attribution

It flipped from failing to passing in the same run as the two skill-path Playthroughs, and a plausible
mechanism exists (the recruiter scoreboard renders sim metadata sourced from the content layer iter-24
re-pointed). Plausible is not measured.

This milestone's most expensive recurring error is exactly this inference — iter-18 attached "downstream of
the unserved content layer" to two defects on the strength of a plausible mechanism, and iter-19's diff
refuted it. The rule that caught it was refusing to reason from *"the layer I changed was broken"* to
*"therefore this symptom was downstream of it."*

So it is recorded as an open question with its evidence (it was failing at iter-19, it passes now, the only
intervening changes are iter-24's re-point and iter-25's runner path). Whoever closes it should name the
read path and check it, not cite this coincidence.
