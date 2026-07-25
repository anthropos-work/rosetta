---
iter: 01
milestone: M254
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
created: 2026-07-24
---

# M254 · iter-01 — bootstrap tok (author TOK-01)

**Type:** tok (bootstrap) — iter-01 of the milestone; authors the FIRST strategy. No prior iters, no
stalled strategy to revise. Does NOT terminate the invocation (Phase 5 §2) — the loop continues into
iter-02 (first tik) under TOK-01.

## Inputs
- `overview.md` — the verbatim a–h exit gate (cold reset-to-seed on billion, driven from a tailnet peer,
  0 platform edits).
- `roadmap.md` § "#### M254 —" — the intra-milestone LANE decomposition (DRIVE serial → read-only
  confirmation sweeps fan-out → mutating serial tail).
- Protocol docs: `corpus/ops/verification.md` + `corpus/ops/demo/tailscale-serve.md` +
  `corpus/ops/demo/coverage-protocol.md` + `corpus/ops/demo/playthroughs.md`.
- The M221/M236/M244 prove-on-billion lineage + the driving-prove-on-billion playbook memory.

## Baseline (measured at bootstrap)
- **Pin:** `july-jitter-m253-studio-first-paint` (`b8969c0`) — the cumulative rext code-of-record; ALL 7
  milestone tags (m246-harden … m253) are ancestors of rext main HEAD; the tag is ON ORIGIN (rung-zero
  satisfied). This is what the billion demo consumes.
- **billion:** reachable (`ssh marco@billion`, up 14 days). `docker ps` EMPTY — no demo up (clean slate for
  the cold bring-up). `tailscale serve` config persists on the offset ports (13000 apps/web · 13001
  apps/hiring · 13077 academy · 15050 cosmo · 18082 backend · 19000 studio · 17700 cockpit · 15400 FAPI).
- **Gate:** 0/8 parts confirmed live (all carry forward from M246–M253's local-provisional gates).

## Initial strategy → TOK-01 (see milestone-root decisions.md)
One gate-cluster per tik, measure→confirm→fix-forward, following the LANE decomposition. Any defect a gate
part surfaces routes to rext (or a sha-pinned demopatch) — 0 platform-repo edits; commit + tag +
`git push --tags` to origin before re-pinning billion.

## Next-tik direction (iter-02)
**The DRIVE (gate a):** cold reset-to-seed `/demo-up` on billion at pin `july-jitter-m253-studio-first-paint`
via `up-injected.sh 1 --public-host billion.taildc510.ts.net`, run foreground-blocking inside a tracked
background Bash (never detach-and-yield on billion; never kill a mid-build). Assert gate (a): builds + comes
up GREEN on the consolidated platform (3 subgraphs, skillpath-in-app), health 200 + casbin > 0, fresh green
`autoverify.json`. This is the enabling precondition — every downstream proof gates on it.
