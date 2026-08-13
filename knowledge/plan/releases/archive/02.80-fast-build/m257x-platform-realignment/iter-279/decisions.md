# iter-279 — decisions

_(populated as the iter runs)_

## `D-M257x-279-1` — fix the SOURCE, and retract the reason iter-278 gave for not doing it

iter-278 repaired `platform-alignment.md` §8's claim that `DEMO_ADVANCE_CLONES=pinned` reads the
**canonical** pin. That sentence was **copied verbatim from `clone_pin_guard.py`'s docstring**, which
still said it. `D-M257x-278-6` declined to fix the source because *"an rext edit advances the clone past
the sha this very iter has just reconciled the corpus to."*

**That reason does not survive the same iter.** iter-278 then took exactly that cost for the
census-denominator fix — commit, tag, push, verify-on-origin, advance the clone, re-point the corpus —
and the loop was minutes of work. **A decision that declines a fix on a cost the same iter proved
affordable is a decision that has been refuted by its own author**, and leaving it standing would tell
the next reader the loop is expensive when it is not.

**Decision: take the fix, and record the retraction rather than quietly reversing it.** Repairing a copy
while leaving the source is how the claim gets copied forward a second time — which is precisely how this
one reached the corpus.

## `D-M257x-279-2` — a duration is a measurement noun, and that is NOT a change to make in passing

The docstring repair tripped `TheNounVocabularyIsMeasuredNotAssumed` twice. The first word,
`numerators`, was classified into `_NOT_NOUNS` on the documented `iters`/`series` precedent: the number
beside it is a DENOMINATOR and the word opens the next clause, so the adjacency is punctuation, not
quantification, and nothing in this repo ever tallies numerators.

The second, `minutes`, is different and was **refused**. A duration genuinely IS a measurement noun, so
the honest home is `_MEASURED_NOUNS` — and putting it there makes **every `N minutes` in the repo a
tracked measurement literal**. That is plausibly the right change and it is emphatically not one to make
inside a documentation iter to unblock a commit. Routed as
`ROUTE-M257x-279-durations-are-unclassified-measurement-nouns`; the incidental figure was dropped from
the comment instead, because it was never load-bearing.

**Note the reach this exposes:** the residual arm scans `stack-core/` only, and `~20 minutes` sits
un-graded in three `dev-stack`/`demo-stack` test files today. The vocabulary's zero is a zero **over
`stack-core`**, which is `§5` rule 60 applied to the instrument that enforces it.

## `D-M257x-279-3` — quoting the defect committed it, for the third consecutive iter

The first repair of the `minutes` finding **re-created it**, by quoting the offending clause verbatim
inside the comment explaining why it was wrong. iter-277 did this with an elided route id (its footnote
re-tripped `route_disposition_guard`); iter-278 did it with a closed route id (`D-M257x-278-7`); this is
the third, and the first where the guard is a *vocabulary* rather than a registry.

**The generalisation, now that there are three:** these fences read PROSE, and prose *about* a defect
contains the defect. **An explanation must describe the offending construct, never reproduce it** — the
duration here is written in words for exactly that reason.

## `D-M257x-279-4` — the commit message was mangled by the shell, and it is recorded rather than force-fixed

The rext commit was written with `git commit -m "…"` in **double quotes containing backticks**, so the
shell ran command substitution on them: two bullet labels — the classified word and the refused one —
were **executed and deleted**, leaving `*  -> _NOT_NOUNS` and `*  -> REFUSED` with their subjects gone.
`command not found` appeared in the output and the commit was already pushed before it was read.

**Decision: do NOT force-push.** The forbidden-ops list is not conditional on the fix being small, and a
mangled message is a cosmetic loss where a rewritten published history is a real one. The full reasoning
is durable in `decisions.md` and `progress.md` here, which is where a reader looks anyway.

**The lesson generalises past git.** This iter's other three findings are all *"the medium altered the
message"* — a fence reading prose about itself, a vocabulary closing on the comment that widened it, an
explanation reproducing the construct it explains. **A commit message is prose passed through a shell,
and backticks are code there.** Use a heredoc or `-F -` for any message containing them.
