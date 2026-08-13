**Type:** tik, under `TOK-05`. Answers `CHECK-M257x-iter92-fenced-claim-restatements`.

# iter-93 — fence the HEDGE, not the sentence

## Why a fence rather than another repair

iter-92 repaired six restatements of one claim by hand, and **the repair leaked twice while doing it** —
`repair_leak_guard` RED on the iter commit, RED again on the first repair, GREEN only on the second. That is
the measured case for TOK-05 read literally: hand-repair of a claim does not hold; the predicate under it
has to be made unrepresentable.

The predicate: **a claim about a repo in no clone set must say that it is not a measurement.**

## What landed

`stack-core/unreadable_repo_claim_guard.py` — every corpus mention of a `module.*_euwest1` construct must
carry an unmeasurable marker **in its own paragraph**. Those modules are declared in
`infrastructure/terraform/production/services.tf`, and `infrastructure` has never been in any clone set.

Registered in `guard_family.py` — mandatory, not optional: the family's reconciliation is bidirectional, so
a guard on disk with no invocation entry makes the whole family exit 2. **Family is now 17 members.**

Live reading: `all 4 module.*_euwest1 mention(s) are marked unmeasurable (infrastructure is in no clone set)`.

Four design decisions, each tested rather than asserted:

- **The marker is a SET OF PHRASES, not a mandated token** — a fence requiring a magic string teaches people
  to type the string. The property is *that the reader is told*.
- **The scope is the PARAGRAPH** — a marker three screens away would launder a flat assertion; a fixed
  ±1-line window would false-RED the real corpus, since every one of these claims is a wrapped blockquote.
  Both directions tested.
- **The guard re-measures its own premise and RETIRES ITSELF** — if an `infrastructure` clone ever appears,
  it prints *PREMISE LIFTED — go and MEASURE those declarations, then retire this fence.* §8 rule 3 turned
  on a fence's own premise: a guard still demanding a hedge after the hedge became unnecessary would be
  pinning the current shape of our ignorance.
- **Anti-vacuity: no mentions ⇒ exit 2**, never green — the rung that stops it rotting the day the modules
  are renamed.

## The skip that proved the anti-vacuity rung necessary, inside the iter that wrote it

The live-corpus control was first written with a hardcoded `parents[3]` — which is `.agentspace/` in the
authoring copy. It **silently SKIPPED**, and the suite printed `OK (skipped=1)`. It now walks up to whichever
ancestor owns `corpus/architecture`, correct from both the authoring copy and a per-stack clone.

**A check that skips reads exactly like a check that passes — including when it is the check on the guard
that exists to say so.** Ninth consecutive iteration in which the author of a rule broke it while writing
the thing the rule governs.

## Close — 2026-08-05

**Outcome:** the iter-92 class is now unrepresentable rather than repaired — a tree-wide fence over the
*hedge*, with the marker set, the paragraph scope, the self-retiring premise and the anti-vacuity rung each
covered by a test, and the family grown to 17 members.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5, unchanged.** No reading taken.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (4 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-93-1 … D-M257x-93-5
**Side-deliverables:** none.
**Routes carried forward:** all open CHECKs from iters 90–92 remain; `CHECK-M257x-iter92-fenced-claim-restatements` is CLOSED by this fence for the `module.*_euwest1` case, and the *general* case (any hedge, any subject) is explicitly NOT closed — routed as `CHECK-M257x-iter93-general-hedge-fence`.
**Lessons:**
- **When a fenced document hedges because the evidence is unreachable, the hedge is part of the claim** —
  fence the hedge tree-wide, not the sentence in the one document that carries it. Recorded in the protocol.
- **A fence's own premise is a claim; measure it every run and let the fence retire itself.**
