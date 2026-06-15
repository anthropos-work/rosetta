# M32 — Decisions

_Implementation decisions with rationale, numbered `M32-D1`, `M32-D2`, … . Empty at scaffold; filled during build._

_Pre-decided at design (2026-06-15):_
- _Root-cause fix is `NODE_ENV=production` on the studio-desk override (the additive env block lets the base
  `development` survive → the dev redirect to dead `:9100`). Production serves via `sendFile`, no cross-port redirect._
- _Remove the un-offset `:9100` CORS origin (dead once studio-desk is single-port production) — record as an explicit decision._

## M32-D1 — Route coverage verified: production path covers all dev-block routes (the load-bearing open question)
**Verified by code-read** of `stack-demo/studio-desk/src/index.ts` (the cloned, byte-pristine platform copy) before
trusting the `NODE_ENV=production` flip.

- **Dev block** (`isDevelopment`, lines 148-206): `/` → `/home`; `/home` → `/home.html`; `/catalog`; `/simulation-builder`
  → `/sim-advanced-builder`; `/sim-advanced-builder`; `/sim-guided-builder`; `/builder-skill-path`; `/generation`;
  `/skills`; plus `app.get('*')` forwarding `${req.path}` to the **Vite dev server** on `frontendPort`.
- **Production block** (else, lines 207-272): `/` + `/home` → `sendFile home.html`; `/simulation-builder`(+`.html`) →
  `/sim-advanced-builder`; `/sim-advanced-builder`; `/sim-guided-builder`; `/builder-skill-path`; `/simulations`;
  `/generation`; `/catalog`; `/academy`; `/skills` (all `sendFile`); **`express.static('/', dist/public)`**; `app.get('*')`
  → `sendFile index.html` (SPA fallback).
- **Verdict: NO GAP.** Every dev route has a production equivalent. The dev block's `.html`-extension *targets*
  (`/home.html`, `/catalog.html`, …) live in `dist/public` and are served by the production `express.static` mount, so they
  don't 404. The dev `*` catch-all only ever bounced to the Vite dev server (a process that doesn't exist in the
  container — exactly the dead-port symptom); its production analog (the `index.html` SPA fallback) is strictly better.
  → Flipping `NODE_ENV=production` is safe; no route silently 404s. (The close-time live Playwright smoke confirms it
  on a running container.)

## M32-D2 — Remove the un-offset `:9100` CORS origin
`gen_injected_override.py` backend CORS (line ~249) enumerated `(3000, 3001, 9000, 9100)`. The `9100` origin is **dead**
now studio-desk is single-port production (the browser only ever talks to `9000+offset`; the `9100` frontend port exists
only under `npm run dev`, never in the container). Dropping it to `(3000, 3001, 9000)` removes a no-op allowlist entry —
no behavior change, just truth. The regression test (`test_backend_gets_cors_extra_origins_at_offset`) loses its `19100`
assertion accordingly.

## M32-D3 — Fold the env-masked stale test assertion into the sweep
`test_injection.py:925` (`test_frontend_blocks_parse_to_valid_compose`) asserted studio-desk ports
`["29000:9000", "29100:9100"]` while the generator emits single-port `["29000:9000"]` only. The test is **skipped** when
PyYAML is absent (this env), which masked the inconsistency. Corrected to `["29000:9000"]` as part of the `:9100` sweep so
the assertion matches the generator whenever PyYAML IS present.
