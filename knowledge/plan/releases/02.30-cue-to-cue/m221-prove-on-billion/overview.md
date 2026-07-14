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
