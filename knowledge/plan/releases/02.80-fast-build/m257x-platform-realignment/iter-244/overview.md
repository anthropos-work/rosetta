---
iter: 244
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-10
---

# iter-244 — the sixth runnable-input surface: does the tool the doc tells you to run EXIST?

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)
— *census the mechanical classes; stop sampling them.* This iter censuses a class end-to-end; it does not
sample it.

## Step 0 — re-survey before targeting

`TOK-08`'s next-tik direction is "work the classes in descending measured size", and the user's standing
redirect (2026-08-09) narrows it: *"the goal remains alignment and be able to build a working stack with the
new platform repos."* Five **runnable-input** surfaces have now been censused — `make` targets, `cd`
directories, environment variables, frontend scripts + ports, slash-command invocations — and the last of
them returned **8 of 8 wrong**. The direction is not stale; it names a family, and the family has
unexplored members.

**The member picked, and why it is the right one.** Every censused surface so far graded a runnable
*argument*: which target, which directory, which variable, which flag. **Nothing has ever graded the
runnable SUBJECT — the path of the tool itself.** The corpus tells an operator to run
`rext stack-core/demo_knob_guard.py`; if that path does not exist, the instruction is not merely
mis-parameterised, it is unexecutable.

**And the substrate says nothing is watching it.** Measured before sealing:

| | |
|---|---|
| corpus `.md` files scanned | **92** |
| distinct rext paths cited **bare** (no `:NN` pin) | **140** |
| occurrences of those | **250** |
| distinct rext paths cited **pinned** (`path:NN`) | **27** (57 occurrences) |
| distinct, union | **154** |
| bare-**only** (never pinned anywhere in the corpus) | **127** |

`anchor_construct_guard` (FENCE-M257x-iter45) *does* resolve rext paths — but its subject is
**`file:line`**: it asserts the cited **line** carries a construct. A path with no line has no line to
grade, so the 127 bare-only paths are in **no fence's subject**. `corpus_citation_guard`
(FENCE-M257x-iter117) is explicitly scoped to `corpus/...` intra-corpus paths and excludes everything else.

## Hypothesis

A corpus reference to a rosetta-extensions path is a **runnable subject**, and the class is mechanically
decidable: the path exists in the rext tree or it does not. Because rext has been refactored across
eleven sections and 200+ iters while the corpus was repaired by prose sweeps, some fraction of the 140
names a file that has moved or never existed — and because the class carries no fence, no repair pass has
ever had to notice.

## Pre-registered numeric claims — SEALED IN THIS COMMIT, before the census runs

Clone set named, per iter-241: the census resolves against **`.agentspace/rosetta-extensions` @ `c2d9052`**
(the authoring copy, worktree-clean), and the reach it reports is a property of that clone set.

| id | claim | prediction |
|---|---|---|
| **P-244-1** | distinct bare-referenced rext paths that do **not** resolve to a file in the rext tree | **≤ 12** of 140 |
| **P-244-2** | at least one non-resolving path sits in a **runnable position** — inside a fenced block or immediately after `rext ` | **YES** |
| **P-244-3** | of the **27** paths the corpus cites *with* a line pin (already in `anchor_construct_guard`'s subject), the number that fail to resolve as files | **0 of 27** |
| **P-244-4** | the dominant failure mode is a **rename/move** (a file of that basename exists elsewhere in the tree) rather than a name that never existed | **rename dominates** |
| **P-244-5** | guards in the `stack-core` family that enumerate **bare** rext path references today | **0** |

**Falsification that matters:** if P-244-1 comes back **0**, the class is empty, the fence is not built, and
the iter closes `closed-no-lift` with the census as its deliverable — a censused-to-zero class is a real
result under `TOK-08`, and the substrate table above is what makes the zero believable.

## Phase plan

1. **A** — build the census instrument; enumerate the population; prove the instrument on a known-good and
   a known-bad input (anti-vacuity).
2. **B** — read it. Record the population, the failures, and the reach **with its denominator**.
3. **C** — repair every non-resolving citation at every site (not the first one found).
4. **D** — fence it, so the class stays at zero: a guard with a mutation control and an anti-vacuity
   control, wired into `guard_family`.
5. **E** — re-derive every number **after the last edit**, per the standing rule.

## Escalation conditions

- A non-resolving path whose correct target is **ambiguous** (two plausible files) → do not guess; record
  the ambiguity and route it.
- The census's own instrument disagrees with a second derivation → per iter-175, **union, never
  substitute**.

## Acceptable close-no-lift outcomes

- The census returns **0 non-resolving paths** with a proven-non-vacuous instrument.
- The class turns out to be already covered by an existing fence under a different spelling — in which case
  the deliverable is the measured reach statement, not a second fence.
