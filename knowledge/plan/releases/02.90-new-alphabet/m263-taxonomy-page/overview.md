---
milestone: M263
title: "The taxonomy page is reachable"
milestone_shape: section
status: planned
release: "02.90-new-alphabet"
depends_on: "M261"
parallel_with: "M262, M264"
complexity: medium
last_updated: "2026-08-14"
---

# M263: The taxonomy page is reachable

**Goal:** A hero can navigate and review the taxonomy on a demo, and it is covered so it cannot rot unseen.

## Scope

**In:**
  - prove Library -> Taxonomy renders and walks index -> category -> specialization -> role -> skill on the replayed canon
  - prove a retired id renders `MovedNotice` rather than a bare 404
  - a PLAYTHROUGH asserting the navigation — a state change on click, never presence
  - confirm the hiring-org exclusion is intentional

**Out:**
  - the `internal/tools` Taxonomy tab (admin surface, out of demo scope this release)

## Depends on

M261

## Parallel with

M262, M264

## Open questions

  - Is `showLibrary` true for every seeded demo org, or does some org lose the nav entry?

## KB dependencies

- [`playthroughs.md`](../../../../corpus/ops/demo/playthroughs.md)
- [`frontend-tier.md`](../../../../corpus/ops/demo/frontend-tier.md)

## Constraints (release-wide, non-negotiable)

- **Zero platform-repo edits.** "Update the platform repos" in this release means **pull them fresh**. A need
  that can only be met by a platform edit **escalates**; it does not edit.
- All stack tooling lives in `rosetta-extensions`, built + tested in the authoring copy, tagged, **pushed**,
  then consumed per-stack at a pinned tag.
- Secrets handled values-blind.
