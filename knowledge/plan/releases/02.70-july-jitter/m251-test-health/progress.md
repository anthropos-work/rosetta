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

## M251: Hardening

### Pass 1 — 2026-07-23
**Scope manifest (milestone-touched files — all in the rext repo, all TEST code):**
- `stack-verify/e2e/run-unit.sh` — the unit-spec roster (Section 1). Guard: `stack-verify/tests/test_e2e_collection_integrity.py::UnitSpecsAreExecuted` (already the two-way disk↔roster fence).
- `demo-stack/tests/test_cockpit.py` — 6 re-pointed assertions (overlay ×2 + academy ×4) + 2 class docstrings.
- `demo-stack/tests/test_host_prereqs_m215.py` — 1 re-pointed assertion (`_UI_PORTS` + port-13001).

**Coverage (measure step):** N/A by design — every milestone-touched file is TEST code, not production
code. Line/branch coverage of a test file is not the applicable robustness metric; the right measure for a
re-pointed assertion is **mutation-testing** — does the assertion go RED when the removed behaviour is
re-introduced. Applied that instead (below). Not a skipped-coverage deferral: mutation-verify is a *stronger*
signal than line coverage for this milestone's surface, and it is 4/4 green.

**Mutation-verify (the re-pointed removal-guards actually fire):** temporarily re-introduced each removed
behaviour in the SOURCE (`cockpit.py`, `gen_tailscale_serve.py`), ran the re-pointed test, confirmed RED,
guaranteed-restored (git-clean asserted after). **4/4 guards fire:**
- overlay 30s window re-added (`var STALE_MS = 30000`) → `test_inflight_flag_is_cleared_on_return_no_stale_window` RED ✓
- `removeItem` unguarded (try/catch stripped) → `test_localstorage_access_is_guarded` RED ✓
- per-hero academy `<a class="btn academy" …33077…>` render re-added → all 4 academy removal-guards RED ✓
- hiring `("hiring", 3001)` dropped from the generator's UI tier → `test_public_host_emits_per_port_off_for_all_browser_ports` RED ✓

This validates the "a re-add goes RED, on purpose" claims in decisions.md D-2/D-3/D-4. The re-pointed tests
are self-regression-testing — no additional committed regression test is warranted (a separate test would
only re-assert what these guards already assert). The Section-1 roster is likewise already guarded by the
pre-existing `UnitSpecsAreExecuted` two-way fence.

**Tests added:** 0 (the re-points ARE the guards; mutation-verify confirmed they hold). The Step 2b six-dimension
scan found nothing new worth adding as a committed test — the surface is mechanical test-file re-points with no
new production code, error path, boundary, or perf-critical path introduced.

**Bugs fixed inline:** none (no bug surfaced; all guards behave as specified).

**Flakes stabilized:** none needed. Flake gate: `test_cockpit.py` + `test_host_prereqs_m215.py` **3/3 clean**
(207 passed each, sequential); `test_e2e_collection_integrity` (runs the full 172-test unit sweep) **3/3 clean**.

**Knowledge backfill:** no KB-worthy findings — mutation-verify confirmed the *current* documented behaviour
(overlay removal already in `cockpit-spec.md:271-277`; hiring fronting in `tailscale-serve.md`; academy-link
removal in-code at `cockpit.py:858-862`). Nothing new about the system surfaced.

### Stop condition
Stopped after pass 1: the six-dimension scan found nothing new worth adding (the touched code is test code
whose robustness is proven by mutation-verify, 4/4 green), no coverage delta applies, and the flake gate is
3/3 clean. Further passes would add ceremony, not signal.
