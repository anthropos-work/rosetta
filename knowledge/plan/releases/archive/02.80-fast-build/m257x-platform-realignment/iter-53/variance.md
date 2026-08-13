# iter-53 — the paired reading #11 + #12: the arithmetic, and the confound that dominates it

## Headline

| | reading #11 | reading #12 |
|---|---|---|
| blockers (as run) | **32** | **26** |
| per seat | A 2 · B 7 · C 5 · D 4 · E 5 · F 2 · G 7 | A 3 · B 4 · C 4 · D 2 · E 3 · F 4 · G 6 |
| blockers (re-graded under the canonical rule) | **23** | **23** |

**Matched `m = 12` · union `= 46`** (as run) — **`m = 11` · union `= 35`** (canonical re-grade).

**Chapman `N̂ = (n₁+1)(n₂+1)/(m+1) − 1`:**
- as run: `33 × 27 / 13 − 1 = 67.5` → **`N̂ ≈ 68`**
- canonical re-grade: `24 × 24 / 12 − 1 = 47` → **`N̂ ≈ 47`**

Both are **floors** — heterogeneous detectability biases capture–recapture downward (§5 rule 23).

## THE FINDING, and it is not the number

**The instrument was not frozen. It never was, and nobody could have seen it in a diff.**

§5 rule 22 states that M257x *"froze its instrument at iter-41 and never touched a knob again."* The briefing
that **is** the instrument — its own §"Grading rule — this is the whole instrument" — lived at
`.agentspace/scratch/work-m257x/iter50-briefing.md`: **git-ignored**. It appeared in no commit, no diff, and
no iter directory. iter-53 went looking for it **inside the milestone dir**, did not find it, and
**re-authored the briefing from its description in iter-50's `overview.md`** — which is what *every* prior
pass must also have done, because there was nothing else to do.

So "same briefing" across nine readings has meant *"a briefing re-derived from a one-line summary of the last
one."* That is a knob turning continuously, in the dark, on the one surface §5 rule 22 declared fixed.

The itemized drift is in [`../instrument/README.md`](../instrument/README.md). The load-bearing item is the
**tie-break inversion**: the canonical briefing says *"if you cannot cite the refutation, it is not a
blocker"*; iter-53's as-run briefing said *"when in doubt, book it as a BLOCKER."* The canonical rule resolves
doubt **downward**; the as-run rule resolved it **upward**. Two further carve-outs the canonical rule makes
explicitly — **undercount → MINOR** and **omitted list member → MINOR** — were absent as-run.

**This is §5 rule 22's own warning, arriving from the direction it did not anticipate:** *"every better
instrument found more is a warning, not progress… it has been measuring reach."* Rule 22 assumed instrument
changes would be deliberate. They were not. **The series `25 → 13 → 11 → 17 → 37 → 18 → 7 → 12 → 14 → 7 →
32/26` has an explanation nothing else in this milestone has offered: it is substantially a record of nine
re-authorings of a briefing, not nine measurements of a corpus.**

**Corrective action taken in this iteration:** the canonical briefing is now **committed** at
`../instrument/briefing-canonical-iter41.md`, alongside the as-run one as drift evidence. The instrument is a
versioned file from here on, and a future reading that changes it will change a diff.

## The re-grade, and what it is worth

Every one of the 46 union findings was re-graded against the canonical rule **verbatim** (undercount → MINOR,
omitted list member → MINOR, line drift → MINOR, "cannot cite the refutation" → not a blocker). Result:
**11 of 46 re-grade to MINOR**, giving `23 / 23 / m=11 / union=35 / N̂ ≈ 47`.

**State its weakness plainly.** This re-grade was performed by the orchestrator, **not blind**, with both
readings in hand. It is an adjudication of other seats' findings by someone who knows the answer key. It is
*evidence about the size of the grading confound*; it is **not** a reading, and it must not be reported as
one. A comparable number requires a fresh paired reading at the now-committed canonical briefing.

## The result that survives the confound: RECALL REPLICATES

| experiment | recall #A | recall #B | mean |
|---|---|---|---|
| iter-50 (readings #9/#10, canonical briefing) | 4/14 = **29%** | 4/7 = **57%** | 43% |
| iter-53 as run (readings #11/#12) | 12/32 = **37.5%** | 12/26 = **46.2%** | 42% |
| iter-53 canonical re-grade | 11/23 = **47.8%** | 11/23 = **47.8%** | 48% |

**Per-finding detection probability of a single 7-seat pass sits near 0.45 in every measurement taken of it,
at two different grading rules and on two different trees.** That is the robust quantity. The *count* is
instrument-dependent; the *recall* is not. It is also the quantity clause 5 depends on, and it says the same
thing it said at iter-50, only more firmly: **a single pass returning zero cannot certify a corpus whose
residual is anywhere above single digits.**

Note the canonical re-grade's symmetry — `23/23`, recall identical to three significant figures. iter-50's
prediction 4 (*"the disagreement is roughly symmetric — neither reading is simply better"*) is corroborated
about as cleanly as it could be.

## Induced term

**9 of 46 as run** (U08, U10, U11, U12, U30, U31, U32, U41, U46), classified mechanically against the added
line ranges of `1255998..0e35b1a` plus two recorded judgment calls — method in
[`blocker-ledger.md`](blocker-ledger.md). Under the canonical re-grade, U30 falls to MINOR, giving **8**.

For scale: iter-52 repaired **18 claims** and induced **8–9 new blockers** doing it. That ratio has not
improved under TOK-03 move 3 (*shrink the edit*), and iter-52 already recorded why — deletion was available
for only 2 of 18.

## PRE-REGISTRATIONS — adjudicated, unsoftened

| # | prediction | outcome |
|---|---|---|
| 1 | each of `N₁₁`, `N₁₂` in **[5, 15]** | **REFUTED** — 32 and 26 as run; 23 and 23 re-graded. Outside the band on either grading. |
| 2 | recall **< 60%** for both readings | **HELD** — 37.5% / 46.2% as run; 47.8% / 47.8% re-graded. |
| 3 | `|#11 ∪ #12| > max(N₁₁, N₁₂)` | **HELD** — 46 > 32 as run; 35 > 23 re-graded. Each reading booked findings the other missed on an identical hand. |
| 4 | `N̂` in **[6, 20]** | **REFUTED** — ≈68 as run, ≈47 re-graded. Off by 2–3× on either grading. |
| 5 | induced term **< 8** | **REFUTED** — 9 as run; 8 re-graded, which is not below 8. |

### TOK-03's own pre-registrations, for this iteration

| prediction | outcome |
|---|---|
| **`N̂` below 12** | **REFUTED, decisively** — ≈68 as run, ≈47 under the canonical re-grade. TOK-03 predicted the residual would fall below 12; the best-case reading of this measurement puts it near 47. |
| **induced term below 4** | **REFUTED** — 9 as run, 8 re-graded. |

iter-52 recorded both as *heading toward refutation* and declined to soften them. They are refuted, and the
margin is large enough that the failure is not a near-miss: **TOK-03 move 2 asked to "drive `N̂` down first",
and `N̂` went up.** Whether it went up because the corpus is worse or because the instrument moved is exactly
what the unfrozen briefing makes it impossible to say — which is itself the answer to why move 2 could not
have been managed as designed. **You cannot drive a metric down when its instrument is re-authored between
readings.**

## What this changes for iter-54 — routed, not decided here

1. **The repair target is ambiguous by construction and a human must choose.** Repair the **46** (as-run
   grading) or the **35** (canonical)? The 11 in between are real defects that the canonical instrument calls
   MINOR, and *"YELLOW with 0 blockers"* is admissible under clause 5. Repairing them is not wrong; counting
   them is what is contested.
2. **`N̂ ≈ 47` — or 68 — is 2–3× the ≈23 that TOK-03 called "arithmetically hopeless" for a zero reading.**
   By TOK-03's own move-2 reasoning, clause 5 is further away than when TOK-03 was authored, not closer.
3. **A tok is likely due and this iteration does not pre-empt it.** TOK-03's core prediction is refuted and
   its central metric moved the wrong way. The tik streak is not yet at 3, so the automatic trigger has not
   fired; the judgment belongs to the next planning step, with the instrument finding in hand.
