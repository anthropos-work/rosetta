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
numbering. **Tooling + docs only — zero platform-repo edits.** 4 milestones:

```
M201 ──┐                 (manifest corpus — prose, user-guided)
M202 ──┼──→ M203 ─┐
            M204 ─┴──→ ship
```

**Active milestone:** **M201 — Manifest corpus** (`iterative`, **USER-GUIDED**, complexity large, depends on
none — the manifest is prose). Goal: top-down, **user-directed** curation of the **full goal-aligned Product →
Story → Use Case manifest corpus** — the build+regression contract every coverage milestone (M203/M204)
implements against. Per pass: **outline** (products → stories → use cases) → **validate** (against the real
platform surface + the spec's manifest model) → **write the prose-intent manifest YAML** (spec §5.3, one file
per product). **Explicitly NOT bounded by the current minimal demo seed** — it captures what the goal says must
be proven; where a use case needs preconditions the demo lacks, that **feeds the M202 dedicated-seed expansion**
(noted, not resolved here). Worked **conversationally** + via `/developer-kit:work-mstone-iters`. Exit gate: the
corpus is comprehensively outlined, validated, written as prose-intent YAML (one file per product, each use case
carrying goal + actor + flow + intermediate/final expectations, §5.3 validator passes — ids unique + both-way),
**and the USER signs off the corpus as the complete-enough v2.0 coverage contract.** Records:
[`releases/02.00-opening-night/m201-manifest-corpus/`](releases/02.00-opening-night/m201-manifest-corpus/).

**Phase:** **v2.0 in development** — M201 planned, not yet started. (M202 the foundation is its parallel peer —
prose-only M201 carries no code dependency; M202 builds the validator + dedicated seed to match the corpus.)

**Next up:** **work M201 WITH THE USER** — `/developer-kit:work-mstone-iters` (the user directs each top-down
pass; the first pass outlines the readiest product slice, validates it, writes the prose-intent YAML). The
prose manifest corpus is authorable in **parallel** with the M202 `section` foundation, which then builds the
§5.3 validator + the dedicated seed to match it; M203 ∥ M204 implement Playthroughs against the M201-declared
use cases on the M202 foundation. _(The orchestrator pushes `main` + the `v1.10` tag + the ext tags to origin —
the v1.10 LOCAL close did not push; v2.0 design likewise did not push.)_

**Last shipped:** **v1.10 — 2026-06-27** (`method acting`, 9 milestones M39→M46, tag `v1.10`,
`release/01.10-method-acting` merged `--no-ff` → `main`). The **last release of the v1.x major**; its history +
the full shipped log now live in [`roadmap-legacy.md`](roadmap-legacy.md). Records:
[`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/).

**Paused:** _(none)_

**Standing backlog (unscheduled, cross-release):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes),
DEF-M21-01 (`replayCmd` hermetic test), M25-D9 (dev taxonomy `rc=4`). Pre-existing, tracked in
[`roadmap-vision.md`](roadmap-vision.md); none scheduled. **Future v2 milestones** (Playthroughs pillar, NOT
pre-assigned to a minor): M205 Hiring + tier gates · M206 AI-sim mirror tier · M207 Academy coverage — also in
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
totals (re-measured at first milestone close):
- **Go test funcs (rext):** **1551** total (`Test`+`Fuzz`) at the v1.10 close. Per-module: `alignment` 52 ·
  clerkenstein 270 · stack-seeding 706 · stack-snapshot 363 · stack-secrets 160. (A new `playthroughs` rext
  section arrives in M202; its first tests land at M202 build/close.)
- **Python / TS:** the cockpit `cockpit.py` suite 63 + the demopatch suite 43; `stack-injection` 117. The
  `@playwright/test ^1.49.0` coverage harness (M42e) is the e2e foundation M202 reuses (the first non-Go rext
  dev/test dep).
- **Flake:** **0** at v1.10 close.
- **Supply-chain:** the v1.10 close carried **1 new dep** (`github.com/anthropos-work/ai@v1.40.1`, M45). v2.0 has
  added none yet. The rosetta corpus is docs-only (no package manifest). Lockfile inherited from
  [`releases/archive/01.10-method-acting/dependencies.lock`](releases/archive/01.10-method-acting/dependencies.lock).
- **Alignment gates:** **100%/100%** on all 5 Clerkenstein surfaces at v1.10 close — v2.0 touches no contract
  surface so far.

## Branch model
**v2.0 IN DEVELOPMENT:** `release/02.00-opening-night` cut from `main` 2026-06-28 (LOCAL — origin push is the
orchestrator's step). Milestone branches `m{201,202,203,204}/{slug}` will branch from it at build time. rext code
of record (a SEPARATE repo) gains a `playthroughs` section, authored + tagged per the tooling policy at M202 build.
**v1.10 SHIPPED:** `release/01.10-method-acting` merged `--no-ff` → `main` + tagged `v1.10` at close (LOCAL).
rext code @ tags `method-acting-m39..m46-servegrant-closure`.
**Shipped:** **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` ·
**v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.
(Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-06-28 (**v2.0 "opening night" — M201 Manifest corpus INSERTED as the new first milestone**;
the prior 3 milestones renumbered M201→M202 [foundation], M202→M203 [employee], M203→M204 [manager]. 4 milestones
M201 ∥ M202 → { M203 ∥ M204 } [`Mxyy` numbering]: M201 `iterative`+user-guided manifest corpus is the prose
build+regression contract, authorable in parallel with the M202 `section` foundation. Active milestone now M201;
next: work it WITH THE USER via `/developer-kit:work-mstone-iters`. Future v2 milestones bumped to M205
Hiring/tier-gates, M206 AI-sim-mirror-tier, M207 Academy. Prior: 2026-06-28 v2.0 DESIGNED + PROMOTED — a NEW MAJOR
opening the Playthroughs pillar, branch `release/02.00-opening-night` cut from `main`, from
`spec-drafts/playthroughs/spec.md` v0.3. Headline numbers reset to the v1.10-close inheritance baseline.)_
