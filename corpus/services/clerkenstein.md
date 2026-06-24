# Clerkenstein

**Status:** v0.3 (v1.0 "body double" · M1 + M2 + M2b + M2c `@clerk/express` + M3 deploy/injection; v1.9 "storytelling" M37 multi-identity seat-switch; v1.10 "method acting" M39 roster org-name threading) · **Last updated:** 2026-06-24
**Repo:** `stack-demo/rosetta-extensions/clerkenstein` (gitignored demo scratchpad, its own git) · **Measured by:** the
[alignment framework](../architecture/alignment_testing.md)

> **This is a pointer.** The full, self-contained documentation now lives **in the clerkenstein repo's own
> knowledge base** (added in M2b): start at `stack-demo/rosetta-extensions/clerkenstein/knowledge/kb-index.md`. This page
> keeps only the platform-side orientation + the cross-links a rosetta reader needs — it deliberately does
> **not** duplicate the repo's KB.

> **One monorepo, two clone roles.** `rosetta-extensions` is ONE private monorepo with sections
> (`clerkenstein`, `demo-stack`, `stack-injection`, `stack-core`, `stack-seeding`, `alignment`). It is
> authored / built / tested / aligned in the **authoring copy** at `.agentspace/rosetta-extensions/`, then
> **tagged**, and consumed by each stack via that stack's own **pinned-tag** clone
> (`stack-demo/rosetta-extensions @ <tag>`). So the KB above is **read from the authoring copy**; a running
> stack uses its pinned consumption copy.

## Role (platform-side orientation)

Clerkenstein is a **drop-in mock of the Clerk library** the platform uses — the *same interface*, with all
security and sync **disarmed**. It lets **demo** environments create users / orgs / admins and log in/out
with no Clerk friction (one universal credential, no live API, no webhooks, no rate limits), while platform
repos keep "thinking" they use Clerk with **zero source changes**.

It is the **first mirror produced by the M0 alignment process** (not a hand-built mock): its fidelity is
*measured* as a 0–100% alignment score against a Clerk **Alignment DNA**, driven to the gate — **100%
critical / 100% overall** on **all five** measured surfaces: the Go surface (22/22 genes, `clerk-sdk-go/v2
@ v2.6.0`, M1), the JS/FAPI surface (9/9 genes, `@clerk/clerk-js` v5 / `@clerk/nextjs` v6, M2), the
**multi-identity seat-switch** surface (9/9 genes, `clerk-multi-1` — the v1.9 M37 registry + active-seat
selection, so a demo can present as any seeded hero; the multi-session FAPI semantics real clerk-js exhibits
with `single_session_mode=false`), the **`@clerk/express`** Node-backend surface (9/9 genes, `@clerk/express`
^1.3.47, M2c — RS256/JWKS, the genuine SDK *satisfied*, not reimplemented), and the **deployment/injection**
surface (7/7 genes, `clerk-deploy-1` — the disarmed `colony/authn/provider/clerk` drop-in compiles against
the platform's real `colony @ v0.34.3` and satisfies its contract; added after **M3** showed *behavioural*
alignment ≠ *deployability* — see [`alignment_testing.md`](../architecture/alignment_testing.md#what-alignment-proves--and-what-it-doesnt-the-m3-lesson)).
The DNAs + mirror + goldens + runners live in the clerkenstein repo; the `/align-dna` + `/align-run`
skills + the [`alignment_testing.md`](../architecture/alignment_testing.md) doc live in rosetta, while the
`alignctl` harness is the `rosetta-extensions/alignment/` section (a sibling of `clerkenstein/`).

> **What "100%" means (and doesn't).** The score measures the mirror as *indistinguishable from the source
> goldens*. Those goldens are **hand-authored / hybrid** (decision M1-D1) — the reference behavior derived
> from the real libraries' documented + observed semantics (and, for `@clerk/express`, confirmed by driving
> the *genuine* SDK), **not** captured from a live, network-connected real-Clerk tenant. So 100% means "the
> mirror reproduces the behavior we encoded as the reference," not "diffed byte-for-byte against a running
> Clerk instance." Re-capturing goldens against a live source on a Clerk version bump is the M1b drift loop's
> job. This is the right bar for a *demo* mock; it is not a conformance certificate against production Clerk.

## Repo structure (library-named, since M2b)

The repo is organised **one dir per mocked dependency** (M2b reorg, decision M2b-D2):

| Dir | Mocks | What it is |
|---|---|---|
| `authn/` | `colony/authn` | the provider twin — **verifies** session JWTs (offline) |
| `clerk-backend/` | `clerk-sdk-go/v2` | fake Backend API + the in-memory org store, merged |
| `clerk-frontend/` | `@clerk/clerk-js` + `@clerk/nextjs` | fake Frontend API + publishable-key codec — **mints** JWTs |
| `clerk-webhook/` | `svix` | the signed-webhook injector |
| `shared/` | — | universal-key HS256 JWT (the mint side + verify side agree here) |
| `deploy/` | `colony/authn/provider/clerk` | the disarmed provider drop-in — **deployable** into a vendored colony fork (compiles against real `colony @ v0.34.3`) |
| `cmd/` | — | standalone binaries: `mintpk` (authoritative publishable-key minter) · `fake-fapi` / `fake-bapi` (standalone fake servers for demos; `fake-fapi` loads `FAKE_FAPI_ROSTER` for M37 multi-identity) |
| `alignment/` | — | the measurement harness: `cmd/{clerkrun,jsfapirun,multirun,expressrun,deployrun}` + `dna/` (five) + `golden{,-js,-multi,-express,-deploy}/` + `scripts/` |

The browser-login → backend-verify coherence chain runs through `shared`: `clerk-frontend` mints the
HS256 universal-key JWT, `authn` verifies that exact token — pinned by the JS DNA's
`SessionToken/decoded-identity` gene (operator `exact`).

### Multi-identity

**(v1.9 M37)** — `clerk-frontend` now holds a **users/orgs registry** (replacing the single
`DefaultDemoUser`) + an **active-seat selection** so a demo can **switch the active browser identity** among
the seeded heroes/orgs (the M35 stories roster) — the seat-switch the presenter cockpit's "login as" needs.
Selection is **server-authoritative** (the FAPI holds the active key, so the client view, `/v1/me`, the
token mint, and the handshake cookies all resolve the same hero): `?__clerk_identity=<key>` on the handshake
(the cockpit's [Login as] deep-link) + the `/v1/demo/{identities,select}` control plane. The single-identity
path is byte-identical (a one-member registry). Measured by the `clerk-multi-1` DNA (`alignment/cmd/multirun`,
9 genes, 100%/100%) — a *new measured surface* that holds while the existing four stay green.

**Roster org-name threading (v1.10 M39).** The roster now carries each hero's **story org name + slug**, so a
logged-in hero's **top bar reads her real company** (e.g. "Cervato Systems") instead of the hardcoded
"Clerkenstein Demo Org". The thread is a **paired change** kept in lockstep by the roster's
`DisallowUnknownFields` decoder — the producer (`stack-seeding/seeders/roster.go`) and the consumer
(`clerk-frontend`) add the same two `org_name`/`org_slug` snake_case fields in one change, and the rext repo is
re-tagged as a whole:

- **Producer** — `RosterIdentity` (roster.go) gains `org_name`/`org_slug`, filled in `BuildRoster` from
  `st.Org.Name` + the single-sourced `orgSlugFor` (the **same** slug rule `OrgSeeder` writes to
  `public.organizations.slug`, so the roster-carried org and the seeded org can never disagree).
- **Consumer** — `RosterEntry` (`clerk-frontend/registry.go`) gains the matching `org_name`/`org_slug` and
  threads them through `toDemoUser` into `DemoUser` (`resources.go`); `DemoUser.orgMemberships()` renders them
  on the FAPI org resource (`/v1/me` → the SDK's active-org → the top bar).
- **No-roster default fallback** — an empty `OrgName`/`OrgSlug` (the `DefaultDemoUser`, or any roster that omits
  the fields) falls back to the `orgNameDefault`/`orgSlugDefault` constants (`"Clerkenstein Demo Org"` /
  `"clerkenstein-demo"`), so the single-identity path is **byte-identical** and a pre-M39 roster still loads
  (the decoder rejects *unknown* fields, not *missing* ones — forward-compatible).

Alignment held: the **multi-identity** (`clerk-multi-1`) + **JS/FAPI** (`clerk-js-5`) surfaces stay **9/9,
100%/100%** (the `DefaultDemoUser` goldens are unchanged — they take the default-name fallback).

**`@clerk/express` (M2c) added no new dir** — it's a *consumer* (a Node backend verifier we satisfy), so
its support is **additive**: an RS256 path (RS256 minting in `shared/` + a real JWKS from `clerk-frontend/`
+ read endpoints in `clerk-backend/`), measured by the `alignment/cmd/expressrun` runner driving the
**genuine `@clerk/backend`** — the same "verify against the real library" discipline `clerk-webhook/` uses
with `svix`. `@clerk/express` verifies RS256-via-JWKS and rejects HS256, so the RS256 path is additive (the
HS256 seams + M1/M2 gates stay green).

**The deployment/injection surface (M3) *did* add `deploy/` + `cmd/`.** Unlike the `authn/` twin (which
mocks the standalone `colony/authn` interface), the platform actually consumes `colony/authn/provider/clerk`
*inside* the `colony` module. So the **deployable** drop-in lives in `deploy/colony-authn/`: the disarmed
provider — same package, same `Clerk` type, same `NewProvider(apiKey)` signature — compiled against the
platform's **real** `colony @ v0.34.3` so an injected demo app accepts Clerkenstein-minted tokens with zero
source changes. It is **identity-agnostic** (straight-through claim mapping — it extracts whatever the token
carries, not a hard-coded user). Its contract is checked at *compile time* and scored by the
`alignment/cmd/deployrun` runner (the `clerk-deploy-1` DNA). `cmd/` ships the supporting standalone tools:
`mintpk` (the authoritative publishable-key minter) and `fake-fapi` / `fake-bapi` (standalone fake servers
for demos).

## Read next (in the clerkenstein repo)

- **`knowledge/kb-index.md`** — the KB entry point (scope, architecture, alignment, injection, coverage).
- **`knowledge/scope.md`** — what it is/isn't + the disarmed-by-design properties.
- **`knowledge/architecture.md`** — the dir layout, public API surface, and the universal-key JWT flow.
- **`knowledge/injection.md`** — the four per-library injection recipes (each labelled built+gated /
  spike-proven / recipe-only) for disarming the platform's Clerk with no platform-code change.
- **`knowledge/alignment.md`** — how fidelity is measured against a pinned Clerk version + the **drift
  runbook** (M1b: `gate.sh` / `drift-check.sh` exit-code contract / weekly CI; re-`/align-dna` +
  re-`/align-run` on a Clerk bump). `ALIGN_DIR` default is `../../alignment` (the sibling section; scripts
  live at `alignment/scripts/`).
- **`knowledge/coverage-index.md`** — per-package test coverage + known gaps.
- Per-library `README.md` in each dir for the code-level entry point.

## See also (rosetta)
- [Alignment Testing](../architecture/alignment_testing.md) — the framework that measures this mirror.
- [Clerk integration](clerk-integration.md) — the real Clerk surface Clerkenstein mirrors.
- [Frontend architecture](../architecture/frontend_architecture.md) · [next-web-app](next-web-app.md) — the
  `@clerk/nextjs` consumers the `clerk-frontend/` server stands in for.
- [Webhook setup](../ops/webhook_setup.md) — the real Clerk webhook path the `clerk-webhook/` injector replays into.
