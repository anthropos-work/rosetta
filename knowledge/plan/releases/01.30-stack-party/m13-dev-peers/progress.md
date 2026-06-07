# M13 — Progress

**Shape:** section · **Status:** planned

## Section checklist (from overview Scope.In)
- [x] Dev bring-up spawns a per-stack Directus (reuse M10 provision.go) + dev CMS repointed at it _(§2: `dev-setdress.sh` emits the M10 bootstrap→replay→boot recipe via the new `stack-snapshot/cmd/provision-plan` runner — makes the M10 `directus.ProvisionPlan`/`EnvContract` contract executable — + firewall-checks the per-stack Directus env (prod-Directus abort); CMS repoint = the `DIRECTUS_BASE_ADDR` offset-port env in the recipe. Live boot = operator step, M9b/M10 discipline.)_
- [x] Auto-run `stacksnap replay` (taxonomy + directus) on dev build, cache-first; `--no-snapshot` escape _(§2: `dev-setdress.sh` runs `stacksnap replay --surface taxonomy|directus --stack dev-N` — cache-first by construction (replay resolves cache; never captures); a stale/missing cache is a warning, not a failure. Wired default-on into `dev-stack up`; `--no-snapshot` skips Directus+replay (seed only), `--no-setdress` skips the whole pass.)_
- [x] `dev-min` seed preset (~1 org + ~10 users + minimal activity) applied on build _(§1: `stack-seeding/presets/dev-min.seed.yaml` — 10 users/1mo/Dev Sandbox, dev@anthropos.test admin; pinned in `presets_test.go`. §2: applied on build via `dev-setdress.sh` → `stackseed --stack dev-N --seed dev-min`.)_
- [x] n=0-dev-reset guard preserved _(unchanged in `stackseed` [`main.go:180-181`, `--reset` refuses N=0 without `--force`]; §2 ADDS a second n=0 guard in `dev-setdress.sh` — refuses to auto-set-dress N=0 without `--force`, so an auto-seed never touches the main dev stack.)_
- [ ] Delivers: seeding-spec.md (dev-min + dev auto-seed) + snapshot-spec.md (dev replay target + local Directus) _(§3)_

## Final review
_(filled at close)_
