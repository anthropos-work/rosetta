# M4 — Retro

**Summary:** Collapsed the repo constellation into **two repos** before it exploded — `rosetta` (the platform
corpus + dev-env skills) and a new **`rosetta-extensions`** monorepo holding `clerkenstein/` + `demo-stack/` (with
full history via `git subtree`), a lightweight `knowledge/` nav, and room for the coming `stack-injection` /
`dev-stack` / `stack-seeding` sections. rosetta thinned to pointers; the two old standalone repos removed (local +
org). Zero platform-repo change.

## Incidents this cycle
- **P1 — the +1-depth path break (caught by verify, fixed before deletion).** Moving the tooling one level deeper
  (`anthropos-demo/rosetta-demo/` → `…/rosetta-extensions/demo-stack/`) silently broke every up-reaching relative
  path: `$HERE/../..` (REPO_ROOT/DEV) no longer reached `rosetta/`, and `up-injected.sh` built the fakes from the
  *old* clerkenstein path. The 78 unit tests passed (they don't run the I/O orchestrators); the **verify gate's
  old-path grep** caught it. Fixed (`../../..` + the monorepo sibling), re-verified. **Lesson:** verify the
  orchestrators, not just the suites, after a depth-changing move.
- **P2 — `delete_repo` scope wall.** The `gh` token couldn't delete org repos (missing scope); the user deleted them
  via the web UI. Not a defect — a permissions boundary; surfaced cleanly with the exact unblock command.

## What went well
- **Verify-then-delete discipline** — proved the full history was inside the monorepo (every SHA, both repos) and
  pushed *before* removing anything; deleted local-first then org, each step leaving a redundant copy. No data risk
  despite an irreversible operation.
- **`git subtree` was the right tool** — one command per repo, full history grafted, the cross-repo `../../clerkenstein`
  ref survived (siblings stayed siblings), deploy gate stayed 100%/100% throughout.
- **Adapting the lifecycle honestly** — drove the git-surgery + deletion under direct control instead of
  fire-and-forget background agents; the verify step (the "harden" analogue) is exactly what caught the path bug.

## What didn't / constraints
- The non-operational doc/DNA `ref` mentions of the old paths took a second sweep to fully clean (the first grep was
  `--include`-scoped). All 0 now.
- Couldn't complete the org deletion autonomously (scope) — a one-step hand-off to the user.

## Carried forward → M5
- The shared **port-offset engine** home (own section vs in `stack-injection`) — Fate-3 to M5.
- M5 extracts the reusable `stack-injection` layer next (demo ON / dev OFF), then M6 dev-stack, M7 seeding, M8 recipes.

## Metrics
See [metrics.json](metrics.json). Monorepo: 73 commits (full history of 2 repos preserved), 0 old-path refs, PRIVATE;
demo-stack 78 tests + deploy gate 7/7 green under the new paths; 2 old repos deleted. rosetta: thinned to 2 pointers.
