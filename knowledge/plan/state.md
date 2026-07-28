---
active_release: "v2.8 «fast build» — IN DEVELOPMENT (branch release/02.80-fast-build, designed 2026-07-27; adversarially plan-reviewed + revised same day). The time-to-ready release: from nothing, to live, to provably live, fast. Measure the machine and spend it deliberately (build bench + two checked-in measured host profiles + one HARD headroom assert), sharpen the Playthrough suite (faster · effective · covered), collapse the demo/dev bring-up 666 s → ≤ 360 s, then bake the Playthroughs into the bring-up so a stack comes up AND proves itself. 4 milestones M255 (HARD barrier) → M256 → M257 → M258, strictly serial. Tooling + docs only, 0 platform-repo edits."
active_branch: "release/02.80-fast-build"
active_milestone: "M256 — playthrough sharpening (iterative) — PLANNED, not started. Bootstrap tok is pre-seeded: releases/02.80-fast-build/evidence/playthrough-map.md"
last_closed: "M255 — 2026-07-28 (build-bench & host-headroom; section, HARD barrier, VERDICT GO)"
phase: "M255 CLOSED + merged into release/02.80-fast-build. Next: M256 (iterative) via /developer-kit:work-mstone-iters or :build-mstone-iters. M256's code half is fully local; only its suite-p50 clause needs a host (billion = standing sign-off rule)."
last_updated: "2026-07-28"
---

# State

**v2.8 "fast build" IN DEVELOPMENT** — designed 2026-07-27 via `/developer-kit:design-roadmap`, branch
`release/02.80-fast-build` cut from `main`, all 4 milestone dirs scaffolded.

**M255 — the HARD barrier — is CLOSED. Verdict GO.** Merged into `release/02.80-fast-build` on 2026-07-28.
The measurement floor exists (gated baseline **n=3 p50 `666.29 s`**, authoritative in
`hostprofiles/billion.json`) and L1 is proven real (hiring image **4.84 GB → 379 MB**, export leg
**146.8 s → 2.9 s**). Full closure narrative: [`roadmap.md`](roadmap.md) § M255 → Closure.

**Next: M256 — playthrough sharpening** (`iterative`, not started). Its bootstrap tok is **pre-seeded** —
`releases/02.80-fast-build/evidence/playthrough-map.md`. Its **code half is entirely local** (the fake-FAPI
seat-isolation enabler, the per-spec mutation classification, the negative-control mechanism, the onboarding ×5
+ org-admin ×4 Playthroughs); only clause 1's suite-p50 measurement needs a host.

> **The release thesis** (in full in [`roadmap.md`](roadmap.md) § Active — v2.8). Two problems, one spine —
> *time to ready*. **(1)** A `--purge` + `demo-up` cycle takes ~11 min, of which **UI-tier builds are 66 % and
> image export/unpack alone 43 %**, while the box never exceeds load 4.90/8 — not a CPU problem but serialised
> I/O plus a deliberate serialisation whose stated reason no longer applies. **(2)** The Playthrough suite is
> **18/18 green while the demo still has things that don't work**: 1 of 18 proves a WRITE, none has a negative
> control, every actor is `enterprise`, `blocked`/`error` outcomes are **0**, and ant-academy / onboarding /
> org-admin / talk-to-data sit at **zero**. Of the 28 curated use cases, **16 uncovered, 12 with no milestone
> home**. Evidence: [`build-annotation.md`](releases/02.80-fast-build/evidence/build-annotation.md) ·
> [`playthrough-map.md`](releases/02.80-fast-build/evidence/playthrough-map.md).

## v2.8 shape

**M255 build-bench & host-headroom** (section, HARD barrier) ✅ BUILT · GO · **awaiting close** →
**M256 playthrough sharpening** → **M257 first-light build** (666 s → ≤ 360 s p50) → **M258 proven-live build**
(up AND self-proven, ≤ 480 s p50) — all iterative, strictly serial by the user's order (*sharpen the detector
before changing what it detects*). Execution graph, parallelism matrix and risk map:
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

## Plan review (2026-07-27) — 23 agents, 7 lenses, 35 findings, 14 verified

**APPROVE WITH REQUIRED EDITS — all applied.** No blocking defect survived; 10 refuted. Six decisions
followed, **D-v28-6 … D-v28-11**, held in full in [`roadmap.md`](roadmap.md) § "Design decisions from the
adversarial plan review": the `hostprofile` auto-planner **CUT** (no downstream consumer — it measured and
changed nothing, the very defect this release retracts) → two measured host profiles + one failing assert wired
into M257's gate · the shared-clone race **downgraded** to union-apply + a guard test · the truly-cold bench
variant **CUT** · M256's speed clause re-cut (120 s was impossible — one LLM-bound test is ~120 s of the 228 s)
· the §8.5 retraction moved to M257 · "keep it secure" given three clauses and an owner.

**Two findings survived and are written into the milestones:** the **fake-FAPI global seat** (one active seat
per stack, no cookie scoping — M256's parallel lane needs an enabler built, and `storageState` reuse does not
isolate it), and **M258's missing world contract** (the naive composition leaves a test world behind a
presenter cockpit of dead CTAs — the state M254 left `billion` in — and the first-draft gate passed anyway).

## Headline numbers (M255 close, 2026-07-28)

- **Python (rext):** **1505 pass / 2 skip / 0 fail** (1507 tests; stack-core + demo-stack + stack-injection),
  counts from JUnit XML not grepped stdout. stack-core alone **226 → 272** over M255.
- **Go (rext):** **2023** test funcs (v2.7: 2019, +4) · **0 of 6** modules failing. NB rext is not one Go
  module — each section has its own `go.mod`, so `./...` from the root fails; run them individually.
- **Flake: 0** (3 sequential full runs). **Platform-repo edits: 0. Net-new deps: 0.**

### New v2.8 baselines to beat

| | Baseline | v2.8 target |
|---|---|---|
| Cold `--purge` + `demo-up` (billion) | **666.29 s** (n=3 p50; min 658.15) | **≤ 360 s p50** (M257) · stretch 300 s |
| Playthrough suite (billion) | **228 s**, dominated by one ~120 s LLM-bound test | median per-PT **≤ 5 s** + post-coverage suite **≤ 200 s**, LLM lane separate, 0 flake ×3 (M256) |
| Mutating Playthroughs (mutate **and read back**) | **1** of 18 (17 UNCLASSIFIED, ≥1 mutates) | **≥ 5** (M256) |
| `blocked` / `error` outcomes | **0** | **≥ 1 `blocked`** (M256) |
| Curated UCs with no milestone home | **12** | **0** — landed or written verdict (M256) |
| Composed up-and-proven cycle | *does not exist* | **≤ 480 s p50**, zero standing red (M258) |

## Recently closed milestones (max 5)
- **M255 — 2026-07-28** · build-bench & host-headroom (section, HARD barrier) · **VERDICT GO** · baseline
  n=3 p50 666.29 s · 3 Fate-1 landed, 4 Fate-3 → M257, 0 escape-hatch · 0 platform edits.

## Recently shipped releases (older → roadmap.md / roadmap-legacy.md)
- **v2.7 "july jitter" — 2026-07-25** (tag `v2.7`) — re-ground + fidelity + field-hardening; M246→M254;
  prove-on-billion a–h live; **zero carry-forward**; 0 platform edits.
- **v2.6** 2026-07-23 (`v2.6`) · **v2.5** 2026-07-20 (`v2.5`).

## Standing backlog (fated destinations)
- **Consumed by v2.8:** the reserved Playthroughs futures **M206** (ai-sim mirror tier + the 3 employee
  deepening legs) and **M207** (academy coverage) are **re-fated inside M256's clause 3** — each gets a written
  verdict (named future milestone or drop) rather than a sixth consecutive re-reservation.
- **DROPPED:** DEF-M250-01 `participants_filter` (D18) · DEF-M215-03(a)/F11 (design-time) · DEF-M239-01 (v2.6).
- **Still unscheduled (vision):** DEF-M10-01 (S3/Bunny voice media — voice presence-only) · DEF-M21-01 ·
  CAVEAT-1 · M314b (platform) · **M205** residual (tier gates + ATS).

## Process flags (do NOT auto-push)
- 🚫 **`billion` — STANDING RULE (user, 2026-07-28; supersedes the dated 48 h freeze).** **Do not touch it, and
  do not even probe its status, until there is a real need for that environment.** Using it requires the
  **user's sign-off, per occasion** — and only when the work genuinely cannot be done locally. **Try locally
  first**, accepting this laptop's limits (Docker VM ~9.7 GiB vs the documented 12 GB UI-tier floor). This
  overrides every gate, protocol and default here, including `verification.md`'s prove-on-billion lineage.
  **Consequence for v2.8:** M256's suite-p50 clause, M257's whole exit gate and M258's composed cycle are all
  billion-measured and therefore **cannot fire without sign-off** — but the *code half* of each is local and
  unblocked (M256's seat-isolation enabler + mutation classification + negative controls + the onboarding ×5 /
  org-admin ×4 Playthroughs; M257's lever implementation). Only their measurement needs the host.
  **Last known state (2026-07-28, do not re-probe):** demo-1 up, 16 containers, cockpit serving the deeplink
  build at `https://demo1.anthropos.work:17700`, rext pinned `cockpit-deeplinks-v1`.
- 📌 **Provenance of every billion-measured M255 number: taken 09:59–11:37Z 2026-07-27, PRE-freeze.** Reported
  third-party activity starts **~13:11Z — 94 min after the last rep** (reps 11:11/11:26/11:37), so there is no
  overlap and nothing is contaminated (user-confirmed). Corroborated twice: three totals across two sessions
  cluster within 2 % (**658/666/672 s**), and rep-02's 881 s outlier is *causally* attributed — 206 s of its
  215 s excess in two sub-phases, matched to a reclaim evicting 7 records / 356.8 MB (contention smears, it
  does not concentrate). **On the first post-freeze campaign, re-confirm the three timing-derived claims:**
  the n=3 p50, spike (a)'s 146.8 s → 2.9 s export, spike (d)'s disk-bound attribution. **The barrier verdict
  needs no re-confirmation** — `4.84 GB → 379 MB` is bytes on disk, not a stopwatch.
- **v2.7 is merged to `main` + tagged `v2.7` LOCALLY; NOT pushed to origin** — the user runs origin publishes
  on their own cadence. **v2.5**'s and **v2.6**'s merges + tags are likewise local-only.
- rext code-of-record: the authoring copy is on `main`, and M255's tooling ships at
  **`fast-build-m255-buildbench-2`** (buildbench + both measured host profiles + the union-apply fence + the
  `Read at` anchor fence) — **pushed to origin**. The earlier `-1`/unsuffixed tags are on origin too.
- **Rung zero:** `git push --tags` is part of shipping a tool. Verify a tag is on **origin** before any
  prove-it-live step.

## ▶ RESUME HERE (paused 2026-07-27 ~14:40Z · resumed + re-synced 2026-07-28 — read this first)

**2026-07-28 re-sync (all local, no remote host touched):** everything is now **on origin** — the roadmap
branches `release/02.80-fast-build` + `m255/build-bench-host-headroom` and the `wip/v2.8-m255-paused` tags in
both repos had **never been pushed** and existed only on the laptop; they are pushed now. `main` was merged
into both branches, so the cockpit-deeplink work done during the pause is here too (the deeplink spec + the
re-pinned `latency-budget.md` citations). M255's claims were re-verified against the
code: **10/10 checklist items, every deliverable present**. **No committed work was lost.** ⚠️ Correction (2026-07-28): the M256 startup's archived-milestone sweep
deleted `work-m255/` including `conflict-preserve/`. The reverted cockpit commit is still a reachable git
object (`git show 37260f1` in rext) — but the zombie's *uncommitted* buildbench patches are gone, and with
them its two unevaluated salvage ideas: a **need-based LRU reclaim** and a **`build_cache_cap_gib` profile
field**. Neither is in any commit. They are recorded here by name so they can be re-derived if wanted; the
implementation sketch is not recoverable. **Process note:** the archived-milestone sweep destroys artifacts
that a closed milestone's own docs point to — the sweep and the "preserved under …" claim are incompatible.

> **The interlude:** while M255 was paused, a cockpit **story-deeplink** feature was built and shipped to
> `main` in both repos (rext `c755214`, tag `cockpit-deeplinks-v1` on origin; rosetta `bf3f9bc`). It is live on
> billion's demo1 cockpit and on a local `demo-2`. It is **not** part of v2.8's scope and owes M255 nothing —
> noted here only so the merge commits on this branch are not mistaken for milestone work.



**Stable resting point.** Both trees committed and clean; rext tagged **and pushed to origin**; suites green.
Nothing is half-written. **M255 is BUILT + HARDENED but deliberately NOT CLOSED** — the user halted roadmap
work for the `billion` freeze, so `/developer-kit:close-milestone` was not run.

| | |
|---|---|
| Branch | `m255/build-bench-host-headroom` @ `d983cdc` (4 commits, unmerged into `release/02.80-fast-build`) |
| rext | `ca4253c`, tag **`fast-build-m255-buildbench-3`** — on origin, rung-zero verified |
| Tests | stack-core **226 → 272**; final green **1198 pass / 1 skip** (stack-core + demo-stack) |
| Coverage | buildbench 71→74 %, union_apply_guard 90→92 %, total 78→80 % over a denominator +97 statements |

**Resume order — the first step needs no host:**
1. **`/developer-kit:close-milestone M255`** — fully local. This is the natural next action.
2. Then **M256's local half**, which the freeze does not block: the fake-FAPI seat-isolation enabler, the
   per-spec `MUTATES`/`READ-ONLY`/`UNKNOWN` classification, the negative-control mechanism, authoring the
   onboarding ×5 + org-admin ×4 Playthroughs, and the coverage verdicts. Only M256's **measurement** needs the
   host.
3. **After the freeze lifts (~2026-07-29):** re-confirm the three timing-derived claims (see the Provenance
   flag above), then M256's gate → M257 → M258.

**M255's five harden items were re-fated at its close (2026-07-28)** — "M255 harden resume" was not a named
milestone and could not survive the close. **Landed (Fate 1):** the plan-number mirror
fence (`stack-core/tests/test_baseline_mirror_fence.py`) — `billion.json`'s `gated_baseline.total_p50_s` is now
the single source, 8 prose mirrors fenced against it, RED-proven. **Fate 3 → M257** (recorded in its
`overview.md` § Inherited from the M255 close): `run_campaign` rep-body coverage · `demo_knob_guard` anchor
mutants · `_manifest_lists` silent truncation · the `laptop` profile's provisional field. Each attaches to M257
because M257 is the milestone that exercises it.

⚠️ **M257 is tighter than first drafted.** Its `re_scope_trigger` was re-derived **480 → 420 s**: at 480 it
would only have fired if L1+L2+L3 returned under 186 s, so a 60 s gate miss could never have tripped it.
666.29 → 360 needs **306.29 s** against **300–350 s** of levers, so **at L1's conservative end the big three
miss the gate** and L4/L5/L7/L8/L10 are load-bearing, not padding.

_Last updated: 2026-07-27 ~14:40Z — **M255 BUILT + HARDENED, barrier verdict GO**, PAUSED at a stable
resting point for the `billion` freeze. Both trees clean, rext tagged on origin, suites green, milestone
unmerged. Resume at ▶ RESUME HERE above._
