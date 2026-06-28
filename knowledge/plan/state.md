---
active_release: "v2.0 opening night"
active_branch: "release/02.00-opening-night"
active_milestone: "M201"
phase: "v2.0 in development"
last_updated: "2026-06-28"
---

# State

**Active release:** **v2.0 "opening night" — IN DEVELOPMENT** (designed 2026-06-28 via
`/developer-kit:design-roadmap`; branch `release/02.00-opening-night` cut from `main`). A **NEW MAJOR** — opens
the **Playthroughs** pillar: functional-flow *testing*, a manifest-driven deterministic e2e suite that *pretends
to be the human* and proves the platform's core user journeys **actually work** end-to-end (the **functional**
sibling of v1.x's M42 **presence** coverage sweep). Governed by the consolidated capability spec
[`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md) (v0.3). v2+ uses **`Mxyy`** milestone
numbering. **Tooling + docs only — zero platform-repo edits.** 3 milestones:

```
M201 ──→ M202 ─┐
          M203 ─┴──→ ship
```

**Active milestone:** **M201 — Playthroughs Foundation** (`section`, complexity medium, depends on none —
reuses the M42 harness + seeding machinery). Goal: stand up the `playthroughs` rext section on the shared M42
e2e foundation, proven by **one trivial end-to-end Playthrough** (login → /profile → assert hero identity).
In-scope: the manifest model + light validator (both-way id integrity + precondition-coverage, datadna-gated);
the per-surface locator/landmark page-object layer (1 surface to start); the dedicated decoupled seed preset
(entitlement tiers + multi-org-private); the reset-to-seed lifecycle + serial-default runner; the 4-state
reporting map (passing / failing / unimplemented / unimplementable-without-platform-edit); one trivial proof
Playthrough. Delivers → a corpus runbook graduating the spec (`corpus/ops/demo/playthroughs.md`). Records:
[`releases/02.00-opening-night/m201-foundation/`](releases/02.00-opening-night/m201-foundation/).

**Phase:** **v2.0 in development** — M201 planned, not yet started.

**Next up:** **`/developer-kit:build-milestone`** — build **M201** (the `section` foundation; it gates the two
`iterative` vantage-coverage milestones M202 ∥ M203). _(The orchestrator pushes `main` + the `v1.10` tag + the
ext tags to origin — the v1.10 LOCAL close did not push; v2.0 design likewise did not push.)_

**Last shipped:** **v1.10 — 2026-06-27** (`method acting`, 9 milestones M39→M46, tag `v1.10`,
`release/01.10-method-acting` merged `--no-ff` → `main`). The **last release of the v1.x major**; its history +
the full shipped log now live in [`roadmap-legacy.md`](roadmap-legacy.md). Records:
[`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/).

**Paused:** _(none)_

**Standing backlog (unscheduled, cross-release):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes),
DEF-M21-01 (`replayCmd` hermetic test), M25-D9 (dev taxonomy `rc=4`). Pre-existing, tracked in
[`roadmap-vision.md`](roadmap-vision.md); none scheduled. **Future v2 milestones** (Playthroughs pillar, NOT
pre-assigned to a minor): M204 Hiring + tier gates · M205 AI-sim mirror tier · Academy coverage — also in
`roadmap-vision.md`.

## Recently shipped releases
- **v1.10 "method acting"** — **2026-06-27**, tag `v1.10`. The **believable-profile release + the presenter-grade
  / scalable-generation extension**: a logged-in hero reads as a fully-fleshed person, proven by a **Playwright
  SEMANTIC coverage gate** at BOTH vantages cold (M42e employee / M42m manager), extended with M43 cockpit UX,
  M44 whole-roster completeness, M45 a cheap-LLM generation engine (first new dep, `ai@v1.40.1`), M46 org-scale
  fill. 9 milestones. Zero platform-repo edits; all 5 Clerkenstein gates 100%/100%. The **last v1.x release** —
  its detail + the full shipped log are in [`roadmap-legacy.md`](roadmap-legacy.md). Records:
  [`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/).
- **v1.9 "storytelling"** — **2026-06-23**, tag `v1.9`. The declarative **Stories & Heroes** seeding engine + a
  presenter cockpit. 5 `section` milestones M34→M38.
  Records: [`releases/archive/01.90-storytelling/`](releases/archive/01.90-storytelling/).
- **v1.8 "understudy"** — **2026-06-15**, tag `v1.8`. Self-contained-demo release (a box with only `stack-demo/`
  runs a demo end-to-end). Single `section` milestone M26.
  Records: [`releases/archive/01.80-understudy/`](releases/archive/01.80-understudy/).
- **Earlier v1.x** (v1.0 … v1.7) — the full shipped table is in
  [`roadmap-legacy.md`](roadmap-legacy.md) § Shipped releases.

## Headline numbers (v2.0 — inheriting the v1.10-close baseline; no v2 work yet)
The v2.0 baseline is the v1.10-close inheritance — no v2.0 milestone has built yet, so these are the carried-over
totals (re-measured at first M201 close):
- **Go test funcs (rext):** **1551** total (`Test`+`Fuzz`) at the v1.10 close. Per-module: `alignment` 52 ·
  clerkenstein 270 · stack-seeding 706 · stack-snapshot 363 · stack-secrets 160. (A new `playthroughs` rext
  section arrives in M201; its first tests land at M201 build/close.)
- **Python / TS:** the cockpit `cockpit.py` suite 63 + the demopatch suite 43; `stack-injection` 117. The
  `@playwright/test ^1.49.0` coverage harness (M42e) is the e2e foundation M201 reuses (the first non-Go rext
  dev/test dep).
- **Flake:** **0** at v1.10 close.
- **Supply-chain:** the v1.10 close carried **1 new dep** (`github.com/anthropos-work/ai@v1.40.1`, M45). v2.0 has
  added none yet. The rosetta corpus is docs-only (no package manifest). Lockfile inherited from
  [`releases/archive/01.10-method-acting/dependencies.lock`](releases/archive/01.10-method-acting/dependencies.lock).
- **Alignment gates:** **100%/100%** on all 5 Clerkenstein surfaces at v1.10 close — v2.0 touches no contract
  surface so far.

## Branch model
**v2.0 IN DEVELOPMENT:** `release/02.00-opening-night` cut from `main` 2026-06-28 (LOCAL — origin push is the
orchestrator's step). Milestone branches `m{201,202,203}/{slug}` will branch from it at build time. rext code of
record (a SEPARATE repo) gains a `playthroughs` section, authored + tagged per the tooling policy at M201 build.
**v1.10 SHIPPED:** `release/01.10-method-acting` merged `--no-ff` → `main` + tagged `v1.10` at close (LOCAL).
rext code @ tags `method-acting-m39..m46-servegrant-closure`.
**Shipped:** **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` ·
**v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.
(Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-06-28 (**v2.0 "opening night" DESIGNED + PROMOTED** via `/developer-kit:design-roadmap` — a
NEW MAJOR opening the Playthroughs pillar; 3 milestones M201 → { M202 ∥ M203 } [`Mxyy` numbering], branch
`release/02.00-opening-night` cut from `main`. v1.x history rotated to `roadmap-legacy.md`; the active roadmap
holds v2.0 only. Designed from `spec-drafts/playthroughs/spec.md` v0.3. Headline numbers reset to the v1.10-close
inheritance baseline. Next: `/developer-kit:build-milestone` — M201.)_
