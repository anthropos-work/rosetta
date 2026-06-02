# Plan — Context

This directory holds the **active** planning artifacts for **Project Rosetta**. It was bootstrapped
on 2026-06-02 to put rosetta on the developer-kit planning lifecycle; no version has shipped under
this lifecycle yet, so there is no `plan-legacy/` archive.

## Files

- [`roadmap.md`](roadmap.md) — active version (its milestones, execution graph, risks) + shipped-version archive
- [`roadmap-vision.md`](roadmap-vision.md) — future versions + proposals not yet in active development
- [`state.md`](state.md) — current/next milestone, last update
- `{version}/overview.md` + `{version}/m{N}-{slug}/overview.md` — full design + per-milestone scope for the **active** version (present only while a version is in development). _**v1.0 "body double" (Clerkenstein) was designed 2026-06-02** but its milestone dirs are **not yet scaffolded** and its `release/01.00-body-double` branch is **not yet cut** (Phase 8 deferred by user choice). They appear here once scaffolded or once `/developer-kit:build-milestone` runs._

## Conventions

- One directory per milestone, named `m{N}-{slug}/`
- Each milestone dir has at minimum an `overview.md`. As the milestone progresses, optional companion files: `progress.md` (deliverable checklist), `decisions.md` (implementation choices with rationale), `spec-notes.md` (technical details).
- Status values: `planned` → `in-progress` → `complete`
- Milestone numbering (this is rosetta's first version): **flat sequential** — M1, M2, … ; B-milestones append `b` (e.g. M2b).
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

**Active:** **v1.0 "body double"** (Clerkenstein) — designed 2026-06-02, **branch not yet cut**. The
`release/01.00-body-double` branch + milestone dirs are created when v1.0 is scaffolded (Phase 8,
deferred) or on the first `/developer-kit:build-milestone`. v1.1 "show floor" (M3–M5) is next, staged
in `roadmap-vision.md`.

## Project note

Rosetta is the **documentation corpus** for the Anthropos platform (architecture guides + the
`/setup-platform` / `/start-platform` / `/update-platform` skills that build the local *dev*
environment). The planning lifecycle tracked here governs **extensions to rosetta itself** — the
first being a second corpus + skill set for building disposable, fully-seeded **demo** environments.
It does **not** track changes to the Anthropos platform repos (those live under `anthropos-work`).
