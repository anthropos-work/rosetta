---
active_release: "v2.4 casting call — the recruiter-vantage / hiring-org release (designed 2026-07-15; ALL 5 MILESTONES CLOSED 2026-07-17 — awaiting /developer-kit:close-release)"
active_branch: "release/02.40-casting-call"
active_milestone: "(between milestones — v2.4 complete; awaiting /developer-kit:close-release [release→main merge + v2.4 tag])"
last_closed: "M226 — 2026-07-17"
phase: "v2.4 casting call — ALL 5 MILESTONES CLOSED (M222→M226). M226 opening-night closed-on-gate: the 7-condition hiring gate proven live on billion over the tailnet (default /demo-up, no flags, 2 cold cycles + orchestrator re-verify; recruiter p95 < 5 s = 3rd vantage). NEXT: /developer-kit:close-release."
last_updated: "2026-07-17"
---

# State

**v2.4 "casting call" — ALL 5 MILESTONES CLOSED 2026-07-17; AWAITING `/developer-kit:close-release`** (designed
2026-07-15 via `/developer-kit:design-roadmap`; branch `release/02.40-casting-call` cut from `main`; tag will be
`v2.4`). The **recruiter-vantage / hiring-org release**: a **NET-NEW** 4th, **HIRING** demo org on the presenter
cockpit where **45 candidates audition on the same 5 positions and a manager compares them side by side**, distinct
from the three workforce orgs. **M226 "opening night" CLOSED 2026-07-17 (closed-on-gate — the FINAL milestone)**: on
`billion.taildc510.ts.net`, a **default `/demo-up` (no flags)** hit the **7-condition hiring gate reproducibly across
2 clean cold reset-to-seed cycles** + independently orchestrator-re-verified from this Mac — 5 admin + 45 candidate,
recruiter comparison 42 × each of 5 positions (junk=0), 2 usable candidate profiles, reads as hiring, **recruiter p95
click→ACCESS 1.74 s < 5 s (the 3rd measured vantage)**, coexists with the 3 workforce orgs, **0 platform edits**. 5
cross-machine findings surfaced + fixed live (all tooling/harness/seed). **The `release → main` merge + the `v2.4` tag
remain `/developer-kit:close-release`'s job — the USER runs it.**

## Active release — v2.4 "casting call" (all milestones closed; awaiting close-release)

**Theme.** *The recruiter's vantage — 45 candidates audition on the same 5 positions; a manager compares them side by
side, distinct from the three workforce orgs on the cockpit.* **Reverses v2.3's D-DESIGN-4** (*"no hiring org, none
will be built"*): the comparison surface ships in the **dockerized `apps/web`/`apps/hiring`** the demo already builds,
not a Vercel-only app. **Consumes the recruiter/seeder half of the reserved vision M205** (Stripe-tier-gate + ATS half
stays a residual vision reservation).

**Shape — 5 milestones, largely SEQUENTIAL — ALL CLOSED:**

- **M222 read-the-room** (`section`) — ✅ **DONE 2026-07-15 (GO).** The hiring-model spike + `hiring.md` + the
  `is_hiring` gate; the mirror-table trap named (score = `local_jobsimulation_sessions`).
- **M223 casting-the-ensemble** (`section`) — ✅ **DONE 2026-07-16.** The 4th story Meridian Talent (5 admin + 45
  candidate) + `HiringConfigSeeder` + `HiringFunnelSeeder` → the `local_jobsimulation_sessions` MIRROR pair.
- **M224 the-callback** (`iterative`) — ✅ **DONE 2026-07-16 (closed-on-gate).** The TOK-02 two-app demo (genuine
  `apps/hiring` as a 2nd UI container); Results scoreboard 20/page × 43 comparable; Clerkenstein `isHiring` (`/align-run`
  100/100); 4 hiring demo-patches.
- **M225 dress-the-set** (`section`) — ✅ **DONE 2026-07-17.** Auto-set-dress bring-up GUARD + hiring coverage gate
  (3 seats MET) + one GREEN recruiter playthrough on pt-world Org D "Kestrel Hiring Group".
- **M226 opening-night** (`iterative`) — ✅ **DONE 2026-07-17 (closed-on-gate — FINAL).** The 7-condition live billion
  proof, default `/demo-up`, no flags, 2 cold cycles + orchestrator re-verify; recruiter p95 < 5 s (3rd vantage);
  0 platform edits. The `billion` demo **LEFT UP** as the live-proof artifact.

**Hard constraints (carried, unchanged):** **zero platform-repo edits** — a platform-source render gate routes to a
sha-pinned `demopatch` or escalates; all stack-operating tooling lives in `rosetta-extensions` (authored in
`.agentspace/rosetta-extensions/`, tagged, consumed per-stack at a pinned tag).

Full design (file:line-cited): [`releases/02.40-casting-call/design-notes.md`](releases/02.40-casting-call/design-notes.md).
Milestone contracts: [`releases/02.40-casting-call/`](releases/02.40-casting-call/).

## Recently closed (milestones, newest first — max 5)

- **M226 opening-night** — 2026-07-17 (closed-on-gate; FINAL v2.4 milestone). 7-condition hiring gate proven live on
  `billion`; recruiter p95 1.74 s < 5 s; 5 findings fixed live; Go funcs 1887→**1888**; flake 0; 0 platform edits.
- **M225 dress-the-set** — 2026-07-17 (section, complete). Auto-set-dress guard + coverage gate (3 seats) + 1 recruiter
  playthrough on pt-world Org D.
- **M224 the-callback** — 2026-07-16 (closed-on-gate). Two-app `apps/hiring` demo; 20/page × 43 comparable; `isHiring`.
- **M223 casting-the-ensemble** — 2026-07-16 (section). 4th hiring story + HiringConfig/HiringFunnel seeders.
- **M222 read-the-room** — 2026-07-15 (section, GO). Hiring-model spike + `hiring.md` + `is_hiring` gate.

## Recently shipped (releases, newest first — max 3)

- **v2.3 "cue to cue"** — 2026-07-15 (tag `v2.3`). The **presenter-speed** release: click→ACCESS < 5 s proven live 8/8
  on `billion`, no flags — login p95 2.11 s / 1.31 s vs a ~39/38 s baseline (~18×). Demo comes up green, full,
  remote-default-on. 5 milestones; tooling + docs only, 0 platform edits. The `billion` demo LEFT LIVE.
- **v2.2 "panorama"** — 2026-07-12 (tag `v2.2`). External-shareability / Tailscale-serve: stacks reachable from another
  tailnet machine over one trusted HTTPS origin. First live remote Linux-VM deploy.
- **v2.1 "quick change"** — 2026-07-09 (tag `v2.1`). The skiller-in-app re-ground (skiller → `app`/`public`, RPC →
  backend, 4 subgraphs); proved dev-up + demo-up cold.

## Headline numbers (v2.4 M226 close, 2026-07-17)
- **p95 click→ACCESS (the gate):** **recruiter 1.74 s** (M226, the 3rd vantage) · employee **2.11 s** / manager
  **1.31 s** (v2.3), all vs the **< 5000 ms** gate, on `billion` over the tailnet, cold reset-to-seed.
- **Go test funcs:** **1888** (+1 vs M225's 1887 — the M226 `hiring_count_harden_test.go` net-5-admin/45-candidate
  fence, RED-proven; all modules `go vet` clean, **0 Go failures**).
- **M226-touched deterministic suites (re-run GREEN at close):** stack-seeding `go test ./...` OK · `test_injection.py`
  **145 passed / 8 skipped** (`TestTailscaleServe` 34) · stack-verify/e2e `tsc --noEmit` exit 0 · **flake 5/5**. The
  live 7-condition gate was proven on `billion` (2 cold cycles + orchestrator re-verify); NOT re-brought-up at close.
- **Inherited (non-milestone) carries:** demo-stack **8 pre-existing fail** (test-debt backlog) + the M204 assign-WRITE
  declared TODO → both routed to the v2.4 **release close**.
- **Alignment (Clerkenstein):** **100% / 100% critical** (M226 touched no alignment surface).
- **Flake:** **0** (milestone-owned). **Platform-repo edits:** **0.** **Supply chain:** GREEN — 0 net-new direct deps.

## D17 — the carried-forward signature hazard (v2.4 discipline)

**D17: *a status artifact that outlives the thing it describes, and is then read as evidence.*** **The keeper:** ***a
named hazard is not a fence; only an executable probe binds.*** Directly lived again at M226: the 7-condition gate is
proven by an executable render/latency probe on `billion`, never by "the seed wrote the rows." Full arc:
[`releases/archive/02.30-cue-to-cue/release-retro.md`](releases/archive/02.30-cue-to-cue/release-retro.md).

## Branch model / shipped tags
**v2.4 ALL MILESTONES CLOSED:** `release/02.40-casting-call` (cut from `main` 2026-07-15); milestone branches
`m222/read-the-room … m226/opening-night` all merged `--no-ff` into the release branch. **The `release → main` merge +
the `v2.4` tag are `/developer-kit:close-release` Phase 11's job** (the USER runs it). rext code-of-record
`casting-call-m226-close` (`7032aea`); the `billion` demo LEFT UP. **v2.3 CLOSED** (the `v2.3` tag done at its
close-release). **Shipped tags:** **v2.3** `v2.3` · **v2.2** `v2.2` · **v2.1** `v2.1` · **v2.0** `v2.0` · **v1.10b**
`v1.10.1` · **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.3b**
`v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`. (Full detail:
[`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog (unscheduled, cross-release)
- **Test-debt + declared TODO (carried, non-gate; routed to the v2.4 RELEASE close):** (a) **8 pre-existing demo-stack
  failures** — 6 × `test_cockpit.py` (4 removed-academy-CTA + 2 v2.3.1 overlay-JS) + `test_purge` + `test_reap`;
  HEAD-identical, in files M222–M226 never touched, predating v2.4 → a future demo-stack test-debt harden pass; (b) the
  **M204 `assign-and-track.UC1` assign-WRITE** declared in-manifest `unimplemented` build-reference gap → its
  declared-TODO fate at release close. Both re-confirmed fresh at M225 + M226 close.
- **M226 Finding-3 (Fate 3, non-gate-blocking):** the **pre-bind serve reap** (clear stale `tailscale serve` fronts on
  offset ports before bind) — a bring-up-path change on a **live-only** surface needing a live re-prove (forbidden at
  close; billion left UP); **self-resolves in the default flow** → a follow-up build-iter / the next `prove-on-<VM>`
  milestone. Recorded in the M226 deferral audit (DEF-M226-01).
- **v2.3 tail carries → v2.4 side track (non-gate; signed off at v2.3 close-release):** **F4** (academy grid renders 0
  cards — fix is in the `ant-academy` **platform repo**, out of zero-edit scope) · **BURNIN-M221-dev-public-host** ·
  **F-M220-4** · **PROBE-M218-c3-rerun** (the last three need live infra). Parked; not folded into a milestone chain.
- **Plan hygiene → next close-release:** `metrics-history.md` still lacks **v2.0 + v2.2** rows.
- **Older, still unscheduled:** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01 (`replayCmd` hermetic
  test), CAVEAT-1 (clean-box literal full `/dev-up`), M314b (prod frozen-read hydration). Reserved **Playthroughs
  futures** M206–M207 stay in vision; **M205**'s tier-gate/ATS half is a residual vision reservation. All tracked in
  [`roadmap-vision.md`](roadmap-vision.md).

_Last updated: 2026-07-17 (**M226 "opening night" CLOSED** via /developer-kit:close-milestone — iterative,
closed-on-gate; merged `--no-ff` into `release/02.40-casting-call`. **The FINAL v2.4 milestone**: the 7-condition
hiring gate proven live on `billion` over the tailnet, default `/demo-up`, no flags, 2 cold cycles + orchestrator
re-verify; recruiter p95 1.74 s < 5 s = the 3rd measured vantage; 5 cross-machine findings fixed live; Go funcs
1887→**1888**, flake **5/5**, **0 platform-repo edits**. Deferral audit **YELLOW** [Finding-3 Fate-3 forward-routed;
2 inherited carries → v2.4 release close]. rext code-of-record `casting-call-m226-close` [`7032aea`]; the `billion`
demo LEFT UP as the live-proof artifact. **v2.4 all 5 milestones closed → NEXT: /developer-kit:close-release** — the
`release → main` merge + `v2.4` tag are the USER's to run.)_
