---
active_release: "v2.3 cue to cue — the presenter-speed release (designed 2026-07-13)"
active_branch: "release/02.30-cue-to-cue"
active_milestone: "M220 (cue sheet) — IN PROGRESS on m220/cue-sheet. S0+S1+S2 DONE (docs/fences); S5+S6 DONE (proven on a live demo: the academy session-poisoning is FIXED — A/B green from a tailnet peer; zero egress; clerk-js from disk). S3, S4, S7 NOT started (each needs a live demo, run one at a time)"
last_closed: "M219 — 2026-07-14"
phase: "M220 S0–S2 + S5 + S6 landed on m220/cue-sheet. S5/S6 proven on a live billion demo (0 ejects, 0 egress, alignment 100%/100% both surfaces). NEXT: S3, S4, S7 — each needs a LIVE DEMO. Then M221 (last)."
last_updated: "2026-07-14"
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

**Active milestone:** **(between milestones).** **M219 "readiness renders" CLOSED 2026-07-14** — merged `--no-ff`
into `release/02.30-cue-to-cue` (`e7a33c6`). rext code-of-record: **`cue-to-cue-m219-final`**. **`closed-complete`.**

> ### ✅ M219 — THE USER'S KICKOFF BAR IS MET ON BOTH COUNTS
> *"each element and sub section of readiness **filled** spot data"* + *"the **right** dashboards/pages … **not the
> old legacy ones**"* — proven on **5 cold reset-to-seed rebuilds**, each independently evidenced. Manager: **all 8
> readiness sections PASS**. Employee: **both heroes** `failingSections=0`. **Zero** pointers resolve to a legacy
> surface. **Zero platform-repo edits.**
>
> **TWO of the milestone's OWN premises were REFUTED by measurement**, and **the planned demo-patch was WITHDRAWN,
> not built.** All 3 demo pointers had targeted an **UNLINKED ORPHAN** page; the employee surface **has no route of
> its own** (embedded in `/home`) — *which is why route-crawling never found it*. Junk skills **ORG-WIDE** (the
> claimed-tail pool ran dry and topped up from the taxonomy's **alphabetical head**) — **the flat tier is DELETED**.
> A hero rendered **ROLE-LESS**. The manager's **4 interview-findings blocks rendered HEADINGS OVER NOTHING** — they
> read a table **no seeder ever wrote**; now seeded, the **disclosed EXCEPTION deleted**, floor **120 → 900**.
>
> **Alignment is now 100% / 100% critical (27/27)** — M218's deliberately-RED gene (**F-11**) is **fixed** and
> **retained as a permanent fence**. Both absence-gates landed too.

> **⚠ M219's battery carries TWO DISCLOSED CAVEATS — neither hidden, neither rounded away.**
> **(1) Not one uncontested consecutive run.** Two further runs went RED on **battery INTEGRITY** — an
> **ORCHESTRATION error, not a demo defect**: two batteries ran **concurrently against the single demo host** and
> one **purged the stack mid-measurement**, so a gate went **UNEXECUTED** (*an unexecuted gate is a **FINDING**, not
> a pass*). **No demo defect was observed in anything that was measured.**
> **(2) The code that GRADED is not the code that SHIPPED** — `c6648d1` (`aiReadinessStep1Score` **double-rounded**)
> is a **SEED-PATH** change landed **after** the graded tag, and **per D13 that restarts the battery count**. Small,
> strictly corrective — **and still not allowed to be rounded away.** → **`REPROVE-M221-battery-at-final-code`**.

> **⚠ THE RELEASE'S HEADLINE GATE WAS MET AT M218 — click→ACCESS 2413 ms / 1767 ms vs < 5000 ms** (worst-case p95,
> **5 GENUINE cold reset-to-seed cycles**, both vantages, 50/50 ACCESS, on `billion` over the tailnet). **From
> 39.45 s / 38.30 s — a ~16× improvement.** Two root causes, **both in demo tooling, zero platform edits**.

> **⚠ Carried into M221 (from M217):** the **pre-bind reap has still never run live** — it was dead code during
> M217's green proof run on `billion`. Fixed and unit-proven; **not field-proven**. M221 re-proves everything on
> the box.

**Phase:** **M219 closed + merged — next: `/developer-kit:build-milestone` on M220**, then **M221** (last).

**Next up:** **M220 "cue sheet"** (the `/demo-up` doc fix + remote **opt-out** + `safety.md` Part 3), then **M221
"prove it on billion"** — which re-proves **every** gate on the VM, over the tailnet, **with no flags**.
**M220 now carries TWO 🔴 demo-BREAKING escalations from M219** (below).

**Recently closed:** **M219** readiness-renders — 2026-07-14 (`closed-complete`; 13 bugs; the user's bar met on
both counts) · **M218** seat-change — 2026-07-14 (`closed-on-gate`; 21 bugs; the <5 s gate MET) · **M217**
clean-stage — 2026-07-13 (gate MET on `billion`; 32 bugs).

## ⚠️ M220 inherits two DEMO-BREAKING defects (M219 close, Fate-3 — its `overview.md` is EDITED)

- 🔴 **The academy POISONS the demo session.** `:13077` runs its own **keyless Clerk**; visiting it returns
  `Set-Cookie: __session=; Expires=1970` (**deletes the demo session**) and `__client_uat=0;
  Domain=taildc510.ts.net` — **domain-wide, not port-scoped** (**cookies scope by HOST, not by PORT**). So **a
  presenter who clicks the AI Academy link is LOGGED OUT of the demo** into `ERR_TOO_MANY_REDIRECTS`, and **every
  employee coverage sweep aborts**. Proven by controlled A/B from a peer. → M220 item **(i)**, severity **RAISED**.
- 🔴 **studio-desk `:19000` → 302 → `:13000/login`** — clicking *"Anthropos Studio"* **ejects the presenter from
  the demo**. An **in-demo dead end**, which is why the sweep's prod-eject detector never flagged it. → M220 item
  **(j)**, added at close.

*Both ship with **honest fences that deliberately report RED** until M220 lands — a half-fix that reports green is
worse than no fix.*

## D17 — the release's signature hazard (name it, then stop walking into it)

**D17: *a status artifact that outlives the thing it describes, and is then read as evidence.*** Named at the M218
close. It has now bitten **~14 times across M217→M219** — **five new instances inside M219 alone, several found by
the tooling in ITSELF**: `run-coverage.sh` printed *"coverage report written to …"* **unconditionally** and re-read
whatever JSON was on disk, so a spec that threw first presented the **PREVIOUS run's numbers** — *it nearly graded
an M219 rebuild on hours-old data from the **old, broken** stack*. **17 existing tests asserted the junk-fallback AS
THE CONTRACT** — **they were not missing the bug, they were guarding it.** The academy launcher's **fixtures had
encoded the broken behaviour**, so the suite was **GREEN for four releases while the academy 502'd**.

**The generalized lessons:** ***"the role classifies" ≠ "the pool is big enough"*** · ***"it resolves" ≠ "it has
skills"*** · ***"it serves" ≠ "it renders"*** · ***"the pool resolved" ≠ "the content is sane"*** · ***an errored SQL
query is not "zero rows"*** · ***an unexecuted gate is a FINDING, not a pass.***

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
| **M220** | Cue sheet — `/demo-up` defaults: the doc fix + remote **opt-out** + `safety.md` **Part 3** | `section` | medium | M217 |
| **M221** | Prove it on billion — every gate, on the VM, over the tailnet, **no flags** | `iterative` | large | M217–M220 |

## Headline numbers (M219 close, 2026-07-14)
- **p95 click→ACCESS (the release's headline gate, set at M218):** **2413 ms** (employee) / **1767 ms** (manager)
  vs the **< 5000 ms** gate, over **5 genuine cold reset-to-seed cycles**. **Baseline 39.45 s / 38.30 s.**
  *A latency number without its environment is not a measurement:* `billion`, Linux VM, 7.3 GiB RAM, tailnet origin.
- **Python tests:** **903** (0 fail, 12 skip) — demo-stack 506 · stack-injection 186 · stack-verify 114 ·
  stack-core 97. **+16 vs M218.** Counted from **JUnit XML**, never grepped stdout.
- **Go test funcs:** **1821** (+37 vs M218's 1784, **by the same method** — pinned in `metrics.json` so future
  closes compare like with like). **0 failures across all 6 modules.**
- **Alignment (Clerkenstein Go surface):** **100% overall / 100% critical (27/27)** — gate ≥95/=100 ⇒ **MET**.
  M218's deliberately-RED gene (**F-11**, the fabricated org eid) is **FIXED**, and the gene is **retained as a
  permanent fence**. **An unmeasurable surface now FAILS LOUD** (exit `3` ≠ exit `2`) — it used to score nothing
  and pass.
- **Demo-patches:** **7** (M219 added `next-web-aireadiness-flag-gate`). **Platform-repo edits:** **0**.
- **Fences:** graded by **MUTATION**, not by re-running them green — **7 mutations, 7 REDs**. *A fence which passes
  against both the pre- and post-fix code is theatre.*
- **Flake:** **0**.
- **Known issue (flagged, not a regression):** the **`dev-stack` suite fails on this box** — environmental (needs a
  live Postgres on `:15432`; this box's `.agentspace/secrets` is also incomplete). Identical at v2.2, M217, M218.

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
- **M216** (dev-path Tailscale parity) — **CONSUMED by M220(d)** per D-DESIGN-3. It is the release's declared
  scope-flex lever (it drops back to a reservation only if M220 bloats).
- **Plan hygiene → next close-release:** `metrics-history.md` has **no v2.2 row**; no release-scope deferral audit
  exists for v2.1 or v2.2.
- **Older, still unscheduled:** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01 (`replayCmd` hermetic
  test), CAVEAT-1 (clean-box literal full `/dev-up`), M314b (prod frozen-read hydration — a prod-team follow-up).
  Reserved **Playthroughs futures** M205–M207 stay in vision. All tracked in
  [`roadmap-vision.md`](roadmap-vision.md).

_Last updated: 2026-07-14 (M219 close — retro + metrics + Completeness Ledger written; roadmap flipped to `done`;
5 Fate-3 routes attached, both receiving `overview.md` files EDITED; zero escape-hatch deferrals. **Next: M220**,
which now carries two 🔴 demo-BREAKING escalations.)_
