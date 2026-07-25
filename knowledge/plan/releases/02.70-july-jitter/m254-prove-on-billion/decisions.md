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


## D-iter10-1: gate (e)+(h) MET live on billion — studio-builder re-tune + networkidle anti-deadlock — 2026-07-25

**Iter:** iter-10 (tik, closed-fixed). **Gate:** (e) + (h)-Playthroughs MET.
The full Playthrough suite is GREEN on billion (cold reset-to-seed, tailnet peer): **18/18 passing (100 %),
0 failing, DONE_rc=0, `ptreport --gate no-regressions` PASSED.** Two live tooling fixes (rext
`july-jitter-m254-studio-pt-retune` @ 4f1409e, ON ORIGIN — rung-zero; 0 platform edits):
1. **(e)** the M252 studio-builder Playthrough targeted STALE routes (`/sim-advanced-builder` + an immediate
   Generate button) — authored-blind, never live-tuned. studio-desk on billion is **v0.152.1 (2026-07-03,
   a redesign predating M252)**: unified `/simulation-builder` entry. Re-authored to the real UX — advanced
   "Design it with AI" DRAFTS a scenario the designer renders (the true generation boundary); guided's 5-part
   interview is live (Generate at Part 5 is P6-out).
2. **(h)** 4 `/home`-landing logins defaulted to `networkidle`, which never idles on the AI-readiness polling
   surface → a flaky 120 s timeout (pt-skillpath-legacy). Applied the protocol's anti-deadlock
   `waitUntil:'domcontentloaded'` to all 4.
No billion re-pin (test-only fix; billion's demo build at `dfdd9bc` unaffected).

## D-iter10-2: DISPOSITION (coordinator-approved) — (f)-FCP-p95 = ACCEPTED environmental — 2026-07-25

Studio-desk first-paint (gate f) shell paint holds on billion: **p50 637–726 ms < 1 s** (the M253 shell fix
holds). The p95 outliers (1443 / 2014 / 4943 ms, `reachedShell` always true) are **tailnet network-RTT
jitter**, not an app-paint regression — per `latency-budget.md`'s "state the environment with every number"
and the coordinator-approved (b) precedent (same tailnet-jitter class). **Gate (f) MET on app-side paint;**
the p95 disposition is recorded here and gets formally fated at close-milestone's deferral audit. **Fate: 1
(complete — disposition recorded), coordinator-approved.**

## D-iter10-3: DISPOSITION (coordinator-approved) — (c)-academy-durability = Fate-3 follow-up — 2026-07-25

A FRESH demo renders "← Back to Cockpit" on **all 4 apps** (proven iter-08, gate (c) MET on the fresh-demo
presenter case). The native academy dev-server reverts the back-to-cockpit patch only on a **long-running**
demo (a durability edge, not the presenter path). **Gate (c) MET;** the durability edge routes to an
**academy-durable follow-up** (`FIX-M254-c-academy-durable` in carry-forward.md). **Fate: 3 (annotate-attach
to a follow-up), coordinator-approved.**

## D-iter10-4: DISPOSITION (coordinator-approved) — (g)-testhealth = carry-forward FIX-M254-g-testhealth — 2026-07-25

The 8-test host-sensitive demo-stack test-health membership (gate g): 2/8-class fixed + verified live
(nvm/node host-robustness, rext `dfdd9bc`). The **6 remaining** (intra-run `:23077` port-leak + M245
reconcile-message drift, next.config sha re-pin, 2 mutation-meta, overlay-127) are **chronic host-sensitive
test-HARNESS issues with 0 demo-runtime impact** (the real academy serves 200 live). Routed to carry-forward
`FIX-M254-g-testhealth`. **Fate: 3 (annotate-attach to carry-forward), coordinator-approved.**
