---
iteration_type: tik
iter_shape: production-fix
status: planned
created: 2026-06-25
---

# iter-15 — T1: RE-SPECIALIZE Maya → DevOps Engineer (user-confirmed Platform/DevOps)

**Active strategy reference:** TOK-10 (persona-believability-by-root-cause). This tik executes the
run-5 T1 task: re-specialize Maya from "Backend Software Engineer" (run-4's too-mild pick) to a
**Platform / DevOps** role per the user's confirmed decision (design-plan §USER DECISIONS #1, answered
2026-06-25).

**Step 0 re-survey:** run-4 landed the coherent-skill machinery (curated_pools.go + persona/profile
top-up) at role "Backend Software Engineer". The TOK-10 next-direction named P0→P3 (done). The user's
ANSWERED decision re-targets the role to Platform/DevOps. Confirmed live on demo-3: Maya = "Backend
Software Engineer", picture = SVG cartoon, org logo_url NULL. The curated machinery is in place; T1
re-points it at a DevOps role + DevOps sub-pool. Not a re-scope — same TOK-10 strategy, the
user-answered role target.

**Cluster / target identified:** Maya's role label + skill set + bio/CV are coded around backend
software. The user chose Platform/DevOps. Queried the REAL public taxonomy: every job_role caps at
exactly 10 role-core skills (so specialization = coherence + title, not depth — the curated sub-pool
drives breadth). **Chose role = "DevOps Engineer"** — its 10 role-core skills are the cleanest
platform/DevOps set (Containerization (Docker, Kubernetes), CI/CD Pipelines, Infrastructure-as-Code
(IaC), Automation Tools, Scripting, Version Control/Git, Security Best Practices, Performance Tuning,
Troubleshooting and Debugging, Agile Methodologies). Resolved a ~40-name DevOps curated allow-list
against the live taxonomy (Ansible/Terraform/Jenkins/IaC/Incident Response/Cloud Architecture
Design/Cloud Cost Management/Secrets Management/Monitoring and Logging/Real-Time Monitoring all
RESOLVE; Kubernetes/Prometheus/Grafana/Helm/GitOps/Istio/Service Mesh/Observability do NOT — they
drop harmlessly, never fabricated).

**Hypothesis:** Re-label Maya's role to "DevOps Engineer" in the preset; add a dedicated `curatedDevOps`
sub-pool to curated_pools.go (resolved against real taxonomy in allow-list order); the existing
curatedCategoryForRole already routes "devops"/"platform"/"sre"/"cloud"/"infrastructure" → a software-
family category, so split that family so a DevOps role draws the DevOps sub-pool FIRST. Update her
bio/CV (profile.go's experienceTitle/fieldOfStudy + the preset annotation) to an infra/SRE→DevOps
progression. Re-seed demo-3 → her /home "latest skills" + /profile/skills + career timeline read as a
coherent platform/DevOps engineer.

**Expected lift:** qualitative (per TOK-10 — the P7 semantic harness isn't built yet). Before→after:
Maya's verified ~12 + claimed ~18 shift from generic-backend (Java/Node.js/REST) to platform/DevOps-
coherent (Docker/K8s, Terraform, IaC, CI/CD, Ansible, Incident Response, Cloud, Monitoring); her CV
reads as an infra/DevOps career; closure stays GREEN.

**Phase plan:** Phase B/C/D of the coverage protocol adapted to TOK-10's targeted-probe measurement
(no full sweep — the per-tik metric is the authenticated DB probe of Maya's role + skills + CV on
demo-3, re-seeded).

**Escalation conditions:** if the DevOps role doesn't resolve in the public taxonomy (it does — 10
role-core confirmed), or the curated DevOps names don't resolve enough to fill ~12 verified + ~18
claimed coherently → fall back to widening the list (still real-taxonomy-resolved). No platform edits.

**Acceptable close-no-lift:** n/a — the role + skills are confirmed resolvable; this lands.
