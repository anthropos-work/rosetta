# M256 — retro

**playthrough sharpening** · `iterative` · closed **2026-07-30**, `closed-on-gate` · 32 iters, 3 harden passes

## Summary

The suite was **18/18 green while the demo still had things that did not work**. That is not a paradox but a
structural property, and the milestone's job was to remove it: make the detector one you can trust.

It ends at **30 live Playthroughs + 1 verdicted TODO**, mutating **1 → 12**, negative controls **0 → 28 of 30**,
`blocked` outcomes **0 → 1**, **31/31** written verdicts with **0 `unimplementable`**, and two whole
products — **org-admin** and **onboarding**, the latter *the first thing every real user does* — that no e2e
suite had ever touched and that the milestone's own pre-flight audit had priced as impossible.

But the deliverable people should read this file for is not the count. It is **~43 checks that reported success
without having checked**, and the fact that **two of the fences shipped to catch that class were themselves
instances of it**.

## Incidents this cycle

| | what | disposition |
|---|---|---|
| **P2** | **A real 245 s suite flake** (iter-06) traced to an **unreachable retry loop** — an unbounded first `click` inheriting the whole test budget. | Fixed; regression-pinned in the bounded-interaction fence. |
| **P2** | **`FLAKE-M256-assign-under-bloated-policy`** (iter-13). Recurred on a Phase-D run. The stated hypothesis was **wrong** and two more plausible ones were refuted by probe; the trace named the real cause — the assign modal is **row-scoped**, so a members-table re-render unmounts it. Time decomposed exactly as the retry ladder's own bounds (3×15 + 20 + 15 = **84 s**). | Fixed three ways, recovery proven deterministically, 3× gate re-run 0 flake. **Diagnosed, not re-rolled.** |
| **P2** | **A batch of three reddened by `PLATFORM-M256-keyrole-nondeterminism`** — a succession key-role card appears on 4 of 5 loads at role occupancy 2. A 45 s timeout did **not** fix it, because the cause was seed occupancy, not the clock. | **Reported and diagnosed rather than re-rolled away. No timeout was bumped.** Mitigated seed-side (hero roles pairwise distinct); routed to the platform. |
| **P2** | **`demo-2-jobsimulation-1` and `demo-2-cms-1` self-terminated cleanly (`Exited 0`)** on their DB-health monitors after an un-clean Postgres restart, and nothing restarted them. `docker ps` showed 14/16 "Up", the app surfaced **no error**, and every jobsimulation surface rendered 20 content-free rows — which a `rows > 0` assertion passes. Cost **an hour of Playthrough diagnosis spent disbelieving a correct assertion.** | Fate 3 → **M257**, because its gate reads `autoverify green / 0 warnings` and would pass on exactly this stack. |
| **P3** | **The stack's PINNED `stackseed` predates M256's persona fields**, so `--reset` truncates the world and the re-seed then **fails**, leaving demo-2 EMPTY. Hit at this close on the first verification run. | Recovered by shadowing the authoring build on `PATH`. Harden-final had done the same **without writing the failure down** — now recorded here and in `progress.md`. |
| — | **Eight sub-agent deaths** (internet drop, session limit, stall watchdog, four API 529s). | **Zero work lost** — every iter committed and pushed **both** repos before the next began. |

## What went well

- **The gate was re-cut in the open, twice, against the measurement that falsified it.** Clause 1's `≤ 0.79×`
  was shown to sit **inside its own 2.04× noise floor** (six batches, one host, code untouched since iter-03).
  The flattering denominator was available — the original-16 subset read `0.7063×`, **inside** the gate — and
  was **refused** as hand-picked. Two earlier "clause 1 MET" readings were **retracted** as favourable samples.
- **Escalating worked, including when the escalation was wrong.** iter-20 refused to grant itself the
  permission whose enforcement was under test, and escalated with both options and the single missing fact. The
  answer refuted its own recommendation — the demo's seed was incomplete, not the platform. Had it guessed, it
  would have filed a platform bug that does not exist; had it granted itself the permission blind, it would
  have papered over a real fidelity defect.
- **Two refusals shipped as deliverables.** iter-18 declined to build an onboarding journey whose green depends
  on **scraping a real person's public profile** — flaky by construction, and its RED would read as an
  Anthropos regression. That became a machine-checked `will-not-build` verdict, and `ptreport` now renders it in
  place of a generic gap line. *A test whose failure lies is worse than no test.*
- **Nothing accumulated.** Zero standing red across 32 iters (D-v28-3), and every iter committed and pushed
  both repos before the next started — which is the entire reason eight agent deaths cost nothing.
- **Measurement beat argument, repeatedly.** Seven confident conclusions were refuted by driving the thing:
  three from iter-05, the studio false-green diagnosis (the blamed string **never renders**, so the routed fix
  was a **no-op that would have shipped as a fix**), and the coordinator's own `D99` mechanism.

## What didn't

- **The fences built to catch the class were instances of the class.** The **liveness fence counted
  `not.toBeVisible()` as proof the page was alive** — an absence assertion serving as the liveness witness,
  precisely what it exists to prevent — and the **bounded-interaction fence never scanned the retry loop it is
  named after** (D25's subject is a *counted* `for` loop; the pattern matched only `for(;;)`/`while(true)`, and
  was satisfied by an unrelated loop 200 lines away). Both were **latent**, which is why a fence that only runs
  over the current corpus could not find either.
- **A flat yield was mistaken for progress until it was counted.** 6 → 11 → 9 across three harden passes. A
  flat yield does not mean nearly done; it means **the seam is broad**. Naming that is what made
  `HARDEN-CAP-ACCEPTED-D105` an honest stop rather than a shrug.
- **The routing table became a place chronics went to be re-typed rather than re-decided.** The deferral audit
  returned **RED**: 8 rows had targets that no longer existed, 6 read OPEN for landed work, and the longest
  chronic was carried **~10 times across 25 iters** — twice promoted in prose to *"the ONLY thing standing
  between clause 2 and 25 of 25"* and still not landed. A ledger that reads open for closed work hides the
  ones that are genuinely open.
- **`PT-M256-resume-fixture-pair` outlived its own premise by ~23 iters.** It existed so two use cases could
  share the cost of one fixture. The fixture landed at iter-18 and the second use case became a verdict — and
  nobody re-read the row.
- **Two coordinator process gaps, named because they are mine.** (1) iter-21 was closed **by hand** during API
  529 spawn failures, and its running-ledger line was never written — found only by this close's iter-ledger
  audit. (2) The verification of that same iter's uncommitted work was `ptvalidate` + `tsc` + *"the spec is a
  real test"*, and it **missed gofmt dirt** and **missed that snapshot Phase 5 had no test call sites at all**.
  The only check that finds the second is the milestone's own thesis — ***delete it and see whether anything
  fails*** — and I did not apply it to the work I was verifying.
- **A source read was relayed as an observation.** iter-29's `file:line` reading reached the next iter as if it
  had been seen happening; iter-31 measured the surface and found the routed repair **impossible** (there is no
  Role screen). Booked as **D118**: *citing a `file:line` does not make it an observation.*

## The finding that changes how conclusions get drawn

**`grep` goes SILENT, not loud, on a file with a stray control byte.**

`tests/url-shapes.unit.spec.ts` carried a raw NUL and 0x01 inside an adversarial fuzz input. `file(1)` called
the file `data`; `grep(1)` treated it as **binary and reported no matches at all** — no error, no warning, no
change to its exit code. Node, Playwright and `tsc` were unaffected: its **79 tests ran and passed throughout**.

So the **largest unit pin in the harness — 43 KB, 79 tests — was invisible to every shell-based sweep over
`tests/`**, and invisible in the only way that cannot be noticed: as an *absence of hits*.

During this close, **two independent reviewers swept the harness with grep, both received silence for that
file, and both concluded that `looksLikeCurrentPosition` — a predicate the file imports and exercises by
name — was dead code with zero consumers.** One recommended deleting it.

**Every *"nothing references this, so it is dead"* conclusion in this project rests on grep reporting
honestly.** It usually does. When it does not, it fails in the direction that looks like a clean result. The
class is now fenced over `tests/` and `lib/` (the fence's first catch was its own docstring), but the
transferable rule is about method, not about that file: **a negative conclusion from a search tool needs the
denominator asserted** — how many files were read — exactly as this milestone's own fences learned to assert a
not-vacuous floor. An unasserted denominator is how a search that read nothing looks identical to a search
that found nothing.

## Carried forward

Full per-item table in `decisions.md` § *DEFERRAL GATE*; destinations are **named milestones**, not phrases —
M255's close routed four items to *"M255 harden resume"*, which was not a milestone, and its own retro says
that should have been rejected when written.

| destination | items |
|---|---|
| **M258** | the studio false-green bundle (fix + 2 withheld controls + the LLM-lane doc premise) · the **9 remaining standing mutants** · the **11 lower-severity harden-3 scan findings** (risk of needing re-fating there is stated in its `overview.md`) · `FIX-M257-content-stories-pair-count` · `ptvalidate` invoked nowhere (which also closes the runner gate's permissive half) · **`BIND_HOST` / `D-M255-7`**, whose M255 routing was declared and never applied |
| **M257** | `FIX-M256-demo2-service-self-termination` · `FIX-M256-autoverify-fapi-libressl` — both **gate-relevant**, not parked |
| **the platform** | `PLATFORM-M256-onboarding-step-not-resumed` (2 defects) · `DEFECT-M256-silent-forbidden-mutation` **+ its sweep residual, marked INFERRED not measured** · `PLATFORM-M256-keyrole-nondeterminism` · `PLATFORM-M256-cv-upload-never-parses` — all four now in the net-new [`platform-defect-register.md`](../../../platform-defect-register.md) with `file:line`, because there was **no register anywhere in this repo** and they would have archived with this milestone |
| **DROPPED** | `PT-M256-resume-fixture-pair` — premise dissolved |
| **awaiting the user's signature** | `PERF-M256-parallel-lane` · `PT-M257-self-evaluation` · `PT-M257-talk-to-data` — each a roadmap call, not a routing one |

## Metrics delta

From [`metrics.json`](metrics.json):

| | M255 close | M256 close |
|---|---|---|
| Go test funcs (rext) | 2 023 | **2 130** (+107) |
| Go modules failing | 0 of 6 | **0 of 6** (58 packages ok) |
| Python (same 3-suite scope) | 1 505 passed | **1 552** passed (+47) · four-suite total **1 723 / 2 skip** |
| Playwright | 204 in 42 files | **209 in 43 files** · unit 169 → **174** |
| Live Playthroughs | 18 | **30** (+1 verdicted TODO) |
| Mutating (mutate **and** read back) | 1 | **12** |
| Negative controls | 0 | **28 of 30** |
| `blocked` outcomes | 0 | **1** |
| Flake | 0 | **0** (3× cold reset-to-seed, rc `0/0/0`) |
| Platform-repo edits | 0 | **0** |
| Net-new deps | 0 | **0** |

**Environment stated with every number** (`D-v28-13`): `Kirality-Mac-Pro-6.local`, darwin 25.1.0, Docker VM
**9.70 GiB** against the documented 12 GB UI-tier floor; `demo-2` offset 20000, localhost/http,
`--no-public-host`. **Per `D-v28-12`, no number here may be quoted as comparable to billion's 228 s** — the
absolute re-measure is M258's.
