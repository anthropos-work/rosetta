# iter-255 — decisions

## `D-M257x-255-1` — declaring a precondition WEAKENS every mutation proof that depends on that test firing

The finding of the iter, and it was made by the tooling rather than by me.

`test_m257x_mechanical_fences_mutation_battery` mutates a fence and asserts the suite goes RED. Once
iters 254–255 gave its detecting arms a clone-set precondition, a fresh checkout **skips** them — and the
battery's failure text changed, mid-repair, from *"the declared-GREEN control went RED"* to:

> `THEATRE: mutant 'anchor-header-lookahead-dropped' left the suite GREEN.`

**That verdict is correct.** On that tree the mutant genuinely is undetected, because its detector is not
running. The battery was not broken by the repair; it *measured* it.

So the rule is not *"the battery needs clones too"*. It is: **a mutation battery is evidence only where
the arms that would detect the mutation actually run**, and every precondition added anywhere in a suite
narrows the trees on which that battery's GREEN means anything. Skipping the battery where its detectors
skip is the only honest option — the alternative is a green mutation proof that proves nothing, which is
`§5` rule 77's lesson arriving from a completely different direction.

The reason string on that skip is the longest in the class, deliberately: it is the one a future reader
is most likely to delete as redundant.

## `D-M257x-255-2` — `PR-3` refuted upward: the class reached ZERO in one iter

I predicted the residual 12 would not close this iter, on the grounds that 12 heterogeneous members is
more than one iter of honest reading. It closed. Recorded because the milestone is careful to book
optimistic misses and should book pessimistic ones on the same terms — **an estimate that was wrong in the
comfortable direction is still wrong**, and the cause is legible: after iter-254 established the two
predicates and the reading discipline, each remaining member cost a `--tb=short` read and three lines.
**The expensive part was the first two; the marginal cost collapsed.**

## `D-M257x-255-3` — the failure that PROPOSES a destructive repair

`test_no_exemption_outlives_its_site` asserts that no exemption has outlived its site — *"an exemption
for a site that no longer matches is fiction carried as reassurance."* On a fresh checkout the findings
set shrinks, so two live exemptions read as stale and are named:
`CLAUDE.md:226:internal/jobsimulation/runner` and
`frontend_architecture.md:39:NEXT_PUBLIC_BACKEND_API_URL`. **Both are real.**

This is a distinct severity within the class and worth separating from it. Most members merely mislead;
this one **instructs**, and the instruction is to delete two correct exemptions. The same shape appears in
`test_01_no_undeclared_markdown_citation_resolves_nowhere`, whose assertion text ends *"**Report it as a
corpus defect — do not widen a class to absorb it**"* — on an unprovisioned box that is 15 defect reports
against a corpus that is right.

**A fresh-checkout-hostile test is not only a false alarm; some of them are false alarms with a work
order attached.**

## `D-M257x-255-4` — `PR-1` REFUTED: the third precondition never appeared

Two preconditions cover the whole 22-member class: the clone set (19) and `node_modules` (3). I predicted
a third cause or a non-environmental member **would** surface; neither did, across the full population read
one failure at a time — so the sealed prediction is **refuted**, and the milestone gets a cleaner answer
than the one it bet on. (This entry first read *"`PR-1` held"*, which inverted the seal: the sealed CLAIM
was *"no third appears"* and my PREDICTION was that it was false. Corrected before the iter closed —
grading a pre-registration is itself a derived figure, and this milestone's rate on those is about 1 in 3.) That is a **bounded negative** of the kind iter-252 recorded — the class is closed,
not merely unexhausted, because every member was examined rather than sampled.

## `D-M257x-255-5` — the ceiling arrow breached its own ceiling in BOTH iters that wrote one

iter-254 recorded this as the expected shape on one instance. It recurred immediately, so it is now a
property rather than an anecdote: writing the arrow that raises the `COMMENT_LITERAL_CEILING` adds a
comment literal, so convergence always takes two passes. Folded into the arrow text itself so the next
author expects it instead of re-discovering it.
