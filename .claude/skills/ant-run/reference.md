# Anthropos Run Skill - Technical Reference

## Architecture

This skill implements a **start-and-verify** loop:
1. Check current state of services
2. Start services in dependency order
3. Verify health after each start
4. Document issues and improvements

## File References

### Documentation Sources
- `corpus/ops/platform-run/run_guide.md` - Master run guide
- `corpus/ops/platform-run/run_checklist.md` - Progress tracking checklist
- `corpus/architecture/` - Platform architecture context

### Working Directory
- `anthropos-dev/` - Git-ignored scratchpad for all platform activities

## Service Dependency Order

```
PostgreSQL → Redis → Sentinel → Backend → CMS → Other Core Services → GraphQL → Frontend → Studio
```

### Startup Graph

```
               ┌──────────┐
               │PostgreSQL│
               └────┬─────┘
                    │
               ┌────▼─────┐
               │  Redis   │
               └────┬─────┘
                    │
          ┌─────────┼─────────┐
          │         │         │
     ┌────▼───┐ ┌───▼────┐ ┌──▼──┐
     │Sentinel│ │Backend │ │ CMS │
     └────┬───┘ └───┬────┘ └──┬──┘
          │         │         │
          └─────────┼─────────┘
                    │
        ┌───────────┼───────────┐
        │           │           │
   ┌────▼───┐  ┌────▼───┐  ┌────▼────┐
   │Skiller │  │Skillpath│  │Chronos  │
   └────────┘  └─────────┘  └─────────┘
        │           │           │
        └───────────┼───────────┘
                    │
               ┌────▼────┐
               │ GraphQL │
               └────┬────┘
                    │
               ┌────▼────┐
               │Frontend │
               └─────────┘
```

## Tool Usage Strategy

### Read-Only Tools (Verification)
- `Bash` for health checks (`docker ps`, `curl`, `pg_isready`)
- `Read` for documentation and logs

### Write Tools (Execution & Documentation)
- `Bash` for starting services (after user confirmation)
- `Edit` for updating documentation (auto-improvement)

### Interactive Tools
- `AskUserQuestion` for all confirmations before:
  - Starting Docker services
  - Starting frontend dev server
  - Stopping services
  - Data-destructive operations

## RUN STEP Guidelines Implementation

### Guideline 1: Check Before Start

```bash
# Check if Docker is running
docker info > /dev/null 2>&1 && echo "Docker is running" || echo "Docker NOT running"

# Check existing containers
docker ps -a --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}"

# Check port availability
lsof -i :3000 || echo "Port 3000 available"
```

### Guideline 2: Verify After Start

```bash
# Container health
docker ps --filter "name=ant-rosetta-backend" --format "{{.Status}}"
# Expected: "Up X seconds (healthy)" or "Up X minutes"

# PostgreSQL
docker exec ant-rosetta-postgresql-1 pg_isready -U postgres
# Expected: "/var/run/postgresql:5432 - accepting connections"

# Redis
docker exec ant-rosetta-redis-1 redis-cli ping
# Expected: "PONG"

# HTTP health check
curl -s http://localhost:8082/health || echo "Backend not responding"
curl -s http://localhost:5050/health || echo "GraphQL not responding"
```

### Guideline 3: Handle Conflicts

```bash
# Find what's using a port
lsof -i :3000

# Kill process (after user confirmation)
kill -9 <PID>

# Or change port
PORT=3001 pnpm dev:web
```

### Guideline 4: Document Improvements

```typescript
// Pattern: Update guide immediately
Edit({
  file_path: "corpus/ops/platform-run/run_guide.md",
  old_string: "docker compose -p ant-rosetta up -d",
  new_string: "docker compose -p ant-rosetta up -d\n\n*Verification*: `docker ps --filter \"name=ant-rosetta\"`"
})
```

## Docker Commands Reference

### Start Commands

```bash
# Infrastructure only
docker compose -p ant-rosetta up -d postgresql redis

# Full backend stack
docker compose -p ant-rosetta --profile graphql up -d

# Minimal backend
docker compose -p ant-rosetta up -d sentinel backend cms

# Specific service
docker compose -p ant-rosetta up -d backend
```

### Status Commands

```bash
# All containers
docker ps --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

# Service logs
docker compose -p ant-rosetta logs -f backend

# All logs
docker compose -p ant-rosetta logs -f
```

### Stop Commands

```bash
# Stop (keep data)
docker compose -p ant-rosetta down

# Stop and remove volumes (data loss!)
docker compose -p ant-rosetta down -v

# Stop single service
docker compose -p ant-rosetta stop backend
```

## Health Check Patterns

### Pattern 1: PostgreSQL Health

```bash
# Check container
docker ps --filter "name=ant-rosetta-postgresql" --format "{{.Status}}"

# Check connections
docker exec ant-rosetta-postgresql-1 pg_isready -U postgres

# Check schemas
docker exec ant-rosetta-postgresql-1 psql -U postgres -c "\dn"
```

### Pattern 2: Redis Health

```bash
# Check container
docker ps --filter "name=ant-rosetta-redis" --format "{{.Status}}"

# Ping
docker exec ant-rosetta-redis-1 redis-cli ping

# Info
docker exec ant-rosetta-redis-1 redis-cli info
```

### Pattern 3: Backend Service Health

```bash
# Check container
docker ps --filter "name=ant-rosetta-backend" --format "{{.Status}}"

# HTTP health
curl -s http://localhost:8082/health

# Logs if unhealthy
docker compose -p ant-rosetta logs --tail=50 backend
```

### Pattern 4: Frontend Health

```bash
# Check process
pgrep -f "next-web-app" || echo "Frontend not running"

# HTTP check
curl -s http://localhost:3000 > /dev/null && echo "Frontend responding" || echo "Frontend not responding"
```

## Error Recovery Patterns

### Pattern 1: Container Won't Start

```yaml
Symptoms: Container exits immediately
Diagnosis:
  1. Check logs: docker compose -p ant-rosetta logs backend
  2. Look for: missing env vars, connection errors, file not found
Recovery:
  - Missing env: Check platform/.env
  - Connection error: Ensure dependencies started first
  - File not found: Rebuild: docker compose -p ant-rosetta build backend
```

### Pattern 2: Port Already in Use

```yaml
Symptoms: "bind: address already in use"
Diagnosis:
  1. Find process: lsof -i :3000
  2. Identify: Is it another Anthropos env? Another app?
Recovery:
  - Kill process: kill -9 <PID>
  - Use different port: PORT=3001 pnpm dev:web
  - Stop other env: docker compose -p other-project down
```

### Pattern 3: Database Connection Failed

```yaml
Symptoms: "connection refused" to PostgreSQL
Diagnosis:
  1. Check container: docker ps --filter "name=postgresql"
  2. Check readiness: docker exec ant-rosetta-postgresql-1 pg_isready -U postgres
Recovery:
  - Not running: docker compose -p ant-rosetta up -d postgresql
  - Not ready: Wait 10-30 seconds, retry
  - Wrong config: Check DATABASE_URL in .env
```

### Pattern 4: Frontend Can't Connect to Backend

```yaml
Symptoms: "Network error" or "Connection refused" in browser
Diagnosis:
  1. Check backend: curl http://localhost:8082/health
  2. Check GraphQL: curl http://localhost:5050/health
Recovery:
  - Backend down: Start backend services
  - CORS issue: Check .env configuration
  - Wrong URL: Verify NEXT_PUBLIC_* env vars
```

## Progress Tracking

### Two-Level Tracking System

**Level 1: Local Checklist** (`anthropos-dev/run_progress.md`)
- Copied from run_checklist.md at start
- Updated as each service starts
- Contains issues encountered

**Level 2: TodoWrite Tool**
- High-level phase tracking
- Blockers requiring user input

Example TodoWrite tracking:
```markdown
- [x] Phase 1: Docker environment verified
- [x] Phase 2: PostgreSQL started and healthy
- [x] Phase 2: Redis started and healthy
- [ ] Phase 3: Backend services starting
- [ ] Phase 4: Frontend accessible
```

## Scenario Handlers

### Scenario: Fresh Start

```typescript
// Full startup sequence
1. Check Docker: docker info
2. Start infra: docker compose -p ant-rosetta up -d postgresql redis
3. Wait for health: pg_isready, redis-cli ping
4. Start backend: docker compose -p ant-rosetta --profile graphql up -d
5. Wait for health: curl health endpoints
6. Start frontend: pnpm dev:web
7. Verify: Open http://localhost:3000
```

### Scenario: Resume After Restart

```typescript
// Abbreviated startup
1. Check Docker: docker info
2. Check containers: docker ps --filter "name=ant-rosetta"
3. If stopped: docker compose -p ant-rosetta up -d
4. If running: Skip to frontend
5. Start frontend: pnpm dev:web
```

### Scenario: Just Frontend

```typescript
// Frontend-only
1. Verify backend: curl http://localhost:5050/health
2. If healthy: cd next-web-app && pnpm dev:web
3. If not: Start backend first
```

## Success Validation

At the end of startup, verify:

1. **Infrastructure**: `docker exec ant-rosetta-postgresql-1 pg_isready` returns success
2. **Backend**: `curl http://localhost:8082/health` returns 200
3. **GraphQL**: `curl http://localhost:5050/health` returns 200
4. **Frontend**: Browser loads http://localhost:3000

## Skill Invocation

```bash
# Full startup
/ant-run
/ant-run full

# Specific scenario
/ant-run resume        # After computer restart
/ant-run frontend      # Just frontend (backend assumed running)
/ant-run minimal       # Minimal backend only

# With target
/ant-run stop          # Stop all services
/ant-run status        # Check what's running
```

## Integration with Project Rosetta

This skill complements:
- `/ant-setup` - Initial environment build (run this first)
- `/ant-update` - Sync code and dependencies before running
- `/ant-integrate` - Document issues found during run

The skill follows the **Recreation Standard**: anyone following the run_guide.md should be able to start the platform reliably.
