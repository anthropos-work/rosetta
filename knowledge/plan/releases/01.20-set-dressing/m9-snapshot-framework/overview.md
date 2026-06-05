---
milestone: M9
slug: snapshot-framework
version: v1.2 "set dressing"
milestone_shape: section
status: planned
created: 2026-06-05
complexity: large
delivers: corpus/ops/snapshot-spec.md (new) + extends corpus/architecture/alignment_testing.md (snapshot-fidelity dimension)
---

# M9 — Snapshot framework + fidelity DNA + taxonomy snapshot

## Goal
A generic **capture → serialize → replay** mechanism for *snapshot-class* surfaces (reference data captured once
from a source, replayed per-stack), an **alignment extension that measures replay fidelity** (source-vs-replay),
and the first surface — **taxonomy** — proving it end-to-end, driving data-DNA coverage from `waived` to its
first snapshot-seeded surface.

## Why section (not iterative)
The deliverables are writable up front: the snapshot contract + format, the isolation rules for capture/replay,
the data-DNA fidelity-gene class, and one known surface (taxonomy, Postgres→Postgres). The fidelity gate is a
per-surface acceptance check (does the replay match the captured source?), not an emergent-path gate. This is
framework work like M7a+M7b — decomposable into a fixed checklist.

## Scope
- **In:**
  - The **snapshot contract + portable format** — a capture manifest + payload describing how a surface is
    serialized, versioned, and pinned per stack (the M0 record/replay discipline applied to data).
  - The **snapshot-store abstraction** honoring the isolation boundary: capture = a *privileged read* from a
    source; replay writes **only** to per-stack-isolated stores. Extends M7a's `isolation/` (`AssertClean` must
    cover snapshot replay; capture is read-only + audited).
  - The **data-DNA extension** — a new surface status `snapshot-seeded` + a **snapshot-fidelity gene class**
    (source-vs-replay row-count / structural conformance / referential integrity / **embedding-dimension
    integrity**) added to the `datadna` harness.
  - The **`stacksnap` CLI** (`capture` / `replay`) + `datadna` recognizing snapshot surfaces.
  - The **taxonomy snapshot seeder** — Postgres→Postgres capture + bulk-`COPY` replay of skiller
    categories/specializations/skills/roles + pgvector embeddings, gated on its fidelity gene. Drives data-DNA
    coverage waived→taxonomy-seeded.
- **Out:**
  - The Directus content snapshot (M10).
  - Recipes / presets / corpus product layer (M11).
  - AI-generated content + external shareability (v1.3).

## Depends on
v1.1's **M7a** (isolation guard + the `COPY` perf path + the seeder DAG) + **M7b** (the data-DNA harness this
extends). **Parallel with:** none (gates M10 + M11).

## Open questions (resolve during build)
- The snapshot **source**: a captured golden committed to the extensions repo vs a live privileged read from a
  reference DB at capture-time → pick the reproducible, offline-replayable path (lean: committed golden, M0
  discipline; the data-DNA `diff` flags schema drift on a bump).
- **Embedding capture fidelity**: carry pgvector vectors verbatim vs recompute → verbatim (offline +
  deterministic; recompute is AI-content, v1.3). Gate with the embedding-dimension fidelity gene.
- Snapshot **versioning** when the skiller schema drifts (the data-DNA `diff` already flags it — confirm the
  refresh workflow).

## KB dependencies (read as contract)
- `corpus/ops/seeding-spec.md` (the seeding framework + the 3-layer isolation boundary + the perf path)
- `corpus/architecture/alignment_testing.md` (the M0 alignment discipline + M7b's data dimension to extend)
- `corpus/services/skiller.md` (the taxonomy schema: categories/specializations/skills/roles + pgvector embeddings)
- `corpus/ops/staging_from_dump.md` (the snapshot precedent — pg_dump restore + isolation-rebind/verify patterns)

## Delivers → knowledge/corpus/ops/snapshot-spec.md (new) + corpus/architecture/alignment_testing.md
- **`corpus/ops/snapshot-spec.md`** (net-new): the capture/replay contract, the portable format, the isolation
  rules for snapshot-class surfaces, per-stack injection + versioning.
- **extends `corpus/architecture/alignment_testing.md`**: the snapshot-fidelity dimension (source-vs-replay gene
  class) alongside the behavioral (v1.0) + structural data-DNA (v1.1 M7b) dimensions.
