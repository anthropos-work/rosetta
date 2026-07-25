---
title: "Deferral Audit — M252 close (studio-desk builder enablement)"
date: 2026-07-24
scope: milestone
invoked-by: close-milestone
---

## Verdict
**GREEN** — no repeat deferrals; no aged-out items; every item has a clear fate decision. Both open
carries are Fate-2 to **M254 "prove on billion"**, the release's designed terminal live-prove milestone
(barrier → fan-out → prove-on-billion), which is the intended architecture for a live-confirmation carry,
not scope erosion. Same shape as this release's M250 close (CARRY-M250-01 Fate-2 → M254, audit GREEN).

## Summary
- Total deferrals in scope: 3 (2 open carries + 1 landed-Fate-1)
- Single deferrals: 2
- Repeat deferrals: 0
- Aged-out: 0
- Chronic patterns flagged: 0

## Deferral Inventory

- id: CARRY-M252-01
  item: "5 PRE-EXISTING stack-verify unit failures (TestAutoVerify + TestDirectusCheapWins) — M245 academy `/library/` cheap-win added without updating 3 older curl stubs"
  origin_milestone: M252 (surfaced at build; root cause M245, prior release v2.6)
  first_deferred_on: 2026-07-24
  last_seen_in: m252/decisions.md:53
  destination: "M254 exit-gate part (g) — the live/docker-gated test-health tests green"
  reason_recorded: "byte-identical pre/post M252 (HEAD 584f1fe); M252 adds 0 failures; out-of-subject academy-harness reconciliation carrying its own cross-class-stub-gating risk; belongs to the test-health / prove-on-billion domain, not the studio-builder milestone"
  partial_attempted: no

- id: CARRY-M252-02
  item: "the LIVE end-to-end builder-generate Playthrough drive (the ~10-min async Generate) — the studio sim-builders actually generate, proven live on billion"
  origin_milestone: M252
  first_deferred_on: 2026-07-24
  last_seen_in: m252/decisions.md (recorded this close) + commit cedde09 message
  destination: "M254 exit-gate parts (e) [studio sim-builders generate — builder Playthrough green] + (h) [live Playthroughs green]"
  reason_recorded: "the wiring is PROVEN live (op1 on demo-2: container carries the AI keys + boots ProviderHealth chain azure-openai->openai->anthropic); the Playthrough is AUTHORED + committed + ptvalidate 18-live/0-TODO + unit-green; the ~10-min async live RUN is flaky locally and, per the release design, ALL live-proof is centralized at the terminal M254 milestone; user close-mandate 2026-07-24"
  partial_attempted: no

- id: KB-1 (LANDED — not an open deferral)
  item: "secrets-spec.md M50 'AI keys waived/optional' imprecise for the studio-desk AI genes (they are required·standard)"
  origin_milestone: M252
  destination: "LANDED as the M252 secrets-spec.md deliverable — the studio-desk AI carve-out (the KB-1 correction) is in the committed doc"
  reason_recorded: "Fate-1 complete; tracked at KB-fidelity pre-flight, delivered in cedde09"
  partial_attempted: no

## Repeat-Deferral Patterns
None. CARRY-M252-01 has no prior appearance (M251 test-health — the disjoint run-unit/cockpit re-point area —
has zero mention of autoverify/TestDirectusCheapWins/the academy `/library/` cheap-win). CARRY-M252-02 is
studio-specific and first-appearance. The three release-level Fate-2→M254 carries (CARRY-M248-01,
CARRY-M250-01, CARRY-M252-01/02) are all distinct items, not the same item re-deferred.

## Aging
No aged-out items. Both carries were created today against a live, open M254 whose exit gate explicitly
owns them; M254 has not closed; the affected areas were not re-touched by a later closed milestone.

## Fate-1 Investigation

### CARRY-M252-01
- **Fate-1 (land now, complete) feasible:** no
- **If no:** Fate 2 — M254 gate (g) already owns "the live/docker-gated test-health tests green," and
  close-release runs a release-level quality review. The fix (mirror the `*"/library/")` stub arm into the
  3 older stubs) is an academy-harness reconciliation out-of-subject for a studio-builder milestone and
  carries its own cross-class stub-gating risk. Landing it here would smuggle unrelated test-harness work
  into a studio wiring milestone. Genuine domain fit is M254 (g) / release-close quality review.

### CARRY-M252-02
- **Fate-1 (land now, complete) feasible:** no
- **If no:** Fate 2 — M254 is the release's terminal, single-driver live-prove milestone (the M221/M236/M244
  lineage). Its exit gate parts (e) + (h) already commit to the studio builder Playthrough green + the live
  Playthroughs green on billion. The release deliberately centralizes ALL cold-reset-to-seed live-proof at
  M254 (the "billion-last" design); a per-milestone live-drive would duplicate M254's single-box serial
  drive. The M252 BUILD deliverables (wiring + DNA assert + Playthrough artifacts) are complete; only the
  live RUN — M254's chartered domain — remains.

## Recommendations
- CARRY-M252-01 → **LAND-NEXT** (Fate 2, confirmed covered by M254 gate (g)). No plan edit — M254 already owns it.
- CARRY-M252-02 → **LAND-NEXT** (Fate 2, confirmed covered by M254 gate (e)+(h)). No plan edit — M254 already owns it.
- KB-1 → already **LAND-NOW** (Fate 1) — delivered in the M252 secrets-spec.md deliverable.

## Applied Changes
- Recorded CARRY-M252-02 (the live builder-generate drive → Fate-2 → M254 (e)+(h)) in `m252/decisions.md`.
- No sibling `overview.md` edits (both carries are Fate 2 — M254 already owns them via its exit gate; no
  Fate-3 annotation needed).

## Blocking Items (require user decision)
None. GREEN.
