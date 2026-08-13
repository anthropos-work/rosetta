---
milestone: M257x
iter: 29
---

**Type:** tik (under `TOK-01: instrument first, then follow`)

# iter-29 — the flake hypothesis is REFUTED: clause 2's instrument is deterministic

## What was done

Three full `--reset` Playthrough runs against an **unchanged** build (no source change of any kind between
them; rext pinned at `fast-build-m257x-iter-27`, platform origin `2adcf71`). Run A is iter-28's; B and C
were run back to back for this iter. Each run's own JSON artifact was preserved separately as
`runA/B/C.json` — which is also the standing defence against
`FIX-M257x-iter27-scoped-run-clobbers-binding-report`, since otherwise only the last would survive.

    runA: 25/31 passing (80.6%)   5 failing
    runB: 25/31 passing (80.6%)   5 failing
    runC: 25/31 passing (80.6%)   5 failing

**Three matching headline counts prove nothing** — `5` three times could be fifteen different failures
(iter-19's rule, and this iter exists precisely because two ids had already moved without explanation). The
measurement is the id-level diff:

    $ diff failA.txt failB.txt   ->  IDENTICAL
    $ diff failB.txt failC.txt   ->  IDENTICAL

    pt-activity-drilldown
    pt-onboarding-hiring-candidate
    pt-orgadmin-role-create
    pt-workforce-funnel
    pt-workforce-succession

Union = intersection; the symmetric difference is **empty**. Zero bistable ids across three runs.

## What this refutes, and what it leaves standing

**Refuted: `CHECK-M257x-iter28-clause2-flake-component`.** iter-28 proposed that clause 2 carried an
unquantified flake component and that a `30/0/0` conjunction might therefore be unsatisfiable *reliably*.
Measured across three runs, that is false at the id level. **This is the first evidence in the milestone
that clause 2's instrument returns the same answer twice** — nobody had ever run the suite twice against
the same build, because every prior full run was separated by a landed fix, so a moving id was always
explicable and the question never arose.

That is worth more than the flake rate would have been: a `30/0/0` gate over a *non-deterministic* suite
would be a target that could be hit by luck, and iter-14's withdrawn three-green-cycles is this milestone's
own precedent for exactly that failure. Clause 2 is a meaningful conjunction.

**Left standing, and now HARDER, not softer:** the two un-attributed flips are **not** flakes.

| id | flipped at | status after this iter |
|---|---|---|
| `hiring.recruiter-comparison.UC1` | iter-26 | a one-way transition needing a real account; still open |
| `pt-assignment-assign` | iter-28 | flipped and **stayed** flipped across three runs |

iter-28 offered a hydrating-grid race as the plausible mechanism for `pt-assignment-assign` and declined to
claim it. **That candidate is now substantially weakened by measurement**: a hydration race would be
bistable, and this one is not. So something made a persistent change to that surface's state. The space of
explanations has narrowed from "timing" to "state", which is the useful half.

Naming the honest residual: iter-27 wrote hero rows into `public.job_simulation_feedbacks` **and** the
app-side `public.job_simulation_sessions` rows that carry them, so the hero's session footprint changed in
the same release the affordance count changed. That is a **coincidence of timing, not an attribution** — no
query has been run against the affordance surface, and this milestone has already had one such inference
refuted an iter after it was made (iter-19 on iter-15's caveat).

## Consequence for the remaining work

Only the **intersection** is worth targeting with a fix, and it is now exactly five stable ids. Every one of
them is reproducible on demand, which is the precondition for the read-side digs iter-27 routed.

## Close — 2026-08-01

**Outcome:** clause 2's instrument is deterministic across three full reset runs on an unchanged build —
identical sorted-id failing sets, zero bistable ids. The flake hypothesis is refuted and the two
un-attributed flips are reclassified from "probably timing" to "persistent state change, unexplained".
**Type:** tik
**Status:** closed-fixed (the planned deliverable was the characterization, and it landed with the id-level
evidence behind it; the null result was declared acceptable in the overview *before* the runs)
**Gate:** NOT MET (3 of 5 — clause 2 stable at `25 / 5 / 1`, needs `30 / 0 / 0`)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (platform origin `2adcf71`
re-checked at open and close; occurrence stays 1 of 2) — (4) user-blocker: n — (5) cap-reached: n —
(6) protocol-stop: n — Outcome: continue
**Decisions:** none beyond the close (a measurement tik).
**Side-deliverables:** per-run artifacts preserved separately, which is the working practice that makes
`FIX-M257x-iter27-scoped-run-clobbers-binding-report` survivable until it is fixed.

**Routes carried forward:**

| item | why | target |
|---|---|---|
| ~~`CHECK-M257x-iter28-clause2-flake-component`~~ **CLOSED — REFUTED by measurement** | Zero bistable ids across three full reset runs on an unchanged build. | — |
| `CHECK-M257x-iter28-assignment-flip-is-stateful` (supersedes `CHECK-M257x-iter27-assignment-affordance-count`) | `pt-assignment-assign` flipped and **stayed** flipped, so the hydration-race candidate is weakened and the cause is a persistent state change. Note the un-attributed coincidence: iter-27 changed the hero's session footprint in the same release. **Measure the affordance query; do not infer from the coincidence.** | later tik |
| `FIX-M257x-iter27-succession-hero-not-rendered` | The best-evidenced remaining failure and now **reproducible on demand** (3/3 runs). Her interview row EXISTS, FK'd to her real session; the row the spec wants is a computed projection, so the question is what the app derives. | next tik |
| `FIX-M257x-iter27-funnel-card-role-missing` | Also 3/3 reproducible. Her card renders; only the role text inside it is missing while the DB carries it on three axes. DOM/locator-shaped. | next tik |
| `FIX-M257x-iter27-scoped-run-clobbers-binding-report` | Unchanged; nothing in the report file distinguishes a binding full run from an advisory scoped one. | next tik |
| `CHECK-M257x-iter27-drilldown-target-coupling` | 3/3 reproducible. | later tik |

**Lessons:**

- **Run the instrument twice before trusting what it says about the system.** Fourteen iters treated clause
  2's number as a fact about the build. It is — but that was an assumption until this iter, and the cost of
  checking was two unattended runs. TOK-01 said *instrument first*; this is the part of that strategy the
  milestone had skipped for its most-quoted metric.
- **Declare the null result acceptable BEFORE the measurement.** iter-29's `overview.md` recorded that
  determinism would be "equally valuable and not a failure of the iter". Writing that first is what stops a
  refuted hypothesis from being quietly re-framed as a disappointment, or worse, re-run until it fires.
- **A refutation can make a residual harder rather than softer.** Killing the flake explanation did not
  clear the two flips — it removed their easiest excuse.
