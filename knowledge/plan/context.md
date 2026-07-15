# Plan — Context

This directory holds the **active** planning artifacts for **Project Rosetta**. It was bootstrapped
on 2026-06-02 to put rosetta on the developer-kit planning lifecycle. **`state.md` is the live source of
truth** — this file is the stable orientation/conventions doc; when the two disagree, `state.md` wins.

**Status (2026-07-15):** **v2.4 "casting call" IN DEVELOPMENT** — the **recruiter-vantage / hiring-org release**
(branch `release/02.40-casting-call`, designed 2026-07-15 via `/developer-kit:design-roadmap`; **5 milestones
M222 → M223 → M224 → M225 → M226**, largely sequential; tag `v2.4`). A **NET-NEW** release that adds a 4th, **HIRING**
demo org to the presenter cockpit (45 candidates audition on the same 5 positions; a manager compares them side by
side), **reversing v2.3's D-DESIGN-4** — the candidate-comparison surface ships in the dockerized `apps/web`, not the
Vercel-only `apps/hiring`. **M222 is a HARD go/no-go barrier** (escalates if the surface is `apps/hiring`-only).
Binding user decisions: genuine `is_hiring=true`, real replayed positions + a realistic non-degenerate funnel,
cockpit heroes 1 manager + 2 candidates. **Consumes the recruiter/seeder half of the reserved M205.** Tooling + docs
only — zero platform-repo edits. _Prior:_ **v2.3 "cue to cue" SHIPPED 2026-07-15** (tag `v2.3`) — the
**presenter-speed release** (branch `release/02.30-cue-to-cue`, designed 2026-07-13 via
`/developer-kit:design-roadmap`; **5 milestones M217 → { M218 ∥ M219 ∥ M220 } → M221**). A **field-hardening
release** (the v1.3b / v1.10b / v2.1 lineage)
triggered by a **live presenter defect**: a cockpit hero-swap takes **1–2 MINUTES**. The investigation cleared
Clerkenstein (the handshake leg is a zero-I/O 303) and found the wall **entirely downstream — already measured
in-repo** (76 s members grid, 84 s router latency, a 180 s AI-readiness read) while **the corpus asserted in 4 places
that login is "~2–5 s we can't shorten"**, and **two `app` perf demo-patches silently REFUSE on sha-drift with the
reason piped to `/dev/null`**. Binding user decisions: the <5 s gate is on **ACCESS (authenticated + interactive
shell), not full first-page render**; demo remote-access flips to **opt-out** (dev stays opt-in — consuming the
reserved **M216**); the three story orgs are **the three that already exist** (no hiring org). **Tooling + docs only —
zero platform-repo edits.** _Prior:_ **v2.2 "panorama" SHIPPED 2026-07-12** (tag `v2.2`) — the **external-shareability
/ Tailscale-serve release**
(branch `release/02.20-panorama`, designed 2026-07-11 via `/developer-kit:design-roadmap`; **4 milestones M212 →
{ M213 ∥ M214 } → M215** (+ opt M216); opt-in default-off, HTTPS-everywhere under one MagicDNS origin, demo-first;
tooling + docs + an opt-in flag only, a 2-item patch tail via the rext mechanism; the sanctioned re-proposal of the
dropped v1.4 Tailscale/ingress seed). Prior: **v1.0 … v1.10 + v1.10b + v2.0 + v2.1 SHIPPED** (the whole **v1.x
major** tagged `v1.0` … `v1.10`
+ the `v1.10.1` backfill; the first v2.x releases **v2.0 "opening night"** (`v2.0`) + **v2.1 "quick change"** (`v2.1`) shipped; records archived
under [`releases/archive/`](releases/archive/), v1.x history in [`roadmap-legacy.md`](roadmap-legacy.md)). **v2.1
"quick change" SHIPPED 2026-07-09 (tag `v2.1`)** — the **skiller-in-app re-ground**, a **field-hardening release**
(v1.3b "dress rehearsal" / v1.10b "fit-up" lineage) triggered by a **landed platform structural change** (designed 2026-07-08 via
`/developer-kit:design-roadmap`; branch `release/02.10-quick-change`; tag `v2.1`; **4 milestones M208 → M209 → M210
→ M211**, strictly sequential): the `skiller` service + its DB schema merged into `app` (domain → the **`public`**
schema, table names unchanged `skiller.X → public.X`; RPC → `backend`; the skiller GraphQL subgraph gone → **4
subgraphs**; skiller repo/container removed). It re-fits the **rext tooling** (the `skiller.<table>` → `public.<table>`
re-point + recapture), the **corpus** (completing the colleague's `origin/docs/skiller-in-app-merge` sweep in
lockstep), and the **stacks** (re-sync to the merged platform), then **proves `/dev-up` + `/demo-up` still work** on
the merged platform (the iterative M211 acceptance gate). **Tooling + docs + stack-re-sync only — zero platform-repo
edits** (the platform already did its half). The prior release, **v2.0 "opening night"** (SHIPPED 2026-07-02, tag
`v2.0`), opened the **Playthroughs** pillar (functional-flow e2e *testing*; M201 ∥ M202 → { M203 ∥ M204 }; 10 live
Playthroughs GREEN on cold reset-to-seed). Genuinely-deferred work stays **unscheduled backlog** — DEF-M10-01 (cloud
store / S3 blob bytes), DEF-M21-01 (replayCmd hermetic test), M25-D9 (dev taxonomy rc=4 — now on the M208 re-sync
path → opportunistic Fate-1), M314b (prod frozen-read hydration) ([`roadmap-vision.md`](roadmap-vision.md)); the
reserved **Playthroughs futures** M205–M207 stay in vision (v2.1 takes M208+). Live state: [`state.md`](state.md).
(The active [`roadmap.md`](roadmap.md) holds the v2.x major — v2.0 "Done" + v2.1 "In Development"; v1.x history is in
[`roadmap-legacy.md`](roadmap-legacy.md).)

## Files

- [`roadmap.md`](roadmap.md) — the **active major** (its milestones, execution graph, risks)
- [`roadmap-legacy.md`](roadmap-legacy.md) — the **retired v1.x** roadmap (M0 … M46, all SHIPPED) — created at
  the v2.0 opening when the v1.x major retired
- [`roadmap-vision.md`](roadmap-vision.md) — future versions + future v2 milestones + proposals not yet in active development
- [`state.md`](state.md) — current/next milestone, last update
- `releases/{VV.VV}-{codename}/m{N}-{slug}/overview.md` (active version) → `releases/archive/{VV.VV}-{codename}/…` (shipped). _The **ACTIVE v2.3 dirs** are [`releases/02.30-cue-to-cue/`](releases/02.30-cue-to-cue/) (`m217-clean-stage/`, `m218-seat-change/`, `m219-readiness-renders/`, `m220-cue-sheet/`, `m221-prove-on-billion/`), scaffolded by the 2026-07-13 `/developer-kit:design-roadmap` run. The **shipped v2.2 dirs** are archived under [`releases/archive/02.20-panorama/`](releases/archive/02.20-panorama/) (`m212-public-host-knob/`, `m213-auth-over-tailnet/`, `m214-origins-and-links/`, `m215-prove-on-odyssey/`), scaffolded by the 2026-07-11 `/developer-kit:design-roadmap` run + archived at the 2026-07-12 close (M216 stays roadmap-only until promoted). The **shipped v2.1 dirs** are under [`releases/archive/02.10-quick-change/`](releases/archive/02.10-quick-change/) (`m208-resync-groundtruth/`, `m209-rext-reground/`, `m210-corpus-reground/`, `m211-bringup-acceptance/`), scaffolded by the 2026-07-08 `/developer-kit:design-roadmap` run. The **shipped v2.0 dirs** are archived under [`releases/archive/02.00-opening-night/`](releases/archive/02.00-opening-night/) (`m201-manifest-corpus/` + the foundation/coverage dirs); the **shipped v1.10b dirs** under [`releases/archive/01.10b-fit-up/`](releases/archive/01.10b-fit-up/) (`m47-resync-recapture/` … `m53-cold-rebuild-acceptance/`). v1.x's shipped milestone dirs are archived under `releases/archive/01.{00..10}-{codename}/`, each with overview/progress/decisions/retro/metrics._

## Conventions

- One directory per milestone, named `m{N}-{slug}/`
- Each milestone dir has at minimum an `overview.md`. As the milestone progresses, optional companion files: `progress.md` (deliverable checklist), `decisions.md` (implementation choices with rationale), `spec-notes.md` (technical details).
- Status values: `planned` → `in-progress` → `done` → `archived` (terminal, set at release close).
- Milestone numbering — **v1.x (the first major) = flat sequential** (M0 … M46, closed + archived; detail below). **v2+ = the `Mxyy` scheme** (`M` + major digit + two-digit milestone): **v2.0 "opening night" = M201 ∥ M202 → { M203 ∥ M204 }** (M201 `iterative`+user-guided manifest corpus ∥ M202 `section` Playthroughs-foundation → M203/M204 `iterative` per-vantage coverage — the new-major scheme `context.md` had reserved for "a future *major* v2+"); **v2.1 "quick change" = M208→M211** (strictly sequential); **v2.2 "panorama" = M212 → { M213 ∥ M214 } → M215** (+ optional M216) — the next free `Mxyy` band after M211, reserved Playthroughs futures M205–M207 kept in vision; **v2.3 "cue to cue" = M217 → { M218 ∥ M219 ∥ M220 } → M221** — the counter resumes at **M217** because **M216 stayed reserved** (dev-path Tailscale parity) and was never scaffolded; v2.3's **M220(d) CONSUMES that reserved scope** (the dev-side opt-in `--public-host`, per D-DESIGN-3) rather than renumbering, so **M216 is retired as a reservation, not built as a milestone**; **v2.4 "casting call" = M222 → M223 → M224 → M225 → M226** (largely sequential — the next free `Mxyy` band after M221; the reserved **M205** is CONSUMED-in-intent by v2.4's recruiter/seeder half [tier-gate + ATS half residual], **M206–M207** stay reserved in vision) — the recruiter-vantage / hiring-org release adds a **4th story org** (the HIRING org, `is_hiring=true`) to the multi-org Stories & Heroes model, gated on a new `is_hiring` blueprint field + the `narrative: hiring` discriminator (M222/M223 `section`; M224/M226 `iterative`; M225 `section`). The v1.x flat-counter detail follows: M0, M1, …, M20, v1.5 = **M21→M25**; **M26** = self-contained-demo, now **v1.8 "understudy"** (re-implemented onto current `main` from the orphaned `m26/self-contained-demo` ext branch @ `25ab855` / tag `prop-room-m26`, authored 2026-06-14 — that orphan is the spec, NOT merged: it predates v1.6/v1.7 which rewrote the same files); v1.6 "stage door" = **M27→M30**; v1.7 "house lights" = **M31→M32** (M33 ant-academy liveness → backlog); v1.8 "understudy" = **M26** (the reserved slot); v1.9 "storytelling" = **M34→M38** (the seeding/Stories-&-Heroes release — the counter resumes at M34, the next free number after the M33 backlog slot); v1.10 "method acting" = **M39→M46** (the believable-profile release **M39→M42m** [the counter resumes at M39; **M42e/M42m** are an `e`/`m` **persona-pair split** [employee/manager] of one planned coverage milestone — the second split-suffix use after M7a/M7b/M7c], **extended 2026-06-26 with M43→M46** — the presenter-grade / scalable-generation extension: M43 cockpit-UX + M44 profile-completeness [`section`] → M45 generation-engine + M46 org-scale-fill [`iterative`]); **v1.10b "fit-up" = M47→M53** (the interposed field-hardening backfill — the flat counter, thought closed at M46, **RE-OPENS** here because backfill work is a `.1` patch of v1.10, not a v2 `Mxyy` milestone: M47 re-sync/recapture · M48 corpus-reground · M49 bring-up-hardening [all `section`] · M50 content-fill ∥-conceptually · M51 AI-readiness-org [both `iterative`] · M52 seed-manifest · M53 cold-rebuild [`section`]) (the milestone counter never resets — M26 was reserved for the self-contained-demo effort at v1.6 design, so stage-door begins at M27 even though M26 ships LATER, in v1.8; the version *number* jumps 1.3b→1.5 after the v1.4 removal; there is no `M5xx` scheme — that two-digit `Mxyy` scheme only begins at a future *major* v2+). A letter suffix has two uses: (1) a milestone **inserted after** the fact (M1b drift CI, M2b consolidation, M2c the iterative `@clerk/express` feature); and (2) a **split** of one planned milestone into a sequential mini-arc (**M7a → M7b → M7c** = the former M7 "seeding" split into framework+safety / data-DNA / fleet, 2026-06-04). Context disambiguates which.
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

**Active:** **v2.4 "casting call" IN DEVELOPMENT** (branch `release/02.40-casting-call`, cut from `main` 2026-07-15;
tag will be `v2.4`). The **recruiter-vantage / hiring-org release** — *a 4th, HIRING org where 45 candidates audition
on the same 5 positions and a manager compares them side by side, distinct from the three workforce orgs on the
cockpit.* **5 milestones (largely sequential):** **M222** read-the-room [`section` — the hiring-model spike + author
`corpus/services/hiring.md` (BLIND AREA) + prove-by-rendering the comparison surface is demo-servable + land the
`is_hiring` gate; a **HARD go/no-go barrier** — escalates if the surface is `apps/hiring`-only] → **M223**
casting-the-ensemble [`section` — the `HiringSeeder`: 5 admins + 45 candidates on 5 shared REAL replayed positions,
the realistic funnel] → **M224** the-callback [`iterative` — cockpit heroes (1 manager + 2 candidates) + Clerkenstein
`isHiring` wiring + drive the comparison render to green: ≥40 comparable non-junk rows per each of 5 sims] → **M225**
dress-the-set [`section` — auto-set-dress `job_position` replay + a hiring coverage manifest + a hiring playthrough]
→ **M226** opening-night [`iterative` — the 7-condition gate live on `billion`, default `/demo-up`, no flags].
Records under [`releases/02.40-casting-call/`](releases/02.40-casting-call/); full proposal at
`releases/02.40-casting-call/design-notes.md`. **Next:** **`/developer-kit:build-milestone`** → **M222**.
_Prior:_ **v2.3 "cue to cue" SHIPPED 2026-07-15** (branch `release/02.30-cue-to-cue`, tag `v2.3`). The
**presenter-speed release** — *a presenter swaps heroes in under 5 seconds, on a demo that
comes up green, fully-loaded, and remotely reachable by default.* **5 milestones:** **M217** clean-stage [`section`
— **a hard barrier**: the last real run's cockpit **crashed on a leaked port** (stale manifest → dead clerk-ids), all
3 snapshot replays **SKIPPED**, autoverify **FAILED**, `jobsimulation` **exits(1)**, and both perf demo-patches
**REFUSED** — *no number taken before M217 lands is trustworthy*] → { **M218** seat-change [`iterative` — the gate:
**p95 click→ACCESS < 5 s**, both vantages, 5 cold reset-to-seed runs] ∥ **M219** readiness-renders [`section` — the
AI-readiness **seed is a verified no-op**; the gap is that it does not **render**] ∥ **M220** cue-sheet [`section` —
the "2 orgs"→3 doc fix, the missing defaults table, the **remote opt-out flip** + **`safety.md` Part 3, the exposure
side**] } → **M221** prove-on-billion [`iterative` — every gate, on the VM, over the tailnet, **no flags**]. **M218
merges first of the parallel three** (it and M220 both touch `up-injected.sh`). Records under
[`releases/02.30-cue-to-cue/`](releases/02.30-cue-to-cue/); the full file:line-cited gap analysis at
`.agentspace/scratch/roadmap-research-2026-07-13.md`. **Next:** **`/developer-kit:build-milestone`** → **M217**.
_Prior shipped record below:_ **v2.2 "panorama" SHIPPED 2026-07-12** (tag `v2.2`; branch
`release/02.20-panorama` merged `--no-ff` → `main`; rext code-of-record `v2.2` = `39e8013`; `.agentspace/rext.tag`
bumped `v2.1`→`v2.2`). The external-shareability / Tailscale-serve release — **M212 → { M213 ∥ M214 } → M215** (+
optional M216, not scaffolded); opt-in default-off, HTTPS-everywhere, demo-first; **tooling + docs + an opt-in flag
only — zero platform-repo edits, 0 net-new deps**; the sanctioned re-proposal of the dropped v1.4 Tailscale/ingress
seed. The first live remote Linux-VM demo deploy, both vantages green on a trusted cert. Records
`releases/archive/02.20-panorama/`; next: **`/developer-kit:design-roadmap`** for the next release. _Prior shipped
record below:_ **v2.1 "quick change" SHIPPED 2026-07-09** (tag
`v2.1`; designed 2026-07-08 via `/developer-kit:design-roadmap`;
branch `release/02.10-quick-change` cut from `main`; tag `v2.1`). The **skiller-in-app re-ground** — a
**field-hardening release** (the v1.3b "dress rehearsal" / v1.10b "fit-up" lineage) triggered by a **landed platform
structural change**: the `skiller` service + its DB schema merged into `app` (domain → the **`public`** schema, table
names unchanged `skiller.X → public.X`; RPC → `backend`; the skiller GraphQL subgraph gone → **4 subgraphs**; skiller
repo/container removed). **4 milestones, strictly sequential:** **M208** Re-sync & merged-schema ground-truth
[`section` ⚠, the foundation — `make pull` both stacks to the merged platform + remove the vestigial `skiller/`
clones + re-migrate against `public` + pin the merge fact-sheet] → **M209** rext tooling re-ground [`section`, the
`skiller.<table>` → `public.<table>` re-point across stack-snapshot + stack-seeding + the small modules, narrow the
cache-key digest, verify the capture column list, recapture the snapshot, tag rext `v2.1`] → **M210** Corpus + skills
re-ground [`section`, complete the colleague's `origin/docs/skiller-in-app-merge` sweep in lockstep with M209 — flip
the rext-facing tooling-doc bodies to `public.*`] → **M211** Bring-up acceptance [`iterative`, the exit gate:
`/dev-up` + `/demo-up` GREEN cold on the merged platform, 0 residual skiller-schema refs, M42 coverage + v2.0
Playthroughs GREEN]. Strictly sequential (the user's execution choice — single-substrate-safe). Tooling + docs +
stack-re-sync only — zero platform-repo edits (the platform already did its half). Records under
[`releases/archive/02.10-quick-change/`](releases/archive/02.10-quick-change/); designed from the user's skiller-merge briefing + the
colleague's docs branch + the 7-agent blast-radius workflow (`wf_08b6bf4a`). **Last shipped:** **v2.0 "opening
night"** (2026-07-02, tag `v2.0`; the first v2.x release — the Playthroughs pillar; records archived under
[`releases/archive/02.00-opening-night/`](releases/archive/02.00-opening-night/)).
**Next:** **`/developer-kit:build-milestone`** → **M212** (the single host knob) — v2.2 is designed, its release
branch is cut, and all four milestone contracts are scaffolded; M212 opens the build. _(Live state:
[`state.md`](state.md). Backlog:
[`roadmap-vision.md`](roadmap-vision.md).)_

## Project note

Rosetta is the **documentation corpus** for the Anthropos platform (architecture guides + the
`/dev-up` / `/dev-down` / `/stack-update` skills that build, run, and sync the local *dev*
environment — converged from the former setup/start/update skills in v1.3/M14). The planning lifecycle tracked here governs **extensions to rosetta itself** — the
first being a second corpus + skill set for building disposable, fully-seeded **demo** environments.
It does **not** track changes to the Anthropos platform repos (those live under `anthropos-work`).
