---
title: "Deferral Audit — milestone M210 (Corpus + skills re-ground)"
date: 2026-07-08
scope: milestone
invoked-by: close-milestone
---

## Verdict
**GREEN**

- No repeat deferrals; no chronic pattern; no aged-out item requiring a fresh block.
- Every open item is a **single**, freshly-fated deferral with a concrete in-release home.
- The four M210-destined defers (KB-1/2/3 + the M209 body-flip set) **LANDED at their destination this milestone** — they come OFF the open list as COMPLETED, not still-open.

## Summary
- Total deferrals in scope (v2.1, milestones M208→M210 + M211-planned): **11 recorded across the release**
- Resolved-at-destination by M210 (now closed): **7** (DEF-M208-03/04/05 + DEF-M209-01/02/03/05)
- Still-open (confirmed routings): **4** (2× → M211, 1× → M211 exit-gate, 1× → close-release rext roll)
- Single deferrals: **11** · Repeat: **0** · Chronic: **0** · Aged-out: **0** · Escape-hatch: **0**
- M210-originated deferrals: **0** (every M210 scope item landed Fate-1)

## Deferral Inventory

### Resolved-at-destination by M210 (come OFF the open list)
| id | item | origin | destination | resolution |
|---|---|---|---|---|
| DEF-M208-03 | backend.md pre-merge consumer prose | M208 | Fate-2 → M210 | LANDED — M210 §1 (adopted skiller-in-app arch/service half; reconciled backend.md vs the M208 fact-sheet) |
| DEF-M208-04 | skiller.md standalone body | M208 | Fate-2 → M210 | LANDED — M210 §1 (skiller.md reframed as merged/legacy stub → backend.md) |
| DEF-M208-05 | 5→4 subgraphs in CLAUDE.md / graphql-wundergraph.md | M208 | Fate-2 → M210 | LANDED — M210 §5/§6 (skill files + CLAUDE.md catalog; 0 "5 subgraphs" leftover) |
| DEF-M209-01 | snapshot-spec.md skiller.* body-flip | M209 | Fate-2 → M210 | LANDED — M210 §3 |
| DEF-M209-02 | seeding-spec.md prose | M209 | Fate-2 → M210 | LANDED — M210 §3 |
| DEF-M209-03 | safety.md firewall evidence row | M209 | Fate-2 → M210 | LANDED — M210 §3 |
| DEF-M209-05 | conceptual bare-word skiller comments/labels (incl. profile-completeness node-id prose) | M209 | Fate-2 → M210 | LANDED — M210 §2/§3 (profile-completeness node-id prose fixed to "taxonomy node-ids … in the public schema") |

### Still-open (confirmed, single, freshly-fated)
```yaml
- id: DEF-M208-01
  item: "clean cold reset-db does not bootstrap the extensions schema (pgvector/pg_trgm/gin_trgm_ops) before migrate + a PG-readiness race (the M25-D9 class)"
  origin_milestone: M208
  first_deferred_on: 2026-07-08
  last_seen_in: m208/retro.md:53, m211/overview.md (In: pinned)
  destination: "Fate-3 -> M211 (overview pinned) + M209 Risk-2 cross-ref"
  reason_recorded: "did NOT fall out as a trivial Fate-1 on the re-migrate path; a bring-up-tooling requirement, correctly re-fated to M211"
  partial_attempted: no
- id: DEF-M208-02
  item: "stack-dev hand-assembled .env lacks INVITATION_HMAC_SECRET (backend Exited(0) on cold containerized run)"
  origin_milestone: M208
  first_deferred_on: 2026-07-08
  last_seen_in: m208/retro.md:54
  destination: "Fate-2 -> M211 / /stack-secrets"
  reason_recorded: "per-stack .env completeness gap, not merge-caused"
  partial_attempted: no
- id: DEF-M209-04
  item: "recapture the public.* taxonomy from merged-prod into .agentspace/snapshots (bump capture version)"
  origin_milestone: M209
  first_deferred_on: 2026-07-08
  last_seen_in: m209/metrics.json, m211/overview.md (exit-gate)
  destination: "Fate-3 -> M211 (tooling READY; no local COPY-byte capture source -> a data op M211 owns)"
  reason_recorded: "the tooling is re-grounded + tagged; capture needs a sanctioned COPY-byte source that only the M211 bring-up provides"
  partial_attempted: no
- id: TEST-1  # M209 D-close-2
  item: "rext stack-seeding/README.md test-count drift (says 496/8 pkgs; actual ~788/13)"
  origin_milestone: M209
  first_deferred_on: 2026-07-08
  last_seen_in: m209/metrics.json close_review, state.md standing-backlog
  destination: "v2.1 rext roll (close-release) / M211 rext re-tag"
  reason_recorded: "PRE-EXISTING since M41 (v1.10); M209 did not touch the README; rext deliberately FROZEN at 2f06e78 for the M209 close -> reconcile at the rext roll, not mid-release"
  partial_attempted: no
```

### Non-deferrals (recorded for completeness — not on the ledger)
- **Live-DB integration test of `pg.Conn.SchemaVersion` (M209)** — ruled out **on merits** (the fake-Conn matcher suite covers the contract; a live-DB test adds no signal a bring-up gate doesn't already give). A design decision, not a backlog punt. Not a deferral.
- **Push-gated KEEP** (origin push of `main` + tags + rext consumption re-pin) — an **administrative** user gate (standing push-gate), explicitly "an administrative KEEP, not a deferral" per state.md. Not a deferral.
- **Cross-release standing backlog** (DEF-M10-01 cloud SnapshotStore, DEF-M21-01 replayCmd hermetic test, M314b prod frozen-read hydration) — tracked in `roadmap-vision.md`, predate v2.1, out of milestone scope.

## Repeat-Deferral Patterns
**None.** No item was deferred across ≥2 distinct v2.1 milestones without resolution. The M210-destined set (KB-1/2/3 + the M209 body-flip defers) was fated once (at M208/M209) and **landed at its destination this close** — the definition of a deferral working correctly, not a repeat. M25-D9 (DEF-M208-01) traces to a v1.x decision but received its **first concrete v2.1 routing** at M208 (Fate-3 → M211); it is a single active routing, not a repeat-without-resolution.

## Aging
**No aged-out items.** Triggers checked per item:
- *Deferred across ≥2 milestones* — none (each open item fated exactly once).
- *Deferred ≥3 months ago* — all four open items fated **2026-07-08** (this week). TEST-1's underlying drift is old (M41 ≈ 2026-06-27, ~2 weeks) but the **deferral decision** is fresh (M209, conscious rext-freeze) — under the 3-month line.
- *Destination milestone closed without landing* — M211 has not started; close-release has not run. No trigger.
- *Area touched substantively by a later milestone* — M210 touched the corpus and **resolved** KB-1/2/3 (a landing, not a stale-context problem); the rext README (TEST-1) area was untouched (rext frozen). No trigger.

## Fate-1 Investigation (still-open items)
- **DEF-M208-01 / M25-D9** — Fate-1 now feasible: **no**. It is bring-up-tooling work (extensions-schema bootstrap + PG-readiness) that only manifests on a live cold stand-up; M210 is a docs-only milestone with no bring-up surface. Correct home: **M211** (owns it in `overview.md`).
- **DEF-M208-02 / INVITATION_HMAC_SECRET** — Fate-1 now feasible: **no**. A per-stack `.env`/secret-provisioning gap; belongs to `/stack-secrets` + the M211 bring-up. M210 touches no `.env`.
- **DEF-M209-04 / recapture** — Fate-1 now feasible: **no**. Requires a sanctioned COPY-byte capture source that only the M211 live bring-up provides; it is a data operation, not a docs edit. M211 exit-gate owns it.
- **TEST-1 / rext README drift** — Fate-1 now feasible: **no** (in this milestone). The count lives in the **rosetta-extensions** repo, which was deliberately frozen at `2f06e78` for the M209 close; touching it now would re-open a tagged rext. Correct home: the **v2.1 rext roll at close-release**.

## Recommendations
| item | verdict |
|---|---|
| DEF-M208-01 (M25-D9 extensions bootstrap) | **LAND-NEXT** (Fate-3, M211 — already pinned; confirm, no edit) |
| DEF-M208-02 (INVITATION_HMAC_SECRET) | **LAND-NEXT** (Fate-2, M211 / /stack-secrets — confirm) |
| DEF-M209-04 (recapture public.* taxonomy) | **LAND-NEXT** (Fate-3, M211 exit-gate — confirm) |
| TEST-1 (rext README drift) | **LAND-NEXT** (Fate-2, close-release rext roll — confirm; rext frozen this close) |

All four are **confirm-only** (Fate-2/Fate-3 already owned by a downstream milestone or close-release). No plan edits required; no LAND-NOW that M210's docs-only scope could carry; no DROP; no escape-hatch.

## Applied Changes
None. Every still-open item is a **confirmation** of an existing, correctly-recorded routing to a downstream milestone (M211) or close-release; no `overview.md` edit, no new decision, no roadmap mutation needed. The four M210-destined defers were resolved by the milestone's own commits (§1/§2/§3/§5/§6), already recorded in `progress.md` + `decisions.md`.

## Blocking Items (require user decision)
**None.** No repeat-deferral, no aged-out item, no escape-hatch. Gate returns **SEVERITY=clear**.
