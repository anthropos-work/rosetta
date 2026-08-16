---
iteration: 01
iteration_type: tik
status: closed-fixed
milestone: M265
date: 2026-08-16
---

# M265 iter-01 — prove it live, and find what six green gates could not see

## Type selection (Phase 0)

**tik.** Not a bootstrap tok, and the reason is recorded rather than assumed.

`/developer-kit:build-mstone-iters` Phase 0 rule 1 makes iter-01 a bootstrap tok **when no prior iter
dirs exist**. Here `iter-01/` already existed (scaffolded by `/developer-kit:design-roadmap` Phase 8)
while the milestone-root `decisions.md` carried **0 `TOK-*` entries** — Before-You-Start case **(b)**,
the bootstrap-less shape. The rule's precondition is false, so it does not fire.

Classifying this iter as a tok anyway would have been false in substance: **a tok does not move the
gate, and this iter moved it from unmeasured to MET.** It is a tik, run without a TOK chain, planning
from the milestone's `overview.md` + the protocol doc directly.

## Active strategy reference

**None — no TOK chain (case (b) legacy-shaped milestone).** Planning came from `overview.md`'s five
gate clauses and `corpus/ops/verification.md`, per Before-You-Start step 5 case (b).

## Pre-flight gates (evaluated, not skipped silently)

| Gate | Verdict | Why |
|---|---|---|
| **Phase 0b** — KB-fidelity | **SKIP** | The rule runs it on iter-01's *bootstrap tok*; this is a tik in case (b), which the rule explicitly excludes. Independently moot here: a gate that blocks *before implementation* has nothing left to block — the implementation is complete and measured. |
| **Phase 0d** — tooling pre-flight | **SKIP** | Triggers on wiring new artifacts through a multi-stage generate/validate pipeline. This iter is code-fix + live-measurement work; the bring-up and the Playthrough suite ARE the gates, not a separate pipeline. |

## Cluster / target identified

The whole gate. v2.9 had shipped **five** milestones of taxonomy realignment (M259–M264), every one of
them closed green — and **nothing in the release had measured a rendered surface**. This iter's target
was the gap that made possible: prove the canon live on a cold demo AND a cold dev stack, or find what
the green gates could not see.

## Hypothesis

A cold bring-up on the new canon would surface a failure set that unit tests, row counts and liveness
probes structurally cannot: **hollow success** — a stack that is up, healthy, correctly counted, and
useless.

## Expected lift

All five gate clauses measured MET on a cold bring-up.

## Phase plan

Per `corpus/ops/verification.md`: cold `/demo-up` → auto-verify → the Playthrough batch gate →
`/dev-up` → seed-closure DNA → live navigation check. Measure every clause; fix what the measurement
surfaces; re-measure.

## Escalation conditions

- A defect requiring a **platform-repo edit** → escalate (this corpus takes 0 platform edits).
- A clause unmeasurable on this host → record as NOT MET (the v2.8 M258 lesson: a projection is not a
  measurement).

## Acceptable close-no-lift outcomes

A clause proven unmeasurable with the reason recorded would have been a complete iter even with the
gate unmet. That did not occur — all five were measured.
