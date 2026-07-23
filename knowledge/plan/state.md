---
active_release: "v2.6 «sound check» — SHIPPED 2026-07-23 (the reliability / field-hardening release: make everything that's built actually get built + provisioned, proven live on billion). 8 milestones M237→M244, all closed; branch release/02.60-sound-check; tag v2.6. release→main merge + v2.6 tag = close-release Phase 11 (orchestrator; not yet done)."
active_branch: "release/02.60-sound-check (awaiting Phase 11 release→main merge + v2.6 tag)"
active_milestone: "(between releases)"
last_closed: "M244 — 2026-07-23 (prove-on-billion, iterative closed-on-gate 8/8); v2.6 release closed 2026-07-23"
phase: "between releases — awaiting /developer-kit:design-roadmap"
last_updated: "2026-07-23"
---

# State

**Between releases.** v2.6 "sound check" **SHIPPED 2026-07-23** — the reliability / field-hardening release
(the v1.3b / v1.10b / v2.1 / v2.3 lineage): *make everything that's built actually get built + provisioned,
proven live on `billion`.* 8 milestones **M237 → M244**, tooling + docs only, **0 platform-repo edits**. No
milestone is active; the next step is **`/developer-kit:design-roadmap`** to design **v2.7** (v2.7 is NOT yet
designed — do not treat it as active).

> ⚠️ **For the orchestrator:** the `release/02.60-sound-check → main` merge + the **`v2.6`** tag are
> **close-release Phase 11** and **not yet done** (this paperwork is Phases 5–10). Also still open from v2.5:
> v2.5's `release → main` merge + `v2.5` tag are **LOCAL-ONLY**, not pushed to origin (R5) — flag to the user,
> do not auto-push.

## v2.6 headline — proved live on billion

M244's multi-part exit gate (a–h) discharged **8/8 GREEN** on `billion`, cold reset-to-seed, driving from a
tailnet peer: ORG-CLEAN 0 tokens · content-stories **47/47** of the 49-pair denominator (2 voice player cells
presence-only) · the **40 live-browser specs** executed green (24 stack-verify + **16/16 Playthroughs**, 96
cases) · anon academy twin · serve-reap 7→0 · 3 v2.3 drift-carries · interview alignment · all 6 v2.6 fixes +
**p95 1.46 s / 1.31 s**. This discharged v2.5's headline debt (its `29/29` shipped unit-proven, never
live-re-proven). Shape: **barrier → parallel fixes → prove-on-billion** — clone-freshness + a 7-state pin
model (M237); academy Start-404 + language (M238); talk-to-data FULL via real AWS Bedrock creds (M239);
content-stories fidelity + voice presence-only (M240); EN/IT language axis (M241); cockpit UX (M242); the
FIRST mutating Playthrough `pt-assignment-assign` (M243). (Detail: [`roadmap.md`](roadmap.md) § Done — v2.6.)

## Recently shipped (releases, newest first — max 3)

- **v2.6 "sound check"** — 2026-07-23 (tag `v2.6`). The **reliability / field-hardening** release: proved the
  whole v2.6 feature + v2.5's headline live on `billion`, cold, **0 platform edits**. 8 milestones M237→M244.
  47/47 content-stories · 16/16 Playthroughs · p95 1.46/1.31 s. rext code-of-record **`498b1a5`** (tag
  `sound-check-m244-content-sweep-robustness`, on origin). The `billion` demo LEFT UP (green at the m244 pin).
- **v2.5 "the playbill"** — 2026-07-20 (tag `v2.5`). The **content-vantage** release: 29/29 landable
  (session × action) pairs + academy grid filled (65 cards), 0 platform edits across 8 milestones. **Its
  29/29 is now LIVE-RE-PROVEN by v2.6/M244** (47/47 of the grown 49-pair denominator).
- **v2.4 "casting call"** — 2026-07-18 (tag `v2.4`). The **recruiter-vantage / hiring-org** release: a 4th
  hiring org on the cockpit (45 candidates on 5 shared positions), proven live on `billion` (recruiter p95
  1.27 s).

## Headline numbers (v2.6 close, 2026-07-23)
- **Go:** **2010** reproducible platform test funcs (v2.5 1976, **+34**). 2461 testcases / 0 failed, 6 modules.
- **TypeScript (unit):** **257** `*.unit.spec.ts` (v2.5 196, **+61**) + the **40 live-browser specs EXECUTED
  GREEN live on `billion`** (24 stack-verify + 16/16 Playthroughs).
- **Python (rext demo-stack):** **839 pass / 8 standing fail** (host-sensitive; **0 real defects**).
- **content-stories:** **47/47** landed live of the 49-pair denominator (2 voice player presence-only).
- **p95 click→ACCESS (cold on `billion`):** employee **1.46 s** · manager **1.31 s** (< 5 s gate).
- **Flake: 0.** **Alignment (Clerkenstein): 100% / 100% critical.** **Supply chain: GREEN — 0 net-new deps.**
- **Platform-repo edits: 0** (all 8 milestones).

## Shipped tags
**v2.6** (branch `release/02.60-sound-check`; merge+tag = Phase 11) · **v2.5** · **v2.4** · **v2.3** · **v2.2** ·
**v2.1** · **v2.0** · **v1.10b** … · **v1.0**. (Detail: [`roadmap.md`](roadmap.md) § Done + [`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog → v2.7 "test-health" (the sole carry out of v2.6)

> **Every item has a FATED destination — a MILESTONE, not "a sweep"/"the next close"/"standing backlog".**

- **→ v2.7 "test-health" (NAMED milestone; user fate 2026-07-23 "tag now, carry to v2.7"):** the **8 standing
  demo-stack test failures** (6 `test_cockpit` academy/overlay + `test_public_host` port-13001 + `test_purge`
  docker-integration — pre-M244 inherited, **0 real defects**, host-dependent) **+** the rext
  `stack-verify/e2e/run-unit.sh` **roster nit** (2 M244-harden-added specs `content-denominator` +
  `run-discrete` unrostered → runner exits 2 + `UnitSpecsAreExecuted` guard RED; all 9 pass when run
  directly). Recorded in [`roadmap-vision.md`](roadmap-vision.md) § v2.6 → v2.7 carry.
- **DROPPED (not carried):** `DEF-M239-01` (ENOSPC loud-build-fail) — the disk-full class is already covered
  by M239's pre-flight disk-measure; un-validatable belt-and-braces.
- **Older, still unscheduled:** `DEF-M215-03(a)`/`F11` (seed hero identity-key vs display-name mismatch,
  cosmetic) · DEF-M10-01 (CONSUMED by M240 for the document facet) · DEF-M21-01 · CAVEAT-1 · M314b · **M205**
  residual · Playthroughs futures **M206–M207** (in vision).

_Last updated: 2026-07-23 — v2.6 "sound check" SHIPPED (M237→M244 all closed); between releases, awaiting /developer-kit:design-roadmap for v2.7._
