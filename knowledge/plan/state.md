---
active_release: "v2.8 «fast build» — IN DEVELOPMENT (branch release/02.80-fast-build, designed 2026-07-27). Time-to-ready: from nothing, to live, to provably live, fast. **5** milestones M255 → M256 → **M257x** → M257 → M258, strictly serial; M257x was INSERTED 2026-07-31 and **M257 is PAUSED behind it**. Tooling + docs only, 0 platform edits. Detail: roadmap.md § v2.8."
active_branch: "release/02.80-fast-build"
active_milestone: "M257x — platform re-alignment (iterative), 127 iters + 26 harden passes closed (branch m257x/platform-realignment). Re-align BOTH rosetta (corpus) and rosetta-extensions (tooling) to platform @ origin HEAD. Gate 4 of 5; clause 5 the only one open. M257 PAUSED behind it after 3 iters."
last_closed: "M256 — 2026-07-30"
phase: "M257x iter loop, **127 iters** (+ 26 harden passes). **Gate 4 of 5**; clause 5 open, `P` UNMEASURED, no reading taken. **Headline, and it is a REFRAME: the corpus is UNDER-CITED, not unfounded** — iter-124 triaged the sealed consequence class (`|C1|`). **Quote the AUDIT-CORRECTED split with its denominator or quote none**: `cite` ≈ **86.6 %** (printed 96.2 %), `hedge`+`drop` ≈ **12.2 %**, and **`fix` = 4 printed is a FLOOR, ≈ 11 ≈ 3.2 % recall-corrected** — falsity is not syntactic. So *"1,151 uncited assertions"* is ~298 whose evidence sits in a clone nobody linked, ~42 genuinely uncheckable from here, and ~11 that are false: three diseases, three treatments, only the third urgent. Both pre-registered branches were sealed before a single verdict and **neither fired**. Detail: `iter-124/audit.md`; live status: `progress.md`."
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

_Counts and the headline are in the frontmatter; this section is the clause-by-clause state only._

- **Clauses 1–2 — CLOSED** at platform `0c91421`, **MET WITH DISCLOSURE that travels with them forever**:
  a freshly built stack failed the first full run in 2 of 2 attempts, so never a clean pass.
- **Clauses 3–4 — hold**, asserted by fences that are watched going RED, not by inspection.
- **Clause 5 — the only open one**, met ONLY by a reading that returns zero. Five graded readings
  (iters 98–119) established that **inflow ≈ outflow** — repair efficacy is confirmed, yet the residual
  does not fall, because clone advance and the repair's own induction feed it. Each reading's arithmetic
  is owned by its own `iter-NN/adjudication.md`; none is current.

**There is NO active strategy, and that is the state — not an omission.** `TOK-07` was refuted at iter-116
by its own pre-registration and `TOK-08` (the USER's re-scope) at iter-119 the same way; `TOK-08`'s sealed
rule **bars a successor**, so there is **no `TOK-09`** (chain in `decisions.md`). **iter-122 added an
INSTRUMENT, not a strategy** — `F4` books any sentence treating the census as clause 5's grader as a defect.

**Chapman is RETIRED**; only floors survive (**≥ 46 at `194361e4`**). The residual stays unmeasured above
a floor — **1,908** tier-1 pairs, **1,160** tier-2 (baseline 1,164; the ratchet holds). The two
measurements that govern how any of it may be quoted are a standing rule now, below.

**The live detail lives in the milestone's own docs — `state.md` is the index, not the narrative**
(`context.md` § state.md contract). All paths below are under
`releases/02.80-fast-build/m257x-platform-realignment/`:

- [`progress.md`](releases/02.80-fast-build/m257x-platform-realignment/progress.md) — **authoritative live
  status**: every iter's findings and the carried items
- [`decisions.md`](releases/02.80-fast-build/m257x-platform-realignment/decisions.md) — `D-N` + the
  `TOK-01…08` strategy chain
- [`gate-clauses-1-2/README.md`](releases/02.80-fast-build/m257x-platform-realignment/gate-clauses-1-2/README.md)
  — **owner of every clause 1 & 2 number**: the five cold-cycle timings, the reflog freeze-proof, the
  disclosed first-run failure
- [`iter-124/audit.md`](releases/02.80-fast-build/m257x-platform-realignment/iter-124/audit.md) — **owner
  of the under-cited reframe and of every corrected figure in `phase:`**: the seeded 30-of-344 hand audit,
  R3 100 % / R4 66.7 %, both sealed branches checked
- [`iter-123/progress.md`](releases/02.80-fast-build/m257x-platform-realignment/iter-123/progress.md) —
  the org census: 93 repos re-derived, five corrections, the `infrastructure` read, the 89-row re-pin
- [`iter-122/adjudication.md`](releases/02.80-fast-build/m257x-platform-realignment/iter-122/adjudication.md)
  — **the CLAIM CENSUS**: 525/525 adjudicated, two fired falsifications, its own three defects
- [`iter-119/adjudication.md`](releases/02.80-fast-build/m257x-platform-realignment/iter-119/adjudication.md)
  — the last graded reading (`P = 22`) + the test–retest that retired the series

## Standing rules (outlive the milestone — do NOT move these into `phase:`)

Each rule's derivation lives at its owner; these are the headlines a reader must not miss.
**`§5`** = [`corpus/ops/platform-alignment.md`](../../corpus/ops/platform-alignment.md) § 5, which owns the
numbered rules in full.

- **§5 rule 44 — no single search tool is safe.** Gitignored-but-tracked files, NUL-bearing source and
  nested untracked repos each blind a *different* instrument. Use the mechanized per-tree path
  (`anchor_construct_guard._clone_of`), never a hand rule — **rule 44's own worked recipe returns 2 where
  it publishes 22.**
- **TWO INSTRUMENTS, and conflating them has cost this milestone repeatedly.** Clause 3's is the **guard
  family**; clause 5's is the **graded read** (frozen, sha `3858ec53…`, one commit ever). A guard going
  green says nothing about clause 5, and a reading says nothing about clause 3.
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
- **M256 — 2026-07-30** · playthrough sharpening (iterative) · **`closed-on-gate`**, 32 iters / 3 harden
  passes · 18 → **30 live Playthroughs** · **all three gate clauses proved unmeetable as first authored**
  · full narrative: `roadmap.md` § M256.
- **M255 — 2026-07-28** · build-bench & host-headroom (section, HARD barrier) · **VERDICT GO** · baseline
  n=3 p50 666.29 s on `billion` · full narrative: `roadmap.md` § M255.

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
  [`platform-defect-register.md`](platform-defect-register.md), **7** entries with `file:line`
  — **3 of them M257x** (the prod-bucket compose default, the fail-open authz gate, the dev-login
  session-mint routes). It had **zero** M257x entries until iter-102.
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

_Last updated 2026-08-07 — M257x iter-128. Both trees clean and pushed; `demo-1` up._

> **Budgets:** file + frontmatter + all six fields **within budget**; the **body is over and is not
> trimmable by the contract's own method** — measured, derived and routed in
> [`iter-128/progress.md`](releases/02.80-fast-build/m257x-platform-realignment/iter-128/progress.md) § 1.
