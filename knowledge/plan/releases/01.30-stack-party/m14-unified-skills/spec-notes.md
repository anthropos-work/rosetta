# M14 — Spec notes

Technical notes accumulate here during build.

## Pre-flight audits — §1 (dev-up/dev-down)
KB-fidelity (Phase 0b, 2026-06-07): **GREEN**. Report: `kb-fidelity-audit.md`. The docs being renamed ARE
the deliverable (SKILL.md files); every driven CLI exists + is stable in `.agentspace/rosetta-extensions/`
(`dev-stack` up/down/gen/status, `rosetta-demo` status, `stackseed`, `stacksnap`, `stack_registry.py`).
Reference blast radius enumerated: ~30 rosetta files + 4 extension-clone files. No blind areas, no stale
load-bearing claims. Sha at audit: pre-§1 HEAD (3f53205-rooted branch).

### Topic → doc → code triples (fast future-audit start)
- `dev-up` ← `setup-platform` SKILL + `start-platform` SKILL → `setup_guide.md`/`run_guide.md` + `dev-stack up` + `dev-setdress.sh`
- `dev-down` (new) → `dev-stack down`
- `stack-list` ← `demo-status` SKILL → `rosetta-demo status` + `stack_registry.py`
- `stack-seed` ← `demo-seed` SKILL → `stack-seeding/cmd/stackseed` (already `--stack dev-N|demo-N`)
- `stack-snapshot` ← `demo-snapshot` SKILL → `stack-snapshot/cmd/stacksnap` (already replays into dev-N or demo-N)
- `stack-update` ← `update-platform` SKILL → `make pull/migrate/up` (per-stack platform dir)
- `demo-up`/`demo-down` retained → `demo-stack/up-injected.sh` + `rosetta-demo`

### M14-Q1 resolution anchor
Boundary already documented in `dev-stack/dev-stack` header + `dev-stack/README.md`: first-time heavy machine
setup (tool install, org clone, ~15-25min build) = the `setup_guide.md` one-time path; `dev-up` drives that
for stack N=0/first-build AND spins additional isolated `dev-N` via `dev-stack up`. Consolidation is of the
*operation skills*, not a re-implementation of the build.

## Skill map (after v1.3)
_Lifecycle: dev-up, dev-down, demo-up, demo-down. Generic ops: stack-list, stack-seed, stack-snapshot, stack-update.
Removed: setup-platform, start-platform, update-platform, demo-status, demo-seed, demo-snapshot (hard rename, no aliases)._

## dev-up consolidation
_setup-platform + start-platform → dev-up; drives the M13 dev bring-up (local Directus + snapshot + light seed)._

## Reference sweep
_CLAUDE.md skill table, READMEs, corpus/ops/ guides, demo/ recipes._
