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

## Pass 2 — 2026-07-08 — final

**Iters hardened this pass:** all milestone-touched code (final cumulative scope). Targeted the two holes Pass 1 flagged.
**Tiks covered since prior pass:** all iters in milestone.
**Coverage delta on touched files:**
- `dev-stack/migrate-dev.sh` (iter-17): static-only → **+ LIVE behavior harness** (4 docker-gated tests). The cold-DB init is now BEHAVIORALLY proven (not just statically fenced) — the same live coverage class its mirror `migrate-demo.sh` already had.
- `stack-verify/e2e/lib/coverage-manifest.ts` + `tests/coverage.spec.ts` (iter-13): **0 → 3 unit drift guards** on the cross-port academy routing (the fix had landed only in the live sweep hook + an exported descriptor, with no drift guard on either).
**Tests added:**
- iter-17 → `dev-stack/tests/test_migrate_dev_live.py` (NEW, 4 docker-gated tests): survives the casbin race + seeds the policy, actually bootstraps the M25-D9 `extensions` schema + pgvector/pg_trgm/pgcrypto LIVE, re-run idempotent (no duplicate policy), non-fatal with absent `DEV_SENTINELC`/`DEV_BACKENDC`. Runs the REAL script against a throwaway pgvector container via env overrides (never the dev stack; all psql via `docker exec`, no host port). Ran green in ~11s on this box.
- iter-13 → `stack-verify/e2e/tests/coverage-manifest.unit.spec.ts` (+3): the `isAcademyPort` predicate in lock-step with coverage.spec.ts's `port % 10000 === 3077`, the `ANT_ACADEMY_HOME_SECTION` distinctness from `STUDIO_DESK_HOME_SECTION` (the false-FAIL the fix prevents), and the academy home real-content floor (main→body region + positive text floor — the gate is not weakened).
**Bugs surfaced + fixed inline:** none — correct-but-untested surfaces.
**Flakes stabilized:** the migrate-dev live harness now behaviorally fences the casbin-race survival + the non-fatal-absent-sentinel path under `set -e` (the "migrate race" the orchestrator flagged) — previously proven only by hand.
**Stop condition:** continue-to-next-pass — one symmetry edge remains: next-web's unreadable-endpoint→reuse arm is tested for studio-desk but not next-web.

## Pass 3 — 2026-07-08 — final

**Iters hardened this pass:** all milestone-touched code (final cumulative scope). Filled the last edge + ran the stabilization scan.
**Tiks covered since prior pass:** all iters in milestone.
**Coverage delta on touched files:**
- `demo-stack/up-injected.sh` next-web reuse guard (iters 09/10/12): **+1 edge** — the unreadable-endpoint→reuse arm (the `[ -n "$got_ep" ]` fall-through), previously covered only for studio-desk.
**Tests added:**
- iters 09/10/12 → `demo-stack/tests/test_frontend_build.py` (+1): `test_tag_guard_reuses_next_web_when_baked_endpoint_is_unreadable` — a present image with an unreadable baked endpoint (an old pre-offset-build-args image) must REUSE, never rebuild/`image rm` (destructive), and log `(unverifiable)`.
**Bugs surfaced + fixed inline:** none.
**Flakes stabilized:** none.
**Stop condition:** continue-to-next-pass — the final scan of the remaining M211-touched surfaces found NO further holes: the iter-15 demopatch URL re-pin (2 YAMLs) is covered by `test_demopatch.py` (49 tests, the hash-loader path), the iter-10 build-scratch re-sync by `test_injected_build_scratch_is_resynced_each_bringup`, the iter-12 academy-URL bake by `test_next_web_bakes_demo_local_academy_url`, and the cross-iter up-injected 09/10/12 integration by the tag-guard + scratch-resync + academy-bake test set. Pass 4 = full verification + the 3× flake gate → expect stabilized.

## Pass 4 — 2026-07-08 — final

**Iters hardened this pass:** all milestone-touched code (final cumulative scope). Verification + stabilization measurement pass — no new tests (the Phase 2 dimension scan found nothing new across all M211-touched surfaces).
**Tiks covered since prior pass:** all iters in milestone.
**Coverage delta on touched files:** 0 — the scan re-swept the milestone footprint (the shell tooling `migrate-dev.sh` / `run-playthroughs.sh` / `up-injected.sh`, the `playthroughs/manifest` Go gate, the `coverage-manifest.ts` TS routing, the demopatch YAML) and found every branch already fenced by passes 1–3 + the pre-existing suites.
**Tests added:** none (measurement/verification pass).
**Verification run (Phase 5), all in the target trees:**
- Go: `go test ./manifest/` **ok** (includes the 4 `TestRunnerSafety_*` gates incl. the new `RosterRefreshGate`); `go vet` clean.
- Python demo-stack: `test_frontend_build` **65 tests OK** (64 → 65 with the next-web unreadable-reuse edge).
- Python dev-stack: `TestMigrateDevStaticFence` + `TestShellcheck` + `test_migrate_dev_live` **18 tests OK** (11 static + 3 shellcheck + 4 live docker).
- TS: `coverage-manifest.unit.spec.ts` **32 passed** (29 → 32 with the academy routing/descriptor pins).
- `shellcheck` + `bash -n` clean on `migrate-dev.sh` + `run-playthroughs.sh`.
- **Flake gate:** `TestMigrateDevLive` (the only runtime-nondeterministic new test — it spins throwaway containers) ran **3 consecutive clean** (~9–10s each). The static/logic new tests are deterministic.
**Bugs surfaced + fixed inline:** none across the whole session.
**Flakes stabilized:** none needed (the live harness was stable on the first 3× gate).
**Knowledge backfill:** none — this harden added regression fences, not new behavior/semantics; the corpus already documents the M211 tooling (`corpus/ops/demo/playthroughs.md` the iter-16 roster refresh; `corpus/ops/setup_guide.md` + `.claude/skills/dev-up/SKILL.md` the iter-17 `migrate-dev.sh` cold DB-init).
**Out-of-scope (unchanged, not chased — per orchestrator + close-review caveat):** the pre-existing `test_dev_stack.py` CLI-subprocess failures (an incomplete local `.agentspace/secrets` source trips the secret-coverage pre-flight; unrelated to M211 — the standalone `migrate-dev.sh` static+live tests are green regardless). Left for CLOSE to adjudicate.
**Stop condition:** stabilized — coverage delta < 2% (0 new tests this pass) AND the Phase 2 dimension scan found nothing new; full suites green; flake gate 3× clean.
