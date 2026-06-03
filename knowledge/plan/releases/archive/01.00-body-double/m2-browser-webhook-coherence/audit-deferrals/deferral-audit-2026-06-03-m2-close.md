---
title: "Deferral Audit — M2 close (milestone scope)"
date: 2026-06-03
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals (no item deferred across ≥2 milestones).
- No aged-out items (every item ≤1 day old; the one forward-routed item is branch topology, not scope).
- All items have a clear fate decision.

## Summary
- Total deferrals in scope: 3 (relevant) + several non-deferrals re-classified
- Single deferrals: 1 (M0-D6, branch topology — destined to the imminent release close)
- Repeat deferrals: 0
- Chronic patterns flagged: 0

## Deferral Inventory

```yaml
- id: DEF-M1-01
  item: "orgclient injection (fake-Clerk-API-server / BAPI) + in-memory store thread-safety"
  origin_milestone: M1
  first_deferred_on: 2026-06-03
  last_seen_in: m1-clerkenstein-backend/decisions.md:52 (M1-D2); progress.md:33 (Fate 3 → M2)
  destination: "M2 (Fate 3 — roadmap M2 In-list updated at M1 close)"
  reason_recorded: "orgclient is app-internal + networked → can't go.mod-replace like authn; needs a
    fake-API-server shared with M2's JS side. Store thread-safety relevant only at injection time."
  partial_attempted: no

- id: DEF-M0-01
  item: "feat/demo-environment → main branch reconciliation"
  origin_milestone: M0
  first_deferred_on: 2026-06-02
  last_seen_in: m0-alignment-framework/decisions.md:36 (M0-D6)
  destination: "release close (/developer-kit:close-release) — explicitly tracked there"
  reason_recorded: "release branch was cut from feat/demo-environment (not main) to preserve planning;
    main reconciliation deferred to release close. M0-D6: 'No three-fate item — branch topology, not scope.'"
  partial_attempted: no

- id: DEF-M2-01
  item: "real-dev-Clerk-for-browser fallback (kept as escape hatch, un-exercised)"
  origin_milestone: M2
  first_deferred_on: 2026-06-03
  last_seen_in: m2/overview.md:39; spec-notes.md:38; decisions.md:17 (M2-D1)
  destination: "n/a — documented contingency, not pending work"
  reason_recorded: "spike resolved Fate 1 (config-only override, no fork); fallback retired to escape-hatch
    documentation only. No work item is outstanding."
  partial_attempted: no
```

**Re-classified as non-deferrals (not work pushed forward):**
- M2 harden "defensive / unreachable residuals" (progress.md) — explicitly NOT tested because the branches
  can't fire through the disarmed paths; documented, not deferred. No shallow coverage-box tests by design.
- M1 authn `go.mod replace` injection recipe — landed Fate 1 in M1 close (blended into clerkenstein.md).

## Repeat-Deferral Patterns
None. DEF-M1-01 was deferred exactly once (M1→M2) and **landed in full in M2** (verified below). No item
crosses ≥2 milestones; no destination was re-pushed forward without resolution.

## Fate-1 Investigation

### DEF-M1-01 — orgclient injection + store thread-safety
- **Fate-1 (land now, complete) feasible:** already landed — RESOLVED in M2.
- **Verification:** `bapi/server.go` (BAPI fake-API-server, 10 methods), `orgclient/store.go` (`sync.Mutex`
  guarding every mutator), `orgclient/concurrency_test.go` (5 `-race` concurrency tests). Recorded as
  M2-D2 (Fate 1). The full carry-forward landed; nothing partial.

### DEF-M0-01 — feat/demo-environment → main reconciliation
- **Fate-1 (land now in M2) feasible:** no — and correctly so. M2's scope is the JS/webhook code + records +
  corpus docs; the branch reconciliation is a **release-level** merge operation that belongs to
  `/developer-kit:close-release` (M2's declared next step). M0-D6 already names that destination.
- **Which fate:** LAND-NEXT (in-release) — owned by the imminent release close, not a new milestone.
  This is branch topology, explicitly flagged by M0-D6 as "not a three-fate scope item."

### DEF-M2-01 — real-dev-Clerk fallback
- **Fate-1 feasible:** n/a — no outstanding work. The spike (M2-D1) resolved the override as config-only;
  the fallback is documentation of a contingency, not deferred implementation. Nothing to land or route.

## Recommendations
- **DEF-M1-01:** LAND-NOW — already satisfied in M2 (M2-D2). No action; record as resolved.
- **DEF-M0-01:** LAND-NEXT — confirmed owned by `/developer-kit:close-release` (per M0-D6). No plan edit
  needed; the destination is already recorded. Surface to the user as the expected next step after M2 close.
- **DEF-M2-01:** No fate needed — not a work deferral.

## Applied Changes
None required. No item needed re-fating, no plan edits, no new decision records. DEF-M1-01 is verifiably
landed; DEF-M0-01 is correctly destined to the release close that immediately follows this milestone close.

## Blocking Items (require user decision)
None. No repeat deferrals, no aged-out items, no chronic patterns. Verdict GREEN — close-milestone proceeds.
