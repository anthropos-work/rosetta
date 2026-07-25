**Type:** tik — TOK-01 cluster 4 (mutating tail). (c)-academy render + the recovery re-bring-up.

# M254 · iter-08 — progress

## Close — 2026-07-25

**Outcome:** A (c)-academy re-heal attempt cascaded into a self-inflicted demo disruption (crashed academy →
killed tailscaled [systemd-recovered] → mis-killed the web+hiring container procs → lost port mappings). Cleanly
**recovered via a full cold reset-to-seed re-bring-up** (GREEN, up_rc=0, autoverify green:true/0 warnings/ts
08:16:34Z) — which turned net-positive: a fresh cold reset-to-seed demo. **(a) re-confirmed green**;
**(c)-academy renders on the fresh demo** → **(c) MET (4/4 apps render LIVE + prod-eject side)**.

**Type:** tik
**Status:** closed-fixed-partial
**Gate:** (c) **MET** (4/4 render + prod-eject); (a) re-confirmed. (e) + (h)-Playthroughs routed to iter-09.
**Phase 5 grading:** (1) gate-met: n (overall) — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (the disruption was self-inflicted + cleanly recovered via a standard re-bring-up; no user decision changed what code lands — 0 rext/platform edits this iter) — (5) cap-reached: n (tik 4 of the session) — (6) protocol-stop: n — Outcome: continue (→ iter-09).
**Decisions:** D1 (c)-academy re-heal caused a disruption (root-caused); D2 recovery = cold reset-to-seed re-bring-up (re-pin dfdd9bc + rext.tag bump past the M217 guard); D3 (a) re-confirmed + (c) MET 4/4 on the fresh demo; D4 native-academy durability still a routed robustness gap (fresh bring-up renders it; a long-running demo can lose it).
**Side-deliverables:** none (0 rext/platform edits; recovery re-bring-up + measurement only).
**Routes carried forward:** (e) builder Playthrough; (h) Playthroughs + live-browser specs; (f)-FCP-p95 disposition; (g)-testhealth batch; (c)-academy-durability robustness (native academy loses the item on a long-running demo — routed to the g-testhealth/academy-durable follow-up).
**Lessons:** (1) Never `fuser -k` a port that `tailscale serve` fronts — it kills tailscaled and drops the whole tailnet (SSH + origins); tailscaled is systemd-managed so it auto-restarts, but the blast radius is total. (2) Docker container `next-server` procs are visible in the host PID namespace — killing "by name/pid" from the host kills the CONTAINER's app; always use `docker` verbs on demo containers. (3) A `docker start` after a container's proc was host-killed can come back WITHOUT its host port mappings (`docker port` empty) — a full re-bring-up is the reliable restore. (4) Re-applying a native-academy source patch to a running next-dev is unsafe (HMR crash); render it via a fresh launch / re-bring-up. (5) A recovery re-bring-up doubles as the gate's ideal fresh-cold-reset-to-seed state — the disruption became net progress ((a) reconfirm + (c) 4/4).
