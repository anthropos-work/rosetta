---
title: "Deferral Audit — M249 cross-app navigation (close)"
date: 2026-07-24
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals; no chronic patterns; no aged-out items. M249 deferred nothing of its own; the one
  surfaced adjacent item has a clear, already-owned fate (Fate 2 → M254).

## Summary
- Total deferrals in scope: 1 (adjacent, not an M249 deferral)
- Single deferrals: 0 (M249 own ledger: Deferred none / Dropped none)
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0

## Deferral Inventory

M249's own ledger is empty:
- `progress.md` § Completeness Ledger → Deferred: _(none)_ · Dropped: _(none)_
- All 5 sections checked in `progress.md`.
- `overview.md` Open questions all resolved in `decisions.md` (rewrite+add per D2/D5; no
  `DEMO_NO_BACK_TO_COCKPIT` knob — the fail-closed conditional render IS the opt-out, so a knob is
  redundant; demo-path only, the cockpit is demo-only). None deferred.
- `spec-notes.md`: no open deferrals.
- No TODO/FIXME/HACK in the touched corpus docs (the code-of-record lives in rext; committed + hardened).

One adjacent item surfaced during the build (not an M249 deferral):

```yaml
- id: DEF-M249-ADJ-01
  item: "2 pre-existing test_ant_academy* failures (launcher/reap flakiness + .env.production.local overlay-extractor bug)"
  origin_milestone: M251   # test-health domain; M249 only re-observed them
  first_deferred_on: 2026-07-23   # fated at M251 close
  last_seen_in: m249-cross-app-nav/decisions.md § "Pre-existing test failures surfaced"
  destination: "M254 (gate parts (g) live-box test-health + (h) live-browser re-prove)"
  reason_recorded: "NOT M249 regressions (verified identical on committed rext HEAD); M251's test domain, not M249's patch domain; subset of the 8 live/env/docker-gated demo-stack failures fated Fate 2 -> M254 at M251 close"
  partial_attempted: no
```

## Repeat-Deferral Patterns
None. The adjacent item was routed exactly once (M251 close → M254); M249 re-observed it and confirms the
routing. It is not being re-deferred.

## Fate-1 Investigation

### DEF-M249-ADJ-01 — "2 pre-existing test_ant_academy* failures"
- **Fate-1 (land now, complete) feasible:** no
- **If no:** Fate 2 (already-owned by M254). These are live/env/docker-gated demo-stack tests
  (spawn real fake HTTP servers on offset ports; execute a bash overlay-extractor under `set -u`). They
  require a live box to green — precisely M254's terrain (gate parts (g)+(h)). They are also **out of M249's
  patch domain** (M251 owns test-health; M254 owns the live re-prove) per the release coordination guardrail.
  M249's diff adds no launcher/reap/bind logic (git-diff-confirmed) and its own `TestAntAcademyWiring`
  covers the `NEXT_PUBLIC_COCKPIT_URL` addition. Landing them here would be cross-domain scope creep.

## Recommendations

### DEF-M249-ADJ-01 → LAND-NEXT (Fate 2 — confirmed covered by M254)
Confirm, don't edit the sibling plan. M254 already owns the 8-failure live-gated set (M251 close routed it
there; M254 gate parts (g)+(h)). Recorded in M249 `decisions.md` § "Pre-existing test failures surfaced" as a
Fate-2 confirm. No `overview.md` edit.

## Applied Changes
- `m249-cross-app-nav/decisions.md` — appended a Fate-2 → M254 confirm line to the existing
  "Pre-existing test failures surfaced" section (no sibling `overview.md` edit).
- This report.

## Blocking Items (require user decision)
None. GREEN — proceed.
