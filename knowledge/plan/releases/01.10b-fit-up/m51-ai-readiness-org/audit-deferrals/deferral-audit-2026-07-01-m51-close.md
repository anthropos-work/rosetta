---
title: "Deferral Audit — M51 close (milestone scope)"
date: 2026-07-01
scope: milestone
invoked-by: close-milestone
---

## Verdict
RED → **CLEARED** (2026-07-01, user fate decision)

- On first pass this audit was RED: one REPEAT deferral (academy F6, deferred M50→M51) reached M51 close
  unresolved, and a repeat deferral requires an explicit per-item user fate decision before close proceeds.
- **Resolution:** the user fated academy F6 **LAND-NEXT → M53** (Fate-3 override of M53's `No new feature code`
  scope-guard). M53 already cold-rebuilds the demo from scratch — the natural place to seed + verify academy
  content on a clean build. Recorded in M51 `decisions.md` (D-CLOSE-1); M53 `overview.md` Scope `In:` now carries
  the F6 item. The repeat-defer now has an explicit fate → **the RED is cleared and the close proceeds.**
- Every other item is cleanly fated (Fate-1 done / Fate-2 owned / documented prod-finding / cross-release KEEP
  already signed off).

## Summary
- Total deferrals in scope: 5 candidate items
- Single deferrals (cleanly fated, no action): 4
- Repeat deferrals (blocking): 1 (academy F6)
- Chronic/aged patterns flagged: 1 (academy F6 = 2nd deferral)

## Deferral Inventory

- id: DEF-M51-F6 (academy)
  item: "ant-academy course content + hero academy menu-link + non-anonymous academy session"
  origin_milestone: M50 (field review finding F6)
  first_deferred_on: 2026-06-30 (M50 close, D-CLOSE-3, Fate-3 → M51)
  last_seen_in: m51/overview.md:50-58 (annotated candidate scope); NOT done in any M51 iter
  destination: "M51 candidate scope (annotated) — not executed"
  reason_recorded: "not on any M50 gate path; academy is a seeding/content surface (M51's domain)"
  partial_attempted: no
  → REPEAT (M50 → M51), now unresolved at M51 close.

- id: COLD-accept
  item: "COLD reset-to-seed acceptance (whole release from cold)"
  origin_milestone: M50 (D-CLOSE-2)
  destination: "M53 (Fate-2, user-decided) — M53 overview `In:` owns the from-cold rebuild acceptance"
  status: CLEAN — Fate-2, target confirmed. No action.

- id: repin
  item: "consumption-clone re-pin to the release rext tag + .agentspace/rext.tag bump"
  origin_milestone: M47/M49/M50 (push-gated KEEP, release-level)
  destination: "M53 (delivers the .agentspace/rext.tag bump at v1.10.1)"
  status: CLEAN — advances by design each milestone; box-level authoritative bump at M53. Not a repeat-defer.

- id: M314b (prod loadMembers / frozen_tags)
  item: "prod frozen-read still hydrates whole org; would need loadMembers bounded / a frozen_tags column"
  origin_milestone: M51 iter-08 (root-caused), iter-09 (documented)
  destination: "documented as a disclosed demo-perf relaxation + prod-finding in coverage-protocol.md:465,495-496"
  status: CLEAN — a documented prod-finding, NOT a demo fix (correctly out of the tooling's scope). No deferral.

- id: standing-backlog
  item: "DEF-M10-01 (cloud SnapshotStore/S3), DEF-M21-01 (replayCmd hermetic test), M25-D9 (dev taxonomy rc=4)"
  origin_milestone: cross-release (pre-existing)
  destination: "roadmap-vision.md:98-108 — unscheduled cross-release backlog"
  status: CLEAN — orthogonal to v1.10b; already tracked + signed-off historically; not aged into this release. No action.

## Repeat-Deferral Patterns

### REPEAT: "academy F6 — course content + menu-link + non-anonymous session"
- **First deferred:** M50 close, 2026-06-30 (D-CLOSE-3), Fate-3 → M51. Reason: "not on any M50 gate path."
- **Deferred again:** M51 close, 2026-07-01. Reason: M51's 9 iters were consumed ENTIRELY by the AI-readiness
  dashboard perf saga (active-cycle → closed-cycle → deep-link → app read-path demo-patch); the academy work was
  never on the AI-readiness manager coverage-gate path and was not executed. The M51 gate is MET without it.
- **Current destination:** unresolved — annotated to M51 but not done.
- **Time in limbo:** ~1 day across 2 milestones (M50 → M51). Classified **DRIFT_DEFER** (the item is real and keeps
  being routed to whichever milestone's theme is nearest, but its actual theme — a separate ant-academy content
  surface — doesn't match either the M51 AI-readiness perf theme or the M52/M53 consolidation/acceptance themes).

## Fate-1 Investigation

### DEF-M51-F6 — "academy F6"
- **Fate-1 (land now, complete) in M51 close:** NO. The academy content (course chapters / skill-paths + the hero
  menu-link + a non-anonymous session) is a substantial separate seeding/content + wiring surface on the standalone
  Vercel-deployed ant-academy app — genuinely disjoint from M51's AI-readiness dashboard work. Landing it fully at
  close would be new emergent scope, not finishing the milestone's deliverable. A partial ("wire the menu-link only")
  is a disguised deferral the three-fate rule rejects.
- **In-release Fate-2/Fate-3 home:** WEAK. M52 (`Out: new seeding behavior` — it only *expresses* seeding auditably)
  and M53 (`No new feature code (acceptance only)` / `it does not become new scope here`) BOTH explicitly disclaim
  new content work. Routing F6 into either would require overriding their stated scope guards.
- **Cross-release home:** STRONG. `roadmap-vision.md:90` already carries **M207 — Academy coverage** (Playthroughs
  over the separate ant-academy deployment) as a future v2 milestone. The academy is a standalone deployment with its
  own content model; it is out of v1.10b's field-hardening theme (re-sync + demo-up hardening + showcase-org) and not
  on any v1.10b gate path.
- **Verdict:** this is a **user fate decision** (blocking). The honest options are (a) KEEP-DEFERRED-WITH-SIGNOFF →
  a future release (M207 / a v2 academy milestone) as an escape-hatch, or (b) LAND-NEXT via a Fate-3 override of
  M52/M53's scope, or (c) LAND-NOW in M51 close (rejected above as new emergent scope), or (d) DROP.
  **Recommended: (a) KEEP-DEFERRED-WITH-SIGNOFF → v2.0 M207 Academy coverage.** The academy is a coherent separate
  surface already owned by a named future milestone; v1.10b is field-hardening + showcase-org, and the F6 items are
  academy-deployment content, not demo-stack seeding.

## Recommendations

| Item | Recommendation | Fate (RESOLVED) |
|---|---|---|
| DEF-M51-F6 (academy) | auditor recommended KEEP-DEFERRED-WITH-SIGNOFF → v2.0 M207; **user chose LAND-NEXT → M53** | **Fate-3 → M53** (RED cleared) |
| COLD-accept | confirm M53 owns it | Fate-2 (no action) |
| repin | confirm M53 delivers the bump | Fate-2/KEEP (no action) |
| M314b | confirm documented in coverage-protocol.md | not a deferral (documented) |
| standing-backlog | leave in roadmap-vision.md | KEEP (no action) |

## Applied Changes
- **DEF-M51-F6 → LAND-NEXT → M53 (Fate-3).** User fate decision (2026-07-01). The auditor's default recommendation
  was KEEP-DEFERRED-WITH-SIGNOFF → v2.0 M207, but the user chose to LAND-NEXT into **M53** (cold-rebuild acceptance)
  instead — M53 already destroys + cold-rebuilds the demo, so it is the natural place to seed + verify academy
  content on a clean build. Applied:
  - `m53-cold-rebuild-acceptance/overview.md` Scope `In:` now carries the academy-F6 item (course content + hero
    academy menu-link + non-anonymous academy session; AI chat documented-as-absent per the AI-keys policy) with the
    M50→M51→M53 handoff-chain rationale + the Fate-3-override note.
  - `m51-ai-readiness-org/decisions.md` D-CLOSE-1 records the repeat-defer resolution (context → options → choice →
    why).
  This clears the RED verdict; no residual blocking item.

## Blocking Items (require user decision)
**None remaining.** DEF-M51-F6 was the sole blocker; it is resolved (LAND-NEXT → M53, Fate-3, user-decided). The
close proceeds.
