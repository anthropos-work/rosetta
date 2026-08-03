---
iter: 50
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-03
---

# iter-50 — the variance experiment: read the SAME tree twice, repair nothing

**Active strategy:** [`TOK-02: fence the prose the way the anchors are fenced`](../decisions.md#tok-02-fence-the-prose-the-way-the-anchors-are-fenced--2026-08-02)
— this iteration runs **step 5 only**, twice-over. It builds no fence and repairs nothing. It executes
the one experiment [`platform-alignment.md` §5 rule 22](../../../../../corpus/ops/platform-alignment.md)
explicitly prescribes and this milestone has never run:

> **Measure the variance FIRST, by reading the same tree twice with no repair between; that is a cheap
> experiment and this milestone paid eight passes to learn it.**

## Step 0 — re-survey before targeting (mandatory, done before this plan settled)

- **Platform origin re-fetched at open: `2adcf714`, unchanged.** Re-scope trigger stays at **occurrence 1
  of 2**.
- **The 14 blockers of `FIX-M257x-iter49-blocker-set` are unrepaired**, at the anchors iter-49 recorded.
  Target is current.
- **The corpus under audit has not changed since reading #9 was taken.**
  `git diff --stat 47c9b7d..HEAD -- corpus/ CLAUDE.md .claude/` is **empty**; the two commits since
  (`0f9e14f`, `57dfbfd`) touch `knowledge/plan/**` only. The 40 files are byte-identical to the tree
  seats A–F read.
- **The ground-truth clones are byte-identical to iter-49's**, all twelve re-read at open:
  `app 5ba17044` · `app/studio aeec036a` · `platform 2adcf714` · `next-web-app bb3313bc` ·
  `sentinel 88bc5592` · `storage 4ce8ece5` · `messenger fa47850d` · `cms ca50c817` ·
  `graphql-wundergraph 60c229f3` · `roadrunner 87d8d443` · `jobsimulation 462343b0` ·
  `studio-desk 14a5442a` · `ant-academy 9c3843cd`.

**So the experiment is available at the price of ONE reading, and only right now.** Reading #9 is
already the first half. The moment the 14 are repaired the tree moves and the paired same-tree
measurement is gone — the perishability of §5 rule 21, applied to the instrument instead of the corpus.

## Why this, and not the repair-then-read cycle

The brief for this run named the honest position: the induced class has moved out of mechanical reach
(iter-49: paraphrase leak 3 · overshoot-in-new-text 3 · wrong-mechanism 1), and the series
`25 → 13 → 11 → 17 → 37 → 18 → 7 → 12 → 14` sits at **±5**. Another turn of repair-then-read produces a
tenth number inside that band and settles nothing.

**The quantity that has never been measured is the one the whole gate rests on.** Clause 5 is met by *a
reading that returns zero*. Whether that is even a statement about the corpus depends entirely on the
per-finding detection probability of a single 7-seat pass — and the only evidence about it today is
circumstantial: iter-47 read zero pre-existing where iter-48 then booked ten, seven of which predated the
milestone and **sat inside seats' own assigned file sets during the pass that reported zero**. That is
suggestive. It is not a measurement, because the tree changed between those two readings.

This iteration removes that confound.

## Instrument — HELD FIXED at iter-41's, on every knob, identically to reading #9

- **Seven auditors** — six full-read partitions (A–F) + one adversarial diff seat (G).
- **Same briefing** — every §5 rule, each file's repair history, the same blocker/minor grading rule.
- **Same partition METHOD** — files sorted by line count descending, snake-dealt A→F then F→A.
- **All 40 files read in full, top-to-bottom**, under a per-file `wc -l` positive control.
- **Seat G reads the identical diff** — `2fc633a..47c9b7d -- corpus/`, which is exactly the working-tree
  diff seat G read at #9, now committed. Same 11 files, same 19 hunks.
- **Fresh seats, blind.** No seat is told what reading #9 found, or that a prior reading exists. A seat
  that could see the answer key would measure agreement, not detection.

### Partition (computed, 40 files / 9,326 lines) — and it is IDENTICAL to #9's by construction

The corpus has not changed, so the size-sort deals the same hand. Verified against iter-47's published
table: all six sets match name-for-name.

| auditor | lines | files |
|---|---|---|
| **A** | 1742 | external_services · backend · graphql-wundergraph · academy-backend · coursebuilder · skiller · TEMPLATE |
| **B** | 1598 | ai-readiness · ai_architecture · security_compliance · ai-labs · clerk-integration · customerio-sync · architecture/README |
| **C** | 1528 | alignment_testing · architecture_overview · cms · sentinel · messenger · services/README · db-backup |
| **D** | 1481 | studio-room · clerkenstein · chronos · roadrunner · next-web-app · gotenberg · intelligence |
| **E** | 1479 | service_taxonomy · hiring · shared_libraries · storage · askengine · dependency_map |
| **F** | 1498 | ant-academy · studio-desk · jobsimulation · platform-migration-status · skillpath · frontend_architecture |
| **G** | (diff) | adversarial diff-read of `2fc633a..47c9b7d -- corpus/` — iter-49's own repair |

**Identical partition is the point.** Holding the *method* fixed across a changing corpus deals a
different hand (iter-47 measured 20 of 40 files changing seats). Here the corpus is fixed, so the hand is
fixed too, and any disagreement between the two readings is **seat variance and nothing else** — no
partition confound, no ground-truth confound, no corpus confound.

## Hypothesis

A single 7-seat pass has a per-finding detection probability well below 1, and the two readings will
**disagree far more than they agree**. If so, the count is not the corpus's residual — it is a sample from
it — and "a reading returns zero" cannot certify the corpus at a residual anywhere near the instrument's
own variance.

## PRE-REGISTERED PREDICTIONS — written before any seat is launched, before any report is read

The primary output of this iteration is not a count. It is the **overlap**.

1. **Count.** Reading #10 returns `N₁₀` in **[9, 19]** — 14 ± 5, the band the eight-reading series
   supports. A value outside it refutes the ±5 characterization itself.
2. **Overlap — the decisive one.** **Fewer than 7 of reading #9's 14 blockers are re-found by reading
   #10** (i.e. recall < 50%). Derivation: if per-finding recall were high, iter-48 could not have booked
   seven months-old defects sitting in the file sets of a pass that had just reported zero pre-existing.
3. **Union.** **|#9 ∪ #10| > 14** — reading #10 books at least one blocker reading #9 did not, in the
   unchanged tree #9 had just read in full.
4. **Directional.** The disagreement is **roughly symmetric** — neither reading is simply "better". A
   large asymmetry would mean the seats differ in quality rather than the instrument being noisy, which
   is a different finding and a different remedy.

**Consequence, registered in advance so it cannot be re-framed afterwards.** If predictions 2 and 3 hold,
then a single pass returning zero is consistent with a corpus carrying a residual of roughly `N/recall`,
and **clause 5's instrument is demonstrably incapable of certifying zero** — not because the corpus is
dirty, but because the reading is a sample. That is a statement about the gate's *instrument*, which the
user's ruling did not fix. The ruling fixed the **clause**: met only by a reading that returns zero
blockers. This iteration does not re-open the clause, does not propose re-cutting it, and does not
propose closing at 4 of 5.

**If predictions 2 and 3 are REFUTED** — high overlap, small union growth — that is the better outcome and
the more valuable one: it says the instrument is precise, that 14 is a real count of a real residual, and
that the repair-then-read cycle should simply continue until it reaches zero. The experiment is designed
so that either answer changes what the milestone does next.

## Phase plan

| step | work | done when |
|---|---|---|
| 1 | Verify the tree and clones are unchanged since reading #9; compute and cross-check the partition | done at Step 0 above — corpus diff empty, 12 clone shas matched, partition matches iter-47's published table name-for-name |
| 2 | Capture the **fixture** for the 14 before anything can perish (read-only) | `iter-50/fixture-14.md` records each of the 14 with its anchor and adjudication, taken from iter-49's ledger |
| 3 | Launch 7 blind seats at the frozen instrument, identical partition | all seven raw reports in `iter-50/raw/` |
| 4 | Adjudicate #10 to a blocker ledger, then compute the **paired overlap** against #9 | `blocker-ledger.md` + `variance.md` with recall, union, and the per-seat breakdown |

## Escalation conditions

- A seat that cannot read its files, or whose positive control fails → re-run that seat; a partial pass is
  not a reading (§5 rule 8 — a check that SKIPS reads exactly like a check that PASSES).
- A platform commit landing mid-iteration → `EXIT_REASON: re-scope-trigger` (occurrence 2 of 2).
- **No repair happens in this iteration under any circumstance.** A repair destroys the measurement.

## Acceptable close-no-lift outcomes

**This iteration cannot move the primary metric and is not trying to.** It repairs nothing, so the blocker
count is expected to stay at or near 14. Pre-declared here so the close cannot be graded as a shortfall:
the deliverable is the variance measurement, and the honest close is `closed-fixed` if the paired reading
is taken and adjudicated, `closed-no-lift` only if the reading cannot be completed.

**It is a no-progress tik on the primary metric by design**, which puts the no-prog streak at 3 and makes
the next iter a **triggered tok**. That is the intended sequencing: the tok that follows will be authored
on a measured detection probability instead of on speculation about one.
