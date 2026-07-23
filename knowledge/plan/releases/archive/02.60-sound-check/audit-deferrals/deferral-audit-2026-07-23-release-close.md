---
title: "Deferral Audit — v2.6 «sound check» (release close)"
date: 2026-07-23
scope: release
invoked-by: close-release
release: v2.6
branch: release/02.60-sound-check
milestones: [M237, M238, M239, M240, M241, M242, M243, M244]
---

## Verdict
RED

RED because the release carries **1 CHRONIC repeat-deferral pattern** (the standing demo-stack test debt,
ridden ≥5 v2.6 milestones) **+ 2 AGED_OUT items** (both had a named destination — M244 — that has now closed
without landing them). Per the aging policy, the prior deferral authority is **revoked**: both items require a
**fresh user fate-decision this pass**. This is not a "something broke" RED — there are **0 real defects, 0
platform edits, 0 pin drift, flake 0** — it is the designed terminal gate finally firing on debt that was
explicitly routed here. This is the last chance to pull these forward before they escape to backlog un-fated.

## Summary
- Total deferrals in scope (whole release): **4**
- Resolved during the release (leave the ledger): **2** — reap-17700 (Fate-1 LANDED, M244 iter-10);
  DEF-M240-01 real-video exhibit (dispositioned to player-presence-only, M244 iter-07)
- Open at close: **2**
- Single deferrals: 0
- Repeat deferrals: **2** (standing demo-stack test debt · DEF-M239-01)
- Chronic patterns flagged: **1** (standing demo-stack test debt)
- AGED_OUT (blocking, require fresh fate): **2**

Also confirmed landed this release (not open): **DEF-M235-03 / M204 assign-WRITE** (Fate-1 LANDED at M243 —
the sole in-manifest Playthrough TODO filled). No `RELEASE-SCOPE-DEFER:` escape-hatch decisions and no
`carry-forward.md` exist in the release — every routed item pointed at this close, not a future release.

## Deferral Inventory

```yaml
# --- RESOLVED during v2.6 (recorded for audit trail; leave the ledger) ---
- id: reap-17700 (M239-D13)
  item: "test_reap hardcoded cockpit port 17700 → collides with a live demo-1 cockpit"
  origin_milestone: M239
  status: RESOLVED — Fate-1 LANDED at M244 iter-10 (free-port default; test_reap 41/41). LEAVES ledger.
- id: DEF-M240-01
  item: "content-stories real-video (voice) exhibit — Bunny.net signed exhibit-by-reference"
  origin_milestone: M240
  status: DISPOSITIONED — Bunny keys absent on billion → 2 voice player cells to player-presence-only per the
    user-pre-approved (2026-07-22) fallback; manager view retained; denominator 49→47. LEAVES ledger.

# --- OPEN at release close (require fresh fate) ---
- id: STANDING-DEMOSTACK-TESTDEBT (M238-D5)
  item: "8 standing demo-stack python test failures — 0 real defects, host-dependent, discharge by editing tests"
  composition: "6 test_cockpit academy/overlay (stale assertions vs deliberately-changed M218/M238 behaviour +
    academy live-premise) · test_public_host (port-13001 hiring-app expectation added by M220, SUT untouched) ·
    test_purge (docker-integration, needs a live docker box)"
  origin_milestone: pre-M244 (rext 04babf8, 2026-07-15); carried M238→M244, re-baselined at v2.5 close
  first_deferred_on: 2026-07-15
  destinations_over_time: "M238→M239→M240→M241→M242→M243→M244 (all Fate-2, fresh-dated), then M244→close-release"
  reason_recorded: "0 real defects, 0 pin drift, host-dependent; discharge by editing the tests"
  partial_attempted: yes — M244 final harden fixed the ONE M244-introduced fence (M215 no-pipe), 9→8
  flags: [REPEAT, CHRONIC_DEFER, AGED_OUT]
- id: DEF-M239-01
  item: "make the demo build FAIL LOUDLY on ENOSPC (build-time loud-abort follow-on to M239's pre-flight disk-measure)"
  origin_milestone: M239 (D12)
  first_deferred_on: 2026-07-21
  destinations_over_time: "M239→M244 (Fate-3), then M244→close-release (Fate-3)"
  reason_recorded: "build-path change needing a real ENOSPC to validate; not landable on a clean gate-met close"
  partial_attempted: no
  flags: [REPEAT, AGED_OUT]
```

## Repeat-Deferral Patterns

### CHRONIC: "Standing demo-stack test debt" (M238-D5)
- **First deferred:** pre-M244 (rext 04babf8, 2026-07-15); re-baselined at v2.5 close.
- **Deferred again:** each of M238·M239·M240·M241·M242·M243·M244 closes (all fresh-dated Fate-2/3).
- **Current count:** **8** (6 test_cockpit academy/overlay + test_public_host port-13001 + test_purge docker).
- **Time in limbo:** ridden ≥5 v2.6 milestones (+ carried from v2.5); named expiry has always been "the release close" = **now**.
- **Pattern:** `CHRONIC_DEFER` — reason stable across every defer ("0 real defects, host-dependent, discharge by editing").

### AGED: "DEF-M239-01 ENOSPC loud-build-fail"
- **First deferred:** M239, 2026-07-21 (Fate-3 → M244).
- **Ageing trigger:** destination milestone (M244) closed without landing it.
- **Reason unchanged:** needs a real ENOSPC to validate a build-path change.

## Fate-1 Investigation

### STANDING-DEMOSTACK-TESTDEBT (8) — SPLIT is possible
- **Fate-1 (land now, complete) feasible:** PARTIAL — yes for the mechanical majority, no for a host/docker-gated minority.
- **Landable-now subset (~6):** the 6 `test_cockpit` academy/overlay failures + the `test_public_host` port-13001
  reconcile are **pure test-assertion updates** — re-point stale assertions at the deliberately-changed
  M218/M238/M220 behaviour. This is **rext tooling test-edit work, 0 platform edits, mechanical.** It fits a
  late-merge branch cleanly and finally discharges the chronic bulk. The context that made it "not proof-milestone
  scope" (M244's `0 platform edits` proof charter) no longer binds at release close.
- **Not-landable-now subset (~2):** `test_purge` (docker-integration) and any live-serve-gated assertion can only
  be **validated** on a live docker box / running demo — un-provable in a docs-only close. Landing unvalidatable
  edits is the disguised-partial the three-fate rule rejects.
- **Enumerated failure last time:** every prior defer cited M244's proof charter. That reason expires at the release close.

### DEF-M239-01 — "ENOSPC loud-build-fail"
- **Fate-1 feasible:** no. It is a **build-path change** validatable only by inducing a real ENOSPC on a build box —
  not reproducible in a clean docs-only close. Note the disk-full CLASS is **already defended**: M239 landed the
  pre-flight disk-**measure** (the demo aborts before build when the VM is low). This item is the redundant
  belt-and-braces loud-abort follow-on, of marginal residual value.

## Recommendations
> These are AGED_OUT / CHRONIC → the prior authority is revoked; each REQUIRES a fresh user fate-decision at this
> close (LAND-NOW / DROP / KEEP-DEFERRED-WITH-SIGNOFF). The options below are teed up with a recommendation; the
> orchestrator surfaces them to the user for the actual call.

1. **STANDING-DEMOSTACK-TESTDEBT (8)** → **RECOMMEND: LAND-NOW the mechanical subset on a late-merge branch +
   KEEP-DEFERRED-WITH-SIGNOFF the host/docker-gated residue → v2.7.**
   - LAND-NOW (late-merge branch, 0 platform edits): the 6 `test_cockpit` stale-assertion updates + the
     `test_public_host` port-13001 reconcile. Discharges the chronic bulk instead of re-deferring the whole
     bundle a 6th time — the disciplined answer to a `CHRONIC_DEFER`.
   - KEEP-DEFERRED-WITH-SIGNOFF → **v2.7 (a NAMED milestone, not "standing backlog")**: only `test_purge`
     (docker-integration) + any live-serve-gated test, because they need a live docker box to validate.
   - **Fallback if no late-merge branch is opened:** KEEP-DEFERRED-WITH-SIGNOFF the whole set → **v2.7 named
     milestone** (the v2.5-close rule: a fate needs a MILESTONE, not a sweep/next-close/standing-backlog). Do NOT
     re-defer to an unnamed "standing backlog" — that is the exact anti-pattern the aging policy forbids.

2. **DEF-M239-01 ENOSPC loud-build-fail** → **RECOMMEND: DROP** (with a decision-record rationale), fallback
   **KEEP-DEFERRED-WITH-SIGNOFF → v2.7.**
   - DROP rationale: the disk-full class is **already covered** by M239's pre-flight disk-measure; the loud-build-fail
     is redundant belt-and-braces of marginal value, and it is un-validatable without a real ENOSPC. Cutting it
     retires an aged item honestly rather than parking un-validatable code.
   - KEEP-DEFERRED-WITH-SIGNOFF → v2.7 is the acceptable thorough alternative if the belt-and-braces is wanted —
     with a fresh dated reason (needs an ENOSPC validation box) and a NAMED v2.7 milestone.

## Applied Changes
- This release-level report filed under `audit-deferrals/` (release root).
- No plan edits applied by this audit: both items require a fresh user fate-decision first (RED → orchestrator
  surfaces to the user). The chosen fates get recorded by the orchestrator at close (LAND-NOW → a late-merge
  branch + task; DROP → a `D#` decision + roadmap update; KEEP-DEFERRED-WITH-SIGNOFF → a `RELEASE-SCOPE-DEFER:`
  decision + `roadmap-vision.md` entry under the named v2.7 milestone).
- `knowledge/plan/state.md` "Standing backlog" already carries both with the close-release destination; it is
  updated to the user's chosen fate at Phase 10.

## Blocking Items (require user decision)
Both items below are AGED_OUT (destination M244 closed without landing) and one is CHRONIC — the prior deferral
authority is revoked. Each needs an explicit fresh verb this pass:

1. **STANDING-DEMOSTACK-TESTDEBT (8 failures)** — CHRONIC + AGED_OUT.
   Recommended: LAND-NOW the ~6 mechanical test-assertion fixes (late-merge branch) + KEEP-DEFERRED-WITH-SIGNOFF
   the ~2 docker/live-gated → v2.7 named milestone.
2. **DEF-M239-01 (ENOSPC loud-build-fail)** — AGED_OUT.
   Recommended: DROP (disk-full class already covered by M239 pre-flight disk-measure), fallback
   KEEP-DEFERRED-WITH-SIGNOFF → v2.7 named milestone.

`SEVERITY=blocker` until each has a recorded fresh fate. No code/tag change is gated on the fixes themselves —
the release is otherwise GREEN (0 real defects, 0 platform edits, flake 0) — but the tag should not land before
these two fates are recorded, per the terminal-audit contract.
