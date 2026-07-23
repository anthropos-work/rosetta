---
title: "Deferral Audit — M244 prove-on-billion (milestone close)"
date: 2026-07-23
scope: milestone
invoked-by: close-milestone
---

## Verdict
YELLOW

M244 is the **iterative terminal closer** of v2.6 "sound check" and a **proof milestone** (`delivers: none`,
0 platform edits). It introduced **zero new deferrals**, **RESOLVED two inherited carries** (reap-17700
LANDED as Fate-1; DEF-M240-01 dispositioned), and **reconciled the one M244-introduced test-fence** (the M215
no-pipe form) during the final harden. The two remaining open items are **inherited, pre-M244, unchanged**,
and their named expiry — first written as "M244" across the M238–M243 closes — resolves, now that M244 is
closing proof-only, to the **release-level deferral audit** (`/developer-kit:close-release` Phase 1b, the
designed final expiry, now one step away). Both are re-fated **LAND-NEXT (Fate-3) → close-release** with a
clear note. No RED conditions raised to a blocker: M244 neither introduced these nor is scoped to fix them
(fixing 6 academy/overlay cockpit tests + a docker-integration purge test + an ENOSPC build-path change is
platform/tooling-fix work, structurally not "prove-on-billion" work), and the terminal release audit is the
correct, adjacent home. This is an explicit re-fate, not a silent auto-accept.

## Summary
- Total deferrals in scope: **4** (2 resolved this milestone · 2 open, inherited)
- Single deferrals: 0 net-new
- Repeat deferrals: **2** (standing demo-stack test debt [chronic] · DEF-M239-01 [aged])
- Chronic patterns flagged: **1** (M238-D5 standing demo-stack test debt — ridden ≥5 v2.6 milestones)

## Deferral Inventory

```yaml
- id: reap-17700 (M239-D13)
  item: "test_a_RACED_listener_exits_silently hardcodes cockpit port 17700 → collides with a live demo-1"
  origin_milestone: M239
  destination: M244
  status: RESOLVED (Fate-1, iter-10) — _reap_with_stubs now defaults to a guaranteed-free port; proven
    load-bearing with :17700 held; test_reap.py 41/41. rext b38ad75 (tag 7dbad4b). LEAVES the ledger.

- id: DEF-M240-01
  item: "content-stories real-video (voice) exhibit — re-pin a hiring-voice cell to a recorded session +
    provision Bunny.net signing keys + exhibit-by-reference render"
  origin_milestone: M240
  destination: M244 (Fate-3, user pre-approved 2026-07-22 — 'land it live IF the keys are reachable on
    billion, else keep voice presence-only')
  status: DISPOSITIONED (iter-07) — Bunny keys ABSENT on billion → the 2 voice player cells (hire-voice-fail,
    asmt-voice-pass-en) dispositioned to PLAYER-presence-only per the pre-approved fallback; manager view
    retained (not a drop); denominator 49→47 landing. The pre-approved conditional resolved by its else-branch.
    LEAVES the ledger.

- id: M238-D5 (standing demo-stack test debt)
  item: "Standing demo-stack python test failures — pre-M244 inherited, SUT provably untouched by M244"
  origin_milestone: pre-M244 (rext 04babf8, 2026-07-15); carried M238→M244
  first_deferred_on: 2026-07-15 (re-confirmed each M238–M243 close)
  last_seen_in: state.md "Standing backlog" (RE-BASELINED); rebaseline-standing-failures.md
  destination: M244 (named expiry) → RE-FATE → close-release
  reason_recorded: "0 real defects, 0 pin drift, 0 new; host-dependent; discharge by editing the tests"
  partial_attempted: no (M244 final harden fixed the ONE M244-introduced fence [M215 no-pipe], 9→8)

- id: DEF-M239-01
  item: "make the demo build FAIL LOUDLY on ENOSPC (the disk-full class that surfaced as a cryptic
    `redis exited(1)`; M239 fixed the pre-flight disk-measure, this is the build-time loud-abort follow-on)"
  origin_milestone: M239 (D12)
  first_deferred_on: 2026-07-21
  last_seen_in: overview.md In: (M244 scope); progress.md "Remaining open carry"
  destination: M244 → RE-FATE → close-release
  reason_recorded: "build-path change needing a real ENOSPC to validate; deliberately not landed on a clean
    gate-met close"
  partial_attempted: no
```

## Repeat-Deferral Patterns

### CHRONIC: "Standing demo-stack test debt" (M238-D5)
- **First deferred:** pre-M244 (rext 04babf8, 2026-07-15), re-baselined at v2.5 close.
- **Deferred again:** each of M238 · M239 · M240 · M241 · M242 · M243 closes (all Fate-2 → M244, fresh-dated).
- **Current count:** **8** (was 9). 6 `test_cockpit` academy/overlay + `test_public_host` (port-13001) +
  `test_purge` (docker-integration). The 9th (M215 no-pipe fence) was **M244-introduced by the final harden's
  iter-25 describe-form change and reconciled inline in Pass 8** (Fate-1) — so the standing set is now **8,
  all provably pre-M244** (SUT untouched by M244; the final harden characterized each).
- **Time in limbo:** ridden ≥5 v2.6 milestones; named expiry has always been "M244."
- **Pattern:** `CHRONIC_DEFER` (reason is stable — "0 real defects, host-dependent, discharge by editing").

### AGED: "DEF-M239-01 ENOSPC loud-build-fail"
- **First deferred:** M239, 2026-07-21 (Fate-3 → M244).
- **Ageing trigger:** destination milestone (M244) closing without landing it.
- **Reason unchanged:** needs a real ENOSPC to validate a build-path change; not landable on a clean gate-met
  proof close.

## Fate-1 Investigation

### M238-D5 — "Standing demo-stack test debt (8)"
- **Fate-1 (land now, complete) feasible:** no.
- **Why:** M244 is a **proof milestone** — its charter is `0 platform edits` and its work is *driving + asserting*
  a cold billion bring-up, NOT authoring demo-stack test fixes. Landing these requires editing 6 academy/overlay
  cockpit tests + a docker-integration purge test — tooling-fix work outside M244's scope, and 2 of the 8 need a
  live stack / docker-integration harness. The final harden proved the SUT was untouched by M244 and fixed the
  only M244-introduced fence. Landing 8 unrelated inherited test fixes inside a proof close would be scope creep,
  not Fate-1 discipline.
- **Fate:** **Fate-3 → close-release.** The release-level deferral audit (close-release Phase 1b, release scope,
  "the last chance to pull them forward before they escape to backlog," extra scrutiny) is the genuine final
  expiry for release-wide test debt. It is the very next step after this milestone close.

### DEF-M239-01 — "ENOSPC loud-build-fail"
- **Fate-1 feasible:** no.
- **Why:** it is a **build-path change** (make the demo image build abort loudly when the Docker VM runs out of
  disk) that can only be *validated* by inducing a real ENOSPC on a build box — not reproducible or provable
  inside a clean gate-met proof close on a healthy billion. Landing code you cannot validate this pass is exactly
  the disguised-partial the three-fate rule rejects.
- **Fate:** **Fate-3 → close-release** (or a follow-on v2.x tooling milestone the release close designates).

## Recommendations
1. **M238-D5 standing demo-stack test debt (8)** → **LAND-NEXT (Fate-3) → `/developer-kit:close-release`.**
2. **DEF-M239-01 ENOSPC loud-build-fail** → **LAND-NEXT (Fate-3) → `/developer-kit:close-release`.**

Both re-fated forward to the terminal release-level deferral audit — the designed final expiry — with the
explicit note that M244 (proof-only, 0 platform edits) neither introduced nor was scoped to fix them.

## Applied Changes
- Recorded D1 (Deferral re-audit) in `../decisions.md` re-fating both items LAND-NEXT (Fate-3) → close-release.
- This report filed under `audit-deferrals/`.
- No plan edits to a sibling milestone (M244 is terminal; the forward target is the release close, not a
  milestone `overview.md`); state.md's "Standing backlog" already carries both with a fated destination and
  is updated at Phase 10 to point the expiry at the release close.

## Blocking Items (require user decision)

**None raised to blocker.** The two repeat/aged items are pre-M244 inherited, structurally un-landable inside a
proof-only milestone, and have a clean, adjacent forward home (the terminal release-close deferral audit — the
designed final expiry). This is an explicit Fate-3 re-fate with notes, not a silent auto-accept, and is the
disposition the milestone/release design prescribes. Verdict **YELLOW → `SEVERITY=warning`**; close proceeds.
The chronic pattern (M238-D5) is flagged for close-release's extra-scrutiny audit so it does not escape to
backlog un-fated.
