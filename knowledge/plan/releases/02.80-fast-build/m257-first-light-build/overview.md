---
milestone_shape: iterative
milestone: M257
title: "first-light build"
status: planned
release: v2.8 "fast build"
exit_gate: "A cold-images `demo-down --purge` + `demo-up` reaches `autoverify green:true / 0 warnings` in p50 <= 360 s across 3 consecutive cycles on billion (baseline: the M255-measured n=3 p50 666.29 s — a 46% cut), 0 platform-repo edits, all 7 demopatch guards (G1-G7) passing, AND two FALSIFIABLE asserts that FAIL the gate when tripped (D-v28-6, D-v28-11): HEADROOM — peak load1 <= cores-2 AND peak summed heap commitment <= 80% of the host budget AND free disk >= floor + projected image bytes, read from the sampler (NOT 'sampled, not asserted'); ISOLATION — no built image contains another stack's baked publishable key or offset origin, asserted by post-build image inspect (L1/L3 change exactly the layers that carry them). Stretch: <= 300 s."
iteration_protocol_ref: corpus/ops/demo/build-budget.md
re_scope_trigger: "If after L1 + L2 + L3 the p50 is still > 420 s, escalate rather than grind. RE-DERIVED at the M255 harden against the real arithmetic: the gate needs 666.29 - 360 = 306.29 s, and once M255 re-priced L2 from ~200 s to <=45 s the three big levers are worth 200-250 (L1) + <=45 (L2) + ~55 (L3) = 300-350 s -- so at L1's conservative end they miss the gate BY THEMSELVES. A p50 above 420 s means the three delivered under ~250 s, i.e. at least 50 s short of even their low estimate, and the remaining small levers (L4/L5/L7/L8/L10, ~93-158 s combined) are then being asked to cover a shortfall they were never priced for: the residual is structural (host I/O, the containerd snapshotter) or a lever did not land. The ORIGINAL trigger was 480 s, set when L2 looked like ~200 s and the three looked like ~505 s against a 306 s need -- against the measured prices it would only have fired if the levers returned under 186 s, barely half, so a 60 s gate miss could never have tripped it."
depends_on: [M256]
parallel_with: []
complexity: very-large
created: 2026-07-27
last_updated: 2026-07-27
---

# M257 — first-light build  (`iterative`)

**Status:** `planned` · **Shape:** `iterative` · **Complexity:** very-large · **Release:** v2.8 "fast build"
**Depends on:** M256 (sharpen the detector before changing what it detects)

> **Revised 2026-07-27** after the adversarial plan review: the gate's headroom clause became **falsifiable**
> (it read "sampled, not asserted" — i.e. it measured and changed nothing, the exact defect this release
> retracts), an **image-isolation assert** was added, the **§8.5 corpus retraction moved here** from M255, L6
> moved **out** to M255, and the dev-path rationale was corrected. See [`roadmap.md`](../../../roadmap.md)
> § "Design decisions from the adversarial plan review" (D-v28-6, D-v28-10, D-v28-11).

## Goal

Collapse the cold demo bring-up so going live is a coffee, not a lunch — **spending the machine deliberately,
never exhausting it**, and without weakening a single safety guard.

## Baseline (`billion` 8 vCPU / 7.3 GiB / x86_64, cold-images + warm layer cache)

**The gate measures against the M255 `n=3` campaign, not the `n=1` annotation.** Both are listed because the
annotation is where the lever ranking comes from; the campaign is what the exit gate is a percentage of.

| | **n=3 p50 — THE BASELINE** | n=1 annotation |
|---|---|---|
| Total cycle | **666.29 s (11 m 06 s)** | 672.4 s (11 m 12 s) |
| Bring-up | 633.15 s (95.0 %) | 650.7 s (96.8 %) |
| **UI-tier image builds (3)** | **436.1 s — 65.5 %** | 446.4 s — 66.4 % |
| **Image export/unpack alone** | **307.5 s — 46.2 %** | 288.4 s — 42.9 % |
| peak load1 | **4.06 / 4.56 / 4.22 of 8** | 4.90 of 8 · avg 2.26 · peak RAM 74 % |

The two agree to **0.9 %**. Campaign artefacts: `billion:/home/devops/panorama/m255/campaign/`; the protocol,
the reclaim caveat and the per-sub-phase table are in
[`corpus/ops/demo/build-budget.md`](../../../../corpus/ops/demo/build-budget.md).

> **Two M255 findings this milestone must carry.** (1) **`turbo --env-mode=loose` is mandatory** — Turbo 2
> defaults to `strict` and filters `NEXT_PRIVATE_STANDALONE` out before `next build` sees it, so the flag
> silently no-ops and the build stays green with the old 4.84 GB image. (2) **L2 is re-priced down to ≲45 s and
> sequenced AFTER L1** (L1 deletes the two export legs L2 existed to overlap), and **the hiring recruiter
> Playthrough must be re-verified after union-apply is flipped on** — D-v28-7's "inert outlier" premise was
> false; hiring's behaviour does change.

**This is not a CPU problem.** It is serialised I/O (writing 9.4 GB of Next.js image to disk) plus a
**deliberate serialisation**: `build_frontends()` has exactly one conditional (`NO_UI`); the RAM pre-flight it
cites is **cosmetic** (`preflight_vm_ram()` declares its vars `local`, assigns no global, returns no verdict);
and the Go builds it was meant to avoid overlapping **finish 1.1 s before the UI tier starts**. Meanwhile the
UI tier runs at `:1877` and `compose up` at `:1924`, so postgres boot, **4 atlas migrations**, snapshot replay
and the seed **idle for ~7.5 minutes**.

## Levers, ranked by measured seconds recoverable

| | Lever | Est. saving | Shape |
|---|---|---|---|
| **L1** | Multi-stage the two Next images — ship `.next/standalone` + static instead of the full build tree with dev deps. 4.77 GB → a few hundred MB collapses the 141.9 s + 136.7 s export **and** the 85.7 s unconditional unpack leg (L9) | **~200–250 s** | rext-owned Dockerfile. **No config edit and no demopatch needed**: `ENV NEXT_PRIVATE_STANDALONE=1` flips Next 16's frozen `defaultConfig` (`output: !!process.env.NEXT_PRIVATE_STANDALONE ? \'standalone\' : undefined`) because **no app `next.config` sets `output`** (verified ×4). Private Next API; fallback = a `next.config.mjs` demopatch per app |
| **L2** | Build `next-web` ∥ `hiring` **and** reorder the UI tier to overlap `compose up`. **Sequenced AFTER L1**, which deletes the two ~140 s export legs L2 existed to overlap | **≲45 s** — re-priced DOWN from ~200 s by M255 spike (d), **measured, not estimated** | rext `up-injected.sh` under M255's **union-apply rule** (D-v28-7): apply the union of both manifest sets once, build both in parallel from the one clone, revert once — with the `urls.ts` chain reverted **pubweb-before-studio** (D-M255-4: *neither* build reverts in strict LIFO and the two orders differ, so "revert LIFO" was never the invariant; `union_apply_guard.py` asserts the real one). **Only the export legs can overlap** — the headroom assert derives `max_parallel_ui_lanes = 1`, so the two *compile* legs cannot run concurrently on `billion` at all |
| **L3** | Manifests-first `COPY` so the `pnpm install` layer survives a source-only change (every demopatch is one). The layer has **never once been reused** — 16 entries × 4.029 GB = **61 % of the whole build cache**, every one `Usage count: 1` | **~55 s** + an ~8 GB/cycle leak | demopatch / rext Dockerfile. Must copy root `package.json` + `pnpm-lock.yaml` + `pnpm-workspace.yaml` **plus all 16 workspace `package.json`s** — a naive `COPY package*.json ./` breaks `--frozen-lockfile` |
| **L4** | Drop `--concurrency=1` from `pnpm turbo build` — the value comes **from the checked-in host profile**, not a hardcoded 8-core assumption (D-v28-6) | ~20–35 s | build-arg or demopatch |
| **L5** | Speed the taxonomy replay (78.0 s / 330,261 rows + 2 pgvector reindexes): index-after-COPY in one pass, `UNLOGGED`-then-`SET LOGGED`, or a pre-built PG data dir. **The chief win on the `/dev-up` path** (`dev-setdress.sh:299`/`:357` run the same `stacksnap replay`) | ~30–50 s | rext `stack-snapshot` |
| **L7** | Multi-stage `studio-desk` — 1.71 GB shipping a full dev toolchain (32,568 JS/CSS files, 266 MB) to serve a Vite bundle | ~8 s | demopatch / rext Dockerfile |
| **L8** | Cache the Directus bootstrap + restart — 15.6 s of pure container-boot latency, the most compressible slice of set-dress | ~15 s | rext |
| **L10** | Serial fat: ~12 serial `git fetch`es · 23 serial `demopatch revert` shells · Go tooling compiled 4–5×/bring-up (`stacksecrets` into a throwaway `mktemp -d`, `stackseed` **twice**) · 4 independent `atlas migrate apply` targets run serially · the entire tailscale-serve plan re-emitted to add one port | ~20–50 s | rext |

*(**L6** — scheduled BuildKit prune — **moved to M255** as campaign hygiene: M255's own bench campaign is the
first thing that would exhaust the ~4–5 cycle disk runway. **L9** — the 85.7 s unconditional unpack leg — is
folded into L1; it is not a build flag.)*

> **L1 + L2 + L3 buy ~300–350 s against the 306.29 s the gate needs — and that margin is THIN.** They attack
> the same 436 s block from three angles (smaller images to export, exported concurrently, with dependency
> layers that actually survive), taking the cycle from 11 m 06 s to roughly **5–6 m**. But do the arithmetic
> before trusting the headline: **666.29 → 360 needs 306.29 s**, and after M255 re-priced L2 from ~200 s to
> **≲45 s** the big three are worth **200–250 (L1) + ≲45 (L2) + ~55 (L3)** — i.e. **300 s at L1's
> conservative end, which MISSES the gate on its own.** The first draft of this section was written when L2
> looked like ~200 s and the three therefore looked like ~505 s, a 200 s cushion. There is no cushion.
> **L4 / L5 / L7 / L8 / L10 (~93–158 s combined) are load-bearing, not garnish** — plan to land some of them,
> and treat "L1+L2+L3 and we are done" as the optimistic branch rather than the plan.

## Also relevant (operational, not a lever)

**`--purge` defeats every image cache, including 5 hidden ones.** `rosetta-demo:336-341` removes the three UI
images **and** 5 that `compose up` then rebuilds inline (postgresql, graphql, sentinel, storage, roadrunner) —
15.3 s warm here, 120–300 s if truly cold — with **no per-service log file to attribute them to**. So the three
cache-reuse checks (`:562`, `:849`, `:1077`) can **never** hit on a purge cycle. Plain `rosetta-demo down N`
(no `--purge`) keeps the images and makes a re-up cost seconds — the fast-cycle option whenever a wiped DB is
not required. **Document it; do not make it the default** (a wiped DB is usually the point).

## Also in scope — the §8.5 corpus retraction (D-v28-10, moved here from M255)

Landing **once**, with the *achieved* numbers, so `frontend-tier.md` is rewritten a single time.
**Enumerated** mirror set (the first draft said "all four docs" and named none):
`corpus/ops/demo/frontend-tier.md` **×4 sites** — `:231`, `:249`, `:262`, `:271` — plus
`corpus/ops/demo/README.md:139` and `CLAUDE.md:318`.
**Gated by a grep assertion** for the retracted strings: the first draft cited
`stack-core/demo_knob_guard.py` as the machine fence, but that guard matches `${DEMO_*:-default}` knobs and
`case` arms and **structurally cannot see prose numbers**. `demo-up-defaults.md` carries none of these claims
and is **not** in the set. The claims:
*"the ~3.7 GB build cache"* → **105.4 GB** (~28× off) · *"~3 min per frontend"* → right for the two Next apps,
**~7× wrong** for studio-desk, and `frontend-tier.md` mentions **"hiring" zero times in 623 lines** ·
*"~3.7 GB first build"* (`up-injected.sh:794`) → measured **4.77 / 4.67 GB** · studio *"pure memory
starvation, not a slow build"* → refuted (export/unpack is 288.4 s; the box never exceeded load 4.90/8).

## Dev path

`/dev-up` shares **L5 / L10** (set-dress + tooling, not the UI tier). **Measured and reported at each iter;
not separately gated** — **because the UI tier has no dev counterpart**: the main dev stack runs next-web
**natively** (`dev-up` SKILL.md:69-76) and `dev-N` defaults to the frontend-free `graphql` profile, so the
446 s / 66.4 % block simply does not exist there. *(The first draft justified this with "the demo path is
where 96.8 % of the wall-clock is" — a misuse: 96.8 % is bring-up as a share of the **demo** cycle, which says
nothing about demo-vs-dev.)*

## Shape (why iterative)

L1's cost depends on M255 spike (a) · L2's real win depends on spike (d) · L3's value is bounded by the
measured 61 %-of-cache figure. The path is measurement-driven by construction: measure → attribute → one lever
→ re-measure at n ≥ 3.

## Hard constraints

- **Zero platform-repo edits.** L1/L3/L4/L7 all touch Dockerfiles in canonical repos → each lands as a
  sha-pinned `demopatch` or an **rext-owned Dockerfile** in the shape `hiring.Dockerfile` already sanctions.
- **All 7 demopatch guards (G1–G7) still pass** after every lever.
- **Per-stack image isolation** — cache **layers**, never **images**. A reused image would carry another
  stack's baked publishable key and offset origin. **Now a falsifiable gate clause** (D-v28-11): asserted by
  post-build image inspect, because L1/L3 change exactly the layers that carry them.
- **The M255 headroom assert is a gate clause, not an observation** (D-v28-6) — peak load1, peak summed heap
  commitment and free disk are read from the sampler and **FAIL** the gate when breached.
- **`ENV NODE_OPTIONS` is not a usable seam** for a per-lane V8 ceiling: `apps/web/package.json:98` and
  `apps/hiring/package.json:92` re-assign `--max_old_space_size=8192` **inline** for the `next build` child.
- `git push --tags` is part of shipping a tool (rung zero).

## KB dependencies

`corpus/ops/demo/build-budget.md` (M255) · `corpus/ops/demo/frontend-tier.md` ·
`corpus/ops/demo/demopatch-spec.md` · `corpus/ops/demo/demo-up-defaults.md` · `corpus/ops/rosetta_demo.md` ·
`corpus/ops/idempotency.md` · `corpus/ops/safety.md` · `corpus/ops/snapshot-spec.md` (L5)

**Delivers → `corpus/ops/demo/frontend-tier.md`** (rewritten **once**, with achieved numbers: real image
anatomy, the multi-stage shape, hiring's existence — plus the enumerated §8.5 retraction + its grep gate)
**Delivers → `corpus/ops/demo/build-budget.md`** (the achieved numbers, per host)
