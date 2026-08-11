# iter-106 — closeout

**Shape:** tik · `iter_shape: fence` · **`TOK-06` step 1** — `FIX-M257x-iter103-drift-fence-gap`, which
iter-103 ranked **above** the read-union repair.

## The one-line answer

**The 61 % inflow now has a watcher, and on its first committed run it fired once — correctly.**
`clone_drift_guard` names `sentinel` as **2 commits past everything the corpus cites**, and those two
commits are a dependency bump that produced **both** of iter-103's booked pin-drift predicates. Zero false
positives; it parsed no prose to find it.

## What the fence asserts, and the assertion it deliberately refused

**D1 — cited-clone advance.** For every cloned repo the corpus cites by sha, at least one cited sha IS the
clone's current HEAD. When none is, the repo has moved past everything the corpus knows about it — reported
with the commit distance and the sites at risk.

**D2 — pin agreement, conservative.** A site naming `` `<repo>/go.mod` `` **and** a `<module> v<semver>`
must agree with that `go.mod`; a site pinned to a ref the clone is not at is **UNMEASURED and named**.

**The refused assertion is the interesting one** (`D-M257x-106-1`). *"A version the corpus states must equal
the clone's"* was designed, measured, and rejected: `shared_libraries.md:85` states `sentinel v1.200.0` and
cites `sentinel/go.mod:9 @ 88bc5592`, and **`git show 88bc5592:go.mod` reads `proto v1.200.0`** — at that ref
the claim was TRUE. §5 rules 41/44 make it ref-scoped, so a fence calling it false asserts something it never
measured. **A fence that cries wolf gets suppressed**, and this milestone has the receipt for what that
costs: a silently-refused perf patch shipped a 76 s members grid for four releases.

So D1 reports the **advance**, which is mechanically true and accuses no sentence: *a whole repo's worth of
change has never been looked at.*

## No baseline file — the baseline is derived from the corpus's own citations

A checked-in `repo → last-reconciled-sha` map is **§2's hand-maintained tuple in a new costume**, and its
first value would have to be *asserted*, since iter-103 proved the corpus is not reconciled to the current
clones.

Instead every backticked sha in `corpus/**` is resolved with `git cat-file -t` against every clone. **A sha
resolves in exactly the repo that contains it**, so attribution is exact with no naming convention and no
list. Measured live: **103 distinct shas · 13 of 14 clones cited · 0 ambiguous** (asserted as a test — an
ambiguous attribution would make every finding meaningless).

## What it caught, verified down to the commit

```
sentinel  is at f2c46190, which the corpus never cites — 2 commit(s) past the nearest of 2 cited shas
          5 citing sites: dependency_map.md:50 · shared_libraries.md:57 · :85 · :213 · seeding-spec.md:592

  88036d7  chore(deps): update dependencies to latest versions
  f2c4619  chore(version): v0.24.2
  go.mod   colony v0.34.3 -> v0.35.2  ·  proto v1.200.0 -> v1.210.0
```

**Both** of iter-103's booked pin-drift predicates are downstream of that single advance —
`clerkenstein.md:275`'s *"sentinel … still on `v0.34.3`"* and `shared_libraries.md:85`'s *"the live skew is
two … `sentinel v1.200.0`"*. The fence found the **cause** of two independently-booked reading findings,
mechanically, and the two reading passes had to find the **effects** separately.

## Reach — stated, and asserted as a test (§5 rule 46)

D1 catches: **a whole repo advanced unreviewed.**
D1 does **not** catch: one site cites HEAD while five others are stale — one fresh citation reconciles the
repo. `app` reads reconciled=YES on that basis while iter-103 booked 6 `ai`-fold anchors inside it.

The OK line says this in its own words, and a test asserts the OK line says it, because otherwise *"no
drift"* reads as *"the corpus is current"* — a different and much larger claim.

## Controls — TOK-06's binding clause, discharged

**20 tests, all green.**

- **Mutation battery, 7 mutants** — the unmutated fence RED on the fixture (the control on the controls) ·
  the `reconciled` discriminator removed would report every cited repo · citing HEAD must CLEAR it, or the
  fence can never be satisfied and gets suppressed · exit 1 on drift, exit 0 under `--report` · **no corpus
  is `CANNOT RUN` (exit 2), never clean** · the sha floor is `{7,40}` and stated · both clone roots are read.
- **Anti-vacuity against the SUBJECT, not the inputs** (§8's iter-94 rule): the LIVE corpus and LIVE clone
  set must yield **≥ 20 distinct shas, ≥ 5 attributed repos, ≥ 5 clones, and 0 ambiguous**. A recognizer
  that quietly stopped matching would fail these, not pass vacuously.
- **Real git repos in every fixture** — the fence resolves shas with `cat-file`, so a mock would prove
  nothing.
- **Family placement asserted:** `clone_drift_guard` is in `guard_family.census()` and reconciles in both
  directions (§8 rule 1's derived registry).

## D2 ships conservative, and says so

It graded **1 pin** and named **7 unmeasured** sites on the live corpus. One graded pin is close to vacuous,
and the response is to **print the denominator** rather than widen the recognizer until it produces numbers
(`D-M257x-106-4`). Two widenings were tried and dropped — a nearest-neighbour association is a coin flip
dressed as a measurement, and grading against any repo on the line is sound but empty.

**The reason it is empty is a corpus-side finding, not a fence-side one:** the corpus writes pins as
`<repo> <version>` with the **module** implied by a table heading, so the module token is not on the line at
all. Recorded as a convention in §8's new fifth-layer section — write `<module> <version>
(`<repo>/go.mod:N`)`, all three on one line. That is the fence-facing half of §5 rule 44: naming the tree
and the ref makes a claim *settleable*; putting them where a checker can see them makes it *checked*.

## The family is committed RED, deliberately

`14 GREEN · 1 RED · 0 could-not-check · 3 not-run` over **18** members. **Read it correctly:** clause 3's
fence (`platform_alignment_guard`) is GREEN, clause 4's schema fence is untouched, and the RED is a **new**
member finding **pre-existing** drift. It is not a regression in either clause.

Committed that way on purpose (`D-M257x-106-3`): clause 4's own wording is *"asserted by a FENCE that is
watched going RED, not by inspection"*, and TOK-06 puts repair at step 3 precisely so the repair has
something watching it. Its subject is **5 sites in 2 predicates**, both already inside
`FIX-M257x-iter103-read-union` — so step 3 clears it, and **if step 3 does not, that is itself the finding.**

## Tests

**`stack-core`: 957 passed · 1 failed.** The +20 are this iter's. The single failure is the same
`test_claim_twin_guard_iter48_answer_key.py::test_02` proven pre-existing at iter-105 — reproduced three ways
against the run-open tree — and routed as `FIX-M257x-iter105-claimtwin-green-twin-refire`.

## Gate

**Unchanged at 4 of 5.** No `N` movement claimed: this is clause 3's instrument, never clause 5's. Clause 5
was not re-cut, narrowed, reinterpreted or argued.

## Housekeeping

Zero platform-repo edits. `stack-demo/**` untouched — the clones were **read** (`rev-parse`, `cat-file`,
`rev-list`, `show`), never fetched or modified. rext stays on `main`; **no tag cut** — nothing here must be
consumed by a stack.

## Close — 2026-08-06

**Outcome:** the drift fence lands and fires once, correctly — `sentinel` 2 commits past every cited sha,
the cause of both booked pin-drift predicates, found without parsing a sentence. 20 new tests, 7-mutant
battery, anti-vacuity against the live tree.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: **n — the
family's 1 RED is this iter's own new fence catching pre-existing drift, which is the deliverable
(`D-M257x-106-3`), not a test-gate failure; the suite's 1 failure is the iter-105 pre-existing one** —
(5) cap-reached: n (2 tiks) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-106-1` (refuse to adjudicate truth-at-a-ref — and why that is the design) · `-2`
(no baseline file; `cat-file` attribution is exact) · `-3` (committed RED, deliberately, with the
suppression risk named) · `-4` (D2 conservative; its low yield reported, two widenings tried and dropped) ·
`-5` (the finding verified downstream to the commit, not just reported)
**Side-deliverables:** none.
**Routes carried forward:**
- **TOK-06 step 2 — the induction checks**, now carrying two items: the two shapes iter-103 measured
  (centralised-wording control + post-repair line-offset check) **and**
  `FIX-M257x-iter105-claimtwin-green-twin-refire`, which is the same class one layer down.
- **TOK-06 step 3 — repair the 33**, which must clear `clone_drift_guard`'s 5 sentinel sites. **The fence is
  now step 3's answer key**; a repair that leaves it RED has not finished.
- **`DOC-M257x-iter106-pin-claim-convention`** *(net-new, landed as a convention, not yet applied)* — pin
  claims should be written `<module> <version> (`<repo>/go.mod:N`)`. Applying it corpus-wide is a step-3
  rider, not a separate item.
- Unmoved: `FIX-M257x-iter56-assignment-flake`, `FIX-M257x-iter103-assignment-context-bleed`,
  `DEF-M257x-iter103-aws-bind-provenance`, `DEF-M257x-iter101-briefing-rext-tree`, `RF-2/3/7–14`, the five
  pass-22 items.
**Lessons:** **the strongest fence in this iter is the one that refuses to say the interesting thing.** D1
never calls a sentence false — it says a repo moved — and that is exactly why it has zero false positives
and why it found the *cause* of two findings two full reading passes could only find the *effects* of.
Generalisable: **when a claim is ref-scoped and your instrument is not, assert the movement, not the
falsehood.** The reading can adjudicate truth; a fence's comparative advantage is noticing that the ground
shifted.
