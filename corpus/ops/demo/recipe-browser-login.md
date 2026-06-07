# Recipe — Interactive browser login (Clerk-free)

**Goal.** Complete the *interactive* demo: open a browser, log in as the demo user with **no real Clerk**, and
land in a **seeded** org where authorized routes return **200**. This is where the two M3-deferred injection
recipes land — the **`api.clerk.com` cert-redirect** (so the backend's orgclient trusts the fake BAPI over
TLS) and the **browser-login walk-through** (the frontend points at the fake FAPI via a minted publishable
key).

**Prerequisite.** A stack up (`/demo-up N`) and seeded (`/stack-seed N`) — the demo identity
`user_clerkenstein` must exist as a member (otherwise login authenticates but org-gated routes 403). See
[`recipe-enterprise-onboarding.md`](recipe-enterprise-onboarding.md).

## The three Clerk seams (all disarmed; full reference: clerkenstein `knowledge/injection.md`)

| Seam | Consumer | Disarm |
|---|---|---|
| **Backend session verify** | the 5 Go services (`authn`) | `go.mod replace` with the disarmed `colony/authn` (done by `apply-authn.sh` at bring-up — **already wired**) |
| **Backend org client** | `app/internal/clerk/orgclient` → `api.clerk.com` | **redirect `api.clerk.com` → the fake BAPI** (`clerk-backend`): container DNS / `extra_hosts` + a **TLS cert the app container trusts** (this recipe) |
| **Browser login** | `@clerk/nextjs` / `@clerk/clerk-js` | a **minted publishable key** encoding the fake FAPI host → the SDK talks to the fake FAPI (`clerk-frontend`), config-only, no SDK fork (this recipe) |

M3 proved the **authn seam live** (a running app accepts a Clerkenstein token on a protected route — 403, not
401, before seeding). This recipe finishes the *interactive* loop.

## A — the `api.clerk.com` cert-redirect (backend orgclient → fake BAPI)

The orgclient is app-internal + networked, so it can't be `go.mod replace`d; it's disarmed by **redirecting its
host**. The fake BAPI (`demo-N-fake-bapi`) serves the `api.clerk.com/v1` surface; the one redirect catches both
the SDK sub-clients and the three raw-HTTP methods.

1. **Route the host.** Add `api.clerk.com` → the fake BAPI to the app/cms/etc. containers' resolution
   (`extra_hosts:` in the injected compose override, or `/etc/hosts` in the container).
2. **Trust the TLS.** The SDK calls `https://api.clerk.com/v1`, so the app container must **trust the fake
   BAPI's certificate** for that host — mount a cert for `api.clerk.com` signed by a CA the container trusts
   (add the CA to the container trust store), and serve it from the fake BAPI. The full mechanism +
   cert-generation steps are in the clerkenstein repo: **`clerk-backend/doc.go`**.

> Status: the fake BAPI server + its behavior are **built and alignment-gated** (Go gate + real-SDK test); the
> DNS/cert *wiring into a live demo stack* is this documented recipe (the "recipe-only" item from M3, landed in
> M8). The backend authn seam (the 403-not-401 proof) already runs without it — the cert-redirect is needed for
> the orgclient (org/membership reads), not for token verification.

## B — the browser-login walk-through (frontend → fake FAPI)

1. **Mint the publishable key** for the demo's fake FAPI host (`localhost:5400+N·10000`). The key is
   `pk_test_<base64(host$)>`, byte-identical to Clerkenstein's authoritative `MintPublishableKey`; the demo
   tooling's `inject.py` (`mint_pk`) emits it.
2. **Rebuild the frontend with it.** Set `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` (next-web-app) /
   `VITE_CLERK_PUBLISHABLE_KEY` (studio-desk) to the minted key and bring up the frontend. `@clerk/clerk-js`
   decodes the host from the key and talks to the fake FAPI — **no real Clerk, no SDK fork**.
3. **Log in.** Open the frontend; the fake FAPI serves the `DefaultDemoUser` session — `user_clerkenstein` /
   `demo@anthropos.test`, **admin** of `org_clerkenstein`. The browser is now authenticated.
4. **Land in the seeded org.** Because `/stack-seed` created `user_clerkenstein` as a seeded admin member (+ the
   casbin grant + the global Sentinel policy), authorized routes return **200** — the populated workforce, not
   a 403 wall.

## Verifying without a browser
The same identity can be exercised headlessly: mint a session token with the universal key
(`clerkenstein/shared.Mint`) for the `DefaultDemoUser` claims and call an authorized GraphQL/REST route — it
returns **200** with the seeded data (this is exactly the M7a login→200 proof; `membershipsCount` returns the
seeded member count). Use this for scripted smoke tests of a seeded stack.
