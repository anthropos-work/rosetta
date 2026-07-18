---
title: "Deferral Audit — M226 opening-night (close)"
date: 2026-07-17
scope: milestone
invoked-by: close-milestone
---

## Verdict
**YELLOW** — one NEW single deferral (Finding-3, self-resolving, non-gate-blocking) + the two INHERITED carries
already re-fated fresh at the M225 close **today** (2026-07-17) and routed to the v2.4 **release** close. No
repeat-deferral of *promised milestone work* without a decision; no chronic/silent scope erosion. M226's own gate
work landed complete (7/7 gate MET, 2 cold cycles + orchestrator re-verify).

## Summary
- Total deferrals in scope: **3** (1 NEW M226-owned + 2 INHERITED carries)
- Single deferrals: 3
- Repeat deferrals (of promised feature work): 0
- Chronic patterns flagged: 0

## Deferral Inventory

### M226's own ledger — CLEAN except Finding-3
- Every gate condition (C1–C7) landed **Fate-1**: proven GREEN on `billion` over the tailnet, reproducibly across
  2 clean cold reset-to-seed cycles + independently orchestrator-re-verified from this Mac (C2 5/5 sims × 42
  candidates junk=0; C5 recruiter p95 1.74 s < 5 s). 5 cross-machine findings (F1 serve-hiring, F2 count 5+45, F3
  surgical-orphan, F4 C2 harness race, F5 C2 cold/tailnet probe-budget) were surfaced + fixed live — all
  tooling/harness/seed, **0 platform edits**. R4 (compare-drawer cold render) documented in `latency-budget.md` as
  a warm-up transient, **not** a gate-condition violation — recorded, not deferred.
- `overview.md` `Out:` = "nothing net-new; this is the acceptance closer" — no deferred scope.

### DEF-M226-01 — Finding-3, the pre-bind serve reap (NEW this milestone)
```yaml
- id: DEF-M226-01
  item: "pre-bind `tailscale serve` reap — clear stale serve fronts on the demo's offset ports before bind (closes the M215 F12 out-of-band-serve window)"
  origin_milestone: M226 (surfaced iter-04)
  first_deferred_on: 2026-07-17
  last_seen_in: m226 hardening-ledger.md:69-77 (Fate 3 — routed forward, NOT gate-blocking)
  destination: "a follow-up build-iter with a live re-prove, or the next prove-on-<VM> milestone"
  reason_recorded: "bring-up-path behavioral change on a live-only surface (no tailscaled in the build env → not deterministically testable here) needing a live re-prove on billion — which the harden constraint forbids (demo is UP; no re-bring-up). SELF-RESOLVES in the default flow (teardown already emits the reset)."
  partial_attempted: no
```

### DEF-CARRY-A — 8 pre-existing demo-stack test failures (inherited, re-fated TODAY at M225)
```yaml
- id: DEF-CARRY-A
  item: "8 pre-existing demo-stack test failures (6× test_cockpit.py [4 removed-academy-CTA + 2 v2.3.1 overlay-JS] + test_purge + test_reap)"
  origin_milestone: pre-v2.4 (v2.3.1/v2.3.2 cockpit hotfixes); first carried explicitly at M224 close (D6)
  first_deferred_on: 2026-07-16 (M224 D6)
  last_seen_in: M225 audit-deferrals/deferral-audit-2026-07-17-m225-close.md (fresh re-fate 2026-07-17); HEAD-identical in rext demo-stack tests
  destination: "v2.4 RELEASE close (Phase 1b, release scope + extra scrutiny) → a future demo-stack test-debt harden pass"
  reason_recorded: "HEAD-identical; in files M224/M225/M226 never touched; predates v2.4; outside the hiring domain — fixing = scope-bleed"
  partial_attempted: no
```

### DEF-CARRY-B — the M204 assign-WRITE declared TODO (inherited, re-fated TODAY at M225)
```yaml
- id: DEF-CARRY-B
  item: "assignment-monitoring.assign-and-track.UC1 — the assign-WRITE half (two-backend org-admin WRITE flow)"
  origin_milestone: M204 (v2.0 opening night)
  first_deferred_on: 2026-07-02 (M204, declared in-manifest TODO)
  last_seen_in: M225 audit (fresh re-fate 2026-07-17); corpus/ops/demo/playthroughs.md; rext playthroughs manifest (reports `unimplemented`)
  destination: "v2.4 RELEASE close (its declared-TODO fate)"
  reason_recorded: "declared, manifest-tracked build-reference gap (reports `unimplemented`); not silent erosion; M226 touched no playthroughs-manifest surface"
  partial_attempted: no
```

## Repeat-Deferral Patterns
**None of the blocking kind.**
- **DEF-M226-01 (Finding-3)** is NEW this milestone — a single, fresh deferral. Not a repeat.
- **DEF-CARRY-A** has now been carried M224→M225→M226 (trips the aging trigger), but it is an **inherited-failure
  carry** in untouched, unrelated files (the `CHRONIC_DEFER of a wanted feature` pattern the gate blocks on does
  not apply). It has a conscious decision (M224 D6), a destination, and was **freshly re-fated TODAY** (2026-07-17)
  at the M225 close with standing user sign-off. M226 touched **no** demo-stack test file.
- **DEF-CARRY-B** is a **declared in-manifest TODO** (reports `unimplemented`) carried since v2.0 — explicit,
  tracked, surfaced in every count. Its designated re-fate point is the **release** close, now imminent.

## Fate-1 Investigation

### DEF-M226-01 (Finding-3)
- **Fate-1 (land now, complete) feasible:** no.
- **Why:** wiring a pre-bind `--reset` into `up-injected.sh` is a bring-up-path behavioral change on a **live-only**
  surface — no `tailscaled` exists in the build/test env, so it is not deterministically testable here, and proving
  it requires a **live re-prove on `billion`**, which the close constraint forbids (the billion demo is intentionally
  LEFT UP as the live-proof artifact; no re-bring-up). The inverse machinery (`reset_commands`/`reset_script`) already
  exists and is now **unit-fenced** (incl. hiring). The gap **self-resolves in the default flow** (teardown already
  emits the reset).
- **Fate:** Fate 3 — routed forward to a follow-up build-iter with a live re-prove, or the next `prove-on-<VM>`
  milestone. Non-gate-blocking; recorded in the M226 Gate Outcome Ledger.

### DEF-CARRY-A
- **Fate-1 feasible:** no. The failures live in the rext **demo-stack** section (`test_cockpit`/purge/reap), a domain
  M226 never touched (M226 = live acceptance proof; code touched only tailscale-serve + seed count + e2e harness).
  Fixing them at an M226 close is scope-bleed into an unrelated module.
- **Fate:** KEEP-DEFERRED (carry) — re-confirmed today; the v2.4 **release close** is the mandated re-fate gate.

### DEF-CARRY-B
- **Fate-1 feasible:** no (out of M226's scope entirely). The assign-WRITE half is a two-backend org-admin WRITE flow,
  a declared build-reference gap owned by the Playthroughs manifest.
- **Fate:** KEEP-DEFERRED (declared TODO) — its explicit fate belongs to the v2.4 release close.

## Recommendations
- **DEF-M226-01 → LAND-NEXT / KEEP-DEFERRED (Fate 3, follow-up).** Fresh decision dated **2026-07-17**: a live-only
  bring-up-path change not testable in the close env; self-resolves in the default flow; not gate-blocking. Route to a
  follow-up build-iter (live re-prove) or the next `prove-on-<VM>` milestone. Recorded in the Gate Outcome Ledger.
- **DEF-CARRY-A → KEEP-DEFERRED-WITH-SIGNOFF (carry).** Re-confirmed **2026-07-17**: unchanged (M226 touched no
  demo-stack test file); destination **v2.4 release close** → future demo-stack test-debt harden pass.
- **DEF-CARRY-B → KEEP-DEFERRED (declared TODO).** Re-confirmed **2026-07-17**: unchanged; its declared-TODO fate is
  a v2.4 **release-close** decision (close-release Phase 1b — the last-chance gate).

## Applied Changes
- This report authored (fresh dated fate for the new Finding-3 + re-confirmation of both inherited carries).
- No plan mutations, no code fixes, no roadmap edits: Finding-3 has a durable home in the M226 hardening-ledger +
  Gate Outcome Ledger; carries A/B already have durable homes (M224 D6 + state.md standing backlog for A; the
  Playthroughs manifest + `playthroughs.md` for B). All three are inherited by the v2.4 close-release Phase 1b audit
  with `release` scope + the "extra scrutiny" mandate.

## Blocking Items (require user decision)
**None.** M226's own gate work landed complete (7/7 MET); Finding-3 is a single, fresh, self-resolving,
non-gate-blocking deferral with Fate 3; both inherited carries have conscious decisions, destinations, and standing
sign-off, re-confirmed this pass and routed to the imminent v2.4 release close. Verdict **YELLOW** →
`SEVERITY=warning`; the milestone close proceeds.
