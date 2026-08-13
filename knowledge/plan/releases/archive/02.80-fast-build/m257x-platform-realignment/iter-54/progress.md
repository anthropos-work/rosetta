# iter-54 — TOK-04

**Type:** tok (triggered — by the milestone's own `re_scope_trigger`, occurrence 2 of 2, plus a direct
user ruling; **not** by the 3-no-prog streak, which was checked and does not apply — iter-53 closed
`closed-fixed`).

Three jobs, in the order the user set them: **re-survey the platform → reassess honestly → author TOK-04.**

## Job 1 — the platform re-survey (committed at `6485151`)

`stack-demo/platform` fast-forwarded `2adcf714 → ef32d4cd` (0 behind origin). Zero platform-repo edits.
The clause-3 fence run **before** any corpus edit went **RED on 3 real direction-B departures** — cms,
jobsimulation, roadrunner leaving `repos.yml` — then GREEN once the map was updated. Narrative:
[`platform-before-after.md`](platform-before-after.md); canonical rows:
`corpus/architecture/platform-migration-status.md`. **81 new drift sites across 21 files, from 3 commits.**

## Job 2 — the honest reassessment ([`reassessment.md`](reassessment.md))

**The gate is 2 of 5 against origin HEAD, not the booked 4 of 5.** Nothing regressed; two clauses went
stale by the gate's own wording.

| clause | verdict | note |
|---|---|---|
| 1 — three green cold cycles | **STALE** | met at `2adcf71` (iter-14, `progress.md:19`). `d11a403` removed 3 compose services; `6060315` changed the bring-up timing contract (`start_period: 120s`). ~35 min to restore |
| 2 — Playthroughs `30/0/0` | **STALE**, and it **recorded no ref** | met at iter-37, whose progress file contains **no platform sha** — ref inferred from iter-36. ~5 min to restore (iter-32: 4 m 50 s incl. reset) |
| 3 — the fenced map | **MET**, re-met today | first **non-staged** catch a milestone fence has made. Caveat: two false prose claims found in the same pass, one written by this iter — the guard fences *membership*, not prose, and says so in its own header |
| 4 — zero writes to a dead schema | **MET — under test** | the derived set tracked the removal with **zero human action**. Verified live: pairs `app:public`, schemas `extensions sentinel public`, transitional debt **empty** |
| 5 — KB-fidelity | **NOT MET, further away** | series non-comparable (unfrozen instrument); recall ≈43–48%; union 46 (as-run) / 35 (canonical), 9 of 46 induced; **+81 today**. Net **−72** for the cycle |

**The two halves of the predecessor's dying claim were both verified, and both held with a correction.**
Clause 4 did self-correct with zero human action — proven by *running* the derivation against the new
`repos.yml`, not by reading it. And the over-claim was real: *"the armed failure is now armed"* cited
`migrate-demo.sh:81-85` / `:106`, code that this milestone's **own iter-02** (rext `54bccf7`) deleted.
Corrected in two files, visibly (`D-M257x-54-3`). A second dead claim was found beside it — §5's
`storage`/`messenger` watch-signal was already true and could never fire (`D-M257x-54-4`).

## Job 3 — TOK-04, in the milestone-root `decisions.md`

**`TOK-04: pin the target, or stop calling it a measurement`** — the pinning-and-tracking policy the
`re_scope_trigger` prescribed by name.

Four rules: **P1** every measurement states its refs in the artifact · **P2** every instrument is a
committed file, and nothing it depends on lives on a git-ignored path · **P3** the platform ref is chosen,
recorded and re-checked at open *and* close, and the iter that detects a move re-points it *in that iter* ·
**P4** derive, else fence, else declare it prose-under-review — an order now **measured on a single event**
(derived: zero human action · fenced: caught it unaided 3/3 · hand-maintained prose: falsified in a day).

The class is generalized beyond the platform repo because it has bitten **four** of this milestone's own
instruments: the rext pin (git-ignored), the audit briefing (git-ignored), the platform clone (free to
move), and — found today — **clause 2's gate-meeting run, which recorded no ref at all**.

TOK-03's three moves are all kept (union-of-two, pre-commit double-reads, smaller edits); its **premise** is
refuted — the residual is a flow, not a stock.

## Close — 2026-08-03

**Outcome:** TOK-04 authored — a pinning-and-tracking policy covering **every** movable input, not just the
platform repo; the gate honestly re-read at **2 of 5** against origin HEAD; clause 4 verified as the
milestone's first thesis-level result tested by an unarranged event; and a false claim this iter committed
30 minutes earlier found by re-measurement and corrected in place.
**Type:** tok
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: **y** — (3) re-scope: n (already fired at iter-53; this tok IS its answer) — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: **exit-2**
**Decisions:** D-M257x-54-1 … D-M257x-54-7 (iter-local); `TOK-04` (milestone-root)
**Side-deliverables:** none — the two claim corrections are in-scope repair of job 1's own output, not
unrelated side-fixes.
**Routes carried forward:**
- **iter-55** — re-establish the ref baseline: clause 2 (~5 min) then clause 1 (3 cycles, ~35 min) against
  `ef32d4c`, each with a P1 `refs:` block. Pre-registered: **both green**.
- **iter-56** — the 81 sites repaired as **one derived-and-fenced class**, with two blind pre-commit diff
  seats.
- **`FENCE-M257x-iter54-refs-block`** — P1 as a guard: fail a `closed-*` iter claiming gate movement with
  no `refs:` block.
- **`CHECK-M257x-iter54-gitignored-instrument-sweep`** — P2's corollary: enumerate every instrument input
  on a git-ignored path. `.agentspace/rext.tag` is the known one; assume it is not the only one.
- **`CHECK-M257x-iter54-map-needs-prose-under-review-marking`** — P4's third category is invisible in the
  map, which is why a false §5 narrative read as authoritative as a fenced row.
- Unchanged and still open: `FIX-M257x-iter53-union-set` (46 vs 35 — user decision `D-M257x-53-5`),
  `CHECK-M257x-iter53-instrument-was-never-frozen`, `FIX-M257x-iter52-mirror-pair-leak`,
  `CHECK-M257x-iter52-second-ai-manager`, `CHECK-M257x-iter50-audited-zero-is-evidence`,
  `FENCE-M257x-iter52-stem-collision`, `CHECK-M257x-iter35-seeder-writes-one-instant`, RF-13, RF-2/3/7–12,
  harden residue iters 27–30 / 32–34 / 36–41, root `CLAUDE.md` outside the 40-file partition,
  `CHECK-M257x-iter38-ai-act-classification` (needs an owner **outside** this milestone).

**Lessons:**
1. **An input that can change without appearing in a diff is not a controlled input.** Four instruments of
   this milestone were bitten by it; three were git-ignored files and one was an unrecorded external ref.
   Generalized into `platform-alignment.md` as TOK-04's P1/P2.
2. **Derived > fenced > prose is not a preference, it is now a measurement** — three approaches, one event,
   one day: zero human action / unaided 3-for-3 catch / falsified.
3. **A milestone's own repair activity is a defect source with a measurable rate** (~50%). Today it produced
   a false claim in the very artifact built to stop false claims, within 30 minutes of committing it, and it
   escaped only because the pre-commit double-read TOK-03 mandated was skipped.
4. **Quoting a prior iter's finding forward is not evidence.** iter-01's time bomb was real when written and
   defused by iter-02; the claim survived anyway because nobody re-ran it. Re-measure, or cite the
   re-measurement.
