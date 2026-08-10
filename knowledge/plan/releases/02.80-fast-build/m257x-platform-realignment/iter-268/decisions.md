# iter-268 — decisions

## Pre-registrations — SEALED BEFORE THE MEASUREMENT

Sealed in this iter's first commit, corpus at `c33d2c8`. Known at seal time: the six directory names exist
under `stack-demo/`. Nothing about their provenance, freshness or use has been read.

**PR-1 — the population is 6 in `stack-demo/` and 0 in `stack-dev/`.**
Repos present on disk but absent from `repos.yml`, excluding the legitimately-unlisted ones
(`ant-academy`, `rosetta-extensions`, `app/studio`). *Risk:* `stack-dev/` carries `studio-room`, which may
or may not fall inside the definition, and the demo count may differ from six.

**PR-2 — nothing fetches them today.**
`ensure-clones.sh` (the demo bootstrap) clones only what `repos.yml` lists plus its explicit `ant-academy`
phase, so the six are leftovers. *Risk:* real and the most valuable branch — if the bootstrap still names
any of them, the user's *"no longer treated as part of the project"* is **false on disk**, not just untidy.

**PR-3 — nothing builds from them.**
No compose service, no build context and no rext write path names any of the six on a current bring-up.
*Risk:* falsifiable; a lingering `build: ../cms` would be a live defect.

**PR-4 — the corpus instructs cloning none of them outside a marked-historical context.**
i.e. iter-265's fence population is complete for the *acquisition* verb specifically. Predicts **0**.
*Risk:* `make init` prose and per-service "Local Development" sections are exactly where such an
instruction hides, and iter-265 graded markers, not verbs.

**PR-5 — they are fossils, not fresh.**
Each of the six clones' HEAD commit predates the `repos.yml` change that removed it. *Risk:* if any is
newer, something fetched it *after* the removal, which points at a live code path rather than a leftover.

## Escalation clause (pre-registered)

**Nothing is deleted in this iter, whatever the outcome.** A stale clone is the evidence; tidying it
destroys the measurement, mutates a workspace nobody asked to change, and would make PR-5 unre-runnable.
If a repair requires removing a directory, it is described and routed, never performed.
