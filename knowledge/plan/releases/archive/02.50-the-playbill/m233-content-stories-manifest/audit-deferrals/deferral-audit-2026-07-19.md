---
title: "Deferral Audit — milestone (M233 close)"
date: 2026-07-19
scope: milestone
invoked-by: close-milestone
---

## Verdict
YELLOW

- No RED: no undecided repeat-deferral originating in this release's milestones. Every in-release deferral has a
  confirmed in-release destination (Fate 2/Fate 3), and the one cross-release chronic pattern (demo-stack test-debt)
  carries an existing, user-accepted release-level home.
- YELLOW (not GREEN): one chronic cross-release repeat pattern is flagged (demo-stack pre-existing test failures) and
  a handful of single/Fate-3 routings are tracked with accepted destinations.

## Summary
- Total deferrals in scope: 6 records (2 M233-own + 3 inherited M230 Fate-3 clusters + 1 cross-release chronic carry)
- Single deferrals: 5 (2 M233 Fate-2 + 3 M230 Fate-3)
- Repeat deferrals: 1 (demo-stack test-debt — cross-release, v2.4→v2.5)
- Chronic patterns flagged: 1 (CHRONIC_DEFER — demo-stack test-debt)

## Deferral Inventory

```yaml
- id: DEF-M233-01
  item: "bring-up export wiring (up-injected.sh exports content-manifest.json) + cockpit tab read + content-player-<idx> seat registration"
  origin_milestone: M233
  first_deferred_on: 2026-07-19
  last_seen_in: m233/decisions.md:43-50 (Cross-milestone handoffs)
  destination: "M234 (Fate-2, already planned)"
  reason_recorded: "M233 delivers the --content-export verb + honesty-gated manifest; the tab render + seat REGISTRATION is M234's In: list"
  partial_attempted: no
- id: DEF-M233-02
  item: "non-simulation product player-path builders (skill-path / academy)"
  origin_milestone: M233
  first_deferred_on: 2026-07-19
  last_seen_in: m233/decisions.md:52-58
  destination: "M234/M235 (Fate-2, already planned)"
  reason_recorded: "the M233 registry is schema-complete for all 4 products; the fixture carries only simulation, so only the sim player-path builder is exercised; skill-path/academy builders need route fields that land with M234/M235 fixture additions; playerResultPath fail-closes on them until then"
  partial_attempted: no
- id: DEF-M230-01
  item: "formal cold-/demo-up ANT_ACADEMY rendered-card-count gate proof"
  origin_milestone: M230
  first_deferred_on: 2026-07-19
  last_seen_in: m230/carry-forward.md cluster-1
  destination: "M235 (cold reset-to-seed) + M236 (prove-on-billion) — Fate-3, homed"
  reason_recorded: "fill mechanism runtime-proven standalone (59 cards, exact code path); formal cold-/demo-up sweep folds into M235/M236's cold bring-ups"
  partial_attempted: no
- id: DEF-M230-02
  item: "local next-web clone re-anchor (2 drifted demo-patch manifests)"
  origin_milestone: M230
  first_deferred_on: 2026-07-19
  last_seen_in: m230/carry-forward.md cluster-2
  destination: "M235/M236 cold-/demo-up prerequisite — Fate-3, homed"
  reason_recorded: "local clone drift; general demo-hygiene, not academy-specific"
  partial_attempted: no
- id: DEF-M230-03
  item: "getPublicCatalogView 2nd manifest for anonymous academy /library + /free routes"
  origin_milestone: M230
  first_deferred_on: 2026-07-19
  last_seen_in: m230/carry-forward.md cluster-3
  destination: "M235 next-iter queue — Fate-3, homed"
  reason_recorded: "M230 patch fills the employee-authed home grid; anon routes are a faithful follow-on not needed by M230's gate"
  partial_attempted: no
- id: DEF-REL-testdebt
  item: "14 pre-existing demo-stack test failures (drifted next-web urls.ts patches + cockpit/host/purge)"
  origin_milestone: (pre-v2.4; carried v2.4→v2.5)
  first_deferred_on: pre-2026-07-15
  last_seen_in: state.md § Standing backlog + task-prompt standing-carry note
  destination: "v2.5 release-close re-anchor (KEEP-DEFERRED-WITH-SIGNOFF at release scope)"
  reason_recorded: "HEAD-identical failures in files M229–M233 never touched; a demo-stack test-debt harden pass batched at release close, not piecemeal per-milestone"
  partial_attempted: no
```

## Repeat-Deferral Patterns

### REPEAT / CHRONIC_DEFER: "demo-stack pre-existing test failures (test-debt backlog)"
- **First deferred:** pre-v2.4 (8 failures), routed to v2.4 release close.
- **Deferred again:** v2.5 (now 14 failures — drifted next-web urls.ts patches + cockpit/host/purge), routed to v2.5 release-close re-anchor.
- **Current destination:** v2.5 `/developer-kit:close-release` re-anchor pass (batch resolution).
- **Time in limbo:** ~1 release cycle across ≥2 release closes.
- **Pattern signal:** `CHRONIC_DEFER` — same "pre-existing, HEAD-identical, in files this release never touched" reason.
  **Resolution decision EXISTS** (batch re-anchor at release close), so it is a tracked chronic carry, not silent
  erosion. It is a **cross-release, release-scoped** item; M233 (docs + manifest tooling) touched **zero** demo-stack
  test files, so it is out of M233's milestone scope to fix. Flagged here for visibility; the fate is the release
  close's to execute.

## Fate-1 Investigation

### DEF-M233-01 — bring-up export wiring + cockpit read + seat registration
- **Fate-1 (land now) feasible:** no.
- **If no:** Fate-2 — M234 already owns it. M234's `overview.md` `In:` list carries "Per-product sections rendering
  the M233 manifest" + "mint/resolve per-session player seats via roster.go + Clerkenstein". Verified at this close.
  Landing the render/wiring in M233 would be building M234's deliverable early — a scope-boundary violation, not a
  completeness gain. M233's export verb + honesty gate is the complete, self-contained deliverable.

### DEF-M233-02 — non-simulation player-path builders
- **Fate-1 feasible:** no.
- **If no:** Fate-2 — M234's `In:` list explicitly covers the academy section (D5, real seeded progress) and AI-labs
  (D4, presence-only); the skill-path builder needs fixture route fields that arrive with M234/M235's fixture
  additions. The M233 resolver fail-closes (records a clear reason, never fabricates) on a non-simulation link-bearing
  session until then — the correct, complete behavior for M233's fixture (simulation-only). No partial slice.

### DEF-M230-01 / -02 / -03 — M230 carry-forward clusters
- **Fate-1 feasible:** no (all three).
- **If no:** Fate-3, already homed in M230's `carry-forward.md` with confirmed destinations (M235/M236). The
  destination milestones have **not yet closed**, so these are **not aged-out** (no "destination closed without landing"
  trigger). M233 does not touch the academy fill surface, so it did not incidentally unblock any of them. Confirmed
  covered; no re-fate needed at M233 close.

### DEF-REL-testdebt — demo-stack test-debt
- **Fate-1 feasible:** no.
- **If no:** KEEP-DEFERRED-WITH-SIGNOFF (release scope), destination confirmed (v2.5 release-close re-anchor). M233
  touched no demo-stack test files; a fresh landing here would be out-of-scope batch test-debt work belonging to the
  release close. Reason unchanged since v2.4 (HEAD-identical, untouched files). Flagged as chronic; not a M233-close
  blocker.

## Recommendations
- DEF-M233-01 → **LAND-NEXT** (Fate-2, M234 — confirmed, no plan edit).
- DEF-M233-02 → **LAND-NEXT** (Fate-2, M234/M235 — confirmed, no plan edit).
- DEF-M230-01 → **LAND-NEXT** (Fate-3, M235/M236 — homed in carry-forward.md).
- DEF-M230-02 → **LAND-NEXT** (Fate-3, M235/M236 — homed).
- DEF-M230-03 → **LAND-NEXT** (Fate-3, M235 — homed).
- DEF-REL-testdebt → **KEEP-DEFERRED-WITH-SIGNOFF** (release scope; existing v2.4 sign-off, re-confirmed as a v2.5
  standing carry → v2.5 release-close re-anchor). No new M233-close decision required.

## Applied Changes
- None to plan files: every in-release item already has a confirmed home (Fate-2 in M234's `In:` list, verified;
  Fate-3 in M230's `carry-forward.md`, destinations open). No `overview.md` edit needed (no Fate-3 orphan). The M233
  handoff records already sit in `m233/decisions.md`. This audit report is the record.

## Blocking Items (require user decision)
- **None.** No repeat-deferral originating in this release lacks a resolution decision. The one cross-release chronic
  pattern (demo-stack test-debt) has an existing user-accepted release-level destination and is out of M233's scope
  to fix — it is flagged for the v2.5 release close, not blocked here.
