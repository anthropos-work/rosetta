# Clerk Integration

> **Single source of truth for how the Anthropos platform uses Clerk.** Other docs
> mention Clerk in their own context (per-service usage, ops/webhook setup); this
> page is the cross-cutting picture — what Clerk is used for, how, which repos depend
> on it, and how each one integrates.

## High-Level Summary (For PMs & Non-Engineers)

**Clerk** ([clerk.com](https://clerk.com)) is the platform's external **identity provider** (SaaS). It owns three things: **authentication** (who the user is — login, sessions, JWTs), **user identity** (profile data), and **organizations** (company workspaces, their members, and each member's role). Multi-tenancy starts here: a user's Clerk **organization** is what scopes their data in the platform.

Clerk is **not** the platform's permission engine. For the backend, *deciding what a user may do* is **[Sentinel](./sentinel.md)**'s job (Casbin RBAC/ABAC). Clerk feeds Sentinel the raw material (who's in which org, with what role) but doesn't make backend allow/deny calls itself. The **frontend and standalone apps**, however, **do** use Clerk's `org:admin` role and org-membership directly to gate features (enterprise/admin tooling, content access) — there's no Sentinel call on that path.

## What Clerk Is Used For

| Concern | Used? | Detail |
|---------|:-----:|--------|
| **Authentication** | ✅ | Issues session JWTs; verified platform-wide against Clerk's JWKS. The "are you signed in?" gate. |
| **Identity / profile** | ✅ | User id (internal `eid` + Clerk subject), email, first/last name. |
| **Organizations & membership** | ✅ | Clerk Organizations are the source of truth for which org a user belongs to. |
| **Org roles** | ✅ | `org_role` rides in the JWT. **Server-side it is the bare `admin` / `basic_member`** (what `AuthRole()` and Next.js `auth()` server routes see — e.g. `metabase/route.ts` checks `orgRole === 'admin'`); the **client `useAuth().orgRole`** is prefixed (`org:admin`). Server checks should match the bare form or accept both. |
| **Authorization (backend)** | ⚠️ indirect | Backend permission decisions are made by **Sentinel**, not Clerk. Clerk org roles are *synced into* Sentinel at webhook time, then Sentinel decides. |
| **Authorization (frontend/standalone)** | ✅ direct | next-web-app, studio-desk, and ant-academy read Clerk's `org:admin` / membership **directly** to allow/deny features (no Sentinel). |

> **One-line answer to "auth only, or authz too?"** — Clerk handles **authentication + identity + organization/role management** everywhere. It is **not** the backend authorization engine (Sentinel is, fed by Clerk via sync), but the **frontend/standalone apps authorize directly off Clerk org roles/membership**.

## How It Works (Deep Dive)

### 1. Authentication — the `authn` library
All Go services authenticate through the shared **`authn`** library (now shipped inside **colony** as `colony/authn`; see [Shared Libraries → authn](../architecture/shared_libraries.md#authn)). Its Clerk provider:
- Verifies the session JWT against Clerk's **JWKS** (`clerk-sdk-go/v2` `jwt.Verify` + `jwks.Client`, 1-minute leeway), then `jwt.Decode`s the claims.
- The HTTP/Echo middleware returns **401** on an invalid/missing token and otherwise injects the authenticated `User` into request context. That's authentication/route-protection — no role check.

### 2. Identity & org claims (custom session token)
To avoid Clerk API round-trips, the platform puts custom claims on the Clerk session token; `authn` reads them:

| Claim | Maps to |
|-------|---------|
| `eid` | `User.ID()` — internal Anthropos UUID (Clerk `external_id`) |
| (JWT subject) | `User.AuthID()` — Clerk user id |
| `email`, `firstname`, `lastname` | profile (lazy `user.Client.Get` fallback if absent) |
| `org` (public-metadata map, its `eid`) | `Organization.ID()` — internal org UUID |
| `org_id` | `Organization.AuthID()` — Clerk org id |
| `org_role` | `Organization.AuthRole()` — returned **verbatim**. The backend session token carries the **bare** form (`admin` / `basic_member`); only the client-side `useAuth().orgRole` is prefixed (`org:admin`). Match the bare form server-side, or accept both. |

`AuthRole()` is **exposed as a getter and never enforced inside `authn`** — consumers decide what to do with it.

### 3. Authorization — who actually decides
| Layer | Clerk's role here | Decision maker |
|-------|-------------------|----------------|
| `colony/authn` | verify JWT, surface claims | — (authentication only) |
| **Backend `app`** + jobsimulation, cms, skiller, skillpath | authenticate; supply `org_id` for tenant scoping | **[Sentinel](./sentinel.md)** via Connect-RPC (`OrgCheckUserPermission`, `CheckFeature`, …). `AuthRole()` has **zero** call sites on the allow/deny path. |
| **Sentinel** | not used at all (no Clerk/authn import) | Sentinel's own Casbin policy store |
| storage, messenger | — | no auth |
| **next-web-app / studio-desk / ant-academy** | authenticate **and** authorize | **local app code** reading Clerk `org:admin` / membership |

### 4. Org/role sync — how Clerk reaches Sentinel
The **backend `app`** receives **svix-verified Clerk webhooks** (`user.*`, `organization.*`, `organizationInvitation.*`, `organizationMembership.*`). On membership change it:
1. upserts the local user/org/membership, and
2. **translates the Clerk role** (`admin` / `basic_member`) and **pushes it into Sentinel** (`OrgAddUserToRole` / `OrgReplaceUserRole` / `OrgRemoveUserFromRole`).

So Clerk org roles become Sentinel roles at **sync time** — that's the only way Clerk influences backend authorization. The backend also writes membership/user metadata **back** to Clerk via `clerk-sdk-go/v2` (`organizationmembership`, `user`). See [Webhook Setup](../ops/webhook_setup.md) for local tunnel configuration.

## Dependent Repos & How They Integrate

| Repo | Clerk SDK / lib | Integration |
|------|-----------------|-------------|
| **colony** (`/authn`) | `clerk-sdk-go/v2` (`jwt`, `jwks`, `user`) | The auth core: JWT/JWKS verification + claim extraction. Imported by every authenticated Go service. |
| **app** (backend) | `clerk-sdk-go/v2` (`organization`, `organizationmembership`, `organizationinvitation`, `user`, `signintoken`) + `svix-webhooks/go` + `colony/authn` | Authn via colony; svix webhook → user/org/membership sync → seed Sentinel roles; org/role writes back to Clerk. |
| **jobsimulation, cms, skiller, skillpath** | `colony/authn` (`clerk-sdk-go/v2` transitively) | Authenticate only; authorization → Sentinel. |
| **storage, messenger** | — | No Clerk / no auth. |
| **sentinel** | — (none) | Does **not** import Clerk/authn; pure Casbin authorization. |
| **next-web-app** (`apps/web`, `apps/hiring`, `apps/integration`) | `@clerk/nextjs` `^6.39` (+ `@clerk/localizations`, `@clerk/types`) | All three Next.js apps use the same `@clerk/nextjs` setup — `clerkMiddleware` (`auth.protect()`) for **authentication only**. In `apps/web`: `/enterprise` is gated **client-side** (`EnterpriseWrapper` redirects non-admins via `useAuth().orgRole`), and `/api/metabase` **server-side** via `auth()` on the bare `orgRole === 'admin'` **AND** a specific hardcoded org id. |
| **next-web-app** (mobile) | `@clerk/clerk-expo` `~2.6` | Mobile sign-in/session (Expo). |
| **studio-desk** | `@clerk/clerk-js` `^5.52` (frontend) + `@clerk/express` `^1.3` (backend) | Client + server auth; **all admin tooling gated on `org:admin`** (`checkEnterpriseAndAdmin` server-side, `UserService.isAdmin()` client-side). |
| **ant-academy** | `@clerk/nextjs` `^7.2` | `clerkMiddleware` in `proxy.js`; **requires ≥1 org membership** (`REQUIRE_ORGANIZATION_MEMBERSHIP` → `/no-organization`). |

## Configuration (Keys)

Clerk credentials live in `platform/.env` (backend) and each app's own env. The shapes:

```bash
# Backend / server
CLERK_SECRET_KEY=sk_...
CLERK_WEBHOOK_SECRET=whsec_...        # svix signature verification (app webhooks)

# Frontend (publishable, framework-prefixed)
NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=pk_...   # next-web-app, ant-academy
VITE_CLERK_PUBLISHABLE_KEY=pk_...          # studio-desk
EXPO_PUBLIC_CLERK_PUBLISHABLE_KEY=pk_...   # mobile
```

Use **separate Clerk applications** for dev vs production, and never commit keys. The full sign-in/up URL variables and the local webhook tunnel are covered in the ops guides below.

## Related Documentation
- [External Services → Clerk](../architecture/external_services.md#clerk-authentication-service) — Clerk in the external-services catalog (overview)
- [Sentinel](./sentinel.md) — the authorization engine fed by Clerk's synced roles
- [Shared Libraries → authn](../architecture/shared_libraries.md#authn) — the Clerk JWT library internals
- [Webhook Setup](../ops/webhook_setup.md) — local Clerk webhook tunnel
- [Staging Clerk](../ops/staging-clerk.md) — dev Clerk app, shared login, the `clerk-fetch-fix.js` patch
- Per-service usage: [backend](./backend.md), [studio-desk](./studio-desk.md), [ant-academy](./ant-academy.md), [frontend](../architecture/frontend_architecture.md)
