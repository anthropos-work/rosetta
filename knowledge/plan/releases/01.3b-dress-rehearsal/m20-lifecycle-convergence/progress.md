# M20 — Progress

_Section checklist. Closure = all boxes land + `/developer-kit:close-milestone` GREEN. The closing milestone of v1.3b → then `/developer-kit:close-release`._

## Deliverables
- [x] **Set-dress chaining** — `up-injected.sh` runs the M13 `dev-setdress` pass (now stack-type-aware via `--stack-type demo`) after migrate, before the M18 verify; default-on + non-fatal; `DEMO_NO_SETDRESS=1` escape. ONE engine, two lifecycles. (#M20-D1)
- [x] **Atomicity contract** — seed ALWAYS runs after the cache-first replay (the seed is the floor); a catalog-only stack 403s, a replay miss degrades to structural-only but still logs in 200; retry-safe via the M17 guards. (#M20-D3)
- [x] **Cold-start capture** — net-new `corpus/ops/snapshot-cold-start.md`: the DSN-export / restore-a-`pg_dump`-then-`--dsn` workflow. MCP-adapter spike RESOLVED → **document-only** (the MCP returns JSON rows, not COPY bytes; an adapter would re-serialize COPY format for zero gain). (#M20-D4)
- [x] **demo auto-set-dress preset** — `small-200` (vs dev's `dev-min`), wired as the demo default in the stack-type-aware engine; an explicit `--seed` overrides. (#M20-D2)
- [x] **`corpus/ops/snapshot-cold-start.md`** authored; `safety.md` §2.7 + demo recipes + `demo-up` skill updated; cross-linked from snapshot-spec/db-access/recipe/CLAUDE.md. (`demo-down` needs no change — teardown is unaffected by auto-set-dress.)

## Verification
- [x] A `demo-N` chain wired so a warm cache comes up auto-set-dressed (real catalog + a `small-200` seeded org); the seed-after-replay floor → login 200, not 403. (Engine path proven by the DevSetdress suite against the real script; live demo bring-up = the operator/close-time smoke.)
- [x] Non-fatal proven: a cold cache (replay rc=1) warns + still seeds; behavioral test `test_chain_is_behaviourally_non_fatal_under_set_e` + `test_cache_miss_is_non_fatal_seed_still_runs` (dev+demo).
- [x] Prod-safety held — capture is NEVER on a bring-up (replay-only, pinned by `test_capture_is_never_invoked_on_a_bring_up` + `test_*_replays_both_surfaces_then_seeds`'s `assertNotIn capture`); per-stack Directus firewall-checked; n=0 guard intact; the M15 safety.md drift guards (read+write) stay GREEN after the §2.7 edit.
- [x] Go/py/shellcheck clean; flake 0. (Python 343 pass / 11 live-deselected; Go snapshot+seeding suites green; shellcheck clean on both scripts.)

## Notes
- The "code" is in `rosetta-extensions` (extensions commit `e4d2f9b`, tagged `dress-rehearsal-m20`); the rosetta side is documentation (the cold-start runbook + safety §2.7 + skill/recipe updates). Both committed.
- Tests added: +7 dev-setdress (stack-type/atomicity) + +9 demo chain (7 static fence + 2 behavioral) = +16 (build), then +6 (harden, below) = **+22**.

## M20: Hardening

### Pass 1 — 2026-06-09
**Scope manifest (milestone-touched code, ext commit `e4d2f9b`):**
- `dev-stack/dev-setdress.sh` ← `dev-stack/tests/test_dev_stack.py` (class `DevSetdress` + `DevStackSetdressWiring`)
- `demo-stack/up-injected.sh` (the M20 chain) ← `demo-stack/tests/test_frontend_build.py` (class `TestSetdressChainContract`)
- rosetta branch = docs only (no testable code); the `safety.md` §2.7 edit is pinned by the Go drift guards in `stack-seeding/isolation/safety_doc_drift_test.go` (re-confirmed GREEN).

**Coverage delta (behavioral branch coverage — bash has no native line tool; measured as die/branch-vs-test closure on touched files):**
- `dev-setdress.sh`: every `die`/branch mapped to a test EXCEPT the env-guarded `need go` (L50) — was ~8/11 die-branches pinned → **11/11 meaningful** (the L50 `missing dependency: go` is environment-guarded scaffolding, not Fate-1; shellcheck covers its syntax).
- `up-injected.sh` chain: the SUCCESS path + the resolved-offset-DSN threading went from **static-fence-only → behaviourally pinned**.

**Tests added (+6, all Fate-1 deepening; no production code changed):**
- `dev-stack/tests/test_dev_stack.py` (+5 in `DevSetdress`): `test_non_integer_n_is_rejected_before_any_cli`, `test_missing_n_prints_usage`, `test_trailing_stack_type_flag_dies_with_a_clear_message`, `test_provision_recipe_failure_aborts_before_replay` (the plain-recipe `die`, distinct from the `--check-env` firewall abort), `test_capture_never_runs_even_on_the_cache_miss_degraded_path` (extends the build's happy-path capture-never pin to the degraded branch).
- `demo-stack/tests/test_frontend_build.py` (+1 in `TestSetdressChainContract`): `test_chain_success_passes_clean_and_threads_the_offset_dsn` — behavioural SUCCESS path (no warning + the engine invoked + the resolved `5432+OFFSET` DSN reaches the engine env).

**Bugs fixed inline:** none — the build's logic was correct on every probed path. The harden surfaced no defects, only test-coverage gaps in error/degraded/success branches.

**Mutation pins (proof the new safety tests are meaningful, not line-ticking):**
- Offset-DSN test: mutating the chain's DSN to the base port `5432` (the prod-safety regression) **fails** the new behavioural test while the existing static body fence still **passes** — proof the behavioural test catches what the fence can't.
- Capture-never-degraded test: adding a `stacksnap capture` fallback to the cache-miss branch **fails** the new test. Both scripts restored byte-identical after each mutation.

**Knowledge backfill:** no KB-worthy *new* findings — every invariant the new tests pin (capture-never, offset-isolated DSN, the atomicity seed-floor, the firewall abort) is already documented in `corpus/ops/safety.md` §2.7 + `corpus/ops/snapshot-cold-start.md`. The harden reinforced existing documented invariants as tests; it did not surface undocumented behavior. (Question asked, answer recorded.)

### Pass 2 — 2026-06-09 (confirmation scan, no tests added)
Full six-dimension re-scan over both scripts: test depth, edge cases, and error paths are fully covered after Pass 1; no build-phase `fix`/`bug` commits to regression-pin (the M20 scripts had none); boundary fuzzing and perf benchmarks would be bloat for this small, fully-enumerated integer/enum input surface with no parser/deserializer and no documented latency SLA. No new gap worth a test.

### Stop condition
Scan clean + coverage delta < 2% residual (only the env-guarded `need go`) + 0 flakes. Stopped after 1 deepening pass + 1 confirmation scan — adding more would be test bloat. dev-stack 45→**50**, demo chain 38→**39** (whole-suite collected count rises with the build's already-counted M20 tests; harden net +6). shellcheck clean on both scripts; Go drift guards GREEN.

## M20: Final Review

_Close review 2026-06-09. The cleanest-but-one shape: 0 scope · 0 code · 0 adversarial · 1 docs · 0 tests · 4 decision-triage. Prod-safety invariant verified intact._

### Scope
- [x] All 4 overview In-items → delivered Fate-1 (chaining/atomicity/cold-start/preset); all progress boxes [x]; 0 TODO/FIXME in touched code; 0 gaps.

### Code Quality
- [x] shellcheck + py_compile + bash -n CLEAN on both touched scripts; one-engine-two-lifecycles (no demo fork); `rc=$?` cache-miss logic verified correct. 0 findings.

### Adversarial (Phase 2c)
- [x] 5 scenarios (non-numeric N, trailing-flag under set -u, demo-N=0+force, demo-engine-fork, capture-sneak) — ALL already test-pinned. Record the subsection in `decisions.md`.

### Documentation
- [x] [DOC-1] root `CLAUDE.md` `/demo-up` skill-table row omits the M20 auto-set-dress that the `/dev-up` row advertises — add it (consistency with the convergence narrative). All other docs accurate; 0 broken xrefs.

### Tests & Benchmarks
- [x] Py dev-stack 50 + demo-stack 84 PASS; all 4 Go modules PASS; M15 safety.md drift guards GREEN after §2.7; counts reconciled (Go 736 unchanged, Py 338→360 +22). 0 gaps.

### Decision Triage
- [x] M20-D1 → blend present in `safety.md` §2.7 (one-engine reuse); add `(#M20-D1)` ref-tag.
- [x] M20-D2 → blend present in `demo/README.md` (small-200 default); add `(#M20-D2)` ref-tag.
- [x] M20-D3 → blend present in `safety.md` §2.7 (atomicity floor); add `(#M20-D3)` ref-tag.
- [x] M20-D4 → blend present + already tagged in `snapshot-cold-start.md` (MCP-not-a-capture-source). No action.
