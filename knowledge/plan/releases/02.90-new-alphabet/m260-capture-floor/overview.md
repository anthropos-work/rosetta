---
milestone: M260
title: "The floor comes down"
milestone_shape: section
status: complete
release: "02.90-new-alphabet"
depends_on: "M259"
parallel_with: "none"
complexity: medium
last_updated: "2026-08-14"
---

# M260: The floor comes down

**Goal:** No Rosetta tool asserts a taxonomy SIZE it did not measure this run.

## Scope

**In:**
  - `MinRows` DERIVED from the capture rather than pinned at 40000 — preserving the under-capture protection the floor exists for (an empty or mis-filtered capture must still abort)
  - re-ground every taxonomy-size assumption across `stack-snapshot`, `stack-seeding`, `stack-verify`
  - a fence that fails when a bare taxonomy count is re-pinned into source

**Out:**
  - re-capturing (M261)
  - the seed's own refs (M262)

## Depends on

M259

## Parallel with

none

## Open questions

  - Is there a legitimate floor at all once the canon is governed, or does the guard become a shape check?

## KB dependencies

- [`snapshot-spec.md`](../../../../corpus/ops/snapshot-spec.md)
- [`verification.md`](../../../../corpus/ops/verification.md)

## Constraints (release-wide, non-negotiable)

- **Zero platform-repo edits.** "Update the platform repos" in this release means **pull them fresh**. A need
  that can only be met by a platform edit **escalates**; it does not edit.
- All stack tooling lives in `rosetta-extensions`, built + tested in the authoring copy, tagged, **pushed**,
  then consumed per-stack at a pinned tag.
- Secrets handled values-blind.
