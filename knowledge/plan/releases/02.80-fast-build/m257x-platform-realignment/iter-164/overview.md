---
iteration_type: tik
status: closed-fixed
---

# iter-164 — two shared anchor helpers, both pinned to a spelling

**Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).
Both targets are `§5` rule 71's family — *a fence pinned to a SPELLING is not pinned to a PROPERTY* —
and both were surfaced by iter-163 **while improving the guarded code**, which is where all ~7 prior
instances of that rule came from.

## Step 0 — re-survey

iter-163 closed with three routed helper defects. Two are in scope here and both were re-measured
before any code was written:

- **`anchor_construct_guard`'s content-free clause is `_CLOSER_ONLY = [\s})\];,]*` — a set of
  CHARACTERS.** A shell script closes a block with a *word*. Probed corpus-wide: **1 instance over
  684 resolved in-range anchors** (`up-injected.sh:2494`, a bare `fi`). Small, real, and stated as
  one.
- **`anchor_subject_census` used `_block_bounds` on source files.** Two iter-163 exemptions blamed
  that helper. Re-read at source: `_block_bounds` returns a *prose* block by design, and the caller
  was wrong. **That framing is retracted in this iter, not softened.**

## Hypothesis

Replacing both spellings with the property they meant will (a) catch the shell terminator, and
(b) let mechanism absorb exemptions that are currently human declarations.

## Expected lift

Instrument correctness. No `P`/`N` reading (`§9`: UNMEASURED, not unmoved). Success = the terminator
class enumerated + repaired + fenced, and the census's **declared** exemption count falling because a
sharper predicate replaced the declarations.

**Falsifiable:** if the sharper block absorbs nothing, iter-163's exemptions were right on their
merits and the retraction is wrong.

## Phase plan

- **A** — language-aware terminator clause; repair the one instance; fence both directions.
- **B** — source-aware enclosing block; reconcile the exemption table against what mechanism now
  covers; fence both directions.

## Escalation conditions

- A repair needs the surrounding CLAIM adjudicated → **tripwire**: repair the pointer, route the claim.
- A clause would duplicate an existing one → drop it rather than let two clauses own one property.

## Acceptable close-no-lift outcomes

The terminator class is 1 and the block change absorbs nothing — the measurement stands and the
exemptions are ratified rather than removed.
