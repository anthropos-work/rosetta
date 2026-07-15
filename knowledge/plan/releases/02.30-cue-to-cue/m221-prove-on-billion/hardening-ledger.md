# Hardening Ledger — M221 prove-on-billion

Milestone shape: **iterative** (iters 01–06, gate MET on the iter-06 cold r4 cycle).
Final harden pass runs cumulative-scope across all milestone-touched code as the last work before
`/developer-kit:close-milestone`. rext authoring clone: `.agentspace/rosetta-extensions` (branch `main`).

## Pass 1 — 2026-07-15 — final

**Iters hardened this pass:** all milestone-touched code (cumulative --final scope).
**Tiks covered since prior pass:** all iters in milestone (first + only harden pass — no prior ledger).
**Scope manifest (milestone-touched, `6a75b5b..HEAD` in rext + the rosetta plan dir):**
- demo-stack: `hostlock.sh`, `reap.sh`, `ant-academy.sh`, `backend_api_url_server_reader_guard.py`, `up-injected.sh`, `rosetta-demo` + tests `test_host_isolation.py` / `test_reap.py` / `test_ant_academy.py` / `test_backend_api_url_twin.py`
- stack-injection: `exposure_claim_guard.py` + `tests/test_exposure_claim_guard.py`
- stack-snapshot: `cmd/stacksnap/main.go` (F1 `workspaceRootFrom` + loud-empty diagnostic) + `main_m221_test.go`
- dev-stack: `dev-setdress.sh` + `tests/test_dev_stack.py`

**Fence RED/mutation-proof audit (dimension 3+4 — every fence M221 shipped):**
- host-lock (`test_host_isolation.py`): 4 fences × targeted mutants (check-then-write TOCTOU, generic-refuse, silent-reclaim, no-rm-release) + atomic-race (24 acquirers → exactly 1 winner) + wiring order. **Load-bearing, honest.**
- academy loopback-bind (`test_ant_academy.py::TestAntAcademyPublicHostBind` + `test_exposure_claim_guard.py::HostNativeBindTest`): RED anchor (`bind_args=(-H 127.0.0.1)` vs pre-fix bare `bind_args=()` → next-dev 0.0.0.0) + derivation reads the tool default + `assertNotIn('bind_args=();')`. **Load-bearing, honest.**
- host-native exposure-guard extension (`exposure_claim_guard.py`): RED-proof vs pre-fix bind + unparseable/empty → return 2 (FINDING, not pass). **Load-bearing, honest.**
- native supervisor reap (`test_reap.py`): mutation (excise the reap block → stale academy wins) + supervisor→worker tree + **detached** supervisor by port-anchored identity + foreign-parent-never-walked + empty-pattern refused. **Load-bearing, honest.**
- backend-url server-reader guard (`test_backend_api_url_twin.py`): mutation vs REAL source (strip `'use client'` from a shipped app-router page → flagged) + zero-readers → return 2. **Load-bearing, honest.**
- F1 store resolver + loud-empty diagnostic (`main_m221_test.go`): empty-subdir-shadow RED-fence + back-compat + loud `st.List()==0` wrong-root diagnostic. **Load-bearing, honest.**

**Bugs surfaced + fixed inline (Fate 1):**
- **SUITE HONESTY (false green):** `demo-stack/tests/test_reap.py` had its `if __name__ == "__main__": unittest.main()` block MID-FILE, with 5 TestCase classes defined below it. `python3 tests/test_reap.py` ran only 21 of 41 tests and printed "Ran 21 tests ... OK" — a false all-clear silently omitting the 20 adversarial-error-path + suite-honesty fences (the exact tests that fence this milestone's reap work). Moved the block to the true end of file. Verified: direct run **21 → 41**; pytest (canonical path) unchanged at 41. (commit `a0f8615`) — D17: a status line read as evidence after the thing it described changed.
- **F-M221-06b (routed-forward residual, landed):** `stack-verify/e2e/run-latency.sh` hardcoded `http://` for the cockpit URL → 400 against M220's HTTPS-fronted `--public-host` cockpit. Added `LATENCY_SCHEME` env (default `http`; localhost unchanged; invalid → refuse rc2). shellcheck clean + construction smoke-test (default→`http://localhost:17700`, `LATENCY_SCHEME=https`→`https://…:17700`). (commit `a0f8615`)

**Tests added (dimension 2 — edge case):**
- stack-snapshot → `main_m221_test.go`: `TestReplay_StoreWithEmptySurfaceSubdirIsStillLoud` (1 Go edge test) — the "exists != populated" hazard one directory level DEEPER than the iter-05 residual (`snapshots/<surface>/` empty from an interrupted capture). Proves the load-bearing invariant holds: the deeper shadow trips the replay-time `st.List()==0` net and is LOUD/non-fatal, so it can never degrade the catalog SILENTLY.

**Coverage delta on touched files:** no single cross-language number (Go + bash + python tooling; coverage used as a finder per the skill). Concrete: F1 resolver path gains a depth-2 edge test; `run-latency.sh` scheme path gains a construction proof; `test_reap.py` direct-run execution 21 → 41 tests (honesty, not new source coverage).

**Flakes stabilized:** none surfaced (see Pass 2 flake gate).

**Knowledge backfill:** none required — findings are test-hygiene/tooling; the `LATENCY_SCHEME` knob is self-documented in `run-latency.sh`'s header (where an operator running it looks); `tailscale-serve.md` has no run-latency reference to cross-link, so adding one would be out-of-mandate scope. The D17 `__main__`/false-green instance is recorded for the retro.

**Stop condition:** continue-to-next-pass — 3 Fate-1 items landed this pass; a confirmation sweep is needed to measure "no new findings" against them.

## Pass 2 — 2026-07-15 — final

**Iters hardened this pass:** all milestone-touched code (confirmation sweep after Pass 1's fixes).
**Work:** re-ran all touched suites post-edit; second honesty scan for the X-tests-not-X / grep-proves-nothing patterns; flake gate on the new/edited tests.

**Second honesty scan — nothing new surfaced:**
- Swept all 6 M221-touched test files for a misplaced `__main__` (only `test_reap.py` had it — fixed Pass 1; the other 5 have the block correctly at EOF).
- `test_dev_stack.py::test_store_root_prefers_a_cache_bearing_ancestor…` reproduces the resolver in a fixture BUT drift-locks it against the shipped `dev-setdress.sh` (`assertIn('[ -n "$(ls -A …)" ]', SETDRESS)` — "keep the shell copy in lock-step or this fence is a lie"). Honest.
- The static source-scan fences that could match their own docs strip comment lines first (`test_reap.py::TestNoSilentSwallows`, `test_reap.py::test_the_compose_range_preflight_is_WIRED`, `exposure_claim_guard` fence-nesting). The "grep-for-a-call-proves-nothing" class is explicitly fenced by `test_reap.py::TestUpInjectedActuallyHasTheFunctions` (resolves the function at runtime, not just greps the string).

**Suites GREEN post-edit:**
- stack-snapshot Go: all 13 packages ok (uncached, `-count=1`).
- demo-stack pytest: 650 passed / 4 skipped / 9 subtests (identical to pre-edit — the `__main__` move is invisible to pytest).
- dev-stack pytest: 118 passed / 4 skipped in 111 s — the "spins forever" ghost (discharged M220) STAYS discharged; no hang.
- stack-injection exposure: 49 passed.

**Flake gate (3 consecutive clean runs of new/edited tests):**
- new Go edge test (`Replay_StoreWithEmptySurface…` + F1 neighbours): 3/3 clean.
- reordered `test_reap.py` (process-spawning): 3/3 clean (41 passed each: 28 s / 31 s / 42 s).

**Coverage delta on touched files:** < 2% (Pass 2 added no tests — confirmation only).

**Flakes stabilized:** none — no flake observed across the gate.

**Knowledge backfill:** none.

**Stop condition:** stabilized — no new findings + all touched suites green + flake gate 3/3. Final-mode pass complete; unblocks `/developer-kit:close-milestone`.
