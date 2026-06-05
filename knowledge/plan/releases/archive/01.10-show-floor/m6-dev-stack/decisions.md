# M6 — Decisions

## M6-D1 — `gen_override.py` extracted to a shared `stack-core/` section (build, 2026-06-04)
Settles the M5-routed open question. The port-offset / `!override` multi-instance engine `gen_override.py` was
the shared primitive both demo + dev need, so it moved out of `demo-stack/lib/` into a new **`stack-core/`**
section (consumed by both). Extract-on-second-consumer: it stayed in demo-stack through M3–M5 (one consumer),
and moved the moment dev-stack became the second. `clone_repos.py` (clone-at-release-tag) **stayed in demo-stack**
— it's demo-specific (per-stack clones); dev uses the platform compose directly (M6-D3). The `rosetta-demo` CLI
was repointed (`../stack-core/gen_override.py`); `TestGenOverride` moved to `stack-core/tests/`.

## M6-D2 — dev-stack scoped to the proven value, not speculative multi-dev (build, 2026-06-04)
The M4 research flagged "N concurrent dev boxes" as **unproven demand**, and the base `rosetta-demo up` path
already does isolated, non-injected multi-instance. So M6 delivered the *proven* value — the shared-engine
extraction + a focused dev tool + **optional** injection — and deliberately did **not** build a full
multi-concurrent-dev system or a separate dev-injection machinery. `dev-stack --inject` **reuses
`../stack-injection`** (the M5 layer); the full injected dev bring-up (per-service disarmed-colony rebuild) is
**identical to demo's** and reuses `demo-stack/up-injected.sh`'s flow with a `dev-N` project — no duplication.
Default is **dev-OFF** (real Clerk from `platform/.env`).

## M6-D3 — dev uses the platform compose directly; no per-stack clones by default (build, 2026-06-04)
`demo-stack` clones the service repos per-demo (M3-D1, for injection + per-demo divergence). `dev-stack` does
**not** — a dev stack brings up the platform compose directly under `-p dev-N` (real Clerk, no rebuild needed),
reusing the engineer's existing dev clones / images. (If injected dev needs per-stack rebuilds, it reuses the
demo flow per M6-D2.) The dev clone strategy beyond this (shared vs per-stack) is left to proven need.

## Adversarial review
- **Did the stack-core extraction break demo?** The repoint + test-split were verified (demo-stack 5 + stack-core
  4 green; the CLI's `gen` path now calls `../stack-core/gen_override.py`). Same depth-fragility lesson as
  M4-D4/M5 — caught by the suites.
- **Can `dev-stack` touch the main dev stack?** No — `-p dev-N` scoped + `guard_n` hard-refuses any N resolving to
  `platform/.env`'s `COMPOSE_PROJECT_NAME` (the `anthropos` project). Pinned by `test_dev_stack.py`.

## Open (routed to M7/M8)
- Whether a `/dev-up`-style skill is wanted in rosetta (mirroring `/demo-*`) — M8 discoverability, only if dev-stack
  proves used. The full injected-dev path (if demand appears) — a future increment, not v1.1.
