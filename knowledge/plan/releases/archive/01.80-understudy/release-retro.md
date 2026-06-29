# Release Retro: v1.8 "understudy"

**Shipped:** 2026-06-15 · **Tag:** `v1.8` · **Milestones:** M26 (single `section`)
**Theme:** the self-contained-demo release — `stack-demo/` gets its own platform clone set so a box with only
`stack-demo/` (no `stack-dev/`) runs a demo end-to-end.

## What shipped
A demo now builds **entirely** from `stack-demo`'s own clones. The headline gap is closed: `up-injected.sh` no
longer borrows `stack-dev`'s code (`src="$DEV/$svc"`, `PLAT="$DEV/platform"`) — a new `ensure-clones.sh` bootstraps
`stack-demo`'s peer clone set from GitHub, the build SOURCE + the compose dir moved to `stack-demo` (D-MAIN), and
dev-image reuse became an explicit `--reuse-dev-images` opt-in. The long-standing `CLAUDE.md` "true peer with its
own clone set" claim is now true in code. v1.6 (stack-secrets/M30) + v1.7 (M31 mkcert, M32 studio-desk) preserved.

## Incidents this cycle
**P0/P1: none.** No production bug surfaced across build/harden/close. The two milestone-close findings were both
documentation/comment-accuracy (Fate-1):
- The D-MAIN PLAT path-move left two stale "legacy stack-dev/.env base" comments in `up-injected.sh` (code correct,
  comments lied) — re-worded.
- `safety.md` lacked the M26 sanctioned cross-stack-read invariant (ensure-clones' `.env` *seed* is the sole
  `stack-dev` read, never the build SOURCE) — blended into §2.7 (#M26-D4).

Harden found **no production bugs** — both harness iterations during the +15 functional/shell-seam tests were
fixture bugs (macOS realpath depth + shebang PATH), fixed in the test.

## Cross-milestone patterns
- **Single-milestone release** — no cross-milestone patterns to surface. The re-implementation-not-merge call (the
  orphan predated v1.6/v1.7 and would have reverted them) was the defining design decision; verified by a
  3-agent + adversarial spec workflow before any code, which is why the build landed clean.
- **Doc-coherence lag at a path-move** recurred as a *class*: a structural change (D-MAIN PLAT → stack-demo;
  dropping the anthropos-dev fallback on demo scripts) left adjacent prose stale in 4 spots (caught at
  close-release Phase 3, all Fate-1: the rosetta_demo.md M16 list, the demo-up SKILL ensure-clones gap, the
  "legacy stack-dev/.env base" claim in two skill docs). Lesson: when a milestone moves a load-bearing path, sweep
  the *prose that names the old path* in the same pass, not just the code comments.

## Carry-forward
- **User-authorized live field-bake** on a freshly-emptied `stack-demo/` — the user chose "straight to
  close-release"; M26 satisfied the observable-behavior gate by composition (M31/M32 precedent). The live run
  (ensure-clones really clones from GitHub → all images build from stack-demo → verified UP, no stack-dev in the
  build graph) remains an optional confirmation, not a release blocker.
- Backlog unchanged + orthogonal: M33 ant-academy demo liveness (repro-first), DEF-M10-01 (cloud store/S3 blob
  bytes), DEF-M21-01 (replayCmd hermetic test), M25-D9 (dev taxonomy rc=4) — see `roadmap-vision.md`.

## Metrics delta (vs v1.7)
- **Go:** 1027 → **1027** (+0 — zero `.go` touched).
- **Python:** 471 → **501** (+30 on the two touched suites: demo-stack 110→138, stack-injection 111→113).
- **Flake:** 0 (triple-clean 3/3 + milestone 5/5 each).
- **Supply-chain:** GREEN — 0 new deps (stdlib + shell + docs).
- **Alignment:** 100%/100% on all 4 Clerkenstein surfaces (untouched).

## Stats delta
Phase 8c snapshot: `knowledge/journal/stats/2026-06-15.json` (5th snapshot; prior was v1.5-close — the v1.6/v1.7
closes didn't snapshot). Valid metrics: **399 commits** (51 merges, 332 co-authored), **+65 since v1.5-close**, 148
in the last 7 days, churn 25%, **39/39 milestones done**, project age 144 days. **Known limitation:** the script's
code/doc **line** auto-detect reports 0 for this repo — the executable code lives in the gitignored
`.agentspace/rosetta-extensions/` (so it's invisible to a repo-root scan) and the corpus/`knowledge/` markdown tree
isn't matched by the script's doc-layout heuristic. The real net change this release (from `git diff`): ext
**+985/−161** (one new 115-line shell unit + repoints + the +30 tests); rosetta **+1117/−79** (doc-half + planning +
records). Tooling + docs only — zero platform-repo edits. (The 0-line auto-detect is a pre-existing script-vs-layout
gap, not a v1.8 regression — same 0s in the prior snapshots.)
