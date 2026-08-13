---
iteration_type: tik
status: closed-fixed
opened: 2026-08-02
closed: 2026-08-02
---

# iter-34 — the clause-5 confirming pass

**Active strategy reference:** `TOK-01: instrument first, then follow` (milestone-root `decisions.md`).
Clause 5 is a *corpus* clause; TOK-01's step 4 ("then the corpus — the migration-status map with its two
states per row, and the reconciliation sweep") is the phase this iter sits in. The instrument (the full
top-to-bottom read with a per-file positive control) was built at iter-33; this iter *uses* it.

## Step 0 — re-survey before targeting

Run at open, 2026-08-02 00:33–00:34 CEST:

| check | result |
|---|---|
| rosetta `m257x/platform-realignment` @ `5e37bb1`, tree clean, 0 behind `main` | ✅ |
| platform origin HEAD re-fetched | **`2adcf71` — UNCHANGED.** Re-scope trigger stays at occurrence 1 of 2 |
| rext pin `.agentspace/rext.tag` = `fast-build-m257x-iter-31b`, matches origin + clone | ✅ |
| `stackseed` binary mtime `Aug 1 23:49` post-dates the pinned tag's commit | ✅ |
| clause-5 target still meaningful? | **YES.** iter-33 measured 6 blockers on its last reading and fixed them; no confirming pass has run. The clause is one measurement away and nothing has absorbed it. |

The TOK-directed target is **not** stale. No substitution.

## Cluster / target identified

`MEASURE-M257x-iter34-clause5-confirming-pass` — the single outstanding action between clause 5 and MET,
routed forward by iter-33's close. Clause 5 reads:

> KB-fidelity audit **GREEN, or YELLOW with 0 blockers**, over `corpus/services/**` + `corpus/architecture/**`.

iter-33's last reading returned **6 blockers**, not 0. A clause is not met by an absent measurement
(the `25 → 27` mistake iter-32 diagnosed). This iter runs the measurement.

## Hypothesis

The corpus is now materially closer to platform origin HEAD than it was at iter-33 open (25 blockers
closed across 13 files), and the residual is small but **not zero** — because (a) the instrument is a
full read, not vocabulary-bound, so it does not converge by exhausting itself, and (b) iter-33's own
pass-2 measured a **24% self-inflicted rate** on the repair text, which is a property of corrective
sweeps in general, not of that sweep in particular.

## PRE-REGISTERED PREDICTIONS (written 2026-08-02 00:36, before any report exists)

Two predictions, so one can be wrong — the discipline that corrected the milestone's model of its own
corpus at iter-33.

**P1 — count.** The confirming pass returns **1–5 blockers**, most likely 2–3.
Reasoning: pass 1 found 19 over an unswept corpus; pass 2 found 6 over the 13 *swept* files. Both are
fixed. What remains is (i) inter-rater residual — different readers grade differently and pass 1's
readers were not exhaustive, and (ii) any new prose defect the pass-2 corrections introduced. Neither
mechanism plausibly yields 0, and neither plausibly yields ≥10 given 25 real closures.
*Falsified by:* 0, or ≥6.

**P2 — location.** The largest residual cluster lands in the **27 files the sweep did NOT touch**, and
specifically in the 18 "remaining service docs" that pass-1 group 5 reported as **1 blocker in 18 files**.
Reasoning: that density (0.06 blockers/file) is an order of magnitude below every other group
(1.0, 0.75, 0.67, 0.86). An 18-file group is also the largest reading load per agent. **Under-detection
is a more parsimonious explanation than genuine cleanliness.**
*Falsified by:* the majority of blockers landing in the 13 swept files, or the residual distributing
evenly.

**If P1 returns 0 (or YELLOW with 0 blockers), clause 5 is MET** and the gate goes to 4 of 5 — and P1
is falsified in the direction that closes a clause. That outcome is recorded as a *refutation*, not
retro-fitted as a success.

## Expected lift

Clause 5: NOT MET → **MET**, conditional on a zero-blocker reading. If blockers > 0, the lift is the
*measurement plus the closures*, and clause 5 stays NOT MET with a fourth pass routed forward — iter-33's
refusal to grade on an absent measurement is precedent and it binds this iter symmetrically.

## Phase plan

1. **Phase A — re-derive ground truth.** Confirm platform origin unchanged at `2adcf71`; reuse
   `iter-33/iter33-groundtruth.md` as the shared brief (its facts were derived against that exact sha),
   plus a confirming-pass addendum.
2. **Phase B — the read.** 5 read-only sub-agents, **re-partitioned** so no agent inherits pass 1's group
   boundaries (correlated blind spots do not survive a different cut), each ~1 700 lines, each mixing
   swept and unswept files, with a mandatory `wc -l` positive control per file.
3. **Phase C — grade.** Verify every reported blocker against platform source *before* acting on it
   (the milestone's standing re-derive rule). Fix by evidence rank.
4. **Phase D — close.** Grade clause 5 honestly on the reading actually taken.

## Escalation conditions

- A platform commit lands (re-fetch at close) → re-scope trigger fires, exit.
- A reported blocker turns out to require a platform-repo edit → escalate; 0 platform edits is binding.
- Blockers > 0 → close with the count, route the next pass forward. **Not** a user-blocker.

## Acceptable close-no-lift outcomes

A reading that returns blockers and is fixed but leaves clause 5 unconfirmed is still a complete iter:
the deliverable is *the measurement*, and a measurement that refutes P1 upward is the most informative
outcome available.

## Group partition (re-cut, deliberately unlike pass 1)

| group | files | lines |
|---|---|---|
| A | external_services\*, studio-desk, ant-academy, dependency_map\*, architecture/README, skiller | 1 716 |
| B | ai-readiness\*, alignment_testing, service_taxonomy\*, services/README, TEMPLATE, db-backup, intelligence | 1 698 |
| C | studio-room, clerkenstein\*, architecture_overview\*, chronos, platform-migration-status, gotenberg, customerio-sync | 1 688 |
| D | hiring\*, backend\*, cms\*, jobsimulation\*, graphql-wundergraph, roadrunner, sentinel, ai-labs | 1 710 |
| E | shared_libraries\*, security_compliance\*, storage\*, ai_architecture, academy-backend, coursebuilder, messenger, clerk-integration, next-web-app, askengine, frontend_architecture, skillpath | 1 718 |

`*` = touched by iter-33's sweep. **8 530 lines, 40 files, 100% coverage.**
