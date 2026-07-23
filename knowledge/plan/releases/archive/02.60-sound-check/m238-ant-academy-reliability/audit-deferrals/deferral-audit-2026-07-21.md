---
title: "Deferral Audit — M238 (milestone scope)"
date: 2026-07-21
scope: milestone
invoked-by: close-milestone
---

## Verdict
**YELLOW**

Single/known deferrals only, each with a clear fate. One standing item (the demo-stack test debt) is a
REPEAT/AGED-OUT pattern — surfaced as a signal, re-fated with a fresh dated decision (Fate-2 → M244); it is
NOT release-scope-breaking, so NOT an escape-hatch and does NOT wake the user.

## Summary
- Total deferrals in scope: **4**
- Single deferrals: **2** (D4; #4→M239 inherited)
- Repeat / aged-out deferrals: **1** (the standing demo-stack test debt)
- Converted deferred → landed this milestone: **1** (KB-1's inventory-fence hygiene gap)
- Chronic patterns flagged: **0 CHRONIC** · **1 repeat** (test-debt, DRIFT-free — same 8 across milestones)

## Deferral Inventory

- **DEF-M238-D4** — full `coverage.spec.ts` billion sweep. origin M238. destination: **M244** (its exit
  gate (c) "the 39 live-browser specs execute green" already owns `coverage.spec.ts`). reason: heavy live
  browser run; probe is unit-proven + live premise validated on billion. partial_attempted: no. → **Fate 2
  (already owned), single, clean.**
- **DEF-M238-KB1-fence** — "no directory-driven fence guards the patch-manifest inventory" (surfaced in the
  Phase 0b KB-fidelity audit / KB-1). origin M238. **RESOLVED THIS MILESTONE:** the harden Pass 2 BUILT it
  (`demo-stack/tests/test_patch_inventory.py` — exact 15 + per-repo pin, mutation-proven RED). → **converted
  deferred → LAND-NOW (done); no longer a deferral.**
- **STANDING — demo-stack test debt (the 8)** — 8 pre-existing demo-stack test failures re-surfaced by
  M238's harden full-sweep. origin: re-baselined at v2.5/M236 (2026-07-20,
  `releases/archive/02.50-the-playbill/m236-prove-on-billion/rebaseline-standing-failures.md`). last_seen_in:
  M238 progress.md "Pre-existing baseline note". destination: **M244** (state.md standing backlog). reason:
  test-side debt asserting deliberately-changed behaviour + one host-conditional. partial_attempted: no.
  → **REPEAT / AGED-OUT → re-fated Fate 2 → M244, fresh dated note (below).**
- **INHERITED — #4 library-flash → M239** (from M237 close). Not in M238 scope; M239 owns it and is still
  open. → **Fate 2, confirmed still owned; no M238 action.**

## Repeat-Deferral Patterns

### REPEAT / AGED-OUT: "standing demo-stack test debt (the 8)"
- **First characterised:** M236 close continuation, 2026-07-20 (the re-baseline doc) — the carried "14 / 6
  pin-drift" label was refuted: 6 were a dirty clone, 0 pin drift; **8 test-side failures remain, 0 product
  defects.**
- **Routed:** v2.5 release close / v2.6 design → standing backlog → **M244** ("one live bring-up discharges
  most of these"), fresh-dated 2026-07-20/07-21 in state.md.
- **Re-surfaced:** M238 harden full-sweep (779 tests), 2026-07-21 — **the identical 8**, composition
  UNCHANGED: `test_cockpit.py` ×6 (4 M234-academy-link-semantics + 2 M218 overlay-JS 30 s-window),
  `test_host_prereqs_m215.py` ×1 (M224 hiring port 13001 not in `_UI_PORTS`), `test_purge.py` ×1
  (macOS-environmental, self-documented). **NONE in a file M238 touched → 0 M238 regressions.**
- **Time in limbo:** carried across M236 → M237 → M238 (≥ 2 milestones → AGED-OUT trigger); feature area
  (demo-stack tests) touched by M238 (added `test_patch_inventory.py` / `test_academy_fs_published_body.py`),
  another aging trigger.
- **Pattern class:** **not CHRONIC** — the composition is stable and each failure is precisely diagnosed
  (test-side, deliberately-changed behaviour). It is a **standing hygiene debt that keeps riding** because no
  v2.6 milestone is a "test-debt cleanup" milestone and M244 is the natural sweep point (it runs the full
  suite live). The signal worth surfacing: it has now ridden **three** consecutive v2.6-adjacent milestones.

## Fate-1 Investigation

### DEF-M238-D4 — full `coverage.spec.ts` billion sweep
- **Fate-1 feasible:** no — it's a heavy authenticated live-browser crawl on billion, out of a docs+tooling
  milestone's altitude; the probe logic IS unit-proven (139 e2e unit specs) and its live premise was
  validated directly on billion this milestone.
- **Fate:** **Fate 2** — M244's exit gate (c) already owns the live `coverage.spec.ts` execution. No plan edit.

### STANDING — the 8 demo-stack failures
- **Fate-1 (land now, complete) feasible in M238:** **no — out of scope.** The edits (re-point 4 academy
  assertions at M234 semantics, invert/retire 2 overlay-JS assertions, add `13001` to one expected list,
  `skipUnless(Linux)` on the purge test) are all cheap and product-risk-free (the re-baseline doc's Fate-1
  recommendation), but they are UNRELATED to ant-academy reliability. Landing them in M238 would be
  scope-bleed, not Fate-1 discipline — the three-fate rule's Fate-1 default is the milestone's OWN
  deliverable scope.
- **Fate:** **Fate 2 → M244** (already on the books in state.md's standing backlog). **Elevated note for
  M244:** these are cheap LAND-able TEST edits — M244 should discharge them by EDITING the tests, not only
  "via a live bring-up"; 6 of the 8 (`test_cockpit`, `test_host_prereqs`) don't touch a live stack at all.
  They have now ridden three v2.6-adjacent milestones; M244 is the expiry point.

## Recommendations
1. **DEF-M238-D4** → **Fate 2 (confirm, no edit)** — owned by M244 exit gate (c).
2. **DEF-M238-KB1-fence** → **already LANDED** (harden Pass 2); close the item.
3. **Standing 8** → **Fate 2 → M244** with a FRESH DATED re-confirmation (2026-07-21): identical set, 0 drift,
   0 M238 regressions, still test-side debt. Recorded as M238 decision **D5**. Surface the repeat pattern to
   the orchestrator (YELLOW).
4. **Inherited #4** → **Fate 2, confirmed still owned by M239** (open); no M238 action.

## Applied Changes
- This report (dated 2026-07-21, M238 `audit-deferrals/`).
- M238 `decisions.md` **D5** — the standing-8 re-fate (Fate 2 → M244) + D4/KB-1-fence confirmations.
- No `progress.md` LAND-NOW additions (no Fate-1 items — KB-1 fence already landed in harden).
- No target-plan edits (all routings are Fate 2 already-owned; state.md standing-backlog line gets a fresh
  date at Phase 10 per the state.md contract).

## Blocking Items (require user decision)
**None.** The one REPEAT/AGED-OUT item (standing 8) is re-fated Fate 2 → M244 with a fresh dated decision and
is not release-scope-breaking (per close orchestration: YELLOW, surfaced, not escape-hatch, no user wake).
