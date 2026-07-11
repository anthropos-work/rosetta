# M212 — spec-notes

The `localhost`/`127.0.0.1` substitution surface (from the `wf_bea3be47` feasibility investigation, file:line).
All sites are in **rext tooling** — zero platform-repo files.

## The knob
- `STACK_PUBLIC_HOST` (default `localhost`). Surfaced as `/demo-up N --public-host <magicdns>`.
- Insert `HOST="${STACK_PUBLIC_HOST:-localhost}"` at `demo-stack/up-injected.sh:44` (right after `OFFSET=$((N*10000))`).

## Browser-facing sites to substitute (`localhost` → `$HOST`)
| Site | File:line | Baked? |
|------|-----------|--------|
| next-web build-args (WUNDERGRAPH/BACKEND_API/HOSTING_URL) | up-injected.sh:264-266 | build-time |
| next-web `.env.local` STUDIO_URL / ACADEMY_URL / PUBLIC_WEBSITE_URL | up-injected.sh:185,192,198 | build-time (ACADEMY via urls.ts native; STUDIO/PUBLIC_WEBSITE via shipped demopatches) |
| studio-desk build-args (VITE_GRAPHQL_ENDPOINT / VITE_WEB_APP_URL) | up-injected.sh:293,295 | build-time |
| pk mint `--fapi-host` | up-injected.sh:590 → inject.py:26-37,60-68 | build-time (pk baked) |
| cockpit `--app-base` / `--fapi-host` / `--academy-base` + `--host` | up-injected.sh:927-943; cockpit.py:534 | runtime |
| ant-academy `NEXT_PUBLIC_STUDIO_URL` | ant-academy.sh:122 (native `next dev` :71,:146) | build-time (native → env+restart) |
| Directus `demo_web` content-URL rewrite | up-injected.sh:876,879-881 | runtime (DB write) |
| `want_ep` cache-validators (must invalidate on HOST) | up-injected.sh:165,275 | — |

## Plumbing (host param, emission deferred to M214)
- `gen_injected_override.py`: thread `host` through `build_lines` (254-407) / `frontend_lines` (202-251). CORS
  (`:304-306`) + Clerk sign-in/web-app URLs (`:232-233`) EMIT in M214.

## DO NOT re-point (host-loopback control plane — not browser-facing)
- `dev-setdress.sh:54` BASE_DSN, `:129` directus_addr; `up-injected.sh:725` setdress_dsn. Leave as `localhost`.

## Registry
- `stack_registry.py:197-202` record `{type,n,ports,status,created}` → add `external_host` (additive).
