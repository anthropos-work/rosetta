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

## D-M257x-268-1 — a sibling registry survived the sweep that fixed its twin

`demo-stack/ensure-clones.sh:310` opens its studio-consumer list with a **hardcoded**
`_studio_repos="cms"` and derives every other member from `repos.yml`. Its comment states the intent —
*"cms goes first so the sanctioned `make init-studio` stays the fetch that actually happens"* — which was
true while `cms` was a live service and has not been since `d11a403`. Guarded only by
`[ -d "$_sdir" ] || continue`, so it is **dormant on a fresh box and live on any box carrying a
`stack-demo/cms/`**. On this box it is live: `stack-demo/cms/studio` is populated, i.e. the studio runtime
`app` builds with was fetched by a decommissioned repo's Makefile and copied across as the donor.

**The clone SET was already fenced, and correctly.** `clone_pin_guard.py` derives the allowed key set from
`repos.yml` and asserts both directions; **iter-222 used it to delete five phantom pin keys, `cms` among
them.** That repair was complete for the registry it covered. The studio-consumer list is a **second
registry, one file over**, and the sweep did not reach it — `platform-alignment.md` §5's *"a named-consumer
list survives the merge that moved the consumer"* (iter-23), occurring inside the repo that wrote the rule
down, against a file a sibling registry's repair had already been run over (§10 iter-194: *a registry that
supersedes a list must reach everything the list does*).

**Why it survived four releases:** it is a **preference, not a dependency**, and preferences do not fail.
On a fresh box the `[ -d ]` guard skips it and the correct branch runs; on a stale box it silently wins.
Both outcomes look like success, and neither produces a log line an operator would question.

Corpus repaired at `corpus/services/cms.md` (the measurement, the asymmetry, and *"do not read a
`stack-demo/cms/` as inert"*). The tooling fix — derive the whole list from `repos.yml`, letting the
`git clone $STUDIO_REPO` branch that mirrors `app`'s CI `additional_repo` be the fetch — needs a tag **and
a pin bump**, so it is routed rather than spent: `FIX-M257x-268-ensure-clones-hardcodes-cms-as-studio-fetcher`.
It should land in the same tag as `FIX-M257x-262-dev-path-needs-the-studio-acquisition`, which is the same
file and the same subject from the other side: **dev has no studio handling at all; demo has it anchored on
a corpse.**
