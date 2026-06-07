# M13 — Decisions

_Implementation decisions with rationale. ID scheme: M13-D1, M13-D2, …_

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M13-D1 | `dev-min` size = 10 users / 1 month / 1 org (vs demo small-200's 200/3mo). | Resolves M13-Q3. 10 is the floor that still exercises the role mix (~1 admin + ~6 members + ~3 candidates) so authz/memberships/activity render; smaller rounds the mix to noise. Keeps a fresh dev stack "never empty" without a demo-scale seed cost (<1s). Fixed admin = `dev@anthropos.test` (the local dev login identity → authorized routes 200). | 2026-06-07 |
| M13-D2 | Make the M10 Directus provision contract EXECUTABLE via a new `stack-snapshot/cmd/provision-plan` runner; the dev bring-up shells it rather than re-encoding the recipe in shell. | The M10 `directus.ProvisionPlan`/`EnvContract`/`Validate` was library/test-only. Reusing it (M13's mandate) the right way = one source of truth for the bootstrap→replay→boot recipe AND the prod-Directus firewall (`--check-env`). Re-encoding the recipe in `dev-setdress.sh` would fork the contract. The runner is consumed by dev now and is available to demo. | 2026-06-07 |
| M13-D3 | Set-dressing is default-ON but NON-FATAL on `dev-stack up`; a stale/missing snapshot cache is a WARNING, not a failure. Resolves M13-Q1/Q2. | Lean (M13-Q2): default-on with `--no-snapshot` (skip Directus+replay, seed only) + `--no-setdress` (skip all). Heaviness (M13-Q1): a fresh stack may not be migrated when `up` returns, so set-dressing must not fail the bring-up — the stack comes UP, re-run `dev-setdress.sh` after migration (cache-first + idempotent). Cache-miss-non-fatal: capture is a separate privileged release-time prod read; a dev with no cache degrades to a structural-only world but still seeds (the seed is the guaranteed floor). | 2026-06-07 |
| M13-D4 (bug) | `build_cli` resolved the cmd package one dir too high (`cd dirname/..` → module root's parent-relative path). | Found in §2 PR review by running the REAL build path (stub tests short-circuit build_cli). Fixed to `build_cli(name, module-dir, cmd-pkg)` → `go build ./cmd/<x>` from the module root. Regression test drives the real build. | 2026-06-07 |

## Open at design (to resolve during build)
- M13-Q1: dev bring-up heaviness (mitigate: cache-first snapshot, minimal seed). — resolved M13-D3.
- M13-Q2: auto-snapshot default-on vs opt-in (lean: default-on, `--no-snapshot`). — resolved M13-D3.
- M13-Q3: `dev-min` preset exact size. — resolved M13-D1.

## Adversarial review (close — Phase 2c)

Scenarios considered against the M13 surface (`provision-plan` runner · `dev-setdress.sh` · the `dev-stack`
set-dress wiring · the `dev-min` preset). Each is the *scenario*, not the fix — recorded so future reviewers
see what was probed. All four were already handled (no new code changes from the adversarial pass; the
documented finding — the missing README index row — came from Phase 3, not here).

1. **`--check-env` validates the *derived* Directus env, but `stacksnap replay` runs against `$BASE_DSN` —
   could replay write to a non-offset (base/prod-adjacent) Postgres while the firewall passed on a different
   address?** No. `replay` derives its real write-target via `pg.DSNForOffset(baseDSN, n)`
   (`stacksnap/main.go:281`), which **replaces the port with the stack offset** (`5432 + n·10000`). So even
   though `dev-setdress.sh` passes the un-offset `$BASE_DSN` (default `localhost:5432`), replay writes only to
   the per-stack-isolated offset Postgres — never the base port. The `provision-plan --check-env` Directus env
   is offset-derived the same way (`base_port=$((5432 + N*10000))`). The two are consistent by construction;
   replay's target can never be the base/prod port. Pinned indirectly by `pg`'s `DSNForOffset` tests +
   `TestPrintPlan_*` (offset-only env contracts).

2. **`build_cli` builds CLIs into a `mktemp -d` `BIN_DIR`; two concurrent `dev-stack up` runs on one machine.**
   No shared-dir race: each `dev-setdress.sh` invocation gets its OWN `mktemp -d` (`BIN_DIR` default), so two
   runs build into disjoint dirs. The `[ -x "$BIN_DIR/$name" ] && return 0` short-circuit only ever sees this
   run's dir (or a test-injected stub via `DEV_SETDRESS_BIN`). No cross-run clobber.

3. **`--no-snapshot` skips the Directus firewall check but the seed still runs against `$BASE_DSN` — is the
   seed itself prod-protected when the firewall step was skipped?** Yes. `stackseed` derives its own
   offset-port target from `--stack dev-N` and runs its own 3-layer isolation guard (`isolation.AssertClean`)
   after every run, independent of the snapshot step. Skipping the Directus firewall (a snapshot/Directus
   concern) never weakens the seed's own write-side guard. The seed targets the offset Postgres regardless.

4. **The `provision-plan` recipe printer accepts `--stack anthropos` (N=0) and prints a recipe.** Intended +
   pinned (`TestPrintPlan_RecipePrinterIsStackAgnostic_IncludingN0`): the printer is text-only — it never
   writes a store — so it is correctly stack-agnostic. The N=0 *refusal* is the caller's responsibility and is
   enforced at two layers downstream (`dev-setdress.sh` refuses N=0 without `--force`; `stacksnap`/`stackseed`
   independently refuse it). The printer's only invariant — never wire the SHARED prod Directus host for ANY
   N — holds (asserted across stacks incl. N=0).
