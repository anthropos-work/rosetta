---
active_release: "(between releases) — v2.3 «cue to cue» SHIPPED 2026-07-15 (tag v2.3). The next release awaits /developer-kit:design-roadmap."
active_branch: "release/02.30-cue-to-cue (release→main merge + the v2.3 tag = /developer-kit:close-release Phase 11, pending)"
active_milestone: "(between releases)"
last_closed: "v2.3 «cue to cue» — 2026-07-15 (M221 was the final milestone; the release is closed)"
phase: "between releases — awaiting /developer-kit:design-roadmap"
last_updated: "2026-07-15"
---

# State

**Between releases.** **v2.3 "cue to cue" is CLOSED** (release-level review done, fixes applied, metrics +
retro + changelog written, deferrals signed off, roadmap rotated). The only step left is the mechanical
**`release/02.30-cue-to-cue → main` merge + the `v2.3` tag** — **`/developer-kit:close-release` Phase 11's job,
pending**. No milestone or release is active. **Next: `/developer-kit:design-roadmap`** decides + cuts the next
release (it owns `active_release`).

**Next up:** **`/developer-kit:design-roadmap`** — design the next release (candidate scope: the 4 v2.3 tail
carries → v2.4, and/or a new user-driven theme). See [`roadmap-vision.md`](roadmap-vision.md) § v2.4.

## Recently shipped (releases, newest first — max 3)

- **v2.3 "cue to cue"** — 2026-07-15 (tag `v2.3`). The **presenter-speed** release: **click→ACCESS < 5 s proven
  live 8/8 on `billion`** over the tailnet, no flags — **login p95 2.11 s / 1.31 s** vs a ~39/38 s baseline (~18×).
  Demo comes up green, full, remote-default-on; AI-readiness renders filled; `safety.md` Part 3; the ~24-instance
  **D17** thread told honestly. 5 milestones M217→{M218∥M219∥M220}→M221; tooling + docs only, 0 platform edits.
  The `billion` demo is **LEFT LIVE**.
- **v2.2 "panorama"** — 2026-07-12 (tag `v2.2`). External-shareability / Tailscale-serve: dev/demo stacks
  reachable from another machine on a tailnet over one trusted HTTPS origin. Opt-in default-off (flipped to
  demo-default-on at v2.3 M220). First live remote Linux-VM deploy.
- **v2.1 "quick change"** — 2026-07-09 (tag `v2.1`). The skiller-in-app re-ground: re-fit tooling + corpus +
  stacks to the merged platform (skiller → `app`/`public`, RPC → backend, 4 subgraphs); proved dev-up + demo-up
  cold.

## Headline numbers (v2.3 close, 2026-07-15)
- **p95 click→ACCESS (the release's headline gate, set M218, re-proven live at M221 with NO flags):** **2.11 s**
  (employee `maya-thriving`) / **1.31 s** (manager `dan-manager`) vs the **< 5000 ms** gate, on `billion` over the
  tailnet, cold reset-to-seed. **Baseline 39.45 s / 38.30 s** (~18×). *A latency number without its environment is
  not a measurement:* Linux VM, 7.3 GiB RAM, tailnet origin.
- **Go test funcs:** **1831** (+82 vs v2.2's 1749 by the same `grep -c '^func Test'` method; 0 failures across all
  6 modules; `go vet` clean).
- **Python tests:** **1341** (0 fail, 16 skip) via JUnit XML — demo-stack 663 · stack-injection 260 · stack-core
  182 · dev-stack 122 · stack-verify 114. v2.2 was 668 across 3 sections (the whole tree completes for the first
  time as of M220).
- **TS e2e:** **151** (69 stack-verify/e2e + 82 playthroughs/e2e via `playwright test --list`; +27 vs v2.2's 124 —
  M219's differently-measured 94 is NOT used, so no false 124→94 regression).
- **Alignment (Clerkenstein Go surface):** **100% / 100% critical (27/27)** — held.
- **Flake:** **0** (close-release triple-clean 3/3, randomized order). **Platform-repo edits:** **0** (rosetta diff
  108 files, all docs/planning). **Supply chain:** GREEN — 0 net-new direct deps; the only change is `x/crypto
  v0.51.0 → v0.52.0` (indirect, landed M218); `govulncheck` clean all 6.

## D17 — the release's signature hazard (carried into the next release's discipline)

**D17: *a status artifact that outlives the thing it describes, and is then read as evidence.*** Named at the M218
close; it bit **~24 times across M217→M221** — and it kept turning inward, **the fences catching themselves**
(a torn-down stack still reporting "green"; a test named for a check it wasn't running; the exposure fence blind to
host-native listeners; `run-coverage.sh` re-printing the previous run's numbers; the dev-stack suite's
"environmental" excuse sitting in this file's own headline numbers for four releases with **one missing env var**
underneath). **The keeper:** ***a named hazard is not a fence; only an executable probe binds.*** Full arc:
[`releases/archive/02.30-cue-to-cue/release-retro.md`](releases/archive/02.30-cue-to-cue/release-retro.md).

## Branch model / shipped tags
**v2.3 CLOSED (awaiting the Phase-11 merge+tag):** `release/02.30-cue-to-cue` cut from `main` (2026-07-13); all 5
milestone branches `m217/clean-stage … m221/prove-on-billion` merged `--no-ff` and deleted. The `release → main`
merge + the `v2.3` tag are `/developer-kit:close-release` Phase 11 (pending). **Shipped tags:** **v2.2** `v2.2` ·
**v2.1** `v2.1` · **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` ·
**v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` ·
**v1.1** `v1.1` · **v1.0** `v1.0`. (Full detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog (unscheduled, cross-release)
- **v2.3 tail carries → v2.4 (non-gate; user signed off 2026-07-15 at close-release; landing spot
  [`roadmap-vision.md`](roadmap-vision.md) § v2.4):** **F4** (academy grid renders 0 cards — fix is in the
  `ant-academy` **platform repo**, out of zero-platform-edit scope) · **BURNIN-M221-dev-public-host** (dev-path
  live burn-in) · **F-M220-4** (academy re-run on a live public-host demo) · **PROBE-M218-c3-rerun** (router-403
  re-check) — the last three all need **live infra**.
- **Plan hygiene → next close-release:** `metrics-history.md` still lacks **v2.0 + v2.2** rows (the v2.3 row landed
  at this close); no release-scope deferral audit exists for v2.1 or v2.2.
- **Older, still unscheduled:** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01 (`replayCmd` hermetic
  test), CAVEAT-1 (clean-box literal full `/dev-up`), M314b (prod frozen-read hydration — a prod-team follow-up).
  Reserved **Playthroughs futures** M205–M207 stay in vision. All tracked in
  [`roadmap-vision.md`](roadmap-vision.md).

_Last updated: 2026-07-15 (v2.3 "cue to cue" CLOSED via /developer-kit:close-release — docs re-synced [academy
loopback-bind LANDED, remote-flip contradiction fixed], 4 tail carries signed off → v2.4, release metrics + retro +
dependencies.lock + CHANGELOG written, triple-clean 3/3, roadmap rotated to Done + release dir archived, state
rotated to between-releases. The release→main merge + the `v2.3` tag remain /developer-kit:close-release Phase 11.
The `billion` demo is LEFT LIVE. **NEXT: /developer-kit:design-roadmap.**)_
