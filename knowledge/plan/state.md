---
active_release: "v2.4 casting call — the recruiter-vantage / hiring-org release (designed 2026-07-15; ALL milestones M222→M228 CLOSED 2026-07-18 — v2.4 is CODE-COMPLETE; the release→main merge + the v2.4 tag are /developer-kit:close-release's job, which the USER runs on explicit sign-off)"
active_branch: "release/02.40-casting-call"
active_milestone: "(between milestones — v2.4 code-complete; all milestones M222→M228 closed; awaiting the USER's /developer-kit:close-release)"
last_closed: "M228 — 2026-07-18"
phase: "v2.4 CODE-COMPLETE. All 7 milestones M222→M228 closed + merged into release/02.40-casting-call. M228 second-night closed-on-gate 2026-07-18: the corrected (M227) hiring demo re-proven LIVE on billion, 7/7 conditions — render 5/5 per-sim (8,8,9,9,8, each ≥ floor 6, junk=0), 2 candidate heroes usable, recruiter p95 click→ACCESS 1.27 s, hiring-only, 4 orgs coexist, 0 platform edits. iter-03 also fixed F1/F2/F3 (FeedbackSeeder + SuccessionSeeder guard gap the deterministic M227 test missed — caught by the live re-prove). NEXT: the USER runs /developer-kit:close-release (release→main merge + v2.4 tag) on explicit sign-off. NOT tagged until then."
last_updated: "2026-07-18"
---

# State

**v2.4 "casting call" — ALL milestones M222→M228 CLOSED; v2.4 CODE-COMPLETE, awaiting `/developer-kit:close-release`**
(designed 2026-07-15 via `/developer-kit:design-roadmap`; branch `release/02.40-casting-call` cut from `main`; tag
will be `v2.4`). The **recruiter-vantage / hiring-org release**: a **NET-NEW** 4th, **HIRING** demo org on the
presenter cockpit where **45 candidates audition on the same 5 positions and a manager compares them side by side**,
distinct from the three workforce orgs. **RE-OPENED 2026-07-17** for believability fixes from live feedback (the demo
worked on `billion` but didn't fully *read* as real), then **CLOSED OUT 2026-07-18**: **M227 "the notes"** landed the
4 seed/content believability fixes + gate retune (deterministic), and **M228 "second night"** re-proved the corrected
demo **LIVE on `billion`** — 7/7 conditions, closed-on-gate. **v2.4 is now code-complete — the `release → main` merge +
the `v2.4` tag remain `/developer-kit:close-release`'s job; the USER runs it on explicit sign-off.**

## Between milestones — v2.4 code-complete, awaiting `/developer-kit:close-release`

All 7 v2.4 milestones (M222→M228) are closed and merged into `release/02.40-casting-call`. **No milestone is active.**
The believability re-open is resolved: the corrected hiring demo reads right AND is proven live on `billion` (M228,
7/7, closed-on-gate 2026-07-18). **Next: the USER runs `/developer-kit:close-release`** — it reviews the full release,
merges `release/02.40-casting-call → main`, and tags `v2.4`. NOT tagged until then.

## Active release — v2.4 "casting call" (M222→M227 closed; M228 remains, then close-release)

**Theme.** *The recruiter's vantage — 45 candidates audition on the same 5 positions; a manager compares them side by
side, distinct from the three workforce orgs on the cockpit.* **Reverses v2.3's D-DESIGN-4** (*"no hiring org, none
will be built"*): the comparison surface ships in the **dockerized `apps/web`/`apps/hiring`** the demo already builds,
not a Vercel-only app. **Consumes the recruiter/seeder half of the reserved vision M205.**

**Shape — 7 milestones (RE-OPENED); ALL M222→M228 CLOSED:**

- **M222 read-the-room** (`section`) — ✅ DONE 2026-07-15 (GO). Hiring-model spike + `hiring.md` + the `is_hiring`
  gate; the mirror-table trap named (score = `local_jobsimulation_sessions`).
- **M223 casting-the-ensemble** (`section`) — ✅ DONE 2026-07-16. Meridian Talent (5 admin + 45 candidate) +
  `HiringConfigSeeder` + `HiringFunnelSeeder` → the `local_jobsimulation_sessions` MIRROR pair.
- **M224 the-callback** (`iterative`) — ✅ DONE 2026-07-16 (closed-on-gate). The two-app demo (genuine `apps/hiring`
  as a 2nd UI container); Results scoreboard × 43 comparable; Clerkenstein `isHiring`; 4 hiring demo-patches.
- **M225 dress-the-set** (`section`) — ✅ DONE 2026-07-17. Auto-set-dress GUARD + hiring coverage gate (3 seats MET)
  + one GREEN recruiter playthrough on pt-world Org D.
- **M226 opening-night** (`iterative`) — ✅ DONE 2026-07-17 (closed-on-gate). The 7-condition live billion proof,
  default `/demo-up`, no flags; recruiter p95 < 5 s (3rd vantage). The `billion` demo LEFT UP.
- **M227 the-notes** (`section`) — ✅ DONE 2026-07-17. 4 believability seed/content fixes + gate retune `≥40→≥6`,
  deterministically proven; live re-prove → M228. **0 platform edits.**
- **M228 second-night** (`iterative`) — ✅ DONE 2026-07-18 (closed-on-gate). The corrected demo re-proven live on
  `billion`: 7/7 conditions, render 5/5 per-sim (8,8,9,9,8), 2 heroes usable, recruiter p95 1.27 s. iter-03 also
  fixed the F1/F2/F3 seeder-guard gap the deterministic M227 test missed (caught by the live re-prove). **0 platform edits.**

**Hard constraints (carried, unchanged):** **zero platform-repo edits** — a platform-source render gate routes to a
sha-pinned `demopatch` or escalates; all stack-operating tooling lives in `rosetta-extensions` (authored in
`.agentspace/rosetta-extensions/`, tagged, consumed per-stack at a pinned tag).

Full design (file:line-cited): [`releases/02.40-casting-call/design-notes.md`](releases/02.40-casting-call/design-notes.md).
Milestone contracts: [`releases/02.40-casting-call/`](releases/02.40-casting-call/).

## Recently closed (milestones, newest first — max 5)

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
- **M225 dress-the-set** — 2026-07-17 (section, complete). Auto-set-dress guard + coverage gate (3 seats) + 1
  recruiter playthrough on pt-world Org D.
- **M224 the-callback** — 2026-07-16 (closed-on-gate). Two-app `apps/hiring` demo; 20/page × 43 comparable; `isHiring`.

## Recently shipped (releases, newest first — max 3)

- **v2.3 "cue to cue"** — 2026-07-15 (tag `v2.3`). The **presenter-speed** release: click→ACCESS < 5 s proven live 8/8
  on `billion`, no flags — login p95 2.11 s / 1.31 s vs a ~39/38 s baseline (~18×). Demo comes up green, full,
  remote-default-on. 5 milestones; tooling + docs only, 0 platform edits. The `billion` demo LEFT LIVE.
- **v2.2 "panorama"** — 2026-07-12 (tag `v2.2`). External-shareability / Tailscale-serve: stacks reachable from another
  tailnet machine over one trusted HTTPS origin. First live remote Linux-VM deploy.
- **v2.1 "quick change"** — 2026-07-09 (tag `v2.1`). The skiller-in-app re-ground (skiller → `app`/`public`, RPC →
  backend, 4 subgraphs); proved dev-up + demo-up cold.

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
