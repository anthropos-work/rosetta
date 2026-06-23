---
title: "Deferral Audit — v1.9 \"storytelling\" release close (release scope)"
date: 2026-06-23
scope: release
invoked-by: close-release
---

## Verdict
GREEN

- **Zero open deferrals across the release.** Every item v1.9 surfaced reached a terminal Fate-1 landing
  inside the release. No repeat-deferral pattern, no chronic pattern, no aged-out item, **zero escape-hatch
  (`RELEASE-SCOPE-DEFER`) entries.** The inherited unscheduled backlog is unchanged, orthogonal to v1.9's
  surface, and freshly GREEN-audited at every milestone close.

## Summary
- Total deferrals in scope (v1.9 own): 2 (both terminally landed Fate-1 within the release)
- Single deferrals open at close: 0
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out items: 0
- Escape-hatch (cross-release) entries: 0
- Inherited unscheduled backlog (carried, not re-deferred): 4 items (M33, DEF-M10-01, DEF-M21-01..04, M25-D9)

## Deferral Inventory

### v1.9's own items (both terminally resolved Fate-1 in-release)

```yaml
- id: DEF-M34-01   # = #M34-D7
  item: "multi-hero index-collision guard + short-role-pool flat top-up"
  origin_milestone: M34
  first_deferred_on: 2026-06-23
  last_seen_in: m34-verified-skill-chain/decisions.md:76 (D-M34-7, Fate-3 → M35)
  destination: M35 (Fate-3 annotate — M35's overview.md In: already owned the multi-org roster)
  reason_recorded: "M34-benign at 1-hero; reachable once M35 adds the trio — route to the roster milestone
    rather than over-build M34's vertical slice."
  partial_attempted: no
  RESOLUTION: "LANDED IN FULL as D-M35-4 at M35 close (BOTH parts: the len(heroes)<=size validation +
    declaration-order collision-free slots + warning; the short-role-pool flat top-up). The chosen fix was
    strictly better than the routed warn-on-hash idea. Fate-3 routed-AND-delivered within the release."

- id: DEF-M38-01   # = M38-D7 → M38-D8
  item: "employee heroes export org_role=admin; org_role should follow vantage (end-user→member, manager→admin)"
  origin_milestone: M38
  first_deferred_on: 2026-06-23 (Fate-3 routed to the v1.9 close-review by design)
  last_seen_in: m38-presenter-cockpit/decisions.md:54 (M38-D7) + :78 (re-fated to M38-D8)
  destination: v1.9 close-review (the orchestrator explicitly routed the triage here)
  reason_recorded: "single-sourced from roleForIndex (M35 seam); a correct fix moves three writes in lockstep
    [membership row + casbin g2 + roster/JWT claim]; orchestrator flagged 'weigh here vs close-review, do NOT
    force a large change.' Behaviorally consistent, vantage-imprecise."
  partial_attempted: "a crashed prior close attempt left an unfinished users.go edit — the M38 close
    COMPLETED it (not a partial landing)."
  RESOLUTION: "Re-fated Fate-3 → Fate-1 LAND-NOW at M38 close = M38-D8. Single roleForHero helper at the M35
    seam, both call-sites (users.go seed loop + roster.go BuildRoster) in lockstep, pinned by
    TestBuildRoster_OrgRoleVantageFaithfulAndLockstep + TestRoleForHero. The close-review's code-quality +
    adversarial scans BOTH caught that the crashed attempt left roster.go on the OLD roleForIndex — completed
    properly. As the LAST milestone there was no later in-release milestone to route to; Fate-1 was the only
    correct outcome. Zero platform-repo edits preserved."
```

### Inherited unscheduled backlog (carried — NOT re-deferred by v1.9)

```yaml
- M33 — ant-academy demo liveness (repro-first; deferred v1.7 design 2026-06-15)
- DEF-M10-01 — cloud SnapshotStore / S3 media blob bytes (re-signed → backlog at v1.5 design 2026-06-11)
- DEF-M21-01..04 — prop-room residuals (replayCmd hermetic test landed v1.5 close so it survives merge)
- M25-D9 — dev-N taxonomy replay rc=4 (dev migrate-ordering nuance; non-fatal)
```

All four are in `roadmap-vision.md` § "Unscheduled backlog". All orthogonal to v1.9's
`stack-seeding` / `clerkenstein` / `demo-stack` / `stack-injection` surface. GREEN-audited at the M34, M35,
M36, M37, and M38 closes (the most recent M37/M38, 2026-06-23 — same day, well inside the 3-month aging
window). None re-touched by any v1.9 milestone.

## Repeat-Deferral Patterns
None.

- The two v1.9 items are each **single** findings that terminated in a Fate-1 landing inside the release
  (one Fate-3-routed-and-delivered, one Fate-3-re-fated-to-Fate-1). Neither was pushed forward a second time.
- The inherited backlog has not been re-deferred — it is **carried** (a stable, signed-off unscheduled
  backlog with no in-development target), audited GREEN repeatedly, unchanged. Carrying a backlog item with a
  fixed "no target version" home is not a repeat-deferral; a repeat-deferral is the same item pushed across
  ≥2 milestones' *in-scope* ledgers. Neither inherited item entered any v1.9 milestone's scope.

## Aging
No aged-out items. The two v1.9 items were created and resolved on 2026-06-23. The inherited backlog items
each have a stable, re-signed home and were last GREEN-audited the same day (M37/M38 close); no aging trigger
(≥2 milestones in-scope / ≥3 months / destination-closed-without-landing / area-touched-by-later-milestone)
fired — v1.9 did not touch the snapshot-store / S3 / dev-migrate / ant-academy surfaces those items live on.

## Fate-1 Investigation

### DEF-M34-01 (#M34-D7) — RESOLVED Fate-1-within-release
- **Fate-1 feasible:** n/a — already terminally landed. Fate-3-routed to M35, delivered in full as D-M35-4.

### DEF-M38-01 (M38-D7/D8) — RESOLVED Fate-1
- **Fate-1 feasible:** YES — verified, and **already landed** at M38 close as M38-D8 (the vantage-faithful
  roleForHero single-source + two-call-site lockstep + regression tests). Bounded, tooling-only, zero
  platform-repo edits. The whole point of the Stories & Heroes release.

### Inherited backlog (DEF-M10-01 / DEF-M21-01..04 / M25-D9 / M33)
- **Fate-1 feasible now:** NO — each is orthogonal to v1.9's surface; landing any would WIDEN release scope
  beyond the believable-demo-narrative thesis (S3 read access not wired for DEF-M10-01; the replayCmd seam /
  dev-migrate ordering / ant-academy native-process liveness are different subsystems). No v1.9 context
  changed their calculus. **KEEP** — unchanged unscheduled backlog, no fresh decision required (not aged,
  not repeat, not re-touched).

## Recommendations
1. **DEF-M34-01 → already LAND-NOW-complete** (landed as D-M35-4). No action.
2. **DEF-M38-01 → already LAND-NOW-complete** (landed as M38-D8). No action.
3. **Inherited backlog → KEEP** — unchanged, orthogonal, freshly GREEN. No action; stays in
   `roadmap-vision.md` § Unscheduled backlog.

## Applied Changes
None required. Every v1.9 deferral already terminated Fate-1 in-release at its milestone close; the inherited
backlog is unchanged and correctly homed in `roadmap-vision.md`. This release-scope audit confirms the
milestone-level audits (M34/M35/M36/M37/M38, all GREEN) compose to a GREEN release with **zero escape-hatch
deferrals** — the release ships clean.

## Blocking Items (require user decision)
None. Zero repeat-deferrals, zero aged-out, zero open escape-hatch. The audit is GREEN — `/developer-kit:close-release`
may proceed.
