# iter-207 — decisions

## `D-M257x-207-1` — the exclusion is SIZED, not re-litigated

The three literal censuses drop every `test_*.py` at file grain. The rationale for the exclusion — a
test's `assertEqual(len(rows), 7)` is an expectation on purpose — is **correct for the asserted half**
and this iter does not touch it. Changing an assertion is a different kind of work with a different
blast radius.

What the iter changes is that the exclusion now has a **size** and an **enumeration**: 462 rows / 314
`standing` across 79 modules in five sections, re-derived on every run. `§5`: *a CORRECT exclusion is
still a defect while it is silent.*

**Not chosen:** removing `path.name.startswith("test_")` from the three siblings. That would fold the
two populations together and destroy the one distinction the exclusion was built to draw. Two censuses
over two named populations keeps both facts.

## `D-M257x-207-2` — three historical figures repaired by SELF-DATING, not by re-pinning

Rows 4–6 of the pre-registered subset are readings taken at a past iter and repaired since. Two repairs
were available:

1. **Re-pin** — update the numbers to today's tree. Rejected: it re-arms the trap (iter-206's rule,
   *de-literalise beats re-pin*), and today's number is stale tomorrow.
2. **Self-date** — put the provenance marker inside the sentence, beside the figure. Chosen.

This is `§5` iter-111, *a verdict with a GRAMMAR states its provenance INSIDE the payload*, and iter-98,
*write the retraction in the vocabulary the fence enumerates*. It also fixes the classification
mechanically — all three now grade `dated` — which is a **consequence** of the repair being right, not
the reason for it. The distinction matters: moving a marker to change a classifier's verdict *without*
the sentence genuinely being historical would be gaming the instrument.

Row 4 additionally had a **live** error (25/23 against 27/27). The repair states the live pair is
`RP.discover_fences()` and lets the arms derive it, rather than writing 27 into the prose.

## `D-M257x-207-3` — the ceiling was taken from the census, never from the probe

This iter's pre-flight probe read docstrings and comments only and sized the fourth site-kind at **301**.
The census — which reads *every* string literal outside a `print(...)`, and a test module's assertion
messages are full of them — reads **462**. Taking the ceiling from the probe would have reproduced, in
the sizing of the fourth site-kind, the exact narrowing that made the fourth site-kind necessary: it is
`D-M257x-203-2` for the fourth time on this module.

The wrong first draft is **kept as the arrow's left operand** (`301 → 462`) rather than deleted, so the
mistake is visible in the block the ratchet arm reads.

## `D-M257x-207-4` — the standing-class figure is re-derived, and the correction is APPENDED

Harden pass 50 published *"the standing class is 147, not 168."* Re-derived here at rext HEAD `5f4b779`
by unpacking `git archive HEAD` and censusing that tree with **its own** `derivation_registry.py`: the
class is **157** (docstrings 80 + comments 77, 314 distinct rows total). The ceilings shipped in that
same commit — 179 and 137 — are consistent with 314 rows and inconsistent with pass 50's 164 + 117.

Pass 50's method was sound and is quoted rather than blamed: it held iter-206's vocabulary fixed *"so
the only variable is the defect."* The failure is at the hand-off — the narrow-lens reading travelled to
the milestone level as the class's size with no qualifier attached.

**Appended as a marked correction**, never substituted, exactly as pass 50 treated iters 205/206. Third
milestone-level value for one class: **168 → 147 → 157**.
