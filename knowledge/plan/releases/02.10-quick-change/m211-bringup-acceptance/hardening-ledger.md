# Hardening Ledger — M211 Bring-up acceptance

M211 is a two-repo, tooling-heavy iterative milestone: the touched code is almost entirely in the
rext authoring copy (`.agentspace/rosetta-extensions/`) — the bring-up + demo/playthroughs shell/TS
tooling fixed across iters 09–17. Test-deepening commits land in the **rext** repo; this ledger + any
doc backfill land on the **rosetta** `m211/bringup-acceptance` branch.

**Coverage instrumentation note.** The M211 delta is shell scripts (`migrate-dev.sh`,
`run-playthroughs.sh`, `up-injected.sh`), a Go doc-drift package (`playthroughs/manifest`), TS harness
libs (`coverage-manifest.ts`), and demopatch YAML. Shell scripts are not line-instrumented (the tests
are static-string drift fences + `bash`/`shellcheck` subprocess + live docker harnesses); coverage is
therefore reported as **guarded invariants / files-brought-under-test**, not stmt/branch %, except for
the Go package (which `go test -cover` measures).

## Pass 1 — 2026-07-08 — final

**Iters hardened this pass:** all milestone-touched code (final cumulative scope; iters 01–17). This
pass targeted the two ZERO-coverage holes surfaced by the cumulative sweep.
**Tiks covered since prior pass:** all iters in milestone (first harden pass — ledger did not exist).
**Coverage delta on touched files:**
- `dev-stack/migrate-dev.sh` (NEW, iter-17): **0 → 11 static invariants fenced + added to the shellcheck sweep** (was in NO test). Mirror `migrate-demo.sh` already had static + live coverage; this file had neither.
- `playthroughs/e2e/run-playthroughs.sh` (iter-16 roster-refresh delta): **0 → 8 drift-gate assertions** on the roster-refresh block (its sibling concerns — reset contract, sentinel reload, reporter override — were already gated; the M211 roster refresh was not). `playthroughs/manifest` Go package stays fully covered.
**Tests added:**
- iter-17 → `dev-stack/tests/test_dev_stack.py`: `TestMigrateDevStaticFence` (11 tests) + `TestShellcheck.test_migrate_dev_is_shellcheck_clean` (1). Fences: strict-pipefail precondition, casbin-count `|| echo 0` set-e resilience, init_policy empty-guard, `wait_pg`/`pg_isready`/`wait_sentinel_running`, schema/extension `|| log` guard, the M25-D9 extensions (vector/pgcrypto/pg_trgm in `extensions`), the 4 merged services + `skiller` never a migrate pair (M209), the absent-clone skip, the isolated cold-proof env overrides (`DEV_PGC`/`DEV_PGPORT`/`DEV_CLONES`/…), and a cross-tooling parity pin vs `migrate-demo.sh`.
- iter-16 → `playthroughs/manifest/runner_safety_test.go`: `TestRunnerSafety_RosterRefreshGate` (8 assertions). Fences: roster re-export from THIS seed, same-`$SEED` constraint, fake-FAPI/BAPI restart, `docker inspect` mount discovery, the `|| echo 000` set-e curl guard in the FAPI-readiness poll (the exact iter-16 flake fix), non-fatal-for-a-roster-native-demo, after-the-reset ordering, and the `unknown_identity` rationale-comment pin.
**Bugs surfaced + fixed inline:** none — the M211 tooling behaves correctly; these were untested-but-correct surfaces (the iter loop's per-symptom discipline never swept them). The additions are pure regression fences.
**Flakes stabilized:** none this pass (Pass 2 exercises the migrate race live).
**Stop condition:** continue-to-next-pass — two known holes remain: (1) `migrate-dev.sh` has no LIVE behavior test (its mirror `migrate-demo.sh` has `test_migrate_race_live.py`; iter-17 proved migrate-dev only by hand); (2) the iter-13 `ANT_ACADEMY_HOME_SECTION` cross-port destination descriptor has no unit drift guard.
