---
name: ant-run
description: Start and manage the Anthropos platform locally with health verification
argument-hint: [scenario or 'full']
---

# Anthropos Platform Run

Execute the Anthropos platform startup by following `corpus/ops/platform-run/run_guide.md` while applying RUN STEP guidelines and verifying service health.

## Your Mission

1. **Follow the guide**: Use `corpus/ops/platform-run/run_guide.md` as your source of truth
2. **Apply RUN STEP guidelines**: Check before start, verify after start, handle conflicts
3. **Track progress locally**: Copy checklist to `anthropos-dev/run_progress.md` and update as you go
4. **Auto-improve docs**: Update run_guide.md when you discover issues or better approaches

## RUN STEP Guidelines (Apply to Every Step)

### 1. Check Before Start
Verify if a service is already running before attempting to start:
```bash
docker ps --filter "name=ant-rosetta" --format "{{.Names}}"
```

### 2. Verify After Start
Confirm service health with explicit checks:
```bash
# Container health
docker ps --filter "name=ant-rosetta-backend" --format "{{.Status}}"

# Service health
curl -s http://localhost:8082/health || echo "Not responding"
```

### 3. Handle Conflicts
If ports are in use, identify and resolve:
```bash
lsof -i :3000
```

### 4. Document Improvements
**Immediately update run_guide.md** when you discover:
- Missing health checks
- Better startup sequences
- New error scenarios
- Environment-specific issues

## Initial Setup

1. Copy checklist: `cp corpus/ops/platform-run/run_checklist.md anthropos-dev/run_progress.md`
2. Read `corpus/ops/platform-run/run_guide.md` to understand the process
3. Ensure setup is complete (or run `/ant-setup` first)
4. Navigate to `anthropos-dev/` workspace

## Startup Sequence

Follow this order for reliable startup:

### Phase 1: Docker Environment
```bash
# Verify Docker is running
docker info > /dev/null 2>&1 && echo "Docker is running" || echo "Start Docker first"

# Check existing containers
docker ps -a --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}"
```

### Phase 2: Infrastructure (PostgreSQL + Redis)
```bash
cd anthropos-dev/platform
docker compose -p ant-rosetta up -d postgresql redis

# Verify
docker exec ant-rosetta-postgresql-1 pg_isready -U postgres
docker exec ant-rosetta-redis-1 redis-cli ping
```

### Phase 3: Backend Services
```bash
# Full stack (recommended)
docker compose -p ant-rosetta --profile graphql up -d

# Or minimal
docker compose -p ant-rosetta up -d sentinel backend cms
```

### Phase 4: Frontend
```bash
cd anthropos-dev/next-web-app
pnpm dev:web
```

### Phase 5: Studio (Optional)
```bash
cd anthropos-dev/studio-desk
npm run dev
```

## Request Confirmation

**ALWAYS ask user before**:
- Starting Docker services
- Starting the frontend development server
- Starting studio services
- Stopping existing services

Use AskUserQuestion tool.

## Error Handling

1. **Do NOT skip errors** - resolve them first
2. Document error message verbatim
3. Check logs: `docker compose -p ant-rosetta logs [service]`
4. Research solution
5. Add to run_guide.md troubleshooting section
6. Continue

## Progress Tracking

Use TodoWrite for high-level tracking:
```markdown
- [x] Copied checklist to anthropos-dev/
- [x] Docker environment verified
- [x] Infrastructure services running
- [ ] Backend services running
- [ ] Frontend accessible
```

Use local checklist (`anthropos-dev/run_progress.md`) for detailed step tracking.

## Common Scenarios

### Resume After Restart
```bash
cd anthropos-dev/platform
docker compose -p ant-rosetta up -d postgresql redis
docker compose -p ant-rosetta --profile graphql up -d
cd ../next-web-app && pnpm dev:web
```

### Quick Frontend Development
```bash
# Backend already running
cd anthropos-dev/next-web-app && pnpm dev:web
```

### Database Reset
```bash
cd anthropos-dev/platform
docker compose -p ant-rosetta down -v
docker compose -p ant-rosetta up -d postgresql redis
# Re-apply migrations
```

## Stopping Services

```bash
# Stop frontend: Ctrl+C

# Stop Docker services (keeps data)
docker compose -p ant-rosetta down

# Stop and remove data (WARNING)
docker compose -p ant-rosetta down -v
```

## Critical Rules

- Work in `anthropos-dev/` scratchpad only
- Use `-p ant-rosetta` for Docker isolation
- Ask before starting/stopping services
- Verify health after every start
- Follow the guide - don't improvise unless needed
- Update docs immediately when improvements found

## Success Criteria

Platform running when:
1. Docker shows healthy containers
2. PostgreSQL accepting connections
3. Redis responding to ping
4. GraphQL gateway at http://localhost:5050
5. Frontend at http://localhost:3000
6. Studio-Desk at http://localhost:3100 (if needed)

## Service URLs

| Service | URL | Health Check |
|---------|-----|--------------|
| Frontend | http://localhost:3000 | Browser loads |
| GraphQL | http://localhost:5050 | `/health` endpoint |
| Backend | http://localhost:8082 | `/health` endpoint |
| Studio-Desk | http://localhost:3100 | Browser loads |
| Directus | http://localhost:8055 | Browser loads |

**Follow `corpus/ops/platform-run/run_guide.md` as your primary reference. Apply these guidelines to start the platform reliably.**
