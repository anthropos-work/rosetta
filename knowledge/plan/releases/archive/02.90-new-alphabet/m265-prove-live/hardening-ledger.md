# Hardening Ledger — M265 Prove it live

## Pass 1 — 2026-08-16 — final

**Iters hardened this pass:** all milestone-touched code (final mode, cumulative scope)
**Tiks covered since prior pass:** all iters in milestone (1 tik; no prior pass — ledger created here)

**Scope manifest.** 7 files in `rosetta` (6 md + 1 json — docs/planning, whose test surface is the
guard/fence family) and 19 files across 5 `rosetta-extensions` sections (the code, shipped as tags
`v2.9.10-rext` → `v2.9.17-rext`). Cross-iter integration check: **N/A — single-iter milestone**, so
there are no multi-iter files whose fix interactions need pinning.

**Coverage delta on touched files:**
- `stack-snapshot/cmd/stacksnap`: 75.8% -> 77.0% stmts
- `stack-snapshot/realign`: 92.4% (unchanged — the iter tested this package thoroughly; the gap was
  around it, not in it)
- `stack-core/seed_role_guard.py`: no test file -> 10 tests

**Tests added:**
- iter-01 -> `stack-snapshot/cmd/stacksnap/realign_adapter_test.go`: 4 unit + 2 compile-time
  interface assertions + 1 call-site-shape assertion
- iter-01 -> `stack-core/tests/test_seed_role_guard.py`: 10 unit (4 reach, 5 verdict, 1 contract)

**What the coverage-as-a-finder scan surfaced.** The iter wrote a full suite for the package it was
fixing (`realign`, 92.4%) and none for the **seams it added around it**: `pg.QueryPairs` and
`realignAdapter` were both net-new with **0 test references**, and `seed_role_guard.py` shipped with
no test file while 38 sibling guards have one.

The most load-bearing new test is the pair of compile-time assertions on `realignAdapter`. Losing
`ListColumns` does NOT break the call site — `realign.Run` still accepts the value as an `Execer` and
then refuses at RUNTIME with a message about verification. The assertion converts that into a
build-time failure, and a mutation check confirms it: renaming the method fails the test binary with
*"realignAdapter does not implement realign.ColumnLister (missing method ListColumns)"*.

**Bugs surfaced + fixed inline:**
- `seed_role_guard`'s role regex matched only the sibling-key pin shape (`role: X`) and **missed the
  YAML list-item shape** (`- role: X`). Every seed in the repo uses the first shape today, so it was
  invisible — but for this guard a missed pin is not a smaller result, it is a **false GREEN**: a
  positive claim that no seed pins a retired role. Widened, with the reason recorded at the regex.
  Fixable-inline: single subsystem, ~8 lines, and a direct corollary of the test that found it.

**Flakes stabilized:** none (no flakes surfaced; the Go suites and the guard suite are deterministic
and were run repeatedly during the pass).

**Knowledge backfill:** none this pass. The protocol doc (`corpus/ops/verification.md`) already
carries the M258 batch-gate lineage this milestone followed, and iter-01's lessons were recorded in
the iter's own `progress.md` § Lessons rather than generalised — they are about *verifier design*
(a verifier narrower than its subject reports clean; a probe's denominator must share the scope of
what it qualifies), which is broader than this protocol and belongs to the corpus, where §6.4/§6.5 of
`taxonomy-canon.md` already carry them.

**Stop condition:** continue-to-next-pass — pass 1 filled three named gaps and fixed one guard bug;
the delta must be measured against a second pass before "stabilized" can be claimed.

## Pass 2 — 2026-08-16 — final

**Iters hardened this pass:** all milestone-touched code (final mode, cumulative scope — dimensions 2
and 3, edge cases + error paths, on the surfaces pass 1 did not reach)
**Tiks covered since prior pass:** 0 new tiks (same iter footprint; a second pass over it, per the
stop-condition rule that a delta cannot be computed from one pass)

**Coverage delta on touched files:**
- `dev-stack/migrate-dev.sh`: 0 assertions on the net-new positional-N contract -> 4 static fences +
  3 behavioural exit-code checks
- `stack-snapshot`: unchanged from pass 1 (77.0% / 92.4%) — pass 2's findings were on the shell and
  fence surfaces, not the Go ones

**Tests added:**
- iter-01 -> `dev-stack/tests/test_dev_stack.py` (`TestMigrateDevStaticFence`): 4 static fences —
  positional-N parsed + validated, dev-N offset-port derivation, the shared-clone-set rule, and the
  `::1` hazard (no `localhost` in the atlas DSN)

**What the dimension scan surfaced.** `migrate-dev.sh`'s net-new positional-`N` handling had **zero**
test coverage — the fix whose absence made the entire dev path unmigratable, and whose failure
presented three layers from its cause (no migration -> backend `Exited(1)` on an enforcer panic ->
no taxonomy -> "the catalog is empty"). A `test_migrate_dev_live.py` exists and a
`TestMigrateDevStaticFence` exists; neither mentioned the new contract.

Behavioural verification beyond the static fences: `notanumber`, `-1` and `3.5` each exit **3** with
the usage message.

**Bugs surfaced + fixed inline:**
- `test_isolated_cold_proof_env_overrides` was pinned to the LITERAL `${DEV_PGC:-anthropos-postgresql-1}`.
  M265's derived defaults preserved the property it protects (no N -> the main dev stack) but moved
  the spelling, so the fence went red on a correct change. Worse in the other direction: a rewrite
  that kept the spelling while changing the branch reaching it would have gone **green**. Re-pointed
  at the property (`_DEFAULT_PGC`/`_DEFAULT_PGPORT`/`_DEFAULT_BACKENDC` in the no-N branch).
  Fixable-inline: one test, ~15 lines, a direct corollary of the change that surfaced it.

**Flakes stabilized:** none. The full `dev-stack` suite — **114 tests** — runs clean in 81.9 s.

**Knowledge backfill:** none. Pass 2's finding is about fence design (*pin the property, not the
spelling*), which is already the operating rule in this repo's fence family rather than a new fact
about the platform; it is recorded at the fence itself, where the next editor will read it.

**Stop condition:** stabilized — coverage delta between pass 1 and pass 2 is 0.0% on the Go
surfaces, the dimension scan across all 6 dimensions found no further untested surface in the
milestone-touched footprint, and every suite is green (Go: `stack-snapshot` + `stack-seeding`;
Python: 114 dev-stack + 10 seed-role-guard; the guard family's own verdicts re-checked).
