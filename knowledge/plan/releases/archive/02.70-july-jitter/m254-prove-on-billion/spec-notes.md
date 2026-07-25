# M254 — Spec notes

Iterative milestone (the closer). Accumulates per-iter / per-gate design notes for the live re-prove on
`billion` — the cold reset-to-seed bring-up, the multi-part exit-gate (a–h) evidence, the read-only
confirmation-sweep findings, and the serial mutating tail — recorded during the iter loop.

**Iteration protocol:** `corpus/ops/verification.md` + `corpus/ops/demo/tailscale-serve.md` +
`corpus/ops/demo/coverage-protocol.md` + `corpus/ops/demo/playthroughs.md` (bring-up →
measure→confirm→fix-forward).

## Pre-flight audits — iter-01
- **Phase 0b KB-fidelity: GREEN-by-inheritance** (no standalone `audit-kb-fidelity` spawn). M254 is a terminal
  PROOF milestone (`Delivers: none`) — it authors no new knowledge doc; its "knowledge" is the 4 protocol docs
  (`verification.md`, `tailscale-serve.md`, `coverage-protocol.md`, `playthroughs.md`), each continuously
  updated + validated through M246–M253 (all just closed). The live cold bring-up + gate sweep is itself a
  stronger fidelity check than a static pre-audit; any doc drift surfaces as an iter finding at measurement
  time. See iter-01/decisions.md D1. Standing verdict for the milestone unless a triggered tok redirects into
  an un-covered subsystem.

## The DRIVE — cold reset-to-seed bring-up on billion (single-driver serial)
- **Pin:** `july-jitter-m253-studio-first-paint` (`b8969c0`) — cumulative code-of-record (all 7 milestone tags
  are ancestors of rext main HEAD), confirmed on origin (rung-zero).
- **billion state at bootstrap:** reachable (up 14 days), `docker ps` empty, no `~/stack-*` (clean slate), the
  `tailscale serve` config persists on offset ports (13000 web · 13001 hiring · 13077 academy · 15050 cosmo ·
  18082 backend · 19000 studio · 17700 cockpit · 15400 FAPI).
- **Bring-up command (on the VM):** `STACK_PUBLIC_HOST=billion.taildc510.ts.net bash demo-stack/up-injected.sh
  1 --public-host billion.taildc510.ts.net` — run foreground-blocking inside a tracked background Bash; never
  detach-and-yield; never kill a mid-build.
- _(Live evidence recorded during iter-02+.)_

## The multi-part exit gate (a–h) — live evidence
_(TBD during build.)_

## Read-only confirmation sweeps (content-stories ∥ coverage ∥ probes)
_(TBD during build.)_

## Mutating drift-carries + seed-destroying Playthroughs (serial tail)
_(TBD during build.)_
