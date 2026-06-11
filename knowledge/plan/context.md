# Plan — Context

This directory holds the **active** planning artifacts for **Project Rosetta**. It was bootstrapped
on 2026-06-02 to put rosetta on the developer-kit planning lifecycle. **`state.md` is the live source of
truth** — this file is the stable orientation/conventions doc; when the two disagree, `state.md` wins.

**Status (2026-06-11):** **v1.0 … v1.3b SHIPPED** (tagged `v1.0` / `v1.1` / `v1.2` / `v1.3` / `v1.3.1`; records archived
under [`releases/archive/`](releases/archive/)). **v1.5 "prop room" is now IN DEVELOPMENT** (designed 2026-06-11; branch
`release/01.50-prop-room`) — the **local-Directus release**: every stack today reads public content **live from prod**;
v1.5 stands up a **local Directus per stack** serving the **captured public library** (real images preserved via prod
public asset links), demo-default + dev-opt-in. 5 milestones M21→M25 (structure capture [iterative] → executed
provisioning + lifecycle → content cutover + referential closure → docs + hygiene → field bake). **Tooling + docs only —
zero platform-repo edits; capture stays read-only / public-only / prod-untouched.** It closes the gap the 2026-06-11
correction named (v1.3's "local Directus" was print-only). Genuinely-deferred work stays **unscheduled backlog** —
chiefly DEF-M10-01 (cloud store / S3 blob bytes), re-signed at v1.5 design with its sting removed (real images via prod
links); the ex-v1.4 seeds (AI content / shareability / more mirrors) + the deploy-CI gate + the dev-up pre-warm question
were **dropped from tracking** ([`roadmap-vision.md`](roadmap-vision.md)). Live state: [`state.md`](state.md). (No
`roadmap-legacy.md` yet — that appears when a whole *major* version retires; v1.3…v1.5 are the same major.)

## Files

- [`roadmap.md`](roadmap.md) — active version (its milestones, execution graph, risks) + shipped-version archive
- [`roadmap-vision.md`](roadmap-vision.md) — future versions + proposals not yet in active development
- [`state.md`](state.md) — current/next milestone, last update
- `releases/{VV.VV}-{codename}/m{N}-{slug}/overview.md` (active version) → `releases/archive/{VV.VV}-{codename}/…` (shipped). _**v1.0's six milestone dirs (M0/M1/M1b/M2/M2b/M2c) are closed and archived** under [`releases/archive/01.00-body-double/`](releases/archive/01.00-body-double/), each with overview/progress/decisions/retro/metrics. The next active version's dirs appear under `releases/{VV.VV}-{codename}/` when `/developer-kit:design-roadmap` scaffolds it._

## Conventions

- One directory per milestone, named `m{N}-{slug}/`
- Each milestone dir has at minimum an `overview.md`. As the milestone progresses, optional companion files: `progress.md` (deliverable checklist), `decisions.md` (implementation choices with rationale), `spec-notes.md` (technical details).
- Status values: `planned` → `in-progress` → `done` → `archived` (terminal, set at release close).
- Milestone numbering (this is rosetta's first major version — v1.x): **flat sequential** — M0, M1, …, M20, and v1.5 continues the same flat run as **M21→M25** (the version *number* jumps 1.3b→1.5 after the v1.4 removal, but the milestone counter does **not** reset — there is no `M5xx` scheme; that two-digit `Mxyy` scheme only begins at a future *major* v2+). A letter suffix has two uses: (1) a milestone **inserted after** the fact (M1b drift CI, M2b consolidation, M2c the iterative `@clerk/express` feature); and (2) a **split** of one planned milestone into a sequential mini-arc (**M7a → M7b → M7c** = the former M7 "seeding" split into framework+safety / data-DNA / fleet, 2026-06-04). Context disambiguates which.
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

**Active:** **v1.5 "prop room" IN DEVELOPMENT** (designed 2026-06-11; branch `release/01.50-prop-room` cut from `main`).
The **local-Directus release** — 5 strictly-sequential milestones M21→M25: **M21** structure capture (`iterative` —
close the M10 collection-schema gap so the `directus` replay exits 0 and a booted Directus serves the captured catalog)
→ **M22** executed provisioning + per-stack Directus lifecycle (compose-service, idempotent, verified, torn-down) →
**M23** content cutover + referential closure (re-point `DIRECTUS_BASE_ADDR` local; close the taxonomy↔content dangling
gap) → **M24** docs convergence + the 4-item hygiene strand → **M25** field bake (observable-behavior gate on the 16 GB
box). Build M21 with **`/developer-kit:build-mstone-iters`** (iterative), M22→M25 with `/developer-kit:build-milestone`
(section). v1.3b "dress rehearsal" SHIPPED 2026-06-09 (tag `v1.3.1`) was the prior release. _(Live state:
[`state.md`](state.md). Backlog / dropped items: [`roadmap-vision.md`](roadmap-vision.md).)_

## Project note

Rosetta is the **documentation corpus** for the Anthropos platform (architecture guides + the
`/dev-up` / `/dev-down` / `/stack-update` skills that build, run, and sync the local *dev*
environment — converged from the former setup/start/update skills in v1.3/M14). The planning lifecycle tracked here governs **extensions to rosetta itself** — the
first being a second corpus + skill set for building disposable, fully-seeded **demo** environments.
It does **not** track changes to the Anthropos platform repos (those live under `anthropos-work`).
