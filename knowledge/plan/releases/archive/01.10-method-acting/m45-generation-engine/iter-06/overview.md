---
iter: 06
milestone: M45
iteration_type: tik
status: closed-fixed
date: 2026-06-26
---

# iter-06 — tik — the `GeneratedBatchSeeder` (component 5)

## Active strategy reference
TOK-01 (inside-out fixtures-first build). Fifth tik (the cap fires after this iter).

## Re-survey
TOK-01 named the `GeneratedBatchSeeder` as iter-06's target; not absorbed. Target current.

## Cluster / target identified
The seeder is the LAST code component before the real gate-proving run: it closes the cache→DB path and
is WHERE the CODE-owns-structure / AI-owns-content boundary is enforced (the drop-not-fabricate seam). It
reuses the existing resolvers (`resolveJobRoleRefs` / `resolveNamedSkillRefs`) unchanged + the Seeder/DAG
machinery, so it slots in as a registered surface.

## Hypothesis
A `GeneratedBatchSeeder` (surface 'generated-batch', DependsOn users+taxonomy, PerStackIsolated) that
reads the cache, parses each envelope, resolves the role + claimed-skill NAMES through the existing
resolvers (a non-resolving role → blank label, never a fabricated J-; a non-resolving skill → DROPPED, no
row), and writes deterministic users/memberships/user_skills rows — enforces the boundary correctly and is
unit-testable against a fake Conn + a temp cache (no key, no DB).

## Expected lift
Infrastructure + the boundary enforcement. No empirical valid-JSON-rate move (that's the REAL run). The
drop-not-fabricate seam + the closure-green guarantee land + are unit-proven (a hallucinated skill name
produces NO row; a real one produces a row with a REAL node-id).

## Phase plan
Protocol §4b: build (component 5) → register in cmd/stackseed → unit-test (seeds members, drops
non-resolving skills, blank-label/nil unresolvable role, drops un-generated members, no-batch no-op,
empty-cache, the surface/deps/isolation contract, the audit records) → full-suite + race regression →
close.

## Escalation conditions
- The resolver call path can't enforce drop-not-fabricate without a resolver change → user-blocker (a
  resolver change is out of M45's reuse-unchanged scope). (Did not happen — the resolvers already return
  zero values for non-matches; the seeder just skips them.)

## Acceptable close-no-lift outcomes
N/A — build tik; closes `closed-fixed` on the seeder landing green.

## Note (cap)
This is tik #5 of the session. After it closes, the 5-tik cap fires (Phase 5 §5) → the call exits
`cap-reached`. The engine is now CODE-COMPLETE (services/ai → blueprint.Batch → cache → cmd/gen-batch →
GeneratedBatchSeeder, all unit-proven against fixtures). The NEXT call does the empirical gate-proving:
the REAL N=20 capped batch (OPENAI_KEY from stack-demo/platform/.env, values-blind) measuring the
valid-JSON rate + resolution + 0-collision + cost-bounded, then the live $0 reseed on demo-3 via the
tagged consumption clone.
