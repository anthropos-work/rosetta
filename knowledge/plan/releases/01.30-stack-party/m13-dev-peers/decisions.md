# M13 — Decisions

_Implementation decisions with rationale. ID scheme: M13-D1, M13-D2, …_

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M13-D1 | `dev-min` size = 10 users / 1 month / 1 org (vs demo small-200's 200/3mo). | Resolves M13-Q3. 10 is the floor that still exercises the role mix (~1 admin + ~6 members + ~3 candidates) so authz/memberships/activity render; smaller rounds the mix to noise. Keeps a fresh dev stack "never empty" without a demo-scale seed cost (<1s). Fixed admin = `dev@anthropos.test` (the local dev login identity → authorized routes 200). | 2026-06-07 |
| M13-D2 | Make the M10 Directus provision contract EXECUTABLE via a new `stack-snapshot/cmd/provision-plan` runner; the dev bring-up shells it rather than re-encoding the recipe in shell. | The M10 `directus.ProvisionPlan`/`EnvContract`/`Validate` was library/test-only. Reusing it (M13's mandate) the right way = one source of truth for the bootstrap→replay→boot recipe AND the prod-Directus firewall (`--check-env`). Re-encoding the recipe in `dev-setdress.sh` would fork the contract. The runner is consumed by dev now and is available to demo. | 2026-06-07 |
| M13-D3 | Set-dressing is default-ON but NON-FATAL on `dev-stack up`; a stale/missing snapshot cache is a WARNING, not a failure. Resolves M13-Q1/Q2. | Lean (M13-Q2): default-on with `--no-snapshot` (skip Directus+replay, seed only) + `--no-setdress` (skip all). Heaviness (M13-Q1): a fresh stack may not be migrated when `up` returns, so set-dressing must not fail the bring-up — the stack comes UP, re-run `dev-setdress.sh` after migration (cache-first + idempotent). Cache-miss-non-fatal: capture is a separate privileged release-time prod read; a dev with no cache degrades to a structural-only world but still seeds (the seed is the guaranteed floor). | 2026-06-07 |
| M13-D4 (bug) | `build_cli` resolved the cmd package one dir too high (`cd dirname/..` → module root's parent-relative path). | Found in §2 PR review by running the REAL build path (stub tests short-circuit build_cli). Fixed to `build_cli(name, module-dir, cmd-pkg)` → `go build ./cmd/<x>` from the module root. Regression test drives the real build. | 2026-06-07 |

## Open at design (to resolve during build)
- M13-Q1: dev bring-up heaviness (mitigate: cache-first snapshot, minimal seed).
- M13-Q2: auto-snapshot default-on vs opt-in (lean: default-on, `--no-snapshot`).
- M13-Q3: `dev-min` preset exact size.
