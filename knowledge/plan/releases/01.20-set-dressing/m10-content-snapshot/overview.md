---
milestone: M10
slug: content-snapshot
version: v1.2 "set dressing"
milestone_shape: section
status: planned
created: 2026-06-05
complexity: large
delivers: extends corpus/ops/snapshot-spec.md (Directus content path + store decision) + updates corpus/ops/seeding-spec.md (content surface promoted waived → snapshot-seeded)
---

# M10 — Directus content snapshot-replay

## Goal
Capture the shared-Directus content library and replay it into a **per-stack content store** — never touching
shared Directus — taking data-DNA coverage to **100% of the full catalog** (the last `waived` surface promoted to
`snapshot-seeded` + fidelity-gated).

## Why section (not iterative)
Once the per-stack content-store decision is made (an up-front architecture fork, not an emergent path), the
capture/replay is mechanical and reuses M9's framework. The fork is the defining risk, resolved in a first
spike — but it's a single decision, not a measurement-driven loop. Section with an early decision spike.

## Scope
- **In:**
  - The **per-stack content-store decision** resolved + built — the defining fork: a per-stack Directus container
    fed from the captured snapshot **vs** direct replay into the per-stack Directus-Postgres backing DB (Directus
    stores its data in Postgres → stays in the per-stack-isolated class). Resolve in the first iter/spike.
  - The **content capture** — export shared-Directus collections + media references (a privileged read,
    isolation-clean, audited).
  - The **content replay seeder** — wired into M9's snapshot framework + the M7a seeder DAG, respecting the
    isolation guard (writes only to per-stack-isolated stores).
  - The **content fidelity gene** in the data-DNA (source-vs-replay conformance for the content surface).
  - The **`sim_id` / `skill_path_id` / `resource_id` linkage** so the v1.1 session/assignment seeders' content
    refs (currently free values, no FK) resolve against real replayed content.
- **Out:**
  - AI-generated / authored content (v1.3 — M10 replays *real* captured content, it does not generate).
  - Recipes / presets (M11).
  - External shareability (v1.3).

## Depends on
**M9** (the snapshot framework + the fidelity DNA + the `stacksnap` CLI it reuses). **Parallel with:** none
(M11 curates its output; the last build milestone before close).

## Open questions (resolve during build)
- The **content-store fork** (per-stack Directus container vs direct Directus-Postgres replay) — the load-bearing
  decision; resolve in the first spike.
- **Media/blobs** in-scope vs refs-only for the demo MVP — S3-private is per-stack-isolated, so blobs *can* be
  replayed; confirm at build.
- How much of the Directus collection set the demo needs — the believable subset (the M7c "reachable" discipline).

## KB dependencies (read as contract)
- `corpus/ops/snapshot-spec.md` (M9's capture/replay contract + format)
- `corpus/services/cms.md` (Directus + the `cms` schema — Similarity/StudioDocument/StudioTask)
- `corpus/ops/seeding-spec.md` (the isolation guard + the session/assignment content refs)
- `corpus/services/jobsimulation.md` + `corpus/services/skillpath.md` (the consumers of `sim_id` / `skill_path_id`)

## Delivers → corpus/ops/snapshot-spec.md (extend) + corpus/ops/seeding-spec.md (update)
- **extends `corpus/ops/snapshot-spec.md`**: the Directus content path + the resolved per-stack content-store
  decision.
- **updates `corpus/ops/seeding-spec.md`**: the `content` surface promoted from `waived` to `snapshot-seeded`;
  data-DNA coverage now 100% of the full catalog.
