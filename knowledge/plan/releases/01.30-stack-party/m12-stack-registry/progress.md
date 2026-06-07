# M12 — Progress

**Shape:** section · **Status:** planned

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

## Final review
_(filled at close)_
