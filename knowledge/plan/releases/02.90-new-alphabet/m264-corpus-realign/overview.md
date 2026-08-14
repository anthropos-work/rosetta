---
milestone: M264
title: "The corpus tells the truth"
milestone_shape: section
status: planned
release: "02.90-new-alphabet"
depends_on: "M259 (numbers), M261 (observed replay counts)"
parallel_with: "M262, M263"
complexity: medium
last_updated: "2026-08-14"
---

# M264: The corpus tells the truth

**Goal:** No corpus page states a taxonomy figure or module fact the new canon has falsified.

## Scope

**In:**
  - the ~17 count-claims across shared_libraries.md (canonical figures AND the AKB contradiction table), ai_architecture.md (x5, incl. embeddings row counts), architecture_overview.md, service_taxonomy.md, toolchain_overview.md, org-repos.md, CLAUDE.md
  - the `taxonomy`-folded-into-`app` correction — `app/go.mod` five org modules -> ZERO first-party (`app` `e72f18199`)
  - re-grade the four taxonomy repos' "dormant, DECIDE" verdicts

**Out:**
  - authoring the source-pipeline doc (M259's `Delivers ->`)

## Depends on

M259 (numbers), M261 (observed replay counts)

## Parallel with

M262, M263

## Open questions

  - Does the '60K skills UNVERIFIED / 18K roles REFUTED' adjudication survive the canon, or is it now moot?

## KB dependencies

- [`shared_libraries.md`](../../../../corpus/architecture/shared_libraries.md)
- [`ai_architecture.md`](../../../../corpus/architecture/ai_architecture.md)
- [`org-repos.md`](../../../../corpus/architecture/org-repos.md)

## Constraints (release-wide, non-negotiable)

- **Zero platform-repo edits.** "Update the platform repos" in this release means **pull them fresh**. A need
  that can only be met by a platform edit **escalates**; it does not edit.
- All stack tooling lives in `rosetta-extensions`, built + tested in the authoring copy, tagged, **pushed**,
  then consumed per-stack at a pinned tag.
- Secrets handled values-blind.
