# iter-53 — decisions

## `D-M257x-53-1` — the union is 46 as run, 35 re-graded, and BOTH numbers are published

The as-run grading is what the seats actually applied and is therefore the honest primary. The canonical
re-grade is published beside it because it is the only number comparable to readings #9/#10. **Neither is
suppressed and neither is presented as the other.** The re-grade's weakness — performed by the orchestrator,
not blind, with both readings in hand — is stated at the point of use in `variance.md`, not in a footnote.

## `D-M257x-53-2` — the matching rule was fixed BEFORE the matching, and its two hard calls are recorded

*Two findings match iff they assert the same defect about the same passage — same anchor and same predicate.*
Fixing it first matters because `m` is the denominator of the whole estimate and a permissive matcher deflates
`N̂`. The two calls it did not decide cleanly (`external_services.md:616/618` → matched;
`:634` twice → not matched) are written into `blocker-ledger.md` so a later reader can re-adjudicate them
rather than inherit them. Flipping both moves `N̂` by roughly ±5, which does not change any conclusion here.

## `D-M257x-53-3` — the instrument is now a COMMITTED file, and this is the iteration's most durable output

`knowledge/plan/.../m257x-platform-realignment/instrument/briefing-canonical-iter41.md` is the recovered
canonical briefing, committed. `briefing-iter53-AS-RUN.md` is preserved beside it as drift evidence, marked
not-for-reuse. Every future reading uses the canonical file **verbatim**; a reading that changes the
instrument now changes a diff.

**This closes a hole that had been open since iter-41 and was invisible by construction** — the file was
git-ignored, so no commit, no review and no fence could ever have seen it move. Recorded in the protocol as
§5 rule 25, and rule 22's false "never touched a knob" sentence corrected in place.

## `D-M257x-53-4` — NO repair was performed, deliberately, and the fixture is intact

TOK-03 move 1 is *read, read, repair* — the repair belongs to iter-54. Repairing here would have spent the
46-finding answer key and the 14 raw seat reports, which are this milestone's only fixture with a known
answer for the tree at `0e35b1a` (§5 rule 21's perishability clause). The two `corpus/ops/platform-alignment.md`
edits in this commit are **protocol evolution**, mandated by the iteration skill and made outside the 40-file
audit partition; they repair no audited claim.

## `D-M257x-53-5` — the repair target for iter-54 is NOT decided here

46 or 35 is a question about what clause 5 counts, and the canonical instrument calls the 11 in between
MINOR — which *"YELLOW with 0 blockers"* admits. Choosing on the orchestrator's own authority is precisely
the single-seat adjudication that put a false `32` into the corpus at iter-49 and a false `31→32` repair into
it at iter-52. It is surfaced as a user-blocker instead.
