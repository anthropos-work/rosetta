# WITHDRAWN — `av-cycle{1,2,3}.json` do not support gate clause 1

**Withdrawn at M257x iter-17, 2026-08-01.** The three files remain on disk as the record of what was
measured; they are **not** evidence for the clause and must not be cited as such.

## What they say

Three `{"green":true,"warnings":0}` verdicts at `11:43:02Z`, `11:53:04Z`, `12:03:33Z` (iter-14), each after
a verified `demo-down 1 --purge`, each on distinct monotonic timestamps. **All of that is still true.** They
were not fabricated and the three-consecutive-cold-cycles procedure was executed honestly.

## Why they do not support the clause

The instrument could not see the failure. `autoverify`'s only Directus check counted rows in
`directus.directus_collections` — the **registry** table. It asked what the content model had *registered*,
never whether the running Directus would *serve* an item (protocol §5 rule 14: **REGISTERED is not SERVED**).
The bring-up transcript said `directus=skipped(error)` on all three cycles at the same time, and nothing
reconciled the two.

## What replaced them, and what it found

iter-17 re-ran the procedure with `probe_directus_serves_content` (harden pass 1) in the **consumed** tag —
the pin was itself stale until iter-16, which is why this could not have been done earlier. The probe picks
a non-`directus_*` collection the stack's own Postgres says holds rows, then asks the running Directus for
an item over HTTP.

**Cycle 1 result — `evidence/av-iter17-cycle1.json`, `{"green":false,"warnings":1}` at `16:03:35Z`:**

```
✗ directus-serves-content  fail: anon GET /items/task_sub_checks -> 403 — the running Directus
                                 holds the content but serves it to nobody (public-role grants never applied)
```

Same stack, same procedure, same moment — the registry check `✓ directus.directus_collections = 21` passed
in that very run. The two checks disagree because they are about different things, and the clause needed the
second one.

**Cycles 2 and 3 were deliberately not run.** The clause requires three *consecutive* green cycles; cycle 1
was red, so the outcome is already determined and two further cycles would produce ~22 minutes of data that
cannot change it. Running them to fill in a table would be the same instinct that produced these three files.
