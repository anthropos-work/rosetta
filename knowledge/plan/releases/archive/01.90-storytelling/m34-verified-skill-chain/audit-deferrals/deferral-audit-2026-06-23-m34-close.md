---
title: "Deferral Audit — M34 close"
date: 2026-06-23
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals; M34 introduced **zero** new deferrals; all surfaced items have a clear fate.

## Summary
- Total deferrals in scope: 0 (M34-originated)
- Single deferrals: 0
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Inherited backlog items reviewed for aging: 4 (all orthogonal — see below)

## Deferral Inventory

**M34-originated: none.** M34's `decisions.md` "Items surfaced + their fate" section records three
items, all already-fated (no deferral written):

| Item | Fate | Status |
|---|---|---|
| `stack.stories.yaml` model, multi-org, hero trio, cockpit | **Fate 2** (already planned) | Owned by M35–M38 `overview.md` (spec §6b / D13). M34's additive `personas` field is the foundation. No deferral. |
| Dashboard surfaces (`membership_skills`, tags, target-roles, feedback) | **Fate 2** (already planned) | Owned by M36 (`overview.md` In: list). M34 is the skill-profile vertical slice only (operator-locked priority). No deferral. |
| `validation_attempt_result_id` FK omission | **Fate 1, landed** | Caught by the integration test, fixed in-section, unit-test-guarded (regression assertion `persona_test.go:181-196`). A PR-review-class issue, explicitly "not a deferral." |

The only `defer`-language hits in M34 docs are (a) `kb-fidelity-audit.md` describing the *historical*
M7c fleet-scope boundary (not an M34 deferral) and (b) `decisions.md` lines stating "no deferral
written" and "not a deferral."

## Repeat-Deferral Patterns
None. (M34 is the first milestone of v1.9; no prior in-release milestone has closed, so there is no
in-release deferral ledger to repeat against.)

## Fate-1 Investigation

The two Fate-2 items (the Stories & Heroes model / multi-org / trio / cockpit, and the dashboard
surfaces) are **genuinely out of M34's vertical-slice scope** and **already owned** by named,
existing-in-this-release milestones (M35–M38 dirs present, `overview.md` In: lists confirmed). A
complete Fate-1 landing now is not feasible (and not desired) — M34's deliberate scope is one hero's
verified-skill spine, the foundation those milestones build on. No partial-slice temptation; the
`personas` blueprint field is the *complete* M34 foundation, not a stub.

## Inherited-backlog aging check (v1.9 design-time backlog, re-checked at M34 close)

The release carries 4 pre-existing backlog items, all recorded in `roadmap-vision.md` with signed
dates + destinations. Re-checked against the aging triggers:

| Item | Area | Aged into M34? |
|---|---|---|
| M33 — ant-academy demo liveness | ant-academy UI (repro-first) | No — orthogonal to seeding |
| DEF-M10-01 — cloud SnapshotStore + S3 media blob bytes | snapshot asset plane | No — orthogonal to seeding |
| DEF-M21-01 — `replayCmd` conn-seam hermetic test | snapshot replay wiring | No — orthogonal to seeding |
| M25-D9 — dev-N taxonomy replay rc=4 | directus/migrate ordering | No — orthogonal to seeding |

None of the 4 touches the verified-skill chain / `stack-seeding` seeders surface M34 modified, so none
is aged-out by the "area touched by a later milestone" trigger. They remain correctly parked in
`roadmap-vision.md`. (This mirrors the v1.9 Phase-0 design-time audit's GREEN verdict.)

## Recommendations
No action. All M34 items are already correctly fated (2× Fate-2 already-owned, 1× Fate-1 landed).
The 4 inherited backlog items stay parked (no fresh fate decision owed — none aged out).

## Applied Changes
None required — the ledger is already clean. This audit is the recorded confirmation.

## Blocking Items (require user decision)
None.
