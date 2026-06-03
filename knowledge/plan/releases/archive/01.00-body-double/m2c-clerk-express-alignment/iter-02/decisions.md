# M2c / iter-02 — decisions

## iter-02-D1 — operator + criticality choices for the express DNA
- **ExpressAuth** (critical, 5 variants): operator `error_class` — the middleware's accept/reject is an
  error-class outcome (`valid` → no error; `expired`/`malformed`/`bad-signature`/`no-token` → their
  reject classes). This is the load-bearing capability (does the real SDK accept our token?).
- **ExtractIdentity** (critical): operator `exact` — the extracted `getAuth()` identity must match the
  platform claim set exactly (`sub`/`eid`/`email`/`org_id`/`org_role`).
- **JWKS** (critical): operator `shape` — the served JWKS has the RSA-key structure (`kty`/`alg`/`use`/`kid`/`n`/`e`).
- **ClerkClientBAPI** (standard, 2 variants): operator `shape` — integration confirmations vs the existing
  `clerk-backend` mock (M2c-D4); not new behavior → standard weight.
