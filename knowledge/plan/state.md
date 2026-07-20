---
active_release: "v2.5 «the playbill» — the content-vantage release (designed 2026-07-19): fill the empty ant-academy (Thread A) + a 2nd «Content stories» cockpit tab of played sessions per content product with as-player/as-manager login-and-land, cloned from anonymized real prod sessions, VPN-scoped, source-pinned. 8 milestones M229→M236, spike-first. Branch release/02.50-the-playbill; tag will be v2.5."
active_branch: "release/02.50-the-playbill"
active_milestone: "M236 prove-on-billion (PLANNED — iterative, medium; depends on M235): re-prove the whole content-vantage feature LIVE on the billion Tailscale VM — both cockpit tabs usable end-to-end from a 2nd tailnet machine on a cold reset-to-seed (Content-stories sessions render real content for player + manager; the academy grid renders real cards), p95 click→ACCESS < 5 s, 0 platform edits, VPN-scoped. AUTHORS the new content-stories seat-login coverage/Playthrough plumbing (M235's USER-BLOCKER-M235-02 — the exact-path/hero-crawl harness can't reach the dynamic-URL cockpit-seat-reached result pages) + works the M235 per-section live-calibration checklists + the M230 carry-forward live items. THE FINAL v2.5 milestone."
last_closed: "M235 — 2026-07-20"
phase: "v2.5 in development. M229–M235 CLOSED. M235 prove-it-lands CLOSED (iterative, closed-incomplete/pragmatic — the LIVE (session×action)-lands gate routes to M236 by design). Built + unit-proven, 0 platform edits: the provably-clean 13-session simulation matrix + all 3 non-simulation sections (skill-path real progress + mirror; ai-labs presence-only; academy /library/<slug> CTA) via content_nonsim.go → manifest 4 products/18 sessions, honesty gates GREEN. 2 user-blockers resolved (M235-01 scrub-leaked-zero-names → hardened+re-captured provably-clean; M235-02 coverage-descriptor-missing → 'build non-sim, then close'). rext playbill-m235-hardened @ 60eff14; Go 1939→1974, flake 5/5. carry-forward 3 clusters all Fate-3→M236 (applied 54eaefe). NEXT: M236 prove-on-billion (FINAL v2.5). Standing carry: 14 pre-existing demo-stack fails (REPEAT) → v2.5 release-close."
last_updated: "2026-07-20"
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

## Active milestone — M236 "prove-on-billion" (PLANNED — iterative, medium; depends on M235) · **FINAL v2.5 milestone**

**Goal.** Re-prove the whole content-vantage feature **LIVE on the `billion` Tailscale VM** (the house pattern that
closed M215/M221/M226/M228). **Exit gate:** both cockpit tabs work live on `billion` — Content-stories sessions render
real content for player + manager, the academy grid renders real cards — cold reset-to-seed, **p95 click→ACCESS < 5 s,
0 platform edits, tailnet-only**. **Iterative** (`verification.md` + `tailscale-serve.md`). It **AUTHORS the new
content-stories seat-login coverage/Playthrough plumbing** (M235's USER-BLOCKER-M235-02 — the exact-path/hero-crawl
harness can't reach the dynamic-URL, cockpit-seat-reached result pages; author + calibrate live), works the **M235
per-section live-calibration checklists**, and lands the **M230 carry-forward live items** — all Fate-3 from M235's
`carry-forward.md`. Then the `release → main` merge + `v2.5` tag are `/developer-kit:close-release`'s job. **0 platform
edits.**

## Active release — v2.5 "the playbill" (8 milestones, spike-first)

**Shape:** `M229 → M230 → M231 (HARD go/no-go) → M232 → M233 → M234 → M235 → M236`. M229 ∥ M231 research overlaps;
M230 (academy fill) must land before M235's academy section. Full milestone designs + the safety/decision record:
`roadmap.md` § Active — v2.5. **Hard constraint:** zero platform-repo edits (a runtime-computed result page that won't
render from a seeded row routes to a sha-pinned `demopatch` or escalates).

## Recently closed (milestones, newest first — max 5)

- **M235 prove-it-lands** — 2026-07-20 (iterative, closed-incomplete/pragmatic; LIVE gate → M236 by design). Built +
  unit-proven, **0 platform edits**: the provably-clean **13-session simulation matrix** + all **3 non-simulation
  sections** (skill-path real progress + mirror; ai-labs presence-only; academy `/library/<slug>` CTA) via a code-owned
  registry `content_nonsim.go` → manifest **4 products / 18 sessions**, honesty gates GREEN. 2 user-blockers resolved
  (**M235-01** scrub-leaked-zero-names → hardened+re-captured provably-clean; **M235-02** coverage-descriptor-missing
  → "build non-sim, then close"). rext `playbill-m235-hardened @ 60eff14`; Go **1939→1974**, flake 5/5. Deferral audit
  YELLOW/0-blockers. **carry-forward.md** 3 clusters all Fate-3 → M236 (applied, `54eaefe`). Full narrative: roadmap `### M235`.
- **M234 content-stories-cockpit-tab** — 2026-07-19 (section, closed-complete). The **render half** of Content
  stories: `cockpit.py`'s 2nd "Content stories" tab reads the M233 `content-manifest.json` (per-product sections,
  per-session rows, two login-and-land CTAs, AI-labs presence-only, academy direct-origin link); `roster.go` appends
  `content-player-<idx>` seats single-sourced via the new `storyPopulationNames`; `up-injected.sh` wires
  `--content-export`/`--content-manifest` (non-fatal). Renderer handles ALL dispositions, **unit-proven** (Python
  249/6-pre-existing/0-new; Go +8 1931→1939; flake 5/5 both stacks). rext tags `playbill-m234-content-tab` @ 7f55eb4
  + `-hardened` @ fd457bf. Close near-clean (2 record fixes), deferral audit YELLOW/0-blockers, 0 platform edits;
  non-sim fixtures + prove-it-lands = M235 (Fate-2, homed).
- **M233 content-stories-manifest** — 2026-07-19 (section, closed-complete). The **manifest half** of Content
  stories: `BuildContentProducts` projects a `content_products[]` menu (per product, played sessions with
  player+manager seats + result paths + `has_manager_view` + app_base + icon) SINGLE-SOURCED from the M232 fixture;
  honesty-gated (`CanonicalFileMatchesProjection` + teeth) so `content-manifest.json` can't drift; fail-closed
  (drop-with-reason, fails loud). Emitted by `stackseed --content-export`; open question resolved (separate JSON,
  `#D-M233-1`). Flat-index-survives-drops seat invariant pinned. rext tags `playbill-m233-content-manifest` @ 9f0ab1c
  + `-hardened` @ c30fee3. Close near-clean (1 fix), deferral audit YELLOW, flake 5/5, 0 platform edits.
- **M232 session-clone-sourcing-seeder** — 2026-07-19 (section, closed-complete). The ContentStorySeeder
  **COPIES real prod sessions** (feedback/transcript/submission/interview report/node-ids) + best-effort PII
  scrub (names/org→placeholders, emails/phones/urls redacted) + re-tenant + source-pin (rext tag
  `playbill-m232-sections-copyreal`); interview flags via 2 demopatches; `safety.md` §3.8 = data-controller-
  accepted residual-risk, VPN-scoped. A synthesize-first build was REWORKED to copy-real per user decision.
  Guardrails flake 5/5. 0 platform edits.
- **M231 content-stories-feasibility-spike** — 2026-07-19 (section, closed-complete, **GO**). The Thread-B
  go/no-go barrier: delivered `content-stories-routes.md` (result-route map + prove-by-render + sourcing/anon
  contract + modality catalog). Central risk resolved — sim result page reads a persisted DB row (seedable).
  Sim+Skill-path GO, Interview GO w/ flag demo-patch (D3→M232), AI-labs OUT/presence-only (D4→M234), Academy IN
  (D5→M234). Fixed 3 stale corpus claims inline. 0 platform edits.
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

## Headline numbers (v2.5 M235 close, 2026-07-20)
- **Go test funcs (whole rext repo, `git grep '^func Test'`):** **1974** (M235 +35 vs M234's fd457bf **1939**;
  the REPRODUCIBLE method — v1.11 incident-#6 counting discipline). `go build`+`go vet` clean.
- **M235-touched suites (re-run GREEN at close):** Go `scrub`+`contentsession`+`seeders`+`cmd/content-capture`+
  `cmd/stackseed` honesty gates OK · **flake 5/5** (Go shuffle `-count=1`). Harden Pass 1+2 `--final` stabilized
  (scrub 94.6→100%, contentsession 93.7→94.7%, seeders 95.9→96.1%, content-capture 0→28.7%; 0 bugs). Close fixes
  docs/records only — **no rext code change** (code-of-record `60eff14` unchanged).
- **p95 click→ACCESS (the standing gate, carried — M235 touched no perf surface):** recruiter 1.27 s (M228) ·
  employee 2.11 s / manager 1.31 s (v2.3), all vs the < 5000 ms gate.
- **Inherited (non-milestone) carries:** demo-stack **14 pre-existing fail** (test-debt backlog, REPEAT v2.4→v2.5;
  M235's slice = 6 `test_cockpit.py`, added 0 NEW) + the M204 assign-WRITE declared TODO → both routed to the v2.5
  **release close** re-anchor.
- **Alignment (Clerkenstein):** **100% / 100% critical** (M235 touched no alignment surface).
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

_Last updated: 2026-07-20 (**M235 "prove-it-lands" CLOSED** via /developer-kit:close-milestone — iterative,
closed-incomplete/pragmatic; merged `--no-ff` into `release/02.50-the-playbill`. Built + unit-proven, **0 platform
edits**: the provably-clean 13-session simulation matrix + all 3 non-simulation sections (`content_nonsim.go`) →
manifest 4 products/18 sessions, both honesty gates GREEN. 2 user-blockers resolved [M235-01 scrub-leaked-zero-names
→ hardened+re-captured; M235-02 coverage-descriptor-doesn't-exist → build-non-sim-then-close]. Go **1939→1974**, flake
**5/5**, harden Pass1+2 stabilized. Close near-clean [adversarial subsection + 2 `(#M235-B1/B2)` back-refs, no rext code
change; `60eff14`], deferral audit **YELLOW**/0-blockers. carry-forward.md 3 clusters all **Fate-3 → M236** (applied,
`54eaefe`). The LIVE `(session×action)`-lands gate is **M236 by design**. **NEXT: M236 prove-on-billion — the FINAL
v2.5 milestone; then `/developer-kit:close-release`.**)_
