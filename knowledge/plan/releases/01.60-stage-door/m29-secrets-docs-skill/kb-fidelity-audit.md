---
title: "KB Fidelity Audit — M29 (Docs + /stack-secrets skill + corpus wiring)"
date: 2026-06-14
scope: milestone:M29
invoked-by: build-milestone
---

## Verdict
GREEN (blind area is the milestone's own deliverable — covered by the `delivers:` line)

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| secret provisioning + secret-coverage DNA | `corpus/ops/secrets-spec.md` (NET-NEW, this milestone authors it) | `.agentspace/rosetta-extensions/stack-secrets/` (provision/, secretdna/, source/, cmd/stacksecrets) | BLIND-AREA → covered by `delivers:` |
| `/stack-secrets` skill | `.claude/skills/stack-secrets/SKILL.md` (NET-NEW, this milestone authors it) | drives `stacksecrets` @ tag `stage-door-m28` | DOC-ONLY (to author) → covered by `delivers:` |
| spec-doc pattern to mirror | `corpus/ops/seeding-spec.md`, `corpus/ops/snapshot-spec.md` | `stack-seeding/`, `stack-snapshot/` | PAIRED (reference only — not modified) |
| skill pattern to mirror | `.claude/skills/stack-seed/SKILL.md` | `stack-seeding/cmd/stackseed` | PAIRED (reference only) |
| CLAUDE.md skill-table + doc-index conventions | `CLAUDE.md` | n/a | PAIRED (extend) |
| values-blind safety contract | `corpus/ops/safety.md` | `stack-seeding/isolation/`, `stack-secrets/provision/` | PAIRED (extend with a new clause; §2.8 already added by M28-harden) |
| manual-copy prose + line-447 TODO | `corpus/ops/setup_guide.md` | n/a | PAIRED (retire prose, point to skill) |
| README-index guard | `corpus/README.md` (the index) | developer-kit guard | PAIRED (must index secrets-spec.md) |

## Fidelity Findings
None. The reference docs (`seeding-spec.md`, `snapshot-spec.md`, `stack-seed/SKILL.md`, `safety.md`) and CLAUDE.md
are all present, current, and accurately describe their code. The ext `stack-secrets` engine at tag `stage-door-m28`
(`9742126`) is the authoritative source the new doc/skill describe; its README + package doc-comments are the
ground-truth (provision blanks `DIRECTUS_TOKEN` on non-prod, demo overlay mints Clerk keys, N=0 guard, alias families,
waived classes, the 0/1/3 exit contract). No stale claim exists because the doc doesn't exist yet.

## Completeness Gaps
The blind area itself: secret provisioning is a fully-built engine (M27/M28, 160 Go tests, tag `stage-door-m28`)
with **zero corpus coverage** — only incidental "secret" mentions in setup_guide/safety/directus-local/staging-*/
webhook_setup. M29 closes exactly this gap, which is its declared purpose.

## Applied Fixes
None applied during the audit (no stale claims to fix). The blind area is resolved by the milestone's own deliverables,
not by an inline backfill: the milestone `overview.md` frontmatter carries
`delivers: corpus/ops/secrets-spec.md (net-new) + .claude/skills/stack-secrets/ + CLAUDE.md/setup_guide.md/safety.md edits`,
and the roadmap entry `Delivers →` names `corpus/ops/secrets-spec.md`. This is the sanctioned BLIND-AREA resolution
(Phase 4: "add a `Delivers → knowledge/{path}` line" — already present), not a RED block.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed. The only finding is the secret-provisioning blind area, which is the milestone's declared deliverable
(`delivers:` line present in `overview.md` + `Delivers →` in the roadmap). M29 authoring `secrets-spec.md` + the
`/stack-secrets` skill IS the backfill. No stale load-bearing claims, no critical undocumented behavior in code the
milestone modifies (M29 modifies docs/skills, not the ext engine). Build may enter Phase 1.
