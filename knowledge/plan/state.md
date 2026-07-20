---
active_release: "v2.5 «the playbill» — the content-vantage release (designed 2026-07-19): fill the empty ant-academy (Thread A) + a 2nd «Content stories» cockpit tab of played sessions per content product with as-player/as-manager login-and-land, cloned from anonymized real prod sessions, VPN-scoped, source-pinned. 8 milestones M229→M236, spike-first. Branch release/02.50-the-playbill; tag will be v2.5."
active_branch: "release/02.50-the-playbill"
active_milestone: "(between milestones — v2.5 M229→M236 all closed AND MERGED; NEXT is /developer-kit:close-release, UNBLOCKED)"
last_closed: "M236 — 2026-07-20"
phase: "v2.5 COMPLETE — all 8 milestones M229→M236 closed. M236 prove-on-billion gate MET cold on billion (29/29 landable pairs both vantages · 65 academy cards, 0 Draft chips · hero p95 1.22/1.51 s vs 5 s · cold reset-to-seed, no intervention · 0 platform edits). Denominator CORRECTED 31→29 (2 skill-path manager pairs target a next-web surface that renders 'Coming soon'), which also exposed a false PASS. Close fixed 39 findings incl. 5 unnamed stack-core doc-truth-guard failures + a tautological membership-key test. rext playbill-m236-close-fixes. ✅ M236 MERGED 2026-07-20 after the close continuation discharged the deferral RED: user fate = RE-BASELINE the standing test-debt carry now, decide at release close. Re-baselined — the carried 14 reproduces exactly, falls to 8 on a clean stable-main clone set, 0 are real defects and 0 are pin drift (that diagnosis is REFUTED; 6 were a dirty clone carrying leftover demo patches, and the remedy the old label implied would have re-pinned 5 manifests to patched content). See m236/rebaseline-standing-failures.md + carry-forward.md. Two upstream findings raised: F-M236-CLOSE-1 — /demo-up rebuilds images from clones it NEVER updates (make init is skip-if-present; app was 249 commits behind main, next-web-app 202, identically on both boxes; this was the user-reported stale left menu and the generator of the whole pin-drift-looking class) and F-M236-CLOSE-2 — the R1 pristine sweep covers 3 manifests of ~15. Cold bring-up at stable main on billion: autoverify GREEN, 0 patches refused. Deferral audit now YELLOW."
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

## Between milestones — v2.5 is feature-COMPLETE; the release close is BLOCKED

All 8 milestones (M229→M236) are closed. **M236 met its gate cold on `billion`** and is the final one.
**It has not merged.** `/developer-kit:audit-deferrals` returned **RED**: the standing **14 pre-existing
demo-stack test failures** are a genuine repeat-deferral across **10 milestones and 2 releases**, and their
declared destination — the **v2.4 release close** — already fired once without landing them (an AGED_OUT
trigger no audit had recorded). There is no further milestone to defer into. **Blocked pending an explicit
user fate: LAND-NOW / DROP / KEEP-DEFERRED-WITH-SIGNOFF (fresh dated decision).** The set must be
**re-baselined** first — it drifted 8 → 14 under a fixed label and the class changed (stale tests →
`pre_sha256` pin drift), so 14 is wrong in both directions. See
`releases/02.50-the-playbill/m236-prove-on-billion/decisions.md` **CLOSE-D2**.

Once fated: merge `m236/prove-on-billion` → `release/02.50-the-playbill`, then `/developer-kit:close-release`
does the `release → main` merge + the `v2.5` tag.

## Active release — v2.5 "the playbill" (8 milestones, spike-first)

**Shape:** `M229 → M230 → M231 (HARD go/no-go) → M232 → M233 → M234 → M235 → M236`. M229 ∥ M231 research overlaps;
M230 (academy fill) must land before M235's academy section. Full milestone designs + the safety/decision record:
`roadmap.md` § Active — v2.5. **Hard constraint:** zero platform-repo edits (a runtime-computed result page that won't
render from a seeded row routes to a sha-pinned `demopatch` or escalates).

## Recently closed (milestones, newest first — max 5)

- **M236 prove-on-billion** — 2026-07-20 (iterative, **closed-on-gate**; merge HELD on the audit RED above).
  **FINAL v2.5 milestone.** 10 iters (1 bootstrap tok + 9 tiks, one day). Gate MET cold on `billion`:
  **29/29** landable (session × action) pairs both vantages · **65** academy cards, 0 Draft chips · hero p95
  **1.22 s / 1.51 s** vs 5 s · cold reset-to-seed, no intervention · **0 platform edits**. **Denominator
  CORRECTED 31 → 29** — the 2 skill-path manager pairs target a next-web surface that renders "Coming soon"
  (table commented out, `userData` null); the correction also exposed a **false PASS** off a definition-only
  header. Close fixed 39 findings, incl. **5 unnamed stack-core failures** (both cross-repo doc-truth guards
  — 4-orgs-since-M223 and the Thread-A-gating `DEMO_NO_ACADEMY_FILL` knob — red and *correct* for 3
  milestones) and a **tautological** membership-key test that could not fail. rext
  `playbill-m236-close-fixes`. Full narrative: roadmap `### M236`.
- **M235 prove-it-lands** — 2026-07-20 (iterative, closed-incomplete/pragmatic; LIVE gate → M236 by design).
  Built + unit-proven, **0 platform edits**: the 13-session simulation matrix + all 3 non-simulation sections
  via `content_nonsim.go` → manifest **4 products / 18 sessions**, honesty gates GREEN. 2 user-blockers
  resolved. rext `playbill-m235-hardened @ 60eff14`; Go 1939→1974. carry-forward 3 clusters all Fate-3 → M236.
- **M234 content-stories-cockpit-tab** — 2026-07-19 (section, closed-complete). The **render half**: the 2nd
  "Content stories" tab reads the M233 manifest (per-product sections, two login-and-land CTAs, AI-labs
  presence-only); `roster.go` appends `content-player-<idx>` seats. Python 249, Go +8, flake 5/5.
- **M233 content-stories-manifest** — 2026-07-19 (section, closed-complete). The **manifest half**:
  `BuildContentProducts` projects `content_products[]` single-sourced from the M232 fixture, honesty-gated so
  it can't drift, **fail-closed** (drop-with-reason, fails loud). Emitted by `stackseed --content-export`.
- **M232 session-clone-sourcing-seeder** — 2026-07-19 (section, closed-complete). The `ContentStorySeeder`
  **copies real prod sessions** + best-effort PII scrub + re-tenant + source-pin; interview flags via 2
  demopatches; `safety.md` §3.8 = data-controller-accepted residual risk, VPN-scoped.

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

## Headline numbers (v2.5 M236 close, 2026-07-20)
- **Go test funcs (whole rext repo, `git grep '^func Test'`):** **1976** (M236 +2 vs M235's **1974** — the
  membership-key contract pins). Full Go suite **2459 pass / 0 fail** across 6 modules / 58 packages.
- **Python:** **1391 pass / 2 fail / 8 skip.** stack-verify **132 → 141** (+6 runner/networkidle pins, +3
  e2e collection-integrity guards). stack-core **147/5 → 150/2** — the 5 doc-truth-guard failures FIXED at
  the M236 close; the 2 remaining (`test_m220_mutation_battery`) verified pre-existing at
  `playbill-m236-hardened` via a detached worktree.
- **TypeScript/Playwright:** stack-verify e2e **127** unit + playthroughs e2e **69**, `tsc` clean both.
  M236 harness specs **0 → 66**. Live-stack specs (24 + 15) not run without a demo up.
- **p95 click→ACCESS (the standing gate):** employee **1.22 s** · manager **1.51 s** (M236, COLD on
  `billion`) · recruiter 1.27 s (M228), all vs the < 5000 ms gate. *The cold stack measured ~2× faster than
  the warm one it replaced — iter-09's 3.15/2.71 s are the pessimistic pair; 1.2 s is not steady state.*
- **Alignment (Clerkenstein):** **100% / 100% critical.**
- **Flake:** **0** (milestone-owned). **Platform-repo edits:** **0** (verified per-clone). **Supply chain:**
  GREEN — 0 net-new direct deps.
- **Standing carry (MEASURED, not copied):** **14** pre-existing demo-stack failures. Briefed as 14, measured
  as **19**, of which **5 were fixed** at the M236 close → back to exactly 14. **REPEAT-DEFERRAL, escalated.**

## D17 — the carried-forward signature hazard (v2.4 discipline)

**D17: *a status artifact that outlives the thing it describes, and is then read as evidence.*** **The keeper:** ***a
named hazard is not a fence; only an executable probe binds.*** M227 lived it again: the 4 fixes are proven by
deterministic executable fences (incl. the load-bearing real-`UsersSeeder` population write-path fence that catches a
silent revert), never by "the seed wrote the rows." Full arc:
[`releases/archive/02.30-cue-to-cue/release-retro.md`](releases/archive/02.30-cue-to-cue/release-retro.md).

## Branch model / shipped tags
**v2.5 M229→M235 merged** `--no-ff` into `release/02.50-the-playbill` (cut from `main` 2026-07-19).
**M236 is CLOSED-ON-GATE but NOT MERGED** — held on the deferral-audit RED above; its branch
`m236/prove-on-billion` is intact. Then `release → main` + the `v2.5` tag are `/developer-kit:close-release`
Phase 11's job (the USER runs it). rext code-of-record **`playbill-m236-close-fixes`** (pushed to origin —
note M217's FATAL pin guard means a tag is only consumable once it is **on origin**, not merely created).
The `billion` demo LEFT UP. **Shipped tags:** **v2.4** `v2.4` · **v2.3** `v2.3` · **v2.2** `v2.2` · **v2.1**
`v2.1` · **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` ·
**v1.7** `v1.7` · **v1.6** `v1.6` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1`
· **v1.0** `v1.0`. (Full detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog (unscheduled, cross-release)
- **⛔ Test-debt (REPEAT-DEFERRAL — BLOCKS the v2.5 release close; needs a USER fate):** **14 pre-existing
  demo-stack failures** — 6 × `test_cockpit.py` + 2 × `test_demopatch.py` + 4 × `test_ssr_origin_chain.py` +
  `test_host_prereqs_m215` + `test_purge`; all environment-dependent (need a live clone tree / a
  container-owned data dir), in files M229–M236 never touched. **First seen M224 (2026-07-17) as *8*;
  carried through 10 milestones and 2 releases.** Its declared destination — *the v2.4 release close* —
  **fired on 2026-07-18 without landing it** (shipped as a known issue, re-anchored on v2.5). The set drifted
  8 → 14 under a fixed label and the *class* changed (stale-tests → `pre_sha256` pin drift, 6 of them), so
  **re-baseline against HEAD before choosing a fate.** Also folded in: `test_m220_mutation_battery` (2, the
  unmutated subject fails its own suite). Fate required: LAND-NOW / DROP / KEEP-DEFERRED-WITH-SIGNOFF.
  → `m236-prove-on-billion/decisions.md` **CLOSE-D2**.
- **v2.5 close queue (Fate 3, non-blocking):** `ACADEMY-M236-iter08-public-catalog-twin` (anonymous
  `/library` + `/free` render 0 cards) + the `apps/web` non-offset `:5050` client GraphQL endpoint — batch
  both into one next-web rebuild + one live re-prove; the **M204 assign-WRITE** declared in-manifest TODO;
  and a live re-run of the content-stories sweep against the close-modified harness (**CLOSE-D3** — the
  29/29 reading was taken on `playbill-m236-hardened`, before 10 close fixes landed in that harness;
  unit-proven, not live-re-proven).
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

_Last updated: 2026-07-20 — M236 closed-on-gate (merge held on the deferral-audit RED); v2.5 feature-complete._
