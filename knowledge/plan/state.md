---
active_release: "(none — v2.9 «new alphabet» SHIPPED 2026-08-16, tag `v2.9`. Awaiting /developer-kit:design-roadmap for the next release.)"
active_branch: "main"
active_milestone: "(between releases)"
last_closed: "M265 — 2026-08-16"
phase: "between releases — awaiting /developer-kit:design-roadmap. v2.9 shipped 2026-08-16 (tag `v2.9`): the taxonomy went 43,584/22,511 -> 3,562/706 and the corpus, tooling and demo learned it, proven cold on a demo AND a dev stack with the Playthrough suite at 222/0. It caught a demo whose simulation library rendered ZERO cards while every probe was green (187 of 302 replayed content node-ids retired). ⚠️ SIX of seven milestones shipped WITHOUT the close-milestone lifecycle — recorded, not back-filled (release-retro.md)."
last_updated: "2026-08-16"
---

# State

**v2.9 "new alphabet" IN DEVELOPMENT** — designed 2026-08-14, branch `release/02.90-new-alphabet` cut from
`main`. **0 of 7 milestones closed.** Full narrative + every citation lives in [`roadmap.md`](roadmap.md)
§ v2.9 — not here.

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

**M259 «canon ground truth»** — the HARD go/no-go barrier. Pull every platform repo fresh into **both** clone
sets, read the canon + the **redirect map** + the retired-id guard, measure the real new counts, and reconcile
the three lineages (`skills-and-job-roles` 12,201 · prod 42,790 · canon ~4k). It delivers
`corpus/architecture/taxonomy-canon.md` — the corpus has **no doc anchor for the taxonomy source pipeline at
all**, a Phase-0b blind area.

## Phase

**Designed, not started.** Every fact below was measured 2026-08-14 against each repo's `origin/main`; the
citations are in [`roadmap.md`](roadmap.md) § v2.9.

**The break is in OUR tooling, not the platform's.** `stack-snapshot/taxonomy/taxonomy.go:104` pins
`MinRows: 40000` on `public.skills`; `capture/capture.go:392` aborts *"refusing to persist a broken
snapshot"*. A ~4k canon cannot be captured → no snapshot → no set-dress → **no demo and no `dev-N`**. M260
owns it, and must preserve the under-capture protection the floor exists for.

**The second risk is silent, which makes it the dangerous one.** The seed presets pin **8 literal job-role
names**; `TaxonomyRefs` returns an empty pool for one that no longer resolves and `PersonaSeeder` **skips
enrichment rather than fabricating**. The closure gene gained a population witness at M256, so a **total**
wipe-out fails loudly — a **partial** one (3 of 7 heroes) still passes. M262 adds the per-hero richness floor.

**Two things already handle the swap correctly** and need no work: `autoverify.sh:305` asserts
`public.skills > 0` (not a floor), and the batch-prompt cache is keyed on the taxonomy capture version, so
generated members invalidate on re-capture.

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

## v2.9 shape

**M259** (section, HARD barrier) → **M260** → **M261** → { **M262** ∥ **M263** ∥ **M264** } → **M265**
(iterative closer). The three parallel milestones touch disjoint trees — `stack-seeding` ·
`playthroughs`+`stack-verify` · `corpus/` — so the merge surface is index/link-level only. Per-milestone
scope: [`roadmap.md`](roadmap.md) § v2.9.

## Binding user decisions (2026-08-14)

- **`D-v29-1` — "update the platform repos" means PULL THEM FRESH, never commit to them.** The zero-platform-edit
  rule is unchanged and binding for this release.
- **`D-v29-2` — the `/taxonomy` gate is "navigable and real", not "reachable".** A hero walks index → category →
  specialization → role → skill on the replayed canon, proven by a Playthrough. v2.8 shipped an academy that
  rendered perfectly and could not hydrate precisely because nothing clicked anything on it.
- **`D-v29-3` — regeneration ceiling $200, and it is REAL API SPEND.** `gen-batch` runs through the platform's
  `ai` module (gpt-4o-mini via Azure/OpenAI keys); there is **no subscription-quota path** for it. M262 **prices
  the regeneration and reports the number before spending it** rather than spending up to the ceiling on its own
  judgement; anything approaching $200 is reviewed first.

## Recently closed milestones (max 5)

_(none this release — v2.8's entries moved to [`roadmap.md`](roadmap.md) at its close.)_

## Recently shipped releases (older → roadmap.md / roadmap-legacy.md)
- **v2.9 "new alphabet"** — 2026-08-16 (tag `v2.9`) — taxonomy realignment; M259→M265;
  **43,584/22,511 → 3,562/706** skills/roles; Playthrough suite **222/0** cold; content realignment
  repaired **515 refs → 0**; 0 platform edits, 0 dependency changes.
  ⚠️ **6 of 7 milestones shipped WITHOUT the close-milestone lifecycle** — recorded, not back-filled
  (`releases/archive/02.90-new-alphabet/release-retro.md`).
- **v2.8 "fast build" — 2026-08-13** (tag `v2.8`) — time-to-ready; M255→M256→M257x→M257→M258;
  **450 s → 286.99 s** cold bring-up on `macmini` (gate 360 s, stretch 300 s); the journey suite now runs
  **inside** every bring-up; 11.54 GB reclaimed at zero build-time cost; 0 platform edits, 0 net-new deps.
  ⚠️ **M257x + M258 closed by USER RULING, not on gate**; M258 clause 3 never measured clean. Real
  carry-forward → `roadmap-vision.md` § v2.8.
- **v2.7 "july jitter" — 2026-07-25** (tag `v2.7`) — re-ground + fidelity + field-hardening; M246→M254;
  prove-on-billion a–h live; **zero carry-forward**; 0 platform edits.

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
  [`m255…/progress.md`](releases/archive/02.80-fast-build/m255-build-bench-host-headroom/progress.md) § Provenance
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
