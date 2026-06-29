# Plan — Context

This directory holds the **active** planning artifacts for **Project Rosetta**. It was bootstrapped
on 2026-06-02 to put rosetta on the developer-kit planning lifecycle. **`state.md` is the live source of
truth** — this file is the stable orientation/conventions doc; when the two disagree, `state.md` wins.

**Status (2026-06-28):** **v1.0 … v1.10 SHIPPED** (the whole **v1.x major**; tagged `v1.0` … `v1.10`; records
archived under [`releases/archive/`](releases/archive/), full history in
[`roadmap-legacy.md`](roadmap-legacy.md)), and **v2.0 "opening night" is IN DEVELOPMENT** — a **NEW MAJOR**
(designed 2026-06-28 via `/developer-kit:design-roadmap`; branch `release/02.00-opening-night`). v2.0 opens the
**Playthroughs** pillar: functional-flow *testing* — a manifest-driven deterministic e2e suite that *pretends to
be the human* and proves the platform's core user journeys **actually work** end-to-end (the **functional**
sibling of v1.x's M42 **presence** coverage sweep). 4 milestones **M201 ∥ M202 → { M203 ∥ M204 }** (`Mxyy`
numbering; M201 `iterative`+user-guided manifest corpus ∥ M202 `section` foundation → M203/M204 `iterative`
per-vantage coverage), governed by the consolidated capability
spec [`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md) v0.3. **Tooling + docs only — zero
platform-repo edits.** The last v1.x release, **v1.10 "method acting"** (SHIPPED 2026-06-27, tag `v1.10`), was the
believable-profile release + the presenter-grade/scalable-generation extension (M39→M46). Genuinely-deferred work
stays **unscheduled backlog** — DEF-M10-01 (cloud store / S3 blob bytes), DEF-M21-01 (replayCmd hermetic test),
M25-D9 (dev taxonomy rc=4) ([`roadmap-vision.md`](roadmap-vision.md)). Live state: [`state.md`](state.md).
(**[`roadmap-legacy.md`](roadmap-legacy.md) now exists** — the v1.x major retired at the v2.0 opening, so its
roadmap rotated there; the active [`roadmap.md`](roadmap.md) holds v2.0 only.)

## Files

- [`roadmap.md`](roadmap.md) — the **active major** (its milestones, execution graph, risks)
- [`roadmap-legacy.md`](roadmap-legacy.md) — the **retired v1.x** roadmap (M0 … M46, all SHIPPED) — created at
  the v2.0 opening when the v1.x major retired
- [`roadmap-vision.md`](roadmap-vision.md) — future versions + future v2 milestones + proposals not yet in active development
- [`state.md`](state.md) — current/next milestone, last update
- `releases/{VV.VV}-{codename}/m{N}-{slug}/overview.md` (active version) → `releases/archive/{VV.VV}-{codename}/…` (shipped). _The **active v2.0 dirs** are under [`releases/02.00-opening-night/`](releases/02.00-opening-night/) (`m201-foundation/`, `m202-employee-coverage/`, `m203-manager-coverage/`), scaffolded by the 2026-06-28 `/developer-kit:design-roadmap` run. v1.x's milestone dirs are closed + archived under `releases/archive/01.{00..10}-{codename}/`, each with overview/progress/decisions/retro/metrics._

## Conventions

- One directory per milestone, named `m{N}-{slug}/`
- Each milestone dir has at minimum an `overview.md`. As the milestone progresses, optional companion files: `progress.md` (deliverable checklist), `decisions.md` (implementation choices with rationale), `spec-notes.md` (technical details).
- Status values: `planned` → `in-progress` → `done` → `archived` (terminal, set at release close).
- Milestone numbering — **v1.x (the first major) = flat sequential** (M0 … M46, closed + archived; detail below). **v2+ = the `Mxyy` scheme** (`M` + major digit + two-digit milestone): **v2.0 "opening night" = M201 ∥ M202 → { M203 ∥ M204 }** (M201 `iterative`+user-guided manifest corpus ∥ M202 `section` Playthroughs-foundation → M203/M204 `iterative` per-vantage coverage — the new-major scheme `context.md` had reserved for "a future *major* v2+"). The v1.x flat-counter detail follows: M0, M1, …, M20, v1.5 = **M21→M25**; **M26** = self-contained-demo, now **v1.8 "understudy"** (re-implemented onto current `main` from the orphaned `m26/self-contained-demo` ext branch @ `25ab855` / tag `prop-room-m26`, authored 2026-06-14 — that orphan is the spec, NOT merged: it predates v1.6/v1.7 which rewrote the same files); v1.6 "stage door" = **M27→M30**; v1.7 "house lights" = **M31→M32** (M33 ant-academy liveness → backlog); v1.8 "understudy" = **M26** (the reserved slot); v1.9 "storytelling" = **M34→M38** (the seeding/Stories-&-Heroes release — the counter resumes at M34, the next free number after the M33 backlog slot); v1.10 "method acting" = **M39→M46** (the believable-profile release **M39→M42m** [the counter resumes at M39; **M42e/M42m** are an `e`/`m` **persona-pair split** [employee/manager] of one planned coverage milestone — the second split-suffix use after M7a/M7b/M7c], **extended 2026-06-26 with M43→M46** — the presenter-grade / scalable-generation extension: M43 cockpit-UX + M44 profile-completeness [`section`] → M45 generation-engine + M46 org-scale-fill [`iterative`]) (the milestone counter never resets — M26 was reserved for the self-contained-demo effort at v1.6 design, so stage-door begins at M27 even though M26 ships LATER, in v1.8; the version *number* jumps 1.3b→1.5 after the v1.4 removal; there is no `M5xx` scheme — that two-digit `Mxyy` scheme only begins at a future *major* v2+). A letter suffix has two uses: (1) a milestone **inserted after** the fact (M1b drift CI, M2b consolidation, M2c the iterative `@clerk/express` feature); and (2) a **split** of one planned milestone into a sequential mini-arc (**M7a → M7b → M7c** = the former M7 "seeding" split into framework+safety / data-DNA / fleet, 2026-06-04). Context disambiguates which.
- Milestone **shapes** can be mixed within a version: `section` (fixed checklist) or `iterative` (measurable exit gate, uncertain path). v1.0 has both — **M0/M1b/M2/M2b are section; M1 and M2c are iterative** (alignment-score gates).
- Date format throughout: ISO `YYYY-MM-DD`
- **Stack workspaces & extension tooling (v1.2):** each gitignored `stack-*/` dir spans one full local stack — its platform service repos **plus** its own clone of rosetta-extensions. The scratchpad rename convention: `anthropos-dev/` → `stack-dev/` (dev), `anthropos-demo/` → `stack-demo/` (demo), `anthropos-dev-2/` → `stack-dev-2/` (secondary dev), with future `stack-stage/` and `stack-tests/`. rosetta-extensions has **two clone roles**: (a) an **authoring** copy at `.agentspace/rosetta-extensions/` — spawned on demand to read/build/**test** tooling, then committed + **tagged**; and (b) **per-stack consumption** copies `stack-<role>/rosetta-extensions @ <tag>` — each stack consumes the tooling at a pinned tag. **Policy:** v1.2 extension code is built+tested in the authoring copy, tagged, then consumed per-stack — never scattered in the rosetta corpus, never authored ad-hoc inside a stack dir. rosetta = read-only doc corpus + dev-env skills; rosetta-extensions = the executable stack tooling.

## Workflow

The standard milestone lifecycle uses the developer-kit skills:

1. `/developer-kit:design-roadmap` — design the version + create branch + scaffold milestone dirs
2. `/developer-kit:build-milestone` — work on a milestone (creates `m{N}/<slug>` branch from the release branch, accumulates commits)
3. `/developer-kit:harden-milestone` — review + close gaps before merging
4. `/developer-kit:close-milestone` — merge milestone → release branch
5. `/developer-kit:close-release` — merge release branch → main, tag

The canonical flow: the `release/{VV.VV}-{codename}` branch is created at design time (the
`/developer-kit:design-roadmap` invocation) so milestone branches have a parent from M1 onward.

**Active:** **v2.0 "opening night" — IN DEVELOPMENT** (designed 2026-06-28 via `/developer-kit:design-roadmap`;
branch `release/02.00-opening-night` cut from `main`). A **NEW MAJOR** opening the **Playthroughs** pillar:
functional-flow *testing* — a manifest-driven deterministic e2e suite that *pretends to be the human* and proves
the platform's core user journeys **actually work** end-to-end (the **functional** sibling of v1.x's M42
**presence** coverage sweep). 4 milestones **M201 ∥ M202 → { M203 ∥ M204 }** (`Mxyy` numbering): **M201** Manifest
corpus [`iterative`, **user-guided**] — top-down user-directed curation of the goal-aligned Product → Story → Use
Case manifest YAML corpus (the build+regression contract, one file per product; prose-only so authorable in
parallel; signed off by the user as the complete-enough v2.0 coverage contract) ∥ **M202** Playthroughs Foundation
[`section`] — the `playthroughs` rext section on the shared M42 e2e foundation (manifest model + light validator +
per-surface page-object layer + dedicated decoupled seed + reset-to-seed serial-runner + the 4-state reporting map
+ one trivial proof Playthrough; builds the validator + seed to match the M201 corpus) → **M203** employee-vantage
coverage ∥ **M204** manager-vantage coverage [both `iterative`, per-vantage exit gates: every declared use case
green on a COLD reset-to-seed demo stack, 0 false-fails over 5 reset runs]. Tooling + docs only — zero
platform-repo edits (an un-drivable surface *escalates* via `unimplementable-without-platform-edit`, it never
edits the platform). Governed by the
consolidated capability spec [`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md) v0.3; milestone
records under [`releases/02.00-opening-night/`](releases/02.00-opening-night/). **Last shipped:** **v1.10 "method
acting"** (2026-06-27, tag `v1.10`; the last v1.x release; archived under
[`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/); full v1.x roadmap in
[`roadmap-legacy.md`](roadmap-legacy.md)). **Next:** **`/developer-kit:work-mstone-iters`** — work M201 (the
`iterative`, user-guided manifest corpus) WITH THE USER; ∥ `/developer-kit:build-milestone` M202 (the `section`
foundation; it gates M203 ∥ M204). _(Live state: [`state.md`](state.md). Backlog: [`roadmap-vision.md`](roadmap-vision.md).)_

## Project note

Rosetta is the **documentation corpus** for the Anthropos platform (architecture guides + the
`/dev-up` / `/dev-down` / `/stack-update` skills that build, run, and sync the local *dev*
environment — converged from the former setup/start/update skills in v1.3/M14). The planning lifecycle tracked here governs **extensions to rosetta itself** — the
first being a second corpus + skill set for building disposable, fully-seeded **demo** environments.
It does **not** track changes to the Anthropos platform repos (those live under `anthropos-work`).
