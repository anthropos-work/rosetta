# M18 — Progress

_Section checklist. Closure = all boxes land + `/developer-kit:close-milestone` GREEN._

## Deliverables
- [x] **Offset/project awareness** in `stack-verify` — derives the `demo-N`/`dev-N` prefix + offset ports from the registry. (`lib/target.sh` + `services.sh` + `readiness.sh`; offset cross-checked against the registry's recorded ports.)
- [x] **Service/profile filter** — checks only what was brought up; reduced bring-ups don't false-`down`. (`STACK_SERVICES` ∩ the SERVICES array.)
- [x] **`$DEVDIR` → `$STACK_ROOT` bug fixed** (`repos/run.sh`, `census/inventory.sh`).
- [x] **Cheap-win asserts** at bring-up tail — `/api/health` + `casbin_rules > 0` on offset ports, warn loudly (non-fatal). (Inside `live/autoverify.sh`.)
- [x] **Auto-wired scoped `verify live`** at bring-up tail — default-on + non-fatal (both `up-injected.sh` + `dev-stack` `cmd_up`).
- [x] **`corpus/ops/verification.md`** authored (the auto-verify contract + offset/scope model); indexed in ops README + CLAUDE.md + cross-linked from rosetta_demo.md.

## Verification
- [x] A regression test on an offset `demo-N` fixture proves verify targets the right ports (no false `down`). (`TestOffsetAwareness`: demo-3 backend → 38082, not 8082.)
- [x] The ISSUE-7 scenario (casbin_rules=0) is caught by the cheap-win assert at bring-up tail. (`TestAutoVerify::test_casbin_zero_is_caught`.)
- [x] Non-fatal proven: a verify failure warns but does not abort a good bring-up. (`TestAutoVerify::test_failing_verify_is_non_fatal` → exit 0.)
- [x] shellcheck + py_compile clean; flake 0. (9 scripts shellcheck-clean; 97 tests across 3 suites; new suite 32/32 ×3 deterministic.)

## Notes
- All 6 sections landed. Extensions tests: stack-verify **+32 net-new** (new `tests/test_verify.py`); demo-stack 27 + dev-stack 38 unchanged + still green (the M16/M17 fences over the edited bring-up scripts held).
- PR-review fix A1 (correctness): the offset/recorded-port cross-check used a broken `port//10000==n` decade lane that would false-warn on roadrunner's high base (10400→20400 for n=1); replaced with a base-band check + 4 regression tests.
- Tag `dress-rehearsal-m18` + push: handled at the build-step tail (extensions on `main` past `dress-rehearsal-m17`).
