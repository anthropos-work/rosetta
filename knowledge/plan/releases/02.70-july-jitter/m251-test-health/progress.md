# M251 — Progress

## Sections
- [x] **run-unit roster fix** — added `content-denominator.unit.spec.ts` + `run-discrete.unit.spec.ts` to the `UNIT_SPECS` roster in `stack-verify/e2e/run-unit.sh` (7 → 9). `run-unit.sh` exits 0 (172 unit tests green); `test_e2e_collection_integrity` 8/8 GREEN incl. `UnitSpecsAreExecuted`. rext commit `cf53426`.
- [x] **Mechanical python re-points** — re-pointed the 6 `test_cockpit` academy/overlay assertions + the `test_public_host` (`test_host_prereqs_m215::TestF12ServeResetGenerator::test_public_host_emits_per_port_off_for_all_browser_ports`) port-13001 assertion at the deliberately-changed behaviour (overlay 30s-window removal @ M218; per-hero academy link removal 2026-07-15; hiring 3001→13001 UI-tier fronting @ M226). `test_cockpit.py` + `test_host_prereqs_m215.py` = **207 passed, 0 failed**. rext commit `e9e29a1`.

## Completeness Ledger

### Deferred
- **Optional `corpus/ops/verification.md` demo-stack-suite + run-unit-roster index anchor** (overview.md "Delivers →") → **M247** (Fate 3-adjacent: M247 owns the corpus reground + `verification.md`; authoring here would collide with its concurrent lane). Not a blind area — the code it would index exists + is exercised. (decisions.md)
- **The out-of-scope live/env/docker-gated demo-stack failures** (`test_purge` + academy launcher/reap ×5 + `test_host_isolation` mutant-term-trap + `test_ant_academy_clerk_wiring` overlay = **8** on a stackless box) → **M254** (Fate 2 — M254 gate parts (g)+(h) own the live-box test-health + live-browser proofs). Flagged: M254's overview names "~2" but the real live-gated set is 8 — the closer should expect the larger set. (decisions.md)

### Dropped
- (none)
