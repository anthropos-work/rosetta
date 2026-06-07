# Zero to Hero: Anthropos Development Setup Guide

This guide is designed to take you from a **fresh computer** (or a clean folder) to a fully running Anthropos development environment.

## 1. Prerequisites

Before we write any code, ensure you have the following tools installed.

### Setup Execution Guidelines (STEP RUN)

When executing each step of this setup guide, follow these guidelines to ensure reliability and documentation:

1. **Verify Before Install**: Run verification commands to check if a tool is already installed before attempting installation.
2. **Verify After Install**: After installation, run verification commands to confirm successful installation.
3. **Request Confirmation**: Always ask for user confirmation before installing or modifying system tools.
4. **Document Verification Commands**: Update this guide with verification commands for each tool if not already present.
5. **Continuous Improvement**: If you discover new steps or issues during setup, document them in this guide for future users.

### Essential Tools

### Essential Tools

#### 1. OS-Specific Setup

<details open>
<summary><strong>MacOS</strong></summary>

We recommend using [Homebrew](https://brew.sh/) for package management.

1.  **Git**: `brew install git`
    *   *Verification*: `git --version`
2.  **Docker Desktop**: [Install Docker Desktop for Mac](https://www.docker.com/products/docker-desktop/).
    *   *Verification*: `docker --version && docker compose version`
3.  **Visual Studio Code**: [Install VS Code](https://code.visualstudio.com/).
    *   *Verification*: `code --version`
4.  **Go** (v1.23+): `brew install go`
    *   *Verification*: `go version`.
5.  **Node.js** (v24+ required) & **pnpm**:
    *   `next-web-app/package.json` declares `"engines": { "node": ">=24.0.0" }`. Older Node versions will fail `pnpm install` with `WARN  Unsupported engine`.
    *   **Recommended**: Use [nvm](https://github.com/nvm-sh/nvm) to manage Node versions:
        ```bash
        curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash
        source ~/.nvm/nvm.sh
        nvm install 24
        nvm use 24
        nvm alias default 24    # optional: make it the default for new shells
        ```
    *   Alternative: `brew install node` (installs latest LTS)
    *   `corepack enable` or `npm install -g pnpm`
    *   *Verification*: `node --version` should print `v24.x.x`, then `pnpm --version`.
6.  **Build Tools**:
    *   Ensure XCode CLI tools are installed: `xcode-select --install`
    *   Ensure XCode CLI tools are installed: `xcode-select --install`
    *   *Verification*: `xcode-select -p`
7.  **Python** (v3.8+ for Studio-Room):
    *   `brew install python`
    *   *Verification*: `python3 --version`
8.  **Atlas** (Database Schema Manager):
    *   `curl -sSf https://atlasgo.sh | sh`
    *   *Verification*: `atlas version`

</details>

<details>
<summary><strong>Linux (Ubuntu/Debian)</strong></summary>

1.  **Git & Build Tools**:
    ```bash
    sudo apt-get update
    sudo apt-get install -y git build-essential
    ```
2.  **Docker**:
    *   Remove conflicting packages and install Docker Engine: [Official Guides](https://docs.docker.com/engine/install/ubuntu/).
    *   **Crucial Step**: Manage Docker as a non-root user (so you don't need `sudo`):
        ```bash
        sudo usermod -aG docker $USER
        newgrp docker
        ```
    *   *Verification*: `docker ps`.
3.  **Visual Studio Code**: [Install VS Code](https://code.visualstudio.com/docs/setup/linux).
4.  **Go** (v1.23+):
    *   [Official Install](https://go.dev/doc/install) is recommended to get the latest version, as apt repos are often outdated.
    *   *Verification*: `go version`.
5.  **Node.js** (v24+ required) & **pnpm**:
    *   `next-web-app/package.json` declares `"engines": { "node": ">=24.0.0" }`.
    *   **Recommended**: Use [nvm](https://github.com/nvm-sh/nvm) to manage Node versions:
        ```bash
        curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash
        source ~/.nvm/nvm.sh
        nvm install 24
        nvm use 24
        ```
    *   `corepack enable` or `npm install -g pnpm`
6.  **Python** (v3.8+ for Studio-Room):
    *   `sudo apt-get install python3 python3-pip python3-venv`
    *   *Verification*: `python3 --version`

</details>

---

## Automated Setup with Claude Code

If you're using **Claude Code**, you can automate this entire setup process using the `/dev-up` skill (which
consolidates the former setup-platform + start-platform — it builds the environment AND starts it):

```bash
/dev-up
```

The skill will:
*   Execute each step with verification before and after
*   Request your confirmation before installing tools or making changes
*   Track progress using TodoWrite
*   Create ops reports when issues are discovered

See [`.claude/skills/dev-up/SKILL.md`](../../.claude/skills/dev-up/SKILL.md) for details.

---

## 2. GitHub SSH Access Setup

Before cloning repositories, you need SSH access to the `anthropos-work` GitHub organization.

### Check Existing SSH Keys

First, verify if you already have SSH keys configured:

```bash
ls -al ~/.ssh
```

Look for files named `id_rsa.pub`, `id_ed25519.pub`, or similar `.pub` files.

### Test GitHub SSH Connection

```bash
ssh -T git@github.com
```

*Expected output*: `Hi [username]! You've successfully authenticated, but GitHub does not provide shell access.`
*Note*: Exit code 1 is normal for this command - it indicates successful authentication.

### If SSH Keys Don't Exist

1. **Generate a new SSH key**:
   ```bash
   ssh-keygen -t ed25519 -C "your_email@example.com"
   ```
   Press Enter to accept the default file location. Optionally set a passphrase.

2. **Start the ssh-agent**:
   ```bash
   eval "$(ssh-agent -s)"
   ```

3. **Add your SSH key to the ssh-agent** (macOS):
   ```bash
   ssh-add --apple-use-keychain ~/.ssh/id_ed25519
   ```

4. **Copy your public key**:
   ```bash
   pbcopy < ~/.ssh/id_ed25519.pub
   ```

5. **Add the SSH key to your GitHub account**:
   - Go to GitHub Settings > SSH and GPG keys
   - Click "New SSH key"
   - Paste your key and save

6. **Request access to anthropos-work organization**:
   - Contact the Engineering Manager to be added to the `anthropos-work` GitHub organization.

*Verification*: `ssh -T git@github.com` should show successful authentication.

---

## 3. Workspace Setup

We will create a dedicated workspace to house all the microservices and the frontend.

1.  Open your terminal.
2.  Navigate to the `rosetta` directory in this repository.
3.  Enter the scratchpad directory:
    ```bash
    cd stack-dev
    ```
    *(Note: This directory is git-ignored, so you can clone anything here without messing up the main repo).*

    **Stack-workspace convention**: each local stack lives in its own gitignored `stack-<role>/` dir — `stack-dev` (the dev stack we're building here), `stack-demo` (demo), `stack-dev-2` (a secondary dev stack), and future `stack-stage`/`stack-tests`. Each `stack-*/` dir spans one full local stack: its platform service repos **plus its own clone of `rosetta-extensions`** pinned to a tag.

    **`rosetta-extensions` clone roles**: all code/scripts that operate the corpus/platform on a spawned stack live in `rosetta-extensions` (never scattered in the rosetta corpus, never authored ad-hoc inside a stack dir). It has two clone roles:
    - **Authoring copy** at `.agentspace/rosetta-extensions/` — spawned on demand to read/build/**test** tooling, then commit + **tag**.
    - **Per-stack consumption copy** `stack-dev/rosetta-extensions @ <tag>` — each stack consumes the tooling at a pinned tag.

---

## 4. Cloning Repositories

The platform uses a **Makefile-driven workflow**. The `platform` repo is the orchestration hub, and `make init` automatically clones all required service repos as sibling directories.

### Clone the Platform Repo
```bash
git clone git@github.com:anthropos-work/platform.git
```
*Verification*: `ls platform/Makefile platform/repos.yml` should show both files.

### Clone All Service Repos (Automated)
The `make init` command reads `repos.yml` and clones all repos that don't already exist:
```bash
cd platform
make init
```
*Verification*: `make status` should list all repos with their branch and status.

This clones the repos declared in `platform/repos.yml` as siblings of `platform/`. As of 2026-05:

| Repo | Type | Has Migrations |
|------|------|---------------|
| `app` | Go backend | Yes (public schema) |
| `cms` | Go backend | Yes (cms schema) |
| `jobsimulation` | Go backend | Yes (jobsimulation schema) |
| `skiller` | Go backend | Yes (skiller schema) |
| `skillpath` | Go backend | Yes (skillpath schema) |
| `sentinel` | Go backend | No |
| `storage` | Go backend | No |
| `messenger` | Go backend | No |
| `roadrunner` | Go backend | No |
| `next-web-app` | Node.js (pnpm) | No |
| `studio-desk` | Node.js (npm) | No |
| `ant-academy` | Node.js (npm) — Next.js 16 + Expo, runs natively only | No |
| `graphql-wundergraph` | Node.js (npm) | No |

> **Note**: `chronos` and `intelligence` were removed from local orchestration (platform commits `045857c`, `fdfa189`). Their repos still exist on GitHub but `make init` no longer clones them. `customerio-sync` is built directly from its GitHub URL by docker-compose and is also not cloned locally. See [Service Taxonomy](../architecture/service_taxonomy.md) for current orchestration details.

### Initialize CMS Studio Submodule

The CMS service requires the Studio-Room Python project inside `cms/studio/` for its Docker build. This is **not** included in `make init` and must be cloned separately:

```bash
cd ../cms
make init-studio
```
*Verification*: `ls cms/studio/requirements.txt` should show the file exists.

### How Local Builds Work

**All services build from local directories.** Docker Compose uses `context: ../service` to build each service from its local clone using `Dockerfile.dev` (fast dev builds with BuildKit cache mounts).

This means:
- Every service **requires a local clone** to build
- `make init` handles cloning everything (except CMS studio submodule — see above)
- Changes to local code are picked up on `make up` (which runs `--build`)

### Optional Repos

These are not in `repos.yml` but useful for development:

```bash
# Knowledge Base (Claude Code plugin for AI-assisted development)
git clone git@github.com:anthropos-work/anthropos-knowledge-base.git

# Internal experiments hub
git clone git@github.com:anthropos-work/experiments.git
```

---

## 5. Environment Configuration

### The `.env` File
All services share a **single centralized `.env` file** located in the `platform` directory.

> **IMPORTANT**: You must obtain the master `.env` values from the Engineering Manager or the 1Password Vault "Engineering/Env".

1.  **Create the environment file**:
    ```bash
    cd platform
    cp .env_example .env
    ```
2.  **Populate secrets**: Edit `platform/.env` and fill in all required secret values from 1Password or the Engineering Manager.

    **Critical Keys Required**:
    *   `GH_PAT` (GitHub Personal Access Token — required for Docker builds)
    *   `CLERK_SECRET_KEY` (Auth — backend services)
    *   `OPENAI_KEY` (AI services)
    *   `AZURE_OPENAI_KEY` & `AZURE_OPENAI_ENDPOINT_URL` (Optional, Azure OpenAI)
    *   `AZURE_API_KEY` & `AZURE_ENDPOINT` (Optional, Azure Cognitive Services)
    *   `VITE_CLERK_PUBLISHABLE_KEY` (Only needed for Studio-Desk via Docker)
    *   `CLERK_WEBHOOK_SECRET` (Only needed if using Clerk webhooks)

3.  **Verification**: `ls -la platform/.env` should show the file exists.

> **STAGING SAFETY — outbound email kill switch.** If you intend to restore a prod DB dump into this stack (see [staging_from_dump.md](staging_from_dump.md)), `BREVO_KEY` **must be blank** before `make up`. The dump contains real customer emails and `messenger` will send them on any flow that triggers a notification. Blank the key in `platform/.env`:
>
> ```bash
> sed -i.bak 's/^BREVO_KEY=.*/BREVO_KEY=/' platform/.env
> ```
>
> Apply the same caution to `CUSTOMERIO_*`, `HEYGEN_WEBHOOK_SECRET`, `BUNNY_*`, `LIVEKIT_*`, `ELEVENLABS_*` if you don't intend to exercise those integrations.

**Note**: The docker-compose configuration uses this single `.env` file for all services (backend, cms, jobsimulation, etc.). Studio-Desk and Next.js also read from this `.env` when run via Docker profiles. Individual service repositories do not need their own `.env` files when running via Docker.

### Studio-Desk Environment (Only for Native Development)

If running Studio-Desk **outside Docker** (natively), it requires its own `.env` file:

```bash
cd ../studio-desk
cp .env.example .env
# Copy CLERK_SECRET_KEY, VITE_CLERK_PUBLISHABLE_KEY, OPENAI_KEY from platform/.env
```

**Note**: When running Studio-Desk via Docker (`make up PROFILE=studio-desk`), the platform `.env` is used automatically.

### Ant Academy Environment (Native Only — Not in Docker)

Ant Academy is the internal learning portal and **always** runs natively (no docker-compose profile). The Next.js app reads from `code/.env`, **not** the repo root:

```bash
cd ../ant-academy/code
cp .env.example .env
# Fill from platform/.env where possible:
#   NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY  (same as VITE_CLERK_PUBLISHABLE_KEY)
#   CLERK_SECRET_KEY
#   OPENAI_API_KEY            (server-side; for /api/ai/chat)
#   ANTHROPIC_API_KEY         (server-side; for /api/ai/chat)
#   FONTAWESOME_NPM_AUTH_TOKEN  (issued by the team; required for npm install)
# Optional:
#   REQUIRE_ORGANIZATION_MEMBERSHIP=0  # disable org-gate for solo local dev
```

See [Ant Academy service doc](../services/ant-academy.md) for details. Ant Academy is deployed to Vercel; nothing platform-side needs to change to run it.

### Clerk Webhook Setup (Optional)

If you need Clerk user and organization data to sync to your local database (required for user management features), you'll need to set up a webhook tunnel.

**Quick start** (no account required):
```bash
npx localtunnel --port 8082
```

Then configure the webhook in the Clerk dashboard.

**For full instructions**, see [Webhook Setup Guide](webhook_setup.md).

**Note**: This step can be skipped if you're only working on frontend features with existing test accounts.

---

## 6. Running the Platform (Docker via Makefile)

The platform uses a **Makefile** as the single entry point for all developer operations.

### Starting the Services

1.  Navigate to the platform directory:
    ```bash
    cd platform
    ```

2.  **Start all backend services** (default `graphql` profile):
    ```bash
    make up
    ```
    This builds from local repos and starts: PostgreSQL, Redis, Sentinel, Backend, CMS, Skiller, Skillpath, Storage, Jobsimulation, Roadrunner, Gotenberg, and the GraphQL/Cosmo Router.

    *Note*: First run may take several minutes as Docker builds images. Ensure your SSH agent is running (`ssh-add -l`).

3.  **Verification**:
    ```bash
    make ps
    ```
    You should see all services running. PostgreSQL and Redis should show as healthy.

### Prepare PostgreSQL Schemas (First Run Only)

After the first `make up`, PostgreSQL is running but missing schemas required by Sentinel and migrations. Create them now:

```bash
# Create pgvector extensions (required by CMS and Skiller migrations)
docker compose exec postgresql psql -U postgres -c "CREATE SCHEMA IF NOT EXISTS extensions; CREATE EXTENSION IF NOT EXISTS vector SCHEMA extensions; CREATE EXTENSION IF NOT EXISTS pg_trgm SCHEMA extensions;"

# Create Sentinel schema (required for Casbin authorization)
docker compose exec postgresql psql -U postgres -c "CREATE SCHEMA IF NOT EXISTS sentinel;"

# Restart Sentinel (it was crash-looping without its schema)
docker compose restart sentinel
```

*Verification*: `make ps` should show Sentinel with `Up` status (not `Restarting`).

> **Note**: These schemas only need to be created once. They persist across `make down` / `make up` cycles. Only `make reset-db` requires re-creating them.

### Database Migrations

After the first startup, apply database schemas:
```bash
make migrate
```
This automatically runs Atlas migrations for all repos that have `migrations: true` in `repos.yml` (currently: app, cms, jobsimulation, skiller, skillpath).

*Verification*: Commands should complete without errors.

To migrate a single service:
```bash
make migrate S=cms
```

### Profiles

Start specific service groups instead of the full stack:

| Command | What it starts |
|---------|---------------|
| `make up` | All backend + GraphQL router (default) |
| `make up PROFILE=backend` | Backend (app) only |
| `make up PROFILE=cms` | CMS only |
| `make up PROFILE=frontend` | Next.js in Docker |
| `make up PROFILE=studio-desk` | Studio-Desk in Docker |
| `make up-all` | Everything |

Base services (PostgreSQL, Redis, Sentinel) always start regardless of profile.

### Ongoing Operations

For daily platform operations (starting, stopping, rebuilding services), see the [Run Guide](run_guide.md).

---

## 7. Running Frontend

The Next.js frontend is a monorepo with multiple apps. Each app needs its own `.env` file.

### Configure Environment Files

<!-- TODO: Improve keys management — currently each app needs manual .env setup with keys copied from platform/.env. Consider a shared env solution or a script to automate this. -->

1.  Navigate to the frontend repo:
    ```bash
    cd ../next-web-app
    ```

2.  **Create the web app `.env`**:
    ```bash
    cp apps/web/.env.example apps/web/.env
    ```

3.  **Populate keys** in `apps/web/.env`: Copy `CLERK_SECRET_KEY` from `platform/.env` and set `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` to the Clerk publishable key from 1Password or the Engineering Manager.

    *Note*: The GraphQL and Backend URLs already default to `localhost:5050` and `localhost:8082` which are correct for local development.

    *Verification*: `ls apps/web/.env` should show the file exists.

### Install and Run

4.  Install dependencies:
    ```bash
    pnpm install
    ```

5.  Run the development server:
    ```bash
    pnpm dev
    ```
    This starts all apps in the monorepo (web, hiring, integration).

6.  Open `http://localhost:3000` (or the port shown in terminal).
    - Web app: http://localhost:3000
    - Hiring app: http://localhost:3001
    - Integration app: http://localhost:3002

---

## 8. Running Studio-Desk (Design Tool)

Studio-Desk is the simulation design tool - a required part of the full platform for content creation workflows.

1.  Navigate to studio-desk:
    ```bash
    cd ../studio-desk
    ```

2.  Install dependencies:
    ```bash
    npm install
    ```

3.  Start the development server:
    ```bash
    npm run dev
    ```
    This starts both the frontend (port 9100) and backend (port 9000). Ports are configurable via `.env`.

4.  Access at `http://localhost:9100`

    *Verification*: You should see the Studio-Desk login page (uses Clerk authentication).

---

## 9. Running Studio-Room (Optional - On-Demand)

Studio-Room is the AI generation pipeline. It runs **on-demand** for specific generation tasks, not as a persistent service.

1.  Navigate to studio-room:
    ```bash
    cd ../studio-room
    ```
2.  Install requirements:
    ```bash
    pip3 install -r requirements.txt
    ```
3.  Run a test generation:
    ```bash
    python3 gen.py --media simulation --template default
    ```

---

## 10. Troubleshooting

### "Generated code missing" / "command not found: make"
If running Go services locally (outside Docker), you may hit errors about missing files.
*   **MacOS**: Ensure `xcode-select --install` is run.
*   **Linux**: Ensure `sudo apt-get install build-essential`.

Resolution:
```bash
cd [service-name]
make setup
make gen
```

### "Connection Refused" / Docker Issues
*   **General**: Ensure Docker containers are running (`make ps` or `docker compose ps`). If a service is failing, check logs: `make logs S=service_name`.
*   **Linux Permission Denied**: If you see "permission denied while trying to connect to the Docker daemon", you likely skipped the `usermod` step. Run `sudo usermod -aG docker $USER`, then log out and back in (or `newgrp docker`).

### "SyntaxError: Unexpected identifier 'assert'" (Frontend - Legacy)
This issue occurred with older versions of the frontend that used `import ... assert { type: 'json' }` syntax removed in Node.js v22+. The frontend has since been updated and now works with Node.js v22+. If you encounter this error on an old branch, switch to the latest `main` branch.

### "schema 'extensions' does not exist" (Atlas migrations)
CMS and Skiller services require the pgvector extension for vector embeddings.
*   **Solution**: The custom PostgreSQL image (built from `platform/postgresql/`) should include pgvector. If missing:
    ```bash
    docker compose exec postgresql psql -U postgres -c "CREATE SCHEMA IF NOT EXISTS extensions; CREATE EXTENSION IF NOT EXISTS vector SCHEMA extensions;"
    ```
*   Then retry: `make migrate`

### Sentinel Crashing / Restarting
Sentinel requires its own database schema for Casbin authorization.
*   **Solution**: Create the sentinel schema:
    ```bash
    docker compose exec postgresql psql -U postgres -c "CREATE SCHEMA IF NOT EXISTS sentinel;"
    ```
*   Then restart sentinel: `docker compose restart sentinel`

### Port Already In Use
If you have another Docker stack running, ports may conflict.
*   **Solution**: Stop the other stack first, or run `make down` to stop the current stack.

### Docker Build Fails with "Permission denied (publickey)"
Docker builds services from local repos but needs SSH access for Go module downloads.
*   **Solution**: Ensure your SSH agent is running with keys loaded:
    ```bash
    # Check if SSH agent has keys
    ssh-add -l

    # If "no identities" or "agent not running", start it and add your key:
    eval "$(ssh-agent -s)"
    ssh-add ~/.ssh/id_ed25519

    # Verify GitHub access
    ssh -T git@github.com
    ```
*   Also ensure `GH_PAT` is set in `platform/.env`
*   Then retry: `make up`

### Full Database Reset
If your database is corrupted or you want a clean start:
```bash
make reset-db
```
This removes PostgreSQL data, restarts the container, and re-runs all migrations.

### "Permission denied" when starting Postgres after a fresh checkout
Bitnami Postgres runs as uid 1001 inside the container. The bind-mount root (`platform/data/postgresql`) is created by Docker as root and the container can't write to it. Pre-create with the right ownership before first start:
```bash
sudo mkdir -p platform/data/postgresql && sudo chown -R 1001:1001 platform/data/postgresql
docker compose up -d postgresql
```

### CMS image build fails on `COPY studio/` / `pip install studio/requirements.txt`
The `studio/` submodule has been removed from `cms/main` but `cms/Dockerfile.dev` still references it. Edit `cms/Dockerfile.dev` and remove these two lines:
```dockerfile
COPY studio/ ./studio/
RUN pip install --no-cache-dir -r studio/requirements.txt
```
The Go binary runs without the Python studio runner.

### `studio-desk` fails to bind host port 9100
Conflicts with `node_exporter` (Prometheus monitoring) if you have any observability stack running on the box. Edit `platform/docker-compose.yml`:
```yaml
studio-desk:
  ports:
    - "9101:9100"   # was 9100:9100
```

### Next.js (`next-web-app`) build crashes with `STRIPE_SECRET_KEY is not configured`
Next.js statically evaluates server routes (e.g. `/api/create-subscription`) at build time and reads from `process.env`. Compose's `env_file` is runtime-only — the build step doesn't see it. Drop a gitignored `next-web-app/apps/web/.env.production` containing the keys those routes touch (`STRIPE_SECRET_KEY`, `OPENAI_API_KEY`, `AZURE_OPENAI_*`, `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY`, `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`, `NEXT_PUBLIC_BACKEND_API_URL`, `NEXT_PUBLIC_HOSTING_URL`) before `docker compose build`.

### Backend GraphQL "404 Not Found" at `/graphql`
The backend GraphQL endpoint is `/graphql/query`. The `/graphql` path returns the embedded Apollo Sandbox playground UI. CORS preflight + auth happen at `/query`.

### Tailscale / cross-device access — frontend points at `localhost`
`NEXT_PUBLIC_*` vars are baked into the Next.js client bundle at build time. To make the staging accessible from another device on your Tailscale network, set `PUBLIC_HOST=100.x.y.z` (your host's Tailscale IP) in `platform/.env` BEFORE `docker compose build next-web-app`, then add the Tailscale origin to the backend CORS allowlist (`app/internal/cors/cors.go`) and to the dev Clerk app's allowed origins.

### Logged in but Members table empty / "forbidden" on enterprise routes
You restored a prod DB dump and the engineer-to-Clerk rebind is incomplete. See [staging_from_dump.md](staging_from_dump.md) — typically you need to (a) apply the colony patch, (b) sync `sentinel.casbin_rules.g2` from `public.memberships`, and (c) restart sentinel.

### Blank page on first sign-in / "Clerk: infinite redirect loop" in next-web-app logs
Clerk's Next.js middleware reads `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` from `process.env` at runtime even though the same value is baked into the client bundle at build time. If only `VITE_CLERK_PUBLISHABLE_KEY` is in the runtime env (which is what compose's `env_file: .env` provides for Studio-Desk), the server-side Clerk init fails and trips the "infinite redirect loop" detector — pages render blank until the cookie cache rescues them.

Fix: add the Clerk vars to `next-web-app`'s runtime `environment:` block in `platform/docker-compose.yml`:

```yaml
next-web-app:
  environment:
    # ... existing ...
    - NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=${NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY}
    - NEXT_PUBLIC_CLERK_SIGN_IN_URL=${NEXT_PUBLIC_CLERK_SIGN_IN_URL:-/login}
    - NEXT_PUBLIC_CLERK_SIGN_UP_URL=${NEXT_PUBLIC_CLERK_SIGN_UP_URL:-/sign-up}
    - NEXT_PUBLIC_CLERK_AFTER_SIGN_IN_URL=${NEXT_PUBLIC_CLERK_AFTER_SIGN_IN_URL:-/}
    - NEXT_PUBLIC_CLERK_AFTER_SIGN_UP_URL=${NEXT_PUBLIC_CLERK_AFTER_SIGN_UP_URL:-/onboarding}
```

Then `docker compose up -d next-web-app` to restart with the new env (no rebuild needed — these are runtime-only vars, the bundle was already correct). After the fix you should see all five `NEXT_PUBLIC_CLERK_*` keys in `docker compose exec next-web-app env`.

If you still see blank pages in your browser after the fix, clear cookies for the staging origin (`ithacastaging`, `calypsostaging`, or your equivalent) — stale dev-browser cookies bound to a prior origin can keep the redirect loop alive client-side.

### Frontend slow to load over Tailscale
By default the root layout loads Plausible, Google Tag Manager, BetterStack, analytics.bellasio.com, and PostHog — ~10 third-party blocking requests that drag page-load time over Tailscale and pollute prod analytics with staging traffic. Disable them per [staging_from_dump.md](staging_from_dump.md) — set `NEXT_PUBLIC_DISABLE_ANALYTICS=true` in `platform/.env`, blank `POSTHOG_API_KEY`/`POSTHOG_SERVER_SIDE_KEY`, rebuild `next-web-app`. Production leaves the flag unset and keeps analytics on.

---

## 11. Maintenance Guidelines

This `setup_guide.md` and the `/dev-up` Claude skill (which consolidates the former setup-platform + start-platform) are interconnected documents that must be maintained together.

### When You Update This Setup Guide

If you modify the setup process (add/remove/reorder steps), you must update:

1.  **Dev-Up Skill** (`.claude/skills/dev-up/SKILL.md`): Update phase definitions, step sequences, and verification commands
2.  **This Guide**: Ensure all steps have verification commands documented

### General Guidelines

*   **OS-Specific Differences**: When a step differs between macOS and Linux, ensure this guide reflects the appropriate commands/tools for each OS
*   **Agent-Friendly**: Ensure all documents remain parseable and clear for autonomous agents
*   **Verification Commands**: Every installation step should have a documented verification command
