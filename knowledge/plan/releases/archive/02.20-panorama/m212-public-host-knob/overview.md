---
milestone: M212
slug: public-host-knob
version: v2.2 "panorama"
milestone_shape: section
status: archived
created: 2026-07-11
last_updated: 2026-07-11
complexity: medium
depends_on: none
delivers: the STACK_PUBLIC_HOST knob + the /demo-up opt-in flag, threaded through every rext browser-facing emitter (the spec-notes seed for the NEW corpus/ops/demo/tailscale-serve.md, authored in M214)
issues: localhost/127.0.0.1 is hardcoded and SCATTERED across rext emitters (up-injected.sh ~23×, gen_injected_override.py ~8×, cockpit.py ~11×, inject.py ~3×); there is no single host variable today
---

# M212 — The single host knob

## Goal
Introduce **ONE explicit, opt-in browser-facing host knob** — `STACK_PUBLIC_HOST` (default `localhost` → the stack
is **byte-identical to today when unset**) — surfaced as an explicit `/demo-up` flag, and thread it through **every
rext emitter** that currently bakes `localhost`/`127.0.0.1` into a browser-facing value. Scoped **strictly to
browser-facing values** — host-loopback control-plane calls (set-dress DSNs, `directus_addr`) stay `localhost`.

## Why section
The deliverables are enumerable up front: the substitution surface is a **known, file:line-mapped list** (see
`spec-notes.md`). The work is "add a `HOST` var + a flag, replace localhost with `$HOST` at N documented sites,
thread it into 4 emitters." No exploratory path — the exit is verifiable by a dry `up-injected.sh` diff.

## Design rule — OPT-IN, default off (user directive, 2026-07-11)
Public/external access is **never default-on**. The knob defaults to `localhost`; external reach must be
**explicitly requested at build time** — passed as a flag to `/demo-up` (and the underlying scripts), e.g.
`/demo-up N --public-host billion.taildc510.ts.net`. Unset ⇒ current behaviour exactly (the regression contract
M215 asserts).

## Scope
- **In:**
  - `up-injected.sh`: read `HOST="${STACK_PUBLIC_HOST:-localhost}"` right after `OFFSET` is computed (`:44`), then
    substitute `$HOST` for `localhost` in the frontend build-args (next-web `:264-266`; studio-desk `:293,:295`),
    the `.env.local` overlays (`STUDIO_URL :185`, `ACADEMY_URL :192`, `PUBLIC_WEBSITE_URL :198`), and the
    post-replay Directus `demo_web` content-URL rewrite (`:876,:879-881`).
  - **Cache-validity fix:** the `want_ep` build-arg cache-validators (`:165,:275`) must invalidate on a **HOST**
    change, not just an offset change — else a stale `localhost`-baked image gets reused. (Top-risk item.)
  - `inject.py`: pass `--fapi-host "$HOST:$((5400+OFFSET))"` (the pk mint is a pure fn of the host; the
    `parse_pk==host` self-check fails loudly on a bad round-trip). *(Cert + validator semantics are M213.)*
  - `gen_injected_override.py`: thread a `host` param (default `'localhost'`) through `build_lines`/`frontend_lines`
    — **plumbing only** here; the CORS + Clerk-URL *emission* using it is M214.
  - `cockpit`: launch with `--host 0.0.0.0` (fixes the ONE `127.0.0.1` listener, `cockpit.py:534` /
    `up-injected.sh:927-943`) + pass `$HOST` into `--app-base`/`--fapi-host`/`--academy-base`.
  - `ant-academy.sh`: substitute `$HOST` in `NEXT_PUBLIC_STUDIO_URL` (`:122`); confirm `next dev` binds `0.0.0.0`
    (add `-H 0.0.0.0` if not).
  - `stack_registry.py`: additive `external_host` field on the record (`allocate()`/set-ports) so `/stack-list`
    shows the reachable URL and verify/teardown can reuse it.
  - **Operator surface:** the `--public-host` flag on `/demo-up` → `STACK_PUBLIC_HOST`, plumbed to the scripts.
- **Out:** the TLS cert swap + the HTTPS reverse proxy + pk validation (M213); the CORS emission + the patch tail +
  the corpus recipe (M214); the live cross-machine acceptance (M215); dev-path parity + `/dev-up` flag (optional
  M216).

## Depends / parallel
- **Depends on:** none (opens the release).
- **Parallel with:** none — M213 + M214 both consume this knob.

## Delivers → knowledge
No new doc lands here; the substitution surface is captured in `spec-notes.md` as the seed for the M214-authored
`corpus/ops/demo/tailscale-serve.md`.

## Open questions
- Single stack per VM (offset 0 / clean base ports) or multiple (offset in every URL)? Default assumption: **one
  demo per VM**; the knob works either way, confirmed live in M215.
- Does `next dev` (ant-academy, Next 16) bind `0.0.0.0` by default on the target VM, or need `-H 0.0.0.0`?

## KB dependencies
`corpus/ops/rosetta_demo.md` (bring-up + offset), `corpus/ops/demo/frontend-tier.md` (the demo UI build),
`corpus/services/clerkenstein.md` (pk/FAPI), `corpus/ops/snapshot-spec.md` (the `demo_web` rewrite context).
