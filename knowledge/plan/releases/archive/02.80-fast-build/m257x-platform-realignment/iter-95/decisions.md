# iter-95 decisions

## `D-M257x-95-1` — the reading was TAKEN; N = 13; the gate does not move

Runs 54–58 deferred the reading four times, the last time on a reason iter-80 had already named as a
conflation: **clause 3's instrument is the guard family; clause 5's is the graded READ.** The READ
instrument was never touched (sha256 `3858ec53…`, one commit ever). Preconditions re-derived, not
inherited. Reading taken. **N = 13.** Clause 5 stays open, gate stays **4 of 5**.

## `D-M257x-95-2` — the graded set is scoped to the clause, and the basis change is declared

Clause 5 reads *"GREEN, or YELLOW with 0 blockers, over `corpus/services/**` + `corpus/architecture/**`"*.
So N counts **upheld · BLOCKER-grade · in-scope · deduped**. Minors do not block by the clause's own
words; `CLAUDE.md`, `corpus/ops/**` and `.claude/skills/**` findings are real, reported, and outside
the clause.

**This is a different basis than the `140 → 43` series used** (all upheld findings). Recorded as a
**declared re-baseline of the count**, alongside a **continuous upheld rate** (92.1 → 93.0 → 92.7 %).
The alternative — quoting 51 to look continuous with 140 — would have compared two different questions.

## `D-M257x-95-3` — `storage.md:58` is counted in N, and the hold is reported both ways

`DEF-M257x-iter80-storage-prod-bucket` holds `:55`/`:154`/`:181`. Re-derivation found `:55` **correct**,
`:156`/`:181` **historically fenced**, and **`:58`** — not among the held lines — the only present-tense
false sentence. It is counted **in** N, with **N = 12 stated explicitly** for the reading where the user
intends the hold to cover the hazard class rather than the three named lines. **The user's item, the
user's call; not decided here.**

## `D-M257x-95-4` — an absence is established only by `git grep` at a named ref

The environment's recursive `grep` is `ugrep --ignore-files`; a `.gitignore` entry silently hides
**tracked** files. `grep -rn mistralai app/studio/` → 1; `git grep mistralai <ref>` → 2. This produced a
false clearance in this very reading and very likely authored the false corpus claim originally.

**Every absence-claim in this milestone taken with a recursive `grep` is suspect, and the bias is toward
under-counting.** Routed to the protocol as a §5 rule. It also means **N is a floor twice over**: once
from capture–recapture, once from the search tool.

## `D-M257x-95-5` — 6-of-6 pre-registration is a warning, not a win

Every prediction held. iter-76 graded 2 of 5 and iter-53 graded 2 of 5, and both learned more.
Prediction 1's band `[0,12]` was nearly unfalsifiable and prediction 4 landed within a tenth of a point.
**Tighten the bands next reading** — a prediction written to be safe is not a prediction, and this
pre-registration drifted toward safe.
