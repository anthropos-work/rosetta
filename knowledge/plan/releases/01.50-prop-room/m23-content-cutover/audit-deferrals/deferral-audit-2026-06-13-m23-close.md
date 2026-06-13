---
title: "Deferral Audit — milestone M23"
date: 2026-06-13
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- **No repeat deferrals.** No item appears in two distinct milestones' deferral ledgers unresolved.
- **Two inherited M21 items RESOLVED this milestone** (the strongest possible audit outcome): DEF-M21-03
  (directus_files ref capture) **landed Fate-1 in M23 §4**; DEF-M21-04 (referential closure / the 20 dangling
  relations) **landed Fate-1 in M23 §5** (the relations were subsumed by M21's 26-collection structure capture
  per M21-D7; the external content→taxonomy refs are now closed + MEASURED by the new cross-surface gene). Both
  drop off the ledger.
- **One genuine prod DATA residual surfaced + named** (`K-AIFUNX-E658`) — NOT a deferral of tooling work; the
  tooling completed its job (measure the dangling ref + name it). It's a prod data-quality inconsistency the
  operator owns; uncloseable by tooling without breaching the tenant firewall or editing prod (both forbidden).
  Fated below as a documented KNOWN-ISSUE, not a punt.
- **No chronic / aged-out patterns.** The remaining inherited items (DEF-M21-01, DEF-M21-02, DEF-M10-01) are all
  days-old (first recorded 2026-06-06..13), none crossed a *closed* destination milestone, none touched by M23.

## Summary
- Total deferrals in scope: 4 standing (2 inherited M21 follow-ups + 1 cross-release backlog + 1 new prod-data known-issue)
- Resolved this milestone: 2 (DEF-M21-03, DEF-M21-04 — both Fate-1 landed)
- Single deferrals: 4
- Repeat deferrals: 0
- Chronic / aged-out patterns flagged: 0

## Deferral Inventory

```yaml
# --- RESOLVED THIS MILESTONE (drop off the ledger) ---
- id: DEF-M21-03
  item: "directus_files ref capture — wire the dead media.go into directus.Surface() for real-image asset refs"
  origin_milestone: M21
  first_deferred_on: 2026-06-11
  destination: "M23 §4 (asset plane / content cutover)"
  m23_status: "RESOLVED — landed Fate-1, ext 2b8e9a0. New REFERENCED-SUBSET admissibility kind; firewall admits
    iff Filter==ReferencedFilesFilter; ClearByDelete DELETE-before-TRUNCATE for the directus_settings FK. Refs
    only (blob bytes stay DEF-M10-01). +13 tests. KB-1 resolved by §6. Drops off the ledger."

- id: DEF-M21-04
  item: "referential closure of the 20 dangling relations (content→taxonomy)"
  origin_milestone: M21
  first_deferred_on: 2026-06-11
  destination: "M23 §5 (referential closure, Fate-2)"
  m23_status: "RESOLVED — landed Fate-1, ext 4cb8786. The 20 relation defs were subsumed by M21's 26-collection
    structure capture (M21-D7); M23 owns the EXTERNAL content→taxonomy node-id refs, now closed by full-taxonomy
    capture (org_id IS NULL — already the state) + MEASURED by the new OpSnapshotCrossSurfaceClosure gene. Drops
    off the ledger."

# --- STANDING (inherited, unchanged, not aged) ---
- id: DEF-M21-01
  item: "replayCmd conn-seam — hermetic replayCmd wiring test needs an injectable connector"
  origin_milestone: M21
  first_deferred_on: 2026-06-13
  destination: "tooling-debt follow-up (replayCmd-seam build iter)"
  reason_recorded: "architectural, >50 lines; touches the load-bearing replay path"
  m23_status: "NOT touched by M23 — M23 grew replay.go with the ClearByDelete flag (a struct-field add, not a
    signature/seam refactor) so the conn-seam follow-up stands unchanged, not aged (5 days old)."

- id: DEF-M21-02
  item: "serve-live-integration harness — automated integration harness for serve-row render SQL"
  origin_milestone: M21
  first_deferred_on: 2026-06-13
  destination: "M25 field-bake (live observable-behavior gate)"
  reason_recorded: "needs a live directus Postgres; render is hermetically unit-tested + hand-validated live"
  m23_status: "Unchanged. M23's cross-surface gene is run against the REPLAYED directus↔skiller pair (live-DB),
    measure-snapshot is NOT in the bring-up path so it never blocks UP — consistent with the M25 live-gate home
    for the live-integration harness. Not aged."

- id: DEF-M10-01
  item: "S3 media blob BYTES + cloud SnapshotStore"
  origin_milestone: M10 (v1.2)
  first_deferred_on: 2026-06-06
  destination: "backlog (unscheduled)"
  reason_recorded: "user-facing sting removed by v1.5's real-images-via-prod-links posture; asset plane on prod"
  m23_status: "M23 explicitly stays REFS-ONLY (BlobBytesAvailable()==false, M23-D3) — directus_files captures
    the asset-ref UUIDs that resolve to prod public asset-plane URLs; the blob BYTES stay backlog by design.
    Re-affirmed by M23's posture, not re-deferred. Days/weeks old, orthogonal. No fresh decision required."

# --- NEW: prod DATA-quality residual (NOT a tooling deferral) ---
- id: M23-RESIDUAL-K-AIFUNX-E658
  item: "1 prod dangling content→taxonomy ref: 2 PUBLIC published sims reference K-AIFUNX-E658, which exists ONLY
    as a customer-scoped skill (org f9e88e97…)"
  origin_milestone: M23 (surfaced by the §5 closure gene; recorded M23-D5)
  first_seen_on: 2026-06-13
  destination: "documented KNOWN-ISSUE — operator-owned prod data correction, outside tooling scope"
  reason_recorded: "Uncloseable by tooling: capturing the customer node would breach the tenant firewall;
    editing prod is forbidden (zero platform/prod edits). The gene MEASURES + NAMES it rather than silently
    producing an empty picker — that IS the honest tooling resolution."
  partial_attempted: "n/a — not a tooling task"
```

## Repeat-Deferral Patterns
None. No item appears in two distinct milestones' ledgers unresolved. M23 carried nothing forward of its own;
it RESOLVED two inherited M21 items (DEF-M21-03, -04) in full and re-deferred none.

## Fate-1 Investigation

### M23-RESIDUAL-K-AIFUNX-E658 — the 1 prod dangling content→taxonomy ref
- **Fate-1 (close it in M23) feasible:** no — and "no" is correct, not a punt.
- **If no:** closing the dangle has only two mechanical resolutions, both forbidden by the release's hard
  constraints: (a) capture the customer-scoped node K-AIFUNX-E658 — a tenant-firewall BREACH (the firewall
  exists precisely to never capture customer data); (b) re-tag/remove the bad ref on the 2 public sims in prod —
  a PROD EDIT (v1.5 is zero platform/prod edits). The tooling's job was to make this VISIBLE not assumed, and it
  did: the gene measures the dangling count, fails non-blockingly at criticality `standard`, and names the
  sample node. The fix is a prod DATA correction the operator owns.
- **Fate:** **KNOWN-ISSUE (operator-owned), documented.** Recorded in M23-D5 + snapshot-spec.md. Not a release
  blocker, not an escape-hatch deferral (no tooling work is being punted — the tooling is complete).

### DEF-M21-01 — replayCmd conn-seam
- **Fate-1 feasible:** no — still architectural; M23 only added a struct field to replay.go, not the seam refactor.
- **Fate:** KEEP (tracked tooling-debt follow-up). Unchanged.

### DEF-M21-02 — serve-live-integration harness
- **Fate-1 feasible:** no — needs live Directus + CI infra; M25 field-bake home unchanged.
- **Fate:** LAND-NEXT (M25). Unchanged.

### DEF-M10-01 — S3 blob bytes + cloud store
- **Fate-1 feasible:** no — cross-release backlog; M23 stays refs-only by design (M23-D3), re-affirming the posture.
- **Fate:** KEEP-DEFERRED (release-level backlog, re-signed 2026-06-11). No fresh M23 decision required.

## Recommendations
- DEF-M21-03 → **RESOLVED (Fate-1, landed M23 §4)** — drops off the ledger.
- DEF-M21-04 → **RESOLVED (Fate-1, landed M23 §5)** — drops off the ledger.
- M23-RESIDUAL-K-AIFUNX-E658 → **KNOWN-ISSUE (operator-owned prod data correction)** — documented in M23-D5 +
  snapshot-spec.md; measured + named by the gene, not silently empty. Not a tooling deferral, not a blocker.
- DEF-M21-01 → **KEEP (tracked follow-up)** — unchanged; not unblocked or required by M23.
- DEF-M21-02 → **LAND-NEXT (M25 field-bake)** — unchanged.
- DEF-M10-01 → **KEEP-DEFERRED** — release-level backlog, re-signed 2026-06-11; M23 re-affirms refs-only.

## Applied Changes
- None required to plan files. The two resolved items (DEF-M21-03/-04) are already reflected as landed in M23
  progress.md §4/§5 + decisions.md (M23-D3/D4/D5). The K-AIFUNX-E658 known-issue is already recorded in
  M23-D5 + corpus/ops/snapshot-spec.md. The audit surfaces; nothing needed re-fating or re-routing.

## Blocking Items (require user decision)
None. Zero repeat deferrals, zero aged-out items. Two inherited items RESOLVED; the 1 prod-data residual is an
operator-owned KNOWN-ISSUE the tooling correctly measures + names. **GREEN — close proceeds.**
