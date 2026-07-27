---
active_release: "v2.8 «fast build» — IN DEVELOPMENT (branch release/02.80-fast-build, designed 2026-07-27; adversarially plan-reviewed + revised same day). The time-to-ready release: from nothing, to live, to provably live, fast. Measure the machine and spend it deliberately (build bench + two checked-in measured host profiles + one HARD headroom assert), sharpen the Playthrough suite (faster · effective · covered), collapse the demo/dev bring-up 666 s → ≤ 360 s, then bake the Playthroughs into the bring-up so a stack comes up AND proves itself. 4 milestones M255 (HARD barrier) → M256 → M257 → M258, strictly serial. Tooling + docs only, 0 platform-repo edits."
active_branch: "release/02.80-fast-build"
active_milestone: "M255 — build-bench & host-headroom (section, HARD go/no-go barrier) — BUILT; all 10 checklist items closed; BARRIER VERDICT = GO; n=3 baseline measured. Ready for /developer-kit:close-milestone."
last_closed: "M254 — 2026-07-25 (prove-on-billion; iterative, closed-on-gate); v2.7 release closed 2026-07-25"
phase: "M255 built on branch m255/build-bench-host-headroom — BARRIER PASSED (spike (a): export 146.8 s -> 2.9 s, image 4.84 GB -> 379 MB) and the GATED BASELINE MEASURED: n=3 p50 666.29 s on billion, 3/3 green, superseding the n=1 672.4 s. buildbench + measured host profiles + the union-apply fence shipped in rext at tag fast-build-m255-buildbench-2 (on origin); build-budget.md + safety.md §3.5.4 landed. Next: /developer-kit:close-milestone M255, then M256."
last_updated: "2026-07-27"
---

# State

**v2.8 "fast build" IN DEVELOPMENT** — designed 2026-07-27 via `/developer-kit:design-roadmap`, branch
`release/02.80-fast-build` cut from `main`, all 4 milestone dirs scaffolded.

**M255 — the HARD barrier — is BUILT and PASSED.** Verdict **GO**: the multi-stage `.next/standalone`
prototype takes the hiring image **4.84 GB → 379 MB** and its export step **146.8 s → 2.9 s**, so L1 does not
collapse. Two riders for M257, both measured rather than assumed: `turbo --env-mode=loose` is mandatory (the
private flag silently no-ops without it), and **L2 must be re-priced from ~200 s to ≲45 s and sequenced after
L1** — the export leg is serial and single-stream (peak load1 3.75/8, peak disk `%util` 63.4 %), and neither
host fits two concurrent Next.js build lanes. Next step: **`/developer-kit:close-milestone M255`**, then M256.

**And the measurement floor exists now.** The `n=3` campaign ran green on `billion` (3/3 `rc=0`,
`autoverify green / 0 warnings`, headroom OK): **the gated baseline is `p50 = 666.29 s`** (min 658.15, max
881.01), which supersedes the `n=1` **672.4 s** it lands within **0.9 %** of. UI-tier builds **65.5 %**,
export+unpack alone **46.2 %** — both n=1 headlines survive. The headroom model was *validated*, not merely
applied: it predicted 5,400 MiB of commitment and the reps peaked at 5,446 / 5,579 / 5,398 MB (**~3 %**).
The campaign also produced its own finding — see the reclaim correction below.

> **The release thesis.** Two standing problems, one spine — *time to ready*.
>
> **(1) A `/demo-down --purge` + `/demo-up` cycle takes 11 m 12 s.** Measured and instrumented on `billion`
> ([`evidence/build-annotation.md`](releases/02.80-fast-build/evidence/build-annotation.md), 2026-07-27, n=1):
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
> no milestone home anywhere**. Full map: [`evidence/playthrough-map.md`](releases/02.80-fast-build/evidence/playthrough-map.md).

## v2.8 shape — barrier → strictly serial → self-proving closer

```
M255 build-bench & host-headroom ── HARD BARRIER (section)   ✅ BUILT · GO   ⬅ AWAITING CLOSE
       │   measurement floor (n=3 p50 666.29 s) · headroom assert · union-apply rule · 3 spikes
       ▼
M256 playthrough sharpening (iterative)      faster · effective · covered
       ▼
M257 first-light build (iterative)           666 s → ≤ 360 s p50, cold, on billion
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

## Plan review (2026-07-27) — 23 agents, 7 lenses, 35 findings, 14 adversarially verified

**Verdict: APPROVE WITH REQUIRED EDITS — all applied.** No blocking defect survived verification; 10 findings
were refuted (several with useful residue, folded in). Six design decisions followed — **D-v28-6 … D-v28-11**,
recorded in [`roadmap.md`](roadmap.md) § "Design decisions from the adversarial plan review":

- **D-v28-6** — `hostprofile` the auto-planner **CUT** → two checked-in measured host profiles + one **failing**
  assert, **wired into M257's gate**. It had *no consumer anywhere downstream* and M257's clause read "sampled,
  not asserted" — i.e. it measured and changed nothing, **the exact defect this release retracts**.
- **D-v28-7** — the shared-clone patch race **downgraded from `blocks-release` to a paragraph + a guard test**:
  of 11 manifests, 5 are byte-identical shared, 5 target disjoint `apps/*` trees, and the 1 outlier is inert by
  its own header. Rule: **union-apply once, build parallel, revert once LIFO.**
  > **M255 CORRECTED both italicised claims.** The outlier is **not inert** — `WUNDERGRAPH_SSR_ENDPOINT` *is*
  > set on the hiring container and `apps/hiring` *does* import the patched module, so union-apply changes
  > hiring's behaviour (for the better: it inherits the M218 SSR fix). **M257 must re-verify the hiring
  > recruiter Playthrough after flipping it on.** And **neither build reverts in strict LIFO**, nor do their
  > two revert orders match — so "revert once LIFO" is not the invariant. The narrower one that *is* true is
  > now machine-fenced in both builds and both phases: the `urls.ts` chain applies studio→pubweb and reverts
  > pubweb→studio.
- **D-v28-8** — the truly-cold bench variant **CUT** (12 cold cycles ≈ 2.5–3 h on a ~4–5 cycle disk runway,
  testing the wrong hypothesis) → replaced by a **15-minute experiment on the rext-owned `hiring.Dockerfile`**.
- **D-v28-9** — M256's speed clause **re-cut**: `≤ 120 s` was arithmetically impossible (the suite is dominated
  by one irreducible ~2–3 min live-LLM test) and measured the wrong denominator.
- **D-v28-10** — the §8.5 retraction **moved to M257** (one rewrite, achieved numbers), mirror set
  **enumerated**, gated by a **grep assertion** — `demo_knob_guard.py` cannot see prose numbers.
- **D-v28-11** — **"keep it secure" now has three clauses and an owner**; the word appeared nowhere in the draft.

**Two findings survived verification and are now written into the milestones:** the **fake-FAPI global seat**
(one active seat / `signedIn` / `sessID` per stack, no cookie scoping — so M256's parallel lane needs an enabler
built, and `storageState` reuse does not isolate it), and **M258's missing world contract** (the naive
composition leaves a test world behind a presenter cockpit full of dead CTAs — exactly the state M254 left
`billion` in — and the first-draft gate passed anyway).

**Evidence base committed** to `releases/02.80-fast-build/evidence/` — `.agentspace/` is git-ignored, and both
files are cited as source-of-record from six sites.

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
| Cold `--purge` + `demo-up` cycle (billion) | **666.29 s p50** (M255, n=3, 3/3 green; supersedes the n=1 672.4 s) | **≤ 360 s p50** (M257) · stretch 300 s |
| Playthrough suite wall-clock (billion) | **228 s** (3.8 min), **dominated by one ~120 s LLM-bound test** | **median per-PT ≤ 5 s** + **post-coverage suite p50 ≤ 200 s**, LLM lane budgeted separately, 0 flake ×3 (M256) |
| Mutating Playthroughs (mutate **and read back**) | **1** of 18 (17 UNCLASSIFIED, ≥1 demonstrably mutates) | **≥ 5** (M256) |
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
- rext code-of-record: the authoring copy is on `main`, and M255's tooling ships at
  **`fast-build-m255-buildbench-2`** (buildbench + both measured host profiles + the union-apply fence + the
  `Read at` anchor fence) — **pushed to origin**. The earlier `-1`/unsuffixed tags are on origin too.
- **Rung zero:** `git push --tags` is part of shipping a tool. Verify a tag is on **origin** before any
  prove-it-live step.

_Last updated: 2026-07-27 — **M255 BUILT, barrier verdict GO**, on `m255/build-bench-host-headroom`. The
measurement floor exists: `buildbench` + two measured host profiles + a hard headroom assert + the union-apply
fence, shipped in rext at `fast-build-m255-buildbench-2` (on origin), and the **n=3 gated baseline p50
666.29 s**. Four corpus claims corrected along the way (D-v28-7's "inert" premise, `demopatch-spec.md` §4's
LIFO/4-manifest facts, `frontend-tier.md`'s argument against the barrier's own build shape, and the reclaim
protocol's `until=24h` reasoning). Next: `/developer-kit:close-milestone M255`, then M256._
