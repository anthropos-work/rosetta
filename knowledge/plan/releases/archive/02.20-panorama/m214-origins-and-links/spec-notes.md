# M214 — spec-notes

The origins + links surface (from `wf_bea3be47`, file:line).

## CORS (config)
- Backend allowlist hardcoded (`app/internal/cors/cors.go:36-44,66-74`) — no MagicDNS. Runtime hook
  `CORS_EXTRA_ORIGINS` appended when `!IsProduction()` (`:24,:78-82`); backend runs `ENVIRONMENT=development`.
- Injector emits localhost-only today: `gen_injected_override.py:304-306`
  `fe_origins = ",".join(f"http://localhost:{p+offset}" for p in (3000,3001,9000))`. Extend to also emit the HTTPS
  MagicDNS origin(s). Tests assert localhost-only at `tests/test_injection.py:382-394,511-515`.
- Already permissive (no change): Cosmo `allow_origins:["*"]` (`graphql-wundergraph/config.compose.yaml:30-31`);
  Clerkenstein reflects Origin + allow-credentials (`clerk-frontend/server.go:90-113`); studio-desk `cors()` wildcard.

## Cross-surface link ejects
- studio-desk runtime `CLERK_SIGN_IN_URL` / `WEB_APP_URL` localhost (`gen_injected_override.py:232-233`) → `$HOST`.
- `VITE_CLERK_SIGN_IN_URL` bake gap: not a declared Dockerfile ARG; falls back to un-offset `localhost:3000/login`
  (`studio-desk app/services/config.ts:3`). Declare the ARG or bake a build-context `.env`.

## Patch tail (via rext apply-*.sh / demopatch — NOT clone edits)
- **REQUIRED — ant-academy `allowedDevOrigins`:** `ant-academy/code/next.config.js:9`
  `['ithacastaging.taildc510.ts.net','100.120.254.65']` — add the MagicDNS host. Live `next dev` blocks otherwise.
  New `apply-*.sh` mirroring `stack-injection/apply-app-authz-skip.sh:40-88` (path+anchor+pre/post_sha256+marker,
  drift-refuse, idempotent, non-fatal).
- **CONDITIONAL — next-web `urls.ts`:** `core-js/src/constants/urls.ts:4-11` WEB_APP_URL/HIRING_APP_URL — no
  override, not demopatched → prod-eject. New demopatch (mirror `next-web-studio-url`) only if flows use those links.
- **Shipped demopatches (confirm MagicDNS value):** `next-web-studio-url`, `next-web-public-website-url`.

## No change (verified)
- Asset plane stays prod: `DIRECTUS_PUBLIC_BASE_ADDR=https://content.anthropos.work`,
  `MEDIA_URL=https://media.anthropos.work` — allowlisted in `next.config.mjs remotePatterns:37-108`. Loads over
  Tailscale unchanged. `/api/e` same-origin gate (`apps/web/src/app/api/e/route.ts:84-93`) holds under one origin.

## Pre-flight audits — CORS + links + patch tail (M214)
Phase 0b (2026-07-11): **KB-FIDELITY YELLOW** — `kb-fidelity-audit.md`. Topic→doc→code triples verified. No
stale load-bearing claim; the sole blind area (`tailscale-serve.md`) is the declared `Delivers →` deliverable.

Verified code facts (the implementation contract):
- **Origin shape (D-PROXY-2 settled):** M213's proxy fronts each browser-facing PLAINTEXT port with `tailscale
  serve --bg --https=<offsetport> http://127.0.0.1:<offsetport>` — PRESERVING the offset port. So the browser
  origin is `https://$HOST:<offsetport>` (NOT a port-less 443). M214's emission is http→https on the SAME port.
- **CORS:** `gen_injected_override.py:311` emits `http://localhost:{3000,3001,9000}+offset` only. `cors.go:78-82`
  honors `CORS_EXTRA_ORIGINS` (comma-split) when `!IsProduction()`; backend runs `ENVIRONMENT=development`.
- **studio-desk redirects:** `gen_injected_override.py:235-236` pin `CLERK_SIGN_IN_URL`/`WEB_APP_URL` at
  `http://localhost:3000+offset`. `VITE_CLERK_SIGN_IN_URL` bake gap: studio-desk `Dockerfile.dev` declares
  ARGs for the 4 other VITE vars but NOT this one; `config.ts:3` defaults to un-offset `http://localhost:3000/login`.
  **studio-desk `.dockerignore` EXCLUDES `.env*`** (root-level), and `envDir=/app` (context root) — so a naive
  build-context `.env` is dropped; needs a `.dockerignore` re-include (negation) + a `.env.production` overlay
  (or the Dockerfile ARG — but that's a platform edit). next-web's `apps/web/.env.local` works because it's
  NESTED (not matched by the root-only `.env*` in its tooling `.dockerignore`).
- **up-injected.sh browser-facing bakes are ALL `http://$HOST`** (232,239,245,311-313,322,340,342 + the two
  `want_ep` cache validators 212/322 + `demo_web` 1018 + cockpit 1069/1077). Under HTTPS-everywhere these are
  mixed-content → the scheme flip (a single `SCHEME` var: `https` when `STACK_PUBLIC_HOST` set, else `http`;
  byte-identical when unset) is the substance of the mixed-content check + the "demopatches carry the value cleanly".
- **ant-academy:** `next.config.js:9` hardcodes `allowedDevOrigins: ['ithacastaging.taildc510.ts.net',
  '100.120.254.65']` (a DIFFERENT tailnet host); runs `next dev` natively from `stack-demo/ant-academy/code`
  (NOT a Docker image — so the patch hooks at the ant-academy.sh launch, not the inject loop). Pristine
  `next.config.js` sha256 = `6837cab95fad4cf2265734c53ae6a95d0200370f67a75e93729d50387850a8e3`.
- **next-web `urls.ts` (item-5 decision, EVIDENCE):** `WEB_APP_URL`/`HIRING_APP_URL` = `NEXT_PUBLIC_NODE_ENV`
  ternary → prod (`app./hiring.anthropos.work`), no per-URL override. Their `apps/web` usages are PUBLIC
  marketing chrome (`PublicHeader/Footer/BlackFridayBanner` — anonymous only), PDF/SEO metadata (non-nav), the
  Clerk.provider `HOSTING_URL` FALLBACK (the demo BAKES `NEXT_PUBLIC_HOSTING_URL` so the fallback is dead), and
  HIRING-product features (share-sim/invite/start-sim — not a Workforce demo flow). The M42e+M42m coverage
  sweeps gate at **0 prod-ejects** and surfaced only `STUDIO_URL`+`PUBLIC_WEBSITE_URL` (both fixed) — NOT
  `WEB_APP_URL`/`HIRING_APP_URL` → the demo's target flows do NOT traverse them. Under HTTPS-everywhere they'd be
  https-prod (not mixed-content), only a prod-eject on flows the demo never exercises. **Decision: documented
  residual, NOT a new demopatch** (adding one for an unrendered link = speculative; coverage-protocol makes
  "add a demopatch" a re-scope trigger only when a 0-eject sweep surfaces the escape — it hasn't).
