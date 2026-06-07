# M13 — Progress

**Shape:** section · **Status:** planned

## Section checklist (from overview Scope.In)
- [x] Dev bring-up spawns a per-stack Directus (reuse M10 provision.go) + dev CMS repointed at it _(§2: `dev-setdress.sh` emits the M10 bootstrap→replay→boot recipe via the new `stack-snapshot/cmd/provision-plan` runner — makes the M10 `directus.ProvisionPlan`/`EnvContract` contract executable — + firewall-checks the per-stack Directus env (prod-Directus abort); CMS repoint = the `DIRECTUS_BASE_ADDR` offset-port env in the recipe. Live boot = operator step, M9b/M10 discipline.)_
- [x] Auto-run `stacksnap replay` (taxonomy + directus) on dev build, cache-first; `--no-snapshot` escape _(§2: `dev-setdress.sh` runs `stacksnap replay --surface taxonomy|directus --stack dev-N` — cache-first by construction (replay resolves cache; never captures); a stale/missing cache is a warning, not a failure. Wired default-on into `dev-stack up`; `--no-snapshot` skips Directus+replay (seed only), `--no-setdress` skips the whole pass.)_
- [x] `dev-min` seed preset (~1 org + ~10 users + minimal activity) applied on build _(§1: `stack-seeding/presets/dev-min.seed.yaml` — 10 users/1mo/Dev Sandbox, dev@anthropos.test admin; pinned in `presets_test.go`. §2: applied on build via `dev-setdress.sh` → `stackseed --stack dev-N --seed dev-min`.)_
- [x] n=0-dev-reset guard preserved _(unchanged in `stackseed` [`main.go:180-181`, `--reset` refuses N=0 without `--force`]; §2 ADDS a second n=0 guard in `dev-setdress.sh` — refuses to auto-set-dress N=0 without `--force`, so an auto-seed never touches the main dev stack.)_
- [x] Delivers: seeding-spec.md (dev-min + dev auto-seed) + snapshot-spec.md (dev replay target + local Directus) _(§3: `seeding-spec.md` — the shipped-presets table + the dev-min/dev-auto-seed subsection + the two-layer n=0 guard; `snapshot-spec.md` — the "Dev as a full-fidelity peer (M13)" section + dev in the scope/replay-CLI lines. Stale-claim fix: cloud/S3 store → v1.4. All anchors verified.)_

## M13: Hardening

### Pass 1 — 2026-06-07
**Scope manifest (milestone-touched source, `git diff be1c979..HEAD` minus tests/docs/yaml):**
- `stack-snapshot/cmd/provision-plan/main.go` — the new M13 runner (recipe printer + `--check-env` firewall). Tests: `main_test.go`.
- `dev-stack/dev-setdress.sh` — the set-dressing helper (per-stack Directus recipe + cache-first replay + dev-min seed). Tests: `tests/test_dev_stack.py::DevSetdress`.
- `dev-stack/dev-stack` — the `up` → set-dress wiring. Tests: `tests/test_dev_stack.py::DevStackContract` (static) + (new) `DevStackSetdressWiring` (integration).
- `stack-seeding/presets/dev-min.seed.yaml` — declarative data (no logic). Pinned by `blueprint/presets_test.go` (parse/validate + floor + ordered-size).
- New-unit handbook check: `provision-plan` is a sub-command of the existing `stack-snapshot` module (not a new top-level unit); it's documented in `corpus/ops/snapshot-spec.md` §"Dev as a full-fidelity peer" + the dev-stack README. No missing handbook.

**Coverage delta (milestone-touched files):**
- `cmd/provision-plan/main.go`: statements 90.7% → **93.0%**; `run` 94.4% → **100%**, `checkEnvContract` → **100%** (remaining: the `main` os.Exit shim + a defensive `printPlan` branch now pinned-unreachable by an invariant test).
- `directus/provision.go` (M10 surface reused by the runner): already **100%** (ProvisionPlan/DefaultEnvContract/Validate/sanitizeEnv).
- `blueprint` (dev-min preset loader): **100%**.
- `dev-stack` bash: no native coverage tool (bash) — exercised by subprocess integration tests (the project's documented strategy for the CLIs); pytest funcs 33 → **38**.

**Tests added:**
- `cmd/provision-plan/main_test.go`: +5 (`TestUsage_EmptyStackArgIsRejected`, `TestUsage_UnknownFlagIsUsageError`, `TestCheckEnv_RejectsBlankArgs` [5 sub-cases, incl. 2 regression], `TestPrintPlan_RecipePrinterIsStackAgnostic_IncludingN0`, `TestDefaultEnvContract_AlwaysProdSafe_AcrossPorts` [invariant]).
- `dev-stack/tests/test_dev_stack.py`: +5 — `DevStackSetdressWiring` (3: non-fatal failed-set-dress + the re-run-hint flag-fidelity regression + `--no-setdress` skips) + `DevSetdress` (2: trailing-value-flag regression + value-flag happy-path).

**Bugs fixed inline (3 — all Fate-1, committed with their pinning regression tests):**
- `provision-plan --check-env` green-lit a **whitespace-only** Directus env as "prod-safe" (the exact-`""` guard let whitespace slip past into `Validate`, which only rejects an empty DSN — not a blank BaseAddr). Now trims+fails-closed. Commit `1e8a510`.
- `dev-stack` set-dress **re-run hint always appended `--no-snapshot`** (`${no_snapshot:+…}` always expands since the var is `0`/`1`, both non-empty) — would silently skip the snapshot on a manual re-run. Commit `b654573`.
- `dev-setdress.sh` trailing `--dsn`/`--seed` (no value) leaked a raw `$2: unbound variable` (`set -u`) instead of a clean usage message. Added `needval`. Commit `3404863`.

**Flakes stabilized:** none observed (0 flakes across 3 sequential runs of the new tests).

**Knowledge backfill:** no KB-worthy findings. The three bugs were implementation gaps *below* the already-correct documented contract (`snapshot-spec.md` §"Dev as a full-fidelity peer" describes the firewall as fail-closed + the non-fatal cache-miss + the escapes); none of the fixes change the documented behavior — they make the code match it. Verified no corpus passage referenced the buggy behavior.

### Stop condition
Pass 1 only. Surface is small (4 source files); the full Step-2b scan is exhausted — every logic-bearing path is covered or pinned-unreachable, the only uncovered lines are the conventional `os.Exit` shim + an invariant-guarded defensive branch. 3 real bugs found, fixed, and pinned (each fails on pre-fix code). A pass 2 would yield <2% delta with no new qualitative gaps. Both modules `-race` clean (`-count=1`); gofmt + vet clean; both CLIs shellcheck-clean; py_compile clean; 0 flakes.

## Final review
_(filled at close)_
