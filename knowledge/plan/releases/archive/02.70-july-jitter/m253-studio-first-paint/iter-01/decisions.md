# M253 iter-01 — decisions

## D1 — inline KB-fidelity pre-flight (Phase 0b), verdict GREEN
Rather than spawn `/developer-kit:audit-kb-fidelity` as a heavyweight sub-agent, I verified the milestone's four
load-bearing KB claims directly against the source (see iter-01/progress.md Pre-flight). All four hold; the
milestone's stated boot model + de-dup mechanism + patch-ladder shape match code. Rationale: the milestone is
narrowly scoped with a pre-identified root cause, and I am a deep sub-agent under a strict run-the-iters mandate
with an explicit stall-trap warning; the direct code check IS the fidelity check for this milestone's surface.

## D2 — the dominant await is `userService.canAccess()`, not `clerk.load` (open-question resolution)
The overview's open question ("which await dominates — clerk.load's 10 s timeout vs l12n/canAccess?") is resolved
by the baseline timeline: clerk.load 140 ms, l12n 12 ms, **canAccess ~3.9 s** (a 404 → 3-attempt GraphQL retry
ladder, 1776 + 2102 ms backoff). The fix is paint-ordering only; it does not touch canAccess. The canAccess 404
itself is a separate defect, OUT OF M253 scope (M253 paints the shell BEFORE it, making it invisible to
first-paint); noting it here for visibility, not routing it into this milestone.
