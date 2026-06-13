# M22 — Progress

**Status:** building. **Shape:** section.

## Section checklist
_Derived from `overview.md` § Scope at build start. One section per cohesive unit of work; ticked as it lands._

### §1 — Executed provisioning in the shared engine (ext) ✓
The load-bearing core: replace `dev-setdress.sh`'s print-only block with an **executed** bootstrap → apply-structure → replay → boot, gated `--local-content` (demo default-on / dev opt-in; `N=0` behind `--force`). The `EnvContract.Validate` firewall becomes a **load-bearing executed gate** (hard-abort before any write if the env resolves to prod). Non-fatal: a failed boot degrades to the prod-read path with an honest status line.
- [x] `--local-content` flag through `dev-setdress.sh` + default-on for demo, opt-in for dev
- [x] executed bootstrap (CREATE SCHEMA + `node cli.js bootstrap`) — guarded, idempotent
- [x] executed apply-structure + replay (reuse the M21 auto-provision-during-replay path)
- [x] executed boot (the compose service comes up) + the load-bearing executed env gate
- [x] non-fatal degrade to prod-read with an honest ⚠ status line
- [x] tests (executor unit tests, the firewall executed-gate target-class pin)
_Landed `0c61003` (ext) — attempt 1._

### §2 — Compose-service emission into the per-stack override (ext) ✓
Emit the Directus container as a **compose service** in the stack's override (offset port `8055+N·10000`, joins the app-network, named `<project>-directus-1`) — reusing the proven `gen_injected_override.py` append-block generator (demo) — so `demo-down`/`dev-down`, the port registry, and `stack-verify` cover it with no bespoke lifecycle code.
- [x] directus service block in `gen_injected_override.py` (offset port, network, image pin, env)
- [x] dev-path emission (`stack-core/gen_override.py` or the dev override generator) for `--local-content`
- [x] register the Directus offset port in the unified registry
- [x] tests (generator emits the block; port offset; idempotent re-gen)
_Landed `06d5064` (ext) — attempt 1._

### §3 — Idempotent re-provision guards (ext) ✓
Bootstrap-on-non-empty + container-name-conflict guards, matching the M17 re-run contract — a re-provision converges or fails loudly, never half-applies.
- [x] bootstrap-on-non-empty guard (skip bootstrap if the directus schema is already populated)
- [x] container-name-conflict guard (re-up reuses the existing container, no clash)
- [x] tests (re-run converges; the guards fire on the precondition)
_Landed `9b4390b` (ext) — attempt 1._

### §4 — Directus verify probes (ext)
A SERVICES row + `/server/health`, `directus` added to the expected-schemas list, a **"registered collections > 0"** cheap-win (the silent-failure analog of the casbin assert), and a **no-prod-read env assert**.
- [ ] `directus` SERVICES row + `/server/health` probe (offset/scope-aware)
- [ ] `directus` in the readiness expected-schemas list
- [ ] "registered collections > 0" cheap-win assert in `autoverify.sh`
- [ ] no-prod-read env assert (the served DIRECTUS_BASE_ADDR is the offset instance, not prod)
- [ ] tests (probe rows; cheap-win pass/fail; the env assert)

### §5 — 12 GB-VM preflight accounting (ext)
Extend the preflight accounting to include the Directus container (a non-fatal headroom note).
- [ ] preflight accounting includes the directus container
- [ ] tests (the preflight note covers directus)

### §6 — Docs (rosetta)
- [ ] `corpus/ops/directus-local.md` — the lifecycle half (container/compose/port/teardown + idempotent re-provision + verify probes); fix KB-1 (the `provision.go:108` anchor)
- [ ] `corpus/ops/verification.md` — the directus probe rows
- [ ] `corpus/ops/idempotency.md` — the re-provision re-run rows
- [ ] `corpus/ops/rosetta_demo.md` — registry/teardown note for the directus container

## Build log
_(append per build session)_

- **2026-06-13 (attempt 1, crashed mid-§4):** §1 `0c61003`, §2 `06d5064`, §3 `9b4390b` landed in ext. Phase 0b KB-fidelity GREEN. Network crash (`FailedToOpenSocket`) mid-§4.
- **2026-06-13 (attempt 2, resume-in-place):** reconciled §1–3 checkboxes to the committed ext code; committed the scaffold; finishing §4→§6.
