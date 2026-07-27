---
milestone_shape: section
milestone: M255
title: "build-bench & host-profiler (the barrier)"
status: planned
release: v2.8 "fast build"
depends_on: []
parallel_with: []
complexity: large
barrier: true
created: 2026-07-27
last_updated: 2026-07-27
---

# M255 — build-bench & host-profiler  (`section`, HARD go/no-go barrier)

**Status:** `planned` · **Shape:** `section` · **Complexity:** large · **Release:** v2.8 "fast build"
**Depends on:** — (it is the barrier)

## Goal

Establish the **measurement floor** and the **safety contract** every later speed lever is judged against —
including the profiler that reads the machine it is running on and decides how much of it to use.

## Why this is a barrier

Everything after this milestone is a **number**, and today there is exactly **one measurement, at n=1**
([`.agentspace/build-annotation.md`](../../../../.agentspace/build-annotation.md), `billion`, 2026-07-27,
cold-images / warm-layer-cache). Three of the four ranked levers have their value gated on an unanswered
question. Committing M257's exit gate before answering them would be guessing.

## Scope

### In

1. **`buildbench`** — rext `stack-core`. Generalize the one-off `reset-clean-build.sh` driver into a
   repeatable, instrumented harness:
   - **n ≥ 3** per configuration.
   - Two variants: **cold-images** (`--purge`, layer cache warm — the sanctioned path the annotation measured)
     and **truly-cold** (`docker builder prune -af` first — never yet measured).
   - Machine-readable per-phase ledger (`@@PHASE@@` markers → JSON) + the 10 s resource sampler.
   - Runs on **`billion`** and on the **laptop**.

2. **`hostprofile`** — rext `stack-core`. Probe cores / RAM / swap / free disk / IO and emit a **build plan**:
   UI-tier lane count · `turbo --concurrency` · per-lane V8 heap ceiling · BuildKit parallelism.
   Under a written **headroom reserve contract**:
   - hold back **≥ 2 cores or 20 %, whichever is greater**;
   - keep total RAM commitment **≤ 80 %**, counting the **summed `--max-old-space-size` ceilings**, not just
     current usage (today only `next-web` + `hiring` carry `4096`, and it is a V8 *ceiling*, not a reservation);
   - **refuse and downshift loudly** rather than over-subscribe;
   - abort if free disk < floor + projected image bytes;
   - on macOS read the **Docker VM allocation**, never host totals.

3. **The safe-parallelism contract** — written, before any lever uses it. The load-bearing item: `next-web` and
   `hiring` build from the **same clone** `ctx=$DEMO/next-web-app` (`up-injected.sh:490`, `:1008`) with
   **different demopatch sets**. Naive parallelism races patch apply/revert and defeats **G2** (drift-refuse),
   **G4** (idempotent), **G5** (self-revert). Resolve as: separate build-scratch clones, or a patch-set
   superset with per-image filtering. Also state the per-stack image-isolation invariant (cache **layers**,
   never **images**).

4. **Four spikes**, each answered with evidence:
   - **(a)** Does `next-web` already have a **multi-stage / production Dockerfile sibling** upstream? *If yes,
     L1 — the single biggest lever — is a **selection** change, not a new Dockerfile.*
   - **(b)** The **truly-cold** baseline (sizes the worst case and L3's real value).
   - **(c)** The **laptop** baseline (M1 Pro, 10 cores / 16 GiB, arm64; Docker VM allocation is the real budget).
   - **(d)** Is peak load1 **4.90 of 8** a plateau or an **I/O ceiling**? *If BuildKit is I/O-throttled, L2's
     parallel win is smaller than the naive ~200 s.*

5. **`corpus/ops/demo/build-budget.md`** — net-new; the §0b blind area. What "fast" means for a bring-up, the
   per-phase attribution model, the baseline, the gate, the headroom contract. Modelled on
   `latency-budget.md` (including its **state the environment with every number** rule).

6. **The §8.5 corpus retraction**, under the v2.7 **C1 mirrored-count rule** (all four docs in one commit;
   `demo-up-defaults.md` is machine-fenced both ways by `stack-core/demo_knob_guard.py`):
   - *"the ~3.7 GB build cache"* → **105.4 GB** (~28× off) — and `DEMO_DISK_MIN_GIB=20` is sized against the
     wrong number;
   - *"~3 min per frontend"* → right for the two Next apps, **~7× wrong** for studio-desk; and
     `frontend-tier.md` mentions **"hiring" zero times in 623 lines**, so the total undercounts by a whole
     208 s image;
   - *"~3.7 GB first build"* (`up-injected.sh:794`) → measured **4.77 GB** / **4.67 GB**;
   - studio *"pure memory starvation, **not a slow build**"* → refuted (export/unpack is 288.4 s; the box never
     exceeded load 4.90/8).

7. **The §8.6 cert hazard.** `$STACK/certs` **survives `--purge`** and the whole mint block is guarded on
   `[ ! -f $CERTS/fapi.crt ]` (`up-injected.sh:1859`), so billion's `tailscale cert` minted **2026-07-11** has
   never been re-minted. A 90-day cert silently expires around **2026-10-09**. Ship an expiry-aware re-mint.

### Out

- Any actual speed lever — that is M257.
- Host `daemon.json` changes. Disabling the containerd snapshotter would orphan all 26 existing images — a bad
  trade, and it is host config, not rext, not demopatch.
- Dev-path-specific work (M257 reports it; nothing here is dev-only).

## Barrier condition

**If the truly-cold measurement shows the export/unpack cost is NOT image-size-driven, L1 collapses** and
M257's exit gate must be re-cut **with the user** before M257 starts.

## KB dependencies

`corpus/ops/demo/frontend-tier.md` · `corpus/ops/demo/demopatch-spec.md` · `corpus/ops/demo/demo-up-defaults.md`
· `corpus/ops/demo/latency-budget.md` (the budget-doc precedent) · `corpus/ops/verification.md` ·
`corpus/ops/safety.md` · `corpus/ops/rosetta_demo.md`

**Delivers → `corpus/ops/demo/build-budget.md`** (net-new)
**Delivers → `corpus/ops/demo/frontend-tier.md`** (the §8.5 retraction, mirrored ×4)

## Source

[`.agentspace/build-annotation.md`](../../../../.agentspace/build-annotation.md) — the measured, instrumented
phase-by-phase anatomy (2026-07-27, `billion`, n=1) this whole release is designed from.
