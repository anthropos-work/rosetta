**Type:** tik — the paired reading, under [`TOK-03`](../decisions.md#tok-03-repair-the-union-shrink-the-estimator-make-the-edits-smaller--2026-08-03) move 1.

## Work

- **Step 0 re-survey.** Platform origin re-fetched: `2adcf714`, **unchanged** — re-scope trigger stays at
  occurrence 1 of 2. All 13 ground-truth clones re-read and byte-identical to iter-50's. Scope re-enumerated:
  **40 files / 9,395 lines** (was 9,326 — iter-52's repair added 69 net lines).
- **Partition recomputed** from current line counts by the fixed method (size-sort descending, snake-dealt
  A→F then F→A). The hand moved — `security_compliance` A↔B, `graphql-wundergraph` A→B, `clerk-integration`
  B→C, `messenger` C→B — because the corpus changed, which is the honest consequence of a fixed method over a
  moving tree and is recorded rather than engineered away.
- **14 blind seats launched in parallel** — readings #11 and #12, identical hand, identical diff for seat G
  (`1255998..0e35b1a`), every seat fresh, barred from `knowledge/plan/**` and from `.agentspace/scratch/`,
  and told nothing of the other reading. All 14 raw reports in [`raw/`](raw/); every seat reported per-file
  `wc -l` positive controls and no seat skipped a file.
- **Adjudicated** to [`blocker-ledger.md`](blocker-ledger.md) (the 46-row union, each row anchored, each
  flagged induced/pre-existing by a mechanical rule) and [`variance.md`](variance.md) (the arithmetic).
- **Induced term classified mechanically** against the added-line ranges of the repair diff, with the two
  judgment calls named.
- **The instrument recovered and committed** — see below.

## Result

| | reading #11 | reading #12 |
|---|---|---|
| blockers (as run) | **32** | **26** |
| per seat | A 2 · B 7 · C 5 · D 4 · E 5 · F 2 · G 7 | A 3 · B 4 · C 4 · D 2 · E 3 · F 4 · G 6 |
| blockers (canonical re-grade) | **23** | **23** |

**as run:** `m = 12` · union **46** · recall **37.5% / 46.2%** · Chapman **`N̂ ≈ 68`**
**canonical re-grade:** `m = 11` · union **35** · recall **47.8% / 47.8%** · Chapman **`N̂ ≈ 47`**
Both floors. **Induced: 9 of 46 (8 re-graded)** — iter-52 repaired 18 claims and induced 8–9 doing it.

## The finding — it is not the number

**The instrument was never frozen, and nobody could have seen it.** §5 rule 22 asserts M257x *"froze its
instrument at iter-41 and never touched a knob again."* The briefing that **is** the instrument lived in
`.agentspace/scratch/work-m257x/iter50-briefing.md` — **git-ignored**, in no commit, in no iter dir. This
iteration looked for it inside the milestone directory, did not find it, and **re-authored it from its
one-line description in iter-50's `overview.md`** — which is what every prior pass must also have done,
because there was nothing else available.

The drift concentrated in one clause. Canonical: *"if you cannot cite the refutation, it is not a blocker"*,
with **undercount**, **omitted list member** and **line drift** carved out as MINOR explicitly. As-run:
*"when in doubt, book it as a BLOCKER"*, with none of them carved out. **The canonical rule resolves doubt
downward; the as-run rule resolved it upward.** Re-grading the union against the canonical rule verbatim
takes 46 → 35 and `N̂` 68 → 47, so roughly half the jump over #9/#10 is grading drift and half is not —
**and neither half is a statement about the corpus.**

**Corrective action, landed in this iteration:** the canonical briefing is committed at
[`../instrument/briefing-canonical-iter41.md`](../instrument/briefing-canonical-iter41.md) with the as-run
one beside it as evidence; the protocol gains **§5 rule 25** (*an instrument that is described rather than
stored is not frozen*); and rule 22's false "never touched a knob" sentence is corrected in place.

## The result that survives the confound

**Recall replicates at ≈0.45 in every measurement ever taken of it** — 29%/57% (iter-50), 37.5%/46.2% (as
run), 47.8%/47.8% (re-graded): means of 43%, 42%, 48%, across two grading rules and two trees. **The count is
instrument-dependent; the recall is not.** And the canonical re-grade's `23/23` symmetry corroborates
iter-50's prediction 4 — neither reading is simply better — about as cleanly as it could be.

## Close — 2026-08-03

**Outcome:** the paired reading returned **32 and 26**, matching on 12, union **46**, `N̂ ≈ 68` (≈47 under a
canonical re-grade) — **`N̂` went UP, not down** — and the reason it cannot be attributed to the corpus is the
iteration's real finding: **the "frozen" instrument was a git-ignored file that every pass re-authored.** It
is now committed, and the protocol says so at §5 rule 25.
**Type:** tik
**Status:** **closed-fixed** — the planned scope was *take two blind readings at the frozen instrument,
adjudicate the union, recompute `N̂`, repair nothing*. All of it landed: 14 seats, all files read under
positive controls, the union anchored, the induced term classified mechanically, `N̂` re-derived with its
method stated, and the fixture left intact. Four of five pre-registrations were refuted, which is a recorded
result of the planned work and not a failure to do it.
**Gate:** **NOT MET** — clause 5 requires a reading that returns zero; this iteration deliberately took no
gate attempt and the residual estimate rose. Gate stands at **4 of 5**.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik; the streak stands at 1, since
iter-52 closed-fixed) — (3) re-scope: n (platform origin `2adcf714` at open **and** close; occurrence 1 of 2)
— (4) **user-blocker: y** — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n —
Outcome: **exit-4**
**Decisions:** [`D-M257x-53-1`](decisions.md) … `-53-5`.
**Side-deliverables:** the committed instrument (`../instrument/`) — recovered canonical briefing + as-run
drift evidence + a README naming the itemized drift. Protocol-evolution: `platform-alignment.md` **§5 rule
25** added and **rule 22** corrected in place, shipped in the same commit as the lesson.
**Routes carried forward:**
- **`FIX-M257x-iter53-union-set`** — the 46 (or 35), anchored in [`blocker-ledger.md`](blocker-ledger.md).
  **The target count is a user decision** (`D-M257x-53-5`) — see the user-blocker below.
- **`CHECK-M257x-iter53-instrument-was-never-frozen`** — every count this milestone has published
  (`25 → 13 → 11 → 17 → 37 → 18 → 7 → 12 → 14 → 7`) was taken with a re-authored briefing. **The series is
  not a comparable series.** Somebody must decide whether the milestone's narrative needs re-baselining, or
  whether the recall figure (which does replicate) is the only quantity worth carrying forward.
- Unchanged and still open: **`FIX-M257x-iter52-mirror-pair-leak`** (40 sites / 12 files, including a
  `verification.md` **gate floor asserting rows in a DROPPED table** — a live gate checking something that
  cannot exist) · **`CHECK-M257x-iter52-second-ai-manager`** (still unconfirmed; do not repair on one seat's
  word) · `FENCE-M257x-iter52-stem-collision` · `FIX-M257x-iter52-alignment-doc-s6` ·
  `FENCE-M257x-iter50-consecutive-audit-mode` · `CHECK-M257x-iter50-audited-zero-is-evidence` ·
  **`CHECK-M257x-iter35-seeder-writes-one-instant`** (highest-value non-gate item) · RF-13 · RF-2/3/7–12 ·
  harden residue iters 27–30, 32–34, 36–41 · `CHECK-M257x-iter38-ai-act-classification` (needs an owner
  outside this milestone) · the root `CLAUDE.md`, still outside the 40-file partition so no reading books it.

**Lessons:**
1. **An instrument kept as prose is re-sharpened every time it is used.** §5 rule 22 told this milestone to
   distrust a count when it *deliberately* sharpens the instrument. The real failure was worse and quieter:
   the instrument had no stored form, so "held fixed" was a claim about nothing, and nine readings measured
   nine re-authorings. **Store the instrument as a versioned file before building any metric on it.** Now
   §5 rule 25.
2. **git-ignored is invisible to every guard this milestone owns.** The same root cause iter-01 found for the
   *pins* (`.agentspace/rext.tag` git-ignored, so drift never appeared in a diff) was sitting under the
   *audit* the whole time. **When a milestone finds that a git-ignored file broke its detection, sweep for
   the others in the same breath** — the class recurs across surfaces, not just across releases.
3. **Recall replicates; counts do not.** Three measurements at two grading rules on two trees all put a
   single 7-seat pass near 0.45. When one quantity survives a confound that destroys another, build on the
   one that survived.
4. **A tie-break clause is a bigger knob than the whole rest of a briefing.** One sentence — resolve doubt
   downward vs upward — moved the count 46 → 35 and `N̂` 68 → 47, more than any other difference between the
   two briefings. Grading rules should state their tie-break explicitly and treat it as load-bearing.
