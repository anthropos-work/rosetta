**Type:** tik (declared multi-step shape — fence, repair, read — per §Phase-2 carve-out)

# iter-48 — fence the leak, repair the seven, read again

Three planned steps, all three landed.

1. **`FENCE-M257x-iter48-repair-leak` built FIRST and watched RED** against the iter-46 fixture before any
   repair — §5 rule 21. On that fixture it reports 5 sites: 2 are iter-47 blockers, **1 is a real defect no
   auditor reported** (`external_services.md:565`, in a file six of seven seats read top-to-bottom), and 2
   are benign and named in a test rather than tuned away. The 2-in-5 false-positive rate is **pinned by
   `test_the_false_positive_count_is_pinned`**, so tuning it to zero fails as loudly as letting it grow.
   K was **measured, not chosen** (`D-M257x-48-2`): the longest common run a real editorial rewrite leaves
   is 8, so K=8 — a comfortable round number above that is blind to the ordinary shape of a leak.
   Twenty mutants, five inversions, one no-op control required to survive; **one further mutant survived
   the first run** (`--json` could be emptied with the suite still green) and was **closed, not booked**
   (`D-M257x-48-5`).

2. **The seven repaired by CLAIM, not by file** (§5 rule 19), tree-wide, with the commit-time
   post-condition active. Committed at `cabc3b1`; the leak fence runs **GREEN on this repair's own diff**.

3. **The eighth clause-5 reading** — seven seats, iter-41's instrument frozen on every knob, all 40 files
   top-to-bottom with per-file `wc -l` positive controls, plus the diff seat.

## The reading — 12 blockers, and the headline is the split, not the count

Full adjudication in [`blocker-ledger.md`](blocker-ledger.md). Every finding re-derived against
`app @ 5ba17044` before acceptance; **12 raw → 12 unique → 12 held** (no duplicates collapsed, unlike
iter-41's 21 → 18).

| class | n |
|---|---|
| **PRE** — predates M257x entirely (authored 2026-03-02 .. 2026-07-23) | **7** |
| **MILE** — authored by earlier iters of this milestone (23, 34) | **3** |
| **THIS** — induced by the repair this pass just made | **2** |

> ### iter-47's "the pre-existing residual measured ZERO" is REFUTED.
> **2 induced, 10 not.** iter-47 measured the same split as **7 induced / 0 not**, one repair earlier,
> **with the same instrument**. It is a complete inversion.

Seven of the twelve sit in text neither repair touched, *inside seats' assigned file sets* —
`external_services.md:662` (LiveKit agent names, authored 2026-03-02), `hiring.md:189-196` (a NOT NULL,
UNIQUE, undefaulted `token` column missing from a "minimal write-set", which **iter-47 saw and booked a
MINOR**), `dependency_map.md:19` (Storage's row contradicted by its own twin doc), `ai-readiness.md:410`
and `ai_architecture.md:202` (both authored in June, both found only now, by seat B).

**So iter-47's zero was a property of the reading, not of the corpus.** The two-term model it introduced —
corpus term plus repair term — is right about the arithmetic and wrong about one of its inputs. That is a
statement about the **instrument**, and it is the most valuable thing this iteration produced.

**Not one of the 7 pre-existing findings was reachable by any shipped fence**, traced individually rather
than assumed: `repair_leak_guard` is verbatim-only (`D-M257x-48-4` pins paraphrase out of reach,
`D-M257x-48-9` pins number-only corrections out of reach *and* shows lowering K buys two false positives
without catching them); `claim_twin_guard` only fires on claims already in a ledger; `anchor_construct_guard`
resolves anchors, and all 12 sit at **valid** ones.

## Deliverables

- `FENCE-M257x-iter48-repair-leak` + its battery, its published false-positive rate, and its **two pinned
  limits** — a fence believed to cover a class it does not cover is worse than no fence.
- The seven repaired; the claim-twin ratchet taken 12 sites → 0 by that repair.
- The eighth reading, adjudicated, with the induced/pre-existing split and the instrument that produced it.
- **`claim_twin_iter48/`** — the perishable answer-key fixture for the new blocker set (18 red / 18 green,
  pinned at rosetta `cabc3b1`), captured **before** any repair can destroy it. It is the only artifact that
  can support the claim *a full seven-auditor reading missed these while they sat in its own file sets.*

## Close — 2026-08-02

**Outcome:** eighth clause-5 reading returns **12** adjudicated blockers (10 under clause 5's literal file
scope) — **2 induced by this pass's repair, 10 not, 7 predating the milestone**. iter-47's zero
pre-existing residual is refuted; the residual is instrument-dependent.
**Type:** tik
**Status:** closed-fixed — all three planned steps landed (fence + repair + reading). The reading is a
measurement and legitimately failed to close clause 5; `overview.md` pre-declared that this is not a no-lift.
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (tik; metric moved at iter-47 so the 3-no-prog
streak cannot fire) — (3) re-scope: n (platform origin `2adcf71` unchanged at open **and** close; trigger
stays at occurrence 1 of 2) — (4) user-blocker: **y** — (5) cap-reached: n — (6) protocol-stop: n —
Outcome: **exit-4**
**Decisions:** D-M257x-48-1 .. D-M257x-48-9 (see [`decisions.md`](decisions.md))
**Side-deliverables:** none — the fixture capture is planned scope (§5 rule 21), not a side discovery.
**Routes carried forward:**
- `FENCE-M257x-iter49-numeric-leak` (Fate 3) — shingle over a **word-level diff** of removed-vs-added text
  so a `16 → 23` correction yields a *changed-claim* form. Reaches 2 of the 7 pre-existing findings; owner
  = next iter of M257x. Recorded in the fence docstring **and pinned by a test**, per `D-M257x-48-9`.
- The 12 adjudicated blockers, unrepaired by deliberate instruction — the orchestrator asked for the honest
  number and split before any repair iteration opens.
- `ops/demo/media-substrate-spec.md:33-34` (seat B, correctly routed out of its own file set) — *"never in
  prod S3"* is false about **capture**; Bunny is the serving layer. Outside clause 5's scope.
- Seat B's 14 minors (anchor drift, an ambiguous bare filename resolving to two real files).

**Escalation:** `overview.md` pre-declared that a non-zero reading *"again entirely repair-induced"* is a
user question. This reading is the **opposite** — mostly *not* repair-induced — which is the stronger
version of the same question, so it escalates for the same reason. **`EXIT_REASON: user-blocker`.**

**Lessons:**
1. **A frozen instrument is not a precise instrument.** Three readings of one corpus with identical seats,
   briefing and partition returned 18, 7, 12. The run-to-run variance is **larger than the residual being
   chased**, so a reading that returns zero is evidence about the reading. Clause 5 asks for exactly that
   reading. → folded into `corpus/ops/platform-alignment.md` §5 as **rule 22**.
2. **An audit pass that only MEASURES still changes the tree**, because its report is an input to a fence
   (`D-M257x-48-6`). iter-47 correctly repaired nothing and shipped 8 RED tests. This pass measured that
   consequence at its own close rather than leaving it for the next hand-off to discover.
3. **The author of a correction violated the rule while writing it for the sixth consecutive iteration**
   (#8). Five of eight passes have produced a blocker in text written to explain a correction. That class
   still has nothing behind it but the author.
