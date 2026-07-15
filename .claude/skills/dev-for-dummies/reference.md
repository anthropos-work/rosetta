# dev-for-dummies — technical reference

Copy-pasteable recipes for the technical phases of [`SKILL.md`](SKILL.md). Read the relevant section before
running a phase. All examples use **`demo-1`** (N=1, offset `+10000`) and the host
`calypsostaging.taildc510.ts.net` (this box). Substitute your real N and FQDN.

Shell shorthands used below:
```bash
N=1; OFF=$((N*10000))                               # port offset for demo-N
HOST=$(tailscale status --json | jq -r '.Self.DNSName' | sed 's/\.$//')   # calypsostaging.taildc510.ts.net
PLAT=stack-demo/platform                            # compose base (demo builds from stack-demo's OWN clones)
STACK=stack-demo/rosetta-extensions/demo-stack/stacks/demo-$N   # per-stack injected dir (holds .env.demo-$N)
DC="docker compose -p demo-$N -f $PLAT/docker-compose.yml -f $STACK/docker-compose.injected.yml"
```

> **Fatal vs non-fatal (apply everywhere).** A step tagged **FATAL** stops the run with one plain-English line
> ("I couldn't X, so I've stopped — here's what to do"). A step tagged **NON-FATAL** warns in plain English and
> continues. Never report "all set" if a FATAL step failed, and never abort a good stack over a NON-FATAL one.

---

## Port map — every host port is `base + N×10000`

N=0 is the main dev stack (base ports). `demo-N` / `dev-N` add `N×10000`.

| Service (compose name) | Base host port | demo-1 | Who talks to it |
|---|---|---|---|
| postgresql | 5432 | 15432 | services (+ you, read-only) |
| redis | 6379 | 16379 | services |
| **next-web-app** (frontend) | **3000** (hiring 3001) | **13000** | **browser** ← smooth target |
| **studio-desk** | **9000** | **19000** | **browser** ← smooth target |
| **ant-academy** (native) | **3077** | **13077** | **browser** ← smooth target |
| cosmo / graphql router | 5050 | 15050 | browser + subgraph fan-out |
| **backend (`app`)** REST | **8082** (RPC 8081/8083) | **18082** | router, other services |
| cms | 8090 (RPC 8091) | 18090 | router, other services |
| jobsimulation | 8400 (RPC 8401) | 18400 | router, other services |
| skillpath | 8100 (RPC 8101) | 18100 | router, other services |
| sentinel | 8087 | 18087 | services (authz) |
| presenter cockpit | 7700 | 17700 | browser (plain HTTP — see caveat) |
| directus (only `--local-content`) | 8055 | 18055 | cms |
| fake-FAPI (Clerkenstein) | 5400 | 15400 | browser (own TLS) |

> **Frontend/UI targets are the smooth path** — only the *browser* talks to them, so a native process on the
> host serves them directly. **Backend targets** (app/cms/jobsimulation/skillpath) are consumed by *other
> containers* (the router) by Docker service name, which a host-native process can't provide — see § *Backend
> targets* for the two caveats (infra endpoints + router federation).

**Valid TARGET repos** (must exist as `stack-demo/<repo>`):
`app cms jobsimulation skillpath next-web-app studio-desk ant-academy messenger storage sentinel roadrunner
graphql-wundergraph`. **`hiring` is NOT a repo** — it's `apps/hiring` inside `next-web-app` (run `pnpm
dev:hiring` on `3001+OFF`). Always `test -d stack-demo/<repo>` first; if it fails, reject with this list
(FATAL for that target).

---

## Discovering the public host (this box's MagicDNS FQDN)

Not auto-derived by the tooling — you supply it. Discover it at runtime:
```bash
HOST=$(tailscale status --json | jq -r '.Self.DNSName' | sed 's/\.$//')   # → calypsostaging.taildc510.ts.net
HOST=${HOST:-$(tailscale status | awk 'NR==1{print $3}')}                 # fallback if jq is unavailable
```
Must be a **dotted MagicDNS FQDN** (clerk-js needs a dotted pk host + a secure/https context). A bare name is
refused. If tailscale is down / no FQDN, ask the user or fall back to a `localhost` demo (warn: no remote access).

---

## Host prereqs (one-time)

The bring-up pre-flights / fails loud on most of these; confirm before Phases 4–5.

- **Node ≥ 24 on PATH (FATAL if missing).** The frontend targets need node 24 (`next-web-app` pins
  `engines.node >=24`). **Do not assume `nvm`** — on a fresh box it may be absent and the system node may be
  older (this box ships node 20; nvm is not installed). Check `node -v` in the **same shell tmux will use**
  (`bash -lc 'node -v'`, which loads the user's version manager if they have one). If it's `< 24`, STOP and tell
  the user in one line to put a node ≥ 24 on PATH (their version manager, e.g. `nvm install 24 && nvm alias
  default 24`, or a system install), then resume. Never let it silently degrade to an old node.
- **pnpm** on PATH (the frontend uses pnpm; `npm`/`yarn` are blocked in the repo). `corepack enable` if needed.
- **`--public-host` only:** **Tailscale operator (F1)** — `sudo tailscale set --operator=$USER` so the un-sudo'd
  `tailscale cert`/`serve` mint a **trusted** Let's Encrypt cert (without it the cert silently falls back to
  local-trust-only mkcert and a *remote* browser sees it untrusted). Verify: `tailscale cert $HOST` works
  without sudo. Plus **Go 1.25.x**, **atlas CLI**, a keyless **ssh-agent**, and the `.agentspace/snapshots`
  cache — already present on this box.

---

## Run a target live — **frontend** (next-web-app) — native `next dev` + HOT-RELOAD over the **tailnet HTTPS URL**

> **CORRECTED 2026-07-14 (proven — remote tailnet access + hot-reload, no SSH tunnel).** Serve the dev server over
> **HTTPS directly** (`next dev --experimental-https` with the demo's **Tailscale cert**), bound to the tailnet,
> and do **NOT** put `tailscale serve` in front of it. It's then reachable at the **same `https://<host>:<offset>`
> URL as the rest of the demo**. Why direct-HTTPS: @clerk/nextjs's middleware rewrites every request to a
> same-origin URL; serving HTTPS directly keeps the origin consistent (`https://<host>:<offset>`), so the rewrite
> stays **relative → resolved internally**. `tailscale serve` (HTTPS→plain-HTTP) creates a host/proto mismatch that
> makes Next **self-proxy** the rewrite into a 500 loop (`Failed to proxy … wrong version number`). *(Pure on-box
> dev also works with `-H localhost` over `http://localhost:<offset>` — a secure context, loop-free — but that
> isn't tailnet-reachable; use the direct-HTTPS form to match the rest of the demo.)*

```bash
# 0. PREREQ (FATAL): node >=24 on PATH in a login shell (see Host prereqs).
bash -lc 'v=$(node -v 2>/dev/null|sed s/v//); [ "${v%%.*}" -ge 24 ] 2>/dev/null' || { echo "STOP: node>=24"; exit 3; }

# 1. Worktree + branch (never edit stack-demo/next-web-app directly). NEW=-b; RESUME drops -b.
git -C stack-demo/next-web-app worktree add -b feat/<name> ../.worktrees/next-web-app-feat-<name>
WT=stack-demo/.worktrees/next-web-app-feat-<name>

# 2. Capture the CONTAINER's exact Clerk env (login must match), THEN stop it + free the port from tailscale serve
#    (the dev server binds this port with HTTPS DIRECTLY).
docker exec demo-$N-next-web-app-1 printenv > /tmp/cenv.txt
$DC stop next-web-app
tailscale serve --https=$((3000+OFF)) off

# 3. Assemble $WT/apps/web/.env.local. NEXT_PUBLIC_* -> the TAILNET HTTPS host:offset (router/backend are F12'd +
#    tailscale-served on :$((5050+OFF))/$((8082+OFF))). Mirror the container's server-side Clerk keys. Point
#    CLERK_API_URL at the fake-bapi's REACHABLE IP — the host /etc/hosts `api.clerk.com` alias goes STALE on
#    re-bring-up (new docker IP) => the #1 login failure (`resolve handshake: fetch failed ECONNREFUSED`).
PK=$(grep -E '^NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=' "$STACK/.env.demo-$N" | cut -d= -f2-)
BIP=$(docker inspect demo-$N-fake-bapi-1 --format '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}')
{ echo "NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=$PK"
  echo "NEXT_PUBLIC_HOSTING_URL=https://$HOST:$((3000+OFF))"
  echo "NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=https://$HOST:$((5050+OFF))/graphql"
  echo "NEXT_PUBLIC_BACKEND_API_URL=https://$HOST:$((8082+OFF))"
  echo "DIRECTUS_PUBLIC_BASE_ADDR=https://content.anthropos.work"
  grep -E '^CLERK_SECRET_KEY=|^CLERK_JWT_KEY=|^CLERK_PUBLISHABLE_KEY=|^CLERK_WEBHOOK_SECRET=' /tmp/cenv.txt  # values-blind
  echo "CLERK_API_URL=http://$BIP:443"      # fake-bapi is plain HTTP on :443
  echo "STRIPE_SECRET_KEY=sk_test_dummy"    # current code has a module-eval throw; a dummy unblocks SSR
} > "$WT/apps/web/.env.local"

# 4. Reuse the demo's Tailscale cert (its SAN already covers $HOST — minted by /demo-up for the fake-FAPI);
#    or `tailscale cert --cert-file c.crt --key-file c.key $HOST`.
CERT="$(pwd)/$STACK/certs/fapi.crt"; KEY="$(pwd)/$STACK/certs/fapi.key"

# 5. Run native with HOT-RELOAD over HTTPS, bound to 0.0.0.0 (tailnet-reachable). Origin stays https://$HOST:<off>
#    so Clerk's middleware rewrite is RELATIVE/internal (no self-proxy loop). Do NOT tailscale-serve this port.
tmux new-session -d -s dfd-web-$N -c "$(pwd)/$WT/apps/web" \
  "bash -lc 'NODE_TLS_REJECT_UNAUTHORIZED=0 pnpm install && pnpm exec next dev --experimental-https \
     --experimental-https-cert $CERT --experimental-https-key $KEY -H 0.0.0.0 -p $((3000+OFF)) --turbopack'"

# 6. Reach it at the SAME tailnet URL as the rest of the demo: https://$HOST:$((3000+OFF)) (no SSH tunnel). LOG IN
#    via the cockpit handshake (redirect to the tailnet host — Chromium trusts the LE tailscale cert):
#      https://$HOST:$((5400+OFF))/v1/client/handshake?__clerk_identity=<hero-key>&redirect_url=https://$HOST:$((3000+OFF))/<path>
# Verify (NON-FATAL) in a REAL browser: log in, open the page, edit a string, confirm it hot-reloads in ~seconds.
# GOTCHA: cockpit `jump_to` links can 404 under dev (/enterprise/workforce/ai-readiness -> real route /ai-readiness).
```
**studio-desk** / **ant-academy**: same direct-HTTPS-on-the-tailnet principle for any native dev server that uses
Clerk — serve HTTPS with the tailscale cert on the offset port, never `tailscale serve` it.

---

## Run a target live — **backend Go targets** (app / cms / jobsimulation / skillpath) — TWO caveats

This is the **harder, more caveated path** — be honest with the user. Two separate problems:

1. **Infra endpoints are NOT in `.env`.** `platform/.env` has no `DB_CONNECTION`/`REDIS_ADDR`/`*_RPC_ADDR` — those
   are injected per-service in `docker-compose.yml` and point at **Docker service names** (`postgresql:5432`,
   `redis:6379`) a host process can't resolve. A native `go run .` therefore reaches **nothing** unless you
   **rewrite** them to the demo's offset host ports.
2. **Router federation.** The cosmo router (a container) fans out to subgraphs by service name
   (`http://backend:8082/...`). A host-native process can't answer that. Reaching a host process needs
   `extra_hosts: ["host.docker.internal:host-gateway"]` on the router — **a demo-tooling change we do NOT
   improvise** (SKILL Phase 7). So: **direct-to-service dev works** (hit the native service's own REST/RPC on its
   offset port); **full browser→router→native-subgraph federation** needs that wiring — **flag it as a tooling
   gap to raise, don't hack the stack.**

```bash
git -C stack-demo/app worktree add -b feat/<name> ../.worktrees/app-feat-<name>   # drop -b to resume
WT=stack-demo/.worktrees/app-feat-<name>
$DC stop backend                                   # frees 18082
# Build a native env FILE the tmux pane sources — a detached tmux session does NOT inherit THIS shell's exports
# once a tmux server is already running (it will be, if a frontend target launched first), so bake the env into a
# file rather than `export`ing here. platform/.env has SECRETS only; the DB/redis/RPC HOSTS are compose service
# names → REWRITE to demo-N offset host ports (the one thing that makes a host `go run .` reach anything).
ENVF="$(pwd)/.agentspace/dev-for-dummies/env-app-$N.sh"; mkdir -p "$(dirname "$ENVF")"   # gitignored; sources by ref
{ echo "set -a"
  echo ". $(pwd)/stack-demo/platform/.env"
  echo ". $(pwd)/$STACK/.env.demo-$N 2>/dev/null || true"
  echo "DB_CONNECTION='postgresql://postgres@localhost:$((5432+OFF))/postgres?sslmode=disable&search_path=public'"
  echo "REDIS_ADDR='localhost:$((6379+OFF))'"
  echo "SKILLER_RPC_ADDR='http://localhost:$((8083+OFF))'; CMS_RPC_ADDR='http://localhost:$((8091+OFF))'"
  echo "PORT=$((8082+OFF))"
  echo "set +a"; } > "$ENVF"
#   (Confirm the exact var NAMES + search_path against the service's docker-compose.yml `environment:` block — they
#    vary per service: backend→search_path=public, sentinel→sentinel. cms → ../cms, jobsimulation → ../jobsimulation,
#    skillpath → ../skillpath.)
tmux new-session -d -s dfd-app-$N -c "$(pwd)/$WT" \
  "bash -lc '. $ENVF; make setup && make gen && go run .'"   # cms/jobsimulation/skillpath: skip make setup && make gen
```

---

## Wrap up / cleanup — when the feature is truly done (user-initiated, AFTER the Phase 9 ritual)

Do NOT leave dangling state. After the PR is opened and reviewed, offer the user this teardown (per target, then
the demo). All plain-English, confirm before each:

```bash
# 1. Stop the native process
tmux kill-session -t dfd-web-$N            # (dfd-app-$N for a backend target)
# 2. Drop the remote proxy for that port (offset-scoped)
tailscale serve --https=$((3000+OFF)) off
# 3. Put the stack back the way it was — EITHER restore the container you stopped …
$DC up -d --no-deps next-web-app
#    … OR, if you're done with the whole demo, tear it down (also clears its tailscale serve + registry slot):
/demo-down $N
# 4. Remove the worktree + (optionally) the branch once the PR is merged
git -C stack-demo/next-web-app worktree remove ../.worktrees/next-web-app-feat-<name>
git -C stack-demo/next-web-app branch -d feat/<name>     # only after merge
# 5. Delete (or archive) the session manifest so a future run doesn't offer to resume finished work
rm .agentspace/dev-for-dummies/session-<slug>.yaml
```
> **Restart-container vs tear-down-demo is a real choice.** If you might come back to this feature, restore the
> container (step 3a) and keep the demo. If you're finished, `/demo-down $N` (3b) frees ~10–12 GB.

---

## Known gotchas (this box / v2.2 — apply ONLY if the demo verify surfaces them; all NON-FATAL)

From the box memory + `corpus/ops/demo/tailscale-serve.md` (F1–F13). All are demo-*infra* fixes, not TARGET
edits — they stay within `/demo-up`'s domain (SKILL Phase 7 forbids improvising outside a TARGET):

1. **`tailscale serve` shadowed by 0.0.0.0 container binds.** `tailscale serve` can't own the tailnet IP:port
   while a container binds `0.0.0.0:<port>`. Stopping a container frees it (which you do for a TARGET anyway);
   for containers you keep, prefix `127.0.0.1:` on the serve-fronted ports in `$STACK/docker-compose.injected.yml`
   + `$DC up -d --no-deps --force-recreate <svc>`, then `tailscale serve reset` and re-serve each port.
2. **Snapshot-cache digest mismatch → empty taxonomy.** If the catalog comes up empty, the cache predates the
   skiller→app merge; re-capture per `corpus/ops/snapshot-cold-start.md` and re-run set-dress (a confirmed
   prod read — public-only via `AssertPublicOnly`).
3. **Backend loses its Docker network endpoint** → `$DC up -d --no-deps --force-recreate backend`.
4. **Hand-teardown leaves serve config** — if you ever `docker rm` by hand, `tailscale serve reset` before the
   next `--public-host` up. `/demo-down` clears per-port serve automatically.

---

## Session manifest — `.agentspace/dev-for-dummies/session-<feature-slug>.yaml`

Written in Phase 6, read in Phase 3 (resume). Keep it human-readable — a person skims this to recognise their
own setup. `allowed_edit_roots` is the **mechanical guardrail** (SKILL Phase 7): before any edit, the agent
asserts the file path is under one of these.

```yaml
feature: ai-readiness-export           # or the fix scope
kind: feat                             # feat | fix
created_human: "Mon 14 Jul 2026, 15:32 (Europe/Rome)"
created_utc: "2026-07-14T13:32:07Z"
updated_human: "Mon 14 Jul 2026, 16:10"   # bump on resume
session:
  model: "Opus 4.8"
  effort: "max (ultrahigh)"
demo:
  n: 1
  public_host: "calypsostaging.taildc510.ts.net"
  app_url:     "https://calypsostaging.taildc510.ts.net:13000"
  cockpit_url: "http://calypsostaging.taildc510.ts.net:17700"   # plain HTTP; may not be remote-reachable
targets:
  - repo: next-web-app
    branch: feat/ai-readiness-export
    worktree: stack-demo/.worktrees/next-web-app-feat-ai-readiness-export
    native_port: 13000
    tmux: dfd-web-1
    live_url: "https://calypsostaging.taildc510.ts.net:13000"
allowed_edit_roots:                    # the ONLY paths this session may edit (SKILL Phase 7)
  - stack-demo/.worktrees/next-web-app-feat-ai-readiness-export
notes: "Frontend only. Backend untouched. Edits reflect on save."
```

**Resume checks (Phase 3):**
- Worktree dir exists + `git -C stack-demo/<repo> worktree list` shows the branch → **reuse** (never re-add).
  If the branch exists but the worktree dir was wiped, re-attach WITHOUT `-b`:
  `git -C stack-demo/<repo> worktree add stack-demo/.worktrees/<repo>-<slug> feat/<name>`.
- `/stack-list` shows `demo-N` **up** → reuse as-is; do **not** re-run `/demo-up` (a bare re-run re-does the
  slow set-dress/seed and can bounce the peers your native TARGET depends on). Only run `/demo-up N
  --public-host <host>` when the demo is **down**.
- tmux session alive (`tmux has-session -t <name>`) → attach; else relaunch per the recipe above.
