# iter-13 decisions (P2)

## D1 — the exact LegacySkills struct shape (confirmed from the platform schema, read-only)
`app/internal/data/ent/skills/skills.go`:
```go
type LegacySkills struct { Skills []*LegacySkill `json:"skills"` }
type LegacySkill  struct { Level int `json:"level"`; Name string `json:"name"` }
```
The ent `legacy_skills` field (`mixin.go LegacySkillsMixin`, StorageKey "skills", default `'{}'`) is mixed into
user_experience AND user_education (also user_content/volunteer/project/certification — same envelope). So the
`skills` column MUST be `{"skills":[{"level":N,"name":"Go"}]}`, NOT a bare `["Go"]`. A bare array fails to
unmarshal → the Experience/Education GraphQL resolver errors → the federated `timelineGrouped` nulls → /profile
+ /home perpetual skeleton (even though the rows exist). Read-only diagnosis; zero platform edit.

## D2 — legacySkillsJSON envelope (the fix)
`profile.go`: NEW `legacySkillsJSON([]string)` emits the envelope with a fixed `level=3` (a credible mid claim;
the believable claimed-vs-verified gap is carried by the user_skills tail, not per-entry timeline levels) +
escaped names (a quote/backslash in a name would break the whole timeline read). Replaced both writes
(experiences :~186, educations :~213); empty slice → `{"skills":[]}` (a valid empty envelope, not `[]`).
Removed the now-orphaned `jsonStringArray` (the bare-array writer). Updated the schema-landmine comment (:~40).
Tests: `profile_test.go` asserts the rows round-trip into the LegacySkills struct shape (level 1–5, non-empty
name); the escaping harden test repointed from `jsonStringArray` to `legacySkillsJSON` (envelope + escaping +
empty-envelope). `go test ./...` GREEN; `go vet` clean.

## Measurement (live demo-3)
Re-seed is idempotent (ON CONFLICT DO NOTHING), so a hero's OLD bare-array timeline rows aren't overwritten —
to MEASURE the fix on the live demo, each end-user hero's claimed-tail user_skills + evidences (which FK the
experiences via user_skill_experience/education — the DB CHECK blocks deleting an experience while a claimed
skill points at it) were deleted FIRST, then the experiences + educations, then re-seeded. Order matters:
claimed user_skills + evidences → experiences/educations → re-seed. Final stack-wide: 0 bare-array, 12 envelope
experiences + 4 envelope educations, closure PASS. The AUTHORITATIVE clean reproduce is the P8 fresh demo-up.

## Result (measured on demo-3)
Maya's experiences now read e.g.
`{"skills":[{"level":3,"name":"Cloud Platform Expertise (e.g., AWS, Azure, GCP)"},{"level":3,"name":"Containerization (Docker, Kubernetes)"},…]}`
— the timeline resolver parses it, `timelineGrouped` populates, and the 3 work + 1 education entries render
(no skeleton). Titles are role-coherent ("Backend Software Engineer" / "Junior Backend Software Engineer") from
the P1 specialization.
