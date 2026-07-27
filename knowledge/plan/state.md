---
active_release: "v2.8 «fast build» — IN DEVELOPMENT (branch release/02.80-fast-build, designed 2026-07-27). The time-to-ready release: from nothing, to live, to provably live, fast. Measure the machine and spend it deliberately (build bench + host-capacity profiler under an explicit headroom reserve contract), sharpen the Playthrough suite (faster · effective · covered), collapse the demo/dev bring-up 672 s → ≤ 360 s, then bake the Playthroughs into the bring-up so a stack comes up AND proves itself. 4 milestones M255 (HARD barrier) → M256 → M257 → M258, strictly serial. Tooling + docs only, 0 platform-repo edits."
active_branch: "release/02.80-fast-build"
active_milestone: "M255 — build-bench & host-profiler (section, HARD go/no-go barrier) — planned, not started"
last_closed: "M254 — 2026-07-25 (prove-on-billion; iterative, closed-on-gate); v2.7 release closed 2026-07-25"
phase: "v2.8 designed + scaffolded; next step is /developer-kit:work-milestone (or :build-milestone) on M255"
last_updated: "2026-07-27"
---

# State

**v2.8 "fast build" IN DEVELOPMENT** — designed 2026-07-27 via `/developer-kit:design-roadmap`, branch
`release/02.80-fast-build` cut from `main`, all 4 milestone dirs scaffolded. **No milestone has started.**
Next step: **M255**, the HARD go/no-go barrier.

> **The release thesis.** Two standing problems, one spine — *time to ready*.
>
> **(1) A `/demo-down --purge` + `/demo-up` cycle takes 11 m 12 s.** Measured and instrumented on `billion`
> ([`.agentspace/build-annotation.md`](../../.agentspace/build-annotation.md), 2026-07-27, n=1):
> **UI-tier image builds are 66.4 % of the cycle and image export/unpack ALONE is 42.9 %** — while the box
> **never exceeded load 4.90 of 8 cores** with RAM at 74 %. This is **not** a CPU problem. It is serialised I/O
> (writing 9.4 GB of Next.js image to disk) plus a **deliberate serialisation whose stated reason no longer
> applies**: `build_frontends()` has exactly one conditional (`NO_UI`), the RAM pre-flight it cites is
> **cosmetic** (declares its vars `local`, returns no verdict, nothing branches on it), and the Go builds it was
> meant to avoid overlapping **finish 1.1 s before the UI tier starts**. Meanwhile the UI tier runs *before*
> `compose up`, so postgres boot, 4 atlas migrations, snapshot replay and the seed **idle for ~7.5 minutes**.
>
> **(2) The Playthrough suite is 18/18 green while the demo still has things that don't work.** Structural, not
> paradoxical: **1 of 18 Playthroughs proves a WRITE**; the other 17 are render-presence proofs — and on a
> *seeded* demo the read path was never the half in doubt. No Playthrough has a **negative control**. Every
> journey stops at a boundary. Every actor is `entitlement: enterprise`; **outcome `blocked`: 0, outcome
> `error`: 0** — nothing proves the platform correctly says *no*. Whole surfaces sit at zero: **ant-academy 0,
> onboarding 0, org-admin 0, talk-to-data 0**. Of the M201 curated 28 use cases, **16 are uncovered and 12 have
> no milestone home anywhere**. Full map: [`.agentspace/playthrough-map.md`](../../.agentspace/playthrough-map.md).

## v2.8 shape — barrier → strictly serial → self-proving closer

```
M255 build-bench & host-profiler ── HARD BARRIER (section)   ⬅ NEXT
       │   measurement floor · headroom contract · safe-parallelism contract · 4 spikes
       ▼
M256 playthrough sharpening (iterative)      faster · effective · covered
       ▼
M257 first-light build (iterative)           672 s → ≤ 360 s p50, cold, on billion
       ▼
M258 proven-live build (iterative, closer)   up AND self-proven, ≤ 480 s p50
```

Serial by the user's explicit order — **sharpen the detector before changing the thing it detects**. A 2-host
`M256 ∥ M257` option exists (billion holding a stable demo while M257 iterates locally) but forfeits that
benefit; the rext authoring copy is also a singleton. Full milestone detail, parallelism matrix and risk map:
[`roadmap.md`](roadmap.md) § Active — v2.8.

## Binding user decisions (2026-07-27)

- **D-v28-1** — codename **"fast build"** (user's call, overrode the proposed "call time").
- **D-v28-2** — **M255 stays a `section` barrier** (design call): three iterative milestones chasing time
  targets with one n=1 measurement between them would be guessing.
- **D-v28-3** — **batch-gate semantics: no accumulating red.** A run drives the **full batch to completion** —
  never stops at a step, never retries to hide a flake. At **batch end**, a non-empty red set **escalates to
  the user for renegotiation** (fix, or an explicit written disposition). **Zero standing red** is the
  invariant. This settles the strict-vs-non-fatal question for M258.
- **D-v28-4** — coverage scope (design call): **land onboarding (5) + org-admin (4)**; the remaining 3 un-homed
  UCs plus the 5-release-old M206/M207 reservations each get a **written verdict**. No silent gaps.
- **D-v28-5** — the **cockpit logout / Back-to-Cockpit double-click defect is FIXED** in M256 (same seat-switch
  machinery every Playthrough drives) and **deliberately gets NO Playthrough** (user's call).

## Headline numbers (inherited from the v2.7 close, 2026-07-25 — the v2.8 baseline)

- **Go:** **2019** rext test funcs. Runtime: stack-seeding 1192 pass / playthroughs 131 pass / 0 fail.
- **TypeScript (unit):** **292** executed / 0 fail (stack-verify/e2e 178 + playthroughs/e2e 114); `tsc` clean.
- **Python (rext):** demo-stack 910 pass / 0 fail / 1 skip; stack-injection 258 / 9 skip.
- **Live on billion:** Playthroughs **18/18** · content-stories **45/45 landable + 4 voice presence-only** ·
  p95 click→ACCESS **1.43 s emp / 1.41 s mgr** · studio first-paint p50 **637–726 ms**.
- **Flake: 0.** **Supply chain: GREEN.** **Platform-repo edits: 0.**

### New v2.8 baselines to beat

| | Baseline | v2.8 target |
|---|---|---|
| Cold `--purge` + `demo-up` cycle (billion) | **672 s** | **≤ 360 s p50** (M257) · stretch 300 s |
| Playthrough suite wall-clock (billion) | **228 s** (3.8 min) | **≤ 120 s p50**, 0 flake ×3 (M256) |
| Mutating Playthroughs | **1** of 18 | **≥ 5** (M256) |
| `blocked` / `error` outcomes | **0** | **≥ 1 `blocked`** (M256) |
| Curated UCs with no milestone home | **12** | **0** — landed or written verdict (M256) |
| Composed up-and-proven cycle | *does not exist* | **≤ 480 s p50**, zero standing red (M258) |

## Recently shipped releases (max 3; older → roadmap.md / roadmap-legacy.md)
- **v2.7 "july jitter" — 2026-07-25** (tag `v2.7`) — re-ground + fidelity + field-hardening; 9 milestones
  M246→M254; prove-on-billion a–h live; **zero carry-forward**; 0 platform edits.
- **v2.6 "sound check" — 2026-07-23** (tag `v2.6`) — reliability / field-hardening; 8 milestones M237→M244.
- **v2.5 "the playbill" — 2026-07-20** (tag `v2.5`) — content-vantage; 8 milestones M229→M236.

## Standing backlog (fated destinations)
- **Consumed by v2.8:** the reserved Playthroughs futures **M206** (ai-sim mirror tier + the 3 employee
  deepening legs) and **M207** (academy coverage) are **re-fated inside M256's clause 3** — each gets a written
  verdict (named future milestone or drop) rather than a sixth consecutive re-reservation.
- **DROPPED:** DEF-M250-01 `participants_filter` (D18) · DEF-M215-03(a)/F11 (design-time) · DEF-M239-01 (v2.6).
- **Still unscheduled (vision):** DEF-M10-01 (S3/Bunny voice media — voice presence-only) · DEF-M21-01 ·
  CAVEAT-1 · M314b (platform) · **M205** residual (tier gates + ATS).

## Process flags (do NOT auto-push)
- **v2.7 is merged to `main` + tagged `v2.7` LOCALLY; NOT pushed to origin** — the user runs origin publishes
  on their own cadence. **v2.5**'s and **v2.6**'s merges + tags are likewise local-only.
- **A stray `(M245)` commit** sits on `main` (post-v2.6 academy docs, untracked in the plan).
- rext code-of-record: the authoring copy is at `july-jitter-v27-close-followups @ a5b1288`, on origin. The
  `fix/studio` studio-FAPI fix is at `rosetta-extensions@bc65850`.
- **Rung zero:** `git push --tags` is part of shipping a tool. Verify a tag is on **origin** before any
  prove-it-live step.

_Last updated: 2026-07-27 — v2.8 "fast build" DESIGNED + scaffolded (branch cut, 4 milestone dirs, 5 binding
decisions). Next: M255, the barrier._
