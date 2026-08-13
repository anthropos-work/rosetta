# iter-159 — decisions

## `D-M257x-159-1` — the discriminator is the HAYSTACK, not the needle

The naive predicate (iter-155: 2,854 candidates / 766 expression-shaped over 109 files) asked what was
being **asserted**. A needle cannot carry the answer: `"directus"` is a domain value in one test and a
copied fragment of source in another, and the string is identical.

The property is in what is being **searched**. An assertion is a candidate iff its haystack transitively
derives from reading a file that is **git-tracked**, carries a **source extension**, and has **not passed
through a parser**. Each clause is load-bearing and each is RED-proofed by mutation (4 failures / 1
failure out of 19).

Clause 3 is the one worth naming: `§8`'s prescribed repair is *assert against a parsed construct, never a
whole-file substring*, and iter-158 applied it. Without clause 3 the instrument flags that repair —
**a census that condemns the fix it exists to produce**.

## `D-M257x-159-2` — grade a census at the grain of its claim

**The first proof of this instrument was a false refutation, and the instrument was not what was wrong.**

Graded by *file-level candidate count before vs after the repair commit*, discrimination read **1 of 4**,
printing *"measures the file, not the pin"*. That reading was an artifact of the measurement: the repairs
each removed **one** pinned assertion from files carrying 9, 51 and 56 legitimate source-text assertions,
so the count cannot move meaningfully. Re-graded at the grain the claim is actually about — *was the
assertion the repair DELETED among the lines the predicate flagged?* — the same predicate on the same
data scores **4 of 4**, naming L410, L252, L763+766 and L360.

**A coarse measurement had refuted a correct instrument**, and one more step would have recorded that
refutation. This milestone has repeatedly found instruments that could not fail; this is the mirror
image — an instrument graded by a test too blunt to see it succeed.

## `D-M257x-159-3` — the class is TWO classes, and the second one is unfenced

The labeled set's two recall misses were investigated at source rather than absorbed into a percentage.
Neither reads a source file at all, so no haystack clause can ever see them. The class splits by **where
the spelling lives**:

| sub-signature | instances | instrument |
|---|---|---|
| **(a) in the haystack** — searches checked-in code as raw text | **4 of 7** | this census, 4/4 discrimination |
| **(b) in the value** — a hand-written literal the test SUPPLIES or EXPECTS, duplicating what the subject derives | **3 of 7** | **none** |

`§5` rule 71's prescribed structural repair — *derive the expectation from the same source the code
derives from* — is aimed squarely at **(b)**. The rule named the repair for the half nothing was
enumerating.

**The measured recall (4/6 = 67 %) is reported on the pre-registered denominator and NOT retro-fitted.**
Re-declaring the two misses "blind" after seeing them would print 4/4. The taxonomy is the finding; the
inflated number would have hidden it.

## `D-M257x-159-4` — an instrument is not a family member until it is at zero

A `FENCE_KIND = "census"` declaration was written and `repair_postcondition.py` refused the entire
registry: *"guard(s) declare no legal FENCE_KIND … spelling_pin_census.py (declares 'census'; legal:
('postcondition', 'standalone'))"*. **That is iter-157's declaration-driven registry failing closed on an
unknown kind, unplanned and in production — a live RED-proof of the previous iter**, recorded as a
side-deliverable.

The declaration was removed **on its merits rather than by widening `LEGAL_KINDS`**. A family member is a
fence that is **at zero and kept there**; this census stands at 961. Enrolling a permanently-RED member
would train readers to skip the family — iter-155's *"a warning nobody can act on trains readers to skip
real ones"*, one level up. The promotion path (sweep to zero → declare exemptions → add the kind →
declare it) is written into the module docstring in that order, so the next reader does not re-litigate it.
