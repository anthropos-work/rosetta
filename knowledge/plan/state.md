---
active_release: "(between releases — v2.5 «the playbill» shipped 2026-07-20, tag v2.5; next release awaits /developer-kit:design-roadmap)"
active_branch: "release/02.50-the-playbill"
active_milestone: "(between releases)"
last_closed: "M236 — 2026-07-20"
phase: "between releases — awaiting /developer-kit:design-roadmap"
last_updated: "2026-07-20"
---

# State

**Between releases.** **v2.5 "the playbill" — CLOSED, tag `v2.5` (2026-07-20).** The **content-vantage release** —
two threads on the mature demo/cockpit machinery: **A** filled the empty **ant-academy** grid (65 real cards, 0 Draft
chips, via the real DB-authoritative path); **B** added a 2nd **"Content stories"** cockpit tab of **played sessions**
per content product, each with **as-player / as-manager** login-and-land, cloned from **anonymized real production
sessions**, re-tenanted, source-pinned. All 8 milestones (M229→M236) closed + merged into `release/02.50-the-playbill`;
tooling + docs only, **0 platform-repo edits**. The `release → main` merge + the `v2.5` tag are the final mechanical
close step (`/developer-kit:close-release` Phase 11, the USER runs it).

**Next:** `/developer-kit:design-roadmap` for **v2.6**, whose **declared first work** is reserved **`M237 —
re-prove-on-billion`** (the deferred live re-prove of v2.5's headline number + the 39 unexecuted live-browser specs +
the batched M237 carries — see [`roadmap-vision.md`](roadmap-vision.md) § v2.5 → v2.6 carry). **M237 opens with a
read-only `ORG-CLEAN` settling check** (standing backlog below).

## ⚠️ The headline shipped UNIT-PROVEN, not LIVE-RE-PROVEN — read this before quoting it

**`29/29` is UNIT-PROVEN, NOT LIVE-RE-PROVEN.** It was measured live on `billion` at rext tag
`playbill-m236-hardened`; the close then fixed **~10 defects in that same harness**, and the close review found more —
incl. a **grader that mis-shapes one of the 29 pairs** and **four unit-spec files no runner executes** (among them the
only literal pin of `29`). **The measuring instrument changed after the measurement, and has not faced a browser
since.** Also unverified: **39 live-browser specs** (24 stack-verify + 15 playthroughs) — the whole prove-by-render
layer. This is a **conscious user decision** (tag now, re-prove as v2.6's first work), not an impossibility: `billion`
is up and reachable. Ledger: `CLOSE-D3` + `T-3` → reserved **`M237`**.

## Recently closed (milestones, newest first — max 5)

- **M236 prove-on-billion** — 2026-07-20 (iterative, closed-on-gate). **FINAL v2.5 milestone.** Gate MET cold on
  `billion`: **29/29** landable (session × action) pairs both vantages · **65** academy cards, 0 Draft chips · hero
  p95 **1.22 / 1.51 s** vs 5 s · cold reset-to-seed, no intervention · **0 platform edits**. Denominator CORRECTED
  31 → 29 (2 skill-path manager pairs target a next-web "Coming soon" surface), which also exposed a false PASS.
- **M235 prove-it-lands** — 2026-07-20 (iterative, closed-incomplete/pragmatic; LIVE gate → M236 by design). The
  13-session simulation matrix + all 3 non-sim sections → manifest **4 products / 18 sessions**, honesty gates GREEN.
- **M234 content-stories-cockpit-tab** — 2026-07-19 (section). The 2nd "Content stories" tab reads the M233
  manifest; `roster.go` appends `content-player-<idx>` seats.
- **M233 content-stories-manifest** — 2026-07-19 (section). `BuildContentProducts` projects `content_products[]`,
  honesty-gated, fail-closed. Emitted by `stackseed --content-export`.
- **M232 session-clone-sourcing-seeder** — 2026-07-19 (section). `ContentStorySeeder` copies real prod sessions +
  best-effort PII scrub + re-tenant + source-pin; `safety.md` §3.8 = data-controller-accepted residual, VPN-scoped.

## Recently shipped (releases, newest first — max 3)

- **v2.5 "the playbill"** — 2026-07-20 (tag `v2.5`). The **content-vantage** release: **29/29** landable
  (session × action) pairs live on `billion` both vantages + the academy grid filled (65 cards), **0 platform edits**
  across 8 milestones. ⚠️ **The 29/29 is UNIT-PROVEN, NOT LIVE-RE-PROVEN** — the harness was fixed ~10× after the
  measurement; the live re-prove is v2.6's first work (reserved `M237`).
- **v2.4 "casting call"** — 2026-07-18 (tag `v2.4`). The **recruiter-vantage / hiring-org** release: a 4th hiring org
  on the cockpit (45 candidates on 5 shared positions), proven live on `billion` (M228 7/7, recruiter p95 1.27 s).
- **v2.3 "cue to cue"** — 2026-07-15 (tag `v2.3`). The **presenter-speed** release: click→ACCESS < 5 s proven live
  8/8 on `billion` — login p95 2.11 s / 1.31 s vs a ~39/38 s baseline. Remote default-on.

## Headline numbers (v2.5, as measured at close 2026-07-20)
- **Go:** **1976** test funcs (`git grep '^func Test'`, whole rext repo; +97 vs v2.4's reproducible 1879).
  **2461 testcases / 0 failed** across 6 modules via JUnit XML.
- **Python:** **1409 testcases** — 1399 pass / **8 deterministic** (standing test-side debt, 0 real defects) / 2 flake
  / 8 skip, 5 sections.
- **TypeScript/Playwright:** **196 unit specs executed / 0 failed**; **235 collected** — **39 live-browser specs
  (24 + 15) NOT executed** (see the caveat above).
- **p95 click→ACCESS:** employee **1.22 s** · manager **1.51 s** (M236, COLD on `billion`), vs the < 5000 ms gate.
- **Flake:** **0** — 2 surfaced at Phase 4b, FIXED test-side + verified 3/3 serial-green (not waived).
- **Alignment (Clerkenstein):** **100% / 100% critical** (not re-scored). **Platform-repo edits:** **0** (per-clone).
  **Supply chain:** GREEN — **0 net-new deps**.

## D17 — the carried-forward signature hazard

**D17: *a status artifact that outlives the thing it describes, and is then read as evidence.*** **The keeper:**
***a named hazard is not a fence; only an executable probe binds.*** Full arc:
[`releases/archive/02.30-cue-to-cue/release-retro.md`](releases/archive/02.30-cue-to-cue/release-retro.md).
**v2.5's own thesis is the sibling:** ***a check can report success while proving nothing*** — nine instances across
M235–M236, and the close found the class **alive inside itself** (see the headline caveat).

## Branch model / shipped tags
**v2.5 M229→M236 all merged** `--no-ff` into `release/02.50-the-playbill` (cut from `main` 2026-07-19). The `release →
main` merge + the `v2.5` tag are the final mechanical close step (Phase 11, the USER runs it). rext code-of-record
**`playbill-m236-close-fixes`** (on origin). The `billion` demo LEFT UP. **Shipped tags:** **v2.5** · **v2.4** ·
**v2.3** · **v2.2** · **v2.1** · **v2.0** · **v1.10b** `v1.10.1` · **v1.10** · **v1.9** · **v1.8** · **v1.7** ·
**v1.6** · **v1.3b** `v1.3.1` · **v1.3** · **v1.2** · **v1.1** · **v1.0**. (Detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog (unscheduled, cross-release)

> **Every item below has a FATED destination.** Full per-item reasoning, handler, and destination-still-valid check:
> [`releases/archive/02.50-the-playbill/release-deferrals.md`](releases/archive/02.50-the-playbill/release-deferrals.md).
> **Class named at this close: a fate needs a MILESTONE — not "a sweep", "the next close", or "standing backlog".**

- **Standing demo-stack test debt — RE-BASELINED.** **8 on macOS · 7 expected on Linux** (clean stable-`main` clone
  set), **0 real defects**, **0 `pre_sha256` pin drift** (that diagnosis REFUTED; the carried **14** was a
  **dirty-clone** reading). **The count is host-dependent — always state the host.** →
  `m236-prove-on-billion/rebaseline-standing-failures.md`.
- **→ reserved `M237` (v2.6's FIRST work — one live bring-up discharges most of these).** **FIRST task:** the
  read-only **`ORG-CLEAN`** settling check — resolve the 13 copied session exhibits' source-org names via one prod
  query + an offline `scrub.OrgTokens`/`SurvivingToken` pass over the already-committed fixtures (no re-capture, no
  names in any transcript; recurrence is structurally prevented, the residual is VPN-scoped + data-controller-
  accepted). Then: `CLOSE-D3` (the 29/29 live re-prove) · the **39 unexecuted live-browser specs** ·
  `ACADEMY-M236-iter08-public-catalog-twin` (anon `/library` + `/free` render 0 cards — the surviving half of `F4`) ·
  the `apps/web` non-offset `:5050` client GraphQL endpoint · **`DEF-M226-01`** (pre-bind serve reap — AGED OUT
  TWICE; M237 must **TEST** the "self-resolves" claim or the item is **DROPPED**) · **`BURNIN-M221-dev-public-host` ·
  `F-M220-4` · `PROBE-M218-c3-rerun`** (the v2.3 tail, reclassified `DRIFT_DEFER`) · the M232 **interview
  plan-section alignment** assertion.
- **→ reserved `M238`:** **`DEF-M235-03`** / M204 **assign-WRITE** in-manifest TODO — ~10 routings across 5 releases;
  **fresh dated KEEP-DEFERRED-WITH-SIGNOFF 2026-07-20**. **Expiry: if `M238` is not designed into v2.6 or v2.7, DROP.**
- **→ next `stack-seeding` build-iter:** **`DEF-M215-03(a)` / `F11`** — seed hero identity-key vs generated profile
  display-name mismatch (cosmetic). Enumerated by id so it stays findable.
- **Older, still unscheduled (re-confirmed 2026-07-20):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes),
  DEF-M21-01 (`replayCmd` hermetic test), CAVEAT-1 (clean-box literal full `/dev-up`), M314b (prod frozen-read
  hydration), **M205**-residual (tier gates + ATS). Playthroughs futures **M206–M207** stay in vision. All tracked
  in [`roadmap-vision.md`](roadmap-vision.md).

_Last updated: 2026-07-20 — v2.5 "the playbill" CLOSED (tag `v2.5`); between releases, awaiting `/developer-kit:design-roadmap` for v2.6._
