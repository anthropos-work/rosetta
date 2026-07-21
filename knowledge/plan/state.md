---
active_release: "v2.6 «sound check» — the reliability / field-hardening release (designed 2026-07-20): make everything that's built actually get built + provisioned. Barrier→parallel-fixes→prove-on-billion, 8 milestones M237→M244. Branch release/02.60-sound-check; tag will be v2.6."
active_branch: "release/02.60-sound-check"
active_milestone: "M238 ∥ M239 ∥ M240 ∥ M243 — the post-barrier fix fan-out (M237 barrier CLOSED 2026-07-21); M238 ant-academy reliability leads"
last_closed: "M237 — 2026-07-21 (v2.6 clean-stage barrier)"
phase: "v2.6 in development — M237 (HARD go/no-go barrier) CLOSED: fetch-verified clone-freshness + 7-state pin model + R1-all-14, both dogfooded on billion; the '202-behind' premise REFUTED (clones 0-2 behind, frontend current). Fan-out M238∥M239∥M240∥M243 now unblocked; #2 academy-language→M238, #4 library-flash→M239 (Fate-2 owned)."
last_updated: "2026-07-21"
---

# State

**v2.6 "sound check" — IN DEVELOPMENT.** The **reliability / field-hardening release** (the v1.3b / v1.10b / v2.1 /
v2.3 lineage), designed 2026-07-20, triggered by **live demo defects** — *"still not all gets built and provisioned as
expected."* The job: make everything that's *built* actually *build + provision* on a fresh box. House shape **barrier →
parallel fixes → prove-on-billion**. **8 milestones M237 → M244**, branch `release/02.60-sound-check` (cut from **local**
`main`); tag will be **`v2.6`**. **Tooling + docs only — zero platform-repo edits.**

```
M237 clean stage (HARD go/no-go barrier)
  ├─▶ M238 academy reliability ─────┐
  ├─▶ M239 enterprise surfaces ─────┤
  ├─▶ M240 content-fidelity ─┐      │   (HARD media-safety gate)
  │      └─▶ M241 language ─┐ │      │
  │            └─▶ M242 cockpit-UX
  ├─▶ M243 assign-WRITE ────────────┤   [realizes reserved M238]
  └─────────────────────────────────▶ M244 prove-on-billion (iterative closer) [realizes reserved M237]
```

- **M237 clean stage** [`section`, HARD go/no-go] — ✅ **CLOSED 2026-07-21.** Fetch-verified clone-freshness in
  `ensure-clones.sh` (never suppressed-stderr; a real 7-state pin model) + F-M236-CLOSE-2 (R1 sweeps all 14 manifests);
  fresh-clone demo dogfooded on `billion` + a confirmed-defect ledger (#1 RESOLVED, #4→M239 re-scoped, #2→M238). The
  "202-behind" premise REFUTED — clones were 0-2 behind, frontend current.
- **M238 academy reliability** · **M239 enterprise surfaces** · **M240 content-stories fidelity** · **M243 assign-WRITE
  Playthrough** — the parallel fix fan-out off M237.
- **M241 language** (serial after M240) · **M242 cockpit UX** (serial after M241).
- **M244 prove-on-billion** [`iterative`, terminal] — re-prove v2.5's headline `29/29` AND every v2.6 fix live on
  `billion`, cold reset-to-seed. Multi-part exit gate (a–h), opens with the read-only `ORG-CLEAN` settling check.

**3 binding user decisions (2026-07-20):** **(1) talk-to-data → FULL** — real AWS Bedrock creds via `/stack-secrets` +
a secret-coverage DNA extension for `app` (not just a flag; `../hyper-studio/.env.example` key set) [M239]; **(2) media
→ PORT IT** — capture + re-host the Chime/S3 voice recording + document blobs, behind a **HARD internal PII gate**
(fresh data-controller sign-off + a `safety.md` raw-media amendment + a voice/document anonymization decision — a voice
cannot be token-scrubbed — **before any customer audio lands in a demo**) [M240]; **(3) language → EN-only fallback per
tuple** — M241 opens with a read-only IT-session pool-count go/no-go query [M241].

**Next:** the **post-barrier fix fan-out**, now unblocked by the M237 close — **M238 ant-academy reliability ∥ M239
enterprise surfaces ∥ M240 content-stories fidelity ∥ M243 assign-WRITE Playthrough** (M241 serial after M240, M242 after
M241). Each fix is now scoped against the fresh, correctly-built demo M237 established. The v2.5 headline live re-prove
stays reserved for the terminal **M244** (which opens with the read-only `ORG-CLEAN` settling check, standing backlog below).

## ⚠️ The v2.5 headline shipped UNIT-PROVEN, not LIVE-RE-PROVEN — v2.6/M244 re-proves it

**`29/29` is UNIT-PROVEN, NOT LIVE-RE-PROVEN.** It was measured live on `billion` at rext tag
`playbill-m236-hardened`; the close then fixed **~10 defects in that same harness**, incl. a **grader that mis-shapes
one of the 29 pairs** and **four unit-spec files no runner executes** (among them the only literal pin of `29`). The
measuring instrument changed after the measurement. Also unverified: **39 live-browser specs** (24 stack-verify + 15
playthroughs) — the whole prove-by-render layer. **This is v2.6's `M244` job** (the realized reserved `M237`): `billion`
is up and reachable. Ledger: `CLOSE-D3` + `T-3`.

## Recently closed (milestones, newest first — max 5)

- **M237 clean stage** — 2026-07-21 (section, HARD go/no-go barrier; **opens v2.6**). Fetch-verified clone-freshness +
  7-state pin model + R1-all-14-manifests (F-M236-CLOSE-1/-2), dogfooded on `billion`. The **"202-behind" premise
  REFUTED** — clones 0-2 behind, frontend current. Ledger: #1 RESOLVED, #4→M239, #2→M238. 160 tests, 0 platform edits.
- **M236 prove-on-billion** — 2026-07-20 (iterative, closed-on-gate). **FINAL v2.5 milestone.** Gate MET cold on
  `billion`: **29/29** landable (session × action) pairs both vantages · **65** academy cards, 0 Draft chips · hero
  p95 **1.22 / 1.51 s** vs 5 s · cold reset-to-seed, no intervention · **0 platform edits**. Denominator CORRECTED
  31 → 29.
- **M235 prove-it-lands** — 2026-07-20 (iterative, closed-incomplete/pragmatic; LIVE gate → M236 by design). The
  13-session simulation matrix + all 3 non-sim sections → manifest **4 products / 18 sessions**, honesty gates GREEN.
- **M234 content-stories-cockpit-tab** — 2026-07-19 (section). The 2nd "Content stories" tab reads the M233
  manifest; `roster.go` appends `content-player-<idx>` seats.
- **M233 content-stories-manifest** — 2026-07-19 (section). `BuildContentProducts` projects `content_products[]`,
  honesty-gated, fail-closed. Emitted by `stackseed --content-export`.

## Recently shipped (releases, newest first — max 3)

- **v2.5 "the playbill"** — 2026-07-20 (tag `v2.5`). The **content-vantage** release: **29/29** landable
  (session × action) pairs live on `billion` both vantages + the academy grid filled (65 cards), **0 platform edits**
  across 8 milestones. ⚠️ **The 29/29 is UNIT-PROVEN, NOT LIVE-RE-PROVEN** — the harness was fixed ~10× after the
  measurement; the live re-prove is v2.6's `M244` (realized reserved `M237`).
- **v2.4 "casting call"** — 2026-07-18 (tag `v2.4`). The **recruiter-vantage / hiring-org** release: a 4th hiring org
  on the cockpit (45 candidates on 5 shared positions), proven live on `billion` (M228 7/7, recruiter p95 1.27 s).
- **v2.3 "cue to cue"** — 2026-07-15 (tag `v2.3`). The **presenter-speed** release: click→ACCESS < 5 s proven live
  8/8 on `billion` — login p95 2.11 s / 1.31 s vs a ~39/38 s baseline. Remote default-on.

## Headline numbers (v2.5, as measured at close 2026-07-20 — the v2.6 baseline)
- **Go:** **1976** test funcs (`git grep '^func Test'`, whole rext repo). **2461 testcases / 0 failed** across 6 modules.
- **Python:** **1409 testcases** — 1399 pass / 8 deterministic / 2 flake / 8 skip.
- **TypeScript/Playwright:** **196 unit specs executed / 0 failed**; **39 live-browser specs (24 + 15) NOT executed**
  (see the caveat above — v2.6/M244 executes them).
- **p95 click→ACCESS:** employee **1.22 s** · manager **1.51 s** (M236, COLD on `billion`), vs the < 5000 ms gate.
- **Flake:** **0.** **Alignment (Clerkenstein):** **100% / 100% critical**. **Platform-repo edits:** **0**. **Supply
  chain:** GREEN — 0 net-new deps (v2.6 M239 adds a Bedrock **secret class**, not a dep).

## D17 — the carried-forward signature hazard
**D17: *a status artifact that outlives the thing it describes, and is then read as evidence.*** **The keeper:**
***a named hazard is not a fence; only an executable probe binds.*** **v2.5's sibling thesis:** ***a check can report
success while proving nothing*** — the class was found alive inside itself (the headline caveat). v2.6/M244's gate is
built to bind it: `ORG-CLEAN` runs FIRST, the 39 specs must actually execute, `DEF-M226-01` must be TESTED or DROPPED.

## Branch model / shipped tags
**v2.6** branch `release/02.60-sound-check` cut from **local** `main` 2026-07-20. ⚠️ **v2.5's `release → main` merge +
`v2.5` tag are LOCAL-ONLY** — `main` + tag not yet pushed to origin (R5; flag to user, do not auto-push). rext
code-of-record at v2.5 close **`playbill-m236-close-fixes`** (on origin). The `billion` demo LEFT UP. **Shipped tags:**
**v2.5** · **v2.4** · **v2.3** · **v2.2** · **v2.1** · **v2.0** · **v1.10b** `v1.10.1` · **v1.10** · **v1.9** · **v1.8** ·
**v1.7** · **v1.6** · **v1.3b** `v1.3.1` · **v1.3** · **v1.2** · **v1.1** · **v1.0**. (Detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog (unscheduled, cross-release)

> **Every item has a FATED destination.** Full per-item reasoning:
> [`releases/archive/02.50-the-playbill/release-deferrals.md`](releases/archive/02.50-the-playbill/release-deferrals.md).
> **Class named at v2.5 close: a fate needs a MILESTONE — not "a sweep", "the next close", or "standing backlog".**
> **v2.6 remap:** the reserved `M237` re-prove is **realized as `M244`**; the reserved `M238` assign-WRITE is
> **realized as `M243`**.

- **Standing demo-stack test debt — RE-BASELINED.** **8 on macOS · 7 expected on Linux** (clean stable-`main` clone
  set), **0 real defects**, **0 pin drift**. **Host-dependent — always state the host.** →
  `releases/archive/02.50-the-playbill/m236-prove-on-billion/rebaseline-standing-failures.md`.
- **→ `M244` (v2.6's live closer — one live bring-up discharges most of these).** `M237` opens with the read-only
  **`ORG-CLEAN`** settling check (13 copied session exhibits — resolve source-org names via one prod query + an offline
  `scrub.OrgTokens`/`SurvivingToken` pass over the committed fixtures; 0 surviving tokens or each dispositioned;
  VPN-scoped + data-controller-accepted). `M244` then discharges: `CLOSE-D3` (the 29/29 live re-prove) · the **39
  unexecuted live-browser specs** · `ACADEMY-M236-iter08-public-catalog-twin` (anon `/library` + `/free` render 0 cards)
  · the `apps/web` non-offset `:5050` client GraphQL endpoint · **`DEF-M226-01`** (pre-bind serve reap — AGED OUT TWICE;
  `M244` must **TEST** the "self-resolves" claim or the item is **DROPPED**) · **`BURNIN-M221-dev-public-host` ·
  `F-M220-4` · `PROBE-M218-c3-rerun`** (the v2.3 `DRIFT_DEFER` tail) · the M232 **interview plan-section** alignment
  assertion.
- **→ `M243` (assign-WRITE):** **`DEF-M235-03`** / M204 **assign-WRITE** in-manifest TODO — ~10 routings across 5
  releases; **fresh dated KEEP-DEFERRED-WITH-SIGNOFF 2026-07-20**. **Expiry: if `M243` does not land it, DROP.**
- **→ next `stack-seeding` build-iter:** **`DEF-M215-03(a)` / `F11`** — seed hero identity-key vs generated profile
  display-name mismatch (cosmetic). Enumerated by id so it stays findable.
- **Older, still unscheduled (re-confirmed 2026-07-20):** **DEF-M10-01** (cloud SnapshotStore / S3 blob bytes — **now
  likely CONSUMED by M240's media-porting**, per user decision 2), DEF-M21-01 (`replayCmd` hermetic test), CAVEAT-1
  (clean-box literal full `/dev-up`), M314b (prod frozen-read hydration), **M205**-residual (tier gates + ATS).
  Playthroughs futures **M206–M207** stay in vision.

_Last updated: 2026-07-21 — v2.6 "sound check" IN DEVELOPMENT (branch `release/02.60-sound-check`); the **M237
clean-stage barrier is CLOSED**, the M238∥M239∥M240∥M243 fix fan-out is unblocked._
