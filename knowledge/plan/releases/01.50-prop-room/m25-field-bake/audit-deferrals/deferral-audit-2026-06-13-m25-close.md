---
title: "Deferral Audit — milestone M25"
date: 2026-06-13
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- **DEF-M21-02 (serve-live-integration harness) RESOLVED in M25 — Fate-1, by observation.** This is the
  strongest possible outcome for the one inherited item whose confirmed Fate-2 destination *was* this milestone.
  M21 routed it to M25 as "the live observable-behavior gate"; M25 delivered exactly that — the local Directus
  was driven against a real structure-bearing capture and **observably serves** (DB-1 demo-1 offset 18055 +
  DB-2 dev-2 offset 28055, both curl-proven: `/server/health` 200, `/items/simulations` returns real published
  rows, asset URL resolves to `content.anthropos.work`). Evidence: `db1-serve-evidence.txt`, `db2-serve-evidence.txt`.
  The serve-row render SQL is now exercised against live Directus Postgres — precisely the live integration
  DEF-M21-02 needed. It **drops off the ledger.**
- **The other two standing items are unchanged and not aged by M25.** DEF-M21-01 (replayCmd conn-seam) and
  DEF-M10-01 (S3 blob bytes) stand exactly as the M24-close audit fated them; M25 touched neither code area
  (M25's ext fixes are in `firewall`/`directus`/`media`/`cmd/stacksnap` capture paths + `demo-stack`
  `GOTOOLCHAIN` — not `replay.go`'s conn-seam nor the asset-plane store).
- **Two M25-originated minor deferrals, both environmental, both fated this pass (single, not repeat, not aged).**
  M25-DEF-01 (full-UI Playwright render) and M25-DEF-02 (dev-2 taxonomy `rc=4`) — neither is a v1.5 tooling
  defect; both have the core behavior proven by an alternate path. Fated below; recorded fresh in `decisions.md`
  (M25-D8, M25-D9). No repeat pattern (first appearance), no chronic flag.

## Summary
- Total deferrals in scope: 3 standing inherited (DEF-M21-01, DEF-M21-02, DEF-M10-01) + 2 M25-originated
- Resolved this milestone: 1 (DEF-M21-02 — Fate-1 by live observation; drops off the ledger)
- New deferrals from M25: 2 (M25-DEF-01, M25-DEF-02 — both single, both fated this pass)
- Single deferrals: 4 (the 2 remaining standing + the 2 new)
- Repeat deferrals: 0
- Chronic / aged-out patterns flagged: 0

## Deferral Inventory

```yaml
# --- RESOLVED THIS MILESTONE (drops off the ledger) ---
- id: DEF-M21-02
  item: "serve-live-integration harness — exercise the serve-row render SQL against a live Directus Postgres"
  origin_milestone: M21
  first_deferred_on: 2026-06-13
  destination: "M25 field-bake (live observable-behavior gate) — its confirmed Fate-2 home"
  reason_recorded: "needs a live directus Postgres; render is hermetically unit-tested + hand-validated live"
  m25_status: "RESOLVED Fate-1. M25 drove the serve-row render against a real structure-bearing capture on a
    live per-stack Directus; it observably serves (DB-1 + DB-2 curl-proven). Drops off the ledger."
  partial_attempted: no

# --- STANDING (inherited, unchanged by M25, not aged) ---
- id: DEF-M21-01
  item: "replayCmd conn-seam — hermetic replayCmd wiring test needs an injectable connector"
  origin_milestone: M21
  first_deferred_on: 2026-06-13
  destination: "tooling-debt follow-up (replayCmd-seam build iter)"
  reason_recorded: "architectural, >50 lines; touches the load-bearing replay path"
  m25_status: "NOT touched by M25 — M25's ext fixes are firewall/directus/media/cmd capture paths + demo-stack
    GOTOOLCHAIN; replay.go conn-seam untouched. Stands unchanged, not aged (no closed destination milestone)."
  partial_attempted: no

- id: DEF-M10-01
  item: "S3 media blob BYTES + cloud SnapshotStore"
  origin_milestone: M10 (v1.2)
  first_deferred_on: 2026-06-06
  destination: "backlog (unscheduled)"
  reason_recorded: "asset plane on prod (real images via prod links); blob BYTES stay backlog by design"
  m25_status: "Unchanged. M25's DB-1 evidence re-confirms the refs-only posture works as designed — the served
    sim's cover resolves to content.anthropos.work (asset plane prod), data plane local. No fresh decision."
  partial_attempted: no

# --- M25-ORIGINATED (new this milestone, single, fated this pass) ---
- id: M25-DEF-01
  item: "full-UI Playwright render — next-web/studio-desk rendered-page screenshot of the local-Directus catalog"
  origin_milestone: M25
  first_deferred_on: 2026-06-13
  destination: "backlog (environmental — VM budget), NOT a v1.5 tooling deliverable"
  reason_recorded: "the ~10 GiB practical VM ceiling on this 16 GB host (M25-D2) can't co-host the full UI tier
    (next-web build spikes ~3.7 GB) with a backend stack; DB-1/DB-2 ran --no-ui. The core observable behavior
    (local Directus SERVES the catalog; asset plane stays prod) is curl-proven against the exact surface a
    browser calls (cms + per-stack Directus)."
  partial_attempted: "n/a — the rendered-page tier is a host-budget constraint, not a tooling gap"

- id: M25-DEF-02
  item: "dev-2 taxonomy replay rc=4 (target schema empty) — a dev migrate-ordering nuance"
  origin_milestone: M25
  first_deferred_on: 2026-06-13
  destination: "tooling-debt follow-up (dev migrate-ordering), NOT a v1.5 deliverable"
  reason_recorded: "non-fatal, unrelated to the content-serve path. DB-2's directus content-serve (the core
    done-bar) is GREEN; the rc=4 is taxonomy-surface migrate-ordering on the opt-in dev-2 stack only — the
    directus replay exits 0 and serves. Pre-existing dev-stack ordering nuance, not introduced by v1.5 content work."
  partial_attempted: "n/a — orthogonal to the milestone's content-serve charter"
```

## Repeat-Deferral Patterns
None. No item appears unresolved in two distinct milestones' ledgers. DEF-M21-02 RESOLVED in its destination
milestone (the ideal). The two M25-originated items are first-appearance. DEF-M21-01 and DEF-M10-01 stand with
intact, never-re-routed destinations.

## Fate-1 Investigation

### DEF-M21-02 — serve-live-integration harness
- **Fate-1 (land now in M25) feasible:** YES — and landed. M25 *is* the live observable-behavior gate. Driving
  the serve-row render against a real structure-bearing capture on a live per-stack Directus, then proving the
  served rows by curl (DB-1 + DB-2), is the live integration this item needed. The render SQL went from
  hermetic-unit + hand-validated to observably-serving-real-prod-shape.
- **Fate:** RESOLVED (Fate-1). Drops off the ledger.

### DEF-M21-01 — replayCmd conn-seam
- **Fate-1 (land now in M25) feasible:** no — M25 is a live field-bake, not a refactor milestone; the conn-seam
  is a >50-line architectural change to the load-bearing replay path, not touched or unblocked by M25.
- **Fate:** KEEP (tracked tooling-debt follow-up). Unchanged.

### DEF-M10-01 — S3 blob bytes + cloud store
- **Fate-1 (land now in M25) feasible:** no — cross-release backlog, orthogonal to a field-bake. M25's DB-1
  evidence re-confirms the deliberate refs-only posture serves correctly (asset plane = prod links).
- **Fate:** KEEP-DEFERRED (release-level backlog, re-signed 2026-06-11). No fresh decision required.

### M25-DEF-01 — full-UI Playwright render
- **Fate-1 (land now in M25) feasible:** no — this is a **host-budget** constraint (M25-D2: 12 GB VM doesn't
  boot on this 16 GB box; ~10 GiB practical max), not a tooling defect. The full UI tier can't be co-resident
  with a backend stack on this box. Landing it "now" would require different hardware, not different code.
  Critically, the **core observable behavior is already proven** — the local Directus serves the catalog and the
  asset plane stays prod, curl-verified against cms + the per-stack Directus (the surface a browser calls). The
  rendered-page screenshot would add presentation confirmation, not behavioral confirmation.
- **Fate:** DROP as a v1.5 deliverable / KEEP-as-environmental-backlog. Not a v1.5 tooling commitment — the
  done-bar "the browser shows content served by the local Directus" is satisfied at the behavior level by the
  data-plane curl proof (M25-D2 records this substitution explicitly, accepted at build time). Recorded fresh
  as **M25-D8**. No cross-release tooling work is owed; if a full-UI render is ever wanted it needs a bigger
  box, not a v1.5 code change.

### M25-DEF-02 — dev-2 taxonomy rc=4
- **Fate-1 (land now in M25) feasible:** no — it's a pre-existing dev-stack **migrate-ordering** nuance on the
  taxonomy surface of the opt-in dev-2 stack, **unrelated to the content-serve path** M25 charters. DB-2's
  directus content-serve (the done-bar) is GREEN; the directus replay exits 0 and serves. Diagnosing/fixing the
  taxonomy migrate-ordering is dev-stack tooling debt outside the field-bake's content-serve scope, and is
  non-fatal (the stack comes up and serves).
- **Fate:** KEEP (tracked tooling-debt follow-up, dev migrate-ordering). Recorded fresh as **M25-D9**. Not a
  v1.5 content-release deliverable; surfaced so it isn't lost.

## Recommendations
- DEF-M21-02 → **RESOLVED (Fate-1, landed M25 — live serve proven)** — drops off the ledger.
- DEF-M21-01 → **KEEP (tracked follow-up)** — unchanged; not unblocked by M25.
- DEF-M10-01 → **KEEP-DEFERRED** — release-level backlog, re-signed 2026-06-11; refs-only re-confirmed by DB-1.
- M25-DEF-01 → **DROP-as-v1.5-deliverable / environmental backlog** — behavior proven by curl; render is a
  host-budget item, not a tooling commitment. Fresh decision M25-D8.
- M25-DEF-02 → **KEEP (tracked dev migrate-ordering follow-up)** — non-fatal, orthogonal to content-serve.
  Fresh decision M25-D9.

## Applied Changes
- `decisions.md` — two fresh decisions recorded for the M25-originated items: **M25-D8** (full-UI Playwright
  render → environmental backlog, behavior proven by the curl data-plane evidence) and **M25-D9** (dev-2
  taxonomy rc=4 → tracked dev migrate-ordering follow-up, non-fatal, orthogonal to the content-serve charter).
- No plan-file edits required for the standing items — DEF-M21-02 is reflected as resolved by the DB-1/DB-2
  evidence + progress.md; DEF-M21-01 and DEF-M10-01 are already correctly recorded by the M24-close audit and
  need no re-fating. The audit surfaces; the two new decisions are the only ledger writes.

## Blocking Items (require user decision)
None. DEF-M21-02 RESOLVED Fate-1 in its destination milestone (the ideal). The two M25-originated items are
single (first-appearance), environmental, with the core behavior proven by an alternate path — fated this pass
with fresh decisions (M25-D8 DROP-as-deliverable, M25-D9 KEEP-follow-up). Zero repeat, zero chronic, zero
aged-out. **GREEN — close proceeds.**
