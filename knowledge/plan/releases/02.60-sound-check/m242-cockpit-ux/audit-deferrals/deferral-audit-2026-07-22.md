---
title: "Deferral Audit — milestone M242 (cockpit UX)"
date: 2026-07-22
scope: milestone
invoked-by: close-milestone
---

## Verdict
YELLOW

Single/inherited deferrals only, every one with an accepted destination and a **named expiry milestone**
(M244 / M243). M242 introduced **zero new deferrals** — a clean complete section close (all 3 UX sections +
spec-extension + delivers landed + 2 harden passes). The one repeat/aging pattern (the standing demo-stack
test debt) is **already fated Fate-2 → M244 with M244 as the confirmed expiry**, re-affirmed with fresh dates
at the M238/M239/M240/M241 closes; M242 surfaces the same 6-of-9 subset (`test_cockpit.py`) but requires **no
new user decision**. No RED conditions.

## Summary
- Total deferrals in scope: 5 (0 new from M242 · 5 inherited from M237–M241, all already fated)
- New M242 deferrals: **0**
- Single deferrals: 3 (DEF-M240-01, DEF-M239-01/D12, M239-D13)
- Repeat/chronic patterns: 1 (the standing demo-stack test debt, M238-D5) — **already fated + named expiry**
- Blocking (unresolved) items: **0**

## Deferral Inventory

M242's own ledger (`decisions.md`, `progress.md`, `spec-notes.md`, code TODO scan):

- **No new deferrals.** All 3 UX sections (row-regroup / header tab / role-tinted avatar) + the spec extension
  + Delivers checked off; 2 harden passes closed 2 toothless tests + added a WCAG contrast pin. The only churn
  M242 made to existing tests is **Fate-1**: 3 stale `.stitle` → `.cttitle` title asserts updated (the exact
  container the tuple regroup moved) — not a deferral.
- Code TODO/FIXME/HACK scan of the M242-touched files (`cockpit.py`, `tests/test_cockpit.py`): **NONE**.
- **PRE-EXISTING, NOT M242-caused: the 6 red `demo-stack/tests/test_cockpit.py` tests** (recorded in M242
  `decisions.md` D7). Confirmed against M242's 3 touched surfaces (content-tab render / header / hero avatar):
  **none fall inside**. The 4 academy-link tests assert the per-hero `[Academy]` link removed 2026-07-15
  (`cockpit.py:717-721`); the 2 overlay-JS tests assert the 30s window + 3 try-blocks removed by M218/v2.3.1 —
  a **different** cockpit area. **0 new from M242.** These are a **subset of the already-fated standing
  demo-stack test debt** (M238-D5) — NOT a new M242 deferral.

Inherited, still-open v2.6 deferrals (unchanged from the M241 audit; carried by M237–M241):

```yaml
- id: M238-D5 (standing demo-stack test debt)
  item: "9 standing demo-stack test failures (6 academy/overlay test_cockpit + test_host_prereqs_m215 + test_purge + test_reap reap-17700), 0 real defects"
  origin_milestone: M238 (re-baselined; predates v2.6 — the v2.5/M236 set)
  first_deferred_on: 2026-07-21 (v2.6 re-fate; carried from v2.5)
  last_seen_in: state.md Standing backlog + M242 test_cockpit.py run (6 of the 9 surfaced)
  destination: "M244 (Fate-2 — 'M244 is the expiry point'; discharge by editing the tests; 6 of 9 need no live stack)"
  reason_recorded: "identical set re-baselined, 0 regressions; host-dependent; cheap test edits"
  partial_attempted: no

- id: DEF-M240-01
  item: "content-stories real-video exhibit (by-reference Bunny CDN render)"
  origin_milestone: M240
  first_deferred_on: 2026-07-22
  last_seen_in: roadmap.md M240 block + state.md
  destination: "M244 (Fate-3, USER PRE-APPROVED 2026-07-22 — land IF Bunny recording keys reachable on billion, else keep presence-only)"
  reason_recorded: "BUNNY_RECORDING_CDN_TOKEN_KEY + PULL_ZONE_HOST genuinely absent from this box's dev-stack"
  partial_attempted: no (voice presence-only IS the shipped deliverable)

- id: DEF-M239-01 / M239-D12
  item: "fail the BUILD loudly on ENOSPC (the 2nd F1 candidate)"
  origin_milestone: M239
  first_deferred_on: 2026-07-21
  last_seen_in: roadmap.md M239 block + state.md
  destination: "M244 (Fate-3)"
  reason_recorded: "F1 landed the VM-disk pre-flight; the loud-build-fail-on-ENOSPC half routed forward"
  partial_attempted: no

- id: M239-D13 (reap-17700)
  item: "9th demo-stack failure — test_a_RACED_listener_exits_silently hardcoded port 17700 test-isolation collision"
  origin_milestone: M239
  first_deferred_on: 2026-07-21
  last_seen_in: roadmap.md M239 block + state.md
  destination: "M244 (Fate-3, with a fix recipe)"
  reason_recorded: "root-caused to a test-isolation collision vs a live demo-1 cockpit; reap.sh itself correct"
  partial_attempted: no
```

(Note: `DEF-M235-03`/M204 assign-WRITE is owned by **M243**, an in-flight sibling of M242 — not M242's to
re-fate. It is fresh-dated KEEP-DEFERRED-WITH-SIGNOFF 2026-07-20 with M243 as its expiry.)

## Repeat-Deferral Patterns

### REPEAT (chronic): "standing demo-stack test debt" (M238-D5)
- **First deferred:** v2.5/M236 (re-baselined), carried into v2.6.
- **Re-confirmed:** M238 close (D5) · M239 close · M240 close · M241 close · **M242 close (this pass)** — the
  identical set each time, **0 regressions** at each close. M242 surfaced the same 6-of-9 subset (`test_cockpit.py`).
- **Current destination:** M244 (Fate-2), explicitly named as **the expiry point** — "M244 should discharge
  them by editing the tests (cheap; 6 of 9 need no live stack)."
- **Time in limbo:** ridden **≥5 v2.6 milestones** (M238→M242).
- **Classification:** the pattern IS chronic, BUT it carries a **resolution decision with a named expiry
  milestone** (M244) re-affirmed with fresh dates at every recent close. It is NOT an unresolved repeat, and it
  does NOT hit an AGED-OUT-unresolved trigger: the destination milestone (M244, terminal) has not yet closed,
  and M242's touched surfaces (content-tab render / header / avatar) do NOT include the academy-link/overlay-JS
  area these 6 tests cover (D7). Under the aging policy it is re-fated fresh (2026-07-22) → **YELLOW, not RED.**
  Flag for M244: the ride-count is now ≥5 — M244 is the correct and final expiry point.

## Fate-1 Investigation

### M238-D5 — the standing 6-of-9 that M242 surfaced
- **Fate-1 (land now in M242) feasible:** no.
- **Why:** these are ACADEMY-LINK + OVERLAY-JS test residuals (M218/M238 surfaces). M242 touches the cockpit
  content-tab render / header / hero-avatar surfaces only. Editing academy/overlay tests inside M242 is
  scope-creep into unrelated milestones' domains — exactly what the three-fate rule routes elsewhere. Fate-2:
  **M244 already owns them** (named expiry). Confirm, don't edit.

### DEF-M240-01, DEF-M239-01, M239-D13
- **Fate-1 (land now in M242) feasible:** no — each belongs to a different domain (media exhibit / build-disk
  pre-flight / demo-stack reap test), none in M242's render-layer scope. All already Fate-3 → M244 with fresh dates.

## Recommendations

| Item | Verdict | Home |
|---|---|---|
| M242's own ledger | (nothing to fate) | 0 new deferrals — clean complete close |
| M238-D5 standing debt (6-of-9 M242 surfaced) | **LAND-NEXT (Fate-2)** | M244 — already-owned, named expiry; no edit |
| DEF-M240-01 real-video exhibit | **LAND-NEXT (Fate-3)** | M244 — user pre-approved 2026-07-22 |
| DEF-M239-01 ENOSPC loud-build-fail | **LAND-NEXT (Fate-3)** | M244 |
| M239-D13 reap-17700 collision | **LAND-NEXT (Fate-3)** | M244 |

No LAND-NOW, no DROP, no new KEEP-DEFERRED-WITH-SIGNOFF required for M242.

## Applied Changes

None required. M242 introduced no new deferrals; every inherited item already carries a recorded fate with a
named expiry milestone (M244/M243) and fresh confirmation dates from the M238–M241 closes. The 6 pre-existing
`test_cockpit.py` failures are re-confirmed Fate-2 → M244 (subset of M238-D5) — recorded in M242's
`decisions.md` D7, no new decision needed.

## Blocking Items (require user decision)

**None.** No unresolved repeat deferral; the one chronic pattern (M238-D5) is already fated with a named expiry
(M244) and fresh re-confirmation. Verdict YELLOW → `SEVERITY=warning`, close proceeds.
