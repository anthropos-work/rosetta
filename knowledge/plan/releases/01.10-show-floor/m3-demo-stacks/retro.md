# M3 — Retro

**Summary:** Built the disposable demo-stack layer for v1.1 — bring up `demo-N` as an isolated Anthropos
stack alongside the dev stack, Clerkenstein-wired, on offset ports, killable cleanly, **with zero
read-only-platform change**. Tooling lives in a new gitignored `anthropos-demo/demo-stacks/` repo (the
clerkenstein pattern); the `/demo-*` skills + `corpus/ops/demo_stacks.md` are rosetta-tracked. All 5
sections built + hardened (12 unit tests). **Acceptance met (M3-D5): demo-1 ran isolated alongside the dev
stack — up→status→down, dev untouched.**

## Incidents / findings this cycle
- **Docker VM disk full → redis "No space left on device"** during the live proof. Reclaimed 21 GB of
  build cache (safe — `docker builder prune`, no images/containers/volumes touched); redis then came up.
  Environment limit, not a tooling defect (postgres came up fine).
- **The `!override` collision bug (caught in review before any `up`).** A plain compose override *appends*
  to `ports`/`volumes`, so demo-1 would have re-bound the base port (5432) and collided with the dev
  stack. Fixed by emitting the override with Compose's `!override` tag (replaces the sequence). The merged
  `docker compose config` made it visible. This is the single most important detail in the engine.
- **shellcheck SC2193** — a dead literal guard (`"demo-$n" != "anthropos"`) → replaced with a real compare
  against the platform `.env`'s configured `COMPOSE_PROJECT_NAME` (so the dev-stack guard tracks reality).
- **py3.9 system Python** — `str | None` annotations broke the tests; `from __future__ import annotations`.

## What went well
- **Driving the Docker work directly (not an autonomous sub-agent)** kept the user's running dev stack
  provably safe — every op `-p demo-N`-scoped, `down` hard-refuses the dev project. demo-1's whole
  lifecycle ran with the dev stack (12 containers, postgres healthy 9h) untouched.
- **Grounding the engine in the real compose first** (24 hard-coded ports, one project name, one bind-mount)
  made the additive fix obvious and the milestone a confident `section`.
- **The clone-at-release-tag resolver** handled the org's mixed tag convention (bare `0.1.0` + `v1.282.0`)
  cleanly on the first pass.

## What didn't / constraints
- **16 GB host + Docker's ~8 GB VM + the dev stack already up** → can't run two full stacks, or even one
  *full* (12-service) stack comfortably. Accepted (M3-D5): the live proof is one minimal demo alongside
  the dev stack; the full-scale + end-to-end Clerkenstein browser-login proofs are documented + wired but
  verify on a bigger Docker VM.
- **Per-demo full clones (M3-D1)** are disk-heavy; only a 2-repo clone was exercised live (disk-tight box).

## Carried forward → a bigger box / M4–M5
- Full 12-service single stack + ≥2 concurrent stacks + end-to-end Clerkenstein browser login (the wiring
  is built; needs a bigger Docker VM to verify).
- `max-N` concrete bound + a `/demo-up` memory/disk budget-check (documented as a knob; enforce in M4/M5).
- The express-gate CI carry-forward (from v1.0) still routed to M5.

## Metrics
See [metrics.json](metrics.json). Tooling: ~3 Python modules + 1 bash CLI + 3 skills + 1 ops guide; 12 unit
tests green; shellcheck/py-compile clean. demo-stacks repo: 5 commits.
