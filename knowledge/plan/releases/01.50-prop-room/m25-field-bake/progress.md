# M25 — Progress

**Status:** planned (not started). **Shape:** section.

## Section checklist
_The 5 live done-bars from `overview.md` § Scope — each a real stack run on this box, binary pass/fail._

- [ ] **DB-1 — Fresh `/demo-up`**: browser shows content served by the LOCAL Directus (data plane local,
  offset port) with REAL images (asset plane prod, `content.anthropos.work`) + verify net GREEN incl. the new
  Directus probes.
- [ ] **DB-2 — `/dev-up 2 --local-content`**: same observable behavior on an opt-in dev stack; confirm N=0
  stays on the prod-read path untouched.
- [ ] **DB-3 — Re-run everything twice**: idempotency live (re-provision + replay + seed safe + convergent).
- [ ] **DB-4 — Cold-start capture once**: structure + rows captured together from a restored dump.
- [ ] **DB-5 — Clean teardown** (`/demo-down`, `/dev-down 2`): reclaims the Directus container + its port; the
  registry is honest.

## Build log

### Session 2026-06-13 (field-bake)
- Phase 0b KB-fidelity: **YELLOW** — 3 parallel audits, all v1.5 behaviors ALIGNED; 2 arg-hint drifts fixed
  inline (`demo-up` fictional `--full`/env-var flags; `dev-up` missing `--local-content`). Report:
  `kb-fidelity-audit.md`.
