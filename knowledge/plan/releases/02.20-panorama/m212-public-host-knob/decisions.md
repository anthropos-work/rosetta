# M212 — decisions

_(implementation choices with rationale, recorded as the milestone is built)_

## D-DESIGN-1 — opt-in, default off
`STACK_PUBLIC_HOST` defaults to `localhost`; external reach requires an explicit `/demo-up --public-host` flag.
**Why:** user directive (2026-07-11) — public access must be explicitly requested at build time, never ambient.

## D-IMPL-1 — the FAPI host keeps a SEPARATE dotted default (`127.0.0.1`), not `localhost`
The overview §5 shorthand says pass `--fapi-host "$HOST:…"`, but `HOST` defaults to `localhost` and the pk mint
**rejects** a dotless host (`@clerk/backend`'s `assertValidPublishableKey`; up-injected.sh:604-606). The byte-identity
contract (progress.md closure: "unset ⇒ byte-identical to today", and today the FAPI host is `127.0.0.1`) is the hard
gate, so the two can't both be literal. **Resolution:** a second derived var `FAPI_HOST="${STACK_PUBLIC_HOST:-127.0.0.1}"`.
Unset ⇒ `127.0.0.1` (byte-identical); a set MagicDNS host is dotted → flows through `HOST` and `FAPI_HOST` identically.
Used at the pk-mint call + the cockpit fapi-host. **Why:** preserve byte-identity AND keep the pk valid — the only
site where the knob's unset default legitimately differs from `HOST`.

## D-IMPL-2 — external *binding* (0.0.0.0) is gated on the knob, not unconditional
The overview §7 reads "cockpit: launch with `--host 0.0.0.0`", but binding all interfaces unconditionally would make
every demo's cockpit + academy LAN-reachable — a default-on external listener, which D-DESIGN-1 forbids. **Resolution:**
`BIND_HOST=0.0.0.0` **only** when `STACK_PUBLIC_HOST` is set; else omit `--host`/`-H` so cockpit.py's `127.0.0.1` default
+ `next dev`'s own default stand (byte-identical). External reach is exactly what `--public-host` requests, so binding
0.0.0.0 there is intended; without it, loopback-only (today). cockpit.py needed **no** change — its `--host` flag
already exists + is used at the bind.

## D-IMPL-3 — gen_injected_override host is a wired-but-unemitted M214 seam
§6 threads `host` through `build_lines`/`frontend_lines` + a `--public-host` flag, but the CORS + Clerk-URL **emission**
(the only host-bearing sites) stays `localhost`. That emission is **Fate-2, owned by M214** — and depends on M213's
HTTPS/reverse-proxy shape (a proxied origin is `https://<host>` with **no** offset port, so emitting `http://<host>:<port>`
now would just be undone). So M212 wires the seam end-to-end (up-injected → gen → build_lines) and M214 flips only the
emission. With `host` set the generated override is **byte-identical** to the default (a pinned test).

## D-IMPL-4 — registry recording is opt-in + reconcile-then-upsert
up-injected.sh has **no** registry interaction today (it relies on the registry's `docker ps` adoption). To surface the
reachable URL in `/stack-list` without changing the default path, the new `set-host` call runs **only** when the knob is
set (byte-identical when unset) and is **non-fatal**. `set_host()` reconciles-then-upserts so it adopts the
just-started `demo-N` project (or creates a minimal reserved row) rather than silently dropping the host.

## KB-1 — clerkenstein.md omits the dotted-FAPI-host constraint (Phase 0b, YELLOW)
`corpus/services/clerkenstein.md` documents the pk minter (`mintpk`, alignment-gated codec) but not the
constraint that the FAPI host must be **dotted** — `@clerk/backend`'s `assertValidPublishableKey` rejects a
dotless `localhost`, which is why the FAPI host defaults to `127.0.0.1` (not `localhost`) and why a **MagicDNS**
host works for the pk. Thoroughly commented in code (`up-injected.sh:581-586`, `inject.py:50`). **Fate 2/3 —
homed to M214** (`corpus/ops/demo/tailscale-serve.md`, the release's declared KB anchor): explaining why a
MagicDNS origin validates for the pk is exactly that doc's remit. Non-blocking for M212 (the `127.0.0.1` default
is preserved byte-identically). Full audit: `kb-fidelity-audit.md`.

## Adversarial review (Phase 2c — close, 2026-07-11)
_Scenarios considered against the milestone-touched modules. Each is handled, pinned by a test, or a documented
input-contract with downstream validation. Recorded per the close-milestone Phase 2c contract (the scenario, not
just the fix). No unhandled defect surfaced; no code change required._

1. **up-injected.sh flag parse — flag-before-N / unknown-arg / missing-value.** `up-injected.sh --public-host X`
   sets `N="--public-host"`, which fails the `^[0-9]+$` integer check (exit 1); an unknown argument hits the `*)`
   reject arm; `--public-host` with no value trips the `${2:?}` guard. All three are pinned by
   `test_frontend_build.py` (harden Pass 2). Fails loudly, never silently mis-parses.
2. **Empty-string exported `STACK_PUBLIC_HOST` precedence.** `export STACK_PUBLIC_HOST="${STACK_PUBLIC_HOST:-}"`
   always exports (empty when unset). Every downstream gate uses `:-` / `-n` (`HOST="${STACK_PUBLIC_HOST:-localhost}"`,
   `FAPI_HOST="${STACK_PUBLIC_HOST:-127.0.0.1}"`, `[ -n "${STACK_PUBLIC_HOST:-}" ]`), so an empty export is
   byte-identical to unset — localhost host, 127.0.0.1 FAPI, no external bind, no registry write. Holds in the child
   `ant-academy.sh` too (it re-derives from the exported var).
3. **The pk round-trip is symmetric — a host with an embedded port would pass the self-check.** `mint_pk`/`parse_pk`
   round-trip ANY string, so `FAPI_HOST=host:port` (a contract violation — the knob is a BARE host) mints a pk that
   round-trips OK yet fails downstream at `@clerk/backend`'s dotted-host `assertValidPublishableKey`. Mitigations:
   the knob's documented contract is a bare MagicDNS host (`billion.taildc510.ts.net`); `mint_pk` hard-rejects a `$`
   (the empty-`PK_DEMO` guard then aborts the bring-up loudly); STRICT host-format/pk validation is **M213's**
   declared scope (its `assertValidPublishableKey` dotted-host gate + base64 round-trip). Not an M212 regression —
   identical pre-knob behaviour for a malformed `--fapi-host`.
4. **`want_ep` cache-validator staleness on a HOST change.** The build-arg cache-validator must invalidate on a HOST
   change, not only an OFFSET change (else a stale localhost-baked image is silently reused on a `--public-host`
   stack). `want_ep` embeds `$HOST`; pinned by `test_frontend_build.py` rebuild-on-host-change / reuse-on-host-match
   for BOTH next-web and studio-desk (harden Pass 1 — the overview's top-risk item).

## D-CLOSE-1 — the demo-stack README `test_tooling` count-drift is a rext-frozen residual, routed to v2.2 close-release
`demo-stack/README.md:66` quotes `test_tooling.py (50 tests)`; ground truth is **111**. This is pre-existing drift
accumulated across prior releases — M212 modified ONE assertion in `test_tooling.py` (the `$HOST`/`$FAPI_HOST`
parametrization) and added NO tests, so it is **not** an M212 regression. The Phase 4 step-6 reconciliation would
normally fix a count drift in-place, but the file lives in the **frozen rext tag `panorama-m212` @ `770f81b`**,
which this rosetta-only M212 close must not advance or re-tag (the orchestrator reserved the rext re-tag + the
box-level `.agentspace/rext.tag` bump for `/developer-kit:close-release` when the whole v2.2 release ships; M212 is
not the last milestone). **Fate 2 — routed to v2.2 `/developer-kit:close-release`**, which legitimately re-tags /
advances rext and runs the release-level knowledge/doc-hygiene pass (its Phase 3b consolidation). Surfaced in the
completeness ledger + retro so it is not silently elided. Stays in-release (not an escape-hatch); not a repeat or
aged-out deferral.
