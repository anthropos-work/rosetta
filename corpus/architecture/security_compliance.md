# Security & Compliance

This document describes the security architecture, data protection measures, and compliance posture of the Anthropos platform.

## High-Level Summary (For PMs & Non-Engineers)

Anthropos follows a **defense-in-depth** approach to security. All customer data is stored and processed in **EU-West-1 (Ireland)** by default. AI providers default to EU endpoints — but there is **no ordered EU-first fallback chain** (`external_services.md:579`): the US path is a PostHog **flag** (`flag_use_azure_us`), not a fallback rung, and Course Builder's `ANTHROPIC_API_KEY` path leaves the EU entirely. See [EU Data Residency](#eu-data-residency). The platform is **GDPR-compliant** with a Data Processing Agreement (DPA v1.4) and 18 approved sub-processors. AI Simulations are classified as **Limited Risk** under the EU AI Act — **but the stated reason for that classification does not hold at platform HEAD** (see [EU AI Act](#eu-ai-act) below): the rubric *arithmetic* is deterministic, but **most** of the per-check pass/fail verdicts it counts are produced by an LLM. *Most*, not all — deterministic `EngineTextDiff` checks are the exception, and "all verdicts are AI" is the opposite error. **The legal classification itself is a question for counsel; this corpus only records that the stated technical premise is false.**

Key guarantees:
- EU data residency (primary)
- Multi-tenant data isolation at database, authorization, and identity layers
- 90-day auto-deletion of personal data post-contract
- Durability is **RDS Multi-AZ + an hourly AWS Backup plan with PITR**. The separate offsite `db-backup` job is **deployed but untriggered** since `7dd1b80` (2025-05-29) — **this line read *"full DB backups every 6 hours to three geographically separate locations"* until M257x iter-124**, and every element of it was wrong: two destinations not three, never Azure, no cadence at all (the disabled value was `rate(12 hours)`; *"6 h"* never had a source), and nothing fires it. See [`db-backup.md`](../services/db-backup.md)
- No direct SSH to production; all access via Tailscale VPN

---

## Network Security

### VPC Architecture
- **VPC CIDR**: 10.0.0.0/16 with Multi-AZ deployment
- **Public subnets**: Application Load Balancer (ALB). ⚠️ **The Cosmo Router was listed here until M257x
  iter-115 and the only readable evidence contradicts it.** Re-derived across **all eight** service terraform
  trees **then in the clone set** (`app`, `sentinel`, `graphql-wundergraph`, `messenger`, `cms`, `roadrunner`,
  `storage`, `jobsimulation` — ⚠️ **`sentinel` left the clone set at platform `766df6c`**, v11.0, which folded
  it into `app` and deleted its `repos.yml` entry; the tree was read when it was still cloned and the reading
  is unaffected, but seven of the eight are now frozen repos rather than a live clone set):
  the token `public_subnet` occurs **0 times**, and **every one of the eight**
  passes `private_subnets_ids = var.platform_private_subnets_ids` — the router at
  `graphql-wundergraph@60c229f3:terraform/main.tf:31`, with **no public-subnet argument of any kind**. The
  router uses the same `base_service` module as `app` (`:11`) and `app` passes the same private ids, so these
  two bullets singled the router out for a placement it shares with `backend`, which the next bullet files as
  private. **Residual, stated rather than hidden:** `infrastructure` (which *defines* `base_service`) has never
  been in a clone set and `use_fargate = false` (`:13`) puts tasks on cluster instances this corpus cannot see —
  so the module could in principle place them elsewhere. What is measurable says private; what was published
  said public, with no ref. **Also note the router is gone from local dev entirely** (platform `2adcf71`) and
  the repo is archived on GitHub
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
> Measured at `app` `ad9f3c49` (== `origin/main`, 2026-08-06): of **135** Ent schemas (139 `.go` files, 4 of which declare no schema), **30** *mention*
> `OrganizationMixin{}` — the mixin that carries the privacy `Policy()` (`mixin.go:126`) — but only **29** *use* it: `user_resource.go:22` reads `// OrganizationMixin{},  // We need to work on this`, and a commented line compiles into nothing. **The predicate is *use*, not *mention*.** Seven use
> `OrganizationIDMixin{}`, explicitly *"a plain nullable organization_id column"* with **no policy** — **and
> a further 18 declare a plain `organization_id` field with neither mixin**. Two of those 18 are policed by
> other means: `org_membership.go` declares its own fail-closed org `Policy()` (`:172-188`, ending in
> `privacy.AlwaysDenyRule()`), and `academy_feedback.go` carries `UserMixin{}`, whose `Policy()`
> (`mixin.go:98`) applies a row-level **owner** filter (`rule.FilterOwnerRule()`) — scoped by *user*, not by
> organization.
>
> **So: 31 schemas auto-filter by ORGANIZATION** — **29** `OrganizationMixin{}` users, plus **two** that
> declare their own: `Membership` (`org_membership.go:172`, `rule.AllowCurrentOrgEdgesOrSkipRule()`) and
> **`Organization` itself** (`organization.go:56`, `rule.FilterSameOrganizations()` at `:96`, and it uses
> neither mixin). `User` also declares its own `Policy()` (`user.go:116`) but filters by **user**, not
> organization, so it is correctly excluded.
>
> ⚠️ **The total is 31, but the old derivation of it was wrong — and so was the audit that "corrected" it.**
> `grep -c 'OrganizationMixin{}' schema/*.go` returns **30**; one of them,
> `user_resource.go:22`, is **commented out** (`// OrganizationMixin{},  // We need to work on this`), so the
> live count is **29**. The long-standing *"30 mixin users + `Membership` = 31"* was therefore right by two
> compensating errors: it over-counted the mixin set by one and omitted `Organization`. M257x iter-49's audit
> booked the total as **32**; iter-52 repaired the corpus to 32 and its two pre-commit readers **independently
> refuted it** — 29 + 2 = **31**. The number is restored; the derivation is now the correct one.
> **Re-derive the SET, not the sum, and exclude commented lines when you do** — a `grep -c` over Go source
> counts code that does not compile into anything. `user_resources` is user-scoped, NOT org-scoped by the ORM.
> And **23 carry an
> `organization_id` with no policy of any kind.** Sixteen of the 23 have neither mixin — `org_subscription.go`, `organization_settings.go`,
> `organization_feature.go`, `api_key.go`, `lab_session.go`, `interview_aggregated_report.go`,
> `admin_audit_log.go`, `job_simulation_session.go`, `jobsimulation_feedback.go`,
> `ai_readiness_diagnose_narrative.go`, `ai_readiness_recommendation.go`, `assignment_invitation_link.go`,
> `job_role_skill_suggestion_cache.go`, `org_membership_invitation.go`, `org_sim_link.go`,
> `profile_history.go` — **and those 16 are the rows most likely to be missed by an audit**: they look
> org-scoped and are not policed. **The remaining 7 of the 23** are the **7
> `OrganizationIDMixin{}`** users named above (`category`, `jobrole`, `similarity`, `skill`,
> `specialization`, `studio_document`, `studio_task`), which carry an `organization_id` and declare no
> `Policy()` either. An earlier revision of this paragraph closed with *"the remainder … carry no org column
> by design"*, which excluded those 7 from its own count **three lines after naming them as unpoliced** — a
> contradiction inside a single blockquote, and in the direction that reads as *"isolation is handled"*.
> **The remainder — 135 − 31 policed − 23 unpoliced-with-`organization_id` = 81 schemas — is NOT uniformly "global reference data."** Most of it carries no org column by design, but at least **four** members are per-TENANT and unfiltered on that axis: `lab.go:63`, `academy_chapter.go:84` and `academy_skill_path.go:78` each declare `field.String("tenant_eid")`, and `skill_path_session.go:43` declares `field.UUID("tenant_id", …)` (all @ `app` `ad9f3c49`). None of the four carries an `organization_id` or either org mixin; the first three declare **no `Policy()` at all**, and the fourth is filtered by `UserMixin{}` — by *user* (owner, `mixin.go:98`), never by its `tenant_id`.
> A fifth, `academy_feedback.go`, falls in that remainder by the same arithmetic and **does** carry an `organization_id` (`:129`); it is out of the 23 only because `UserMixin{}` (`:64`) polices it by owner. **Scoping the tenant-scoped four is the caller's job**, exactly as it is for the jobsim fan-out and the taxonomy named below. The sentence replaced here — a **later** slip than the one named just above, and in the same *"isolation is handled"* direction — read *"The genuine remainder — global reference data with no `organization_id` at all — is what carries no org column by design"*, which silently re-classed those four tenant-scoped schemas as reference data. Booked M257x iter-99, repaired iter-102.
>
> **Re-measured M257x iter-46, at `app` @ `5ba17044`:** 139 `.go` files in `internal/data/ent/schema/`;
> `OrganizationMixin{}` in 30, `OrganizationIDMixin{}` in 7, a plain `organization_id` with neither mixin in
> 18 (a 19th hit, `skiller_mixins.go`, is a mixin definition and not a schema); and **only FOUR files in the
> whole directory declare any `Policy()` at all** — `organization.go`, `mixin.go`, `user.go`,
> `org_membership.go`. Two auditors disagreed on the base count (17+7 vs 16+7); the measurement says **16**,
> because `org_membership.go` polices itself and `academy_feedback.go` is filtered by owner.
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
- Ent privacy policies auto-filter by organization on **31** schemas — the **29** live `OrganizationMixin{}` users (a 30th is commented out at `user_resource.go:22`), plus `Membership` and `Organization`, which each declare their own
- Cross-tenant reads are prevented at the query level **on those tables**; elsewhere isolation is
  enforced by Layer 2 (Sentinel) and by explicit query scoping, not by the ORM

### Layer 2: Authorization

> ⚠️ **CORRECTED M257x iter-120 — this layer said Sentinel *"validates **every** API request"* and that
> *"authorization checks happen before **any** data access."* Both are false, and they are false in the
> direction that makes the platform sound MORE protected than it is.** The same class as the
> [`clerk-integration.md`](../services/clerk-integration.md)'s **Sign-in tokens** bullet — the one whose
> own text now reads *"this bullet used to say **"only"**, and it was false"* — and the `cms.md` `bash -c`
> inversion: an absolute quantifier over a
> security surface, published unhedged. Layer 1 directly above has been re-measured four times down to
> an exact schema count and carries its caveat; **Layer 2 sat underneath it, unhedged and unfenced,
> asserting a blanket that does not exist.**
>
> **The platform's own source says so, in a comment written as a post-mortem of this exact misreading**
> (`app/internal/web/backend/graphql/graph/resolver_skiller_taxonomy_authz.go:53-66` @ `app` `ad9f3c49`):
> the M207 skiller-in-app port dropped skiller's per-resolver guards and *"leaned on app's blanket
> `AuthorizationMiddleware` — but that gate is keyed on a `userId` operation variable and **FAILS OPEN**
> for taxonomy operations … That left every taxonomy read/write reachable by any authenticated caller
> (**cross-tenant IDOR + privilege escalation**)."* It ends: ***"Do NOT rely on the blanket gate for this
> surface — it fails open here."*** (That specific hole was closed by restoring per-resolver checks; the
> **general** statement about the gate is what this doc got wrong.)

**What is actually enforced.** `AuthorizationMiddleware`'s own doc comment says it *"gates every
operation **on a viewer**"* (`app/internal/authorization/gqlauthz/gqlauthz.go:149` @ `app` `ad9f3c49`) —
an **authentication** gate. The single Sentinel call is `OrgCheckUserPermission` at `:222`, and **six
paths reach the resolver before it**:

| `gqlauthz.go` | condition | effect |
|---|---|---|
| `:160-161` | the operation failed to parse/validate | `next(ctx)` |
| `:174-178` | no viewer **and** the op is `@public` / federation / dev introspection | `next(ctx)` |
| `:190-191` | **viewer has no active org** (`org == nil \|\| org.ID() == uuid.Nil`) | `next(ctx)` |
| `:196-197` | `errUnknownTarget` — the operation carries no `userId` **variable** | `next(ctx)` |
| `:202-203` | target is nil, or the target **is** the viewer | `next(ctx)` |
| `:209-219` | `@resolverAuthorized` — grants `authorization.Allow` **without calling Sentinel** | `next(ctx)` |

So an operation is Sentinel-checked only when it (i) carries a `userId` **GraphQL variable** — the target
is one hardcoded variable name, `grapqlTargetVar = "userId"` (`gqlauthz/target.go:11`) — (ii) whose value
differs from the viewer's own id, and (iii) the viewer has a non-nil active org. **An id inlined as a
document literal rather than passed as a variable does not reach Sentinel at all.**

**And there is no BLANKET authz middleware on the REST surface.** `git grep -nE
'AuthzMiddleware|EchoAuthzMiddleware' -- '*.go'` over `app` @ `ad9f3c49` returns **0**, and REST
authorization is opt-in **per group or per handler**, never applied to the surface as a whole:

| Echo group | `backend.go` | middleware stack |
|---|---|---|
| `/api` | `:121-141` | `cors` + `swagger` + `authn` (10 skipped paths: 8 webhooks, `/api/skills`, `/api/health`) |
| `/ask` | `:171-173` | `cors` + `authn` — the `admin/auto-rules/*` routes are gated **per handler** by `requireAdmin`, which the source says in place (`:183-185`) |
| `/assignment-builder` | `:194-196` | `cors` + `authn` + a 2 MB body limit |
| `/admin/backfill` | `:210-212` | `cors` + `authn` — *"gated on the admin Sentinel permission"* per handler (`:207-209`) |
| `/coursebuilder` | `:229-232` | `cors` + `authn` + **`cbGate`** |
| `/credits` | `:273-276` | `cors` + `authn` + **`cbGate`** |

> ⚠️ **CORRECTED AGAIN at run 81 — and this time the DENOMINATOR was wrong, not the quantifier.** The
> table above enumerates only the groups declared **inside `backend.go`**. `app` mounts **ELEVEN**
> non-test Echo groups on the **one** REST instance (`internal/web/web.go:124-163`), and the five below
> were invisible to every previous reading of this paragraph. **Three never touch the Clerk `authn`
> middleware at all, and one has no authentication whatsoever.** All at `app` `ad9f3c498`:

| Echo group | declared | middleware stack |
|---|---|---|
| `/api/invitations` | `internal/invitations/handlers.go:31` (mounted `web.go:148`) | **`cors` is the only middleware — and that is NOT the same as no authentication.** ⚠️ **Corrected at run 82 — see the box below this table.** The credential is the path segment: a 256-bit `base64url(HMAC-SHA256(email\|org_id\|invited_at, INVITATION_HMAC_SECRET))` (`internal/invitations/token.go:29-34`), **checked before any data is returned** — `invite.go:159` / `:194` filter on the stored `token` column and a miss returns `404 not_found` |
| `/content/admin` | `internal/web/backend/content_admin.go:35` (mounted `backend.go:294`) | **no Clerk `authn`** — a bearer shared secret (`ACADEMY_CONTENT_API_TOKEN`) is the entire gate |
| `/v1/labs` | `internal/web/backend/labs_admin.go:31` (mounted `backend.go:306`) | **no Clerk `authn`** — a group-level org **API key + `labs:write` scope** check |
| `/academy/embeddings` | `internal/web/backend/academy_embeddings_admin.go:41` (mounted `backend.go:300`) | `cors` + `authn` |
| `/api/workforce` | `internal/web/backend/emailpreview/handler.go:66` (mounted `web.go:162`) | `cors` + `authn` — **grouped off the ROOT `e`**, so despite the `/api/` prefix it does **not** inherit the `/api` group's swagger/authn stack |

> ⚠️ **CORRECTED AGAIN at run 82 — this time the alarm was too loud, not too quiet.** Run 81's row for
> `/api/invitations` read *"`cors` ONLY — no authentication"*, and that overstates it. **Settled at `app`
> `ad9f3c498` by reading the `RegisterRoutes` call site and the manager, not the mount comment:**
>
> - **What is absent is Clerk, and deliberately.** `RegisterRoutes(srv.e, cors.EchoCORSMiddleware(...))`
>   (`internal/web/web.go:148`) passes exactly one middleware, because next-web-app renders the
>   invite-landing and unsubscribe pages **before the caller has an account** and calls these
>   cross-origin (`web.go:145-146`). The group mounts `GET /:token` and `POST /:token/opt-out`
>   (`internal/invitations/handlers.go:31-33`).
> - **A credential IS required, and it is checked before anything is disclosed.** The token is minted at
>   invite time as `base64url(HMAC-SHA256(email|org_id|invited_at, INVITATION_HMAC_SECRET))`
>   (`internal/invitations/token.go:29-34`; `main.go:423-427` refuses to boot without the secret) and
>   stored on the row. `GetInviteDetailsByToken` and `OptOutByToken` filter on it —
>   `Where(membershipinvitation.Token(token))` at `internal/invitations/invite.go:159` and `:194` — so a
>   non-matching token yields `404 not_found` / `already_opted_out` and **no row, no email, no org name**.
> - **The mechanism is a stored bearer capability, not a re-verified signature.** `TokenManager.ValidateToken`
>   (`token.go:38`, constant-time) is called by **nothing outside its own test** — measured repo-wide,
>   `git grep -n ValidateToken ad9f3c498 -- '*.go'` returns 8 hits, all in `token.go` + `token_test.go`.
>   The HMAC supplies 256 bits of unguessability and determinism; the *check* is an equality match on an
>   indexed column.
> - **Both handlers bypass the Ent privacy layer on purpose** — `privacy.DecisionContext(ctx, privacy.Allow)`
>   at `invite.go:157` and `:190` — and the source states the model in its own words:
>   *"It backs a public endpoint — **token possession is the authorization**"* (`:154-155`, `:187-188`).
>
> **So the accurate sentence is *token-authenticated, deliberately pre-login* — not *unauthenticated*.**
> No defect is filed: the design is stated, the credential is required, and the disclosure it gates
> (invited email, org name, inviter name) is the content of the invitation the token was mailed with.
> *(Two observations that are not defects and are recorded so nobody re-derives them: the token is
> written to application logs on both the miss path and the opt-out success path — `handlers.go:59`,
> `:62`, `:99` and `invite.go:203`, `:218`; and `/api/invitations` is the second of two public token endpoints, the other
> being the root-mounted one in the next box, which **does** verify the signature.)*

> ⚠️ **AND THE ENUMERATION ITSELF IS NARROWER THAN THE SURFACE.** *"Eleven groups"* is correct **for
> groups** — re-derived independently at run 82, `git grep -nE '\.Group\("' ad9f3c498 -- '*.go' | grep -v
> _test.go` returns exactly those 11. But `app` also mounts **eight routes directly on the root `e`,
> inside no group at all**, so no group-level middleware statement reaches them:
>
> | root-mounted route | declared | gate |
> |---|---|---|
> | `/graphql/query` | `backend.go:317` | **no Echo middleware** — authn/authz happen inside the GraphQL chain (the *"fail-open"* viewer gate above) |
> | `/api/webhook/directus` | `backend.go:324` | handler self-authenticates via the shared secret; mounted only when the Directus edge is configured |
> | `/ai-readiness/unsubscribe/:token` | `internal/aireadiness/notifications/handlers.go:41` (mounted `web.go:153`) | `cors` + **the HMAC signature is verified in-handler**, `401` on mismatch (`handlers.go:71-81`, verifier at `:156`) — the invitations group's sibling, and the contrast is instructive: **this one re-derives the signature, `/api/invitations` matches a stored token** |
> | `/api/schema.json` | `backend.go:117` | **none** — serves the OpenAPI document; registered on the root before the `/api` group, so it inherits none of that group's stack |
> | `/content/catalog.json` | `internal/web/backend/content.go:23` | **none, by design** — `content_admin.go:32` says so: it *"stays open by design"* |
> | `/graphql` (Apollo Sandbox) | `backend.go:320` | **`colony.Development` only** |
> | `/api/*` (API docs) | `backend.go:314` | **`colony.Development` only** |
> | **`/v1/labs/:slug/workspace.tar.gz`** | `internal/web/backend/labs_admin.go:40` (wired unconditionally at `backend.go:306`) | **OPTIONAL auth, deliberately** — it is mounted on the root `e`, *outside* the `/v1/labs` write group and therefore outside that group's `apiKeyAuthMiddleware(…, "labs:write")`. The file says so in its own words: *"Serve is OUTSIDE the write group — it has OPTIONAL auth (a public Lab's workspace is served to anyone; a tenant-private Lab requires a key with access)"* (`:36-39`). It serves a **workspace tarball**, and it is the URL the control-plane fetches at boot (`CP_WORKSPACE_BASE_URL`) |
>
> **Rule 57 applied to rule 57's own repair — twice now.** Run 81 widened the search from one file to
> the whole service and got the *group* count right; it did not widen from *groups* to *routes*. Then
> the routes row said **seven** and that was wrong too: **the eighth is the Labs workspace tarball**,
> added M257x iter-136 after three independent adjudicators converged on the count at iter-135 and one
> of them located the route. **The seat that first reported the miscount named the WRONG route** — it
> proposed `/ai-readiness/unsubscribe/:token`, which this table already contained — so a repair driven
> by that report would have changed nothing. *Grade a count claim by re-enumerating, never by accepting
> the reporter's candidate.*
>
> The honest shape of the REST surface is **11 groups + 8 ungrouped root mounts**. **Of the eight: two
> are open by design (`/api/schema.json`, `/content/catalog.json`), two are development-only, three
> self-authenticate in-handler (`/graphql/query` via the GraphQL chain, `/api/webhook/directus` via the
> shared secret, `/ai-readiness/unsubscribe/:token` via the HMAC), and one — the tarball — carries
> optional auth with the tenancy decision inside the handler.** That enumeration is stated rather than
> summarised on purpose: `D-M257x-121-2` records this milestone publishing a *new* absolute quantifier
> over a security surface inside a repair whose whole subject was absolute quantifiers.

> **`cbGate` is also not the only group-level authorization**: `/v1/labs` carries a group-level *scope*
> check and `/content/admin` a group-level shared-secret check. **This is the THIRD correction to this
> one paragraph** — iter-120 over-stated it, iter-121 corrected the quantifier, and run 81 found the
> *denominator* had been six all along because the enumeration only ever read one file. **A count is only
> as wide as the search that produced it**, and the two earlier repairs both re-derived from
> `backend.go` because that is where the previous sentence pointed.

> ⚠️ **CORRECTED at iter-121, in the OTHER direction.** iter-120's own repair of this paragraph said
> *"every Echo group … and nothing else"*, and cited `:230-231` / `:274-275` — **each one line short of
> the third middleware**. `cbGate := courseBuilderAccessGate(authorizationManager)` (`backend.go:227`,
> defined `internal/web/backend/gate.go:27-49`) **is** a Sentinel-backed group middleware: it requires a
> user, a non-nil active org, and `OrgCheckFeaturePermission(OrgFeatureMembersEdit, orgID)`, returning
> 401/403 before the handler. Two of the eleven groups carry it. The conclusion — no blanket, authorization
> is opt-in — survives; the absolute quantifier did not. **Same defect class as the sentence it replaced,
> pointing the other way**, and a citation that stops one line short of its own subject is exactly the
> wrong-construct class `anchor_construct_guard` does not detect.

**The honest statement of Layer 2:**
- **Sentinel is the centralized authorization *engine*** (Casbin RBAC/ABAC, PostgreSQL-backed policy
  store), and its policies are centrally managed and auditable — that part always held.
- **It is not a blanket applied to every request.** The GraphQL middleware is a *viewer* gate with a
  narrow cross-user check bolted on; the REST surface has **no blanket authz middleware** — two of its
  **eleven** groups opt into `cbGate`, two more carry a non-Clerk group gate (an API-key scope, a shared
  secret), **one (`/api/invitations`) is token-authenticated and deliberately pre-login** rather than
  Clerk-gated, and the rest authorize per handler or not at all. **And the group count is not the whole
  surface**: seven further routes are mounted on the root `e` inside no group, two of them open by
  design (`/api/schema.json`, `/content/catalog.json`) and two Development-only. *(Said "six" until
  run 81 — an under-count of the REST attack surface by five groups; said `/api/invitations` had "no
  authentication at all" until run 82, which is the opposite error, and both boxes above record it.)*
- **Where isolation actually comes from** is the three layers *together* — Layer 1 on the 31 schemas
  that declare a policy, per-resolver and per-handler checks elsewhere, and Layer 3's org-scoped
  session. Reading Layer 2 as universal is what let the M207 port ship an IDOR.
- **Nothing here should be read as a live vulnerability claim.** It is a claim about what this
  document may assert. The named taxonomy hole is fixed; the surfaces that rely on per-resolver
  checks have not been enumerated by this corpus, and that gap is now stated rather than papered over.

### Layer 3: Identity
- **Clerk** JWT tokens include organization context
- Sessions are org-scoped — users can only access their active organization
- Organization switching **re-mints the session JWT with the new org claim — it is NOT a re-authentication.** It is a client-side `clerk.setActive({ organization })` call (`next-web-app/apps/{web,hiring}/src/hooks/useOrgSelection.tsx:94`, `useResolveActiveOrg.tsx:107`, `useActivateMembershipOrg.tsx:81`); no credential is re-presented and no sign-out occurs. The isolation that follows comes from the new claim, not from a fresh proof of identity

---

## Backup & Disaster Recovery

| Aspect | Detail |
|:-------|:-------|
| **Full backups** | **Not currently running.** The `db-backup` job (**Bash**, `pg_dump` → **S3 + a Hetzner Storage Box — two destinations, never Azure**) is deployed but its EventBridge trigger has been commented out since `7dd1b80` (2025-05-29), and production pins that commit. **This row read *"Every 6 hours → S3, Azure, Hetzner (Germany)"* until M257x iter-124.** [`db-backup.md`](../services/db-backup.md) |
| **Point-in-time recovery** | RDS automated backups + an **hourly AWS Backup plan** — this, not the offsite job, is what carries durability today. What the stalled `db-backup` job costs is the **offsite, non-AWS leg**, not recoverability |
| **Primary region** | EU-West-1 (Ireland) |
| **DR site** | US AWS region |
| **Deployment** | Multi-AZ with auto-scaling |
| **CDN** | Worldwide (Vercel for frontend) |

**RETRACTED 2026-08-07.** `db-backup` does **not** run on a schedule and does **not** write three geographies. Its EventBridge rule and target have been commented out since `7dd1b80` (2025-05-29) — the commit production pins (`infrastructure/terraform/production/services.tf:571`, `ref=v0.3.3`) — and it has only ever had **two** destinations, S3 and a Hetzner Storage Box; **Azure appears in none of the 157 objects the repo has ever contained**. **Durability today is AWS-native only**: RDS `multi_az = true` with `backup_retention_period = 7` plus an hourly AWS Backup plan with continuous PITR (`infrastructure/modules/core/storage/rds.tf:6,19,78-89`). **The offsite, non-AWS copy has not been written for over a year**, so *"resilient to a full AWS region failure"* no longer holds. Full derivation: [`db-backup.md`](../services/db-backup.md).

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
| **AI Token Tracking** | Centralized usage, latency, and cost tracking in **`app/internal/aiusage`** — **not** the shared `ai` library, which only wraps providers (consistent with the shared **`ai`** library's own README + [`ai_architecture.md`](ai_architecture.md)). **The line pin was REMOVED at M257x iter-126, not re-pinned.** It read a bare `README.md:21`, which is ambiguous across every repo; qualifying it to the `ai` repo's README at line 21 made it *worse* — the resolver bound it to **`studio-desk` @ `41ee357`**, a repo with nothing to do with the shared Go library, and landed on a blank line. **`ai` is a private Go module that no stack clones** — not in `repos.yml`, pulled at Docker build via `GOPRIVATE` — so **no `file:line` into it is verifiable from here**, and an anchor that resolves to the wrong repo is more expensive than no anchor at all |

- Structured logging uses Go `slog` + Sentry integration
- ECS auto-scales on CPU/memory metrics

---

## Compliance

### EU Data Residency
- **Primary region**: EU-West-1 (Ireland)
- AI provider clients are **EU-resident by default** — Azure OpenAI EU (`ai.go:262-266`), AWS Bedrock pinned
  to `eu-west-1` (`:85-88`). **Not "routed through EU endpoints first"**: there is no ordered EU-first
  fallback chain (`external_services.md:579`), and the wording mattered because the two US paths *inside
  the AI manager* — the two the bullets below cover — are a flag and a retry target, which a "first"
  implies are tried only after an EU option fails. Corrected M257x iter-46. **Those two are not the whole
  set**: [`external_services.md:602-607`](./external_services.md) enumerates **four live** ways a request leaves
  the EU, the other two being `ANTHROPIC_API_KEY` and an authored sequence with `ai_vendor` unset — the
  latter reaching direct US OpenAI unconditionally, on the first attempt. A fifth arm, **Studio-Room's own
  `openai` `TARGET SERVICE`**, is a bare client against `https://api.openai.com` that **no shipped config
  selects** (all three `app/studio/configs/*.ini` pin `azure`). Scope corrected M257x iter-48,
  count corrected to five at iter-49 and to four-live-plus-one-latent at iter-52
- **⚠️ "EU-first" is not "EU-only", and the US path is a FLAG, not a fallback.** `getClient` swaps
  `azureClientEu` → **`azureClientUs`** whenever the PostHog flag **`flag_use_azure_us`** is enabled
  (`app/internal/jobsimulation/ai/ai.go:263-277`). That is a deliberate switch that can route live
  simulation traffic to a US region with no error condition involved. Direct OpenAI is additionally used as
  the **retry target on HTTP 429** (`isThrottlingError`, `:129` / `:166` / `:325`)
- **Anthropic is reached through AWS Bedrock `eu-west-1` from the AI manager (`:85-95`) — but "Anthropic
  Direct is not used at all" is FALSE at platform HEAD.** Course Builder routes **every** model call to
  first-party `api.anthropic.com` whenever `ANTHROPIC_API_KEY` is set:
  `app/internal/coursebuilder/bedrock.go:109-112` @ `2035f9a` (`newUnderlyingClient` → `NewAnthropicClientWithModel`; re-pinned M257x iter-126),
  with `ModelBackendName()` (`:100`) returning `"anthropic-api"` to say so. That is a **US-terminating**
  path outside the Bedrock EU region, selected by an env var rather than a flag — so it is not covered by
  the `flag_use_azure_us` caveat below. [`external_services.md:569`](./external_services.md) carries the
  provider row and [`coursebuilder.md`](../services/coursebuilder.md)'s **LLM usage — the backend is SELECTED AT START-UP** bullet calls it *"the shipped path"*; this section said the opposite.
  Corrected M257x iter-46 — *the anchor said `:489`, which is a TypeScript codegen comment, because it was
  transcribed from a blocker ledger instead of re-derived; corrected iter-48*
- No customer data stored in US **by default** — but the residency guarantee is contingent on
  `flag_use_azure_us` being off, and that is a runtime flag, not a build-time property. **Check the flag
  before asserting residency.**

### EU AI Act
- AI Simulations classified as **Limited Risk** (not High Risk)
- Stated reason: AI is used for conversation/generation only; scoring is deterministic (rubric-based, 0-100 scale), NOT AI-scored
- Stated consequence of that classification: transparency obligations only, not the strict requirements of High Risk systems

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
>
> **Both bullets above are what is STATED, not what this corpus asserts** — including the consequence
> bullet. It previously sat *after* this blockquote, at column 0, drawing the operative legal consequence
> from the classification the blockquote had just retracted three lines earlier; the retraction had been
> spliced into the middle of the list and the list resumed on the far side of it. Moved back inside the
> stated-rationale list so the retraction governs it. Repaired M257x iter-46.

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
