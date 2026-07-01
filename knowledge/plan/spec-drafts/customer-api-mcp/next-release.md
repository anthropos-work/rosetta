# Customer API + MCP — Next / Out-of-scope (parking lot)

> **Status:** Draft · spec-draft · 2026-07-01
> Companion to [`spec.md`](spec.md). Anything **out of scope for defining + first-building the Customer API +
> MCP pillar (R1)** lands here: deliberate **non-goals** and **later** work.

## Purpose

Two kinds of thing park here:

1. **Non-goals** — things that are explicitly **not** the customer-API pillar, so the pillar stays focused on
   *curated, principal-audited programmatic access* ([`spec.md`](spec.md) §8).
2. **Later** — pillar scope, but not R1: R2–R6 (or beyond).

When something is genuinely in-scope for a later release, park it here with its target release. When it becomes
a non-goal, park it here with the reason.

## How to use

- One row per item. Keep it short.
- **Target** = the intended release (`R2`, `R3`, …) for *later* items, or `non-goal` if out of scope forever.

## Parked items

| # | Item | Target | Notes |
|---|------|:------:|-------|
| 1 | **MCP shell over the R1 catalog** — the MCP server binary (UC8 FIRST-MCP, UC9) | R2 | Fast-follow. The R1 catalog is MCP-tool-shaped by design (P8); R2 lights up the shell on the same handlers. |
| 2 | **ND-JSON streaming for large list endpoints** | R3 | Cursor pagination suffices for the R1 UCs. |
| 3 | **GraphQL projection over the catalog** | R3 | Beyond REST — a customer-driven query surface. |
| 4 | **Server-side aggregations for common report shapes** | R3 | The "roll-up" shapes customers hand-write today. |
| 5 | **W1 writes cluster** (member CRUD, assignment CRUD, org-structure updates — UC10–UC12) | R4 | Safe writes on the R1 floor. Full audit + per-action rate limits + entitlement gates. |
| 6 | **W2 writes cluster** (verified-skill emit, session launch, webhooks — UC13) | R5 | Advanced writes; tighter entitlement gates. Includes the per-customer webhook subscription (UC13). |
| 7 | **SDK code-generation (TS + Python)** | R6 | OpenAPI supports it; the effort is real and warrants its own release beat. |
| 8 | **Hosted MCP variant** | R6 | R2 ships MCP as a customer-hosted binary; a hosted variant is R6. |
| 9 | **Formal SLA + status page + on-call rotation** | R6 | GA hardening. |
| 10 | **Cross-browser / device-matrix testing of the docs site** | R6 | Documentation-site polish, not customer-API polish. |
| 11 | **Service-to-service internal RPC replacement** | non-goal | The internal Connect-RPC surface stays as it is. The customer API is a curated façade, not a replacement. |
| 12 | **Billing / metering-for-revenue** | non-goal | Metering-for-limits (rate-limit budgets) is in R1; metering-for-invoicing is a separate program. |
| 13 | **Partner marketplace / directory** | non-goal | A directory of third-party apps built on the API is a future business decision, not a spec item. |
| 14 | **Public webhook broker / event bus** | non-goal | R5's webhook subscription is per-customer, not a public event bus. |
| 15 | **Anthropos-hosted "AI ask" endpoint** (natural-language over the catalog) | later — needs MCP-mature | Customers who want this today can run their own MCP client on their LLM of choice; a first-party hosted endpoint is a distinct product decision. |
| 16 | **Authoring APIs (Studio Desk, Academy)** | non-goal | The pillar exposes consumption + operations; authoring is Studio's surface. |
| 17 | **Admin cross-tenant reads for platform operators** | non-goal | A distinct, platform-internal surface. Explicitly *not* the customer API (P2). |

> Add new rows above. Promote an item into scope only by moving it into [`spec.md`](spec.md) / a spec section.

