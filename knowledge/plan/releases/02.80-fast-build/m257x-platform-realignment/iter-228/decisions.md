# iter-228 — decisions

## D-M257x-228-1 — the three ACTIVE clones are recorded as behind, and deliberately NOT advanced

`app` is 28 commits behind `origin/main`, `next-web-app` 12, `ant-academy` 9.

iter-224 fast-forwarded the four **archived** repos without hesitation: nothing builds from them, so the
only effect was to make the graded substrate match the cited substrate. **These three are different.**
`app` and `next-web-app` are build sources for a demo stack; advancing them changes what a bring-up
produces, and `ant-academy` runs natively against the same world.

That decision already has an owner: `ROUTE-M257x-222-pin-advance-needs-a-reproof`, which gates the pin
advance behind gate clause 1's three cold cycles. iter-223 established that the 23 demopatch anchors
survive an advance — a necessary condition, not a sufficient one. **This iter supplies the missing
numbers to that route and changes nothing else.**

Advancing them here would have been a silent, unreproved change to the stack's build inputs, made by an
iter whose planned scope was a citation census.

## D-M257x-228-2 — the census's own wrong-operand first pass is published

The first pass compared each clone's **local HEAD** to the corpus's sha vocabulary and found all 13 cited.
The question was about the **origin tip**, and for a clone that is behind those are different commits.

The wrong table is described in the close section rather than replaced silently, because its failure mode
is the dangerous kind: it returns a **clean bill of health**. *"All 13 tips are cited"* would have been
published as reassurance, derived from a column that could not answer the question — the same shape as
`clone_drift_guard`'s own limitation, reproduced by hand one hour after measuring it.

## D-M257x-228-3 — the moving-label repair keeps the original claim scoped to its measured ref

`CLAUDE.md` read *"origin/main is now `ad9f3c49`, and all five anchors resolve identically there."*

The naive repair — swap the sha — would have silently re-attributed *"all five anchors resolve
identically"* to `3eaadae6`, a ref at which **that was never measured**. The correction instead pins the
original claim to the ref it was measured at (*"origin/main was `ad9f3c49` when this was written"*) and
states the current tip separately.

**A stale ref is repaired by adding the current one, never by moving the old claim onto it** — the
measurement travelled with the old sha, and it does not transfer.
