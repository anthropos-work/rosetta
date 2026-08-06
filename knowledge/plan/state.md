---
active_release: "v2.8 «fast build» — IN DEVELOPMENT (branch release/02.80-fast-build, designed 2026-07-27). Time-to-ready: from nothing, to live, to provably live, fast. **5** milestones M255 → M256 → **M257x** → M257 → M258, strictly serial; M257x was INSERTED 2026-07-31 and **M257 is PAUSED behind it**. Tooling + docs only, 0 platform edits. Detail: roadmap.md § v2.8."
active_branch: "release/02.80-fast-build"
active_milestone: "M257x — platform re-alignment (iterative) — IN PROGRESS, 103 iters + 22 harden passes closed (branch m257x/platform-realignment). Re-align BOTH rosetta (corpus) and rosetta-extensions (tooling) to platform @ origin HEAD. Gate 4 of 5: clauses 1–2 CLOSED by the concurrent lane at 0c91421 (clause 2 is MET WITH DISCLOSURE — a fresh stack failed the first full run 29/1 in 2 of 2, never a clean pass), clauses 3–4 hold, CLAUSE 5 is the only open one and is met ONLY by a reading that returns zero. iter-103's 14-seat double reading returned N = 33 against a pre-sealed rule whose ≥23 branch reads THE BURN-DOWN LEG DOES NOT REACH THE RESIDUAL — 22 predicates then, 22 now. M257 PAUSED behind it after 3 iters."
last_closed: "M256 — 2026-07-30"
phase: "M257x ITER LOOP, local to the new Mac (D-v28-15). **Gate 4 of 5.** Clauses 1–2 were closed by the CONCURRENT LANE at platform `0c91421` — clause 1 on five consecutive cold cycles (green:true / 0 warnings, 621/402/370/370/371 s, refs proven from the reflog); **clause 2 is MET WITH DISCLOSURE and the disclosure is part of the claim: a freshly built stack failed the first full run 29/1 in 2 of 2 attempts, so it is NEVER recorded as a clean pass.** Clauses 3–4 hold. **CLAUSE 5 is the only open one**, met ONLY by a reading that returns zero (ruled FOUR times; never re-cut, narrowed or argued). **iter-103 returned `N = 33`** on a 14-seat double reading over the repaired tree, against a rule SEALED before any seat was dealt (`≤16` works · `17–22` ambiguous · **`≥23` DOES NOT REACH**) → **THE BURN-DOWN LEG DOES NOT REACH THE RESIDUAL.** **By predicate the pool did not move: 22 then, 22 now**; by anchor 24 → 33. Repair efficacy is nonetheless CONFIRMED (21 of iter-101's 22 predicates closed). `N` held up because two inflows feed the residual that repair does not touch: **clone advance** (61 % of `N` is drift — version literals, `go.mod` pins, line offsets — which NO guard fences) and **the repair's own induction** (7/33 anchors in prose iter-102 wrote, incl. a false canonical sentence multiplied to 5 sites). **Inflow ≈ outflow: running the loop faster does not close clause 5. CHAPMAN IS RETIRED** — its independence assumption measured 17 % then 61 % on one byte-identical instrument, so `N̂ ≈ 103` and `≈ 35` are both unusable; only the FLOOR survives (≥24 at `8f04d3a`, ≥33 at `e6aed2e`). **Next action: a repair iter that pays `FIX-M257x-iter103-read-union` AND lands `FIX-M257x-iter103-drift-fence-gap`, because repairing 20 drift anchors without a fence re-arms the class at the next clone advance.** **Live detail lives in the milestone `progress.md`. This field is a POINTER, per `context.md` § state.md contract; do not grow it back."
last_updated: "2026-08-06"
---

# State

**v2.8 "fast build" IN DEVELOPMENT** — designed 2026-07-27, branch `release/02.80-fast-build` cut from `main`.
**2 of 4 milestones closed.** Full narratives live in [`roadmap.md`](roadmap.md) § M255 / § M256 — not here.

## Hosts (D-v28-15, 2026-07-31 — supersedes D-v28-14)

- 🎬 **`billion` — the OFFICIAL host.** Demo deployment only. Not for development or testing.
- 💻 **dev/test = LOCAL to the new Mac.** The old laptop and **`odysseus` are both retired** from this project.
- ⚠️ **Host-class mismatch, and it blocks M257:** billion is **x86_64/containerd**, a Mac is **arm64/overlay2**.
  M255 measured **4.84 GB vs 2.88 GB** for the same Dockerfile — **the Mac pays no unpack leg**, which is exactly
  what M257's L1 (~200–250 s) optimises. **M257's speed gate is un-measurable on the sanctioned hosts as
  written.** M257x is largely unaffected: its gate is correctness, not seconds.
- 🔧 **New-Mac bootstrap:** `.agentspace/rext.tag` is **git-ignored** → a fresh clone has no pin, and a mismatch
  is **FATAL** in `ensure-clones.sh:94-101`. Create it deliberately. (On the old box the SoT was 63 commits
  behind `main`, so `/demo-up` aborted there.)


## Active milestone

**M257x — platform re-alignment** (`iterative`, **IN PROGRESS**). Find how far the microservices→`app`
consolidation has got, write it where it cannot rot, and make **both** repos work against the platform as it is.
**Third occurrence of one class** (v2.1 skiller · v2.7 skillpath · now jobsimulation), each re-derived from
scratch — so `corpus/ops/platform-alignment.md` is a deliverable, not a formality.

**M257 — first-light build: PAUSED** behind it. Banked, not to be redone: odysseus provisioned (rc=0, 16/16),
both gate-honesty instruments landed with mutation-proven controls, B1+B2 fixed, mirror fence parameterised by
host. Still owed: the odysseus baseline, and `INVESTIGATE-M257-load1-48` — peak `load1` **48.7** vs HEADROOM
clause 1's limit of **6**.

## Phase

M257x iter loop, **103 iters closed**. Gate **4 of 5**. Clauses **1 and 2 were closed by the concurrent
lane at platform `0c91421`** — clause 1 on five consecutive cold cycles (`green:true` / 0 warnings,
621/402/370/370/371 s, refs proven from the reflog); clause 2 **MET WITH DISCLOSURE**
(`{passing:30, failing:0, unimplementable:0, unimplemented:1}`, `binding:true, scoped:false`, two stacks) —
**and the disclosure travels with it forever: a freshly built stack failed the first full run 29/1 in 2 of 2
attempts. It is never recorded as a clean pass.** Clauses 3 and 4 hold. **Clause 5 is the only open one.**

**iter-103 answered the burn-down question and the answer is no.** A fresh 14-seat double reading over the
repaired tree returned **`N = 33`** against a **pre-registered, pre-sealed** rule (`≤16` works · `17–22`
ambiguous · **`≥23` DOES NOT REACH**). **The burn-down leg does not reach the residual.**
By **predicate** the pool did not move at all — **22 then, 22 now**; by anchor it rose **24 → 33**. Repair
efficacy is nonetheless **confirmed**: 21 of iter-101's 22 predicates are closed (band #3 held at 1). `N`
stayed up because two inflows feed the residual that repair does not touch — **clone advance** (61 % of `N`
is drift no guard fences: version literals, `go.mod` pins, line offsets) and **the repair's own induction**
(7 of 33 anchors sit in prose iter-102 wrote, incl. a false canonical sentence multiplied to 5 sites).
**Inflow ≈ outflow; running the loop faster does not close clause 5.**

**Chapman is RETIRED for this milestone.** Its independence assumption has now been measured at both
extremes on one byte-identical instrument — `m`/union **17 %** (iter-101) then **61 %** (iter-103) — so
`N̂ ≈ 103` and `N̂ ≈ 35` are both unusable. **Only the floor survives: ≥ 24 at `8f04d3a`, ≥ 33 at `e6aed2e`.**
The series 16.7 → 29.4 → 45.2 → ~103 stays four **corrections to an underestimate**, not four measurements
of growth — but stop quoting a point estimate from it. **A zero reading is not near.**

**The live detail lives in the milestone's own docs — `state.md` is the index, not the narrative**
(`context.md` § state.md contract):

- [`m257x…/progress.md`](releases/02.80-fast-build/m257x-platform-realignment/progress.md) — authoritative
  live status, every iter's findings, the carried items
- [`iter-103/adjudication.md`](releases/02.80-fast-build/m257x-platform-realignment/iter-103/adjudication.md) — **the current reading** (`N = 33`; the burn-down verdict, the composition finding, the Chapman retirement)
- [`iter-101/adjudication.md`](releases/02.80-fast-build/m257x-platform-realignment/iter-101/adjudication.md) — the replicate, whose 24 iter-103 measures overlap against
- [`iter-99/adjudication.md`](releases/02.80-fast-build/m257x-platform-realignment/iter-99/adjudication.md) — the reading before that
- [`iter-98/discovery-pool.md`](releases/02.80-fast-build/m257x-platform-realignment/iter-98/discovery-pool.md) — the pool measurement

## Standing rules (outlive the milestone — do NOT move these into `phase:`)

- **Rule 44 — no single search tool is safe.** Gitignored-but-tracked files, NUL-bearing source, and nested
  untracked repos each blind a *different* instrument. Use the mechanized per-tree path
  (`anchor_construct_guard._clone_of`), never a hand rule — **rule 44's own worked recipe returns 2 where it
  publishes 22.** Three mechanisms, no single tool.
- **TWO INSTRUMENTS, and conflating them has cost this milestone repeatedly.** Clause 3's instrument is the
  **guard family**; clause 5's is the **graded read** (frozen, sha `3858ec53…`, one commit ever). A guard
  going green says nothing about clause 5, and a reading says nothing about clause 3.
- **A guard's green is only as strong as its own load-bearing word.** `anchor_construct_guard` reported
  *"every **resolvable** anchor names a construct"* and was GREEN while ≥7 upheld findings resolved to the
  wrong construct. When a guard qualifies its own claim, the qualifier is the blind spot — read it first.
- **RE-SCOPE TRIGGER: occurrence 3, NOT firing.** It fired at iter-53 and the remedy shipped as TOK-04;
  35+ clean iters since, so *"two CONSECUTIVE invalidated attempts"* is false on its own words.

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
