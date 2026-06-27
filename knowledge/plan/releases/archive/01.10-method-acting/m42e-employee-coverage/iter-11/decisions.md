# iter-11 decisions (P0 baseline)

## B1 — P0 before-state (live demo-3, Maya Chen `d26e0467-31d9-5f73-96ad-525dfbb38d11`)

Probed via `docker exec demo-3-postgresql-1 psql`. The 4 P1–P3 roots all confirmed live:

- **P1 (skill draw):** Maya's 30 distinct VERIFIED skills = **10 coherent role-core** (Agile Methodologies,
  Data Structures and Algorithms, Java, Microservices Architecture, Node.js, Performance Optimization, Python,
  RESTful Services and APIs, Server-side Programming Languages, SQL) **+ 20 flat-pool JUNK** (15Five, 17Track,
  24-hour dietary recall, 25Live, 2Checkout, 2D/3D Animation, 2D Materials Engineering, 3D Bioprinting in
  Dentistry, 3dcart, 360° Video for Exposure Therapy, …). Exactly the `resolveHeroSkills` flat-pool
  `ORDER BY node_id` head bug — `verified: 30` in the preset vs a 10-skill Backend Developer role ⇒ 20 junk
  top-ups. (Also 60 claimed-but-unverified, same junk-head source via `combinedNamedPool`.)
- **P2 (skeleton):** 3 work_experiences + 1 user_education EXIST with real role-coherent skill NAMES — but the
  `skills` json is a **bare array** `["Java","Microservices Architecture",…]`, while backend
  `skills.LegacySkills` demands `{"skills":[{"level":N,"name":"Java"}]}`. The federation unmarshal error nulls
  `timelineGrouped` → /profile + /home perpetual skeleton. High-confidence P2 fix.
- **P3 (activity):** `personal_assignments` = **0** for Maya (→ /home no path pills, Paths=0);
  skillpath sessions = **0 completed** (1 active + 2 in_progress version-"1" seeded + 21 pending version-today
  runtime-enrolled; max progress 78) → "Skill Paths Completed"=0; `user_bookmarks` = **0**.

Schemas (demo-3, all `public` unless noted):
- `personal_assignments(id, created_at, updated_at, due_date, status='active', resource_id uuid NOT NULL,
  resource_type, user_id uuid NOT NULL)`.
- `personal_assignment_sessions(id, created_at, updated_at, assignment_id, user_id, session_id, progress bigint,
  started_at, ended_at)`.
- `user_bookmarks(id, created_at, resource_id uuid NOT NULL, resource_type, user_id)`.
- `skillpath.skill_path_sessions` (the SkillpathSessionsSeeder target) — make ≥1 of Maya's seeded rows
  `completed` (progress=100).

## B2 — specialized-role candidate analysis (the user-facing P1 character decision)

EVERY public job_role caps at exactly **10 role-skills** (the taxonomy is 10-core-per-role) — so specialization
does NOT add skill depth; it adds COHERENCE + a senior title. Compared the 10-skill SETS of the candidates:

| Role | Character of its 10 skills |
|---|---|
| Backend Developer (current) | Java/Python/SQL/Node.js/Microservices/REST/DSA/Server-side/Perf/Agile — solid but generic-junior |
| **Backend Software Engineer** | + Docker/Kubernetes, CI/CD, Cloud (AWS/Azure/GCP), API Dev, Unit Testing, SQL+NoSQL — a modern SENIOR backend stack |
| Distributed Systems Engineer | Kafka, Distributed Algorithms, Distributed System Design, High Availability, Scalable Architecture, Consistency — specialized senior |
| Software Architect | System/Scalable Architecture Design, Technical Leadership, SDLC, Cloud — principal/architect |
| Senior Software Engineer | + Code Review, CI/CD, Debugging, Software Systems Design |

**CHOSEN: Backend Software Engineer** (see B3 in iter-12). It is the strongest believable SENIOR-backend
specialization: a coherent modern stack (Kubernetes, CI/CD, Cloud, NoSQL, Unit Testing) that keeps
Java/Node.js/SQL continuity with the current role (so the work-history stays coherent) AND extends naturally to
the curated claimed tail (Kafka, gRPC→Redis, Terraform, Observability, System Design…). Surfaced for the user to
confirm/tweak.

## B3 — curated software claimed-pool resolves in real taxonomy (no-fabrication-safe)

Of the design-plan's curated software claimed names, **27 resolve** to real public node-ids (so closure stays
green): Amazon Web Services (AWS) K-AMAWEB-7855, API Development and Integration K-APIDEV-F24A, Caching
Strategies K-CACSTR-8F8B, Cloud Platform Expertise K-CLOPLA-EEA1, Code Review K-CODREV-0524, Containerization
(Docker, Kubernetes) K-CONDOC-A287, CI/CD K-CONINT-5045, Database Administration and Management K-DATADM-8B7C,
Debugging K-DEBUGG-BC02, Distributed Algorithms K-DISALG-973C, Distributed System Design K-DISSYS-9981, GraphQL
K-GRAPHQ-32FD, High Availability Design K-HIGAVA-0363, Kafka K-KAFKAX-3C46, Load Balancing K-LOABAL-7500,
Microservices Architecture K-MICARC-F5A5, Performance Tuning K-PERTUN-99CC, PostgreSQL K-POSTGR-E59F, Redis
K-REDISX-E92E, Scalable Architecture Design K-SCAARC-DEE6, Secure SDLC K-SECSOF-2EA7, SQL and NoSQL Database
Management K-SQLNOS-B06A, System Architecture Design K-SYSARC-7700, System Design K-SYSDES-EC18, Technical
Leadership K-TECLEA-FC01, Terraform K-TERRAF-433D, Unit Testing K-UNITES-9263.

NON-resolving (would be fabrications — DO NOT use by name): bare "Kubernetes", "gRPC", "Domain-Driven Design",
"Observability", "Event-Driven Architecture", "Message Queuing", "Infrastructure as Code", bare "Docker". The
allow-list is resolved by NAME against the live taxonomy at run time, so a non-resolving name is silently
dropped (never fabricated) — closure stays green.
