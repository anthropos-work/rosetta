---
milestone: M12
slug: stack-registry
version: v1.3 "stack party"
milestone_shape: section
status: planned
created: 2026-06-07
complexity: medium
delivers: updates corpus/ops/rosetta_demo.md (the unified registry + first-available-N model)
---

# M12 — Unified stack registry + first-available-N allocation

## Goal
One shared registry across **dev + demo** that tracks live stacks and allocates the **lowest free N**, so bring-ups
never collide on ports: building `dev, demo, dev, demo, demo` yields `dev-1, demo-2, dev-3, demo-4, demo-5`.

## Why section
The deliverables are enumerable: the registry schema, the allocator, the up/teardown wiring. No emergent path.

## Scope
- **In:**
  - A **unified stack registry** spanning both kinds (extend the demo `registry.json` → records `{type: dev|demo, N, ports, status, created}`) in `stack-core` (or a small new `stack-registry` shared module).
  - A **first-available-N allocator** — reconcile the registry against live `docker ps`, return the lowest free N. Race-safe (locked registry write).
  - The up-paths **accept an explicit N OR auto-allocate** the next free one; teardown frees the slot.
  - Consumed by both `dev-stack` and `demo-stack` bring-up (the wiring; the generic skills come in M14).
- **Out:** the skill renames (M14); dev local-Directus/snapshot/seed (M13).

## Depends on
v1.1's `stack-core` (port-offset engine) + the demo-stack `registry.json`. **Parallel with:** M13 (feasible; lean sequential so M13's dev bring-up consumes the registry).

## Open questions (resolve during build)
- Registry-of-record (a lockfile) vs `docker ps`-derived — lean: registry is the record, `docker ps` reconciles used-N.
- How a manually-started stack (outside the skills) is reconciled into the registry.
- Lock mechanism for concurrent `up`s (flock vs an atomic write).

## KB dependencies (read as contract)
- `corpus/ops/rosetta_demo.md` (the demo registry + port-offset model)
- the `stack-core` / `demo-stack` / `dev-stack` extension sections (the existing port engine + the demo registry.json)

## Delivers → corpus/ops/rosetta_demo.md
- Updates `corpus/ops/rosetta_demo.md` with the **unified cross-type registry** + the **first-available-N** allocation model.
