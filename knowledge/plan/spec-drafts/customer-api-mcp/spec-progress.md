# Customer API + MCP — Spec Progress (open points tracker)

> **Status:** Draft · spec-draft · 2026-07-01 (tracks [`spec.md`](spec.md) `v0.1`)
> Tracker + decision log for the spec. Decisions are worked **one at a time** and recorded here.

**Legend:** 🔴 not decided · 🟡 discussing / proposed · ✅ decided · ⏭️ deferred (→ [`next-release.md`](next-release.md))

| # | Topic (plain English) | Status | Decision |
|---|------------------------|:------:|----------|
| A | What is the capability? | ✅ | A **versioned, tenant-isolated, principal-audited** programmatic contract to Anthropos, exposed as one contract in two shells: **REST façade** for scripts/integrations and **MCP server** for AI agents. [`spec.md`](spec.md) §1. |
| B | Read-first vs write-first | ✅ | **Read-first.** R1 = reads + audit + rate-limits + API-key self-serve; writes stage in W1 (R4) + W2 (R5). No write ships without the audit + rate-limit + entitlement floor. [`spec.md`](spec.md) §3 P1, §4.5, §6. |
| C | The auth-vendor swap principle | ✅ | **Auth-layer independence.** Internal `Principal` DTO + `IdentityProvider` adapter port; Clerk is one adapter. **No `clerk.*` import above the adapter package.** Enforced by lint + review. [`spec.md`](spec.md) §3 P3, §5.4. |
| D | One contract or two? | ✅ | **One contract, two shells.** REST + MCP are two projections of the **same resource catalog**; adding a resource lights up both shells; auth + audit + rate-limits fire once, in the shared layer. [`spec.md`](spec.md) §3 P5. |
| E | The MVP (smallest customer-usable release) | ✅ | **R1 = read-only REST + API-key self-serve + audit + rate-limit floor + docs lite.** FIRST-USABLE UCs: UC1–UC4 (roster, path progress, verified skill, org chart). MCP fast-follows in R2 on the same catalog. [`spec.md`](spec.md) §6.2. |
| F | The resource model | ✅ | Four-level: **Product → Resource → Action → Tool**. One tool = one REST endpoint + one MCP tool + one audit row shape + one rate-limit bucket. [`spec.md`](spec.md) §4.1. |
| G | The catalog first pass | ✅ | Products: **People, Learning, Verification, Simulations, Audit, Access**. Actions: `read` (R1), `w1` (R4), `w2` (R5), `admin` (R1, Access only). [`spec.md`](spec.md) §4.2. |
| H | The customer use-case list | ✅ | 13 UCs, FIRST-USABLE = UC1–UC4 (R1 self-serve), FIRST-MCP = UC8 (R2). [`spec.md`](spec.md) §4.4. |
| I | Write staging model | ✅ | **W1 / W2 / admin** tiers. W1 (R4) = safe writes (member/assignment/org-structure). W2 (R5) = verification emit + session launch + webhooks; tighter entitlement gates, per-action rate limits. Admin (R1) = API-key management only. [`spec.md`](spec.md) §4.5. |
| J | Where the façade lives | ✅ | **Inside `app`** as an `internal/customerapi/` package (not a new microservice in R1). MCP server lives in its own repo `anthropos-mcp-server` (distributable binary), but its tool set is generated from the platform-owned catalog. [`spec.md`](spec.md) §5.3. |
| K | API-key shape + lifecycle | ✅ | `ak_live_<random>` / `ak_test_<random>`; 256-bit random tail; stored **hashed** at rest; scoped-at-mint; mint / rotate (with grace) / revoke; every lifecycle event audited. Self-serve UI in Workforce. [`spec.md`](spec.md) §5.5. |
| L | Audit surface | ✅ | Append-only Postgres `customer_api.audit_events`; **90 days hot + S3 archive**; **read surface for customers** (`GET /v1/audit-events`). W2 writes carry only `input_hash` (never raw payload). [`spec.md`](spec.md) §5.6. |
| M | Rate limits | ✅ | Redis token-bucket, per-`Principal.id` + `rate_limit_bucket`. Default budgets (60 req/min, 10k req/day); per-tenant override via platform-internal admin. Standard `X-RateLimit-*` headers + 429 + `Retry-After`. [`spec.md`](spec.md) §5.6. |
| N | Versioning contract | ✅ | URL-versioned `/v{N}/...`. **Additive changes** don't bump the version; breaking changes need a new `v{N+1}` + deprecation window (RFC 8594 `Deprecation` header). [`spec.md`](spec.md) §3 P4. |
| O | Docs discipline | ✅ | **Documented is shipped** (P9): a resource without an OpenAPI entry + an MCP-manifest entry + a quickstart is not shipped. Docs land in the same PR. [`spec.md`](spec.md) §3 P9. |
| P | Cross-tenant isolation testing | ✅ | Dedicated cross-tenant isolation gauntlet suite runs every read endpoint under Org A's principal and asserts **zero** Org B rows leak. Ships **at R1**, not later (per P2). [`spec.md`](spec.md) §5.7. |
| Q | Curated vs passthrough | ✅ | **Curated façade** (P7). A resource in the catalog exists because a customer UC motivates it; an internal RPC is not exposed just because it exists internally. [`spec.md`](spec.md) §3 P7. |
| R | MCP-first posture | ✅ | Every R1 read resource is designed **MCP-tool-shaped from the start** (name, description, JSON-schema I/O, safety category). R2's MCP server is a build, not a retrofit. [`spec.md`](spec.md) §3 P8, §5.2. |
| S | R1 milestone shape | ✅ | 4 milestones, sequential: **M301** discovery + identity seam (`section`) → **M302** access primitive (`section`) → **M303** REST reads gateway (`iterative`) → **M304** customer surface + docs lite (`section`). [`spec.md`](spec.md) §6.3. |
| T | Mutation gap posture | ✅ | Real-mutation gap analysis is Appendix A. **Missing mutations are escalated, never shimmed** — the customer API never invents a mutation the platform doesn't own. Mirrors Playthroughs' `unimplementable-without-platform-edit` state. [`spec.md`](spec.md) Appendix A. |
| 1 | API-key hashing algorithm (argon2id vs bcrypt) | 🟡 | Decide at **M302 build**. Leaning **argon2id** (modern default; OWASP guidance). |
| 2 | OpenAPI vs homegrown catalog format for the machine source | 🟡 | Decide at **M301 build**. Leaning **OpenAPI 3.1 + small `x-anthropos-*` extension** for MCP fields. |
| 3 | MCP hosted vs customer-hosted default | 🟡 | Decide at **R2 design**. Leaning **customer-hosted binary** (customer holds the key, hosts the binary); hosted variant deferred to R6. |
| 4 | ND-JSON streaming for large list endpoints | ⏭️ | R3. Not R1 scope — cursor pagination suffices for the R1 UCs. |
| 5 | GraphQL projection over the catalog | ⏭️ | R3. |
| 6 | Public webhook broker | ⏭️ | Non-goal — per-customer webhook subscriptions are in R5 (§4.5), a public event bus is not a spec item ([`next-release.md`](next-release.md)). |
| 7 | SDK code-generation (TS + Python) | ⏭️ | R6. |
| 8 | Formal SLA + status page + on-call | ⏭️ | R6. |
| 9 | Partner marketplace / directory | ⏭️ | Non-goal ([`next-release.md`](next-release.md)). |
| 10 | Billing / metering-for-revenue | ⏭️ | Non-goal — metering-for-limits is in R1; metering-for-invoicing is a separate program ([`next-release.md`](next-release.md)). |

---

## Decision log

### Points A–T — the capability, model & MVP (decided 2026-07-01, from the founding brief)

The founding direction fixed the spine: **read-first, writes-staged; one contract, two shells; auth-vendor-
independent; docs is shipped; audit + rate-limit floor at R1, not later.** The MVP is the smallest slice a real
customer can use end-to-end on day one (UC1–UC4 self-serve reads + the API-key self-serve UI + docs lite), on a
foundation (Principal + IdentityProvider + audit ledger + rate-limit bucket) that every subsequent release
inherits. Writes stage in W1 (R4) + W2 (R5) behind that floor. MCP is not a retrofit — it's a fast-follow (R2)
on the same R1 catalog, and every R1 resource is MCP-tool-shaped from the start.

### Rationale — read-first (Point B)

Reads are safe by construction (no state change → no undo needed), immediately valuable (they replace real
support tickets — G4), and force the design of the whole non-write floor (Principal + audit + rate-limits +
API-key + tenant isolation) into the very first release. Once that floor exists, writes are cheap-to-add
one-tier-at-a-time (W1 → W2). The opposite ordering — writes first — would ship the highest-blast-radius surface
on the shakiest floor. Ruled out.

### Rationale — auth-vendor independence (Point C)

Anthropos runs on Clerk today; if the SSO/identity vendor ever changes (customer demand, pricing, compliance),
we do not want to rewrite every customer-API handler. The internal `Principal` DTO + `IdentityProvider` adapter
port contain the change to one package. **The lint rule "no `clerk.*` import in `internal/customerapi/`" is the
mechanism** — a code-review reminder is not enough; the machine enforces it. This is P3, load-bearing.

### Rationale — one contract, two shells (Point D)

An MCP server that reimplements the REST handlers doubles the audit + rate-limit + isolation surface area, and
guarantees they will drift. **Both shells delegate to the same resource-catalog handler; auth + audit + rate-
limits fire in the shared layer, once.** Adding a resource lights up both shells atomically. This is what makes
MCP a fast-follow (R2) rather than a parallel program.

### Rationale — the MVP shape (Point E)

The FIRST-USABLE flag on UC1–UC4 is deliberate: **roster, path progress, verified skill, org chart** are the
four reads that cover the vast majority of "please export X for us" support tickets today (G4). Landing these on
the audit + rate-limit + API-key floor means the customer's *first* API call is on the *permanent* foundation —
no v0-vs-v1 migration story, no "we'll re-do the auth later." That the MCP shell can then re-project the same
catalog in R2 is the payoff of the shared-layer investment (D).

### Rationale — the M303 iterative shape (Point S)

M303 (REST reads gateway) is the only R1 milestone that is `iterative`, not `section`. The reason: **each
resource ships closed on its own end-to-end gate** — OpenAPI entry + contract test + cross-tenant isolation
test + audit row + rate-limit fire. The *set* of resources to ship is declarable (UC1–UC7's reads); getting each
resource **through** its gate against the real internal RPC + real Postgres + real Redis budget is where the
uncertainty lives (missing-scope discovery, isolation edge cases, rate-limit tuning). One iter = one resource;
the exit gate is UC1–UC7 all green with 0 cross-tenant leakage over N runs. The other three milestones are
`section` (they have known, enumerable checklists — the adapter port, the key primitive, the docs surface).

### Rationale — the mutation gap posture (Point T)

Appendix A is the honest inventory: some W1/W2 actions map cleanly to existing internal RPCs; some are
partial-fits that need a small platform-repo edit; one (`audit.webhook.subscribe`) is a genuine missing platform
capability. **The customer API never invents a mutation the platform doesn't own** — a `missing` entry is
escalated to the platform roadmap (its own release beat), and the customer-API endpoint stays behind an
`unimplemented` state until the mutation lands. This mirrors Playthroughs' `unimplementable-without-platform-
edit` state — a zero-invention escape valve.

