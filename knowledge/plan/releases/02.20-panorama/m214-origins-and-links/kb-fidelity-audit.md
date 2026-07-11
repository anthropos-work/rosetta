---
title: "KB Fidelity Audit — M214 origins-and-links"
date: 2026-07-11
scope: milestone:M214
invoked-by: build-milestone
---

## Verdict
YELLOW

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| CORS `CORS_EXTRA_ORIGINS` emission | `corpus/ops/demo/frontend-tier.md` §CORS (202-214) | `stack-injection/gen_injected_override.py:307-312`; `app/internal/cors/cors.go:24,66-82` | PAIRED |
| studio-desk runtime redirects (`CLERK_SIGN_IN_URL`/`WEB_APP_URL`) | `frontend-tier.md` (62); `clerkenstein.md` | `gen_injected_override.py:225-236` | PAIRED |
| `VITE_CLERK_SIGN_IN_URL` bake gap | `frontend-tier.md` (studio-desk build) | `up-injected.sh:317-345`; studio-desk `Dockerfile.dev`, `app/services/config.ts:3`, `vite.config.ts` | PAIRED |
| ant-academy `allowedDevOrigins` patch (via apply-*.sh) | patch-mechanism precedent `apply-app-authz-skip.sh`; `frontend-tier.md` (176); `coverage-protocol.md` (200) | `ant-academy/code/next.config.js:9`; `ant-academy.sh`; `demo-stack/patches/manifest_loader.py` | PAIRED |
| next-web `urls.ts` app./hiring. ejects | `coverage-protocol.md` (200-205, the fix-surface routing table + prod-eject gate) | `next-web-app/packages/core-js/src/constants/urls.ts:4-11` | PAIRED |
| tailscale-serve remote-access recipe | `corpus/ops/demo/tailscale-serve.md` (NEW) | `gen_tailscale_serve.py` (M213); `up-injected.sh` public-host tail | BLIND-AREA (declared `Delivers →`) |
| safety framing (opt-in, Tailscale = access control) | `corpus/ops/safety.md` | (tooling policy) | PAIRED |

## Fidelity Findings
1. **CORS example reflects current localhost-only code (ALIGNED, not stale).** `frontend-tier.md:207-208` shows `CORS_EXTRA_ORIGINS=http://localhost:13000,…` — byte-accurate to `gen_injected_override.py:311` today. The header comment "# each entry carries its own scheme+host" already anticipates scheme variation. M214 extends BOTH code and this example together (declared doc deliverable). Not misleading to the implementer (code is the truth).
2. **studio-desk requireAuth-fallback claim ALIGNED.** `frontend-tier.md:62` states the override pins `CLERK_SIGN_IN_URL`/`WEB_APP_URL` at the offset next-web — matches `gen_injected_override.py:235-236`. M214 makes the emission host/scheme-aware; the doc claim stays true (still "the offset next-web", now over https for a public host).
3. **`cors.go` honors `CORS_EXTRA_ORIGINS` in non-production ALIGNED.** Verified `cors.go:78-82` (`!environment.IsProduction()` guard + comma-split `parseExtraOrigins`), backend runs `ENVIRONMENT=development`. The doc's "documented runtime hook, not a code path the demo adds" is accurate.
4. **coverage-protocol prod-eject gate ALIGNED (the item-5 evidence).** `coverage-protocol.md:53,66,94,108,200` document the `0 prod-eject escapes` gate + per-`<a href>` host classification + the JS-constant-behind-`NEXT_PUBLIC_NODE_ENV` "re-scope trigger" rule. This is the authoritative source for the `urls.ts` decision.

## Completeness Gaps
1. **`tailscale-serve.md` (NEW)** — the blind area. Satisfied by the milestone's `overview.md` `Delivers → NEW corpus/ops/demo/tailscale-serve.md` line — authored in M214's Document phase. NOT a blocker.
2. **`rosetta_demo.md`** has no `--public-host` coverage yet (grep empty) — a declared M214 doc update.
3. **`clerkenstein.md`** covers the M213 tailscale cert + dotted-pk (108,123-130) but not yet the CORS origin / `allowedDevOrigins` / per-port-serve origin shape — a declared M214 doc update.
4. **`frontend-tier.md`** CORS example + studio-desk redirect + the new `allowedDevOrigins` patch — declared M214 doc updates.
5. Carry-forward from M213 Phase 0b **KB-4** (frontend-tier "Browser-trusted FAPI cert (M31)" callout describing the pre-M213 mkcert world) — homed to M214's Document phase.

## Applied Fixes
None inline — every gap is a declared M214 doc deliverable authored during the milestone's Document phase (not a pre-existing stale claim to correct now).

## Open Items (require user decision)
None.

## Gate Result
YELLOW: proceed with tracking. No blind area that isn't already a declared `Delivers →` deliverable; no stale load-bearing claim that would mislead the implementer (docs match current code; M214 extends code + docs together). KB items recorded in `decisions.md` (KB-1..KB-3). SEVERITY: warning.
