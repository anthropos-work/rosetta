# iter-58 — progress

**Type:** tik
**Active strategy:** TOK-04 (P1/P2/P3/P4) + protocol §7 rules 4 and 5.

## Phase A — verify the baseline (the hand-off's designated first act)

Full `stack-core` suite, cold, nothing else running:

```
refs:
  rext:     ab81527 (main) — pre-advance HEAD was 28c99d0, tag fast-build-m257x-iter-57
  rosetta:  1937e1f
  taken:    2026-08-04T07:39Z -> 07:47Z (443.8 s)
  command:  cd .agentspace/rosetta-extensions/stack-core && python3 -m unittest discover -s tests -q
  verdict:  Ran 610 tests — FAILED (failures=1)
```

**PASS — and the stated baseline was wrong in a way that had to be measured, not assumed.**
The single red is the **iter-48 perishable answer-key fixture**
(`test_claim_twin_guard_iter48_answer_key.TestIter48AnswerKey`), which is the expected one. **Not spent.**

The count, however, came back **610** against a stated baseline of **1F / 599**, under an instruction that
*"anything other than 1F/599 is iter-57's regression."* It is not a regression, and the arithmetic closes to
the unit: `test_platform_alignment_guard.py` held **19** tests at `28c99d0^` and holds **30** now, so
iter-57 added **+11**, and `599 + 11 = 610`. iter-57's own close reports the same delta from the other side
(*"Tests 19 → 30"*). The baseline was the **pre-iter-57 count quoted in a hand-off written after iter-57
landed**. Full record: `D-M257x-58-2`.

Blast-radius check the hand-off asked for by name: `repair_postcondition.py` reads `FENCE_KIND` statically
out of the module iter-57 edited. `FENCE_KIND = "standalone"` still at `platform_alignment_guard.py:63`,
the AST reader at `repair_postcondition.py:160-174` still resolves it, **`test_repair_postcondition*` →
52 tests OK.** Coupling held.

**Baseline corrected for future hand-offs: `stack-core` = 1F / 610.**

## Phase B — advance both pins (§7 rule 4)

Two pins were being deliberately held back, each for a stated reason, and both reasons were conditional on
*"an iteration that takes a cold cycle."* This is that iteration.

| pin | from | to | held back by | why now |
|---|---|---|---|---|
| `demo-stack/clones.pin.json` → `app` | `v1.365.0` | **`v1.366.0`** | iter-57: *"changes what a bring-up consumes; `demo-1` is live clause-1/2 evidence"* | `demo-1` is `Exited (255)`; clause 1's evidence is its checked-in verdicts, not its containers |
| `.agentspace/rext.tag` | `fast-build-m257x-iter-56` | **`fast-build-m257x-iter-58`** | `D-M257x-57-6`: assertion F is a corpus guard no bring-up consumes, so advancing it was **unobservable** | a cold cycle is exactly the condition under which it becomes observable |

The rext advance is **two tags in one step** (56 → 58), and carries iter-57's assertion F plus this iter's
pin edit. `fast-build-m257x-iter-58` is **annotated, pushed, and verified on origin**
(`git ls-remote --tags origin` → `6e69caa … refs/tags/fast-build-m257x-iter-58` + its `^{}` deref
`ab81527`) — CLAUDE.md's *tagging is not publishing* rung zero, which cost M236 an entire iteration.

**What the app advance contains** — measured at open, not recalled; the full table is in `overview.md` and
the commit body. Headline: **0 migrations, 0 destructive DDL, 0 new hard-required config, 0 new env reads.**
The one `+log.Fatalf` in the diff is the pre-existing *"can't init clerk events manager"* fatal **moved**
below `orgManager`/`assignmentManager` so the `user.created` webhook can force-join a hiring token-link
candidate; same line, same error, new position. That was read out of the full `main.go` diff, because a
grep for `+log.Fatal` alone reports it as new — which is the shape of a false positive that would have
stopped this advance for no reason.

**Registered in the bring-up's own provenance record**, not just asserted here —
`stack-demo/clones.lock.json` after the cycle-1 `ensure-clones`:

```json
"app": { "pin_state": "pinned-tag", "pinned_at": "v1.366.0",
         "sha": "b948604ff86125a4e83516fbe356f210ddfc3809", "behind": null, "fetch_ok": true }
```

and `==> [ensure-clones] rext pin: consuming rosetta-extensions @ fast-build-m257x-iter-58 (matches
.agentspace/rext.tag)` — the M217 FATAL pin guard agreeing, rather than this iter claiming agreement.

> **`behind: null` is the iter-01 finding, still true and still unfixed**: a pinned clone reports no
> drift-distance at all. It is correct-by-construction for a tag pin and is *not* a defect of this advance —
> but it is the reason `.agentspace/rext.tag` being git-ignored (P2's named offender) still matters.

## Phase C — prove it cold (§7 rule 5)

**Cycle 1 went RED and it had to be attributed before anything else happened.** Three builds died on
`DeadlineExceeded: context deadline exceeded`, ~4 minutes in. The available reading — *"we advanced `app`
and the build broke"* — was wrong, and taking it would have reverted a good pin. The decisive fact is that
one of the three failures was `fake-bapi`, whose entire Dockerfile body is `COPY ${BIN} /server`: **a
two-line image that only copies a host-cross-compiled binary cannot be broken by an `app` version bump.**
All three died at `[internal] load metadata for docker.io/library/…`, before any layer ran. Measured
immediately after: the first `imagetools inspect` after Docker Desktop's boot took **26.9 s**; the next
three took **3 s** each. The daemon had been started ~90 s before the run (the machine was rebooted this
morning) and the bring-up put three concurrent metadata resolutions onto a cold connection. Environmental,
transient, **not evidence about the advance** — `D-M257x-58-3`, with `FIX-M257x-iter58-cold-daemon-registry`
routed. The 3-consecutive-green count restarts at zero: a failed cycle is not counted, skipped or resumed.

Then, from a fresh `down --purge` each time, verdicts read from **the bring-up's own `autoverify.json`**
(the iter-11 vantage lesson — never a standalone re-run):

```
refs (identical for all three cycles):
  platform:  0dab54dfac6beacdef54a671e2500d3940fd7329   (origin/main; re-fetched at open, level)
  app:       v1.366.0 (b948604f)                        (== app origin/main; ADVANCED this iter)
  rext:      fast-build-m257x-iter-58 (ab81527)         (tag verified ON ORIGIN; clone re-pinned)
  rosetta:   1937e1f + this iter's work
```

| cycle | teardown | verdict | timestamp |
|---|---|---|---|
| A | `rc=0`, survivors **0** | `{"project":"demo-1","offset":10000,"warnings":0,"green":true}` | `2026-08-04T08:03:34Z` |
| B | `rc=0`, survivors **0** | same, `warnings:0 green:true` | `2026-08-04T08:12:49Z` |
| C | `rc=0`, survivors **0** | same, `warnings:0 green:true` | `2026-08-04T08:22:37Z` |

**Three consecutive cold cycles, green, 0 warnings, distinct timestamps** (the M236 stale-verdict guard).
All 14 asserts pass on each, including `sentinel.casbin_rules = 1251`, `public.skills = 42790`,
directus per-stack-local, `demo-patches: all applied (none refused, none skipped)`, and
`frontend builds: ok (the running images are this run's)`.

**Clause 1 → MET at the advanced pins.**

**Pre-registration 2 was refuted, and by my own stale number rather than by the platform.** It predicted
**15** containers; the measured count is **11**, on all three cycles, with
`docker ps | grep -ci 'graphql\|router\|wunder\|cosmo'` → **0**. The half that mattered (*the router stays
deleted*) holds. The number was iter-14's, lifted from a summary instead of derived, and four services have
left the default profile since — `cms` and `jobsimulation` folded into `app`, `storage` to
`profiles: [storage-legacy]`, `messenger` dropped from `all`. 15 − 4 = 11, exactly. `D-M257x-58-4`.

## Phase D — clause 2, and the rider

```
refs:      as Phase C, on the cycle-C stack
command:   playthroughs/e2e/run-playthroughs.sh 1 --reset
verdict:   {"failing": 0, "passing": 30, "unimplementable-without-platform-edit": 0, "unimplemented": 1}
           total 31, coverage 0.9677
```

**`passing=30 failing=0`.** The single `unimplemented` is the declared `will-not-build` **verdict** on
`onboarding.enterprise-workforce-standard.UC1` (its only advancing path scrapes a live third-party profile,
so its RED would read as a product regression) — the machine-checked verdict M256 landed, not a gap.

**Clause 2 → MET at the advanced pins.**

### The rider — `FIX-M257x-iter56-assignment-flake` — and it is NOT DECIDED

iter-56 routed it with a specific instruction: *"Measure first whether app `850917d7` bears on it."*
`850917d7 fix(assignments): scope the join fall-through and tighten the already-member match` **is** in the
`v1.365.0..v1.366.0` range, so this iteration could answer it for free. Pre-registration 4 predicted **no**
— the flake is a test-side settle race (`toBe(before - 1)` over a baseline sampled while the members grid
is still settling; iter-56 observed `16 → 14`), which a server-side join fix does not touch.

**Observed: `pt-assignment-assign` passed every time it was run at `v1.366.0`** — in the full suite (4.4 s)
and in each repeat.

**That does not decide it, and recording it as if it did would be the milestone's own worst habit.** A flake
that manifests occasionally passes repeatedly under *both* hypotheses. Deciding it needs a **failure-rate**
at each version, and nobody has one at either — iter-56 has a single observed failure at `v1.365.0`, this
iter has a handful of passes at `v1.366.0`, and those two facts are compatible with "fixed" and with
"unchanged." **Pre-registration 4: not decided. The route stands, sharpened.**

**The repeat measurement also broke its own instrument, which is the more useful finding.** Running
`run-playthroughs.sh --grep pt-assignment-assign` makes the reporter mark **every other declared
Playthrough** `failing — no test outcome … (the declared Playthrough did not run / was not found)` and the
no-regressions gate FAIL. That is the M236 fail-closed rule (*an empty ledger is a FAILURE, not a 0/0
pass*) working exactly as designed — and it means **the suite has no supported way to run a subset and get
a usable verdict**, which is precisely what a flake investigation needs. Routed as
`FIX-M257x-iter58-grep-vs-failclosed`. The full suite was then re-run so the on-disk `last-report.json`
matches this section's claim rather than the `--grep` run — the `CHECK-M257x-iter56-stale-autoverify-twin`
class, avoided rather than created.

## Phase D (cont.) — the guard re-run, and what the advance induced

All four platform-alignment guards were re-run from the repo root against the real corpus.
**`anchor_construct_guard` went RED**, unaided, on a file nobody had touched:
`corpus/services/storage.md:9` cites `app/main.go:983`, **which is now a bare `}`**.

Traced: the cited call `storage.NewClient(…, storagens.CMS)` was at `:988` at `v1.365.0` and is at `:992`
at `v1.366.0` — the `clerkEventsManager` block move nets **+4 lines** for everything below ~452.
**Repaired minimally** (`:983 → :992`, one number, pointing at the construct the prose quotes rather than
at the comment above it) and watched **GREEN**.

Then the class was measured rather than assumed (`instrument/measure_mainline_shift.py`, committed per P2):

| | |
|---|---|
| `main.go:N` citations across `corpus/**`, `.claude/**`, `CLAUDE.md` | **23** |
| landing on different content at `v1.366.0` than at `v1.365.0` | **22** |
| caught by `anchor_construct_guard` | **1** — a **4.5% catch rate** |

**`FIX-M257x-iter57-within-block-drift` and `CHECK-M257x-iter57-anchor-guard-bare-class` are both
confirmed and sized.** iter-57 predicted the blind spot from structure and could not measure it; it is
**21 of 22**. The fence fires only when a cited line is not a construct *at all* — `storage.md:9` was
caught because its new line is a bare `}`. The other 21 land on comments and statements that still look
like good anchors; `app/main.go:1196-1202` now opens on `mux.Handle(skillerv1connect…` instead of the
messenger/cms comment it was chosen for, **in six separate files**.

**MOVED is a lower bound on instability, not a count of newly-false claims** — an unknown share was already
stale from iter-56's 37-commit advance, which nobody re-checked the corpus against. Full reasoning, and why
the advance nonetheless stays, in `D-M257x-58-5`.

## Close — 2026-08-04

**Outcome:** both deliberately-held-back pins are **advanced and proven cold** — `app v1.365.0 → v1.366.0`
and rext `iter-56 → iter-58` — with clauses 1 and 2 re-established at refs that are, for the first time in
this milestone, **current on every axis at once**. The advance then did what no reading had managed: it
**moved the platform under the corpus while an instrument was watching**, and the fence caught a site
unaided. Measuring what it caught produced the iteration's real finding — **22 of 23 `main.go` citations
moved and the fence saw 1 of them**, sizing iter-57's predicted blind spot at 21 of 22.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue

**Gate reading at close, against platform `0dab54d`** (P3 re-checked at close 08:30:23Z — platform origin
**unchanged**, app origin **unchanged at `v1.366.0`**, so the measurements stand rather than being
invalidated by construction):

| clause | reading | basis |
|---|---|---|
| 1 — 3 cold cycles green | **MET** | **re-proven this iter at the advanced pins**: 08:03:34Z / 08:12:49Z / 08:22:37Z, `warnings:0 green:true`, distinct timestamps, each from a fresh `down --purge` with survivors 0 |
| 2 — full Playthrough suite | **MET** | **re-proven this iter at the advanced pins**: `passing=30 failing=0 unimplemented=1` (the declared `will-not-build` verdict), reproduced twice |
| 3 — the migration-status map | **MET** | guards re-run: `platform_alignment_guard` OK (assertion F resolved 49 citations, 0 unresolvable) |
| 4 — zero writes to a dropped schema | **MET** | unchanged; guards re-run rc=0 |
| 5 — KB-fidelity | **NOT MET** | net **negative** this iter — 1 repaired, up to 21 newly falsified by the advance |

**4 of 5**, unchanged in count. What changed is that clauses 1 and 2 are no longer stale-by-construction
under the gate's own *"against origin HEAD"* wording.

**Decisions:** `D-M257x-58-1` … `D-M257x-58-5` (`iter-58/decisions.md`)

**Side-deliverables:** none. Every edit was inside the planned four-line scope; the one corpus edit
(`storage.md:9`) is the Phase-D guard repair, not an unrelated fix.

**Routes carried forward:**
- **`FIX-M257x-iter58-mainline-shift`** — the 21 remaining moved citations, as **ONE derived class** (TOK-04
  change 3). Shared predicate: *a line citation into `app/main.go` below ~452*. Shared remedy: **cite the
  construct, resolve the line at check time.** `instrument/measure_mainline_shift.py` is the seed. The
  strategic finding to carry with it: **a line-number citation into a fast-moving file is a claim with a
  half-life measured in days.**
- `FIX-M257x-iter58-cold-daemon-registry` — a bring-up started shortly after the Docker daemon dies ~4 min
  in with `DeadlineExceeded` and no mention of a registry. One warm-up `imagetools inspect` in the
  pre-flight both warms the path and fails early with a message that names the cause. **Pair it with
  `FIX-M257x-iter56-preflight-fails-late`** — same class: preconditions validated late or not at all.
- `FIX-M257x-iter58-grep-vs-failclosed` — `run-playthroughs.sh --grep` makes the fail-closed reporter mark
  every non-selected Playthrough `failing`, so **the suite has no supported way to run a subset and get a
  usable verdict**. That is exactly what a flake investigation needs.
- `FIX-M257x-iter58-empty-stdout-class` — §5 rules 1/2 are worded about *searches*; the class is **any
  command whose failure mode is empty stdout** (`D-M257x-58-1`).
- `CHECK-M257x-iter58-derive-preregistrations` — pre-registered quantities are prose in the one document
  that is supposed to be refutable (`D-M257x-58-4`).
- `FIX-M257x-iter58-baseline-refs` — the milestone's carried test baselines have no provenance line
  (`D-M257x-58-2`). **`stack-core` is 1F / 610**, not 599.
- `FIX-M257x-iter56-assignment-flake` — **NOT DECIDED**, sharpened: deciding it needs a **failure rate** at
  each app version, and nobody has one at either.
- Unchanged and still open: `FIX-M257x-iter57-within-block-drift` (now **sized**: 21 of 22),
  `CHECK-M257x-iter57-anchor-guard-bare-class` (now **sized**), `FIX-M257x-iter56-preflight-fails-late`,
  `FIX-M257x-iter56-evidence-gitignore`, `CHECK-M257x-iter56-directus-race-uncertified`,
  `CHECK-M257x-iter56-stale-autoverify-twin`, `FIX-M257x-iter55-stranded-demopatch-revert`, the 81 drift
  sites / 21 files, `FIX-M257x-iter53-union-set` (**a user decision, `D-M257x-53-5`**),
  `FENCE-M257x-iter54-refs-block`, `CHECK-M257x-iter52-second-ai-manager`, RF-2/3/7–13, root `CLAUDE.md`,
  `CHECK-M257x-iter38-ai-act-classification`.

**Lessons:**
1. **Schema-safety and citation-safety are unrelated properties, and only the first was ever measured.**
   An advance with 0 migrations, 0 destructive DDL and 0 removed contract — the shape this protocol calls
   safest, and correctly — moved 22 of 23 `main.go` citations. §7 rule 4's checklist (*what does the
   advance contain?*) asks only about the runtime. **It should also ask what the advance moves in the
   files the corpus cites.** That is a one-line addition to a checklist that has otherwise held up well.
2. **Three inherited numbers failed in one iteration, and the third was written by this iteration.** A
   container count read through a dead daemon; a test count read from before the commit it described; and
   a pre-registration I copied from a summary rather than derived. §5 rule 20's *"the author of a
   newly-written rule violated it while writing it"* now has a third instance — recorded by the author who
   had just written the rule down twice. The remedy is not care; it is P4: **all three quantities were
   derivable, and the derived one (autoverify's `all 11 expected container(s)`) was right.**
3. **A green bring-up after a red one is not the same evidence as three greens.** Cycle 1 died on
   `DeadlineExceeded`, and the available reading — *"we advanced `app` and the build broke"* — was wrong.
   What settled it in one step was noticing that one of the three failures was a **two-line Dockerfile
   whose entire body is `COPY ${BIN} /server`**: an image with no build in it cannot be broken by a
   version bump. **Look for the failing case that is too simple to be guilty.**
4. **A single pass of a flaky test decides nothing, and reporting it as a result is the habit this
   milestone keeps catching.** `pt-assignment-assign` passed every run at `v1.366.0`; that is compatible
   with "fixed" and with "unchanged". The instrument built to settle it — `--grep` — then broke, in a way
   that is the fail-closed rule working correctly. **Pre-registration 4: not decided.**
