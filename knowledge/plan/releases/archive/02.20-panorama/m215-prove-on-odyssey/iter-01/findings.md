# M215 iter-01 (tik-1) — findings from the first live bring-up on billion

**Strategy (bootstrap):** drive M215 directly (not via background sub-agents — live shared-infra work). Tik-1 goal:
a trimmed bring-up on billion (`DEMO_NO_UI=1 DEMO_NO_LOCAL_CONTENT=1`, backend + Clerkenstein only) with
`--public-host billion.taildc510.ts.net`, to de-risk clone + amd64 Go builds + migrate + Clerkenstein FAPI +
`tailscale cert`/`serve` before adding next-web (tik-2) for the actual remote login.

**Consumption:** rext @ `41a28aa` (M212–M214 + the colleague `/sign-in` 404-fix). Secrets scp'd to
`~/panorama/.agentspace/secrets` (user-approved 3.a). Layout on billion: `~/panorama/{stack-demo/rosetta-extensions,
.agentspace/secrets, .agentspace/rext.tag}`.

## Findings (real "run-it-on-a-VM" issues — the point of M215)

- **F1 — `tailscale cert` needs elevation; the M213 tooling calls it un-sudo'd.** `up-injected.sh::gen_tailscale_fapi_cert`
  runs `tailscale cert` without sudo (else falls back to mkcert = local-trust-only → a REMOTE browser sees an
  untrusted cert). Fix: a one-time **`sudo tailscale set --operator=devops`** on the VM makes `tailscale cert` +
  `tailscale serve` run as the deploy user without sudo (the deployment assumption the tooling bakes in).
  → **DOC/PREREQ finding** for `tailscale-serve.md` (VM prereq: set the tailscale operator).

- **F2 — the host needs Go for the rext orchestration tooling.** `stacksecrets`/`stacksnap`/`stackseed` are Go and run
  ON the host; a bare VM (billion) had none → secret provisioning skipped → `no usable platform .env` → abort. The
  Docker builds compile Go in-image, but the host tooling needs Go too. Fix: install **Go 1.25.12** on the VM
  (matches rext's `toolchain go1.25.12`). → **DOC/PREREQ finding** (VM prereq: Go on the host). (The stale odyssey KB
  claims VMs have Go 1.26; billion did not.)

- **F3 — `git tag --list | head -1` SIGPIPEs → 141 → `set -e` aborts the bring-up (REXT BUG, fixed).**
  `up-injected.sh:546` resolved the build tag via `git -C "$src" tag --list 'v*' --sort=-v:refname | head -1`. On a
  many-tag repo (`app` has ~337 v-tags), `head` closing the pipe SIGPIPEs `git` → pipeline exits 141 under
  `set -o pipefail` → `set -e` kills the whole bring-up (even though the tag was captured). Races on tag-count+I/O,
  so it reproduces every time on billion (full clone) but not reliably locally. **Fixed** in the authoring copy →
  `git for-each-ref --sort=-v:refname --format='%(refname:short)' --count=1 'refs/tags/v*'` (no pipe, git limits to
  one line itself). Needs a commit + a regression test (the tik's rext deliverable) + re-pin.

- **F4 — the platform compose build needs an ssh-agent (`ssh: default`) even though module pulls use the PAT.** The
  6 injected services build fine via the `GH_ACCESS_TOKEN` build-arg, but `make up`'s `docker compose build`
  (buildx bake) fails at definition-load: `invalid empty ssh agent socket: make sure SSH_AUTH_SOCK is set` — the
  platform `docker-compose.yml` build blocks declare `ssh: default` (there's an upstream
  `chore/drop-ssh-default-compose-directives` branch acknowledging it). On a bare Linux host with no ssh-agent,
  bake refuses. Fix: **start a keyless ssh-agent** (`eval "$(ssh-agent -s)"`) before the build — it satisfies the
  bake directive; the PAT still does the real private-module pulls (no key needed). → **REXT fix** (up-injected.sh
  should ensure `SSH_AUTH_SOCK` before the compose build, keyless-agent if absent) + **DOC/PREREQ**.

- **F5 (non-fatal, note) — two demopatches refused on the current `app` tag.** `app: target-role authz-skip apply
  failed/refused` + `app: ai-readiness loadMembers bound apply failed/refused` (both non-fatal — the demo works,
  just with the slower per-member Sentinel fan-out / unbounded AI-readiness hydration). Likely sha-drift: the
  demopatches' pinned pre-hash no longer matches the current `app` release. Not a Linux/remote issue per se, but
  surfaced on this fresh full-clone build → note for a demopatch re-anchor (separate from the remote-deploy story).

- **F6 — Linux bind-mount data-dir perms (Bitnami UID 1001 can't write a root-owned host dir).** `demo-1-postgresql`
  exited(1): `mkdir: cannot create directory '/bitnami/postgresql/data': Permission denied`. On macOS Docker Desktop
  the bind-mount perms are remapped; on **native Linux** the container's non-root UID can't write the host data dir
  (Docker created it as root). Fix: `sudo chmod -R 777 $STACK/data` (or chown to the Bitnami UID). → **REXT fix**
  (pre-create the data dirs with open perms on Linux before compose up) + **DOC**.

- **F7 — every seeder failed because migrate hadn't populated the schema** (root cause = F8).

- **F8 — the host needs the `atlas` migration CLI; billion had none.** `migrate-demo.sh` runs `atlas migrate apply`
  and treats failure as a non-fatal "migration warnings" → schemas created but **0 tables** → every seeder failed
  `relation public.X does not exist`. Fix: install **atlas** (`curl -sSfL release.ariga.io/atlas/atlas-linux-amd64-latest`).
  Same class as F2 (Go). → **DOC/PREREQ** (VM prereq: atlas) + consider making `migrate-demo.sh` FAIL LOUD on
  `atlas`-not-found (the "non-fatal warning" masked a total migration failure).

- **F9 — the taxonomy (42,790 public skills) is set-dressed from the SNAPSHOT CACHE, not migrations.** After migrate,
  `public.skills = 0`; the seed logs `taxonomy=skipped(cache-miss)`. billion has no `.agentspace/snapshots` cache, so
  the taxonomy + Directus content are empty (library/skills surfaces sparse). Identity heroes still seed fine (201
  users, 1 org). → for the content/skills surfaces a later tik scp's the snapshot cache to the VM (or captures). Note
  in the runbook: the VM needs the snapshot cache for full content.

## Status — TIK-1 PROVEN ✅ (2026-07-11)
After F1–F8, the trimmed backend + Clerkenstein tier is UP + healthy + migrated (91 public tables) + seeded (201
users / 1 org / 4543 rows, isolation clean) on billion, and — the make-or-break — **a REMOTE machine (kirality's
Mac) reaches it over Tailscale with a real trusted Let's Encrypt cert (verify=0, no CA install):** backend
`/api/health` → `"OK"` 200, router `/health` → 200, FAPI `/v1/client` → 200 — all `verify=0`. `tailscale serve`
fronts the plaintext services per-port; the FAPI serves its own TLS (M213 D-PROXY-2). **The M213/M214 remote-auth
foundation works end-to-end on a real Linux VM.** Findings F1–F8 are the host-prereq + rext-fix set the propagation
gate must land. Next — **tik-2:** add next-web (drop `DEMO_NO_UI`) + drive an actual browser login as a seeded hero
from the Mac.

## TIK-2 PROVEN ✅ — remote login + real journey over Tailscale (2026-07-11)
next-web + studio-desk brought up on billion (fronted by `tailscale serve` on `:13000`/`:19000`); RAM held (5.7 GB
avail). **A headless Chromium on kirality's Mac (a DIFFERENT tailnet machine) drove the full login:** opened the
cockpit (`http://billion.taildc510.ts.net:17700`, 9 heroes) → navigated hero `maya-thriving`'s [Log in as]
handshake (`…:15400/v1/client/handshake?__clerk_identity=maya-thriving&redirect_url=…:13000/profile`) → the
Clerkenstein FAPI set the session + redirected → **landed authenticated at `https://billion.taildc510.ts.net:13000/profile`,
title "Anthropos | Profile"**, with the hero's full profile rendered (work history, certifications, education,
projects — the M41 ProfileSeeder depth). **`ignoreHTTPSErrors:false` (the LE cert is genuinely trusted, no
override), 0 console errors, 0 functional request failures** (the 19 `requestfailed` are all third-party
analytics/ads — google-analytics, doubleclick, linkedin, the `/api/e` meta-pixel shim — blocked/aborted in a
headless browser, NOT app failures). This is the **M215 employee-vantage exit-gate proof**: a teammate on another
tailnet machine logs in + completes a real journey, trusted HTTPS, zero localhost ejects, assets rendering.

- **F11 (cosmetic, seed) — hero identity vs profile-name mismatch.** Logged in as `maya-thriving`; the top-nav shows
  "Maya" (org workspace) but the profile person renders as "Sven Park" (`sven.park1@northwind.com`). A seed-data
  naming inconsistency (the identity key ≠ the generated profile name), not a functional issue — login + render
  work. Note for a seed polish; unrelated to the Linux/remote story.

## Both vantages PROVEN ✅
- **Manager vantage (2026-07-11):** logged in as `dan-manager` over Tailscale → landed at
  `…:13000/enterprise/workforce?tab=skills-verification` ("Workforce Intelligence"), fully rendered with real
  seeded structural data (221 members, 445 AI sims / 67.6% pass, 47 skill paths, 224 hours, 164 certs, 16
  languages), nav shows "Dan", `hero-name-in-page=true`, 0 console errors, 0 functional req failures. The
  skills-mapped/verified funnel reads 0 (F9 — empty taxonomy). Employee (Maya) + manager (Dan) both proven.

## Cold reset-to-seed capstone (2026-07-11) — validates the fixes + surfaces F12
Synced the fixed rext (`panorama-m215`) to billion and ran a **cold** (wiped-DB) one-shot `--public-host` bring-up.
The three auto-fixes fired exactly as designed (log-confirmed): `host pre-flight OK — Go + atlas present` (F1/F2/F8),
`F4: started a keyless ssh-agent`, `F6: pre-created Linux bind-mount data dirs (chmod 777)`. No manual steps. It
then hit **F12**:

- **F12 — the demo teardown does NOT reset `tailscale serve` → a re-deploy port-conflicts (REXT gap).** `tailscale
  serve` binds the tailnet IP `:<offsetport>` (a real listener); on teardown the serve config PERSISTS (grep: serve
  is only in the up path, never reset in `rosetta-demo down`/up-injected). So a re-up's new backend container can't
  bind `0.0.0.0:18082` (overlaps the leftover `100.x:18082` serve listener) → `address already in use` → the
  backend/dependent services never start. Fix: (a) `rosetta-demo down` should `tailscale serve reset` (or per-port
  off) for the demo's ports; (b) defensively, the up-path serve step should reset the demo's ports before
  (re)configuring (idempotent re-up). Manual unblock: `tailscale serve reset`. → **REXT fix + test + runbook note**
  (added to the propagation close-gate).

## Cold reset-to-seed — PROVEN ✅ (clean run, 2026-07-11)
With serve reset (F12) + the fixed tooling, a **wiped-DB one-shot** `--public-host` bring-up ran end-to-end with NO
manual steps: the auto-fixes fired (pre-flight OK, F4 ssh-agent, F6 data-dirs), 14 containers up, `tailscale serve`
fronting 5 ports, seed 12,245 rows / 541 users / 9-hero roster, `/api/health 200`, casbin 1150. From the Mac the
fresh demo is login-ready (next-web `/` → 307 to the billion-baked FAPI sign-in; backend health 200 verify=0). **The
"reproducibly on a cold reset-to-seed" gate item is met.**

- **F13 (secondary, not Linux-specific) — `jobsimulation` container exits(1) on startup.** Its binary printed its CLI
  help + exited (got no run/serve subcommand). Off the login/profile/dashboard path (both hero journeys rendered
  fine); it would affect the AI-Simulations surface. Likely a service-command/compose or version-drift issue that
  would hit any demo, not a remote/Linux finding. → investigate separately (a demo-service fix), not part of the
  Linux/remote propagation.

## Remaining toward the FULL exit gate (later tiks)
- The **taxonomy/library/skills** surfaces (currently sparse — F9: no snapshot cache on billion → scp the
  `.agentspace/snapshots` cache or capture, then re-set-dress).
- The automated multi-journey Playwright suite pointed at the remote origin (reuse M42/M202) for a repeatable gate.
- **The propagation close-gate** (F1–F8 → tools/KB/skills) per `propagation-checklist.md`.
