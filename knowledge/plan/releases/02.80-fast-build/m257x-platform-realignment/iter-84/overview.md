---
iter: 84
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-05
---

# iter-84 — adjudicate the 43, and re-derive the eleven discharges

**Type:** tik, under `TOK-05`. Discharges the routes iter-83 opened and the one iter-82 left binding.

## Step 0 — re-survey

Metric re-derived at open: gate **4 of 5**; clause 5 the only open clause, graded only by a reading that
returns zero. Both repos clean and **pushed** (`rosetta 36cca77`, `rext 24819f0`, remote shas verified).
All 6 corpus guards exit 0. Ground-truth clones re-derived at this open — `platform 0dab54df`,
`app b948604f`, `cms ca50c817`, `jobsimulation 462343b0`, `roadrunner 87d8d443`, `storage 4ce8ece5`,
`messenger fa47850d`, `sentinel 88bc5592`, `next-web-app bb3313bc`, `studio-desk 14a5442a`,
`ant-academy 9c3843cd`, `graphql-wundergraph 60c229f3`, `.agentspace/rosetta-extensions ab81527a`.

> **⚠️ `.agentspace/rosetta-extensions` reads `ab81527a` in `stack-demo/`, and the authoring copy is at
> `24819f08`.** These are two different clones of the same repo. A corpus citation into
> `rosetta-extensions` grades against the **authoring copy** (`24819f08`) unless the citing block pins a
> ref — the largest class in the 43 is rext anchors, so this distinction is load-bearing, not pedantic.

**Target confirmed still meaningful:** the 43 are unadjudicated and no repair has touched them.

## Cluster / target identified

Three routes, one dependency chain — **adjudication is the precondition for repair** (iter-80, binding),
and **the eleven discharges are unproven** (iter-83, measured).

## Hypothesis

The 43 will *not* collapse. iter-80 measured the seventh routed count at **92.1 % upheld**, the first
that did not collapse, and iter-82's union is drawn by the same frozen instrument from a corpus that has
had exactly one repair pass since. **Pre-registered:** ≥ 70 % upheld. A collapse below 50 % would mean
the instrument regressed after iter-81's repair and would itself be the finding.

Second, independent: **re-deriving the eleven predicates by membership will find survivors in more
predicates than P4.** iter-83 measured 38 unreached anchors across 16 files and both count styles;
if only P4 has survivors, iter-83's generality claim is wrong and must be retracted.

## Expected lift

None on the gate metric — clause 5 moves only on a reading, and no reading is taken here. The
deliverables are the per-anchor ledger, the eleven membership verdicts, and a work list for iter-85.

## Phase plan — three planned lines (declared)

1. **Adjudicate the 43** — 4 parallel adjudicators over disjoint packets, each re-deriving from the
   clones, never from a seat's note (iter-76 method note 2) and never from another document
   (TOK-05: *two documents that agree are not two witnesses*).
2. **Emit a PER-ANCHOR ledger** — `FIX-M257x-iter83-adjudication-has-no-per-anchor-ledger`. iter-76
   recorded rejection *mechanisms* and counts but never per-anchor verdicts, which is why reach can only
   be graded against *booked* rather than *upheld*. Not repeating that one iteration later.
3. **Re-derive the eleven discharges as MEMBERSHIP questions** — enumerate each predicate's legal set
   over the corpus and check every member. Does **not** require the frozen instrument.

## Escalation conditions

- **Seat-ref discipline is on its 3rd occurrence.** Every adjudicator is told explicitly to grade a
  ref-pinned claim **at the ref the claim names**, and that the clones sit at *older* shas than some
  cited refs. A 4th occurrence is escalated, not absorbed.
- If an adjudicator cannot settle an anchor from source, it returns **UNSETTLED** with the reason. An
  unsettled anchor is never silently upheld or silently dropped.
- `DEF-M257x-iter80-storage-prod-bucket` stays **held**; `storage.md:55,:154,:181` unchanged.

## Acceptable close-no-lift outcomes

If the 43 collapse below 50 % upheld, the iter closes with that falsification recorded — it would mean
the instrument's post-repair signal is mostly noise, which is a finding about the gate's own instrument
and more important than a work list.

## Out of scope, routed

- **Repairing by predicate** → iter-85, gated on this adjudication.
- `CHECK-M257x-iter83-standalone-is-the-forgettable-class` → iter-85.
- Re-running the frozen instrument. Not done in this run.
