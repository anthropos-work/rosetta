# M32 — Progress

_Section checklist. Closure when all boxes land. Stub at scaffold._

## Deliverables
- [ ] `gen_injected_override.py` FRONTENDS studio-desk dict: add `NODE_ENV=production` (+ `FRONTEND_PORT=9000`)
- [ ] regression assertion in `test_injection.py` (~820-857): `NODE_ENV=production` in the studio-desk env block
- [ ] Playwright smoke: studio-desk `/home` + a couple routes serve via the production `sendFile` path (no dead-port 302; no 404s)
- [ ] `:9100` sweep — demo-up SKILL (`:9100+`→`:9000+`)
- [ ] `:9100` sweep — `frontend-tier.md:21` (drop dead `:9100` frontend port → single-port `9000`+offset)
- [ ] `:9100` sweep — `gen_injected_override.py:249` CORS (remove the un-offset `9100` origin + decision note)
- [ ] README-index guard exit 0
- [ ] ext tag `house-lights-m32`

## Notes
_(append build notes here)_
