# M16 — Progress

_Section checklist. Closure = all boxes land + `/developer-kit:close-milestone` GREEN._

## Deliverables
- [ ] **Push the stranded fixes** — push `547de17`+`ed72e94` to `origin`; re-tag `dress-rehearsal-m16` (retire local-only `stack-party-devpath-fix`); re-consume per-stack; verify the consumed copy matches.
- [ ] **Stack-core rename migration** — `stack-dev` is the documented default; `anthropos-dev` demoted to a single "legacy alias" mention.
- [ ] **Prose sweep** — `demo-stack/README.md:12`, `demo-stack/GUIDE.md:17`, `dev-stack/README.md:73`, `stack-core/gen_override.py:4`.
- [ ] **GUIDE.md header truth** — remote-exists / 13-tests / `/stack-list` / v1.3 (was: no-remote / 78 / `/demo-status` / v1.1·M3).
- [ ] **pytest doc fix** — `pytest tests/ -v` + the 3.11/3.12 note (`demo-stack/GUIDE.md:167`).
- [ ] **`rosetta-extensions/knowledge/` KB** refreshed where it repeats any of the above; version-jump expectation noted (ISSUE-5).
- [ ] **rosetta corpus** — residual `anthropos-dev` swept in `corpus/`; consolidated `stack-dev` layout note (cross-linked from `rosetta_demo.md`).

## Verification
- [ ] `bash -n` / `py_compile` / shellcheck clean on every touched script.
- [ ] pytest suite green via the documented invocation.
- [ ] `grep -rn 'anthropos-dev'` shows only the intentional back-compat-alias mention(s).
- [ ] `origin` and the consumed per-stack copy agree at the new tag.

## Notes
_(build notes appended here)_
