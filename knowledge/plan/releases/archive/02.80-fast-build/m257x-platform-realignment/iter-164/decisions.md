# iter-164 — decisions

## `D-M257x-164-1` — iter-163's exemption reasons are RETRACTED, not softened

Two of iter-163's nine declared exemptions read *"`_block_bounds` closed the block 2 lines early —
a defect of the shared helper, not the doc."* Re-read at source: `_block_bounds`'s docstring says it
returns **the PROSE block** — a run of contiguous non-blank lines — and it was extracted verbatim at
iter-100 precisely so one definition of "block" serves the markdown callers. **It does exactly what
it says. `anchor_subject_census` applied a prose predicate to Go source, where a blank line inside a
function is normal.** The caller was wrong.

The repair is a caller-side `enclosing_block(tlines, n, target)` that dispatches on the target's
suffix: source files get the top-level declaration (back to a column-0 opener, forward to its
column-0 terminator); everything else keeps the prose block, untouched.

**Why the retraction matters beyond two lines of text.** A declared exemption that names the wrong
cause is worse than no exemption: it is a *routed defect* against an innocent component, and
`FIX-M257x-iter163-block-bounds-under-reaches-by-two` would have sent a future iter to widen a helper
that four other guards depend on. The route is withdrawn.

## `D-M257x-164-2` — an over-reaching ACCEPTANCE clause hides findings, and this one hid a real candidate

The interesting direction is not that the sharper block absorbed exemptions. It is that it **surfaced
a new candidate the loose block had been swallowing**: `demo-up-defaults.md:77` cites
`up-injected.sh:43`, and the prose block for a 2,700-line shell script ran from line 1 to line 154, so
*any* literal in the first 154 lines counted as "inside the block."

Every guard in this milestone has been audited for *can it fire* and, since iter-161, for *can it
still show that it fires*. **Nobody had audited an ACCEPTANCE clause for over-reach.** A too-generous
"this is fine" rule produces a green with no finding to look at, which is the same failure as a
guard that cannot fire, arriving from the opposite side. The new clause is tested in **both**
directions, including the control that the prose block would *not* have covered the case — without
which the fix is unfalsifiable.

Net effect, and the direction is the point: **declared exemptions 9 → 5.** Four human declarations
were replaced by mechanism; the fifth is net-new and graded at source.

## `D-M257x-164-3` — the terminator clause nearly shipped `;;`, which would have RENAMED a class

The first draft of `_TERMINATOR_WORD` was `(fi|esac|done|;;|end)`. `;;` is punctuation and
`_CLOSER_ONLY = [\s})\];,]*` already matches it — so the clause would have looked broader while
changing nothing, **except the class name reported for every `;;` anchor**, which four releases of
recorded verdicts spell `anchor-on-closing-delimiter`. Caught by the test that asserts the two
classes stay distinct.

The measurement is reported as it came out: **1 instance over 684 resolved in-range anchors.** A
class of one is still worth a clause when the clause is four words and the instance is real — but it
is stated as one, not generalised into a claim about shell scripts in the corpus.

**Tripwire honoured on the repair.** `verification.md:662`'s `fi` anchor is repaired to `:2714`, the
line that passes `--services` — unambiguously the subject the prose names. The *truth* of the
surrounding claim is a separate question, and its other anchor (`services.sh:43-44`, a range this
census does not read) looks suspect too. That is a third line of investigation; it is routed, not
opened.
