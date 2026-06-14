# M22 — Spec notes

_Technical detail that doesn't belong in `overview.md` (code maps, contracts, edge cases). Accumulates during build._

## Pre-flight audits — §1 (Executed provisioning)

KB-fidelity audit (2026-06-13): **GREEN**. Report: `kb-fidelity-audit.md`. Sha at audit: `58b810a` (ext) / `a313171` (rosetta HEAD at build start). One incidental completeness gap (KB-1, the `provision.go:108` anchor) tracked for Phase 5; zero blind areas, zero stale load-bearing claims. Reusable across all M22 sections (single subsystem: the shared `dev-setdress` engine + its ext-repo callees + the 3 rosetta docs).

## Topic → doc → code triples (audit start cache)

| Topic | Doc | Code |
|---|---|---|
| Executed provisioning | `snapshot-spec.md` §store-fork, `directus-local.md` | `dev-stack/dev-setdress.sh::snapshot_step`, `stack-snapshot/directus/provision.go`, `cmd/provision-plan/main.go` |
| Compose-service emission | `directus-local.md` §future | `stack-injection/gen_injected_override.py::build_lines/frontend_lines`, `stack-core/gen_override.py` |
| Idempotent re-provision | `idempotency.md` | `cmd/stacksnap/autoprovision.go::tryAutoProvision`, `dev-setdress.sh` |
| Directus verify probes | `verification.md` | `stack-verify/lib/services.sh::SERVICES`, `live/autoverify.sh`, `lib/readiness.sh` |
| 12 GB preflight | `demo/frontend-tier.md` | `demo-stack/up-injected.sh::preflight_vm_ram` |

## Code map — the surfaces M22 touches

- **`dev-stack/dev-setdress.sh`** — `snapshot_step()` (L99-154) currently: builds `stacksnap`+`provision-plan`, prints the recipe via `provision-plan`, firewall-checks the env (`--check-env`), then runs `stacksnap replay` for taxonomy+directus. M22 turns the print into an **executed** bootstrap→apply→replay→boot, gated by `--local-content` (demo default-on, dev opt-in), emitting the compose service.
- **`stack-snapshot/directus/provision.go`** — `ProvisionPlan` (the recipe steps, illustrative `Command`s), `EnvContract`/`DefaultEnvContract`/`Validate` (the firewall + offset derivation: directus port = basePort + 8055 - 5432). This is the contract the executed steps consume.
- **`stack-snapshot/cmd/stacksnap/autoprovision.go`** — `tryAutoProvision` (M21) already applies structure on a bootstrap-gap during replay. So once Directus is bootstrapped, `stacksnap replay --surface directus` provisions+loads in one step — the executor only needs bootstrap + boot around it.
- **`stack-injection/gen_injected_override.py`** — `build_lines` appends net-new service blocks (frontends, fake-fapi/bapi) with offset ports + `mem_limit` + `networks`. The directus block follows this exact pattern. **demo** path.
- **`stack-core/gen_override.py`** — ports/volumes only today. The **dev** path; M22 (or M23) adds the directus service emission for dev opt-in.
- **`stack-verify/lib/services.sh`** — `SERVICES` registry (base ports, offset applied centrally by `service_rows`). Add a `directus` row.
- **`stack-verify/live/autoverify.sh`** — cheap-win asserts (`/api/health`, `casbin_rules>0`). Add a directus "registered collections > 0" cheap-win + a no-prod-read env assert.
- **`demo-stack/up-injected.sh`** — `preflight_vm_ram` (12 GB). Extend the accounting note for the directus container.
