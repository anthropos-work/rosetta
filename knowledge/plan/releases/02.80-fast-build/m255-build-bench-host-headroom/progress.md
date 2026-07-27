# M255 — progress

## Section checklist

- [x] 1. `buildbench` harness (rext `stack-core`) — n≥3 on billion, **cold-images variant only**; JSON phase ledger + 10 s sampler; **every entry records the invocation + full `DEMO_*` env snapshot**; one informational n=1 laptop run — **campaign RUN and LANDED: `n=3 p50 666.29 s`, 3/3 green** (see § the baseline, below)
- [x] 2. Campaign protocol + reclaim — hard-failing pre-rep disk/cache assert · reclaim step between reps (**L6 promoted here**) · per-rep `docker system df` declaration · **`DEMO_DISK_MIN_GIB` re-sized** · ENOSPC→"redis exited (1)" signature noted
- [x] 3. Host profiles + headroom assert — `stack-core/hostprofiles/{billion,laptop}.json` measured + checked in; a **failing** sampler assert (load1 / summed heap / free disk); decision recorded reconciling "fail loudly" vs the never-block-a-bring-up pre-flight contract
- [x] 4. The **union-apply** parallelism rule + guard test (shared members byte-identical; non-shared under disjoint `apps/*` or waived inert). "Separate clones" option deleted
- [x] 5a. **Spike (a) — the 15-min L1 experiment** on the rext-owned `hiring.Dockerfile`; measure the export delta; record `NEXT_PRIVATE_STANDALONE=1` + its demopatch fallback  ← **THE BARRIER DECIDER**
- [x] 5d. Spike (d) — is peak load1 4.90/8 a plateau or an I/O ceiling?
- [x] 5e. Spike (e) — host-vs-peer topology for M258's composed command
- [x] 6. `corpus/ops/demo/build-budget.md` (net-new; the blind area)
- [x] 7. Security + cert hazard (**non-gating**) — expiry-aware re-mint + the paired `corpus/ops/safety.md` §3 amendment
- [x] **BARRIER VERDICT** recorded — **GO**

---

## The baseline — **`n=3 p50 = 666.29 s`** (2026-07-27)

The milestone's other deliverable: v2.8 now has a **measurement floor**, where before it had one `n=1` number
from a one-off shell script that lived on a single box and was never in version control.

`buildbench run 1 --reps 3 --profile billion --public-host billion.taildc510.ts.net --label m255-baseline`.
**3/3 reps `rc=0`, `autoverify green / 0 warnings`, `phases_complete`, zero missing anchors, headroom OK.**
Artefacts: `billion:/home/devops/panorama/m255/campaign/`.

| | p50 | min | max |
|---|---|---|---|
| **total cycle** | **666.29 s** | 658.15 s | 881.01 s |
| UI-tier image builds (3) | **436.1 s — 65.5 %** | | |
| image export + unpack alone | **307.5 s — 46.2 %** | | |

**It supersedes the n=1 672.4 s, which it lands within 0.9 % of** — so no lever was mis-ranked, and **M257's
exit gate is now a percentage of a real number** (re-pinned in its `overview.md`, along with the roadmap,
`state.md`, `context.md`, `demo/README.md` and `CLAUDE.md`).

**The headroom model was validated, not just applied.** Clause 2 predicted `1 × 3900 + 1500 = 5400 MiB`; the
reps peaked at **5446 / 5579 / 5398 MB** — inside **~3 %**. `lane_heap_measured_peak_mib = 3900` now rests on
a campaign rather than on the single annotated run, and that is recorded in `billion.json` itself.

### The campaign found a fifth wrong corpus claim — this doc's own reclaim reasoning

rep-02 cost **881.01 s (32 % above p50)**, and 206 s of the 215 s excess is one thing: studio-desk's cache
chain was **evicted by the reclaim step**, so it paid a full 136.5 s `npm ci` the neighbouring reps got free.
The protocol justified `--filter until=24h` on the grounds that *"CACHED-step records are touched by every rep,
so their last-used clock keeps resetting and they survive."* Measured, that is **false**: rep-01 served that
chain from cache and the very next reclaim evicted it anyway (**7 records, 356.8 MB → +173 s**).

It is a **one-off**, not a per-rep tax (the following reclaim pruned **0 B / 0 records**), which is why the
corrected protocol now says: **budget a warm-up rep, report `p50`, never the mean** (the mean here is 735.2 s
— 10 % high, describing an eviction rather than a bring-up), and **`n ≥ 3` is a floor**: at `n = 2` this
campaign would have reported **~773 s** and every v2.8 lever would have been priced against a number that does
not exist. Two further protocol figures were corrected the same way: per-rep cache growth is **+1.7–2.2 GB**,
not the claimed ~11.6 GiB, and the binding disk constraint is the **~18 GiB mid-cycle transient**, not the
~2 GiB a steady rep nets.

One honest caveat entered alongside: **1 of 221 samples hit 100 % disk `%util`**. Average util stayed
20.8–23.9 %, so spike (d)'s "not an I/O ceiling" reading stands — but *"zero samples ≥ 90 %"* is an `n=1`
claim and is now marked as one.

---

## Barrier verdict — **GO** (2026-07-27)

**Spike (a) is the decider, and it did not merely pass — it beat the estimate.** Prototyping the multi-stage
`.next/standalone` shape on the rext-owned `hiring.Dockerfile` (no demopatch, no platform-edit question):

| | shipped single-stage | multi-stage standalone |
|---|---|---|
| image | **4.84 GB** | **379 MB** (−92.2 %) |
| **export step** | **146.8 s** | **2.9 s** (−98 %) |
| build wall | 225 s | **49 s** |

**L1 does not collapse. M257's exit gate does not need re-cutting on L1's account** — the annotated baseline
pays 278.6 s of export across the two Next.js images, so L1's ~200–250 s estimate is conservative. The
standalone image also **boots and Clerkenstein-redirects correctly** (`307 → …:15400/sign-in`), so this is a
functional shape, not just a smaller one.

**Two riders M257 must carry, both discovered here:**

1. **`turbo --env-mode=loose` is mandatory.** Turbo 2 defaults to `strict`, which filters
   `NEXT_PRIVATE_STANDALONE` out before `next build` sees it — the flag then **silently no-ops** and the
   build is green with the old 4.84 GB image. Keep the `RUN test -f …/standalone/apps/hiring/server.js`
   guard.
2. **L2 must be re-priced, downward, and sequenced AFTER L1.** Spike (d) refutes both the plateau and the
   I/O-ceiling hypotheses (peak load1 3.75/8; peak disk `%util` 63.4 %, **zero** samples ≥90 %). L2's value
   was overlapping two ~140 s export legs — **L1 deletes them**. After L1 the hiring image costs ~49 s of
   which ~44 s is `next build`, so **L2 buys ≲45 s, not ~200 s**; and the two *compile* legs cannot overlap
   on `billion` at all, because the headroom assert derives `max_parallel_ui_lanes = 1`. This is a **gate
   input for M257**, not a barrier failure.

---

## What else the milestone found (and fixed)

The Phase-0b KB-fidelity audit came back **RED** with three blockers, all resolved here (Fate 1):

- **D-v28-7's *"inert for the hiring image"* premise is FALSE.** The variable it depends on *is* set on the
  hiring container and `apps/hiring` imports the patched module. Union-apply survives — the change is
  beneficial (hiring inherits the M218 SSR fix) — but "beneficial" is not "inert", and **M257 must re-verify
  the hiring recruiter Playthrough** after flipping it on.
- **`demopatch-spec.md` §4 misstated both facts the union-apply rule rests on** — a "4-manifest union" that
  is 7, and a "revert LIFO, identical on both" that neither build does. Corrected, and the narrow invariant
  that *does* hold is now machine-fenced in both builds and both phases.
- **`frontend-tier.md` argued against the barrier's own build shape**, calling `output:'standalone'` a
  forbidden upstream PR while rext has owned `hiring.Dockerfile` since M224. Three shapes now documented.

Plus two unfenced doc-rot classes, both closed with a guard clause and RED-proven tests: the **mirrored knob
count** (three sites said "all 27" against 29 parsed) and the **`Read at` column** (**28 of 29** citations
drifted while the guard reported green, because it compared names and never anchors — now regenerable with
`demo_knob_guard.py --fix`).

And one bug of our own making, caught live: `buildbench`'s sampler stored its stop-Event on `self._stop`,
shadowing `threading.Thread._stop`, so a campaign ran the full bring-up and then died writing the ledger. It
cost one 11-minute cycle — whose numbers were recovered with `buildbench parse`, and which turned out to be
the run that answers spike (d).

Plus one fake fence found in our own new code: `union_apply_guard`'s clause (i) shipped as
`_sha(p) != _sha(p)` — a tautology that could never fire. Replaced with the proposition it was reaching for
(a shared member must be **one** manifest file, and both builds must resolve the slug to the same path),
RED-proven.

## Shipped

- **rext** `fast-build-m255-buildbench-2` — **pushed to origin and rung-zero verified** by a fresh
  `git clone --branch <tag>` from origin carrying the new tooling. `buildbench.py`, both measured host
  profiles, `union_apply_guard.py`, the `demo_knob_guard.py` anchor fence (`--fix`). **226 stack-core tests
  pass**; full tooling suite green at **1418 passed / 2 skipped**.
- **corpus** `corpus/ops/demo/build-budget.md` (net-new) + the `safety.md` §3.5.4 cert amendment, and the
  baseline re-pinned across `demo/README.md`, `CLAUDE.md`, `roadmap.md`, `context.md`, `state.md` and M257's
  `overview.md`.
- **0 platform-repo edits.**

## Cut from the first draft (see roadmap.md § design decisions)

- ~~`hostprofile` auto-planner~~ → replaced by item 3 (D-v28-6)
- ~~truly-cold bench variant~~ → replaced by spike (a); optional one-shot post-M257 (D-v28-8). *One
  truly-cold data point arrived anyway, informationally: the laptop run had an empty BuildKit cache.*
- ~~§8.5 prose retraction~~ → moved to M257 so `frontend-tier.md` moves once (D-v28-10). *Annotated in place
  so M257 rewrites the model, not just the numbers.*
