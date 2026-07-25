---
active_release: "v2.7 «july jitter» — FULLY BUILT (designed 2026-07-23; all 9 milestones M246→M254 closed 2026-07-23..25). The re-ground + fidelity + field-hardening release: realign the demo + corpus to the platform's TRUE current state (skillpath fully decommissioned into app → 3 subgraphs; jobsim mid-merge; net-new app domains; the rext seeder re-pointed off skillpath.*) and fix six field defects (content-stories manager link · cross-app Back-to-Cockpit · studio prod-eject · AI-readiness fidelity · studio builder keys · studio blank-page). Barrier → 7-lane fan-out → prove-on-billion. Branch release/02.70-july-jitter; **awaiting the user's /developer-kit:close-release** (release→main merge + tag v2.7)."
active_branch: "release/02.70-july-jitter (cut from main 2026-07-23) — all 9 milestones merged in"
active_milestone: "(between milestones — v2.7 fully built; awaiting /developer-kit:close-release)"
last_closed: "M254 — 2026-07-25 (prove-on-billion; iterative, closed-on-gate)"
phase: "v2.7 fully built — all 9 milestones M246→M254 closed + merged to release/02.70-july-jitter; awaiting the user's /developer-kit:close-release (release→main merge + v2.7 tag). Do NOT auto-merge or auto-tag."
last_updated: "2026-07-25"
---

# State

**v2.7 "july jitter" — FULLY BUILT, awaiting close-release.** All **9 milestones M246 → M254** are closed and
merged into `release/02.70-july-jitter` (tooling + docs only, **0 platform-repo edits**). The
**re-ground + fidelity + field-hardening** release (the v1.3b / v2.1 / v2.3 / v2.6 lineage): *realign the demo +
corpus to the platform's true current state, and fix what drifted.* The terminal milestone **M254
(prove-on-billion, `iterative`) closed `closed-on-gate` 2026-07-25** — the whole release re-proven live on the
`billion` Tailscale VM, cold reset-to-seed, driven + asserted from a tailnet peer: all 8 exit-gate parts (a–h)
MET (Playthroughs 18/18, latency p95 1.43/1.41 s < 5 s, studio first-paint p50 < 1 s, content-stories 45/45
landable, AI-readiness both vantages, 4/4 Back-to-Cockpit, 0 platform edits). **Next step is the user's**
`/developer-kit:close-release` (the release-level review + `release→main` merge + `v2.7` tag).

> **The headline finding (release thesis):** the skiller→app merge (v2.1) was **one step of a "consolidate every
> runtime engine into app" program.** `app` is **~386 commits** ahead of the old stack pin; **skillpath is now
> FULLY decommissioned into `app`** (M501–M507: gone from `repos.yml`/compose/subgraphs → **3 subgraphs**,
> sessions → `public.skill_path_sessions`); **jobsimulation is mid-merge** (dormant, the next shoe); and `app`
> grew undocumented domains (coursebuilder, AI Labs + credits/stripe, askengine, a server-owned academy). v2.7
> re-grounded the corpus + demo to this topology and re-pointed the rext seeder off `skillpath.*` before it broke.

## v2.7 shape — barrier → 7-lane fan-out → prove-on-billion  (all 9 done)

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
`roadmap.md`'s `### M{N}` blocks — NOT here (state.md contract).

## Binding decisions (2026-07-23)
1. **Scope + codename** — expanded beyond the pre-reserved "test-health" to a full re-ground + fidelity release; codename **"july jitter"** (departs the stagecraft lineage, user's explicit choice).
2. **Re-ground depth** — **full bump to current `origin/main` + prove + author the 4 new fact sheets**.
3. **M250 shape → `iterative`** — the 8→31 re-derivation + net-new directus-write set-dress + live-render believability made the path exploratory.
4. **DEF-M215-03(a)/F11 → DROPPED** — tripped its own drop-if-survives-another-release condition.

## Headline numbers (v2.7 close-of-M254, 2026-07-25 — reset at the v2.7 close-release)
- **Go:** **2010** reproducible platform test funcs (unchanged; v2.7 is tooling + docs only).
- **TypeScript (unit):** **271** `*.unit.spec.ts` (257 at v2.6 + 14 from the M254 harden) + 40 live-browser specs (24 stack-verify + 16 Playthroughs).
- **Python (rext demo-stack):** **909 collected / 900 pass** — the 8 fail + 1 error are the `FIX-M254-g-testhealth` host-sensitive test-harness carry (test_ant_academy launcher/reap + clerk-wiring overlay + Linux-only test_purge + docker image-guard); **0 milestone regressions, 0 demo-runtime impact** (real academy serves 200).
- **Playthroughs:** **18/18** live on billion. **content-stories:** 45/45 landable ALL LANDED + 4 voice presence-only (Bunny-keyless demo box, `DEF-M240-01`). **p95 click→ACCESS (billion):** employee 1.43 s · manager 1.41 s. **studio first-paint (billion):** p50 637–726 ms < 1 s.
- **Flake: 0.** **Alignment (Clerkenstein): 100% / 100% critical.** **Supply chain: GREEN.** **Platform-repo edits: 0.**
- rext code-of-record (M254): tag `july-jitter-m254-patch-inventory-fence @ 02ac973` (all 6 M254 tags on origin, rung-zero).

## Recently closed (max 5; older → roadmap.md ### M{N} blocks)
- **M254 — 2026-07-25** (prove-on-billion, `iterative`, closed-on-gate) — v2.7 re-proven live on billion, a–h MET (Playthroughs 18/18, p95 1.43/1.41 s, studio p50 <1 s, content-stories 45/45, AI-readiness both vantages, 0 platform edits); close landed FIX-M254-h (demopatch inventory fence 21→23); audit YELLOW.
- **M253 — 2026-07-24** (studio-desk first-paint, `iterative`, closed-on-gate) — first-meaningful-paint 4669 → p95 817 ms (demo-2 laptop, ~5.7× under <1 s); 2 demopatches on the M249 ladder; rext `july-jitter-m253-studio-first-paint @ b8969c0`.
- **M252 — 2026-07-24** (studio-builder enablement) — demo studio carries its own AI keys via `env_file`-only wire; studio-desk's first Playthroughs entry (16→18); rext `july-jitter-m252-studio-builder @ d80db9f`.
- **M250 — 2026-07-24** (AI-readiness fidelity, `iterative`) — real 31-skill default + track-keyed sims + directus set-dress; live-green both vantages; rext `july-jitter-m250-iter07 @ 584f1fe`.
- **M249 — 2026-07-24** (cross-app navigation) — "← Back to Cockpit" ×4 apps + studio prod-eject fix; first-ever studio-desk source patch trio; rext `july-jitter-m249-harden @ 8ab5192`.

## Standing backlog carried INTO / OUT OF v2.7 (fated destinations)
- **v2.7-originated follow-ups (rext tooling, non-blocking, 0 platform edits):** `FIX-M254-c-academy-durable` (make the native-academy Back-to-Cockpit patch durable across the dev-server lifecycle) · `FIX-M254-g-testhealth` (6 host-sensitive demo-stack test-harness tests) · `(b)-voice manager_presence_only` (content denominator 47→45 + flag + re-seed) · verify.sh stale `skillpath` default drop · studio-desk billion re-pin → `july-jitter-m254-studio-pt-retune` on the next full cold re-prove. All Fate-3, coordinator-approved (see M254 carry-forward.md).
- **DROPPED at v2.7 design:** **DEF-M215-03(a)/F11** (cosmetic hero identity-key) — tripped its drop-if-survives condition; **DEF-M239-01** (ENOSPC loud-build) — dropped at v2.6 close.
- **Still unscheduled (vision):** DEF-M10-01 (S3/Bunny voice media — voice presence-only; document facet consumed by v2.6 M240) · DEF-M21-01 · CAVEAT-1 · M314b (platform) · **M205** residual (tier gates + ATS) · Playthroughs futures **M206–M207**.

## Process flags (do NOT auto-push)
- **v2.7 is FULLY BUILT but NOT closed as a release.** `/developer-kit:close-milestone` merged M254 → `release/02.70-july-jitter`; it did **NOT** merge `release → main` or tag `v2.7` — that is the user's separate `/developer-kit:close-release` step.
- **v2.5's** `release→main` merge + `v2.5` tag are **LOCAL-ONLY**, not pushed to origin (R5).
- **A stray `(M245)` commit** sits on `main` (post-v2.6 academy docs, untracked in the plan) — v2.7 numbering started at **M246** to skip it.
- The user runs the v2.5 / v2.6 / v2.7 origin publishes on their own cadence.

_Last updated: 2026-07-25 — M254 (prove-on-billion, `iterative`) CLOSED-on-gate + merged to release/02.70-july-jitter (a–h MET live on billion; Playthroughs 18/18, p95 1.43/1.41 s, studio p50 <1 s; close landed FIX-M254-h fence 21→23, audit YELLOW; 0 platform edits). **v2.7 fully built — all 9 milestones M246→M254 closed; awaiting the user's /developer-kit:close-release.**_
