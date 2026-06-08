# M16 — Progress

_Section checklist. Closure = all boxes land + `/developer-kit:close-milestone` GREEN._

## Deliverables
- [x] **Push the stranded fixes** — pushed `547de17`+`ed72e94` (+the M16 doc commits) to `origin`; tagged `dress-rehearsal-m16` (local-only `stack-party-devpath-fix` superseded, not deleted); re-consumed `stack-demo/rosetta-extensions`; consumed copy matches.
- [x] **Stack-core rename migration** — `stack-dev` is the documented default everywhere; `anthropos-dev` demoted to the single intentional "legacy alias" (the back-compat fallback line in the scripts + the `clone_repos.py` help).
- [x] **Prose sweep** — `demo-stack/README.md`, `demo-stack/GUIDE.md`, `dev-stack/README.md`, `stack-core/gen_override.py` docstring → `stack-dev/platform`.
- [x] **GUIDE.md header truth** — remote-exists / 13-tests / `/stack-list` / v1.3 (was: no-remote / 78 / `/demo-status` / v1.1·M3).
- [x] **pytest doc fix** — `pytest tests/ -v` + the 3.11/3.12 note (`demo-stack/GUIDE.md`).
- [x] **`rosetta-extensions/knowledge/` KB** — version-jump expectation + per-milestone tag scheme noted (`knowledge/README.md`, ISSUE-5).
- [x] **rosetta corpus** — `corpus/` had 0 stray `anthropos-dev` (sweep = verify no-op); consolidated `stack-dev` layout + back-compat note added to `corpus/ops/rosetta_demo.md` (the doc itself = the cross-link anchor).

## Verification
- [x] `bash -n` / `py_compile` / shellcheck clean on every touched script (4 shell scripts + `gen_override.py`).
- [x] pytest suite green via the documented invocation — `pytest tests/ -v` → **13/13**.
- [x] `grep -rn 'anthropos-dev'` shows only the intentional back-compat-alias mentions (5 script fallback lines + 1 help-text in extensions; the explanatory note in corpus).
- [x] `origin` and the consumed per-stack copy agree at the new tag — origin tag→commit = authoring HEAD = consumed HEAD = `44edc09`.

## Notes
- Extensions-side work lands as commits in the SEPARATE `.agentspace/rosetta-extensions` repo (gitignored from rosetta) — see spec-notes "Publish result". The rosetta `m16/land-fixes` branch carries only the corpus note + these tracking docs.
- PR review surfaced 5 same-class stale facts in `demo-stack/README.md` (missed by the first sweep — they used `anthropos-demo/`); all landed Fate-1 in extensions commit `44edc09`.
- Behavior/idempotency work (ISSUE-11/14) + frontend (ISSUE-8/9) are M17+ — not pulled in.
