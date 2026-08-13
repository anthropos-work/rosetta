# iter-67 — decisions

## `D-M257x-67-1` — G7: membership is decidable in the table construct

*"Service X is in selection Y"* is a predicate with a derivable legal set —
`compose.beyond_floor(tok)` — and until this iter nothing checked it. G1 asks whether a documented
token is legal and selects *something*; G3 checks the default's *count*. A profile-reference row could
name the correct profile and the wrong service list for as long as anyone cared to leave it.

Three design choices, each derived rather than tuned:

- **The services column is found by its header**, in any position — the same move iter-63 made for the
  profile column, and for the same reason: position is layout, the header is identity.
- **Both directions are reported.** MISSING (the profile starts it, the row omits it) and NOT STARTED
  (the row names it, the profile does not start it). A membership claim can be wrong by addition as
  easily as by omission — iter-66's own RPC-edge inventory was wrong both ways at once.
- **A prose cell is UNREACHED, not an empty claim.** *"all backend services"* yields no service tokens.
  Treating that as "claims nothing, therefore agrees" is the check-that-skips-wearing-the-voice-of-a-
  check-that-passes shape this milestone has found repeatedly.

**Live GREEN at 12 checked rows of 22.** iter-62 repaired every profile table by hand, so G7 locks a
correct state rather than catching a defect. That is the right outcome for a fence built one iteration
after its class was repaired, and reporting it as a catch would be the honesty failure the protocol
keeps naming.

## `D-M257x-67-2` — the corpus's most important profile row was invisible, defeated by its own spelling

`` | `core` *(default — `PROFILE ?= core`)* | backend, gotenberg | `` — the row that states the default
selection, in the file every agent reads — yielded **no token** to `_cell_profile_tokens`, and had
therefore been invisible to **G1** as well as to the brand-new G7. The cell strips a bare `(default)`
and backticks; this qualifier is *emphasised* and carries its **own** backticks, so neither pattern
matched.

Fixed by stripping a trailing parenthetical **with any surrounding emphasis** before the default-mark
— derived from the format, not from this one spelling. Rows checked **10 → 12**, profile sites
**91 → 94**.

**Third instance this milestone of the same lesson**: iter-61 (the noun-phrase construct), iter-63
(the profile column's position), and now the qualifier's spelling. **Each time, the fence was GREEN
and the reason was that it could not see the site.** A GREEN reading is a claim about the fence's
reach as much as about the corpus.

## `D-M257x-67-3` — the mutation battery caught a weak TEST, not a weak rule

Of three mutants, two confirmed the assertion. The third — reverting the qualifier strip — produced
**zero** failures, because `test_a_default_qualifier_does_not_hide_the_row` asserted
`membership_rows >= 1`, which passes on the base fixture corpus alone. **The test could not fail.**

Re-written to assert the **delta** (`base + 1`), it goes RED under the same mutant.

§8 rule 2 says mutation-verify the fixtures too. This is what that buys: the battery's most valuable
result was not "the rule works" but "one of your tests does not test anything." A rule can be audited
by reading it; an unfalsifiable test looks exactly like a passing one.
