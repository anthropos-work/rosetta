# M2c / iter-05 — decisions

## iter-05-D1 — runner architecture (Go-mint + embedded-Node-verify + bapi-httptest)
`expressrun` (Go) mints the scenario RS256 tokens (`shared.MintRS256`) and shells to an **embedded**
`verify.js` (`//go:embed`) that drives the **real** `@clerk/backend.verifyToken` networkless via `jwtKey`
(`NODE_PATH` from `$EXPRESSRUN_NODE_PATH`). The `ExpressAuth` error classes are **normalized from
`@clerk/backend`'s actual reasons** (`token-expired`→`expired`, `…-signature`→`bad-signature`,
`token-invalid`→`no-token`, parse errors→`malformed`). `ExtractIdentity` compares the extracted
`{sub,eid,email,org_id,org_role}` (operator `exact`). The `ClerkClientBAPI` genes `httptest` the existing
`clerk-backend` mock. The differential, load-bearing genes (ExpressAuth ×5 + ExtractIdentity) have
independent hand-authored source goldens; JWKS + ClerkClientBAPI confirm structure (shape).

## iter-05-D2 — additive bapi read handlers
`GET /v1/organizations/{id}` + `…/memberships` (+ `Store.GetOrganization`/`ListMemberships`) added to
`clerk-backend` for the 2 `ClerkClientBAPI` integration genes. **Additive** — the 10 write methods are
unchanged, so the Go gate stays 22/22 (verified). `getOrganization` returns the org with
`public_metadata.eid` (studio-desk's `loadOwnedPath` tenant key).
