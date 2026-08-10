---
iter: 278
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
---

# iter-278 — pay the advance's corpus debt in the iter that reads it

**Type:** tik · **Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them)
— *census the mechanical classes; stop sampling them.*

## Step 0 — re-survey (mandatory, before targeting)

`clone_drift_guard` re-run at this tree: **RED**, one finding.

    [D1 advanced] rosetta-extensions is at 0a8674e74, which the corpus never cites —
    22 commit(s) past the nearest of 12 cited sha(s); 44 citing site(s)

Guard family at open: **23 GREEN · 1 RED · 0 could-not-check · 11 not-run.** The single RED is the
target. iter-277's named highest-value route is confirmed live and is **not** absorbed.

## Cluster / target identified

`FIX-M257x-277-corpus-cites-a-rext-sha-that-no-longer-exists`. Created by iter-276: shipping the
Playthrough fix advanced `rosetta-extensions` past every sha the corpus cites. iter-277's Lesson 3 —
*"the tooling half and the documentation half of this milestone are coupled; that debt should be paid
in the same iter that creates it"* — is the thing this iter tests on itself.

## Hypothesis

**The 44 sites are not the repair.** D1's own docstring says it "does NOT adjudicate truth-at-a-ref",
and `advance_impact_census` places a block-pinned citation **out of subject** by contract (§5 rules
41/44). If that holds, renumbering the 44 would be iter-256's error repeated — *27 "obvious"
renumbering repairs reverted byte-for-byte, seven of which had moved **correct** citations onto
comments.*

What D1 **does** assert is that a whole repo's worth of change has never been looked at. So the repair
is a **review of the 22 commits for corpus impact**, and the sha lands where that review actually
changed a claim.

## Expected lift

`clone_drift_guard` GREEN, earned by ≥1 substantive repair — never by prose *about* the drift, which
is this fence's own documented weakness (`FIX-M257x-iter107-drift-fence-satisfiable-by-prose`: *"writing
about the drift satisfies the drift fence"*).

## Pre-registrations — sealed in this iter's FIRST commit, before any repair

- **PR-1** — every one of the 44 sha-citing sites is a **ref-scoped** claim, so **0** of the 44 is
  falsified by the advance and **0** is renumbered by this iter.
- **PR-2** — **0** corpus-cited rext *path* is dead at HEAD (`rext_path_guard`'s population, not a
  second parser's).
- **PR-3** — the corpus sites citing files the advance **touched** contain **≥ 1** claim the advance
  made false. If **0**, the honest close is `closed-no-lift`: the drift would be real and its corpus
  cost nil, and the fence would be measuring mentions rather than currency.
- **PR-4** — after the repair `clone_drift_guard` is GREEN **and** every newly-written HEAD sha sits on
  a sentence about the **new state**, not on a sentence about the drift. Graded by reading, and
  recorded either way.
- **PR-5** — `route_disposition_guard` stays GREEN. The closed-route correction lands in `corpus/**`,
  which is not that guard's subject; if it reddens, the correction was written in the registry's
  grammar by accident.

## Phase plan

A. census the affected population · B. grade each site (ref-scoped vs gradeable-at-HEAD) · C. repair
what the advance falsified · D. re-measure the family · E. close.

## Escalation conditions

A repair that cannot be grounded without renumbering an anchor whose clock is undecidable → **route it,
do not guess** (`D-M257x-122-5`, and iter-256's revert).

## Acceptable close-no-lift outcomes

PR-3 returning **0** — the advance cost the corpus nothing measurable. That is a real finding about the
fence's reach and would be reported as one.
