---
iteration_type: tik
milestone: M254
iter: 08
status: closed-fixed-partial
---

# M254 · iter-08 — (c)-academy render + recovery re-bring-up (disruption + net-positive)

**Type:** tik · **Active strategy:** TOK-01 — cluster 4 (mutating tail: c-academy + e builder + h Playthroughs).

## Cluster / target
Complete (c) render (academy, the 4th app), then the mutating tail (e builder Playthrough, h Playthroughs).

## What happened (honest)
Attempted the (c)-academy re-heal (re-apply the Back-to-Cockpit patch to the RUNNING native academy). This
**cascaded into a self-inflicted demo disruption**: the re-apply crashed the academy's next-dev; relaunch
attempts hit EADDRINUSE; a `sudo fuser -k 13077/tcp` killed **tailscaled** (systemd auto-restarted it — SSH
recovered); then I mis-killed 2 host-visible `next-server` procs that were the WEB + HIRING **container**
processes → the containers exited and lost their host port mappings on restart → web/hiring 502. **Recovered
cleanly** with a full **cold reset-to-seed re-bring-up** at the latest rext pin — which turned the disruption
**net-positive**: a fresh cold reset-to-seed demo (the gate's ideal state).

## Realized
- Recovery re-bring-up GREEN (up_rc=0, autoverify green:true / 0 warnings / ts 08:16:34Z). Fresh demo, all
  origins serving.
- **(a) RE-CONFIRMED green** on a fresh cold reset-to-seed.
- **(c)-academy renders on the FRESH demo** (Back-to-Cockpit item present + href → :17700) → **(c) MET: 4/4
  apps render LIVE** (next-web + hiring + studio + academy) + the prod-eject side (0 escapes/133 + 0 studio
  prod-ejects).
- (e) builder Playthrough + (h) Playthroughs routed to iter-09 (the budget went to recovery).

## Escalation
The disruption was self-inflicted + cleanly recoverable via a standard re-bring-up — driven to recovery, not
escalated. 0 platform edits throughout.

## Lessons
Never `fuser -k` a port that `tailscale serve` fronts (it kills tailscaled). Container `next-server` procs are
visible in the host PID namespace — never kill demo procs by name/pid from the host; use `docker` verbs. A
native-academy patch re-apply to a running next-dev is unsafe — a fresh launch (or a full re-bring-up) is the
right way to (re)render it.
