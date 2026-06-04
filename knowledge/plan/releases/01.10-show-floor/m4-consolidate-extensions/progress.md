# M4 — progress (section checklist)

**Milestone:** M4 — Consolidate into the `rosetta-extensions` monorepo · **Shape:** section · **Status:** in-progress (2026-06-04)

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

## Blocked (one open scope item)
- [ ] **Delete the old ORG repos** `anthropos-work/clerkenstein` + `anthropos-work/rosetta-demo` (M4-D2). **Blocked on
  the `delete_repo` GitHub scope** — the current `gh` token has `repo`/`read:org`/`workflow`/`gist` but not
  `delete_repo`, and the scope can't be granted non-interactively. Their full history is safely inside the monorepo,
  so this is the only step left. **Unblock:** `gh auth refresh -h github.com -s delete_repo` (then retry `gh repo
  delete`), or delete via the GitHub web UI (Settings → Delete repository).

## Decisions
M4-D1 git subtree (history-preserving) · M4-D2 delete (not archive) the old repos · M4-D3 alignment framework stays
in rosetta. (Details in [decisions.md] — to be written at close.)
