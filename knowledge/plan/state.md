---
active_release: "(between releases) — v2.0 opening night SHIPPED (tag v2.0); awaiting /developer-kit:design-roadmap"
active_branch: "main"
active_milestone: "(between releases)"
last_closed: "v2.0 opening night — 2026-07-02 (tag v2.0, 4 milestones M201..M204) — the Playthroughs pillar"
phase: "between releases — awaiting /developer-kit:design-roadmap"
last_updated: "2026-07-02"
---

# State

**Active release:** **(between releases).** **v2.0 "opening night"** — the **Playthroughs** pillar — SHIPPED
2026-07-02 (tag `v2.0`). No release is active; the next release awaits `/developer-kit:design-roadmap`.

**Last shipped:** **v2.0 "opening night" — 2026-07-02, tag `v2.0`** (rext code-of-record rolled to `v2.0`). The
**Playthroughs** pillar: a Playthrough is an automated actor that IS the user — logs in as a seeded hero, plays a
real journey end-to-end, and proves the platform delivered the outcome (**function**, vs the M42 coverage
protocol's **presence**). **4 milestones M201 ∥ M202 → { M203 ∥ M204 }**, all merged `--no-ff` →
`release/02.00-opening-night` → `main`. Corpus: **10 live Playthroughs (6 employee + 4 manager) GREEN on cold
reset-to-seed, 1 declared in-manifest TODO** (the assign-WRITE half → Fate-2). Tooling + docs only — zero platform
edits, zero new deps. Records: [`releases/archive/02.00-opening-night/`](releases/archive/02.00-opening-night/).

**Active milestone:** **(between releases).** All four v2.0 milestones are closed + merged. No milestone is active.

**Phase:** **between releases — awaiting `/developer-kit:design-roadmap`.** v2.0 shipped clean (close-release: all 9
review sweeps GREEN, no blockers; one M201 stale-label reconcile + 5 small land-now fixes, all Fate-1).

**Next up:** **run `/developer-kit:design-roadmap`** to open the next release. Candidates in `roadmap-vision.md`:
**M205** Hiring + tier gates · **M206** AI-sim mirror tier + the M203 carried edge UCs · **M207** Academy coverage.

**Push-gated KEEP (the user's manual step):** origin has NOT received `main` + tags `v1.10.1` + `v2.0` + the rext
tags (`fit-up-m47..m52`, `v1.10.1`, `opening-night-m201..m204`, and the new **`v2.0`** rext roll). Local closes
deliberately do not push; this is the user's gate. The box-level re-pin (`.agentspace/rext.tag`) stays at
`v1.10.1` (demo-1's pin; the playthroughs section ran from the authoring copy — a coverage release needs no
consumption re-pin). An administrative KEEP, not a deferral.

**Standing backlog (unscheduled, cross-release):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01
(`replayCmd` hermetic test), M25-D9 (dev taxonomy `rc=4`), **M314b** (prod frozen-read whole-org hydration — a
prod-team follow-up, not tooling). All tracked in [`roadmap-vision.md`](roadmap-vision.md); none scheduled.

## Recently shipped releases
- **v2.0 "opening night"** — **2026-07-02**, tag `v2.0`. The **Playthroughs** pillar: manifest corpus (M201) →
  foundation (M202) → employee (M203) + manager (M204) coverage. **10 live Playthroughs GREEN on cold reset-to-seed**
  + 1 in-manifest TODO. 4 milestones M201..M204. **The first v2.x release.** Records:
  [`releases/archive/02.00-opening-night/`](releases/archive/02.00-opening-night/).
- **v1.10b "fit-up"** — **2026-07-01**, tag `v1.10.1`. Interposed field-hardening backfill: re-sync + recapture,
  corpus re-ground, from-cold `/demo-up` hardening, content/AI-readiness-org seeding fill, one auditable seed+gen
  manifest, then a from-cold destroy-and-rebuild acceptance. 7 milestones M47..M53. Records:
  [`releases/archive/01.10b-fit-up/`](releases/archive/01.10b-fit-up/).
- **v1.10 "method acting"** — **2026-06-27**, tag `v1.10`. The believable-profile release + presenter-grade /
  scalable-generation extension; Playwright SEMANTIC coverage gate at both vantages cold; 9 milestones (M39→M46).
  The **last v1.x release** — detail in [`roadmap-legacy.md`](roadmap-legacy.md). Records:
  [`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/).

_(Earlier v1.x — v1.0 … v1.9 — full shipped table in [`roadmap-legacy.md`](roadmap-legacy.md) § Shipped releases.)_

## Headline numbers (v2.0 — final)
- **playthroughs Go test funcs:** **105** (101 Test + 4 Fuzz across manifest/report/ptvalidate/ptreport). **rext
  total 1745** across 6 modules (playthroughs is the new 6th). `go vet ./...` clean; 3/3 + 5/5 shuffled clean.
- **TS:** playthroughs **58** unit specs (url-shapes 46 + stack-env 12) + **10 browser Playthroughs** (6 employee +
  4 manager) — 5/5 cold reset-to-seed deterministic each vantage.
- **Flake:** **0** (Go 3/3 & 5/5; TS 3/3 & 5/5; browser 5/5 cold-reset).
- **Coverage:** 94.8–100% stmt on the new module (invariant-pinning at harden, not %-gains).
- **Supply-chain:** **0 net-new deps** (`ai v1.40.1` unchanged, confined to stack-seeding). The v1.10b-inherited
  HIGH `x/net` CVE-2026-39821 **resolved** (v0.55.0 in stack-seeding). 0 GPL/AGPL. Lockfile:
  [`releases/archive/02.00-opening-night/dependencies.lock`](releases/archive/02.00-opening-night/dependencies.lock).
- **Alignment gates:** **100%/100%** on all 5 Clerkenstein surfaces (carried from v1.10; v2.0 touched no contract surface).

## Branch model / shipped tags
**v2.0 SHIPPED:** `release/02.00-opening-night` cut from `main` 2026-06-28 → merged `--no-ff` → `main` + tagged
`v2.0` at close (LOCAL — origin push is the user's step). M201/M202/M203/M204 all merged `--no-ff` → the release
branch. rext authoring rolled to **`v2.0`** (from `opening-night-m204` @ `c81c6dd` + the close prune); tagged
`opening-night-m201..m204` per milestone. Consumption pin (`.agentspace/rext.tag`) stays `v1.10.1`.
**Shipped tags:** **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` ·
**v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` ·
**v1.1** `v1.1` · **v1.0** `v1.0`. (Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-07-02 (v2.0 "opening night" CLOSED — the Playthroughs pillar shipped, tag `v2.0`, 4 milestones
M201..M204; 10 live Playthroughs GREEN on cold reset-to-seed + 1 in-manifest TODO. Next: `/developer-kit:design-roadmap`.)_
