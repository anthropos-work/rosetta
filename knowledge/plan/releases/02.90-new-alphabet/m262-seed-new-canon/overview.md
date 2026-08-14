---
milestone: M262
title: "The seed speaks the new canon"
milestone_shape: section
status: planned
release: "02.90-new-alphabet"
depends_on: "M261"
parallel_with: "M263, M264"
complexity: large
last_updated: "2026-08-14"
---

# M262: The seed speaks the new canon

**Goal:** Every seeded hero's skill chain resolves in the new canon, and a PARTIAL failure is loud.

## Scope

**In:**
  - remap seeded refs THROUGH the redirect map
  - re-resolve the 8 literal job-role names the presets pin (Account Executive, Backend Developer, Engineering Manager, Sales Manager, Data Analyst, DevOps Engineer, Business Operations Analyst, Talent Acquisition Specialist)
  - add a PER-HERO RICHNESS FLOOR to the closure gene
  - PRICE the AI-profile regeneration before spending it, then re-run `gen-batch` under an explicit `--max-cost`

**Out:**
  - the taxonomy browser (M263)
  - corpus counts (M264)

## Depends on

M261

## Parallel with

M263, M264

## Open questions

  - How much of the 7-table verified-skill fan-out survives a remap vs needs re-seeding?
  - What does the regeneration actually price at?

## KB dependencies

- [`stories-spec.md`](../../../../corpus/ops/demo/stories-spec.md)
- [`seeding-spec.md`](../../../../corpus/ops/seeding-spec.md)
- [`ai-generation-spec.md`](../../../../corpus/ops/demo/ai-generation-spec.md)

## Constraints (release-wide, non-negotiable)

- **Zero platform-repo edits.** "Update the platform repos" in this release means **pull them fresh**. A need
  that can only be met by a platform edit **escalates**; it does not edit.
- All stack tooling lives in `rosetta-extensions`, built + tested in the authoring copy, tagged, **pushed**,
  then consumed per-stack at a pinned tag.
- Secrets handled values-blind.
