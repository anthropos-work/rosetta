# Release Retro: v2.7 "july jitter"

**Shipped:** 2026-07-25 · **Milestones:** M246 → M254 (9) · **Shape:** barrier → 7-lane fan-out → prove-on-billion · **0 platform-repo edits.**

## What this release was
The **re-ground + fidelity + field-hardening** release. The headline finding drove it: the skiller→app merge (v2.1) was **one step of a "consolidate every runtime engine into app" program** — `app` is ~386 commits ahead of the old stack pin, **skillpath is now fully decommissioned into `app`** (→ 3 subgraphs, sessions in `public.skill_path_sessions`), jobsim is mid-merge, and `app` grew four undocumented domains. v2.7 re-grounded the corpus + demo to that topology (re-pointing the rext seeder off `skillpath.*` before it broke) and fixed six field defects (content-stories manager link · cross-app Back-to-Cockpit · studio prod-eject · AI-readiness fidelity · studio builder keys · studio blank-page). Every fix was proven **live on `billion`** at M254.

## P1 incidents

1. **The 7.5-hour orchestration stall (P1, process).** During M254 prove-on-billion, a build-iters sub-agent backgrounded a cold reset-to-seed and yielded; the background op completed but did **not** re-invoke the agent, and the coordinator had stood down its staleness watchdog → the milestone sat dormant overnight. **Root cause:** in this environment a `run_in_background`+yield wait does not reliably re-invoke a sub-agent (twice, even "hardened with a marker"). **Fix:** the FOREGROUND bounded-poll pattern (detached op + sentinel + a foreground `Bash timeout` polling loop that blocks-and-returns, no yield) + **never** standing down the coordinator watchdog. Recorded in agent memory so the next prove-on-billion doesn't repeat it. No release-artifact impact — the work was intact, only wall-clock was lost.

2. **Self-inflicted container mis-kill, cleanly recovered (P1, ops).** M254 iter-08 tried to hot-re-heal the native-academy patch on a running demo; a `fuser -k` killed `tailscaled` (systemd auto-restarted it) and the host-visible web+hiring container processes lost their port mappings. **Recovered** via a cold reset-to-seed re-bring-up (which became net-positive — the fresh demo re-proved gates a + c). **Lessons recorded:** never `fuser -k` a tailscale-serve-fronted port; use `docker` verbs on demo containers; render a native-academy patch via a fresh launch, not a hot re-apply.

## Cross-milestone patterns

1. **Mirrored-count drift (the dominant review finding).** Three independent count-drifts surfaced across the release: the demopatch inventory **shipped RED for a full milestone span** (M253 bumped the doc header 21→23 but not the `test_patch_inventory.py` fence — caught by M254 harden, landed FIX-M254-h), the Playthrough count (16→18 lagged in `CLAUDE.md`/`README.md`), and the AI-readiness KPI tiles (4→5). **Instructive case:** the demopatch count *had* a fence and still shipped RED because the fence wasn't bumped in the same commit as the doc. **Discipline captured** (release-review C1 + `demopatch-spec.md §5`): any count mirrored in >1 doc or backed by a fence must move with all mirrors + its fence in one commit, and the mirror set must be enumerated where the authority lives.

2. **Blind-authored tooling drifts against a redesigned surface.** M252's studio-builder Playthrough was authored *without a live render in hand* ("orchestrator live-tunes later") and never live-tuned; by M254 the studio-desk on billion was a redesign (`v0.152.1`, unified `/simulation-builder` entry) and the Playthrough tested dead routes. Live-proof caught it; re-tuned to the real UX. **Lesson:** a tooling artifact authored blind against a UI must be live-tuned before it's trusted, or it silently rots.

3. **Live-proof earns its keep.** Prove-on-billion caught what unit tests could not: a demopatch anchor that moved in the platform consolidation, the studio redesign drift, and the studio `--public-host` "session-carry defect" that turned out to be a measurement artifact (wrong test hero). The pattern from v2.5 M236 / v2.6 M244 held.

## Carry-forwards
**Zero.** At close-release the user elected to LAND all five terminal-milestone follow-ups rather than defer them:
- `FIX-M254-g-testhealth` — the host-sensitive demo-stack test-harness carry **eliminated** (8 fail → 910 pass / 1 documented Linux-only skip).
- `FIX-M254-c-academy-durable` — reapply lifecycle made idempotent + logged.
- `(b)-voice manager_presence_only` — formalized (denominator 47→45, 4 voice cells presence-only, honesty gate green).
- `verify.sh` stale-`skillpath` default + the rext-hygiene inert set — removed.

## Metrics delta (vs v2.6)
Go test funcs 2010→**2019**; TS unit 257→**292**; Python demo-stack **910 pass / 0 fail / 1 skip** (the 8-fail carry cleared); flake **0**; triple-clean **PASS** (byte-identical ×3); supply-chain **GREEN** (0 net-new deps); **0 platform-repo edits**. Live gates: Playthroughs 18/18, p95 click→ACCESS 1.43/1.41 s, studio p50 <1 s, content-stories 45/45 + 4 voice presence-only, AI-readiness both vantages. Full table: `knowledge/plan/metrics-history.md`.

## Stats delta
Release-close snapshot: `knowledge/journal/stats/2026-07-25.json` (5 total snapshots; +193 commits vs the 2026-07-20 snapshot; 1210 total commits, 25/25 milestones). Note: project-stats reads 0 for CODE/TEST/DOCS because rosetta is docs-only and the code-of-record lives in the rext repo, which project-stats doesn't traverse from the rosetta root — the authoritative test metrics are in metrics-history.md.

## What went well
- The barrier → fan-out shape worked: M246's HARD go/no-go correctly gated the fan-out; the two demo-free lanes (M247 docs ∥ M251 test-health) parallelized via worktrees.
- The re-ground itself was clean — the hard structural facts (3 subgraphs, skillpath-in-app, 4 new fact sheets, 27 service docs) landed correctly across ~30 files; the residue was only secondary count/reference lag.
- Zero platform-repo edits held across all 9 milestones; every platform-side need routed to a sha-pinned demopatch or a tooling fix.
