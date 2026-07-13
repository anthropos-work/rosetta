---
active_release: "v2.3 cue to cue — the presenter-speed release (designed 2026-07-13)"
active_branch: "release/02.30-cue-to-cue"
active_milestone: "(between milestones) — M217 CLOSED; M218 ∥ M219 ∥ M220 are next"
last_closed: "M217 clean-stage — 2026-07-13 (v2.3 cue to cue; exit gate MET on billion; 32 bugs)"
phase: "M217 merged into release/02.30-cue-to-cue — next: /developer-kit:build-milestone on M218 (or M219/M220 in parallel)"
last_updated: "2026-07-13"
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

**Active milestone:** **(between milestones).** **M217 "clean stage" CLOSED 2026-07-13** — merged
`--no-ff` into `release/02.30-cue-to-cue`. rext code-of-record: **`cue-to-cue-m217`**. **32 bugs fixed** (build 0
→ harden ×3: 24 → close: 8). **Exit gate MET on `billion`** (cold reset-to-seed: `autoverify: OK`, 3/3 replays
exit 0, 2/2 app patches applied [one self-healed], jobsimulation serving, content plane local). Records:
[`releases/02.30-cue-to-cue/m217-clean-stage/`](releases/02.30-cue-to-cue/m217-clean-stage/).

> **⚠ Carried into M221:** the **pre-bind reap has still never run live**. The close review found that
> `up-injected.sh` called `reap_port` **without sourcing `reap.sh`** — so the milestone's headline deliverable
> was dead code (exit 127, swallowed) *during the green proof run on `billion`*. It is fixed and unit-proven;
> it is **not field-proven**. M221 re-proves everything on the box.

**Phase:** **M217 closed + merged — next: `/developer-kit:build-milestone` on M218** (the headline latency
milestone), or **M219 ∥ M220** in parallel.

**Next up:** **{ M218 ∥ M219 ∥ M220 }** in parallel (M218 merges first of the three — it and M220 both touch
`up-injected.sh`), then **M221**. **M218 may now measure** — and `autoverify.json` is the signal it gates on, so it
can never again measure a broken stack.

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
| **M217** | Clean stage — a demo that comes up **green** (the hard barrier) | `section` | medium | — |
| **M218** | Seat change — **click→ACCESS < 5 s**, both vantages | `iterative` | large | M217 |
| **M219** | Readiness renders — the AI-readiness story is **visible** (the seed is a **verified no-op**) | `section` | medium | M217 |
| **M220** | Cue sheet — `/demo-up` defaults: the doc fix + remote **opt-out** + `safety.md` **Part 3** | `section` | medium | M217 |
| **M221** | Prove it on billion — every gate, on the VM, over the tailnet, **no flags** | `iterative` | large | M217–M220 |

## Headline numbers (M217 close, 2026-07-13)
- **Python tests:** **867** (0 fail, 11 skip) — demo-stack 462 · stack-injection 194 · stack-verify 109 ·
  stack-core 97 · dev-stack(new) 5. Counted from **JUnit XML**, never grepped stdout.
- **Go test funcs:** **1750** (+1 vs v2.2 **measured by the same method**, which gives 1749 at v2.2). *The 1772
  recorded at the v2.2 close used a DIFFERENT method and is not comparable — the method is now pinned in
  `metrics.json` so future closes compare like with like.*
- **Flake:** **0** (5/5 clean, sequential). **Platform-repo edits:** **0**.
- **Known issue (flagged, not a regression):** **33 pre-existing `dev-stack` failures** — environmental (they
  need a live Postgres on `:15432`), **identical count at v2.2**. M217 took that suite from 55 → 60 passing.
- **NEW metric this release:** p95 click→ACCESS latency — **not yet measured**; M218 iter-01 sets the baseline.


## Branch model / shipped tags
**v2.3 IN DEVELOPMENT:** `release/02.30-cue-to-cue` cut from `main` at `/developer-kit:design-roadmap` (2026-07-13);
milestone branches `m217/clean-stage` … `m221/prove-on-billion` branch from it.
**Shipped tags:** **v2.2** `v2.2` · **v2.1** `v2.1` · **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` ·
**v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3**
`v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`. (Full detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog (unscheduled, cross-release)
- **The v2.2 residuals are now FOLDED INTO v2.3 (Fate-1), not standing:** DEF-M215-01 (app demopatch sha re-anchor)
  → **M217**; DEF-M215-02 (remote-VM snapshot cache) → **M217 + M221**; DEF-M215-03(b) (committed remote-origin
  Playwright gate) → **M218 + M221**; DEF-M215-04 (`jobsimulation` exits(1)) → **M217**; the `x/crypto@v0.52.0` bump
  → **M218's rext roll**.
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
