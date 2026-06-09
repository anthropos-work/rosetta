# M17 — Progress

_Section checklist. Closure = all boxes land + `/developer-kit:close-milestone` GREEN._

## Deliverables
- [x] **`set -e` first-run-race audit** across `up-injected.sh` / `rosetta-demo` / `dev-stack` / `dev-setdress.sh`; risky spots fixed (3 ISSUE-7 kin sites: GH_PAT loud-guard, DEV_PROJECT ×2 `|| true` fallback). 4th site (schema-create) surfaced by the live harness + fixed.
- [x] **"wait-for-sentinel-ready"** defense-in-depth before migrate (`wait_pg` via `pg_isready`/`SELECT 1` + `wait_sentinel_running`, both bounded + non-fatal).
- [x] **`stacksnap replay` re-run guard** — per-stack-isolated TRUNCATE-then-reload (child-first) before COPY; 2nd replay REPLACES, no abort/double.
- [x] **`stackseed` casbin g2 + det-UUID `ON CONFLICT`** — idempotent COPY-to-temp merge (all 7 seeders) + casbin `WHERE NOT EXISTS`; 2nd seed inserts 0.
- [x] **`--reset` truncate list fixed** to the full FK-ordered fleet (+ targeted casbin g2 DELETE, not TRUNCATE — preserves global policy).
- [x] **Tested idempotency contract** per component (replay e2e re-run gate, seed re-seed-inserts-nothing, casbin guard, live docker migrate-race harness).

## Verification
- [x] Go `-race -count=1` + gofmt + `go vet` clean on `stack-seeding` + `stack-snapshot`.
- [x] A regression test per guard, mutation-pinned (removing the guard fails the test — verified for replay clear, casbin, idempotent-copy).
- [x] TRUNCATE target-class test proves it can only hit a per-stack-isolated offset target (`truncateForReplaySQL` shape-pin + `DSNForOffset` offset chain).
- [x] flake 0 (5× shuffle, both Go modules).

## Notes
- Net-new `corpus/ops/idempotency.md` delivered + wired into demo/README, root CLAUDE.md, snapshot-spec, seeding-spec, safety.md (bidirectional).
- M15 safety-doc drift guards stay GREEN after the safety.md "See also" addition.
- Python: demo-stack 27 (test_tooling, +6 over M16's 21) + 3 live docker (test_migrate_race_live, Docker-gated).
- The live harness surfaced + fixed a 4th latent ISSUE-7-class site (schema-create `docker exec` under set -e).

## M17: Hardening

### Pass 1 — 2026-06-09
**Scope manifest (milestone-touched, from `git diff e6161b0..dcef026`):** rosetta = docs-only (idempotency.md + 5 wired parents + tracking) — no testable code. Testable M17 logic lives in `.agentspace/rosetta-extensions` across: stack-seeding/{pg,seeders,cmd/stackseed,seeder} (Go), stack-snapshot/{replay,cmd/stacksnap} (Go), demo-stack (Python tooling + live docker harness + 4 bring-up shell scripts). The build phase wrote substantial co-located tests; the pure SQL builders were already 100% covered.

**Coverage delta (milestone-touched files):**
- stack-seeding/pg: 52.0% -> 55.3% statements (+3.3 — the `CopyRowsIdempotent` validation/conflict-check branch now exercised hermetically). Remaining uncovered = DB-touching (`Connect`/`Exec`/`CopyIn`/live merge), proven by the live docker harness.
- stack-snapshot/replay 100% (unchanged), cmd/stacksnap 77.0% (unchanged — remainder is live-DB adapter wiring).
- seeders 95.2% (unchanged — at the hermetic ceiling).

**Tests added:** 7 (all mutation-pins; finder-not-goal — they deepen behavior on already-covered lines)
- pg/pg_test.go: 2 error-path (empty-conflictCol rejected before pool; validation-precedence) + 3 edge (single-column merge, conflictCol-not-in-cols, createTempLikeSQL) + 1 fuzz (seed-side injection sweep over adversarial identifiers — the analogue of the replay TRUNCATE sweep).
- seeders/seeders_test.go: 2 correctness (casbin g2 dedup predicate must compare ALL 7 columns p_type+v0..v5; exact-repeat-suppressed vs fresh-stack-inserts).
- cmd/stacksnap/adapters_test.go: 1 edge (truncateForReplaySQL degenerate/empty identifiers stay a single-relation TRUNCATE, never a destructive class).

**Bugs fixed inline:** none — the build-phase logic held under every new probe.

**Flakes stabilized:** none observed.

**Knowledge backfill:** no KB-worthy findings — every invariant pinned (per-stack TRUNCATE injection-safety, idempotent-COPY `ON CONFLICT (id)`, casbin `WHERE NOT EXISTS`, full `--reset` fleet) was ALREADY documented in `corpus/ops/idempotency.md` + decisions M17-D4/D6/D7. The harden deepened test ENFORCEMENT of existing claims, surfacing no new behavioral truth.

### Pass 2 — 2026-06-09
**Coverage delta:** cmd/stackseed 47.0% (unchanged — the new pins are over the package-level `resetTables` const, zero executable statements; the rest is live-DB code covered by the docker harness).

**Tests added:** 2 (behavioral pins on the destructive `--reset` surface)
- cmd/stackseed/main_test.go: every `resetTables` entry must be exactly `schema.table` (an unqualified entry + raw `TRUNCATE … CASCADE` could hit the wrong schema via search_path) + no-duplicates.

**Bugs fixed inline:** none. **Flakes:** none.

**Knowledge backfill:** none — the `--reset clears the full fleet` claim is already in idempotency.md §`--reset`.

### Stop condition
Stopped after Pass 2 (well under the 5-pass cap). The six-dimension scan is exhausted for what's hermetically testable — error paths, edge cases, injection-fuzz, and the casbin/reset correctness seams are now mutation-pinned; the remaining uncovered statements are all live-DB code already proven by the Docker-gated `test_migrate_race_live.py` harness + the recordingConn idempotency model. Coverage deltas are negligible going forward and a Pass 3 would only yield shallow tests. Flake gate: 3/3 clean sequential runs, both Go modules, `-race`. Extensions commits: `ea32daf` (Pass 1), `0d36251` (Pass 2).
