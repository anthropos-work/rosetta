---
milestone: M4
slug: consolidate-extensions
version: v1.1 "show floor"
milestone_shape: section
status: done
completed: 2026-06-04
created: 2026-06-04
last_updated: 2026-06-04
delivers: rosetta-extensions/ (monorepo) + corpus pointers
---

# M4 — Consolidate into the `rosetta-extensions` monorepo

## Goal
Collapse the repo constellation into **two repos** — `rosetta` (the platform corpus + dev-env skills) and a new
**`rosetta-extensions`** monorepo (the operations-extension sections) — before it explodes into complexity.
Move `clerkenstein` + `rosetta-demo` in **with full git history**, stand up a lightweight `knowledge/` nav,
thin rosetta's corpus to pointers, and **remove the two old standalone repos completely** (org + local). Zero
platform-repo change.

## Why a section milestone
The shape is fully grounded — it's git surgery + a doc reorg, every step concrete and checklistable. The only
real risk is sequencing the irreversible deletion safely (verify-then-delete).

## Scope
### In
- **The `rosetta-extensions` monorepo** — a new **private** org repo `anthropos-work/rosetta-extensions`,
  cloned locally under the gitignored `anthropos-demo/rosetta-extensions/` (replacing the two separate dirs).
- **`git subtree` import, history-preserving** (M4-D1): `clerkenstein` → `clerkenstein/`; `rosetta-demo` →
  `demo-stack/`. Every v1.0 + M3 commit stays reachable. (Cross-repo relative refs survive — demo-stack +
  clerkenstein remain siblings, so `../../clerkenstein` still resolves.)
- **`rosetta-extensions/knowledge/`** — the lightweight nav: a `README.md`/index that lists the sections
  (`clerkenstein`, `demo-stack`, + the future `stack-injection` / `dev-stack` / `stack-seeding`), clarifies each
  one's role/scope, and **points into each section's own knowledge** (clerkenstein's KB, demo-stack's docs) —
  no duplication.
- **Thin rosetta's corpus to references:** `corpus/ops/rosetta_demo.md` → a short pointer ("the demo-stack
  orchestration lives in `rosetta-extensions/demo-stack/`; this corpus documents the platform, not the stacks");
  `corpus/services/clerkenstein.md` → a pointer to `rosetta-extensions/clerkenstein/`. Repoint the `/demo-*`
  skills + any `anthropos-demo/{clerkenstein,rosetta-demo}` path refs to the new monorepo paths.
- **Verify before cutover:** every suite passes under the new paths — clerkenstein **218** Go test/fuzz funcs
  (`-race`) + the 4 alignment gates (incl. deploy 7/7), demo-stack **78** tests (incl. the cross-repo mint
  drift-guard) + shellcheck. The monorepo is pushed + the history is confirmed present.
- **Remove the old repos completely (M4-D2, user-directed):** **only after** the monorepo is pushed + verified,
  delete `anthropos-work/clerkenstein` + `anthropos-work/rosetta-demo` (org, via `gh repo delete`) **and** the
  local `anthropos-demo/clerkenstein` + `anthropos-demo/rosetta-demo` dirs. History is preserved inside the
  monorepo, so the originals are redundant. Final `gh repo delete` confirmed at the moment of deletion.
- **Delivers** the monorepo + `rosetta-extensions/knowledge/README.md` + the thinned rosetta pointers.

### Out
- Extracting the reusable `stack-injection` layer — **M5**.
- `dev-stack` tooling — **M6**. `stack-seeding` (declarative seeding) — **M7**. Recipes/polish — **M8**.
- Moving the alignment framework — it **stays in rosetta** (M4-D3); only its consumer (clerkenstein) moves.
- Any change to a read-only platform repo (`anthropos-dev/*`).

## Decisions locked at design (2026-06-04)
- **M4-D1 — `git subtree` (history-preserving), not `filter-repo`** (user call): import each repo into its subdir
  with full history via `git subtree add --prefix=…`. Old commits show root paths pre-import (the merge relocates
  them); good enough. (filter-repo for pristine per-commit subdir paths only if later wanted.)
- **M4-D2 — DELETE the old repos, not archive** (user-directed, supersedes the design's "archive read-only"):
  remove `anthropos-work/{clerkenstein,rosetta-demo}` from the org + locally, after the monorepo verifies.
- **M4-D3 — the alignment framework stays in rosetta** (user call): `test/alignment/` + `/align-*` +
  `alignment_testing.md` are a generic capability rosetta ships; extensions only holds clerkenstein (the consumer).

## Open questions (resolve later)
- The shared **port-offset / multi-instance engine** (`gen_override.py`): its own section vs folded into
  `stack-injection` vs a shared lib — settle in **M5**.
- Whether **M4 + M5 merge** if the consolidation lands small — decide after M4's `git subtree` step.

## Depends on
M3 (rosetta-demo exists) + v1.0 (clerkenstein exists), both pushed to the org. **Blocks** M5–M8 (everything
rehomes through the monorepo).

## Exit (section)
All `In:` deliverables land: the monorepo exists with both subtrees (history intact), the knowledge nav is up,
rosetta is thinned to pointers, every suite passes under the new paths, the monorepo is pushed, and the two old
repos are gone (org + local). Documented in `rosetta-extensions/knowledge/README.md`.
