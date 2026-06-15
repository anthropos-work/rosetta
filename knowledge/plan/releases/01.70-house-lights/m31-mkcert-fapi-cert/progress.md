# M31 — Progress

_Section checklist. Closure when all boxes land. Stub at scaffold._

## Deliverables
- [x] `up-injected.sh` step 3a-bis: mkcert branch (idempotent `-install` + mint `127.0.0.1 localhost ::1`) inside the keep-existing guard
- [x] openssl fallback kept verbatim (mkcert-absent AND mint-failure paths); non-fatal throughout
- [x] `DEMO_NO_MKCERT=1` opt-out parsed + honored (mirrors the DEMO_NO_* flags)
- [x] code comments retired (`up-injected.sh` 3a-bis header, `gen_injected_override.py:295`) + the dev-N shared-helper forward-note
- [x] `corpus/ops/demo/recipe-browser-login.md §B` rewritten manual→automatic (+ security/remote/Firefox/expiry notes)
- [x] `frontend-tier.md` cert one-liner
- [x] demo-up SKILL browser-login note (+ `DEMO_NO_MKCERT` in the argument-hint)
- [x] verified: ZERO change needed to `gen_injected_override.py` / `inject.py` / `fake-fapi/main.go` (path-only consumers — M31-D4)
- [x] README-index guard exit 0
- [x] ext tag `house-lights-m31`
- [ ] (close-time verify) a fresh real browser renders next-web `/home` signed-in, no proceed-anyway

## Notes
- **Ext commit** `6565ef8` (on ext `m31/mkcert-fapi-cert`, from base `868a68a`): the mkcert branch + the
  `gen_openssl_fapi_cert` factored fallback + `DEMO_NO_MKCERT` parse + the comment retire/forward-note +
  the `gen_injected_override.py:295` comment touch-up + a new `FapiCertStep` test class (8 tests).
- **Rosetta doc commit** `359eee4` (on `m31/mkcert-fapi-cert`): recipe-browser-login.md §B rewrite +
  frontend-tier.md one-liner + demo-up SKILL note.
- **Tests:** `FapiCertStep` 8/8 pass (5 static body-pins + functional all-4-branches + idempotency, values-blind,
  stub mkcert, no live docker). Full `test_tooling.py` 47/47 + `test_frontend_build.py` 42/42 (lib-only sourcing
  still works — the cert step's fn sits below the seam). `shellcheck 0.11.0` clean. README-index guard exit 0.
- **Cert-consumer no-touch verified** (M31-D4): `fake-fapi/main.go` + `gen_injected_override.py` mount +
  `inject.py` all reference the cert by path only — a trusted cert at the same path serves identically.
- **Why the last box stays open:** the live "fresh real browser renders next-web signed-in, no proceed-anyway"
  is the **close-time** verify (a live `/demo-up` + browser, not run at build per the milestone contract). The
  static + functional tests prove the cert-mint logic; the render proof is close-milestone's job.
