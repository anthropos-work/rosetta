# iter-17 — decisions

## D-M257x-17-1: clause 1's three verdicts are WITHDRAWN, not re-run

**Decision.** `evidence/av-cycle{1,2,3}.json` stay on disk with a sibling
`av-cycle1.json.WITHDRAWN.md` stating that they do not support the clause and must not be cited for it.
The milestone's clause count is corrected **2 of 5 → 1 of 5**.

**Why not delete them.** They are an accurate record of a procedure that was executed honestly — three
verified purges, three distinct monotonic timestamps, three real `green:true` reads. Deleting them would
erase the most useful part of the finding: *this is what a correctly-executed measurement with a blind
instrument looks like*. The withdrawal note is the artifact; the files are its evidence.

**Why not "partially met".** The clause is a conjunction over three cycles of a green verdict. One cycle is
red. There is no fractional reading of it, and inventing one would be the thing this iter exists to stop.

**Why correcting downward is recorded as progress.** Everything decided since iter-14 — including the
ordering of iters 15 and 16 — was decided under "clause 1 is met." A milestone is worth less with a false
2-of-5 than a true 1-of-5.

## D-M257x-17-2: cycles 2 and 3 not run

**Decision.** After cycle 1 returned `green:false`, the remaining two cold cycles were not executed.

**Why.** The clause requires three *consecutive* green cycles; a red cycle 1 determines the outcome. Two
further cycles cost ~22 minutes and could not change the verdict — they would only populate a table.

**The risk accepted.** Cycle 1 is a single observation, and a flake would be indistinguishable from a defect
on n=1. Two things make that acceptable rather than sloppy: the failure is **mechanistically explained** end
to end (bootstrap fails → grants never applied → anon 403) rather than merely observed, and the same 403 was
independently seen by iter-15 from the opposite direction. If iter-18's fix does not change the result, the
flake hypothesis returns and n=1 is where to look first.

## D-M257x-17-3: the bootstrap fix is ROUTED, not landed here

**Decision.** `FIX-M257x-iter17-directus-bootstrap-blind` and
`CHECK-M257x-iter17-setdress-verdict-contradiction` go to iter-18. iter-17 lands no source change.

**Why route.** This iter's declared scope was *measure clause 1 through an honest instrument*, and it did
that. Landing the capture-and-classify fix would be a second line; proving it needs another cold cycle
(~11 min plus teardown), and the interesting question it opens — *should a failed bootstrap stop the replay,
given the pass currently announces prod-read fallback and then replays into the local Directus anyway?* — is
a design decision that deserves its own hypothesis rather than being decided in the last hour of an iter
that set out to measure something else. Scope-creep tripwire respected.

**Why iter-18 is genuinely unblocked.** The fix's shape is already known — it is iter-16's RF-1 fix applied
to a third file (`>/dev/null 2>&1` → capture, classify, report) — and its first measurement is a single cold
cycle whose expected output is a named bootstrap error. That is a well-formed next iter, not a vague route.

**Deliberately NOT concluded:** that the bootstrap failure and the replay-anyway behaviour are one bug. It
is the strongest hypothesis and it is written down as one, but the diagnosis the fix produces is what should
decide it. Asserting it now would be this milestone's own dominant class — a claim reported without being
measured — committed in the write-up of an iter about exactly that.
