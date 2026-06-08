# M18 — Progress

_Section checklist. Closure = all boxes land + `/developer-kit:close-milestone` GREEN._

## Deliverables
- [ ] **Offset/project awareness** in `stack-verify` — derives the `demo-N`/`dev-N` prefix + offset ports from the registry.
- [ ] **Service/profile filter** — checks only what was brought up; reduced bring-ups don't false-`down`.
- [ ] **`$DEVDIR` → `$STACK_ROOT` bug fixed** (`repos/run.sh:108`, `census/inventory.sh:75`).
- [ ] **Cheap-win asserts** at bring-up tail — `/api/health` + `casbin_rules > 0` on offset ports, warn loudly (non-fatal).
- [ ] **Auto-wired scoped `verify live`** at bring-up tail — default-on + non-fatal (both `up-injected.sh` + the dev path).
- [ ] **`corpus/ops/verification.md`** authored (the auto-verify contract + offset/scope model).

## Verification
- [ ] A regression test on an offset `demo-N` fixture proves verify targets the right ports (no false `down`).
- [ ] The ISSUE-7 scenario (casbin_rules=0) is caught by the cheap-win assert at bring-up tail.
- [ ] Non-fatal proven: a verify failure warns but does not abort a good bring-up.
- [ ] shellcheck + py_compile clean; flake 0.

## Notes
_(build notes appended here)_
