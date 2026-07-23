---
title: "Deferral Audit — M240 content-stories fidelity (close)"
date: 2026-07-22
scope: milestone
invoked-by: close-milestone
---

## Verdict
YELLOW — single deferrals with user-accepted destinations; one tracked repeat (the standing demo-stack
debt) already fated with a named expiry. No un-fated repeat, no chronic pattern, no RED blocker.

## Summary
- Total deferrals in scope: 4
- Single deferrals: 3 (M240 video exhibit; DEF-M239-01; the 9th reap collision)
- Repeat deferrals: 1 (the standing demo-stack test debt — fated Fate-2 → M244, expiry named)
- Chronic patterns flagged: 0

## Deferral Inventory

```yaml
- id: DEF-M240-01
  item: "Content-stories VIDEO exhibit — re-pin a hiring-voice cell to a recorded session, provision the
         Bunny.net recording signing keys, wire the exhibit-by-reference render, flip chime_status='completed'"
  origin_milestone: M240
  first_deferred_on: 2026-07-22
  last_seen_in: progress.md (Defect 2 section, [~]) + decisions.md (Bunny-key search result)
  destination: "M244 (Fate-3, user pre-approved 2026-07-22)"
  reason_recorded: "the Bunny.net recording signing keys (BUNNY_RECORDING_CDN_TOKEN_KEY +
         BUNNY_RECORDING_PULL_ZONE_HOST) are genuinely absent from this box's entire dev-stack; flipping
         chime_status without them = a broken 500 player (regression vs faithful not_available). Voice
         presence-only IS the v2.6 deliverable per user decision 2026-07-22."
  partial_attempted: no   # held atomic — the posture/spec/gender-contract landed; the port did not

- id: DEF-M239-01
  item: "Make the demo build fail loudly on ENOSPC (the disk-full class that surfaced as a cryptic redis exited(1))"
  origin_milestone: M239
  first_deferred_on: 2026-07-21
  last_seen_in: M244 overview.md In-list (already homed)
  destination: "M244 (Fate-3, M239-D12)"
  reason_recorded: "build-time loud-abort follow-on to M239's VM-disk pre-flight fix; non-blocking"
  partial_attempted: no

- id: DEF-M239-13
  item: "9th standing failure — test_a_RACED… hardcodes cockpit port 17700, collides with a live demo-1 (test-isolation)"
  origin_milestone: M239
  first_deferred_on: 2026-07-21
  last_seen_in: M244 overview.md In-list (already homed, fix recipe in M239 decisions.md D13)
  destination: "M244 (Fate-3, M239-D13)"
  reason_recorded: "reap.sh is correct; the test hardcodes a port that collides with a running demo; a clean
         reset-to-seed resolves it but land the isolation fix; non-blocking"
  partial_attempted: no

- id: STANDING-DEMO-STACK-8
  item: "8 (macOS) / 7 (Linux) standing demo-stack test failures — host-dependent, 0 real defects, 0 pin drift"
  origin_milestone: M236 (re-baselined) → re-confirmed M238, M239
  first_deferred_on: 2026-07-20 (re-baselined) ; M238-D5 fresh dated 2026-07-21
  last_seen_in: state.md Standing backlog + M244 (owner)
  destination: "M244 (Fate-2, M238-D5; M244 = the named expiry point)"
  reason_recorded: "identical set re-surfaced 0 regressions across M238/M239/M240; M244 should discharge by
         EDITING the tests (6 of 8 need no live stack), not only via a live bring-up"
  partial_attempted: no
```

## Repeat-Deferral Patterns

### REPEAT: "standing demo-stack test debt (8 macOS / 7 Linux)"
- **First deferred:** M236, 2026-07-20 (re-baselined as 0-defect host-dependent debt)
- **Deferred again:** M238, 2026-07-21 (M238-D5, fresh dated, Fate-2 → M244)
- **Re-confirmed:** M239, 2026-07-21 (identical set, 0 M239 regressions)
- **Seen (not re-deferred) at:** M240 close — M240 did NOT touch the demo-stack test harness (M240 code = `stack-seeding`); the set does not reproduce from M240's work, so M240 adds no new deferral here.
- **Current destination:** M244 (the v2.6 live closer), with the expiry named explicitly in state.md ("Now ridden ≥3 v2.6-adjacent milestones — M244 is the expiry point").
- **Time in limbo:** ~2 days across 3 v2.6 milestones; the expiry milestone (M244) exists, is scoped, and is imminent.
- **Pattern class:** DRIFT-tracked-with-named-expiry, NOT CHRONIC_DEFER — it carries a fresh dated Fate-2 decision and a named milestone destination + expiry, and it is host-state debt (not feature scope) with 0 real defects.

## Fate-1 Investigation

### DEF-M240-01 — "content-stories VIDEO exhibit"
- **Fate-1 (land now, complete) feasible:** no
- **If no:** Fate-3 → M244. A complete landing requires the Bunny.net recording signing keys, which are
  genuinely absent from this box's entire dev-stack (verified values-blind: no populated value in any real
  `.env`, the `.agentspace/secrets` source, or the compose — only key-name templates). Flipping
  `chime_status='completed'` without them renders a 500 player — strictly WORSE than the faithful
  `not_available`. Building the port against media that cannot be provisioned or playback-verified is
  untestable scaffolding the three-fate rule forbids. The keys may be reachable on `billion` (the M244 host),
  so M244 is the right home: land the exhibit live IF the keys resolve there, else keep voice presence-only.
  User pre-approved the M244 routing 2026-07-22. NOT escape-hatch (in-release, Fate-3).

### DEF-M239-01 / DEF-M239-13 — "ENOSPC loud-abort" / "reap 17700 collision"
- **Fate-1 feasible:** no (not M240's domain — demo-stack build/test harness, not `stack-seeding`)
- **If no:** Fate-3, already HOMED in M244's In-list (M239-D12 / M239-D13). No action needed at M240 close.

### STANDING-DEMO-STACK-8 — "standing demo-stack test debt"
- **Fate-1 feasible:** no (not M240's domain; M240 touched `stack-seeding`, not the demo-stack test harness)
- **If no:** Fate-2, already OWNED by M244 (M238-D5, re-confirmed M239). M244 = the named expiry. No new
  decision at M240 close; the item is not a fresh un-fated repeat.

## Recommendations
- **DEF-M240-01** → **LAND-NEXT** (Fate-3 → M244). Apply: add to M244 overview.md In-list; record the handoff
  in M240 decisions.md. User pre-approved 2026-07-22 — no escape-hatch prompt.
- **DEF-M239-01, DEF-M239-13** → **LAND-NEXT** (Fate-3, already in M244 In-list). No edit — confirmed covered.
- **STANDING-DEMO-STACK-8** → **LAND-NEXT** (Fate-2, already owned by M244, expiry named). No edit — confirmed
  covered; surfaced as the legitimate tracked-repeat YELLOW.

## Applied Changes
- M244 `overview.md` In-list: added the DEF-M240-01 video-exhibit bullet (Fate-3, user pre-approved).
- M240 `decisions.md`: recorded the DEF-M240-01 → M244 handoff + the voice-presence-only close disposition.

## Blocking Items (require user decision)
None. No un-fated repeat, no aged-out item requiring fresh fate this pass, no escape-hatch. The standing
demo-stack repeat carries a fresh dated Fate-2 decision with a named expiry milestone (M244) — surfaced as
YELLOW, not RED. DEF-M240-01's M244 routing is user-pre-approved (2026-07-22), applied without a new prompt.
