# Observability — the tier this corpus documented nowhere

**Until 2026-08-07 this corpus had no observability documentation at all.**
`git grep -i grafana -- corpus/ CLAUDE.md` returned **0 files**; the only `prometheus` hits anywhere
were a port-9100 collision warning against `node_exporter` in the setup guides. Meanwhile the platform
has had live outside-in monitoring since **2026-07-07**, and the answer to *"is production up?"* has been
computed somewhere this corpus never named.

It lives in **`anthropos-work/ant-observability`** (HEAD `b49eb7af`, 2026-08-05). That repo holds **two
different things**, and conflating them is the first mistake to avoid.

## 1. What the platform actually emits — measure this before believing any dashboard

> **Every `app` anchor in the table below is pinned to `app` `ad9f3c49`, and to that ref ALONE.** Not
> because it is the newest — it is not; `origin/main` was `3eaadae6` on 2026-08-07 — but because
> **a block that names two refs is `ambiguous` to the citation resolver, which then falls back and grades
> every anchor in the block against a file the block did not mean** (M257x run-53; the `storage` row of
> [`platform-migration-status.md`](../architecture/platform-migration-status.md) omits its older line
> numbers for the same reason). This table cost one RED to learn it twice: an earlier revision cited
> `main.go:278` @ `3eaadae6` **and** named `ad9f3c49` in the same cell, and `anchor_construct_guard`
> resolved the anchor onto a closing `)`. **The two anchors here each moved by exactly one line between
> the refs** (`:273`/`:277` here, `:274`/`:278` at `3eaadae6`) — the construct is identical at both, and
> the drift is a one-line insert upstream, not a change to the thing being cited.

| Signal | Status | Evidence |
|---|---|---|
| **Metrics** | **NONE.** The platform exposes no `/metrics` endpoint and registers no Prometheus collectors | `promhttp` / `/metrics` / `prometheus` grep over `app` and `colony` `.go` → empty. `prometheus/client_golang` is `// indirect` at `app/go.mod:200`. Independently corroborated by the monitoring repo's own backlog: `product-monitoring/prometheus-scrape.yml:265` — *"profile import duration histogram — **needs `/metrics` in colony** (phase 2b)"* |
| **Traces** | **Sentry-protocol only, 15 % sample. No OpenTelemetry** | `app/main.go:278` `colony.WithLoggingTracing(0.15, 0.15)` → `colony/logging.go:54-55`. `go.opentelemetry.io/otel` is `// indirect` at `app/go.mod:232` and imported by **zero** source files in either repo |
| **Errors** | Sentry SDK, **production-only** | `app/main.go:273` `SentryDSN: os.Getenv("SENTRY_DSN")`; `colony/logging.go:47` early-returns unless `IsProduction()` or `FORCE_SENTRY` |
| **Frontend errors** | `@sentry/nextjs` ^10.57.0 across web / hiring / integration / maintenance; `@sentry/react-native` on mobile | `next-web-app/apps/*/package.json`; `apps/web/sentry.{edge,server}.config.ts`, `instrumentation-client.ts` |
| **Health** | `/_meta`, `/_meta/ready`, `/_asynq/` | `app/internal/meta/server.go:23-25` |

> **The Sentry reconciliation.** `CLAUDE.md` has always said colony gives you *"logging+Sentry"*. What it
> never said is that **the DSN points at a self-hosted GlitchTip**, not at sentry.io —
> `ant-observability/glitchtip/README.md:1-6`: *"speaks the **Sentry wire protocol**… the only change is
> pointing the DSN at this instance instead of Sentry's cloud."* This corpus half-knew it
> ([`studio-desk.md:301`](../services/studio-desk.md) mentions a *"self-hosted GlitchTip Sentry
> endpoint"*), but no architecture doc said the platform's **error tier is GlitchTip on a Proxmox VM**.
>
> **Not measurable from any clone, and therefore not asserted:** which receiver is live.
> `next-web-app/packages/core-js/src/monitoring/tunnel.ts:1-2` documents itself as a **Better Stack**
> tunnel and is host-agnostic — it forwards to `<trustedDsn.host>/api/<projectId>/envelope/`. **Which
> host that is, is a DSN environment value.** [`next-web-app.md:76`](../services/next-web-app.md)'s
> *"tunnels Sentry/Better Stack events"* is correct as far as it goes, and that is as far as it can go
> from here.

**The consequence for anyone debugging:** there is **no metrics pipeline out of the platform**. Anything
that looks like a platform metric is a **synthetic probe result**, produced outside the platform by
§ 3 below.

## 2. Half one — Claude Code usage telemetry (not the platform)

The repo self-identifies as a **singularity** asset: `CLAUDE.md:1` *"Singularity Observability"*, `:3-4`
*"The repo is named for our internal 'singularity' tooling line; what it **observes** is Claude Code
usage across the team."*

OTel Collector → Prometheus + Loki → Grafana (`stack/docker-compose.yml`:
`otel/opentelemetry-collector-contrib:0.153.0`, `prom/prometheus:v3.1.0`, `grafana/loki:3.3.2`,
`grafana/grafana:11.4.0`). **The OTLP source is every engineer's Claude Code CLI, not the platform** —
`stack/otel-collector/config.yaml:1-3`; engineers opt in per `docs/onboarding.md:20-30`.

**Do not read this half as platform telemetry.** It is the reason the repo's name misleads.

## 3. Half two — `product-monitoring/`: the platform tier, live since 2026-07-07

**Every probe asserts on BODY CONTENT, because status codes structurally cannot work here.** This is the
single most reusable finding in the repo (`docs/product-monitoring.md:13-30`):

- gqlgen answers errors as **`200 {"errors":…}`**;
- Next.js ISR serves a `FourOhFour` render at **200, cached for up to an hour**;
- the 2026-07-05 wedge kept the **load-balancer health check green** while the ent pool was exhausted.

**Targets** (`product-monitoring/prometheus-scrape.yml`): `anthropos.work/library` (`:45`) ·
`gql.anthropos.work/graphql/query` (`:71`, `:92`) · `api.anthropos.work/graphql/query` (`:123`) ·
VPC-internal `backend.internal.anthropos:8083/_meta/ready` (`:145`) and `:8080/graphql/query` (`:180`) ·
`sentinel.internal.anthropos:8080/` (`:215`) · and an **asynq redis-exporter reading production
ElastiCache directly** (`product-monitoring/docker-compose.yml:40-43` — O(1) lookups on 65 fixed keys,
security-group-gated, no AUTH).

**Deployment:** a dedicated Proxmox VM on the **odyssey** host (`infra/README.md:1-20`, 2 vCPU / 4 GB /
60 GB), Docker Compose, reachable **only over Tailscale** (`README.md:7-9`). No CI. GlitchTip's exposure
is deliberately split: the dashboard via `tailscale serve` (tailnet-only), **event ingest publicly via
`tailscale funnel`** behind an ingest-only Caddy filter allowing exactly
`^/api/[0-9]+/(envelope|store|security)/?$` and 404-ing everything else
(`glitchtip/ingest-proxy/Caddyfile:24-35`).

> ### ⚠️ A production READ path that no safety document enumerates
> The asynq exporter connects to **production ElastiCache**. It is bounded (65 fixed keys, O(1),
> SG-gated) and it is not Rosetta tooling, so it falsifies nothing in [`safety.md`](safety.md) — that
> contract is scoped to `rosetta-extensions` (`safety.md:3`, `:36-37`). It is named here for the same
> reason `sim-qa`'s write path is named in [`org-repos.md`](../architecture/org-repos.md) § 9: **the
> corpus's prod-access guarantees are tooling-scoped, and a reader will otherwise over-read them as
> org-wide.**

## 4. Two rules worth importing into this corpus's own probe design

Both come from `product-monitoring/`'s operational record and both apply directly to
[`verification.md`](verification.md):

1. **Never gate on a status code where the tier can return 200 for a failure.** Assert on body content.
   The three concrete shapes above are the proof, and the third — a green LB health check over an
   exhausted connection pool — is exactly the class `verification.md`'s cheap-win asserts exist to catch.
2. **A synthetic probe is production traffic.** Probe noise was roughly **90 %** of GlitchTip issue #92's
   volume (`docs/product-monitoring.md:134-163`). A monitor that pollutes the error stream it monitors
   degrades the signal it was added to protect.

## 5. External corroboration — an outside team's telemetry agrees with the fenced map

Worth recording because independent agreement is rare and cheap to lose:

- `product-monitoring/prometheus-scrape.yml:27-33` records the **router dropped 2026-07-31**, standardises
  on `gql.anthropos.work`, and instructs *"Do NOT add a router probe back"* — matching
  [`graphql-wundergraph.md`](../services/graphql-wundergraph.md).
- `:194-205` records **storage + messenger folded 2026-08-04**, with the measured **last responses at
  `11:16:59Z` and `11:31:59Z`** — matching the v9.0 banner in `CLAUDE.md`, from a completely different
  instrument.

## Related
- [`org-repos.md`](../architecture/org-repos.md) — the register this doc was created from
- [`verification.md`](verification.md) — Rosetta's own bring-up probe set
- [`safety.md`](safety.md) — the tooling safety contract, and its scope
- [`db-access.md`](db-access.md) — the documented production read path
