---
active_release: "v2.3 cue to cue — the presenter-speed release (designed 2026-07-13)"
active_branch: "release/02.30-cue-to-cue"
active_milestone: "M221 (prove-on-billion) — the LAST milestone. Iterative. Re-prove EVERY v2.3 gate on the remote VM over the tailnet with NO FLAGS: p95 click->ACCESS <5s both vantages, full replayed catalog, 3 orgs incl. AI-readiness, remote reach default-on. Inherits M217-M220 field routes (incl. the 4 M220 Fate-3 routes: 2 academy + dev burn-in)."
last_closed: "M220 — 2026-07-15"
phase: "M220 CLOSED (closed-complete, merged --no-ff into release/02.30-cue-to-cue). NEXT + LAST: M221. Do NOT merge to main / tag v2.3 until M221 closes."
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

**Active milestone:** **M221 "prove it on billion" — the LAST milestone** (`iterative`, large). Re-prove **every**
v2.3 gate on the **remote VM, over the tailnet, with NO FLAGS**: p95 click→ACCESS **< 5 s** both vantages, the full
replayed catalog (no skipped surface), all 3 orgs incl. AI-readiness, remote reach **default-on**. It inherits the
field routes from M217–M220 — including M220's **4 Fate-3 routes**: `FIX-M221-academy-empty-catalog`,
`FIX-M221-academy-loopback-bind` (+ extend the exposure fence to host-native listeners), `F-M220-4`
(academy re-runnable on a live public-host demo), `BURNIN-M221-dev-public-host`.

> ### ✅ M220 CLOSED — the user's ask was three things, and TWO were already true (the docs lied)
> *"Pull all data"* + *"seed 3 orgs"* were **already default-on**; the `"2 orgs"` lie (**11 sites, 4 releases**)
> is why the user believed the seeding ask was unmet — fixed + **fenced** (a doc-vs-code guard that counts the
> preset). The **one genuine flip** — remote reach **default-on**, opt-out `--no-public-host` — **landed (S3) and
> is live-proven** on `billion`, both vantages, on a trusted LE cert, cold reset-to-seed reproducible, 0 ejects.
> **The invariant that holds:** a box with **no** Tailscale is **byte-identical to localhost** (proven by
> genuinely removing `tailscale`, tripwire-fenced — `SCHEME`/`BIND_HOST` share one predicate, so a half-satisfied
> public path is strictly worse than localhost). Also: academy stops **poisoning the session** (S5); **zero
> egress** on an authed load (7 third parties killed, S6); clerk-js **from disk** (was unbounded CDN on the login
> path); cockpit **behind HTTPS** (S4); `safety.md` **Part 3 — the exposure axis** (S1). The dev-side **opt-in**
> `--public-host` **discharged M216** without pulling the scope-flex lever (S7). **studio-desk (j) was NOT a bug**
> — M219's *"302 → /login"* was a **cookieless curl**, mis-read (D15). **Zero platform-repo edits.**
>
> **Harden + close found their own defects, several in the tooling itself (D17):** the **17-mutant battery** M220
> *claimed but never committed* (unfalsifiable — the exact D17) is now a test, and it found the milestone's own
> **HARD INVARIANT unfenced** (the fence asserted a **re-typed copy** of the derivation; H-1) + an **untested
> `|| true`** (H-3). Close resolved a **5-milestone chronic deferral**: the dev-stack suite's *"environmental"*
> failure was **one missing env var** (D31) — so the **whole rext Python suite completes for the first time**
> (1215 tests, 0 fail). Adversarial: rung 3 accepted any dotted string as a hostname (D32). **Deferral audit
> YELLOW, 0 blocking, 0 escape-hatch.**

> **⚠ THE RELEASE'S HEADLINE GATE WAS MET AT M218 — click→ACCESS 2413 ms / 1767 ms vs < 5000 ms** (worst-case p95,
> **5 GENUINE cold reset-to-seed cycles**, both vantages, on `billion` over the tailnet). **From 39.45 s / 38.30 s
> — a ~16× improvement.** M221 re-proves it (and everything) live, **with no flags**.

> **⚠ Carried into M221:** the **pre-bind reap has never run live** (M217, dead code during the green proof run);
> **two disclosed M219 caveats** (the battery was not one uncontested consecutive run — an orchestration error;
> and the graded code ≠ the shipped code, `aiReadinessStep1Score` double-rounded → `REPROVE-M221-battery-at-final-code`);
> and M218's `PROBE-F7`/`C-3`. All settled only **on the box**, which is M221's whole job.

**Phase:** **M220 CLOSED** (closed-complete, merged `--no-ff`). **Next + LAST: M221.** Do **NOT** merge to `main`
or tag `v2.3` until M221 closes.

**Next up:** **M221 "prove it on billion"** (the last milestone) — re-prove **every** gate on the VM, over the
tailnet, **with no flags**, then `/developer-kit:close-release`.

**Recently closed:** **M220** cue-sheet — 2026-07-15 (`closed-complete`; the "2 orgs" lie + remote opt-out +
`safety.md` Part 3; a 5-milestone chronic deferral resolved) · **M219** readiness-renders — 2026-07-14
(`closed-complete`; 13 bugs) · **M218** seat-change — 2026-07-14 (`closed-on-gate`; the <5 s gate MET).

## D17 — the release's signature hazard (name it, then stop walking into it)

**D17: *a status artifact that outlives the thing it describes, and is then read as evidence.*** Named at the M218
close. It has now bitten **~20 times across M217→M220** — and in M220 the pattern turned inward: **the fences kept
catching themselves.** The **17-mutant battery M220 *claimed but never committed*** was itself a D17 (a result
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
| **M221** | Prove it on billion — every gate, on the VM, over the tailnet, **no flags** | `iterative` | large | M217–M220 |

## Headline numbers (M220 close, 2026-07-15)
- **p95 click→ACCESS (the release's headline gate, set at M218):** **2413 ms** (employee) / **1767 ms** (manager)
  vs the **< 5000 ms** gate, over **5 genuine cold reset-to-seed cycles**. **Baseline 39.45 s / 38.30 s.**
  *A latency number without its environment is not a measurement:* `billion`, Linux VM, 7.3 GiB RAM, tailnet origin.
  **M221 re-proves it live, with no flags.**
- **Python tests:** **1215** (0 fail, 16 skip, 39 subtests) — demo-stack 583 · stack-injection 246 · stack-core
  152 · **dev-stack 120** · stack-verify 114. Counted from **JUnit XML**, never grepped stdout. **The whole rext
  Python suite completes for the FIRST time** — the `dev-stack` suite was un-runnable (`rc=124`) for 5 milestones
  until M220-D31 (**one missing env var**, not "environmental"). Like-for-like on M219's 4 counted sections:
  **1095 vs 903 = +192**; dev-stack 120 now visible.
- **Go test funcs:** **1827** (+6 vs M219's 1821, **same method**). **0 failures across all 6 modules; `go vet`
  clean.**
- **Fences graded by MUTATION** — M220's **17-mutant battery is now COMMITTED** (S3/S7 claimed 12+9 but committed
  neither; that was itself a D17). 17 mutants, 17 RED, with three anti-theatre assertions (baseline-green /
  mutants-actually-change-the-file / signatures-not-a-constant). *A fence that passes against the bug is theatre —
  three of M220's were, caught only by the mutation run.*
- **Alignment (Clerkenstein Go surface):** **100% / 100% critical (27/27)** — MET; re-run at M220 S6 after
  touching the FAPI (clerk-js from disk).
- **Demo-patches:** **8** (M220 added `next-web-no-thirdparty`). **Platform-repo edits:** **0**.
- **Flake:** **0** (5 sequential randomized runs).

## Branch model / shipped tags
**v2.3 IN DEVELOPMENT:** `release/02.30-cue-to-cue` cut from `main` at `/developer-kit:design-roadmap` (2026-07-13);
milestone branches `m217/clean-stage` … `m221/prove-on-billion` branch from it.
**Shipped tags:** **v2.2** `v2.2` · **v2.1** `v2.1` · **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` ·
**v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3**
`v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`. (Full detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog (unscheduled, cross-release)
- **The v2.2 residuals are FOLDED INTO v2.3 (Fate-1).** ✅ **DISCHARGED:** DEF-M215-01 → M217 · DEF-M215-04 → M217 ·
  **DEF-M215-03(b)** → **M218 built it** (M221 runs it remotely) · **`x/crypto@v0.52.0`** → **LANDED at the M218
  close**. **STILL OPEN:** DEF-M215-02 (remote-VM snapshot cache) → **M221**.
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

_Last updated: 2026-07-15 (M220 close — retro + metrics + Completeness Ledger written; roadmap flipped to `done`;
merged `--no-ff` into `release/02.30-cue-to-cue`; rext tag `cue-to-cue-m220-final`. Deferral audit YELLOW, 0
blocking, 0 escape-hatch — the 5-milestone "environmental" dev-stack carry LANDED Fate 1 (one env var). 4 Fate-3
routes attached to M221's `overview.md` (EDITED) + the stale `devstack-test-spin` route retracted. **Next + LAST:
M221** — do NOT merge to `main` or tag `v2.3` until it closes.)_
