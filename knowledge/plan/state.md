---
active_release: "v2.3 cue to cue — the presenter-speed release (designed 2026-07-13)"
active_branch: "release/02.30-cue-to-cue"
active_milestone: "(between milestones — v2.3 ALL 5 milestones M217-M221 CLOSED + merged into release/02.30-cue-to-cue). Next: /developer-kit:close-release (release->main merge + v2.3 tag)."
last_closed: "M221 — 2026-07-15"
phase: "M221 CLOSED (closed-on-gate, merged --no-ff into release/02.30-cue-to-cue) — the FINAL milestone of v2.3. All 5 milestones done. NEXT: /developer-kit:close-release tags v2.3. The billion demo is LEFT LIVE (intentional deliverable)."
last_updated: "2026-07-15"
---

# State

**Active release:** **v2.3 "cue to cue"** — the **presenter-speed release**. Designed 2026-07-13 via
`/developer-kit:design-roadmap`; branch `release/02.30-cue-to-cue` cut from `main`; **5 milestones M217 →
{ M218 ∥ M219 ∥ M220 } → M221**; tag will be `v2.3`. **Tooling + docs only — zero platform-repo edits.**

**Theme:** *a presenter swaps heroes in **under 5 seconds**, on a demo that comes up **green, fully-loaded, and
remotely reachable by default**.*

**Trigger — a live presenter defect** (user, 2026-07-13): *"I click a user, then it takes 1 or 2 minutes to access
the actual platform. … For a demo stack it is key that a presenter can swap from one hero to another with little
time."* The investigation found **Clerkenstein is INNOCENT**, that **the walls were already measured in this repo
and nobody looked**, and that **the corpus asserted the opposite in 4 places** (*"~2–5 s, which we can't
shorten"* — booked as M43-D5 with **zero deferrals recorded**, so it never entered a ledger across four releases).
Full gap analysis: [`.agentspace/scratch/roadmap-research-2026-07-13.md`](../../.agentspace/scratch/roadmap-research-2026-07-13.md).
Per-milestone closure narratives live in [`roadmap.md`](roadmap.md) — **not here**.

**Between milestones — v2.3 is CODE-COMPLETE. All 5 milestones (M217–M221) are closed and merged into
`release/02.30-cue-to-cue`.** The next lifecycle step is **`/developer-kit:close-release`** — the release-level
review + the `release/02.30-cue-to-cue → main` merge + the **`v2.3`** tag. No milestone is active.

> ### ✅ M221 CLOSED — the FINAL milestone: every v2.3 gate proven live on `billion`, over the tailnet, with NO FLAGS
> Gate **MET 8/8** on the iter-06 FINAL cold no-flag r4 cycle, browser-graded from a tailnet peer: **login p95
> maya-thriving 2.11 s / dan-manager 1.31 s** (both < 5 s, ACCESS 5/5); full catalog replayed (skills **42,790**,
> the F1 store-root **shadow** fixed live); 3 orgs; **Dana `/ai-readiness` 900-char browser check PASSED**; Ben
> STARTED; Aria COMPLETED; remote **default-on** no-flag; **0 platform edits**. M219 readiness fold-in all MET;
> F10 field-exercised; seed isolation CLEAN. **Reproducibly** = two independent cold r4 cycles (iter-05 + iter-06)
> at the same rext code, per the user's one-cycle pragmatic mandate. **The `billion` demo is LEFT LIVE** (cockpit
> `https://billion.taildc510.ts.net:17700`, app `:13000`) as an intentional final deliverable — not torn down.
> Final harden (2 passes → stabilized) caught its own D17: **`test_reap.py` printed "OK" after running only 21 of
> 41 tests** on a direct run (a mid-file `unittest.main()`) — fixed; **F-M221-06b landed** (`run-latency.sh`
> `LATENCY_SCHEME` for the HTTPS-fronted remote cockpit); an **F1 depth-2** edge pins that even a deeper empty-store
> shadow is loud-not-silent. rext code-of-record **`cue-to-cue-m221-final`** (live-graded at `-r4`). Deferral audit
> **YELLOW** — 4 non-gate tail carries (F4 academy-render + 3 live-infra probes) route to **v2.4**, sign-off at
> close-release.

> **⚠ THE RELEASE'S HEADLINE GATE was MET at M218 and RE-PROVEN at M221 live with no flags** — click→ACCESS **p95
> 2.11 s / 1.31 s** vs the **< 5000 ms** gate (from **39.45 s / 38.30 s** baseline — a ~18× improvement), on
> `billion` over the tailnet, cold reset-to-seed.

**Phase:** **M221 CLOSED** (`closed-on-gate`, merged `--no-ff`) — the FINAL milestone. **v2.3 is code-complete.**
**Next: `/developer-kit:close-release`** (release→main merge + the `v2.3` tag). Do NOT hand-merge/tag here.

**Next up:** **`/developer-kit:close-release`** — the release-level review, the `release/02.30-cue-to-cue → main`
merge, and the **`v2.3`** tag. It also owns the cross-release sign-off for M221's 4 tail carries → v2.4.

**Recently closed:** **M221** prove-on-billion — 2026-07-15 (`closed-on-gate`; 8/8 live on `billion`, no flags,
demo LEFT LIVE) · **M220** cue-sheet — 2026-07-15 (`closed-complete`; the "2 orgs" lie + remote opt-out +
a 5-milestone chronic resolved) · **M219** readiness-renders — 2026-07-14 (`closed-complete`; 13 bugs) · **M218**
seat-change — 2026-07-14 (`closed-on-gate`; the <5 s gate MET).

## D17 — the release's signature hazard (name it, then stop walking into it)

**D17: *a status artifact that outlives the thing it describes, and is then read as evidence.*** Named at the M218
close. It has now bitten **~24 times across M217→M221** — and it kept turning inward: **the fences kept catching
themselves.** M221's own instances: the M217 pre-bind reap **shipped but ran in zero tests**; the snapshot-cache
**shadow** ("cache miss" that was a wrong store-root); the exposure fence **blind to host-native listeners**;
`run-latency.sh`'s hardcoded `http://` cockpit URL (broke against M220's HTTPS-fronting); a `test_reap.py` that
printed **"Ran 21 tests … OK"** while silently omitting 20; and a **process** instance — the orchestrator
prematurely closed iter-05 on a **stale r3** snapshot while the agent was still re-proving at r4 (the host-lock
stopped DATA corruption, not BOOKKEEPING corruption; reconciled loudly in `3c64af1`). The **17-mutant battery M220 *claimed but never committed*** was itself a D17 (a result
surviving only as a sentence in a progress note); once committed it found the milestone's own **HARD INVARIANT
asserted against a re-typed COPY of itself** (mutating the shipped `SCHEME` to `https` left all 23 tests green).
The **image cache** matched a stale image so a new demo-patch **never baked while grading green**. **Five "reuse"
tests exercised the REBUILD path.** And the deepest: the dev-stack suite's *"environmental"* excuse **sat in this
file's own headline numbers for four releases**, re-characterised by every milestone that touched it — **one
missing env var underneath** (M220-D31).

**The generalized lessons:** ***"the doc says N" ≠ "the code ships N"*** · ***"it serves" ≠ "it renders" ≠ "the
session survives"*** · ***"the flag exists" ≠ "the flag works"*** · ***"the test is named X" ≠ "the test tests
X"*** · ***"the binary is installed" ≠ "it mints a cert"*** · ***"tailscaled handed me a string" ≠ "it is a
hostname"*** · ***an errored command is not "zero results"*** · ***a uniform result across unrelated mutations is a
constant, not a result*** · ***"we deferred this before" is not a reason to defer it again — a re-characterised
excuse is where a chronic bug hides.***

## User decisions taken at design time (binding)

| # | Decision |
|---|----------|
| **D-DESIGN-1** | **The <5 s gate is on ACCESS, not full first-page render.** ACCESS := the authenticated shell is rendered + interactive with the hero's identity present. The 200-member grid's data-load is **reported, never gated**. |
| **D-DESIGN-2** | **Fix properly via rext/Clerkenstein first; a new sha-pinned demo-patch is allowed if genuinely required.** A **platform-repo edit is NEVER in bounds.** |
| **D-DESIGN-3** | **Remote access: OPT-OUT for `/demo-up`, OPT-IN for `/dev-up`.** **Supersedes v2.2's D-DESIGN-1** for the demo path only. Consumes the reserved **M216** → folded into **M220(d)** as a declared scope-flex lever. |
| **D-DESIGN-4** | **The three story orgs are the three that already exist** (ai-transformation / sales-ramp / ai-readiness). **There is no hiring org and none will be built.** |

## Milestones

| # | Name | Shape | Complexity | Depends on |
|---|------|-------|------------|------------|
| ~~**M217**~~ | ✅ **DONE** — Clean stage: a demo that comes up **green** | `section` | medium | — |
| ~~**M218**~~ | ✅ **DONE** — Seat change: click→ACCESS **2413/1767 ms** vs < 5 s | `iterative` | large | M217 |
| ~~**M219**~~ | ✅ **DONE** — Readiness renders: **filled**, on the **current** surfaces | `section` | medium | M217 |
| ~~**M220**~~ | ✅ **DONE** — Cue sheet: the "2 orgs" lie + remote **opt-out** + `safety.md` **Part 3** | `section` | medium | M217 |
| ~~**M221**~~ | ✅ **DONE** — Prove it on billion: **8/8 live**, over the tailnet, **no flags**, demo LEFT LIVE | `iterative` | large | M217–M220 |

## Headline numbers (M221 close, 2026-07-15)
- **p95 click→ACCESS (the release's headline gate, set at M218 — RE-PROVEN live at M221 with NO flags):**
  **2.11 s** (employee `maya-thriving`) / **1.31 s** (manager `dan-manager`) vs the **< 5000 ms** gate, on the
  iter-06 cold r4 cycle over the tailnet HTTPS origin (reproduced on iter-05's r4 cycle). **Baseline 39.45 s /
  38.30 s.** *A latency number without its environment is not a measurement:* `billion`, Linux VM, 7.3 GiB RAM,
  tailnet origin.
- **Python tests:** **1341** (0 fail, 16 skip) — demo-stack 663 · stack-injection 260 · stack-core 182 ·
  dev-stack 122 · stack-verify 114. Counted from **JUnit XML**, never grepped stdout. **M221-attributable: +96**
  across the 3 touched sections (demo-stack +80, stack-injection +14, dev-stack +2). **stack-core re-counts to 182
  (M220 recorded 152)** — a pre-existing count drift *surfaced* by the fresh reconciliation; M221 did not touch it.
- **Go test funcs:** **1831** (+4 vs M220's 1827, **same method** — the 4 F1 store-resolver tests). **0 failures
  across all 6 modules; `go vet` clean.**
- **Fences graded by MUTATION** — M220's **17-mutant battery is now COMMITTED** (S3/S7 claimed 12+9 but committed
  neither; that was itself a D17). 17 mutants, 17 RED, with three anti-theatre assertions (baseline-green /
  mutants-actually-change-the-file / signatures-not-a-constant). *A fence that passes against the bug is theatre —
  three of M220's were, caught only by the mutation run.*
- **Alignment (Clerkenstein Go surface):** **100% / 100% critical (27/27)** — MET; re-run at M220 S6 after
  touching the FAPI (clerk-js from disk).
- **Demo-patches:** **8** (M220 added `next-web-no-thirdparty`). **Platform-repo edits:** **0**.
- **Flake:** **0** (5 sequential randomized runs).

## Branch model / shipped tags
**v2.3 CODE-COMPLETE (awaiting release close):** `release/02.30-cue-to-cue` cut from `main` at
`/developer-kit:design-roadmap` (2026-07-13); **all 5 milestone branches `m217/clean-stage` … `m221/prove-on-billion`
merged `--no-ff` and deleted.** The `release → main` merge + the `v2.3` tag are `/developer-kit:close-release`'s.
**Shipped tags:** **v2.2** `v2.2` · **v2.1** `v2.1` · **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` ·
**v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3**
`v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`. (Full detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog (unscheduled, cross-release)
- **The v2.2 residuals are FOLDED INTO v2.3 (Fate-1) — ALL DISCHARGED.** DEF-M215-01 → M217 · DEF-M215-04 → M217 ·
  **DEF-M215-03(b)** → M218 built it, **M221 ran it remotely** (the committed remote-origin latency gate) ·
  **`x/crypto@v0.52.0`** → LANDED at the M218 close · **DEF-M215-02** (remote-VM snapshot cache) → **DISCHARGED at
  M221** (the F1 store-root **shadow** fix — `public.skills 0 → 42,790` live on `billion`).
- **M221 tail carries → v2.4 (non-gate; sign-off owed at `/developer-kit:close-release`):** **F4** (academy grid
  renders 0 cards — fix is in the `ant-academy` **platform repo**, out of v2.3's zero-platform-edit scope) ·
  **BURNIN-M221-dev-public-host** (dev-path live burn-in) · **F-M220-4** (academy re-run on a live public-host
  demo) · **PROBE-M218-c3-rerun** (router-403 re-check, needs the box).
- ~~**M216** (dev-path Tailscale parity)~~ — ✅ **CONSUMED AND DISCHARGED by M220 S7** (2026-07-14, D28). The
  declared scope-flex lever was **NOT pulled**: S3 had already built + fenced the 6-rung ladder, so S7 was the
  thin wiring job the plan hoped for (a flag, a default, cross-section reuse — the ladder is *reused*, never
  forked). `/dev-up` now has `--public-host <host>|auto` (+ `DEV_PUBLIC_HOST`), **opt-in**, with a no-flag path
  proven **byte-identical** (zero `tailscale` invocations, tripwire-fenced). **M216 is retired as a
  reservation, never built as a milestone** — exactly as `roadmap-vision.md` said it would be.
- **Plan hygiene → next close-release:** `metrics-history.md` has **no v2.2 row**; no release-scope deferral audit
  exists for v2.1 or v2.2.
- **Older, still unscheduled:** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01 (`replayCmd` hermetic
  test), CAVEAT-1 (clean-box literal full `/dev-up`), M314b (prod frozen-read hydration — a prod-team follow-up).
  Reserved **Playthroughs futures** M205–M207 stay in vision. All tracked in
  [`roadmap-vision.md`](roadmap-vision.md).

_Last updated: 2026-07-15 (M221 close — the FINAL milestone. Gate Outcome Ledger + retro + metrics written; roadmap
flipped to `done`; merged `--no-ff` into `release/02.30-cue-to-cue`; rext tag `cue-to-cue-m221-final`. Final harden
2 passes → stabilized (false-green `test_reap.py` fix, F-M221-06b landed, F1 depth-2 edge). Deferral audit YELLOW,
0 blocking; 4 non-gate tail carries → v2.4 (sign-off at close-release). The `billion` demo is LEFT LIVE. **v2.3 is
code-complete — NEXT: `/developer-kit:close-release`** for the release→main merge + the `v2.3` tag.)_
