# Customer API + MCP — Programmatic-Access Spec

> **Status:** Consolidated draft `v0.1` · spec-draft · 2026-07-01
> **Companions:** [`spec-progress.md`](spec-progress.md) (decision tracker + log) · [`next-release.md`](next-release.md) (out-of-scope / parking lot) · [`vision.md`](vision.md) (north-star + long-horizon posture)
> **Brand:** *Anthropos Public API* + *Anthropos MCP Server* — one contract, two shells (REST-and-later-GraphQL for scripts, MCP for AI agents).

This spec defines a **new pillar** for the Anthropos platform: **programmatic, customer-facing access** to
Anthropos data and operations, exposed as (1) a versioned public **HTTP API** and (2) a **Model-Context-Protocol
(MCP) server** over the same contract. It defines the **capability, the model, the principles, and the tech
approach** — it does **not** implement any endpoint or write path (that is the work this spec governs).

The pillar is deliberately **read-first** (self-serve reporting, AI-agent discovery), with a **strict, staged
opening** of write surfaces behind an audit + rate-limit + entitlement floor.

---

## 1. Overview

### 1.1 North star — what it is

Today Anthropos is reachable only through its own UIs (Workforce, Academy, Studio, the mobile app). A customer
who wants to *"pull our roster into our HRIS every night"* or *"let our AI agent answer questions about our
learners"* has no sanctioned path. Support tickets fill the gap; the actual code path a partner would use does
not exist.

The Customer API + MCP pillar closes that gap. It gives customers and their AI agents **one contract, two
shells** — a REST façade for scripts + integrations, and an MCP server for LLM-driven consumption — over the
**same resource model, the same auth, the same audit log, and the same rate limits.**

**How it feels:**
- A **customer developer** signs into Workforce, mints an API key with a scope, and hits `GET /v1/organizations/{id}/members` from `curl`. It returns *their* org's members, and only theirs.
- An **AI agent** running a customer-installed Anthropos MCP server answers *"who on the team has verified skill X?"* by calling the same resource under the same principal.
- A **compliance officer** exports the audit log for the last 30 days and sees every read + write, keyed by principal + scope + resource.

### 1.2 The goal

A **versioned, tenant-isolated, principal-audited** programmatic contract that:

- exposes the **read surface** first (safe, low-risk, immediately useful for reporting + AI agents),
- opens **writes** in staged tiers behind an audit + rate-limit + entitlement floor (never a free-for-all),
- is **auth-vendor-independent** — the API depends on an internal `Principal` abstraction, not on Clerk's
  identity model, so a future SSO/identity swap is contained,
- **doubles** as: a customer-facing SDK/docs foundation, an MCP source for AI agents, and the internal contract
  the platform's own surfaces can consolidate on over time.

### 1.3 Scope of this spec

**In scope:** the **capability definition**, the **auth-layer-independence principle**, the **API resource
catalog** (Products → Resources → Actions), the **release roadmap** (R1..Rn — what ships when), the
**mutation gap analysis** (which write actions exist as platform mutations today vs which the platform must
add), the **MVP definition** (the smallest customer-usable slice), and the **tech approach** (REST façade over
existing Connect-RPC / GraphQL, MCP over the same, key management, rate limits, audit).

**Explicitly out of scope here:** implementing any endpoint or MCP tool; SDK code-generation; partner-marketplace
mechanics; billing/metering-for-revenue (metering-for-limits is in). Enumerating the *full* endpoint list per
release lands in the milestone specs, not here.

---

## 2. Goals

The eight outcomes this pillar exists to achieve. Each subsequent design decision must serve at least one.

| # | Goal | Kind |
|---|------|------|
| **G1** | **Programmatic customer access** — a customer developer can script recurring HR operations (roster pull, path assignment, verification query) without a human clicking Workforce. | product |
| **G2** | **AI-agent access via MCP** — an AI agent (Claude, or a customer-hosted LLM) can drive Anthropos as an MCP source: discover resources, read state, take sanctioned actions. | product |
| **G3** | **Ecosystem enablement** — HR-tech partners (HRIS mirrors, LMS bridges, SIEM sinks) can integrate against a stable, documented contract instead of a bespoke agreement. | business |
| **G4** | **Reduce internal support load** — self-serve reads (roster, progress, verified skills, audit) replace the "please export X for us" tickets that operators handle today. | business |
| **G5** | **SDK / docs foundation** — one contract carries the SDK, the quickstarts, and the API reference; every future paying tier lands on it, not on N per-customer scripts. | business |
| **G6** | **Audit + rate-limit baseline** — every access (read + write) is keyed by principal + scope + resource + timestamp; every principal has a rate-limit budget. Compliance-critical writes have this **from R1**, not "added later." | platform |
| **G7** | **Auth-vendor independence** — the API depends on an internal `Principal` abstraction (org, user, scope, entitlement tier), **not** on Clerk-specific claims/IDs. Swapping the identity vendor changes one adapter, not every consumer. | architecture |
| **G8** | **MCP-first surface for AI-native product** — establish Anthropos as a first-class MCP source, not an API that later "gets" an MCP wrapper. The MCP contract ships alongside R2, on the R1 read foundation. | strategy |

**Non-goal (explicitly):** replicating the internal service-to-service Connect-RPC surface as the customer API.
This is a **curated, versioned, principal-scoped façade** — not a wire-level passthrough. (See §7.)

---

## 3. Principles

> These are the load-bearing contract. A new endpoint / MCP tool that violates a principle is wrong even if it
> works. Every reviewer holds to all of them.

- **P1 — Read-first, writes-staged.** The R1 surface is **read-only**. Writes open in **W1 / W2 tiers** (§4.5),
  each behind the audit + rate-limit floor. No write ships without an audit row + a rate-limit budget +
  entitlement gating. A "safe" write is still a staged write.

- **P2 — Tenant isolation is non-negotiable.** Every request is scoped to a `Principal`, and every resource
  fetch is filtered by that principal's `organization_id`. **No cross-tenant read path exists** in the customer
  API — even for an "admin" principal (that is a distinct, platform-internal surface, not this one). Enforced at
  the resource layer, not just at the auth layer.

- **P3 — Auth-layer independence (Clerk-swap principle).** The API depends on an **internal `Principal`
  abstraction** (§5.4) — org id, user id, scope set, entitlement tier — resolved by an **`IdentityProvider`
  adapter port**. Clerk is one implementation. **No Anthropos code above the adapter references Clerk types,
  claims, or IDs.** A future SSO / identity swap replaces the adapter, not the endpoints.

- **P4 — Versioned + additive.** Every endpoint lives under `/v{N}/...`. **Additive changes** (new field, new
  endpoint) do **not** bump the version. Breaking changes require a new `v{N+1}` and a documented deprecation
  window. Response envelopes include a `Deprecation` header per RFC 8594 when a resource is on the sunset path.

- **P5 — One contract, two shells.** The **REST façade** and the **MCP server** are two projections of the
  **same resource catalog**. An MCP tool that reads members is not a separate implementation — it delegates to
  the same read handler under the same principal. Consequences: adding a resource lights it up on both shells;
  auth + audit + rate-limits fire once, in the shared layer.

- **P6 — Every call audited, every principal budgeted.** A request that is **not** written to the audit ledger,
  or **not** counted against a rate-limit budget, is a bug. The audit surface is a first-class read resource
  (customers export their own audit log — G4). Rate limits are per-principal-per-window with sane defaults +
  per-tenant override.

- **P7 — Curated, not-a-wire-passthrough.** The customer API is a **curated façade over** the internal
  Connect-RPC / GraphQL / DB layer — never a proxy. A resource in the catalog exists because a **customer use
  case** motivates it (§4.4); an internal RPC is not exposed because it exists internally.

- **P8 — MCP is a first-class shell, not an afterthought.** Every read resource that lands in R1 is designed
  MCP-tool-shaped from the start (name, description, JSON-schema input/output, safety category). The MCP server
  in R2 is a **build**, not a **retrofit**.

- **P9 — Documented is shipped.** A resource without an OpenAPI (REST) + MCP-manifest entry + a quickstart
  example is **not shipped**. Docs land with the endpoint, in the same PR, on the same release. G5 depends on
  this being a contract, not a wish.

---

## 4. The model & vocabulary

### 4.1 Resource hierarchy

The customer API is organized in a four-level hierarchy that mirrors how a customer *thinks* about their data:

```
Product        (1) a platform capability area          (People, Learning, Verification, Simulations, Audit)
└─ Resource    (2) a noun the customer manipulates     (Member, Skill Path Assignment, Verified Skill, Session)
   └─ Action   (3) a verb over that resource           (list, get, create, update, delete, webhook.subscribe)
      └─ Tool  (4) the exposed unit                    (REST endpoint + MCP tool + audit row + rate-limit key)
```

- **Product** — a platform product / capability area under the API. The top-level grouping (used to organize
  docs + MCP tool namespaces).
- **Resource** — the **atomic customer-visible noun**. It carries an id, a schema, an audit key, and a resource
  owner (`organization_id`).
- **Action** — a verb over a resource. Actions are typed **`read`** (safe by default), **`w1`** (safe writes,
  see §4.5), **`w2`** (advanced writes), **`admin`** (org-scoped admin, e.g. rotate a key).
- **Tool** — the **atomic unit of contract**. One tool = one REST endpoint + one MCP tool + one audit row shape
  + one rate-limit bucket. **Two shells, one tool** (P5).

### 4.2 API resource catalog (indicative, first pass)

The catalog is the spec's **build reference** for the roadmap in §6. It is deliberately **shorter than what the
platform can technically expose** — it starts from customer use cases (§4.4), not from the internal RPC surface.

| Product | Resource | Actions (planned) | First release |
|---|---|---|---|
| **People** | `organization`, `member`, `role`, `manager-report-line` | `list`, `get` (R1); `w1: member.create`, `member.update`, `member.deactivate` (R4) | R1 read / R4 W1 |
| **Learning** | `skill-path`, `skill-path-assignment`, `skill-path-progress`, `certificate` | `list`, `get` (R1); `w1: assignment.create`, `assignment.reassign` (R4) | R1 read / R4 W1 |
| **Verification** | `user-skill` (claimed), `verified-skill`, `skill-taxonomy-node` | `list`, `get` (R1); `w2: verified-skill.emit` (R5, gated) | R1 read / R5 W2 |
| **Simulations** | `simulation-session`, `simulation-blueprint`, `session-result` | `list`, `get` (R1); `w2: session.launch` (R5) | R1 read / R5 W2 |
| **Audit** | `audit-event` | `list`, `get` (R1); `webhook.subscribe` (R5) | R1 read / R5 webhook |
| **Access** | `api-key`, `scope`, `rate-limit-budget` | `list`, `get`, `admin: create/rotate/revoke` (R1) | R1 admin |

**W1 (safe writes, R4)**: the "roster + assignment" cluster — high value, low blast radius, well-understood
platform mutations. **W2 (advanced writes, R5)**: verification emissions + session launches — larger blast
radius, tighter entitlement gates, per-action rate limits.

### 4.3 What a Resource declares (the tool contract)

Every catalog entry declares a **tool contract** — the atomic unit of §4.1. The contract carries:

| Field | Meaning |
|---|---|
| `id` | Stable identifier (`people.member.list`, `learning.assignment.create`). The 1:1 link across REST + MCP + docs + audit. |
| `product` / `resource` / `action` | The §4.1 coordinates. |
| `action_type` | `read` / `w1` / `w2` / `admin` (§4.1). |
| `principal_scope` | The scope-set token(s) required (§5.4). |
| `entitlement_tier` | The `Principal.tier` a caller must hold (`free` / `paying` / `enterprise` / `partner`). |
| `input_schema` | JSON Schema (REST body + MCP tool input). |
| `output_schema` | JSON Schema (REST response + MCP tool output). |
| `rate_limit_bucket` | The rate-limit key + default window/quota (§5.6). |
| `audit_shape` | The audit row shape (which fields land in the audit ledger — never the full payload for W2). |
| `docs` | OpenAPI ref + MCP-manifest ref + quickstart pointer (P9 enforcement). |

### 4.4 Customer use cases (the source of truth for the catalog)

The catalog is derived from a **numbered use-case list** — each row is a real thing a real customer does. Use
cases marked **FIRST-USABLE** are the R1 MVP surface (§6.2).

| # | Use case | Persona / JTBD | Value | Kind | Ships |
|---|---|---|---|:---:|:---:|
| **UC1** | List all members of my org | HR ops · "keep our HRIS in sync" | replace nightly CSV export | READ | R1 **FIRST-USABLE** |
| **UC2** | Fetch skill-path progress for a team | HR ops · "show weekly training pulse" | replace the manual roll-up | READ | R1 **FIRST-USABLE** |
| **UC3** | Fetch verified-skill state per employee | HR ops · "compliance report" | replace one-off SQL | READ | R1 **FIRST-USABLE** |
| **UC4** | Pull org structure + reporting lines | HR ops · "sync org chart into HRIS" | | READ | R1 **FIRST-USABLE** |
| **UC5** | Manager pulls team competency snapshot for review | Manager · "1:1 prep" | replace UI clicking | READ | R1 |
| **UC6** | Compliance officer exports the audit log | Compliance · "SIEM feed" | new capability | READ | R1 |
| **UC7** | Ecosystem partner mirrors Anthropos org data into HRIS | Partner · "keep systems in lockstep" | ecosystem enablement (G3) | READ | R1 |
| **UC8** | AI agent (Claude) answers *"who has verified skill X?"* over MCP | AI agent (G2, G8) | new capability | READ (MCP) | R2 **FIRST-MCP** |
| **UC9** | AI agent recommends *"what should Bob take next?"* over MCP | AI agent (G2) | new capability | READ (MCP) | R2 |
| **UC10** | Onboard a batch of new employees | HR ops · "new-hire ingest" | replace batch UI | WRITE (W1) | R4 |
| **UC11** | Assign a skill path to a team | HR ops / L&D · "assign training" | replace UI | WRITE (W1) | R4 |
| **UC12** | Update org structure / reassign a manager | HR ops · "reorg" | replace UI | WRITE (W1) | R4 |
| **UC13** | Ecosystem app launches an AI simulation on behalf of a user | Partner · "embed sim in partner app" | new capability | WRITE (W2) | R5 |

The **FIRST-USABLE flag** lands on UC1–UC4 — enough that a customer can *do something real end-to-end* the day
R1 ships (roster + progress + verified skill + org chart, all self-serve). **UC8 is the FIRST-MCP** — R2's proof
that the MCP shell works on the R1 read foundation. Writes stay parked until R4/R5 behind the audit floor (P1).

### 4.5 Write staging — W1 / W2 / admin

Writes never ship in the same release as the read foundation. The staging:

- **`admin`** (R1) — org-scoped admin over the **Access** product only (mint/rotate/revoke API keys, view
  rate-limit budgets). No customer-data mutation. Ships with R1 because the reads need keys.
- **`w1`** (R4) — the **safe writes cluster**: create/update/deactivate members, create/reassign skill-path
  assignments, update org structure. Well-understood platform mutations, low blast radius, high customer value.
- **`w2`** (R5) — the **advanced writes cluster**: emit a verified skill, launch a simulation session, subscribe
  to webhooks. Larger blast radius, tighter entitlement gates (`paying` or `enterprise` only), per-action rate
  limits, in some cases require a signed provenance claim (e.g. verified-skill emission).

**No write endpoint bypasses the audit floor or the rate-limit budget** (P6). A write's audit row records
principal + resource + action + input hash (never the raw input for W2, per privacy).

---

## 5. Tech approach

### 5.1 The REST façade

- **Layer:** a new **`app`-hosted** REST layer (customer-API) that lives *above* the internal Connect-RPC +
  GraphQL surface and delegates to them per-endpoint. **Not** a new microservice in R1 (v3.0's rule: minimize
  new services; the façade is a package inside `app`).
- **Envelope:** JSON responses, HTTP-standard status codes, RFC 7807 `problem+json` for errors, cursor-based
  pagination (`cursor` + `next_cursor`), `ETag` + `If-None-Match` for cacheable reads, RFC 8594 `Deprecation`
  header on sunset-path resources.
- **Versioning:** URL-versioned (`/v1/...`). Additive changes never bump the version (P4).
- **Rate limiting:** shared token-bucket keyed by `Principal.id` + `rate_limit_bucket`, backed by Redis (already
  in the stack).
- **Content negotiation:** JSON only in R1; ND-JSON streaming for large list endpoints is a R3 candidate.

### 5.2 The MCP server

- **Shell:** an **MCP server** (per the MCP specification) that exposes the R1 catalog as tools. Delegates to the
  same handlers as the REST façade — one contract, two shells (P5).
- **Discovery:** the MCP `tools/list` return is generated **from the catalog** — a resource-action entry with an
  `input_schema` is an MCP tool with that schema. No hand-maintained tool list.
- **Auth:** the MCP client presents the same API key (via the sanctioned MCP auth header). The server resolves
  the `Principal` via the same `IdentityProvider` adapter as REST (§5.4).
- **Safety category:** every MCP tool declares a `safety` field (`safe-read`, `mutating-w1`, `mutating-w2`,
  `admin`) so a hosting MCP client can gate the tool per its own policy.
- **Deployment:** the MCP server is a **binary**, distributable as a Docker image + a `npx anthropos-mcp-server`
  wrapper (R2). Customer-hosted by default; a hosted variant is R6.

### 5.3 Where it lives

- **REST façade + shared handler layer:** in `app` (`internal/customerapi/` package). Shares the internal-service
  RPC + DB clients; adds the auth-layer independence adapter, the resource catalog registry, the audit + rate-
  limit middleware.
- **MCP server:** its own repo `anthropos-work/anthropos-mcp-server` (Go, sharing the Connect-RPC client
  code-gen from `proto`). Lives outside the platform monorepo because it's a distributable binary — but its tool
  set is **generated from the platform-owned catalog**, not authored.
- **Docs:** OpenAPI spec generated from the catalog; hosted on `docs.anthropos.work/api/v1/` (a new surface in
  R1, minimally styled). MCP manifest hosted at `mcp.anthropos.work/manifest.json`.

### 5.4 Auth-layer independence — the Principal + IdentityProvider adapter

The load-bearing architectural principle (G7 / P3). The internal contract:

```
Principal {
  id                 string        // stable internal id, NOT a Clerk id
  organization_id    string
  user_id            string?       // null for org-scoped API-key principals
  scopes             []Scope       // e.g. ["people:read", "learning:read"]
  entitlement_tier   Tier          // free / paying / enterprise / partner
  identity_source    string        // "clerk" (today) / "saml:acme" (future)
}
```

**Every customer-API handler receives a `Principal`, never a Clerk claim.** The `IdentityProvider` port has one
adapter implementation today (`ClerkIdentityProvider`) plus one for API keys (`ApiKeyIdentityProvider`). A future
identity swap adds an implementation; the customer-API code above the port does **not** change.

**Concrete forbiddens** (P3, enforced by lint + review):
- No `clerk.*` import in `internal/customerapi/`.
- No Clerk user id, org id, or session id in any response body.
- No `sub` / `org_id` claim read outside the adapter package.

The **API-key primitive** is a first-class Principal source (§5.5) — a mint/rotate/revoke path that produces a
`Principal` at request time, sitting alongside the Clerk-JWT path (which the customer API accepts too, for
Workforce's own scripts). Both routes converge on the same `Principal` and the same handler.

### 5.5 API keys — the customer-facing credential

- **Shape:** `ak_live_<random>` / `ak_test_<random>`. Prefix routes to the environment; the random tail is
  256-bit, stored **hashed** (bcrypt/argon2id — decision at build), never in plaintext at rest.
- **Scoping:** at mint time the customer selects a **scope set** (e.g. `people:read`, `learning:read`,
  `audit:read`). Least-privilege is the default UI (no "all scopes" button).
- **Lifecycle:** mint → rotate (issues a new key, revokes the old with a grace window) → revoke (immediate).
  Every lifecycle event is audited.
- **UI:** the R1 MVP includes a **customer self-serve** API-key page in Workforce (list, mint, rotate, revoke,
  view usage). Not the studio-desk-style separate surface — inline in Workforce settings.
- **Rate-limit budget:** each key has a default per-window budget (R1 defaults: 60 req/min, 10k req/day). Tenant
  overrides through the platform-internal admin surface (not the customer API).

### 5.6 Audit + rate limits

- **Audit ledger:** append-only Postgres table `customer_api.audit_events` (columns: `id`, `ts`, `principal_id`,
  `organization_id`, `resource_id`, `action`, `status`, `input_hash`, `latency_ms`, `client_ip`, `user_agent`).
  W2 writes carry only an input hash + a length; the raw input is not persisted.
- **Retention:** 90 days hot in Postgres, older shipped to S3 (mirrors the `execution_traces` retention pattern
  from other services).
- **Read surface:** the audit ledger *is* a customer resource — `GET /v1/audit-events?since=…` (UC6).
- **Rate limits:** shared token-bucket in Redis, keyed by `Principal.id` + `rate_limit_bucket`. Response includes
  `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset`. HTTP 429 with a `Retry-After` on exhaustion.

### 5.7 Testing posture

- **Contract tests:** every catalog entry has a contract test (input/output schema conformance, principal
  isolation, rate-limit fires, audit row lands).
- **Cross-tenant isolation gauntlet:** a dedicated suite that runs every read endpoint under Org A's principal
  and asserts **zero** Org B rows leak. Ships **at R1**, not later (P2).
- **MCP conformance:** R2's MCP server ships with an MCP-spec-conformance suite (tool discovery, schema, error
  shape).
- **Playthroughs sibling:** the R1 customer flows (UC1–UC4 self-serve) are also candidate Playthroughs on the
  v2.0 foundation once R1 lands — but this is a follow-on, not R1 scope.

---

## 6. Release roadmap

The pillar is a multi-release program. R1 is the MVP; each subsequent R adds one clean tier.

### 6.1 The roadmap

| Release | Codename | Goal | Ships |
|---|---|---|---|
| **R1 — this spec's MVP** | **`v3.0` "open house"** | **Read-only REST + API-key self-serve + audit + rate-limit floor.** The FIRST-USABLE UCs (UC1–UC4, plus UC5–UC7 read). Docs site v1. | R1 |
| **R2** | *(v3.1, tbd)* | **MCP shell** over the R1 read catalog. UC8 (FIRST-MCP) + UC9. The MCP server binary. | R2 |
| **R3** | *(v3.2, tbd)* | **Query enrichments** — GraphQL projection over the catalog; ND-JSON streaming for large lists; server-side aggregations for common report shapes. | R3 |
| **R4** | *(v3.3, tbd)* | **W1 writes GA** — the safe writes cluster (member CRUD, assignment CRUD, org structure updates — UC10–UC12). Full audit + per-action rate-limits + entitlement gates. | R4 |
| **R5** | *(v3.4, tbd)* | **W2 writes + webhooks** — verified-skill emissions, session launches, webhook subscriptions (UC13). Tighter entitlement gates + signed provenance where applicable. | R5 |
| **R6** | *(v3.5, tbd)* | **GA hardening + SLA** — a hosted MCP variant, SDK code-gen (TS + Python), formal SLA + status page + on-call. | R6 |

**Coexistence:** through R1–R5 the customer API + Workforce UI coexist — every action available on the API is
also available in Workforce; the UI is not deprecated for customers who prefer it. R6 does not change this.

### 6.2 The MVP (R1 — the smallest customer-usable slice)

**In R1:** the FIRST-USABLE reads (UC1–UC4), the audit + rate-limit floor, API-key self-serve mint/rotate/revoke,
minimal docs. Just enough for a customer to script the nightly HRIS pull.

**Explicitly deferred from R1:**
- Writes of any kind → **R4/R5**.
- The MCP shell → **R2** (fast-follow, on the R1 foundation — same catalog, no new resources).
- Aggregations, streaming, GraphQL projection → **R3**.
- SDK code-gen, hosted MCP, SLA → **R6**.

**Why this shape:** it is the **shortest end-to-end customer-usable slice** — a real customer replaces a real
support ticket with a real curl call on day one — while establishing the audit + rate-limit + Principal +
API-key floor that every subsequent release inherits. **Read-first + audit-from-R1 = writes-cheap-later** (P1
+ P6).

### 6.3 R1 milestone shape (governed by this spec's `/developer-kit:design-roadmap` run)

Registered in [`../../roadmap-vision.md`](../../roadmap-vision.md) as a **proposal for v3.0**. Four milestones,
sequential:

- **M301 — Discovery + Identity seam** (`section`) — the API resource catalog registry; the `Principal` DTO +
  `IdentityProvider` adapter port; the Clerk adapter; the `ClerkGuardrails` lint (no `clerk.*` import above the
  adapter). No public endpoint yet.
- **M302 — Access primitive** (`section`) — the API-key mint/rotate/revoke path; the `ApiKeyIdentityProvider`;
  the audit ledger table + append-only middleware; the rate-limit middleware (Redis token-bucket). No customer
  data endpoint yet — the primitive is proven with a `/v1/access/whoami` echo.
- **M303 — REST reads gateway** (`iterative`) — the R1 read catalog (People + Learning + Verification + Audit),
  one resource at a time, each closed on: OpenAPI entry + contract test + cross-tenant isolation test + audit
  row + rate-limit fire. **Exit gate:** UC1–UC7 all green on an integration stack with 0 cross-tenant leakage
  over N runs.
- **M304 — Customer surface + docs lite** (`section`) — the Workforce self-serve API-key UI (list / mint /
  rotate / revoke / usage) + the `docs.anthropos.work/api/v1/` docs site (OpenAPI-generated + hand-written
  quickstart for UC1–UC4).

Execution graph: **M301 → M302 → M303 → M304** (strictly sequential; each depends on the prior).

---

## 7. Relationship to existing Rosetta / Anthropos capabilities

| Capability | Relationship |
|---|---|
| **Internal Connect-RPC / GraphQL federation** | The **substrate**. The customer API delegates to it — never proxies it. Curated, not passthrough (P7). |
| **Sentinel (authz)** | The customer API's `Principal` carries scopes; the handler asks Sentinel per-resource. Sentinel is unchanged. |
| **Clerk (authn)** | One of two `IdentityProvider` adapters (§5.4). The customer API depends on the port, not on Clerk. |
| **Playthroughs (v2.0)** | Follow-on — R1's UC1–UC4 customer flows are candidate Playthroughs on v2.0's foundation, once R1 lands. Not R1 scope. |
| **Path migration (R2 spec-draft)** | Independent. The customer API exposes skill-path *reads* regardless of which engine is behind them (P7's curated principle absorbs the engine coexistence). |
| **Studio Desk / Academy authoring** | Out of scope — authoring is not a customer-API concern; the API exposes *consumption + operations*. |

---

## 8. Out of scope for this pillar (non-goals)

Anti-pillar. Anything on this list stays out (parked in [`next-release.md`](next-release.md)):

- **Service-to-service internal RPC replacement** — the internal Connect-RPC surface stays as it is; the
  customer API is a *curated façade*, not a replacement.
- **Billing / metering-for-revenue** — metering-for-limits (rate-limit budgets) is in R1; metering-for-invoicing
  is a separate program.
- **Partner marketplace** — a directory of third-party apps built on the API is a future business decision, not
  a spec item.
- **Public webhook broker** — the R5 webhook subscription is per-customer, not a public event bus.
- **SDK code-generation** — deferred to R6 (the OpenAPI spec supports it; the effort is a real one and lives on
  its own release beat).
- **Hosted MCP endpoint** — R2 ships the MCP server as a customer-hosted binary. A hosted variant is R6.

---

## 9. Open / to-confirm

Tracked in [`spec-progress.md`](spec-progress.md). The load-bearing decisions **for R1 (the MVP)** are settled
here: the goal set (§2), the principles (§3), the resource model + write staging (§4), the auth-layer
independence contract (§5.4), the API-key primitive (§5.5), the audit + rate-limit floor (§5.6), the R1
milestone shape (§6.3).

**Carried as open build items** (decided *in shape*, work deferred):

- **API-key hashing algorithm** (argon2id vs bcrypt) — decide at M302 build.
- **OpenAPI vs a homegrown catalog format** for the machine source — decide at M301 build; leaning OpenAPI 3.1
  with a small `x-anthropos-*` extension for MCP fields.
- **MCP hosted vs customer-hosted default** — decide at R2 design; leaning customer-hosted (customer holds the
  key, hosts the binary).

**Deferred after this spec:** the R2–R6 milestone plans (each gets its own `/developer-kit:design-roadmap` run
on its own release beat).

---

## Appendix A — Real-mutation gap analysis

The customer API is a curated façade; it can only expose a mutation the platform actually implements. This
appendix maps each planned write action to its underlying platform mutation and records the gap.

**Method:** cross-reference every W1 / W2 action in §4.2 against the internal Connect-RPC + GraphQL surface
(from `corpus/services/*.md`). Mark as `exists` (fully implemented), `partial` (implemented but not exposed to
this shape of caller), `missing` (would require a platform-repo edit).

| Action | Underlying mutation | Status | Notes |
|---|---|---|---|
| `people.member.create` | `app`: `Backend.CreateMembership` + `Skiller.LinkUser` | **exists** | Full path via Workforce today. |
| `people.member.update` | `app`: `Backend.UpdateMember` | **exists** | |
| `people.member.deactivate` | `app`: `Backend.DeactivateMembership` | **exists** | Soft-delete; audit row required. |
| `people.org-structure.update` | `app`: `Backend.SetManagerReport` | **exists** | |
| `learning.assignment.create` | `skillpath`: `Skillpath.AssignPath` | **exists** | |
| `learning.assignment.reassign` | `skillpath`: `Skillpath.ReassignPath` | **partial** | Exists via internal admin; not exposed to customer-scoped principal. Needs handler that accepts a `Principal` and enforces scope. |
| `verification.verified-skill.emit` | `skiller`: `Skiller.EmitVerifiedSkill` | **partial** | Emission is engine-driven today (sim completion); a customer-driven emission needs a **provenance-claim** field on the RPC (small platform edit). |
| `simulations.session.launch` | `jobsimulation`: `JobSimulation.StartSession` | **exists** | Requires blueprint id + user id. |
| `audit.webhook.subscribe` | *(none)* | **missing** | Requires a webhook-broker + a subscription record — non-trivial platform edit (R5). |

**Escalation policy (mirrors Playthroughs' `unimplementable-without-platform-edit`):** a `missing` mutation
**does not** get quietly shimmed in the customer API — it is **escalated** to the platform roadmap, and the
customer-API endpoint stays behind an `unimplemented` state until the platform mutation lands. The customer API
never invents a mutation the platform doesn't own.

