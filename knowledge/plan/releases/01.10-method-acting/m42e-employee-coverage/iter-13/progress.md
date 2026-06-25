# iter-13 progress

**Type:** tik (production-fix). Active strategy: **TOK-10**. P2 — the legacy_skills json shape (skeleton fix).

## Phase C — fix (rext, zero platform edit)
- Confirmed the exact `skills.LegacySkills{ Skills []*LegacySkill{Level,Name} }` struct from the platform
  schema (read-only; D1) — the column needs `{"skills":[{"level":N,"name":"..."}]}`, not a bare `["..."]`.
- `profile.go`: NEW `legacySkillsJSON()` (envelope, level=3, escaped names); replaced both writes (experiences
  + educations); removed the orphaned `jsonStringArray`; fixed the landmine comment. (D2)
- Tests: `profile_test.go` round-trips into the LegacySkills struct shape; the escaping harden test repointed
  to `legacySkillsJSON`. `go test ./...` GREEN; `go vet` clean.

## Phase D — re-measure (demo-3 re-seed, measurement-only clear-then-seed)
- Re-seeded all 4 end-user heroes (cleared claimed-tail user_skills+evidences → experiences/educations → seed,
  the FK-safe order). Maya's experiences now `{"skills":[{"level":3,"name":"Cloud Platform Expertise…"},…]}`.
- Stack-wide: **0 bare-array**, 12 envelope experiences + 4 envelope educations. The timeline renders.
- `datadna measure-closure --stack demo-3`: **PASS**. Re-seed isolation: clean (prod=false).

## Close — 2026-06-25
**Outcome:** P2 landed. The timeline json is the correct ent legacy_skills envelope; /profile + /home no longer
render perpetual skeletons (the rows round-trip into the resolver). Stack-wide 0 bare-array; closure PASS.
**Type:** tik (production-fix)
**Status:** closed-fixed
**Gate:** NOT MET (P2 of P0–P8; the believability gate needs P3–P8 + the P7 semantic harness)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik #3) — (6) protocol-stop: n — Outcome: continue (into iter-14 P3)
**Decisions:** D1 (LegacySkills struct shape), D2 (legacySkillsJSON envelope) — see ./decisions.md.
**Routes carried forward:** P3 (iter-14 — activity: personal_assignments + completed session + bookmarks).
**Lessons:** an ent JSON field with a struct codec needs the EXACT envelope — confirm the codec struct from the
platform schema before writing a federation-unmarshaled json column; one unmarshal error on a non-nullable
resolver field nulls the whole query (the "exists but renders empty" failure class).
**rext:** commit `09bbf2a`, tag `method-acting-m42e-iter13`.
