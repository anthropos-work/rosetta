---
milestone: M257x
iter: 06
---

# iter-06 — decisions

## D-M257x-6 — for a co-written PAIR, the correct re-point is REMOVAL, with the replacement asserted in place

`platform-alignment.md` §7 rule 2 forbids re-pointing a write "to nothing", because a deleted write is
trivially satisfiable and silently re-empties the surface (the M219/M222 render-gate trap). The
pre-computed input for this iter flagged the session PAIR as a decision rather than a conclusion for
exactly that reason.

Measured: **three** seeders (`persona_write.go`, `content_stories.go`, `hiring_funnel.go`) wrote the same
rows, under the same ids, into `jobsimulation.sessions` **and** `public.job_simulation_sessions`. M257
had added the app-side half beside the legacy one instead of replacing it. Re-pointing the legacy half
would have made both steps the identical statement.

**Decision: delete the legacy half.** Rule 2 is satisfied *in situ* — the replacement is the very next
step in the same slice, which is a stronger form of "asserted" than adding a separate check would be.
The two seeders with no app-side half (`jobsim_sessions.go`, `ai_readiness_funnel.go`) were genuinely
re-pointed instead.

**Corollary that cost real work:** cross-schema writes have no FK, so their order was free. Once both
land in `public` the FK is real and the order is load-bearing. `public.job_simulation_sessions` had to
move to the FRONT of both flush slices, and `actors`/`interactions` had to move ABOVE it in
`resetTables`. Recorded in the protocol (§7 rule 2's new paragraph).

## D-M257x-7 — the PROD-read paths keep naming `jobsimulation`, and that is correct, not an oversight

`cmd/content-capture/main.go` (12 queries) and `contentsession/sourcing.go` (4) read the **production**
database read-only at authoring time. `D-M257x-3` established that the map needs two states per row, and
this is the case that proves it: on a fresh local stack the `jobsimulation` schema is **never created**,
while in production it is **still present**, pending platform M710. The same relation name is dead in one
environment and live in the other.

Re-pointing these at `public` would break content capture against prod today in exchange for nothing.
They stay, and they are deliberately outside the new fence's scored sections. `sourcing_test.go`'s
assertion was rewritten by the mechanical sweep and reverted by hand.

**Tracked:** when platform M710 executes, these are the sites that move. That is the first item the next
detection sweep (§4 cadence, "watch the named next fold") should check.

## D-M257x-8 — the write-target fence scores `stack-seeding` only, and says why in the file

The obvious generalisation was to score `stack-snapshot` too, since it holds the one remaining
transitional write (`cms.similarities`). It was tried and rejected: `stack-snapshot`'s replay surfaces
legitimately name schemas rext's migrate step does not and should not create — `directus` (created by
Directus's own bootstrap) and `ref` (rext-owned). Scoring them makes the fence permanently RED, and the
natural repair — an allow-list — is protocol **Trap A** in miniature: you tune the check until it also
stops catching the thing it was built for.

The narrow fence is the honest one. `cms.similarities` is not left unguarded: the replay already fails
LOUD with a named precondition (`the stack's "cms" schema is missing/empty`, rc=4), which is a working
signal, and the re-point is tracked as `REPOINT-M257x-cms-similarity-writes`. The exclusion and its
reasoning are written into the fence's own source so the next reader does not re-litigate it.
