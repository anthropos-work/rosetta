---
title: "Deferral Audit — M9a close (milestone scope)"
date: 2026-06-06
scope: milestone
invoked-by: close-milestone
---

## Verdict
**GREEN** — M9a is the FIRST milestone of v1.2, so there are no inherited in-release milestone
deferrals. M9a's own ledger carries one designed-out-of-scope item (the v1.3 cloud/S3 `SnapshotStore`
backend) with a clean architectural seam, and confirms its three "Out:" surfaces are owned Fate-2 by
M9b/M10/M11. No repeat-deferral, no chronic pattern, no aged-out item (everything is hours old). Not a
blocker. Close proceeds.

## Summary
- Total deferrals in scope: 4
- Single deferrals: 4 · Repeat (≥2 milestones, unresolved): 0 · Chronic: 0 · Aged-out: 0
- LAND-NEXT (Fate 2, already owned in-release): 3
- KEEP-DEFERRED-WITH-SIGNOFF (escape hatch, v1.3 — pre-signed by user at design): 1
- LAND-NOW (Fate 1): 0 (none applicable — all framework scope already landed in M9a)
- Code TODO/FIXME/HACK in milestone-touched source: **0**

## Deferral Inventory + Fate verdicts

| # | Item | Origin | Recorded | Fate | Destination |
|---|---|---|---|---|---|
| 1 | **Taxonomy surface** — public skiller catalog capture/replay + pgvector index rebuild | M9a overview `Out:` | 2026-06-05 | **LAND-NEXT (Fate 2)** | **M9b** — `In:` list owns it (overview confirms: "Prove the M9a framework on the real ~2.1 GB taxonomy surface") |
| 2 | **Directus content surface** — public template library → per-stack content store | M9a overview `Out:` | 2026-06-05 | **LAND-NEXT (Fate 2)** | **M10** — `In:` list owns it (overview confirms: "Capture the public Directus content library … per-stack content store") |
| 3 | **Recipes / presets / corpus product layer** | M9a overview `Out:` | 2026-06-05 | **LAND-NEXT (Fate 2)** | **M11** — `In:` list owns it (presets + recipes + `/demo-snapshot` skill) |
| 4 | **Cloud/S3 `SnapshotStore` backend** (+ AI-content, shareability) | M9a-D4 / M9a-Q4→D | 2026-06-06 | **KEEP-DEFERRED-WITH-SIGNOFF (escape hatch)** | **v1.3** — user-signed at design (2026-06-05/06); clean seam: the `SnapshotStore` interface (5 methods) makes the swap a backend re-implementation. localfs is the only M9a backend by design. |

## Repeat-Deferral Patterns
None. Items #1–#3 are scope partitioning across the v1.2 milestone chain (M9a frames, M9b/M10 fill the
surfaces, M11 curates) — the M7a→M7c precedent, not erosion. Item #4 (cloud store) is a deliberate
v1.3 layering decision (localfs now, cloud swap later behind a stable interface) signed by the user at
design time — a forward-design KEEP, not a chronic/drift defer.

## Fate-1 Investigation
No LAND-NOW candidate. M9a's framework scope (extension + capture-source policy + tenant firewall +
`.agentspace` manifest store + fidelity-DNA + `/db-query` port + reference surface) all landed in this
milestone — all 9 section checkboxes checked. The four inventory items are correctly out-of-milestone:
#1–#3 belong to later v1.2 milestones that explicitly own them (Fate 2, no plan edit needed); #4 is a
pre-signed cross-release punt with an architectural seam already built (the `SnapshotStore` interface).

## Carry-from-design check (v1.1 → v1.2 carry-forwards, for context)
The v1.2 design-time audit (`.agentspace/scratch/deferral-audit-2026-06-05.md`, GREEN) routed 6 v1.1
carry-forwards. Status at M9a close: #1 snapshot-mechanism = NOW IMPLEMENTING (M9a–M11, the active
work); #2 AI-content + #3 shareability = confirmed v1.3 (roadmap-vision); #4 deploy/express CI gate =
largely landed in M8 (wired into clerkenstein `alignment.yml`, validated 9/9 locally) — the residual
"needs an org-secret runner" piece is environmental, a candidate not a committed deferral; #5/#6 = KEEP
(design non-goals). No carry-forward regressed into a repeat.

## Recommendations
- #1–#3 → **LAND-NEXT (Fate 2)**: confirmed already owned by M9b/M10/M11 `In:` lists. No plan edit.
- #4 → **KEEP-DEFERRED-WITH-SIGNOFF**: v1.3, user-signed at design; interface seam in place. No action.

## Applied Changes
None required. All inventory items already have correct fate records (M9a-D4 / overview `Out:` /
downstream `In:` lists). No plan mutation, no new decision needed.

## Blocking Items (require user decision)
None. No repeats-without-resolution, no aged-out, no chronic. **close-milestone proceeds.**
