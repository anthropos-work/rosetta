---
active_release: "v2.3 cue to cue — the presenter-speed release (designed 2026-07-13)"
active_branch: "release/02.30-cue-to-cue"
active_milestone: "(between milestones) — M218 CLOSED; M219 ∥ M220 are next (parallel), then M221"
last_closed: "M218 — 2026-07-14"
phase: "M218 merged into release/02.30-cue-to-cue — next: /developer-kit:build-milestone on M219 ∥ M220, then M221 (last)"
last_updated: "2026-07-14"
---

# State

**Active release:** **v2.3 "cue to cue"** — the **presenter-speed release**. Designed 2026-07-13 via
`/developer-kit:design-roadmap`; branch `release/02.30-cue-to-cue` cut from `main`; **5 milestones M217 →
{ M218 ∥ M219 ∥ M220 } → M221**; tag will be `v2.3`. **Tooling + docs only — zero platform-repo edits.**

**Theme:** *a presenter swaps heroes in **under 5 seconds**, on a demo that comes up **green, fully-loaded, and
remotely reachable by default**.*

**Trigger — a live presenter defect** (user, 2026-07-13): *"I click a user, then it takes 1 or 2 minutes to access
the actual platform. Once logged in it works normally. For a demo stack it is key that a presenter can swap from one
hero to another with little time."*

**What the investigation found** (10-agent workflow + a dedicated residual-Clerk audit; full gap analysis, every
claim file:line-cited, at [`.agentspace/scratch/roadmap-research-2026-07-13.md`](../../.agentspace/scratch/roadmap-research-2026-07-13.md)):

1. **Clerkenstein is INNOCENT.** The cockpit CTA is a plain `<a href>`; the FAPI handshake is an in-memory mint + a
   303 with **zero I/O**; `grep time.Sleep` across `clerkenstein/` → **0 hits**. A cold Next.js compile is also
   REFUTED (the demo runs a **production** build). **Every surviving suspect lives AFTER the 303.**
2. **The walls were already measured, in this repo, and nobody looked** — 76.7 s members grid, 84 s router
   max-latency, a 180 s AI-readiness read that never completes (`stack-verify/e2e/lib/section-assert.ts:63-73`;
   `corpus/ops/demo/coverage-protocol.md:203,411,471`).
3. **The corpus asserts the opposite in 4 places** — *"~2–5 s, which we can't shorten"* — booked as **M43-D5** with
   **zero deferrals recorded**, so it never entered a ledger and was never revisited across v1.10b/v2.0/v2.1/v2.2.
   **v2.3 formally re-opens it.**
4. **Two `app` perf demo-patches silently REFUSE on sha-drift, on every run** (pinned @ app v1.295/v1.315; the box
   runs **v1.337**) — **with the refusal reason piped to `/dev/null`** (`up-injected.sh:701,717`).
5. **The last real run on `billion` was not a valid measurement** — the cockpit **crashed** on a leaked port and the
   bring-up **logged "serving" anyway**, so the operator drove a **stale predecessor**; 2 of 3 snapshot replays were
   cache misses (directus was rc=4, a *different* fault); autoverify **FAILED** with the failing probe's identity
   **discarded**; `jobsimulation` **exits(1)**. **M217 is a hard barrier: no number taken before it lands is
   trustworthy.**

**⚠️ M217's KB-fidelity gate came back RED (2026-07-13) — and it was right.** 14 load-bearing stale claims, **three
inside the milestone's own overview**. The worst: the drafted `jobsimulation` fix (`command: serve`) would have
**actively broken the service** — its cobra **root `RunE` IS the server**; the real cause is a `$HOME/.aws/credentials`
bind that Docker auto-creates as an empty **directory**, hard-erroring the AWS SDK and making cobra print its usage
block. Also corrected: the stale cockpit carries **no dead clerk-ids** (that mechanism does not exist), and **two**,
not three, replays were cache misses. All three corrected; gate cleared to **YELLOW**. Report:
[`releases/02.30-cue-to-cue/m217-clean-stage/kb-fidelity-audit.md`](releases/02.30-cue-to-cue/m217-clean-stage/kb-fidelity-audit.md)
— **its §5 is the ground truth the build works from, not the corpus docs.**

**Live finding:** the `/demo-down` run on `billion` earlier that day **left an orphaned cockpit alive** (pid 83214,
`0.0.0.0:17700`) — an unauthenticated hero-vending panel pointing at a deleted database. Killed manually. That is
**S2's defect, caught live**: teardown reaps by PID only, discards `kill`'s status, and prints success regardless.

**Active milestone:** **(between milestones).** **M218 "seat change" CLOSED 2026-07-14** — merged `--no-ff`
into `release/02.30-cue-to-cue`. rext code-of-record: **`cue-to-cue-m218`**. **`closed-on-gate`.**

> ### ✅ THE HEADLINE GATE OF THE RELEASE IS MET — click→ACCESS **2413 ms / 1767 ms** vs **< 5000 ms**
> Worst-case p95 over **5 GENUINE cold reset-to-seed cycles**, both vantages, **50/50** ACCESS, on `billion`
> over the tailnet. **From 39.45 s / 38.30 s — a ~16× improvement.** Two root causes, **both in demo tooling,
> zero platform edits**: (1) next-web's SSR GraphQL origin was a **build-inlined PUBLIC url that blackholes
> from inside the container** (~37.5 s); (2) Clerkenstein's fake BAPI served **a hardcoded STUB user to every
> hero**, so `app` refused `userPreferences` and a retry ladder burned ~6 s per render.
>
> **⚠ The gate number is 2413 ms — NOT the 1456 ms iter-04 reported.** That was measured on a **warm DB**:
> **F-9**, `/demo-down --purge` had **never purged on any Linux host**, so every "cold" bring-up for days
> reused the same database and images. The gate was **ungradeable** until that was fixed. Records:
> [`releases/02.30-cue-to-cue/m218-seat-change/`](releases/02.30-cue-to-cue/m218-seat-change/).

> **⚠ Alignment now reads 97.2%, not 100% — deliberately.** A **permanently-visible RED gene**
> (`MembershipOrgIdentity/real-org-eid`, **F-11/D16**) names a real divergence: the fake BAPI fabricates the
> org's external id. It was kept red rather than hidden by omitting the field — **a silently-omitted field is
> exactly how the user-level stub survived four releases.** Critical is **100%**; the gate (≥95/=100) is
> **MET**. → **`FIX-M219-bapi-org-eid`** (needs a runtime change **+ a fresh 5-cycle battery**).

> **⚠ Carried into M221 (from M217):** the **pre-bind reap has still never run live** — it was dead code
> (unsourced `reap.sh`, exit 127 swallowed) *during M217's green proof run on `billion`*. Fixed and
> unit-proven; **not field-proven**. M221 re-proves everything on the box.

**Phase:** **M218 closed + merged — next: `/developer-kit:build-milestone` on M219 ∥ M220** (parallel), then
**M221** (last).

**Next up:** **{ M219 ∥ M220 }** in parallel, then **M221 "prove it on billion"** — which re-proves **every**
gate on the VM, over the tailnet, **with no flags**. M218 merged first of the three parallel milestones (it
and M220 both touch `up-injected.sh`).

**Recently closed:** **M218** seat-change — 2026-07-14 (`closed-on-gate`; 21 bugs) · **M217** clean-stage —
2026-07-13 (gate MET on `billion`; 32 bugs).

## User decisions taken at design time (binding)

| # | Decision |
|---|----------|
| **D-DESIGN-1** | **The <5 s gate is on ACCESS, not full first-page render.** ACCESS := the authenticated shell is rendered + interactive with the hero's identity present. The 200-member grid's data-load is **reported, never gated** — this sidesteps the platform-side DataLoader defect instead of fighting it. |
| **D-DESIGN-2** | **Fix properly via rext/Clerkenstein first; a new sha-pinned demo-patch is allowed if genuinely required.** A **platform-repo edit is NEVER in bounds.** |
| **D-DESIGN-3** | **Remote access: OPT-OUT for `/demo-up`, OPT-IN for `/dev-up`.** **Supersedes v2.2's D-DESIGN-1** (*"public reach is never default-on"*) **for the demo path only**. Consumes the reserved **M216** (the dev-side flag does not exist today) → folded into **M220(d)** as a declared scope-flex lever. |
| **D-DESIGN-4** | **The three story orgs are the three that already exist** (ai-transformation / sales-ramp / ai-readiness). **There is no hiring org and none will be built** — it would need unmapped domain tables **+ the `hiring-app` frontend, which is not in the demo UI tier**. A separate release (reserved M205) if it ever revives. |

## Milestones

| # | Name | Shape | Complexity | Depends on |
|---|------|-------|------------|------------|
| ~~**M217**~~ | ✅ **DONE** — Clean stage: a demo that comes up **green** | `section` | medium | — |
| ~~**M218**~~ | ✅ **DONE** — Seat change: click→ACCESS **2413/1767 ms** vs < 5 s | `iterative` | large | M217 |
| **M219** | Readiness renders — the AI-readiness story is **visible** (the seed is a **verified no-op**) | `section` | medium | M217 |
| **M220** | Cue sheet — `/demo-up` defaults: the doc fix + remote **opt-out** + `safety.md` **Part 3** | `section` | medium | M217 |
| **M221** | Prove it on billion — every gate, on the VM, over the tailnet, **no flags** | `iterative` | large | M217–M220 |

## Headline numbers (M218 close, 2026-07-14)
- **NEW — p95 click→ACCESS:** **2413 ms** (employee) / **1767 ms** (manager) vs the **< 5000 ms** gate, over
  **5 genuine cold reset-to-seed cycles**, 50/50 ACCESS. **Baseline 39.45 s / 38.30 s.** *A latency number
  without its environment is not a measurement:* `billion`, Linux VM, 7.3 GiB RAM, over the tailnet origin.
- **Python tests:** **887** (0 fail, 12 skip) — demo-stack 495 · stack-injection 186 · stack-verify 109 ·
  stack-core 97. Counted from **JUnit XML**, never grepped stdout.
- **Go test funcs:** **1784** (+34 vs M217's 1750, **by the same method** — pinned in `metrics.json` so
  future closes compare like with like). **0 failures across all 6 modules.**
- **Alignment (Clerkenstein Go surface):** **97.2% overall / 100% critical** (26/27) — gate ≥95/=100 ⇒ **MET**.
  The 2.8% is **one deliberately RED gene** (see above). **Only 4 of 5 surfaces are measurable** — `expressrun`
  is dependency-gated and exits **rc=2 with NO score** on a box without `@clerk/express` `node_modules`, which
  nothing treated as a failure (**pre-existing**; → `TEST-M219-expressrun-dep-gate`).
- **Flake:** **0** (5/5 clean, sequential). **Platform-repo edits:** **0**.
- **Known issue (flagged, not a regression):** the **`dev-stack` suite fails on this box** — environmental
  (needs a live Postgres on `:15432`; this box's `.agentspace/secrets` is also incomplete). Identical at v2.2
  and M217. **M218 does not touch `dev-stack/`.**

## Branch model / shipped tags
**v2.3 IN DEVELOPMENT:** `release/02.30-cue-to-cue` cut from `main` at `/developer-kit:design-roadmap` (2026-07-13);
milestone branches `m217/clean-stage` … `m221/prove-on-billion` branch from it.
**Shipped tags:** **v2.2** `v2.2` · **v2.1** `v2.1` · **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` ·
**v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3**
`v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`. (Full detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog (unscheduled, cross-release)
- **The v2.2 residuals are FOLDED INTO v2.3 (Fate-1).** ✅ **DISCHARGED:** DEF-M215-01 (app demopatch sha
  re-anchor) → M217 · DEF-M215-04 (`jobsimulation` exits(1)) → M217 · **DEF-M215-03(b)** (committed
  remote-origin Playwright gate) → **M218 built it** (the `stack-verify/e2e/` latency harness; M221 runs it
  remotely) · **`x/crypto@v0.52.0`** → **LANDED at the M218 close** (it was a *repeat + aged-out* deferral;
  the close **is** M218's rext roll — alignment gate scores identically pre/post ⇒ behaviour-neutral).
  **STILL OPEN:** DEF-M215-02 (remote-VM snapshot cache) → **M221**.
- **M216** (dev-path Tailscale parity) — **CONSUMED by M220(d)** per D-DESIGN-3. No longer a reservation; it is the
  release's declared scope-flex lever (it drops back to a reservation only if M220 bloats).
- **Plan hygiene → next close-release:** `metrics-history.md` has **no v2.2 row**; no release-scope deferral audit
  exists for v2.1 or v2.2.
- **Older, still unscheduled:** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01 (`replayCmd` hermetic
  test), CAVEAT-1 (clean-box literal full `/dev-up`), M314b (prod frozen-read hydration — a prod-team follow-up).
  Reserved **Playthroughs futures** M205–M207 stay in vision. All tracked in
  [`roadmap-vision.md`](roadmap-vision.md).

_Last updated: 2026-07-13 (`/developer-kit:design-roadmap` — v2.3 "cue to cue" designed; branch
`release/02.30-cue-to-cue` cut from `main`; 5 milestone dirs scaffolded under `releases/02.30-cue-to-cue/`; 4 binding
user decisions recorded. Deferral preflight: **YELLOW** — 4 open v2.2 residuals, all folded Fate-1 into this release,
none repeat-deferred. KB preflight: **6 blind areas** identified, each assigned a `Delivers →` owner.)_
