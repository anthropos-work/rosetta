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

## M18: Hardening

### Pass 1 — 2026-06-09
**Scope manifest (milestone-touched, all in `rosetta-extensions`):** `stack-verify/lib/target.sh`, `lib/services.sh`, `lib/readiness.sh`, `live/autoverify.sh`, `live/verify.sh`, `repos/run.sh`, `census/inventory.sh`; bring-up wiring `demo-stack/up-injected.sh` + `dev-stack/dev-stack`; test `stack-verify/tests/test_verify.py`. Rosetta `m18/verify-net` side = docs only (no testable code). Coverage tool: bash tooling is path-enumerated via the Python harness (no kcov; matches M16/M17 bash-tooling harden practice) — coverage as candidate-finder, not a numeric gate.

**Bug found + fixed inline:**
- **The readiness phase ignored `STACK_SERVICES`** (`live/verify.sh`). The build phase scoped only the *liveness* phase (`service_rows`); the 6 deep readiness probes ran **unconditionally**, so a reduced bring-up (`--services "postgresql redis"`) false-`down`ed graphql/gotenberg/sentinel/storage — the exact ISSUE-12b wall-of-false-downs M18 exists to prevent (contradicting decision M18-D2). Fix: thread the backing service name into `run_readiness` + gate on `target_service_selected` (mirrors `service_rows`). Commit `2f412a3`.

**Tests added (32 → 60, +28):**
- `TestReadinessScopeFilter` (5) — the regression (reduced bring-up skips absent deep probes) + static fence that verify.sh threads the service names.
- `TestOffsetMatrixSweep` (4) — the **full** port-offset matrix (every base × N), roadrunner's high-base 10400 across the decade boundary (the A1 trap), no host-port collisions, dev-N/demo-N share the offset.
- `TestCrossCheckEdges` (6) — single-in-band-among-junk, all-out-of-band warn, empty ports, malformed registry JSON, missing file, roadrunner-only — all non-fatal.
- `TestTargetHelperBoundaries` (8) — empty/non-numeric N, demo-0, multi-dash, exact-match (skill ≠ skiller), ragged whitespace.
- `TestAutoVerifyEdges` (5) — non-numeric/empty casbin (fail-closed warn), unknown arg, `--offset` override, health-only failure — all exit 0.

**Knowledge backfill:** `corpus/ops/verification.md` — the readiness-phase bullet now states **both** phases honour the `STACK_SERVICES` scope filter (the readiness phase skips an out-of-scope deep probe via the same `target_service_selected` gate). The doc previously implied per-phase scoping that the readiness phase didn't actually do — the fix makes the doc true.

### Pass 2 — 2026-06-09
**No production changes (pure coverage deepening — the code was correct, branches just untested).**

**Tests added (60 → 76, +16):**
- `TestProbeService` (12) — the per-service liveness classifier had **no** direct test. All 4 probe kinds × status branches: docker (healthy / running-no-healthcheck / not-running), tcp (ok / fail), http|http-200 (000 refused / 2xx up / 4xx up-for-plain / http-200 mismatch down / unexpected-code down), unknown-kind fallthrough.
- `TestContainerForProject` (3) — the `${base/#anthropos-/…}` anchored replace swaps only the leading prefix (embedded "anthropos" preserved); main-dev-stack path is a verbatim no-op.
- `TestRunShResolution` (1) — **behavioral** proof of the `$DEVDIR → $STACK_ROOT` fix: run.sh resolves an absent repo to `<STACK_ROOT>/<repo>` (not the old `/<repo>`), driven against a temp STACK_ROOT.

**Knowledge backfill:** no KB-worthy findings (pure test deepening of correct code).

### Pass 3 — 2026-06-09
**No production changes.**

**Tests added (76 → 79, +3):**
- `TestInventoryShResolution` (1) — the inventory.sh half of the `$DEVDIR → $STACK_ROOT` fix had only a static fence. Behavioral proof: a present clone under `$STACK_ROOT` is walked + its planted test counted; an absent repo is `not-cloned`.
- `TestAutoVerifyCrossCheckIntegration` (2) — build tested `target_cross_check` in isolation; this pins the autoverify↔cross-check **integration** end-to-end (a registry/offset mismatch surfaces the warning through autoverify *and* still exits 0; a matching registry is silent).

**Knowledge backfill:** no KB-worthy findings.

### Stop condition
Loop stopped after Pass 3: the full Step 2b re-scan found nothing new worth adding (every milestone-touched seam — offset matrix, both-phase scope filter, cross-check edges, target boundaries, autoverify failure modes, probe_service, container resolution, run.sh + inventory.sh resolution, bring-up wiring — is pinned), coverage deltas negligible, and the flake gate was clean (79 passed × 3 consecutive sequential runs, deterministic). One real bug found + fixed inline (readiness scope filter). Final: `stack-verify/tests/test_verify.py` **32 → 79 (+47)**.
