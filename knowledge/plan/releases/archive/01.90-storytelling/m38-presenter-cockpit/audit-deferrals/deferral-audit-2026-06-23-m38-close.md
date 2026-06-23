---
title: "Deferral Audit — M38 close (milestone scope)"
date: 2026-06-23
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals. No chronic patterns. No aged-out items. M38's one surfaced item (M38-D7, the
  employee-hero `org_role=admin` fidelity nuance) is a *fresh, single* finding the close-review owns and
  re-fates here to **LAND-NOW (Fate 1)** — not a repeat push-forward, so it is not a Phase-4 blocker.

## Summary
- Total deferrals in scope: 1 (M38-D7) + the inherited release backlog (carried, unchanged)
- Single deferrals: 1 (M38-D7)
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out items: 0

## Deferral Inventory

### M38's own ledger
- **DEF-M38-01 / M38-D7** — "the employee heroes (Maya/Tom/Sara/Nick) export `org_role=admin`; a hero's
  `org_role` should follow her `vantage` (end-user→member, manager→admin)."
  - origin_milestone: M38 (surfaced by the harden pass 2026-06-23)
  - first_deferred_on: 2026-06-23 (Fate-3 routed to close-review — i.e. THIS audit, by design)
  - last_seen_in: `m38-presenter-cockpit/decisions.md` § M38-D7 + `progress.md` § Hardening "Decision recorded"
  - destination: v1.9 close-review (this close) — the orchestrator explicitly routed the triage here
  - reason_recorded: "the `org_role` is single-sourced from `roleForIndex` (M35 seam); a correct fix must move
    three writes in lockstep [membership row + casbin g2 + roster/JWT claim] — NOT M38-touched code. Bounded
    but the M35 role-assignment seam; orchestrator flagged 'weigh here vs close-review, do NOT force a large
    change.' Behaviorally consistent (claim ≡ seeded DB), vantage-imprecise — fidelity nuance, not a bug."
  - partial_attempted: no (a partial in-progress edit to users.go existed from a crashed prior close attempt;
    re-run completes it fully — not a partial landing)

### Inherited release backlog (carried from prior closes — unchanged, not re-touched by M38)
- DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01..04 (prop-room residuals), M25-D9, M33 —
  all GREEN-audited at the M34 / M35 / M36 / M37 closes (the most recent was M37, 2026-06-23 — same day,
  well inside the 3-month aging window). All orthogonal to the `demo-stack`/`stack-seeding`/`stack-injection`
  surfaces M38 touched; none re-touched by M38. No aging trigger fires.

## Repeat-Deferral Patterns
None. M38-D7 is a fresh single finding (one prior touch, the harden pass that surfaced it). The inherited
backlog has not been re-deferred — it is carried, audited GREEN repeatedly, and unchanged.

## Fate-1 Investigation

### DEF-M38-01 / M38-D7 — "employee heroes export org_role=admin"
- **Fate-1 (land now, complete) feasible:** YES — bounded and safe. Verified in code this close:
  - The fix is ONE new helper `roleForHero(i, n, mix, hero)` in `seeders/users.go` — vantage-faithful for a
    hero slot (`hero.IsManager()`→`admin`, else `member`; `Persona.IsManager()` == `Vantage=="manager"`,
    already defined `blueprint/blueprint.go:200`), `roleForIndex` fallback for non-hero slots.
  - TWO call-sites swap to it: `users.go` seed loop (the membership row + the casbin `g2` grant — already
    edited in the working tree) and `roster.go:93` `BuildRoster` (the exported `org_role` claim — every roster
    entry IS a hero `h`, so the pointer is always non-nil → vantage-faithful). cockpit.go does NOT emit
    `org_role` (it's a vantage-keyed menu projection), so it is NOT part of the role triple.
  - A regression test asserts the three writes agree per hero (membership row role == casbin grant role ==
    roster `org_role`), and that a manager hero → `admin` / an end-user hero → `member`.
  - It does NOT widen release scope (still tooling/seeder-only, zero platform-repo edits), and v1.9 is the
    Stories & Heroes release — landing it makes the employee-vs-manager vantage faithful, which is the whole
    point of the release. The close-review is the right place: it can touch the M35 seam (the close reviews the
    whole milestone-branch + the rext code), and this is the LAST milestone so there is no "later milestone of
    this release" to route to.
- **Decision:** LAND-NOW (Fate 1). See M38-D8 in `decisions.md`; lands in this close's Phase 7.

### Inherited release backlog (DEF-M10-01 / DEF-M21-01..04 / M25-D9 / M33)
- **Fate-1 feasible now:** no — all are orthogonal to M38's surface; each is owned by its own release's
  backlog / a future release, audited GREEN at M37 close < 24h ago. No new context from M38 changes their
  calculus. KEEP (no fresh decision required — not aged, not repeat, not re-touched).

## Recommendations
1. **DEF-M38-01 / M38-D7 → LAND-NOW (Fate 1)** — implement the vantage-faithful `org_role` at the M35 seam
   (`roleForHero` single source + the two call-sites in lockstep + regression test) in this close. The
   cleanest outcome for the final milestone.
2. **Inherited backlog → KEEP** — unchanged, orthogonal, freshly GREEN at M37 close. No action.

## Applied Changes
- M38-D7 re-fated from Fate-3 (route-to-close-review) to **Fate 1 (land now)** — recorded as M38-D8 in
  `decisions.md`; the fix lands in the close's Phase 7.
- No `overview.md` edits (no Fate-2/Fate-3 routing; the item lands here).
- No `RELEASE-SCOPE-DEFER` (no escape-hatch).

## Blocking Items (require user decision)
None. M38-D7 is re-fated to LAND-NOW (Fate 1) — no escape-hatch, no repeat, no aged-out. Audit is GREEN.
