---
active_release: "v2.8 «fast build» — ALL 5 MILESTONES CLOSED (branch release/02.80-fast-build), awaiting /developer-kit:close-release. M255 → M256 → M257x → M257 → M258, strictly serial. Tooling + docs only, 0 platform edits. Detail: roadmap.md § v2.8."
active_branch: "release/02.80-fast-build"
active_milestone: "(between milestones — the release is complete; next action is /developer-kit:close-release for v2.8)"
last_closed: "M258 — 2026-08-12"
phase: "**v2.8 is feature-complete — M258 closed 2026-08-12 `closed-incomplete`, ACHIEVED BY USER RULING, NOT on gate.** Clauses 1/2/4/5 proven; **clause 3 (composed p50 ≤ 480 s over 3 cold cycles) NOT MET and never to be recorded as met** — 840.01 s is instrument-rejected, 401.60 s is a projection, and the ~290 s warm-cache cycle was deliberately not banked. `END-M258-one-stack` MET: demo-4 is the only stack, built by the fixed tooling from the newest mains, and it proved itself in its own bring-up. Next: /developer-kit:close-release."
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

**`/developer-kit:close-release` for v2.8 "fast build"** — all five milestones closed, branch
`release/02.80-fast-build`. It inherits **one conscious block fate** rather than a scatter of silent punts:
[`m258…/carry-forward.md`](releases/02.80-fast-build/m258-proven-live-build/carry-forward.md) +
[`deferrals-audit.md`](releases/02.80-fast-build/m258-proven-live-build/deferrals-audit.md) name every item
with its fate. **0 escape-hatch deferrals** at the milestone close.

**Three items are owed the user's explicit fate there** (they cross a release boundary, which revokes a
deferral's authority): `F2` (`ptvalidate` unwired, M256→M257→M258) · `PROFILE-M257-provisional-fields`
(M255→M257→M258) · `RATCHET-M257-literal-ceilings-breached` (**pre-existing breach of 8** — 249 vs a
ceiling of 240; **never raised by anyone, at any point**).

⚠️ **The one number v2.8 never took is a clean composed p50** — not engineering, a quiet box. The arithmetic
already fits (247.79 s bring-up + ~129–179 s batch ≈ 377–427 s vs 480), and `LEVER-M257-L5-setdress` is
**still unspent**, now with a named target (the taxonomy replay, ~88 % of `set_dress`).

## Phase

**v2.8 is feature-complete.** M258 closed 2026-08-12 **`closed-incomplete` — ACHIEVED BY USER RULING, NOT
ON GATE.** Clauses 1/2/4/5 proven (1, 2 and 5 re-proven on the final stack in its own bring-up); **clause 3
— composed p50 ≤ 480 s over 3 cold cycles — NOT MET, and it must never be recorded as met.** ⚠️ **Read it
like M257x's `TOK-09`, not like M257**, whose gate fired on its own terms one day earlier. Full narrative:
[`roadmap.md`](roadmap.md) § M258.

**Every M258 number carries its status or must not be quoted:** **840.01 s** instrument-rejected (3/3
`headroom=FAIL`) · **401.60 s** a PROJECTION, never one measured cycle · **~290 s** deliberately **not
banked** (warm-cache, missing the export/unpack leg that is 46.2 % of a cold one) · **179.37 s** is
`batch_gate`'s own p50, **inside M256's 200 s budget while contended** — the batch half is not what is slow.

**`END-M258-one-stack` MET, and it is the USER'S stack.** `demo-4` is the only stack up, built by the
**fixed** tooling from the newest platform mains (`platform` `766df6c` · `app` `c52dbc51e`), and it proved
itself in the same command. **Cockpit `http://localhost:47700`. Do not tear down, re-seed, restart or
reset it.**

⚠️ **Open safety item, inherited by the release close and still not closed:** a demo reached the
**production** S3 bucket and only an **IAM policy on an account we do not control** refused it. Containment
is proven by a unit test on the emitter and **on no running stack**. Owner: `corpus/ops/safety.md` + M257x
`carry-forward.md` cluster 1.

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

**M255** (section, HARD barrier) ✅ *VERDICT GO* → **M256** ✅ `closed-on-gate` → **M257x** ✅
`closed-incomplete` *(user ruling)* → **M257** ✅ `closed-on-gate` *(449.51 → **286.99 s** p50, gate 360)* →
**M258** ✅ `closed-incomplete` *(user ruling — clause 3 never measured clean)*. **ALL 5 CLOSED.** Strictly
serial by the user's order — *sharpen the detector before changing what it detects*. Per-milestone detail:
[`roadmap.md`](roadmap.md).

## Binding user decisions (2026-07-27, + later)

**`D-v28-1` … `D-v28-15` are held in full in [`roadmap.md`](roadmap.md)** — including the two that produced
standing rules: **`D-v28-3`** (batch-gate semantics — **zero standing red** is the invariant) and
**`D-v28-12`/`D-v28-13`** (***never gate a relative statistic without first measuring its variance***).
M256's close ratifications (`D103`, `D104`, the iter-31/32 deviation) are in that file's § M256 closure;
`HARDEN-CAP-ACCEPTED-D105` is in [`m256…/decisions.md`](releases/02.80-fast-build/m256-playthrough-sharpening/decisions.md).

## Recently closed milestones (max 5)

_Trimmed to the last 3 days per the state.md contract; older entries live in `roadmap.md`'s `### M{N}` blocks._

- **M258 — 2026-08-12** · proven-live build (iterative, **the closer**) · **`closed-incomplete` —
  ACHIEVED BY USER RULING, NOT ON GATE.** Clauses 1/2/4/5 proven; **clause 3 (composed p50 ≤ 480 s over 3
  cold cycles) NOT MET and never to be recorded as met** — the 840.01 s figure is instrument-rejected, and
  the last iter's refusal to bank a flattering ~290 s warm-cache cycle stands as the honest reading. 20
  iters (18 tiks + 2 toks, one **user-directed**), 5 harden passes STABILIZED. The bring-up now ends in a
  **batch gate** that exits non-zero on a red set while leaving the stack UP, plus the world-contract
  restore leg. **`END-M258-one-stack` MET** — `demo-4` proved itself in its own bring-up. The 15-red
  escalation was **closed, not carried** (sentinel folded into `app`; `batch_seconds` **629 → 129** — the
  suite was slow because it was broken). **11.54 GB reclaimed at zero build-time cost**, build cache
  untouched. The close found **eight** more defects, two of them fail-opens to GREEN. 0 escape-hatch
  deferrals. `roadmap.md` § M258.
- **M257 — 2026-08-12** · first-light build (iterative) · **`closed-on-gate` — THE GATE FIRED** on its own
  terms. p50 **286.99 s** on `macmini` vs 360 s (under the 300 s stretch), n=3, every clause green on all
  three reps **including both falsifiable ones**. One lever of eight priced (L1) cleared it alone.
  0 escape-hatch deferrals. `roadmap.md` § M257.
- **M257x — 2026-08-11** · platform re-alignment (iterative) · **`closed-incomplete` — BY USER RULING
  (`TOK-09`), NOT on gate.** Clauses 1–4 proven; **clause 5 out of scope by that ruling — never met, never
  measured clean.** 288 iters / 73 harden passes. The microservices→`app` map is machine-fenced against
  `repos.yml` both ways, and all 93 org repos are enumerated. `roadmap.md` § M257x.

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
- rext code-of-record: authoring copy on `main`; **M258's tooling ships at `fast-build-m258-close` @
  `d06a56d` — on origin, rung-zero verified** (`main` pushed to the same sha).
- **Rung zero:** `git push --tags` is part of shipping a tool. Verify a tag is on **origin** before any
  prove-it-live step.
- **`run-playthroughs.sh` is BINDING since M256** — a full run exits non-zero when ptreport's gate is unmet
  (advisory on a scoped run). Anything that ran the suite and trusted a zero exit is now genuinely gated.

_Last updated 2026-08-12 — the M258 close, and v2.8's last milestone. Both trees clean; `rosetta-extensions` pushed + tagged `fast-build-m258-close`, `rosetta` local-only by the user's cadence._

> **Budget note:** the M258 close had to TRIM this file three times to meet the **15,360-byte** cap
> (16,097 → 15,634 → 15,448 → here). **No budget was raised to fit it.** Derivation + the re-raise guard:
> [`context.md` § state.md contract](context.md).
