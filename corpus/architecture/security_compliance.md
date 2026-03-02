# Security & Compliance

This document describes the security architecture, data protection measures, and compliance posture of the Anthropos platform.

## High-Level Summary (For PMs & Non-Engineers)

Anthropos follows a **defense-in-depth** approach to security. All customer data is stored and processed in **EU-West-1 (Ireland)** by default. AI providers are routed through EU endpoints first. The platform is **GDPR-compliant** with a Data Processing Agreement (DPA v1.4) and 18 approved sub-processors. AI Simulations are classified as **Limited Risk** under the EU AI Act because scoring is deterministic (rubric-based), not AI-generated.

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
- Every table has an `organization_id` foreign key
- Ent ORM policies auto-filter all queries by organization
- No cross-tenant data access is possible at the query level

### Layer 2: Authorization
- **Sentinel** service validates every API request using Casbin (RBAC/ABAC)
- Authorization checks happen before any data access
- Policies are centrally managed and auditable

### Layer 3: Identity
- **Clerk** JWT tokens include organization context
- Sessions are org-scoped — users can only access their active organization
- Organization switching requires re-authentication

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
| **AI Token Tracking** | Centralized usage, latency, and cost tracking via shared `ai` library |

- Structured logging uses Go `slog` + Sentry integration
- ECS auto-scales on CPU/memory metrics

---

## Compliance

### EU Data Residency
- **Primary region**: EU-West-1 (Ireland)
- AI providers are routed through EU endpoints first (Azure OpenAI EU, AWS Bedrock EU, Mistral EU)
- US providers (OpenAI Direct, Anthropic Direct) used only as fallback
- No customer data stored in US by default

### EU AI Act
- AI Simulations classified as **Limited Risk** (not High Risk)
- Reason: AI is used for conversation/generation only; **scoring is deterministic** (rubric-based, 0-100 scale), NOT AI-scored
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
