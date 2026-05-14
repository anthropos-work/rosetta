# Staging Clerk — Dev App, Shared Login, and the Monkey-Patch

The dev Clerk app shared across every Anthropos personal staging, the shared test login that lets any engineer log into any staging, and the load-bearing `clerk-fetch-fix.js` that makes Server Components talk to Clerk inside Docker on HTTP.

Companion to:

- [`staging-bringup.md`](./staging-bringup.md) — fresh-VM onboarding.
- [`staging-sync.md`](./staging-sync.md) — daily sync routine.

---

## The dev Clerk app

| Field            | Value                                                      |
| ---------------- | ---------------------------------------------------------- |
| Name             | `national-elk-17`                                          |
| Instance ID      | `ins_3DIWRX3yChqGneZ9gI0JpGvHPKA`                          |
| API base         | `https://api.clerk.com`                                    |
| Publishable key  | `pk_test_…` (in `platform/.env` as `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY`) |
| Secret key       | `sk_test_…` (in `platform/.env` as `CLERK_SECRET_KEY`)     |
| Dashboard        | https://dashboard.clerk.com (use Stefano's email to access) |

**Shared across every staging on the tailnet.** Ithaca, Calypso, your `wip-*` — all point at the same dev Clerk app. Users you create show up everywhere.

**This is dev, not prod.** Production has a separate Clerk app — the dev keys do NOT work in prod, and prod keys must NEVER be used on a staging machine (any user-mutating action you trigger from staging would hit real users).

---

## Shared test login

```
Email:    stefano@anthropos.work
Password: chichi88kora
```

**This account is a shared cross-engineer testing instrument.** Anyone on the team can log into any staging using these credentials, and they "become" Stefano for testing — Clerk authenticates, the dev DB has a `users` row keyed to Stefano's prod UUID, and the casbin grants are wired so Stefano is admin in all the orgs he was admin of in the prod snapshot.

There's also `elisa@anthropos.work` (Elisa's account, same dev Clerk app) for testing as Elisa. Add more shared accounts as the team grows — see [§ Adding a shared test account](#adding-a-shared-test-account).

**Caveat: dev Clerk is shared, so all sessions are visible across the team.** Don't put real-customer data into the dev Clerk user pool (names, emails, phone numbers from prod imports — fine; sensitive PII or test fixtures you don't want a teammate to see — not fine). The dev Clerk dashboard shows every user, session, org, and event to everyone with dashboard access.

### Playwright usage

For automated smoke tests (the sync routine's `anthropos-staging-smoke.js`) use the real password — it exercises the real-user latency profile and catches Clerk-JS bootstrap regressions:

```js
await page.goto(`${SMOKE_URL}/login`);
await page.fill('input[name="identifier"]', SMOKE_EMAIL);
await page.click('button:has-text("Continue")');
await page.fill('input[name="password"]', SMOKE_PASSWORD);
await page.click('button:has-text("Continue")');
await page.waitForSelector('text=Hi, Stefano', { timeout: 120_000 });
```

For programmatic CI flows where Clerk's "new device" challenge blocks you (Quirk #13), use a one-shot sign-in token instead:

```bash
TOKEN=$(curl -s -X POST https://api.clerk.com/v1/sign_in_tokens \
  -H "Authorization: Bearer $CLERK_SECRET_KEY" -H "Content-Type: application/json" \
  -d "{\"user_id\":\"user_3DIYdXgwlr0Q0R12qDNbk4z95aZ\",\"expires_in_seconds\":600}" \
  | python3 -c "import json,sys;print(json.load(sys.stdin)['token'])")
echo "$SMOKE_URL/login?__clerk_ticket=$TOKEN"
```

(That `user_3DI…` is Stefano's Clerk user id in the dev app.) Don't use this for human-like smoke tests — it bypasses a chunk of the Clerk JS bootstrap and hides real-user latency.

### Adding a shared test account

If the team wants a third shared identity (e.g. for an SRE rotation):

```bash
export CLERK_SECRET=sk_test_…

curl -s -X POST https://api.clerk.com/v1/users \
  -H "Authorization: Bearer $CLERK_SECRET" -H "Content-Type: application/json" \
  -d '{
        "email_address": ["sre@anthropos.work"],
        "password": "<strong-password>",
        "first_name": "SRE",
        "last_name": "Test",
        "skip_password_checks": true
      }'
```

Then on each staging host, run the rebind procedure (§ 3 of [`staging_from_dump.md`](./staging_from_dump.md)) to wire the new Clerk user to a DB user UUID and the casbin grants. Document the new account here.

---

## Allowed origins

Clerk only accepts requests from origins in its `allowed_origins` list. Currently (2026-05-14) there are 13:

```
http://localhost:3000
http://localhost:3001
http://localhost:5050
http://anthropos.local:3000
http://ithaca:3000
http://ithacastaging:3000
https://ithacastaging.taildc510.ts.net
http://calypso:3000
http://calypsostaging:3000
https://calypsostaging.taildc510.ts.net
http://100.120.254.65:3000
http://100.83.121.80:3000
https://anthropos.work
```

### Adding a new staging host

When you spin up `wip-<initials>`, you need to add **three** things in three different places — Clerk allowed_origins, Tailscale ACL `hosts:`, and backend CORS. Skip any one and Clerk / the backend / the browser will reject your origin.

1. **Tailscale ACL `hosts:`** — see [`staging-bringup.md` § Tailnet membership](./staging-bringup.md#tailnet-membership). Add both `wip-<initials>` and `wip-<initials>staging` aliases.

2. **Clerk `allowed_origins`** — PATCH the dev instance:

   ```bash
   export CLERK_SECRET=sk_test_…

   # Get current allowed_origins
   curl -s "https://api.clerk.com/v1/instance" \
     -H "Authorization: Bearer $CLERK_SECRET" \
     | python3 -c "import json,sys;print('\n'.join(json.load(sys.stdin)['allowed_origins']))"

   # Append your new origin (preserve all existing entries)
   curl -s -X PATCH "https://api.clerk.com/v1/instance" \
     -H "Authorization: Bearer $CLERK_SECRET" -H "Content-Type: application/json" \
     -d '{
       "allowed_origins": [
         "http://localhost:3000",
         "http://localhost:3001",
         "http://localhost:5050",
         "http://anthropos.local:3000",
         "http://ithaca:3000",
         "http://ithacastaging:3000",
         "https://ithacastaging.taildc510.ts.net",
         "http://calypso:3000",
         "http://calypsostaging:3000",
         "https://calypsostaging.taildc510.ts.net",
         "http://100.120.254.65:3000",
         "http://100.83.121.80:3000",
         "https://anthropos.work",
         "http://wip-mn:3000",
         "http://wip-mnstaging:3000"
       ]
     }'
   ```

3. **Backend CORS** — until [`anthropos-work/app#feat/cors-extra-origins-env`](https://github.com/anthropos-work/app/pulls) lands, edit `app/internal/cors/cors.go` `colony.Development` branch on your staging clone and rebuild backend. Once merged, set `CORS_EXTRA_ORIGINS=http://wip-<initials>staging:3000` in `platform/.env` and `docker compose restart backend` — no rebuild needed.

Test the round-trip:

```bash
curl -sI -X OPTIONS "https://api.clerk.com/v1/client/handshake" \
  -H "Origin: http://wip-<initials>staging:3000" \
  -H "Access-Control-Request-Method: POST" | head -5
# expect: 200 with Access-Control-Allow-Origin echoing your origin
```

---

## Pitfalls that bit us

These are the Clerk-related symptoms that ate days of bringup time. Each has a clear root cause and a fix.

### Symptom: `UND_ERR_CONNECT_TIMEOUT` from Server Components

**What you see:** First request to a protected route (`/home`, `/profile`, `/settings`) after a fresh container start works. Every subsequent request renders `error.tsx` ("Ops, something went wrong. Try again.") and the server log spits:

```
clerkError: api_response_error { errors: [{code:"unexpected_error", message:"fetch failed"}] }
UND_ERR_CONNECT_TIMEOUT (api.clerk.com:443, 10s)
```

User-visible symptom: "login page never moves on" — the Clerk `<SignIn>` ticket exchange completes and cookies are set, but the post-login redirect target itself errors.

**Root cause:** A stale-state issue inside Node 24's embedded undici when invoked from Next.js 15.5's wrapped fetch. Plain `node -e "fetch('https://api.clerk.com/v1/users/<id>')"` from inside the same container, with the same `CLERK_SECRET_KEY`, succeeds repeatedly. Only the Next.js Server Component fetch path is broken.

Verified non-fixes (in case you're tempted to redo the diagnosis):

- `NODE_OPTIONS=--dns-result-order=ipv4first`
- `NODE_OPTIONS=--no-network-family-autoselection`
- `extra_hosts: api.clerk.com → 104.18.37.202` (the IPv4-only Cloudflare anycast)

None fix it. The Docker bridge is IPv4-only (`platform_app-network`, `EnableIPv6=false`), so IPv6 isn't the culprit either.

**Fix in place:** a bootstrap file that replaces `globalThis.fetch` with the installed undici 7.25.0 `fetch`, using a global dispatcher that disables connection re-use (every outbound request opens a fresh TCP/TLS connection).

Loaded via `NODE_OPTIONS=--require=/clerk-fetch-fix.js`, mounted into the container via compose:

```yaml
# platform/docker-compose.yml -- next-web-app service
next-web-app:
  volumes:
    - ./clerk-fetch-fix.js:/clerk-fetch-fix.js:ro
  environment:
    - NODE_OPTIONS=--max-old-space-size=4096 --require=/clerk-fetch-fix.js
```

Confirm it loaded:

```bash
docker compose logs --since 10m next-web-app | grep "clerk-fetch-fix"
# expect: [clerk-fetch-fix] globalThis.fetch replaced (undici 7.25.0, no keepalive)
```

The file itself lives at `/home/devops/platform/clerk-fetch-fix.js` on Ithaca (and the equivalent path on every other staging). It's hand-edited (Stefano's edits to the Secure-cookie section in particular are load-bearing) — copy verbatim, don't refactor.

Full content (verbatim from Ithaca, 2026-05-14):

```js
// Workaround for a fetch bug in this staging stack:
//
//   Symptom: every request to a protected route (`/home`, `/profile`, etc.)
//   renders the error fallback ("Ops, something went wrong") because the
//   `(authenticated)/layout.tsx` Server Component calls `currentUser()` and
//   `auth().getToken()` which fail with `UND_ERR_CONNECT_TIMEOUT` to
//   api.clerk.com:443 after the first request.
//
// Plain `node -e "fetch(...)"` from the same container hits api.clerk.com
// fine, repeatedly. Only the Next.js Server Component fetch path is broken.
// `--dns-result-order=ipv4first`, `--no-network-family-autoselection`, and
// hardcoding `api.clerk.com -> IPv4` in `extra_hosts` all failed to fix it —
// suggesting a stale-state issue inside Next.js's wrapped fetch / Node's
// embedded undici.
//
// Workaround: replace `globalThis.fetch` with the installed undici 7.25.0
// `fetch`, and set a global dispatcher that disables connection re-use (every
// request gets a fresh TCP/TLS connection, exactly like a one-off `node -e`).
// Loaded via `NODE_OPTIONS=--require=...` so it runs before Next.js starts.
const path = '/app/node_modules/.pnpm/undici@7.25.0/node_modules/undici';
const { fetch: undiciFetch, Agent, setGlobalDispatcher } = require(path);

const dispatcher = new Agent({
  keepAliveTimeout: 1,
  keepAliveMaxTimeout: 1,
  pipelining: 0,
  connect: { timeout: 30_000 },
});
setGlobalDispatcher(dispatcher);

globalThis.fetch = undiciFetch;

console.log('[clerk-fetch-fix] globalThis.fetch replaced (undici 7.25.0, no keepalive)');

// ---------------------------------------------------------------------------
// Clerk-Secure-cookie fix for HTTP staging (ithacastaging / calypsostaging)
// ---------------------------------------------------------------------------
//
//   Symptom: after login, /home (and other protected pages) load *very slowly*
//   (60+ seconds) and the three-dot loading spinner stays for minutes. The
//   server logs spam:
//     Clerk: Refreshing the session token resulted in an infinite redirect loop.
//
//   Root cause: Clerk's FAPI handshake endpoint always returns cookies with
//   `Secure; SameSite=None`. The staging stack serves over plain HTTP
//   (`http://ithacastaging:3000`), so the browser silently drops every Secure
//   cookie. The middleware then never sees the session cookie on the next
//   request, kicks off another handshake, and after 3 retries gives up and
//   returns "signed-out". Every protected RSC prefetch hits this 3x cycle,
//   stacking up many seconds of latency before the page finally renders via
//   the client-side Clerk JS (which writes cookies without `Secure` via
//   document.cookie and eventually unblocks).
//
//   Fix: intercept `Set-Cookie` header writes for Clerk-prefixed cookies and
//   strip `Secure` + downgrade `SameSite=None` to `SameSite=Lax` when the
//   instance is a dev/test instance (pk_test) running over HTTP. The
//   downgraded cookies persist in the browser, the middleware sees them on
//   the next request, the handshake loop never triggers, and pages render
//   immediately.
//
//   This is a dev-only workaround — production runs on HTTPS and the original
//   Secure cookies are exactly what we want. We guard on `pk_test_` so the
//   patch is a no-op in production.
const isDevClerk =
  (process.env.NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY || '').startsWith('pk_test_') ||
  (process.env.CLERK_SECRET_KEY || '').startsWith('sk_test_');

if (isDevClerk) {
  const CLERK_COOKIE_RE = /^(__client_uat|__session|__clerk_db_jwt|__clerk_handshake|__clerk_handshake_nonce|__clerk)(_[A-Za-z0-9\-]+)?=/;

  const downgradeCookie = (value) => {
    if (typeof value !== 'string') return value;
    // Quick exit if it doesn't look like a Clerk cookie or isn't Secure.
    if (!CLERK_COOKIE_RE.test(value)) return value;
    let v = value;
    // Strip "; Secure" (any case, optional whitespace) — anywhere in the string.
    v = v.replace(/;\s*Secure(?=\s*;|\s*$)/gi, '');
    // Downgrade "SameSite=None" to "SameSite=Lax". Lax is sent on top-level
    // navigations, which is exactly what middleware sees.
    v = v.replace(/SameSite=None/gi, 'SameSite=Lax');
    return v;
  };

  // Patch Headers.prototype.append and .set so anything written via the
  // standard web-Headers API gets the downgrade. Next.js / Clerk middleware
  // both use this path.
  for (const method of ['append', 'set']) {
    const orig = Headers.prototype[method];
    Headers.prototype[method] = function (name, value) {
      if (typeof name === 'string' && name.toLowerCase() === 'set-cookie') {
        value = downgradeCookie(value);
      }
      return orig.call(this, name, value);
    };
  }

  // Also patch Node's ServerResponse.setHeader / appendHeader (used by some
  // adapter code paths that bypass the web Headers object).
  try {
    const http = require('node:http');
    for (const method of ['setHeader', 'appendHeader']) {
      const proto = http.ServerResponse.prototype;
      if (typeof proto[method] !== 'function') continue;
      const orig = proto[method];
      proto[method] = function (name, value) {
        if (typeof name === 'string' && name.toLowerCase() === 'set-cookie') {
          if (Array.isArray(value)) {
            value = value.map(downgradeCookie);
          } else if (typeof value === 'string') {
            value = downgradeCookie(value);
          }
        }
        return orig.call(this, name, value);
      };
    }
  } catch (e) {
    console.warn('[clerk-fetch-fix] could not patch ServerResponse:', e?.message);
  }

  console.log(
    '[clerk-fetch-fix] dev Clerk detected — Secure flag will be stripped from Clerk cookies for HTTP staging compatibility'
  );
}
```

**Do NOT ship this to production.** Disabling undici connection re-use globally is fine on a staging singleton but wasteful at production traffic. The upstream bug (Next.js 15.5 Server Components + Clerk SDK 2.33.2 + Node 24 + Docker IPv4-only bridge → `UND_ERR_CONNECT_TIMEOUT`) is worth a Clerk + Next.js issue report.

**Skip-worktree on the compose patch:** the `docker-compose.yml` next-web-app block (volumes mount + `NODE_OPTIONS`) is one of the long-lived skip-worktree files on every staging clone. Without it, the `--require` never fires and you get the timeout chain.

### Symptom: Secure-cookie drop on HTTP staging

**What you see:** After login, `/home` (and other protected pages) load *very slowly* (60+ seconds), the three-dot loading spinner stays for minutes, and the server log spams `Clerk: Refreshing the session token resulted in an infinite redirect loop.`

**Root cause:** Clerk's FAPI handshake endpoint always returns cookies with `Secure; SameSite=None`. The staging stack serves over plain HTTP (`http://<host>staging:3000`), so the browser silently drops every Secure cookie. The middleware then never sees the session cookie, kicks off another handshake, and after 3 retries gives up and returns "signed-out". Every protected RSC prefetch hits this 3x cycle.

**Fix:** the same `clerk-fetch-fix.js` above also intercepts `Set-Cookie` writes for Clerk-prefixed cookies and strips `Secure` + downgrades `SameSite=None` to `SameSite=Lax` when the instance is `pk_test_` (dev). The `if (isDevClerk)` block in the file handles this. Make sure the file is the verbatim Ithaca version — the cookie regex covers `__client_uat`, `__session`, `__clerk_db_jwt`, `__clerk_handshake`, `__clerk_handshake_nonce`, `__clerk`.

**Production-safe:** the `pk_test_` / `sk_test_` guard makes this a no-op in production.

### Symptom: login page's `localStorage.clear()` corrupting Clerk state

**What you see:** Slow / janky login experience after a fresh client load. The `<SignIn>` component sometimes loses track of partially-completed factor exchanges and asks you to re-enter credentials mid-flow.

**Root cause:** `next-web-app/apps/web/src/app/(unauthenticated)/login/[[...login]]/page.tsx` calls `window.localStorage.clear()` and `window.sessionStorage.clear()` in a mount `useEffect`. Clerk JS stores transient handshake state in `localStorage` (`__clerk_*` keys), so clearing it on every login-page mount corrupts in-flight sign-ins.

**Fix:** the 2026-05-14 cleanup opened [`anthropos-work/next-web-app#fix/login-page-no-localstorage-clear`](https://github.com/anthropos-work/next-web-app/pulls) to remove the clear calls. Until merged, you can safely apply the same edit on your staging clone via skip-worktree.

### Symptom: first-load latency 60-100s in dev

**What you see:** After a fresh sign-in, the dashboard takes 60-100 seconds to fully render. There's a ~30s gap with no network activity after redirect off `/login` where Clerk JS is bootstrapping in dev mode, then GraphQL fan-out begins.

**Root cause:** Clerk dev-mode has heavier JS bootstrap than prod, plus the cold backend GraphQL subgraphs warm up serially on first hit.

**Not a regression — wait ≥90 seconds.** Subsequent navigations within the same session are fast. Worth investigating as a follow-up (likely needs Clerk dev-mode flag or SDK upgrade), but treat a stuck three-dot spinner under 90s as "still loading," not "bug."

When writing Playwright tests, use selector waits, not time waits:

```js
await page.waitForSelector('text=Hi, Stefano', { timeout: 120_000 });
// NOT: await page.waitForLoadState('networkidle')
// NOT: await page.waitForTimeout(30_000)
```

### Symptom: "new device" challenge blocks CI

**What you see:** Programmatic / Playwright sign-ins from a new IP get a Clerk email-code challenge that breaks automation.

**Root cause:** Clerk's "new device" challenge is enabled on `national-elk-17`.

**Fix for CI:** generate a one-shot sign-in token (see [§ Playwright usage](#playwright-usage)) and visit `/login?__clerk_ticket=<token>`. Token consumes once, creates a session, no challenge.

**Fix for humans:** receive the email code once on each new device. Clerk then trusts the device.

---

## Customizing the dev Clerk session token (recommended)

Clerk's session token doesn't include your DB UUID by default, so colony's auth path has to fetch your user/org from the Clerk Backend API on every request. Clerk rate-limits aggressively (HTTP 429), and the first few page loads after a fresh login can fail with "organization mismatch".

**Fix:** in the Clerk dashboard → Sessions → "Customize session token", add:

```json
{
  "eid": "{{user.external_id}}",
  "email": "{{user.primary_email_address}}",
  "firstname": "{{user.first_name}}",
  "lastname": "{{user.last_name}}",
  "org": {
    "eid": "{{org.public_metadata.eid}}"
  }
}
```

This is dashboard-only as of 2026-05; there's no public REST endpoint to script it. Once set, every issued session token carries the DB UUIDs in its claims and colony's auth path can skip the per-request Backend API fetch.

### Dev app currently has no custom template (2026-05-14)

Verified during the Ithaca repair: `national-elk-17` is issuing **bare v2 session tokens** with no custom claims attached. The dashboard "Customize session token" template appears empty (or was never configured). Programmatic workarounds tried and failed:

- `PATCH /v1/instance` with `{"session_token_template": {"claims": …}}` returns HTTP 204 but the next-issued JWT has none of the custom claims. The endpoint is silently accepted-and-ignored — session-token customization is genuinely dashboard-only.
- `POST /v1/jwt_templates` with name `colony` creates a template (returns a `jtmp_…` id) but only takes effect if the **frontend** explicitly calls `getToken({template: 'colony'})`. `next-web-app` doesn't — it uses the bare session token via Clerk middleware. Extending the frontend to use a named template would also require teaching the backend to accept template-issued JWTs (different audience / signing rules), so it's not a small change.

Until someone with dashboard access logs in and applies the template, **plan on the lazy Clerk Backend API fetch path** described in [`staging-bringup.md` Quirk #11](./staging-bringup.md#bringup-quirks-consolidated-as-a-procedural-narrative). The colony v2-JWT patch already implements this with a process-wide eid cache (`clerk-org-id → DB uuid`), so per-request load on the Clerk API is bounded to one-fetch-per-distinct-org-per-process-lifetime.

### Anatomy of a v2 session token

The token shape colony has to read is now v2 — flat in `v: 1` (`org_id`, `org_role`, `org.eid`), nested in `v: 2` under a single `o` key. Both shapes live in the wild because Clerk only flipped the default for new apps; older apps can be migrated via dashboard toggle, but `national-elk-17` was created post-flip and only emits v2.

```jsonc
// What a freshly-issued dev session token (decoded) looks like today:
{
  "exp": 1778764223,
  "iat": 1778764163,
  "iss": "https://national-elk-17.clerk.accounts.dev",
  "nbf": 1778764153,
  "o": {                                       // ← v2 nesting
    "id":  "org_3DIcUoSDXZjxmD82ipAh7xMijhi",  // ← v1 name was "org_id"
    "rol": "admin",                            // ← v1 name was "org_role"
    "slg": "while-true-srl-1777973042707468570"
  },
  // NOTE: "public_metadata" is NOT included by default in v2 tokens —
  // the dashboard custom template is the only built-in way to inject it.
  // colony falls back to `OrganizationClient.Get()` to resolve eid.
  "sid": "sess_3DiTGfLL0JDk4tnHwN1m6cdPZrp",
  "sts": "active",
  "sub": "user_3DIYdXgwlr0Q0R12qDNbk4z95aZ",
  "v": 2                                       // ← format version marker
}
```

You can capture a real one with:

```bash
SESS=sess_…                              # from Clerk dashboard → Sessions, or browser cookies
curl -s -X POST "https://api.clerk.com/v1/sessions/$SESS/tokens" \
  -H "Authorization: Bearer $CLERK_SECRET_KEY" \
  | python3 -c "import json,sys,base64; t=json.load(sys.stdin)['jwt']; \
    print(json.dumps(json.loads(base64.urlsafe_b64decode(t.split('.')[1] + '==')), indent=2))"
```

If a token you capture has `"v": 1` or top-level `org_id`/`org_role`, your colony build doesn't need the v2 fallback. If it's `"v": 2` with nested `o.*`, the colony build that resolves `GetOrganization()` must support v2 (vendored patch or the future upstream fix).

---

## Related

- [`staging-bringup.md`](./staging-bringup.md) — fresh-VM onboarding.
- [`staging-sync.md`](./staging-sync.md) — daily sync routine.
- [`staging_from_dump.md`](./staging_from_dump.md) — engineer-rebind reference (creates Clerk users + remaps DB).
- Ant-singularity catalog: [`auto-anthropos-staging-dev-loop.md`](https://github.com/stefano-anthropos/ant-singularity/blob/main/knowledge/singularity-catalog/auto-anthropos-staging-dev-loop.md) — full dev-loop blueprint.
