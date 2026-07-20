# M229 "academy content-model re-ground" — Retro

## Summary
Corrected the materially-stale `ant-academy.md` (and 3 sibling docs) from the false "no backend / static JSON /
only Clerk" model to the true DB-authoritative catalog (grid → academy subgraph over GraphQL → `emptyCatalogView()`
on failure). The KB-fidelity prerequisite that unblocks the whole v2.5 academy thread — and the doc whose wrongness
mis-routed the F4 empty-grid carry into the platform repo for a full release.

## Incidents this cycle
- None. (The milestone's value is *preventing* the recurring incident: a wrong doc sending a real fix to the wrong repo.)

## What went well
- Every code claim was spot-verified against the actual `stack-demo/ant-academy/code/` source before it was written —
  no doc-on-doc drift. The 6-agent design-roadmap research had already code-cited the model, so the rewrite was fast + accurate.
- Phase 3's corpus-wide scan caught the *same* false claim in `run_guide.md` + `CLAUDE.md` — fixed as Fate-1 rather than left to rot.

## What didn't
- The false claim had propagated to 4 docs; a single-file mental model of "the doc" would have missed 3 of them. The
  corpus-wide grep at close is what made it complete.

## Carried forward
- None. M230 (production-faithful fill) + M231 (content-stories spike) are already-planned siblings, not carries.

## Metrics delta
- 4 docs corrected · 0 platform edits · 0 tests (docs milestone) · 0 flakes. See `metrics.json`.
