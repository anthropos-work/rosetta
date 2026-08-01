---
milestone: M257x
iter: 27
iteration_type: tik
status: in-progress
created: 2026-08-01
---

# iter-27 — `CHECK-M257x-iter15-manager-reads-empty`, measured PER ID

**Active strategy reference:** `TOK-01: instrument first, then follow` (milestone-root `decisions.md`) —
the only TOK on the chain; no triggered tok has ever fired in this milestone.

## Step 0 — re-survey (mandatory), and what it corrected

`demo-1` is UP, 15 containers, carrying iter-24's Directus re-point and iter-26's `pt-world` reset.
Re-read the failing set from the run's own `report/last-run.json` rather than from the hand-off:

    total 209, failing 7   (23 live / 7 failing / 1 unimplemented — matches iter-26)

**The hand-off's ENUMERATION of the four-that-share-a-signature is wrong on one row**, and the correction
is the §5 *re-derive the enumeration* rule applied to this milestone's own hand-off:

| hand-off said | the report says |
|---|---|
| `assignment-monitoring.assign-and-track.UC2` — "the seeded hero is among the per-member results — got 0" | that text belongs to **`pt-activity-drilldown`** |
| — | `pt-assignment-assign` fails on something else entirely: `expect(received).toBe(expected)` **Expected 15 / Received 14** (an affordance COUNT, not a hero-absence) |

So the four sharing *"a manager-vantage read reports the seeded hero as absent"* are
**`pt-workforce-org-feedback` · `pt-workforce-succession` · `pt-workforce-funnel` · `pt-activity-drilldown`**,
and `pt-assignment-assign` is a **singleton**, not a member of the class.

## Cluster / target identified

The four hero-absence failures. Measured **per id** — the hand-off's own instruction, and iter-19's lesson
(a shared symptom is not a shared cause).

## Hypothesis

The hand-off's cheap hypothesis was *the dropped `local_*` session mirrors*. **It is already refuted for at
least two of the four before a line is written** — see the two measurements below — so the working
hypothesis is instead: **these four are NOT one bug**, and the split is *seed-side absence* vs *read-side
invisibility*.

Measured at re-survey (both reproducible from the live demo-1 DB):

1. **`pt-workforce-org-feedback` — the row is genuinely NOT SEEDED.** The hero is population slot 1
   (`deterministicUUID("demo-1:story:pt-org-a:user:1")` == `23f24e3f-38fb-5027-9e07-2ef49a644af5`, exact),
   and `feedback.go:116`'s share draw for slot 1 is **0.8305 against a 0.45 share** → excluded,
   deterministically and permanently (the prefix is pinned by `stack: demo-1` **in the seed YAML**, so it
   cannot vary by stack). Reproduced independently in Python: the predicted in-share slots
   `[7,12,24,27,28,34,36,38,40]` are **exactly** the 9 rows the org actually has.
2. **`pt-workforce-succession` — the row IS seeded.** `public.interview_extraction_results` holds the
   hero's row (`957d5253-…`, FK'd to her real session `f22d182d-…`), because `succession.go:114` reads
   `if !isHero && !memberInShare(...)` — an explicit hero exemption. So its failure is **read-side**.

## Expected lift

**Not predicted as a clause-2 number, deliberately.** At most one of the four is fixable seed-side in this
iter, and iter-24's lesson (a fix addressing one of ≥4 causes) applies. The iter's claimable deliverable is
(a) a landed, family-consistent seeder fix proven at the DATA layer, and (b) **four separately-attributed
mechanisms** replacing one unattributed cluster.

## Phase plan

1. Finish the per-id diagnosis for `pt-workforce-funnel` and `pt-activity-drilldown` (the two not yet split).
2. Land the seed-side fix for the one that is genuinely un-seeded, in the shape the sibling seeders already
   use, with a fence + a mutation battery carrying a declared-GREEN no-op control (§8 rule 5).
3. Prove it live at the data layer on `demo-1` (re-seed → the hero's rows exist), and run the ONE affected
   spec scoped for **diagnosis only** (§ the ptreport gate binds only on a full run — a scoped run is never
   quoted as a clause-2 number).
4. Route the read-side ones with named handlers and their measured mechanism.

## Escalation conditions

- If the seed-side fix needs a platform edit → **stop**, route, do not edit (v2.8's binding constraint).
- If a second platform commit lands at origin → re-scope trigger occurrence 2 → exit.
- If the diagnosis opens a **third unplanned** line of investigation → scope-creep tripwire; land what is
  complete and route the rest.

## Acceptable close-no-lift outcomes

Four separately-attributed mechanisms with evidence, even if no clause-2 number moves, is a complete iter —
the cluster was the deliverable the hand-off named, and splitting it is what makes the next fix targetable.

## Budget note (iter-25's mistake, not to be repeated)

A **full** Playthrough run is the clause-2 instrument and costs an entire iteration (serial, `workers:1`).
It is **out of scope here** and routed as its own measurement.
