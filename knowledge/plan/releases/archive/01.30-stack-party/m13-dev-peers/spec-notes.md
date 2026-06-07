# M13 — Spec notes

Technical notes accumulate here during build.

## Pre-flight audits — §1 (dev bring-up: Directus + snapshot + seed)
**KB-fidelity (Phase 0b): GREEN** — report `kb-fidelity-audit.md`. All M13 KB deps PAIRED + ALIGNED.
Topic → doc → code triples (verified true before coding):
- Local per-stack Directus on dev → `snapshot-spec.md` § per-stack Directus store fork → `stack-snapshot/directus/provision.go` (`ProvisionPlan`/`EnvContract`/`Validate`)
- Auto-snapshot replay (cache-first) → `snapshot-spec.md` § stacksnap CLI / store → `stack-snapshot/cmd/stacksnap/main.go` `replayCmd` (cache-hit-or-fail; `pg.ParseStackN` already parses `dev-N`)
- `dev-min` preset + dev auto-seed → `seeding-spec.md` § blueprint/CLI → `stack-seeding/presets/*.seed.yaml` + `blueprint/blueprint.go` (strict `KnownFields`)
- n=0-dev-reset guard → `seeding-spec.md` § CLI → `stack-seeding/cmd/stackseed/main.go:180-181` (`if n == 0 && !force`)
- dev bring-up + unified registry → `dev-stack/README.md` + `rosetta_demo.md` → `dev-stack/dev-stack` + `stack-core/stack_registry.py`

KEY FINDINGS that shape the build:
- **`stacksnap replay` already targets `dev-N`** (no replay-code change). "cache-first" == call `replay` (resolves cache; never captures).
- **M10 provision.go is a declarative plan + env contract** (bootstrap→replay→boot), not a container runner. M13 wires the dev bring-up to execute/document this.
- **Platform compose has NO Directus service** (prod Directus is external `content.anthropos.work`). Per-stack Directus is a standalone `docker run` (the provision plan), not part of `gen_override.py`.

## Local per-stack Directus on dev
_Reuse M10 `stack-snapshot/directus/provision.go`; repoint dev CMS at the per-stack Directus._

## Auto-snapshot on dev build
_`stacksnap replay` taxonomy + directus, cache-first; `--no-snapshot` escape._

**Implemented (§2):** `dev-setdress.sh` runs `stacksnap replay --surface {taxonomy,directus} --stack dev-N`.
Cache-first is INHERENT — `replay` resolves cache-hit vs stale and never captures (a miss exits 1 telling you
to capture; the script treats that as a warning and proceeds to seed). No replay-code change was needed (replay
already accepts `dev-N`; `pg.ParseStackN("dev-3")==3`). The per-stack Directus surface additionally needs its
container booted (bootstrap→replay→boot) — emitted as the operator recipe, not booted in-build (M9b/M10 discipline).

## dev-min seed preset
_~1 org + ~10 users + minimal activity; n=0 reset guard preserved._

**Implemented (§1+§2):** preset `stack-seeding/presets/dev-min.seed.yaml`; applied on build via
`dev-setdress.sh` → `stackseed --stack dev-N --seed dev-min`. n=0 guard: TWO layers — the existing `stackseed
--reset` n=0 refusal (unchanged) + a NEW `dev-setdress.sh` n=0 refusal (never auto-set-dress the main dev stack
without `--force`). The seed targets the offset-port Postgres (`5432 + N*10000`; verified dev-1 → 15432 live).

## §2 wiring map (the dev bring-up)
- `dev-stack up` → (after bring-up/inject) `dev-setdress.sh $n [--no-snapshot]`, default-on, NON-FATAL.
  Flags: `--no-snapshot` (skip Directus+replay, seed only) · `--no-setdress` (skip all).
- `dev-setdress.sh` → builds CLIs into a mktemp `BIN_DIR` (tests inject stubs via `DEV_SETDRESS_BIN`):
  `provision-plan` (recipe + `--check-env` firewall) · `stacksnap` (replay) · `stackseed` (dev-min).
- NEW Go runner `stack-snapshot/cmd/provision-plan` makes M10 `directus.ProvisionPlan`/`EnvContract` executable.
- Tests: `dev-stack/tests/test_dev_stack.py` +11 (set-dress class + contract + real-build regression); the
  Go runner +6. Hermetic via stubbed CLIs; one go-guarded real-build test pins build_cli.
