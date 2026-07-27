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

---

## M255: Hardening

### Pass 1 — 2026-07-27  (single pass; **halted early at the user's 48 h `billion` freeze**, see Stop condition)

**Scope manifest** (milestone-touched, for `/developer-kit:close-milestone` Phase 4 to reuse):

| module | files | tests |
|---|---|---|
| rext `stack-core` | `buildbench.py` · `union_apply_guard.py` · `demo_knob_guard.py` · `hostprofiles/{billion,laptop}.json` · `README.md` | `tests/test_buildbench.py` · `tests/test_union_apply_guard.py` · `tests/test_demo_knob_guard.py` · **`tests/test_m255_mutation_battery.py` (net-new)** |
| rext `demo-stack` | `up-injected.sh` | `tests/test_tooling.py` · `tests/test_frontend_build.py` · `tests/test_cert_remint_m255.py` |
| corpus | `build-budget.md` (net-new) · `safety.md` · `demopatch-spec.md` · `frontend-tier.md` · `demo-up-defaults.md` · `demo/README.md` · `CLAUDE.md` · `tailscale-serve.md` | the 5 corpus guards |
| plan | `context.md` · `roadmap.md` · `state.md` · M255 + M257 dirs | — |

**Coverage delta (milestone-touched rext modules):**

| file | statements | stmts total | newly covered |
|---|---|---|---|
| `buildbench.py` | **71 % → 74 %** | 597 → 694 | the whole report/verdict/headroom-refusal path, previously unexercised |
| `union_apply_guard.py` | **90 % → 92 %** | 134 → 150 | clause 0 + the divergent-declaration path |
| `demo_knob_guard.py` | 93 % (unchanged — nothing new landed here) | 193 | — |
| **total** | **78 % → 80 %** | 924 → 1037 | |

**Read the denominator, not just the percentage.** buildbench gained **97 statements** of new fail-closed
logic in this pass, so +3 points of coverage is ~110 newly-covered statements, not 18. The single largest
remaining gap is unchanged and is listed under Stop condition: the non-dry-run rep body.

**stack-core tests: 226 → 272.**

**Bugs fixed inline — the theme is `an empty result reported as a pass`:**

1. **`buildbench report` exited 0 on a campaign that measured nothing.** Three ABORTED reps still write
   ledgers, so `reps == 3`, and `return 0 if rep.get("reps") else 2` handed back **success** for a campaign
   whose `p50` was literally `None` — from the harness whose entire purpose is to refuse un-measured
   numbers, and in violation of its own documented exit-code contract (*"1 = … the campaign report is
   red"*). `build_report` now computes `ok` + an enumerated `red[]`, and the CLI returns 0/1/2 accordingly.
2. **The headroom assert SKIPPED any clause whose input was `None`.** A rep whose sampler died wrote zero
   samples, handed `peak_load1=None` in, and got `ok=True` — a green headroom verdict on a host nothing had
   measured. Added `require_measured`, so a caller declares which inputs are mandatory and a `None` becomes
   a failure. (`pre_rep_assert` legitimately cannot measure load1 *yet*; that is a different fact from
   unmeasured-because-broken, and only the caller knows which.)
3. **`docker_vm_disk_avail_gib() or host_disk_avail_gib()` — and `0.0` is falsy.** A Docker VM with **zero**
   bytes free is the exact ENOSPC state clause 3 exists to catch, and it fell through to host `/`, which has
   plenty; the assert PASSED on a full engine disk. Two call sites; both now go through one
   `effective_disk_avail_gib()` that distinguishes `None` (probe failed) from `0.0` (probe answered).
4. **A stale `autoverify.json` was read as a fresh green.** `autoverify.sh:376-381` writes `ts` and says in
   as many words that it exists so *a grader can SEE that the verdict predates the stack it describes* —
   buildbench IS that grader and never looked. A rep whose bring-up dies before autoverify runs is exactly
   the path that inherits the previous rep's green. `read_verdict()` is fail-closed on missing / unparseable
   / undatable / predating, and `parse_verdict_ts` is **explicitly UTC** (M236's green-gate bug was this
   same string parsed as LOCAL, failing OPEN west of UTC).
5. **`phases_complete` was printed as a warning and then ignored.** The rep still counted toward the
   campaign, contradicting this module's own fail-closed contract. Extracted `rep_is_ok()` — six clauses,
   each a way for a rep to look fine and mean nothing.
6. **`union_apply_guard`'s "DIVERGENT PATH" clause was the `_sha(p) != _sha(p)` tautology again.** It
   compared `web_decls[slug]` against `str(_manifest_path(slug))` where `web_decls` was built as
   `{s: str(_manifest_path(s))}` — *the same tautology, one indirection longer, shipped by the pass whose
   purpose was to delete tautologies.* The captured `file` group was being discarded, so a per-image variant
   that is **actually wired in** (`patches/<slug>/<slug>.hiring.yaml`) was invisible: every other clause
   re-derives `<slug>.yaml` by construction and would have audited a file the build does not use. Added
   `_manifest_decls()` (parses the filename) + clause 0 (canonical-filename) + DIVERGENT DECLARATION.
7. **`reclaim()`'s docstring still carried the claim the campaign REFUTED** — *"CACHED-step records are
   touched by every rep, so their clock keeps resetting and they survive"*. Measured: rep-01 served
   studio-desk's chain from cache and the next reclaim evicted it anyway (7 records, 356.8 MB → +173 s).
   Rewritten around what is actually true (one-off, not a per-rep tax → warm-up rep → p50, never the mean →
   `n ≥ 3` is a floor), and `parse_reclaimed_bytes()` + `_reclaim_attribution()` now make the link
   mechanical: the report NAMES the reclaim that explains an outlier, or says the outlier is **unexplained**.
8. **`load_host_profile` accepted a profile with no provenance.** A test spot-checked the two shipped files
   for `measured`; nothing stopped a third arriving without it, or with `cores: "eight"`. D-v28-6 cut the
   auto-planner precisely so no number in a profile is a guess — the loader now requires `measured` +
   `budget_source` and type-checks the five numerics.

**Tests added: 46** (7 rep-ok · 8 report/gateability · 3 reclaim-attribution · 3 byte-parsing ·
7 verdict-freshness incl. the TZ regression · 5 headroom-unmeasured · 4 falsy-zero disk · 6 profile
validation · 2 union-apply · 1 mutation battery [3 cases, 10 mutants]).

**The mutation battery is the pass's real deliverable.** `tests/test_m255_mutation_battery.py` reverts each
fix to its pre-fix form in a staged copy and asserts the suite goes RED — carrying M220's three anti-theatre
assertions (baseline GREEN · every mutant changes bytes · signatures not all identical) plus a **fourth**:
mutants reverting *different* fences must produce *different* signatures, while two halves of the *same*
fence may coincide. **It caught two bugs in itself on its first two runs** — a signature parser that returned
the literal string `"FAILED"` for every mutant (so all ten looked identical), and an over-strict uniqueness
rule that demanded a test which cannot exist. A fence nobody has watched go red is not a fence; that now
includes this one.

**Knowledge backfill:** the `stack-core/README.md` guard index gains a row for the battery. No corpus doc
needed a change — every finding was in rext code or its own docstrings, and the two corpus-facing facts
(the refuted reclaim reasoning; `p50`-never-mean) were **already** correct in `build-budget.md`, which is
what let the code's stale docstring be spotted as a divergence.

**Flakes stabilized:** none observed — 3 consecutive clean runs of the new tests.

### Stop condition

**Halted after pass 1 by explicit user instruction** (a 48 h operational freeze on `billion`, 2026-07-27 →
~2026-07-29, halting roadmap work), not by the stabilization rule. Coverage signal was still meaningful when
the pass stopped. Remaining work is recorded as **Fate 3 → M255 harden resume**:

| item | what it would do |
|---|---|
| `run_campaign` rep-body coverage | the non-dry-run block (`buildbench.py` ~763-850) is still the largest uncovered region; drive it with a faked `Popen`/`Sampler`/docker probe set so the staleness, unmeasured-sampler and phase-table paths are proven end-to-end and not only unit-wise |
| a plan-number mirror fence | `666.29` is mirrored in 8 prose sites and the `D-v28-N` range in 5; the range had **already** rotted in 4 of them (fixed here by hand). Derive the baseline from `hostprofiles/billion.json` and the range from `roadmap.md`, and fence both — the C1 mirrored-count class, which this release has now paid for twice |
| `demo_knob_guard` anchor-fence mutants | the `Read at` fence is the third M255 guard and has **no** battery entry; add mutants for the anchor comparison and the `--fix` regenerator |
| `_manifest_lists` body extraction | `text.find("\n}\n")` truncates a build function at the first column-0 `}`; currently masked by the pinned 11/5/6 count test, but the truncation would be silent |
| the `laptop` profile's provisional field | `projected_image_gib` is the one non-measured number in either profile and says so only in prose; make it a machine-declared `provisional_fields` list the loader can surface |
