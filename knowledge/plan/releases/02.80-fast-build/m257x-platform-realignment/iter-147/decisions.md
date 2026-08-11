# iter-147 — decisions

## `D-M257x-147-1` — the census is INVERTED, and the inversion is what found the defect

iter-146's route asked for the same search over more tokens: grep `skiller`, `skillpath`, `chronos`,
`intelligence`, `storage`, `messenger`, `customerio-sync` the way iter-146 grepped `5050`. Run, that
population is **74 lines / 0 emitters** — registry rows, tests about those rows, and port arithmetic.

The substitution is to enumerate **what the emitters announce or choose**, and grade each against the
platform. Stated as a rule because the two searches have different reach and the difference is not one of
thoroughness:

> **A token census can only find a value that is WRONG. It cannot find one that is ABSENT.**

iter-147's defect is an **empty** compose profile. There is no string to grep for; the defect is the
absence of any string at all, and it presents as `--profiles ` followed immediately by the next flag. No
widening of the token list would ever have reached it — which is the honest reason the route is closed by
substitution rather than by execution.

**The routed token question is NOT dropped.** It is answered with its measurement (0 emitters over 74
lines) and re-routed narrowed: `SURVEY-M257x-iter147-absent-value-class` asks the same *inverted* question
of the other choice-points (`--services`, `--ref`, `--data-root`, the `STACK_*` scope variables).

## `D-M257x-147-2` — the bring-up REFUSES; the teardown keeps its fallback. The asymmetry is deliberate

`rosetta-demo`'s own `cmd_down` already derives its profile (iter-55) and **falls through** to the
un-profiled compose scope when derivation fails — the F-9 rule, *a teardown must never die halfway*,
because an aborted teardown leaves containers holding ports and a registry slot leaked.

`derive_profile` on the bring-up path does the opposite: it `die`s. A bring-up has the inverse obligation
— **announcing a stack that is not there is worse than refusing to bring one up** — and it is the
disposition `up-injected.sh:2156-2157` and `dev-stack:89` already take. Recorded as a decision because
the same function name now carries two dispositions in one file, and a future reader will otherwise read
the difference as an inconsistency to "clean up".

## `D-M257x-147-3` — the fixture was REPAIRED, not bypassed: the empty stub was the defect, encoded

Four `test_tooling.py::RosettaDemoRegistry` tests went RED on the fix. The cause is not the fix:
`setUp` created the stub platform's `docker-compose.yml` with `open(...).close()` — an **empty file** —
and then asserted `rosetta-demo up` returned **0**. That is the defect written down as expected
behaviour: a platform directory from which no profile can be derived produced a *successful* bring-up.

The available repairs were (a) pass `--profile` in the four tests, stepping around the derivation, or
(b) make the stub a faithful platform. **(b)**, and the reason is `§8`'s: a fixture that cannot reach the
code path under test converts a fence into decoration. The stub now carries
`services:\n  backend:\n    profiles: [core, backend, all]` — the anchor row `platform_topology`
actually reads — so those four tests keep exercising the derivation instead of avoiding it.

**Graded before repaired**, per `D-M257x-144-2`: 13 failed at first run, **9 of them pre-existing** and
identical by name to iter-145's demo-stack set; **4 were mine**. The count was not quoted as a backlog
before the names were read.

## `D-M257x-147-4` — the invented word is removed from the announcement AND from the registry

`${profile:-base}` printed `profile='base'` to the operator on success and wrote `"profile":"base"` into
the unified stack registry, which `/stack-list` reads. **Compose has no `base` profile.** The word was
this script's own name for *no profile at all*, and it reads to a human as the name of a profile — so
the one artifact that could have disclosed the hollow stack instead described it as an ordinary one.

Both substitutions are deleted rather than re-worded: `profile` cannot be empty on this path any more, so
a `:-` default here can only re-introduce the euphemism. `services` keeps its `:-all-in-profile`, because
an empty `services` genuinely does mean every service in the profile — the distinction the fence encodes.
