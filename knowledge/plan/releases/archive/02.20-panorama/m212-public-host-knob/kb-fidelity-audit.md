---
title: "KB Fidelity Audit — M212 the single host knob"
date: 2026-07-11
scope: milestone:M212
invoked-by: build-milestone
---

## Verdict
YELLOW — proceed with tracking. No stale load-bearing claim; the one blind area (the
`STACK_PUBLIC_HOST` knob's own doc) is explicitly homed to M214's `corpus/ops/demo/tailscale-serve.md`
(M212 overview: "No new doc lands here"). One incidental completeness gap tracked as KB-1.

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Browser-facing URL baking (offset ports) | `corpus/ops/demo/frontend-tier.md` · `corpus/ops/rosetta_demo.md` | `demo-stack/up-injected.sh` (build-args, `.env.local` overlays, `want_ep`) · `stack-injection/gen_injected_override.py` | PAIRED |
| pk / FAPI host (Clerkenstein) | `corpus/services/clerkenstein.md` | `stack-injection/inject.py` (`mint_pk`) · `up-injected.sh:590,928` (fapi-host `127.0.0.1`) | PAIRED |
| `demo_web` content-URL rewrite | `corpus/ops/demo/coverage-protocol.md` (the prod-eject gate; the milestone cross-ref to `snapshot-spec.md` is a near-miss) | `up-injected.sh:876-881` | PAIRED |
| `STACK_PUBLIC_HOST` host knob | — (homed to M214 `corpus/ops/demo/tailscale-serve.md`) | (net-new this milestone) | BLIND-AREA (homed to M214) |

## Fidelity Findings

1. **frontend-tier.md — baked URLs use `localhost` + offset ports** — Source: `frontend-tier.md` §"How the pk + URLs are baked" (L143-218), §"Offset-origin CORS". Expected: build-args/overlays bake `http://localhost:<base+offset>`; CORS emits `CORS_EXTRA_ORIGINS=http://localhost:1300…`. Actual: exactly that (up-injected.sh:185-198,264-266,293-295; gen_injected_override.py:305). Verdict: **ALIGNED**. M212 preserves these verbatim when the knob is unset (byte-identical); the host-parametric emission (CORS + Clerk-URL) is explicitly M214. Fix owner: M214 updates the doc when the emission becomes host-parametric.

2. **rosetta_demo.md — offset/cockpit port model** — Source: `rosetta_demo.md` (offset scheme, `:7700+` cockpit). Actual: matches up-injected.sh. Verdict: **ALIGNED**.

3. **clerkenstein.md — pk/FAPI minter** — Source: `clerkenstein.md` (L57-115: `mintpk` the authoritative minter, the codec alignment-gated, `clerk-frontend` mints JWTs). Actual: `inject.py::mint_pk` mirrors `MintPublishableKey` (test-gated), host-parametric. Verdict: **ALIGNED** (no stale claim). Completeness gap: the doc does not state the *dotted-host* constraint (`assertValidPublishableKey` rejects dotless `localhost` → the FAPI host defaults to `127.0.0.1`, not `localhost`). That constraint is thoroughly commented in code (up-injected.sh:581-586, inject.py:50). It is the very reason a MagicDNS host works for the pk — the natural home is M214's `tailscale-serve.md`. Tracked as **KB-1** (non-critical; the M212 default preserves `127.0.0.1`).

4. **coverage-protocol.md — content-URL rewrite / prod-eject** — Source: `coverage-protocol.md` L200 ("Out-of-demo link (escape)"; the M50 rewrite to the offset port). Actual: up-injected.sh:876-881 rewrites `anthropos.work` → `http://localhost:<3000+offset>`. Verdict: **ALIGNED**. M212 only swaps the target host (`localhost` → `$HOST`); the rewrite stays demo-local (byte-identical when unset).

## Completeness Gaps

1. **(KB-1, incidental)** clerkenstein.md omits the dotted-FAPI-host constraint. Homed to M214's `tailscale-serve.md` (Fate 2/3 — a future milestone of this release owns the doc). Not load-bearing for M212 (the `127.0.0.1` default is preserved; the constraint is enforced+commented in code).

## Applied Fixes
None applied inline — no doc is stale for M212 (each describes the localhost-default behavior M212 preserves when the knob is unset). The one blind area and the one completeness gap are both homed to M214.

## Open Items (require user decision)
None.

## Gate Result
YELLOW: proceed with tracking. KB-1 recorded in `decisions.md`; the `STACK_PUBLIC_HOST` doc blind area is a declared M214 deliverable (`tailscale-serve.md`).
