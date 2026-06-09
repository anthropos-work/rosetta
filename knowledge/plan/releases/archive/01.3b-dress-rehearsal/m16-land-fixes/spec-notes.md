# M16 — Spec Notes

_Technical notes accumulated during build — mechanisms, file paths (with line cites), gotchas, and the concrete shape of each change. Populated by `/developer-kit:build-milestone`. The verified code locations from the design-time research are in the milestone `overview.md` and `.agentspace/demo-up-issue.md`._

## Pre-flight audits — Push the stranded fixes (first section)
- **Phase 0b KB-fidelity verdict: GREEN** (SEVERITY clear). Report: `kb-fidelity-audit.md` (this dir).
- 6 STALE findings — all are M16's own enumerated deliverables (78→13 tests; "no remote"→remote exists; `/demo-status`→`/stack-list`; v1.1/M3→v1.3; `anthropos-dev/platform`→`stack-dev`). 2 net-new completeness items (corpus layout note; version-jump note) — both M16 deliverables. 0 blind areas, 0 open items.

## Topic → doc → code triples (verified at audit time, sha 0e04d38)
- **stack-dev layout + back-compat fallback** → `corpus/ops/rosetta_demo.md` (+ CLAUDE.md) → `.agentspace/rosetta-extensions/{demo-stack/up-injected.sh:18, migrate-demo.sh:13, rosetta-demo:24/159; dev-stack/dev-stack:24}` (the `stack-dev → anthropos-dev` fallback — the intentional single legacy-alias). `clone_repos.py:96` help already says "stack-dev/ … legacy: anthropos-dev/".
- **The two applied fixes** → `.agentspace/demo-up-issue.md` ISSUE-1/7 → extensions commits `547de17` (devpath), `ed72e94` (migrate-race). Local-only tag `stack-party-devpath-fix` @ `ed72e94`. Authoring copy is 2 ahead of `origin/main` (`a31d70b`).
- **demo-stack header facts** → `demo-stack/GUIDE.md:5,161,164` → `demo-stack/tests/test_tooling.py` (13 collected/pass; `test_inject_scripts.py` GONE), `origin` remote exists.
- **pytest invocation** → `demo-stack/GUIDE.md:166-167` → `pytest` entrypoint (this box: `python3`=3.9.6 no pytest module; `pytest`=8.4.2 works).
- **prose `anthropos-dev/platform` (default)** → `demo-stack/GUIDE.md:17`, `demo-stack/README.md:12`, `dev-stack/README.md:73`, `stack-core/gen_override.py:4`.
- **version-jump note** → `.agentspace/rosetta-extensions/knowledge/README.md` (net-new).
- **corpus residual sweep** → `corpus/` already has **0** `anthropos-dev` refs (sweep is a verify no-op).

## Consumption inventory
- Only one per-stack consumption clone with a git repo: `stack-demo/rosetta-extensions` @ tag `stack-party-devpath-fix` (= `ed72e94`). `stack-dev/` has **no** rosetta-extensions clone. → after push+re-tag, re-consume only `stack-demo/rosetta-extensions` to `dress-rehearsal-m16`.
- Auth: `gh auth status` logged in as `kiralise` (scopes incl. `repo`, `workflow`); `git ls-remote origin` succeeds → push expected to work.

## Publish result (the outward-facing action — SUCCEEDED)
- **extensions commits on `main`** (since `origin` baseline `a31d70b`): `547de17` (devpath fix, pre-applied) · `ed72e94` (migrate-race fix, pre-applied) · `f1ee81a` (M16 doc truth-up) · `44edc09` (M16 PR-review README sweep).
- **`git push origin main`**: `a31d70b..f1ee81a` then fast-forward `f1ee81a..44edc09`. **No branch force-push.**
- **Tag `dress-rehearsal-m16`**: created at `f1ee81a`, then moved forward to `44edc09` (tag-ref force only) to capture the PR-review sweep. Pushed. Annotated tag object `23bddee` → commit `44edc09`.
- **Local-only `stack-party-devpath-fix` tag**: superseded (unused), not deleted, never pushed.
- **Re-consumed**: `stack-demo/rosetta-extensions` checked out at `dress-rehearsal-m16` (`44edc09`). Three-way agreement: origin tag→commit = authoring HEAD = consumed HEAD = `44edc09`.
- **Verified in consumed copy**: `up-injected.sh:18` stack-dev resolve + `migrate-demo.sh:60` `|| echo 0` race guard present; corrected GUIDE (13 tests) + README present; consumed scripts shellcheck-clean.

## PR-review findings (all Fate-1, landed in `44edc09`)
- `demo-stack/README.md` carried 5 same-class stale facts the first sweep missed (it used `anthropos-demo/`, not `anthropos-dev/`): stale title/version label; the `anthropos-demo/ ... own git` consumption claim (no remote); **two** `78 unit tests (55+23)` repeats; and a pre-M5/M6 `lib/` layout list (`gen_override.py`/`gen_injected_override.py`/`inject.py` were extracted to `stack-core/` + `stack-injection/`; `lib/` now holds only `clone_repos.py`) + a stale `lib/gen_override.py` how-it-works path.
- All cited paths in the corrected README verified to exist. `clerkenstein/knowledge/glossary.md:19` `anthropos-demo` mention left untouched (clerkenstein section's own KB, out of M16 scope — M16-D5).
