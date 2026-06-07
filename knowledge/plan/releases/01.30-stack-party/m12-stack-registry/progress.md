# M12 — Progress

**Shape:** section · **Status:** archived (completed 2026-06-07)

## Section checklist (from overview Scope.In)
- [x] Unified stack registry (dev+demo, type/N/ports/status) in stack-core — `stack-core/stack_registry.py` (§1)
- [x] First-available-N allocator (registry + `docker ps` reconcile, lowest free N), race-safe — `allocate()` under flock (§1)
- [x] Up accepts explicit-N OR auto-allocates; teardown frees the slot — both CLIs `up [N]` + `down` release (§2)
- [x] dev-stack + demo-stack bring-up consume the registry — `reg_cli allocate/release/set-ports` (§2)
- [x] Delivers: corpus/ops/rosetta_demo.md (unified registry + first-available-N) — new model section (§3) + stack-core/README.md

## M12: Hardening

**Scope manifest (milestone-touched, code-under-test):**
- `stack-core/stack_registry.py` (NEW) — the unified registry + `allocate()`/`release()`/`set_ports()`/`list_stacks()` + the `docker ps` reconcile + CLI. Tests: `stack-core/tests/test_stack_registry.py`.
- `dev-stack/dev-stack` (shell CLI) — allocator wiring (`reg_allocate`/`reg_release`, ERR trap, `$STACK_REGISTRY`). Tests: `dev-stack/tests/test_dev_stack.py`.
- `demo-stack/rosetta-demo` (shell CLI) — allocator wiring (`ureg_allocate`/`ureg_release`, ERR trap). Tests: `demo-stack/tests/test_tooling.py`.
- Docs (no tests): `stack-core/README.md`, `dev-stack/README.md`, `demo-stack/GUIDE.md`, `knowledge/README.md`, `README.md`.

### Pass 1 — 2026-06-07
**Coverage delta (`stack_registry.py`):**
- Statements: 92% → 98% (8 miss → 2 miss)
- Branches: 9 partial → 2 partial
- Remaining 2 misses (313, 317) are structurally unreachable: the `main()` fallthrough `return 1` (argparse `required=True` guarantees a valid subcommand) + the `if __name__ == "__main__"` guard. Not tested — would manufacture coverage.

**Tests added (`test_stack_registry.py`, +19):**
- 1 concurrency: cross-process `allocate()` — 12 separate OS processes contending on the real `fcntl.flock` produce the distinct set {1..12}, the M12 headline race guarantee at OS-lock level (not just threads).
- edge cases: `_used_n` skips non-int `n`; `_reconcile` drops malformed rows while keeping reserved-but-dark N held; `first_free_n` honours `start`; `_used_n` unions registry ∪ live.
- error/no-op paths: `set_ports` on a missing/malformed record is a silent no-op (no file fabricated); `live_projects` swallows rc!=0 / `OSError` / blank lines to `set()`.
- `list_stacks(reconcile=False)` never consults docker; empty registry → `[]`.
- `$STACK_REGISTRY`/`registry_path` override == fully isolated N-pools + parent-dir auto-create.
- CLI: `set-ports` command; explicit-N=0 → exit 2 (not traceback); missing subcommand → usage error.

**Bugs fixed inline:** none — the registry was correct under every probe.
**Flakes stabilized:** none.

### Pass 2 — 2026-06-07
**Coverage delta:** CLIs are I/O-bound (offline-uncoverable); guarded by shellcheck + the allocation contract per the milestone's stated test policy. No line-coverage tool on the bash.

**Tests added (+2, one per CLI):**
- `dev-stack` + `rosetta-demo`: the **real** ERR-trap-frees path. The pre-existing `test_guard_failure_after_allocation_releases_the_slot` dies at the *pre*-allocation guard (explicit N == dev project), so it never reserves a slot — the post-allocation `trap "…release $n" ERR` was untested. New tests auto-allocate (skipping the pre-guard), then a stubbed `docker compose … up` fails (exit 1) *after* the reservation; assert the trap frees the slot and the freed N is re-allocatable. Delta-verified the trap is load-bearing (N persists on success, vanishes on failed up). Both CLIs stay shellcheck-clean.

**Bugs fixed inline:** none. **Flakes:** none.

### Pass 3 — 2026-06-07
**Coverage delta:** 98% → 98% (behavior coverage, not line coverage).

**Tests added (`test_stack_registry.py`, +2):**
- explicit-N rejected when an **adopted-live** stack (docker-ps-only, no registry row) holds that N — the docker-ps adds-only contract from the explicit-N angle (was only asserted for auto-allocation).
- an adopted record is releasable like any managed one (slot frees on teardown).

**Bugs fixed inline:** none. **Flakes:** none.

**Knowledge backfill:** no KB-worthy findings across the 3 passes — every probe confirmed documented behavior (the M12 decisions D1–D3 + the module's own docstring already capture the race guard, the adds-only reconcile, and the atomic/corrupt-recovery semantics). The one discovery (the prior ERR-trap test not reaching the post-allocation trap) is a test-quality fact pinned by the new regression tests, not a system invariant needing a doc.

### Stop condition
Loop stopped after Pass 3: full six-dimension scan found nothing new worth adding, coverage delta < 2% (98% steady; the 2 residual misses are structurally unreachable), and zero flakes (3 consecutive clean runs of the new tests). 0 bugs surfaced — the implementation held under every stress probe. Test funcs on touched files: `test_stack_registry.py` 28 → **49** (+21), `test_dev_stack.py` 21 → **22**, `test_tooling.py` 12 → **13**.

## M12: Final Review

Close review (2026-06-07). **3 findings total** · 0 scope · 0 code-quality must-fix · 0 docs · 1 test ·
2 decision-triage. Deferral re-audit **GREEN** (`audit-deferrals/deferral-audit-2026-06-07-m12-close.md` —
1 inherited DEF-M10-01 → v1.4 signed/unchanged, 0 repeat, 0 aged, 0 new; M12 added zero deferrals).

### Scope
- [x] All 5 section checkboxes delivered Fate-1; `overview.md` `Out:` items (skill renames → M14, dev
      peers → M13) confirmed-covered, no silent drops. No TODO/FIXME/HACK in any M12-touched file.

### Code Quality
- [x] Both shell CLIs shellcheck-CLEAN; `stack_registry.py` py_compile-CLEAN. Consistent patterns across
      dev-stack + demo-stack (allocate→trap→reguard→up→untrap→set-ports; `auto` keyword; `$STACK_REGISTRY`
      override; dual-view status). No dead code, no cross-module reach-in, no resource leaks. **0 must-fix.**

### Documentation
- [x] Per-unit handbook contract satisfied: `stack-core/README.md` documents `stack_registry.py`; indexed
      in both `README.md` + `knowledge/README.md` with M12 provenance. KB-1 resolved (GUIDE.md "registry
      assigns N" rewritten to the true unified-allocator prose). `corpus/ops/rosetta_demo.md` delivery
      accurate; all relative links resolve. No numeric test-count in any handbook → nothing to reconcile.

### Tests & Benchmarks
- [x] [Phase 2c] Added `test_concurrent_explicit_N_collision_exactly_one_wins` — the explicit-N analogue of
      the cross-process auto-allocation race (6 procs claim `--n 5` → exactly 1 wins, 5 exit-2). Pins the
      flock'd validate-then-reserve under contention (previously only single-threaded). stack-core 53→**54**.

### Decision Triage
- [x] M12-D1/D2/D3 → **archive** (maintainer rationale; the user-facing "why" of the unified registry +
      first-available-N + adds-only reconcile already lives in `corpus/ops/rosetta_demo.md` + `stack-core/README.md`).
- [x] KB-1 → already resolved + documented (the §3 GUIDE.md rewrite); no further blend needed.

### Adversarial review (Phase 2c)
- [x] 4 scenarios recorded in `decisions.md` (concurrent explicit-N collision · stale `.lock` no-wedge ·
      empty-override `set_ports([])` · corrupt-registry recovery) — each verified handled; 1 pinned with a
      new regression test, 3 already covered. **0 bugs surfaced.**
