# iter-12 decisions (P1)

## D1 — curated_pools.go (the no-fabrication-safe coherent top-up)
NEW `stack-seeding/seeders/curated_pools.go`: per-category (software / sales) curated skill-NAME allow-lists,
a role→category classifier (`curatedCategoryForRole`), and `resolveCuratedPools` which resolves each list's
NAMES → real public (node_id, name) **in allow-list order**, dropping any name that doesn't resolve. Threaded
into `taxonomyRefs.curated` (persona side) + `namedSkillRefs.curated` (profile side), resolved once per run.
`resolveHeroSkills` (persona.go) + `combinedNamedPool` (profile.go) now top up CURATED-first, flat-pool last.
Closure stays GREEN (every emitted node-id is real; a non-resolving curated name is simply absent — never
fabricated). Proven by `datadna measure-closure` PASS on demo-3.

## D2 — SPECIALIZE Maya: Backend Developer → Backend Software Engineer (the user-facing character choice)
Per the user decision (specialize, not align) + iter-11 B2. Backend Software Engineer is a coherent senior-
backend public role (10 role-skills: Agile, API Dev, Cloud Platform Expertise, Containerization (Docker, K8s),
CI/CD, Node.js, Problem Solving, Programming Languages, SQL+NoSQL, Unit Testing) that keeps continuity with the
old Backend Developer skills (so the work-history stays coherent) and extends naturally to the curated claimed
tail (Kafka, Redis, Terraform, GraphQL, Distributed System Design, System Design…). NAMED in the run output for
the user to confirm/tweak.

**Maya's resulting coherent skill set (measured on demo-3):**
- 12 VERIFIED: Agile Methodologies, API Development and Integration, Cloud Platform Expertise (AWS/Azure/GCP),
  Containerization (Docker, Kubernetes), CI/CD, Kafka, Node.js, Problem Solving, Programming Languages, Redis,
  SQL and NoSQL Database Management, Unit Testing.
- 18 CLAIMED (unverified, the gap): Amazon Web Services (AWS), Caching Strategies, Code Review, Database
  Administration and Management, Debugging, Distributed Algorithms, Distributed System Design, GraphQL, High
  Availability Design, Load Balancing, Microservices Architecture, Performance Tuning, PostgreSQL, Scalable
  Architecture Design, Secure SDLC, System Architecture Design, System Design, Terraform.
- ZERO junk (was 20/30 junk before). Preset: `verified: 12, mapped: 18`.

## D3 — the claimed tail honors EffectiveMapped() per-hero (was a flat const 60)
`claimedTailCount(p)` resolves the per-hero claimed count from the blueprint's `mapped:` field (Maya 18,
Tom 18, Sara 14, Nick 16) instead of the hardcoded `claimedTailPerHero = 60`. A believable per-hero
claimed-vs-verified gap, not a 60-skill junk-padded tail. `profile_test.go` updated to the per-hero semantics
(the two tail-magnitude assertions now check `claimedTailCount(hero)`).

## D4 — sales-pool generalization (Sara — Account Executive)
The curated sales allow-list (24 resolving clean sales skills: Prospecting, Lead Generation, Cold Calling,
Negotiation, Account Planning/Lifecycle/Health/Performance Management, Client Relationship Management, Customer
Success Management, Sales Strategy, Territory Management, Persuasion…). Account Executive role ∪ sales curated =
27 coherent. Sara `verified: 28→12, mapped: 20→14` to fit. Measured: 12 verified + 14 claimed, all sales-
coherent, ZERO junk, closure PASS. The fix is general (the classifier handles software AND sales; the role
fallback handles unclassified roles unchanged).

## Sizing rule (lesson)
A hero's `verified + mapped` must stay ≤ the role's COHERENT union (role-skills ∪ curated-pool), or the
overflow spills into flat junk. Backend Software Engineer ∪ curated = 31; Account Executive ∪ curated = 27.
Widened the curated software pool (+19 resolving names: Java/Python/SQL/Node.js/Azure/GCP/MongoDB/MySQL/Git/
Data Modeling/Concurrency/…) to give headroom.

## Measurement note (live demo-3)
The persona/verified-chain rows are keyed by skill INDEX (deterministicUUID), so an idempotent re-seed
(ON CONFLICT DO NOTHING) does NOT overwrite a hero's OLD skill rows. To MEASURE the new draw cleanly on the
LIVE demo, the hero's user_skills + user_skill_evidences + verified-chain session rows were deleted first
(a measurement-only DB op on demo-3, columns: sessions.owner_id, local_jobsimulation_sessions.user_id), then
re-seeded. The AUTHORITATIVE clean reproduce is the P8 fresh demo-up (no manual delete) — the seeder produces
the coherent set from scratch there.
