---
iter: 48
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-02
---

# iter-48 — fence the leak, repair the seven, read again

**Active strategy reference:**
[`TOK-02`](../decisions.md#tok-02-fence-the-prose-the-way-the-anchors-are-fenced--2026-08-02).
TOK-02's five steps are complete and iter-47 graded them at **18 → 7**. This iteration is **TOK-02's own
pattern applied once more** — *fence, then repair, then read* — not a new strategy. It is a tik, not a tok:
the metric moved at iter-47, so the 3-no-prog trigger cannot fire.

## Step 0 — re-survey before targeting

Platform origin HEAD re-fetched at open: **`2adcf71`, unchanged** (re-scope trigger stays at occurrence 1
of 2). iter-47's seven blockers re-confirmed live before touching anything: `repair_postcondition.py`
against the working tree reported **12 sites over exactly those 7 locations**, so the target is current,
not stale.

**Hand-off correction (the 18th consecutive).** The hand-off records the `stack-core` baseline as
**14F/491**. Measured at rext HEAD `3ff8118` against rosetta `72298dd` on a pristine `git archive HEAD`
tree, it is **22F/491**: iter-47 committed its blocker-ledger, that ledger registered 12 new claim-twin
sites, and **the ratchet went RED in 8 tests** — which iter-47's own close did not measure. Not a defect
introduced here, and repairing the seven is what clears it.

## The three planned steps (this is a declared multi-step shape, per §Phase-2 carve-out)

1. **Build `FENCE-M257x-iter48-repair-leak` FIRST, and watch it go RED against the current corpus** —
   before repairing, because repairing destroys the fixture (§5 rule 21). iter-47 named it and it needs no
   adjudication: *for every claim a commit changes, grep the whole tree for the old form.*
2. **Repair the seven by CLAIM, not by file** (§5 rule 19), tree-wide, with the commit-time
   post-condition active.
3. **ONE full 7-auditor read at iter-41's frozen instrument** — same seat count, same briefing, same
   partition, all 40 files top-to-bottom, plus the diff seat. That reading alone grades clause 5.

## Pre-registered predictions — written before any auditor reported

Predictions have been the most informative artifact in this milestone: four refuted, one confirmed. These
are recorded here, in the commit that precedes the read, so they cannot be tuned afterwards.

| # | prediction | grading |
|---|---|---|
| **headline** | **3 blockers.** Fewer than iter-47's 7, and **not zero** | |
| 1 | **Zero of the verbatim self-contradiction class** ("repaired at one site, left standing at another"). The leak fence is GREEN on this repair's own diff, so a survival of this class can only be a **paraphrase**, which measurement showed is out of that fence's reach | |
| 2 | **At least one blocker in text written to EXPLAIN a correction.** This has held in five consecutive iterations, including for the author of the rule forbidding it. I have shortened that surface deliberately — linking to canonical anchors instead of restating them — but I have not mechanized it | |
| 3 | **If there is a residual, `ai_architecture.md:42-68` carries it.** That block is the most heavily rewritten in this repair, it is the one auditor G said to rewrite against `external_services.md:541` rather than de novo, and it has produced a blocker in three consecutive iterations | |
| consequent | **clause 5 does NOT close** | |

**Why not zero, stated plainly.** Two of the three defect classes iter-47 measured now have a machine
behind them — the verbatim leak (this iteration's fence) and the returning adjudicated claim
(`claim_twin_guard`, whose ratchet this repair took from 12 sites to 0). The third — *"is this newly
written sentence true?"* — has nothing behind it but the author, and it produced **4 of iter-47's 7**.
Predicting zero would be predicting that the one unmechanized class stopped occurring for the first time
in eight passes, on the pass where the corpus was hardest to write about.

## Escalation conditions

- Reading returns **zero** → clause 5 met, gate 5 of 5, `EXIT_REASON: gate-met`.
- Reading returns **non-zero and again entirely repair-induced** → `EXIT_REASON: user-blocker`. A corpus
  that is clean except for the act of cleaning it is a question for the user, not another cycle.
- A second platform commit invalidating an alignment attempt → `EXIT_REASON: re-scope-trigger`.

## Acceptable close-no-lift outcomes

None — this iteration ships a fence and a repair, so it lands deliverables regardless of the reading. The
reading is a measurement and may legitimately fail to close clause 5; that is not a no-lift.
