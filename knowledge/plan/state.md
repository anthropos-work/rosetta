---
active_release: "v2.5 «the playbill» — the content-vantage release (designed 2026-07-19): fill the empty ant-academy (Thread A) + a 2nd «Content stories» cockpit tab of played sessions per content product with as-player/as-manager login-and-land, cloned from anonymized real prod sessions, VPN-scoped, source-pinned. 8 milestones M229→M236, spike-first. Branch release/02.50-the-playbill; tag will be v2.5."
active_branch: "release/02.50-the-playbill"
active_milestone: "(between milestones — v2.5 M229→M236 all closed AND MERGED; NEXT is /developer-kit:close-release, UNBLOCKED)"
last_closed: "M236 — 2026-07-20"
phase: "v2.5 COMPLETE + IN CLOSE. All 8 milestones M229→M236 closed and MERGED. M236's gate MET cold on billion (29/29 landable pairs both vantages · 65 academy cards, 0 Draft chips · hero p95 1.22/1.51 s vs 5 s · cold reset-to-seed, no intervention · 0 platform edits). Denominator CORRECTED 31→29 (2 skill-path manager pairs target a next-web surface that renders 'Coming soon'), which also exposed a false PASS. The deferral RED that held M236's merge is DISCHARGED: user fate = RE-BASELINE the standing test-debt carry, decide at release close. Re-baselined — the carried 14 is a DIRTY-clone reading; on a clean stable-main clone set it is 8 on macOS / 7 expected on Linux, 0 real defects, 0 pin drift (that diagnosis REFUTED). /developer-kit:close-release is RUNNING: Phase 1b RED (22 items, 8 aged-out, 3 unhomed) is discharged by the release-scope deferral ledger at releases/02.50-the-playbill/release-deferrals.md. HEADLINE CAVEAT, recorded honestly: the 29/29 is UNIT-PROVEN, NOT LIVE-RE-PROVEN — the harness was fixed ~10 times AFTER the measurement, and the live re-prove is deferred to v2.6 as its first work (reserved M237) by explicit user decision."
last_updated: "2026-07-20"
---

# State

**v2.5 "the playbill" — FEATURE-COMPLETE, IN RELEASE CLOSE** (designed 2026-07-19; branch
`release/02.50-the-playbill` cut from `main`; tag will be `v2.5`). The **content-vantage release** — two threads on the
mature demo/cockpit machinery: **A** fills the empty **ant-academy** grid (it rendered 0 cards because the catalog is
DB-authoritative and a demo neither sets the GraphQL endpoint nor holds academy rows → `emptyCatalogView()`); **B**
adds a 2nd **"Content stories"** cockpit tab listing **played sessions** per content product (Simulation · Skill-path
legacy · Skill-path new · AI-labs), each with **as-player / as-manager** login-and-land actions, cloned from
**anonymized real production sessions**, non-manager-played, re-tenanted, **source-pinned by prod session-id**.
**User decisions (2026-07-19):** real customer-session sourcing accepted (data-controller call) · demos
**VPN/tailnet-scoped**, release **amends `safety.md` Part 3** · academy fill **production-faithful** (no "Draft" chip)
· AI-labs + academy section **scoped by the M231 spike**. Tooling + docs only, **0 platform-repo edits**.

## Where we are — the close is RUNNING, nothing is blocked

All 8 milestones (M229→M236) are **closed and merged** into `release/02.50-the-playbill`. **M236 met its gate cold on
`billion`.** The merge-holding deferral RED is **discharged**: the user's fate (2026-07-20) was *re-baseline the
standing test-debt carry now, decide at release close* — executed, and the decision is taken (below).

`/developer-kit:close-release` is in progress. Phase 1b returned **RED** (22 items · 8 aged-out · 3 unhomed); it is
discharged by the **release-scope deferral ledger**
[`releases/02.50-the-playbill/release-deferrals.md`](releases/02.50-the-playbill/release-deferrals.md) — every item
fated, with its concrete why-Fate-1/2/3-failed, a named handler, and a destination-still-valid check. Remaining:
`release → main` merge + the `v2.5` tag (Phase 11, the USER runs it).

## ⚠️ The headline number ships unverified-live — read this before quoting it

**`29/29` is UNIT-PROVEN, NOT LIVE-RE-PROVEN.** It was measured live on `billion` at rext tag
`playbill-m236-hardened`; the close then fixed **~10 defects in that same harness**, and the close review found more —
including a **grader that mis-shapes one of the 29 pairs** and **four unit-spec files no runner executes** (among them
the only literal pin of `29`). **The measuring instrument changed after the measurement, and has never faced a
browser since.** Also unverified: **39 live-browser specs** (24 stack-verify + 15 playthroughs) — the whole
prove-by-render layer. This is a **conscious user decision** (tag now, re-prove as v2.6's first work), not an
impossibility: `billion` is up and reachable. Record: `CLOSE-D3` + `T-3` in the ledger; destination **reserved
`M237 — re-prove-on-billion`**.

## Active release — v2.5 (8 milestones, spike-first)

**Shape:** `M229 → M230 → M231 (HARD go/no-go) → M232 → M233 → M234 → M235 → M236`. Full designs + the
safety/decision record: `roadmap.md` § Active — v2.5. Shipped-release detail moved to
[`roadmap-archive-v2.0-v2.4.md`](roadmap-archive-v2.0-v2.4.md) at this close. **Hard constraint:** zero
platform-repo edits.

## Recently closed (milestones, newest first — max 5)

- **M236 prove-on-billion** — 2026-07-20 (iterative, **closed-on-gate**, merged). **FINAL v2.5 milestone.**
  10 iters (1 bootstrap tok + 9 tiks, one day). Gate MET cold on `billion`: **29/29** landable (session ×
  action) pairs both vantages · **65** academy cards, 0 Draft chips · hero p95 **1.22 s / 1.51 s** vs 5 s ·
  cold reset-to-seed, no intervention · **0 platform edits**. **Denominator CORRECTED 31 → 29** — the 2
  skill-path manager pairs target a next-web surface that renders "Coming soon" (table commented out,
  `userData` null); the correction also exposed a **false PASS** off a definition-only header. Close fixed
  39 findings, incl. **5 unnamed stack-core failures** (both cross-repo doc-truth guards red and *correct*
  for 3 milestones) and a **tautological** membership-key test that could not fail. Two upstream findings
  raised: **F-M236-CLOSE-1** (`/demo-up` rebuilds images from clones it NEVER updates — `app` 249 commits
  behind `main`, `next-web-app` 202, identically on both boxes; this was the user-reported stale left menu
  and the generator of the whole pin-drift-looking class) and **F-M236-CLOSE-2** (the R1 pristine sweep
  covers 3 manifests of ~15). rext `playbill-m236-close-fixes`.
- **M235 prove-it-lands** — 2026-07-20 (iterative, closed-incomplete/pragmatic; LIVE gate → M236 by design).
  Built + unit-proven, **0 platform edits**: the 13-session simulation matrix + all 3 non-simulation sections
  via `content_nonsim.go` → manifest **4 products / 18 sessions**, honesty gates GREEN. rext
  `playbill-m235-hardened @ 60eff14`. carry-forward 3 clusters all Fate-3 → M236.
- **M234 content-stories-cockpit-tab** — 2026-07-19 (section, closed-complete). The **render half**: the 2nd
  "Content stories" tab reads the M233 manifest; `roster.go` appends `content-player-<idx>` seats.
- **M233 content-stories-manifest** — 2026-07-19 (section, closed-complete). The **manifest half**:
  `BuildContentProducts` projects `content_products[]` single-sourced from the M232 fixture, honesty-gated,
  **fail-closed**. Emitted by `stackseed --content-export`.
- **M232 session-clone-sourcing-seeder** — 2026-07-19 (section, closed-complete). The `ContentStorySeeder`
  **copies real prod sessions** + best-effort PII scrub + re-tenant + source-pin; interview flags via 2
  demopatches; `safety.md` §3.8 = data-controller-accepted residual risk, VPN-scoped.

## Recently shipped (releases, newest first — max 3)

- **v2.4 "casting call"** — 2026-07-18 (tag `v2.4`). The **recruiter-vantage / hiring-org** release: a 4th hiring org
  on the cockpit (45 candidates on 5 shared positions), proven live on `billion` (M228 7/7, recruiter p95 1.27 s).
  ⚠️ **Its close issued false greens and ran no release-scope deferral audit** — the structural cause of 7 of v2.5's
  8 aged-out items; its archived `release-review.md` is annotated in place, not rewritten.
- **v2.3 "cue to cue"** — 2026-07-15 (tag `v2.3`). The **presenter-speed** release: click→ACCESS < 5 s proven live 8/8
  on `billion` — login p95 2.11 s / 1.31 s vs a ~39/38 s baseline (~18×). Remote default-on.
- **v2.2 "panorama"** — 2026-07-12 (tag `v2.2`). External-shareability / Tailscale-serve. First live remote
  Linux-VM deploy.

## Headline numbers (v2.5, as measured at close 2026-07-20)
- **Go:** **1976** test funcs (`git grep '^func Test'`, whole rext repo; +97 vs v2.4's reproducible **1879**).
  **2461 testcases / 0 failed** across 6 modules via JUnit XML.
- **Python:** **1409 testcases** — 1399 pass / **8 deterministic fail** / 2 flake / 8 skip, 5 sections.
- **TypeScript/Playwright:** **196 unit specs executed / 0 failed** (e2e 127 + playthroughs 69); **235
  collected** — **39 live-browser specs (24 + 15) NOT executed** (see the caveat above).
- **p95 click→ACCESS:** employee **1.22 s** · manager **1.51 s** (M236, COLD on `billion`) · recruiter 1.27 s
  (M228), all vs the < 5000 ms gate. *The cold stack measured ~2× faster than the warm one it replaced —
  iter-09's 3.15/2.71 s are the pessimistic pair; 1.2 s is not steady state.*
- **Alignment (Clerkenstein):** **100% / 100% critical** (not re-scored this close).
- **Flake:** **2** — first non-zero since v1.0; both pre-existing test-infra timing flakes, neither in v2.5 code.
  **Platform-repo edits:** **0** (verified per-clone). **Supply chain:** GREEN — 0 net-new deps.

## D17 — the carried-forward signature hazard

**D17: *a status artifact that outlives the thing it describes, and is then read as evidence.*** **The keeper:**
***a named hazard is not a fence; only an executable probe binds.*** Full arc:
[`releases/archive/02.30-cue-to-cue/release-retro.md`](releases/archive/02.30-cue-to-cue/release-retro.md).
**v2.5's own thesis is the sibling:** ***a check can report success while proving nothing*** — nine instances across
M235–M236, and the close found the class **alive inside itself** (see the headline caveat).

## Branch model / shipped tags
**v2.5 M229→M236 all merged** `--no-ff` into `release/02.50-the-playbill` (cut from `main` 2026-07-19). `release →
main` + the `v2.5` tag are `/developer-kit:close-release` Phase 11's job (the USER runs it). rext code-of-record
**`playbill-m236-close-fixes`** (pushed to origin — M217's FATAL pin guard means a tag is only consumable once it is
**on origin**, not merely created). The `billion` demo LEFT UP. **Shipped tags:** **v2.4** · **v2.3** · **v2.2** ·
**v2.1** · **v2.0** · **v1.10b** `v1.10.1` · **v1.10** · **v1.9** · **v1.8** · **v1.7** · **v1.6** · **v1.3b**
`v1.3.1` · **v1.3** · **v1.2** · **v1.1** · **v1.0**. (Detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog (unscheduled, cross-release)

> **Every item below has a FATED destination.** Full per-item reasoning, handler, and destination-still-valid check:
> [`releases/02.50-the-playbill/release-deferrals.md`](releases/02.50-the-playbill/release-deferrals.md).
> **Class named at this close: "an arch-doc pass" is not a milestone — neither is "a sweep", "the next close", or
> "standing backlog". A fate needs a MILESTONE.**

- **Standing demo-stack test debt — RE-BASELINED, no longer a mystery.** **8 on macOS · 7 expected on Linux**
  (clean stable-`main` clone set), **0 real defects**, **0 `pre_sha256` pin drift** — that diagnosis is **REFUTED**,
  and its implied remedy would have re-pinned five manifests to *patched* content. The carried **14** was a
  **dirty-clone** reading (6 leftover applied demo patches reporting as failures; they did not reproduce at this
  close, confirming the `stack-demo` clone set is pristine). **The count is host-dependent — always state the host.**
  → `m236-prove-on-billion/rebaseline-standing-failures.md`.
- **→ reserved `M237` (v2.6's FIRST work — one live bring-up discharges all of these):** `CLOSE-D3` (the 29/29 live
  re-prove) · the **39 unexecuted live-browser specs** · `ACADEMY-M236-iter08-public-catalog-twin` (anonymous
  `/library` + `/free` render 0 cards — **this is the surviving half of `F4`; the signed-in half LANDED at M230/M236,
  so `F4` is retired as an id**) · the `apps/web` non-offset `:5050` client GraphQL endpoint · **`DEF-M226-01`**
  (pre-bind serve reap — **AGED OUT TWICE**, M228 then M236, because its destination was the *phrase* "the next
  prove-on-VM"; M237 must **TEST** the "self-resolves in the default flow" claim, or the item is **DROPPED**) ·
  **`BURNIN-M221-dev-public-host` · `F-M220-4` · `PROBE-M218-c3-rerun`** (the v2.3 tail carries, reclassified
  `DRIFT_DEFER` — their "needs live infra" rationale is **now false**) · the M232 **interview plan-section
  alignment** assertion (a Fate-2 "confirmed-covered" whose covering gate never ran).
- **→ reserved `M238`:** **`DEF-M235-03`** / M204 **assign-WRITE** in-manifest TODO — ~10 routings across 5
  releases; **fresh dated KEEP-DEFERRED-WITH-SIGNOFF 2026-07-20**. **Expiry: if `M238` is not designed into v2.6 or
  v2.7, DROP it.**
- **→ next `stack-seeding` build-iter:** **`DEF-M215-03(a)` / `F11`** — seed hero identity-key vs generated
  profile display-name mismatch (cosmetic). Routed to "standing backlog" at the v2.2 close and **never written into
  one** — invisible for three releases until this close found it. Now enumerated by id.
- **Older, still unscheduled (re-confirmed 2026-07-20):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes),
  DEF-M21-01 (`replayCmd` hermetic test), CAVEAT-1 (clean-box literal full `/dev-up`), M314b (prod frozen-read
  hydration), **M205**-residual (tier gates + ATS). Playthroughs futures **M206–M207** stay in vision. All tracked
  in [`roadmap-vision.md`](roadmap-vision.md).

_Last updated: 2026-07-20 — v2.5 feature-complete, all milestones merged, `/developer-kit:close-release` running._
