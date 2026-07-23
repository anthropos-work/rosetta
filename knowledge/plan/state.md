---
active_release: "v2.7 «july jitter» — IN DEVELOPMENT (designed 2026-07-23). The re-ground + fidelity + field-hardening release: realign the demo + corpus to the platform's TRUE current state (the skiller→app merge was one step of a consolidate-every-runtime-engine-into-app program — skillpath now fully decommissioned into app, jobsim mid-merge), and fix what drifted (content-stories manager link · cross-app Back-to-Cockpit · studio logo prod-eject · AI-readiness fidelity · studio builder keys · studio blank-page). 9 milestones M246→M254, barrier → 7-lane fan-out → prove-on-billion. Branch release/02.70-july-jitter; tag will be v2.7."
active_branch: "release/02.70-july-jitter (cut from main 2026-07-23)"
active_milestone: "M251 (fan-out lane; building in an isolated worktree off post-M246 HEAD)"
last_closed: "M246 — 2026-07-23 (v2.7 re-sync & re-point barrier; go/no-go PASS — seeder re-pointed to public.skill_path_sessions, cold demo GREEN on the consolidated platform [561 rows, 3 subgraphs, 0 skillpath], 9-row drift ledger → M247; rext tag july-jitter-m246-harden on origin)"
phase: "M246 closed (barrier PASSED + merged to release/02.70-july-jitter) — M247–M254 fan out off post-M246 HEAD"
last_updated: "2026-07-23"
---

# State

**v2.7 "july jitter" — DESIGNED + IN DEVELOPMENT** (designed 2026-07-23 via `/developer-kit:design-roadmap`
from `.agentspace/annotation.md` field notes + 3 research workflows). The **re-ground + fidelity +
field-hardening** release (the v1.3b / v2.1 / v2.3 / v2.6 lineage): *realign the demo + corpus to the
platform's true current state, and fix what drifted.* **9 milestones M246 → M254**, tooling + docs only,
**0 platform-repo edits**. Branch `release/02.70-july-jitter` cut from `main`; tag will be `v2.7`. **M246 (the
HARD go/no-go re-sync barrier) is CLOSED — go/no-go PASS** (cold demo GREEN on the consolidated platform, 561
rows in `public.skill_path_sessions`, 3 subgraphs, 0 skillpath; the seeder re-point + demo clone-pin bump
landed; 9-row drift ledger → M247). The **M247–M252 fan-out** now branches from post-M246 HEAD.

> **The headline finding:** the skiller→app merge (v2.1) was **one step of a "consolidate every runtime engine
> into app" program.** `app` is **~386 commits** ahead of the stack pin; **skillpath is now FULLY decommissioned
> into `app`** (M501–M507: gone from `repos.yml`/compose/subgraphs → **3 subgraphs**, sessions →
> `public.skill_path_sessions`); **jobsimulation is mid-merge** (dormant, the next shoe); and `app` grew
> undocumented domains (coursebuilder, AI Labs + credits/stripe, askengine, a server-owned academy). The corpus
> asserts skillpath as live Tier-1 in ~30 files, and **rext `stack-seeding` writes to `skillpath.skill_path_sessions`
> in live seeder code** → breaks the instant a stack-update crosses M507 (the v2.1 class, repeating).

## v2.7 shape — barrier → 7-lane fan-out → prove-on-billion

```
M246 re-sync & re-point (HARD go/no-go barrier)
  ├─▶ M247 corpus re-ground ────────┐
  ├─▶ M248 content-mgr-link ────────┤
  ├─▶ M249 cross-app-nav ───────────┼──▶ M253 studio-first-paint (extends M249's studio patch ladder)
  ├─▶ M250 ai-readiness (iterative) ┤
  ├─▶ M251 test-health ─────────────┤
  ├─▶ M252 studio-builder-enable ───┤
  └────────────────────────────────────▶ M254 prove-on-billion (iterative closer)
```

- **M246** re-sync & re-point (`section`, HARD go/no-go) — re-point rext seeder `skillpath.*→public.*`, bump the
  **demo** clone pins to current `origin/main`, prove `/demo-up` green on the consolidated platform, emit a drift ledger.
- **M247** corpus re-ground (`section`) — skillpath→app redirect + **3-subgraph** truth + 4 new fact sheets
  (coursebuilder / AI Labs+credits / askengine / academy-backend) + refresh `ai-readiness.md`. Internal: core-lanes ∥ + reconcile-tail.
- **M248** content-stories manager result-link (`section`, small) — manager CTA → `/sim/<slug>/<userId>/result/<sessionId>` (the real built manager view).
- **M249** cross-app navigation (`section`) — "← Back to Cockpit" ×4 apps + studio logo/prod-eject fix. **Owns the first-ever studio-desk source patch machinery.**
- **M250** AI-readiness fidelity (**`iterative`**, marquee) — 31 canonical skills + 2 track-keyed named sims + evaluated-skills directus set-dress + skill distribution.
- **M251** test-health (`section`, the reserved v2.6→v2.7 carry) — `run-unit.sh` roster + ~6 mechanical `test_cockpit`/`test_public_host` re-points.
- **M252** studio-desk builder enablement (`section`) — AI-key **demo-wiring** (not a DNA/secrets gap) + DNA hardening + builder Playthrough. (talk-to-data Bedrock double-checked → COMPLETE.)
- **M253** studio-desk first-paint (**`iterative`**, deps M249) — shell-before-awaits + no-thirdparty demopatches, **<1s FCP** gate.
- **M254** prove-on-billion (**`iterative`**, closer) — re-prove the whole release live on `billion`, cold reset-to-seed, p95 < 5 s.

## Binding decisions (2026-07-23)
1. **Scope + codename** — expanded beyond the pre-reserved "test-health" to a full re-ground + fidelity release; codename **"july jitter"** (departs the stagecraft lineage, user's explicit choice).
2. **Re-ground depth** — **full bump to current `origin/main` + prove + author the 4 new fact sheets** (per "update repo to the new status quo").
3. **M250 shape → `iterative`** — the 8→31 arithmetic re-derivation + net-new directus-write set-dress + live-render believability make the path exploratory.
4. **DEF-M215-03(a)/F11 → DROPPED** — tripped its own drop-if-survives-another-release condition (v2.6 shipped without it).

## Parallel-build strategy (baked into each milestone's overview)
- **7-lane worktree fan-out** off M246: M247-core ∥ M248 ∥ M249→M253 ∥ M250 ∥ M251 ∥ M252. Run builds concurrently as `work-milestone --worktree=<path>`; **serialize the closes**.
- **All fan-out worktrees branch from post-M246 HEAD.** Merge order: M251 → {M248, M250} → M249 → M253 → M252 → M247-reconcile → M254.
- **9 coordination rules** (shared files: `cmd/stackseed/main.go` M248∥M250 · `run-unit.sh` roster M248→M251 · `CLAUDE.md` sole-owner M247 · `up-injected.sh build_frontend_studio_desk` M249→M253 · studio spec docs reconciled M247-tail · rung-zero every push).
- **Honest speedup:** ~3–4× on the build phase; ~1.5–2× end-to-end (the live proofs M246/M250/M253/M254 are the serial floor). Full detail in each `overview.md` + `roadmap.md § Active — v2.7`.
- **Environment (billion-last):** M246→M253 build + meet their gates on a **LOCAL `demo-N`** (this box); **`billion` stays untouched** (its v2.6 demo left up) until **M254**, the closer. Caveat: M253's <1 s FCP is tailnet-sensitive → its local gate is provisional, re-confirmed on billion at M254.

## Headline numbers (inherited from v2.6 close, 2026-07-23 — reset at v2.7 close)
- **Go:** **2010** reproducible platform test funcs. 2461 testcases / 0 failed, 6 modules.
- **TypeScript (unit):** **257** `*.unit.spec.ts` + 40 live-browser specs (24 stack-verify + 16 Playthroughs).
- **Python (rext demo-stack):** **839 pass / 8 standing fail** (host-sensitive; 0 real defects — the M251 target).
- **content-stories:** 47/47 landed of the 49-pair denominator. **p95 click→ACCESS:** employee 1.46 s · manager 1.31 s.
- **Flake: 0.** **Alignment (Clerkenstein): 100% / 100% critical.** **Supply chain: GREEN.** **Platform-repo edits: 0.**
- rext code-of-record @ v2.6: `498b1a5` (tag `sound-check-m244-content-sweep-robustness`).

## Standing backlog carried INTO v2.7 (fated destinations)
- **→ M251 (NAMED milestone):** the 8 standing demo-stack test failures (6 mechanical `test_cockpit`/overlay + `test_public_host` port-13001; ~2 docker/live-gated `test_purge` ride M254) + the `run-unit.sh` `UNIT_SPECS` roster nit (2 unrostered M244-harden specs → runner exit 2).
- **DROPPED at v2.7 design:** **DEF-M215-03(a)/F11** (cosmetic hero identity-key vs display-name) — tripped its drop-if-survives condition; **DEF-M239-01** (ENOSPC loud-build) — dropped at v2.6 close (disk-full class already covered).
- **Still unscheduled (vision):** DEF-M10-01 (S3 media — the document facet was consumed by v2.6 M240; voice presence-only) · DEF-M21-01 · CAVEAT-1 · M314b (platform) · **M205** residual (tier gates + ATS) · Playthroughs futures **M206–M207**.

## Process flags (do NOT auto-push)
- **v2.5's** `release→main` merge + `v2.5` tag are **LOCAL-ONLY**, not pushed to origin (R5).
- **A stray `(M245)` commit** sits on `main` (post-v2.6 academy docs, untracked in the plan) — v2.7 numbering starts at **M246** to skip it.
- The user runs the v2.5/v2.6/v2.7 origin publishes on their own cadence.

_Last updated: 2026-07-23 — M246 (the HARD go/no-go re-sync barrier) CLOSED + merged to release/02.70-july-jitter (go/no-go PASS); next is the M247–M252 fan-out off post-M246 HEAD._
