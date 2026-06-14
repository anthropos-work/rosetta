---
milestone: M24
slug: docs-hygiene
version: v1.5 "prop room"
milestone_shape: section
status: archived
created: 2026-06-11
last_updated: 2026-06-13
complexity: small
delivers: corpus-wide truth-up + 3 stale-doc corrections (rosetta) + README-index-row guard + zero-critical-genes guard + stats-scope fix + Go pin bump (rosetta-extensions)
backlog_refs: NEW-5 (Go toolchain bump), NEW-6 (README index-row guard), NEW-11 (zero-critical-genes guard), NEW-14 (stats-script scope fix)
---

# M24 — Docs convergence + hygiene strand

## Goal
Make the whole corpus tell the new truth (stacks are content-self-contained), and absorb the four small aged-out
hygiene items the deferral audit surfaced — so v1.5 leaves the repo honest and the backlog cleaner.

## Why section
Every deliverable is a known, bounded edit. The doc sweep follows the finished M21–M23 cutover; the four hygiene
items are each small + independently landable. Build with `/developer-kit:build-milestone`.

## Repo split
- **`rosetta`**: the corpus-wide truth-up + the three stale-doc corrections.
- **`rosetta-extensions`** (authoring → tag `prop-room-m24` → consume): the README-index-row guard, the
  zero-critical-genes guard, the stats-scope fix, the Go pin bump.

## Scope
- **In (`rosetta` docs):**
  - Rewrite the `snapshot-spec.md` known-state block (the per-stack Directus is now real; exit-4 semantics
    redefined), the `safety.md` §2 deltas, finish `corpus/ops/directus-local.md`.
  - **Correct the stale local-Directus claims** in `corpus/architecture/external_services.md` (image 10.10.1 +
    admin/password + a compose snippet — all **false**; the platform compose has no directus service — verified),
    `service_taxonomy.md:242-260`, `quick_ops.md:162`.
  - Sweep the "print-only / exit-4 / reads live from prod" language across the skills + `CLAUDE.md` (via
    `/update-knowledge`).
- **In (`rosetta-extensions` — the hygiene strand, each small + independently landable):**
  - **(a)** bump the Go toolchain pin to **go1.25.11+** (clears the 12 called-stdlib advisories; **lazy rebuild** —
    no dedicated rebuild session, per the user).
  - **(b)** a **corpus README index-row guard** (a lint that fails when a new doc lacks its directory-README index
    row — the recurring miss in 3 straight releases; v1.5 ships new docs, its exact protected class).
  - **(c)** the **alignment zero-critical-genes guard** (`dna.Validate`/`compare.pct` treat a zero-critical DNA as
    100% — verified still absent at `compare.go:247-252` / `dna.go:168-183`; reject/flag it).
  - **(d)** the **`/project-stats` scope fix** (stop scanning the gitignored `stack-*/` platform clones that inflate
    the absolutes).
- **Out:** anything in M21–M23; the DROPPED items (AI content, shareability, more mirrors, deploy-CI gate, dev-up
  pre-warm).

## Depends on / parallel with
Depends on: M23 (the docs describe the finished cutover). Parallel with: none (but the four hygiene items are
internally independent — land in any order).

## Open questions
None material (the four hygiene items are self-contained).

## KB dependencies
The full `corpus/ops/` + `corpus/architecture/` + `corpus/services/` set touched by the cutover; the `alignment` +
stats tooling for the hygiene items.

## Risk (low)
The stale-doc corrections must be verified against the *actual* platform compose, not assumed (already verified
once: no directus service exists in `stack-dev/platform/docker-compose.yml`).
