---
iter: 12
iteration_type: tik
iter_shape: production-fix
status: closed-fixed
---

# iter-12 — P1: core skill-draw fix + specialize Maya's role

**Type:** tik (production-fix). **Active strategy: TOK-10** (persona-believability-by-root-cause). The
design-plan's root cause #1 + the user decision to SPECIALIZE Maya's role.

## Cluster / target identified
P1 from TOK-10's next-tik direction: the flat-pool junk-head top-up (`resolveHeroSkills` persona.go +
`combinedNamedPool` profile.go drawing past the 10 role-skills from `tax.flat ORDER BY node_id` → 15Five /
3D Dental Anatomy / 3dcart). iter-11 B1 confirmed 20/30 of Maya's verified skills were junk. Plus the user
decision: SPECIALIZE Maya's role (not align).

## Hypothesis
A curated, category-coherent skill-NAME allow-list (software / sales), resolved against the real taxonomy at
run time and topped up BEFORE the flat pool, replaces the junk with coherent skills while keeping closure green.
Specializing Maya to a richer senior role + curating her counts to fit the coherent pool makes her a believable
senior backend engineer.

## Phase plan
A (baseline from iter-11) → C (fix in rext: curated_pools.go + persona/profile top-up + role specialize) →
test → re-seed demo-3 (measurement) → D (re-measure: Maya + Sara skills + closure) → E (close).

## Close — 2026-06-25
**Outcome:** P1 landed. Maya specialized to **Backend Software Engineer**; the curated-pool top-up replaced
ALL junk. Measured on demo-3: Maya = 12 verified + 18 claimed, Sara = 12 verified + 14 claimed — all coherent,
zero junk; `datadna measure-closure` PASS for both (no fabrication).
**Type:** tik (production-fix)
**Status:** closed-fixed
**Gate:** NOT MET (P1 of P0–P8; the believability gate needs P4–P8 + the P7 semantic harness). The P1 ROOT is
closed — the skill-coherence dimension of persona self-consistency now holds for Maya + Sara.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik #2 of session) — (6) protocol-stop: n — Outcome: continue (into iter-13 P2)
**Decisions:** see ./decisions.md (D1 curated_pools design; D2 role-specialization to Backend Software
Engineer; D3 per-hero claimed tail via EffectiveMapped; D4 sales-pool generalization for Sara).
**Side-deliverables:** the `mapped:` blueprint field is now LOAD-BEARING for the claimed tail (was a flat 60),
a coherent improvement beyond the strict P1 scope but in-line with the believability gate.
**Routes carried forward:** P2 (iter-13 — the legacy_skills json shape / skeleton), P3 (iter-14 — activity).
The aspirational curated names (Kubernetes/gRPC/Observability/DDD) that don't yet resolve are kept in the
allow-list for a future taxonomy capture (they drop harmlessly today).
**Lessons:** (1) a curated allow-list resolved-by-NAME against the real taxonomy is the no-fabrication-safe way
to get COHERENT skills past a role's 10-skill cap (closure stays green; a non-resolving name simply drops).
(2) Size a hero's verified+claimed counts to the role's COHERENT union (role ∪ curated), or the overflow spills
into flat junk — Backend Software Engineer ∪ curated = 31, Account Executive ∪ curated = 27; keep verified+claimed
under that. (3) The `mapped:` field should drive the claimed tail per-hero (not a flat const).
