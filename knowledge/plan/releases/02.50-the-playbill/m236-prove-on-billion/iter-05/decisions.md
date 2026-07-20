# iter-05 — Decisions

## D1 — root cause: the manager route's last segment is a MEMBERSHIP id

`InsightsJobSimulationBySessions(ctx, organizationID, jobSimulationID, membershipsID uuid.UUID, options)`
calls `r.app.OrgManager.GetMembership(ctx, membershipsID)`. M233's projection put `owner.UserID` there.
`GetMembership` on a user id → `ent: membership not found` → the **whole query** errors to `data: null`.

Both symptoms fall out of that one failure:
- the player's name is read from that query's payload → renders **`undefined undefined`**;
- the attempts table is that query's payload → renders **`No data`**.

**Why it survived a milestone:** the page's *header* is served by a **different** query (the sim
definition), so it renders correctly — real sim title, real "2 skills measured". The page therefore looks
populated. Any check that asked "does this page have content?" passed it. Only a check that asked "does the
attempts **table** have rows?" could catch it — which is why iter-04's `manager-dashboard` shape asserts on
the table and explicitly treats the header as chrome.

Fixed in tooling (zero platform edits): membership ids are already deterministic from the same
`(prefix, index)` the projection holds (`deterministicUUID("<prefix>:membership:<index>")`, `users.go:199`),
so `ownerSlot.MembershipID` is computed identically and both projections emit it. The derived id matched
the live-queried one exactly — a useful confirmation that the seeder's determinism contract holds across
the seed/export boundary.

## D2 — the test asserted the defective contract (this is why it shipped)

`content_manifest_test.go` asserted:

> `session %q manager path must end in the owner user-id %q`

That is the bug, written down as the expected behaviour and passing green. The projection was consistent
with its test; the test was consistent with an assumption nobody had checked against the resolver.

Corrected to assert the **membership** id, plus an explicit negative guard that fails if a user id ever
reappears in that position. The negative guard is deliberate: the positive assertion alone would be
satisfied by a future refactor that reverts the field, if the derivation happened to change shape.

**Generalizable:** a test that encodes an interface contract nobody verified end-to-end is a hypothesis
wearing a test's clothes. It converts an unchecked assumption into apparent evidence — which is strictly
worse than having no test, because it stops the question being asked. Pair such assertions with one live
verification (as iter-04's calibration step does for renders).

## D3 — the "pre-existing demo-wide defect" conclusion was premature and is retracted

Step 1 probed a **hero's** drill-down (Maya Chen) and saw the identical failure, which pointed hard at a
demo-wide platform defect — a conclusion reinforced by `coverage.spec.ts`'s manager sample-rule comment
recording the very same symptom as normal (*"the leaf renders only the dashboard tab chrome, textLen ~70"*).

**It was wrong.** The hero URL was constructed by hand from `owner_id` — i.e. with a **user id**, the same
mistake the manifest was making. The probe reproduced my own bug and I read it as independent corroboration.

Recorded because the failure mode is subtle and self-confirming: a hand-built probe that shares the
suspected defect will always "confirm" it. **A probe intended to discriminate between two hypotheses must
not be constructed from the artifact under suspicion.** This is the same class of error as iter-04's
`skipPaths` finding — an observation promoted to a fact about the platform — arrived at from the opposite
direction.

Whether the *real* manager UI navigation (which presumably passes the membership id correctly) was ever
broken is **not established** and is not claimed anywhere in this milestone's record.
