---
milestone: M10
slug: content-snapshot
version: v1.2 "set dressing"
milestone_shape: section
status: archived
created: 2026-06-05
refined: 2026-06-06
last_updated: 2026-06-06
complexity: large
delivers: extends corpus/ops/snapshot-spec.md (Directus content path + store decision) + updates corpus/ops/seeding-spec.md (content surface promoted waived → snapshot-seeded)
---

# M10 — Directus content snapshot-replay

## Goal
Capture the **public** Directus content library (global simulation / skill-path templates) and replay it into a
**per-stack content store** — never touching shared Directus, never capturing customer content — taking data-DNA
coverage to **100% of the full catalog** (the last `waived` surface promoted to `snapshot-seeded` + fidelity-gated).

## Source correction (prod evidence, 2026-06-06)
Live-prod research sharpened M10's source — see
[`.agentspace/scratch/roadmap-research-2026-06-06.md`](../../../../../.agentspace/scratch/roadmap-research-2026-06-06.md):
- The app-Postgres `cms` schema is **not** the content source. It holds only `studio_documents` / `studio_tasks` /
  `similarities` (+ similarity join tables) — and `studio_documents` + `studio_tasks` are **100% org-scoped
  customer data (0 public rows)**, so they are **excluded entirely** (note #3). They were never the target.
- The **public content library** (global simulation / skill-path templates) lives in the **separate self-hosted
  Directus store** (`content.anthropos.work`, its own Postgres). That store, filtered to **public/global templates
  only**, is M10's real source.

## Why section (not iterative)
Once the per-stack content-store decision is made (an up-front architecture fork, not an emergent path), the
capture/replay reuses M9a's framework + the tenant firewall. The fork is the defining risk, resolved in a first
spike — a single decision, not a measurement-driven loop.

## Scope
- **In:**
  - The **per-stack content-store decision** resolved + built — the defining fork: a per-stack Directus container
    fed from the captured snapshot **vs** direct replay into the per-stack Directus-Postgres backing DB (Directus
    stores its data in Postgres → stays in the per-stack-isolated class). Resolve in the first iter/spike.
  - The **public content capture** — export the **public/global** Directus template collections + media references
    from the separate Directus store (a privileged read via M9a's capture-source policy + tenant firewall;
    public-templates-only, isolation-clean, audited).
  - The **content replay seeder** — wired into M9a's snapshot framework + the M7a seeder DAG, respecting the
    isolation guard (writes only to per-stack-isolated stores).
  - The **content fidelity + public-only genes** in the data-DNA (source-vs-replay conformance for the content surface).
  - The **`sim_id` / `skill_path_id` / `resource_id` linkage** so the v1.1 session/assignment seeders' content refs
    (currently free values, no FK) resolve against the real replayed **public** templates.
- **Out:**
  - App-Postgres `cms.studio_*` customer content (excluded — tenant data, note #3).
  - AI-generated / authored content (v1.3 — M10 replays *real* captured public content, it does not generate).
  - Recipes / presets (M11). External shareability (v1.3).

## Depends on
**M9a + M9b** (the snapshot framework + the fidelity DNA + the `stacksnap` CLI + the proven taxonomy surface).
**Parallel with:** none (M11 curates its output; the last build milestone before close).

## Open questions (resolve during build)
- The **content-store fork** (per-stack Directus container vs direct Directus-Postgres replay) — the load-bearing
  decision; resolve in the first spike.
- Where the public Directus templates physically live (the separate Directus Postgres) + how the public/global
  subset is identified (public flag vs a global org sentinel) — confirm against the Directus store at build.
- **Media/blobs** in-scope vs refs-only for the demo MVP — S3-private is per-stack-isolated, so blobs *can* be
  replayed; confirm at build.
- How much of the Directus collection set the demo needs — the believable subset (the M7c "reachable" discipline).

## KB dependencies (read as contract)
- `corpus/ops/snapshot-spec.md` (M9a's capture/replay contract + the capture-source policy + the tenant firewall)
- `corpus/services/cms.md` (Directus + the content model) + `corpus/ops/db-access.md` (the Directus store connection)
- `corpus/ops/seeding-spec.md` (the isolation guard + the session/assignment content refs)
- `corpus/services/jobsimulation.md` + `corpus/services/skillpath.md` (the consumers of `sim_id` / `skill_path_id`)

## Delivers → corpus/ops/snapshot-spec.md (extend) + corpus/ops/seeding-spec.md (update)
- **extends `corpus/ops/snapshot-spec.md`**: the public-Directus content path + the resolved per-stack
  content-store decision.
- **updates `corpus/ops/seeding-spec.md`**: the `content` surface promoted from `waived` to `snapshot-seeded`;
  data-DNA coverage now 100% of the full catalog.
</content>
