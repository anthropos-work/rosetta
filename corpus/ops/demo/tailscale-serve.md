# Remote stack access over Tailscale — the `--public-host` runbook

**Make a stack reachable from another machine on your Tailscale tailnet** — run it on a Tailscale VM
(e.g. `billion.taildc510.ts.net` on the odyssey Proxmox host) and a teammate with Tailscale up browses the whole
demo end-to-end, in their own browser, over HTTPS. This is the **external-shareability** surface of v2.2
"panorama".

> **Both stack families are covered here, with OPPOSITE defaults** (v2.3's D-DESIGN-3):
> **`/demo-up` is DEFAULT-ON** (Steps 0–7 below — opt out with `--no-public-host`) and
> **`/dev-up` is OPT-IN** (**[Step 8](#step-8--the-dev-path-same-reach-opposite-default-v23-m220-s7)** — opt in
> with `--public-host auto`). The **capability ladder is one implementation shared by both**; only the default
> differs. A dev box that passes no flag makes **zero** `tailscale` calls.

> **The demo-patch mechanism is specified in [`demopatch-spec.md`](demopatch-spec.md).** It is the sanctioned **zero-platform-edit escape hatch**: patch the demo's own ephemeral clone before the image build, revert after — the canonical repos are never touched. Read it before adding or re-pinning a patch. Since M217 the gate is **self-healing**: the *anchor* is the contract, the whole-file sha is only a baseline.

> ## ⚠️ v2.3 M220 S3 — THIS IS NOW **DEFAULT-ON** (opt-out), NOT OPT-IN
>
> This page was written for v2.2, where remote reach was **opt-in** behind `--public-host`. **As of v2.3 M220 S3
> (D-DESIGN-3, superseding v2.2's D-DESIGN-1) a bare `/demo-up N` auto-discovers the host and brings the demo up
> reachable.** The flag still works (it forces a host and skips discovery); the new thing is that you no longer
> need it. Opt out with **`--no-public-host`** (or `DEMO_NO_PUBLIC_HOST=1`), which does not even probe.
>
> **Auto-discovery is capability-gated, never presence-probed** — six rungs: `tailscale` on PATH →
> `BackendState == Running` → a **dotted** `.Self.DNSName` → `MagicDNSEnabled` → no operator/sudo denial on
> `tailscale serve status` → **`tailscale cert` actually MINTS**. *"The binary exists" is not "it works."*
>
> **🔴 Any failed rung ⇒ an EMPTY host ⇒ the plain localhost demo, byte-identical to v2.2**, plus one loud line
> naming the fix. Never a *partial* public path: `SCHEME` and `BIND_HOST` share the same `-n $STACK_PUBLIC_HOST`
> predicate, so a half-satisfied public path bakes `https://` URLs against plain-HTTP listeners and the demo
> **does not load at all**. Contract + rationale: [`../safety.md`](../safety.md) **§3.5**.
>
> Everything below — the topology, the cert, the serve proxies, the walkthrough — is **unchanged and still
> correct**. Only the *trigger* changed: what you had to ask for, the bring-up now offers.

> ### The co-derivation is a NAMED invariant, not a convention (M220 harden)
>
> The sentence above — *"`SCHEME` and `BIND_HOST` share the same predicate"* — is the load-bearing safety
> argument for the whole default-on flip. It lives in **one function**, `derive_public_host_vars()` in
> `up-injected.sh`, which is where both are derived and where the dotless-pk refusal and the FAPI topology
> guard also sit. It is a function *specifically so the fence can call it*: it used to be straight-line
> top-level code, and because the lib-only test seam executes that block at **source** time — before a test
> can resolve a host — the fence **re-typed the three derivation lines inside the test**, beneath a docstring
> claiming it did not. Mutating the shipped `SCHEME` to an unconditional `https` therefore left the entire
> suite **green**: the test was comparing a copy of the predicate against itself.
>
> **If you change how `HOST` / `FAPI_HOST` / `SCHEME` / `BIND_HOST` are derived, change them inside that
> function.** A second derivation site anywhere else re-creates the exact hole, and it is invisible to a
> green suite. RED-proven by mutants D5/D6 in `stack-core/tests/test_m220_mutation_battery.py`.

> ### The ladder ALWAYS exits 0 — which is why a broken `tailscale` cannot test the fallback (M220 harden)
>
> `tailscale_autohost.py`'s contract is **exit 0, always**: a failed rung is the *documented fallback*, not an
> error, and the module converts every internal crash into the empty host. A direct consequence, and it is not
> obvious: **a broken/absent `tailscale` binary can never make the ladder process exit non-zero.** So it can
> never exercise the `|| true` on the caller's side.
>
> The `|| true` in `resolve_public_host` / `resolve_dev_public_host` is still **load-bearing** — with
> `set -euo pipefail` and a bare `pub="$(resolve_… )"` assignment (which, unlike `local pub=$(…)`, does *not*
> mask the exit status), a non-zero collapses the whole bring-up. But the only thing that can produce one is
> the **ladder process itself** dying: a missing or broken `python3`, an OOM, or an `ImportError`/`SyntaxError`
> raised at *import* time — i.e. before `main()`'s `try/except` exists to catch it.
>
> **A test that stubs a failing `tailscale` is therefore testing the rung ladder, not the fallback.** To test
> the fallback you must break the *interpreter or the module*. The dev path shipped with only the former and
> the `|| true` went unexercised (mutant V4 survived the whole suite).

Historically (v2.2) this was **opt-in and default-off**: a plain `/demo-up N` was byte-identical to a localhost
demo, and remote reach was requested explicitly with **one flag** — `--public-host <magicdns>`. In both eras
**Tailscale itself is the access control** (only tailnet members can reach the host; there is no public internet
listener) — and note that this was *never* what gated the container ports, which have always been published on
`0.0.0.0` (see the correction at the bottom of this page, and `safety.md` §3.1).

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
> throughout — tooling + docs + the opt-in flag only (**three** platform-family files ride the **existing** rext
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
| 2 | **Go 1.25.x** (matches rext's `toolchain go1.25.12`) | the host rext tooling is Go; without it secret provisioning is skipped → `no usable platform .env` → abort | `curl -sSfL https://go.dev/dl/go1.25.12.linux-amd64.tar.gz \| sudo tar -C /usr/local -xz` (the standard install; `/usr/local/go/bin` is added to `PATH` by the login profile). **Before installing anything, see F2b below — a "Go NOT on PATH" pre-flight failure is usually a LOGIN-SHELL problem, not a missing Go** | **F2**, **F2b** |
| 3 | **atlas CLI** | `migrate-demo.sh` runs `atlas migrate apply`; without it the schema is created with **0 tables** and every seeder fails `relation public.X does not exist` | `curl -sSfL https://release.ariga.io/atlas/atlas-linux-amd64-latest -o atlas && sudo install -m755 atlas /usr/local/bin` | **F8** |
| 4 | **Tailscale operator** | so the bring-up's un-sudo'd `tailscale cert` / `tailscale serve` run as the deploy user; else the cert falls back to `mkcert` = **local-trust-only** and a remote browser distrusts it | one-time `sudo tailscale set --operator=<deploy-user>` | **F1** |
| 5 | **An ssh-agent** | the platform compose declares `ssh: default`, so `buildx bake` needs `SSH_AUTH_SOCK` at definition-load — even though the private-module pulls use the `GH_PAT`, not the agent. A **keyless** agent suffices | the bring-up **auto-starts one if absent** (`eval "$(ssh-agent -s)"`); no key needed | **F4** |
| 6 | **The snapshot cache** (content surfaces only) | the taxonomy (~42,790 public skills) + Directus content are **set-dressed from the snapshot cache**, not migrations. Without `.agentspace/snapshots` on the VM, `public.skills=0` and the library/skills surfaces are sparse (identity/profile/dashboard still render fully) | `scp` the `.agentspace/snapshots` cache to the VM, or capture per [`../snapshot-cold-start.md`](../snapshot-cold-start.md) | **F9** |

The canonical prereq list + these install commands also live in
[`../setup_guide.md`](../setup_guide.md) §"Linux host prerequisites (for a remote/VM demo over Tailscale)".

> ### 🔴 Run every remote bring-up through a LOGIN shell — or this table will lie to you (F2b, M236 iter-03)
>
> ```bash
> ssh <host> 'bash -lc "<the bring-up command>"'     # -l = login shell; sources the profile
> ```
>
> A non-interactive `ssh host 'cmd'` sources **no login profile**, so everything the profile puts on `PATH` —
> including `/usr/local/go/bin` — is simply absent. The host pre-flight then reports **"Go NOT on PATH …
> install Go 1.25.x"** on a box where Go *is* installed at the exact pinned version.
>
> **Why this is a trap and not a footnote:** `atlas` (prereq #3) lives in `/usr/local/bin`, which *is* on the
> default non-login `PATH`, so it **passes**. One prereq green and one red reads as *prereq-specific* ("Go is
> missing") rather than *shell-specific* ("nothing from the profile is on `PATH`") — so a **shell** problem
> presents as a **prereq** problem, and row #2's install command is exactly the wrong next step. **Disprove it
> in one command before installing anything:**
>
> ```bash
> ssh <host> 'bash -lc "go version"'                 # green here + red in pre-flight ⇒ PATH, not prereqs
> ```
>
> Full write-up, including the general rule (*a tool's absence and a tool's invisibility produce identical
> output, and only one is fixed by installing anything*):
> [`../verification.md` § Drive every remote bring-up through a LOGIN shell](../verification.md#pre-flight-rung-zero--can-the-host-even-obtain-the-thing-under-test-v25-m236).

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

**The cockpit login (the interactive proof).** Open the presenter cockpit at **`https://$HOST:17700`**
(`7700+off`) — on a `--public-host` demo it is fronted by `tailscale serve` behind the **same trusted MagicDNS
cert** as every other browser-facing port (M220 S4; see the correction at § "the last plain-HTTP surface is
gone" below). It lists the seeded heroes; each **[Log in as]** is a link to the FAPI handshake:

> ⚠️ **This step previously read `http://$HOST:17700` — "deliberately *not* fronted by `tailscale serve`".**
> That was true up to v2.3 M220 S3 and is **false now**: `gen_tailscale_serve.py` carries `('cockpit', 7700)`
> on its own `DEMO_STORIES`-gated axis and `up-injected.sh` fronts it. Following the old text verbatim points
> the operator at the wrong scheme on the demo's **entry point** — the one page a presenter actually opens.
> On a **localhost** demo (no `--public-host`) the cockpit is plain `http://localhost:17700`, as it always was.

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

# Step 8 — The DEV path: same reach, opposite default (v2.3 M220 S7)

Everything above is the **demo** path, where remote reach is **default-on**. A **dev** stack can be made
reachable too — but you must **ask**:

```bash
DEV=stack-dev/rosetta-extensions/dev-stack

"$DEV/dev-stack" up 2                              # ← nothing happens. No tailscale, no probe, no serve.
"$DEV/dev-stack" up 2 --public-host auto           # ← opt IN: walk the ladder, adopt this box's MagicDNS host
"$DEV/dev-stack" up 2 --public-host box.tail.ts.net  # ← ...or name it outright (no probe; you have spoken)
DEV_PUBLIC_HOST=auto "$DEV/dev-stack" up 2         # ← the env form of the same opt-in

"$DEV/dev-stack" down 2                            # ← clears this stack's serve ports, exactly like /demo-down
```

**Why the asymmetry** — it is **v2.3's D-DESIGN-3**, in the user's words: *"opt-out at build time for
`demo-up`, **opt-in** at build time for `stack up`"*. A demo exists to be shown to someone else. A dev box does
not, and — unlike a demo — **its content is not guaranteed synthetic** (a dev stack reads content live from prod
by default). See [`../safety.md`](../safety.md) **§3.5.3**.

| | remote reach | escape hatch | env form |
|---|---|---|---|
| **`/demo-up N`** | **DEFAULT-ON** | `--no-public-host` | `DEMO_NO_PUBLIC_HOST=1` |
| **`/dev-up N`** | **OFF** | `--public-host auto` \| `<fqdn>` | `DEV_PUBLIC_HOST` |

**What is shared, and what differs:**

- **The ladder is the SAME code** — `demo-stack/tailscale_autohost.py`, all six rungs, reused cross-section
  (not reimplemented). Same rungs, same order, same verdict, same fix-lines; only the words on stderr change
  (`dev-up:` instead of `demo-up:`). **The one difference is the default, and it lives in the caller.**
- **The fallback is the same and just as hard.** Any failed rung ⇒ an **empty host** ⇒ the localhost dev stack
  that has always worked, plus one loud line naming the fix. Never a half-satisfied public path.
- **Pass no flag ⇒ byte-identical to before the feature existed.** Zero `tailscale` invocations — it does not
  probe and decline, it does not look. (Fenced with a tripwire stub that fails the test if `tailscale` is
  called at all.)
- **`DEV_PUBLIC_HOST`, not `STACK_PUBLIC_HOST`.** `up-injected.sh` *exports* the latter, so an inherited value
  would otherwise flip a dev stack public with no flag on the command line. Dev has its own namespace.
- **Only the ports your `--profile` actually publishes are fronted** (default `graphql` ⇒ backend API + Cosmo
  GraphQL). The demo's fixed registry does not apply: `tailscale serve` **binds** what it fronts, so fronting a
  port with no listener would hold it against the next bring-up.
- **No cockpit, no Clerkenstein.** A dev stack authenticates against **real Clerk** and has no presenter
  launcher. §3.2's *"unauthenticated, authz-weakened build"* is a description of a **demo**, not of this.

> **The dev up-path pre-reset runs BEFORE `docker compose up`** — which the demo path (ADV-1) cannot do, since
> it resolves its host later. So a stale serve listener from a previous dev stack is cleared *before* the
> containers try to bind, rather than after.

---

# Part 2 — How it works

## The topology — HTTPS everywhere, one MagicDNS host, per offset port

A demo runs its browser-facing services on **offset ports** (`base + N*10000`): next-web `3000+off`, **the
apps/hiring 2nd app `3001+off`** (the TOK-02 two-app hiring demo — v2.4 "casting call"), the Cosmo GraphQL router
`5050+off`, the backend REST `8082+off`, studio-desk `9000+off`, ant-academy `3077+off`, and the fake Clerk FAPI
`5400+off`. Under `--public-host`, each is reached over **HTTPS on the MagicDNS host at the same offset port**:

```
teammate's browser ── https://billion.taildc510.ts.net:13000 ──▶  tailscale serve ──▶ http://127.0.0.1:13000  (next-web / apps/web)
                   ── https://billion.taildc510.ts.net:13001 ──▶  tailscale serve ──▶ http://127.0.0.1:13001  (apps/hiring — the recruiter's 2nd app)
                   ── https://billion.taildc510.ts.net:15050 ──▶  tailscale serve ──▶ http://127.0.0.1:15050  (cosmo)
                   ── https://billion.taildc510.ts.net:18082 ──▶  tailscale serve ──▶ http://127.0.0.1:18082  (backend)
                   ── https://billion.taildc510.ts.net:19000 ──▶  tailscale serve ──▶ http://127.0.0.1:19000  (studio-desk)
                   ── https://billion.taildc510.ts.net:13077 ──▶  tailscale serve ──▶ http://127.0.0.1:13077  (ant-academy, native)
                   ── https://billion.taildc510.ts.net:15400 ──▶  fake-FAPI's OWN TLS (tailscale cert)          (Clerkenstein)
```

> **The apps/hiring port (`3001+off`) joined the serve front at M226 "opening night" (Finding-1).** It was added
> to `gen_tailscale_serve.py`'s `UI_BROWSER_FACING` registry (same UI-tier lifecycle as next-web, dropped under
> `--no-ui`). Before M226 the hiring 2nd app (added at M224) was **reachable only on localhost** — a recruiter's
> cockpit CTA lands on `https://<host>:3001+off`, which had **no HTTPS listener over the tailnet**, so the
> recruiter vantage worked on the dev box yet was **dead cross-machine**. The billion proof surfaced it (the
> M215/M221 "last breakage is cross-machine" lesson); the default demo bring-up now fronts it.

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
| **ant-academy** (`NEXT_PUBLIC_STUDIO_URL`, `next dev` bind, `allowedDevOrigins`) | `http://localhost:…`, **`127.0.0.1:13077` bind** (`-H 127.0.0.1` on the localhost path since v2.3 M221 F-M220-5 — was `*:13077`, because `next dev`'s own default is `0.0.0.0`), hardcoded origins | `https://$HOST:…`, `-H 0.0.0.0`, MagicDNS host admitted | `ant-academy.sh` + the `ant-academy-dev-origins` patch |
| **FAPI cert** | `mkcert`/openssl (local trust) | `tailscale cert` (tailnet-wide trust) | `up-injected.sh` (M213) |
| **Cockpit bind** | `127.0.0.1` (loopback, `cockpit.py --host` default) | `0.0.0.0` | `up-injected.sh` (M212) |
| **ant-academy bind** | **`127.0.0.1`** (loopback, `-H 127.0.0.1` — landed v2.3 M221 F-M220-5; was `*:13077` from `next dev`'s own `0.0.0.0` default) | `0.0.0.0` (`-H 0.0.0.0`) | `ant-academy.sh` (M212; loopback default M221) |
| **Registry** | no interaction | records `external_host` for `/stack-list` | `up-injected.sh` (M212) |

**Mixed-content clean.** With HTTPS-everywhere, no browser-facing call resolves to plain `http://` — the scheme
flip covers every baked endpoint, redirect, and cross-surface link; the asset plane is already prod-HTTPS.

> **v2.3 M220 S4 — the last plain-HTTP surface is gone.** This paragraph used to end: *"The **one** deliberate
> plain-http surface is the **presenter cockpit's own page** (`7700+off`): it is not in `tailscale serve`'s front
> list … fronting the cockpit too is a live-acceptance polish left as an accepted future enhancement."*
>
> **It is no longer a polish, and it is no longer excluded.** `gen_tailscale_serve.py` now carries
> `('cockpit', 7700)` on its **own axis** (gated on `DEMO_STORIES`, *not* on `--no-ui` — the cockpit runs on a
> `--no-ui` stories demo, so filing it under the UI tier would have left a **live** cockpit unfronted), and
> `up-injected.sh` fronts it with the same trusted MagicDNS cert as every other browser-facing port.
>
> **Why it mattered enough to fix now:** S3 made remote reach **default-on**. Leaving the cockpit on plain HTTP
> then becomes the worst possible combination — the demo's **entry point**, the *one* page a presenter actually
> opens, is the *one* page not behind the trusted cert.
>
> **Ordering matters, and it is not cosmetic.** `tailscale serve` binds the tailnet IP `:<port>` as a **real
> listener**, so fronting `:7700` *before* the cockpit binds makes the cockpit's own bind fail `EADDRINUSE` —
> you would "fix" its exposure by killing it (the same contention M215 F12 hit for `ant-academy`). So the
> bring-up's **first** serve apply passes `--no-cockpit`, and a **second**, idempotent apply fronts the cockpit
> only after its `/healthz` answers. Bind first, front second. The **reset** plan, by contrast, always includes
> `:7700` — otherwise a re-up over a stale serve config finds the port held and the cockpit cannot start at all.
>
> **This is transport, not authentication** — do not over-read it. The cockpit remains a one-click,
> password-free *"become any seeded hero"* launcher; it is now behind the tailnet's TLS + device mesh rather than
> in cleartext. See [`../safety.md`](../safety.md) **§3.5.2**.

## The tailscale-cert FAPI (the Clerk-free login over a real cert)

The Clerk-free browser login routes through Clerkenstein's **fake FAPI** over HTTPS (clerk-js always reaches the
FAPI over `https://`, derived from the publishable key). For a remote demo the FAPI cert is minted with
**`tailscale cert <magicdns>`** — a real Let's Encrypt MagicDNS cert **trusted tailnet-wide with no per-machine
CA install**, exactly what `mkcert` cannot give a *remote* browser. The fake-FAPI serves its **own** TLS with
that cert on `5400+off` (so it is excluded from the `tailscale serve` proxy — double-fronting would double-TLS).
The consumer mount is path-only (`<stack>/certs/fapi.{crt,key}`), so it is a drop-in at the same paths as the
local mkcert/openssl cert. Falls back to the local mkcert/openssl path (local trust only) if `tailscaled` isn't
up **or** the tailscale operator isn't set (F1 — set it, Step 0 #4). The LE cert is **90-day**, so a long-lived
stack still needs a renew-then-reload step eventually.

> **CORRECTION (M220 S3, measured).** This paragraph used to add: *"`tailscale cert` re-issues on re-run"*.
> **It does not.** Two back-to-back mints on `billion` (2026-07-14) returned the **identical certificate serial**
> (`05777C48…`) in **0.01 s** each, with **zero** new ACME orders in `tailscaled`'s journal. tailscaled serves
> the cert from its own cache and only re-orders near expiry.
>
> **Why this was load-bearing, not trivia.** M220 S3 makes remote reach **default-on**, and its capability
> ladder **mints a cert as rung 6 — on every bring-up**. If the old claim had been true, default-on would burn a
> Let's Encrypt **duplicate-certificate** slot per `demo-up` — and since `ts.net` is a **PSL entry**, that bucket
> is **per-tailnet**, shared by every box on it. A mint failure then silently degrades to a local-trust cert a
> remote browser rejects. The flip was gated on **settling this empirically** rather than trusting the sentence.

Full cert story + caveats:
[`recipe-browser-login.md`](recipe-browser-login.md) §B step 2 and
[`../../services/clerkenstein.md`](../../services/clerkenstein.md).

## The patch tail — THREE platform-family files, all via the existing sha-pinned mechanism

> ⚠️ **M219 close: this section listed only TWO, and omitted the one patch that exists *because of* `--public-host`.**
> **`next-web-ssr-graphql-origin`** (M218) is the fix for the **38-second login**: the SSR pass fetched the public
> MagicDNS origin **from inside the container**, where the tailnet IP **blackholes** (ts-input drops the SYN-ACK on
> the docker bridge) → ~37 s per authenticated render. That defect **only manifests on a `--public-host` demo** —
> i.e. exactly the flow this runbook documents — and the remote-access runbook never mentioned it. See
> [`demopatch-spec.md` § 5](demopatch-spec.md) for the row.

**Three** files in the platform **family** aren't reachable by the pure config/env layer, so they ride the
**existing rext sha-pinned patch mechanism** (drift-refuse, single-occurrence anchor, idempotent, non-fatal)
applied to the demo's **ephemeral clone** — **never a checked-in platform clone, never a canonical repo edit**:

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

3. **next-web SSR GraphQL origin (`next-web-ssr-graphql-origin`, M218) — the one that exists *because of*
   `--public-host`.** `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` is a single build-time constant serving two consumers
   with **incompatible** reachability: the **browser** needs the public origin
   (`https://<public-host>:15050+off/graphql`), the **SSR pass** needs the container origin
   (`http://graphql:8080/graphql`). Because `NEXT_PUBLIC_*` is build-inlined into the *server* bundle too, the
   SSR pass fetched the public URL **from inside the container**, where the tailnet IP **blackholes** (ts-input
   drops the SYN-ACK on the docker bridge) → undici's 10 s connect timeout × 3 attempts + 6 s backoff
   ≈ **37.5 s per authenticated render**, on both vantages (they block on the same shared authenticated layout).
   Measured on `billion`: employee p95 **39.45 s**, manager **38.30 s**, against a 5 s gate. The patch gives the
   server-side client its own container origin. **This defect manifests *only* on a `--public-host` demo** — the
   exact flow this runbook documents. Manifest:
   `rosetta-extensions/demo-stack/patches/next-web-ssr-graphql-origin/`; budget + attribution model:
   [`latency-budget.md`](latency-budget.md).

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
| **F2** — the host needs Go for the rext tooling | no Go → secret provisioning skipped → `no usable platform .env` → abort | **pre-flight + prereq** — install Go 1.25.x (Step 0 #2). ⚠️ **Confirm it is genuinely missing first — see F2b**; the same symptom is produced by a non-login shell |
| **F2b** — a remote bring-up run over a **non-login** shell reports a **false "Go NOT on PATH"** | `ssh host 'cmd'` sources no profile, so `/usr/local/go/bin` is absent while `atlas` (in `/usr/local/bin`) passes — a **shell** problem presenting as a **prereq** problem, and F2's remedy text pushes the operator into installing a second Go | **protocol (documented)** — always `ssh <host> 'bash -lc "…"'`; disprove with `ssh <host> 'bash -lc "go version"'` before believing any remote "prereq missing" verdict. See Step 0's login-shell callout + [`../verification.md` § PRE-FLIGHT RUNG ZERO](../verification.md#pre-flight-rung-zero--can-the-host-even-obtain-the-thing-under-test-v25-m236) *(handler `DOC-M236-iterTBD-protocol-backfill`)* |
| **F3** — `git tag --list \| head -1` SIGPIPE → 141 → `set -e` aborts | reproduces on a many-tag repo (`app` ~337 v-tags) | **rext fix (shipped)** — pipe-less `git for-each-ref --count=1` in `up-injected.sh` |
| **F4** — buildx bake needs `SSH_AUTH_SOCK` (`ssh: default`) even though pulls use the PAT | a bare host with no ssh-agent fails at definition-load | **auto-handled** — the bring-up starts a **keyless** ssh-agent when absent (the PAT still does the real pulls) |
| **F5** — two `app` demopatches refused (target-role authz-skip, ai-readiness loadMembers) | sha-drift on the current `app` tag; **non-fatal** (demo works, slower per-member fan-out) | ✅ **RESOLVED (M217 self-healing anchor gate).** *The anchor is the contract; the whole-file sha is only a baseline* — a drifted sha with an intact anchor now **self-heals** and applies (`demopatch-spec.md` §6). **Do not chase a re-pin; it no longer exists.** M219 (F-7) further confirmed the `loadmembers` patch is not dead. |
| **F6** — Linux bind-mount data-dir perms (Bitnami UID 1001 can't write a root-owned host dir) | `mkdir: /bitnami/postgresql/data: Permission denied`; macOS Docker Desktop remaps, native Linux does not | **auto-handled** — the bring-up pre-creates the data dirs writable (manual fix was `sudo chmod -R 777 $STACK/data`) |
| **F7/F8** — the host needs the `atlas` CLI; without it `migrate` creates 0 tables → every seeder fails | `migrate-demo.sh` treated `atlas`-missing as a non-fatal warning that masked a total migration failure | **pre-flight + prereq** — install atlas (Step 0 #3); the bring-up now fails loud on `atlas`-not-found |
| **F9** — the taxonomy is set-dressed from the snapshot cache, not migrations | no `.agentspace/snapshots` on the VM → `public.skills=0`, sparse library/skills surfaces | **prereq (optional)** — scp/capture the cache (Step 4); identity/profile/dashboard work without it |
| **F11** — hero identity vs profile-name mismatch (cosmetic) | e.g. logged in as `maya-thriving`, profile person renders as a different generated name | **known issue (seed polish)** — login + render work; unrelated to the remote story |
| **F12** — the teardown didn't reset `tailscale serve` → a re-deploy port-conflicts | `serve` binds the tailnet IP `:<offsetport>` as a listener that persists past `compose down`; the next deploy fails `address already in use` (surfaced on the billion cold reset-to-seed) | **rext fix (shipped)** — `/demo-down` resets THIS demo's serve ports (per-port `--https=<port> off`, offset-scoped), + a defensive up-path pre-reset (idempotent re-up); non-fatal, no-op on localhost. Manual by-hand-teardown unblock: `tailscale serve reset` (Step 7) |

> **Numbering note.** The ledger skips **F10** (unused). **F13** — a jobsimulation-service startup crash — is
> **off the proven journey path** and out of this runbook's host-deploy scope (it would hit any demo, remote or
> local); it is recorded in the milestone findings ledger + routed to standing backlog, not baked here.

## Safety framing

> ### 🔴 CORRECTION (v2.3 "cue to cue", M220) — this section used to state the opposite, and it was FALSE
>
> Until M220 this section read (quoted as a historical artifact — **do not treat as current**):
>
> ```text
> RETRACTED — FALSE (shipped v2.2, corrected v2.3 M220):
>   "Opt-in, default-off. … A bare `/demo-up N` binds loopback only and is byte-identical to today."
>   "…no open 0.0.0.0-on-the-LAN surprise beyond the tailnet. Binding `0.0.0.0` is gated on the knob
>    precisely so it is never ambient."
> ```
>
> **Both sentences were false.** `stack-injection/gen_injected_override.py` emits **every** published port as
> a bare `"<hostport>:<target>"` pair — with **no `127.0.0.1` prefix** — at all three emitters (`directus_lines`,
> `frontend_lines`, `build_lines`). **Docker's default bind for a bare `host:container` mapping is `0.0.0.0`.**
>
> The doc had even said so itself, 200 lines above (§ "Why per-port `serve`…"): *"`docker-proxy` binds the demo's
> offset ports on **`0.0.0.0`**."* The reassuring half is the one people quoted.
>
> **A shipped safety doc that understates real exposure is worse than no doc** — no doc at least prompts you to
> go and look. The correction below is now fenced by `stack-injection/exposure_claim_guard.py`, which derives the
> bind by *running* the emitters and fails if any doc denies it.

- **⚠️ CONTAINER PORTS ARE PUBLISHED ON ALL INTERFACES — on EVERY `demo-up`, TODAY, flag or no flag.** Every
  demo container's offset port is bound on **`0.0.0.0`**, with **or without** `--public-host`. This is **not**
  introduced by remote access; it has been true of every demo since the injected override existed. Anyone who can
  route to the host's IP can reach the demo's ports directly, bypassing `tailscale serve` entirely.
  - **On Linux this bypasses your host firewall.** Docker installs its own iptables rules in the `DOCKER` chain,
    which are consulted *before* `ufw`/`firewalld`'s. A `ufw deny` on the port does **not** block it.
  - **What `--public-host` actually adds** is *not* the exposure — it is the **trusted-HTTPS origin** (a real
    `tailscale cert`) and the per-port `tailscale serve` proxy that make the demo *usable* by a tailnet peer. The
    ports were already reachable; the flag makes them *browsable* (Clerk needs a secure context).
  - **Therefore the exposure delta of a default-on remote flip is far smaller than this doc used to imply** — the
    LAN/tailnet-IP exposure is already there. See [`../safety.md`](../safety.md) **Part 3 — the exposure side**
    for the full contract, and for what a demo actually *is* (an unauthenticated, authz-weakened build).
- **`BIND_HOST` gates only the two HOST-NATIVE servers** — the presenter cockpit and ant-academy, which run as
  plain host processes, not containers. `BIND_HOST` is `""` by default and `0.0.0.0` under a public host, and it
  **does not touch a single container**. But `""` means *"pass no `-H` and let each server keep its own default"*,
  and the two servers *used to* differ: the **cockpit** (`cockpit.py --host`) defaults to **`127.0.0.1`**
  (loopback), while **ant-academy** (`next dev`) has an OWN default of **`0.0.0.0`** — so until M221 **the academy
  was world-published on `*:13077` on every demo, flag or no flag**, exactly like the containers. **✅ LANDED v2.3
  M221 (F-M220-5):** `ant-academy.sh` now passes **`-H 127.0.0.1`** on the localhost path (`-H 0.0.0.0` only under
  a public host), so **both** host-native servers now bind loopback by default on a localhost demo (the
  *"loopback by default"* framing is no longer cockpit-only — see [`../safety.md`](../safety.md) Part 3). The demo
  **container** ports remain `0.0.0.0` by design — that half is unchanged.
- **Tailscale is the access control for the *published origin*.** The MagicDNS host + `tailscale serve` listeners
  are reachable **only** to members of your tailnet: no public-internet listener, no port-forward. This is a real
  guarantee — it is just **narrower** than "nothing else is exposed". The teammate's client must keep Tailscale
  **MagicDNS on** (the default) for the `<magicdns>` name to resolve.
- **Zero platform-repo edits.** The whole surface is rext tooling + this doc + the opt-in flag. The **three**
  platform-family patches (§"The patch tail") touch only the demo's **ephemeral** clone via the sha-pinned
  mechanism (drift-refuse fails loud on an upstream change; reverted on teardown), never a canonical repo.
- **The demo's data-isolation guarantees are unchanged.** Remote reach changes the *origin/scheme*, not the data
  plane: the tenant-data firewall, the per-stack isolated Postgres, and the never-write-prod boundary all hold
  exactly as documented in [`../safety.md`](../safety.md) (Parts 1 and 2). **This is the load-bearing mitigation:
  the ports are open, and there is nothing real behind them.**
  > 🔴 **Except on a CONTENT-STORY demo (v2.5) — where that last sentence is false.** A content-story demo
  > carries the copied, best-effort-scrubbed free-text of **real production sessions**; there *is* something
  > real behind the ports. The never-write-prod boundary is untouched, but *"nothing behind the door"* is not
  > available as a mitigation there, and the **VPN/tailnet scope becomes THE control rather than a comfort** —
  > which is precisely what this runbook provides. See [`../safety.md` §3.8](../safety.md) (the exception) and
  > **§3.3.1** (what carries the exposure argument once it applies). **Do not expose a content-story demo
  > outside a tailnet/VPN.**

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
  `urls.ts` residual decision), **and the content-stories `(session × action)` LANDS sweep (v2.5 M236)**, which
  is driven ACROSS the tailnet: the harness runs locally and only the stack is remote
  (`run-content-stories.sh <N> --host <magicdns>`). Two host-path rules M236 established — the fake FAPI is
  **always** `https` even when the app origin is not, and a `--public-host` stack must be measured with
  `LATENCY_SCHEME=https` or the latency runner grades the wrong origin.
- **v2.5 M236 — the second live proof on `billion`.** The whole content-vantage feature (both cockpit tabs,
  29/29 landable pairs, 65 academy cards, hero p95 1.22 s / 1.51 s) was proven end-to-end from a second
  tailnet machine on a cold reset-to-seed. The remote-reach recipe in this doc is the one it used, unchanged.
- Design decisions: `knowledge/plan/releases/archive/02.20-panorama/` — M212 (the knob), M213 (TLS/proxy/pk, D-PROXY-2 /
  M213-D-SCHEME-1), M214 (origins & links, M214-D-SCHEME-1 / D-VITE-SIGNIN-1 / D-URLS-1), M215 (the live
  acceptance — the full finding ledger F1–F13 at `m215-prove-on-odyssey/iter-01/findings.md`; F13 = the
  out-of-scope jobsimulation-service crash, off the proven journey path).
</content>
