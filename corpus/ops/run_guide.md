# Running the Anthropos Platform Locally

This guide takes you from a **completed setup** to a **running local platform** that you can access in your browser.

> **Prerequisites**: This guide assumes you have completed the [Setup Guide](setup_guide.md). All tools must be installed and repositories cloned to `stack-dev/`.
>
> `stack-dev/` is the DEV stack dir under the `stack-*` convention — a gitignored workspace spanning the dev stack's platform service repos plus its own clone of rosetta-extensions. Any stack tooling it runs comes from its tagged `stack-dev/rosetta-extensions` clone (authored in `.agentspace/rosetta-extensions/`, then committed + tagged), not ad-hoc scripts in rosetta.

## Quick Reference: Service URLs

Once running, access these URLs in your browser:

| Service | URL | Description |
|---------|-----|-------------|
| **Frontend (Web App)** | http://localhost:3000 | Main user-facing application |
| **Studio-Desk** | http://localhost:9100 | Simulation design tool |
| **Ant Academy** | http://localhost:3077 | Internal learning portal (`@anthropos.work` only) |
| **GraphQL Playground** | http://localhost:5050 | API gateway (Cosmo Router) |
| **Backend API** | http://localhost:8082 | Backend service (Connect RPC) |

---

## Run Execution Guidelines (RUN STEP)

When executing each step of this guide, follow these principles:

1. **Check Before Start**: Verify if a service is already running before attempting to start it.
2. **Verify After Start**: Confirm the service started successfully with health checks.
3. **Handle Conflicts**: If a port is already in use, identify and resolve the conflict.
4. **Document Issues**: Update this guide when you discover problems or better approaches.

---

## 1. Verify Docker Environment

Before starting any services, ensure Docker is running.

### Check Docker Status

```bash
docker info > /dev/null 2>&1 && echo "Docker is running" || echo "Docker is NOT running"
```

*Expected*: `Docker is running`

### If Docker is Not Running

**macOS**: Open Docker Desktop from Applications, or:
```bash
open -a Docker
```
Wait 30-60 seconds for Docker to initialize, then re-verify.

**Linux**: Start the Docker daemon:
```bash
sudo systemctl start docker
```

### Check for Existing Containers

```bash
cd stack-dev/platform
make ps
```

This shows any existing containers from previous runs.

---

## 2. Start Backend Services

The platform uses a **Makefile** as the single entry point. Infrastructure (PostgreSQL, Redis) and Sentinel start automatically with any profile.

### Navigate to Platform Directory

```bash
cd stack-dev/platform
```

### Option A: Start Full Backend (Recommended)

This starts all backend services + GraphQL router (default `graphql` profile):

```bash
make up
```

This starts: PostgreSQL, Redis, Sentinel, Backend, CMS, Storage, Jobsimulation, Roadrunner, Gotenberg, and GraphQL/Cosmo Router. (Skillpath is no longer a separate service — its engine merged into Backend/`app`, "skillpath-in-app".)

*Note*: First run may take several minutes as Docker builds images from local repos.

### Option B: Start Specific Profile

Start only the services you need:

```bash
make up PROFILE=backend    # Backend (app) only
make up PROFILE=cms        # CMS only
make up-all                # Everything including frontend and studio-desk
```

See [Setup Guide](setup_guide.md#profiles) for the full profile list.

### Verify Backend Services

```bash
make ps
```

*Expected*: All started containers showing `Up` status. PostgreSQL and Redis should show `(healthy)`.

### Check Service Logs (if issues)

```bash
# View logs for a specific service
make logs S=backend

# View logs for all services
make logs
```

### Health Check: GraphQL Gateway

```bash
curl -s http://localhost:5050/health || echo "GraphQL not responding"
```

*Expected*: Health check response or API info.

---

## 4. Start Frontend (Next.js Web App)

**Required** — the frontend always runs natively (not in Docker for the `graphql` profile). It must be
started in a **tmux session** so it survives Claude Code session closure.

### Verify Node.js Version

The frontend requires **Node.js v24+**.

```bash
node --version
```

*Expected*: `v24.x.x`

### Install Dependencies (if needed)

```bash
ls stack-dev/next-web-app/node_modules > /dev/null 2>&1 && echo "Dependencies installed" || (cd stack-dev/next-web-app && pnpm install)
```

### Start in tmux

```bash
# Idempotent — skips if the session already exists
tmux has-session -t anthropos-web 2>/dev/null || \
  tmux new-session -d -s anthropos-web -c "$(pwd)/stack-dev/next-web-app" 'pnpm dev:web'
```

*Expected*: tmux session `anthropos-web` starts; Next.js serves on http://localhost:3000

### Attach to view logs

```bash
tmux attach -t anthropos-web   # detach with Ctrl+B D
```

### Verify Frontend

Open http://localhost:3000 in your browser.

*Expected*: Anthropos login page or dashboard (depending on auth state).

### Troubleshooting: Port Already in Use

```bash
lsof -i :3000
```

Options:
1. Stop the conflicting process
2. Kill the existing tmux session and restart: `tmux kill-session -t anthropos-web`

---

## 5. Start Studio-Desk (Design Tool)

Studio-Desk is the simulation design tool used by content creators.

### Option A: Run in Docker (Recommended)

```bash
cd stack-dev/platform
make up PROFILE=studio-desk
```

This starts Studio-Desk in Docker along with its dependencies (GraphQL, CMS).

### Option B: Run Natively

For development with hot-reloading:

```bash
cd stack-dev/studio-desk
npm install
npm run dev
```

*Note*: Requires its own `.env` file when running natively. See [Setup Guide](setup_guide.md#studio-desk-environment-only-for-native-development).

*Expected*:
- Frontend: http://localhost:9100
- Backend: http://localhost:9000

### Verify Studio-Desk

Open http://localhost:9100 in your browser.

*Expected*: Studio-Desk login page (uses Clerk authentication).

---

## 5.1 Start Ant Academy (Optional — Internal Learning Portal)

Ant Academy is the standalone Next.js 16 / Expo learning portal for `@anthropos.work` employees. It runs **natively only** — no docker-compose profile. It authenticates via **Clerk** and — since v0.5.1 — **reads its course catalog from the platform academy subgraph over GraphQL** (`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`); without that backend it still boots but the catalog grid **renders empty** (see [`../services/ant-academy.md` § The Content Model](../services/ant-academy.md#the-content-model--db-authoritative-catalog-v051-m7)).

### Navigate

```bash
cd stack-dev/ant-academy/code
```

> **Note**: the Next.js app lives in `code/`, not the repo root. The repo-root `.env` is for content-authoring tooling, **not** the React app.

### Verify Node Version

```bash
node --version
```

*Expected*: `v22.x.x` or newer (declared in `code/package.json` `engines.node: ">=22"`).

> If you use nvm and the rest of the platform is on v24, that works too — Ant Academy's `>=22` is a lower bound, not a pin.

### Verify Env

```bash
ls .env > /dev/null 2>&1 && echo ".env present" || echo "Copy .env.example -> .env and fill keys (see setup_guide.md)"
```

### Install Dependencies (first run only)

```bash
npm install
```

> **Font Awesome Pro**: no token required. The FA Pro icons are vendored/self-hosted in `code/public/assets/fontawesome/` (webfonts + `css/all.min.css`), so `npm install` pulls only from the public registry — a token-less install works. `FONTAWESOME_NPM_AUTH_TOKEN` lingers in `.env.example` but is vestigial; you don't need it to install or run Ant Academy.

### Start in tmux

```bash
# Idempotent — skips if the session already exists
tmux has-session -t anthropos-academy 2>/dev/null || \
  tmux new-session -d -s anthropos-academy -c "$(pwd)" 'npm run dev'
```

*Expected*: tmux session `anthropos-academy` starts; dev server on **http://localhost:3077** (not 3000 — that port is reserved for `next-web-app`).

### Attach to view logs

```bash
tmux attach -t anthropos-academy   # detach with Ctrl+B D
```

### Verify

Open <http://localhost:3077>. You should land on the Clerk sign-in page; sign in with an `@anthropos.work` Google account.

### Troubleshooting

| Symptom | Cause | Fix |
|---------|-------|-----|
| Redirect to `/no-organization` | Org-membership gate (default ON) | Set `REQUIRE_ORGANIZATION_MEMBERSHIP=0` in `code/.env` for solo dev |
| Sign-in fails with "domain not allowed" | Account not on `@anthropos.work` | Sign in with a work account, or update the Clerk app's allowed domain |
| 401 from `/api/ai/chat` | Missing `OPENAI_API_KEY` / `ANTHROPIC_API_KEY` server-side | Fill `code/.env` with the server keys (different from the in-app Cosmo's localStorage key) |

### Mobile App (Optional)

```bash
cd ../mobile
pnpm install
pnpm run dev:web    # web preview on http://localhost:8555
```

---

## 5.5 Start Webhook Tunnel (Optional)

If you need Clerk webhooks to sync user/organization data to your local database, start a tunnel:

```bash
# Quick start - no account needed
npx localtunnel --port 8082
```

Copy the URL (e.g., `https://gentle-flies-think.loca.lt`) and configure it in Clerk Dashboard > Webhooks with path `/api/webhook/clerk`.

**When is this needed?**
- Creating new users/organizations
- Testing membership changes
- Working on user management features

**When can you skip this?**
- Pure frontend development with existing test accounts
- Working on features that don't involve user/org sync

For detailed setup (including more reliable alternatives like ngrok), see [Webhook Setup Guide](webhook_setup.md).

---

## 6. Start Studio-Room (Optional - On-Demand)

Studio-Room is the AI-powered generation pipeline. Unlike other services, it runs **on-demand** for specific generation tasks, not as a persistent service.

### Navigate to Studio-Room Directory

```bash
cd stack-dev/studio-room
```

### Verify Python Environment
```bash
python3 --version
```

### Install Dependencies (if needed)
```bash
pip3 install -r requirements.txt
```

### Run a Generation
```bash
python3 gen.py --media simulation --template default
```

---

## 7. Stopping Services

### Stop native processes (tmux sessions)

```bash
tmux kill-session -t anthropos-web      # next-web-app
tmux kill-session -t anthropos-academy  # ant-academy (if running)
tmux kill-session -t anthropos-backend  # native backend (if developed locally)
```

### Stop Docker Services

```bash
cd stack-dev/platform

# Stop all services (keeps data)
make down
```

### Check All Services Stopped

```bash
make ps                   # Docker containers
tmux list-sessions        # tmux sessions (should be empty)
```

*Expected*: No containers listed, no tmux sessions.

---

## 8. Common Scenarios

### Scenario: Resume After Computer Restart

1. Start Docker Desktop
2. Start Docker services: `cd stack-dev/platform && make up`
3. Start frontend in tmux:
   ```bash
   tmux new-session -d -s anthropos-web -c "$(pwd)/stack-dev/next-web-app" 'pnpm dev:web'
   ```
4. (Optional) Start Studio-Desk in Docker: `make up PROFILE=studio-desk`

> **Note**: PostgreSQL schemas (extensions, sentinel) persist across restarts. You do NOT need to re-create them unless you run `make reset-db`.

### Scenario: Quick Frontend Development

If you only need to work on the frontend and backend is already running:

```bash
# Verify backend is up
cd stack-dev/platform && make ps

# Start frontend in tmux (idempotent)
tmux has-session -t anthropos-web 2>/dev/null || \
  tmux new-session -d -s anthropos-web -c "$(pwd)/stack-dev/next-web-app" 'pnpm dev:web'
```

### Scenario: Develop a Single Service Natively

To stop a service's Docker container and run it locally in a tmux session:

```bash
cd stack-dev/platform
make dev S=backend       # Stops the Docker container for backend/app

# Start natively in tmux so it survives Claude session closure
tmux new-session -d -s anthropos-backend -c "$(pwd)/../app" 'go run .'
tmux attach -t anthropos-backend   # tail logs; Ctrl+B D to detach
```

### Scenario: Database Reset

To wipe the database and start fresh:

```bash
cd stack-dev/platform
make reset-db
```

This removes PostgreSQL data, restarts the container, and re-runs all migrations automatically.

> **Important**: After `reset-db`, you must re-create the PostgreSQL schemas before migrations will succeed. See [Setup Guide — Prepare PostgreSQL Schemas](setup_guide.md#prepare-postgresql-schemas-first-run-only).

### Scenario: Update a Single Service

To rebuild and restart just one service after code changes:

```bash
cd stack-dev/platform
make up   # Runs --build automatically
```

---

## 9. Troubleshooting

### "Connection Refused" Errors

**Symptoms**: Frontend can't connect to backend, API calls fail.

**Diagnosis**:
```bash
make ps              # Check what's running
make logs S=backend  # Check backend logs
```

**Common Fixes**:
- Start the backend services: `make up`
- Check `.env` file has correct values
- Verify GraphQL gateway is running

### Port Conflicts

**Symptoms**: Service fails to start with "address already in use" error.

**Diagnosis**:
```bash
lsof -i :3000  # Find what's using a port
```

**Common Fixes**:
- Stop the conflicting process
- Run `make down` then `make up`

### Docker Build Failures

**Symptoms**: `make up` fails during image build.

**Common Fixes**:
- Ensure SSH agent is running: `eval "$(ssh-agent -s)" && ssh-add`
- Check GitHub access: `ssh -T git@github.com`
- Ensure `GH_PAT` is set in `.env`
- Clean Docker cache: `docker system prune -f`

### Database Connection Errors

**Symptoms**: Backend services fail with database errors.

**Diagnosis**:
```bash
docker compose exec postgresql pg_isready -U postgres
docker compose exec postgresql psql -U postgres -c "\dn"
```

**Fix**: Reset the database: `make reset-db`

---

## 10. Maintenance Guidelines

This guide and the `/dev-up` skill (which consolidates the former start-platform + setup-platform) are interconnected documents.

### When You Update This Guide

If you modify the run process:

1. **Update Skills**: Sync changes with `.claude/skills/dev-up/SKILL.md`
2. **Document Issues**: Add troubleshooting entries for new problems

### Continuous Improvement

When you discover:
- Missing verification commands → Add them
- Better startup sequences → Update the guide
- New error scenarios → Add to troubleshooting

The goal is that following this guide reliably starts the platform every time.
