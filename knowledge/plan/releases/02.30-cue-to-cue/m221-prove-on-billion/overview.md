---
milestone: M221
slug: prove-on-billion
version: v2.3 "cue to cue"
milestone_shape: iterative
status: planned
created: 2026-07-13
last_updated: 2026-07-13
complexity: large
depends_on: M217, M218, M219, M220
exit_gate: "On billion.taildc510.ts.net, a DEFAULT /demo-up N (NO FLAGS) yields, reproducibly on a cold reset-to-seed: (1) p95 click→ACCESS < 5 s for BOTH maya-thriving and dan-manager, measured over the TAILNET origin; (2) the full replayed catalog — taxonomy + directus + sim-embeddings, NO SKIPPED surface; (3) all 3 story orgs seeded incl. AI-readiness; (4) Dana sees a FILLED AI-readiness page; (5) Ben's from-scratch STARTED workflow is visible on his dashboard; (6) Aria's COMPLETED state renders; (7) remote access came up BY DEFAULT, no flag passed; (8) ZERO platform-repo edits."
iteration_protocol_ref: corpus/ops/verification.md (+ corpus/ops/demo/coverage-protocol.md + corpus/ops/demo/playthroughs.md — the remote-origin cold-reset-to-seed gates, as in M215)
delivers: every requirement of v2.3 proven live on the remote VM over the tailnet + the committed remote-origin Playwright gate that v2.2 owed (DEF-M215-03(b))
---

# M221 — Prove it on billion

## Goal
Every requirement of this release, verified **on the remote VM, over the tailnet, with no flags passed**.

## Exit gate (measurable)

On **`billion.taildc510.ts.net`**, a **DEFAULT** `/demo-up N` — **no flags** — yields, **reproducibly on a cold
reset-to-seed**:

1. **p95 click→ACCESS < 5 s** for **both** `maya-thriving` (employee) and `dan-manager` (manager), measured **over
   the tailnet origin** — the extra TLS/`tailscale serve` proxy hop is **inside** the budget, not excluded from it.
   (ACCESS as defined in M218: authenticated shell rendered + interactive + hero identity present. In-page data
   completion is **reported, not gated** — D-DESIGN-1.)
2. **The full replayed catalog** — taxonomy **+** directus content **+** sim-embeddings, with **NO SKIPPED surface**
   (the last real run skipped all three).
3. **All 3 story orgs** seeded, including the AI-readiness org.
4. **Dana** (manager) sees a **FILLED** AI-readiness page.
5. **Ben's** from-scratch **STARTED** AI-readiness workflow is **visible on his dashboard**.
6. **Aria's COMPLETED** state renders.
7. **Remote access came up BY DEFAULT** — no flag was passed (D-DESIGN-3).
8. **ZERO platform-repo edits.**

## Why iterative (not section)
The direct analogue is **M215 "prove-on-odyssey" (7.1 h, direct-drive)**: the reconfiguration is fully specified by
the upstream milestones, but **the last breakages only surface on a live cross-machine run**. A fixed `In:` list
would be speculative. Expect the same **direct-drive** shape (one canonical `iter-01/findings.md` rather than a long
tik/tok chain) — live shared infra does not reward speculative iteration.

## Iteration protocol
`corpus/ops/verification.md` + the coverage/playthroughs gates run **from a remote origin** — bring up → drive from
a second tailnet machine → capture every eject/block/warning/timing → fix in the M217/M218/M219/M220 surface →
re-run. Tik/tok until the gate holds on a cold reset-to-seed.

> **No new platform edits invented during iteration.** A surfaced platform-source hardcode routes to a **NEW
> sha-pinned demo-patch** (D-DESIGN-2) or **escalates**. It never gets edited.

## Inherited from M217 (Fate-3, added at the M217 close)

- **The pre-bind reap has never run live.** M217's close review found that `up-injected.sh` called `reap_port`
  **without sourcing `reap.sh`** — so the milestone's headline deliverable was dead code (exit 127, swallowed by
  `|| true`) **during the green proof run on `billion`**. It is fixed and unit-proven; it is **not field-proven**.
  **M221 must exercise it**: leave a cockpit listening on the offset port, re-run `/demo-up`, and assert the
  pre-bind reap actually reclaims it.
- **The compose-range preflight** (`assert_ports_free`) was likewise wired only at the close — never run in the
  field. Same treatment.
- **The freshness preflight** (built at the close) has never aborted a real bring-up. Verify it fires by
  temporarily breaking an anchor.

## Inherited from M218 (Fate-3, added at the M218 close, 2026-07-14)

Two items that can only be settled **on the box, over the tailnet** — which is precisely this milestone's job.

- **`PROBE-M218-backend-api-url-twin` (F-7) — the loaded gun.** `NEXT_PUBLIC_BACKEND_API_URL` bakes to
  `https://billion…:18082` and was **measured at 10,553 ms → `UND_ERR_CONNECT_TIMEOUT` from inside the
  container** — **the exact C-1 shape that cost M218 37.5 s per render**. It is dormant **only** because every
  current reader is client-side (M218 **D10**). A single future server-side read re-introduces the
  38-second login. **DoD:** either the address is made reachable from inside the container, or the
  server-side origin is made explicit the way `WUNDERGRAPH_SSR_ENDPOINT` was — and a probe **fences** it so a
  server-side reader can never silently appear. Verify **on `billion`**, where it actually blackholes.
- **`PROBE-M218-c3-rerun` (C-3) — now exercisable for the first time.** Until M218 landed, the SSR fetch died
  *upstream* of the router, so the federation was never exercised and C-3 could not be measured. It can be now:
  the Cosmo router **is** logging **cms/Directus 403s** (`getSkillPaths`, `_entities JobSimulation`) on the
  **CONTENT** path. Not on the login path (so it never threatened M218's gate), but it directly threatens **this
  milestone's gate item (2): "the full replayed catalog — no SKIPPED surface."** Root-cause the 403s (a
  serve-grant on the replayed Directus is the prime suspect) and fix or explain them.

## Inherited from M219 (Fate-3, added at the M219 close, 2026-07-14)

- **`GUARD-M221-host-isolation` — two agents can run cycles against ONE demo host, and nothing stops them.**
  M219's 5-cycle cold battery was corrupted **by an orchestration error, not a demo defect**: two batteries were
  run **concurrently against the single demo host**, and one of them purged the stack **mid-measurement** while
  the other was reading it. Cycle 5's `no-junk-skills` gate consequently went **UNEXECUTED** — and an unexecuted
  gate is a **FINDING, not a pass** (D17). No demo defect was observed in anything that *was* measured, and the
  five graded greens are each independently evidenced, but the audit trail now carries a permanent, disclosed
  asterisk: **they are not a single uncontested consecutive run.**

  The demo host is a **singleton shared resource** and the tooling treats it as if it were private. The same
  class already bit the coverage harness twice this milestone: `run-coverage.sh`'s out-dir was keyed on vantage
  only, so a Northwind sweep **silently overwrote** a Cervato sweep's `coverage-report.json` (fixed, M219 S1);
  and its report was re-read **without a freshness check**, so a concurrent writer's numbers could be graded as
  yours (fixed in the M219 harden pass, `TestCoverageReportFreshness`). Both are the same root cause as the
  battery corruption: **no mutual exclusion on a shared host.**

  **DoD:** a cycle **cannot start** while another holds the host. Either a **host lock** (an advertised,
  stale-tolerant lockfile on the demo host, taken for the life of a cycle and named with the holder's identity),
  or a **per-cycle stack `N`** so concurrent runs are isolated by construction. A second concurrent cycle must
  **fail loud** with the holder's identity — never queue silently, and never proceed. **This is a prerequisite
  for M221's own gate**, which is itself a multi-cycle battery on the single `billion` host: without it, this
  milestone can corrupt its own evidence exactly as M219 did.

- **`FIX-M221-reap-native-academy` — `down --purge` does not reap the native academy.** The academy is a
  **host-native** `next dev` process (not a container), so `down --purge` leaves it holding `:13077` **across
  cycles**. The next bring-up's academy then dies `EADDRINUSE` while **the OLD process keeps answering the
  port** — and the launcher log says *"the academy process DIED"* while the port still serves. Another
  *serves ≠ works* case, and it silently makes cycle N+1 measure cycle N's process. Folds naturally into the
  pre-bind-reap item above (same `reap.sh` surface) — but note that the reap must cover the **native** processes
  (cockpit **and** academy), not only the container ports.

- **`REPROVE-M221-battery-at-final-code` — the code that GRADED M219 is not the code that SHIPPED it.**
  M219's 5-cycle cold battery was graded at rext tag **`cue-to-cue-m219-r8`**. **Two commits landed after it**,
  during harden/close:
  - `b5bf65b` — the coverage-harness stale-report fix. **Harness only** — it *grades* the demo, it is not *in* it.
  - `c6648d1` — **`aiReadinessStep1Score` double-rounded.** It routed through a 0–100 intermediate
    (`round(round(held/total*100) * 30/100)`) while the platform's `computeTier1` rounds **once**; they disagree on
    **3 of the 14** reachable `heldWeight`s (2.5 → 11 vs 12 · 4.0 → 19 vs 18 · 5.5 → 26 vs 25). **This is a
    SEED-PATH change** — and **per M218 D13, a seed-path change RESTARTS the battery count** (M219 D-M219-15
    honoured that rule once already, re-running a 35-minute battery rather than argue a fix was "behaviourally
    identical").

  **The delta is small and strictly corrective** — a seeded Step-1 score off by one point for some members, and the
  fix makes the seed **agree with the platform** where it previously disagreed. It does **not** affect whether any
  section renders. **M219 did not round it away, and neither should this milestone.** It matters more *after* M219,
  which seeds an **active** (live-recomputed) cycle alongside the **closed** (frozen) one: a double-rounded frozen
  score means the **same member reports two different Step-1 scores depending on which cycle is viewed**.

  **DoD:** M219's readiness gates are re-proven **at final code** as part of this milestone's own cold
  reset-to-seed battery — it rebuilds the demo from scratch on the VM anyway, so this costs **no extra cycle**.
  Concretely, fold M219's five graded gates into this battery's assertions: **no junk skills** (enumerate every
  distinct claimed skill name), **hero role titles resolve**, `ai_readiness_cycles == 2`,
  `interview_aggregated_reports` **non-empty**, and the readiness sections **filled** (900-char floor). Note the
  dependency: this battery is **itself** exposed to the host-isolation hazard above — **`GUARD-M221-host-isolation`
  lands first, or this re-proof can corrupt its own evidence exactly as M219's did.**

## Inherited from M220 (Fate-3, added at the M220 close, 2026-07-15)

> ### ✅ `FIX-M221-devstack-test-spin` is **DISCHARGED — do not work it.**
> It was routed here during M220 S0–S2 as *"`test_dev_stack.py` **BUSY-SPINS forever** — 8 min wall, 145 % CPU,
> `rc=124`; a whole-repo run never completes, and **this will block the v2.3 close**."* **That description was
> wrong, and the item is now closed.** It was never a spin. It was two things, both fixed:
> - **(a)** both subprocess harnesses ran the **real M28 secret pre-flight**, which **compiles a Go binary** (the
>   "145 % CPU") and reads the developer's own `.agentspace/secrets`, dying `rc=1` **before reaching the code
>   under test** → fixed at M220 S7 (**D29**).
> - **(b)** `DevSetdressLocalContent.run_sd` never set **`DEV_SETDRESS_USE_STUB_BINS=1`** — the one flag
>   `build_cli` consults before honouring the stubs — so a **unit** test did a real `go build` and ran the real
>   `stackseed` against a **real Postgres that was never meant to be there** (20 tests, 486 s, **19 failures**,
>   all `connection refused`) → fixed at the M220 close (**D31**). **One environment variable.**
>
> **The whole `dev-stack` suite now runs: 116 passed · 4 skipped · 127 s.** The whole rext Python suite completes
> for the first time in the release: **1208 tests · 0 failures.** **The v2.3 close is no longer blocked by it.**
>
> It had been carried as *"a known issue — **environmental**"* in `state.md` since **v2.2**, across M217, M218
> and M219, **re-characterised every single time** and investigated by none. That is **D17 in the release's own
> headline numbers**: a status artifact that outlived the thing it described and was read as evidence. This
> paragraph stays here as the retraction — **a route that is quietly deleted teaches nobody.**

**Four genuine routes from M220, all needing a live box (which is this milestone's whole job):**

- **`FIX-M221-academy-empty-catalog` (F-M220-2).** The academy now renders and the demo session **survives** it
  (S5), but its **catalog is EMPTY**: *"0 PATHS / 0 COURSES / No adventures here… yet"* (**348 chars**) while its
  own clone HAS content — `[build-catalog]` emits **2,705 entries across 419 public chapters**. The home reads the
  **LOCAL** catalog, and `[build-local-catalog]` emits **0** (`code/ucourses/local-catalog.generated.js` = **368
  bytes**). **Not the session bug, and not caused by the Clerkenstein wiring** — a separate content-pipeline
  defect. **M219's 400-char content floor is therefore honestly RED and was deliberately NOT weakened.** A fix
  likely needs a new rext-owned demo-patch (the `ant-academy-dev-origins` precedent); the repo itself is out of
  bounds.

- **`FIX-M221-academy-loopback-bind` (F-M220-5).** **ant-academy binds `*:13077`** and answers HTTP 200 on the
  tailnet IP **even on a localhost demo** — `BIND_HOST=""` passes no `-H`, and **`next dev`'s own default is
  `0.0.0.0`**. So the academy is world-published on **every** demo, exactly like the containers. **This is the S0
  lie one layer up**, and it survived S0/S1 because `exposure_claim_guard` only ever knew about the three
  **container** port emitters — *an exposure fence blind to a whole class of listener reports a confident,
  quietly incomplete pass.* **The docs were corrected in M220 (Fate 1);** the **code** fix (`-H 127.0.0.1` when
  `BIND_HOST` is empty) is here, because it would change the very localhost path S3's HARD INVARIANT is fenced
  on. **Extend `exposure_claim_guard` to the host-native listeners at the same time** — otherwise the fence stays
  blind to the class that just bit it twice.

- **`F-M220-4` — `ant-academy.sh` is not re-runnable on a live public-host demo.** A **standalone** re-run *after*
  `tailscale serve` is configured **cannot bind** (serve holds the tailnet IPv4/IPv6 on the same port), and a
  standalone re-run **without** `STACK_PUBLIC_HOST` silently bakes **localhost** URLs into a public-host demo.
  Worked around by hand in M220 (`serve off` → launch → restore). **DoD:** re-derive the host from the registry;
  cycle the serve proxy around the bind.

- **`BURNIN-M221-dev-public-host` — the dev `--public-host` is FENCED, not LIVE-PROVEN.** S7's ladder and serve
  generator were proven live on `billion` (S3) and in v2.2, and the ladder is **reused, not forked** — but the
  **~60 lines of net-new dev wiring** (`resolve_dev_public_host`, `gen_serve_plans`, the teardown reset) are
  covered by the tripwire fence + 13 mutants **and nothing else**. It has never brought up a real remote dev
  stack. **DoD:** one `dev-stack up N --public-host auto` on the VM, reached from a peer; and one **no-flag**
  `dev-stack up` proving **zero** `tailscale` invocations on a box that HAS tailscale (the invariant the whole
  opt-in rests on, verified in the field rather than against a stub).

## Also lands
- **DEF-M215-03(b)** — the **committed, repeatable remote-origin Playwright gate** that v2.2 owed. Note that the
  latency gate **cannot be a Playthrough** (Playthroughs declare perf a **NON-GOAL**), so it is a **new
  `stack-verify` surface** — which M218 builds and this milestone runs remotely.
- **The 7.3 GiB RAM question** (**C-6**). `billion` has **7.325 GiB** vs the documented **12 GiB** floor and the
  tooling warns every run. **Measure `docker stats` + `free -h` DURING a login before blaming code** — this may be a
  pure VM resize, in which case it is an infra fix, not a code fix. Decide and record it.

## Known remote-specific hazards (from M215's findings, F1–F12)
- The **teardown must reset `tailscale serve`** (F12) — verify the shipped fix actually fires on this box; the rext
  clone there was **behind the fix** as of 2026-07-13.
- The **cockpit must now be fronted** (M220e) — it was the one plain-HTTP surface.
- **`tailscale cert` re-issue / LE rate limits** (M220's open question) — a default-on flip calls the mint on **every
  fresh demo-N**.

## KB dependencies
- `corpus/ops/demo/tailscale-serve.md` (the remote-access runbook + the F1–F12 finding set)
- `corpus/ops/verification.md` · `corpus/ops/demo/coverage-protocol.md` · `corpus/ops/demo/playthroughs.md`
- `corpus/ops/demo/latency-budget.md` ← **authored by M218** (the gate definition this milestone enforces remotely)
- `corpus/ops/safety.md` Part 3 ← **authored by M220** (the exposure contract this milestone runs under)
