# Remote demo access over Tailscale — the `--public-host` runbook

**Make a demo stack reachable from another machine on your Tailscale tailnet** — run a demo on a Tailscale VM
(e.g. `billion.taildc510.ts.net` on the odyssey Proxmox host) and a teammate with Tailscale up browses the whole
demo end-to-end, in their own browser, over HTTPS. This is the **external-shareability** surface of v2.2
"panorama".

> **The demo-patch mechanism is specified in [`demopatch-spec.md`](demopatch-spec.md).** It is the sanctioned **zero-platform-edit escape hatch**: patch the demo's own ephemeral clone before the image build, revert after — the canonical repos are never touched. Read it before adding or re-pinning a patch. Since M217 the gate is **self-healing**: the *anchor* is the contract, the whole-file sha is only a baseline.

It is **opt-in and default-off**: a plain `/demo-up N` is byte-identical to today (localhost only, no external
listener). Remote reach is requested explicitly with **one flag** — `--public-host <magicdns>` — and **Tailscale
itself is the access control** (only tailnet members can reach the host; there is no public internet exposure).

> **This is a PROVEN recipe, not a plan.** M215 executed it live on a real Linux VM (`billion`, odyssey Proxmox
> host) on **2026-07-11**: a teammate on a **different** tailnet machine logged in as a seeded hero and completed
> a real journey over trusted HTTPS — for **both** the employee vantage (`maya-thriving` → `/profile`) and the
> manager vantage (`dan-manager` → `/enterprise/workforce`), `ignoreHTTPSErrors:false`, 0 console errors, 0
> functional request failures. Everything below is grounded in that run; the host-prereq + rext-fix set it
> surfaced (findings F1–F12) is baked into this runbook.

> **Scope of this doc.** Two halves. **Part 1 — the runbook:** stand up a remote demo on a **fresh Linux VM**
> unaided (prereqs, GitHub-via-PAT, secrets, workspace, the `--public-host` bring-up, verify, teardown).
> **Part 2 — how it works:** the topology, what `--public-host` flips, the tailscale-cert FAPI, the CORS +
> cross-surface-link tail, and the safety framing. The knob plumbing is M212, the TLS/proxy/pk is M213, the
> origins-and-links layer is M214, and the **live cross-machine acceptance** is M215. Zero platform-repo edits
> throughout — tooling + docs + the opt-in flag only (two platform-family files ride the **existing** rext
> sha-pinned patch mechanism; see §"The patch tail").

---

## TL;DR — the one command

```bash
# on the Tailscale VM (tailscaled up + logged in), from the rext demo-stack section:
STACK_PUBLIC_HOST=billion.taildc510.ts.net \
  bash demo-stack/up-injected.sh 1 --public-host billion.taildc510.ts.net
#   → a teammate on the tailnet opens  https://billion.taildc510.ts.net:13000  and browses demo-1 end-to-end
```

On a box with a full rosetta checkout + Claude Code, the equivalent operator surface is the skill:
`/demo-up 1 --public-host billion.taildc510.ts.net` — it drives the exact same `up-injected.sh`. On a **bare
headless VM** you run the `bash demo-stack/up-injected.sh` form directly (that is what the M215 live run used).

`--public-host` maps to the `STACK_PUBLIC_HOST` env knob (either form works). Unset ⇒ `localhost` (byte-identical
to a normal demo). The host **must be a dotted MagicDNS FQDN** — a dotless bare name is refused at bring-up
(`@clerk/backend`'s publishable-key host must be dotted; see [`../../services/clerkenstein.md`](../../services/clerkenstein.md)
§"Remote HTTPS over the tailnet").

`/stack-list` then shows the reachable external URL for the stack (the registry records `external_host` on the
public path — opt-in, non-fatal).

---

# Part 1 — The runbook (a fresh Linux VM → a remote demo)

A bare Ubuntu VM needs a few **host-side** prereqs a dev box doesn't — the Docker builds compile Go in-image, but
the rext orchestration tooling (`stacksecrets`/`stacksnap`/`stackseed`, all Go) + `atlas migrate` + the cert/serve
steps run on the **host**. Then it's clone → secrets → bring up → verify.

## Step 0 — VM host prerequisites (proven on odyssey `billion`, 2026-07-11)

| # | Prereq | Why | Install | Finding |
|---|--------|-----|---------|---------|
| 1 | **Docker + Compose** | builds + runs the whole stack | present on the odyssey VMs | — |
| 2 | **Go 1.25.x** (matches rext's `toolchain go1.25.12`) | the host rext tooling is Go; without it secret provisioning is skipped → `no usable platform .env` → abort | `curl -sSfL https://go.dev/dl/go1.25.12.linux-amd64.tar.gz \| sudo tar -C /usr/local -xz` then add `/usr/local/go/bin` to `PATH` | **F2** |
| 3 | **atlas CLI** | `migrate-demo.sh` runs `atlas migrate apply`; without it the schema is created with **0 tables** and every seeder fails `relation public.X does not exist` | `curl -sSfL https://release.ariga.io/atlas/atlas-linux-amd64-latest -o atlas && sudo install -m755 atlas /usr/local/bin` | **F8** |
| 4 | **Tailscale operator** | so the bring-up's un-sudo'd `tailscale cert` / `tailscale serve` run as the deploy user; else the cert falls back to `mkcert` = **local-trust-only** and a remote browser distrusts it | one-time `sudo tailscale set --operator=<deploy-user>` | **F1** |
| 5 | **An ssh-agent** | the platform compose declares `ssh: default`, so `buildx bake` needs `SSH_AUTH_SOCK` at definition-load — even though the private-module pulls use the `GH_PAT`, not the agent. A **keyless** agent suffices | the bring-up **auto-starts one if absent** (`eval "$(ssh-agent -s)"`); no key needed | **F4** |
| 6 | **The snapshot cache** (content surfaces only) | the taxonomy (~42,790 public skills) + Directus content are **set-dressed from the snapshot cache**, not migrations. Without `.agentspace/snapshots` on the VM, `public.skills=0` and the library/skills surfaces are sparse (identity/profile/dashboard still render fully) | `scp` the `.agentspace/snapshots` cache to the VM, or capture per [`../snapshot-cold-start.md`](../snapshot-cold-start.md) | **F9** |

The canonical prereq list + these install commands also live in
[`../setup_guide.md`](../setup_guide.md) §"Linux host prerequisites (for a remote/VM demo over Tailscale)".

**What the tooling now does for you** (so a bare VM doesn't re-trip the M215 findings):

- **Pre-flights + fails loud** on the three host-toolchain prereqs — **Go** (F2), **atlas** (F8), and the
  **tailscale operator** (F1). A missing one aborts with a clear message instead of silently producing an empty
  schema or a locally-trusted cert.
- **Auto-handles** the two Linux-only footguns: it **auto-starts a keyless ssh-agent** when `SSH_AUTH_SOCK` is
  unset (F4), and it **pre-creates the bind-mount data dirs with open perms** so the Bitnami Postgres container
  (UID 1001) can write a host dir Docker would otherwise create root-owned (F6 — the manual fix was
  `sudo chmod -R 777 $STACK/data`; **not** needed on macOS, where Docker Desktop remaps the perms).

> **Cross-repo follow-up — the `kb-ant-business` `odyssey` skill / `reference_devserver.md` is STALE.** The M215
> `billion` run surfaced that it lists **4** VMs (there are now ~**13**) and claims the VMs ship **Go 1.26**
> (`billion` had **none** — hence the Go host-prereq above). That doc lives in a **different repo, out of the
> rosetta corpus** — flagged here for whoever owns the odyssey KB to refresh; not fixed from here.

> **12 GB Docker-VM floor (UI tier).** The full UI tier's next-web build spikes to ~3.7 GB. The
> [`frontend-tier.md`](frontend-tier.md) §"The 12 GB Docker-VM prerequisite" floor applies; on the M215 run RAM
> held (5.7 GB free) with next-web + studio-desk up. A RAM-tight VM can do a backend-only first pass with
> `DEMO_NO_UI=1` (Step 5).

## Step 1 — GitHub access without an org SSH key (PAT-over-HTTPS)

A fresh VM has no org SSH key, but `ensure-clones.sh` / `make init` clone the private `anthropos-work` repos over
`git@github.com:` by default. Rather than provision an SSH key, use the **`GH_PAT`** (already in the secret
bundle, Step 3) to clone over HTTPS:

```bash
git config --global url."https://github.com/".insteadOf git@github.com:
git config --global credential.helper store          # or: cache
# then prime the credential store once so the PAT is used for github.com HTTPS pulls
```

With the `insteadOf` rewrite + a credential store, every `git@github.com:anthropos-work/<repo>.git` clone the
bring-up issues resolves to `https://github.com/anthropos-work/<repo>.git` and authenticates with the PAT. The
**same** `GH_PAT` then reaches the Docker builds as the `GH_ACCESS_TOKEN` build-arg (pulling the private Go
modules in-image). See [`../setup_github_guide.md`](../setup_github_guide.md) for the canonical SSH path (the
alternative when you do have an org key).

## Step 2 — The workspace layout on the VM

Lay out a single root (`<root>`, e.g. `~/panorama`) with **only** the rext consumption clone + the agentspace:

```
<root>/
  stack-demo/
    rosetta-extensions/          # the pinned-tag consumption clone (Step 2a)
  .agentspace/
    secrets/                     # the secret source (Step 3)
    rext.tag                     # a one-line pinned tag string (Step 2a)
    snapshots/                   # the snapshot cache (Step 0 #6 / Step 4, optional)
```

The bring-up derives `REPO_ROOT` relative to the `stack-demo/rosetta-extensions` clone (so `<root>` is the parent
of `stack-demo`) and **clones the platform + all peer repos into `stack-demo/`** itself — a `stack-demo`-only box
(no `stack-dev`) brings a demo up end-to-end (the v1.8 "understudy" self-contained model; see
[`../rosetta_demo.md`](../rosetta_demo.md) §"A demo builds from its OWN clone set").

**Step 2a — clone rext at the pinned tag.** Record the release's pinned tag in `.agentspace/rext.tag` (a bare
one-line tag string; `#`-comments + blank lines + CRLF tolerated), then check the consumption clone out at it:

```bash
mkdir -p <root>/stack-demo <root>/.agentspace
echo "<panorama-tag>" > <root>/.agentspace/rext.tag       # the tag carrying M212–M215 + the F1–F8 fixes
git clone https://github.com/anthropos-work/rosetta-extensions.git <root>/stack-demo/rosetta-extensions
# M217: `git fetch --tags` is MANDATORY. A fresh clone does NOT necessarily carry the tag you are about to
# check out, and a bare `checkout <tag>` then dies `pathspec did not match` — or, worse, silently leaves the
# clone on a bare sha. THE OMISSION OF THIS LINE IS EXACTLY HOW `billion` ENDED UP ON AN UNTAGGED COMMIT
# (panorama-m214-3-g41a28aa) that then warned about itself on every bring-up for a whole release.
git -C <root>/stack-demo/rosetta-extensions fetch --tags origin
git -C <root>/stack-demo/rosetta-extensions checkout -f "$(cat <root>/.agentspace/rext.tag)"
git -C <root>/stack-demo/rosetta-extensions describe --tags --exact-match   # MUST print the pinned tag
```

> **The pin is now enforced, not suggested.** Since M217 a mismatch between the clone's checkout and
> `.agentspace/rext.tag` **aborts the bring-up** (`DEMO_ALLOW_UNPINNED_REXT=1` to override). Detached HEAD is the
> correct end state — `ensure-clones.sh` keys on `git describe --tags --exact-match`, so leaving the clone on a
> branch trips the guard even when the content is right.

`.agentspace/rext.tag` is the single source-of-truth both `/demo-up` and `ensure-clones.sh` read (M49 #1). The
consumed tag **must** carry the M215 host fixes — the pre-flights (F1/F2/F8), the auto ssh-agent (F4), the
pre-created data dirs (F6), and the `git for-each-ref` build-tag resolver (F3: the old
`git tag --list | head -1` SIGPIPE'd → 141 → `set -e` aborted the bring-up on a many-tag repo like `app`'s ~337
v-tags; **fixed** to a pipe-less `git for-each-ref --count=1`).

## Step 3 — Provision the secrets (values-blind)

Copy the curated secret source onto the VM (e.g. `scp` it to `<root>/.agentspace/secrets`, user-approved), then
let the bring-up's auto-provision step assemble each repo's `.env` from it. The rext `stacksecrets` provisioner
(Go, run on the host) writes `stack-demo/platform/.env` (and the other per-repo targets) **values-blind** — no
verb ever reads, echoes, or logs a secret value. The `GH_PAT` in the bundle both authenticates the HTTPS clones
(Step 1) and rides into the Docker builds as `GH_ACCESS_TOKEN`.

The secret source is laid out **by repo** (`<root>/.agentspace/secrets/<repo>/<target-file>`); the full layout +
the 6-repo/56-gene coverage DNA + the `DIRECTUS_TOKEN`-stays-blank safety are in
[`../secrets-spec.md`](../secrets-spec.md). `/demo-up` runs `/stack-secrets` as an auto-provision step, so a demo
is self-sourced from `.agentspace/secrets`; you can also pre-run it explicitly with `/stack-secrets demo-1`.

> **The `.env`-presence guard runs AFTER provision (M49 #3).** On a `stack-demo`-only VM there is no `stack-dev`
> to seed `platform/.env` from, so the guard that aborts on "no usable platform .env" fires only after
> `/stack-secrets` has had its chance to write it from `.agentspace/secrets`. A box with **neither** a
> `stack-dev` seed **nor** a usable secret source aborts loud here (the genuine unprovisionable case).

## Step 4 — (optional) The snapshot cache, for content surfaces

Identity/profile/dashboard/workforce surfaces work with **zero** content set-dressing. But the
**taxonomy/library/skills** surfaces are set-dressed from the snapshot cache (F9), so if you want them populated,
put the cache on the VM before bring-up:

```bash
scp -r <local>/.agentspace/snapshots  <root>/.agentspace/snapshots     # from a box that has captured it
```

Without it, the seed logs `taxonomy=skipped(cache-miss)` and `public.skills=0` (the manager funnel's
mapped/verified counts read 0), while identity heroes still seed fully (201 users / 1 org on the M215 run). To
fill the cache from scratch when you have none, follow [`../snapshot-cold-start.md`](../snapshot-cold-start.md).
Skip content entirely with `DEMO_NO_LOCAL_CONTENT=1` (Step 5) to read content live from prod
(`content.anthropos.work`, already HTTPS + allowlisted).

## Step 5 — The bring-up (`--public-host`)

```bash
cd <root>/stack-demo/rosetta-extensions
STACK_PUBLIC_HOST=billion.taildc510.ts.net \
  bash demo-stack/up-injected.sh 1 --public-host billion.taildc510.ts.net
```

**Trim flags** (env knobs, all opt-in) for a first pass or a RAM-tight VM:

| Flag | Effect | Use it for |
|------|--------|-----------|
| `DEMO_NO_UI=1` | backend + Clerkenstein only, no frontend build / academy | a fast backend-only de-risk pass, or a RAM-tight VM (the M215 tik-1 pass) |
| `DEMO_NO_LOCAL_CONTENT=1` | no per-stack Directus; content read live from prod | skip the snapshot-cache dependency (Step 4) |
| `DEMO_NO_STORIES=1` | the legacy structural small-200 seed instead of the Stories & Heroes world | — |
| `DEMO_NO_VERIFY=1` | skip the bring-up-tail auto-verify | — |

**What the bring-up does, in order** (the `--public-host` knob threads through every step): pre-flight the host
prereqs (Go/atlas/operator) → `ensure-clones.sh` (bootstrap-clone platform + `make init` the peers into
`stack-demo/`) → provision `.env` from `.agentspace/secrets` → pre-create the Linux data dirs writable (F6) →
ensure a keyless ssh-agent (F4) → build the injected Go services + (unless `DEMO_NO_UI`) the two frontends with
the **offset + `https://$HOST` URLs + the minted pk baked** → `migrate-demo.sh` (`atlas migrate apply`) → mint the
FAPI cert via **`tailscale cert`** → generate + apply the **`tailscale serve`** per-port plan → set-dress
(snapshot replay + the Stories seed) → reload Sentinel's Casbin policy → launch the native academy + presenter
cockpit → the non-fatal auto-verify on the offset ports. `/stack-list` then shows the reachable URL.

## Step 6 — Verify (the exact curls + the cockpit login)

Offset ports are `base + N·10000`; for `demo-1` (N=1) the offset is `+10000`. From **any tailnet PEER**, the
plaintext services are fronted by `tailscale serve` over the trusted cert, and the FAPI serves its own TLS — so
**no `-k`/`--insecure`** is needed (the whole point: a genuinely trusted Let's Encrypt cert):

```bash
HOST=billion.taildc510.ts.net
curl -s  https://$HOST:18082/api/health                                   # backend  (8082+off) → OK
curl -s -o /dev/null -w '%{http_code}\n' https://$HOST:15050/health       # cosmo    (5050+off) → 200
curl -s -o /dev/null -w '%{http_code}\n' https://$HOST:15400/v1/client    # FAPI-own-TLS (5400+off) → 200
```

All three answered `verify=0` (cert trusted, no CA install) from a **remote** Mac on the M215 run — the
make-or-break proof that the M213/M214 remote-auth foundation works on a real Linux VM.

> ### ⚠️ NOT from the VM itself — `tailscale serve` is bypassed on the loopback path (M219)
>
> This section used to read *"from any tailnet machine **(or the VM itself)**"*. **The parenthetical is false**,
> and it cost M219 a full false-RED sweep before it was diagnosed.
>
> `docker-proxy` binds the demo's offset ports on **`0.0.0.0`**, which includes the VM's own `100.x` tailscale
> address. A connection originating **on the VM** to `https://<magicdns>:<port>` therefore lands on the **kernel
> socket** — the container, speaking plain HTTP — instead of being intercepted by `tailscaled`'s `serve` layer,
> which is what terminates TLS. Plain HTTP answering a TLS handshake yields:
>
> ```
> curl: (35) OpenSSL/3.0.13: error:0A00010B:SSL routines::wrong version number
> ```
>
> Measured on `billion`: **from the VM, https on `:13000`, `:15050` and `:18082` ALL fail TLS; from a tailnet
> peer all three answer 307/200/200.** From a *peer*, WireGuard delivers the packet to `tailscaled`, which serves
> the trusted cert — which is why `tailscale serve status` can list a mapping that nevertheless does not apply to
> traffic you originate locally.
>
> **Consequence for testing.** A `--public-host` demo bakes the MagicDNS origin into the frontend build, so the
> app's own GraphQL client calls `https://<magicdns>:15050/graphql`. Drive that app from a browser **on the VM**
> and every GraphQL call dies `ERR_SSL_PROTOCOL_ERROR`, every page renders a permanent loading spinner, and every
> content assert fails for reasons that have nothing to do with the product. **Browser-driven suites (the
> coverage sweep, the Playthroughs) must run from a tailnet PEER** — see
> [`coverage-protocol.md` § WHERE you run the sweep is part of the test](coverage-protocol.md).

**The cockpit login (the interactive proof).** Open the presenter cockpit at `http://$HOST:17700` (`7700+off`,
plain HTTP — it is deliberately *not* fronted by `tailscale serve`; navigating from an http launcher page to the
https demo surfaces is not mixed content). It lists the seeded heroes; each **[Log in as]** is a link to the FAPI
handshake:

```
https://$HOST:15400/v1/client/handshake?__clerk_identity=<hero>&redirect_url=https://$HOST:13000/<jump>
```

The fake FAPI sets the session and redirects into next-web, landing the hero authenticated at
`https://$HOST:13000/<jump>`. Proven live from a remote headless Chromium (`ignoreHTTPSErrors:false`, 0 console
errors) for **both** vantages:

- **Employee** — `maya-thriving` → `/profile` (the M41 ProfileSeeder depth: work history, certs, education,
  projects rendered).
- **Manager** — `dan-manager` → `/enterprise/workforce?tab=skills-verification` ("Workforce Intelligence",
  fully rendered with real seeded structural data — 221 members, 445 AI sims, 47 skill paths, …).

For a browserless scripted smoke, mint a universal-key session and call an authorized route — see
[`recipe-browser-login.md`](recipe-browser-login.md) §"Verifying without a browser".

## Step 7 — Teardown

```bash
/demo-down 1          # or: bash demo-stack/rosetta-demo down 1
```

Stops the containers, reaps the native cockpit + academy (reverting the ant-academy patch), and frees the
registry slot; the dev stack (if any) is untouched. Add `--purge` to also drop the stack's images + data dir.

**The teardown also RESETS this demo's `tailscale serve` ports (F12).** `tailscale serve` binds the tailnet IP
`:<offsetport>` as a REAL listener whose config **persists past `docker compose down`** (it is node-level, not
per-container) — so without a reset a re-deploy on the same offset ports fails `address already in use` (the new
backend can't bind `0.0.0.0:<offsetport>`). `/demo-down` clears **just this demo's** browser-facing offset ports
(per-port `tailscale serve --https=<port> off`, offset-scoped so a co-resident `demo-N`'s serve is never
clobbered), gated on the demo having been public + tailscale present, non-fatal. The up-path also pre-clears
these ports before (re)configuring serve, so a re-up is idempotent. Byte-identical / no-op on a localhost
bring-up and where `tailscale` is absent.

> **If you ever tear a public-host demo down BY HAND** (e.g. `docker rm` the containers instead of `/demo-down`),
> the `tailscale serve` config is **not** cleared — run **`tailscale serve reset`** before the next `--public-host`
> deploy, or it will port-conflict. (`/demo-down` does this per-port for you; the blanket `reset` is the manual
> catch-all.)

---

# Part 2 — How it works

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
origin is not one, so HTTPS on the app origin is effectively required, not cosmetic (M213-D-SCHEME-1).

**Why per-port, not a single port-less `https://<host>`?** M213's reverse proxy is **`tailscale serve`** run
**per port**, PRESERVING the offset-port scheme (M213 decision D-PROXY-2): each browser-facing plaintext service
gets `tailscale serve --bg --https=<offsetport> http://127.0.0.1:<offsetport>`. So the only thing that changes
between a localhost demo and a remote demo is **`http://localhost:<port>` → `https://<magicdns>:<port>`** — same
port, `http`→`https`. It is **not** a single 443. (`tailscale serve` was chosen over a bundled Caddy for **zero
net-new dependency** — the `tailscale` CLI is already on every tailnet VM — and because it auto-terminates TLS
with the node's cert. The proxy plan is emitted by
`rosetta-extensions/stack-injection/gen_tailscale_serve.py`; the fake-FAPI is **excluded** from the proxy — it
serves its own TLS, see below.)

The **asset plane stays on prod HTTPS, unchanged**: `DIRECTUS_PUBLIC_BASE_ADDR=https://content.anthropos.work` +
`MEDIA_URL=https://media.anthropos.work` are already HTTPS and allowlisted in next-web's `next.config.mjs`
`remotePatterns`, so browser images load over Tailscale with no change and no mixed-content.

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
acceptance polish left as an accepted future enhancement (M215 shipped with the cockpit deliberately plain-HTTP).

## The tailscale-cert FAPI (the Clerk-free login over a real cert)

The Clerk-free browser login routes through Clerkenstein's **fake FAPI** over HTTPS (clerk-js always reaches the
FAPI over `https://`, derived from the publishable key). For a remote demo the FAPI cert is minted with
**`tailscale cert <magicdns>`** — a real Let's Encrypt MagicDNS cert **trusted tailnet-wide with no per-machine
CA install**, exactly what `mkcert` cannot give a *remote* browser. The fake-FAPI serves its **own** TLS with
that cert on `5400+off` (so it is excluded from the `tailscale serve` proxy — double-fronting would double-TLS).
The consumer mount is path-only (`<stack>/certs/fapi.{crt,key}`), so it is a drop-in at the same paths as the
local mkcert/openssl cert. Falls back to the local mkcert/openssl path (local trust only) if `tailscaled` isn't
up **or** the tailscale operator isn't set (F1 — set it, Step 0 #4). The LE cert is **90-day**; `tailscale cert`
re-issues on re-run, and a long-lived stack needs a renew-then-reload step. Full cert story + caveats:
[`recipe-browser-login.md`](recipe-browser-login.md) §B step 2 and
[`../../services/clerkenstein.md`](../../services/clerkenstein.md).

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

## What the tooling auto-handles, pre-flights, and fails loud on (the M215 finding set)

The live `billion` run surfaced the exact host-prereq + rext-fix set a fresh Linux VM needs. Where each landed:

| Finding | What it was | Resolution class |
|---|---|---|
| **F1** — `tailscale cert` needs elevation; the tooling calls it un-sudo'd | else the cert silently falls back to `mkcert` (local-trust-only) and a remote browser distrusts it | **pre-flight + prereq** — set the tailscale operator (Step 0 #4) |
| **F2** — the host needs Go for the rext tooling | no Go → secret provisioning skipped → `no usable platform .env` → abort | **pre-flight + prereq** — install Go 1.25.x (Step 0 #2) |
| **F3** — `git tag --list \| head -1` SIGPIPE → 141 → `set -e` aborts | reproduces on a many-tag repo (`app` ~337 v-tags) | **rext fix (shipped)** — pipe-less `git for-each-ref --count=1` in `up-injected.sh` |
| **F4** — buildx bake needs `SSH_AUTH_SOCK` (`ssh: default`) even though pulls use the PAT | a bare host with no ssh-agent fails at definition-load | **auto-handled** — the bring-up starts a **keyless** ssh-agent when absent (the PAT still does the real pulls) |
| **F5** — two `app` demopatches refused (target-role authz-skip, ai-readiness loadMembers) | sha-drift on the current `app` tag; **non-fatal** (demo works, slower per-member fan-out) | **known issue** — a demopatch re-anchor, separate from the remote story |
| **F6** — Linux bind-mount data-dir perms (Bitnami UID 1001 can't write a root-owned host dir) | `mkdir: /bitnami/postgresql/data: Permission denied`; macOS Docker Desktop remaps, native Linux does not | **auto-handled** — the bring-up pre-creates the data dirs writable (manual fix was `sudo chmod -R 777 $STACK/data`) |
| **F7/F8** — the host needs the `atlas` CLI; without it `migrate` creates 0 tables → every seeder fails | `migrate-demo.sh` treated `atlas`-missing as a non-fatal warning that masked a total migration failure | **pre-flight + prereq** — install atlas (Step 0 #3); the bring-up now fails loud on `atlas`-not-found |
| **F9** — the taxonomy is set-dressed from the snapshot cache, not migrations | no `.agentspace/snapshots` on the VM → `public.skills=0`, sparse library/skills surfaces | **prereq (optional)** — scp/capture the cache (Step 4); identity/profile/dashboard work without it |
| **F11** — hero identity vs profile-name mismatch (cosmetic) | e.g. logged in as `maya-thriving`, profile person renders as a different generated name | **known issue (seed polish)** — login + render work; unrelated to the remote story |
| **F12** — the teardown didn't reset `tailscale serve` → a re-deploy port-conflicts | `serve` binds the tailnet IP `:<offsetport>` as a listener that persists past `compose down`; the next deploy fails `address already in use` (surfaced on the billion cold reset-to-seed) | **rext fix (shipped)** — `/demo-down` resets THIS demo's serve ports (per-port `--https=<port> off`, offset-scoped), + a defensive up-path pre-reset (idempotent re-up); non-fatal, no-op on localhost. Manual by-hand-teardown unblock: `tailscale serve reset` (Step 7) |

> **Numbering note.** The ledger skips **F10** (unused). **F13** — a jobsimulation-service startup crash — is
> **off the proven journey path** and out of this runbook's host-deploy scope (it would hit any demo, remote or
> local); it is recorded in the milestone findings ledger + routed to standing backlog, not baked here.

## Safety framing

- **Opt-in, default-off.** No demo is externally reachable unless `--public-host` is passed. A bare `/demo-up N`
  binds loopback only and is byte-identical to today.
- **Tailscale is the access control.** The host is reachable **only** to members of your tailnet — there is no
  public-internet exposure, no port-forward, no open 0.0.0.0-on-the-LAN surprise beyond the tailnet. Binding
  `0.0.0.0` is gated on the knob precisely so it is never ambient. The teammate's client must keep Tailscale
  **MagicDNS on** (the default) for the `<magicdns>` name to resolve.
- **Zero platform-repo edits.** The whole surface is rext tooling + this doc + the opt-in flag. The two
  platform-family patches touch only the demo's **ephemeral** clone via the sha-pinned mechanism (drift-refuse
  fails loud on an upstream change; reverted on teardown), never a canonical repo.
- **The demo's data-isolation guarantees are unchanged.** Remote reach changes the *origin/scheme*, not the data
  plane: the tenant-data firewall, the per-stack isolated Postgres, and the never-write-prod boundary all hold
  exactly as documented in [`../safety.md`](../safety.md).

---

## See also

- [`../rosetta_demo.md`](../rosetta_demo.md) — the demo lifecycle + the `--public-host` operator surface + the
  self-contained `stack-demo` clone-set model.
- [`frontend-tier.md`](frontend-tier.md) — the UI tier build (offset URLs, the CORS `CORS_EXTRA_ORIGINS`, the
  studio-desk requireAuth fallback) — the HTTPS/remote deltas are cross-referenced there.
- [`../setup_guide.md`](../setup_guide.md) — §"Linux host prerequisites (for a remote/VM demo over Tailscale)" (the canonical
  prereq list Step 0 mirrors).
- [`../secrets-spec.md`](../secrets-spec.md) — the values-blind secret provisioning (`/stack-secrets`) the VM
  runs from `.agentspace/secrets`.
- [`../../services/clerkenstein.md`](../../services/clerkenstein.md) — the fake FAPI/BAPI, the tailscale-cert
  remote path, and the dotted-publishable-key host rule.
- [`recipe-browser-login.md`](recipe-browser-login.md) — the interactive Clerk-free login + the full cert story +
  the browserless smoke.
- [`../snapshot-cold-start.md`](../snapshot-cold-start.md) — filling the snapshot cache on a fresh box (F9).
- [`../setup_github_guide.md`](../setup_github_guide.md) — the canonical SSH GitHub path (the alternative to the
  PAT-over-HTTPS clone in Step 1).
- [`coverage-protocol.md`](coverage-protocol.md) — the 0-prod-eject believability gate (the evidence base for the
  `urls.ts` residual decision).
- Design decisions: `knowledge/plan/releases/archive/02.20-panorama/` — M212 (the knob), M213 (TLS/proxy/pk, D-PROXY-2 /
  M213-D-SCHEME-1), M214 (origins & links, M214-D-SCHEME-1 / D-VITE-SIGNIN-1 / D-URLS-1), M215 (the live
  acceptance — the full finding ledger F1–F13 at `m215-prove-on-odyssey/iter-01/findings.md`; F13 = the
  out-of-scope jobsimulation-service crash, off the proven journey path).
</content>
