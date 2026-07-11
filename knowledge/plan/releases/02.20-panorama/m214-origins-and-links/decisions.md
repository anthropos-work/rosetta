# M214 — decisions

_(implementation choices with rationale, recorded as the milestone is built)_

## D-PATCH-1 — the patch tail rides the existing sha-pinned mechanism
The one required platform-family change (ant-academy `allowedDevOrigins`) and the conditional one (next-web
`urls.ts`) go through the rext `apply-*.sh` / demopatch surface (drift-refuse, idempotent, non-fatal), never a raw
clone edit. **Why:** the platform stays read-only (CLAUDE.md hard rule); drift-refusal makes an upstream change fail
loudly by design.
