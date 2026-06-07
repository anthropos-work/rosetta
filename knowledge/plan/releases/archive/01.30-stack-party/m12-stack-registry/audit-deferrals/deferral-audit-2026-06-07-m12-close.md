---
title: "Deferral Audit — M12 stack-registry close (release v1.3 \"stack party\")"
date: 2026-06-07
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals; no chronic patterns; no aged-out items requiring fresh re-fate.
- M12 (the first milestone of v1.3) added **zero** deferrals — every scope item landed Fate 1.
- The one inherited cross-release item (DEF-M10-01) was consciously re-fated by the user on 2026-06-07
  during v1.3 design (v1.3 → v1.4), and is parked in `roadmap-vision.md § v1.4 seeds` with a fresh,
  dated rationale — not a stale rubber-stamp.

## Summary
- Total deferrals in scope: **1** (inherited, cross-release, signed)
- Single deferrals: 1
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out (require fresh decision): 0
- New deferrals introduced by M12: **0**

## Deferral Inventory

### Inherited (from v1.2, cross-release)
```yaml
- id: DEF-M10-01
  item: "Directus S3 media blob *bytes* (vs refs-only) + the cloud SnapshotStore backend (M9a-D4)"
  origin_milestone: M10 (v1.2 "set dressing")
  first_deferred_on: 2026-06-06
  last_seen_in: knowledge/plan/roadmap-vision.md:23 (§ v1.4 seeds)
  destination: "v1.4 (re-scoped from v1.3 by user 2026-06-07)"
  reason_recorded: "media/provision API built + unit-tested but unreachable from any entrypoint; gated on
    S3-read access not wired here. v1.2 shipped refs + structure (the floor). Was Fate-2 → v1.3; the user
    scoped v1.3 = the dev/demo-convergence 'stack party' (2026-06-07) and moved the cloud/blob seeds to v1.4."
  partial_attempted: no
```

### M12-introduced
_(none)_ — every M12 scope item (unified registry, first-available-N allocator, both up-paths wired, teardown
frees the slot, the `corpus/ops/rosetta_demo.md` delivery) landed Fate 1 in-milestone. No TODO/FIXME/HACK
markers in any M12-touched file (`stack_registry.py`, both shell CLIs, the test suites).

## Repeat-Deferral Patterns
_(none)_ — DEF-M10-01 has been seen in exactly two distinct *destinations* (v1.3 then v1.4) but never
*re-deferred to dodge work*: the move was a deliberate release-scoping decision by the user (v1.3's theme is
the convergence, not the cloud-store), recorded in the roadmap + `roadmap-vision.md` on 2026-06-07. The item
remains a clean, single, signed cross-release punt. No CHRONIC_DEFER, no DRIFT_DEFER.

## Fate-1 Investigation

### DEF-M10-01 — "S3 media blob bytes + cloud SnapshotStore backend"
- **Fate-1 (land now, complete) feasible:** no
- **If no:** **escape-hatch (cross-release → v1.4), already signed.** A complete landing requires wired
  S3-read access (AWS creds in eu-west-1) + the cloud-store backend swap — genuinely outside v1.3's
  dev/demo-convergence scope, and outside M12's registry/allocator scope entirely (different surface:
  snapshot media, not stack allocation). The user explicitly scoped v1.3 to exclude the cloud/blob seeds
  on 2026-06-07. No new context in M12 unblocks it (M12 touches the port registry, not snapshot media).
- **Re-fate authority:** fresh, dated 2026-06-07 (the v1.3 design decision), parked in `roadmap-vision.md`.

### The `stack-*` skill renames (NOT a deferral — clarification)
The generic `stack-list`/`stack-seed`/`stack-snapshot`/`stack-update` renames + `dev-up`/`dev-down` are
**M14's owned scope** (`roadmap.md § M14`, `In:` list) — a Fate-2 / confirmed-covered item, not a deferral.
M12's `overview.md` correctly lists them under `Out:` with the M14 destination. No action.

## Recommendations
- **DEF-M10-01** → **KEEP-DEFERRED-WITH-SIGNOFF** (escape hatch, already recorded). Confirmed home: v1.4
  (`roadmap-vision.md § v1.4 seeds`). Reason refreshed at v1.3 design time. No fresh sign-off needed this
  pass — the user issued the v1.3→v1.4 re-scope decision two days into this release's life (2026-06-07),
  well within the freshness window.

## Applied Changes
_(none)_ — GREEN verdict, no inline edits required. M12 introduced no deferrals; the single inherited item
is already correctly parked with a current, signed destination. This report is the audit trail.

## Blocking Items (require user decision)
_(none)_
