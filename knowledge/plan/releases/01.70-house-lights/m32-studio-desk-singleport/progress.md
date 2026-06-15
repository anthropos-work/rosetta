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

## M32: Hardening

### Pass 1 — 2026-06-15
Small, two-file code surface (`gen_injected_override.py` + `tests/test_injection.py`) + a doc sweep — scanned all
six dimensions in-thread. Coverage tool: stdlib `unittest`; no per-file coverage instrument exists in this layer
(behavior-pinned suite). Ran both PyYAML modes: `python3.11` (PyYAML 6.0.2) for the structural-parse tier, the
default `python3` (3.14, no PyYAML) for the env-skip tier.

**Scope manifest (milestone-touched, vs `house-lights-m31`):**
- `stack-injection/gen_injected_override.py` — studio-desk override `NODE_ENV=production` + `FRONTEND_PORT=9000`; backend CORS `9100` removal. Tests: `test_injection.py` (co-located, `TestFrontendTier` + `TestGenInjectedOverride`).
- `stack-injection/tests/test_injection.py` — the regression assertions themselves.
- rosetta docs: `corpus/ops/demo/frontend-tier.md`, `.claude/skills/demo-up/SKILL.md` (the `:9100` sweep — verified, not re-edited).

**Verifications (the harden focus, all GREEN):**
- **Override-merge invariant — mutation-checked 4 ways.** `test_studio_desk_env_pins_node_env_production` correctly FAILS if (A) `NODE_ENV=production` dropped, (B) `FRONTEND_PORT=9000` dropped, (C) both pins moved to next-web, (D) next-web *additionally* gains `NODE_ENV` (the studio-desk-scoped half — fails on the `next-web NODE_ENV == []` assertion). Studio-desk-scoped in both directions; no false-pass.
- **CORS removal — emitted set verified exactly** `{3000,3001,9000}+offset` at offsets 0 / 10000 / 30000; no `9100`; other origins intact + offset-correct.
- **Generated-compose validity (PyYAML)** — studio-desk service parses to a valid compose dict, single-port `9000:9000`+offset, env carries the pins, no `9100` in the block; 6 PyYAML-gated tests pass.
- **Doc consistency** — `frontend-tier.md` + demo-up SKILL `:9100`→`:9000` sweep complete; zero demo/offset `9100` claims (no `19100/29100`); remaining `9100` mentions are the removal-narrative. The dev/native/base-compose `9100` refs elsewhere in the corpus are correctly out of M32's demo-only scope (`9100` IS the real Vite dev port under `npm run dev`).
- **Suite + guards** — 88/88 with PyYAML (0 skipped), 88 (8 env-skipped) without; full suite stable across 3 consecutive runs (fuzz/random); `py_compile` clean on 3.11 + 3.14; README-index guard exit 0.

**Tests added:**
- `gen_injected_override.py` (CORS): +1 exact-set regression assertion in `test_backend_gets_cors_extra_origins_at_offset` (the origins must equal `[3000,3001,9000]+offset` in emit order).

**Bugs fixed inline:**
- None (no production bug). One **test-coverage gap closed** (commit `7b17c39`): the CORS test had only membership asserts (`13000`/`19000` present, `19100` absent) and did NOT pin the full surviving set — over-removing the `3001` next-web origin passed the entire 88-test suite (mutation-confirmed). The new exact-set assertion catches over-removal of any kept origin AND an accidental re-add (`9100` or otherwise). This closes the over-removal risk the M32-D2 CORS edit introduced.

**Flakes stabilized:** None observed (3 consecutive clean runs of the M32-touched tests + 3 clean full-suite runs).

**Knowledge backfill:** No KB-worthy findings this pass — the override-merge mechanism, route coverage, and the `9100` removal rationale are already documented in `frontend-tier.md` (the M32 single-port narrative) + `decisions.md` (M32-D1/D2/D4). The gap closed was a test-rigour gap, not a system truth.

### Stop condition
Single pass: the full six-dimension scan found exactly one genuine gap (the CORS exact-set), it was closed
Fate-1, and a re-scan surfaced nothing further worth adding on this two-file surface. No coverage tooling to
delta against (behavior-pinned `unittest` suite); the mutation checks ARE the coverage evidence. The close-time
live Playwright smoke remains OPEN by its discipline (needs a live studio-desk; route coverage already
code-read-verified at #M32-D1) — not a harden-pass item.
