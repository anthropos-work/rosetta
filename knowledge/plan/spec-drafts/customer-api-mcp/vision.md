# Customer API + MCP — Vision & Long-Horizon Posture

> **Status:** Draft · spec-draft · 2026-07-01
> Companion to [`spec.md`](spec.md). This is the *why-it-matters* + *where-it-goes* piece; the *what-and-how*
> lives in [`spec.md`](spec.md).

## 1. Why this pillar exists

Anthropos today is reachable **only through its own UIs**. A customer developer who wants to script an
integration, or an AI agent that wants to answer a question about a customer's workforce, has no sanctioned
path. The gap is closed today by support tickets ("please export X for us") and one-off internal SQL — neither
scales, both leak the tenant-isolation boundary through humans.

This pillar closes that gap **for real**: a versioned, tenant-isolated, principal-audited programmatic contract
that a customer developer or a customer's AI agent can trust to do a real job — replace the export ticket,
answer the roster question, mirror the org chart, launch the simulation.

## 2. The strategic bet

Three concentric bets, in order of certainty:

1. **Table stakes** — enterprise HR-tech buyers **expect** a public API. Not having one is a losing-deal reason
   today. R1 is the minimum stake at the table.
2. **Ecosystem** — HR-tech partners (HRIS mirrors, LMS bridges, SIEM sinks) will build against a stable public
   contract if one exists. Every partner integration reduces the internal integration burden and grows the
   platform's reach. R4/R5 unlock partner-driven writes.
3. **AI-native** — the biggest bet: Anthropos becomes **an MCP source**. Not "an API that has an MCP wrapper" —
   a platform designed for AI agents to consume as a first-class client. That's the R2 posture (P8): every R1
   resource is MCP-tool-shaped from the start. If AI agents become the dominant consumer of enterprise data
   surfaces, being MCP-native at the point they do matters more than the API elegance.

## 3. What we deliberately don't chase

- **Not a wire-level RPC replacement.** The internal Connect-RPC surface stays as it is. The customer API is a
  curated façade — a smaller, safer, better-documented surface than the internal RPC.
- **Not a general-purpose data warehouse export.** Customers who want raw analytics can pipe the API into their
  warehouse. Anthropos does not become a warehouse.
- **Not a partner marketplace.** A directory of third-party apps built on the API is a future business decision
  ([`next-release.md`](next-release.md)) — the pillar just provides the contract.
- **Not a billing surface.** Metering for rate-limits is in R1; metering for invoicing is a separate program.

## 4. The auth-vendor-swap horizon (why P3 matters far beyond R1)

Anthropos runs on Clerk today. Enterprise customers regularly ask for SSO / SAML / custom IdP — sometimes the
sales cycle *depends* on it. If the identity vendor ever changes, we do not want to rewrite every customer-API
handler. The **`Principal` DTO + `IdentityProvider` adapter port** (§5.4 in [`spec.md`](spec.md)) is a small
architectural investment now that pays back many-fold when the swap becomes real.

This is the same lesson the platform learned with Directus (content) and could have learned earlier with LiveKit
(voice): **contain the vendor at a port; consume via a DTO; never let the vendor leak.** The customer-API pillar
is the first surface where this rule is *machine-enforced* (the "no `clerk.*` above the adapter" lint), because
it is the surface most exposed to a future swap.

## 5. The multi-release arc

The [`spec.md`](spec.md) roadmap (§6.1) is a **six-release program** (R1..R6). The arc:

- **R1 "open house"** — reads + audit + rate-limits + API-key self-serve + docs lite. **The MVP.** A customer
  replaces one export ticket with one curl call on day one.
- **R2** — the MCP shell over the R1 catalog. Agents can consume Anthropos as an MCP source.
- **R3** — query enrichments (GraphQL projection, streaming, aggregations). The API becomes richer without
  changing shape.
- **R4** — W1 writes (safe writes cluster). Roster + assignment + org structure. The first real "our systems
  drive Anthropos" release.
- **R5** — W2 writes + webhooks. Verified-skill emissions, session launches, webhook subscriptions. The full
  read+write ecosystem contract.
- **R6** — GA hardening (SDK code-gen, hosted MCP, SLA, status page, on-call). The pillar as a product.

Each release inherits the R1 floor (Principal + audit + rate-limits + API-key + tenant isolation) — the
architectural investment amortizes across the whole arc. **This is why R1 spends its scope on the floor + a
narrow read surface, not on breadth.**

## 6. Success signals

Per release:

- **R1** — N customers hitting `GET /v1/organizations/{id}/members` in the first month replaces N support
  tickets; the audit log surface is exported by at least one compliance officer; zero cross-tenant leakage
  incidents.
- **R2** — at least one AI agent (Claude, a partner LLM) demonstrably answers customer roster + progress
  questions via the Anthropos MCP server.
- **R4** — at least one customer scripts new-hire onboarding through the W1 writes cluster.
- **R5** — at least one HR-tech partner ships a bidirectional integration (reads + writes).
- **R6** — SDK downloads, hosted-MCP tenant count, SLA hit rate.

The **top-line business signal** is deals won because we have a public API + MCP source, and deals not lost
because we don't. That signal begins at R1 and compounds through R6.

## 7. Where this connects to the rest of Anthropos

- **Playthroughs (v2.0)** — R1's UC1–UC4 customer flows become Playthroughs on the v2.0 foundation once R1
  lands. The API being deterministic + audited makes it *easier* to Playthrough than the UI is.
- **Path migration (R2 spec-draft)** — orthogonal. The customer API exposes skill-path reads regardless of which
  engine is behind them.
- **Sentinel (authz)** — unchanged. The customer API's `Principal` carries scopes; the handler asks Sentinel
  per-resource.
- **Clerk (authn)** — contained. One of two adapters (the other being `ApiKeyIdentityProvider`).
- **Studio / Academy authoring** — orthogonal. Authoring is Studio's surface; the customer API is consumption +
  operations.

