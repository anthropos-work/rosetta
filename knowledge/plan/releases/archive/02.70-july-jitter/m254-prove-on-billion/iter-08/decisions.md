# M254 · iter-08 — decisions

## D1 — the (c)-academy re-heal caused a self-inflicted disruption (root-caused)
Re-applying `apply-ant-academy-back-to-cockpit.sh apply` to the RUNNING native academy triggered a next-dev HMR
recompile that crashed the academy (502). Relaunch attempts hit `EADDRINUSE 0.0.0.0:13077` (zombie next-server
workers held the port). A `sudo fuser -k 13077/tcp` killed **tailscaled** (pid 3549 — it fronts :13077 via
`tailscale serve`), dropping the tailnet (SSH + all origins); systemd auto-restarted it. Then I `sudo kill`ed 2
host-visible `next-server` procs believing they were academy workers — they were the **web-app + hiring-app
container** processes (root-owned, parents `pnpm @anthropos/web-app|hiring-app`, visible in the host PID
namespace). The containers exited and, on `docker start`, came back Up + Ready internally but with **no host
port mappings** (`docker port` empty) → web/hiring 502.

## D2 — recovery: full cold reset-to-seed re-bring-up (re-pin + rext.tag bump)
Recovered with the coordinator's reset-to-seed flow (down 1 --purge → tailscale serve reset → advance-pinned
up-injected → autoverify). First attempt hit the **M217 FATAL pin guard** (clone checked out to
`july-jitter-m254-academy-nonode-hostrobust` but `/home/devops/panorama/.agentspace/rext.tag` still pinned
`aireadiness-repoint`); teardown had already succeeded. Bumped `rext.tag` to match, re-kicked → **GREEN**
(up_rc=0, autoverify green:true / 0 warnings / ts 08:16:34Z). billion now pinned to dfdd9bc (all M254 fixes).

## D3 — (a) re-confirmed + (c) MET (4/4) on the fresh demo
The recovery re-bring-up IS a fresh cold reset-to-seed (the gate's ideal). (a) re-confirmed green. (c)-academy
renders live on the fresh demo (ACADEMY_BTC:1, href → https://billion.taildc510.ts.net:17700/). Combined with
next-web + hiring (iter-05) + studio (iter-06), **(c) render side is 4/4 LIVE** + the prod-eject side
(0 escapes/133 + 0 studio prod-ejects) → **(c) MET**.

## D4 — native-academy Back-to-Cockpit durability: routed robustness gap (not gate-blocking)
A FRESH cold reset-to-seed renders the academy item (the launch path applies it + next-dev compiles it — proven
here). The durability gap (a long-running demo can lose it after an out-of-band clone revert, healed only on an
`ant-academy.sh` re-invoke) is a separate robustness issue routed to the academy-durable / g-testhealth
follow-up. The GATE ("cold reset-to-seed … all 4 apps") is satisfied.
