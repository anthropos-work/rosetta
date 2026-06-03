# M2c / iter-04 — decisions

## iter-04-D1 — additive RS256 validated against the real SDK; runner is Go-mints + Node-verifies
The genuine `@clerk/backend.verifyToken` accepted a `shared.MintRS256` token via the **networkless
`jwtKey`** path (`shared.RS256PublicKeyPEM()`), returning `sub`/`org_id`/`org_role` — **no `iss`/`azp`
tuning required** (those are only checked when `authorizedParties`/issuer options are passed). This
**confirms TOK-01's additive strategy**: the HS256 seams (authn/clerk-frontend/shared HS256) need no
migration; the `clerk-express/` surface uses RS256 exclusively and the real SDK is satisfied.

**Runner shape (confirms M2c-D5):** a Go runner (`alignment/cmd/expressrun`) mints the scenario tokens
(`shared.MintRS256`) and shells to a co-located Node verifier (`verify.js`) that drives the real
`@clerk/backend`. `@clerk/backend` resolves from studio-desk's `node_modules` via `NODE_PATH`; iter-05
decides whether to keep that shim or give the runner its own `package.json`.
