---
iter: 01
milestone: M226
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
date: 2026-07-17
---

# iter-01 (bootstrap tok) — author the billion-proof strategy

**Type:** tok (bootstrap). No gate progress; authors TOK-01 (the first strategy) + a read-only recon of billion.

## Inputs consumed
- `overview.md` (the 7-condition exit gate), `spec-notes.md`, the milestone `progress.md` next-iter note.
- Protocol: `corpus/ops/verification.md` + `corpus/ops/demo/coverage-protocol.md` + `corpus/ops/demo/latency-budget.md`.
- The prove-on-billion lineage: `knowledge/plan/releases/archive/02.30-cue-to-cue/m221-prove-on-billion/` (+ M215 prove-on-odyssey), `corpus/ops/demo/tailscale-serve.md`, `corpus/ops/safety.md`.
- Phase 0b KB-fidelity audit: **GREEN** (`kb-fidelity-audit.md`).
- A careful **synchronous read-only** recon of billion (peer-side `tailscale`/`ping` + one read-only `ssh devops@billion`).

## Recon result (billion current state, 2026-07-17)
- **Reachable**: 100.110.136.3, idle, ping ~59 ms. SSH user is **`devops`** (Tailscale-SSH `kirality` mapping fails).
- **A stale demo is UP**: `demo-1-*` (16 containers) "Up 2 days" at rext tag **`cue-to-cue-v2.3.2`** — the v2.3
  panorama demo (NO hiring org). Holds base ports 13000/15050/18082/19000/17700/13077 + a live `tailscale serve`
  config on those ports. **Must be cleanly torn down first** (incl. serve reset + academy respawner reap, the M221
  F5/F5b/F12 discipline) before a fresh casting-call bring-up.
- **Host prereqs GREEN**: Go 1.25.12, atlas v1.2.4-canary, tailscale 1.98.8. Stack root `~/panorama/stack-demo/`.
- **C-6 memory risk (sharper here)**: 7.3 GiB RAM (4.5 Gi avail + 14 Gi free swap); disk 40 G free (80% used). The
  hiring demo adds a **2nd UI container** (`apps/hiring`) beyond the panorama demo → more pressure. Watch OOM.
- **rext on billion = `cue-to-cue-v2.3.2`** → cut over to the casting-call code-of-record `casting-call-m225-harden`.

## Strategy authored
See milestone-root `decisions.md` → **TOK-01: reprove-hiring-on-billion**. In one line: cold-teardown the stale
v2.3 demo → retag billion's rext + platform to the casting-call code-of-record → run a **default `/demo-up 1`
(no flags)** synchronously on billion (remote-reach default-on) → measure the 7-condition gate **from this Mac** →
attribute + fix any gap (R1 render re-surface / R4 45×5 hydration latency) via tooling / a sha-pinned demo-patch
re-proven live; a platform-repo edit is never in bounds → escalate.

## Close
See `progress.md`.
