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

## dev-min seed preset
_~1 org + ~10 users + minimal activity; n=0 reset guard preserved._
