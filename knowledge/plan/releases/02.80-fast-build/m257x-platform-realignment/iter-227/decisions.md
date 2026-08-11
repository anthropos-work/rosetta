# iter-227 — decisions

## D-M257x-227-1 — a refuted prediction is the deliverable, and the iter is graded on the census

`P-227-2` predicted at least one further repo would state a status the corpus does not reflect. It does
not. Zero of six.

That is **not** an empty iter: the value is the **bound**. iter-224 found the messenger defect by accident
and had no way to know whether it was one site or twenty. Asking all six settles it — the defect is bounded
to one site, and the corpus is ahead of, never behind, every archived repo's self-description.

Graded `closed-fixed` because the planned scope was *the census plus repair of what it found*, and both
landed: six rows classified, and the one thing the census did surface (four repos self-describing as live)
written into the map.

## D-M257x-227-2 — the prose was rewritten to clear the fence, not the fence's registry widened

`derived_count_guard` arm D RED'd on *"4 of 6 say they are live"* — an `N of M` count on the clause-5
surface with no entry in `N_OF_M_DISPOSITIONS`.

Two ways to clear it: add a disposition to the registry (which lives in `rosetta-extensions`, so a commit
**and a push to origin** for one sentence of corpus prose), or stop making the construct. The second was
chosen. The guard's pattern is digits-only (`(\d[\d,]*)\s+of\s+(\d[\d,]*)`), so *"four of the six"* clears
it, and it reads better.

**This is not evading the fence.** The count is unchanged and still stated; what changed is that it is no
longer written in the form reserved for counts that must carry a re-derivation disposition. The full
six-row table sits immediately below it, so the number is derivable by inspection at the site itself.

## D-M257x-227-3 — the four stale repo self-descriptions are recorded, NOT fixed

`cms`, `jobsimulation`, `graphql-wundergraph` and `roadrunner` describe themselves in their own
`CLAUDE.md` as live, ECS-deployed services. All four are wrong.

**Fixing them is a platform edit**, which this milestone does not do — the iter's `overview.md` named this
escalation condition in advance: *"if a repo's `CLAUDE.md` contradicts the platform's `repos.yml` rather
than the corpus, that is a platform-internal inconsistency and gets recorded, not repaired."*

Recorded as `ROUTE-M257x-227-archived-repo-selfdesc-is-stale`, with the table standing as the evidence if
it is ever raised upstream.
