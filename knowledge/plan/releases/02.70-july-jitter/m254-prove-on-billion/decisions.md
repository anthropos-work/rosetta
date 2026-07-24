# M254 — Decisions

_(Implementation decisions with rationale, D-numbered, recorded during build. TOK entries — the
strategy-evolution chain — live here at the milestone root; intra-iter decisions live in each `iter-NN/decisions.md`.)_

## TOK-01: cluster-per-tik live re-prove (DRIVE → read-only sweeps → mutating tail) — 2026-07-24

**Tok type:** bootstrap (iter-01)
**Initial strategy:** Discharge the multi-part a–h exit gate **one gate-cluster per tik**, following the
roadmap's intra-milestone LANE decomposition, measure→confirm→fix-forward until every part is GREEN cold on
`billion`. The clusters, in dependency order:
1. **The DRIVE (gate a) — single-driver serial, the critical path.** Cold reset-to-seed `/demo-up` on billion
   at pin `july-jitter-m253-studio-first-paint` (`up-injected.sh 1 --public-host billion.taildc510.ts.net`,
   run foreground-blocking inside a tracked background Bash — never detach-and-yield on billion, never kill a
   mid-build). Assert: builds + comes up GREEN on the consolidated platform (**3 subgraphs, skillpath-in-app**),
   health 200 + casbin > 0, **fresh green `autoverify.json`**. Every downstream proof gates on this.
2. **Read-only confirmation sweeps (gates b, c, d, g + part of h) — fan out ~3 concurrent tailnet-PEER sweeps
   against ONE bring-up.** content-stories sweep (b: manager CTA lands on `/sim` per-session manager view,
   non-empty; part of h) ∥ coverage sweep (c: Back-to-Cockpit in all 4 apps + studio logo/back/logout → stack
   app, 0 prod-ejects; d: AI-readiness faithful per M250's gate both vantages incl. the 3 manager-dashboard
   drift-fix sections; part of h) ∥ probes (g: the **8** live/docker-gated demo-stack test-health tests).
   Asserted from THIS workstation (a tailnet peer), never from the VM.
3. **Latency solo (gates f + part of h) — quiet system, no concurrent load.** studio first-paint < 1 s **cold
   p95** (5 cold loads, fresh-green autoverify); p95 click→ACCESS < 5 s hero vantages. `LATENCY_SCHEME=https`
   mandatory; the latency runner needs `STACK_DIR`; gate on fresh-green `autoverify.json`, never networkidle.
4. **Mutating / seed-destroying serial tail (gate e + rest of h) — after the read-only sweeps.** studio
   builder Playthrough green (the ~10-min async generate — assert the completion BOUNDARY); the live-browser
   specs + Playthroughs green. These mutate/re-seed so they cannot share the read-only bring-up.
Plus the M247-reconcile tail (CLAUDE.md/README "16→18" playthrough-count mirror) if it surfaces.

**Rationale:** the gate's own LANE decomposition (roadmap M254) dictates this order — the DRIVE is the
un-shardable critical path that every other proof depends on (fresh green autoverify precondition); the
read-only sweeps parallelize on ONE bring-up; latency must be solo or the p95 is polluted; the mutating tail
can't share a read-only bring-up. Every defect routes to rext / a sha-pinned demopatch (0 platform edits),
committed + tagged + pushed to origin (rung-zero) before re-pinning billion. This is the proven M221/M236/M244
shape.

**Strategy class:** new-direction (bootstrap — no prior strategy to compare against).
**Distance-to-gate context:** gate = 0/8 parts confirmed live at bootstrap (all carry forward from
M246–M253's local-provisional gates); the metric is the count of a–h parts GREEN cold on billion. The gate
reads binary-per-part (0–8) and can look FLAT across productive within-cluster iters — drive by real per-part
evidence, not the coarse counter (M244 lesson; a benign triggered tok ~1 per ~5 iters is expected).
**Next-tik direction:** iter-02 = the DRIVE. Kick the cold bring-up on billion at the pin; assert gate (a).

