# Plan — Context

This directory holds the **active** planning artifacts for **Project Rosetta**. It was bootstrapped
on 2026-06-02 to put rosetta on the developer-kit planning lifecycle. **`state.md` is the live source of
truth** — this file is the stable orientation/conventions doc; when the two disagree, `state.md` wins.

**Status (2026-07-20):** **v2.5 "the playbill" FEATURE-COMPLETE, IN RELEASE CLOSE** — the **content-vantage
release** (branch `release/02.50-the-playbill`, designed 2026-07-19 via `/developer-kit:design-roadmap`; **8 milestones
M229 → M230 → M231 (HARD go/no-go) → M232 → M233 → M234 → M235 → M236**, spike-first; tag will be `v2.5`). All 8 are
**closed and MERGED**; `/developer-kit:close-release` is running (the `release → main` merge + tag are its Phase 11,
which the USER runs). Two threads on the mature demo/cockpit machinery: **A** fills the empty **ant-academy** grid
(DB-authoritative catalog; production-faithful fill); **B** adds a 2nd **"Content stories"** cockpit tab of **played
sessions** per content product, each with **as-player / as-manager** login-and-land, **cloned from anonymized real
production sessions**, re-tenanted, **source-pinned by prod session-id**. Binding user decisions: real customer-session
sourcing **accepted** as the data-controller call · demos **VPN/tailnet-scoped**, the release **amends `safety.md`
Part 3** · academy fill production-faithful (no "Draft" chip). Tooling + docs only — **zero platform-repo edits**.
M236's gate was met cold on `billion` (**29/29** landable pairs both vantages · 65 academy cards · hero p95
1.22/1.51 s vs 5 s), with the denominator **corrected 31 → 29**.

> ⚠️ **Read `state.md` for the headline caveat before quoting v2.5's number.** The `29/29` is **UNIT-PROVEN, NOT
> LIVE-RE-PROVEN** — the harness was fixed ~10 times *after* the measurement — and **39 live-browser specs went
> unexecuted at the close**. The live re-prove is deferred to **v2.6 as its first work** (reserved **`M237`**) by
> explicit user decision. Full ledger:
> [`releases/02.50-the-playbill/release-deferrals.md`](releases/02.50-the-playbill/release-deferrals.md).

_Prior:_ **v2.4 "casting call" SHIPPED 2026-07-18** (tag `v2.4`) — the **recruiter-vantage / hiring-org release**: a
4th, **HIRING** demo org (45 candidates auditioning on the same 5 positions, compared side by side), reversing v2.3's
D-DESIGN-4; 7 milestones **M222 → M226**, RE-OPENED 2026-07-17 for believability → **M227 → M228**. It consumed the
recruiter/seeder half of the reserved vision **M205**. ⚠️ **Its close ran no release-scope deferral audit and issued
false greens** — the structural cause of 7 of v2.5's 8 aged-out items; its archived `release-review.md` carries a
dated correction annotation. _Prior:_ **v2.3 "cue to cue" SHIPPED 2026-07-15** (tag `v2.3`) — the **presenter-speed
release**, a field-hardening release (the v1.3b / v1.10b / v2.1 lineage) triggered by a live presenter defect (a
cockpit hero-swap took 1–2 MINUTES). The wall was **entirely downstream and already measured in-repo** while the
corpus asserted in 4 places that login was "~2–5 s we can't shorten". Binding decisions: the <5 s gate is on **ACCESS**
(authenticated + interactive shell), not full first-page render; demo remote-access flips to **opt-out** (dev stays
opt-in — consuming the reserved **M216**). _Prior:_ **v2.2 "panorama" SHIPPED 2026-07-12** (tag `v2.2`) — the
**external-shareability / Tailscale-serve release** (4 milestones M212 → { M213 ∥ M214 } → M215 (+ opt M216); opt-in
default-off, HTTPS-everywhere under one MagicDNS origin, demo-first; the sanctioned re-proposal of the dropped v1.4
Tailscale/ingress seed). _Prior:_ **v2.1 "quick change" SHIPPED 2026-07-09** (tag `v2.1`) — the **skiller-in-app
re-ground**, a field-hardening release triggered by a landed platform structural change (the `skiller` service + its
DB schema merged into `app` → the **`public`** schema, RPC → `backend`, the skiller subgraph gone → **4 subgraphs**);
4 milestones **M208 → M211**, strictly sequential. _Prior:_ **v2.0 "opening night" SHIPPED 2026-07-02** (tag `v2.0`) —
the first v2.x release, which opened the **Playthroughs** pillar (functional-flow e2e *testing*; M201 ∥ M202 →
{ M203 ∥ M204 }; 10 live Playthroughs GREEN on cold reset-to-seed, **1 declared in-manifest TODO** — the assign-WRITE
half, still open and now routed to reserved **`M238`**). The whole **v1.x major** (v1.0 … v1.10 + the `v1.10.1`
backfill) shipped before it; records archived under [`releases/archive/`](releases/archive/), v1.x history in
[`roadmap-legacy.md`](roadmap-legacy.md).

Genuinely-deferred work stays **unscheduled backlog** — DEF-M10-01 (cloud store / S3 blob bytes), DEF-M21-01
(replayCmd hermetic test), CAVEAT-1, M314b (prod frozen-read hydration), `DEF-M215-03(a)`/`F11` (seed hero
identity-key polish) ([`roadmap-vision.md`](roadmap-vision.md)); the reserved Playthroughs futures **M206–M207** and
the **M205**-residual (tier gates + ATS pipeline) stay in vision. Live state: [`state.md`](state.md). (The active
[`roadmap.md`](roadmap.md) holds the **active** release only — the shipped `## Done` sections for v1.10b + v2.0 → v2.4
were split out to [`roadmap-archive-v2.0-v2.4.md`](roadmap-archive-v2.0-v2.4.md) at the v2.5 close, under the
`roadmap-legacy.md` precedent; v1.x history stays in [`roadmap-legacy.md`](roadmap-legacy.md).)

## Files

- [`roadmap.md`](roadmap.md) — the **active major** (its milestones, execution graph, risks)
- [`roadmap-legacy.md`](roadmap-legacy.md) — the **retired v1.x** roadmap (M0 … M46, all SHIPPED) — created at
  the v2.0 opening when the v1.x major retired
- [`roadmap-vision.md`](roadmap-vision.md) — future versions + future v2 milestones + proposals not yet in active development
- [`state.md`](state.md) — current/next milestone, last update
- `releases/{VV.VV}-{codename}/m{N}-{slug}/overview.md` (active version) → `releases/archive/{VV.VV}-{codename}/…` (shipped). _The **ACTIVE v2.5 dirs** are [`releases/02.50-the-playbill/`](releases/02.50-the-playbill/) (`m229-academy-content-model-re-ground/` … `m236-prove-on-billion/`), scaffolded by the 2026-07-19 `/developer-kit:design-roadmap` run, plus the release-root artifacts `metrics.json`, `dependencies.lock`, `release-review.md`, **`release-deferrals.md`** (the per-item deferral ledger, new at this close) and `audit-deferrals/`. Every SHIPPED release's dirs live under [`releases/archive/`](releases/archive/) — `02.40-casting-call/`, `02.30-cue-to-cue/`, `02.20-panorama/`, `02.10-quick-change/`, `02.00-opening-night/`, `01.10b-fit-up/`, and `01.{00..10}-{codename}/` for the v1.x major, each with overview/progress/decisions/retro/metrics. **They are archived at release close** — if a version's dirs are still under `releases/` it has not shipped._
- [`roadmap-archive-v2.0-v2.4.md`](roadmap-archive-v2.0-v2.4.md) — the shipped `## Done` sections for **v1.10b + v2.0 → v2.4**, split out of `roadmap.md` at the v2.5 close (2026-07-20) under the `roadmap-legacy.md` precedent, when `roadmap.md` had reached 2,079 lines / 203 KB with ~77% of it shipped history

## Conventions

- One directory per milestone, named `m{N}-{slug}/`
- Each milestone dir has at minimum an `overview.md`. As the milestone progresses, optional companion files: `progress.md` (deliverable checklist), `decisions.md` (implementation choices with rationale), `spec-notes.md` (technical details).
- Status values: `planned` → `in-progress` → `done` → `archived` (terminal, set at release close).
- Milestone numbering — **v1.x (the first major) = flat sequential** (M0 … M46, closed + archived; detail below). **v2+ = the `Mxyy` scheme** (`M` + major digit + two-digit milestone): **v2.0 "opening night" = M201 ∥ M202 → { M203 ∥ M204 }** (M201 `iterative`+user-guided manifest corpus ∥ M202 `section` Playthroughs-foundation → M203/M204 `iterative` per-vantage coverage — the new-major scheme `context.md` had reserved for "a future *major* v2+"); **v2.1 "quick change" = M208→M211** (strictly sequential); **v2.2 "panorama" = M212 → { M213 ∥ M214 } → M215** (+ optional M216) — the next free `Mxyy` band after M211, reserved Playthroughs futures M205–M207 kept in vision; **v2.3 "cue to cue" = M217 → { M218 ∥ M219 ∥ M220 } → M221** — the counter resumes at **M217** because **M216 stayed reserved** (dev-path Tailscale parity) and was never scaffolded; v2.3's **M220(d) CONSUMES that reserved scope** (the dev-side opt-in `--public-host`, per D-DESIGN-3) rather than renumbering, so **M216 is retired as a reservation, not built as a milestone**; **v2.4 "casting call" = M222 → M223 → M224 → M225 → M226** (largely sequential — the next free `Mxyy` band after M221; the reserved **M205** is CONSUMED-in-intent by v2.4's recruiter/seeder half [tier-gate + ATS half residual], **M206–M207** stay reserved in vision) — the recruiter-vantage / hiring-org release adds a **4th story org** (the HIRING org, `is_hiring=true`) to the multi-org Stories & Heroes model, gated on a new `is_hiring` blueprint field + the `narrative: hiring` discriminator (M222/M223 `section`; M224/M226 `iterative`; M225 `section`), **RE-OPENED 2026-07-17 for believability → M227 → M228**; **v2.5 "the playbill" = M229 → M230 → M231 (HARD go/no-go) → M232 → M233 → M234 → M235 → M236** (spike-first, the next free `Mxyy` band after M228 — the content-vantage release; M229/M231–M234 `section`, M230/M235/M236 `iterative`). **RESERVED for v2.6, numbered at the v2.5 close (2026-07-20) rather than at design time:** **M237** re-prove-on-billion (v2.6's declared FIRST work — the deferred live re-prove of v2.5's headline `29/29` + the 39 unexecuted live-browser specs + six batched carries) and **M238** playthrough-assign-write. These are **reservations** under the same precedent as vision M205–M207 and the retired M216, and may be renumbered at the v2.6 `/developer-kit:design-roadmap` run — the deliberate exception to *"a milestone gets its number at design time"*, taken because the v2.5 close named the class ***a fate needs a MILESTONE, not a phase, a pass, or "the next X"*** (three of that release's eight aged-out items failed exactly that way, and one aged out TWICE against the phrase "the next prove-on-VM"). The v1.x flat-counter detail follows: M0, M1, …, M20, v1.5 = **M21→M25**; **M26** = self-contained-demo, now **v1.8 "understudy"** (re-implemented onto current `main` from the orphaned `m26/self-contained-demo` ext branch @ `25ab855` / tag `prop-room-m26`, authored 2026-06-14 — that orphan is the spec, NOT merged: it predates v1.6/v1.7 which rewrote the same files); v1.6 "stage door" = **M27→M30**; v1.7 "house lights" = **M31→M32** (M33 ant-academy liveness → backlog); v1.8 "understudy" = **M26** (the reserved slot); v1.9 "storytelling" = **M34→M38** (the seeding/Stories-&-Heroes release — the counter resumes at M34, the next free number after the M33 backlog slot); v1.10 "method acting" = **M39→M46** (the believable-profile release **M39→M42m** [the counter resumes at M39; **M42e/M42m** are an `e`/`m` **persona-pair split** [employee/manager] of one planned coverage milestone — the second split-suffix use after M7a/M7b/M7c], **extended 2026-06-26 with M43→M46** — the presenter-grade / scalable-generation extension: M43 cockpit-UX + M44 profile-completeness [`section`] → M45 generation-engine + M46 org-scale-fill [`iterative`]); **v1.10b "fit-up" = M47→M53** (the interposed field-hardening backfill — the flat counter, thought closed at M46, **RE-OPENS** here because backfill work is a `.1` patch of v1.10, not a v2 `Mxyy` milestone: M47 re-sync/recapture · M48 corpus-reground · M49 bring-up-hardening [all `section`] · M50 content-fill ∥-conceptually · M51 AI-readiness-org [both `iterative`] · M52 seed-manifest · M53 cold-rebuild [`section`]) (the milestone counter never resets — M26 was reserved for the self-contained-demo effort at v1.6 design, so stage-door begins at M27 even though M26 ships LATER, in v1.8; the version *number* jumps 1.3b→1.5 after the v1.4 removal; there is no `M5xx` scheme — that two-digit `Mxyy` scheme only begins at a future *major* v2+). A letter suffix has two uses: (1) a milestone **inserted after** the fact (M1b drift CI, M2b consolidation, M2c the iterative `@clerk/express` feature); and (2) a **split** of one planned milestone into a sequential mini-arc (**M7a → M7b → M7c** = the former M7 "seeding" split into framework+safety / data-DNA / fleet, 2026-06-04). Context disambiguates which.
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

**Active:** **v2.5 "the playbill" — FEATURE-COMPLETE, IN RELEASE CLOSE** (branch `release/02.50-the-playbill`, cut
from `main` 2026-07-19; tag will be `v2.5`). The **content-vantage release** — *show the platform's content, played.*
**8 milestones, spike-first:** **M229** academy-content-model-re-ground [`section`] → **M230** academy-demo-fill
[`iterative`] → **M231** content-stories-feasibility-spike [`section` — a **HARD go/no-go barrier**; verdict **GO**:
the sim result page reads a **PERSISTED** row, so a cloned session renders] → **M232** session-clone-sourcing-seeder
[`section` — the `ContentStorySeeder` copies real prod sessions + best-effort PII scrub + re-tenant + source-pin] →
**M233** content-stories-manifest [`section` — `BuildContentProducts`, honesty-gated, fail-closed] → **M234**
content-stories-cockpit-tab [`section` — the render half + `content-player-<idx>` seats] → **M235** prove-it-lands
[`iterative`, closed-incomplete/pragmatic — LIVE gate → M236 by design] → **M236** prove-on-billion [`iterative`,
closed-on-gate — **29/29** landable pairs cold on `billion`, denominator corrected 31 → 29]. All 8 **closed and
MERGED**; `/developer-kit:close-release` is running. Records under
[`releases/02.50-the-playbill/`](releases/02.50-the-playbill/).
**Next:** **`/developer-kit:close-release` Phase 11** — the `release → main` merge + the `v2.5` tag (**the USER runs
it**), then `/developer-kit:design-roadmap` for **v2.6**, whose **declared first work** is the reserved **`M237 —
re-prove-on-billion`** (the deferred live re-prove of v2.5's headline number + the 39 unexecuted live-browser specs +
six batched carries — see [`roadmap-vision.md`](roadmap-vision.md) § v2.5 → v2.6 carry).

_Shipped-release narratives are no longer duplicated here_ — they live one paragraph each in `roadmap.md`
§ Shipped releases, in full in [`roadmap-archive-v2.0-v2.4.md`](roadmap-archive-v2.0-v2.4.md) (v1.10b + v2.0 → v2.4)
and [`roadmap-legacy.md`](roadmap-legacy.md) (the v1.x major), and per-release under
[`releases/archive/`](releases/archive/). This file kept a growing inline chain of them for six releases; at the v2.5
close (finding `D-3`) that chain was found **two releases stale** — it still called v2.4 "IN DEVELOPMENT" and pointed
"Next" at M222, a milestone that had shipped two releases earlier. **The orientation doc must not restate status it
does not own.** _(Live state: [`state.md`](state.md). Backlog + deferral destinations:
[`roadmap-vision.md`](roadmap-vision.md).)_

## Project note

Rosetta is the **documentation corpus** for the Anthropos platform (architecture guides + the
`/dev-up` / `/dev-down` / `/stack-update` skills that build, run, and sync the local *dev*
environment — converged from the former setup/start/update skills in v1.3/M14). The planning lifecycle tracked here governs **extensions to rosetta itself** — the
first being a second corpus + skill set for building disposable, fully-seeded **demo** environments.
It does **not** track changes to the Anthropos platform repos (those live under `anthropos-work`).
