---
title: "Deferral Audit — M246 (close)"
date: 2026-07-23
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals (M246 is the FIRST milestone of v2.7 — nothing inherited).
- No aged-out items (every record was first written today, 2026-07-23).
- Every item has a clear fate decision routed through the durable, checked-in
  **drift ledger** (`drift-ledger.md`), which M247 formally consumes (`depends_on: [M246]`).

## Summary
- Total deferrals in scope: 9 (the drift-ledger rows D-01..D-09) + 0 from progress/decisions
- Single deferrals: 9
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0
- Blocking items: 0

## Deferral Inventory

**M246's own progress/decisions carry ZERO scope deferrals.** All 4 declared sections are checked off
and delivered (Fate 1). The Completeness Ledger `### Deferred` and `### Dropped` sections are empty.
`overview.md`'s `Out:` items (corpus reconciliation, fidelity fixes, platform edits, dev-worktree
touches) were **always out** — never In:-scope that moved — so they are not deferrals.

The only forward-routed items are the **9 confirmed-drift rows** — M246's *designed output* as a HARD
go/no-go barrier. Each is a divergence between what the corpus/tooling asserts and what the consolidated
platform now is, surfaced by the cold `/demo-up` prove, and routed to its domain milestone.

| id | item | origin | first_deferred | destination | inert? |
|---|---|---|---|---|---|
| DEF-M246-01 (D-01) | Corpus asserts skillpath live Tier-1 / 4th subgraph (~30 files) | M246 | 2026-07-23 | M247 | doc-drift |
| DEF-M246-02 (D-02) | rext `gen_injected_override.py` `INJECTED` dict dormant skillpath key ("4 injected") | M246 | 2026-07-23 | M247-triage / rext-hygiene | inert (no compose service; comment made truthful M246 `88bcdb8`) |
| DEF-M246-03 (D-03) | `test_injection.py` pins skillpath-as-injected + models already-merged skiller | M246 | 2026-07-23 | M251 / rext-hygiene | inert (behavioural-test fixture; not needed for green) |
| DEF-M246-04 (D-04) | `exposure_claim_guard.py` `_cfg` lists skillpath:8095 | M246 | 2026-07-23 | M251 / rext-hygiene | inert (test-only fixture; mirrors D-03) |
| DEF-M246-05 (D-05) | Stale `stack-demo/skillpath/` clone dir lingers | M246 | 2026-07-23 | housekeeping (no milestone) | cosmetic (never re-cloned/built) |
| DEF-M246-06 (D-06) | `up-injected.sh:458` audit-prose comment names historical skiller/skillpath | M246 | 2026-07-23 | M247 | cosmetic prose |
| DEF-M246-07 (D-07) | AI-readiness `loadMembers` perf demopatch anchor stale (file moved in bump) | M246 | 2026-07-23 | M250 (its domain) | non-fatal perf (~180 s unbounded hydration) |
| DEF-M246-08 (D-08) | fake-FAPI http-vs-TLS cheap-win probe artifact + end-to-end login not re-run | M246 | 2026-07-23 | M251 (probe fix) / M254 (login re-prove) | probe artifact (container Up, roster loaded) |
| DEF-M246-09 (D-09) | AI Academy peripheral (2/6 env keys short) not serving | M246 | 2026-07-23 | standing/peripheral (opt. M251/M254) | benign peripheral (non-fatal by design) |

## Repeat-Deferral Patterns
None. M246 is the release opener; every record is a first-pass routing. No item has been deferred across
≥2 milestones, none is ≥3 months old, no destination milestone has closed, no later milestone has yet
touched these areas. Zero REPEAT groups, zero CHRONIC_DEFER, zero DRIFT_DEFER.

## Fate-1 Investigation
**Fate-1 (land in M246) is correctly NO for all 9 rows — by DESIGN, not by time pressure.**

M246 is a HARD go/no-go barrier. Its declared scope (seeder re-point · clone pins · comment fix ·
de-skillpath the LIVE bring-up path · prove green · emit the ledger) is fully delivered as Fate 1.
Landing the drift rows in M246 would (a) violate the milestone's own `Out:` scope ("the corpus doc
reconciliation (M247); any fidelity fix"), (b) pre-empt M247's designated triage role (M247
`depends_on: [M246]` = the ledger), and (c) for D-01 specifically, duplicate ~30 files of work that is
literally M247's declared `In:` scope. The barrier's contract is: *surface drift, route it, do not fix it
here.* Every row is genuinely a sibling-milestone concern.

## Recommendations

- **DEF-M246-01 (D-01) → LAND-NEXT (Fate 2).** M247's `In:` list explicitly owns "convert
  `skillpath.md` → merged-into-app redirect" + "re-point every 4-subgraphs → 3 + reclassify skillpath
  not-a-live-service across ~30 echo files." Confirmed covered; no plan edit.
- **DEF-M246-02/03/04 (D-02/03/04) → LAND-NEXT (Fate 2, via the drift-ledger handoff).** Inert rext
  hygiene (no compose service matches → no functional impact; the barrier is GREEN). Durably captured in
  the checked-in `drift-ledger.md`, which M247 formally consumes and triages ("some rows name a
  better-fit milestone"). M247's triage assigns the final rext-hygiene home (M247-adjacent or M251).
  No overview edit — pre-empting M247's triage on inert cosmetic items would be premature roadmap churn.
- **DEF-M246-05/09 (D-05/D-09) → no milestone action (housekeeping / standing peripheral).** Cosmetic;
  removed on next clean workspace / optional M251/M254 follow-up. Durably ledgered.
- **DEF-M246-06 (D-06) → LAND-NEXT (Fate 2).** Cosmetic audit-prose; M247 (corpus/prose re-ground).
- **DEF-M246-07 (D-07) → LAND-NEXT (Fate 2, its domain M250).** M250 (AI-readiness fidelity, iterative)
  owns the AI-readiness surface end-to-end; its live measure→triage→fix→re-render loop ("0 prod-ejects,
  closure green", manager-vantage fidelity) encounters + re-anchors the stale perf demopatch. Ledgered
  to M250.
- **DEF-M246-08 (D-08) → LAND-NEXT (Fate 2).** The end-to-end login re-prove is M254 gate-part (h)
  (live-browser specs + p95 click→ACCESS). The cheap-win probe http-vs-TLS fix is M251/M254 test-health.
  Ledgered.

**No Fate-1 landings required. No Fate-3 roadmap mutations made** (no sibling `overview.md` `In:` list was
edited — every row routes via the durable drift-ledger handoff / M247's designated triage). No escape-hatch
(cross-release) deferrals — nothing is release-scope-breaking.

## Applied Changes
- Report written to this path. The confirmed-covered mapping is already recorded in M246 `decisions.md`
  (KB-4, KB-5, D-3, D-4) and in `drift-ledger.md` (the durable Fate-2 tracking vehicle). No sibling plan
  edits. No new blocking decisions required.

## Blocking Items (require user decision)
None.
