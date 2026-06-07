---
milestone: M13
slug: dev-peers
version: v1.3 "stack party"
milestone_shape: section
status: archived
created: 2026-06-07
last_updated: 2026-06-07
complexity: large
delivers: updates corpus/ops/seeding-spec.md (dev-min preset + dev auto-seed) + corpus/ops/snapshot-spec.md (dev as a replay target + local-Directus-on-dev)
---

# M13 — Dev stacks as full-fidelity peers (local Directus + auto-snapshot + light seed)

## Goal
A freshly-built dev stack gets the demo treatment — its **own local per-stack Directus** (no longer pointing at
shared prod), an **auto-snapshot** on build that fills the real public reference data (taxonomy + the now-local
content), and a **light default seed** so it is never empty.

## Why section
Decomposable up front (reuses M10's per-stack Directus + the M9/M7 snapshot/seed framework). The fidelity gate is a
per-surface acceptance check. No emergent path.

## Scope
- **In:**
  - Wire the **dev bring-up to spawn a per-stack Directus** (reuse M10's `stack-snapshot/directus/provision.go`) + repoint the dev CMS at it (no longer shared prod Directus).
  - **Auto-run `stacksnap replay`** (taxonomy + directus) as part of dev build — cache-first (replay, not capture) so it's fast; `--no-snapshot` to skip.
  - A **`dev-min` seed preset** (smaller than demo's `small-200` — ~1 org + ~10 users + minimal activity) applied on build.
  - Keep the **n=0-dev-reset guard** (never wipe the main dev stack without `--force`).
- **Out:** the generic skills (M14); media blob bytes (v1.4 — refs-only, as for demo).

## Depends on
**M12** (consumes the registry for dev-N) + v1.2's **M10** (per-stack Directus + content snapshot) + v1.1's **M6** (dev-stack) + **M7** (seeding). **Parallel with:** none (gates M14).

## Open questions (resolve during build)
- Does spawning Directus + snapshot + seed make dev bring-up too heavy? (mitigate: cache-first snapshot, minimal seed, reuse M10's provision).
- Auto-snapshot default-on vs opt-in for dev (lean: default-on, `--no-snapshot` to skip).
- `dev-min` preset exact size (1 org / ~10 users / how much activity).

## KB dependencies (read as contract)
- `corpus/ops/snapshot-spec.md` (capture/replay + the per-stack Directus store fork)
- `corpus/ops/seeding-spec.md` (presets + the dev/n=0 guard)
- `corpus/services/cms.md` (Directus) + the `dev-stack` extension section

## Delivers → corpus/ops/seeding-spec.md + corpus/ops/snapshot-spec.md
- Updates `seeding-spec.md` (the `dev-min` preset + dev auto-seed on build).
- Updates `snapshot-spec.md` (dev as a replay target + the local-Directus-on-dev path).
