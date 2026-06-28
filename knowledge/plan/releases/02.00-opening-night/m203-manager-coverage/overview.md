---
milestone: M203
slug: manager-coverage
version: v2.0 "opening night"
milestone_shape: iterative
exit_gate: "Same shape as M202, manager-vantage: every declared MANAGER-vantage use case has a PASSING Playthrough on a COLD reset-to-seed demo stack (Dan's core journeys: Workforce funnel + member roster; member drill-down via the activity-dashboard; succession/at-risk via the Growth tab signals), with 0 false-fails over 5 consecutive reset runs."
iteration_protocol_ref: knowledge/plan/spec-drafts/playthroughs/spec.md (the consolidated capability spec; the M201-delivered runbook corpus/ops/demo/playthroughs.md supersedes it as the protocol once M201 graduates it)
status: planned
created: 2026-06-28
last_updated: 2026-06-28
complexity: large
delivers: passing manager-vantage Playthroughs (in the `playthroughs` rext section) for Dan's core journeys — Workforce funnel + member roster, member drill-down (activity-dashboard), succession/at-risk (Growth tab) — added to the manifest + the per-surface page-object layer, on the M201 foundation; plus the corpus/runbook updates as use-cases land
depends_on: M201
spec_ref: knowledge/plan/spec-drafts/playthroughs/spec.md (the consolidated capability spec, v0.3)
---

# M203 — Manager-vantage coverage

## Goal
**Dan's** core **manager** journeys play green as Playthroughs on the M201 foundation:
- **Workforce funnel + member roster** — the mapped→verified funnel + the member list.
- **Member drill-down** — the per-member **activity-dashboard**.
- **Succession / at-risk** — the **Growth tab** signals (succession + at-risk).

A Playthrough **is the user** (P1): drive the real antd UI as Dan would, assert **user-observable outcomes** (the
goal achieved), never pixel/copy specifics (P2).

## Exit gate (objective, machine-verifiable)
**Same shape as M202, manager-vantage** — **every declared manager-vantage use case has a PASSING Playthrough on
a COLD reset-to-seed demo stack** (Dan's journeys above), **with 0 false-fails over 5 consecutive reset runs.** A
run starts from the seeded baseline; after a mutating suite the world is reset via the real `--reset` path (the
M201 reset-to-seed lifecycle — additive re-seed is NOT a reset). The 5-run zero-false-fail bar proves determinism
under mutation.

## Why iterative (not section)
Same as M202 — the manager use-cases are **declarable**, but getting them green against the **real manager antd
UI** (the landmark layer) + the **assertion boundary** is **exploratory**: the failure modes are discovered BY
the attempt. Each iter declares/plays a use case → it fails against the real UI → diagnose → fix in
`rosetta-extensions` (the page-object registry / the dedicated seed / the manifest) → re-play, until the gate is
GREEN for the manager vantage. Build with `/developer-kit:build-mstone-iters`.

## Iteration protocol
The loop is the playthroughs spec / the **M201-delivered runbook**:
[`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../spec-drafts/playthroughs/spec.md) **until M201
graduates it to** `corpus/ops/demo/playthroughs.md`, which then becomes the authoritative protocol.

## Re-scope trigger
A surface that **can't** be driven without a platform edit (the `unimplementable-without-platform-edit` state) →
**escalate, don't edit** the platform (the P3 zero-edit line; the platform repos are read-only this release).
Mirrors the M42m/coverage-sweep re-scope trigger.

## Depends on / Parallel with
- **Depends on:** **M201** — reuses its foundation (the manifest + validator, the per-surface page-object layer,
  the dedicated decoupled seed, the reset-to-seed serial-default runner, the 4-state reporting).
- **Parallel with:** **M202** (employee-vantage). The two share the **landmark-registry + locator index** (the
  M201 page-object layer); each vantage adds its own surfaces/anchors — an **additive merge surface**, not a
  conflicting one. Both are `iterative` and advance independently toward their own exit gates.

## KB dependencies
Read as contract:
- [`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../spec-drafts/playthroughs/spec.md) — the capability
  spec (the principles P1–P6, the manifest model, the assertion-boundary posture §5.8).
- `corpus/ops/demo/playthroughs.md` (NEW, M201) — the runbook / iteration protocol once M201 graduates the spec.
- [`corpus/ops/demo/coverage-protocol.md`](../../../../../corpus/ops/demo/coverage-protocol.md) — the M42m
  precedent (the manager-vantage measure→triage→fix loop the iterative shape mirrors).
- [`corpus/ops/demo/stories-spec.md`](../../../../../corpus/ops/demo/stories-spec.md) — the manager hero (Dan) +
  the org Workforce-Intelligence dashboard surfaces (M36) the manager Playthroughs play through.

## Delivers →
- Passing **manager-vantage Playthroughs** for Dan's Workforce-funnel / member-drill-down / succession journeys,
  in the `playthroughs` rext section (added to the manifest + the per-surface page-object layer, on the M201
  foundation).
- Corpus/runbook updates as use-cases land (the manifest's manager surface + any landmark-registry additions).
