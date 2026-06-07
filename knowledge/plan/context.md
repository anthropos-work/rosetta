# Plan — Context

This directory holds the **active** planning artifacts for **Project Rosetta**. It was bootstrapped
on 2026-06-02 to put rosetta on the developer-kit planning lifecycle. **`state.md` is the live source of
truth** — this file is the stable orientation/conventions doc; when the two disagree, `state.md` wins.

**Status (2026-06-07):** **v1.0 "body double"** + **v1.1 "show floor"** + **v1.2 "set dressing"** SHIPPED (tagged
`v1.0` / `v1.1` / `v1.2`; records archived under [`releases/archive/`](releases/archive/)). **v1.2 "set dressing"**
delivered the **dedicated `stack-snapshot` extension**: M9a (framework + capture-safety + tenant firewall +
`.agentspace` store + `/db-query` port + fidelity-DNA) → M9b (taxonomy surface) → M10 (public Directus content) →
M11 (recipes + `/demo-snapshot` + corpus), all section-shaped, lifting M7c's `waived` taxonomy + content to **100%
data-DNA coverage** (capturing **public** data only, from a **safe, low-impact source** — default a prod `pg_dump`,
fallback a throttled read-only pull). **v1.3 "stack party" is now IN DEVELOPMENT** on `release/01.30-stack-party` (designed 2026-06-07) — the **dev/demo convergence**: M12 (unified first-available-N stack registry) → M13 (dev stacks as first-class peers: local Directus + auto-snapshot + light seed) → M14 (generic `stack-*` skills, hard-renamed) → M15 (safety & security doc + dual-repo KB). Former v1.3 seeds (cloud store / S3 blobs, AI content, shareability) → v1.4. Live
state: [`state.md`](state.md). (No `roadmap-legacy.md` yet — that appears when a whole *major* version retires; v1.2
is the same major.)

## Files

- [`roadmap.md`](roadmap.md) — active version (its milestones, execution graph, risks) + shipped-version archive
- [`roadmap-vision.md`](roadmap-vision.md) — future versions + proposals not yet in active development
- [`state.md`](state.md) — current/next milestone, last update
- `releases/{VV.VV}-{codename}/m{N}-{slug}/overview.md` (active version) → `releases/archive/{VV.VV}-{codename}/…` (shipped). _**v1.0's six milestone dirs (M0/M1/M1b/M2/M2b/M2c) are closed and archived** under [`releases/archive/01.00-body-double/`](releases/archive/01.00-body-double/), each with overview/progress/decisions/retro/metrics. The next active version's dirs appear under `releases/{VV.VV}-{codename}/` when `/developer-kit:design-roadmap` scaffolds it._

## Conventions

- One directory per milestone, named `m{N}-{slug}/`
- Each milestone dir has at minimum an `overview.md`. As the milestone progresses, optional companion files: `progress.md` (deliverable checklist), `decisions.md` (implementation choices with rationale), `spec-notes.md` (technical details).
- Status values: `planned` → `in-progress` → `done` → `archived` (terminal, set at release close).
- Milestone numbering (this is rosetta's first version): **flat sequential** — M0, M1, M2, … . A letter suffix has two uses: (1) a milestone **inserted after** the fact (M1b drift CI, M2b consolidation, M2c the iterative `@clerk/express` feature); and (2) a **split** of one planned milestone into a sequential mini-arc (**M7a → M7b → M7c** = the former M7 "seeding" split into framework+safety / data-DNA / fleet, 2026-06-04). Context disambiguates which.
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

**Active:** **v1.3 "stack party"** — in development on `release/01.30-stack-party` (designed 2026-06-07). The
**dev/demo convergence**: M12 (unified stack registry + first-available-N) → M13 (dev peers: local Directus +
auto-snapshot + light seed) → M14 (generic `stack-*` skills + `dev-up`/`dev-down`, hard-renamed) → M15 (safety &
security doc + dual-repo KB), strictly sequential. v1.0 / v1.1 / v1.2 shipped (tags `v1.0` / `v1.1` / `v1.2`); the
former v1.3 seeds → v1.4. _(Live milestone/branch state: [`state.md`](state.md). v1.4 proposals:
[`roadmap-vision.md`](roadmap-vision.md).)_

## Project note

Rosetta is the **documentation corpus** for the Anthropos platform (architecture guides + the
`/dev-up` / `/dev-down` / `/stack-update` skills that build, run, and sync the local *dev*
environment — converged from the former setup/start/update skills in v1.3/M14). The planning lifecycle tracked here governs **extensions to rosetta itself** — the
first being a second corpus + skill set for building disposable, fully-seeded **demo** environments.
It does **not** track changes to the Anthropos platform repos (those live under `anthropos-work`).
