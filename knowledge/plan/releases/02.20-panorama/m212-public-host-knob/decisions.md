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
