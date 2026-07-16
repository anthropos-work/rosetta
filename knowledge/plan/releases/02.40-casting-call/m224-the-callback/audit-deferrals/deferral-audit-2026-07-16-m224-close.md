---
title: "Deferral Audit — milestone (M224 the-callback, close)"
date: 2026-07-16
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals; no chronic patterns; no aged-out items.
- Every M224 scope-delivery gap surfaced at close (the demopatch-spec fold-in + three overlooked
  `Delivers → knowledge/corpus` sections) is **LAND-NOW (Fate-1)** in this close — none is deferred.
- The single carry is a set of **8 pre-existing, HEAD-identical test failures in files M224 never touched** —
  an inherited known-issue flagged for sign-off per the close-milestone Phase 4/8 carry mechanism, not a punt of
  M224's own work.

## Summary
- Total deferrals in scope: **5** (4 resolved-at-close Fate-1 landings + 1 inherited known-issue carry)
- Single deferrals: 1 (the inherited test-debt carry)
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0

Inherited from prior v2.4 milestones: **none pending.** DEF-M222-01 (replay `directus.job_position`) was applied
as **Fate-3/DROPPED** into M223 (0 rows captured; the scoreboard does not read the entity) and M223 closed with
S3 confirming the drop. M223 recorded **no new deferrals** (S6 folded into S2; its one forward-note — `roleForHero`
returning `member` — was M224 scope and was CONSUMED by M224 iter-11's `endUserHeroRole`→candidate fork).

## Deferral Inventory

```yaml
- id: DEF-M224-01
  item: "demopatch-spec.md §4/§5 — note the HIRING image bakes 4 patches (2 net-new apps/hiring + 2 chained shared urls.ts)"
  origin_milestone: M224
  first_deferred_on: 2026-07-16
  last_seen_in: m224/progress.md 'Next' fold-in (a) ; hardening-ledger.md 'Knowledge backfill'
  destination: "LAND-NOW (Fate-1) at close"
  reason_recorded: "left in-lane at harden rather than double-authored; a small in-domain doc update to do properly at close"
  partial_attempted: no

- id: DEF-M224-02
  item: "cockpit-spec.md — the M224 hero-trio + hiring-vantage delivery (Delivers-> corpus)"
  origin_milestone: M224
  first_deferred_on: 2026-07-16   # surfaced by this close's Phase-1 scope review
  last_seen_in: m224/overview.md 'Delivers ->' ; kb-fidelity-audit.md (named an M224 Delivers deliverable)
  destination: "LAND-NOW (Fate-1) at close"
  reason_recorded: "not previously logged as a deferral — an overlooked Delivers-> corpus section; content exists verified in spec-notes.md"
  partial_attempted: no

- id: DEF-M224-03
  item: "clerkenstein.md — the isHiring FAPI org publicMetadata emission mechanism (Delivers-> corpus / BLIND-AREA fill)"
  origin_milestone: M224
  first_deferred_on: 2026-07-16
  last_seen_in: m224/spec-notes.md 'Clerkenstein isHiring wiring' ; kb-fidelity-audit.md (BLIND-AREA = M224 Delivers deliverable)
  destination: "LAND-NOW (Fate-1) at close"
  reason_recorded: "overlooked corpus fill; iter-02 landed the wiring + the /align-run record; content exists verified in spec-notes.md"
  partial_attempted: no

- id: DEF-M224-04
  item: "hiring.md render-path — reflect the shipped TOK-02 reality (render via the REAL apps/hiring two-app demo, not an apps/web re-skin)"
  origin_milestone: M224
  first_deferred_on: 2026-07-16
  last_seen_in: corpus/services/hiring.md:201 (frames the scoreboard as apps/web only — pre-TOK-02)
  destination: "LAND-NOW (Fate-1) at close"
  reason_recorded: "stale/incomplete vs shipped — TOK-02 pivoted the render to the real apps/hiring; a Phase-3 stale-claim correction"
  partial_attempted: no

- id: DEF-M224-05
  item: "8 pre-existing test failures — 6 demo-stack/tests/test_cockpit.py (4 removed-academy-CTA + 2 v2.3.1 overlay-JS) + test_purge + test_reap"
  origin_milestone: M224   # first FORMALLY observed here (M222/M223 were stack-seeding/hiring.md milestones; they did not run the demo-stack python suite)
  first_deferred_on: 2026-07-16
  last_seen_in: hardening-ledger.md Pass-1 'Verification (Phase 5)' ; HEAD-identical across the milestone boundary
  destination: "CARRY — inherited known-issue → standing test-debt backlog (routed to a future demo-stack test-debt harden pass)"
  reason_recorded: "HEAD-identical, in files M224 never touched; predate v2.4 (v2.3.1/v2.3.2 cockpit hotfixes removed the academy CTA + changed the overlay JS without updating the tests); not the render gate; non-blocking to the gate"
  partial_attempted: no
```

## Repeat-Deferral Patterns
None. Every record is a first-and-only occurrence created at this close. DEF-M224-05 is the first close to *run*
the demo-stack python suite in v2.4, so it is a first observation, not a repeat.

## Fate-1 Investigation

### DEF-M224-01…04 — the four corpus/doc landings
- **Fate-1 (land now, complete) feasible:** **yes** — all four are in-domain documentation fold-ins whose content
  is already verified in the milestone's own `spec-notes.md` / `decisions.md` / `hardening-ledger.md` (file:line
  cited). No code, no new investigation. Landed in full this close (no partial slices).
- **Fate applied:** **Fate-1.**

### DEF-M224-05 — the 8 pre-existing test failures
- **Fate-1 (fix now) feasible:** **no, and correctly so.** These live in files **M224 never touched**
  (`test_cockpit.py`, `test_purge`, `test_reap`) and encode v2.3.x cockpit/demo-down behavior (the removed academy
  CTA, the v2.3.1 overlay-JS hotfix, purge/reap env-deps) — **outside M224's hiring render-loop domain.** Fixing
  them is separate test-debt, not M224 work; landing it here would be scope-bleed, not Fate-1 discipline.
- **Fate applied:** **CARRY as a flagged inherited known-issue** (close-milestone Phase 4/8 mechanism) →
  **standing test-debt backlog**, to be picked up by a future demo-stack test-debt harden pass. Non-blocking to the
  M224 exit gate (the gate is a live render probe, GREEN 3-cold-run + 4/4 flake). **Not a release-scope-breaking
  punt of desired M224 work** — so no escape-hatch sign-off is triggered; it is an inherited-failure carry, flagged
  in the Gate Outcome Ledger for visibility.
- **Aging:** not aged out (first observation 2026-07-16).

## Recommendations
- **DEF-M224-01 → LAND-NOW (Fate-1):** demopatch-spec §4/§5 fold-in. **Applied this close.**
- **DEF-M224-02 → LAND-NOW (Fate-1):** cockpit-spec.md hero-trio + hiring-vantage section. **Applied this close.**
- **DEF-M224-03 → LAND-NOW (Fate-1):** clerkenstein.md isHiring FAPI section. **Applied this close.**
- **DEF-M224-04 → LAND-NOW (Fate-1):** hiring.md render-path TOK-02 correction. **Applied this close.**
- **DEF-M224-05 → CARRY (inherited known-issue):** flag in the Gate Outcome Ledger + the standing backlog; route to
  a future demo-stack test-debt harden pass. Do not fix in M224.

## Applied Changes
- The four Fate-1 landings are executed in the close's Phase-7 fix pass (corpus docs) and recorded in
  `../decisions.md` (D2–D5, the deferral→active conversions).
- The DEF-M224-05 carry is recorded in the Gate Outcome Ledger (`../progress.md`) and the release standing backlog
  (`state.md`); `../decisions.md` D6 notes the carry.

## Blocking Items (require user decision)
None. No repeat / chronic / aged-out deferrals. Verdict GREEN.
