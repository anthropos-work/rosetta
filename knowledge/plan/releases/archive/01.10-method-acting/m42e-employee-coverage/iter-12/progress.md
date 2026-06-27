# iter-12 progress

**Type:** tik (production-fix). Active strategy: **TOK-10**. P1 — core skill-draw + role specialization.

## Phase C — fix (rext, zero platform edit)
- NEW `curated_pools.go` (D1): per-category curated skill-NAME allow-lists (software / sales), role→category
  classifier, `resolveCuratedPools` (resolve names → real node-ids in allow-list order, drop non-resolvers).
- `persona.go resolveHeroSkills` + `profile.go combinedNamedPool` top up CURATED-first then flat (last resort).
- `taxonomyref.go` / `skillref_named.go` thread `curated curatedPools`.
- `profile.go`: claimed tail honors `EffectiveMapped()` per-hero via `claimedTailCount` (D3).
- SPECIALIZE Maya → Backend Software Engineer; curate counts (verified 12, mapped 18). Sara → curated sales
  pool (verified 12, mapped 14). (D2, D4 — stories.seed.yaml.)
- Tests: NEW `curated_pools_test.go` (classifier, allow-list-order resolution, drop-not-fabricate,
  curated-before-flat); `profile_test.go` per-hero tail. `go test ./...` GREEN; `go vet` clean.

## Phase D — re-measure (demo-3 re-seed, measurement-only clear-then-seed)
- Maya: 12 verified (Agile/API Dev/Cloud/Docker-K8s/CI-CD/Kafka/Node.js/Problem Solving/Programming
  Languages/Redis/SQL+NoSQL/Unit Testing) + 18 claimed (AWS/Caching/Code Review/DB Admin/Debugging/Distributed
  Algorithms+Design/GraphQL/HA Design/Load Balancing/Microservices/Perf Tuning/PostgreSQL/Scalable+System
  Arch/Secure SDLC/System Design/Terraform). **ZERO junk** (was 20/30).
- Sara: 12 verified + 14 claimed, all sales-coherent. **ZERO junk**.
- `datadna measure-closure --stack demo-3`: **PASS** for both heroes (no dangling refs — no fabrication).
- Re-seed isolation: clean (prod=false).

## Close — 2026-06-25
**Outcome:** P1 landed. Maya = a believable senior **Backend Software Engineer** (12 coherent verified + 18
coherent claimed); Sara = a coherent Account Executive (12 + 14). All junk gone; closure PASS.
**Type:** tik (production-fix)
**Status:** closed-fixed
**Gate:** NOT MET (P1 of P0–P8; the believability gate needs P4–P8 + the P7 semantic harness)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik #2) — (6) protocol-stop: n — Outcome: continue (into iter-13 P2)
**Decisions:** D1–D4 — see ./decisions.md.
**Side-deliverables:** `mapped:` field now load-bearing for the claimed tail (per-hero).
**Routes carried forward:** P2 (iter-13 — json shape), P3 (iter-14 — activity).
**Lessons:** curated-by-NAME allow-list = no-fabrication-safe coherence past the 10-skill role cap; size
verified+claimed ≤ role∪curated coherent union to avoid flat-junk spill; `mapped:` drives the tail per-hero.
**rext:** commit `19cab88`, tag `method-acting-m42e-iter12`.
