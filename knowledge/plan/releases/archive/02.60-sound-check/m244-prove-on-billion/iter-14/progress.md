**Type:** tik (run 6, tik 2). Active strategy: TOK-02.

# iter-14 — progress

## Seed-model finding (decides the whole gate-(c) + remaining-gate ordering)
The 16 Playthroughs seed the dedicated **pt-world** (`playthroughs/e2e/seed/pt-world.seed.yaml`, 3 private orgs) and `run-playthroughs.sh --reset` does a full FK-ordered TRUNCATE + re-seed + roster swap — which **DESTROYS the content/stories demo seed** currently live on billion. The 24 stack-verify specs (coverage/calibrate/persona/verify/m2xx/talk-to-data/…) run against the DEMO/stories seed. So: do ALL demo-seed gate work (24 stack-verify, f, h, d) BEFORE any pt-world reset, and run the Playthroughs LAST (so the pt-world reset happens once, at the end, with no demo re-bring-up). iter-14 therefore drove the demo-seed coverage sweep, read-only, no reset.

## Coverage sweep — both hero vantages, live on billion (`COVERAGE_HOST` + https, foreground ~5.7m each)
- **employee (maya-thriving):** reachable 59/150 · failingSections **0** · personaFailures **0** · escapes **0** · notReached **0** · frontier EXHAUSTED · **crossPortFailures 1** → GATE NOT MET.
- **manager (dan-manager):** reachable 70/150 · failingSections **0** · personaFailures **0** · escapes **0** · notReached **0** · frontier EXHAUSTED · **crossPortFailures 1** → GATE NOT MET.

Both vantages fail on the **identical** single cross-port follow: `/home` → `https://billion…:13077/` (the ant-academy home), detail: *"ant-academy home marker absent (empty: region missing required content: AI Academy — re-asserted 6× over the heavy-grid budget, genuinely below bar)"*. Everything else (every workforce/enterprise page, persona self-consistency, 0 prod-eject escapes, all manifest pages reached) is GREEN.

## The (c)↔(d) coupling — the pivotal finding
The one coverage failure IS the gate-(d) academy-empty root cause (iter-09: the anon academy grid reads a tenant-filtered/empty catalog; `/free` renders 43 cards but `/library` + the home render empty). So **gate (c)'s coverage half cannot go fully green until the gate-(d) ant-academy demopatch lands** — the academy fix unblocks BOTH the coverage cross-port (c) AND gate (d). This RE-ORDERS the remaining work: do gate (d) NEXT (iter-15), then re-run the coverage sweep to green, then the discrete stack-verify specs, then the Playthroughs last on pt-world.

## Close — 2026-07-22

**Outcome:** Coverage sweep (both hero vantages) is GREEN on billion EXCEPT one shared failure — the cross-port follow to the empty ant-academy home (:13077), the gate-(d) academy-empty root cause. Gate (c) does NOT tick (needs the academy fix first). Documented the (c)↔(d) coupling + re-ordered the remaining gates. Metric stays **4/8**.
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET (gate part (c) blocked on the gate-(d) academy fix)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik 2/5) — (6) protocol-stop: n — Outcome: **continue**
**Decisions:** D1, D2 (iter-14/decisions.md).
**Side-deliverables:** none.
**Routes carried forward:**
- iter-15 = **gate (d)** ant-academy demopatch (feed the FS catalog to the anon `/library` + home grid; iter-09 root-cause) + re-bake academy on billion → unblocks the coverage cross-port (c) AND ticks (d). Includes the `/free` 2-Draft-chip sub-finding.
- then finish gate (c): re-run the coverage sweep (→ green), run the discrete stack-verify specs, run the 16 Playthroughs LAST (pt-world reset).
**Lessons:** a "prove the specs green" gate can be COUPLED to a sibling gate through a cross-port follow — the workforce coverage sweep follows into the academy app, so the academy's own defect fails the workforce gate. Measure before sequencing: iter-14's sweep turned "run gate c then gate d" into "gate d unblocks gate c," saving a wasted discrete-spec pass against a known-blocked academy.
