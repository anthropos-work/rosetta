---
title: "KB Fidelity Audit — M14 (Unified stack-* skills + dev-up/dev-down)"
date: 2026-06-07
scope: milestone:M14
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| `setup-platform` skill (→ folds into `dev-up`) | `.claude/skills/setup-platform/SKILL.md` + `corpus/ops/setup_guide.md` | n/a (skill is the doc; drives `make`/setup_guide) | PAIRED |
| `start-platform` skill (→ folds into `dev-up`) | `.claude/skills/start-platform/SKILL.md` + `corpus/ops/run_guide.md` | n/a (drives `make up`) | PAIRED |
| `update-platform` skill (→ `stack-update`) | `.claude/skills/update-platform/SKILL.md` + `corpus/ops/update_guide.md` | n/a (drives `make pull`/`migrate`/`up`) | PAIRED |
| `demo-status` skill (→ `stack-list`) | `.claude/skills/demo-status/SKILL.md` | `.agentspace/rosetta-extensions/demo-stack/rosetta-demo` (`status`) + `stack-core/stack_registry.py` | PAIRED |
| `demo-seed` skill (→ `stack-seed`) | `.claude/skills/demo-seed/SKILL.md` | `.agentspace/rosetta-extensions/stack-seeding/cmd/stackseed` | PAIRED |
| `demo-snapshot` skill (→ `stack-snapshot`) | `.claude/skills/demo-snapshot/SKILL.md` | `.agentspace/rosetta-extensions/stack-snapshot/cmd/stacksnap` | PAIRED |
| `demo-up`/`demo-down` (retained) | `.claude/skills/demo-up,down/SKILL.md` + `corpus/ops/rosetta_demo.md` | `demo-stack/up-injected.sh` + `rosetta-demo` | PAIRED |
| `dev-up`/`dev-down` (net-new) | — (delivered by M14) | `dev-stack/dev-stack` (`up`/`down`) + `dev-stack/dev-setdress.sh` (M13) + `corpus/ops/setup_guide.md`/`run_guide.md` | DOC-ONLY (target exists; skill is the deliverable) |
| Unified registry (drives target detection) | `corpus/ops/rosetta_demo.md` (M12 update) | `stack-core/stack_registry.py` | PAIRED |

## Fidelity Findings
1. **CLAUDE.md skill table vs reality — ALIGNED.** All 14 table rows match the live `.claude/skills/` dirs; the 6 to-be-renamed entries describe the current pre-rename state accurately.
2. **`demo-snapshot` SKILL already accepts `dev-N` — ALIGNED.** The skill body already documents replay into `demo-N` *or* `dev-N`, so the `stack-snapshot` rename is a name change, not a capability gap.
3. **`dev-up` setup-boundary (M14-Q1) — ALIGNED / well-anchored.** `dev-stack/dev-stack` header + `dev-stack/README.md` already state the boundary: the heavy first-time machine setup (tool install, org clone, ~15-25min build) is driven by `setup_guide.md`; `dev-stack up N` spins up *additional isolated* dev-N copies. So `dev-up` consolidates the two skills' *operation*, the first-time prereq stays a documented one-time path. No stale claim — the contract is already written.
4. **Extension-clone references to old skill names — load-bearing for the sweep, not stale-now.** `dev-stack/dev-stack:5`, `dev-stack/README.md:6`, `demo-stack/GUIDE.md:5,74`, `stack-verify/repos/run.sh:29` reference `/setup-platform` / `/start-platform` / `/update-platform` / `/demo-*`. These are accurate *today*; the rename makes them dead refs → in M14's "update every reference" section (extensions-clone side, committed under the same discipline).

## Completeness Gaps
None critical. `dev-up`/`dev-down` are DOC-ONLY by design — they are the milestone's net-new deliverable, and their drive target (`dev-stack` CLI + the M13 set-dress pass + setup/run guides) is fully documented. No blind area.

## Applied Fixes
None needed inline — all docs accurately describe the current pre-rename state. The renames + reference sweep ARE the milestone work (Phase 1), not pre-flight backfill.

## Open Items (require user decision)
None. M14-Q1/Q2/Q3 are build-time design choices (resolve in implementation), not KB blind areas.

## Gate Result
GREEN: proceed to Phase 1. The docs being renamed are themselves the deliverable; every CLI the renamed skills drive exists and is stable in `.agentspace/rosetta-extensions/`. The reference blast radius (~30 rosetta files + 4 extension-clone files) is enumerated and is in-scope for the §4 reference-sweep section.
