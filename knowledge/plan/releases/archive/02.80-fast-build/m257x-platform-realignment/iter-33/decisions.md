# iter-33 decisions

## D-M257x-33-1 — the pre-registered prediction was HALF WRONG, and the wrong half is the finding

The overview recorded two predictions before any report was read:

| prediction | outcome |
|---|---|
| total blockers **10–25** | **19** — inside the range |
| the router drop will be the **largest single cluster** | **REFUTED — 0 blockers, 2 minor captions** |

The router sweep had landed. Every large file already carried a banner, an inline prod-only fence or a
HISTORICAL block, and four of the five audit groups reported it explicitly clean without being asked to
confirm a negative.

**What the corpus actually had is a different class: derived-fact rot.** Every doc states *who is merged*
correctly — the `ServiceDocStatusFence` from harden pass 6 holds — while naming **tables the platform
dropped or renamed**, **packages that were split out**, and **"routed forward to M219/M220" items for work
that already shipped. None of it uses merged/live/gone vocabulary, so no term-scoped sweep could have
reached any of it. That is the same shape as the studio-room archetype, generalised: **the status layer is
now fenced, and the layer underneath it is not.**

This is recorded as a *correction to the milestone's model of its own corpus*, not as a lucky miss. The
prediction was written down precisely so that being wrong about it would have to be said out loud.

## D-M257x-33-2 — the blocker bar was fixed BEFORE reading, and it excluded a lot

> **BLOCKER = false at platform origin HEAD *and* acting on it would misdirect real work.**

Explicitly not blockers: claims true at HEAD however stale they feel; fenced historical or prod-only
content; a merged-service doc opening with its standing ⚠ banner; a service with no doc at all. This bar is
what makes "19" a number rather than an opinion, and it is why 52 further findings were graded **minor**
and left. The gate's own wording — *"GREEN, or YELLOW with 0 blockers"* — means the minors do not block it.

## D-M257x-33-3 — a corrective sweep is audited like any other text

iter-22's precedent is that a sweep introduced a defect its own pass could not see. **This iteration
reproduced that three times, by hand, within minutes of applying the sweep:**

1. `jobsimulation.md` — an anchor that began mid-sentence left *"The manager view does / reads the same
   table."*
2. `clerkenstein.md` — *"all five surfaces at 100% is false on **two** counts"* became false itself once
   one of the two counts was fixed. It is now one.
3. `clerkenstein.md` — the fake BAPI *"fabricates"* the wrong org eid, present tense, in a paragraph now
   headed ✅ RESOLVED.

None of the three would have been caught by re-running the guards, and all three were introduced by the
correction rather than surviving it. So the fixes were then handed to a **second, adversarial audit pass**
over the 13 changed files whose brief is explicitly *"catch any NEW false claim the correction itself
introduced"* — the same discipline the milestone applies to a fence (a mutant must die) applied to prose.

**Clause 5 is not graded on "19 found, 19 fixed."** That would be a probe satisfying itself (§5 rule 7).
It is graded on a re-measurement.

## D-M257x-33-4 — the sweep harness is two-phase, and that was a real bug not a precaution

The iter-22 harness shape (enumerated `(file, old, new)`; `old` must match **exactly once**; 0 and 2+ both
fail loudly; deliberately non-idempotent) was reused. **The first cut of it had a genuine defect:** it
wrote each file as it finished, guarded by a *global* findings list, so a broken anchor in the last file
could leave the first file already written and the tree half-swept. Fixed to validate every anchor in
every file first and write only if all pass. Caught by reading the harness, not by running it — it would
have been silent on a clean run.

## D-M257x-33-5 — target substitution, on a measured justification

The hand-off named `CHECK-M257x-iter27-drilldown-target-coupling` as the next target. Substituted for
clause 5 under the same TOK, because **iter-32 measured the binding suite at 4 min 50 s rather than the
~40 min every hand-off had assumed**: a clause-2 fix iter can now afford to close with its own binding
read, so it no longer needs a session to itself, while clause 5 needs a long serial read that does. The
drilldown target is carried forward untouched.

## D-M257x-33-6 — no rext change, no re-pin

Clause 5 is entirely a rosetta-corpus clause. No rext runtime source was touched; the pin stays
`fast-build-m257x-iter-31b`.
