---
title: "Deferral Audit — M52 close (milestone scope)"
date: 2026-07-01
scope: milestone
invoked-by: close-milestone
---

## Verdict
**GREEN**

- No repeat-deferrals originating in M52; the one prior REPEAT in this release (academy F6) was already
  resolved with an explicit user LAND-NEXT fate at M51 close (D-CLOSE-1). No chronic pattern. Every item in
  scope has a clear, current fate decision.

## Summary
- Total deferrals in scope: **4** (1 M52-originated + 3 inherited release-level carries visible at M52 close)
- Single deferrals: **4**
- Repeat deferrals: **0** _(the academy-F6 REPEAT was resolved at M51 close; it is a settled LAND-NEXT, not an open repeat)_
- Chronic patterns flagged: **0**
- Aged-out: **0**

## Deferral Inventory

```yaml
- id: DEF-M52-01
  item: "up-injected.sh end-to-end glue (exports the consolidated manifest via --manifest-export/--gen-seed + passes --seed-manifest to the cockpit at bring-up) has no shell unit-test harness"
  origin_milestone: M52
  first_deferred_on: 2026-07-01
  last_seen_in: m52-seed-manifest/progress.md:107 (Pass-3 out-of-scope note)
  destination: "M53 (cold-rebuild acceptance)"
  reason_recorded: "No shell unit harness exists; its behavior is already fenced by the Go CLI tests (doManifestExport ±--gen-seed, unwritable-out) + the Python cockpit tests (present/absent non-fatal). End-to-end exercise belongs to M53's cold-rebuild proof."
  partial_attempted: no

- id: DEF-M52-02 (inherited — release-level, not M52-originated)
  item: "Consumption-clone re-pin to the release rext tag + .agentspace/rext.tag bump"
  origin_milestone: M47
  first_deferred_on: 2026-06-29
  last_seen_in: m51-ai-readiness-org/progress.md:66
  destination: "M53 (push-gated KEEP; authoritative box-level bump)"
  reason_recorded: "Box-level authoritative bump at M53 (v1.10.1); needs origin pushes. Explicitly 'not a repeat-defer' — a push-gated administrative KEEP."
  partial_attempted: no

- id: DEF-M52-03 (inherited — release-level, user-decided at M50)
  item: "COLD reset-to-seed acceptance (whole release proven from cold)"
  origin_milestone: M50
  first_deferred_on: 2026-06-30
  last_seen_in: m51-ai-readiness-org/progress.md:65
  destination: "M53"
  reason_recorded: "M53 In: owns the from-cold rebuild acceptance (user-decided at M50, D-CLOSE-2). Fate-2 carry-forward, not escape-hatch."
  partial_attempted: no

- id: DEF-M52-04 (inherited — resolved REPEAT, user-decided at M51)
  item: "Academy F6 (course content + hero menu-link + non-anonymous academy session)"
  origin_milestone: M50
  first_deferred_on: 2026-06-30
  last_seen_in: m53-cold-rebuild-acceptance/overview.md:44-50
  destination: "M53 (LAND-NEXT, Fate-3 override of M53's 'No new feature code' guard)"
  reason_recorded: "Was a REPEAT (M50 Fate-3→M51, not executed). Resolved at M51 close (D-CLOSE-1) with an explicit user LAND-NEXT → M53; M53 overview In: now owns it. Settled."
  partial_attempted: no
```

## Repeat-Deferral Patterns
**None open.** The only historical REPEAT in this release (academy F6, deferred M50→M51 without landing)
was detected + resolved at the M51 close deferral re-audit
(`m51-ai-readiness-org/audit-deferrals/deferral-audit-2026-07-01-m51-close.md`), with an explicit
user LAND-NEXT → M53 fate (D-CLOSE-1). It carries a settled destination and is no longer an open repeat.

## Fate-1 Investigation

### DEF-M52-01 — up-injected.sh end-to-end glue
- **Fate-1 (land now, complete) feasible:** no
- **Why not:** up-injected.sh is bring-up orchestration glue — a `[ -f ]` gen-seed guard, a non-fatal export
  fallback, and `--seed-manifest` arg-wiring. Its constituent behaviors ARE already unit-fenced: the Go CLI
  tests exercise `doManifestExport` with and without `--gen-seed` and an unwritable output path; the Python
  cockpit tests exercise the manifest-present and manifest-absent (non-fatal fallback) paths. What is NOT
  unit-testable is the *composition* — a real `/demo-up` running the whole script end-to-end. That has no
  standalone unit surface in either stack and building a shell harness for a 3-line bring-up hook in this
  milestone would be a scaffold, not a real feature landing.
- **Which fate:** **Fate-2 (already owned by M53).** M53's `overview.md` `In:` explicitly asserts (line 43)
  "the cockpit **[Download manifest]** returns the **complete inlined** `seed-generation-manifest.yaml` (M52)"
  on a cold `/demo-up` (line 38: set-dress + seed + verify + cockpit all complete). A cold `/demo-up` runs
  up-injected.sh — so the end-to-end exercise is genuinely already in M53's acceptance checklist. **No plan
  edit needed** (confirm-only).

### DEF-M52-02 — consumption-clone re-pin
- **Fate-1 feasible:** no (push-gated; requires origin pushes the orchestrator still owes).
- **Which fate:** **Fate-2 (already owned by M53, push-gated KEEP).** Authoritative box-level bump at M53.
  Explicitly not a repeat-defer — an administrative KEEP tracked release-wide since M47. No edit.

### DEF-M52-03 — COLD reset-to-seed acceptance
- **Fate-1 feasible:** no (M52 is fix-on-live; the cold proof is M53's defining work by design).
- **Which fate:** **Fate-2 (already owned by M53).** User-decided at M50 (D-CLOSE-2). No edit.

### DEF-M52-04 — Academy F6
- **Fate-1 feasible:** no (M52 owns only the seed+gen manifest; academy content is out of M52's domain).
- **Which fate:** **Settled LAND-NEXT → M53** (user-decided at M51 close, D-CLOSE-1; M53 `In:` owns it). No edit.

## Recommendations
- **DEF-M52-01** → **LAND-NEXT (Fate-2)** — confirmed covered by M53's cold-rebuild acceptance. No plan edit.
- **DEF-M52-02** → **LAND-NEXT (Fate-2, push-gated KEEP)** — M53. No plan edit.
- **DEF-M52-03** → **LAND-NEXT (Fate-2)** — M53. No plan edit.
- **DEF-M52-04** → **LAND-NEXT (Fate-3, settled)** — M53. Already recorded (D-CLOSE-1). No plan edit.

## Applied Changes
None. Every in-scope item already has a current, correct fate with its destination (M53) reflected in
`m53-cold-rebuild-acceptance/overview.md`. This audit is a confirm-only pass; no `overview.md`/`decisions.md`
edits were required.

## Blocking Items (require user decision)
**None.** No open repeat-deferrals, no aged-out items, no escape-hatch entries. Verdict GREEN — close proceeds.
