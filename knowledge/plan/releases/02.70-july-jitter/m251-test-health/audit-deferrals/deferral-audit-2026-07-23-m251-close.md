---
title: "Deferral Audit — M251 (close)"
date: 2026-07-23
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals — every item is a first-pass routing dated today (2026-07-23), across 0 prior milestones.
- No aged-out items — nothing ≥ 3 months old; no destination milestone (M247, M254) has closed.
- No chronic / drift patterns.
- Every item has a clear Fate-2 decision to an already-planned in-release milestone (M247 or M254). No
  Fate-1 landing owed in M251; no escape-hatch (cross-release) deferral; no sibling-plan edit.

## Summary
- Total deferrals in scope: 6 (2 M251-own + 4 inherited-from-M246 that named M251 as a possible dest)
- Single deferrals: 6
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0
- Blocking items: 0

## Deferral Inventory

M251 delivered both declared sections as Fate 1 (run-unit roster fix + the 6 mechanical `test_cockpit` /
`test_public_host` re-points — all checked off, all tests green). The forward-routed items are:

| id | item | origin | first_deferred | destination | fate |
|---|---|---|---|---|---|
| DEF-M251-01 | The 8 live/env/docker-gated demo-stack failures (`test_purge` + `test_ant_academy` launcher/reap ×5 + `test_host_isolation` mutant-term-trap + `test_ant_academy_clerk_wiring` overlay) — fail only on a stackless box | M251 | 2026-07-23 | M254 (gate g+h) | Fate 2 |
| DEF-M251-02 | Optional `corpus/ops/verification.md` demo-stack-suite + run-unit-roster index anchor (overview "Delivers →") | M251 | 2026-07-23 | M247 (owns the corpus reground / `verification.md`) | Fate 2 |
| DEF-M246-03 (inherited) | `test_injection.py` pins skillpath-as-injected + models already-merged skiller | M246 | 2026-07-23 | M247 (drift-ledger triage) | Fate 2 |
| DEF-M246-04 (inherited) | `exposure_claim_guard.py` `_cfg` lists skillpath:8095 | M246 | 2026-07-23 | M247 (drift-ledger triage) | Fate 2 |
| DEF-M246-08 (inherited) | fake-FAPI http-vs-TLS cheap-win probe artifact + end-to-end login not re-run | M246 | 2026-07-23 | M254 gate (h) | Fate 2 |
| DEF-M246-09 (inherited) | AI Academy peripheral (2/6 env keys short) not serving | M246 | 2026-07-23 | M254 / standing | Fate 2 |

## Repeat-Deferral Patterns
None. M251 is a fresh milestone in v2.7; every record is a first-pass routing dated today. No item has been
deferred across ≥ 2 milestones, none is ≥ 3 months old, neither destination milestone (M247, M254) has
closed, and no later milestone has yet touched these areas. Zero REPEAT / CHRONIC_DEFER / DRIFT_DEFER.

## Fate-1 Investigation
**Fate-1 (land in M251) is correctly NO for all 6 rows — by scope boundary, not by time pressure.**

- **DEF-M251-01 (8 live-gated).** They require a live/docker box to even reproduce (`test_purge` needs a
  real container-owned 0700 dir; the academy launcher/reap need live process-lifecycle; the clerk-wiring
  overlay sources a stack env). M251 built + gated on a stackless local box, so a Fate-1 landing here would
  be unverifiable. M254 (the live billion prove) owns "the docker/live-gated test-health tests green" (gate
  g) + "the live-browser specs green" (gate h). **Correction flagged to the M254 driver:** the real
  live-gated failing set is **8**, not the "~2" M254's overview estimates — see Recommendations.
- **DEF-M251-02 (verification.md anchor).** Optional deliverable. Authoring it in M251 would collide with
  M247's concurrent lane on the same ops doc (`verification.md`), which M247 owns as part of the corpus
  reground. Not a blind area — the code it would index exists + is exercised (proven green this close).
- **DEF-M246-03 / -04 (inherited, injection test fixtures).** Out of M251's design-confirmed scope; they
  live in `stack-injection/tests/` — a module M251 does not own (M251 owns only demo-stack *health/inventory*
  tests). M246's own audit routes them through the durable `drift-ledger.md` handoff M247 consumes. Inert
  (no compose service matches → the barrier is already green). Pulling them into M251 would also force a
  re-tag of the already-pushed code-of-record (`july-jitter-m251-test-health` @ e9e29a1) — a scope-creep at
  close time the three-fate rule rejects.
- **DEF-M246-08 (inherited, fake-FAPI probe).** The http-vs-TLS cheap-win probe-fix is only *verifiable*
  against a live fake-FAPI container; the end-to-end login re-prove is explicitly M254 gate (h). M251 lacks
  a live box, so this correctly rides M254.
- **DEF-M246-09 (inherited, academy peripheral).** Benign, non-fatal-by-design peripheral; the academy
  surface is a M254 live-gate concern.

## Recommendations
- **DEF-M251-01 → LAND-NEXT (Fate 2, M254).** Confirmed covered by M254 gate parts (g)+(h). No plan edit
  from this lane. ⚠️ **Flag for the M254 driver / orchestrator:** M254's `overview.md` names "~2
  docker/live-gated test-health tests"; the real stackless-box live-gated set is **8**
  (`test_purge` + academy launcher/reap ×5 + `test_host_isolation` mutant-term-trap +
  `test_ant_academy_clerk_wiring`). M254's overview should be corrected to "8" when M254 runs — NOT edited
  from this close window (M254 is not the active milestone; cross-lane collision avoidance).
- **DEF-M251-02 → LAND-NEXT (Fate 2, M247).** M247 owns `verification.md` in the corpus reground; the
  optional demo-stack-suite index anchor folds into that lane. Confirmed covered; no plan edit (M247 is
  closing concurrently — do not touch its files).
- **DEF-M246-03 / -04 → LAND-NEXT (Fate 2, M247).** Confirmed covered by M247's drift-ledger triage
  (M246's designated handoff). No plan edit.
- **DEF-M246-08 → LAND-NEXT (Fate 2, M254 gate h).** Confirmed. No plan edit.
- **DEF-M246-09 → LAND-NEXT (Fate 2, M254 / standing).** Confirmed. No plan edit.

**No Fate-1 landings required. No Fate-3 roadmap mutations made** (no sibling `overview.md` `In:` edited).
**No escape-hatch (cross-release) deferrals** — nothing is release-scope-breaking.

## Applied Changes
- This report written.
- M251 `decisions.md` gained an "Inherited-deferral re-audit (M251 close, Phase 1b)" subsection recording
  the four inherited-item Fate-2 confirmations (D-03/D-04 → M247; D-08/D-09 → M254) + the aging check.
- The **8-not-~2** live-gated count correction is recorded in M251 `decisions.md` + `progress.md` and
  flagged here + in the close report for the serialized orchestrator to apply to M254's overview when M254
  runs. No sibling file edited from this lane.

## Blocking Items (require user decision)
None.
