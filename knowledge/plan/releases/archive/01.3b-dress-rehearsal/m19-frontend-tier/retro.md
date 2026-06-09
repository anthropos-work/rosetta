# M19 — Retro

_The demo-up frontend tier — the 4th + largest milestone of v1.3b "dress rehearsal", closed 2026-06-09._

## Summary

M19 completed the demo family: `/demo-up` now brings up the full UI tier — **next-web-app + studio-desk**
per-demo (a **cached** Docker image from the **unmodified** platform Dockerfile, offset ports, minted-pk +
offset-URL baked) + **ant-academy** natively (Clerk-free via `BENCHMARK_VISUAL_BYPASS`) — so a demo is actually
demoable on a 16 GB Mac, with a non-fatal 12 GB VM pre-flight, a `--no-ui` escape, and the frontend ports in
M18's verify net. The defining constraint, the **zero-platform-edit invariant**, held end-to-end: the minted pk
rides a gitignored `apps/web/.env.local`, the `.dockerignore` is tooling-owned + transient + non-clobber +
`RETURN`-trap-removed (byte-clean even on a failed build), the platform repos are a build *context* only with
**unmodified** Dockerfiles. The close was the cleanest possible shape — 6 findings, all decision-triage, with
0 scope / 0 code / 0 docs / 0 tests.

## Incidents This Cycle

- **None at close.** No P2 flakes (flake gate 5/5 deterministic), no regressions (full extensions suite 338/338,
  Go 736 unchanged). The harden pass fixed 0 production bugs — the load-bearing invariants all held under the
  new (harder) tests; the one residual observation (a stale academy pidfile on a failed launch) was judged
  **harmless by design** (`--stop` + re-launch both gate on `kill -0`), so no fix was warranted.
- **Earlier, during build (recorded in decisions.md, resolved before close):** the §4 native-academy launch
  exposed the classic background-job **pipe-hang** — a daemon that keeps the caller's stdout/stderr pipe open
  blocks any `capture_output` caller until it exits. Fixed by detaching all three stdio fds on the launch
  subshell (#M19-D9), with a regression test (`test_daemon_that_dies_immediately_falls_back_and_exits_zero`).

## What Went Well

- **The verified plan paid off.** `.agentspace/demo-up-frontend-plan.md` (the pre-build real-Docker
  investigation) meant the `In:` list was writable with confidence — no scope surprises in the largest v1.3b
  milestone. The ARG contracts (next-web's no-pk-ARG vs studio-desk's pk-ARG) were verified against the real
  Dockerfiles at design time, so the pk-path split (#M19-D3) was a known quantity, not a build-time discovery.
- **The zero-platform-edit invariant was pinned, not asserted.** `TestZeroPlatformRepoEdit` stands up a *real
  git repo* as the build context and asserts `git status` stays empty after the (stubbed) build — including the
  failed/aborted path via the `RETURN`-scoped trap — plus a `git check-ignore` fence on the pk overlay.
  Mutation-verified: dropping the trap cleanup or routing the pk to a tracked path both *fail* the guard.
- **The close was friction-free.** Everything that was going to be the deliverable was already done + hardened;
  the only close work was the standard decision-triage ref-tagging + the adversarial-scenario records.

## What Didn't

- **Live browser verification stayed an operator path, not a milestone gate.** Docker is unavailable on this box,
  so the live ~3.7 GB build + browser-login couldn't be exercised in CI; the milestone gate is the tooling-level
  proof (the override emits both frontends with minted-pk-baked images; the build assembles the offset URLs + pk;
  the verify net probes them). This is the honest M19 resource calibration — the same shape as M18's live-docker
  surfaces — but it means the first *real* `/demo-up N` with the UI tier is still the true end-to-end test.

## Carried Forward

- **Nothing M19-internal.** All 8 deliverables landed Fate 1; 0 routed / dropped / escape-hatch-deferred.
- **True zero-rebuild (one image, port/pk switched at runtime)** stays a documented **OUT** boundary — it edits
  platform repos (forbidden) → an optional **user-owned upstream PR** (the v1.4 deploy-CI precedent), not a
  deferral. The honest residual (one ~3-min cached build per new `demo-N`) is the accepted tooling-only cost.
- **DEF-M10-01** (cloud snapshot store + S3 media blob bytes → v1.4, signed) — inherited, untouched by M19
  (which touched the frontend bring-up surface, not the snapshot-store/S3 area), not aged out.

## Metrics Delta (from metrics.json)

- **Go test funcs:** 736 → **736** (+0 — no Go touched).
- **Python collected:** 273 → **338** (+65 — the UI-tier suites: per-demo build + zero-platform-edit guard,
  native-academy launcher/fallback, frontend-emitting override, frontend-port registration).
- **Coverage:** `gen_injected_override.py` (the one pure-Python touched file) **98%** (saturated — only the
  `__main__` guard miss).
- **Flake:** **0** (5/5 clean, sequential, on the 193-test milestone-touched set).
- **Findings at close:** 6, all decision-triage (0 scope / 0 code / 0 docs / 0 tests).
- **Extensions:** tag `dress-rehearsal-m19` reconciled `32b1ae8 → 4f96ddd`, force-pushed to origin,
  `stack-demo` re-consumed (3-way agreement).
