# M29 — Spec notes

_Technical detail accumulated during build. Stub at scaffold; sections seeded from the overview scope._

## Pre-flight audits — secrets-spec.md (first section)
- **Phase 0b KB-fidelity:** **GREEN** (2026-06-14). Report:
  `knowledge/plan/releases/01.60-stage-door/m29-secrets-docs-skill/kb-fidelity-audit.md`.
  One BLIND-AREA finding (secret provisioning has zero corpus coverage) — resolved by the milestone's own
  `delivers:` line (`corpus/ops/secrets-spec.md` net-new + `/stack-secrets` skill). No stale claims; reference docs
  (seeding-spec / snapshot-spec / stack-seed skill / safety.md / CLAUDE.md) all current. Ground-truth = ext
  `stack-secrets` @ tag `stage-door-m28` (`9742126`). All M29 sections share this subsystem → reuse across sections.

## Topic → doc → code triples (for future audits)
- secret provisioning → `corpus/ops/secrets-spec.md` → `stack-secrets/{provision,secretdna,source,cmd/stacksecrets}`
- `/stack-secrets` skill → `.claude/skills/stack-secrets/SKILL.md` → `stacksecrets` @ `stage-door-m28`
- values-blind safety → `corpus/ops/safety.md` (new clause; §2.8 = M28-harden shell invariant, do not clobber)
- manual-copy retire → `corpus/ops/setup_guide.md` (studio-desk §, ant-academy §, next-web §, line-439 TODO)

## corpus/ops/secrets-spec.md outline
_(source-dir/zip layout · the secret-DNA · per-repo target-file map · values-blind safety statement · alias/collision
rules · waived-class rationale.)_

## /stack-secrets skill shape
_(argument-hint, default source, the read-spec→confirm→build-tagged-binary→run→report flow; mirror /stack-seed.)_

## CLAUDE.md wiring
_(skill-table row · Key-Documentation-Locations entry · Interconnected-Documentation list update.)_

## setup_guide.md + safety.md edits
_(retire the manual-copy prose + delete the line-447 TODO; add the never-echo / PreflightEnv-emitting clause.)_
