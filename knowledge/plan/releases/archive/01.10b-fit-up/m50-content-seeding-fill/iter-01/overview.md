---
iter: 1
milestone: M50
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
created: 2026-06-30
---

# iter-01 — bootstrap tok (author initial strategy)

**Type:** tok (bootstrap)

## Job
Author M50's first strategy from the milestone `overview.md`, the annotation field-review, the four KB-contract
docs, the existing rext seeder fleet, and a re-diagnosis on the **FRESH demo-1** (the orchestration's critical
constraint — many annotation gaps were observed on the OLD STALE pre-M47/M48 demo and may already render now).
Does NOT terminate the call (Phase 5 §2); the loop continues into iter-02 (the first tik) under TOK-01.

## Inputs consumed
- `overview.md` (exit gate = M42 coverage gate both vantages on a COLD demo, 0 prod-eject escapes).
- `.agentspace/annotation.md` (the field review: Maya academy/activities/xp/skill-paths/home; Dan
  workforce-tabs/assignments/members/studio-link).
- `coverage-protocol.md` (the iteration protocol + fix-surface routing table + manifest model).
- `stories-spec.md` (7-table chain + M36 dashboard seeders), `profile-completeness-spec.md` (M41/M44
  rubric), `seeding-spec.md` (isolation/closure).
- The rext `stack-seeding/seeders/` fleet (20 seeders; **no `member_languages.go`**).
- A live re-diagnosis of demo-1's Postgres (`:15432`) — see spec-notes "Re-diagnosis".

## Phase 0b — KB-fidelity gate
Ran `/developer-kit:audit-kb-fidelity --milestone=M50` → **YELLOW** (proceed; gaps = known context).
Recorded in `spec-notes.md` (F1 gate-MET reconciliation + F2 HeroActivitySeeder-exists are the two must-knows).

## Initial strategy → TOK-01
See milestone-root `decisions.md` TOK-01. In short: the metric is the M42 coverage sweep `(failingSections,
escapes)` per vantage; the gate is `(0,0)` both vantages on a COLD demo. Iterate by **re-seeding against the
running demo-1** (light) per tik — observe (sweep) → triage by the routing table → fix the highest-leverage
seed cluster in rext `stack-seeding` → re-seed → re-sweep. Reserve the heavy COLD reset-to-seed for the
exit-gate proof only (machine-aware: 9 GiB Docker VM + dev stack co-resident).

## Re-survey baseline framing
The protocol's primary metric is a sweep result; iter-02 (first tik) runs the **baseline sweeps** (employee +
manager) to fix the true starting `(failingSections, escapes)` before targeting a cluster. The re-diagnosis
already gives the seed-level gap inventory (spec-notes), so iter-02 can target the highest-leverage cluster
immediately after the baseline confirms which gaps actually surface in the rendered sweep.

## Outcome
TOK-01 authored. Strategy class new-direction. Continues into iter-02 (tik).
