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
| **Org roles** | ✅ | `org_role` rides in the JWT. **Server-side it is the bare `admin` / `basic_member` / `content_creator`** (what `AuthRole()` and Next.js `auth()` server routes see — e.g. `metabase/route.ts` checks `orgRole === 'admin'`); the **client `useAuth().orgRole`** is prefixed (`org:admin`). Server checks should match the bare form or accept both. |
| **Authorization (backend)** | ⚠️ indirect | Backend permission decisions are made by **Sentinel**, not Clerk. Clerk org roles are *synced into* Sentinel at webhook time, then Sentinel decides. |
| **Authorization (frontend/standalone)** | ✅ direct | next-web-app, studio-desk, and ant-academy read Clerk's `org:admin` / membership **directly** to allow/deny features (no Sentinel). |

> **One-line answer to "auth only, or authz too?"** — Clerk handles **authentication + identity + organization/role management** everywhere. It is **not** the backend authorization engine (Sentinel is, fed by Clerk via sync), but the **frontend/standalone apps authorize directly off Clerk org roles/membership**.

## Clerk Feature Surface — used vs. not used

Clerk's catalog is large; the platform uses a **focused subset**. Unused features are dashboard-toggled SaaS capabilities that cost nothing by being off — the real integration effort is in the design (custom claims, org-role sync), not the feature breadth.

**Used**
- **Authentication** — session JWT + JWKS verification (backend). Sign-in/up UI: prebuilt `<SignIn>` / `<SignUp>` + `<UserProfile>` (next-web-app), `<SignIn>` (ant-academy web); studio-desk delegates to the web app; ant-academy mobile uses a custom email+password form.
- **Session JWT as the API bearer** — `getToken()` is called with **no template** (the default token); the custom claims are baked into that default token via dashboard config.
- **Organizations + memberships** — active org, `setActive`, `publicMetadata.eid` → tenant id (core to multi-tenancy).
- **Org roles** — **three**, not two: `admin` / `basic_member` / **`content_creator`** (each also in a
  prefixed `org:` form) — coarse RBAC gating; synced into Sentinel. All six constants:
  `app/internal/clerk/orgclient/roles/roles.go:12-17` @ `ad9f3c498`. **This bullet said
  `admin` / `basic_member` until run 81**, while this same file named `content_creator` 120 lines lower
  (the studio-desk row) — the enumeration and the counter-example sat in one document.
- **Org invitations** — backend create/revoke (+ a hand-rolled bulk call); frontend list/accept.
- **Webhooks** (svix, 12 event types) — Clerk → Postgres + Sentinel sync.
- **Backend API** — org/membership/invitation CRUD, user-create CLI, `external_id` + metadata write-back, lookups.
- **User/org metadata** — `unsafeMetadata` (trial/stripe flags), `publicMetadata` (eid, isHiring, role).
- **Sign-in tokens** — **five** live minting sites, one product and four harnesses. The product one is app-native admin impersonation (chosen over Enterprise-tier Actor Tokens). **See the enumeration below — this bullet used to say "only", and it was false.**
- **Localization** — 8 locales (`@clerk/localizations`).

**Not used** (available but untouched)
MFA / TOTP / passkeys · SAML / Enterprise SSO · device & multi-session management · Clerk Billing (billing is Stripe) · custom org **permissions** / `has()` / `<Protect>` (the gating is by *role*, and there are **three** — see above) · Clerk-native impersonation / Actor Tokens · Waitlist / GoogleOneTap / Web3 · `@clerk/themes` · most prebuilt components (`<UserButton>`, `<OrganizationSwitcher>`, `<OrganizationProfile>`, … — org UI is custom).

> **⚠️ `OAuth / social` was in that "Not used" list until run 81, with the parenthetical *"mobile is
> email+password only"*, and BOTH halves are false.** `next-web-app`'s Expo app renders Google **and**
> Microsoft Clerk OAuth buttons — `apps/mobile/app/(public)/sign-in.tsx:70,75` and `sign-up.tsx:86,91`
> drive `apps/mobile/components/OAuthButton.tsx:29` (`useOAuth` → `startOAuthFlow`). Verified at
> `origin/main` `f97ba6599` **and** at `8297c684`, the ref this page pins — so it is not clone drift.
> Two consequences a reader was being denied: **two external IdP connections (Google, Microsoft) must be
> enabled on the Clerk instance**, and the identity attack surface includes a federated path. The mobile
> app is a **paused PoC** (see the SDK table below), so this is wired-and-shipped code rather than a
> live production login — state both halves, not either alone.

> **⚠️ Caveat:** ant-academy **mobile sign-in aborts on any second factor** (`ClerkSignInForm.tsx`), so enabling MFA in the Clerk dashboard would **break mobile login**. Treat MFA as an explicit on/off decision, not a silent dashboard toggle.

### Sign-in tokens — every minting site, enumerated

> ⚠️ **CORRECTED M257x iter-120 — this list said sign-in tokens are minted "only" for app-native admin
> impersonation. There are FIVE minting sites in the clone set, not one.** The word was a
> **security-surface understatement**: a reader auditing where a token that mints a full session for an
> arbitrary user can be created was sent to one site out of five. Booked and upheld at iter-116 and again
> at iter-119; repaired here by enumeration rather than by anchor. **Two of the five had never been named
> in any prior reading either** — the enumeration found them, the ledger did not.

A Clerk **sign-in token** (Backend API `POST /v1/sign_in_tokens`) is a one-shot ticket that a browser
exchanges via `signIn.create({ strategy: 'ticket' })` for a **genuine session as the named user** — no
password, no 2FA, no CAPTCHA. It is the strongest impersonation primitive in the integration, so the
complete list matters. Measured across the whole clone set with three independent instruments
(`/usr/bin/grep -rn`, per-clone `git grep`, and a Python `os.walk`, per rule 44 — all three agree):

| # | site | ref | what it is | gate |
|---|------|-----|-----------|------|
| 1 | `app/internal/admin/impersonation/manager.go:101` (`m.signInTokenCl.Create`) | `app` `ad9f3c49` | **The product feature** — app-native admin impersonation, the one this page always documented | A Sentinel permission check + a non-nil actor, then an audit row on **every** attempt, success or failure (`manager.go:1-4`). **The permission it checks is `ActionObjectTaxonomy` / `UserActionWrite`** (`internal/web/backend/graphql/graph/resolver_admin_audit.go:19-23`) — a **taxonomy-write** permission, not a dedicated impersonation one |
| 2 | `next-web-app/apps/web/src/app/api/dev/login-as/route.ts:79` | `next-web-app` `8297c684` | Dev "log in as a real Clerk user" harness → `/dev/accept` | `DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production'` (`apps/web/src/lib/devLogin.ts:28`); hard-404 otherwise |
| 3 | `next-web-app/e2e/auth.setup.ts:72` | `next-web-app` `8297c684` | **Playwright e2e auth setup** — mints a ticket rather than driving the password form | **No `NODE_ENV` gate** — it is a test-runner file, never in an app build, but it runs against a **real Clerk instance** with a real `CLERK_SECRET_KEY` |
| 4 | `studio-desk/src/routes/dev.ts:83` | `studio-desk` `41ee357` | Dev login harness → `dev-accept.html` | `DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production'` (`src/lib/devLogin.ts:33`) |
| 5 | `ant-academy/code/app/api/dev/login-as/route.js:78` | `ant-academy` `22df69dd` | Dev login harness → `/dev/accept` | `DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production'` (`code/src/lib/devLogin.js:29`); hard-404 otherwise |

**Sites 2–5 mint a session for an arbitrary user found by email** (`users.getUserList({ emailAddress })`
then `createSignInToken({ userId })`) — none of them is admin-scoped, and none of them is "admin
impersonation". Sites 2, 4 and 5 hard-404 outside local dev, so a **production-scoped** reading of the old
sentence was defensible; the sentence was not production-scoped, and site 3 is not `NODE_ENV`-gated at all.

**Site 3 is the one worth reading twice.** `e2e/auth.setup.ts:55-62` states its own reason: the e2e
account *"enforces 2FA (email_code as second factor); password signin returns `needs_second_factor` and
never produces a session"*, and *"Clerk treats it as fully authenticated and **skips both factors**."* A
sign-in token is therefore a **documented, deliberate second-factor bypass** against a real Clerk instance
— which is precisely the kind of fact an emphatic **only** on one unrelated site hides.

**A sixth consumer is an operator recipe, not checked-in code**, and this corpus already documents it in
three other places: the staging/CI `curl -s -X POST https://api.clerk.com/v1/sign_in_tokens` bypass for
Clerk's "new device" challenge — [`staging-clerk.md:58`](../ops/staging-clerk.md),
[`staging_from_dump.md:421`](../ops/staging_from_dump.md), and
[`staging-bringup.md:528`](../ops/staging-bringup.md) (*"Quirk #13"* — cite the **quirk number**, which is
stable). It mints against whichever instance the operator's key points at.

> **All three line pins moved again on 2026-08-07 (M257x iter-130) and this is the third recorded move.**
> `staging_from_dump.md` 384 → **421** and `staging-bringup.md` 468 → **528**, both because the routed
> residual repair grew the sections above them; the `staging-clerk.md:58` pin held. The paragraph already
> told its reader to cite the stable *quirk number* rather than the line — **so quote the anchor, not the
> offset**. **But the three anchors are NOT one anchor, and saying they were is the defect this note
> now records** (M257x iter-131 P19, repaired iter-133): **two** of the three are the literal
> `curl -s -X POST https://api.clerk.com/v1/sign_in_tokens` (`staging-clerk.md:58`,
> `staging_from_dump.md:421`); the third, `staging-bringup.md:528`, is a **prose bullet** —
> *"**Quirk #13 …** Bypass with `POST /v1/sign_in_tokens`"* — which carries neither `curl` nor the
> host. A reader who greps the literal finds two of three and concludes a pin rotted.
>
> **The robust re-derivation is the shared substring, and only that:** `grep -n sign_in_tokens
> corpus/ops/*.md` returns all three in one command. A line number is a convenience; **the grep is the
> contract, and the grep term must be the part every site actually shares.**
>
> **Recorded rather than quietly fixed, because of where it happened:** the over-claim was written into
> the sentence whose entire purpose was to make a citation robust against drift. *An anchor claim needs
> the same enumeration discipline as the claims it anchors.*

**The "chosen over Enterprise-tier Actor Tokens" half is correct** and stays: `manager.go:6-11` says so in
the source's own words (*"Actor Tokens would require Enterprise"*) @ `app` `ad9f3c49`.

**Corpus cross-references for the harness half** (this page previously mentioned it nowhere):
[`ant-academy.md:334`](./ant-academy.md) (the `DEV_LOGIN_ENABLED` public-route pair),
[`studio-desk.md:90`](./studio-desk.md) (*"Dev-only acceptance harness"*), and — added at **iter-121**, the
**pair-half iter-120 left open** — [`next-web-app.md` § the two minting sites in this repo](./next-web-app.md#the-two-clerk-sign-in-token-minting-sites-in-this-repo-added-m257x-iter-121). That page held **2 of the 5 sites, including the
only ungated one**, and said nothing about either; a repaired enumeration whose per-repo pages stay silent
is half a repair.

**The platform-side reading of sites 2–5 is FILED, not asserted here** (M257x iter-121):
`knowledge/plan/platform-defect-register.md` →
`PLATFORM-M257x-dev-login-routes-mint-a-full-session-for-any-email-behind-one-NODE_ENV-boolean`. This page
documents what the integration *does*; the register holds what a platform engineer would want to change.

## How It Works (Deep Dive)

### 1. Authentication — the `authn` library
Every Go service that authenticates a user does so through the shared **`authn`** library (now shipped inside **colony** as `colony/authn`; see [Shared Libraries → authn](../architecture/shared_libraries.md#authn)). Since the merges that is **`app` alone**: `repos.yml` now lists just **two** Go repos, `app` and `sentinel`, and only `app` imports `colony/authn` — measured, 129 files in `app` vs **0** in `sentinel`, consistent with the tables at §3 and § *Dependent Repos* below. (At `0dab54d` there were four, `storage` and `messenger` alongside them, and the same measurement gave 0 for each; `838d907` removed both entries.) Its Clerk provider:
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
| `org_role` | `Organization.AuthRole()` — returned **verbatim**. The backend session token carries the **bare** form (`admin` / `basic_member` / `content_creator`); only the client-side `useAuth().orgRole` is prefixed (`org:admin`). Match the bare form server-side, or accept both. |

`AuthRole()` is **exposed as a getter and never enforced inside `authn`** — consumers decide what to do with it.

### 3. Authorization — who actually decides
| Layer | Clerk's role here | Decision maker |
|-------|-------------------|----------------|
| `colony/authn` | verify JWT, surface claims | — (authentication only) |
| **Backend `app`** (incl. the folded jobsimulation + cms domains) | authenticate; supply `org_id` for tenant scoping | **[Sentinel](./sentinel.md)** via Connect-RPC (`OrgCheckUserPermission`, `CheckFeature`, …). `AuthRole()` has **zero** call sites on the allow/deny path. |
| **Sentinel** | not used at all (no Clerk/authn import) | Sentinel's own Casbin policy store |
| storage, messenger | — | no auth |
| **next-web-app / studio-desk / ant-academy** | authenticate **and** authorize | **local app code** reading Clerk `org:admin` / membership |

### 4. Org/role sync — how Clerk reaches Sentinel
The **backend `app`** receives **svix-verified Clerk webhooks** (`user.*`, `organization.*`, `organizationInvitation.*`, `organizationMembership.*`). On membership change it:
1. upserts the local user/org/membership, and
2. **translates the Clerk role** (`admin` / `basic_member` / **`content_creator`** — all three, via
   `TranslateRoleFromClerkToAnt`) and **pushes it into Sentinel** (`OrgAddUserToRole` /
   `OrgReplaceUserRole` / `OrgRemoveUserFromRole`). `app/internal/clerk/events/org_handler.go:229,:249`
   @ `ad9f3c498` — the role synced into Sentinel is whatever the translation returns, so the
   **synced-role set a Sentinel policy reviewer must reason about has three members, not two.**

So Clerk org roles become Sentinel roles at **sync time** — that's the only way Clerk influences backend authorization. The backend also writes membership/user metadata **back** to Clerk via `clerk-sdk-go/v2` (`organizationmembership`, `user`). See [Webhook Setup](../ops/webhook_setup.md) for local tunnel configuration.

## Dependent Repos & How They Integrate

Clerk ships a **separate package per framework** (Go, Next.js, Express, browser-JS, Expo). Each is a **thin adapter over Clerk's shared core** — one server-side core and one browser-side core — **not** a separate auth implementation. The column below lists the **package each repo actually installs** (its declared dependency); the internal/transitive modules those adapters pull in are inferable from the lockfile and intentionally not tracked here.

| Repo / app | Installed Clerk package(s) | What it's for |
|------------|----------------------------|---------------|
| **colony** (`/authn`) — the shared auth core; among the live Go services, imported by **`app` only** | `clerk-sdk-go/v2` | Verifies the session JWT (JWKS) + reads claims. sentinel / storage / messenger import it **0** times. |
| **app** (backend) | `clerk-sdk-go/v2`, `svix-webhooks/go` | Authn (via colony) + org/membership/invitation Backend-API writes + svix-verified webhook sync → Postgres + Sentinel. Also carries the folded jobsimulation + cms domains, which authenticate through the same middleware. |
| **jobsimulation, cms** — *frozen legacy repos* (no compose service, no `repos.yml` entry since `d11a403`) | *(none direct — was via `colony/authn`)* | Historical: authenticate only; authorization → Sentinel. Nothing starts them. |
| **storage, messenger** | — | No Clerk / no auth. |
| **sentinel** | — | Does **not** use Clerk; pure Casbin authorization. |
| **next-web-app** — `apps/web`, `apps/hiring`, `apps/integration` | `@clerk/nextjs` (+ `@clerk/localizations`) | Next.js App Router auth: `clerkMiddleware` route protection, `useAuth().getToken()` bearer, org/role gating (see deep-dive). |
| **next-web-app** — `apps/mobile` | `@clerk/clerk-expo` | Expo / React Native session (paused PoC). |
| **next-web-app** — `e2e` | `@clerk/testing`, `@clerk/backend` | **Test-only**: Playwright bypass token + programmatic user/session seeding. |
| **studio-desk** | `@clerk/clerk-js` (frontend), `@clerk/express` (backend) | Vanilla-TS browser SDK + Express middleware; admin tooling gated on the `STUDIO_ACCESS_ROLES` set — `org:admin` **and** `content_creator`. |
| **ant-academy** — web | `@clerk/nextjs` | `clerkMiddleware` in `proxy.js`; **requires ≥1 org membership** (`REQUIRE_ORGANIZATION_MEMBERSHIP` → `/no-organization`). |
| **ant-academy** — mobile | `@clerk/clerk-expo` | Expo session-only gate (custom email+password form). |

> **SDK versions:** `@clerk/nextjs` is **NOT aligned** across the four surfaces — this line asserted *"all four on `^6.39.2`"* until M257x iter-108. Measured at each clone's own HEAD: `next-web-app` `apps/{web,hiring,integration}` are all on **`^6.39.6`** (`8297c684`), while `ant-academy` (`code/package.json:52`) is on **`^6.39.2`** (`22df69dd`) — **three on one version, one a patch-range behind**. (`next-web-app/apps/maintenance` declares no `@clerk/nextjs`, so it is not a fourth surface.) **`@clerk/clerk-expo` is NOT aligned:** `next-web-app/apps/mobile/package.json:6` pins **`~2.6.18`** while `ant-academy/mobile/package.json:18` pins **`~2.19.36`** — thirteen minor versions apart, on the two mobile surfaces. The Go side has **drifted again**: `app/go.mod:31` @ `5ba17044` reads **`clerk-sdk-go/v2 v2.7.0`**, not the `v2.6.0` this doc previously asserted for both. Since `v2.6.0` is *the version the Clerkenstein Alignment DNA targets*, re-verify `colony`'s pin and the DNA before trusting an alignment score (`CHECK-M257x-iter22-clerk-sdk-drift`).

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
