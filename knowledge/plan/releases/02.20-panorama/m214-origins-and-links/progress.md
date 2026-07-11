# M214 — progress

Section checklist (closure = a MagicDNS-origin browser is admitted by CORS, cross-surface links resolve, the
required patch lands via the sha-pinned mechanism, and the recipe doc exists).

- [x] `CORS_EXTRA_ORIGINS` emission includes the HTTPS MagicDNS origin(s) + injection tests updated
- [x] studio-desk runtime `CLERK_SIGN_IN_URL`/`WEB_APP_URL` host-substituted (https for a public host)
- [x] `VITE_CLERK_SIGN_IN_URL` bake gap resolved (gitignored `.env.production.local` overlay; no Dockerfile ARG)
- [x] ant-academy `allowedDevOrigins` NEW `apply-*.sh` patch (sha-pinned, drift-refuse, idempotent, non-fatal)
- [x] next-web `urls.ts` WEB_APP_URL/HIRING_APP_URL — **documented residual** (evidence: 0-eject sweeps never surfaced them; D-URLS-1)
- [x] shipped demopatches carry the MagicDNS value (the `$SCHEME`/`$HOST` flip in `up-injected.sh`)
- [x] mixed-content check (no browser-facing http under HTTPS-everywhere — the scheme flip covers every surface)
- [x] NEW `corpus/ops/demo/tailscale-serve.md` + cross-ref updates (rosetta_demo, frontend-tier, clerkenstein, demo/README index)

**All sections landed.** rext code + tests: `panorama-m214` (commits `bf3edd1` CORS+redirects, `ca4cb0b`
scheme-flip + VITE bake, `4599a2d` ant-academy patch). rosetta docs + plan on `m214/origins-and-links`.
