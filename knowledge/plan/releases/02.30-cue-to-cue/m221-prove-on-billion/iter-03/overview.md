---
iter: 3
milestone: M221
iteration_type: tik
iter_shape: tik
status: closed-fixed
opened: 2026-07-15
closed: 2026-07-15
strategy_ref: TOK-01 (../decisions.md)
---

# iter-03 (tik) — Phase B: the off-box demo-hygiene cluster

**Type:** tik, under TOK-01 ("fix-then-prove, host-isolation FIRST"). Phase A (iter-02) landed the host lock.
This tik lands **Phase B** — three tightly-coupled **off-box** fixes so that Phase C's live cold-reset-to-seed
battery on `billion` grades the **SHIPPED** code, not a graded-then-changed artifact (M219's `REPROVE` /
graded≠shipped lesson). All three are **fence-able off-box**; the live proof is Phase C's job.

## The three fixes (all inherited carries)

### (1) `FIX-M221-academy-loopback-bind` (F-M220-5) — the S0 lie one layer up
`ant-academy.sh` passed **no `-H`** on the localhost path, and `next dev`'s **own** default bind is `0.0.0.0`
— so the academy was world-published on `*:13077` on **every** localhost demo, answering HTTP 200 on the
tailnet IP. It survived S0/S1 because `exposure_claim_guard` only ever knew the three **container** port
emitters — *an exposure fence blind to a whole class of listener reports a confident, quietly incomplete pass*
(D17).

- **Code:** on the empty-`BIND_HOST` (localhost) path, pass **`-H 127.0.0.1`**; the public path
  (`BIND_HOST=0.0.0.0`, deliberately world-facing behind the tailnet) is untouched. Surgical — the fenced
  M220 S3 HARD INVARIANT is on **`up-injected.sh`'s derivation trio** (`BIND_HOST`/`HOST`/`SCHEME`) + no
  tailscale probe, NOT on the academy's literal `-H`; the trio stays `BIND_HOST=""` on localhost. The only
  change is the academy's bind interface **tightening** `0.0.0.0`→`127.0.0.1` — a strict de-exposure.
- **Guard:** extend `exposure_claim_guard` to the **host-native** listeners (cockpit + academy). It must
  derive their **effective** bind on a localhost demo — and the whole reason the bug hid is that the effective
  bind is the *tool default*, not what the script writes: `next dev` defaults to `0.0.0.0`, `cockpit.py`'s
  argparse `--host` defaults to `127.0.0.1`. So the guard reads the launch args AND the tool defaults, and
  asserts both host-native listeners bind **loopback** on a localhost demo.

### (2) `FIX-M221-reap-native-academy` — a stale native academy across cycles
The academy is a host-native `next dev` (not a container), so `down --purge` cannot compose-reach it. The
reap machinery already exists (M217 `reap_native_ports` covers cockpit **and** academy by port; M220 S5(i)
added the `ant-academy.sh` **pre-bind** reap). **But the pre-bind wiring was UNFENCED** — an unfenced fix is
exactly the D17 shape that regresses silently. This tik lands the fence: the pre-bind reap is wired **before
the launch**, is **port-scoped** (a co-resident demo-M academy on a different offset port is never touched),
and reclaims a stale academy-identity listener.

### (3) `PROBE-M218-backend-api-url-twin` (F-7) — fence the loaded gun
`NEXT_PUBLIC_BACKEND_API_URL` bakes to the public-host origin (`up-injected.sh:697`) and blackholes from
inside the container (measured 10,553 ms → `UND_ERR_CONNECT_TIMEOUT`) — the C-1 shape that cost M218 37.5 s
per render. It is dormant **only** because all readers are client-side (M218 D10). **DoD = the FENCE**, not a
live re-measure (that's Phase C): a guard that fails loud if a **server-side** reader of the var ever appears
(an app-router server component without `'use client'`, a `route.ts` handler, `getServerSideProps`,
middleware, or a `server-only` module). Modelled on M218's `WUNDERGRAPH_SSR_ENDPOINT` fence
(`test_ssr_origin_chain.py`).

## The fences (RED-proven — the release's core discipline)
Every fence FAILS against pre-fix / mutant code and PASSES after. Homed in the owning rext section
(`demo-stack/tests/`, `stack-injection/tests/`). Mutation-proved where practical.

## Out of scope (this tik)
Phase C's live `billion` battery; `PROBE-M218-c3-rerun` (needs the box); the academy empty-catalog patch
(`FIX-M221-academy-empty-catalog`, a separate content-pipeline defect → its own tik); the dev `--public-host`
burn-in. This tik does **not** require a live cycle.

## Distance to gate
The 8-condition milestone gate is unmeasured on this code; this tik does not move it directly. Its sub-gate:
the three fences RED-proven pre-fix and GREEN after — so Phase C grades the shipped code.
