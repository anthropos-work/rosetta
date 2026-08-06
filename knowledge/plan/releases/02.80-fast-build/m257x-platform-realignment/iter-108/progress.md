**Type:** tik · `iter_shape: repair` · **`TOK-06` step 3** — *repair the 33, and the repair is now watched.*

# iter-108 — the union is paid, by predicate, under two fences that both fired on it

## The tok-trigger fired on its literal words and was resolved, not waved through

The last three tiks — **105, 106, 107** — are three consecutive tiks with **no `N` movement**, which is
Phase 0 rule 2's streak on its face. It does **not** fire, for a reason that is a measurement distinction
and not a convenience: rule 2 defines no-progress as *"the metric did not move … (zero or net-negative
delta)"*, and **a delta requires two measurements. Those three iters took none** — `TOK-06` puts the read at
step 4 **on purpose**, and each of the three says in its own close that no `N` movement is claimed. The
metric is **UNMEASURED, not unmoved.**

Grading "not measured" as "did not move" asserts something nobody measured — §8's *grade the cannot-tell*,
pointed at the skill's own trigger. The substantive check agrees: a triggered tok revises a **stalled**
strategy, and `TOK-06` is **3 of 5 steps in, on schedule, with its two metric-moving steps not yet run**.
Revising it here would revise it *before its own evidence exists*. `D-M257x-108-1`, and **codified in the
protocol doc** so the next agent inherits a rule rather than re-deriving a judgement call.

## What was repaired

**22 predicates**, by **predicate** (TOK-05's unit, vindicated by iter-103's band #3), across **23 files**.

**The anchor list was DERIVED, never hand-assembled** (§5 rule 19) — through
`repair_reach_guard.read_ledger()` over `iter-103/raw/`, *the same code path that grades the repair*, so the
input and its grader cannot disagree: **48 booked blocks / 14 seat reports → 31 distinct primary anchors**.
Published as [`derived-ledger.md`](derived-ledger.md). **A hand list would have missed
`shared_libraries.md:128`** — the derivation caught it, and it turned out to be the one finding worth the
most (below).

| class | predicates | what changed |
|---|---|---|
| **drift** | 9 | colony split **CLOSED** (`app` + `sentinel` both `v0.35.2`) · proto skew is **ZERO**, not two · `next` is **`~16.2.12`** — a *tilde*, not `^16.2.7` · `@clerk/nextjs` is **NOT aligned** (3 × `^6.39.6` vs ant-academy `^6.39.2`) · clerkenstein's *"sentinel is still on `v0.34.3`"* softener is false · 5 line-anchors re-derived |
| **iter-102-induced** | 4 | the `:8081` cardinality (5 sites) · the `:321` citers (4 sites, incl. `backend.md:54`) |
| **never-true** | 4 | `storage.md` never contained the text `dependency_map` credited it with · `cert.py` is the **certification-benchmark CLI**, not signing · there is **no "academy subgraph"** (5 sites) · a demo academy does **not** serve its committed FS catalog |
| **unclassified** | 5 | mistral constructor · jobsim wiring fatal · `.env.example` line · hiring's non-uniform offset · the dropped prod-terraform assertion |

## The two induction shapes were refused BY CONSTRUCTION, not by intention

**1 — the canonical wording is not re-multiplied.** iter-102 published *"one occurrence anywhere in the
clone set"* to **5 anchors**; the literal has **six**, five inside the sentence's own denominator.
Re-derived here **before a word was written**: **6 occurrences** (`app` 1, `rosetta-extensions` 5), **0 in
any `.tf`**, **44 tracked `.tf` across 13 repos** — the old sentence was self-refuting at its own stated
scope. The fix is **structural, not editorial**: the count is now derived **once**, in `backend.md`, and
`cms.md` / `jobsimulation.md` **point at it**. *A pointer cannot carry a false cardinality to five places.*
The multiplier is not corrected — it is removed.

**2 — prose above a cited line re-points the citers.** The single edit that added lines
(`architecture_overview.md`'s markdownManager correction) was made **first**, the target **re-measured
after** (`:335`), and all four `:321` citers re-pointed — including **`backend.md:54`, which the 14-seat
double reading missed in BOTH passes**. Every other edit to that file was made **one-line-for-one-line** so
it could not shift again.

## Both fences fired ON THIS REPAIR — twice — and that is the deliverable

`TOK-06` put repair at step 3 *"precisely so the repair has something watching it."* It did.

- **`repair_postcondition` refused the repair commit itself.** §8's iter-102 post-mortem carried bare
  `:321`/`:331` numbers; once the wording moved to `:335` those became citations onto a **blank line**,
  indistinguishable from live ones. Post-mortem line numbers are now written **with their ref**.
- **`anchor_offset_guard` went RED on the repair, one commit old** — 2 rotted citations at
  `platform-alignment.md:1844`/`:1864`, both pointing at `shared_libraries.md:85`, whose claim the repair
  had just fixed. **§8's own write-up of the drift fence quoted the drift as live corpus text.** Repaired
  in `e688843`.

**The class landed on the document that teaches the rule, one iteration after the fence shipped. Fifth
occurrence in four days, and the fifth time what caught it was re-running rather than reasoning.**

## Grading — by machine

| instrument | verdict |
|---|---|
| `anchor_offset_guard` (`f8be5a1^..e688843`) | **OK** — 13 graded, **0 rotted**, 7 CANNOT-TELL |
| `repair_reach_guard` (same range, ledger = `iter-103/raw/`) | **raw reach 46/47 = 97.9 %** |
| `clone_drift_guard` | **OK** — 14/14 clones, and the **2 gradeable pins now MATCH** |
| `repair_postcondition` (pre-commit, both commits) | **OK** — publishes no adjudicated claim the baseline lacked |

**The 7 CANNOT-TELLs were each checked by hand**, as the guard's own OK line instructs — the guard reports
that class rather than asserting it (`D-M257x-107-2`). **All 7 are correct post-move numbers**:
`architecture_overview.md:335` (all four citers), `backend.md:302`, `backend.md:54`, and the deliberately
historical `hiring.md:93`.

### The residue is a REJECTION, and that is the honest number

The single unreached booking is **`shared_libraries.md:128`, `r25-G B3`** — and adjudicator 4 **REJECTED**
it, class **`wrong-tree`**: it graded `app/internal/ai/` at `ad9f3c49` (app's post-fold diverged fork)
instead of the `ai` module at the `v1.40.2` the section's own pin row names, *readable in the same clone at
`1e457fa70`*, where all three booked claims hold verbatim.

> **Raw reach 46/47 = 97.9 %. Reach over the repair's actual input — the UPHELD union — is 46/46 = 100 %.**
> `repair_reach_guard` grades all **48** booked blocks, not the 22 upheld predicates, so its denominator
> includes findings the adjudicators threw out. **This is iter-102's residue result reproduced exactly**:
> the apparent miss is a claim that came out **true**.

## What did NOT get repaired, deliberately

- **`hiring.md:93`** — a **historical** anchor (*what iter-39 found*). Re-pointing it would falsify the
  record. **A guard cannot tell a live citation from a record of where something once was**, which is why
  it reads as ROT and is left standing with its ref made explicit.
- **3 of the 5 iter-107 rotted citations** — re-checked individually and they **resolve correctly today**;
  only `backend.md:241` was genuinely rotted. Bulk-bumping all five would have broken three working
  citations. **The 5 were graded one at a time, not applied as a patch.**

### `backend.md`'s four-address anchor has now rotted THREE times

`:187` (→ iter-98) → `:241` (→ iter-103) → **`:302`** today. Re-pointed, and the passage now says
**cite the claim, not the line** — an anchor with a measured recurrence rate is a design problem, not an
accident.

## `FIX-M257x-iter107-drift-fence-satisfiable-by-prose` — NOT closed, and why

The brief asked for a **derived** discriminator rather than an exclusion list. **I could not find one, and
I am not shipping a list.**

The fence's D1 asserts *"every cited clone's HEAD is a commit the corpus cites."* A doc that merely
**mentions** a sha satisfies it. Separating *"this sha dates a claim"* from *"this sha is being talked
about"* is a question about **what the author meant by writing it** — and **intent is not in the
repository**, which is the identical wall `D-M257x-107-2` hit when `anchor_offset_guard` could not separate
*post-move-and-correct* from *pre-move-and-stale*. Every discriminator I could construct (require an `@`,
require an adjacent `file:line`, require a table cell) is a **shape allow-list wearing a derivation's
clothes** — §2's hand-maintained tuple, which this milestone has rejected twice.

**Left open, with the known-limitation test still pinning it.** Note this iter's green is nonetheless
**earned rather than satisfied-by-prose**: the drift was *actually repaired*, and the 2 gradeable pins
**match**.

## Close — 2026-08-06

**Outcome:** the read union is paid by predicate — 22 predicates / 23 files; machine reach **46/46 = 100 %
of the upheld union** (raw 46/47, the miss a REJECTED `wrong-tree` finding); `anchor_offset_guard` **GREEN**
on the repair's own range and `clone_drift_guard` **GREEN with its pins matching**. Both new fences **fired
on this repair before it was allowed to stand**, which is what step 3 was sequenced after step 2 to test.
**No `N` is claimed** — measuring here would be repair inside the measuring pass; that is step 4.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-108-1` (the tok-trigger resolution + protocol codification) · `D-M257x-108-2` (the
`:8081` multiplier removed structurally, not corrected) · `D-M257x-108-3` (the 5 rotted citations graded
individually; historical anchors are not re-pointed) · `D-M257x-108-4` (reach is reported twice — raw, and
over the upheld union) · `D-M257x-108-5` (the drift-fence limitation stays open; no exclusion list)
**Side-deliverables:** `rosetta-extensions` `680e852` — **`anchor_offset_guard` false-greened its own answer
key on a bare rev.** `git diff <sha>` is *sha vs the working tree*, so the bare form graded **0 of 33
citations seen and printed OK**, where `<sha>^..<sha>` is **RED with 10 findings**. All 18 existing tests
passed throughout because every one of them used the explicit form. Fixed by normalizing (the semantics
`repair_reach_guard` already had); **+3 tests, mutation control verified firing**. Found by pre-flighting
the fence that was about to grade this iter — Phase 0d paying for itself.
**Routes carried forward:**
- `FIX-M257x-iter107-drift-fence-satisfiable-by-prose` → **Fate 3** → a future iter: no derived
  discriminator found; reasoning recorded above; **no exclusion list added**.
- `FIX-M257x-iter103-read-union` → **DISCHARGED** by this iter (100 % of the upheld union).
- `FIX-M257x-iter107-unbooked-rot` → **DISCHARGED** — 1 of 5 repaired, 3 verified already-correct,
  1 ruled historical-and-must-not-move.
- **`TOK-06` step 4 (the read)** → the next iter. Unstarted by design.
**Lessons:**
1. **A fence you shipped last iter will fire on the repair you ship this iter, and it will be right.**
   Both did. Budget for it: a repair commit is not done when it is written, it is done when the fences
   that watch it are green.
2. **Remove a multiplier structurally; do not correct it.** A corrected canonical sentence is still a
   canonical sentence, and the next author will re-multiply it. A *pointer* cannot carry a cardinality.
3. **Grade a bulk finding one item at a time.** 3 of the 5 "rotted" citations were fine, and one must
   never move. A five-item route is five findings, not one.
4. **A bare rev is not a range** — and the invocation an operator types first was the one no test covered.
