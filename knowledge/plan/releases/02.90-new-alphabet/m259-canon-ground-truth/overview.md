---
milestone: M259
title: "Canon ground truth"
milestone_shape: section
status: planned
release: "02.90-new-alphabet"
depends_on: "none"
parallel_with: "none"
complexity: medium
last_updated: "2026-08-14"
---

# M259: Canon ground truth

**HARD go/no-go barrier** — downstream milestones are sized against its findings; a GO/NO-GO verdict is a deliverable.

**Goal:** Establish what the new taxonomy actually IS — from the repos and a prod read — before any downstream milestone is sized against a guess.

## Scope

**In:**
  - pull EVERY platform repo fresh into BOTH clone sets (`stack-dev/`, `stack-demo/`) — pull, never commit
  - read the canon, the **redirect map**, and the retired-id guard in `app` (`0b5cef2d2`, `34b5b9635`)
  - measure the real new counts: skills, roles, specializations, categories, and the embeddings-pruning effect
  - establish whether the redirect map is TOTAL or PARTIAL, and what a retired id with no successor should become
  - reconcile the three lineages: `skills-and-job-roles` 12,201 · prod 42,790 · canon ~4k

**Out:**
  - changing any rext code (M260+)
  - editing any platform repo, ever

## Depends on

none

## Parallel with

none

## Open questions

  - Is the redirect map total or partial?
  - Does a retired skill with no successor exist, and what should a seeded ref pointing at one become?
  - Did taxonomy v2 add columns/tables the capture firewall has never seen?

## KB dependencies

- [`shared_libraries.md#taxonomy-figures`](../../../../corpus/architecture/shared_libraries.md#taxonomy-figures)
- [`snapshot-spec.md`](../../../../corpus/ops/snapshot-spec.md)
- [`org-repos.md`](../../../../corpus/architecture/org-repos.md)

## Delivers

corpus/architecture/taxonomy-canon.md — the corpus has NO doc anchor for the taxonomy source pipeline at all (Phase-0b blind area). `org-repos.md` lists all four candidate repos as "dormant >=18 months, DECIDE", which the revamp contradicts.
## Constraints (release-wide, non-negotiable)

- **Zero platform-repo edits.** "Update the platform repos" in this release means **pull them fresh**. A need
  that can only be met by a platform edit **escalates**; it does not edit.
- All stack tooling lives in `rosetta-extensions`, built + tested in the authoring copy, tagged, **pushed**,
  then consumed per-stack at a pinned tag.
- Secrets handled values-blind.
