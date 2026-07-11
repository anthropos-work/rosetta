---
milestone: M214
slug: origins-and-links
version: v2.2 "panorama"
milestone_shape: section
status: archived
created: 2026-07-11
last_updated: 2026-07-11
complexity: medium
depends_on: M212
delivers: NEW corpus/ops/demo/tailscale-serve.md (the remote-access recipe) + the MagicDNS origin admitted everywhere a browser call is gated + the bounded patch tail landed via the existing rext patch mechanism
issues: backend CORS_EXTRA_ORIGINS is localhost-only; ant-academy allowedDevOrigins hardcodes a DIFFERENT tailnet host (live next dev); next-web urls.ts WEB_APP_URL/HIRING_APP_URL have no override (prod-eject risk)
---

# M214 — Origins & links

## Goal
Admit the MagicDNS/HTTPS origin **everywhere** a browser→backend or cross-surface call is gated, **close the
cross-surface link ejects**, and land the **bounded patch tail** via the **existing** rext patch mechanism (never a
raw clone edit) — then author the remote-access recipe so the flow is documented.

## Why section
Enumerable: extend one CORS emission, fix two runtime redirect envs, add one confirmed patch + one conditional
patch, author one doc. Every item is file:line-mapped (`spec-notes.md`).

## Scope
- **In (config — rext tooling):**
  - **CORS:** extend `gen_injected_override.py:304-306` so `fe_origins` also emits the HTTPS MagicDNS origin(s)
    at the offset ports (keep the `localhost` entries for on-host use); update the injection tests that assert
    localhost-only (`tests/test_injection.py:382-394,511-515`). The backend honors `CORS_EXTRA_ORIGINS` at runtime
    in non-production (`cors.go:24,78-82`; backend runs `ENVIRONMENT=development`).
  - **studio-desk redirects:** `$HOST` into runtime `CLERK_SIGN_IN_URL`/`WEB_APP_URL`
    (`gen_injected_override.py:232-233`); resolve the `VITE_CLERK_SIGN_IN_URL` bake gap (declare the ARG or a
    build-context `.env` so it isn't the un-offset `localhost:3000/login` default).
- **In (patch — via the sha-pinned `apply-*.sh` / demopatch surface, NOT clone edits):**
  - **ant-academy `allowedDevOrigins` (REQUIRED):** author a NEW `apply-*.sh` patch (mirroring
    `apply-app-authz-skip.sh`: path+anchor+`pre_sha256`/`post_sha256`+marker, drift-refuse, idempotent, non-fatal)
    adding the MagicDNS host to `ant-academy/code/next.config.js:9` (currently hardcoded to a *different*
    `taildc510` host + IP; live `next dev` blocks cross-origin dev requests otherwise).
  - **next-web `urls.ts` WEB_APP_URL / HIRING_APP_URL (CONDITIONAL):** only if target demo flows traverse those
    links — add a demopatch mirroring `next-web-studio-url`; else record as documented residual (they prod-eject).
  - Confirm the two **already-shipped** demopatches (`next-web-studio-url`, `next-web-public-website-url`) carry the
    MagicDNS baked value cleanly.
  - **Mixed-content check:** with M213's HTTPS-everywhere origin, verify no browser-facing call resolves to plain
    `http://` (asset plane stays prod HTTPS — no change; `next.config` `remotePatterns` already allowlists it).
- **Out:** the knob (M212); TLS/proxy/pk (M213); the live cross-machine acceptance (M215).

## Depends / parallel
- **Depends on:** M212 (the `$HOST` plumbing in `gen_injected_override.py`).
- **Parallel with:** **M213** — disjoint files (CORS/patches vs cert/proxy/pk), additive merge.

## Delivers → knowledge
- **NEW `corpus/ops/demo/tailscale-serve.md`** — the remote-access recipe: the opt-in `--public-host` flag, the
  HTTPS-everywhere topology, the tailscale-cert FAPI, the CORS + patch tail, the "teammate on the tailnet browses
  it" walkthrough, and the safety framing (Tailscale = the access control; opt-in default-off).
- Updates: `corpus/ops/rosetta_demo.md`, `corpus/ops/demo/frontend-tier.md`, `corpus/services/clerkenstein.md`.

## Open questions
- Exact origin string the browser uses under HTTPS-everywhere (one `https://<magicdns>` vs per-port) — must match
  `CORS_EXTRA_ORIGINS` + `allowedDevOrigins` byte-for-byte. Settled by M213's proxy topology; verified in M215.
- Do target demo flows actually traverse `app.`/`hiring.` links (decides whether the conditional `urls.ts` patch is
  needed)?

## KB dependencies
`corpus/ops/safety.md` (the tenant-firewall + values-blind framing), `corpus/ops/rosetta_demo.md`,
`corpus/ops/demo/frontend-tier.md`, `corpus/services/next-web-app.md`, the rext `stack-injection/apply-*.sh`
patch-mechanism precedent.
