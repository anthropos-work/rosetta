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
2. **The fake FAPI serves browser-trusted HTTPS — automatically.** `@clerk/clerk-js` + `clerkMiddleware` **always**
   reach the FAPI over `https://` (the host comes from the pk, prefixed `https://`), so the fake FAPI **terminates
   TLS** with a cert for the FAPI host. `up-injected.sh` (step 3a-bis) mints the cert into `<stack>/certs`; the
   override mounts it (`FAKE_FAPI_TLS_CERT/KEY`). **The bring-up makes the cert browser-trusted for you (M31):** when
   [`mkcert`](https://github.com/FiloSottile/mkcert) is on `PATH` it runs `mkcert -install` (idempotent) + mints a
   leaf for `127.0.0.1 localhost ::1`, so a fresh browser renders the signed-in app with **no proceed-anyway**. The
   bring-up keeps a pre-existing cert, so the trusted one survives re-ups. **No manual cert step is needed** — with
   the historical caveats below.
   - **First-ever `mkcert -install` on a fresh machine may prompt once for your OS password** (a GUI keychain write
     to add mkcert's local CA to the trust store). It's a one-time, machine-wide prompt; thereafter `-install` is a
     silent no-op. This is the only residual manual touch, and only on a brand-new box.
   - **openssl fallback (proceed-anyway).** If mkcert is **not installed** (or you set `DEMO_NO_MKCERT=1`, or a
     mkcert mint fails), the bring-up degrades to the **openssl self-signed** cert — byte-compatible, valid TLS, but
     **untrusted**, so the browser shows a warning and you click **"proceed anyway"** once (or import/trust the
     `<stack>/certs/fapi.crt` directly). This still works for automated (Playwright `ignoreHTTPSErrors`) verify.
     Install mkcert (`brew install mkcert`) to get the zero-touch path.
   - **Security note — a dev CA in your trust store.** `mkcert -install` adds mkcert's **local CA private key** to
     your OS (and, if `certutil` is present, Firefox) trust store. That is a real, if small, **trust expansion** —
     anything signed by that CA is trusted on your machine until you `mkcert -uninstall`. If you'd rather not, set
     **`DEMO_NO_MKCERT=1`** to force the openssl/proceed-anyway path; nothing else changes.
   - **Remote / VM demos.** For a **local, same-machine** demo, `mkcert -install` trusts only the **machine the
     bring-up runs on** — browse from a *different* machine and its browser hits the untrusted cert (proceed-anyway,
     or import the CA). **M213 (v2.2 "panorama") removes that limit for a tailnet demo:** bring the demo up with
     `/demo-up --public-host <magicdns>` (e.g. `billion.taildc510.ts.net`) and the FAPI cert is minted via
     **`tailscale cert`** — a real **Let's Encrypt** cert trusted **tailnet-wide with no per-machine CA install**, so
     a teammate's browser on the tailnet renders the signed-in app with **no proceed-anyway**. Same output paths
     (`<stack>/certs/fapi.{crt,key}`), so the mount + `ListenAndServeTLS` are unchanged; a mint failure (no
     `tailscaled` / not logged in) falls back to the mkcert/openssl **local-trust** path (non-fatal). The LE cert is
     **90-day** — `tailscale cert` re-issues on re-run; a long-lived stack needs a renew-then-reload step (M215). The
     one-clean-HTTPS-origin reverse proxy (`tailscale serve`) that fronts the *rest* of the browser surface lands in
     **M214** (`tailscale-serve.md`).
   - **Firefox needs `certutil`.** mkcert wires Chrome/Safari via the OS keychain automatically; **Firefox** has its
     own trust store and only picks up the CA when `certutil` is installed at `-install` time
     (`brew install nss`). Without it, Firefox falls back to proceed-anyway.
   - **Cert expiry.** The keep-existing guard never re-mints, and it has **no expiry check** — a long-lived stack
     could outlive its cert (openssl: 825 days; mkcert leaf: ~2.25 years) and silently re-blank. If a previously
     working demo suddenly bounces to `/login`, **`rm <stack>/certs/fapi.crt`** and re-up — the bring-up regenerates
     a fresh cert.
3. **The dev-instance handshake.** An unauthenticated load hits `clerkMiddleware`, which **307-redirects** to
   `https://<fapi>/v1/client/handshake?…&format=nonce`. The fake FAPI signs the demo user in and **303-bounces** back
   to the app with `?__clerk_handshake=<token>` carrying the `Set-Cookie` directives (`__session` + `__client_uat` +
   `__clerk_db_jwt` — the dev-browser cookie is what breaks the `dev-browser-missing` redirect loop). The fake FAPI
   also **proxies `clerk-js`** (`/npm/...`) and serves `/v1/environment` + `/v1/client`. The clerk-js proxy is the
   fake FAPI's **one outbound-egress dependency** — it fetches the bundle from **`cdn.jsdelivr.net`**, so the FAPI
   container needs outbound HTTPS to it (a `--public-host` bring-up runs a **non-fatal** host-side egress pre-check +
   warns if it's blocked). On a locked-down network, point the FAPI at a mirror with **`FAKE_FAPI_CLERKJS_CDN`**
   (M213).
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
> needs the FAPI to be **browser-trusted HTTPS** (M31: minted via mkcert at bring-up — see step 2), complete the
> **handshake** (nonce + dev-browser cookie), mint an **RS256** session the Node SDKs accept, and include the
> **`sid`** claim the client derives state from. All four are wired by Clerkenstein + the demo injection; the full
> JWT/handshake flow is the clerkenstein knowledge base (`knowledge/architecture.md` § Universal-key JWT /
> `knowledge/injection.md`).

## Verifying without a browser
The same identity can be exercised headlessly: mint a session token with the universal key
(`clerkenstein/shared.Mint`) for the `DefaultDemoUser` claims and call an authorized GraphQL/REST route — it
returns **200** with the seeded data (this is exactly the M7a login→200 proof; `membershipsCount` returns the
seeded member count). Use this for scripted smoke tests of a seeded stack.
