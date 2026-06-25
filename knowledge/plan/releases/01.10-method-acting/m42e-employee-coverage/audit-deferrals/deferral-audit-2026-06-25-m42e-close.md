---
title: "Deferral Audit — milestone M42e (close)"
date: 2026-06-25
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals. M42e added **zero** escape-hatch deferrals of its own; its manager-vantage
  findings are a clean Fate-2/Fate-3 forward-route to the already-scaffolded sibling milestone M42m
  (not a repeat-defer — M42m is the manager half of the same coverage-gate pair). The one inherited
  release-level deferral (DEF-M40-01) is RESOLVED for the employee vantage by this milestone and
  remains a confirmed single Fate-2 for the manager vantage. 0 aged-out, 0 chronic.

## Summary
- Total deferrals in scope: 1 inherited (DEF-M40-01) + 0 new escape-hatch
- Single deferrals: 1 (DEF-M40-01)
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out items: 0

## Deferral Inventory

Sources walked: M42e `progress.md` (running ledger + no `## M42e: Final Review` defer subsection
at audit time), `decisions.md` (AVATAR-LICENSING-BLOCKER, RE-SCOPE-TRIGGER iter-09, TOK-10, TOK-01),
`overview.md` (no `Out:` regressions), `spec-notes.md` (scaffold TODOs — all delivered as
`coverage-protocol.md` + the harness), 23 `iter-NN/` dirs, and the inherited release ledger
(M39/M40/M41 close audits + state.md standing backlog).

```yaml
- id: DEF-M40-01
  item: 'KPI "AI simulations completed" = 0 on the profile/dashboard surface'
  origin_milestone: M40
  first_deferred_on: 2026-06-24
  last_seen_in: m41 close audit (2026-06-24) — Fate-2 → M42e/M42m, unchanged
  destination: "M42e (employee vantage) + M42m (manager vantage) coverage gates"
  reason_recorded: "the KPI reads jobsimulation directly (no CMS dependency); a separate
    frontend/auth-context concern owned by the iterative coverage milestones whose gate
    encompasses a non-zero completed-KPI on a reachable page"
  partial_attempted: no
```

Two M42e `decisions.md` items examined and confirmed NOT to be open deferrals:

- **AVATAR-LICENSING-BLOCKER (iter-16):** a user-blocker that was **RESOLVED in-milestone** —
  the user chose option (a) synthetic non-existent-person photoreal faces (iter-18 landed the
  `photo_avatar` synthetic-face generator; the menu==profile + real-photo persona assert PASSES
  in the final gate). Not an open deferral.
- **RE-SCOPE-TRIGGER iter-09 (`/reimport-profile` → linkedin.com/help):** RESOLVED in-milestone —
  the editorial/help-link allow-rule (disclose-not-hide presenter notes) was applied; the final
  employee gate reports `escapes=0`. Not an open deferral.

The `spec-notes.md` `TODO:` lines (cockpit login / crawl / per-page assertions / report schema /
fix surfaces / the protocol doc) are **build-time scaffold placeholders authored at milestone
creation** — every one is now implemented (`coverage-protocol.md` shipped; the `stack-verify/e2e/`
harness shipped with `cockpit-login.ts`, `crawl.ts`, `section-assert.ts`, `persona-assert.ts`,
`empty-states.ts`, `coverage-manifest.ts`, `review-html.ts`). They are scaffold debris, not
deferred work (flagged for cosmetic cleanup in the Phase-3 docs review, not a deferral).

## Repeat-Deferral Patterns
None. DEF-M40-01 has been seen in exactly one origin milestone (M40); it does not recur as a
*new* deferral in M41 or M42e — it is the same single item being progressively resolved along its
Fate-2 destination chain. M42e introduced no escape-hatch deferral.

## Fate-1 Investigation

### DEF-M40-01 — KPI "AI simulations completed" = 0 (the profile stats-row KPI)
- **Fate-1 (land now, complete) for the EMPLOYEE vantage:** **RESOLVED by M42e.** The
  semantic coverage manifest carries a `/profile` stats-row descriptor asserting the
  "...Completed" KPIs render as real semantic content (`minMeaningfulLen: 60`), and the
  **employee gate PASSES that section** (0 failing sections). iter-14's `HeroActivitySeeder`
  seeds the activity rows that drive the stats row (≥1 completed skill_path_session → "Skill
  Paths Completed">0; personal_assignments + bookmarks populate the activity surfaces). The
  per-session AI-sim *result deep-link* is a genuinely runtime-computed artifact (iter-04
  falsified the seedable-row hypothesis) and is correctly crawl-scoped out — but the *KPI count*
  on the stats row is now a passing populated section. So the **employee half of DEF-M40-01 is
  Fate-1 landed** within M42e.
- **Manager vantage:** the Workforce-dashboard KPIs live on the `/workforce/*` M36 pages, which
  the M42e final manager smoke-sweep found **unreached (notReached=5)** because the manager nav
  doesn't link them yet. That is the core M42m content+nav work. **Fate-2 → M42m** (confirmed —
  M42m's exit gate is identical to M42e's, run as a manager hero, and explicitly covers the
  Workforce dashboard funnel; no plan edit needed, M42m already owns it).
- **Aging:** no trigger fired. The destination M42m is the next active milestone, not closed;
  the surface was substantively advanced (not stalled) by M42e.

### M42e manager-vantage findings (surfaced by the final smoke-sweep — M42m input)
Not deferrals of M42e (employee) scope — they are **manager-scope discoveries** that belong to
M42m by design (the two milestones are the employee/manager halves of one coverage-gate pair):
- **139 `studio.anthropos.work` escapes** = one baked left-nav prod link rendered across pages →
  the demo-injection **link-rewriting** fix surface. **Fate-3 → M42m** (M42m overview annotated;
  manager-manifest descriptors authored, `calibrated:false` until workforce renders).
- **5 unreached `/workforce/*` M36 pages** = the core manager content+nav work. **Fate-2 → M42m**
  (its exit gate already owns the Workforce dashboard surfaces).
- **team-roster `/user/<id>` fan-out (frontier CAPPED +79)** = needs a representative-sample
  crawl rule + higher cap. **Fate-3 → M42m** (annotated).

These are a clean forward-route to a sibling milestone of the SAME release, not a punt — M42m is
scaffolded, its overview carries the manager scope, and its manager-manifest descriptors are
already authored. NOT a repeat-defer (M42e never owned the manager vantage).

## Recommendations

| Item | Verdict | Rationale |
|------|---------|-----------|
| DEF-M40-01 (employee half — KPI on profile stats row) | **LAND-NOW** (Fate-1, already done in M42e) | The employee gate passes the `/profile` stats-row section; the KPIs render real content. No further action. |
| DEF-M40-01 (manager half — Workforce KPIs) | **LAND-NEXT** (Fate-2 → M42m) | Owned by M42m's identical exit gate (manager vantage; covers `/workforce/*`). Confirmed unchanged; no plan edit. |
| M42e manager escapes (139 studio links) | **LAND-NEXT** (Fate-3 → M42m) | Link-rewriting fix surface; M42m overview annotated, manager descriptors authored. |
| M42e unreached `/workforce/*` (5 pages) | **LAND-NEXT** (Fate-2 → M42m) | M42m's core content+nav work; already in its scope. |
| M42e team-roster fan-out (sample rule) | **LAND-NEXT** (Fate-3 → M42m) | Crawl representative-sample rule; M42m annotated. |
| Standing backlog (DEF-M10-01, DEF-M21-01, M25-D9) | unchanged | Unscheduled, orthogonal to v1.10; not in scope, not aged into v1.10. |

## Applied Changes
- No plan edits required at this audit. DEF-M40-01's employee half is resolved in-milestone;
  its manager half + M42e's manager-vantage findings are confirmed Fate-2/Fate-3 to M42m, which
  already owns them (overview annotated during the iter loop; manager-manifest descriptors
  authored in `coverage-manifest.ts`). The M42e `decisions.md` AVATAR + RE-SCOPE-TRIGGER entries
  are RESOLVED-in-milestone, not open deferrals.

## Blocking Items (require user decision)
None. No repeat deferrals, no aged-out items, no chronic patterns, no escape-hatch deferrals.
Verdict **GREEN**.
