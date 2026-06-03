# Plan — Context

This directory holds the **active** planning artifacts for **Project Rosetta**. It was bootstrapped
on 2026-06-02 to put rosetta on the developer-kit planning lifecycle. **`state.md` is the live source of
truth** — this file is the stable orientation/conventions doc; when the two disagree, `state.md` wins.

**Status (2026-06-03):** **v1.0 "body double"** (Clerkenstein) is **feature-complete** — all six milestones
(M0 → M1 → M1b → M2 → M2b → M2c) are closed and merged onto `release/01.00-body-double`, and the release is
in **`/developer-kit:close-release`** (merge → `main` + tag `v1.0`). No major version has fully shipped +
tagged yet, so there is no `roadmap-legacy.md` / `plan-legacy/` archive (it appears at the first tag).

## Files

- [`roadmap.md`](roadmap.md) — active version (its milestones, execution graph, risks) + shipped-version archive
- [`roadmap-vision.md`](roadmap-vision.md) — future versions + proposals not yet in active development
- [`state.md`](state.md) — current/next milestone, last update
- `releases/{VV.VV}-{codename}/m{N}-{slug}/overview.md` — full design + per-milestone scope. _**v1.0's six milestone dirs are scaffolded and closed** under [`releases/01.00-body-double/`](releases/01.00-body-double/) (M0/M1/M1b/M2/M2b/M2c); the `release/01.00-body-double` branch is cut and carries all their merges. At release close these archive to `releases/archive/01.00-body-double/`._

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

**Active:** **v1.0 "body double"** (Clerkenstein) — **feature-complete, in release-close.** All six
milestones (M0→M2c) closed + merged onto `release/01.00-body-double`; `/developer-kit:close-release` merges
it → `main` and tags `v1.0`. **v1.1 "show floor"** (M3–M5: disposable stacks + seeding + recipes) is next,
staged in [`roadmap-vision.md`](roadmap-vision.md) and promoted to active by `/developer-kit:design-roadmap`
after v1.0 tags. _(For the live milestone/branch state, read [`state.md`](state.md).)_

## Project note

Rosetta is the **documentation corpus** for the Anthropos platform (architecture guides + the
`/setup-platform` / `/start-platform` / `/update-platform` skills that build the local *dev*
environment). The planning lifecycle tracked here governs **extensions to rosetta itself** — the
first being a second corpus + skill set for building disposable, fully-seeded **demo** environments.
It does **not** track changes to the Anthropos platform repos (those live under `anthropos-work`).
