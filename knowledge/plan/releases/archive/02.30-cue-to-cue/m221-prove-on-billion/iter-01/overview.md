---
iter: 1
milestone: M221
iteration_type: tok
tok_flavor: bootstrap
iter_shape: tok
status: closed-fixed
opened: 2026-07-15
closed: 2026-07-15
---

# iter-01 (bootstrap tok) — the strategy to prove v2.3 on billion

**Type:** tok (bootstrap). Authors TOK-01, the strategy the gate-proving tiks run under. No gate progress
directly; its deliverable is the strategy + the ordering that keeps the proof from corrupting its own evidence.

## Inputs
- `overview.md` (the 8-condition exit gate + the inherited carries from M217/M218/M219/M220)
- `iteration_protocol_ref`: `corpus/ops/verification.md` + coverage-protocol.md + playthroughs.md, run **from a
  remote origin** (a second tailnet machine driving `billion`)
- The KB-fidelity gate (iter-01, Phase 0b): **YELLOW**, no blind areas; every load-bearing carry verified against
  `cue-to-cue-m220-r6` (the code M221 runs). Report: `../kb-fidelity-audit.md`.
- The direct analogue: **M215 "prove-on-odyssey"** (7.1 h, direct-drive) — the reconfiguration is fully specified
  upstream; the last breakages only surface on a live cross-machine run. Expect the same **direct-drive** shape.

## The initial strategy (see TOK-01 in ../decisions.md for the canonical record)

**Fix-then-prove, in a dependency order that the overview makes non-negotiable — and under the operating rules
this session paid for in hours of churn.**

### The ordering constraint (load-bearing)
The overview states it twice: **`GUARD-M221-host-isolation` lands FIRST.** M221's own gate is a multi-cycle cold
battery on the *single* `billion` host. M219 proved that two agents on one host corrupt the evidence — a cycle
purged the stack mid-measurement and a gate went UNEXECUTED (a FINDING, not a pass). Without a host lock, this
milestone can corrupt its own proof exactly as M219 did. So the guard is a **prerequisite for the battery**, not
a parallel item.

### The three phases (tiks run under this)
- **Phase A — the host-isolation guard (off-box, lands first).** A stale-tolerant lockfile on the demo host,
  taken for the life of a cycle, named with the holder's identity; a second concurrent cycle **fails loud** with
  the holder's identity — never queues, never proceeds. Fence: RED-proven that a second cycle is refused.
- **Phase B — the off-box code fixes** (no live demo needed to *write* them; the battery *proves* them):
  `FIX-M221-academy-loopback-bind` (`-H 127.0.0.1` + extend `exposure_claim_guard` to host-native listeners —
  the fence is blind to the class that bit it twice); `FIX-M221-reap-native-academy` (reap the native
  cockpit+academy by port, same `reap.sh` surface); `FIX-M221-academy-empty-catalog` (the local-catalog pipeline
  emits 0 — likely a new sha-pinned demo-patch); `PROBE-M218-backend-api-url-twin` (a fence so a server-side
  reader of the blackholed backend URL can never silently appear).
- **Phase C — the live cold-reset-to-seed battery on `billion`, driven from a peer.** A DEFAULT `/demo-up N` (no
  flags), proving all 8 gate conditions + folding in M219's five readiness gates re-proven **at final code**
  (`REPROVE-M221-battery-at-final-code`). Also exercises the never-field-run M217 deliverables (pre-bind reap,
  `assert_ports_free`, freshness preflight), the F12 teardown-serve-reset, and settles the 7.3 GiB RAM question
  (C-6) by measurement before blaming code. Plus `BURNIN-M221-dev-public-host` (one live remote dev stack + one
  no-flag proof of zero tailscale calls on a tailscale-capable box).

### The operating rules (paid for in M218/M219 churn — non-negotiable)
- **ONE agent against `billion` at a time.** This is *why* Phase A exists; it also applies to the humans/agents
  driving it. Two concurrent things on the host corrupted M219's audit trail.
- **Drive bring-ups SYNCHRONOUSLY, foreground, bounded polls. NO detached/background scripts on the host.** Every
  orphan left running wiped a stack mid-measurement.
- **NEVER kill a build mid-bake** — it strands the demopatches and the next image ships silently unpatched. If
  ever killed, restore the demo clone to pristine and verify 0-dirty.
- **Assert from a tailnet peer, never on-host** (an on-host run fakes a 21-section RED via an SSL artifact).
- **An empty / absent / unexecuted result is a FINDING, not a pass** (D17 — ~20 hits and counting).
- **Every fence RED-proven pre-fix.** A fence that passes against the bug is theatre — it happened repeatedly.

## Distance to gate
Gate = the 8 conditions on a cold reset-to-seed with NO FLAGS. Starting distance: **unmeasured on this code** —
the last real `billion` run (M215-era) skipped all three catalog surfaces and predates M218/M219/M220. Phase C's
first battery cycle is the baseline.

## Next-tik direction
iter-02 (first tik) = **Phase A: land `GUARD-M221-host-isolation`** — the prerequisite. Off-box; RED-prove that a
second concurrent cycle is refused with the holder's identity.
