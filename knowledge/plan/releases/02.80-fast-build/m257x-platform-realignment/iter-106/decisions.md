# iter-106 — decisions

## `D-M257x-106-1` — the fence does NOT adjudicate truth-at-a-ref, and that is the design rather than a limitation

The obvious drift fence is *"a version the corpus states must equal the clone's."* It was designed, measured
against the live tree, and **rejected**.

`shared_libraries.md:85` states `sentinel v1.200.0` and cites `sentinel/go.mod:9 @ 88bc5592`. **At that ref
the claim was true** — verified: `git show 88bc5592:go.mod` reads `proto v1.200.0`. §5 rules 41 and 44 make
it a ref-scoped claim, and a fence calling it *false* would be asserting something it never measured.

That is not a pedantic point. **A fence that cries wolf gets suppressed, and a suppressed fence is worse than
no fence** — this milestone has the receipt: a silently-refused perf patch shipped a 76 s members grid for
four releases (`demopatch-spec.md`).

So the assertion is the one that *is* mechanically true and *is* the inflow:

> **This repo has moved past every commit the corpus knows about, by N commits, and here are the citing sites.**

It accuses no sentence. It says a whole repo's worth of change has never been looked at.

## `D-M257x-106-2` — there is no baseline file, because the baseline is DERIVED from the corpus's own citations

The natural implementation is a checked-in `repo → last-reconciled-sha` map. It is also **§2's
hand-maintained tuple in a new costume**: it drifts, it gets re-accepted absent-mindedly, and its *first*
value would have to be asserted rather than measured — there is no honest sha to seed it with, because
iter-103 proved the corpus is **not** reconciled to the current clones.

Instead: every backticked sha-shaped token in `corpus/**` is resolved with `git cat-file -t` against every
clone. **A sha resolves in exactly the repo that contains that object**, so attribution is exact and needs
no naming convention, no per-doc rule, and nothing to keep in step. Measured on the live tree: **103 distinct
shas, 0 ambiguous** (asserted as a test, not assumed — an ambiguous attribution would make every finding
meaningless).

The corpus's own citations are the baseline, and they cannot go stale relative to the corpus because they
are read out of it on every run.

## `D-M257x-106-3` — the fence is committed RED, deliberately

The family now reads **`14 GREEN · 1 RED`**, and the RED is this new fence. Three reasons it is committed
that way rather than repaired first:

1. **Clause 4's own wording** is *"asserted by a FENCE that is watched going RED, not by inspection."* A
   fence never seen RED on real drift in a committed state has not been watched.
2. **TOK-06's sequence puts repair at step 3**, two iters out, precisely so the repair has something
   watching it. Repairing now would leave step 3 with no answer key.
3. iter-83's lesson: **a guard that was not run reads exactly like a guard that passed.** The inverse also
   holds — a guard that has only ever been green on a tree with known drift in it is a guard nobody has
   evidence about.

**Read the family verdict correctly:** clause 3's fence (`platform_alignment_guard`) is GREEN and clause 4's
schema fence is untouched. The RED is a *new* member finding *pre-existing* drift. It is not a regression in
either gate clause.

**The risk, named:** a permanently-RED member invites suppression. The mitigation is that its subject is
**5 sites in 2 predicates**, both already booked in `FIX-M257x-iter103-read-union`, so step 3 clears it — and
if step 3 does not, that is itself the finding.

## `D-M257x-106-4` — D2 ships conservative, and its low yield is REPORTED rather than dressed up

The pin rule fires only where the site names `` `<repo>/go.mod` `` — the file that literally holds the
answer — and only when exactly one such repo is named on the line. On the live corpus that graded **1 pin**
and named **7 unmeasured** sites.

One graded pin is close to vacuous, and the honest response is to print the denominator (it does) rather
than widen the recognizer until it produces numbers. Two widenings were tried and dropped:

| widening | why dropped |
|---|---|
| associate a `<module> v<semver>` with the nearest `<repo>/go.mod` on a multi-repo line | it guesses. `shared_libraries.md:85` puts four repo/pin pairs and two platform refs on one line; a nearest-neighbour rule is a coin flip dressed as a measurement |
| grade a pin against *any* repo named on the line | sound but empty — the corpus writes `<repo> <version>` with the **module** implied by a table heading, so the module token is not on the line at all |

**The second one is the real finding, and it is corpus-side, not fence-side:** the corpus's pin-claim
*writing convention* is unreadable to any line-scoped checker. Recorded in §8's new fifth-layer section as a
convention — write `<module> <version> (`<repo>/go.mod:N`)`, all three on one line — which is the
fence-facing half of §5 rule 44.

## `D-M257x-106-5` — the fence's first finding was verified downstream, not just reported

A guard that names a repo is easy to believe and easy to be wrong about. The finding was checked to the
commit:

```
sentinel  88bc5592..f2c46190  =  2 commits
  88036d7  chore(deps): update dependencies to latest versions
  f2c4619  chore(version): v0.24.2
  go.mod:  colony v0.34.3 -> v0.35.2 ;  proto v1.200.0 -> v1.210.0
```

**Both** of iter-103's booked pin-drift predicates — `clerkenstein.md:275`'s *"sentinel … still on
`v0.34.3`"* and `shared_libraries.md:85`'s *"the live skew is two … `sentinel v1.200.0`"* — are downstream of
that single advance. The fence found the cause of two independently-booked findings **without parsing one
sentence**, which is the strongest available evidence that D1 is watching the right thing.
