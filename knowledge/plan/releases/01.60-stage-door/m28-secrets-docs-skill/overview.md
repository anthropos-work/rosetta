---
milestone: M28
slug: secrets-docs-skill
version: v1.6 "stage door"
milestone_shape: section
status: planned
created: 2026-06-14
last_updated: 2026-06-14
complexity: medium
delivers: corpus/ops/secrets-spec.md (net-new) + .claude/skills/stack-secrets/ + CLAUDE.md/setup_guide.md/safety.md edits
backlog_refs: (none — closes the Phase 0b KB blind area)
---

# M28 — Docs + `/stack-secrets` skill + corpus wiring

## Goal
Make the feature discoverable and give it a corpus home: author `corpus/ops/secrets-spec.md`, add the
`/stack-secrets` skill + CLAUDE.md skill-table row, retire the manual-copy prose + the `setup_guide.md:447`
TODO, extend `safety.md`.

## Why section
Pure authoring + wiring against a finished engine (M27). The doc/skill patterns to mirror already exist
(`seeding-spec.md` + `/stack-seed`, `snapshot-spec.md` + `/stack-snapshot`). Build with
`/developer-kit:build-milestone`.

## Repo split
- **`rosetta`**: `corpus/ops/secrets-spec.md` (net-new), `.claude/skills/stack-secrets/SKILL.md`, the CLAUDE.md
  skill-table + doc-index + interconnected-docs rows, the `setup_guide.md` + `safety.md` edits.
- **`rosetta-extensions`**: none new (the skill drives the M26/M27 binary at its pinned tag).

## Scope
- **In:**
  - Author **`corpus/ops/secrets-spec.md`** — the blueprint the skill reads: the source-dir/zip layout, the
    secret-DNA, the per-repo target-file map, the values-blind safety statement, the alias/collision rules, the
    `waived`-class rationale.
  - New **`/stack-secrets`** skill (`argument-hint [dev-N|demo-N] [--from DIR|ZIP] [--check|--provision|--status]`,
    default source `.agentspace/secrets`), mirroring `/stack-seed` (read spec → confirm non-prod target → build the
    tagged-clone binary → run the verb → report values-blind).
  - **CLAUDE.md** skill-table row + Key-Documentation-Locations entry + Interconnected-Documentation list update.
  - **Extend** `setup_guide.md` (delete the manual-copy prose + the line-447 TODO, point to the skill) and
    `safety.md` (add the never-echo / `PreflightEnv`-emitting clause to the safety contract).
- **Out:** the build-from-stack-dev observable-behavior validation (M29).

## Depends on
M27 (the engine + gate the docs/skill describe).

## Parallel with
None.

## Estimated complexity
medium.

## Open questions
- One skill (author + measure) vs an `/align`-style pair. **Default:** one `/stack-secrets` skill; pre-flight
  `check` rides inside `/dev-up` + `/demo-up`.

## KB dependencies
- `corpus/ops/seeding-spec.md` + `corpus/ops/snapshot-spec.md` — the spec-doc + skill pattern to mirror.
- `CLAUDE.md` — the skill-table + doc-index conventions.

## Delivers →
`corpus/ops/secrets-spec.md` (net-new) + `.claude/skills/stack-secrets/` + CLAUDE.md / setup_guide.md / safety.md edits.
