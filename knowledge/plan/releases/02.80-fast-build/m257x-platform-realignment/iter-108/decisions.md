# iter-108 — decisions

## `D-M257x-108-1` — the 3-no-prog tok-trigger does not fire on UNMEASURED, and this is now protocol

**The streak is literally present**: iters 105, 106 and 107 are three consecutive tiks (iter-104's tok
opened the window) and none moved `N`.

**It does not fire.** Rule 2 defines no-progress as *"the metric did not move in any of those 3 tiks (zero
or net-negative delta)"* — and **a delta requires two measurements**. Those three iters took **none**;
`TOK-06` sequences the read at **step 4, last**, and each of the three states in its own close that no `N`
movement is claimed. **The metric is UNMEASURED, not unmoved**, so the trigger's precondition is
unestablished.

This is §8's *grade the cannot-tell* (iter-91; re-applied at iter-107 when `anchor_offset_guard` refused to
assert a class it could not decide) turned on the skill's own trigger. **Grading "not measured" as "did not
move" asserts something nobody measured.**

**The substantive check agrees with the formal one.** A triggered tok exists to revise a **stalled**
strategy; `TOK-06` is **3 of 5 steps executed, on its declared schedule**, with both metric-moving steps
still ahead. Revising it now would revise it *before its own evidence exists* — step 4's reading **is** the
evidence a revision would need. Firing here would also mean **no declared multi-step strategy with more
than three non-metric steps could ever be executed**, because the tok would terminate the call mid-sequence
every time.

**Codified in `corpus/ops/platform-alignment.md` § 9** so the next agent inherits a rule instead of
re-deriving a judgement call (the skill's protocol-evolution guideline: a lesson that generalises beyond
the iter goes into the protocol doc, in the same commit).

## `D-M257x-108-2` — the `:8081` multiplier is REMOVED, not corrected

iter-102 published *"one occurrence anywhere in the clone set"* to **5 anchors**. The literal has **six**,
five of them inside the sentence's own 13-repo/44-`.tf` denominator — **self-refuting at its own stated
scope**.

**A corrected canonical sentence is still a canonical sentence.** Rewriting the count at five sites would
leave the exact mechanism that turned one authoring error into a five-anchor defect, and the next author
would re-multiply it.

So the fix is **structural**: the count is derived **once**, in `backend.md`, and `cms.md` /
`jobsimulation.md` **point at it** instead of restating it. **A pointer cannot carry a false cardinality to
five places.**

Re-derived before any wording was written (the brief's rider — *a centralised wording is an instrument and
needs a control*): **6 occurrences** (`app` 1, `rosetta-extensions` 5), **0 in any `.tf`**, **44 tracked
`.tf` files across 13 repos**.

## `D-M257x-108-3` — a bulk route is graded one item at a time; historical anchors are never re-pointed

`FIX-M257x-iter107-unbooked-rot` routed **5** rotted citations. Graded individually:

| citation | verdict |
|---|---|
| `service_taxonomy.md` → `platform-migration-status.md:89` | **resolves correctly today** — untouched |
| `services/README.md` → `platform-migration-status.md:101` | **resolves correctly today** — untouched |
| `shared_libraries.md` → `storage.md:115` | **rotted** → `:129` |
| `platform-alignment.md` → `backend.md:241` | **rotted** → `:302` (third rot of this one anchor) |
| `platform-alignment.md` → `hiring.md:93` | **HISTORICAL — must not move** |

**Bulk-bumping all five would have broken three working citations and falsified one historical record.**

The last row is the general finding: *"iter-39 had two auditors find `hiring.md:93` independently"* is a
record of **what iter-39 found**, not a live citation. Re-pointing it would make the sentence false about
the past. **A guard cannot distinguish a live citation from a record of where something once was** — the
same undecidability `D-M257x-107-2` hit. The anchor is left, with its historical status now explicit.

**Corollary, applied to §8's own post-mortems:** an unpinned `file:line` in a retrospective is
indistinguishable from a live claim — `repair_postcondition` and `anchor_offset_guard` each refused a
commit over exactly that this iter. Retrospective line numbers now carry the ref they were true at, which
is `TOK-04`'s rule applied to the document that teaches it.

## `D-M257x-108-4` — reach is reported TWICE: raw, and over the upheld union

`repair_reach_guard` grades **all 48 booked blocks** in the ledger, including findings the adjudicators
**rejected**. Reporting only its headline understates the repair; reporting only the filtered number hides
the instrument's actual output.

- **Raw reach: 46/47 = 97.9 %.**
- **Over the repair's actual input (the 22 upheld predicates): 46/46 = 100 %.**

The single unreached booking is `shared_libraries.md:128` / `r25-G B3`, **REJECTED by adjudicator 4, class
`wrong-tree`** — it graded `app/internal/ai/` at `ad9f3c49` (app's post-fold diverged fork) rather than the
`ai` module at the `v1.40.2` the section's own pin row names, readable in the same clone at `1e457fa70`,
where all three booked claims hold verbatim.

**This reproduces iter-102's residue result exactly**: the apparent miss is a claim that came out **true**.

## `D-M257x-108-5` — the drift-fence limitation stays OPEN; no exclusion list

`FIX-M257x-iter107-drift-fence-satisfiable-by-prose` asks whether a **derived** discriminator can replace
the known limitation that *writing about the drift satisfies the drift fence*.

**No derived discriminator was found, and none is being faked.** D1 asserts *"every cited clone's HEAD is a
commit the corpus cites"*; a doc that merely **mentions** a sha satisfies it. Separating *"this sha dates a
claim"* from *"this sha is being discussed"* is a question about **what the author meant**, and **intent is
not in the repository** — the identical wall `D-M257x-107-2` hit. Every candidate rule (require an `@`,
require an adjacent `file:line`, require a table cell) is a **shape allow-list in a derivation's clothes**,
i.e. §2's hand-maintained tuple, rejected twice in this milestone.

**Left open, pinned by its known-limitation test.** Noting that this iter's green is **earned** rather than
satisfied-by-prose: the drift was actually repaired and the 2 gradeable pins **match**.

## Side-deliverable — `rext 680e852`: a bare rev is not a range

Not planned scope; recorded separately so it does not blur the close status.

`anchor_offset_guard` accepted a **bare rev** and graded the wrong comparison. `git diff <sha>` is *<sha>
versus the working tree*, not versus its parent — so on `cd16967`, **this guard's own pinned answer key**:

```
cd16967             60 changed files,  0 graded of 33 seen,   0 findings, exit 0   OK
cd16967^..cd16967   53 changed files, 17 graded of 33 seen,  10 findings, exit 1   RED
```

**All 18 existing tests passed throughout**, because every one of them passes the explicit `^..` form. The
defect lived exactly in the shorter invocation an operator reaches for first — and it defeats the module's
own anti-vacuity rule (§5 rule 8), whose existing guard only catches an **empty** range: this run is not
empty, it is full of the **wrong** comparison, so nothing could report the difference.

Fixed by normalizing to `<rev>^..<rev>` — the semantics `repair_reach_guard` already had; **the asymmetry
between two fences graded by the same operator was itself the hazard.** +3 tests (21 total), including an
anti-vacuity control that fires: with `normalize_range` reverted in-memory the bare form returns **0
findings / 0 graded**, so the new tests fail.

**Found by Phase 0d — pre-flighting the fence that was about to grade this iter.** Without it, this iter's
own `anchor_offset_guard` verdict would have been a false green.
