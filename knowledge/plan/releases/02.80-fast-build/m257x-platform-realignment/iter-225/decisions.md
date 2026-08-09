# iter-225 — decisions

## D-M257x-225-1 — a SIZING answer is the deliverable; a fabricated cycle would not be

The orchestrating brief is explicit: *"a cold bring-up is expensive and may not fit one iter; sizing it
honestly is a legitimate iter, fabricating it is not. Never claim a stack works on the strength of a
fence."*

This iter asked the cheapest **disqualifying** question first — *is there a measured host profile for the
only host the release is allowed to develop and test on?* — and the answer (no) settles clause 1's
attemptability without building a container.

**No cycle was run and none is claimed.** The wall-clock cost of a cold cycle on this host remains
**unmeasured**, and it must stay unmeasured rather than be scaled from `billion`'s 666.29 s or
`laptop.json`'s numbers: `laptop.json`'s own `arch_changes_the_image_size` note records the same
Dockerfile producing 2.88 GB on arm64 and 4.84 GB on x86_64, and `billion.json`'s baseline note says a
wall-clock figure *"does NOT transfer to another host."*

## D-M257x-225-2 — the refuted prediction is published WITH its mechanism

`P-225-2` predicted `buildbench` would refuse to grade on this host, citing `require_measured`. It does
not. The prediction is recorded as **REFUTED** in the close table with the reason spelled out — clause
zero fails a `None` *measurement input*, which is a dead sampler, not a missing or inapplicable profile.

The distinction matters beyond the grade: the corpus describes `require_measured` as *"the clause a fresh
host hits FIRST"*, which is true for one kind of fresh host and false for another. Publishing the refuted
prediction with its mechanism is what makes that ambiguity visible; silently restating the finding would
have hidden the doc's own imprecision.

## D-M257x-225-3 — the two instrument repairs are ROUTED, not landed

Both are real and neither is opened here:

- **A profile-vs-host identity check** in `buildbench` (`ROUTE-M257x-225-profile-vs-host-identity-check`).
- **The `role`/notes strings** in `billion.json` and `laptop.json`, which cite the superseded `D-v28-14`
  and name the retired `odysseus` as gate host
  (`ROUTE-M257x-225-hostprofile-role-strings-name-a-retired-gate-host`).

Reasons, in order: the user's redirect ranks corpus + working stack above the instruments that grade
them; each is a distinct line of investigation beyond this iter's planned census + repair scope (the
scope-creep tripwire's own example); and the second requires a `rosetta-extensions` commit **plus a push
to origin**, which is a shipping step, not an edit.

The mechanism for each is written into the close section so the next handler re-derives nothing.

## D-M257x-225-4 — `knowledge/plan/` citations are NOT re-pointed by a corpus edit

The ~40-line insertion into `build-budget.md` shifts every line below 128. **0 corpus-scoped citers of
`build-budget.md:NN` exist**; all 55 hits live in `knowledge/plan/releases/.../iter-NNN/` records.

Those are **frozen evidence of what a past iter read at a past ref** and are deliberately left alone.
Re-pinning them would rewrite history to match the present, which is the opposite of what an iter record
is for — and iter-139 already recorded one of them (`build-budget.md:394` → `:319`) as *observed rot*,
which is exactly the sort of datum a re-pin would erase.
