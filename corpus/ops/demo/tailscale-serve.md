# Remote demo access over Tailscale — the `--public-host` recipe

**Make a demo stack reachable from another machine on your Tailscale tailnet** — run a demo on a Tailscale VM
(e.g. `billion.taildc510.ts.net` on the odyssey Proxmox host) and a teammate with Tailscale up browses the whole
demo end-to-end, in their own browser, over HTTPS. This is the **external-shareability** surface of v2.2
"panorama".

It is **opt-in and default-off**: a plain `/demo-up N` is byte-identical to today (localhost only, no external
listener). Remote reach is requested explicitly with **one flag** — `--public-host <magicdns>` — and **Tailscale
itself is the access control** (only tailnet members can reach the host; there is no public internet exposure).

> **Scope of this doc.** The *recipe* + the *why-it-works*: the topology, what `--public-host` flips, the
> tailscale-cert FAPI, the CORS + cross-surface-link tail, and the teammate walkthrough. The knob plumbing is
> M212, the TLS/proxy/pk is M213, this origins-and-links layer is M214, and the **live cross-machine burn-down**
> (executing it on a real tailnet VM) is M215. Zero platform-repo edits throughout — tooling + docs + the opt-in
> flag only (two platform-family files ride the **existing** rext sha-pinned patch mechanism; see §"The patch
> tail").

---

## TL;DR — the one command

```bash
# on the Tailscale VM (tailscaled up + logged in), from stack-demo/rosetta-extensions:
/demo-up 1 --public-host billion.taildc510.ts.net
#   → a teammate on the tailnet opens  https://billion.taildc510.ts.net:13000  and browses demo-1 end-to-end
```

`--public-host` maps to the `STACK_PUBLIC_HOST` env knob (either form works). Unset ⇒ `localhost` (byte-identical
to a normal demo). The host **must be a dotted MagicDNS FQDN** — a dotless bare name is refused at bring-up
(`@clerk/backend`'s publishable-key host must be dotted; see [`../../services/clerkenstein.md`](../../services/clerkenstein.md)
§"Remote HTTPS over the tailnet").

`/stack-list` then shows the reachable external URL for the stack (the registry records `external_host` on the
public path — opt-in, non-fatal).

---

## The topology — HTTPS everywhere, one MagicDNS host, per offset port

A demo runs its browser-facing services on **offset ports** (`base + N*10000`): next-web `3000+off`, the Cosmo
GraphQL router `5050+off`, the backend REST `8082+off`, studio-desk `9000+off`, ant-academy `3077+off`, and the
fake Clerk FAPI `5400+off`. Under `--public-host`, each is reached over **HTTPS on the MagicDNS host at the same
offset port**:

```
teammate's browser ── https://billion.taildc510.ts.net:13000 ──▶  tailscale serve ──▶ http://127.0.0.1:13000  (next-web)
                   ── https://billion.taildc510.ts.net:15050 ──▶  tailscale serve ──▶ http://127.0.0.1:15050  (cosmo)
                   ── https://billion.taildc510.ts.net:18082 ──▶  tailscale serve ──▶ http://127.0.0.1:18082  (backend)
                   ── https://billion.taildc510.ts.net:19000 ──▶  tailscale serve ──▶ http://127.0.0.1:19000  (studio-desk)
                   ── https://billion.taildc510.ts.net:13077 ──▶  tailscale serve ──▶ http://127.0.0.1:13077  (ant-academy, native)
                   ── https://billion.taildc510.ts.net:15400 ──▶  fake-FAPI's OWN TLS (tailscale cert)          (Clerkenstein)
```

**Why HTTPS everywhere?** Clerk's `clerk-js` needs a **secure context** (Web Crypto) — a plain-`http://` MagicDNS
origin is not one, so HTTPS on the app origin is effectively required, not cosmetic (M213 decision D-SCHEME-1).

**Why per-port, not a single port-less `https://<host>`?** M213's reverse proxy is **`tailscale serve`** run
**per port**, PRESERVING the offset-port scheme (M213 decision D-PROXY-2): each browser-facing plaintext service
gets `tailscale serve --bg --https=<offsetport> http://127.0.0.1:<offsetport>`. So the only thing that changes
between a localhost demo and a remote demo is **`http://localhost:<port>` → `https://<magicdns>:<port>`** — same
port, `http`→`https`. (`tailscale serve` was chosen over a bundled Caddy for **zero net-new dependency** — the
`tailscale` CLI is already on every tailnet VM — and because it auto-terminates TLS with the node's cert. The
proxy plan is emitted by `rosetta-extensions/stack-injection/gen_tailscale_serve.py`; the fake-FAPI is
**excluded** from the proxy — it serves its own TLS, see below.)

The **asset plane stays on prod HTTPS, unchanged**: `DIRECTUS_PUBLIC_BASE_ADDR=https://content.anthropos.work` +
`MEDIA_URL=https://media.anthropos.work` are already HTTPS and allowlisted in next-web's `next.config.mjs`
`remotePatterns`, so browser images load over Tailscale with no change and no mixed-content.

---

## What `--public-host` flips (and what stays byte-identical)

Everything below is gated on the knob: **unset ⇒ byte-identical to a normal localhost demo.** A set MagicDNS host
flips exactly these, all derived from **one scheme predicate** (`https` for a dotted host, `http` for
localhost) so no site can drift:

| Surface | localhost demo | `--public-host` demo | Where |
|---|---|---|---|
| **Backend CORS** (`CORS_EXTRA_ORIGINS`) | `http://localhost:{3000,3001,9000}+off` | + `https://$HOST:{3000,3001,9000}+off` (the localhost trio is **kept** for on-host use) | `gen_injected_override.py` (`browser_scheme`); `app/internal/cors/cors.go` honors it in non-production |
| **studio-desk redirects** (`CLERK_SIGN_IN_URL` / `WEB_APP_URL`) | `http://localhost:3000+off` | `https://$HOST:3000+off` | `gen_injected_override.py` (runtime env) |
| **Baked browser URLs** (`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`, `_BACKEND_API_URL`, `_HOSTING_URL`, `_STUDIO_URL`, `_ACADEMY_URL`, `_PUBLIC_WEBSITE_URL`, `VITE_GRAPHQL_ENDPOINT`, `VITE_WEB_APP_URL`) | `http://localhost:…` | `https://$HOST:…` | `up-injected.sh` (`$SCHEME`) — the image cache-validators embed `$SCHEME` too, so an http-baked image is rebuilt under an https host |
| **studio-desk SPA sign-in** (`VITE_CLERK_SIGN_IN_URL`) | *(was the un-offset `localhost:3000/login` default)* → now `http://localhost:3000+off/login` | `https://$HOST:3000+off/login` | `up-injected.sh` — baked via a gitignored `.env.production.local` overlay (no Dockerfile ARG; see §"The patch tail") |
| **ant-academy** (`NEXT_PUBLIC_STUDIO_URL`, `next dev` bind, `allowedDevOrigins`) | `http://localhost:…`, loopback bind, hardcoded origins | `https://$HOST:…`, `-H 0.0.0.0`, MagicDNS host admitted | `ant-academy.sh` + the `ant-academy-dev-origins` patch |
| **FAPI cert** | `mkcert`/openssl (local trust) | `tailscale cert` (tailnet-wide trust) | `up-injected.sh` (M213) |
| **Cockpit / academy bind** | loopback | `0.0.0.0` | `up-injected.sh` / `ant-academy.sh` (M212) |
| **Registry** | no interaction | records `external_host` for `/stack-list` | `up-injected.sh` (M212) |

**Mixed-content clean.** With HTTPS-everywhere, no browser-facing call resolves to plain `http://` — the scheme
flip covers every baked endpoint, redirect, and cross-surface link; the asset plane is already prod-HTTPS. The
**one** deliberate plain-http surface is the **presenter cockpit's own page** (port `7700+off`): it is not in
`tailscale serve`'s front list, so it serves plain HTTP — an http launcher page linking/POSTing to the https demo
surfaces is fine (a navigation from http to https is not mixed content). Fronting the cockpit too is a live-
acceptance polish for M215.

---

## The tailscale-cert FAPI (the Clerk-free login over a real cert)

The Clerk-free browser login routes through Clerkenstein's **fake FAPI** over HTTPS (clerk-js always reaches the
FAPI over `https://`, derived from the publishable key). For a remote demo the FAPI cert is minted with
**`tailscale cert <magicdns>`** — a real Let's Encrypt MagicDNS cert **trusted tailnet-wide with no per-machine
CA install**, exactly what `mkcert` cannot give a *remote* browser. The fake-FAPI serves its **own** TLS with
that cert on `5400+off` (so it is excluded from the `tailscale serve` proxy — double-fronting would double-TLS).
The consumer mount is path-only (`<stack>/certs/fapi.{crt,key}`), so it is a drop-in at the same paths as the
local mkcert/openssl cert. Falls back to the local mkcert/openssl path (local trust only) if `tailscaled` isn't
up. Full cert story + caveats: [`recipe-browser-login.md`](recipe-browser-login.md) §B and
[`../../services/clerkenstein.md`](../../services/clerkenstein.md).

---

## The patch tail — two platform-family files, both via the existing sha-pinned mechanism

Two files in the platform **family** aren't reachable by the pure config/env layer, so they ride the **existing
rext sha-pinned patch mechanism** (drift-refuse, single-occurrence anchor, idempotent, non-fatal) applied to the
demo's **ephemeral clone** — **never a checked-in platform clone, never a canonical repo edit**:

1. **ant-academy `allowedDevOrigins` (required).** `next dev` blocks cross-origin dev requests from a host not in
   `code/next.config.js` `allowedDevOrigins` — which hardcodes a *different* tailnet host. The
   **`ant-academy-dev-origins`** demo-patch rewrites that array to also read an env var
   (`ANT_ACADEMY_ALLOWED_DEV_ORIGIN`), keeping the original entries (behavior-identical when unset, upstream-safe
   — the same shape as the `next-web-studio-url` demopatch). The host is supplied at `next dev` launch via the
   env; `ant-academy.sh` applies the patch before launch (gated on the public host) and reverts it on `--stop`.
   Manifest: `rosetta-extensions/demo-stack/patches/ant-academy-dev-origins/`; helper:
   `stack-injection/apply-ant-academy-dev-origins.sh` (apply|revert).

2. **studio-desk `VITE_CLERK_SIGN_IN_URL` (the SPA sign-in bake).** studio-desk's Dockerfile declares no ARG for
   it, so the SPA sign-in redirect falls back to the un-offset `http://localhost:3000/login`. Declaring the ARG is
   a platform-repo edit; a naive build-context `.env` is dropped by studio-desk's `.dockerignore` (`.env*`).
   Instead a gitignored **`.env.production.local`** overlay (vite loads it in production mode) bakes
   `$SCHEME://$HOST:3000+off/login`, admitted past the `.env*` exclusion by a *transient* `!.env.production.local`
   re-include — both reverted on the build's trap (clone left git-clean). This fixes the un-offset `:3000` default
   for **every** demo, and is https for a public host.

The **two already-shipped demopatches** — `next-web-studio-url` + `next-web-public-website-url` — carry the
MagicDNS baked value cleanly: their values are baked by `up-injected.sh` as `$SCHEME://$HOST:…`, so under a public
host they resolve demo-local over HTTPS (no prod-eject, no mixed-content).

### Documented residual — next-web `WEB_APP_URL` / `HIRING_APP_URL` (#M214-D-URLS-1)

These `urls.ts` constants are `NEXT_PUBLIC_NODE_ENV` ternaries → prod (`app.`/`hiring.anthropos.work`) with no
per-URL override, so they *would* prod-eject if traversed. **They are a documented residual, not patched** —
decided with evidence: the demo's target flows do **not** render them as off-demo links. The M42e (employee) +
M42m (manager) coverage sweeps gate at **0 prod-ejects** and surfaced only `STUDIO_URL` + `PUBLIC_WEBSITE_URL`
(both fixed); `WEB_APP_URL`/`HIRING_APP_URL` never surfaced. Their `apps/web` usages are public marketing chrome
(anonymous-only), PDF/SEO metadata (non-navigation), a dead Clerk.provider fallback (the demo bakes
`NEXT_PUBLIC_HOSTING_URL`), and hiring-product features (not a Workforce demo flow). Under HTTPS-everywhere they
would be https-prod (not mixed-content), only an eject on flows the demo never exercises. If a future coverage
sweep ever surfaces one of these hosts, the fix is a demopatch mirroring `next-web-studio-url` — the mechanism is
proven and ready. (See [`coverage-protocol.md`](coverage-protocol.md) §"fix-surface routing".)

---

## Walkthrough — a teammate browses the demo

1. **On the VM** (Tailscale up + logged in, ≥ 12 GB Docker VM, the repos cloned):
   `/demo-up 1 --public-host billion.taildc510.ts.net`. The bring-up mints the tailscale cert, generates + applies
   the `tailscale serve` plan, bakes the offset+https URLs, applies the two patches, seeds the stories world, and
   launches the native academy + cockpit. `/stack-list` shows the reachable URL.
2. **On the teammate's laptop** (Tailscale up, a member of the same tailnet): open
   `https://billion.taildc510.ts.net:13000` — the demo's next-web, over a trusted cert. Log in through the
   Clerkenstein flow (or use the presenter cockpit at `http://billion.taildc510.ts.net:17700` → pick a hero →
   **[Log in as]**). Browse the profile, the enterprise Workforce-Intelligence dashboards, the library, Studio
   (`:19000`), and the Academy (`:13077`) — all demo-local, all over HTTPS.
3. **Tear down:** `/demo-down 1` — stops the containers, reaps the native cockpit + academy (reverting the
   ant-academy patch), frees the registry slot; the dev stack is untouched.

---

## Safety framing

- **Opt-in, default-off.** No demo is externally reachable unless `--public-host` is passed. A bare `/demo-up N`
  binds loopback only and is byte-identical to today.
- **Tailscale is the access control.** The host is reachable **only** to members of your tailnet — there is no
  public-internet exposure, no port-forward, no open 0.0.0.0-on-the-LAN surprise beyond the tailnet. Binding
  `0.0.0.0` is gated on the knob precisely so it is never ambient.
- **Zero platform-repo edits.** The whole surface is rext tooling + this doc + the opt-in flag. The two
  platform-family patches touch only the demo's **ephemeral** clone via the sha-pinned mechanism (drift-refuse
  fails loud on an upstream change; reverted on teardown), never a canonical repo.
- **The demo's data-isolation guarantees are unchanged.** Remote reach changes the *origin/scheme*, not the data
  plane: the tenant-data firewall, the per-stack isolated Postgres, and the never-write-prod boundary all hold
  exactly as documented in [`../safety.md`](../safety.md).

---

## See also

- [`../rosetta_demo.md`](../rosetta_demo.md) — the demo lifecycle + the `--public-host` operator surface.
- [`frontend-tier.md`](frontend-tier.md) — the UI tier build (offset URLs, the CORS `CORS_EXTRA_ORIGINS`, the
  studio-desk requireAuth fallback) — the HTTPS/remote deltas are cross-referenced there.
- [`../../services/clerkenstein.md`](../../services/clerkenstein.md) — the fake FAPI/BAPI, the tailscale-cert
  remote path, and the dotted-publishable-key host rule.
- [`recipe-browser-login.md`](recipe-browser-login.md) — the interactive Clerk-free login + the full cert story.
- [`coverage-protocol.md`](coverage-protocol.md) — the 0-prod-eject believability gate (the evidence base for the
  `urls.ts` residual decision).
- Design decisions: `knowledge/plan/releases/02.20-panorama/` — M212 (the knob), M213 (TLS/proxy/pk, D-PROXY-2 /
  D-SCHEME-1), M214 (origins & links, D-SCHEME-1 / D-VITE-SIGNIN-1 / D-URLS-1), M215 (the live acceptance).
