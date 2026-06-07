# M14 — Progress

**Shape:** section · **Status:** build complete (all 5 sections landed)

## Section checklist (from overview Scope.In)
- [x] `dev-up` (consolidate setup-platform + start-platform → dev bring-up, drives M13 flow) + `dev-down` — §1, commit `e6c2da6`
- [x] Hard-rename ops skills → `stack-list`/`stack-seed`/`stack-snapshot`/`stack-update` (dev-N|demo-N target) — §2, commit `d6afb70`
- [x] Remove old skill dirs (setup-platform, start-platform, update-platform, demo-status/seed/snapshot) — §3, commit `d6afb70`
- [x] Update every reference: CLAUDE.md skill table, READMEs, corpus/ops/ guides, demo/ recipes — §4, commits `3b8e80e` (rosetta) + `b37e831` (extensions)
- [x] `demo-up`/`demo-down` retained + aligned with `dev-up`/`dev-down` — §5, commit `23b2398`

## Build notes
- KB-fidelity Phase 0b: **GREEN** (`kb-fidelity-audit.md`) — docs ARE the deliverable; every driven CLI exists.
- PR review (Phase 3): **1 finding** (D1 — `stack-seed --preset` is a skill-level shorthand, not a CLI flag),
  fixed inline (commit `28cb5c4`, → M14-D5). A/B/C clean; all links resolve; renamed skills' argument-hints
  verified against the `dev-stack`/`stackseed`/`stacksnap` CLIs.
- Decisions: M14-D1…D6 (D2/D3/D4 resolve Q1/Q2/Q3; D5 = the review fix; D6 = CHANGELOG convention).
- Extensions: 1 doc-only commit on `main` (`b37e831`) for the 4 extension-clone references. No tag needed
  (no CLI/section change); the renamed skills drive the existing M12/M13 binaries at their pinned tags.
- Three-fate guard: **0 deferrals** — every item (incl. the `--preset` finding) landed Fate-1 in M14.

## Final review
_(filled at close)_
