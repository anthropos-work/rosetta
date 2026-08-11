---
active_release: "v2.8 «fast build» — IN DEVELOPMENT (branch release/02.80-fast-build, designed 2026-07-27). Time-to-ready: from nothing, to live, to provably live, fast. **5** milestones M255 → M256 → M257x → **M257** → M258, strictly serial. **3 of 5 closed.** Tooling + docs only, 0 platform edits. Detail: roadmap.md § v2.8."
active_branch: "release/02.80-fast-build"
active_milestone: "M257 «first-light build» (`iterative`, branch m257/first-light-build) — RUNNING, 5 iters closed. Gate re-pointed at `macmini` at iter-05, targets unchanged."
last_closed: "M257x — 2026-08-11"
phase: "M257 in progress. M257x closed 2026-08-11 by USER RULING (`TOK-09`), NOT on gate — clauses 1–4 met and proven, **clause 5 OUT OF SCOPE by that ruling and never met**. M257's gate named the retired `odysseus` and was ungradeable; iter-05 re-pointed it at `macmini` (a stale-reference repair — every target survived) and retracted the *'a Mac pays no unpack leg'* claim that paused the milestone. Next: `FIX-M257-load1-units-vm`, then the `n ≥ 3` **contended** baseline that fills `macmini.json`'s `gated_baseline`, then levers. Carry-forward from M257x → M258: 11 items / 5 clusters + 1 block fate."
last_updated: "2026-08-11"
---

# State

**v2.8 "fast build" IN DEVELOPMENT** — designed 2026-07-27, branch `release/02.80-fast-build` cut from `main`.
**3 of 5 milestones closed.** Full narratives live in [`roadmap.md`](roadmap.md) § M255 / § M256 / § M257x — not here.

## Hosts (D-v28-15, 2026-07-31 — supersedes D-v28-14)

- 🎬 **`billion` — the OFFICIAL host.** Demo deployment only. Not for development or testing.
- 💻 **dev/test = LOCAL to the new Mac.** The old laptop and **`odysseus` are both retired** from this project.
- ✅ **~~Host-class mismatch blocks M257~~ — RETRACTED 2026-08-11 (M257 iter-04 measured, iter-05 wrote it
  down).** This bullet read: *"a Mac is **arm64/overlay2** … **the Mac pays no unpack leg**, which is exactly
  what M257's L1 (~200–250 s) optimises. M257's speed gate is **un-measurable** on the sanctioned hosts as
  written."* **False for this machine.** The Mac mini runs the **containerd image store** and pays a
  size-proportional unpack leg — measured 0.8 s @ 256 MB → 3.0 s @ 1024 MB, and **56.6 s export + 19.3 s
  unpack** on the real 4.12 GB hiring image. The generalisation came from the **retired M1 Pro laptop**, a
  different machine; *"a Mac"* is a class, and the fact that mattered is a per-machine Docker Desktop toggle.
  **Do not re-derive this from `docker info`** — it prints `Storage Driver: overlayfs` here, which is exactly
  what made the wrong reading look right; grade it with a probe. So **L1 keeps a substantial price locally
  (~136–152 s), and M257's gate IS measurable on the host that exists** — re-pointed at `macmini` at M257
  iter-05. What still holds: billion is x86_64 and this host arm64 (**4.84 vs 4.12 GB**, a ~15 % gap, not the
  ~40 % the laptop suggested), and **seconds measured here do not transfer to billion**.
- 🔧 **New-Mac bootstrap:** `.agentspace/rext.tag` is **git-ignored** → a fresh clone has no pin, and a mismatch
  is **FATAL** in `ensure-clones.sh:94-101`. Create it deliberately. (On the old box the SoT was 63 commits
  behind `main`, so `/demo-up` aborted there.)


## Next up

**M257 — first-light build** (`iterative`), **UNPAUSED** by M257x's close and **RUNNING** since iter-04.
**Its exit gate has been re-cut** (iter-05): it named `odysseus`, retired by `D-v28-15` on 2026-07-31, so it
could not be graded at all — the host reference now names **`macmini`**, the local M4 Pro Mac mini, whose
profile *was* measured and checked in at iter-04 (`stack-core/hostprofiles/macmini.json`, deliberately
without a `gated_baseline`). **Every target survived the re-cut unchanged** — p50 ≤ 360 s over 3 consecutive
cold cycles, 0 platform-repo edits, G1–G7, both falsifiable asserts, the ≤ 300 s stretch. Banked from its
closed iters, not to be redone: both gate-honesty instruments landed with mutation-proven controls, B1+B2
fixed, the mirror fence parameterised by host, the host measured. Still owed: the `n ≥ 3` baseline —
**taken on a permanently contended box and labelled so** — and `FIX-M257-load1-units-vm`, which is what
`INVESTIGATE-M257-load1-48` became: clause 1 grades **host** `load1` against a **VM-allocation** core count,
computing a limit of **6** here where the correct one is **10**, so it **fails closed**.

**M257x carry-forward lands in M258**, not M257 — 11 items in 5 root-cause clusters + 1 block fate. Owner:
[`m257x…/carry-forward.md`](releases/02.80-fast-build/m257x-platform-realignment/carry-forward.md).
The one with an operational deadline: **the tooling fixes from M257x's close are NOT on a pushed tag**, so
no stack can obtain them until someone tags and `git push --tags`. *Tagging is not publishing.*

## Phase

Between milestones. **M257x closed 2026-08-11 by USER RULING (`TOK-09`) — not on gate.** Clauses 1–4 met
and proven; **clause 5 is OUT OF SCOPE by that ruling, was never met, and must never be reported as
"measured clean".** Full closure narrative: [`roadmap.md`](roadmap.md) § M257x. Clause-by-clause table,
metrics and routing live in that milestone's own `carry-forward.md` / `metrics.json` / `retro.md`.

⚠️ **Open safety item, routed to M258 and not closed:** a demo reached the **production** S3 bucket and only
an **IAM policy on an account we do not control** refused it. The containment is proven by a unit test on
the emitter and **on no running stack**; both currently-running stacks still carry the pointer, and the
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
**M256 playthrough sharpening** ✅ **done 2026-07-30, `closed-on-gate`** → **M257x platform re-alignment** ✅ **done 2026-08-11, `closed-incomplete` (user ruling, not gate)** → **M257 first-light build**
(`billion` 666 s → ≤ 360 s p50) → **M258 proven-live build** (up AND self-proven, ≤ 480 s p50). Strictly serial by the
user's order — *sharpen the detector before changing what it detects*.

## Binding user decisions (2026-07-27, + later)

**`D-v28-1` … `D-v28-15` are held in full in [`roadmap.md`](roadmap.md)** — including the two that produced
standing rules: **`D-v28-3`** (batch-gate semantics — **zero standing red** is the invariant) and
**`D-v28-12`/`D-v28-13`** (***never gate a relative statistic without first measuring its variance***).
M256's close ratifications (`D103`, `D104`, the iter-31/32 deviation) are in that file's § M256 closure;
`HARDEN-CAP-ACCEPTED-D105` is in [`m256…/decisions.md`](releases/02.80-fast-build/m256-playthrough-sharpening/decisions.md).

## Recently closed milestones (max 5)

_Trimmed to the last 3 days per the state.md contract; older entries live in `roadmap.md`'s `### M{N}` blocks._

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
- rext code-of-record: authoring copy on `main`; M256's tooling ships at **`fast-build-m256-close`** @ `ce345e1`
  — **on origin, rung-zero verified** — which is `fast-build-m256-harden-final` plus the close's **3** commits.
- **Rung zero:** `git push --tags` is part of shipping a tool. Verify a tag is on **origin** before any
  prove-it-live step.
- **`run-playthroughs.sh` is BINDING since M256** — a full run exits non-zero when ptreport's gate is unmet
  (advisory on a scoped run). Anything that ran the suite and trusted a zero exit is now genuinely gated.

_Last updated 2026-08-07 — M257x iter-129. Both trees clean and pushed; `demo-1` up._

> **Budgets: every one of them met** — file 14,935/15,360 · frontmatter 1,656/1,860 · body 13,281/13,500 ·
> all six fields in budget. The body budget was **raised once against a measurement** at iter-129 (12,000 →
> 13,500, and the frontmatter 2,600 → 1,860 so the two now sum *exactly* to the file cap, which the old
> triple did not). Derivation, the two probes that were narrower than their own conclusion, and the
> **re-raise guard**: [`context.md` § state.md contract](context.md).
