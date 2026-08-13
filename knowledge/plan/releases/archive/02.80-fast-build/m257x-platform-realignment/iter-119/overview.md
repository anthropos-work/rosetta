---
iter: 119
milestone: M257x
iteration_type: tik
iter_shape: reading
status: closed
opened: 2026-08-07
---

# iter-119 — the grading reading: `TOK-08`'s own arithmetic, against the number

**Active strategy reference:** [`TOK-08`: census the mechanical classes; stop sampling
them](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07) — milestone-root
`decisions.md`, **supplied by the USER on 2026-08-07**, not authored by an agent, and **binding**.

`TOK-08` put the reading *after* the sweep and pre-registered its own falsification as arithmetic before a
single sweep line was written. **The sweep is complete. This iter is the grading.**

## Step 0 — re-survey before targeting (mandatory)

| check | result |
|---|---|
| Is the full mechanical sweep actually complete? | **Yes.** Class list fixed in iter-117's `overview.md`, may only grow, **did not grow**. Class 1 (intra-corpus citation) 1,520 enumerated / 0 findings / **100 %** reach; class 2 (platform-source citation) 1,070 candidates / 0 findings / **63.1 %** reach, denominator named. Both ship a mutation control and an anti-vacuity control that can fire |
| Is `TOK-08`'s trigger therefore armed? | **Yes** — iter-118's close states it in those words |
| Is the baseline still the right one? | **Yes — `P = 37` / `N = 41` at corpus `f581de09`**, sealed at `85f6f1c` before iter-116's first seat |
| Has the subject moved under us? | **No.** 14 platform clones at the same shas for a **fourth** consecutive reading; the in-scope corpus is **5 in-place lines** and **zero net lines** from `f581de09` |
| Is a successor strategy in scope? | **NO.** `TOK-08` forbids one on the refutation branch, in the user's own words. **There is no `TOK-09`** |

**Type selection: tik.** The 3-no-prog tok-trigger cannot fire — of the last three tiks, **117 and 118
took no reading**, and §9's refinement reads UNMEASURED as UNMEASURED, not as unmoved. Both said so in
their own closes, in those words, and `TOK-08` declared the read-last sequence in advance, which is that
rule's second guard-rail. **Separately and independently: a triggered tok is forbidden here by the user's
binding instruction**, so even were the trigger armed it could not be honoured.

## This iter: the reading

**Hypothesis — and it is `TOK-08`'s, not a fresh one.** If the mechanically-decidable classes were the
bulk of the standing pool, censusing them exhaustively drains it and `P` falls by half or more. If they
were not, `P` holds and enumeration-first is refuted alongside repair-and-read.

**The instrument is copied verbatim, not redesigned.** `instrument/briefing-iter76-AS-RUN.md`, sha
`3858ec53…`, re-checked after copying and after the addendum was appended, `diff` empty both times.
14 seats (7 × 2 blind readings of an identical partition) + 4 adjudicators — the iter-95/97/99/101/103/109/116
shape.

**Expected lift: none is claimed, and that is pre-registered.** `D-M257x-117-2` recorded *before* the
sweep's second class that little `P` movement should be expected, because the machine-reachable half of
class 1 is largely disjoint from the half the readings book. **A flat `P` is a PREDICTED outcome here, and
`TOK-08`'s branch fires on it anyway.** Both are true at once and neither softens the other.

**The structural gift of this reading, disclosed before the number:** the partition is **bit-identical to
iter-116's** (proven by re-running the partitioner at both refs) and the clones are frozen, so this is the
milestone's **first true seat-level replicate** — and **test-retest reliability becomes measurable for the
first time** (band #3). Whatever the primary returns, the sheet also returns the first measurement of how
much a single reading of this corpus can be trusted.

**Phase plan.** (0) re-derive ground truth + reproduce the partition at both refs + run the guard family
with its fence tree printed. (1) **seal the pre-registration in this iter's FIRST commit**, before a seat
is dealt. (2) deal 7 seats, commit each verbatim; deal 7 more, commit each verbatim. (3) adjudicate in 4
groups, commit the verdicts unedited. (4) compute `P` and `N`, grade all 14 bands and `TOK-08`'s branch.
(5) close; state the invocation with every count.

**Escalation conditions.** `P ≥ 19` → **`TOK-08` REFUTED**: report it first and loudest, author **no**
successor strategy, exit `re-scope-trigger` for a user scope decision. A finding that inverts a shipped
security property (the iter-115 `bash -c` class) → grade by consequence and report it plainly regardless
of class. Any repair temptation inside the pass → refuse; route it.

**Acceptable close-no-lift outcomes.** **A refuted `TOK-08` is a first-class, pre-authorized outcome, not
a failure** — the user stated in advance they would carry it. A reading that returns a flat `P` with a
high band-#3 overlap is a *successful measurement of a standing pool*; a reading that returns a flat `P`
with a LOW band-#3 overlap is a more important finding still, about the instrument rather than the corpus.
