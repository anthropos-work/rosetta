---
active_release: "v2.8 «fast build» — IN DEVELOPMENT (branch release/02.80-fast-build, designed 2026-07-27). Time-to-ready: from nothing, to live, to provably live, fast. **5** milestones M255 → M256 → **M257x** → M257 → M258, strictly serial; M257x was INSERTED 2026-07-31 and **M257 is PAUSED behind it**. Tooling + docs only, 0 platform edits. Detail: roadmap.md § v2.8."
active_branch: "release/02.80-fast-build"
active_milestone: "M257x — platform re-alignment (iterative) — **AWAITING USER SCOPE DECISION**, 119 iters + 25 harden passes closed (branch m257x/platform-realignment). Re-align BOTH rosetta (corpus) and rosetta-extensions (tooling) to platform @ origin HEAD. Gate 4 of 5; clause 5 the only one open. `TOK-08` REFUTED at iter-119. M257 PAUSED behind it after 3 iters."
last_closed: "M256 — 2026-07-30"
phase: "M257x AWAITING A **USER SCOPE DECISION**. **Gate 4 of 5**; clause 5 open and met ONLY by a reading that returns zero — four user rulings, never re-cut. **TWO strategies have now been refuted by their own pre-registered arithmetic**: `TOK-07` (repair-and-read) at iter-116 (`P = 37` vs `P >= 15`) and `TOK-08` (the USER's enumerate-then-read re-scope) at iter-119 (`P = 22` vs `P >= 19`). Per `TOK-08`'s own sealed rule **NO successor strategy is authored — there is no TOK-09.** iter-119 also took the milestone's FIRST test-retest reading: it re-found only **13 of iter-116's 37 (35.1 %)** over a corpus 5 in-place lines away, so **`P` fell while the floor ROSE**. Chapman retired — floors only (**>= 46 at `194361e4`**). Detail: milestone `progress.md` + `decisions.md` § TOK-08 OUTCOME. POINTER field per `context.md` § state.md contract; do not grow it back."
last_updated: "2026-08-07"
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

M257x iter loop, **108 iters closed**. Gate **4 of 5**.

- **Clauses 1–2 — CLOSED** by the concurrent lane at platform `0c91421`. **Clause 2 is MET WITH DISCLOSURE
  and the disclosure travels with it forever: a freshly built stack failed the first full run in 2 of 2
  attempts, so it is never recorded as a clean pass.** Every cycle timing, the reflog freeze-proof and the
  suite verdicts live in the dossier — see the pointer list below.
- **Clauses 3–4 — hold**, asserted by fences that are watched going RED, not by inspection.
- **Clause 5 — the only open one**, met ONLY by a reading that returns zero. iter-103's 14-seat double
  reading over the repaired tree returned **`N = 33`** against a **pre-registered, pre-sealed** rule →
  **the burn-down leg does not reach the residual.** By **predicate** the pool did not move — **22 then,
  22 now**; by anchor 24 → 33. Repair efficacy is nonetheless **confirmed** (21 of 22 predicates closed).
  `N` stayed up because two inflows feed the residual that repair does not touch: **clone advance** (61 %)
  and **the repair's own induction** (21 %). **Inflow ≈ outflow; running the loop faster does not close it.**

**Active strategy — `TOK-06: fence the inflows before repairing again`** (iter-104, a **deliberate**,
non-terminating tok). It changes the ORDER of the loop, not the instrument and not the unit of repair:
**(0) guard-tree provenance → (1) the drift fence → (2) the induction checks → (3) repair the 33 →
(4) read LAST.** **Steps 0–3 LANDED in iters 105–108** — three net-new rext fence modules, each with a
mutation control and an anti-vacuity control that can fire, then **the union paid BY PREDICATE** (iter-108:
22 predicates / 23 files; machine reach **46/46 = 100 % of the upheld union**, raw 46/47 — the single miss a
**REJECTED `wrong-tree`** finding, iter-102's residue result reproduced). **Both fences fired ON that repair
and were cleared.** **Only step 4 remains: the read**, deliberately unstarted.

**Chapman is RETIRED for this milestone.** Its independence assumption measured **17 %** then **61 %** on
one byte-identical instrument, so every point estimate derived from it is unusable. **Only the floor
survives: ≥ 24 at `8f04d3a`, ≥ 33 at `e6aed2e`.** **A zero reading is not near.**

**The live detail lives in the milestone's own docs — `state.md` is the index, not the narrative**
(`context.md` § state.md contract):

- [`m257x…/progress.md`](releases/02.80-fast-build/m257x-platform-realignment/progress.md) — authoritative
  live status, every iter's findings, the carried items
- [`m257x…/decisions.md`](releases/02.80-fast-build/m257x-platform-realignment/decisions.md) — `D-N` +
  the `TOK-01…06` strategy chain, incl. TOK-06 in full
- [`gate-clauses-1-2/README.md`](releases/02.80-fast-build/m257x-platform-realignment/gate-clauses-1-2/README.md)
  — **the owner of every clause 1 & 2 number**: the five cold-cycle timings, the reflog freeze-proof, and
  the disclosed first-run failure
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

**`D-v28-1` … `D-v28-15` are held in full in [`roadmap.md`](roadmap.md)** — including the two that produced
standing rules: **`D-v28-3`** (batch-gate semantics — **zero standing red** is the invariant) and
**`D-v28-12`/`D-v28-13`** (***never gate a relative statistic without first measuring its variance***).
M256's close ratifications (`D103`, `D104`, the iter-31/32 deviation) are in that file's § M256 closure;
`HARDEN-CAP-ACCEPTED-D105` is in [`m256…/decisions.md`](releases/02.80-fast-build/m256-playthrough-sharpening/decisions.md).

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
