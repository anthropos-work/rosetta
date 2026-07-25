---
title: "M253 Carry-Forward — Routes from studio-desk first-paint"
date: 2026-07-24
status: archived
close_status: closed-on-gate
gate_target: "cold demo (state environment — laptop vs tailnet), first-meaningful-paint < 1000 ms (the .page-skeleton header+sidemenu shell visible) AND no blank > 1 s, p95 over 5 consecutive cold loads; never gate on networkidle; always gate on a fresh-green autoverify.json"
gate_achieved: "skeleton-visible p95 817 ms (p50 743, max 817) over 5 consecutive cold loads on demo-2 (LOCAL LAPTOP), 5/5 reached the shell, 0 login bounces; baseline 4669 ms (~5.7x faster) — numerical gate MET"
gate_distance: "numerical gate MET on the local-bootstrap charter; residual = the fully-green COLD-p95 confirmation on billion (fresh-green autoverify.json)"
---

## TL;DR
M253 **closed on gate** — studio-desk first-meaningful-paint dropped **4669 ms → p95 817 ms** on demo-2 (local
laptop), 5/5 cold loads, 0 login bounces, decisively under the < 1000 ms budget. The fix is a pure zero-platform-edit
paint-ordering demopatch pair on the M249 ladder. **One coordination-split route** carries forward: the **fully-green
COLD-p95 confirmation on billion**. This is **not a gate miss** — it is the deliberate **coordination-rule-9 split**
(M253 bootstraps the FCP loop on a LOCAL demo because two live-measured iteratives can't share billion's RAM; the
cold-on-billion confirmation is chartered to M254). It routes **Fate 2 → M254**, whose exit gate part **(f)** already
names it verbatim.

## Root-cause clusters

### Cluster 1 (CARRY-M253-01): fresh-green COLD-p95 confirmation on billion
- **Affected items:** re-measure the studio-desk first-paint FCP gate on a **freshly brought-up, fully set-dressed
  COLD billion demo** with a **green `autoverify.json`** — the formal fresh-green clause of the exit gate (satisfied
  numerically on demo-2, but demo-2 is warm/partially-set-dressed so a fully-green verdict there is unrelated to
  studio first-paint and unachievable).
- **Root cause:** NOT a defect — a **deliberate coordination split**. Per coordination rule 9 (overview.md), two
  live-measured iteratives (M250 + M253) cannot share billion's RAM, so M253 was chartered to **bootstrap on a local
  demo** with the cold-p95 confirmation reserved for the closer. The numerical gate is **fully MET** on demo-2
  (p95 817 ms < 1000 ms, 5/5 cold, 0 bounce); the code-of-record fix rides M249's ladder and is baked into the
  studio image at build. Nothing to build — only to *observe cold + green at scale on billion*.
- **Estimated scope:** zero new build; one cold reset-to-seed studio-FCP sweep on billion via `run-studio-fcp.sh`,
  gated on a fresh-green `autoverify.json`.
- **Fate:** **Fate 2** — already owned by M254 (no roadmap edit; the target's plan already covers it).
- **Target milestone:** **M254** (the terminal closer) — `overview.md:7` exit_gate part **(f)** "studio first-paint
  **< 1 s cold p95**" + `:87` "(f) studio **first-paint < 1 s cold p95** — (← M253)". `M254 depends_on: [… M253 …]`.
- **Provenance:** established at authoring time (overview.md coordination rule 9); confirmed numerically iter-02;
  routed at iter-02/iter-03 close and re-affirmed Fate-2 at the M253 close deferral audit.

## Projected post-resolution state
When M254 brings a demo up cold + fully set-dressed on billion and runs the studio-FCP sweep against a fresh-green
`autoverify.json`, first-meaningful-paint is expected to hold **p95 < 1000 ms** (no blank > 1 s) over 5 consecutive
cold loads — closing exit gate part (f). Confidence is high: the number is already met with margin on the local box
(817 ms vs 1000 ms) and the fix is a deterministic paint-ordering change baked into the image; the only unknown is the
tailnet-VM environment (which latency-budget.md requires be stated with every number). If billion measured slower, M254
triages it under its own iterative loop (no M253 re-open needed).

## Cross-references
- Gate Outcome Ledger: ./progress.md (§ Gate Outcome Ledger)
- Iter ledger: ./progress.md (§ Running ledger)
- Decisions: ./decisions.md (TOK-01) + iter-02/decisions.md (D5 green-gate → M254)
- Deferral audit: ./audit-deferrals/deferral-audit-2026-07-24-m253-close.md
- Iteration protocol used: corpus/ops/demo/latency-budget.md + corpus/ops/demo/coverage-protocol.md
- Delivers (KB): corpus/ops/demo/latency-budget.md (studio first-paint budget) + corpus/ops/demo/demopatch-spec.md (the 2 patches) + corpus/services/studio-desk.md (MPA boot model)
- Code-of-record: rosetta-extensions @ july-jitter-m253-studio-first-paint (b8969c0, on origin)
