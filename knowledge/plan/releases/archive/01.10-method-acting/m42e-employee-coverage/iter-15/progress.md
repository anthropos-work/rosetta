# iter-15 progress

**Type:** tik (production-fix). Active strategy: **TOK-10**. T1 — RE-SPECIALIZE Maya → DevOps Engineer
(user-confirmed Platform/DevOps, design-plan §USER DECISIONS #1).

## Phase B — role decision (real public taxonomy)
- Queried skiller for Platform/DevOps/SRE/Cloud roles: ALL public job_roles cap at exactly 10 role-core
  skills (the taxonomy is 10-core-per-role) → specialization = COHERENCE + TITLE, not depth; the curated
  sub-pool drives breadth (the same finding curated_pools.go already documents). D1.
- CHOSE role = **"DevOps Engineer"** — its 10 role-core resolve as the cleanest platform/DevOps set:
  Containerization (Docker, Kubernetes), CI/CD Pipelines, Infrastructure-as-Code (IaC), Automation Tools,
  Scripting, Version Control/Git, Security Best Practices, Performance Tuning, Troubleshooting and Debugging,
  Agile Methodologies.
- Resolved a ~44-name DevOps curated allow-list against the LIVE public taxonomy: Ansible / Terraform /
  Jenkins / IaC / Container Orchestration / Incident Response / Cloud Architecture Design / Cloud Cost
  Management / Cloud Security/Compliance/Governance / Secrets Management / Monitoring and Logging / Real-Time
  Monitoring and Alerts / Networking Fundamentals / root-cause analysis ALL RESOLVE; Kubernetes / Prometheus /
  Grafana / Helm / GitOps / Istio / Service Mesh / Observability / Chaos Engineering / Linux Administration do
  NOT (kept aspirationally, drop harmlessly — no fabrication). D2.

## Phase C — fix (rext, zero platform edit)
- `curated_pools.go`: NEW `curatedDevOps` category + `curatedDevOpsSkills` allow-list (verified block +
  aspirational tail); wired into `curatedSkillsFor`; `curatedCategoryForRole` now routes platform/DevOps/SRE/
  cloud/infra roles → `curatedDevOps` BEFORE the generic software family (so a role that is both engineering
  AND infra leans DevOps). D2/D3.
- `profile.go`: `experienceTitle` now role-aware — a DevOps role gets the infra ladder (Junior Cloud Engineer
  → Site Reliability Engineer → DevOps Engineer); NEW `isDevOpsRoleTitle` helper. The bio summary
  (roleSummaryFor) auto-reads "devops engineer…" off the resolved role; the degree field stays Computer
  Science (engineer family). D4.
- `presets/stories.seed.yaml`: Maya's role `Backend Software Engineer` → `DevOps Engineer`; annotation +
  comments re-pointed to platform/DevOps. D5.
- Tests: updated `TestCuratedCategoryForRole` (SRE + Cloud Platform Engineer now correctly → curatedDevOps;
  added DevOps/Platform/Cloud/Infra role assertions); NEW `TestResolveCuratedPools_DevOps` (allow-list order
  + drop + category isolation). `go test ./...` GREEN; build clean.

## Phase D — re-measure (demo-3 reset + clean re-seed)
- A re-specialization needs a `--reset` (user_skills is COPY-idempotent-on-id → the old backend skill_id
  survives an additive re-seed at the same slot id). Reset + clean re-seed: 51734 rows, isolation clean
  (prod=false). [NOTE: the reset cleared the casbin g2/g3 grants — a Sentinel reload is needed for the sim
  entitlement; P5 fixes that for fresh demo-up; here it sets up the P5 proof.]
- Maya AFTER: role **DevOps Engineer**; VERIFIED 12 — Agile / Automation Tools / Containerization (Docker,
  Kubernetes) / Container Orchestration with Kubernetes / CI/CD Pipelines / IaC / Performance Tuning /
  Scripting / Security Best Practices / Terraform / Troubleshooting and Debugging / Version Control/Git
  (ALL platform/DevOps-core, ZERO junk, ZERO generic-backend leakage). CLAIMED 18 — AWS / Ansible / Cloud
  Architecture Design / Cloud Cost Management / Cloud Platform Expertise / Cloud Security/Compliance/Governance
  / Git / GCP / Incident Response / Jenkins / Azure / Monitoring and Logging / Networking Fundamentals /
  Network Security / root-cause analysis / Real-Time Monitoring / Scripting for Automation / Secrets Management
  (ALL DevOps-coherent).
- CV arc: **Junior Cloud Engineer (2019) → Site Reliability Engineer (2021) → DevOps Engineer (2024, current)**
  — a believable infra→DevOps progression. EDU: B.Sc. Computer Science, KTH. BIO: "Maya is a devops engineer
  focused on shipping reliable, well-scoped work."
- `datadna measure-closure --stack demo-3`: **PASS** (every seeded verified-skill node-id resolves — no
  fabrication).

## Close — 2026-06-25
**Outcome:** T1 landed. Maya re-specialized Backend Software Engineer → **DevOps Engineer**: 12 verified + 18
claimed, ALL platform/DevOps-coherent (Docker/K8s, Terraform, IaC, CI/CD, Ansible, Incident Response, Cloud,
Monitoring, Secrets), ZERO junk; CV reads as an infra→DevOps career; closure PASS.
**Type:** tik (production-fix)
**Status:** closed-fixed
**Gate:** NOT MET (T1 of the run-5 T1→P4→P5 scope; the believability gate needs P4–P8 + the P7 semantic harness)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1st tik) — (6) protocol-stop: n — Outcome: continue (→ iter-16 P4)
**Decisions:** D1 (10-core-per-role → coherence not depth), D2 (DevOps role + curated sub-pool), D3
(category routing order), D4 (DevOps CV arc), D5 (preset role) — see ./decisions.md.
**Routes carried forward:** P4 (avatar real photos + org logo), P5 (FATAL Sentinel-reload reproducibility) —
this run; P6/P7/P8 — later runs.
**Lessons:** a role RE-specialization (not just a first specialization) requires a `--reset` before re-seed —
user_skills is COPY-idempotent on a deterministic slot id, so an additive re-seed keeps the OLD skill_id at
that slot. A fresh demo-up resets implicitly, so this only bites a live re-seed (measurement). Folded into the
coverage-protocol re-apply note.
**rext:** commit `bba7f67`, tag `method-acting-m42e-iter15`.
