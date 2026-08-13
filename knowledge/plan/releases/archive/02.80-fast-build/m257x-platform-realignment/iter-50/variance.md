# The variance experiment — two readings, one tree, no repair between

**The experiment `platform-alignment.md` §5 rule 22 prescribes and this milestone had never run.**

> *Measure the variance FIRST, by reading the same tree twice with no repair between; that is a cheap
> experiment and this milestone paid eight passes to learn it.*

It cost **one** reading, because reading #9 was already the first half — and it was available exactly once,
in the window between iter-49's reading and the repair of its 14. That window is now spent.

## The control — what was held identical, and how it was verified

| held fixed | verification |
|---|---|
| the 40 audited files | `git diff --stat 47c9b7d..HEAD -- corpus/ CLAUDE.md .claude/` → **empty**. The only commits between are `knowledge/plan/**` |
| the ground truth | all 13 clone shas re-read at open, each identical to the ones #9's seats recorded |
| the partition | corpus unchanged ⇒ the size-sort deals the **same hand**. Cross-checked name-for-name against iter-47's published table: all six sets match |
| seat count, briefing, grading rule, read-in-full discipline, per-file `wc -l` control | replicated |
| seat G's subject | the identical diff — `2fc633a..47c9b7d -- corpus/`, which is exactly the working-tree diff G read at #9, now committed |
| **blindness** | fresh seats; no seat told a prior reading existed; every seat barred from `knowledge/plan/**` |

**There is no partition confound, no ground-truth confound and no corpus confound.** Whatever the two
readings disagree about is the instrument.

## The result

|  | reading #9 (iter-49) | reading #10 (iter-50) |
|---|---|---|
| blockers held | **14** | **7** |
| per seat | A 1 · B 1 · C 2 · D 2 · E 3 · F 0 · G 5 | A 1 · B 0 · C 0 · D 1 · E 0 · F 0 · G 5 |
| induced / pre-existing | 7 / 7 | 4 / 3 |

**Matched pairs (found by BOTH): 4 of #9's 14.**

| #9 | #10 | claim |
|---|---|---|
| #9 | 4 | `jobsimulation.sessions` asserted to exist, unqualified, in a local-dev section |
| #10 + #11 | 3 | the `flag_use_realtime_openai` leak (#9 split the three sites into two rows; #10 filed one) |
| #13 | 7 | `token` is *the only* required-and-undefaulted column |

**Found by #9 only: 10.** #1, #2, #3, #4, #5, #6, #7, #8, #12, #14.
**Found by #10 only: 4.** #1 (the `.env` variable names), #2 (studio-room's *"only outbound API call"*),
#5 (`hiring.md:241-242`, the opposite-polarity twin of a claim #9 raised), #6 (`dependency_map.md:19`).

### Every pre-registered prediction, graded

| # | prediction | result |
|---|---|---|
| 1 | count in **[9, 19]** | **REFUTED** — 7. The ±5 band around 14 does not contain it; the *series* range 7–18 does |
| 2 | **fewer than 7 of 14 re-found** (recall < 50%) | **HELD** — 4 of 14, **recall 28.6%** |
| 3 | **union > 14** | **HELD** — **18** |
| 4 | disagreement roughly **symmetric** | **REFUTED** — 10 vs 4. Both directions non-empty, but #9 was 2.5× richer |

Two of four, and the two that held are the two the experiment was for.

## What the numbers estimate

Two independent samples from one population with a shared per-finding detection probability is a
capture–recapture design. With `n₁ = 14`, `n₂ = 7`, `m = 4`:

- **Lincoln–Petersen:** `N̂ = n₁n₂/m` = **24.5**
- **Chapman (the less-biased form):** `((n₁+1)(n₂+1)/(m+1)) − 1` = **23**

> **The tree that reading #9 measured at 14 and reading #10 measured at 7 is estimated to carry ~23–25
> blockers. Two full 7-seat readings named 18 of them between them, and neither named more than 14.**
> Implied per-reading recall: **58%** and **29%**; mean ≈ **43%**.

Stated as the caveat it deserves: capture–recapture assumes the two samples are independent and that every
finding is equally detectable. Neither is exactly true — seat G's subject overlaps both readings' highest-
yield surface, and some defects are plainly easier to see than others. **Heterogeneous detectability biases
`N̂` DOWNWARD**, so ~23–25 is a **floor**, not a point estimate. The direction of the error is the
uncomfortable one.

## What it explains — the series was never noise around zero

The five readings taken at the frozen instrument read `18 → 7 → 12 → 14 → 7`, and the milestone has spent
four iterations treating the variation as noise around a shrinking residual. It is not.

**A repair pass can only repair what a reading NAMES.** If a reading names ~43% of the pool and the repair
of what it names induces new defects at the rate this milestone has measured (9, then 7, then 2, then 7,
then 4), the process has a **fixed point**, and it sits exactly where the series has been sitting.

That reframes four iterations of results at once:

- **iter-47's "zero pre-existing" is fully explained** — not a converged corpus, a low-recall draw. iter-48
  then booked ten, seven of them months older than the milestone, sitting in the file sets of the pass that
  had reported zero. This experiment reproduces that mechanism under control.
- **"Every better instrument found more" (rule 22) is the same phenomenon** seen from the other side.
- **The repair-then-read cycle cannot reach zero by iteration**, because each turn leaves behind the ~57%
  the reading did not see, and adds to it.

## What it says about clause 5 — precisely, and without re-opening it

Clause 5 is met by a reading that returns zero blockers. The user has ruled twice; this iteration does not
re-open the clause, does not propose re-cutting it, and does not propose closing at 4 of 5.

What the measurement adds is an arithmetic that was previously unavailable:

1. **A zero reading is not cheap to obtain by luck.** With recall ≈ 0.43, the chance of a single pass
   missing all of a residual of 23 is ~`(1−0.43)²³` ≈ 1 in 10⁵. **So a zero, if one ever arrives, is
   strong evidence** — this refutes the weaker reading of rule 22's *"a zero is the least trustworthy
   reading"*. It is untrustworthy relative to a residual of ~7, not relative to a residual of ~23.
2. **But the residual must actually approach zero for a zero to be drawable**, and the current cycle
   converges to a fixed point around 20, not to 0.
3. **Therefore the binding constraint on clause 5 is not the reading. It is the repair's coverage.** The
   gate asks for a zero; the method supplies repairs for 43% of the pool per pass.

That is a statement about the *method*, which the user's ruling did not fix, and it is the input the next
strategy revision needs.

## The instrument this hands forward

The paired design is itself the deliverable, and it is cheap once the discipline exists:

- **Two blind readings of the same tree cost one extra reading and yield three things a single reading
  cannot**: a recall estimate, a residual estimate with a known-direction bias, and a **union** that is
  strictly better than either reading alone.
- **The union is the right repair target.** Repairing #9 ∪ #10 = 18 covers 78% of the estimated pool,
  against 61% for #9 alone and 30% for #10 alone.
- **`N̂` has a floor of zero by construction** — the property rule 22 asks a gate metric to have. When two
  independent readings of an unchanged tree agree on nothing because there is nothing to find, `m`, `n₁`
  and `n₂` all go to zero together. It is not proposed here as a replacement for clause 5; it is recorded
  as the estimator the milestone did not have.

## The sharpest single observation

Three seats in reading #10 independently re-derived the `31 of 135` org-filtering count, each recorded it
as a **positively audited zero**, and all three were **wrong** — the true count is 32, because
`schema/organization.go:56` declares its own `Policy()` with `rule.FilterSameOrganizations()` and uses
neither mixin (measured at adjudication: 30 mixin users, 4 own-`Policy()` schemas).

Each seat verified the **arithmetic the document showed** rather than the **predicate the document
claimed**. That is §5 rule 17, written by this milestone, violated three times in one pass by auditors
briefed on it.

> **An audited zero can be wrong, and a wrong audited zero is worse than a silence** — a silence is
> uninformative; a wrong audited zero is evidence pointing the wrong way, and it is exactly what a pass
> reporting zero blockers is made of.
