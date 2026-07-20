---
title: "Deferral Audit — milestone (M234 close)"
date: 2026-07-19
scope: milestone
invoked-by: close-milestone
---

## Verdict
YELLOW

- Single deferrals only, each with a user-accepted in-release destination.
- One CHRONIC repeat (the 14 pre-existing demo-stack test failures) — already user-dispositioned and homed at the v2.5 release-close re-anchor (re-confirmed at M232/M233 close); **not** a fresh unfated repeat, so **not** RED-blocking. Flagged for transparency.
- 0 blocking items requiring a fresh user decision at this close.

## Summary
- Total deferrals in scope: 4
- Single deferrals: 3 (all Fate-2/Fate-3, homed)
- Repeat deferrals: 1 (CHRONIC — the 14 demo-stack test-debt failures)
- Chronic patterns flagged: 1

## Deferral Inventory

```yaml
- id: DEF-M234-01
  item: "Non-simulation product fixtures (ai-labs / academy / skill-path) + prove-every-CTA-lands live"
  origin_milestone: M234
  first_deferred_on: 2026-07-19
  last_seen_in: m234/decisions.md:46 (D-M234-5); content-stories-spec.md §6/§7.6
  destination: "M235 prove-it-lands (In: + exit_gate) [+ M230 for academy]"
  reason_recorded: "M234 is the render HALF; the renderer handles all dispositions (unit-proven). Populating INTERESTING real-shaped fixtures per product + proving every CTA lands on a non-empty result page is M235's whole purpose."
  partial_attempted: no

- id: DEF-M234-02   # subset of the chronic
  item: "6 pre-existing test_cockpit.py failures (removed per-hero academy-CTA x4 + v2.3.1 overlay-30s x2)"
  origin_milestone: pre-v2.5 (v2.3.1 / v2.4 era)
  first_deferred_on: 2026-07-17 (v2.4 M224 standing-backlog entry)
  last_seen_in: m234/decisions.md:59 ("Pre-existing 6-fail cockpit carry — UNCHANGED (Fate-2, release-close)")
  destination: "v2.5 release-close test-debt re-anchor (part of the 14-fail chronic)"
  reason_recorded: "Stale tests for intentionally-removed/changed behavior; M234 adds 0 new failures (verified 249 pass / 6 fail, exactly those 6). Release-close scoped."
  partial_attempted: no

- id: DEF-M230-carry   # inherited, homed in M235
  item: "Formal cold-/demo-up academy card-count sweep + local next-web re-anchor + getPublicCatalogView anon-routes"
  origin_milestone: M230
  first_deferred_on: 2026-07-19
  last_seen_in: m235/overview.md:32-36 (In: — Inherited from M230, Fate-3); m230/carry-forward.md
  destination: "M235/M236"
  reason_recorded: "M230 gate MET-BY-PROXY; the formal cold-demo-up sweep + re-anchor belong in the prove-it-lands / prove-on-billion milestones."
  partial_attempted: no

- id: DEF-M231-carry   # inherited, homed in M235
  item: "Live billion render corroboration of content-story sessions"
  origin_milestone: M231
  first_deferred_on: 2026-07-19
  last_seen_in: m231/progress.md:13; m235/exit_gate + m236
  destination: "M235 (render) / M236 (billion)"
  reason_recorded: "No content-story sessions seeded on billion yet — that's the M232+ build; live corroboration is prove-it-lands/prove-on-billion."
  partial_attempted: no
```

## Repeat-Deferral Patterns

### REPEAT (CHRONIC_DEFER): "14 pre-existing demo-stack test failures"
- **First deferred:** v2.4 M224, 2026-07-17, reason: "8 pre-existing, non-milestone test failures in files the milestone never touched"
- **Grew + deferred again:** v2.4 close-scoping → "14 pre-existing demo-stack failures" standing carry (6 x test_cockpit.py removed-academy-CTA/overlay + test_purge + test_reap + others)
- **Re-confirmed:** v2.5 M232 close (YELLOW), M233 close (YELLOW), M234 close (this audit)
- **Current destination:** v2.5 **release-close** test-debt re-anchor (state.md standing backlog + Headline numbers both home it)
- **Time in limbo:** ~2 days across releases; a HEAD-identical pre-existing test-debt backlog, not milestone-caused
- **Fate:** KEEP-DEFERRED-WITH-SIGNOFF — user-dispositioned at the v2.4/v2.5 release scoping; the final re-anchor + sign-off is a v2.5 **release-close** deliverable (M236's close-release). M234 adds **0** new failures.

## Fate-1 Investigation

### DEF-M234-01 — non-sim fixtures + prove-it-lands
- **Fate-1 (land now, complete) feasible:** no
- **If no:** Fate-2 (already-owned by M235). M235's `In:` list ("seed the full required set per type", "drive every session×action via a Playthrough + coverage descriptor") + its `exit_gate` ("each product either passes or is declared with a documented fate, AI-labs feasibility answered explicitly") + the M230-inherited academy bullet own it in full. Landing fixtures + a live browser prove-it-lands loop in M234 would duplicate M235's entire iterative scope — a mis-scoping, not a Fate-1 slice. The renderer already handles every disposition (unit-proven at M234), so nothing is stranded.

### DEF-M234-02 — 6 pre-existing cockpit fails (chronic subset)
- **Fate-1 feasible:** no (not M234-introduced; stale tests for intentionally-removed behavior across the demo-stack test-debt backlog)
- **If no:** KEEP-DEFERRED-WITH-SIGNOFF via the chronic (release-close). Fixing 6 stale tests piecemeal here would fragment the 14-fail test-debt harden the release-close owns.

### DEF-M230-carry / DEF-M231-carry — inherited academy + live-render
- **Fate-1 feasible:** no (require a live cold `/demo-up` + billion — out of M234's unit-scoped render half)
- **If no:** Fate-2/Fate-3, homed in M235 (`In:`) / M236 (billion). Confirmed present in M235's overview.

## Recommendations
- **DEF-M234-01** → **LAND-NEXT** (Fate-2, M235) — confirmed covered; already recorded in m234/decisions.md D-M234-5. No plan edit needed.
- **DEF-M234-02** → **KEEP-DEFERRED-WITH-SIGNOFF** (release-close, chronic) — already recorded in m234/decisions.md + state.md standing backlog. No fresh wake per prior user disposition.
- **DEF-M230-carry** → **LAND-NEXT** (Fate-3, M235/M236) — confirmed present in m235/overview.md `In:`.
- **DEF-M231-carry** → **LAND-NEXT** (Fate-2/3, M235/M236) — confirmed.

## Applied Changes
- No new edits required — every fate decision is already recorded:
  - DEF-M234-01: m234/decisions.md `D-M234-5` (+ content-stories-spec.md §6/§7.6).
  - DEF-M234-02 + the chronic: m234/decisions.md "Pre-existing 6-fail cockpit carry" + state.md standing backlog + Headline numbers.
  - DEF-M230/M231 carries: m235/overview.md `In:` + m230/carry-forward.md.
- This report is the audit trail for the M234 close.

## Blocking Items (require user decision)
- None. The one chronic repeat is already user-dispositioned and homed at the v2.5 release-close re-anchor; per the standing disposition it is not re-woken here. It surfaces at the v2.5 **close-release** Phase 1b for the final fresh sign-off.
