# M32 — Progress

_Section checklist. Closure when all boxes land. Stub at scaffold._

## Deliverables
- [x] `gen_injected_override.py` FRONTENDS studio-desk dict: add `NODE_ENV=production` (+ `FRONTEND_PORT=9000`)
- [x] regression assertion in `test_injection.py`: `test_studio_desk_env_pins_node_env_production` (both content paths, studio-desk-only) + block-shape + CORS tests updated
- [ ] Playwright smoke: studio-desk `/home` + a couple routes serve via the production `sendFile` path (no dead-port 302; no 404s) — **CLOSE-TIME** (needs a live studio-desk; route coverage verified by code-read at build, #M32-D1)
- [x] `:9100` sweep — demo-up SKILL (`:9100+`→`:9000+`)
- [x] `:9100` sweep — `frontend-tier.md` (drop dead `:9100` frontend port → single-port `9000`+offset; port row + example + CORS emit + verify registry)
- [x] `:9100` sweep — `gen_injected_override.py` CORS (remove the un-offset `9100` origin + decision note #M32-D2)
- [x] README-index guard exit 0
- [x] ext tag `house-lights-m32` (@ `107599c`)

## Notes
- **Phase 0b KB-fidelity: GREEN** — route coverage verified by code-read; 4 doc/test/CORS staleness items = the planned sweep. Report: `kb-fidelity-audit.md`.
- **Route coverage (the load-bearing open question): VERIFIED via code-read** of `stack-demo/studio-desk/src/index.ts` — the production `sendFile` path covers every dev-block route (+ `express.static` over `dist/public` serving the `.html` targets + an `index.html` SPA fallback); no 404 gap (#M32-D1).
- **Latent bugs fixed Fate-1** (env-masked YAML tests, surfaced under PyYAML): `test_frontend_blocks_parse_to_valid_compose` stale `29100:9100` ports (#M32-D3) + `test_a_plain_service_parses_to_ports_only` predating the universal fix16/17 `DIRECTUS_TOKEN=` strip. Both pre-existing on the M31 tag.
- **Tests:** 88/88 pass with PyYAML (0 skipped), 88 (8 env-skipped) without. `py_compile` clean.
- **Only the close-time live Playwright smoke remains** (by its discipline — a live studio-desk; no demo is up now).
