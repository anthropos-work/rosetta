# iter-285 — decisions

## D-M257x-285-1 — refute the premise before diagnosing the mechanism

The defect was reported as an AI-readiness regression. One `GROUP BY status` over
`public.ai_readiness_cycles` returned `active|1 closed|1` — the M219 contract, intact. That single query
moved the search from "which read path broke" to "which world is this", which is where the defect was.
**A bug report names a symptom and usually names a subsystem; the second half is a hypothesis.**

## D-M257x-285-2 — built-but-never-armed is its own class, and this is its first named instance

`cockpit.py --roster` was written at v2.8 M256 for precisely the failure the user hit, is covered by
tests, and is cited by `run-playthroughs.sh` **twice** as the reason a stale manifest fails closed. It was
never passed by `up-injected.sh`. The milestone already knows *parsed-but-never-read is invisible*; this
is the same shape one level up — **implemented, tested, quoted, and not invoked.** The generalisation
worth carrying: **grep for the flag at the CALL SITE, never for the feature in the module.**

## D-M257x-285-3 — a separate array for the cockpit's roster, not a reuse of the override generator's

`roster_flag` already carried the same path to `gen_injected_override.py`. Reusing it in the cockpit
launch would work today and break silently the moment either consumer renamed its flag — and the failure
mode would again be a **disabled check**, not an error. Two consumers, two arrays, one source path.

## D-M257x-285-4 — the log-out swap rides the SAME null guard as the item it replaces

`backToCockpitMenuItem ? null : mapItem(logOutMenuItem, 0)`. Anything else — a separate env check, a
build-time flag — could take Log out away from a real deployment. Tying the removal to the *same*
expression that adds the item makes the off-demo dropdown byte-identical by construction rather than by
care, and the anti-vacuity arm pins that an unconditional Log out has not survived beside it.

## D-M257x-285-5 — `demo-2` is neither restarted nor re-seeded, and the choice is disclosed

Restarting the cockpit would make its seats match the database — i.e. show `pt-world`, not the demo world
the user is looking for. Restoring the demo world needs a re-seed, which destroys whatever they are
validating. **Both are decisions with visible consequences for a session in progress, so both are the
user's.** Recorded as a route.
