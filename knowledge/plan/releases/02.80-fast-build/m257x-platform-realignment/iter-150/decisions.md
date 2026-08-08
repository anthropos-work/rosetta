# iter-150 — decisions

## `D-M257x-150-1` — a declared partition can still have a DERIVED completeness

`blocking_state_guard` splits the iteration protocol's exit enum into `BLOCKING_FIELDS` (routes out of the
loop) and `NON_BLOCKING_FIELDS` (ends a session, not the milestone). The comment claimed the split was
*"derived from the iteration protocol's own Phase-5 grading"*; both tuples were hand-typed.

**The obvious repair — make it derive — is wrong.** Which side a condition falls on is a judgement about
what that condition *means*; no parse can make it, and a mechanism that pretended to would be the false
claim again with machinery attached.

**Decision: split the claim into its decidable and undecidable halves.** The partition stays declared and
the comment says so, with the reason. The **completeness** of the partition is derived: `run()` already
computes `seen`, the set of fields actually graded across the milestone — the only place in this
repository where the protocol's enum is written out field by field — so subtracting both tuples from it
yields anything neither classifies.

Reported as a **finding**, not a `could-not-run`: an unclassified field means the protocol moved, and
refusing to run would suppress the blocking gradings the guard exists to surface. The finding names the
first iter that graded it and how many do, so it is actionable without re-deriving.

**Why it is not hypothetical.** `budget-exhausted` entered the exit enum on **2026-08-06** — three
sessions had reported a clean budget stop, the enum had no value for it, each was told to emit
`user-blocker`, and each was flagged as a mis-grade. Someone hand-added the new name here. Had they not,
an entire exit class would have passed through this guard unclassified, and "treated as non-blocking by
omission" is indistinguishable from "decided to be safe" once it is in the output.

## `D-M257x-150-2` — hand-grade a keyword census before publishing its number

The parse-level census returned **30** literal constants whose comment block contains a derivation word.
Graded individually: **9** make a self-directed claim, and **1** is a defect. The other 21 use the word
for something else entirely — most often that a *different* module derives FROM the constant (all nine
`FENCE_KIND` declarations read *"read STATICALLY by `repair_postcondition.py` to derive the fence
registry"*), or that the constant is a fixture for a derivation under test, or that the comment discusses
derivation as the design choice being rejected, or that the sentence is about a value defined nearby.

**Decision: publish the graded number, and publish the raw one beside it as the instrument's reading, never as the finding.**
30 → 9 → 1 is a 30× difference between what a token count would have reported and what is there. This
milestone has already withdrawn one number in place for exactly that error (iter-138's *"127 rotted pins
/ 57.2 %"*, refuted at iter-139 on a pre-registered sample and withdrawn at all three publishing sites);
the difference here is that the grading happened before publication.

## `D-M257x-150-3` — narrowing an existing control for a new mechanism is disclosed in-line

`test_MUT_shrinking_BLOCKING_FIELDS_to_user_blocker_LOSES_the_finding` asserted `res["findings"] == []`
under a deliberately-shrunk partition. With the completeness check writing to the same list, that mutation
now also — **correctly** — reports `re-scope` and `protocol-stop` as classified by neither tuple.

**Decision: scope the control to the mechanism it isolates (`f["iter"] != "(partition)"`), with the reason
written at the assertion**, and assert the `(partition)` findings in their own pair of controls rather than
letting them be absorbed anywhere. A mutation control that silently starts tolerating a second mechanism's
output is a control that has stopped isolating; the narrowing has to be legible at the line that does it.
