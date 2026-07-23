---
active_release: "v2.6 «sound check» — the reliability / field-hardening release (designed 2026-07-20): make everything that's built actually get built + provisioned. Barrier→parallel-fixes→prove-on-billion, 8 milestones M237→M244. Branch release/02.60-sound-check; tag will be v2.6. ALL 8 CLOSED — ready for /developer-kit:close-release."
active_branch: "release/02.60-sound-check"
active_milestone: "(between milestones — v2.6 COMPLETE: all 8 M237→M244 closed; next step is /developer-kit:close-release, run by the user)"
last_closed: "M244 — 2026-07-23 (prove-on-billion, iterative closed-on-gate 8/8)"
phase: "v2.6 COMPLETE — all 8 milestones M237→M244 CLOSED (M244 the iterative terminal closer: gate MET 8/8 live on billion, cold reset-to-seed, 0 platform edits). release/02.60-sound-check is ready for /developer-kit:close-release (release→main merge + v2.6 tag). No milestone is active."
last_updated: "2026-07-23"
---

# State

**v2.6 "sound check" — ALL MILESTONES CLOSED, ready for release close.** The **reliability / field-hardening release**
(the v1.3b / v1.10b / v2.1 / v2.3 lineage), designed 2026-07-20, triggered by **live demo defects** — *"still not all
gets built and provisioned as expected."* House shape **barrier → parallel fixes → prove-on-billion**. **8 milestones
M237 → M244**, branch `release/02.60-sound-check` (cut from **local** `main`); tag will be **`v2.6`**. **Tooling + docs
only — zero platform-repo edits.** All 8 closed; the next step is **`/developer-kit:close-release`** (release→main merge
+ the `v2.6` tag — the USER runs it).

```
M237 clean stage (HARD go/no-go barrier) ✅
  ├─▶ M238 academy reliability ✅
  ├─▶ M239 enterprise surfaces ✅
  ├─▶ M240 content-fidelity ✅ → M241 language ✅ → M242 cockpit-UX ✅
  ├─▶ M243 assign-WRITE ✅   [realizes reserved M238]
  └──────────────────────────▶ M244 prove-on-billion ✅ (iterative closer, gate MET 8/8) [realizes reserved M237]
```

- **M237–M243** — ✅ CLOSED (M237–M239 2026-07-21, M240–M243 2026-07-22). Barrier + 6 parallel fixes: clone-freshness +
  7-state pin model (M237); academy chapter-body FS-published demopatch fixing Start-404 + language (M238); talk-to-data
  FULL via real AWS Bedrock creds, proven live (M239); content-stories fidelity + voice presence-only (M240); EN/IT
  language axis (M241); cockpit render/CSS (M242); the FIRST MUTATING Playthrough `pt-assignment-assign` (M243).
  (Detail: roadmap.md.)
- **M244 prove-on-billion** [`iterative`, terminal, realizes reserved M237] — ✅ **CLOSED 2026-07-23, gate MET 8/8** live
  on `billion` cold reset-to-seed, **0 platform edits**: (a) ORG-CLEAN 0 tokens · (b) content-stories **47/47** of the
  49-pair denominator (2 voice player cells presence-only) · (c) the **40 live-browser specs** (24 stack-verify +
  **16/16 Playthroughs**) · (d) anon academy `/library`+`/free` twin · (e) serve-reap 7→0 · (f) 3 v2.3 drift-carries
  incl. BURNIN-M221 · (g) interview plan-section alignment · (h) all 6 v2.6 fixes + **p95 1.46 s / 1.31 s**. 27 iters
  (24 tiks / 3 toks). Close: 2 doc reconciliations, 0 code fixes; deferral audit YELLOW. (Detail: roadmap.md M244.)

**Next:** run **`/developer-kit:close-release`** — the release-level review + `release/02.60-sound-check → main` merge +
the **`v2.6`** tag. It also runs the terminal deferral audit that is the expiry for the two carries M244 routed forward
(standing-8 demo-stack test debt + DEF-M239-01 ENOSPC loud-build-fail).

## v2.5's headline `29/29` — NOW LIVE-RE-PROVEN by M244

The v2.5 close shipped `29/29` **unit-proven, not live-re-proven** (the harness was fixed ~10× after the measurement).
**v2.6/M244 discharged that debt:** the content-stories sweep ran **live on `billion`** at the grown denominator —
**47/47** landed of the 49 pairs (M241's EN/IT growth; 2 Bunny-absent voice **player** cells presence-only) — and the
**40 live-browser specs executed green** (24 stack-verify + 16/16 Playthroughs, 96 cases in one clean full run). The
measuring-instrument-changed-after-measurement hazard (`CLOSE-D3` / `T-3`, the D17 signature-hazard thesis) is closed.

## Recently closed (milestones, newest first — max 5)

- **M244 prove-on-billion** — 2026-07-23 (iterative, terminal, realizes reserved M237; closed-on-gate). Re-proved the
  whole v2.6 feature + v2.5's headline live on `billion` cold reset-to-seed, gate MET 8/8, **0 platform edits**. 47/47
  content-stories of 49 · 40 live specs (16/16 Playthroughs) · p95 1.46/1.31 s. 27 iters (24 tiks / 3 toks). A real
  iter-25 finding fixed durably: the demo image must compile the **pinned** ref, not the highest fetched tag. rext tag
  `sound-check-m244-content-sweep-robustness` @ `498b1a5`. Deferral audit YELLOW (standing-8 + DEF-M239-01 →
  close-release); 2 doc fixes, 0 code fixes. (Detail: roadmap.md M244.)
- **M243 assign-WRITE Playthrough** — 2026-07-22 (section, realizes reserved M238). The FIRST MUTATING Playthrough
  (`pt-assignment-assign`): a manager assigns a skill path with a deadline → a real `organization_assignments` row
  written + read back. Closed the ~10-routing `DEF-M235-03`/M204 carry (5 releases) as Fate-1; 15→16 live / 0 TODO.
  (Detail: roadmap.md M243.)
- **M242 cockpit UX** — 2026-07-22 (section). Render/CSS-only: tuple-regrouped Content-stories rows + header-resident
  tab selector (byte-identical invariant held) + role-tinted hero avatars (AA). Close landed 1 latent D3 fix. (Detail:
  roadmap.md M242.)
- **M241 content-stories language** — 2026-07-22 (section, pool-count go/no-go). Real per-session `cs.Language` (was
  hard-coded `english`; 11/13 pins Italian) + cockpit **EN|IT toggle** + fail-closed gate. Fixture 13→23, denom 29→49.
  (Detail: roadmap.md M241.)
- **M240 content-stories fidelity** — 2026-07-22 (section, HARD media-safety gate). 5 fidelity fixes + **voice
  presence-only** (`not_available` IS the deliverable); real-video exhibit → M244 (`DEF-M240-01`, dispositioned).
  (Detail: roadmap.md M240.)

## Recently shipped (releases, newest first — max 3)

- **v2.5 "the playbill"** — 2026-07-20 (tag `v2.5`). The **content-vantage** release: 29/29 landable (session × action)
  pairs + academy grid filled (65 cards), 0 platform edits across 8 milestones. **Its 29/29 is now LIVE-RE-PROVEN by
  v2.6/M244** (47/47 of the grown 49-pair denominator, live on `billion`).
- **v2.4 "casting call"** — 2026-07-18 (tag `v2.4`). The **recruiter-vantage / hiring-org** release: a 4th hiring org on
  the cockpit (45 candidates on 5 shared positions), proven live on `billion` (M228 7/7, recruiter p95 1.27 s).
- **v2.3 "cue to cue"** — 2026-07-15 (tag `v2.3`). The **presenter-speed** release: click→ACCESS < 5 s proven live 8/8
  on `billion` — login p95 2.11 s / 1.31 s vs a ~39/38 s baseline. Remote default-on.

## Headline numbers (v2.6, at M244 close 2026-07-23 — the release-close baseline)
- **Go:** **2005** platform test funcs (v2.6 release-cumulative M237–M243; M244 `delivers:none` = 0 platform code). 2461
  testcases / 0 failed across 6 modules. **Platform-repo edits: 0** (all 8 milestones).
- **Python (rext demo-stack):** full suite **839 pass / 8 standing fail** (was 9; M244's final harden fixed the one
  M244-introduced M215 fence). The standing-8 (6 `test_cockpit` academy/overlay + `test_public_host` port-13001 +
  `test_purge` docker-integration) are pre-M244 inherited, host-dependent, **0 real defects** → close-release.
- **TypeScript/Playwright:** the **40 live-browser specs EXECUTED GREEN live on `billion`** (24 stack-verify + 16/16
  Playthroughs, 96 cases; M244 — the v2.5-deferred execution). rext unit: playthroughs 85 + stack-verify 172.
- **content-stories:** **47/47** landed live of the 49-pair denominator (2 voice player presence-only).
- **p95 click→ACCESS (M244, COLD on `billion`):** employee **1.46 s** · manager **1.31 s**, vs the < 5000 ms gate.
- **Flake:** **0.** **Alignment (Clerkenstein):** **100% / 100% critical.** **Supply chain:** GREEN — 0 net-new deps.

## Branch model / shipped tags
**v2.6** branch `release/02.60-sound-check` (cut from **local** `main` 2026-07-20) — **all 8 milestones closed; ready for
`/developer-kit:close-release`** (release→main + `v2.6` tag). ⚠️ **v2.5's `release → main` merge + `v2.5` tag are
LOCAL-ONLY** — not yet pushed to origin (R5; flag to user, do not auto-push). rext code-of-record at M244 close
**`498b1a5`** (consumption tag `sound-check-m244-content-sweep-robustness`, on origin; origin/main synced). The `billion`
demo LEFT UP (green at the m244 pin). **Shipped tags:** **v2.5** · **v2.4** · **v2.3** · **v2.2** · **v2.1** · **v2.0** ·
**v1.10b** … · **v1.0**. (Detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog (unscheduled, cross-release)

> **Every item has a FATED destination.** Full per-item reasoning:
> [`releases/archive/02.50-the-playbill/release-deferrals.md`](releases/archive/02.50-the-playbill/release-deferrals.md).
> **Class named at v2.5 close: a fate needs a MILESTONE — not "a sweep", "the next close", or "standing backlog".**

- **→ `/developer-kit:close-release` (v2.6's release close, the terminal expiry).** Two M244-routed carries (M244 close
  D1, Fate-3): **(1) standing demo-stack test debt — 8 fails** (6 `test_cockpit` academy/overlay + `test_public_host`
  port-13001 + `test_purge` docker-integration), pre-M244 inherited, 0 real defects, host-dependent, ridden ≥5 v2.6
  milestones — the release close is the final expiry (M244 is proof-only, structurally can't fix them). **(2)
  `DEF-M239-01`** — make the demo build **fail loudly on ENOSPC** (a build-path change validatable only by inducing a
  real ENOSPC). → `rebaseline-standing-failures.md`.
- **✅ RESOLVED (M244):** **reap-17700** (M239-D13) LANDED (iter-10); **DEF-M240-01** (real-video exhibit) dispositioned
  to voice player-presence-only (iter-07, the pre-approved else-branch); **DEF-M226-01** serve-reap TESTED (7→0, iter-11).
- **✅ RESOLVED (M243):** **`DEF-M235-03`** / M204 assign-WRITE — LANDED as `pt-assignment-assign` (Fate-1).
- **→ next `stack-seeding` build-iter:** **`DEF-M215-03(a)` / `F11`** — seed hero identity-key vs generated display-name
  mismatch (cosmetic). Enumerated by id so it stays findable.
- **Older, still unscheduled:** DEF-M10-01 (S3 blob bytes — CONSUMED by M240 for the document facet), DEF-M21-01
  (`replayCmd` hermetic test), CAVEAT-1 (clean-box full `/dev-up`), M314b (prod frozen-read hydration), **M205**-residual
  (tier gates + ATS). Playthroughs futures **M206–M207** stay in vision.

_Last updated: 2026-07-23 — v2.6 "sound check" ALL 8 MILESTONES CLOSED (M237→M244); ready for /developer-kit:close-release._
