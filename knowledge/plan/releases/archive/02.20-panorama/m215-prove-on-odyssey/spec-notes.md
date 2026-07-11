# M215 — spec-notes

_(accumulates iteration-protocol-specific technical notes as iters run)_

## Target environment (verified live 2026-07-11)
- **Host `odyssey`** (Proxmox, Hetzner): SSH `root@100.112.141.83` (Tailscale), Tailscale 1.98.8; runs **13 VMs**
  (the `odyssey`-skill KB listing 4 is stale). **128 GB physical, ~180 GB configured (1.4× over-committed), ~94%
  used → only ~7 GB free host-wide.**
- **VM `billion`** (the ONLY VM we may touch — user constraint 2026-07-11): Ubuntu 24.04, Docker 29.6.1 + Compose
  v5.3.0, MagicDNS `billion.taildc510.ts.net`, **`tailscale cert` proven remotely trusted (`verify=0`, no CA
  install)**, ~168 GB free disk. RAM: **32 GB ceiling / 8 GB balloon floor, currently pinned at the 8 GB floor**
  (host saturated). **16 GB swap net added** (persistent `/swapfile`, `swappiness=10`) → a ~12 GB demo won't OOM
  (spills to swap, slower). Host-RAM reclaim is **off-limits** (can't touch other VMs), so M215 runs **swap-backed
  and/or trimmed** (e.g. next-web only, skip studio-desk + academy) rather than the full UI tier when RAM is tight.
- Reach billion over Tailscale SSH (`devops@billion`); a second tailnet machine (or headless browser) drives the UI.
- MagicDNS enabled tailnet-wide (suffix `taildc510.ts.net`); the HTTPS cert is bound to the **name**, so clients must
  resolve via MagicDNS (default-on when Tailscale is installed) — the raw `100.x` IP would fail cert matching.

## Remote-drive harness
- Reuse the M42 e2e Playwright harness (`rext stack-verify/e2e/`) + the M202 Playthrough seat-switch (cockpit hero
  login), pointed at the remote `https://billion.taildc510.ts.net` origin instead of localhost.

## Failure taxonomy to capture per run
localhost-eject · prod-eject · CORS block · cert-untrusted · mixed-content · auth-incomplete (secure-context /
cookie same-site) · asset-missing · RAM/OOM.
