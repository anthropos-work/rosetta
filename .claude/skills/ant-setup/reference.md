# Anthropos Setup Skill - Technical Reference

Quick reference for verification commands and error recovery. For full setup instructions, see `corpus/ops/setup_guide.md`.

## File References

| Type | Path |
|------|------|
| Master Guide | `corpus/ops/setup_guide.md` |
| Ops Reports | `anthropos-dev/ops-reports/` |
| Working Dir | `anthropos-dev/` |

## Verification Commands

```bash
# Prerequisites
git --version
docker --version && docker compose version
go version
node --version && pnpm --version
python3 --version
atlas version
xcode-select -p  # macOS only

# GitHub SSH
ssh -T git@github.com

# Docker services
docker ps --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}"

# Database
docker exec ant-rosetta-postgresql-1 pg_isready -U postgres
docker exec ant-rosetta-redis-1 redis-cli ping

# Services
curl -s http://localhost:3000 > /dev/null && echo "Frontend OK"
curl -s http://localhost:5050/health && echo "GraphQL OK"
curl -s http://localhost:8082/health && echo "Backend OK"
```

## Common Docker Commands

```bash
# Start infrastructure
docker compose -p ant-rosetta up -d postgresql redis

# Prepare PostgreSQL extensions (before migrations)
docker exec ant-rosetta-postgresql-1 psql -U postgres -c "CREATE SCHEMA IF NOT EXISTS extensions; CREATE EXTENSION IF NOT EXISTS vector SCHEMA extensions; CREATE EXTENSION IF NOT EXISTS pg_trgm SCHEMA extensions;"
docker exec ant-rosetta-postgresql-1 psql -U postgres -c "CREATE SCHEMA IF NOT EXISTS sentinel;"

# Start full backend
docker compose -p ant-rosetta --profile graphql up -d

# View logs
docker compose -p ant-rosetta logs -f [service]
```

## Atlas Migrations

```bash
# From anthropos-dev/
(cd backend && atlas migrate apply --env local)
(cd cms && atlas migrate apply --env local)
(cd jobsimulation && atlas migrate apply --env local)
(cd skiller && atlas migrate apply --env local)
(cd skillpath && atlas migrate apply --env local)
(cd chronos && atlas migrate apply --env local)
```

## Error Recovery Patterns

### Permission Denied (Docker on Linux)

```bash
# Check group membership
groups | grep docker

# Fix
sudo usermod -aG docker $USER
newgrp docker

# Verify
docker ps
```

### Port Already in Use

```bash
# Find what's using the port
lsof -i :3000

# Options:
# - Kill process: kill -9 <PID>
# - Change port: PORT=3001 pnpm dev
```

### Missing pgvector Extension

```bash
# Create extensions schema
docker exec ant-rosetta-postgresql-1 psql -U postgres -c "CREATE SCHEMA IF NOT EXISTS extensions; CREATE EXTENSION IF NOT EXISTS vector SCHEMA extensions;"

# Retry migration
(cd cms && atlas migrate apply --env local)
```

### Docker Build Fails (SSH)

```bash
# Check SSH agent
ssh-add -l

# If no identities, add key
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_ed25519

# Verify GitHub access
ssh -T git@github.com
```

## Ops Report Template

When creating `anthropos-dev/ops-reports/op_YYYYMMDD_HHMMSS_setup_<topic>.md`:

```markdown
# Ops Report: [Brief Title]

**Date**: YYYY-MM-DD HH:MM
**Skill**: /ant-setup
**OS**: [macOS 14.x / Ubuntu 22.04 / etc.]
**Phase**: [Prerequisites / Repos / Docker / Frontend / etc.]

## Issue Encountered
[Exact error message]

## Context
[What was being done, what commands ran]

## Resolution
[How fixed, or "Unresolved"]

## Suggested Documentation Update
[What to add/change in setup_guide.md]
```

## Repository Structure Post-Setup

```
anthropos-dev/
├── platform/           # Docker config + .env
├── backend/            # Main API (cloned from app)
├── cms/
│   └── studio/         # Studio-room for Docker build
├── jobsimulation/
├── skiller/
├── next-web-app/       # Frontend
├── studio-desk/
├── studio-room/
└── ops-reports/        # Operational feedback
```

## Related Skills

| Skill | Use When |
|-------|----------|
| `/ant-run` | Start platform after setup complete |
| `/ant-update` | Sync code/deps after initial setup |
| `/ant-integrate` | Process ops-reports into corpus |
