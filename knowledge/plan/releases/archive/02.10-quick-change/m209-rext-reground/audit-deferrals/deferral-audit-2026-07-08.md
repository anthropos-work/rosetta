---
title: "Deferral Audit — milestone:M209"
date: 2026-07-08
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat-deferrals in the scope-erosion sense; every item has a clear, in-release fate decision with its
  destination milestone's plan edit already applied.
- No aged-out items: every M209 deferral is fresh, today-dated (2026-07-08), routed to a still-open milestone
  that OWNS the work by release design; no destination has closed without landing.
- Inherited M208 deferrals (5, all Fate-2/3) re-checked — still GREEN, destinations (M210/M211) intact.

## Summary
- Total deferrals in scope (M209 + inherited M208): **10** (M209: **5** · inherited M208: **5**)
- Single deferrals: **10**
- Repeat deferrals: **0**
- Chronic patterns flagged: **0**
- Aged-out items requiring fresh decision: **0**

## Deferral Inventory (M209)

```yaml
- id: DEF-M209-01
  item: "corpus/ops/snapshot-spec.md describes the taxonomy surface as skiller.* (19 mentions) — STALE vs public.*"
  origin_milestone: M209
  first_deferred_on: 2026-07-08
  last_seen_in: decisions.md:9 (KB-1) / kb-fidelity-audit.md:26
  destination: "M210 (corpus body-flip) — Fate 2 (M210 overview names snapshot-spec.md, 26 mentions)"
  reason_recorded: "chartered lockstep body-flip; not load-bearing for M209's code re-ground (M209 grounds against M208's empirically-verified merged schema, not the doc bodies). Flipping here would collide with M210."
  partial_attempted: no

- id: DEF-M209-02
  item: "corpus/ops/seeding-spec.md prose ('public skiller catalog / taxonomy') — STALE-prose"
  origin_milestone: M209
  first_deferred_on: 2026-07-08
  last_seen_in: decisions.md:11 (KB-2) / kb-fidelity-audit.md:33
  destination: "M210 (corpus body-flip) — Fate 2 (M210 overview names seeding-spec.md)"
  reason_recorded: "same chartered lockstep; 0 skiller.<table> query refs currently, only narrative prose flips in M210."
  partial_attempted: no

- id: DEF-M209-03
  item: "corpus/ops/safety.md firewall evidence row 'skiller.skills 42,763 public' — STALE (schema + count, now ~42,790)"
  origin_milestone: M209
  first_deferred_on: 2026-07-08
  last_seen_in: decisions.md:12 (KB-3) / kb-fidelity-audit.md:30
  destination: "M210 (corpus body-flip) — Fate 2 (M210 overview names safety.md firewall row explicitly)"
  reason_recorded: "chartered lockstep; firewall CODE unchanged (schema-agnostic organization_id IS NULL)."
  partial_attempted: no

- id: DEF-M209-04
  item: "recapture of merged-prod public.* taxonomy into .agentspace/snapshots/ (the DATA op)"
  origin_milestone: M209
  first_deferred_on: 2026-07-08
  last_seen_in: decisions.md:39 / spec-notes.md:34
  destination: "M211 (bring-up acceptance) — Fate 3; M211 exit gate already names it (Fate-2 coverage) + a pinned 'Pre-surfaced recapture prerequisite' section"
  reason_recorded: "the tooling CODE re-ground is complete (capture/replay ready for public.*, digest narrowed, MinRows floor, column list verified). The actual recapture is operationally gated: values-blind-confirmed NO marco_read/prod-read Postgres DSN in .agentspace/secrets or platform/.env; merged stack-dev PG holds 0 taxonomy rows; the postgres MCP returns JSON not COPY bytes. Investigated Fate-1 (a real capture) — genuinely infeasible here. A data prerequisite M211's bring-up owns."
  partial_attempted: no

- id: DEF-M209-05
  item: "a handful of CONCEPTUAL bare-word 'skiller' code comments/labels remain (seeder/dna narrative — 'resolves in skiller', the ↔skiller closure-probe label)"
  origin_milestone: M209
  first_deferred_on: 2026-07-08
  last_seen_in: decisions.md:48 (scope-boundary note)
  destination: "M210 (doc-narrative) — Fate 2"
  reason_recorded: "the done-bar is '0 skiller.<table> QUERIES in production' — VERIFIED 0 (dotted swap + grep). These remaining mentions are conceptual narrative, not schema-existence claims and not queries; the factually-WRONG 'missing skiller schema' comments WERE fixed. Rewording concept-narrative is M210 territory."
  partial_attempted: no
```

**Non-deferral (recorded for completeness, carries no deferral authority):** a live-DB integration test of
`pg.Conn.SchemaVersion` was **ruled out on merits** — the uncovered `pg.Conn` remainder is the DB-integration
layer (integration-only by design); a DB-mutating test contradicts the read-only tooling ethos; the narrowing
behavior is proven DB-free (SchemaVersionSQL builder 100% covered). Recorded in progress.md stop-condition +
decisions.md. This is a design stance, not a deferred scope item.

## Repeat-Deferral Patterns
**None.** M209's KB-1/2/3 (rext-facing ops/spec docs: snapshot-spec.md, seeding-spec.md, safety.md) are
**distinct files** from M208's DEF-M208-03/04/05 (architecture/service docs: backend.md, skiller.md,
graphql-wundergraph.md). Both sets route to **M210** — but M210 is a **dedicated, chartered milestone whose
entire job is the corpus body-flip**, not a repeated punt of the same item. This is the release's designed
sequential structure (M208 ground-truth → M209 rext → **M210 corpus** → M211 bring-up), a textbook Fate-2
("a future milestone of this release already owns it"), NOT scope erosion. No item appears twice; no
CHRONIC_DEFER / DRIFT_DEFER.

## Aging
No AGED_OUT items. Every M209 deferral is dated 2026-07-08 with a fresh fate decision. The destination
milestones (M210, M211) are both still `planned` and OPEN — M210 is literally the next milestone. No
"deferred-to milestone closed without landing" trigger fires. The recapture (DEF-M209-04) is a first-time,
concretely-scoped assignment with a pre-surfaced-prerequisite section in M211's overview, not a re-punt.

## Fate-1 Investigation

### DEF-M209-01/-02/-03/-05 — corpus doc-body / narrative staleness → M210
- **Fate-1 (land now, complete) feasible:** no (by design)
- **If no:** Fate 2 → M210. Flipping the corpus bodies at M209 would collide with M210's chartered lockstep
  deliverable and land doc changes ahead of / entangled with the tooling they must stay in lockstep with
  (roadmap Risk 4). M210 explicitly enumerates each of these files. M209's charter is the CODE re-ground; the
  done-bar (0 `skiller.<table>` queries) is met.

### DEF-M209-04 — recapture the public taxonomy → M211
- **Fate-1 feasible:** no — genuinely investigated. No valid COPY-byte capture source is provisioned locally
  (values-blind-checked). The tooling code is READY; the missing piece is a data source, which is a bring-up-time
  operational concern M211 owns.
- **If no:** Fate 3 → M211 (its exit gate already names the recapture — Fate-2 coverage — plus a pinned
  "Pre-surfaced recapture prerequisite" section so M211's first tik doesn't re-discover the blocker).

## Recommendations
| Item | Fate | Destination | Verb |
|---|---|---|---|
| DEF-M209-01 | Fate 2 | M210 | LAND-NEXT |
| DEF-M209-02 | Fate 2 | M210 | LAND-NEXT |
| DEF-M209-03 | Fate 2 | M210 | LAND-NEXT |
| DEF-M209-04 | Fate 3 | M211 | LAND-NEXT |
| DEF-M209-05 | Fate 2 | M210 | LAND-NEXT |

No LAND-NOW (none feasible as a complete Fate-1 within M209's charter — the code re-ground IS the charter and it
landed complete), no DROP, no KEEP-DEFERRED-WITH-SIGNOFF (zero cross-release / escape-hatch punts).

## Applied Changes
No new plan edits required by this audit — every route was already applied during the M209 build and is confirmed
intact:
- M210 `overview.md` — enumerates snapshot-spec.md / seeding-spec.md / safety.md body-flips — present.
- M211 `overview.md` — "Pre-surfaced recapture prerequisite" section (DEF-M209-04 Fate-3 landing) + corrected
  ~42,790 count — present.
- M209 `decisions.md` — KB-1/2/3, recapture, and scope-boundary records — present.
This report is the audit artifact; it records the confirmations, it does not re-open settled routes.

## Blocking Items (require user decision)
None. Zero repeat-deferrals, zero aged-out items with revoked authority, zero escape-hatch punts. The close may
proceed.
