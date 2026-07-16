---
active_release: "v2.4 casting call — the recruiter-vantage / hiring-org release (designed 2026-07-15)"
active_branch: "release/02.40-casting-call"
active_milestone: "M226 opening-night (planned — iterative: prove the 7-condition hiring gate live on billion over the tailnet, default /demo-up, no flags — recruiter p95 click→ACCESS < 5 s as a 3rd measured vantage)"
last_closed: "M225 — 2026-07-17"
phase: "M225 CLOSED (section, complete; merged --no-ff into release/02.40-casting-call; the demo-integration milestone — hiring auto-set-dress guard + coverage gate [3 seats MET] + one GREEN recruiter playthrough + pt-world Org D); next M226 opening-night (the live billion proof)"
last_updated: "2026-07-17"
---

# State

**v2.4 "casting call" — IN DEVELOPMENT** (designed 2026-07-15 via `/developer-kit:design-roadmap`; branch
`release/02.40-casting-call` cut from `main`; tag will be `v2.4`). The **recruiter-vantage / hiring-org release**:
a **NET-NEW** 4th, **HIRING** demo org on the presenter cockpit where **45 candidates audition on the same 5
positions and a manager compares them side by side**, distinct from the three workforce orgs. **M222 CLOSED
2026-07-15 (GO).** **M223 "casting the ensemble" CLOSED 2026-07-16** (the 4th hiring story + the HiringConfig/
HiringFunnel seeders → the `local_jobsimulation_sessions` MIRROR pair). **M224 "the callback" CLOSED 2026-07-16
(closed-on-gate)** — the render payoff: via the **TOK-02 two-app demo** (the genuine `apps/hiring` as a 2nd UI
container), the recruiter's Results scoreboard paints **20/page × 43** comparable candidates per each of the 5
shared sims (faithful pagination, GATE-DECISION D1), non-degenerate, 0 junk, 0 prod-eject, ≥3 cold runs + 4/4
flake; the hero trio (Rae/Cara/Cody) resolves; Clerkenstein emits org `publicMetadata.isHiring` (conditional-emit,
`/align-run` 100/100). **M225 "dress the set" CLOSED 2026-07-17** (section, complete) — the hiring org comes up
**auto-set-dressed on a default `/demo-up`**: the S1 bring-up guard + the S2 coverage gate (**3 seats MET**,
`manifestFor` 3-arg + `profileGated`) + **one GREEN recruiter playthrough** (`pt-hiring-recruiter-compare`) on
**pt-world Org D "Kestrel Hiring Group"**; the S4 corpus docs. **Next: `/developer-kit:build-mstone-iters` → M226
opening-night (prove the 7-condition hiring gate live on `billion`).**

## Active release — v2.4 "casting call"

**Theme.** *The recruiter's vantage — 45 candidates audition on the same 5 positions; a manager compares them side
by side, distinct from the three workforce orgs on the cockpit.*

**Reverses v2.3's D-DESIGN-4** (*"no hiring org, none will be built"*): the stated blocker — *"needs the `hiring-app`
frontend, not in the demo UI tier"* — was **refuted**. The candidate-comparison surface ships in the **dockerized
`apps/web` (Workforce)** app the demo already builds (not the Vercel-only `apps/hiring`), and the domain primitives
(`organizations.is_hiring`, the `candidate` role, `SIMULATION_TYPE_HIRING` sessions) already exist. **Consumes the
recruiter/seeder half of the reserved vision M205** (Stripe-tier-gate + ATS-pipeline half stays a residual vision
reservation).

**Shape — 5 milestones, largely SEQUENTIAL** (each consumes the prior's output):

```
M222 ──→ M223 ──→ M224 ──→ M225 ──→ M226
read the casting  the      dress    opening
room     ensemble callback the set  night
section  section  iter     section  iter
```

- **M222 read-the-room** (`section`) — ✅ **DONE 2026-07-15 (GO).** The hiring-model spike: delivered
  `corpus/services/hiring.md` (BLIND AREA), **proved by rendering** the comparison surface is demo-servable +
  renders a comparable score from seedable data (BA-3 refuted), landed the `is_hiring` gate + `narrative: hiring`.
  **The go/no-go barrier cleared** — the mirror-table trap named (score = `local_jobsimulation_sessions`, not
  `jobsimulation.sessions`); no HiringSeeder/funnel (→ M223).
- **M223 casting-the-ensemble** (`section`) — the `HiringSeeder`: exactly 5 `admin` + 45 `candidate`, 5 shared REAL
  replayed positions (`directus.job_position` snapshot extension), the realistic funnel, closure/isolation wiring.
- **M224 the-callback** (`iterative`) — ✅ **DONE 2026-07-16 (closed-on-gate).** Cockpit hero trio (Rae/Cara/Cody)
  + Clerkenstein `isHiring` FAPI (conditional-emit, `/align-run` 100/100) + the **TOK-02 two-app demo** (the real
  `apps/hiring` as a 2nd UI container). The recruiter's Results scoreboard paints **20/page × 43** per each of the
  5 sims (faithful pagination, D1), non-degenerate, 0 junk, 0 prod-eject, ≥3 cold + 4/4 flake. 4 hiring-image
  demo-patches; 0 platform edits.
- **M225 dress-the-set** (`section`) — ✅ **DONE 2026-07-17.** Hiring auto-set-dress bring-up GUARD + a hiring
  coverage manifest (3 seats GATE MET, `manifestFor` 3-arg + `profileGated` on `apps/hiring`) + **one GREEN recruiter
  playthrough** (`pt-hiring-recruiter-compare`) on **pt-world Org D "Kestrel Hiring Group"**. No `job_position` replay
  (KB-1). 0 platform edits.
- **M226 opening-night** (`iterative`) — prove the 7-condition gate live on `billion` over the tailnet, default
  `/demo-up`, no flags (recruiter p95 click→ACCESS < 5 s as a 3rd measured vantage).

**Binding user decisions (2026-07-15):**
- **Genuine hiring org** (`is_hiring=true` end-to-end) — Clerkenstein `isHiring` wiring IS in scope (M224); M222
  de-risks the flag-flip blast radius (R5, the two `isEnterprise` definitions). A break → sha-pinned demo-patch or
  ESCALATE, never a platform edit.
- **Real replayed positions + a realistic non-degenerate funnel** — MOST candidates assessed on the 5 shared
  positions (comparison populated + comparable), SOME assigned-not-taken. NOT a flat 225-session grid.
- **Cockpit heroes = 1 manager + 2 candidates**, login-only. The 2 candidates show two funnel states (one assessed,
  one only-assigned); the candidate heroes render usable assessed `/profile` surfaces.

**Go / no-go note.** M222 is a hard barrier: the comparison surface is *inferred* (not render-proven) to be in the
dockerized `apps/web`. If M222's render-probe shows it is **`apps/hiring`-only**, the release **ESCALATES rather than
proceeding** — showing a Vercel-only app would be a large net-new build + a platform edit, out of the zero-edit wall.

**Hard constraints (carried, unchanged):** **zero platform-repo edits** — a platform-source render gate routes to a
sha-pinned `demopatch` or escalates; all stack-operating tooling lives in `rosetta-extensions` (authored in
`.agentspace/rosetta-extensions/`, tagged, consumed per-stack at a pinned tag).

Full design (file:line-cited proposal): [`releases/02.40-casting-call/design-notes.md`](releases/02.40-casting-call/design-notes.md).
Milestone contracts: [`releases/02.40-casting-call/`](releases/02.40-casting-call/).

## Recently shipped (releases, newest first — max 3)

- **v2.3 "cue to cue"** — 2026-07-15 (tag `v2.3`). The **presenter-speed** release: **click→ACCESS < 5 s proven
  live 8/8 on `billion`** over the tailnet, no flags — **login p95 2.11 s / 1.31 s** vs a ~39/38 s baseline (~18×).
  Demo comes up green, full, remote-default-on; AI-readiness renders filled; `safety.md` Part 3; the ~24-instance
  **D17** thread told honestly. 5 milestones M217→{M218∥M219∥M220}→M221; tooling + docs only, 0 platform edits.
  The `billion` demo is **LEFT LIVE** (a free M222 render substrate).
- **v2.2 "panorama"** — 2026-07-12 (tag `v2.2`). External-shareability / Tailscale-serve: dev/demo stacks
  reachable from another machine on a tailnet over one trusted HTTPS origin. Opt-in default-off (flipped to
  demo-default-on at v2.3 M220). First live remote Linux-VM deploy.
- **v2.1 "quick change"** — 2026-07-09 (tag `v2.1`). The skiller-in-app re-ground: re-fit tooling + corpus +
  stacks to the merged platform (skiller → `app`/`public`, RPC → backend, 4 subgraphs); proved dev-up + demo-up
  cold.

## Headline numbers (v2.4 M225 close, 2026-07-17)
- **p95 click→ACCESS (v2.3 headline gate):** **2.11 s** (employee) / **1.31 s** (manager) vs the **< 5000 ms** gate,
  on `billion` over the tailnet, cold reset-to-seed. v2.4 extends this to a **3rd (recruiter) vantage** at M226.
- **Go test funcs:** **1887** (+2 vs M224's 1885 — the M225 playthroughs `corpus_test.go` pin + `hiring_isolation_test.go`
  pt-world Org D isolation invariant; all modules `go vet` clean, **0 Go failures**).
- **M225-touched suites (re-verified GREEN at merge base):** `test_verify.py` **124** (incl `shellcheck`; +4 autoverify
  floor-boundary fences) · TS unit **61** (stack-verify/e2e) + **69** (playthroughs/e2e) · both `tsc --noEmit` clean ·
  **flake 5/5** · the live hiring playthrough independently orchestrator-verified GREEN on a clean reset-to-seed.
- **Inherited (non-milestone) carries:** demo-stack **650 pass / 8 pre-existing fail** (test-debt backlog) + the M204
  assign-WRITE declared TODO → both re-fated to the v2.4 release close.
- **Alignment (Clerkenstein):** **100% / 100% critical** — held since M224 (M225 touched no alignment surface).
- **Flake:** **0** (milestone-owned). **Platform-repo edits:** **0.** **Supply chain:** GREEN — 0 net-new direct deps.

## D17 — the carried-forward signature hazard (v2.4 discipline)

**D17: *a status artifact that outlives the thing it describes, and is then read as evidence.*** It bit ~24 times
across M217→M221. **The keeper:** ***a named hazard is not a fence; only an executable probe binds.*** Directly
relevant to v2.4 M224: the render gate must be proven by an executable render probe, never by "the seed wrote the
rows." Full arc: [`releases/archive/02.30-cue-to-cue/release-retro.md`](releases/archive/02.30-cue-to-cue/release-retro.md).

## Branch model / shipped tags
**v2.4 IN DEVELOPMENT:** `release/02.40-casting-call` cut from `main` (2026-07-15); milestone branches
`m222/read-the-room … m226/opening-night` will branch from it per milestone. **v2.3 CLOSED** (the mechanical
`release/02.30-cue-to-cue → main` merge + the `v2.3` tag remain `/developer-kit:close-release` Phase 11's job).
**Shipped tags:** **v2.2** `v2.2` · **v2.1** `v2.1` · **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` ·
**v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` ·
**v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`. (Full detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog (unscheduled, cross-release)
- **v2.3 tail carries → v2.4 (non-gate; user signed off 2026-07-15 at v2.3 close-release; a SEPARATE track from the
  hiring theme — they do NOT overlap it):** **F4** (academy grid renders 0 cards — fix is in the `ant-academy`
  **platform repo**, out of zero-platform-edit scope) · **BURNIN-M221-dev-public-host** (dev-path live burn-in) ·
  **F-M220-4** (academy re-run on a live public-host demo) · **PROBE-M218-c3-rerun** (router-403 re-check) — the
  last three all need **live infra**. Parked or ride v2.4 as a side track; not folded into the milestone chain.
- **Test-debt + declared TODO (carried, non-gate; M224 D6 + M225 D-AUDIT — both re-fated fresh 2026-07-17, routed to
  the v2.4 release close):** (a) 8 pre-existing demo-stack failures — 6 × `test_cockpit.py` (4 removed-academy-CTA + 2
  v2.3.1 overlay-JS) + `test_purge` + `test_reap`; HEAD-identical, in files M224/M225 never touched, predating v2.4 →
  a future demo-stack test-debt harden pass; (b) the **M204 `assign-and-track.UC1` assign-WRITE** declared in-manifest
  `unimplemented` build-reference gap → its declared-TODO fate at release close.
- **Plan hygiene → next close-release:** `metrics-history.md` still lacks **v2.0 + v2.2** rows.
- **Older, still unscheduled:** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01 (`replayCmd` hermetic
  test), CAVEAT-1 (clean-box literal full `/dev-up`), M314b (prod frozen-read hydration — a prod-team follow-up).
  Reserved **Playthroughs futures** M206–M207 stay in vision; **M205** is now CONSUMED-in-intent by v2.4 (recruiter/
  seeder half) with the tier-gate/ATS half residual. All tracked in [`roadmap-vision.md`](roadmap-vision.md).

_Last updated: 2026-07-17 (**M225 "dress the set" CLOSED** via /developer-kit:close-milestone — section, complete;
merged `--no-ff` into `release/02.40-casting-call`. The demo-integration milestone: hiring auto-set-dress bring-up
GUARD + a hiring coverage gate [**3 seats MET**, `manifestFor` 3-arg + `profileGated` on `apps/hiring`] + **one GREEN
recruiter playthrough** [`pt-hiring-recruiter-compare` on pt-world Org D "Kestrel Hiring Group"] + the S4 corpus docs.
Go funcs 1885→**1887**, `test_verify.py` 124, TS unit 61+69, flake **5/5**, **0 platform-repo edits**. Deferral audit
**YELLOW** [0 new; 2 inherited carries re-fated + routed to release close]. rext code-of-record
`casting-call-m225-harden` [`be431c3`]; live tag `casting-call-m225-sections`. M225 is NOT the final v2.4 milestone —
no main-merge/tag. **NEXT: /developer-kit:build-mstone-iters → M226 opening-night.**)_
