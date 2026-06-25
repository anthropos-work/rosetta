---
iter: 13
iteration_type: tik
iter_shape: production-fix
status: closed-fixed
---

# iter-13 — P2: the legacy_skills json shape (the timeline skeleton fix)

**Type:** tik (production-fix). **Active strategy: TOK-10**. The design-plan's root cause #2 (highest-confidence
fix): the /profile + /home perpetual skeleton.

## Cluster / target identified
P2 from TOK-10's next-tik direction. iter-11 B1: the timeline rows EXIST (3 work + 1 education) but the `skills`
json is a BARE array `["Java",...]` while the platform ent `legacy_skills` field unmarshals into
`skills.LegacySkills{ Skills []*LegacySkill{Level int; Name string} }`, i.e. `{"skills":[{"level":N,"name":"Java"}]}`.

## Hypothesis
The bare array fails to unmarshal into the LegacySkills struct → the GraphQL timeline resolver errors → the
federated `timelineGrouped` nulls → /profile + /home render perpetual skeletons. Writing the correct envelope
makes the timeline render.

## Phase plan
A (baseline iter-11) → confirm the exact struct shape from the platform schema (read-only) → C (fix profile.go:
`legacySkillsJSON` helper + both writes + comment + tests) → re-seed demo-3 → D (re-measure: envelope shape +
stack-wide bare-array=0 + closure) → E (close).

## Close — 2026-06-25
**Outcome:** P2 landed. `legacySkillsJSON()` writes the correct `{"skills":[{"level":3,"name":"..."}]}` envelope
for experiences + educations. Measured on demo-3 (re-seeded all 4 end-user heroes): 0 bare-array timelines
stack-wide, 12 envelope experiences + 4 envelope educations, closure PASS. The timeline now renders (not
skeleton).
**Type:** tik (production-fix)
**Status:** closed-fixed
**Gate:** NOT MET (P2 of P0–P8; the believability gate needs P3–P8 + the P7 semantic harness). The skeleton
ROOT is closed.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik #3 of session) — (6) protocol-stop: n — Outcome: continue (into iter-14 P3)
**Decisions:** see ./decisions.md (D1 the exact LegacySkills struct shape from the platform schema; D2 the
legacySkillsJSON envelope + level=3 + removing the orphaned jsonStringArray).
**Routes carried forward:** P3 (iter-14 — activity: personal_assignments + completed session + bookmarks).
**Lessons:** an ent JSON field with a struct codec (here `skills.LegacySkills`) needs the EXACT envelope, not a
convenient bare array — confirm the codec struct from the platform schema (read-only) before writing a json
column the federation unmarshals; a single unmarshal error on a non-nullable resolver field nulls the whole
query (the "everything exists but renders empty" failure class — same family as the iter-07 federation-null).
