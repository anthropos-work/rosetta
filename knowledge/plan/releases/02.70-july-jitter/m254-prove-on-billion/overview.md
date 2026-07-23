---
milestone_shape: iterative
milestone: M254
title: "prove on billion (the closer)"
status: planned
release: v2.7 "july jitter"
exit_gate: "Cold reset-to-seed on billion, driven from a tailnet peer, 0 platform edits: (a) the re-grounded stack builds + comes up GREEN on the consolidated platform (3 subgraphs, skillpath-in-app); (b) the content-stories manager CTA lands on the /sim per-session manager result view (non-empty) for sim products; (c) Back-to-Cockpit works in all 4 apps + studio logo/back/logout resolve to the stack app (0 prod-ejects); (d) the AI-readiness page faithful per M250 gate, live, both vantages; (e) the studio sim-builders generate (builder Playthrough green); (f) studio first-paint < 1 s cold p95; (g) the 8 live/docker-gated demo-stack test-health tests green (count confirmed 8 at M251/M247 close, host-sensitive membership — NOT the earlier ~2 estimate); (h) the live-browser specs + content-stories sweep + Playthroughs green; p95 click-to-ACCESS < 5 s hero vantages."
iteration_protocol_ref: corpus/ops/verification.md
re_scope_trigger: "5 consecutive toks without a viable strategy -> user-strategic-replan"
depends_on: [M247, M248, M249, M250, M251, M252, M253]
complexity: medium
created: 2026-07-23
---

# M254 — prove on billion  (`iterative`, the closer)

**Status:** `planned`  ·  **Shape:** `iterative` (the closer)  ·  **Complexity:** medium (iterative)  ·  **Release:** v2.7 "july jitter"  ·  **Depends on:** M247, M248, M249, M250, M251, M252, M253 (all fixes)

## Goal
Re-prove the whole release **live on `billion`**, cold reset-to-seed. Every v2.7 fix — the re-ground, the
content-stories manager link, the cross-app Back-to-Cockpit + studio prod-eject fix, the AI-readiness fidelity,
the studio builder keys, and the studio first-paint — proven together on the `billion` Tailscale VM, driven and
asserted from a tailnet peer, **0 platform edits**. The terminal live-proof of the release (the M221/M236/M244
prove-on-billion lineage).

## Shape (why this shape)
`iterative`. Live-proof is **measurement-driven** (the M221/M236/M244 lineage): a cold `billion` bring-up is the
critical path and highest-risk step, and every downstream proof (content-stories sweep, coverage, Playthroughs,
latency) gates on a **fresh green `autoverify.json`**. The multi-part exit gate (a–h) is discharged one cluster
per tik, measure→confirm→fix-forward, iterating until all parts are green cold. A section shape can't express the
"bring it up, measure, route the residual, re-measure" loop the live box demands.

## Scope
### In
- **The DRIVE (single-driver serial)** — bring up a fresh demo on `billion`, **cold reset-to-seed**, at the v2.7
  rext pin; drive + assert every fix from a **tailnet PEER** (this workstation), recording live evidence. The
  enabling foundation: every downstream proof gates on a fresh green `autoverify.json`, so a green cold bring-up
  is the precondition. Same billion-safety rules as the lineage (one driver, rung-zero pin-on-**origin** before
  billion consumes it, assert from a peer never the VM, never kill a mid-build).
- **The multi-part exit gate (a–h)** — run + confirm each part live (see ## Exit gate). This is the spec.
- **Read-only confirmation sweeps (fan-out)** — **content-stories ∥ coverage ∥ probes** fan out as concurrent
  tailnet peers against **ONE** bring-up (latency measured solo, quiet system). The ~1.4–1.8× confirmation-window
  parallelism.
- **Mutating drift-carries + seed-destroying Playthroughs (serial tail)** — the parts that mutate or re-seed stay
  a **serial tail** after the read-only sweeps (they can't share a bring-up).

### Out
- **New feature work** — all built by M247–M253. M254 CALIBRATES + PROVES live; it does not re-build.

## Dependencies & parallelism
- **Depends on:** M247, M248, M249, M250, M251, M252, M253 (all fixes). **Terminal** — the last milestone of the
  release; every gate part traces to an upstream milestone's delivery.
- **Parallel with:** **none** (terminal — it proves what the fan-out built).
- **Intra-milestone LANE decomposition:** **~1.4–1.8× on the confirmation window only.**
  - **The DRIVE is single-driver serial** (bring-up + drive from a tailnet peer) — the critical path, un-shardable.
  - **Read-only confirmation sweeps fan out** — **content-stories ∥ coverage ∥ probes** run as concurrent tailnet
    peers against ONE bring-up; **latency runs solo** (quiet system, or the p95 is polluted).
  - **The mutating drift-carries + the seed-destroying Playthroughs stay a serial tail** (they mutate / re-seed,
    so they cannot share the read-only bring-up).
  - **Recommended subagents:** **1 serial driver** for the bring-up + the mutating tail; during the confirmation
    window fan out **~3 read-only peer sweeps** (content-stories ∥ coverage ∥ probes) with latency taken solo.
- **Coordination (release-level):** rung-zero every push — rext tags on **origin** before the `billion` re-pin
  (rule 7; M236 lost an iter to an unpushed tag). Live-box contention — M250 + M253 (both live-measured
  iteratives) serialize on one `billion` demo upstream (rule 9); M254 runs after both land. **Merge/close order:**
  M251 → { M248, M250 } → M249 → M253 → M252 → M247-reconcile → **M254** (the terminal close).

## KB dependencies
- `corpus/ops/verification.md`
- `corpus/ops/demo/tailscale-serve.md`
- `corpus/ops/demo/coverage-protocol.md`
- `corpus/ops/demo/playthroughs.md`

## Delivers
- **None** (proof milestone — it re-proves the release live; it authors no corpus doc, extends the
  coverage/Playthrough manifests + burns in the confirmation).

## Exit gate
Cold reset-to-seed on `billion`, driven from a tailnet peer, **0 platform edits**:
- **(a)** the re-grounded stack **builds + comes up GREEN** on the consolidated platform (**3 subgraphs,
  skillpath-in-app**) — (← M246 / M247);
- **(b)** the content-stories **manager CTA lands on the `/sim` per-session manager result view** (**non-empty**)
  for the sim products — (← M248);
- **(c)** "← Back to Cockpit" works in **all 4 apps** + the studio logo/back/logout **resolve to the stack app**
  (**0 prod-ejects**) — (← M249);
- **(d)** the **AI-readiness page faithful** per M250's gate, **live, both vantages** — (← M250);
- **(e)** the studio **sim-builders generate** (the **builder Playthrough green**) — (← M252);
- **(f)** studio **first-paint < 1 s cold p95** — (← M253);
- **(g)** the **8 live/docker-gated demo-stack test-health tests green** (count confirmed 8 at M251 close, host-sensitive membership — `test_purge` + the `test_ant_academy` launcher/reap set + a host-isolation/clerk-wiring assertion; NOT the earlier ~2 estimate) — (← M251);
- **(h)** the **live-browser specs + content-stories sweep + Playthroughs green**; **p95 click→ACCESS < 5 s** hero
  vantages — (M254 itself, the overall live re-prove).

## Iteration protocol
- **Protocol:** `corpus/ops/verification.md` + `corpus/ops/demo/tailscale-serve.md` +
  `corpus/ops/demo/coverage-protocol.md` + `corpus/ops/demo/playthroughs.md` (bring-up →
  measure→confirm→fix-forward, one gate-cluster per tik; always gate on a fresh green `autoverify.json`, never on
  `networkidle`).
- **Re-scope trigger:** **5 consecutive toks without a viable strategy → user-strategic-replan.**

## Open questions
- **None blocking** — the multi-part exit gate (a–h) is the spec.
