---
title: "Deferral Audit — M16 close (v1.3b dress rehearsal)"
date: 2026-06-08
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

No repeat deferrals; no aged-out items; the single forward-routed item has a clear, already-owned Fate-2 home (M17). M16 is the release's first milestone — nothing inherited *within* v1.3b. The one release-level inherited escape-hatch (DEF-M10-01) is orthogonal to M16, untouched, and not aged out.

## Summary
- Total deferrals in scope: 1 in-release (Fate-2 confirm) + 1 inherited cross-release (escape-hatch, untouched)
- Single deferrals: 1
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0

## Deferral Inventory

1. **Live docker-harness migrate-race BEHAVIOR test** (the runtime half of the ISSUE-7 fence)
   - origin_milestone: M16 (surfaced in the harden pass, Pass 2)
   - first_deferred_on: 2026-06-08
   - last_seen_in: `m16-land-fixes/progress.md` (Pass 2 `TestMigrateRaceGuard` note: "The LIVE docker race test is routed to M17 (Fate 2 — M17 owns idempotency/race + the harness); this is the static half.")
   - destination: M17 (Bring-up re-run safety — idempotency + first-run race)
   - reason_recorded: the static regression fence (`TestMigrateRaceGuard`, 3 tests) landed in M16; the live docker-harness behavior test belongs with M17's race/idempotency work which owns the harness
   - partial_attempted: no — the M16 deliverable (the static fence) is complete; the live behavior test is a *different* test class, not a slice of the same one

2. **DEF-M10-01 — S3 media blob bytes + cloud `SnapshotStore` backend** (inherited, cross-release)
   - origin_milestone: M10 (v1.2)
   - first_deferred_on: 2026-06-07 (v1.2 close, signed escape-hatch)
   - last_seen_in: `knowledge/plan/roadmap-vision.md:28` (v1.4 staged)
   - destination: v1.4 (signed, KEEP-DEFERRED-WITH-SIGNOFF)
   - reason_recorded: cloud store / S3 blob bytes is a v1.4 feature; v1.3b is tooling/docs-only field-hardening — orthogonal
   - partial_attempted: no

## Repeat-Deferral Patterns
None. Item 1 is a single, newly-routed Fate-2. Item 2 has a single signed escape-hatch destination, unchanged since v1.2 close.

## Fate-1 Investigation

### Item 1 — live docker-migrate-race behavior test
- **Fate-1 (land now, complete) feasible:** no
- **Why not Fate-1:** the live behavior test requires a running docker harness + a live Postgres race window — M16 is a docs/publish/rename milestone with NO docker harness in scope. M17 ("Bring-up re-run safety") explicitly owns the `set -e` first-run-race audit (sweeps `up-injected.sh`/`rosetta-demo`/`dev-stack`/`dev-setdress.sh` for the exact ISSUE-7 class) AND the per-component tested idempotency contract that "trips on a real 2nd run" — i.e. it brings the runtime harness. Landing the live test in M16 would duplicate M17's harness setup and split the race-behavior coverage across two milestones.
- **Fate applies:** **Fate 2** — already owned by M17's `In:` scope (verified: `m17-rerun-safety/overview.md:30-39` — the `set -e` first-run-race audit + "an explicit, tested idempotency contract per component (the guard trips on a real 2nd run)"). No `overview.md` EDIT needed — the home already enumerates this work.

### Item 2 — DEF-M10-01
- **Fate-1 feasible:** no — cross-release feature (cloud store / S3 bytes), explicitly OUT of v1.3b ("zero platform feature work; tooling+docs only").
- **Aging check:** no trigger fired. v1.3b just opened (M16 is the first milestone); M16 touched the rename/doc surface, NOT the snapshot-store/S3 area; <1 day elapsed in this release; destination (v1.4) not yet a closed milestone. Authority intact.
- **Fate applies:** escape-hatch, unchanged. No fresh decision required.

## Recommendations
1. **Item 1 → LAND-NEXT (Fate 2).** Confirmed covered by M17. No plan edit (M17 already owns it). Record the confirm in M16 `decisions.md`.
2. **Item 2 → KEEP-DEFERRED-WITH-SIGNOFF (unchanged).** Signed v1.2 escape-hatch, not aged, orthogonal to M16. No action.

## Applied Changes
- M16 `decisions.md`: add M16-D7 recording the Fate-2 confirm for the live docker-migrate-race behavior test → M17 (no M17 overview edit — already owned).
- No roadmap edits. No roadmap-vision edits.

## Blocking Items (require user decision)
None.
