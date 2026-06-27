**Type:** tik

# iter-06 — the GeneratedBatchSeeder (cache → rows; the CODE-vs-AI boundary)

Built component (5) of TOK-01's chain — the LAST code component. The engine is now CODE-COMPLETE.

## Work done
- **NEW `seeders/generated_batch.go`** — `GeneratedBatchSeeder` (surface 'generated-batch', DependsOn
  users+taxonomy, PerStackIsolated). Reads the prompt-hash cache (the SAME batch dir cmd/gen-batch wrote
  — it rebuilds the same mother prompts via `EffectiveBatches(reservedHeroNamesForSeed)`), parses each
  cached envelope, and writes deterministic `public.users` + `public.memberships` + `public.user_skills`
  (claimed) rows via `CopyRowsIdempotent` + audit records. THE BOUNDARY: the role NAME resolves via
  `resolveJobRoleRefs.forName` (non-resolving → NULL role label via nodeIDOrNil/nameOrNil, the
  batch-assigned role as a fallback, NEVER a fabricated J-); each claimed-skill NAME resolves via the
  named-skill pools (`resolveClaimedSkillNames`: role pool then flat pool by case-insensitive name; a
  non-resolving name is DROPPED — no row). Generated members occupy a HIGH index band
  (generatedBaseIndex=100000 + per-batch band) so they never collide with curated users. Email is the
  envelope's email_local (sanitized) or name-derived.
- **`cmd/stackseed/main.go`** — registered `GeneratedBatchSeeder{}` in the DAG (zero-value → default
  cache root + unversioned; the standard flow).
- **NEW `seeders/generated_batch_test.go`** — 9 tests (fake Conn + temp cache, no key/DB): seeds members
  (3 users/3 memberships/9 claimed skills, the real resolved role node-id), DROPS non-resolving claimed
  skills (only the real one rows, with a REAL node-id), blank-label on a non-resolving role (assigned-role
  fallback) + NULL on a truly-unresolvable role, drops un-generated (cache-miss) members, no-batch no-op,
  empty-cache→0 rows, the surface/deps/isolation contract, the audit-records-every-write provability.

## Measurements
- `go build ./...` + `go vet ./...` clean; `gofmt -l` clean.
- `go test ./seeders/... -run TestGeneratedBatch` GREEN; full `go test ./...` GREEN; `-race ./seeders/...`
  GREEN. Test funcs 623 -> **632** (+9).

## Close — 2026-06-26

**Outcome:** Component (5) landed — the `GeneratedBatchSeeder` closes the cache→DB path and ENFORCES the
CODE-owns-structure / AI-owns-content boundary (drop-not-fabricate, unit-proven: a hallucinated skill
produces NO row; a real one a row with a REAL node-id; a non-resolving role → NULL label). The engine is
now CODE-COMPLETE + fully fixture/fake-proven. Gate still 0/5 EMPIRICALLY (the real valid-JSON rate +
resolution + collision + cost + live $0-reseed are the NEXT call's gate-proving run), but every component
exists + every gate metric is measurable.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (0/5 empirically — engine CODE-COMPLETE + fixture/fake-proven; the REAL capped batch + the live $0-reseed on demo-3 remain — the NEXT call)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: Y (5th tik of the session — iter-02..06) — (6) protocol-stop: n — Outcome: exit-5
**Decisions:** D1 (reuse resolvers unchanged), D2 (high index band, no curated collision), D3 (first-org for the bounded batch — per-story routing is M46)
**Side-deliverables (if any):** none
**Routes carried forward:** NEXT CALL (a fresh tik after the cap) — the EMPIRICAL gate-proving: run a REAL
N=20 capped `gen-batch` (OPENAI_KEY from stack-demo/platform/.env, values-blind, --max-cost ~$0.10,
gpt-4o-mini) → measure valid-JSON rate (≥95%?), taxonomy-resolution (closure green, 0 fabrication),
0 hero-collision, cost ≤ ceiling; then the byte-identical $0 re-seed; then re-seed on a fresh demo-3 via
the tagged consumption clone. THEN the gate is gradable empirically.
**Lessons:** the resolvers' existing zero-value-on-no-match contract is EXACTLY the drop-not-fabricate
seam — the seeder needed ZERO resolver changes, just "skip the zero value." Reusing the `namedSkillConn`
test fake (already serves the 6-array named-skill + 2-array job-role reads) made the seeder testable
without a new fake. The generated-member index band (100000+) is the clean way to avoid curated-user
collision without coordinating indices.
