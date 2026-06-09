# M18 — Retro

_The verification safety net. Closed 2026-06-09 → `release/01.3b-dress-rehearsal`._

## Summary
M18 made a bring-up self-verifying: `stack-verify` now targets an offset stack + scopes to the services
actually brought up, and runs (default-on, NON-FATAL) at both bring-up tails — cheap-win `/api/health` +
`sentinel.casbin_rules > 0` asserts (the ISSUE-7 silent-403 catcher) then the full offset/project/scope-aware
probe set. So "UP" now carries a real promise. The defining constraint was the **load-bearing non-fatal
invariant** — the auto-verify must never block a genuinely-good bring-up and never systematically false-`down`
a healthy offset or reduced-profile stack — which the close verified on three axes (always-exit-0 wrapper +
`|| true` call sites; offset derived-from-known and cross-checked against the registry's recorded ports via a
base-port band, not a drift-prone formula; both liveness AND readiness honour the scope filter).

## Incidents This Cycle
- **P2 (harden Pass 1) — the readiness phase ignored `STACK_SERVICES`.** The build phase scoped only the
  liveness phase (`service_rows`); the 6 deep readiness probes ran unconditionally, so a reduced bring-up
  (`--services "postgresql redis"`) false-`down`ed graphql/gotenberg/sentinel/storage — the exact ISSUE-12b
  wall-of-false-downs M18 exists to prevent, contradicting decision M18-D2. Fixed inline in harden (commit
  `2f412a3`): `run_readiness` threads the backing service name + gates on `target_service_selected`. +
  `TestReadinessScopeFilter`.
- **P2 (close Phase 2c, FINDING-A1) — a non-numeric offset crashed the arithmetic under `set -u`.** An
  invalid `--offset`/`STACK_OFFSET` (e.g. an operator/`/test-platform` typo) flowed verbatim from
  `target_resolve_offset` into three `$(( base + offset ))` sites → "unbound variable", **silently skipping**
  the cheap-win asserts. The non-fatal invariant still held (the wrapper exited 0), but the asserts vanished
  with a confusing message. Fixed at the single resolution boundary (M18-D7): `target_resolve_offset`
  validates `^[0-9]+$`, else warns non-fatally + derives from the project's N. Regression pinned at unit +
  integration level.
- No flakes, no regressions in the existing Go/Python suites.

## What Went Well
- **The non-fatal invariant was designed in, not bolted on.** `autoverify.sh` is structurally non-fatal
  (no `set -e`, every probe wrapped, always `exit 0`) with belt-and-suspenders `|| true` at both call sites.
  Even the two bugs found (readiness filter, A1) never broke the invariant — they were correctness/clarity
  gaps, not abort paths.
- **The base-port-band cross-check (M18-D5)** sidestepped the roadrunner high-base decade trap the build
  PR-review (A1) caught — a single resolution helper (`target.sh`) keeps the SERVICES table one source of
  truth and applies the offset centrally.
- **A net-new test suite from zero** (stack-verify had only Playwright e2e before): 82 stdlib-unittest tests
  driving the bash via subprocess, every milestone-touched seam pinned.

## What Didn't
- **Per-phase scoping was asserted in the doc before it was true in the code.** The build doc implied both
  phases scoped; only liveness did. Harden caught it, but it's a reminder that "the doc says X" is not
  evidence "the code does X" — the harden/close test pins are what make the claim true.
- **The offset wasn't validated at its single boundary in the build phase**, so a non-numeric value could
  reach three arithmetic sites. Fixing once at `target_resolve_offset` (not at each site) was the right shape;
  it just should have been there from the start given the helper already centralized resolution.

## Carried Forward
- **Frontend-port verification → M19** (Fate-2, already owned — M19's overview line: "Register the frontend
  ports so M18's verify net covers them"). A clean scope boundary, not a deferral.
- **DEF-M10-01** (S3 blob bytes + cloud store → v1.4, signed) — inherited, orthogonal to M18 (which touched
  the verify surface, NOT the snapshot-store/S3 area), not aged. Unchanged.

## Metrics Delta (from metrics.json)
- Go test funcs: **736** (unchanged — M18 touched no Go).
- Python test funcs: 191 → **273** collected (+82 — the net-new `stack-verify/tests/test_verify.py`; 32 build
  → 79 harden → 82 close).
- Flake: **0** (5/5 sequential on the touched suite, deterministic). shellcheck (9 scripts) + py_compile CLEAN.
- Findings: 3 (1 code-quality must-fix A1, 1 test regression, 1 decision-triage) — all Fate-1, landed.
- Deferral re-audit: **GREEN** (0 new, 0 repeat, 0 aged-out).
- Extensions tag: **`dress-rehearsal-m18`** @ `777723a` (reconciled from `594b9cf`).
