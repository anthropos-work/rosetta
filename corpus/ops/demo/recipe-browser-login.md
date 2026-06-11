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

`/demo-up` bakes this end-to-end; you just open the browser. What it does, and *why* each piece is needed (the
fake FAPI must satisfy the **full Clerk dev-instance handshake**, not just serve a session — that's the part the
early "mint a pk and log in" sketch missed):

1. **Mint the publishable key** for the demo's fake FAPI host. The key is `pk_test_<base64(host$)>`, byte-identical
   to Clerkenstein's `MintPublishableKey` (`inject.py`'s `mint_pk` emits it). **The host is `127.0.0.1:5400+N·10000`,
   not `localhost`** — `@clerk/backend`'s pk validator requires a **dot** in the decoded host, so a dotless
   `localhost` pk is rejected as invalid (a 500 on every request).
2. **The fake FAPI serves HTTPS.** `@clerk/clerk-js` + `clerkMiddleware` **always** reach the FAPI over `https://`
   (the host comes from the pk, prefixed `https://`), so the fake FAPI **terminates TLS** with a cert for the FAPI
   host. `up-injected.sh` generates the cert into `<stack>/certs`; the override mounts it (`FAKE_FAPI_TLS_CERT/KEY`).
   The openssl-generated cert is **self-signed**, so the browser won't trust it out of the box. **One-time operator
   step (pick one):** run `mkcert -install` *and* mint a cert for `127.0.0.1` into `<stack>/certs` (the bring-up
   keeps a pre-existing cert, so a mkcert-issued one survives re-ups) — `mkcert -install` needs your machine
   password to add its local CA to the OS/Firefox trust stores; **or** import/trust the generated self-signed cert
   directly. Without a trusted cert the browser blocks clerk-js's cross-origin FAPI calls and the app bounces back
   to `/login`. (`mkcert -install` *alone* does not help — its CA never signed the openssl cert.)
3. **The dev-instance handshake.** An unauthenticated load hits `clerkMiddleware`, which **307-redirects** to
   `https://<fapi>/v1/client/handshake?…&format=nonce`. The fake FAPI signs the demo user in and **303-bounces** back
   to the app with `?__clerk_handshake=<token>` carrying the `Set-Cookie` directives (`__session` + `__client_uat` +
   `__clerk_db_jwt` — the dev-browser cookie is what breaks the `dev-browser-missing` redirect loop). The fake FAPI
   also **proxies `clerk-js`** (`/npm/...`) and serves `/v1/environment` + `/v1/client`.
4. **`__session` is RS256, verified networklessly.** The Node SDKs (`@clerk/nextjs`, `@clerk/express`) **reject
   HS256** and verify the session as RS256 via `CLERK_JWT_KEY` (the fixed demo public key, supplied as **runtime
   container env** — filled per-demo into `.env.demo-N` by `up-injected.sh`, not build-baked) or the
   **BAPI `/v1/jwks`** (reachable from the app *container* via the `api.clerk.com` alias — sidesteps the
   localhost split-horizon). The disarmed Go `authn` accepts **both** algs (`shared.ParseAny`), so the same RS256
   browser token also works as the backend API bearer. The minted token carries a **`sid`** (session id) claim —
   without it `@clerk/nextjs`'s client `useDerivedAuth` sees a user with no session and throws *"Invalid state"* on
   the first render.
5. **Log in + land in the seeded org.** Open the frontend (`http://localhost:3000+N·10000`); it auto-signs-in as
   `DefaultDemoUser` — `user_clerkenstein` / `demo@anthropos.test`, **admin** of `org_clerkenstein`. Because the
   auto-set-dress seeded that identity as an admin member (+ its casbin grant + the global Sentinel policy),
   authorized routes return **200** — the populated workforce, not a 403 wall.

> **Why this is more than "mint a pk."** A pk alone points clerk-js at the fake FAPI, but a real dev-instance login
> needs the FAPI to be **browser-trusted HTTPS**, complete the **handshake** (nonce + dev-browser cookie), mint an
> **RS256** session the Node SDKs accept, and include the **`sid`** claim the client derives state from. All four
> are wired by Clerkenstein + the demo injection; the full JWT/handshake flow is the clerkenstein knowledge base
> (`knowledge/architecture.md` § Universal-key JWT / `knowledge/injection.md`).

## Verifying without a browser
The same identity can be exercised headlessly: mint a session token with the universal key
(`clerkenstein/shared.Mint`) for the `DefaultDemoUser` claims and call an authorized GraphQL/REST route — it
returns **200** with the seeded data (this is exactly the M7a login→200 proof; `membershipsCount` returns the
seeded member count). Use this for scripted smoke tests of a seeded stack.
