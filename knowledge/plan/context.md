# Plan — Context

This directory holds the **active** planning artifacts for **Project Rosetta**. It was bootstrapped
on 2026-06-02 to put rosetta on the developer-kit planning lifecycle. **`state.md` is the live source of
truth** — this file is the stable orientation/conventions doc; when the two disagree, `state.md` wins.

**Status (2026-06-15):** **v1.0 … v1.6 SHIPPED** (tagged `v1.0` / `v1.1` / `v1.2` / `v1.3` / `v1.3.1` / `v1.5` / `v1.6`;
records archived under [`releases/archive/`](releases/archive/)). **v1.7 "house lights" is now IN DEVELOPMENT** (designed
2026-06-15; branch `release/01.70-house-lights`) — a **demo-UI-hardening release**: a fresh browser at a demo's offset UI
renders the working app with **zero manual steps**. Triggered by a live defect (next-web at `http://localhost:33000`
blanked because clerk-js's handshake to the fake FAPI hit an untrusted self-signed cert). 2 milestones M31→M32, both
`section`-shaped: **M31** automates a locally-trusted **mkcert** FAPI cert into the demo bring-up (openssl fallback +
`DEMO_NO_MKCERT` opt-out; fake BAPI is plain HTTP → out of scope) → **M32** fixes the sibling studio-desk
`:9100`-dead-redirect (a 1-line `NODE_ENV=production` override fix + the `:9100` doc/CORS sweep). **Tooling + docs only —
zero platform-repo edits.** Prior: v1.6 "stage door" (the secret-provisioning release, M27→M30) shipped 2026-06-14.
Genuinely-deferred work stays **unscheduled backlog** — M33 (ant-academy demo liveness, repro-first), M26 (self-contained
demo, orphaned ext effort), DEF-M10-01 (cloud store / S3 blob bytes), DEF-M21-01 (replayCmd hermetic test), M25-D9 (dev
taxonomy rc=4) ([`roadmap-vision.md`](roadmap-vision.md)). Live state: [`state.md`](state.md). (No `roadmap-legacy.md`
yet — that appears when a whole *major* version retires; v1.3…v1.7 are the same major.)

## Files

- [`roadmap.md`](roadmap.md) — active version (its milestones, execution graph, risks) + shipped-version archive
- [`roadmap-vision.md`](roadmap-vision.md) — future versions + proposals not yet in active development
- [`state.md`](state.md) — current/next milestone, last update
- `releases/{VV.VV}-{codename}/m{N}-{slug}/overview.md` (active version) → `releases/archive/{VV.VV}-{codename}/…` (shipped). _**v1.0's six milestone dirs (M0/M1/M1b/M2/M2b/M2c) are closed and archived** under [`releases/archive/01.00-body-double/`](releases/archive/01.00-body-double/), each with overview/progress/decisions/retro/metrics. The next active version's dirs appear under `releases/{VV.VV}-{codename}/` when `/developer-kit:design-roadmap` scaffolds it._

## Conventions

- One directory per milestone, named `m{N}-{slug}/`
- Each milestone dir has at minimum an `overview.md`. As the milestone progresses, optional companion files: `progress.md` (deliverable checklist), `decisions.md` (implementation choices with rationale), `spec-notes.md` (technical details).
- Status values: `planned` → `in-progress` → `done` → `archived` (terminal, set at release close).
- Milestone numbering (this is rosetta's first major version — v1.x): **flat sequential** — M0, M1, …, M20, v1.5 = **M21→M25**; **M26** = the orphaned `self-contained-demo` ext work (branch `m26/self-contained-demo` @ `25ab855`, tag `prop-room-m26`, made 2026-06-14 — "make demo stacks self-contained"; **untracked, unmerged, pending its own roadmap home** per user decision); v1.6 "stage door" = **M27→M30**; and v1.7 "house lights" = **M31→M32** (M33 ant-academy liveness → backlog) (the milestone counter never resets — M26 was consumed by the self-contained-demo effort, so stage-door begins at M27; the version *number* jumps 1.3b→1.5 after the v1.4 removal; there is no `M5xx` scheme — that two-digit `Mxyy` scheme only begins at a future *major* v2+). A letter suffix has two uses: (1) a milestone **inserted after** the fact (M1b drift CI, M2b consolidation, M2c the iterative `@clerk/express` feature); and (2) a **split** of one planned milestone into a sequential mini-arc (**M7a → M7b → M7c** = the former M7 "seeding" split into framework+safety / data-DNA / fleet, 2026-06-04). Context disambiguates which.
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

**Active:** **v1.7 "house lights" IN DEVELOPMENT** (designed 2026-06-15; branch `release/01.70-house-lights` cut from
`main`). The **demo-UI-hardening release** — 2 strictly-sequential `section` milestones M31→M32: **M31** mkcert-trusted
FAPI cert (automate a locally-trusted `mkcert` cert into the demo bring-up so a fresh browser at next-web renders instead
of blanking; openssl fallback + `DEMO_NO_MKCERT` opt-out; fake BAPI plain HTTP → out of scope) → **M32** studio-desk
single-port/production (a 1-line `NODE_ENV=production` override fix so studio-desk stops 302-ing to the dead `:9100`, +
the `:9100` doc/CORS sweep). Build both with **`/developer-kit:build-milestone`**. ant-academy demo liveness → backlog
(M33, repro-first). v1.6 "stage door" SHIPPED 2026-06-14 (tag `v1.6`) was the prior release. _(Live state:
[`state.md`](state.md). Backlog: [`roadmap-vision.md`](roadmap-vision.md).)_

## Project note

Rosetta is the **documentation corpus** for the Anthropos platform (architecture guides + the
`/dev-up` / `/dev-down` / `/stack-update` skills that build, run, and sync the local *dev*
environment — converged from the former setup/start/update skills in v1.3/M14). The planning lifecycle tracked here governs **extensions to rosetta itself** — the
first being a second corpus + skill set for building disposable, fully-seeded **demo** environments.
It does **not** track changes to the Anthropos platform repos (those live under `anthropos-work`).
