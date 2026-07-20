---
title: "Deferral Audit — M235 prove-it-lands (close)"
date: 2026-07-20
scope: milestone
invoked-by: close-milestone
---

## Verdict
YELLOW

- One CHRONIC repeat (the 14 pre-existing demo-stack test failures) — already user-dispositioned and homed at the
  v2.5 **release-close** re-anchor (re-confirmed at M232/M233/M234 close). **Not** a fresh unfated repeat, so **not**
  RED-blocking. Flagged for transparency.
- All of M235's OWN carry-forward routes are Fate-3 → M236, **already applied** (M236 `overview.md` `In:` edited in
  iter-08, commit `54eaefe`, user-authorized) with decision records in `decisions.md` (USER-BLOCKER-M235-02
  resolution). No fresh escape-hatch; no NEW unfated repeat introduced by M235.

## Summary
- Total deferrals in scope: 4 (1 chronic-repeat group + 1 release-level TODO + M235's own Fate-3 route cluster + M231's live-corroboration chain-through)
- Single deferrals: 2 (M204 assign-WRITE TODO; M231 live-render corroboration → chained through M235 build → M236)
- Repeat deferrals: 1 (CHRONIC — the 14 demo-stack test-debt failures)
- Chronic patterns flagged: 1

## Deferral Inventory

```yaml
- id: DEF-M235-01   # M235's own carry-forward cluster
  item: "The LIVE (session × action)-lands proof on a cold reset-to-seed (both vantages, 0 ejects) + the NEW content-stories seat-login coverage/Playthrough plumbing + per-section live-calibration checklists (skill-path / ai-labs / academy) + M230 carry-forward live items (ANT_ACADEMY coverage descriptor, next-web clone re-anchor, getPublicCatalogView 2nd manifest)"
  origin_milestone: M235
  first_deferred_on: 2026-07-20
  last_seen_in: m235/decisions.md (USER-BLOCKER-M235-02 resolution) ; m236/overview.md In:
  destination: "M236 prove-on-billion (Fate-3, applied — In: edited iter-08, commit 54eaefe)"
  reason_recorded: "Needs a running stack + a live seeded render to author + CALIBRATE the seat-login sweep; authoring blind ships an incorrect load-bearing harness. User-authorized Fate-3."
  partial_attempted: no  # offline-buildable scope landed IN FULL + unit-proven; only the live measurement routes
- id: DEF-M235-02   # subset of the release-wide chronic (M235's slice)
  item: "6 test_cockpit.py demo-stack failures (part of the 14-fail chronic) — M218 overlay / M53 academy-link stale assertions"
  origin_milestone: M234 (subset first seen); chronic first_deferred v2.4 M224
  first_deferred_on: 2026-07-17 (v2.4 M224 standing-backlog entry)
  last_seen_in: m235/hardening-ledger.md (Verification / Phase 5) — "128 tests, 6 stale failures … unchanged by this pass"
  destination: "v2.5 release-close test-debt re-anchor (part of the 14-fail chronic)"
  reason_recorded: "Stale assertions for intentionally-removed/changed behavior; M235 adds 0 NEW failures (only in-scope stale-assertion updates). Release-close scoped, user-dispositioned."
  partial_attempted: no
- id: DEF-M235-03   # release-level, not M235's
  item: "M204 assign-and-track.UC1 assign-WRITE declared in-manifest unimplemented TODO"
  origin_milestone: M204 (v2.0)
  first_deferred_on: 2026-07 (v2.0)
  last_seen_in: state.md Standing backlog ; roadmap carries to v2.5 release-close
  destination: "v2.5 release-close declared-TODO fate"
  reason_recorded: "Declared-TODO build-reference gap; release-close fate."
  partial_attempted: no
- id: DEF-M235-04   # chain-through, resolved
  item: "M231 live-billion render corroboration of prove-by-render"
  origin_milestone: M231
  first_deferred_on: 2026-07-19
  last_seen_in: m231/progress.md:13 ("deferred to M235 prove-it-lands")
  destination: "chained: M235 built the seedable substrate (landed) → live proof now M236 (Fate-3, DEF-M235-01)"
  reason_recorded: "No content-story sessions seeded on billion yet — that's the M232+ build; M235 built it, M236 proves live."
  partial_attempted: no
```

## Repeat-Deferral Patterns

### REPEAT (CHRONIC): "14 pre-existing demo-stack test failures"
- **First deferred:** v2.4 M224, 2026-07-17, reason: "stale tests for intentionally-removed/changed behavior; predates the touched milestones"
- **Deferred again:** M232 / M233 / M234 close (2026-07-19), reason: identical — HEAD-identical stale assertions, milestone adds 0 new
- **Again:** M235 close (2026-07-20) — M235 touched 0 of them net-new; only in-scope stale-assertion updates
- **Current destination:** v2.5 **release-close** test-debt re-anchor (the correct home for release-spanning test-debt)
- **Time in limbo:** ~3 days across v2.4-tail + M232→M235 (4 milestones)
- **Pattern:** `CHRONIC_DEFER` — identical reason each time. **BUT user-dispositioned with an explicit destination (release-close)** since M233; a fresh un-fated repeat this is NOT. Per the aging policy the destination milestone (v2.5 release-close) has **not yet closed**, so the deferral authority is intact.

## Fate-1 Investigation

### DEF-M235-01 — M235's own carry-forward cluster (live proof + seat-login plumbing)
- **Fate-1 (land now, complete) feasible:** no
- **If no:** Fate 3 (already applied → M236). A complete landing here is infeasible offline — it needs a running stack + a live seeded render to author AND calibrate the seat-login sweep against real selectors / mirror-table scoreboard / per-session score-feedback fences. Authoring blind ships an *incorrect* (not merely uncalibrated) load-bearing harness. The offline-buildable half (fixture matrix + non-sim seeders + manifest projection) landed IN FULL and is unit-proven — this is not a partial slice of one feature, it's the correct offline/live split.

### DEF-M235-02 — 6 test_cockpit.py failures (chronic subset)
- **Fate-1 feasible:** no (release-scoped test-debt) — user-dispositioned, homed at v2.5 release-close.

### DEF-M235-03 — M204 assign-WRITE TODO
- **Fate-1 feasible:** no — release-level declared-TODO, homed at v2.5 release-close. Not M235's domain.

### DEF-M235-04 — M231 live-render corroboration
- **Fate-1 feasible:** resolved by chain — M235 built the seedable substrate (landed); the live corroboration is now DEF-M235-01 → M236.

## Recommendations
- **DEF-M235-01** → LAND-NEXT (Fate 3, **already applied** — M236 `overview.md` `In:` edited iter-08). No further action.
- **DEF-M235-02** → KEEP-DEFERRED (user-dispositioned; v2.5 release-close home; not re-escalated per the standing decision).
- **DEF-M235-03** → LAND-NEXT (Fate 2 — v2.5 release-close already owns it; no edit).
- **DEF-M235-04** → resolved (chained into DEF-M235-01).

## Applied Changes
- None required this audit. DEF-M235-01's Fate-3 application already landed in iter-08 (M236 `overview.md` `In:` + `decisions.md` USER-BLOCKER-M235-02 resolution). The chronic carry's disposition is unchanged from M234's audit.

## Blocking Items (require user decision)
- **None.** The one CHRONIC repeat is user-dispositioned with a live destination (v2.5 release-close, not yet closed); all M235-own routes are Fate-3-applied. No fresh un-fated repeat, no fresh escape-hatch. **Not RED.**
