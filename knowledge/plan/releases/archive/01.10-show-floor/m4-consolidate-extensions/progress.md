# M4 — progress (section checklist)

**Milestone:** M4 — Consolidate into the `rosetta-extensions` monorepo · **Shape:** section · **Status:** done (2026-06-04)

## Done
- [x] **`rosetta-extensions` monorepo created** — private org repo `anthropos-work/rosetta-extensions`, cloned
  locally at `anthropos-demo/rosetta-extensions/`. README + `knowledge/README.md` nav skeleton.
- [x] **`git subtree` import, history-preserving** — `clerkenstein` → `clerkenstein/` + `rosetta-demo` → `demo-stack/`.
  **Full history proven present** (73 commits = 50 ck + 16 rd + init + merges + the M4 fixups; every spot-checked SHA
  reachable, incl. the oldest M1 commits).
- [x] **`knowledge/` nav** — `rosetta-extensions/knowledge/README.md` lists the sections + points into each one's KB.
- [x] **Thinned rosetta to pointers** — `corpus/ops/rosetta_demo.md` → pointer (the guide moved to
  `demo-stack/GUIDE.md`); `corpus/services/clerkenstein.md` → repointed; `/demo-*` skills' CLI path repointed.
- [x] **Path-depth fixes (caught by the verify gate)** — the tooling moved +1 level deeper, so `$HERE/../..` (REPO_ROOT,
  DEV) → `$HERE/../../..`, and `up-injected.sh` CLERK → the monorepo sibling `$HERE/../clerkenstein`. Verified:
  REPO_ROOT→rosetta, DEV+CLERK resolve, shellcheck clean.
- [x] **Verified under the new paths** — demo-stack **78 tests** green, clerkenstein suites green, **deploy gate
  100%/100% (7/7)**, monorepo self-contained (0 old-path refs), pushed.
- [x] **Local old dirs removed** — `anthropos-demo/{clerkenstein,rosetta-demo}` deleted; `anthropos-demo/` now holds
  only `rosetta-extensions/`.

- [x] **Old ORG repos deleted** — `anthropos-work/clerkenstein` + `anthropos-work/rosetta-demo` both gone (404).
  The `gh` token lacked the `delete_repo` scope, so the user deleted them via the GitHub web UI; their full history
  lives inside `rosetta-extensions`. `anthropos-work/rosetta-extensions` (PRIVATE, 73 commits) is the sole survivor.

## Completeness Ledger (section)
- **Done (Fate 1):** all `In:` items — the monorepo + history-preserving subtree import, the `knowledge/` nav, the
  rosetta thinning (pointers + repoints), the path-depth fixes (caught by verify), full verification under the new
  paths (78 demo-stack + clerkenstein suites + deploy gate 7/7), the monorepo push, and the old-repo removal (local
  + org).
- **Routed (Fate 3) → M5:** the shared port-offset engine's home; the M4+M5-merge question (now moot — M4 closed
  standalone, M5 proceeds as planned).
- **Dropped / escape-hatch:** none.

## Decisions
M4-D1 git subtree (history-preserving) · M4-D2 delete (not archive) the old repos · M4-D3 alignment framework stays
in rosetta · M4-D4 path-depth fix for the +1 monorepo level (verify-caught). See [decisions.md](decisions.md).

## Verification
monorepo 73 commits (full history of both repos, proven); demo-stack **78 tests** + clerkenstein suites green;
**deploy gate 100%/100% (7/7)** held throughout; shellcheck clean; monorepo self-contained (0 old-path refs); the
two old org repos deleted (404). dev stack untouched.
