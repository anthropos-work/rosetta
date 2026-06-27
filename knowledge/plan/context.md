# Plan — Context

This directory holds the **active** planning artifacts for **Project Rosetta**. It was bootstrapped
on 2026-06-02 to put rosetta on the developer-kit planning lifecycle. **`state.md` is the live source of
truth** — this file is the stable orientation/conventions doc; when the two disagree, `state.md` wins.

**Status (2026-06-23):** **v1.0 … v1.9 SHIPPED** (tagged `v1.0` / `v1.1` / `v1.2` / `v1.3` / `v1.3.1` / `v1.5` /
`v1.6` / `v1.7` / `v1.8` / `v1.9`; records archived under [`releases/archive/`](releases/archive/)), and **no
version is currently in development** — the next awaits a `/developer-kit:design-roadmap` run. v1.9 "storytelling"
(SHIPPED 2026-06-23, tag `v1.9`) was the **believable-demo-narrative release** — it converted the placeholder
seeder into a declarative **Stories & Heroes** engine: each *story* is one org with a thriving/struggling/manager
**hero** trio, seeded via the real **verified-skill chain** (the 7-table jobsim→`user_skills`→`user_skill_evidences`
fan-out) so the two product Musts — the individual **skill profile** and the org **Workforce dashboard** — tell one
coherent story, plus a standalone **presenter cockpit** (log in as a hero + jump to the right screen). 5 `section`
milestones **M34→M38**, designed from the adversarially-verified spec [`.agentspace/seeding_gaps.md`]. **Tooling +
docs only — zero platform-repo edits.** Prior: v1.8 "understudy" (the self-contained-demo release, M26) shipped
2026-06-15. Genuinely-deferred work stays **unscheduled backlog** — DEF-M10-01 (cloud store / S3 blob bytes),
DEF-M21-01 (replayCmd hermetic test), M25-D9 (dev taxonomy rc=4) ([`roadmap-vision.md`](roadmap-vision.md)).
A post-v1.9 **demo-hardening pass** shipped at rext tag `storytelling-postfix-1` (tooling + docs only) — it
**resolved M33** (ant-academy/cockpit demo liveness, the session-detach fix) and made DEMO_STORIES the demo
default. Live state: [`state.md`](state.md). (No `roadmap-legacy.md` yet —
that appears when a whole *major* version retires; v1.3…v1.9 are the same major.)

## Files

- [`roadmap.md`](roadmap.md) — active version (its milestones, execution graph, risks) + shipped-version archive
- [`roadmap-vision.md`](roadmap-vision.md) — future versions + proposals not yet in active development
- [`state.md`](state.md) — current/next milestone, last update
- `releases/{VV.VV}-{codename}/m{N}-{slug}/overview.md` (active version) → `releases/archive/{VV.VV}-{codename}/…` (shipped). _**v1.0's six milestone dirs (M0/M1/M1b/M2/M2b/M2c) are closed and archived** under [`releases/archive/01.00-body-double/`](releases/archive/01.00-body-double/), each with overview/progress/decisions/retro/metrics. The next active version's dirs appear under `releases/{VV.VV}-{codename}/` when `/developer-kit:design-roadmap` scaffolds it._

## Conventions

- One directory per milestone, named `m{N}-{slug}/`
- Each milestone dir has at minimum an `overview.md`. As the milestone progresses, optional companion files: `progress.md` (deliverable checklist), `decisions.md` (implementation choices with rationale), `spec-notes.md` (technical details).
- Status values: `planned` → `in-progress` → `done` → `archived` (terminal, set at release close).
- Milestone numbering (this is rosetta's first major version — v1.x): **flat sequential** — M0, M1, …, M20, v1.5 = **M21→M25**; **M26** = self-contained-demo, now **v1.8 "understudy"** (re-implemented onto current `main` from the orphaned `m26/self-contained-demo` ext branch @ `25ab855` / tag `prop-room-m26`, authored 2026-06-14 — that orphan is the spec, NOT merged: it predates v1.6/v1.7 which rewrote the same files); v1.6 "stage door" = **M27→M30**; v1.7 "house lights" = **M31→M32** (M33 ant-academy liveness → backlog); v1.8 "understudy" = **M26** (the reserved slot); v1.9 "storytelling" = **M34→M38** (the seeding/Stories-&-Heroes release — the counter resumes at M34, the next free number after the M33 backlog slot); v1.10 "method acting" = **M39→M46** (the believable-profile release **M39→M42m** [the counter resumes at M39; **M42e/M42m** are an `e`/`m` **persona-pair split** [employee/manager] of one planned coverage milestone — the second split-suffix use after M7a/M7b/M7c], **extended 2026-06-26 with M43→M46** — the presenter-grade / scalable-generation extension: M43 cockpit-UX + M44 profile-completeness [`section`] → M45 generation-engine + M46 org-scale-fill [`iterative`]) (the milestone counter never resets — M26 was reserved for the self-contained-demo effort at v1.6 design, so stage-door begins at M27 even though M26 ships LATER, in v1.8; the version *number* jumps 1.3b→1.5 after the v1.4 removal; there is no `M5xx` scheme — that two-digit `Mxyy` scheme only begins at a future *major* v2+). A letter suffix has two uses: (1) a milestone **inserted after** the fact (M1b drift CI, M2b consolidation, M2c the iterative `@clerk/express` feature); and (2) a **split** of one planned milestone into a sequential mini-arc (**M7a → M7b → M7c** = the former M7 "seeding" split into framework+safety / data-DNA / fleet, 2026-06-04). Context disambiguates which.
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

**Active:** **v1.10 "method acting" — IN DEVELOPMENT (EXTENDED with M43→M46)** (designed 2026-06-24, **extended
2026-06-26**, via `/developer-kit:design-roadmap`; branch `release/01.10-method-acting` cut from `main`). The
**believable-profile release** (M39–M42m, all CLOSED) — make each hero hold up under a close-up when a presenter
*Logs in as* them: profile identity (org name + role + real-face avatar) + the content-surface unblock (library +
activity feed via one per-stack-Directus serve-grant) + profile depth (work/education/deep role-aligned skills) +
**100% per-vantage demo coverage** proven by a **Playwright** sweep (zero empty pages, zero out-of-demo escapes):
{ M39 ∥ M40 } → M41 → **M42e** → **M42m** (M42e/M42m `iterative`; the rest `section`). **Extended 2026-06-26 with
M43→M46** — the **presenter-grade / scalable-generation extension**: M43 cockpit-UX polish + M44 profile-completeness
(whole-roster members+managers; DATA DENSITY only) [`section`] → M45 generation-engine (cheap-LLM `cmd/gen-batch` +
prompt-keyed cache) → M46 org-scale-fill + gen-batch preview [`iterative`]; execution **{ M43 ∥ M44 } → M45 → M46**,
close-release AFTER M46. Tooling + docs only — zero platform-repo edits. Designed from the live-demo review
[`.agentspace/profile_gaps.md`](../../.agentspace/profile_gaps.md) + workflow `w7t4wq2z4`, and the v1.10-extend
research note [`.agentspace/scratch/roadmap-research-2026-06-26.md`](../../.agentspace/scratch/roadmap-research-2026-06-26.md);
milestone records under [`releases/01.10-method-acting/`](releases/01.10-method-acting/). **Last shipped:** **v1.9
"storytelling"** (2026-06-23, tag `v1.9`; archived under
[`releases/archive/01.90-storytelling/`](releases/archive/01.90-storytelling/)). **Next:**
**`/developer-kit:build-milestone`** (M43 ∥ M44 — the section pair; start M44, it gates M45). _(Live state:
[`state.md`](state.md). Backlog: [`roadmap-vision.md`](roadmap-vision.md).)_

## Project note

Rosetta is the **documentation corpus** for the Anthropos platform (architecture guides + the
`/dev-up` / `/dev-down` / `/stack-update` skills that build, run, and sync the local *dev*
environment — converged from the former setup/start/update skills in v1.3/M14). The planning lifecycle tracked here governs **extensions to rosetta itself** — the
first being a second corpus + skill set for building disposable, fully-seeded **demo** environments.
It does **not** track changes to the Anthropos platform repos (those live under `anthropos-work`).
