# M215 — spec-notes

_(accumulates iteration-protocol-specific technical notes as iters run)_

## Target environment (verified live 2026-07-11)
- **Host `odyssey`** (Proxmox, Hetzner): SSH `root@100.112.141.83` (Tailscale), Tailscale 1.98.8; runs 7 VMs.
- **VM `billion`:** Ubuntu 24.04, Docker 29.6.1 + Compose v5.3.0, MagicDNS `billion.taildc510.ts.net`,
  `tailscale cert` available, ~174 GB free disk, **~13 GB RAM available (ballooned — host over-committed)**.
- Reach a VM directly over Tailscale SSH (`devops@<vm-tailscale-ip>`); a second tailnet machine drives the browser.
- MagicDNS enabled tailnet-wide (suffix `taildc510.ts.net`).

## Remote-drive harness
- Reuse the M42 e2e Playwright harness (`rext stack-verify/e2e/`) + the M202 Playthrough seat-switch (cockpit hero
  login), pointed at the remote `https://billion.taildc510.ts.net` origin instead of localhost.

## Failure taxonomy to capture per run
localhost-eject · prod-eject · CORS block · cert-untrusted · mixed-content · auth-incomplete (secure-context /
cookie same-site) · asset-missing · RAM/OOM.
