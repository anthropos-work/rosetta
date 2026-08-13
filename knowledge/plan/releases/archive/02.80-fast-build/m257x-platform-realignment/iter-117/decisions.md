# iter-117 — decisions

## `D-M257x-117-1` — the class lives in the ANCHOR arm, the one arm no reader exercises

The census enumerated **1,520** intra-corpus citations over 92 documents and found **8** false. All eight
are broken `#fragment` anchors. **Zero** of 1,337 path citations and **zero** of 4 explicit line-pins.

That distribution is the finding, not the count. A reader following a broken fragment still lands on the
right file, sees plausible content, and moves on — so the defect is invisible to exactly the process this
milestone has been running for weeks. Four graded readings of this corpus booked none of these eight.

**The rule:** when choosing what to census, ask which part of a citation a human never checks. That is
where a machine buys something a reading cannot.

## `D-M257x-117-2` — the machine-reachable half of class 1 is largely DISJOINT from the sampled half

iter-116's band #7 measured 10 wrong-construct intra-corpus citations. Those are **construct** defects — a
pin naming lines that hold something else. This census measured that shape directly and it is **not
machine-reachable at scale**: of **387** lines carrying a bare `` `:NN` `` pin, only **4** name exactly one
corpus document and no other source path. The rest resolve to a platform file named earlier in the same
sentence (`app/main.go:524`, then `` `:525` ``) or to a port.

**Recorded before the reading that will grade `TOK-08`, deliberately.** The honest projection is that
closing class 1's mechanical half will **not** move `P` by much. That is early evidence bearing on
`TOK-08`'s own refutation branch, and the iter's `overview.md` pre-registered it as an acceptable outcome
rather than a failure. Discovering it *after* the grading reading would have read as the method failing,
when what is really narrower is the method's SCOPE — the readings sample a class whose mechanical subset
is smaller than the class.

**This does not license re-cutting anything.** Clause 5 is untouched, `TOK-08`'s threshold is untouched,
and the sweep continues to class 2. The projection is on the record so that a `P` that barely moves reads
as *predicted*, not as *explained afterwards*.

## `D-M257x-117-3` — under-flag rather than false-RED, and annotate every exclusion with its cost

Four drafts of the census ran before one line of corpus prose was repaired, and three existed only to kill
a false-positive class:

| draft rule | measured false REDs |
|---|---|
| a backticked basename is a relative path | ~180 |
| a bare `` `:NN` `` resolves against the last-named doc | 256 |
| `knowledge/…` / `.claude/…` are rosetta paths | 5 |
| an anchor is a markdown heading (only) | 22 |

A fence shipping on any one of them turns ~460 correct citations RED, and §8 rule 6 says where that
ends — it gets disabled, and then it defends nothing.

**The decision:** each exclusion is (a) taken deliberately, (b) annotated in the guard's docstring **with
the number it cost**, and (c) pinned by a named regression test. A rule adopted by reasoning is a rule
nobody can audit later; a rule adopted by measurement carries its own evidence. The guard therefore
UNDER-enumerates by design, and says so — the correct direction for a fence, and the opposite of the
reach-metric defect iter-114 fixed, where a denominator flattered itself by shrinking.

## `D-M257x-117-4` — the mutation control caught a silent vacuity bug in the fence it was written for

On its first run, the cross-document anchor mutation did not fire. The cause was not the fixture: `run()`
took an unresolved repo root while `resolve_link()` called `.resolve()` on relative targets, so the two
sides of the anchor lookup disagreed the moment the tree was reached through a symlink — `/var` →
`/private/var` on macOS is enough. `tp not in anchors` then **skipped the C2 cross-document arm entirely**
and the guard printed a clean pass over a check it never ran.

Eight vacuous fences have been caught in this milestone. This is the ninth, it was caught **before the
fence shipped**, and it was caught by the control the protocol makes mandatory rather than by a later
reading. Recorded because it is the strongest available argument for the rule that produced it: **every
fence ships with a mutation control that can actually fire.**
