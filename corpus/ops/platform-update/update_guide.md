# Updating the Anthropos Platform

This guide covers keeping your local development environment **up-to-date** with the latest code, dependencies, and database schemas.

> **Prerequisites**: This guide assumes you have completed the [Platform Setup Guide](../platform-setup/setup_guide.md) and have a working local environment.

> **Companion Checklist**: Copy [update_checklist.md](./update_checklist.md) to your workspace to track progress.

## When to Update

Run through this guide when:
- Starting work after being away (daily/weekly sync)
- A teammate mentions breaking changes
- You see dependency errors or schema mismatches
- Before starting work on a new feature

---

## Update Execution Guidelines (UPDATE STEP)

When executing each step:

1. **Check Current State**: Verify what needs updating before making changes.
2. **Pull Before Build**: Always fetch latest code before rebuilding.
3. **Handle Conflicts**: If git conflicts occur, resolve before proceeding.
4. **Verify After Update**: Confirm services still work after updates.

---

## 1. Stop Running Services

Before updating, stop any running services to avoid conflicts.

### Check What's Running

```bash
# Check Docker services
docker ps --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}"

# Check for running Node processes (frontend)
pgrep -f "next-web-app" || echo "No frontend running"
```

### Stop Services

```bash
# Stop Docker services
cd anthropos-dev/platform
docker compose -p ant-rosetta down

# Stop frontend (Ctrl+C in terminal, or)
pkill -f "pnpm dev:web" 2>/dev/null || true
```

---

## 2. Update Repository Code

Pull latest changes from all repositories.

### Navigate to Workspace

```bash
cd anthropos-dev
```

### Update All Repositories

Run these commands to pull latest code:

```bash
# Platform configuration
(cd platform && git pull origin main)

# Backend services
(cd backend && git pull origin main)
(cd cms && git pull origin main)
(cd jobsimulation && git pull origin main)
(cd skiller && git pull origin main)

# Frontend
(cd next-web-app && git pull origin main)

# Studio services
(cd studio-desk && git pull origin main)
(cd studio-room && git pull origin main)

# CMS studio dependency (keep in sync with studio-room)
(cd cms/studio && git pull origin main)
```

*Note*: The parentheses `(...)` ensure you return to the current directory after each command.

### Quick Update Script

For convenience, you can run all pulls in sequence:

```bash
for repo in platform backend cms jobsimulation skiller next-web-app studio-desk studio-room cms/studio; do
  echo "Updating $repo..."
  (cd "$repo" 2>/dev/null && git pull origin main) || echo "  Skipped: $repo not found"
done
```

### Handle Git Conflicts

If you see merge conflicts:

```bash
cd [repo-with-conflict]
git status  # See conflicting files

# Option 1: Stash your changes and pull
git stash
git pull origin main
git stash pop  # Reapply your changes

# Option 2: Reset to remote (WARNING: loses local changes)
git fetch origin
git reset --hard origin/main
```

---

## 3. Update Dependencies

After pulling code, install any new or updated dependencies.

### Frontend (Next.js Monorepo)

```bash
cd anthropos-dev/next-web-app
pnpm install
```

*Verification*: No errors during install, `node_modules` updated.

### Studio-Desk

```bash
cd anthropos-dev/studio-desk
npm install
```

### Studio-Room (Python)

```bash
cd anthropos-dev/studio-room
pip3 install -r requirements.txt --upgrade
```

### Go Services (if building locally)

Go modules update automatically on build, but you can force an update:

```bash
cd anthropos-dev/backend
go mod download
go mod tidy
```

---

## 4. Run Database Migrations

If schema changes were made, apply new migrations.

### Check Migration Status

```bash
# Ensure PostgreSQL is running
docker compose -p ant-rosetta up -d postgresql
sleep 5  # Wait for PostgreSQL to be ready
```

### Apply Migrations

```bash
cd anthropos-dev

# Backend schema
(cd backend && atlas migrate apply --env local)

# CMS schema
(cd cms && atlas migrate apply --env local)

# Jobsimulation schema
(cd jobsimulation && atlas migrate apply --env local)

# Skiller schema
(cd skiller && atlas migrate apply --env local)
```

*Expected*: Each command shows applied migrations or "No migration files to apply".

### If Migrations Fail

Common issues:

**Schema already exists**: The migration may have been partially applied.
```bash
# Check current schema state
docker exec ant-rosetta-postgresql-1 psql -U postgres -c "\dn"
```

**Connection refused**: PostgreSQL not ready yet.
```bash
# Wait and retry
docker exec ant-rosetta-postgresql-1 pg_isready -U postgres
```

**Dirty migration state**: Reset migration history (use with caution).
```bash
# Check atlas migration status
cd [service]
atlas migrate status --env local
```

---

## 5. Rebuild Docker Images

If Dockerfiles or Go code changed, rebuild the affected services.

### Rebuild Specific Service

```bash
cd anthropos-dev/platform
docker compose -p ant-rosetta build backend  # or cms, jobsimulation, etc.
```

### Rebuild All Services

```bash
docker compose -p ant-rosetta build
```

*Note*: This can take several minutes as it pulls and compiles code.

### Force Fresh Build (no cache)

```bash
docker compose -p ant-rosetta build --no-cache
```

---

## 6. Verify Updated Environment

After updating, verify everything works.

### Start Services

```bash
cd anthropos-dev/platform

# Start infrastructure
docker compose -p ant-rosetta up -d postgresql redis

# Start backend
docker compose -p ant-rosetta --profile graphql up -d
```

### Health Checks

```bash
# Check all containers are healthy
docker ps --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}"

# Test backend
curl -s http://localhost:8082/health || echo "Backend not responding"

# Test GraphQL
curl -s http://localhost:5050/health || echo "GraphQL not responding"
```

### Start and Test Frontend

```bash
cd anthropos-dev/next-web-app
pnpm dev:web
```

Open http://localhost:3000 and verify the application loads.

---

## 7. Quick Update Scenarios

### Scenario: Daily Sync (Quick)

Minimal update when starting work:

```bash
cd anthropos-dev

# Pull main repos
(cd platform && git pull origin main)
(cd backend && git pull origin main)
(cd next-web-app && git pull origin main)

# Update frontend deps
(cd next-web-app && pnpm install)

# Restart services
cd platform
docker compose -p ant-rosetta up -d
```

### Scenario: Weekly Sync (Full)

Comprehensive update:

```bash
cd anthropos-dev

# Stop everything
cd platform && docker compose -p ant-rosetta down && cd ..

# Pull all repos
for repo in platform backend cms jobsimulation skiller next-web-app studio-desk studio-room cms/studio; do
  (cd "$repo" 2>/dev/null && git pull origin main) || true
done

# Update all dependencies
(cd next-web-app && pnpm install)
(cd studio-desk && npm install)
(cd studio-room && pip3 install -r requirements.txt --upgrade)

# Apply migrations
cd platform && docker compose -p ant-rosetta up -d postgresql && sleep 5 && cd ..
(cd backend && atlas migrate apply --env local)
(cd cms && atlas migrate apply --env local)
(cd jobsimulation && atlas migrate apply --env local)
(cd skiller && atlas migrate apply --env local)

# Rebuild and start
cd platform
docker compose -p ant-rosetta --profile graphql up -d --build
```

### Scenario: After Major Release

When significant changes are announced:

```bash
cd anthropos-dev

# Full stop
cd platform && docker compose -p ant-rosetta down -v && cd ..

# Fresh pull all repos
for repo in platform backend cms jobsimulation skiller next-web-app studio-desk studio-room; do
  (cd "$repo" 2>/dev/null && git fetch origin && git reset --hard origin/main) || true
done

# Clean and reinstall deps
(cd next-web-app && rm -rf node_modules && pnpm install)
(cd studio-desk && rm -rf node_modules && npm install)

# Fresh database
cd platform && docker compose -p ant-rosetta up -d postgresql redis && sleep 10 && cd ..
(cd backend && atlas migrate apply --env local)
(cd cms && atlas migrate apply --env local)
(cd jobsimulation && atlas migrate apply --env local)
(cd skiller && atlas migrate apply --env local)

# Fresh build
cd platform && docker compose -p ant-rosetta --profile graphql up -d --build
```

---

## 8. Troubleshooting

### "Package not found" After Pull

Dependencies changed but weren't installed.

```bash
# Frontend
cd next-web-app && pnpm install

# Studio-Desk
cd studio-desk && npm install
```

### Migration Conflicts

Schema out of sync with code.

```bash
# Check migration status
cd [service]
atlas migrate status --env local

# If stuck, may need to reset (WARNING: data loss)
docker compose -p ant-rosetta down -v
# Then re-apply all migrations
```

### Docker Build Errors After Update

Image cache may be stale.

```bash
# Clean Docker cache
docker system prune -f

# Rebuild without cache
docker compose -p ant-rosetta build --no-cache
```

### Git Conflicts During Pull

Local changes conflict with remote.

```bash
# See what's conflicting
git status

# Option 1: Keep your changes
git stash
git pull origin main
git stash pop

# Option 2: Discard local changes
git checkout -- .
git pull origin main
```

### Frontend Compilation Errors After Update

TypeScript types may be out of sync.

```bash
cd next-web-app
pnpm clean  # Clean build artifacts
pnpm install  # Reinstall deps
pnpm dev:web  # Try again
```

---

## 9. Maintenance Guidelines

This guide follows the same iterative improvement pattern as other ops guides:

- **Document Issues**: When you discover new update scenarios, add them.
- **Keep Scripts Updated**: If repo structure changes, update the commands.
- **Track Patterns**: Common issues should become troubleshooting entries.

### Related Guides

- [Platform Setup](../platform-setup/setup_guide.md) - Initial environment build
- [Platform Run](../platform-run/run_guide.md) - Starting the platform
