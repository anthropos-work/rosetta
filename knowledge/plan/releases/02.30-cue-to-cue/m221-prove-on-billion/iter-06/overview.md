---
iter: 6
milestone: M221
iteration_type: tik
iter_shape: tik
status: closed-fixed
opened: 2026-07-15
closed: 2026-07-15
gate: MET
strategy_ref: TOK-01 (../decisions.md)
---

# iter-06 (tik) — the FINAL live demo cycle on `billion`, LEFT RUNNING

**Type:** tik, under TOK-01 ("fix-then-prove, host-isolation FIRST"). iter-05 landed the r4 correction and
the r4 re-prove ≈ MET the 8-condition gate on ONE cold cycle (PROVISIONAL — reproducibility + the
browser-level Dana grade remained). This tik runs **exactly ONE** DEFAULT cold `up-injected.sh 1` (NO FLAGS)
at the r4 code on `billion`, grades all 8 gate conditions + folds in M219's readiness gates thoroughly
(incl. the **browser-grade** of Dana's `/ai-readiness` 900-char floor that the prior cycle only DB-confirmed),
and — per the user's explicit ask — **LEAVES THE STACK UP as the final live demo** (no teardown).

The user wants exactly ONE cycle (not a reproducibility battery) + a working live demo left on the box. The
close treats iter-05's prior r4 cycle + this cycle together as satisfying the gate's "reproducibly on a cold
reset-to-seed" pragmatically.

## The cycle

A DEFAULT `up-injected.sh 1`, NO FLAGS, cold reset-to-seed, remote-by-default auto-discovery, driven +
graded from a tailnet peer (this Mac). Cold proof: PG_VERSION mtime from INSIDE the container > T0. Green
proof: fresh `autoverify.json` ts.

## The grade — 8 gate conditions + M219 readiness fold-in + F10/F4

_(gate table + measurements filled at close — see `progress.md`.)_

## Then LEFT LIVE
No teardown. Stack UP, `tailscale serve` fronting it, reachable from the peer (a hero login works
end-to-end). Live URLs recorded in the close.
