# Plan — Context

This directory holds the **active** planning artifacts for **Project Rosetta**. It was bootstrapped
on 2026-06-02 to put rosetta on the developer-kit planning lifecycle. **`state.md` is the live source of
truth** — this file is the stable orientation/conventions doc; when the two disagree, `state.md` wins.

**Status (2026-06-03):** **v1.0 "body double"** SHIPPED (tagged `v1.0`; records archived under
[`releases/archive/01.00-body-double/`](releases/archive/01.00-body-double/)). **v1.1 "show floor" is now IN
DEVELOPMENT** on `release/01.10-show-floor` — M3 (demo stacks) → M4 (seeding) → M5 (recipes), all section-shaped,
built additively under the gitignored `anthropos-demo/`. Live state: [`state.md`](state.md). (No `roadmap-legacy.md`
yet — that appears when a whole *major* version retires; v1.1 is the same major.)

## Files

- [`roadmap.md`](roadmap.md) — active version (its milestones, execution graph, risks) + shipped-version archive
- [`roadmap-vision.md`](roadmap-vision.md) — future versions + proposals not yet in active development
- [`state.md`](state.md) — current/next milestone, last update
- `releases/{VV.VV}-{codename}/m{N}-{slug}/overview.md` (active version) → `releases/archive/{VV.VV}-{codename}/…` (shipped). _**v1.0's six milestone dirs (M0/M1/M1b/M2/M2b/M2c) are closed and archived** under [`releases/archive/01.00-body-double/`](releases/archive/01.00-body-double/), each with overview/progress/decisions/retro/metrics. The next active version's dirs appear under `releases/{VV.VV}-{codename}/` when `/developer-kit:design-roadmap` scaffolds it._

## Conventions

- One directory per milestone, named `m{N}-{slug}/`
- Each milestone dir has at minimum an `overview.md`. As the milestone progresses, optional companion files: `progress.md` (deliverable checklist), `decisions.md` (implementation choices with rationale), `spec-notes.md` (technical details).
- Status values: `planned` → `in-progress` → `done` → `archived` (terminal, set at release close).
- Milestone numbering (this is rosetta's first version): **flat sequential** — M0, M1, M2, … ; a letter suffix marks a milestone **inserted after** the fact (M1b drift CI, M2b consolidation, M2c the iterative `@clerk/express` feature).
- Milestone **shapes** can be mixed within a version: `section` (fixed checklist) or `iterative` (measurable exit gate, uncertain path). v1.0 has both — **M0/M1b/M2/M2b are section; M1 and M2c are iterative** (alignment-score gates).
- Date format throughout: ISO `YYYY-MM-DD`

## Workflow

The standard milestone lifecycle uses the developer-kit skills:

1. `/developer-kit:design-roadmap` — design the version + create branch + scaffold milestone dirs
2. `/developer-kit:build-milestone` — work on a milestone (creates `m{N}/<slug>` branch from the release branch, accumulates commits)
3. `/developer-kit:harden-milestone` — review + close gaps before merging
4. `/developer-kit:close-milestone` — merge milestone → release branch
5. `/developer-kit:close-release` — merge release branch → main, tag

The canonical flow: the `release/{VV.VV}-{codename}` branch is created at design time (the
`/developer-kit:design-roadmap` invocation) so milestone branches have a parent from M1 onward.

**Active:** **v1.1 "show floor"** — in development on `release/01.10-show-floor` (designed 2026-06-03). M3
(disposable demo stacks) → M4 (declarative seeding) → M5 (corpus + recipes + polish), strictly sequential, all
additive under `anthropos-demo/`. v1.0 "body double" shipped (tag `v1.0`, 2026-06-03). _(Live milestone/branch
state: [`state.md`](state.md). v2+ proposals: [`roadmap-vision.md`](roadmap-vision.md).)_

## Project note

Rosetta is the **documentation corpus** for the Anthropos platform (architecture guides + the
`/setup-platform` / `/start-platform` / `/update-platform` skills that build the local *dev*
environment). The planning lifecycle tracked here governs **extensions to rosetta itself** — the
first being a second corpus + skill set for building disposable, fully-seeded **demo** environments.
It does **not** track changes to the Anthropos platform repos (those live under `anthropos-work`).
