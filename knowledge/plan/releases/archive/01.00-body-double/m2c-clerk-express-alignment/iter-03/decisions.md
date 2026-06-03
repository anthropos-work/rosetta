# M2c / iter-03 — decisions

## iter-03-D1 — fixed demo RSA keypair (disarmed by design), additive
A **fixed** 2048-bit RSA key is embedded in `shared/rsa.go` (PKCS#8 PEM) — the RS256 analog of
`universalSecret`: reproducible offline, **never production**. RS256 is **additive**: `jwt.go`'s HS256 path
is untouched, so `authn` + the existing browser tokens stay HS256 and M1/M2 gates stay green (verified:
Go 22/22, JS 9/9, full suite 7/7).

## iter-03-D2 — JWKS served by clerk-frontend; crypto in shared/
The key material (private key + `MintRS256` + `JWKS()` + `RS256PublicKeyPEM()`) lives in `shared/` — one
home for the crypto primitives, like the HS256 key. `clerk-frontend`'s `/.well-known/jwks.json` now serves
`shared.JWKS()` (a real RSA key) instead of the empty set — safe because the JS DNA has no JWKS gene and
clerk-js doesn't cryptographically verify it. `RS256PublicKeyPEM()` is exposed for the Node runner's
networkless (`jwtKey`) verification path.
