---
title: "KB Fidelity Audit — M44 profile completeness"
date: 2026-06-26
scope: milestone:M44
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Verified-skill chain + `user_level` self-rating | `corpus/ops/demo/stories-spec.md` §"claimed-vs-verified gap" | `stack-seeding/seeders/persona.go`, `persona_write.go` | PAIRED |
| Profile depth (timeline + claimed tail) | `stories-spec.md` §M41, `corpus/ops/seeding-spec.md` | `stack-seeding/seeders/profile.go` | PAIRED |
| Seeding isolation contract + write surfaces | `seeding-spec.md` §"production-isolation boundary" | `stack-seeding/isolation/`, `seeders/users.go` | PAIRED |
| Coverage semantic gate / Playwright per-section asserts | `corpus/ops/demo/coverage-protocol.md` | `stack-verify/e2e/` | PAIRED |
| Photo avatar on every member | `coverage-protocol.md` (persona self-consistency), roster.go header | `seeders/avatar.go`, `seeders/users.go:156` | PAIRED (CODE present + doc accurate) |
| `user_certificates` / `user_projects` seeding surfaces | — (M44 `Delivers →`) | none yet (M44 authors) | DOC-ONLY-to-be (milestone deliverable, NOT blind) |
| "complete profile" rubric (`profile-completeness-spec.md`) | `corpus/ops/demo/profile-completeness-spec.md` (NEW, M44 Delivers) | — | DOC-ONLY (milestone deliverable) |

## Fidelity Findings
1. **`user_level` mechanic (stories-spec.md §"claimed-vs-verified gap")** — Expected: `PersonaSeeder` sets `user_skill_evidences.user_level` per `self_eval_bias`. Actual: `persona.go:212 selfEvalLevel()` + `persona_write.go upsertEvidenceSQL()` set `user_level` on both INSERT and UPDATE paths. **ALIGNED.**
2. **Manager skip (stories-spec.md l.214 "a manager has no personal timeline — skipped"; seeding-spec.md l.282)** — Expected: managers skipped in PersonaSeeder + ProfileSeeder. Actual: `persona.go:121` + `profile.go:125` both `if p.IsManager() { continue }`. **ALIGNED** — this is current-truth that M44 §C will deliberately CHANGE (and the docs update is part of M44's `Delivers →`), not a stale claim.
3. **ProfileSeeder surfaces (seeding-spec.md l.281-297)** — Expected: companies + user_experiences + user_educations + claimed-tail user_skills/evidences, all `PerStackIsolated`, with the live-schema landmines (company uuid NOT NULL FK, DATE from/to, lowercase location_type enum, legacy_skills JSON envelope). Actual: matches `profile.go` exactly. **ALIGNED.**
4. **Photo avatar coverage** — Expected (overview §D, drawn from the pre-M42e G4 grounding): "extend `photoAvatarDataURI` to EVERY member (not heroes-only)". Actual: `users.go:156` ALREADY calls `photoAvatarDataURI(uid)` for every population user since M42e P4. **The avatar half of §D is already satisfied in code** — §D reduces to verification + the bulk-member shallow career. Recorded as a build decision (not a doc-staleness blocker; the overview reflects an older grounding doc, not a wrong doc claim).

## Completeness Gaps
- None critical. The new write surfaces (`user_certificates`, `user_projects`) and the manager-personal-data / bulk-member-career extensions have no existing doc anchor by design — they are the milestone's own deliverables (`profile-completeness-spec.md` NEW + updates to `seeding-spec.md` / `stories-spec.md`), authored in M44's Docs section.

## Applied Fixes
- None needed inline. Triples recorded in `spec-notes.md`.

## Open Items (require user decision)
- None.

## Gate Result
GREEN — every milestone-dependency topic is PAIRED and ALIGNED; the only un-doc'd topics are M44's own `Delivers →`. `build-milestone` may proceed to Phase 1.
