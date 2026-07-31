---
active_release: "v2.8 «fast build» — IN DEVELOPMENT (branch release/02.80-fast-build, designed 2026-07-27; adversarially plan-reviewed + revised same day). The time-to-ready release: from nothing, to live, to provably live, fast. Measure the machine and spend it deliberately (build bench + two checked-in measured host profiles + one HARD headroom assert), sharpen the Playthrough suite (faster · effective · covered), collapse the demo/dev bring-up 666 s → ≤ 360 s, then bake the Playthroughs into the bring-up so a stack comes up AND proves itself. 4 milestones M255 (HARD barrier) → M256 → M257 → M258, strictly serial. Tooling + docs only, 0 platform-repo edits."
active_branch: "release/02.80-fast-build"
active_milestone: "M257 — first-light build (iterative) — NOT STARTED. Gate: cold --purge + demo-up, from `billion`'s 666.29 s baseline → ≤ 360 s p50 measured on **odysseus** (D-v28-14 moved the gate host; odysseus's own baseline is UNMEASURED and M257 owes it), autoverify green / 0 warnings. re_scope_trigger re-derived 480 → 420 s."
last_closed: "M256 — 2026-07-30"
phase: "Between milestones. M256 closed-on-gate + merged into release/02.80-fast-build. Both repos clean + pushed; rext tag fast-build-m256-harden-final on origin. demo-2 up (16 containers, pt-world re-seeded, drifted cockpit fixture sha 99e2f315 restored). Next: /developer-kit:work-milestone --milestone=M257."
last_updated: "2026-07-30"
---

# State

**v2.8 "fast build" IN DEVELOPMENT** — designed 2026-07-27, branch `release/02.80-fast-build` cut from `main`.
**2 of 4 milestones closed.** Full narratives live in [`roadmap.md`](roadmap.md) § M255 / § M256 — not here.

## Hosts (D-v28-14, 2026-07-31)

- 🔒 **`billion` — THE DEMO MACHINE.** Deploying a final working demo ONLY. **Not available for development
  or testing.** Do not use it to build, measure, or iterate. (This replaces the old per-occasion sign-off
  rule and is stricter, not looser.)
- ✅ **`odysseus` — the dev/test host.** `devops@100.110.67.14` / `odysseus.taildc510.ts.net` (Tailscale).
  User-signed-off for this milestone **and later ones**, including moving to a nearer-production machine
  when a milestone needs it. Measured 2026-07-31: **8 cores / 7 GiB / x86_64 Linux 6.8, 189 G free,
  Docker 29.6.2, NO Go** (the rext tooling is Go — install it).
- 💻 **laptop** — 10 cores / 16 GiB. Free as of 2026-07-31 (demo-2 torn down, 244 GiB free, 0 containers).
  M255's headroom model **refuses** it for two concurrent Next.js build lanes.

**`666.29 s` is BILLION's baseline and does not transfer.** M257 must measure odysseus's own (n >= 3) and
check in `odysseus.json` before pricing any lever against it.


## Active milestone

**M257 — first-light build** (`iterative`, not started). Collapse the cold `--purge` + `demo-up` cycle
from **`billion`'s 666.29 s → ≤ 360 s p50 on `odysseus`**. Levers L1–L10 are ranked by measured seconds in its `overview.md`; **L1 is already
proven real** (M255: hiring image 4.84 GB → 379 MB, export leg 146.8 s → 2.9 s).

⚠️ **Tighter than first drafted.** Its `re_scope_trigger` was re-derived **480 → 420 s**: at 480 it could only
fire if L1+L2+L3 returned under 186 s, so a 60 s gate miss could never have tripped it. 666.29 → 360 needs
**306.29 s** against **300–350 s** of levers, so **at L1's conservative end the big three miss the gate** and
L4/L5/L7/L8/L10 are load-bearing, not padding.

**It inherits 6 items** — 4 from the M255 close (`run_campaign` rep-body coverage · `demo_knob_guard` anchor
mutants · `_manifest_lists` truncation · the `laptop` provisional field) and **2 from the M256 close, both
gate-relevant**: `FIX-M256-demo2-service-self-termination` (two services self-terminate `Exited 0` while
`docker ps` shows 14/16 "Up" and every grid renders 20 content-free rows — M257's gate would go **green on a
half-dead stack**) and `FIX-M256-autoverify-fapi-libressl` (warns *"NOBODY CAN LOG IN"* about a working stack;
M257's gate counts autoverify warnings). Both in its `overview.md`.

## Phase

Between milestones. M256 is closed and merged; nothing is half-written.

## v2.8 shape

**M255 build-bench & host-headroom** (section, HARD barrier) ✅ **done 2026-07-28, VERDICT GO** →
**M256 playthrough sharpening** ✅ **done 2026-07-30, `closed-on-gate`** → **M257 first-light build**
(`billion` 666 s → ≤ 360 s p50) → **M258 proven-live build** (up AND self-proven, ≤ 480 s p50). Strictly serial by the
user's order — *sharpen the detector before changing what it detects*.

## Binding user decisions (2026-07-27, + later)

- **D-v28-1** codename **"fast build"** · **D-v28-2** M255 stays a `section` barrier.
- **D-v28-3** — **batch-gate semantics: no accumulating red.** A run drives the **full batch to completion**;
  at batch end a non-empty red set **escalates for renegotiation**. **Zero standing red** is the invariant.
- **D-v28-4** coverage scope: land onboarding (5) + org-admin (4); every remaining uncovered curated UC gets a
  **written verdict**. **D-v28-5** the cockpit logout defect fixed in M256, deliberately with no Playthrough.
- **D-v28-6 … D-v28-11** from the adversarial plan review (23 agents, 7 lenses, 35 findings, 14 verified,
  10 refuted) — held in full in [`roadmap.md`](roadmap.md).
- **D-v28-12 / D-v28-13** — M256's clause 1 re-cut **twice**: relative for a local host, then **per-leg rather
  than per-suite**, once the threshold was measured to sit **inside its own 2.04× noise floor**. The standing
  rule it produced: ***never gate a relative statistic without first measuring its variance.***
- **HARDEN-CAP-ACCEPTED-D105** (2026-07-30) — M256's final harden accepted un-stabilized, residuals enumerated.
- **M256 close ratifications** (standing delegation, overrule preserved to release close): **D103**, **D104**,
  and the **iter-31/32 deviation**. See `m256…/decisions.md`.

## Headline numbers (M256 close, 2026-07-30)

- **Python (rext):** **1,723 pass / 2 skip / 0 fail** over four suites (stack-core 287 · demo-stack 1000 ·
  stack-verify 171 · stack-injection 267), counts from **JUnit XML, never grepped stdout**, one invocation per
  suite with rc captured into a variable. Same-scope vs M255's three-suite 1,505: **1,552 (+47)**.
- **Go (rext):** **2,130** test funcs (M255: 2,023, **+107**) · **0 of 6** modules failing, 58 packages ok.
  NB rext is **not one Go module** — `go test ./...` from the root fails by design; run each section.
- **Playwright:** **209 tests in 43 files** (was 204/42); unit **174** (was 169).
- **Playthroughs: 30 live + 1 verdicted TODO** · mutating **12** · negative controls **28 of 30** ·
  `blocked` **1** · written verdicts **31/31, 0 `unimplementable`** · `ptvalidate` VALID rc 0.
- **Flake: 0** (3 consecutive cold reset-to-seed, rc `0/0/0`). **Platform-repo edits: 0. Net-new deps: 0.**

### v2.8 baselines to beat

| | Baseline | v2.8 target |
|---|---|---|
| Cold `--purge` + `demo-up` (billion) | **666.29 s** (n=3 p50; min 658.15) | **≤ 360 s p50** (M257) · stretch 300 |
| Composed up-and-proven cycle | *does not exist* | **≤ 480 s p50**, zero standing red (M258) |
| Playthrough suite, ABSOLUTE on billion | **228 s** (dominated by one LLM-bound test) | re-measure = **reporting only**, M258 |

## Recently closed milestones (max 5)
- **M256 — 2026-07-30** · playthrough sharpening (iterative) · **`closed-on-gate`**, 32 iters / 3 harden passes
  · 18 → **30 live Playthroughs**, mutating 1 → 12, controls 0 → 28/30 · **all three gate clauses proved
  unmeetable as first authored** · ~43 checks that reported success without checking · 4 Fate-1, 10 Fate-3
  (6 → M258, 2 → M257), 1 drop, **3 awaiting the user's signature**, 0 escape-hatch · 0 platform edits.
- **M255 — 2026-07-28** · build-bench & host-headroom (section, HARD barrier) · **VERDICT GO** · baseline
  n=3 p50 666.29 s on `billion` · 3 Fate-1, 4 Fate-3 → M257, 0 escape-hatch · 0 platform edits.

## Recently shipped releases (older → roadmap.md / roadmap-legacy.md)
- **v2.7 "july jitter" — 2026-07-25** (tag `v2.7`) — re-ground + fidelity + field-hardening; M246→M254;
  prove-on-billion a–h live; **zero carry-forward**; 0 platform edits.
- **v2.6** 2026-07-23 (`v2.6`) · **v2.5** 2026-07-20 (`v2.5`).

## Standing backlog (fated destinations)
- **Consumed by v2.8:** the reserved Playthrough futures **M206** / **M207** were re-fated inside M256's
  clause 3 — each has a **written verdict** rather than a sixth consecutive re-reservation.
- **Awaiting the user's signature** (M256 close, roadmap calls not routing ones): `PERF-M256-parallel-lane`
  (needs a cookie-scoped Clerkenstein registry or one fake-FAPI per worker — a real build; no v2.8 gate needs
  it) · `PT-M257-self-evaluation` (re-homing an M206 reservation) · `PT-M257-talk-to-data` (blocked on
  `ask_*` migrations **plus live Bedrock credentials**).
- **Platform defects** (Rosetta cannot fix; zero platform edits binding) → the net-new
  [`platform-defect-register.md`](platform-defect-register.md), 4 entries with `file:line`.
- **DROPPED:** DEF-M250-01 · DEF-M215-03(a)/F11 · DEF-M239-01 · `PT-M256-resume-fixture-pair` (premise dissolved).
- **Still unscheduled (vision):** DEF-M10-01 (S3/Bunny voice media) · DEF-M21-01 · CAVEAT-1 · M314b (platform)
  · **M205** residual.

## Process flags (do NOT auto-push)

- 📌 **Provenance of every billion-measured M255 number: taken 09:59–11:37Z 2026-07-27, PRE-freeze**, with no
  overlap with third-party activity (user-confirmed; three totals across two sessions cluster within 2 % —
  658/666/672 s). **On the first post-freeze campaign, re-confirm three timing-derived claims:** the n=3 p50,
  spike (a)'s 146.8 → 2.9 s export, spike (d)'s disk-bound attribution. The **barrier verdict needs no
  re-confirmation** — 4.84 GB → 379 MB is bytes on disk, not a stopwatch.
- ⚠️ **The stack's PINNED `stackseed` can be older than the authoring copy.** M256 added three `Persona`
  fields, so a `--reset` with the stack's binary **truncates the world and then fails to re-seed**, leaving it
  EMPTY. Shadow the authoring build on `PATH` for any reset. Cost the M256 close one run before it was found.
- **v2.7 is merged to `main` + tagged `v2.7` LOCALLY; NOT pushed to origin** — the user runs origin publishes
  on their own cadence. **v2.5** and **v2.6** are likewise local-only.
- rext code-of-record: authoring copy on `main`; M256's tooling ships at **`fast-build-m256-close`** @ `ce345e1`
  — **on origin, rung-zero verified** — which is `fast-build-m256-harden-final` plus the close's **3** commits.
- **Rung zero:** `git push --tags` is part of shipping a tool. Verify a tag is on **origin** before any
  prove-it-live step.
- **`run-playthroughs.sh` is BINDING since M256** — a full run exits non-zero when ptreport's gate is unmet
  (advisory on a scoped run). Anything that ran the suite and trusted a zero exit is now genuinely gated.

_Last updated 2026-07-30 — M256 closed-on-gate and merged. Stable resting point: both trees clean, rext tagged
on origin, suites green, demo-2 up with its drifted cockpit fixture restored (sha 99e2f315)._
