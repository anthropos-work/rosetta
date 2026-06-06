---
milestone: M11
slug: richer-recipes
version: v1.2 "set dressing"
milestone_shape: section
status: planned
created: 2026-06-05
complexity: medium
delivers: refreshes corpus/ops/demo/ (recipes + presets to full-fidelity) + the /demo-snapshot skill + the CLAUDE.md skill table
---

# M11 — Richer-world recipes, presets + corpus polish

## Goal
The product/discoverability layer that closes v1.2 (the M8-analog): make the full-fidelity worlds *usable +
discoverable* — refresh presets + recipes so a demo curator gets a real-taxonomy, real-content world out of the
box, and update the corpus to reflect 100% coverage.

## Why section (not iterative)
Pure curation/discoverability over M9a/M9b/M10's shipped surfaces — a finite checklist (presets, recipes, a skill,
corpus cross-links). Same shape as v1.1's M8.

## Scope
- **In:**
  - Refreshed **seed presets** (small/mid/large) that now include the taxonomy + content snapshots.
  - An updated **`corpus/ops/demo/` recipe family** — the end-to-end recipes now showcase a *set-dressed* world:
    real skills in the catalog, real simulations/skill-paths behind the seeded sessions.
  - A **`/demo-snapshot` skill** (or an extended `/demo-seed`) for capture/replay.
  - Cross-linking + corpus updates (the data-DNA now reads **100%**); the CLAUDE.md skill table.
  - **Release-close hygiene** carry — any small items surfaced in M9/M10.
- **Out:**
  - New snapshot surfaces (M9/M10 own them).
  - AI-content + external shareability (v1.3).

## Depends on
**M9a + M9b + M10** (curates their output). **Parallel with:** none (the closing milestone before
`/developer-kit:close-release`).

## Open questions (resolve during build)
- Whether snapshot **capture** is a curator step or a manifest-cached refresh (decide alongside M9a's
  capture-source policy — cache-hit by default, refresh on schema drift).
- **`/demo-seed` extension vs a new `/demo-snapshot` skill** (driving the `stacksnap` CLI) — pick the cleaner UX.

## KB dependencies (read as contract)
- `corpus/ops/demo/README.md` + the recipe files (the family to refresh)
- `corpus/ops/seeding-spec.md` + `corpus/ops/snapshot-spec.md` (the now-complete surface catalog)
- the `/demo-seed` skill (the extension/companion baseline)

## Delivers → corpus/ops/demo/ (refresh) + the /demo-snapshot skill + CLAUDE.md
- refreshes **`corpus/ops/demo/`** (recipes + presets to full-fidelity worlds).
- the **`/demo-snapshot` skill** (capture/replay) + its CLAUDE.md skill-table entry.
