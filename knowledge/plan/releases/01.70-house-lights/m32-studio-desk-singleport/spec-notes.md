# M32 — Spec notes

_Technical detail. Stub at scaffold; seeded from the design._

## Root cause (verified)
The override (`gen_injected_override.py`) correctly publishes single-port `9000:9000` for studio-desk, but the base
platform `docker-compose.yml` ships `NODE_ENV=development` + `FRONTEND_PORT=9100`, and the override's per-frontend env
block is **additive** (deliberately NOT `!override`), so `NODE_ENV=development` survives → `studio-desk/src/index.ts`
`isProduction=false` → the dev block runs `res.redirect('http://localhost:9100/home')` — a dead port (only `9000`→offset
is mapped).

## The fix
Add `NODE_ENV=production` (+ `FRONTEND_PORT=9000`) to the studio-desk dict in `gen_injected_override.py` FRONTENDS
(~lines 90-96). Production → the `sendFile` path (src/index.ts ~211-265), no cross-port redirect.

## Verify the production path covers all routes
The dev block also handled `.html`-extension redirects (~152-204). The Playwright smoke (/home + a couple routes) must
confirm the production `sendFile` path serves them, or some pages 404.

## `:9100` sweep
- demo-up SKILL: `:9100+` → `:9000+`.
- `frontend-tier.md:21`: drop the dead `:9100` frontend port → single-port `9000`+offset.
- `gen_injected_override.py:249`: CORS — remove the un-offset `9100` origin (explicit decision note; dead now).

## Regression test
`stack-injection/tests/test_injection.py` near the single-port assertions (~820-857): assert `NODE_ENV=production` in the
studio-desk env block.
