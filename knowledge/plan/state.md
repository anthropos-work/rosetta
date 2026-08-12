---
active_release: "v2.8 «fast build» — IN DEVELOPMENT (branch release/02.80-fast-build, designed 2026-07-27). Time-to-ready: from nothing, to live, to provably live, fast. **5** milestones M255 → M256 → M257x → M257 → **M258**, strictly serial. **4 of 5 closed.** Tooling + docs only, 0 platform edits. Detail: roadmap.md § v2.8."
active_branch: "release/02.80-fast-build"
active_milestone: "M258 «proven-live build» (`iterative`, the closer) — NOT STARTED. Gate: one cold command brings the stack up AND drives the full Playthrough batch to completion, zero standing red, total p50 ≤ 480 s over 3 cold cycles."
last_closed: "M257 — 2026-08-12"
phase: "Between milestones. **M257 closed 2026-08-12 `closed-on-gate` — the gate FIRED**, p50 286.99 s on `macmini` vs a 360 s gate (under the 300 s stretch), every clause green on all 3 reps including both falsifiable ones. Do not read it like M257x, which closed by ruling. One lever landed of eight priced. Next: M258, the closer — and it inherits from BOTH: M257x's 11 items / 5 clusters + 1 block fate, and M257's 15 Fate-3 items led by `LEVER-M257-L5-setdress` (`set_dress` is now the largest phase at 28.6 %)."
last_updated: "2026-08-12"
---

# State

**v2.8 "fast build" IN DEVELOPMENT** — designed 2026-07-27, branch `release/02.80-fast-build` cut from `main`.
**3 of 5 milestones closed.** Full narratives live in [`roadmap.md`](roadmap.md) § M255 / § M256 / § M257x — not here.

## Hosts (D-v28-15, 2026-07-31 — supersedes D-v28-14)

- 🎬 **`billion` — the OFFICIAL host.** Demo deployment only. Not for development or testing.
- 💻 **dev/test = LOCAL to the new Mac.** The old laptop and **`odysseus` are both retired** from this project.
- ✅ **~~Host-class mismatch blocks M257~~ — RETRACTED, and M257 then MET its gate here.** It read *"a Mac
  is arm64/**overlay2** … pays no unpack leg … the gate is **un-measurable**."* **False for this machine**:
  it runs the **containerd image store** and pays a size-proportional unpack leg (**56.6 s export + 19.3 s
  unpack** on the real 4.12 GB hiring image). The generalisation came from the **retired M1 Pro laptop** —
  *"a Mac"* is a class; the fact that mattered is a per-machine Docker Desktop toggle. **Never re-derive
  this from `docker info`**, which prints `Storage Driver: overlayfs` here and is what made the wrong
  reading look right — **grade it with a probe.** L1 was worth **−141.63 s**; the cycle closed at **p50
  286.99 s**. Still true: **4.84 vs 4.12 GB** for the identical Dockerfile, and **seconds here do not
  transfer to billion**, whose post-L1 cycle is **unmeasured**.
- 🔧 **New-Mac bootstrap:** `.agentspace/rext.tag` is **git-ignored** → a fresh clone has no pin, and a mismatch
  is **FATAL** in `ensure-clones.sh:94-101`. Create it deliberately. (On the old box the SoT was 63 commits
  behind `main`, so `/demo-up` aborted there.)


## Next up

**M258 — proven-live build** (`iterative`, **the closer**). Gate: one cold command brings the stack up **and**
drives the full Playthrough batch to completion with **zero standing red**, at **total p50 ≤ 480 s** over 3
consecutive cold reset-to-seed cycles, 0 platform-repo edits, **and the stack left in a presenter-usable
world** (the world contract — decide (a) pt-world-native vs (b) restore-after at iter-01).

**480 s is a sum of two ceilings** (360 + 200), reachable *"only if M257 spends part of its unspent
levers."* **M257 landed at 286.99 s on ONE lever**, so that reserve is real — the largest piece being
`LEVER-M257-L5-setdress` (**`set_dress` is now the largest phase at 82.04 s = 28.6 %**, priced ~30–50 s and
ranked fifth).

**M258 inherits from BOTH, and both now appear in its own `overview.md`** — which was the gap: M257x's
routing had **zero** mentions there until this close (the `BIND_HOST` failure, from the file documenting
it). From **M257x**: 11 items / 5 clusters + 1 block fate
([`carry-forward.md`](releases/02.80-fast-build/m257x-platform-realignment/carry-forward.md)); **cluster 4 is
half-discharged** — what survives is that **`buildbench` asserts no elapsed-time threshold** and M258's gate
is a p50 number. From **M257**: 15 Fate-3 items, **0 escape-hatch**, led by L5 and
`FIX-M257-dockerignore-env-pattern-unpaired`, whose tidy one-line fix **bakes the real Clerk key**.

## Phase

Between milestones. **M257 closed 2026-08-12 `closed-on-gate` — the gate FIRED on its own terms.** p50
**286.99 s** on `macmini` (n=3) against a **360 s** gate and **under the 300 s stretch**, with `autoverify
green:true / 0 warnings`, **HEADROOM OK 3/3** and **ISOLATION OK 3/3**, identity `match` ×3, 0 platform-repo
edits. **Both falsifiable clauses actually fired earlier in the milestone**, and the headline was re-graded
twice under code that changed after it was taken. ⚠️ **Do not read this close like M257x's**, which closed
`closed-incomplete` **by user ruling** one day earlier with **clause 5 never met**. Full narrative:
[`roadmap.md`](roadmap.md) § M257.

⚠️ **Open safety item, routed to M258 and still not closed:** a demo reached the **production** S3 bucket and
only an **IAM policy on an account we do not control** refused it. The containment is proven by a unit test
on the emitter and **on no running stack**; both currently-running stacks still carry the pointer, and the
dev-side strip is demo-only. Owner: `corpus/ops/safety.md` + M257x `carry-forward.md` cluster 1.

## Standing rules (outlive the milestone — do NOT move these into `phase:`)

Each rule's derivation lives at its owner; these are the headlines a reader must not miss.
**`§5`** = [`corpus/ops/platform-alignment.md`](../../corpus/ops/platform-alignment.md) § 5, which owns the
numbered rules in full.

- **§5 rule 44 — no single search tool is safe.** Gitignored-but-tracked files, NUL-bearing source and
  nested untracked repos each blind a *different* instrument. Use the mechanized per-tree path
  (`anchor_construct_guard._clone_of`), never a hand rule — **rule 44's own worked recipe returns 2 where
  it publishes 22.**
- **TWO INSTRUMENTS, and conflating them cost M257x repeatedly.** The guard family and the graded read
  measure different things: a guard going green says nothing about corpus sentence-level fidelity, and a
  reading says nothing about the fences. (M257x's clause 5 — the graded read — was placed **out of scope
  by the user**, `TOK-09`. It was never met. Do not resurrect it as a claim of cleanliness.)
- **A guard's green is only as strong as its own load-bearing word** — when a guard qualifies its own
  claim, the qualifier IS the blind spot. `anchor_construct_guard` said *"every **resolvable** anchor"*
  while ≥7 upheld findings resolved to the wrong construct; it now PRINTS its floor (`KNOWN_WEAKNESS`),
  widening measured and declined (`D-M257x-121-4`). **Every count of that class is a floor.**
- **Read the SUBSTRATE before believing a defect** — a checkout behind its own fetched `origin/main` does
  not merely fail to confirm a claim, **it manufactures evidence against a true one** (6 of 13 clones
  stale; 4 adjudicators booked a true claim as contradicted). **And the mirror:** reading at HEAD instead
  of the census ref **inflated** the re-pin backlog 3.1× (iter-126). Derivation: `D-M257x-122-4`.
- **NEVER quote this milestone's read-derived error rate as corpus-wide**, and quote its two governing
  measurements together or neither: test–retest recall **~35 %** (iter-119), and the published rate is a
  **hunted-sample artifact — 0.70 % over 427 exhaustive adjudications vs ≥ 13.3 % hunted**, over-stating
  the population **~19×** (`D-M257x-122-3`). The corpus publishes no error rate at all; keep it that way.
- **§5 rule 53 — every mutation asserts it APPLIED** before its result is interpreted, and is a control
  only for the clause it can **isolate**. Three of harden pass 26's silently failed to apply, each reading
  as *"the controls survive."*
- **§5 rule 51 — state the invocation AND the expected wall time with every whole-suite count.**
- **§5 rule 54 — a correction that reaches ONE cell is not a correction.** It must reach every site
  publishing the retracted claim: the router (24 sites), `db-backup` (3), `cms`'s *"assert neither"* (5).
- **RE-SCOPE TRIGGER: occurrence 3, NOT firing.** It fired at iter-53 and the remedy shipped as TOK-04;
  35+ clean iters since, so *"two CONSECUTIVE invalidated attempts"* is false on its own words.

## v2.8 shape

**M255 build-bench & host-headroom** (section, HARD barrier) ✅ **done 2026-07-28, VERDICT GO** →
**M256 playthrough sharpening** ✅ **done 2026-07-30, `closed-on-gate`** → **M257x platform re-alignment** ✅ **done 2026-08-11, `closed-incomplete` (user ruling, not gate)** → **M257 first-light build** ✅ **done 2026-08-12, `closed-on-gate`** (`macmini` 449.51 → **286.99 s** p50, gate 360) → **M258 proven-live build** (up AND self-proven, ≤ 480 s p50). Strictly serial by the
user's order — *sharpen the detector before changing what it detects*.

## Binding user decisions (2026-07-27, + later)

**`D-v28-1` … `D-v28-15` are held in full in [`roadmap.md`](roadmap.md)** — including the two that produced
standing rules: **`D-v28-3`** (batch-gate semantics — **zero standing red** is the invariant) and
**`D-v28-12`/`D-v28-13`** (***never gate a relative statistic without first measuring its variance***).
M256's close ratifications (`D103`, `D104`, the iter-31/32 deviation) are in that file's § M256 closure;
`HARDEN-CAP-ACCEPTED-D105` is in [`m256…/decisions.md`](releases/02.80-fast-build/m256-playthrough-sharpening/decisions.md).

## Recently closed milestones (max 5)

_Trimmed to the last 3 days per the state.md contract; older entries live in `roadmap.md`'s `### M{N}` blocks._

- **M257 — 2026-08-12** · first-light build (iterative) · **`closed-on-gate` — THE GATE FIRED.** p50
  **286.99 s** on `macmini` vs 360 s (under the 300 s stretch), n=3, every clause green on all three reps
  **including both falsifiable ones**. 9 iters (7 tiks + 2 toks); **one lever of eight priced (L1) cleared it
  alone** (UI tier −141.63 s; images 4.04 GB → 417 MB and 3.94 GB → 380 MB). The close also landed the §8.5
  corpus retraction with achieved numbers + the grep gate that had never existed, and found **three
  fail-opens in the gate instrument itself**. 0 escape-hatch deferrals. `roadmap.md` § M257.
- **M257x — 2026-08-11** · platform re-alignment (iterative) · **`closed-incomplete` — CLOSED BY USER
  RULING (`TOK-09`), NOT on gate.** Clauses 1–4 met and proven; **clause 5 out of scope by that ruling —
  never met, never measured clean.** 288 iters / 73 harden passes. The microservices→`app` map is now
  machine-fenced against `repos.yml` in both directions, all 93 org repos are enumerated, and the demo +
  dev stacks both build from current `main`. Full narrative: `roadmap.md` § M257x.

## Recently shipped releases (older → roadmap.md / roadmap-legacy.md)
- **v2.7 "july jitter" — 2026-07-25** (tag `v2.7`) — re-ground + fidelity + field-hardening; M246→M254;
  prove-on-billion a–h live; **zero carry-forward**; 0 platform edits.
- **v2.6** 2026-07-23 (`v2.6`) · **v2.5** 2026-07-20 (`v2.5`).

## Standing backlog (fated destinations — this is an INDEX; each row's owner is the link)

Every item below is owned by the linked file. iter-129 closed the last three that were owned by **this
section and nothing else**, so no fate now depends on a field the next close overwrites.

| class | owner |
|---|---|
| **Consumed by v2.8** — the M256-clause-3 re-fating of the reserved Playthrough futures M206 / M207, each with a written verdict rather than a sixth re-reservation | [`roadmap-vision.md`](roadmap-vision.md) |
| **Awaiting the user's signature** — `PERF-M256-parallel-lane` · `PT-M257-self-evaluation` · `PT-M257-talk-to-data` (roadmap calls, not routing ones) | [`roadmap.md`](roadmap.md) § M256 closure |
| **Platform defects** Rosetta cannot fix (zero platform edits binding) — **7** entries with `file:line`, **3 of them M257x**; it had **zero** until iter-102 | [`platform-defect-register.md`](platform-defect-register.md) |
| **DROPPED** + **still unscheduled** — incl. the three moved out of here at iter-129 (`DEF-M250-01` · `CAVEAT-1` · `PT-M256-resume-fixture-pair`) | [`roadmap-vision.md`](roadmap-vision.md) § Unscheduled backlog |

## Process flags (do NOT auto-push)

- 📌 **Provenance of every `billion`-measured M255 number now lives with the milestone that took it** —
  [`m255…/progress.md`](releases/02.80-fast-build/m255-build-bench-host-headroom/progress.md) § Provenance
  (moved there at iter-129, its first durable owner): the window, the pre-freeze condition, the 658/666/672
  cluster, and the three claims owed a re-confirmation on the first post-freeze campaign.
- ⚠️ **The stack's PINNED `stackseed` can be older than the authoring copy.** M256 added three `Persona`
  fields, so a `--reset` with the stack's binary **truncates the world and then fails to re-seed**, leaving it
  EMPTY. Shadow the authoring build on `PATH` for any reset. Cost the M256 close one run before it was found.
- **v2.7 is merged to `main` + tagged `v2.7` LOCALLY; NOT pushed to origin** — the user runs origin publishes
  on their own cadence. **v2.5** and **v2.6** are likewise local-only.
- rext code-of-record: authoring copy on `main`; **M257's tooling ships at `fast-build-m257-close` @ `679a5f7`
  — on origin, rung-zero verified** (`main` pushed to the same sha), which is `fast-build-m257-harden-1`
  plus the close's **2** commits.
- **Rung zero:** `git push --tags` is part of shipping a tool. Verify a tag is on **origin** before any
  prove-it-live step.
- **`run-playthroughs.sh` is BINDING since M256** — a full run exits non-zero when ptreport's gate is unmet
  (advisory on a scoped run). Anything that ran the suite and trusted a zero exit is now genuinely gated.

_Last updated 2026-08-12 — the M257 close. Both trees clean; `rosetta-extensions` pushed + tagged, `rosetta` local-only by the user's cadence._

> **Budgets: every one of them met** — file 14,709/15,360 · frontmatter 1,214/1,860 · body 13,489/13,500 ·
> all six fields in budget. The body budget was **raised once against a measurement** at iter-129 (12,000 →
> 13,500, and the frontmatter 2,600 → 1,860 so the two now sum *exactly* to the file cap, which the old
> triple did not). Derivation, the two probes that were narrower than their own conclusion, and the
> **re-raise guard**: [`context.md` § state.md contract](context.md).
