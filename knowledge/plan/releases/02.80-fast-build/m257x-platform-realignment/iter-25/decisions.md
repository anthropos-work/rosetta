# iter-25 decisions

## D-M257x-25-1 — the reset arm REFUSES; the other two sites keep degrading

Three call sites, two treatments, deliberately.

The **reset** arm exits 2 when `stackseed` is missing. Continuing past a failed reset is what makes a
measurement incomparable, and the whole point of `--reset` is that the world is known. A loud refusal is
strictly better than a run whose number cannot be trusted.

The **roster export** and **cockpit-manifest export** keep their documented **non-fatal** behaviour, now
behind an `-x` guard so a missing binary produces their own warning rather than a shell error. They are
refresh steps for surfaces a Playthrough run may not touch at all (the file's own comment: *"Invisible to all
23 Playthroughs, because not one of them clicks a cockpit button"*), and turning them fatal would make a
legitimate shape — a demo with no cockpit manifest on disk — unrunnable.

The asymmetry is the point: **fail closed where the failure invalidates the measurement, degrade where it
does not.** What was wrong before was not the non-fatality; it was that the degradation printed as a shell
error from a line number, which names no consequence.

## D-M257x-25-2 — no clause-2 number is claimed from a 65-of-209 run

The overview wrote the escalation condition before the run started, precisely so this decision would not be
made under the temptation of a nearly-finished suite: *"record what was measured, route the rest; do not
quote a partial run as a clause-2 number."*

Two reasons it would be wrong here, beyond incompleteness:

1. **The denominator is the point.** Clause 2 is `30 live / 0 failing / 0 error`. A partial run cannot
   distinguish "these are the only failures" from "these are the failures so far", and the ptreport gate is
   binding **only** on a full run (`run-playthroughs.sh:300-307`) for exactly this reason.
2. **This run predates the second-pass fix.** Its roster and cockpit manifest describe the pre-reset world.
   iter-26 must **re-run**, not read this log — reading it would be the iter-15 error (comparing worlds)
   wearing new clothes.

The three failing ids observed are recorded in `progress.md` as an observation with its provenance attached,
and explicitly not as evidence. The suggestive absence of the skill-path failures is recorded the same way —
it is the result iter-24 predicts, which is exactly why it must not be counted before the full run.
