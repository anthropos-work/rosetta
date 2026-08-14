---
milestone: M261
title: "Load the new canon"
milestone_shape: section
status: planned
release: "02.90-new-alphabet"
depends_on: "M260"
parallel_with: "none"
complexity: medium
last_updated: "2026-08-14"
---

# M261: Load the new canon

**Goal:** A stack replays the new canon, cold, from a fresh capture.

## Scope

**In:**
  - re-capture from a safe prod source under the existing capture-source policy
  - PURGE rather than refresh the snapshot cache — node-ids moved, so a stale artifact is not merely old, it is WRONG
  - replay proven on a cold stack; the cold-start runbook updated
  - confirm the batch-prompt cache invalidates on the new capture version (it is keyed on it — expected to work, must be OBSERVED)
  - **[from D-M259-3, measured for you by M260]** add the net-new taxonomy tables to the capture
    surface and prove they replay. M260 measured their scoping so this milestone does not re-derive it:
    **`skill_redirect`** (`id`, `created_at`, `old_node_id`, `skill_id`, `source`, `score`, `review`) and
    **`job_role_redirect`** (same shape, `job_role_id`) carry **ZERO org-scoping** — they are canon-level
    `PureReference` tables. **`category_translation`** and **`specialization_translation`** exist and are
    parent-scoped like the two translation tables already captured. **`taxonomy_canon_state` was NOT found
    in `app/internal/data/ent/schema/`** at `4bccda085` despite appearing in the taxonomy-v2 commit range —
    confirm what it is before declaring it, rather than declaring it blind.
    M260 deliberately did NOT declare these: a column list that no capture has ever validated is a
    confident guess, and validating it is this milestone's whole job.

**Out:**
  - seeded data (M262)

## Depends on

M260

## Parallel with

none

## Open questions

  - Does the capture-source policy still hold for the new canon's tables?

## KB dependencies

- [`snapshot-cold-start.md`](../../../../corpus/ops/snapshot-cold-start.md)
- [`cache-spec.md`](../../../../corpus/ops/demo/cache-spec.md)
- [`safety.md`](../../../../corpus/ops/safety.md)

## Constraints (release-wide, non-negotiable)

- **Zero platform-repo edits.** "Update the platform repos" in this release means **pull them fresh**. A need
  that can only be met by a platform edit **escalates**; it does not edit.
- All stack tooling lives in `rosetta-extensions`, built + tested in the authoring copy, tagged, **pushed**,
  then consumed per-stack at a pinned tag.
- Secrets handled values-blind.
