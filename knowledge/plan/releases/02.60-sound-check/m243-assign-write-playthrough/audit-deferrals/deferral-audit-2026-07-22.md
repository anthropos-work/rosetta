---
title: "Deferral Audit — milestone M243 (assign-WRITE Playthrough)"
date: 2026-07-22
scope: milestone
invoked-by: close-milestone
---

## Verdict
YELLOW

M243 introduced **zero new deferrals** and **RESOLVED one standing-backlog carry as Fate-1**: it LANDED
`DEF-M235-03` / M204 assign-WRITE — the ~10-routing item that was fresh-dated **KEEP-DEFERRED-WITH-SIGNOFF
2026-07-20 with M243 as its named expiry** ("if M243 does not land it, DROP"). M243 landed it in full
(`pt-assignment-assign`, live GREEN + DB-verified), so that item leaves the ledger — a deferral discharged,
the opposite of erosion. The remaining in-scope deferrals are all **inherited, unchanged, already fated
Fate-2/Fate-3 → M244** with M244 as the confirmed expiry, re-affirmed with fresh dates at the M238–M242
closes. No RED conditions; consistent with the M238–M242 close verdicts (all YELLOW).

## Summary
- Total deferrals in scope: 4 (0 new from M243 · 4 inherited from M237–M242, all already fated → M244)
- New M243 deferrals: **0**
- Deferrals RESOLVED by M243 (Fate-1): **1** (`DEF-M235-03` assign-WRITE — LANDED)
- Single deferrals: 3 (DEF-M240-01, DEF-M239-01/D12, M239-D13)
- Repeat/chronic patterns: 1 (the standing demo-stack test debt, M238-D5) — **already fated + named expiry**
- Blocking (unresolved) items: **0**

## Deferral Inventory

M243's own ledger (`decisions.md` D1–D6, `progress.md`, `spec-notes.md`, code TODO scan):

- **No new deferrals.** All 5 sections (UC1 fill · `/enterprise/assignments` page object · `pt-world`
  precondition · the live spec · Delivers) + all supporting items checked off; 1 harden pass closed the
  `isOnAssignments` predicate-parity gap and empirically mutation-verified the two Go honesty-teeth. D1–D6 are
  implementation-choice records (assign surface/flow, anti-toothless read-back, lockstep precondition, antd-v6
  Select lesson, corpus-pin reversal, dev-run roster refresh) — none is a defer note.
- Code TODO/FIXME/HACK scan of the M243-touched files (`assignments-page.ts`, `assignment-assign.spec.ts`,
  `url-shapes.ts`, `url-shapes.unit.spec.ts`, `seed-worlds.yaml`, `corpus_test.go`): **NONE**. The manifest
  `playthrough: TODO` M243 was chartered to fill is **filled** (`ptvalidate`: 16 live / 0 TODO).
- **The live re-drive on `billion` is NOT a deferral** — it is M244's declared scope by the roadmap
  (`overview.md` Out: "the re-prove-on-billion live drive (M244 executes it)"; M244 exit-gate (h) "every v2.6
  fix proven live"). Fate-2, already owned. The build already proved `pt-assignment-assign` live GREEN on
  demo-1 (7.9 s) + DB-verified (`organization_assignments` 6→7); M244 re-proves it cold on `billion` with the
  rest of the suite.

Inherited, still-open v2.6 deferrals (unchanged from the M242 audit; carried by M237–M242):

```yaml
- id: M238-D5 (standing demo-stack test debt)
  item: "9 standing demo-stack test failures (6 academy/overlay test_cockpit + test_host_prereqs_m215 + test_purge + test_reap reap-17700), 0 real defects"
  origin_milestone: M238 (re-baselined; predates v2.6 — the v2.5/M236 set)
  first_deferred_on: 2026-07-21 (v2.6 re-fate; carried from v2.5)
  last_seen_in: state.md Standing backlog + M242-close full-suite run (839 pass / 9 fail)
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

## Resolved This Milestone (Fate-1 — the ledger shrank)

```yaml
- id: DEF-M235-03 / M204 assign-WRITE in-manifest TODO
  item: "the assign-WRITE half of assignment-monitoring.assign-and-track (UC1) — a two-backend org-admin WRITE Playthrough"
  origin_milestone: M204 (v2.0); routed ~10 times across 5 releases
  status_at_this_audit: RESOLVED — LANDED by M243 (Fate-1)
  landed_as: "pt-assignment-assign (e2e/tests/assignment-assign.spec.ts); manifest UC1 non-TODO; corpus 15→16 live / 0 TODO"
  proof: "live GREEN on demo-1 (7.9s) + DB-verified organization_assignments skill_path/active 6→7 (build); Go+TS+ptvalidate green (close)"
  prior_fate: "KEEP-DEFERRED-WITH-SIGNOFF 2026-07-20, expiry 'if M243 does not land it, DROP'"
```

## Repeat-Deferral Patterns

### REPEAT (chronic): "standing demo-stack test debt" (M238-D5)
- **First deferred:** v2.5/M236 (re-baselined), carried into v2.6.
- **Re-confirmed:** M238 (D5) · M239 · M240 · M241 · M242 closes — the identical set each time, **0 regressions**.
- **Current destination:** M244 (Fate-2), explicitly named as **the expiry point** — "M244 should discharge
  them by editing the tests (cheap; 6 of 9 need no live stack)."
- **Time in limbo:** ridden **≥5 v2.6 milestones** (M238→M242); M243 is a docs-only-corpus close that touched
  **no** demo-stack test (all M243 code is in `rext playthroughs/`), so it surfaces none of the 9 and adds none.
- **Classification:** the pattern IS chronic, BUT it carries a **resolution decision with a named expiry
  milestone** (M244, the terminal milestone — the discharge point is now the immediately-next milestone),
  re-affirmed with fresh dates at every recent close. It is NOT an unresolved repeat and does NOT hit an
  AGED-OUT-unresolved trigger: the destination milestone (M244) has not yet closed, and M243 did not touch the
  demo-stack test area. Under the aging policy it is re-fated fresh (2026-07-22) → **YELLOW, not RED.** Flag for
  M244: the ride-count is ≥5 — M244 is the correct and final expiry point (discharge or DROP).

## Fate-1 Investigation

### M238-D5 — the 9 standing demo-stack failures
- **Fate-1 (land now in M243) feasible:** no.
- **Why:** these are ACADEMY-LINK + OVERLAY-JS + host-prereq/purge/reap residuals in the `demo-stack` Python
  suite. **M243 touched none of that** — every M243 code change is in the `rext playthroughs/` section (Go
  manifest + TS e2e), and the corpus side is docs-only. Editing demo-stack tests inside M243 is scope-creep into
  unrelated milestones' domains — exactly what the three-fate rule routes elsewhere. Fate-2: **M244 already owns
  them** (named, now-adjacent expiry). Confirm, don't edit.

### DEF-M240-01, DEF-M239-01, M239-D13
- **Fate-1 (land now in M243) feasible:** no — each belongs to a different domain (media exhibit / build-disk
  pre-flight / demo-stack reap test), none in M243's Playthroughs scope. All already Fate-3 → M244 with fresh dates.

## Recommendations

| Item | Verdict | Home |
|---|---|---|
| M243's own ledger | (nothing to fate) | 0 new deferrals — clean complete section close |
| DEF-M235-03 assign-WRITE | **LAND-NOW (Fate-1) — DONE** | landed by M243 as `pt-assignment-assign`; leaves the backlog |
| M238-D5 standing debt (9 fails) | **LAND-NEXT (Fate-2)** | M244 — already-owned, named (adjacent) expiry; no edit |
| DEF-M240-01 real-video exhibit | **LAND-NEXT (Fate-3)** | M244 — user pre-approved 2026-07-22 |
| DEF-M239-01 ENOSPC loud-build-fail | **LAND-NEXT (Fate-3)** | M244 |
| M239-D13 reap-17700 collision | **LAND-NEXT (Fate-3)** | M244 |

No new KEEP-DEFERRED-WITH-SIGNOFF required for M243. One standing KEEP-DEFERRED item (DEF-M235-03) converted to
**LAND-NOW/Fate-1 and discharged** this milestone.

## Applied Changes

None required to the plan. M243 introduced no new deferrals; every remaining inherited item already carries a
recorded fate with a named expiry milestone (M244) and fresh confirmation dates from the M238–M242 closes.
`DEF-M235-03`'s discharge is recorded in the roadmap/state Standing-backlog update at close (Phase 10) — the
item's expiry condition ("land it or DROP") is met by LANDING it.

## Blocking Items (require user decision)

**None.** No unresolved repeat deferral; the one chronic pattern (M238-D5) is already fated with a named,
now-adjacent expiry (M244) and fresh re-confirmation, and the standing item M243 was chartered to resolve is
resolved. Verdict YELLOW → `SEVERITY=warning`, close proceeds.
