# Staging Bringup — From Fresh VM to Working Personal Staging

This is the **spine doc** for setting up a personal Anthropos staging environment that mirrors Stefano's on Ithaca: full prod-shaped data, the platform stack running on Tailscale, your engineer account bound to a dev Clerk app, and PR-ready repos.

**Audience:** a new engineer (or AI agent) joining the Anthropos team who needs their own staging tomorrow. Recreation standard: read this end-to-end and you should be able to `gh pr open` against the current `main` from your staging within a working day.

**Companion docs:**

- [`staging_from_dump.md`](./staging_from_dump.md) — original prod-dump restore procedure (more verbose; this doc supersedes/integrates it).
- [`staging-sync.md`](./staging-sync.md) — the daily sync routine that keeps every staging on `origin/main`.
- [`staging-clerk.md`](./staging-clerk.md) — the shared dev Clerk app and the cross-engineer test login.

**Sections:**

- [§ 0 Mental model](#0-mental-model)
- [§ 1 Prerequisites](#1-prerequisites)
- [§ 2 Clone repos and lay out the workspace](#2-clone-repos-and-lay-out-the-workspace)
- [§ 3 Environment file](#3-environment-file)
- [§ 4 Live data — restoring the prod pg_dump](#4-live-data--restoring-the-prod-pg_dump)
- [§ 4.5 Apply pending Atlas migrations](#45-apply-pending-atlas-migrations) — **read before §5**, prod dump is always behind `main` migrations
- [§ 5 Bring up the stack](#5-bring-up-the-stack) — incl. **Quirk #11 (colony v1+v2 JWT)**, **Quirk #3 (cms studio submodule)**, and the rest of the 19-quirk narrative
- [§ 6 Engineer rebind — make Clerk match your DB](#6-engineer-rebind--make-clerk-match-your-db)
- [§ 7 Optional HTTPS via `tailscale serve`](#7-optional-https-via-tailscale-serve)
- [§ 8 Install the sync routine](#8-install-the-sync-routine)
- [§ 9 Apply skip-worktree hygiene](#9-apply-skip-worktree-hygiene)
- [§ 10 Smoke check](#10-smoke-check)
- [§ 10.5 Known schema drifts (expected on staging)](#105-known-schema-drifts-expected-on-staging)
- [§ 11 You're done](#11-youre-done)

---

## 0. Mental model

What you're building (per-engineer, on a Tailscale-attached VM):

```
+------------------ <yourhost>.taildc510.ts.net (Tailscale) ------------------+
|                                                                            |
|  /home/<you>/platform/        docker compose orchestrator                  |
|  /home/<you>/{app,cms,...}    15 service-repo sibling clones (always main) |
|  /home/<you>/rosetta/         this corpus                                  |
|  /home/<you>/ant-singularity/ agent fleet & operations docs                |
|                                                                            |
|  docker stack on :3000 (next-web-app) / :5050 (graphql) / :8082 (backend)  |
|       └── postgres restored from a 12 GB prod pg_dump                      |
|       └── auth via dev Clerk app `national-elk-17` (shared with all eng)   |
|                                                                            |
|  systemd-user timers:                                                      |
|    anthropos-staging-drift.timer  (hourly drift check)                     |
|    anthropos-staging-sync.timer   (daily 06:00 UTC force-sync to main)     |
+----------------------------------------------------------------------------+
```

**Hard rule:** every repo on your staging is always on `origin/main` HEAD. No feature branches on staging, ever. If you need to test a feature branch against prod-shape data, do it from an `--unmanaged` agentspace workspace on your laptop, or spin up a separate `wip-<initials>` host. See [`staging-sync.md`](./staging-sync.md#what-if-a-developer-wants-to-test-a-feature-branch-on-staging) for the reasoning.

---

## 1. Prerequisites

### Tailnet membership

Your VM must be on the Anthropos Tailscale tailnet (`taildc510.ts.net`). Once it is, ask Stefano or an admin to add a friendly alias to the ACL `hosts:` block — e.g. `wip-mn → 100.x.y.z` AND `wip-mnstaging → 100.x.y.z` (same IP, two aliases). Anything you want browsers to trust must be a real DNS name, not just an IP.

```hcl
{
  "hosts": {
    "ithaca":         "100.120.254.65",
    "ithacastaging":  "100.120.254.65",
    "calypso":        "100.83.121.80",
    "calypsostaging": "100.83.121.80",
    "wip-mn":         "100.x.y.z",
    "wip-mnstaging":  "100.x.y.z"
  }
}
```

Edit at https://login.tailscale.com/admin/acls/file. Resolution is instant tailnet-wide.

### GitHub HTTPS credentials

```bash
gh auth login              # follow the device prompts; scope: repo + read:org
gh auth setup-git          # configures the HTTPS credential helper
gh auth token > ~/.gh-token  # used as GH_PAT for docker build args
```

Why HTTPS, not SSH: the `Makefile` and every service Dockerfile uses `GH_ACCESS_TOKEN` over `git config insteadOf https://`. The SSH agent path is vestigial (Phase 2-D of the 2026-05-14 cleanup strips the `ssh: ["default"]` directives from `docker-compose.yml`). HTTPS works on every host with no key gymnastics.

### Docker + system

- Linux (Ubuntu 22.04 LTS+ recommended). Apple Silicon laptops work for dev but the staging fleet is x86_64 Linux.
- Docker Engine + `docker compose` v2.20+.
- `psql` client (`apt install postgresql-client-16`).
- ≥30 GB free disk (12 GB dump + restored DB + docker images).
- ≥16 GB RAM is ample. (The earlier "~10-12 GB" figure was an unmeasured over-estimate — the `graphql`-profile dev stack idles at **~0.9 GB measured**; `--profile all` adds only a handful more services. RAM is not the constraint it was once assumed to be — see [`rosetta_demo.md`](rosetta_demo.md) § Resource budget.)
- `node` + `npm` 20+ on the host (only for running the Playwright smoke script outside Docker).

### Group membership

If your VM has the team's analytics-and-reports unix group set up (for `db-query` from the host), add your user:

```bash
sudo usermod -aG analytics-and-reports $USER
# log out + back in for it to take
```

---

## 2. Clone repos and lay out the workspace

The platform's `Makefile init` target does the heavy lifting (it clones every repo in `repos.yml` as a sibling of `platform/`). To match the layout the sync routine expects, also clone `rosetta` and `ant-singularity` alongside.

```bash
cd ~
git clone https://github.com/anthropos-work/platform.git
cd platform
make init                  # clones app/, cms/, jobsimulation/, ...
                           # uses GH_PAT under-the-hood via the gh-cli helper

cd ~
git clone https://github.com/anthropos-work/rosetta.git
git clone https://github.com/stefano-anthropos/ant-singularity.git
git clone https://github.com/anthropos-work/anthropos-knowledge-base.git
```

**Quirk #1** — `make init` may issue `git clone git@github.com:` (SSH). If yours doesn't have `gh auth setup-git` configured, you'll see prompts for SSH keys. Fix by editing `Makefile` to `git clone https://github.com/` (or land the upstream PR that does this) before re-running `make init`. The dockerfiles themselves use `GH_PAT` over HTTPS — no SSH agent needed.

**Quirk #2** — the compose service `customerio-sync` originally builds from `git@github.com:anthropos-work/customerio-sync.git#main` (Docker daemon doesn't have your GitHub creds). On staging clones this is patched to `context: ../customerio-sync` and the repo is cloned locally. Upstream `platform/docker-compose.yml` may still carry the SSH form — update it locally:

```yaml
customerio-sync:
  build:
    context: ../customerio-sync
    # remove: context: git@github.com:anthropos-work/customerio-sync.git#main
```

You will end up with this layout:

```
/home/<you>/
├── platform/                      # orchestrator (Makefile, docker-compose.yml, .env)
├── app/                           # Go backend (CORS, GraphQL gateway)
├── cms/                           # Go content management
├── skillpath/                     # Go skill-path runtime
├── jobsimulation/                 # Go AI simulations service
├── sentinel/                      # Go authz (casbin)
├── storage/                       # Go S3-shim
├── messenger/                     # Go transactional email (Brevo)
├── roadrunner/                    # Go scheduler
├── customerio-sync/               # Go marketing-email sync
├── next-web-app/                  # Next.js 15 frontend monorepo
├── studio-desk/                   # TypeScript content design tool
├── graphql-wundergraph/           # GraphQL federation gateway
├── rosetta/                       # this corpus
├── anthropos-knowledge-base/      # knowledge layer
└── ant-singularity/               # agent fleet (this node)
```

---

## 3. Environment file

Copy `platform/.env.example` to `platform/.env` and fill these (ask Stefano for the dev secrets if you don't have them):

```bash
# Required for stack to come up
GH_PAT=ghp_…                       # = gh auth token; reused as docker build arg
NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=pk_test_…  # dev Clerk app `national-elk-17`
CLERK_SECRET_KEY=sk_test_…

# Public-host name baked into next-web-app at build time (the Tailscale alias)
PUBLIC_HOST=<your>staging          # e.g. wip-mnstaging (must be in ACL hosts:)

# Outbound-email kill switch (mandatory — prod dump has real customer emails)
BREVO_KEY=

# Analytics kill switch (don't pollute prod dashboards from staging traffic)
NEXT_PUBLIC_DISABLE_ANALYTICS=true
POSTHOG_API_KEY=
POSTHOG_SERVER_SIDE_KEY=

# Build-time vars Next.js statically evaluates (more in Quirk #4)
STRIPE_SECRET_KEY=sk_test_…
OPENAI_API_KEY=sk-…
AZURE_OPENAI_ENDPOINT=…
AZURE_OPENAI_API_KEY=…

# Required if you want Talk to Data (/ask/stream) to actually answer questions.
# Without these, the SSE stream opens fine but the agentic loop fails with
# "no EC2 IMDS role found" (the AWS SDK falls through to instance-metadata
# lookup, which times out on Tailscale VMs). Use the prod-equivalent inference
# profile region — eu-west-1 today.
AWS_ACCESS_KEY_ID=AKIA…
AWS_SECRET_ACCESS_KEY=…
AWS_REGION=eu-west-1
```

Production keys are NOT used here. Only dev/test keys. The AWS creds are an exception — Talk to Data calls Bedrock via the prod inference profile (no dev tenancy yet), so use the dedicated staging IAM user Stefano keeps for this. Ask Stefano if you don't have those.

**Quirk #4** — Next.js 15 statically evaluates server routes at build time (`/api/create-subscription`, `/api/wundergraph/*`). Compose `env_file` is **runtime-only**, so build-time evaluation will crash with `STRIPE_SECRET_KEY is not configured` etc. Drop a gitignored `.env.production` into `next-web-app/apps/web/`:

```bash
cat > ~/next-web-app/apps/web/.env.production <<EOF
STRIPE_SECRET_KEY=sk_test_…
OPENAI_API_KEY=sk-…
AZURE_OPENAI_ENDPOINT=…
AZURE_OPENAI_API_KEY=…
NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=pk_test_…
NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=http://<yourhost>staging:5050/graphql
NEXT_PUBLIC_BACKEND_API_URL=http://<yourhost>staging:8082
NEXT_PUBLIC_HOSTING_URL=http://<yourhost>staging:3000
EOF
```

`.env.production` is in `.gitignore` already; add it to `.git/info/exclude` as well to make `git status` clean (see [`staging-sync.md` § Skip-worktree handling](./staging-sync.md#skip-worktree-handling)).

---

## 4. Live data — restoring the prod pg_dump

Stefano keeps a recent prod dump at `/home/devops/backup_22934492224-317.sql` on Ithaca (12 GB, plain SQL, pg_dump 16). To pull it to your VM:

```bash
scp ithaca:/home/devops/backup_*.sql ~/prod_dump.sql
```

(Or ask Stefano for an alternative S3-signed URL if SCP is blocked.) Then:

```bash
cd ~/platform

# Bring up only Postgres first
docker compose up -d postgresql
until docker compose exec -T postgresql pg_isready -U postgres > /dev/null 2>&1; do sleep 1; done

# Restore (~10-15 min on a stock VM)
cat ~/prod_dump.sql | docker compose exec -T postgresql psql -U postgres -d postgres -v ON_ERROR_STOP=0 > /tmp/restore.log 2>&1
```

### Expected warnings during restore (all harmless)

- **Quirk #9** — `ERROR: role "<name>" does not exist` for `backend`, `cms`, `skiller`, `chronos`, `customerio`, `simulator`, `sentinel`, `skillsgateway`, `skillpath`. These are GRANT/ALTER OWNER statements that no-op against a fresh box. Data tables load fine. Pre-create the roles if you want clean output: `CREATE ROLE backend; CREATE ROLE cms; ...`.
- **Quirk #14** — `invalid command \restrict` / `\unrestrict` at the very start and end of the file. PG 16 client emits these tokens, PG 15 client doesn't recognize them. Cosmetic. Strip with `sed -i '/^\\\\\(un\)\?restrict\b/d' ~/prod_dump.sql` if it bothers you.

### Sanity-check the restore

**Important: the dump restores into the default `postgres` database, not into a separate `anthropos` DB.** The Bitnami Postgres image creates `postgres` as the bootstrap DB, the dump's top-level `\connect postgres` lands all data there, and so the schemas you care about (`public`, `sentinel`, `cms`, `skillpath`, `jobsimulation` — plus the legacy `skiller` schema older dumps carry, pre-dating the skiller→app merge) all sit *inside* `postgres`. Running `docker compose exec -T postgresql psql -U postgres -c '\l'` listing only `postgres / template0 / template1` is the **expected** post-restore shape — it is **not** evidence that the restore failed. The 2026-05-14 Ithaca repair burned an hour on this misread; don't repeat it.

To actually sanity-check the data, you must query *inside* the `postgres` DB (note the `-d postgres`):

```bash
docker compose exec -T postgresql psql -U postgres -d postgres -c "
  SELECT 'users' tbl, COUNT(*) FROM public.users
  UNION ALL SELECT 'organizations', COUNT(*) FROM public.organizations
  UNION ALL SELECT 'memberships', COUNT(*) FROM public.memberships
  UNION ALL SELECT 'casbin_rules', COUNT(*) FROM sentinel.casbin_rules;
"
```

You should see thousands of users (~5,951 in the 2026-05-14 dump), hundreds of orgs (~250), thousands of memberships (~3,597), and thousands of casbin rules (~5,168). If those numbers are zero or the schemas don't exist, *that's* a failed restore — re-check `/tmp/restore.log` for `ERROR` lines that aren't covered by Quirk #9 / #14.

### If restore fails completely

**Quirk #8** — the Bitnami Postgres bind-mount needs uid 1001. Wipe and re-init:

```bash
docker compose down
sudo rm -rf data/postgresql
sudo mkdir -p data/postgresql && sudo chown -R 1001:1001 data/postgresql
docker compose up -d postgresql
# then retry the restore
```

You can re-restore later (after an upstream schema migration breaks something) by repeating the wipe+restore cycle. No need to re-do steps 5-7 since the dev Clerk binding is per-row in the DB, but you WILL need to re-do step 6 (engineer rebind) since the dump's `clerk_id` columns get overwritten.

---

## 4.5. Apply pending Atlas migrations

**Read this before §5.** The prod dump is taken at a single point in time — every migration applied to prod *after* that snapshot date is sitting in the service repos' `terraform/migrations/` directories but has **not** been replayed against your restored DB. On 2026-05-14 the gap was 11 migrations across 3 services (app: 6, skiller: 3, jobsimulation: 2) including `ask_conversations` / `ask_messages` (which Talk to Data writes to) and `skill_translations` (which the skiller subgraph reads). Without these migrations, you bring the stack up "successfully" but Talk to Data 500s with `relation "ask_conversations" does not exist` the first time anyone tries to send a query.

Neither `make migrate` (no longer wired on staging clones) nor the daily `anthropos-staging-sync` ([`staging-sync.md`](./staging-sync.md)) runs Atlas. **This is an operator responsibility** — every fresh bringup, and every time prod ships new migrations you want reflected on staging.

### Install Atlas (one-time)

```bash
curl -sSf https://atlasgo.sh | sh        # installs /usr/local/bin/atlas
atlas version                            # confirm — 0.30+ or v1.x both work
```

### Apply per service

Each Go service that owns DB schema has its own `terraform/migrations/` directory and its own Atlas config under `terraform/atlas.hcl` declaring the `local` env. The pattern for each:

```bash
cd ~/<service>

# See what's pending against the running staging DB
atlas migrate status --env local \
  --url "postgresql://postgres@localhost:5432/postgres?sslmode=disable&search_path=<schema>"

# Apply
atlas migrate apply --env local \
  --url "postgresql://postgres@localhost:5432/postgres?sslmode=disable&search_path=<schema>"
```

Per-service `<schema>` values (this is the `search_path` Atlas writes the `atlas_schema_revisions` table into and treats as default):

| Service           | `search_path=` | Why                                                          |
| ----------------- | -------------- | ------------------------------------------------------------ |
| `app`             | `public`       | Owns `users`, `organizations`, `memberships`, `ask_conversations`, `ask_messages`, audit logs — plus the merged skiller taxonomy (skills, job roles, translations, embeddings) since July 2026 |
| `cms`             | `cms`          | Directus / content schema                                    |
| `skillpath`       | `skillpath`    | Skill-path runtime                                           |
| `jobsimulation`   | `jobsimulation`| Job sims, interview extraction results                       |

Apply in any order — schemas don't cross-reference at the migration level. `sentinel` is not in the table because it uses raw Casbin schema management, not Atlas.

### Baselining (only if `atlas migrate apply` complains about non-linear history)

You may see:

```
Error: pending migration files are not in a linear order: 20240304140158.sql
```

This means a migration file exists in the repo but is not recorded in `atlas_schema_revisions`. Sometimes the file's effects are already in the DB (applied via a later migration that overlapped, or out-of-band). To check whether it's truly missing or just unrecorded, run the relevant `\d <table>` and inspect for columns the file would have added.

If the effects are already present, re-baseline to the latest *real* migration **before** the out-of-order file's revision, then apply normally:

```bash
# Set the revisions table to a known-applied version so Atlas re-bases on it.
# Use a version that you can confirm IS applied (latest non-conflicting one).
atlas migrate set <latest-applied-version> --env local --url "..."

# OR, use --exec-order linear-skip to bypass the linearity check and apply
# the newer migrations cleanly while leaving the out-of-order file flagged.
atlas migrate apply --exec-order linear-skip --env local --url "..."
```

**Gotcha (cost an hour on Ithaca):** `atlas migrate set <old-version>` *truncates* the revisions table down to that version. If you pick a version too far back, Atlas will then try to re-apply already-applied intermediate migrations and fail again. The safer move is `atlas migrate set <latest-applied-version>` to keep the table consistent, or `--exec-order linear-skip` as above. The 2026-05-14 Ithaca run used `linear-skip`; Calypso used `migrate set` to the latest applied revision — both worked.

### Restart the services whose schemas moved

```bash
cd ~/platform
docker compose restart backend jobsimulation
# replace with whatever services you applied migrations for
```

A `restart` is enough — the Go services re-open DB connections on boot. No rebuild needed unless the migration came with a code change (it usually does, but the daily sync's `docker compose build` covers that path).

### Reference

The 2026-05-14 Calypso bringup applied:

- **app** (6 migrations, public): `20260414085836` (drop legacy triggers) → `20260506145258` (`ask_query_lessons`).
- **skiller** (3 migrations, skiller): `20260417103036` → `20260511094027` (`job_role_translations`, `skill_translations` + multilingual tsvector indexes).
- **jobsimulation** (2 migrations, jobsimulation): `20260402145459` (`interview_extraction_results`) → `20260409131539` (`summary` jsonb column).
- **cms** + **skillpath**: already up to date.

The per-service apply logs from that run live at `/tmp/calypso-migrations/{app,skiller,jobsimulation}-apply.log` on Calypso (kept for ~30 days). Use them as expected-output reference when something looks off.

---

## 5. Bring up the stack

```bash
cd ~/platform
docker compose --profile all up --build -d
```

Wait 5-15 min for all 15 services to report healthy:

```bash
docker compose ps --format "table {{.Service}}\t{{.Status}}"
```

If something crashes, check its logs (`docker compose logs <svc> --tail 50`). Most failures map to one of the bringup quirks below.

### Bringup quirks consolidated as a procedural narrative

This is the integrated form of the 19 quirks Stefano discovered during the Ithaca bringup. As you run `make up` / `docker compose up`, here's what will break and what to do:

1. **Quirk #1 — Makefile uses SSH** — already addressed in §2. Patch `git clone git@github.com:` → `https://github.com/`.

2. **Quirk #2 — `customerio-sync` builds from a git URL** — already addressed in §2. Use `context: ../customerio-sync`.

3. **Quirk #3 — `cms/Dockerfile.dev` references removed `studio/` submodule** — `COPY studio/` (line ~39) and `RUN pip install -r studio/requirements.txt` (line ~42) fail with `not found`. Comment them out (mark the lines with `# Staging patch (Quirk #3)` so the next operator knows why); the Go binary runs fine without the Python studio runner. The 2026-05-14 cleanup opened [`anthropos-work/cms#fix/dockerfile-remove-studio-submodule`](https://github.com/anthropos-work/cms/pulls) to fix upstream — **PR is still open and unmerged as of 2026-05-14**, so the patch must be re-applied on every fresh clone and is one of the long-lived skip-worktree files on each staging. When the PR lands, the daily sync's `git reset --hard origin/main` will drop the staging-local comments naturally and the skip-worktree entry on `cms/Dockerfile.dev` can be removed.

4. **Quirks #4 + #5 — Next.js needs build-time env vars** — already addressed in §3. Drop `apps/web/.env.production` before first build. Make sure `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` is also in compose's runtime `environment:` block (not just `env_file:`) — see Quirk #15.

5. **Quirk #6 — Backend CORS doesn't trust your staging origin.** Hardcoded in `app/internal/cors/cors.go`. Two ways:
   - **Until [`anthropos-work/app#feat/cors-extra-origins-env`](https://github.com/anthropos-work/app/pulls) is merged:** unmark skip-worktree, edit `cors.go` to append `"http://<yourhost>staging:3000"` (and `:8000`, `:9000` if you use them) in the `colony.Development` branch, rebuild backend.
   - **Once merged:** set `CORS_EXTRA_ORIGINS=http://<yourhost>staging:3000` in `platform/.env` and restart backend — no rebuild needed.

6. **Quirk #7 — `studio-desk` host port 9100 collides with `node_exporter`** if you run any observability stack (Prometheus etc.). Remap to `9101:9100` in `platform/docker-compose.yml`:

   ```yaml
   studio-desk:
     ports:
       - "9101:9100"   # was 9100:9100
   ```

7. **Quirks #8, #9, #14 — Postgres bind-mount + restore warnings** — already addressed in §4.

8. **Quirk #10 — Backend GraphQL endpoint is `/graphql/query`**, not `/graphql`. The `/graphql` path returns Apollo Sandbox UI; CORS preflight + auth happen at `/query`. The Wundergraph router (`:5050`) federates these into `/5050/graphql`. Tools that expect `/graphql` directly need to know.

9. **Quirk #11 — `colony` has two separate Clerk auth bugs.** Both bite every staging today; the working fix on Ithaca + Calypso is a single vendored copy of `colony` that patches both. **Read both halves before reaching for the vendor recipe** — the recipe is identical, but knowing what each piece fixes is what lets you keep it pruned over time.

   **Bug 1 — nil-client / nil-email panic** (`colony@v0.34.0`). `authn/provider/clerk.GetUser` constructs `&User{}` without wiring the `client` field; `u.client.Get()` panics if the JWT lacks custom claims (which dev Clerk apps don't ship by default). Same panic in `Email()` when `PrimaryEmailAddressID` is nil. The 2026-05-14 cleanup opened [`anthropos-work/colony#fix/clerk-getuser-nil-client`](https://github.com/anthropos-work/colony/pulls) to fix upstream — that branch fixes the panic but NOT bug 2 below.

   **Bug 2 — Clerk v2-JWT claim shape unsupported.** New Clerk apps (including `national-elk-17`) default to a v2 session-token format (`"v": 2`) that nests org info under a single `o` key:

   ```json
   {
     "o": { "id": "org_…", "rol": "admin", "slg": "while-true-srl-…" },
     "sub": "user_…",
     "sid": "sess_…",
     "v": 2
   }
   ```

   `colony` v0.34.0 (and the `fix/clerk-getuser-nil-client` branch on top) only reads v1 names: `org_id`, `org_role`, `org.eid`. On a v2 token, `GetOrganization()` returns `nil` → every `colony.User.GetOrganization()`-gated endpoint (Talk to Data `/ask/*`, Members listing, Workforce Intelligence, anything else REST-backed) 403s with `missing organization context`. The GraphQL ent privacy layer is unaffected, so the bug *looks* like "REST is broken, GraphQL is fine," which is misleading — it's a single common root cause.

   The dashboard "Customize session token" template (per [`staging-clerk.md`](./staging-clerk.md#customizing-the-dev-clerk-session-token-recommended)) could in principle inject `org.eid` directly and sidestep bug 2, but the dev app currently has no template configured, the REST `PATCH /v1/instance` for `session_token_template` is silently dropped (dashboard-only as of 2026-05), and `POST /v1/jwt_templates` creates a template that only takes effect if the frontend explicitly calls `getToken({template: 'colony'})` — which it doesn't. So a code-side fix is the only viable path until upstream colony lands.

   **The working fix on Ithaca + Calypso** is a 309-LOC patch on top of `fix/clerk-getuser-nil-client` (saved at `/tmp/colony-v2-jwt-patch.diff` on Ithaca, 2026-05-14). Three pieces, all in `authn/provider/clerk/`:

   - `clerk.go`: `Clerk` struct gains an `orgClient OrganizationClient` field; `NewProvider` wires `organization.NewClient(cfg)` into it (reuses the same `BackendConfig` already used by `user.Client` + `jwks.Client`).
   - `clerk_user.go`: new constant `claimOrgV2 = "o"`, new `OrganizationClient` interface (one-method `Get(ctx, idOrSlug)`), new process-wide `orgEidCache` (clerk-org-id → eid). `GetOrganization()` reads v1 claims first, then falls back to `tokenClaims.Extra["o"].{id, rol, public_metadata.eid}` for any field v1 didn't supply. If `orgEid` is still empty (the realistic case — v2 tokens omit `public_metadata` by default), a new `lookupOrgEid(clerkOrgID)` helper fetches the org from the Clerk Backend API, reads `public_metadata.eid` from `clerk.Organization.PublicMetadata` (a `json.RawMessage`), caches it, and returns it. Errors return `""` quietly so the caller sees the existing "no org context" → 403, never a panic.
   - `clerk_user_test.go`: three new tests (`TestUser_ValidOrgClaims_V2`, `TestUser_V2Claims_LazyFetchEID` with cache-hit assertion, `TestUser_V1ClaimsStillWork` regression check). All pass; full clerk suite stays green.

   **Until both upstream changes land (the nil-client PR + a v2-claim follow-up that doesn't exist yet)**, vendor it. Same Ithaca recipe as before — the layout in §6 / Dockerfile.dev `COPY vendor-colony` lines / `replace` directive in each `go.mod` is unchanged; just make sure the source `colony` tree you copy from has the patch applied:

   ```bash
   # one-time on each new staging clone
   cd ~ && git clone https://github.com/anthropos-work/colony.git
   cd colony && git checkout fix/clerk-getuser-nil-client
   git apply /path/to/colony-v2-jwt-patch.diff
   go test ./authn/provider/clerk/...      # all green incl. v2 tests
   for svc in app cms jobsimulation; do
     rm -rf ~/$svc/vendor-colony
     cp -r ~/colony ~/$svc/vendor-colony
     rm -rf ~/$svc/vendor-colony/.git
     grep -q "replace github.com/anthropos-work/colony => ./vendor-colony" ~/$svc/go.mod \
       || printf '\nreplace github.com/anthropos-work/colony => ./vendor-colony\n' >> ~/$svc/go.mod
     grep -q "vendor-colony" ~/$svc/Dockerfile.dev \
       || sed -i '/^COPY go.sum/a COPY vendor-colony ./vendor-colony' ~/$svc/Dockerfile.dev
     (cd ~/$svc && go mod tidy)
     (cd ~/$svc && git update-index --skip-worktree go.mod go.sum Dockerfile.dev)
     grep -q "^vendor-colony/" ~/$svc/.git/info/exclude \
       || echo "vendor-colony/" >> ~/$svc/.git/info/exclude
   done
   cd ~/platform && docker compose up --build -d backend cms jobsimulation
   ```

   Three services vendor `colony` because all three call `colony.User.GetOrganization()`: `app`, `cms`, `jobsimulation` (`skiller` used to be the fourth, until its July 2026 merge into `app`). `go.mod` colony pins (v0.34.0 / v0.33.2 / v0.33.0 respectively) stay unchanged — the `replace` directive overrides them. The `go.mod`, `go.sum`, `Dockerfile.dev` get `skip-worktree`'d so the daily sync doesn't reset them; `vendor-colony/` goes in `.git/info/exclude` so it survives `reset --hard` as an unknown file.

   **Why two PRs, not one:** the nil-client PR is small and obviously correct — likely lands soon. The v2-claim fallback needs a follow-up upstream because it adds a real dependency on the Clerk Backend API at request time (rate-limited), and the long-term fix is the dashboard session-token template plus a colony switch that prefers it. Until then, the lazy fetch + cache works fine at staging-scale traffic.

10. **Quirk #12 — Dev Clerk needs Organizations enabled + per-user/org `external_id` set.** Documented as the rebind procedure in §6 below.

11. **Quirk #13 — Dev Clerk "new device" sign-in challenge** blocks programmatic login. Bypass with `POST /v1/sign_in_tokens` for Playwright / CI. Real-user login through the form is fine (Clerk emails the code on first sign-in, then trusts the device).

12. **Quirk #15 — `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` must be in next-web-app's runtime `environment:` block, not just `env_file:`.** Clerk middleware reads it from `process.env` at runtime. If only `VITE_CLERK_PUBLISHABLE_KEY` is in the runtime env, Clerk's server-side init falls into the "infinite redirect loop" detector → blank pages. Fix: list `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` explicitly in compose's `next-web-app.environment:` array. Plus the four sign-in/up URL vars. Restart container — no rebuild needed (runtime-only). Sibling client-side symptom: stale `__clerk_db_jwt` cookies from a prior origin keep the loop alive after the env fix; clear cookies for the staging origin to recover.

13. **Quirk #16 — Disable third-party analytics on staging.** `apps/web/src/app/layout.tsx` eagerly loads ~10 third-party blocking scripts (Plausible, GTM → GA + FB + LinkedIn + Google Ads, BetterStack, analytics.bellasio.com, plus PostHog). Drags page-load over Tailscale and pollutes prod analytics with staging traffic. Already addressed in §3 via `NEXT_PUBLIC_DISABLE_ANALYTICS=true`.

14. **Quirk #17 — Tailscale ACL `hosts:` for friendly aliases** — addressed in §1. Each new alias also needs: (a) Clerk allowed_origins, (b) backend CORS (or `CORS_EXTRA_ORIGINS` env once the upstream PR lands).

15. **Quirk #18 — `ssh: ["default"]` in compose breaks builds on hosts without `SSH_AUTH_SOCK`.** Vestigial — Dockerfiles use `GH_ACCESS_TOKEN` over HTTPS. On Calypso (no SSH agent in the shell), build fails immediately with `invalid empty ssh agent socket`. Strip with `sed -i '/ssh: \["default"\]/d' docker-compose.yml`. The 2026-05-14 cleanup opened [`anthropos-work/platform#chore/drop-ssh-default-compose-directives`](https://github.com/anthropos-work/platform/pulls) to fix upstream.

16. **Quirk #19 — Staging clones use `skip-worktree` + `.git/info/exclude`** to keep long-lived staging-only patches invisible to upstream PRs. Full mechanics in [`staging-sync.md` § Skip-worktree handling](./staging-sync.md#skip-worktree-handling). Idempotent script lives at `/tmp/staging_clean.sh` on Ithaca + Calypso. After: `git status` is clean, `git add .` only stages real edits, files stay on disk so docker builds keep working.

17. **`next-web-app` Clerk fetch monkey-patch (post-#19 addition).** Server Components in `next-web-app` hit `UND_ERR_CONNECT_TIMEOUT` to `api.clerk.com` from inside Docker after the first request. Plus on HTTP staging (`http://<host>:3000`), Clerk's Secure cookies get dropped by the browser → infinite redirect loop. Fix: copy `clerk-fetch-fix.js` from Ithaca verbatim, mount + `NODE_OPTIONS=--require=`. Full details and the file's content in [`staging-clerk.md` § Pitfalls](./staging-clerk.md#pitfalls-that-bit-us). Load-bearing — do not skip.

---

## 6. Engineer rebind — make Clerk match your DB

After restoring the prod dump, every `users.clerk_id` and `organizations.clerk_id` in your DB points at **prod** Clerk IDs that don't exist in the dev Clerk app. If you log in now, Clerk authenticates you but the backend can't find your user row → blank /home, no admin context, all enterprise routes redirect to /profile.

The fix is the engineer-rebind procedure documented at length in `corpus/ops/staging_from_dump.md` § 3 — **read that file end-to-end before continuing.** It's still the canonical reference for:

1. Creating your Clerk user in the dev app, setting `external_id` to your DB UUID, rewriting `public.users.clerk_id`.
2. Enabling Organizations on the dev Clerk app, creating matching dev orgs, setting `public_metadata.eid`, rewriting `public.organizations.clerk_id`.
3. Syncing `sentinel.casbin_rules.g2` from `public.memberships` so the authz layer recognizes you as admin.
4. (Optional, recommended) Customizing the dev Clerk session token to embed `eid`, `email`, `firstname`, `lastname`, and `org.eid` claims — avoids per-request Clerk API fetches and 429 rate-limit pain.

**Shortcut: use Stefano's account.** If you don't need your own user to exist in the DB, you can just log in as `stefano@anthropos.work / chichi88kora` — that's the shared cross-engineer test login. See [`staging-clerk.md` § Shared test login](./staging-clerk.md#shared-test-login).

The full singularity-catalog blueprint for this loop lives at [`ant-singularity/knowledge/singularity-catalog/auto-anthropos-staging-dev-loop.md`](https://github.com/stefano-anthropos/ant-singularity/blob/main/knowledge/singularity-catalog/auto-anthropos-staging-dev-loop.md) — it has the SQL one-liners and the Clerk REST snippets in one place.

---

## 7. Optional: HTTPS via `tailscale serve`

Stefano set this up on Ithaca as `https://ithacastaging.taildc510.ts.net`. The dev Clerk app has these origins in `allowed_origins` already:

```bash
sudo tailscale serve --bg https://ithacastaging.taildc510.ts.net http://localhost:3000
sudo tailscale serve status
```

Replicate on your host:

```bash
sudo tailscale serve --bg https://<yourhost>.taildc510.ts.net http://localhost:3000
```

Then ask Stefano to add `https://<yourhost>.taildc510.ts.net` to the Clerk `allowed_origins` list (see [`staging-clerk.md` § Adding a new staging host](./staging-clerk.md#adding-a-new-staging-host)).

**Caveat: graphql/backend env-vars are baked HTTP at build time.** `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` and `NEXT_PUBLIC_BACKEND_API_URL` get baked into the Next.js bundle pointing at `http://<host>:5050|8082/...`. Browser → HTTPS frontend → HTTP backend = Mixed Content blocking → blank dashboards. Use the plain `http://<yourhost>staging:3000` URL for end-to-end testing until those vars are HTTPS too (and the backend has TLS).

---

## 8. Install the sync routine

This is what makes your staging "stay alive" — daily force-pull to `origin/main` + rebuild + smoke test. Mandatory on every staging host. Full details: [`staging-sync.md`](./staging-sync.md).

```bash
# Copy from Ithaca (or any existing staging host)
scp -r ithaca:~/.local/bin/anthropos-staging-{sync,drift,smoke}.{sh,js} ~/.local/bin/
scp -r ithaca:~/.config/systemd/user/anthropos-staging-{sync,drift}.{service,timer} ~/.config/systemd/user/

# Allow systemd-user units to run while you're logged out
sudo loginctl enable-linger $USER

# Enable + start both timers
systemctl --user daemon-reload
systemctl --user enable --now anthropos-staging-sync.timer anthropos-staging-drift.timer
```

Verify:

```bash
systemctl --user list-timers anthropos-staging-*
# expect both ACTIVE + a next-trigger time within 24h

~/.local/bin/anthropos-staging-drift.sh
cat ~/.local/state/anthropos-staging-sync/drift.summary
# expect a one-liner with drift=0 once everything's in sync
```

---

## 9. Apply skip-worktree hygiene

Once everything's running, mark your local patches `skip-worktree` so agent commits stay clean:

```bash
# For each repo with uncommitted edits:
for repo in app cms next-web-app platform; do
  cd ~/$repo
  git diff --name-only | while read f; do
    git update-index --skip-worktree -- "$f"
  done
done

# For untracked staging-only dirs:
cd ~/app && echo 'vendor-colony/' >> .git/info/exclude
cd ~/cms && echo 'vendor-colony/' >> .git/info/exclude
cd ~/next-web-app && echo 'apps/web/.env.production' >> .git/info/exclude
cd ~/platform && echo 'clerk-fetch-fix.js' >> .git/info/exclude
```

Full mechanics + recovery: [`staging-sync.md` § Skip-worktree handling](./staging-sync.md#skip-worktree-handling).

---

## 10. Smoke check

Final-step verification. Log in via Playwright (the same script the sync routine uses):

```bash
SMOKE_URL=http://<yourhost>staging:3000 \
SMOKE_EMAIL=stefano@anthropos.work \
SMOKE_PASSWORD=chichi88kora \
node ~/.local/bin/anthropos-staging-smoke.js
```

Pass criteria:

- HTTP 200 on `/login`, no console errors.
- After form submission, redirect to `/home` within 120s.
- `/home` renders the dashboard greeting ("Hi, Stefano" — yes, Stefano, because you logged in as Stefano; if you rebound your own account in §6 it'll be "Hi, <Your Name>").
- Workforce Intelligence sidebar item is present.

If smoke fails, check `docker compose logs --since 5m next-web-app` for `UND_ERR_CONNECT_TIMEOUT`, `infinite redirect loop`, or `clerkError`. Most of the time the culprit is the clerk-fetch-fix not being loaded (see [`staging-clerk.md`](./staging-clerk.md#symptom-und_err_connect_timeout-from-server-components)) or a Tailscale alias / Clerk allowed_origin / backend CORS gap.

---

## 10.5. Known schema drifts (expected on staging)

These are *expected* failures on a current staging — they're not bringup mistakes and don't need a local fix. They show up because `next-web-app@main` and `app@main` aren't perfectly in lockstep at every moment of every day, and the daily sync force-resets both to whatever HEAD looks like at 06:00 UTC. The drifts are upstream-coordination matters; flag them in `#anthropos-eng` and wait.

| Drift                                                                       | Symptom                                                                                                              | Backend response                                                                                                            | First seen   |
| --------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | ------------ |
| `Membership.jobRole` / `Membership.targetRole` / `Info.jobRole` / `Query.recapUserSkills` pass a `language: ContentLanguage` argument the backend doesn't expose | Members page renders "0 / 50 — No data" despite 47+ active members in the DB; Workforce Intelligence is unaffected | `Unknown argument "language" on field …` from Wundergraph; some skillpath subgraph queries also return 422 (`Failed to fetch from Subgraph 'skillpath'`) | 2026-05-14   |

**Do not patch this locally.** Resolving it means regenerating `next-web-app/packages/graphql/` against the current backend schema (a frontend team change) or extending the schema to accept the arg (a backend team change) — neither is a staging-local edit. If Members 0/50 is blocking demo work *today*, you can either (a) demo a different surface (Workforce Intelligence renders fine), or (b) check out matching versions of `next-web-app` and `app` from before the drift on a non-staging workspace (NOT on a staging clone — see [`staging-sync.md` § What if I want to test a feature branch](./staging-sync.md#what-if-i-want-to-test-a-feature-branch-with-prod-shape-data)).

When a drift here is resolved, delete its row.

---

## 11. You're done

You now have:

- A Tailscale-attached staging serving `http://<yourhost>staging:3000` with full prod-shaped data.
- A working Clerk login (yours or Stefano's shared account).
- Daily auto-sync of every repo to `origin/main` at 06:00 UTC + hourly drift check.
- Skip-worktree hygiene so any PR you open from this clone is clean.

To open a real PR upstream from your staging clone:

```bash
cd ~/<repo>
git checkout -b fix/something
# edit code (skip-worktree files won't appear in git status; that's intentional)
git add <file> && git commit -m "fix: something"
git push -u origin fix/something
gh pr create --base main --title "fix: something" --body "…"
```

`git status` shows only what you actually changed, not the staging-only patches. PRs are clean.

---

## Related

- [`staging_from_dump.md`](./staging_from_dump.md) — verbose original (this doc supersedes it; staging_from_dump remains as the engineer-rebind reference).
- [`staging-sync.md`](./staging-sync.md) — daily sync routine + skip-worktree mechanics.
- [`staging-clerk.md`](./staging-clerk.md) — dev Clerk app, shared test login, monkey-patch.
- [`setup_guide.md`](./setup_guide.md) — original setup guide (no prod-dump path).
- [`run_guide.md`](./run_guide.md) — day-to-day operations.
- [`update_guide.md`](./update_guide.md) — pulling latest code (now superseded by the auto-sync routine on staging hosts).
- [Ant-singularity catalog entry](https://github.com/stefano-anthropos/ant-singularity/blob/main/knowledge/singularity-catalog/auto-anthropos-staging-dev-loop.md) — org-level workflow framing.
