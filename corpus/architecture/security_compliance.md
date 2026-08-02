# Security & Compliance

This document describes the security architecture, data protection measures, and compliance posture of the Anthropos platform.

## High-Level Summary (For PMs & Non-Engineers)

Anthropos follows a **defense-in-depth** approach to security. All customer data is stored and processed in **EU-West-1 (Ireland)** by default. AI providers are routed through EU endpoints first. The platform is **GDPR-compliant** with a Data Processing Agreement (DPA v1.4) and 18 approved sub-processors. AI Simulations are classified as **Limited Risk** under the EU AI Act — **but the stated reason for that classification does not hold at platform HEAD** (see [EU AI Act](#eu-ai-act) below): the rubric *arithmetic* is deterministic, the per-check pass/fail verdicts it counts are produced by an LLM.

Key guarantees:
- EU data residency (primary)
- Multi-tenant data isolation at database, authorization, and identity layers
- 90-day auto-deletion of personal data post-contract
- Full DB backups every 6 hours to three geographically separate locations
- No direct SSH to production; all access via Tailscale VPN

---

## Network Security

### VPC Architecture
- **VPC CIDR**: 10.0.0.0/16 with Multi-AZ deployment
- **Public subnets**: Application Load Balancer (ALB), Cosmo Router
- **Private subnets**: All microservices (no direct internet access)
- **Data subnets**: PostgreSQL RDS, Redis ElastiCache
- **Controls**: Network ACLs, Security Groups, least-privilege rules

### Developer Access
- **Tailscale VPN** for dev team and GitHub Actions self-hosted runners
- VPN provides secure access to private subnets
- No direct SSH to production instances
- Mandatory MFA for admin access

---

## Transport & Data Encryption

### In Transit
- TLS 1.3 (ECDHE key exchange, RSA, AES-128+, SHA256+)
- All service-to-service communication encrypted

### At Rest
- AES-256 on RDS, EBS, and S3 (AWS KMS managed keys)
- Encryption enabled by default on all storage

---

## Access Management

| Mechanism | Purpose |
|:----------|:--------|
| **AWS Secrets Manager** | DB credentials, API keys |
| **AWS Parameter Store** | Encrypted configuration |
| **IAM Roles** | Role-based, least-privilege access |
| **CloudTrail** | Audit trail for all AWS API calls |
| **GitHub 2FA** | Mandatory for all org members |
| **Branch Protection** | Main branch requires code review |

---

## Multi-Tenant Data Isolation

Three layers of isolation ensure tenant data cannot leak:

### Layer 1: Database

> **⚠️ Isolation is NOT automatic across the whole schema — do not rely on it as a blanket guarantee.**
> Measured at `app` HEAD: of **135** Ent schemas (139 `.go` files, 4 of which declare no schema), only
> **30** use `OrganizationMixin{}` — the one that carries the privacy `Policy()` (`mixin.go:126`). Seven use
> `OrganizationIDMixin{}`, explicitly *"a plain nullable organization_id column"* with **no policy** — **and
> a further 18 declare a plain `organization_id` field with neither mixin**. Two of those 18 are policed by
> other means: `org_membership.go` declares its own fail-closed org `Policy()` (`:172-188`, ending in
> `privacy.AlwaysDenyRule()`), and `academy_feedback.go` carries `UserMixin{}`, whose `Policy()`
> (`mixin.go:98`) applies a row-level **owner** filter (`rule.FilterOwnerRule()`) — scoped by *user*, not by
> organization.
>
> **So: 31 schemas auto-filter by ORGANIZATION** (the 30 mixin users + `Membership`), and **16 carry an
> `organization_id` with no policy of any kind**: `org_subscription.go`, `organization_settings.go`,
> `organization_feature.go`, `api_key.go`, `lab_session.go`, `interview_aggregated_report.go`,
> `admin_audit_log.go`, `job_simulation_session.go`, `jobsimulation_feedback.go`,
> `ai_readiness_diagnose_narrative.go`, `ai_readiness_recommendation.go`, `assignment_invitation_link.go`,
> `job_role_skill_suggestion_cache.go`, `org_membership_invitation.go`, `org_sim_link.go`,
> `profile_history.go`. **Those 16 are the rows most likely to be missed by an audit**: they look org-scoped
> and are not policed. The remainder (the taxonomy, and other global reference data) carry no org column by
> design.
>
> **⚠️ This fence has now been wrong FOUR times. Re-derive it; do not quote it.** v1 asserted a blanket
> guarantee. M257x iter-33 over-swung to "the non-mixin schemas never mention organization at all". iter-34
> found `org_membership.go` listed as unpoliced when it polices itself. iter-34's *own* correction then
> listed `academy_feedback.go` as unpoliced when `UserMixin` filters it. The first three failed toward
> *"isolation is handled"*; the fourth failed toward alarm — **the error direction is not stable, so neither
> reading is safe without measuring.**
>
> Derivation (restrict to schema files — a bare `*.go` glob pulls in `skiller_mixins.go` and returns 19):
>
> ```sh
> cd app/internal/data/ent/schema
> ALL=$(grep -l 'ent.Schema' *.go | sort)
> comm -23 <(echo "$ALL") <(grep -lE 'Organization(ID)?Mixin\{\}' *.go | sort) \
>   | xargs grep -l '"organization_id"'          # -> 18
> # then subtract any that declare their own Policy() or carry UserMixin{} -> 16
> ```
> The platform states this itself: `job_simulation_session.go:5` — *"L2: NO Ent privacy Policy;
> owner/org/tenant are plain fields"* — and `jobrole.go:18` / `category.go:15` note the taxonomy is
> deliberately globally readable. **Scoping on the jobsim fan-out and the taxonomy is the caller's job.**

- Org-scoped tables carry an `organization_id` column
- Ent privacy policies auto-filter by organization on **31** schemas — the 30 using `OrganizationMixin{}` plus `Membership`, which declares its own
- Cross-tenant reads are prevented at the query level **on those tables**; elsewhere isolation is
  enforced by Layer 2 (Sentinel) and by explicit query scoping, not by the ORM

### Layer 2: Authorization
- **Sentinel** service validates every API request using Casbin (RBAC/ABAC)
- Authorization checks happen before any data access
- Policies are centrally managed and auditable

### Layer 3: Identity
- **Clerk** JWT tokens include organization context
- Sessions are org-scoped — users can only access their active organization
- Organization switching **re-mints the session JWT with the new org claim — it is NOT a re-authentication.** It is a client-side `clerk.setActive({ organization })` call (`next-web-app/apps/{web,hiring}/src/hooks/useOrgSelection.tsx:94`, `useResolveActiveOrg.tsx:107`, `useActivateMembershipOrg.tsx:81`); no credential is re-presented and no sign-out occurs. The isolation that follows comes from the new claim, not from a fresh proof of identity

---

## Backup & Disaster Recovery

| Aspect | Detail |
|:-------|:-------|
| **Full backups** | Every 6 hours → S3, Azure, Hetzner (Germany) |
| **Point-in-time recovery** | RDS automated backups |
| **Primary region** | EU-West-1 (Ireland) |
| **DR site** | US AWS region |
| **Deployment** | Multi-AZ with auto-scaling |
| **CDN** | Worldwide (Vercel for frontend) |

The `db-backup` service runs on a schedule, dumping PostgreSQL to three geographically separate locations for resilience.

---

## Server & Runtime Security

- All infrastructure provisioned via **Terraform** (Infrastructure as Code)
- Containers rebuilt from fresh base images regularly
- Monthly patch updates; critical patches can be accelerated
- Git tags trigger automated deployments
- Critical services require manual deployment approval
- ECS health checks every 30 seconds with automated rollback on failure

---

## Monitoring, Logging & Incidents

| Tool | Purpose |
|:-----|:--------|
| **CloudWatch** | Metrics, dashboards, alarms; structured logs with 90-day retention |
| **Sentry** | Error tracking, performance monitoring, cron job monitoring |
| **PostHog** | Product analytics |
| **Better Stack** | Incident escalation, uptime monitoring |
| **AI Token Tracking** | Centralized usage, latency, and cost tracking in **`app/internal/aiusage`** — **not** the shared `ai` library, which only wraps providers (consistent with `README.md:21` + `ai_architecture.md`) |

- Structured logging uses Go `slog` + Sentry integration
- ECS auto-scales on CPU/memory metrics

---

## Compliance

### EU Data Residency
- **Primary region**: EU-West-1 (Ireland)
- AI providers are routed through EU endpoints **first** — Azure OpenAI EU (`ai.go:262-266`), AWS Bedrock
  pinned to `eu-west-1` (`:85-88`)
- **⚠️ "EU-first" is not "EU-only", and the US path is a FLAG, not a fallback.** `getClient` swaps
  `azureClientEu` → **`azureClientUs`** whenever the PostHog flag **`flag_use_azure_us`** is enabled
  (`app/internal/jobsimulation/ai/ai.go:263-277`). That is a deliberate switch that can route live
  simulation traffic to a US region with no error condition involved. Direct OpenAI is additionally used as
  the **retry target on HTTP 429** (`isThrottlingError`, `:129` / `:166` / `:325`)
- **"Anthropic Direct" is not used at all** — Anthropic is reached exclusively through **AWS Bedrock
  `eu-west-1`** (`:85-95`)
- No customer data stored in US **by default** — but the residency guarantee is contingent on
  `flag_use_azure_us` being off, and that is a runtime flag, not a build-time property. **Check the flag
  before asserting residency.**

### EU AI Act
- AI Simulations classified as **Limited Risk** (not High Risk)
- Stated reason: AI is used for conversation/generation only; scoring is deterministic (rubric-based, 0-100 scale), NOT AI-scored

> **⚠️ THE STATED REASON IS FALSE AT PLATFORM HEAD, AND THIS IS A COMPLIANCE CLAIM.** It is a conjunction
> and **both conjuncts fail**. The *aggregation* is deterministic arithmetic — `calculateSkillScore`
> (`app/internal/jobsimulation/simulator/validation/v3/validator/skills.go:53-64`) counts booleans and
> `:75` divides. **The booleans it counts are LLM output.** The validator registers exactly ONE check
> engine — but **cite the DISPATCH, not that map**: `checkerEngines` is stored and never read, so it is
> not the mechanism. The real path is the hardcoded switch at
> `internal/jobsimulation/simulator/validation/basevalidator/criterion.go:127` → `validateLLM` →
> `NewLLMBulkChecker(c.logger)` (`:428`), which sends
> `basevalidator/templates/checkValidationBulk.tmpl` — a prompt asking a model to *"assess whether the
> `<asset>` … meets or does not meet"* each check and to return `{"check_id", "feedback", "success"}`. So
> "AI is used for conversation/generation only" is also false.
>
> **Not ALL verdicts are LLM-produced, and the honest claim is "most":** `EngineTextDiff` checks run
> deterministically alongside them (`criterion.go:168` dispatches `validateCodeDiff`; `:450-475` sets
> `success` from a pure string comparison, no model), and both result sets are appended together.
>
> **What follows is a question for counsel, not for this corpus**: a system that judges workers and
> candidates sits near Annex III. **Do not cite this section as evidence of a Limited-Risk
> classification** — re-derive it. Measured M257x iter-38; the same false premise was stated
> independently in `ai_architecture.md` and is corrected there too.
- This classification means transparency obligations only, not the strict requirements of High Risk systems

### GDPR / CCPA
- **90-day auto-deletion** of personal data post-contract termination
- CV data is never used for AI training
- Data Processing Agreement (DPA v1.4) with 18 approved sub-processors
- Data subject access/deletion requests supported

### Sub-Processors

Key sub-processors include:
- **AI**: OpenAI, Anthropic, Mistral
- **Voice/Recording**: LiveKit, AWS Chime
- **Infrastructure**: AWS, Vercel
- **Auth**: Clerk
- **Analytics**: PostHog, Sentry
- **Email**: Brevo (Sendinblue)

---

## Related Documentation
- [Architecture Overview](./architecture_overview.md)
- [AI Architecture](./ai_architecture.md)
- [Service Taxonomy](./service_taxonomy.md)
- [Rosetta Tooling Safety Contract](../ops/safety.md) — how the demo/dev stack **tooling** stays safe (never reads
  customer data, never touches prod); the layer *above* the platform's own tenant-isolation posture described here.
