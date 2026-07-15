# iter-03 (tik) — progress

**Type:** tik, under TOK-01 (Phase B). Lands the three off-box demo-hygiene fixes so Phase C's live battery
grades the SHIPPED code (M219's graded≠shipped lesson).

## What changed

### (1) `FIX-M221-academy-loopback-bind` (F-M220-5)
- **`demo-stack/ant-academy.sh`** — the localhost `next dev` bind is now an **explicit `-H 127.0.0.1`**
  (`bind_args=(-H 127.0.0.1); [ -n STACK_PUBLIC_HOST ] && bind_args=(-H 0.0.0.0)`). The old
  `bind_args=()` passed no `-H`, and next dev's OWN default is `0.0.0.0` → the academy was world-published on
  `*:$PORT` on EVERY localhost demo. The comment's "byte-identical to today / next's own default" framing was
  the bug talking; replaced. The public path is untouched.
- **`stack-injection/exposure_claim_guard.py`** — extended to the two **host-native** listeners (the class it
  was blind to). New `host_native_localhost_binds()` / `find_host_native_exposures()` derive each listener's
  **effective** localhost bind = the launch flag OR **the tool's own default**, read from source:
  `NEXT_DEV_DEFAULT_BIND=0.0.0.0` (named + commented) and `cockpit.py`'s parsed argparse `--host` default
  (`127.0.0.1`). `main()` fails (rc=1) on a non-loopback host-native bind and rc=2 on an unparseable construct
  (a FINDING, per the fence's own doctrine).

### (2) `FIX-M221-reap-native-academy`
- The reap CODE already shipped (M217 `reap_native_ports` covers cockpit+academy by port; M220 S5(i) added the
  `ant-academy.sh` **pre-bind** reap) — but the pre-bind wiring was **UNFENCED**: the launcher harness never
  copied `reap.sh` into its section dir, so `[ -f "$HERE/reap.sh" ]` was always false and the reap block ran in
  **zero** tests. Landed the fence in **`demo-stack/tests/test_ant_academy.py`** (`TestAntAcademyPreBindReap`):
  static wiring (reap sourced + `reap_port "$PORT"` port-scoped identity, BEFORE the launch), behavioural
  reclaim (a stale academy-identity listener on the offset port is reaped → the new academy binds), cross-N
  (a listener on a different offset port survives), + a **mutation** test (excise the reap block → the stale
  academy wins + the launcher reports not-serving).

### (3) `PROBE-M218-backend-api-url-twin` (F-7)
- **`demo-stack/backend_api_url_server_reader_guard.py`** (new) — scans the next-web clone and **fails loud**
  if a **server-side** reader of `NEXT_PUBLIC_BACKEND_API_URL` appears (route handler · app-router server
  component w/o `'use client'` · middleware · `getServerSideProps`/`getStaticProps`/`getInitialProps` ·
  `server-only` import). Enforces M218 D10 instead of restating it. Exit 2 on a zero-reader tree (a FINDING).
- **`demo-stack/tests/test_backend_api_url_twin.py`** (new) — classify unit tests, synthetic-tree scan (each
  RED form + the clean form), the **real-clone gate** (skip-not-pass if absent — ran GREEN: 12 readers, 0
  server-side), a **real-source mutation** (strip `'use client'` from a shipped app-router page → flagged), and
  main() exit codes. Modelled on M218's `test_ssr_origin_chain.py`.
- READMEs updated (`stack-injection/README.md`, `demo-stack/README.md`) for discoverability.

## Fences RED-proven (the release's core discipline — none is theatre)
- **(1)** Reverted the real `ant-academy.sh` bind to the pre-fix `bind_args=()` → `find_host_native_exposures()`
  flagged `academy:0.0.0.0` and **3 committed fences FAILED** (`test_bind_defaults_to_loopback_…`,
  `test_the_real_shipped_ant_academy_binds_loopback`, `test_the_real_tree_has_no_host_native_exposure`);
  restored → green. Plus unit RED-proofs on the pre-fix vs fixed snippet + the "reads the REAL tool default,
  not an assumed loopback" cockpit proof, and a `main()` rc=1/rc=2 pair.
- **(2)** Committed mutation `test_MUTATION_without_the_reap_block_the_stale_academy_wins` — with the reap block
  excised, the stale academy survives and the launcher reports DIED/NEVER-ANSWERED (the exact billion failure).
- **(3)** Committed `MainExitTest.test_server_reader_returns_1` (a `route.ts` reader → exit 1),
  `test_MUTATION_stripping_use_client_from_a_real_page_is_caught` (real source), and a CLI RED demo
  (synthetic route handler → exit 1, loud). GREEN on the real clone.

## Regression / HARD INVARIANT
- **Full sweep GREEN:** `demo-stack/tests` + `stack-injection/tests` = **892 passed, 12 skipped, 0 failed**;
  `stack-core/tests` = **152 passed** (no cross-section impact — no stack-core ref to academy bind / host-native
  exposure). `py_compile` clean; shellcheck (via `test_ant_academy.py`) clean.
- **M220 S3 HARD INVARIANT not disturbed:** `test_public_host_flip.py::TestTheFallbackIsByteIdentical` stays
  green. Fix (1) touches only `ant-academy.sh`'s `bind_args`, not `up-injected.sh`'s derivation trio
  (`BIND_HOST=""`/`HOST=localhost`/`SCHEME=http`) or the tailscale probe — the only change is the academy's
  bind interface **tightening** `0.0.0.0`→`127.0.0.1`, a strict de-exposure the invariant never covered.
- The GUIDE test-count doc-truth test (`TestGuideDocTruth`) counts a curated `test_tooling.py` subset — my
  additions are outside it, so no count drift.

## rext roll
Committed on `main`, tagged **`cue-to-cue-m221-r2`**, pushed to `anthropos-work/rosetta-extensions`.

## Off-box
No live `billion` cycle needed — all three are fence-able off-box (per the operating rules). The live proofs
are Phase C's battery: the academy loopback bind + the reap firing on a real re-up + the backend-api-url
blackhole re-measure. This tik makes the battery grade the shipped code.

## Close — 2026-07-15

**Outcome:** Landed the Phase B off-box demo-hygiene cluster: the academy binds loopback on a localhost demo
(F-M220-5) and the exposure fence now sees the host-native listeners; the `ant-academy.sh` pre-bind reap is
fenced (it shipped M217+M220 but ran in zero tests); and the dormant `NEXT_PUBLIC_BACKEND_API_URL` server-side
blackhole (F-7) is fenced so it can never silently fire. All three RED-proven pre-fix; the full suites are
green. Phase C's live battery will now grade the shipped code.
**Type:** tik
**Status:** closed-fixed
**Gate:** sub-gate MET — the three fences are RED-proven pre-fix (fix 1: 3 committed fences fail on the reverted
bind + the guard flags `academy:0.0.0.0`; fix 2: the committed mutation shows the stale academy wins without
the reap; fix 3: a server-side reader → exit 1, incl. a real-source mutation) and GREEN after
(892+152 passed). The milestone's 8-condition gate is **not** measured by this tik — Phase B lands the code
Phase C's live battery will measure; N/A for the milestone gate this iter.
**Phase 5 grading:** (1) gate-met: n (milestone gate not measured by this tik — Phase B is prep for the Phase C
battery) — (2) triggered-tok: n (tik made measurable progress: 3 fixes + fences landed) — (3) re-scope: n —
(4) user-blocker: n — (5) cap-reached: n (2 tiks this milestone) — (6) protocol-stop: n — **Outcome: continue**
(Phase B done → Phase C, the live `billion` battery, is next).
**Decisions:** D-M221-03a..h (iter-03/decisions.md).
**Side-deliverables:** the exposure fence now enforces a whole new listener class (host-native), closing the
blind spot that let F-M220-5 hide from S0/S1; and the `ant-academy.sh` pre-bind reap gained its first test
coverage (the launcher harness now copies `reap.sh`, so the reap block is actually exercised).
**Routes carried forward:** Phase C (the live cold-reset-to-seed battery on `billion`, NO FLAGS) + its folded
M219 readiness re-proof; and the still-off-box carries NOT in this tik's scope — `FIX-M221-academy-empty-catalog`
(a separate content-pipeline defect → its own tik), `FIX-M221-devstack-…`(discharged), `F-M220-4`
(ant-academy re-runnability on a live public-host demo), `BURNIN-M221-dev-public-host`,
`PROBE-M218-c3-rerun` (needs the box). Per TOK-01.
**Lessons:** an exposure fence that enumerates listeners must enumerate ALL of them — the academy's world-publish
hid for a whole release inside the ONE listener class the fence didn't know (D17: a confident, quietly
incomplete pass). And "the reap code exists" ≠ "the reap is tested": the pre-bind reap shipped across two
milestones and ran in zero tests because the harness silently skipped it (`[ -f reap.sh ]` false) — a fix with
no failing-when-removed test is indistinguishable from a fix that was reverted.

## Next iter
- iter-04: **Phase C — the live cold-reset-to-seed battery on `billion`** (a DEFAULT `/demo-up N`, NO FLAGS,
  driven from a tailnet peer): prove all 8 gate conditions + fold in M219's five readiness gates re-proven at
  final code. First cycle establishes the baseline. `GUARD-M221-host-isolation` (iter-02) + the Phase B fixes
  (iter-03) are the prerequisites now in place.
