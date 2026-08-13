# iter-143 — decisions

## `D-M257x-143-1` — an AUDIT is a predicate too, and it needs a control that is not another reading

**THE HEADLINE.** iter-142 established *audit the predicate before the repair, not after the
publication* (`D-M257x-142-1`), and it is right: iter-138 published first and was audited to 0-for-12;
iter-142 audited first and repaired 44 of 44. **iter-143 followed that instruction exactly and the
audit was still wrong on 9 of 92.**

What caught it was **not** a second reading. It was pushing the reader's own TRUE set through the
guard's `classify()` — the machine that was going to consume the verdict anyway — and finding nine
`anchor-out-of-range` verdicts against files whose line counts make the citation impossible.

> **A reading is evidence about a corpus. It is not evidence about itself.**

The reader's error was **systematic, single-mechanism, and one-directional**: every one of the nine was
a bare file *mention* sitting nearer to the anchor than the *qualified citation* actually governing it,
and every one of them inflated the predicate's apparent precision (90.2 % → 74.5 %, **15.7 points**).
The sites a reader mis-grades are the ones that look right, which is why the bias runs **towards
shipping**. Landed as `§5` rule 65.

## `D-M257x-143-2` — head inference over the `(bare)` bucket is REFUSED, with its number

Built, hand-read at 100 %, measured: **57.6 %** raw, **77.3 %** at best (32.1 % recall). Not fence
quality. **Not shipped.**

This is an **answer** to `FIX-M257x-iter138-anchor-rot-fence`'s re-specification (*head resolution
first*), not a failure to deliver on it. The route asked whether the orphan bucket can be resolved;
the measured answer is **not mechanically, at fence quality, on this construct** — and the route is
re-specified accordingly rather than left open as though untried.

The alternative that would have "worked" — a numeric cut at `n < 3000`, which separates the population
perfectly — was **declined as a tuned constant on a 92-site denominator**. Two such constants are
already routed open from iter-142 (`-path-arm-window`, `-tier-b-underflag`). A third is a pattern.

## `D-M257x-143-3` — count the LOUD and the SILENT failure modes separately

The guard's own comment justified refusing the bucket by naming **ports**. Measured over the 39 false
admits: **21 ports** (loud — out-of-range, shows as a RED) and **16 wrong-head** (silent — a real line
anchor booked against a file the sentence never named, which can land on a real construct and PASS).

The comment was not false; it named the half that announces itself and omitted the half that does not.
**When you decline a widening, decline it for the hazard that cannot be seen when it is wrong** — a
justification that only covers the visible failure mode will be re-litigated by the next person who
notices the visible mode is manageable. Retracted in place, as an explanation, with the numbers.

## `D-M257x-143-4` — a reach gain that requires no inference is a different KIND of change

`_CODE_SUFFIX` had not been measured since iter-73 chose it. **32** no-slash `name.EXT:NNN` citations
were **invisible** to `_QUALIFIED` — not unresolved, invisible. Seven suffixes admitted, three
declined.

The distinction that makes this safe, and that the iter is careful to keep visible: **a qualified
citation carries its own path, so admitting it decides nothing the prose had not already said.** The
head inference, by contrast, *decides* something the prose left open. Same guard, same axis, same
iter — opposite risk profiles, and they must not be graded by the same standard.

The declined `de` (`u422950.your-storagebox.de:23`) is kept in the source comment **as the
counter-example that prices the list**: it is a `HOST:PORT` that matches the citation shape exactly,
and admitting it would book a hostname as a file at a line number that is really a port. The list is
an allow-list of suffixes belonging to **files** rather than to **names** — that property, not its
length, is the safety argument.

## `D-M257x-143-5` — publish a census that moved with BOTH readings

This iter changed the corpus (`_CODE_SUFFIX` made citations visible; the protocol-doc edit added a
mention), so the `(bare)` census moved **384 → 380** *inside the iter*. Both readings are published.

Quoting only the closing figure would present an **intervention** as a **measurement** — the same
error class as iter-138's, one level up. The four reason counts also always sum to the `(bare)` head
the same run prints, which is the census's own arithmetic control: a reason that stops being counted
surfaces as a mismatch, never as a quiet zero. Asserted by a test on the real corpus.

## `D-M257x-143-6` — a returned tuple's ARITY is a published interface

The census shipped, first, as an eighth member of `anchor_construct_guard.run()`'s return tuple. The
guard was GREEN, its own module was 24/24, and the three suites chosen as the change-derived scope
were 106/0. **The whole-suite run then returned 31 failed**, 30 of them mine: `test_iter45_mechanical_fences`
unpacks that tuple **positionally at six call sites**, so every one raised `too many values to unpack
(expected 7)`, and the mutation battery that runs the module cascaded three more.

Two things are worth separating, because only one of them is the mistake everybody notices:

1. **The design was wrong, and the right one was already in the file.** `RESOLVE_ROUTES` and
   `NOT_CITATIONS` are module-level accumulators cleared by `run()` at entry — the same shape, for
   the same purpose, ten lines above. Widening the tuple was the lazier reach past an idiom the
   module had already established. `BARE_REFUSALS` now follows it.
2. **The scope selection was wrong in a way that felt principled.** I picked the scoped suites by
   *"what consumes the changed return value"* and named three. `test_iter45_mechanical_fences`
   consumes it too — it is the one that unpacks it positionally, which is precisely the consumption
   pattern an arity change breaks. **A change-derived scope is only as good as the derivation**, and
   mine was a recollection of imports rather than a search for call sites.

Regression tests pin both halves: `len(run(root)) == 7`, and the accumulator is asserted **per-run,
never cumulative** — an accumulator not cleared at entry doubles on a second run, and `D-M257x-143-5`'s
arithmetic control would then pass on a wrong number.

**And the meta-point, which is this iter's second instance of one pattern.** iter-142's real miss was
caught by `repair_leak_guard` — a guard on a *different axis*, run over the commit. iter-143's real
miss was caught by the *whole suite* — a check on a different axis from the one being changed. Neither
was caught by anything inside the thing being built, and in both cases everything inside it was green.
This is `D-M257x-143-1` (*an audit is a predicate too*) arriving a third time, from a third direction:
**the check that catches you is never the one you designed while making the change.**
