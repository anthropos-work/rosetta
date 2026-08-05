---
active_release: "v2.8 «fast build» — IN DEVELOPMENT (branch release/02.80-fast-build, designed 2026-07-27; adversarially plan-reviewed + revised same day). The time-to-ready release: from nothing, to live, to provably live, fast. Measure the machine and spend it deliberately (build bench + two checked-in measured host profiles + one HARD headroom assert), sharpen the Playthrough suite (faster · effective · covered), collapse the demo/dev bring-up 666 s → ≤ 360 s, then bake the Playthroughs into the bring-up so a stack comes up AND proves itself. **5** milestones M255 (HARD barrier) → M256 → **M257x** → M257 → M258, strictly serial. **M257x was INSERTED 2026-07-31** — the platform is migrating its microservices back into `app`, and both rosetta and rosetta-extensions must be re-aligned before a bring-up can be honestly measured; **M257 is PAUSED behind it**. Tooling + docs only, 0 platform-repo edits."
active_branch: "release/02.80-fast-build"
active_milestone: "M257x — platform re-alignment (iterative) — IN PROGRESS, 97 iters + 22 harden passes closed (branch m257x/platform-realignment). Find where the microservices→app consolidation actually stands, then make BOTH rosetta (corpus) and rosetta-extensions (tooling) work against platform @ origin HEAD. M257 is PAUSED behind it after 3 iters."
last_closed: "M256 — 2026-07-30"
phase: "M257x ITER LOOP, local to the new Mac (D-v28-15). **Gate 4 of 5.** Clauses 1–4 HOLD. **Clause 5 (KB-fidelity) is the only open one** and the user has ruled FOUR times it is met ONLY by a reading returning zero — never re-cut, reinterpreted or argued. **iter-96 REPAIRED by predicate; iter-97 RE-READ and it returned `N` = 20** (17 predicates), from 58 booked across 14 seats, 54 upheld by 4 independent adjudicators (**93.1 %**, against 92.1 / 93.0 / 92.7 — four adjudications inside 1.0 points, precision has NOT drifted). **N went 13 → 20 and that is reported as the result, not framed:** only **3 of 20 are recurrences** of a repaired predicate, so the repair held and the reading went deeper; iter-95 said its 13 was **a floor twice over** and it was. Chapman **N̂ ≈ 29.3**, per-pass **41 %/44 %**, union **≈68 %** → **~9 still unfound**. **The class composition INVERTED, which is the repair's signature:** the platform-drift class M257x exists to fix fell **7-of-13 → 1-of-20 (5 %)**; what dominates now is the corpus's claims about ITSELF — wrong-construct intra-corpus citations, stale currency pins (`v1.363.2`), enumerations short a member. **iter-96 repaired 51 SITES from 13 booked anchors** (adjudicators had estimated ≥8 unbooked twins; it was 38) and fenced them: `claim_twin_guard` 101 → **114** adjudicated claims, GREEN. **But 3 of 51 escaped as PARAPHRASES** — the fence matches quoted verbatim forms, so `backend.md:13` (*'still pending'*), `service_taxonomy.md` (archive dates) and `backend.md:33` (skiller stream, **left standing in the SAME FILE whose `:127` was repaired**) are live while the guard is green. **The fence is necessary and measurably not sufficient: residual ~6 %.** Pre-registration **3 of 7 HELD** (iter-95 was 6 of 6 and booked as a warning): every FAILURE was a magnitude guess, every HOLD a mechanism claim — precision, recall, class composition and induced-defect rate all landed. **Two of the 20 were CAUSED by the repair**, and one is the hazard `D-M257x-96-5` named one iter earlier (an in-file `:543` self-anchor left behind when the cross-file ones were re-pointed). **Rule 44's own '1,178 NULs' is FALSE — the file has 1** (`grep -c` counts LINES; the zsh `$'\\x00'` pattern degenerates to empty), caught by the adversarial seat in BOTH readings: the rule about lying instruments was written with one. Guard family **14 GREEN · 0 RED** over 17 (3 input-gated report *'Nothing was checked; this is not GREEN'*). **rext Python suite ran on this box for the first time — pytest was never installed — 909 pass / 1 fail, the failure PRE-EXISTING and reproduced with this run's changes reverted:** `claim_ledger.py` has **no supersession**, so iter-49's claim (itself REFUTED at iter-52) stays armed against the sentence that is now true → `DEF-M257x-iter96-ledger-supersession`. **TWO INSTRUMENTS, AND CONFLATING THEM HAS COST THIS MILESTONE REPEATEDLY:** clause 3's is the GUARD FAMILY; clause 5's is the graded READ (instrument frozen, sha `3858ec53…`, one commit ever). **STANDING SEARCH RULE, AMENDED at iter-96 because the old one was WORSE than the defect:** `git grep` at a named ref is necessary and NOT sufficient — three mechanisms hide a tracked file (12 gitignored-but-tracked text files; 2 NUL-bearing sources invisible to BOTH tools; **2 NESTED untracked repos, `app/studio` + `cms/studio`, where `git grep` at the HOST ref is a guaranteed zero**). On the `mistralai` predicate the three instruments returned **1 / 0 / 22** and the **0 was the ref-named `git grep`**. Now `platform-alignment.md` §5 rule 44 + mechanized in `anchor_construct_guard._clone_of`. **RE-SCOPE TRIGGER — read before quoting:** occurrence 3, **NOT firing** (it fired at iter-53, remedy shipped as TOK-04; occurrence 3 is 33+ clean iters later, so 'two CONSECUTIVE invalidated attempts' is false on its own words)."
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

M257x iter loop, **97 iters closed**. Gate **4 of 5**; clause 5 open — the iter-97 re-read returned
**N = 20**. See the `phase:` field above for the live detail; the reading is
[`iter-97/adjudication.md`](releases/02.80-fast-build/m257x-platform-realignment/iter-97/adjudication.md).

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
