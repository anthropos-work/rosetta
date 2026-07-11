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
