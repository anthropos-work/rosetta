---
milestone_shape: iterative
milestone: M257
title: "first-light build"
status: planned
release: v2.8 "fast build"
exit_gate: "A cold-images `demo-down --purge` + `demo-up` reaches `autoverify green:true / 0 warnings` in p50 <= 360 s across 3 consecutive cycles on billion (baseline: measured 672 s — a 46% cut), with the M255 headroom reserve contract never breached (sampled, not asserted), 0 platform-repo edits, and all 7 demopatch guards (G1-G7) still passing. Stretch: <= 300 s."
iteration_protocol_ref: corpus/ops/demo/build-budget.md
re_scope_trigger: "If after L1 + L2 + L3 the p50 is still > 480 s, the remaining cost is structural (host I/O or the containerd snapshotter) — escalate rather than grind."
depends_on: [M256]
parallel_with: []
complexity: very-large
created: 2026-07-27
last_updated: 2026-07-27
---

# M257 — first-light build  (`iterative`)

**Status:** `planned` · **Shape:** `iterative` · **Complexity:** very-large · **Release:** v2.8 "fast build"
**Depends on:** M256 (sharpen the detector before changing what it detects)

## Goal

Collapse the cold demo bring-up so going live is a coffee, not a lunch — **spending the machine deliberately,
never exhausting it**, and without weakening a single safety guard.

## Baseline (measured, n=1, `billion` 8 vCPU / 7.3 GiB / x86_64, cold-images + warm layer cache)

| | |
|---|---|
| Total cycle | **672.4 s (11 m 12 s)** |
| Bring-up | 650.7 s (96.8 %) |
| **UI-tier image builds (3)** | **446.4 s — 66.4 %** |
| **Image export/unpack alone** | **288.4 s — 42.9 %** |
| peak load1 | **4.90 of 8 cores** · avg 2.26 · peak RAM 74 % |

**This is not a CPU problem.** It is serialised I/O (writing 9.4 GB of Next.js image to disk) plus a
**deliberate serialisation**: `build_frontends()` has exactly one conditional (`NO_UI`); the RAM pre-flight it
cites is **cosmetic** (`preflight_vm_ram()` declares its vars `local`, assigns no global, returns no verdict);
and the Go builds it was meant to avoid overlapping **finish 1.1 s before the UI tier starts**. Meanwhile the
UI tier runs at `:1877` and `compose up` at `:1924`, so postgres boot, **4 atlas migrations**, snapshot replay
and the seed **idle for ~7.5 minutes**.

## Levers, ranked by measured seconds recoverable

| | Lever | Est. saving | Shape |
|---|---|---|---|
| **L1** | Multi-stage the two Next images — ship `.next/standalone` + static instead of the full build tree with dev deps. 4.77 GB → a few hundred MB collapses the 141.9 s + 136.7 s export **and** the 85.7 s unconditional unpack leg (L9) | **~200–250 s** | demopatch / rext-owned Dockerfile. **M255 spike (a)** may make it a *selection* change |
| **L2** | Build `next-web` ∥ `hiring` **and** reorder the UI tier to overlap `compose up` | **~200 s** | rext `up-injected.sh` + the M255 safe-parallelism contract |
| **L3** | Manifests-first `COPY` so the `pnpm install` layer survives a source-only change (every demopatch is one). The layer has **never once been reused** — 16 entries × 4.029 GB = **61 % of the whole build cache**, every one `Usage count: 1` | **~55 s** + an ~8 GB/cycle leak | demopatch / rext Dockerfile. Must copy root `package.json` + `pnpm-lock.yaml` + `pnpm-workspace.yaml` **plus all 16 workspace `package.json`s** — a naive `COPY package*.json ./` breaks `--frozen-lockfile` |
| **L4** | Drop `--concurrency=1` from `pnpm turbo build` on an 8-core host | ~20–35 s | build-arg or demopatch |
| **L5** | Speed the taxonomy replay (78.0 s / 330,261 rows + 2 pgvector reindexes): index-after-COPY in one pass, `UNLOGGED`-then-`SET LOGGED`, or a pre-built PG data dir. **Also the main `/dev-up` win** | ~30–50 s | rext `stack-snapshot` |
| **L7** | Multi-stage `studio-desk` — 1.71 GB shipping a full dev toolchain (32,568 JS/CSS files, 266 MB) to serve a Vite bundle | ~8 s | demopatch / rext Dockerfile |
| **L8** | Cache the Directus bootstrap + restart — 15.6 s of pure container-boot latency, the most compressible slice of set-dress | ~15 s | rext |
| **L10** | Serial fat: ~12 serial `git fetch`es · 23 serial `demopatch revert` shells · Go tooling compiled 4–5×/bring-up (`stacksecrets` into a throwaway `mktemp -d`, `stackseed` **twice**) · 4 independent `atlas migrate apply` targets run serially · the entire tailscale-serve plan re-emitted to add one port | ~20–50 s | rext |
| **L6** | Prune BuildKit on a schedule. **Not a time win** — the documented ENOSPC failure mode surfaces as a cryptic *"redis exited (1)"* (M239-F1), and with L3's leak the real runway is **~4–5 cycles**, not 15–20 | 0 s (risk) | ops / rext teardown |

> **L1 + L2 + L3 plausibly take the cycle from ~11 m to ~4–5 m** — they attack the same 446 s block from three
> angles (smaller images to export, exported concurrently, with dependency layers that actually survive).

## Also relevant (operational, not a lever)

**`--purge` defeats every image cache, including 5 hidden ones.** `rosetta-demo:336-341` removes the three UI
images **and** 5 that `compose up` then rebuilds inline (postgresql, graphql, sentinel, storage, roadrunner) —
15.3 s warm here, 120–300 s if truly cold — with **no per-service log file to attribute them to**. So the three
cache-reuse checks (`:562`, `:849`, `:1077`) can **never** hit on a purge cycle. Plain `rosetta-demo down N`
(no `--purge`) keeps the images and makes a re-up cost seconds — the fast-cycle option whenever a wiped DB is
not required. **Document it; do not make it the default** (a wiped DB is usually the point).

## Dev path

`/dev-up` shares **L5 / L6 / L10** (set-dress + tooling, not the UI tier). **Measured and reported at each
iter; not separately gated** — the demo path is where 96.8 % of the wall-clock is.

## Shape (why iterative)

L1's cost depends entirely on M255 spike (a) · L2's real win depends on spike (d) · L3's value depends on
spike (b). The path is measurement-driven by construction: measure → attribute → one lever → re-measure at
n ≥ 3.

## Hard constraints

- **Zero platform-repo edits.** L1/L3/L4/L7 all touch Dockerfiles in canonical repos → each lands as a
  sha-pinned `demopatch` or an **rext-owned Dockerfile** in the shape `hiring.Dockerfile` already sanctions.
- **All 7 demopatch guards (G1–G7) still pass** after every lever.
- **Per-stack image isolation** — cache **layers**, never **images**. A reused image would carry another
  stack's baked publishable key and offset origin.
- **The M255 headroom contract is never breached**, sampled during every gate run.
- `git push --tags` is part of shipping a tool (rung zero).

## KB dependencies

`corpus/ops/demo/build-budget.md` (M255) · `corpus/ops/demo/frontend-tier.md` ·
`corpus/ops/demo/demopatch-spec.md` · `corpus/ops/demo/demo-up-defaults.md` · `corpus/ops/rosetta_demo.md` ·
`corpus/ops/idempotency.md` · `corpus/ops/safety.md` · `corpus/ops/snapshot-spec.md` (L5)

**Delivers → `corpus/ops/demo/frontend-tier.md`** (substantially revised: real image anatomy, the multi-stage
shape, hiring's existence)
**Delivers → `corpus/ops/demo/build-budget.md`** (the achieved numbers, per host)
