---
iter: 248
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
---

# iter-248 — "the family is green" is not "the suite is green", measured and then made unsayable

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

## Step 0 — re-survey

iter-245 opened `ROUTE-M257x-245-guard-family-green-is-not-suite-green` on a **two-for-two** observation,
and it is the only route this run opened that names a defect in **how this milestone works** rather than
in the corpus:

* **iter-239** added `skill_invocation_guard` to the family without indexing it — `test_fence_registry_population` went RED and stayed RED. Found by **iter-244**.
* **iter-240** edited `setup_guide.md`, moving the line `platform-alignment.md:1054` cites — `test_anchor_subject_census` went RED and stayed RED. Found by **iter-245**.

Both iters closed reporting **`guard_family`: N GREEN / 0 RED**. Both were telling the truth. **`guard_family`
runs the guards; it does not run the ~2,000-test suite behind them**, and nothing in its output says so.

**A verdict that is true and reads as a stronger claim than it makes is this milestone's most-repeated
finding.** It has been booked against reach denominators, range anchors, host-identity rollups and
could-not-resolve buckets. Here it is in the runner the milestone uses to certify its own iters.

## Hypothesis

The two inherited REDs were not carelessness; they were a **legible-output defect**. An operator reading
`29 GREEN · 0 RED` has been handed a number that answers a different question from the one they asked.
Two deliverables follow, and only one of them is a measurement:

1. **Measure** the current state — run the whole `stack-core` suite, which nothing in this run's first
   four iters did in full, and find any *further* inherited REDs.
2. **Make the conflation unsayable** — `guard_family` must state, on every run, that it is not the suite,
   and name the suite it is not.

## Pre-registered numeric claims — SEALED IN THIS COMMIT

| id | claim | prediction |
|---|---|---|
| **P-248-1** | further inherited REDs in the full `stack-core` suite, beyond the two already repaired | **0–2** |
| **P-248-2** | the suite's passing count is **≥ 1,985** (the harden-11 whole-section baseline this milestone quotes) plus this run's net-new cases | **YES** |
| **P-248-3** | `guard_family`'s output today contains **no** statement that it is not the test suite | **CONFIRMED** |
| **P-248-4** | the two repaired REDs stay repaired (regression) | **YES** |

**Falsification:** if the full suite is green with **zero** further inherited REDs, the *measurement* half
found nothing — and the disclosure half is still the deliverable, because the two REDs iters 244/245 found
are already proof that the conflation costs something. The iter says which half paid.

## Phase plan

A — run the full suite (started before this probe was written; the wall-clock is the reason).
B — triage anything RED against the pre-run tree, to separate inherited from self-inflicted.
C — the disclosure in `guard_family`. D — test it. E — re-derive last.

## Escalation

A RED that this run's own commits caused is **not** an inherited RED and must be repaired as this iter's
own defect, not reported as a finding about iters 239/240.
