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

### §4 — Directus verify probes (ext) ✓
A SERVICES row + `/server/health`, `directus` added to the expected-schemas list, a **"registered collections > 0"** cheap-win (the silent-failure analog of the casbin assert), and a **no-prod-read env assert**.
- [x] `directus` SERVICES row + `/server/health` probe (offset/scope-aware)
- [x] `directus` in the readiness expected-schemas list (container-presence-gated — prod-read stays clean)
- [x] "registered collections > 0" cheap-win assert in `autoverify.sh` (+ a readiness `directus-collections` probe)
- [x] no-prod-read env assert (the served Directus DB resolves to the offset instance, not prod)
- [x] tests (probe rows; cheap-win pass/fail; the env assert; readiness schema-gate; +2 demo-stack chain fixes)
_Landed `7235181` (ext) — attempt 2._

### §5 — 12 GB-VM preflight accounting (ext) ✓
Extend the preflight accounting to include the Directus container (a non-fatal headroom note).
- [x] preflight accounting includes the directus container (UI-independent runtime note; +1 GiB)
- [x] tests (the preflight note covers directus; on/off; fires-under-no-ui; dropped-on-no-local-content)
_Landed `94399e9` (ext) — attempt 2._

### §6 — Docs (rosetta) ✓
- [x] `corpus/ops/directus-local.md` — the lifecycle half (container/compose/port/teardown + idempotent re-provision + verify probes); fix KB-1 (the `provision.go:108` anchor → prose/symbol ref)
- [x] `corpus/ops/verification.md` — the directus probe rows (cheap-wins + readiness liveness/serve rows)
- [x] `corpus/ops/idempotency.md` — the re-provision re-run rows (verdict table + engineers subsection)
- [x] `corpus/ops/rosetta_demo.md` — registry/teardown note for the directus compose service
- [x] (collateral, Fate-1) retired the now-stale print-only claims in `snapshot-spec.md`, `safety.md`, `demo/README.md`
_Landed `0ab823a` (rosetta) — attempt 2._

## Build log
_(append per build session)_

- **2026-06-13 (attempt 1, crashed mid-§4):** §1 `0c61003`, §2 `06d5064`, §3 `9b4390b` landed in ext. Phase 0b KB-fidelity GREEN. Network crash (`FailedToOpenSocket`) mid-§4.
- **2026-06-13 (attempt 2, resume-in-place):** reconciled §1–3 checkboxes to the committed ext code; committed the scaffold; finished §4 (verify probes, `7235181`), §5 (preflight, `94399e9`), §6 (docs, `0ab823a`). **All sections done.** Collateral Fate-1 fixes: studio-desk verify-port test (M22-D2), 3 stale print-only doc claims (M22-D1). Tests at exit: 102 stack-verify / 66 dev-stack / 87 demo-stack green; shellcheck clean.

## M22: Hardening

### Pass 1–4 — 2026-06-13

**Scope manifest (Phase 1 — milestone-touched code, ext repo `58b810a..94399e9`):**
M22's production code is entirely in the gitignored ext repo (`.agentspace/rosetta-extensions/`); the rosetta worktree holds docs only (`corpus/ops/*`, no executable surface). Two stacks of concern: shell engines (subprocess-tested) + Python generators (directly coverable). Every touched source file has a co-located test.

| Source (ext) | Stack | Co-located tests | M22 surface |
|---|---|---|---|
| `dev-stack/dev-setdress.sh` | shell | `dev-stack/tests/test_dev_stack.py` | executed provision (bootstrap→replay→boot), firewall gate, non-fatal degrade |
| `dev-stack/dev-stack` | shell | same | `--local-content` threading (override + setdress + verify scope) |
| `stack-core/gen_override.py` | python | `stack-core/tests/test_gen_override.py` | `directus_lines` + `to_yaml(with_directus)` + `main()` summary |
| `stack-injection/gen_injected_override.py` | python | `stack-injection/tests/test_injection.py` | demo directus compose-service block |
| `stack-verify/lib/{services,readiness}.sh`, `live/{autoverify,verify}.sh` | shell | `stack-verify/tests/test_verify.py` | SERVICES row, schema-gate, collections probe, cheap-wins, env-assert |
| `demo-stack/up-injected.sh` | shell | `demo-stack/tests/test_frontend_build.py` | 12 GB preflight directus runtime note |

**Coverage (directly-coverable Python generators):**
- `stack-core/gen_override.py`: statements 62% -> **85%** (+23) — `main()` + `--with-directus` summary now covered; residual is the docker subprocess (`resolved_config`) + pre-M22 `build_override` edges.
- `stack-injection/gen_injected_override.py`: **99%** at entry (only the `__main__` guard uncovered) — no change needed.
- Shell engines (dev-setdress, autoverify, readiness, up-injected) are not Python-coverable; deepened behaviorally via subprocess harnesses (stubbed docker/CLIs recording argv).

**Tests added (+11 total; 0 production bugs surfaced):**
- `test_dev_stack.py` (+8): restart-failure non-fatal warn; N=0+`--force` executed provision (base band) + N=0 refusal-before-provision; `--no-snapshot`+`--local-content` `set -u` guard; directus-replay-skip restart-gating; verify `--services` directus-conditional scope (body-grep); `--local-content`/`--no-local-content` last-wins ordering.
- `test_verify.py` (+2): DB_CONNECTION_STRING-unreadable env-assert skip (not a false prod-leak); non-numeric collections count hits the numeric guard (warn, not crash).
- `test_gen_override.py` (+2): `main()` end-to-end via `resolved_config` monkeypatch — with/without `--with-directus`, the summary accounting.
- + coverage instrumentation: gitignore for `.coverage*`/`coverage.xml`/`htmlcov/`/`.pytest_cache/` (pytest-cov 7.1.0 already on the interpreter).

**Bugs fixed inline:** none — the build phase's tests were sound; this pass deepened error/edge/`main()` coverage the build-minimum didn't reach.

**Flakes stabilized:** none observed (3 consecutive clean runs of the new tests; shell subprocess harnesses are deterministic — stubbed docker/CLIs, no timing/random/shared-port surface).

**Knowledge backfill:** no KB-worthy findings — every behavior the new tests pin (the executed-provision flow, the firewall gate, the non-fatal degrade, the verify cheap-wins, the last-wins flag contract) is already documented in `corpus/ops/directus-local.md` / `verification.md` / `idempotency.md` from the §6 docs commit. No new invariant, edge semantic, or threshold was discovered — the tests confirm documented behavior rather than surfacing undocumented behavior.

**Commits (ext):** `08fc875` (gitignore instrument), `e782458` (Pass 1 dev-stack), `fa11c52` (Pass 2 stack-verify), `d62c685` (Pass 3 stack-core), `93ad686` (Pass 4 dev-stack last-wins).

**Suite totals at exit:** stack-verify 104 (+2), dev-stack 73 (+7), demo-stack 87, stack-core 61 (+2), stack-injection 93 (+8 skipped, env-gated). All green; shellcheck clean.

### Stop condition
Scan clean at Pass 4 — the full six-dimension sweep (test depth / edge / error paths / regression / fuzz / perf) found only one thin arg-parse edge in Pass 4, coverage deltas on the directly-coverable surface stabilized (85%/99%, residual all out-of-M22-scope), the shell engines are exhaustively exercised behaviorally, and zero flakes across runs. Stopped at 4 of 5 passes.
