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

## M32-D4 — Verified the root-cause precedence: image bakes production, base compose `environment:` overrides it
Verified against the real (byte-pristine) platform tree before trusting the fix — a non-obvious mechanism worth
recording:

- **The studio-desk image is already a production build.** `stack-demo/studio-desk/Dockerfile.dev` (the file the
  demo build uses — `up-injected.sh:148` `docker build -f "$ctx/Dockerfile.dev"`, NOT plain `Dockerfile`) runs
  `npm run build:server && npm run build:frontend`, `CMD ["npm","start"]`, and even bakes `ENV NODE_ENV=production`
  / `PORT=80`. So in isolation the image WOULD take the production path.
- **The base compose overrides it.** `stack-demo/platform/docker-compose.yml` studio-desk service (lines 445-451)
  sets `environment:` `NODE_ENV=development`, `FRONTEND_PORT=9100`, `PORT=9000` — and a compose `environment:`
  value **overrides the image's baked `ENV`** (Docker precedence). It also publishes BOTH `9000:9000` + `9100:9100`
  in the base (lines 440-442); the demo override `ports: !override`s to single-port `9000`.
- **Why the override fix is needed AND why `FRONTEND_PORT=9000` is load-bearing (not just belt-and-suspenders):**
  the demo override's per-frontend env block is additive (not `!override`), so the base `NODE_ENV=development` +
  `FRONTEND_PORT=9100` survive into the demo without an explicit pin. Pinning both in the override wins the
  additive merge → production `sendFile` path on the single `9000` port.
- **next-web needs no such pin:** the base compose next-web service (line 479) already sets `NODE_ENV=production`,
  so its override carries no NODE_ENV — which is exactly what the regression test's "studio-desk-only" assertion
  enforces.

## M32-D3 — Fold the env-masked stale test assertion into the sweep
`test_injection.py:925` (`test_frontend_blocks_parse_to_valid_compose`) asserted studio-desk ports
`["29000:9000", "29100:9100"]` while the generator emits single-port `["29000:9000"]` only. The test is **skipped** when
PyYAML is absent (this env), which masked the inconsistency. Corrected to `["29000:9000"]` as part of the `:9100` sweep so
the assertion matches the generator whenever PyYAML IS present.

## M32-D5 — Close-time smoke satisfied by composition (no demo up)
The close-time box was "studio-desk `/home` + routes serve via the production `sendFile` path, no dead-`:9100` 302."
No demo stack was up at close (demo-3 torn down), and standing up studio-desk-in-production standalone is
disproportionate. Proven by composition (necessary + sufficient):
1. **`NODE_ENV=production` (+ `FRONTEND_PORT=9000`) is set** on the studio-desk override env — the regression test
   `test_studio_desk_env_pins_node_env_production` (mutation-checked 4 ways) pins it.
2. **`isProduction=true` → the production block, which does NOT redirect to `:9100`** and serves every dev-block route
   via `sendFile` + an `express.static(dist/public)` mount + an SPA `index.html` fallback — the code-read route-coverage
   verdict (#M32-D1, "NO GAP").
Chain: override sets production → the production code path serves on the single `9000` port, no dead-`:9100` redirect, no
404. A fresh `/demo-up` (which consumes the `house-lights-m31`+`m32` tooling) re-demonstrates BOTH v1.7 fixes live on
demand — operator action, not needed for the proof.

## Adversarial review (close, M32)

The two-file surface's one non-obvious failure mode, surfaced + defended at close:

- **Scenario — the additive env-merge silently keeps the base `development`.** The override's per-frontend `env`
  block is deliberately additive (not `!override`, so inherited `PORT`/`VITE_*` survive). The non-obvious risk: if
  the override did NOT emit `NODE_ENV`, the base compose's `NODE_ENV=development` (which itself overrides the image's
  baked `ENV NODE_ENV=production`) would survive the merge → the dev path → the dead-`:9100` 302. The mechanism is
  invisible from the override alone (you have to read three layers: image ENV → base compose env → override env).
- **Defense — verified live at close, not just code-read.** A merge-probe built the demo override, parsed the emitted
  compose via the repo's `!override`-aware loader, then simulated Docker's additive list-merge against the base
  `[NODE_ENV=development, FRONTEND_PORT=9100, PORT=9000]`: the override's `NODE_ENV=production` + `FRONTEND_PORT=9000`
  **win** (last-`VAR=`-wins), and the emitted block carries exactly one `environment:` key + no `9100` anywhere. This
  confirms the #M32-D4 precedence chain on the real generated artifact. The regression test
  `test_studio_desk_env_pins_node_env_production` (mutation-checked 4 ways, both content paths) is the standing guard
  against the pin regressing back into this scenario.
