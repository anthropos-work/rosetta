# iter-198 — decisions

## `D-M257x-198-1` — exposure is measured at CITED-LINE grain, and the two weaker grains are kept visible

Three grains were available and each answers a different question:

| grain | this tree | what it answers |
|---|---|---|
| clones behind origin/main | **6** | is the substrate current |
| pairs materialized from the clone set | **949** | how much evidence *could* be affected |
| files drifted HEAD..origin/main | **30** | which files *could* be affected |
| **pairs whose cited lines differ** | **19** | **which excerpts actually read different bytes** |

The last is the one a reader of an excerpt needs, and it is 2 % of the grain the warning implied. The
`same` bucket (**1** here) is reported rather than folded away precisely because it is the difference
between *"this file changed"* and *"this citation changed"*; without it, `on_drifted_file` would be
quoted as the exposure and would over-state it — by one pair today and by an unknown amount on any other
tree. `absent` (**0**) is separate for the same reason: a cited file that does not exist at origin is a
different defect from one whose lines moved.

## `D-M257x-198-2` — the direction is UNDETERMINED, and the prior single-direction wording is RETRACTED

`KNOWN_WEAKNESS` (5) and `substrate_of`'s docstring both stated the false-RED direction as *the*
consequence. The retraction is not a hedge added for safety; it is a correction:

- **false RED** — the corpus states what is true at `origin/main`; the stale excerpt contradicts it.
  Observed: 2 false DOES-NOT-SUPPORT verdicts in iter-122's first adjudication.
- **false GREEN** — the corpus states what *was* true; the stale excerpt **confirms** it. Not observed,
  and structurally the *more* likely of the two, because **the corpus and the clone fell behind
  together**. A corpus written against `storage @ 4ce8ece` and adjudicated against `storage @ 4ce8ece`
  agrees with itself.

Deciding which case a pair is in requires adjudicating the claim. `F4` is the standing statement that no
regex in this family reaches that. So the module now says *substrate-dependent, direction undetermined*
and names the pairs, which is what it can support.

**A retraction inherits every weakness of the claim it retracts** (live rule). This one is text, and the
arms pinning it are text assertions — they pin a **spelling**, not a **property** (`§5` r70/71). Stated
in the fence's own docstring rather than left for a later pass to find: the defect *was* a spelling, and
the stronger arm would need the adjudicator (4) says does not exist.

## `D-M257x-198-3` — the verdict-scope statement is DERIVED, after the asserted version was false

First draft: *"the exit code grades the tier-2 ratchet, computed from the corpus text and not from any
clone's contents."* Checked against `census`: every tier-2 candidate passes through
`has_subject_token(sent, _live_names(root, clones_root))`, and `_live_names` reads clone directory names
plus `platform/repos.yml` and `platform/docker-compose.yml`. **The exit code is substrate-dependent.**

Repaired by making it a condition rather than a sentence: `NAME_SOURCE_FILES` declares the two files,
`drifted_files` says whether either moved, and the report prints one of two branches — the ⚠ branch
("the verdict below is substrate-dependent too; this is the state in which the guard would be stating
its own invalidity") or the clear branch, which names the check it ran. Today it is the clear branch:
`platform` is not stale, and directory names do not move with commits.

The near-miss is the point. The sentence was written **to be careful about exactly this**, in an iter
whose subject is instruments that misdescribe themselves, and it was wrong on the first attempt. Prose
about what a program reads is a claim like any other.

## `D-M257x-198-4` — the exit code is NOT changed, and that is a choice with a reason

The batch's rule is *an instrument that states its own invalidity must not exit 0*. Applied here it
would mean: exposure > 0 → non-zero exit. Not taken:

1. **It would be the wrong rule for this instrument.** The exposure invalidates **tier-1 excerpts**;
   the exit code grades the **tier-2 ratchet**, and `D-M257x-198-3`'s derived check says the ratchet's
   inputs have not drifted. Failing the ratchet for a tier-1 condition is a category error, and it is the
   `§9` trap harden pass 45 documented — *a good repair can destroy the proof the instrument fires*.
2. **This milestone may not fix the cause.** The clone set belongs to a live demo stack (`substrate_of`'s
   docstring). A RED whose only remedy is out of scope trains the operator to ignore it — the same
   argument the module already makes at `:846-853` about `--update-baseline`.

What replaces it is disclosure with a **named condition**, so the day the ratchet's own inputs drift, the
report says so in its own words rather than leaving a reader to infer it.

## `D-M257x-198-5` — the fence is offline and synthetic; no arm depends on the demo clone set

Every arm in `test_claim_census_substrate_m257x.py` builds its own stale clone with `git init` + a bare
origin + a `reset --hard HEAD~1`, so the battery runs on a machine with no `stack-demo/` at all — and,
more importantly, its result does not change when someone updates the clones.

Deliberate: an arm asserting *"19 pairs are exposed"* against the live tree would be **green today and
meaningless tomorrow**, and would go RED the moment somebody did the right thing and pulled. The live
number belongs in the report (where it is a reading), not in a test (where it would be a frozen literal —
`SURVEY-M257x-h45-…`'s class). What the tests pin is that each of the three buckets **can be reached**,
which is the property that makes today's 19 believable.
