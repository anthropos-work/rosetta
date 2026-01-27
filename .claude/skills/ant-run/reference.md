# Anthropos Run Skill - Technical Reference

Quick reference for health checks, error recovery, and Docker commands. For full startup instructions, see `corpus/ops/run_guide.md`.

## File References

| Type | Path |
|------|------|
| Master Guide | `corpus/ops/run_guide.md` |
| Ops Reports | `anthropos-dev/ops-reports/` |
| Working Dir | `anthropos-dev/` |

## Service Dependency Graph

```
               PostgreSQL
                    │
                  Redis
                    │
          ┌─────────┼─────────┐
          │         │         │
       Sentinel  Backend     CMS
          │         │         │
          └─────────┼─────────┘
                    │
        ┌───────────┼───────────┐
        │           │           │
     Skiller   Skillpath    Chronos
        │           │           │
        └───────────┼───────────┘
                    │
                GraphQL
                    │
                Frontend
```

## Quick Health Checks

```bash
# Docker running?
docker info > /dev/null 2>&1 && echo "OK" || echo "Start Docker"

# PostgreSQL ready?
docker exec ant-rosetta-postgresql-1 pg_isready -U postgres

# Redis responding?
docker exec ant-rosetta-redis-1 redis-cli ping

# Backend health?
curl -s http://localhost:8082/health

# GraphQL health?
curl -s http://localhost:5050/health

# Container status
docker ps --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}"
```

## Common Docker Commands

```bash
# Start infrastructure
docker compose -p ant-rosetta up -d postgresql redis

# Start full backend
docker compose -p ant-rosetta --profile graphql up -d

# Start minimal backend
docker compose -p ant-rosetta up -d sentinel backend cms

# View logs
docker compose -p ant-rosetta logs -f [service]

# Stop (keep data)
docker compose -p ant-rosetta down

# Stop and delete data
docker compose -p ant-rosetta down -v
```

## Error Recovery Patterns

### Container Won't Start

```bash
# Check logs
docker compose -p ant-rosetta logs backend

# Common causes:
# - Missing env vars → Check platform/.env
# - Connection error → Start dependencies first
# - File not found → Rebuild: docker compose -p ant-rosetta build backend
```

### Port Already in Use

```bash
# Find what's using the port
lsof -i :3000

# Options:
# - Kill process: kill -9 <PID> (ask user first)
# - Use different port: PORT=3001 pnpm dev:web
```

### Database Connection Failed

```bash
# Check if PostgreSQL container is running
docker ps --filter "name=ant-rosetta-postgresql"

# Check if accepting connections
docker exec ant-rosetta-postgresql-1 pg_isready -U postgres

# If not running, start it:
docker compose -p ant-rosetta up -d postgresql
```

### Frontend Can't Connect to Backend

```bash
# Verify backend is healthy
curl http://localhost:8082/health
curl http://localhost:5050/health

# Check backend logs
docker compose -p ant-rosetta logs backend
```

## Progress Tracking Template

```typescript
TodoWrite({
  todos: [
    { content: "Verify Docker environment", status: "in_progress", activeForm: "Verifying Docker" },
    { content: "Start infrastructure (PostgreSQL, Redis)", status: "pending", activeForm: "Starting infrastructure" },
    { content: "Start backend services", status: "pending", activeForm: "Starting backend" },
    { content: "Verify GraphQL gateway", status: "pending", activeForm: "Verifying GraphQL" },
    { content: "Start frontend", status: "pending", activeForm: "Starting frontend" },
    { content: "Verify platform accessible", status: "pending", activeForm: "Verifying platform" }
  ]
})
```

## Service URLs

| Service | URL | Health |
|---------|-----|--------|
| Frontend | http://localhost:3000 | Browser loads |
| GraphQL | http://localhost:5050 | `/health` |
| Backend | http://localhost:8082 | `/health` |
| Studio-Desk | http://localhost:9100 | Browser loads |
| Directus | http://localhost:8055 | Browser loads |

## Ops Report Template

When creating `anthropos-dev/ops-reports/op_YYYYMMDD_HHMMSS_run_<topic>.md`:

```markdown
# Ops Report: [Brief Title]

**Date**: YYYY-MM-DD HH:MM
**Skill**: /ant-run
**OS**: [macOS 14.x / Ubuntu 22.04 / etc.]
**Phase**: [Docker / Infrastructure / Backend / Frontend]

## Issue Encountered
[Exact error message]

## Context
[What was being done, what commands ran]

## Resolution
[How fixed, or "Unresolved"]

## Suggested Documentation Update
[What to add/change in run_guide.md]
```

## Related Skills

| Skill | Use When |
|-------|----------|
| `/ant-setup` | First-time environment setup |
| `/ant-update` | Sync code/deps before running |
| `/ant-integrate` | Process ops-reports into corpus |
