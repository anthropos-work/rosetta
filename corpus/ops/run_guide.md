# Running the Anthropos Platform Locally

This guide takes you from a **completed setup** to a **running local platform** that you can access in your browser.

> **Prerequisites**: This guide assumes you have completed the [Setup Guide](setup_guide.md). All tools must be installed and repositories cloned to `anthropos-dev/`.

## Quick Reference: Service URLs

Once running, access these URLs in your browser:

| Service | URL | Description |
|---------|-----|-------------|
| **Frontend (Web App)** | http://localhost:3000 | Main user-facing application |
| **Studio-Desk** | http://localhost:9100 | Simulation design tool |
| **GraphQL Playground** | http://localhost:5050 | API gateway (Wundergraph) |
| **Directus CMS** | http://localhost:8055 | Content management (if running) |
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

### Check for Existing Anthropos Containers

```bash
docker ps -a --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
```

This shows any existing containers from previous runs. They may be:
- **Running**: Already up and healthy
- **Exited**: Stopped, can be restarted
- **Not present**: Need to be started fresh

---

## 2. Start Infrastructure Services

The platform requires PostgreSQL and Redis to be running before any backend services.

### Navigate to Platform Directory

```bash
cd anthropos-dev/platform
```

### Start PostgreSQL and Redis

```bash
docker compose -p ant-rosetta up -d postgresql redis
```

### Verify Infrastructure Health

```bash
# Check containers are running
docker ps --filter "name=ant-rosetta-postgresql" --filter "name=ant-rosetta-redis" --format "table {{.Names}}\t{{.Status}}"
```

*Expected*: Both containers showing `Up` with `(healthy)` status.

```bash
# Test PostgreSQL connection
docker exec ant-rosetta-postgresql-1 pg_isready -U postgres
```

*Expected*: `/var/run/postgresql:5432 - accepting connections`

```bash
# Test Redis connection
docker exec ant-rosetta-redis-1 redis-cli ping
```

*Expected*: `PONG`

---

## 3. Start Backend Services

With infrastructure running, start the core backend services.

### Option A: Start All Backend Services (Recommended)

This starts the full backend stack needed for the web application:

```bash
docker compose -p ant-rosetta --profile graphql up -d
```

This starts: `sentinel`, `backend`, `cms`, `skiller`, `skillpath`, `storage`, `chronos`, `jobsimulation`, `intelligence`, and `graphql`.

*Note*: First run may take several minutes as Docker builds images from source.

### Option B: Start Minimal Backend (Faster)

For a lighter setup, start only essential services:

```bash
docker compose -p ant-rosetta up -d sentinel backend cms
```

*Note*: Some frontend features may not work without the full stack.

### Verify Backend Services

```bash
docker ps --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
```

*Expected*: All started containers showing `Up` status.

### Check Service Logs (if issues)

```bash
# View logs for a specific service
docker compose -p ant-rosetta logs -f backend

# View logs for all services
docker compose -p ant-rosetta logs -f
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

Studio-Desk is the simulation design tool used by content creators. It runs alongside the frontend.

### Navigate to Studio-Desk Directory

```bash
cd anthropos-dev/studio-desk
```

### Install Dependencies (if needed)

```bash
ls node_modules > /dev/null 2>&1 && echo "Dependencies installed" || npm install
```

### Verify Environment File

Studio-Desk requires its own `.env` file with Clerk and OpenAI credentials.

```bash
ls .env > /dev/null 2>&1 && echo ".env exists" || echo "Missing: Copy .env.example to .env"
```

If missing, create it:
```bash
cp .env.example .env
```

Then populate the required keys from `platform/.env`:
- `CLERK_SECRET_KEY` and `CLERK_PUBLISHABLE_KEY` (copy from platform)
- `OPENAI_API_KEY` (copy from platform)

### Start Studio-Desk

```bash
npm run dev
```

*Expected*:
- Frontend starts on http://localhost:9100 (configurable via `FRONTEND_PORT` in `.env`)
- Backend starts on http://localhost:9000 (configurable via `PORT` or `BACKEND_PORT` in `.env`)

### Verify Studio-Desk

Open http://localhost:9100 in your browser.

*Expected*: Studio-Desk login page (uses Clerk authentication).

### Troubleshooting: Port Already in Use

```bash
# Find what's using port 9100 or 9000
lsof -i :9100
lsof -i :9000
```

To use different ports, edit `studio-desk/.env`:
```
FRONTEND_PORT=9200
PORT=9001
```

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

### Stop Studio-Desk
Press `Ctrl+C` in the terminal running `npm run dev`.

### Stop Docker Services

```bash
cd anthropos-dev/platform

# Stop all Anthropos services (keeps data)
docker compose -p ant-rosetta down

# Stop and remove volumes (WARNING: deletes database data)
docker compose -p ant-rosetta down -v
```

### Check All Services Stopped

```bash
docker ps --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}"
```

*Expected*: No containers listed, or all showing `Exited`.

---

## 8. Common Scenarios

### Scenario: Resume After Computer Restart

1. Start Docker Desktop
2. Start infrastructure: `docker compose -p ant-rosetta up -d postgresql redis`
3. Start backend: `docker compose -p ant-rosetta --profile graphql up -d`
4. Start frontend: `cd next-web-app && pnpm dev:web`
5. Start Studio-Desk: `cd studio-desk && npm run dev`

### Scenario: Quick Frontend Development

If you only need to work on the frontend and backend is already running:

```bash
# Verify backend is up
docker ps --filter "name=ant-rosetta-backend" --format "{{.Status}}"

# Start frontend
cd anthropos-dev/next-web-app && pnpm dev:web
```

### Scenario: Database Reset

To wipe the database and start fresh:

```bash
cd anthropos-dev/platform

# Stop everything and remove volumes
docker compose -p ant-rosetta down -v

# Start fresh
docker compose -p ant-rosetta up -d postgresql redis

# Re-apply migrations
(cd ../backend && atlas migrate apply --env local)
(cd ../cms && atlas migrate apply --env local)
(cd ../jobsimulation && atlas migrate apply --env local)
(cd ../skiller && atlas migrate apply --env local)

# Start services
docker compose -p ant-rosetta --profile graphql up -d
```

### Scenario: Update a Single Service

To rebuild and restart just one service:

```bash
docker compose -p ant-rosetta up -d --build backend
```

---

## 9. Troubleshooting

### "Connection Refused" Errors

**Symptoms**: Frontend can't connect to backend, API calls fail.

**Diagnosis**:
```bash
# Check if backend is running
docker ps --filter "name=ant-rosetta-backend"

# Check backend logs
docker compose -p ant-rosetta logs backend
```

**Common Fixes**:
- Start the backend services (see Section 3)
- Check `.env` file has correct values
- Verify GraphQL gateway is running

### Port Conflicts

**Symptoms**: Service fails to start with "address already in use" error.

**Diagnosis**:
```bash
# Find what's using a port (example: 3000)
lsof -i :3000
```

**Common Fixes**:
- Stop the conflicting process
- Use a different port
- Stop other Anthropos environments

### Docker Build Failures

**Symptoms**: `docker compose up` fails during image build.

**Diagnosis**:
```bash
# Check Docker build logs
docker compose -p ant-rosetta logs --tail=100
```

**Common Fixes**:
- Ensure SSH agent is running: `eval "$(ssh-agent -s)" && ssh-add`
- Check GitHub access: `ssh -T git@github.com`
- Clean Docker cache: `docker system prune -f`

### "Module Not Found" in Frontend

**Symptoms**: Next.js fails to start with import errors.

**Fix**:
```bash
cd anthropos-dev/next-web-app
pnpm install
```

### Database Connection Errors

**Symptoms**: Backend services fail with database errors.

**Diagnosis**:
```bash
# Check PostgreSQL status
docker exec ant-rosetta-postgresql-1 pg_isready -U postgres

# Check if schemas exist
docker exec ant-rosetta-postgresql-1 psql -U postgres -c "\dn"
```

**Fix**: Re-apply migrations (see "Database Reset" scenario above).

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
