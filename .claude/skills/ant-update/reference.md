# Anthropos Update Skill - Technical Reference

Quick reference for update commands and error recovery. For full update instructions, see `corpus/ops/update_guide.md`.

## File References

| Type | Path |
|------|------|
| Master Guide | `corpus/ops/update_guide.md` |
| Ops Reports | `anthropos-dev/ops-reports/` |
| Working Dir | `anthropos-dev/` |

## Update Scenarios

| Scenario | Use When | Scope |
|----------|----------|-------|
| Daily | Starting daily work | Main repos, frontend deps |
| Weekly | Weekly maintenance, after being away | All repos, all deps, migrations |
| Full | Breaking changes, major version bumps | Fresh pull, clean install, rebuild |

## Quick State Check

```bash
# Check running services
docker ps --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}"

# Check git status across repos
cd anthropos-dev
for repo in platform backend cms jobsimulation skiller next-web-app; do
  echo "=== $repo ==="
  (cd "$repo" && git status -s) 2>/dev/null || echo "  Not found"
done

# Check for pending migrations
(cd backend && atlas migrate status --env local)
```

## Common Update Commands

```bash
# Stop services before pulling
cd anthropos-dev/platform
docker compose -p ant-rosetta down

# Pull all repos
cd anthropos-dev
for repo in platform backend cms jobsimulation skiller next-web-app studio-desk studio-room cms/studio; do
  (cd "$repo" 2>/dev/null && git pull origin main) || echo "Skipped: $repo"
done

# Update dependencies
(cd next-web-app && pnpm install)
(cd studio-desk && npm install)
(cd studio-room && pip3 install -r requirements.txt --upgrade)

# Apply migrations (ensure PostgreSQL is running)
docker compose -p ant-rosetta up -d postgresql && sleep 5
(cd backend && atlas migrate apply --env local)
(cd cms && atlas migrate apply --env local)
(cd jobsimulation && atlas migrate apply --env local)
(cd skiller && atlas migrate apply --env local)

# Rebuild and start
cd platform
docker compose -p ant-rosetta --profile graphql up -d --build
```

## Error Recovery Patterns

### Git Conflict During Pull

```bash
# Option 1: Stash local changes
git stash
git pull origin main
git stash pop

# Option 2: Discard local changes (ask user first)
git fetch origin
git reset --hard origin/main
```

### Package Not Found After Pull

```bash
# Frontend
cd next-web-app && pnpm install

# Studio-Desk
cd studio-desk && npm install

# If still failing
rm -rf node_modules && pnpm install
```

### Migration Failed

```bash
# Check migration status
atlas migrate status --env local

# If dirty state, may need fresh database
docker compose -p ant-rosetta down -v
docker compose -p ant-rosetta up -d postgresql && sleep 5
(cd backend && atlas migrate apply --env local)
```

### Docker Build Fails

```bash
# Clear cache
docker system prune -f

# Rebuild without cache
docker compose -p ant-rosetta build --no-cache
```

### Service Won't Start After Update

```bash
# Check logs
docker compose -p ant-rosetta logs [service]

# Common causes:
# - Missing env vars → Check platform/.env
# - Schema mismatch → Apply pending migrations
# - Stale image → Rebuild: docker compose -p ant-rosetta build [service]
```

## Progress Tracking Template

```typescript
TodoWrite({
  todos: [
    { content: "Stop running services", status: "in_progress", activeForm: "Stopping services" },
    { content: "Update repository code", status: "pending", activeForm: "Pulling repositories" },
    { content: "Install dependencies", status: "pending", activeForm: "Installing dependencies" },
    { content: "Apply database migrations", status: "pending", activeForm: "Applying migrations" },
    { content: "Rebuild Docker images", status: "pending", activeForm: "Rebuilding Docker" },
    { content: "Verify services healthy", status: "pending", activeForm: "Verifying health" }
  ]
})
```

## Ops Report Template

When creating `anthropos-dev/ops-reports/op_YYYYMMDD_HHMMSS_update_<topic>.md`:

```markdown
# Ops Report: [Brief Title]

**Date**: YYYY-MM-DD HH:MM
**Skill**: /ant-update
**OS**: [macOS 14.x / Ubuntu 22.04 / etc.]
**Phase**: [Stop / Pull / Dependencies / Migrations / Rebuild / Verify]

## Issue Encountered
[Exact error message]

## Context
[What was being done, what commands ran]

## Resolution
[How fixed, or "Unresolved"]

## Suggested Documentation Update
[What to add/change in update_guide.md]
```

## Repository Update Order

```
1. platform      (Docker configs, shared .env)
2. backend       (main API)
3. cms           (content management)
4. jobsimulation
5. skiller
6. next-web-app  (frontend)
7. studio-desk
8. studio-room
9. cms/studio    (symlinked studio-room)
```

## Related Skills

| Skill | Use When |
|-------|----------|
| `/ant-setup` | First-time environment setup |
| `/ant-run` | Start platform after updating |
| `/ant-integrate` | Process ops-reports into corpus |
