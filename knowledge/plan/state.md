---
active_release: "v2.7 «july jitter» — SHIPPED 2026-07-25 (tag v2.7; release→main merged). The re-ground + fidelity + field-hardening release: realign the demo + corpus to the platform's TRUE current state (skillpath fully decommissioned into app → 3 subgraphs; jobsim mid-merge; net-new app domains; the rext seeder re-pointed off skillpath.*) and fix six field defects (content-stories manager link · cross-app Back-to-Cockpit · studio prod-eject · AI-readiness fidelity · studio builder keys · studio blank-page). 9 milestones M246→M254, barrier → 7-lane fan-out → prove-on-billion; tooling + docs only, 0 platform-repo edits. Close-release landed all 5 terminal follow-ups (zero carry-forward)."
active_branch: "main (release/02.70-july-jitter merged + deleted at close-release)"
active_milestone: "(between releases)"
last_closed: "M254 — 2026-07-25 (prove-on-billion; iterative, closed-on-gate); v2.7 release closed 2026-07-25"
phase: "between releases — awaiting /developer-kit:design-roadmap (v2.8 not yet designed)"
last_updated: "2026-07-25"
---

# State

**Between releases.** v2.7 "july jitter" **SHIPPED 2026-07-25** (tag `v2.7`; `release/02.70-july-jitter`
merged to `main` + deleted). The **re-ground + fidelity + field-hardening** release (the v1.3b / v2.1 / v2.3 /
v2.6 lineage): *realign the demo + corpus to the platform's true current state, and fix what drifted.* **9
milestones M246 → M254**, tooling + docs only, **0 platform-repo edits**. No milestone is active; the next step
is **`/developer-kit:design-roadmap`** to design **v2.8** (NOT yet designed — do not treat it as active).

> **The headline finding (release thesis):** the skiller→app merge (v2.1) was **one step of a "consolidate every
> runtime engine into app" program.** `app` is **~386 commits** ahead of the old stack pin; **skillpath is now
> FULLY decommissioned into `app`** (M501–M507: gone from `repos.yml`/compose/subgraphs → **3 subgraphs**,
> sessions → `public.skill_path_sessions`); **jobsimulation is mid-merge** (dormant, the next shoe); and `app`
> grew undocumented domains (coursebuilder, AI Labs + credits/stripe, askengine, a server-owned academy). v2.7
> re-grounded the corpus + demo to this topology and re-pointed the rext seeder off `skillpath.*` before it broke.

## v2.7 shape — barrier → 7-lane fan-out → prove-on-billion  (all 9 SHIPPED)

```
M246 re-sync & re-point (HARD go/no-go barrier)  ✅
  ├─▶ M247 corpus re-ground ─────────┐ ✅
  ├─▶ M248 content-mgr-link ─────────┤ ✅
  ├─▶ M249 cross-app-nav ────────────┼──▶ M253 studio-first-paint ✅
  ├─▶ M250 ai-readiness (iterative) ─┤ ✅
  ├─▶ M251 test-health ──────────────┤ ✅
  ├─▶ M252 studio-builder-enable ────┤ ✅
  └─────────────────────────────────────▶ M254 prove-on-billion (iterative closer) ✅
```

Full per-milestone closure narratives (decisions, deliverables, test deltas, code-of-record tags) live in
`roadmap.md`'s `## Done — v2.7` `### M{N}` blocks; the release review + retro + metrics live in
`releases/archive/02.70-july-jitter/`.

## Headline numbers (v2.7 close-release, 2026-07-25 — triple-clean PASS)
- **Go:** **2019** rext test funcs (`git grep -E '^func Test'`, +9 vs v2.6's 2010; tooling+docs release). Runtime: stack-seeding 1192 pass / playthroughs 131 pass / 0 fail; `gofmt`/`go vet`/`go build` clean.
- **TypeScript (unit):** **292** executed / 0 fail (stack-verify/e2e 178 + playthroughs/e2e 114); `tsc` clean. + 40 live-browser specs (proven at M254; need a running stack).
- **Python (rext):** **demo-stack 910 pass / 0 fail / 1 skip** (the v2.6 8-fail `g-testhealth` carry ELIMINATED at close); stack-injection 258 / 9 skip; stack-verify 149 pass / 5 environmental-fail (live local stack, academy down — byte-identical ×3, not regressions).
- **Live on billion (M254):** Playthroughs **18/18**; content-stories **45/45 landable + 4 voice presence-only** (denominator 47→45 formalized at close); p95 click→ACCESS **1.43 s emp / 1.41 s mgr** (<5 s); studio first-paint **p50 637–726 ms** (<1 s); AI-readiness both vantages; Back-to-Cockpit 4/4, 0 prod-ejects.
- **Flake: 0** (triple-clean byte-identical ×3). **Supply chain: GREEN** (0 net-new deps). **Platform-repo edits: 0.**

## Recently shipped releases (max 3; older → roadmap.md / roadmap-legacy.md)
- **v2.7 "july jitter" — 2026-07-25** (tag `v2.7`) — re-ground + fidelity + field-hardening; 9 milestones M246→M254; prove-on-billion a–h live; **zero carry-forward** (all 5 terminal follow-ups landed at close); 0 platform edits.
- **v2.6 "sound check" — 2026-07-23** (tag `v2.6`) — reliability / field-hardening; 8 milestones M237→M244; prove-on-billion 8/8 live.
- **v2.5 "the playbill" — 2026-07-20** (tag `v2.5`) — content-vantage: empty-academy fill + the content-stories cockpit tab; 8 milestones M229→M236.

## Standing backlog (fated destinations)
- **v2.7 close-release LANDED (zero carry-forward):** all 5 terminal follow-ups the user elected to land — `FIX-M254-g-testhealth` (carry eliminated, demo-stack 8-fail → 910 pass/1 skip), `FIX-M254-c-academy-durable` (idempotent reapply + logging), `(b)-voice manager_presence_only` (denominator 47→45 formalized), `verify.sh` skillpath default drop, rext-hygiene inert set. rext close tags on origin: `july-jitter-v27-close-rext-cleanup @ e61b604` + `july-jitter-v27-close-followups @ a5b1288`.
- **DROPPED:** DEF-M250-01 `participants_filter` (D18) · DEF-M215-03(a)/F11 (design-time) · DEF-M239-01 (v2.6).
- **Still unscheduled (vision):** DEF-M10-01 (S3/Bunny voice media — voice presence-only) · DEF-M21-01 · CAVEAT-1 · M314b (platform) · **M205** residual (tier gates + ATS) · Playthroughs futures **M206–M207**.

## Process flags (do NOT auto-push)
- **v2.7 is merged to `main` + tagged `v2.7` LOCALLY at close-release; NOT pushed to origin** — the user runs origin publishes on their own cadence.
- **v2.5's** `release→main` merge + `v2.5` tag are **LOCAL-ONLY**, not pushed to origin (R5). **v2.6** likewise merged to main locally (commit `5cf51ed`); confirm the `v2.6` tag + origin push on the user's cadence.
- **A stray `(M245)` commit** sits on `main` (post-v2.6 academy docs, untracked in the plan) — v2.7 numbering started at **M246** to skip it.
- rext code-of-record: all v2.7 `july-jitter-*` tags are on **origin** (rung-zero verified); the rext authoring copy is at `july-jitter-v27-close-followups @ a5b1288`.

_Last updated: 2026-07-25 — v2.7 "july jitter" SHIPPED (close-release: release→main merged, tag v2.7, all 5 terminal follow-ups landed [zero carry-forward], triple-clean PASS, 0 platform edits). Awaiting `/developer-kit:design-roadmap` for v2.8._
