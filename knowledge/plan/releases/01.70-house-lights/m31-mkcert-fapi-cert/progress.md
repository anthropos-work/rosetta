# M31 — Progress

_Section checklist. Closure when all boxes land. Stub at scaffold._

## Deliverables
- [ ] `up-injected.sh` step 3a-bis: mkcert branch (idempotent `-install` + mint `127.0.0.1 localhost ::1`) inside the keep-existing guard
- [ ] openssl fallback kept verbatim (mkcert-absent AND mint-failure paths); non-fatal throughout
- [ ] `DEMO_NO_MKCERT=1` opt-out parsed + honored (mirrors the DEMO_NO_* flags)
- [ ] code comments retired (`up-injected.sh:337-342`, `gen_injected_override.py:295`) + the dev-N shared-helper forward-note
- [ ] `corpus/ops/demo/recipe-browser-login.md §B` rewritten manual→automatic (+ security/remote/Firefox/expiry notes)
- [ ] `frontend-tier.md` cert one-liner
- [ ] demo-up SKILL browser-login note
- [ ] verified: ZERO change needed to `gen_injected_override.py` / `inject.py` / `fake-fapi/main.go`
- [ ] README-index guard exit 0
- [ ] ext tag `house-lights-m31`
- [ ] (close-time verify) a fresh real browser renders next-web `/home` signed-in, no proceed-anyway

## Notes
_(append build notes here)_
