---
title: "Deferral Audit — M227 the-notes (close)"
date: 2026-07-17
scope: milestone
invoked-by: close-milestone
---

## Verdict
**YELLOW** — one NEW single deferral (section-5 LOCAL live-render re-prove → M228, already planned, clean
Fate-2) + the same INHERITED carries M225/M226 already routed to the v2.4 **release close**, re-confirmed fresh
today. No M227-native repeat; no chronic/silent scope erosion. Every item has a conscious, imminent destination
(M228, or the release close — both one step away). Consistent with the M225 (YELLOW) + M226 (YELLOW) audits.

## Summary
- Total deferrals in scope: 4
- Single (M227-native): 1 (DEF-M227-01)
- Inherited carries (re-confirmed): 3 (DEF-CARRY-A, DEF-CARRY-B, DEF-M226-01)
- Chronic patterns flagged: 0

## Deferral Inventory

```yaml
- id: DEF-M227-01
  item: "Section 5 — the LOCAL live full-stack render/coverage/playthrough re-prove on the corrected data"
  origin_milestone: M227
  first_deferred_on: 2026-07-17
  last_seen_in: m227-the-notes/decisions.md D5 + progress.md §Section 5
  destination: "M228 second-night (billion re-prove — already planned in roadmap.md)"
  reason_recorded: "ENVIRONMENTAL, not code: a /demo-down --purge removed the working demo's images, the cold
    rebuild hit ENOSPC, builder-prune evicted the go-build cache → ~35-min cold recompiles + buildx wedged under
    host CPU contention. The DATA correctness of all 4 fixes is proven DETERMINISTICALLY by the unit+regression
    suite (stronger than ad-hoc SQL); the GREEN live render/coverage/playthrough is M228's explicit deliverable."
  partial_attempted: no   # the DATA is fully proven deterministically; the live RENDER is wholly routed to M228,
                          # not sliced. This is a clean fate-boundary, not a disguised partial landing.

- id: DEF-CARRY-A
  item: "8 pre-existing demo-stack test failures (6× test_cockpit + test_purge + test_reap)"
  origin_milestone: pre-v2.4 (HEAD-identical; in files M222–M227 never touched)
  first_deferred_on: 2026-07-17 (surfaced/tracked at M225)
  last_seen_in: state.md §Standing backlog + m226 deferral-audit DEF-CARRY-A
  destination: "v2.4 RELEASE close (future demo-stack test-debt harden pass)"
  reason_recorded: "Pre-existing, out of every v2.4 milestone's scope; predates the release. Re-confirmed fresh."
  partial_attempted: no

- id: DEF-CARRY-B
  item: "M204 assign-and-track.UC1 assign-WRITE — declared in-manifest `unimplemented` (ptvalidate) TODO"
  origin_milestone: M204 (v2.0)
  first_deferred_on: 2026-07-17 (re-confirmed at M225/M226)
  last_seen_in: state.md §Standing backlog + m226 deferral-audit
  destination: "v2.4 RELEASE close (its declared-TODO fate)"
  reason_recorded: "A declared, intentional build-reference gap; out of M227 scope (M227 touched no playthrough
    WRITE path). Re-confirmed fresh."
  partial_attempted: no

- id: DEF-M226-01
  item: "Pre-bind serve reap — clear stale `tailscale serve` fronts on offset ports before bind (Finding-3)"
  origin_milestone: M226
  first_deferred_on: 2026-07-17
  last_seen_in: m226 deferral-audit DEF-M226-01
  destination: "next prove-on-VM milestone (M228) / a follow-up build-iter — self-resolves in the default flow"
  reason_recorded: "Fate 3, non-gate-blocking; a bring-up-path change on a live-only surface needing a live
    re-prove (forbidden at close; billion left UP). M227 is a LOCAL, tooling-only milestone that neither touched
    nor could re-prove this surface."
  partial_attempted: no
```

## Repeat-Deferral Patterns

**DEF-CARRY-A / DEF-CARRY-B** have now been carried across M225 → M226 → M227 (≥2 milestones), so by the letter of
Phase 2 they are REPEAT. But they are **not silent scope erosion**: each carry was a *conscious, recorded*
re-fating (M225 YELLOW audit, M226 YELLOW audit, and this pass), every one is **entirely outside every v2.4
milestone's touched surface** (pre-existing demo-stack tests / a v2.0 playthrough WRITE path — M227 touched only
`stack-seeding` + the render probe + `hiring.yaml`), and both have a **definite, imminent home one step away**: the
v2.4 **release close**, the mandated last-chance re-fate gate. This is Fate-2 (an owning gate already holds them),
not `CHRONIC_DEFER` limbo. No new repeat pattern is introduced by M227.

## Fate-1 Investigation

### DEF-M227-01 — LOCAL live-render re-prove
- **Fate-1 (land now, complete) feasible:** no.
- **Why:** the live render requires a working full-stack demo bring-up; the local Docker box is environmentally
  wedged (ENOSPC + evicted build cache + buildx contention) and is recovering. Forcing it here would be a partial
  (data-only, which is ALREADY done deterministically) or an unreliable/blocked bring-up. The billion VM (clean,
  warm cache) is the correct venue.
- **Fate applied:** **Fate 2 (LAND-NEXT, already owned by M228).** Confirmed: M228's exit gate in `roadmap.md`
  explicitly covers the **corrected** data live — cond (2) "each candidate on exactly 1 sim, ≥ the retuned floor
  (≥6) per position", cond (3) "candidate profiles usable (external emails, matched avatars)", cond (4) "reads as
  hiring, **hiring-only content**". All four M227 believability fixes are M228's render-proof targets. No plan edit
  needed (M228 already owns it).

### DEF-CARRY-A / DEF-CARRY-B — inherited test-debt + declared TODO
- **Fate-1 feasible:** no (out of M227 scope; would require unrelated demo-stack / playthrough-WRITE work).
- **Fate applied:** **Fate 2 (LAND-NEXT → the v2.4 release close).** Re-confirmed fresh 2026-07-17, unchanged from
  M225/M226. The release close is the "extra-scrutiny last-chance" gate the audit-deferrals policy names for exactly
  these items.

### DEF-M226-01 — pre-bind serve reap (Finding-3)
- **Fate-1 feasible:** no (live-only surface; M227 is local/tooling-only and cannot re-prove it).
- **Fate applied:** **Fate 2/3 → M228** (the next prove-on-VM run) — self-resolves in the default flow;
  non-gate-blocking. Unchanged from M226.

## Recommendations
- DEF-M227-01 → **LAND-NEXT (Fate 2, M228)** — already owned; no plan edit.
- DEF-CARRY-A → **LAND-NEXT (Fate 2, v2.4 release close)** — re-confirmed.
- DEF-CARRY-B → **LAND-NEXT (Fate 2, v2.4 release close)** — re-confirmed.
- DEF-M226-01 → **LAND-NEXT (Fate 2/3, M228 / next prove-on-VM)** — re-confirmed.

## Applied Changes
- This report authored under `m227-the-notes/audit-deferrals/`.
- M227 `decisions.md` D5 already records DEF-M227-01's Fate-2 routing to M228 (no edit needed).
- No new milestone-plan edits: M228 already owns DEF-M227-01 (its exit gate covers the corrected-data render);
  the two inherited carries + DEF-M226-01 keep their standing destinations.

## Blocking Items (require user decision)
None. No M227-native repeat; the inherited repeats are consciously-tracked, out-of-scope, and one step from their
mandated re-fate gate (the release close). SEVERITY=warning (YELLOW).
