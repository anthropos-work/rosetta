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
