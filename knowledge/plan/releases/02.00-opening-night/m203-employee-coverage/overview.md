---
milestone: M203
slug: employee-coverage
version: v2.0 "opening night"
milestone_shape: iterative
exit_gate: "Every declared EMPLOYEE-vantage use case has a PASSING Playthrough on a COLD reset-to-seed demo stack (the 3 employee stories — Maya's core journeys: Skill Paths browse→enroll→complete→verify-skill; AI Simulations chat/code launch→complete→score-in-range NON-voice; Profile verified-skill chart + claimed-vs-verified gap + work/education timeline), with 0 false-fails over 5 consecutive reset runs."
iteration_protocol_ref: knowledge/plan/spec-drafts/playthroughs/spec.md (the consolidated capability spec; the M202-delivered runbook corpus/ops/demo/playthroughs.md supersedes it as the protocol once M202 graduates it)
status: archived
created: 2026-06-28
last_updated: 2026-07-02
complexity: large
delivers: passing employee-vantage Playthroughs (in the `playthroughs` rext section) for Maya's core journeys — Skill Paths, AI Simulations (NON-voice), Profile — added to the manifest + the per-surface page-object layer, on the M202 foundation; plus the corpus/runbook updates as use-cases land
depends_on: M202
spec_ref: knowledge/plan/spec-drafts/playthroughs/spec.md (the consolidated capability spec, v0.3)
---

# M203 — Employee-vantage coverage

## Goal
**Maya's** core **employee** journeys play green as Playthroughs on the M202 foundation:
- **Skill Paths** — browse → enroll → complete → verify-skill.
- **AI Simulations** — chat/code launch → complete → score-in-range (**NON-voice**: voice/recording is the M206
  mirror tier, future; chat/code/document sims are playable as-is, asserted at the launch/completion boundary
  per spec §5.8).
- **Profile** — the verified-skill chart + the claimed-vs-verified gap + the work/education timeline.

A Playthrough **is the user** (P1): drive the real antd UI as Maya would, assert **user-observable outcomes**
(the goal achieved), never pixel/copy specifics (P2).

## Exit gate (objective, machine-verifiable)
**Every declared employee-vantage use case has a PASSING Playthrough on a COLD reset-to-seed demo stack** (the 3
employee stories — the journeys above), **with 0 false-fails over 5 consecutive reset runs.** A run starts from
the seeded baseline; after a mutating suite the world is reset via the real `--reset` path (the M202
reset-to-seed lifecycle — additive re-seed is NOT a reset). The 5-run zero-false-fail bar proves determinism
under mutation.

## Why iterative (not section)
The use-cases are **declarable** (the journeys above), but getting them green against the **real antd UI** (the
landmark layer) + the **AI-sim assertion boundary** is **exploratory** — the failure modes (a thin-a11y surface,
an async projection's settle timing, a flow that ejects out of the demo, a seed gap) are **discovered BY** the
attempt, not enumerable up front. Each iter declares/plays a use case → it fails against the real UI → diagnose
(landmark anchor / locator / seed / boundary) → fix in `rosetta-extensions` (the page-object registry / the
dedicated seed / the manifest) → re-play, until the gate is GREEN for the employee vantage. Like M42e. Build with
`/developer-kit:build-mstone-iters`.

## Iteration protocol
The loop is the playthroughs spec / the **M202-delivered runbook**:
[`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../spec-drafts/playthroughs/spec.md) **until M202
graduates it to** `corpus/ops/demo/playthroughs.md`, which then becomes the authoritative protocol.

## Re-scope trigger
A surface that **can't** be driven without a platform edit (the `unimplementable-without-platform-edit` state) →
**escalate, don't edit** the platform (the P3 zero-edit line; the platform repos are read-only this release).
Mirrors the M42e/coverage-sweep re-scope trigger.

## Depends on / Parallel with
- **Depends on:** **M202** — reuses its foundation (the manifest + validator, the per-surface page-object layer,
  the dedicated decoupled seed, the reset-to-seed serial-default runner, the 4-state reporting). The use cases it
  proves are declared in the **M201 manifest corpus**.
- **Parallel with:** **M204** (manager-vantage). The two share the **landmark-registry + locator index** (the
  M202 page-object layer); each vantage adds its own surfaces/anchors — an **additive merge surface**, not a
  conflicting one. Both are `iterative` and advance independently toward their own exit gates.

## KB dependencies
Read as contract:
- [`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../spec-drafts/playthroughs/spec.md) — the capability
  spec (the principles P1–P6, the manifest model, the assertion-boundary posture §5.8).
- [`../m201-manifest-corpus/overview.md`](../m201-manifest-corpus/overview.md) — the **M201 manifest corpus**, the
  prose use-case declarations the employee-vantage Playthroughs implement against.
- `corpus/ops/demo/playthroughs.md` (NEW, M202) — the runbook / iteration protocol once M202 graduates the spec.
- [`corpus/ops/demo/coverage-protocol.md`](../../../../../corpus/ops/demo/coverage-protocol.md) — the M42e
  precedent (the measure→triage→fix loop the iterative shape mirrors) + the shared e2e foundation.
- [`corpus/ops/demo/stories-spec.md`](../../../../../corpus/ops/demo/stories-spec.md) — the verified-skill chain
  + the employee hero (Maya) the Playthroughs log in as + play through.

## Delivers →
- Passing **employee-vantage Playthroughs** for Maya's Skill Paths / AI Simulations (NON-voice) / Profile
  journeys, in the `playthroughs` rext section (added to the manifest + the per-surface page-object layer, on the
  M202 foundation).
- Corpus/runbook updates as use-cases land (the manifest's employee surface + any landmark-registry additions).
