# iter-116 decisions

## `D-M257x-116-1` — the guard family is RED at this reading's open, and the RED is FALSE

`platform_predicate_guard` check G10 reports `corpus/services/sentinel.md:5` as declaring the wrong
compose-service count, grading it at `d11a403` (8 file-local / 10 effective). **The sentence names
`0c91421d`, where the pair is `(5, 7)` — measured with the guard's own `compose_counts_at` helper — and
all five of its line anchors resolve. The corpus claim is TRUE at the ref it names.**

The defect is in the guard: `sentinel.md:5` is one long wrapped paragraph carrying **two** platform refs,
and G10 takes `_REF_PINNED.search(cell)` — the **first** match in the window — so it dates the claim by
whichever ref appears earliest rather than by the ref the claim names. That is §5 rule 33 violated by the
guard that enforces it.

**Routed as `FIX-M257x-iter116-predicate-guard-takes-first-pin-in-block`. NOT repaired here** — no repair
is taken inside a measuring pass (`pre-registration.md` binding condition 3). Disclosed to the seats in
the addendum so a seat that notices the same site books what it measured rather than what a guard said.

**It does not block the reading, and the reason is stated rather than assumed:** this iter lands zero code
and zero corpus edits. A RED gate blocks an iter because code must not land on top of one; there is
nothing to land, and the RED's subject was opened at source and holds.

**Second guard in two iters caught resolving a claim against the wrong thing** — the first being
`FIX-M257x-iter115-anchor-guard-resolves-fixture-paths-live`. That is a pattern worth naming even though
neither is a corpus defect.
