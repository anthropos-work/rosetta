# M254 — Retro (prove on billion — the closer)

## Summary
The terminal live re-prove of **v2.7 "july jitter"**: the whole release proven together on the `billion` Tailscale
VM, **cold reset-to-seed, driven + asserted from a tailnet peer, 0 platform edits** (the M221/M236/M244
prove-on-billion lineage). **Closed `closed-on-gate` 2026-07-25** — all 8 exit-gate parts (a–h) MET live:
(a) autoverify `green:true`/0 warnings on the consolidated platform (3 subgraphs, skillpath-in-app);
(b) content-stories **45/45 landable ALL LANDED** + 4 voice presence-only; (c) **4/4 apps** render "← Back to
Cockpit" + studio prod-eject side, 0 escapes; (d) AI-readiness faithful **both vantages** (8/8 sections + 3
drift-fixes); (e) **both studio builders green**; (f) studio first-paint **p50 637–726 ms < 1 s**; (g) 2/8-class
test-health fixed live; (h) **latency p95 1.43 s employee / 1.41 s manager < 5 s** + **Playthroughs 18/18**. 10 iters
(1 bootstrap tok + 9 tiks) + 1 final harden pass (stabilized, +22 tests). Code-of-record: 6 rext tags on origin
(rung-zero), billion demo pinned `july-jitter-m254-academy-nonode-hostrobust @ dfdd9bc`. Deferral audit **YELLOW**
(0 blockers). **0 platform-repo edits.**

## Incidents This Cycle
- **P1 (self-inflicted, cleanly recovered — iter-08).** A (c)-academy re-heal excursion cascaded into a demo
  disruption: re-applying the back-to-cockpit patch to the RUNNING native academy crashed its next-dev; relaunch hit
  `EADDRINUSE`; **`sudo fuser -k 13077/tcp` killed tailscaled** (SSH dropped, systemd auto-restarted it) and then
  mis-killed the web+hiring **container** `next-server` procs (visible on the host PID namespace), which lost their
  host port mappings on `docker start`. **Recovery:** a full cold reset-to-seed re-bring-up at the latest rext pin —
  **net-positive** (a fresh demo that re-confirmed gates a + c). **Lessons (now in the discipline):** never `fuser -k`
  a tailscale-serve-fronted port; a container's `next-server` proc is visible on the host PID ns — don't kill by name.
- **P2 (caught pre-ship, iter-02) — the AI-readiness demopatch had drifted into a silent no-op.** The
  `app-aireadiness-snapshot-loadmembers` patch still targeted the pre-consolidation `internal/workforce/ai_readiness.go`;
  the app consolidation moved it to `internal/aireadiness/readiness.go`, so the patch silently skipped → autoverify
  `green:false` (gate a). Re-authored (path/anchor/shas) at iter-03; regression-fenced in the harden pass.
- **P2 (caught at iter-10) — the M252 studio-builder Playthrough targeted stale routes.** studio-desk on billion is
  **v0.152.1 (2026-07-03, a redesign predating M252)** — unified `/simulation-builder`; the authored-blind
  never-live-tuned Playthrough hit dead routes. Re-tuned to the real UX; both builders green.
- **P3 — a measurement artifact masqueraded as a defect (gate f, iter-06).** The "studio session-carry" symptom was
  the FCP probe defaulting to `maya-thriving` (an employee whom studio's `checkEnterpriseAndAdmin` correctly bounces);
  admin heroes reach the shell. Fixed the default identity (maya → dan-manager). Not a real defect.
- **P3 (surfaced by the final harden pass) — a RED-at-HEAD sibling-milestone regression.** `test_patch_inventory` fence
  asserted 21/studio-desk-3 while `patches/` held 23/5 — **root cause M253** (`b8969c0` bumped the doc header but not
  the directory-driven fence). Re-fated Fate-3 → **Fate-1 LANDED at close** (fence 21→23 + `demopatch-spec.md §5`
  reconciliation).

## What Went Well
- **The gate held its shape.** 1 bootstrap tok, 0 triggered toks — no 3-consecutive-no-progress stall across 9 tiks.
  Every gate part discharged by real per-part live evidence, not the coarse counter (the M244 lesson).
- **Measure-first repeatedly beat the hypothesis.** Gate (f)'s "session-carry defect" was a wrong-hero measurement
  artifact; gate (e)'s failure was stale-Playthrough drift vs a live redesign, not a broken builder. Reading the live
  render before "fixing" saved two speculative fixes.
- **Rung-zero discipline held on every push.** All 6 rext fix tags verified on origin before billion consumed them;
  no repeat of the M236 unpushed-tag class.
- **Self-inflicted disruption recovered without a platform edit or data loss** — the cold reset-to-seed lifecycle IS
  the recovery path; it left a fresh demo that advanced the gate.
- **A RED-at-HEAD did not ship.** The close re-fated the harden-surfaced fence drift to Fate-1 and landed it in full
  (rext + corpus reconciled together, per the doc's own "bump both" contract).

## What Didn't
- **The (c)-academy re-heal excursion was avoidable** — the durability edge (a long-running native academy reverting
  the patch) is real but the presenter path (a fresh demo) renders fine; the excursion to "fix it live" caused the P1.
  Routed to `FIX-M254-c-academy-durable` (make the reapply lifecycle durable) instead of live-poking a running demo.
- **The demo-stack test-health membership is chronically host-sensitive.** 6 tests (test_ant_academy launcher/reap +
  clerk-wiring overlay + Linux-only test_purge + docker image-guard) are test-harness environment-sensitivity, not
  demo defects (real academy serves 200) — but they've now ridden two releases as a live-gated carry. Routed to
  `FIX-M254-g-testhealth` for a proper per-test disentangle.
- **The M253 header-only inventory bump left a latent RED for a full milestone.** The `demopatch-spec.md §5` contract
  says "bump the table AND the fence together" — M253 bumped only the header, and three secondary refs (144/218-219/238)
  plus the rext fence stayed stale until this close caught them. The contract is right; the execution slipped.

## Carried Forward
- **`FIX-M254-c-academy-durable` (Fate 3)** — make the native-academy Back-to-Cockpit patch durable across the
  next-dev lifecycle (rext `stack-injection`/`ant-academy.sh` reapply). 0 platform edits.
- **`FIX-M254-g-testhealth` (Fate 3)** — per-test disentangle + fix of the 6 host-sensitive demo-stack tests in rext
  `demo-stack`. 0 platform edits, 0 demo-runtime impact.
- **`(b)-voice manager_presence_only` (Fate 3)** — content denominator 47→45 + a `manager_presence_only` flag + re-seed
  (the surface renders; no fabricated CTA). The underlying voice-media blocker is the accepted vision item DEF-M10-01.
- **verify.sh stale `skillpath` default (Fate 3)** — drop the decommissioned `skillpath` from verify.sh's default
  service list (bring-up already passes explicit `--services`).
- **studio-desk billion re-pin → `july-jitter-m254-studio-pt-retune` (Fate 3, note)** — on the next full cold re-prove,
  pin billion's rext past the test-only pt-retune + fence commits so the whole toolchain is current.

## Metrics Delta
- **Live on billion (cold reset-to-seed, tailnet peer):** Playthroughs **18/18** (100%); latency p95 click→ACCESS
  **1.43 s employee / 1.41 s manager** (< 5 s); studio first-paint **p50 637–726 ms < 1 s** (p95 = accepted
  environmental tailnet jitter); content-stories **45/45 landable ALL LANDED** + 4 voice presence-only; AI-readiness
  both vantages (8/8 + 3 drift-fixes); 4/4 Back-to-Cockpit, 0 prod-ejects.
- **Test suites (this Mac):** demo-stack Python **909 collected / 900 pass** (+40 collected vs M251's 869; the 8+1
  remaining = the g-testhealth host-sensitive carry, 0 milestone regressions); TypeScript unit +14 (harden);
  `test_patch_inventory` 2 RED → 5/5 green (close fix); playthroughs `tsc` clean.
- **Flake: 0.** **Platform-repo edits: 0.** rext code-of-record: `july-jitter-m254-patch-inventory-fence @ 02ac973`
  (all 6 M254 tags on origin, rung-zero).
