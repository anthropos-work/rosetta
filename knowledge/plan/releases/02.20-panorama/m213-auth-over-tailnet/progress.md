# M213 — progress

Section checklist (closure = the fake FAPI serves a trusted `tailscale cert` for the MagicDNS host, the re-minted
pk validates, and the browser surface is reachable under one HTTPS origin — proven locally before the M215 live run).

- [ ] cert mint swaps mkcert/openssl → `tailscale cert` when `STACK_PUBLIC_HOST != localhost` (same output paths)
- [ ] pk re-mint validates (`assertValidPublishableKey` dotted-host pass + base64 round-trip)
- [ ] reverse proxy (`tailscale serve` / Caddy) fronts the browser-facing ports on the MagicDNS host, one origin
- [ ] topology guard: FAPI same host as app; PSL status of taildc510.ts.net verified
- [ ] build-rebuild guard trips on HOST change (no stale localhost image)
- [ ] fake-fapi egress to cdn.jsdelivr.net confirmed
- [ ] tests (cert-path selection, pk validation, proxy config generation)
