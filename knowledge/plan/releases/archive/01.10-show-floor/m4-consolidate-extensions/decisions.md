# M4 — Decisions

## M4-D1 — `git subtree`, history-preserving (user-chosen, 2026-06-04)
Import `clerkenstein` + `rosetta-demo` into the monorepo via `git subtree add --prefix=… <repo> main` (no
`--squash`) — each repo's full history grafted under its subdir via a merge commit. Chosen over `git filter-repo`
(which rewrites every commit so files always show the subdir path) for simplicity: the history is fully reachable
(73 commits = 50 ck + 16 rd + init + merges + fixups; every SHA verified present), at the cost that pre-import
commits record root paths (the merge relocates them) — `git log -- <subdir>/` shows only the merge, but the granular
history is in the log. Good enough; filter-repo remains an option if pristine per-commit subdir paths are ever wanted.

## M4-D2 — DELETE the old repos, not archive (user-directed, 2026-06-04)
The design proposed "archive read-only"; the user directed **complete removal**. After the monorepo was pushed +
verified to hold the full history, the local `anthropos-demo/{clerkenstein,rosetta-demo}` dirs were `rm`'d and the
org repos `anthropos-work/{clerkenstein,rosetta-demo}` deleted (the `gh` token lacked `delete_repo`, so the user
deleted them via the web UI). Safe: their history is redundant — preserved inside `rosetta-extensions`.

## M4-D3 — the alignment framework stays in rosetta (user-chosen, 2026-06-04)
`test/alignment/` (alignctl) + the `/align-dna`/`/align-run` skills + `corpus/architecture/alignment_testing.md`
stay in rosetta — a generic, engine-agnostic measurement capability rosetta ships. Only its consumer
(`clerkenstein`, which holds the DNAs + runners) moved into `rosetta-extensions`.

## M4-D4 — path-depth fix for the +1 monorepo level (build, 2026-06-04)
The tooling moved from `anthropos-demo/rosetta-demo/` to `anthropos-demo/rosetta-extensions/demo-stack/` — one level
deeper. The verify gate caught that the up-reaching path computations were off by one (the 78 unit tests didn't —
they don't exercise the I/O orchestrators). Fixes: `$HERE/../..` → `$HERE/../../..` (REPO_ROOT/DEV reach
`rosetta/anthropos-dev` again, in `up-injected.sh`, `migrate-demo.sh`, the `rosetta-demo` CLI); `up-injected.sh`'s
`CLERK` → the monorepo sibling `$HERE/../clerkenstein` (was the now-deleted `$REPO_ROOT/anthropos-demo/clerkenstein`).
`apply-authn.sh`'s `../../clerkenstein` was already correct (`inject/` sits one deeper). **Lesson:** a repo move that
changes nesting depth silently breaks every up-reaching relative path — verify the I/O orchestrators, not just the
unit suites.

## Adversarial review
- **Could the deletion lose history?** Mitigated by verify-then-delete: proved all 66 source commits reachable in the
  monorepo (+ pushed to origin) *before* removing the originals. The deletion order was local-first, then org —
  each step left a redundant copy until the monorepo was confirmed authoritative.
- **Do any consumers still point at the old repos?** Swept rosetta (operational refs → repointed) + the monorepo
  (0 old-path refs) before deleting. The demo-stack tests' `../../clerkenstein` resolves within the monorepo
  (siblings), verified green.

## Open (routed to M5)
- The shared port-offset/multi-instance engine's home (own section vs folded into `stack-injection`).
- (The M4+M5-merge question is moot — M4 closed standalone.)
