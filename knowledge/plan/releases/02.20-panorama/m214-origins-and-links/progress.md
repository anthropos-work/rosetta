# M214 — progress

Section checklist (closure = a MagicDNS-origin browser is admitted by CORS, cross-surface links resolve, the
required patch lands via the sha-pinned mechanism, and the recipe doc exists).

- [ ] `CORS_EXTRA_ORIGINS` emission includes the HTTPS MagicDNS origin(s) + injection tests updated
- [ ] studio-desk runtime `CLERK_SIGN_IN_URL`/`WEB_APP_URL` host-substituted
- [ ] `VITE_CLERK_SIGN_IN_URL` bake gap resolved
- [ ] ant-academy `allowedDevOrigins` NEW `apply-*.sh` patch (sha-pinned, drift-refuse, idempotent, non-fatal)
- [ ] next-web `urls.ts` WEB_APP_URL/HIRING_APP_URL — conditional demopatch or documented residual
- [ ] shipped demopatches carry the MagicDNS value
- [ ] mixed-content check (no browser-facing http under HTTPS-everywhere)
- [ ] NEW `corpus/ops/demo/tailscale-serve.md` + cross-ref updates
