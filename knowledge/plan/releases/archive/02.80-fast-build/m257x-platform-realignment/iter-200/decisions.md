# iter-200 — decisions

## `D-M257x-200-1` — the memo population is ENUMERATED, never asserted to zero

The reflex, given the milestone's habits, is a fence reading `assertEqual([], memoised_disk_readers(REXT))`.
It would be wrong.

These caches are deliberate and load-bearing. `anchor_construct_guard._git_out`'s own docstring records
the measurement: *"Uncached, reading at a ref took `anchor_construct_guard` from ~1 s to 10.9 s on the
live [tree]"*. A fence demanding their removal would trade a tenfold slowdown for a property no guard
currently violates.

What the enumeration buys is different and cheaper: **a future mutation control written against one of
these guards meets a list instead of a surprise.** The arms assert the census is non-empty and that both
shapes (`lru`, `module-dict`) are represented — i.e. that the instrument looked — and stop there.

A no-fire control is included for the same reason: a module-level dict that is only ever **read** is a
constant table, not a memo. Without that arm the census would call every lookup table a cache and its
count would carry no information.

## `D-M257x-200-2` — the repeat predicate's blind spot is measured, not disclaimed

`mutation_rewrite_sites` finds a repeated write to **one target expression inside one function**. That is
the unambiguous mutation shape, and it misses the split one — a helper stages, the test mutates; a
battery calls its stager once per case.

A prose disclaimer would have been the easy move, and `§5` has a name for what that is worth (*a
NOT-REACHED clause is a measurement or it is a mood*). So the blind spot is bounded from above by a
second census: **35 test functions write a `.py` path at all**, against **0** repeated-write `.py` sites.
The gap between those two numbers *is* the blind spot, in the open, with an arm asserting the upper bound
stays materially larger than the lower.

Reading the 35 by hand: four are the mutation batteries' `_stage` helpers, which are safe because each
mutation gets a **fresh tmp directory** — a cache entry from a prior case cannot be hit because the
absolute path differs. **Safe by isolation, not by discipline**, and nothing in those modules says so.
Routed as `SURVEY-M257x-iter200-battery-stagers-are-safe-by-isolation-not-by-discipline` rather than
"fixed", because there is nothing broken to fix — only an unstated invariant, and stating it is a
different iter's work than answering h42.

## `D-M257x-200-3` — h42 is closed with a STRUCTURAL answer, and that is a stronger close than a sample

The route could have been closed by re-running a handful of mutation controls with an mtime bump and
observing no change in verdict. That would be a sample, and this milestone has spent nine iters learning
what a sample is worth against a standing pool.

The census answers it structurally instead: **rule 77's hazard requires re-imported Python source, and no
repeated in-place rewrite in this repo targets Python source.** That is a property of the population, not
an observation about some of it, and it stays true under re-running because the fence re-derives it.

The second half is why the close is not a formality. **Rule 77 is correct and complete about bytecode,
and bytecode is the smaller of the two ways a mutation control here can fail to re-read.** Auditing
compliance with a rule is not the same as auditing the thing the rule exists to prevent — and the audit
this route asked for surfaced the gap only because it went looking for the mechanism rather than for the
rule's wording.
