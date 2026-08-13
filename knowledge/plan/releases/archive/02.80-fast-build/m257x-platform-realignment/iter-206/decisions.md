# iter-206 — decisions

## `D-M257x-206-1` — four of six were stale, and this run made two of them stale itself

Six standing comment figures in `claim_census_guard.py` describe populations the module computes.
Derived:

| comment | said | derives | verdict |
|---|---|---|---|
| `:230` the ABBREV population *"in these 40 files"* | 40 | **41** | stale |
| `:976` *"unresolved for a class of 2"* | 2 | **1** | stale — **iter-202's own** |
| `:981` *"Measured over the corpus: 2 pairs"* | 2 | **1** | stale — **iter-202's own** |
| `:1275` *"a wholesale warning over 949 pairs"* | 949 | **1,015** | stale since iter-202 |
| `:548` *"21,610 files indexed, 50,357 pruned"* | — | matches | current |
| `:551` *"12 names"* (the prune list) | — | matches | current |

**The two marked *iter-202's own* are the finding.** iter-202 measured the wrong-repo class at **2**,
wrote it into two comments, and then — four paragraphs further down the same iter — fixed the extension
truncation that made the second member exist at all. `ant-academy.md`'s `code/public/catalog.js` was a
**parser artifact**: the file is `catalog.json`. The class was **1** before that iter closed, and its own
comments still said 2 four iters later.

A number can go stale **inside the iter that wrote it**, and the author is the least likely person to
re-read the sentence they just wrote.

## `D-M257x-206-2` — `949` is not re-pinned; the sentence stops carrying a number

The obvious repair is `949` → `1,015`. That buys one iter of correctness and re-arms the same trap: the
value moves whenever the clone set or the parser moves, and nothing reads a comment.

Instead the sentence now points at **`_exp["under_clones"]`**, which the code prints two lines below, and
says out loud that it read `949` from iter-198 until iter-202 moved it and **nothing noticed for four
iters**. *A number a comment does not carry cannot go stale* — the same reasoning iter-199 applied to
printed totals, which is where this whole class began.

The other three are genuine constants of the current tree, so they are corrected in place **and fenced**
rather than removed.

## `D-M257x-206-3` — the fence recomputes, and it checks its own numbers are distinguishable

`TheModulesOwnCommentFiguresAreDERIVED` derives each figure from the live census and asserts the comment
text carries it; the `949` arm asserts the **absence** of the literal and the presence of the pointer.

Plus an anti-vacuity arm nothing else would have given: the three derived values (**41**, **1**,
**1,015**) must be **pairwise distinct**. If two ever coincided, a comment carrying the wrong one of them
would satisfy the wrong arm and the fence would pass on a defect — the *agreeing reconstruction*
failure, one layer down.

## `D-M257x-206-4` — the standing class now has a verified subset, and it is 6 of 168

Before this iter, **one** of the 168 standing figures had been derived (`basename_index`, iter-203).
Now **seven** — six here plus that one. The route's wording changes from *sized but unverified* to
*sized, with a named verified subset of 7 and 161 unverified*, which is the first time the class has had
a numerator at all.

**And the hit rate on the first verified subset was 4 in 6.** That is not an estimate of the remaining
161 — this subset was selected for being *derivable*, and derivable figures sit in the modules under
active edit, which is exactly where staleness concentrates. Stating the selection bias is the point:
a 67 % hit rate on a suspicion-selected subset is a floor on nothing and a rate for nothing, the same
discipline `P` has carried since iter-116.
