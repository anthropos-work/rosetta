---
active_release: "v2.5 «the playbill» — the content-vantage release (designed 2026-07-19): fill the empty ant-academy (Thread A) + a 2nd «Content stories» cockpit tab of played sessions per content product with as-player/as-manager login-and-land, cloned from anonymized real prod sessions, VPN-scoped, source-pinned. 8 milestones M229→M236, spike-first. Branch release/02.50-the-playbill; tag will be v2.5."
active_branch: "release/02.50-the-playbill"
active_milestone: "M231 content-stories-feasibility-spike (planned — section, HARD go/no-go: discover per-product player+manager result routes + prove-by-render which land from seedable rows vs runtime-blank; confirm the prod-session sourcing+anonymization mechanism; catalog public sims by modality; rule AI-labs + the academy section in/out — the barrier before the whole Thread-B build chain)."
last_closed: "M230 — 2026-07-19"
phase: "v2.5 in development. M229 (doc re-ground) + M230 (academy demo-fill) CLOSED 2026-07-19. M230 closed-incomplete/pragmatic: the Option C academy-fs-published-fallback demo-patch (rext tag playbill-m230-academy-fs-published) is built + runtime-proven (59 cards, 0 chips); the formal cold-/demo-up card-count gate + next-web re-anchor folded to M235/M236 (Fate-3, homed). M231 content-stories feasibility spike (HARD go/no-go) BUILT 2026-07-19 on branch m231/content-stories-feasibility-spike (5 sections; deliverable corpus/ops/demo/content-stories-routes.md): Thread B is GO — sim result = persisted read (seedable); Sim+Skill-path GO, Interview GO behind a PostHog-flag demo-patch (D3→M232), AI-labs OUT (presence-only, D4→M234), Academy IN. KB-fidelity YELLOW. 0 platform edits. NEXT: harden + /developer-kit:close-milestone M231."
last_updated: "2026-07-19"
---

# State

**v2.5 "the playbill" — IN DEVELOPMENT** (designed 2026-07-19 via `/developer-kit:design-roadmap`; branch
`release/02.50-the-playbill` cut from `main`; tag will be `v2.5`). The **content-vantage release** — two threads on the
mature demo/cockpit machinery: **A** fills the empty **ant-academy** grid (it renders 0 cards because the catalog is
DB-authoritative and a demo neither sets the GraphQL endpoint nor holds academy rows → `emptyCatalogView()`; the corpus
mis-documented this); **B** adds a 2nd **"Content stories"** cockpit tab listing **played sessions** per content product
(Simulation · Skill-path legacy · Skill-path new · AI-labs), each with **as-player / as-manager** login-and-land actions,
cloned from **anonymized real production sessions**, non-manager-played, re-tenanted, **source-pinned by prod session-id**.
**User decisions (2026-07-19):** real customer-session sourcing accepted (data-controller call) · demos **VPN/tailnet-scoped**,
release **amends `safety.md` Part 3** to the honest posture · academy fill **production-faithful** (no "Draft" chip) · AI-labs
+ academy section **scoped by the M231 spike**. Tooling + docs only, **0 platform-repo edits**.

## Active milestone — M231 "content-stories feasibility spike" (planned, section — HARD go/no-go)

**Goal.** The barrier before the whole Thread-B build chain (mirrors v2.4's M222): BEFORE building anything, discover
the exact per-product **player + manager result routes** and **prove-by-render** which land from seedable rows vs are
runtime-computed-blank; confirm the **prod-session sourcing + anonymization** mechanism (read → pick interesting → pin
by prod session-id); catalog captured public sims by **modality** (confirm ≥2 voice + 1 code + 1 document assessment
sources exist); and rule **AI-labs + the academy content-product section** in or out. **Delivers →**
`corpus/ops/demo/content-stories-routes.md` (the manager-view eligibility matrix + result-route map + modality catalog
+ AI-labs verdict + the sourcing/anonymization contract). **Needs prod DB read access** (`/db-query`) for the
route/session scouting. M230 landed the academy fill (Thread A). Next: `/developer-kit:build-milestone` → M231.

## Active release — v2.5 "the playbill" (8 milestones, spike-first)

**Shape:** `M229 → M230 → M231 (HARD go/no-go) → M232 → M233 → M234 → M235 → M236`. M229 ∥ M231 research overlaps;
M230 (academy fill) must land before M235's academy section. Full milestone designs + the safety/decision record:
`roadmap.md` § Active — v2.5. **Hard constraint:** zero platform-repo edits (a runtime-computed result page that won't
render from a seeded row routes to a sha-pinned `demopatch` or escalates).

## Recently closed (milestones, newest first — max 5)

- **M230 academy-demo-fill** — 2026-07-19 (iterative, closed-incomplete/pragmatic). The Option C
  `academy-fs-published-fallback` demo-patch (rext tag `playbill-m230-academy-fs-published`) built + runtime-proven
  (59 real cards, 0 Draft chips, exact DB-authoritative code path, byte-clean revert; 14 unit tests, flake 3/3). Gate
  MET-BY-PROXY; the formal cold-`/demo-up` card-count sweep + next-web re-anchor folded to M235/M236 (Fate-3, homed). 0 platform edits.
- **M229 academy-content-model-re-ground** — 2026-07-19 (section, closed-complete). Corrected 4 docs (`ant-academy.md`
  + `frontend-tier.md` + `run_guide.md` + `CLAUDE.md`) from the false "no backend / static JSON / only Clerk" model to
  the DB-authoritative catalog (grid → academy subgraph over GraphQL → `emptyCatalogView()` on failure) + fixed the F4
  mis-attribution. The KB-fidelity prerequisite for the v2.5 academy thread. Code-verified; 0 platform edits.
- **M228 second-night** — 2026-07-18 (iterative, closed-on-gate). The corrected demo re-proven live on `billion`:
  7/7 conditions, render 5/5 per-sim (8,8,9,9,8, each ≥ floor 6, junk=0), 2 candidate heroes usable, recruiter p95
  click→ACCESS **1.27 s**, hiring-only, 4 orgs coexist. iter-03 fixed F1/F2/F3 (FeedbackSeeder + SuccessionSeeder
  guard gap the deterministic M227 test missed — caught by the LIVE re-prove) + hardened the render probe for the
  intercepting-route drawer (`RENDER_ONLY_SIM`). rext seeders 96.8% cov, flake 3/3; **0 platform edits**.
- **M227 the-notes** — 2026-07-17 (section, complete). 4 believability seed/content fixes deterministically proven +
  write-path-fenced: hiring-only content, external candidate emails, 1-sim/candidate (~8/position, gate retuned
  `≥40→≥6` everywhere), gender-matched avatars. Fix #1/#2/#4 mechanisms blended into corpus at close. Go funcs
  1888→**1902**; flake 5/5; 0 platform edits. Live re-prove → M228 (Fate-2). Deferral audit YELLOW.
- **M226 opening-night** — 2026-07-17 (closed-on-gate). 7-condition hiring gate proven live on `billion`; recruiter
  p95 1.74 s < 5 s; 5 findings fixed live; Go funcs 1887→1888; 0 platform edits.

## Recently shipped (releases, newest first — max 3)

- **v2.4 "casting call"** — 2026-07-18 (tag `v2.4`). The **recruiter-vantage / hiring-org** release: a 4th hiring org
  on the cockpit (45 candidates on 5 shared positions, compared side by side), proven live on `billion` (M228 7/7,
  recruiter p95 1.27 s), reads believably (hiring-only, external emails, 1-sim/candidate, matched avatars). 7
  milestones M222→M228; tooling + docs only, **0 platform edits**.
- **v2.3 "cue to cue"** — 2026-07-15 (tag `v2.3`). The **presenter-speed** release: click→ACCESS < 5 s proven live 8/8
  on `billion`, no flags — login p95 2.11 s / 1.31 s vs a ~39/38 s baseline (~18×). Demo comes up green, full,
  remote-default-on. 5 milestones; tooling + docs only, 0 platform edits. The `billion` demo LEFT LIVE.
- **v2.2 "panorama"** — 2026-07-12 (tag `v2.2`). External-shareability / Tailscale-serve: stacks reachable from another
  tailnet machine over one trusted HTTPS origin. First live remote Linux-VM deploy.

## Headline numbers (v2.4 M227 close, 2026-07-17)
- **Go test funcs:** **1902** (+14 vs M226's 1888 — the M227 fences across `hiring_scope_test`, `candidate_email_test`,
  `gender_test` ×2, `hiring_funnel_test`, and the 4 harden fences in `users_m227_test.go`; all modules `go vet` clean).
- **M227-touched deterministic suites (re-run GREEN at close):** stack-seeding `go test ./...` OK (13 pkgs) · tsc
  `stack-verify/e2e` exit 0 · **flake 5/5** (seeders, shuffle, `-count=1`). The 4 fixes are proven DETERMINISTICALLY;
  the LOCAL live render was env-blocked → M228.
- **p95 click→ACCESS (the standing gate):** recruiter 1.74 s (M226, 3rd vantage) · employee 2.11 s / manager 1.31 s
  (v2.3), all vs the < 5000 ms gate. (M228 re-proves it on the corrected demo.)
- **Inherited (non-milestone) carries:** demo-stack **8 pre-existing fail** (test-debt backlog) + the M204 assign-WRITE
  declared TODO → both routed to the v2.4 **release close**.
- **Alignment (Clerkenstein):** **100% / 100% critical** (M227 touched no alignment surface).
- **Flake:** **0** (milestone-owned). **Platform-repo edits:** **0.** **Supply chain:** GREEN — 0 net-new direct deps.

## D17 — the carried-forward signature hazard (v2.4 discipline)

**D17: *a status artifact that outlives the thing it describes, and is then read as evidence.*** **The keeper:** ***a
named hazard is not a fence; only an executable probe binds.*** M227 lived it again: the 4 fixes are proven by
deterministic executable fences (incl. the load-bearing real-`UsersSeeder` population write-path fence that catches a
silent revert), never by "the seed wrote the rows." Full arc:
[`releases/archive/02.30-cue-to-cue/release-retro.md`](releases/archive/02.30-cue-to-cue/release-retro.md).

## Branch model / shipped tags
**v2.4 M222→M227 CLOSED; M228 REMAINS:** `release/02.40-casting-call` (cut from `main` 2026-07-15); milestone branches
`m222/read-the-room … m227/the-notes` all merged `--no-ff` into the release branch. **M228 second-night is the last
milestone; then the `release → main` merge + the `v2.4` tag are `/developer-kit:close-release` Phase 11's job** (the
USER runs it). rext code-of-record `casting-call-m227-sections` (`63c3e8d`); the `billion` demo LEFT UP. **Shipped
tags:** **v2.3** `v2.3` · **v2.2** `v2.2` · **v2.1** `v2.1` · **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10**
`v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.3b** `v1.3.1` · **v1.3** `v1.3`
· **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`. (Full detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog (unscheduled, cross-release)
- **Test-debt + declared TODO (carried, non-gate; routed to the v2.4 RELEASE close):** (a) **8 pre-existing demo-stack
  failures** — 6 × `test_cockpit.py` (4 removed-academy-CTA + 2 v2.3.1 overlay-JS) + `test_purge` + `test_reap`;
  HEAD-identical, in files M222–M227 never touched, predating v2.4 → a future demo-stack test-debt harden pass; (b) the
  **M204 `assign-and-track.UC1` assign-WRITE** declared in-manifest `unimplemented` build-reference gap → its
  declared-TODO fate at release close. Re-confirmed fresh at M225 + M226 + M227.
- **M226 Finding-3 (Fate 3, non-gate-blocking):** the **pre-bind serve reap** (clear stale `tailscale serve` fronts on
  offset ports before bind) — a bring-up-path change on a **live-only** surface needing a live re-prove;
  **self-resolves in the default flow** → a follow-up build-iter / **M228** (the next prove-on-VM). DEF-M226-01.
- **v2.3 tail carries → v2.4 side track (non-gate; signed off at v2.3 close-release):** **F4** (academy grid renders 0
  cards — fix is in the `ant-academy` **platform repo**, out of zero-edit scope) · **BURNIN-M221-dev-public-host** ·
  **F-M220-4** · **PROBE-M218-c3-rerun** (the last three need live infra). Parked; not folded into a milestone chain.
- **Plan hygiene → next close-release:** `metrics-history.md` still lacks **v2.0 + v2.2** rows.
- **Older, still unscheduled:** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01 (`replayCmd` hermetic
  test), CAVEAT-1 (clean-box literal full `/dev-up`), M314b (prod frozen-read hydration). Reserved **Playthroughs
  futures** M206–M207 stay in vision; **M205**'s tier-gate/ATS half is a residual vision reservation. All tracked in
  [`roadmap-vision.md`](roadmap-vision.md).

_Last updated: 2026-07-17 (**M227 "the notes" CLOSED** via /developer-kit:close-milestone — section; merged `--no-ff`
into `release/02.40-casting-call`. The 4 believability seed/content fixes + gate retune `≥40→≥6`, all proven
DETERMINISTICALLY + write-path-fenced; fix #1/#2/#4 mechanisms blended into corpus at close [#M227-D1/#M227-D2/#M227-D4].
Go funcs 1888→**1902**, flake **5/5**, tsc exit 0, **0 platform-repo edits**. Deferral audit **YELLOW** [1 new single
Fate-2: the LOCAL live-render re-prove → M228; 2 inherited carries → release close]. Section 5's live render was
environment-blocked → **M228 "second night"** (the corrected-demo billion re-prove, already planned). rext
code-of-record `casting-call-m227-sections` [`63c3e8d`]; the `billion` demo LEFT UP. **NEXT: M228 second-night → then
the USER's /developer-kit:close-release** — the `release → main` merge + `v2.4` tag are the USER's to run.)_
