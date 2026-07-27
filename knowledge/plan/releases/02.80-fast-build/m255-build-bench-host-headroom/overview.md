---
milestone_shape: section
milestone: M255
title: "build-bench & host-headroom (the barrier)"
status: planned
release: v2.8 "fast build"
depends_on: []
parallel_with: []
complexity: medium-large
barrier: true
created: 2026-07-27
last_updated: 2026-07-27
---

# M255 — build-bench & host-headroom  (`section`, HARD go/no-go barrier)

**Status:** `planned` · **Shape:** `section` · **Complexity:** medium-large · **Release:** v2.8 "fast build"
**Depends on:** — (it is the barrier)

> **Revised 2026-07-27** after a 23-agent adversarial plan review. Three deliverables were **cut or replaced**
> (`hostprofile` the auto-planner → D-v28-6; the truly-cold bench variant → D-v28-8; the §8.5 prose retraction
> → moved to M257, D-v28-10), one **downgraded** from an architecture decision to a paragraph + a guard test
> (D-v28-7), and two **added** (a campaign/reclaim protocol, and security). Rationale for each lives in
> [`roadmap.md`](../../../roadmap.md) § "Design decisions from the adversarial plan review".

## Goal

Establish the **measurement floor**, the **headroom assert**, and the **parallelism rule** that every later
speed lever is judged against.

## Why this is a barrier

Everything after this milestone is a **number**, and today there is exactly **one measurement, at n=1**
([`../evidence/build-annotation.md`](../evidence/build-annotation.md), `billion`, 2026-07-27, cold-images /
warm-layer-cache). The largest lever (L1) has its value gated on an unanswered question. Committing M257's
exit gate before answering it would be guessing.

## Scope

### In

1. **`buildbench`** — rext `stack-core`. Generalize the one-off `reset-clean-build.sh` driver into a
   repeatable, instrumented harness:
   - **n ≥ 3 on `billion`, cold-images variant only** (`--purge`, layer cache warm — the sanctioned path the
     annotation measured). **The truly-cold variant is CUT** (D-v28-8, see Out).
   - Machine-readable per-phase ledger (`@@PHASE@@` markers → JSON) + the 10 s resource sampler.
   - **Every ledger entry records the invocation and a full `DEMO_*` env snapshot**, so a gate run is
     self-describing. Neither `autoverify.json` (which emits only `project/offset/warnings/green/ts`,
     `stack-verify/live/autoverify.sh:381-385`) nor the phase log currently records which services were in
     scope — so a UI-on cycle is indistinguishable from a `DEMO_NO_UI=1` cycle after the fact.
   - **One informational `n=1` laptop run** — recorded in `build-budget.md`, **not gated** (every exit gate in
     the release is measured on `billion`).

2. **Campaign protocol + reclaim.** The first draft's campaign was **not executable**: each rep leaks ~2 G of
   disk and orphans ~11.6 G of cache, against a **~4–5 cycle runway** on billion's 38–40 G free — while the
   guard that should catch it (`DEMO_DISK_MIN_GIB=20`, `up-injected.sh:277` at design time) is **both
   mis-sized and non-fatal** ("non-fatal: continuing"), and a mid-campaign ENOSPC presents as the cryptic
   *"redis exited (1)"* (M239-F1) — i.e. it looks like a lever broke something. Deliver:
   - a **hard-failing** pre-rep disk/cache assert (`floor + projected image bytes`);
   - an explicit **reclaim step between reps** — **L6 (scheduled BuildKit prune) is promoted from M257's lever
     table into this milestone as campaign hygiene**;
   - per-rep declaration of the starting `docker system df` state;
   - **`DEMO_DISK_MIN_GIB` re-sized** in the same commit (the number it is sized against is ~28× wrong);
   - a note in `build-budget.md` naming the ENOSPC → *"redis exited (1)"* signature.

3. **Host profiles + headroom assert** (D-v28-6 — replaces the `hostprofile` auto-planner):
   - `stack-core/hostprofiles/billion.json` + `laptop.json` — **measured, checked in**. On macOS the budget is
     the **Docker VM allocation**, never host totals.
   - **One assert** in buildbench's existing sampler that **FAILS** when: peak load1 > cores − 2, **or** peak
     summed heap commitment > 80 % of the host budget, **or** free disk < floor + projected image bytes.
   - **Record a decision** in `decisions.md` reconciling *"fail loudly"* against the codebase's standing
     **never-block-a-bring-up** pre-flight contract: the assert gates **buildbench and
     the M257 gate**, it does **not** block an operator's bring-up. Those are different contracts and the
     first draft conflated them.
   - **Note the seam that does NOT work:** `ENV NODE_OPTIONS` cannot lower the per-lane V8 ceiling —
     `apps/web/package.json:98` and `apps/hiring/package.json:92` re-assign
     `NODE_OPTIONS=--max_old_space_size=8192` **inline** for the `next build` child, overriding the Dockerfile
     ENV. (The first draft's `4096` was the ENV, not the effective ceiling.)

4. **The union-apply parallelism rule** (D-v28-7 — downgraded from an architecture decision). `next-web` and
   `hiring` build from the **same clone** `ctx=$DEMO/next-web-app` (`up-injected.sh:490`, `:1008`) with
   different demopatch sets, which naive parallelism would race against **G2** (drift-refuse), **G4**
   (idempotent) and **G5** (self-revert). Inspection settles it without an architecture: of the **11 distinct
   manifests** (`:496-509` next-web ×9, `:1020-1047` hiring ×7), **5 are the same manifests on the same shared
   files**, **5 target disjoint trees** (`apps/web/**` ×3 vs `apps/hiring/**` ×2), and the one shared-package
   outlier — `next-web-ssr-graphql-origin` → `packages/graphql/src/server/server.graphql.ts` — is **inert for
   the hiring image by its own manifest header** (behaviour-identical when `WUNDERGRAPH_SSR_ENDPOINT` is
   unset; it only prepends to an existing `||` chain).
   **Deliver:** the rule — *apply the union once, build both images in parallel from the single clone, revert
   once LIFO* (preserving the studio→pubweb `urls.ts` chain, and removing one apply/revert cycle of the chained
   pair — exactly where G2 drift-refusals historically bite) — **plus a guard test** asserting (i) the two
   lists' shared members are byte-identical manifests and (ii) every non-shared member's `path:` is under a
   disjoint `apps/*` tree or explicitly waived as inert. **Delete the "separate build-scratch clones" option.**
   Also state the per-stack image-isolation invariant (cache **layers**, never **images**) that M257 then
   asserts.

5. **Three spikes**, each answered with evidence:
   - **(a) The 15-minute L1 experiment — the barrier's decider.** Prototype the multi-stage shape on
     **`demo-stack/frontend/hiring.Dockerfile`, which rext already owns outright** (no demopatch, no
     platform-edit question), and **measure the export delta directly**. This **replaces** the truly-cold
     campaign as the barrier test — it answers the L1 hypothesis with real evidence and banks ~100 s of the
     lever, instead of spending ~2.5–3 h of serial machine time before a second is saved.
     Record that standalone output is enabled by **`ENV NEXT_PRIVATE_STANDALONE=1`** — Next 16's frozen
     `defaultConfig` reads `output: !!process.env.NEXT_PRIVATE_STANDALONE ? 'standalone' : undefined`, and **no
     app `next.config` sets `output`** (verified across all four), so this needs **zero source edits and zero
     demopatches**. It is a Next-**private** API; the fallback is a `next.config.mjs` demopatch per app.
   - **(d) Plateau or I/O ceiling?** Is peak load1 4.90/8 a plateau — if BuildKit is I/O-throttled, L2's
     parallel win is smaller than the naive ~200 s, and M257's gate must be re-cut.
   - **(e) NEW — host-vs-peer topology for M258.** A `--public-host` demo **cannot be browsed from its own
     host**: docker-proxy binds `0.0.0.0`, so a connection from the host to its own tailscale IP hits the
     kernel socket and **bypasses `tailscale serve`** (which terminates TLS) → `ERR_SSL_PROTOCOL_ERROR`
     (`run-playthroughs.sh:56-72`). And `--public-host` is **default-on** for the demo path (D-DESIGN-3). So
     M258's *"one cold command on billion"* may be unachievable as literally worded, or achievable only under
     `--no-public-host` — which proves the composition in a mode the presenter never uses. A ~20-minute answer
     that determines M258's gate text, moved forward here because the barrier already runs spikes.

6. **`corpus/ops/demo/build-budget.md`** — net-new; the §0b blind area. What "fast" means for a bring-up, the
   per-phase attribution model, the baseline, the gate, the headroom contract, the ENOSPC signature, and the
   `NEXT_PRIVATE_STANDALONE` note. Modelled on `latency-budget.md`, including its **state the environment with
   every number** rule.

7. **Security (D-v28-11) + the §8.6 cert hazard — explicitly NON-GATING hygiene.** `$STACK/certs` **survives
   `--purge`** and the whole mint block is guarded on `[ ! -f $CERTS/fapi.crt ]` alone,
   so billion's `tailscale cert` minted **2026-07-11** has never been re-minted. A 90-day cert silently expires
   around **2026-10-09** — and the failure path drops to `gen_local_fapi_cert` with **only a warning**, so a
   remote browser silently loses trust. Ship:
   - the **expiry-aware re-mint**, and
   - the paired **`corpus/ops/safety.md` §3 amendment** covering the renewal path and that silent-fallback
     failure mode.
   *Kept here rather than routed to `/developer-kit:bugfix` because it pairs with the safety amendment; marked
   non-gating so it never sits on the barrier's critical path.*

### Out

- **Any actual speed lever** — that is M257.
- **The §8.5 prose retraction** — moved to M257 (D-v28-10), so `frontend-tier.md` is rewritten **once** with
  achieved numbers rather than twice. *(The `DEMO_DISK_MIN_GIB` re-size stays here — item 2 needs it.)*
- **The truly-cold bench variant** (D-v28-8). It doubled the campaign and tested the wrong hypothesis: the warm
  run already yields three (size, unpack) points — **8.03 / 8.05 / 5.73 s per GB** — and studio-desk paid
  **9.8 s of unpack with zero new layer bytes exported**, which is precisely the size-driven, cache-independent
  signature the barrier needed. Survives as an **optional one-shot after M257** if L3's value is disputed.
- **Host `daemon.json` changes.** Disabling the containerd snapshotter would orphan all 26 existing images —
  and it is host config, not rext, not demopatch.
- Dev-path-specific work (M257 reports it; nothing here is dev-only).

## Barrier condition

**Spike (a) is the decider. If the multi-stage prototype on `hiring.Dockerfile` does not materially cut the
export leg, L1 collapses** and M257's exit gate must be re-cut **with the user** before M257 starts.

## KB dependencies

`corpus/ops/demo/frontend-tier.md` · `corpus/ops/demo/demopatch-spec.md` ·
`corpus/ops/demo/latency-budget.md` (the budget-doc precedent) · `corpus/ops/verification.md` ·
`corpus/ops/safety.md` · `corpus/ops/demo/tailscale-serve.md` · `corpus/ops/rosetta_demo.md`

**Delivers → `corpus/ops/demo/build-budget.md`** (net-new)
**Delivers → `corpus/ops/safety.md`** (§3 cert-renewal amendment — the release's named security owner)

## Source

[`../evidence/build-annotation.md`](../evidence/build-annotation.md) — the measured, instrumented
phase-by-phase anatomy (2026-07-27, `billion`, n=1) this whole release is designed from. **Committed into the
release directory** at the plan review's request: it is cited as source-of-record from six sites, and
`.agentspace/` is git-ignored (`.gitignore:138`), so the barrier verdict would otherwise be underivable and
`close-release` would archive dangling links.
