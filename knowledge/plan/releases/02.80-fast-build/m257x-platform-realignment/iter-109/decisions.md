# iter-109 — decisions

## `D-M257x-109-1` — a seat-commit subject named one seat and carried two. Recorded, not rewritten.

`57d34f6`'s subject reads *"seat C committed VERBATIM on landing"*. It carries **`r27-C.md` AND
`r27-D.md`** — seat D landed in the window between the notification and the `git add`, and a
directory-scoped `git add` swept both.

**The substance of the discipline held and was verified**: neither file was read, edited or graded before
it was committed, and `git show --name-only` on all five seat commits reconstructs exactly which seat
landed in which commit — `ed50c78` F, `bd7b088` B, `9ca8270` A, **`57d34f6` C+D**. The *record* is
recoverable; the *subject line* is wrong.

**Not amended.** Rewriting the commit would erase the evidence of the slip, which is the opposite of what
this milestone is for. A commit subject is a claim like any other, and the correction belongs beside it
rather than on top of it.

**Lesson, and it generalises past this iter:** `git add <dir>` under a concurrent producer stages whatever
has appeared, not what you believe appeared. When the commit *subject* is an assertion about scope, stage
the **named paths**, not the directory — otherwise the message and the content are free to disagree
silently. Same shape as `D-M257x-108-4`'s reach-denominator lesson: state what you measured over.

## `D-M257x-109-2` — the corpus HEAD moved mid-open; the read scope provably did not

Recorded in full in [`ground-truth.md`](ground-truth.md) as a disclosed correction. `2e3443d` → `08cfbd8`
(a concurrent lane closing iter-108) between the 20:10 measurement and the 20:18 seal. Identical
`corpus/services` + `corpus/architecture` tree hashes at `2e3443d`, `08cfbd8` and `ac48e5b`, with
`e6aed2e` as a **firing negative control**. `08cfbd8` is +72/−0, all under `knowledge/plan/**`, which is
barred to every seat. **Subject unmoved; corpus under audit is `ac48e5b`.**

## `D-M257x-109-3` — ARRIVAL vs DETECTION, resolved as DETECTION

The pre-registration named this confound *before* the number existed, which is the only reason it could be
resolved rather than argued. iter-103 measured 61 % of `N` as platform-drift and `TOK-06` read that as an
**inflow**. This reading held all 14 clones at the identical sha — **nothing arrived** — and still measured
**~33 % drift** among upheld blockers. Band #8 (`≤ 25 %`) was cut hard on purpose and **failed**.

**Consequence for the strategy, stated plainly:** `TOK-06`'s premise — *"inflow is comparable to outflow"* —
was a correct measurement of the residual's **composition** read as a measurement of its **flow**. The two
are not the same, and only freezing the subject could separate them.

**What survives:** the induction half was real and its fences worked (**21 % → 5.6 %**, band #10 held at 2
of 36). Steps 0–2 are not reverted and not wasted. What changes is their **rank**.

## `D-M257x-109-4` — repair scope is DETECTION-BOUNDED (the structural finding)

iter-108 repaired by **predicate** and was graded **46/46 = 100 %** of the upheld union. Both facts stand.
Its anchor list was nevertheless derived from `iter-103/raw/` — **what the previous reading detected**.

**A predicate's site list ≠ a reading's detection list.** With per-pass recall on this instrument at
33–83 %, ~100 % reach against a detection-bounded ledger is fully compatible with leaving twins standing.
Two were measured directly (`external_services.md:554` eleven lines from the repaired `:565`;
`ai_architecture.md:34` against the repaired `:95`/`:99`), and **one of them is now a self-contradiction**
because the repair fixed one side of a pair and the other side still asserts the old line.

> **Fixing one site of a pair is worse than fixing neither.** A single consistent falsehood becomes a
> corpus that contradicts itself, which is a strictly harder defect for a reader to resolve.

**Binding on the next repair:** the anchor set is re-derived **from the corpus, per predicate**, never from
a reading's raw output, and `repair_reach_guard`'s denominator must say which set it graded over
(`D-M257x-108-4`'s lesson, one level up).
