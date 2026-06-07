# M12 — Retro

**Milestone:** Unified stack registry + first-available-N allocation · **Shape:** section · **Closed:** 2026-06-07
**Branch:** `m12/stack-registry` → merged `--no-ff` into `release/01.30-stack-party` · **Tag:** `stack-party-m12` @ `be1c979`

## Summary

The first milestone of v1.3 "stack party" and the **only non-Go surface to date** — the unified registry is
Python (`stack-core/stack_registry.py`) + shell-CLI wiring, not the snapshot/seeding Go modules. It closes the
v1.1-era collision bug where `dev-N` and `demo-N` both computed `base + N·OFFSET` and collided on every port:
N is now a **single shared resource across both kinds**, allocated first-available from one registry. Built in
3 sections (registry+allocator → both up-paths wired → docs), hardened in 3 passes, closed clean. The headline
race guarantee (no two concurrent `up`s pick the same N) is proven at the **OS-process level** — not just
threads — by a 12-process auto-allocation test plus a close-added 6-process explicit-N collision test.

## Incidents This Cycle

**None.** Zero bugs across the 3 harden passes and the 4 close-phase adversarial scenarios — the registry held
under every probe (cross-process race, stale `.lock`, corrupt registry, empty override, explicit-N contention).
No flakes (5/5 sequential at close; 3 consecutive clean during harden). No regressions. No P2s.

The one *test-quality* finding (not a defect): the pre-existing ERR-trap test died at the *pre*-allocation
guard, so it never exercised the post-allocation `trap … release ERR` — fixed in harden Pass 2 by auto-allocating
to skip the pre-guard, then failing a stubbed `up` after the reservation. The trap was load-bearing and correct;
only the test was shallow.

## What Went Well

- **Design questions resolved cleanly to decisions.** All 3 open questions (registry-of-record vs docker-ps,
  the lock mechanism, manual-stack reconciliation) converged to M12-D1/D2/D3 during build — no thrash, no
  re-litigation at close. The "adds-only reconcile" rule (docker-ps never *subtracts* used-N) is the elegant
  race guard that makes a reserved-but-not-yet-started stack survive the docker-ps lag.
- **Both CLIs stayed symmetric.** dev-stack and demo-stack got the identical allocate→trap→reguard→up→
  untrap→set-ports flow, the same `auto` keyword and `$STACK_REGISTRY` override — no divergence to reconcile.
- **KB-1 converged to truth exactly as planned.** The GUIDE.md "registry assigns N" claim was aspirational at
  design (flagged YELLOW in the Phase 0b KB-fidelity audit); M12 *built* the allocator that makes it true, and
  the §3 doc pass rewrote the prose. The DOC-ONLY-because-it's-the-deliverable pattern worked.
- **OS-process-level race testing.** The contention tests spawn real `python3 stack_registry.py` subprocesses
  on the real `fcntl.flock`, with docker-bin pointed at a nonexistent path so the flock+atomic-write is the
  *only* guard — the strongest possible probe of the production scenario (two independent `up` invocations).

## What Didn't

- **Nothing material.** A clean, well-scoped foundation milestone. The only friction was a non-existent linter
  (`ruff` not installed locally) — fell back to `py_compile` + shellcheck, which sufficed for a stdlib-only module.
- **Bash isn't line-coverable offline**, so the CLI allocator wiring is guarded by shellcheck + the allocation
  contract + the real ERR-trap regression tests rather than a coverage %. Acceptable per the milestone's stated
  test policy; noted so M13+ doesn't expect a CLI coverage number.

## Carried Forward

- **DEF-M10-01 (S3 media blob bytes + cloud SnapshotStore backend)** → **v1.4** (inherited from v1.2, re-scoped
  v1.3→v1.4 by the user on 2026-06-07 during design; parked in `roadmap-vision.md § v1.4 seeds`, signed). M12
  added **zero new deferrals** — every scope item landed Fate 1.
- **The generic `stack-*` skill renames** (`/demo-status` → `/stack-list`, etc.) that *surface* this registry →
  **M14** (already in M14's `In:` list — Fate-2 confirmed-covered, not a deferral).

## Metrics Delta

(Source: `metrics.json`.)

- **Go test funcs:** 708 → **708** (unchanged — M12 touched no Go).
- **Python test funcs (new M12 surface):** **89** across the 3 M12-touched suites — stack-core 54 (registry +
  gen_override; `test_stack_registry.py` 0 → **50**), dev-stack 22 (+1), demo-stack 13 (+1). +52 net new.
- **Coverage:** `stack_registry.py` 92% → **98%** statements (2 residual misses structurally unreachable).
- **Flake:** **0** (5/5 sequential). **Lint:** both CLIs shellcheck-CLEAN; `stack_registry.py` py_compile-CLEAN.
- **Review:** 3 findings (0 scope / 0 code must-fix / 0 docs / 1 test / 2 decision-triage); 4 adversarial
  scenarios recorded; deferral re-audit **GREEN**.
