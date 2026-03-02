# Running the Anthropos Platform Locally

This guide takes you from a **completed setup** to a **running local platform** that you can access in your browser.

> **Prerequisites**: This guide assumes you have completed the [Setup Guide](setup_guide.md). All tools must be installed and repositories cloned to `anthropos-dev/`.

## Quick Reference: Service URLs

Once running, access these URLs in your browser:

| Service | URL | Description |
|---------|-----|-------------|
| **Frontend (Web App)** | http://localhost:3000 | Main user-facing application |
| **Studio-Desk** | http://localhost:9100 | Simulation design tool |
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
cd anthropos-dev/platform
make ps
```

This shows any existing containers from previous runs.

---

## 2. Start Backend Services

The platform uses a **Makefile** as the single entry point. Infrastructure (PostgreSQL, Redis) and Sentinel start automatically with any profile.

### Navigate to Platform Directory

```bash
cd anthropos-dev/platform
```

### Option A: Start Full Backend (Recommended)

This starts all backend services + GraphQL router (default `graphql` profile):

```bash
make up
```

This starts: PostgreSQL, Redis, Sentinel, Backend, CMS, Skiller, Skillpath, Storage, Chronos, Jobsimulation, Intelligence, Roadrunner, and GraphQL/Cosmo Router.

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

The frontend runs outside Docker for faster development iteration.

### Verify Node.js Version (Critical)

The frontend requires **Node.js v20 LTS**. Node.js v22+ removed the `import ... assert` syntax used by the codebase.

```bash
node --version
```

*Expected*: `v20.x.x`

**If using Node.js v22+ or v24+**: Switch to Node 20 using nvm:
```bash
nvm use 20
```

If nvm is not installed or Node 20 is not available:
```bash
nvm install 20
nvm use 20
```

### Navigate to Frontend Directory

```bash
cd anthropos-dev/next-web-app
```

### Install Dependencies (if needed)

Check if `node_modules` exists:
```bash
ls node_modules > /dev/null 2>&1 && echo "Dependencies installed" || echo "Run: pnpm install"
```

If dependencies are missing:
```bash
pnpm install
```

### Option A: Development Mode (Recommended)

Hot-reloading enabled, faster iteration:
```bash
pnpm dev:web
```

*Expected*: Server starts on http://localhost:3000

### Option B: Production Mode

Build first, then serve (faster runtime, no hot-reload):
```bash
pnpm build:web && pnpm start:web
```

### Verify Frontend

Open http://localhost:3000 in your browser.

*Expected*: Anthropos login page or dashboard (depending on auth state).

### Troubleshooting: Port Already in Use

```bash
# Find what's using port 3000
lsof -i :3000
```

Options:
1. Stop the conflicting process
2. Use a different port: `PORT=3001 pnpm dev:web`

---

## 5. Start Studio-Desk (Design Tool)

Studio-Desk is the simulation design tool used by content creators.

### Option A: Run in Docker (Recommended)

```bash
cd anthropos-dev/platform
make up PROFILE=studio-desk
```

This starts Studio-Desk in Docker along with its dependencies (GraphQL, CMS).

### Option B: Run Natively

For development with hot-reloading:

```bash
cd anthropos-dev/studio-desk
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
cd anthropos-dev/studio-room
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

### Stop Frontend
Press `Ctrl+C` in the terminal running `pnpm dev:web`.

### Stop Studio-Desk (Native)
Press `Ctrl+C` in the terminal running `npm run dev`.

### Stop Docker Services

```bash
cd anthropos-dev/platform

# Stop all services (keeps data)
make down
```

### Check All Services Stopped

```bash
make ps
```

*Expected*: No containers listed.

---

## 8. Common Scenarios

### Scenario: Resume After Computer Restart

1. Start Docker Desktop
2. Start backend: `cd anthropos-dev/platform && make up`
3. Start frontend: `cd anthropos-dev/next-web-app && pnpm dev:web`
4. (Optional) Start Studio-Desk: `cd anthropos-dev/studio-desk && npm run dev`

### Scenario: Quick Frontend Development

If you only need to work on the frontend and backend is already running:

```bash
# Verify backend is up
cd anthropos-dev/platform && make ps

# Start frontend
cd ../next-web-app && pnpm dev:web
```

### Scenario: Develop a Single Service Natively

To stop a service's Docker container and run it locally:

```bash
cd anthropos-dev/platform
make dev S=cms           # Stops the Docker container for CMS
cd ../cms
go run .                 # Run natively with hot-reload
```

### Scenario: Database Reset

To wipe the database and start fresh:

```bash
cd anthropos-dev/platform
make reset-db
```

This removes PostgreSQL data, restarts the container, and re-runs all migrations automatically.

### Scenario: Update a Single Service

To rebuild and restart just one service after code changes:

```bash
cd anthropos-dev/platform
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

This guide and the `/ant-run` skill are interconnected documents.

### When You Update This Guide

If you modify the run process:

1. **Update Skills**: Sync changes with `.claude/skills/ant-run/SKILL.md`
2. **Document Issues**: Add troubleshooting entries for new problems

### Continuous Improvement

When you discover:
- Missing verification commands → Add them
- Better startup sequences → Update the guide
- New error scenarios → Add to troubleshooting

The goal is that following this guide reliably starts the platform every time.
