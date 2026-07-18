# Hardening Ledger — M226 Opening night

_Final-mode harden pass after the 7-condition gate fired (MET on `billion`, 2 clean cold cycles, +
independently re-verified from this Mac). Cumulative-scope sweep across all milestone-touched code, the last
work before `/developer-kit:close-milestone`. The billion demo stayed UP throughout — harden operated on the
deterministic go/python/tsc surface only; no re-bring-up._

## Pass 1 — 2026-07-17 — final

**Iters hardened this pass:** all milestone-touched code (iter-02 → iter-05; iter-01 is the bootstrap tok — no
code). Cumulative footprint = `be431c3..HEAD` in the rext authoring clone (6 files):
`stack-injection/gen_tailscale_serve.py`, `stack-seeding/blueprint/presets_test.go`,
`stack-seeding/presets/stories.seed.yaml`, `stack-verify/e2e/run-latency.sh`,
`stack-verify/e2e/tests/m224-candidate-heroes.spec.ts`, `stack-verify/e2e/tests/render-hiring-comparison.spec.ts`.

**Tiks covered since prior pass:** all iters in milestone (first + only harden pass — no prior ledger entry).

**Coverage delta on touched files:**
- `stack-seeding/seeders/{users,persona}.go` role funcs (`roleForHero`/`roleForIndex`/`endUserHeroRole`/
  `personaIndexMapForStory`): **100.0% → 100.0% stmts** (already saturated). The gap here was **not**
  statement coverage — it was a **missing end-to-end SEMANTIC invariant**: nothing asserted the NET 5-admin/
  45-candidate outcome on the SHIPPED preset (the gate-C1 count). Now pinned + RED-proven (see below).
- `stack-injection/gen_tailscale_serve.py`: **~94% stmts** from the canonical `TestTailscaleServe` suite
  (unchanged by the fix — the added hiring lines were already in the covered path). The gap here was a
  **stale EXACT-SET assertion**, not an uncovered line (see the bug below).
- `stack-verify/e2e/*.spec.ts`: `tsc --noEmit` **clean (exit 0)** — the M226 remote-capability env plumbing
  (`RENDER_TEST_TIMEOUT_MS`, `CANDIDATE_HOST`/`CANDIDATE_OFFSET`/`CANDIDATE_APP_SCHEME`) type-checks.
- `stack-verify/e2e/run-latency.sh`: **shellcheck-clean (error severity)**; no deterministic bash-unit surface
  exists in rext (no `.bats` harness anywhere) — the recruiter-vantage arg-parse is live-exercised by the C5
  gate. Recorded as live-only, not a coverage gap.

**Tests added:**
- iter-04 (seed count fix) → **NEW** `stack-seeding/seeders/hiring_count_harden_test.go`:
  `TestHiringPreset_NetRoleDistributionIs5Admin45Candidate` — loads the SHIPPED `stories.seed.yaml`, walks all
  50 population slots through the REAL assignment path (`personaIndexMapForStory` → `roleForHero` →
  `roleForIndex`/`endUserHeroRole`) and asserts the gate-C1 NET distribution: **EXACTLY 5 admin + 45 candidate
  + 0 member**, plus the mechanism (band=7, exactly 2 candidate-hero overrides). 1 integration test.
  **RED-proven**: with the shipped ratio temporarily flipped to the old `admin: 0.10` it reports *"NET admins =
  3, want 5 … NET candidates = 47, want 45"* — the exact M226 bug — then reverts clean.
- iter-02 (the tailscale-serve hiring front) → `stack-injection/tests/test_injection.py` `TestTailscaleServe`:
  **NEW** `test_apps_hiring_is_fronted_and_rides_the_ui_tier_lifecycle` — apps/hiring (:3001+off → 13001) is
  fronted by default, dropped under `--no-ui`, and cleared by the reset (the M215 F12 inverse). 1 edge/integration
  test + attribution anchor.

**Bugs surfaced + fixed inline (Fate 1):**
- **3 deterministic M226 regressions in `TestTailscaleServe`** (`test_injection.py`) — the hiring-front change
  (iter-02, `ee1bdf2`) added a **7th** browser-facing port but three exact-count/exact-set assertions still
  pinned **6**: `test_public_host_fronts_all_browser_facing_ports` (`len==6` + fixed 6-port set),
  `test_only_ports_omitted_is_byte_identical_to_before_s7` (`len==6`), and
  `test_main_out_writes_executable_and_reports_port_count` (`"6 port(s)"`). The iter loop measured LIVE on
  billion and never re-ran the deterministic Python suite, so these slipped through green iters. Updated all
  three to the 7-port default (incl. `--https=13001`) with M226 attribution comments. **This is the canonical
  final-harden catch: a cross-iter deterministic regression the per-iter live measurement couldn't see.**
  (Commit with the tests above.)

**Flakes stabilized:** none (all added/changed tests are pure-deterministic — no network, no timing). Flake gate:
**3/3 consecutive clean** for both the Go net-role fence and the Python `TestTailscaleServe` suite (34 tests).

**Knowledge backfill:**
- `corpus/ops/demo/latency-budget.md` — folded in the **R4 finding** (the milestone's declared `delivers`): the
  candidate-comparison drawer's COLD first render is a **warm-up transient** (~2.5 min first sim, warms to
  ~2.4 s), **NOT a gate violation** — C2 gates on data-present-and-renders (with the env-tunable
  `RENDER_TEST_TIMEOUT_MS` budget absorbing the cold transient so it can't false-fail), C5 gates on
  login→ACCESS (not the drawer). Also corroborated the recruiter vantage with the orchestrator's independent
  **p95 1.74 s** re-verify from this Mac.
- `corpus/ops/demo/tailscale-serve.md` — already documents the apps/hiring :3001+off serve front (folded at
  iter-05, incl. the M226 Finding-1 note); verified present, no edit needed.

**Deferred (Fate 3 — routed forward, NOT gate-blocking):**
- **Finding-3 — the pre-bind serve reap** (clear stale `tailscale serve` fronts on the demo's offset ports
  before bind, closing the M215 F12 out-of-band-serve window). The inverse machinery (`reset_commands`/
  `reset_script`) already exists and is now **unit-fenced** (incl. hiring). Wiring a pre-bind `--reset` into
  `up-injected.sh` is a **bring-up-path behavioral change on a live-only surface** (no tailscaled in the build
  env → not deterministically testable here) that would need a **live re-prove on billion** — which the harden
  constraint forbids (the demo is UP; no re-bring-up). **Self-resolves in the default flow** (teardown already
  emits the reset). Fate: land in a follow-up build-iter with a live re-prove, or at the next `prove-on-<VM>`
  milestone. Not gate-blocking.
- **Pre-existing carries (NOT M226 regressions; HEAD-identical):** the 8 demo-stack failures
  (6 `test_cockpit` + `test_purge` + `test_reap`) and the 1 `ptvalidate` M204 assign-WRITE TODO. Untouched
  per scope; noted for `/developer-kit:close-milestone` deferral triage.

**Stop condition:** **stabilized** — the cumulative sweep across all 6 milestone-touched files is complete; the
deterministic surface is fully fenced (Go role invariant pinned + RED-proven; the 3 serve regressions fixed +
an explicit hiring fence; specs tsc-green; bash shellcheck-clean) or is live-only (bash/TS-e2e) with no
deterministic delta to gain. The dimension scan surfaced exactly one bug class (fixed inline) + one missing
invariant (fenced); the confirming re-measurement (3× flake gate + full-suite re-run: 256 passed / 8 skipped,
seeders+blueprint green) showed **zero new findings and zero coverage movement**. Ready for
`/developer-kit:close-milestone`.
