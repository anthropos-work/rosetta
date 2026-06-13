# M22 — Retro

## Summary
M22 turned M21's print-only per-stack-Directus recipe into an **executed**, idempotent, prod-safe bring-up
step: a per-stack Directus now boots as a **compose service** (offset port, app-network, torn down with the
stack), provisioned `bootstrap → apply-structure → replay → boot`, verified, demo default-on / dev opt-in
(`--local-content`). The load-bearing design move was making the `EnvContract.Validate` firewall a
**load-bearing executed gate** (hard-abort before any write if the env resolves to prod), with a runtime
no-prod-read backstop in `autoverify` — the M17-for-TRUNCATE discipline applied to the executed provision.
Non-fatal throughout: any provision failure degrades to the prod-read path with an honest `⚠` status line.
6 sections, all Fate-1. The maintainability constraint held: the only steps the recipe executes itself are the
one-shot bootstrap + the post-replay restart compose can't express — teardown / registry / verify cover the
container for free.

## Incidents This Cycle
- **P3 (process, not code) — attempt-1 build crash + resume-in-place.** Attempt 1 crashed mid-§4 with a
  network error (`FailedToOpenSocket`), an environment blip, not a code fault. Recovered by a resume-in-place:
  §1–3 were already committed in attempt 1; attempt 2 reconciled the §1–3 checkboxes to the committed ext
  code, committed the scaffold, and finished §4–6. No work lost, no double-apply. Mitigation already in place
  (the per-section commit discipline made the resume clean).
- **0 code incidents / 0 regressions / 0 flakes.** The 4-pass hardening sweep surfaced 0 production bugs; the
  close adversarial pass (7 scenarios) found 0 new findings (all already test-pinned); the flake gate ran 5/5
  clean.

## What Went Well
- **The compose-service-not-`docker run` call paid off.** Because the Directus is a compose service, M22 wrote
  **zero bespoke lifecycle code** — `demo-down`/`dev-down` teardown, the unified port registry, and
  `stack-verify`'s naming convention all cover it for free. The override generators (`gen_injected_override.py`
  demo + `gen_override.py` dev) just appended a service block in the proven pattern.
- **The firewall-as-executed-gate was unusually well-tested.** Every adversarial scenario the close pass could
  construct (write-before-gate, `set -u`, prod-write-on-degrade, false-safe gate, half-bootstrap, serve-nothing,
  N=0) was already pinned by a named test at build/harden time — the milestone shipped with its safety surface
  exhaustively exercised.
- **The half-bootstrap sentinel guard.** Probing the `directus_collections` *sentinel* (not a blanket
  `directus_*` count) makes a crash-then-rerun re-bootstrap to converge instead of skipping onto a broken
  schema — a genuinely subtle correctness nuance caught + tested during build, not at close.

## What Didn't
- **The CLAUDE.md doc-list drifted (DOC-2).** `directus-local.md` (net-new at M21) is indexed in
  `corpus/ops/README.md` but not in CLAUDE.md's "Demo Environments" prose list — a discoverability gap. This is
  the recurring index-row-miss class; it's correctly routed to M24's CLAUDE.md sweep (Fate-2), but it's a
  reminder that the per-unit handbook contract should check *every* index, not just the directory README.
- **The ops-README index row carried a stale future-claim (DOC-1).** The row still said the lifecycle half
  "lands in M22/M23" after M22 landed it — caught + fixed Fate-1 at close, but it's a status claim that should
  have been updated in the §6 docs commit.

## Carried Forward
- **DOC-2** (CLAUDE.md doc-list adds `directus-local.md`) → **M24** (the corpus-wide CLAUDE.md sweep — already
  in M24's `In:`).
- **3 nice-to-have refactors** (shared `DIRECTUS_BASE_PORT` const across the two emission generators; a shared
  verify-svc-list helper; a `readiness.sh` docker-availability comment) → **M24** hygiene strand (cross-module
  tidy).
- **DEF-M21-02** (automated serve-live-integration harness) → **M25** field-bake (M22 booting a live per-stack
  Directus made this nearly free, as the M21 audit predicted).
- Inherited unchanged: **DEF-M21-03** (`directus_files` ref capture) + **DEF-M21-04** (referential closure) →
  **M23**; **DEF-M21-01** (replayCmd conn-seam) → tracked follow-up; **DEF-M10-01** (S3 blob bytes) → backlog.

## Metrics Delta
(from `metrics.json`)
- **Go test funcs:** 795 → **795** (+0 — M22 touches no Go).
- **Python collected:** 360 → **418** (+58, +8 env-gated skip). Per-module: stack-core 54→61, dev-stack 50→73,
  demo-stack 84→87, stack-injection 85→93, stack-verify 87→104.
- **Coverage:** `gen_override.py` 62% → **85%**; `gen_injected_override.py` 99%; shell engines subprocess-tested.
- **Flake:** **0** (5/5 sequential, dev-stack + stack-verify). **Harden bugs surfaced:** 0. **Adversarial-new:** 0.
- **Close findings:** 8 (0 scope · 0 must-fix · 3 should-fix · 2 docs · 0 tests · 0 blend). **Deferral audit:** GREEN.
